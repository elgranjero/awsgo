package iam

// SendDelegationToken is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Sends the exchange token for an accepted delegation request.
//
// The exchange token is sent to the partner via an asynchronous notification
// channel, established by the partner.
//
// The delegation request must be in the ACCEPTED state when calling this API.
// After the SendDelegationToken API call is successful, the request transitions
// to a FINALIZED state and cannot be rolled back. However, a user may reject an
// accepted request before the SendDelegationToken API is called.
//
// For more details, see [Managing Permissions for Delegation Requests].
//
// [Managing Permissions for Delegation Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies-temporary-delegation.html#temporary-delegation-managing-permissions
