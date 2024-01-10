package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_learn/3_model/3/pb"
	"log"
)

func main() {
	//1.创建与服务器连接
	conn, err := grpc.DialContext(
		context.Background(),
		":50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	defer conn.Close()

	//2.创建客户端
	client := pb.NewCalculateServiceClient(conn)

	// 创建一个 context 对象，用于控制 RPC 调用的行为
	ctx := context.Background()

	//创建一个客户端流
	stream, err := client.Sum(ctx)

	//stream发送数据
	err = stream.Send(&pb.SumRequest{Number: 10})
	err = stream.Send(&pb.SumRequest{Number: 90})
	err = stream.Send(&pb.SumRequest{Number: -150})
	err = stream.Send(&pb.SumRequest{Number: 0})
	if err != nil {
		log.Fatalf("Error sending SumRequest: %v", err)
	}

	// 完成请求后关闭发送端并接收服务端的唯一响应
	sumResponse, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving SumResponse: %v", err)
	}
	log.Printf("计算结果: %d", sumResponse.GetNumber())
}
