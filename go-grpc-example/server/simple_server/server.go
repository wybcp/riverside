package main

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"riverside/go-grpc-example/interceptor"
	"riverside/go-grpc-example/pkg/ginit"
	"riverside/go-grpc-example/pkg/gtls"
	pb "riverside/go-grpc-example/proto"
)

// SearchService 结构
type SearchService struct {
}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {

	if ctx.Err() == context.Canceled {
		log.Println("server receive cancel signal")
		return nil, status.Errorf(codes.Canceled, "SearchService Search canceled")
	}

	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

func main() {
	err := ginit.InitViper()
	if err != nil {
		log.Fatalf("init.InitViper err:", err)
	}
	tlsServer := gtls.ServerTLS{
		CertFile: viper.GetString("tls.CERT_FILE"),
		KeyFile:  viper.GetString("tls.KEY_FILE"),
	}
	//log.Fatal(viper.GetString("server.CERT_FILE"))

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
	lis, err := net.Listen("tcp", ":"+viper.GetString("port.SIMPLE"))
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	if err := server.Serve(lis); err != nil {
		log.Fatalf("server err: %v", err)
	}
}
