package route53resolver

// DeleteResolverEndpoint is generated as a reference stub.
// Executable command wiring lives under cmd/route53resolver.go.
//
// Deletes a Resolver endpoint. The effect of deleting a Resolver endpoint depends
// on whether it's an inbound or an outbound Resolver endpoint:
//
// - Inbound: DNS queries from your network are no longer routed to the DNS
// service for the specified VPC.
//
// - Outbound: DNS queries from a VPC are no longer routed to your network.
