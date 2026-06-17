package neptune

// DeleteDBCluster is generated as a reference stub.
// Executable command wiring lives under cmd/neptune.go.
//
// The DeleteDBCluster action deletes a previously provisioned DB cluster. When
// you delete a DB cluster, all automated backups for that DB cluster are deleted
// and can't be recovered. Manual DB cluster snapshots of the specified DB cluster
// are not deleted.
//
// Note that the DB Cluster cannot be deleted if deletion protection is enabled.
// To delete it, you must first set its DeletionProtection field to False .
