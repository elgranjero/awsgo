package route53

// DeleteTrafficPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/route53.go.
//
// Deletes a traffic policy.
//
// When you delete a traffic policy, Route 53 sets a flag on the policy to
// indicate that it has been deleted. However, Route 53 never fully deletes the
// traffic policy. Note the following:
//
// - Deleted traffic policies aren't listed if you run [ListTrafficPolicies].
//
// - There's no way to get a list of deleted policies.
//
// - If you retain the ID of the policy, you can get information about the
// policy, including the traffic policy document, by running [GetTrafficPolicy].
//
// [ListTrafficPolicies]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_ListTrafficPolicies.html
// [GetTrafficPolicy]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetTrafficPolicy.html
