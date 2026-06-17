package cloudformation

// RegisterType is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Registers an extension with the CloudFormation service. Registering an
// extension makes it available for use in CloudFormation templates in your Amazon
// Web Services account, and includes:
//
// - Validating the extension schema.
//
// - Determining which handlers, if any, have been specified for the extension.
//
// - Making the extension available for use in your account.
//
// For more information about how to develop extensions and ready them for
// registration, see [Creating resource types using the CloudFormation CLI]in the CloudFormation Command Line Interface (CLI) User Guide.
//
// You can have a maximum of 50 resource extension versions registered at a time.
// This maximum is per account and per Region. Use [DeregisterType]to deregister specific
// extension versions if necessary.
//
// Once you have initiated a registration request using RegisterType, you can use DescribeTypeRegistration to monitor
// the progress of the registration request.
//
// Once you have registered a private extension in your account and Region, use [SetTypeConfiguration]
// to specify configuration properties for the extension. For more information, see
// [Edit configuration data for extensions in your account]in the CloudFormation User Guide.
//
// [SetTypeConfiguration]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_SetTypeConfiguration.html
// [Edit configuration data for extensions in your account]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-set-configuration.html
// [Creating resource types using the CloudFormation CLI]: https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-types.html
// [DeregisterType]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DeregisterType.html
