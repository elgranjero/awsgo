package iam

// SimulateCustomPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Simulate how a set of IAM policies and optionally a resource-based policy works
// with a list of API operations and Amazon Web Services resources to determine the
// policies' effective permissions. The policies are provided as strings.
//
// The simulation does not perform the API operations; it only checks the
// authorization to determine if the simulated policies allow or deny the
// operations. You can simulate resources that don't exist in your account.
//
// If you want to simulate existing policies that are attached to an IAM user,
// group, or role, use [SimulatePrincipalPolicy]instead.
//
// Context keys are variables that are maintained by Amazon Web Services and its
// services and which provide details about the context of an API query request.
// You can use the Condition element of an IAM policy to evaluate context keys. To
// get the list of context keys that the policies require for correct simulation,
// use [GetContextKeysForCustomPolicy].
//
// If the output is long, you can use MaxItems and Marker parameters to paginate
// the results.
//
// The IAM policy simulator evaluates statements in the identity-based policy and
// the inputs that you provide during simulation. The policy simulator results can
// differ from your live Amazon Web Services environment. We recommend that you
// check your policies against your live Amazon Web Services environment after
// testing using the policy simulator to confirm that you have the desired results.
// For more information about using the policy simulator, see [Testing IAM policies with the IAM policy simulator]in the IAM User
// Guide.
//
// [GetContextKeysForCustomPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetContextKeysForCustomPolicy.html
// [Testing IAM policies with the IAM policy simulator]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_testing-policies.html
// [SimulatePrincipalPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_SimulatePrincipalPolicy.html
