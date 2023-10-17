package rpc

import (
	context "context"
	"errors"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// AsGRPCStatusError will unwrap the errors until it finds one with a
// GRPCStatus().
func AsGRPCStatusError(err error) (error, bool) {
	var rpcError interface {
		error
		GRPCStatus() *status.Status
	}
	ok := errors.As(err, &rpcError)
	if !ok {
		return err, false
	}
	return rpcError, true
}

// UnaryServerInterceptor returns a new unary server interceptors that hoists
// rpc errors to be the error returned.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		err, _ = AsGRPCStatusError(err)
		return resp, err
	}
}

// StreamServerInterceptor returns a new streaming server interceptor that
// hoists rpc errors to be the error returned.
func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		err := handler(srv, stream)
		err, _ = AsGRPCStatusError(err)
		return err
	}
}
