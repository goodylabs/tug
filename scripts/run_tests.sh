#!/bin/bash

export TUG_ENV="testing"

gotestsum $@ ./tests/...
