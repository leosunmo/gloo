
---
title: "solo_jwt_authn.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `envoy.config.filter.http.solo_jwt_authn.v2` 
#### Types:


- [SoloJwtAuthnPerRoute](#solojwtauthnperroute)
- [ClaimToHeader](#claimtoheader)
- [ClaimToHeaders](#claimtoheaders)
  



##### Source File: [github.com/solo-io/gloo/projects/gloo/api/external/envoy/jwt/solo_jwt_authn.proto](https://github.com/solo-io/gloo/blob/master/projects/gloo/api/external/envoy/jwt/solo_jwt_authn.proto)





---
### SoloJwtAuthnPerRoute



```yaml
"requirement": string
"claimsToHeaders": map<string, .envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders>
"clearRouteCache": bool
"payloadInMetadata": string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `requirement` | `string` |  |  |
| `claimsToHeaders` | `map<string, .envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders>` |  |  |
| `clearRouteCache` | `bool` | clear the route cache if claims were added to the header |  |
| `payloadInMetadata` | `string` |  |  |




---
### ClaimToHeader

 
If this is specified, one of the claims will be copied to a header
and the route cache will be cleared.

```yaml
"claim": string
"header": string
"append": bool

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `claim` | `string` |  |  |
| `header` | `string` |  |  |
| `append` | `bool` |  |  |




---
### ClaimToHeaders



```yaml
"claims": []envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `claims` | [[]envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader](../solo_jwt_authn.proto.sk#claimtoheader) |  |  |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->
