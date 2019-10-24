# Public/Private split Gloo Gateway
This is a basic public/private deployment utilizing a single Gateway-controller to control two separate proxies, one for private traffic and one for public traffic.

It's assumed that ingress in to the Kubernetes cluster is taken care of externally, such as a cloud provider's loadbalancer.

Common deployment artifacts such as CRDs, RBAC config and monitoring/observability has been excluded for brevity.

# Relationships
We're only using one Gateway-controller to consume multiple Gateway CRD resources. We are also only using one Settings resource to configure both of the Gateways/Proxies.
