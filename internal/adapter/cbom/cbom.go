// Package cbom converts a CycloneDX 1.6 or 1.7 CBOM, as written by any tool,
// into a benchmark prediction file. The mapping is documented in README.md
// next to this file.
package cbom

import (
	"encoding/json"
	"fmt"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/scanoss/crypto-benchmark/internal/algo"
	"github.com/scanoss/crypto-benchmark/internal/groundtruth"
	"github.com/scanoss/crypto-benchmark/internal/prediction"
)

// specVersions are the CycloneDX versions that define cryptographic-asset
// components with cryptoProperties and evidence.occurrences.
var specVersions = []string{"1.6", "1.7"}

// Options says how to read one CBOM.
type Options struct {
	Tool, ToolVersion, Entry string
	// SourceRoot is the absolute directory the tool scanned. An absolute
	// location under it becomes relative to it; without it, absolute
	// locations cannot be matched.
	SourceRoot string
	// Cdxgen reads the locations cdxgen keeps in cdx:*crypto:sourceLocation
	// properties instead of evidence.occurrences, and the source directory
	// of a Go module replaced by a local path.
	Cdxgen bool
}

// Result is the prediction file plus one note per component or location
// left out.
type Result struct {
	File  prediction.File
	Notes []string
}

type bom struct {
	BOMFormat    string       `json:"bomFormat"`
	SpecVersion  string       `json:"specVersion"`
	Metadata     metadata     `json:"metadata"`
	Components   []component  `json:"components"`
	Dependencies []dependency `json:"dependencies"`
}

type metadata struct {
	Component *component `json:"component"`
}

type component struct {
	Type             string            `json:"type"`
	Name             string            `json:"name"`
	BOMRef           string            `json:"bom-ref"`
	PURL             string            `json:"purl"`
	CryptoProperties *cryptoProperties `json:"cryptoProperties"`
	Evidence         *struct {
		Occurrences []occurrence `json:"occurrences"`
	} `json:"evidence"`
	Properties []property  `json:"properties"`
	Components []component `json:"components"`
}

type cryptoProperties struct {
	AssetType           string `json:"assetType"`
	AlgorithmProperties *struct {
		Primitive              string `json:"primitive"`
		Mode                   string `json:"mode"`
		Padding                string `json:"padding"`
		Curve                  string `json:"curve"`
		EllipticCurve          string `json:"ellipticCurve"`
		ParameterSetIdentifier string `json:"parameterSetIdentifier"`
	} `json:"algorithmProperties"`
	RelatedCryptoMaterialProperties *struct {
		Type string `json:"type"`
		Size int    `json:"size"`
	} `json:"relatedCryptoMaterialProperties"`
}

type occurrence struct {
	Location string `json:"location"`
	Line     int    `json:"line"`
}

type property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type dependency struct {
	Ref       string   `json:"ref"`
	DependsOn []string `json:"dependsOn"`
	Provides  []string `json:"provides"`
}

// location is one place a component was found: a path as the tool wrote
// it and a line, 0 when the tool gave none.
type location struct {
	path string
	line int
}

// cdxgenLocation matches the properties in which cdxgen records crypto
// source locations: cdx:crypto:sourceLocation for JavaScript, and
// cdx:golem:..., cdx:rusi:... and the like for its per-language analyzers.
var cdxgenLocation = regexp.MustCompile(`^cdx:([a-z]+:)?crypto:sourceLocation$`)

