package storagegateway

// RefreshCache is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Refreshes the cached inventory of objects for the specified file share. This
// operation finds objects in the Amazon S3 bucket that were added, removed, or
// replaced since the gateway last listed the bucket's contents and cached the
// results. This operation does not import files into the S3 File Gateway cache
// storage. It only updates the cached inventory to reflect changes in the
// inventory of the objects in the S3 bucket. This operation is only supported in
// the S3 File Gateway types.
//
// You can subscribe to be notified through an Amazon CloudWatch event when your
// RefreshCache operation completes. For more information, see [Getting notified about file operations] in the Amazon S3
// File Gateway User Guide. This operation is Only supported for S3 File Gateways.
//
// When this API is called, it only initiates the refresh operation. When the API
// call completes and returns a success code, it doesn't necessarily mean that the
// file refresh has completed. You should use the refresh-complete notification to
// determine that the operation has completed before you check for new files on the
// gateway file share. You can subscribe to be notified through a CloudWatch event
// when your RefreshCache operation completes.
//
// Throttle limit: This API is asynchronous, so the gateway will accept no more
// than two refreshes at any time. We recommend using the refresh-complete
// CloudWatch event notification before issuing additional requests. For more
// information, see [Getting notified about file operations]in the Amazon S3 File Gateway User Guide.
//
// - Wait at least 60 seconds between consecutive RefreshCache API requests.
//
// - If you invoke the RefreshCache API when two requests are already being
// processed, any new request will cause an InvalidGatewayRequestException error
// because too many requests were sent to the server.
//
// The S3 bucket name does not need to be included when entering the list of
// folders in the FolderList parameter.
//
// For more information, see [Getting notified about file operations] in the Amazon S3 File Gateway User Guide.
//
// [Getting notified about file operations]: https://docs.aws.amazon.com/filegateway/latest/files3/monitoring-file-gateway.html#get-notification
