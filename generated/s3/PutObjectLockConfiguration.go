package s3

// PutObjectLockConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Places an Object Lock configuration on the specified bucket. The rule specified
// in the Object Lock configuration will be applied by default to every new object
// placed in the specified bucket. For more information, see [Locking Objects].
//
// - The DefaultRetention settings require both a mode and a period.
//
// - The DefaultRetention period can be either Days or Years but you must select
// one. You cannot specify Days and Years at the same time.
//
// - You can enable Object Lock for new or existing buckets. For more
// information, see [Configuring Object Lock].
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Configuring Object Lock]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-configure.html
// [Locking Objects]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html
