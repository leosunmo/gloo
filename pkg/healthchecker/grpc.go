package healthchecker

import (
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"

	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type grpcHealthChecker struct {
	srv *health.Server
	ok  uint32
}

var _ HealthChecker = new(grpcHealthChecker)

func NewGrpc(grpcHealthServer *health.Server) *grpcHealthChecker {
	ret := &grpcHealthChecker{}
	ret.ok = 1

	ret.srv = grpcHealthServer
	ret.srv.SetServingStatus("TestService", healthpb.HealthCheckResponse_SERVING)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM)

	go func() {
		<-sigterm
		atomic.StoreUint32(&ret.ok, 0)
		ret.srv.SetServingStatus("TestService", healthpb.HealthCheckResponse_NOT_SERVING)
	}()

	return ret
}

func (hc *grpcHealthChecker) Fail() {
	atomic.StoreUint32(&hc.ok, 0)
	hc.srv.SetServingStatus("TestService", healthpb.HealthCheckResponse_NOT_SERVING)
}

func (hc *grpcHealthChecker) GetServer() *health.Server {
	return hc.srv
}
