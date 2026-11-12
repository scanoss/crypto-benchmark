// Package groundtruth parses groundtruth/<suite>/<entry-id>.yaml files.
package groundtruth

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"

	"github.com/scanoss/crypto-benchmark/internal/check"
)

const SchemaVersion = "1"

type ScopeKind string

const (
	Exhaustive ScopeKind = "exhaustive" // the whole entry was annotated
	Files      ScopeKind = "files"      // only Scope.Files were annotated
	Pool       ScopeKind = "pool"       // only pooled tool candidates; recall is relative
)

// AssetType is CycloneDX 1.7 cryptoProperties.assetType.
type AssetType string

const (
	Algorithm             AssetType = "algorithm"
	Certificate           AssetType = "certificate"
	Protocol              AssetType = "protocol"
	RelatedCryptoMaterial AssetType = "related-crypto-material"
)

// AssetTypes lists every AssetType.
var AssetTypes = []AssetType{Algorithm, Certificate, Protocol, RelatedCryptoMaterial}

// Primitive is CycloneDX 1.7 algorithmProperties.primitive.
type Primitive string

// Primitives lists every Primitive.
var Primitives = []Primitive{
	"drbg", "mac", "block-cipher", "stream-cipher", "signature", "hash", "pke", "xof",
	"kdf", "key-agree", "kem", "ae", "combiner", "key-wrap", "other", "unknown",
}

// MaterialType is CycloneDX 1.7 relatedCryptoMaterialProperties.type.
type MaterialType string

// MaterialTypes lists every MaterialType.
var MaterialTypes = []MaterialType{
	"private-key", "public-key", "secret-key", "key", "ciphertext", "signature", "digest",
	"initialization-vector", "nonce", "seed", "salt", "shared-secret", "tag", "additional-data",
	"password", "credential", "token", "other", "unknown",
}

// Depth counts hops from project code: 0 in the project, 1 one hop, 2 two or more.
type Depth int

type Scope struct {
	Kind  ScopeKind
	Files []string // set only when Kind == Files
	// Components, when set, limits the dependency side of the scope to these
	// package URLs: predictions in any other component are ignored. Project
	// code stays in scope. Suite B uses it to leave out the internals of a
	// third-party crypto library that a fixture library calls.
	Components []string
}

// Origin is Project or Dependency.
type Origin interface{ isOrigin() }

type Project struct{}

type Dependency struct {
	Component string // package URL
}

func (Project) isOrigin()    {}
func (Dependency) isOrigin() {}

// Reachability is Unknown or Known; only Known carries a depth.
type Reachability interface{ isReachability() }

type Unknown struct{}

type Known struct {
	Reachable bool
	Depth     Depth
	// Distance is the exact dependency distance (0 project, 1 direct, 3 for
	// app -> A -> B -> C); -1 when the file does not state it. Depth is
	// Distance capped at 2.
	Distance int
}

func (Unknown) isReachability() {}
func (Known) isReachability()   {}

type Lines struct{ Start, End int }

type Parameters struct {
	Mode    string
	Padding string
	KeySize int // 0 when absent
	Curve   string
	Hash    string // the digest a signature or MAC computes
}

type Asset struct {
	ID    string
	File  string
	Lines Lines
	// Anchors are extra line ranges of the same asset: the call that names
	// the algorithm for the same object in the same function, when Lines is
	// the call that performs the operation. A prediction matches the asset
	// when it overlaps Lines or any anchor. Empty for most assets.
	Anchors      []Lines
	Type         AssetType
	Primitive    Primitive    // empty when absent; only with Type == Algorithm
	MaterialType MaterialType // empty when absent; only with Type == RelatedCryptoMaterial
	Algorithm    string
	Parameters   Parameters
	Origin       Origin
	Reachability Reachability
	Path         []string
	// Object is shared by the assets that act on one crypto object; empty
	// when the asset is its own object.
	Object string
	Notes  string
}

type File struct {
	Entry  string
	Scope  Scope
	Assets []Asset
}

type rawFile struct {
	SchemaVersion string     `yaml:"schema_version"`
	Entry         string     `yaml:"entry"`
	Scope         rawScope   `yaml:"scope"`
	Assets        []rawAsset `yaml:"assets"`
}

type rawScope struct {
	Kind       string   `yaml:"kind"`
	Files      []string `yaml:"files"`
	Components []string `yaml:"components"`
}

type rawAsset struct {
	ID           string        `yaml:"id"`
	File         string        `yaml:"file"`
	Lines        []int         `yaml:"lines"`
	Anchors      [][]int       `yaml:"anchors"`
	AssetType    string        `yaml:"asset_type"`
	Primitive    string        `yaml:"primitive"`
	MaterialType string        `yaml:"material_type"`
	Algorithm    string        `yaml:"algorithm"`
	Parameters   rawParameters `yaml:"parameters"`
	Origin       string        `yaml:"origin"`
	Component    string        `yaml:"component"`
	Reachable    string        `yaml:"reachable"`
	Depth        *int          `yaml:"depth"`
	Distance     *int          `yaml:"distance"`
	Path         []string      `yaml:"path"`
	Object       string        `yaml:"object"`
	Notes        string        `yaml:"notes"`
}

