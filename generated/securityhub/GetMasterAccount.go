package securityhub

// GetMasterAccount is generated as a reference stub.
// Executable command wiring lives under cmd/securityhub.go.
//
// This method is deprecated. Instead, use GetAdministratorAccount .
//
// The Security Hub CSPM console continues to use GetMasterAccount . It will
// eventually change to use GetAdministratorAccount . Any IAM policies that
// specifically control access to this function must continue to use
// GetMasterAccount . You should also add GetAdministratorAccount to your policies
// to ensure that the correct permissions are in place after the console begins to
// use GetAdministratorAccount .
//
// Provides the details for the Security Hub CSPM administrator account for the
// current member account.
//
// Can be used by both member accounts that are managed using Organizations and
// accounts that were invited manually.
//
// Deprecated: This API has been deprecated, use GetAdministratorAccount API
// instead.
