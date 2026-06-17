package workmail

// TestAvailabilityConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/workmail.go.
//
// Performs a test on an availability provider to ensure that access is allowed.
// For EWS, it verifies the provided credentials can be used to successfully log
// in. For Lambda, it verifies that the Lambda function can be invoked and that the
// resource access policy was configured to deny anonymous access. An anonymous
// invocation is one done without providing either a SourceArn or SourceAccount
// header.
//
// The request must contain either one provider definition ( EwsProvider or
// LambdaProvider ) or the DomainName parameter. If the DomainName parameter is
// provided, the configuration stored under the DomainName will be tested.
