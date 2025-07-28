#!/bin/bash

set -euo pipefail

source ./devops/ci_cd_vars.sh

IP_ADDRESS=unix:///var/run/docker.sock

DOCKER_STACK_FILE="devops/localhost/docker-stack.yml"
STACK_NAME="lorem"
SERVICES_IN_STACK=("app" "cron")
