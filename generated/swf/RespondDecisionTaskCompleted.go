package swf

// RespondDecisionTaskCompleted is generated as a reference stub.
// Executable command wiring lives under cmd/swf.go.
//
// Used by deciders to tell the service that the DecisionTask identified by the taskToken has
// successfully completed. The decisions argument specifies the list of decisions
// made while processing the task.
//
// A DecisionTaskCompleted event is added to the workflow history. The
// executionContext specified is attached to the event in the workflow execution
// history.
//
// # Access Control
//
// If an IAM policy grants permission to use RespondDecisionTaskCompleted , it can
// express permissions for the list of decisions in the decisions parameter. Each
// of the decisions has one or more parameters, much like a regular API call. To
// allow for policies to be as readable as possible, you can express permissions on
// decisions as if they were actual API calls, including applying conditions to
// some parameters. For more information, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
