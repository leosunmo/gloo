
---
title: "gateway.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `gateway.solo.io.v2` 
#### Types:


- [Gateway](#gateway) **Top-Level Resource**
- [HttpGateway](#httpgateway)
- [TcpGateway](#tcpgateway)
  



##### Source File: [github.com/solo-io/gloo/projects/gateway/api/v2/gateway.proto](https://github.com/solo-io/gloo/blob/master/projects/gateway/api/v2/gateway.proto)





---
### Gateway

 
A Gateway describes a single Listener (bind address:port)
and the routing configuration to upstreams that are reachable via a specific port on the Gateway Proxy itself.

```yaml
"ssl": bool
"bindAddress": string
"bindPort": int
"plugins": .gloo.solo.io.ListenerPlugins
"status": .core.solo.io.Status
"metadata": .core.solo.io.Metadata
"useProxyProto": .google.protobuf.BoolValue
"httpGateway": .gateway.solo.io.v2.HttpGateway
"tcpGateway": .gateway.solo.io.v2.TcpGateway
"gatewayProxyName": string
"proxyNames": []string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `ssl` | `bool` | if set to false, only use virtual services without ssl configured. if set to true, only use virtual services with ssl configured. |  |
| `bindAddress` | `string` | the bind address the gateway should serve traffic on |  |
| `bindPort` | `int` | bind ports must not conflict across gateways for a single proxy |  |
| `plugins` | [.gloo.solo.io.ListenerPlugins](../../../../gloo/api/v1/plugins.proto.sk#listenerplugins) | top level plugin configuration for all routes on the gateway |  |
| `status` | [.core.solo.io.Status](../../../../../../solo-kit/api/v1/status.proto.sk#status) | Status indicates the validation status of this resource. Status is read-only by clients, and set by gloo during validation |  |
| `metadata` | [.core.solo.io.Metadata](../../../../../../solo-kit/api/v1/metadata.proto.sk#metadata) | Metadata contains the object metadata for this resource |  |
| `useProxyProto` | [.google.protobuf.BoolValue](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/bool-value) | Enable ProxyProtocol support for this listener |  |
| `httpGateway` | [.gateway.solo.io.v2.HttpGateway](../gateway.proto.sk#httpgateway) |  |  |
| `tcpGateway` | [.gateway.solo.io.v2.TcpGateway](../gateway.proto.sk#tcpgateway) |  |  |
| `gatewayProxyName` | `string` | deprecated: use proxyNames |  |
| `proxyNames` | `[]string` | Names of the [`Proxy`](https://gloo.solo.io/v1/github.com/solo-io/gloo/projects/gloo/api/v1/proxy.proto.sk/) resources to generate from this gateway. If other gateways exist which point to the same proxy, Gloo will join them together. Proxies have a one-to-many relationship with Envoy bootstrap configuration. In order to connect to Gloo, the Envoy bootstrap configuration sets a `role` in the [node metadata](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/core/base.proto#envoy-api-msg-core-node) Envoy instances announce their `role` to Gloo, which maps to the `{{ .Namespace }}~{{ .Name }}` of the Proxy resource. The template for this value can be seen in the [Gloo Helm chart](https://github.com/solo-io/gloo/blob/master/install/helm/gloo/templates/9-gateway-proxy-configmap.yaml#L22) Note: this field also accepts fields written in camel-case. They will be converted to kebab-case in the Proxy name. This allows use of the [Gateway Name Helm value](https://github.com/solo-io/gloo/blob/master/install/helm/gloo/values-gateway-template.yaml#L47) for this field Defaults to `["gateway-proxy-v2"]` |  |




---
### HttpGateway



```yaml
"virtualServices": []core.solo.io.ResourceRef
"virtualServiceSelector": map<string, string>
"plugins": .gloo.solo.io.HttpListenerPlugins

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `virtualServices` | [[]core.solo.io.ResourceRef](../../../../../../solo-kit/api/v1/ref.proto.sk#resourceref) | names of the the virtual services, which contain the actual routes for the gateway if the list is empty, all virtual services will apply to this gateway (with accordance to tls flag above). |  |
| `virtualServiceSelector` | `map<string, string>` | Select virtual services by their label. This will apply only to virtual services in the same namespace as the gateway resource. only one of `virtualServices` or `virtualServiceSelector` should be provided |  |
| `plugins` | [.gloo.solo.io.HttpListenerPlugins](../../../../gloo/api/v1/plugins.proto.sk#httplistenerplugins) | http gateway configuration |  |




---
### TcpGateway



```yaml
"destinations": []gloo.solo.io.TcpHost
"plugins": .gloo.solo.io.TcpListenerPlugins

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `destinations` | [[]gloo.solo.io.TcpHost](../../../../gloo/api/v1/proxy.proto.sk#tcphost) | Name of the destinations the gateway can route to |  |
| `plugins` | [.gloo.solo.io.TcpListenerPlugins](../../../../gloo/api/v1/plugins.proto.sk#tcplistenerplugins) | tcp gateway configuration |  |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->
