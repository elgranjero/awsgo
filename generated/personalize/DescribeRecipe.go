package personalize

// DescribeRecipe is generated as a reference stub.
// Executable command wiring lives under cmd/personalize.go.
//
// Describes a recipe.
//
// A recipe contains three items:
//
// - An algorithm that trains a model.
//
// - Hyperparameters that govern the training.
//
// - Feature transformation information for modifying the input data before
// training.
//
// Amazon Personalize provides a set of predefined recipes. You specify a recipe
// when you create a solution with the [CreateSolution]API. CreateSolution trains a model by using
// the algorithm in the specified recipe and a training dataset. The solution, when
// deployed as a campaign, can provide recommendations using the [GetRecommendations]API.
//
// [CreateSolution]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateSolution.html
// [GetRecommendations]: https://docs.aws.amazon.com/personalize/latest/dg/API_RS_GetRecommendations.html
