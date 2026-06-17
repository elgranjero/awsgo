package ec2

// GetInstanceTypesFromInstanceRequirements is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Returns a list of instance types with the specified instance attributes. You
// can use the response to preview the instance types without launching instances.
// Note that the response does not consider capacity.
//
// When you specify multiple parameters, you get instance types that satisfy all
// of the specified parameters. If you specify multiple values for a parameter, you
// get instance types that satisfy any of the specified values.
//
// For more information, see [Preview instance types with specified attributes], [Specify attributes for instance type selection for EC2 Fleet or Spot Fleet], and [Spot placement score] in the Amazon EC2 User Guide, and [Creating mixed instance groups using attribute-based instance type selection] in the
// Amazon EC2 Auto Scaling User Guide.
//
// [Specify attributes for instance type selection for EC2 Fleet or Spot Fleet]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet-attribute-based-instance-type-selection.html
// [Preview instance types with specified attributes]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet-attribute-based-instance-type-selection.html#ec2fleet-get-instance-types-from-instance-requirements
// [Creating mixed instance groups using attribute-based instance type selection]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-asg-instance-type-requirements.html
// [Spot placement score]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-placement-score.html
