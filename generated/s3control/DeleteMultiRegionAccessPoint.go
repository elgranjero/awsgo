package s3control

// DeleteMultiRegionAccessPoint is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// This operation is not supported by directory buckets.
//
// Deletes a Multi-Region Access Point. This action does not delete the buckets
// associated with the Multi-Region Access Point, only the Multi-Region Access
// Point itself.
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
// The following actions are related to DeleteMultiRegionAccessPoint :
//
// [CreateMultiRegionAccessPoint]
//
// [DescribeMultiRegionAccessPointOperation]
//
// [GetMultiRegionAccessPoint]
//
// [ListMultiRegionAccessPoints]
//
// [GetMultiRegionAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetMultiRegionAccessPoint.html
// [ListMultiRegionAccessPoints]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListMultiRegionAccessPoints.html
// [DescribeMultiRegionAccessPointOperation]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DescribeMultiRegionAccessPointOperation.html
// [CreateMultiRegionAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateMultiRegionAccessPoint.html
// [Multi-Region Access Point restrictions and limitations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/MultiRegionAccessPointRestrictions.html
