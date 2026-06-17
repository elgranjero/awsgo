package ec2

// ModifyInstanceAttribute is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Modifies the specified attribute of the specified instance. You can specify
// only one attribute at a time.
//
// Note: Using this action to change the security groups associated with an
// elastic network interface (ENI) attached to an instance can result in an error
// if the instance has more than one ENI. To change the security groups associated
// with an ENI attached to an instance that has multiple ENIs, we recommend that
// you use the ModifyNetworkInterfaceAttributeaction.
//
// To modify some attributes, the instance must be stopped. For more information,
// see [Modify a stopped instance]in the Amazon EC2 User Guide.
//
// [Modify a stopped instance]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_ChangingAttributesWhileInstanceStopped.html
