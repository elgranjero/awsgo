package ssoadmin

// AddRegion is generated as a reference stub.
// Executable command wiring lives under cmd/ssoadmin.go.
//
// Adds a Region to an IAM Identity Center instance. This operation initiates an
// asynchronous workflow to replicate the IAM Identity Center instance to the
// target Region. The Region status is set to ADDING at first and changes to ACTIVE
// when the workflow completes.
//
// To use this operation, your IAM Identity Center instance and the target Region
// must meet the requirements described in the [IAM Identity Center User Guide].
//
// The following actions are related to AddRegion :
//
// [RemoveRegion]
//
// [DescribeRegion]
//
// [ListRegions]
//
// [IAM Identity Center User Guide]: https://docs.aws.amazon.com/singlesignon/latest/userguide/multi-region-iam-identity-center.html#multi-region-prerequisites
// [RemoveRegion]: https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_RemoveRegion.html
// [DescribeRegion]: https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_DescribeRegion.html
// [ListRegions]: https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_ListRegions.html
