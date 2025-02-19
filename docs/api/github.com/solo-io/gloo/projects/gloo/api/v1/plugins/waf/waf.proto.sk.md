
---
title: "waf.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `waf.plugins.gloo.solo.io` 
#### Types:


- [Settings](#settings)
- [CoreRuleSet](#coreruleset)
- [VhostSettings](#vhostsettings)
- [RouteSettings](#routesettings)
  



##### Source File: [github.com/solo-io/gloo/projects/gloo/api/v1/plugins/waf/waf.proto](https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/plugins/waf/waf.proto)





---
### Settings



```yaml
"disabled": bool
"coreRuleSet": .waf.plugins.gloo.solo.io.CoreRuleSet
"ruleSets": []envoy.config.filter.http.modsecurity.v2.RuleSet

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `disabled` | `bool` | disable waf on this listener |  |
| `coreRuleSet` | [.waf.plugins.gloo.solo.io.CoreRuleSet](../waf.proto.sk#coreruleset) | Add owasp core rule set if nil will not be added |  |
| `ruleSets` | [[]envoy.config.filter.http.modsecurity.v2.RuleSet](../../../../external/envoy/waf/waf.proto.sk#ruleset) | custom rule sets rules to add |  |




---
### CoreRuleSet



```yaml
"customSettingsString": string
"customSettingsFile": string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `customSettingsString` | `string` | String representing the core rule set custom config options |  |
| `customSettingsFile` | `string` | String representing the core rule set custom config options |  |




---
### VhostSettings



```yaml
"disabled": bool
"settings": .waf.plugins.gloo.solo.io.Settings

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `disabled` | `bool` | disable waf on this virtual host |  |
| `settings` | [.waf.plugins.gloo.solo.io.Settings](../waf.proto.sk#settings) |  |  |




---
### RouteSettings



```yaml
"disabled": bool
"settings": .waf.plugins.gloo.solo.io.Settings

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `disabled` | `bool` | disable waf on this route |  |
| `settings` | [.waf.plugins.gloo.solo.io.Settings](../waf.proto.sk#settings) |  |  |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->
