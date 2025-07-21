#!/bin/bash

export BASE_DIR=$(pwd)
export DEVOPS_DIR=".example-dir"

gotestsum $@ ./...
