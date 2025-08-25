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
arch=$(uname -m)

case "$arch" in
    x86_64)
        arch="amd64"
        ;;
    aarch64 | arm64)
        arch="arm64"
        ;;
    *)
        echo "Unsupported architecture: $arch"
        exit 1
        ;;
esac

artifact_url=$(curl -s "$releaseUrl" | jq -r ".assets[] | select(.name | test(\"${os_type}-${arch}\")) | .browser_download_url")

if [[ -z "$artifact_url" ]]; then
    echo "No compatible binary found for ${os_type}-${arch}"
    exit 1
fi

echo "Downloading Tug binary for ${os_type}-${arch}..."

curl --fail --location --progress-bar --compressed --retry 3 --retry-delay 5 \
  --max-time 10 -o "$TUG_BIN_PATH/tug" "$artifact_url"

chmod +x "$TUG_BIN_PATH/tug"

echo "Tug has been installed successfully!"
