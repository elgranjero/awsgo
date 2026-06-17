package iam

// GetContextKeysForCustomPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Gets a list of all of the context keys referenced in the input policies. The
// policies are supplied as a list of one or more strings. To get the context keys
// from policies associated with an IAM user, group, or role, use [GetContextKeysForPrincipalPolicy].
//
// Context keys are variables maintained by Amazon Web Services and its services
// that provide details about the context of an API query request. Context keys can
// be evaluated by testing against a value specified in an IAM policy. Use
// GetContextKeysForCustomPolicy to understand what key names and values you must
// supply when you call [SimulateCustomPolicy]. Note that all parameters are shown in unencoded form
// here for clarity but must be URL encoded to be included as a part of a real HTML
// request.
//
// [GetContextKeysForPrincipalPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetContextKeysForPrincipalPolicy.html
// [SimulateCustomPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_SimulateCustomPolicy.html
