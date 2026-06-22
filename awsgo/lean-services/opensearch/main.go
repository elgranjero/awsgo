package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/opensearch"
)

var fields_accept_inbound_connection = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
}

var fields_add_data_source = []leanruntime.Field{
	{Name: "DataSourceType", Flag: "data-source-type", Type: "types.DataSourceType", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_add_direct_query_data_source = []leanruntime.Field{
	{Name: "DataSourceName", Flag: "data-source-name", Type: "*string", Required: true},
	{Name: "DataSourceType", Flag: "data-source-type", Type: "types.DirectQueryDataSourceType", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "OpenSearchArns", Flag: "open-search-arns", Type: "[]string", Required: true},
	{Name: "TagList", Flag: "tag-list", Type: "[]types.Tag", Required: false},
}

var fields_add_tags = []leanruntime.Field{
	{Name: "ARN", Flag: "arn", Type: "*string", Required: true},
	{Name: "TagList", Flag: "tag-list", Type: "[]types.Tag", Required: true},
}

var fields_associate_package = []leanruntime.Field{
	{Name: "AssociationConfiguration", Flag: "association-configuration", Type: "*types.PackageAssociationConfiguration", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "PackageID", Flag: "package-id", Type: "*string", Required: true},
	{Name: "PrerequisitePackageIDList", Flag: "prerequisite-package-id-list", Type: "[]string", Required: false},
}

var fields_associate_packages = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "PackageList", Flag: "package-list", Type: "[]types.PackageDetailsForAssociation", Required: true},
}

var fields_authorize_vpc_endpoint_access = []leanruntime.Field{
	{Name: "Account", Flag: "account", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Service", Flag: "service", Type: "types.AWSServicePrincipal", Required: false},
}

var fields_cancel_domain_config_change = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_cancel_service_software_update = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_create_application = []leanruntime.Field{
	{Name: "AppConfigs", Flag: "app-configs", Type: "[]types.AppConfig", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DataSources", Flag: "data-sources", Type: "[]types.DataSource", Required: false},
	{Name: "IamIdentityCenterOptions", Flag: "iam-identity-center-options", Type: "*types.IamIdentityCenterOptionsInput", Required: false},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "TagList", Flag: "tag-list", Type: "[]types.Tag", Required: false},
}

var fields_create_domain = []leanruntime.Field{
	{Name: "AIMLOptions", Flag: "aiml-options", Type: "*types.AIMLOptionsInput", Required: false},
	{Name: "AccessPolicies", Flag: "access-policies", Type: "*string", Required: false},
	{Name: "AdvancedOptions", Flag: "advanced-options", Type: "map[string]string", Required: false},
	{Name: "AdvancedSecurityOptions", Flag: "advanced-security-options", Type: "*types.AdvancedSecurityOptionsInput", Required: false},
	{Name: "AutoTuneOptions", Flag: "auto-tune-options", Type: "*types.AutoTuneOptionsInput", Required: false},
	{Name: "ClusterConfig", Flag: "cluster-config", Type: "*types.ClusterConfig", Required: false},
	{Name: "CognitoOptions", Flag: "cognito-options", Type: "*types.CognitoOptions", Required: false},
	{Name: "DomainEndpointOptions", Flag: "domain-endpoint-options", Type: "*types.DomainEndpointOptions", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "EBSOptions", Flag: "ebs-options", Type: "*types.EBSOptions", Required: false},
	{Name: "EncryptionAtRestOptions", Flag: "encryption-at-rest-options", Type: "*types.EncryptionAtRestOptions", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "IPAddressType", Flag: "ip-address-type", Type: "types.IPAddressType", Required: false},
	{Name: "IdentityCenterOptions", Flag: "identity-center-options", Type: "*types.IdentityCenterOptionsInput", Required: false},
	{Name: "LogPublishingOptions", Flag: "log-publishing-options", Type: "map[string]types.LogPublishingOption", Required: false},
	{Name: "NodeToNodeEncryptionOptions", Flag: "node-to-node-encryption-options", Type: "*types.NodeToNodeEncryptionOptions", Required: false},
	{Name: "OffPeakWindowOptions", Flag: "off-peak-window-options", Type: "*types.OffPeakWindowOptions", Required: false},
	{Name: "SnapshotOptions", Flag: "snapshot-options", Type: "*types.SnapshotOptions", Required: false},
	{Name: "SoftwareUpdateOptions", Flag: "software-update-options", Type: "*types.SoftwareUpdateOptions", Required: false},
	{Name: "TagList", Flag: "tag-list", Type: "[]types.Tag", Required: false},
	{Name: "VPCOptions", Flag: "vpc-options", Type: "*types.VPCOptions", Required: false},
}

var fields_create_index = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: true},
	{Name: "IndexSchema", Flag: "index-schema", Type: "document.Interface", Required: true},
}

