package auditmanager

// DeleteAssessmentReport is generated as a reference stub.
// Executable command wiring lives under cmd/auditmanager.go.
//
// Deletes an assessment report in Audit Manager.
//
// When you run the DeleteAssessmentReport operation, Audit Manager attempts to
// delete the following data:
//
// - The specified assessment report that’s stored in your S3 bucket
//
// - The associated metadata that’s stored in Audit Manager
//
// If Audit Manager can’t access the assessment report in your S3 bucket, the
// report isn’t deleted. In this event, the DeleteAssessmentReport operation
// doesn’t fail. Instead, it proceeds to delete the associated metadata only. You
// must then delete the assessment report from the S3 bucket yourself.
//
// This scenario happens when Audit Manager receives a 403 (Forbidden) or 404 (Not
// Found) error from Amazon S3. To avoid this, make sure that your S3 bucket is
// available, and that you configured the correct permissions for Audit Manager to
// delete resources in your S3 bucket. For an example permissions policy that you
// can use, see [Assessment report destination permissions]in the Audit Manager User Guide. For information about the issues
// that could cause a 403 (Forbidden) or 404 (Not Found ) error from Amazon S3, see [List of Error Codes]
// in the Amazon Simple Storage Service API Reference.
//
// [List of Error Codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#ErrorCodeList
// [Assessment report destination permissions]: https://docs.aws.amazon.com/audit-manager/latest/userguide/security_iam_id-based-policy-examples.html#full-administrator-access-assessment-report-destination
