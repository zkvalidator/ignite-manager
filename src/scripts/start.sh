#!/bin/bash

DAEMON="$1"

NAMESPACE_ID=$(openssl rand -hex 8)
echo "namespace: $NAMESPACE_ID"
DA_BLOCK_HEIGHT=$(curl http://0.0.0.0:26650/block | jq -r '.result.block.header.height')
echo "block height: $DA_BLOCK_HEIGHT"

$DAEMON start \
  --rollkit.aggregator true \
  --rollkit.da_layer celestia \
  --rollkit.da_config='{"base_url":"http://localhost:26659","timeout":60000000000,"fee":6000,"gas_limit":6000000}' \
  --rollkit.namespace_id $NAMESPACE_ID \
  --rollkit.da_start_height $DA_BLOCK_HEIGHT
