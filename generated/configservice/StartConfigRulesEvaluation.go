package configservice

// StartConfigRulesEvaluation is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Runs an on-demand evaluation for the specified Config rules against the last
// known configuration state of the resources. Use StartConfigRulesEvaluation when
// you want to test that a rule you updated is working as expected.
// StartConfigRulesEvaluation does not re-record the latest configuration state for
// your resources. It re-runs an evaluation against the last known state of your
// resources.
//
// You can specify up to 25 Config rules per request.
//
// An existing StartConfigRulesEvaluation call for the specified rules must
// complete before you can call the API again. If you chose to have Config stream
// to an Amazon SNS topic, you will receive a ConfigRuleEvaluationStarted
// notification when the evaluation starts.
//
// You don't need to call the StartConfigRulesEvaluation API to run an evaluation
// for a new rule. When you create a rule, Config evaluates your resources against
// the rule automatically.
//
// The StartConfigRulesEvaluation API is useful if you want to run on-demand
// evaluations, such as the following example:
//
// - You have a custom rule that evaluates your IAM resources every 24 hours.
//
// - You update your Lambda function to add additional conditions to your rule.
//
// - Instead of waiting for the next periodic evaluation, you call the
// StartConfigRulesEvaluation API.
//
// - Config invokes your Lambda function and evaluates your IAM resources.
//
// - Your custom rule will still run periodic evaluations every 24 hours.
