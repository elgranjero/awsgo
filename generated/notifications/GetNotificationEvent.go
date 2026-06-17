package notifications

// GetNotificationEvent is generated as a reference stub.
// Executable command wiring lives under cmd/notifications.go.
//
// Returns a specified NotificationEvent .
//
// User Notifications stores notifications in the individual Regions you register
// as notification hubs and the Region of the source event rule.
// GetNotificationEvent only returns notifications stored in the same Region in
// which the action is called. User Notifications doesn't backfill notifications to
// new Regions selected as notification hubs. For this reason, we recommend that
// you make calls in your oldest registered notification hub. For more information,
// see [Notification hubs]in the Amazon Web Services User Notifications User Guide.
//
// [Notification hubs]: https://docs.aws.amazon.com/notifications/latest/userguide/notification-hubs.html
