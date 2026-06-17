package iam

// AcceptDelegationRequest is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Accepts a delegation request, granting the requested temporary access.
//
// Once the delegation request is accepted, it is eligible to send the exchange
// token to the partner. The [SendDelegationToken]API has to be explicitly called to send the
// delegation token.
//
// At the time of acceptance, IAM records the details and the state of the
// identity that called this API. This is the identity that gets mapped to the
// delegated credential.
//
// An accepted request may be rejected before the exchange token is sent to the
// partner.
//
// [SendDelegationToken]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_SendDelegationToken.html
