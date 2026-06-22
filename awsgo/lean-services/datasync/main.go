package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/datasync"
)

var fields_cancel_task_execution = []leanruntime.Field{
	{Name: "TaskExecutionArn", Flag: "task-execution-arn", Type: "*string", Required: true},
}

var fields_create_agent = []leanruntime.Field{
	{Name: "ActivationKey", Flag: "activation-key", Type: "*string", Required: true},
	{Name: "AgentName", Flag: "agent-name", Type: "*string", Required: false},
	{Name: "SecurityGroupArns", Flag: "security-group-arns", Type: "[]string", Required: false},
	{Name: "SubnetArns", Flag: "subnet-arns", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.TagListEntry", Required: false},
	{Name: "VpcEndpointId", Flag: "vpc-endpoint-id", Type: "*string", Required: false},
}

var fields_create_location_azure_blob = []leanruntime.Field{
	{Name: "AccessTier", Flag: "access-tier", Type: "types.AzureAccessTier", Required: false},
	{Name: "AgentArns", Flag: "agent-arns", Type: "[]string", Required: false},
	{Name: "AuthenticationType", Flag: "authentication-type", Type: "types.AzureBlobAuthenticationType", Required: true},
	{Name: "BlobType", Flag: "blob-type", Type: "types.AzureBlobType", Required: false},
	{Name: "CmkSecretConfig", Flag: "cmk-secret-config", Type: "*types.CmkSecretConfig", Required: false},
	{Name: "ContainerUrl", Flag: "container-url", Type: "*string", Required: true},
	{Name: "CustomSecretConfig", Flag: "custom-secret-config", Type: "*types.CustomSecretConfig", Required: false},
	{Name: "SasConfiguration", Flag: "sas-configuration", Type: "*types.AzureBlobSasConfiguration", Required: false},
	{Name: "Subdirectory", Flag: "subdirectory", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.TagListEntry", Required: false},
}

var fields_create_location_efs = []leanruntime.Field{
	{Name: "AccessPointArn", Flag: "access-point-arn", Type: "*string", Required: false},
	{Name: "Ec2Config", Flag: "ec2-config", Type: "*types.Ec2Config", Required: true},
	{Name: "EfsFilesystemArn", Flag: "efs-filesystem-arn", Type: "*string", Required: true},
	{Name: "FileSystemAccessRoleArn", Flag: "file-system-access-role-arn", Type: "*string", Required: false},
	{Name: "InTransitEncryption", Flag: "in-transit-encryption", Type: "types.EfsInTransitEncryption", Required: false},
	{Name: "Subdirectory", Flag: "subdirectory", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.TagListEntry", Required: false},
}

var fields_create_location_fsx_lustre = []leanruntime.Field{
	{Name: "FsxFilesystemArn", Flag: "fsx-filesystem-arn", Type: "*string", Required: true},
	{Name: "SecurityGroupArns", Flag: "security-group-arns", Type: "[]string", Required: true},
	{Name: "Subdirectory", Flag: "subdirectory", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.TagListEntry", Required: false},
}

var fields_create_location_fsx_ontap = []leanruntime.Field{
	{Name: "Protocol", Flag: "protocol", Type: "*types.FsxProtocol", Required: true},
	{Name: "SecurityGroupArns", Flag: "security-group-arns", Type: "[]string", Required: true},
	{Name: "StorageVirtualMachineArn", Flag: "storage-virtual-machine-arn", Type: "*string", Required: true},
	{Name: "Subdirectory", Flag: "subdirectory", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.TagListEntry", Required: false},
}

var fields_create_location_fsx_open_zfs = []leanruntime.Field{
	{Name: "FsxFilesystemArn", Flag: "fsx-filesystem-arn", Type: "*string", Required: true},
	{Name: "Protocol", Flag: "protocol", Type: "*types.FsxProtocol", Required: true},
	{Name: "SecurityGroupArns", Flag: "security-group-arns", Type: "[]string", Required: true},
	{Name: "Subdirectory", Flag: "subdirectory", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.TagListEntry", Required: false},
}

var fields_create_location_fsx_windows = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "FsxFilesystemArn", Flag: "fsx-filesystem-arn", Type: "*string", Required: true},
	{Name: "Password", Flag: "password", Type: "*string", Required: true},
	{Name: "SecurityGroupArns", Flag: "security-group-arns", Type: "[]string", Required: true},
	{Name: "Subdirectory", Flag: "subdirectory", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.TagListEntry", Required: false},
	{Name: "User", Flag: "user", Type: "*string", Required: true},
}

var fields_create_location_hdfs = []leanruntime.Field{
	{Name: "AgentArns", Flag: "agent-arns", Type: "[]string", Required: true},
	{Name: "AuthenticationType", Flag: "authentication-type", Type: "types.HdfsAuthenticationType", Required: true},
	{Name: "BlockSize", Flag: "block-size", Type: "*int32", Required: false},
	{Name: "KerberosKeytab", Flag: "kerberos-keytab", Type: "[]byte", Required: false},
	{Name: "KerberosKrb5Conf", Flag: "kerberos-krb5-conf", Type: "[]byte", Required: false},
	{Name: "KerberosPrincipal", Flag: "kerberos-principal", Type: "*string", Required: false},
	{Name: "KmsKeyProviderUri", Flag: "kms-key-provider-uri", Type: "*string", Required: false},
	{Name: "NameNodes", Flag: "name-nodes", Type: "[]types.HdfsNameNode", Required: true},
	{Name: "QopConfiguration", Flag: "qop-configuration", Type: "*types.QopConfiguration", Required: false},
	{Name: "ReplicationFactor", Flag: "replication-factor", Type: "*int32", Required: false},
	{Name: "SimpleUser", Flag: "simple-user", Type: "*string", Required: false},
	{Name: "Subdirectory", Flag: "subdirectory", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.TagListEntry", Required: false},
}

var fields_create_location_nfs = []leanruntime.Field{
	{Name: "MountOptions", Flag: "mount-options", Type: "*types.NfsMountOptions", Required: false},
	{Name: "OnPremConfig", Flag: "on-prem-config", Type: "*types.OnPremConfig", Required: true},
	{Name: "ServerHostname", Flag: "server-hostname", Type: "*string", Required: true},
	{Name: "Subdirectory", Flag: "subdirectory", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.TagListEntry", Required: false},
}

var fields_create_location_object_storage = []leanruntime.Field{
	{Name: "AccessKey", Flag: "access-key", Type: "*string", Required: false},
	{Name: "AgentArns", Flag: "agent-arns", Type: "[]string", Required: false},
	{Name: "BucketName", Flag: "bucket-name", Type: "*string", Required: true},
	{Name: "CmkSecretConfig", Flag: "cmk-secret-config", Type: "*types.CmkSecretConfig", Required: false},
	{Name: "CustomSecretConfig", Flag: "custom-secret-config", Type: "*types.CustomSecretConfig", Required: false},
	{Name: "SecretKey", Flag: "secret-key", Type: "*string", Required: false},
	{Name: "ServerCertificate", Flag: "server-certificate", Type: "[]byte", Required: false},
	{Name: "ServerHostname", Flag: "server-hostname", Type: "*string", Required: true},
	{Name: "ServerPort", Flag: "server-port", Type: "*int32", Required: false},
	{Name: "ServerProtocol", Flag: "server-protocol", Type: "types.ObjectStorageServerProtocol", Required: false},
	{Name: "Subdirectory", Flag: "subdirectory", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.TagListEntry", Required: false},
}

var fields_create_location_s3 = []leanruntime.Field{
	{Name: "AgentArns", Flag: "agent-arns", Type: "[]string", Required: false},
	{Name: "S3BucketArn", Flag: "s3-bucket-arn", Type: "*string", Required: true},
	{Name: "S3Config", Flag: "s3-config", Type: "*types.S3Config", Required: true},
	{Name: "S3StorageClass", Flag: "s3-storage-class", Type: "types.S3StorageClass", Required: false},
	{Name: "Subdirectory", Flag: "subdirectory", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.TagListEntry", Required: false},
}

var fields_create_location_smb = []leanruntime.Field{
	{Name: "AgentArns", Flag: "agent-arns", Type: "[]string", Required: true},
	{Name: "AuthenticationType", Flag: "authentication-type", Type: "types.SmbAuthenticationType", Required: false},
	{Name: "CmkSecretConfig", Flag: "cmk-secret-config", Type: "*types.CmkSecretConfig", Required: false},
	{Name: "CustomSecretConfig", Flag: "custom-secret-config", Type: "*types.CustomSecretConfig", Required: false},
	{Name: "DnsIpAddresses", Flag: "dns-ip-addresses", Type: "[]string", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "KerberosKeytab", Flag: "kerberos-keytab", Type: "[]byte", Required: false},
	{Name: "KerberosKrb5Conf", Flag: "kerberos-krb5-conf", Type: "[]byte", Required: false},
	{Name: "KerberosPrincipal", Flag: "kerberos-principal", Type: "*string", Required: false},
	{Name: "MountOptions", Flag: "mount-options", Type: "*types.SmbMountOptions", Required: false},
	{Name: "Password", Flag: "password", Type: "*string", Required: false},
	{Name: "ServerHostname", Flag: "server-hostname", Type: "*string", Required: true},
	{Name: "Subdirectory", Flag: "subdirectory", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.TagListEntry", Required: false},
	{Name: "User", Flag: "user", Type: "*string", Required: false},
}

var fields_create_task = []leanruntime.Field{
	{Name: "CloudWatchLogGroupArn", Flag: "cloud-watch-log-group-arn", Type: "*string", Required: false},
	{Name: "DestinationLocationArn", Flag: "destination-location-arn", Type: "*string", Required: true},
	{Name: "Excludes", Flag: "excludes", Type: "[]types.FilterRule", Required: false},
	{Name: "Includes", Flag: "includes", Type: "[]types.FilterRule", Required: false},
	{Name: "ManifestConfig", Flag: "manifest-config", Type: "*types.ManifestConfig", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Options", Flag: "options", Type: "*types.Options", Required: false},
	{Name: "Schedule", Flag: "schedule", Type: "*types.TaskSchedule", Required: false},
	{Name: "SourceLocationArn", Flag: "source-location-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.TagListEntry", Required: false},
	{Name: "TaskMode", Flag: "task-mode", Type: "types.TaskMode", Required: false},
	{Name: "TaskReportConfig", Flag: "task-report-config", Type: "*types.TaskReportConfig", Required: false},
}

var fields_delete_agent = []leanruntime.Field{
	{Name: "AgentArn", Flag: "agent-arn", Type: "*string", Required: true},
}

var fields_delete_location = []leanruntime.Field{
	{Name: "LocationArn", Flag: "location-arn", Type: "*string", Required: true},
}

var fields_delete_task = []leanruntime.Field{
	{Name: "TaskArn", Flag: "task-arn", Type: "*string", Required: true},
}

var fields_describe_agent = []leanruntime.Field{
	{Name: "AgentArn", Flag: "agent-arn", Type: "*string", Required: true},
}

var fields_describe_location_azure_blob = []leanruntime.Field{
	{Name: "LocationArn", Flag: "location-arn", Type: "*string", Required: true},
}

var fields_describe_location_efs = []leanruntime.Field{
	{Name: "LocationArn", Flag: "location-arn", Type: "*string", Required: true},
}

var fields_describe_location_fsx_lustre = []leanruntime.Field{
	{Name: "LocationArn", Flag: "location-arn", Type: "*string", Required: true},
}

var fields_describe_location_fsx_ontap = []leanruntime.Field{
	{Name: "LocationArn", Flag: "location-arn", Type: "*string", Required: true},
}

var fields_describe_location_fsx_open_zfs = []leanruntime.Field{
	{Name: "LocationArn", Flag: "location-arn", Type: "*string", Required: true},
}

var fields_describe_location_fsx_windows = []leanruntime.Field{
	{Name: "LocationArn", Flag: "location-arn", Type: "*string", Required: true},
}

var fields_describe_location_hdfs = []leanruntime.Field{
	{Name: "LocationArn", Flag: "location-arn", Type: "*string", Required: true},
}

var fields_describe_location_nfs = []leanruntime.Field{
	{Name: "LocationArn", Flag: "location-arn", Type: "*string", Required: true},
}

var fields_describe_location_object_storage = []leanruntime.Field{
	{Name: "LocationArn", Flag: "location-arn", Type: "*string", Required: true},
}

var fields_describe_location_s3 = []leanruntime.Field{
	{Name: "LocationArn", Flag: "location-arn", Type: "*string", Required: true},
}

var fields_describe_location_smb = []leanruntime.Field{
	{Name: "LocationArn", Flag: "location-arn", Type: "*string", Required: true},
}

var fields_describe_task = []leanruntime.Field{
	{Name: "TaskArn", Flag: "task-arn", Type: "*string", Required: true},
}

var fields_describe_task_execution = []leanruntime.Field{
	{Name: "TaskExecutionArn", Flag: "task-execution-arn", Type: "*string", Required: true},
}

var fields_list_agents = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_locations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.LocationFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_task_executions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TaskArn", Flag: "task-arn", Type: "*string", Required: false},
}

var fields_list_tasks = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.TaskFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_start_task_execution = []leanruntime.Field{
	{Name: "Excludes", Flag: "excludes", Type: "[]types.FilterRule", Required: false},
	{Name: "Includes", Flag: "includes", Type: "[]types.FilterRule", Required: false},
	{Name: "ManifestConfig", Flag: "manifest-config", Type: "*types.ManifestConfig", Required: false},
	{Name: "OverrideOptions", Flag: "override-options", Type: "*types.Options", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.TagListEntry", Required: false},
	{Name: "TaskArn", Flag: "task-arn", Type: "*string", Required: true},
	{Name: "TaskReportConfig", Flag: "task-report-config", Type: "*types.TaskReportConfig", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.TagListEntry", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "Keys", Flag: "keys", Type: "[]string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_update_agent = []leanruntime.Field{
	{Name: "AgentArn", Flag: "agent-arn", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_location_azure_blob = []leanruntime.Field{
	{Name: "AccessTier", Flag: "access-tier", Type: "types.AzureAccessTier", Required: false},
	{Name: "AgentArns", Flag: "agent-arns", Type: "[]string", Required: false},
	{Name: "AuthenticationType", Flag: "authentication-type", Type: "types.AzureBlobAuthenticationType", Required: false},
	{Name: "BlobType", Flag: "blob-type", Type: "types.AzureBlobType", Required: false},
	{Name: "CmkSecretConfig", Flag: "cmk-secret-config", Type: "*types.CmkSecretConfig", Required: false},
	{Name: "CustomSecretConfig", Flag: "custom-secret-config", Type: "*types.CustomSecretConfig", Required: false},
	{Name: "LocationArn", Flag: "location-arn", Type: "*string", Required: true},
	{Name: "SasConfiguration", Flag: "sas-configuration", Type: "*types.AzureBlobSasConfiguration", Required: false},
	{Name: "Subdirectory", Flag: "subdirectory", Type: "*string", Required: false},
}

var fields_update_location_efs = []leanruntime.Field{
	{Name: "AccessPointArn", Flag: "access-point-arn", Type: "*string", Required: false},
	{Name: "FileSystemAccessRoleArn", Flag: "file-system-access-role-arn", Type: "*string", Required: false},
	{Name: "InTransitEncryption", Flag: "in-transit-encryption", Type: "types.EfsInTransitEncryption", Required: false},
	{Name: "LocationArn", Flag: "location-arn", Type: "*string", Required: true},
	{Name: "Subdirectory", Flag: "subdirectory", Type: "*string", Required: false},
}

var fields_update_location_fsx_lustre = []leanruntime.Field{
	{Name: "LocationArn", Flag: "location-arn", Type: "*string", Required: true},
	{Name: "Subdirectory", Flag: "subdirectory", Type: "*string", Required: false},
}

var fields_update_location_fsx_ontap = []leanruntime.Field{
	{Name: "LocationArn", Flag: "location-arn", Type: "*string", Required: true},
	{Name: "Protocol", Flag: "protocol", Type: "*types.FsxUpdateProtocol", Required: false},
	{Name: "Subdirectory", Flag: "subdirectory", Type: "*string", Required: false},
}

var fields_update_location_fsx_open_zfs = []leanruntime.Field{
	{Name: "LocationArn", Flag: "location-arn", Type: "*string", Required: true},
	{Name: "Protocol", Flag: "protocol", Type: "*types.FsxProtocol", Required: false},
	{Name: "Subdirectory", Flag: "subdirectory", Type: "*string", Required: false},
}

var fields_update_location_fsx_windows = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "LocationArn", Flag: "location-arn", Type: "*string", Required: true},
	{Name: "Password", Flag: "password", Type: "*string", Required: false},
	{Name: "Subdirectory", Flag: "subdirectory", Type: "*string", Required: false},
	{Name: "User", Flag: "user", Type: "*string", Required: false},
}

var fields_update_location_hdfs = []leanruntime.Field{
	{Name: "AgentArns", Flag: "agent-arns", Type: "[]string", Required: false},
	{Name: "AuthenticationType", Flag: "authentication-type", Type: "types.HdfsAuthenticationType", Required: false},
	{Name: "BlockSize", Flag: "block-size", Type: "*int32", Required: false},
	{Name: "KerberosKeytab", Flag: "kerberos-keytab", Type: "[]byte", Required: false},
	{Name: "KerberosKrb5Conf", Flag: "kerberos-krb5-conf", Type: "[]byte", Required: false},
	{Name: "KerberosPrincipal", Flag: "kerberos-principal", Type: "*string", Required: false},
	{Name: "KmsKeyProviderUri", Flag: "kms-key-provider-uri", Type: "*string", Required: false},
	{Name: "LocationArn", Flag: "location-arn", Type: "*string", Required: true},
	{Name: "NameNodes", Flag: "name-nodes", Type: "[]types.HdfsNameNode", Required: false},
	{Name: "QopConfiguration", Flag: "qop-configuration", Type: "*types.QopConfiguration", Required: false},
	{Name: "ReplicationFactor", Flag: "replication-factor", Type: "*int32", Required: false},
	{Name: "SimpleUser", Flag: "simple-user", Type: "*string", Required: false},
	{Name: "Subdirectory", Flag: "subdirectory", Type: "*string", Required: false},
}

var fields_update_location_nfs = []leanruntime.Field{
	{Name: "LocationArn", Flag: "location-arn", Type: "*string", Required: true},
	{Name: "MountOptions", Flag: "mount-options", Type: "*types.NfsMountOptions", Required: false},
	{Name: "OnPremConfig", Flag: "on-prem-config", Type: "*types.OnPremConfig", Required: false},
	{Name: "ServerHostname", Flag: "server-hostname", Type: "*string", Required: false},
	{Name: "Subdirectory", Flag: "subdirectory", Type: "*string", Required: false},
}

var fields_update_location_object_storage = []leanruntime.Field{
	{Name: "AccessKey", Flag: "access-key", Type: "*string", Required: false},
	{Name: "AgentArns", Flag: "agent-arns", Type: "[]string", Required: false},
	{Name: "CmkSecretConfig", Flag: "cmk-secret-config", Type: "*types.CmkSecretConfig", Required: false},
	{Name: "CustomSecretConfig", Flag: "custom-secret-config", Type: "*types.CustomSecretConfig", Required: false},
	{Name: "LocationArn", Flag: "location-arn", Type: "*string", Required: true},
	{Name: "SecretKey", Flag: "secret-key", Type: "*string", Required: false},
	{Name: "ServerCertificate", Flag: "server-certificate", Type: "[]byte", Required: false},
	{Name: "ServerHostname", Flag: "server-hostname", Type: "*string", Required: false},
	{Name: "ServerPort", Flag: "server-port", Type: "*int32", Required: false},
	{Name: "ServerProtocol", Flag: "server-protocol", Type: "types.ObjectStorageServerProtocol", Required: false},
	{Name: "Subdirectory", Flag: "subdirectory", Type: "*string", Required: false},
}

var fields_update_location_s3 = []leanruntime.Field{
	{Name: "LocationArn", Flag: "location-arn", Type: "*string", Required: true},
	{Name: "S3Config", Flag: "s3-config", Type: "*types.S3Config", Required: false},
	{Name: "S3StorageClass", Flag: "s3-storage-class", Type: "types.S3StorageClass", Required: false},
	{Name: "Subdirectory", Flag: "subdirectory", Type: "*string", Required: false},
}

var fields_update_location_smb = []leanruntime.Field{
	{Name: "AgentArns", Flag: "agent-arns", Type: "[]string", Required: false},
	{Name: "AuthenticationType", Flag: "authentication-type", Type: "types.SmbAuthenticationType", Required: false},
	{Name: "CmkSecretConfig", Flag: "cmk-secret-config", Type: "*types.CmkSecretConfig", Required: false},
	{Name: "CustomSecretConfig", Flag: "custom-secret-config", Type: "*types.CustomSecretConfig", Required: false},
	{Name: "DnsIpAddresses", Flag: "dns-ip-addresses", Type: "[]string", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "KerberosKeytab", Flag: "kerberos-keytab", Type: "[]byte", Required: false},
	{Name: "KerberosKrb5Conf", Flag: "kerberos-krb5-conf", Type: "[]byte", Required: false},
	{Name: "KerberosPrincipal", Flag: "kerberos-principal", Type: "*string", Required: false},
	{Name: "LocationArn", Flag: "location-arn", Type: "*string", Required: true},
	{Name: "MountOptions", Flag: "mount-options", Type: "*types.SmbMountOptions", Required: false},
	{Name: "Password", Flag: "password", Type: "*string", Required: false},
	{Name: "ServerHostname", Flag: "server-hostname", Type: "*string", Required: false},
	{Name: "Subdirectory", Flag: "subdirectory", Type: "*string", Required: false},
	{Name: "User", Flag: "user", Type: "*string", Required: false},
}

var fields_update_task = []leanruntime.Field{
	{Name: "CloudWatchLogGroupArn", Flag: "cloud-watch-log-group-arn", Type: "*string", Required: false},
	{Name: "Excludes", Flag: "excludes", Type: "[]types.FilterRule", Required: false},
	{Name: "Includes", Flag: "includes", Type: "[]types.FilterRule", Required: false},
	{Name: "ManifestConfig", Flag: "manifest-config", Type: "*types.ManifestConfig", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Options", Flag: "options", Type: "*types.Options", Required: false},
	{Name: "Schedule", Flag: "schedule", Type: "*types.TaskSchedule", Required: false},
	{Name: "TaskArn", Flag: "task-arn", Type: "*string", Required: true},
	{Name: "TaskReportConfig", Flag: "task-report-config", Type: "*types.TaskReportConfig", Required: false},
}

var fields_update_task_execution = []leanruntime.Field{
	{Name: "Options", Flag: "options", Type: "*types.Options", Required: true},
	{Name: "TaskExecutionArn", Flag: "task-execution-arn", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"cancel-task-execution": {
			Name:   "cancel-task-execution",
			Fields: fields_cancel_task_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelTaskExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_task_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelTaskExecution(ctx, input)
			},
		},
		"create-agent": {
			Name:   "create-agent",
			Fields: fields_create_agent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAgentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_agent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAgent(ctx, input)
			},
		},
		"create-location-azure-blob": {
			Name:   "create-location-azure-blob",
			Fields: fields_create_location_azure_blob,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLocationAzureBlobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_location_azure_blob, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLocationAzureBlob(ctx, input)
			},
		},
		"create-location-efs": {
			Name:   "create-location-efs",
			Fields: fields_create_location_efs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLocationEfsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_location_efs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLocationEfs(ctx, input)
			},
		},
		"create-location-fsx-lustre": {
			Name:   "create-location-fsx-lustre",
			Fields: fields_create_location_fsx_lustre,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLocationFsxLustreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_location_fsx_lustre, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLocationFsxLustre(ctx, input)
			},
		},
		"create-location-fsx-ontap": {
			Name:   "create-location-fsx-ontap",
			Fields: fields_create_location_fsx_ontap,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLocationFsxOntapInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_location_fsx_ontap, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLocationFsxOntap(ctx, input)
			},
		},
		"create-location-fsx-open-zfs": {
			Name:   "create-location-fsx-open-zfs",
			Fields: fields_create_location_fsx_open_zfs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLocationFsxOpenZfsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_location_fsx_open_zfs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLocationFsxOpenZfs(ctx, input)
			},
		},
		"create-location-fsx-windows": {
			Name:   "create-location-fsx-windows",
			Fields: fields_create_location_fsx_windows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLocationFsxWindowsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_location_fsx_windows, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLocationFsxWindows(ctx, input)
			},
		},
		"create-location-hdfs": {
			Name:   "create-location-hdfs",
			Fields: fields_create_location_hdfs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLocationHdfsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_location_hdfs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLocationHdfs(ctx, input)
			},
		},
		"create-location-nfs": {
			Name:   "create-location-nfs",
			Fields: fields_create_location_nfs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLocationNfsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_location_nfs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLocationNfs(ctx, input)
			},
		},
		"create-location-object-storage": {
			Name:   "create-location-object-storage",
			Fields: fields_create_location_object_storage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLocationObjectStorageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_location_object_storage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLocationObjectStorage(ctx, input)
			},
		},
		"create-location-s3": {
			Name:   "create-location-s3",
			Fields: fields_create_location_s3,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLocationS3Input{}
				if _, err := leanruntime.ApplyInput(input, fields_create_location_s3, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLocationS3(ctx, input)
			},
		},
		"create-location-smb": {
			Name:   "create-location-smb",
			Fields: fields_create_location_smb,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLocationSmbInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_location_smb, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLocationSmb(ctx, input)
			},
		},
		"create-task": {
			Name:   "create-task",
			Fields: fields_create_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTask(ctx, input)
			},
		},
		"delete-agent": {
			Name:   "delete-agent",
			Fields: fields_delete_agent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAgentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_agent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAgent(ctx, input)
			},
		},
		"delete-location": {
			Name:   "delete-location",
			Fields: fields_delete_location,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_location, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLocation(ctx, input)
			},
		},
		"delete-task": {
			Name:   "delete-task",
			Fields: fields_delete_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTask(ctx, input)
			},
		},
		"describe-agent": {
			Name:   "describe-agent",
			Fields: fields_describe_agent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAgentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_agent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAgent(ctx, input)
			},
		},
		"describe-location-azure-blob": {
			Name:   "describe-location-azure-blob",
			Fields: fields_describe_location_azure_blob,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLocationAzureBlobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_location_azure_blob, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLocationAzureBlob(ctx, input)
			},
		},
		"describe-location-efs": {
			Name:   "describe-location-efs",
			Fields: fields_describe_location_efs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLocationEfsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_location_efs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLocationEfs(ctx, input)
			},
		},
		"describe-location-fsx-lustre": {
			Name:   "describe-location-fsx-lustre",
			Fields: fields_describe_location_fsx_lustre,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLocationFsxLustreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_location_fsx_lustre, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLocationFsxLustre(ctx, input)
			},
		},
		"describe-location-fsx-ontap": {
			Name:   "describe-location-fsx-ontap",
			Fields: fields_describe_location_fsx_ontap,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLocationFsxOntapInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_location_fsx_ontap, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLocationFsxOntap(ctx, input)
			},
		},
		"describe-location-fsx-open-zfs": {
			Name:   "describe-location-fsx-open-zfs",
			Fields: fields_describe_location_fsx_open_zfs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLocationFsxOpenZfsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_location_fsx_open_zfs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLocationFsxOpenZfs(ctx, input)
			},
		},
		"describe-location-fsx-windows": {
			Name:   "describe-location-fsx-windows",
			Fields: fields_describe_location_fsx_windows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLocationFsxWindowsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_location_fsx_windows, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLocationFsxWindows(ctx, input)
			},
		},
		"describe-location-hdfs": {
			Name:   "describe-location-hdfs",
			Fields: fields_describe_location_hdfs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLocationHdfsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_location_hdfs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLocationHdfs(ctx, input)
			},
		},
		"describe-location-nfs": {
			Name:   "describe-location-nfs",
			Fields: fields_describe_location_nfs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLocationNfsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_location_nfs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLocationNfs(ctx, input)
			},
		},
		"describe-location-object-storage": {
			Name:   "describe-location-object-storage",
			Fields: fields_describe_location_object_storage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLocationObjectStorageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_location_object_storage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLocationObjectStorage(ctx, input)
			},
		},
		"describe-location-s3": {
			Name:   "describe-location-s3",
			Fields: fields_describe_location_s3,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLocationS3Input{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_location_s3, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLocationS3(ctx, input)
			},
		},
		"describe-location-smb": {
			Name:   "describe-location-smb",
			Fields: fields_describe_location_smb,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLocationSmbInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_location_smb, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLocationSmb(ctx, input)
			},
		},
		"describe-task": {
			Name:   "describe-task",
			Fields: fields_describe_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTask(ctx, input)
			},
		},
		"describe-task-execution": {
			Name:   "describe-task-execution",
			Fields: fields_describe_task_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTaskExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_task_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTaskExecution(ctx, input)
			},
		},
		"list-agents": {
			Name:   "list-agents",
			Fields: fields_list_agents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAgentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_agents, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAgents(ctx, input)
				}
				var results []*svc.ListAgentsOutput
				p := svc.NewListAgentsPaginator(client, input)
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
		"list-locations": {
			Name:   "list-locations",
			Fields: fields_list_locations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLocationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_locations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLocations(ctx, input)
				}
				var results []*svc.ListLocationsOutput
				p := svc.NewListLocationsPaginator(client, input)
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
		"list-task-executions": {
			Name:   "list-task-executions",
			Fields: fields_list_task_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTaskExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_task_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTaskExecutions(ctx, input)
				}
				var results []*svc.ListTaskExecutionsOutput
				p := svc.NewListTaskExecutionsPaginator(client, input)
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
		"list-tasks": {
			Name:   "list-tasks",
			Fields: fields_list_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTasks(ctx, input)
				}
				var results []*svc.ListTasksOutput
				p := svc.NewListTasksPaginator(client, input)
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
		"start-task-execution": {
			Name:   "start-task-execution",
			Fields: fields_start_task_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartTaskExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_task_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartTaskExecution(ctx, input)
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
		"update-agent": {
			Name:   "update-agent",
			Fields: fields_update_agent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAgentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_agent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAgent(ctx, input)
			},
		},
		"update-location-azure-blob": {
			Name:   "update-location-azure-blob",
			Fields: fields_update_location_azure_blob,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLocationAzureBlobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_location_azure_blob, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLocationAzureBlob(ctx, input)
			},
		},
		"update-location-efs": {
			Name:   "update-location-efs",
			Fields: fields_update_location_efs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLocationEfsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_location_efs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLocationEfs(ctx, input)
			},
		},
		"update-location-fsx-lustre": {
			Name:   "update-location-fsx-lustre",
			Fields: fields_update_location_fsx_lustre,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLocationFsxLustreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_location_fsx_lustre, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLocationFsxLustre(ctx, input)
			},
		},
		"update-location-fsx-ontap": {
			Name:   "update-location-fsx-ontap",
			Fields: fields_update_location_fsx_ontap,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLocationFsxOntapInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_location_fsx_ontap, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLocationFsxOntap(ctx, input)
			},
		},
		"update-location-fsx-open-zfs": {
			Name:   "update-location-fsx-open-zfs",
			Fields: fields_update_location_fsx_open_zfs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLocationFsxOpenZfsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_location_fsx_open_zfs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLocationFsxOpenZfs(ctx, input)
			},
		},
		"update-location-fsx-windows": {
			Name:   "update-location-fsx-windows",
			Fields: fields_update_location_fsx_windows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLocationFsxWindowsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_location_fsx_windows, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLocationFsxWindows(ctx, input)
			},
		},
		"update-location-hdfs": {
			Name:   "update-location-hdfs",
			Fields: fields_update_location_hdfs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLocationHdfsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_location_hdfs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLocationHdfs(ctx, input)
			},
		},
		"update-location-nfs": {
			Name:   "update-location-nfs",
			Fields: fields_update_location_nfs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLocationNfsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_location_nfs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLocationNfs(ctx, input)
			},
		},
		"update-location-object-storage": {
			Name:   "update-location-object-storage",
			Fields: fields_update_location_object_storage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLocationObjectStorageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_location_object_storage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLocationObjectStorage(ctx, input)
			},
		},
		"update-location-s3": {
			Name:   "update-location-s3",
			Fields: fields_update_location_s3,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLocationS3Input{}
				if _, err := leanruntime.ApplyInput(input, fields_update_location_s3, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLocationS3(ctx, input)
			},
		},
		"update-location-smb": {
			Name:   "update-location-smb",
			Fields: fields_update_location_smb,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLocationSmbInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_location_smb, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLocationSmb(ctx, input)
			},
		},
		"update-task": {
			Name:   "update-task",
			Fields: fields_update_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTask(ctx, input)
			},
		},
		"update-task-execution": {
			Name:   "update-task-execution",
			Fields: fields_update_task_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTaskExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_task_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTaskExecution(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("datasync", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
