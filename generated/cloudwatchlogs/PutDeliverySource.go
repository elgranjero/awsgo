package cloudwatchlogs

// PutDeliverySource is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Creates or updates a logical delivery source. A delivery source represents an
// Amazon Web Services resource that sends logs to an logs delivery destination.
// The destination can be CloudWatch Logs, Amazon S3, Firehose or X-Ray for sending
// traces.
//
// To configure logs delivery between a delivery destination and an Amazon Web
// Services service that is supported as a delivery source, you must do the
// following:
//
// - Use PutDeliverySource to create a delivery source, which is a logical object
// that represents the resource that is actually sending the logs.
//
// - Use PutDeliveryDestination to create a delivery destination, which is a
// logical object that represents the actual delivery destination. For more
// information, see [PutDeliveryDestination].
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
// If you use this operation to update an existing delivery source, all the
// current delivery source parameters are overwritten with the new parameter values
// that you specify.
//
// [PutDeliveryDestination]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliveryDestination.html
// [Enabling logging from Amazon Web Services services.]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-logs-and-resource-policy.html
// [CreateDelivery]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CreateDelivery.html
// [PutDeliveryDestinationPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliveryDestinationPolicy.html
