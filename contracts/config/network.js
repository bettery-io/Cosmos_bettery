const key = require("./key");
const networkProvider = `https://ropsten.infura.io/v3/${key.infuraKey}`;
const networkId = 3;

module.exports = {
    networkProvider,
    networkId
}