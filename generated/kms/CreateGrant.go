package kms

// CreateGrant is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Adds a grant to a KMS key.
//
// A grant is a policy instrument that allows Amazon Web Services principals to
// use KMS keys in cryptographic operations. It also can allow them to view a KMS
// key (DescribeKey ) and create and manage grants. When authorizing access to a KMS key,
// grants are considered along with key policies and IAM policies. Grants are often
// used for temporary permissions because you can create one, use its permissions,
// and delete it without changing your key policies or IAM policies.
//
// For detailed information about grants, including grant terminology, see [Grants in KMS] in the
// Key Management Service Developer Guide . For examples of creating grants in
// several programming languages, see [Use CreateGrant with an Amazon Web Services SDK or CLI].
//
// The CreateGrant operation returns a GrantToken and a GrantId .
//
// - When you create, retire, or revoke a grant, there might be a brief delay,
// usually less than five minutes, until the grant is available throughout KMS.
// This state is known as eventual consistency. Once the grant has achieved
// eventual consistency, the grantee principal can use the permissions in the grant
// without identifying the grant.
//
// However, to use the permissions in the grant immediately, use the GrantToken
//
// that CreateGrant returns. For details, see [Using a grant token]in the Key Management Service
// Developer Guide .
//
// - The CreateGrant operation also returns a GrantId . You can use the GrantId
// and a key identifier to identify the grant in the RetireGrantand RevokeGrantoperations. To find the
// grant ID, use the ListGrantsor ListRetirableGrantsoperations.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. To perform this operation on a KMS key in a different
// Amazon Web Services account, specify the key ARN in the value of the KeyId
// parameter.
//
// Required permissions: [kms:CreateGrant] (key policy)
//
// Related operations:
//
// # ListGrants
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
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [Grants in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/grants.html
// [kms:CreateGrant]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Use CreateGrant with an Amazon Web Services SDK or CLI]: https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_CreateGrant_section.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
//
// [Using a grant token]: https://docs.aws.amazon.com/kms/latest/developerguide/using-grant-token.html
