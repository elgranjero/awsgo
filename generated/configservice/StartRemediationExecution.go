package configservice

// StartRemediationExecution is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Runs an on-demand remediation for the specified Config rules against the last
// known remediation configuration. It runs an execution against the current state
// of your resources. Remediation execution is asynchronous.
//
// You can specify up to 100 resource keys per request. An existing
// StartRemediationExecution call for the specified resource keys must complete
// before you can call the API again.
