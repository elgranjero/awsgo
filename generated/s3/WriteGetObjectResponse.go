package s3

// WriteGetObjectResponse is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Passes transformed objects to a GetObject operation when using Object Lambda
// access points. For information about Object Lambda access points, see [Transforming objects with Object Lambda access points]in the
// Amazon S3 User Guide.
//
// This operation supports metadata that can be returned by [GetObject], in addition to
// RequestRoute , RequestToken , StatusCode , ErrorCode , and ErrorMessage . The
// GetObject response metadata is supported so that the WriteGetObjectResponse
// caller, typically an Lambda function, can provide the same metadata when it
// internally invokes GetObject . When WriteGetObjectResponse is called by a
// customer-owned Lambda function, the metadata returned to the end user GetObject
// call might differ from what Amazon S3 would normally return.
//
// You can include any number of metadata headers. When including a metadata
// header, it should be prefaced with x-amz-meta . For example,
// x-amz-meta-my-custom-header: MyCustomValue . The primary use case for this is to
// forward GetObject metadata.
//
// Amazon Web Services provides some prebuilt Lambda functions that you can use
// with S3 Object Lambda to detect and redact personally identifiable information
// (PII) and decompress S3 objects. These Lambda functions are available in the
// Amazon Web Services Serverless Application Repository, and can be selected
// through the Amazon Web Services Management Console when you create your Object
// Lambda access point.
//
// Example 1: PII Access Control - This Lambda function uses Amazon Comprehend, a
// natural language processing (NLP) service using machine learning to find
// insights and relationships in text. It automatically detects personally
// identifiable information (PII) such as names, addresses, dates, credit card
// numbers, and social security numbers from documents in your Amazon S3 bucket.
//
// Example 2: PII Redaction - This Lambda function uses Amazon Comprehend, a
// natural language processing (NLP) service using machine learning to find
// insights and relationships in text. It automatically redacts personally
// identifiable information (PII) such as names, addresses, dates, credit card
// numbers, and social security numbers from documents in your Amazon S3 bucket.
//
// Example 3: Decompression - The Lambda function S3ObjectLambdaDecompression, is
// equipped to decompress objects stored in S3 in one of six compressed file
// formats including bzip2, gzip, snappy, zlib, zstandard and ZIP.
//
// For information on how to view and use these functions, see [Using Amazon Web Services built Lambda functions] in the Amazon S3
// User Guide.
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Transforming objects with Object Lambda access points]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/transforming-objects.html
// [Using Amazon Web Services built Lambda functions]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/olap-examples.html
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
