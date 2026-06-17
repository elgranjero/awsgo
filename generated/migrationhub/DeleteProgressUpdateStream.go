package migrationhub

// DeleteProgressUpdateStream is generated as a reference stub.
// Executable command wiring lives under cmd/migrationhub.go.
//
// Deletes a progress update stream, including all of its tasks, which was
// previously created as an AWS resource used for access control. This API has the
// following traits:
//
// - The only parameter needed for DeleteProgressUpdateStream is the stream name
// (same as a CreateProgressUpdateStream call).
//
// - The call will return, and a background process will asynchronously delete
// the stream and all of its resources (tasks, associated resources, resource
// attributes, created artifacts).
//
// - If the stream takes time to be deleted, it might still show up on a
// ListProgressUpdateStreams call.
//
// - CreateProgressUpdateStream , ImportMigrationTask , NotifyMigrationTaskState
// , and all Associate[*] APIs related to the tasks belonging to the stream will
// throw "InvalidInputException" if the stream of the same name is in the process
// of being deleted.
//
// - Once the stream and all of its resources are deleted,
// CreateProgressUpdateStream for a stream of the same name will succeed, and
// that stream will be an entirely new logical resource (without any resources
// associated with the old stream).
