package databasemigrationservice

// StartReplicationTaskAssessment is generated as a reference stub.
// Executable command wiring lives under cmd/databasemigrationservice.go.
//
// Starts the replication task assessment for unsupported data types in the
//
// source database.
//
// You can only use this operation for a task if the following conditions are true:
//
// - The task must be in the stopped state.
//
// - The task must have successful connections to the source and target.
//
// If either of these conditions are not met, an InvalidResourceStateFault error
// will result.
//
// For information about DMS task assessments, see [Creating a task assessment report] in the Database Migration
// Service User Guide.
//
// [Creating a task assessment report]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.AssessmentReport.html
