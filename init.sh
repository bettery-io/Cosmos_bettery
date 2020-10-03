#!/bin/bash
rm -r ~/.betterycli
rm -r ~/.betteryd

betteryd init local --chain-id=bettery

betterycli config keyring-backend test
betterycli config chain-id bettery
betterycli config output json
betterycli config indent true
betterycli config trust-node true

betterycli keys add validator
betterycli keys add me
betterycli keys add you
betterycli keys add she
betterycli keys add he
betterycli keys add we
betterycli keys add they

betteryd add-genesis-account $(betterycli keys show validator -a) 1000foo,100000000stake
betteryd add-genesis-account $(betterycli keys show me -a) 1000foo
betteryd add-genesis-account $(betterycli keys show you -a) 1000foo
betteryd add-genesis-account $(betterycli keys show she -a) 1000foo
betteryd add-genesis-account $(betterycli keys show he -a) 1000foo
betteryd add-genesis-account $(betterycli keys show we -a) 1000foo
betteryd add-genesis-account $(betterycli keys show they -a) 1000foo

betteryd gentx --name validator --keyring-backend test
betteryd collect-gentxs