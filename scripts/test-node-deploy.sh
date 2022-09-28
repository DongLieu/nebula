#!/bin/bash

KEY="test0"
TEST_KEY_AMOUNT=5
CHAINID="nebula-1"
KEYRING="test"
MONIKER="localtestnet"
KEYALGO="secp256k1"
LOGLEVEL="info"

# retrieve all args
WILL_RECOVER=0
WILL_INSTALL=0
WILL_CONTINUE=0
# $# is to check number of arguments
if [ $# -gt 0 ];
then
    # $@ is for getting list of arguments
    for arg in "$@"; do
        case $arg in
        --recover)
            WILL_RECOVER=1
            shift
            ;;
        --install)
            WILL_INSTALL=1
            shift
            ;;
        --continue)
            WILL_CONTINUE=1
            shift
            ;;
        *)
            printf >&2 "wrong argument somewhere"; exit 1;
            ;;
        esac
    done
fi

# continue running if everything is configured
if [ $WILL_CONTINUE -eq 1 ];
then
    # Start the node (remove the --pruning=nothing flag if historical queries are not needed)
    nebulad start --pruning=nothing --log_level $LOGLEVEL --minimum-gas-prices=0.0001unebula
    exit 1;
fi

# validate dependencies are installed
command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }
command -v toml > /dev/null 2>&1 || { echo >&2 "toml not installed. More info: https://github.com/mrijken/toml-cli"; exit 1; }

# install nebulad if not exist
if [ $WILL_INSTALL -eq 0 ];
then 
    command -v nebulad > /dev/null 2>&1 || { echo >&1 "installing nebulad"; make install; }
else
    echo >&1 "installing nebulad"
    rm -rf $HOME/.nebula*
    make install
fi

nebulad config keyring-backend $KEYRING
nebulad config chain-id $CHAINID

# determine if user wants to recorver or create new
if [ $WILL_RECOVER -eq 0 ];
then
    nebulad keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO
else
    nebulad keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --recover
fi

echo >&1 "\n"

# init chain
nebulad init $MONIKER --chain-id $CHAINID

# Change parameter token denominations to unebula
cat $HOME/.nebula/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="unebula"' > $HOME/.nebula/config/tmp_genesis.json && mv $HOME/.nebula/config/tmp_genesis.json $HOME/.nebula/config/genesis.json
cat $HOME/.nebula/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="unebula"' > $HOME/.nebula/config/tmp_genesis.json && mv $HOME/.nebula/config/tmp_genesis.json $HOME/.nebula/config/genesis.json
cat $HOME/.nebula/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="unebula"' > $HOME/.nebula/config/tmp_genesis.json && mv $HOME/.nebula/config/tmp_genesis.json $HOME/.nebula/config/genesis.json
cat $HOME/.nebula/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="unebula"' > $HOME/.nebula/config/tmp_genesis.json && mv $HOME/.nebula/config/tmp_genesis.json $HOME/.nebula/config/genesis.json

# Set gas limit in genesis
# cat $HOME/.nebula/config/genesis.json | jq '.consensus_params["block"]["max_gas"]="10000000"' > $HOME/.nebula/config/tmp_genesis.json && mv $HOME/.nebula/config/tmp_genesis.json $HOME/.nebula/config/genesis.json

# enable rest server and swagger
toml set --toml-path $HOME/.nebula/config/app.toml api.swagger true
toml set --toml-path $HOME/.nebula/config/app.toml api.enable true

# Allocate genesis accounts (cosmos formatted addresses)
nebulad add-genesis-account $KEY 1000000000000000unebula,1000000000000uusdt --keyring-backend $KEYRING

for i in $(seq 1 $TEST_KEY_AMOUNT)
do
    nebulad keys add $(printf "%s%d" "$KEY" "$i") --keyring-backend $KEYRING --algo $KEYALGO
    nebulad add-genesis-account $(printf "%s%d" "$KEY" "$i") 1000000000000unebula,1000000000000000uusdt --keyring-backend $KEYRING
done

# Sign genesis transaction
nebulad gentx $KEY 1000000unebula --keyring-backend $KEYRING --chain-id $CHAINID

# Collect genesis tx
nebulad collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
nebulad validate-genesis

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
nebulad start --pruning=nothing --log_level $LOGLEVEL --minimum-gas-prices=0.0001unebula
