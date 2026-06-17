package s3

// GetBucketNotificationConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Returns the notification configuration of a bucket.
//
// If notifications are not enabled on the bucket, the action returns an empty
// NotificationConfiguration element.
//
// By default, you must be the bucket owner to read the notification configuration
// of a bucket. However, the bucket owner can use a bucket policy to grant
// permission to other users to read this configuration with the
// s3:GetBucketNotification permission.
//
// When you use this API operation with an access point, provide the alias of the
// access point in place of the bucket name.
//
// When you use this API operation with an Object Lambda access point, provide the
// alias of the Object Lambda access point in place of the bucket name. If the
// Object Lambda access point alias in a request is not valid, the error code
// InvalidAccessPointAliasError is returned. For more information about
// InvalidAccessPointAliasError , see [List of Error Codes].
//
// For more information about setting and reading the notification configuration
// on a bucket, see [Setting Up Notification of Bucket Events]. For more information about bucket policies, see [Using Bucket Policies].
//
// The following action is related to GetBucketNotification :
//
// [PutBucketNotification]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Using Bucket Policies]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html
// [Setting Up Notification of Bucket Events]: https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html
// [List of Error Codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#ErrorCodeList
// [PutBucketNotification]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketNotification.html
