package auditmanager

// BatchImportEvidenceToAssessmentControl is generated as a reference stub.
// Executable command wiring lives under cmd/auditmanager.go.
//
// Adds one or more pieces of evidence to a control in an Audit Manager
// assessment.
//
// You can import manual evidence from any S3 bucket by specifying the S3 URI of
// the object. You can also upload a file from your browser, or enter plain text in
// response to a risk assessment question.
//
// The following restrictions apply to this action:
//
// - manualEvidence can be only one of the following: evidenceFileName ,
// s3ResourcePath , or textResponse
//
// - Maximum size of an individual evidence file: 100 MB
//
// - Number of daily manual evidence uploads per control: 100
//
// - Supported file formats: See [Supported file types for manual evidence]in the Audit Manager User Guide
//
// For more information about Audit Manager service restrictions, see [Quotas and restrictions for Audit Manager].
//
// [Supported file types for manual evidence]: https://docs.aws.amazon.com/audit-manager/latest/userguide/upload-evidence.html#supported-manual-evidence-files
// [Quotas and restrictions for Audit Manager]: https://docs.aws.amazon.com/audit-manager/latest/userguide/service-quotas.html
