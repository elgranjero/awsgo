package storagegateway

// UpdateSMBFileShare is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Updates a Server Message Block (SMB) file share. This operation is only
// supported for S3 File Gateways.
//
// To leave a file share field unchanged, set the corresponding input field to
// null.
//
// File gateways require Security Token Service (Amazon Web Services STS) to be
// activated to enable you to create a file share. Make sure that Amazon Web
// Services STS is activated in the Amazon Web Services Region you are creating
// your file gateway in. If Amazon Web Services STS is not activated in this Amazon
// Web Services Region, activate it. For information about how to activate Amazon
// Web Services STS, see [Activating and deactivating Amazon Web Services STS in an Amazon Web Services Region]in the Identity and Access Management User Guide.
//
// File gateways don't support creating hard or symbolic links on a file share.
//
// [Activating and deactivating Amazon Web Services STS in an Amazon Web Services Region]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html
