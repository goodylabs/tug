#!/bin/bash

mkdir -p bin

go build -o bin/tug main.go

mv bin/tug "$HOME/.tug/bin/tug"