var fields_create_outbound_connection = []leanruntime.Field{
	{Name: "ConnectionAlias", Flag: "connection-alias", Type: "*string", Required: true},
	{Name: "ConnectionMode", Flag: "connection-mode", Type: "types.ConnectionMode", Required: false},
	{Name: "ConnectionProperties", Flag: "connection-properties", Type: "*types.ConnectionProperties", Required: false},
	{Name: "LocalDomainInfo", Flag: "local-domain-info", Type: "*types.DomainInformationContainer", Required: true},
	{Name: "RemoteDomainInfo", Flag: "remote-domain-info", Type: "*types.DomainInformationContainer", Required: true},
}

var fields_create_package = []leanruntime.Field{
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "PackageConfiguration", Flag: "package-configuration", Type: "*types.PackageConfiguration", Required: false},
	{Name: "PackageDescription", Flag: "package-description", Type: "*string", Required: false},
	{Name: "PackageEncryptionOptions", Flag: "package-encryption-options", Type: "*types.PackageEncryptionOptions", Required: false},
	{Name: "PackageName", Flag: "package-name", Type: "*string", Required: true},
	{Name: "PackageSource", Flag: "package-source", Type: "*types.PackageSource", Required: true},
	{Name: "PackageType", Flag: "package-type", Type: "types.PackageType", Required: true},
	{Name: "PackageVendingOptions", Flag: "package-vending-options", Type: "*types.PackageVendingOptions", Required: false},
}

var fields_create_vpc_endpoint = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainArn", Flag: "domain-arn", Type: "*string", Required: true},
	{Name: "VpcOptions", Flag: "vpc-options", Type: "*types.VPCOptions", Required: true},
}

var fields_delete_application = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_data_source = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_direct_query_data_source = []leanruntime.Field{
	{Name: "DataSourceName", Flag: "data-source-name", Type: "*string", Required: true},
}

var fields_delete_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_delete_inbound_connection = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
}

var fields_delete_index = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: true},
}

var fields_delete_outbound_connection = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
}

var fields_delete_package = []leanruntime.Field{
	{Name: "PackageID", Flag: "package-id", Type: "*string", Required: true},
}

var fields_delete_vpc_endpoint = []leanruntime.Field{
	{Name: "VpcEndpointId", Flag: "vpc-endpoint-id", Type: "*string", Required: true},
}

var fields_describe_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_describe_domain_auto_tunes = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_domain_change_progress = []leanruntime.Field{
	{Name: "ChangeId", Flag: "change-id", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_describe_domain_config = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_describe_domain_health = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_describe_domain_nodes = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_describe_domains = []leanruntime.Field{
	{Name: "DomainNames", Flag: "domain-names", Type: "[]string", Required: true},
}

var fields_describe_dry_run_progress = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DryRunId", Flag: "dry-run-id", Type: "*string", Required: false},
	{Name: "LoadDryRunConfig", Flag: "load-dry-run-config", Type: "*bool", Required: false},
}

var fields_describe_inbound_connections = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_instance_type_limits = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: true},
	{Name: "InstanceType", Flag: "instance-type", Type: "types.OpenSearchPartitionInstanceType", Required: true},
}

var fields_describe_outbound_connections = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_packages = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.DescribePackagesFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_reserved_instance_offerings = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReservedInstanceOfferingId", Flag: "reserved-instance-offering-id", Type: "*string", Required: false},
}

var fields_describe_reserved_instances = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReservedInstanceId", Flag: "reserved-instance-id", Type: "*string", Required: false},
}

var fields_describe_vpc_endpoints = []leanruntime.Field{
	{Name: "VpcEndpointIds", Flag: "vpc-endpoint-ids", Type: "[]string", Required: true},
}

