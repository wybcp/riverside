package interceptor

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"google.golang.org/grpc"
	"runtime/debug"
)

func LoggingInterceptor(ctx context.Context,req interface{},info *grpc.UnaryServerInfo,handler grpc.UnaryHandler)(interface{}, error){
	log.Printf("gRPC method: %s,%v",info.FullMethod,req)
	resp,err:=handler(ctx,req)
	log.Printf("gRPC method: %s,%v",info.FullMethod,resp)
	return resp,err
}
func RecoveryInterceptor(ctx context.Context,req interface{},info *grpc.UnaryServerInfo,handler grpc.UnaryHandler)(resp interface{},err error){
	log.Printf("gRPC method: %s,%v",info.FullMethod,req)

	log.Printf("gRPC method: %s,%v",info.FullMethod,resp)
	defer func() {
		if e:=recover();e!=nil{
			debug.PrintStack()
			err=status.Errorf(codes.Internal,"panic err:%v",e)
		}
	}()
	return handler(ctx,req)
}
