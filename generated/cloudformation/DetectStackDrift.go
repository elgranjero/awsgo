package cloudformation

// DetectStackDrift is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Detects whether a stack's actual configuration differs, or has drifted, from
// its expected configuration, as defined in the stack template and any values
// specified as template parameters. For each resource in the stack that supports
// drift detection, CloudFormation compares the actual configuration of the
// resource with its expected template configuration. Only resource properties
// explicitly defined in the stack template are checked for drift. A stack is
// considered to have drifted if one or more of its resources differ from their
// expected template configurations. For more information, see [Detect unmanaged configuration changes to stacks and resources with drift detection].
//
// Use DetectStackDrift to detect drift on all supported resources for a given
// stack, or DetectStackResourceDriftto detect drift on individual resources.
//
// For a list of stack resources that currently support drift detection, see [Resource type support for imports and drift detection].
//
// DetectStackDrift can take up to several minutes, depending on the number of
// resources contained within the stack. Use DescribeStackDriftDetectionStatusto monitor the progress of a detect
// stack drift operation. Once the drift detection operation has completed, use DescribeStackResourceDriftsto
// return drift information about the stack and its resources.
//
// When detecting drift on a stack, CloudFormation doesn't detect drift on any
// nested stacks belonging to that stack. Perform DetectStackDrift directly on the
// nested stack itself.
//
// [Resource type support for imports and drift detection]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/resource-import-supported-resources.html
// [Detect unmanaged configuration changes to stacks and resources with drift detection]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift.html
