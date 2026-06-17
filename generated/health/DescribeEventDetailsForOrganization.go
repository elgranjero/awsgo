package health

// DescribeEventDetailsForOrganization is generated as a reference stub.
// Executable command wiring lives under cmd/health.go.
//
// Returns detailed information about one or more specified events for one or more
// Amazon Web Services accounts in your organization. This information includes
// standard event data (such as the Amazon Web Services Region and service), an
// event description, and (depending on the event) possible metadata. This
// operation doesn't return affected entities, such as the resources related to the
// event. To return affected entities, use the [DescribeAffectedEntitiesForOrganization]operation.
//
// Before you can call this operation, you must first enable Health to work with
// Organizations. To do this, call the [EnableHealthServiceAccessForOrganization]operation from your organization's
// management account.
//
// When you call the DescribeEventDetailsForOrganization operation, specify the
// organizationEventDetailFilters object in the request. Depending on the Health
// event type, note the following differences:
//
// - To return event details for a public event, you must specify a null value
// for the awsAccountId parameter. If you specify an account ID for a public
// event, Health returns an error message because public events aren't specific to
// an account.
//
// - To return event details for an event that is specific to an account in your
// organization, you must specify the awsAccountId parameter in the request. If
// you don't specify an account ID, Health returns an error message because the
// event is specific to an account in your organization.
//
// For more information, see [Event].
//
// This operation doesn't support resource-level permissions. You can't use this
// operation to allow or deny access to specific Health events. For more
// information, see [Resource- and action-based conditions]in the Health User Guide.
//
// [Resource- and action-based conditions]: https://docs.aws.amazon.com/health/latest/ug/security_iam_id-based-policy-examples.html#resource-action-based-conditions
// [DescribeAffectedEntitiesForOrganization]: https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeAffectedEntitiesForOrganization.html
// [Event]: https://docs.aws.amazon.com/health/latest/APIReference/API_Event.html
// [EnableHealthServiceAccessForOrganization]: https://docs.aws.amazon.com/health/latest/APIReference/API_EnableHealthServiceAccessForOrganization.html
