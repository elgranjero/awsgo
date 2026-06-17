package cloudwatch

// GetAlarmMuteRule is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatch.go.
//
// Retrieves details for a specific alarm mute rule.
//
// This operation returns complete information about the mute rule, including its
// configuration, status, targeted alarms, and metadata.
//
// The returned status indicates the current state of the mute rule:
//
// - SCHEDULED: The mute rule is configured and will become active in the future
//
// - ACTIVE: The mute rule is currently muting alarm actions
//
// - EXPIRED: The mute rule has passed its expiration date and will no longer
// become active
//
// # Permissions
//
// To retrieve details for a mute rule, you need the cloudwatch:GetAlarmMuteRule
// permission on the alarm mute rule resource.
