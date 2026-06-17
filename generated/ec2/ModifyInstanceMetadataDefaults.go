package ec2

// ModifyInstanceMetadataDefaults is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Modifies the default instance metadata service (IMDS) settings at the account
// level in the specified Amazon Web Services  Region.
//
// To remove a parameter's account-level default setting, specify no-preference .
// If an account-level setting is cleared with no-preference , then the instance
// launch considers the other instance metadata settings. For more information, see
// [Order of precedence for instance metadata options]in the Amazon EC2 User Guide.
//
// [Order of precedence for instance metadata options]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-instance-metadata-options.html#instance-metadata-options-order-of-precedence
