package swf

// StartWorkflowExecution is generated as a reference stub.
// Executable command wiring lives under cmd/swf.go.
//
// Starts an execution of the workflow type in the specified domain using the
// provided workflowId and input data.
//
// This action returns the newly started workflow execution.
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
// - tagList.member.0 : The key is swf:tagList.member.0 .
//
// - tagList.member.1 : The key is swf:tagList.member.1 .
//
// - tagList.member.2 : The key is swf:tagList.member.2 .
//
// - tagList.member.3 : The key is swf:tagList.member.3 .
//
// - tagList.member.4 : The key is swf:tagList.member.4 .
//
// - taskList : String constraint. The key is swf:taskList.name .
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
