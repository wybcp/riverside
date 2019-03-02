package main

import (
	"context"
	"log"
	"net"
	"riverside/go-grpc-example/interceptor"
	"riverside/go-grpc-example/pkg/gtls"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	pb "riverside/go-grpc-example/proto"
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
	certFle := "src/riverside/go-grpc-example/conf/server.pem"
	keyFile := "src/riverside/go-grpc-example/conf/server.key"
	tlsServer := gtls.ServerTLS{
		CertFile: certFle,
		KeyFile:  keyFile,
	}
	c, err := tlsServer.GetTLSCredentials()
	if err != nil {
		log.Fatalf("tlsServer.GetTLSCredentials err:", err)
	}

	opts := []grpc.ServerOption{
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
