#!/bin/bash

DAEMON="$1"

NAMESPACE_ID=$(openssl rand -hex 8)
echo "namespace: $NAMESPACE_ID"

DA_BLOCK_HEIGHT=""
while [ -z "$DA_BLOCK_HEIGHT" ] || [ "$DA_BLOCK_HEIGHT" == "null" ] || [ "$DA_BLOCK_HEIGHT" -lt 2 ]; do
  echo "waiting for block height..."
  sleep 1
  # DA_BLOCK_HEIGHT=$(curl -s http://0.0.0.0:26650/block | jq -r '.result.block.header.height')
  DA_BLOCK_HEIGHT=$(curl -s http://local-celestia-devnet:26657/block | jq -r '.result.block.header.height')
done
echo "block height: $DA_BLOCK_HEIGHT"

$DAEMON start \
  --rollkit.aggregator true \
  --rollkit.da_layer celestia \
  --rollkit.da_config='{"base_url":"http://local-celestia-devnet:26659","timeout":60000000000,"fee":6000,"gas_limit":6000000}' \
  --rollkit.namespace_id $NAMESPACE_ID \
  --rollkit.da_start_height $DA_BLOCK_HEIGHT
