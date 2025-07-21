#!/bin/bash

set -euo pipefail

source ./devops/ci_cd_vars.sh

TARGET_IP="167.99.198.9"
DOCKER_STACK_FILE="devops/production/docker-stack.yml"
STACK_NAME="crappy"
SERVICES_IN_STACK=("app" "cron")
