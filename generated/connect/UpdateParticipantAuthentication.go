package connect

// UpdateParticipantAuthentication is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Instructs Amazon Connect to resume the authentication process. The subsequent
// actions depend on the request body contents:
//
// - If a code is provided: Connect retrieves the identity information from
// Amazon Cognito and imports it into Connect Customer Profiles.
//
// - If an error is provided: The error branch of the Authenticate Customer
// block is executed.
//
// The API returns a success response to acknowledge the request. However, the
// interaction and exchange of identity information occur asynchronously after the
// response is returned.
