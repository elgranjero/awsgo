package kinesisanalytics

// StopApplication is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisanalytics.go.
//
// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Stops the application from processing input data. You can stop an application
// only if it is in the running state. You can use the [DescribeApplication]operation to find the
// application state. After the application is stopped, Amazon Kinesis Analytics
// stops reading data from the input, the application stops processing data, and
// there is no output written to the destination.
//
// This operation requires permissions to perform the
// kinesisanalytics:StopApplication action.
//
// [DescribeApplication]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html
