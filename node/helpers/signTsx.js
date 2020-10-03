const sig = require('@tendermint/sig');
const { readFileSync } = require('fs');
const path = require('path');
const cocmosCong = require('../config/cosmosConfig');
const demonPatn = require('../config/path')
const axios = require("axios");


const getWalet = () => {
    const mnemonic = readFileSync(path.join(__dirname, '../secret'), 'utf-8')
    return sig.createWalletFromMnemonic(mnemonic);
}

const signTransaction = async (tx) => {
    const wallet = getWalet();
    let { sequence, account_number } = await getSequence()

    const signMeta = {
        account_number: String(account_number),
        chain_id: cocmosCong.chain_id,
        sequence: String(sequence)
    };

    return sig.signTx(tx.value, signMeta, wallet);
}


const getSequence = async () => {
    let wallet = getWalet()
    let send = await axios.get(demonPatn.path + '/auth/accounts/' + wallet.address
    ).catch((err) => {
        console.log(err.response.data.error);
    })
    if (send) {
        let sequence = send.data.result.value.sequence;
        let account_number = send.data.result.value.account_number;
        return { sequence, account_number }
    }
}

module.exports = {
    signTransaction
}

