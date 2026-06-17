package appconfig

// CreateConfigurationProfile is generated as a reference stub.
// Executable command wiring lives under cmd/appconfig.go.
//
// Creates a configuration profile, which is information that enables AppConfig to
// access the configuration source. Valid configuration sources include the
// following:
//
// - Configuration data in YAML, JSON, and other formats stored in the AppConfig
// hosted configuration store
//
// - Configuration data stored as objects in an Amazon Simple Storage Service
// (Amazon S3) bucket
//
// - Pipelines stored in CodePipeline
//
// - Secrets stored in Secrets Manager
//
// - Standard and secure string parameters stored in Amazon Web Services Systems
// Manager Parameter Store
//
// - Configuration data in SSM documents stored in the Systems Manager document
// store
//
// A configuration profile includes the following information:
//
// - The URI location of the configuration data.
//
// - The Identity and Access Management (IAM) role that provides access to the
// configuration data.
//
// - A validator for the configuration data. Available validators include either
// a JSON Schema or an Amazon Web Services Lambda function.
//
// For more information, see [Create a Configuration and a Configuration Profile] in the AppConfig User Guide.
//
// [Create a Configuration and a Configuration Profile]: http://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-configuration-and-profile.html
