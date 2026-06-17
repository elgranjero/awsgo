package personalizeruntime

// GetRecommendations is generated as a reference stub.
// Executable command wiring lives under cmd/personalizeruntime.go.
//
// Returns a list of recommended items. For campaigns, the campaign's Amazon
// Resource Name (ARN) is required and the required user and item input depends on
// the recipe type used to create the solution backing the campaign as follows:
//
// - USER_PERSONALIZATION - userId required, itemId not used
//
// - RELATED_ITEMS - itemId required, userId not used
//
// Campaigns that are backed by a solution created using a recipe of type
// PERSONALIZED_RANKING use the API.
//
// For recommenders, the recommender's ARN is required and the required item and
// user input depends on the use case (domain-based recipe) backing the
// recommender. For information on use case requirements see [Choosing recommender use cases].
//
// [Choosing recommender use cases]: https://docs.aws.amazon.com/personalize/latest/dg/domain-use-cases.html
