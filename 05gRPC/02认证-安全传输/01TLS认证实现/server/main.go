package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"

	pb "github.com/go-tutorial/05gRPC/01HelloWorld/proto"
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
	// 加载 TLS 证书 路径需要修改
	creds, err := credentials.NewServerTLSFromFile(
		"E:\\OneDrive\\web3\\go\\go-tutorial\\05gRPC\\certs\\test.pem",
		"E:\\OneDrive\\web3\\go\\go-tutorial\\05gRPC\\certs\\test.key")
	if err != nil {
		log.Fatalf("Failed to load TLS credentials: %v", err)
	}
	// 创建grpc服务
	grpcServer := grpc.NewServer(grpc.Creds(creds))

	// 开启端口
	listen, _ := net.Listen("tcp", ":8080")

	// 在grpc服务端中注册我们自己编写的服务
	pb.RegisterSayHelloServer(grpcServer, &server{})
	// 启动服务
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}

}
