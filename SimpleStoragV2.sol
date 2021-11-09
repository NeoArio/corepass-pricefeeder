// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.4.16 <0.9.0;

contract SimpleStorage {
    uint storedData;
    event ItemSet(string key, uint value);

    function set(uint x) public {
        emit ItemSet("event", x);
        storedData = x;
        
    }

    function get() public view returns (uint) {
        return storedData;
    }
}
