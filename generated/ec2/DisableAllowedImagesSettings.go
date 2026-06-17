package ec2

// DisableAllowedImagesSettings is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Disables Allowed AMIs for your account in the specified Amazon Web Services
// Region. When set to disabled , the image criteria in your Allowed AMIs settings
// do not apply, and no restrictions are placed on AMI discoverability or usage.
// Users in your account can launch instances using any public AMI or AMI shared
// with your account.
//
// The Allowed AMIs feature does not restrict the AMIs owned by your account.
// Regardless of the criteria you set, the AMIs created by your account will always
// be discoverable and usable by users in your account.
//
// For more information, see [Control the discovery and use of AMIs in Amazon EC2 with Allowed AMIs] in Amazon EC2 User Guide.
//
// [Control the discovery and use of AMIs in Amazon EC2 with Allowed AMIs]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-allowed-amis.html
