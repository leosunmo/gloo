package server

import (
	"fmt"

	envoyals "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v2"
)

// server is used to implement envoyals.AccessLogServiceServer.
type Server struct {
}
var _ envoyals.AccessLogServiceServer = new(Server)

func (server *Server) StreamAccessLogs(srv envoyals.AccessLogService_StreamAccessLogsServer) error {
	msg, err := srv.Recv()
	if err != nil {
		return err
	}
	fmt.Println(msg)
	return nil
}


func NewServer() *Server {
	var s Server
	return &s
}
