package kms

// ListRetirableGrants is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Returns information about all grants in the Amazon Web Services account and
// Region that have the specified retiring principal.
//
// You can specify any principal in your Amazon Web Services account. The grants
// that are returned include grants for KMS keys in your Amazon Web Services
// account and other Amazon Web Services accounts. You might use this operation to
// determine which grants you may retire. To retire a grant, use the RetireGrantoperation.
//
// For detailed information about grants, including grant terminology, see [Grants in KMS] in the
// Key Management Service Developer Guide . For examples of creating grants in
// several programming languages, see [Use CreateGrant with an Amazon Web Services SDK or CLI].
//
// Cross-account use: You must specify a principal in your Amazon Web Services
// account. This operation returns a list of grants where the retiring principal
// specified in the ListRetirableGrants request is the same retiring principal on
// the grant. This can include grants on KMS keys owned by other Amazon Web
// Services accounts, but you do not need kms:ListRetirableGrants permission (or
// any other additional permission) in any Amazon Web Services account other than
// your own.
//
// Required permissions: [kms:ListRetirableGrants] (IAM policy) in your Amazon Web Services account.
//
// KMS authorizes ListRetirableGrants requests by evaluating the caller account's
// kms:ListRetirableGrants permissions. The authorized resource in
// ListRetirableGrants calls is the retiring principal specified in the request.
// KMS does not evaluate the caller's permissions to verify their access to any KMS
// keys or grants that might be returned by the ListRetirableGrants call.
//
// Related operations:
//
// # CreateGrant
//
// # ListGrants
//
// # RetireGrant
//
// # RevokeGrant
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [kms:ListRetirableGrants]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Grants in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/grants.html
// [Use CreateGrant with an Amazon Web Services SDK or CLI]: https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_CreateGrant_section.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
