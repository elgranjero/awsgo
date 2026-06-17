package databasemigrationservice

// ModifyReplicationConfig is generated as a reference stub.
// Executable command wiring lives under cmd/databasemigrationservice.go.
//
// Modifies an existing DMS Serverless replication configuration that you can use
// to start a replication. This command includes input validation and logic to
// check the state of any replication that uses this configuration. You can only
// modify a replication configuration before any replication that uses it has
// started. As soon as you have initially started a replication with a given
// configuiration, you can't modify that configuration, even if you stop it.
//
// Other run statuses that allow you to run this command include FAILED and
// CREATED. A provisioning state that allows you to run this command is
// FAILED_PROVISION.
