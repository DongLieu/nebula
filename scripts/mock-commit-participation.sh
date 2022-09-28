#!/bin/bash

KEY="test2"
CHAINID="nebula-1"
KEYRING="test"
SLEEP_TIME="10s"

# Participation information
PROJECT_ID=$1
TOKEN=100000000uusdt

RES=$(nebulad tx ido commit-participation $PROJECT_ID $TOKEN --from $KEY --chain-id $CHAINID --keyring-backend $KEYRING -y --gas auto --fees 100unebula -o json)
echo $RES

# sleep for chain to update
sleep "$SLEEP_TIME"

RAW_LOG=$(nebulad query tx "$(echo $RES | jq -r .txhash)" --chain-id "$CHAINID" -o json | jq -r .raw_log)
echo $RAW_LOG

# query information on ido
echo "============ IDO ============"
nebulad q ido ido $PROJECT_ID

# query information of account
echo "============ ACCOUNT ============"
nebulad q bank balances $(nebulad keys show $KEY -a --keyring-backend $KEYRING)