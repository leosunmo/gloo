package e2e_test

import (
	"context"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/gogo/protobuf/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/als"

	gatewayv2 "github.com/solo-io/gloo/projects/gateway/pkg/api/v2"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/gloo/test/services"
	"github.com/solo-io/gloo/test/v1helpers"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

var _ = Describe("Gateway", func() {

	var (
		ctx            context.Context
		cancel         context.CancelFunc
		testClients    services.TestClients
		writeNamespace string
	)

	Describe("in memory", func() {

		BeforeEach(func() {
			ctx, cancel = context.WithCancel(context.Background())
			defaults.HttpPort = services.NextBindPort()
			defaults.HttpsPort = services.NextBindPort()

			writeNamespace = "gloo-system"
			ro := &services.RunOptions{
				NsToWrite: writeNamespace,
				NsToWatch: []string{"default", writeNamespace},
				WhatToRun: services.What{
					DisableGateway: false,
					DisableFds:     true,
					DisableUds:     true,
				},
			}

			testClients = services.RunGlooGatewayUdsFds(ctx, ro)

			// wait for the two gateways to be created.
			Eventually(func() (gatewayv2.GatewayList, error) {
				return testClients.GatewayClient.List(writeNamespace, clients.ListOpts{})
			}, "10s", "0.1s").Should(HaveLen(2))
		})

		AfterEach(func() {
			cancel()
		})

		Context("Access Logs", func() {

			var (
				envoyInstance *services.EnvoyInstance
				tu            *v1helpers.TestUpstream
			)

			TestUpstreamReachable := func() {
				v1helpers.TestUpstreamReachable(defaults.HttpPort, tu, nil)
			}

			BeforeEach(func() {
				ctx, cancel = context.WithCancel(context.Background())
				var err error
				envoyInstance, err = envoyFactory.NewEnvoyInstance()
				Expect(err).NotTo(HaveOccurred())

				tu = v1helpers.NewTestHttpUpstream(ctx, envoyInstance.LocalAddr())

				_, err = testClients.UpstreamClient.Write(tu.Upstream, clients.WriteOpts{})
				Expect(err).NotTo(HaveOccurred())

				err = envoyInstance.RunWithRole(writeNamespace+"~gateway-proxy-v2", testClients.GlooPort)
				Expect(err).NotTo(HaveOccurred())
			})

			AfterEach(func() {
				if envoyInstance != nil {
					_ = envoyInstance.Clean()
				}
			})

			Context("Grpc", func() {
				var (
					gw *gatewayv2.Gateway
				)

				BeforeEach(func() {
					gatewaycli := testClients.GatewayClient
					var err error
					gw, err = gatewaycli.Read("gloo-system", "gateway", clients.ReadOpts{})
					Expect(err).NotTo(HaveOccurred())
				})
				AfterEach(func() {
					gatewaycli := testClients.GatewayClient
					var err error
					gw, err = gatewaycli.Read("gloo-system", "gateway", clients.ReadOpts{})
					Expect(err).NotTo(HaveOccurred())
					gw.Plugins = nil
					_, err = gatewaycli.Write(gw, clients.WriteOpts{OverwriteExisting: true})
					Expect(err).NotTo(HaveOccurred())
				})

				It("can stream access logs")
			})

			Context("File", func() {
				var (
					gw   *gatewayv2.Gateway
					path string
				)

				var checkLogs = func(ei *services.EnvoyInstance, logsPresent func(logs string) bool) error {
					var (
						logs string
						err  error
					)

					if ei.UseDocker {
						logs, err = ei.Logs()
						if err != nil {
							return err
						}
					} else {
						file, err := os.OpenFile(ei.AccessLogs, os.O_RDONLY, 0777)
						if err != nil {
							return err
						}
						var byt []byte
						byt, err = ioutil.ReadAll(file)
						if err != nil {
							return err
						}
						logs = string(byt)
					}

					if logs == "" {
						return errors.Errorf("logs should not be empty")
					}
					if !logsPresent(logs) {
						return errors.Errorf("no access logs present")
					}
					return nil
				}

				BeforeEach(func() {
					gatewaycli := testClients.GatewayClient
					var err error
					gw, err = gatewaycli.Read("gloo-system", "gateway", clients.ReadOpts{})
					Expect(err).NotTo(HaveOccurred())
					path = "/dev/stdout"
					if !envoyInstance.UseDocker {
						tmpfile, err := ioutil.TempFile("", "")
						Expect(err).NotTo(HaveOccurred())
						path = tmpfile.Name()
						envoyInstance.AccessLogs = path
					}
				})
				AfterEach(func() {
					gatewaycli := testClients.GatewayClient
					var err error
					gw, err = gatewaycli.Read("gloo-system", "gateway", clients.ReadOpts{})
					Expect(err).NotTo(HaveOccurred())
					gw.Plugins = nil
					_, err = gatewaycli.Write(gw, clients.WriteOpts{OverwriteExisting: true})
					Expect(err).NotTo(HaveOccurred())
				})
				It("can create string access logs", func() {
					gw.Plugins = &gloov1.ListenerPlugins{
						AccessLoggingService: &als.AccessLoggingService{
							AccessLog: []*als.AccessLog{
								{
									OutputDestination: &als.AccessLog_FileSink{
										FileSink: &als.FileSink{
											Path: path,
											OutputFormat: &als.FileSink_StringFormat{
												StringFormat: "",
											},
										},
									},
								},
							},
						},
					}

					gatewaycli := testClients.GatewayClient
					_, err := gatewaycli.Write(gw, clients.WriteOpts{OverwriteExisting: true})
					Expect(err).NotTo(HaveOccurred())
					up := tu.Upstream
					vs := getTrivialVirtualServiceForUpstream("default", up.Metadata.Ref())
					_, err = testClients.VirtualServiceClient.Write(vs, clients.WriteOpts{})
					Expect(err).NotTo(HaveOccurred())
					TestUpstreamReachable()

					Eventually(func() error {
						var logsPresent = func(logs string) bool {
							return strings.Contains(logs, `"POST /1 HTTP/1.1" 200`)
						}
						return checkLogs(envoyInstance, logsPresent)
					}, time.Second*30, time.Second/2).ShouldNot(HaveOccurred())
				})
				It("can create json access logs", func() {
					gw.Plugins = &gloov1.ListenerPlugins{
						AccessLoggingService: &als.AccessLoggingService{
							AccessLog: []*als.AccessLog{
								{
									OutputDestination: &als.AccessLog_FileSink{
										FileSink: &als.FileSink{
											Path: path,
											OutputFormat: &als.FileSink_JsonFormat{
												JsonFormat: &types.Struct{
													Fields: map[string]*types.Value{
														"protocol": {
															Kind: &types.Value_StringValue{
																StringValue: "%PROTOCOL%",
															},
														},
														"method": {
															Kind: &types.Value_StringValue{
																StringValue: "%REQ(:METHOD)%",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					}
					gatewaycli := testClients.GatewayClient
					_, err := gatewaycli.Write(gw, clients.WriteOpts{OverwriteExisting: true})
					Expect(err).NotTo(HaveOccurred())
					up := tu.Upstream
					vs := getTrivialVirtualServiceForUpstream("default", up.Metadata.Ref())
					_, err = testClients.VirtualServiceClient.Write(vs, clients.WriteOpts{})
					Expect(err).NotTo(HaveOccurred())

					TestUpstreamReachable()
					Eventually(func() error {
						var logsPresent = func(logs string) bool {
							return strings.Contains(logs, `{"method":"POST","protocol":"HTTP/1.1"}`) ||
								strings.Contains(logs, `{"protocol":"HTTP/1.1","method":"POST"}`)
						}
						return checkLogs(envoyInstance, logsPresent)
					}, time.Second*30, time.Second/2).ShouldNot(HaveOccurred())
				})
			})
		})
	})
})
