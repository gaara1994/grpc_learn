package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_learn/4_alts/pb"
)

func main() {
	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
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
