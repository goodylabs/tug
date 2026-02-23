#!/bin/bash

go run -ldflags="-X 'github.com/goodylabs/tug/pkg/config.TugEnv=development'" main.go $@
