// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Store {
    event ItemSet(bytes32 key, bytes32 value);

    string public version;
    mapping (bytes32 => bytes32) private _items;

    constructor(string memory _version) {
        version = _version;
    }

    function setItem(bytes32 key, bytes32 value) external {
        _items[key] = value;
        emit ItemSet(key, value);
    }

    function getItem(bytes32 key) external view returns (bytes32) {
        return _items[key];
    }
}
