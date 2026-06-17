package signer

// StartSigningJob is generated as a reference stub.
// Executable command wiring lives under cmd/signer.go.
//
// Initiates a signing job to be performed on the code provided. Signing jobs are
// viewable by the ListSigningJobs operation. Note the following requirements:
//
// - You must create an Amazon S3 source bucket. For more information, see [Creating a Bucket]in
// the Amazon S3 Getting Started Guide.
//
// - Your S3 source bucket must be version enabled.
//
// - You must create an S3 destination bucket. AWS Signer uses your S3
// destination bucket to write your signed code.
//
// - You specify the name of the source and destination buckets when calling the
// StartSigningJob operation.
//
// - You must ensure the S3 buckets are from the same Region as the signing
// profile. Cross-Region signing isn't supported.
//
// - You must also specify a request token that identifies your request to
// Signer.
//
// You can call the DescribeSigningJob and the ListSigningJobs actions after you call StartSigningJob .
//
// For a Java example that shows how to use this action, see [StartSigningJob].
//
// [Creating a Bucket]: http://docs.aws.amazon.com/AmazonS3/latest/gsg/CreatingABucket.html
// [StartSigningJob]: https://docs.aws.amazon.com/signer/latest/developerguide/api-startsigningjob.html
