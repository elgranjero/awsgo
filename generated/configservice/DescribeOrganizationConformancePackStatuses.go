package configservice

// DescribeOrganizationConformancePackStatuses is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Provides organization conformance pack deployment status for an organization.
//
// The status is not considered successful until organization conformance pack is
// successfully deployed in all the member accounts with an exception of excluded
// accounts.
//
// When you specify the limit and the next token, you receive a paginated
// response. Limit and next token are not applicable if you specify organization
// conformance pack names. They are only applicable, when you request all the
// organization conformance packs.
