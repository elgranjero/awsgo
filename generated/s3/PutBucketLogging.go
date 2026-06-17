package s3

// PutBucketLogging is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// End of support notice: As of October 1, 2025, Amazon S3 has discontinued
// support for Email Grantee Access Control Lists (ACLs). If you attempt to use an
// Email Grantee ACL in a request after October 1, 2025, the request will receive
// an HTTP 405 (Method Not Allowed) error.
//
// This change affects the following Amazon Web Services Regions: US East (N.
// Virginia), US West (N. California), US West (Oregon), Asia Pacific (Singapore),
// Asia Pacific (Sydney), Asia Pacific (Tokyo), Europe (Ireland), and South America
// (São Paulo).
//
// This operation is not supported for directory buckets.
//
// Set the logging parameters for a bucket and to specify permissions for who can
// view and modify the logging parameters. All logs are saved to buckets in the
// same Amazon Web Services Region as the source bucket. To set the logging status
// of a bucket, you must be the bucket owner.
//
// The bucket owner is automatically granted FULL_CONTROL to all logs. You use the
// Grantee request element to grant access to other people. The Permissions
// request element specifies the kind of access the grantee has to the logs.
//
// If the target bucket for log delivery uses the bucket owner enforced setting
// for S3 Object Ownership, you can't use the Grantee request element to grant
// access to others. Permissions can only be granted using policies. For more
// information, see [Permissions for server access log delivery]in the Amazon S3 User Guide.
//
// Grantee Values You can specify the person (grantee) to whom you're assigning
// access rights (by using request elements) in the following ways. For examples of
// how to specify these grantee values in JSON format, see the Amazon Web Services
// CLI example in [Enabling Amazon S3 server access logging]in the Amazon S3 User Guide.
//
// - By the person's ID:
//
// <>ID<><>GranteesEmail<>
//
// DisplayName is optional and ignored in the request.
//
// - By Email address:
//
// <>Grantees(at)email.com<>
//
// The grantee is resolved to the CanonicalUser and, in a response to a
//
// GETObjectAcl request, appears as the CanonicalUser.
//
// - By URI:
//
// <>http://acs.amazonaws.com/groups/global/AuthenticatedUsers<>
//
// To enable logging, you use LoggingEnabled and its children request elements. To
// disable logging, you use an empty BucketLoggingStatus request element:
//
// For more information about server access logging, see [Server Access Logging] in the Amazon S3 User
// Guide.
//
// For more information about creating a bucket, see [CreateBucket]. For more information about
// returning the logging status of a bucket, see [GetBucketLogging].
//
// The following operations are related to PutBucketLogging :
//
// [PutObject]
//
// [DeleteBucket]
//
// [CreateBucket]
//
// [GetBucketLogging]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Permissions for server access log delivery]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/enable-server-access-logging.html#grant-log-delivery-permissions-general
// [DeleteBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html
// [GetBucketLogging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLogging.html
// [PutObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html
// [Enabling Amazon S3 server access logging]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/enable-server-access-logging.html
// [Server Access Logging]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/ServerLogs.html
