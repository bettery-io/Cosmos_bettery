const sig = require('@tendermint/sig');
const redis = require("redis");
const util = require('util');
const { readFileSync } = require('fs');
const path = require('path');
const cocmosCong = require('../config/cosmosConfig');

const redisUrl = path.redisUrl;
const client = redis.createClient(redisUrl);
client.hget = util.promisify(client.hget);

const signTransaction = async (tx) => {
    const mnemonic = readFileSync(path.join(__dirname, '../secret'), 'utf-8')
    const wallet = sig.createWalletFromMnemonic(mnemonic);

    const value = await client.hget(cocmosCong.chain_id, "sequence");
    let sequence = value ? value : "0";

    const signMeta = {
        account_number: '1',
        chain_id: 'cosmos',
        sequence: sequence
    };

    sequence = Number(sequence) + 1;

    client.hset(cocmosCong.chain_id, "sequence", String(sequence));

    return sig.signTx(tx, signMeta, wallet);
}

module.exports = {
    signTransaction
}

