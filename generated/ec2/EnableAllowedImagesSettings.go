package ec2

// EnableAllowedImagesSettings is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Enables Allowed AMIs for your account in the specified Amazon Web Services
// Region. Two values are accepted:
//
// - enabled : The image criteria in your Allowed AMIs settings are applied. As a
// result, only AMIs matching these criteria are discoverable and can be used by
// your account to launch instances.
//
// - audit-mode : The image criteria in your Allowed AMIs settings are not
// applied. No restrictions are placed on AMI discoverability or usage. Users in
// your account can launch instances using any public AMI or AMI shared with your
// account.
//
// The purpose of audit-mode is to indicate which AMIs will be affected when
//
// Allowed AMIs is enabled . In audit-mode , each AMI displays either
// "ImageAllowed": true or "ImageAllowed": false to indicate whether the AMI will
// be discoverable and available to users in the account when Allowed AMIs is
// enabled.
//
// The Allowed AMIs feature does not restrict the AMIs owned by your account.
// Regardless of the criteria you set, the AMIs created by your account will always
// be discoverable and usable by users in your account.
//
// For more information, see [Control the discovery and use of AMIs in Amazon EC2 with Allowed AMIs] in Amazon EC2 User Guide.
//
// [Control the discovery and use of AMIs in Amazon EC2 with Allowed AMIs]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-allowed-amis.html
