package kinesisanalytics

// DiscoverInputSchema is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisanalytics.go.
//
// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Infers a schema by evaluating sample records on the specified streaming source
// (Amazon Kinesis stream or Amazon Kinesis Firehose delivery stream) or S3 object.
// In the response, the operation returns the inferred schema and also the sample
// records that the operation used to infer the schema.
//
// You can use the inferred schema when configuring a streaming source for your
// application. For conceptual information, see [Configuring Application Input]. Note that when you create an
// application using the Amazon Kinesis Analytics console, the console uses this
// operation to infer a schema and show it in the console user interface.
//
// This operation requires permissions to perform the
// kinesisanalytics:DiscoverInputSchema action.
//
// [Configuring Application Input]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-input.html
