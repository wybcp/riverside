package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
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
	collector, err := zipkin.NewHTTPCollector(viper.GetString("port.SIMPLE_ZIPKIN_PORT"))
	if err != nil {
		log.Fatalf("zipkin.NewHTTPCollector err:%v", err)
	}
	recorder:=zipkin.NewRecorder(collector,true,viper.GetString("port.ZIPKIN_RECORD_HOST_PORT"),viper.GetString("tls.SERVER_NAME"))
	tracer,err:=zipkin.NewTracer(
		recorder,zipkin.ClientServerSameSpan(true),
	)
	if err != nil {
		log.Fatalf("zipkin.NewTracer err:%v", err)
	}


	clientTLS:=gtls.ClientTLS{
		CertFile:viper.GetString("tls.CERT_FILE"),
		KeyFile:viper.GetString("tls.KEY_FILE"),
		CaFile:viper.GetString("tls.CA_FILE"),
		ServerName:viper.GetString("tls.SERVER_NAME"),
	}

	c ,err:= clientTLS.GetTLSCredentialsByCA()
	if err != nil {
		log.Fatalf("clientTLS.GetTLSCredentialsByCA err:%v",err)
	}

	conn, err := grpc.Dial(":"+viper.GetString("port.SIMPLE_ZIPKIN_PORT"), grpc.WithTransportCredentials(c),
		grpc.WithUnaryInterceptor(
			otgrpc.OpenTracingClientInterceptor(tracer,otgrpc.LogPayloads()),
		),)
	if err != nil {
		log.Fatalf("grpc.Dial err:%v", err)
	}
	defer conn.Close()

	//deadline
	//d:=time.Now().Add(time.Duration(5*time.Second))
	//ctx,cancel:=context.WithDeadline(context.Background(),d)
	//defer cancel()

	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &pb.SearchRequest{
		Request: "gRPC zipkin",
	})
	if err != nil {
		//statusErr,ok:=status.FromError(err)
		//if ok {
		//	if statusErr.Code()==codes.DeadlineExceeded {
		//		log.Fatalln("client.Search err:deadline")
		//	}
		//}
		log.Fatalf("client.Search err:%v", err)
	}
	log.Printf("response :%s", resp.GetResponse())
}
