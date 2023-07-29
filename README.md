[![Build and run tests](https://github.com/loikg/hedera-cli/actions/workflows/ci.yml/badge.svg?event=push)](https://github.com/loikg/hedera-cli/actions/workflows/ci.yml)

# hedera-cli

This project is a work in progress and is not production ready. Use it at your own risks.

hedera-cli make it easy to interact with the hedera blockchain from the command line.
It can connect to a local hedera node, testnet and mainnet.
Operators and networks can be configured in the config file located at $HOME/.hedera-cli.yaml by default.

## Install

```sh
go install github.com/loikg/hedera-cli@latest
```

## Configuration

By default hedera-cli looks for a config file in `$HOME/.hedera-cli.yaml`.
This can be overrided with the `--config` flag.

You need to configure the operator account for the networks you wich to interact with.

```yaml
networks:
  local:
    operatorId: "0.0.1022"
    operatorKey: "xxxx"
  testnet:
    operatorId: "0.0.1022"
    operatorKey: "xxxx"
  mainnet:
    operatorId: "0.0.1022"
    operatorKey: "xxxx"
```

By default hedera-cli will use the `local` network unless you specify a network with `--network`.

## Usage

```
hedera-cli make it easy to interact with the hedera blockchain form the command line.
It can connect to a local hedera node, testnet and mainnet.
Operator and network can be configured in the config file located at $HOME/.hedera-cli.yaml by default.

Usage:
  hedera-cli [command]

Available Commands:
  account     Create, update, delete accounts
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  keygen      Create a private key.
  token       Create, update, delete fungible and non fungible tokens
  version     Display version

Flags:
      --config string    config file (default is $HOME/.hedera-cli.yaml)
  -h, --help             help for hedera-cli
      --network string   Network to connect to either local,testnet or mainnet
      --verbose          enable debug mesage useful for debugging

Use "hedera-cli [command] --help" for more information about a command.
```
