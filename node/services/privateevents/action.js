const path = require("../../config/path");
const config = require("../../config/cosmosConfig");
const axios = require("axios");
const helper = require('../../helpers/signTsx');

const participate = async (req, res) => {
    let send = await axios.post(path.path + '/privateevent/participate', req.body, config.header
    ).catch((err) => {
        res.status(400);
        res.send(err.response.data);
    })

    if (send) {
        let transaction = await helper.signTransaction(send.data);

        send.data.value.signatures = transaction.signatures

        let sendTxs = {
            "tx": send.data.value,
            "mode": "async"
        }

        let txs = await axios.post(path.path + '/txs', sendTxs, config.header
        ).catch((err) => {
            res.status(400);
            res.send(err.response.data);
        })
        if (txs) {
            res.status(200);
            res.send({ "status": "ok" });
        }
    }
}

const validate = async (req, res) => {
    let send = await axios.post(path.path + '/privateevent/validate', req.body, config.header
    ).catch((err) => {
        res.status(400);
        res.send(err.response.data);
    })

    if (send) {
        let transaction = await helper.signTransaction(send.data);

        send.data.value.signatures = transaction.signatures

        let sendTxs = {
            "tx": send.data.value,
            "mode": "async"
        }

        let txs = await axios.post(path.path + '/txs', sendTxs, config.header
        ).catch((err) => {
            res.status(400);
            res.send(err.response.data);
        })
        if (txs) {
            res.status(200);
            res.send({ "status": "ok" });
        }
    }
}

module.exports = {
    participate,
    validate
}