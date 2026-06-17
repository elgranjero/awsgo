package marketplacemetering

// MeterUsage is generated as a reference stub.
// Executable command wiring lives under cmd/marketplacemetering.go.
//
// As a seller, your software hosted in the buyer's Amazon Web Services account
// uses this API action to emit metering records directly to Amazon Web Services
// Marketplace. You must use the following buyer Amazon Web Services account
// credentials to sign the API request.
//
// - For Amazon EC2 deployments, your software must use the [IAM role for Amazon EC2]to sign the API call
// for MeterUsage API operation.
//
// - For Amazon EKS deployments, your software must use [IAM roles for service accounts (IRSA)]to sign the API call for
// the MeterUsage API operation. Using [EKS Pod Identity], the node role, or long-term access keys
// is not supported.
//
// - For Amazon ECS deployments, your software must use [Amazon ECS task IAM]role to sign the API
// call for the MeterUsage API operation. Using the node role or long-term access
// keys are not supported.
//
// - For Amazon Bedrock AgentCore Runtime deployments, your software must use
// the [AgentCore Runtime execution role]to sign the API call for the MeterUsage API operation. Long-term access
// keys are not supported.
//
// The handling of MeterUsage requests varies between Amazon Bedrock AgentCore
// Runtime and non-Amazon Bedrock AgentCore deployments.
//
// - For non-Amazon Bedrock AgentCore Runtime deployments, you can only report
// usage once per hour for each dimension. For AMI-based products, this is per
// dimension and per EC2 instance. For container products, this is per dimension
// and per ECS task or EKS pod. You can't modify values after they're recorded. If
// you report usage before a current hour ends, you will be unable to report
// additional usage until the next hour begins. The Timestamp request parameter
// is rounded down to the hour and used to enforce this once-per-hour rule for
// idempotency. For requests that are identical after the Timestamp is rounded
// down, the API is idempotent and returns the metering record ID.
//
// - For Amazon Bedrock AgentCore Runtime deployments, you can report usage
// multiple times per hour for the same dimension. You do not need to aggregate
// metering records by the hour. You must include an idempotency token in the
// ClientToken request parameter. If using an Amazon SDK or the Amazon Web
// Services CLI, you must use the latest version which automatically includes an
// idempotency token in the ClientToken request parameter so that the request is
// processed successfully. The Timestamp request parameter is not rounded down to
// the hour and is not used for duplicate validation. Requests with duplicate
// Timestamps are aggregated as long as the ClientToken is unique.
//
// If you submit records more than six hours after events occur, the records won't
// be accepted. The timestamp in your request determines when an event is recorded.
//
// You can optionally include multiple usage allocations, to provide customers
// with usage data split into buckets by tags that you define or allow the customer
// to define.
//
// For Amazon Web Services Regions that support MeterUsage , see [MeterUsage Region support for Amazon EC2] and [MeterUsage Region support for Amazon ECS and Amazon EKS].
//
// [MeterUsage Region support for Amazon ECS and Amazon EKS]: https://docs.aws.amazon.com/marketplace/latest/APIReference/metering-regions.html#meterusage-region-support-ecs-eks
// [Amazon ECS task IAM]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html
// [AgentCore Runtime execution role]: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/runtime-permissions.html#runtime-permissions-execution
// [IAM role for Amazon EC2]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/iam-roles-for-amazon-ec2.html
// [IAM roles for service accounts (IRSA)]: https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html
// [EKS Pod Identity]: https://docs.aws.amazon.com/eks/latest/userguide/pod-identities.html
// [MeterUsage Region support for Amazon EC2]: https://docs.aws.amazon.com/marketplace/latest/APIReference/metering-regions.html#meterusage-region-support-ec2
