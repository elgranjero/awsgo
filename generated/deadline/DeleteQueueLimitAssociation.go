package deadline

// DeleteQueueLimitAssociation is generated as a reference stub.
// Executable command wiring lives under cmd/deadline.go.
//
// Removes the association between a queue and a limit. You must use the
// UpdateQueueLimitAssociation operation to set the status to
// STOP_LIMIT_USAGE_AND_COMPLETE_TASKS or STOP_LIMIT_USAGE_AND_CANCEL_TASKS . The
// status does not change immediately. Use the GetQueueLimitAssociation operation
// to see if the status changed to STOPPED before deleting the association.
