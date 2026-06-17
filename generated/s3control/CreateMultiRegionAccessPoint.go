package s3control

// CreateMultiRegionAccessPoint is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// This operation is not supported by directory buckets.
//
// Creates a Multi-Region Access Point and associates it with the specified
// buckets. For more information about creating Multi-Region Access Points, see [Creating Multi-Region Access Points]in
// the Amazon S3 User Guide.
//
// This action will always be routed to the US West (Oregon) Region. For more
// information about the restrictions around working with Multi-Region Access
// Points, see [Multi-Region Access Point restrictions and limitations]in the Amazon S3 User Guide.
//
// This request is asynchronous, meaning that you might receive a response before
// the command has completed. When this request provides a response, it provides a
// token that you can use to monitor the status of the request with
// DescribeMultiRegionAccessPointOperation .
//
// The following actions are related to CreateMultiRegionAccessPoint :
//
// [DeleteMultiRegionAccessPoint]
//
// [DescribeMultiRegionAccessPointOperation]
//
// [GetMultiRegionAccessPoint]
//
// [ListMultiRegionAccessPoints]
//
// [Creating Multi-Region Access Points]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/CreatingMultiRegionAccessPoints.html
// [DeleteMultiRegionAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteMultiRegionAccessPoint.html
// [GetMultiRegionAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetMultiRegionAccessPoint.html
// [ListMultiRegionAccessPoints]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListMultiRegionAccessPoints.html
// [DescribeMultiRegionAccessPointOperation]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DescribeMultiRegionAccessPointOperation.html
// [Multi-Region Access Point restrictions and limitations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/MultiRegionAccessPointRestrictions.html
