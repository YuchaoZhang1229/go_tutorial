## gRPC核心概念
gRPC 是一个现代化、高性能、跨语言的 RPC 框架，它通过严格的接口定义和强大的代码生成能力，简化了分布式系统中服务间的通信

### 1. 什么是gRPC?
- **目标**：让一台计算机上的程序能够轻松地调用另一台计算机上的程序中的方法，就像调用本地方法一样
- **核心**：通过**定义服务接口**来实现。先定义好有什么方法，方法的参数和返回值是什么，然后客户端和服务器都遵守这个约定。

### 2. 如何定义服务？（使用 Protocol Buffers）
- **IDL**：使用一种叫做 Protocol Buffers 的语言来定义服务和消息的结构。它简单、高效，并且能生成多种编程语言的代码
- **示例**：
```protobuf
// 定义一个服务
service HelloService {
  // 定义一个方法。客户端发送 HelloRequest，服务器返回 HelloResponse。
  rpc SayHello (HelloRequest) returns (HelloResponse);
}

// 定义请求消息的结构
message HelloRequest {
  string greeting = 1; // 一个字符串类型的字段，数字 1 是字段编号
}

// 定义响应消息的结构
message HelloResponse {
  string reply = 1;
}
```