package cloudwatch

// PutMetricAlarm is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatch.go.
//
// Creates or updates an alarm and associates it with the specified metric, metric
// math expression, anomaly detection model, or Metrics Insights query. For more
// information about using a Metrics Insights query for an alarm, see [Create alarms on Metrics Insights queries].
//
// Alarms based on anomaly detection models cannot have Auto Scaling actions.
//
// When this operation creates an alarm, the alarm state is immediately set to
// INSUFFICIENT_DATA . The alarm is then evaluated and its state is set
// appropriately. Any actions associated with the new state are then executed.
//
// When you update an existing alarm, its state is left unchanged, but the update
// completely overwrites the previous configuration of the alarm.
//
// If you are an IAM user, you must have Amazon EC2 permissions for some alarm
// operations:
//
// - The iam:CreateServiceLinkedRole permission for all alarms with EC2 actions
//
// - The iam:CreateServiceLinkedRole permissions to create an alarm with Systems
// Manager OpsItem or response plan actions.
//
// The first time you create an alarm in the Amazon Web Services Management
// Console, the CLI, or by using the PutMetricAlarm API, CloudWatch creates the
// necessary service-linked role for you. The service-linked roles are called
// AWSServiceRoleForCloudWatchEvents and
// AWSServiceRoleForCloudWatchAlarms_ActionSSM . For more information, see [Amazon Web Services service-linked role].
//
// Each PutMetricAlarm action has a maximum uncompressed payload of 120 KB.
//
// # Cross-account alarms
//
// You can set an alarm on metrics in the current account, or in another account.
// To create a cross-account alarm that watches a metric in a different account,
// you must have completed the following pre-requisites:
//
// - The account where the metrics are located (the sharing account) must
// already have a sharing role named CloudWatch-CrossAccountSharingRole. If it does
// not already have this role, you must create it using the instructions in Set up
// a sharing account in [Cross-account cross-Region CloudWatch console]. The policy for that role must grant access to the ID
// of the account where you are creating the alarm.
//
// - The account where you are creating the alarm (the monitoring account) must
// already have a service-linked role named AWSServiceRoleForCloudWatchCrossAccount
// to allow CloudWatch to assume the sharing role in the sharing account. If it
// does not, you must create it following the directions in Set up a monitoring
// account in [Cross-account cross-Region CloudWatch console].
//
// [Amazon Web Services service-linked role]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts.html#iam-term-service-linked-role
// [Create alarms on Metrics Insights queries]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Create_Metrics_Insights_Alarm.html
// [Cross-account cross-Region CloudWatch console]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Cross-Account-Cross-Region.html#enable-cross-account-cross-Region
