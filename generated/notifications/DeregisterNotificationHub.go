package notifications

// DeregisterNotificationHub is generated as a reference stub.
// Executable command wiring lives under cmd/notifications.go.
//
// Deregisters a NotificationConfiguration in the specified Region.
//
// You can't deregister the last NotificationHub in the account. NotificationEvents
// stored in the deregistered NotificationConfiguration are no longer be visible.
// Recreating a new NotificationConfiguration in the same Region restores access
// to those NotificationEvents .
