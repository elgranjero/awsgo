package fsx

// DeleteSnapshot is generated as a reference stub.
// Executable command wiring lives under cmd/fsx.go.
//
// Deletes an Amazon FSx for OpenZFS snapshot. After deletion, the snapshot no
// longer exists, and its data is gone. Deleting a snapshot doesn't affect
// snapshots stored in a file system backup.
//
// The DeleteSnapshot operation returns instantly. The snapshot appears with the
// lifecycle status of DELETING until the deletion is complete.
