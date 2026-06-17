package personalize

// UpdateRecommender is generated as a reference stub.
// Executable command wiring lives under cmd/personalize.go.
//
// Updates the recommender to modify the recommender configuration. If you update
// the recommender to modify the columns used in training, Amazon Personalize
// automatically starts a full retraining of the models backing your recommender.
// While the update completes, you can still get recommendations from the
// recommender. The recommender uses the previous configuration until the update
// completes. To track the status of this update, use the latestRecommenderUpdate
// returned in the [DescribeRecommender]operation.
//
// [DescribeRecommender]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeRecommender.html
