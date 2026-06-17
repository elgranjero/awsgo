package health

// DescribeAffectedEntities is generated as a reference stub.
// Executable command wiring lives under cmd/health.go.
//
// Returns a list of entities that have been affected by the specified events,
// based on the specified filter criteria. Entities can refer to individual
// customer resources, groups of customer resources, or any other construct,
// depending on the Amazon Web Services service. Events that have impact beyond
// that of the affected entities, or where the extent of impact is unknown, include
// at least one entity indicating this.
//
// At least one event ARN is required.
//
// - This API operation uses pagination. Specify the nextToken parameter in the
// next request to return more results.
//
// - This operation supports resource-level permissions. You can use this
// operation to allow or deny access to specific Health events. For more
// information, see [Resource- and action-based conditions]in the Health User Guide.
//
// [Resource- and action-based conditions]: https://docs.aws.amazon.com/health/latest/ug/security_iam_id-based-policy-examples.html#resource-action-based-conditions
