#!/bin/bash

set -euo pipefail

source ./devops/ci_cd_vars.sh

TARGET_IP="64.226.87.6"
DOCKER_STACK_FILE="devops/staging/docker-stack.yml"
STACK_NAME="crappy"
SERVICES_IN_STACK=("app" "cron")
