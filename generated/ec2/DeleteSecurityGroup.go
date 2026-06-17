package ec2

// DeleteSecurityGroup is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Deletes a security group.
//
// If you attempt to delete a security group that is associated with an instance
// or network interface, is referenced by another security group in the same VPC,
// or has a VPC association, the operation fails with DependencyViolation .
