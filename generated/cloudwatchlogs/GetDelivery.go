package cloudwatchlogs

// GetDelivery is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Returns complete information about one logical delivery. A delivery is a
// connection between a [delivery source]and a [delivery destination].
//
// A delivery source represents an Amazon Web Services resource that sends logs to
// an logs delivery destination. The destination can be CloudWatch Logs, Amazon S3,
// or Firehose. Only some Amazon Web Services services support being configured as
// a delivery source. These services are listed in [Enable logging from Amazon Web Services services.]
//
// You need to specify the delivery id in this operation. You can find the IDs of
// the deliveries in your account with the [DescribeDeliveries]operation.
//
// [delivery destination]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliveryDestination.html
// [delivery source]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliverySource.html
// [Enable logging from Amazon Web Services services.]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-logs-and-resource-policy.html
// [DescribeDeliveries]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeDeliveries.html
