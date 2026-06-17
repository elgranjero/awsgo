package workspaces

// RebootWorkspaces is generated as a reference stub.
// Executable command wiring lives under cmd/workspaces.go.
//
// Reboots the specified WorkSpaces.
//
// You cannot reboot a WorkSpace unless its state is AVAILABLE , UNHEALTHY , or
// REBOOTING . Reboot a WorkSpace in the REBOOTING state only if your WorkSpace
// has been stuck in the REBOOTING state for over 20 minutes.
//
// This operation is asynchronous and returns before the WorkSpaces have rebooted.
