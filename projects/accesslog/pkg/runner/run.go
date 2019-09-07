package runner

import (
	"context"
	"fmt"
	"net"

	pb "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v2"
	"github.com/solo-io/gloo/pkg/healthchecker"
	"github.com/solo-io/gloo/projects/accesslog/pkg/loggingservice"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/stats"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func init() {
	view.Register(ocgrpc.DefaultServerViews...)
}

func Run() {
	clientSettings := NewSettings()
	ctx := context.Background()

	if clientSettings.DebugPort != 0 {

		debugPort := fmt.Sprintf("%d", clientSettings.DebugPort)
		// TODO(yuval-k): we need to start the stats server before calling contextutils
		// need to think of a better way to express this dependency, or preferably, fix it.
		stats.StartStatsServerWithPort(debugPort)
	}

	service := loggingservice.NewServer(false, nil)

	err := RunWithSettings(ctx, service, clientSettings)

	if err != nil {
		if ctx.Err() == nil {
			// not a context error - panic
			panic(err)
		}
	}
}

func RunWithSettings(ctx context.Context, service *loggingservice.Server, clientSettings Settings) error {
	ctx = contextutils.WithLogger(ctx, "extauth")

	err := StartAccessLog(ctx, clientSettings, service)
	if ctx.Err() != nil {
		return ctx.Err()
	}
	return err
}

func StartAccessLog(ctx context.Context, clientSettings Settings, service *loggingservice.Server) error {
	srv := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))

	pb.RegisterAccessLogServiceServer(srv, service)
	hc := healthchecker.NewGrpc("AccessLog", health.NewServer())
	healthpb.RegisterHealthServer(srv, hc.GetServer())
	reflection.Register(srv)

	logger := contextutils.LoggerFrom(ctx)
	logger.Infow("Starting access-log server")

	addr := fmt.Sprintf(":%d", clientSettings.ServerPort)
	runMode := "gRPC"
	network := "tcp"

	logger.Infof("access-log server running in [%s] mode, listening at [%s]", runMode, addr)
	lis, err := net.Listen(network, addr)
	if err != nil {
		logger.Errorw("Failed to announce on network", zap.Any("mode", runMode), zap.Any("address", addr), zap.Any("error", err))
		return err
	}
	go func() {
		<-ctx.Done()
		srv.Stop()
		_ = lis.Close()
	}()

	return srv.Serve(lis)
}
