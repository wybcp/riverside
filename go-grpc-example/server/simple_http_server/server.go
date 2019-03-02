package main

import (
	"context"

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

// PORT 服务端口
const PORT = "9003"

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

func main() {
	certFle := "src/riverside/go-grpc-example/conf/server.pem"
	keyFile := "src/riverside/go-grpc-example/conf/server.key"
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
		),
	}
	mux := GetHTTPServerMux()
	server := grpc.NewServer(opts...)

	pb.RegisterSearchServiceServer(server, &SearchService{})
	err = http.ListenAndServeTLS(":"+PORT,
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
