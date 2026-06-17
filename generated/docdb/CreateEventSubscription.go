package docdb

// CreateEventSubscription is generated as a reference stub.
// Executable command wiring lives under cmd/docdb.go.
//
// Creates an Amazon DocumentDB event notification subscription. This action
// requires a topic Amazon Resource Name (ARN) created by using the Amazon
// DocumentDB console, the Amazon SNS console, or the Amazon SNS API. To obtain an
// ARN with Amazon SNS, you must create a topic in Amazon SNS and subscribe to the
// topic. The ARN is displayed in the Amazon SNS console.
//
// You can specify the type of source ( SourceType ) that you want to be notified
// of. You can also provide a list of Amazon DocumentDB sources ( SourceIds ) that
// trigger the events, and you can provide a list of event categories (
// EventCategories ) for events that you want to be notified of. For example, you
// can specify SourceType = db-instance , SourceIds = mydbinstance1, mydbinstance2
// and EventCategories = Availability, Backup .
//
// If you specify both the SourceType and SourceIds (such as SourceType =
// db-instance and SourceIdentifier = myDBInstance1 ), you are notified of all the
// db-instance events for the specified source. If you specify a SourceType but do
// not specify a SourceIdentifier , you receive notice of the events for that
// source type for all your Amazon DocumentDB sources. If you do not specify either
// the SourceType or the SourceIdentifier , you are notified of events generated
// from all Amazon DocumentDB sources belonging to your customer account.
