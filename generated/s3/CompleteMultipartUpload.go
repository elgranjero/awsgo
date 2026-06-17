package s3

// CompleteMultipartUpload is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// Completes a multipart upload by assembling previously uploaded parts.
//
// You first initiate the multipart upload and then upload all parts using the [UploadPart]
// operation or the [UploadPartCopy]operation. After successfully uploading all relevant parts of
// an upload, you call this CompleteMultipartUpload operation to complete the
// upload. Upon receiving this request, Amazon S3 concatenates all the parts in
// ascending order by part number to create a new object. In the
// CompleteMultipartUpload request, you must provide the parts list and ensure that
// the parts list is complete. The CompleteMultipartUpload API operation
// concatenates the parts that you provide in the list. For each part in the list,
// you must provide the PartNumber value and the ETag value that are returned
// after that part was uploaded.
//
// The processing of a CompleteMultipartUpload request could take several minutes
// to finalize. After Amazon S3 begins processing the request, it sends an HTTP
// response header that specifies a 200 OK response. While processing is in
// progress, Amazon S3 periodically sends white space characters to keep the
// connection from timing out. A request could fail after the initial 200 OK
// response has been sent. This means that a 200 OK response can contain either a
// success or an error. The error response might be embedded in the 200 OK
// response. If you call this API operation directly, make sure to design your
// application to parse the contents of the response and handle it appropriately.
// If you use Amazon Web Services SDKs, SDKs handle this condition. The SDKs detect
// the embedded error and apply error handling per your configuration settings
// (including automatically retrying the request as appropriate). If the condition
// persists, the SDKs throw an exception (or, for the SDKs that don't use
// exceptions, they return an error).
//
// Note that if CompleteMultipartUpload fails, applications should be prepared to
// retry any failed requests (including 500 error responses). For more information,
// see [Amazon S3 Error Best Practices].
//
// You can't use Content-Type: application/x-www-form-urlencoded for the
// CompleteMultipartUpload requests. Also, if you don't provide a Content-Type
// header, CompleteMultipartUpload can still return a 200 OK response.
//
// For more information about multipart uploads, see [Uploading Objects Using Multipart Upload] in the Amazon S3 User Guide.
//
// Directory buckets - For directory buckets, you must make requests for this API
// operation to the Zonal endpoint. These endpoints support virtual-hosted-style
// requests in the format
// https://amzn-s3-demo-bucket.s3express-zone-id.region-code.amazonaws.com/key-name
// . Path-style requests are not supported. For more information about endpoints
// in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information
// about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// Permissions
// - General purpose bucket permissions - For information about permissions
// required to use the multipart upload API, see [Multipart Upload and Permissions]in the Amazon S3 User Guide.
//
// If you provide an [additional checksum value]in your MultipartUpload requests and the object is encrypted
//
// with Key Management Service, you must have permission to use the kms:Decrypt
// action for the CompleteMultipartUpload request to succeed.
//
// - Directory bucket permissions - To grant access to this API operation on a
// directory bucket, we recommend that you use the [CreateSession]CreateSession API operation
// for session-based authorization. Specifically, you grant the
// s3express:CreateSession permission to the directory bucket in a bucket policy
// or an IAM identity-based policy. Then, you make the CreateSession API call on
// the bucket to obtain a session token. With the session token in your request
// header, you can make API requests to this operation. After the session token
// expires, you make another CreateSession API call to generate a new session
// token for use. Amazon Web Services CLI or SDKs create session and refresh the
// session token automatically to avoid service interruptions when a session
// expires. For more information about authorization, see [CreateSession]CreateSession .
//
// If the object is encrypted with SSE-KMS, you must also have the
//
// kms:GenerateDataKey and kms:Decrypt permissions in IAM identity-based policies
// and KMS key policies for the KMS key.
//
// Special errors
//
// - Error Code: EntityTooSmall
//
// - Description: Your proposed upload is smaller than the minimum allowed
// object size. Each part must be at least 5 MB in size, except the last part.
//
// - HTTP Status Code: 400 Bad Request
//
// - Error Code: InvalidPart
//
// - Description: One or more of the specified parts could not be found. The
// part might not have been uploaded, or the specified ETag might not have matched
// the uploaded part's ETag.
//
// - HTTP Status Code: 400 Bad Request
//
// - Error Code: InvalidPartOrder
//
// - Description: The list of parts was not in ascending order. The parts list
// must be specified in order by part number.
//
// - HTTP Status Code: 400 Bad Request
//
// - Error Code: NoSuchUpload
//
// - Description: The specified multipart upload does not exist. The upload ID
// might be invalid, or the multipart upload might have been aborted or completed.
//
// - HTTP Status Code: 404 Not Found
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// The following operations are related to CompleteMultipartUpload :
//
// [CreateMultipartUpload]
//
// [UploadPart]
//
// [AbortMultipartUpload]
//
// [ListParts]
//
// [ListMultipartUploads]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Uploading Objects Using Multipart Upload]: https://docs.aws.amazon.com/AmazonS3/latest/dev/uploadobjusingmpu.html
// [Amazon S3 Error Best Practices]: https://docs.aws.amazon.com/AmazonS3/latest/dev/ErrorBestPractices.html
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [ListParts]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html
// [UploadPart]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html
// [additional checksum value]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_Checksum.html
// [UploadPartCopy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPartCopy.html
// [CreateMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html
// [AbortMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html
// [ListMultipartUploads]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListMultipartUploads.html
// [Multipart Upload and Permissions]: https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
//
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
