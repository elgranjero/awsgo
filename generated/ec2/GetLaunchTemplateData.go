package ec2

// GetLaunchTemplateData is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Retrieves the configuration data of the specified instance. You can use this
// data to create a launch template.
//
// This action calls on other describe actions to get instance information.
// Depending on your instance configuration, you may need to allow the following
// actions in your IAM policy: DescribeSpotInstanceRequests ,
// DescribeInstanceCreditSpecifications , DescribeVolumes , and
// DescribeInstanceAttribute . Or, you can allow describe* depending on your
// instance requirements.
