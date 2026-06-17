package cloudwatchlogs

// CreateDelivery is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Creates a delivery. A delivery is a connection between a logical delivery
// source and a logical delivery destination that you have already created.
//
// Only some Amazon Web Services services support being configured as a delivery
// source using this operation. These services are listed as Supported [V2
// Permissions] in the table at [Enabling logging from Amazon Web Services services.]
//
// A delivery destination can represent a log group in CloudWatch Logs, an Amazon
// S3 bucket, a delivery stream in Firehose, or X-Ray.
//
// To configure logs delivery between a supported Amazon Web Services service and
// a destination, you must do the following:
//
// - Create a delivery source, which is a logical object that represents the
// resource that is actually sending the logs. For more information, see [PutDeliverySource].
//
// - Create a delivery destination, which is a logical object that represents
// the actual delivery destination. For more information, see [PutDeliveryDestination].
//
// - If you are delivering logs cross-account, you must use [PutDeliveryDestinationPolicy]in the destination
// account to assign an IAM policy to the destination. This policy allows delivery
// to that destination.
//
// - Use CreateDelivery to create a delivery by pairing exactly one delivery
// source and one delivery destination.
//
// You can configure a single delivery source to send logs to multiple
// destinations by creating multiple deliveries. You can also create multiple
// deliveries to configure multiple delivery sources to send logs to the same
// delivery destination.
//
// To update an existing delivery configuration, use [UpdateDeliveryConfiguration].
//
// [PutDeliveryDestination]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliveryDestination.html
// [PutDeliverySource]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliverySource.html
// [Enabling logging from Amazon Web Services services.]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-logs-and-resource-policy.html
// [PutDeliveryDestinationPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliveryDestinationPolicy.html
// [UpdateDeliveryConfiguration]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_UpdateDeliveryConfiguration.html
