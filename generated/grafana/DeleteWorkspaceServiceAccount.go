package grafana

// DeleteWorkspaceServiceAccount is generated as a reference stub.
// Executable command wiring lives under cmd/grafana.go.
//
// Deletes a workspace service account from the workspace.
//
// This will delete any tokens created for the service account, as well. If the
// tokens are currently in use, the will fail to authenticate / authorize after
// they are deleted.
//
// Service accounts are only available for workspaces that are compatible with
// Grafana version 9 and above.
