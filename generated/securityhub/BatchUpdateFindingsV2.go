package securityhub

// BatchUpdateFindingsV2 is generated as a reference stub.
// Executable command wiring lives under cmd/securityhub.go.
//
// Used by customers to update information about their investigation into a
// finding. Requested by delegated administrator accounts or member accounts.
// Delegated administrator accounts can update findings for their account and their
// member accounts. Member accounts can update findings for their account.
// BatchUpdateFindings and BatchUpdateFindingV2 both use
// securityhub:BatchUpdateFindings in the Action element of an IAM policy
// statement. You must have permission to perform the
// securityhub:BatchUpdateFindings action. Updates from BatchUpdateFindingsV2
// don't affect the value of f inding_info.modified_time ,
// finding_info.modified_time_dt , time , time_dt for a finding .
