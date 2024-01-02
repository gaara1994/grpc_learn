package main

import (
	"google.golang.org/grpc"
	"grpc_learn/2_grpc/product"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051") // 端口号自定义
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// 创建并配置你的 ProductServiceServer 实例
	productServiceServer := &product.ProductServiceServer{}

	// 注册 ProductService 服务到 gRPC 服务器
	product.RegisterProductServiceServer(grpcServer, productServiceServer)

	log.Printf("Starting gRPC server on port :%d", 50051)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
