pragma solidity ^0.8.26;

// 当你调用 setItem函数并传入一个键（key）和值（value）时，合约会执行以下操作：
// 1. 更新状态：将传入的 value存储到 items映射中与传入的 key对应的位置，即 items[key] = value。
// 2. 发出事件：触发 ItemSet(key, value)事件。这个事件会被记录在区块链的日志中，任何监听此事件的应用都能接收到这次操作的通知
// 由于 items映射被声明为 public，Solidity 编译器会自动为它生成一个同名的getter函数。这意味着你不需要编写额外的函数，就可以直接通过传入 key来查询 items映射中该键对应的值

contract Store {
  event ItemSet(bytes32 indexed key, bytes32 value);

  string public version;
  mapping (bytes32 => bytes32) public items;

  constructor(string memory _version) {
    version = _version;
  }

  function setItem(bytes32 key, bytes32 value) external {
    items[key] = value;
    emit ItemSet(key, value);
  }
}