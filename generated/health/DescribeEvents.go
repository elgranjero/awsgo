package health

// DescribeEvents is generated as a reference stub.
// Executable command wiring lives under cmd/health.go.
//
// Returns information about events that meet the specified filter criteria.
//
// Events are returned in a summary form and do not include the detailed
// description, any additional metadata that depends on the event type, or any
// affected resources. To retrieve that information, use the [DescribeEventDetails]and [DescribeAffectedEntities] operations.
//
// If no filter criteria are specified, all events are returned. Results are
// sorted by lastModifiedTime , starting with the most recent event.
//
// - When you call the DescribeEvents operation and specify an entity for the
// entityValues parameter, Health might return public events that aren't specific
// to that resource. For example, if you call DescribeEvents and specify an ID
// for an Amazon Elastic Compute Cloud (Amazon EC2) instance, Health might return
// events that aren't specific to that resource or service. To get events that are
// specific to a service, use the services parameter in the filter object. For
// more information, see [Event].
//
// - This API operation uses pagination. Specify the nextToken parameter in the
// next request to return more results.
//
// [DescribeEventDetails]: https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventDetails.html
// [DescribeAffectedEntities]: https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeAffectedEntities.html
// [Event]: https://docs.aws.amazon.com/health/latest/APIReference/API_Event.html
