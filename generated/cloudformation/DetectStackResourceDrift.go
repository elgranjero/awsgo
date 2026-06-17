package cloudformation

// DetectStackResourceDrift is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Returns information about whether a resource's actual configuration differs, or
// has drifted, from its expected configuration, as defined in the stack template
// and any values specified as template parameters. This information includes
// actual and expected property values for resources in which CloudFormation
// detects drift. Only resource properties explicitly defined in the stack template
// are checked for drift. For more information about stack and resource drift, see [Detect unmanaged configuration changes to stacks and resources with drift detection]
// .
//
// Use DetectStackResourceDrift to detect drift on individual resources, or DetectStackDrift to
// detect drift on all resources in a given stack that support drift detection.
//
// Resources that don't currently support drift detection can't be checked. For a
// list of resources that support drift detection, see [Resource type support for imports and drift detection].
//
// [Resource type support for imports and drift detection]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/resource-import-supported-resources.html
// [Detect unmanaged configuration changes to stacks and resources with drift detection]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift.html
