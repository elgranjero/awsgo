package mq

// CreateBroker is generated as a reference stub.
// Executable command wiring lives under cmd/mq.go.
//
// Creates a broker. Note: This API is asynchronous.
//
// To create a broker, you must either use the AmazonMQFullAccess IAM policy or
// include the following EC2 permissions in your IAM policy.
//
// - ec2:CreateNetworkInterface
//
// This permission is required to allow Amazon MQ to create an elastic network
//
// interface (ENI) on behalf of your account.
//
// - ec2:CreateNetworkInterfacePermission
//
// This permission is required to attach the ENI to the broker instance.
//
// - ec2:DeleteNetworkInterface
//
// - ec2:DeleteNetworkInterfacePermission
//
// - ec2:DetachNetworkInterface
//
// - ec2:DescribeInternetGateways
//
// - ec2:DescribeNetworkInterfaces
//
// - ec2:DescribeNetworkInterfacePermissions
//
// - ec2:DescribeRouteTables
//
// - ec2:DescribeSecurityGroups
//
// - ec2:DescribeSubnets
//
// - ec2:DescribeVpcs
//
// For more information, see [Create an IAM User and Get Your Amazon Web Services Credentials] and [Never Modify or Delete the Amazon MQ Elastic Network Interface] in the Amazon MQ Developer Guide.
//
// [Never Modify or Delete the Amazon MQ Elastic Network Interface]: https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/connecting-to-amazon-mq.html#never-modify-delete-elastic-network-interface
// [Create an IAM User and Get Your Amazon Web Services Credentials]: https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/amazon-mq-setting-up.html#create-iam-user
