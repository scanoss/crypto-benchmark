#!/usr/bin/env bash
# Builds one suite B fixture with its package manager, the way a user
# installs a published dependency, and runs the application once.
#
#   scripts/prepare-suite-b.sh ENTRY_DIR OUT
#
# ENTRY_DIR is a manifest path such as fixtures/go/suite-b/transitive/app-reachable.
# Its scenario directory holds the libraries (liba, libb, libc) that both
# variants share. OUT receives a copy of the libraries and of this app, and
# the app is built and run there:
#
#   go          go.mod require + replace of each library module; a
#               third-party module (a web framework) comes from the Go
#               module proxy at the version go.sum pins
#   python      each library built into a wheel (pip wheel); every other
#               pinned line of the app's requirements.txt (a web framework)
#               downloaded from PyPI into the same wheel directory; the
#               requirements.txt then installed from that directory alone
#               into OUT/venv, so pip never asks PyPI for a cbench- name
#   java        each library installed with its sources jar into the Maven
#               repository OUT/../m2 (shared by every entry; artifacts are
#               unique per scenario); the app built against it. Needs a
#               full JDK 21 in JAVA_HOME
#   javascript  each library packed with npm pack into OUT/dist; the app's
#               package.json depends on the tarballs and npm install
#               writes package-lock.json and node_modules
#   rust        cargo path dependencies; registry crates fetched at the
#               versions Cargo.lock pins, then built offline
#
# OUT/app is a git repository of its own, so no .gitignore of an enclosing
# repository hides the app from a scanner. OUT/scan.env holds the
# environment a dependency scan of OUT/app needs (VIRTUAL_ENV, MAVEN_OPTS).
set -euo pipefail

if [ $# -ne 2 ]; then
	echo "usage: $0 ENTRY_DIR OUT" >&2
	exit 2
fi
entry=$(realpath "$1")
out=$(realpath -m "$2")
scenario=$(dirname "$entry")
lang=$(basename "$(dirname "$(dirname "$scenario")")")

rm -rf "$out"
mkdir -p "$out"
libs=()
for lib in liba libb libc; do
	if [ -d "$scenario/$lib" ]; then
		cp -R "$scenario/$lib" "$out/$lib"
		libs+=("$lib")
	fi
done
cp -R "$entry" "$out/app"
: >"$out/scan.env"

# deepest returns the libraries deepest first (libc, libb, liba).
deepest() {
	local i
	for ((i = ${#libs[@]} - 1; i >= 0; i--)); do echo "${libs[i]}"; done
}

case $lang in
go)
	# The app's go.mod replaces ../liba; the copy keeps that layout.
	(cd "$out/app" && GOFLAGS=-mod=mod go run . >"$out/run.txt")
	;;
python)
	python3 -m venv "$out/venv"
	for lib in "${libs[@]}"; do
		"$out/venv/bin/pip" wheel --quiet --no-deps --wheel-dir "$out/wheels" "$out/$lib"
	done
	if grep -qv '^cbench-' "$out/app/requirements.txt"; then
		grep -v '^cbench-' "$out/app/requirements.txt" >"$out/third-party.txt"
		"$out/venv/bin/pip" download --quiet --no-deps --dest "$out/wheels" -r "$out/third-party.txt"
	fi
	"$out/venv/bin/pip" install --quiet --no-index --find-links "$out/wheels" -r "$out/app/requirements.txt"
	(cd "$out/app" && "$out/venv/bin/python" main.py >"$out/run.txt")
	echo "VIRTUAL_ENV=$out/venv" >"$out/scan.env"
	;;
java)
	repo="$(dirname "$out")/m2"
	mkdir -p "$repo"
	for lib in $(deepest); do
		(cd "$out/$lib" && mvn -q -B "-Dmaven.repo.local=$repo" install)
	done
	(cd "$out/app" && mvn -q -B "-Dmaven.repo.local=$repo" package \
		dependency:build-classpath "-Dmdep.outputFile=$out/classpath.txt")
	main=$(sed -n 's:.*<mainClass>\(.*\)</mainClass>.*:\1:p' "$out/app/pom.xml")
	java -cp "$out/app/target/classes:$(cat "$out/classpath.txt")" "$main" >"$out/run.txt"
	# The build output is not source; keep the scanned tree to what a
	# checkout holds.
	rm -rf "$out/app/target"
	echo "MAVEN_OPTS=-Dmaven.repo.local=$repo" >"$out/scan.env"
	[ -z "${JAVA_HOME:-}" ] || echo "JAVA_HOME=$JAVA_HOME" >>"$out/scan.env"
	;;
javascript)
	mkdir -p "$out/dist"
	for lib in "${libs[@]}"; do
		(cd "$out/$lib" && npm pack --quiet --pack-destination "$out/dist" >/dev/null)
	done
	(cd "$out/app" && npm install --quiet --no-audit --no-fund >/dev/null && node index.js >"$out/run.txt")
	;;
rust)
	(cd "$out/app" && cargo fetch --locked --quiet &&
		CARGO_NET_OFFLINE=true cargo run --quiet --offline >"$out/run.txt")
	rm -rf "$out/app/target" "$out"/lib*/target
	echo "CARGO_NET_OFFLINE=true" >"$out/scan.env"
	;;
*)
	echo "$0: unknown language $lang (from $entry)" >&2
	exit 1
	;;
esac

git -C "$out/app" init -q
echo "prepared $lang $(basename "$scenario")/$(basename "$entry"): $(head -c 120 "$out/run.txt")"
