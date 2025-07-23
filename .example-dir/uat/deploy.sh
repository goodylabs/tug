#!/bin/bash

set -euo pipefail

source ./devops/ci_cd_vars.sh

IP_ADDRESS="121.122.123.124"
DOCKER_STACK_FILE="devops/uat/docker-stack.yml"
STACK_NAME="crappy"
SERVICES_IN_STACK=("app" "cron")
