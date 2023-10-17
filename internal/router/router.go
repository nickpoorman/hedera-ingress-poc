package router

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"sync"

	"github.com/google/uuid"
	"github.com/nickpoorman/hedera-ingress/internal/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	ErrNoAddresses = errors.New("no addresses provided")
)

type Config struct {
	// The router instance name.
	RouterName string

	// Deployment is a value to help us distinguish between
	// deployments of routers.
	Deployment string

	ServerTLSConfig *tls.Config
	ClientTLSConfig *tls.Config

	// BindAddr is the IP address : port the services bind to.
	BindAddr string

	// If we ever want to support ACLs.
	// ACLModelFile  string
	// ACLPolicyFile string

	// The logger the router will use and pass down.
	Logger *zap.Logger

	// DataDir stores data for the router.
	DataDir string

	// NFT Based Limits
	NFTBasedLimits bool
}

// RPCAddr returns a bound RPC address as the host that grpc is bound to and the
// RPC port.
func (c Config) RPCAddr() (string, string, error) {
	host, port, err := net.SplitHostPort(c.BindAddr)
	if err != nil {
		return "", "", fmt.Errorf("failed to split host and port: %w", err)
	}
	return host, port, nil
}

// Router binds grpc for ingress requests to the router.
type Router struct {
	Config Config

	logger   *zap.Logger
	listener net.Listener

	// backendRouter *backend.Router
	backendGrpcConn *grpc.ClientConn
	server          *grpc.Server

	shutdown     bool
	shutdownCh   chan struct{}
	shutdownLock sync.RWMutex
}

func NewRouter(config Config, backendGrpcConn *grpc.ClientConn) (*Router, error) {
	r := &Router{
		Config:          config,
		backendGrpcConn: backendGrpcConn,
		shutdownCh:      make(chan struct{}),
	}
	setup := []func() error{
		r.setupLogger,
		// r.setupBackends,
		r.setupTCP,
		r.setupServer,
	}
	for _, fn := range setup {
		if err := fn(); err != nil {
			return nil, err
		}
	}
	return r, nil
}

func (r *Router) setupLogger() error {
	id := r.Config.RouterName
	r.logger = r.Config.Logger.With(zap.String("router", uuid.New().String())).Named(fmt.Sprintf("router-%s", id))
	return nil
}

// func (r *Router) setupBackends() error {
// 	r.backendRouter = backend.NewRouter(backend.Config{Logger: r.logger})
// 	return nil
// }

func (r *Router) setupTCP() error {
	addr, err := net.ResolveTCPAddr("tcp", r.Config.BindAddr)
	if err != nil {
		return err
	}
	rpcAddr := fmt.Sprintf(
		"%s:%d",
		addr.IP.String(),
		addr.Port,
	)
	ln, err := net.Listen("tcp", rpcAddr)
	if err != nil {
		return err
	}
	r.listener = ln
	return nil
}

func (r *Router) setupServer() error {
	// If we ever want to support ACLs.
	// authorizer := auth.New(
	// 	r.Config.ACLModelFile,
	// 	r.Config.ACLPolicyFile,
	// )

	serverConfig := &server.Config{
		// TODO: Add HAPI server config.
		Logger:         r.logger,
		NFTBasedLimits: r.Config.NFTBasedLimits,
	}
	var opts []grpc.ServerOption
	if r.Config.ServerTLSConfig != nil {
		creds := credentials.NewTLS(r.Config.ServerTLSConfig)
		opts = append(opts, grpc.Creds(creds))
	}
	var err error
	r.server, err = server.NewGRPCServer(serverConfig, r.backendGrpcConn, opts...)
	if err != nil {
		return err
	}
	go func() {
		r.logger.Debug("server listening on", zap.String("address", r.listener.Addr().String()))
		if err := r.server.Serve(r.listener); err != nil {
			r.logger.Error("server serve error", zap.Error(err))
			if err := r.Shutdown(); err != nil {
				r.logger.Error("server shutdown error", zap.Error(err))
			}
		}
	}()
	return err
}

func (r *Router) Shutdown() error {
	r.shutdownLock.Lock()
	defer r.shutdownLock.Unlock()
	if r.shutdown {
		return nil
	}
	r.shutdown = true
	close(r.shutdownCh)

	shutdown := []func() error{
		func() error {
			r.logger.Debug("stopping grpc gracefully")
			r.server.GracefulStop()
			r.logger.Debug("stopped grpc gracefully")
			return nil
		},
		func() error {
			r.logger.Debug("closing listener")
			r.listener.Close()
			r.logger.Debug("closed listener")
			return nil
		},
		// Close backends after the grpc server has stopped accepting requests.
		// func() error {
		// 	r.logger.Debug("closing backends")
		// 	r.backendRouter.Close()
		// 	r.logger.Debug("closed backends")
		// 	return nil
		// },
		func() error {
			// Ignore these errors.
			// Uber made the same change within the core of the logger
			// implementation.
			// See: https://github.com/uber-go/zap/issues/328
			_ = r.logger.Sync()
			return nil
		},
	}
	for _, fn := range shutdown {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}
