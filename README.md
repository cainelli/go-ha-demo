* VirtualServices need to be generated for all Services and bind to the internal ingress gateway.
* Existing VirtualServices need to add the internal ingress gateway to the list of `gateways`
* IstioOperator creates a Classic Load Balancer. Need to remove it.
* Terraform/Cloudformation SercurityGroups/Routing/Peering between two VPCs.
* CloudFormation NLB pointing to internal NLB





Issue with loop requests

* In the scope of service entries it’s used for discovering pod ips and not to limit where configuration is applied.
* Sidecar doesn't work (at least not in istio-system) for filtering out ServiceEntries.
* `exportTo` doesn't seem to work in ServiceEntry. Is always visible for
