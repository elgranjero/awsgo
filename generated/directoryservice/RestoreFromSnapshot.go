package directoryservice

// RestoreFromSnapshot is generated as a reference stub.
// Executable command wiring lives under cmd/directoryservice.go.
//
// Restores a directory using an existing directory snapshot.
//
// When you restore a directory from a snapshot, any changes made to the directory
// after the snapshot date are overwritten.
//
// This action returns as soon as the restore operation is initiated. You can
// monitor the progress of the restore operation by calling the DescribeDirectoriesoperation with the
// directory identifier. When the DirectoryDescription.Stage value changes to
// Active , the restore operation is complete.
