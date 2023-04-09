#!/bin/bash

# VALIDATOR_NAME=distrifab
# CHAIN_ID=distrifab
# KEY_NAME=distrifab
# DAEMON=/Users/peac/go/bin/distrifabd

DAEMON="$1"
CHAIN_ID="$2"
VALIDATOR_NAME="$3"
KEY_NAME="$4"

CHAINFLAG="--chain-id ${CHAIN_ID}"
TOKEN_AMOUNT="10000000000000000000000000stake"
STAKING_AMOUNT="1000000000stake"

$DAEMON tendermint unsafe-reset-all
$DAEMON init $VALIDATOR_NAME --chain-id $CHAIN_ID
$DAEMON keys add $KEY_NAME --keyring-backend test
$DAEMON add-genesis-account $KEY_NAME $TOKEN_AMOUNT --keyring-backend test
$DAEMON gentx $KEY_NAME $STAKING_AMOUNT --chain-id $CHAIN_ID --keyring-backend test
$DAEMON collect-gentxs
exec start.sh
