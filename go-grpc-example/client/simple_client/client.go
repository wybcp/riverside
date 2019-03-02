package main

import (
	"context"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"riverside/go-grpc-example/pkg/ginit"
	"riverside/go-grpc-example/pkg/gtls"
	pb "riverside/go-grpc-example/proto"
	"time"
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

	//deadline
	d:=time.Now().Add(time.Duration(20*time.Second))
	ctx,cancel:=context.WithDeadline(context.Background(),d)
	defer func() {
		log.Println("client sent deadline cancel signal")
		cancel()
	}()

	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(ctx, &pb.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		statusErr,ok:=status.FromError(err)
		if ok {
			if statusErr.Code()==codes.DeadlineExceeded {
				log.Fatalln("client.Search err:deadline")
			}
		}
		log.Fatalf("client.Search err:%v", err)
	}
	log.Printf("response :%s", resp.GetResponse())
}
