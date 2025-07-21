#!/bin/bash

set -euo pipefail

source ./devops/ci_cd_vars.sh

TARGET_IP="127.0.0.1"
DOCKER_STACK_FILE="devops/localhost/docker-stack.yml"
STACK_NAME="crappy"
SERVICES_IN_STACK=("app" "cron")
