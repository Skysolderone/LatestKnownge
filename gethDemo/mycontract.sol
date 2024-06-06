//SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract MyContract {
    uint256 private value;

    function setValue(uint256 _value) public {
        value = _value;
    }
    function getValue() public view returns (uint256) {
        return value;
    }
}
