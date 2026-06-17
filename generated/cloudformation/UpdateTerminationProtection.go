package cloudformation

// UpdateTerminationProtection is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Updates termination protection for the specified stack. If a user attempts to
// delete a stack with termination protection enabled, the operation fails and the
// stack remains unchanged. For more information, see [Protect a CloudFormation stack from being deleted]in the CloudFormation User
// Guide.
//
// For [nested stacks], termination protection is set on the root stack and can't be changed
// directly on the nested stack.
//
// [nested stacks]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-nested-stacks.html
// [Protect a CloudFormation stack from being deleted]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-protect-stacks.html
