package kms

// ListGrants is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Gets a list of all grants for the specified KMS key.
//
// You must specify the KMS key in all requests. You can filter the grant list by
// grant ID or grantee principal.
//
// For detailed information about grants, including grant terminology, see [Grants in KMS] in the
// Key Management Service Developer Guide . For examples of creating grants in
// several programming languages, see [Use CreateGrant with an Amazon Web Services SDK or CLI].
//
// The GranteePrincipal field in the ListGrants response usually contains the user
// or role designated as the grantee principal in the grant. However, when the
// grantee principal in the grant is an Amazon Web Services service, the
// GranteePrincipal field contains the [service principal], which might represent several different
// grantee principals.
//
// Cross-account use: Yes. To perform this operation on a KMS key in a different
// Amazon Web Services account, specify the key ARN in the value of the KeyId
// parameter.
//
// Required permissions: [kms:ListGrants] (key policy)
//
// Related operations:
//
// # CreateGrant
//
// # ListRetirableGrants
//
// # RetireGrant
//
// # RevokeGrant
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [service principal]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html#principal-services
// [Grants in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/grants.html
// [Use CreateGrant with an Amazon Web Services SDK or CLI]: https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_CreateGrant_section.html
// [kms:ListGrants]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
