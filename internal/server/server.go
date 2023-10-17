package server

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/nickpoorman/hedera-ingress/internal/server/hapi"
	"github.com/nickpoorman/hedera-ingress/rpc"
)

type Config struct {
	// Add HAPI server config.
	Logger         *zap.Logger
	NFTBasedLimits bool
}

func NewGRPCServer(config *Config, backendGrpcConn *grpc.ClientConn, grpcOpts ...grpc.ServerOption) (
	*grpc.Server,
	error,
) {
	logger := config.Logger.With(zap.String("server", uuid.New().String())).Named("server")
	zapOpts := []grpc_zap.Option{
		grpc_zap.WithDurationField(
			func(duration time.Duration) zapcore.Field {
				return zap.Int64(
					"grpc.time_ns",
					duration.Nanoseconds(),
				)
			},
		),
	}

	// Add the grpc instrumentation.
	grpc.StatsHandler(otelgrpc.NewServerHandler())
	grpcOpts = append(grpcOpts,
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	grpcOpts = append(grpcOpts,
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_ctxtags.StreamServerInterceptor(),
				grpc_zap.StreamServerInterceptor(logger, zapOpts...),

				// If we ever want to support ACLs.
				// Inject authorization info into context.
				// grpc_auth.StreamServerInterceptor(proxyServer.Authenticate),

				rpc.StreamServerInterceptor(),
			)),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(),
				grpc_zap.UnaryServerInterceptor(logger, zapOpts...),

				// If we ever want to support ACLs.
				// Inject authorization info into context.
				// grpc_auth.UnaryServerInterceptor(proxyServer.Authenticate),

				rpc.UnaryServerInterceptor(),
			)),
		grpc.StatsHandler(&ocgrpc.ServerHandler{}),
	)

	gsrv := grpc.NewServer(grpcOpts...)

	// Register health server.
	hsrv := health.NewServer()
	hsrv.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(gsrv, hsrv)

	// Register any other gRPC servers we want...

	// Register the HAPI grpc servers.
	hapiServer := hapi.NewHapiServer(config.NFTBasedLimits)
	if err := hapiServer.RegisterServices(gsrv, backendGrpcConn); err != nil {
		return gsrv, fmt.Errorf("hapi register services: %w", err)
	}

	return gsrv, nil
}
