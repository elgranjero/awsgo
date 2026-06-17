package storagegateway

// EvictFilesFailingUpload is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Starts a process that cleans the specified file share's cache of file entries
// that are failing upload to Amazon S3. This API operation reports success if the
// request is received with valid arguments, and there are no other cache clean
// operations currently in-progress for the specified file share. After a
// successful request, the cache clean operation occurs asynchronously and reports
// progress using CloudWatch logs and notifications.
//
// If ForceRemove is set to True , the cache clean operation will delete file data
// from the gateway which might otherwise be recoverable. We recommend using this
// operation only after all other methods to clear files failing upload have been
// exhausted, and if your business need outweighs the potential data loss.
