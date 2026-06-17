package kinesisanalytics

// AddApplicationReferenceDataSource is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisanalytics.go.
//
// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Adds a reference data source to an existing application.
//
// Amazon Kinesis Analytics reads reference data (that is, an Amazon S3 object)
// and creates an in-application table within your application. In the request, you
// provide the source (S3 bucket name and object key name), name of the
// in-application table to create, and the necessary mapping information that
// describes how data in Amazon S3 object maps to columns in the resulting
// in-application table.
//
// For conceptual information, see [Configuring Application Input]. For the limits on data sources you can add to
// your application, see [Limits].
//
// This operation requires permissions to perform the
// kinesisanalytics:AddApplicationOutput action.
//
// [Limits]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/limits.html
// [Configuring Application Input]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-input.html
