package iot

// DeletePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/iot.go.
//
// Deletes the specified policy.
//
// A policy cannot be deleted if it has non-default versions or it is attached to
// any certificate.
//
// To delete a policy, use the DeletePolicyVersion action to delete all non-default versions of the
// policy; use the DetachPolicyaction to detach the policy from any certificate; and then use
// the DeletePolicy action to delete the policy.
//
// When a policy is deleted using DeletePolicy, its default version is deleted
// with it.
//
// Because of the distributed nature of Amazon Web Services, it can take up to
// five minutes after a policy is detached before it's ready to be deleted.
//
// Requires permission to access the [DeletePolicy] action.
//
// [DeletePolicy]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
