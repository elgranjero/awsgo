package codecommit

// CreateUnreferencedMergeCommit is generated as a reference stub.
// Executable command wiring lives under cmd/codecommit.go.
//
// Creates an unreferenced commit that represents the result of merging two
// branches using a specified merge strategy. This can help you determine the
// outcome of a potential merge. This API cannot be used with the fast-forward
// merge strategy because that strategy does not create a merge commit.
//
// This unreferenced merge commit can only be accessed using the GetCommit API or
// through git commands such as git fetch. To retrieve this commit, you must
// specify its commit ID or otherwise reference it.
