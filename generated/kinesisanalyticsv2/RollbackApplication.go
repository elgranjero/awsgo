package kinesisanalyticsv2

// RollbackApplication is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisanalyticsv2.go.
//
// Reverts the application to the previous running version. You can roll back an
// application if you suspect it is stuck in a transient status or in the running
// status.
//
// You can roll back an application only if it is in the UPDATING , AUTOSCALING ,
// or RUNNING statuses.
//
// When you rollback an application, it loads state data from the last successful
// snapshot. If the application has no snapshots, Managed Service for Apache Flink
// rejects the rollback request.
