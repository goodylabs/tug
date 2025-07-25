#!/bin/bash

set -euo pipefail

source ./devops/ci_cd_vars.sh

IP_ADDRESS="<ip_address_uat>"
DOCKER_STACK_FILE="devops/uat/docker-stack.yml"
STACK_NAME="sit-amet"
SERVICES_IN_STACK=("app" "cron")
