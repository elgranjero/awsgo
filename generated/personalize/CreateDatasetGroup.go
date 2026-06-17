package personalize

// CreateDatasetGroup is generated as a reference stub.
// Executable command wiring lives under cmd/personalize.go.
//
// Creates an empty dataset group. A dataset group is a container for Amazon
// Personalize resources. A dataset group can contain at most three datasets, one
// for each type of dataset:
//
// - Item interactions
//
// - Items
//
// - Users
//
// - Actions
//
// - Action interactions
//
// A dataset group can be a Domain dataset group, where you specify a domain and
// use pre-configured resources like recommenders, or a Custom dataset group, where
// you use custom resources, such as a solution with a solution version, that you
// deploy with a campaign. If you start with a Domain dataset group, you can still
// add custom resources such as solutions and solution versions trained with
// recipes for custom use cases and deployed with campaigns.
//
// A dataset group can be in one of the following states:
//
// - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// - DELETE PENDING
//
// To get the status of the dataset group, call [DescribeDatasetGroup]. If the status shows as CREATE
// FAILED, the response includes a failureReason key, which describes why the
// creation failed.
//
// You must wait until the status of the dataset group is ACTIVE before adding a
// dataset to the group.
//
// You can specify an Key Management Service (KMS) key to encrypt the datasets in
// the group. If you specify a KMS key, you must also include an Identity and
// Access Management (IAM) role that has permission to access the key.
//
// # APIs that require a dataset group ARN in the request
//
// [CreateDataset]
//
// [CreateEventTracker]
//
// [CreateSolution]
//
// # Related APIs
//
// [ListDatasetGroups]
//
// [DescribeDatasetGroup]
//
// [DeleteDatasetGroup]
//
// [CreateDataset]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDataset.html
// [ListDatasetGroups]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListDatasetGroups.html
// [CreateSolution]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateSolution.html
// [DescribeDatasetGroup]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeDatasetGroup.html
// [CreateEventTracker]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateEventTracker.html
// [DeleteDatasetGroup]: https://docs.aws.amazon.com/personalize/latest/dg/API_DeleteDatasetGroup.html
