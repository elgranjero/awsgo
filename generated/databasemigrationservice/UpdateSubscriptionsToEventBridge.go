package databasemigrationservice

// UpdateSubscriptionsToEventBridge is generated as a reference stub.
// Executable command wiring lives under cmd/databasemigrationservice.go.
//
// Migrates 10 active and enabled Amazon SNS subscriptions at a time and converts
// them to corresponding Amazon EventBridge rules. By default, this operation
// migrates subscriptions only when all your replication instance versions are
// 3.4.5 or higher. If any replication instances are from versions earlier than
// 3.4.5, the operation raises an error and tells you to upgrade these instances to
// version 3.4.5 or higher. To enable migration regardless of version, set the
// Force option to true. However, if you don't upgrade instances earlier than
// version 3.4.5, some types of events might not be available when you use Amazon
// EventBridge.
//
// To call this operation, make sure that you have certain permissions added to
// your user account. For more information, see [Migrating event subscriptions to Amazon EventBridge]in the Amazon Web Services
// Database Migration Service User Guide.
//
// [Migrating event subscriptions to Amazon EventBridge]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Events.html#CHAP_Events-migrate-to-eventbridge
