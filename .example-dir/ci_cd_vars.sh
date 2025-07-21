#!/bin/bash

set -euxo pipefail

ir_url="ir.goodylabs.com"
ir_project="work.life"

export IMAGE_NAME="crappy"
export COMMIT_HASH="${CI_COMMIT_SHA:-$(git rev-parse HEAD)}"

export IMAGE_URL="${ir_url}/${ir_project}/${IMAGE_NAME}:${COMMIT_HASH}"
