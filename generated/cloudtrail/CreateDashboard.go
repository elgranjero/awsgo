package cloudtrail

// CreateDashboard is generated as a reference stub.
// Executable command wiring lives under cmd/cloudtrail.go.
//
// Creates a custom dashboard or the Highlights dashboard.
//
// - Custom dashboards - Custom dashboards allow you to query events in any
// event data store type. You can add up to 10 widgets to a custom dashboard. You
// can manually refresh a custom dashboard, or you can set a refresh schedule.
//
// - Highlights dashboard - You can create the Highlights dashboard to see a
// summary of key user activities and API usage across all your event data stores.
// CloudTrail Lake manages the Highlights dashboard and refreshes the dashboard
// every 6 hours. To create the Highlights dashboard, you must set and enable a
// refresh schedule.
//
// CloudTrail runs queries to populate the dashboard's widgets during a manual or
// scheduled refresh. CloudTrail must be granted permissions to run the StartQuery
// operation on your behalf. To provide permissions, run the PutResourcePolicy
// operation to attach a resource-based policy to each event data store. For more
// information, see [Example: Allow CloudTrail to run queries to populate a dashboard]in the CloudTrail User Guide.
//
// To set a refresh schedule, CloudTrail must be granted permissions to run the
// StartDashboardRefresh operation to refresh the dashboard on your behalf. To
// provide permissions, run the PutResourcePolicy operation to attach a
// resource-based policy to the dashboard. For more information, see [Resource-based policy example for a dashboard]in the
// CloudTrail User Guide.
//
// For more information about dashboards, see [CloudTrail Lake dashboards] in the CloudTrail User Guide.
//
// [CloudTrail Lake dashboards]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/lake-dashboard.html
// [Example: Allow CloudTrail to run queries to populate a dashboard]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/security_iam_resource-based-policy-examples.html#security_iam_resource-based-policy-examples-eds-dashboard
// [Resource-based policy example for a dashboard]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/security_iam_resource-based-policy-examples.html#security_iam_resource-based-policy-examples-dashboards
