package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/odb"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// odbCmd represents the odb command
var _odbCmd = &cobra.Command{
	Use:   "odb",
	Short: "AWS odb CLI",
	Run: func(cmd *cobra.Command, args []string) {
		_awsOutput = resolveAWSOutput(_awsProfile, cmd.Flags().Changed("output"))
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := odb.NewFromConfig(cfg)
		if _odbAcceptMarketplaceRegistration {
			odb_AcceptMarketplaceRegistration(cfg, client)
			return
		}
		if _odbAssociateIamRoleToResource {
			odb_AssociateIamRoleToResource(cfg, client)
			return
		}
		if _odbCreateCloudAutonomousVmCluster {
			odb_CreateCloudAutonomousVmCluster(cfg, client)
			return
		}
		if _odbCreateCloudExadataInfrastructure {
			odb_CreateCloudExadataInfrastructure(cfg, client)
			return
		}
		if _odbCreateCloudVmCluster {
			odb_CreateCloudVmCluster(cfg, client)
			return
		}
		if _odbCreateOdbNetwork {
			odb_CreateOdbNetwork(cfg, client)
			return
		}
		if _odbCreateOdbPeeringConnection {
			odb_CreateOdbPeeringConnection(cfg, client)
			return
		}
		if _odbDeleteCloudAutonomousVmCluster {
			odb_DeleteCloudAutonomousVmCluster(cfg, client)
			return
		}
		if _odbDeleteCloudExadataInfrastructure {
			odb_DeleteCloudExadataInfrastructure(cfg, client)
			return
		}
		if _odbDeleteCloudVmCluster {
			odb_DeleteCloudVmCluster(cfg, client)
			return
		}
		if _odbDeleteOdbNetwork {
			odb_DeleteOdbNetwork(cfg, client)
			return
		}
		if _odbDeleteOdbPeeringConnection {
			odb_DeleteOdbPeeringConnection(cfg, client)
			return
		}
		if _odbDisassociateIamRoleFromResource {
			odb_DisassociateIamRoleFromResource(cfg, client)
			return
		}
		if _odbGetCloudAutonomousVmCluster {
			odb_GetCloudAutonomousVmCluster(cfg, client)
			return
		}
		if _odbGetCloudExadataInfrastructure {
			odb_GetCloudExadataInfrastructure(cfg, client)
			return
		}
		if _odbGetCloudExadataInfrastructureUnallocatedResources {
			odb_GetCloudExadataInfrastructureUnallocatedResources(cfg, client)
			return
		}
		if _odbGetCloudVmCluster {
			odb_GetCloudVmCluster(cfg, client)
			return
		}
		if _odbGetDbNode {
			odb_GetDbNode(cfg, client)
			return
		}
		if _odbGetDbServer {
			odb_GetDbServer(cfg, client)
			return
		}
		if _odbGetOciOnboardingStatus {
			odb_GetOciOnboardingStatus(cfg, client)
			return
		}
		if _odbGetOdbNetwork {
			odb_GetOdbNetwork(cfg, client)
			return
		}
		if _odbGetOdbPeeringConnection {
			odb_GetOdbPeeringConnection(cfg, client)
			return
		}
		if _odbInitializeService {
			odb_InitializeService(cfg, client)
			return
		}
		if _odbListAutonomousVirtualMachines {
			odb_ListAutonomousVirtualMachines(cfg, client)
			return
		}
		if _odbListCloudAutonomousVmClusters {
			odb_ListCloudAutonomousVmClusters(cfg, client)
			return
		}
		if _odbListCloudExadataInfrastructures {
			odb_ListCloudExadataInfrastructures(cfg, client)
			return
		}
		if _odbListCloudVmClusters {
			odb_ListCloudVmClusters(cfg, client)
			return
		}
		if _odbListDbNodes {
			odb_ListDbNodes(cfg, client)
			return
		}
		if _odbListDbServers {
			odb_ListDbServers(cfg, client)
			return
		}
		if _odbListDbSystemShapes {
			odb_ListDbSystemShapes(cfg, client)
			return
		}
		if _odbListGiVersions {
			odb_ListGiVersions(cfg, client)
			return
		}
		if _odbListOdbNetworks {
			odb_ListOdbNetworks(cfg, client)
			return
		}
		if _odbListOdbPeeringConnections {
			odb_ListOdbPeeringConnections(cfg, client)
			return
		}
		if _odbListSystemVersions {
			odb_ListSystemVersions(cfg, client)
			return
		}
		if _odbListTagsForResource {
			odb_ListTagsForResource(cfg, client)
			return
		}
		if _odbRebootDbNode {
			odb_RebootDbNode(cfg, client)
			return
		}
		if _odbStartDbNode {
			odb_StartDbNode(cfg, client)
			return
		}
		if _odbStopDbNode {
			odb_StopDbNode(cfg, client)
			return
		}
		if _odbTagResource {
			odb_TagResource(cfg, client)
			return
		}
		if _odbUntagResource {
			odb_UntagResource(cfg, client)
			return
		}
		if _odbUpdateCloudExadataInfrastructure {
			odb_UpdateCloudExadataInfrastructure(cfg, client)
			return
		}
		if _odbUpdateOdbNetwork {
			odb_UpdateOdbNetwork(cfg, client)
			return
		}
		if _odbUpdateOdbPeeringConnection {
			odb_UpdateOdbPeeringConnection(cfg, client)
			return
		}

	},
}

var (
	_odbAcceptMarketplaceRegistration                     bool
	_odbAssociateIamRoleToResource                        bool
	_odbCreateCloudAutonomousVmCluster                    bool
	_odbCreateCloudExadataInfrastructure                  bool
	_odbCreateCloudVmCluster                              bool
	_odbCreateOdbNetwork                                  bool
	_odbCreateOdbPeeringConnection                        bool
	_odbDeleteCloudAutonomousVmCluster                    bool
	_odbDeleteCloudExadataInfrastructure                  bool
	_odbDeleteCloudVmCluster                              bool
	_odbDeleteOdbNetwork                                  bool
	_odbDeleteOdbPeeringConnection                        bool
	_odbDisassociateIamRoleFromResource                   bool
	_odbGetCloudAutonomousVmCluster                       bool
	_odbGetCloudExadataInfrastructure                     bool
	_odbGetCloudExadataInfrastructureUnallocatedResources bool
	_odbGetCloudVmCluster                                 bool
	_odbGetDbNode                                         bool
	_odbGetDbServer                                       bool
	_odbGetOciOnboardingStatus                            bool
	_odbGetOdbNetwork                                     bool
	_odbGetOdbPeeringConnection                           bool
	_odbInitializeService                                 bool
	_odbListAutonomousVirtualMachines                     bool
	_odbListCloudAutonomousVmClusters                     bool
	_odbListCloudExadataInfrastructures                   bool
	_odbListCloudVmClusters                               bool
	_odbListDbNodes                                       bool
	_odbListDbServers                                     bool
	_odbListDbSystemShapes                                bool
	_odbListGiVersions                                    bool
	_odbListOdbNetworks                                   bool
	_odbListOdbPeeringConnections                         bool
	_odbListSystemVersions                                bool
	_odbListTagsForResource                               bool
	_odbRebootDbNode                                      bool
	_odbStartDbNode                                       bool
	_odbStopDbNode                                        bool
	_odbTagResource                                       bool
	_odbUntagResource                                     bool
	_odbUpdateCloudExadataInfrastructure                  bool
	_odbUpdateOdbNetwork                                  bool
	_odbUpdateOdbPeeringConnection                        bool

	_odbAutonomousDataStorageSizeInTBs       string
	_odbAvailabilityZone                     string
	_odbAvailabilityZoneId                   string
	_odbAwsIntegration                       string
	_odbBackupSubnetCidr                     string
	_odbClientSubnetCidr                     string
	_odbClientToken                          string
	_odbCloudAutonomousVmClusterId           string
	_odbCloudExadataInfrastructureId         string
	_odbCloudVmClusterId                     string
	_odbClusterName                          string
	_odbComputeCount                         string
	_odbCpuCoreCount                         string
	_odbCpuCoreCountPerNode                  string
	_odbCrossRegionS3RestoreSourcesToDisable []string
	_odbCrossRegionS3RestoreSourcesToEnable  []string
	_odbCustomDomainName                     string
	_odbCustomerContactsToSendToOCI          string
	_odbDataCollectionOptions                string
	_odbDataStorageSizeInTBs                 string
	_odbDatabaseServerType                   string
	_odbDbNodeId                             string
	_odbDbNodeStorageSizeInGBs               string
	_odbDbServerId                           string
	_odbDbServers                            []string
	_odbDefaultDnsPrefix                     string
	_odbDeleteAssociatedResources            string
	_odbDescription                          string
	_odbDisplayName                          string
	_odbGiVersion                            string
	_odbHostname                             string
	_odbIamRoleArn                           string
	_odbIsLocalBackupEnabled                 string
	_odbIsMtlsEnabledVmCluster               string
	_odbIsSparseDiskgroupEnabled             string
	_odbKmsAccess                            string
	_odbKmsPolicyDocument                    string
	_odbLicenseModel                         string
	_odbMaintenanceWindow                    string
	_odbMarketplaceRegistrationToken         string
	_odbMaxResults                           string
	_odbMemoryPerOracleComputeUnitInGBs      string
	_odbMemorySizeInGBs                      string
	_odbNextToken                            string
	_odbOciIdentityDomain                    string
	_odbOdbNetworkId                         string
	_odbOdbPeeringConnectionId               string
	_odbPeerNetworkCidrsToBeAdded            []string
	_odbPeerNetworkCidrsToBeRemoved          []string
	_odbPeerNetworkId                        string
	_odbPeerNetworkRouteTableIds             []string
	_odbPeeredCidrsToBeAdded                 []string
	_odbPeeredCidrsToBeRemoved               []string
	_odbResourceArn                          string
	_odbS3Access                             string
	_odbS3PolicyDocument                     string
	_odbScanListenerPortNonTls               string
	_odbScanListenerPortTcp                  string
	_odbScanListenerPortTls                  string
	_odbShape                                string
	_odbSshPublicKeys                        []string
	_odbStorageCount                         string
	_odbStorageServerType                    string
	_odbStsAccess                            string
	_odbStsPolicyDocument                    string
	_odbSystemVersion                        string
	_odbTagKeys                              []string
	_odbTags                                 string
	_odbTimeZone                             string
	_odbTotalContainerDatabases              string
	_odbZeroEtlAccess                        string
)

