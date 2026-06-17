package resourcegroupstaggingapi

// GetResources is generated as a reference stub.
// Executable command wiring lives under cmd/resourcegroupstaggingapi.go.
//
// Returns all the tagged or previously tagged resources that are located in the
// specified Amazon Web Services Region for the account.
//
// Depending on what information you want returned, you can also specify the
// following:
//
// - Filters that specify what tags and resource types you want returned. The
// response includes all tags that are associated with the requested resources.
//
// - Information about compliance with the account's effective tag policy. For
// more information on tag policies, see [Tag Policies]in the Organizations User Guide.
//
// This operation supports pagination, where the response can be sent in multiple
// pages. You should check the PaginationToken response parameter to determine if
// there are additional results available to return. Repeat the query, passing the
// PaginationToken response parameter value as an input to the next request until
// you recieve a null value. A null value for PaginationToken indicates that there
// are no more results waiting to be returned.
//
// GetResources does not return untagged resources.
//
// To find untagged resources in your account, use Amazon Web Services Resource
// Explorer with a query that uses tag:none . For more information, see [Search query syntax reference for Resource Explorer].
//
// [Search query syntax reference for Resource Explorer]: https://docs.aws.amazon.com/resource-explorer/latest/userguide/using-search-query-syntax.html
// [Tag Policies]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_tag-policies.html
