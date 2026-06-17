package grafana

// CreateWorkspaceServiceAccount is generated as a reference stub.
// Executable command wiring lives under cmd/grafana.go.
//
// Creates a service account for the workspace. A service account can be used to
// call Grafana HTTP APIs, and run automated workloads. After creating the service
// account with the correct GrafanaRole for your use case, use
// CreateWorkspaceServiceAccountToken to create a token that can be used to
// authenticate and authorize Grafana HTTP API calls.
//
// You can only create service accounts for workspaces that are compatible with
// Grafana version 9 and above.
//
// For more information about service accounts, see [Service accounts] in the Amazon Managed Grafana
// User Guide.
//
// For more information about the Grafana HTTP APIs, see [Using Grafana HTTP APIs] in the Amazon Managed
// Grafana User Guide.
//
// [Service accounts]: https://docs.aws.amazon.com/grafana/latest/userguide/service-accounts.html
// [Using Grafana HTTP APIs]: https://docs.aws.amazon.com/grafana/latest/userguide/Using-Grafana-APIs.html
