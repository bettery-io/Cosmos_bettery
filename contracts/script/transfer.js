const { readFileSync } = require('fs');
const path = require('path');
const Web3 = require("web3");
const config = require('../config/network');
const BridgeContract = require('../build/contracts/Bridge.json')

async function connectToNetwork(provider) {
    let web3 = new Web3(provider);
    let privateKey = readFileSync(path.join(__dirname, '../secret'), 'utf-8')
    const prKey = web3.eth.accounts.privateKeyToAccount('0x' + privateKey);
    await web3.eth.accounts.wallet.add(prKey);
    let accounts = await web3.eth.accounts.wallet;
    let account = accounts[0].address;
    return { web3, account };
}

async function connectToContract() {
    let provider = config.networkProvider;
    let networkId = config.networkId;
    let { web3, account } = await connectToNetwork(provider);
    let abi = BridgeContract.abi;
    let address = BridgeContract.networks[networkId].address;
    return new web3.eth.Contract(abi, address, { from: account });
}

(async function depost() {
    let web3 = new Web3();
    let amount = web3.utils.toWei("0.01", 'ether');
    let contract = await connectToContract();    
    let gasEstimate = await contract.methods.depositEth().estimateGas();
    let transfer = await contract.methods.depositEth().send({
        amount: amount,
        gas: gasEstimate
    })
    console.log(transfer)
})();