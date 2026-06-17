package cloudwatch

// PutAlarmMuteRule is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatch.go.
//
// Creates or updates an alarm mute rule.
//
// Alarm mute rules automatically mute alarm actions during predefined time
// windows. When a mute rule is active, targeted alarms continue to evaluate
// metrics and transition between states, but their configured actions (such as
// Amazon SNS notifications or Auto Scaling actions) are muted.
//
// You can create mute rules with recurring schedules using cron expressions or
// one-time mute windows using at expressions. Each mute rule can target up to 100
// specific alarms by name.
//
// If you specify a rule name that already exists, this operation updates the
// existing rule with the new configuration.
//
// # Permissions
//
// To create or update a mute rule, you must have the cloudwatch:PutAlarmMuteRule
// permission on two types of resources: the alarm mute rule resource itself, and
// each alarm that the rule targets.
//
// For example, If you want to allow a user to create mute rules that target only
// specific alarms named "WebServerCPUAlarm" and "DatabaseConnectionAlarm", you
// would create an IAM policy with one statement granting
// cloudwatch:PutAlarmMuteRule on the alarm mute rule resource (
// arn:aws:cloudwatch:[REGION]:123456789012:alarm-mute:* ), and another statement
// granting cloudwatch:PutAlarmMuteRule on the targeted alarm resources (
// arn:aws:cloudwatch:[REGION]:123456789012:alarm:WebServerCPUAlarm and
// arn:aws:cloudwatch:[REGION]:123456789012:alarm:DatabaseConnectionAlarm ).
//
// You can also use IAM policy conditions to allow targeting alarms based on
// resource tags. For example, you can restrict users to create/update mute rules
// to only target alarms that have a specific tag key-value pair, such as
// Team=TeamA .
