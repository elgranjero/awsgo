package xray

// CreateSamplingRule is generated as a reference stub.
// Executable command wiring lives under cmd/xray.go.
//
// Creates a rule to control sampling behavior for instrumented applications.
// Services retrieve rules with [GetSamplingRules], and evaluate each rule in ascending order of
// priority for each request. If a rule matches, the service records a trace,
// borrowing it from the reservoir size. After 10 seconds, the service reports back
// to X-Ray with [GetSamplingTargets]to get updated versions of each in-use rule. The updated rule
// contains a trace quota that the service can use instead of borrowing from the
// reservoir.
//
// [GetSamplingTargets]: https://docs.aws.amazon.com/xray/latest/api/API_GetSamplingTargets.html
// [GetSamplingRules]: https://docs.aws.amazon.com/xray/latest/api/API_GetSamplingRules.html
