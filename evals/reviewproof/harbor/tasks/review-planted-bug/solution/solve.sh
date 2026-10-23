#!/bin/bash
cat > /app/review.json <<'EOF'
{"findings": [{"id": "", "file": "limiter.go", "line": 27, "verdict": "fix", "evidence": "limiter.go:27 starts a new window but never clears l.counts. New(1, time.Minute): Allow at t=0 is true, Allow at t=2m is false.", "comment": "issue: counts never reset when a new window starts\n\nlimiter.go:27 moves windowStart but keeps l.counts, so a user who hits Max stays blocked forever. Clear counts when the window rolls over."}]}
EOF
