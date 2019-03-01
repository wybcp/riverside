package main

import (
	"context"
	"google.golang.org/grpc/credentials"
	"log"
	"net"

	pb "riverside/go-grpc-example/proto"

	"google.golang.org/grpc"
)

// SearchService 结构
type SearchService struct {
}

// PORT 服务端口
const PORT = "9001"

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}


func main() {
	c, err := credentials.NewServerTLSFromFile("src/riverside/go-grpc-example/conf/server.pem", "src/riverside/go-grpc-example/conf/server.key")
	if err != nil {
		log.Fatalf("credentials.NewClientTLSFromFile err:", err)
	}
	server := grpc.NewServer(grpc.Creds(c))
	pb.RegisterSearchServiceServer(server, &SearchService{})
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	if err := server.Serve(lis); err != nil {
		log.Fatalf("server err: %v", err)
	}
}
