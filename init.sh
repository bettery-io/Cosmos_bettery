#!/bin/bash
rm -r ~/.betterycli
rm -r ~/.betteryd

betteryd init mynode --chain-id bettery

betterycli config keyring-backend test
betterycli config chain-id bettery
betterycli config trust-node true
betterycli config indent true
betterycli config output json

betterycli keys add validator
betterycli keys add me
betterycli keys add you
betterycli keys add she
betterycli keys add he
betterycli keys add we
betterycli keys add they

betteryd add-genesis-account $(betterycli keys show validator -a) 1000000000stake,10000000000000000000000000000000000000000000000000000000000000000000000000000bet
betteryd add-genesis-account $(betterycli keys show me -a) 10000000000000000000bet
betteryd add-genesis-account $(betterycli keys show you -a) 10000000000000000000bet
betteryd add-genesis-account $(betterycli keys show she -a) 10000000000000000000bet
betteryd add-genesis-account $(betterycli keys show he -a) 10000000000000000000bet
betteryd add-genesis-account $(betterycli keys show we -a) 10000000000000000000bet
betteryd add-genesis-account $(betterycli keys show they -a) 10000000000000000000bet

betteryd gentx --name validator --keyring-backend test
betteryd collect-gentxs