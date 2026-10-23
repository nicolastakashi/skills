#!/bin/bash
# Counts the trials in a Harbor job where the agent opened the testproof skill.
# Usage: ./activation.sh jobs/<job>
set -euo pipefail
total=0; hit=0
for trial in "$1"/*/; do
  [ -d "${trial}agent" ] || continue
  total=$((total+1))
  if grep -rEq '"name": ?"Skill".{0,200}"skill": ?"testproof"' "${trial}agent"; then
    hit=$((hit+1)); echo "hit   $(basename "$trial")"
  else
    echo "miss  $(basename "$trial")"
  fi
done
echo "activation: $hit of $total"
