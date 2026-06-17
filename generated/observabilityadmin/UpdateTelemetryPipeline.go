package observabilityadmin

// UpdateTelemetryPipeline is generated as a reference stub.
// Executable command wiring lives under cmd/observabilityadmin.go.
//
// Updates the configuration of an existing telemetry pipeline.
//
// The following attributes cannot be updated after pipeline creation:
//
// - Pipeline name - The pipeline name is immutable
//
// - Pipeline ARN - The ARN is automatically generated and cannot be changed
//
// - Source type - Once a pipeline is created with a specific source type (such
// as S3, CloudWatch Logs, GitHub, or third-party sources), it cannot be changed to
// a different source type
//
// Processors can be added, removed, or modified. However, some processors are not
// supported for third-party pipelines and cannot be added through updates.
//
// # Source-Specific Update Rules
//
// CloudWatch Logs Sources (Vended and Custom)  Updatable: sts_role_arn
//
// Fixed: data_source_name , data_source_type , sink (must remain (at)original )
//
// S3 Sources (Crowdstrike, Zscaler, SentinelOne, Custom)  Updatable: All SQS
// configuration parameters, sts_role_arn , codec settings, compression type,
// bucket ownership settings, sink log group
//
// Fixed: notification_type , aws.region
//
// GitHub Audit Logs  Updatable: All Amazon Web Services Secrets Manager
// attributes, scope (can switch between ORGANIZATION/ENTERPRISE), organization or
// enterprise name, range , authentication credentials (PAT or GitHub App)
//
// Microsoft Sources (Entra ID, Office365, Windows)  Updatable: All Amazon Web
// Services Secrets Manager attributes, tenant_id , workspace_id (Windows only),
// OAuth2 credentials ( client_id , client_secret )
//
// Okta Sources (SSO, Auth0)  Updatable: All Amazon Web Services Secrets Manager
// attributes, domain , range , OAuth2 credentials ( client_id , client_secret )
//
// Palo Alto Networks  Updatable: All Amazon Web Services Secrets Manager
// attributes, hostname , basic authentication credentials ( username , password )
//
// ServiceNow CMDB  Updatable: All Amazon Web Services Secrets Manager attributes,
// instance_url , range , OAuth2 credentials ( client_id , client_secret )
//
// Wiz CNAPP  Updatable: All Amazon Web Services Secrets Manager attributes, region
// , range , OAuth2 credentials ( client_id , client_secret )
