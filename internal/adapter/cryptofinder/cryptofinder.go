// Package cryptofinder converts a crypto-finder interim report, and
// optionally its call graph export, into a benchmark prediction file.
// The mapping is documented in README.md next to this file.
package cryptofinder

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/scanoss/crypto-benchmark/internal/algo"
	"github.com/scanoss/crypto-benchmark/internal/check"
	"github.com/scanoss/crypto-benchmark/internal/groundtruth"
	"github.com/scanoss/crypto-benchmark/internal/prediction"
)

// ToolName is the tool name every crypto-finder prediction file carries.
const ToolName = "crypto-finder"

// Interim report versions whose fields this adapter reads: dependency_info.purl
// and occurrence_key arrived in 1.5. 1.8 adds the per-asset pqc block, which
// is decoded and left out of the prediction: it classifies an algorithm for
// post-quantum risk and no scored field reads it. 1.9 adds rules.lease, the
// signed lease of a remote ruleset; it is decoded and left out too, since the
// ruleset version and checksum already identify the rules.
var interimVersions = []string{"1.5", "1.6", "1.7", "1.8", "1.9"}

// Call graph export versions carrying finding_graphs[].reachability (6.8+)
// and finding_graphs[].dependency (6.17+).
var callgraphVersions = []string{"6.17", "6.18"}

// interimReport mirrors schemas/interim-report-schema.json in the
// crypto-finder repository. It is decoded strictly: an undeclared property
// means the contract moved and the adapter must be reviewed.
type interimReport struct {
	Version  string    `json:"version"`
	Tool     toolInfo  `json:"tool"`
	Rules    rulesInfo `json:"rules"`
	Findings []finding `json:"findings"`
}

type toolInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type rulesInfo struct {
	Source         string      `json:"source"`
	Name           string      `json:"name"`
	Version        string      `json:"version"`
	ChecksumSHA256 string      `json:"checksum_sha256"`
	Lease          *rulesLease `json:"lease"`
}

type rulesLease struct {
	ExpiresAt    string `json:"expires_at"`
	SigningKeyID string `json:"signing_key_id"`
}

type finding struct {
	FilePath string  `json:"file_path"`
	Language string  `json:"language"`
	Assets   []asset `json:"cryptographic_assets"`
}

type ruleInfo struct {
	ID       string `json:"id"`
	Message  string `json:"message"`
	Severity string `json:"severity"`
	Version  string `json:"version"`
}

type asset struct {
	StartLine           int                `json:"start_line"`
	EndLine             int                `json:"end_line"`
	StartCol            int                `json:"start_col"`
	EndCol              int                `json:"end_col"`
	TerminalStartCol    int                `json:"terminal_start_col"`
	TerminalEndCol      int                `json:"terminal_end_col"`
	ConditionedValue    string             `json:"conditioned_value"`
	Match               string             `json:"match"`
	Rules               []ruleInfo         `json:"rules"`
	Rule                *ruleInfo          `json:"rule"`
	Status              string             `json:"status"`
	Metadata            map[string]string  `json:"metadata"`
	ParameterConditions json.RawMessage    `json:"parameter_conditions"`
	OID                 string             `json:"oid"`
	FindingID           string             `json:"finding_id"`
	OccurrenceKey       string             `json:"occurrence_key"`
	Source              string             `json:"source"`
	DependencyInfo      *dependencyInfo    `json:"dependency_info"`
	PURL                string             `json:"purl"`
	PQC                 *pqcClassification `json:"pqc"`
}

// pqcClassification mirrors the interim schema's PQCClassification (v1.8+).
// Nothing reads it; it is declared so strict decoding accepts the block.
type pqcClassification struct {
	Status         string        `json:"status"`
	Family         string        `json:"family"`
	Deadline       int           `json:"deadline"`
	Deadlines      []pqcDeadline `json:"deadlines"`
	DatasetVersion string        `json:"dataset_version"`
}

type pqcDeadline struct {
	Source          string `json:"source"`
	DeprecatedAfter int    `json:"deprecated_after"`
	DisallowedAfter int    `json:"disallowed_after"`
	MigrateBy       int    `json:"migrate_by"`
	Citation        string `json:"citation"`
}

