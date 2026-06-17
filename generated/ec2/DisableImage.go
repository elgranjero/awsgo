package ec2

// DisableImage is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Sets the AMI state to disabled and removes all launch permissions from the AMI.
// A disabled AMI can't be used for instance launches.
//
// A disabled AMI can't be shared. If an AMI was public or previously shared, it
// is made private. If an AMI was shared with an Amazon Web Services account,
// organization, or Organizational Unit, they lose access to the disabled AMI.
//
// A disabled AMI does not appear in [DescribeImages] API calls by default.
//
// Only the AMI owner can disable an AMI.
//
// You can re-enable a disabled AMI using [EnableImage].
//
// For more information, see [Disable an AMI] in the Amazon EC2 User Guide.
//
// [DescribeImages]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeImages.html
// [Disable an AMI]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/disable-an-ami.html
// [EnableImage]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_EnableImage.html
