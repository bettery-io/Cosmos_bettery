const Bridge = artifacts.require("Bridge");

module.exports = async (deployer, network) => {
    if (network === 'ropsten') {
        return deployer.deploy(Bridge);
    }
}