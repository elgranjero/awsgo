package sfn

// ValidateStateMachineDefinition is generated as a reference stub.
// Executable command wiring lives under cmd/sfn.go.
//
// Validates the syntax of a state machine definition specified in [Amazon States Language] (ASL), a
// JSON-based, structured language.
//
// You can validate that a state machine definition is correct without creating a
// state machine resource.
//
// Suggested uses for ValidateStateMachineDefinition :
//
// - Integrate automated checks into your code review or Continuous Integration
// (CI) process to check state machine definitions before starting deployments.
//
// - Run validation from a Git pre-commit hook to verify the definition before
// committing to your source repository.
//
// Validation will look for problems in your state machine definition and return a
// result and a list of diagnostic elements.
//
// The result value will be OK when your workflow definition can be successfully
// created or updated. Note the result can be OK even when diagnostic warnings are
// present in the response. The result value will be FAIL when the workflow
// definition contains errors that would prevent you from creating or updating your
// state machine.
//
// The list of [ValidateStateMachineDefinitionDiagnostic] data elements can contain zero or more WARNING and/or ERROR
// elements.
//
// The ValidateStateMachineDefinition API might add new diagnostics in the future,
// adjust diagnostic codes, or change the message wording. Your automated processes
// should only rely on the value of the result field value (OK, FAIL). Do not rely
// on the exact order, count, or wording of diagnostic messages.
//
// [Amazon States Language]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html
// [ValidateStateMachineDefinitionDiagnostic]: https://docs.aws.amazon.com/step-functions/latest/apireference/API_ValidateStateMachineDefinitionDiagnostic.html
