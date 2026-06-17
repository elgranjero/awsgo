package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// lightsailCmd represents the lightsail command
var _lightsailCmd = &cobra.Command{
	Use:   "lightsail",
	Short: "AWS lightsail CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := lightsail.NewFromConfig(cfg)
		if _lightsailAllocateStaticIp {
			lightsail_AllocateStaticIp(cfg, client)
			return
		}
		if _lightsailAttachCertificateToDistribution {
			lightsail_AttachCertificateToDistribution(cfg, client)
			return
		}
		if _lightsailAttachDisk {
			lightsail_AttachDisk(cfg, client)
			return
		}
		if _lightsailAttachInstancesToLoadBalancer {
			lightsail_AttachInstancesToLoadBalancer(cfg, client)
			return
		}
		if _lightsailAttachLoadBalancerTlsCertificate {
			lightsail_AttachLoadBalancerTlsCertificate(cfg, client)
			return
		}
		if _lightsailAttachStaticIp {
			lightsail_AttachStaticIp(cfg, client)
			return
		}
		if _lightsailCloseInstancePublicPorts {
			lightsail_CloseInstancePublicPorts(cfg, client)
			return
		}
		if _lightsailCopySnapshot {
			lightsail_CopySnapshot(cfg, client)
			return
		}
		if _lightsailCreateBucket {
			lightsail_CreateBucket(cfg, client)
			return
		}
		if _lightsailCreateBucketAccessKey {
			lightsail_CreateBucketAccessKey(cfg, client)
			return
		}
		if _lightsailCreateCertificate {
			lightsail_CreateCertificate(cfg, client)
			return
		}
		if _lightsailCreateCloudFormationStack {
			lightsail_CreateCloudFormationStack(cfg, client)
			return
		}
		if _lightsailCreateContactMethod {
			lightsail_CreateContactMethod(cfg, client)
			return
		}
		if _lightsailCreateContainerService {
			lightsail_CreateContainerService(cfg, client)
			return
		}
		if _lightsailCreateContainerServiceDeployment {
			lightsail_CreateContainerServiceDeployment(cfg, client)
			return
		}
		if _lightsailCreateContainerServiceRegistryLogin {
			lightsail_CreateContainerServiceRegistryLogin(cfg, client)
			return
		}
		if _lightsailCreateDisk {
			lightsail_CreateDisk(cfg, client)
			return
		}
		if _lightsailCreateDiskFromSnapshot {
			lightsail_CreateDiskFromSnapshot(cfg, client)
			return
		}
		if _lightsailCreateDiskSnapshot {
			lightsail_CreateDiskSnapshot(cfg, client)
			return
		}
		if _lightsailCreateDistribution {
			lightsail_CreateDistribution(cfg, client)
			return
		}
		if _lightsailCreateDomain {
			lightsail_CreateDomain(cfg, client)
			return
		}
		if _lightsailCreateDomainEntry {
			lightsail_CreateDomainEntry(cfg, client)
			return
		}
		if _lightsailCreateGUISessionAccessDetails {
			lightsail_CreateGUISessionAccessDetails(cfg, client)
			return
		}
		if _lightsailCreateInstanceSnapshot {
			lightsail_CreateInstanceSnapshot(cfg, client)
			return
		}
		if _lightsailCreateInstances {
			lightsail_CreateInstances(cfg, client)
			return
		}
		if _lightsailCreateInstancesFromSnapshot {
			lightsail_CreateInstancesFromSnapshot(cfg, client)
			return
		}
		if _lightsailCreateKeyPair {
			lightsail_CreateKeyPair(cfg, client)
			return
		}
		if _lightsailCreateLoadBalancer {
			lightsail_CreateLoadBalancer(cfg, client)
			return
		}
		if _lightsailCreateLoadBalancerTlsCertificate {
			lightsail_CreateLoadBalancerTlsCertificate(cfg, client)
			return
		}
		if _lightsailCreateRelationalDatabase {
			lightsail_CreateRelationalDatabase(cfg, client)
			return
		}
		if _lightsailCreateRelationalDatabaseFromSnapshot {
			lightsail_CreateRelationalDatabaseFromSnapshot(cfg, client)
			return
		}
		if _lightsailCreateRelationalDatabaseSnapshot {
			lightsail_CreateRelationalDatabaseSnapshot(cfg, client)
			return
		}
		if _lightsailDeleteAlarm {
			lightsail_DeleteAlarm(cfg, client)
			return
		}
		if _lightsailDeleteAutoSnapshot {
			lightsail_DeleteAutoSnapshot(cfg, client)
			return
		}
		if _lightsailDeleteBucket {
			lightsail_DeleteBucket(cfg, client)
			return
		}
		if _lightsailDeleteBucketAccessKey {
			lightsail_DeleteBucketAccessKey(cfg, client)
			return
		}
		if _lightsailDeleteCertificate {
			lightsail_DeleteCertificate(cfg, client)
			return
		}
		if _lightsailDeleteContactMethod {
			lightsail_DeleteContactMethod(cfg, client)
			return
		}
		if _lightsailDeleteContainerImage {
			lightsail_DeleteContainerImage(cfg, client)
			return
		}
		if _lightsailDeleteContainerService {
			lightsail_DeleteContainerService(cfg, client)
			return
		}
		if _lightsailDeleteDisk {
			lightsail_DeleteDisk(cfg, client)
			return
		}
		if _lightsailDeleteDiskSnapshot {
			lightsail_DeleteDiskSnapshot(cfg, client)
			return
		}
		if _lightsailDeleteDistribution {
			lightsail_DeleteDistribution(cfg, client)
			return
		}
		if _lightsailDeleteDomain {
			lightsail_DeleteDomain(cfg, client)
			return
		}
		if _lightsailDeleteDomainEntry {
			lightsail_DeleteDomainEntry(cfg, client)
			return
		}
		if _lightsailDeleteInstance {
			lightsail_DeleteInstance(cfg, client)
			return
		}
		if _lightsailDeleteInstanceSnapshot {
			lightsail_DeleteInstanceSnapshot(cfg, client)
			return
		}
		if _lightsailDeleteKeyPair {
			lightsail_DeleteKeyPair(cfg, client)
			return
		}
		if _lightsailDeleteKnownHostKeys {
			lightsail_DeleteKnownHostKeys(cfg, client)
			return
		}
		if _lightsailDeleteLoadBalancer {
			lightsail_DeleteLoadBalancer(cfg, client)
			return
		}
		if _lightsailDeleteLoadBalancerTlsCertificate {
			lightsail_DeleteLoadBalancerTlsCertificate(cfg, client)
			return
		}
		if _lightsailDeleteRelationalDatabase {
			lightsail_DeleteRelationalDatabase(cfg, client)
			return
		}
		if _lightsailDeleteRelationalDatabaseSnapshot {
			lightsail_DeleteRelationalDatabaseSnapshot(cfg, client)
			return
		}
		if _lightsailDetachCertificateFromDistribution {
			lightsail_DetachCertificateFromDistribution(cfg, client)
			return
		}
		if _lightsailDetachDisk {
			lightsail_DetachDisk(cfg, client)
			return
		}
		if _lightsailDetachInstancesFromLoadBalancer {
			lightsail_DetachInstancesFromLoadBalancer(cfg, client)
			return
		}
		if _lightsailDetachStaticIp {
			lightsail_DetachStaticIp(cfg, client)
			return
		}
		if _lightsailDisableAddOn {
			lightsail_DisableAddOn(cfg, client)
			return
		}
		if _lightsailDownloadDefaultKeyPair {
			lightsail_DownloadDefaultKeyPair(cfg, client)
			return
		}
		if _lightsailEnableAddOn {
			lightsail_EnableAddOn(cfg, client)
			return
		}
		if _lightsailExportSnapshot {
			lightsail_ExportSnapshot(cfg, client)
			return
		}
		if _lightsailGetActiveNames {
			lightsail_GetActiveNames(cfg, client)
			return
		}
		if _lightsailGetAlarms {
			lightsail_GetAlarms(cfg, client)
			return
		}
		if _lightsailGetAutoSnapshots {
			lightsail_GetAutoSnapshots(cfg, client)
			return
		}
		if _lightsailGetBlueprints {
			lightsail_GetBlueprints(cfg, client)
			return
		}
		if _lightsailGetBucketAccessKeys {
			lightsail_GetBucketAccessKeys(cfg, client)
			return
		}
		if _lightsailGetBucketBundles {
			lightsail_GetBucketBundles(cfg, client)
			return
		}
		if _lightsailGetBucketMetricData {
			lightsail_GetBucketMetricData(cfg, client)
			return
		}
		if _lightsailGetBuckets {
			lightsail_GetBuckets(cfg, client)
			return
		}
		if _lightsailGetBundles {
			lightsail_GetBundles(cfg, client)
			return
		}
		if _lightsailGetCertificates {
			lightsail_GetCertificates(cfg, client)
			return
		}
		if _lightsailGetCloudFormationStackRecords {
			lightsail_GetCloudFormationStackRecords(cfg, client)
			return
		}
		if _lightsailGetContactMethods {
			lightsail_GetContactMethods(cfg, client)
			return
		}
		if _lightsailGetContainerAPIMetadata {
			lightsail_GetContainerAPIMetadata(cfg, client)
			return
		}
		if _lightsailGetContainerImages {
			lightsail_GetContainerImages(cfg, client)
			return
		}
		if _lightsailGetContainerLog {
			lightsail_GetContainerLog(cfg, client)
			return
		}
		if _lightsailGetContainerServiceDeployments {
			lightsail_GetContainerServiceDeployments(cfg, client)
			return
		}
		if _lightsailGetContainerServiceMetricData {
			lightsail_GetContainerServiceMetricData(cfg, client)
			return
		}
		if _lightsailGetContainerServicePowers {
			lightsail_GetContainerServicePowers(cfg, client)
			return
		}
		if _lightsailGetContainerServices {
			lightsail_GetContainerServices(cfg, client)
			return
		}
		if _lightsailGetCostEstimate {
			lightsail_GetCostEstimate(cfg, client)
			return
		}
		if _lightsailGetDisk {
			lightsail_GetDisk(cfg, client)
			return
		}
		if _lightsailGetDiskSnapshot {
			lightsail_GetDiskSnapshot(cfg, client)
			return
		}
		if _lightsailGetDiskSnapshots {
			lightsail_GetDiskSnapshots(cfg, client)
			return
		}
		if _lightsailGetDisks {
			lightsail_GetDisks(cfg, client)
			return
		}
		if _lightsailGetDistributionBundles {
			lightsail_GetDistributionBundles(cfg, client)
			return
		}
		if _lightsailGetDistributionLatestCacheReset {
			lightsail_GetDistributionLatestCacheReset(cfg, client)
			return
		}
		if _lightsailGetDistributionMetricData {
			lightsail_GetDistributionMetricData(cfg, client)
			return
		}
		if _lightsailGetDistributions {
			lightsail_GetDistributions(cfg, client)
			return
		}
		if _lightsailGetDomain {
			lightsail_GetDomain(cfg, client)
			return
		}
		if _lightsailGetDomains {
			lightsail_GetDomains(cfg, client)
			return
		}
		if _lightsailGetExportSnapshotRecords {
			lightsail_GetExportSnapshotRecords(cfg, client)
			return
		}
		if _lightsailGetInstance {
			lightsail_GetInstance(cfg, client)
			return
		}
		if _lightsailGetInstanceAccessDetails {
			lightsail_GetInstanceAccessDetails(cfg, client)
			return
		}
		if _lightsailGetInstanceMetricData {
			lightsail_GetInstanceMetricData(cfg, client)
			return
		}
		if _lightsailGetInstancePortStates {
			lightsail_GetInstancePortStates(cfg, client)
			return
		}
		if _lightsailGetInstanceSnapshot {
			lightsail_GetInstanceSnapshot(cfg, client)
			return
		}
		if _lightsailGetInstanceSnapshots {
			lightsail_GetInstanceSnapshots(cfg, client)
			return
		}
		if _lightsailGetInstanceState {
			lightsail_GetInstanceState(cfg, client)
			return
		}
		if _lightsailGetInstances {
			lightsail_GetInstances(cfg, client)
			return
		}
		if _lightsailGetKeyPair {
			lightsail_GetKeyPair(cfg, client)
			return
		}
		if _lightsailGetKeyPairs {
			lightsail_GetKeyPairs(cfg, client)
			return
		}
		if _lightsailGetLoadBalancer {
			lightsail_GetLoadBalancer(cfg, client)
			return
		}
		if _lightsailGetLoadBalancerMetricData {
			lightsail_GetLoadBalancerMetricData(cfg, client)
			return
		}
		if _lightsailGetLoadBalancerTlsCertificates {
			lightsail_GetLoadBalancerTlsCertificates(cfg, client)
			return
		}
		if _lightsailGetLoadBalancerTlsPolicies {
			lightsail_GetLoadBalancerTlsPolicies(cfg, client)
			return
		}
		if _lightsailGetLoadBalancers {
			lightsail_GetLoadBalancers(cfg, client)
			return
		}
		if _lightsailGetOperation {
			lightsail_GetOperation(cfg, client)
			return
		}
		if _lightsailGetOperations {
			lightsail_GetOperations(cfg, client)
			return
		}
		if _lightsailGetOperationsForResource {
			lightsail_GetOperationsForResource(cfg, client)
			return
		}
		if _lightsailGetRegions {
			lightsail_GetRegions(cfg, client)
			return
		}
		if _lightsailGetRelationalDatabase {
			lightsail_GetRelationalDatabase(cfg, client)
			return
		}
		if _lightsailGetRelationalDatabaseBlueprints {
			lightsail_GetRelationalDatabaseBlueprints(cfg, client)
			return
		}
		if _lightsailGetRelationalDatabaseBundles {
			lightsail_GetRelationalDatabaseBundles(cfg, client)
			return
		}
		if _lightsailGetRelationalDatabaseEvents {
			lightsail_GetRelationalDatabaseEvents(cfg, client)
			return
		}
		if _lightsailGetRelationalDatabaseLogEvents {
			lightsail_GetRelationalDatabaseLogEvents(cfg, client)
			return
		}
		if _lightsailGetRelationalDatabaseLogStreams {
			lightsail_GetRelationalDatabaseLogStreams(cfg, client)
			return
		}
		if _lightsailGetRelationalDatabaseMasterUserPassword {
			lightsail_GetRelationalDatabaseMasterUserPassword(cfg, client)
			return
		}
		if _lightsailGetRelationalDatabaseMetricData {
			lightsail_GetRelationalDatabaseMetricData(cfg, client)
			return
		}
		if _lightsailGetRelationalDatabaseParameters {
			lightsail_GetRelationalDatabaseParameters(cfg, client)
			return
		}
		if _lightsailGetRelationalDatabaseSnapshot {
			lightsail_GetRelationalDatabaseSnapshot(cfg, client)
			return
		}
		if _lightsailGetRelationalDatabaseSnapshots {
			lightsail_GetRelationalDatabaseSnapshots(cfg, client)
			return
		}
		if _lightsailGetRelationalDatabases {
			lightsail_GetRelationalDatabases(cfg, client)
			return
		}
		if _lightsailGetSetupHistory {
			lightsail_GetSetupHistory(cfg, client)
			return
		}
		if _lightsailGetStaticIp {
			lightsail_GetStaticIp(cfg, client)
			return
		}
		if _lightsailGetStaticIps {
			lightsail_GetStaticIps(cfg, client)
			return
		}
		if _lightsailImportKeyPair {
			lightsail_ImportKeyPair(cfg, client)
			return
		}
		if _lightsailIsVpcPeered {
			lightsail_IsVpcPeered(cfg, client)
			return
		}
		if _lightsailOpenInstancePublicPorts {
			lightsail_OpenInstancePublicPorts(cfg, client)
			return
		}
		if _lightsailPeerVpc {
			lightsail_PeerVpc(cfg, client)
			return
		}
		if _lightsailPutAlarm {
			lightsail_PutAlarm(cfg, client)
			return
		}
		if _lightsailPutInstancePublicPorts {
			lightsail_PutInstancePublicPorts(cfg, client)
			return
		}
		if _lightsailRebootInstance {
			lightsail_RebootInstance(cfg, client)
			return
		}
		if _lightsailRebootRelationalDatabase {
			lightsail_RebootRelationalDatabase(cfg, client)
			return
		}
		if _lightsailRegisterContainerImage {
			lightsail_RegisterContainerImage(cfg, client)
			return
		}
		if _lightsailReleaseStaticIp {
			lightsail_ReleaseStaticIp(cfg, client)
			return
		}
		if _lightsailResetDistributionCache {
			lightsail_ResetDistributionCache(cfg, client)
			return
		}
		if _lightsailSendContactMethodVerification {
			lightsail_SendContactMethodVerification(cfg, client)
			return
		}
		if _lightsailSetIpAddressType {
			lightsail_SetIpAddressType(cfg, client)
			return
		}
		if _lightsailSetResourceAccessForBucket {
			lightsail_SetResourceAccessForBucket(cfg, client)
			return
		}
		if _lightsailSetupInstanceHttps {
			lightsail_SetupInstanceHttps(cfg, client)
			return
		}
		if _lightsailStartGUISession {
			lightsail_StartGUISession(cfg, client)
			return
		}
		if _lightsailStartInstance {
			lightsail_StartInstance(cfg, client)
			return
		}
		if _lightsailStartRelationalDatabase {
			lightsail_StartRelationalDatabase(cfg, client)
			return
		}
		if _lightsailStopGUISession {
			lightsail_StopGUISession(cfg, client)
			return
		}
		if _lightsailStopInstance {
			lightsail_StopInstance(cfg, client)
			return
		}
		if _lightsailStopRelationalDatabase {
			lightsail_StopRelationalDatabase(cfg, client)
			return
		}
		if _lightsailTagResource {
			lightsail_TagResource(cfg, client)
			return
		}
		if _lightsailTestAlarm {
			lightsail_TestAlarm(cfg, client)
			return
		}
		if _lightsailUnpeerVpc {
			lightsail_UnpeerVpc(cfg, client)
			return
		}
		if _lightsailUntagResource {
			lightsail_UntagResource(cfg, client)
			return
		}
		if _lightsailUpdateBucket {
			lightsail_UpdateBucket(cfg, client)
			return
		}
		if _lightsailUpdateBucketBundle {
			lightsail_UpdateBucketBundle(cfg, client)
			return
		}
		if _lightsailUpdateContainerService {
			lightsail_UpdateContainerService(cfg, client)
			return
		}
		if _lightsailUpdateDistribution {
			lightsail_UpdateDistribution(cfg, client)
			return
		}
		if _lightsailUpdateDistributionBundle {
			lightsail_UpdateDistributionBundle(cfg, client)
			return
		}
		if _lightsailUpdateDomainEntry {
			lightsail_UpdateDomainEntry(cfg, client)
			return
		}
		if _lightsailUpdateInstanceMetadataOptions {
			lightsail_UpdateInstanceMetadataOptions(cfg, client)
			return
		}
		if _lightsailUpdateLoadBalancerAttribute {
			lightsail_UpdateLoadBalancerAttribute(cfg, client)
			return
		}
		if _lightsailUpdateRelationalDatabase {
			lightsail_UpdateRelationalDatabase(cfg, client)
			return
		}
		if _lightsailUpdateRelationalDatabaseParameters {
			lightsail_UpdateRelationalDatabaseParameters(cfg, client)
			return
		}

	},
}

