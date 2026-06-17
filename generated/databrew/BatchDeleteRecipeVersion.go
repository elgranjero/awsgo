package databrew

// BatchDeleteRecipeVersion is generated as a reference stub.
// Executable command wiring lives under cmd/databrew.go.
//
// Deletes one or more versions of a recipe at a time.
//
// The entire request will be rejected if:
//
// - The recipe does not exist.
//
// - There is an invalid version identifier in the list of versions.
//
// - The version list is empty.
//
// - The version list size exceeds 50.
//
// - The version list contains duplicate entries.
//
// The request will complete successfully, but with partial failures, if:
//
// - A version does not exist.
//
// - A version is being used by a job.
//
// - You specify LATEST_WORKING , but it's being used by a project.
//
// - The version fails to be deleted.
//
// The LATEST_WORKING version will only be deleted if the recipe has no other
// versions. If you try to delete LATEST_WORKING while other versions exist (or if
// they can't be deleted), then LATEST_WORKING will be listed as partial failure
// in the response.
