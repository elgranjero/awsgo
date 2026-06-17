package s3control

// CreateAccessPoint is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// Creates an access point and associates it to a specified bucket. For more
// information, see [Managing access to shared datasets with access points]or [Managing access to shared datasets in directory buckets with access points] in the Amazon S3 User Guide.
//
// To create an access point and attach it to a volume on an Amazon FSx file
// system, see [CreateAndAttachS3AccessPoint]in the Amazon FSx API Reference.
//
// S3 on Outposts only supports VPC-style access points.
//
// For more information, see [Accessing Amazon S3 on Outposts using virtual private cloud (VPC) only access points] in the Amazon S3 User Guide.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to CreateAccessPoint :
//
// [GetAccessPoint]
//
// [DeleteAccessPoint]
//
// [ListAccessPoints]
//
// [ListAccessPointsForDirectoryBuckets]
//
// [ListAccessPointsForDirectoryBuckets]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessPointsForDirectoryBuckets.html
// [ListAccessPoints]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessPoints.html
// [CreateAndAttachS3AccessPoint]: https://docs.aws.amazon.com/fsx/latest/APIReference/API_CreateAndAttachS3AccessPoint.html
// [Managing access to shared datasets with access points]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points.html
// [GetAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPoint.html
// [Accessing Amazon S3 on Outposts using virtual private cloud (VPC) only access points]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [DeleteAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPoint.html
// [Managing access to shared datasets in directory buckets with access points]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-directory-buckets.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPoint.html#API_control_CreateAccessPoint_Examples
