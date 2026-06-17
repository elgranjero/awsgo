package appfabric

// DeleteIngestionDestination is generated as a reference stub.
// Executable command wiring lives under cmd/appfabric.go.
//
// Deletes an ingestion destination.
//
// This deletes the association between an ingestion and it's destination. It
// doesn't delete previously ingested data or the storage destination, such as the
// Amazon S3 bucket where the data is delivered. If the ingestion destination is
// deleted while the associated ingestion is enabled, the ingestion will fail and
// is eventually disabled.
