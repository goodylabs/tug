# TUG

![tug](https://raw.githubusercontent.com/goodylabs/tug/refs/heads/main/assets/images/tug-cli-logo-256x256.png)

CLI tool for reducing developers' life pain when they need to work with Docker.

# Installation / Update

```bash
curl -s https://raw.githubusercontent.com/goodylabs/tug/refs/heads/main/scripts/download_script.sh | bash -s
```

# Commands

## Basic

```bash
# Basic
tug --version
tug --help
```

## PM2

Allows to interact with pm2 environments on remote servers, using configuration from `./ecosystem.config.js`.

```bash
# interactive env selection
tug pm2

# already selected env (examples)
tug pm2 staging
tug pm2 staging_RO
tug pm2 staging_RO
```

## Docker

Allows to interact with Docker environments on remote servers, using configuration from `./devops/<env>/deploy.sh`.

```bash
tug docker [env]

# examples:
tug docker staging
tug docker uat
tug docker prod

```
