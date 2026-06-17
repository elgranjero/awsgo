package personalize

// DescribeRecommender is generated as a reference stub.
// Executable command wiring lives under cmd/personalize.go.
//
// Describes the given recommender, including its status.
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
// When the status is CREATE FAILED , the response includes the failureReason key,
// which describes why.
//
// The modelMetrics key is null when the recommender is being created or deleted.
//
// For more information on recommenders, see [CreateRecommender].
//
// [CreateRecommender]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateRecommender.html
