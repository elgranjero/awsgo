package ec2

// DescribeAccountAttributes is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Describes attributes of your Amazon Web Services account. The following are the
// supported account attributes:
//
// - default-vpc : The ID of the default VPC for your account, or none .
//
// - max-instances : This attribute is no longer supported. The returned value
// does not reflect your actual vCPU limit for running On-Demand Instances. For
// more information, see [On-Demand Instance Limits]in the Amazon Elastic Compute Cloud User Guide.
//
// - max-elastic-ips : The maximum number of Elastic IP addresses that you can
// allocate.
//
// - supported-platforms : This attribute is deprecated.
//
// - vpc-max-elastic-ips : The maximum number of Elastic IP addresses that you
// can allocate.
//
// - vpc-max-security-groups-per-interface : The maximum number of security
// groups that you can assign to a network interface.
//
// The order of the elements in the response, including those within nested
// structures, might vary. Applications should not assume the elements appear in a
// particular order.
//
// [On-Demand Instance Limits]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-on-demand-instances.html#ec2-on-demand-instances-limits
