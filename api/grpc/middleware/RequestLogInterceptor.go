// Package middleware contains grpc interceptors that are applied to both, grpc and http calls.
package middleware

import (
	"context"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type RequestLogInterceptor struct{}

// NewRequestLogInterceptor -
func NewRequestLogInterceptor() *RequestLogInterceptor {
	return &RequestLogInterceptor{}
}

// Unary -
func (r *RequestLogInterceptor) Unary() grpc.ServerOption {
	return grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		glog.Info("Hello from RequestLogInterceptor")
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			glog.Infof("INTERCEPTOR OLé! md method: %s ", md.Get("method"))
		}

		return handler(ctx, req)
	})
}
