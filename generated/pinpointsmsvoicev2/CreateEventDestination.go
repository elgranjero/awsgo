package pinpointsmsvoicev2

// CreateEventDestination is generated as a reference stub.
// Executable command wiring lives under cmd/pinpointsmsvoicev2.go.
//
// Creates a new event destination in a configuration set.
//
// An event destination is a location where you send message events. The event
// options are Amazon CloudWatch, Amazon Data Firehose, or Amazon SNS. For example,
// when a message is delivered successfully, you can send information about that
// event to an event destination, or send notifications to endpoints that are
// subscribed to an Amazon SNS topic.
//
// You can only create one event destination at a time. You must provide a value
// for a single event destination using either CloudWatchLogsDestination ,
// KinesisFirehoseDestination or SnsDestination . If an event destination isn't
// provided then an exception is returned.
//
// Each configuration set can contain between 0 and 5 event destinations. Each
// event destination can contain a reference to a single destination, such as a
// CloudWatch or Firehose destination.
