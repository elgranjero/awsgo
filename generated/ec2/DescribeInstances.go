package ec2

// DescribeInstances is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Describes the specified instances or all instances.
//
// If you specify instance IDs, the output includes information for only the
// specified instances. If you specify filters, the output includes information for
// only those instances that meet the filter criteria. If you do not specify
// instance IDs or filters, the output includes information for all instances,
// which can affect performance. We recommend that you use pagination to ensure
// that the operation returns quickly and successfully.
//
// The response includes SQL license exemption status information for instances
// registered with the SQL LE service, providing visibility into license exemption
// configuration and status.
//
// If you specify an instance ID that is not valid, an error is returned. If you
// specify an instance that you do not own, it is not included in the output.
//
// Recently terminated instances might appear in the returned results. This
// interval is usually less than one hour.
//
// If you describe instances in the rare case where an Availability Zone is
// experiencing a service disruption and you specify instance IDs that are in the
// affected zone, or do not specify any instance IDs at all, the call fails. If you
// describe instances and specify only instance IDs that are in an unaffected zone,
// the call works normally.
//
// The Amazon EC2 API follows an eventual consistency model. This means that the
// result of an API command you run that creates or modifies resources might not be
// immediately available to all subsequent commands you run. For guidance on how to
// manage eventual consistency, see [Eventual consistency in the Amazon EC2 API]in the Amazon EC2 Developer Guide.
//
// We strongly recommend using only paginated requests. Unpaginated requests are
// susceptible to throttling and timeouts.
//
// The order of the elements in the response, including those within nested
// structures, might vary. Applications should not assume the elements appear in a
// particular order.
//
// [Eventual consistency in the Amazon EC2 API]: https://docs.aws.amazon.com/ec2/latest/devguide/eventual-consistency.html
