package quicksight

// StartDashboardSnapshotJobSchedule is generated as a reference stub.
// Executable command wiring lives under cmd/quicksight.go.
//
// Starts an asynchronous job that runs an existing dashboard schedule and sends
// the dashboard snapshot through email.
//
// Only one job can run simultaneously in a given schedule. Repeated requests are
// skipped with a 202 HTTP status code.
//
// For more information, see [Scheduling and sending Amazon Quick Sight reports by email] and [Configuring email report settings for a Amazon Quick Sight dashboard] in the Amazon Quick Sight User Guide.
//
// [Configuring email report settings for a Amazon Quick Sight dashboard]: https://docs.aws.amazon.com/quicksight/latest/user/email-reports-from-dashboard.html
// [Scheduling and sending Amazon Quick Sight reports by email]: https://docs.aws.amazon.com/quicksight/latest/user/sending-reports.html
