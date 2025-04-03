// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26; // solidity ^0.8.0 recommended

import {JackalInterface} from "./JackalInterface.sol"; // import bridge interface

contract StorageDrawer {
    JackalInterface internal jackalBridge; // variable for interface
    mapping(address => string[]) public cabinet; // variable for file owners

    constructor(address _jackalAddress) {
        // constructor takes bridge address
        jackalBridge = JackalInterface(_jackalAddress); // use jackal bridge
    }

    function upload(string memory merkle, uint64 filesize) public payable {
        // takes file and size
        jackalBridge.postFileFrom{value: msg.value}(msg.sender, merkle, filesize, "", 30); // call bridge method
        // method list https://github.com/JackalLabs/mulberry/blob/main/forge/src/JackalInterface.sol
        cabinet[msg.sender].push(merkle); // record file owner
    }

    function fileAddress(address _addr, uint256 _index) public view returns (string memory) {
        // takes address and index
        require(_index < cabinet[_addr].length); // check file exists for address
        return cabinet[_addr][_index]; // return file
    }

    function fileCount(address _addr) public view returns (uint256) {
        // takes address
        return cabinet[_addr].length; // return file count
    }
}
