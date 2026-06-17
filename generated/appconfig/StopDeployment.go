package appconfig

// StopDeployment is generated as a reference stub.
// Executable command wiring lives under cmd/appconfig.go.
//
// Stops a deployment. This API action works only on deployments that have a
// status of DEPLOYING , unless an AllowRevert parameter is supplied. If the
// AllowRevert parameter is supplied, the status of an in-progress deployment will
// be ROLLED_BACK . The status of a completed deployment will be REVERTED .
// AppConfig only allows a revert within 72 hours of deployment completion.
