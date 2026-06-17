package iam

// SimulatePrincipalPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Simulate how a set of IAM policies attached to an IAM entity works with a list
// of API operations and Amazon Web Services resources to determine the policies'
// effective permissions. The entity can be an IAM user, group, or role. If you
// specify a user, then the simulation also includes all of the policies that are
// attached to groups that the user belongs to. You can simulate resources that
// don't exist in your account.
//
// You can optionally include a list of one or more additional policies specified
// as strings to include in the simulation. If you want to simulate only policies
// specified as strings, use [SimulateCustomPolicy]instead.
//
// You can also optionally include one resource-based policy to be evaluated with
// each of the resources included in the simulation for IAM users only.
//
// The simulation does not perform the API operations; it only checks the
// authorization to determine if the simulated policies allow or deny the
// operations.
//
// Note: This operation discloses information about the permissions granted to
// other users. If you do not want users to see other user's permissions, then
// consider allowing them to use [SimulateCustomPolicy]instead.
//
// Context keys are variables maintained by Amazon Web Services and its services
// that provide details about the context of an API query request. You can use the
// Condition element of an IAM policy to evaluate context keys. To get the list of
// context keys that the policies require for correct simulation, use [GetContextKeysForPrincipalPolicy].
//
// If the output is long, you can use the MaxItems and Marker parameters to
// paginate the results.
//
// The IAM policy simulator evaluates statements in the identity-based policy and
// the inputs that you provide during simulation. The policy simulator results can
// differ from your live Amazon Web Services environment. We recommend that you
// check your policies against your live Amazon Web Services environment after
// testing using the policy simulator to confirm that you have the desired results.
// For more information about using the policy simulator, see [Testing IAM policies with the IAM policy simulator]in the IAM User
// Guide.
//
// [GetContextKeysForPrincipalPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetContextKeysForPrincipalPolicy.html
// [Testing IAM policies with the IAM policy simulator]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_testing-policies.html
// [SimulateCustomPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_SimulateCustomPolicy.html
