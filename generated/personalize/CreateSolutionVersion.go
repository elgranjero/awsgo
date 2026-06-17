package personalize

// CreateSolutionVersion is generated as a reference stub.
// Executable command wiring lives under cmd/personalize.go.
//
// Trains or retrains an active solution in a Custom dataset group. A solution is
// created using the [CreateSolution]operation and must be in the ACTIVE state before calling
// CreateSolutionVersion . A new version of the solution is created every time you
// call this operation.
//
// # Status
//
// A solution version can be in one of the following states:
//
// - CREATE PENDING
//
// - CREATE IN_PROGRESS
//
// - ACTIVE
//
// - CREATE FAILED
//
// - CREATE STOPPING
//
// - CREATE STOPPED
//
// To get the status of the version, call [DescribeSolutionVersion]. Wait until the status shows as ACTIVE
// before calling CreateCampaign .
//
// If the status shows as CREATE FAILED, the response includes a failureReason
// key, which describes why the job failed.
//
// # Related APIs
//
// [ListSolutionVersions]
//
// [DescribeSolutionVersion]
//
// [ListSolutions]
//
// [CreateSolution]
//
// [DescribeSolution]
//
// [DeleteSolution]
//
// [DescribeSolutionVersion]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeSolutionVersion.html
// [DeleteSolution]: https://docs.aws.amazon.com/personalize/latest/dg/API_DeleteSolution.html
// [CreateSolution]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateSolution.html
// [ListSolutionVersions]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListSolutionVersions.html
// [ListSolutions]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListSolutions.html
// [DescribeSolution]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeSolution.html
