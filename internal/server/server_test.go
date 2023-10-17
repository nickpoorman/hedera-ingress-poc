package server

// import (
// 	"context"

// 	"io/ioutil"
// 	"net"
// 	"os"
// 	"testing"
// 	"time"

// 	"github.com/nickpoorman/hedera-ingress/internal/router"
// 	"github.com/nickpoorman/hedera-ingress/rpc"

// 	"github.com/stretchr/testify/require"
// 	"go.uber.org/zap"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials"

// 	"flag"

// 	"go.opencensus.io/examples/exporter"

// 	"github.com/nickpoorman/hedera-ingress/internal/config"
// 	"github.com/nickpoorman/hedera-ingress/internal/test/bindport"

// 	healthpb "google.golang.org/grpc/health/grpc_health_v1"
// )

// // debug/TestMain...
// var debug = flag.Bool("debug", false, "Enable observability for debugging.")

// func TestMain(m *testing.M) {
// 	flag.Parse()
// 	if *debug {
// 		logger, err := zap.NewDevelopment()
// 		if err != nil {
// 			panic(err)
// 		}
// 		zap.ReplaceGlobals(logger)
// 	}
// 	os.Exit(m.Run())
// }

// func TestServer(t *testing.T) {
// 	logger, err := zap.NewDevelopment()
// 	require.Nil(t, err)
// 	logger = logger.Named("TestServer")
// 	defer logger.Sync()

// 	for scenario, fn := range map[string]func(
// 		t *testing.T,
// 		clients clients,
// 		config *Config,
// 	){
// 		// Hapi server
// 		"forward a message to/from hapi succeeds": testForward,
// 		"healthcheck succeeds":                    testHealthCheck,
// 	} {
// 		t.Run(scenario, func(t *testing.T) {
// 			clients, config, teardown := setupTest(t, logger, nil)
// 			defer teardown()
// 			fn(t, clients, config)
// 		})
// 	}
// }

// type clients struct {
// 	Root   rpc.HapiClient
// 	Nobody rpc.HapiClient
// 	Health healthpb.HealthClient
// }

// func setupTest(t *testing.T, logger *zap.Logger, fn func(*Config)) (
// 	clients clients,
// 	cfg *Config,
// 	teardown func(),
// ) {
// 	t.Helper()

// 	ports := bindport.Get(3)

// 	rpcAddr := &net.TCPAddr{IP: []byte{127, 0, 0, 1}, Port: ports[0]}

// 	tlsConfig, err := config.SetupTLSConfig(config.TLSConfig{
// 		CertFile:      config.ServerCertFile,
// 		KeyFile:       config.ServerKeyFile,
// 		CAFile:        config.CAFile,
// 		Server:        true,
// 		ServerAddress: rpcAddr.IP.String(),
// 	})
// 	require.NoError(t, err)
// 	serverCreds := credentials.NewTLS(tlsConfig)

// 	l, err := net.Listen("tcp", rpcAddr.String())
// 	require.NoError(t, err)

// 	dataDir, err := ioutil.TempDir("", "server-test-hapi")
// 	require.NoError(t, err)
// 	defer os.RemoveAll(dataDir)

// 	router, err := router.NewRouter(logger, "test-router-0", l.Addr().String(), dataDir)
// 	require.NoError(t, err)

// 	var telemetryExporter *exporter.LogExporter
// 	if *debug {
// 		metricsLogFile, err := ioutil.TempFile("", "metrics-*.log")
// 		require.NoError(t, err)
// 		t.Logf("metrics log file: %s", metricsLogFile.Name())

// 		tracesLogFile, err := ioutil.TempFile("", "traces-*.log")
// 		require.NoError(t, err)
// 		t.Logf("traces log file: %s", tracesLogFile.Name())

// 		telemetryExporter, err = exporter.NewLogExporter(exporter.Options{
// 			MetricsLogFile:    metricsLogFile.Name(),
// 			TracesLogFile:     tracesLogFile.Name(),
// 			ReportingInterval: time.Second,
// 		})
// 		require.NoError(t, err)
// 		err = telemetryExporter.Start()
// 		require.NoError(t, err)
// 	}

// 	cfg = &Config{
// 		HapiConfig: &hapiServer.Config{
// 			GetHapiServerer: &HapiMock{},
// 			Logger:          logger,
// 		},
// 		Logger: logger,
// 	}
// 	if fn != nil {
// 		fn(cfg)
// 	}

// 	server, err := NewGRPCServer(cfg, grpc.Creds(serverCreds))
// 	require.NoError(t, err)

// 	go func() {
// 		if err := server.Serve(l); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	newHapiClient := func(crtPath, keyPath string) (
// 		*grpc.ClientConn,
// 		rpc.HapiClient,
// 		[]grpc.DialOption,
// 	) {
// 		tlsConfig, err := config.SetupTLSConfig(config.TLSConfig{
// 			CertFile:      crtPath,
// 			KeyFile:       keyPath,
// 			CAFile:        config.CAFile,
// 			Server:        false,
// 			ServerAddress: rpcAddr.IP.String(),
// 		})
// 		require.NoError(t, err)
// 		tlsCreds := credentials.NewTLS(tlsConfig)
// 		opts := []grpc.DialOption{grpc.WithTransportCredentials(tlsCreds)}
// 		conn, err := grpc.Dial(l.Addr().String(), opts...)
// 		require.NoError(t, err)
// 		client := rpc.NewHapiClient(conn)
// 		return conn, client, opts
// 	}

// 	var nobodyConn *grpc.ClientConn
// 	nobodyConn, clients.Nobody, _ = newHapiClient(config.NobodyClientCertFile, config.NobodyClientKeyFile)

// 	clients.Health = healthpb.NewHealthClient(nobodyConn)

// 	return clients, cfg, func() {
// 		server.Stop()
// 		nobodyConn.Close()
// 		l.Close()
// 		if telemetryExporter != nil {
// 			time.Sleep(1500 * time.Millisecond)
// 			telemetryExporter.Stop()
// 			telemetryExporter.Close()
// 		}
// 		router.Close()
// 	}
// }

// func testForward(t *testing.T, clients clients, config *Config) {
// 	//ctx := context.Background()

// 	// Send a gRPC request to hapi and get the response.
// 	t.Fatal("not yet implemented")
// }

// func testHealthCheck(
// 	t *testing.T,
// 	clients clients,
// 	config *Config,
// ) {
// 	ctx := context.Background()
// 	res, err := clients.Health.Check(ctx, &healthpb.HealthCheckRequest{})
// 	require.NoError(t, err)
// 	require.Equal(t, healthpb.HealthCheckResponse_SERVING, res.Status)
// }

// type hapiMock struct{}

// func (m *hapiMock) GetHapiClusterServers() ([]*rpc.LogClusterServer, error) {
// 	panic("not implemented for test")
// }
