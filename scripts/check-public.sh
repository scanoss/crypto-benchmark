#!/usr/bin/env bash
# Fails when the tree holds anything that must not be in a public repository:
# backlog IDs, private repository names, local home paths, key-like strings,
# the internal ruleset name, or detection-rule IDs.
#
#   scripts/check-public.sh
#
# It searches tracked files and untracked files that are not ignored, so it
# also catches a file before it is added. Exit 0 when clean, 1 on any hit.
set -euo pipefail

cd "$(dirname "$0")/.."

# The test PKI of the rustls fixtures is a throwaway key pair made for the
# benchmark (fixtures/rust/README.md); it is the only private key allowed.
allowed_keys=(
	':!fixtures/rust/rustls/certs/localhost.key'
	':!fixtures/rust/rustls-webpki/certs/localhost.key'
	':!fixtures/rust/tokio-rustls/certs/localhost.key'
)

# Each check: a name, then an extended regular expression. Matching is
# case-insensitive only where the line says so.
checks=(
	'backlog ID|\b(CF|CR|BM)-[0-9]+\b'
	'private repository name|crypto[_-]rules|crypto-benchmark-backlog|scanoss/roadmap'
	'backlog reference (case-insensitive)|i:\bbacklog\b'
	'local home path|/home/[A-Za-z0-9_.-]+/|/Users/[A-Za-z0-9_.-]+/|~/(src|\.scanoss)\b|[A-Za-z]:\\Users\\'
	'key-like string|\bsk-[A-Za-z0-9_-]{20,}|\bghp_[A-Za-z0-9]{30,}|github_pat_[A-Za-z0-9_]{20,}|\bAKIA[0-9A-Z]{16}\b|\bxox[abpr]-[A-Za-z0-9-]{10,}|SCANOSS_API_KEY *[:=] *["'"'"']?[A-Za-z0-9]'
	'private key block|^-----BEGIN ([A-Z]+ )?PRIVATE KEY-----$'
	'internal ruleset name|i:\bdca\b|"ruleset" *:|ruleset=[a-z]'
	'detection rule ID|[a-z0-9_-]+\.(algorithm|protocol|certificate|related-crypto-material)\.[a-z0-9_-]+\.(c|cpp|go|java|javascript|typescript|python|rust)\.'
)

status=0
for check in "${checks[@]}"; do
	name=${check%%|*}
	pattern=${check#*|}
	flags=(-n -I -E)
	if [[ $pattern == i:* ]]; then
		flags+=(-i)
		pattern=${pattern#i:}
	fi
	exclude=(':!scripts/check-public.sh')
	if [[ $name == 'private key block' ]]; then
		exclude+=("${allowed_keys[@]}")
	fi
	if hits=$(git grep --untracked "${flags[@]}" -e "$pattern" -- . "${exclude[@]}"); then
		echo "check-public: $name:" >&2
		echo "$hits" | cut -c1-200 >&2
		status=1
	fi
done

if [ "$status" -eq 0 ]; then
	echo "check-public: clean"
fi
exit "$status"
