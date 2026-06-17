package efs

// CreateAccessPoint is generated as a reference stub.
// Executable command wiring lives under cmd/efs.go.
//
// Creates an EFS access point. An access point is an application-specific view
// into an EFS file system that applies an operating system user and group, and a
// file system path, to any file system request made through the access point. The
// operating system user and group override any identity information provided by
// the NFS client. The file system path is exposed as the access point's root
// directory. Applications using the access point can only access data in the
// application's own directory and any subdirectories. A file system can have a
// maximum of 10,000 access points unless you request an increase. To learn more,
// see [Mounting a file system using EFS access points].
//
// If multiple requests to create access points on the same file system are sent
// in quick succession, and the file system is near the limit of access points, you
// may experience a throttling response for these requests. This is to ensure that
// the file system does not exceed the stated access point limit.
//
// This operation requires permissions for the elasticfilesystem:CreateAccessPoint
// action.
//
// Access points can be tagged on creation. If tags are specified in the creation
// action, IAM performs additional authorization on the
// elasticfilesystem:TagResource action to verify if users have permissions to
// create tags. Therefore, you must grant explicit permissions to use the
// elasticfilesystem:TagResource action. For more information, see [Granting permissions to tag resources during creation].
//
// [Mounting a file system using EFS access points]: https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html
// [Granting permissions to tag resources during creation]: https://docs.aws.amazon.com/efs/latest/ug/using-tags-efs.html#supported-iam-actions-tagging.html
