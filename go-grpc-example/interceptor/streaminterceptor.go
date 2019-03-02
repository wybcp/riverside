package interceptor

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"runtime/debug"
)

func LoggingStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("before handling.gRPC method: %s", info.FullMethod)
	err := handler(srv, ss)
	log.Printf("after handling.gRPC method: %s,err: %v", info.FullMethod, err)
	return err
}
func RecoveryStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	log.Printf("before handling.gRPC method: %s", info.FullMethod)
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "panic err:%v", e)
		}
	}()
	err = handler(srv, ss)
	log.Printf("after handling.gRPC method: %s,%v", info.FullMethod, err)

	return
}
