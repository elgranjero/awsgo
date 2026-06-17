package s3

// PutBucketVersioning is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// When you enable versioning on a bucket for the first time, it might take a
// short amount of time for the change to be fully propagated. While this change is
// propagating, you might encounter intermittent HTTP 404 NoSuchKey errors for
// requests to objects created or updated after enabling versioning. We recommend
// that you wait for 15 minutes after enabling versioning before issuing write
// operations ( PUT or DELETE ) on objects in the bucket.
//
// Sets the versioning state of an existing bucket.
//
// You can set the versioning state with one of the following values:
//
// Enabled—Enables versioning for the objects in the bucket. All objects added to
// the bucket receive a unique version ID.
//
// Suspended—Disables versioning for the objects in the bucket. All objects added
// to the bucket receive the version ID null.
//
// If the versioning state has never been set on a bucket, it has no versioning
// state; a [GetBucketVersioning]request does not return a versioning state value.
//
// In order to enable MFA Delete, you must be the bucket owner. If you are the
// bucket owner and want to enable MFA Delete in the bucket versioning
// configuration, you must include the x-amz-mfa request header and the Status and
// the MfaDelete request elements in a request to set the versioning state of the
// bucket.
//
// If you have an object expiration lifecycle configuration in your non-versioned
// bucket and you want to maintain the same permanent delete behavior when you
// enable versioning, you must add a noncurrent expiration policy. The noncurrent
// expiration lifecycle configuration will manage the deletes of the noncurrent
// object versions in the version-enabled bucket. (A version-enabled bucket
// maintains one current and zero or more noncurrent object versions.) For more
// information, see [Lifecycle and Versioning].
//
// The following operations are related to PutBucketVersioning :
//
// [CreateBucket]
//
// [DeleteBucket]
//
// [GetBucketVersioning]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [DeleteBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html
// [Lifecycle and Versioning]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html#lifecycle-and-other-bucket-config
// [GetBucketVersioning]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketVersioning.html