// Registers the Amazon Web Services Marketplace token for your Amazon Web
// Services account to activate your Oracle Database(at)Amazon Web Services
// subscription.
func odb_AcceptMarketplaceRegistration(cfg aws.Config, client *odb.Client) {
	input := &odb.AcceptMarketplaceRegistrationInput{
		// MarketplaceRegistrationToken: *string, // Required
	}

	if len(_odbMarketplaceRegistrationToken) > 0 {
		input.MarketplaceRegistrationToken = aws.String(_odbMarketplaceRegistrationToken)
	}

	if resp, err := client.AcceptMarketplaceRegistration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates an Amazon Web Services Identity and Access Management (IAM) service
// role with a specified resource to enable Amazon Web Services service
// integration.
func odb_AssociateIamRoleToResource(cfg aws.Config, client *odb.Client) {
	input := &odb.AssociateIamRoleToResourceInput{
		// AwsIntegration: types.SupportedAwsIntegration, // Required
		// IamRoleArn: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_odbAwsIntegration) > 0 {
		if err := assignInputField(input, "AwsIntegration", _odbAwsIntegration); err != nil {
			log.Errorf("invalid --aws-integration: %s", err.Error())
			return
		}
	}
	if len(_odbIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_odbIamRoleArn)
	}
	if len(_odbResourceArn) > 0 {
		input.ResourceArn = aws.String(_odbResourceArn)
	}

	if resp, err := client.AssociateIamRoleToResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Autonomous VM cluster in the specified Exadata infrastructure.
func odb_CreateCloudAutonomousVmCluster(cfg aws.Config, client *odb.Client) {
	input := &odb.CreateCloudAutonomousVmClusterInput{
		// AutonomousDataStorageSizeInTBs: *float64, // Required
		// CloudExadataInfrastructureId: *string, // Required
		// CpuCoreCountPerNode: *int32, // Required
		// DisplayName: *string, // Required
		// MemoryPerOracleComputeUnitInGBs: *int32, // Required
		// OdbNetworkId: *string, // Required
		// TotalContainerDatabases: *int32, // Required
	}

	if len(_odbAutonomousDataStorageSizeInTBs) > 0 {
		if err := assignInputField(input, "AutonomousDataStorageSizeInTBs", _odbAutonomousDataStorageSizeInTBs); err != nil {
			log.Errorf("invalid --autonomous-data-storage-size-in-tbs: %s", err.Error())
			return
		}
	}
	if len(_odbCloudExadataInfrastructureId) > 0 {
		input.CloudExadataInfrastructureId = aws.String(_odbCloudExadataInfrastructureId)
	}
	if len(_odbCpuCoreCountPerNode) > 0 {
		if err := assignInputField(input, "CpuCoreCountPerNode", _odbCpuCoreCountPerNode); err != nil {
			log.Errorf("invalid --cpu-core-count-per-node: %s", err.Error())
			return
		}
	}
	if len(_odbDisplayName) > 0 {
		input.DisplayName = aws.String(_odbDisplayName)
	}
	if len(_odbMemoryPerOracleComputeUnitInGBs) > 0 {
		if err := assignInputField(input, "MemoryPerOracleComputeUnitInGBs", _odbMemoryPerOracleComputeUnitInGBs); err != nil {
			log.Errorf("invalid --memory-per-oracle-compute-unit-in-gbs: %s", err.Error())
			return
		}
	}
	if len(_odbOdbNetworkId) > 0 {
		input.OdbNetworkId = aws.String(_odbOdbNetworkId)
	}
	if len(_odbTotalContainerDatabases) > 0 {
		if err := assignInputField(input, "TotalContainerDatabases", _odbTotalContainerDatabases); err != nil {
			log.Errorf("invalid --total-container-databases: %s", err.Error())
			return
		}
	}
	if len(_odbClientToken) > 0 {
		input.ClientToken = aws.String(_odbClientToken)
	}
	if len(_odbDbServers) > 0 {
		input.DbServers = append([]string(nil), _odbDbServers...)
	}
	if len(_odbDescription) > 0 {
		input.Description = aws.String(_odbDescription)
	}
	if len(_odbIsMtlsEnabledVmCluster) > 0 {
		if err := assignInputField(input, "IsMtlsEnabledVmCluster", _odbIsMtlsEnabledVmCluster); err != nil {
			log.Errorf("invalid --is-mtls-enabled-vm-cluster: %s", err.Error())
			return
		}
	}
	if len(_odbLicenseModel) > 0 {
		if err := assignInputField(input, "LicenseModel", _odbLicenseModel); err != nil {
			log.Errorf("invalid --license-model: %s", err.Error())
			return
		}
	}
	if len(_odbMaintenanceWindow) > 0 {
		if err := assignInputField(input, "MaintenanceWindow", _odbMaintenanceWindow); err != nil {
			log.Errorf("invalid --maintenance-window: %s", err.Error())
			return
		}
	}
	if len(_odbScanListenerPortNonTls) > 0 {
		if err := assignInputField(input, "ScanListenerPortNonTls", _odbScanListenerPortNonTls); err != nil {
			log.Errorf("invalid --scan-listener-port-non-tls: %s", err.Error())
			return
		}
	}
	if len(_odbScanListenerPortTls) > 0 {
		if err := assignInputField(input, "ScanListenerPortTls", _odbScanListenerPortTls); err != nil {
			log.Errorf("invalid --scan-listener-port-tls: %s", err.Error())
			return
		}
	}
	if len(_odbTags) > 0 {
		if err := assignInputField(input, "Tags", _odbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_odbTimeZone) > 0 {
		input.TimeZone = aws.String(_odbTimeZone)
	}

	if resp, err := client.CreateCloudAutonomousVmCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Exadata infrastructure.
func odb_CreateCloudExadataInfrastructure(cfg aws.Config, client *odb.Client) {
	input := &odb.CreateCloudExadataInfrastructureInput{
		// ComputeCount: *int32, // Required
		// DisplayName: *string, // Required
		// Shape: *string, // Required
		// StorageCount: *int32, // Required
	}

	if len(_odbComputeCount) > 0 {
		if err := assignInputField(input, "ComputeCount", _odbComputeCount); err != nil {
			log.Errorf("invalid --compute-count: %s", err.Error())
			return
		}
	}
	if len(_odbDisplayName) > 0 {
		input.DisplayName = aws.String(_odbDisplayName)
	}
	if len(_odbShape) > 0 {
		input.Shape = aws.String(_odbShape)
	}
	if len(_odbStorageCount) > 0 {
		if err := assignInputField(input, "StorageCount", _odbStorageCount); err != nil {
			log.Errorf("invalid --storage-count: %s", err.Error())
			return
		}
	}
	if len(_odbAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_odbAvailabilityZone)
	}
	if len(_odbAvailabilityZoneId) > 0 {
		input.AvailabilityZoneId = aws.String(_odbAvailabilityZoneId)
	}
	if len(_odbClientToken) > 0 {
		input.ClientToken = aws.String(_odbClientToken)
	}
	if len(_odbCustomerContactsToSendToOCI) > 0 {
		if err := assignInputField(input, "CustomerContactsToSendToOCI", _odbCustomerContactsToSendToOCI); err != nil {
			log.Errorf("invalid --customer-contacts-to-send-to-oci: %s", err.Error())
			return
		}
	}
	if len(_odbDatabaseServerType) > 0 {
		input.DatabaseServerType = aws.String(_odbDatabaseServerType)
	}
	if len(_odbMaintenanceWindow) > 0 {
		if err := assignInputField(input, "MaintenanceWindow", _odbMaintenanceWindow); err != nil {
			log.Errorf("invalid --maintenance-window: %s", err.Error())
			return
		}
	}
	if len(_odbStorageServerType) > 0 {
		input.StorageServerType = aws.String(_odbStorageServerType)
	}
	if len(_odbTags) > 0 {
		if err := assignInputField(input, "Tags", _odbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCloudExadataInfrastructure(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a VM cluster on the specified Exadata infrastructure.
func odb_CreateCloudVmCluster(cfg aws.Config, client *odb.Client) {
	input := &odb.CreateCloudVmClusterInput{
		// CloudExadataInfrastructureId: *string, // Required
		// CpuCoreCount: *int32, // Required
		// DisplayName: *string, // Required
		// GiVersion: *string, // Required
		// Hostname: *string, // Required
		// OdbNetworkId: *string, // Required
		// SshPublicKeys: []string, // Required
	}

	if len(_odbCloudExadataInfrastructureId) > 0 {
		input.CloudExadataInfrastructureId = aws.String(_odbCloudExadataInfrastructureId)
	}
	if len(_odbCpuCoreCount) > 0 {
		if err := assignInputField(input, "CpuCoreCount", _odbCpuCoreCount); err != nil {
			log.Errorf("invalid --cpu-core-count: %s", err.Error())
			return
		}
	}
	if len(_odbDisplayName) > 0 {
		input.DisplayName = aws.String(_odbDisplayName)
	}
	if len(_odbGiVersion) > 0 {
		input.GiVersion = aws.String(_odbGiVersion)
	}
	if len(_odbHostname) > 0 {
		input.Hostname = aws.String(_odbHostname)
	}
	if len(_odbOdbNetworkId) > 0 {
		input.OdbNetworkId = aws.String(_odbOdbNetworkId)
	}
	if len(_odbSshPublicKeys) > 0 {
		input.SshPublicKeys = append([]string(nil), _odbSshPublicKeys...)
	}
	if len(_odbClientToken) > 0 {
		input.ClientToken = aws.String(_odbClientToken)
	}
	if len(_odbClusterName) > 0 {
		input.ClusterName = aws.String(_odbClusterName)
	}
	if len(_odbDataCollectionOptions) > 0 {
		if err := assignInputField(input, "DataCollectionOptions", _odbDataCollectionOptions); err != nil {
			log.Errorf("invalid --data-collection-options: %s", err.Error())
			return
		}
	}
	if len(_odbDataStorageSizeInTBs) > 0 {
		if err := assignInputField(input, "DataStorageSizeInTBs", _odbDataStorageSizeInTBs); err != nil {
			log.Errorf("invalid --data-storage-size-in-tbs: %s", err.Error())
			return
		}
	}
	if len(_odbDbNodeStorageSizeInGBs) > 0 {
		if err := assignInputField(input, "DbNodeStorageSizeInGBs", _odbDbNodeStorageSizeInGBs); err != nil {
			log.Errorf("invalid --db-node-storage-size-in-gbs: %s", err.Error())
			return
		}
	}
	if len(_odbDbServers) > 0 {
		input.DbServers = append([]string(nil), _odbDbServers...)
	}
	if len(_odbIsLocalBackupEnabled) > 0 {
		if err := assignInputField(input, "IsLocalBackupEnabled", _odbIsLocalBackupEnabled); err != nil {
			log.Errorf("invalid --is-local-backup-enabled: %s", err.Error())
			return
		}
	}
	if len(_odbIsSparseDiskgroupEnabled) > 0 {
		if err := assignInputField(input, "IsSparseDiskgroupEnabled", _odbIsSparseDiskgroupEnabled); err != nil {
			log.Errorf("invalid --is-sparse-diskgroup-enabled: %s", err.Error())
			return
		}
	}
	if len(_odbLicenseModel) > 0 {
		if err := assignInputField(input, "LicenseModel", _odbLicenseModel); err != nil {
			log.Errorf("invalid --license-model: %s", err.Error())
			return
		}
	}
	if len(_odbMemorySizeInGBs) > 0 {
		if err := assignInputField(input, "MemorySizeInGBs", _odbMemorySizeInGBs); err != nil {
			log.Errorf("invalid --memory-size-in-gbs: %s", err.Error())
			return
		}
	}
	if len(_odbScanListenerPortTcp) > 0 {
		if err := assignInputField(input, "ScanListenerPortTcp", _odbScanListenerPortTcp); err != nil {
			log.Errorf("invalid --scan-listener-port-tcp: %s", err.Error())
			return
		}
	}
	if len(_odbSystemVersion) > 0 {
		input.SystemVersion = aws.String(_odbSystemVersion)
	}
	if len(_odbTags) > 0 {
		if err := assignInputField(input, "Tags", _odbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_odbTimeZone) > 0 {
		input.TimeZone = aws.String(_odbTimeZone)
	}

	if resp, err := client.CreateCloudVmCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an ODB network.
func odb_CreateOdbNetwork(cfg aws.Config, client *odb.Client) {
	input := &odb.CreateOdbNetworkInput{
		// ClientSubnetCidr: *string, // Required
		// DisplayName: *string, // Required
	}

	if len(_odbClientSubnetCidr) > 0 {
		input.ClientSubnetCidr = aws.String(_odbClientSubnetCidr)
	}
	if len(_odbDisplayName) > 0 {
		input.DisplayName = aws.String(_odbDisplayName)
	}
	if len(_odbAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_odbAvailabilityZone)
	}
	if len(_odbAvailabilityZoneId) > 0 {
		input.AvailabilityZoneId = aws.String(_odbAvailabilityZoneId)
	}
	if len(_odbBackupSubnetCidr) > 0 {
		input.BackupSubnetCidr = aws.String(_odbBackupSubnetCidr)
	}
	if len(_odbClientToken) > 0 {
		input.ClientToken = aws.String(_odbClientToken)
	}
	if len(_odbCrossRegionS3RestoreSourcesToEnable) > 0 {
		input.CrossRegionS3RestoreSourcesToEnable = append([]string(nil), _odbCrossRegionS3RestoreSourcesToEnable...)
	}
	if len(_odbCustomDomainName) > 0 {
		input.CustomDomainName = aws.String(_odbCustomDomainName)
	}
	if len(_odbDefaultDnsPrefix) > 0 {
		input.DefaultDnsPrefix = aws.String(_odbDefaultDnsPrefix)
	}
	if len(_odbKmsAccess) > 0 {
		if err := assignInputField(input, "KmsAccess", _odbKmsAccess); err != nil {
			log.Errorf("invalid --kms-access: %s", err.Error())
			return
		}
	}
	if len(_odbKmsPolicyDocument) > 0 {
		input.KmsPolicyDocument = aws.String(_odbKmsPolicyDocument)
	}
	if len(_odbS3Access) > 0 {
		if err := assignInputField(input, "S3Access", _odbS3Access); err != nil {
			log.Errorf("invalid --s3-access: %s", err.Error())
			return
		}
	}
	if len(_odbS3PolicyDocument) > 0 {
		input.S3PolicyDocument = aws.String(_odbS3PolicyDocument)
	}
	if len(_odbStsAccess) > 0 {
		if err := assignInputField(input, "StsAccess", _odbStsAccess); err != nil {
			log.Errorf("invalid --sts-access: %s", err.Error())
			return
		}
	}
	if len(_odbStsPolicyDocument) > 0 {
		input.StsPolicyDocument = aws.String(_odbStsPolicyDocument)
	}
	if len(_odbTags) > 0 {
		if err := assignInputField(input, "Tags", _odbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_odbZeroEtlAccess) > 0 {
		if err := assignInputField(input, "ZeroEtlAccess", _odbZeroEtlAccess); err != nil {
			log.Errorf("invalid --zero-etl-access: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOdbNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a peering connection between an ODB network and a VPC.
// A peering connection enables private connectivity between the networks for
// application-tier communication.
func odb_CreateOdbPeeringConnection(cfg aws.Config, client *odb.Client) {
	input := &odb.CreateOdbPeeringConnectionInput{
		// OdbNetworkId: *string, // Required
		// PeerNetworkId: *string, // Required
	}

	if len(_odbOdbNetworkId) > 0 {
		input.OdbNetworkId = aws.String(_odbOdbNetworkId)
	}
	if len(_odbPeerNetworkId) > 0 {
		input.PeerNetworkId = aws.String(_odbPeerNetworkId)
	}
	if len(_odbClientToken) > 0 {
		input.ClientToken = aws.String(_odbClientToken)
	}
	if len(_odbDisplayName) > 0 {
		input.DisplayName = aws.String(_odbDisplayName)
	}
	if len(_odbPeerNetworkCidrsToBeAdded) > 0 {
		input.PeerNetworkCidrsToBeAdded = append([]string(nil), _odbPeerNetworkCidrsToBeAdded...)
	}
	if len(_odbPeerNetworkRouteTableIds) > 0 {
		input.PeerNetworkRouteTableIds = append([]string(nil), _odbPeerNetworkRouteTableIds...)
	}
	if len(_odbTags) > 0 {
		if err := assignInputField(input, "Tags", _odbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOdbPeeringConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Autonomous VM cluster.
func odb_DeleteCloudAutonomousVmCluster(cfg aws.Config, client *odb.Client) {
	input := &odb.DeleteCloudAutonomousVmClusterInput{
		// CloudAutonomousVmClusterId: *string, // Required
	}

	if len(_odbCloudAutonomousVmClusterId) > 0 {
		input.CloudAutonomousVmClusterId = aws.String(_odbCloudAutonomousVmClusterId)
	}

	if resp, err := client.DeleteCloudAutonomousVmCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified Exadata infrastructure. Before you use this operation,
// make sure to delete all of the VM clusters that are hosted on this Exadata
// infrastructure.
func odb_DeleteCloudExadataInfrastructure(cfg aws.Config, client *odb.Client) {
	input := &odb.DeleteCloudExadataInfrastructureInput{
		// CloudExadataInfrastructureId: *string, // Required
	}

	if len(_odbCloudExadataInfrastructureId) > 0 {
		input.CloudExadataInfrastructureId = aws.String(_odbCloudExadataInfrastructureId)
	}

	if resp, err := client.DeleteCloudExadataInfrastructure(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified VM cluster.
func odb_DeleteCloudVmCluster(cfg aws.Config, client *odb.Client) {
	input := &odb.DeleteCloudVmClusterInput{
		// CloudVmClusterId: *string, // Required
	}

	if len(_odbCloudVmClusterId) > 0 {
		input.CloudVmClusterId = aws.String(_odbCloudVmClusterId)
	}

	if resp, err := client.DeleteCloudVmCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified ODB network.
func odb_DeleteOdbNetwork(cfg aws.Config, client *odb.Client) {
	input := &odb.DeleteOdbNetworkInput{
		// DeleteAssociatedResources: *bool, // Required
		// OdbNetworkId: *string, // Required
	}

	if len(_odbDeleteAssociatedResources) > 0 {
		if err := assignInputField(input, "DeleteAssociatedResources", _odbDeleteAssociatedResources); err != nil {
			log.Errorf("invalid --delete-associated-resources: %s", err.Error())
			return
		}
	}
	if len(_odbOdbNetworkId) > 0 {
		input.OdbNetworkId = aws.String(_odbOdbNetworkId)
	}

	if resp, err := client.DeleteOdbNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an ODB peering connection.
// When you delete an ODB peering connection, the underlying VPC peering
// connection is also deleted.
func odb_DeleteOdbPeeringConnection(cfg aws.Config, client *odb.Client) {
	input := &odb.DeleteOdbPeeringConnectionInput{
		// OdbPeeringConnectionId: *string, // Required
	}

	if len(_odbOdbPeeringConnectionId) > 0 {
		input.OdbPeeringConnectionId = aws.String(_odbOdbPeeringConnectionId)
	}

	if resp, err := client.DeleteOdbPeeringConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates an Amazon Web Services Identity and Access Management (IAM)
// service role from a specified resource to disable Amazon Web Services service
// integration.
func odb_DisassociateIamRoleFromResource(cfg aws.Config, client *odb.Client) {
	input := &odb.DisassociateIamRoleFromResourceInput{
		// AwsIntegration: types.SupportedAwsIntegration, // Required
		// IamRoleArn: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_odbAwsIntegration) > 0 {
		if err := assignInputField(input, "AwsIntegration", _odbAwsIntegration); err != nil {
			log.Errorf("invalid --aws-integration: %s", err.Error())
			return
		}
	}
	if len(_odbIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_odbIamRoleArn)
	}
	if len(_odbResourceArn) > 0 {
		input.ResourceArn = aws.String(_odbResourceArn)
	}

	if resp, err := client.DisassociateIamRoleFromResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specific Autonomous VM cluster.
func odb_GetCloudAutonomousVmCluster(cfg aws.Config, client *odb.Client) {
	input := &odb.GetCloudAutonomousVmClusterInput{
		// CloudAutonomousVmClusterId: *string, // Required
	}

	if len(_odbCloudAutonomousVmClusterId) > 0 {
		input.CloudAutonomousVmClusterId = aws.String(_odbCloudAutonomousVmClusterId)
	}

	if resp, err := client.GetCloudAutonomousVmCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified Exadata infrastructure.
func odb_GetCloudExadataInfrastructure(cfg aws.Config, client *odb.Client) {
	input := &odb.GetCloudExadataInfrastructureInput{
		// CloudExadataInfrastructureId: *string, // Required
	}

	if len(_odbCloudExadataInfrastructureId) > 0 {
		input.CloudExadataInfrastructureId = aws.String(_odbCloudExadataInfrastructureId)
	}

	if resp, err := client.GetCloudExadataInfrastructure(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about unallocated resources in a specified Cloud Exadata
// Infrastructure.
func odb_GetCloudExadataInfrastructureUnallocatedResources(cfg aws.Config, client *odb.Client) {
	input := &odb.GetCloudExadataInfrastructureUnallocatedResourcesInput{
		// CloudExadataInfrastructureId: *string, // Required
	}

	if len(_odbCloudExadataInfrastructureId) > 0 {
		input.CloudExadataInfrastructureId = aws.String(_odbCloudExadataInfrastructureId)
	}
	if len(_odbDbServers) > 0 {
		input.DbServers = append([]string(nil), _odbDbServers...)
	}

	if resp, err := client.GetCloudExadataInfrastructureUnallocatedResources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified VM cluster.
func odb_GetCloudVmCluster(cfg aws.Config, client *odb.Client) {
	input := &odb.GetCloudVmClusterInput{
		// CloudVmClusterId: *string, // Required
	}

	if len(_odbCloudVmClusterId) > 0 {
		input.CloudVmClusterId = aws.String(_odbCloudVmClusterId)
	}

	if resp, err := client.GetCloudVmCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified DB node.
func odb_GetDbNode(cfg aws.Config, client *odb.Client) {
	input := &odb.GetDbNodeInput{
		// CloudVmClusterId: *string, // Required
		// DbNodeId: *string, // Required
	}

	if len(_odbCloudVmClusterId) > 0 {
		input.CloudVmClusterId = aws.String(_odbCloudVmClusterId)
	}
	if len(_odbDbNodeId) > 0 {
		input.DbNodeId = aws.String(_odbDbNodeId)
	}

	if resp, err := client.GetDbNode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified database server.
func odb_GetDbServer(cfg aws.Config, client *odb.Client) {
	input := &odb.GetDbServerInput{
		// CloudExadataInfrastructureId: *string, // Required
		// DbServerId: *string, // Required
	}

	if len(_odbCloudExadataInfrastructureId) > 0 {
		input.CloudExadataInfrastructureId = aws.String(_odbCloudExadataInfrastructureId)
	}
	if len(_odbDbServerId) > 0 {
		input.DbServerId = aws.String(_odbDbServerId)
	}

	if resp, err := client.GetDbServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the tenancy activation link and onboarding status for your Amazon Web
// Services account.
func odb_GetOciOnboardingStatus(cfg aws.Config, client *odb.Client) {
	input := &odb.GetOciOnboardingStatusInput{}

	if resp, err := client.GetOciOnboardingStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified ODB network.
func odb_GetOdbNetwork(cfg aws.Config, client *odb.Client) {
	input := &odb.GetOdbNetworkInput{
		// OdbNetworkId: *string, // Required
	}

	if len(_odbOdbNetworkId) > 0 {
		input.OdbNetworkId = aws.String(_odbOdbNetworkId)
	}

	if resp, err := client.GetOdbNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an ODB peering connection.
func odb_GetOdbPeeringConnection(cfg aws.Config, client *odb.Client) {
	input := &odb.GetOdbPeeringConnectionInput{
		// OdbPeeringConnectionId: *string, // Required
	}

	if len(_odbOdbPeeringConnectionId) > 0 {
		input.OdbPeeringConnectionId = aws.String(_odbOdbPeeringConnectionId)
	}

	if resp, err := client.GetOdbPeeringConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initializes the ODB service for the first time in an account.
func odb_InitializeService(cfg aws.Config, client *odb.Client) {
	input := &odb.InitializeServiceInput{}

	if len(_odbOciIdentityDomain) > 0 {
		if err := assignInputField(input, "OciIdentityDomain", _odbOciIdentityDomain); err != nil {
			log.Errorf("invalid --oci-identity-domain: %s", err.Error())
			return
		}
	}

	if resp, err := client.InitializeService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all Autonomous VMs in an Autonomous VM cluster.
func odb_ListAutonomousVirtualMachines(cfg aws.Config, client *odb.Client) {
	input := &odb.ListAutonomousVirtualMachinesInput{
		// CloudAutonomousVmClusterId: *string, // Required
	}

	if len(_odbCloudAutonomousVmClusterId) > 0 {
		input.CloudAutonomousVmClusterId = aws.String(_odbCloudAutonomousVmClusterId)
	}
	if len(_odbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _odbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_odbNextToken) > 0 {
		input.NextToken = aws.String(_odbNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAutonomousVirtualMachines(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*odb.ListAutonomousVirtualMachinesOutput
	p := odb.NewListAutonomousVirtualMachinesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists all Autonomous VM clusters in a specified Cloud Exadata infrastructure.
func odb_ListCloudAutonomousVmClusters(cfg aws.Config, client *odb.Client) {
	input := &odb.ListCloudAutonomousVmClustersInput{}

	if len(_odbCloudExadataInfrastructureId) > 0 {
		input.CloudExadataInfrastructureId = aws.String(_odbCloudExadataInfrastructureId)
	}
	if len(_odbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _odbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_odbNextToken) > 0 {
		input.NextToken = aws.String(_odbNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCloudAutonomousVmClusters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*odb.ListCloudAutonomousVmClustersOutput
	p := odb.NewListCloudAutonomousVmClustersPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns information about the Exadata infrastructures owned by your Amazon Web
// Services account.
func odb_ListCloudExadataInfrastructures(cfg aws.Config, client *odb.Client) {
	input := &odb.ListCloudExadataInfrastructuresInput{}

	if len(_odbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _odbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_odbNextToken) > 0 {
		input.NextToken = aws.String(_odbNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCloudExadataInfrastructures(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*odb.ListCloudExadataInfrastructuresOutput
	p := odb.NewListCloudExadataInfrastructuresPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns information about the VM clusters owned by your Amazon Web Services
// account or only the ones on the specified Exadata infrastructure.
func odb_ListCloudVmClusters(cfg aws.Config, client *odb.Client) {
	input := &odb.ListCloudVmClustersInput{}

	if len(_odbCloudExadataInfrastructureId) > 0 {
		input.CloudExadataInfrastructureId = aws.String(_odbCloudExadataInfrastructureId)
	}
	if len(_odbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _odbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_odbNextToken) > 0 {
		input.NextToken = aws.String(_odbNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCloudVmClusters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*odb.ListCloudVmClustersOutput
	p := odb.NewListCloudVmClustersPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns information about the DB nodes for the specified VM cluster.
func odb_ListDbNodes(cfg aws.Config, client *odb.Client) {
	input := &odb.ListDbNodesInput{
		// CloudVmClusterId: *string, // Required
	}

	if len(_odbCloudVmClusterId) > 0 {
		input.CloudVmClusterId = aws.String(_odbCloudVmClusterId)
	}
	if len(_odbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _odbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_odbNextToken) > 0 {
		input.NextToken = aws.String(_odbNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDbNodes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*odb.ListDbNodesOutput
	p := odb.NewListDbNodesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns information about the database servers that belong to the specified
// Exadata infrastructure.
func odb_ListDbServers(cfg aws.Config, client *odb.Client) {
	input := &odb.ListDbServersInput{
		// CloudExadataInfrastructureId: *string, // Required
	}

	if len(_odbCloudExadataInfrastructureId) > 0 {
		input.CloudExadataInfrastructureId = aws.String(_odbCloudExadataInfrastructureId)
	}
	if len(_odbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _odbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_odbNextToken) > 0 {
		input.NextToken = aws.String(_odbNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDbServers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*odb.ListDbServersOutput
	p := odb.NewListDbServersPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns information about the shapes that are available for an Exadata
// infrastructure.
func odb_ListDbSystemShapes(cfg aws.Config, client *odb.Client) {
	input := &odb.ListDbSystemShapesInput{}

	if len(_odbAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_odbAvailabilityZone)
	}
	if len(_odbAvailabilityZoneId) > 0 {
		input.AvailabilityZoneId = aws.String(_odbAvailabilityZoneId)
	}
	if len(_odbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _odbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_odbNextToken) > 0 {
		input.NextToken = aws.String(_odbNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDbSystemShapes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*odb.ListDbSystemShapesOutput
	p := odb.NewListDbSystemShapesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns information about Oracle Grid Infrastructure (GI) software versions
// that are available for a VM cluster for the specified shape.
func odb_ListGiVersions(cfg aws.Config, client *odb.Client) {
	input := &odb.ListGiVersionsInput{}

	if len(_odbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _odbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_odbNextToken) > 0 {
		input.NextToken = aws.String(_odbNextToken)
	}
	if len(_odbShape) > 0 {
		input.Shape = aws.String(_odbShape)
	}

	if disablePaginator() {
		if resp, err := client.ListGiVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*odb.ListGiVersionsOutput
	p := odb.NewListGiVersionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns information about the ODB networks owned by your Amazon Web Services
// account.
func odb_ListOdbNetworks(cfg aws.Config, client *odb.Client) {
	input := &odb.ListOdbNetworksInput{}

	if len(_odbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _odbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_odbNextToken) > 0 {
		input.NextToken = aws.String(_odbNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOdbNetworks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*odb.ListOdbNetworksOutput
	p := odb.NewListOdbNetworksPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists all ODB peering connections or those associated with a specific ODB
// network.
func odb_ListOdbPeeringConnections(cfg aws.Config, client *odb.Client) {
	input := &odb.ListOdbPeeringConnectionsInput{}

	if len(_odbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _odbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_odbNextToken) > 0 {
		input.NextToken = aws.String(_odbNextToken)
	}
	if len(_odbOdbNetworkId) > 0 {
		input.OdbNetworkId = aws.String(_odbOdbNetworkId)
	}

	if disablePaginator() {
		if resp, err := client.ListOdbPeeringConnections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*odb.ListOdbPeeringConnectionsOutput
	p := odb.NewListOdbPeeringConnectionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns information about the system versions that are available for a VM
// cluster for the specified giVersion and shape .
func odb_ListSystemVersions(cfg aws.Config, client *odb.Client) {
	input := &odb.ListSystemVersionsInput{
		// GiVersion: *string, // Required
		// Shape: *string, // Required
	}

	if len(_odbGiVersion) > 0 {
		input.GiVersion = aws.String(_odbGiVersion)
	}
	if len(_odbShape) > 0 {
		input.Shape = aws.String(_odbShape)
	}
	if len(_odbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _odbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_odbNextToken) > 0 {
		input.NextToken = aws.String(_odbNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSystemVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*odb.ListSystemVersionsOutput
	p := odb.NewListSystemVersionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns information about the tags applied to this resource.
func odb_ListTagsForResource(cfg aws.Config, client *odb.Client) {
	input := &odb.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_odbResourceArn) > 0 {
		input.ResourceArn = aws.String(_odbResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Reboots the specified DB node in a VM cluster.
func odb_RebootDbNode(cfg aws.Config, client *odb.Client) {
	input := &odb.RebootDbNodeInput{
		// CloudVmClusterId: *string, // Required
		// DbNodeId: *string, // Required
	}

	if len(_odbCloudVmClusterId) > 0 {
		input.CloudVmClusterId = aws.String(_odbCloudVmClusterId)
	}
	if len(_odbDbNodeId) > 0 {
		input.DbNodeId = aws.String(_odbDbNodeId)
	}

	if resp, err := client.RebootDbNode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the specified DB node in a VM cluster.
func odb_StartDbNode(cfg aws.Config, client *odb.Client) {
	input := &odb.StartDbNodeInput{
		// CloudVmClusterId: *string, // Required
		// DbNodeId: *string, // Required
	}

	if len(_odbCloudVmClusterId) > 0 {
		input.CloudVmClusterId = aws.String(_odbCloudVmClusterId)
	}
	if len(_odbDbNodeId) > 0 {
		input.DbNodeId = aws.String(_odbDbNodeId)
	}

	if resp, err := client.StartDbNode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the specified DB node in a VM cluster.
func odb_StopDbNode(cfg aws.Config, client *odb.Client) {
	input := &odb.StopDbNodeInput{
		// CloudVmClusterId: *string, // Required
		// DbNodeId: *string, // Required
	}

	if len(_odbCloudVmClusterId) > 0 {
		input.CloudVmClusterId = aws.String(_odbCloudVmClusterId)
	}
	if len(_odbDbNodeId) > 0 {
		input.DbNodeId = aws.String(_odbDbNodeId)
	}

	if resp, err := client.StopDbNode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies tags to the specified resource.
func odb_TagResource(cfg aws.Config, client *odb.Client) {
	input := &odb.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_odbResourceArn) > 0 {
		input.ResourceArn = aws.String(_odbResourceArn)
	}
	if len(_odbTags) > 0 {
		if err := assignInputField(input, "Tags", _odbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes tags from the specified resource.
func odb_UntagResource(cfg aws.Config, client *odb.Client) {
	input := &odb.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_odbResourceArn) > 0 {
		input.ResourceArn = aws.String(_odbResourceArn)
	}
	if len(_odbTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _odbTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the properties of an Exadata infrastructure resource.
func odb_UpdateCloudExadataInfrastructure(cfg aws.Config, client *odb.Client) {
	input := &odb.UpdateCloudExadataInfrastructureInput{
		// CloudExadataInfrastructureId: *string, // Required
	}

	if len(_odbCloudExadataInfrastructureId) > 0 {
		input.CloudExadataInfrastructureId = aws.String(_odbCloudExadataInfrastructureId)
	}
	if len(_odbMaintenanceWindow) > 0 {
		if err := assignInputField(input, "MaintenanceWindow", _odbMaintenanceWindow); err != nil {
			log.Errorf("invalid --maintenance-window: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCloudExadataInfrastructure(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates properties of a specified ODB network.
func odb_UpdateOdbNetwork(cfg aws.Config, client *odb.Client) {
	input := &odb.UpdateOdbNetworkInput{
		// OdbNetworkId: *string, // Required
	}

	if len(_odbOdbNetworkId) > 0 {
		input.OdbNetworkId = aws.String(_odbOdbNetworkId)
	}
	if len(_odbCrossRegionS3RestoreSourcesToDisable) > 0 {
		input.CrossRegionS3RestoreSourcesToDisable = append([]string(nil), _odbCrossRegionS3RestoreSourcesToDisable...)
	}
	if len(_odbCrossRegionS3RestoreSourcesToEnable) > 0 {
		input.CrossRegionS3RestoreSourcesToEnable = append([]string(nil), _odbCrossRegionS3RestoreSourcesToEnable...)
	}
	if len(_odbDisplayName) > 0 {
		input.DisplayName = aws.String(_odbDisplayName)
	}
	if len(_odbKmsAccess) > 0 {
		if err := assignInputField(input, "KmsAccess", _odbKmsAccess); err != nil {
			log.Errorf("invalid --kms-access: %s", err.Error())
			return
		}
	}
	if len(_odbKmsPolicyDocument) > 0 {
		input.KmsPolicyDocument = aws.String(_odbKmsPolicyDocument)
	}
	if len(_odbPeeredCidrsToBeAdded) > 0 {
		input.PeeredCidrsToBeAdded = append([]string(nil), _odbPeeredCidrsToBeAdded...)
	}
	if len(_odbPeeredCidrsToBeRemoved) > 0 {
		input.PeeredCidrsToBeRemoved = append([]string(nil), _odbPeeredCidrsToBeRemoved...)
	}
	if len(_odbS3Access) > 0 {
		if err := assignInputField(input, "S3Access", _odbS3Access); err != nil {
			log.Errorf("invalid --s3-access: %s", err.Error())
			return
		}
	}
	if len(_odbS3PolicyDocument) > 0 {
		input.S3PolicyDocument = aws.String(_odbS3PolicyDocument)
	}
	if len(_odbStsAccess) > 0 {
		if err := assignInputField(input, "StsAccess", _odbStsAccess); err != nil {
			log.Errorf("invalid --sts-access: %s", err.Error())
			return
		}
	}
	if len(_odbStsPolicyDocument) > 0 {
		input.StsPolicyDocument = aws.String(_odbStsPolicyDocument)
	}
	if len(_odbZeroEtlAccess) > 0 {
		if err := assignInputField(input, "ZeroEtlAccess", _odbZeroEtlAccess); err != nil {
			log.Errorf("invalid --zero-etl-access: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateOdbNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the settings of an Oracle Database(at)Amazon Web Services peering
// connection. You can update the display name and add or remove CIDR blocks from
// the peering connection.
func odb_UpdateOdbPeeringConnection(cfg aws.Config, client *odb.Client) {
	input := &odb.UpdateOdbPeeringConnectionInput{
		// OdbPeeringConnectionId: *string, // Required
	}

	if len(_odbOdbPeeringConnectionId) > 0 {
		input.OdbPeeringConnectionId = aws.String(_odbOdbPeeringConnectionId)
	}
	if len(_odbDisplayName) > 0 {
		input.DisplayName = aws.String(_odbDisplayName)
	}
	if len(_odbPeerNetworkCidrsToBeAdded) > 0 {
		input.PeerNetworkCidrsToBeAdded = append([]string(nil), _odbPeerNetworkCidrsToBeAdded...)
	}
	if len(_odbPeerNetworkCidrsToBeRemoved) > 0 {
		input.PeerNetworkCidrsToBeRemoved = append([]string(nil), _odbPeerNetworkCidrsToBeRemoved...)
	}

	if resp, err := client.UpdateOdbPeeringConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_odbCmd)
	_odbCmd.Flags().SortFlags = false

	_odbCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_odbCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_odbCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_odbCmd.Flags().StringVarP(&_odbAutonomousDataStorageSizeInTBs, "autonomous-data-storage-size-in-tbs", "", "", "Autonomous Data Storage Size In Tbs")
	_odbCmd.Flags().StringVarP(&_odbAvailabilityZone, "availability-zone", "", "", "Availability Zone")
	_odbCmd.Flags().StringVarP(&_odbAvailabilityZoneId, "availability-zone-id", "", "", "Availability Zone ID")
	_odbCmd.Flags().StringVarP(&_odbAwsIntegration, "aws-integration", "", "", "AWS Integration")
	_odbCmd.Flags().StringVarP(&_odbBackupSubnetCidr, "backup-subnet-cidr", "", "", "Backup Subnet CIDR")
	_odbCmd.Flags().StringVarP(&_odbClientSubnetCidr, "client-subnet-cidr", "", "", "Client Subnet CIDR")
	_odbCmd.Flags().StringVarP(&_odbClientToken, "client-token", "", "", "Client Token")
	_odbCmd.Flags().StringVarP(&_odbCloudAutonomousVmClusterId, "cloud-autonomous-vm-cluster-id", "", "", "Cloud Autonomous Vm Cluster ID")
	_odbCmd.Flags().StringVarP(&_odbCloudExadataInfrastructureId, "cloud-exadata-infrastructure-id", "", "", "Cloud Exadata Infrastructure ID")
	_odbCmd.Flags().StringVarP(&_odbCloudVmClusterId, "cloud-vm-cluster-id", "", "", "Cloud Vm Cluster ID")
	_odbCmd.Flags().StringVarP(&_odbClusterName, "cluster-name", "", "", "Cluster Name")
	_odbCmd.Flags().StringVarP(&_odbComputeCount, "compute-count", "", "", "Compute Count")
	_odbCmd.Flags().StringVarP(&_odbCpuCoreCount, "cpu-core-count", "", "", "CPU Core Count")
	_odbCmd.Flags().StringVarP(&_odbCpuCoreCountPerNode, "cpu-core-count-per-node", "", "", "CPU Core Count Per Node")
	_odbCmd.Flags().StringSliceVarP(&_odbCrossRegionS3RestoreSourcesToDisable, "cross-region-s3-restore-sources-to-disable", "", nil, "Cross Region S3 Restore Sources To Disable")
	_odbCmd.Flags().StringSliceVarP(&_odbCrossRegionS3RestoreSourcesToEnable, "cross-region-s3-restore-sources-to-enable", "", nil, "Cross Region S3 Restore Sources To Enable")
	_odbCmd.Flags().StringVarP(&_odbCustomDomainName, "custom-domain-name", "", "", "Custom Domain Name")
	_odbCmd.Flags().StringVarP(&_odbCustomerContactsToSendToOCI, "customer-contacts-to-send-to-oci", "", "", "Customer Contacts To Send To Oci")
	_odbCmd.Flags().StringVarP(&_odbDataCollectionOptions, "data-collection-options", "", "", "Data Collection Options")
	_odbCmd.Flags().StringVarP(&_odbDataStorageSizeInTBs, "data-storage-size-in-tbs", "", "", "Data Storage Size In Tbs")
	_odbCmd.Flags().StringVarP(&_odbDatabaseServerType, "database-server-type", "", "", "Database Server Type")
	_odbCmd.Flags().StringVarP(&_odbDbNodeId, "db-node-id", "", "", "DB Node ID")
	_odbCmd.Flags().StringVarP(&_odbDbNodeStorageSizeInGBs, "db-node-storage-size-in-gbs", "", "", "DB Node Storage Size In Gbs")
	_odbCmd.Flags().StringVarP(&_odbDbServerId, "db-server-id", "", "", "DB Server ID")
	_odbCmd.Flags().StringSliceVarP(&_odbDbServers, "db-servers", "", nil, "DB Servers")
	_odbCmd.Flags().StringVarP(&_odbDefaultDnsPrefix, "default-dns-prefix", "", "", "Default DNS Prefix")
	_odbCmd.Flags().StringVarP(&_odbDeleteAssociatedResources, "delete-associated-resources", "", "", "Delete Associated Resources")
	_odbCmd.Flags().StringVarP(&_odbDescription, "description", "", "", "Description")
	_odbCmd.Flags().StringVarP(&_odbDisplayName, "display-name", "", "", "Display Name")
	_odbCmd.Flags().StringVarP(&_odbGiVersion, "gi-version", "", "", "Gi Version")
	_odbCmd.Flags().StringVarP(&_odbHostname, "hostname", "", "", "Hostname")
	_odbCmd.Flags().StringVarP(&_odbIamRoleArn, "iam-role-arn", "", "", "IAM Role ARN")
	_odbCmd.Flags().StringVarP(&_odbIsLocalBackupEnabled, "is-local-backup-enabled", "", "", "Is Local Backup Enabled")
	_odbCmd.Flags().StringVarP(&_odbIsMtlsEnabledVmCluster, "is-mtls-enabled-vm-cluster", "", "", "Is Mtls Enabled Vm Cluster")
	_odbCmd.Flags().StringVarP(&_odbIsSparseDiskgroupEnabled, "is-sparse-diskgroup-enabled", "", "", "Is Sparse Diskgroup Enabled")
	_odbCmd.Flags().StringVarP(&_odbKmsAccess, "kms-access", "", "", "KMS Access")
	_odbCmd.Flags().StringVarP(&_odbKmsPolicyDocument, "kms-policy-document", "", "", "KMS Policy Document")
	_odbCmd.Flags().StringVarP(&_odbLicenseModel, "license-model", "", "", "License Model")
	_odbCmd.Flags().StringVarP(&_odbMaintenanceWindow, "maintenance-window", "", "", "Maintenance Window")
	_odbCmd.Flags().StringVarP(&_odbMarketplaceRegistrationToken, "marketplace-registration-token", "", "", "Marketplace Registration Token")
	_odbCmd.Flags().StringVarP(&_odbMaxResults, "max-results", "", "", "Max Results")
	_odbCmd.Flags().StringVarP(&_odbMemoryPerOracleComputeUnitInGBs, "memory-per-oracle-compute-unit-in-gbs", "", "", "Memory Per Oracle Compute Unit In Gbs")
	_odbCmd.Flags().StringVarP(&_odbMemorySizeInGBs, "memory-size-in-gbs", "", "", "Memory Size In Gbs")
	_odbCmd.Flags().StringVarP(&_odbNextToken, "next-token", "", "", "Next Token")
	_odbCmd.Flags().StringVarP(&_odbOciIdentityDomain, "oci-identity-domain", "", "", "Oci Identity Domain")
	_odbCmd.Flags().StringVarP(&_odbOdbNetworkId, "odb-network-id", "", "", "Odb Network ID")
	_odbCmd.Flags().StringVarP(&_odbOdbPeeringConnectionId, "odb-peering-connection-id", "", "", "Odb Peering Connection ID")
	_odbCmd.Flags().StringSliceVarP(&_odbPeerNetworkCidrsToBeAdded, "peer-network-cidrs-to-be-added", "", nil, "Peer Network Cidrs To Be Added")
	_odbCmd.Flags().StringSliceVarP(&_odbPeerNetworkCidrsToBeRemoved, "peer-network-cidrs-to-be-removed", "", nil, "Peer Network Cidrs To Be Removed")
	_odbCmd.Flags().StringVarP(&_odbPeerNetworkId, "peer-network-id", "", "", "Peer Network ID")
	_odbCmd.Flags().StringSliceVarP(&_odbPeerNetworkRouteTableIds, "peer-network-route-table-ids", "", nil, "Peer Network Route Table Ids")
	_odbCmd.Flags().StringSliceVarP(&_odbPeeredCidrsToBeAdded, "peered-cidrs-to-be-added", "", nil, "Peered Cidrs To Be Added")
	_odbCmd.Flags().StringSliceVarP(&_odbPeeredCidrsToBeRemoved, "peered-cidrs-to-be-removed", "", nil, "Peered Cidrs To Be Removed")
	_odbCmd.Flags().StringVarP(&_odbResourceArn, "resource-arn", "", "", "Resource ARN")
	_odbCmd.Flags().StringVarP(&_odbS3Access, "s3-access", "", "", "S3 Access")
	_odbCmd.Flags().StringVarP(&_odbS3PolicyDocument, "s3-policy-document", "", "", "S3 Policy Document")
	_odbCmd.Flags().StringVarP(&_odbScanListenerPortNonTls, "scan-listener-port-non-tls", "", "", "Scan Listener Port Non TLS")
	_odbCmd.Flags().StringVarP(&_odbScanListenerPortTcp, "scan-listener-port-tcp", "", "", "Scan Listener Port TCP")
	_odbCmd.Flags().StringVarP(&_odbScanListenerPortTls, "scan-listener-port-tls", "", "", "Scan Listener Port TLS")
	_odbCmd.Flags().StringVarP(&_odbShape, "shape", "", "", "Shape")
	_odbCmd.Flags().StringSliceVarP(&_odbSshPublicKeys, "ssh-public-keys", "", nil, "SSH Public Keys")
	_odbCmd.Flags().StringVarP(&_odbStorageCount, "storage-count", "", "", "Storage Count")
	_odbCmd.Flags().StringVarP(&_odbStorageServerType, "storage-server-type", "", "", "Storage Server Type")
	_odbCmd.Flags().StringVarP(&_odbStsAccess, "sts-access", "", "", "Sts Access")
	_odbCmd.Flags().StringVarP(&_odbStsPolicyDocument, "sts-policy-document", "", "", "Sts Policy Document")
	_odbCmd.Flags().StringVarP(&_odbSystemVersion, "system-version", "", "", "System Version")
	_odbCmd.Flags().StringSliceVarP(&_odbTagKeys, "tag-keys", "", nil, "Tag Keys")
	_odbCmd.Flags().StringVarP(&_odbTags, "tags", "", "", "Tags")
	_odbCmd.Flags().StringVarP(&_odbTimeZone, "time-zone", "", "", "Time Zone")
	_odbCmd.Flags().StringVarP(&_odbTotalContainerDatabases, "total-container-databases", "", "", "Total Container Databases")
	_odbCmd.Flags().StringVarP(&_odbZeroEtlAccess, "zero-etl-access", "", "", "Zero Etl Access")

	_odbCmd.Flags().BoolVarP(&_odbAcceptMarketplaceRegistration, "accept-marketplace-registration", "", false, "Accept Marketplace Registration")
	_odbCmd.Flags().BoolVarP(&_odbAssociateIamRoleToResource, "associate-iam-role-to-resource", "", false, "Associate IAM Role To Resource")
	_odbCmd.Flags().BoolVarP(&_odbCreateCloudAutonomousVmCluster, "create-cloud-autonomous-vm-cluster", "", false, "Create Cloud Autonomous Vm Cluster")
	_odbCmd.Flags().BoolVarP(&_odbCreateCloudExadataInfrastructure, "create-cloud-exadata-infrastructure", "", false, "Create Cloud Exadata Infrastructure")
	_odbCmd.Flags().BoolVarP(&_odbCreateCloudVmCluster, "create-cloud-vm-cluster", "", false, "Create Cloud Vm Cluster")
	_odbCmd.Flags().BoolVarP(&_odbCreateOdbNetwork, "create-odb-network", "", false, "Create Odb Network")
	_odbCmd.Flags().BoolVarP(&_odbCreateOdbPeeringConnection, "create-odb-peering-connection", "", false, "Create Odb Peering Connection")
	_odbCmd.Flags().BoolVarP(&_odbDeleteCloudAutonomousVmCluster, "delete-cloud-autonomous-vm-cluster", "", false, "Delete Cloud Autonomous Vm Cluster")
	_odbCmd.Flags().BoolVarP(&_odbDeleteCloudExadataInfrastructure, "delete-cloud-exadata-infrastructure", "", false, "Delete Cloud Exadata Infrastructure")
	_odbCmd.Flags().BoolVarP(&_odbDeleteCloudVmCluster, "delete-cloud-vm-cluster", "", false, "Delete Cloud Vm Cluster")
	_odbCmd.Flags().BoolVarP(&_odbDeleteOdbNetwork, "delete-odb-network", "", false, "Delete Odb Network")
	_odbCmd.Flags().BoolVarP(&_odbDeleteOdbPeeringConnection, "delete-odb-peering-connection", "", false, "Delete Odb Peering Connection")
	_odbCmd.Flags().BoolVarP(&_odbDisassociateIamRoleFromResource, "disassociate-iam-role-from-resource", "", false, "Disassociate IAM Role From Resource")
	_odbCmd.Flags().BoolVarP(&_odbGetCloudAutonomousVmCluster, "get-cloud-autonomous-vm-cluster", "", false, "Get Cloud Autonomous Vm Cluster")
	_odbCmd.Flags().BoolVarP(&_odbGetCloudExadataInfrastructure, "get-cloud-exadata-infrastructure", "", false, "Get Cloud Exadata Infrastructure")
	_odbCmd.Flags().BoolVarP(&_odbGetCloudExadataInfrastructureUnallocatedResources, "get-cloud-exadata-infrastructure-unallocated-resources", "", false, "Get Cloud Exadata Infrastructure Unallocated Resources")
	_odbCmd.Flags().BoolVarP(&_odbGetCloudVmCluster, "get-cloud-vm-cluster", "", false, "Get Cloud Vm Cluster")
	_odbCmd.Flags().BoolVarP(&_odbGetDbNode, "get-db-node", "", false, "Get DB Node")
	_odbCmd.Flags().BoolVarP(&_odbGetDbServer, "get-db-server", "", false, "Get DB Server")
	_odbCmd.Flags().BoolVarP(&_odbGetOciOnboardingStatus, "get-oci-onboarding-status", "", false, "Get Oci Onboarding Status")
	_odbCmd.Flags().BoolVarP(&_odbGetOdbNetwork, "get-odb-network", "", false, "Get Odb Network")
	_odbCmd.Flags().BoolVarP(&_odbGetOdbPeeringConnection, "get-odb-peering-connection", "", false, "Get Odb Peering Connection")
	_odbCmd.Flags().BoolVarP(&_odbInitializeService, "initialize-service", "", false, "Initialize Service")
	_odbCmd.Flags().BoolVarP(&_odbListAutonomousVirtualMachines, "list-autonomous-virtual-machines", "", false, "List Autonomous Virtual Machines")
	_odbCmd.Flags().BoolVarP(&_odbListCloudAutonomousVmClusters, "list-cloud-autonomous-vm-clusters", "", false, "List Cloud Autonomous Vm Clusters")
	_odbCmd.Flags().BoolVarP(&_odbListCloudExadataInfrastructures, "list-cloud-exadata-infrastructures", "", false, "List Cloud Exadata Infrastructures")
	_odbCmd.Flags().BoolVarP(&_odbListCloudVmClusters, "list-cloud-vm-clusters", "", false, "List Cloud Vm Clusters")
	_odbCmd.Flags().BoolVarP(&_odbListDbNodes, "list-db-nodes", "", false, "List DB Nodes")
	_odbCmd.Flags().BoolVarP(&_odbListDbServers, "list-db-servers", "", false, "List DB Servers")
	_odbCmd.Flags().BoolVarP(&_odbListDbSystemShapes, "list-db-system-shapes", "", false, "List DB System Shapes")
	_odbCmd.Flags().BoolVarP(&_odbListGiVersions, "list-gi-versions", "", false, "List Gi Versions")
	_odbCmd.Flags().BoolVarP(&_odbListOdbNetworks, "list-odb-networks", "", false, "List Odb Networks")
	_odbCmd.Flags().BoolVarP(&_odbListOdbPeeringConnections, "list-odb-peering-connections", "", false, "List Odb Peering Connections")
	_odbCmd.Flags().BoolVarP(&_odbListSystemVersions, "list-system-versions", "", false, "List System Versions")
	_odbCmd.Flags().BoolVarP(&_odbListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_odbCmd.Flags().BoolVarP(&_odbRebootDbNode, "reboot-db-node", "", false, "Reboot DB Node")
	_odbCmd.Flags().BoolVarP(&_odbStartDbNode, "start-db-node", "", false, "Start DB Node")
	_odbCmd.Flags().BoolVarP(&_odbStopDbNode, "stop-db-node", "", false, "Stop DB Node")
	_odbCmd.Flags().BoolVarP(&_odbTagResource, "tag-resource", "", false, "Tag Resource")
	_odbCmd.Flags().BoolVarP(&_odbUntagResource, "untag-resource", "", false, "Untag Resource")
	_odbCmd.Flags().BoolVarP(&_odbUpdateCloudExadataInfrastructure, "update-cloud-exadata-infrastructure", "", false, "Update Cloud Exadata Infrastructure")
	_odbCmd.Flags().BoolVarP(&_odbUpdateOdbNetwork, "update-odb-network", "", false, "Update Odb Network")
	_odbCmd.Flags().BoolVarP(&_odbUpdateOdbPeeringConnection, "update-odb-peering-connection", "", false, "Update Odb Peering Connection")

}
