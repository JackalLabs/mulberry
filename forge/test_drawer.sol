// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

// forge/src/JackalInterface.sol

interface JackalInterface {
    function postFileFrom(address from, string memory merkle, uint64 filesize, string memory note, uint64 expires)
        external
        payable;
    function postFile(string memory merkle, uint64 filesize, string memory note, uint64 expires) external payable;
    function buyStorageFrom(
        address from,
        string memory for_address,
        uint64 duration_days,
        uint64 size_bytes,
        string memory referral
    ) external payable;
    function buyStorage(string memory for_address, uint64 duration_days, uint64 size_bytes, string memory referral)
        external
        payable;
    function deleteFileFrom(address from, string memory merkle, uint64 start) external;
    function deleteFile(string memory merkle, uint64 start) external;
    function requestReportFormFrom(
        address from,
        string memory prover,
        string memory merkle,
        string memory owner,
        uint64 start
    ) external;
    function requestReportForm(string memory prover, string memory merkle, string memory owner, uint64 start)
        external;
    function postKey(string memory key) external;
    function postKeyFrom(address from, string memory key) external;
    function deleteFileTree(string memory hash_path, string memory account) external;
    function deleteFileTreeFrom(address from, string memory hash_path, string memory account) external;
    function provisionFileTree(string memory editors, string memory viewers, string memory tracking_number) external;
    function provisionFileTreeFrom(
        address from,
        string memory editors,
        string memory viewers,
        string memory tracking_number
    ) external;
    function postFileTree(
        string memory account,
        string memory hash_parent,
        string memory hash_child,
        string memory contents,
        string memory viewers,
        string memory editors,
        string memory tracking_number
    ) external;
    function postFileTreeFrom(
        address from,
        string memory account,
        string memory hash_parent,
        string memory hash_child,
        string memory contents,
        string memory viewers,
        string memory editors,
        string memory tracking_number
    ) external;
    function addViewers(
        string memory viewer_ids,
        string memory viewer_keys,
        string memory for_address,
        string memory file_owner
    ) external;
    function addViewersFrom(
        address from,
        string memory viewer_ids,
        string memory viewer_keys,
        string memory for_address,
        string memory file_owner
    ) external;
    function removeViewers(string memory viewer_ids, string memory for_address, string memory file_owner) external;
    function removeViewersFrom(
        address from,
        string memory viewer_ids,
        string memory for_address,
        string memory file_owner
    ) external;
    function resetViewers(string memory for_address, string memory file_owner) external;
    function resetViewersFrom(address from, string memory for_address, string memory file_owner) external;
    function changeOwner(string memory for_address, string memory file_owner, string memory new_owner) external;
    function changeOwnerFrom(address from, string memory for_address, string memory file_owner, string memory new_owner)
        external;
    function addEditors(
        string memory editor_ids,
        string memory editor_keys,
        string memory for_address,
        string memory file_owner
    ) external;
    function addEditorsFrom(
        address from,
        string memory editor_ids,
        string memory editor_keys,
        string memory for_address,
        string memory file_owner
    ) external;
    function removeEditors(string memory editor_ids, string memory for_address, string memory file_owner) external;
    function removeEditorsFrom(
        address from,
        string memory editor_ids,
        string memory for_address,
        string memory file_owner
    ) external;
    function resetEditors(string memory for_address, string memory file_owner) external;
    function resetEditorsFrom(address from, string memory for_address, string memory file_owner) external;
    function createNotification(string memory to, string memory contents, string memory private_contents) external;
    function createNotificationFrom(
        address from,
        string memory to,
        string memory contents,
        string memory private_contents
    ) external;
    function deleteNotification(string memory notification_from, uint64 time) external;
    function deleteNotificationFrom(address from, string memory notification_from, uint64 time) external;
    function blockSenders(string[] memory to_block) external;
    function blockSendersFrom(address from, string[] memory to_block) external;
}

// forge/src/StorageDrawer.sol

 // solidity ^0.8.0 recommended

 // import bridge interface

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
        require(dayCount > 31, "day count must be more than 31");
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

