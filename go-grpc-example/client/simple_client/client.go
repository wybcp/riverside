package main

import (
	"context"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"riverside/go-grpc-example/pkg/ginit"
	"riverside/go-grpc-example/pkg/gtls"
	pb "riverside/go-grpc-example/proto"
)

func main() {
	err:= ginit.InitViper()
	if err != nil {
		log.Fatalf("init.InitViper err:", err)
	}
	clientTLS:=gtls.ClientTLS{
		CertFile:viper.GetString("tls.CERT_FILE"),
		ServerName:viper.GetString("tls.SERVER_NAME"),
	}
	c, err := clientTLS.GetTLSCredentials()
	if err != nil {
		log.Fatalf("clientTLS.GetTLSCredentials err:", err)
	}

	conn, err := grpc.Dial(":"+viper.GetString("port.SIMPLE"), grpc.WithTransportCredentials(c))
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
