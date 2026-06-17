package wafv2

// DeleteWebACL is generated as a reference stub.
// Executable command wiring lives under cmd/wafv2.go.
//
// Deletes the specified WebACL.
//
// You can only use this if ManagedByFirewallManager is false in the web ACL.
//
// Before deleting any web ACL, first disassociate it from all resources.
//
// - To retrieve a list of the resources that are associated with a web ACL, use
// the following calls:
//
// - For Amazon CloudFront distributions, use the CloudFront call
// ListDistributionsByWebACLId . For information, see [ListDistributionsByWebACLId]in the Amazon CloudFront
// API Reference.
//
// - For all other resources, call ListResourcesForWebACL.
//
// - To disassociate a resource from a web ACL, use the following calls:
//
// - For Amazon CloudFront distributions, provide an empty web ACL ID in the
// CloudFront call UpdateDistribution . For information, see [UpdateDistribution]in the Amazon
// CloudFront API Reference.
//
// - For all other resources, call DisassociateWebACL.
//
// [ListDistributionsByWebACLId]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_ListDistributionsByWebACLId.html
// [UpdateDistribution]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_UpdateDistribution.html
