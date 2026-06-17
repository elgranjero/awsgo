package cloudfront

// UpdateDomainAssociation is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// We recommend that you use the UpdateDomainAssociation API operation to move a
// domain association, as it supports both standard distributions and distribution
// tenants. [AssociateAlias]performs similar checks but only supports standard distributions.
//
// Moves a domain from its current standard distribution or distribution tenant to
// another one.
//
// You must first disable the source distribution (standard distribution or
// distribution tenant) and then separately call this operation to move the domain
// to another target distribution (standard distribution or distribution tenant).
//
// To use this operation, specify the domain and the ID of the target resource
// (standard distribution or distribution tenant). For more information, including
// how to set up the target resource, prerequisites that you must complete, and
// other restrictions, see [Moving an alternate domain name to a different standard distribution or distribution tenant]in the Amazon CloudFront Developer Guide.
//
// [AssociateAlias]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_AssociateAlias.html
// [Moving an alternate domain name to a different standard distribution or distribution tenant]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/CNAMEs.html#alternate-domain-names-move
