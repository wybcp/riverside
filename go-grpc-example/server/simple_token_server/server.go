package main

import (
	"context"
	"github.com/spf13/viper"
	"riverside/go-grpc-example/pkg/ginit"

	"log"

	"net/http"
	"riverside/go-grpc-example/interceptor"
	"riverside/go-grpc-example/pkg/gtls"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	pb "riverside/go-grpc-example/proto"
)

// SearchService 结构
type SearchService struct {

}


func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + "token Server"}, nil
}

func main() {
	err:= ginit.InitViper()
	if err != nil {
		log.Fatalf("init.InitViper err:", err)
	}
	certFle := viper.GetString("tls.CERT_FILE")
	keyFile := viper.GetString("tls.KEY_FILE")

	tlsServer := gtls.ServerTLS{
		CertFile: certFle,
		KeyFile:  keyFile,
	}
	c, err := tlsServer.GetTLSCredentials()
	if err != nil {
		log.Fatalf("tlsServer.GetTLSCredentials err:", err)
	}
	opts := []grpc.ServerOption{
		grpc.Creds(c),
		grpc_middleware.WithUnaryServerChain(
			interceptor.LoggingInterceptor,
			interceptor.RecoveryInterceptor,
			interceptor.TokenInterceptor,
		),
	}
	mux := GetHTTPServerMux()
	server := grpc.NewServer(opts...)

	pb.RegisterSearchServiceServer(server, &SearchService{})

	err = http.ListenAndServeTLS(":"+viper.GetString("port.SIMPLE_TOKEN"),
		certFle,
		keyFile,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				server.ServeHTTP(w, r)
			} else {
				mux.ServeHTTP(w, r)
			}
			return
		}),
	)
	if err != nil {
		log.Fatalf("http.ListenAndServeTLS err:", err)
	}

}

func GetHTTPServerMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("http server :go-grpc-example"))
	})
	return mux
}
