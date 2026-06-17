package kinesisanalyticsv2

// StopApplication is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisanalyticsv2.go.
//
// Stops the application from processing data. You can stop an application only if
// it is in the running status, unless you set the Force parameter to true .
//
// You can use the DescribeApplication operation to find the application status.
//
// Managed Service for Apache Flink takes a snapshot when the application is
// stopped, unless Force is set to true .
