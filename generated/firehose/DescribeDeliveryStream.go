package firehose

// DescribeDeliveryStream is generated as a reference stub.
// Executable command wiring lives under cmd/firehose.go.
//
// Describes the specified Firehose stream and its status. For example, after your
// Firehose stream is created, call DescribeDeliveryStream to see whether the
// Firehose stream is ACTIVE and therefore ready for data to be sent to it.
//
// If the status of a Firehose stream is CREATING_FAILED , this status doesn't
// change, and you can't invoke CreateDeliveryStreamagain on it. However, you can invoke the DeleteDeliveryStream
// operation to delete it. If the status is DELETING_FAILED , you can force
// deletion by invoking DeleteDeliveryStreamagain but with DeleteDeliveryStreamInput$AllowForceDelete set to true.
