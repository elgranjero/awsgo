package s3

// HeadBucket is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// You can use this operation to determine if a bucket exists and if you have
// permission to access it. The action returns a 200 OK HTTP status code if the
// bucket exists and you have permission to access it. You can make a HeadBucket
// call on any bucket name to any Region in the partition, and regardless of the
// permissions on the bucket, you will receive a response header with the correct
// bucket location so that you can then make a proper, signed request to the
// appropriate Regional endpoint.
//
// If the bucket doesn't exist or you don't have permission to access it, the HEAD
// request returns a generic 400 Bad Request , 403 Forbidden , or 404 Not Found
// HTTP status code. A message body isn't included, so you can't determine the
// exception beyond these HTTP response codes.
//
// Authentication and authorization  General purpose buckets - Request to public
// buckets that grant the s3:ListBucket permission publicly do not need to be
// signed. All other HeadBucket requests must be authenticated and signed by using
// IAM credentials (access key ID and secret access key for the IAM identities).
// All headers with the x-amz- prefix, including x-amz-copy-source , must be
// signed. For more information, see [REST Authentication].
//
// Directory buckets - You must use IAM credentials to authenticate and authorize
// your access to the HeadBucket API operation, instead of using the temporary
// security credentials through the CreateSession API operation.
//
// Amazon Web Services CLI or SDKs handles authentication and authorization on
// your behalf.
//
// Permissions
//
// - General purpose bucket permissions - To use this operation, you must have
// permissions to perform the s3:ListBucket action. The bucket owner has this
// permission by default and can grant this permission to others. For more
// information about permissions, see [Managing access permissions to your Amazon S3 resources]in the Amazon S3 User Guide.
//
// - Directory bucket permissions - You must have the s3express:CreateSession
// permission in the Action element of a policy. By default, the session is in
// the ReadWrite mode. If you want to restrict the access, you can explicitly set
// the s3express:SessionMode condition key to ReadOnly on the bucket.
//
// For more information about example bucket policies, see [Example bucket policies for S3 Express One Zone]and [Amazon Web Services Identity and Access Management (IAM) identity-based policies for S3 Express One Zone]in the Amazon S3
//
// User Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// You must make requests for this API operation to the Zonal endpoint. These
// endpoints support virtual-hosted-style requests in the format
// https://bucket-name.s3express-zone-id.region-code.amazonaws.com . Path-style
// requests are not supported. For more information about endpoints in Availability
// Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information about endpoints in
// Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Amazon Web Services Identity and Access Management (IAM) identity-based policies for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam-identity-policies.html
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [REST Authentication]: https://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html
// [Example bucket policies for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam-example-bucket-policies.html
// [Managing access permissions to your Amazon S3 resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