var (
	_lightsailAllocateStaticIp                        bool
	_lightsailAttachCertificateToDistribution         bool
	_lightsailAttachDisk                              bool
	_lightsailAttachInstancesToLoadBalancer           bool
	_lightsailAttachLoadBalancerTlsCertificate        bool
	_lightsailAttachStaticIp                          bool
	_lightsailCloseInstancePublicPorts                bool
	_lightsailCopySnapshot                            bool
	_lightsailCreateBucket                            bool
	_lightsailCreateBucketAccessKey                   bool
	_lightsailCreateCertificate                       bool
	_lightsailCreateCloudFormationStack               bool
	_lightsailCreateContactMethod                     bool
	_lightsailCreateContainerService                  bool
	_lightsailCreateContainerServiceDeployment        bool
	_lightsailCreateContainerServiceRegistryLogin     bool
	_lightsailCreateDisk                              bool
	_lightsailCreateDiskFromSnapshot                  bool
	_lightsailCreateDiskSnapshot                      bool
	_lightsailCreateDistribution                      bool
	_lightsailCreateDomain                            bool
	_lightsailCreateDomainEntry                       bool
	_lightsailCreateGUISessionAccessDetails           bool
	_lightsailCreateInstanceSnapshot                  bool
	_lightsailCreateInstances                         bool
	_lightsailCreateInstancesFromSnapshot             bool
	_lightsailCreateKeyPair                           bool
	_lightsailCreateLoadBalancer                      bool
	_lightsailCreateLoadBalancerTlsCertificate        bool
	_lightsailCreateRelationalDatabase                bool
	_lightsailCreateRelationalDatabaseFromSnapshot    bool
	_lightsailCreateRelationalDatabaseSnapshot        bool
	_lightsailDeleteAlarm                             bool
	_lightsailDeleteAutoSnapshot                      bool
	_lightsailDeleteBucket                            bool
	_lightsailDeleteBucketAccessKey                   bool
	_lightsailDeleteCertificate                       bool
	_lightsailDeleteContactMethod                     bool
	_lightsailDeleteContainerImage                    bool
	_lightsailDeleteContainerService                  bool
	_lightsailDeleteDisk                              bool
	_lightsailDeleteDiskSnapshot                      bool
	_lightsailDeleteDistribution                      bool
	_lightsailDeleteDomain                            bool
	_lightsailDeleteDomainEntry                       bool
	_lightsailDeleteInstance                          bool
	_lightsailDeleteInstanceSnapshot                  bool
	_lightsailDeleteKeyPair                           bool
	_lightsailDeleteKnownHostKeys                     bool
	_lightsailDeleteLoadBalancer                      bool
	_lightsailDeleteLoadBalancerTlsCertificate        bool
	_lightsailDeleteRelationalDatabase                bool
	_lightsailDeleteRelationalDatabaseSnapshot        bool
	_lightsailDetachCertificateFromDistribution       bool
	_lightsailDetachDisk                              bool
	_lightsailDetachInstancesFromLoadBalancer         bool
	_lightsailDetachStaticIp                          bool
	_lightsailDisableAddOn                            bool
	_lightsailDownloadDefaultKeyPair                  bool
	_lightsailEnableAddOn                             bool
	_lightsailExportSnapshot                          bool
	_lightsailGetActiveNames                          bool
	_lightsailGetAlarms                               bool
	_lightsailGetAutoSnapshots                        bool
	_lightsailGetBlueprints                           bool
	_lightsailGetBucketAccessKeys                     bool
	_lightsailGetBucketBundles                        bool
	_lightsailGetBucketMetricData                     bool
	_lightsailGetBuckets                              bool
	_lightsailGetBundles                              bool
	_lightsailGetCertificates                         bool
	_lightsailGetCloudFormationStackRecords           bool
	_lightsailGetContactMethods                       bool
	_lightsailGetContainerAPIMetadata                 bool
	_lightsailGetContainerImages                      bool
	_lightsailGetContainerLog                         bool
	_lightsailGetContainerServiceDeployments          bool
	_lightsailGetContainerServiceMetricData           bool
	_lightsailGetContainerServicePowers               bool
	_lightsailGetContainerServices                    bool
	_lightsailGetCostEstimate                         bool
	_lightsailGetDisk                                 bool
	_lightsailGetDiskSnapshot                         bool
	_lightsailGetDiskSnapshots                        bool
	_lightsailGetDisks                                bool
	_lightsailGetDistributionBundles                  bool
	_lightsailGetDistributionLatestCacheReset         bool
	_lightsailGetDistributionMetricData               bool
	_lightsailGetDistributions                        bool
	_lightsailGetDomain                               bool
	_lightsailGetDomains                              bool
	_lightsailGetExportSnapshotRecords                bool
	_lightsailGetInstance                             bool
	_lightsailGetInstanceAccessDetails                bool
	_lightsailGetInstanceMetricData                   bool
	_lightsailGetInstancePortStates                   bool
	_lightsailGetInstanceSnapshot                     bool
	_lightsailGetInstanceSnapshots                    bool
	_lightsailGetInstanceState                        bool
	_lightsailGetInstances                            bool
	_lightsailGetKeyPair                              bool
	_lightsailGetKeyPairs                             bool
	_lightsailGetLoadBalancer                         bool
	_lightsailGetLoadBalancerMetricData               bool
	_lightsailGetLoadBalancerTlsCertificates          bool
	_lightsailGetLoadBalancerTlsPolicies              bool
	_lightsailGetLoadBalancers                        bool
	_lightsailGetOperation                            bool
	_lightsailGetOperations                           bool
	_lightsailGetOperationsForResource                bool
	_lightsailGetRegions                              bool
	_lightsailGetRelationalDatabase                   bool
	_lightsailGetRelationalDatabaseBlueprints         bool
	_lightsailGetRelationalDatabaseBundles            bool
	_lightsailGetRelationalDatabaseEvents             bool
	_lightsailGetRelationalDatabaseLogEvents          bool
	_lightsailGetRelationalDatabaseLogStreams         bool
	_lightsailGetRelationalDatabaseMasterUserPassword bool
	_lightsailGetRelationalDatabaseMetricData         bool
	_lightsailGetRelationalDatabaseParameters         bool
	_lightsailGetRelationalDatabaseSnapshot           bool
	_lightsailGetRelationalDatabaseSnapshots          bool
	_lightsailGetRelationalDatabases                  bool
	_lightsailGetSetupHistory                         bool
	_lightsailGetStaticIp                             bool
	_lightsailGetStaticIps                            bool
	_lightsailImportKeyPair                           bool
	_lightsailIsVpcPeered                             bool
	_lightsailOpenInstancePublicPorts                 bool
	_lightsailPeerVpc                                 bool
	_lightsailPutAlarm                                bool
	_lightsailPutInstancePublicPorts                  bool
	_lightsailRebootInstance                          bool
	_lightsailRebootRelationalDatabase                bool
	_lightsailRegisterContainerImage                  bool
	_lightsailReleaseStaticIp                         bool
	_lightsailResetDistributionCache                  bool
	_lightsailSendContactMethodVerification           bool
	_lightsailSetIpAddressType                        bool
	_lightsailSetResourceAccessForBucket              bool
	_lightsailSetupInstanceHttps                      bool
	_lightsailStartGUISession                         bool
	_lightsailStartInstance                           bool
	_lightsailStartRelationalDatabase                 bool
	_lightsailStopGUISession                          bool
	_lightsailStopInstance                            bool
	_lightsailStopRelationalDatabase                  bool
	_lightsailTagResource                             bool
	_lightsailTestAlarm                               bool
	_lightsailUnpeerVpc                               bool
	_lightsailUntagResource                           bool
	_lightsailUpdateBucket                            bool
	_lightsailUpdateBucketBundle                      bool
	_lightsailUpdateContainerService                  bool
	_lightsailUpdateDistribution                      bool
	_lightsailUpdateDistributionBundle                bool
	_lightsailUpdateDomainEntry                       bool
	_lightsailUpdateInstanceMetadataOptions           bool
	_lightsailUpdateLoadBalancerAttribute             bool
	_lightsailUpdateRelationalDatabase                bool
	_lightsailUpdateRelationalDatabaseParameters      bool

	_lightsailAcceptBundleUpdate                         string
	_lightsailAccess                                     string
	_lightsailAccessKeyId                                string
	_lightsailAccessLogConfig                            string
	_lightsailAccessRules                                string
	_lightsailAddOnRequest                               string
	_lightsailAddOnType                                  string
	_lightsailAddOns                                     string
	_lightsailAlarmName                                  string
	_lightsailAppCategory                                string
	_lightsailApplyImmediately                           string
	_lightsailAttachedDiskMapping                        string
	_lightsailAttributeName                              string
	_lightsailAttributeValue                             string
	_lightsailAutoMounting                               string
	_lightsailAvailabilityZone                           string
	_lightsailBlueprintId                                string
	_lightsailBucketName                                 string
	_lightsailBundleId                                   string
	_lightsailCaCertificateIdentifier                    string
	_lightsailCacheBehaviorSettings                      string
	_lightsailCacheBehaviors                             string
	_lightsailCertificateAlternativeNames                []string
	_lightsailCertificateDomainName                      string
	_lightsailCertificateName                            string
	_lightsailCertificateProvider                        string
	_lightsailCertificateStatuses                        string
	_lightsailComparisonOperator                         string
	_lightsailContactEndpoint                            string
	_lightsailContactProtocols                           string
	_lightsailContainerName                              string
	_lightsailContainers                                 string
	_lightsailCors                                       string
	_lightsailCustomImageName                            string
	_lightsailDatapointsToAlarm                          string
	_lightsailDate                                       string
	_lightsailDefaultCacheBehavior                       string
	_lightsailDeployment                                 string
	_lightsailDigest                                     string
	_lightsailDisableBackupRetention                     string
	_lightsailDiskName                                   string
	_lightsailDiskPath                                   string
	_lightsailDiskSnapshotName                           string
	_lightsailDistributionName                           string
	_lightsailDomainEntry                                string
	_lightsailDomainName                                 string
	_lightsailDomainNames                                []string
	_lightsailDurationInMinutes                          string
	_lightsailEmailAddress                               string
	_lightsailEnableBackupRetention                      string
	_lightsailEnableObjectVersioning                     string
	_lightsailEndTime                                    string
	_lightsailEvaluationPeriods                          string
	_lightsailExpectedFingerprint                        string
	_lightsailFilterPattern                              string
	_lightsailFinalRelationalDatabaseSnapshotName        string
	_lightsailForce                                      string
	_lightsailForceDelete                                string
	_lightsailForceDeleteAddOns                          string
	_lightsailHealthCheckPath                            string
	_lightsailHttpEndpoint                               string
	_lightsailHttpProtocolIpv6                           string
	_lightsailHttpPutResponseHopLimit                    string
	_lightsailHttpTokens                                 string
	_lightsailImage                                      string
	_lightsailIncludeAvailabilityZones                   string
	_lightsailIncludeCertificateDetails                  string
	_lightsailIncludeConnectedResources                  string
	_lightsailIncludeCors                                string
	_lightsailIncludeDefaultKeyPair                      string
	_lightsailIncludeInactive                            string
	_lightsailIncludeRelationalDatabaseAvailabilityZones string
	_lightsailInstanceName                               string
	_lightsailInstanceNames                              []string
	_lightsailInstancePort                               string
	_lightsailInstanceSnapshotName                       string
	_lightsailInstances                                  string
	_lightsailIpAddressType                              string
	_lightsailIsDisabled                                 string
	_lightsailIsEnabled                                  string
	_lightsailKeyPairName                                string
	_lightsailLabel                                      string
	_lightsailLoadBalancerName                           string
	_lightsailLogStreamName                              string
	_lightsailMasterDatabaseName                         string
	_lightsailMasterUserPassword                         string
	_lightsailMasterUsername                             string
	_lightsailMetricName                                 string
	_lightsailMonitoredResourceName                      string
	_lightsailNotificationEnabled                        string
	_lightsailNotificationTriggers                       string
	_lightsailOperationId                                string
	_lightsailOrigin                                     string
	_lightsailPageToken                                  string
	_lightsailParameters                                 string
	_lightsailPasswordVersion                            string
	_lightsailPeriod                                     string
	_lightsailPortInfo                                   string
	_lightsailPortInfos                                  string
	_lightsailPower                                      string
	_lightsailPreferredBackupWindow                      string
	_lightsailPreferredMaintenanceWindow                 string
	_lightsailPrivateRegistryAccess                      string
	_lightsailProtocol                                   string
	_lightsailProtocols                                  string
	_lightsailPublicDomainNames                          string
	_lightsailPublicEndpoint                             string
	_lightsailPublicKeyBase64                            string
	_lightsailPubliclyAccessible                         string
	_lightsailReadonlyAccessAccounts                     []string
	_lightsailRelationalDatabaseBlueprintId              string
	_lightsailRelationalDatabaseBundleId                 string
	_lightsailRelationalDatabaseName                     string
	_lightsailRelationalDatabaseSnapshotName             string
	_lightsailResourceArn                                string
	_lightsailResourceName                               string
	_lightsailResourceType                               string
	_lightsailRestoreDate                                string
	_lightsailRestoreTime                                string
	_lightsailRotateMasterUserPassword                   string
	_lightsailScale                                      string
	_lightsailServiceName                                string
	_lightsailSizeInGb                                   string
	_lightsailSkipFinalSnapshot                          string
	_lightsailSourceDiskName                             string
	_lightsailSourceInstanceName                         string
	_lightsailSourceRegion                               string
	_lightsailSourceRelationalDatabaseName               string
	_lightsailSourceResourceName                         string
	_lightsailSourceSnapshotName                         string
	_lightsailStartFromHead                              string
	_lightsailStartTime                                  string
	_lightsailState                                      string
	_lightsailStaticIpName                               string
	_lightsailStatistics                                 string
	_lightsailSubjectAlternativeNames                    []string
	_lightsailTagKeys                                    []string
	_lightsailTags                                       string
	_lightsailTargetSnapshotName                         string
	_lightsailThreshold                                  string
	_lightsailTlsPolicyName                              string
	_lightsailTreatMissingData                           string
	_lightsailUnit                                       string
	_lightsailUseDefaultCertificate                      string
	_lightsailUseLatestRestorableAutoSnapshot            string
	_lightsailUseLatestRestorableTime                    string
	_lightsailUserData                                   string
	_lightsailVersioning                                 string
	_lightsailViewerMinimumTlsProtocolVersion            string
)

// Allocates a static IP address.
func lightsail_AllocateStaticIp(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.AllocateStaticIpInput{
		// StaticIpName: *string, // Required
	}

	if len(_lightsailStaticIpName) > 0 {
		input.StaticIpName = aws.String(_lightsailStaticIpName)
	}

	if resp, err := client.AllocateStaticIp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches an SSL/TLS certificate to your Amazon Lightsail content delivery
// network (CDN) distribution.
//
// After the certificate is attached, your distribution accepts HTTPS traffic for
// all of the domains that are associated with the certificate.
//
// Use the CreateCertificate action to create a certificate that you can attach to
// your distribution.
//
// Only certificates created in the us-east-1 Amazon Web Services Region can be
// attached to Lightsail distributions. Lightsail distributions are global
// resources that can reference an origin in any Amazon Web Services Region, and
// distribute its content globally. However, all distributions are located in the
// us-east-1 Region.
func lightsail_AttachCertificateToDistribution(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.AttachCertificateToDistributionInput{
		// CertificateName: *string, // Required
		// DistributionName: *string, // Required
	}

	if len(_lightsailCertificateName) > 0 {
		input.CertificateName = aws.String(_lightsailCertificateName)
	}
	if len(_lightsailDistributionName) > 0 {
		input.DistributionName = aws.String(_lightsailDistributionName)
	}

	if resp, err := client.AttachCertificateToDistribution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches a block storage disk to a running or stopped Lightsail instance and
// exposes it to the instance with the specified disk name.
//
// The attach disk operation supports tag-based access control via resource tags
// applied to the resource identified by disk name . For more information, see the [Amazon Lightsail Developer Guide]
// .
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_AttachDisk(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.AttachDiskInput{
		// DiskName: *string, // Required
		// DiskPath: *string, // Required
		// InstanceName: *string, // Required
	}

	if len(_lightsailDiskName) > 0 {
		input.DiskName = aws.String(_lightsailDiskName)
	}
	if len(_lightsailDiskPath) > 0 {
		input.DiskPath = aws.String(_lightsailDiskPath)
	}
	if len(_lightsailInstanceName) > 0 {
		input.InstanceName = aws.String(_lightsailInstanceName)
	}
	if len(_lightsailAutoMounting) > 0 {
		if err := assignInputField(input, "AutoMounting", _lightsailAutoMounting); err != nil {
			log.Errorf("invalid --auto-mounting: %s", err.Error())
			return
		}
	}

	if resp, err := client.AttachDisk(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches one or more Lightsail instances to a load balancer.
// After some time, the instances are attached to the load balancer and the health
// check status is available.
//
// The attach instances to load balancer operation supports tag-based access
// control via resource tags applied to the resource identified by load balancer
// name . For more information, see the [Lightsail Developer Guide].
//
// [Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_AttachInstancesToLoadBalancer(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.AttachInstancesToLoadBalancerInput{
		// InstanceNames: []string, // Required
		// LoadBalancerName: *string, // Required
	}

	if len(_lightsailInstanceNames) > 0 {
		input.InstanceNames = append([]string(nil), _lightsailInstanceNames...)
	}
	if len(_lightsailLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_lightsailLoadBalancerName)
	}

	if resp, err := client.AttachInstancesToLoadBalancer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches a Transport Layer Security (TLS) certificate to your load balancer.
// TLS is just an updated, more secure version of Secure Socket Layer (SSL).
//
// Once you create and validate your certificate, you can attach it to your load
// balancer. You can also use this API to rotate the certificates on your account.
// Use the AttachLoadBalancerTlsCertificate action with the non-attached
// certificate, and it will replace the existing one and become the attached
// certificate.
//
// The AttachLoadBalancerTlsCertificate operation supports tag-based access
// control via resource tags applied to the resource identified by load balancer
// name . For more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_AttachLoadBalancerTlsCertificate(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.AttachLoadBalancerTlsCertificateInput{
		// CertificateName: *string, // Required
		// LoadBalancerName: *string, // Required
	}

	if len(_lightsailCertificateName) > 0 {
		input.CertificateName = aws.String(_lightsailCertificateName)
	}
	if len(_lightsailLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_lightsailLoadBalancerName)
	}

	if resp, err := client.AttachLoadBalancerTlsCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches a static IP address to a specific Amazon Lightsail instance.
