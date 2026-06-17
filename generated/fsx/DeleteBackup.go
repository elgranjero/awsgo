package fsx

// DeleteBackup is generated as a reference stub.
// Executable command wiring lives under cmd/fsx.go.
//
// Deletes an Amazon FSx backup. After deletion, the backup no longer exists, and
// its data is gone.
//
// The DeleteBackup call returns instantly. The backup won't show up in later
// DescribeBackups calls.
//
// The data in a deleted backup is also deleted and can't be recovered by any
// means.