type dependencyInfo struct {
	Module  string `json:"module"`
	Version string `json:"version"`
	PURL    string `json:"purl"`
}

// callgraph holds only what the adapter reads from the export; the rest of
// that large document is not decoded.
type callgraph struct {
	SchemaVersion string         `json:"schema_version"`
	FindingGraphs []findingGraph `json:"finding_graphs"`
}

type findingGraph struct {
	FindingID     string             `json:"finding_id"`
	OccurrenceKey string             `json:"occurrence_key"`
	Reachability  string             `json:"reachability"`
	Dependency    *findingDependency `json:"dependency"`
}

type findingDependency struct {
	Relationship string            `json:"relationship"`
	Path         []json.RawMessage `json:"path"`
}

// graphKey joins an interim asset to its finding graph: finding_id alone is
// not unique, (finding_id, occurrence_key) is.
type graphKey struct{ findingID, occurrenceKey string }

// Result is the prediction file plus one note per interim asset left out.
type Result struct {
	File    prediction.File
	Skipped []string
}

// Adapt converts an interim report into the prediction file for entry.
// callgraphJSON is nil when the scan exported no call graph; the result
// then reports no reachability.
func Adapt(entry string, interimJSON, callgraphJSON []byte) (Result, error) {
	var rep interimReport
	if err := check.DecodeStrictJSON(interimJSON, &rep); err != nil {
		return Result{}, fmt.Errorf("cryptofinder: interim report: %w", err)
	}
	if !slices.Contains(interimVersions, rep.Version) {
		return Result{}, fmt.Errorf("cryptofinder: interim report version %q is not one of %s", rep.Version, strings.Join(interimVersions, "|"))
	}
	if rep.Tool.Name != ToolName {
		return Result{}, fmt.Errorf("cryptofinder: interim report tool.name is %q, want %q", rep.Tool.Name, ToolName)
	}
	graphs, err := parseCallgraph(callgraphJSON)
	if err != nil {
		return Result{}, err
	}
	res := Result{File: prediction.File{
		Tool:                 prediction.Tool{Name: ToolName, Version: rep.Tool.Version, Config: config(rep.Rules)},
		Entry:                entry,
		ReachabilityReported: graphs != nil,
		Assets:               []prediction.Asset{},
	}}
	for _, f := range rep.Findings {
		for _, a := range f.Assets {
			p, skip := convert(f.FilePath, a, graphs)
			if skip != "" {
				res.Skipped = append(res.Skipped, fmt.Sprintf("%s:%d: %s", f.FilePath, a.StartLine, skip))
				continue
			}
			res.File.Assets = append(res.File.Assets, p)
		}
	}
	return res, nil
}

// config records which ruleset produced the scan. The ruleset name is left
// out: the version and checksum identify a remote ruleset on their own.
func config(r rulesInfo) map[string]string {
	c := map[string]string{}
	for k, v := range map[string]string{"ruleset_source": r.Source, "ruleset_version": r.Version, "ruleset_checksum": r.ChecksumSHA256} {
		if v != "" {
			c[k] = v
		}
	}
	return c
}

func parseCallgraph(data []byte) (map[graphKey]findingGraph, error) {
	if data == nil {
		return nil, nil
	}
	var cg callgraph
	if err := json.Unmarshal(data, &cg); err != nil {
		return nil, fmt.Errorf("cryptofinder: call graph: %w", err)
	}
	if !slices.Contains(callgraphVersions, cg.SchemaVersion) {
		return nil, fmt.Errorf("cryptofinder: call graph schema_version %q is not one of %s", cg.SchemaVersion, strings.Join(callgraphVersions, "|"))
	}
	graphs := make(map[graphKey]findingGraph, len(cg.FindingGraphs))
	for _, g := range cg.FindingGraphs {
		k := graphKey{g.FindingID, g.OccurrenceKey}
		if _, dup := graphs[k]; dup {
			return nil, fmt.Errorf("cryptofinder: call graph has two finding graphs for finding_id %q, occurrence_key %q", g.FindingID, g.OccurrenceKey)
		}
		graphs[k] = g
	}
	return graphs, nil
}

