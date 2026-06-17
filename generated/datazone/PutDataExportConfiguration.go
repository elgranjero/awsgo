package datazone

// PutDataExportConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/datazone.go.
//
// Creates data export configuration details.
//
// If you want to temporarily disable export and later re-enable it for the same
// domain, use the --no-enable-export flag to disable and the --enable-export flag
// to re-enable. This preserves the configuration and allows you to re-enable
// export without deleting S3 table.
//
// You can enable asset metadata export for only one domain per account per
// Region. To enable export for a different domain, complete the following steps:
//
// - Delete the export configuration for the currently enabled domain using the
// DeleteDataExportConfiguration operation.
//
// - Delete the asset S3 table under the aws-sagemaker-catalog S3 table bucket.
// We recommend backing up the S3 table before deletion.
//
// - Call the PutDataExportConfiguration API to enable export for the new domain.
