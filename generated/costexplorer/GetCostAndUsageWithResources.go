package costexplorer

// GetCostAndUsageWithResources is generated as a reference stub.
// Executable command wiring lives under cmd/costexplorer.go.
//
// Retrieves cost and usage metrics with resources for your account. You can
// specify which cost and usage-related metric, such as BlendedCosts or
// UsageQuantity , that you want the request to return. You can also filter and
// group your data by various dimensions, such as SERVICE or AZ , in a specific
// time range. For a complete list of valid dimensions, see the [GetDimensionValues]operation.
// Management account in an organization in Organizations have access to all member
// accounts.
//
// Hourly granularity is only available for EC2-Instances (Elastic Compute Cloud)
// resource-level data. All other resource-level data is available at daily
// granularity.
//
// This is an opt-in only feature. You can enable this feature from the Cost
// Explorer Settings page. For information about how to access the Settings page,
// see [Controlling Access for Cost Explorer]in the Billing and Cost Management User Guide.
//
// [GetDimensionValues]: https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_GetDimensionValues.html
// [Controlling Access for Cost Explorer]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/ce-access.html
