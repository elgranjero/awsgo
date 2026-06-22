package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/iot"
)

var fields_accept_certificate_transfer = []leanruntime.Field{
	{Name: "CertificateId", Flag: "certificate-id", Type: "*string", Required: true},
	{Name: "SetAsActive", Flag: "set-as-active", Type: "bool", Required: false},
}

var fields_add_thing_to_billing_group = []leanruntime.Field{
	{Name: "BillingGroupArn", Flag: "billing-group-arn", Type: "*string", Required: false},
	{Name: "BillingGroupName", Flag: "billing-group-name", Type: "*string", Required: false},
	{Name: "ThingArn", Flag: "thing-arn", Type: "*string", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: false},
}

var fields_add_thing_to_thing_group = []leanruntime.Field{
	{Name: "OverrideDynamicGroups", Flag: "override-dynamic-groups", Type: "bool", Required: false},
	{Name: "ThingArn", Flag: "thing-arn", Type: "*string", Required: false},
	{Name: "ThingGroupArn", Flag: "thing-group-arn", Type: "*string", Required: false},
	{Name: "ThingGroupName", Flag: "thing-group-name", Type: "*string", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: false},
}

var fields_associate_sbom_with_package_version = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PackageName", Flag: "package-name", Type: "*string", Required: true},
	{Name: "Sbom", Flag: "sbom", Type: "*types.Sbom", Required: true},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: true},
}

var fields_associate_targets_with_job = []leanruntime.Field{
	{Name: "Comment", Flag: "comment", Type: "*string", Required: false},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "NamespaceId", Flag: "namespace-id", Type: "*string", Required: false},
	{Name: "Targets", Flag: "targets", Type: "[]string", Required: true},
}

var fields_attach_policy = []leanruntime.Field{
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "Target", Flag: "target", Type: "*string", Required: true},
}

var fields_attach_principal_policy = []leanruntime.Field{
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "Principal", Flag: "principal", Type: "*string", Required: true},
}

var fields_attach_security_profile = []leanruntime.Field{
	{Name: "SecurityProfileName", Flag: "security-profile-name", Type: "*string", Required: true},
	{Name: "SecurityProfileTargetArn", Flag: "security-profile-target-arn", Type: "*string", Required: true},
}

var fields_attach_thing_principal = []leanruntime.Field{
	{Name: "Principal", Flag: "principal", Type: "*string", Required: true},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
	{Name: "ThingPrincipalType", Flag: "thing-principal-type", Type: "types.ThingPrincipalType", Required: false},
}

