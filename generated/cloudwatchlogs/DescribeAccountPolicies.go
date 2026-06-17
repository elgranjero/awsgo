package cloudwatchlogs

// DescribeAccountPolicies is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Returns a list of all CloudWatch Logs account policies in the account.
//
// To use this operation, you must be signed on with the correct permissions
// depending on the type of policy that you are retrieving information for.
//
// - To see data protection policies, you must have the
// logs:GetDataProtectionPolicy and logs:DescribeAccountPolicies permissions.
//
// - To see subscription filter policies, you must have the
// logs:DescribeSubscriptionFilters and logs:DescribeAccountPolicies permissions.
//
// - To see transformer policies, you must have the logs:GetTransformer and
// logs:DescribeAccountPolicies permissions.
//
// - To see field index policies, you must have the logs:DescribeIndexPolicies
// and logs:DescribeAccountPolicies permissions.
