package kinesisanalyticsv2

// UpdateApplicationMaintenanceConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisanalyticsv2.go.
//
// Updates the maintenance configuration of the Managed Service for Apache Flink
// application.
//
// You can invoke this operation on an application that is in one of the two
// following states: READY or RUNNING . If you invoke it when the application is in
// a state other than these two states, it throws a ResourceInUseException . The
// service makes use of the updated configuration the next time it schedules
// maintenance for the application. If you invoke this operation after the service
// schedules maintenance, the service will apply the configuration update the next
// time it schedules maintenance for the application. This means that you might not
// see the maintenance configuration update applied to the maintenance process that
// follows a successful invocation of this operation, but to the following
// maintenance process instead.
//
// To see the current maintenance configuration of your application, invoke the DescribeApplication
// operation.
//
// For information about application maintenance, see [Managed Service for Apache Flink for Apache Flink Maintenance].
//
// This operation is supported only for Managed Service for Apache Flink.
//
// [Managed Service for Apache Flink for Apache Flink Maintenance]: https://docs.aws.amazon.com/kinesisanalytics/latest/java/maintenance.html
