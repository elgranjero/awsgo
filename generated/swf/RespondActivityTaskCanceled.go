package swf

// RespondActivityTaskCanceled is generated as a reference stub.
// Executable command wiring lives under cmd/swf.go.
//
// Used by workers to tell the service that the ActivityTask identified by the taskToken was
// successfully canceled. Additional details can be provided using the details
// argument.
//
// These details (if provided) appear in the ActivityTaskCanceled event added to
// the workflow history.
//
// Only use this operation if the canceled flag of a RecordActivityTaskHeartbeat request returns true and if
// the activity can be safely undone or abandoned.
//
// A task is considered open from the time that it is scheduled until it is
// closed. Therefore a task is reported as open while a worker is processing it. A
// task is closed after it has been specified in a call to RespondActivityTaskCompleted,
// RespondActivityTaskCanceled, RespondActivityTaskFailed, or the task has [timed out].
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [timed out]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dg-basic.html#swf-dev-timeout-types
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
