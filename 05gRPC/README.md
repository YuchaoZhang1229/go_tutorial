### 一、gRPC介绍
**单体架构**：所有功能打包为一个整体，单点故障可能就引起整体系统崩溃

**微服务架构**：每个服务独立部署，互不影响，故障被隔离在单个服务内，系统整体可用性更好


| 特性         | 单体架构 (Monolithic Architecture)       | 微服务架构 (Microservices Architecture) |
|------------| ------------------------------------------------------------ |------------------------------------|
| **部署方式**       | 所有功能打包为一个整体，统一部署 | 每个服务独立部署，互不影响                      |
| **技术栈**        | 通常采用统一的技术栈 | 不同服务可采用最不同的技术（多语言、多框架）             |
| **扩展性** | 整体扩展，可能造成资源浪费 | 可按需单独扩展高负载的服务，资源利用率高               |
| **故障影响** | 单点故障可能导致整个系统崩溃 | 故障被隔离在单个服务内，系统整体可用性更高              |
| **开发与维护**      | 简单，适合小型项目；但项目庞大后复杂度高，维护成本高 | 分布式系统复杂，对开发、测试、运维和监控要求高 |



**RPC的全称是 Remote Procedure Call**，远程过程调用，这是一种协议，使得你可以像本地调用一样直接调用一个远程的函数
client 与 server 沟通的过程
1. client 发送数据（以**字节流**的方式）
2. 服务端接受并解析。根据约定直到要执行什么，然后把结果响应给客户端
   RPC就是将上述过程封装，规范化

gRPC 采用 “**服务定义**” 来约定远 程服务。先通过一种与开发语言无关的方式（如 Protobuf）明确描述**服务名**、**可用方法及各方法的参数与返回值**。此后，gRPC 会**屏蔽底层通信细节**：客户端只需要直接调用定义好的接口，而服务端则只需专注实现这些方法的具体业务逻辑



### 二、Protocol Buffer介绍
**Protocol Buffers**（简称 Protobuf）是 Google 开发的一种**语言中立**、**平台中立**的序列化结构化数据的机制。它通过 .proto文件定义数据结构，并利用编译器生成多种编程语言的代码，实现高效的数据序列化与反序列化
- **代码生成工具**和**序列化工具**
- **优势**
  - 序列化体积相比Json和XML都小，适合网络传输
  - 支持跨平台多语言
  - 消息格式升级和兼容性还不错
  - 序列化反序列化很快
- **工作原理**
  1. 定义数据结构：在 .proto文件中编写消息格式 
  2. 编译生成代码：使用 protoc编译器将 .proto文件转换为目标语言的数据访问类 
  3. 序列化/反序列化：调用生成类的接口将对象序列化为二进制数据，或反向解析

### 三、安装插件
1. Protobuf
```shell
winget install protobuf
protoc --version # Ensure compiler version is 3+
```

2. 安装gRPC核心库
```shell
go get google.golang.org/grpc
```
   
3. 安装protocol编译器
```shell
# 安装protobuf生成插件
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 确保GOPATH/bin在PATH中，以便使用安装的插件
export PATH="$PATH:$(go env GOPATH)/bin"
```

### 四、安装插件

#### 1. Proto文件编写
```protobuf
// 这是在说明我们使用的是proto3语法。
syntax ="proto3";

// 这部分的内容是关于最后生成的go文件是处在哪个目录哪个包中，.代表在当前目录生成，service代表了生成的go文件的包名是service。
option go_package = ".;service";

// 然后我们需要定义一个服务，在这个服务中需要有一个方法，这个方法可以接受客户端的参数，再返回服务端的响应。
// 其实很容易可以看出，我们定义了一个 service，称为 SayHeLLo，这个服务中有一个rpc方法，名为SayHeLLo。
// 这个方法会发送一个 HelloRequest，然后返回一个 HelloResponse。
service SayHello {
rpc SayHello(HelloRequest) returns (HelloResponse) {}
}
// message 关键字，其实你可以理解为 Golang 中的结构体。
// 这里比较特别的是变量后面的"赋值"。 注意，这里并不是赋值，而是在定义这个变量在这个 message 中的位置。
message HelloRequest {
  string requestName = 1;
  // int64 age = 2;
}
message HelloResponse {
  string responseMsg = 1;
}
```

```shell
protoc --go_out=. hello.proto        // 生成 hello.pb.go
protoc --go-grpc_out=. hello.proto   // 生成 hello_grpc.pb.go
```

#### 2. 服务端代码编写
- 创建 gRPC Server 对象, 你可以理解为他是 Server 端的抽象对象
- server（其包含需要被调用的服务端接口）注册到 gRPC Server 的内部注册中心。这样可以在接受到请求时，通过内部的服务发现，发现该服务端接口并转接进行逻辑处理
- 创建 Listen, 监听 TCP 端口
- gRPC Server 开始 lis.Accept, 直到 Stop
```go
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"

	pb "github.com/go-tutorial/05gRPC/01HelloWorld/server/proto"
)

// hello server
type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Printf("server: SayHello")
	return &pb.HelloResponse{ResponseMsg: "hello" + req.RequestName}, nil
}

func main() {
	// 开启端口
	listen, _ := net.Listen("tcp", ":8080")
	// 创建grpc服务
	grpcServer := grpc.NewServer()
	// 在grpc服务端中注册我们自己编写的服务
	pb.RegisterSayHelloServer(grpcServer, &server{})
	// 启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}

}

```

#### 3. 客户端代码编写
- 创建与给定目标（服务端）的连接交互
- 创建server的客户端对象
- 发送RPC请求，等待同步响应，得到回调后返回响应结果
- 输出响应结果

```go
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	pb "github.com/go-tutorial/05gRPC/01HelloWorld/client/proto"
)

func main() {
	// 连接到server端，此处禁用安全传输，没有加密和验证
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 建立连接
	client := pb.NewSayHelloClient(conn)

	// 执行rpc调用 (这个方法在服务器端来实现并返回结果)
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "rpc"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	fmt.Println(resp.GetResponseMsg())

}

```