package s3control

// ListAccessPoints is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// This operation is not supported by directory buckets.
//
// Returns a list of the access points. You can retrieve up to 1,000 access points
// per call. If the call returns more than 1,000 access points (or the number
// specified in maxResults , whichever is less), the response will include a
// continuation token that you can use to list the additional access points.
//
// Returns only access points attached to S3 buckets by default. To return all
// access points specify DataSourceType as ALL .
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to ListAccessPoints :
//
// [CreateAccessPoint]
//
// [DeleteAccessPoint]
//
// [GetAccessPoint]
//
// [CreateAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPoint.html
// [GetAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPoint.html
// [DeleteAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPoint.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPoint.html#API_control_GetAccessPoint_Examples
