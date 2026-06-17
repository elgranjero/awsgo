package cloudwatchlogs

// PutDeliveryDestination is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Creates or updates a logical delivery destination. A delivery destination is an
// Amazon Web Services resource that represents an Amazon Web Services service that
// logs can be sent to. CloudWatch Logs, Amazon S3, and Firehose are supported as
// logs delivery destinations and X-Ray as the trace delivery destination.
//
// To configure logs delivery between a supported Amazon Web Services service and
// a destination, you must do the following:
//
// - Create a delivery source, which is a logical object that represents the
// resource that is actually sending the logs. For more information, see [PutDeliverySource].
//
// - Use PutDeliveryDestination to create a delivery destination in the same
// account of the actual delivery destination. The delivery destination that you
// create is a logical object that represents the actual delivery destination.
//
// - If you are delivering logs cross-account, you must use [PutDeliveryDestinationPolicy]in the destination
// account to assign an IAM policy to the destination. This policy allows delivery
// to that destination.
//
// - Use CreateDelivery to create a delivery by pairing exactly one delivery
// source and one delivery destination. For more information, see [CreateDelivery].
//
// You can configure a single delivery source to send logs to multiple
// destinations by creating multiple deliveries. You can also create multiple
// deliveries to configure multiple delivery sources to send logs to the same
// delivery destination.
//
// Only some Amazon Web Services services support being configured as a delivery
// source. These services are listed as Supported [V2 Permissions] in the table at [Enabling logging from Amazon Web Services services.]
//
// If you use this operation to update an existing delivery destination, all the
// current delivery destination parameters are overwritten with the new parameter
// values that you specify.
//
// [PutDeliverySource]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliverySource.html
// [Enabling logging from Amazon Web Services services.]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-logs-and-resource-policy.html
// [CreateDelivery]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CreateDelivery.html
// [PutDeliveryDestinationPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliveryDestinationPolicy.html
