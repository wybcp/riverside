package simple_server

import (
	"context"
	"log"
	"net"

	pb "riverside/go-grpc-example/proto"

	"google.golang.org/grpc"
)

// SearchService 结构
type SearchService struct {
}

// PORT 服务端口
const PORT = "9001"

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}
func main() {
	server := grpc.NewServer()
	pb.RegisterSearchServiceServer(server, &SearchService{})
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	if err:=server.Serve(lis);err!=nil{
		log.Fatalf("server err: %v", err)
	}
}
