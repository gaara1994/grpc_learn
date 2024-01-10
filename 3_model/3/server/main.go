package main

import (
	"google.golang.org/grpc"
	"grpc_learn/3_model/3/pb"
	"io"
	"log"
	"net"
)

type CalculateServiceServer struct {
	pb.UnimplementedCalculateServiceServer
}

func (s *CalculateServiceServer) Sum(stream pb.CalculateService_SumServer) error {
	var sum int64

	for {
		//从stream中读取数据
		req, err := stream.Recv()
		if err == io.EOF {
			// 当没有更多消息时，返回最终结果
			return stream.SendAndClose(&pb.SumResponse{Number: sum})
		}
		if err != nil {
			log.Fatalf("Failed to receive SumRequest: %v", err)
		}
		// 对接收到的每个 number 进行累加
		sum += req.Number
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCalculateServiceServer(grpcServer, &CalculateServiceServer{})
	log.Printf("Server listening at %v", lis.Addr())

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
