package ec2

// CreateTrafficMirrorSession is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates a Traffic Mirror session.
//
// A Traffic Mirror session actively copies packets from a Traffic Mirror source
// to a Traffic Mirror target. Create a filter, and then assign it to the session
// to define a subset of the traffic to mirror, for example all TCP traffic.
//
// The Traffic Mirror source and the Traffic Mirror target (monitoring appliances)
// can be in the same VPC, or in a different VPC connected via VPC peering or a
// transit gateway.
//
// By default, no traffic is mirrored. Use [CreateTrafficMirrorFilter] to create filter rules that specify
// the traffic to mirror.
//
// [CreateTrafficMirrorFilter]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTrafficMirrorFilter.html
