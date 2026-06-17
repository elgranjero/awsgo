package cloudformation

// DetectStackSetDrift is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Detect drift on a StackSet. When CloudFormation performs drift detection on a
// StackSet, it performs drift detection on the stack associated with each stack
// instance in the StackSet. For more information, see [Performing drift detection on CloudFormation StackSets].
//
// DetectStackSetDrift returns the OperationId of the StackSet drift detection
// operation. Use this operation id with DescribeStackSetOperationto monitor the progress of the drift
// detection operation. The drift detection operation may take some time, depending
// on the number of stack instances included in the StackSet, in addition to the
// number of resources included in each stack.
//
// Once the operation has completed, use the following actions to return drift
// information:
//
// - Use DescribeStackSetto return detailed information about the stack set, including detailed
// information about the last completed drift operation performed on the StackSet.
// (Information about drift operations that are in progress isn't included.)
//
// - Use ListStackInstancesto return a list of stack instances belonging to the StackSet,
// including the drift status and last drift time checked of each instance.
//
// - Use DescribeStackInstanceto return detailed information about a specific stack instance,
// including its drift status and last drift time checked.
//
// You can only run a single drift detection operation on a given StackSet at one
// time.
//
// To stop a drift detection StackSet operation, use StopStackSetOperation.
//
// [Performing drift detection on CloudFormation StackSets]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-drift.html
