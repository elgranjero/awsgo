package connectcases

// SearchAllRelatedItems is generated as a reference stub.
// Executable command wiring lives under cmd/connectcases.go.
//
// Searches for related items across all cases within a domain. This is a global
// search operation that returns related items from multiple cases, unlike the
// case-specific [SearchRelatedItems]API.
//
// # Use cases
//
// Following are common uses cases for this API:
//
// - Find cases with similar issues across the domain. For example, search for
// all cases containing comments about "product defect" to identify patterns and
// existing solutions.
//
// - Locate all cases associated with specific contacts or orders. For example,
// find all cases linked to a contactArn to understand the complete customer
// journey.
//
// - Monitor SLA compliance across cases. For example, search for all cases with
// "Active" SLA status to prioritize remediation efforts.
//
// Important things to know
//
// - This API returns case identifiers, not complete case objects. To retrieve
// full case details, you must make additional calls to the [GetCase]API for each
// returned case ID.
//
// - This API searches across related items content, not case fields. Use the [SearchCases]
// API to search within case field values.
//
// Endpoints: See [Amazon Connect endpoints and quotas].
//
// [GetCase]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_GetCase.html
// [SearchCases]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_SearchCases.html
// [Amazon Connect endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/connect_region.html
// [SearchRelatedItems]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_SearchRelatedItems.html
