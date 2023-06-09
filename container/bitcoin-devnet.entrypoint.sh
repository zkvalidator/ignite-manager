#!/bin/bash

# set -e
# set -x

# Set up the config for regtest (local network)
  # -rpcport=$RPCPORT \
  # -rest=1 \
bitcoind \
  -chain=$CHAIN \
  -rpcbind=0.0.0.0:$RPCPORT \
  -rpcallowip=0.0.0.0/0 \
  -rpcuser=$RPCUSER \
  -rpcpassword=$RPCPASS \
  -fallbackfee=0.000001 \
  -txindex=1 \
  &

sleep 3

# create a wallet
bitcoin-cli \
  -$CHAIN \
  -rpcport=$RPCPORT \
  -rpcuser=$RPCUSER \
  -rpcpassword=$RPCPASS \
  createwallet w1

export COINBASE=$(\
  bitcoin-cli \
  -$CHAIN \
  -rpcport=$RPCPORT \
  -rpcuser=$RPCUSER \
  -rpcpassword=$RPCPASS \
  getnewaddress \
)

# generate a new address and mine 101 blocks
bitcoin-cli \
  -$CHAIN \
  -rpcport=$RPCPORT \
  -rpcuser=$RPCUSER \
  -rpcpassword=$RPCPASS \
  generatetoaddress 101 $COINBASE

export ADDRESS=$(\
  bitcoin-cli \
  -$CHAIN \
  -rpcport=$RPCPORT \
  -rpcuser=$RPCUSER \
  -rpcpassword=$RPCPASS \
  getnewaddress \
)

# mine a block every second
while :
do
  echo "Generate a new block `date '+%d/%m/%Y %H:%M:%S'`"
  bitcoin-cli \
    -$CHAIN \
    -rpcport=$RPCPORT \
    -rpcuser=$RPCUSER \
    -rpcpassword=$RPCPASS \
    generatetoaddress 1 $ADDRESS
  sleep 1
done
