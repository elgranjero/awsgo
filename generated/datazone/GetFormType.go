package datazone

// GetFormType is generated as a reference stub.
// Executable command wiring lives under cmd/datazone.go.
//
// Gets a metadata form type in Amazon DataZone.
//
// Form types define the structure and validation rules for collecting metadata
// about assets in Amazon DataZone. They act as templates that ensure consistent
// metadata capture across similar types of assets, while allowing for
// customization to meet specific organizational needs. Form types can include
// required fields, validation rules, and dependencies, helping maintain
// high-quality metadata that makes data assets more discoverable and usable.
//
// - The form type with the specified identifier must exist in the given domain.
//
// - The domain must be valid and active.
//
// - User must have permission on the form type.
//
// - The form type should not be deleted or in an invalid state.
//
// One use case for this API is to determine whether a form field is indexed for
// search.
//
// A searchable field will be annotated with (at)amazon.datazone#searchable . By
// default, searchable fields are indexed for semantic search, where related query
// terms will match the attribute value even if they are not stemmed or keyword
// matches. If a field is indexed technical identifier search, it will be annotated
// with (at)amazon.datazone#searchable(modes:["TECHNICAL"]) . If a field is indexed
// for lexical search (supports stemmed and prefix matches but not semantic
// matches), it will be annotated with
// (at)amazon.datazone#searchable(modes:["LEXICAL"]) .
//
// A field storing glossary term IDs (which is filterable) will be annotated with
// (at)amazon.datazone#glossaryterm("${glossaryId}") .
