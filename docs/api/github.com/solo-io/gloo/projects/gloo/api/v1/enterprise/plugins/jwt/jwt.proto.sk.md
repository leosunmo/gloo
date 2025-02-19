
---
title: "jwt.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `jwt.plugins.gloo.solo.io` 
#### Types:


- [RemoteJwks](#remotejwks)
- [LocalJwks](#localjwks)
- [Jwks](#jwks)
- [TokenSource](#tokensource)
- [HeaderSource](#headersource)
- [ClaimToHeader](#claimtoheader)
- [Provider](#provider)
- [VhostExtension](#vhostextension)
- [RouteExtension](#routeextension)
  



##### Source File: [github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/plugins/jwt/jwt.proto](https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/enterprise/plugins/jwt/jwt.proto)





---
### RemoteJwks



```yaml
"url": string
"upstreamRef": .core.solo.io.ResourceRef
"cacheDuration": .google.protobuf.Duration

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `url` | `string` | The url used when accessing the upstream for Json Web Key Set. This is used to set the host and path in the request |  |
| `upstreamRef` | [.core.solo.io.ResourceRef](../../../../../../../../../solo-kit/api/v1/ref.proto.sk#resourceref) | The Upstream representing the Json Web Key Set server |  |
| `cacheDuration` | [.google.protobuf.Duration](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/duration) | Duration after which the cached JWKS should be expired. If not specified, default cache duration is 5 minutes. |  |




---
### LocalJwks



```yaml
"key": string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `key` | `string` | Inline key. this can be json web key, key-set or PEM format. |  |




---
### Jwks



```yaml
"remote": .jwt.plugins.gloo.solo.io.RemoteJwks
"local": .jwt.plugins.gloo.solo.io.LocalJwks

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `remote` | [.jwt.plugins.gloo.solo.io.RemoteJwks](../jwt.proto.sk#remotejwks) | Use a remote JWKS server |  |
| `local` | [.jwt.plugins.gloo.solo.io.LocalJwks](../jwt.proto.sk#localjwks) | Use an inline JWKS |  |




---
### TokenSource

 
Describes the location of a JWT token

```yaml
"headers": []jwt.plugins.gloo.solo.io.TokenSource.HeaderSource
"queryParams": []string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `headers` | [[]jwt.plugins.gloo.solo.io.TokenSource.HeaderSource](../jwt.proto.sk#headersource) | Try to retrieve token from these headers |  |
| `queryParams` | `[]string` | Try to retrieve token from these query params |  |




---
### HeaderSource

 
Describes how to retrieve a JWT from a header

```yaml
"header": string
"prefix": string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `header` | `string` | The name of the header. for example, "authorization" |  |
| `prefix` | `string` | Prefix before the token. for example, "Bearer " |  |




---
### ClaimToHeader

 
Allows copying verified claims to headers sent upstream

```yaml
"claim": string
"header": string
"append": bool

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `claim` | `string` | Claim name. for example, "sub" |  |
| `header` | `string` | The header the claim will be copied to. for example, "x-sub". |  |
| `append` | `bool` | If header exist, append to it, or set it. |  |




---
### Provider



```yaml
"jwks": .jwt.plugins.gloo.solo.io.Jwks
"audiences": []string
"issuer": string
"tokenSource": .jwt.plugins.gloo.solo.io.TokenSource
"keepToken": bool
"claimsToHeaders": []jwt.plugins.gloo.solo.io.ClaimToHeader

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `jwks` | [.jwt.plugins.gloo.solo.io.Jwks](../jwt.proto.sk#jwks) | The source for the keys to validate JWTs. |  |
| `audiences` | `[]string` | An incoming JWT must have an 'aud' claim and it must be in this list. |  |
| `issuer` | `string` | Issuer of the JWT. the 'iss' claim of the JWT must match this. |  |
| `tokenSource` | [.jwt.plugins.gloo.solo.io.TokenSource](../jwt.proto.sk#tokensource) | Where to find the JWT of the current provider. |  |
| `keepToken` | `bool` | Should the token forwarded upstream. if false, the header containing the token will be removed. |  |
| `claimsToHeaders` | [[]jwt.plugins.gloo.solo.io.ClaimToHeader](../jwt.proto.sk#claimtoheader) | What claims should be copied to upstream headers. |  |




---
### VhostExtension



```yaml
"jwks": .jwt.plugins.gloo.solo.io.Jwks
"audiences": []string
"issuer": string
"providers": map<string, .jwt.plugins.gloo.solo.io.Provider>

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `jwks` | [.jwt.plugins.gloo.solo.io.Jwks](../jwt.proto.sk#jwks) | The source for the keys to validate JWTs. Deprecated: this field is deprecated, use `providers` instead. |  |
| `audiences` | `[]string` | An incoming JWT must have an 'aud' claim and it must be in this list. Deprecated: this field is deprecated, use `providers` instead. |  |
| `issuer` | `string` | Issuer of the JWT. the 'iss' claim of the JWT must match this. Deprecated: this field is deprecated, use `providers` instead. |  |
| `providers` | `map<string, .jwt.plugins.gloo.solo.io.Provider>` | Auth providers can be used instead of the fields above where more than one is required. if this list is provided the fields above are ignored. |  |




---
### RouteExtension



```yaml
"disable": bool

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `disable` | `bool` | Disable JWT checks on this route. |  |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->
