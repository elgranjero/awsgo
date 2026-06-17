package synthetics

// DescribeCanaries is generated as a reference stub.
// Executable command wiring lives under cmd/synthetics.go.
//
// This operation returns a list of the canaries in your account, along with full
// details about each canary.
//
// This operation supports resource-level authorization using an IAM policy and
// the Names parameter. If you specify the Names parameter, the operation is
// successful only if you have authorization to view all the canaries that you
// specify in your request. If you do not have permission to view any of the
// canaries, the request fails with a 403 response.
//
// You are required to use the Names parameter if you are logged on to a user or
// role that has an IAM policy that restricts which canaries that you are allowed
// to view. For more information, see [Limiting a user to viewing specific canaries].
//
// [Limiting a user to viewing specific canaries]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Restricted.html
