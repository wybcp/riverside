package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"log"
	"net"
	"riverside/go-grpc-example/pkg/ginit"
	"riverside/go-grpc-example/pkg/gtls"
	pb "riverside/go-grpc-example/proto"

	zipkin "github.com/openzipkin/zipkin-go-opentracing"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/spf13/viper"

	"google.golang.org/grpc"
)

type SearchService struct {
}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

func main() {
	err := ginit.InitViper()
	if err != nil {
		log.Fatalf("init.InitViper err:", err)
	}
	collector, err := zipkin.NewHTTPCollector(viper.GetString("port.SIMPLE_ZIPKIN_PORT"))
	if err != nil {
		log.Fatalf("zipkin.NewHTTPCollector err:%v", err)
	}
	recorder:=zipkin.NewRecorder(collector,true,viper.GetString("port.ZIPKIN_RECORD_HOST_PORT"),viper.GetString("tls.SERVER_NAME"))

	tracer,err:=zipkin.NewTracer(
		recorder,zipkin.ClientServerSameSpan(false),
	)
	if err != nil {
		log.Fatalf("zipkin.NewTracer err:%v", err)
	}
	
	serverTLS := gtls.ServerTLS{
		CertFile: viper.GetString("tls.CERT_FILE"),
		KeyFile:  viper.GetString("tls.KEY_FILE"),
		CaFile:   viper.GetString("tls.CA_FILE"),
	}
	c, err := serverTLS.GetTLSCredentialsByCA()
	if err != nil {
		log.Fatalf("serverTLS.GetTLSCredentialsByCA err:%v", err)
	}
	opts := []grpc.ServerOption{
		grpc.Creds(c),
		grpc_middleware.WithUnaryServerChain(
			otgrpc.OpenTracingServerInterceptor(tracer,otgrpc.LogPayloads()),
		),
	}

	server := grpc.NewServer(opts...)
	pb.RegisterSearchServiceServer(server, &SearchService{})
	lis, err := net.Listen("tcp", ":"+viper.GetString("port.SIMPLE_ZIPKIN_PORT"))
	if err != nil {
		log.Fatalf("net listen err:%v", err)
	}
	server.Serve(lis)
}
