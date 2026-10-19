pragma solidity ^0.7.0;

contract Vulnerable {
    address public owner;
    uint256 public balance;

    function withdraw() public {
        (bool success, ) = msg.sender.call{value: balance}("");
        balance = 0;
    }

    function destroy() public {
        selfdestruct(payable(msg.sender));
    }
}
