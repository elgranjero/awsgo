package appconfig

// CreateExtension is generated as a reference stub.
// Executable command wiring lives under cmd/appconfig.go.
//
// Creates an AppConfig extension. An extension augments your ability to inject
// logic or behavior at different points during the AppConfig workflow of creating
// or deploying a configuration.
//
// You can create your own extensions or use the Amazon Web Services authored
// extensions provided by AppConfig. For an AppConfig extension that uses Lambda,
// you must create a Lambda function to perform any computation and processing
// defined in the extension. If you plan to create custom versions of the Amazon
// Web Services authored notification extensions, you only need to specify an
// Amazon Resource Name (ARN) in the Uri field for the new extension version.
//
// - For a custom EventBridge notification extension, enter the ARN of the
// EventBridge default events in the Uri field.
//
// - For a custom Amazon SNS notification extension, enter the ARN of an Amazon
// SNS topic in the Uri field.
//
// - For a custom Amazon SQS notification extension, enter the ARN of an Amazon
// SQS message queue in the Uri field.
//
// For more information about extensions, see [Extending workflows] in the AppConfig User Guide.
//
// [Extending workflows]: https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html
