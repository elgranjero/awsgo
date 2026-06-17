package backup

// CreateLegalHold is generated as a reference stub.
// Executable command wiring lives under cmd/backup.go.
//
// Creates a legal hold on a recovery point (backup). A legal hold is a restraint
// on altering or deleting a backup until an authorized user cancels the legal
// hold. Any actions to delete or disassociate a recovery point will fail with an
// error if one or more active legal holds are on the recovery point.
