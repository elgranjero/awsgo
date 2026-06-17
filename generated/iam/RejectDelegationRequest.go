package iam

// RejectDelegationRequest is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Rejects a delegation request, denying the requested temporary access.
//
// Once a request is rejected, it cannot be accepted or updated later. Rejected
// requests expire after 7 days.
//
// When rejecting a request, an optional explanation can be added using the Notes
// request parameter.
//
// For more details, see [Managing Permissions for Delegation Requests].
//
// [Managing Permissions for Delegation Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies-temporary-delegation.html#temporary-delegation-managing-permissions
