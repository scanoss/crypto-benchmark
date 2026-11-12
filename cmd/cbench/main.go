// Command cbench validates and fetches the crypto-finder accuracy benchmark corpus.
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/scanoss/crypto-benchmark/internal/adapter/cbom"
	"github.com/scanoss/crypto-benchmark/internal/adapter/cryptofinder"
	"github.com/scanoss/crypto-benchmark/internal/check"
	"github.com/scanoss/crypto-benchmark/internal/corpus"
	"github.com/scanoss/crypto-benchmark/internal/fetch"
	"github.com/scanoss/crypto-benchmark/internal/prediction"
	"github.com/scanoss/crypto-benchmark/internal/score"
	"github.com/scanoss/crypto-benchmark/internal/validate"
)

const usage = `usage:
  cbench validate [--root DIR] [--overlay DIR]
  cbench fetch [--root DIR] [--overlay DIR] [--only ID,...] [--dest DIR]
  cbench adapt crypto-finder --entry ID --input INTERIM.json [--callgraph CALLGRAPH.json] --out FILE [--root DIR] [--overlay DIR]
  cbench adapt cbom --tool NAME --tool-version V --entry ID --input BOM.json --out FILE [--source-root DIR] [--cdxgen] [--root DIR] [--overlay DIR]
  cbench score --tool NAME [--root DIR] [--overlay DIR] [--predictions DIR] [--format md|json] [--out FILE]
`

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprint(stderr, usage)
		return 2
	}
	fs := flag.NewFlagSet("cbench "+args[0], flag.ContinueOnError)
	fs.SetOutput(stderr)
	root := fs.String("root", ".", "benchmark repository root")
	overlay := fs.String("overlay", "", "private overlay directory with the same layout")
	switch args[0] {
	case "validate":
		if fs.Parse(args[1:]) != nil {
			return 2
		}
		return printReport(validate.Run(*root, *overlay), stdout)
	case "fetch":
		only := fs.String("only", "", "comma-separated entry ids to fetch")
		dest := fs.String("dest", "work/corpus", "directory to materialize repositories into")
		if fs.Parse(args[1:]) != nil {
			return 2
		}
		return runFetch(*root, *overlay, *only, *dest, stdout, stderr)
	case "adapt":
		if len(args) < 2 {
			fmt.Fprint(stderr, usage)
			return 2
		}
		return runAdaptCommand(fs, args[1], args[2:], root, overlay, stderr)
	case "score":
		tool := fs.String("tool", "", "tool whose predictions to score")
		preds := fs.String("predictions", "predictions", "directory holding <tool>/<entry-id>.json")
		format := fs.String("format", "md", "output format: md or json")
		out := fs.String("out", "", "write the report to FILE instead of stdout")
		if fs.Parse(args[1:]) != nil {
			return 2
		}
		if *tool == "" || (*format != "md" && *format != "json") {
			fmt.Fprintln(stderr, "score: --tool is required and --format must be md or json")
			return 2
		}
		return runScore(*root, *overlay, *preds, *tool, *format, *out, stdout, stderr)
	default:
		fmt.Fprint(stderr, usage)
		return 2
	}
}

func printReport(r *check.Report, w io.Writer) int {
	errs := 0
	for _, p := range r.Problems {
		fmt.Fprintln(w, p)
		if p.Severity == check.Error {
			errs++
		}
	}
	fmt.Fprintf(w, "%d error(s), %d warning(s)\n", errs, len(r.Problems)-errs)
	if errs > 0 {
		return 1
	}
	return 0
}

func runFetch(root, overlay, only, dest string, stdout, stderr io.Writer) int {
	r := &check.Report{}
	m := corpus.Load(root, overlay, r)
	if r.HasErrors() {
		printReport(r, stderr)
		return 1
	}
	want := map[string]bool{}
	for _, id := range strings.Split(only, ",") {
		if id = strings.TrimSpace(id); id == "" {
			continue
		}
		if _, ok := m.Lookup(id); !ok {
			fmt.Fprintf(stderr, "--only: %q is not a manifest entry\n", id)
			return 2
		}
		want[id] = true
	}
	code := 0
	for _, e := range m.Entries {
		if len(want) > 0 && !want[e.ID] {
			continue
		}
		src, ok := e.Source.(corpus.Remote)
		if !ok {
			continue
		}
		out, err := fetch.Fetch(context.Background(), e.ID, src, dest)
		switch {
		case err != nil:
			fmt.Fprintln(stderr, err)
			code = 1
		case out == fetch.AlreadyPresent:
			fmt.Fprintf(stdout, "present %s @ %s\n", e.ID, src.SHA)
		default:
			fmt.Fprintf(stdout, "fetched %s @ %s\n", e.ID, src.SHA)
		}
	}
	return code
}

