[![Build and run tests](https://github.com/loikg/hedera-cli/actions/workflows/ci.yml/badge.svg?event=push)](https://github.com/loikg/hedera-cli/actions/workflows/ci.yml)

# hedera-cli

This project is a work in progress and is not production ready. Use at your own risks.

hedera-cli make it easy to interact with the hedera blockchain from the command line.
It can connect to a local hedera node, testnet and mainnet.
Operator and network can be configured in the config file located at $HOME/.hedera-cli.yaml by default.

## Install

```sh
go install github.com/loikg/hedera-cli@latest
```

## Usage

```
Usage:
  hedera-cli [command]

Available Commands:
  account     Create, update, delete accounts
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  token       Create, update, delete fungible and non fungible tokens

Flags:
      --config string   config file (default is $HOME/.hedera-cli.yaml)
  -h, --help            help for hedera-cli

Use "hedera-cli [command] --help" for more information about a command.
```
