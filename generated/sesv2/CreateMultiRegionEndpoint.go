package sesv2

// CreateMultiRegionEndpoint is generated as a reference stub.
// Executable command wiring lives under cmd/sesv2.go.
//
// Creates a multi-region endpoint (global-endpoint).
//
// The primary region is going to be the AWS-Region where the operation is
// executed. The secondary region has to be provided in request's parameters. From
// the data flow standpoint there is no difference between primary and secondary
// regions - sending traffic will be split equally between the two. The primary
// region is the region where the resource has been created and where it can be
// managed.
