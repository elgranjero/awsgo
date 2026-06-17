package kinesisanalytics

// ListApplications is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisanalytics.go.
//
// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Returns a list of Amazon Kinesis Analytics applications in your account. For
// each application, the response includes the application name, Amazon Resource
// Name (ARN), and status.
//
// If the response returns the HasMoreApplications value as true,
//
// you can send another request by adding the ExclusiveStartApplicationName in the
// request body, and set the value of this to the last application name from the
// previous response.
//
// If you want detailed information about a specific application, use [DescribeApplication].
//
// This operation requires permissions to perform the
// kinesisanalytics:ListApplications action.
//
// [DescribeApplication]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html
