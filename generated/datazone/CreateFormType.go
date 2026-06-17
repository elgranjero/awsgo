package datazone

// CreateFormType is generated as a reference stub.
// Executable command wiring lives under cmd/datazone.go.
//
// Creates a metadata form type.
//
// Prerequisites:
//
// - The domain must exist and be in an ENABLED state.
//
// - The owning project must exist and be accessible.
//
// - The name must be unique within the domain.
//
// For custom form types, to indicate that a field should be searchable, annotate
// it with (at)amazon.datazone#searchable . By default, searchable fields are indexed
// for semantic search, where related query terms will match the attribute value
// even if they are not stemmed or keyword matches. To indicate that a field should
// be indexed for lexical search (which disables semantic search but supports
// stemmed and partial matches), annotate it with
// (at)amazon.datazone#searchable(modes:["LEXICAL"]) . To indicate that a field should
// be indexed for technical identifier search (for more information on technical
// identifier search, see: [https://aws.amazon.com/blogs/big-data/streamline-data-discovery-with-precise-technical-identifier-search-in-amazon-sagemaker-unified-studio/]), annotate it with
// (at)amazon.datazone#searchable(modes:["TECHNICAL"]) .
//
// To denote that a field will store glossary term ids (which are filterable via
// the Search/SearchListings APIs), annotate it with
// (at)amazon.datazone#glossaryterm("${GLOSSARY_ID}") , where ${GLOSSARY_ID} is the
// id of the glossary that the glossary terms stored in the field belong to.
//
// [https://aws.amazon.com/blogs/big-data/streamline-data-discovery-with-precise-technical-identifier-search-in-amazon-sagemaker-unified-studio/]: https://aws.amazon.com/blogs/big-data/streamline-data-discovery-with-precise-technical-identifier-search-in-amazon-sagemaker-unified-studio/
