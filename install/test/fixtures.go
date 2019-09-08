package test

var confWithoutTracing = `
node:
  cluster: gateway
  id: "{{.PodName}}.{{.PodNamespace}}"
  metadata:
    # role's value is the key for the in-memory xds cache (projects/gloo/pkg/xds/envoy.go)
    role: "{{.PodNamespace}}~gateway-proxy-v2"
static_resources:
  listeners:
    - name: prometheus_listener
      address:
        socket_address:
          address: 0.0.0.0
          port_value: 8081
      filter_chains:
        - filters:
            - name: envoy.http_connection_manager
              config:
                codec_type: auto
                stat_prefix: prometheus
                route_config:
                  name: prometheus_route
                  virtual_hosts:
                    - name: prometheus_host
                      domains:
                        - "*"
                      routes:
                        - match:
                            path: "/ready"
                            headers:
                            - name: ":method"
                              exact_match: GET
                          route:
                            cluster: admin_port_cluster
                        - match:
                            prefix: "/metrics"
                            headers:
                            - name: ":method"
                              exact_match: GET
                          route:
                            prefix_rewrite: "/stats/prometheus"
                            cluster: admin_port_cluster
                http_filters:
                  - name: envoy.router
                    config: {} # if $spec.stats # if $spec.tracing


  clusters:
  - name: gloo.gloo-system.svc.cluster.local:9977
    alt_stat_name: xds_cluster
    connect_timeout: 5.000s
    load_assignment:
      cluster_name: gloo.gloo-system.svc.cluster.local:9977
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              socket_address:
                address: gloo.gloo-system.svc.cluster.local
                port_value: 9977
    http2_protocol_options: {}
    upstream_connection_options:
      tcp_keepalive: {}
    type: STRICT_DNS
    respect_dns_ttl: true
  - name: admin_port_cluster
    connect_timeout: 5.000s
    type: STATIC
    lb_policy: ROUND_ROBIN
    load_assignment:
      cluster_name: admin_port_cluster
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              socket_address:
                address: 127.0.0.1
                port_value: 19000 # if $spec.stats

dynamic_resources:
  ads_config:
    api_type: GRPC
    grpc_services:
    - envoy_grpc: {cluster_name: gloo.gloo-system.svc.cluster.local:9977}
    rate_limit_settings: {}
  cds_config:
    ads: {}
  lds_config:
    ads: {}
admin:
  access_log_path: /dev/null
  address:
    socket_address:
      address: 127.0.0.1
      port_value: 19000 # if (empty $spec.configMap.data) ## allows full custom # range $name, $spec := .Values.gatewayProxies# if .Values.gateway.enabled
`

var confWithTracingProvider = `
node:
  cluster: gateway
  id: "{{.PodName}}.{{.PodNamespace}}"
  metadata:
    # role's value is the key for the in-memory xds cache (projects/gloo/pkg/xds/envoy.go)
    role: "{{.PodNamespace}}~gateway-proxy-v2"
static_resources:
  listeners:
    - name: prometheus_listener
      address:
        socket_address:
          address: 0.0.0.0
          port_value: 8081
      filter_chains:
        - filters:
            - name: envoy.http_connection_manager
              config:
                codec_type: auto
                stat_prefix: prometheus
                route_config:
                  name: prometheus_route
                  virtual_hosts:
                    - name: prometheus_host
                      domains:
                        - "*"
                      routes:
                        - match:
                            path: "/ready"
                            headers:
                            - name: ":method"
                              exact_match: GET
                          route:
                            cluster: admin_port_cluster
                        - match:
                            prefix: "/metrics"
                            headers:
                            - name: ":method"
                              exact_match: GET
                          route:
                            prefix_rewrite: "/stats/prometheus"
                            cluster: admin_port_cluster
                http_filters:
                  - name: envoy.router
                    config: {} # if $spec.stats
  clusters:
  - name: gloo.gloo-system.svc.cluster.local:9977
    alt_stat_name: xds_cluster
    connect_timeout: 5.000s
    load_assignment:
      cluster_name: gloo.gloo-system.svc.cluster.local:9977
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              socket_address:
                address: gloo.gloo-system.svc.cluster.local
                port_value: 9977
    http2_protocol_options: {}
    upstream_connection_options:
      tcp_keepalive: {}
    type: STRICT_DNS
    respect_dns_ttl: true # if $spec.tracing.cluster # if $spec.tracing
  - name: admin_port_cluster
    connect_timeout: 5.000s
    type: STATIC
    lb_policy: ROUND_ROBIN
    load_assignment:
      cluster_name: admin_port_cluster
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              socket_address:
                address: 127.0.0.1
                port_value: 19000 # if $spec.stats
tracing:
  http:
    another: line
    trace: spec
     # if $spec.tracing.provider # if $spec.tracing
dynamic_resources:
  ads_config:
    api_type: GRPC
    grpc_services:
    - envoy_grpc: {cluster_name: gloo.gloo-system.svc.cluster.local:9977}
    rate_limit_settings: {}
  cds_config:
    ads: {}
  lds_config:
    ads: {}
admin:
  access_log_path: /dev/null
  address:
    socket_address:
      address: 127.0.0.1
      port_value: 19000 # if (empty $spec.configMap.data) ## allows full custom # range $name, $spec := .Values.gatewayProxies# if .Values.gateway.enabled
`

