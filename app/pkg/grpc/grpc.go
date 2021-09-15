package grpc

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

type clientWithContextKeyType struct{ name string }

var ClientWithContextKey = clientWithContextKeyType{"_client_with_context"}

var _ ClientConn = (*clientConn)(nil)

type ClientConn interface {
	i()
	Conn() *grpc.ClientConn
}

type clientConn struct {
	conn *grpc.ClientConn
}

func New() (ClientConn, error) {
	// TODO 需从配置文件中获取 暂未实现
	//target := "127.0.0.1:8006"
	//secret := "abcdef"
	//
	//clientInterceptor := NewClientInterceptor(func(message []byte) (authorization string, err error) {
	//	return GenerateSign(secret, message)
	//})
	//
	//conn, err := grpclient.New(target,
	//	grpclient.WithKeepAlive(keepAlive),
	//	grpclient.WithDialTimeout(time.Second*5),
	//	grpclient.WithUnaryInterceptor(clientInterceptor.UnaryInterceptor),
	//)
	//
	//return &clientConn{
	//	conn: conn,
	//}, err

	return nil, nil
}

func (c *clientConn) i() {}

func (c *clientConn) Conn() *grpc.ClientConn {
	return c.conn
}

func ContextWithValueAndTimeout(value interface{}, duration time.Duration) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), duration)
	return context.WithValue(ctx, ClientWithContextKey, value)
}
