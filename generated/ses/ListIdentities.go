package ses

// ListIdentities is generated as a reference stub.
// Executable command wiring lives under cmd/ses.go.
//
// Returns a list containing all of the identities (email addresses and domains)
// for your Amazon Web Services account in the current Amazon Web Services Region,
// regardless of verification status.
//
// You can execute this operation no more than once per second.
//
// It's recommended that for successive pagination calls of this API, you continue
// to the use the same parameter/value pairs as used in the original call, e.g., if
// you used IdentityType=Domain in the the original call and received a NextToken
// in the response, you should continue providing the IdentityType=Domain
// parameter for further NextToken calls; however, if you didn't provide the
// IdentityType parameter in the original call, then continue to not provide it for
// successive pagination calls. Using this protocol will ensure consistent results.