var fields_dissociate_package = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "PackageID", Flag: "package-id", Type: "*string", Required: true},
}

var fields_dissociate_packages = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "PackageList", Flag: "package-list", Type: "[]string", Required: true},
}

var fields_get_application = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_compatible_versions = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: false},
}

var fields_get_data_source = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_default_application_setting = []leanruntime.Field{}

var fields_get_direct_query_data_source = []leanruntime.Field{
	{Name: "DataSourceName", Flag: "data-source-name", Type: "*string", Required: true},
}

var fields_get_domain_maintenance_status = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaintenanceId", Flag: "maintenance-id", Type: "*string", Required: true},
}

var fields_get_index = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: true},
}

var fields_get_package_version_history = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PackageID", Flag: "package-id", Type: "*string", Required: true},
}

var fields_get_upgrade_history = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_upgrade_status = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_list_applications = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Statuses", Flag: "statuses", Type: "[]types.ApplicationStatus", Required: false},
}

var fields_list_data_sources = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_list_direct_query_data_sources = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_domain_maintenances = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.MaintenanceType", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.MaintenanceStatus", Required: false},
}

var fields_list_domain_names = []leanruntime.Field{
	{Name: "EngineType", Flag: "engine-type", Type: "types.EngineType", Required: false},
}

var fields_list_domains_for_package = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PackageID", Flag: "package-id", Type: "*string", Required: true},
}

var fields_list_instance_type_details = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: true},
	{Name: "InstanceType", Flag: "instance-type", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RetrieveAZs", Flag: "retrieve-azs", Type: "*bool", Required: false},
}

var fields_list_packages_for_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_scheduled_actions = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags = []leanruntime.Field{
	{Name: "ARN", Flag: "arn", Type: "*string", Required: true},
}

var fields_list_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_vpc_endpoint_access = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_vpc_endpoints = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_vpc_endpoints_for_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_purchase_reserved_instance_offering = []leanruntime.Field{
	{Name: "InstanceCount", Flag: "instance-count", Type: "*int32", Required: false},
	{Name: "ReservationName", Flag: "reservation-name", Type: "*string", Required: true},
	{Name: "ReservedInstanceOfferingId", Flag: "reserved-instance-offering-id", Type: "*string", Required: true},
}

var fields_put_default_application_setting = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
	{Name: "SetAsDefault", Flag: "set-as-default", Type: "*bool", Required: true},
}

var fields_reject_inbound_connection = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
}

