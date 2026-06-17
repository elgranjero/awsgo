package efs

// PutFileSystemPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/efs.go.
//
// Applies an Amazon EFS FileSystemPolicy to an Amazon EFS file system. A file
// system policy is an IAM resource-based policy and can contain multiple policy
// statements. A file system always has exactly one file system policy, which can
// be the default policy or an explicit policy set or updated using this API
// operation. EFS file system policies have a 20,000 character limit. When an
// explicit policy is set, it overrides the default policy. For more information
// about the default file system policy, see [Default EFS file system policy].
//
// EFS file system policies have a 20,000 character limit.
//
// This operation requires permissions for the
// elasticfilesystem:PutFileSystemPolicy action.
//
// [Default EFS file system policy]: https://docs.aws.amazon.com/efs/latest/ug/iam-access-control-nfs-efs.html#default-filesystempolicy
