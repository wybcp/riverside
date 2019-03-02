package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"riverside/go-grpc-example/pkg/gtls"
	pb "riverside/go-grpc-example/proto"
)

const PORT = "9003"

func main() {
	clientTLS:=gtls.ClientTLS{
		CertFile:"src/riverside/go-grpc-example/conf/server.pem",
		ServerName:"wyb",
	}
	c, err := clientTLS.GetTLSCredentials()
	if err != nil {
		log.Fatalf("clientTLS.GetTLSCredentials err:", err)
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
