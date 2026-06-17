package iot

// CreatePolicyVersion is generated as a reference stub.
// Executable command wiring lives under cmd/iot.go.
//
// Creates a new version of the specified IoT policy. To update a policy, create a
// new policy version. A managed policy can have up to five versions. If the policy
// has five versions, you must use DeletePolicyVersionto delete an existing version before you create
// a new one.
//
// Optionally, you can set the new version as the policy's default version. The
// default version is the operative version (that is, the version that is in effect
// for the certificates to which the policy is attached).
//
// Requires permission to access the [CreatePolicyVersion] action.
//
// [CreatePolicyVersion]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
