package synthetics

// CreateCanary is generated as a reference stub.
// Executable command wiring lives under cmd/synthetics.go.
//
// Creates a canary. Canaries are scripts that monitor your endpoints and APIs
// from the outside-in. Canaries help you check the availability and latency of
// your web services and troubleshoot anomalies by investigating load time data,
// screenshots of the UI, logs, and metrics. You can set up a canary to run
// continuously or just once.
//
// Do not use CreateCanary to modify an existing canary. Use [UpdateCanary] instead.
//
// To create canaries, you must have the CloudWatchSyntheticsFullAccess policy. If
// you are creating a new IAM role for the canary, you also need the iam:CreateRole
// , iam:CreatePolicy and iam:AttachRolePolicy permissions. For more information,
// see [Necessary Roles and Permissions].
//
// Do not include secrets or proprietary information in your canary names. The
// canary name makes up part of the Amazon Resource Name (ARN) for the canary, and
// the ARN is included in outbound calls over the internet. For more information,
// see [Security Considerations for Synthetics Canaries].
//
// [UpdateCanary]: https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_UpdateCanary.html
// [Necessary Roles and Permissions]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Roles
// [Security Considerations for Synthetics Canaries]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/servicelens_canaries_security.html
