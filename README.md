[![Build and run tests](https://github.com/loikg/hedera-cli/actions/workflows/ci.yml/badge.svg?event=push)](https://github.com/loikg/hedera-cli/actions/workflows/ci.yml)

# hedera-cli

This project is a work in progress and is not production ready. Use it at your own risks.

hedera-cli make it easy to interact with the hedera blockchain from the command line.
It can connect to a local hedera node, testnet and mainnet.

## Install

```sh
go install github.com/loikg/hedera-cli@latest
```

## Usage

```
NAME:
   hedera-cli - hedera-cli make it easy to interact with the hedera blockchain

USAGE:
   hedera-cli [global options] command [command options] [arguments...]

DESCRIPTION:
   hedera-cli make it easy to interact with the hedera blockchain form the command line.
   It can connect to a local hedera node, testnet and mainnet.
   Operator and network can be configured inYou can either provide the
   --network, --operator-id, --operator-key or use the HEDERA_NETWORK, HEDERA_OPERATOR_ID,
   HEDERA_OPERATOR_KEY environment variables.

COMMANDS:
   account, a  Manage hedera accounts
   keygen, kg  Create a private key.
   token, t    Create, update, delete fungible and non fungible tokens
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --network value       Hedera network to connect to, possible values are local,testnet,mainnet [$HEDERA_NETWORK, $NETWORK]
   --operator-id value   Hedera operator account id [$HEDERA_OPERATOR_ID, $OPERATOR_ID]
   --operator-key value  Hedera operator account private key [$HEDERA_OPERATOR_KEY, $OPERATOR_KEY]
   --help, -h            show help
```

## Examples

### Create an account

```shell
export HEDERA_NETWORK=local
export HEDERA_OPERATOR_ID=0.0.1022
export HEDERA_OPERATOR_KEY=a608e2130a0a3cb34f86e757303c862bee353d9ab77ba4387ec084f881d420d4

hedera-cli account create
```

### Create a fungible token

```shell
export HEDERA_NETWORK=local
export HEDERA_OPERATOR_ID=0.0.1022
export HEDERA_OPERATOR_KEY=a608e2130a0a3cb34f86e757303c862bee353d9ab77ba4387ec084f881d420d4

hedera-cli token create \
		--balance 100 \
		--decimals 2 \
		--name myToken \
		--symbol MTK \
		--type ft \
		--supply-type infinite \
		--treasury-id 0.0.1022 \
		--treasury-key $HEDERA_OPERATOR_KEY
```
