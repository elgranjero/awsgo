package ec2

// CreateVpcEndpointServiceConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates a VPC endpoint service to which service consumers (Amazon Web Services
// accounts, users, and IAM roles) can connect.
//
// Before you create an endpoint service, you must create one of the following for
// your service:
//
// - A [Network Load Balancer]. Service consumers connect to your service using an interface endpoint.
//
// - A [Gateway Load Balancer]. Service consumers connect to your service using a Gateway Load Balancer
// endpoint.
//
// If you set the private DNS name, you must prove that you own the private DNS
// domain name.
//
// For more information, see the [Amazon Web Services PrivateLink Guide].
//
// [Gateway Load Balancer]: https://docs.aws.amazon.com/elasticloadbalancing/latest/gateway/
// [Network Load Balancer]: https://docs.aws.amazon.com/elasticloadbalancing/latest/network/
// [Amazon Web Services PrivateLink Guide]: https://docs.aws.amazon.com/vpc/latest/privatelink/
