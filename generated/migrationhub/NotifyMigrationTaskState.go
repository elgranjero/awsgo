package migrationhub

// NotifyMigrationTaskState is generated as a reference stub.
// Executable command wiring lives under cmd/migrationhub.go.
//
// Notifies Migration Hub of the current status, progress, or other detail
// regarding a migration task. This API has the following traits:
//
// - Migration tools will call the NotifyMigrationTaskState API to share the
// latest progress and status.
//
// - MigrationTaskName is used for addressing updates to the correct target.
//
// - ProgressUpdateStream is used for access control and to provide a namespace
// for each migration tool.
