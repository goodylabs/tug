#!/bin/bash

set -euo pipefail

source ./devops/ci_cd_vars.sh

TARGET_IP="<ip_address>"
DOCKER_STACK_FILE="devops/staging/docker-stack.yml"
STACK_NAME="dolor"
SERVICES_IN_STACK=("app" "cron")
