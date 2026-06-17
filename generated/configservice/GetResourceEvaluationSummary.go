package configservice

// GetResourceEvaluationSummary is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Returns a summary of resource evaluation for the specified resource evaluation
// ID from the proactive rules that were run. The results indicate which evaluation
// context was used to evaluate the rules, which resource details were evaluated,
// the evaluation mode that was run, and whether the resource details comply with
// the configuration of the proactive rules.
//
// To see additional information about the evaluation result, such as which rule
// flagged a resource as NON_COMPLIANT, use the [GetComplianceDetailsByResource]API. For more information, see the [Examples]
// section.
//
// [GetComplianceDetailsByResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_GetComplianceDetailsByResource.html
// [Examples]: https://docs.aws.amazon.com/config/latest/APIReference/API_GetResourceEvaluationSummary.html#API_GetResourceEvaluationSummary_Examples
