package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/alts"
	"grpc_learn/4_alts/pb"
	"log"
	"net"
)

type GreeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s GreeterServer) SayHello(context.Context, *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "服务端返回的消息"}, nil
}
func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	//gRPC 服务器可以使用 ALTS 凭据来允许客户端连接到它们
	altsTC := alts.NewServerCreds(alts.DefaultServerOptions())
	grpcServer := grpc.NewServer(grpc.Creds(altsTC))
	pb.RegisterGreeterServer(grpcServer, &GreeterServer{})

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}

}
