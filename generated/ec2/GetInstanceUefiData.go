package ec2

// GetInstanceUefiData is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// A binary representation of the UEFI variable store. Only non-volatile variables
// are stored. This is a base64 encoded and zlib compressed binary value that must
// be properly encoded.
//
// When you use [register-image] to create an AMI, you can create an exact copy of your variable
// store by passing the UEFI data in the UefiData parameter. You can modify the
// UEFI data by using the [python-uefivars tool]on GitHub. You can use the tool to convert the UEFI data
// into a human-readable format (JSON), which you can inspect and modify, and then
// convert back into the binary format to use with register-image.
//
// For more information, see [UEFI Secure Boot] in the Amazon EC2 User Guide.
//
// [UEFI Secure Boot]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/uefi-secure-boot.html
// [python-uefivars tool]: https://github.com/awslabs/python-uefivars
// [register-image]: https://docs.aws.amazon.com/cli/latest/reference/ec2/register-image.html
