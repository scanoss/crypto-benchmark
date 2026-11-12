#!/usr/bin/env bash
# Runs the pinned sonar-cryptography plugin on one corpus entry and writes its
# CBOM to OUT. The exact commands, with every version, go to OUT.command, and
# the scanner's log to OUT.log.
#
#   scripts/run-sonar-cryptography.sh ENTRY_DIR OUT
#
# The scan runs headless in the CBOMkit-action container, which drives the
# sonar-cryptography detection rules through the SonarQube Java, Python and Go
# frontends without a SonarQube server. Convert OUT with:
#
#   cbench adapt cbom --tool sonar-cryptography --tool-version 1.7.0 \
#     --entry ID --input OUT --out FILE
#
# The language comes from the entry's root: go.mod is Go, pom.xml or
# build.gradle(.kts) is Java, otherwise Python. Set SONAR_CRYPTO_LANGUAGE to
# java, python or go to choose. A Java entry must be built first: the scanner
# reads compiled classes and jars from inside ENTRY_DIR, and the script refuses
# a Java entry without .class files.
#
# DOCKER names the docker binary (default docker). SONAR_CRYPTO_IMAGE replaces
# the pinned image, for a mirror of it; its versions are checked all the same.
set -euo pipefail

readonly PLUGIN_VERSION=1.7.0
readonly CBOMKIT_LIB_VERSION=1.3.0
readonly ACTION_VERSION=v2.3.0
readonly PINNED_IMAGE=ghcr.io/cbomkit/cbomkit-action@sha256:47c2505267dc5f055152616ffa0d0076b7d0824fe9e0de670b572d1ebad2fc4b
readonly JAR=/cbomkit-action/CBOMkit-action.jar
readonly WS=/scan/ws

if [ $# -ne 2 ]; then
	echo "usage: $0 ENTRY_DIR OUT" >&2
	exit 2
fi
entry=$(realpath "$1")
out=$(realpath -m "$2")
record="$out.command"
docker=${DOCKER:-docker}
image=${SONAR_CRYPTO_IMAGE:-$PINNED_IMAGE}
mkdir -p "$(dirname "$out")"

die() {
	echo "$0: $*" >&2
	exit 1
}

# Prints "groupId/artifactId version" for each component embedded in the jar.
embedded_versions() {
	"$docker" run --rm --network none --entrypoint sh "$image" -c \
		"cd /tmp && for c in $*; do jar xf $JAR META-INF/maven/\$c/pom.properties && echo \$c \$(sed -n 's/^version=//p' META-INF/maven/\$c/pom.properties); done"
}

want="com.ibm/sonar-cryptography-plugin $PLUGIN_VERSION
org.pqca/cbomkit-lib $CBOMKIT_LIB_VERSION
org.pqca/CBOMkit-action $ACTION_VERSION"
got=$(embedded_versions com.ibm/sonar-cryptography-plugin org.pqca/cbomkit-lib org.pqca/CBOMkit-action || true)
[ "$got" = "$want" ] || die "want in $image:
$want
got:
$got"

language=${SONAR_CRYPTO_LANGUAGE:-}
if [ -z "$language" ]; then
	if [ -f "$entry/go.mod" ]; then
		language=go
	elif [ -f "$entry/pom.xml" ] || [ -f "$entry/build.gradle" ] || [ -f "$entry/build.gradle.kts" ]; then
		language=java
	else
		language=python
	fi
fi

# cbomkit-lib indexes a directory that is not a module by scanning only the
# first leaf directory it finds, so the root of the scanned copy must be a
# module: a Python entry without a build file gets an empty setup.cfg.
case "$language" in
python)
	marker=
	[ -e "$entry/pyproject.toml" ] || [ -e "$entry/setup.cfg" ] || [ -e "$entry/setup.py" ] || marker=setup.cfg
	;;
go)
	[ -f "$entry/go.mod" ] || die "a Go entry needs go.mod at its root"
	marker=
	;;
java)
	[ -d "$entry/src" ] || die "a Java entry needs src/ next to its build file"
	[ -n "$(find "$entry" -name '*.class' -print -quit)" ] ||
		die "no .class files under $entry: build the Java entry first"
	marker=
	;;
*) die "unknown language '$language'" ;;
esac
work=$(mktemp -d)
trap 'rm -rf "$work"' EXIT
mkdir "$work/out"
cp -a "$entry" "$work/ws"

{
	echo "# image $image"
	echo "# sonar-cryptography-plugin $PLUGIN_VERSION, cbomkit-lib $CBOMKIT_LIB_VERSION, CBOMkit-action $ACTION_VERSION"
	echo "# language $language"
	echo "cp -a ENTRY_DIR WS"
	[ -z "$marker" ] || echo "touch WS/$marker"
} >"$record"
[ -z "$marker" ] || touch "$work/ws/$marker"

# Occurrence locations are written relative to the JVM's working directory, so
# the scan runs inside the workspace. CBOMKIT_EXCLUDE matches nothing: the
# default excludes test files, and the benchmark scores every file.
cmd=("$docker" run --rm --network none --user "$(id -u):$(id -g)" -e HOME=/tmp -w "$WS"
	-v "$work/ws:$WS:ro" -v "$work/out:/scan/out"
	-e GITHUB_WORKSPACE="$WS" -e CBOMKIT_OUTPUT_DIR=/scan/out
	-e CBOMKIT_LANGUAGES="$language" -e CBOMKIT_GENERATE_MODULE_CBOMS=false
	-e 'CBOMKIT_EXCLUDE=(?!)' "$image")
printf '%q ' "${cmd[@]}" | sed "s|$work/ws|WS|g; s|$work/out|OUTDIR|g; s|--user [0-9]*:[0-9]*|--user \"\$(id -u):\$(id -g)\"|; s| $||" >>"$record"
echo >>"$record"
echo "cp OUTDIR/cbom.json OUT" >>"$record"

"${cmd[@]}" >"$out.log" 2>&1 || die "scan failed, see $out.log"
[ -s "$work/out/cbom.json" ] || die "no cbom.json written, see $out.log"
cp "$work/out/cbom.json" "$out"