var fields_cancel_audit_mitigation_actions_task = []leanruntime.Field{
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_cancel_audit_task = []leanruntime.Field{
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_cancel_certificate_transfer = []leanruntime.Field{
	{Name: "CertificateId", Flag: "certificate-id", Type: "*string", Required: true},
}

var fields_cancel_detect_mitigation_actions_task = []leanruntime.Field{
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_cancel_job = []leanruntime.Field{
	{Name: "Comment", Flag: "comment", Type: "*string", Required: false},
	{Name: "Force", Flag: "force", Type: "bool", Required: false},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "ReasonCode", Flag: "reason-code", Type: "*string", Required: false},
}

var fields_cancel_job_execution = []leanruntime.Field{
	{Name: "ExpectedVersion", Flag: "expected-version", Type: "*int64", Required: false},
	{Name: "Force", Flag: "force", Type: "bool", Required: false},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "StatusDetails", Flag: "status-details", Type: "map[string]string", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

var fields_clear_default_authorizer = []leanruntime.Field{}

var fields_confirm_topic_rule_destination = []leanruntime.Field{
	{Name: "ConfirmationToken", Flag: "confirmation-token", Type: "*string", Required: true},
}

var fields_create_audit_suppression = []leanruntime.Field{
	{Name: "CheckName", Flag: "check-name", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExpirationDate", Flag: "expiration-date", Type: "*time.Time", Required: false},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*types.ResourceIdentifier", Required: true},
	{Name: "SuppressIndefinitely", Flag: "suppress-indefinitely", Type: "*bool", Required: false},
}

var fields_create_authorizer = []leanruntime.Field{
	{Name: "AuthorizerFunctionArn", Flag: "authorizer-function-arn", Type: "*string", Required: true},
	{Name: "AuthorizerName", Flag: "authorizer-name", Type: "*string", Required: true},
	{Name: "EnableCachingForHttp", Flag: "enable-caching-for-http", Type: "*bool", Required: false},
	{Name: "SigningDisabled", Flag: "signing-disabled", Type: "*bool", Required: false},
	{Name: "Status", Flag: "status", Type: "types.AuthorizerStatus", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TokenKeyName", Flag: "token-key-name", Type: "*string", Required: false},
	{Name: "TokenSigningPublicKeys", Flag: "token-signing-public-keys", Type: "map[string]string", Required: false},
}

var fields_create_billing_group = []leanruntime.Field{
	{Name: "BillingGroupName", Flag: "billing-group-name", Type: "*string", Required: true},
	{Name: "BillingGroupProperties", Flag: "billing-group-properties", Type: "*types.BillingGroupProperties", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_certificate_from_csr = []leanruntime.Field{
	{Name: "CertificateSigningRequest", Flag: "certificate-signing-request", Type: "*string", Required: true},
	{Name: "SetAsActive", Flag: "set-as-active", Type: "bool", Required: false},
}

var fields_create_certificate_provider = []leanruntime.Field{
	{Name: "AccountDefaultForOperations", Flag: "account-default-for-operations", Type: "[]types.CertificateProviderOperation", Required: true},
	{Name: "CertificateProviderName", Flag: "certificate-provider-name", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "LambdaFunctionArn", Flag: "lambda-function-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_command = []leanruntime.Field{
	{Name: "CommandId", Flag: "command-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "MandatoryParameters", Flag: "mandatory-parameters", Type: "[]types.CommandParameter", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "types.CommandNamespace", Required: false},
	{Name: "Payload", Flag: "payload", Type: "*types.CommandPayload", Required: false},
	{Name: "PayloadTemplate", Flag: "payload-template", Type: "*string", Required: false},
	{Name: "Preprocessor", Flag: "preprocessor", Type: "*types.CommandPreprocessor", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_custom_metric = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "MetricName", Flag: "metric-name", Type: "*string", Required: true},
	{Name: "MetricType", Flag: "metric-type", Type: "types.CustomMetricType", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_dimension = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "StringValues", Flag: "string-values", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "types.DimensionType", Required: true},
}

var fields_create_domain_configuration = []leanruntime.Field{
	{Name: "ApplicationProtocol", Flag: "application-protocol", Type: "types.ApplicationProtocol", Required: false},
	{Name: "AuthenticationType", Flag: "authentication-type", Type: "types.AuthenticationType", Required: false},
	{Name: "AuthorizerConfig", Flag: "authorizer-config", Type: "*types.AuthorizerConfig", Required: false},
	{Name: "ClientCertificateConfig", Flag: "client-certificate-config", Type: "*types.ClientCertificateConfig", Required: false},
	{Name: "DomainConfigurationName", Flag: "domain-configuration-name", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: false},
	{Name: "ServerCertificateArns", Flag: "server-certificate-arns", Type: "[]string", Required: false},
	{Name: "ServerCertificateConfig", Flag: "server-certificate-config", Type: "*types.ServerCertificateConfig", Required: false},
	{Name: "ServiceType", Flag: "service-type", Type: "types.ServiceType", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TlsConfig", Flag: "tls-config", Type: "*types.TlsConfig", Required: false},
	{Name: "ValidationCertificateArn", Flag: "validation-certificate-arn", Type: "*string", Required: false},
}

var fields_create_dynamic_thing_group = []leanruntime.Field{
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: false},
	{Name: "QueryString", Flag: "query-string", Type: "*string", Required: true},
	{Name: "QueryVersion", Flag: "query-version", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "ThingGroupName", Flag: "thing-group-name", Type: "*string", Required: true},
	{Name: "ThingGroupProperties", Flag: "thing-group-properties", Type: "*types.ThingGroupProperties", Required: false},
}

var fields_create_fleet_metric = []leanruntime.Field{
	{Name: "AggregationField", Flag: "aggregation-field", Type: "*string", Required: true},
	{Name: "AggregationType", Flag: "aggregation-type", Type: "*types.AggregationType", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: false},
	{Name: "MetricName", Flag: "metric-name", Type: "*string", Required: true},
	{Name: "Period", Flag: "period", Type: "*int32", Required: true},
	{Name: "QueryString", Flag: "query-string", Type: "*string", Required: true},
	{Name: "QueryVersion", Flag: "query-version", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Unit", Flag: "unit", Type: "types.FleetMetricUnit", Required: false},
}

var fields_create_job = []leanruntime.Field{
	{Name: "AbortConfig", Flag: "abort-config", Type: "*types.AbortConfig", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DestinationPackageVersions", Flag: "destination-package-versions", Type: "[]string", Required: false},
	{Name: "Document", Flag: "document", Type: "*string", Required: false},
	{Name: "DocumentParameters", Flag: "document-parameters", Type: "map[string]string", Required: false},
	{Name: "DocumentSource", Flag: "document-source", Type: "*string", Required: false},
	{Name: "JobExecutionsRetryConfig", Flag: "job-executions-retry-config", Type: "*types.JobExecutionsRetryConfig", Required: false},
	{Name: "JobExecutionsRolloutConfig", Flag: "job-executions-rollout-config", Type: "*types.JobExecutionsRolloutConfig", Required: false},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "JobTemplateArn", Flag: "job-template-arn", Type: "*string", Required: false},
	{Name: "NamespaceId", Flag: "namespace-id", Type: "*string", Required: false},
	{Name: "PresignedUrlConfig", Flag: "presigned-url-config", Type: "*types.PresignedUrlConfig", Required: false},
	{Name: "SchedulingConfig", Flag: "scheduling-config", Type: "*types.SchedulingConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetSelection", Flag: "target-selection", Type: "types.TargetSelection", Required: false},
	{Name: "Targets", Flag: "targets", Type: "[]string", Required: true},
	{Name: "TimeoutConfig", Flag: "timeout-config", Type: "*types.TimeoutConfig", Required: false},
}

var fields_create_job_template = []leanruntime.Field{
	{Name: "AbortConfig", Flag: "abort-config", Type: "*types.AbortConfig", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "DestinationPackageVersions", Flag: "destination-package-versions", Type: "[]string", Required: false},
	{Name: "Document", Flag: "document", Type: "*string", Required: false},
	{Name: "DocumentSource", Flag: "document-source", Type: "*string", Required: false},
	{Name: "JobArn", Flag: "job-arn", Type: "*string", Required: false},
	{Name: "JobExecutionsRetryConfig", Flag: "job-executions-retry-config", Type: "*types.JobExecutionsRetryConfig", Required: false},
	{Name: "JobExecutionsRolloutConfig", Flag: "job-executions-rollout-config", Type: "*types.JobExecutionsRolloutConfig", Required: false},
	{Name: "JobTemplateId", Flag: "job-template-id", Type: "*string", Required: true},
	{Name: "MaintenanceWindows", Flag: "maintenance-windows", Type: "[]types.MaintenanceWindow", Required: false},
	{Name: "PresignedUrlConfig", Flag: "presigned-url-config", Type: "*types.PresignedUrlConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TimeoutConfig", Flag: "timeout-config", Type: "*types.TimeoutConfig", Required: false},
}

var fields_create_keys_and_certificate = []leanruntime.Field{
	{Name: "SetAsActive", Flag: "set-as-active", Type: "bool", Required: false},
}

var fields_create_mitigation_action = []leanruntime.Field{
	{Name: "ActionName", Flag: "action-name", Type: "*string", Required: true},
	{Name: "ActionParams", Flag: "action-params", Type: "*types.MitigationActionParams", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_ota_update = []leanruntime.Field{
	{Name: "AdditionalParameters", Flag: "additional-parameters", Type: "map[string]string", Required: false},
	{Name: "AwsJobAbortConfig", Flag: "aws-job-abort-config", Type: "*types.AwsJobAbortConfig", Required: false},
	{Name: "AwsJobExecutionsRolloutConfig", Flag: "aws-job-executions-rollout-config", Type: "*types.AwsJobExecutionsRolloutConfig", Required: false},
	{Name: "AwsJobPresignedUrlConfig", Flag: "aws-job-presigned-url-config", Type: "*types.AwsJobPresignedUrlConfig", Required: false},
	{Name: "AwsJobTimeoutConfig", Flag: "aws-job-timeout-config", Type: "*types.AwsJobTimeoutConfig", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Files", Flag: "files", Type: "[]types.OTAUpdateFile", Required: true},
	{Name: "OtaUpdateId", Flag: "ota-update-id", Type: "*string", Required: true},
	{Name: "Protocols", Flag: "protocols", Type: "[]types.Protocol", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetSelection", Flag: "target-selection", Type: "types.TargetSelection", Required: false},
	{Name: "Targets", Flag: "targets", Type: "[]string", Required: true},
}

var fields_create_package = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "PackageName", Flag: "package-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_package_version = []leanruntime.Field{
	{Name: "Artifact", Flag: "artifact", Type: "*types.PackageVersionArtifact", Required: false},
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "PackageName", Flag: "package-name", Type: "*string", Required: true},
	{Name: "Recipe", Flag: "recipe", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: true},
}

var fields_create_policy = []leanruntime.Field{
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: true},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_policy_version = []leanruntime.Field{
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: true},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "SetAsDefault", Flag: "set-as-default", Type: "bool", Required: false},
}

var fields_create_provisioning_claim = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_create_provisioning_template = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "PreProvisioningHook", Flag: "pre-provisioning-hook", Type: "*types.ProvisioningHook", Required: false},
	{Name: "ProvisioningRoleArn", Flag: "provisioning-role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TemplateBody", Flag: "template-body", Type: "*string", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.TemplateType", Required: false},
}

var fields_create_provisioning_template_version = []leanruntime.Field{
	{Name: "SetAsDefault", Flag: "set-as-default", Type: "bool", Required: false},
	{Name: "TemplateBody", Flag: "template-body", Type: "*string", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_create_role_alias = []leanruntime.Field{
	{Name: "CredentialDurationSeconds", Flag: "credential-duration-seconds", Type: "*int32", Required: false},
	{Name: "RoleAlias", Flag: "role-alias", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_scheduled_audit = []leanruntime.Field{
	{Name: "DayOfMonth", Flag: "day-of-month", Type: "*string", Required: false},
	{Name: "DayOfWeek", Flag: "day-of-week", Type: "types.DayOfWeek", Required: false},
	{Name: "Frequency", Flag: "frequency", Type: "types.AuditFrequency", Required: true},
	{Name: "ScheduledAuditName", Flag: "scheduled-audit-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetCheckNames", Flag: "target-check-names", Type: "[]string", Required: true},
}

var fields_create_security_profile = []leanruntime.Field{
	{Name: "AdditionalMetricsToRetain", Flag: "additional-metrics-to-retain", Type: "[]string", Required: false},
	{Name: "AdditionalMetricsToRetainV2", Flag: "additional-metrics-to-retain-v2", Type: "[]types.MetricToRetain", Required: false},
	{Name: "AlertTargets", Flag: "alert-targets", Type: "map[string]types.AlertTarget", Required: false},
	{Name: "Behaviors", Flag: "behaviors", Type: "[]types.Behavior", Required: false},
	{Name: "MetricsExportConfig", Flag: "metrics-export-config", Type: "*types.MetricsExportConfig", Required: false},
	{Name: "SecurityProfileDescription", Flag: "security-profile-description", Type: "*string", Required: false},
	{Name: "SecurityProfileName", Flag: "security-profile-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_stream = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Files", Flag: "files", Type: "[]types.StreamFile", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_thing = []leanruntime.Field{
	{Name: "AttributePayload", Flag: "attribute-payload", Type: "*types.AttributePayload", Required: false},
	{Name: "BillingGroupName", Flag: "billing-group-name", Type: "*string", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
	{Name: "ThingTypeName", Flag: "thing-type-name", Type: "*string", Required: false},
}

var fields_create_thing_group = []leanruntime.Field{
	{Name: "ParentGroupName", Flag: "parent-group-name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "ThingGroupName", Flag: "thing-group-name", Type: "*string", Required: true},
	{Name: "ThingGroupProperties", Flag: "thing-group-properties", Type: "*types.ThingGroupProperties", Required: false},
}

var fields_create_thing_type = []leanruntime.Field{
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "ThingTypeName", Flag: "thing-type-name", Type: "*string", Required: true},
	{Name: "ThingTypeProperties", Flag: "thing-type-properties", Type: "*types.ThingTypeProperties", Required: false},
}

var fields_create_topic_rule = []leanruntime.Field{
	{Name: "RuleName", Flag: "rule-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "*string", Required: false},
	{Name: "TopicRulePayload", Flag: "topic-rule-payload", Type: "*types.TopicRulePayload", Required: true},
}

var fields_create_topic_rule_destination = []leanruntime.Field{
	{Name: "DestinationConfiguration", Flag: "destination-configuration", Type: "*types.TopicRuleDestinationConfiguration", Required: true},
}

var fields_delete_account_audit_configuration = []leanruntime.Field{
	{Name: "DeleteScheduledAudits", Flag: "delete-scheduled-audits", Type: "bool", Required: false},
}

var fields_delete_audit_suppression = []leanruntime.Field{
	{Name: "CheckName", Flag: "check-name", Type: "*string", Required: true},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*types.ResourceIdentifier", Required: true},
}

var fields_delete_authorizer = []leanruntime.Field{
	{Name: "AuthorizerName", Flag: "authorizer-name", Type: "*string", Required: true},
}

var fields_delete_billing_group = []leanruntime.Field{
	{Name: "BillingGroupName", Flag: "billing-group-name", Type: "*string", Required: true},
	{Name: "ExpectedVersion", Flag: "expected-version", Type: "*int64", Required: false},
}

var fields_delete_ca_certificate = []leanruntime.Field{
	{Name: "CertificateId", Flag: "certificate-id", Type: "*string", Required: true},
}

var fields_delete_certificate = []leanruntime.Field{
	{Name: "CertificateId", Flag: "certificate-id", Type: "*string", Required: true},
	{Name: "ForceDelete", Flag: "force-delete", Type: "bool", Required: false},
}

var fields_delete_certificate_provider = []leanruntime.Field{
	{Name: "CertificateProviderName", Flag: "certificate-provider-name", Type: "*string", Required: true},
}

var fields_delete_command = []leanruntime.Field{
	{Name: "CommandId", Flag: "command-id", Type: "*string", Required: true},
}

var fields_delete_command_execution = []leanruntime.Field{
	{Name: "ExecutionId", Flag: "execution-id", Type: "*string", Required: true},
	{Name: "TargetArn", Flag: "target-arn", Type: "*string", Required: true},
}

var fields_delete_custom_metric = []leanruntime.Field{
	{Name: "MetricName", Flag: "metric-name", Type: "*string", Required: true},
}

var fields_delete_dimension = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_domain_configuration = []leanruntime.Field{
	{Name: "DomainConfigurationName", Flag: "domain-configuration-name", Type: "*string", Required: true},
}

var fields_delete_dynamic_thing_group = []leanruntime.Field{
	{Name: "ExpectedVersion", Flag: "expected-version", Type: "*int64", Required: false},
	{Name: "ThingGroupName", Flag: "thing-group-name", Type: "*string", Required: true},
}

var fields_delete_fleet_metric = []leanruntime.Field{
	{Name: "ExpectedVersion", Flag: "expected-version", Type: "*int64", Required: false},
	{Name: "MetricName", Flag: "metric-name", Type: "*string", Required: true},
}

var fields_delete_job = []leanruntime.Field{
	{Name: "Force", Flag: "force", Type: "bool", Required: false},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "NamespaceId", Flag: "namespace-id", Type: "*string", Required: false},
}

var fields_delete_job_execution = []leanruntime.Field{
	{Name: "ExecutionNumber", Flag: "execution-number", Type: "*int64", Required: true},
	{Name: "Force", Flag: "force", Type: "bool", Required: false},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "NamespaceId", Flag: "namespace-id", Type: "*string", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

var fields_delete_job_template = []leanruntime.Field{
	{Name: "JobTemplateId", Flag: "job-template-id", Type: "*string", Required: true},
}

var fields_delete_mitigation_action = []leanruntime.Field{
	{Name: "ActionName", Flag: "action-name", Type: "*string", Required: true},
}

var fields_delete_ota_update = []leanruntime.Field{
	{Name: "DeleteStream", Flag: "delete-stream", Type: "bool", Required: false},
	{Name: "ForceDeleteAWSJob", Flag: "force-delete-aws-job", Type: "bool", Required: false},
	{Name: "OtaUpdateId", Flag: "ota-update-id", Type: "*string", Required: true},
}

var fields_delete_package = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PackageName", Flag: "package-name", Type: "*string", Required: true},
}

var fields_delete_package_version = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PackageName", Flag: "package-name", Type: "*string", Required: true},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: true},
}

var fields_delete_policy = []leanruntime.Field{
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
}

var fields_delete_policy_version = []leanruntime.Field{
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "PolicyVersionId", Flag: "policy-version-id", Type: "*string", Required: true},
}

var fields_delete_provisioning_template = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_delete_provisioning_template_version = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "VersionId", Flag: "version-id", Type: "*int32", Required: true},
}

var fields_delete_registration_code = []leanruntime.Field{}

var fields_delete_role_alias = []leanruntime.Field{
	{Name: "RoleAlias", Flag: "role-alias", Type: "*string", Required: true},
}

var fields_delete_scheduled_audit = []leanruntime.Field{
	{Name: "ScheduledAuditName", Flag: "scheduled-audit-name", Type: "*string", Required: true},
}

var fields_delete_security_profile = []leanruntime.Field{
	{Name: "ExpectedVersion", Flag: "expected-version", Type: "*int64", Required: false},
	{Name: "SecurityProfileName", Flag: "security-profile-name", Type: "*string", Required: true},
}

var fields_delete_stream = []leanruntime.Field{
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: true},
}

var fields_delete_thing = []leanruntime.Field{
	{Name: "ExpectedVersion", Flag: "expected-version", Type: "*int64", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

var fields_delete_thing_group = []leanruntime.Field{
	{Name: "ExpectedVersion", Flag: "expected-version", Type: "*int64", Required: false},
	{Name: "ThingGroupName", Flag: "thing-group-name", Type: "*string", Required: true},
}

var fields_delete_thing_type = []leanruntime.Field{
	{Name: "ThingTypeName", Flag: "thing-type-name", Type: "*string", Required: true},
}

var fields_delete_topic_rule = []leanruntime.Field{
	{Name: "RuleName", Flag: "rule-name", Type: "*string", Required: true},
}

var fields_delete_topic_rule_destination = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_v2_logging_level = []leanruntime.Field{
	{Name: "TargetName", Flag: "target-name", Type: "*string", Required: true},
	{Name: "TargetType", Flag: "target-type", Type: "types.LogTargetType", Required: true},
}

var fields_deprecate_thing_type = []leanruntime.Field{
	{Name: "ThingTypeName", Flag: "thing-type-name", Type: "*string", Required: true},
	{Name: "UndoDeprecate", Flag: "undo-deprecate", Type: "bool", Required: false},
}

var fields_describe_account_audit_configuration = []leanruntime.Field{}

var fields_describe_audit_finding = []leanruntime.Field{
	{Name: "FindingId", Flag: "finding-id", Type: "*string", Required: true},
}

var fields_describe_audit_mitigation_actions_task = []leanruntime.Field{
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_describe_audit_suppression = []leanruntime.Field{
	{Name: "CheckName", Flag: "check-name", Type: "*string", Required: true},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*types.ResourceIdentifier", Required: true},
}

var fields_describe_audit_task = []leanruntime.Field{
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_describe_authorizer = []leanruntime.Field{
	{Name: "AuthorizerName", Flag: "authorizer-name", Type: "*string", Required: true},
}

var fields_describe_billing_group = []leanruntime.Field{
	{Name: "BillingGroupName", Flag: "billing-group-name", Type: "*string", Required: true},
}

var fields_describe_ca_certificate = []leanruntime.Field{
	{Name: "CertificateId", Flag: "certificate-id", Type: "*string", Required: true},
}

var fields_describe_certificate = []leanruntime.Field{
	{Name: "CertificateId", Flag: "certificate-id", Type: "*string", Required: true},
}

var fields_describe_certificate_provider = []leanruntime.Field{
	{Name: "CertificateProviderName", Flag: "certificate-provider-name", Type: "*string", Required: true},
}

var fields_describe_custom_metric = []leanruntime.Field{
	{Name: "MetricName", Flag: "metric-name", Type: "*string", Required: true},
}

var fields_describe_default_authorizer = []leanruntime.Field{}

var fields_describe_detect_mitigation_actions_task = []leanruntime.Field{
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_describe_dimension = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_domain_configuration = []leanruntime.Field{
	{Name: "DomainConfigurationName", Flag: "domain-configuration-name", Type: "*string", Required: true},
}

var fields_describe_encryption_configuration = []leanruntime.Field{}

var fields_describe_endpoint = []leanruntime.Field{
	{Name: "EndpointType", Flag: "endpoint-type", Type: "*string", Required: false},
}

var fields_describe_event_configurations = []leanruntime.Field{}

var fields_describe_fleet_metric = []leanruntime.Field{
	{Name: "MetricName", Flag: "metric-name", Type: "*string", Required: true},
}

var fields_describe_index = []leanruntime.Field{
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: true},
}

var fields_describe_job = []leanruntime.Field{
	{Name: "BeforeSubstitution", Flag: "before-substitution", Type: "bool", Required: false},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_describe_job_execution = []leanruntime.Field{
	{Name: "ExecutionNumber", Flag: "execution-number", Type: "*int64", Required: false},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

var fields_describe_job_template = []leanruntime.Field{
	{Name: "JobTemplateId", Flag: "job-template-id", Type: "*string", Required: true},
}

var fields_describe_managed_job_template = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "TemplateVersion", Flag: "template-version", Type: "*string", Required: false},
}

var fields_describe_mitigation_action = []leanruntime.Field{
	{Name: "ActionName", Flag: "action-name", Type: "*string", Required: true},
}

var fields_describe_provisioning_template = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_describe_provisioning_template_version = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "VersionId", Flag: "version-id", Type: "*int32", Required: true},
}

var fields_describe_role_alias = []leanruntime.Field{
	{Name: "RoleAlias", Flag: "role-alias", Type: "*string", Required: true},
}

var fields_describe_scheduled_audit = []leanruntime.Field{
	{Name: "ScheduledAuditName", Flag: "scheduled-audit-name", Type: "*string", Required: true},
}

var fields_describe_security_profile = []leanruntime.Field{
	{Name: "SecurityProfileName", Flag: "security-profile-name", Type: "*string", Required: true},
}

var fields_describe_stream = []leanruntime.Field{
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: true},
}

var fields_describe_thing = []leanruntime.Field{
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

var fields_describe_thing_group = []leanruntime.Field{
	{Name: "ThingGroupName", Flag: "thing-group-name", Type: "*string", Required: true},
}

var fields_describe_thing_registration_task = []leanruntime.Field{
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_describe_thing_type = []leanruntime.Field{
	{Name: "ThingTypeName", Flag: "thing-type-name", Type: "*string", Required: true},
}

var fields_detach_policy = []leanruntime.Field{
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "Target", Flag: "target", Type: "*string", Required: true},
}

var fields_detach_principal_policy = []leanruntime.Field{
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "Principal", Flag: "principal", Type: "*string", Required: true},
}

var fields_detach_security_profile = []leanruntime.Field{
	{Name: "SecurityProfileName", Flag: "security-profile-name", Type: "*string", Required: true},
	{Name: "SecurityProfileTargetArn", Flag: "security-profile-target-arn", Type: "*string", Required: true},
}

var fields_detach_thing_principal = []leanruntime.Field{
	{Name: "Principal", Flag: "principal", Type: "*string", Required: true},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

var fields_disable_topic_rule = []leanruntime.Field{
	{Name: "RuleName", Flag: "rule-name", Type: "*string", Required: true},
}

var fields_disassociate_sbom_from_package_version = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PackageName", Flag: "package-name", Type: "*string", Required: true},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: true},
}

var fields_enable_topic_rule = []leanruntime.Field{
	{Name: "RuleName", Flag: "rule-name", Type: "*string", Required: true},
}

var fields_get_behavior_model_training_summaries = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SecurityProfileName", Flag: "security-profile-name", Type: "*string", Required: false},
}

var fields_get_buckets_aggregation = []leanruntime.Field{
	{Name: "AggregationField", Flag: "aggregation-field", Type: "*string", Required: true},
	{Name: "BucketsAggregationType", Flag: "buckets-aggregation-type", Type: "*types.BucketsAggregationType", Required: true},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: false},
	{Name: "QueryString", Flag: "query-string", Type: "*string", Required: true},
	{Name: "QueryVersion", Flag: "query-version", Type: "*string", Required: false},
}

var fields_get_cardinality = []leanruntime.Field{
	{Name: "AggregationField", Flag: "aggregation-field", Type: "*string", Required: false},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: false},
	{Name: "QueryString", Flag: "query-string", Type: "*string", Required: true},
	{Name: "QueryVersion", Flag: "query-version", Type: "*string", Required: false},
}

var fields_get_command = []leanruntime.Field{
	{Name: "CommandId", Flag: "command-id", Type: "*string", Required: true},
}

var fields_get_command_execution = []leanruntime.Field{
	{Name: "ExecutionId", Flag: "execution-id", Type: "*string", Required: true},
	{Name: "IncludeResult", Flag: "include-result", Type: "*bool", Required: false},
	{Name: "TargetArn", Flag: "target-arn", Type: "*string", Required: true},
}

var fields_get_effective_policies = []leanruntime.Field{
	{Name: "CognitoIdentityPoolId", Flag: "cognito-identity-pool-id", Type: "*string", Required: false},
	{Name: "Principal", Flag: "principal", Type: "*string", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: false},
}

var fields_get_indexing_configuration = []leanruntime.Field{}

var fields_get_job_document = []leanruntime.Field{
	{Name: "BeforeSubstitution", Flag: "before-substitution", Type: "bool", Required: false},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_logging_options = []leanruntime.Field{}

var fields_get_ota_update = []leanruntime.Field{
	{Name: "OtaUpdateId", Flag: "ota-update-id", Type: "*string", Required: true},
}

var fields_get_package = []leanruntime.Field{
	{Name: "PackageName", Flag: "package-name", Type: "*string", Required: true},
}

var fields_get_package_configuration = []leanruntime.Field{}

var fields_get_package_version = []leanruntime.Field{
	{Name: "PackageName", Flag: "package-name", Type: "*string", Required: true},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: true},
}

var fields_get_percentiles = []leanruntime.Field{
	{Name: "AggregationField", Flag: "aggregation-field", Type: "*string", Required: false},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: false},
	{Name: "Percents", Flag: "percents", Type: "[]float64", Required: false},
	{Name: "QueryString", Flag: "query-string", Type: "*string", Required: true},
	{Name: "QueryVersion", Flag: "query-version", Type: "*string", Required: false},
}

var fields_get_policy = []leanruntime.Field{
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
}

var fields_get_policy_version = []leanruntime.Field{
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "PolicyVersionId", Flag: "policy-version-id", Type: "*string", Required: true},
}

var fields_get_registration_code = []leanruntime.Field{}

var fields_get_statistics = []leanruntime.Field{
	{Name: "AggregationField", Flag: "aggregation-field", Type: "*string", Required: false},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: false},
	{Name: "QueryString", Flag: "query-string", Type: "*string", Required: true},
	{Name: "QueryVersion", Flag: "query-version", Type: "*string", Required: false},
}

var fields_get_thing_connectivity_data = []leanruntime.Field{
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

var fields_get_topic_rule = []leanruntime.Field{
	{Name: "RuleName", Flag: "rule-name", Type: "*string", Required: true},
}

var fields_get_topic_rule_destination = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_v2_logging_options = []leanruntime.Field{
	{Name: "Verbose", Flag: "verbose", Type: "bool", Required: false},
}

var fields_list_active_violations = []leanruntime.Field{
	{Name: "BehaviorCriteriaType", Flag: "behavior-criteria-type", Type: "types.BehaviorCriteriaType", Required: false},
	{Name: "ListSuppressedAlerts", Flag: "list-suppressed-alerts", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SecurityProfileName", Flag: "security-profile-name", Type: "*string", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: false},
	{Name: "VerificationState", Flag: "verification-state", Type: "types.VerificationState", Required: false},
}

var fields_list_attached_policies = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "Recursive", Flag: "recursive", Type: "bool", Required: false},
	{Name: "Target", Flag: "target", Type: "*string", Required: true},
}

var fields_list_audit_findings = []leanruntime.Field{
	{Name: "CheckName", Flag: "check-name", Type: "*string", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "ListSuppressedFindings", Flag: "list-suppressed-findings", Type: "bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*types.ResourceIdentifier", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: false},
}

var fields_list_audit_mitigation_actions_executions = []leanruntime.Field{
	{Name: "ActionStatus", Flag: "action-status", Type: "types.AuditMitigationActionsExecutionStatus", Required: false},
	{Name: "FindingId", Flag: "finding-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_list_audit_mitigation_actions_tasks = []leanruntime.Field{
	{Name: "AuditTaskId", Flag: "audit-task-id", Type: "*string", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "FindingId", Flag: "finding-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
	{Name: "TaskStatus", Flag: "task-status", Type: "types.AuditMitigationActionsTaskStatus", Required: false},
}

var fields_list_audit_suppressions = []leanruntime.Field{
	{Name: "AscendingOrder", Flag: "ascending-order", Type: "bool", Required: false},
	{Name: "CheckName", Flag: "check-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*types.ResourceIdentifier", Required: false},
}

var fields_list_audit_tasks = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
	{Name: "TaskStatus", Flag: "task-status", Type: "types.AuditTaskStatus", Required: false},
	{Name: "TaskType", Flag: "task-type", Type: "types.AuditTaskType", Required: false},
}

var fields_list_authorizers = []leanruntime.Field{
	{Name: "AscendingOrder", Flag: "ascending-order", Type: "bool", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "Status", Flag: "status", Type: "types.AuthorizerStatus", Required: false},
}

var fields_list_billing_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NamePrefixFilter", Flag: "name-prefix-filter", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_ca_certificates = []leanruntime.Field{
	{Name: "AscendingOrder", Flag: "ascending-order", Type: "bool", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: false},
}

var fields_list_certificate_providers = []leanruntime.Field{
	{Name: "AscendingOrder", Flag: "ascending-order", Type: "bool", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_certificates = []leanruntime.Field{
	{Name: "AscendingOrder", Flag: "ascending-order", Type: "bool", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_certificates_by_ca = []leanruntime.Field{
	{Name: "AscendingOrder", Flag: "ascending-order", Type: "bool", Required: false},
	{Name: "CaCertificateId", Flag: "ca-certificate-id", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_command_executions = []leanruntime.Field{
	{Name: "CommandArn", Flag: "command-arn", Type: "*string", Required: false},
	{Name: "CompletedTimeFilter", Flag: "completed-time-filter", Type: "*types.TimeFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "types.CommandNamespace", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "StartedTimeFilter", Flag: "started-time-filter", Type: "*types.TimeFilter", Required: false},
	{Name: "Status", Flag: "status", Type: "types.CommandExecutionStatus", Required: false},
	{Name: "TargetArn", Flag: "target-arn", Type: "*string", Required: false},
}

var fields_list_commands = []leanruntime.Field{
	{Name: "CommandParameterName", Flag: "command-parameter-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "types.CommandNamespace", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_custom_metrics = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_detect_mitigation_actions_executions = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: false},
	{Name: "ViolationId", Flag: "violation-id", Type: "*string", Required: false},
}

var fields_list_detect_mitigation_actions_tasks = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_list_dimensions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_domain_configurations = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "ServiceType", Flag: "service-type", Type: "types.ServiceType", Required: false},
}

var fields_list_fleet_metrics = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_indices = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_job_executions_for_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.JobExecutionStatus", Required: false},
}

var fields_list_job_executions_for_thing = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NamespaceId", Flag: "namespace-id", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.JobExecutionStatus", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

var fields_list_job_templates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_jobs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NamespaceId", Flag: "namespace-id", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.JobStatus", Required: false},
	{Name: "TargetSelection", Flag: "target-selection", Type: "types.TargetSelection", Required: false},
	{Name: "ThingGroupId", Flag: "thing-group-id", Type: "*string", Required: false},
	{Name: "ThingGroupName", Flag: "thing-group-name", Type: "*string", Required: false},
}

var fields_list_managed_job_templates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: false},
}

var fields_list_metric_values = []leanruntime.Field{
	{Name: "DimensionName", Flag: "dimension-name", Type: "*string", Required: false},
	{Name: "DimensionValueOperator", Flag: "dimension-value-operator", Type: "types.DimensionValueOperator", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MetricName", Flag: "metric-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

var fields_list_mitigation_actions = []leanruntime.Field{
	{Name: "ActionType", Flag: "action-type", Type: "types.MitigationActionType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_ota_updates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OtaUpdateStatus", Flag: "ota-update-status", Type: "types.OTAUpdateStatus", Required: false},
}

var fields_list_outgoing_certificates = []leanruntime.Field{
	{Name: "AscendingOrder", Flag: "ascending-order", Type: "bool", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_package_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PackageName", Flag: "package-name", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.PackageVersionStatus", Required: false},
}

var fields_list_packages = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_policies = []leanruntime.Field{
	{Name: "AscendingOrder", Flag: "ascending-order", Type: "bool", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_policy_principals = []leanruntime.Field{
	{Name: "AscendingOrder", Flag: "ascending-order", Type: "bool", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
}

var fields_list_policy_versions = []leanruntime.Field{
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
}

var fields_list_principal_policies = []leanruntime.Field{
	{Name: "AscendingOrder", Flag: "ascending-order", Type: "bool", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "Principal", Flag: "principal", Type: "*string", Required: true},
}

var fields_list_principal_things = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Principal", Flag: "principal", Type: "*string", Required: true},
}

var fields_list_principal_things_v2 = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Principal", Flag: "principal", Type: "*string", Required: true},
	{Name: "ThingPrincipalType", Flag: "thing-principal-type", Type: "types.ThingPrincipalType", Required: false},
}

var fields_list_provisioning_template_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_list_provisioning_templates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_related_resources_for_audit_finding = []leanruntime.Field{
	{Name: "FindingId", Flag: "finding-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_role_aliases = []leanruntime.Field{
	{Name: "AscendingOrder", Flag: "ascending-order", Type: "bool", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_sbom_validation_results = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PackageName", Flag: "package-name", Type: "*string", Required: true},
	{Name: "ValidationResult", Flag: "validation-result", Type: "types.SbomValidationResult", Required: false},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: true},
}

var fields_list_scheduled_audits = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_security_profiles = []leanruntime.Field{
	{Name: "DimensionName", Flag: "dimension-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MetricName", Flag: "metric-name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_security_profiles_for_target = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Recursive", Flag: "recursive", Type: "bool", Required: false},
	{Name: "SecurityProfileTargetArn", Flag: "security-profile-target-arn", Type: "*string", Required: true},
}

var fields_list_streams = []leanruntime.Field{
	{Name: "AscendingOrder", Flag: "ascending-order", Type: "bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_targets_for_policy = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
}

var fields_list_targets_for_security_profile = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SecurityProfileName", Flag: "security-profile-name", Type: "*string", Required: true},
}

var fields_list_thing_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NamePrefixFilter", Flag: "name-prefix-filter", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ParentGroup", Flag: "parent-group", Type: "*string", Required: false},
	{Name: "Recursive", Flag: "recursive", Type: "*bool", Required: false},
}

var fields_list_thing_groups_for_thing = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

var fields_list_thing_principals = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

var fields_list_thing_principals_v2 = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
	{Name: "ThingPrincipalType", Flag: "thing-principal-type", Type: "types.ThingPrincipalType", Required: false},
}

var fields_list_thing_registration_task_reports = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReportType", Flag: "report-type", Type: "types.ReportType", Required: true},
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_list_thing_registration_tasks = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.Status", Required: false},
}

var fields_list_thing_types = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ThingTypeName", Flag: "thing-type-name", Type: "*string", Required: false},
}

var fields_list_things = []leanruntime.Field{
	{Name: "AttributeName", Flag: "attribute-name", Type: "*string", Required: false},
	{Name: "AttributeValue", Flag: "attribute-value", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ThingTypeName", Flag: "thing-type-name", Type: "*string", Required: false},
	{Name: "UsePrefixAttributeValue", Flag: "use-prefix-attribute-value", Type: "bool", Required: false},
}

var fields_list_things_in_billing_group = []leanruntime.Field{
	{Name: "BillingGroupName", Flag: "billing-group-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_things_in_thing_group = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Recursive", Flag: "recursive", Type: "bool", Required: false},
	{Name: "ThingGroupName", Flag: "thing-group-name", Type: "*string", Required: true},
}

var fields_list_topic_rule_destinations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_topic_rules = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RuleDisabled", Flag: "rule-disabled", Type: "*bool", Required: false},
	{Name: "Topic", Flag: "topic", Type: "*string", Required: false},
}

var fields_list_v2_logging_levels = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TargetType", Flag: "target-type", Type: "types.LogTargetType", Required: false},
}

var fields_list_violation_events = []leanruntime.Field{
	{Name: "BehaviorCriteriaType", Flag: "behavior-criteria-type", Type: "types.BehaviorCriteriaType", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "ListSuppressedAlerts", Flag: "list-suppressed-alerts", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SecurityProfileName", Flag: "security-profile-name", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: false},
	{Name: "VerificationState", Flag: "verification-state", Type: "types.VerificationState", Required: false},
}

var fields_put_verification_state_on_violation = []leanruntime.Field{
	{Name: "VerificationState", Flag: "verification-state", Type: "types.VerificationState", Required: true},
	{Name: "VerificationStateDescription", Flag: "verification-state-description", Type: "*string", Required: false},
	{Name: "ViolationId", Flag: "violation-id", Type: "*string", Required: true},
}

var fields_register_ca_certificate = []leanruntime.Field{
	{Name: "AllowAutoRegistration", Flag: "allow-auto-registration", Type: "bool", Required: false},
	{Name: "CaCertificate", Flag: "ca-certificate", Type: "*string", Required: true},
	{Name: "CertificateMode", Flag: "certificate-mode", Type: "types.CertificateMode", Required: false},
	{Name: "RegistrationConfig", Flag: "registration-config", Type: "*types.RegistrationConfig", Required: false},
	{Name: "SetAsActive", Flag: "set-as-active", Type: "bool", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VerificationCertificate", Flag: "verification-certificate", Type: "*string", Required: false},
}

var fields_register_certificate = []leanruntime.Field{
	{Name: "CaCertificatePem", Flag: "ca-certificate-pem", Type: "*string", Required: false},
	{Name: "CertificatePem", Flag: "certificate-pem", Type: "*string", Required: true},
	{Name: "SetAsActive", Flag: "set-as-active", Type: "*bool", Required: false},
	{Name: "Status", Flag: "status", Type: "types.CertificateStatus", Required: false},
}

var fields_register_certificate_without_ca = []leanruntime.Field{
	{Name: "CertificatePem", Flag: "certificate-pem", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.CertificateStatus", Required: false},
}

var fields_register_thing = []leanruntime.Field{
	{Name: "Parameters", Flag: "parameters", Type: "map[string]string", Required: false},
	{Name: "TemplateBody", Flag: "template-body", Type: "*string", Required: true},
}

var fields_reject_certificate_transfer = []leanruntime.Field{
	{Name: "CertificateId", Flag: "certificate-id", Type: "*string", Required: true},
	{Name: "RejectReason", Flag: "reject-reason", Type: "*string", Required: false},
}

var fields_remove_thing_from_billing_group = []leanruntime.Field{
	{Name: "BillingGroupArn", Flag: "billing-group-arn", Type: "*string", Required: false},
	{Name: "BillingGroupName", Flag: "billing-group-name", Type: "*string", Required: false},
	{Name: "ThingArn", Flag: "thing-arn", Type: "*string", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: false},
}

var fields_remove_thing_from_thing_group = []leanruntime.Field{
	{Name: "ThingArn", Flag: "thing-arn", Type: "*string", Required: false},
	{Name: "ThingGroupArn", Flag: "thing-group-arn", Type: "*string", Required: false},
	{Name: "ThingGroupName", Flag: "thing-group-name", Type: "*string", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: false},
}

var fields_replace_topic_rule = []leanruntime.Field{
	{Name: "RuleName", Flag: "rule-name", Type: "*string", Required: true},
	{Name: "TopicRulePayload", Flag: "topic-rule-payload", Type: "*types.TopicRulePayload", Required: true},
}

var fields_search_index = []leanruntime.Field{
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueryString", Flag: "query-string", Type: "*string", Required: true},
	{Name: "QueryVersion", Flag: "query-version", Type: "*string", Required: false},
}

var fields_set_default_authorizer = []leanruntime.Field{
	{Name: "AuthorizerName", Flag: "authorizer-name", Type: "*string", Required: true},
}

var fields_set_default_policy_version = []leanruntime.Field{
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "PolicyVersionId", Flag: "policy-version-id", Type: "*string", Required: true},
}

var fields_set_logging_options = []leanruntime.Field{
	{Name: "LoggingOptionsPayload", Flag: "logging-options-payload", Type: "*types.LoggingOptionsPayload", Required: true},
}

var fields_set_v2_logging_level = []leanruntime.Field{
	{Name: "LogLevel", Flag: "log-level", Type: "types.LogLevel", Required: true},
	{Name: "LogTarget", Flag: "log-target", Type: "*types.LogTarget", Required: true},
}

var fields_set_v2_logging_options = []leanruntime.Field{
	{Name: "DefaultLogLevel", Flag: "default-log-level", Type: "types.LogLevel", Required: false},
	{Name: "DisableAllLogs", Flag: "disable-all-logs", Type: "bool", Required: false},
	{Name: "EventConfigurations", Flag: "event-configurations", Type: "[]types.LogEventConfiguration", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_start_audit_mitigation_actions_task = []leanruntime.Field{
	{Name: "AuditCheckToActionsMapping", Flag: "audit-check-to-actions-mapping", Type: "map[string][]string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "Target", Flag: "target", Type: "*types.AuditMitigationActionsTaskTarget", Required: true},
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_start_detect_mitigation_actions_task = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "[]string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "IncludeOnlyActiveViolations", Flag: "include-only-active-violations", Type: "*bool", Required: false},
	{Name: "IncludeSuppressedAlerts", Flag: "include-suppressed-alerts", Type: "*bool", Required: false},
	{Name: "Target", Flag: "target", Type: "*types.DetectMitigationActionsTaskTarget", Required: true},
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
	{Name: "ViolationEventOccurrenceRange", Flag: "violation-event-occurrence-range", Type: "*types.ViolationEventOccurrenceRange", Required: false},
}

var fields_start_on_demand_audit_task = []leanruntime.Field{
	{Name: "TargetCheckNames", Flag: "target-check-names", Type: "[]string", Required: true},
}

var fields_start_thing_registration_task = []leanruntime.Field{
	{Name: "InputFileBucket", Flag: "input-file-bucket", Type: "*string", Required: true},
	{Name: "InputFileKey", Flag: "input-file-key", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "TemplateBody", Flag: "template-body", Type: "*string", Required: true},
}

var fields_stop_thing_registration_task = []leanruntime.Field{
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_test_authorization = []leanruntime.Field{
	{Name: "AuthInfos", Flag: "auth-infos", Type: "[]types.AuthInfo", Required: true},
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: false},
	{Name: "CognitoIdentityPoolId", Flag: "cognito-identity-pool-id", Type: "*string", Required: false},
	{Name: "PolicyNamesToAdd", Flag: "policy-names-to-add", Type: "[]string", Required: false},
	{Name: "PolicyNamesToSkip", Flag: "policy-names-to-skip", Type: "[]string", Required: false},
	{Name: "Principal", Flag: "principal", Type: "*string", Required: false},
}

var fields_test_invoke_authorizer = []leanruntime.Field{
	{Name: "AuthorizerName", Flag: "authorizer-name", Type: "*string", Required: true},
	{Name: "HttpContext", Flag: "http-context", Type: "*types.HttpContext", Required: false},
	{Name: "MqttContext", Flag: "mqtt-context", Type: "*types.MqttContext", Required: false},
	{Name: "TlsContext", Flag: "tls-context", Type: "*types.TlsContext", Required: false},
	{Name: "Token", Flag: "token", Type: "*string", Required: false},
	{Name: "TokenSignature", Flag: "token-signature", Type: "*string", Required: false},
}

var fields_transfer_certificate = []leanruntime.Field{
	{Name: "CertificateId", Flag: "certificate-id", Type: "*string", Required: true},
	{Name: "TargetAwsAccount", Flag: "target-aws-account", Type: "*string", Required: true},
	{Name: "TransferMessage", Flag: "transfer-message", Type: "*string", Required: false},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_account_audit_configuration = []leanruntime.Field{
	{Name: "AuditCheckConfigurations", Flag: "audit-check-configurations", Type: "map[string]types.AuditCheckConfiguration", Required: false},
	{Name: "AuditNotificationTargetConfigurations", Flag: "audit-notification-target-configurations", Type: "map[string]types.AuditNotificationTarget", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_update_audit_suppression = []leanruntime.Field{
	{Name: "CheckName", Flag: "check-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExpirationDate", Flag: "expiration-date", Type: "*time.Time", Required: false},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*types.ResourceIdentifier", Required: true},
	{Name: "SuppressIndefinitely", Flag: "suppress-indefinitely", Type: "*bool", Required: false},
}

var fields_update_authorizer = []leanruntime.Field{
	{Name: "AuthorizerFunctionArn", Flag: "authorizer-function-arn", Type: "*string", Required: false},
	{Name: "AuthorizerName", Flag: "authorizer-name", Type: "*string", Required: true},
	{Name: "EnableCachingForHttp", Flag: "enable-caching-for-http", Type: "*bool", Required: false},
	{Name: "Status", Flag: "status", Type: "types.AuthorizerStatus", Required: false},
	{Name: "TokenKeyName", Flag: "token-key-name", Type: "*string", Required: false},
	{Name: "TokenSigningPublicKeys", Flag: "token-signing-public-keys", Type: "map[string]string", Required: false},
}

var fields_update_billing_group = []leanruntime.Field{
	{Name: "BillingGroupName", Flag: "billing-group-name", Type: "*string", Required: true},
	{Name: "BillingGroupProperties", Flag: "billing-group-properties", Type: "*types.BillingGroupProperties", Required: true},
	{Name: "ExpectedVersion", Flag: "expected-version", Type: "*int64", Required: false},
}

var fields_update_ca_certificate = []leanruntime.Field{
	{Name: "CertificateId", Flag: "certificate-id", Type: "*string", Required: true},
	{Name: "NewAutoRegistrationStatus", Flag: "new-auto-registration-status", Type: "types.AutoRegistrationStatus", Required: false},
	{Name: "NewStatus", Flag: "new-status", Type: "types.CACertificateStatus", Required: false},
	{Name: "RegistrationConfig", Flag: "registration-config", Type: "*types.RegistrationConfig", Required: false},
	{Name: "RemoveAutoRegistration", Flag: "remove-auto-registration", Type: "bool", Required: false},
}

var fields_update_certificate = []leanruntime.Field{
	{Name: "CertificateId", Flag: "certificate-id", Type: "*string", Required: true},
	{Name: "NewStatus", Flag: "new-status", Type: "types.CertificateStatus", Required: true},
}

var fields_update_certificate_provider = []leanruntime.Field{
	{Name: "AccountDefaultForOperations", Flag: "account-default-for-operations", Type: "[]types.CertificateProviderOperation", Required: false},
	{Name: "CertificateProviderName", Flag: "certificate-provider-name", Type: "*string", Required: true},
	{Name: "LambdaFunctionArn", Flag: "lambda-function-arn", Type: "*string", Required: false},
}

var fields_update_command = []leanruntime.Field{
	{Name: "CommandId", Flag: "command-id", Type: "*string", Required: true},
	{Name: "Deprecated", Flag: "deprecated", Type: "*bool", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
}

var fields_update_custom_metric = []leanruntime.Field{
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "MetricName", Flag: "metric-name", Type: "*string", Required: true},
}

var fields_update_dimension = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "StringValues", Flag: "string-values", Type: "[]string", Required: true},
}

var fields_update_domain_configuration = []leanruntime.Field{
	{Name: "ApplicationProtocol", Flag: "application-protocol", Type: "types.ApplicationProtocol", Required: false},
	{Name: "AuthenticationType", Flag: "authentication-type", Type: "types.AuthenticationType", Required: false},
	{Name: "AuthorizerConfig", Flag: "authorizer-config", Type: "*types.AuthorizerConfig", Required: false},
	{Name: "ClientCertificateConfig", Flag: "client-certificate-config", Type: "*types.ClientCertificateConfig", Required: false},
	{Name: "DomainConfigurationName", Flag: "domain-configuration-name", Type: "*string", Required: true},
	{Name: "DomainConfigurationStatus", Flag: "domain-configuration-status", Type: "types.DomainConfigurationStatus", Required: false},
	{Name: "RemoveAuthorizerConfig", Flag: "remove-authorizer-config", Type: "bool", Required: false},
	{Name: "ServerCertificateConfig", Flag: "server-certificate-config", Type: "*types.ServerCertificateConfig", Required: false},
	{Name: "TlsConfig", Flag: "tls-config", Type: "*types.TlsConfig", Required: false},
}

var fields_update_dynamic_thing_group = []leanruntime.Field{
	{Name: "ExpectedVersion", Flag: "expected-version", Type: "*int64", Required: false},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: false},
	{Name: "QueryString", Flag: "query-string", Type: "*string", Required: false},
	{Name: "QueryVersion", Flag: "query-version", Type: "*string", Required: false},
	{Name: "ThingGroupName", Flag: "thing-group-name", Type: "*string", Required: true},
	{Name: "ThingGroupProperties", Flag: "thing-group-properties", Type: "*types.ThingGroupProperties", Required: true},
}

var fields_update_encryption_configuration = []leanruntime.Field{
	{Name: "EncryptionType", Flag: "encryption-type", Type: "types.EncryptionType", Required: true},
	{Name: "KmsAccessRoleArn", Flag: "kms-access-role-arn", Type: "*string", Required: false},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
}

var fields_update_event_configurations = []leanruntime.Field{
	{Name: "EventConfigurations", Flag: "event-configurations", Type: "map[string]types.Configuration", Required: false},
}

var fields_update_fleet_metric = []leanruntime.Field{
	{Name: "AggregationField", Flag: "aggregation-field", Type: "*string", Required: false},
	{Name: "AggregationType", Flag: "aggregation-type", Type: "*types.AggregationType", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExpectedVersion", Flag: "expected-version", Type: "*int64", Required: false},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: true},
	{Name: "MetricName", Flag: "metric-name", Type: "*string", Required: true},
	{Name: "Period", Flag: "period", Type: "*int32", Required: false},
	{Name: "QueryString", Flag: "query-string", Type: "*string", Required: false},
	{Name: "QueryVersion", Flag: "query-version", Type: "*string", Required: false},
	{Name: "Unit", Flag: "unit", Type: "types.FleetMetricUnit", Required: false},
}

var fields_update_indexing_configuration = []leanruntime.Field{
	{Name: "ThingGroupIndexingConfiguration", Flag: "thing-group-indexing-configuration", Type: "*types.ThingGroupIndexingConfiguration", Required: false},
	{Name: "ThingIndexingConfiguration", Flag: "thing-indexing-configuration", Type: "*types.ThingIndexingConfiguration", Required: false},
}

var fields_update_job = []leanruntime.Field{
	{Name: "AbortConfig", Flag: "abort-config", Type: "*types.AbortConfig", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "JobExecutionsRetryConfig", Flag: "job-executions-retry-config", Type: "*types.JobExecutionsRetryConfig", Required: false},
	{Name: "JobExecutionsRolloutConfig", Flag: "job-executions-rollout-config", Type: "*types.JobExecutionsRolloutConfig", Required: false},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "NamespaceId", Flag: "namespace-id", Type: "*string", Required: false},
	{Name: "PresignedUrlConfig", Flag: "presigned-url-config", Type: "*types.PresignedUrlConfig", Required: false},
	{Name: "TimeoutConfig", Flag: "timeout-config", Type: "*types.TimeoutConfig", Required: false},
}

var fields_update_mitigation_action = []leanruntime.Field{
	{Name: "ActionName", Flag: "action-name", Type: "*string", Required: true},
	{Name: "ActionParams", Flag: "action-params", Type: "*types.MitigationActionParams", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_update_package = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DefaultVersionName", Flag: "default-version-name", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "PackageName", Flag: "package-name", Type: "*string", Required: true},
	{Name: "UnsetDefaultVersion", Flag: "unset-default-version", Type: "*bool", Required: false},
}

var fields_update_package_configuration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "VersionUpdateByJobsConfig", Flag: "version-update-by-jobs-config", Type: "*types.VersionUpdateByJobsConfig", Required: false},
}

var fields_update_package_version = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.PackageVersionAction", Required: false},
	{Name: "Artifact", Flag: "artifact", Type: "*types.PackageVersionArtifact", Required: false},
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "PackageName", Flag: "package-name", Type: "*string", Required: true},
	{Name: "Recipe", Flag: "recipe", Type: "*string", Required: false},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: true},
}

var fields_update_provisioning_template = []leanruntime.Field{
	{Name: "DefaultVersionId", Flag: "default-version-id", Type: "*int32", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "PreProvisioningHook", Flag: "pre-provisioning-hook", Type: "*types.ProvisioningHook", Required: false},
	{Name: "ProvisioningRoleArn", Flag: "provisioning-role-arn", Type: "*string", Required: false},
	{Name: "RemovePreProvisioningHook", Flag: "remove-pre-provisioning-hook", Type: "*bool", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_update_role_alias = []leanruntime.Field{
	{Name: "CredentialDurationSeconds", Flag: "credential-duration-seconds", Type: "*int32", Required: false},
	{Name: "RoleAlias", Flag: "role-alias", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_update_scheduled_audit = []leanruntime.Field{
	{Name: "DayOfMonth", Flag: "day-of-month", Type: "*string", Required: false},
	{Name: "DayOfWeek", Flag: "day-of-week", Type: "types.DayOfWeek", Required: false},
	{Name: "Frequency", Flag: "frequency", Type: "types.AuditFrequency", Required: false},
	{Name: "ScheduledAuditName", Flag: "scheduled-audit-name", Type: "*string", Required: true},
	{Name: "TargetCheckNames", Flag: "target-check-names", Type: "[]string", Required: false},
}

var fields_update_security_profile = []leanruntime.Field{
	{Name: "AdditionalMetricsToRetain", Flag: "additional-metrics-to-retain", Type: "[]string", Required: false},
	{Name: "AdditionalMetricsToRetainV2", Flag: "additional-metrics-to-retain-v2", Type: "[]types.MetricToRetain", Required: false},
	{Name: "AlertTargets", Flag: "alert-targets", Type: "map[string]types.AlertTarget", Required: false},
	{Name: "Behaviors", Flag: "behaviors", Type: "[]types.Behavior", Required: false},
	{Name: "DeleteAdditionalMetricsToRetain", Flag: "delete-additional-metrics-to-retain", Type: "bool", Required: false},
	{Name: "DeleteAlertTargets", Flag: "delete-alert-targets", Type: "bool", Required: false},
	{Name: "DeleteBehaviors", Flag: "delete-behaviors", Type: "bool", Required: false},
	{Name: "DeleteMetricsExportConfig", Flag: "delete-metrics-export-config", Type: "bool", Required: false},
	{Name: "ExpectedVersion", Flag: "expected-version", Type: "*int64", Required: false},
	{Name: "MetricsExportConfig", Flag: "metrics-export-config", Type: "*types.MetricsExportConfig", Required: false},
	{Name: "SecurityProfileDescription", Flag: "security-profile-description", Type: "*string", Required: false},
	{Name: "SecurityProfileName", Flag: "security-profile-name", Type: "*string", Required: true},
}

var fields_update_stream = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Files", Flag: "files", Type: "[]types.StreamFile", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: true},
}

var fields_update_thing = []leanruntime.Field{
	{Name: "AttributePayload", Flag: "attribute-payload", Type: "*types.AttributePayload", Required: false},
	{Name: "ExpectedVersion", Flag: "expected-version", Type: "*int64", Required: false},
	{Name: "RemoveThingType", Flag: "remove-thing-type", Type: "bool", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
	{Name: "ThingTypeName", Flag: "thing-type-name", Type: "*string", Required: false},
}

var fields_update_thing_group = []leanruntime.Field{
	{Name: "ExpectedVersion", Flag: "expected-version", Type: "*int64", Required: false},
	{Name: "ThingGroupName", Flag: "thing-group-name", Type: "*string", Required: true},
	{Name: "ThingGroupProperties", Flag: "thing-group-properties", Type: "*types.ThingGroupProperties", Required: true},
}

var fields_update_thing_groups_for_thing = []leanruntime.Field{
	{Name: "OverrideDynamicGroups", Flag: "override-dynamic-groups", Type: "bool", Required: false},
	{Name: "ThingGroupsToAdd", Flag: "thing-groups-to-add", Type: "[]string", Required: false},
	{Name: "ThingGroupsToRemove", Flag: "thing-groups-to-remove", Type: "[]string", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: false},
}

var fields_update_thing_type = []leanruntime.Field{
	{Name: "ThingTypeName", Flag: "thing-type-name", Type: "*string", Required: true},
	{Name: "ThingTypeProperties", Flag: "thing-type-properties", Type: "*types.ThingTypeProperties", Required: false},
}

var fields_update_topic_rule_destination = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.TopicRuleDestinationStatus", Required: true},
}

var fields_validate_security_profile_behaviors = []leanruntime.Field{
	{Name: "Behaviors", Flag: "behaviors", Type: "[]types.Behavior", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-certificate-transfer": {
			Name:   "accept-certificate-transfer",
			Fields: fields_accept_certificate_transfer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptCertificateTransferInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_certificate_transfer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptCertificateTransfer(ctx, input)
			},
		},
		"add-thing-to-billing-group": {
			Name:   "add-thing-to-billing-group",
			Fields: fields_add_thing_to_billing_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddThingToBillingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_thing_to_billing_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddThingToBillingGroup(ctx, input)
			},
		},
		"add-thing-to-thing-group": {
			Name:   "add-thing-to-thing-group",
			Fields: fields_add_thing_to_thing_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddThingToThingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_thing_to_thing_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddThingToThingGroup(ctx, input)
			},
		},
		"associate-sbom-with-package-version": {
			Name:   "associate-sbom-with-package-version",
			Fields: fields_associate_sbom_with_package_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateSbomWithPackageVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_sbom_with_package_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateSbomWithPackageVersion(ctx, input)
			},
		},
		"associate-targets-with-job": {
			Name:   "associate-targets-with-job",
			Fields: fields_associate_targets_with_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateTargetsWithJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_targets_with_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateTargetsWithJob(ctx, input)
			},
		},
		"attach-policy": {
			Name:   "attach-policy",
			Fields: fields_attach_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachPolicy(ctx, input)
			},
		},
		"attach-principal-policy": {
			Name:   "attach-principal-policy",
			Fields: fields_attach_principal_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachPrincipalPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_principal_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachPrincipalPolicy(ctx, input)
			},
		},
		"attach-security-profile": {
			Name:   "attach-security-profile",
			Fields: fields_attach_security_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachSecurityProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_security_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachSecurityProfile(ctx, input)
			},
		},
		"attach-thing-principal": {
			Name:   "attach-thing-principal",
			Fields: fields_attach_thing_principal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachThingPrincipalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_thing_principal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachThingPrincipal(ctx, input)
			},
		},
		"cancel-audit-mitigation-actions-task": {
			Name:   "cancel-audit-mitigation-actions-task",
			Fields: fields_cancel_audit_mitigation_actions_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelAuditMitigationActionsTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_audit_mitigation_actions_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelAuditMitigationActionsTask(ctx, input)
			},
		},
		"cancel-audit-task": {
			Name:   "cancel-audit-task",
			Fields: fields_cancel_audit_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelAuditTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_audit_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelAuditTask(ctx, input)
			},
		},
		"cancel-certificate-transfer": {
			Name:   "cancel-certificate-transfer",
			Fields: fields_cancel_certificate_transfer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelCertificateTransferInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_certificate_transfer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelCertificateTransfer(ctx, input)
			},
		},
		"cancel-detect-mitigation-actions-task": {
			Name:   "cancel-detect-mitigation-actions-task",
			Fields: fields_cancel_detect_mitigation_actions_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelDetectMitigationActionsTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_detect_mitigation_actions_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelDetectMitigationActionsTask(ctx, input)
			},
		},
		"cancel-job": {
			Name:   "cancel-job",
			Fields: fields_cancel_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelJob(ctx, input)
			},
		},
		"cancel-job-execution": {
			Name:   "cancel-job-execution",
			Fields: fields_cancel_job_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelJobExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_job_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelJobExecution(ctx, input)
			},
		},
		"clear-default-authorizer": {
			Name:   "clear-default-authorizer",
			Fields: fields_clear_default_authorizer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ClearDefaultAuthorizerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_clear_default_authorizer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ClearDefaultAuthorizer(ctx, input)
			},
		},
		"confirm-topic-rule-destination": {
			Name:   "confirm-topic-rule-destination",
			Fields: fields_confirm_topic_rule_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ConfirmTopicRuleDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_confirm_topic_rule_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ConfirmTopicRuleDestination(ctx, input)
			},
		},
		"create-audit-suppression": {
			Name:   "create-audit-suppression",
			Fields: fields_create_audit_suppression,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAuditSuppressionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_audit_suppression, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAuditSuppression(ctx, input)
			},
		},
		"create-authorizer": {
			Name:   "create-authorizer",
			Fields: fields_create_authorizer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAuthorizerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_authorizer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAuthorizer(ctx, input)
			},
		},
		"create-billing-group": {
			Name:   "create-billing-group",
			Fields: fields_create_billing_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBillingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_billing_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBillingGroup(ctx, input)
			},
		},
		"create-certificate-from-csr": {
			Name:   "create-certificate-from-csr",
			Fields: fields_create_certificate_from_csr,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCertificateFromCsrInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_certificate_from_csr, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCertificateFromCsr(ctx, input)
			},
		},
		"create-certificate-provider": {
			Name:   "create-certificate-provider",
			Fields: fields_create_certificate_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCertificateProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_certificate_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCertificateProvider(ctx, input)
			},
		},
		"create-command": {
			Name:   "create-command",
			Fields: fields_create_command,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCommandInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_command, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCommand(ctx, input)
			},
		},
		"create-custom-metric": {
			Name:   "create-custom-metric",
			Fields: fields_create_custom_metric,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCustomMetricInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_custom_metric, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCustomMetric(ctx, input)
			},
		},
		"create-dimension": {
			Name:   "create-dimension",
			Fields: fields_create_dimension,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDimensionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_dimension, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDimension(ctx, input)
			},
		},
		"create-domain-configuration": {
			Name:   "create-domain-configuration",
			Fields: fields_create_domain_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDomainConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_domain_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDomainConfiguration(ctx, input)
			},
		},
		"create-dynamic-thing-group": {
			Name:   "create-dynamic-thing-group",
			Fields: fields_create_dynamic_thing_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDynamicThingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_dynamic_thing_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDynamicThingGroup(ctx, input)
			},
		},
		"create-fleet-metric": {
			Name:   "create-fleet-metric",
			Fields: fields_create_fleet_metric,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFleetMetricInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_fleet_metric, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFleetMetric(ctx, input)
			},
		},
		"create-job": {
			Name:   "create-job",
			Fields: fields_create_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateJob(ctx, input)
			},
		},
		"create-job-template": {
			Name:   "create-job-template",
			Fields: fields_create_job_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateJobTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_job_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateJobTemplate(ctx, input)
			},
		},
		"create-keys-and-certificate": {
			Name:   "create-keys-and-certificate",
			Fields: fields_create_keys_and_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateKeysAndCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_keys_and_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateKeysAndCertificate(ctx, input)
			},
		},
		"create-mitigation-action": {
			Name:   "create-mitigation-action",
			Fields: fields_create_mitigation_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMitigationActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_mitigation_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMitigationAction(ctx, input)
			},
		},
		"create-ota-update": {
			Name:   "create-ota-update",
			Fields: fields_create_ota_update,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOTAUpdateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ota_update, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOTAUpdate(ctx, input)
			},
		},
		"create-package": {
			Name:   "create-package",
			Fields: fields_create_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePackage(ctx, input)
			},
		},
		"create-package-version": {
			Name:   "create-package-version",
			Fields: fields_create_package_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePackageVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_package_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePackageVersion(ctx, input)
			},
		},
		"create-policy": {
			Name:   "create-policy",
			Fields: fields_create_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePolicy(ctx, input)
			},
		},
		"create-policy-version": {
			Name:   "create-policy-version",
			Fields: fields_create_policy_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePolicyVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_policy_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePolicyVersion(ctx, input)
			},
		},
		"create-provisioning-claim": {
			Name:   "create-provisioning-claim",
			Fields: fields_create_provisioning_claim,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProvisioningClaimInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_provisioning_claim, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProvisioningClaim(ctx, input)
			},
		},
		"create-provisioning-template": {
			Name:   "create-provisioning-template",
			Fields: fields_create_provisioning_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProvisioningTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_provisioning_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProvisioningTemplate(ctx, input)
			},
		},
		"create-provisioning-template-version": {
			Name:   "create-provisioning-template-version",
			Fields: fields_create_provisioning_template_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProvisioningTemplateVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_provisioning_template_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProvisioningTemplateVersion(ctx, input)
			},
		},
		"create-role-alias": {
			Name:   "create-role-alias",
			Fields: fields_create_role_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRoleAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_role_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRoleAlias(ctx, input)
			},
		},
		"create-scheduled-audit": {
			Name:   "create-scheduled-audit",
			Fields: fields_create_scheduled_audit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateScheduledAuditInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_scheduled_audit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateScheduledAudit(ctx, input)
			},
		},
		"create-security-profile": {
			Name:   "create-security-profile",
			Fields: fields_create_security_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSecurityProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_security_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSecurityProfile(ctx, input)
			},
		},
		"create-stream": {
			Name:   "create-stream",
			Fields: fields_create_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStream(ctx, input)
			},
		},
		"create-thing": {
			Name:   "create-thing",
			Fields: fields_create_thing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateThingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_thing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateThing(ctx, input)
			},
		},
		"create-thing-group": {
			Name:   "create-thing-group",
			Fields: fields_create_thing_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateThingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_thing_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateThingGroup(ctx, input)
			},
		},
		"create-thing-type": {
			Name:   "create-thing-type",
			Fields: fields_create_thing_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateThingTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_thing_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateThingType(ctx, input)
			},
		},
		"create-topic-rule": {
			Name:   "create-topic-rule",
			Fields: fields_create_topic_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTopicRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_topic_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTopicRule(ctx, input)
			},
		},
		"create-topic-rule-destination": {
			Name:   "create-topic-rule-destination",
			Fields: fields_create_topic_rule_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTopicRuleDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_topic_rule_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTopicRuleDestination(ctx, input)
			},
		},
		"delete-account-audit-configuration": {
			Name:   "delete-account-audit-configuration",
			Fields: fields_delete_account_audit_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccountAuditConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_account_audit_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccountAuditConfiguration(ctx, input)
			},
		},
		"delete-audit-suppression": {
			Name:   "delete-audit-suppression",
			Fields: fields_delete_audit_suppression,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAuditSuppressionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_audit_suppression, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAuditSuppression(ctx, input)
			},
		},
		"delete-authorizer": {
			Name:   "delete-authorizer",
			Fields: fields_delete_authorizer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAuthorizerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_authorizer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAuthorizer(ctx, input)
			},
		},
		"delete-billing-group": {
			Name:   "delete-billing-group",
			Fields: fields_delete_billing_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBillingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_billing_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBillingGroup(ctx, input)
			},
		},
		"delete-ca-certificate": {
			Name:   "delete-ca-certificate",
			Fields: fields_delete_ca_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCACertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ca_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCACertificate(ctx, input)
			},
		},
		"delete-certificate": {
			Name:   "delete-certificate",
			Fields: fields_delete_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCertificate(ctx, input)
			},
		},
		"delete-certificate-provider": {
			Name:   "delete-certificate-provider",
			Fields: fields_delete_certificate_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCertificateProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_certificate_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCertificateProvider(ctx, input)
			},
		},
		"delete-command": {
			Name:   "delete-command",
			Fields: fields_delete_command,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCommandInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_command, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCommand(ctx, input)
			},
		},
		"delete-command-execution": {
			Name:   "delete-command-execution",
			Fields: fields_delete_command_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCommandExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_command_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCommandExecution(ctx, input)
			},
		},
		"delete-custom-metric": {
			Name:   "delete-custom-metric",
			Fields: fields_delete_custom_metric,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCustomMetricInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_custom_metric, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCustomMetric(ctx, input)
			},
		},
		"delete-dimension": {
			Name:   "delete-dimension",
			Fields: fields_delete_dimension,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDimensionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_dimension, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDimension(ctx, input)
			},
		},
		"delete-domain-configuration": {
			Name:   "delete-domain-configuration",
			Fields: fields_delete_domain_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDomainConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_domain_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDomainConfiguration(ctx, input)
			},
		},
		"delete-dynamic-thing-group": {
			Name:   "delete-dynamic-thing-group",
			Fields: fields_delete_dynamic_thing_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDynamicThingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_dynamic_thing_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDynamicThingGroup(ctx, input)
			},
		},
		"delete-fleet-metric": {
			Name:   "delete-fleet-metric",
			Fields: fields_delete_fleet_metric,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFleetMetricInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_fleet_metric, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFleetMetric(ctx, input)
			},
		},
		"delete-job": {
			Name:   "delete-job",
			Fields: fields_delete_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteJob(ctx, input)
			},
		},
		"delete-job-execution": {
			Name:   "delete-job-execution",
			Fields: fields_delete_job_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteJobExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_job_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteJobExecution(ctx, input)
			},
		},
		"delete-job-template": {
			Name:   "delete-job-template",
			Fields: fields_delete_job_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteJobTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_job_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteJobTemplate(ctx, input)
			},
		},
		"delete-mitigation-action": {
			Name:   "delete-mitigation-action",
			Fields: fields_delete_mitigation_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMitigationActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_mitigation_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMitigationAction(ctx, input)
			},
		},
		"delete-ota-update": {
			Name:   "delete-ota-update",
			Fields: fields_delete_ota_update,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOTAUpdateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ota_update, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOTAUpdate(ctx, input)
			},
		},
		"delete-package": {
			Name:   "delete-package",
			Fields: fields_delete_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePackage(ctx, input)
			},
		},
		"delete-package-version": {
			Name:   "delete-package-version",
			Fields: fields_delete_package_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePackageVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_package_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePackageVersion(ctx, input)
			},
		},
		"delete-policy": {
			Name:   "delete-policy",
			Fields: fields_delete_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePolicy(ctx, input)
			},
		},
		"delete-policy-version": {
			Name:   "delete-policy-version",
			Fields: fields_delete_policy_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePolicyVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_policy_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePolicyVersion(ctx, input)
			},
		},
		"delete-provisioning-template": {
			Name:   "delete-provisioning-template",
			Fields: fields_delete_provisioning_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProvisioningTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_provisioning_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProvisioningTemplate(ctx, input)
			},
		},
		"delete-provisioning-template-version": {
			Name:   "delete-provisioning-template-version",
			Fields: fields_delete_provisioning_template_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProvisioningTemplateVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_provisioning_template_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProvisioningTemplateVersion(ctx, input)
			},
		},
		"delete-registration-code": {
			Name:   "delete-registration-code",
			Fields: fields_delete_registration_code,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRegistrationCodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_registration_code, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRegistrationCode(ctx, input)
			},
		},
		"delete-role-alias": {
			Name:   "delete-role-alias",
			Fields: fields_delete_role_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRoleAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_role_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRoleAlias(ctx, input)
			},
		},
		"delete-scheduled-audit": {
			Name:   "delete-scheduled-audit",
			Fields: fields_delete_scheduled_audit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteScheduledAuditInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_scheduled_audit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteScheduledAudit(ctx, input)
			},
		},
		"delete-security-profile": {
			Name:   "delete-security-profile",
			Fields: fields_delete_security_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSecurityProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_security_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSecurityProfile(ctx, input)
			},
		},
		"delete-stream": {
			Name:   "delete-stream",
			Fields: fields_delete_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStream(ctx, input)
			},
		},
		"delete-thing": {
			Name:   "delete-thing",
			Fields: fields_delete_thing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteThingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_thing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteThing(ctx, input)
			},
		},
		"delete-thing-group": {
			Name:   "delete-thing-group",
			Fields: fields_delete_thing_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteThingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_thing_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteThingGroup(ctx, input)
			},
		},
		"delete-thing-type": {
			Name:   "delete-thing-type",
			Fields: fields_delete_thing_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteThingTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_thing_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteThingType(ctx, input)
			},
		},
		"delete-topic-rule": {
			Name:   "delete-topic-rule",
			Fields: fields_delete_topic_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTopicRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_topic_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTopicRule(ctx, input)
			},
		},
		"delete-topic-rule-destination": {
			Name:   "delete-topic-rule-destination",
			Fields: fields_delete_topic_rule_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTopicRuleDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_topic_rule_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTopicRuleDestination(ctx, input)
			},
		},
		"delete-v2-logging-level": {
			Name:   "delete-v2-logging-level",
			Fields: fields_delete_v2_logging_level,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteV2LoggingLevelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_v2_logging_level, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteV2LoggingLevel(ctx, input)
			},
		},
		"deprecate-thing-type": {
			Name:   "deprecate-thing-type",
			Fields: fields_deprecate_thing_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeprecateThingTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deprecate_thing_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeprecateThingType(ctx, input)
			},
		},
		"describe-account-audit-configuration": {
			Name:   "describe-account-audit-configuration",
			Fields: fields_describe_account_audit_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountAuditConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_account_audit_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccountAuditConfiguration(ctx, input)
			},
		},
		"describe-audit-finding": {
			Name:   "describe-audit-finding",
			Fields: fields_describe_audit_finding,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAuditFindingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_audit_finding, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAuditFinding(ctx, input)
			},
		},
		"describe-audit-mitigation-actions-task": {
			Name:   "describe-audit-mitigation-actions-task",
			Fields: fields_describe_audit_mitigation_actions_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAuditMitigationActionsTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_audit_mitigation_actions_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAuditMitigationActionsTask(ctx, input)
			},
		},
		"describe-audit-suppression": {
			Name:   "describe-audit-suppression",
			Fields: fields_describe_audit_suppression,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAuditSuppressionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_audit_suppression, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAuditSuppression(ctx, input)
			},
		},
		"describe-audit-task": {
			Name:   "describe-audit-task",
			Fields: fields_describe_audit_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAuditTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_audit_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAuditTask(ctx, input)
			},
		},
		"describe-authorizer": {
			Name:   "describe-authorizer",
			Fields: fields_describe_authorizer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAuthorizerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_authorizer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAuthorizer(ctx, input)
			},
		},
		"describe-billing-group": {
			Name:   "describe-billing-group",
			Fields: fields_describe_billing_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBillingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_billing_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBillingGroup(ctx, input)
			},
		},
		"describe-ca-certificate": {
			Name:   "describe-ca-certificate",
			Fields: fields_describe_ca_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCACertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_ca_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCACertificate(ctx, input)
			},
		},
		"describe-certificate": {
			Name:   "describe-certificate",
			Fields: fields_describe_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCertificate(ctx, input)
			},
		},
		"describe-certificate-provider": {
			Name:   "describe-certificate-provider",
			Fields: fields_describe_certificate_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCertificateProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_certificate_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCertificateProvider(ctx, input)
			},
		},
		"describe-custom-metric": {
			Name:   "describe-custom-metric",
			Fields: fields_describe_custom_metric,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCustomMetricInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_custom_metric, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCustomMetric(ctx, input)
			},
		},
		"describe-default-authorizer": {
			Name:   "describe-default-authorizer",
			Fields: fields_describe_default_authorizer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDefaultAuthorizerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_default_authorizer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDefaultAuthorizer(ctx, input)
			},
		},
		"describe-detect-mitigation-actions-task": {
			Name:   "describe-detect-mitigation-actions-task",
			Fields: fields_describe_detect_mitigation_actions_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDetectMitigationActionsTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_detect_mitigation_actions_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDetectMitigationActionsTask(ctx, input)
			},
		},
		"describe-dimension": {
			Name:   "describe-dimension",
			Fields: fields_describe_dimension,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDimensionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_dimension, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDimension(ctx, input)
			},
		},
		"describe-domain-configuration": {
			Name:   "describe-domain-configuration",
			Fields: fields_describe_domain_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDomainConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_domain_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDomainConfiguration(ctx, input)
			},
		},
		"describe-encryption-configuration": {
			Name:   "describe-encryption-configuration",
			Fields: fields_describe_encryption_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEncryptionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_encryption_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEncryptionConfiguration(ctx, input)
			},
		},
		"describe-endpoint": {
			Name:   "describe-endpoint",
			Fields: fields_describe_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEndpoint(ctx, input)
			},
		},
		"describe-event-configurations": {
			Name:   "describe-event-configurations",
			Fields: fields_describe_event_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEventConfigurationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_event_configurations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEventConfigurations(ctx, input)
			},
		},
		"describe-fleet-metric": {
			Name:   "describe-fleet-metric",
			Fields: fields_describe_fleet_metric,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFleetMetricInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_fleet_metric, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFleetMetric(ctx, input)
			},
		},
		"describe-index": {
			Name:   "describe-index",
			Fields: fields_describe_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeIndex(ctx, input)
			},
		},
		"describe-job": {
			Name:   "describe-job",
			Fields: fields_describe_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeJob(ctx, input)
			},
		},
		"describe-job-execution": {
			Name:   "describe-job-execution",
			Fields: fields_describe_job_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeJobExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_job_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeJobExecution(ctx, input)
			},
		},
		"describe-job-template": {
			Name:   "describe-job-template",
			Fields: fields_describe_job_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeJobTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_job_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeJobTemplate(ctx, input)
			},
		},
		"describe-managed-job-template": {
			Name:   "describe-managed-job-template",
			Fields: fields_describe_managed_job_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeManagedJobTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_managed_job_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeManagedJobTemplate(ctx, input)
			},
		},
		"describe-mitigation-action": {
			Name:   "describe-mitigation-action",
			Fields: fields_describe_mitigation_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMitigationActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_mitigation_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeMitigationAction(ctx, input)
			},
		},
		"describe-provisioning-template": {
			Name:   "describe-provisioning-template",
			Fields: fields_describe_provisioning_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProvisioningTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_provisioning_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProvisioningTemplate(ctx, input)
			},
		},
		"describe-provisioning-template-version": {
			Name:   "describe-provisioning-template-version",
			Fields: fields_describe_provisioning_template_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProvisioningTemplateVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_provisioning_template_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProvisioningTemplateVersion(ctx, input)
			},
		},
		"describe-role-alias": {
			Name:   "describe-role-alias",
			Fields: fields_describe_role_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRoleAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_role_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRoleAlias(ctx, input)
			},
		},
		"describe-scheduled-audit": {
			Name:   "describe-scheduled-audit",
			Fields: fields_describe_scheduled_audit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeScheduledAuditInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_scheduled_audit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeScheduledAudit(ctx, input)
			},
		},
		"describe-security-profile": {
			Name:   "describe-security-profile",
			Fields: fields_describe_security_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSecurityProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_security_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSecurityProfile(ctx, input)
			},
		},
		"describe-stream": {
			Name:   "describe-stream",
			Fields: fields_describe_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStream(ctx, input)
			},
		},
		"describe-thing": {
			Name:   "describe-thing",
			Fields: fields_describe_thing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeThingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_thing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeThing(ctx, input)
			},
		},
		"describe-thing-group": {
			Name:   "describe-thing-group",
			Fields: fields_describe_thing_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeThingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_thing_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeThingGroup(ctx, input)
			},
		},
		"describe-thing-registration-task": {
			Name:   "describe-thing-registration-task",
			Fields: fields_describe_thing_registration_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeThingRegistrationTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_thing_registration_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeThingRegistrationTask(ctx, input)
			},
		},
		"describe-thing-type": {
			Name:   "describe-thing-type",
			Fields: fields_describe_thing_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeThingTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_thing_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeThingType(ctx, input)
			},
		},
		"detach-policy": {
			Name:   "detach-policy",
			Fields: fields_detach_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachPolicy(ctx, input)
			},
		},
		"detach-principal-policy": {
			Name:   "detach-principal-policy",
			Fields: fields_detach_principal_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachPrincipalPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_principal_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachPrincipalPolicy(ctx, input)
			},
		},
		"detach-security-profile": {
			Name:   "detach-security-profile",
			Fields: fields_detach_security_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachSecurityProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_security_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachSecurityProfile(ctx, input)
			},
		},
		"detach-thing-principal": {
			Name:   "detach-thing-principal",
			Fields: fields_detach_thing_principal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachThingPrincipalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_thing_principal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachThingPrincipal(ctx, input)
			},
		},
		"disable-topic-rule": {
			Name:   "disable-topic-rule",
			Fields: fields_disable_topic_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableTopicRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_topic_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableTopicRule(ctx, input)
			},
		},
		"disassociate-sbom-from-package-version": {
			Name:   "disassociate-sbom-from-package-version",
			Fields: fields_disassociate_sbom_from_package_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateSbomFromPackageVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_sbom_from_package_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateSbomFromPackageVersion(ctx, input)
			},
		},
		"enable-topic-rule": {
			Name:   "enable-topic-rule",
			Fields: fields_enable_topic_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableTopicRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_topic_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableTopicRule(ctx, input)
			},
		},
		"get-behavior-model-training-summaries": {
			Name:   "get-behavior-model-training-summaries",
			Fields: fields_get_behavior_model_training_summaries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBehaviorModelTrainingSummariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_behavior_model_training_summaries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetBehaviorModelTrainingSummaries(ctx, input)
				}
				var results []*svc.GetBehaviorModelTrainingSummariesOutput
				p := svc.NewGetBehaviorModelTrainingSummariesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-buckets-aggregation": {
			Name:   "get-buckets-aggregation",
			Fields: fields_get_buckets_aggregation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketsAggregationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_buckets_aggregation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketsAggregation(ctx, input)
			},
		},
		"get-cardinality": {
			Name:   "get-cardinality",
			Fields: fields_get_cardinality,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCardinalityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cardinality, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCardinality(ctx, input)
			},
		},
		"get-command": {
			Name:   "get-command",
			Fields: fields_get_command,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCommandInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_command, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCommand(ctx, input)
			},
		},
		"get-command-execution": {
			Name:   "get-command-execution",
			Fields: fields_get_command_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCommandExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_command_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCommandExecution(ctx, input)
			},
		},
		"get-effective-policies": {
			Name:   "get-effective-policies",
			Fields: fields_get_effective_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEffectivePoliciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_effective_policies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEffectivePolicies(ctx, input)
			},
		},
		"get-indexing-configuration": {
			Name:   "get-indexing-configuration",
			Fields: fields_get_indexing_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIndexingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_indexing_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIndexingConfiguration(ctx, input)
			},
		},
		"get-job-document": {
			Name:   "get-job-document",
			Fields: fields_get_job_document,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJobDocumentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_job_document, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJobDocument(ctx, input)
			},
		},
		"get-logging-options": {
			Name:   "get-logging-options",
			Fields: fields_get_logging_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLoggingOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_logging_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLoggingOptions(ctx, input)
			},
		},
		"get-ota-update": {
			Name:   "get-ota-update",
			Fields: fields_get_ota_update,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOTAUpdateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ota_update, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOTAUpdate(ctx, input)
			},
		},
		"get-package": {
			Name:   "get-package",
			Fields: fields_get_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPackage(ctx, input)
			},
		},
		"get-package-configuration": {
			Name:   "get-package-configuration",
			Fields: fields_get_package_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPackageConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_package_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPackageConfiguration(ctx, input)
			},
		},
		"get-package-version": {
			Name:   "get-package-version",
			Fields: fields_get_package_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPackageVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_package_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPackageVersion(ctx, input)
			},
		},
		"get-percentiles": {
			Name:   "get-percentiles",
			Fields: fields_get_percentiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPercentilesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_percentiles, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPercentiles(ctx, input)
			},
		},
		"get-policy": {
			Name:   "get-policy",
			Fields: fields_get_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPolicy(ctx, input)
			},
		},
		"get-policy-version": {
			Name:   "get-policy-version",
			Fields: fields_get_policy_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPolicyVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_policy_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPolicyVersion(ctx, input)
			},
		},
		"get-registration-code": {
			Name:   "get-registration-code",
			Fields: fields_get_registration_code,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRegistrationCodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_registration_code, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRegistrationCode(ctx, input)
			},
		},
		"get-statistics": {
			Name:   "get-statistics",
			Fields: fields_get_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStatisticsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_statistics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStatistics(ctx, input)
			},
		},
		"get-thing-connectivity-data": {
			Name:   "get-thing-connectivity-data",
			Fields: fields_get_thing_connectivity_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetThingConnectivityDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_thing_connectivity_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetThingConnectivityData(ctx, input)
			},
		},
		"get-topic-rule": {
			Name:   "get-topic-rule",
			Fields: fields_get_topic_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTopicRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_topic_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTopicRule(ctx, input)
			},
		},
		"get-topic-rule-destination": {
			Name:   "get-topic-rule-destination",
			Fields: fields_get_topic_rule_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTopicRuleDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_topic_rule_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTopicRuleDestination(ctx, input)
			},
		},
		"get-v2-logging-options": {
			Name:   "get-v2-logging-options",
			Fields: fields_get_v2_logging_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetV2LoggingOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_v2_logging_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetV2LoggingOptions(ctx, input)
			},
		},
		"list-active-violations": {
			Name:   "list-active-violations",
			Fields: fields_list_active_violations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListActiveViolationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_active_violations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListActiveViolations(ctx, input)
				}
				var results []*svc.ListActiveViolationsOutput
				p := svc.NewListActiveViolationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-attached-policies": {
			Name:   "list-attached-policies",
			Fields: fields_list_attached_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAttachedPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_attached_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAttachedPolicies(ctx, input)
				}
				var results []*svc.ListAttachedPoliciesOutput
				p := svc.NewListAttachedPoliciesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-audit-findings": {
			Name:   "list-audit-findings",
			Fields: fields_list_audit_findings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAuditFindingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_audit_findings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAuditFindings(ctx, input)
				}
				var results []*svc.ListAuditFindingsOutput
				p := svc.NewListAuditFindingsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-audit-mitigation-actions-executions": {
			Name:   "list-audit-mitigation-actions-executions",
			Fields: fields_list_audit_mitigation_actions_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAuditMitigationActionsExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_audit_mitigation_actions_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAuditMitigationActionsExecutions(ctx, input)
				}
				var results []*svc.ListAuditMitigationActionsExecutionsOutput
				p := svc.NewListAuditMitigationActionsExecutionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-audit-mitigation-actions-tasks": {
			Name:   "list-audit-mitigation-actions-tasks",
			Fields: fields_list_audit_mitigation_actions_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAuditMitigationActionsTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_audit_mitigation_actions_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAuditMitigationActionsTasks(ctx, input)
				}
				var results []*svc.ListAuditMitigationActionsTasksOutput
				p := svc.NewListAuditMitigationActionsTasksPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-audit-suppressions": {
			Name:   "list-audit-suppressions",
			Fields: fields_list_audit_suppressions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAuditSuppressionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_audit_suppressions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAuditSuppressions(ctx, input)
				}
				var results []*svc.ListAuditSuppressionsOutput
				p := svc.NewListAuditSuppressionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-audit-tasks": {
			Name:   "list-audit-tasks",
			Fields: fields_list_audit_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAuditTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_audit_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAuditTasks(ctx, input)
				}
				var results []*svc.ListAuditTasksOutput
				p := svc.NewListAuditTasksPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-authorizers": {
			Name:   "list-authorizers",
			Fields: fields_list_authorizers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAuthorizersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_authorizers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAuthorizers(ctx, input)
				}
				var results []*svc.ListAuthorizersOutput
				p := svc.NewListAuthorizersPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-billing-groups": {
			Name:   "list-billing-groups",
			Fields: fields_list_billing_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBillingGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_billing_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBillingGroups(ctx, input)
				}
				var results []*svc.ListBillingGroupsOutput
				p := svc.NewListBillingGroupsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-ca-certificates": {
			Name:   "list-ca-certificates",
			Fields: fields_list_ca_certificates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCACertificatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ca_certificates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCACertificates(ctx, input)
				}
				var results []*svc.ListCACertificatesOutput
				p := svc.NewListCACertificatesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-certificate-providers": {
			Name:   "list-certificate-providers",
			Fields: fields_list_certificate_providers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCertificateProvidersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_certificate_providers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListCertificateProviders(ctx, input)
			},
		},
		"list-certificates": {
			Name:   "list-certificates",
			Fields: fields_list_certificates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCertificatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_certificates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCertificates(ctx, input)
				}
				var results []*svc.ListCertificatesOutput
				p := svc.NewListCertificatesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-certificates-by-ca": {
			Name:   "list-certificates-by-ca",
			Fields: fields_list_certificates_by_ca,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCertificatesByCAInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_certificates_by_ca, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCertificatesByCA(ctx, input)
				}
				var results []*svc.ListCertificatesByCAOutput
				p := svc.NewListCertificatesByCAPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-command-executions": {
			Name:   "list-command-executions",
			Fields: fields_list_command_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCommandExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_command_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCommandExecutions(ctx, input)
				}
				var results []*svc.ListCommandExecutionsOutput
				p := svc.NewListCommandExecutionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-commands": {
			Name:   "list-commands",
			Fields: fields_list_commands,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCommandsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_commands, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCommands(ctx, input)
				}
				var results []*svc.ListCommandsOutput
				p := svc.NewListCommandsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-custom-metrics": {
			Name:   "list-custom-metrics",
			Fields: fields_list_custom_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCustomMetricsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_custom_metrics, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCustomMetrics(ctx, input)
				}
				var results []*svc.ListCustomMetricsOutput
				p := svc.NewListCustomMetricsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-detect-mitigation-actions-executions": {
			Name:   "list-detect-mitigation-actions-executions",
			Fields: fields_list_detect_mitigation_actions_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDetectMitigationActionsExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_detect_mitigation_actions_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDetectMitigationActionsExecutions(ctx, input)
				}
				var results []*svc.ListDetectMitigationActionsExecutionsOutput
				p := svc.NewListDetectMitigationActionsExecutionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-detect-mitigation-actions-tasks": {
			Name:   "list-detect-mitigation-actions-tasks",
			Fields: fields_list_detect_mitigation_actions_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDetectMitigationActionsTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_detect_mitigation_actions_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDetectMitigationActionsTasks(ctx, input)
				}
				var results []*svc.ListDetectMitigationActionsTasksOutput
				p := svc.NewListDetectMitigationActionsTasksPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-dimensions": {
			Name:   "list-dimensions",
			Fields: fields_list_dimensions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDimensionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_dimensions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDimensions(ctx, input)
				}
				var results []*svc.ListDimensionsOutput
				p := svc.NewListDimensionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-domain-configurations": {
			Name:   "list-domain-configurations",
			Fields: fields_list_domain_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDomainConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_domain_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDomainConfigurations(ctx, input)
				}
				var results []*svc.ListDomainConfigurationsOutput
				p := svc.NewListDomainConfigurationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-fleet-metrics": {
			Name:   "list-fleet-metrics",
			Fields: fields_list_fleet_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFleetMetricsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_fleet_metrics, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFleetMetrics(ctx, input)
				}
				var results []*svc.ListFleetMetricsOutput
				p := svc.NewListFleetMetricsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-indices": {
			Name:   "list-indices",
			Fields: fields_list_indices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIndicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_indices, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIndices(ctx, input)
				}
				var results []*svc.ListIndicesOutput
				p := svc.NewListIndicesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-job-executions-for-job": {
			Name:   "list-job-executions-for-job",
			Fields: fields_list_job_executions_for_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListJobExecutionsForJobInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_job_executions_for_job, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListJobExecutionsForJob(ctx, input)
				}
				var results []*svc.ListJobExecutionsForJobOutput
				p := svc.NewListJobExecutionsForJobPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-job-executions-for-thing": {
			Name:   "list-job-executions-for-thing",
			Fields: fields_list_job_executions_for_thing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListJobExecutionsForThingInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_job_executions_for_thing, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListJobExecutionsForThing(ctx, input)
				}
				var results []*svc.ListJobExecutionsForThingOutput
				p := svc.NewListJobExecutionsForThingPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-job-templates": {
			Name:   "list-job-templates",
			Fields: fields_list_job_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListJobTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_job_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListJobTemplates(ctx, input)
				}
				var results []*svc.ListJobTemplatesOutput
				p := svc.NewListJobTemplatesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-jobs": {
			Name:   "list-jobs",
			Fields: fields_list_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListJobs(ctx, input)
				}
				var results []*svc.ListJobsOutput
				p := svc.NewListJobsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-managed-job-templates": {
			Name:   "list-managed-job-templates",
			Fields: fields_list_managed_job_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListManagedJobTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_managed_job_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListManagedJobTemplates(ctx, input)
				}
				var results []*svc.ListManagedJobTemplatesOutput
				p := svc.NewListManagedJobTemplatesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-metric-values": {
			Name:   "list-metric-values",
			Fields: fields_list_metric_values,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMetricValuesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_metric_values, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMetricValues(ctx, input)
				}
				var results []*svc.ListMetricValuesOutput
				p := svc.NewListMetricValuesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-mitigation-actions": {
			Name:   "list-mitigation-actions",
			Fields: fields_list_mitigation_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMitigationActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_mitigation_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMitigationActions(ctx, input)
				}
				var results []*svc.ListMitigationActionsOutput
				p := svc.NewListMitigationActionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-ota-updates": {
			Name:   "list-ota-updates",
			Fields: fields_list_ota_updates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOTAUpdatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ota_updates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOTAUpdates(ctx, input)
				}
				var results []*svc.ListOTAUpdatesOutput
				p := svc.NewListOTAUpdatesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-outgoing-certificates": {
			Name:   "list-outgoing-certificates",
			Fields: fields_list_outgoing_certificates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOutgoingCertificatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_outgoing_certificates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOutgoingCertificates(ctx, input)
				}
				var results []*svc.ListOutgoingCertificatesOutput
				p := svc.NewListOutgoingCertificatesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-package-versions": {
			Name:   "list-package-versions",
			Fields: fields_list_package_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPackageVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_package_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPackageVersions(ctx, input)
				}
				var results []*svc.ListPackageVersionsOutput
				p := svc.NewListPackageVersionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-packages": {
			Name:   "list-packages",
			Fields: fields_list_packages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPackagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_packages, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPackages(ctx, input)
				}
				var results []*svc.ListPackagesOutput
				p := svc.NewListPackagesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-policies": {
			Name:   "list-policies",
			Fields: fields_list_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPolicies(ctx, input)
				}
				var results []*svc.ListPoliciesOutput
				p := svc.NewListPoliciesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-policy-principals": {
			Name:   "list-policy-principals",
			Fields: fields_list_policy_principals,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPolicyPrincipalsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_policy_principals, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPolicyPrincipals(ctx, input)
				}
				var results []*svc.ListPolicyPrincipalsOutput
				p := svc.NewListPolicyPrincipalsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-policy-versions": {
			Name:   "list-policy-versions",
			Fields: fields_list_policy_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPolicyVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_policy_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListPolicyVersions(ctx, input)
			},
		},
		"list-principal-policies": {
			Name:   "list-principal-policies",
			Fields: fields_list_principal_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPrincipalPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_principal_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPrincipalPolicies(ctx, input)
				}
				var results []*svc.ListPrincipalPoliciesOutput
				p := svc.NewListPrincipalPoliciesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-principal-things": {
			Name:   "list-principal-things",
			Fields: fields_list_principal_things,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPrincipalThingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_principal_things, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPrincipalThings(ctx, input)
				}
				var results []*svc.ListPrincipalThingsOutput
				p := svc.NewListPrincipalThingsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-principal-things-v2": {
			Name:   "list-principal-things-v2",
			Fields: fields_list_principal_things_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPrincipalThingsV2Input{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_principal_things_v2, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPrincipalThingsV2(ctx, input)
				}
				var results []*svc.ListPrincipalThingsV2Output
				p := svc.NewListPrincipalThingsV2Paginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-provisioning-template-versions": {
			Name:   "list-provisioning-template-versions",
			Fields: fields_list_provisioning_template_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProvisioningTemplateVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_provisioning_template_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProvisioningTemplateVersions(ctx, input)
				}
				var results []*svc.ListProvisioningTemplateVersionsOutput
				p := svc.NewListProvisioningTemplateVersionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-provisioning-templates": {
			Name:   "list-provisioning-templates",
			Fields: fields_list_provisioning_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProvisioningTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_provisioning_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProvisioningTemplates(ctx, input)
				}
				var results []*svc.ListProvisioningTemplatesOutput
				p := svc.NewListProvisioningTemplatesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-related-resources-for-audit-finding": {
			Name:   "list-related-resources-for-audit-finding",
			Fields: fields_list_related_resources_for_audit_finding,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRelatedResourcesForAuditFindingInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_related_resources_for_audit_finding, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRelatedResourcesForAuditFinding(ctx, input)
				}
				var results []*svc.ListRelatedResourcesForAuditFindingOutput
				p := svc.NewListRelatedResourcesForAuditFindingPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-role-aliases": {
			Name:   "list-role-aliases",
			Fields: fields_list_role_aliases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRoleAliasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_role_aliases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRoleAliases(ctx, input)
				}
				var results []*svc.ListRoleAliasesOutput
				p := svc.NewListRoleAliasesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-sbom-validation-results": {
			Name:   "list-sbom-validation-results",
			Fields: fields_list_sbom_validation_results,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSbomValidationResultsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sbom_validation_results, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSbomValidationResults(ctx, input)
				}
				var results []*svc.ListSbomValidationResultsOutput
				p := svc.NewListSbomValidationResultsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-scheduled-audits": {
			Name:   "list-scheduled-audits",
			Fields: fields_list_scheduled_audits,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListScheduledAuditsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_scheduled_audits, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListScheduledAudits(ctx, input)
				}
				var results []*svc.ListScheduledAuditsOutput
				p := svc.NewListScheduledAuditsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-security-profiles": {
			Name:   "list-security-profiles",
			Fields: fields_list_security_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSecurityProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_security_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSecurityProfiles(ctx, input)
				}
				var results []*svc.ListSecurityProfilesOutput
				p := svc.NewListSecurityProfilesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-security-profiles-for-target": {
			Name:   "list-security-profiles-for-target",
			Fields: fields_list_security_profiles_for_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSecurityProfilesForTargetInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_security_profiles_for_target, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSecurityProfilesForTarget(ctx, input)
				}
				var results []*svc.ListSecurityProfilesForTargetOutput
				p := svc.NewListSecurityProfilesForTargetPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-streams": {
			Name:   "list-streams",
			Fields: fields_list_streams,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStreamsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_streams, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStreams(ctx, input)
				}
				var results []*svc.ListStreamsOutput
				p := svc.NewListStreamsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-tags-for-resource": {
			Name:   "list-tags-for-resource",
			Fields: fields_list_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForResourceInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTagsForResource(ctx, input)
				}
				var results []*svc.ListTagsForResourceOutput
				p := svc.NewListTagsForResourcePaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-targets-for-policy": {
			Name:   "list-targets-for-policy",
			Fields: fields_list_targets_for_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTargetsForPolicyInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_targets_for_policy, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTargetsForPolicy(ctx, input)
				}
				var results []*svc.ListTargetsForPolicyOutput
				p := svc.NewListTargetsForPolicyPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-targets-for-security-profile": {
			Name:   "list-targets-for-security-profile",
			Fields: fields_list_targets_for_security_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTargetsForSecurityProfileInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_targets_for_security_profile, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTargetsForSecurityProfile(ctx, input)
				}
				var results []*svc.ListTargetsForSecurityProfileOutput
				p := svc.NewListTargetsForSecurityProfilePaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-thing-groups": {
			Name:   "list-thing-groups",
			Fields: fields_list_thing_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListThingGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_thing_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListThingGroups(ctx, input)
				}
				var results []*svc.ListThingGroupsOutput
				p := svc.NewListThingGroupsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-thing-groups-for-thing": {
			Name:   "list-thing-groups-for-thing",
			Fields: fields_list_thing_groups_for_thing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListThingGroupsForThingInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_thing_groups_for_thing, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListThingGroupsForThing(ctx, input)
				}
				var results []*svc.ListThingGroupsForThingOutput
				p := svc.NewListThingGroupsForThingPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-thing-principals": {
			Name:   "list-thing-principals",
			Fields: fields_list_thing_principals,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListThingPrincipalsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_thing_principals, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListThingPrincipals(ctx, input)
				}
				var results []*svc.ListThingPrincipalsOutput
				p := svc.NewListThingPrincipalsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-thing-principals-v2": {
			Name:   "list-thing-principals-v2",
			Fields: fields_list_thing_principals_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListThingPrincipalsV2Input{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_thing_principals_v2, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListThingPrincipalsV2(ctx, input)
				}
				var results []*svc.ListThingPrincipalsV2Output
				p := svc.NewListThingPrincipalsV2Paginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-thing-registration-task-reports": {
			Name:   "list-thing-registration-task-reports",
			Fields: fields_list_thing_registration_task_reports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListThingRegistrationTaskReportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_thing_registration_task_reports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListThingRegistrationTaskReports(ctx, input)
				}
				var results []*svc.ListThingRegistrationTaskReportsOutput
				p := svc.NewListThingRegistrationTaskReportsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-thing-registration-tasks": {
			Name:   "list-thing-registration-tasks",
			Fields: fields_list_thing_registration_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListThingRegistrationTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_thing_registration_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListThingRegistrationTasks(ctx, input)
				}
				var results []*svc.ListThingRegistrationTasksOutput
				p := svc.NewListThingRegistrationTasksPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-thing-types": {
			Name:   "list-thing-types",
			Fields: fields_list_thing_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListThingTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_thing_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListThingTypes(ctx, input)
				}
				var results []*svc.ListThingTypesOutput
				p := svc.NewListThingTypesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-things": {
			Name:   "list-things",
			Fields: fields_list_things,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListThingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_things, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListThings(ctx, input)
				}
				var results []*svc.ListThingsOutput
				p := svc.NewListThingsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-things-in-billing-group": {
			Name:   "list-things-in-billing-group",
			Fields: fields_list_things_in_billing_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListThingsInBillingGroupInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_things_in_billing_group, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListThingsInBillingGroup(ctx, input)
				}
				var results []*svc.ListThingsInBillingGroupOutput
				p := svc.NewListThingsInBillingGroupPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-things-in-thing-group": {
			Name:   "list-things-in-thing-group",
			Fields: fields_list_things_in_thing_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListThingsInThingGroupInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_things_in_thing_group, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListThingsInThingGroup(ctx, input)
				}
				var results []*svc.ListThingsInThingGroupOutput
				p := svc.NewListThingsInThingGroupPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-topic-rule-destinations": {
			Name:   "list-topic-rule-destinations",
			Fields: fields_list_topic_rule_destinations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTopicRuleDestinationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_topic_rule_destinations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTopicRuleDestinations(ctx, input)
				}
				var results []*svc.ListTopicRuleDestinationsOutput
				p := svc.NewListTopicRuleDestinationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-topic-rules": {
			Name:   "list-topic-rules",
			Fields: fields_list_topic_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTopicRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_topic_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTopicRules(ctx, input)
				}
				var results []*svc.ListTopicRulesOutput
				p := svc.NewListTopicRulesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-v2-logging-levels": {
			Name:   "list-v2-logging-levels",
			Fields: fields_list_v2_logging_levels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListV2LoggingLevelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_v2_logging_levels, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListV2LoggingLevels(ctx, input)
				}
				var results []*svc.ListV2LoggingLevelsOutput
				p := svc.NewListV2LoggingLevelsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-violation-events": {
			Name:   "list-violation-events",
			Fields: fields_list_violation_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListViolationEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_violation_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListViolationEvents(ctx, input)
				}
				var results []*svc.ListViolationEventsOutput
				p := svc.NewListViolationEventsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"put-verification-state-on-violation": {
			Name:   "put-verification-state-on-violation",
			Fields: fields_put_verification_state_on_violation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutVerificationStateOnViolationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_verification_state_on_violation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutVerificationStateOnViolation(ctx, input)
			},
		},
		"register-ca-certificate": {
			Name:   "register-ca-certificate",
			Fields: fields_register_ca_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterCACertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_ca_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterCACertificate(ctx, input)
			},
		},
		"register-certificate": {
			Name:   "register-certificate",
			Fields: fields_register_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterCertificate(ctx, input)
			},
		},
		"register-certificate-without-ca": {
			Name:   "register-certificate-without-ca",
			Fields: fields_register_certificate_without_ca,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterCertificateWithoutCAInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_certificate_without_ca, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterCertificateWithoutCA(ctx, input)
			},
		},
		"register-thing": {
			Name:   "register-thing",
			Fields: fields_register_thing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterThingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_thing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterThing(ctx, input)
			},
		},
		"reject-certificate-transfer": {
			Name:   "reject-certificate-transfer",
			Fields: fields_reject_certificate_transfer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectCertificateTransferInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_certificate_transfer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectCertificateTransfer(ctx, input)
			},
		},
		"remove-thing-from-billing-group": {
			Name:   "remove-thing-from-billing-group",
			Fields: fields_remove_thing_from_billing_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveThingFromBillingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_thing_from_billing_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveThingFromBillingGroup(ctx, input)
			},
		},
		"remove-thing-from-thing-group": {
			Name:   "remove-thing-from-thing-group",
			Fields: fields_remove_thing_from_thing_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveThingFromThingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_thing_from_thing_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveThingFromThingGroup(ctx, input)
			},
		},
		"replace-topic-rule": {
			Name:   "replace-topic-rule",
			Fields: fields_replace_topic_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReplaceTopicRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_replace_topic_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReplaceTopicRule(ctx, input)
			},
		},
		"search-index": {
			Name:   "search-index",
			Fields: fields_search_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_search_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SearchIndex(ctx, input)
			},
		},
		"set-default-authorizer": {
			Name:   "set-default-authorizer",
			Fields: fields_set_default_authorizer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetDefaultAuthorizerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_default_authorizer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetDefaultAuthorizer(ctx, input)
			},
		},
		"set-default-policy-version": {
			Name:   "set-default-policy-version",
			Fields: fields_set_default_policy_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetDefaultPolicyVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_default_policy_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetDefaultPolicyVersion(ctx, input)
			},
		},
		"set-logging-options": {
			Name:   "set-logging-options",
			Fields: fields_set_logging_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetLoggingOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_logging_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetLoggingOptions(ctx, input)
			},
		},
		"set-v2-logging-level": {
			Name:   "set-v2-logging-level",
			Fields: fields_set_v2_logging_level,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetV2LoggingLevelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_v2_logging_level, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetV2LoggingLevel(ctx, input)
			},
		},
		"set-v2-logging-options": {
			Name:   "set-v2-logging-options",
			Fields: fields_set_v2_logging_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetV2LoggingOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_v2_logging_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetV2LoggingOptions(ctx, input)
			},
		},
		"start-audit-mitigation-actions-task": {
			Name:   "start-audit-mitigation-actions-task",
			Fields: fields_start_audit_mitigation_actions_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAuditMitigationActionsTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_audit_mitigation_actions_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAuditMitigationActionsTask(ctx, input)
			},
		},
		"start-detect-mitigation-actions-task": {
			Name:   "start-detect-mitigation-actions-task",
			Fields: fields_start_detect_mitigation_actions_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDetectMitigationActionsTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_detect_mitigation_actions_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDetectMitigationActionsTask(ctx, input)
			},
		},
		"start-on-demand-audit-task": {
			Name:   "start-on-demand-audit-task",
			Fields: fields_start_on_demand_audit_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartOnDemandAuditTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_on_demand_audit_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartOnDemandAuditTask(ctx, input)
			},
		},
		"start-thing-registration-task": {
			Name:   "start-thing-registration-task",
			Fields: fields_start_thing_registration_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartThingRegistrationTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_thing_registration_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartThingRegistrationTask(ctx, input)
			},
		},
		"stop-thing-registration-task": {
			Name:   "stop-thing-registration-task",
			Fields: fields_stop_thing_registration_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopThingRegistrationTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_thing_registration_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopThingRegistrationTask(ctx, input)
			},
		},
		"tag-resource": {
			Name:   "tag-resource",
			Fields: fields_tag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagResource(ctx, input)
			},
		},
		"test-authorization": {
			Name:   "test-authorization",
			Fields: fields_test_authorization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestAuthorizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_authorization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestAuthorization(ctx, input)
			},
		},
		"test-invoke-authorizer": {
			Name:   "test-invoke-authorizer",
			Fields: fields_test_invoke_authorizer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestInvokeAuthorizerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_invoke_authorizer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestInvokeAuthorizer(ctx, input)
			},
		},
		"transfer-certificate": {
			Name:   "transfer-certificate",
			Fields: fields_transfer_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TransferCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_transfer_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TransferCertificate(ctx, input)
			},
		},
		"untag-resource": {
			Name:   "untag-resource",
			Fields: fields_untag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagResource(ctx, input)
			},
		},
		"update-account-audit-configuration": {
			Name:   "update-account-audit-configuration",
			Fields: fields_update_account_audit_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccountAuditConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_account_audit_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccountAuditConfiguration(ctx, input)
			},
		},
		"update-audit-suppression": {
			Name:   "update-audit-suppression",
			Fields: fields_update_audit_suppression,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAuditSuppressionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_audit_suppression, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAuditSuppression(ctx, input)
			},
		},
		"update-authorizer": {
			Name:   "update-authorizer",
			Fields: fields_update_authorizer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAuthorizerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_authorizer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAuthorizer(ctx, input)
			},
		},
		"update-billing-group": {
			Name:   "update-billing-group",
			Fields: fields_update_billing_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBillingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_billing_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBillingGroup(ctx, input)
			},
		},
		"update-ca-certificate": {
			Name:   "update-ca-certificate",
			Fields: fields_update_ca_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCACertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_ca_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCACertificate(ctx, input)
			},
		},
		"update-certificate": {
			Name:   "update-certificate",
			Fields: fields_update_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCertificate(ctx, input)
			},
		},
		"update-certificate-provider": {
			Name:   "update-certificate-provider",
			Fields: fields_update_certificate_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCertificateProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_certificate_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCertificateProvider(ctx, input)
			},
		},
		"update-command": {
			Name:   "update-command",
			Fields: fields_update_command,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCommandInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_command, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCommand(ctx, input)
			},
		},
		"update-custom-metric": {
			Name:   "update-custom-metric",
			Fields: fields_update_custom_metric,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCustomMetricInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_custom_metric, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCustomMetric(ctx, input)
			},
		},
		"update-dimension": {
			Name:   "update-dimension",
			Fields: fields_update_dimension,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDimensionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_dimension, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDimension(ctx, input)
			},
		},
		"update-domain-configuration": {
			Name:   "update-domain-configuration",
			Fields: fields_update_domain_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDomainConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_domain_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDomainConfiguration(ctx, input)
			},
		},
		"update-dynamic-thing-group": {
			Name:   "update-dynamic-thing-group",
			Fields: fields_update_dynamic_thing_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDynamicThingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_dynamic_thing_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDynamicThingGroup(ctx, input)
			},
		},
		"update-encryption-configuration": {
			Name:   "update-encryption-configuration",
			Fields: fields_update_encryption_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEncryptionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_encryption_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEncryptionConfiguration(ctx, input)
			},
		},
		"update-event-configurations": {
			Name:   "update-event-configurations",
			Fields: fields_update_event_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEventConfigurationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_event_configurations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEventConfigurations(ctx, input)
			},
		},
		"update-fleet-metric": {
			Name:   "update-fleet-metric",
			Fields: fields_update_fleet_metric,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFleetMetricInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_fleet_metric, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFleetMetric(ctx, input)
			},
		},
		"update-indexing-configuration": {
			Name:   "update-indexing-configuration",
			Fields: fields_update_indexing_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIndexingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_indexing_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIndexingConfiguration(ctx, input)
			},
		},
		"update-job": {
			Name:   "update-job",
			Fields: fields_update_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateJob(ctx, input)
			},
		},
		"update-mitigation-action": {
			Name:   "update-mitigation-action",
			Fields: fields_update_mitigation_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMitigationActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_mitigation_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMitigationAction(ctx, input)
			},
		},
		"update-package": {
			Name:   "update-package",
			Fields: fields_update_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePackage(ctx, input)
			},
		},
		"update-package-configuration": {
			Name:   "update-package-configuration",
			Fields: fields_update_package_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePackageConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_package_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePackageConfiguration(ctx, input)
			},
		},
		"update-package-version": {
			Name:   "update-package-version",
			Fields: fields_update_package_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePackageVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_package_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePackageVersion(ctx, input)
			},
		},
		"update-provisioning-template": {
			Name:   "update-provisioning-template",
			Fields: fields_update_provisioning_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProvisioningTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_provisioning_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProvisioningTemplate(ctx, input)
			},
		},
		"update-role-alias": {
			Name:   "update-role-alias",
			Fields: fields_update_role_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRoleAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_role_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRoleAlias(ctx, input)
			},
		},
		"update-scheduled-audit": {
			Name:   "update-scheduled-audit",
			Fields: fields_update_scheduled_audit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateScheduledAuditInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_scheduled_audit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateScheduledAudit(ctx, input)
			},
		},
		"update-security-profile": {
			Name:   "update-security-profile",
			Fields: fields_update_security_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSecurityProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_security_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSecurityProfile(ctx, input)
			},
		},
		"update-stream": {
			Name:   "update-stream",
			Fields: fields_update_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStream(ctx, input)
			},
		},
		"update-thing": {
			Name:   "update-thing",
			Fields: fields_update_thing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateThingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_thing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateThing(ctx, input)
			},
		},
		"update-thing-group": {
			Name:   "update-thing-group",
			Fields: fields_update_thing_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateThingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_thing_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateThingGroup(ctx, input)
			},
		},
		"update-thing-groups-for-thing": {
			Name:   "update-thing-groups-for-thing",
			Fields: fields_update_thing_groups_for_thing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateThingGroupsForThingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_thing_groups_for_thing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateThingGroupsForThing(ctx, input)
			},
		},
		"update-thing-type": {
			Name:   "update-thing-type",
			Fields: fields_update_thing_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateThingTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_thing_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateThingType(ctx, input)
			},
		},
		"update-topic-rule-destination": {
			Name:   "update-topic-rule-destination",
			Fields: fields_update_topic_rule_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTopicRuleDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_topic_rule_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTopicRuleDestination(ctx, input)
			},
		},
		"validate-security-profile-behaviors": {
			Name:   "validate-security-profile-behaviors",
			Fields: fields_validate_security_profile_behaviors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ValidateSecurityProfileBehaviorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_validate_security_profile_behaviors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ValidateSecurityProfileBehaviors(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("iot", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
