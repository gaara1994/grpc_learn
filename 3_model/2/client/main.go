package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_learn/3_model/2/pb"
	"io"
	"log"
)

func main() {
	// 1.连接到 gRPC 服务器。本地写测试可以使用不安全的连接
	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal("grpc.DialContext err:", err)

	}
	defer conn.Close()

	// 创建一个 MusicService 的客户端实例
	client := pb.NewMusicServiceClient(conn)

	// 准备请求参数
	req := &pb.ArtistRequest{ArtistName: "贝多芬"}

	// 创建一个上下文
	ctx := context.Background()

	// 调用 ListSongs 流式 RPC 方法并接收响应
	stream, err := client.ListSongs(ctx, req)
	if err != nil {
		log.Fatal("client.ListSongs err", err)
	}
	// 处理流式响应
	for {
		songResponse, err := stream.Recv()
		if err == io.EOF { // 如果没有更多数据可读，则退出循环
			break
		}
		if err != nil { // 其他错误则打印错误信息并退出
			log.Fatalf("接收失败: %v", err)
		}
		fmt.Printf("收到一条歌曲: %s 来自专辑: %s, 发布于: %d\n",
			songResponse.SongName, songResponse.Album, songResponse.ReleaseYear)
	}
	fmt.Println("接收完毕！")
}
