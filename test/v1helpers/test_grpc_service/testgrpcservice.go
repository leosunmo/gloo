package testgrpcservice

import (
	"context"
	"errors"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync/atomic"
	"syscall"
	"time"

	glootest "github.com/solo-io/gloo/test/v1helpers/test_grpc_service/glootest/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// For reflection to work, this uses the golang/proto plugin. To install it, run this command:
//    go get -u github.com/golang/protobuf/protoc-gen-go

// In the unlikely event that you need to re-generate this proto, open this directory in a terminal
// and run the following commands:
//    mkdir -p glootest
//    mkdir -p descriptors
//    protoc -I. --go_out=plugins=grpc:glootest --descriptor_set_out=descriptors/proto.pb protos/glootest.proto
//    protoc -I. --go_out=plugins=grpc:glootest --descriptor_set_out=descriptors/proto-nopkg.pb protos/glootest-nopackage.proto; sed -i 's/package glootest_nopackage/package glootest/' glootest/protos/glootest-nopackage.pb.go

func RunServer(ctx context.Context) *TestGRPCServer {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	srv := newServer()
	hc := NewHealthChecker(health.NewServer())
	healthpb.RegisterHealthServer(grpcServer, hc.grpc)
	glootest.RegisterTestServiceServer(grpcServer, srv)
	glootest.RegisterTestService2Server(grpcServer, srv)
	go grpcServer.Serve(lis)
	time.Sleep(time.Millisecond)

	addr := lis.Addr().String()
	_, portstr, err := net.SplitHostPort(addr)
	if err != nil {
		panic(err)
	}

	port, err := strconv.Atoi(portstr)
	if err != nil {
		panic(err)
	}

	srv.Port = uint32(port)
	srv.HealthChecker = hc

	return srv
}

func newServer() *TestGRPCServer {
	return &TestGRPCServer{
		C: make(chan *glootest.TestRequest),
	}
}

type TestGRPCServer struct {
	C             chan *glootest.TestRequest
	Port          uint32
	HealthChecker *healthChecker
}

// Returns a list of all shelves in the bookstore.
func (s *TestGRPCServer) TestMethod(_ context.Context, req *glootest.TestRequest) (*glootest.TestResponse, error) {
	if req == nil {
		return nil, errors.New("cannot be nil")
	}
	go func() {
		s.C <- req
	}()
	return &glootest.TestResponse{Str: req.Str}, nil
}

func (s *TestGRPCServer) TestMethod2(_ context.Context, req *glootest.TestRequest2) (*glootest.TestResponse2, error) {
	if req == nil {
		return nil, errors.New("cannot be nil")
	}
	go func() {
		s.C <- &glootest.TestRequest{Str: req.Str}
	}()
	return &glootest.TestResponse2{Str: req.Str}, nil
}

type healthChecker struct {
	grpc *health.Server
	ok   uint32
}

func NewHealthChecker(grpcHealthServer *health.Server) *healthChecker {
	ret := &healthChecker{}
	ret.ok = 1

	ret.grpc = grpcHealthServer
	ret.grpc.SetServingStatus("TestService", healthpb.HealthCheckResponse_SERVING)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM)

	go func() {
		<-sigterm
		atomic.StoreUint32(&ret.ok, 0)
		ret.grpc.SetServingStatus("TestService", healthpb.HealthCheckResponse_NOT_SERVING)
	}()

	return ret
}

func (hc *healthChecker) Fail() {
	atomic.StoreUint32(&hc.ok, 0)
	hc.grpc.SetServingStatus("TestService", healthpb.HealthCheckResponse_NOT_SERVING)
}
