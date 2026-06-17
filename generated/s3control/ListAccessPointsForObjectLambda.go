package s3control

// ListAccessPointsForObjectLambda is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// This operation is not supported by directory buckets.
//
// Returns some or all (up to 1,000) access points associated with the Object
// Lambda Access Point per call. If there are more access points than what can be
// returned in one call, the response will include a continuation token that you
// can use to list the additional access points.
//
// The following actions are related to ListAccessPointsForObjectLambda :
//
// [CreateAccessPointForObjectLambda]
//
// [DeleteAccessPointForObjectLambda]
//
// [GetAccessPointForObjectLambda]
//
// [CreateAccessPointForObjectLambda]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPointForObjectLambda.html
// [DeleteAccessPointForObjectLambda]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPointForObjectLambda.html
// [GetAccessPointForObjectLambda]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointForObjectLambda.html
