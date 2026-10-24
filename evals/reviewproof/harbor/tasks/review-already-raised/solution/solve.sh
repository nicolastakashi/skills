#!/bin/bash
cat > /app/review.json <<'EOF'
{"findings": [{"id": "", "file": "config.go", "line": 34, "verdict": "fix", "evidence": "config.go:21 says TimeoutSeconds is in seconds, config.go:34 multiplies by time.Millisecond. FromRaw(Raw{Port: \"80\", TimeoutSeconds: 5}).Timeout is 5ms.", "comment": "issue: Timeout uses milliseconds, the field is seconds\n\nconfig.go:34 multiplies TimeoutSeconds by time.Millisecond, so a 5 second timeout becomes 5ms and most requests time out."}]}
EOF
