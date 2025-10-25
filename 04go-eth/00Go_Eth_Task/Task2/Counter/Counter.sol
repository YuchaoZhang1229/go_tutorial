// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Counter {
    // 状态变量：永久存储在区块链上
    uint256 private count;  // 使用uint256存储计数，private限制仅合约内可访问

    // 事件：记录重要的状态变更
    event CounterIncremented(address indexed sender, uint256 newValue);

    // 构造函数，在合约部署时初始化计数器为0
    constructor() {
        count = 0;
    }

   // 增加计数函数，公开可调用，修改状态需消耗Gas
    function increment() public {
        count += 1;
        emit CounterIncremented(msg.sender, count);  // 触发事件记录调用者和新值
    }

    // 获取当前计数（view函数，只读不消耗Gas）
    function getCount() public view returns (uint256) {
        return count;
    }
}