package bedrock

// ListGuardrails is generated as a reference stub.
// Executable command wiring lives under cmd/bedrock.go.
//
// Lists details about all the guardrails in an account. To list the DRAFT version
// of all your guardrails, don't specify the guardrailIdentifier field. To list
// all versions of a guardrail, specify the ARN of the guardrail in the
// guardrailIdentifier field.
//
// You can set the maximum number of results to return in a response in the
// maxResults field. If there are more results than the number you set, the
// response returns a nextToken that you can send in another ListGuardrails
// request to see the next batch of results.
