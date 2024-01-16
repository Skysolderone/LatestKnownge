// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract SimpleStorage {
    uint storageData;

    function set(uint x) public {
        storageData = x;
    }

    function get() public view returns (uint) {
        return storageData;
    }
}
