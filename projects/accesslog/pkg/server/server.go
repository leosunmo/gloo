package server

import (
	envoyals2 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v2"
)

// server is used to implement envoyauthv2.AuthorizationServer.
type Server struct {
}

func (server *Server) StreamAccessLogs(envoyals2.AccessLogService_StreamAccessLogsServer) error {
	panic("implement me")
}

var _ envoyals2.AccessLogServiceServer = &Server{}

func NewServer() *Server {
	var s Server
	return &s
}
