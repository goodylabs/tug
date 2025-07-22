#!/bin/bash

export BASE_DIR=$(pwd)
export USE_MOCKS=true
export DEVOPS_DIR=".example-dir"

gotestsum $@ ./tests/...
