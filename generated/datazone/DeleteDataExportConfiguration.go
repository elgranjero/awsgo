package datazone

// DeleteDataExportConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/datazone.go.
//
// Deletes data export configuration for a domain.
//
// This operation does not delete the S3 table created by the
// PutDataExportConfiguration operation.
//
// To temporarily disable export without deleting the configuration, use the
// PutDataExportConfiguration operation with the --no-enable-export flag instead.
// This allows you to re-enable export for the same domain using the
// --enable-export flag without deleting S3 table.
