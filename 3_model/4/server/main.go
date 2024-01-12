package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc_learn/3_model/4/pb"
	"io"
	"log"
	"net"
)

type CalculateService struct {
	pb.UnimplementedCalculateServiceServer
}

func (s *CalculateService) BidirectionalSum(stream pb.CalculateService_BidirectionalSumServer) error {
	var num int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to receive SumRequest: %v", err)
		}

		num += req.Number

		// 向客户端发送当前的累加结果
		resp := &pb.SumResponse{Number: num}

		err = stream.Send(resp)
		if err != nil {
			return fmt.Errorf("failed to send SumResponse: %v", err)
		}
	}
	return status.Errorf(codes.Unimplemented, "method BidirectionalSum not implemented")
}

func main() {
	//
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCalculateServiceServer(grpcServer, &CalculateService{})

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}

}
