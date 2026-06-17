package efs

// DeleteFileSystem is generated as a reference stub.
// Executable command wiring lives under cmd/efs.go.
//
// Deletes a file system, permanently severing access to its contents. Upon
// return, the file system no longer exists and you can't access any contents of
// the deleted file system.
//
// You need to manually delete mount targets attached to a file system before you
// can delete an EFS file system. This step is performed for you when you use the
// Amazon Web Services console to delete a file system.
//
// You cannot delete a file system that is part of an EFS replication
// configuration. You need to delete the replication configuration first.
//
// You can't delete a file system that is in use. That is, if the file system has
// any mount targets, you must first delete them. For more information, see DescribeMountTargetsand DeleteMountTarget.
//
// The DeleteFileSystem call returns while the file system state is still deleting
// . You can check the file system deletion status by calling the DescribeFileSystemsoperation, which
// returns a list of file systems in your account. If you pass file system ID or
// creation token for the deleted file system, the DescribeFileSystemsreturns a 404 FileSystemNotFound
// error.
//
// This operation requires permissions for the elasticfilesystem:DeleteFileSystem
// action.
