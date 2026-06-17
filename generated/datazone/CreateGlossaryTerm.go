package datazone

// CreateGlossaryTerm is generated as a reference stub.
// Executable command wiring lives under cmd/datazone.go.
//
// Creates a business glossary term.
//
// A glossary term represents an individual entry within the Amazon DataZone
// glossary, serving as a standardized definition for a specific business concept
// or data element. Each term can include rich metadata such as detailed
// definitions, synonyms, related terms, and usage examples. Glossary terms can be
// linked directly to data assets, providing business context to technical data
// elements. This linking capability helps users understand the business meaning of
// data fields and ensures consistent interpretation across different systems and
// teams. Terms can also have relationships with other terms, creating a semantic
// network that reflects the complexity of business concepts.
//
// Prerequisites:
//
// - Domain must exist.
//
// - Glossary must exist.
//
// - The term name must be unique within the glossary.
//
// - Ensure term does not conflict with existing terms in hierarchy.
