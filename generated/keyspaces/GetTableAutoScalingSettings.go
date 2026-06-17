package keyspaces

// GetTableAutoScalingSettings is generated as a reference stub.
// Executable command wiring lives under cmd/keyspaces.go.
//
// Returns auto scaling related settings of the specified table in JSON format. If
// the table is a multi-Region table, the Amazon Web Services Region specific auto
// scaling settings of the table are included.
//
// Amazon Keyspaces auto scaling helps you provision throughput capacity for
// variable workloads efficiently by increasing and decreasing your table's read
// and write capacity automatically in response to application traffic. For more
// information, see [Managing throughput capacity automatically with Amazon Keyspaces auto scaling]in the Amazon Keyspaces Developer Guide.
//
// GetTableAutoScalingSettings can't be used as an action in an IAM policy.
//
// To define permissions for GetTableAutoScalingSettings , you must allow the
// following two actions in the IAM policy statement's Action element:
//
// - application-autoscaling:DescribeScalableTargets
//
// - application-autoscaling:DescribeScalingPolicies
//
// [Managing throughput capacity automatically with Amazon Keyspaces auto scaling]: https://docs.aws.amazon.com/keyspaces/latest/devguide/autoscaling.html
