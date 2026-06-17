package kinesisanalytics

// AddApplicationOutput is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisanalytics.go.
//
// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Adds an external destination to your Amazon Kinesis Analytics application.
//
// If you want Amazon Kinesis Analytics to deliver data from an in-application
// stream within your application to an external destination (such as an Amazon
// Kinesis stream, an Amazon Kinesis Firehose delivery stream, or an AWS Lambda
// function), you add the relevant configuration to your application using this
// operation. You can configure one or more outputs for your application. Each
// output configuration maps an in-application stream and an external destination.
//
// You can use one of the output configurations to deliver data from your
// in-application error stream to an external destination so that you can analyze
// the errors. For more information, see [Understanding Application Output (Destination)].
//
// Any configuration update, including adding a streaming source using this
// operation, results in a new version of the application. You can use the [DescribeApplication]
// operation to find the current application version.
//
// For the limits on the number of application inputs and outputs you can
// configure, see [Limits].
//
// This operation requires permissions to perform the
// kinesisanalytics:AddApplicationOutput action.
//
// [Limits]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/limits.html
// [Understanding Application Output (Destination)]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-output.html
// [DescribeApplication]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html
