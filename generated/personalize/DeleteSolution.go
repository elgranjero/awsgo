package personalize

// DeleteSolution is generated as a reference stub.
// Executable command wiring lives under cmd/personalize.go.
//
// Deletes all versions of a solution and the Solution object itself. Before
// deleting a solution, you must delete all campaigns based on the solution. To
// determine what campaigns are using the solution, call [ListCampaigns]and supply the Amazon
// Resource Name (ARN) of the solution. You can't delete a solution if an
// associated SolutionVersion is in the CREATE PENDING or IN PROGRESS state. For
// more information on solutions, see [CreateSolution].
//
// [CreateSolution]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateSolution.html
// [ListCampaigns]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListCampaigns.html
