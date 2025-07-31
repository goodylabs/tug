#!/bin/bash

mkdir -p bin

TAG="${GITHUB_REF##*/}"

x_flags=(
    "-X github.com/goodylabs/tug/cmd.version=$TAG"
    "-X github.com/goodylabs/tug/cmd.useDockerCmd=false"
)

ldflags=$(IFS=" "; echo "${x_flags[*]}")

export TUG_ENV="production"

GOOS=linux GOARCH=amd64 go build -ldflags "$ldflags" -o bin/tug-linux-amd64
GOOS=linux GOARCH=arm64 go build -ldflags "$ldflags" -o bin/tug-linux-arm64
GOOS=darwin GOARCH=amd64 go build -ldflags "$ldflags" -o bin/tug-darwin-amd64
GOOS=darwin GOARCH=arm64 go build -ldflags "$ldflags" -o bin/tug-darwin-arm64
