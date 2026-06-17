package kms

// RetireGrant is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Deletes a grant. Typically, you retire a grant when you no longer need its
// permissions. To identify the grant to retire, use a [grant token], or both the grant ID and
// a key identifier (key ID or key ARN) of the KMS key. The CreateGrantoperation returns both
// values.
//
// This operation can be called by the retiring principal for a grant, by the
// grantee principal if the grant allows the RetireGrant operation, and by the
// Amazon Web Services account in which the grant is created. It can also be called
// by principals to whom permission for retiring a grant is delegated.
//
// For detailed information about grants, including grant terminology, see [Grants in KMS] in the
// Key Management Service Developer Guide . For examples of creating grants in
// several programming languages, see [Use CreateGrant with an Amazon Web Services SDK or CLI].
//
// Cross-account use: Yes. You can retire a grant on a KMS key in a different
// Amazon Web Services account.
//
// Required permissions: Permission to retire a grant is determined primarily by
// the grant. For details, see [Retiring and revoking grants]in the Key Management Service Developer Guide.
//
// Related operations:
//
// # CreateGrant
//
// # ListGrants
//
// # ListRetirableGrants
//
// # RevokeGrant
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [grant token]: https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant_token
// [Retiring and revoking grants]: https://docs.aws.amazon.com/kms/latest/developerguide/grant-delete.html
// [Grants in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/grants.html
// [Use CreateGrant with an Amazon Web Services SDK or CLI]: https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_CreateGrant_section.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