func convert(file string, a asset, graphs map[graphKey]findingGraph) (prediction.Asset, string) {
	md := a.Metadata
	typ := groundtruth.AssetType(strings.ToLower(strings.TrimSpace(md["assetType"])))
	if !slices.Contains(groundtruth.AssetTypes, typ) {
		return prediction.Asset{}, fmt.Sprintf("assetType %q is not a CycloneDX asset type", md["assetType"])
	}
	p := prediction.Asset{
		File:      prediction.NormalizePath(file),
		Lines:     groundtruth.Lines{Start: a.StartLine, End: max(a.StartLine, a.EndLine)},
		Type:      typ,
		Algorithm: algorithmName(md),
		Reachable: prediction.NoAnswer,
	}
	switch typ {
	case groundtruth.Algorithm:
		if prim := groundtruth.Primitive(strings.ToLower(strings.TrimSpace(md["algorithmPrimitive"]))); slices.Contains(groundtruth.Primitives, prim) {
			p.Primitive = prim
		}
		p.Params = groundtruth.Parameters{
			Mode:    strings.TrimSpace(md["algorithmMode"]),
			Padding: strings.TrimSpace(md["algorithmPadding"]),
			KeySize: positiveInt(md["algorithmParameterSetIdentifier"]),
			Hash:    strings.TrimSpace(md["algorithmHashFunction"]),
		}
		if c, ok := algo.Curve(md["algorithmParameterSetIdentifier"]); ok {
			p.Params.Curve = c
		}
	case groundtruth.RelatedCryptoMaterial:
		p.Params = groundtruth.Parameters{KeySize: positiveInt(md["materialSize"]), Curve: strings.TrimSpace(md["materialCurve"])}
		if mt := groundtruth.MaterialType(strings.ToLower(strings.TrimSpace(md["materialType"]))); slices.Contains(groundtruth.MaterialTypes, mt) {
			p.MaterialType = mt
		}
	}
	switch a.Source {
	case "", "direct":
		p.Origin = groundtruth.Project{}
	case "dependency":
		if a.DependencyInfo == nil || a.DependencyInfo.PURL == "" {
			return prediction.Asset{}, "dependency finding without dependency_info.purl"
		}
		p.Origin = groundtruth.Dependency{Component: a.DependencyInfo.PURL}
	default:
		return prediction.Asset{}, fmt.Sprintf("source %q is not direct or dependency", a.Source)
	}
	if g, ok := graphs[graphKey{a.FindingID, a.OccurrenceKey}]; ok {
		p.Reachable, p.Depth, p.Distance = reachability(g, p.Origin)
	}
	return p, ""
}

// reachability returns the answer, the depth and, when the call graph states
// it, the exact dependency distance: 0 for project code, else the length of
// dependency.path. A relationship without a path gives a depth only.
func reachability(g findingGraph, origin groundtruth.Origin) (prediction.Answer, *groundtruth.Depth, *int) {
	var answer prediction.Answer
	switch g.Reachability {
	case "reachable":
		answer = prediction.Reachable
	case "unreachable":
		answer = prediction.Unreachable
	default: // unknown, not_applicable, or absent
		return prediction.NoAnswer, nil, nil
	}
	var dist *int
	var d groundtruth.Depth
	_, project := origin.(groundtruth.Project)
	switch {
	case project:
		d, dist = 0, new(int)
	case g.Dependency == nil:
		return answer, nil, nil
	case len(g.Dependency.Path) > 0:
		n := len(g.Dependency.Path)
		d, dist = groundtruth.Depth(min(n, 2)), &n
	case g.Dependency.Relationship == "direct":
		d = 1
	case g.Dependency.Relationship == "transitive":
		d = 2
	default:
		return answer, nil, nil
	}
	return answer, &d, dist
}

// algorithmName is metadata.algorithmName, the specific algorithm (SHA-256),
// and the family (SHA-2) only when the name is absent. An unresolved name
// such as SHA-$bits is kept: the scorer counts it as no algorithm.
func algorithmName(md map[string]string) string {
	if name := strings.TrimSpace(md["algorithmName"]); name != "" {
		return name
	}
	return strings.TrimSpace(md["algorithmFamily"])
}

func positiveInt(s string) int {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil || n <= 0 {
		return 0
	}
	return n
}
