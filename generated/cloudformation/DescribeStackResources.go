package cloudformation

// DescribeStackResources is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Returns Amazon Web Services resource descriptions for running and deleted
// stacks. If StackName is specified, all the associated resources that are part
// of the stack are returned. If PhysicalResourceId is specified, the associated
// resources of the stack that the resource belongs to are returned.
//
// Only the first 100 resources will be returned. If your stack has more resources
// than this, you should use ListStackResources instead.
//
// For deleted stacks, DescribeStackResources returns resource information for up
// to 90 days after the stack has been deleted.
//
// You must specify either StackName or PhysicalResourceId , but not both. In
// addition, you can specify LogicalResourceId to filter the returned result. For
// more information about resources, the LogicalResourceId and PhysicalResourceId ,
// see the [CloudFormation User Guide].
//
// A ValidationError is returned if you specify both StackName and
// PhysicalResourceId in the same request.
//
// [CloudFormation User Guide]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/
