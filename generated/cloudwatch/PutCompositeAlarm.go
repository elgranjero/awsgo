package cloudwatch

// PutCompositeAlarm is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatch.go.
//
// Creates or updates a composite alarm. When you create a composite alarm, you
// specify a rule expression for the alarm that takes into account the alarm states
// of other alarms that you have created. The composite alarm goes into ALARM state
// only if all conditions of the rule are met.
//
// The alarms specified in a composite alarm's rule expression can include metric
// alarms and other composite alarms. The rule expression of a composite alarm can
// include as many as 100 underlying alarms. Any single alarm can be included in
// the rule expressions of as many as 150 composite alarms.
//
// Using composite alarms can reduce alarm noise. You can create multiple metric
// alarms, and also create a composite alarm and set up alerts only for the
// composite alarm. For example, you could create a composite alarm that goes into
// ALARM state only when more than one of the underlying metric alarms are in ALARM
// state.
//
// Composite alarms can take the following actions:
//
// - Notify Amazon SNS topics.
//
// - Invoke Lambda functions.
//
// - Create OpsItems in Systems Manager Ops Center.
//
// - Create incidents in Systems Manager Incident Manager.
//
// It is possible to create a loop or cycle of composite alarms, where composite
// alarm A depends on composite alarm B, and composite alarm B also depends on
// composite alarm A. In this scenario, you can't delete any composite alarm that
// is part of the cycle because there is always still a composite alarm that
// depends on that alarm that you want to delete.
//
// To get out of such a situation, you must break the cycle by changing the rule
// of one of the composite alarms in the cycle to remove a dependency that creates
// the cycle. The simplest change to make to break a cycle is to change the
// AlarmRule of one of the alarms to false .
//
// Additionally, the evaluation of composite alarms stops if CloudWatch detects a
// cycle in the evaluation path.
//
// When this operation creates an alarm, the alarm state is immediately set to
// INSUFFICIENT_DATA . The alarm is then evaluated and its state is set
// appropriately. Any actions associated with the new state are then executed. For
// a composite alarm, this initial time after creation is the only time that the
// alarm can be in INSUFFICIENT_DATA state.
//
// When you update an existing alarm, its state is left unchanged, but the update
// completely overwrites the previous configuration of the alarm.
//
// To use this operation, you must be signed on with the
// cloudwatch:PutCompositeAlarm permission that is scoped to * . You can't create a
// composite alarms if your cloudwatch:PutCompositeAlarm permission has a narrower
// scope.
//
// If you are an IAM user, you must have iam:CreateServiceLinkedRole to create a
// composite alarm that has Systems Manager OpsItem actions.
