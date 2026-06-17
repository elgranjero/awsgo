package personalize

// CreateRecommender is generated as a reference stub.
// Executable command wiring lives under cmd/personalize.go.
//
// Creates a recommender with the recipe (a Domain dataset group use case) you
// specify. You create recommenders for a Domain dataset group and specify the
// recommender's Amazon Resource Name (ARN) when you make a [GetRecommendations]request.
//
// # Minimum recommendation requests per second
//
// A high minRecommendationRequestsPerSecond will increase your bill. We recommend
// starting with 1 for minRecommendationRequestsPerSecond (the default). Track
// your usage using Amazon CloudWatch metrics, and increase the
// minRecommendationRequestsPerSecond as necessary.
//
// When you create a recommender, you can configure the recommender's minimum
// recommendation requests per second. The minimum recommendation requests per
// second ( minRecommendationRequestsPerSecond ) specifies the baseline
// recommendation request throughput provisioned by Amazon Personalize. The default
// minRecommendationRequestsPerSecond is 1 . A recommendation request is a single
// GetRecommendations operation. Request throughput is measured in requests per
// second and Amazon Personalize uses your requests per second to derive your
// requests per hour and the price of your recommender usage.
//
// If your requests per second increases beyond minRecommendationRequestsPerSecond
// , Amazon Personalize auto-scales the provisioned capacity up and down, but never
// below minRecommendationRequestsPerSecond . There's a short time delay while the
// capacity is increased that might cause loss of requests.
//
// Your bill is the greater of either the minimum requests per hour (based on
// minRecommendationRequestsPerSecond) or the actual number of requests. The actual
// request throughput used is calculated as the average requests/second within a
// one-hour window.
//
// We recommend starting with the default minRecommendationRequestsPerSecond ,
// track your usage using Amazon CloudWatch metrics, and then increase the
// minRecommendationRequestsPerSecond as necessary.
//
// # Status
//
// A recommender can be in one of the following states:
//
// - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// - STOP PENDING > STOP IN_PROGRESS > INACTIVE > START PENDING > START
// IN_PROGRESS > ACTIVE
//
// - DELETE PENDING > DELETE IN_PROGRESS
//
// To get the recommender status, call [DescribeRecommender].
//
// Wait until the status of the recommender is ACTIVE before asking the
// recommender for recommendations.
//
// # Related APIs
//
// [ListRecommenders]
//
// [DescribeRecommender]
//
// [UpdateRecommender]
//
// [DeleteRecommender]
//
// [ListRecommenders]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListRecommenders.html
// [GetRecommendations]: https://docs.aws.amazon.com/personalize/latest/dg/API_RS_GetRecommendations.html
// [UpdateRecommender]: https://docs.aws.amazon.com/personalize/latest/dg/API_UpdateRecommender.html
// [DeleteRecommender]: https://docs.aws.amazon.com/personalize/latest/dg/API_DeleteRecommender.html
// [DescribeRecommender]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeRecommender.html
