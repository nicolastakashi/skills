#!/bin/bash
cat > /app/review.json <<'EOF'
{"findings": [{"id": "", "file": "batch.go", "line": 10, "verdict": "fix", "evidence": "batch.go:10 stops while i+size <= len(ids). Chunk([a b c], 2) returns [[a b]] and loses c.", "comment": "issue: Chunk drops the last short batch\n\nThe loop at batch.go:10 stops before a partial batch, so Chunk([a b c], 2) returns [[a b]] and c is never fetched, while the doc says the last batch holds the remainder."}, {"id": "", "file": "batch.go", "line": 26, "verdict": "drop", "evidence": "go.mod sets go 1.24, so i and id are per-iteration. TestFetchAllKeepsOrder passes with go test -race.", "comment": ""}]}
EOF
