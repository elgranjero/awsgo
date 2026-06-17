package ssm

// DeleteOpsItem is generated as a reference stub.
// Executable command wiring lives under cmd/ssm.go.
//
// Delete an OpsItem. You must have permission in Identity and Access Management
// (IAM) to delete an OpsItem.
//
// Note the following important information about this operation.
//
// - Deleting an OpsItem is irreversible. You can't restore a deleted OpsItem.
//
// - This operation uses an eventual consistency model, which means the system
// can take a few minutes to complete this operation. If you delete an OpsItem and
// immediately call, for example, GetOpsItem, the deleted OpsItem might still appear in
// the response.
//
// - This operation is idempotent. The system doesn't throw an exception if you
// repeatedly call this operation for the same OpsItem. If the first call is
// successful, all additional calls return the same successful response as the
// first call.
//
// - This operation doesn't support cross-account calls. A delegated
// administrator or management account can't delete OpsItems in other accounts,
// even if OpsCenter has been set up for cross-account administration. For more
// information about cross-account administration, see [Setting up OpsCenter to centrally manage OpsItems across accounts]in the Systems Manager
// User Guide.
//
// [Setting up OpsCenter to centrally manage OpsItems across accounts]: https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-setting-up-cross-account.html
