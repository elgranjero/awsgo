package rds

// CreateEventSubscription is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Creates an RDS event notification subscription. This operation requires a topic
// Amazon Resource Name (ARN) created by either the RDS console, the SNS console,
// or the SNS API. To obtain an ARN with SNS, you must create a topic in Amazon SNS
// and subscribe to the topic. The ARN is displayed in the SNS console.
//
// You can specify the type of source ( SourceType ) that you want to be notified
// of and provide a list of RDS sources ( SourceIds ) that triggers the events. You
// can also provide a list of event categories ( EventCategories ) for events that
// you want to be notified of. For example, you can specify SourceType =
// db-instance , SourceIds = mydbinstance1 , mydbinstance2 and EventCategories =
// Availability , Backup .
//
// If you specify both the SourceType and SourceIds , such as SourceType =
// db-instance and SourceIds = myDBInstance1 , you are notified of all the
// db-instance events for the specified source. If you specify a SourceType but do
// not specify SourceIds , you receive notice of the events for that source type
// for all your RDS sources. If you don't specify either the SourceType or the
// SourceIds , you are notified of events generated from all RDS sources belonging
// to your customer account.
//
// For more information about subscribing to an event for RDS DB engines, see [Subscribing to Amazon RDS event notification] in
// the Amazon RDS User Guide.
//
// For more information about subscribing to an event for Aurora DB engines, see [Subscribing to Amazon RDS event notification]
// in the Amazon Aurora User Guide.
//
// [Subscribing to Amazon RDS event notification]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Events.Subscribing.html
