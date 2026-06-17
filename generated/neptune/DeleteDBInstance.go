package neptune

// DeleteDBInstance is generated as a reference stub.
// Executable command wiring lives under cmd/neptune.go.
//
// The DeleteDBInstance action deletes a previously provisioned DB instance. When
// you delete a DB instance, all automated backups for that instance are deleted
// and can't be recovered. Manual DB snapshots of the DB instance to be deleted by
// DeleteDBInstance are not deleted.
//
// If you request a final DB snapshot the status of the Amazon Neptune DB instance
// is deleting until the DB snapshot is created. The API action DescribeDBInstance
// is used to monitor the status of this operation. The action can't be canceled or
// reverted once submitted.
//
// Note that when a DB instance is in a failure state and has a status of failed ,
// incompatible-restore , or incompatible-network , you can only delete it when the
// SkipFinalSnapshot parameter is set to true .
//
// You can't delete a DB instance if it is the only instance in the DB cluster, or
// if it has deletion protection enabled.
