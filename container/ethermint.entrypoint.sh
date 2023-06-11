#!/bin/bash

export NAMESPACE_ID=$(openssl rand -hex 8)

DA_CONFIG='{
  "http_post_mode": true,
  "disable_tls": true
}'
DA_CONFIG=$(echo "$DA_CONFIG" | jq --arg host "$RPCHOST:$RPCPORT" '.host = $host' | tr -d '\n' | tr -d ' ')
DA_CONFIG=$(echo "$DA_CONFIG" | jq --arg rpcuser "$RPCUSER" '.user = $rpcuser' | tr -d '\n' | tr -d ' ')
DA_CONFIG=$(echo "$DA_CONFIG" | jq --arg rpcpass "$RPCPASS" '.pass = $rpcpass' | tr -d '\n' | tr -d ' ')
echo "da_config:"
echo $DA_CONFIG | jq

# while :
# do
#   sleep 1
# done

sed -i 's/api = .*/api = "eth,net,web3,personal"/' /root/.ethermintd/config/app.toml

ethermintd keys add alice --keyring-backend test --algo eth_secp256k1
ethermintd keys add bob --keyring-backend test --algo eth_secp256k1
ethermintd keys add charlie --keyring-backend test --algo eth_secp256k1

ethermintd start \
  --rollkit.aggregator true \
  --rollkit.da_layer bitcoin \
  --rollkit.da_config=$DA_CONFIG \
  --rollkit.namespace_id $NAMESPACE_ID \
  --rollkit.da_start_height 1
