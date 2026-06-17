package cloudformation

// DeregisterType is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Marks an extension or extension version as DEPRECATED in the CloudFormation
// registry, removing it from active use. Deprecated extensions or extension
// versions cannot be used in CloudFormation operations.
//
// To deregister an entire extension, you must individually deregister all active
// versions of that extension. If an extension has only a single active version,
// deregistering that version results in the extension itself being deregistered
// and marked as deprecated in the registry.
//
// You can't deregister the default version of an extension if there are other
// active version of that extension. If you do deregister the default version of an
// extension, the extension type itself is deregistered as well and marked as
// deprecated.
//
// To view the deprecation status of an extension or extension version, use [DescribeType].
//
// For more information, see [Remove third-party private extensions from your account] in the CloudFormation User Guide.
//
// [DescribeType]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeType.html
// [Remove third-party private extensions from your account]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-private-deregister-extension.html
