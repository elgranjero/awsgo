package cloudfront

// AssociateAlias is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// The AssociateAlias API operation only supports standard distributions. To move
// domains between distribution tenants and/or standard distributions, we recommend
// that you use the [UpdateDomainAssociation]API operation instead.
//
// Associates an alias with a CloudFront standard distribution. An alias is
// commonly known as a custom domain or vanity domain. It can also be called a
// CNAME or alternate domain name.
//
// With this operation, you can move an alias that's already used for a standard
// distribution to a different standard distribution. This prevents the downtime
// that could occur if you first remove the alias from one standard distribution
// and then separately add the alias to another standard distribution.
//
// To use this operation, specify the alias and the ID of the target standard
// distribution.
//
// For more information, including how to set up the target standard distribution,
// prerequisites that you must complete, and other restrictions, see [Moving an alternate domain name to a different standard distribution or distribution tenant]in the Amazon
// CloudFront Developer Guide.
//
// [UpdateDomainAssociation]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_UpdateDomainAssociation.html
// [Moving an alternate domain name to a different standard distribution or distribution tenant]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/CNAMEs.html#alternate-domain-names-move
