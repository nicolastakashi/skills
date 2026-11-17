#!/bin/bash
cp /solution/percentile.go /app/percentile.go
cp /solution/percentile_test.go /app/percentile_test.go
cd /app && go build ./...
