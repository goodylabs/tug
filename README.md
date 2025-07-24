# TUG

![tug](https://raw.githubusercontent.com/goodylabs/tug/refs/heads/main/assets/images/tug-cli-logo-256x256.png)

CLI tool for reducing developers' life pain when they need to work with Docker.

## Installation

```bash
curl https://raw.githubusercontent.com/goodylabs/tug/refs/heads/main/scripts/download.sh | bash -c
```

## Main features

- Interactive container selection and command execution for remote environments

## Commands

### `developer`

```bash
# Connect to a specific environment
tug developer <environment>
# example:
tug developer staging
tug developer production

# Using flag syntax
tug developer --envDir=staging

# Get help
tug --help
tug developer --help
```
