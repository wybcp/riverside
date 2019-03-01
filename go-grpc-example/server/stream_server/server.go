package main

import (
	"io"
	"log"
	"net"
	pb "riverside/go-grpc-example/proto"

	"google.golang.org/grpc"
)

const PORT = "9002"

type StreamService struct {
}

func main() {
	server := grpc.NewServer()
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
	for   {
		r,err:=stream.Recv()
		if err==io.EOF {
			return stream.SendAndClose(&pb.StreamResponse{Pt:&pb.StreamPoint{Name:"gRPC Stream Server: Record",Value:1}})
		}
		if err!=nil{
			return err
		}
		log.Printf("Stream.Recv pt.Name:%s,pt.Value:%d",r.Pt.Name,r.Pt.Value)
	}
	return nil
}

func (s *StreamService) Route(stream pb.StreamService_RouteServer) error {
	n:=0
	for  {
		//发送
		err:=stream.Send(&pb.StreamResponse{
			Pt:&pb.StreamPoint{
				Name:"gRPC Stream Client:Route",
				Value:int32(n),
			},
		})
		if err != nil {
			return err
		}
	//	接收
		r,err:=stream.Recv()
		if err==io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		n++
		log.Printf("Stream.Recv pt.Name:%s,pt.Value:%d",r.Pt.Name,r.Pt.Value)

	}
	return nil
}
