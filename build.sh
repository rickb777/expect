#!/bin/bash -e
cd "$(dirname $0)"
unset GOPATH

function v
{
  echo "$@"
  "$@"
}

v go mod download
v go mod tidy

v gofmt -l -w -s *.go

v go test ./...

v go vet ./...
