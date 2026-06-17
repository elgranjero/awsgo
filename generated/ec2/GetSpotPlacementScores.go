package ec2

// GetSpotPlacementScores is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Calculates the Spot placement score for a Region or Availability Zone based on
// the specified target capacity and compute requirements.
//
// You can specify your compute requirements either by using
// InstanceRequirementsWithMetadata and letting Amazon EC2 choose the optimal
// instance types to fulfill your Spot request, or you can specify the instance
// types by using InstanceTypes .
//
// For more information, see [Spot placement score] in the Amazon EC2 User Guide.
//
// [Spot placement score]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-placement-score.html
