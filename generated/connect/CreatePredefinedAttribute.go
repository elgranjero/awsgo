package connect

// CreatePredefinedAttribute is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Creates a new predefined attribute for the specified Amazon Connect instance. A
// predefined attribute is made up of a name and a value.
//
// For the predefined attributes per instance quota, see [Amazon Connect quotas].
//
// # Use cases
//
// Following are common uses cases for this API:
//
// - Create an attribute for routing proficiency (for example, agent
// certification) that has predefined values (for example, a list of possible
// certifications). For more information, see [Create predefined attributes for routing contacts to agents].
//
// - Create an attribute for business unit name that has a list of predefined
// business unit names used in your organization. This is a use case where
// information for a contact varies between transfers or conferences. For more
// information, see [Use contact segment attributes].
//
// Endpoints: See [Amazon Connect endpoints and quotas].
//
// [Use contact segment attributes]: https://docs.aws.amazon.com/connect/latest/adminguide/use-contact-segment-attributes.html
// [Amazon Connect endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/connect_region.html
// [Amazon Connect quotas]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html#connect-quotas
// [Create predefined attributes for routing contacts to agents]: https://docs.aws.amazon.com/connect/latest/adminguide/predefined-attributes.html
