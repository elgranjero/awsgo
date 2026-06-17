package health

// DescribeEventDetails is generated as a reference stub.
// Executable command wiring lives under cmd/health.go.
//
// Returns detailed information about one or more specified events. Information
// includes standard event data (Amazon Web Services Region, service, and so on, as
// returned by [DescribeEvents]), a detailed event description, and possible additional metadata
// that depends upon the nature of the event. Affected entities are not included.
// To retrieve the entities, use the [DescribeAffectedEntities]operation.
//
// If a specified event can't be retrieved, an error message is returned for that
// event.
//
// This operation supports resource-level permissions. You can use this operation
// to allow or deny access to specific Health events. For more information, see [Resource- and action-based conditions]in
// the Health User Guide.
//
// [DescribeEvents]: https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEvents.html
// [DescribeAffectedEntities]: https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeAffectedEntities.html
// [Resource- and action-based conditions]: https://docs.aws.amazon.com/health/latest/ug/security_iam_id-based-policy-examples.html#resource-action-based-conditions
