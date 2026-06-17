package datazone

// CreateAssetFilter is generated as a reference stub.
// Executable command wiring lives under cmd/datazone.go.
//
// Creates a data asset filter.
//
// Asset filters provide a sophisticated way to create controlled views of data
// assets by selecting specific columns or applying row-level filters. This
// capability is crucial for organizations that need to share data while
// maintaining security and privacy controls. For example, your database might be
// filtered to show only non-PII fields to certain users, or sales data might be
// filtered by region for different regional teams. Asset filters enable
// fine-grained access control while maintaining a single source of truth.
//
// Prerequisites:
//
// - A valid domain ( --domain-identifier ) must exist.
//
// - A data asset ( --asset-identifier ) must already be created under that
// domain.
//
// - The asset must have the referenced columns available in its schema for
// column-based filtering.
//
// - You cannot specify both ( columnConfiguration , rowConfiguration )at the
// same time.
