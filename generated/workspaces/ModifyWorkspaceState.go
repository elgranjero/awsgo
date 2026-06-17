package workspaces

// ModifyWorkspaceState is generated as a reference stub.
// Executable command wiring lives under cmd/workspaces.go.
//
// Sets the state of the specified WorkSpace.
//
// To maintain a WorkSpace without being interrupted, set the WorkSpace state to
// ADMIN_MAINTENANCE . WorkSpaces in this state do not respond to requests to
// reboot, stop, start, rebuild, or restore. An AutoStop WorkSpace in this state is
// not stopped. Users cannot log into a WorkSpace in the ADMIN_MAINTENANCE state.
