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
