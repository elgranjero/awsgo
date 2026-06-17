package grafana

// CreateWorkspaceServiceAccountToken is generated as a reference stub.
// Executable command wiring lives under cmd/grafana.go.
//
// Creates a token that can be used to authenticate and authorize Grafana HTTP API
// operations for the given [workspace service account]. The service account acts as a user for the API
// operations, and defines the permissions that are used by the API.
//
// When you create the service account token, you will receive a key that is used
// when calling Grafana APIs. Do not lose this key, as it will not be retrievable
// again.
//
// If you do lose the key, you can delete the token and recreate it to receive a
// new key. This will disable the initial key.
//
// Service accounts are only available for workspaces that are compatible with
// Grafana version 9 and above.
//
// [workspace service account]: https://docs.aws.amazon.com/grafana/latest/userguide/service-accounts.html
