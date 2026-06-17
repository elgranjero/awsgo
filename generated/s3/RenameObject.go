package s3

// RenameObject is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// Renames an existing object in a directory bucket that uses the S3 Express One
// Zone storage class. You can use RenameObject by specifying an existing object’s
// name as the source and the new name of the object as the destination within the
// same directory bucket.
//
// RenameObject is only supported for objects stored in the S3 Express One Zone
// storage class.
//
// To prevent overwriting an object, you can use the If-None-Match conditional
// header.
//
// - If-None-Match - Renames the object only if an object with the specified
// name does not already exist in the directory bucket. If you don't want to
// overwrite an existing object, you can add the If-None-Match conditional header
// with the value ‘*’ in the RenameObject request. Amazon S3 then returns a 412
// Precondition Failed error if the object with the specified name already
// exists. For more information, see [RFC 7232].
//
// Permissions  To grant access to the RenameObject operation on a directory
// bucket, we recommend that you use the CreateSession operation for session-based
// authorization. Specifically, you grant the s3express:CreateSession permission
// to the directory bucket in a bucket policy or an IAM identity-based policy.
// Then, you make the CreateSession API call on the directory bucket to obtain a
// session token. With the session token in your request header, you can make API
// requests to this operation. After the session token expires, you make another
// CreateSession API call to generate a new session token for use. The Amazon Web
// Services CLI and SDKs will create and manage your session including refreshing
// the session token automatically to avoid service interruptions when a session
// expires. In your bucket policy, you can specify the s3express:SessionMode
// condition key to control who can create a ReadWrite or ReadOnly session. A
// ReadWrite session is required for executing all the Zonal endpoint API
// operations, including RenameObject . For more information about authorization,
// see [CreateSession]CreateSession . To learn more about Zonal endpoint API operations, see [Authorizing Zonal endpoint API operations with CreateSession] in
// the Amazon S3 User Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
// [RFC 7232]: https://datatracker.ietf.org/doc/rfc7232/
// [Authorizing Zonal endpoint API operations with CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-create-session.html
