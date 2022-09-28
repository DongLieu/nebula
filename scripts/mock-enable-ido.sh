#!/bin/bash

KEY="test0"
CHAINID="nebula-1"
KEYRING="test"
SLEEP_TIME="10s"
API="localhost:1317"

# Enable IDO corresponding to project id
PROJECT_ID=$1
TOKENS="100000000unebula"
LISTING_PRICE="1000000uusdt"
ALLOCATION_RANGE="10000000uusdt-100000000uusdt"
START_TIME=$(date --date="+5 minutes" --rfc-3339=seconds | sed 's/ /T/')

echo $START_TIME

RES=$(nebulad tx ido enable-ido $PROJECT_ID $TOKENS $LISTING_PRICE "$ALLOCATION_RANGE" --start-time $START_TIME --from $KEY --chain-id $CHAINID --keyring-backend $KEYRING -y --gas auto --fees 100unebula -o json)
echo $RES

# sleep for chain to update
sleep "$SLEEP_TIME"

RAW_LOG=$(nebulad query tx "$(echo $RES | jq -r .txhash)" --chain-id "$CHAINID" -o json | jq -r .raw_log)
echo "RAW LOG = $RAW_LOG"

# query information on api
API_RES=$(curl --location --request GET "$API/nebula-labs/nebula/ido/$PROJECT_ID")
echo "API RESPOND = $API_RES"

# query information on project again
echo "============ PROJECT ============"
nebulad q launchpad project $PROJECT_ID

# query information on ido
echo "============ IDO ============"
nebulad q ido ido $PROJECT_ID