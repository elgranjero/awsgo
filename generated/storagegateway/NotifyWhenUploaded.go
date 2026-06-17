package storagegateway

// NotifyWhenUploaded is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Sends you notification through Amazon EventBridge when all files written to
// your file share have been uploaded to Amazon S3.
//
// Storage Gateway can send a notification through Amazon EventBridge when all
// files written to your file share up to that point in time have been uploaded to
// Amazon S3. These files include files written to the file share up to the time
// that you make a request for notification. When the upload is done, Storage
// Gateway sends you notification through EventBridge. You can configure
// EventBridge to send the notification through event targets such as Amazon SNS or
// Lambda function. This operation is only supported for S3 File Gateways.
//
// For more information, see [Getting file upload notification] in the Amazon S3 File Gateway User Guide.
//
// [Getting file upload notification]: https://docs.aws.amazon.com/filegateway/latest/files3/monitoring-file-gateway.html#get-notification
