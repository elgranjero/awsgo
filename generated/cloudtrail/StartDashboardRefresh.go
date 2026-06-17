package cloudtrail

// StartDashboardRefresh is generated as a reference stub.
// Executable command wiring lives under cmd/cloudtrail.go.
//
// Starts a refresh of the specified dashboard.
//
// Each time a dashboard is refreshed, CloudTrail runs queries to populate the
// dashboard's widgets. CloudTrail must be granted permissions to run the
// StartQuery operation on your behalf. To provide permissions, run the
// PutResourcePolicy operation to attach a resource-based policy to each event data
// store. For more information, see [Example: Allow CloudTrail to run queries to populate a dashboard]in the CloudTrail User Guide.
//
// [Example: Allow CloudTrail to run queries to populate a dashboard]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/security_iam_resource-based-policy-examples.html#security_iam_resource-based-policy-examples-eds-dashboard
