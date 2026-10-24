#!/bin/bash
# Scores the task. The logic lives in verify/ (copied from _shared by sync.sh).
cd /tests/verify && go run .
