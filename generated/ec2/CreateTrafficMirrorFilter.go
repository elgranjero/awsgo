package ec2

// CreateTrafficMirrorFilter is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates a Traffic Mirror filter.
//
// A Traffic Mirror filter is a set of rules that defines the traffic to mirror.
//
// By default, no traffic is mirrored. To mirror traffic, use [CreateTrafficMirrorFilterRule] to add Traffic
// Mirror rules to the filter. The rules you add define what traffic gets mirrored.
// You can also use [ModifyTrafficMirrorFilterNetworkServices]to mirror supported network services.
//
// [CreateTrafficMirrorFilterRule]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTrafficMirrorFilterRule.htm
// [ModifyTrafficMirrorFilterNetworkServices]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyTrafficMirrorFilterNetworkServices.html
