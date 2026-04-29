#!/usr/bin/env sh
set -eu

hugo --minify
GOCACHE="${GOCACHE:-$(pwd)/.gocache}" go run scripts/verify-posts.go
