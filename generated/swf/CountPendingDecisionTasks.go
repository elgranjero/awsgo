package swf

// CountPendingDecisionTasks is generated as a reference stub.
// Executable command wiring lives under cmd/swf.go.
//
// Returns the estimated number of decision tasks in the specified task list. The
// count returned is an approximation and isn't guaranteed to be exact. If you
// specify a task list that no decision task was ever scheduled in then 0 is
// returned.
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
// - Constrain the taskList.name parameter by using a Condition element with the
// swf:taskList.name key to allow the action to access only certain task lists.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
