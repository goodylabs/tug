#!/bin/bash

set -euo pipefail

mkdir -p "$HOME/.tug/bin"

TUG_BIN_PATH="$HOME/.tug/bin"

if ! echo "$PATH" | grep -q "$TUG_BIN_PATH"; then
    echo "Please add the following line to your shell configuration file:"
    echo "export PATH=\"\$PATH:$TUG_BIN_PATH\""
    exit 0
fi

releaseUrl="https://api.github.com/repos/goodylabs/tug/releases/latest"

os_type=$(uname -s | tr '[:upper:]' '[:lower:]')
artifact_url=$(curl -s $releaseUrl | jq -r ".assets[] | select(.name | test(\"tug-${os_type}-amd64\"))" | jq ".browser_download_url" -r)

curl -Ls "$artifact_url" -o "$TUG_BIN_PATH/tug"

chmod +x "$TUG_BIN_PATH/tug"

echo "Tug has been installed successfully!"
