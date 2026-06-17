package health

// DescribeEventsForOrganization is generated as a reference stub.
// Executable command wiring lives under cmd/health.go.
//
// Returns information about events across your organization in Organizations. You
// can use the filters parameter to specify the events that you want to return.
// Events are returned in a summary form and don't include the affected accounts,
// detailed description, any additional metadata that depends on the event type, or
// any affected resources. To retrieve that information, use the following
// operations:
//
// [DescribeAffectedAccountsForOrganization]
//
// [DescribeEventDetailsForOrganization]
//
// [DescribeAffectedEntitiesForOrganization]
//
// If you don't specify a filter , the DescribeEventsForOrganizations returns all
// events across your organization. Results are sorted by lastModifiedTime ,
// starting with the most recent event.
//
// For more information about the different types of Health events, see [Event].
//
// Before you can call this operation, you must first enable Health to work with
// Organizations. To do this, call the [EnableHealthServiceAccessForOrganization]operation from your organization's
// management account.
//
// This API operation uses pagination. Specify the nextToken parameter in the next
// request to return more results.
//
// [DescribeEventDetailsForOrganization]: https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventDetailsForOrganization.html
// [DescribeAffectedEntitiesForOrganization]: https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeAffectedEntitiesForOrganization.html
// [DescribeAffectedAccountsForOrganization]: https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeAffectedAccountsForOrganization.html
// [Event]: https://docs.aws.amazon.com/health/latest/APIReference/API_Event.html
// [EnableHealthServiceAccessForOrganization]: https://docs.aws.amazon.com/health/latest/APIReference/API_EnableHealthServiceAccessForOrganization.html
