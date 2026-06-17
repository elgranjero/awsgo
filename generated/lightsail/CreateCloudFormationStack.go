package lightsail

// CreateCloudFormationStack is generated as a reference stub.
// Executable command wiring lives under cmd/lightsail.go.
//
// Creates an AWS CloudFormation stack, which creates a new Amazon EC2 instance
// from an exported Amazon Lightsail snapshot. This operation results in a
// CloudFormation stack record that can be used to track the AWS CloudFormation
// stack created. Use the get cloud formation stack records operation to get a
// list of the CloudFormation stacks created.
//
// Wait until after your new Amazon EC2 instance is created before running the
// create cloud formation stack operation again with the same export snapshot
// record.
