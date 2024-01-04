package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc_learn/3_model/2/pb"
	"log"
	"net"
	"time"
)

type MusicServiceServer struct {
	pb.UnimplementedMusicServiceServer
}

type songDatabase struct {
	songs []pb.SongResponse
}

func (s *MusicServiceServer) ListSongs(in *pb.ArtistRequest, stream pb.MusicService_ListSongsServer) error {
	//假设这是从数据库获取的所有歌曲（此处仅为演示，实际应从数据库查询）
	db := &songDatabase{songs: []pb.SongResponse{
		{SongName: "第一交响曲（C大调 Op.21）", Album: "贝多芬专辑", ReleaseYear: 1796},
		{SongName: "第二交响曲（D大调 Op.36）", Album: "贝多芬专辑", ReleaseYear: 1802},
		{SongName: "第三交响曲（降E大调 Op.55 “英雄”）", Album: "贝多芬专辑", ReleaseYear: 1804},
		{SongName: "第四交响曲（降B大调 Op.60）", Album: "贝多芬专辑", ReleaseYear: 1806},
		{SongName: "第五交响曲（c小调 Op.67 “命运”）", Album: "贝多芬专辑", ReleaseYear: 1808},
		{SongName: "第六交响曲（F大调 Op.68 “田园”）", Album: "贝多芬专辑", ReleaseYear: 1808},
		{SongName: "第七交响曲（A大调 Op.92）", Album: "贝多芬专辑", ReleaseYear: 1811},
		{SongName: "第八交响曲（F大调 Op.93）", Album: "贝多芬专辑", ReleaseYear: 1812},
		{SongName: "第九交响曲（d小调 Op.125 “合唱”）", Album: "贝多芬专辑", ReleaseYear: 1824},
		{SongName: "第十交响曲（未完成）", Album: "贝多芬专辑", ReleaseYear: 0},
		{SongName: "威灵顿的胜利（战争交响曲）（D大调 Op,91）", Album: "贝多芬专辑", ReleaseYear: 1813},
	}}

	//使用流发送
	for _, song := range db.songs {
		time.Sleep(time.Second)
		if err := stream.Send(&song); err != nil {
			return fmt.Errorf("failed to send song: %v", err)
		}
	}
	return nil
}
func main() {
	//1.创建一个监听在 TCP 端口 50051 上的网络监听器（Listener）。当客户端连接到这个端口时，服务器将能够接受和处理请求。
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//2.创建一个新的 gRPC 服务器实例。这是一个实现了 gRPC 协议栈的核心服务器对象，用于注册服务、接收客户端连接以及处理 RPC 请求。
	s := grpc.NewServer()

	//3.注册你在 .proto 文件中定义的服务（这里是 MusicService）的实现到 gRPC 服务器。
	pb.RegisterMusicServiceServer(s, &MusicServiceServer{})

	//4.开始在之前创建的监听器上运行 gRPC 服务器。Serve() 方法会让服务器开始监听传入的连接，并处理来自客户端的 RPC 请求。
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
