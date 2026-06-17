package sfn

// CreateStateMachine is generated as a reference stub.
// Executable command wiring lives under cmd/sfn.go.
//
// Creates a state machine. A state machine consists of a collection of states
// that can do work ( Task states), determine to which states to transition next (
// Choice states), stop an execution with an error ( Fail states), and so on.
// State machines are specified using a JSON-based, structured language. For more
// information, see [Amazon States Language]in the Step Functions User Guide.
//
// If you set the publish parameter of this API action to true , it publishes
// version 1 as the first revision of the state machine.
//
// For additional control over security, you can encrypt your data using a
// customer-managed key for Step Functions state machines. You can configure a
// symmetric KMS key and data key reuse period when creating or updating a State
// Machine. The execution history and state machine definition will be encrypted
// with the key applied to the State Machine.
//
// This operation is eventually consistent. The results are best effort and may
// not reflect very recent updates and changes.
//
// CreateStateMachine is an idempotent API. Subsequent requests won’t create a
// duplicate resource if it was already created. CreateStateMachine 's idempotency
// check is based on the state machine name , definition , type ,
// LoggingConfiguration , TracingConfiguration , and EncryptionConfiguration The
// check is also based on the publish and versionDescription parameters. If a
// following request has a different roleArn or tags , Step Functions will ignore
// these differences and treat it as an idempotent request of the previous. In this
// case, roleArn and tags will not be updated, even if they are different.
//
// [Amazon States Language]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html
