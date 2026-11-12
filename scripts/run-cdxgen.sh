#!/usr/bin/env bash
# Runs the pinned cdxgen on one corpus entry and writes its CBOM to OUT.
# The exact commands, with the cdxgen version, go to OUT.command.
#
#   scripts/run-cdxgen.sh ENTRY_DIR OUT
#
# A directory with go.mod gets cdxgen -t go and then evinse -l go (golem),
# one with Cargo.toml cdxgen -t rust and then evinse -l rust (rusi); any
# other directory gets cbom, whose source analysis covers JavaScript and
# TypeScript. Convert OUT with:
#
#   cbench adapt cbom --tool cdxgen --tool-version 13.6.0 --cdxgen \
#     --source-root "$(realpath ENTRY_DIR)" --entry ID --input OUT --out FILE
#
# By default the tools come from `npm exec` of the pinned package. Set
# CDXGEN_BIN to the node_modules/.bin of an install of the same version to
# skip that.
set -euo pipefail

readonly CDXGEN_VERSION=13.6.0

if [ $# -ne 2 ]; then
	echo "usage: $0 ENTRY_DIR OUT" >&2
	exit 2
fi
entry=$(realpath "$1")
out=$(realpath -m "$2")
record="$out.command"
[ -z "${CDXGEN_BIN:-}" ] || CDXGEN_BIN=$(realpath "$CDXGEN_BIN")
mkdir -p "$(dirname "$out")"

tool() {
	if [ -n "${CDXGEN_BIN:-}" ]; then
		"$CDXGEN_BIN/$1" "${@:2}"
	else
		npm exec --yes --package="@cdxgen/cdxgen@$CDXGEN_VERSION" -- "$@"
	fi
}

run() {
	printf '%q' "$1" >>"$record"
	printf ' %q' "${@:2}" >>"$record"
	echo >>"$record"
	tool "$@"
}

version=$(tool cdxgen --version | grep -o 'CycloneDX Generator [0-9.]*' | head -n1)
if [ "$version" != "CycloneDX Generator $CDXGEN_VERSION" ]; then
	echo "$0: want cdxgen $CDXGEN_VERSION, got '$version'" >&2
	exit 1
fi
echo "# $version" >"$record"

# cdxgen and its analyzers write intermediate slices to the working directory.
work=$(mktemp -d)
trap 'rm -rf "$work"' EXIT
cd "$work"
sbom="$work/sbom.json"
if [ -f "$entry/go.mod" ]; then
	run cdxgen -t go -o "$sbom" "$entry"
	run evinse -i "$sbom" -o "$out" -l go "$entry"
elif [ -f "$entry/Cargo.toml" ]; then
	run cdxgen -t rust -o "$sbom" "$entry"
	run evinse -i "$sbom" -o "$out" -l rust "$entry"
else
	run cbom -o "$out" "$entry"
fi
