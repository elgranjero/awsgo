package iam

// AssociateDelegationRequest is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Associates a delegation request with the current identity.
//
// If the partner that created the delegation request has specified the owner
// account during creation, only an identity from that owner account can call the
// AssociateDelegationRequest API for the specified delegation request. Once the
// AssociateDelegationRequest API call is successful, the ARN of the current
// calling identity will be stored as the ownerId of the request.
//
// If the partner that created the delegation request has not specified the owner
// account during creation, any caller from any account can call the
// AssociateDelegationRequest API for the delegation request. Once this API call is
// successful, the ARN of the current calling identity will be stored as the
// ownerId and the Amazon Web Services account ID of the current calling identity
// will be stored as the ownerAccount of the request.
//
// For more details, see [Managing Permissions for Delegation Requests].
//
// [Managing Permissions for Delegation Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies-temporary-delegation.html#temporary-delegation-managing-permissions
