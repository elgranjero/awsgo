package ec2

// ModifyTrafficMirrorFilterNetworkServices is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Allows or restricts mirroring network services.
//
// By default, Amazon DNS network services are not eligible for Traffic Mirror.
// Use AddNetworkServices to add network services to a Traffic Mirror filter. When
// a network service is added to the Traffic Mirror filter, all traffic related to
// that network service will be mirrored. When you no longer want to mirror network
// services, use RemoveNetworkServices to remove the network services from the
// Traffic Mirror filter.
