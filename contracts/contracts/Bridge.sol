pragma solidity >=0.4.22 <0.8.0;

contract Bridge {
    struct Deposited {
       address owner;
       uint amount;
    }
    mapping(address => Deposited) deposited;
    address[] public depositedAccts;

    event Deposit(address sender, uint amount);

    function depositEth() public payable{
        Deposited memory depos = deposited[msg.sender];
        depos.owner = msg.sender;
        depos.amount = msg.value;
        depositedAccts.push(msg.sender);
        emit Deposit(msg.sender, msg.value);
    }
}
