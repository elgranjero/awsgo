package grafana

// ListWorkspaceServiceAccountTokens is generated as a reference stub.
// Executable command wiring lives under cmd/grafana.go.
//
// Returns a list of tokens for a workspace service account.
//
// This does not return the key for each token. You cannot access keys after they
// are created. To create a new key, delete the token and recreate it.
//
// Service accounts are only available for workspaces that are compatible with
// Grafana version 9 and above.
