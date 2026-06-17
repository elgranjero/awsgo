package resourcegroupstaggingapi

// GetComplianceSummary is generated as a reference stub.
// Executable command wiring lives under cmd/resourcegroupstaggingapi.go.
//
// Returns a table that shows counts of resources that are noncompliant with their
// tag policies.
//
// For more information on tag policies, see [Tag Policies] in the Organizations User Guide.
//
// You can call this operation only from the organization's management account and
// from the us-east-1 Region.
//
// This operation supports pagination, where the response can be sent in multiple
// pages. You should check the PaginationToken response parameter to determine if
// there are additional results available to return. Repeat the query, passing the
// PaginationToken response parameter value as an input to the next request until
// you recieve a null value. A null value for PaginationToken indicates that there
// are no more results waiting to be returned.
//
// [Tag Policies]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_tag-policies.html
