// Package prediction parses and writes predictions/<tool>/<entry-id>.json:
// one tool's crypto assets for one corpus entry, normalized to the
// ground-truth vocabulary so every tool is scored by the same rules.
package prediction

import (
	"encoding/json"
	"fmt"
	"path"
	"path/filepath"

	"github.com/scanoss/crypto-benchmark/internal/check"
	"github.com/scanoss/crypto-benchmark/internal/groundtruth"
)

const SchemaVersion = "1"

// Answer is a tool's reachability answer for one asset.
type Answer string

const (
	Reachable   Answer = "true"
	Unreachable Answer = "false"
	NoAnswer    Answer = "unknown"
)

type Tool struct {
	Name    string
	Version string
	Config  map[string]string
}

type Asset struct {
	File      string // clean, slash-separated, relative to the entry or component root
	Lines     groundtruth.Lines
	Type      groundtruth.AssetType
	Primitive groundtruth.Primitive // empty when absent
	// MaterialType is empty when absent; only with Type == RelatedCryptoMaterial.
	MaterialType groundtruth.MaterialType
	Algorithm    string
	Params       groundtruth.Parameters
	Origin       groundtruth.Origin
	Reachable    Answer
	Depth        *groundtruth.Depth // nil when the tool gives none; never set with NoAnswer
	// Distance is the exact dependency distance behind Depth (3 for app ->
	// A -> B -> C); nil when the tool gives none. Never set without Depth.
	Distance *int
}

type File struct {
	Tool  Tool
	Entry string
	// ReachabilityReported is false when the tool never reports reachability;
	// every asset then answers NoAnswer and reachability metrics are n/a.
	ReachabilityReported bool
	Assets               []Asset
	// Metadata is what the adapter records about the conversion, such as
	// tool findings it could not turn into assets. The scorer ignores it.
	Metadata map[string]string
}

type rawFile struct {
	SchemaVersion        string            `json:"schema_version"`
	Tool                 rawTool           `json:"tool"`
	Entry                string            `json:"entry"`
	ReachabilityReported *bool             `json:"reachability_reported"`
	Assets               []rawAsset        `json:"assets"`
	Metadata             map[string]string `json:"metadata,omitempty"`
}

type rawTool struct {
	Name    string            `json:"name"`
	Version string            `json:"version"`
	Config  map[string]string `json:"config,omitempty"`
}

type rawAsset struct {
	File         string         `json:"file"`
	Lines        []int          `json:"lines"`
	AssetType    string         `json:"asset_type"`
	Primitive    string         `json:"primitive,omitempty"`
	MaterialType string         `json:"material_type,omitempty"`
	Algorithm    string         `json:"algorithm,omitempty"`
	Parameters   *rawParameters `json:"parameters,omitempty"`
	Origin       string         `json:"origin"`
	Component    string         `json:"component,omitempty"`
	Reachable    string         `json:"reachable"`
	Depth        *int           `json:"depth,omitempty"`
	Distance     *int           `json:"distance,omitempty"`
}

type rawParameters struct {
	Mode    string `json:"mode,omitempty"`
	Padding string `json:"padding,omitempty"`
	KeySize *int   `json:"key_size,omitempty"`
	Curve   string `json:"curve,omitempty"`
	Hash    string `json:"hash,omitempty"`
}

// NormalizePath returns p cleaned and slash-separated, the form file paths are compared in.
func NormalizePath(p string) string { return path.Clean(filepath.ToSlash(p)) }

// Parse decodes and checks one prediction file, recording problems in r.
// ok is false when the file could not be decoded at all.
func Parse(file string, data []byte, r *check.Report) (f File, ok bool) {
	var raw rawFile
	if err := check.DecodeStrictJSON(data, &raw); err != nil {
		r.Errorf(file, "", "%v", err)
		return File{}, false
	}
	if raw.SchemaVersion != SchemaVersion {
		r.Errorf(file, "schema_version", "got %q, want %q", raw.SchemaVersion, SchemaVersion)
	}
	if raw.Tool.Name == "" {
		r.Errorf(file, "tool.name", "required")
	}
	if raw.Entry == "" {
		r.Errorf(file, "entry", "required")
	}
	if raw.ReachabilityReported == nil {
		r.Errorf(file, "reachability_reported", "required")
	}
	f = File{
		Tool:                 Tool{Name: raw.Tool.Name, Version: raw.Tool.Version, Config: raw.Tool.Config},
		Entry:                raw.Entry,
		ReachabilityReported: raw.ReachabilityReported != nil && *raw.ReachabilityReported,
		Metadata:             raw.Metadata,
	}
	// A missing flag is already an error; do not repeat it on every asset.
	answersAllowed := raw.ReachabilityReported == nil || *raw.ReachabilityReported
	for i, ra := range raw.Assets {
		f.Assets = append(f.Assets, parseAsset(ra, answersAllowed, file, fmt.Sprintf("assets[%d]", i), r))
	}
	return f, true
}

