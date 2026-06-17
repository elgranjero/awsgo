package datazone

// SearchTypes is generated as a reference stub.
// Executable command wiring lives under cmd/datazone.go.
//
// Searches for types in Amazon DataZone.
//
// Prerequisites:
//
// - The --domain-identifier must refer to an existing Amazon DataZone domain.
//
// - --search-scope must be one of the valid values including: ASSET_TYPE,
// GLOSSARY_TERM_TYPE, DATA_PRODUCT_TYPE.
//
// - The --managed flag must be present without a value.
//
// - The user must have permissions for form or asset types in the domain.
//
// - If using --filters, ensure that the JSON is valid.
//
// - Filters contain correct structure (attribute, value, operator).
