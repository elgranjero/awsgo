package ec2

// EnableImage is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Re-enables a disabled AMI. The re-enabled AMI is marked as available and can be
// used for instance launches, appears in describe operations, and can be shared.
// Amazon Web Services accounts, organizations, and Organizational Units that lost
// access to the AMI when it was disabled do not regain access automatically. Once
// the AMI is available, it can be shared with them again.
//
// Only the AMI owner can re-enable a disabled AMI.
//
// For more information, see [Disable an Amazon EC2 AMI] in the Amazon EC2 User Guide.
//
// [Disable an Amazon EC2 AMI]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/disable-an-ami.html
