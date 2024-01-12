package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_learn/3_model/4/pb"
	"log"
	"time"
)

func main() {
	conn, err := grpc.DialContext(context.Background(), "localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {

	}
	client := pb.NewCalculateServiceClient(conn)

	ctx := context.Background()
	stream, err := client.BidirectionalSum(ctx)
	if err != nil {
		log.Fatalf("Error client.BidirectionalSum(ctx): %v", err)
	}

	for i := 0; i < 10; i++ {
		req := &pb.SumRequest{Number: int64(i)}
		err = stream.Send(req)
		if err != nil {
			log.Fatalf("Error stream.Send(req): %v", err)
		}

		// 从服务端接收中间结果
		resp, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error receiving SumResponse: %v", err)
		}

		log.Printf("累加后的结果: %d", resp.GetNumber())
		time.Sleep(time.Second)
	}

	// 完成发送后关闭流
	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("Error closing send side of the stream: %v", err)
	}

}
