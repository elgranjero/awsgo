package cloudformation

// DescribeStackDriftDetectionStatus is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Returns information about a stack drift detection operation. A stack drift
// detection operation detects whether a stack's actual configuration differs, or
// has drifted, from its expected configuration, as defined in the stack template
// and any values specified as template parameters. A stack is considered to have
// drifted if one or more of its resources have drifted. For more information about
// stack and resource drift, see [Detect unmanaged configuration changes to stacks and resources with drift detection].
//
// Use DetectStackDrift to initiate a stack drift detection operation. DetectStackDrift returns a
// StackDriftDetectionId you can use to monitor the progress of the operation using
// DescribeStackDriftDetectionStatus . Once the drift detection operation has
// completed, use DescribeStackResourceDriftsto return drift information about the stack and its resources.
//
// [Detect unmanaged configuration changes to stacks and resources with drift detection]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift.html