func parseAsset(ra rawAsset, reported bool, file, where string, r *check.Report) Asset {
	fail := func(field, format string, args ...any) {
		r.Errorf(file, where+"."+field, format, args...)
	}
	a := Asset{File: NormalizePath(ra.File), Algorithm: ra.Algorithm}
	if !filepath.IsLocal(ra.File) {
		fail("file", "%q must be relative to the entry or component root", ra.File)
	}
	switch {
	case len(ra.Lines) != 2:
		fail("lines", "must be [start, end], got %v", ra.Lines)
	case ra.Lines[0] < 1 || ra.Lines[0] > ra.Lines[1]:
		fail("lines", "need 1 <= start <= end, got %v", ra.Lines)
	default:
		a.Lines = groundtruth.Lines{Start: ra.Lines[0], End: ra.Lines[1]}
	}
	var err error
	if a.Type, err = check.OneOf(ra.AssetType, groundtruth.AssetTypes...); err != nil {
		fail("asset_type", "%v", err)
	}
	if ra.Primitive != "" {
		if a.Primitive, err = check.OneOf(ra.Primitive, groundtruth.Primitives...); err != nil {
			fail("primitive", "%v", err)
		}
		if a.Type != groundtruth.Algorithm {
			fail("primitive", "only allowed when asset_type is algorithm")
		}
	}
	if ra.MaterialType != "" {
		if a.MaterialType, err = check.OneOf(ra.MaterialType, groundtruth.MaterialTypes...); err != nil {
			fail("material_type", "%v", err)
		}
		if a.Type != groundtruth.RelatedCryptoMaterial {
			fail("material_type", "only allowed when asset_type is related-crypto-material")
		}
	}
	if p := ra.Parameters; p != nil {
		a.Params = groundtruth.Parameters{Mode: p.Mode, Padding: p.Padding, Curve: p.Curve, Hash: p.Hash}
		if p.KeySize != nil {
			if *p.KeySize <= 0 {
				fail("parameters.key_size", "must be > 0, got %d", *p.KeySize)
			}
			a.Params.KeySize = *p.KeySize
		}
	}
	switch ra.Origin {
	case "project":
		a.Origin = groundtruth.Project{}
		if ra.Component != "" {
			fail("component", "only allowed when origin is dependency")
		}
	case "dependency":
		a.Origin = groundtruth.Dependency{Component: ra.Component}
		if ra.Component == "" {
			fail("component", "required when origin is dependency")
		}
	default:
		fail("origin", "%q is not one of project|dependency", ra.Origin)
	}
	if a.Reachable, err = check.OneOf(ra.Reachable, Reachable, Unreachable, NoAnswer); err != nil {
		fail("reachable", "%v", err)
	}
	if !reported && a.Reachable != NoAnswer && err == nil {
		fail("reachable", "must be unknown when reachability_reported is false")
	}
	if ra.Depth != nil {
		switch {
		case a.Reachable == NoAnswer:
			fail("depth", "not allowed when reachable is unknown")
		case *ra.Depth < 0 || *ra.Depth > 2:
			fail("depth", "must be 0, 1 or 2 (2 means two or more hops), got %d", *ra.Depth)
		default:
			d := groundtruth.Depth(*ra.Depth)
			a.Depth = &d
		}
	}
	if ra.Distance != nil {
		switch {
		case ra.Depth == nil:
			fail("distance", "only allowed together with depth")
		case *ra.Distance < 0 || min(*ra.Distance, 2) != *ra.Depth:
			fail("distance", "%d disagrees with depth %d (depth is distance capped at 2)", *ra.Distance, *ra.Depth)
		default:
			d := *ra.Distance
			a.Distance = &d
		}
	}
	return a
}

// Marshal encodes f as an indented prediction file.
func Marshal(f File) ([]byte, error) {
	reported := f.ReachabilityReported
	raw := rawFile{
		SchemaVersion:        SchemaVersion,
		Tool:                 rawTool{Name: f.Tool.Name, Version: f.Tool.Version, Config: f.Tool.Config},
		Entry:                f.Entry,
		ReachabilityReported: &reported,
		Assets:               make([]rawAsset, 0, len(f.Assets)),
		Metadata:             f.Metadata,
	}
	for _, a := range f.Assets {
		ra := rawAsset{
			File:         a.File,
			Lines:        []int{a.Lines.Start, a.Lines.End},
			AssetType:    string(a.Type),
			Primitive:    string(a.Primitive),
			MaterialType: string(a.MaterialType),
			Algorithm:    a.Algorithm,
			Reachable:    string(a.Reachable),
		}
		if p := a.Params; p != (groundtruth.Parameters{}) {
			ra.Parameters = &rawParameters{Mode: p.Mode, Padding: p.Padding, Curve: p.Curve, Hash: p.Hash}
			if p.KeySize > 0 {
				ks := p.KeySize
				ra.Parameters.KeySize = &ks
			}
		}
		switch o := a.Origin.(type) {
		case groundtruth.Project:
			ra.Origin = "project"
		case groundtruth.Dependency:
			ra.Origin, ra.Component = "dependency", o.Component
		}
		if a.Depth != nil {
			d := int(*a.Depth)
			ra.Depth = &d
		}
		if a.Distance != nil {
			d := *a.Distance
			ra.Distance = &d
		}
		raw.Assets = append(raw.Assets, ra)
	}
	out, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("prediction: %w", err)
	}
	return append(out, '\n'), nil
}