type rawParameters struct {
	Mode    string `yaml:"mode"`
	Padding string `yaml:"padding"`
	KeySize *int   `yaml:"key_size"`
	Curve   string `yaml:"curve"`
	Hash    string `yaml:"hash"`
}

// Parse decodes and checks one ground-truth file, recording problems in r.
// ok is false when the file could not be decoded at all.
func Parse(file string, data []byte, r *check.Report) (f File, ok bool) {
	var raw rawFile
	if err := check.DecodeStrict(data, &raw); err != nil {
		r.Errorf(file, "", "%v", err)
		return File{}, false
	}
	if raw.SchemaVersion != SchemaVersion {
		r.Errorf(file, "schema_version", "got %q, want %q", raw.SchemaVersion, SchemaVersion)
	}
	f.Entry = raw.Entry
	f.Scope = parseScope(raw.Scope, file, r)
	ids := map[string]bool{}
	for i, ra := range raw.Assets {
		where := fmt.Sprintf("assets[%d]", i)
		if ids[ra.ID] {
			r.Errorf(file, where+".id", "duplicate asset id %q", ra.ID)
		}
		ids[ra.ID] = true
		f.Assets = append(f.Assets, parseAsset(ra, f.Scope, file, where, r))
	}
	checkObjects(f.Assets, file, r)
	return f, true
}

// checkObjects enforces the annotation guide's object rules: an object has
// two or more assets, all of one asset type, and no related crypto material
// with a material type, which is always its own object.
func checkObjects(assets []Asset, file string, r *check.Report) {
	first := map[string]int{}
	count := map[string]int{}
	for i, a := range assets {
		if a.Object == "" {
			continue
		}
		where := fmt.Sprintf("assets[%d].object", i)
		count[a.Object]++
		if a.MaterialType != "" {
			r.Errorf(file, where, "%q is set on related crypto material with material_type %s, which is its own object", a.Object, a.MaterialType)
		}
		j, seen := first[a.Object]
		switch {
		case !seen:
			first[a.Object] = i
		case assets[j].Type != a.Type:
			r.Errorf(file, where, "object %q mixes asset types %s (%s) and %s", a.Object, assets[j].Type, assets[j].ID, a.Type)
		}
	}
	for i, a := range assets {
		if a.Object != "" && count[a.Object] == 1 {
			r.Errorf(file, fmt.Sprintf("assets[%d].object", i), "object %q has only one asset; leave object out", a.Object)
		}
	}
}

func parseScope(rs rawScope, file string, r *check.Report) Scope {
	kind, err := check.OneOf(rs.Kind, Exhaustive, Files, Pool)
	if err != nil {
		r.Errorf(file, "scope.kind", "%v", err)
	}
	switch {
	case kind == Files && len(rs.Files) == 0:
		r.Errorf(file, "scope.files", "required when scope.kind is files")
	case kind != Files && len(rs.Files) > 0:
		r.Errorf(file, "scope.files", "only allowed when scope.kind is files")
	}
	for i, p := range rs.Files {
		if !filepath.IsLocal(p) {
			r.Errorf(file, fmt.Sprintf("scope.files[%d]", i), "%q must be relative to the entry root", p)
		}
	}
	seen := map[string]bool{}
	for i, c := range rs.Components {
		where := fmt.Sprintf("scope.components[%d]", i)
		switch {
		case !strings.HasPrefix(c, "pkg:"):
			r.Errorf(file, where, "%q is not a package URL", c)
		case seen[c]:
			r.Errorf(file, where, "duplicate component %q", c)
		}
		seen[c] = true
	}
	return Scope{Kind: kind, Files: rs.Files, Components: rs.Components}
}

// parseAnchors checks an asset's extra line ranges: each is [start, end]
// with 1 <= start <= end, overlaps neither lines nor another anchor, and is
// allowed only on algorithm assets, the only kind whose operation call and
// algorithm-naming call the annotation guide lets differ.
func parseAnchors(raw [][]int, a Asset, fail func(i int, format string, args ...any)) []Lines {
	var out []Lines
	overlaps := func(x, y Lines) bool { return x.Start <= y.End && y.Start <= x.End }
	for i, r := range raw {
		switch {
		case a.Type != Algorithm:
			fail(i, "only allowed when asset_type is algorithm")
			continue
		case len(r) != 2:
			fail(i, "must be [start, end], got %v", r)
			continue
		case r[0] < 1 || r[0] > r[1]:
			fail(i, "need 1 <= start <= end, got %v", r)
			continue
		}
		l := Lines{r[0], r[1]}
		if a.Lines != (Lines{}) && overlaps(l, a.Lines) {
			fail(i, "%v overlaps lines [%d, %d]", r, a.Lines.Start, a.Lines.End)
			continue
		}
		if j := slices.IndexFunc(out, func(o Lines) bool { return overlaps(o, l) }); j >= 0 {
			fail(i, "%v overlaps another anchor [%d, %d]", r, out[j].Start, out[j].End)
			continue
		}
		out = append(out, l)
	}
	return out
}

