# TUG

![tug](https://raw.githubusercontent.com/goodylabs/tug/refs/heads/modules-refactor/docs/assets/images/tug-cli-logo-256x256.png)

## The main idea

TUG makes common server operations more easy

each TUG server cmd (docker, swarm, pm2, pystrano) consists of 2 steps:
  1. loading server connection config 
  2. making actions on server you selected

## Example flow:

1. `tug docker`
2. select environment (example: stg)
3. select host 
4. select select action
5. if action requires resource, select resource
6. step back using `ctrl + c`
7. step back using `ctrl + c`

## Demo video:

https://github.com/user-attachments/assets/547564e2-e6df-455f-9520-5adf7152d559

## Navigation

- chose option: arrow keys + Enter
- search: "/" + type phrase
- go back: `ctrl + c`

## Installation

### 1. Install jq

```bash
brew install jq      # macOS
sudo apt install jq  # Ubuntu
sudo pacman -S jq    # Arch Linux
sudo dnf install jq  # Red Hat/CentOS
```

### 2. Prepare your environment

bash/zsh

```bash
rc_file=${HOME}/.$(basename "$SHELL")rc
echo 'export PATH="$HOME/.tug/bin:$PATH"' >> $rc_file
source $rc_file
```

fish

```bash
set rc_file $HOME/.(basename $SHELL)rc
echo 'set -gx PATH $HOME/.tug/bin $PATH' >> $rc_file
source $rc_file
```

### 3. Install TUG

```bash
curl -s https://raw.githubusercontent.com/goodylabs/tug/refs/heads/modules-refactor/scripts/download.sh | bash -s
```

### 4. Check TUG version

```bash
tug --version
```

### 5. Configure TUG for ssh connections

```bash
tug configure
```

## Commands

```bash
tug docker
```

### Docker

**Config source: variable in script `./devops/*/deploy.sh`**

```bash
TARGET_IP    (string)
IP_ADDRESS   (string)
IP_ADDRESSES ([]string)

# Examples
TARGET_IP=192.168.1.100
IP_ADDRESS=192.168.1.100
IP_ADDRESSES=(192.168.1.100 192.168.1.101)
```

**Available actions**

![](docs/assets/images/tug-docker-actions.example.png)

### Swarm (Docker)

**Config source: same as Docker**

**Available actions**

![](docs/assets/images/tug-swarm-actions.example.png)

### PM2

**Config source: `./ecosystem.config.js` or `./ecosystem.config.cjs`**

**Available actions**

![](docs/assets/images/tug-pm2-actions.example.png)

### Flag `--host`

Instead of reading config directly from pm2 / docker project you can manually specify host(s) using `--host` flag and then optionally `--user` flag.

```bash
tug docker --host 123.23.7.254 # (default user: root) 
tug docker --host 123.23.7.254 --user ubuntu
```
