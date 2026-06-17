package mturk

// UpdateNotificationSettings is generated as a reference stub.
// Executable command wiring lives under cmd/mturk.go.
//
// The UpdateNotificationSettings operation creates, updates, disables or
//
// re-enables notifications for a HIT type. If you call the
// UpdateNotificationSettings operation for a HIT type that already has a
// notification specification, the operation replaces the old specification with a
// new one. You can call the UpdateNotificationSettings operation to enable or
// disable notifications for the HIT type, without having to modify the
// notification specification itself by providing updates to the Active status
// without specifying a new notification specification. To change the Active status
// of a HIT type's notifications, the HIT type must already have a notification
// specification, or one must be provided in the same call to
// UpdateNotificationSettings .
