package neptune

// CreateDBCluster is generated as a reference stub.
// Executable command wiring lives under cmd/neptune.go.
//
// Creates a new Amazon Neptune DB cluster.
//
// You can use the ReplicationSourceIdentifier parameter to create the DB cluster
// as a Read Replica of another DB cluster or Amazon Neptune DB instance.
//
// Note that when you create a new cluster using CreateDBCluster directly,
// deletion protection is disabled by default (when you create a new production
// cluster in the console, deletion protection is enabled by default). You can only
// delete a DB cluster if its DeletionProtection field is set to false .
