package clouddirectory

// UpgradeAppliedSchema is generated as a reference stub.
// Executable command wiring lives under cmd/clouddirectory.go.
//
// Upgrades a single directory in-place using the PublishedSchemaArn with schema
// updates found in MinorVersion . Backwards-compatible minor version upgrades are
// instantaneously available for readers on all objects in the directory. Note:
// This is a synchronous API call and upgrades only one schema on a given directory
// per call. To upgrade multiple directories from one schema, you would need to
// call this API on each directory.
