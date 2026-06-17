package route53resolver

// AssociateResolverQueryLogConfig is generated as a reference stub.
// Executable command wiring lives under cmd/route53resolver.go.
//
// Associates an Amazon VPC with a specified query logging configuration. Route 53
// Resolver logs DNS queries that originate in all of the Amazon VPCs that are
// associated with a specified query logging configuration. To associate more than
// one VPC with a configuration, submit one AssociateResolverQueryLogConfig
// request for each VPC.
//
// The VPCs that you associate with a query logging configuration must be in the
// same Region as the configuration.
//
// To remove a VPC from a query logging configuration, see [DisassociateResolverQueryLogConfig].
//
// [DisassociateResolverQueryLogConfig]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_DisassociateResolverQueryLogConfig.html
