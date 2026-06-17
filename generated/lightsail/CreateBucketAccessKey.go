package lightsail

// CreateBucketAccessKey is generated as a reference stub.
// Executable command wiring lives under cmd/lightsail.go.
//
// Creates a new access key for the specified Amazon Lightsail bucket. Access keys
// consist of an access key ID and corresponding secret access key.
//
// Access keys grant full programmatic access to the specified bucket and its
// objects. You can have a maximum of two access keys per bucket. Use the [GetBucketAccessKeys]action
// to get a list of current access keys for a specific bucket. For more information
// about access keys, see [Creating access keys for a bucket in Amazon Lightsail]in the Amazon Lightsail Developer Guide.
//
// The secretAccessKey value is returned only in response to the
// CreateBucketAccessKey action. You can get a secret access key only when you
// first create an access key; you cannot get the secret access key later. If you
// lose the secret access key, you must create a new access key.
//
// [Creating access keys for a bucket in Amazon Lightsail]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-creating-bucket-access-keys
// [GetBucketAccessKeys]: https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_GetBucketAccessKeys.html
