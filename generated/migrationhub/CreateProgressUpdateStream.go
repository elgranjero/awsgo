package migrationhub

// CreateProgressUpdateStream is generated as a reference stub.
// Executable command wiring lives under cmd/migrationhub.go.
//
// Creates a progress update stream which is an AWS resource used for access
// control as well as a namespace for migration task names that is implicitly
// linked to your AWS account. It must uniquely identify the migration tool as it
// is used for all updates made by the tool; however, it does not need to be unique
// for each AWS account because it is scoped to the AWS account.
