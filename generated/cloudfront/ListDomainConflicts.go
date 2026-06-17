package cloudfront

// ListDomainConflicts is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// We recommend that you use the ListDomainConflicts API operation to check for
// domain conflicts, as it supports both standard distributions and distribution
// tenants. [ListConflictingAliases]performs similar checks but only supports standard distributions.
//
// Lists existing domain associations that conflict with the domain that you
// specify.
//
// You can use this API operation to identify potential domain conflicts when
// moving domains between standard distributions and/or distribution tenants.
// Domain conflicts must be resolved first before they can be moved.
//
// For example, if you provide www.example.com as input, the returned list can
// include www.example.com and the overlapping wildcard alternate domain name (
// .example.com ), if they exist. If you provide .example.com as input, the
// returned list can include *.example.com and any alternate domain names covered
// by that wildcard (for example, www.example.com , test.example.com ,
// dev.example.com , and so on), if they exist.
//
// To list conflicting domains, specify the following:
//
// - The domain to search for
//
// - The ID of a standard distribution or distribution tenant in your account
// that has an attached TLS certificate, which covers the specified domain
//
// For more information, including how to set up the standard distribution or
// distribution tenant, and the certificate, see [Moving an alternate domain name to a different standard distribution or distribution tenant]in the Amazon CloudFront
// Developer Guide.
//
// You can optionally specify the maximum number of items to receive in the
// response. If the total number of items in the list exceeds the maximum that you
// specify, or the default maximum, the response is paginated. To get the next page
// of items, send a subsequent request that specifies the NextMarker value from
// the current response as the Marker value in the subsequent request.
//
// [ListConflictingAliases]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_ListConflictingAliases.html
// [Moving an alternate domain name to a different standard distribution or distribution tenant]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/CNAMEs.html#alternate-domain-names-move
