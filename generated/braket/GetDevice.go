package braket

// GetDevice is generated as a reference stub.
// Executable command wiring lives under cmd/braket.go.
//
// Retrieves the devices available in Amazon Braket.
//
// For backwards compatibility with older versions of BraketSchemas, OpenQASM
// information is omitted from GetDevice API calls. To get this information the
// user-agent needs to present a recent version of the BraketSchemas (1.8.0 or
// later). The Braket SDK automatically reports this for you. If you do not see
// OpenQASM results in the GetDevice response when using a Braket SDK, you may need
// to set AWS_EXECUTION_ENV environment variable to configure user-agent. See the
// code examples provided below for how to do this for the AWS CLI, Boto3, and the
// Go, Java, and JavaScript/TypeScript SDKs.
