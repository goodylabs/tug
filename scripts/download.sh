mkdir -p ~/.tug/bin

TUG_BIN_PATH="$HOME/.tug/bin"
TUG_BIN_URL="https://raw.githubusercontent.com/goodylabs/tug/refs/heads/main/scripts/example?token=GHSAT0AAAAAADHX3RHTIOX72H5Z2424ULSQ2ECC73A"

if [[ ":$PATH:" != *":$TUG_BIN_PATH:"* ]]; then

    echo "Adding $TUG_BIN_PATH to PATH"

    if [ "$SHELL" == "/bin/bash" ]; then
        shel_rc_file="$HOME/.bashrc"

    elif [ "$SHELL" == "/bin/zsh" ]; then
        shel_rc_file="$HOME/.zshrc"

    else
        echo "Unsupported shell. Please add the following line to your shell configuration file manually:"
        echo "export PATH=\"\$PATH:\$TUG_BIN_PATH\""
        exit 1
    fi

    echo "export PATH=\"\$PATH:\$TUG_BIN_PATH\"" >> "$shel_rc_file"
    echo "Added to $shel_rc_file"
fi

curl $TUG_BIN_URL > /usr/local/bin/tug

chmod +x /usr/local/bin/tug

if [[ "$(uname)" == "Darwin" ]]; then
    xattr -d com.apple.quarantine /usr/local/bin/tug || echo "No need to remove quarantine attribute"
fi
