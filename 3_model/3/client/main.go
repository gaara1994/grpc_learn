package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"grpc_learn/3_model/3/pb" // 替换为实际protobuf生成包路径
)

type vehicleLocationGenerator struct {
	current int // 用于跟踪已生成的位置数量
}

func (vlg *vehicleLocationGenerator) generate() *pb.VehicleLocation {
	if vlg.current >= 5 { // 当已生成5个位置时返回nil表示没有更多数据
		return nil
	}

	vlg.current++ // 每次生成一个位置后递增计数器

	// 示例中的经纬度可以随机生成或按照实际轨迹设定
	latitude := 37.0 + float64(vlg.current)/10.0
	longitude := -122.0 - float64(vlg.current)/10.0

	location := &pb.VehicleLocation{
		VehicleId: "example_vehicle_id",
		Latitude:  latitude,
		Longitude: longitude,
		Timestamp: time.Now().UnixNano(),
	}

	return location
}

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewVehicleTrackingServiceClient(conn)

	// 创建一个模拟的位置生成器，生成5个位置信息
	vlg := &vehicleLocationGenerator{}

	stream, err := client.RecordVehicleRoute(context.Background())
	if err != nil {
		log.Fatalf("Failed to create RecordVehicleRoute stream: %v", err)
	}

	if err := vlg.sendVehicleRoute(stream); err != nil {
		log.Fatalf("Failed to send vehicle route: %v", err)
	}

	// 当所有位置信息发送完毕后，等待服务器返回的 RouteSummary
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to receive RouteSummary: %v", err)
	}

	fmt.Printf("Received Route Summary: %+v\n", resp)
}

func (vlg *vehicleLocationGenerator) sendVehicleRoute(stream pb.VehicleTrackingService_RecordVehicleRouteClient) error {
	for {
		location := vlg.generate()
		if location == nil { // 没有更多位置数据时退出循环
			break
		}

		if err := stream.Send(location); err != nil {
			return fmt.Errorf("failed to send vehicle location: %v", err)
		}

		// 这里不需要延时，因为我们在generate函数中已经限制了位置点的数量
		time.Sleep(time.Second) // 示例中的延时一秒
	}

	return nil
}
