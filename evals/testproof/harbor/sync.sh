#!/bin/bash
# Copies the shared Dockerfile, test.sh and verifier into every task.
# Harbor only sees a task's own folder, so shared files must be copied in.
# Usage: ./sync.sh          copy
#        ./sync.sh --check  fail if any task is out of date
set -euo pipefail
cd "$(dirname "$0")"
status=0
for task in tasks/*/; do
  pairs=(
    "_shared/Dockerfile:${task}environment/Dockerfile"
    "_shared/test.sh:${task}tests/test.sh"
    "_shared/verify/main.go:${task}tests/verify/main.go"
    "_shared/verify/go.mod:${task}tests/verify/go.mod"
  )
  for p in "${pairs[@]}"; do
    src=${p%%:*}; dst=${p#*:}
    if [ "${1:-}" = "--check" ]; then
      cmp -s "$src" "$dst" || { echo "out of date: $dst"; status=1; }
    else
      mkdir -p "$(dirname "$dst")"; cp "$src" "$dst"
    fi
  done
  [ "${1:-}" = "--check" ] || chmod +x "${task}tests/test.sh"
done
exit $status
