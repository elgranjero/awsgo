package customerprofiles

// CreateDomain is generated as a reference stub.
// Executable command wiring lives under cmd/customerprofiles.go.
//
// Creates a domain, which is a container for all customer data, such as customer
// profile attributes, object types, profile keys, and encryption keys. You can
// create multiple domains, and each domain can have multiple third-party
// integrations.
//
// Each Amazon Connect instance can be associated with only one domain. Multiple
// Amazon Connect instances can be associated with one domain.
//
// Use this API or [UpdateDomain] to enable [identity resolution]: set Matching to true.
//
// To prevent cross-service impersonation when you call this API, see [Cross-service confused deputy prevention] for sample
// policies that you should apply.
//
// It is not possible to associate a Customer Profiles domain with an Amazon
// Connect Instance directly from the API. If you would like to create a domain and
// associate a Customer Profiles domain, use the Amazon Connect admin website. For
// more information, see [Enable Customer Profiles].
//
// Each Amazon Connect instance can be associated with only one domain. Multiple
// Amazon Connect instances can be associated with one domain.
//
// [UpdateDomain]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_UpdateDomain.html
// [Enable Customer Profiles]: https://docs.aws.amazon.com/connect/latest/adminguide/enable-customer-profiles.html#enable-customer-profiles-step1
// [Cross-service confused deputy prevention]: https://docs.aws.amazon.com/connect/latest/adminguide/cross-service-confused-deputy-prevention.html
// [identity resolution]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_GetMatches.html
