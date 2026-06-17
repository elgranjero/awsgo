package cloudfront

// ListConflictingAliases is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// The ListConflictingAliases API operation only supports standard distributions.
// To list domain conflicts for both standard distributions and distribution
// tenants, we recommend that you use the [ListDomainConflicts]API operation instead.
//
// Gets a list of aliases that conflict or overlap with the provided alias, and
// the associated CloudFront standard distribution and Amazon Web Services accounts
// for each conflicting alias. An alias is commonly known as a custom domain or
// vanity domain. It can also be called a CNAME or alternate domain name.
//
// In the returned list, the standard distribution and account IDs are partially
// hidden, which allows you to identify the standard distribution and accounts that
// you own, and helps to protect the information of ones that you don't own.
//
// Use this operation to find aliases that are in use in CloudFront that conflict
// or overlap with the provided alias. For example, if you provide www.example.com
// as input, the returned list can include www.example.com and the overlapping
// wildcard alternate domain name ( .example.com ), if they exist. If you provide
// .example.com as input, the returned list can include *.example.com and any
// alternate domain names covered by that wildcard (for example, www.example.com ,
// test.example.com , dev.example.com , and so on), if they exist.
//
// To list conflicting aliases, specify the alias to search and the ID of a
// standard distribution in your account that has an attached TLS certificate that
// includes the provided alias. For more information, including how to set up the
// standard distribution and certificate, see [Moving an alternate domain name to a different standard distribution or distribution tenant]in the Amazon CloudFront Developer
// Guide.
//
// You can optionally specify the maximum number of items to receive in the
// response. If the total number of items in the list exceeds the maximum that you
// specify, or the default maximum, the response is paginated. To get the next page
// of items, send a subsequent request that specifies the NextMarker value from
// the current response as the Marker value in the subsequent request.
//
// [Moving an alternate domain name to a different standard distribution or distribution tenant]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/CNAMEs.html#alternate-domain-names-move
// [ListDomainConflicts]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_ListDomainConflicts.html
