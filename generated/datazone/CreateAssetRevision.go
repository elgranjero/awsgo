package datazone

// CreateAssetRevision is generated as a reference stub.
// Executable command wiring lives under cmd/datazone.go.
//
// Creates a revision of the asset.
//
// Asset revisions represent new versions of existing assets, capturing changes to
// either the underlying data or its metadata. They maintain a historical record of
// how assets evolve over time, who made changes, and when those changes occurred.
// This versioning capability is crucial for governance and compliance, allowing
// organizations to track changes, understand their impact, and roll back if
// necessary.
//
// Prerequisites:
//
// - Asset must already exist in the domain with identifier.
//
// - formsInput is required when asset has the form type. typeRevision should be
// the latest version of form type.
//
// - The form content must include all required fields (e.g., bucketArn for
// S3ObjectCollectionForm ).
//
// - The owning project of the original asset must still exist and be active.
//
// - User must have write access to the project and domain.
