package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc_learn/3_model/3/pb"
	"io"
	"log"
	"net"
	"time"
)

type VehicleTracking struct {
	pb.UnimplementedVehicleTrackingServiceServer
}

func (vt *VehicleTracking) RecordVehicleRoute(stream pb.VehicleTrackingService_RecordVehicleRouteServer) error {
	var totalDistance float64
	maxSpeed := 0.0
	startTime := time.Now()
	lastLocation := &pb.VehicleLocation{}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Errorf(codes.Internal, "failed to receive vehicle location: %v", err)
		}
		log.Println("收到数据：", req)

		// 计算总距离、最大速度等信息（这里仅为示例，实际情况可能更复杂）
		if lastLocation != nil {
			distance := calculateDistance(lastLocation, req)
			totalDistance += distance
			speed := calculateSpeed(lastLocation.Timestamp, req.Timestamp, distance)
			if speed > maxSpeed {
				maxSpeed = speed
			}
		}
		lastLocation = req

		// 可以在此处进行实时处理或存储到数据库
		// ...
	}

	endTime := time.Now()

	// 计算完成，发送 RouteSummary 给客户端
	summary := &pb.RouteSummary{
		VehicleId:     lastLocation.VehicleId,
		TotalDistance: int32(totalDistance),
		MaxSpeed:      int32(maxSpeed),
		StartTime:     startTime.UnixNano(),
		EndTime:       endTime.UnixNano(),
	}

	if err := stream.SendAndClose(summary); err != nil {
		return fmt.Errorf("failed to send route summary: %v", err)
	}

	return nil
}

// 示例辅助函数，用于计算两点之间的直线距离
func calculateDistance(a, b *pb.VehicleLocation) float64 {
	// 使用 Haversine 公式或其他算法计算经纬度间的距离
	// 此处仅作示例，具体实现需参考地理坐标距离计算方法
	return 0.0
}

// 示例辅助函数，用于计算速度
func calculateSpeed(startTime, endTime int64, distance float64) float64 {
	timeDelta := time.Duration(endTime - startTime).Seconds()
	return distance / timeDelta
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterVehicleTrackingServiceServer(s, &VehicleTracking{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
