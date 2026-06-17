package quicksight

// CreateAccountSubscription is generated as a reference stub.
// Executable command wiring lives under cmd/quicksight.go.
//
// Creates an Amazon Quick Sight account, or subscribes to Amazon Quick Sight Q.
//
// The Amazon Web Services Region for the account is derived from what is
// configured in the CLI or SDK.
//
// Before you use this operation, make sure that you can connect to an existing
// Amazon Web Services account. If you don't have an Amazon Web Services account,
// see [Sign up for Amazon Web Services]in the Amazon Quick Sight User Guide. The person who signs up for Amazon
// Quick Sight needs to have the correct Identity and Access Management (IAM)
// permissions. For more information, see [IAM Policy Examples for Amazon Quick Sight]in the Amazon Quick Sight User Guide.
//
// If your IAM policy includes both the Subscribe and CreateAccountSubscription
// actions, make sure that both actions are set to Allow . If either action is set
// to Deny , the Deny action prevails and your API call fails.
//
// You can't pass an existing IAM role to access other Amazon Web Services
// services using this API operation. To pass your existing IAM role to Amazon
// Quick Sight, see [Passing IAM roles to Amazon Quick Sight]in the Amazon Quick Sight User Guide.
//
// You can't set default resource access on the new account from the Amazon Quick
// Sight API. Instead, add default resource access from the Amazon Quick Sight
// console. For more information about setting default resource access to Amazon
// Web Services services, see [Setting default resource access to Amazon Web Services services]in the Amazon Quick Sight User Guide.
//
// [Setting default resource access to Amazon Web Services services]: https://docs.aws.amazon.com/quicksight/latest/user/scoping-policies-defaults.html
// [IAM Policy Examples for Amazon Quick Sight]: https://docs.aws.amazon.com/quicksight/latest/user/iam-policy-examples.html
// [Passing IAM roles to Amazon Quick Sight]: https://docs.aws.amazon.com/quicksight/latest/user/security_iam_service-with-iam.html#security-create-iam-role
// [Sign up for Amazon Web Services]: https://docs.aws.amazon.com/quicksight/latest/user/setting-up-aws-sign-up.html
