package connect

// CreateTrafficDistributionGroup is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Creates a traffic distribution group given an Amazon Connect instance that has
// been replicated.
//
// The SignInConfig distribution is available only on a default
// TrafficDistributionGroup (see the IsDefault parameter in the [TrafficDistributionGroup] data type). If
// you call UpdateTrafficDistribution with a modified SignInConfig and a
// non-default TrafficDistributionGroup , an InvalidRequestException is returned.
//
// For more information about creating traffic distribution groups, see [Set up traffic distribution groups] in the
// Amazon Connect Administrator Guide.
//
// [Set up traffic distribution groups]: https://docs.aws.amazon.com/connect/latest/adminguide/setup-traffic-distribution-groups.html
// [TrafficDistributionGroup]: https://docs.aws.amazon.com/connect/latest/APIReference/API_TrafficDistributionGroup.html
