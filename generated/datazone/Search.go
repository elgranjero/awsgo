package datazone

// Search is generated as a reference stub.
// Executable command wiring lives under cmd/datazone.go.
//
// Searches for assets in Amazon DataZone.
//
// Search in Amazon DataZone is a powerful capability that enables users to
// discover and explore data assets, glossary terms, and data products across their
// organization. It provides both basic and advanced search functionality, allowing
// users to find resources based on names, descriptions, metadata, and other
// attributes. Search can be scoped to specific types of resources (like assets,
// glossary terms, or data products) and can be filtered using various criteria
// such as creation date, owner, or status. The search functionality is essential
// for making the wealth of data resources in an organization discoverable and
// usable, helping users find the right data for their needs quickly and
// efficiently.
//
// Many search commands in Amazon DataZone are paginated, including search and
// search-types . When the result set is large, Amazon DataZone returns a nextToken
// in the response. This token can be used to retrieve the next page of results.
//
// Prerequisites:
//
// - The --domain-identifier must refer to an existing Amazon DataZone domain.
//
// - --search-scope must be one of: ASSET, GLOSSARY_TERM, DATA_PRODUCT, or
// GLOSSARY.
//
// - The user must have search permissions in the specified domain.
//
// - If using --filters, ensure that the JSON is well-formed and that each
// filter includes valid attribute and value keys.
//
// - For paginated results, be prepared to use --next-token to fetch additional
// pages.
//
// To run a standard free-text search, the searchText parameter must be supplied.
// By default, all searchable fields are indexed for semantic search and will
// return semantic matches for SearchListings queries. To prevent semantic search
// indexing for a custom form attribute, see the [CreateFormType API documentation]. To run a lexical search query,
// enclose the query with double quotes (""). This will disable semantic search
// even for fields that have semantic search enabled and will only return results
// that contain the keywords wrapped by double quotes (order of tokens in the query
// is not enforced). Free-text search is supported for all attributes annotated
// with (at)amazon.datazone#searchable.
//
// To run a filtered search, provide filter clause using the filters parameter. To
// filter on glossary terms, use the special attribute __DataZoneGlossaryTerms . To
// filter on an indexed numeric attribute (i.e., a numeric attribute annotated with
// (at)amazon.datazone#sortable ), provide a filter using the intValue parameter. The
// filters parameter can also be used to run more advanced free-text searches that
// target specific attributes (attributes must be annotated with
// (at)amazon.datazone#searchable for free-text search). Create/update timestamp
// filtering is supported using the special creationTime / lastUpdatedTime
// attributes. Filter types can be mixed and matched to power complex queries.
//
// To find out whether an attribute has been annotated and indexed for a given
// search type, use the GetFormType API to retrieve the form containing the
// attribute.
//
// [CreateFormType API documentation]: https://docs.aws.amazon.com/datazone/latest/APIReference/API_CreateFormType.html
