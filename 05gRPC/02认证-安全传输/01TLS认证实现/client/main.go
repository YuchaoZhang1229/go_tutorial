package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"

	pb "github.com/go-tutorial/05gRPC/01HelloWorld/proto"
)

func main() {
	// 加载 CA 证书 (验证服务器证书) 路径需要修改
	creds, err := credentials.NewClientTLSFromFile("E:\\OneDrive\\web3\\go\\go-tutorial\\05gRPC\\certs\\test.pem", "example.com")
	if err != nil {
		log.Fatalf("Failed to load  TLS credentials %v", err)
	}
	// 连接到server端，此处禁用安全传输，没有加密和验证
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(creds))
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