// Adapt converts a CBOM into the prediction file for opts.Entry.
func Adapt(data []byte, opts Options) (Result, error) {
	var b bom
	if err := json.Unmarshal(data, &b); err != nil {
		return Result{}, fmt.Errorf("cbom: %w", err)
	}
	if b.BOMFormat != "CycloneDX" {
		return Result{}, fmt.Errorf("cbom: bomFormat is %q, want CycloneDX", b.BOMFormat)
	}
	if !slices.Contains(specVersions, b.SpecVersion) {
		return Result{}, fmt.Errorf("cbom: specVersion %q is not one of %s", b.SpecVersion, strings.Join(specVersions, "|"))
	}
	a := newAdapter(b, opts)
	walk(b.Components, "", a.convert)
	res := Result{
		File: prediction.File{
			Tool:   prediction.Tool{Name: opts.Tool, Version: opts.ToolVersion, Config: map[string]string{}},
			Entry:  opts.Entry,
			Assets: a.assets,
			Metadata: map[string]string{
				"components_without_location": strconv.Itoa(a.unlocated),
				"locations_dropped":           strconv.Itoa(a.dropped),
			},
		},
		Notes: a.notes,
	}
	if opts.Cdxgen {
		res.File.Tool.Config["cdxgen_locations"] = "true"
	}
	return res, nil
}

// walk calls f on every component, depth first, with the package URL of the
// nearest library component it is nested under.
func walk(cs []component, parent string, f func(c component, parent string)) {
	for _, c := range cs {
		f(c, parent)
		inner := parent
		if isLibrary(c) {
			inner = c.PURL
		}
		walk(c.Components, inner, f)
	}
}

func isLibrary(c component) bool {
	return (c.Type == "library" || c.Type == "framework") && c.PURL != ""
}

type adapter struct {
	opts      Options
	libraries libraries
	links     map[string]linkers
	assets    []prediction.Asset
	notes     []string
	unlocated int
	dropped   int
}

// linkers are the components whose dependencies entry provides or depends
// on a crypto asset: library package URLs, and whether the project does.
type linkers struct {
	libraries []string
	project   bool
}

func newAdapter(b bom, opts Options) *adapter {
	a := &adapter{opts: opts, libraries: libraries{roots: map[string][]string{}, npm: map[string][]string{}}, links: map[string]linkers{}}
	refs := map[string]component{}
	walk(b.Components, "", func(c component, _ string) {
		refs[c.BOMRef] = c
		if isLibrary(c) {
			a.libraries.add(c, opts)
		}
	})
	project := ""
	if b.Metadata.Component != nil {
		project = b.Metadata.Component.BOMRef
	}
	for _, d := range b.Dependencies {
		lib, isLib := refs[d.Ref]
		for _, target := range append(slices.Clone(d.Provides), d.DependsOn...) {
			l := a.links[target]
			switch {
			case d.Ref == project && project != "":
				l.project = true
			case isLib && isLibrary(lib) && !slices.Contains(l.libraries, lib.PURL):
				l.libraries = append(l.libraries, lib.PURL)
			}
			a.links[target] = l
		}
	}
	return a
}

func (a *adapter) note(c component, format string, args ...any) {
	a.notes = append(a.notes, fmt.Sprintf("%s (%s): %s", c.Name, c.BOMRef, fmt.Sprintf(format, args...)))
}

func (a *adapter) convert(c component, parent string) {
	if c.Type != "cryptographic-asset" {
		return
	}
	if c.CryptoProperties == nil {
		a.note(c, "no cryptoProperties")
		return
	}
	typ := groundtruth.AssetType(strings.ToLower(strings.TrimSpace(c.CryptoProperties.AssetType)))
	if !slices.Contains(groundtruth.AssetTypes, typ) {
		a.note(c, "assetType %q is not a CycloneDX asset type", c.CryptoProperties.AssetType)
		return
	}
	template := prediction.Asset{
		Type:      typ,
		Algorithm: strings.TrimSpace(c.Name),
		Reachable: prediction.NoAnswer,
	}
	template.Primitive, template.Params = parameters(typ, *c.CryptoProperties)
	template.MaterialType = materialType(typ, *c.CryptoProperties)

	locs := a.locations(c)
	if len(locs) == 0 {
		a.unlocated++
		a.note(c, "no occurrence location; it cannot be matched")
		return
	}
	located := false
	for _, l := range locs {
		p, why := a.place(c, parent, l)
		if why != "" {
			a.dropped++
			a.note(c, "%s:%d %s", l.path, l.line, why)
			continue
		}
		asset := template
		asset.File, asset.Lines, asset.Origin = p.file, groundtruth.Lines{Start: l.line, End: l.line}, p.origin
		a.assets = append(a.assets, asset)
		located = true
	}
	if !located {
		a.unlocated++
	}
}

