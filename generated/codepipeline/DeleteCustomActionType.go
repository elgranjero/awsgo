package codepipeline

// DeleteCustomActionType is generated as a reference stub.
// Executable command wiring lives under cmd/codepipeline.go.
//
// Marks a custom action as deleted. PollForJobs for the custom action fails after
// the action is marked for deletion. Used for custom actions only.
//
// To re-create a custom action after it has been deleted you must use a string in
// the version field that has never been used before. This string can be an
// incremented version number, for example. To restore a deleted custom action, use
// a JSON file that is identical to the deleted action, including the original
// string in the version field.
