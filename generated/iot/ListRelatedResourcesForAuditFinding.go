package iot

// ListRelatedResourcesForAuditFinding is generated as a reference stub.
// Executable command wiring lives under cmd/iot.go.
//
// The related resources of an Audit finding. The following resources can be
// returned from calling this API:
//
// - DEVICE_CERTIFICATE
//
// - CA_CERTIFICATE
//
// - IOT_POLICY
//
// - COGNITO_IDENTITY_POOL
//
// - CLIENT_ID
//
// - ACCOUNT_SETTINGS
//
// - ROLE_ALIAS
//
// - IAM_ROLE
//
// - ISSUER_CERTIFICATE
//
// This API is similar to DescribeAuditFinding's [RelatedResources] but provides pagination and is
// not limited to 10 resources. When calling [DescribeAuditFinding]for the intermediate CA revoked for
// active device certificates check, RelatedResources will not be populated. You
// must use this API, ListRelatedResourcesForAuditFinding, to list the
// certificates.
//
// [RelatedResources]: https://docs.aws.amazon.com/iot/latest/apireference/API_DescribeAuditFinding.html
// [DescribeAuditFinding]: https://docs.aws.amazon.com/iot/latest/apireference/API_DescribeAuditFinding.html
