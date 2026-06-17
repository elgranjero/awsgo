package datasync

// CancelTaskExecution is generated as a reference stub.
// Executable command wiring lives under cmd/datasync.go.
//
// Stops an DataSync task execution that's in progress. The transfer of some files
// are abruptly interrupted. File contents that're transferred to the destination
// might be incomplete or inconsistent with the source files.
//
// However, if you start a new task execution using the same task and allow it to
// finish, file content on the destination will be complete and consistent. This
// applies to other unexpected failures that interrupt a task execution. In all of
// these cases, DataSync successfully completes the transfer when you start the
// next task execution.
