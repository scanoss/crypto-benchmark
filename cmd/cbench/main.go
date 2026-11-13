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
	"github.com/scanoss/crypto-benchmark/internal/frame"
	"github.com/scanoss/crypto-benchmark/internal/prediction"
	"github.com/scanoss/crypto-benchmark/internal/score"
	"github.com/scanoss/crypto-benchmark/internal/validate"
)

const usage = `usage:
  cbench validate [--root DIR] [--overlay DIR]
  cbench fetch [--root DIR] [--overlay DIR] [--only ID,...] [--dest DIR]
  cbench frame (--entry ID [--dest DIR] | --dir DIR --language L --commit SHA) [--sample | --grep] [--root DIR] [--overlay DIR]
  cbench adapt crypto-finder --entry ID --input INTERIM.json [--callgraph CALLGRAPH.json] --out FILE [--root DIR] [--overlay DIR]
  cbench adapt cbom --tool NAME --tool-version V --entry ID --input BOM.json --out FILE [--source-root DIR] [--cdxgen] [--root DIR] [--overlay DIR]
  cbench score --tool NAME [--root DIR] [--overlay DIR] [--predictions DIR] [--format md|json] [--out FILE]
  cbench unique --run LABEL=DIR/TOOL [--run ...] [--root DIR] [--overlay DIR] [--format md|json] [--out FILE]
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
	case "frame":
		entry := fs.String("entry", "", "suite C entry id; its checkout is read from --dest")
		dest := fs.String("dest", "work/corpus", "directory cbench fetch materialized repositories into")
		dir := fs.String("dir", "", "repository checkout to read instead of an entry's")
		lang := fs.String("language", "", "language of --dir")
		commit := fs.String("commit", "", "commit --dir is checked out at, the seed of --sample")
		sample := fs.Bool("sample", false, "print the seeded sample of the files scope")
		grep := fs.Bool("grep", false, "print the grep candidate lines of the pool scope")
		if fs.Parse(args[1:]) != nil {
			return 2
		}
		return runFrame(*root, *overlay, frameArgs{*entry, *dest, *dir, *lang, *commit, *sample, *grep}, stdout, stderr)
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
	case "unique":
		var runs runFlags
		fs.Var(&runs, "run", "LABEL=DIR/TOOL: one run's prediction directory, as score reads DIR with --tool TOOL; repeat per run")
		format := fs.String("format", "md", "output format: md or json")
		out := fs.String("out", "", "write the report to FILE instead of stdout")
		if fs.Parse(args[1:]) != nil {
			return 2
		}
		if len(runs) < 2 || (*format != "md" && *format != "json") {
			fmt.Fprintln(stderr, "unique: at least two --run LABEL=DIR/TOOL are required and --format must be md or json")
			return 2
		}
		return runUnique(*root, *overlay, runs, *format, *out, stdout, stderr)
	default:
		fmt.Fprint(stderr, usage)
		return 2
	}
}

// runFlags collects repeated --run LABEL=DIR/TOOL values.
type runFlags [][2]string

func (r *runFlags) String() string { return fmt.Sprint(*r) }
func (r *runFlags) Set(v string) error {
	label, dir, ok := strings.Cut(v, "=")
	if !ok || label == "" || dir == "" {
		return fmt.Errorf("want LABEL=DIR/TOOL, got %q", v)
	}
	*r = append(*r, [2]string{label, filepath.Clean(dir)})
	return nil
}

func runUnique(root, overlay string, runs runFlags, format, out string, stdout, stderr io.Writer) int {
	var in score.Inputs
	var compared []score.Run
	for _, r := range runs {
		loaded, rep := score.Load(root, overlay, filepath.Dir(r[1]), filepath.Base(r[1]))
		if rep.HasErrors() {
			printReport(rep, stderr)
			return 1
		}
		in = loaded
		compared = append(compared, score.Run{Label: r[0], Preds: loaded.Preds})
	}
	rep := score.Unique(in.Manifest, in.Truth, compared)
	data := []byte(score.UniqueMarkdown(rep))
	if format == "json" {
		var err error
		if data, err = score.UniqueJSON(rep); err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
	}
	if out == "" {
		_, err := stdout.Write(data)
		return exitOn(err, stderr)
	}
	return exitOn(os.WriteFile(out, data, 0o644), stderr)
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

type frameArgs struct {
	entry, dest, dir, lang, commit string
	sample, grep                   bool
}

// runFrame prints the tool-independent crypto signal of a suite C checkout
// (METHODOLOGY.md, "Suite C"): the signal files, their seeded sample, or the
// pool's grep candidates.
func runFrame(root, overlay string, a frameArgs, stdout, stderr io.Writer) int {
	if (a.entry == "") == (a.dir == "") || (a.sample && a.grep) {
		fmt.Fprintln(stderr, "frame: give either --entry or --dir, and at most one of --sample and --grep")
		return 2
	}
	if a.entry != "" {
		r := &check.Report{}
		m := corpus.Load(root, overlay, r)
		if r.HasErrors() {
			printReport(r, stderr)
			return 1
		}
		e, ok := m.Lookup(a.entry)
		src, remote := e.Source.(corpus.Remote)
		if !ok || !remote {
			fmt.Fprintf(stderr, "frame: --entry %q is not a url entry of the manifest\n", a.entry)
			return 2
		}
		a.dir, a.lang, a.commit = filepath.Join(a.dest, e.ID), string(e.Language), src.SHA
	}
	lang, err := check.OneOf(a.lang, corpus.Java, corpus.Python, corpus.Go, corpus.Rust, corpus.JavaScript)
	if err != nil {
		fmt.Fprintf(stderr, "frame: --language: %v\n", err)
		return 2
	}
	if a.sample && a.commit == "" {
		fmt.Fprintln(stderr, "frame: --sample needs the commit (--entry, or --commit with --dir)")
		return 2
	}
	if a.grep {
		lines, err := frame.Grep(a.dir, lang)
		if err != nil {
			return exitOn(err, stderr)
		}
		for _, l := range lines {
			fmt.Fprintf(stdout, "%s:%d\n", l.File, l.Line)
		}
		return 0
	}
	f, err := frame.Scan(a.dir, lang)
	if err != nil {
		return exitOn(err, stderr)
	}
	files := f.Signal
	if a.sample {
		files = frame.Sample(f.Signal, a.commit, frame.SampleSize)
	} else {
		fmt.Fprintf(stdout, "# %d %s source files, %d signal files\n", f.SourceFiles, lang, len(f.Signal))
	}
	for _, p := range files {
		fmt.Fprintln(stdout, p)
	}
	return 0
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
