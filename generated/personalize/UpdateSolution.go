package personalize

// UpdateSolution is generated as a reference stub.
// Executable command wiring lives under cmd/personalize.go.
//
// Updates an Amazon Personalize solution to use a different automatic training
// configuration. When you update a solution, you can change whether the solution
// uses automatic training, and you can change the training frequency. For more
// information about updating a solution, see [Updating a solution].
//
// A solution update can be in one of the following states:
//
// CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// To get the status of a solution update, call the [DescribeSolution] API operation and find the
// status in the latestSolutionUpdate .
//
// [Updating a solution]: https://docs.aws.amazon.com/personalize/latest/dg/updating-solution.html
// [DescribeSolution]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeSolution.html
