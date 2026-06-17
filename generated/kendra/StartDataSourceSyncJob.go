package kendra

// StartDataSourceSyncJob is generated as a reference stub.
// Executable command wiring lives under cmd/kendra.go.
//
// Starts a synchronization job for a data source connector. If a synchronization
// job is already in progress, Amazon Kendra returns a ResourceInUseException
// exception.
//
// Re-syncing your data source with your index after modifying, adding, or
// deleting documents from your data source respository could take up to an hour or
// more, depending on the number of documents to sync.
