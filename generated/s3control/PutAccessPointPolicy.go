package s3control

// PutAccessPointPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// Associates an access policy with the specified access point. Each access point
// can have only one policy, so a request made to this API replaces any existing
// policy associated with the specified access point.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to PutAccessPointPolicy :
//
// [GetAccessPointPolicy]
//
// [DeleteAccessPointPolicy]
//
// [GetAccessPointPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointPolicy.html
// [DeleteAccessPointPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPointPolicy.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointPolicy.html#API_control_PutAccessPointPolicy_Examples
