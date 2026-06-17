package s3control

// PutMultiRegionAccessPointPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// This operation is not supported by directory buckets.
//
// Associates an access control policy with the specified Multi-Region Access
// Point. Each Multi-Region Access Point can have only one policy, so a request
// made to this action replaces any existing policy that is associated with the
// specified Multi-Region Access Point.
//
// This action will always be routed to the US West (Oregon) Region. For more
// information about the restrictions around working with Multi-Region Access
// Points, see [Multi-Region Access Point restrictions and limitations]in the Amazon S3 User Guide.
//
// The following actions are related to PutMultiRegionAccessPointPolicy :
//
// [GetMultiRegionAccessPointPolicy]
//
// [GetMultiRegionAccessPointPolicyStatus]
//
// [GetMultiRegionAccessPointPolicyStatus]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetMultiRegionAccessPointPolicyStatus.html
// [GetMultiRegionAccessPointPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetMultiRegionAccessPointPolicy.html
// [Multi-Region Access Point restrictions and limitations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/MultiRegionAccessPointRestrictions.html
