package kinesisanalytics

// StartApplication is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisanalytics.go.
//
// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Starts the specified Amazon Kinesis Analytics application. After creating an
// application, you must exclusively call this operation to start your application.
//
// After the application starts, it begins consuming the input data, processes it,
// and writes the output to the configured destination.
//
// The application status must be READY for you to start an application. You can
// get the application status in the console or using the [DescribeApplication]operation.
//
// After you start the application, you can stop the application from processing
// the input by calling the [StopApplication]operation.
//
// This operation requires permissions to perform the
// kinesisanalytics:StartApplication action.
//
// [DescribeApplication]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html
// [StopApplication]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_StopApplication.html
