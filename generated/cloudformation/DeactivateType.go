package cloudformation

// DeactivateType is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Deactivates a public third-party extension, such as a resource or module, or a
// CloudFormation Hook when you no longer use it.
//
// Deactivating an extension deletes the configuration details that are associated
// with it. To temporarily disable a CloudFormation Hook instead, you can use [SetTypeConfiguration].
//
// Once deactivated, an extension can't be used in any CloudFormation operation.
// This includes stack update operations where the stack template includes the
// extension, even if no updates are being made to the extension. In addition,
// deactivated extensions aren't automatically updated if a new version of the
// extension is released.
//
// To see which extensions are currently activated, use [ListTypes].
//
// [SetTypeConfiguration]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_SetTypeConfiguration.html
// [ListTypes]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListTypes.html
