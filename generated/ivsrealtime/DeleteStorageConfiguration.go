package ivsrealtime

// DeleteStorageConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/ivsrealtime.go.
//
// Deletes the storage configuration for the specified ARN.
//
// If you try to delete a storage configuration that is used by a Composition, you
// will get an error (409 ConflictException). To avoid this, for all Compositions
// that reference the storage configuration, first use StopCompositionand wait for it to
// complete, then use DeleteStorageConfiguration.
