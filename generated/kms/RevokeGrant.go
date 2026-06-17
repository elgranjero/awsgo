package kms

// RevokeGrant is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Deletes the specified grant. You revoke a grant to terminate the permissions
// that the grant allows. For more information, see [Retiring and revoking grants]in the Key Management Service
// Developer Guide .
//
// When you create, retire, or revoke a grant, there might be a brief delay,
// usually less than five minutes, until the grant is available throughout KMS.
// This state is known as eventual consistency. For details, see [Eventual consistency]in the Key
// Management Service Developer Guide .
//
// For detailed information about grants, including grant terminology, see [Grants in KMS] in the
// Key Management Service Developer Guide . For examples of creating grants in
// several programming languages, see [Use CreateGrant with an Amazon Web Services SDK or CLI].
//
// Cross-account use: Yes. To perform this operation on a KMS key in a different
// Amazon Web Services account, specify the key ARN in the value of the KeyId
// parameter.
//
// Required permissions: [kms:RevokeGrant] (key policy).
//
// Related operations:
//
// # CreateGrant
//
// # ListGrants
//
// # ListRetirableGrants
//
// # RetireGrant
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#terms-eventual-consistency
// [kms:RevokeGrant]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Retiring and revoking grants]: https://docs.aws.amazon.com/kms/latest/developerguide/grant-delete.html
// [Grants in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/grants.html
// [Use CreateGrant with an Amazon Web Services SDK or CLI]: https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_CreateGrant_section.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
