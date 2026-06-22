package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/transfer"
)

var fields_create_access = []leanruntime.Field{
	{Name: "ExternalId", Flag: "external-id", Type: "*string", Required: true},
	{Name: "HomeDirectory", Flag: "home-directory", Type: "*string", Required: false},
	{Name: "HomeDirectoryMappings", Flag: "home-directory-mappings", Type: "[]types.HomeDirectoryMapEntry", Required: false},
	{Name: "HomeDirectoryType", Flag: "home-directory-type", Type: "types.HomeDirectoryType", Required: false},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: false},
	{Name: "PosixProfile", Flag: "posix-profile", Type: "*types.PosixProfile", Required: false},
	{Name: "Role", Flag: "role", Type: "*string", Required: true},
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
}

var fields_create_agreement = []leanruntime.Field{
	{Name: "AccessRole", Flag: "access-role", Type: "*string", Required: true},
	{Name: "BaseDirectory", Flag: "base-directory", Type: "*string", Required: false},
	{Name: "CustomDirectories", Flag: "custom-directories", Type: "*types.CustomDirectoriesType", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnforceMessageSigning", Flag: "enforce-message-signing", Type: "types.EnforceMessageSigningType", Required: false},
	{Name: "LocalProfileId", Flag: "local-profile-id", Type: "*string", Required: true},
	{Name: "PartnerProfileId", Flag: "partner-profile-id", Type: "*string", Required: true},
	{Name: "PreserveFilename", Flag: "preserve-filename", Type: "types.PreserveFilenameType", Required: false},
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.AgreementStatusType", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_connector = []leanruntime.Field{
	{Name: "AccessRole", Flag: "access-role", Type: "*string", Required: true},
	{Name: "As2Config", Flag: "as2-config", Type: "*types.As2ConnectorConfig", Required: false},
	{Name: "EgressConfig", Flag: "egress-config", Type: "types.ConnectorEgressConfig", Required: false},
	{Name: "LoggingRole", Flag: "logging-role", Type: "*string", Required: false},
	{Name: "SecurityPolicyName", Flag: "security-policy-name", Type: "*string", Required: false},
	{Name: "SftpConfig", Flag: "sftp-config", Type: "*types.SftpConnectorConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Url", Flag: "url", Type: "*string", Required: false},
}

var fields_create_profile = []leanruntime.Field{
	{Name: "As2Id", Flag: "as2-id", Type: "*string", Required: true},
	{Name: "CertificateIds", Flag: "certificate-ids", Type: "[]string", Required: false},
	{Name: "ProfileType", Flag: "profile-type", Type: "types.ProfileType", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_server = []leanruntime.Field{
	{Name: "Certificate", Flag: "certificate", Type: "*string", Required: false},
	{Name: "Domain", Flag: "domain", Type: "types.Domain", Required: false},
	{Name: "EndpointDetails", Flag: "endpoint-details", Type: "*types.EndpointDetails", Required: false},
	{Name: "EndpointType", Flag: "endpoint-type", Type: "types.EndpointType", Required: false},
	{Name: "HostKey", Flag: "host-key", Type: "*string", Required: false},
	{Name: "IdentityProviderDetails", Flag: "identity-provider-details", Type: "*types.IdentityProviderDetails", Required: false},
	{Name: "IdentityProviderType", Flag: "identity-provider-type", Type: "types.IdentityProviderType", Required: false},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: false},
	{Name: "LoggingRole", Flag: "logging-role", Type: "*string", Required: false},
	{Name: "PostAuthenticationLoginBanner", Flag: "post-authentication-login-banner", Type: "*string", Required: false},
	{Name: "PreAuthenticationLoginBanner", Flag: "pre-authentication-login-banner", Type: "*string", Required: false},
	{Name: "ProtocolDetails", Flag: "protocol-details", Type: "*types.ProtocolDetails", Required: false},
	{Name: "Protocols", Flag: "protocols", Type: "[]types.Protocol", Required: false},
	{Name: "S3StorageOptions", Flag: "s3-storage-options", Type: "*types.S3StorageOptions", Required: false},
	{Name: "SecurityPolicyName", Flag: "security-policy-name", Type: "*string", Required: false},
	{Name: "StructuredLogDestinations", Flag: "structured-log-destinations", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "WorkflowDetails", Flag: "workflow-details", Type: "*types.WorkflowDetails", Required: false},
}

var fields_create_user = []leanruntime.Field{
	{Name: "HomeDirectory", Flag: "home-directory", Type: "*string", Required: false},
	{Name: "HomeDirectoryMappings", Flag: "home-directory-mappings", Type: "[]types.HomeDirectoryMapEntry", Required: false},
	{Name: "HomeDirectoryType", Flag: "home-directory-type", Type: "types.HomeDirectoryType", Required: false},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: false},
	{Name: "PosixProfile", Flag: "posix-profile", Type: "*types.PosixProfile", Required: false},
	{Name: "Role", Flag: "role", Type: "*string", Required: true},
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
	{Name: "SshPublicKeyBody", Flag: "ssh-public-key-body", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_create_web_app = []leanruntime.Field{
	{Name: "AccessEndpoint", Flag: "access-endpoint", Type: "*string", Required: false},
	{Name: "EndpointDetails", Flag: "endpoint-details", Type: "types.WebAppEndpointDetails", Required: false},
	{Name: "IdentityProviderDetails", Flag: "identity-provider-details", Type: "types.WebAppIdentityProviderDetails", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "WebAppEndpointPolicy", Flag: "web-app-endpoint-policy", Type: "types.WebAppEndpointPolicy", Required: false},
	{Name: "WebAppUnits", Flag: "web-app-units", Type: "types.WebAppUnits", Required: false},
}

var fields_create_workflow = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "OnExceptionSteps", Flag: "on-exception-steps", Type: "[]types.WorkflowStep", Required: false},
	{Name: "Steps", Flag: "steps", Type: "[]types.WorkflowStep", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_access = []leanruntime.Field{
	{Name: "ExternalId", Flag: "external-id", Type: "*string", Required: true},
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
}

var fields_delete_agreement = []leanruntime.Field{
	{Name: "AgreementId", Flag: "agreement-id", Type: "*string", Required: true},
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
}

var fields_delete_certificate = []leanruntime.Field{
	{Name: "CertificateId", Flag: "certificate-id", Type: "*string", Required: true},
}

var fields_delete_connector = []leanruntime.Field{
	{Name: "ConnectorId", Flag: "connector-id", Type: "*string", Required: true},
}

var fields_delete_host_key = []leanruntime.Field{
	{Name: "HostKeyId", Flag: "host-key-id", Type: "*string", Required: true},
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
}

var fields_delete_profile = []leanruntime.Field{
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
}

var fields_delete_server = []leanruntime.Field{
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
}

var fields_delete_ssh_public_key = []leanruntime.Field{
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
	{Name: "SshPublicKeyId", Flag: "ssh-public-key-id", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_delete_user = []leanruntime.Field{
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_delete_web_app = []leanruntime.Field{
	{Name: "WebAppId", Flag: "web-app-id", Type: "*string", Required: true},
}

var fields_delete_web_app_customization = []leanruntime.Field{
	{Name: "WebAppId", Flag: "web-app-id", Type: "*string", Required: true},
}

var fields_delete_workflow = []leanruntime.Field{
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
}

var fields_describe_access = []leanruntime.Field{
	{Name: "ExternalId", Flag: "external-id", Type: "*string", Required: true},
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
}

var fields_describe_agreement = []leanruntime.Field{
	{Name: "AgreementId", Flag: "agreement-id", Type: "*string", Required: true},
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
}

var fields_describe_certificate = []leanruntime.Field{
	{Name: "CertificateId", Flag: "certificate-id", Type: "*string", Required: true},
}

var fields_describe_connector = []leanruntime.Field{
	{Name: "ConnectorId", Flag: "connector-id", Type: "*string", Required: true},
}

var fields_describe_execution = []leanruntime.Field{
	{Name: "ExecutionId", Flag: "execution-id", Type: "*string", Required: true},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
}

var fields_describe_host_key = []leanruntime.Field{
	{Name: "HostKeyId", Flag: "host-key-id", Type: "*string", Required: true},
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
}

var fields_describe_profile = []leanruntime.Field{
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
}

var fields_describe_security_policy = []leanruntime.Field{
	{Name: "SecurityPolicyName", Flag: "security-policy-name", Type: "*string", Required: true},
}

var fields_describe_server = []leanruntime.Field{
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
}

var fields_describe_user = []leanruntime.Field{
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_describe_web_app = []leanruntime.Field{
	{Name: "WebAppId", Flag: "web-app-id", Type: "*string", Required: true},
}

var fields_describe_web_app_customization = []leanruntime.Field{
	{Name: "WebAppId", Flag: "web-app-id", Type: "*string", Required: true},
}

var fields_describe_workflow = []leanruntime.Field{
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
}

var fields_import_certificate = []leanruntime.Field{
	{Name: "ActiveDate", Flag: "active-date", Type: "*time.Time", Required: false},
	{Name: "Certificate", Flag: "certificate", Type: "*string", Required: true},
	{Name: "CertificateChain", Flag: "certificate-chain", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InactiveDate", Flag: "inactive-date", Type: "*time.Time", Required: false},
	{Name: "PrivateKey", Flag: "private-key", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Usage", Flag: "usage", Type: "types.CertificateUsageType", Required: true},
}

var fields_import_host_key = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "HostKeyBody", Flag: "host-key-body", Type: "*string", Required: true},
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_import_ssh_public_key = []leanruntime.Field{
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
	{Name: "SshPublicKeyBody", Flag: "ssh-public-key-body", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_list_accesses = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
}

var fields_list_agreements = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
}

var fields_list_certificates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_connectors = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_executions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
}

var fields_list_file_transfer_results = []leanruntime.Field{
	{Name: "ConnectorId", Flag: "connector-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransferId", Flag: "transfer-id", Type: "*string", Required: true},
}

var fields_list_host_keys = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
}

var fields_list_profiles = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProfileType", Flag: "profile-type", Type: "types.ProfileType", Required: false},
}

var fields_list_security_policies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_servers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_users = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
}

var fields_list_web_apps = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_workflows = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_send_workflow_step_state = []leanruntime.Field{
	{Name: "ExecutionId", Flag: "execution-id", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.CustomStepStatus", Required: true},
	{Name: "Token", Flag: "token", Type: "*string", Required: true},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
}

var fields_start_directory_listing = []leanruntime.Field{
	{Name: "ConnectorId", Flag: "connector-id", Type: "*string", Required: true},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "OutputDirectoryPath", Flag: "output-directory-path", Type: "*string", Required: true},
	{Name: "RemoteDirectoryPath", Flag: "remote-directory-path", Type: "*string", Required: true},
}

var fields_start_file_transfer = []leanruntime.Field{
	{Name: "ConnectorId", Flag: "connector-id", Type: "*string", Required: true},
	{Name: "CustomHttpHeaders", Flag: "custom-http-headers", Type: "[]types.CustomHttpHeader", Required: false},
	{Name: "LocalDirectoryPath", Flag: "local-directory-path", Type: "*string", Required: false},
	{Name: "RemoteDirectoryPath", Flag: "remote-directory-path", Type: "*string", Required: false},
	{Name: "RetrieveFilePaths", Flag: "retrieve-file-paths", Type: "[]string", Required: false},
	{Name: "SendFilePaths", Flag: "send-file-paths", Type: "[]string", Required: false},
}

var fields_start_remote_delete = []leanruntime.Field{
	{Name: "ConnectorId", Flag: "connector-id", Type: "*string", Required: true},
	{Name: "DeletePath", Flag: "delete-path", Type: "*string", Required: true},
}

var fields_start_remote_move = []leanruntime.Field{
	{Name: "ConnectorId", Flag: "connector-id", Type: "*string", Required: true},
	{Name: "SourcePath", Flag: "source-path", Type: "*string", Required: true},
	{Name: "TargetPath", Flag: "target-path", Type: "*string", Required: true},
}

var fields_start_server = []leanruntime.Field{
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
}

var fields_stop_server = []leanruntime.Field{
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_test_connection = []leanruntime.Field{
	{Name: "ConnectorId", Flag: "connector-id", Type: "*string", Required: true},
}

var fields_test_identity_provider = []leanruntime.Field{
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
	{Name: "ServerProtocol", Flag: "server-protocol", Type: "types.Protocol", Required: false},
	{Name: "SourceIp", Flag: "source-ip", Type: "*string", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
	{Name: "UserPassword", Flag: "user-password", Type: "*string", Required: false},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_access = []leanruntime.Field{
	{Name: "ExternalId", Flag: "external-id", Type: "*string", Required: true},
	{Name: "HomeDirectory", Flag: "home-directory", Type: "*string", Required: false},
	{Name: "HomeDirectoryMappings", Flag: "home-directory-mappings", Type: "[]types.HomeDirectoryMapEntry", Required: false},
	{Name: "HomeDirectoryType", Flag: "home-directory-type", Type: "types.HomeDirectoryType", Required: false},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: false},
	{Name: "PosixProfile", Flag: "posix-profile", Type: "*types.PosixProfile", Required: false},
	{Name: "Role", Flag: "role", Type: "*string", Required: false},
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
}

var fields_update_agreement = []leanruntime.Field{
	{Name: "AccessRole", Flag: "access-role", Type: "*string", Required: false},
	{Name: "AgreementId", Flag: "agreement-id", Type: "*string", Required: true},
	{Name: "BaseDirectory", Flag: "base-directory", Type: "*string", Required: false},
	{Name: "CustomDirectories", Flag: "custom-directories", Type: "*types.CustomDirectoriesType", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnforceMessageSigning", Flag: "enforce-message-signing", Type: "types.EnforceMessageSigningType", Required: false},
	{Name: "LocalProfileId", Flag: "local-profile-id", Type: "*string", Required: false},
	{Name: "PartnerProfileId", Flag: "partner-profile-id", Type: "*string", Required: false},
	{Name: "PreserveFilename", Flag: "preserve-filename", Type: "types.PreserveFilenameType", Required: false},
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.AgreementStatusType", Required: false},
}

var fields_update_certificate = []leanruntime.Field{
	{Name: "ActiveDate", Flag: "active-date", Type: "*time.Time", Required: false},
	{Name: "CertificateId", Flag: "certificate-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InactiveDate", Flag: "inactive-date", Type: "*time.Time", Required: false},
}

var fields_update_connector = []leanruntime.Field{
	{Name: "AccessRole", Flag: "access-role", Type: "*string", Required: false},
	{Name: "As2Config", Flag: "as2-config", Type: "*types.As2ConnectorConfig", Required: false},
	{Name: "ConnectorId", Flag: "connector-id", Type: "*string", Required: true},
	{Name: "EgressConfig", Flag: "egress-config", Type: "types.UpdateConnectorEgressConfig", Required: false},
	{Name: "LoggingRole", Flag: "logging-role", Type: "*string", Required: false},
	{Name: "SecurityPolicyName", Flag: "security-policy-name", Type: "*string", Required: false},
	{Name: "SftpConfig", Flag: "sftp-config", Type: "*types.SftpConnectorConfig", Required: false},
	{Name: "Url", Flag: "url", Type: "*string", Required: false},
}

var fields_update_host_key = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "HostKeyId", Flag: "host-key-id", Type: "*string", Required: true},
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
}

var fields_update_profile = []leanruntime.Field{
	{Name: "CertificateIds", Flag: "certificate-ids", Type: "[]string", Required: false},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
}

var fields_update_server = []leanruntime.Field{
	{Name: "Certificate", Flag: "certificate", Type: "*string", Required: false},
	{Name: "EndpointDetails", Flag: "endpoint-details", Type: "*types.EndpointDetails", Required: false},
	{Name: "EndpointType", Flag: "endpoint-type", Type: "types.EndpointType", Required: false},
	{Name: "HostKey", Flag: "host-key", Type: "*string", Required: false},
	{Name: "IdentityProviderDetails", Flag: "identity-provider-details", Type: "*types.IdentityProviderDetails", Required: false},
	{Name: "IdentityProviderType", Flag: "identity-provider-type", Type: "types.IdentityProviderType", Required: false},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: false},
	{Name: "LoggingRole", Flag: "logging-role", Type: "*string", Required: false},
	{Name: "PostAuthenticationLoginBanner", Flag: "post-authentication-login-banner", Type: "*string", Required: false},
	{Name: "PreAuthenticationLoginBanner", Flag: "pre-authentication-login-banner", Type: "*string", Required: false},
	{Name: "ProtocolDetails", Flag: "protocol-details", Type: "*types.ProtocolDetails", Required: false},
	{Name: "Protocols", Flag: "protocols", Type: "[]types.Protocol", Required: false},
	{Name: "S3StorageOptions", Flag: "s3-storage-options", Type: "*types.S3StorageOptions", Required: false},
	{Name: "SecurityPolicyName", Flag: "security-policy-name", Type: "*string", Required: false},
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
	{Name: "StructuredLogDestinations", Flag: "structured-log-destinations", Type: "[]string", Required: false},
	{Name: "WorkflowDetails", Flag: "workflow-details", Type: "*types.WorkflowDetails", Required: false},
}

var fields_update_user = []leanruntime.Field{
	{Name: "HomeDirectory", Flag: "home-directory", Type: "*string", Required: false},
	{Name: "HomeDirectoryMappings", Flag: "home-directory-mappings", Type: "[]types.HomeDirectoryMapEntry", Required: false},
	{Name: "HomeDirectoryType", Flag: "home-directory-type", Type: "types.HomeDirectoryType", Required: false},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: false},
	{Name: "PosixProfile", Flag: "posix-profile", Type: "*types.PosixProfile", Required: false},
	{Name: "Role", Flag: "role", Type: "*string", Required: false},
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_update_web_app = []leanruntime.Field{
	{Name: "AccessEndpoint", Flag: "access-endpoint", Type: "*string", Required: false},
	{Name: "EndpointDetails", Flag: "endpoint-details", Type: "types.UpdateWebAppEndpointDetails", Required: false},
	{Name: "IdentityProviderDetails", Flag: "identity-provider-details", Type: "types.UpdateWebAppIdentityProviderDetails", Required: false},
	{Name: "WebAppId", Flag: "web-app-id", Type: "*string", Required: true},
	{Name: "WebAppUnits", Flag: "web-app-units", Type: "types.WebAppUnits", Required: false},
}

var fields_update_web_app_customization = []leanruntime.Field{
	{Name: "FaviconFile", Flag: "favicon-file", Type: "[]byte", Required: false},
	{Name: "LogoFile", Flag: "logo-file", Type: "[]byte", Required: false},
	{Name: "Title", Flag: "title", Type: "*string", Required: false},
	{Name: "WebAppId", Flag: "web-app-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-access": {
			Name:   "create-access",
			Fields: fields_create_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccess(ctx, input)
			},
		},
		"create-agreement": {
			Name:   "create-agreement",
			Fields: fields_create_agreement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAgreementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_agreement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAgreement(ctx, input)
			},
		},
		"create-connector": {
			Name:   "create-connector",
			Fields: fields_create_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConnector(ctx, input)
			},
		},
		"create-profile": {
			Name:   "create-profile",
			Fields: fields_create_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProfile(ctx, input)
			},
		},
		"create-server": {
			Name:   "create-server",
			Fields: fields_create_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateServer(ctx, input)
			},
		},
		"create-user": {
			Name:   "create-user",
			Fields: fields_create_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUser(ctx, input)
			},
		},
		"create-web-app": {
			Name:   "create-web-app",
			Fields: fields_create_web_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWebAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_web_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWebApp(ctx, input)
			},
		},
		"create-workflow": {
			Name:   "create-workflow",
			Fields: fields_create_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkflow(ctx, input)
			},
		},
		"delete-access": {
			Name:   "delete-access",
			Fields: fields_delete_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccess(ctx, input)
			},
		},
		"delete-agreement": {
			Name:   "delete-agreement",
			Fields: fields_delete_agreement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAgreementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_agreement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAgreement(ctx, input)
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
		"delete-connector": {
			Name:   "delete-connector",
			Fields: fields_delete_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConnector(ctx, input)
			},
		},
		"delete-host-key": {
			Name:   "delete-host-key",
			Fields: fields_delete_host_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteHostKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_host_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteHostKey(ctx, input)
			},
		},
		"delete-profile": {
			Name:   "delete-profile",
			Fields: fields_delete_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProfile(ctx, input)
			},
		},
		"delete-server": {
			Name:   "delete-server",
			Fields: fields_delete_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteServer(ctx, input)
			},
		},
		"delete-ssh-public-key": {
			Name:   "delete-ssh-public-key",
			Fields: fields_delete_ssh_public_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSshPublicKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ssh_public_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSshPublicKey(ctx, input)
			},
		},
		"delete-user": {
			Name:   "delete-user",
			Fields: fields_delete_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUser(ctx, input)
			},
		},
		"delete-web-app": {
			Name:   "delete-web-app",
			Fields: fields_delete_web_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWebAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_web_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWebApp(ctx, input)
			},
		},
		"delete-web-app-customization": {
			Name:   "delete-web-app-customization",
			Fields: fields_delete_web_app_customization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWebAppCustomizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_web_app_customization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWebAppCustomization(ctx, input)
			},
		},
		"delete-workflow": {
			Name:   "delete-workflow",
			Fields: fields_delete_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkflow(ctx, input)
			},
		},
		"describe-access": {
			Name:   "describe-access",
			Fields: fields_describe_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccess(ctx, input)
			},
		},
		"describe-agreement": {
			Name:   "describe-agreement",
			Fields: fields_describe_agreement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAgreementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_agreement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAgreement(ctx, input)
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
		"describe-connector": {
			Name:   "describe-connector",
			Fields: fields_describe_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConnector(ctx, input)
			},
		},
		"describe-execution": {
			Name:   "describe-execution",
			Fields: fields_describe_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeExecution(ctx, input)
			},
		},
		"describe-host-key": {
			Name:   "describe-host-key",
			Fields: fields_describe_host_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeHostKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_host_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeHostKey(ctx, input)
			},
		},
		"describe-profile": {
			Name:   "describe-profile",
			Fields: fields_describe_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProfile(ctx, input)
			},
		},
		"describe-security-policy": {
			Name:   "describe-security-policy",
			Fields: fields_describe_security_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSecurityPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_security_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSecurityPolicy(ctx, input)
			},
		},
		"describe-server": {
			Name:   "describe-server",
			Fields: fields_describe_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeServer(ctx, input)
			},
		},
		"describe-user": {
			Name:   "describe-user",
			Fields: fields_describe_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeUser(ctx, input)
			},
		},
		"describe-web-app": {
			Name:   "describe-web-app",
			Fields: fields_describe_web_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWebAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_web_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWebApp(ctx, input)
			},
		},
		"describe-web-app-customization": {
			Name:   "describe-web-app-customization",
			Fields: fields_describe_web_app_customization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWebAppCustomizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_web_app_customization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWebAppCustomization(ctx, input)
			},
		},
		"describe-workflow": {
			Name:   "describe-workflow",
			Fields: fields_describe_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWorkflow(ctx, input)
			},
		},
		"import-certificate": {
			Name:   "import-certificate",
			Fields: fields_import_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportCertificate(ctx, input)
			},
		},
		"import-host-key": {
			Name:   "import-host-key",
			Fields: fields_import_host_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportHostKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_host_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportHostKey(ctx, input)
			},
		},
		"import-ssh-public-key": {
			Name:   "import-ssh-public-key",
			Fields: fields_import_ssh_public_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportSshPublicKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_ssh_public_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportSshPublicKey(ctx, input)
			},
		},
		"list-accesses": {
			Name:   "list-accesses",
			Fields: fields_list_accesses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccessesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_accesses, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccesses(ctx, input)
				}
				var results []*svc.ListAccessesOutput
				p := svc.NewListAccessesPaginator(client, input)
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
		"list-agreements": {
			Name:   "list-agreements",
			Fields: fields_list_agreements,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAgreementsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_agreements, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAgreements(ctx, input)
				}
				var results []*svc.ListAgreementsOutput
				p := svc.NewListAgreementsPaginator(client, input)
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
		"list-connectors": {
			Name:   "list-connectors",
			Fields: fields_list_connectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConnectorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_connectors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConnectors(ctx, input)
				}
				var results []*svc.ListConnectorsOutput
				p := svc.NewListConnectorsPaginator(client, input)
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
		"list-executions": {
			Name:   "list-executions",
			Fields: fields_list_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExecutions(ctx, input)
				}
				var results []*svc.ListExecutionsOutput
				p := svc.NewListExecutionsPaginator(client, input)
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
		"list-file-transfer-results": {
			Name:   "list-file-transfer-results",
			Fields: fields_list_file_transfer_results,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFileTransferResultsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_file_transfer_results, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFileTransferResults(ctx, input)
				}
				var results []*svc.ListFileTransferResultsOutput
				p := svc.NewListFileTransferResultsPaginator(client, input)
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
		"list-host-keys": {
			Name:   "list-host-keys",
			Fields: fields_list_host_keys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHostKeysInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_host_keys, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListHostKeys(ctx, input)
			},
		},
		"list-profiles": {
			Name:   "list-profiles",
			Fields: fields_list_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProfiles(ctx, input)
				}
				var results []*svc.ListProfilesOutput
				p := svc.NewListProfilesPaginator(client, input)
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
		"list-security-policies": {
			Name:   "list-security-policies",
			Fields: fields_list_security_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSecurityPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_security_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSecurityPolicies(ctx, input)
				}
				var results []*svc.ListSecurityPoliciesOutput
				p := svc.NewListSecurityPoliciesPaginator(client, input)
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
		"list-servers": {
			Name:   "list-servers",
			Fields: fields_list_servers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_servers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServers(ctx, input)
				}
				var results []*svc.ListServersOutput
				p := svc.NewListServersPaginator(client, input)
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
		"list-users": {
			Name:   "list-users",
			Fields: fields_list_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUsersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_users, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUsers(ctx, input)
				}
				var results []*svc.ListUsersOutput
				p := svc.NewListUsersPaginator(client, input)
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
		"list-web-apps": {
			Name:   "list-web-apps",
			Fields: fields_list_web_apps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWebAppsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_web_apps, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWebApps(ctx, input)
				}
				var results []*svc.ListWebAppsOutput
				p := svc.NewListWebAppsPaginator(client, input)
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
		"list-workflows": {
			Name:   "list-workflows",
			Fields: fields_list_workflows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkflowsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workflows, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkflows(ctx, input)
				}
				var results []*svc.ListWorkflowsOutput
				p := svc.NewListWorkflowsPaginator(client, input)
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
		"send-workflow-step-state": {
			Name:   "send-workflow-step-state",
			Fields: fields_send_workflow_step_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendWorkflowStepStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_workflow_step_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendWorkflowStepState(ctx, input)
			},
		},
		"start-directory-listing": {
			Name:   "start-directory-listing",
			Fields: fields_start_directory_listing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDirectoryListingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_directory_listing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDirectoryListing(ctx, input)
			},
		},
		"start-file-transfer": {
			Name:   "start-file-transfer",
			Fields: fields_start_file_transfer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartFileTransferInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_file_transfer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartFileTransfer(ctx, input)
			},
		},
		"start-remote-delete": {
			Name:   "start-remote-delete",
			Fields: fields_start_remote_delete,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartRemoteDeleteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_remote_delete, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartRemoteDelete(ctx, input)
			},
		},
		"start-remote-move": {
			Name:   "start-remote-move",
			Fields: fields_start_remote_move,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartRemoteMoveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_remote_move, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartRemoteMove(ctx, input)
			},
		},
		"start-server": {
			Name:   "start-server",
			Fields: fields_start_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartServer(ctx, input)
			},
		},
		"stop-server": {
			Name:   "stop-server",
			Fields: fields_stop_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopServer(ctx, input)
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
		"test-connection": {
			Name:   "test-connection",
			Fields: fields_test_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestConnection(ctx, input)
			},
		},
		"test-identity-provider": {
			Name:   "test-identity-provider",
			Fields: fields_test_identity_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestIdentityProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_identity_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestIdentityProvider(ctx, input)
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
		"update-access": {
			Name:   "update-access",
			Fields: fields_update_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccess(ctx, input)
			},
		},
		"update-agreement": {
			Name:   "update-agreement",
			Fields: fields_update_agreement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAgreementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_agreement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAgreement(ctx, input)
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
		"update-connector": {
			Name:   "update-connector",
			Fields: fields_update_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConnector(ctx, input)
			},
		},
		"update-host-key": {
			Name:   "update-host-key",
			Fields: fields_update_host_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateHostKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_host_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateHostKey(ctx, input)
			},
		},
		"update-profile": {
			Name:   "update-profile",
			Fields: fields_update_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProfile(ctx, input)
			},
		},
		"update-server": {
			Name:   "update-server",
			Fields: fields_update_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateServer(ctx, input)
			},
		},
		"update-user": {
			Name:   "update-user",
			Fields: fields_update_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUser(ctx, input)
			},
		},
		"update-web-app": {
			Name:   "update-web-app",
			Fields: fields_update_web_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWebAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_web_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWebApp(ctx, input)
			},
		},
		"update-web-app-customization": {
			Name:   "update-web-app-customization",
			Fields: fields_update_web_app_customization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWebAppCustomizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_web_app_customization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWebAppCustomization(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("transfer", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
