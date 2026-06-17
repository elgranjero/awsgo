package efs

// CreateFileSystem is generated as a reference stub.
// Executable command wiring lives under cmd/efs.go.
//
// Creates a new, empty file system. The operation requires a creation token in
// the request that Amazon EFS uses to ensure idempotent creation (calling the
// operation with same creation token has no effect). If a file system does not
// currently exist that is owned by the caller's Amazon Web Services account with
// the specified creation token, this operation does the following:
//
// - Creates a new, empty file system. The file system will have an Amazon EFS
// assigned ID, and an initial lifecycle state creating .
//
// - Returns with the description of the created file system.
//
// Otherwise, this operation returns a FileSystemAlreadyExists error with the ID
// of the existing file system.
//
// For basic use cases, you can use a randomly generated UUID for the creation
// token.
//
// The idempotent operation allows you to retry a CreateFileSystem call without
// risk of creating an extra file system. This can happen when an initial call
// fails in a way that leaves it uncertain whether or not a file system was
// actually created. An example might be that a transport level timeout occurred or
// your connection was reset. As long as you use the same creation token, if the
// initial call had succeeded in creating a file system, the client can learn of
// its existence from the FileSystemAlreadyExists error.
//
// For more information, see [Creating a file system] in the Amazon EFS User Guide.
//
// The CreateFileSystem call returns while the file system's lifecycle state is
// still creating . You can check the file system creation status by calling the DescribeFileSystems
// operation, which among other things returns the file system state.
//
// This operation accepts an optional PerformanceMode parameter that you choose
// for your file system. We recommend generalPurpose PerformanceMode for all file
// systems. The maxIO mode is a previous generation performance type that is
// designed for highly parallelized workloads that can tolerate higher latencies
// than the generalPurpose mode. MaxIO mode is not supported for One Zone file
// systems or file systems that use Elastic throughput.
//
// The PerformanceMode can't be changed after the file system has been created.
// For more information, see [Amazon EFS performance modes].
//
// You can set the throughput mode for the file system using the ThroughputMode
// parameter.
//
// After the file system is fully created, Amazon EFS sets its lifecycle state to
// available , at which point you can create one or more mount targets for the file
// system in your VPC. For more information, see CreateMountTarget. You mount your Amazon EFS file
// system on an EC2 instances in your VPC by using the mount target. For more
// information, see [Amazon EFS: How it Works].
//
// This operation requires permissions for the elasticfilesystem:CreateFileSystem
// action.
//
// File systems can be tagged on creation. If tags are specified in the creation
// action, IAM performs additional authorization on the
// elasticfilesystem:TagResource action to verify if users have permissions to
// create tags. Therefore, you must grant explicit permissions to use the
// elasticfilesystem:TagResource action. For more information, see [Granting permissions to tag resources during creation].
//
// [Creating a file system]: https://docs.aws.amazon.com/efs/latest/ug/creating-using-create-fs.html#creating-using-create-fs-part1
// [Amazon EFS: How it Works]: https://docs.aws.amazon.com/efs/latest/ug/how-it-works.html
// [Amazon EFS performance modes]: https://docs.aws.amazon.com/efs/latest/ug/performance.html#performancemodes.html
// [Granting permissions to tag resources during creation]: https://docs.aws.amazon.com/efs/latest/ug/using-tags-efs.html#supported-iam-actions-tagging.html
