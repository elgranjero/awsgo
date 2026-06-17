package kinesisanalytics

// AddApplicationInput is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisanalytics.go.
//
// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Adds a streaming source to your Amazon Kinesis application. For conceptual
// information, see [Configuring Application Input].
//
// You can add a streaming source either when you create an application or you can
// use this operation to add a streaming source after you create an application.
// For more information, see [CreateApplication].
//
// Any configuration update, including adding a streaming source using this
// operation, results in a new version of the application. You can use the [DescribeApplication]
// operation to find the current application version.
//
// This operation requires permissions to perform the
// kinesisanalytics:AddApplicationInput action.
//
// [CreateApplication]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_CreateApplication.html
// [DescribeApplication]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html
// [Configuring Application Input]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-input.html
