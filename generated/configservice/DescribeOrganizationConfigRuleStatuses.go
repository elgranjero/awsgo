package configservice

// DescribeOrganizationConfigRuleStatuses is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Provides organization Config rule deployment status for an organization.
//
// The status is not considered successful until organization Config rule is
// successfully deployed in all the member accounts with an exception of excluded
// accounts.
//
// When you specify the limit and the next token, you receive a paginated
// response. Limit and next token are not applicable if you specify organization
// Config rule names. It is only applicable, when you request all the organization
// Config rules.
