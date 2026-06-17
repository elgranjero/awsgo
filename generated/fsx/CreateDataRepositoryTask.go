package fsx

// CreateDataRepositoryTask is generated as a reference stub.
// Executable command wiring lives under cmd/fsx.go.
//
// Creates an Amazon FSx for Lustre data repository task. A
// CreateDataRepositoryTask operation will fail if a data repository is not linked
// to the FSx file system.
//
// You use import and export data repository tasks to perform bulk operations
// between your FSx for Lustre file system and its linked data repositories. An
// example of a data repository task is exporting any data and metadata changes,
// including POSIX metadata, to files, directories, and symbolic links (symlinks)
// from your FSx file system to a linked data repository.
//
// You use release data repository tasks to release data from your file system for
// files that are exported to S3. The metadata of released files remains on the
// file system so users or applications can still access released files by reading
// the files again, which will restore data from Amazon S3 to the FSx for Lustre
// file system.
//
// To learn more about data repository tasks, see [Data Repository Tasks]. To learn more about linking a
// data repository to your file system, see [Linking your file system to an S3 bucket].
//
// [Data Repository Tasks]: https://docs.aws.amazon.com/fsx/latest/LustreGuide/data-repository-tasks.html
// [Linking your file system to an S3 bucket]: https://docs.aws.amazon.com/fsx/latest/LustreGuide/create-dra-linked-data-repo.html
