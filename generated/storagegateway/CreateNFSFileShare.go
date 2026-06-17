package storagegateway

// CreateNFSFileShare is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Creates a Network File System (NFS) file share on an existing S3 File Gateway.
// In Storage Gateway, a file share is a file system mount point backed by Amazon
// S3 cloud storage. Storage Gateway exposes file shares using an NFS interface.
// This operation is only supported for S3 File Gateways.
//
// S3 File gateway requires Security Token Service (Amazon Web Services STS) to be
// activated to enable you to create a file share. Make sure Amazon Web Services
// STS is activated in the Amazon Web Services Region you are creating your S3 File
// Gateway in. If Amazon Web Services STS is not activated in the Amazon Web
// Services Region, activate it. For information about how to activate Amazon Web
// Services STS, see [Activating and deactivating Amazon Web Services STS in an Amazon Web Services Region]in the Identity and Access Management User Guide.
//
// S3 File Gateways do not support creating hard or symbolic links on a file share.
//
// [Activating and deactivating Amazon Web Services STS in an Amazon Web Services Region]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html
