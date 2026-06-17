package iam

// GetDelegationRequest is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Retrieves information about a specific delegation request.
//
// If a delegation request has no owner or owner account, GetDelegationRequest for
// that delegation request can be called by any account. If the owner account is
// assigned but there is no owner id, only identities within that owner account can
// call GetDelegationRequest for the delegation request. Once the delegation
// request is fully owned, the owner of the request gets a default permission to
// get that delegation request. For more details, see [Managing Permissions for Delegation Requests].
//
// [Managing Permissions for Delegation Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies-temporary-delegation.html#temporary-delegation-managing-permissions
