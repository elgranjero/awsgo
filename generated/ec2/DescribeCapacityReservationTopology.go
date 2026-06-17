package ec2

// DescribeCapacityReservationTopology is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Describes a tree-based hierarchy that represents the physical host placement of
// your pending or active Capacity Reservations within an Availability Zone or
// Local Zone. You can use this information to determine the relative proximity of
// your capacity within the Amazon Web Services network before it is launched and
// use this information to allocate capacity together to support your tightly
// coupled workloads.
//
// Capacity Reservation topology is supported for specific instance types only.
// For more information, see [Prerequisites for Amazon EC2 instance topology]in the Amazon EC2 User Guide.
//
// The Amazon EC2 API follows an eventual consistency model due to the distributed
// nature of the system supporting it. As a result, when you call the
// DescribeCapacityReservationTopology API command immediately after launching
// instances, the response might return a null value for capacityBlockId because
// the data might not have fully propagated across all subsystems. For more
// information, see [Eventual consistency in the Amazon EC2 API]in the Amazon EC2 Developer Guide.
//
// For more information, see [Amazon EC2 topology] in the Amazon EC2 User Guide.
//
// [Prerequisites for Amazon EC2 instance topology]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-topology-prerequisites.html
// [Amazon EC2 topology]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-topology.html
// [Eventual consistency in the Amazon EC2 API]: https://docs.aws.amazon.com/ec2/latest/devguide/eventual-consistency.html
