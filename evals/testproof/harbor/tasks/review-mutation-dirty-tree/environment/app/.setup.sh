#!/bin/sh
# Builds the git history: Clamp without the negative fix is committed,
# the fix and its test stay uncommitted, like a branch in progress.
set -e
cd /app
git init -q
git config user.email eval@example.com
git config user.name eval
mv quota.go /tmp/quota.fixed
mv quota_test.go /tmp/quota_test.new
cat > quota.go <<'EOF'
// Package quota keeps usage numbers inside a quota.
package quota

// Clamp limits n to the range 0..max.
func Clamp(n, max int) int {
	if n > max {
		return max
	}
	return n
}
EOF
git add go.mod quota.go CASE.md
git commit -qm "quota: add Clamp"
mv /tmp/quota.fixed quota.go
mv /tmp/quota_test.new quota_test.go
