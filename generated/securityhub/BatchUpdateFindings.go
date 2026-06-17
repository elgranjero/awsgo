package securityhub

// BatchUpdateFindings is generated as a reference stub.
// Executable command wiring lives under cmd/securityhub.go.
//
// Used by Security Hub CSPM customers to update information about their
//
// investigation into one or more findings. Requested by administrator accounts or
// member accounts. Administrator accounts can update findings for their account
// and their member accounts. A member account can update findings only for their
// own account. Administrator and member accounts can use this operation to update
// the following fields and objects for one or more findings:
//
// - Confidence
//
// - Criticality
//
// - Note
//
// - RelatedFindings
//
// - Severity
//
// - Types
//
// - UserDefinedFields
//
// - VerificationState
//
// - Workflow
//
// If you use this operation to update a finding, your updates don’t affect the
// value for the UpdatedAt field of the finding. Also note that it can take
// several minutes for Security Hub CSPM to process your request and update each
// finding specified in the request.
//
// You can configure IAM policies to restrict access to fields and field values.
// For example, you might not want member accounts to be able to suppress findings
// or change the finding severity. For more information see [Configuring access to BatchUpdateFindings]in the Security Hub
// CSPM User Guide.
//
// [Configuring access to BatchUpdateFindings]: https://docs.aws.amazon.com/securityhub/latest/userguide/finding-update-batchupdatefindings.html#batchupdatefindings-configure-access
