package s3

// PutBucketCors is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Sets the cors configuration for your bucket. If the configuration exists,
// Amazon S3 replaces it.
//
// To use this operation, you must be allowed to perform the s3:PutBucketCORS
// action. By default, the bucket owner has this permission and can grant it to
// others.
//
// You set this configuration on a bucket so that the bucket can service
// cross-origin requests. For example, you might want to enable a request whose
// origin is http://www.example.com to access your Amazon S3 bucket at
// my.example.bucket.com by using the browser's XMLHttpRequest capability.
//
// To enable cross-origin resource sharing (CORS) on a bucket, you add the cors
// subresource to the bucket. The cors subresource is an XML document in which you
// configure rules that identify origins and the HTTP methods that can be executed
// on your bucket. The document is limited to 64 KB in size.
//
// When Amazon S3 receives a cross-origin request (or a pre-flight OPTIONS
// request) against a bucket, it evaluates the cors configuration on the bucket
// and uses the first CORSRule rule that matches the incoming browser request to
// enable a cross-origin request. For a rule to match, the following conditions
// must be met:
//
// - The request's Origin header must match AllowedOrigin elements.
//
// - The request method (for example, GET, PUT, HEAD, and so on) or the
// Access-Control-Request-Method header in case of a pre-flight OPTIONS request
// must be one of the AllowedMethod elements.
//
// - Every header specified in the Access-Control-Request-Headers request header
// of a pre-flight request must match an AllowedHeader element.
//
// For more information about CORS, go to [Enabling Cross-Origin Resource Sharing] in the Amazon S3 User Guide.
//
// The following operations are related to PutBucketCors :
//
// [GetBucketCors]
//
// [DeleteBucketCors]
//
// [RESTOPTIONSobject]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetBucketCors]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketCors.html
// [Enabling Cross-Origin Resource Sharing]: https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html
// [RESTOPTIONSobject]: https://docs.aws.amazon.com/AmazonS3/latest/API/RESTOPTIONSobject.html
// [DeleteBucketCors]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketCors.html
