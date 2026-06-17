package fsx

// CreateDataRepositoryAssociation is generated as a reference stub.
// Executable command wiring lives under cmd/fsx.go.
//
// Creates an Amazon FSx for Lustre data repository association (DRA). A data
// repository association is a link between a directory on the file system and an
// Amazon S3 bucket or prefix. You can have a maximum of 8 data repository
// associations on a file system. Data repository associations are supported on all
// FSx for Lustre 2.12 and 2.15 file systems, excluding scratch_1 deployment type.
//
// Each data repository association must have a unique Amazon FSx file system
// directory and a unique S3 bucket or prefix associated with it. You can configure
// a data repository association for automatic import only, for automatic export
// only, or for both. To learn more about linking a data repository to your file
// system, see [Linking your file system to an S3 bucket].
//
// CreateDataRepositoryAssociation isn't supported on Amazon File Cache resources.
// To create a DRA on Amazon File Cache, use the CreateFileCache operation.
//
// [Linking your file system to an S3 bucket]: https://docs.aws.amazon.com/fsx/latest/LustreGuide/create-dra-linked-data-repo.html
