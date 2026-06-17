package cloudformation

// DescribeStackResourceDrifts is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Returns drift information for the resources that have been checked for drift in
// the specified stack. This includes actual and expected configuration values for
// resources where CloudFormation detects configuration drift.
//
// For a given stack, there will be one StackResourceDrift for each stack resource
// that has been checked for drift. Resources that haven't yet been checked for
// drift aren't included. Resources that don't currently support drift detection
// aren't checked, and so not included. For a list of resources that support drift
// detection, see [Resource type support for imports and drift detection].
//
// Use DetectStackResourceDrift to detect drift on individual resources, or DetectStackDrift to detect drift on all
// supported resources for a given stack.
//
// [Resource type support for imports and drift detection]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/resource-import-supported-resources.html
