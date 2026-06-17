package s3

// PutBucketWebsite is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Sets the configuration of the website that is specified in the website
// subresource. To configure a bucket as a website, you can add this subresource on
// the bucket with website configuration information such as the file name of the
// index document and any redirect rules. For more information, see [Hosting Websites on Amazon S3].
//
// This PUT action requires the S3:PutBucketWebsite permission. By default, only
// the bucket owner can configure the website attached to a bucket; however, bucket
// owners can allow other users to set the website configuration by writing a
// bucket policy that grants them the S3:PutBucketWebsite permission.
//
// To redirect all website requests sent to the bucket's website endpoint, you add
// a website configuration with the following elements. Because all requests are
// sent to another website, you don't need to provide index document name for the
// bucket.
//
// - WebsiteConfiguration
//
// - RedirectAllRequestsTo
//
// - HostName
//
// - Protocol
//
// If you want granular control over redirects, you can use the following elements
// to add routing rules that describe conditions for redirecting requests and
// information about the redirect destination. In this case, the website
// configuration must provide an index document for the bucket, because some
// requests might not be redirected.
//
// - WebsiteConfiguration
//
// - IndexDocument
//
// - Suffix
//
// - ErrorDocument
//
// - Key
//
// - RoutingRules
//
// - RoutingRule
//
// - Condition
//
// - HttpErrorCodeReturnedEquals
//
// - KeyPrefixEquals
//
// - Redirect
//
// - Protocol
//
// - HostName
//
// - ReplaceKeyPrefixWith
//
// - ReplaceKeyWith
//
// - HttpRedirectCode
//
// Amazon S3 has a limitation of 50 routing rules per website configuration. If
// you require more than 50 routing rules, you can use object redirect. For more
// information, see [Configuring an Object Redirect]in the Amazon S3 User Guide.
//
// The maximum request length is limited to 128 KB.
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Hosting Websites on Amazon S3]: https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html
// [Configuring an Object Redirect]: https://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html
