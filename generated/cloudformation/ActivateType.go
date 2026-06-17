package cloudformation

// ActivateType is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Activates a public third-party extension, such as a resource or module, to make
// it available for use in stack templates in your current account and Region. It
// can also create CloudFormation Hooks, which allow you to evaluate resource
// configurations before CloudFormation provisions them. Hooks integrate with both
// CloudFormation and Cloud Control API operations.
//
// After you activate an extension, you can use [SetTypeConfiguration] to set specific properties for
// the extension.
//
// To see which extensions have been activated, use [ListTypes]. To see configuration details
// for an extension, use [DescribeType].
//
// For more information, see [Activate a third-party public extension in your account] in the CloudFormation User Guide. For information
// about creating Hooks, see the [CloudFormation Hooks User Guide].
//
// [DescribeType]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeType.html
// [SetTypeConfiguration]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_SetTypeConfiguration.html
// [ListTypes]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListTypes.html
// [Activate a third-party public extension in your account]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-public-activate-extension.html
// [CloudFormation Hooks User Guide]: https://docs.aws.amazon.com/cloudformation-cli/latest/hooks-userguide/what-is-cloudformation-hooks.html