var confWithTracingProviderCluster = `
node:
  cluster: gateway
  id: "{{.PodName}}.{{.PodNamespace}}"
  metadata:
    # role's value is the key for the in-memory xds cache (projects/gloo/pkg/xds/envoy.go)
    role: "{{.PodNamespace}}~gateway-proxy-v2"
static_resources:
  listeners:
    - name: prometheus_listener
      address:
        socket_address:
          address: 0.0.0.0
          port_value: 8081
      filter_chains:
        - filters:
            - name: envoy.http_connection_manager
              config:
                codec_type: auto
                stat_prefix: prometheus
                route_config:
                  name: prometheus_route
                  virtual_hosts:
                    - name: prometheus_host
                      domains:
                        - "*"
                      routes:
                        - match:
                            path: "/ready"
                            headers:
                            - name: ":method"
                              exact_match: GET
                          route:
                            cluster: admin_port_cluster
                        - match:
                            prefix: "/metrics"
                            headers:
                            - name: ":method"
                              exact_match: GET
                          route:
                            prefix_rewrite: "/stats/prometheus"
                            cluster: admin_port_cluster
                http_filters:
                  - name: envoy.router
                    config: {} # if $spec.stats
  clusters:
  - name: gloo.gloo-system.svc.cluster.local:9977
    alt_stat_name: xds_cluster
    connect_timeout: 5.000s
    load_assignment:
      cluster_name: gloo.gloo-system.svc.cluster.local:9977
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              socket_address:
                address: gloo.gloo-system.svc.cluster.local
                port_value: 9977
    http2_protocol_options: {}
    upstream_connection_options:
      tcp_keepalive: {}
    type: STRICT_DNS
    respect_dns_ttl: true
  - connect_timeout: 1s
    lb_policy: round_robin
    load_assignment:
      cluster_name: zipkin
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              socket_address:
                address: zipkin
                port_value: 1234
    name: zipkin
    type: STRICT_DNS
    respect_dns_ttl: true
   # if $spec.tracing.cluster # if $spec.tracing
  - name: admin_port_cluster
    connect_timeout: 5.000s
    type: STATIC
    lb_policy: ROUND_ROBIN
    load_assignment:
      cluster_name: admin_port_cluster
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              socket_address:
                address: 127.0.0.1
                port_value: 19000 # if $spec.stats
tracing:
  http:
    typed_config:
      '@type': type.googleapis.com/envoy.config.trace.v2.ZipkinConfig
      collector_cluster: zipkin
      collector_endpoint: /api/v1/spans
     # if $spec.tracing.provider # if $spec.tracing
dynamic_resources:
  ads_config:
    api_type: GRPC
    grpc_services:
    - envoy_grpc: {cluster_name: gloo.gloo-system.svc.cluster.local:9977}
    rate_limit_settings: {}
  cds_config:
    ads: {}
  lds_config:
    ads: {}
admin:
  access_log_path: /dev/null
  address:
    socket_address:
      address: 127.0.0.1
      port_value: 19000 # if (empty $spec.configMap.data) ## allows full custom # range $name, $spec := .Values.gatewayProxies# if .Values.gateway.enabled`

