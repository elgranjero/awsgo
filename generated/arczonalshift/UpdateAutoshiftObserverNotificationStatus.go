package arczonalshift

// UpdateAutoshiftObserverNotificationStatus is generated as a reference stub.
// Executable command wiring lives under cmd/arczonalshift.go.
//
// Update the status of autoshift observer notification. Autoshift observer
// notification enables you to be notified, through Amazon EventBridge, when there
// is an autoshift event for zonal autoshift.
//
// If the status is ENABLED , ARC includes all autoshift events when you use the
// EventBridge pattern Autoshift In Progress . When the status is DISABLED , ARC
// includes only autoshift events for autoshifts when one or more of your resources
// is included in the autoshift.
//
// For more information, see [Notifications for practice runs and autoshifts] in the Amazon Application Recovery Controller
// Developer Guide.
//
// [Notifications for practice runs and autoshifts]: https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-autoshift.how-it-works.html#ZAShiftNotification
