package ec2

// DescribeVpcEndpointServices is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Describes available services to which you can create a VPC endpoint.
//
// When the service provider and the consumer have different accounts in multiple
// Availability Zones, and the consumer views the VPC endpoint service information,
// the response only includes the common Availability Zones. For example, when the
// service provider account uses us-east-1a and us-east-1c and the consumer uses
// us-east-1a and us-east-1b , the response includes the VPC endpoint services in
// the common Availability Zone, us-east-1a .
