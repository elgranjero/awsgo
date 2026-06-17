package s3

// ListMultipartUploads is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation lists in-progress multipart uploads in a bucket. An in-progress
// multipart upload is a multipart upload that has been initiated by the
// CreateMultipartUpload request, but has not yet been completed or aborted.
//
// Directory buckets - If multipart uploads in a directory bucket are in progress,
// you can't delete the bucket until all the in-progress multipart uploads are
// aborted or completed. To delete these in-progress multipart uploads, use the
// ListMultipartUploads operation to list the in-progress multipart uploads in the
// bucket and use the AbortMultipartUpload operation to abort all the in-progress
// multipart uploads.
//
// The ListMultipartUploads operation returns a maximum of 1,000 multipart uploads
// in the response. The limit of 1,000 multipart uploads is also the default value.
// You can further limit the number of uploads in a response by specifying the
// max-uploads request parameter. If there are more than 1,000 multipart uploads
// that satisfy your ListMultipartUploads request, the response returns an
// IsTruncated element with the value of true , a NextKeyMarker element, and a
// NextUploadIdMarker element. To list the remaining multipart uploads, you need to
// make subsequent ListMultipartUploads requests. In these requests, include two
// query parameters: key-marker and upload-id-marker . Set the value of key-marker
// to the NextKeyMarker value from the previous response. Similarly, set the value
// of upload-id-marker to the NextUploadIdMarker value from the previous response.
//
// Directory buckets - The upload-id-marker element and the NextUploadIdMarker
// element aren't supported by directory buckets. To list the additional multipart
// uploads, you only need to set the value of key-marker to the NextKeyMarker
// value from the previous response.
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
//
// - General purpose bucket permissions - For information about permissions
// required to use the multipart upload API, see [Multipart Upload and Permissions]in the Amazon S3 User Guide.
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
// Sorting of multipart uploads in response
//
// - General purpose bucket - In the ListMultipartUploads response, the multipart
// uploads are sorted based on two criteria:
//
// - Key-based sorting - Multipart uploads are initially sorted in ascending
// order based on their object keys.
//
// - Time-based sorting - For uploads that share the same object key, they are
// further sorted in ascending order based on the upload initiation time. Among
// uploads with the same key, the one that was initiated first will appear before
// the ones that were initiated later.
//
// - Directory bucket - In the ListMultipartUploads response, the multipart
// uploads aren't sorted lexicographically based on the object keys.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// The following operations are related to ListMultipartUploads :
//
// [CreateMultipartUpload]
//
// [UploadPart]
//
// [CompleteMultipartUpload]
//
// [ListParts]
//
// [AbortMultipartUpload]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Uploading Objects Using Multipart Upload]: https://docs.aws.amazon.com/AmazonS3/latest/dev/uploadobjusingmpu.html
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [ListParts]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html
// [AbortMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html
// [UploadPart]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
// [Multipart Upload and Permissions]: https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html
// [CompleteMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [CreateMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html
