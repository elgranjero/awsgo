package ec2

// DescribeFleetHistory is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Describes the events for the specified EC2 Fleet during the specified time.
//
// EC2 Fleet events are delayed by up to 30 seconds before they can be described.
// This ensures that you can query by the last evaluated time and not miss a
// recorded event. EC2 Fleet events are available for 48 hours.
//
// For more information, see [Monitor fleet events using Amazon EventBridge] in the Amazon EC2 User Guide.
//
// [Monitor fleet events using Amazon EventBridge]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/fleet-monitor.html
