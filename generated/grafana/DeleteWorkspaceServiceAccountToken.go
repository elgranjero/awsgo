package grafana

// DeleteWorkspaceServiceAccountToken is generated as a reference stub.
// Executable command wiring lives under cmd/grafana.go.
//
// Deletes a token for the workspace service account.
//
// This will disable the key associated with the token. If any automation is
// currently using the key, it will no longer be authenticated or authorized to
// perform actions with the Grafana HTTP APIs.
//
// Service accounts are only available for workspaces that are compatible with
// Grafana version 9 and above.
