package kinesisanalyticsv2

// AddApplicationOutput is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisanalyticsv2.go.
//
// Adds an external destination to your SQL-based Kinesis Data Analytics
// application.
//
// If you want Kinesis Data Analytics to deliver data from an in-application
// stream within your application to an external destination (such as an Kinesis
// data stream, a Kinesis Data Firehose delivery stream, or an Amazon Lambda
// function), you add the relevant configuration to your application using this
// operation. You can configure one or more outputs for your application. Each
// output configuration maps an in-application stream and an external destination.
//
// You can use one of the output configurations to deliver data from your
// in-application error stream to an external destination so that you can analyze
// the errors.
//
// Any configuration update, including adding a streaming source using this
// operation, results in a new version of the application. You can use the DescribeApplication
// operation to find the current application version.
