const path = require("../../config/path");
const config = require("../../config/cosmosConfig");
const axios = require("axios");

const getEventById = async (req, res) => {
    var id = req.params.id;
    let send = await axios.get(path.path + `/privateevents/${id}`, config.header
    ).catch((err) => {
        console.log(err.response.data)
        res.status(400);
        res.send(err.response.data);
    })
    if (send) {
        res.status(200);
        res.send(send.data);
    }
}

module.exports = {
    getEventById
}