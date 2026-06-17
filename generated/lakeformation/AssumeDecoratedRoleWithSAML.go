package lakeformation

// AssumeDecoratedRoleWithSAML is generated as a reference stub.
// Executable command wiring lives under cmd/lakeformation.go.
//
// Allows a caller to assume an IAM role decorated as the SAML user specified in
// the SAML assertion included in the request. This decoration allows Lake
// Formation to enforce access policies against the SAML users and groups. This API
// operation requires SAML federation setup in the caller’s account as it can only
// be called with valid SAML assertions. Lake Formation does not scope down the
// permission of the assumed role. All permissions attached to the role via the
// SAML federation setup will be included in the role session.
//
// This decorated role is expected to access data in Amazon S3 by getting
// temporary access from Lake Formation which is authorized via the virtual API
// GetDataAccess . Therefore, all SAML roles that can be assumed via
// AssumeDecoratedRoleWithSAML must at a minimum include
// lakeformation:GetDataAccess in their role policies. A typical IAM policy
// attached to such a role would include the following actions:
//
// - glue:*Database*
//
// - glue:*Table*
//
// - glue:*Partition*
//
// - glue:*UserDefinedFunction*
//
// - lakeformation:GetDataAccess
