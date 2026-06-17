package route53resolver

// DeleteResolverQueryLogConfig is generated as a reference stub.
// Executable command wiring lives under cmd/route53resolver.go.
//
// Deletes a query logging configuration. When you delete a configuration,
// Resolver stops logging DNS queries for all of the Amazon VPCs that are
// associated with the configuration. This also applies if the query logging
// configuration is shared with other Amazon Web Services accounts, and the other
// accounts have associated VPCs with the shared configuration.
//
// Before you can delete a query logging configuration, you must first
// disassociate all VPCs from the configuration. See [DisassociateResolverQueryLogConfig].
//
// If you used Resource Access Manager (RAM) to share a query logging
// configuration with other accounts, you must stop sharing the configuration
// before you can delete a configuration. The accounts that you shared the
// configuration with can first disassociate VPCs that they associated with the
// configuration, but that's not necessary. If you stop sharing the configuration,
// those VPCs are automatically disassociated from the configuration.
//
// [DisassociateResolverQueryLogConfig]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_DisassociateResolverQueryLogConfig.html
