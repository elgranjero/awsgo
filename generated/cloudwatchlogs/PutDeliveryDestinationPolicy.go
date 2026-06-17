package cloudwatchlogs

// PutDeliveryDestinationPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Creates and assigns an IAM policy that grants permissions to CloudWatch Logs to
// deliver logs cross-account to a specified destination in this account. To
// configure the delivery of logs from an Amazon Web Services service in another
// account to a logs delivery destination in the current account, you must do the
// following:
//
// - Create a delivery source, which is a logical object that represents the
// resource that is actually sending the logs. For more information, see [PutDeliverySource].
//
// - Create a delivery destination, which is a logical object that represents
// the actual delivery destination. For more information, see [PutDeliveryDestination].
//
// - Use this operation in the destination account to assign an IAM policy to
// the destination. This policy allows delivery to that destination.
//
// - Create a delivery by pairing exactly one delivery source and one delivery
// destination. For more information, see [CreateDelivery].
//
// Only some Amazon Web Services services support being configured as a delivery
// source. These services are listed as Supported [V2 Permissions] in the table at [Enabling logging from Amazon Web Services services.]
//
// The contents of the policy must include two statements. One statement enables
// general logs delivery, and the other allows delivery to the chosen destination.
// See the examples for the needed policies.
//
// [PutDeliveryDestination]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliveryDestination.html
// [PutDeliverySource]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliverySource.html
// [Enabling logging from Amazon Web Services services.]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-logs-and-resource-policy.html
// [CreateDelivery]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CreateDelivery.html
