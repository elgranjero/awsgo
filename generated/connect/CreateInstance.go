package connect

// CreateInstance is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// This API is in preview release for Amazon Connect and is subject to change.
//
// Initiates an Amazon Connect instance with all the supported channels enabled.
// It does not attach any storage, such as Amazon Simple Storage Service (Amazon
// S3) or Amazon Kinesis. It also does not allow for any configurations on
// features, such as Contact Lens for Amazon Connect.
//
// For more information, see [Create an Amazon Connect instance] in the Amazon Connect Administrator Guide.
//
// Amazon Connect enforces a limit on the total number of instances that you can
// create or delete in 30 days. If you exceed this limit, you will get an error
// message indicating there has been an excessive number of attempts at creating or
// deleting instances. You must wait 30 days before you can restart creating and
// deleting instances in your account.
//
// [Create an Amazon Connect instance]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-instances.html