func parseAsset(ra rawAsset, scope Scope, file, where string, r *check.Report) Asset {
	fail := func(field, format string, args ...any) {
		r.Errorf(file, where+"."+field, format, args...)
	}
	a := Asset{ID: ra.ID, File: ra.File, Algorithm: ra.Algorithm, Path: ra.Path, Object: ra.Object, Notes: ra.Notes}
	if ra.ID == "" {
		fail("id", "required")
	}
	switch {
	case !filepath.IsLocal(ra.File):
		fail("file", "%q must be relative to the entry root", ra.File)
	case scope.Kind == Files && !slices.Contains(scope.Files, ra.File):
		fail("file", "%q is outside scope.files", ra.File)
	}
	switch {
	case len(ra.Lines) != 2:
		fail("lines", "must be [start, end], got %v", ra.Lines)
	case ra.Lines[0] < 1 || ra.Lines[0] > ra.Lines[1]:
		fail("lines", "need 1 <= start <= end, got %v", ra.Lines)
	default:
		a.Lines = Lines{ra.Lines[0], ra.Lines[1]}
	}

	var err error
	if a.Type, err = check.OneOf(ra.AssetType, AssetTypes...); err != nil {
		fail("asset_type", "%v", err)
	}
	a.Anchors = parseAnchors(ra.Anchors, a, func(i int, format string, args ...any) {
		fail(fmt.Sprintf("anchors[%d]", i), format, args...)
	})
	if ra.Primitive != "" {
		if a.Primitive, err = check.OneOf(ra.Primitive, Primitives...); err != nil {
			fail("primitive", "%v", err)
		}
		if a.Type != Algorithm {
			fail("primitive", "only allowed when asset_type is algorithm")
		}
	}
	if ra.MaterialType != "" {
		if a.MaterialType, err = check.OneOf(ra.MaterialType, MaterialTypes...); err != nil {
			fail("material_type", "%v", err)
		}
		if a.Type != RelatedCryptoMaterial {
			fail("material_type", "only allowed when asset_type is related-crypto-material")
		}
	}
	a.Parameters = Parameters{Mode: ra.Parameters.Mode, Padding: ra.Parameters.Padding, Curve: ra.Parameters.Curve, Hash: ra.Parameters.Hash}
	if ks := ra.Parameters.KeySize; ks != nil {
		if *ks <= 0 {
			fail("parameters.key_size", "must be > 0, got %d", *ks)
		}
		a.Parameters.KeySize = *ks
	}

	switch ra.Origin {
	case "project":
		a.Origin = Project{}
		if ra.Component != "" {
			fail("component", "only allowed when origin is dependency")
		}
	case "dependency":
		a.Origin = Dependency{Component: ra.Component}
		if ra.Component == "" {
			fail("component", "required when origin is dependency")
		}
		if len(scope.Components) > 0 && !slices.Contains(scope.Components, ra.Component) {
			fail("component", "%q is not in scope.components", ra.Component)
		}
	default:
		fail("origin", "%q is not one of project|dependency", ra.Origin)
	}

	switch ra.Reachable {
	case "unknown":
		a.Reachability = Unknown{}
		if ra.Depth != nil {
			fail("depth", "not allowed when reachable is unknown")
		}
		if ra.Distance != nil {
			fail("distance", "not allowed when reachable is unknown")
		}
	case "true", "false":
		known := Known{Reachable: ra.Reachable == "true", Distance: -1}
		switch {
		case ra.Depth == nil:
			fail("depth", "required when reachable is %s", ra.Reachable)
		case *ra.Depth < 0 || *ra.Depth > 2:
			fail("depth", "must be 0, 1 or 2 (2 means two or more hops), got %d", *ra.Depth)
		default:
			known.Depth = Depth(*ra.Depth)
		}
		if d := ra.Distance; d != nil {
			_, project := a.Origin.(Project)
			switch {
			case *d < 0:
				fail("distance", "must be >= 0, got %d", *d)
			case project != (*d == 0):
				fail("distance", "must be 0 exactly when origin is project, got %d", *d)
			case ra.Depth != nil && min(*d, 2) != *ra.Depth:
				fail("distance", "%d disagrees with depth %d (depth is distance capped at 2)", *d, *ra.Depth)
			default:
				known.Distance = *d
			}
		}
		a.Reachability = known
	default:
		fail("reachable", "%q is not one of true|false|unknown", ra.Reachable)
	}
	return a
}
