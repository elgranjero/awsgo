package ec2

// CreateTrafficMirrorTarget is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates a target for your Traffic Mirror session.
//
// A Traffic Mirror target is the destination for mirrored traffic. The Traffic
// Mirror source and the Traffic Mirror target (monitoring appliances) can be in
// the same VPC, or in different VPCs connected via VPC peering or a transit
// gateway.
//
// A Traffic Mirror target can be a network interface, a Network Load Balancer, or
// a Gateway Load Balancer endpoint.
//
// To use the target in a Traffic Mirror session, use [CreateTrafficMirrorSession].
//
// [CreateTrafficMirrorSession]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTrafficMirrorSession.htm
