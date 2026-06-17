package sfn

// TestState is generated as a reference stub.
// Executable command wiring lives under cmd/sfn.go.
//
// Accepts the definition of a single state and executes it. You can test a state
// without creating a state machine or updating an existing state machine. Using
// this API, you can test the following:
//
// - A state's [input and output processing]data flow
//
// - An [Amazon Web Services service integration]request and response
//
// - An [HTTP Task]request and response
//
// You can call this API on only one state at a time. The states that you can test
// include the following:
//
// [All Task types]
// - except [Activity]
//
// [Pass]
//
// [Wait]
//
// [Choice]
//
// [Succeed]
//
// [Fail]
//
// The TestState API assumes an IAM role which must contain the required IAM
// permissions for the resources your state is accessing. For information about the
// permissions a state might need, see [IAM permissions to test a state].
//
// The TestState API can run for up to five minutes. If the execution of a state
// exceeds this duration, it fails with the States.Timeout error.
//
// TestState only supports the following when a mock is specified: [Activity tasks], .sync or
// .waitForTaskToken[service integration patterns] , [Parallel], or [Map] states.
//
// [Amazon Web Services service integration]: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-services.html
// [All Task types]: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-task-state.html#task-types
// [Choice]: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-choice-state.html
// [Activity tasks]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-activities.html
// [HTTP Task]: https://docs.aws.amazon.com/step-functions/latest/dg/call-https-apis.html
// [input and output processing]: https://docs.aws.amazon.com/step-functions/latest/dg/test-state-isolation.html#test-state-input-output-dataflow
// [Activity]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-activities.html
// [Parallel]: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-parallel-state.html
// [Succeed]: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-succeed-state.html
// [service integration patterns]: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html
// [Pass]: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-pass-state.html
// [IAM permissions to test a state]: https://docs.aws.amazon.com/step-functions/latest/dg/test-state-isolation.html#test-state-permissions
// [Wait]: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-wait-state.html
// [Map]: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-map-state.html
// [Fail]: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-fail-state.html
