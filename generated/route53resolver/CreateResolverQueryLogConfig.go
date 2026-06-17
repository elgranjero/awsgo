package route53resolver

// CreateResolverQueryLogConfig is generated as a reference stub.
// Executable command wiring lives under cmd/route53resolver.go.
//
// Creates a Resolver query logging configuration, which defines where you want
// Resolver to save DNS query logs that originate in your VPCs. Resolver can log
// queries only for VPCs that are in the same Region as the query logging
// configuration.
//
// To specify which VPCs you want to log queries for, you use
// AssociateResolverQueryLogConfig . For more information, see [AssociateResolverQueryLogConfig].
//
// You can optionally use Resource Access Manager (RAM) to share a query logging
// configuration with other Amazon Web Services accounts. The other accounts can
// then associate VPCs with the configuration. The query logs that Resolver creates
// for a configuration include all DNS queries that originate in all VPCs that are
// associated with the configuration.
//
// [AssociateResolverQueryLogConfig]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_AssociateResolverQueryLogConfig.html
