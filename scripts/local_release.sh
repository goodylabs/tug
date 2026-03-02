#!/bin/bash

mkdir -p bin

export TUG_ENV=production

go build -o bin/tug main.go

mv bin/tug "$HOME/.tug/bin/tug"
