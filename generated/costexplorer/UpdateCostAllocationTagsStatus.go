package costexplorer

// UpdateCostAllocationTagsStatus is generated as a reference stub.
// Executable command wiring lives under cmd/costexplorer.go.
//
// Updates status for cost allocation tags in bulk, with maximum batch size of 20.
// If the tag status that's updated is the same as the existing tag status, the
// request doesn't fail. Instead, it doesn't have any effect on the tag status (for
// example, activating the active tag).
