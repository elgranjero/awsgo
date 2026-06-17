package personalize

// CreateBatchInferenceJob is generated as a reference stub.
// Executable command wiring lives under cmd/personalize.go.
//
// Generates batch recommendations based on a list of items or users stored in
// Amazon S3 and exports the recommendations to an Amazon S3 bucket.
//
// To generate batch recommendations, specify the ARN of a solution version and an
// Amazon S3 URI for the input and output data. For user personalization, popular
// items, and personalized ranking solutions, the batch inference job generates a
// list of recommended items for each user ID in the input file. For related items
// solutions, the job generates a list of recommended items for each item ID in the
// input file.
//
// For more information, see [Creating a batch inference job].
//
// If you use the Similar-Items recipe, Amazon Personalize can add descriptive
// themes to batch recommendations. To generate themes, set the job's mode to
// THEME_GENERATION and specify the name of the field that contains item names in
// the input data.
//
// For more information about generating themes, see [Batch recommendations with themes from Content Generator].
//
// You can't get batch recommendations with the Trending-Now or Next-Best-Action
// recipes.
//
// [Creating a batch inference job]: https://docs.aws.amazon.com/personalize/latest/dg/getting-batch-recommendations.html
// [Batch recommendations with themes from Content Generator]: https://docs.aws.amazon.com/personalize/latest/dg/themed-batch-recommendations.html
