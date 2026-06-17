package cloudwatchlogs

// DeleteDeliverySource is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Deletes a delivery source. A delivery is a connection between a logical
// delivery source and a logical delivery destination.
//
// You can't delete a delivery source if any current deliveries are associated
// with it. To find whether any deliveries are associated with this delivery
// source, use the [DescribeDeliveries]operation and check the deliverySourceName field in the results.
//
// [DescribeDeliveries]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeDeliveries.html
