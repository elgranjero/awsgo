package swf

// PollForActivityTask is generated as a reference stub.
// Executable command wiring lives under cmd/swf.go.
//
// Used by workers to get an ActivityTask from the specified activity taskList . This initiates
// a long poll, where the service holds the HTTP connection open and responds as
// soon as a task becomes available. The maximum time the service holds on to the
// request before responding is 60 seconds. If no task is available within 60
// seconds, the poll returns an empty result. An empty result, in this context,
// means that an ActivityTask is returned, but that the value of taskToken is an
// empty string. If a task is returned, the worker should use its type to identify
// and process it correctly.
//
// Workers should set their client side socket timeout to at least 70 seconds (10
// seconds higher than the maximum time service may hold the poll request).
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