func parameters(typ groundtruth.AssetType, cp cryptoProperties) (groundtruth.Primitive, groundtruth.Parameters) {
	var prim groundtruth.Primitive
	var p groundtruth.Parameters
	switch {
	case typ == groundtruth.Algorithm && cp.AlgorithmProperties != nil:
		ap := cp.AlgorithmProperties
		if v := groundtruth.Primitive(strings.ToLower(strings.TrimSpace(ap.Primitive))); slices.Contains(groundtruth.Primitives, v) {
			prim = v
		}
		p.Mode, p.Padding = strings.TrimSpace(ap.Mode), strings.TrimSpace(ap.Padding)
		// 1.7's ellipticCurve is namespaced (nist/P-256); 1.6's curve is free text.
		p.Curve = strings.TrimSpace(ap.Curve)
		if ec := strings.TrimSpace(ap.EllipticCurve); ec != "" {
			p.Curve = ec[strings.LastIndex(ec, "/")+1:]
		}
		psi := strings.TrimSpace(ap.ParameterSetIdentifier)
		if n, err := strconv.Atoi(psi); err == nil && n > 0 {
			p.KeySize = n
		} else if c, ok := algo.Curve(psi); ok && p.Curve == "" {
			p.Curve = c
		}
	case typ == groundtruth.RelatedCryptoMaterial && cp.RelatedCryptoMaterialProperties != nil:
		p.KeySize = max(cp.RelatedCryptoMaterialProperties.Size, 0)
	}
	return prim, p
}

// materialType is relatedCryptoMaterialProperties.type, lowercased, when it
// is a CycloneDX 1.7 material type.
func materialType(typ groundtruth.AssetType, cp cryptoProperties) groundtruth.MaterialType {
	if typ != groundtruth.RelatedCryptoMaterial || cp.RelatedCryptoMaterialProperties == nil {
		return ""
	}
	mt := groundtruth.MaterialType(strings.ToLower(strings.TrimSpace(cp.RelatedCryptoMaterialProperties.Type)))
	if !slices.Contains(groundtruth.MaterialTypes, mt) {
		return ""
	}
	return mt
}

// locations are the component's evidence.occurrences, or, with Cdxgen and
// no occurrences, its cdxgen source location properties.
func (a *adapter) locations(c component) []location {
	var locs []location
	if c.Evidence != nil {
		for _, o := range c.Evidence.Occurrences {
			if o.Location != "" {
				locs = append(locs, location{o.Location, o.Line})
			}
		}
	}
	if len(locs) > 0 || !a.opts.Cdxgen {
		return locs
	}
	for _, p := range c.Properties {
		if cdxgenLocation.MatchString(p.Name) {
			if l, ok := parseCdxgenLocation(p.Value); ok {
				locs = append(locs, l)
			}
		}
	}
	return locs
}

// parseCdxgenLocation reads FILE:LINE:COLUMN, FILE:LINE or FILE#LINE.
func parseCdxgenLocation(v string) (location, bool) {
	if file, line, ok := strings.Cut(v, "#"); ok {
		n, err := strconv.Atoi(line)
		return location{file, n}, err == nil && file != ""
	}
	file, last, ok := cutLastInt(v)
	if !ok {
		return location{}, false
	}
	if f, line, ok := cutLastInt(file); ok {
		return location{f, line}, f != ""
	}
	return location{file, last}, file != ""
}

func cutLastInt(s string) (string, int, bool) {
	i := strings.LastIndex(s, ":")
	if i < 0 {
		return "", 0, false
	}
	n, err := strconv.Atoi(s[i+1:])
	return s[:i], n, err == nil
}

// placement is where a located component lands: its file, relative to the
// entry root or to the component's source root, and its origin.
type placement struct {
	file   string
	origin groundtruth.Origin
}

