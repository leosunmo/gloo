package translator

import (
	"fmt"

	envoylistener "github.com/envoyproxy/go-control-plane/envoy/api/v2/listener"
	envoyal "github.com/envoyproxy/go-control-plane/envoy/config/filter/accesslog/v2"
	envoyutil "github.com/envoyproxy/go-control-plane/pkg/util"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

func UpstreamToClusterName(upstream core.ResourceRef) string {

	// For non-namespaced resources, return only name
	if upstream.Namespace == "" {
		return upstream.Name
	}

	// Don't use dots in the name as it messes up prometheus stats
	return fmt.Sprintf("%s_%s", upstream.Name, upstream.Namespace)
}

func NewFilterWithConfig(name string, config proto.Message) (envoylistener.Filter, error) {

	s := envoylistener.Filter{
		Name: name,
	}

	if config != nil {
		marshalledConf, err := envoyutil.MessageToStruct(config)
		if err != nil {
			// this should NEVER HAPPEN!
			return envoylistener.Filter{}, err
		}

		s.ConfigType = &envoylistener.Filter_Config{
			Config: marshalledConf,
		}
	}

	return s, nil
}

func NewAccessLogWithConfig(config proto.Message) (envoyal.AccessLog, error) {
	s := envoyal.AccessLog{
		Name: envoyutil.FileAccessLog,
	}

	if config != nil {
		marshalledConf, err := envoyutil.MessageToStruct(config)
		if err != nil {
			// this should NEVER HAPPEN!
			return envoyal.AccessLog{}, err
		}

		s.ConfigType = &envoyal.AccessLog_Config{
			Config: marshalledConf,
		}
	}

	return s, nil
}

func ParseConfig(c configObject, config proto.Message) error {
	any := c.GetTypedConfig()
	if any != nil {
		return types.UnmarshalAny(any, config)
	}
	structt := c.GetConfig()
	if structt != nil {
		return envoyutil.StructToMessage(structt, config)
	}
	return nil
}

type configObject interface {
	GetConfig() *types.Struct
	GetTypedConfig() *types.Any
}
