package keyspaces

// RestoreTable is generated as a reference stub.
// Executable command wiring lives under cmd/keyspaces.go.
//
// Restores the table to the specified point in time within the
// earliest_restorable_timestamp and the current time. For more information about
// restore points, see [Time window for PITR continuous backups]in the Amazon Keyspaces Developer Guide.
//
// Any number of users can execute up to 4 concurrent restores (any type of
// restore) in a given account.
//
// When you restore using point in time recovery, Amazon Keyspaces restores your
// source table's schema and data to the state based on the selected timestamp
// (day:hour:minute:second) to a new table. The Time to Live (TTL) settings are
// also restored to the state based on the selected timestamp.
//
// In addition to the table's schema, data, and TTL settings, RestoreTable
// restores the capacity mode, auto scaling settings, encryption settings, and
// point-in-time recovery settings from the source table. Unlike the table's schema
// data and TTL settings, which are restored based on the selected timestamp, these
// settings are always restored based on the table's settings as of the current
// time or when the table was deleted.
//
// You can also overwrite these settings during restore:
//
// - Read/write capacity mode
//
// - Provisioned throughput capacity units
//
// - Auto scaling settings
//
// - Point-in-time (PITR) settings
//
// - Tags
//
// For more information, see [PITR restore settings] in the Amazon Keyspaces Developer Guide.
//
// Note that the following settings are not restored, and you must configure them
// manually for the new table:
//
// - Identity and Access Management (IAM) policies
//
// - Amazon CloudWatch metrics and alarms
//
// [PITR restore settings]: https://docs.aws.amazon.com/keyspaces/latest/devguide/PointInTimeRecovery_HowItWorks.html#howitworks_backup_settings
// [Time window for PITR continuous backups]: https://docs.aws.amazon.com/keyspaces/latest/devguide/PointInTimeRecovery_HowItWorks.html#howitworks_backup_window
