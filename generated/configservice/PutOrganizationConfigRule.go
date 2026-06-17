package configservice

// PutOrganizationConfigRule is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Adds or updates an Config rule for your entire organization to evaluate if your
// Amazon Web Services resources comply with your desired configurations. For
// information on how many organization Config rules you can have per account, see [Service Limits]
// in the Config Developer Guide.
//
// Only a management account and a delegated administrator can create or update an
// organization Config rule. When calling this API with a delegated administrator,
// you must ensure Organizations ListDelegatedAdministrator permissions are added.
// An organization can have up to 3 delegated administrators.
//
// This API enables organization service access through the EnableAWSServiceAccess
// action and creates a service-linked role
// AWSServiceRoleForConfigMultiAccountSetup in the management or delegated
// administrator account of your organization. The service-linked role is created
// only when the role does not exist in the caller account. Config verifies the
// existence of role with GetRole action.
//
// To use this API with delegated administrator, register a delegated
// administrator by calling Amazon Web Services Organization
// register-delegated-administrator for config-multiaccountsetup.amazonaws.com .
//
// There are two types of rules: Config Managed Rules and Config Custom Rules. You
// can use PutOrganizationConfigRule to create both Config Managed Rules and
// Config Custom Rules.
//
// Config Managed Rules are predefined, customizable rules created by Config. For
// a list of managed rules, see [List of Config Managed Rules]. If you are adding an Config managed rule, you
// must specify the rule's identifier for the RuleIdentifier key.
//
// Config Custom Rules are rules that you create from scratch. There are two ways
// to create Config custom rules: with Lambda functions ([Lambda Developer Guide] ) and with Guard ([Guard GitHub Repository] ), a
// policy-as-code language.
//
// Config custom rules created with Lambda are called Config Custom Lambda Rules
// and Config custom rules created with Guard are called Config Custom Policy
// Rules.
//
// If you are adding a new Config Custom Lambda rule, you first need to create an
// Lambda function in the management account or a delegated administrator that the
// rule invokes to evaluate your resources. You also need to create an IAM role in
// the managed account that can be assumed by the Lambda function. When you use
// PutOrganizationConfigRule to add a Custom Lambda rule to Config, you must
// specify the Amazon Resource Name (ARN) that Lambda assigns to the function.
//
// Prerequisite: Ensure you call EnableAllFeatures API to enable all features in
// an organization.
//
// Make sure to specify one of either OrganizationCustomPolicyRuleMetadata for
// Custom Policy rules, OrganizationCustomRuleMetadata for Custom Lambda rules, or
// OrganizationManagedRuleMetadata for managed rules.
//
// [List of Config Managed Rules]: https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html
// [Lambda Developer Guide]: https://docs.aws.amazon.com/config/latest/developerguide/gettingstarted-concepts.html#gettingstarted-concepts-function
// [Service Limits]: https://docs.aws.amazon.com/config/latest/developerguide/configlimits.html
// [Guard GitHub Repository]: https://github.com/aws-cloudformation/cloudformation-guard
