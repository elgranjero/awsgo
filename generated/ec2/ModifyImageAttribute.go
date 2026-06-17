package ec2

// ModifyImageAttribute is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Modifies the specified attribute of the specified AMI. You can specify only one
// attribute at a time.
//
// To specify the attribute, you can use the Attribute parameter, or one of the
// following parameters: Description , ImdsSupport , or LaunchPermission .
//
// Images with an Amazon Web Services Marketplace product code cannot be made
// public.
//
// To enable the SriovNetSupport enhanced networking attribute of an image, enable
// SriovNetSupport on an instance and create an AMI from the instance.
