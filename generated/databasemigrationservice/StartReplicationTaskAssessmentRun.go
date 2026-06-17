package databasemigrationservice

// StartReplicationTaskAssessmentRun is generated as a reference stub.
// Executable command wiring lives under cmd/databasemigrationservice.go.
//
// Starts a new premigration assessment run for one or more individual assessments
// of a migration task.
//
// The assessments that you can specify depend on the source and target database
// engine and the migration type defined for the given task. To run this operation,
// your migration task must already be created. After you run this operation, you
// can review the status of each individual assessment. You can also run the
// migration task manually after the assessment run and its individual assessments
// complete.
