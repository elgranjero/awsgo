package configservice

// DescribeComplianceByResource is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Indicates whether the specified Amazon Web Services resources are compliant. If
// a resource is noncompliant, this operation returns the number of Config rules
// that the resource does not comply with.
//
// A resource is compliant if it complies with all the Config rules that evaluate
// it. It is noncompliant if it does not comply with one or more of these rules.
//
// If Config has no current evaluation results for the resource, it returns
// INSUFFICIENT_DATA . This result might indicate one of the following conditions
// about the rules that evaluate the resource:
//
// - Config has never invoked an evaluation for the rule. To check whether it
// has, use the DescribeConfigRuleEvaluationStatus action to get the
// LastSuccessfulInvocationTime and LastFailedInvocationTime .
//
// - The rule's Lambda function is failing to send evaluation results to Config.
// Verify that the role that you assigned to your configuration recorder includes
// the config:PutEvaluations permission. If the rule is a custom rule, verify
// that the Lambda execution role includes the config:PutEvaluations permission.
//
// - The rule's Lambda function has returned NOT_APPLICABLE for all evaluation
// results. This can occur if the resources were deleted or removed from the rule's
// scope.