// place resolves one location, or says why it cannot be matched. A path
// inside a dependency's source root decides the origin first; then a
// library the component is nested under; then the one library the
// dependencies section links it to; otherwise the project.
func (a *adapter) place(c component, parent string, l location) (placement, string) {
	if l.line < 1 {
		return placement{}, "has no line"
	}
	p := filepath.ToSlash(l.path)
	if path.IsAbs(p) {
		if a.opts.SourceRoot == "" {
			return placement{}, "is absolute and no source root was given"
		}
		rel, err := filepath.Rel(a.opts.SourceRoot, l.path)
		if err != nil {
			return placement{}, "is not under the source root"
		}
		p = filepath.ToSlash(rel)
	}
	p = path.Clean(p)
	if purl, rest, ok := a.libraries.owner(p); ok {
		return placement{rest, groundtruth.Dependency{Component: purl}}, ""
	}
	var origin groundtruth.Origin = groundtruth.Project{}
	if parent != "" {
		origin = groundtruth.Dependency{Component: parent}
	} else if l := a.links[c.BOMRef]; len(l.libraries) == 1 && !l.project {
		origin = groundtruth.Dependency{Component: l.libraries[0]}
	}
	if !filepath.IsLocal(p) {
		return placement{}, "is outside the source root and no dependency's source root"
	}
	return placement{p, origin}, ""
}

// libraries maps a path relative to the source root to the library whose
// source it holds.
type libraries struct {
	// roots are source directories relative to the source root, for Go
	// modules cdxgen reports as replaced by a local directory.
	roots map[string][]string
	// npm maps a package name, @scope/name or name, to its package URLs;
	// its source is node_modules/<name>.
	npm map[string][]string
}

func (ls libraries) add(c component, opts Options) {
	if name, ok := npmName(c.PURL); ok {
		ls.npm[name] = append(ls.npm[name], c.PURL)
	}
	if !opts.Cdxgen || opts.SourceRoot == "" {
		return
	}
	var localDir, goMod string
	for _, p := range c.Properties {
		switch p.Name {
		case "cdx:go:local_dir":
			localDir = p.Value
		case "internal:SrcGoMod":
			goMod = p.Value
		}
	}
	if localDir == "" || !filepath.IsAbs(goMod) {
		return
	}
	if rel, err := filepath.Rel(opts.SourceRoot, filepath.Dir(goMod)); err == nil {
		root := path.Clean(filepath.ToSlash(rel))
		ls.roots[root] = append(ls.roots[root], c.PURL)
	}
}

// owner returns the library whose source root holds p, and p relative to
// that root, when exactly one library claims the longest matching root.
func (ls libraries) owner(p string) (string, string, bool) {
	best, rest := "", ""
	for root := range ls.roots {
		if r, ok := strings.CutPrefix(p, root+"/"); ok && len(root) > len(best) {
			best, rest = root, r
		}
	}
	if best != "" {
		return single(ls.roots[best], rest)
	}
	i := strings.LastIndex("/"+p, "/node_modules/")
	if i < 0 {
		return "", "", false
	}
	segs := strings.SplitN(p[i+len("node_modules/"):], "/", 3)
	n := 1
	if strings.HasPrefix(segs[0], "@") {
		n = 2
	}
	if len(segs) <= n {
		return "", "", false
	}
	return single(ls.npm[strings.Join(segs[:n], "/")], strings.Join(segs[n:], "/"))
}

func single(purls []string, rest string) (string, string, bool) {
	if len(purls) != 1 {
		return "", "", false
	}
	return purls[0], rest, true
}

// npmName returns the package name in an npm package URL, @scope/name or name.
func npmName(purl string) (string, bool) {
	rest, ok := strings.CutPrefix(purl, "pkg:npm/")
	if !ok {
		return "", false
	}
	rest, _, _ = strings.Cut(rest, "#")
	rest, _, _ = strings.Cut(rest, "?")
	if i := strings.LastIndex(rest, "@"); i > 0 {
		rest = rest[:i]
	}
	return strings.ReplaceAll(rest, "%40", "@"), true
}
