package databasemigrationservice

// CreateEventSubscription is generated as a reference stub.
// Executable command wiring lives under cmd/databasemigrationservice.go.
//
// Creates an DMS event notification subscription.
//
// You can specify the type of source ( SourceType ) you want to be notified of,
// provide a list of DMS source IDs ( SourceIds ) that triggers the events, and
// provide a list of event categories ( EventCategories ) for events you want to be
// notified of. If you specify both the SourceType and SourceIds , such as
// SourceType = replication-instance and SourceIdentifier = my-replinstance , you
// will be notified of all the replication instance events for the specified
// source. If you specify a SourceType but don't specify a SourceIdentifier , you
// receive notice of the events for that source type for all your DMS sources. If
// you don't specify either SourceType nor SourceIdentifier , you will be notified
// of events generated from all DMS sources belonging to your customer account.
//
// For more information about DMS events, see [Working with Events and Notifications] in the Database Migration Service
// User Guide.
//
// [Working with Events and Notifications]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Events.html
