package lightsail

// PutAlarm is generated as a reference stub.
// Executable command wiring lives under cmd/lightsail.go.
//
// Creates or updates an alarm, and associates it with the specified metric.
//
// An alarm is used to monitor a single metric for one of your resources. When a
// metric condition is met, the alarm can notify you by email, SMS text message,
// and a banner displayed on the Amazon Lightsail console. For more information,
// see [Alarms in Amazon Lightsail].
//
// When this action creates an alarm, the alarm state is immediately set to
// INSUFFICIENT_DATA . The alarm is then evaluated and its state is set
// appropriately. Any actions associated with the new state are then executed.
//
// When you update an existing alarm, its state is left unchanged, but the update
// completely overwrites the previous configuration of the alarm. The alarm is then
// evaluated with the updated configuration.
//
// [Alarms in Amazon Lightsail]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-alarms
