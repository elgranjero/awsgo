package customerprofiles

// UpdateDomain is generated as a reference stub.
// Executable command wiring lives under cmd/customerprofiles.go.
//
// Updates the properties of a domain, including creating or selecting a dead
// letter queue or an encryption key.
//
// After a domain is created, the name can’t be changed.
//
// Use this API or [CreateDomain] to enable [identity resolution]: set Matching to true.
//
// To prevent cross-service impersonation when you call this API, see [Cross-service confused deputy prevention] for sample
// policies that you should apply.
//
// To add or remove tags on an existing Domain, see [TagResource]/[UntagResource] .
//
// [CreateDomain]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_CreateDomain.html
// [TagResource]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_TagResource.html
// [Cross-service confused deputy prevention]: https://docs.aws.amazon.com/connect/latest/adminguide/cross-service-confused-deputy-prevention.html
// [UntagResource]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_UntagResource.html
// [identity resolution]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_GetMatches.html
