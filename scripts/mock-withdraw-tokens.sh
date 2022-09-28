#!/bin/bash

KEY="test0"
CHAINID="nebula-1"
KEYRING="test"
SLEEP_TIME="10s"

# Project information
PROJECT_ID=$1

# delete project
RES=$(nebulad tx launchpad withdraw-all-tokens $PROJECT_ID --from $KEY --chain-id $CHAINID --keyring-backend $KEYRING -y --gas auto --fees 100unebula -o json)
echo $RES

# sleep for chain to update
sleep "$SLEEP_TIME"

RAW_LOG=$(nebulad query tx "$(echo $RES | jq -r .txhash)" --chain-id "$CHAINID" -o json | jq -r .raw_log)

echo $RAW_LOG

# query information on ido
echo "============ IDO ============"
nebulad q launchpad project-balances $PROJECT_ID

# query information of account
echo "============ ACCOUNT ============"
nebulad q bank balances $(nebulad keys show $KEY -a --keyring-backend $KEYRING)