package cloudformation

// TestType is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Tests a registered extension to make sure it meets all necessary requirements
// for being published in the CloudFormation registry.
//
// - For resource types, this includes passing all contracts tests defined for
// the type.
//
// - For modules, this includes determining if the module's model meets all
// necessary requirements.
//
// For more information, see [Testing your public extension before publishing] in the CloudFormation Command Line Interface (CLI)
// User Guide.
//
// If you don't specify a version, CloudFormation uses the default version of the
// extension in your account and Region for testing.
//
// To perform testing, CloudFormation assumes the execution role specified when
// the type was registered. For more information, see [RegisterType].
//
// Once you've initiated testing on an extension using TestType , you can pass the
// returned TypeVersionArn into [DescribeType] to monitor the current test status and test
// status description for the extension.
//
// An extension must have a test status of PASSED before it can be published. For
// more information, see [Publishing extensions to make them available for public use]in the CloudFormation Command Line Interface (CLI) User
// Guide.
//
// [DescribeType]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeType.html
// [Testing your public extension before publishing]: https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/publish-extension.html#publish-extension-testing
// [RegisterType]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_RegisterType.html
// [Publishing extensions to make them available for public use]: https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-publish.html