func runScore(root, overlay, preds, tool, format, out string, stdout, stderr io.Writer) int {
	in, r := score.Load(root, overlay, preds, tool)
	if r.HasErrors() {
		printReport(r, stderr)
		return 1
	}
	rep := score.Score(tool, in.Manifest, in.Truth, in.Preds)
	var data []byte
	if format == "json" {
		var err error
		if data, err = score.JSON(rep); err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
	} else {
		data = []byte(score.Markdown(rep))
	}
	if out == "" {
		_, err := stdout.Write(data)
		return exitOn(err, stderr)
	}
	return exitOn(os.WriteFile(out, data, 0o644), stderr)
}

func exitOn(err error, stderr io.Writer) int {
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	return 0
}

// runAdaptCommand parses the flags of `adapt TOOL` and writes the
// prediction file that TOOL's adapter produces.
func runAdaptCommand(fs *flag.FlagSet, tool string, args []string, root, overlay *string, stderr io.Writer) int {
	entry := fs.String("entry", "", "manifest entry id the scan ran on")
	out := fs.String("out", "", "prediction file to write, normally predictions/<tool>/<entry>.json")
	var adapt func(input []byte) (prediction.File, []string, error)
	var input *string
	switch tool {
	case cryptofinder.ToolName:
		input = fs.String("input", "", "crypto-finder interim report (scan --output)")
		callgraph := fs.String("callgraph", "", "crypto-finder call graph export (scan --export-callgraph); without it no reachability is reported")
		adapt = func(interim []byte) (prediction.File, []string, error) {
			var graph []byte
			if *callgraph != "" {
				var err error
				if graph, err = os.ReadFile(*callgraph); err != nil {
					return prediction.File{}, nil, err
				}
			}
			res, err := cryptofinder.Adapt(*entry, interim, graph)
			return res.File, res.Skipped, err
		}
	case "cbom":
		input = fs.String("input", "", "CycloneDX 1.6 or 1.7 CBOM (JSON)")
		name := fs.String("tool", "", "name of the tool that wrote the CBOM; predictions/<tool>/ by convention")
		version := fs.String("tool-version", "", "version of that tool")
		sourceRoot := fs.String("source-root", "", "absolute directory the tool scanned; absolute locations under it become relative")
		cdxgen := fs.Bool("cdxgen", false, "read cdxgen's cdx:*crypto:sourceLocation properties and local Go module directories")
		adapt = func(data []byte) (prediction.File, []string, error) {
			if *name == "" || *version == "" {
				return prediction.File{}, nil, fmt.Errorf("adapt cbom: --tool and --tool-version are required")
			}
			res, err := cbom.Adapt(data, cbom.Options{Tool: *name, ToolVersion: *version, Entry: *entry, SourceRoot: *sourceRoot, Cdxgen: *cdxgen})
			return res.File, res.Notes, err
		}
	default:
		fmt.Fprintf(stderr, "adapt: supported tools are %s and cbom\n", cryptofinder.ToolName)
		return 2
	}
	if fs.Parse(args) != nil {
		return 2
	}
	if *entry == "" || *input == "" || *out == "" {
		fmt.Fprintln(stderr, "adapt: --entry, --input and --out are required")
		return 2
	}
	r := &check.Report{}
	m := corpus.Load(*root, *overlay, r)
	if r.HasErrors() {
		printReport(r, stderr)
		return 1
	}
	if _, ok := m.Lookup(*entry); !ok {
		fmt.Fprintf(stderr, "adapt: --entry %q is not a manifest entry\n", *entry)
		return 2
	}
	data, err := os.ReadFile(*input)
	if err != nil {
		return exitOn(err, stderr)
	}
	f, notes, err := adapt(data)
	if err != nil {
		return exitOn(err, stderr)
	}
	for _, n := range notes {
		fmt.Fprintf(stderr, "skipped %s\n", n)
	}
	data, err = prediction.Marshal(f)
	if err != nil {
		return exitOn(err, stderr)
	}
	if err := os.MkdirAll(filepath.Dir(*out), 0o755); err != nil {
		return exitOn(err, stderr)
	}
	return exitOn(os.WriteFile(*out, data, 0o644), stderr)
}
