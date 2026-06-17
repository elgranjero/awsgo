package securityhub

// UpdateFindings is generated as a reference stub.
// Executable command wiring lives under cmd/securityhub.go.
//
// UpdateFindings is a deprecated operation. Instead of UpdateFindings , use the
// BatchUpdateFindings operation.
//
// The UpdateFindings operation updates the Note and RecordState of the Security
// Hub CSPM aggregated findings that the filter attributes specify. Any member
// account that can view the finding can also see the update to the finding.
//
// Finding updates made with UpdateFindings aren't persisted if the same finding
// is later updated by the finding provider through the BatchImportFindings
// operation. In addition, Security Hub CSPM doesn't record updates made with
// UpdateFindings in the finding history.
