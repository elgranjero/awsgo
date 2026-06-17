package configservice

// DeliverConfigSnapshot is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Schedules delivery of a configuration snapshot to the Amazon S3 bucket in the
// specified delivery channel. After the delivery has started, Config sends the
// following notifications using an Amazon SNS topic that you have specified.
//
// - Notification of the start of the delivery.
//
// - Notification of the completion of the delivery, if the delivery was
// successfully completed.
//
// - Notification of delivery failure, if the delivery failed.
