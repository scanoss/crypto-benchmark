#!/usr/bin/env bash
# Runs crypto-finder on every suite B pair entry of the given languages and
# writes one prediction file per entry.
#
#   scripts/run-crypto-finder-suite-b.sh CF_BIN PREDICTIONS_DIR LANG...
#
# Each entry is first built by scripts/prepare-suite-b.sh into
# work/suite-b/<entry>. The scan is a dependency scan of the built app,
# pinned to the remote ruleset in RULESET_VERSION (default v1.48.0), and
# needs a SCANOSS API key (SCANOSS_API_KEY or crypto-finder's config).
# The exact command of every scan goes to work/suite-b/<entry>/scan.command.
set -euo pipefail

if [ $# -lt 3 ]; then
	echo "usage: $0 CF_BIN PREDICTIONS_DIR LANG..." >&2
	exit 2
fi
cf=$(realpath "$1")
preds=$(realpath -m "$2")
shift 2
root=$(cd "$(dirname "$0")/.." && pwd)
ruleset=${RULESET_VERSION:-v1.48.0}
mkdir -p "$preds/crypto-finder"
cd "$root"

for lang in "$@"; do
	for dir in fixtures/"$lang"/suite-b/*/app-*; do
		scenario=$(basename "$(dirname "$dir")")
		variant=${dir##*app-}
		id="$lang-b-$scenario-$variant"
		work="work/suite-b/$id"
		scripts/prepare-suite-b.sh "$dir" "$work"
		cmd=(env $(cat "$work/scan.env") "$cf" scan --ruleset-version "$ruleset"
			--scan-dependencies --no-dependency-findings-api --findings-cache none
			--export-callgraph "$work/callgraph.json" --output "$work/interim.json" "$work/app")
		printf '%q ' "${cmd[@]}" | sed 's/ $//' >"$work/scan.command"
		echo >>"$work/scan.command"
		"${cmd[@]}" >"$work/scan.log" 2>&1 || {
			echo "scan failed for $id (see $work/scan.log)" >&2
			continue
		}
		go run ./cmd/cbench adapt crypto-finder --entry "$id" --input "$work/interim.json" \
			--callgraph "$work/callgraph.json" --out "$preds/crypto-finder/$id.json"
		echo "scanned $id"
	done
done
