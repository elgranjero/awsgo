package configservice

// PutConformancePack is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Creates or updates a conformance pack. A conformance pack is a collection of
// Config rules that can be easily deployed in an account and a region and across
// an organization. For information on how many conformance packs you can have per
// account, see [Service Limits]in the Config Developer Guide.
//
// When you use PutConformancePack to deploy conformance packs in your account,
// the operation can create Config rules and remediation actions without requiring
// config:PutConfigRule or config:PutRemediationConfigurations permissions in your
// account IAM policies.
//
// This API uses the AWSServiceRoleForConfigConforms service-linked role in your
// account to create conformance pack resources. This service-linked role includes
// the permissions to create Config rules and remediation configurations, even if
// your account IAM policies explicitly deny these actions.
//
// This API creates a service-linked role AWSServiceRoleForConfigConforms in your
// account. The service-linked role is created only when the role does not exist in
// your account.
//
// You must specify only one of the follow parameters: TemplateS3Uri , TemplateBody
// or TemplateSSMDocumentDetails .
//
// # Tags are added at creation and cannot be updated with this operation
//
// PutConformancePack is an idempotent API. Subsequent requests won't create a
// duplicate resource if one was already created. If a following request has
// different tags values, Config will ignore these differences and treat it as an
// idempotent request of the previous. In this case, tags will not be updated,
// even if they are different.
//
// Use [TagResource] and [UntagResource] to update tags after creation.
//
// [TagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_UntagResource.html
// [Service Limits]: https://docs.aws.amazon.com/config/latest/developerguide/configlimits.html
