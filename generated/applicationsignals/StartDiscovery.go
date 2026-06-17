package applicationsignals

// StartDiscovery is generated as a reference stub.
// Executable command wiring lives under cmd/applicationsignals.go.
//
// Enables this Amazon Web Services account to be able to use CloudWatch
// Application Signals by creating the
// AWSServiceRoleForCloudWatchApplicationSignals service-linked role. This service-
// linked role has the following permissions:
//
// - xray:GetServiceGraph
//
// - logs:StartQuery
//
// - logs:GetQueryResults
//
// - cloudwatch:GetMetricData
//
// - cloudwatch:ListMetrics
//
// - tag:GetResources
//
// - autoscaling:DescribeAutoScalingGroups
//
// A service-linked CloudTrail event channel is created to process CloudTrail
// events and return change event information. This includes last deployment time,
// userName, eventName, and other event metadata.
//
// After completing this step, you still need to instrument your Java and Python
// applications to send data to Application Signals. For more information, see [Enabling Application Signals].
//
// [Enabling Application Signals]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Application-Signals-Enable.html
