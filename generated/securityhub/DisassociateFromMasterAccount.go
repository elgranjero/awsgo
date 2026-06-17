package securityhub

// DisassociateFromMasterAccount is generated as a reference stub.
// Executable command wiring lives under cmd/securityhub.go.
//
// This method is deprecated. Instead, use DisassociateFromAdministratorAccount .
//
// The Security Hub CSPM console continues to use DisassociateFromMasterAccount .
// It will eventually change to use DisassociateFromAdministratorAccount . Any IAM
// policies that specifically control access to this function must continue to use
// DisassociateFromMasterAccount . You should also add
// DisassociateFromAdministratorAccount to your policies to ensure that the correct
// permissions are in place after the console begins to use
// DisassociateFromAdministratorAccount .
//
// Disassociates the current Security Hub CSPM member account from the associated
// administrator account.
//
// This operation is only used by accounts that are not part of an organization.
// For organization accounts, only the administrator account can disassociate a
// member account.
//
// Deprecated: This API has been deprecated, use
// DisassociateFromAdministratorAccount API instead.
