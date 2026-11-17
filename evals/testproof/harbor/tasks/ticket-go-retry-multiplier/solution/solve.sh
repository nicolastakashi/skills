#!/bin/bash
cp /solution/policy.go /app/policy.go
cp /solution/policy_test.go /app/policy_test.go
cd /app && go build ./...
