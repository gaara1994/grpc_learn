package main

import (
	"context"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc_learn/5_tls/pb"
	"io/ioutil"
	"log"
)

func main() {
	rootCAs := x509.NewCertPool()
	pem, err := ioutil.ReadFile("../tls_certs/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	if !rootCAs.AppendCertsFromPEM(pem) {
		log.Fatal("Failed to append PEM.")
	}

	creds, err := credentials.NewClientTLSFromFile("../tls_certs/ca.crt", "localhost")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials: %v", err)
	}

	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(creds),
		// 其他可选配置...
	)

	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	ctx := context.Background()
	req := pb.HelloRequest{Name: "张三"}

	rep, err := client.SayHello(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rep.GetMessage())
}
