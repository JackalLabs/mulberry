// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26; // solidity ^0.8.0 recommended

import {JackalInterface} from "./JackalInterface.sol"; // import bridge interface

contract StorageDrawer {
    struct FileDetails {
        address owner;
        string merkle;
        uint start;
    }

    JackalInterface internal jackalBridge; // variable for interface
    mapping(address => FileDetails[]) public cabinet; // variable for file owners

    constructor(address _jackalAddress) {
        // constructor takes bridge address
        jackalBridge = JackalInterface(_jackalAddress); // use jackal bridge
    }

    function upload(string memory merkle, uint64 filesize) public payable {
        upload(merkle, filesize, msg.value);
    }

    function upload(string memory merkle, uint64 filesize, uint value) public payable {
        upload(merkle, filesize, value, 73000); // uploading for 200 years by default
    }

    function upload(string memory merkle, uint64 filesize, uint value, uint64 dayCount) public payable {
        require(dayCount > 30, "day count must be more than 30");
        // takes file and size
        jackalBridge.postFileFrom{value: value}(msg.sender, merkle, filesize, "", dayCount); // call bridge for 200 years
        // method list https://github.com/JackalLabs/mulberry/blob/main/forge/src/JackalInterface.sol

        cabinet[msg.sender].push(FileDetails(msg.sender, merkle, block.number)); // record file owner
    }

    function uploadMany(string[] memory merkle, uint64[] memory filesize, uint[] memory values) public payable {
        for (uint i=0; i<merkle.length; i++) {
            upload(merkle[i], filesize[i], values[i]);
        }
    }

    function uploadMany(string[] memory merkle, uint64[] memory filesize, uint[] memory values, uint64 dayCount) public payable {
        for (uint i=0; i<merkle.length; i++) {
            upload(merkle[i], filesize[i], values[i], dayCount);
        }
    }

    function fileAddress(address _addr, uint256 _index) public view returns (string memory) {
        // takes address and index
        require(_index < cabinet[_addr].length); // check file exists for address
        return cabinet[_addr][_index].merkle; // return file
    }

    function fileStart(address _addr, uint256 _index) public view returns (uint) {
        // takes address and index
        require(_index < cabinet[_addr].length); // check file exists for address
        return cabinet[_addr][_index].start; // return file
    }

    function fileDetails(address _addr, uint256 _index) public view returns (FileDetails memory) {
        // takes address and index
        require(_index < cabinet[_addr].length); // check file exists for address
        return cabinet[_addr][_index]; // return file
    }

    function fileCount(address _addr) public view returns (uint256) {
        // takes address
        return cabinet[_addr].length; // return file count
    }
}
