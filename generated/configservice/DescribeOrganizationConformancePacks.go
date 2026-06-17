package configservice

// DescribeOrganizationConformancePacks is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Returns a list of organization conformance packs.
//
// When you specify the limit and the next token, you receive a paginated
// response.
//
// Limit and next token are not applicable if you specify organization conformance
// packs names. They are only applicable, when you request all the organization
// conformance packs.
//
// # For accounts within an organization
//
// If you deploy an organizational rule or conformance pack in an organization
// administrator account, and then establish a delegated administrator and deploy
// an organizational rule or conformance pack in the delegated administrator
// account, you won't be able to see the organizational rule or conformance pack in
// the organization administrator account from the delegated administrator account
// or see the organizational rule or conformance pack in the delegated
// administrator account from organization administrator account. The
// DescribeOrganizationConfigRules and DescribeOrganizationConformancePacks APIs
// can only see and interact with the organization-related resource that were
// deployed from within the account calling those APIs.
