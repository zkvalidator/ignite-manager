#!/bin/bash

# anvil --port 9545

# export ANVIL_KEY=<anvil-private-key>

# export RPC_URL=http://127.0.0.1:9545

# forge script script/Counter.s.sol:CounterScript --fork-url \
# $RPC_URL  --private-key $ANVIL_KEY --broadcast

# export CONTRACT_ADDRESS=<contract-address>

set -e
set -x

PRIVATE_KEY=$(docker exec ethermint ethermintd keys unsafe-export-eth-key mykey --keyring-backend test)

# /root/.foundry/bin/forge script \
#   script/Counter.s.sol:CounterScript \
#   --rpc-url $RPC_ENDPOINT \
#   --private-key $PRIVATE_KEY \
#   --broadcast \
#   --verify

OUTPUT=$(/root/.foundry/bin/forge create \
  CounterScript \
  --contracts script/Counter.s.sol \
  --private-key $PRIVATE_KEY \
  --rpc-url $RPC_ENDPOINT)

OUTPUT=$(echo "$OUTPUT" | grep "Deployed to")
OUTPUT=$(echo "$OUTPUT" | cut -d':' -f2 | xargs)
echo $OUTPUT

export CONTRACT_ADDRESS=$OUTPUT

/root/.foundry/bin/cast send \
  $CONTRACT_ADDRESS \
  "incrementCounter()" \
  --rpc-url $RPC_ENDPOINT \
  --private-key $PRIVATE_KEY 

/root/.foundry/bin/cast call \
  $CONTRACT_ADDRESS \
  "getCount()(int)" \
  --rpc-url $RPC_ENDPOINT

python -m http.server --bind 0.0.0.0 9000

