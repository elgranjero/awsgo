package cloudwatchlogs

// PutDestination is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Creates or updates a destination. This operation is used only to create
// destinations for cross-account subscriptions.
//
// A destination encapsulates a physical resource (such as an Amazon Kinesis
// stream). With a destination, you can subscribe to a real-time stream of log
// events for a different account, ingested using [PutLogEvents].
//
// Through an access policy, a destination controls what is written to it. By
// default, PutDestination does not set any access policy with the destination,
// which means a cross-account user cannot call [PutSubscriptionFilter]against this destination. To
// enable this, the destination owner must call [PutDestinationPolicy]after PutDestination .
//
// To perform a PutDestination operation, you must also have the iam:PassRole
// permission.
//
// [PutSubscriptionFilter]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutSubscriptionFilter.html
// [PutLogEvents]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutLogEvents.html
// [PutDestinationPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDestinationPolicy.html
