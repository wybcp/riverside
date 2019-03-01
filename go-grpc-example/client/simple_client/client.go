package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	pb "riverside/go-grpc-example/proto"
)

const PORT = "9001"

func main() {
	c, err := credentials.NewClientTLSFromFile("src/riverside/go-grpc-example/conf/server.pem", "wyb")
	if err != nil {
		log.Fatalf("credentials.NewClientTLSFromFile err:", err)
	}

	conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("grpc.Dial err:%v", err)
	}
	defer conn.Close()

	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &pb.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err:%v", err)
	}
	log.Printf("response :%s", resp.GetResponse())
}
