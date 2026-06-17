package cloudwatch

// DeleteAlarmMuteRule is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatch.go.
//
// Deletes a specific alarm mute rule.
//
// When you delete a mute rule, any alarms that are currently being muted by that
// rule are immediately unmuted. If those alarms are in an ALARM state, their
// configured actions will trigger.
//
// This operation is idempotent. If you delete a mute rule that does not exist,
// the operation succeeds without returning an error.
//
// # Permissions
//
// To delete a mute rule, you need the cloudwatch:DeleteAlarmMuteRule permission
// on the alarm mute rule resource.
