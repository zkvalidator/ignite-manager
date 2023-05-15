#!/bin/bash

DAEMON="$1"
NAMESPACE_ID="$2"
BRIDGE_ADDRESS="$3"
APP_ADDRESS="$4"

if [ -z "$NAMESPACE_ID" ] || [ "$NAMESPACE_ID" == "None" ]; then
  NAMESPACE_ID=$(openssl rand -hex 8)
fi
echo "namespace: $NAMESPACE_ID"

DA_CONFIG='{"base_url":"","timeout":60000000000,"fee":6000,"gas_limit":6000000}'
DA_CONFIG=$(echo "$DA_CONFIG" | jq --arg base_url "$BRIDGE_ADDRESS:26659" '.base_url = $base_url' | tr -d '\n' | tr -d ' ')
echo "da_config: $DA_CONFIG"

DA_BLOCK_HEIGHT=""
while [ -z "$DA_BLOCK_HEIGHT" ] || [ "$DA_BLOCK_HEIGHT" == "null" ] || [ "$DA_BLOCK_HEIGHT" -lt 2 ]; do
  echo "waiting for block height..."
  sleep 1
  # DA_BLOCK_HEIGHT=$(curl -s http://0.0.0.0:26650/block | jq -r '.result.block.header.height')
  DA_BLOCK_HEIGHT=$(curl -s "$APP_ADDRESS:26657/block" | jq -r '.result.block.header.height')
done
echo "block height: $DA_BLOCK_HEIGHT"

$DAEMON start \
  --rollkit.aggregator true \
  --rollkit.da_layer celestia \
  --rollkit.da_config=$DA_CONFIG \
  --rollkit.namespace_id $NAMESPACE_ID \
  --rollkit.da_start_height $DA_BLOCK_HEIGHT

# b9529aa114a1f70a