var confWithReadConfig = `
node:
  cluster: gateway
  id: "{{.PodName}}.{{.PodNamespace}}"
  metadata:
    # role's value is the key for the in-memory xds cache (projects/gloo/pkg/xds/envoy.go)
    role: "{{.PodNamespace}}~gateway-proxy-v2"
static_resources:
  listeners:
    - name: prometheus_listener
      address:
        socket_address:
          address: 0.0.0.0
          port_value: 8081
      filter_chains:
        - filters:
            - name: envoy.http_connection_manager
              config:
                codec_type: auto
                stat_prefix: prometheus
                route_config:
                  name: prometheus_route
                  virtual_hosts:
                    - name: prometheus_host
                      domains:
                        - "*"
                      routes:
                        - match:
                            path: "/ready"
                            headers:
                            - name: ":method"
                              exact_match: GET
                          route:
                            cluster: admin_port_cluster
                        - match:
                            prefix: "/metrics"
                            headers:
                            - name: ":method"
                              exact_match: GET
                          route:
                            prefix_rewrite: "/stats/prometheus"
                            cluster: admin_port_cluster
                http_filters:
                  - name: envoy.router
                    config: {} # if $spec.stats # if $spec.tracing
    - name: read_config_listener
      address:
        socket_address:
          address: 0.0.0.0
          port_value: 8082
      filter_chains:
        - filters:
          - name: envoy.http_connection_manager
            config:
              codec_type: auto
              stat_prefix: read_config
              route_config:
                name: read_config_route
                virtual_hosts:
                - name: read_config_host
                  domains:
                  - "*"
                  routes:
                  - match:
                      path: "/ready"
                      headers:
                        - name: ":method"
                          exact_match: GET
                    route:
                      cluster: admin_port_cluster
                  - match:
                      prefix: "/stats"
                      headers:
                        - name: ":method"
                          exact_match: GET
                    route:
                      cluster: admin_port_cluster
                  - match:
                      prefix: "/config_dump"
                      headers:
                        - name: ":method"
                          exact_match: GET
                    route:
                      cluster: admin_port_cluster
              http_filters:
                - name: envoy.router
                  config: {}
  clusters:
  - name: gloo.gloo-system.svc.cluster.local:9977
    alt_stat_name: xds_cluster
    connect_timeout: 5.000s
    load_assignment:
      cluster_name: gloo.gloo-system.svc.cluster.local:9977
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              socket_address:
                address: gloo.gloo-system.svc.cluster.local
                port_value: 9977
    http2_protocol_options: {}
    upstream_connection_options:
      tcp_keepalive: {}
    type: STRICT_DNS
    respect_dns_ttl: true
  - name: admin_port_cluster
    connect_timeout: 5.000s
    type: STATIC
    lb_policy: ROUND_ROBIN
    load_assignment:
      cluster_name: admin_port_cluster
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              socket_address:
                address: 127.0.0.1
                port_value: 19000 # if $spec.stats
dynamic_resources:
  ads_config:
    api_type: GRPC
    grpc_services:
    - envoy_grpc: {cluster_name: gloo.gloo-system.svc.cluster.local:9977}
    rate_limit_settings: {}
  cds_config:
    ads: {}
  lds_config:
    ads: {}
admin:
  access_log_path: /dev/null
  address:
    socket_address:
      address: 127.0.0.1
      port_value: 19000 # if (empty $spec.configMap.data) ## allows full custom # range $name, $spec := .Values.gatewayProxies# if .Values.gateway.enabled`
