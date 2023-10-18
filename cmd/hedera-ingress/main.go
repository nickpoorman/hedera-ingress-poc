package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/nickpoorman/hedera-ingress/internal/config"
	"github.com/nickpoorman/hedera-ingress/internal/router"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

// build is the git version of this program. It is set using build flags in the Makerfile.
var build = "develop"
var serviceName = "hedera-ingress"

func main() {
	cli := &cli{}

	cmd := &cobra.Command{
		Use:     serviceName,
		PreRunE: cli.setupConfig,
		RunE:    cli.run,
	}

	if err := setupFlags(cmd); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

type cli struct {
	cfg cfg
}

type cfg struct {
	router.Config
	BuildVersion    string
	ServerTLSConfig config.TLSConfig
	Backend         struct {
		Addr               string
		Tls                bool
		CaFile             string
		ServerHostOverride string
	}
	ShutdownWaitDuration time.Duration
	TracingEnabled       bool
	MetricsBindAddr      string

	// GroupCache
	// GroupCacheSelf  string
	// GroupCachePeers []string
}

func setupFlags(cmd *cobra.Command) error {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	cmd.Flags().String("config-file", "", "Path to config file.")

	cmd.Flags().String("build-version", build, "The build version.")
	cmd.Flags().String("build-version-description", "Hedera Ingress API", "The build version description.")

	dataDir := path.Join(os.TempDir(), serviceName)
	cmd.Flags().String("data-dir",
		dataDir,
		"Directory to store data.")

	// TODO: Add shutdown kill override to everything.
	cmd.Flags().Duration("shutdown-wait-duration", 5*time.Second, "Duration to wait for graceful shutdown before application kills iteself.")

	cmd.Flags().Bool("enable-tracing", true, "Enable tracing.")

	cmd.Flags().String("metrics-bind-addr", "0.0.0.0:2223", "Address to bind metrics http server")

	cmd.Flags().String("router-name", hostname, "The router name")
	cmd.Flags().String("grpc-bind-addr", "0.0.0.0:50221", "Address to bind gRPC server.")

	cmd.Flags().String("mirror-node-url", "https://mainnet-public.mirrornode.hedera.com", "The mirror node url. i.e. https://mainnet-public.mirrornode.hedera.com")

	// If we ever want to support ACLs.
	// cmd.Flags().String("acl-model-file", "", "Path to ACL model.")
	// cmd.Flags().String("acl-policy-file", "", "Path to ACL policy.")

	// Frontend connections.
	cmd.Flags().String("server-tls-cert-file", "", "Path to server tls cert.")
	cmd.Flags().String("server-tls-key-file", "", "Path to server tls key.")
	cmd.Flags().String("server-tls-ca-file",
		"",
		"Path to server certificate authority.")

	// Backend connections.
	cmd.Flags().Bool("backend-tls", false, "Connection uses TLS if true, else plain TCP.")
	cmd.Flags().String("backend-ca-file", "", "Path to the file containing the CA root cert file.")
	cmd.Flags().String("backend-addr", "127.0.0.1:50213", "The server address in the format of host:port.")
	cmd.Flags().String("backend-server-host-override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake.")

	// NFT based limits.
	cmd.Flags().Bool("nft-based-limits-enabled", true, "Enable NFT based limits.")

	// Group cache so that all nodes can work off consistant state.
	// cmd.Flags().String("groupcache-self", "127.0.0.1:50222", "The address of this machine's group cache.")
	// cmd.Flags().StringSlice("groupcache-peer", []string{"127.0.0.1:50222"}, "The peers this node will connect to for the group cache.")

	return viper.BindPFlags(cmd.Flags())
}

func (c *cli) setupConfig(cmd *cobra.Command, args []string) error {
	var err error

	configFile, err := cmd.Flags().GetString("config-file")
	if err != nil {
		return err
	}
	viper.SetConfigFile(configFile)

	if err = viper.ReadInConfig(); err != nil {
		// it's ok if config file doesn't exist
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	c.cfg.BuildVersion = viper.GetString("build-version")

	c.cfg.Config.DataDir = viper.GetString("data-dir")

	c.cfg.ShutdownWaitDuration = viper.GetDuration("shutdown-wait-duration")

	c.cfg.TracingEnabled = viper.GetBool("enable-tracing")

	c.cfg.MetricsBindAddr = viper.GetString("metrics-bind-addr")

	c.cfg.Config.BindAddr = viper.GetString("grpc-bind-addr")

	c.cfg.Config.RouterName = viper.GetString("router-name")

	// If we ever want to support ACLs.
	// c.cfg.ACLModelFile = viper.GetString("acl-model-file")
	// c.cfg.ACLPolicyFile = viper.GetString("acl-policy-file")

	// TODO: Wire this up to gRPC server.
	c.cfg.ServerTLSConfig.CertFile = viper.GetString("server-tls-cert-file")
	c.cfg.ServerTLSConfig.KeyFile = viper.GetString("server-tls-key-file")
	c.cfg.ServerTLSConfig.CAFile = viper.GetString("server-tls-ca-file")

	if c.cfg.ServerTLSConfig.CertFile != "" &&
		c.cfg.ServerTLSConfig.KeyFile != "" {
		c.cfg.ServerTLSConfig.Server = true
		c.cfg.Config.ServerTLSConfig, err = config.SetupTLSConfig(
			c.cfg.ServerTLSConfig,
		)
		if err != nil {
			return err
		}
	}

	c.cfg.Config.Logger = mustNewDevelopmentLogger()

	// Backend connections.
	c.cfg.Backend.Tls = viper.GetBool("backend-tls")
	c.cfg.Backend.CaFile = viper.GetString("backend-ca-file")
	c.cfg.Backend.Addr = viper.GetString("backend-addr")
	c.cfg.Backend.ServerHostOverride = viper.GetString("backend-server-host-override")

	// NFT based limits.
	c.cfg.Config.NFTBasedLimits = viper.GetBool("nft-based-limits-enabled")

	// Group cache peers.
	// c.cfg.GroupCacheSelf = viper.GetString("groupcache-self")
	// c.cfg.GroupCachePeers = viper.GetStringSlice("groupcache-peer")

	return nil
}

func (c *cli) run(cmd *cobra.Command, args []string) (err error) {
	// Set up tracing.
	if c.cfg.TracingEnabled {
		otelShutdown, err := setupOTelSDK(cmd.Context(), serviceName, c.cfg.BuildVersion)
		if err != nil {
			return fmt.Errorf("setupOTelSDK: %w", err)
		}
		// Handle shutdown properly so nothing leaks.
		defer func() {
			err = errors.Join(err, otelShutdown(context.Background()))
		}()
	}

	// Create a grpc client and dial the backend.
	var opts []grpc.DialOption
	if c.cfg.Backend.Tls {
		if c.cfg.Backend.CaFile == "" {
			return errors.New("TLS is enabled and missing CA file for backend TLS")
		}
		creds, err := credentials.NewClientTLSFromFile(c.cfg.Backend.CaFile, c.cfg.Backend.ServerHostOverride)
		if err != nil {
			return fmt.Errorf("failed to create TLS credentials: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	backendGrpcConn, err := grpc.Dial(c.cfg.Backend.Addr, opts...)
	if err != nil {
		return fmt.Errorf("failed to dial backend: %v", err)
	}
	defer backendGrpcConn.Close()

	router, err := router.NewRouter(c.cfg.Config, backendGrpcConn)
	if err != nil {
		return err
	}

	// HTTP server for metrics.
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	metricsServer := &http.Server{
		Addr:    c.cfg.MetricsBindAddr,
		Handler: mux,
	}
	go func() {
		c.cfg.Config.Logger.Info("serving metrics", zap.String("endpoint", "/metrics"), zap.String("addr", c.cfg.MetricsBindAddr))
		if err := metricsServer.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				c.cfg.Config.Logger.Info("stopped http /metrics server gracefully", zap.String("endpoint", "/metrics"), zap.String("addr", c.cfg.MetricsBindAddr))
			} else {
				c.cfg.Config.Logger.Error("error serving metrics", zap.String("endpoint", "/metrics"), zap.String("addr", c.cfg.MetricsBindAddr), zap.Error(err))
			}

			return
		}
	}()

	// TODO: If we want all the proxies to stay in sync we'll need to spin up etcd.
	//
	//
	//

	// Group cache server.
	// p := groupcache.NewHTTPPool(c.cfg.GroupCacheSelf)
	// if len(c.cfg.GroupCachePeers) > 0 {
	// 	p.Set(c.cfg.GroupCachePeers...)
	// }
	// groupCacheServer := &http.Server{
	// 	Addr:    c.cfg.GroupCacheSelf,
	// 	Handler: p,
	// }
	// go func() {
	// 	c.cfg.Config.Logger.Info("serving group cache", zap.String("addr", c.cfg.GroupCacheSelf), zap.Strings("peers", c.cfg.GroupCachePeers))
	// 	if err := groupCacheServer.ListenAndServe(); err != nil {
	// 		if err == http.ErrServerClosed {
	// 			c.cfg.Config.Logger.Info("stopped http group cache server gracefully", zap.String("addr", c.cfg.GroupCacheSelf), zap.Strings("peers", c.cfg.GroupCachePeers))
	// 		} else {
	// 			c.cfg.Config.Logger.Error("error serving group cache", zap.String("addr", c.cfg.GroupCacheSelf), zap.Strings("peers", c.cfg.GroupCachePeers), zap.Error(err))
	// 		}
	// 	}
	// }()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc

	// Shutdown the router.
	err = errors.Join(err, router.Shutdown())

	// Shutdown the server with a context timeout
	ctx, cancel := context.WithTimeout(context.Background(), c.cfg.ShutdownWaitDuration)
	defer cancel()
	if err := metricsServer.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown: %v\n", err)
		c.cfg.Config.Logger.Error("metrics server failed to shutdown gracefully", zap.String("endpoint", "/metrics"), zap.String("addr", c.cfg.MetricsBindAddr), zap.Error(err))
	}
	// if err := groupCacheServer.Shutdown(ctx); err != nil {
	// 	fmt.Printf("Server forced to shutdown: %v\n", err)
	// 	c.cfg.Config.Logger.Error("group cache server failed to shutdown gracefully", zap.String("addr", c.cfg.GroupCacheSelf), zap.Strings("peers", c.cfg.GroupCachePeers), zap.Error(err))
	// }
	return
}

func mustNewDevelopmentLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return logger
}

// setupOTelSDK bootstraps the OpenTelemetry pipeline.
// If it does not return an error, make sure to call shutdown for proper cleanup.
func setupOTelSDK(ctx context.Context, serviceName, serviceVersion string) (shutdown func(context.Context) error, err error) {
	var shutdownFuncs []func(context.Context) error

	// shutdown calls cleanup functions registered via shutdownFuncs.
	// The errors from the calls are joined.
	// Each registered cleanup will be invoked once.
	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	// handleErr calls shutdown for cleanup and makes sure that all errors are returned.
	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	// Setup resource.
	res, err := newResource(serviceName, serviceVersion)
	if err != nil {
		handleErr(err)
		return
	}

	// Setup trace provider.
	tracerProvider, err := newTraceProvider(ctx, res)
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
	otel.SetTracerProvider(tracerProvider)
	// otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	// Setup meter provider.
	meterProvider, err := newMeterProvider(res)
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
	otel.SetMeterProvider(meterProvider)

	return
}

func newResource(serviceName, serviceVersion string) (*resource.Resource, error) {
	return resource.Merge(resource.Default(),
		resource.NewWithAttributes(semconv.SchemaURL,
			semconv.ServiceName(serviceName),
			semconv.ServiceVersion(serviceVersion),
		))
}

func newTraceProvider(ctx context.Context, res *resource.Resource) (*trace.TracerProvider, error) {
	// Export the trace to std out.
	// traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())

	// Export the trace to the OTLP collector.
	// Override configuration with env vars.
	// https://github.com/open-telemetry/opentelemetry-go/blob/main/exporters/otlp/otlptrace/README.md?plain=1#L43
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter,
			// Default is 5s. Set to 1s for demonstrative purposes.
			trace.WithBatchTimeout(time.Second)),
		trace.WithResource(res),
	)
	return traceProvider, nil
}

func newMeterProvider(res *resource.Resource) (*metric.MeterProvider, error) {
	// Export the trace to std out.
	// metricExporter, err := stdoutmetric.New()

	// The exporter embeds a default OpenTelemetry Reader and
	// implements prometheus.Collector, allowing it to be used as
	// both a Reader and Collector.
	metricExporter, err := prometheus.New()
	if err != nil {
		return nil, err
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithResource(res),
		// metric.WithReader(metric.NewPeriodicReader(metricExporter,
		// 	// Default is 1m. Set to 3s for demonstrative purposes.
		// 	metric.WithInterval(3*time.Second))),
		metric.WithReader(metricExporter),
	)
	return meterProvider, nil
}
