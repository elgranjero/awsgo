package scheduler

// DeleteScheduleGroup is generated as a reference stub.
// Executable command wiring lives under cmd/scheduler.go.
//
// Deletes the specified schedule group. Deleting a schedule group results in
// EventBridge Scheduler deleting all schedules associated with the group. When you
// delete a group, it remains in a DELETING state until all of its associated
// schedules are deleted. Schedules associated with the group that are set to run
// while the schedule group is in the process of being deleted might continue to
// invoke their targets until the schedule group and its associated schedules are
// deleted.
//
// This operation is eventually consistent.
