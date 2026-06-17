package ec2

// CreateInstanceEventWindow is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates an event window in which scheduled events for the associated Amazon EC2
// instances can run.
//
// You can define either a set of time ranges or a cron expression when creating
// the event window, but not both. All event window times are in UTC.
//
// You can create up to 200 event windows per Amazon Web Services Region.
//
// When you create the event window, targets (instance IDs, Dedicated Host IDs, or
// tags) are not yet associated with it. To ensure that the event window can be
// used, you must associate one or more targets with it by using the AssociateInstanceEventWindowAPI.
//
// Event windows are applicable only for scheduled events that stop, reboot, or
// terminate instances.
//
// Event windows are not applicable for:
//
// - Expedited scheduled events and network maintenance events.
//
// - Unscheduled maintenance such as AutoRecovery and unplanned reboots.
//
// For more information, see [Define event windows for scheduled events] in the Amazon EC2 User Guide.
//
// [Define event windows for scheduled events]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/event-windows.html
