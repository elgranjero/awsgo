package cognitoidentityprovider

// ListUserPoolClientSecrets is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Lists all client secrets associated with a user pool app client. Returns
// metadata about the secrets. The response does not include pagination tokens as
// there are only 2 secrets at any given time and we return both with every
// ListUserPoolClientSecrets call. For security reasons, the response never reveals
// the actual secret value in ClientSecretValue.
