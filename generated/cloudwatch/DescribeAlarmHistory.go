package cloudwatch

// DescribeAlarmHistory is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatch.go.
//
// Retrieves the history for the specified alarm. You can filter the results by
// date range or item type. If an alarm name is not specified, the histories for
// either all metric alarms or all composite alarms are returned.
//
// CloudWatch retains the history of an alarm even if you delete the alarm.
//
// To use this operation and return information about a composite alarm, you must
// be signed on with the cloudwatch:DescribeAlarmHistory permission that is scoped
// to * . You can't return information about composite alarms if your
// cloudwatch:DescribeAlarmHistory permission has a narrower scope.
