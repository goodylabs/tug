#!/bin/bash

set -euo pipefail

source ./devops/ci_cd_vars.sh

IP_ADDRESSES=("ip_1" "ip_2" "ip_3")
DOCKER_STACK_FILE="devops/production/docker-stack.yml"
STACK_NAME="ipsum"
SERVICES_IN_STACK=("app" "cron")
