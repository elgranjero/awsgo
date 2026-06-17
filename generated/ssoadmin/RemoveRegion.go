package ssoadmin

// RemoveRegion is generated as a reference stub.
// Executable command wiring lives under cmd/ssoadmin.go.
//
// Removes an additional Region from an IAM Identity Center instance. This
// operation initiates an asynchronous workflow to clean up IAM Identity Center
// resources in the specified additional Region. The Region status is set to
// REMOVING and the Region record is deleted when the workflow completes. The
// request must be made from the primary Region. The target Region cannot be the
// primary Region, and no other add or remove Region workflows can be in progress.
//
// The following actions are related to RemoveRegion :
//
// [AddRegion]
//
// [DescribeRegion]
//
// [ListRegions]
//
// [AddRegion]: https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_AddRegion.html
// [DescribeRegion]: https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_DescribeRegion.html
// [ListRegions]: https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_ListRegions.html
