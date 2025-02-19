
---
title: "service.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `ingress.solo.io` 
#### Types:


- [KubeService](#kubeservice) **Top-Level Resource**
  



##### Source File: [github.com/solo-io/gloo/projects/ingress/api/v1/service.proto](https://github.com/solo-io/gloo/blob/master/projects/ingress/api/v1/service.proto)





---
### KubeService

 
A simple wrapper for a Kubernetes Service Object.

```yaml
"kubeServiceSpec": .google.protobuf.Any
"kubeServiceStatus": .google.protobuf.Any
"metadata": .core.solo.io.Metadata

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `kubeServiceSpec` | [.google.protobuf.Any](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/any) | a raw byte representation of the kubernetes service this resource wraps |  |
| `kubeServiceStatus` | [.google.protobuf.Any](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/any) | a raw byte representation of the service status of the kubernetes service object |  |
| `metadata` | [.core.solo.io.Metadata](../../../../../../solo-kit/api/v1/metadata.proto.sk#metadata) | Metadata contains the object metadata for this resource |  |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->