var fields_remove_tags = []leanruntime.Field{
	{Name: "ARN", Flag: "arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_revoke_vpc_endpoint_access = []leanruntime.Field{
	{Name: "Account", Flag: "account", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Service", Flag: "service", Type: "types.AWSServicePrincipal", Required: false},
}

var fields_start_domain_maintenance = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.MaintenanceType", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "NodeId", Flag: "node-id", Type: "*string", Required: false},
}

var fields_start_service_software_update = []leanruntime.Field{
	{Name: "DesiredStartTime", Flag: "desired-start-time", Type: "*int64", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "ScheduleAt", Flag: "schedule-at", Type: "types.ScheduleAt", Required: false},
}

var fields_update_application = []leanruntime.Field{
	{Name: "AppConfigs", Flag: "app-configs", Type: "[]types.AppConfig", Required: false},
	{Name: "DataSources", Flag: "data-sources", Type: "[]types.DataSource", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_update_data_source = []leanruntime.Field{
	{Name: "DataSourceType", Flag: "data-source-type", Type: "types.DataSourceType", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.DataSourceStatus", Required: false},
}

var fields_update_direct_query_data_source = []leanruntime.Field{
	{Name: "DataSourceName", Flag: "data-source-name", Type: "*string", Required: true},
	{Name: "DataSourceType", Flag: "data-source-type", Type: "types.DirectQueryDataSourceType", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "OpenSearchArns", Flag: "open-search-arns", Type: "[]string", Required: true},
}

var fields_update_domain_config = []leanruntime.Field{
	{Name: "AIMLOptions", Flag: "aiml-options", Type: "*types.AIMLOptionsInput", Required: false},
	{Name: "AccessPolicies", Flag: "access-policies", Type: "*string", Required: false},
	{Name: "AdvancedOptions", Flag: "advanced-options", Type: "map[string]string", Required: false},
	{Name: "AdvancedSecurityOptions", Flag: "advanced-security-options", Type: "*types.AdvancedSecurityOptionsInput", Required: false},
	{Name: "AutoTuneOptions", Flag: "auto-tune-options", Type: "*types.AutoTuneOptions", Required: false},
	{Name: "ClusterConfig", Flag: "cluster-config", Type: "*types.ClusterConfig", Required: false},
	{Name: "CognitoOptions", Flag: "cognito-options", Type: "*types.CognitoOptions", Required: false},
	{Name: "DomainEndpointOptions", Flag: "domain-endpoint-options", Type: "*types.DomainEndpointOptions", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "DryRunMode", Flag: "dry-run-mode", Type: "types.DryRunMode", Required: false},
	{Name: "EBSOptions", Flag: "ebs-options", Type: "*types.EBSOptions", Required: false},
	{Name: "EncryptionAtRestOptions", Flag: "encryption-at-rest-options", Type: "*types.EncryptionAtRestOptions", Required: false},
	{Name: "IPAddressType", Flag: "ip-address-type", Type: "types.IPAddressType", Required: false},
	{Name: "IdentityCenterOptions", Flag: "identity-center-options", Type: "*types.IdentityCenterOptionsInput", Required: false},
	{Name: "LogPublishingOptions", Flag: "log-publishing-options", Type: "map[string]types.LogPublishingOption", Required: false},
	{Name: "NodeToNodeEncryptionOptions", Flag: "node-to-node-encryption-options", Type: "*types.NodeToNodeEncryptionOptions", Required: false},
	{Name: "OffPeakWindowOptions", Flag: "off-peak-window-options", Type: "*types.OffPeakWindowOptions", Required: false},
	{Name: "SnapshotOptions", Flag: "snapshot-options", Type: "*types.SnapshotOptions", Required: false},
	{Name: "SoftwareUpdateOptions", Flag: "software-update-options", Type: "*types.SoftwareUpdateOptions", Required: false},
	{Name: "VPCOptions", Flag: "vpc-options", Type: "*types.VPCOptions", Required: false},
}

var fields_update_index = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: true},
	{Name: "IndexSchema", Flag: "index-schema", Type: "document.Interface", Required: true},
}

var fields_update_package = []leanruntime.Field{
	{Name: "CommitMessage", Flag: "commit-message", Type: "*string", Required: false},
	{Name: "PackageConfiguration", Flag: "package-configuration", Type: "*types.PackageConfiguration", Required: false},
	{Name: "PackageDescription", Flag: "package-description", Type: "*string", Required: false},
	{Name: "PackageEncryptionOptions", Flag: "package-encryption-options", Type: "*types.PackageEncryptionOptions", Required: false},
	{Name: "PackageID", Flag: "package-id", Type: "*string", Required: true},
	{Name: "PackageSource", Flag: "package-source", Type: "*types.PackageSource", Required: true},
}

var fields_update_package_scope = []leanruntime.Field{
	{Name: "Operation", Flag: "operation", Type: "types.PackageScopeOperationEnum", Required: true},
	{Name: "PackageID", Flag: "package-id", Type: "*string", Required: true},
	{Name: "PackageUserList", Flag: "package-user-list", Type: "[]string", Required: true},
}

var fields_update_scheduled_action = []leanruntime.Field{
	{Name: "ActionID", Flag: "action-id", Type: "*string", Required: true},
	{Name: "ActionType", Flag: "action-type", Type: "types.ActionType", Required: true},
	{Name: "DesiredStartTime", Flag: "desired-start-time", Type: "*int64", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "ScheduleAt", Flag: "schedule-at", Type: "types.ScheduleAt", Required: true},
}

var fields_update_vpc_endpoint = []leanruntime.Field{
	{Name: "VpcEndpointId", Flag: "vpc-endpoint-id", Type: "*string", Required: true},
	{Name: "VpcOptions", Flag: "vpc-options", Type: "*types.VPCOptions", Required: true},
}

var fields_upgrade_domain = []leanruntime.Field{
	{Name: "AdvancedOptions", Flag: "advanced-options", Type: "map[string]string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "PerformCheckOnly", Flag: "perform-check-only", Type: "*bool", Required: false},
	{Name: "TargetVersion", Flag: "target-version", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-inbound-connection": {
			Name:   "accept-inbound-connection",
			Fields: fields_accept_inbound_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptInboundConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_inbound_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptInboundConnection(ctx, input)
			},
		},
		"add-data-source": {
			Name:   "add-data-source",
			Fields: fields_add_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddDataSource(ctx, input)
			},
		},
		"add-direct-query-data-source": {
			Name:   "add-direct-query-data-source",
			Fields: fields_add_direct_query_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddDirectQueryDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_direct_query_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddDirectQueryDataSource(ctx, input)
			},
		},
		"add-tags": {
			Name:   "add-tags",
			Fields: fields_add_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddTags(ctx, input)
			},
		},
		"associate-package": {
			Name:   "associate-package",
			Fields: fields_associate_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociatePackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociatePackage(ctx, input)
			},
		},
		"associate-packages": {
			Name:   "associate-packages",
			Fields: fields_associate_packages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociatePackagesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_packages, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociatePackages(ctx, input)
			},
		},
		"authorize-vpc-endpoint-access": {
			Name:   "authorize-vpc-endpoint-access",
			Fields: fields_authorize_vpc_endpoint_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AuthorizeVpcEndpointAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_authorize_vpc_endpoint_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AuthorizeVpcEndpointAccess(ctx, input)
			},
		},
		"cancel-domain-config-change": {
			Name:   "cancel-domain-config-change",
			Fields: fields_cancel_domain_config_change,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelDomainConfigChangeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_domain_config_change, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelDomainConfigChange(ctx, input)
			},
		},
		"cancel-service-software-update": {
			Name:   "cancel-service-software-update",
			Fields: fields_cancel_service_software_update,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelServiceSoftwareUpdateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_service_software_update, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelServiceSoftwareUpdate(ctx, input)
			},
		},
		"create-application": {
			Name:   "create-application",
			Fields: fields_create_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApplication(ctx, input)
			},
		},
		"create-domain": {
			Name:   "create-domain",
			Fields: fields_create_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDomain(ctx, input)
			},
		},
		"create-index": {
			Name:   "create-index",
			Fields: fields_create_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIndex(ctx, input)
			},
		},
		"create-outbound-connection": {
			Name:   "create-outbound-connection",
			Fields: fields_create_outbound_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOutboundConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_outbound_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOutboundConnection(ctx, input)
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
		"create-vpc-endpoint": {
			Name:   "create-vpc-endpoint",
			Fields: fields_create_vpc_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVpcEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpc_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVpcEndpoint(ctx, input)
			},
		},
		"delete-application": {
			Name:   "delete-application",
			Fields: fields_delete_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApplication(ctx, input)
			},
		},
		"delete-data-source": {
			Name:   "delete-data-source",
			Fields: fields_delete_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataSource(ctx, input)
			},
		},
		"delete-direct-query-data-source": {
			Name:   "delete-direct-query-data-source",
			Fields: fields_delete_direct_query_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDirectQueryDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_direct_query_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDirectQueryDataSource(ctx, input)
			},
		},
		"delete-domain": {
			Name:   "delete-domain",
			Fields: fields_delete_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDomain(ctx, input)
			},
		},
		"delete-inbound-connection": {
			Name:   "delete-inbound-connection",
			Fields: fields_delete_inbound_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInboundConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_inbound_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInboundConnection(ctx, input)
			},
		},
		"delete-index": {
			Name:   "delete-index",
			Fields: fields_delete_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIndex(ctx, input)
			},
		},
		"delete-outbound-connection": {
			Name:   "delete-outbound-connection",
			Fields: fields_delete_outbound_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOutboundConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_outbound_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOutboundConnection(ctx, input)
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
		"delete-vpc-endpoint": {
			Name:   "delete-vpc-endpoint",
			Fields: fields_delete_vpc_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVpcEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpc_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVpcEndpoint(ctx, input)
			},
		},
		"describe-domain": {
			Name:   "describe-domain",
			Fields: fields_describe_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDomain(ctx, input)
			},
		},
		"describe-domain-auto-tunes": {
			Name:   "describe-domain-auto-tunes",
			Fields: fields_describe_domain_auto_tunes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDomainAutoTunesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_domain_auto_tunes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDomainAutoTunes(ctx, input)
				}
				var results []*svc.DescribeDomainAutoTunesOutput
				p := svc.NewDescribeDomainAutoTunesPaginator(client, input)
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
		"describe-domain-change-progress": {
			Name:   "describe-domain-change-progress",
			Fields: fields_describe_domain_change_progress,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDomainChangeProgressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_domain_change_progress, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDomainChangeProgress(ctx, input)
			},
		},
		"describe-domain-config": {
			Name:   "describe-domain-config",
			Fields: fields_describe_domain_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDomainConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_domain_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDomainConfig(ctx, input)
			},
		},
		"describe-domain-health": {
			Name:   "describe-domain-health",
			Fields: fields_describe_domain_health,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDomainHealthInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_domain_health, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDomainHealth(ctx, input)
			},
		},
		"describe-domain-nodes": {
			Name:   "describe-domain-nodes",
			Fields: fields_describe_domain_nodes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDomainNodesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_domain_nodes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDomainNodes(ctx, input)
			},
		},
		"describe-domains": {
			Name:   "describe-domains",
			Fields: fields_describe_domains,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDomainsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_domains, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDomains(ctx, input)
			},
		},
		"describe-dry-run-progress": {
			Name:   "describe-dry-run-progress",
			Fields: fields_describe_dry_run_progress,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDryRunProgressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_dry_run_progress, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDryRunProgress(ctx, input)
			},
		},
		"describe-inbound-connections": {
			Name:   "describe-inbound-connections",
			Fields: fields_describe_inbound_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInboundConnectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_inbound_connections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeInboundConnections(ctx, input)
				}
				var results []*svc.DescribeInboundConnectionsOutput
				p := svc.NewDescribeInboundConnectionsPaginator(client, input)
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
		"describe-instance-type-limits": {
			Name:   "describe-instance-type-limits",
			Fields: fields_describe_instance_type_limits,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstanceTypeLimitsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_instance_type_limits, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInstanceTypeLimits(ctx, input)
			},
		},
		"describe-outbound-connections": {
			Name:   "describe-outbound-connections",
			Fields: fields_describe_outbound_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOutboundConnectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_outbound_connections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeOutboundConnections(ctx, input)
				}
				var results []*svc.DescribeOutboundConnectionsOutput
				p := svc.NewDescribeOutboundConnectionsPaginator(client, input)
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
		"describe-packages": {
			Name:   "describe-packages",
			Fields: fields_describe_packages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePackagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_packages, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribePackages(ctx, input)
				}
				var results []*svc.DescribePackagesOutput
				p := svc.NewDescribePackagesPaginator(client, input)
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
		"describe-reserved-instance-offerings": {
			Name:   "describe-reserved-instance-offerings",
			Fields: fields_describe_reserved_instance_offerings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReservedInstanceOfferingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_reserved_instance_offerings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReservedInstanceOfferings(ctx, input)
				}
				var results []*svc.DescribeReservedInstanceOfferingsOutput
				p := svc.NewDescribeReservedInstanceOfferingsPaginator(client, input)
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
		"describe-reserved-instances": {
			Name:   "describe-reserved-instances",
			Fields: fields_describe_reserved_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReservedInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_reserved_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReservedInstances(ctx, input)
				}
				var results []*svc.DescribeReservedInstancesOutput
				p := svc.NewDescribeReservedInstancesPaginator(client, input)
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
		"describe-vpc-endpoints": {
			Name:   "describe-vpc-endpoints",
			Fields: fields_describe_vpc_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpcEndpointsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_vpc_endpoints, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVpcEndpoints(ctx, input)
			},
		},
		"dissociate-package": {
			Name:   "dissociate-package",
			Fields: fields_dissociate_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DissociatePackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_dissociate_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DissociatePackage(ctx, input)
			},
		},
		"dissociate-packages": {
			Name:   "dissociate-packages",
			Fields: fields_dissociate_packages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DissociatePackagesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_dissociate_packages, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DissociatePackages(ctx, input)
			},
		},
		"get-application": {
			Name:   "get-application",
			Fields: fields_get_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApplication(ctx, input)
			},
		},
		"get-compatible-versions": {
			Name:   "get-compatible-versions",
			Fields: fields_get_compatible_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCompatibleVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_compatible_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCompatibleVersions(ctx, input)
			},
		},
		"get-data-source": {
			Name:   "get-data-source",
			Fields: fields_get_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataSource(ctx, input)
			},
		},
		"get-default-application-setting": {
			Name:   "get-default-application-setting",
			Fields: fields_get_default_application_setting,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDefaultApplicationSettingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_default_application_setting, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDefaultApplicationSetting(ctx, input)
			},
		},
		"get-direct-query-data-source": {
			Name:   "get-direct-query-data-source",
			Fields: fields_get_direct_query_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDirectQueryDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_direct_query_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDirectQueryDataSource(ctx, input)
			},
		},
		"get-domain-maintenance-status": {
			Name:   "get-domain-maintenance-status",
			Fields: fields_get_domain_maintenance_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDomainMaintenanceStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_domain_maintenance_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDomainMaintenanceStatus(ctx, input)
			},
		},
		"get-index": {
			Name:   "get-index",
			Fields: fields_get_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIndex(ctx, input)
			},
		},
		"get-package-version-history": {
			Name:   "get-package-version-history",
			Fields: fields_get_package_version_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPackageVersionHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_package_version_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetPackageVersionHistory(ctx, input)
				}
				var results []*svc.GetPackageVersionHistoryOutput
				p := svc.NewGetPackageVersionHistoryPaginator(client, input)
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
		"get-upgrade-history": {
			Name:   "get-upgrade-history",
			Fields: fields_get_upgrade_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUpgradeHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_upgrade_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetUpgradeHistory(ctx, input)
				}
				var results []*svc.GetUpgradeHistoryOutput
				p := svc.NewGetUpgradeHistoryPaginator(client, input)
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
		"get-upgrade-status": {
			Name:   "get-upgrade-status",
			Fields: fields_get_upgrade_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUpgradeStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_upgrade_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUpgradeStatus(ctx, input)
			},
		},
		"list-applications": {
			Name:   "list-applications",
			Fields: fields_list_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_applications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplications(ctx, input)
				}
				var results []*svc.ListApplicationsOutput
				p := svc.NewListApplicationsPaginator(client, input)
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
		"list-data-sources": {
			Name:   "list-data-sources",
			Fields: fields_list_data_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataSourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_data_sources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDataSources(ctx, input)
			},
		},
		"list-direct-query-data-sources": {
			Name:   "list-direct-query-data-sources",
			Fields: fields_list_direct_query_data_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDirectQueryDataSourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_direct_query_data_sources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDirectQueryDataSources(ctx, input)
			},
		},
		"list-domain-maintenances": {
			Name:   "list-domain-maintenances",
			Fields: fields_list_domain_maintenances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDomainMaintenancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_domain_maintenances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDomainMaintenances(ctx, input)
				}
				var results []*svc.ListDomainMaintenancesOutput
				p := svc.NewListDomainMaintenancesPaginator(client, input)
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
		"list-domain-names": {
			Name:   "list-domain-names",
			Fields: fields_list_domain_names,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDomainNamesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_domain_names, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDomainNames(ctx, input)
			},
		},
		"list-domains-for-package": {
			Name:   "list-domains-for-package",
			Fields: fields_list_domains_for_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDomainsForPackageInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_domains_for_package, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDomainsForPackage(ctx, input)
				}
				var results []*svc.ListDomainsForPackageOutput
				p := svc.NewListDomainsForPackagePaginator(client, input)
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
		"list-instance-type-details": {
			Name:   "list-instance-type-details",
			Fields: fields_list_instance_type_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInstanceTypeDetailsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_instance_type_details, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInstanceTypeDetails(ctx, input)
				}
				var results []*svc.ListInstanceTypeDetailsOutput
				p := svc.NewListInstanceTypeDetailsPaginator(client, input)
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
		"list-packages-for-domain": {
			Name:   "list-packages-for-domain",
			Fields: fields_list_packages_for_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPackagesForDomainInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_packages_for_domain, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPackagesForDomain(ctx, input)
				}
				var results []*svc.ListPackagesForDomainOutput
				p := svc.NewListPackagesForDomainPaginator(client, input)
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
		"list-scheduled-actions": {
			Name:   "list-scheduled-actions",
			Fields: fields_list_scheduled_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListScheduledActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_scheduled_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListScheduledActions(ctx, input)
				}
				var results []*svc.ListScheduledActionsOutput
				p := svc.NewListScheduledActionsPaginator(client, input)
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
		"list-tags": {
			Name:   "list-tags",
			Fields: fields_list_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTags(ctx, input)
			},
		},
		"list-versions": {
			Name:   "list-versions",
			Fields: fields_list_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVersions(ctx, input)
				}
				var results []*svc.ListVersionsOutput
				p := svc.NewListVersionsPaginator(client, input)
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
		"list-vpc-endpoint-access": {
			Name:   "list-vpc-endpoint-access",
			Fields: fields_list_vpc_endpoint_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVpcEndpointAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_vpc_endpoint_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListVpcEndpointAccess(ctx, input)
			},
		},
		"list-vpc-endpoints": {
			Name:   "list-vpc-endpoints",
			Fields: fields_list_vpc_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVpcEndpointsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_vpc_endpoints, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListVpcEndpoints(ctx, input)
			},
		},
		"list-vpc-endpoints-for-domain": {
			Name:   "list-vpc-endpoints-for-domain",
			Fields: fields_list_vpc_endpoints_for_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVpcEndpointsForDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_vpc_endpoints_for_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListVpcEndpointsForDomain(ctx, input)
			},
		},
		"purchase-reserved-instance-offering": {
			Name:   "purchase-reserved-instance-offering",
			Fields: fields_purchase_reserved_instance_offering,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PurchaseReservedInstanceOfferingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_purchase_reserved_instance_offering, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PurchaseReservedInstanceOffering(ctx, input)
			},
		},
		"put-default-application-setting": {
			Name:   "put-default-application-setting",
			Fields: fields_put_default_application_setting,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDefaultApplicationSettingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_default_application_setting, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDefaultApplicationSetting(ctx, input)
			},
		},
		"reject-inbound-connection": {
			Name:   "reject-inbound-connection",
			Fields: fields_reject_inbound_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectInboundConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_inbound_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectInboundConnection(ctx, input)
			},
		},
		"remove-tags": {
			Name:   "remove-tags",
			Fields: fields_remove_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveTags(ctx, input)
			},
		},
		"revoke-vpc-endpoint-access": {
			Name:   "revoke-vpc-endpoint-access",
			Fields: fields_revoke_vpc_endpoint_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RevokeVpcEndpointAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_revoke_vpc_endpoint_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RevokeVpcEndpointAccess(ctx, input)
			},
		},
		"start-domain-maintenance": {
			Name:   "start-domain-maintenance",
			Fields: fields_start_domain_maintenance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDomainMaintenanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_domain_maintenance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDomainMaintenance(ctx, input)
			},
		},
		"start-service-software-update": {
			Name:   "start-service-software-update",
			Fields: fields_start_service_software_update,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartServiceSoftwareUpdateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_service_software_update, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartServiceSoftwareUpdate(ctx, input)
			},
		},
		"update-application": {
			Name:   "update-application",
			Fields: fields_update_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApplication(ctx, input)
			},
		},
		"update-data-source": {
			Name:   "update-data-source",
			Fields: fields_update_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataSource(ctx, input)
			},
		},
		"update-direct-query-data-source": {
			Name:   "update-direct-query-data-source",
			Fields: fields_update_direct_query_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDirectQueryDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_direct_query_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDirectQueryDataSource(ctx, input)
			},
		},
		"update-domain-config": {
			Name:   "update-domain-config",
			Fields: fields_update_domain_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDomainConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_domain_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDomainConfig(ctx, input)
			},
		},
		"update-index": {
			Name:   "update-index",
			Fields: fields_update_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIndex(ctx, input)
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
		"update-package-scope": {
			Name:   "update-package-scope",
			Fields: fields_update_package_scope,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePackageScopeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_package_scope, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePackageScope(ctx, input)
			},
		},
		"update-scheduled-action": {
			Name:   "update-scheduled-action",
			Fields: fields_update_scheduled_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateScheduledActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_scheduled_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateScheduledAction(ctx, input)
			},
		},
		"update-vpc-endpoint": {
			Name:   "update-vpc-endpoint",
			Fields: fields_update_vpc_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVpcEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_vpc_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVpcEndpoint(ctx, input)
			},
		},
		"upgrade-domain": {
			Name:   "upgrade-domain",
			Fields: fields_upgrade_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpgradeDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_upgrade_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpgradeDomain(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("opensearch", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
