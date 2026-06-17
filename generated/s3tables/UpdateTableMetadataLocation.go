package s3tables

// UpdateTableMetadataLocation is generated as a reference stub.
// Executable command wiring lives under cmd/s3tables.go.
//
// Updates the metadata location for a table. The metadata location of a table
// must be an S3 URI that begins with the table's warehouse location. The metadata
// location for an Apache Iceberg table must end with .metadata.json , or if the
// metadata file is Gzip-compressed, .metadata.json.gz .
//
// Permissions You must have the s3tables:UpdateTableMetadataLocation permission
// to use this operation.
