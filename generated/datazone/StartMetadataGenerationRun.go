package datazone

// StartMetadataGenerationRun is generated as a reference stub.
// Executable command wiring lives under cmd/datazone.go.
//
// Starts the metadata generation run.
//
// Prerequisites:
//
// - Asset must be created and belong to the specified domain and project.
//
// - Asset type must be supported for metadata generation (e.g., Amazon Web
// Services Glue table).
//
// - Asset must have a structured schema with valid rows and columns.
//
// - Valid values for --type: BUSINESS_DESCRIPTIONS, BUSINESS_NAMES,
// BUSINESS_GLOSSARY_ASSOCIATIONS.
//
// - The user must have permission to run metadata generation in the
// domain/project.
