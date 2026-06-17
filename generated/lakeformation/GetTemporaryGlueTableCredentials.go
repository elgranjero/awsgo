package lakeformation

// GetTemporaryGlueTableCredentials is generated as a reference stub.
// Executable command wiring lives under cmd/lakeformation.go.
//
// Allows a caller in a secure environment to assume a role with permission to
// access Amazon S3. In order to vend such credentials, Lake Formation assumes the
// role associated with a registered location, for example an Amazon S3 bucket,
// with a scope down policy which restricts the access to a single prefix.
//
// To call this API, the role that the service assumes must have
// lakeformation:GetDataAccess permission on the resource.
