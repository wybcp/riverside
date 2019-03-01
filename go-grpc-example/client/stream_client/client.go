package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"io/ioutil"
	"log"
	pb "riverside/go-grpc-example/proto"
)

const PORT = "9002"

func main() {
	cert,err:=tls.LoadX509KeyPair("src/riverside/go-grpc-example/conf/server.pem","src/riverside/go-grpc-example/conf/server.key")
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err:%v",err)
	}
	certPool:=x509.NewCertPool()
	ca,err:=ioutil.ReadFile("src/riverside/go-grpc-example/conf/ca.pem")
	if err != nil {
		log.Fatalf("ioutil.ReadFile err:%v",err)
	}
	if ok:=certPool.AppendCertsFromPEM(ca);!ok{
		log.Fatal("certPool.AppendCertsFromPEM err")
	}

	c:=credentials.NewTLS(&tls.Config{
		Certificates:[]tls.Certificate{cert},
		ServerName:"wyb",
		RootCAs:certPool,
	})


	conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("grpc dial err:%v", err)
	}
	defer conn.Close()

	client := pb.NewStreamServiceClient(conn)
	err = printLists(client, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRPC stream client :List", Value: 2019}})
	if err != nil {
		log.Fatalf("printLists err :%v", err)
	}
	err = printRecord(client, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRPC stream client :Record", Value: 2019}})
	if err != nil {
		log.Fatalf("printRecord err :%v", err)
	}
	err = printRoute(client, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRPC stream client :Route", Value: 2019}})
	if err != nil {
		log.Fatalf("printRoute err :%v", err)
	}
}

func printLists(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream, err := client.List(context.Background(), r)
	if err != nil {
		return err
	}
	for {
		resp, err := stream.Recv()
		//数据流结束标记
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("Resp: pt.name: %s,pt.value:%d", resp.Pt.Name, resp.Pt.Value)
	}
	return nil
}

func printRecord(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream, err := client.Record(context.Background())
	if err != nil {
		return err
	}
	for n := 0; n < 7; n++ {
		err = stream.Send(r)
		if err != nil {
			return err
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	log.Printf("Resp: pt.name: %s,pt.value:%d", resp.Pt.Name, resp.Pt.Value)
	return nil
}

func printRoute(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream,err:=client.Route(context.Background())
	if err != nil {
		return err
	}
	for n := 0; n < 7; n++ {
		err = stream.Send(r)
		if err != nil {
			return err
		}
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("Resp: pt.name: %s,pt.value:%d", resp.Pt.Name, resp.Pt.Value)
	}
	stream.CloseSend()

	return nil
}
