package iam

// GetContextKeysForPrincipalPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Gets a list of all of the context keys referenced in all the IAM policies that
// are attached to the specified IAM entity. The entity can be an IAM user, group,
// or role. If you specify a user, then the request also includes all of the
// policies attached to groups that the user is a member of.
//
// You can optionally include a list of one or more additional policies, specified
// as strings. If you want to include only a list of policies by string, use [GetContextKeysForCustomPolicy]
// instead.
//
// Note: This operation discloses information about the permissions granted to
// other users. If you do not want users to see other user's permissions, then
// consider allowing them to use [GetContextKeysForCustomPolicy]instead.
//
// Context keys are variables maintained by Amazon Web Services and its services
// that provide details about the context of an API query request. Context keys can
// be evaluated by testing against a value in an IAM policy. Use [GetContextKeysForPrincipalPolicy]to understand
// what key names and values you must supply when you call [SimulatePrincipalPolicy].
//
// [GetContextKeysForPrincipalPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetContextKeysForPrincipalPolicy.html
// [GetContextKeysForCustomPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetContextKeysForCustomPolicy.html
// [SimulatePrincipalPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_SimulatePrincipalPolicy.html
