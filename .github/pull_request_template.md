## What this adds or fixes

<!-- Suite, language, library and version, or the ground-truth error it fixes (link the issue). -->

## Checklist

- [ ] The programs build and run, with every dependency pinned to an exact version.
- [ ] Every asset of every changed file is in the ground truth, following the [annotation guide](https://github.com/scanoss/crypto-benchmark/blob/main/docs/annotation-guide.md).
- [ ] `go run ./cmd/cbench validate` passes.
- [ ] `go test ./...` and `scripts/check-public.sh` pass.
- [ ] No tool output, prediction file, real key or secret is included.

A contributed ground truth is the first of two annotations: a maintainer runs a second, independent one and settles every disagreement against the guide before merging ([CONTRIBUTING.md](https://github.com/scanoss/crypto-benchmark/blob/main/CONTRIBUTING.md#what-happens-next)).
