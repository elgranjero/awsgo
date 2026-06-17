package ec2

// ModifyInstanceEventWindow is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Modifies the specified event window.
//
// You can define either a set of time ranges or a cron expression when modifying
// the event window, but not both.
//
// To modify the targets associated with the event window, use the AssociateInstanceEventWindow and DisassociateInstanceEventWindow API.
//
// If Amazon Web Services has already scheduled an event, modifying an event
// window won't change the time of the scheduled event.
//
// For more information, see [Define event windows for scheduled events] in the Amazon EC2 User Guide.
//
// [Define event windows for scheduled events]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/event-windows.html
