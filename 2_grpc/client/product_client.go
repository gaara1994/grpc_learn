package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc_learn/2_grpc/pb"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("grpc.Dial err:", err.Error())
	}
	defer conn.Close()
	c := pb.NewProductServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//r, err := c.GetProduct(ctx, &pb.GetProductRequest{Id: 1})
	//if err != nil {
	//	log.Fatal("c.ListProduct err:", err.Error())
	//}
	//log.Println("获取响应", r)

	r2, err := c.ListProducts(ctx, &pb.ListProductsRequest{Name: "手机"})
	if err != nil {
		log.Fatal("c.ListProduct err:", err.Error())
	}
	log.Println("获取响应2", r2)
	for _, product := range r2.Products {
		log.Println(product)
	}

}
