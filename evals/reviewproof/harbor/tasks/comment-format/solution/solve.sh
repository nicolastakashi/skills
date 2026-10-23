#!/bin/bash
cat > /app/review.json <<'EOF'
{"findings": [{"id": "", "file": "slug.go", "line": 10, "verdict": "fix", "evidence": "slug.go:10 passes n=1 to strings.Replace. Make(\"Release Notes 2024\") returns \"release-notes 2024\".", "comment": "issue: Make replaces only the first space\n\nslug.go:10 calls strings.Replace with n=1, so \"Release Notes 2024\" becomes \"release-notes 2024\" and the URL keeps a space."}]}
EOF
