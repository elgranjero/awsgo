package ivsrealtime

// CreateStorageConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/ivsrealtime.go.
//
// Creates a new storage configuration, used to enable recording to Amazon S3.
// When a StorageConfiguration is created, IVS will modify the S3 bucketPolicy of
// the provided bucket. This will ensure that IVS has sufficient permissions to
// write content to the provided bucket.
