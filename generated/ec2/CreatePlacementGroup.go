package ec2

// CreatePlacementGroup is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates a placement group in which to launch instances. The strategy of the
// placement group determines how the instances are organized within the group.
//
// A cluster placement group is a logical grouping of instances within a single
// Availability Zone that benefit from low network latency, high network
// throughput. A spread placement group places instances on distinct hardware. A
// partition placement group places groups of instances in different partitions,
// where instances in one partition do not share the same hardware with instances
// in another partition.
//
// For more information, see [Placement groups] in the Amazon EC2 User Guide.
//
// [Placement groups]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html
