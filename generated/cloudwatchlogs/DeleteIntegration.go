package cloudwatchlogs

// DeleteIntegration is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Deletes the integration between CloudWatch Logs and OpenSearch Service. If your
// integration has active vended logs dashboards, you must specify true for the
// force parameter, otherwise the operation will fail. If you delete the
// integration by setting force to true , all your vended logs dashboards powered
// by OpenSearch Service will be deleted and the data that was on them will no
// longer be accessible.
