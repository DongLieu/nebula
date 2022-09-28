#!/bin/bash

KEY="test0"
CHAINID="nebula-1"
KEYRING="test"
SLEEP_TIME="10s"

# Project information
PROJECT_TITLE="temp"
PROJECT_INFORMATION="This is a temp project"

RES=$(nebulad tx launchpad create-project $PROJECT_TITLE "$PROJECT_INFORMATION" --from $KEY --chain-id $CHAINID --keyring-backend $KEYRING -y --gas auto --fees 100unebula -o json)
echo $RES

# sleep for chain to update
sleep "$SLEEP_TIME"

RAW_LOG=$(nebulad query tx "$(echo $RES | jq -r .txhash)" --chain-id "$CHAINID" -o json | jq -r .raw_log)

PROJECT_ID=$(echo $RAW_LOG | jq -r ".[0].events[1].attributes[0].value")

echo "PROJECT_ID = $PROJECT_ID"

nebulad q launchpad project $PROJECT_ID