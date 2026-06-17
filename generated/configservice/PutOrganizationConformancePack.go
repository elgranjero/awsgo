package configservice

// PutOrganizationConformancePack is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Deploys conformance packs across member accounts in an Amazon Web Services
// Organization. For information on how many organization conformance packs and how
// many Config rules you can have per account, see [Service Limits]in the Config Developer Guide.
//
// Only a management account and a delegated administrator can call this API. When
// calling this API with a delegated administrator, you must ensure Organizations
// ListDelegatedAdministrator permissions are added. An organization can have up to
// 3 delegated administrators.
//
// When you use PutOrganizationConformancePack to deploy conformance packs across
// member accounts, the operation can create Config rules and remediation actions
// without requiring config:PutConfigRule or config:PutRemediationConfigurations
// permissions in member account IAM policies.
//
// This API uses the AWSServiceRoleForConfigConforms service-linked role in each
// member account to create conformance pack resources. This service-linked role
// includes the permissions to create Config rules and remediation configurations,
// even if member account IAM policies explicitly deny these actions.
//
// This API enables organization service access for
// config-multiaccountsetup.amazonaws.com through the EnableAWSServiceAccess
// action and creates a service-linked role
// AWSServiceRoleForConfigMultiAccountSetup in the management or delegated
// administrator account of your organization. The service-linked role is created
// only when the role does not exist in the caller account. To use this API with
// delegated administrator, register a delegated administrator by calling Amazon
// Web Services Organization register-delegate-admin for
// config-multiaccountsetup.amazonaws.com .
//
// Prerequisite: Ensure you call EnableAllFeatures API to enable all features in
// an organization.
//
// You must specify either the TemplateS3Uri or the TemplateBody parameter, but
// not both. If you provide both Config uses the TemplateS3Uri parameter and
// ignores the TemplateBody parameter.
//
// Config sets the state of a conformance pack to CREATE_IN_PROGRESS and
// UPDATE_IN_PROGRESS until the conformance pack is created or updated. You cannot
// update a conformance pack while it is in this state.
//
// [Service Limits]: https://docs.aws.amazon.com/config/latest/developerguide/configlimits.html
