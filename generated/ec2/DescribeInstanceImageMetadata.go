package ec2

// DescribeInstanceImageMetadata is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Describes the AMI that was used to launch an instance, even if the AMI is
// deprecated, deregistered, made private (no longer public or shared with your
// account), or not allowed.
//
// If you specify instance IDs, the output includes information for only the
// specified instances. If you specify filters, the output includes information for
// only those instances that meet the filter criteria. If you do not specify
// instance IDs or filters, the output includes information for all instances,
// which can affect performance.
//
// If you specify an instance ID that is not valid, an instance that doesn't
// exist, or an instance that you do not own, an error ( InvalidInstanceID.NotFound
// ) is returned.
//
// Recently terminated instances might appear in the returned results. This
// interval is usually less than one hour.
//
// In the rare case where an Availability Zone is experiencing a service
// disruption and you specify instance IDs that are in the affected Availability
// Zone, or do not specify any instance IDs at all, the call fails. If you specify
// only instance IDs that are in an unaffected Availability Zone, the call works
// normally.
//
// The order of the elements in the response, including those within nested
// structures, might vary. Applications should not assume the elements appear in a
// particular order.
