package main

import (
	"context"
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc_learn/5_tls/pb"
	"log"
	"net"
)

type GreeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s GreeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("收到消息：", in.GetName())
	return &pb.HelloReply{Message: "服务端返回的消息"}, nil
}
func main() {
	cert, err := tls.LoadX509KeyPair("../tls_certs/server.crt", "../tls_certs/server.key")
	if err != nil {
		log.Fatal(err)
	}

	creds := credentials.NewServerTLSFromCert(&cert)

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterGreeterServer(grpcServer, &GreeterServer{})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
