#!/bin/bash

KEYRING="test"

rm -rf $HOME/.relayer
rly config init
rly chains add -f relayer/chain-baby.json baby
rly chains add -f relayer/chain-nebula.json nebula

BABY_IBC_KEY=$(rly keys add baby ibc-key | jq -r .address)
NEBULA_IBC_KEY=$(rly keys add nebula ibc-key | jq -r .address)

echo $BABY_IBC_KEY
echo $NEBULA_IBC_KEY

# Fund accounts
babyd tx bank send $(babyd keys show test -a --keyring-backend $KEYRING) "$BABY_IBC_KEY" 500000000ubaby --gas auto --fees 100ubaby --node http://localhost:2281 -y

sleep 5s

nebulad tx bank send $(nebulad keys show test0 -a --keyring-backend $KEYRING) "$NEBULA_IBC_KEY" 50000000unebula --gas auto --fees 100unebula --node http://localhost:26657 -y

sleep 5s

rly q balance baby

sleep 5s

rly q balance nebula

rly paths delete demo_transfer
rly paths new baby-1 nebula-1 demo_transfer

rly transact link demo_transfer

rly paths list

rly start demo_transfer