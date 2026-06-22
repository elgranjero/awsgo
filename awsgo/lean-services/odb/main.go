package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/odb"
)

var fields_accept_marketplace_registration = []leanruntime.Field{
	{Name: "MarketplaceRegistrationToken", Flag: "marketplace-registration-token", Type: "*string", Required: true},
}

var fields_associate_iam_role_to_resource = []leanruntime.Field{
	{Name: "AwsIntegration", Flag: "aws-integration", Type: "types.SupportedAwsIntegration", Required: true},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_create_cloud_autonomous_vm_cluster = []leanruntime.Field{
	{Name: "AutonomousDataStorageSizeInTBs", Flag: "autonomous-data-storage-size-in-tbs", Type: "*float64", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CloudExadataInfrastructureId", Flag: "cloud-exadata-infrastructure-id", Type: "*string", Required: true},
	{Name: "CpuCoreCountPerNode", Flag: "cpu-core-count-per-node", Type: "*int32", Required: true},
	{Name: "DbServers", Flag: "db-servers", Type: "[]string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "IsMtlsEnabledVmCluster", Flag: "is-mtls-enabled-vm-cluster", Type: "*bool", Required: false},
	{Name: "LicenseModel", Flag: "license-model", Type: "types.LicenseModel", Required: false},
	{Name: "MaintenanceWindow", Flag: "maintenance-window", Type: "*types.MaintenanceWindow", Required: false},
	{Name: "MemoryPerOracleComputeUnitInGBs", Flag: "memory-per-oracle-compute-unit-in-gbs", Type: "*int32", Required: true},
	{Name: "OdbNetworkId", Flag: "odb-network-id", Type: "*string", Required: true},
	{Name: "ScanListenerPortNonTls", Flag: "scan-listener-port-non-tls", Type: "*int32", Required: false},
	{Name: "ScanListenerPortTls", Flag: "scan-listener-port-tls", Type: "*int32", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TimeZone", Flag: "time-zone", Type: "*string", Required: false},
	{Name: "TotalContainerDatabases", Flag: "total-container-databases", Type: "*int32", Required: true},
}

var fields_create_cloud_exadata_infrastructure = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "AvailabilityZoneId", Flag: "availability-zone-id", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ComputeCount", Flag: "compute-count", Type: "*int32", Required: true},
	{Name: "CustomerContactsToSendToOCI", Flag: "customer-contacts-to-send-to-oci", Type: "[]types.CustomerContact", Required: false},
	{Name: "DatabaseServerType", Flag: "database-server-type", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "MaintenanceWindow", Flag: "maintenance-window", Type: "*types.MaintenanceWindow", Required: false},
	{Name: "Shape", Flag: "shape", Type: "*string", Required: true},
	{Name: "StorageCount", Flag: "storage-count", Type: "*int32", Required: true},
	{Name: "StorageServerType", Flag: "storage-server-type", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_cloud_vm_cluster = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CloudExadataInfrastructureId", Flag: "cloud-exadata-infrastructure-id", Type: "*string", Required: true},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: false},
	{Name: "CpuCoreCount", Flag: "cpu-core-count", Type: "*int32", Required: true},
	{Name: "DataCollectionOptions", Flag: "data-collection-options", Type: "*types.DataCollectionOptions", Required: false},
	{Name: "DataStorageSizeInTBs", Flag: "data-storage-size-in-tbs", Type: "*float64", Required: false},
	{Name: "DbNodeStorageSizeInGBs", Flag: "db-node-storage-size-in-gbs", Type: "*int32", Required: false},
	{Name: "DbServers", Flag: "db-servers", Type: "[]string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "GiVersion", Flag: "gi-version", Type: "*string", Required: true},
	{Name: "Hostname", Flag: "hostname", Type: "*string", Required: true},
	{Name: "IsLocalBackupEnabled", Flag: "is-local-backup-enabled", Type: "*bool", Required: false},
	{Name: "IsSparseDiskgroupEnabled", Flag: "is-sparse-diskgroup-enabled", Type: "*bool", Required: false},
	{Name: "LicenseModel", Flag: "license-model", Type: "types.LicenseModel", Required: false},
	{Name: "MemorySizeInGBs", Flag: "memory-size-in-gbs", Type: "*int32", Required: false},
	{Name: "OdbNetworkId", Flag: "odb-network-id", Type: "*string", Required: true},
	{Name: "ScanListenerPortTcp", Flag: "scan-listener-port-tcp", Type: "*int32", Required: false},
	{Name: "SshPublicKeys", Flag: "ssh-public-keys", Type: "[]string", Required: true},
	{Name: "SystemVersion", Flag: "system-version", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TimeZone", Flag: "time-zone", Type: "*string", Required: false},
}

var fields_create_odb_network = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "AvailabilityZoneId", Flag: "availability-zone-id", Type: "*string", Required: false},
	{Name: "BackupSubnetCidr", Flag: "backup-subnet-cidr", Type: "*string", Required: false},
	{Name: "ClientSubnetCidr", Flag: "client-subnet-cidr", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CrossRegionS3RestoreSourcesToEnable", Flag: "cross-region-s3-restore-sources-to-enable", Type: "[]string", Required: false},
	{Name: "CustomDomainName", Flag: "custom-domain-name", Type: "*string", Required: false},
	{Name: "DefaultDnsPrefix", Flag: "default-dns-prefix", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "KmsAccess", Flag: "kms-access", Type: "types.Access", Required: false},
	{Name: "KmsPolicyDocument", Flag: "kms-policy-document", Type: "*string", Required: false},
	{Name: "S3Access", Flag: "s3-access", Type: "types.Access", Required: false},
	{Name: "S3PolicyDocument", Flag: "s3-policy-document", Type: "*string", Required: false},
	{Name: "StsAccess", Flag: "sts-access", Type: "types.Access", Required: false},
	{Name: "StsPolicyDocument", Flag: "sts-policy-document", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "ZeroEtlAccess", Flag: "zero-etl-access", Type: "types.Access", Required: false},
}

var fields_create_odb_peering_connection = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "OdbNetworkId", Flag: "odb-network-id", Type: "*string", Required: true},
	{Name: "PeerNetworkCidrsToBeAdded", Flag: "peer-network-cidrs-to-be-added", Type: "[]string", Required: false},
	{Name: "PeerNetworkId", Flag: "peer-network-id", Type: "*string", Required: true},
	{Name: "PeerNetworkRouteTableIds", Flag: "peer-network-route-table-ids", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_cloud_autonomous_vm_cluster = []leanruntime.Field{
	{Name: "CloudAutonomousVmClusterId", Flag: "cloud-autonomous-vm-cluster-id", Type: "*string", Required: true},
}

var fields_delete_cloud_exadata_infrastructure = []leanruntime.Field{
	{Name: "CloudExadataInfrastructureId", Flag: "cloud-exadata-infrastructure-id", Type: "*string", Required: true},
}

var fields_delete_cloud_vm_cluster = []leanruntime.Field{
	{Name: "CloudVmClusterId", Flag: "cloud-vm-cluster-id", Type: "*string", Required: true},
}

var fields_delete_odb_network = []leanruntime.Field{
	{Name: "DeleteAssociatedResources", Flag: "delete-associated-resources", Type: "*bool", Required: true},
	{Name: "OdbNetworkId", Flag: "odb-network-id", Type: "*string", Required: true},
}

var fields_delete_odb_peering_connection = []leanruntime.Field{
	{Name: "OdbPeeringConnectionId", Flag: "odb-peering-connection-id", Type: "*string", Required: true},
}

var fields_disassociate_iam_role_from_resource = []leanruntime.Field{
	{Name: "AwsIntegration", Flag: "aws-integration", Type: "types.SupportedAwsIntegration", Required: true},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_cloud_autonomous_vm_cluster = []leanruntime.Field{
	{Name: "CloudAutonomousVmClusterId", Flag: "cloud-autonomous-vm-cluster-id", Type: "*string", Required: true},
}

var fields_get_cloud_exadata_infrastructure = []leanruntime.Field{
	{Name: "CloudExadataInfrastructureId", Flag: "cloud-exadata-infrastructure-id", Type: "*string", Required: true},
}

var fields_get_cloud_exadata_infrastructure_unallocated_resources = []leanruntime.Field{
	{Name: "CloudExadataInfrastructureId", Flag: "cloud-exadata-infrastructure-id", Type: "*string", Required: true},
	{Name: "DbServers", Flag: "db-servers", Type: "[]string", Required: false},
}

var fields_get_cloud_vm_cluster = []leanruntime.Field{
	{Name: "CloudVmClusterId", Flag: "cloud-vm-cluster-id", Type: "*string", Required: true},
}

var fields_get_db_node = []leanruntime.Field{
	{Name: "CloudVmClusterId", Flag: "cloud-vm-cluster-id", Type: "*string", Required: true},
	{Name: "DbNodeId", Flag: "db-node-id", Type: "*string", Required: true},
}

var fields_get_db_server = []leanruntime.Field{
	{Name: "CloudExadataInfrastructureId", Flag: "cloud-exadata-infrastructure-id", Type: "*string", Required: true},
	{Name: "DbServerId", Flag: "db-server-id", Type: "*string", Required: true},
}

var fields_get_oci_onboarding_status = []leanruntime.Field{}

var fields_get_odb_network = []leanruntime.Field{
	{Name: "OdbNetworkId", Flag: "odb-network-id", Type: "*string", Required: true},
}

var fields_get_odb_peering_connection = []leanruntime.Field{
	{Name: "OdbPeeringConnectionId", Flag: "odb-peering-connection-id", Type: "*string", Required: true},
}

var fields_initialize_service = []leanruntime.Field{
	{Name: "OciIdentityDomain", Flag: "oci-identity-domain", Type: "*bool", Required: false},
}

var fields_list_autonomous_virtual_machines = []leanruntime.Field{
	{Name: "CloudAutonomousVmClusterId", Flag: "cloud-autonomous-vm-cluster-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_cloud_autonomous_vm_clusters = []leanruntime.Field{
	{Name: "CloudExadataInfrastructureId", Flag: "cloud-exadata-infrastructure-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_cloud_exadata_infrastructures = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_cloud_vm_clusters = []leanruntime.Field{
	{Name: "CloudExadataInfrastructureId", Flag: "cloud-exadata-infrastructure-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_db_nodes = []leanruntime.Field{
	{Name: "CloudVmClusterId", Flag: "cloud-vm-cluster-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_db_servers = []leanruntime.Field{
	{Name: "CloudExadataInfrastructureId", Flag: "cloud-exadata-infrastructure-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_db_system_shapes = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "AvailabilityZoneId", Flag: "availability-zone-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_gi_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Shape", Flag: "shape", Type: "*string", Required: false},
}

var fields_list_odb_networks = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_odb_peering_connections = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OdbNetworkId", Flag: "odb-network-id", Type: "*string", Required: false},
}

var fields_list_system_versions = []leanruntime.Field{
	{Name: "GiVersion", Flag: "gi-version", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Shape", Flag: "shape", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_reboot_db_node = []leanruntime.Field{
	{Name: "CloudVmClusterId", Flag: "cloud-vm-cluster-id", Type: "*string", Required: true},
	{Name: "DbNodeId", Flag: "db-node-id", Type: "*string", Required: true},
}

var fields_start_db_node = []leanruntime.Field{
	{Name: "CloudVmClusterId", Flag: "cloud-vm-cluster-id", Type: "*string", Required: true},
	{Name: "DbNodeId", Flag: "db-node-id", Type: "*string", Required: true},
}

var fields_stop_db_node = []leanruntime.Field{
	{Name: "CloudVmClusterId", Flag: "cloud-vm-cluster-id", Type: "*string", Required: true},
	{Name: "DbNodeId", Flag: "db-node-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_cloud_exadata_infrastructure = []leanruntime.Field{
	{Name: "CloudExadataInfrastructureId", Flag: "cloud-exadata-infrastructure-id", Type: "*string", Required: true},
	{Name: "MaintenanceWindow", Flag: "maintenance-window", Type: "*types.MaintenanceWindow", Required: false},
}

var fields_update_odb_network = []leanruntime.Field{
	{Name: "CrossRegionS3RestoreSourcesToDisable", Flag: "cross-region-s3-restore-sources-to-disable", Type: "[]string", Required: false},
	{Name: "CrossRegionS3RestoreSourcesToEnable", Flag: "cross-region-s3-restore-sources-to-enable", Type: "[]string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "KmsAccess", Flag: "kms-access", Type: "types.Access", Required: false},
	{Name: "KmsPolicyDocument", Flag: "kms-policy-document", Type: "*string", Required: false},
	{Name: "OdbNetworkId", Flag: "odb-network-id", Type: "*string", Required: true},
	{Name: "PeeredCidrsToBeAdded", Flag: "peered-cidrs-to-be-added", Type: "[]string", Required: false},
	{Name: "PeeredCidrsToBeRemoved", Flag: "peered-cidrs-to-be-removed", Type: "[]string", Required: false},
	{Name: "S3Access", Flag: "s3-access", Type: "types.Access", Required: false},
	{Name: "S3PolicyDocument", Flag: "s3-policy-document", Type: "*string", Required: false},
	{Name: "StsAccess", Flag: "sts-access", Type: "types.Access", Required: false},
	{Name: "StsPolicyDocument", Flag: "sts-policy-document", Type: "*string", Required: false},
	{Name: "ZeroEtlAccess", Flag: "zero-etl-access", Type: "types.Access", Required: false},
}

var fields_update_odb_peering_connection = []leanruntime.Field{
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "OdbPeeringConnectionId", Flag: "odb-peering-connection-id", Type: "*string", Required: true},
	{Name: "PeerNetworkCidrsToBeAdded", Flag: "peer-network-cidrs-to-be-added", Type: "[]string", Required: false},
	{Name: "PeerNetworkCidrsToBeRemoved", Flag: "peer-network-cidrs-to-be-removed", Type: "[]string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-marketplace-registration": {
			Name:   "accept-marketplace-registration",
			Fields: fields_accept_marketplace_registration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptMarketplaceRegistrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_marketplace_registration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptMarketplaceRegistration(ctx, input)
			},
		},
		"associate-iam-role-to-resource": {
			Name:   "associate-iam-role-to-resource",
			Fields: fields_associate_iam_role_to_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateIamRoleToResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_iam_role_to_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateIamRoleToResource(ctx, input)
			},
		},
		"create-cloud-autonomous-vm-cluster": {
			Name:   "create-cloud-autonomous-vm-cluster",
			Fields: fields_create_cloud_autonomous_vm_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCloudAutonomousVmClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cloud_autonomous_vm_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCloudAutonomousVmCluster(ctx, input)
			},
		},
		"create-cloud-exadata-infrastructure": {
			Name:   "create-cloud-exadata-infrastructure",
			Fields: fields_create_cloud_exadata_infrastructure,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCloudExadataInfrastructureInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cloud_exadata_infrastructure, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCloudExadataInfrastructure(ctx, input)
			},
		},
		"create-cloud-vm-cluster": {
			Name:   "create-cloud-vm-cluster",
			Fields: fields_create_cloud_vm_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCloudVmClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cloud_vm_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCloudVmCluster(ctx, input)
			},
		},
		"create-odb-network": {
			Name:   "create-odb-network",
			Fields: fields_create_odb_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOdbNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_odb_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOdbNetwork(ctx, input)
			},
		},
		"create-odb-peering-connection": {
			Name:   "create-odb-peering-connection",
			Fields: fields_create_odb_peering_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOdbPeeringConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_odb_peering_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOdbPeeringConnection(ctx, input)
			},
		},
		"delete-cloud-autonomous-vm-cluster": {
			Name:   "delete-cloud-autonomous-vm-cluster",
			Fields: fields_delete_cloud_autonomous_vm_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCloudAutonomousVmClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cloud_autonomous_vm_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCloudAutonomousVmCluster(ctx, input)
			},
		},
		"delete-cloud-exadata-infrastructure": {
			Name:   "delete-cloud-exadata-infrastructure",
			Fields: fields_delete_cloud_exadata_infrastructure,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCloudExadataInfrastructureInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cloud_exadata_infrastructure, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCloudExadataInfrastructure(ctx, input)
			},
		},
		"delete-cloud-vm-cluster": {
			Name:   "delete-cloud-vm-cluster",
			Fields: fields_delete_cloud_vm_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCloudVmClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cloud_vm_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCloudVmCluster(ctx, input)
			},
		},
		"delete-odb-network": {
			Name:   "delete-odb-network",
			Fields: fields_delete_odb_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOdbNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_odb_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOdbNetwork(ctx, input)
			},
		},
		"delete-odb-peering-connection": {
			Name:   "delete-odb-peering-connection",
			Fields: fields_delete_odb_peering_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOdbPeeringConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_odb_peering_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOdbPeeringConnection(ctx, input)
			},
		},
		"disassociate-iam-role-from-resource": {
			Name:   "disassociate-iam-role-from-resource",
			Fields: fields_disassociate_iam_role_from_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateIamRoleFromResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_iam_role_from_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateIamRoleFromResource(ctx, input)
			},
		},
		"get-cloud-autonomous-vm-cluster": {
			Name:   "get-cloud-autonomous-vm-cluster",
			Fields: fields_get_cloud_autonomous_vm_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCloudAutonomousVmClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cloud_autonomous_vm_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCloudAutonomousVmCluster(ctx, input)
			},
		},
		"get-cloud-exadata-infrastructure": {
			Name:   "get-cloud-exadata-infrastructure",
			Fields: fields_get_cloud_exadata_infrastructure,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCloudExadataInfrastructureInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cloud_exadata_infrastructure, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCloudExadataInfrastructure(ctx, input)
			},
		},
		"get-cloud-exadata-infrastructure-unallocated-resources": {
			Name:   "get-cloud-exadata-infrastructure-unallocated-resources",
			Fields: fields_get_cloud_exadata_infrastructure_unallocated_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCloudExadataInfrastructureUnallocatedResourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cloud_exadata_infrastructure_unallocated_resources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCloudExadataInfrastructureUnallocatedResources(ctx, input)
			},
		},
		"get-cloud-vm-cluster": {
			Name:   "get-cloud-vm-cluster",
			Fields: fields_get_cloud_vm_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCloudVmClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cloud_vm_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCloudVmCluster(ctx, input)
			},
		},
		"get-db-node": {
			Name:   "get-db-node",
			Fields: fields_get_db_node,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDbNodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_db_node, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDbNode(ctx, input)
			},
		},
		"get-db-server": {
			Name:   "get-db-server",
			Fields: fields_get_db_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDbServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_db_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDbServer(ctx, input)
			},
		},
		"get-oci-onboarding-status": {
			Name:   "get-oci-onboarding-status",
			Fields: fields_get_oci_onboarding_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOciOnboardingStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_oci_onboarding_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOciOnboardingStatus(ctx, input)
			},
		},
		"get-odb-network": {
			Name:   "get-odb-network",
			Fields: fields_get_odb_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOdbNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_odb_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOdbNetwork(ctx, input)
			},
		},
		"get-odb-peering-connection": {
			Name:   "get-odb-peering-connection",
			Fields: fields_get_odb_peering_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOdbPeeringConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_odb_peering_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOdbPeeringConnection(ctx, input)
			},
		},
		"initialize-service": {
			Name:   "initialize-service",
			Fields: fields_initialize_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InitializeServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_initialize_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InitializeService(ctx, input)
			},
		},
		"list-autonomous-virtual-machines": {
			Name:   "list-autonomous-virtual-machines",
			Fields: fields_list_autonomous_virtual_machines,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAutonomousVirtualMachinesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_autonomous_virtual_machines, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAutonomousVirtualMachines(ctx, input)
				}
				var results []*svc.ListAutonomousVirtualMachinesOutput
				p := svc.NewListAutonomousVirtualMachinesPaginator(client, input)
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
		"list-cloud-autonomous-vm-clusters": {
			Name:   "list-cloud-autonomous-vm-clusters",
			Fields: fields_list_cloud_autonomous_vm_clusters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCloudAutonomousVmClustersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cloud_autonomous_vm_clusters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCloudAutonomousVmClusters(ctx, input)
				}
				var results []*svc.ListCloudAutonomousVmClustersOutput
				p := svc.NewListCloudAutonomousVmClustersPaginator(client, input)
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
		"list-cloud-exadata-infrastructures": {
			Name:   "list-cloud-exadata-infrastructures",
			Fields: fields_list_cloud_exadata_infrastructures,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCloudExadataInfrastructuresInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cloud_exadata_infrastructures, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCloudExadataInfrastructures(ctx, input)
				}
				var results []*svc.ListCloudExadataInfrastructuresOutput
				p := svc.NewListCloudExadataInfrastructuresPaginator(client, input)
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
		"list-cloud-vm-clusters": {
			Name:   "list-cloud-vm-clusters",
			Fields: fields_list_cloud_vm_clusters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCloudVmClustersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cloud_vm_clusters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCloudVmClusters(ctx, input)
				}
				var results []*svc.ListCloudVmClustersOutput
				p := svc.NewListCloudVmClustersPaginator(client, input)
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
		"list-db-nodes": {
			Name:   "list-db-nodes",
			Fields: fields_list_db_nodes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDbNodesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_db_nodes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDbNodes(ctx, input)
				}
				var results []*svc.ListDbNodesOutput
				p := svc.NewListDbNodesPaginator(client, input)
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
		"list-db-servers": {
			Name:   "list-db-servers",
			Fields: fields_list_db_servers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDbServersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_db_servers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDbServers(ctx, input)
				}
				var results []*svc.ListDbServersOutput
				p := svc.NewListDbServersPaginator(client, input)
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
		"list-db-system-shapes": {
			Name:   "list-db-system-shapes",
			Fields: fields_list_db_system_shapes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDbSystemShapesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_db_system_shapes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDbSystemShapes(ctx, input)
				}
				var results []*svc.ListDbSystemShapesOutput
				p := svc.NewListDbSystemShapesPaginator(client, input)
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
		"list-gi-versions": {
			Name:   "list-gi-versions",
			Fields: fields_list_gi_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGiVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_gi_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGiVersions(ctx, input)
				}
				var results []*svc.ListGiVersionsOutput
				p := svc.NewListGiVersionsPaginator(client, input)
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
		"list-odb-networks": {
			Name:   "list-odb-networks",
			Fields: fields_list_odb_networks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOdbNetworksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_odb_networks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOdbNetworks(ctx, input)
				}
				var results []*svc.ListOdbNetworksOutput
				p := svc.NewListOdbNetworksPaginator(client, input)
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
		"list-odb-peering-connections": {
			Name:   "list-odb-peering-connections",
			Fields: fields_list_odb_peering_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOdbPeeringConnectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_odb_peering_connections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOdbPeeringConnections(ctx, input)
				}
				var results []*svc.ListOdbPeeringConnectionsOutput
				p := svc.NewListOdbPeeringConnectionsPaginator(client, input)
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
		"list-system-versions": {
			Name:   "list-system-versions",
			Fields: fields_list_system_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSystemVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_system_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSystemVersions(ctx, input)
				}
				var results []*svc.ListSystemVersionsOutput
				p := svc.NewListSystemVersionsPaginator(client, input)
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
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForResource(ctx, input)
			},
		},
		"reboot-db-node": {
			Name:   "reboot-db-node",
			Fields: fields_reboot_db_node,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RebootDbNodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reboot_db_node, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RebootDbNode(ctx, input)
			},
		},
		"start-db-node": {
			Name:   "start-db-node",
			Fields: fields_start_db_node,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDbNodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_db_node, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDbNode(ctx, input)
			},
		},
		"stop-db-node": {
			Name:   "stop-db-node",
			Fields: fields_stop_db_node,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopDbNodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_db_node, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopDbNode(ctx, input)
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
		"update-cloud-exadata-infrastructure": {
			Name:   "update-cloud-exadata-infrastructure",
			Fields: fields_update_cloud_exadata_infrastructure,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCloudExadataInfrastructureInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_cloud_exadata_infrastructure, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCloudExadataInfrastructure(ctx, input)
			},
		},
		"update-odb-network": {
			Name:   "update-odb-network",
			Fields: fields_update_odb_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateOdbNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_odb_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateOdbNetwork(ctx, input)
			},
		},
		"update-odb-peering-connection": {
			Name:   "update-odb-peering-connection",
			Fields: fields_update_odb_peering_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateOdbPeeringConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_odb_peering_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateOdbPeeringConnection(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("odb", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
