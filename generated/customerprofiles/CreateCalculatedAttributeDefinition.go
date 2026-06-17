package customerprofiles

// CreateCalculatedAttributeDefinition is generated as a reference stub.
// Executable command wiring lives under cmd/customerprofiles.go.
//
// Creates a new calculated attribute definition. After creation, new object data
// ingested into Customer Profiles will be included in the calculated attribute,
// which can be retrieved for a profile using the [GetCalculatedAttributeForProfile]API. Defining a calculated
// attribute makes it available for all profiles within a domain. Each calculated
// attribute can only reference one ObjectType and at most, two fields from that
// ObjectType .
//
// [GetCalculatedAttributeForProfile]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_GetCalculatedAttributeForProfile.html
