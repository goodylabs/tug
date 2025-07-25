#!/bin/bash

set -euo pipefail

source ./devops/ci_cd_vars.sh

TARGET_IP="<ip_address_production>"
DOCKER_STACK_FILE="devops/production/docker-stack.yml"
STACK_NAME="ipsum"
SERVICES_IN_STACK=("app" "cron")
