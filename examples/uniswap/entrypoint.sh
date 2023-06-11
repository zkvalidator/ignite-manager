#!/bin/bash

update() {
  cat installation_data.json \
    | jq "$1 = $2" \
    > installation_data.json.tmp \
    && mv installation_data.json.tmp installation_data.json
}

rpc_call() {
  curl -s -X POST --data "{\"jsonrpc\":\"2.0\",\"method\":\"$1\",\"params\":[$2],\"id\":1}" -H "Content-Type: application/json" $RPC_ENDPOINT \
    | jq -r '.result'
#   curl -s -X POST --data '{"jsonrpc":"2.0","method":"personal_newAccount","params":["alice"],"id":1}' -H "Content-Type: application/json" localhost:8545
#   curl -s -X POST --data '{"jsonrpc":"2.0","method":"eth_sendTransaction","params":[{"from":"0xd89a841cfc40bd754bc2986210686042c3e7e50d","to":"0x6a28cbe1371f2f845c64ecc624492c0ae060b71a","value":"0x100000"}],"id":1}' -H "Content-Type: application/json" localhost:8545
}

PUBLIC_KEY_MYKEY=$(rpc_call eth_accounts | jq -r '.[0]')
echo "mykey's pubkey: $PUBLIC_KEY_MYKEY"

# PRIVATE_KEY_ALICE=$(rpc_call personal_newAccount "alice")
PRIVATE_KEY_ALICE=$(docker exec ethermint ethermintd keys unsafe-export-eth-key alice --keyring-backend test)
PUBLIC_KEY_ALICE=$(rpc_call eth_accounts | jq -r '.[1]')
echo "alice's privkey: $PRIVATE_KEY_ALICE"
echo "alice's pubkey: $PUBLIC_KEY_ALICE"

# create account bob
PRIVATE_KEY_BOB=$(docker exec ethermint ethermintd keys unsafe-export-eth-key bob --keyring-backend test)
PUBLIC_KEY_BOB=$(rpc_call eth_accounts | jq -r '.[2]')
echo "bob's privkey: $PRIVATE_KEY_BOB"
echo "bob's pubkey: $PUBLIC_KEY_BOB"

# create account charlie
PRIVATE_KEY_CHARLIE=$(docker exec ethermint ethermintd keys unsafe-export-eth-key charlie --keyring-backend test)
PUBLIC_KEY_CHARLIE=$(rpc_call eth_accounts | jq -r '.[3]')
echo "charlie's privkey: $PRIVATE_KEY_CHARLIE"
echo "charlie's pubkey: $PUBLIC_KEY_CHARLIE"

# fund bob account
rpc_call eth_sendTransaction "{\"from\":\"$PUBLIC_KEY_ALICE\",\"to\":\"$PUBLIC_KEY_BOB\",\"value\":\"0x100000\"}"
rpc_call eth_getBalance "$PUBLIC_KEY_BOB"

# fund charlie account
rpc_call eth_sendTransaction "{\"from\":\"$PUBLIC_KEY_ALICE\",\"to\":\"$PUBLIC_KEY_CHARLIE\",\"value\":\"0x100000\"}"
rpc_call eth_getBalance "$PUBLIC_KEY_CHARLIE"

cd /app/uniswap-interface/how_to_deploy_uniswap

update ".private_key.alice" "$PRIVATE_KEY_ALICE"
update ".public_key.alice" "$PUBLIC_KEY_ALICE"
update ".private_key.bob" "$PRIVATE_KEY_BOB"
update ".public_key.bob" "$PUBLIC_KEY_BOB"
update ".private_key.charlie" "$PRIVATE_KEY_CHARLIE"
update ".public_key.charlie" "$PUBLIC_KEY_CHARLIE"

update ".provider.rpc_endpoint" "$RPC_ENDPOINT"

cd /app/uniswap-interface/how_to_deploy_uniswap/uniswap_v2
node deploy_uniswap_v2.js

cd /app/uniswap-interface/how_to_deploy_uniswap/uniswap_interface
python3 modify_addresses.py
cd /app/uniswap-interface
yarn
