// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";

contract Accum is Ownable{
    constructor (address initialOwner) Ownable(initialOwner){}

    event Adding(address _from, uint _amount);
    event Exceed(uint _amount, string _message);
    uint constant PLANK = 13 ether;

    function send(address sender, uint amount) internal {
        emit Adding(sender, amount);
        if(address(this).balance >= PLANK){
            emit Exceed(address(this).balance, "Plank done!");
        }
    }

    receive() external payable {
        send(msg.sender, msg.value);
    }

    function sendEthers() external payable{
        send(msg.sender, msg.value);
    }

    function withdrawEther() external onlyOwner{
        require(address(this).balance >= PLANK, "Balance lower than plank!");
        payable(owner()).transfer(address(this).balance);
    }

    function checkBalance()external view returns(uint){
        return address(this).balance;
    }
}