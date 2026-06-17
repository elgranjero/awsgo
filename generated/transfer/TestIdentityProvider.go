package transfer

// TestIdentityProvider is generated as a reference stub.
// Executable command wiring lives under cmd/transfer.go.
//
// If the IdentityProviderType of a file transfer protocol-enabled server is
// AWS_DIRECTORY_SERVICE or API_Gateway , tests whether your identity provider is
// set up successfully. We highly recommend that you call this operation to test
// your authentication method as soon as you create your server. By doing so, you
// can troubleshoot issues with the identity provider integration to ensure that
// your users can successfully use the service.
//
// The ServerId and UserName parameters are required. The ServerProtocol , SourceIp
// , and UserPassword are all optional.
//
// Note the following:
//
// - You cannot use TestIdentityProvider if the IdentityProviderType of your
// server is SERVICE_MANAGED .
//
// - TestIdentityProvider does not work with keys: it only accepts passwords.
//
// - TestIdentityProvider can test the password operation for a custom Identity
// Provider that handles keys and passwords.
//
// - If you provide any incorrect values for any parameters, the Response field
// is empty.
//
// - If you provide a server ID for a server that uses service-managed users,
// you get an error:
//
// An error occurred (InvalidRequestException) when calling the
//
// TestIdentityProvider operation: s-server-ID not configured for external auth
//
// - If you enter a Server ID for the --server-id parameter that does not
// identify an actual Transfer server, you receive the following error:
//
// An error occurred (ResourceNotFoundException) when calling the
//
// TestIdentityProvider operation: Unknown server .
//
// It is possible your sever is in a different region. You can specify a region by
//
// adding the following: --region region-code , such as --region us-east-2 to
// specify a server in US East (Ohio).