func lightsail_AttachStaticIp(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.AttachStaticIpInput{
		// InstanceName: *string, // Required
		// StaticIpName: *string, // Required
	}

	if len(_lightsailInstanceName) > 0 {
		input.InstanceName = aws.String(_lightsailInstanceName)
	}
	if len(_lightsailStaticIpName) > 0 {
		input.StaticIpName = aws.String(_lightsailStaticIpName)
	}

	if resp, err := client.AttachStaticIp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Closes ports for a specific Amazon Lightsail instance.
// The CloseInstancePublicPorts action supports tag-based access control via
// resource tags applied to the resource identified by instanceName . For more
// information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_CloseInstancePublicPorts(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CloseInstancePublicPortsInput{
		// InstanceName: *string, // Required
		// PortInfo: *types.PortInfo, // Required
	}

	if len(_lightsailInstanceName) > 0 {
		input.InstanceName = aws.String(_lightsailInstanceName)
	}
	if len(_lightsailPortInfo) > 0 {
		if err := assignInputField(input, "PortInfo", _lightsailPortInfo); err != nil {
			log.Errorf("invalid --port-info: %s", err.Error())
			return
		}
	}

	if resp, err := client.CloseInstancePublicPorts(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Copies a manual snapshot of an instance or disk as another manual snapshot, or
// copies an automatic snapshot of an instance or disk as a manual snapshot. This
// operation can also be used to copy a manual or automatic snapshot of an instance
// or a disk from one Amazon Web Services Region to another in Amazon Lightsail.
//
// When copying a manual snapshot, be sure to define the source region , source
// snapshot name , and target snapshot name parameters.
//
// When copying an automatic snapshot, be sure to define the source region ,
// source resource name , target snapshot name , and either the restore date or
// the use latest restorable auto snapshot parameters.
func lightsail_CopySnapshot(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CopySnapshotInput{
		// SourceRegion: types.RegionName, // Required
		// TargetSnapshotName: *string, // Required
	}

	if len(_lightsailSourceRegion) > 0 {
		if err := assignInputField(input, "SourceRegion", _lightsailSourceRegion); err != nil {
			log.Errorf("invalid --source-region: %s", err.Error())
			return
		}
	}
	if len(_lightsailTargetSnapshotName) > 0 {
		input.TargetSnapshotName = aws.String(_lightsailTargetSnapshotName)
	}
	if len(_lightsailRestoreDate) > 0 {
		input.RestoreDate = aws.String(_lightsailRestoreDate)
	}
	if len(_lightsailSourceResourceName) > 0 {
		input.SourceResourceName = aws.String(_lightsailSourceResourceName)
	}
	if len(_lightsailSourceSnapshotName) > 0 {
		input.SourceSnapshotName = aws.String(_lightsailSourceSnapshotName)
	}
	if len(_lightsailUseLatestRestorableAutoSnapshot) > 0 {
		if err := assignInputField(input, "UseLatestRestorableAutoSnapshot", _lightsailUseLatestRestorableAutoSnapshot); err != nil {
			log.Errorf("invalid --use-latest-restorable-auto-snapshot: %s", err.Error())
			return
		}
	}

	if resp, err := client.CopySnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Lightsail bucket.
// A bucket is a cloud storage resource available in the Lightsail object storage
// service. Use buckets to store objects such as data and its descriptive metadata.
// For more information about buckets, see [Buckets in Amazon Lightsail]in the Amazon Lightsail Developer Guide.
//
// [Buckets in Amazon Lightsail]: https://docs.aws.amazon.com/lightsail/latest/userguide/buckets-in-amazon-lightsail
func lightsail_CreateBucket(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CreateBucketInput{
		// BucketName: *string, // Required
		// BundleId: *string, // Required
	}

	if len(_lightsailBucketName) > 0 {
		input.BucketName = aws.String(_lightsailBucketName)
	}
	if len(_lightsailBundleId) > 0 {
		input.BundleId = aws.String(_lightsailBundleId)
	}
	if len(_lightsailEnableObjectVersioning) > 0 {
		if err := assignInputField(input, "EnableObjectVersioning", _lightsailEnableObjectVersioning); err != nil {
			log.Errorf("invalid --enable-object-versioning: %s", err.Error())
			return
		}
	}
	if len(_lightsailTags) > 0 {
		if err := assignInputField(input, "Tags", _lightsailTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBucket(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new access key for the specified Amazon Lightsail bucket. Access keys
// consist of an access key ID and corresponding secret access key.
//
// Access keys grant full programmatic access to the specified bucket and its
// objects. You can have a maximum of two access keys per bucket. Use the [GetBucketAccessKeys]action
// to get a list of current access keys for a specific bucket. For more information
// about access keys, see [Creating access keys for a bucket in Amazon Lightsail]in the Amazon Lightsail Developer Guide.
//
// The secretAccessKey value is returned only in response to the
// CreateBucketAccessKey action. You can get a secret access key only when you
// first create an access key; you cannot get the secret access key later. If you
// lose the secret access key, you must create a new access key.
//
// [Creating access keys for a bucket in Amazon Lightsail]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-creating-bucket-access-keys
// [GetBucketAccessKeys]: https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_GetBucketAccessKeys.html
func lightsail_CreateBucketAccessKey(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CreateBucketAccessKeyInput{
		// BucketName: *string, // Required
	}

	if len(_lightsailBucketName) > 0 {
		input.BucketName = aws.String(_lightsailBucketName)
	}

	if resp, err := client.CreateBucketAccessKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an SSL/TLS certificate for an Amazon Lightsail content delivery network
// (CDN) distribution and a container service.
//
// After the certificate is valid, use the AttachCertificateToDistribution action
// to use the certificate and its domains with your distribution. Or use the
// UpdateContainerService action to use the certificate and its domains with your
// container service.
//
// Only certificates created in the us-east-1 Amazon Web Services Region can be
// attached to Lightsail distributions. Lightsail distributions are global
// resources that can reference an origin in any Amazon Web Services Region, and
// distribute its content globally. However, all distributions are located in the
// us-east-1 Region.
func lightsail_CreateCertificate(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CreateCertificateInput{
		// CertificateName: *string, // Required
		// DomainName: *string, // Required
	}

	if len(_lightsailCertificateName) > 0 {
		input.CertificateName = aws.String(_lightsailCertificateName)
	}
	if len(_lightsailDomainName) > 0 {
		input.DomainName = aws.String(_lightsailDomainName)
	}
	if len(_lightsailSubjectAlternativeNames) > 0 {
		input.SubjectAlternativeNames = append([]string(nil), _lightsailSubjectAlternativeNames...)
	}
	if len(_lightsailTags) > 0 {
		if err := assignInputField(input, "Tags", _lightsailTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an AWS CloudFormation stack, which creates a new Amazon EC2 instance
// from an exported Amazon Lightsail snapshot. This operation results in a
// CloudFormation stack record that can be used to track the AWS CloudFormation
// stack created. Use the get cloud formation stack records operation to get a
// list of the CloudFormation stacks created.
//
// Wait until after your new Amazon EC2 instance is created before running the
// create cloud formation stack operation again with the same export snapshot
// record.
func lightsail_CreateCloudFormationStack(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CreateCloudFormationStackInput{
		// Instances: []types.InstanceEntry, // Required
	}

	if len(_lightsailInstances) > 0 {
		if err := assignInputField(input, "Instances", _lightsailInstances); err != nil {
			log.Errorf("invalid --instances: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCloudFormationStack(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an email or SMS text message contact method.
// A contact method is used to send you notifications about your Amazon Lightsail
// resources. You can add one email address and one mobile phone number contact
// method in each Amazon Web Services Region. However, SMS text messaging is not
// supported in some Amazon Web Services Regions, and SMS text messages cannot be
// sent to some countries/regions. For more information, see [Notifications in Amazon Lightsail].
//
// [Notifications in Amazon Lightsail]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-notifications
func lightsail_CreateContactMethod(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CreateContactMethodInput{
		// ContactEndpoint: *string, // Required
		// Protocol: types.ContactProtocol, // Required
	}

	if len(_lightsailContactEndpoint) > 0 {
		input.ContactEndpoint = aws.String(_lightsailContactEndpoint)
	}
	if len(_lightsailProtocol) > 0 {
		if err := assignInputField(input, "Protocol", _lightsailProtocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateContactMethod(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Lightsail container service.
// A Lightsail container service is a compute resource to which you can deploy
// containers. For more information, see [Container services in Amazon Lightsail]in the Lightsail Dev Guide.
//
// [Container services in Amazon Lightsail]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-container-services
func lightsail_CreateContainerService(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CreateContainerServiceInput{
		// Power: types.ContainerServicePowerName, // Required
		// Scale: *int32, // Required
		// ServiceName: *string, // Required
	}

	if len(_lightsailPower) > 0 {
		if err := assignInputField(input, "Power", _lightsailPower); err != nil {
			log.Errorf("invalid --power: %s", err.Error())
			return
		}
	}
	if len(_lightsailScale) > 0 {
		if err := assignInputField(input, "Scale", _lightsailScale); err != nil {
			log.Errorf("invalid --scale: %s", err.Error())
			return
		}
	}
	if len(_lightsailServiceName) > 0 {
		input.ServiceName = aws.String(_lightsailServiceName)
	}
	if len(_lightsailDeployment) > 0 {
		if err := assignInputField(input, "Deployment", _lightsailDeployment); err != nil {
			log.Errorf("invalid --deployment: %s", err.Error())
			return
		}
	}
	if len(_lightsailPrivateRegistryAccess) > 0 {
		if err := assignInputField(input, "PrivateRegistryAccess", _lightsailPrivateRegistryAccess); err != nil {
			log.Errorf("invalid --private-registry-access: %s", err.Error())
			return
		}
	}
	if len(_lightsailPublicDomainNames) > 0 {
		if err := assignInputField(input, "PublicDomainNames", _lightsailPublicDomainNames); err != nil {
			log.Errorf("invalid --public-domain-names: %s", err.Error())
			return
		}
	}
	if len(_lightsailTags) > 0 {
		if err := assignInputField(input, "Tags", _lightsailTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateContainerService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a deployment for your Amazon Lightsail container service.
// A deployment specifies the containers that will be launched on the container
// service and their settings, such as the ports to open, the environment variables
// to apply, and the launch command to run. It also specifies the container that
// will serve as the public endpoint of the deployment and its settings, such as
// the HTTP or HTTPS port to use, and the health check configuration.
//
// You can deploy containers to your container service using container images from
// a public registry such as Amazon ECR Public, or from your local machine. For
// more information, see [Creating container images for your Amazon Lightsail container services]in the Amazon Lightsail Developer Guide.
//
// [Creating container images for your Amazon Lightsail container services]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-creating-container-images
func lightsail_CreateContainerServiceDeployment(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CreateContainerServiceDeploymentInput{
		// ServiceName: *string, // Required
	}

	if len(_lightsailServiceName) > 0 {
		input.ServiceName = aws.String(_lightsailServiceName)
	}
	if len(_lightsailContainers) > 0 {
		if err := assignInputField(input, "Containers", _lightsailContainers); err != nil {
			log.Errorf("invalid --containers: %s", err.Error())
			return
		}
	}
	if len(_lightsailPublicEndpoint) > 0 {
		if err := assignInputField(input, "PublicEndpoint", _lightsailPublicEndpoint); err != nil {
			log.Errorf("invalid --public-endpoint: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateContainerServiceDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a temporary set of log in credentials that you can use to log in to the
// Docker process on your local machine. After you're logged in, you can use the
// native Docker commands to push your local container images to the container
// image registry of your Amazon Lightsail account so that you can use them with
// your Lightsail container service. The log in credentials expire 12 hours after
// they are created, at which point you will need to create a new set of log in
// credentials.
//
// You can only push container images to the container service registry of your
// Lightsail account. You cannot pull container images or perform any other
// container image management actions on the container service registry.
//
// After you push your container images to the container image registry of your
// Lightsail account, use the RegisterContainerImage action to register the pushed
// images to a specific Lightsail container service.
//
// This action is not required if you install and use the Lightsail Control
// (lightsailctl) plugin to push container images to your Lightsail container
// service. For more information, see [Pushing and managing container images on your Amazon Lightsail container services]in the Amazon Lightsail Developer Guide.
//
// [Pushing and managing container images on your Amazon Lightsail container services]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-pushing-container-images
func lightsail_CreateContainerServiceRegistryLogin(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CreateContainerServiceRegistryLoginInput{}

	if resp, err := client.CreateContainerServiceRegistryLogin(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a block storage disk that can be attached to an Amazon Lightsail
// instance in the same Availability Zone ( us-east-2a ).
//
// The create disk operation supports tag-based access control via request tags.
// For more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_CreateDisk(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CreateDiskInput{
		// AvailabilityZone: *string, // Required
		// DiskName: *string, // Required
		// SizeInGb: *int32, // Required
	}

	if len(_lightsailAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_lightsailAvailabilityZone)
	}
	if len(_lightsailDiskName) > 0 {
		input.DiskName = aws.String(_lightsailDiskName)
	}
	if len(_lightsailSizeInGb) > 0 {
		if err := assignInputField(input, "SizeInGb", _lightsailSizeInGb); err != nil {
			log.Errorf("invalid --size-in-gb: %s", err.Error())
			return
		}
	}
	if len(_lightsailAddOns) > 0 {
		if err := assignInputField(input, "AddOns", _lightsailAddOns); err != nil {
			log.Errorf("invalid --add-ons: %s", err.Error())
			return
		}
	}
	if len(_lightsailTags) > 0 {
		if err := assignInputField(input, "Tags", _lightsailTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDisk(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a block storage disk from a manual or automatic snapshot of a disk. The
// resulting disk can be attached to an Amazon Lightsail instance in the same
// Availability Zone ( us-east-2a ).
//
// The create disk from snapshot operation supports tag-based access control via
// request tags and resource tags applied to the resource identified by disk
// snapshot name . For more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_CreateDiskFromSnapshot(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CreateDiskFromSnapshotInput{
		// AvailabilityZone: *string, // Required
		// DiskName: *string, // Required
		// SizeInGb: *int32, // Required
	}

	if len(_lightsailAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_lightsailAvailabilityZone)
	}
	if len(_lightsailDiskName) > 0 {
		input.DiskName = aws.String(_lightsailDiskName)
	}
	if len(_lightsailSizeInGb) > 0 {
		if err := assignInputField(input, "SizeInGb", _lightsailSizeInGb); err != nil {
			log.Errorf("invalid --size-in-gb: %s", err.Error())
			return
		}
	}
	if len(_lightsailAddOns) > 0 {
		if err := assignInputField(input, "AddOns", _lightsailAddOns); err != nil {
			log.Errorf("invalid --add-ons: %s", err.Error())
			return
		}
	}
	if len(_lightsailDiskSnapshotName) > 0 {
		input.DiskSnapshotName = aws.String(_lightsailDiskSnapshotName)
	}
	if len(_lightsailRestoreDate) > 0 {
		input.RestoreDate = aws.String(_lightsailRestoreDate)
	}
	if len(_lightsailSourceDiskName) > 0 {
		input.SourceDiskName = aws.String(_lightsailSourceDiskName)
	}
	if len(_lightsailTags) > 0 {
		if err := assignInputField(input, "Tags", _lightsailTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_lightsailUseLatestRestorableAutoSnapshot) > 0 {
		if err := assignInputField(input, "UseLatestRestorableAutoSnapshot", _lightsailUseLatestRestorableAutoSnapshot); err != nil {
			log.Errorf("invalid --use-latest-restorable-auto-snapshot: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDiskFromSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a snapshot of a block storage disk. You can use snapshots for backups,
// to make copies of disks, and to save data before shutting down a Lightsail
// instance.
//
// You can take a snapshot of an attached disk that is in use; however, snapshots
// only capture data that has been written to your disk at the time the snapshot
// command is issued. This may exclude any data that has been cached by any
// applications or the operating system. If you can pause any file systems on the
// disk long enough to take a snapshot, your snapshot should be complete.
// Nevertheless, if you cannot pause all file writes to the disk, you should
// unmount the disk from within the Lightsail instance, issue the create disk
// snapshot command, and then remount the disk to ensure a consistent and complete
// snapshot. You may remount and use your disk while the snapshot status is
// pending.
//
// You can also use this operation to create a snapshot of an instance's system
// volume. You might want to do this, for example, to recover data from the system
// volume of a botched instance or to create a backup of the system volume like you
// would for a block storage disk. To create a snapshot of a system volume, just
// define the instance name parameter when issuing the snapshot command, and a
// snapshot of the defined instance's system volume will be created. After the
// snapshot is available, you can create a block storage disk from the snapshot and
// attach it to a running instance to access the data on the disk.
//
// The create disk snapshot operation supports tag-based access control via
// request tags. For more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_CreateDiskSnapshot(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CreateDiskSnapshotInput{
		// DiskSnapshotName: *string, // Required
	}

	if len(_lightsailDiskSnapshotName) > 0 {
		input.DiskSnapshotName = aws.String(_lightsailDiskSnapshotName)
	}
	if len(_lightsailDiskName) > 0 {
		input.DiskName = aws.String(_lightsailDiskName)
	}
	if len(_lightsailInstanceName) > 0 {
		input.InstanceName = aws.String(_lightsailInstanceName)
	}
	if len(_lightsailTags) > 0 {
		if err := assignInputField(input, "Tags", _lightsailTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDiskSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Lightsail content delivery network (CDN) distribution.
// A distribution is a globally distributed network of caching servers that
// improve the performance of your website or web application hosted on a Lightsail
// instance. For more information, see [Content delivery networks in Amazon Lightsail].
//
// [Content delivery networks in Amazon Lightsail]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-content-delivery-network-distributions
func lightsail_CreateDistribution(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CreateDistributionInput{
		// BundleId: *string, // Required
		// DefaultCacheBehavior: *types.CacheBehavior, // Required
		// DistributionName: *string, // Required
		// Origin: *types.InputOrigin, // Required
	}

	if len(_lightsailBundleId) > 0 {
		input.BundleId = aws.String(_lightsailBundleId)
	}
	if len(_lightsailDefaultCacheBehavior) > 0 {
		if err := assignInputField(input, "DefaultCacheBehavior", _lightsailDefaultCacheBehavior); err != nil {
			log.Errorf("invalid --default-cache-behavior: %s", err.Error())
			return
		}
	}
	if len(_lightsailDistributionName) > 0 {
		input.DistributionName = aws.String(_lightsailDistributionName)
	}
	if len(_lightsailOrigin) > 0 {
		if err := assignInputField(input, "Origin", _lightsailOrigin); err != nil {
			log.Errorf("invalid --origin: %s", err.Error())
			return
		}
	}
	if len(_lightsailCacheBehaviorSettings) > 0 {
		if err := assignInputField(input, "CacheBehaviorSettings", _lightsailCacheBehaviorSettings); err != nil {
			log.Errorf("invalid --cache-behavior-settings: %s", err.Error())
			return
		}
	}
	if len(_lightsailCacheBehaviors) > 0 {
		if err := assignInputField(input, "CacheBehaviors", _lightsailCacheBehaviors); err != nil {
			log.Errorf("invalid --cache-behaviors: %s", err.Error())
			return
		}
	}
	if len(_lightsailCertificateName) > 0 {
		input.CertificateName = aws.String(_lightsailCertificateName)
	}
	if len(_lightsailIpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _lightsailIpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_lightsailTags) > 0 {
		if err := assignInputField(input, "Tags", _lightsailTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_lightsailViewerMinimumTlsProtocolVersion) > 0 {
		if err := assignInputField(input, "ViewerMinimumTlsProtocolVersion", _lightsailViewerMinimumTlsProtocolVersion); err != nil {
			log.Errorf("invalid --viewer-minimum-tls-protocol-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDistribution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a domain resource for the specified domain (example.com).
// The create domain operation supports tag-based access control via request tags.
// For more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_CreateDomain(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CreateDomainInput{
		// DomainName: *string, // Required
	}

	if len(_lightsailDomainName) > 0 {
		input.DomainName = aws.String(_lightsailDomainName)
	}
	if len(_lightsailTags) > 0 {
		if err := assignInputField(input, "Tags", _lightsailTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates one of the following domain name system (DNS) records in a domain DNS
// zone: Address (A), canonical name (CNAME), mail exchanger (MX), name server
// (NS), start of authority (SOA), service locator (SRV), or text (TXT).
//
// The create domain entry operation supports tag-based access control via
// resource tags applied to the resource identified by domain name . For more
// information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_CreateDomainEntry(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CreateDomainEntryInput{
		// DomainEntry: *types.DomainEntry, // Required
		// DomainName: *string, // Required
	}

	if len(_lightsailDomainEntry) > 0 {
		if err := assignInputField(input, "DomainEntry", _lightsailDomainEntry); err != nil {
			log.Errorf("invalid --domain-entry: %s", err.Error())
			return
		}
	}
	if len(_lightsailDomainName) > 0 {
		input.DomainName = aws.String(_lightsailDomainName)
	}

	if resp, err := client.CreateDomainEntry(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates two URLs that are used to access a virtual computer’s graphical user
// interface (GUI) session. The primary URL initiates a web-based Amazon DCV
// session to the virtual computer's application. The secondary URL initiates a
// web-based Amazon DCV session to the virtual computer's operating session.
//
// Use StartGUISession to open the session.
func lightsail_CreateGUISessionAccessDetails(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CreateGUISessionAccessDetailsInput{
		// ResourceName: *string, // Required
	}

	if len(_lightsailResourceName) > 0 {
		input.ResourceName = aws.String(_lightsailResourceName)
	}

	if resp, err := client.CreateGUISessionAccessDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a snapshot of a specific virtual private server, or instance. You can
// use a snapshot to create a new instance that is based on that snapshot.
//
// The create instance snapshot operation supports tag-based access control via
// request tags. For more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_CreateInstanceSnapshot(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CreateInstanceSnapshotInput{
		// InstanceName: *string, // Required
		// InstanceSnapshotName: *string, // Required
	}

	if len(_lightsailInstanceName) > 0 {
		input.InstanceName = aws.String(_lightsailInstanceName)
	}
	if len(_lightsailInstanceSnapshotName) > 0 {
		input.InstanceSnapshotName = aws.String(_lightsailInstanceSnapshotName)
	}
	if len(_lightsailTags) > 0 {
		if err := assignInputField(input, "Tags", _lightsailTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateInstanceSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates one or more Amazon Lightsail instances.
// The create instances operation supports tag-based access control via request
// tags. For more information, see the [Lightsail Developer Guide].
//
// [Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_CreateInstances(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CreateInstancesInput{
		// AvailabilityZone: *string, // Required
		// BlueprintId: *string, // Required
		// BundleId: *string, // Required
		// InstanceNames: []string, // Required
	}

	if len(_lightsailAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_lightsailAvailabilityZone)
	}
	if len(_lightsailBlueprintId) > 0 {
		input.BlueprintId = aws.String(_lightsailBlueprintId)
	}
	if len(_lightsailBundleId) > 0 {
		input.BundleId = aws.String(_lightsailBundleId)
	}
	if len(_lightsailInstanceNames) > 0 {
		input.InstanceNames = append([]string(nil), _lightsailInstanceNames...)
	}
	if len(_lightsailAddOns) > 0 {
		if err := assignInputField(input, "AddOns", _lightsailAddOns); err != nil {
			log.Errorf("invalid --add-ons: %s", err.Error())
			return
		}
	}
	if len(_lightsailCustomImageName) > 0 {
		input.CustomImageName = aws.String(_lightsailCustomImageName)
	}
	if len(_lightsailIpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _lightsailIpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_lightsailKeyPairName) > 0 {
		input.KeyPairName = aws.String(_lightsailKeyPairName)
	}
	if len(_lightsailTags) > 0 {
		if err := assignInputField(input, "Tags", _lightsailTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_lightsailUserData) > 0 {
		input.UserData = aws.String(_lightsailUserData)
	}

	if resp, err := client.CreateInstances(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates one or more new instances from a manual or automatic snapshot of an
// instance.
//
// The create instances from snapshot operation supports tag-based access control
// via request tags and resource tags applied to the resource identified by
// instance snapshot name . For more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_CreateInstancesFromSnapshot(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CreateInstancesFromSnapshotInput{
		// AvailabilityZone: *string, // Required
		// BundleId: *string, // Required
		// InstanceNames: []string, // Required
	}

	if len(_lightsailAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_lightsailAvailabilityZone)
	}
	if len(_lightsailBundleId) > 0 {
		input.BundleId = aws.String(_lightsailBundleId)
	}
	if len(_lightsailInstanceNames) > 0 {
		input.InstanceNames = append([]string(nil), _lightsailInstanceNames...)
	}
	if len(_lightsailAddOns) > 0 {
		if err := assignInputField(input, "AddOns", _lightsailAddOns); err != nil {
			log.Errorf("invalid --add-ons: %s", err.Error())
			return
		}
	}
	if len(_lightsailAttachedDiskMapping) > 0 {
		if err := assignInputField(input, "AttachedDiskMapping", _lightsailAttachedDiskMapping); err != nil {
			log.Errorf("invalid --attached-disk-mapping: %s", err.Error())
			return
		}
	}
	if len(_lightsailInstanceSnapshotName) > 0 {
		input.InstanceSnapshotName = aws.String(_lightsailInstanceSnapshotName)
	}
	if len(_lightsailIpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _lightsailIpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_lightsailKeyPairName) > 0 {
		input.KeyPairName = aws.String(_lightsailKeyPairName)
	}
	if len(_lightsailRestoreDate) > 0 {
		input.RestoreDate = aws.String(_lightsailRestoreDate)
	}
	if len(_lightsailSourceInstanceName) > 0 {
		input.SourceInstanceName = aws.String(_lightsailSourceInstanceName)
	}
	if len(_lightsailTags) > 0 {
		if err := assignInputField(input, "Tags", _lightsailTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_lightsailUseLatestRestorableAutoSnapshot) > 0 {
		if err := assignInputField(input, "UseLatestRestorableAutoSnapshot", _lightsailUseLatestRestorableAutoSnapshot); err != nil {
			log.Errorf("invalid --use-latest-restorable-auto-snapshot: %s", err.Error())
			return
		}
	}
	if len(_lightsailUserData) > 0 {
		input.UserData = aws.String(_lightsailUserData)
	}

	if resp, err := client.CreateInstancesFromSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a custom SSH key pair that you can use with an Amazon Lightsail
// instance.
//
// Use the [DownloadDefaultKeyPair] action to create a Lightsail default key pair in an Amazon Web
// Services Region where a default key pair does not currently exist.
//
// The create key pair operation supports tag-based access control via request
// tags. For more information, see the [Amazon Lightsail Developer Guide].
//
// [DownloadDefaultKeyPair]: https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_DownloadDefaultKeyPair.html
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_CreateKeyPair(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CreateKeyPairInput{
		// KeyPairName: *string, // Required
	}

	if len(_lightsailKeyPairName) > 0 {
		input.KeyPairName = aws.String(_lightsailKeyPairName)
	}
	if len(_lightsailTags) > 0 {
		if err := assignInputField(input, "Tags", _lightsailTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateKeyPair(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Lightsail load balancer. To learn more about deciding whether to load
// balance your application, see [Configure your Lightsail instances for load balancing]. You can create up to 10 load balancers per AWS
// Region in your account.
//
// When you create a load balancer, you can specify a unique name and port
// settings. To change additional load balancer settings, use the
// UpdateLoadBalancerAttribute operation.
//
// The create load balancer operation supports tag-based access control via
// request tags. For more information, see the [Amazon Lightsail Developer Guide].
//
// [Configure your Lightsail instances for load balancing]: https://docs.aws.amazon.com/lightsail/latest/userguide/configure-lightsail-instances-for-load-balancing
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_CreateLoadBalancer(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CreateLoadBalancerInput{
		// InstancePort: int32, // Required
		// LoadBalancerName: *string, // Required
	}

	if len(_lightsailInstancePort) > 0 {
		if err := assignInputField(input, "InstancePort", _lightsailInstancePort); err != nil {
			log.Errorf("invalid --instance-port: %s", err.Error())
			return
		}
	}
	if len(_lightsailLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_lightsailLoadBalancerName)
	}
	if len(_lightsailCertificateAlternativeNames) > 0 {
		input.CertificateAlternativeNames = append([]string(nil), _lightsailCertificateAlternativeNames...)
	}
	if len(_lightsailCertificateDomainName) > 0 {
		input.CertificateDomainName = aws.String(_lightsailCertificateDomainName)
	}
	if len(_lightsailCertificateName) > 0 {
		input.CertificateName = aws.String(_lightsailCertificateName)
	}
	if len(_lightsailHealthCheckPath) > 0 {
		input.HealthCheckPath = aws.String(_lightsailHealthCheckPath)
	}
	if len(_lightsailIpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _lightsailIpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_lightsailTags) > 0 {
		if err := assignInputField(input, "Tags", _lightsailTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_lightsailTlsPolicyName) > 0 {
		input.TlsPolicyName = aws.String(_lightsailTlsPolicyName)
	}

	if resp, err := client.CreateLoadBalancer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an SSL/TLS certificate for an Amazon Lightsail load balancer.
// TLS is just an updated, more secure version of Secure Socket Layer (SSL).
//
// The CreateLoadBalancerTlsCertificate operation supports tag-based access
// control via resource tags applied to the resource identified by load balancer
// name . For more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_CreateLoadBalancerTlsCertificate(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CreateLoadBalancerTlsCertificateInput{
		// CertificateDomainName: *string, // Required
		// CertificateName: *string, // Required
		// LoadBalancerName: *string, // Required
	}

	if len(_lightsailCertificateDomainName) > 0 {
		input.CertificateDomainName = aws.String(_lightsailCertificateDomainName)
	}
	if len(_lightsailCertificateName) > 0 {
		input.CertificateName = aws.String(_lightsailCertificateName)
	}
	if len(_lightsailLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_lightsailLoadBalancerName)
	}
	if len(_lightsailCertificateAlternativeNames) > 0 {
		input.CertificateAlternativeNames = append([]string(nil), _lightsailCertificateAlternativeNames...)
	}
	if len(_lightsailTags) > 0 {
		if err := assignInputField(input, "Tags", _lightsailTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLoadBalancerTlsCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new database in Amazon Lightsail.
// The create relational database operation supports tag-based access control via
// request tags. For more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_CreateRelationalDatabase(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CreateRelationalDatabaseInput{
		// MasterDatabaseName: *string, // Required
		// MasterUsername: *string, // Required
		// RelationalDatabaseBlueprintId: *string, // Required
		// RelationalDatabaseBundleId: *string, // Required
		// RelationalDatabaseName: *string, // Required
	}

	if len(_lightsailMasterDatabaseName) > 0 {
		input.MasterDatabaseName = aws.String(_lightsailMasterDatabaseName)
	}
	if len(_lightsailMasterUsername) > 0 {
		input.MasterUsername = aws.String(_lightsailMasterUsername)
	}
	if len(_lightsailRelationalDatabaseBlueprintId) > 0 {
		input.RelationalDatabaseBlueprintId = aws.String(_lightsailRelationalDatabaseBlueprintId)
	}
	if len(_lightsailRelationalDatabaseBundleId) > 0 {
		input.RelationalDatabaseBundleId = aws.String(_lightsailRelationalDatabaseBundleId)
	}
	if len(_lightsailRelationalDatabaseName) > 0 {
		input.RelationalDatabaseName = aws.String(_lightsailRelationalDatabaseName)
	}
	if len(_lightsailAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_lightsailAvailabilityZone)
	}
	if len(_lightsailMasterUserPassword) > 0 {
		input.MasterUserPassword = aws.String(_lightsailMasterUserPassword)
	}
	if len(_lightsailPreferredBackupWindow) > 0 {
		input.PreferredBackupWindow = aws.String(_lightsailPreferredBackupWindow)
	}
	if len(_lightsailPreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_lightsailPreferredMaintenanceWindow)
	}
	if len(_lightsailPubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _lightsailPubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_lightsailTags) > 0 {
		if err := assignInputField(input, "Tags", _lightsailTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRelationalDatabase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new database from an existing database snapshot in Amazon Lightsail.
// You can create a new database from a snapshot in if something goes wrong with
// your original database, or to change it to a different plan, such as a high
// availability or standard plan.
//
// The create relational database from snapshot operation supports tag-based
// access control via request tags and resource tags applied to the resource
// identified by relationalDatabaseSnapshotName. For more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_CreateRelationalDatabaseFromSnapshot(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CreateRelationalDatabaseFromSnapshotInput{
		// RelationalDatabaseName: *string, // Required
	}

	if len(_lightsailRelationalDatabaseName) > 0 {
		input.RelationalDatabaseName = aws.String(_lightsailRelationalDatabaseName)
	}
	if len(_lightsailAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_lightsailAvailabilityZone)
	}
	if len(_lightsailPubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _lightsailPubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_lightsailRelationalDatabaseBundleId) > 0 {
		input.RelationalDatabaseBundleId = aws.String(_lightsailRelationalDatabaseBundleId)
	}
	if len(_lightsailRelationalDatabaseSnapshotName) > 0 {
		input.RelationalDatabaseSnapshotName = aws.String(_lightsailRelationalDatabaseSnapshotName)
	}
	if len(_lightsailRestoreTime) > 0 {
		if err := assignInputField(input, "RestoreTime", _lightsailRestoreTime); err != nil {
			log.Errorf("invalid --restore-time: %s", err.Error())
			return
		}
	}
	if len(_lightsailSourceRelationalDatabaseName) > 0 {
		input.SourceRelationalDatabaseName = aws.String(_lightsailSourceRelationalDatabaseName)
	}
	if len(_lightsailTags) > 0 {
		if err := assignInputField(input, "Tags", _lightsailTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_lightsailUseLatestRestorableTime) > 0 {
		if err := assignInputField(input, "UseLatestRestorableTime", _lightsailUseLatestRestorableTime); err != nil {
			log.Errorf("invalid --use-latest-restorable-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRelationalDatabaseFromSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a snapshot of your database in Amazon Lightsail. You can use snapshots
// for backups, to make copies of a database, and to save data before deleting a
// database.
//
// The create relational database snapshot operation supports tag-based access
// control via request tags. For more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_CreateRelationalDatabaseSnapshot(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.CreateRelationalDatabaseSnapshotInput{
		// RelationalDatabaseName: *string, // Required
		// RelationalDatabaseSnapshotName: *string, // Required
	}

	if len(_lightsailRelationalDatabaseName) > 0 {
		input.RelationalDatabaseName = aws.String(_lightsailRelationalDatabaseName)
	}
	if len(_lightsailRelationalDatabaseSnapshotName) > 0 {
		input.RelationalDatabaseSnapshotName = aws.String(_lightsailRelationalDatabaseSnapshotName)
	}
	if len(_lightsailTags) > 0 {
		if err := assignInputField(input, "Tags", _lightsailTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRelationalDatabaseSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an alarm.
// An alarm is used to monitor a single metric for one of your resources. When a
// metric condition is met, the alarm can notify you by email, SMS text message,
// and a banner displayed on the Amazon Lightsail console. For more information,
// see [Alarms in Amazon Lightsail].
//
// [Alarms in Amazon Lightsail]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-alarms
func lightsail_DeleteAlarm(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DeleteAlarmInput{
		// AlarmName: *string, // Required
	}

	if len(_lightsailAlarmName) > 0 {
		input.AlarmName = aws.String(_lightsailAlarmName)
	}

	if resp, err := client.DeleteAlarm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an automatic snapshot of an instance or disk. For more information, see
// the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-configuring-automatic-snapshots
func lightsail_DeleteAutoSnapshot(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DeleteAutoSnapshotInput{
		// Date: *string, // Required
		// ResourceName: *string, // Required
	}

	if len(_lightsailDate) > 0 {
		input.Date = aws.String(_lightsailDate)
	}
	if len(_lightsailResourceName) > 0 {
		input.ResourceName = aws.String(_lightsailResourceName)
	}

	if resp, err := client.DeleteAutoSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Amazon Lightsail bucket.
// When you delete your bucket, the bucket name is released and can be reused for
// a new bucket in your account or another Amazon Web Services account.
func lightsail_DeleteBucket(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DeleteBucketInput{
		// BucketName: *string, // Required
	}

	if len(_lightsailBucketName) > 0 {
		input.BucketName = aws.String(_lightsailBucketName)
	}
	if len(_lightsailForceDelete) > 0 {
		if err := assignInputField(input, "ForceDelete", _lightsailForceDelete); err != nil {
			log.Errorf("invalid --force-delete: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteBucket(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an access key for the specified Amazon Lightsail bucket.
// We recommend that you delete an access key if the secret access key is
// compromised.
//
// For more information about access keys, see [Creating access keys for a bucket in Amazon Lightsail] in the Amazon Lightsail Developer
// Guide.
//
// [Creating access keys for a bucket in Amazon Lightsail]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-creating-bucket-access-keys
func lightsail_DeleteBucketAccessKey(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DeleteBucketAccessKeyInput{
		// AccessKeyId: *string, // Required
		// BucketName: *string, // Required
	}

	if len(_lightsailAccessKeyId) > 0 {
		input.AccessKeyId = aws.String(_lightsailAccessKeyId)
	}
	if len(_lightsailBucketName) > 0 {
		input.BucketName = aws.String(_lightsailBucketName)
	}

	if resp, err := client.DeleteBucketAccessKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an SSL/TLS certificate for your Amazon Lightsail content delivery
// network (CDN) distribution.
//
// Certificates that are currently attached to a distribution cannot be deleted.
// Use the DetachCertificateFromDistribution action to detach a certificate from a
// distribution.
func lightsail_DeleteCertificate(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DeleteCertificateInput{
		// CertificateName: *string, // Required
	}

	if len(_lightsailCertificateName) > 0 {
		input.CertificateName = aws.String(_lightsailCertificateName)
	}

	if resp, err := client.DeleteCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a contact method.
// A contact method is used to send you notifications about your Amazon Lightsail
// resources. You can add one email address and one mobile phone number contact
// method in each Amazon Web Services Region. However, SMS text messaging is not
// supported in some Amazon Web Services Regions, and SMS text messages cannot be
// sent to some countries/regions. For more information, see [Notifications in Amazon Lightsail].
//
// [Notifications in Amazon Lightsail]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-notifications
func lightsail_DeleteContactMethod(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DeleteContactMethodInput{
		// Protocol: types.ContactProtocol, // Required
	}

	if len(_lightsailProtocol) > 0 {
		if err := assignInputField(input, "Protocol", _lightsailProtocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteContactMethod(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a container image that is registered to your Amazon Lightsail container
// service.
func lightsail_DeleteContainerImage(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DeleteContainerImageInput{
		// Image: *string, // Required
		// ServiceName: *string, // Required
	}

	if len(_lightsailImage) > 0 {
		input.Image = aws.String(_lightsailImage)
	}
	if len(_lightsailServiceName) > 0 {
		input.ServiceName = aws.String(_lightsailServiceName)
	}

	if resp, err := client.DeleteContainerImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes your Amazon Lightsail container service.
func lightsail_DeleteContainerService(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DeleteContainerServiceInput{
		// ServiceName: *string, // Required
	}

	if len(_lightsailServiceName) > 0 {
		input.ServiceName = aws.String(_lightsailServiceName)
	}

	if resp, err := client.DeleteContainerService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified block storage disk. The disk must be in the available
// state (not attached to a Lightsail instance).
//
// The disk may remain in the deleting state for several minutes.
//
// The delete disk operation supports tag-based access control via resource tags
// applied to the resource identified by disk name . For more information, see the [Amazon Lightsail Developer Guide]
// .
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_DeleteDisk(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DeleteDiskInput{
		// DiskName: *string, // Required
	}

	if len(_lightsailDiskName) > 0 {
		input.DiskName = aws.String(_lightsailDiskName)
	}
	if len(_lightsailForceDeleteAddOns) > 0 {
		if err := assignInputField(input, "ForceDeleteAddOns", _lightsailForceDeleteAddOns); err != nil {
			log.Errorf("invalid --force-delete-add-ons: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteDisk(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified disk snapshot.
// When you make periodic snapshots of a disk, the snapshots are incremental, and
// only the blocks on the device that have changed since your last snapshot are
// saved in the new snapshot. When you delete a snapshot, only the data not needed
// for any other snapshot is removed. So regardless of which prior snapshots have
// been deleted, all active snapshots will have access to all the information
// needed to restore the disk.
//
// The delete disk snapshot operation supports tag-based access control via
// resource tags applied to the resource identified by disk snapshot name . For
// more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_DeleteDiskSnapshot(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DeleteDiskSnapshotInput{
		// DiskSnapshotName: *string, // Required
	}

	if len(_lightsailDiskSnapshotName) > 0 {
		input.DiskSnapshotName = aws.String(_lightsailDiskSnapshotName)
	}

	if resp, err := client.DeleteDiskSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes your Amazon Lightsail content delivery network (CDN) distribution.
func lightsail_DeleteDistribution(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DeleteDistributionInput{}

	if len(_lightsailDistributionName) > 0 {
		input.DistributionName = aws.String(_lightsailDistributionName)
	}

	if resp, err := client.DeleteDistribution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified domain recordset and all of its domain records.
// The delete domain operation supports tag-based access control via resource tags
// applied to the resource identified by domain name . For more information, see
// the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_DeleteDomain(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DeleteDomainInput{
		// DomainName: *string, // Required
	}

	if len(_lightsailDomainName) > 0 {
		input.DomainName = aws.String(_lightsailDomainName)
	}

	if resp, err := client.DeleteDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specific domain entry.
// The delete domain entry operation supports tag-based access control via
// resource tags applied to the resource identified by domain name . For more
// information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_DeleteDomainEntry(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DeleteDomainEntryInput{
		// DomainEntry: *types.DomainEntry, // Required
		// DomainName: *string, // Required
	}

	if len(_lightsailDomainEntry) > 0 {
		if err := assignInputField(input, "DomainEntry", _lightsailDomainEntry); err != nil {
			log.Errorf("invalid --domain-entry: %s", err.Error())
			return
		}
	}
	if len(_lightsailDomainName) > 0 {
		input.DomainName = aws.String(_lightsailDomainName)
	}

	if resp, err := client.DeleteDomainEntry(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Lightsail instance.
// The delete instance operation supports tag-based access control via resource
// tags applied to the resource identified by instance name . For more information,
// see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_DeleteInstance(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DeleteInstanceInput{
		// InstanceName: *string, // Required
	}

	if len(_lightsailInstanceName) > 0 {
		input.InstanceName = aws.String(_lightsailInstanceName)
	}
	if len(_lightsailForceDeleteAddOns) > 0 {
		if err := assignInputField(input, "ForceDeleteAddOns", _lightsailForceDeleteAddOns); err != nil {
			log.Errorf("invalid --force-delete-add-ons: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specific snapshot of a virtual private server (or instance).
// The delete instance snapshot operation supports tag-based access control via
// resource tags applied to the resource identified by instance snapshot name . For
// more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_DeleteInstanceSnapshot(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DeleteInstanceSnapshotInput{
		// InstanceSnapshotName: *string, // Required
	}

	if len(_lightsailInstanceSnapshotName) > 0 {
		input.InstanceSnapshotName = aws.String(_lightsailInstanceSnapshotName)
	}

	if resp, err := client.DeleteInstanceSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified key pair by removing the public key from Amazon Lightsail.
// You can delete key pairs that were created using the [ImportKeyPair] and [CreateKeyPair] actions, as well as
// the Lightsail default key pair. A new default key pair will not be created
// unless you launch an instance without specifying a custom key pair, or you call
// the [DownloadDefaultKeyPair]API.
//
// The delete key pair operation supports tag-based access control via resource
// tags applied to the resource identified by key pair name . For more information,
// see the [Amazon Lightsail Developer Guide].
//
// [ImportKeyPair]: https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_ImportKeyPair.html
// [CreateKeyPair]: https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_CreateKeyPair.html
// [DownloadDefaultKeyPair]: https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_DownloadDefaultKeyPair.html
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_DeleteKeyPair(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DeleteKeyPairInput{
		// KeyPairName: *string, // Required
	}

	if len(_lightsailKeyPairName) > 0 {
		input.KeyPairName = aws.String(_lightsailKeyPairName)
	}
	if len(_lightsailExpectedFingerprint) > 0 {
		input.ExpectedFingerprint = aws.String(_lightsailExpectedFingerprint)
	}

	if resp, err := client.DeleteKeyPair(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the known host key or certificate used by the Amazon Lightsail
// browser-based SSH or RDP clients to authenticate an instance. This operation
// enables the Lightsail browser-based SSH or RDP clients to connect to the
// instance after a host key mismatch.
//
// Perform this operation only if you were expecting the host key or certificate
// mismatch or if you are familiar with the new host key or certificate on the
// instance. For more information, see [Troubleshooting connection issues when using the Amazon Lightsail browser-based SSH or RDP client].
//
// [Troubleshooting connection issues when using the Amazon Lightsail browser-based SSH or RDP client]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-troubleshooting-browser-based-ssh-rdp-client-connection
func lightsail_DeleteKnownHostKeys(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DeleteKnownHostKeysInput{
		// InstanceName: *string, // Required
	}

	if len(_lightsailInstanceName) > 0 {
		input.InstanceName = aws.String(_lightsailInstanceName)
	}

	if resp, err := client.DeleteKnownHostKeys(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Lightsail load balancer and all its associated SSL/TLS certificates.
// Once the load balancer is deleted, you will need to create a new load balancer,
// create a new certificate, and verify domain ownership again.
//
// The delete load balancer operation supports tag-based access control via
// resource tags applied to the resource identified by load balancer name . For
// more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_DeleteLoadBalancer(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DeleteLoadBalancerInput{
		// LoadBalancerName: *string, // Required
	}

	if len(_lightsailLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_lightsailLoadBalancerName)
	}

	if resp, err := client.DeleteLoadBalancer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an SSL/TLS certificate associated with a Lightsail load balancer.
// The DeleteLoadBalancerTlsCertificate operation supports tag-based access
// control via resource tags applied to the resource identified by load balancer
// name . For more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_DeleteLoadBalancerTlsCertificate(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DeleteLoadBalancerTlsCertificateInput{
		// CertificateName: *string, // Required
		// LoadBalancerName: *string, // Required
	}

	if len(_lightsailCertificateName) > 0 {
		input.CertificateName = aws.String(_lightsailCertificateName)
	}
	if len(_lightsailLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_lightsailLoadBalancerName)
	}
	if len(_lightsailForce) > 0 {
		if err := assignInputField(input, "Force", _lightsailForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteLoadBalancerTlsCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a database in Amazon Lightsail.
// The delete relational database operation supports tag-based access control via
// resource tags applied to the resource identified by relationalDatabaseName. For
// more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_DeleteRelationalDatabase(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DeleteRelationalDatabaseInput{
		// RelationalDatabaseName: *string, // Required
	}

	if len(_lightsailRelationalDatabaseName) > 0 {
		input.RelationalDatabaseName = aws.String(_lightsailRelationalDatabaseName)
	}
	if len(_lightsailFinalRelationalDatabaseSnapshotName) > 0 {
		input.FinalRelationalDatabaseSnapshotName = aws.String(_lightsailFinalRelationalDatabaseSnapshotName)
	}
	if len(_lightsailSkipFinalSnapshot) > 0 {
		if err := assignInputField(input, "SkipFinalSnapshot", _lightsailSkipFinalSnapshot); err != nil {
			log.Errorf("invalid --skip-final-snapshot: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteRelationalDatabase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a database snapshot in Amazon Lightsail.
// The delete relational database snapshot operation supports tag-based access
// control via resource tags applied to the resource identified by
// relationalDatabaseName. For more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_DeleteRelationalDatabaseSnapshot(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DeleteRelationalDatabaseSnapshotInput{
		// RelationalDatabaseSnapshotName: *string, // Required
	}

	if len(_lightsailRelationalDatabaseSnapshotName) > 0 {
		input.RelationalDatabaseSnapshotName = aws.String(_lightsailRelationalDatabaseSnapshotName)
	}

	if resp, err := client.DeleteRelationalDatabaseSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detaches an SSL/TLS certificate from your Amazon Lightsail content delivery
// network (CDN) distribution.
//
// After the certificate is detached, your distribution stops accepting traffic
// for all of the domains that are associated with the certificate.
func lightsail_DetachCertificateFromDistribution(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DetachCertificateFromDistributionInput{
		// DistributionName: *string, // Required
	}

	if len(_lightsailDistributionName) > 0 {
		input.DistributionName = aws.String(_lightsailDistributionName)
	}

	if resp, err := client.DetachCertificateFromDistribution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detaches a stopped block storage disk from a Lightsail instance. Make sure to
// unmount any file systems on the device within your operating system before
// stopping the instance and detaching the disk.
//
// The detach disk operation supports tag-based access control via resource tags
// applied to the resource identified by disk name . For more information, see the [Amazon Lightsail Developer Guide]
// .
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_DetachDisk(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DetachDiskInput{
		// DiskName: *string, // Required
	}

	if len(_lightsailDiskName) > 0 {
		input.DiskName = aws.String(_lightsailDiskName)
	}

	if resp, err := client.DetachDisk(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detaches the specified instances from a Lightsail load balancer.
// This operation waits until the instances are no longer needed before they are
// detached from the load balancer.
//
// The detach instances from load balancer operation supports tag-based access
// control via resource tags applied to the resource identified by load balancer
// name . For more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_DetachInstancesFromLoadBalancer(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DetachInstancesFromLoadBalancerInput{
		// InstanceNames: []string, // Required
		// LoadBalancerName: *string, // Required
	}

	if len(_lightsailInstanceNames) > 0 {
		input.InstanceNames = append([]string(nil), _lightsailInstanceNames...)
	}
	if len(_lightsailLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_lightsailLoadBalancerName)
	}

	if resp, err := client.DetachInstancesFromLoadBalancer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detaches a static IP from the Amazon Lightsail instance to which it is attached.
func lightsail_DetachStaticIp(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DetachStaticIpInput{
		// StaticIpName: *string, // Required
	}

	if len(_lightsailStaticIpName) > 0 {
		input.StaticIpName = aws.String(_lightsailStaticIpName)
	}

	if resp, err := client.DetachStaticIp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables an add-on for an Amazon Lightsail resource. For more information, see
// the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-configuring-automatic-snapshots
func lightsail_DisableAddOn(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DisableAddOnInput{
		// AddOnType: types.AddOnType, // Required
		// ResourceName: *string, // Required
	}

	if len(_lightsailAddOnType) > 0 {
		if err := assignInputField(input, "AddOnType", _lightsailAddOnType); err != nil {
			log.Errorf("invalid --add-on-type: %s", err.Error())
			return
		}
	}
	if len(_lightsailResourceName) > 0 {
		input.ResourceName = aws.String(_lightsailResourceName)
	}

	if resp, err := client.DisableAddOn(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Downloads the regional Amazon Lightsail default key pair.
// This action also creates a Lightsail default key pair if a default key pair
// does not currently exist in the Amazon Web Services Region.
func lightsail_DownloadDefaultKeyPair(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.DownloadDefaultKeyPairInput{}

	if resp, err := client.DownloadDefaultKeyPair(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables or modifies an add-on for an Amazon Lightsail resource. For more
// information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-configuring-automatic-snapshots
func lightsail_EnableAddOn(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.EnableAddOnInput{
		// AddOnRequest: *types.AddOnRequest, // Required
		// ResourceName: *string, // Required
	}

	if len(_lightsailAddOnRequest) > 0 {
		if err := assignInputField(input, "AddOnRequest", _lightsailAddOnRequest); err != nil {
			log.Errorf("invalid --add-on-request: %s", err.Error())
			return
		}
	}
	if len(_lightsailResourceName) > 0 {
		input.ResourceName = aws.String(_lightsailResourceName)
	}

	if resp, err := client.EnableAddOn(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Exports an Amazon Lightsail instance or block storage disk snapshot to Amazon
// Elastic Compute Cloud (Amazon EC2). This operation results in an export snapshot
// record that can be used with the create cloud formation stack operation to
// create new Amazon EC2 instances.
//
// Exported instance snapshots appear in Amazon EC2 as Amazon Machine Images
// (AMIs), and the instance system disk appears as an Amazon Elastic Block Store
// (Amazon EBS) volume. Exported disk snapshots appear in Amazon EC2 as Amazon EBS
// volumes. Snapshots are exported to the same Amazon Web Services Region in Amazon
// EC2 as the source Lightsail snapshot.
//
// The export snapshot operation supports tag-based access control via resource
// tags applied to the resource identified by source snapshot name . For more
// information, see the [Amazon Lightsail Developer Guide].
//
// Use the get instance snapshots or get disk snapshots operations to get a list
// of snapshots that you can export to Amazon EC2.
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_ExportSnapshot(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.ExportSnapshotInput{
		// SourceSnapshotName: *string, // Required
	}

	if len(_lightsailSourceSnapshotName) > 0 {
		input.SourceSnapshotName = aws.String(_lightsailSourceSnapshotName)
	}

	if resp, err := client.ExportSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the names of all active (not deleted) resources.
func lightsail_GetActiveNames(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetActiveNamesInput{}

	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetActiveNames(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the configured alarms. Specify an alarm name in your
// request to return information about a specific alarm, or specify a monitored
// resource name to return information about all alarms for a specific resource.
//
// An alarm is used to monitor a single metric for one of your resources. When a
// metric condition is met, the alarm can notify you by email, SMS text message,
// and a banner displayed on the Amazon Lightsail console. For more information,
// see [Alarms in Amazon Lightsail].
//
// [Alarms in Amazon Lightsail]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-alarms
func lightsail_GetAlarms(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetAlarmsInput{}

	if len(_lightsailAlarmName) > 0 {
		input.AlarmName = aws.String(_lightsailAlarmName)
	}
	if len(_lightsailMonitoredResourceName) > 0 {
		input.MonitoredResourceName = aws.String(_lightsailMonitoredResourceName)
	}
	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetAlarms(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the available automatic snapshots for an instance or disk. For more
// information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-configuring-automatic-snapshots
func lightsail_GetAutoSnapshots(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetAutoSnapshotsInput{
		// ResourceName: *string, // Required
	}

	if len(_lightsailResourceName) > 0 {
		input.ResourceName = aws.String(_lightsailResourceName)
	}

	if resp, err := client.GetAutoSnapshots(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the list of available instance images, or blueprints. You can use a
// blueprint to create a new instance already running a specific operating system,
// as well as a preinstalled app or development stack. The software each instance
// is running depends on the blueprint image you choose.
//
// Use active blueprints when creating new instances. Inactive blueprints are
// listed to support customers with existing instances and are not necessarily
// available to create new instances. Blueprints are marked inactive when they
// become outdated due to operating system updates or new application releases.
func lightsail_GetBlueprints(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetBlueprintsInput{}

	if len(_lightsailAppCategory) > 0 {
		if err := assignInputField(input, "AppCategory", _lightsailAppCategory); err != nil {
			log.Errorf("invalid --app-category: %s", err.Error())
			return
		}
	}
	if len(_lightsailIncludeInactive) > 0 {
		if err := assignInputField(input, "IncludeInactive", _lightsailIncludeInactive); err != nil {
			log.Errorf("invalid --include-inactive: %s", err.Error())
			return
		}
	}
	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetBlueprints(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the existing access key IDs for the specified Amazon Lightsail bucket.
// This action does not return the secret access key value of an access key. You
// can get a secret access key only when you create it from the response of the [CreateBucketAccessKey]
// action. If you lose the secret access key, you must create a new access key.
//
// [CreateBucketAccessKey]: https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_CreateBucketAccessKey.html
func lightsail_GetBucketAccessKeys(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetBucketAccessKeysInput{
		// BucketName: *string, // Required
	}

	if len(_lightsailBucketName) > 0 {
		input.BucketName = aws.String(_lightsailBucketName)
	}

	if resp, err := client.GetBucketAccessKeys(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the bundles that you can apply to a Amazon Lightsail bucket.
// The bucket bundle specifies the monthly cost, storage quota, and data transfer
// quota for a bucket.
//
// Use the [UpdateBucketBundle] action to update the bundle for a bucket.
//
// [UpdateBucketBundle]: https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_UpdateBucketBundle.html
func lightsail_GetBucketBundles(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetBucketBundlesInput{}

	if len(_lightsailIncludeInactive) > 0 {
		if err := assignInputField(input, "IncludeInactive", _lightsailIncludeInactive); err != nil {
			log.Errorf("invalid --include-inactive: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetBucketBundles(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the data points of a specific metric for an Amazon Lightsail bucket.
// Metrics report the utilization of a bucket. View and collect metric data
// regularly to monitor the number of objects stored in a bucket (including object
// versions) and the storage space used by those objects.
func lightsail_GetBucketMetricData(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetBucketMetricDataInput{
		// BucketName: *string, // Required
		// EndTime: *time.Time, // Required
		// MetricName: types.BucketMetricName, // Required
		// Period: *int32, // Required
		// StartTime: *time.Time, // Required
		// Statistics: []types.MetricStatistic, // Required
		// Unit: types.MetricUnit, // Required
	}

	if len(_lightsailBucketName) > 0 {
		input.BucketName = aws.String(_lightsailBucketName)
	}
	if len(_lightsailEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _lightsailEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_lightsailMetricName) > 0 {
		if err := assignInputField(input, "MetricName", _lightsailMetricName); err != nil {
			log.Errorf("invalid --metric-name: %s", err.Error())
			return
		}
	}
	if len(_lightsailPeriod) > 0 {
		if err := assignInputField(input, "Period", _lightsailPeriod); err != nil {
			log.Errorf("invalid --period: %s", err.Error())
			return
		}
	}
	if len(_lightsailStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _lightsailStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_lightsailStatistics) > 0 {
		if err := assignInputField(input, "Statistics", _lightsailStatistics); err != nil {
			log.Errorf("invalid --statistics: %s", err.Error())
			return
		}
	}
	if len(_lightsailUnit) > 0 {
		if err := assignInputField(input, "Unit", _lightsailUnit); err != nil {
			log.Errorf("invalid --unit: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetBucketMetricData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about one or more Amazon Lightsail buckets. The information
// returned includes the synchronization status of the Amazon Simple Storage
// Service (Amazon S3) account-level block public access feature for your Lightsail
// buckets.
//
// For more information about buckets, see [Buckets in Amazon Lightsail] in the Amazon Lightsail Developer
// Guide.
//
// [Buckets in Amazon Lightsail]: https://docs.aws.amazon.com/lightsail/latest/userguide/buckets-in-amazon-lightsail
func lightsail_GetBuckets(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetBucketsInput{}

	if len(_lightsailBucketName) > 0 {
		input.BucketName = aws.String(_lightsailBucketName)
	}
	if len(_lightsailIncludeConnectedResources) > 0 {
		if err := assignInputField(input, "IncludeConnectedResources", _lightsailIncludeConnectedResources); err != nil {
			log.Errorf("invalid --include-connected-resources: %s", err.Error())
			return
		}
	}
	if len(_lightsailIncludeCors) > 0 {
		if err := assignInputField(input, "IncludeCors", _lightsailIncludeCors); err != nil {
			log.Errorf("invalid --include-cors: %s", err.Error())
			return
		}
	}
	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetBuckets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the bundles that you can apply to an Amazon Lightsail instance when you
// create it.
//
// A bundle describes the specifications of an instance, such as the monthly cost,
// amount of memory, the number of vCPUs, amount of storage space, and monthly
// network data transfer quota.
//
// Bundles are referred to as instance plans in the Lightsail console.
func lightsail_GetBundles(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetBundlesInput{}

	if len(_lightsailAppCategory) > 0 {
		if err := assignInputField(input, "AppCategory", _lightsailAppCategory); err != nil {
			log.Errorf("invalid --app-category: %s", err.Error())
			return
		}
	}
	if len(_lightsailIncludeInactive) > 0 {
		if err := assignInputField(input, "IncludeInactive", _lightsailIncludeInactive); err != nil {
			log.Errorf("invalid --include-inactive: %s", err.Error())
			return
		}
	}
	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetBundles(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about one or more Amazon Lightsail SSL/TLS certificates.
// To get a summary of a certificate, omit includeCertificateDetails from your
// request. The response will include only the certificate Amazon Resource Name
// (ARN), certificate name, domain name, and tags.
func lightsail_GetCertificates(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetCertificatesInput{}

	if len(_lightsailCertificateName) > 0 {
		input.CertificateName = aws.String(_lightsailCertificateName)
	}
	if len(_lightsailCertificateStatuses) > 0 {
		if err := assignInputField(input, "CertificateStatuses", _lightsailCertificateStatuses); err != nil {
			log.Errorf("invalid --certificate-statuses: %s", err.Error())
			return
		}
	}
	if len(_lightsailIncludeCertificateDetails) > 0 {
		if err := assignInputField(input, "IncludeCertificateDetails", _lightsailIncludeCertificateDetails); err != nil {
			log.Errorf("invalid --include-certificate-details: %s", err.Error())
			return
		}
	}
	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetCertificates(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the CloudFormation stack record created as a result of the create cloud
// formation stack operation.
//
// An AWS CloudFormation stack is used to create a new Amazon EC2 instance from an
// exported Lightsail snapshot.
func lightsail_GetCloudFormationStackRecords(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetCloudFormationStackRecordsInput{}

	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetCloudFormationStackRecords(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the configured contact methods. Specify a protocol in
// your request to return information about a specific contact method.
//
// A contact method is used to send you notifications about your Amazon Lightsail
// resources. You can add one email address and one mobile phone number contact
// method in each Amazon Web Services Region. However, SMS text messaging is not
// supported in some Amazon Web Services Regions, and SMS text messages cannot be
// sent to some countries/regions. For more information, see [Notifications in Amazon Lightsail].
//
// [Notifications in Amazon Lightsail]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-notifications
func lightsail_GetContactMethods(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetContactMethodsInput{}

	if len(_lightsailProtocols) > 0 {
		if err := assignInputField(input, "Protocols", _lightsailProtocols); err != nil {
			log.Errorf("invalid --protocols: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetContactMethods(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about Amazon Lightsail containers, such as the current
// version of the Lightsail Control (lightsailctl) plugin.
func lightsail_GetContainerAPIMetadata(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetContainerAPIMetadataInput{}

	if resp, err := client.GetContainerAPIMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the container images that are registered to your Amazon Lightsail
// container service.
//
// If you created a deployment on your Lightsail container service that uses
// container images from a public registry like Docker Hub, those images are not
// returned as part of this action. Those images are not registered to your
// Lightsail container service.
func lightsail_GetContainerImages(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetContainerImagesInput{
		// ServiceName: *string, // Required
	}

	if len(_lightsailServiceName) > 0 {
		input.ServiceName = aws.String(_lightsailServiceName)
	}

	if resp, err := client.GetContainerImages(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the log events of a container of your Amazon Lightsail container
// service.
//
// If your container service has more than one node (i.e., a scale greater than
// 1), then the log events that are returned for the specified container are merged
// from all nodes on your container service.
//
// Container logs are retained for a certain amount of time. For more information,
// see [Amazon Lightsail endpoints and quotas]in the Amazon Web Services General Reference.
//
// [Amazon Lightsail endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/lightsail.html
func lightsail_GetContainerLog(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetContainerLogInput{
		// ContainerName: *string, // Required
		// ServiceName: *string, // Required
	}

	if len(_lightsailContainerName) > 0 {
		input.ContainerName = aws.String(_lightsailContainerName)
	}
	if len(_lightsailServiceName) > 0 {
		input.ServiceName = aws.String(_lightsailServiceName)
	}
	if len(_lightsailEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _lightsailEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_lightsailFilterPattern) > 0 {
		input.FilterPattern = aws.String(_lightsailFilterPattern)
	}
	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}
	if len(_lightsailStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _lightsailStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetContainerLog(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the deployments for your Amazon Lightsail container service
// A deployment specifies the settings, such as the ports and launch command, of
// containers that are deployed to your container service.
//
// The deployments are ordered by version in ascending order. The newest version
// is listed at the top of the response.
//
// A set number of deployments are kept before the oldest one is replaced with the
// newest one. For more information, see [Amazon Lightsail endpoints and quotas]in the Amazon Web Services General
// Reference.
//
// [Amazon Lightsail endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/lightsail.html
func lightsail_GetContainerServiceDeployments(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetContainerServiceDeploymentsInput{
		// ServiceName: *string, // Required
	}

	if len(_lightsailServiceName) > 0 {
		input.ServiceName = aws.String(_lightsailServiceName)
	}

	if resp, err := client.GetContainerServiceDeployments(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the data points of a specific metric of your Amazon Lightsail container
// service.
//
// Metrics report the utilization of your resources. Monitor and collect metric
// data regularly to maintain the reliability, availability, and performance of
// your resources.
func lightsail_GetContainerServiceMetricData(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetContainerServiceMetricDataInput{
		// EndTime: *time.Time, // Required
		// MetricName: types.ContainerServiceMetricName, // Required
		// Period: *int32, // Required
		// ServiceName: *string, // Required
		// StartTime: *time.Time, // Required
		// Statistics: []types.MetricStatistic, // Required
	}

	if len(_lightsailEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _lightsailEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_lightsailMetricName) > 0 {
		if err := assignInputField(input, "MetricName", _lightsailMetricName); err != nil {
			log.Errorf("invalid --metric-name: %s", err.Error())
			return
		}
	}
	if len(_lightsailPeriod) > 0 {
		if err := assignInputField(input, "Period", _lightsailPeriod); err != nil {
			log.Errorf("invalid --period: %s", err.Error())
			return
		}
	}
	if len(_lightsailServiceName) > 0 {
		input.ServiceName = aws.String(_lightsailServiceName)
	}
	if len(_lightsailStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _lightsailStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_lightsailStatistics) > 0 {
		if err := assignInputField(input, "Statistics", _lightsailStatistics); err != nil {
			log.Errorf("invalid --statistics: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetContainerServiceMetricData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the list of powers that can be specified for your Amazon Lightsail
// container services.
//
// The power specifies the amount of memory, the number of vCPUs, and the base
// price of the container service.
func lightsail_GetContainerServicePowers(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetContainerServicePowersInput{}

	if resp, err := client.GetContainerServicePowers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about one or more of your Amazon Lightsail container
// services.
func lightsail_GetContainerServices(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetContainerServicesInput{}

	if len(_lightsailServiceName) > 0 {
		input.ServiceName = aws.String(_lightsailServiceName)
	}

	if resp, err := client.GetContainerServices(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the cost estimate for a specified resource. A cost
// estimate will not generate for a resource that has been deleted.
func lightsail_GetCostEstimate(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetCostEstimateInput{
		// EndTime: *time.Time, // Required
		// ResourceName: *string, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_lightsailEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _lightsailEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_lightsailResourceName) > 0 {
		input.ResourceName = aws.String(_lightsailResourceName)
	}
	if len(_lightsailStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _lightsailStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetCostEstimate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific block storage disk.
func lightsail_GetDisk(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetDiskInput{
		// DiskName: *string, // Required
	}

	if len(_lightsailDiskName) > 0 {
		input.DiskName = aws.String(_lightsailDiskName)
	}

	if resp, err := client.GetDisk(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific block storage disk snapshot.
func lightsail_GetDiskSnapshot(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetDiskSnapshotInput{
		// DiskSnapshotName: *string, // Required
	}

	if len(_lightsailDiskSnapshotName) > 0 {
		input.DiskSnapshotName = aws.String(_lightsailDiskSnapshotName)
	}

	if resp, err := client.GetDiskSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about all block storage disk snapshots in your AWS account
// and region.
func lightsail_GetDiskSnapshots(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetDiskSnapshotsInput{}

	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetDiskSnapshots(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about all block storage disks in your AWS account and
// region.
func lightsail_GetDisks(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetDisksInput{}

	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetDisks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the bundles that can be applied to your Amazon Lightsail content
// delivery network (CDN) distributions.
//
// A distribution bundle specifies the monthly network transfer quota and monthly
// cost of your distribution.
func lightsail_GetDistributionBundles(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetDistributionBundlesInput{}

	if resp, err := client.GetDistributionBundles(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the timestamp and status of the last cache reset of a specific Amazon
// Lightsail content delivery network (CDN) distribution.
func lightsail_GetDistributionLatestCacheReset(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetDistributionLatestCacheResetInput{}

	if len(_lightsailDistributionName) > 0 {
		input.DistributionName = aws.String(_lightsailDistributionName)
	}

	if resp, err := client.GetDistributionLatestCacheReset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the data points of a specific metric for an Amazon Lightsail content
// delivery network (CDN) distribution.
//
// Metrics report the utilization of your resources, and the error counts
// generated by them. Monitor and collect metric data regularly to maintain the
// reliability, availability, and performance of your resources.
func lightsail_GetDistributionMetricData(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetDistributionMetricDataInput{
		// DistributionName: *string, // Required
		// EndTime: *time.Time, // Required
		// MetricName: types.DistributionMetricName, // Required
		// Period: *int32, // Required
		// StartTime: *time.Time, // Required
		// Statistics: []types.MetricStatistic, // Required
		// Unit: types.MetricUnit, // Required
	}

	if len(_lightsailDistributionName) > 0 {
		input.DistributionName = aws.String(_lightsailDistributionName)
	}
	if len(_lightsailEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _lightsailEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_lightsailMetricName) > 0 {
		if err := assignInputField(input, "MetricName", _lightsailMetricName); err != nil {
			log.Errorf("invalid --metric-name: %s", err.Error())
			return
		}
	}
	if len(_lightsailPeriod) > 0 {
		if err := assignInputField(input, "Period", _lightsailPeriod); err != nil {
			log.Errorf("invalid --period: %s", err.Error())
			return
		}
	}
	if len(_lightsailStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _lightsailStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_lightsailStatistics) > 0 {
		if err := assignInputField(input, "Statistics", _lightsailStatistics); err != nil {
			log.Errorf("invalid --statistics: %s", err.Error())
			return
		}
	}
	if len(_lightsailUnit) > 0 {
		if err := assignInputField(input, "Unit", _lightsailUnit); err != nil {
			log.Errorf("invalid --unit: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetDistributionMetricData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about one or more of your Amazon Lightsail content delivery
// network (CDN) distributions.
func lightsail_GetDistributions(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetDistributionsInput{}

	if len(_lightsailDistributionName) > 0 {
		input.DistributionName = aws.String(_lightsailDistributionName)
	}
	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetDistributions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific domain recordset.
func lightsail_GetDomain(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetDomainInput{
		// DomainName: *string, // Required
	}

	if len(_lightsailDomainName) > 0 {
		input.DomainName = aws.String(_lightsailDomainName)
	}

	if resp, err := client.GetDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of all domains in the user's account.
func lightsail_GetDomains(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetDomainsInput{}

	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetDomains(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns all export snapshot records created as a result of the export snapshot
// operation.
//
// An export snapshot record can be used to create a new Amazon EC2 instance and
// its related resources with the [CreateCloudFormationStack]action.
//
// [CreateCloudFormationStack]: https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_CreateCloudFormationStack.html
func lightsail_GetExportSnapshotRecords(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetExportSnapshotRecordsInput{}

	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetExportSnapshotRecords(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific Amazon Lightsail instance, which is a
// virtual private server.
func lightsail_GetInstance(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetInstanceInput{
		// InstanceName: *string, // Required
	}

	if len(_lightsailInstanceName) > 0 {
		input.InstanceName = aws.String(_lightsailInstanceName)
	}

	if resp, err := client.GetInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns temporary SSH keys you can use to connect to a specific virtual private
// server, or instance.
//
// The get instance access details operation supports tag-based access control via
// resource tags applied to the resource identified by instance name . For more
// information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_GetInstanceAccessDetails(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetInstanceAccessDetailsInput{
		// InstanceName: *string, // Required
	}

	if len(_lightsailInstanceName) > 0 {
		input.InstanceName = aws.String(_lightsailInstanceName)
	}
	if len(_lightsailProtocol) > 0 {
		if err := assignInputField(input, "Protocol", _lightsailProtocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetInstanceAccessDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the data points for the specified Amazon Lightsail instance metric,
// given an instance name.
//
// Metrics report the utilization of your resources, and the error counts
// generated by them. Monitor and collect metric data regularly to maintain the
// reliability, availability, and performance of your resources.
func lightsail_GetInstanceMetricData(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetInstanceMetricDataInput{
		// EndTime: *time.Time, // Required
		// InstanceName: *string, // Required
		// MetricName: types.InstanceMetricName, // Required
		// Period: *int32, // Required
		// StartTime: *time.Time, // Required
		// Statistics: []types.MetricStatistic, // Required
		// Unit: types.MetricUnit, // Required
	}

	if len(_lightsailEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _lightsailEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_lightsailInstanceName) > 0 {
		input.InstanceName = aws.String(_lightsailInstanceName)
	}
	if len(_lightsailMetricName) > 0 {
		if err := assignInputField(input, "MetricName", _lightsailMetricName); err != nil {
			log.Errorf("invalid --metric-name: %s", err.Error())
			return
		}
	}
	if len(_lightsailPeriod) > 0 {
		if err := assignInputField(input, "Period", _lightsailPeriod); err != nil {
			log.Errorf("invalid --period: %s", err.Error())
			return
		}
	}
	if len(_lightsailStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _lightsailStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_lightsailStatistics) > 0 {
		if err := assignInputField(input, "Statistics", _lightsailStatistics); err != nil {
			log.Errorf("invalid --statistics: %s", err.Error())
			return
		}
	}
	if len(_lightsailUnit) > 0 {
		if err := assignInputField(input, "Unit", _lightsailUnit); err != nil {
			log.Errorf("invalid --unit: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetInstanceMetricData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the firewall port states for a specific Amazon Lightsail instance, the
// IP addresses allowed to connect to the instance through the ports, and the
// protocol.
func lightsail_GetInstancePortStates(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetInstancePortStatesInput{
		// InstanceName: *string, // Required
	}

	if len(_lightsailInstanceName) > 0 {
		input.InstanceName = aws.String(_lightsailInstanceName)
	}

	if resp, err := client.GetInstancePortStates(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific instance snapshot.
func lightsail_GetInstanceSnapshot(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetInstanceSnapshotInput{
		// InstanceSnapshotName: *string, // Required
	}

	if len(_lightsailInstanceSnapshotName) > 0 {
		input.InstanceSnapshotName = aws.String(_lightsailInstanceSnapshotName)
	}

	if resp, err := client.GetInstanceSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns all instance snapshots for the user's account.
func lightsail_GetInstanceSnapshots(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetInstanceSnapshotsInput{}

	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetInstanceSnapshots(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the state of a specific instance. Works on one instance at a time.
func lightsail_GetInstanceState(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetInstanceStateInput{
		// InstanceName: *string, // Required
	}

	if len(_lightsailInstanceName) > 0 {
		input.InstanceName = aws.String(_lightsailInstanceName)
	}

	if resp, err := client.GetInstanceState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about all Amazon Lightsail virtual private servers, or
// instances.
func lightsail_GetInstances(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetInstancesInput{}

	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetInstances(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific key pair.
func lightsail_GetKeyPair(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetKeyPairInput{
		// KeyPairName: *string, // Required
	}

	if len(_lightsailKeyPairName) > 0 {
		input.KeyPairName = aws.String(_lightsailKeyPairName)
	}

	if resp, err := client.GetKeyPair(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about all key pairs in the user's account.
func lightsail_GetKeyPairs(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetKeyPairsInput{}

	if len(_lightsailIncludeDefaultKeyPair) > 0 {
		if err := assignInputField(input, "IncludeDefaultKeyPair", _lightsailIncludeDefaultKeyPair); err != nil {
			log.Errorf("invalid --include-default-key-pair: %s", err.Error())
			return
		}
	}
	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetKeyPairs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified Lightsail load balancer.
func lightsail_GetLoadBalancer(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetLoadBalancerInput{
		// LoadBalancerName: *string, // Required
	}

	if len(_lightsailLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_lightsailLoadBalancerName)
	}

	if resp, err := client.GetLoadBalancer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about health metrics for your Lightsail load balancer.
// Metrics report the utilization of your resources, and the error counts
// generated by them. Monitor and collect metric data regularly to maintain the
// reliability, availability, and performance of your resources.
func lightsail_GetLoadBalancerMetricData(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetLoadBalancerMetricDataInput{
		// EndTime: *time.Time, // Required
		// LoadBalancerName: *string, // Required
		// MetricName: types.LoadBalancerMetricName, // Required
		// Period: *int32, // Required
		// StartTime: *time.Time, // Required
		// Statistics: []types.MetricStatistic, // Required
		// Unit: types.MetricUnit, // Required
	}

	if len(_lightsailEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _lightsailEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_lightsailLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_lightsailLoadBalancerName)
	}
	if len(_lightsailMetricName) > 0 {
		if err := assignInputField(input, "MetricName", _lightsailMetricName); err != nil {
			log.Errorf("invalid --metric-name: %s", err.Error())
			return
		}
	}
	if len(_lightsailPeriod) > 0 {
		if err := assignInputField(input, "Period", _lightsailPeriod); err != nil {
			log.Errorf("invalid --period: %s", err.Error())
			return
		}
	}
	if len(_lightsailStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _lightsailStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_lightsailStatistics) > 0 {
		if err := assignInputField(input, "Statistics", _lightsailStatistics); err != nil {
			log.Errorf("invalid --statistics: %s", err.Error())
			return
		}
	}
	if len(_lightsailUnit) > 0 {
		if err := assignInputField(input, "Unit", _lightsailUnit); err != nil {
			log.Errorf("invalid --unit: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetLoadBalancerMetricData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the TLS certificates that are associated with the
// specified Lightsail load balancer.
//
// TLS is just an updated, more secure version of Secure Socket Layer (SSL).
//
// You can have a maximum of 2 certificates associated with a Lightsail load
// balancer. One is active and the other is inactive.
func lightsail_GetLoadBalancerTlsCertificates(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetLoadBalancerTlsCertificatesInput{
		// LoadBalancerName: *string, // Required
	}

	if len(_lightsailLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_lightsailLoadBalancerName)
	}

	if resp, err := client.GetLoadBalancerTlsCertificates(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of TLS security policies that you can apply to Lightsail load
// balancers.
//
// For more information about load balancer TLS security policies, see [Configuring TLS security policies on your Amazon Lightsail load balancers] in the
// Amazon Lightsail Developer Guide.
//
// [Configuring TLS security policies on your Amazon Lightsail load balancers]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-configure-load-balancer-tls-security-policy
func lightsail_GetLoadBalancerTlsPolicies(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetLoadBalancerTlsPoliciesInput{}

	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetLoadBalancerTlsPolicies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about all load balancers in an account.
func lightsail_GetLoadBalancers(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetLoadBalancersInput{}

	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetLoadBalancers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific operation. Operations include events such
// as when you create an instance, allocate a static IP, attach a static IP, and so
// on.
func lightsail_GetOperation(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetOperationInput{
		// OperationId: *string, // Required
	}

	if len(_lightsailOperationId) > 0 {
		input.OperationId = aws.String(_lightsailOperationId)
	}

	if resp, err := client.GetOperation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about all operations.
// Results are returned from oldest to newest, up to a maximum of 200. Results can
// be paged by making each subsequent call to GetOperations use the maximum (last)
// statusChangedAt value from the previous request.
func lightsail_GetOperations(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetOperationsInput{}

	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetOperations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets operations for a specific resource (an instance or a static IP).
func lightsail_GetOperationsForResource(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetOperationsForResourceInput{
		// ResourceName: *string, // Required
	}

	if len(_lightsailResourceName) > 0 {
		input.ResourceName = aws.String(_lightsailResourceName)
	}
	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetOperationsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of all valid regions for Amazon Lightsail. Use the include
// availability zones parameter to also return the Availability Zones in a region.
func lightsail_GetRegions(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetRegionsInput{}

	if len(_lightsailIncludeAvailabilityZones) > 0 {
		if err := assignInputField(input, "IncludeAvailabilityZones", _lightsailIncludeAvailabilityZones); err != nil {
			log.Errorf("invalid --include-availability-zones: %s", err.Error())
			return
		}
	}
	if len(_lightsailIncludeRelationalDatabaseAvailabilityZones) > 0 {
		if err := assignInputField(input, "IncludeRelationalDatabaseAvailabilityZones", _lightsailIncludeRelationalDatabaseAvailabilityZones); err != nil {
			log.Errorf("invalid --include-relational-database-availability-zones: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetRegions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific database in Amazon Lightsail.
func lightsail_GetRelationalDatabase(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetRelationalDatabaseInput{
		// RelationalDatabaseName: *string, // Required
	}

	if len(_lightsailRelationalDatabaseName) > 0 {
		input.RelationalDatabaseName = aws.String(_lightsailRelationalDatabaseName)
	}

	if resp, err := client.GetRelationalDatabase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of available database blueprints in Amazon Lightsail. A
// blueprint describes the major engine version of a database.
//
// You can use a blueprint ID to create a new database that runs a specific
// database engine.
func lightsail_GetRelationalDatabaseBlueprints(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetRelationalDatabaseBlueprintsInput{}

	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetRelationalDatabaseBlueprints(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the list of bundles that are available in Amazon Lightsail. A bundle
// describes the performance specifications for a database.
//
// You can use a bundle ID to create a new database with explicit performance
// specifications.
func lightsail_GetRelationalDatabaseBundles(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetRelationalDatabaseBundlesInput{}

	if len(_lightsailIncludeInactive) > 0 {
		if err := assignInputField(input, "IncludeInactive", _lightsailIncludeInactive); err != nil {
			log.Errorf("invalid --include-inactive: %s", err.Error())
			return
		}
	}
	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetRelationalDatabaseBundles(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of events for a specific database in Amazon Lightsail.
func lightsail_GetRelationalDatabaseEvents(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetRelationalDatabaseEventsInput{
		// RelationalDatabaseName: *string, // Required
	}

	if len(_lightsailRelationalDatabaseName) > 0 {
		input.RelationalDatabaseName = aws.String(_lightsailRelationalDatabaseName)
	}
	if len(_lightsailDurationInMinutes) > 0 {
		if err := assignInputField(input, "DurationInMinutes", _lightsailDurationInMinutes); err != nil {
			log.Errorf("invalid --duration-in-minutes: %s", err.Error())
			return
		}
	}
	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetRelationalDatabaseEvents(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of log events for a database in Amazon Lightsail.
func lightsail_GetRelationalDatabaseLogEvents(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetRelationalDatabaseLogEventsInput{
		// LogStreamName: *string, // Required
		// RelationalDatabaseName: *string, // Required
	}

	if len(_lightsailLogStreamName) > 0 {
		input.LogStreamName = aws.String(_lightsailLogStreamName)
	}
	if len(_lightsailRelationalDatabaseName) > 0 {
		input.RelationalDatabaseName = aws.String(_lightsailRelationalDatabaseName)
	}
	if len(_lightsailEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _lightsailEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}
	if len(_lightsailStartFromHead) > 0 {
		if err := assignInputField(input, "StartFromHead", _lightsailStartFromHead); err != nil {
			log.Errorf("invalid --start-from-head: %s", err.Error())
			return
		}
	}
	if len(_lightsailStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _lightsailStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetRelationalDatabaseLogEvents(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of available log streams for a specific database in Amazon
// Lightsail.
func lightsail_GetRelationalDatabaseLogStreams(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetRelationalDatabaseLogStreamsInput{
		// RelationalDatabaseName: *string, // Required
	}

	if len(_lightsailRelationalDatabaseName) > 0 {
		input.RelationalDatabaseName = aws.String(_lightsailRelationalDatabaseName)
	}

	if resp, err := client.GetRelationalDatabaseLogStreams(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the current, previous, or pending versions of the master user password
// for a Lightsail database.
//
// The GetRelationalDatabaseMasterUserPassword operation supports tag-based access
// control via resource tags applied to the resource identified by
// relationalDatabaseName.
func lightsail_GetRelationalDatabaseMasterUserPassword(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetRelationalDatabaseMasterUserPasswordInput{
		// RelationalDatabaseName: *string, // Required
	}

	if len(_lightsailRelationalDatabaseName) > 0 {
		input.RelationalDatabaseName = aws.String(_lightsailRelationalDatabaseName)
	}
	if len(_lightsailPasswordVersion) > 0 {
		if err := assignInputField(input, "PasswordVersion", _lightsailPasswordVersion); err != nil {
			log.Errorf("invalid --password-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetRelationalDatabaseMasterUserPassword(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the data points of the specified metric for a database in Amazon
// Lightsail.
//
// Metrics report the utilization of your resources, and the error counts
// generated by them. Monitor and collect metric data regularly to maintain the
// reliability, availability, and performance of your resources.
func lightsail_GetRelationalDatabaseMetricData(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetRelationalDatabaseMetricDataInput{
		// EndTime: *time.Time, // Required
		// MetricName: types.RelationalDatabaseMetricName, // Required
		// Period: *int32, // Required
		// RelationalDatabaseName: *string, // Required
		// StartTime: *time.Time, // Required
		// Statistics: []types.MetricStatistic, // Required
		// Unit: types.MetricUnit, // Required
	}

	if len(_lightsailEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _lightsailEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_lightsailMetricName) > 0 {
		if err := assignInputField(input, "MetricName", _lightsailMetricName); err != nil {
			log.Errorf("invalid --metric-name: %s", err.Error())
			return
		}
	}
	if len(_lightsailPeriod) > 0 {
		if err := assignInputField(input, "Period", _lightsailPeriod); err != nil {
			log.Errorf("invalid --period: %s", err.Error())
			return
		}
	}
	if len(_lightsailRelationalDatabaseName) > 0 {
		input.RelationalDatabaseName = aws.String(_lightsailRelationalDatabaseName)
	}
	if len(_lightsailStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _lightsailStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_lightsailStatistics) > 0 {
		if err := assignInputField(input, "Statistics", _lightsailStatistics); err != nil {
			log.Errorf("invalid --statistics: %s", err.Error())
			return
		}
	}
	if len(_lightsailUnit) > 0 {
		if err := assignInputField(input, "Unit", _lightsailUnit); err != nil {
			log.Errorf("invalid --unit: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetRelationalDatabaseMetricData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns all of the runtime parameters offered by the underlying database
// software, or engine, for a specific database in Amazon Lightsail.
//
// In addition to the parameter names and values, this operation returns other
// information about each parameter. This information includes whether changes
// require a reboot, whether the parameter is modifiable, the allowed values, and
// the data types.
func lightsail_GetRelationalDatabaseParameters(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetRelationalDatabaseParametersInput{
		// RelationalDatabaseName: *string, // Required
	}

	if len(_lightsailRelationalDatabaseName) > 0 {
		input.RelationalDatabaseName = aws.String(_lightsailRelationalDatabaseName)
	}
	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetRelationalDatabaseParameters(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific database snapshot in Amazon Lightsail.
func lightsail_GetRelationalDatabaseSnapshot(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetRelationalDatabaseSnapshotInput{
		// RelationalDatabaseSnapshotName: *string, // Required
	}

	if len(_lightsailRelationalDatabaseSnapshotName) > 0 {
		input.RelationalDatabaseSnapshotName = aws.String(_lightsailRelationalDatabaseSnapshotName)
	}

	if resp, err := client.GetRelationalDatabaseSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about all of your database snapshots in Amazon Lightsail.
func lightsail_GetRelationalDatabaseSnapshots(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetRelationalDatabaseSnapshotsInput{}

	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetRelationalDatabaseSnapshots(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about all of your databases in Amazon Lightsail.
func lightsail_GetRelationalDatabases(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetRelationalDatabasesInput{}

	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetRelationalDatabases(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns detailed information for five of the most recent SetupInstanceHttps
// requests that were ran on the target instance.
func lightsail_GetSetupHistory(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetSetupHistoryInput{
		// ResourceName: *string, // Required
	}

	if len(_lightsailResourceName) > 0 {
		input.ResourceName = aws.String(_lightsailResourceName)
	}
	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetSetupHistory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an Amazon Lightsail static IP.
func lightsail_GetStaticIp(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetStaticIpInput{
		// StaticIpName: *string, // Required
	}

	if len(_lightsailStaticIpName) > 0 {
		input.StaticIpName = aws.String(_lightsailStaticIpName)
	}

	if resp, err := client.GetStaticIp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about all static IPs in the user's account.
func lightsail_GetStaticIps(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.GetStaticIpsInput{}

	if len(_lightsailPageToken) > 0 {
		input.PageToken = aws.String(_lightsailPageToken)
	}

	if resp, err := client.GetStaticIps(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports a public SSH key from a specific key pair.
func lightsail_ImportKeyPair(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.ImportKeyPairInput{
		// KeyPairName: *string, // Required
		// PublicKeyBase64: *string, // Required
	}

	if len(_lightsailKeyPairName) > 0 {
		input.KeyPairName = aws.String(_lightsailKeyPairName)
	}
	if len(_lightsailPublicKeyBase64) > 0 {
		input.PublicKeyBase64 = aws.String(_lightsailPublicKeyBase64)
	}

	if resp, err := client.ImportKeyPair(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a Boolean value indicating whether your Lightsail VPC is peered.
func lightsail_IsVpcPeered(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.IsVpcPeeredInput{}

	if resp, err := client.IsVpcPeered(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Opens ports for a specific Amazon Lightsail instance, and specifies the IP
// addresses allowed to connect to the instance through the ports, and the
// protocol.
//
// The OpenInstancePublicPorts action supports tag-based access control via
// resource tags applied to the resource identified by instanceName . For more
// information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_OpenInstancePublicPorts(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.OpenInstancePublicPortsInput{
		// InstanceName: *string, // Required
		// PortInfo: *types.PortInfo, // Required
	}

	if len(_lightsailInstanceName) > 0 {
		input.InstanceName = aws.String(_lightsailInstanceName)
	}
	if len(_lightsailPortInfo) > 0 {
		if err := assignInputField(input, "PortInfo", _lightsailPortInfo); err != nil {
			log.Errorf("invalid --port-info: %s", err.Error())
			return
		}
	}

	if resp, err := client.OpenInstancePublicPorts(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Peers the Lightsail VPC with the user's default VPC.
func lightsail_PeerVpc(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.PeerVpcInput{}

	if resp, err := client.PeerVpc(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates an alarm, and associates it with the specified metric.
// An alarm is used to monitor a single metric for one of your resources. When a
// metric condition is met, the alarm can notify you by email, SMS text message,
// and a banner displayed on the Amazon Lightsail console. For more information,
// see [Alarms in Amazon Lightsail].
//
// When this action creates an alarm, the alarm state is immediately set to
// INSUFFICIENT_DATA . The alarm is then evaluated and its state is set
// appropriately. Any actions associated with the new state are then executed.
//
// When you update an existing alarm, its state is left unchanged, but the update
// completely overwrites the previous configuration of the alarm. The alarm is then
// evaluated with the updated configuration.
//
// [Alarms in Amazon Lightsail]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-alarms
func lightsail_PutAlarm(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.PutAlarmInput{
		// AlarmName: *string, // Required
		// ComparisonOperator: types.ComparisonOperator, // Required
		// EvaluationPeriods: *int32, // Required
		// MetricName: types.MetricName, // Required
		// MonitoredResourceName: *string, // Required
		// Threshold: *float64, // Required
	}

	if len(_lightsailAlarmName) > 0 {
		input.AlarmName = aws.String(_lightsailAlarmName)
	}
	if len(_lightsailComparisonOperator) > 0 {
		if err := assignInputField(input, "ComparisonOperator", _lightsailComparisonOperator); err != nil {
			log.Errorf("invalid --comparison-operator: %s", err.Error())
			return
		}
	}
	if len(_lightsailEvaluationPeriods) > 0 {
		if err := assignInputField(input, "EvaluationPeriods", _lightsailEvaluationPeriods); err != nil {
			log.Errorf("invalid --evaluation-periods: %s", err.Error())
			return
		}
	}
	if len(_lightsailMetricName) > 0 {
		if err := assignInputField(input, "MetricName", _lightsailMetricName); err != nil {
			log.Errorf("invalid --metric-name: %s", err.Error())
			return
		}
	}
	if len(_lightsailMonitoredResourceName) > 0 {
		input.MonitoredResourceName = aws.String(_lightsailMonitoredResourceName)
	}
	if len(_lightsailThreshold) > 0 {
		if err := assignInputField(input, "Threshold", _lightsailThreshold); err != nil {
			log.Errorf("invalid --threshold: %s", err.Error())
			return
		}
	}
	if len(_lightsailContactProtocols) > 0 {
		if err := assignInputField(input, "ContactProtocols", _lightsailContactProtocols); err != nil {
			log.Errorf("invalid --contact-protocols: %s", err.Error())
			return
		}
	}
	if len(_lightsailDatapointsToAlarm) > 0 {
		if err := assignInputField(input, "DatapointsToAlarm", _lightsailDatapointsToAlarm); err != nil {
			log.Errorf("invalid --datapoints-to-alarm: %s", err.Error())
			return
		}
	}
	if len(_lightsailNotificationEnabled) > 0 {
		if err := assignInputField(input, "NotificationEnabled", _lightsailNotificationEnabled); err != nil {
			log.Errorf("invalid --notification-enabled: %s", err.Error())
			return
		}
	}
	if len(_lightsailNotificationTriggers) > 0 {
		if err := assignInputField(input, "NotificationTriggers", _lightsailNotificationTriggers); err != nil {
			log.Errorf("invalid --notification-triggers: %s", err.Error())
			return
		}
	}
	if len(_lightsailTreatMissingData) > 0 {
		if err := assignInputField(input, "TreatMissingData", _lightsailTreatMissingData); err != nil {
			log.Errorf("invalid --treat-missing-data: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutAlarm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Opens ports for a specific Amazon Lightsail instance, and specifies the IP
// addresses allowed to connect to the instance through the ports, and the
// protocol. This action also closes all currently open ports that are not included
// in the request. Include all of the ports and the protocols you want to open in
// your PutInstancePublicPorts request. Or use the OpenInstancePublicPorts action
// to open ports without closing currently open ports.
//
// The PutInstancePublicPorts action supports tag-based access control via
// resource tags applied to the resource identified by instanceName . For more
// information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_PutInstancePublicPorts(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.PutInstancePublicPortsInput{
		// InstanceName: *string, // Required
		// PortInfos: []types.PortInfo, // Required
	}

	if len(_lightsailInstanceName) > 0 {
		input.InstanceName = aws.String(_lightsailInstanceName)
	}
	if len(_lightsailPortInfos) > 0 {
		if err := assignInputField(input, "PortInfos", _lightsailPortInfos); err != nil {
			log.Errorf("invalid --port-infos: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutInstancePublicPorts(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restarts a specific instance.
// The reboot instance operation supports tag-based access control via resource
// tags applied to the resource identified by instance name . For more information,
// see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_RebootInstance(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.RebootInstanceInput{
		// InstanceName: *string, // Required
	}

	if len(_lightsailInstanceName) > 0 {
		input.InstanceName = aws.String(_lightsailInstanceName)
	}

	if resp, err := client.RebootInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restarts a specific database in Amazon Lightsail.
// The reboot relational database operation supports tag-based access control via
// resource tags applied to the resource identified by relationalDatabaseName. For
// more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_RebootRelationalDatabase(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.RebootRelationalDatabaseInput{
		// RelationalDatabaseName: *string, // Required
	}

	if len(_lightsailRelationalDatabaseName) > 0 {
		input.RelationalDatabaseName = aws.String(_lightsailRelationalDatabaseName)
	}

	if resp, err := client.RebootRelationalDatabase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers a container image to your Amazon Lightsail container service.
// This action is not required if you install and use the Lightsail Control
// (lightsailctl) plugin to push container images to your Lightsail container
// service. For more information, see [Pushing and managing container images on your Amazon Lightsail container services]in the Amazon Lightsail Developer Guide.
//
// [Pushing and managing container images on your Amazon Lightsail container services]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-pushing-container-images
func lightsail_RegisterContainerImage(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.RegisterContainerImageInput{
		// Digest: *string, // Required
		// Label: *string, // Required
		// ServiceName: *string, // Required
	}

	if len(_lightsailDigest) > 0 {
		input.Digest = aws.String(_lightsailDigest)
	}
	if len(_lightsailLabel) > 0 {
		input.Label = aws.String(_lightsailLabel)
	}
	if len(_lightsailServiceName) > 0 {
		input.ServiceName = aws.String(_lightsailServiceName)
	}

	if resp, err := client.RegisterContainerImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specific static IP from your account.
func lightsail_ReleaseStaticIp(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.ReleaseStaticIpInput{
		// StaticIpName: *string, // Required
	}

	if len(_lightsailStaticIpName) > 0 {
		input.StaticIpName = aws.String(_lightsailStaticIpName)
	}

	if resp, err := client.ReleaseStaticIp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes currently cached content from your Amazon Lightsail content delivery
// network (CDN) distribution.
//
// After resetting the cache, the next time a content request is made, your
// distribution pulls, serves, and caches it from the origin.
func lightsail_ResetDistributionCache(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.ResetDistributionCacheInput{}

	if len(_lightsailDistributionName) > 0 {
		input.DistributionName = aws.String(_lightsailDistributionName)
	}

	if resp, err := client.ResetDistributionCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends a verification request to an email contact method to ensure it's owned by
// the requester. SMS contact methods don't need to be verified.
//
// A contact method is used to send you notifications about your Amazon Lightsail
// resources. You can add one email address and one mobile phone number contact
// method in each Amazon Web Services Region. However, SMS text messaging is not
// supported in some Amazon Web Services Regions, and SMS text messages cannot be
// sent to some countries/regions. For more information, see [Notifications in Amazon Lightsail].
//
// A verification request is sent to the contact method when you initially create
// it. Use this action to send another verification request if a previous
// verification request was deleted, or has expired.
//
// Notifications are not sent to an email contact method until after it is
// verified, and confirmed as valid.
//
// [Notifications in Amazon Lightsail]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-notifications
func lightsail_SendContactMethodVerification(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.SendContactMethodVerificationInput{
		// Protocol: types.ContactMethodVerificationProtocol, // Required
	}

	if len(_lightsailProtocol) > 0 {
		if err := assignInputField(input, "Protocol", _lightsailProtocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendContactMethodVerification(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the IP address type for an Amazon Lightsail resource.
// Use this action to enable dual-stack for a resource, which enables IPv4 and
// IPv6 for the specified resource. Alternately, you can use this action to disable
// dual-stack, and enable IPv4 only.
func lightsail_SetIpAddressType(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.SetIpAddressTypeInput{
		// IpAddressType: types.IpAddressType, // Required
		// ResourceName: *string, // Required
		// ResourceType: types.ResourceType, // Required
	}

	if len(_lightsailIpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _lightsailIpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_lightsailResourceName) > 0 {
		input.ResourceName = aws.String(_lightsailResourceName)
	}
	if len(_lightsailResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _lightsailResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_lightsailAcceptBundleUpdate) > 0 {
		if err := assignInputField(input, "AcceptBundleUpdate", _lightsailAcceptBundleUpdate); err != nil {
			log.Errorf("invalid --accept-bundle-update: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetIpAddressType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the Amazon Lightsail resources that can access the specified Lightsail
// bucket.
//
// Lightsail buckets currently support setting access for Lightsail instances in
// the same Amazon Web Services Region.
func lightsail_SetResourceAccessForBucket(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.SetResourceAccessForBucketInput{
		// Access: types.ResourceBucketAccess, // Required
		// BucketName: *string, // Required
		// ResourceName: *string, // Required
	}

	if len(_lightsailAccess) > 0 {
		if err := assignInputField(input, "Access", _lightsailAccess); err != nil {
			log.Errorf("invalid --access: %s", err.Error())
			return
		}
	}
	if len(_lightsailBucketName) > 0 {
		input.BucketName = aws.String(_lightsailBucketName)
	}
	if len(_lightsailResourceName) > 0 {
		input.ResourceName = aws.String(_lightsailResourceName)
	}

	if resp, err := client.SetResourceAccessForBucket(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an SSL/TLS certificate that secures traffic for your website. After the
// certificate is created, it is installed on the specified Lightsail instance.
//
// If you provide more than one domain name in the request, at least one name must
// be less than or equal to 63 characters in length.
func lightsail_SetupInstanceHttps(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.SetupInstanceHttpsInput{
		// CertificateProvider: types.CertificateProvider, // Required
		// DomainNames: []string, // Required
		// EmailAddress: *string, // Required
		// InstanceName: *string, // Required
	}

	if len(_lightsailCertificateProvider) > 0 {
		if err := assignInputField(input, "CertificateProvider", _lightsailCertificateProvider); err != nil {
			log.Errorf("invalid --certificate-provider: %s", err.Error())
			return
		}
	}
	if len(_lightsailDomainNames) > 0 {
		input.DomainNames = append([]string(nil), _lightsailDomainNames...)
	}
	if len(_lightsailEmailAddress) > 0 {
		input.EmailAddress = aws.String(_lightsailEmailAddress)
	}
	if len(_lightsailInstanceName) > 0 {
		input.InstanceName = aws.String(_lightsailInstanceName)
	}

	if resp, err := client.SetupInstanceHttps(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a graphical user interface (GUI) session that’s used to access a
// virtual computer’s operating system and application. The session will be active
// for 1 hour. Use this action to resume the session after it expires.
func lightsail_StartGUISession(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.StartGUISessionInput{
		// ResourceName: *string, // Required
	}

	if len(_lightsailResourceName) > 0 {
		input.ResourceName = aws.String(_lightsailResourceName)
	}

	if resp, err := client.StartGUISession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a specific Amazon Lightsail instance from a stopped state. To restart an
// instance, use the reboot instance operation.
//
// When you start a stopped instance, Lightsail assigns a new public IP address to
// the instance. To use the same IP address after stopping and starting an
// instance, create a static IP address and attach it to the instance. For more
// information, see the [Amazon Lightsail Developer Guide].
//
// The start instance operation supports tag-based access control via resource
// tags applied to the resource identified by instance name . For more information,
// see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_StartInstance(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.StartInstanceInput{
		// InstanceName: *string, // Required
	}

	if len(_lightsailInstanceName) > 0 {
		input.InstanceName = aws.String(_lightsailInstanceName)
	}

	if resp, err := client.StartInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a specific database from a stopped state in Amazon Lightsail. To restart
// a database, use the reboot relational database operation.
//
// The start relational database operation supports tag-based access control via
// resource tags applied to the resource identified by relationalDatabaseName. For
// more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_StartRelationalDatabase(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.StartRelationalDatabaseInput{
		// RelationalDatabaseName: *string, // Required
	}

	if len(_lightsailRelationalDatabaseName) > 0 {
		input.RelationalDatabaseName = aws.String(_lightsailRelationalDatabaseName)
	}

	if resp, err := client.StartRelationalDatabase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Terminates a web-based Amazon DCV session that’s used to access a virtual
// computer’s operating system or application. The session will close and any
// unsaved data will be lost.
func lightsail_StopGUISession(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.StopGUISessionInput{
		// ResourceName: *string, // Required
	}

	if len(_lightsailResourceName) > 0 {
		input.ResourceName = aws.String(_lightsailResourceName)
	}

	if resp, err := client.StopGUISession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a specific Amazon Lightsail instance that is currently running.
// When you start a stopped instance, Lightsail assigns a new public IP address to
// the instance. To use the same IP address after stopping and starting an
// instance, create a static IP address and attach it to the instance. For more
// information, see the [Amazon Lightsail Developer Guide].
//
// The stop instance operation supports tag-based access control via resource tags
// applied to the resource identified by instance name . For more information, see
// the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_StopInstance(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.StopInstanceInput{
		// InstanceName: *string, // Required
	}

	if len(_lightsailInstanceName) > 0 {
		input.InstanceName = aws.String(_lightsailInstanceName)
	}
	if len(_lightsailForce) > 0 {
		if err := assignInputField(input, "Force", _lightsailForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}

	if resp, err := client.StopInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a specific database that is currently running in Amazon Lightsail.
// If you don't manually start your database instance after it has been stopped
// for seven consecutive days, Amazon Lightsail automatically starts it for you.
// This action helps ensure that your database instance doesn't fall behind on any
// required maintenance updates.
//
// The stop relational database operation supports tag-based access control via
// resource tags applied to the resource identified by relationalDatabaseName. For
// more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_StopRelationalDatabase(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.StopRelationalDatabaseInput{
		// RelationalDatabaseName: *string, // Required
	}

	if len(_lightsailRelationalDatabaseName) > 0 {
		input.RelationalDatabaseName = aws.String(_lightsailRelationalDatabaseName)
	}
	if len(_lightsailRelationalDatabaseSnapshotName) > 0 {
		input.RelationalDatabaseSnapshotName = aws.String(_lightsailRelationalDatabaseSnapshotName)
	}

	if resp, err := client.StopRelationalDatabase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more tags to the specified Amazon Lightsail resource. Each resource
// can have a maximum of 50 tags. Each tag consists of a key and an optional value.
// Tag keys must be unique per resource. For more information about tags, see the [Amazon Lightsail Developer Guide].
//
// The tag resource operation supports tag-based access control via request tags
// and resource tags applied to the resource identified by resource name . For more
// information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_TagResource(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.TagResourceInput{
		// ResourceName: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_lightsailResourceName) > 0 {
		input.ResourceName = aws.String(_lightsailResourceName)
	}
	if len(_lightsailTags) > 0 {
		if err := assignInputField(input, "Tags", _lightsailTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_lightsailResourceArn) > 0 {
		input.ResourceArn = aws.String(_lightsailResourceArn)
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tests an alarm by displaying a banner on the Amazon Lightsail console. If a
// notification trigger is configured for the specified alarm, the test also sends
// a notification to the notification protocol ( Email and/or SMS ) configured for
// the alarm.
//
// An alarm is used to monitor a single metric for one of your resources. When a
// metric condition is met, the alarm can notify you by email, SMS text message,
// and a banner displayed on the Amazon Lightsail console. For more information,
// see [Alarms in Amazon Lightsail].
//
// [Alarms in Amazon Lightsail]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-alarms
func lightsail_TestAlarm(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.TestAlarmInput{
		// AlarmName: *string, // Required
		// State: types.AlarmState, // Required
	}

	if len(_lightsailAlarmName) > 0 {
		input.AlarmName = aws.String(_lightsailAlarmName)
	}
	if len(_lightsailState) > 0 {
		if err := assignInputField(input, "State", _lightsailState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if resp, err := client.TestAlarm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Unpeers the Lightsail VPC from the user's default VPC.
func lightsail_UnpeerVpc(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.UnpeerVpcInput{}

	if resp, err := client.UnpeerVpc(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified set of tag keys and their values from the specified
// Amazon Lightsail resource.
//
// The untag resource operation supports tag-based access control via request tags
// and resource tags applied to the resource identified by resource name . For more
// information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_UntagResource(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.UntagResourceInput{
		// ResourceName: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_lightsailResourceName) > 0 {
		input.ResourceName = aws.String(_lightsailResourceName)
	}
	if len(_lightsailTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _lightsailTagKeys...)
	}
	if len(_lightsailResourceArn) > 0 {
		input.ResourceArn = aws.String(_lightsailResourceArn)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing Amazon Lightsail bucket.
// Use this action to update the configuration of an existing bucket, such as
// versioning, public accessibility, and the Amazon Web Services accounts that can
// access the bucket.
func lightsail_UpdateBucket(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.UpdateBucketInput{
		// BucketName: *string, // Required
	}

	if len(_lightsailBucketName) > 0 {
		input.BucketName = aws.String(_lightsailBucketName)
	}
	if len(_lightsailAccessLogConfig) > 0 {
		if err := assignInputField(input, "AccessLogConfig", _lightsailAccessLogConfig); err != nil {
			log.Errorf("invalid --access-log-config: %s", err.Error())
			return
		}
	}
	if len(_lightsailAccessRules) > 0 {
		if err := assignInputField(input, "AccessRules", _lightsailAccessRules); err != nil {
			log.Errorf("invalid --access-rules: %s", err.Error())
			return
		}
	}
	if len(_lightsailCors) > 0 {
		if err := assignInputField(input, "Cors", _lightsailCors); err != nil {
			log.Errorf("invalid --cors: %s", err.Error())
			return
		}
	}
	if len(_lightsailReadonlyAccessAccounts) > 0 {
		input.ReadonlyAccessAccounts = append([]string(nil), _lightsailReadonlyAccessAccounts...)
	}
	if len(_lightsailVersioning) > 0 {
		input.Versioning = aws.String(_lightsailVersioning)
	}

	if resp, err := client.UpdateBucket(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the bundle, or storage plan, of an existing Amazon Lightsail bucket.
// A bucket bundle specifies the monthly cost, storage space, and data transfer
// quota for a bucket. You can update a bucket's bundle only one time within a
// monthly Amazon Web Services billing cycle. To determine if you can update a
// bucket's bundle, use the [GetBuckets]action. The ableToUpdateBundle parameter in the
// response will indicate whether you can currently update a bucket's bundle.
//
// Update a bucket's bundle if it's consistently going over its storage space or
// data transfer quota, or if a bucket's usage is consistently in the lower range
// of its storage space or data transfer quota. Due to the unpredictable usage
// fluctuations that a bucket might experience, we strongly recommend that you
// update a bucket's bundle only as a long-term strategy, instead of as a
// short-term, monthly cost-cutting measure. Choose a bucket bundle that will
// provide the bucket with ample storage space and data transfer for a long time to
// come.
//
// [GetBuckets]: https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_GetBuckets.html
func lightsail_UpdateBucketBundle(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.UpdateBucketBundleInput{
		// BucketName: *string, // Required
		// BundleId: *string, // Required
	}

	if len(_lightsailBucketName) > 0 {
		input.BucketName = aws.String(_lightsailBucketName)
	}
	if len(_lightsailBundleId) > 0 {
		input.BundleId = aws.String(_lightsailBundleId)
	}

	if resp, err := client.UpdateBucketBundle(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of your Amazon Lightsail container service, such as
// its power, scale, and public domain names.
func lightsail_UpdateContainerService(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.UpdateContainerServiceInput{
		// ServiceName: *string, // Required
	}

	if len(_lightsailServiceName) > 0 {
		input.ServiceName = aws.String(_lightsailServiceName)
	}
	if len(_lightsailIsDisabled) > 0 {
		if err := assignInputField(input, "IsDisabled", _lightsailIsDisabled); err != nil {
			log.Errorf("invalid --is-disabled: %s", err.Error())
			return
		}
	}
	if len(_lightsailPower) > 0 {
		if err := assignInputField(input, "Power", _lightsailPower); err != nil {
			log.Errorf("invalid --power: %s", err.Error())
			return
		}
	}
	if len(_lightsailPrivateRegistryAccess) > 0 {
		if err := assignInputField(input, "PrivateRegistryAccess", _lightsailPrivateRegistryAccess); err != nil {
			log.Errorf("invalid --private-registry-access: %s", err.Error())
			return
		}
	}
	if len(_lightsailPublicDomainNames) > 0 {
		if err := assignInputField(input, "PublicDomainNames", _lightsailPublicDomainNames); err != nil {
			log.Errorf("invalid --public-domain-names: %s", err.Error())
			return
		}
	}
	if len(_lightsailScale) > 0 {
		if err := assignInputField(input, "Scale", _lightsailScale); err != nil {
			log.Errorf("invalid --scale: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateContainerService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing Amazon Lightsail content delivery network (CDN)
// distribution.
//
// Use this action to update the configuration of your existing distribution.
func lightsail_UpdateDistribution(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.UpdateDistributionInput{
		// DistributionName: *string, // Required
	}

	if len(_lightsailDistributionName) > 0 {
		input.DistributionName = aws.String(_lightsailDistributionName)
	}
	if len(_lightsailCacheBehaviorSettings) > 0 {
		if err := assignInputField(input, "CacheBehaviorSettings", _lightsailCacheBehaviorSettings); err != nil {
			log.Errorf("invalid --cache-behavior-settings: %s", err.Error())
			return
		}
	}
	if len(_lightsailCacheBehaviors) > 0 {
		if err := assignInputField(input, "CacheBehaviors", _lightsailCacheBehaviors); err != nil {
			log.Errorf("invalid --cache-behaviors: %s", err.Error())
			return
		}
	}
	if len(_lightsailCertificateName) > 0 {
		input.CertificateName = aws.String(_lightsailCertificateName)
	}
	if len(_lightsailDefaultCacheBehavior) > 0 {
		if err := assignInputField(input, "DefaultCacheBehavior", _lightsailDefaultCacheBehavior); err != nil {
			log.Errorf("invalid --default-cache-behavior: %s", err.Error())
			return
		}
	}
	if len(_lightsailIsEnabled) > 0 {
		if err := assignInputField(input, "IsEnabled", _lightsailIsEnabled); err != nil {
			log.Errorf("invalid --is-enabled: %s", err.Error())
			return
		}
	}
	if len(_lightsailOrigin) > 0 {
		if err := assignInputField(input, "Origin", _lightsailOrigin); err != nil {
			log.Errorf("invalid --origin: %s", err.Error())
			return
		}
	}
	if len(_lightsailUseDefaultCertificate) > 0 {
		if err := assignInputField(input, "UseDefaultCertificate", _lightsailUseDefaultCertificate); err != nil {
			log.Errorf("invalid --use-default-certificate: %s", err.Error())
			return
		}
	}
	if len(_lightsailViewerMinimumTlsProtocolVersion) > 0 {
		if err := assignInputField(input, "ViewerMinimumTlsProtocolVersion", _lightsailViewerMinimumTlsProtocolVersion); err != nil {
			log.Errorf("invalid --viewer-minimum-tls-protocol-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDistribution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the bundle of your Amazon Lightsail content delivery network (CDN)
// distribution.
//
// A distribution bundle specifies the monthly network transfer quota and monthly
// cost of your distribution.
//
// Update your distribution's bundle if your distribution is going over its
// monthly network transfer quota and is incurring an overage fee.
//
// You can update your distribution's bundle only one time within your monthly
// Amazon Web Services billing cycle. To determine if you can update your
// distribution's bundle, use the GetDistributions action. The ableToUpdateBundle
// parameter in the result will indicate whether you can currently update your
// distribution's bundle.
func lightsail_UpdateDistributionBundle(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.UpdateDistributionBundleInput{}

	if len(_lightsailBundleId) > 0 {
		input.BundleId = aws.String(_lightsailBundleId)
	}
	if len(_lightsailDistributionName) > 0 {
		input.DistributionName = aws.String(_lightsailDistributionName)
	}

	if resp, err := client.UpdateDistributionBundle(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a domain recordset after it is created.
// The update domain entry operation supports tag-based access control via
// resource tags applied to the resource identified by domain name . For more
// information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_UpdateDomainEntry(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.UpdateDomainEntryInput{
		// DomainEntry: *types.DomainEntry, // Required
		// DomainName: *string, // Required
	}

	if len(_lightsailDomainEntry) > 0 {
		if err := assignInputField(input, "DomainEntry", _lightsailDomainEntry); err != nil {
			log.Errorf("invalid --domain-entry: %s", err.Error())
			return
		}
	}
	if len(_lightsailDomainName) > 0 {
		input.DomainName = aws.String(_lightsailDomainName)
	}

	if resp, err := client.UpdateDomainEntry(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the Amazon Lightsail instance metadata parameters on a running or
// stopped instance. When you modify the parameters on a running instance, the
// GetInstance or GetInstances API operation initially responds with a state of
// pending . After the parameter modifications are successfully applied, the state
// changes to applied in subsequent GetInstance or GetInstances API calls. For
// more information, see [Use IMDSv2 with an Amazon Lightsail instance]in the Amazon Lightsail Developer Guide.
//
// [Use IMDSv2 with an Amazon Lightsail instance]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-configuring-instance-metadata-service
func lightsail_UpdateInstanceMetadataOptions(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.UpdateInstanceMetadataOptionsInput{
		// InstanceName: *string, // Required
	}

	if len(_lightsailInstanceName) > 0 {
		input.InstanceName = aws.String(_lightsailInstanceName)
	}
	if len(_lightsailHttpEndpoint) > 0 {
		if err := assignInputField(input, "HttpEndpoint", _lightsailHttpEndpoint); err != nil {
			log.Errorf("invalid --http-endpoint: %s", err.Error())
			return
		}
	}
	if len(_lightsailHttpProtocolIpv6) > 0 {
		if err := assignInputField(input, "HttpProtocolIpv6", _lightsailHttpProtocolIpv6); err != nil {
			log.Errorf("invalid --http-protocol-ipv6: %s", err.Error())
			return
		}
	}
	if len(_lightsailHttpPutResponseHopLimit) > 0 {
		if err := assignInputField(input, "HttpPutResponseHopLimit", _lightsailHttpPutResponseHopLimit); err != nil {
			log.Errorf("invalid --http-put-response-hop-limit: %s", err.Error())
			return
		}
	}
	if len(_lightsailHttpTokens) > 0 {
		if err := assignInputField(input, "HttpTokens", _lightsailHttpTokens); err != nil {
			log.Errorf("invalid --http-tokens: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateInstanceMetadataOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified attribute for a load balancer. You can only update one
// attribute at a time.
//
// The update load balancer attribute operation supports tag-based access control
// via resource tags applied to the resource identified by load balancer name . For
// more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_UpdateLoadBalancerAttribute(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.UpdateLoadBalancerAttributeInput{
		// AttributeName: types.LoadBalancerAttributeName, // Required
		// AttributeValue: *string, // Required
		// LoadBalancerName: *string, // Required
	}

	if len(_lightsailAttributeName) > 0 {
		if err := assignInputField(input, "AttributeName", _lightsailAttributeName); err != nil {
			log.Errorf("invalid --attribute-name: %s", err.Error())
			return
		}
	}
	if len(_lightsailAttributeValue) > 0 {
		input.AttributeValue = aws.String(_lightsailAttributeValue)
	}
	if len(_lightsailLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_lightsailLoadBalancerName)
	}

	if resp, err := client.UpdateLoadBalancerAttribute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows the update of one or more attributes of a database in Amazon Lightsail.
// Updates are applied immediately, or in cases where the updates could result in
// an outage, are applied during the database's predefined maintenance window.
//
// The update relational database operation supports tag-based access control via
// resource tags applied to the resource identified by relationalDatabaseName. For
// more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_UpdateRelationalDatabase(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.UpdateRelationalDatabaseInput{
		// RelationalDatabaseName: *string, // Required
	}

	if len(_lightsailRelationalDatabaseName) > 0 {
		input.RelationalDatabaseName = aws.String(_lightsailRelationalDatabaseName)
	}
	if len(_lightsailApplyImmediately) > 0 {
		if err := assignInputField(input, "ApplyImmediately", _lightsailApplyImmediately); err != nil {
			log.Errorf("invalid --apply-immediately: %s", err.Error())
			return
		}
	}
	if len(_lightsailCaCertificateIdentifier) > 0 {
		input.CaCertificateIdentifier = aws.String(_lightsailCaCertificateIdentifier)
	}
	if len(_lightsailDisableBackupRetention) > 0 {
		if err := assignInputField(input, "DisableBackupRetention", _lightsailDisableBackupRetention); err != nil {
			log.Errorf("invalid --disable-backup-retention: %s", err.Error())
			return
		}
	}
	if len(_lightsailEnableBackupRetention) > 0 {
		if err := assignInputField(input, "EnableBackupRetention", _lightsailEnableBackupRetention); err != nil {
			log.Errorf("invalid --enable-backup-retention: %s", err.Error())
			return
		}
	}
	if len(_lightsailMasterUserPassword) > 0 {
		input.MasterUserPassword = aws.String(_lightsailMasterUserPassword)
	}
	if len(_lightsailPreferredBackupWindow) > 0 {
		input.PreferredBackupWindow = aws.String(_lightsailPreferredBackupWindow)
	}
	if len(_lightsailPreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_lightsailPreferredMaintenanceWindow)
	}
	if len(_lightsailPubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _lightsailPubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_lightsailRelationalDatabaseBlueprintId) > 0 {
		input.RelationalDatabaseBlueprintId = aws.String(_lightsailRelationalDatabaseBlueprintId)
	}
	if len(_lightsailRotateMasterUserPassword) > 0 {
		if err := assignInputField(input, "RotateMasterUserPassword", _lightsailRotateMasterUserPassword); err != nil {
			log.Errorf("invalid --rotate-master-user-password: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRelationalDatabase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows the update of one or more parameters of a database in Amazon Lightsail.
// Parameter updates don't cause outages; therefore, their application is not
// subject to the preferred maintenance window. However, there are two ways in
// which parameter updates are applied: dynamic or pending-reboot . Parameters
// marked with a dynamic apply type are applied immediately. Parameters marked
// with a pending-reboot apply type are applied only after the database is
// rebooted using the reboot relational database operation.
//
// The update relational database parameters operation supports tag-based access
// control via resource tags applied to the resource identified by
// relationalDatabaseName. For more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func lightsail_UpdateRelationalDatabaseParameters(cfg aws.Config, client *lightsail.Client) {
	input := &lightsail.UpdateRelationalDatabaseParametersInput{
		// Parameters: []types.RelationalDatabaseParameter, // Required
		// RelationalDatabaseName: *string, // Required
	}

	if len(_lightsailParameters) > 0 {
		if err := assignInputField(input, "Parameters", _lightsailParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_lightsailRelationalDatabaseName) > 0 {
		input.RelationalDatabaseName = aws.String(_lightsailRelationalDatabaseName)
	}

	if resp, err := client.UpdateRelationalDatabaseParameters(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_lightsailCmd)
	_lightsailCmd.Flags().SortFlags = false

	_lightsailCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_lightsailCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_lightsailCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_lightsailCmd.Flags().StringVarP(&_lightsailAcceptBundleUpdate, "accept-bundle-update", "", "", "Accept Bundle Update")
	_lightsailCmd.Flags().StringVarP(&_lightsailAccess, "access", "", "", "Access")
	_lightsailCmd.Flags().StringVarP(&_lightsailAccessKeyId, "access-key-id", "", "", "Access Key ID")
	_lightsailCmd.Flags().StringVarP(&_lightsailAccessLogConfig, "access-log-config", "", "", "Access Log Config")
	_lightsailCmd.Flags().StringVarP(&_lightsailAccessRules, "access-rules", "", "", "Access Rules")
	_lightsailCmd.Flags().StringVarP(&_lightsailAddOnRequest, "add-on-request", "", "", "Add On Request")
	_lightsailCmd.Flags().StringVarP(&_lightsailAddOnType, "add-on-type", "", "", "Add On Type")
	_lightsailCmd.Flags().StringVarP(&_lightsailAddOns, "add-ons", "", "", "Add Ons")
	_lightsailCmd.Flags().StringVarP(&_lightsailAlarmName, "alarm-name", "", "", "Alarm Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailAppCategory, "app-category", "", "", "App Category")
	_lightsailCmd.Flags().StringVarP(&_lightsailApplyImmediately, "apply-immediately", "", "", "Apply Immediately")
	_lightsailCmd.Flags().StringVarP(&_lightsailAttachedDiskMapping, "attached-disk-mapping", "", "", "Attached Disk Mapping")
	_lightsailCmd.Flags().StringVarP(&_lightsailAttributeName, "attribute-name", "", "", "Attribute Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailAttributeValue, "attribute-value", "", "", "Attribute Value")
	_lightsailCmd.Flags().StringVarP(&_lightsailAutoMounting, "auto-mounting", "", "", "Auto Mounting")
	_lightsailCmd.Flags().StringVarP(&_lightsailAvailabilityZone, "availability-zone", "", "", "Availability Zone")
	_lightsailCmd.Flags().StringVarP(&_lightsailBlueprintId, "blueprint-id", "", "", "Blueprint ID")
	_lightsailCmd.Flags().StringVarP(&_lightsailBucketName, "bucket-name", "", "", "Bucket Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailBundleId, "bundle-id", "", "", "Bundle ID")
	_lightsailCmd.Flags().StringVarP(&_lightsailCaCertificateIdentifier, "ca-certificate-identifier", "", "", "Ca Certificate Identifier")
	_lightsailCmd.Flags().StringVarP(&_lightsailCacheBehaviorSettings, "cache-behavior-settings", "", "", "Cache Behavior Settings")
	_lightsailCmd.Flags().StringVarP(&_lightsailCacheBehaviors, "cache-behaviors", "", "", "Cache Behaviors")
	_lightsailCmd.Flags().StringSliceVarP(&_lightsailCertificateAlternativeNames, "certificate-alternative-names", "", nil, "Certificate Alternative Names")
	_lightsailCmd.Flags().StringVarP(&_lightsailCertificateDomainName, "certificate-domain-name", "", "", "Certificate Domain Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailCertificateName, "certificate-name", "", "", "Certificate Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailCertificateProvider, "certificate-provider", "", "", "Certificate Provider")
	_lightsailCmd.Flags().StringVarP(&_lightsailCertificateStatuses, "certificate-statuses", "", "", "Certificate Statuses")
	_lightsailCmd.Flags().StringVarP(&_lightsailComparisonOperator, "comparison-operator", "", "", "Comparison Operator")
	_lightsailCmd.Flags().StringVarP(&_lightsailContactEndpoint, "contact-endpoint", "", "", "Contact Endpoint")
	_lightsailCmd.Flags().StringVarP(&_lightsailContactProtocols, "contact-protocols", "", "", "Contact Protocols")
	_lightsailCmd.Flags().StringVarP(&_lightsailContainerName, "container-name", "", "", "Container Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailContainers, "containers", "", "", "Containers")
	_lightsailCmd.Flags().StringVarP(&_lightsailCors, "cors", "", "", "Cors")
	_lightsailCmd.Flags().StringVarP(&_lightsailCustomImageName, "custom-image-name", "", "", "Custom Image Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailDatapointsToAlarm, "datapoints-to-alarm", "", "", "Datapoints To Alarm")
	_lightsailCmd.Flags().StringVarP(&_lightsailDate, "date", "", "", "Date")
	_lightsailCmd.Flags().StringVarP(&_lightsailDefaultCacheBehavior, "default-cache-behavior", "", "", "Default Cache Behavior")
	_lightsailCmd.Flags().StringVarP(&_lightsailDeployment, "deployment", "", "", "Deployment")
	_lightsailCmd.Flags().StringVarP(&_lightsailDigest, "digest", "", "", "Digest")
	_lightsailCmd.Flags().StringVarP(&_lightsailDisableBackupRetention, "disable-backup-retention", "", "", "Disable Backup Retention")
	_lightsailCmd.Flags().StringVarP(&_lightsailDiskName, "disk-name", "", "", "Disk Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailDiskPath, "disk-path", "", "", "Disk Path")
	_lightsailCmd.Flags().StringVarP(&_lightsailDiskSnapshotName, "disk-snapshot-name", "", "", "Disk Snapshot Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailDistributionName, "distribution-name", "", "", "Distribution Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailDomainEntry, "domain-entry", "", "", "Domain Entry")
	_lightsailCmd.Flags().StringVarP(&_lightsailDomainName, "domain-name", "", "", "Domain Name")
	_lightsailCmd.Flags().StringSliceVarP(&_lightsailDomainNames, "domain-names", "", nil, "Domain Names")
	_lightsailCmd.Flags().StringVarP(&_lightsailDurationInMinutes, "duration-in-minutes", "", "", "Duration In Minutes")
	_lightsailCmd.Flags().StringVarP(&_lightsailEmailAddress, "email-address", "", "", "Email Address")
	_lightsailCmd.Flags().StringVarP(&_lightsailEnableBackupRetention, "enable-backup-retention", "", "", "Enable Backup Retention")
	_lightsailCmd.Flags().StringVarP(&_lightsailEnableObjectVersioning, "enable-object-versioning", "", "", "Enable Object Versioning")
	_lightsailCmd.Flags().StringVarP(&_lightsailEndTime, "end-time", "", "", "End Time")
	_lightsailCmd.Flags().StringVarP(&_lightsailEvaluationPeriods, "evaluation-periods", "", "", "Evaluation Periods")
	_lightsailCmd.Flags().StringVarP(&_lightsailExpectedFingerprint, "expected-fingerprint", "", "", "Expected Fingerprint")
	_lightsailCmd.Flags().StringVarP(&_lightsailFilterPattern, "filter-pattern", "", "", "Filter Pattern")
	_lightsailCmd.Flags().StringVarP(&_lightsailFinalRelationalDatabaseSnapshotName, "final-relational-database-snapshot-name", "", "", "Final Relational Database Snapshot Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailForce, "force", "", "", "Force")
	_lightsailCmd.Flags().StringVarP(&_lightsailForceDelete, "force-delete", "", "", "Force Delete")
	_lightsailCmd.Flags().StringVarP(&_lightsailForceDeleteAddOns, "force-delete-add-ons", "", "", "Force Delete Add Ons")
	_lightsailCmd.Flags().StringVarP(&_lightsailHealthCheckPath, "health-check-path", "", "", "Health Check Path")
	_lightsailCmd.Flags().StringVarP(&_lightsailHttpEndpoint, "http-endpoint", "", "", "HTTP Endpoint")
	_lightsailCmd.Flags().StringVarP(&_lightsailHttpProtocolIpv6, "http-protocol-ipv6", "", "", "HTTP Protocol IPV6")
	_lightsailCmd.Flags().StringVarP(&_lightsailHttpPutResponseHopLimit, "http-put-response-hop-limit", "", "", "HTTP Put Response Hop Limit")
	_lightsailCmd.Flags().StringVarP(&_lightsailHttpTokens, "http-tokens", "", "", "HTTP Tokens")
	_lightsailCmd.Flags().StringVarP(&_lightsailImage, "image", "", "", "Image")
	_lightsailCmd.Flags().StringVarP(&_lightsailIncludeAvailabilityZones, "include-availability-zones", "", "", "Include Availability Zones")
	_lightsailCmd.Flags().StringVarP(&_lightsailIncludeCertificateDetails, "include-certificate-details", "", "", "Include Certificate Details")
	_lightsailCmd.Flags().StringVarP(&_lightsailIncludeConnectedResources, "include-connected-resources", "", "", "Include Connected Resources")
	_lightsailCmd.Flags().StringVarP(&_lightsailIncludeCors, "include-cors", "", "", "Include Cors")
	_lightsailCmd.Flags().StringVarP(&_lightsailIncludeDefaultKeyPair, "include-default-key-pair", "", "", "Include Default Key Pair")
	_lightsailCmd.Flags().StringVarP(&_lightsailIncludeInactive, "include-inactive", "", "", "Include Inactive")
	_lightsailCmd.Flags().StringVarP(&_lightsailIncludeRelationalDatabaseAvailabilityZones, "include-relational-database-availability-zones", "", "", "Include Relational Database Availability Zones")
	_lightsailCmd.Flags().StringVarP(&_lightsailInstanceName, "instance-name", "", "", "Instance Name")
	_lightsailCmd.Flags().StringSliceVarP(&_lightsailInstanceNames, "instance-names", "", nil, "Instance Names")
	_lightsailCmd.Flags().StringVarP(&_lightsailInstancePort, "instance-port", "", "", "Instance Port")
	_lightsailCmd.Flags().StringVarP(&_lightsailInstanceSnapshotName, "instance-snapshot-name", "", "", "Instance Snapshot Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailInstances, "instances", "", "", "Instances")
	_lightsailCmd.Flags().StringVarP(&_lightsailIpAddressType, "ip-address-type", "", "", "IP Address Type")
	_lightsailCmd.Flags().StringVarP(&_lightsailIsDisabled, "is-disabled", "", "", "Is Disabled")
	_lightsailCmd.Flags().StringVarP(&_lightsailIsEnabled, "is-enabled", "", "", "Is Enabled")
	_lightsailCmd.Flags().StringVarP(&_lightsailKeyPairName, "key-pair-name", "", "", "Key Pair Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailLabel, "label", "", "", "Label")
	_lightsailCmd.Flags().StringVarP(&_lightsailLoadBalancerName, "load-balancer-name", "", "", "Load Balancer Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailLogStreamName, "log-stream-name", "", "", "Log Stream Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailMasterDatabaseName, "master-database-name", "", "", "Master Database Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailMasterUserPassword, "master-user-password", "", "", "Master User Password")
	_lightsailCmd.Flags().StringVarP(&_lightsailMasterUsername, "master-username", "", "", "Master Username")
	_lightsailCmd.Flags().StringVarP(&_lightsailMetricName, "metric-name", "", "", "Metric Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailMonitoredResourceName, "monitored-resource-name", "", "", "Monitored Resource Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailNotificationEnabled, "notification-enabled", "", "", "Notification Enabled")
	_lightsailCmd.Flags().StringVarP(&_lightsailNotificationTriggers, "notification-triggers", "", "", "Notification Triggers")
	_lightsailCmd.Flags().StringVarP(&_lightsailOperationId, "operation-id", "", "", "Operation ID")
	_lightsailCmd.Flags().StringVarP(&_lightsailOrigin, "origin", "", "", "Origin")
	_lightsailCmd.Flags().StringVarP(&_lightsailPageToken, "page-token", "", "", "Page Token")
	_lightsailCmd.Flags().StringVarP(&_lightsailParameters, "parameters", "", "", "Parameters")
	_lightsailCmd.Flags().StringVarP(&_lightsailPasswordVersion, "password-version", "", "", "Password Version")
	_lightsailCmd.Flags().StringVarP(&_lightsailPeriod, "period", "", "", "Period")
	_lightsailCmd.Flags().StringVarP(&_lightsailPortInfo, "port-info", "", "", "Port Info")
	_lightsailCmd.Flags().StringVarP(&_lightsailPortInfos, "port-infos", "", "", "Port Infos")
	_lightsailCmd.Flags().StringVarP(&_lightsailPower, "power", "", "", "Power")
	_lightsailCmd.Flags().StringVarP(&_lightsailPreferredBackupWindow, "preferred-backup-window", "", "", "Preferred Backup Window")
	_lightsailCmd.Flags().StringVarP(&_lightsailPreferredMaintenanceWindow, "preferred-maintenance-window", "", "", "Preferred Maintenance Window")
	_lightsailCmd.Flags().StringVarP(&_lightsailPrivateRegistryAccess, "private-registry-access", "", "", "Private Registry Access")
	_lightsailCmd.Flags().StringVarP(&_lightsailProtocol, "protocol", "", "", "Protocol")
	_lightsailCmd.Flags().StringVarP(&_lightsailProtocols, "protocols", "", "", "Protocols")
	_lightsailCmd.Flags().StringVarP(&_lightsailPublicDomainNames, "public-domain-names", "", "", "Public Domain Names")
	_lightsailCmd.Flags().StringVarP(&_lightsailPublicEndpoint, "public-endpoint", "", "", "Public Endpoint")
	_lightsailCmd.Flags().StringVarP(&_lightsailPublicKeyBase64, "public-key-base64", "", "", "Public Key Base64")
	_lightsailCmd.Flags().StringVarP(&_lightsailPubliclyAccessible, "publicly-accessible", "", "", "Publicly Accessible")
	_lightsailCmd.Flags().StringSliceVarP(&_lightsailReadonlyAccessAccounts, "readonly-access-accounts", "", nil, "Readonly Access Accounts")
	_lightsailCmd.Flags().StringVarP(&_lightsailRelationalDatabaseBlueprintId, "relational-database-blueprint-id", "", "", "Relational Database Blueprint ID")
	_lightsailCmd.Flags().StringVarP(&_lightsailRelationalDatabaseBundleId, "relational-database-bundle-id", "", "", "Relational Database Bundle ID")
	_lightsailCmd.Flags().StringVarP(&_lightsailRelationalDatabaseName, "relational-database-name", "", "", "Relational Database Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailRelationalDatabaseSnapshotName, "relational-database-snapshot-name", "", "", "Relational Database Snapshot Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailResourceArn, "resource-arn", "", "", "Resource ARN")
	_lightsailCmd.Flags().StringVarP(&_lightsailResourceName, "resource-name", "", "", "Resource Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailResourceType, "resource-type", "", "", "Resource Type")
	_lightsailCmd.Flags().StringVarP(&_lightsailRestoreDate, "restore-date", "", "", "Restore Date")
	_lightsailCmd.Flags().StringVarP(&_lightsailRestoreTime, "restore-time", "", "", "Restore Time")
	_lightsailCmd.Flags().StringVarP(&_lightsailRotateMasterUserPassword, "rotate-master-user-password", "", "", "Rotate Master User Password")
	_lightsailCmd.Flags().StringVarP(&_lightsailScale, "scale", "", "", "Scale")
	_lightsailCmd.Flags().StringVarP(&_lightsailServiceName, "service-name", "", "", "Service Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailSizeInGb, "size-in-gb", "", "", "Size In Gb")
	_lightsailCmd.Flags().StringVarP(&_lightsailSkipFinalSnapshot, "skip-final-snapshot", "", "", "Skip Final Snapshot")
	_lightsailCmd.Flags().StringVarP(&_lightsailSourceDiskName, "source-disk-name", "", "", "Source Disk Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailSourceInstanceName, "source-instance-name", "", "", "Source Instance Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailSourceRegion, "source-region", "", "", "Source Region")
	_lightsailCmd.Flags().StringVarP(&_lightsailSourceRelationalDatabaseName, "source-relational-database-name", "", "", "Source Relational Database Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailSourceResourceName, "source-resource-name", "", "", "Source Resource Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailSourceSnapshotName, "source-snapshot-name", "", "", "Source Snapshot Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailStartFromHead, "start-from-head", "", "", "Start From Head")
	_lightsailCmd.Flags().StringVarP(&_lightsailStartTime, "start-time", "", "", "Start Time")
	_lightsailCmd.Flags().StringVarP(&_lightsailState, "state", "", "", "State")
	_lightsailCmd.Flags().StringVarP(&_lightsailStaticIpName, "static-ip-name", "", "", "Static IP Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailStatistics, "statistics", "", "", "Statistics")
	_lightsailCmd.Flags().StringSliceVarP(&_lightsailSubjectAlternativeNames, "subject-alternative-names", "", nil, "Subject Alternative Names")
	_lightsailCmd.Flags().StringSliceVarP(&_lightsailTagKeys, "tag-keys", "", nil, "Tag Keys")
	_lightsailCmd.Flags().StringVarP(&_lightsailTags, "tags", "", "", "Tags")
	_lightsailCmd.Flags().StringVarP(&_lightsailTargetSnapshotName, "target-snapshot-name", "", "", "Target Snapshot Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailThreshold, "threshold", "", "", "Threshold")
	_lightsailCmd.Flags().StringVarP(&_lightsailTlsPolicyName, "tls-policy-name", "", "", "TLS Policy Name")
	_lightsailCmd.Flags().StringVarP(&_lightsailTreatMissingData, "treat-missing-data", "", "", "Treat Missing Data")
	_lightsailCmd.Flags().StringVarP(&_lightsailUnit, "unit", "", "", "Unit")
	_lightsailCmd.Flags().StringVarP(&_lightsailUseDefaultCertificate, "use-default-certificate", "", "", "Use Default Certificate")
	_lightsailCmd.Flags().StringVarP(&_lightsailUseLatestRestorableAutoSnapshot, "use-latest-restorable-auto-snapshot", "", "", "Use Latest Restorable Auto Snapshot")
	_lightsailCmd.Flags().StringVarP(&_lightsailUseLatestRestorableTime, "use-latest-restorable-time", "", "", "Use Latest Restorable Time")
	_lightsailCmd.Flags().StringVarP(&_lightsailUserData, "user-data", "", "", "User Data")
	_lightsailCmd.Flags().StringVarP(&_lightsailVersioning, "versioning", "", "", "Versioning")
	_lightsailCmd.Flags().StringVarP(&_lightsailViewerMinimumTlsProtocolVersion, "viewer-minimum-tls-protocol-version", "", "", "Viewer Minimum TLS Protocol Version")

	_lightsailCmd.Flags().BoolVarP(&_lightsailAllocateStaticIp, "allocate-static-ip", "", false, "Allocate Static IP")
	_lightsailCmd.Flags().BoolVarP(&_lightsailAttachCertificateToDistribution, "attach-certificate-to-distribution", "", false, "Attach Certificate To Distribution")
	_lightsailCmd.Flags().BoolVarP(&_lightsailAttachDisk, "attach-disk", "", false, "Attach Disk")
	_lightsailCmd.Flags().BoolVarP(&_lightsailAttachInstancesToLoadBalancer, "attach-instances-to-load-balancer", "", false, "Attach Instances To Load Balancer")
	_lightsailCmd.Flags().BoolVarP(&_lightsailAttachLoadBalancerTlsCertificate, "attach-load-balancer-tls-certificate", "", false, "Attach Load Balancer TLS Certificate")
	_lightsailCmd.Flags().BoolVarP(&_lightsailAttachStaticIp, "attach-static-ip", "", false, "Attach Static IP")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCloseInstancePublicPorts, "close-instance-public-ports", "", false, "Close Instance Public Ports")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCopySnapshot, "copy-snapshot", "", false, "Copy Snapshot")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCreateBucket, "create-bucket", "", false, "Create Bucket")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCreateBucketAccessKey, "create-bucket-access-key", "", false, "Create Bucket Access Key")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCreateCertificate, "create-certificate", "", false, "Create Certificate")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCreateCloudFormationStack, "create-cloud-formation-stack", "", false, "Create Cloud Formation Stack")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCreateContactMethod, "create-contact-method", "", false, "Create Contact Method")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCreateContainerService, "create-container-service", "", false, "Create Container Service")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCreateContainerServiceDeployment, "create-container-service-deployment", "", false, "Create Container Service Deployment")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCreateContainerServiceRegistryLogin, "create-container-service-registry-login", "", false, "Create Container Service Registry Login")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCreateDisk, "create-disk", "", false, "Create Disk")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCreateDiskFromSnapshot, "create-disk-from-snapshot", "", false, "Create Disk From Snapshot")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCreateDiskSnapshot, "create-disk-snapshot", "", false, "Create Disk Snapshot")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCreateDistribution, "create-distribution", "", false, "Create Distribution")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCreateDomain, "create-domain", "", false, "Create Domain")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCreateDomainEntry, "create-domain-entry", "", false, "Create Domain Entry")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCreateGUISessionAccessDetails, "create-gui-session-access-details", "", false, "Create Gui Session Access Details")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCreateInstanceSnapshot, "create-instance-snapshot", "", false, "Create Instance Snapshot")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCreateInstances, "create-instances", "", false, "Create Instances")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCreateInstancesFromSnapshot, "create-instances-from-snapshot", "", false, "Create Instances From Snapshot")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCreateKeyPair, "create-key-pair", "", false, "Create Key Pair")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCreateLoadBalancer, "create-load-balancer", "", false, "Create Load Balancer")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCreateLoadBalancerTlsCertificate, "create-load-balancer-tls-certificate", "", false, "Create Load Balancer TLS Certificate")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCreateRelationalDatabase, "create-relational-database", "", false, "Create Relational Database")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCreateRelationalDatabaseFromSnapshot, "create-relational-database-from-snapshot", "", false, "Create Relational Database From Snapshot")
	_lightsailCmd.Flags().BoolVarP(&_lightsailCreateRelationalDatabaseSnapshot, "create-relational-database-snapshot", "", false, "Create Relational Database Snapshot")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDeleteAlarm, "delete-alarm", "", false, "Delete Alarm")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDeleteAutoSnapshot, "delete-auto-snapshot", "", false, "Delete Auto Snapshot")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDeleteBucket, "delete-bucket", "", false, "Delete Bucket")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDeleteBucketAccessKey, "delete-bucket-access-key", "", false, "Delete Bucket Access Key")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDeleteCertificate, "delete-certificate", "", false, "Delete Certificate")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDeleteContactMethod, "delete-contact-method", "", false, "Delete Contact Method")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDeleteContainerImage, "delete-container-image", "", false, "Delete Container Image")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDeleteContainerService, "delete-container-service", "", false, "Delete Container Service")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDeleteDisk, "delete-disk", "", false, "Delete Disk")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDeleteDiskSnapshot, "delete-disk-snapshot", "", false, "Delete Disk Snapshot")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDeleteDistribution, "delete-distribution", "", false, "Delete Distribution")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDeleteDomain, "delete-domain", "", false, "Delete Domain")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDeleteDomainEntry, "delete-domain-entry", "", false, "Delete Domain Entry")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDeleteInstance, "delete-instance", "", false, "Delete Instance")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDeleteInstanceSnapshot, "delete-instance-snapshot", "", false, "Delete Instance Snapshot")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDeleteKeyPair, "delete-key-pair", "", false, "Delete Key Pair")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDeleteKnownHostKeys, "delete-known-host-keys", "", false, "Delete Known Host Keys")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDeleteLoadBalancer, "delete-load-balancer", "", false, "Delete Load Balancer")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDeleteLoadBalancerTlsCertificate, "delete-load-balancer-tls-certificate", "", false, "Delete Load Balancer TLS Certificate")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDeleteRelationalDatabase, "delete-relational-database", "", false, "Delete Relational Database")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDeleteRelationalDatabaseSnapshot, "delete-relational-database-snapshot", "", false, "Delete Relational Database Snapshot")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDetachCertificateFromDistribution, "detach-certificate-from-distribution", "", false, "Detach Certificate From Distribution")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDetachDisk, "detach-disk", "", false, "Detach Disk")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDetachInstancesFromLoadBalancer, "detach-instances-from-load-balancer", "", false, "Detach Instances From Load Balancer")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDetachStaticIp, "detach-static-ip", "", false, "Detach Static IP")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDisableAddOn, "disable-add-on", "", false, "Disable Add On")
	_lightsailCmd.Flags().BoolVarP(&_lightsailDownloadDefaultKeyPair, "download-default-key-pair", "", false, "Download Default Key Pair")
	_lightsailCmd.Flags().BoolVarP(&_lightsailEnableAddOn, "enable-add-on", "", false, "Enable Add On")
	_lightsailCmd.Flags().BoolVarP(&_lightsailExportSnapshot, "export-snapshot", "", false, "Export Snapshot")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetActiveNames, "get-active-names", "", false, "Get Active Names")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetAlarms, "get-alarms", "", false, "Get Alarms")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetAutoSnapshots, "get-auto-snapshots", "", false, "Get Auto Snapshots")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetBlueprints, "get-blueprints", "", false, "Get Blueprints")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetBucketAccessKeys, "get-bucket-access-keys", "", false, "Get Bucket Access Keys")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetBucketBundles, "get-bucket-bundles", "", false, "Get Bucket Bundles")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetBucketMetricData, "get-bucket-metric-data", "", false, "Get Bucket Metric Data")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetBuckets, "get-buckets", "", false, "Get Buckets")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetBundles, "get-bundles", "", false, "Get Bundles")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetCertificates, "get-certificates", "", false, "Get Certificates")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetCloudFormationStackRecords, "get-cloud-formation-stack-records", "", false, "Get Cloud Formation Stack Records")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetContactMethods, "get-contact-methods", "", false, "Get Contact Methods")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetContainerAPIMetadata, "get-container-api-metadata", "", false, "Get Container API Metadata")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetContainerImages, "get-container-images", "", false, "Get Container Images")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetContainerLog, "get-container-log", "", false, "Get Container Log")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetContainerServiceDeployments, "get-container-service-deployments", "", false, "Get Container Service Deployments")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetContainerServiceMetricData, "get-container-service-metric-data", "", false, "Get Container Service Metric Data")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetContainerServicePowers, "get-container-service-powers", "", false, "Get Container Service Powers")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetContainerServices, "get-container-services", "", false, "Get Container Services")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetCostEstimate, "get-cost-estimate", "", false, "Get Cost Estimate")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetDisk, "get-disk", "", false, "Get Disk")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetDiskSnapshot, "get-disk-snapshot", "", false, "Get Disk Snapshot")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetDiskSnapshots, "get-disk-snapshots", "", false, "Get Disk Snapshots")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetDisks, "get-disks", "", false, "Get Disks")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetDistributionBundles, "get-distribution-bundles", "", false, "Get Distribution Bundles")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetDistributionLatestCacheReset, "get-distribution-latest-cache-reset", "", false, "Get Distribution Latest Cache Reset")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetDistributionMetricData, "get-distribution-metric-data", "", false, "Get Distribution Metric Data")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetDistributions, "get-distributions", "", false, "Get Distributions")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetDomain, "get-domain", "", false, "Get Domain")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetDomains, "get-domains", "", false, "Get Domains")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetExportSnapshotRecords, "get-export-snapshot-records", "", false, "Get Export Snapshot Records")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetInstance, "get-instance", "", false, "Get Instance")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetInstanceAccessDetails, "get-instance-access-details", "", false, "Get Instance Access Details")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetInstanceMetricData, "get-instance-metric-data", "", false, "Get Instance Metric Data")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetInstancePortStates, "get-instance-port-states", "", false, "Get Instance Port States")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetInstanceSnapshot, "get-instance-snapshot", "", false, "Get Instance Snapshot")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetInstanceSnapshots, "get-instance-snapshots", "", false, "Get Instance Snapshots")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetInstanceState, "get-instance-state", "", false, "Get Instance State")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetInstances, "get-instances", "", false, "Get Instances")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetKeyPair, "get-key-pair", "", false, "Get Key Pair")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetKeyPairs, "get-key-pairs", "", false, "Get Key Pairs")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetLoadBalancer, "get-load-balancer", "", false, "Get Load Balancer")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetLoadBalancerMetricData, "get-load-balancer-metric-data", "", false, "Get Load Balancer Metric Data")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetLoadBalancerTlsCertificates, "get-load-balancer-tls-certificates", "", false, "Get Load Balancer TLS Certificates")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetLoadBalancerTlsPolicies, "get-load-balancer-tls-policies", "", false, "Get Load Balancer TLS Policies")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetLoadBalancers, "get-load-balancers", "", false, "Get Load Balancers")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetOperation, "get-operation", "", false, "Get Operation")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetOperations, "get-operations", "", false, "Get Operations")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetOperationsForResource, "get-operations-for-resource", "", false, "Get Operations For Resource")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetRegions, "get-regions", "", false, "Get Regions")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetRelationalDatabase, "get-relational-database", "", false, "Get Relational Database")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetRelationalDatabaseBlueprints, "get-relational-database-blueprints", "", false, "Get Relational Database Blueprints")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetRelationalDatabaseBundles, "get-relational-database-bundles", "", false, "Get Relational Database Bundles")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetRelationalDatabaseEvents, "get-relational-database-events", "", false, "Get Relational Database Events")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetRelationalDatabaseLogEvents, "get-relational-database-log-events", "", false, "Get Relational Database Log Events")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetRelationalDatabaseLogStreams, "get-relational-database-log-streams", "", false, "Get Relational Database Log Streams")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetRelationalDatabaseMasterUserPassword, "get-relational-database-master-user-password", "", false, "Get Relational Database Master User Password")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetRelationalDatabaseMetricData, "get-relational-database-metric-data", "", false, "Get Relational Database Metric Data")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetRelationalDatabaseParameters, "get-relational-database-parameters", "", false, "Get Relational Database Parameters")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetRelationalDatabaseSnapshot, "get-relational-database-snapshot", "", false, "Get Relational Database Snapshot")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetRelationalDatabaseSnapshots, "get-relational-database-snapshots", "", false, "Get Relational Database Snapshots")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetRelationalDatabases, "get-relational-databases", "", false, "Get Relational Databases")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetSetupHistory, "get-setup-history", "", false, "Get Setup History")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetStaticIp, "get-static-ip", "", false, "Get Static IP")
	_lightsailCmd.Flags().BoolVarP(&_lightsailGetStaticIps, "get-static-ips", "", false, "Get Static Ips")
	_lightsailCmd.Flags().BoolVarP(&_lightsailImportKeyPair, "import-key-pair", "", false, "Import Key Pair")
	_lightsailCmd.Flags().BoolVarP(&_lightsailIsVpcPeered, "is-vpc-peered", "", false, "Is VPC Peered")
	_lightsailCmd.Flags().BoolVarP(&_lightsailOpenInstancePublicPorts, "open-instance-public-ports", "", false, "Open Instance Public Ports")
	_lightsailCmd.Flags().BoolVarP(&_lightsailPeerVpc, "peer-vpc", "", false, "Peer VPC")
	_lightsailCmd.Flags().BoolVarP(&_lightsailPutAlarm, "put-alarm", "", false, "Put Alarm")
	_lightsailCmd.Flags().BoolVarP(&_lightsailPutInstancePublicPorts, "put-instance-public-ports", "", false, "Put Instance Public Ports")
	_lightsailCmd.Flags().BoolVarP(&_lightsailRebootInstance, "reboot-instance", "", false, "Reboot Instance")
	_lightsailCmd.Flags().BoolVarP(&_lightsailRebootRelationalDatabase, "reboot-relational-database", "", false, "Reboot Relational Database")
	_lightsailCmd.Flags().BoolVarP(&_lightsailRegisterContainerImage, "register-container-image", "", false, "Register Container Image")
	_lightsailCmd.Flags().BoolVarP(&_lightsailReleaseStaticIp, "release-static-ip", "", false, "Release Static IP")
	_lightsailCmd.Flags().BoolVarP(&_lightsailResetDistributionCache, "reset-distribution-cache", "", false, "Reset Distribution Cache")
	_lightsailCmd.Flags().BoolVarP(&_lightsailSendContactMethodVerification, "send-contact-method-verification", "", false, "Send Contact Method Verification")
	_lightsailCmd.Flags().BoolVarP(&_lightsailSetIpAddressType, "set-ip-address-type", "", false, "Set IP Address Type")
	_lightsailCmd.Flags().BoolVarP(&_lightsailSetResourceAccessForBucket, "set-resource-access-for-bucket", "", false, "Set Resource Access For Bucket")
	_lightsailCmd.Flags().BoolVarP(&_lightsailSetupInstanceHttps, "setup-instance-https", "", false, "Setup Instance HTTPS")
	_lightsailCmd.Flags().BoolVarP(&_lightsailStartGUISession, "start-gui-session", "", false, "Start Gui Session")
	_lightsailCmd.Flags().BoolVarP(&_lightsailStartInstance, "start-instance", "", false, "Start Instance")
	_lightsailCmd.Flags().BoolVarP(&_lightsailStartRelationalDatabase, "start-relational-database", "", false, "Start Relational Database")
	_lightsailCmd.Flags().BoolVarP(&_lightsailStopGUISession, "stop-gui-session", "", false, "Stop Gui Session")
	_lightsailCmd.Flags().BoolVarP(&_lightsailStopInstance, "stop-instance", "", false, "Stop Instance")
	_lightsailCmd.Flags().BoolVarP(&_lightsailStopRelationalDatabase, "stop-relational-database", "", false, "Stop Relational Database")
	_lightsailCmd.Flags().BoolVarP(&_lightsailTagResource, "tag-resource", "", false, "Tag Resource")
	_lightsailCmd.Flags().BoolVarP(&_lightsailTestAlarm, "test-alarm", "", false, "Test Alarm")
	_lightsailCmd.Flags().BoolVarP(&_lightsailUnpeerVpc, "unpeer-vpc", "", false, "Unpeer VPC")
	_lightsailCmd.Flags().BoolVarP(&_lightsailUntagResource, "untag-resource", "", false, "Untag Resource")
	_lightsailCmd.Flags().BoolVarP(&_lightsailUpdateBucket, "update-bucket", "", false, "Update Bucket")
	_lightsailCmd.Flags().BoolVarP(&_lightsailUpdateBucketBundle, "update-bucket-bundle", "", false, "Update Bucket Bundle")
	_lightsailCmd.Flags().BoolVarP(&_lightsailUpdateContainerService, "update-container-service", "", false, "Update Container Service")
	_lightsailCmd.Flags().BoolVarP(&_lightsailUpdateDistribution, "update-distribution", "", false, "Update Distribution")
	_lightsailCmd.Flags().BoolVarP(&_lightsailUpdateDistributionBundle, "update-distribution-bundle", "", false, "Update Distribution Bundle")
	_lightsailCmd.Flags().BoolVarP(&_lightsailUpdateDomainEntry, "update-domain-entry", "", false, "Update Domain Entry")
	_lightsailCmd.Flags().BoolVarP(&_lightsailUpdateInstanceMetadataOptions, "update-instance-metadata-options", "", false, "Update Instance Metadata Options")
	_lightsailCmd.Flags().BoolVarP(&_lightsailUpdateLoadBalancerAttribute, "update-load-balancer-attribute", "", false, "Update Load Balancer Attribute")
	_lightsailCmd.Flags().BoolVarP(&_lightsailUpdateRelationalDatabase, "update-relational-database", "", false, "Update Relational Database")
	_lightsailCmd.Flags().BoolVarP(&_lightsailUpdateRelationalDatabaseParameters, "update-relational-database-parameters", "", false, "Update Relational Database Parameters")

}
