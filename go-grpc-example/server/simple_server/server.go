package main

import (
	"context"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"riverside/go-grpc-example/interceptor"

	pb "riverside/go-grpc-example/proto"
	"github.com/grpc-ecosystem/go-grpc-middleware"
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
	opts:=[]grpc.ServerOption{
		grpc.Creds(c),
		grpc_middleware.WithUnaryServerChain(
			interceptor.LoggingInterceptor,
			interceptor.RecoveryInterceptor,
			),
	}
	server := grpc.NewServer(opts...)
	pb.RegisterSearchServiceServer(server, &SearchService{})
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	if err := server.Serve(lis); err != nil {
		log.Fatalf("server err: %v", err)
	}
}
