package syncer

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v2 "github.com/solo-io/gloo/projects/gateway/pkg/api/v2"
	"github.com/solo-io/gloo/projects/gateway/pkg/defaults"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

var _ = Describe("TranslatorSyncer unit tests", func() {
	Describe("gatewaysByProxyName", func() {
		It("assigns each gateway once to each proxy by their proxyNames", func() {

			gws := v2.GatewayList{
				{Metadata: core.Metadata{Name: "gw1"}, ProxyNames: nil /*default proxy*/},
				{Metadata: core.Metadata{Name: "gw2"}, ProxyNames: []string{"proxy1", "proxy2"}},
				{Metadata: core.Metadata{Name: "gw3"}, ProxyNames: []string{"proxy1", defaults.GatewayProxyName}},
			}

			gw1, gw2, gw3 := gws[0], gws[1], gws[2]

			byProxy := gatewaysByProxyName(gws)
			Expect(byProxy).To(Equal(map[string]v2.GatewayList{
				defaults.GatewayProxyName: {gw1, gw3},
				"proxy1":                  {gw2, gw3},
				"proxy2":                  {gw2},
			}))
		})
	})
})
