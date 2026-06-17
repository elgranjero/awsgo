package sfn

// DescribeStateMachine is generated as a reference stub.
// Executable command wiring lives under cmd/sfn.go.
//
// Provides information about a state machine's definition, its IAM role Amazon
// Resource Name (ARN), and configuration.
//
// A qualified state machine ARN can either refer to a Distributed Map state
// defined within a state machine, a version ARN, or an alias ARN.
//
// The following are some examples of qualified and unqualified state machine ARNs:
//
// - The following qualified state machine ARN refers to a Distributed Map state
// with a label mapStateLabel in a state machine named myStateMachine .
//
// arn:partition:states:region:account-id:stateMachine:myStateMachine/mapStateLabel
//
// If you provide a qualified state machine ARN that refers to a Distributed Map
//
// state, the request fails with ValidationException .
//
// - The following qualified state machine ARN refers to an alias named PROD .
//
// arn::states:::stateMachine:
//
// If you provide a qualified state machine ARN that refers to a version ARN or an
//
// alias ARN, the request starts execution for that version or alias.
//
// - The following unqualified state machine ARN refers to a state machine named
// myStateMachine .
//
// arn::states:::stateMachine:
//
// This API action returns the details for a state machine version if the
// stateMachineArn you specify is a state machine version ARN.
//
// This operation is eventually consistent. The results are best effort and may
// not reflect very recent updates and changes.
