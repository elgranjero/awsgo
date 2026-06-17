package ec2

// DisableFastLaunch is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Discontinue Windows fast launch for a Windows AMI, and clean up existing
// pre-provisioned snapshots. After you disable Windows fast launch, the AMI uses
// the standard launch process for each new instance. Amazon EC2 must remove all
// pre-provisioned snapshots before you can enable Windows fast launch again.
//
// You can only change these settings for Windows AMIs that you own or that have
// been shared with you.
