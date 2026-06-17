package cloudwatchlogs

// PutRetentionPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Sets the retention of the specified log group. With a retention policy, you can
// configure the number of days for which to retain log events in the specified log
// group.
//
// CloudWatch Logs doesn't immediately delete log events when they reach their
// retention setting. It typically takes up to 72 hours after that before log
// events are deleted, but in rare situations might take longer.
//
// To illustrate, imagine that you change a log group to have a longer retention
// setting when it contains log events that are past the expiration date, but
// haven't been deleted. Those log events will take up to 72 hours to be deleted
// after the new retention date is reached. To make sure that log data is deleted
// permanently, keep a log group at its lower retention setting until 72 hours
// after the previous retention period ends. Alternatively, wait to change the
// retention setting until you confirm that the earlier log events are deleted.
//
// When log events reach their retention setting they are marked for deletion.
// After they are marked for deletion, they do not add to your archival storage
// costs anymore, even if they are not actually deleted until later. These log
// events marked for deletion are also not included when you use an API to retrieve
// the storedBytes value to see how many bytes a log group is storing.
