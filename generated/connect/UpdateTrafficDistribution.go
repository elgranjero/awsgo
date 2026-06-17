package connect

// UpdateTrafficDistribution is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Updates the traffic distribution for a given traffic distribution group.
//
// When you shift telephony traffic, also shift agents and/or agent sign-ins to
// ensure they can handle the calls in the other Region. If you don't shift the
// agents, voice calls will go to the shifted Region but there won't be any agents
// available to receive the calls.
//
// The SignInConfig distribution is available only on a default
// TrafficDistributionGroup (see the IsDefault parameter in the [TrafficDistributionGroup] data type). If
// you call UpdateTrafficDistribution with a modified SignInConfig and a
// non-default TrafficDistributionGroup , an InvalidRequestException is returned.
//
// For more information about updating a traffic distribution group, see [Update telephony traffic distribution across Amazon Web Services Regions] in the
// Amazon Connect Administrator Guide.
//
// [TrafficDistributionGroup]: https://docs.aws.amazon.com/connect/latest/APIReference/API_TrafficDistributionGroup.html
// [Update telephony traffic distribution across Amazon Web Services Regions]: https://docs.aws.amazon.com/connect/latest/adminguide/update-telephony-traffic-distribution.html
