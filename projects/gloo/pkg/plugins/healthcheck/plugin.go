package healthcheck

import (
	"context"

	envoyapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
)

// var pluginStage = plugins.DuringStage(plugins.OutAuthStage)

func NewPlugin() plugins.Plugin {
	return &plugin{}
}

type plugin struct {
	ctx context.Context
}

var _ plugins.UpstreamPlugin = new(plugin)

func (p *plugin) Init(params plugins.InitParams) error {
	p.ctx = params.Ctx
	return nil
}

func (p *plugin) ProcessUpstream(params plugins.Params, in *v1.Upstream, out *envoyapi.Cluster) error {

	return nil
}
