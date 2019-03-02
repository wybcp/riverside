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
type Auth struct {
	AppKey string
	AppSecret string
}

func (a *Auth)GetRequestMetadata(ctx context.Context,uri ...string)(map[string]string,error)  {
	return map[string]string{"app_key":a.AppKey,"app_secret":a.AppSecret},nil
}
func (a *Auth)RequireTransportSecurity()bool  {
	return true
}
func main() {
	err:= ginit.InitViper()
	if err != nil {
		log.Fatalf("init.InitViper err:", err)
	}
	clientTLS := gtls.ClientTLS{
		CertFile:   viper.GetString("tls.CERT_FILE"),
		ServerName: viper.GetString("tls.SERVER_NAME"),
	}
	c, err := clientTLS.GetTLSCredentials()
	if err != nil {
		log.Fatalf("clientTLS.GetTLSCredentials err:", err)
	}
	auth:=Auth{
		AppKey:viper.GetString("token.APP_KEY"),
		AppSecret:viper.GetString("token.APP_SECRET"),
	}
	conn, err := grpc.Dial(":"+viper.GetString("port.SIMPLE_TOKEN"),
		grpc.WithTransportCredentials(c),
		grpc.WithPerRPCCredentials(&auth))
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
