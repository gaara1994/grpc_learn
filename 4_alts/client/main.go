package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/alts"
	"grpc_learn/4_alts/pb"
)

func main() {
	altsTC := alts.NewClientCreds(alts.DefaultClientOptions())
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(altsTC))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//2.客户端实例
	client := pb.NewGreeterClient(conn)

	ctx := context.Background()

	req := pb.HelloRequest{Name: "张三"}

	rep, err := client.SayHello(ctx, &req)
	if err != nil {
		panic(err)
	}

	fmt.Println(rep.GetMessage())

}
