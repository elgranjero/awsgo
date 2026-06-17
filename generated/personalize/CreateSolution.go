package personalize

// CreateSolution is generated as a reference stub.
// Executable command wiring lives under cmd/personalize.go.
//
// By default, all new solutions use automatic training. With automatic training,
// you incur training costs while your solution is active. To avoid unnecessary
// costs, when you are finished you can [update the solution]to turn off automatic training. For
// information about training costs, see [Amazon Personalize pricing].
//
// Creates the configuration for training a model (creating a solution version).
// This configuration includes the recipe to use for model training and optional
// training configuration, such as columns to use in training and feature
// transformation parameters. For more information about configuring a solution,
// see [Creating and configuring a solution].
//
// By default, new solutions use automatic training to create solution versions
// every 7 days. You can change the training frequency. Automatic solution version
// creation starts within one hour after the solution is ACTIVE. If you manually
// create a solution version within the hour, the solution skips the first
// automatic training. For more information, see [Configuring automatic training].
//
// To turn off automatic training, set performAutoTraining to false. If you turn
// off automatic training, you must manually create a solution version by calling
// the [CreateSolutionVersion]operation.
//
// After training starts, you can get the solution version's Amazon Resource Name
// (ARN) with the [ListSolutionVersions]API operation. To get its status, use the [DescribeSolutionVersion].
//
// After training completes you can evaluate model accuracy by calling [GetSolutionMetrics]. When you
// are satisfied with the solution version, you deploy it using [CreateCampaign]. The campaign
// provides recommendations to a client through the [GetRecommendations]API.
//
// Amazon Personalize doesn't support configuring the hpoObjective for solution
// hyperparameter optimization at this time.
//
// # Status
//
// A solution can be in one of the following states:
//
// - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// - DELETE PENDING > DELETE IN_PROGRESS
//
// To get the status of the solution, call [DescribeSolution]. If you use manual training, the
// status must be ACTIVE before you call CreateSolutionVersion .
//
// # Related APIs
//
// [UpdateSolution]
//
// [ListSolutions]
//
// [CreateSolutionVersion]
//
// [DescribeSolution]
//
// [DeleteSolution]
//
// [ListSolutionVersions]
//
// [DescribeSolutionVersion]
//
// [CreateCampaign]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateCampaign.html
// [GetSolutionMetrics]: https://docs.aws.amazon.com/personalize/latest/dg/API_GetSolutionMetrics.html
// [update the solution]: https://docs.aws.amazon.com/personalize/latest/dg/API_UpdateSolution.html
// [ListSolutions]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListSolutions.html
// [Amazon Personalize pricing]: https://aws.amazon.com/personalize/pricing/
// [DescribeSolution]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeSolution.html
// [DescribeSolutionVersion]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeSolutionVersion.html
// [DeleteSolution]: https://docs.aws.amazon.com/personalize/latest/dg/API_DeleteSolution.html
// [UpdateSolution]: https://docs.aws.amazon.com/personalize/latest/dg/API_UpdateSolution.html
// [ListSolutionVersions]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListSolutionVersions.html
// [Creating and configuring a solution]: https://docs.aws.amazon.com/personalize/latest/dg/customizing-solution-config.html
// [GetRecommendations]: https://docs.aws.amazon.com/personalize/latest/dg/API_RS_GetRecommendations.html
// [Configuring automatic training]: https://docs.aws.amazon.com/personalize/latest/dg/solution-config-auto-training.html
// [CreateSolutionVersion]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateSolutionVersion.html
