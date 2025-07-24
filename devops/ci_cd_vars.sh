#!/bin/bash

set -euxo pipefail

ir_url="https://registry.example.com"
ir_project="<project>"

export IMAGE_NAME="<image>"
export COMMIT_HASH="${CI_COMMIT_SHA:-$(git rev-parse HEAD)}"

export IMAGE_URL="${ir_url}/${ir_project}/${IMAGE_NAME}:${COMMIT_HASH}"
