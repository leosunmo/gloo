package loggingservice

import (
	envoyals "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v2"
	"golang.org/x/sync/errgroup"
)

// server is used to implement envoyals.AccessLogServiceServer.

type AlsCallback func(message *envoyals.StreamAccessLogsMessage) error
type AlsCallbackList []AlsCallback

type Server struct {
	ordered   bool
	callbacks AlsCallbackList
}

var _ envoyals.AccessLogServiceServer = new(Server)

func (s *Server) StreamAccessLogs(srv envoyals.AccessLogService_StreamAccessLogsServer) error {
	msg, err := srv.Recv()
	if err != nil {
		return err
	}

	if s.ordered {
		for _, cb := range s.callbacks {
			if err := cb(msg); err != nil {
				return err
			}
		}
	} else {
		eg := errgroup.Group{}
		for _, cb := range s.callbacks {
			cb := cb
			eg.Go(func() error {
				return cb(msg)
			})
		}
		if err := eg.Wait(); err != nil {
			return err
		}
	}
	return nil
}

func NewServer(ordered bool, cb ...AlsCallback) *Server {
	return &Server{
		ordered:   ordered,
		callbacks: cb,
	}
}
