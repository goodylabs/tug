# TUG

![tug](https://raw.githubusercontent.com/goodylabs/tug/refs/heads/main/assets/images/tug-cli-logo-256x256.png)

TUG is a CLI tool that helps you monitor and manage remote environments for a given repo, when you don’t want to bother with manually typing in commands.

## Installation

1. Install jq:

```bash
brew install jq      # macOS
sudo apt install jq  # Ubuntu
sudo pacman -S jq    # Arch Linux
sudo dnf install jq  # Red Hat/CentOS
```

2. Prepare your environment:

```bash
rc_file=${HOME}/.$(basename "$SHELL")rc
echo 'export PATH="$HOME/.tug/bin:$PATH"' >> $rc_file
source $rc_file
```

3. Install TUG:

```bash
curl -s https://raw.githubusercontent.com/goodylabs/tug/refs/heads/main/scripts/download.sh | bash -s
```

4. Configure TUG for ssh connections:

```bash
tug configure
```

# Commands

![tug](https://raw.githubusercontent.com/goodylabs/tug/refs/heads/main/assets/images/tug-help.png)
