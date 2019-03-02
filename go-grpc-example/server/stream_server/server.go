package main

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"io"
	"log"
	"net"
	"riverside/go-grpc-example/interceptor"
	"riverside/go-grpc-example/pkg/gtls"
	pb "riverside/go-grpc-example/proto"

	"google.golang.org/grpc"
)

const PORT = "9002"

type StreamService struct {
}

func main() {
	serverTLS := gtls.ServerTLS{
		CertFile: "src/riverside/go-grpc-example/conf/server.pem",
		KeyFile:  "src/riverside/go-grpc-example/conf/server.key",
		CaFile:   "src/riverside/go-grpc-example/conf/ca.pem",
	}
	c, err := serverTLS.GetTLSCredentialsByCA()
	if err != nil {
		log.Fatalf("serverTLS.GetTLSCredentialsByCA err:%v", err)
	}
	opts := []grpc.ServerOption{
		grpc.Creds(c),
		grpc_middleware.WithStreamServerChain(
			interceptor.LoggingStreamInterceptor,
			interceptor.RecoveryStreamInterceptor,
		),
	}

	server := grpc.NewServer(opts...)
	pb.RegisterStreamServiceServer(server, &StreamService{})
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net listen err:%v", err)
	}
	server.Serve(lis)
}

func (s *StreamService) List(r *pb.StreamRequest, stream pb.StreamService_ListServer) error {
	for n := 0; n < 7; n++ {
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  r.Pt.Name,
				Value: r.Pt.Value + int32(n),
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *StreamService) Record(stream pb.StreamService_RecordServer) error {
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.StreamResponse{Pt: &pb.StreamPoint{Name: "gRPC Stream Server: Record", Value: 1}})
		}
		if err != nil {
			return err
		}
		log.Printf("Stream.Recv pt.Name:%s,pt.Value:%d", r.Pt.Name, r.Pt.Value)
	}
	return nil
}

func (s *StreamService) Route(stream pb.StreamService_RouteServer) error {
	n := 0
	for {
		//发送
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  "gRPC Stream Client:Route",
				Value: int32(n),
			},
		})
		if err != nil {
			return err
		}
		//	接收
		r, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		n++
		log.Printf("Stream.Recv pt.Name:%s,pt.Value:%d", r.Pt.Name, r.Pt.Value)

	}
	return nil
}
