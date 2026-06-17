package configservice

// DescribeComplianceByConfigRule is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Indicates whether the specified Config rules are compliant. If a rule is
// noncompliant, this operation returns the number of Amazon Web Services resources
// that do not comply with the rule.
//
// A rule is compliant if all of the evaluated resources comply with it. It is
// noncompliant if any of these resources do not comply.
//
// If Config has no current evaluation results for the rule, it returns
// INSUFFICIENT_DATA . This result might indicate one of the following conditions:
//
// - Config has never invoked an evaluation for the rule. To check whether it
// has, use the DescribeConfigRuleEvaluationStatus action to get the
// LastSuccessfulInvocationTime and LastFailedInvocationTime .
//
// - The rule's Lambda function is failing to send evaluation results to Config.
// Verify that the role you assigned to your configuration recorder includes the
// config:PutEvaluations permission. If the rule is a custom rule, verify that
// the Lambda execution role includes the config:PutEvaluations permission.
//
// - The rule's Lambda function has returned NOT_APPLICABLE for all evaluation
// results. This can occur if the resources were deleted or removed from the rule's
// scope.
