package health

// DescribeAffectedEntitiesForOrganization is generated as a reference stub.
// Executable command wiring lives under cmd/health.go.
//
// Returns a list of entities that have been affected by one or more events for
// one or more accounts in your organization in Organizations, based on the filter
// criteria. Entities can refer to individual customer resources, groups of
// customer resources, or any other construct, depending on the Amazon Web Services
// service.
//
// At least one event Amazon Resource Name (ARN) and account ID are required.
//
// Before you can call this operation, you must first enable Health to work with
// Organizations. To do this, call the [EnableHealthServiceAccessForOrganization]operation from your organization's
// management account.
//
// - This API operation uses pagination. Specify the nextToken parameter in the
// next request to return more results.
//
// - This operation doesn't support resource-level permissions. You can't use
// this operation to allow or deny access to specific Health events. For more
// information, see [Resource- and action-based conditions]in the Health User Guide.
//
// [Resource- and action-based conditions]: https://docs.aws.amazon.com/health/latest/ug/security_iam_id-based-policy-examples.html#resource-action-based-conditions
// [EnableHealthServiceAccessForOrganization]: https://docs.aws.amazon.com/health/latest/APIReference/API_EnableHealthServiceAccessForOrganization.html
