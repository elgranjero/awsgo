package storagegateway

// DeleteCacheReport is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Deletes the specified cache report and any associated tags from the Storage
// Gateway database. You can only delete completed reports. If the status of the
// report you attempt to delete still IN-PROGRESS, the delete operation returns an
// error. You can use CancelCacheReport to cancel an IN-PROGRESS report.
//
// DeleteCacheReport does not delete the report object from your Amazon S3 bucket.
