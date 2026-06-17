package kinesisvideo

// DeleteEdgeConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisvideo.go.
//
// An asynchronous API that deletes a stream’s existing edge configuration, as
// well as the corresponding media from the Edge Agent.
//
// When you invoke this API, the sync status is set to DELETING . A deletion
// process starts, in which active edge jobs are stopped and all media is deleted
// from the edge device. The time to delete varies, depending on the total amount
// of stored media. If the deletion process fails, the sync status changes to
// DELETE_FAILED . You will need to re-try the deletion.
//
// When the deletion process has completed successfully, the edge configuration is
// no longer accessible.
