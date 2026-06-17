package health

// DescribeAffectedAccountsForOrganization is generated as a reference stub.
// Executable command wiring lives under cmd/health.go.
//
// Returns a list of accounts in the organization from Organizations that are
// affected by the provided event. For more information about the different types
// of Health events, see [Event].
//
// Before you can call this operation, you must first enable Health to work with
// Organizations. To do this, call the [EnableHealthServiceAccessForOrganization]operation from your organization's
// management account.
//
// This API operation uses pagination. Specify the nextToken parameter in the next
// request to return more results.
//
// [Event]: https://docs.aws.amazon.com/health/latest/APIReference/API_Event.html
// [EnableHealthServiceAccessForOrganization]: https://docs.aws.amazon.com/health/latest/APIReference/API_EnableHealthServiceAccessForOrganization.html
