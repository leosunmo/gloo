package surveyutils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/kubernetes"

	"github.com/solo-io/gloo/pkg/cliutil/testutil"
	gatewayv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	. "github.com/solo-io/gloo/projects/gloo/cli/pkg/surveyutils"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

var _ = Describe("Route", func() {

	BeforeEach(func() {
		helpers.UseMemoryClients()

		vsClient := helpers.MustVirtualServiceClient()
		vs := &gatewayv1.VirtualService{
			Metadata: core.Metadata{
				Name:      "vs",
				Namespace: "gloo-system",
			},
			VirtualHost: &gatewayv1.VirtualHost{
				Routes: []*gatewayv1.Route{{
					Matcher: &v1.Matcher{
						PathSpecifier: &v1.Matcher_Prefix{Prefix: "/"},
					}}, {
					Matcher: &v1.Matcher{
						PathSpecifier: &v1.Matcher_Prefix{Prefix: "/r"},
					}},
				},
			},
		}
		_, err := vsClient.Write(vs, clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())

		usClient := helpers.MustUpstreamClient()
		us := &v1.Upstream{
			Metadata: core.Metadata{
				Name:      "gloo-system.some-ns-test-svc-1234",
				Namespace: "gloo-system",
			},
			UpstreamSpec: &v1.UpstreamSpec{
				UpstreamType: &v1.UpstreamSpec_Kube{
					Kube: &kubernetes.UpstreamSpec{
						ServiceName:      "test-svc",
						ServiceNamespace: "some-ns",
						ServicePort:      1234,
					},
				},
			},
		}
		_, err = usClient.Write(us, clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
	})

	It("should select a route", func() {
		testutil.ExpectInteractive(func(c *testutil.Console) {
			c.ExpectString("vsvc prompt:")
			c.SendLine("")
			c.ExpectString("route prompt:")
			c.PressDown()
			c.SendLine("")
			c.ExpectEOF()
		}, func() {
			var opts options.Options
			_, idx, err := SelectRouteInteractive(&opts, "vsvc prompt:", "route prompt:")
			Expect(err).NotTo(HaveOccurred())
			Expect(idx).To(Equal(1))
		})
	})

	It("should populate the correct flags", func() {
		testutil.ExpectInteractive(func(c *testutil.Console) {
			c.ExpectString("Choose a Virtual Service to add the route to")
			c.PressDown()
			c.SendLine("")
			c.ExpectString("where do you want to insert the route in the virtual service's route list?")
			c.SendLine("")
			c.ExpectString("Choose a path match type")
			c.SendLine("")
			c.ExpectString("What path prefix should we match?")
			c.SendLine("")
			c.ExpectString("Add a header matcher for this function (empty to skip)?")
			c.SendLine("")
			c.ExpectString("HTTP Method to match for this route (empty to skip)?")
			c.SendLine("")
			c.ExpectString("Choose the upstream or upstream group to route to:")
			c.SendLine("")
			c.ExpectString("do you wish to add a prefix-rewrite transformation to the route")
			c.SendLine("")
			c.ExpectEOF()
		}, func() {
			var opts options.Options
			err := AddRouteFlagsInteractive(&opts)
			Expect(err).NotTo(HaveOccurred())
			Expect(opts.Metadata.Name).To(Equal("vs"))
			Expect(opts.Metadata.Namespace).To(Equal("gloo-system"))
			Expect(opts.Add.Route.Matcher.PathPrefix).To(Equal("/"))
			Expect(opts.Add.Route.Destination.Upstream.Namespace).To(Equal("gloo-system"))
			Expect(opts.Add.Route.Destination.Upstream.Name).To(Equal("gloo-system.some-ns-test-svc-1234"))

		})
	})
})
