#!/bin/bash

export BASE_DIR=$(pwd)
export TESTING=true

gotestsum $@ ./tests/...
