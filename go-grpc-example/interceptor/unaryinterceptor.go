package interceptor

import (
	"context"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"runtime/debug"
)

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("gRPC method: %s,%v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	log.Printf("gRPC method: %s,%v", info.FullMethod, resp)
	return resp, err
}
func RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Printf("gRPC method: %s,%v", info.FullMethod, req)

	log.Printf("gRPC method: %s,%v", info.FullMethod, resp)
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "panic err:%v", e)
		}
	}()
	return handler(ctx, req)
}
func TokenInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("gRPC token start")
	md,ok:=metadata.FromIncomingContext(ctx)
	if !ok{
		return nil,status.Errorf(codes.Unauthenticated,"自定义认证 token 失败 ，没有传入token信息")
	}
	var appKey  string
	var appSecret string
	if value,ok:=md["app_key"];ok {
		appKey=value[0]
	}
	if value,ok:=md["app_secret"];ok {
		appSecret=value[0]
	}
	a:=Auth{
	}
	if appKey!=a.GetAppKey()||appSecret!=a.GetAppSecret() {
		return nil,status.Errorf(codes.Unauthenticated,"自定义认证 token 失败，token信息不匹配")
	}

	resp, err := handler(ctx, req)
	return resp, err
}
type Auth struct {
	appKey string
	appSecret string
}


func (a *Auth)GetAppKey()string  {
	return viper.GetString("token.APP_KEY")
}
func (a *Auth)GetAppSecret()string  {
	return viper.GetString("token.APP_SECRET")
}
