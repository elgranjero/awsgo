package swf

// DeleteWorkflowType is generated as a reference stub.
// Executable command wiring lives under cmd/swf.go.
//
// Deletes the specified workflow type.
//
// Note: Prior to deletion, workflow types must first be deprecated.
//
// After a workflow type has been deleted, you cannot create new executions of
// that type. Executions that started before the type was deleted will continue to
// run.
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
// - Constrain the following parameters by using a Condition element with the
// appropriate keys.
//
// - workflowType.name : String constraint. The key is swf:workflowType.name .
//
// - workflowType.version : String constraint. The key is
// swf:workflowType.version .
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
