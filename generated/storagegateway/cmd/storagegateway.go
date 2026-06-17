package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// storagegatewayCmd represents the storagegateway command
var _storagegatewayCmd = &cobra.Command{
	Use:   "storagegateway",
	Short: "AWS storagegateway CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := storagegateway.NewFromConfig(cfg)
		if _storagegatewayActivateGateway {
			storagegateway_ActivateGateway(cfg, client)
			return
		}
		if _storagegatewayAddCache {
			storagegateway_AddCache(cfg, client)
			return
		}
		if _storagegatewayAddTagsToResource {
			storagegateway_AddTagsToResource(cfg, client)
			return
		}
		if _storagegatewayAddUploadBuffer {
			storagegateway_AddUploadBuffer(cfg, client)
			return
		}
		if _storagegatewayAddWorkingStorage {
			storagegateway_AddWorkingStorage(cfg, client)
			return
		}
		if _storagegatewayAssignTapePool {
			storagegateway_AssignTapePool(cfg, client)
			return
		}
		if _storagegatewayAssociateFileSystem {
			storagegateway_AssociateFileSystem(cfg, client)
			return
		}
		if _storagegatewayAttachVolume {
			storagegateway_AttachVolume(cfg, client)
			return
		}
		if _storagegatewayCancelArchival {
			storagegateway_CancelArchival(cfg, client)
			return
		}
		if _storagegatewayCancelCacheReport {
			storagegateway_CancelCacheReport(cfg, client)
			return
		}
		if _storagegatewayCancelRetrieval {
			storagegateway_CancelRetrieval(cfg, client)
			return
		}
		if _storagegatewayCreateCachediSCSIVolume {
			storagegateway_CreateCachediSCSIVolume(cfg, client)
			return
		}
		if _storagegatewayCreateNFSFileShare {
			storagegateway_CreateNFSFileShare(cfg, client)
			return
		}
		if _storagegatewayCreateSMBFileShare {
			storagegateway_CreateSMBFileShare(cfg, client)
			return
		}
		if _storagegatewayCreateSnapshot {
			storagegateway_CreateSnapshot(cfg, client)
			return
		}
		if _storagegatewayCreateSnapshotFromVolumeRecoveryPoint {
			storagegateway_CreateSnapshotFromVolumeRecoveryPoint(cfg, client)
			return
		}
		if _storagegatewayCreateStorediSCSIVolume {
			storagegateway_CreateStorediSCSIVolume(cfg, client)
			return
		}
		if _storagegatewayCreateTapePool {
			storagegateway_CreateTapePool(cfg, client)
			return
		}
		if _storagegatewayCreateTapeWithBarcode {
			storagegateway_CreateTapeWithBarcode(cfg, client)
			return
		}
		if _storagegatewayCreateTapes {
			storagegateway_CreateTapes(cfg, client)
			return
		}
		if _storagegatewayDeleteAutomaticTapeCreationPolicy {
			storagegateway_DeleteAutomaticTapeCreationPolicy(cfg, client)
			return
		}
		if _storagegatewayDeleteBandwidthRateLimit {
			storagegateway_DeleteBandwidthRateLimit(cfg, client)
			return
		}
		if _storagegatewayDeleteCacheReport {
			storagegateway_DeleteCacheReport(cfg, client)
			return
		}
		if _storagegatewayDeleteChapCredentials {
			storagegateway_DeleteChapCredentials(cfg, client)
			return
		}
		if _storagegatewayDeleteFileShare {
			storagegateway_DeleteFileShare(cfg, client)
			return
		}
		if _storagegatewayDeleteGateway {
			storagegateway_DeleteGateway(cfg, client)
			return
		}
		if _storagegatewayDeleteSnapshotSchedule {
			storagegateway_DeleteSnapshotSchedule(cfg, client)
			return
		}
		if _storagegatewayDeleteTape {
			storagegateway_DeleteTape(cfg, client)
			return
		}
		if _storagegatewayDeleteTapeArchive {
			storagegateway_DeleteTapeArchive(cfg, client)
			return
		}
		if _storagegatewayDeleteTapePool {
			storagegateway_DeleteTapePool(cfg, client)
			return
		}
		if _storagegatewayDeleteVolume {
			storagegateway_DeleteVolume(cfg, client)
			return
		}
		if _storagegatewayDescribeAvailabilityMonitorTest {
			storagegateway_DescribeAvailabilityMonitorTest(cfg, client)
			return
		}
		if _storagegatewayDescribeBandwidthRateLimit {
			storagegateway_DescribeBandwidthRateLimit(cfg, client)
			return
		}
		if _storagegatewayDescribeBandwidthRateLimitSchedule {
			storagegateway_DescribeBandwidthRateLimitSchedule(cfg, client)
			return
		}
		if _storagegatewayDescribeCache {
			storagegateway_DescribeCache(cfg, client)
			return
		}
		if _storagegatewayDescribeCacheReport {
			storagegateway_DescribeCacheReport(cfg, client)
			return
		}
		if _storagegatewayDescribeCachediSCSIVolumes {
			storagegateway_DescribeCachediSCSIVolumes(cfg, client)
			return
		}
		if _storagegatewayDescribeChapCredentials {
			storagegateway_DescribeChapCredentials(cfg, client)
			return
		}
		if _storagegatewayDescribeFileSystemAssociations {
			storagegateway_DescribeFileSystemAssociations(cfg, client)
			return
		}
		if _storagegatewayDescribeGatewayInformation {
			storagegateway_DescribeGatewayInformation(cfg, client)
			return
		}
		if _storagegatewayDescribeMaintenanceStartTime {
			storagegateway_DescribeMaintenanceStartTime(cfg, client)
			return
		}
		if _storagegatewayDescribeNFSFileShares {
			storagegateway_DescribeNFSFileShares(cfg, client)
			return
		}
		if _storagegatewayDescribeSMBFileShares {
			storagegateway_DescribeSMBFileShares(cfg, client)
			return
		}
		if _storagegatewayDescribeSMBSettings {
			storagegateway_DescribeSMBSettings(cfg, client)
			return
		}
		if _storagegatewayDescribeSnapshotSchedule {
			storagegateway_DescribeSnapshotSchedule(cfg, client)
			return
		}
		if _storagegatewayDescribeStorediSCSIVolumes {
			storagegateway_DescribeStorediSCSIVolumes(cfg, client)
			return
		}
		if _storagegatewayDescribeTapeArchives {
			storagegateway_DescribeTapeArchives(cfg, client)
			return
		}
		if _storagegatewayDescribeTapeRecoveryPoints {
			storagegateway_DescribeTapeRecoveryPoints(cfg, client)
			return
		}
		if _storagegatewayDescribeTapes {
			storagegateway_DescribeTapes(cfg, client)
			return
		}
		if _storagegatewayDescribeUploadBuffer {
			storagegateway_DescribeUploadBuffer(cfg, client)
			return
		}
		if _storagegatewayDescribeVTLDevices {
			storagegateway_DescribeVTLDevices(cfg, client)
			return
		}
		if _storagegatewayDescribeWorkingStorage {
			storagegateway_DescribeWorkingStorage(cfg, client)
			return
		}
		if _storagegatewayDetachVolume {
			storagegateway_DetachVolume(cfg, client)
			return
		}
		if _storagegatewayDisableGateway {
			storagegateway_DisableGateway(cfg, client)
			return
		}
		if _storagegatewayDisassociateFileSystem {
			storagegateway_DisassociateFileSystem(cfg, client)
			return
		}
		if _storagegatewayEvictFilesFailingUpload {
			storagegateway_EvictFilesFailingUpload(cfg, client)
			return
		}
		if _storagegatewayJoinDomain {
			storagegateway_JoinDomain(cfg, client)
			return
		}
		if _storagegatewayListAutomaticTapeCreationPolicies {
			storagegateway_ListAutomaticTapeCreationPolicies(cfg, client)
			return
		}
		if _storagegatewayListCacheReports {
			storagegateway_ListCacheReports(cfg, client)
			return
		}
		if _storagegatewayListFileShares {
			storagegateway_ListFileShares(cfg, client)
			return
		}
		if _storagegatewayListFileSystemAssociations {
			storagegateway_ListFileSystemAssociations(cfg, client)
			return
		}
		if _storagegatewayListGateways {
			storagegateway_ListGateways(cfg, client)
			return
		}
		if _storagegatewayListLocalDisks {
			storagegateway_ListLocalDisks(cfg, client)
			return
		}
		if _storagegatewayListTagsForResource {
			storagegateway_ListTagsForResource(cfg, client)
			return
		}
		if _storagegatewayListTapePools {
			storagegateway_ListTapePools(cfg, client)
			return
		}
		if _storagegatewayListTapes {
			storagegateway_ListTapes(cfg, client)
			return
		}
		if _storagegatewayListVolumeInitiators {
			storagegateway_ListVolumeInitiators(cfg, client)
			return
		}
		if _storagegatewayListVolumeRecoveryPoints {
			storagegateway_ListVolumeRecoveryPoints(cfg, client)
			return
		}
		if _storagegatewayListVolumes {
			storagegateway_ListVolumes(cfg, client)
			return
		}
		if _storagegatewayNotifyWhenUploaded {
			storagegateway_NotifyWhenUploaded(cfg, client)
			return
		}
		if _storagegatewayRefreshCache {
			storagegateway_RefreshCache(cfg, client)
			return
		}
		if _storagegatewayRemoveTagsFromResource {
			storagegateway_RemoveTagsFromResource(cfg, client)
			return
		}
		if _storagegatewayResetCache {
			storagegateway_ResetCache(cfg, client)
			return
		}
		if _storagegatewayRetrieveTapeArchive {
			storagegateway_RetrieveTapeArchive(cfg, client)
			return
		}
		if _storagegatewayRetrieveTapeRecoveryPoint {
			storagegateway_RetrieveTapeRecoveryPoint(cfg, client)
			return
		}
		if _storagegatewaySetLocalConsolePassword {
			storagegateway_SetLocalConsolePassword(cfg, client)
			return
		}
		if _storagegatewaySetSMBGuestPassword {
			storagegateway_SetSMBGuestPassword(cfg, client)
			return
		}
		if _storagegatewayShutdownGateway {
			storagegateway_ShutdownGateway(cfg, client)
			return
		}
		if _storagegatewayStartAvailabilityMonitorTest {
			storagegateway_StartAvailabilityMonitorTest(cfg, client)
			return
		}
		if _storagegatewayStartCacheReport {
			storagegateway_StartCacheReport(cfg, client)
			return
		}
		if _storagegatewayStartGateway {
			storagegateway_StartGateway(cfg, client)
			return
		}
		if _storagegatewayUpdateAutomaticTapeCreationPolicy {
			storagegateway_UpdateAutomaticTapeCreationPolicy(cfg, client)
			return
		}
		if _storagegatewayUpdateBandwidthRateLimit {
			storagegateway_UpdateBandwidthRateLimit(cfg, client)
			return
		}
		if _storagegatewayUpdateBandwidthRateLimitSchedule {
			storagegateway_UpdateBandwidthRateLimitSchedule(cfg, client)
			return
		}
		if _storagegatewayUpdateChapCredentials {
			storagegateway_UpdateChapCredentials(cfg, client)
			return
		}
		if _storagegatewayUpdateFileSystemAssociation {
			storagegateway_UpdateFileSystemAssociation(cfg, client)
			return
		}
		if _storagegatewayUpdateGatewayInformation {
			storagegateway_UpdateGatewayInformation(cfg, client)
			return
		}
		if _storagegatewayUpdateGatewaySoftwareNow {
			storagegateway_UpdateGatewaySoftwareNow(cfg, client)
			return
		}
		if _storagegatewayUpdateMaintenanceStartTime {
			storagegateway_UpdateMaintenanceStartTime(cfg, client)
			return
		}
		if _storagegatewayUpdateNFSFileShare {
			storagegateway_UpdateNFSFileShare(cfg, client)
			return
		}
		if _storagegatewayUpdateSMBFileShare {
			storagegateway_UpdateSMBFileShare(cfg, client)
			return
		}
		if _storagegatewayUpdateSMBFileShareVisibility {
			storagegateway_UpdateSMBFileShareVisibility(cfg, client)
			return
		}
		if _storagegatewayUpdateSMBLocalGroups {
			storagegateway_UpdateSMBLocalGroups(cfg, client)
			return
		}
		if _storagegatewayUpdateSMBSecurityStrategy {
			storagegateway_UpdateSMBSecurityStrategy(cfg, client)
			return
		}
		if _storagegatewayUpdateSnapshotSchedule {
			storagegateway_UpdateSnapshotSchedule(cfg, client)
			return
		}
		if _storagegatewayUpdateVTLDeviceType {
			storagegateway_UpdateVTLDeviceType(cfg, client)
			return
		}

	},
}

var (
	_storagegatewayActivateGateway                       bool
	_storagegatewayAddCache                              bool
	_storagegatewayAddTagsToResource                     bool
	_storagegatewayAddUploadBuffer                       bool
	_storagegatewayAddWorkingStorage                     bool
	_storagegatewayAssignTapePool                        bool
	_storagegatewayAssociateFileSystem                   bool
	_storagegatewayAttachVolume                          bool
	_storagegatewayCancelArchival                        bool
	_storagegatewayCancelCacheReport                     bool
	_storagegatewayCancelRetrieval                       bool
	_storagegatewayCreateCachediSCSIVolume               bool
	_storagegatewayCreateNFSFileShare                    bool
	_storagegatewayCreateSMBFileShare                    bool
	_storagegatewayCreateSnapshot                        bool
	_storagegatewayCreateSnapshotFromVolumeRecoveryPoint bool
	_storagegatewayCreateStorediSCSIVolume               bool
	_storagegatewayCreateTapePool                        bool
	_storagegatewayCreateTapeWithBarcode                 bool
	_storagegatewayCreateTapes                           bool
	_storagegatewayDeleteAutomaticTapeCreationPolicy     bool
	_storagegatewayDeleteBandwidthRateLimit              bool
	_storagegatewayDeleteCacheReport                     bool
	_storagegatewayDeleteChapCredentials                 bool
	_storagegatewayDeleteFileShare                       bool
	_storagegatewayDeleteGateway                         bool
	_storagegatewayDeleteSnapshotSchedule                bool
	_storagegatewayDeleteTape                            bool
	_storagegatewayDeleteTapeArchive                     bool
	_storagegatewayDeleteTapePool                        bool
	_storagegatewayDeleteVolume                          bool
	_storagegatewayDescribeAvailabilityMonitorTest       bool
	_storagegatewayDescribeBandwidthRateLimit            bool
	_storagegatewayDescribeBandwidthRateLimitSchedule    bool
	_storagegatewayDescribeCache                         bool
	_storagegatewayDescribeCacheReport                   bool
	_storagegatewayDescribeCachediSCSIVolumes            bool
	_storagegatewayDescribeChapCredentials               bool
	_storagegatewayDescribeFileSystemAssociations        bool
	_storagegatewayDescribeGatewayInformation            bool
	_storagegatewayDescribeMaintenanceStartTime          bool
	_storagegatewayDescribeNFSFileShares                 bool
	_storagegatewayDescribeSMBFileShares                 bool
	_storagegatewayDescribeSMBSettings                   bool
	_storagegatewayDescribeSnapshotSchedule              bool
	_storagegatewayDescribeStorediSCSIVolumes            bool
	_storagegatewayDescribeTapeArchives                  bool
	_storagegatewayDescribeTapeRecoveryPoints            bool
	_storagegatewayDescribeTapes                         bool
	_storagegatewayDescribeUploadBuffer                  bool
	_storagegatewayDescribeVTLDevices                    bool
	_storagegatewayDescribeWorkingStorage                bool
	_storagegatewayDetachVolume                          bool
	_storagegatewayDisableGateway                        bool
	_storagegatewayDisassociateFileSystem                bool
	_storagegatewayEvictFilesFailingUpload               bool
	_storagegatewayJoinDomain                            bool
	_storagegatewayListAutomaticTapeCreationPolicies     bool
	_storagegatewayListCacheReports                      bool
	_storagegatewayListFileShares                        bool
	_storagegatewayListFileSystemAssociations            bool
	_storagegatewayListGateways                          bool
	_storagegatewayListLocalDisks                        bool
	_storagegatewayListTagsForResource                   bool
	_storagegatewayListTapePools                         bool
	_storagegatewayListTapes                             bool
	_storagegatewayListVolumeInitiators                  bool
	_storagegatewayListVolumeRecoveryPoints              bool
	_storagegatewayListVolumes                           bool
	_storagegatewayNotifyWhenUploaded                    bool
	_storagegatewayRefreshCache                          bool
	_storagegatewayRemoveTagsFromResource                bool
	_storagegatewayResetCache                            bool
	_storagegatewayRetrieveTapeArchive                   bool
	_storagegatewayRetrieveTapeRecoveryPoint             bool
	_storagegatewaySetLocalConsolePassword               bool
	_storagegatewaySetSMBGuestPassword                   bool
	_storagegatewayShutdownGateway                       bool
	_storagegatewayStartAvailabilityMonitorTest          bool
	_storagegatewayStartCacheReport                      bool
	_storagegatewayStartGateway                          bool
	_storagegatewayUpdateAutomaticTapeCreationPolicy     bool
	_storagegatewayUpdateBandwidthRateLimit              bool
	_storagegatewayUpdateBandwidthRateLimitSchedule      bool
	_storagegatewayUpdateChapCredentials                 bool
	_storagegatewayUpdateFileSystemAssociation           bool
	_storagegatewayUpdateGatewayInformation              bool
	_storagegatewayUpdateGatewaySoftwareNow              bool
	_storagegatewayUpdateMaintenanceStartTime            bool
	_storagegatewayUpdateNFSFileShare                    bool
	_storagegatewayUpdateSMBFileShare                    bool
	_storagegatewayUpdateSMBFileShareVisibility          bool
	_storagegatewayUpdateSMBLocalGroups                  bool
	_storagegatewayUpdateSMBSecurityStrategy             bool
	_storagegatewayUpdateSnapshotSchedule                bool
	_storagegatewayUpdateVTLDeviceType                   bool

	_storagegatewayAccessBasedEnumeration               string
	_storagegatewayActivationKey                        string
	_storagegatewayAdminUserList                        []string
	_storagegatewayAuditDestinationARN                  string
	_storagegatewayAuthentication                       string
	_storagegatewayAutomaticTapeCreationRules           string
	_storagegatewayAverageDownloadRateLimitInBitsPerSec string
	_storagegatewayAverageUploadRateLimitInBitsPerSec   string
	_storagegatewayBandwidthRateLimitIntervals          string
	_storagegatewayBandwidthType                        string
	_storagegatewayBucketRegion                         string
	_storagegatewayBypassGovernanceRetention            string
	_storagegatewayCacheAttributes                      string
	_storagegatewayCacheReportARN                       string
	_storagegatewayCaseSensitivity                      string
	_storagegatewayClientList                           []string
	_storagegatewayClientToken                          string
	_storagegatewayCloudWatchLogGroupARN                string
	_storagegatewayDayOfMonth                           string
	_storagegatewayDayOfWeek                            string
	_storagegatewayDefaultStorageClass                  string
	_storagegatewayDescription                          string
	_storagegatewayDeviceType                           string
	_storagegatewayDiskId                               string
	_storagegatewayDiskIds                              []string
	_storagegatewayDomainControllers                    []string
	_storagegatewayDomainName                           string
	_storagegatewayEncryptionType                       string
	_storagegatewayEndpointNetworkConfiguration         string
	_storagegatewayExclusionFilters                     string
	_storagegatewayFileShareARN                         string
	_storagegatewayFileShareARNList                     []string
	_storagegatewayFileShareName                        string
	_storagegatewayFileSharesVisible                    string
	_storagegatewayFileSystemAssociationARN             string
	_storagegatewayFileSystemAssociationARNList         []string
	_storagegatewayFolderList                           []string
	_storagegatewayForceDelete                          string
	_storagegatewayForceDetach                          string
	_storagegatewayForceRemove                          string
	_storagegatewayGatewayARN                           string
	_storagegatewayGatewayCapacity                      string
	_storagegatewayGatewayName                          string
	_storagegatewayGatewayRegion                        string
	_storagegatewayGatewayTimezone                      string
	_storagegatewayGatewayType                          string
	_storagegatewayGuessMIMETypeEnabled                 string
	_storagegatewayHourOfDay                            string
	_storagegatewayInclusionFilters                     string
	_storagegatewayInitiatorName                        string
	_storagegatewayInvalidUserList                      []string
	_storagegatewayKMSEncrypted                         string
	_storagegatewayKMSKey                               string
	_storagegatewayLimit                                string
	_storagegatewayLocalConsolePassword                 string
	_storagegatewayLocationARN                          string
	_storagegatewayMarker                               string
	_storagegatewayMediumChangerType                    string
	_storagegatewayMinuteOfHour                         string
	_storagegatewayNetworkInterfaceId                   string
	_storagegatewayNFSFileShareDefaults                 string
	_storagegatewayNotificationPolicy                   string
	_storagegatewayNumTapesToCreate                     string
	_storagegatewayObjectACL                            string
	_storagegatewayOplocksEnabled                       string
	_storagegatewayOrganizationalUnit                   string
	_storagegatewayPassword                             string
	_storagegatewayPoolARN                              string
	_storagegatewayPoolARNs                             []string
	_storagegatewayPoolId                               string
	_storagegatewayPoolName                             string
	_storagegatewayPreserveExistingData                 string
	_storagegatewayReadOnly                             string
	_storagegatewayRecurrenceInHours                    string
	_storagegatewayRecursive                            string
	_storagegatewayRequesterPays                        string
	_storagegatewayResourceARN                          string
	_storagegatewayRetentionLockTimeInDays              string
	_storagegatewayRetentionLockType                    string
	_storagegatewayRole                                 string
	_storagegatewaySecretToAuthenticateInitiator        string
	_storagegatewaySecretToAuthenticateTarget           string
	_storagegatewaySMBLocalGroups                       string
	_storagegatewaySMBSecurityStrategy                  string
	_storagegatewaySMBACLEnabled                        string
	_storagegatewaySnapshotDescription                  string
	_storagegatewaySnapshotId                           string
	_storagegatewaySoftwareUpdatePreferences            string
	_storagegatewaySourceVolumeARN                      string
	_storagegatewaySquash                               string
	_storagegatewayStartAt                              string
	_storagegatewayStorageClass                         string
	_storagegatewayTagKeys                              []string
	_storagegatewayTags                                 string
	_storagegatewayTapeARN                              string
	_storagegatewayTapeARNs                             []string
	_storagegatewayTapeBarcode                          string
	_storagegatewayTapeBarcodePrefix                    string
	_storagegatewayTapeDriveType                        string
	_storagegatewayTapeSizeInBytes                      string
	_storagegatewayTargetARN                            string
	_storagegatewayTargetName                           string
	_storagegatewayTimeoutInSeconds                     string
	_storagegatewayUserName                             string
	_storagegatewayValidUserList                        []string
	_storagegatewayVolumeARN                            string
	_storagegatewayVolumeARNs                           []string
	_storagegatewayVolumeSizeInBytes                    string
	_storagegatewayVPCEndpointDNSName                   string
	_storagegatewayVTLDeviceARN                         string
	_storagegatewayVTLDeviceARNs                        []string
	_storagegatewayWorm                                 string
)

// Activates the gateway you previously deployed on your host. In the activation
// process, you specify information such as the Amazon Web Services Region that you
// want to use for storing snapshots or tapes, the time zone for scheduled
// snapshots the gateway snapshot schedule window, an activation key, and a name
// for your gateway. The activation process also associates your gateway with your
// account. For more information, see UpdateGatewayInformation.
//
// You must turn on the gateway VM before you can activate your gateway.
func storagegateway_ActivateGateway(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.ActivateGatewayInput{
		// ActivationKey: *string, // Required
		// GatewayName: *string, // Required
		// GatewayRegion: *string, // Required
		// GatewayTimezone: *string, // Required
	}

	if len(_storagegatewayActivationKey) > 0 {
		input.ActivationKey = aws.String(_storagegatewayActivationKey)
	}
	if len(_storagegatewayGatewayName) > 0 {
		input.GatewayName = aws.String(_storagegatewayGatewayName)
	}
	if len(_storagegatewayGatewayRegion) > 0 {
		input.GatewayRegion = aws.String(_storagegatewayGatewayRegion)
	}
	if len(_storagegatewayGatewayTimezone) > 0 {
		input.GatewayTimezone = aws.String(_storagegatewayGatewayTimezone)
	}
	if len(_storagegatewayGatewayType) > 0 {
		input.GatewayType = aws.String(_storagegatewayGatewayType)
	}
	if len(_storagegatewayMediumChangerType) > 0 {
		input.MediumChangerType = aws.String(_storagegatewayMediumChangerType)
	}
	if len(_storagegatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _storagegatewayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayTapeDriveType) > 0 {
		input.TapeDriveType = aws.String(_storagegatewayTapeDriveType)
	}

	if resp, err := client.ActivateGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configures one or more gateway local disks as cache for a gateway. This
// operation is only supported in the cached volume, tape, and file gateway type
// (see [How Storage Gateway works (architecture)].
//
// In the request, you specify the gateway Amazon Resource Name (ARN) to which you
// want to add cache, and one or more disk IDs that you want to configure as cache.
//
// [How Storage Gateway works (architecture)]: https://docs.aws.amazon.com/storagegateway/latest/userguide/StorageGatewayConcepts.html
func storagegateway_AddCache(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.AddCacheInput{
		// DiskIds: []string, // Required
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayDiskIds) > 0 {
		input.DiskIds = append([]string(nil), _storagegatewayDiskIds...)
	}
	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.AddCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more tags to the specified resource. You use tags to add metadata
// to resources, which you can use to categorize these resources. For example, you
// can categorize resources by purpose, owner, environment, or team. Each tag
// consists of a key and a value, which you define. You can add tags to the
// following Storage Gateway resources:
//
// - Storage gateways of all types
//
// - Storage volumes
//
// - Virtual tapes
//
// - NFS and SMB file shares
//
// - File System associations
//
// You can create a maximum of 50 tags for each resource. Virtual tapes and
// storage volumes that are recovered to a new gateway maintain their tags.
func storagegateway_AddTagsToResource(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.AddTagsToResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_storagegatewayResourceARN) > 0 {
		input.ResourceARN = aws.String(_storagegatewayResourceARN)
	}
	if len(_storagegatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _storagegatewayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddTagsToResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configures one or more gateway local disks as upload buffer for a specified
// gateway. This operation is supported for the stored volume, cached volume, and
// tape gateway types.
//
// In the request, you specify the gateway Amazon Resource Name (ARN) to which you
// want to add upload buffer, and one or more disk IDs that you want to configure
// as upload buffer.
func storagegateway_AddUploadBuffer(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.AddUploadBufferInput{
		// DiskIds: []string, // Required
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayDiskIds) > 0 {
		input.DiskIds = append([]string(nil), _storagegatewayDiskIds...)
	}
	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.AddUploadBuffer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configures one or more gateway local disks as working storage for a gateway.
// This operation is only supported in the stored volume gateway type. This
// operation is deprecated in cached volume API version 20120630. Use AddUploadBufferinstead.
//
// Working storage is also referred to as upload buffer. You can also use the AddUploadBuffer
// operation to add upload buffer to a stored volume gateway.
//
// In the request, you specify the gateway Amazon Resource Name (ARN) to which you
// want to add working storage, and one or more disk IDs that you want to configure
// as working storage.
func storagegateway_AddWorkingStorage(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.AddWorkingStorageInput{
		// DiskIds: []string, // Required
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayDiskIds) > 0 {
		input.DiskIds = append([]string(nil), _storagegatewayDiskIds...)
	}
	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.AddWorkingStorage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns a tape to a tape pool for archiving. The tape assigned to a pool is
// archived in the S3 storage class that is associated with the pool. When you use
// your backup application to eject the tape, the tape is archived directly into
// the S3 storage class (S3 Glacier or S3 Glacier Deep Archive) that corresponds to
// the pool.
func storagegateway_AssignTapePool(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.AssignTapePoolInput{
		// PoolId: *string, // Required
		// TapeARN: *string, // Required
	}

	if len(_storagegatewayPoolId) > 0 {
		input.PoolId = aws.String(_storagegatewayPoolId)
	}
	if len(_storagegatewayTapeARN) > 0 {
		input.TapeARN = aws.String(_storagegatewayTapeARN)
	}
	if len(_storagegatewayBypassGovernanceRetention) > 0 {
		if err := assignInputField(input, "BypassGovernanceRetention", _storagegatewayBypassGovernanceRetention); err != nil {
			log.Errorf("invalid --bypass-governance-retention: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssignTapePool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associate an Amazon FSx file system with the FSx File Gateway. After the
// association process is complete, the file shares on the Amazon FSx file system
// are available for access through the gateway. This operation only supports the
// FSx File Gateway type.
func storagegateway_AssociateFileSystem(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.AssociateFileSystemInput{
		// ClientToken: *string, // Required
		// GatewayARN: *string, // Required
		// LocationARN: *string, // Required
		// Password: *string, // Required
		// UserName: *string, // Required
	}

	if len(_storagegatewayClientToken) > 0 {
		input.ClientToken = aws.String(_storagegatewayClientToken)
	}
	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayLocationARN) > 0 {
		input.LocationARN = aws.String(_storagegatewayLocationARN)
	}
	if len(_storagegatewayPassword) > 0 {
		input.Password = aws.String(_storagegatewayPassword)
	}
	if len(_storagegatewayUserName) > 0 {
		input.UserName = aws.String(_storagegatewayUserName)
	}
	if len(_storagegatewayAuditDestinationARN) > 0 {
		input.AuditDestinationARN = aws.String(_storagegatewayAuditDestinationARN)
	}
	if len(_storagegatewayCacheAttributes) > 0 {
		if err := assignInputField(input, "CacheAttributes", _storagegatewayCacheAttributes); err != nil {
			log.Errorf("invalid --cache-attributes: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayEndpointNetworkConfiguration) > 0 {
		if err := assignInputField(input, "EndpointNetworkConfiguration", _storagegatewayEndpointNetworkConfiguration); err != nil {
			log.Errorf("invalid --endpoint-network-configuration: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _storagegatewayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateFileSystem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Connects a volume to an iSCSI connection and then attaches the volume to the
// specified gateway. Detaching and attaching a volume enables you to recover your
// data from one gateway to a different gateway without creating a snapshot. It
// also makes it easier to move your volumes from an on-premises gateway to a
// gateway hosted on an Amazon EC2 instance.
func storagegateway_AttachVolume(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.AttachVolumeInput{
		// GatewayARN: *string, // Required
		// NetworkInterfaceId: *string, // Required
		// VolumeARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayNetworkInterfaceId) > 0 {
		input.NetworkInterfaceId = aws.String(_storagegatewayNetworkInterfaceId)
	}
	if len(_storagegatewayVolumeARN) > 0 {
		input.VolumeARN = aws.String(_storagegatewayVolumeARN)
	}
	if len(_storagegatewayDiskId) > 0 {
		input.DiskId = aws.String(_storagegatewayDiskId)
	}
	if len(_storagegatewayTargetName) > 0 {
		input.TargetName = aws.String(_storagegatewayTargetName)
	}

	if resp, err := client.AttachVolume(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels archiving of a virtual tape to the virtual tape shelf (VTS) after the
// archiving process is initiated. This operation is only supported in the tape
// gateway type.
func storagegateway_CancelArchival(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.CancelArchivalInput{
		// GatewayARN: *string, // Required
		// TapeARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayTapeARN) > 0 {
		input.TapeARN = aws.String(_storagegatewayTapeARN)
	}

	if resp, err := client.CancelArchival(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels generation of a specified cache report. You can use this operation to
// manually cancel an IN-PROGRESS report for any reason. This action changes the
// report status from IN-PROGRESS to CANCELLED. You can only cancel in-progress
// reports. If the the report you attempt to cancel is in FAILED, ERROR, or
// COMPLETED state, the cancel operation returns an error.
func storagegateway_CancelCacheReport(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.CancelCacheReportInput{
		// CacheReportARN: *string, // Required
	}

	if len(_storagegatewayCacheReportARN) > 0 {
		input.CacheReportARN = aws.String(_storagegatewayCacheReportARN)
	}

	if resp, err := client.CancelCacheReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels retrieval of a virtual tape from the virtual tape shelf (VTS) to a
// gateway after the retrieval process is initiated. The virtual tape is returned
// to the VTS. This operation is only supported in the tape gateway type.
func storagegateway_CancelRetrieval(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.CancelRetrievalInput{
		// GatewayARN: *string, // Required
		// TapeARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayTapeARN) > 0 {
		input.TapeARN = aws.String(_storagegatewayTapeARN)
	}

	if resp, err := client.CancelRetrieval(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a cached volume on a specified cached volume gateway. This operation is
// only supported in the cached volume gateway type.
//
// Cache storage must be allocated to the gateway before you can create a cached
// volume. Use the AddCacheoperation to add cache storage to a gateway.
//
// In the request, you must specify the gateway, size of the volume in bytes, the
// iSCSI target name, an IP address on which to expose the target, and a unique
// client token. In response, the gateway creates the volume and returns
// information about it. This information includes the volume Amazon Resource Name
// (ARN), its size, and the iSCSI target ARN that initiators can use to connect to
// the volume target.
//
// Optionally, you can provide the ARN for an existing volume as the
// SourceVolumeARN for this cached volume, which creates an exact copy of the
// existing volume’s latest recovery point. The VolumeSizeInBytes value must be
// equal to or larger than the size of the copied volume, in bytes.
func storagegateway_CreateCachediSCSIVolume(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.CreateCachediSCSIVolumeInput{
		// ClientToken: *string, // Required
		// GatewayARN: *string, // Required
		// NetworkInterfaceId: *string, // Required
		// TargetName: *string, // Required
		// VolumeSizeInBytes: int64, // Required
	}

	if len(_storagegatewayClientToken) > 0 {
		input.ClientToken = aws.String(_storagegatewayClientToken)
	}
	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayNetworkInterfaceId) > 0 {
		input.NetworkInterfaceId = aws.String(_storagegatewayNetworkInterfaceId)
	}
	if len(_storagegatewayTargetName) > 0 {
		input.TargetName = aws.String(_storagegatewayTargetName)
	}
	if len(_storagegatewayVolumeSizeInBytes) > 0 {
		if err := assignInputField(input, "VolumeSizeInBytes", _storagegatewayVolumeSizeInBytes); err != nil {
			log.Errorf("invalid --volume-size-in-bytes: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayKMSEncrypted) > 0 {
		if err := assignInputField(input, "KMSEncrypted", _storagegatewayKMSEncrypted); err != nil {
			log.Errorf("invalid --kms-encrypted: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayKMSKey) > 0 {
		input.KMSKey = aws.String(_storagegatewayKMSKey)
	}
	if len(_storagegatewaySnapshotId) > 0 {
		input.SnapshotId = aws.String(_storagegatewaySnapshotId)
	}
	if len(_storagegatewaySourceVolumeARN) > 0 {
		input.SourceVolumeARN = aws.String(_storagegatewaySourceVolumeARN)
	}
	if len(_storagegatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _storagegatewayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCachediSCSIVolume(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Network File System (NFS) file share on an existing S3 File Gateway.
// In Storage Gateway, a file share is a file system mount point backed by Amazon
// S3 cloud storage. Storage Gateway exposes file shares using an NFS interface.
// This operation is only supported for S3 File Gateways.
//
// S3 File gateway requires Security Token Service (Amazon Web Services STS) to be
// activated to enable you to create a file share. Make sure Amazon Web Services
// STS is activated in the Amazon Web Services Region you are creating your S3 File
// Gateway in. If Amazon Web Services STS is not activated in the Amazon Web
// Services Region, activate it. For information about how to activate Amazon Web
// Services STS, see [Activating and deactivating Amazon Web Services STS in an Amazon Web Services Region]in the Identity and Access Management User Guide.
//
// S3 File Gateways do not support creating hard or symbolic links on a file share.
//
// [Activating and deactivating Amazon Web Services STS in an Amazon Web Services Region]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html
func storagegateway_CreateNFSFileShare(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.CreateNFSFileShareInput{
		// ClientToken: *string, // Required
		// GatewayARN: *string, // Required
		// LocationARN: *string, // Required
		// Role: *string, // Required
	}

	if len(_storagegatewayClientToken) > 0 {
		input.ClientToken = aws.String(_storagegatewayClientToken)
	}
	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayLocationARN) > 0 {
		input.LocationARN = aws.String(_storagegatewayLocationARN)
	}
	if len(_storagegatewayRole) > 0 {
		input.Role = aws.String(_storagegatewayRole)
	}
	if len(_storagegatewayAuditDestinationARN) > 0 {
		input.AuditDestinationARN = aws.String(_storagegatewayAuditDestinationARN)
	}
	if len(_storagegatewayBucketRegion) > 0 {
		input.BucketRegion = aws.String(_storagegatewayBucketRegion)
	}
	if len(_storagegatewayCacheAttributes) > 0 {
		if err := assignInputField(input, "CacheAttributes", _storagegatewayCacheAttributes); err != nil {
			log.Errorf("invalid --cache-attributes: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayClientList) > 0 {
		input.ClientList = append([]string(nil), _storagegatewayClientList...)
	}
	if len(_storagegatewayDefaultStorageClass) > 0 {
		input.DefaultStorageClass = aws.String(_storagegatewayDefaultStorageClass)
	}
	if len(_storagegatewayEncryptionType) > 0 {
		if err := assignInputField(input, "EncryptionType", _storagegatewayEncryptionType); err != nil {
			log.Errorf("invalid --encryption-type: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayFileShareName) > 0 {
		input.FileShareName = aws.String(_storagegatewayFileShareName)
	}
	if len(_storagegatewayGuessMIMETypeEnabled) > 0 {
		if err := assignInputField(input, "GuessMIMETypeEnabled", _storagegatewayGuessMIMETypeEnabled); err != nil {
			log.Errorf("invalid --guess-mime-type-enabled: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayKMSEncrypted) > 0 {
		if err := assignInputField(input, "KMSEncrypted", _storagegatewayKMSEncrypted); err != nil {
			log.Errorf("invalid --kms-encrypted: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayKMSKey) > 0 {
		input.KMSKey = aws.String(_storagegatewayKMSKey)
	}
	if len(_storagegatewayNFSFileShareDefaults) > 0 {
		if err := assignInputField(input, "NFSFileShareDefaults", _storagegatewayNFSFileShareDefaults); err != nil {
			log.Errorf("invalid --nfs-file-share-defaults: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayNotificationPolicy) > 0 {
		input.NotificationPolicy = aws.String(_storagegatewayNotificationPolicy)
	}
	if len(_storagegatewayObjectACL) > 0 {
		if err := assignInputField(input, "ObjectACL", _storagegatewayObjectACL); err != nil {
			log.Errorf("invalid --object-acl: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayReadOnly) > 0 {
		if err := assignInputField(input, "ReadOnly", _storagegatewayReadOnly); err != nil {
			log.Errorf("invalid --read-only: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayRequesterPays) > 0 {
		if err := assignInputField(input, "RequesterPays", _storagegatewayRequesterPays); err != nil {
			log.Errorf("invalid --requester-pays: %s", err.Error())
			return
		}
	}
	if len(_storagegatewaySquash) > 0 {
		input.Squash = aws.String(_storagegatewaySquash)
	}
	if len(_storagegatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _storagegatewayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayVPCEndpointDNSName) > 0 {
		input.VPCEndpointDNSName = aws.String(_storagegatewayVPCEndpointDNSName)
	}

	if resp, err := client.CreateNFSFileShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Server Message Block (SMB) file share on an existing S3 File Gateway.
// In Storage Gateway, a file share is a file system mount point backed by Amazon
// S3 cloud storage. Storage Gateway exposes file shares using an SMB interface.
// This operation is only supported for S3 File Gateways.
//
// S3 File Gateways require Security Token Service (Amazon Web Services STS) to be
// activated to enable you to create a file share. Make sure that Amazon Web
// Services STS is activated in the Amazon Web Services Region you are creating
// your S3 File Gateway in. If Amazon Web Services STS is not activated in this
// Amazon Web Services Region, activate it. For information about how to activate
// Amazon Web Services STS, see [Activating and deactivating Amazon Web Services STS in an Amazon Web Services Region]in the Identity and Access Management User Guide.
//
// File gateways don't support creating hard or symbolic links on a file share.
//
// [Activating and deactivating Amazon Web Services STS in an Amazon Web Services Region]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html
func storagegateway_CreateSMBFileShare(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.CreateSMBFileShareInput{
		// ClientToken: *string, // Required
		// GatewayARN: *string, // Required
		// LocationARN: *string, // Required
		// Role: *string, // Required
	}

	if len(_storagegatewayClientToken) > 0 {
		input.ClientToken = aws.String(_storagegatewayClientToken)
	}
	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayLocationARN) > 0 {
		input.LocationARN = aws.String(_storagegatewayLocationARN)
	}
	if len(_storagegatewayRole) > 0 {
		input.Role = aws.String(_storagegatewayRole)
	}
	if len(_storagegatewayAccessBasedEnumeration) > 0 {
		if err := assignInputField(input, "AccessBasedEnumeration", _storagegatewayAccessBasedEnumeration); err != nil {
			log.Errorf("invalid --access-based-enumeration: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayAdminUserList) > 0 {
		input.AdminUserList = append([]string(nil), _storagegatewayAdminUserList...)
	}
	if len(_storagegatewayAuditDestinationARN) > 0 {
		input.AuditDestinationARN = aws.String(_storagegatewayAuditDestinationARN)
	}
	if len(_storagegatewayAuthentication) > 0 {
		input.Authentication = aws.String(_storagegatewayAuthentication)
	}
	if len(_storagegatewayBucketRegion) > 0 {
		input.BucketRegion = aws.String(_storagegatewayBucketRegion)
	}
	if len(_storagegatewayCacheAttributes) > 0 {
		if err := assignInputField(input, "CacheAttributes", _storagegatewayCacheAttributes); err != nil {
			log.Errorf("invalid --cache-attributes: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayCaseSensitivity) > 0 {
		if err := assignInputField(input, "CaseSensitivity", _storagegatewayCaseSensitivity); err != nil {
			log.Errorf("invalid --case-sensitivity: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayDefaultStorageClass) > 0 {
		input.DefaultStorageClass = aws.String(_storagegatewayDefaultStorageClass)
	}
	if len(_storagegatewayEncryptionType) > 0 {
		if err := assignInputField(input, "EncryptionType", _storagegatewayEncryptionType); err != nil {
			log.Errorf("invalid --encryption-type: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayFileShareName) > 0 {
		input.FileShareName = aws.String(_storagegatewayFileShareName)
	}
	if len(_storagegatewayGuessMIMETypeEnabled) > 0 {
		if err := assignInputField(input, "GuessMIMETypeEnabled", _storagegatewayGuessMIMETypeEnabled); err != nil {
			log.Errorf("invalid --guess-mime-type-enabled: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayInvalidUserList) > 0 {
		input.InvalidUserList = append([]string(nil), _storagegatewayInvalidUserList...)
	}
	if len(_storagegatewayKMSEncrypted) > 0 {
		if err := assignInputField(input, "KMSEncrypted", _storagegatewayKMSEncrypted); err != nil {
			log.Errorf("invalid --kms-encrypted: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayKMSKey) > 0 {
		input.KMSKey = aws.String(_storagegatewayKMSKey)
	}
	if len(_storagegatewayNotificationPolicy) > 0 {
		input.NotificationPolicy = aws.String(_storagegatewayNotificationPolicy)
	}
	if len(_storagegatewayObjectACL) > 0 {
		if err := assignInputField(input, "ObjectACL", _storagegatewayObjectACL); err != nil {
			log.Errorf("invalid --object-acl: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayOplocksEnabled) > 0 {
		if err := assignInputField(input, "OplocksEnabled", _storagegatewayOplocksEnabled); err != nil {
			log.Errorf("invalid --oplocks-enabled: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayReadOnly) > 0 {
		if err := assignInputField(input, "ReadOnly", _storagegatewayReadOnly); err != nil {
			log.Errorf("invalid --read-only: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayRequesterPays) > 0 {
		if err := assignInputField(input, "RequesterPays", _storagegatewayRequesterPays); err != nil {
			log.Errorf("invalid --requester-pays: %s", err.Error())
			return
		}
	}
	if len(_storagegatewaySMBACLEnabled) > 0 {
		if err := assignInputField(input, "SMBACLEnabled", _storagegatewaySMBACLEnabled); err != nil {
			log.Errorf("invalid --smbacl-enabled: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _storagegatewayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayVPCEndpointDNSName) > 0 {
		input.VPCEndpointDNSName = aws.String(_storagegatewayVPCEndpointDNSName)
	}
	if len(_storagegatewayValidUserList) > 0 {
		input.ValidUserList = append([]string(nil), _storagegatewayValidUserList...)
	}

	if resp, err := client.CreateSMBFileShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a snapshot of a volume.
// Storage Gateway provides the ability to back up point-in-time snapshots of your
// data to Amazon Simple Storage (Amazon S3) for durable off-site recovery, and
// also import the data to an Amazon Elastic Block Store (EBS) volume in Amazon
// Elastic Compute Cloud (EC2). You can take snapshots of your gateway volume on a
// scheduled or ad hoc basis. This API enables you to take an ad hoc snapshot. For
// more information, see [Editing a snapshot schedule].
//
// In the CreateSnapshot request, you identify the volume by providing its Amazon
// Resource Name (ARN). You must also provide description for the snapshot. When
// Storage Gateway takes the snapshot of specified volume, the snapshot and
// description appears in the Storage Gateway console. In response, Storage Gateway
// returns you a snapshot ID. You can use this snapshot ID to check the snapshot
// progress or later use it when you want to create a volume from a snapshot. This
// operation is only supported in stored and cached volume gateway type.
//
// To list or delete a snapshot, you must use the Amazon EC2 API. For more
// information, see [DescribeSnapshots]or [DeleteSnapshot] in the Amazon Elastic Compute Cloud API Reference.
//
// Volume and snapshot IDs are changing to a longer length ID format. For more
// information, see the important note on the [Welcome]page.
//
// [Editing a snapshot schedule]: https://docs.aws.amazon.com/storagegateway/latest/userguide/managing-volumes.html#SchedulingSnapshot
// [Welcome]: https://docs.aws.amazon.com/storagegateway/latest/APIReference/Welcome.html
// [DescribeSnapshots]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSnapshots.html
// [DeleteSnapshot]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DeleteSnapshot.html
func storagegateway_CreateSnapshot(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.CreateSnapshotInput{
		// SnapshotDescription: *string, // Required
		// VolumeARN: *string, // Required
	}

	if len(_storagegatewaySnapshotDescription) > 0 {
		input.SnapshotDescription = aws.String(_storagegatewaySnapshotDescription)
	}
	if len(_storagegatewayVolumeARN) > 0 {
		input.VolumeARN = aws.String(_storagegatewayVolumeARN)
	}
	if len(_storagegatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _storagegatewayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a snapshot of a gateway from a volume recovery point. This operation
// is only supported in the cached volume gateway type.
//
// A volume recovery point is a point in time at which all data of the volume is
// consistent and from which you can create a snapshot. To get a list of volume
// recovery point for cached volume gateway, use ListVolumeRecoveryPoints.
//
// In the CreateSnapshotFromVolumeRecoveryPoint request, you identify the volume
// by providing its Amazon Resource Name (ARN). You must also provide a description
// for the snapshot. When the gateway takes a snapshot of the specified volume, the
// snapshot and its description appear in the Storage Gateway console. In response,
// the gateway returns you a snapshot ID. You can use this snapshot ID to check the
// snapshot progress or later use it when you want to create a volume from a
// snapshot.
//
// To list or delete a snapshot, you must use the Amazon EC2 API. For more
// information, see [DescribeSnapshots]or [DeleteSnapshot] in the Amazon Elastic Compute Cloud API Reference.
//
// [DescribeSnapshots]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSnapshots.html
// [DeleteSnapshot]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DeleteSnapshot.html
func storagegateway_CreateSnapshotFromVolumeRecoveryPoint(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.CreateSnapshotFromVolumeRecoveryPointInput{
		// SnapshotDescription: *string, // Required
		// VolumeARN: *string, // Required
	}

	if len(_storagegatewaySnapshotDescription) > 0 {
		input.SnapshotDescription = aws.String(_storagegatewaySnapshotDescription)
	}
	if len(_storagegatewayVolumeARN) > 0 {
		input.VolumeARN = aws.String(_storagegatewayVolumeARN)
	}
	if len(_storagegatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _storagegatewayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSnapshotFromVolumeRecoveryPoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a volume on a specified gateway. This operation is only supported in
// the stored volume gateway type.
//
// The size of the volume to create is inferred from the disk size. You can choose
// to preserve existing data on the disk, create volume from an existing snapshot,
// or create an empty volume. If you choose to create an empty gateway volume, then
// any existing data on the disk is erased.
//
// In the request, you must specify the gateway and the disk information on which
// you are creating the volume. In response, the gateway creates the volume and
// returns volume information such as the volume Amazon Resource Name (ARN), its
// size, and the iSCSI target ARN that initiators can use to connect to the volume
// target.
func storagegateway_CreateStorediSCSIVolume(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.CreateStorediSCSIVolumeInput{
		// DiskId: *string, // Required
		// GatewayARN: *string, // Required
		// NetworkInterfaceId: *string, // Required
		// PreserveExistingData: bool, // Required
		// TargetName: *string, // Required
	}

	if len(_storagegatewayDiskId) > 0 {
		input.DiskId = aws.String(_storagegatewayDiskId)
	}
	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayNetworkInterfaceId) > 0 {
		input.NetworkInterfaceId = aws.String(_storagegatewayNetworkInterfaceId)
	}
	if len(_storagegatewayPreserveExistingData) > 0 {
		if err := assignInputField(input, "PreserveExistingData", _storagegatewayPreserveExistingData); err != nil {
			log.Errorf("invalid --preserve-existing-data: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayTargetName) > 0 {
		input.TargetName = aws.String(_storagegatewayTargetName)
	}
	if len(_storagegatewayKMSEncrypted) > 0 {
		if err := assignInputField(input, "KMSEncrypted", _storagegatewayKMSEncrypted); err != nil {
			log.Errorf("invalid --kms-encrypted: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayKMSKey) > 0 {
		input.KMSKey = aws.String(_storagegatewayKMSKey)
	}
	if len(_storagegatewaySnapshotId) > 0 {
		input.SnapshotId = aws.String(_storagegatewaySnapshotId)
	}
	if len(_storagegatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _storagegatewayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateStorediSCSIVolume(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new custom tape pool. You can use custom tape pool to enable tape
// retention lock on tapes that are archived in the custom pool.
func storagegateway_CreateTapePool(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.CreateTapePoolInput{
		// PoolName: *string, // Required
		// StorageClass: types.TapeStorageClass, // Required
	}

	if len(_storagegatewayPoolName) > 0 {
		input.PoolName = aws.String(_storagegatewayPoolName)
	}
	if len(_storagegatewayStorageClass) > 0 {
		if err := assignInputField(input, "StorageClass", _storagegatewayStorageClass); err != nil {
			log.Errorf("invalid --storage-class: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayRetentionLockTimeInDays) > 0 {
		if err := assignInputField(input, "RetentionLockTimeInDays", _storagegatewayRetentionLockTimeInDays); err != nil {
			log.Errorf("invalid --retention-lock-time-in-days: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayRetentionLockType) > 0 {
		if err := assignInputField(input, "RetentionLockType", _storagegatewayRetentionLockType); err != nil {
			log.Errorf("invalid --retention-lock-type: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _storagegatewayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTapePool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a virtual tape by using your own barcode. You write data to the virtual
// tape and then archive the tape. A barcode is unique and cannot be reused if it
// has already been used on a tape. This applies to barcodes used on deleted tapes.
// This operation is only supported in the tape gateway type.
//
// Cache storage must be allocated to the gateway before you can create a virtual
// tape. Use the AddCacheoperation to add cache storage to a gateway.
func storagegateway_CreateTapeWithBarcode(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.CreateTapeWithBarcodeInput{
		// GatewayARN: *string, // Required
		// TapeBarcode: *string, // Required
		// TapeSizeInBytes: *int64, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayTapeBarcode) > 0 {
		input.TapeBarcode = aws.String(_storagegatewayTapeBarcode)
	}
	if len(_storagegatewayTapeSizeInBytes) > 0 {
		if err := assignInputField(input, "TapeSizeInBytes", _storagegatewayTapeSizeInBytes); err != nil {
			log.Errorf("invalid --tape-size-in-bytes: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayKMSEncrypted) > 0 {
		if err := assignInputField(input, "KMSEncrypted", _storagegatewayKMSEncrypted); err != nil {
			log.Errorf("invalid --kms-encrypted: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayKMSKey) > 0 {
		input.KMSKey = aws.String(_storagegatewayKMSKey)
	}
	if len(_storagegatewayPoolId) > 0 {
		input.PoolId = aws.String(_storagegatewayPoolId)
	}
	if len(_storagegatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _storagegatewayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayWorm) > 0 {
		if err := assignInputField(input, "Worm", _storagegatewayWorm); err != nil {
			log.Errorf("invalid --worm: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTapeWithBarcode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates one or more virtual tapes. You write data to the virtual tapes and then
// archive the tapes. This operation is only supported in the tape gateway type.
//
// Cache storage must be allocated to the gateway before you can create virtual
// tapes. Use the AddCacheoperation to add cache storage to a gateway.
func storagegateway_CreateTapes(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.CreateTapesInput{
		// ClientToken: *string, // Required
		// GatewayARN: *string, // Required
		// NumTapesToCreate: *int32, // Required
		// TapeBarcodePrefix: *string, // Required
		// TapeSizeInBytes: *int64, // Required
	}

	if len(_storagegatewayClientToken) > 0 {
		input.ClientToken = aws.String(_storagegatewayClientToken)
	}
	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayNumTapesToCreate) > 0 {
		if err := assignInputField(input, "NumTapesToCreate", _storagegatewayNumTapesToCreate); err != nil {
			log.Errorf("invalid --num-tapes-to-create: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayTapeBarcodePrefix) > 0 {
		input.TapeBarcodePrefix = aws.String(_storagegatewayTapeBarcodePrefix)
	}
	if len(_storagegatewayTapeSizeInBytes) > 0 {
		if err := assignInputField(input, "TapeSizeInBytes", _storagegatewayTapeSizeInBytes); err != nil {
			log.Errorf("invalid --tape-size-in-bytes: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayKMSEncrypted) > 0 {
		if err := assignInputField(input, "KMSEncrypted", _storagegatewayKMSEncrypted); err != nil {
			log.Errorf("invalid --kms-encrypted: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayKMSKey) > 0 {
		input.KMSKey = aws.String(_storagegatewayKMSKey)
	}
	if len(_storagegatewayPoolId) > 0 {
		input.PoolId = aws.String(_storagegatewayPoolId)
	}
	if len(_storagegatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _storagegatewayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayWorm) > 0 {
		if err := assignInputField(input, "Worm", _storagegatewayWorm); err != nil {
			log.Errorf("invalid --worm: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTapes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the automatic tape creation policy of a gateway. If you delete this
// policy, new virtual tapes must be created manually. Use the Amazon Resource Name
// (ARN) of the gateway in your request to remove the policy.
func storagegateway_DeleteAutomaticTapeCreationPolicy(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DeleteAutomaticTapeCreationPolicyInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.DeleteAutomaticTapeCreationPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the bandwidth rate limits of a gateway. You can delete either the
// upload and download bandwidth rate limit, or you can delete both. If you delete
// only one of the limits, the other limit remains unchanged. To specify which
// gateway to work with, use the Amazon Resource Name (ARN) of the gateway in your
// request. This operation is supported only for the stored volume, cached volume,
// and tape gateway types.
func storagegateway_DeleteBandwidthRateLimit(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DeleteBandwidthRateLimitInput{
		// BandwidthType: *string, // Required
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayBandwidthType) > 0 {
		input.BandwidthType = aws.String(_storagegatewayBandwidthType)
	}
	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.DeleteBandwidthRateLimit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified cache report and any associated tags from the Storage
// Gateway database. You can only delete completed reports. If the status of the
// report you attempt to delete still IN-PROGRESS, the delete operation returns an
// error. You can use CancelCacheReport to cancel an IN-PROGRESS report.
//
// DeleteCacheReport does not delete the report object from your Amazon S3 bucket.
func storagegateway_DeleteCacheReport(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DeleteCacheReportInput{
		// CacheReportARN: *string, // Required
	}

	if len(_storagegatewayCacheReportARN) > 0 {
		input.CacheReportARN = aws.String(_storagegatewayCacheReportARN)
	}

	if resp, err := client.DeleteCacheReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes Challenge-Handshake Authentication Protocol (CHAP) credentials for a
// specified iSCSI target and initiator pair. This operation is supported in volume
// and tape gateway types.
func storagegateway_DeleteChapCredentials(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DeleteChapCredentialsInput{
		// InitiatorName: *string, // Required
		// TargetARN: *string, // Required
	}

	if len(_storagegatewayInitiatorName) > 0 {
		input.InitiatorName = aws.String(_storagegatewayInitiatorName)
	}
	if len(_storagegatewayTargetARN) > 0 {
		input.TargetARN = aws.String(_storagegatewayTargetARN)
	}

	if resp, err := client.DeleteChapCredentials(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a file share from an S3 File Gateway. This operation is only supported
// for S3 File Gateways.
func storagegateway_DeleteFileShare(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DeleteFileShareInput{
		// FileShareARN: *string, // Required
	}

	if len(_storagegatewayFileShareARN) > 0 {
		input.FileShareARN = aws.String(_storagegatewayFileShareARN)
	}
	if len(_storagegatewayForceDelete) > 0 {
		if err := assignInputField(input, "ForceDelete", _storagegatewayForceDelete); err != nil {
			log.Errorf("invalid --force-delete: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteFileShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a gateway. To specify which gateway to delete, use the Amazon Resource
// Name (ARN) of the gateway in your request. The operation deletes the gateway;
// however, it does not delete the gateway virtual machine (VM) from your host
// computer.
//
// After you delete a gateway, you cannot reactivate it. Completed snapshots of
// the gateway volumes are not deleted upon deleting the gateway, however, pending
// snapshots will not complete. After you delete a gateway, your next step is to
// remove it from your environment.
//
// You no longer pay software charges after the gateway is deleted; however, your
// existing Amazon EBS snapshots persist and you will continue to be billed for
// these snapshots. You can choose to remove all remaining Amazon EBS snapshots by
// canceling your Amazon EC2 subscription. If you prefer not to cancel your Amazon
// EC2 subscription, you can delete your snapshots using the Amazon EC2 console.
// For more information, see the [Storage Gateway detail page].
//
// [Storage Gateway detail page]: http://aws.amazon.com/storagegateway
func storagegateway_DeleteGateway(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DeleteGatewayInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.DeleteGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a snapshot of a volume.
// You can take snapshots of your gateway volumes on a scheduled or ad hoc basis.
// This API action enables you to delete a snapshot schedule for a volume. For more
// information, see [Backing up your volumes]. In the DeleteSnapshotSchedule request, you identify the
// volume by providing its Amazon Resource Name (ARN). This operation is only
// supported for cached volume gateway types.
//
// To list or delete a snapshot, you must use the Amazon EC2 API. For more
// information, go to [DescribeSnapshots]in the Amazon Elastic Compute Cloud API Reference.
//
// [Backing up your volumes]: https://docs.aws.amazon.com/storagegateway/latest/userguide/backing-up-volumes.html
// [DescribeSnapshots]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSnapshots.html
func storagegateway_DeleteSnapshotSchedule(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DeleteSnapshotScheduleInput{
		// VolumeARN: *string, // Required
	}

	if len(_storagegatewayVolumeARN) > 0 {
		input.VolumeARN = aws.String(_storagegatewayVolumeARN)
	}

	if resp, err := client.DeleteSnapshotSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified virtual tape. This operation is only supported in the
// tape gateway type.
func storagegateway_DeleteTape(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DeleteTapeInput{
		// GatewayARN: *string, // Required
		// TapeARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayTapeARN) > 0 {
		input.TapeARN = aws.String(_storagegatewayTapeARN)
	}
	if len(_storagegatewayBypassGovernanceRetention) > 0 {
		if err := assignInputField(input, "BypassGovernanceRetention", _storagegatewayBypassGovernanceRetention); err != nil {
			log.Errorf("invalid --bypass-governance-retention: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteTape(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified virtual tape from the virtual tape shelf (VTS). This
// operation is only supported in the tape gateway type.
func storagegateway_DeleteTapeArchive(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DeleteTapeArchiveInput{
		// TapeARN: *string, // Required
	}

	if len(_storagegatewayTapeARN) > 0 {
		input.TapeARN = aws.String(_storagegatewayTapeARN)
	}
	if len(_storagegatewayBypassGovernanceRetention) > 0 {
		if err := assignInputField(input, "BypassGovernanceRetention", _storagegatewayBypassGovernanceRetention); err != nil {
			log.Errorf("invalid --bypass-governance-retention: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteTapeArchive(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a custom tape pool. A custom tape pool can only be deleted if there are
// no tapes in the pool and if there are no automatic tape creation policies that
// reference the custom tape pool.
func storagegateway_DeleteTapePool(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DeleteTapePoolInput{
		// PoolARN: *string, // Required
	}

	if len(_storagegatewayPoolARN) > 0 {
		input.PoolARN = aws.String(_storagegatewayPoolARN)
	}

	if resp, err := client.DeleteTapePool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified storage volume that you previously created using the CreateCachediSCSIVolume or CreateStorediSCSIVolume
// API. This operation is only supported in the cached volume and stored volume
// types. For stored volume gateways, the local disk that was configured as the
// storage volume is not deleted. You can reuse the local disk to create another
// storage volume.
//
// Before you delete a volume, make sure there are no iSCSI connections to the
// volume you are deleting. You should also make sure there is no snapshot in
// progress. You can use the Amazon Elastic Compute Cloud (Amazon EC2) API to query
// snapshots on the volume you are deleting and check the snapshot status. For more
// information, go to [DescribeSnapshots]in the Amazon Elastic Compute Cloud API Reference.
//
// In the request, you must provide the Amazon Resource Name (ARN) of the storage
// volume you want to delete.
//
// [DescribeSnapshots]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSnapshots.html
func storagegateway_DeleteVolume(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DeleteVolumeInput{
		// VolumeARN: *string, // Required
	}

	if len(_storagegatewayVolumeARN) > 0 {
		input.VolumeARN = aws.String(_storagegatewayVolumeARN)
	}

	if resp, err := client.DeleteVolume(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the most recent high availability monitoring test
// that was performed on the host in a cluster. If a test isn't performed, the
// status and start time in the response would be null.
func storagegateway_DescribeAvailabilityMonitorTest(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DescribeAvailabilityMonitorTestInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.DescribeAvailabilityMonitorTest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the bandwidth rate limits of a gateway. By default, these limits are
// not set, which means no bandwidth rate limiting is in effect. This operation is
// supported only for the stored volume, cached volume, and tape gateway types. To
// describe bandwidth rate limits for S3 file gateways, use DescribeBandwidthRateLimitSchedule.
//
// This operation returns a value for a bandwidth rate limit only if the limit is
// set. If no limits are set for the gateway, then this operation returns only the
// gateway ARN in the response body. To specify which gateway to describe, use the
// Amazon Resource Name (ARN) of the gateway in your request.
func storagegateway_DescribeBandwidthRateLimit(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DescribeBandwidthRateLimitInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.DescribeBandwidthRateLimit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the bandwidth rate limit schedule of a gateway. By
// default, gateways do not have bandwidth rate limit schedules, which means no
// bandwidth rate limiting is in effect. This operation is supported only for
// volume, tape and S3 file gateways. FSx file gateways do not support bandwidth
// rate limits.
//
// This operation returns information about a gateway's bandwidth rate limit
// schedule. A bandwidth rate limit schedule consists of one or more bandwidth rate
// limit intervals. A bandwidth rate limit interval defines a period of time on one
// or more days of the week, during which bandwidth rate limits are specified for
// uploading, downloading, or both.
//
// A bandwidth rate limit interval consists of one or more days of the week, a
// start hour and minute, an ending hour and minute, and bandwidth rate limits for
// uploading and downloading
//
// If no bandwidth rate limit schedule intervals are set for the gateway, this
// operation returns an empty response. To specify which gateway to describe, use
// the Amazon Resource Name (ARN) of the gateway in your request.
func storagegateway_DescribeBandwidthRateLimitSchedule(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DescribeBandwidthRateLimitScheduleInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.DescribeBandwidthRateLimitSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the cache of a gateway. This operation is only
// supported in the cached volume, tape, and file gateway types.
//
// The response includes disk IDs that are configured as cache, and it includes
// the amount of cache allocated and used.
func storagegateway_DescribeCache(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DescribeCacheInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.DescribeCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified cache report, including completion
// status and generation progress.
func storagegateway_DescribeCacheReport(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DescribeCacheReportInput{
		// CacheReportARN: *string, // Required
	}

	if len(_storagegatewayCacheReportARN) > 0 {
		input.CacheReportARN = aws.String(_storagegatewayCacheReportARN)
	}

	if resp, err := client.DescribeCacheReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a description of the gateway volumes specified in the request. This
// operation is only supported in the cached volume gateway types.
//
// The list of gateway volumes in the request must be from one gateway. In the
// response, Storage Gateway returns volume information sorted by volume Amazon
// Resource Name (ARN).
func storagegateway_DescribeCachediSCSIVolumes(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DescribeCachediSCSIVolumesInput{
		// VolumeARNs: []string, // Required
	}

	if len(_storagegatewayVolumeARNs) > 0 {
		input.VolumeARNs = append([]string(nil), _storagegatewayVolumeARNs...)
	}

	if resp, err := client.DescribeCachediSCSIVolumes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an array of Challenge-Handshake Authentication Protocol (CHAP)
// credentials information for a specified iSCSI target, one for each
// target-initiator pair. This operation is supported in the volume and tape
// gateway types.
func storagegateway_DescribeChapCredentials(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DescribeChapCredentialsInput{
		// TargetARN: *string, // Required
	}

	if len(_storagegatewayTargetARN) > 0 {
		input.TargetARN = aws.String(_storagegatewayTargetARN)
	}

	if resp, err := client.DescribeChapCredentials(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the file system association information. This operation is only supported
// for FSx File Gateways.
func storagegateway_DescribeFileSystemAssociations(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DescribeFileSystemAssociationsInput{
		// FileSystemAssociationARNList: []string, // Required
	}

	if len(_storagegatewayFileSystemAssociationARNList) > 0 {
		input.FileSystemAssociationARNList = append([]string(nil), _storagegatewayFileSystemAssociationARNList...)
	}

	if resp, err := client.DescribeFileSystemAssociations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns metadata about a gateway such as its name, network interfaces, time
// zone, status, and software version. To specify which gateway to describe, use
// the Amazon Resource Name (ARN) of the gateway in your request.
func storagegateway_DescribeGatewayInformation(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DescribeGatewayInformationInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.DescribeGatewayInformation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns your gateway's maintenance window schedule information, with values for
// monthly or weekly cadence, specific day and time to begin maintenance, and which
// types of updates to apply. Time values returned are for the gateway's time zone.
func storagegateway_DescribeMaintenanceStartTime(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DescribeMaintenanceStartTimeInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.DescribeMaintenanceStartTime(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a description for one or more Network File System (NFS) file shares from
// an S3 File Gateway. This operation is only supported for S3 File Gateways.
func storagegateway_DescribeNFSFileShares(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DescribeNFSFileSharesInput{
		// FileShareARNList: []string, // Required
	}

	if len(_storagegatewayFileShareARNList) > 0 {
		input.FileShareARNList = append([]string(nil), _storagegatewayFileShareARNList...)
	}

	if resp, err := client.DescribeNFSFileShares(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a description for one or more Server Message Block (SMB) file shares from
// a S3 File Gateway. This operation is only supported for S3 File Gateways.
func storagegateway_DescribeSMBFileShares(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DescribeSMBFileSharesInput{
		// FileShareARNList: []string, // Required
	}

	if len(_storagegatewayFileShareARNList) > 0 {
		input.FileShareARNList = append([]string(nil), _storagegatewayFileShareARNList...)
	}

	if resp, err := client.DescribeSMBFileShares(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a description of a Server Message Block (SMB) file share settings from a
// file gateway. This operation is only supported for file gateways.
func storagegateway_DescribeSMBSettings(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DescribeSMBSettingsInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.DescribeSMBSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the snapshot schedule for the specified gateway volume. The snapshot
// schedule information includes intervals at which snapshots are automatically
// initiated on the volume. This operation is only supported in the cached volume
// and stored volume types.
func storagegateway_DescribeSnapshotSchedule(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DescribeSnapshotScheduleInput{
		// VolumeARN: *string, // Required
	}

	if len(_storagegatewayVolumeARN) > 0 {
		input.VolumeARN = aws.String(_storagegatewayVolumeARN)
	}

	if resp, err := client.DescribeSnapshotSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the description of the gateway volumes specified in the request. The
// list of gateway volumes in the request must be from one gateway. In the
// response, Storage Gateway returns volume information sorted by volume ARNs. This
// operation is only supported in stored volume gateway type.
func storagegateway_DescribeStorediSCSIVolumes(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DescribeStorediSCSIVolumesInput{
		// VolumeARNs: []string, // Required
	}

	if len(_storagegatewayVolumeARNs) > 0 {
		input.VolumeARNs = append([]string(nil), _storagegatewayVolumeARNs...)
	}

	if resp, err := client.DescribeStorediSCSIVolumes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a description of specified virtual tapes in the virtual tape shelf
// (VTS). This operation is only supported in the tape gateway type.
//
// If a specific TapeARN is not specified, Storage Gateway returns a description
// of all virtual tapes found in the VTS associated with your account.
func storagegateway_DescribeTapeArchives(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DescribeTapeArchivesInput{}

	if len(_storagegatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _storagegatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayMarker) > 0 {
		input.Marker = aws.String(_storagegatewayMarker)
	}
	if len(_storagegatewayTapeARNs) > 0 {
		input.TapeARNs = append([]string(nil), _storagegatewayTapeARNs...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeTapeArchives(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*storagegateway.DescribeTapeArchivesOutput
	p := storagegateway.NewDescribeTapeArchivesPaginator(client, input)
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

// Returns a list of virtual tape recovery points that are available for the
// specified tape gateway.
//
// A recovery point is a point-in-time view of a virtual tape at which all the
// data on the virtual tape is consistent. If your gateway crashes, virtual tapes
// that have recovery points can be recovered to a new gateway. This operation is
// only supported in the tape gateway type.
func storagegateway_DescribeTapeRecoveryPoints(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DescribeTapeRecoveryPointsInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _storagegatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayMarker) > 0 {
		input.Marker = aws.String(_storagegatewayMarker)
	}

	if disablePaginator() {
		if resp, err := client.DescribeTapeRecoveryPoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*storagegateway.DescribeTapeRecoveryPointsOutput
	p := storagegateway.NewDescribeTapeRecoveryPointsPaginator(client, input)
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

// Returns a description of virtual tapes that correspond to the specified Amazon
// Resource Names (ARNs). If TapeARN is not specified, returns a description of
// the virtual tapes associated with the specified gateway. This operation is only
// supported for the tape gateway type.
//
// The operation supports pagination. By default, the operation returns a maximum
// of up to 100 tapes. You can optionally specify the Limit field in the body to
// limit the number of tapes in the response. If the number of tapes returned in
// the response is truncated, the response includes a Marker field. You can use
// this Marker value in your subsequent request to retrieve the next set of tapes.
func storagegateway_DescribeTapes(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DescribeTapesInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _storagegatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayMarker) > 0 {
		input.Marker = aws.String(_storagegatewayMarker)
	}
	if len(_storagegatewayTapeARNs) > 0 {
		input.TapeARNs = append([]string(nil), _storagegatewayTapeARNs...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeTapes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*storagegateway.DescribeTapesOutput
	p := storagegateway.NewDescribeTapesPaginator(client, input)
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

// Returns information about the upload buffer of a gateway. This operation is
// supported for the stored volume, cached volume, and tape gateway types.
//
// The response includes disk IDs that are configured as upload buffer space, and
// it includes the amount of upload buffer space allocated and used.
func storagegateway_DescribeUploadBuffer(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DescribeUploadBufferInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.DescribeUploadBuffer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a description of virtual tape library (VTL) devices for the specified
// tape gateway. In the response, Storage Gateway returns VTL device information.
//
// This operation is only supported in the tape gateway type.
func storagegateway_DescribeVTLDevices(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DescribeVTLDevicesInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _storagegatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayMarker) > 0 {
		input.Marker = aws.String(_storagegatewayMarker)
	}
	if len(_storagegatewayVTLDeviceARNs) > 0 {
		input.VTLDeviceARNs = append([]string(nil), _storagegatewayVTLDeviceARNs...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeVTLDevices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*storagegateway.DescribeVTLDevicesOutput
	p := storagegateway.NewDescribeVTLDevicesPaginator(client, input)
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

// Returns information about the working storage of a gateway. This operation is
// only supported in the stored volumes gateway type. This operation is deprecated
// in cached volumes API version (20120630). Use DescribeUploadBuffer instead.
//
// Working storage is also referred to as upload buffer. You can also use the
// DescribeUploadBuffer operation to add upload buffer to a stored volume gateway.
//
// The response includes disk IDs that are configured as working storage, and it
// includes the amount of working storage allocated and used.
func storagegateway_DescribeWorkingStorage(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DescribeWorkingStorageInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.DescribeWorkingStorage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disconnects a volume from an iSCSI connection and then detaches the volume from
// the specified gateway. Detaching and attaching a volume enables you to recover
// your data from one gateway to a different gateway without creating a snapshot.
// It also makes it easier to move your volumes from an on-premises gateway to a
// gateway hosted on an Amazon EC2 instance. This operation is only supported in
// the volume gateway type.
func storagegateway_DetachVolume(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DetachVolumeInput{
		// VolumeARN: *string, // Required
	}

	if len(_storagegatewayVolumeARN) > 0 {
		input.VolumeARN = aws.String(_storagegatewayVolumeARN)
	}
	if len(_storagegatewayForceDetach) > 0 {
		if err := assignInputField(input, "ForceDetach", _storagegatewayForceDetach); err != nil {
			log.Errorf("invalid --force-detach: %s", err.Error())
			return
		}
	}

	if resp, err := client.DetachVolume(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables a tape gateway when the gateway is no longer functioning. For example,
// if your gateway VM is damaged, you can disable the gateway so you can recover
// virtual tapes.
//
// Use this operation for a tape gateway that is not reachable or not functioning.
// This operation is only supported in the tape gateway type.
//
// After a gateway is disabled, it cannot be enabled.
func storagegateway_DisableGateway(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DisableGatewayInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.DisableGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates an Amazon FSx file system from the specified gateway. After the
// disassociation process finishes, the gateway can no longer access the Amazon FSx
// file system. This operation is only supported in the FSx File Gateway type.
func storagegateway_DisassociateFileSystem(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.DisassociateFileSystemInput{
		// FileSystemAssociationARN: *string, // Required
	}

	if len(_storagegatewayFileSystemAssociationARN) > 0 {
		input.FileSystemAssociationARN = aws.String(_storagegatewayFileSystemAssociationARN)
	}
	if len(_storagegatewayForceDelete) > 0 {
		if err := assignInputField(input, "ForceDelete", _storagegatewayForceDelete); err != nil {
			log.Errorf("invalid --force-delete: %s", err.Error())
			return
		}
	}

	if resp, err := client.DisassociateFileSystem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a process that cleans the specified file share's cache of file entries
// that are failing upload to Amazon S3. This API operation reports success if the
// request is received with valid arguments, and there are no other cache clean
// operations currently in-progress for the specified file share. After a
// successful request, the cache clean operation occurs asynchronously and reports
// progress using CloudWatch logs and notifications.
//
// If ForceRemove is set to True , the cache clean operation will delete file data
// from the gateway which might otherwise be recoverable. We recommend using this
// operation only after all other methods to clear files failing upload have been
// exhausted, and if your business need outweighs the potential data loss.
func storagegateway_EvictFilesFailingUpload(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.EvictFilesFailingUploadInput{
		// FileShareARN: *string, // Required
	}

	if len(_storagegatewayFileShareARN) > 0 {
		input.FileShareARN = aws.String(_storagegatewayFileShareARN)
	}
	if len(_storagegatewayForceRemove) > 0 {
		if err := assignInputField(input, "ForceRemove", _storagegatewayForceRemove); err != nil {
			log.Errorf("invalid --force-remove: %s", err.Error())
			return
		}
	}

	if resp, err := client.EvictFilesFailingUpload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a file gateway to an Active Directory domain. This operation is only
// supported for file gateways that support the SMB file protocol.
//
// Joining a domain creates an Active Directory computer account in the default
// organizational unit, using the gateway's Gateway ID as the account name (for
// example, SGW-1234ADE). If your Active Directory environment requires that you
// pre-stage accounts to facilitate the join domain process, you will need to
// create this account ahead of time.
//
// To create the gateway's computer account in an organizational unit other than
// the default, you must specify the organizational unit when joining the domain.
func storagegateway_JoinDomain(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.JoinDomainInput{
		// DomainName: *string, // Required
		// GatewayARN: *string, // Required
		// Password: *string, // Required
		// UserName: *string, // Required
	}

	if len(_storagegatewayDomainName) > 0 {
		input.DomainName = aws.String(_storagegatewayDomainName)
	}
	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayPassword) > 0 {
		input.Password = aws.String(_storagegatewayPassword)
	}
	if len(_storagegatewayUserName) > 0 {
		input.UserName = aws.String(_storagegatewayUserName)
	}
	if len(_storagegatewayDomainControllers) > 0 {
		input.DomainControllers = append([]string(nil), _storagegatewayDomainControllers...)
	}
	if len(_storagegatewayOrganizationalUnit) > 0 {
		input.OrganizationalUnit = aws.String(_storagegatewayOrganizationalUnit)
	}
	if len(_storagegatewayTimeoutInSeconds) > 0 {
		if err := assignInputField(input, "TimeoutInSeconds", _storagegatewayTimeoutInSeconds); err != nil {
			log.Errorf("invalid --timeout-in-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.JoinDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the automatic tape creation policies for a gateway. If there are no
// automatic tape creation policies for the gateway, it returns an empty list.
//
// This operation is only supported for tape gateways.
func storagegateway_ListAutomaticTapeCreationPolicies(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.ListAutomaticTapeCreationPoliciesInput{}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.ListAutomaticTapeCreationPolicies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of existing cache reports for all file shares associated with
// your Amazon Web Services account. This list includes all information provided by
// the DescribeCacheReport action, such as report name, status, completion
// progress, start time, end time, filters, and tags.
func storagegateway_ListCacheReports(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.ListCacheReportsInput{}

	if len(_storagegatewayMarker) > 0 {
		input.Marker = aws.String(_storagegatewayMarker)
	}

	if disablePaginator() {
		if resp, err := client.ListCacheReports(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*storagegateway.ListCacheReportsOutput
	p := storagegateway.NewListCacheReportsPaginator(client, input)
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

// Gets a list of the file shares for a specific S3 File Gateway, or the list of
// file shares that belong to the calling Amazon Web Services account. This
// operation is only supported for S3 File Gateways.
func storagegateway_ListFileShares(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.ListFileSharesInput{}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _storagegatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayMarker) > 0 {
		input.Marker = aws.String(_storagegatewayMarker)
	}

	if disablePaginator() {
		if resp, err := client.ListFileShares(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*storagegateway.ListFileSharesOutput
	p := storagegateway.NewListFileSharesPaginator(client, input)
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

// Gets a list of FileSystemAssociationSummary objects. Each object contains a
// summary of a file system association. This operation is only supported for FSx
// File Gateways.
func storagegateway_ListFileSystemAssociations(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.ListFileSystemAssociationsInput{}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _storagegatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayMarker) > 0 {
		input.Marker = aws.String(_storagegatewayMarker)
	}

	if disablePaginator() {
		if resp, err := client.ListFileSystemAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*storagegateway.ListFileSystemAssociationsOutput
	p := storagegateway.NewListFileSystemAssociationsPaginator(client, input)
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

// Lists gateways owned by an Amazon Web Services account in an Amazon Web
// Services Region specified in the request. The returned list is ordered by
// gateway Amazon Resource Name (ARN).
//
// By default, the operation returns a maximum of 100 gateways. This operation
// supports pagination that allows you to optionally reduce the number of gateways
// returned in a response.
//
// If you have more gateways than are returned in a response (that is, the
// response returns only a truncated list of your gateways), the response contains
// a marker that you can specify in your next request to fetch the next page of
// gateways.
func storagegateway_ListGateways(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.ListGatewaysInput{}

	if len(_storagegatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _storagegatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayMarker) > 0 {
		input.Marker = aws.String(_storagegatewayMarker)
	}

	if disablePaginator() {
		if resp, err := client.ListGateways(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*storagegateway.ListGatewaysOutput
	p := storagegateway.NewListGatewaysPaginator(client, input)
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

// Returns a list of the gateway's local disks. To specify which gateway to
// describe, you use the Amazon Resource Name (ARN) of the gateway in the body of
// the request.
//
// The request returns a list of all disks, specifying which are configured as
// working storage, cache storage, or stored volume or not configured at all. The
// response includes a DiskStatus field. This field can have a value of present
// (the disk is available to use), missing (the disk is no longer connected to the
// gateway), or mismatch (the disk node is occupied by a disk that has incorrect
// metadata or the disk content is corrupted).
func storagegateway_ListLocalDisks(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.ListLocalDisksInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.ListLocalDisks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the tags that have been added to the specified resource. This operation
// is supported in storage gateways of all types.
func storagegateway_ListTagsForResource(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_storagegatewayResourceARN) > 0 {
		input.ResourceARN = aws.String(_storagegatewayResourceARN)
	}
	if len(_storagegatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _storagegatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayMarker) > 0 {
		input.Marker = aws.String(_storagegatewayMarker)
	}

	if disablePaginator() {
		if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*storagegateway.ListTagsForResourceOutput
	p := storagegateway.NewListTagsForResourcePaginator(client, input)
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

// Lists custom tape pools. You specify custom tape pools to list by specifying
// one or more custom tape pool Amazon Resource Names (ARNs). If you don't specify
// a custom tape pool ARN, the operation lists all custom tape pools.
//
// This operation supports pagination. You can optionally specify the Limit
// parameter in the body to limit the number of tape pools in the response. If the
// number of tape pools returned in the response is truncated, the response
// includes a Marker element that you can use in your subsequent request to
// retrieve the next set of tape pools.
func storagegateway_ListTapePools(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.ListTapePoolsInput{}

	if len(_storagegatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _storagegatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayMarker) > 0 {
		input.Marker = aws.String(_storagegatewayMarker)
	}
	if len(_storagegatewayPoolARNs) > 0 {
		input.PoolARNs = append([]string(nil), _storagegatewayPoolARNs...)
	}

	if disablePaginator() {
		if resp, err := client.ListTapePools(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*storagegateway.ListTapePoolsOutput
	p := storagegateway.NewListTapePoolsPaginator(client, input)
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

// Lists virtual tapes in your virtual tape library (VTL) and your virtual tape
// shelf (VTS). You specify the tapes to list by specifying one or more tape Amazon
// Resource Names (ARNs). If you don't specify a tape ARN, the operation lists all
// virtual tapes in both your VTL and VTS.
//
// This operation supports pagination. By default, the operation returns a maximum
// of up to 100 tapes. You can optionally specify the Limit parameter in the body
// to limit the number of tapes in the response. If the number of tapes returned in
// the response is truncated, the response includes a Marker element that you can
// use in your subsequent request to retrieve the next set of tapes. This operation
// is only supported in the tape gateway type.
func storagegateway_ListTapes(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.ListTapesInput{}

	if len(_storagegatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _storagegatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayMarker) > 0 {
		input.Marker = aws.String(_storagegatewayMarker)
	}
	if len(_storagegatewayTapeARNs) > 0 {
		input.TapeARNs = append([]string(nil), _storagegatewayTapeARNs...)
	}

	if disablePaginator() {
		if resp, err := client.ListTapes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*storagegateway.ListTapesOutput
	p := storagegateway.NewListTapesPaginator(client, input)
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

// Lists iSCSI initiators that are connected to a volume. You can use this
// operation to determine whether a volume is being used or not. This operation is
// only supported in the cached volume and stored volume gateway types.
func storagegateway_ListVolumeInitiators(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.ListVolumeInitiatorsInput{
		// VolumeARN: *string, // Required
	}

	if len(_storagegatewayVolumeARN) > 0 {
		input.VolumeARN = aws.String(_storagegatewayVolumeARN)
	}

	if resp, err := client.ListVolumeInitiators(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the recovery points for a specified gateway. This operation is only
// supported in the cached volume gateway type.
//
// Each cache volume has one recovery point. A volume recovery point is a point in
// time at which all data of the volume is consistent and from which you can create
// a snapshot or clone a new cached volume from a source volume. To create a
// snapshot from a volume recovery point use the CreateSnapshotFromVolumeRecoveryPointoperation.
func storagegateway_ListVolumeRecoveryPoints(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.ListVolumeRecoveryPointsInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.ListVolumeRecoveryPoints(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the iSCSI stored volumes of a gateway. Results are sorted by volume ARN.
// The response includes only the volume ARNs. If you want additional volume
// information, use the DescribeStorediSCSIVolumesor the DescribeCachediSCSIVolumes API.
//
// The operation supports pagination. By default, the operation returns a maximum
// of up to 100 volumes. You can optionally specify the Limit field in the body to
// limit the number of volumes in the response. If the number of volumes returned
// in the response is truncated, the response includes a Marker field. You can use
// this Marker value in your subsequent request to retrieve the next set of
// volumes. This operation is only supported in the cached volume and stored volume
// gateway types.
func storagegateway_ListVolumes(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.ListVolumesInput{}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _storagegatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayMarker) > 0 {
		input.Marker = aws.String(_storagegatewayMarker)
	}

	if disablePaginator() {
		if resp, err := client.ListVolumes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*storagegateway.ListVolumesOutput
	p := storagegateway.NewListVolumesPaginator(client, input)
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

// Sends you notification through Amazon EventBridge when all files written to
// your file share have been uploaded to Amazon S3.
//
// Storage Gateway can send a notification through Amazon EventBridge when all
// files written to your file share up to that point in time have been uploaded to
// Amazon S3. These files include files written to the file share up to the time
// that you make a request for notification. When the upload is done, Storage
// Gateway sends you notification through EventBridge. You can configure
// EventBridge to send the notification through event targets such as Amazon SNS or
// Lambda function. This operation is only supported for S3 File Gateways.
//
// For more information, see [Getting file upload notification] in the Amazon S3 File Gateway User Guide.
//
// [Getting file upload notification]: https://docs.aws.amazon.com/filegateway/latest/files3/monitoring-file-gateway.html#get-notification
func storagegateway_NotifyWhenUploaded(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.NotifyWhenUploadedInput{
		// FileShareARN: *string, // Required
	}

	if len(_storagegatewayFileShareARN) > 0 {
		input.FileShareARN = aws.String(_storagegatewayFileShareARN)
	}

	if resp, err := client.NotifyWhenUploaded(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Refreshes the cached inventory of objects for the specified file share. This
// operation finds objects in the Amazon S3 bucket that were added, removed, or
// replaced since the gateway last listed the bucket's contents and cached the
// results. This operation does not import files into the S3 File Gateway cache
// storage. It only updates the cached inventory to reflect changes in the
// inventory of the objects in the S3 bucket. This operation is only supported in
// the S3 File Gateway types.
//
// You can subscribe to be notified through an Amazon CloudWatch event when your
// RefreshCache operation completes. For more information, see [Getting notified about file operations] in the Amazon S3
// File Gateway User Guide. This operation is Only supported for S3 File Gateways.
//
// When this API is called, it only initiates the refresh operation. When the API
// call completes and returns a success code, it doesn't necessarily mean that the
// file refresh has completed. You should use the refresh-complete notification to
// determine that the operation has completed before you check for new files on the
// gateway file share. You can subscribe to be notified through a CloudWatch event
// when your RefreshCache operation completes.
//
// Throttle limit: This API is asynchronous, so the gateway will accept no more
// than two refreshes at any time. We recommend using the refresh-complete
// CloudWatch event notification before issuing additional requests. For more
// information, see [Getting notified about file operations]in the Amazon S3 File Gateway User Guide.
//
// - Wait at least 60 seconds between consecutive RefreshCache API requests.
//
// - If you invoke the RefreshCache API when two requests are already being
// processed, any new request will cause an InvalidGatewayRequestException error
// because too many requests were sent to the server.
//
// The S3 bucket name does not need to be included when entering the list of
// folders in the FolderList parameter.
//
// For more information, see [Getting notified about file operations] in the Amazon S3 File Gateway User Guide.
//
// [Getting notified about file operations]: https://docs.aws.amazon.com/filegateway/latest/files3/monitoring-file-gateway.html#get-notification
func storagegateway_RefreshCache(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.RefreshCacheInput{
		// FileShareARN: *string, // Required
	}

	if len(_storagegatewayFileShareARN) > 0 {
		input.FileShareARN = aws.String(_storagegatewayFileShareARN)
	}
	if len(_storagegatewayFolderList) > 0 {
		input.FolderList = append([]string(nil), _storagegatewayFolderList...)
	}
	if len(_storagegatewayRecursive) > 0 {
		if err := assignInputField(input, "Recursive", _storagegatewayRecursive); err != nil {
			log.Errorf("invalid --recursive: %s", err.Error())
			return
		}
	}

	if resp, err := client.RefreshCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes one or more tags from the specified resource. This operation is
// supported in storage gateways of all types.
func storagegateway_RemoveTagsFromResource(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.RemoveTagsFromResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_storagegatewayResourceARN) > 0 {
		input.ResourceARN = aws.String(_storagegatewayResourceARN)
	}
	if len(_storagegatewayTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _storagegatewayTagKeys...)
	}

	if resp, err := client.RemoveTagsFromResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resets all cache disks that have encountered an error and makes the disks
// available for reconfiguration as cache storage. If your cache disk encounters an
// error, the gateway prevents read and write operations on virtual tapes in the
// gateway. For example, an error can occur when a disk is corrupted or removed
// from the gateway. When a cache is reset, the gateway loses its cache storage. At
// this point, you can reconfigure the disks as cache disks. This operation is only
// supported in the cached volume and tape types.
//
// If the cache disk you are resetting contains data that has not been uploaded to
// Amazon S3 yet, that data can be lost. After you reset cache disks, there will be
// no configured cache disks left in the gateway, so you must configure at least
// one new cache disk for your gateway to function properly.
func storagegateway_ResetCache(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.ResetCacheInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.ResetCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves an archived virtual tape from the virtual tape shelf (VTS) to a tape
// gateway. Virtual tapes archived in the VTS are not associated with any gateway.
// However after a tape is retrieved, it is associated with a gateway, even though
// it is also listed in the VTS, that is, archive. This operation is only supported
// in the tape gateway type.
//
// Once a tape is successfully retrieved to a gateway, it cannot be retrieved
// again to another gateway. You must archive the tape again before you can
// retrieve it to another gateway. This operation is only supported in the tape
// gateway type.
func storagegateway_RetrieveTapeArchive(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.RetrieveTapeArchiveInput{
		// GatewayARN: *string, // Required
		// TapeARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayTapeARN) > 0 {
		input.TapeARN = aws.String(_storagegatewayTapeARN)
	}

	if resp, err := client.RetrieveTapeArchive(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the recovery point for the specified virtual tape. This operation is
// only supported in the tape gateway type.
//
// A recovery point is a point in time view of a virtual tape at which all the
// data on the tape is consistent. If your gateway crashes, virtual tapes that have
// recovery points can be recovered to a new gateway.
//
// The virtual tape can be retrieved to only one gateway. The retrieved tape is
// read-only. The virtual tape can be retrieved to only a tape gateway. There is no
// charge for retrieving recovery points.
func storagegateway_RetrieveTapeRecoveryPoint(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.RetrieveTapeRecoveryPointInput{
		// GatewayARN: *string, // Required
		// TapeARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayTapeARN) > 0 {
		input.TapeARN = aws.String(_storagegatewayTapeARN)
	}

	if resp, err := client.RetrieveTapeRecoveryPoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the password for your VM local console. When you log in to the local
// console for the first time, you log in to the VM with the default credentials.
// We recommend that you set a new password. You don't need to know the default
// password to set a new password.
func storagegateway_SetLocalConsolePassword(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.SetLocalConsolePasswordInput{
		// GatewayARN: *string, // Required
		// LocalConsolePassword: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayLocalConsolePassword) > 0 {
		input.LocalConsolePassword = aws.String(_storagegatewayLocalConsolePassword)
	}

	if resp, err := client.SetLocalConsolePassword(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the password for the guest user smbguest . The smbguest user is the user
// when the authentication method for the file share is set to GuestAccess . This
// operation only supported for S3 File Gateways
func storagegateway_SetSMBGuestPassword(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.SetSMBGuestPasswordInput{
		// GatewayARN: *string, // Required
		// Password: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayPassword) > 0 {
		input.Password = aws.String(_storagegatewayPassword)
	}

	if resp, err := client.SetSMBGuestPassword(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Shuts down a Tape Gateway or Volume Gateway. To specify which gateway to shut
// down, use the Amazon Resource Name (ARN) of the gateway in the body of your
// request.
//
// This API action cannot be used to shut down S3 File Gateway or FSx File Gateway.
//
// The operation shuts down the gateway service component running in the gateway's
// virtual machine (VM) and not the host VM.
//
// If you want to shut down the VM, it is recommended that you first shut down the
// gateway component in the VM to avoid unpredictable conditions.
//
// After the gateway is shutdown, you cannot call any other API except StartGateway, DescribeGatewayInformation, and ListGateways.
// For more information, see ActivateGateway. Your applications cannot read from or write to the
// gateway's storage volumes, and there are no snapshots taken.
//
// When you make a shutdown request, you will get a 200 OK success response
// immediately. However, it might take some time for the gateway to shut down. You
// can call the DescribeGatewayInformationAPI to check the status. For more information, see ActivateGateway.
//
// If do not intend to use the gateway again, you must delete the gateway (using DeleteGateway)
// to no longer pay software charges associated with the gateway.
func storagegateway_ShutdownGateway(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.ShutdownGatewayInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.ShutdownGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Start a test that verifies that the specified gateway is configured for High
// Availability monitoring in your host environment. This request only initiates
// the test and that a successful response only indicates that the test was
// started. It doesn't indicate that the test passed. For the status of the test,
// invoke the DescribeAvailabilityMonitorTest API.
//
// Starting this test will cause your gateway to go offline for a brief period.
func storagegateway_StartAvailabilityMonitorTest(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.StartAvailabilityMonitorTestInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.StartAvailabilityMonitorTest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts generating a report of the file metadata currently cached by an S3 File
// Gateway for a specific file share. You can use this report to identify and
// resolve issues if you have files failing upload from your gateway to Amazon S3.
// The report is a CSV file containing a list of files which match the set of
// filter parameters you specify in the request.
//
// The Files Failing Upload flag is reset every 24 hours and during gateway
// reboot. If this report captures the files after the reset, but before they
// become flagged again, they will not be reported as Files Failing Upload.
//
// The following requirements must be met to successfully generate a cache report:
//
// - You must have s3:PutObject and s3:AbortMultipartUpload permissions for the
// Amazon S3 bucket where you want to store the cache report.
//
// - No other cache reports can currently be in-progress for the specified file
// share.
//
// - There must be fewer than 10 existing cache reports for the specified file
// share.
//
// - The gateway must be online and connected to Amazon Web Services.
//
// - The root disk must have at least 20GB of free space when report generation
// starts.
//
// - You must specify at least one value for InclusionFilters or ExclusionFilters
// in the request.
func storagegateway_StartCacheReport(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.StartCacheReportInput{
		// BucketRegion: *string, // Required
		// ClientToken: *string, // Required
		// FileShareARN: *string, // Required
		// LocationARN: *string, // Required
		// Role: *string, // Required
	}

	if len(_storagegatewayBucketRegion) > 0 {
		input.BucketRegion = aws.String(_storagegatewayBucketRegion)
	}
	if len(_storagegatewayClientToken) > 0 {
		input.ClientToken = aws.String(_storagegatewayClientToken)
	}
	if len(_storagegatewayFileShareARN) > 0 {
		input.FileShareARN = aws.String(_storagegatewayFileShareARN)
	}
	if len(_storagegatewayLocationARN) > 0 {
		input.LocationARN = aws.String(_storagegatewayLocationARN)
	}
	if len(_storagegatewayRole) > 0 {
		input.Role = aws.String(_storagegatewayRole)
	}
	if len(_storagegatewayExclusionFilters) > 0 {
		if err := assignInputField(input, "ExclusionFilters", _storagegatewayExclusionFilters); err != nil {
			log.Errorf("invalid --exclusion-filters: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayInclusionFilters) > 0 {
		if err := assignInputField(input, "InclusionFilters", _storagegatewayInclusionFilters); err != nil {
			log.Errorf("invalid --inclusion-filters: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _storagegatewayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayVPCEndpointDNSName) > 0 {
		input.VPCEndpointDNSName = aws.String(_storagegatewayVPCEndpointDNSName)
	}

	if resp, err := client.StartCacheReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a gateway that you previously shut down (see ShutdownGateway). After the gateway
// starts, you can then make other API calls, your applications can read from or
// write to the gateway's storage volumes and you will be able to take snapshot
// backups.
//
// When you make a request, you will get a 200 OK success response immediately.
// However, it might take some time for the gateway to be ready. You should call DescribeGatewayInformation
// and check the status before making any additional API calls. For more
// information, see ActivateGateway.
//
// To specify which gateway to start, use the Amazon Resource Name (ARN) of the
// gateway in your request.
func storagegateway_StartGateway(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.StartGatewayInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.StartGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the automatic tape creation policy of a gateway. Use this to update the
// policy with a new set of automatic tape creation rules. This is only supported
// for tape gateways.
//
// By default, there is no automatic tape creation policy.
//
// A gateway can have only one automatic tape creation policy.
func storagegateway_UpdateAutomaticTapeCreationPolicy(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.UpdateAutomaticTapeCreationPolicyInput{
		// AutomaticTapeCreationRules: []types.AutomaticTapeCreationRule, // Required
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayAutomaticTapeCreationRules) > 0 {
		if err := assignInputField(input, "AutomaticTapeCreationRules", _storagegatewayAutomaticTapeCreationRules); err != nil {
			log.Errorf("invalid --automatic-tape-creation-rules: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.UpdateAutomaticTapeCreationPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the bandwidth rate limits of a gateway. You can update both the upload
// and download bandwidth rate limit or specify only one of the two. If you don't
// set a bandwidth rate limit, the existing rate limit remains. This operation is
// supported only for the stored volume, cached volume, and tape gateway types. To
// update bandwidth rate limits for S3 file gateways, use UpdateBandwidthRateLimitSchedule.
//
// By default, a gateway's bandwidth rate limits are not set. If you don't set any
// limit, the gateway does not have any limitations on its bandwidth usage and
// could potentially use the maximum available bandwidth.
//
// To specify which gateway to update, use the Amazon Resource Name (ARN) of the
// gateway in your request.
func storagegateway_UpdateBandwidthRateLimit(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.UpdateBandwidthRateLimitInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayAverageDownloadRateLimitInBitsPerSec) > 0 {
		if err := assignInputField(input, "AverageDownloadRateLimitInBitsPerSec", _storagegatewayAverageDownloadRateLimitInBitsPerSec); err != nil {
			log.Errorf("invalid --average-download-rate-limit-in-bits-per-sec: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayAverageUploadRateLimitInBitsPerSec) > 0 {
		if err := assignInputField(input, "AverageUploadRateLimitInBitsPerSec", _storagegatewayAverageUploadRateLimitInBitsPerSec); err != nil {
			log.Errorf("invalid --average-upload-rate-limit-in-bits-per-sec: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBandwidthRateLimit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the bandwidth rate limit schedule for a specified gateway. By default,
// gateways do not have bandwidth rate limit schedules, which means no bandwidth
// rate limiting is in effect. Use this to initiate or update a gateway's bandwidth
// rate limit schedule. This operation is supported for volume, tape, and S3 file
// gateways. S3 file gateways support bandwidth rate limits for upload only. FSx
// file gateways do not support bandwidth rate limits.
func storagegateway_UpdateBandwidthRateLimitSchedule(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.UpdateBandwidthRateLimitScheduleInput{
		// BandwidthRateLimitIntervals: []types.BandwidthRateLimitInterval, // Required
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayBandwidthRateLimitIntervals) > 0 {
		if err := assignInputField(input, "BandwidthRateLimitIntervals", _storagegatewayBandwidthRateLimitIntervals); err != nil {
			log.Errorf("invalid --bandwidth-rate-limit-intervals: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.UpdateBandwidthRateLimitSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the Challenge-Handshake Authentication Protocol (CHAP) credentials for
// a specified iSCSI target. By default, a gateway does not have CHAP enabled;
// however, for added security, you might use it. This operation is supported in
// the volume and tape gateway types.
//
// When you update CHAP credentials, all existing connections on the target are
// closed and initiators must reconnect with the new credentials.
func storagegateway_UpdateChapCredentials(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.UpdateChapCredentialsInput{
		// InitiatorName: *string, // Required
		// SecretToAuthenticateInitiator: *string, // Required
		// TargetARN: *string, // Required
	}

	if len(_storagegatewayInitiatorName) > 0 {
		input.InitiatorName = aws.String(_storagegatewayInitiatorName)
	}
	if len(_storagegatewaySecretToAuthenticateInitiator) > 0 {
		input.SecretToAuthenticateInitiator = aws.String(_storagegatewaySecretToAuthenticateInitiator)
	}
	if len(_storagegatewayTargetARN) > 0 {
		input.TargetARN = aws.String(_storagegatewayTargetARN)
	}
	if len(_storagegatewaySecretToAuthenticateTarget) > 0 {
		input.SecretToAuthenticateTarget = aws.String(_storagegatewaySecretToAuthenticateTarget)
	}

	if resp, err := client.UpdateChapCredentials(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a file system association. This operation is only supported in the FSx
// File Gateways.
func storagegateway_UpdateFileSystemAssociation(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.UpdateFileSystemAssociationInput{
		// FileSystemAssociationARN: *string, // Required
	}

	if len(_storagegatewayFileSystemAssociationARN) > 0 {
		input.FileSystemAssociationARN = aws.String(_storagegatewayFileSystemAssociationARN)
	}
	if len(_storagegatewayAuditDestinationARN) > 0 {
		input.AuditDestinationARN = aws.String(_storagegatewayAuditDestinationARN)
	}
	if len(_storagegatewayCacheAttributes) > 0 {
		if err := assignInputField(input, "CacheAttributes", _storagegatewayCacheAttributes); err != nil {
			log.Errorf("invalid --cache-attributes: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayPassword) > 0 {
		input.Password = aws.String(_storagegatewayPassword)
	}
	if len(_storagegatewayUserName) > 0 {
		input.UserName = aws.String(_storagegatewayUserName)
	}

	if resp, err := client.UpdateFileSystemAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a gateway's metadata, which includes the gateway's name, time zone, and
// metadata cache size. To specify which gateway to update, use the Amazon Resource
// Name (ARN) of the gateway in your request.
//
// For gateways activated after September 2, 2015, the gateway's ARN contains the
// gateway ID rather than the gateway name. However, changing the name of the
// gateway has no effect on the gateway's ARN.
func storagegateway_UpdateGatewayInformation(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.UpdateGatewayInformationInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayCloudWatchLogGroupARN) > 0 {
		input.CloudWatchLogGroupARN = aws.String(_storagegatewayCloudWatchLogGroupARN)
	}
	if len(_storagegatewayGatewayCapacity) > 0 {
		if err := assignInputField(input, "GatewayCapacity", _storagegatewayGatewayCapacity); err != nil {
			log.Errorf("invalid --gateway-capacity: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayGatewayName) > 0 {
		input.GatewayName = aws.String(_storagegatewayGatewayName)
	}
	if len(_storagegatewayGatewayTimezone) > 0 {
		input.GatewayTimezone = aws.String(_storagegatewayGatewayTimezone)
	}

	if resp, err := client.UpdateGatewayInformation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the gateway virtual machine (VM) software. The request immediately
// triggers the software update.
//
// When you make this request, you get a 200 OK success response immediately.
// However, it might take some time for the update to complete. You can call DescribeGatewayInformationto
// verify the gateway is in the STATE_RUNNING state.
//
// A software update forces a system restart of your gateway. You can minimize the
// chance of any disruption to your applications by increasing your iSCSI
// Initiators' timeouts. For more information about increasing iSCSI Initiator
// timeouts for Windows and Linux, see [Customizing your Windows iSCSI settings]and [Customizing your Linux iSCSI settings], respectively.
//
// [Customizing your Linux iSCSI settings]: https://docs.aws.amazon.com/storagegateway/latest/userguide/ConfiguringiSCSIClientInitiatorRedHatClient.html#CustomizeLinuxiSCSISettings
// [Customizing your Windows iSCSI settings]: https://docs.aws.amazon.com/storagegateway/latest/userguide/ConfiguringiSCSIClientInitiatorWindowsClient.html#CustomizeWindowsiSCSISettings
func storagegateway_UpdateGatewaySoftwareNow(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.UpdateGatewaySoftwareNowInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.UpdateGatewaySoftwareNow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a gateway's maintenance window schedule, with settings for monthly or
// weekly cadence, specific day and time to begin maintenance, and which types of
// updates to apply. Time configuration uses the gateway's time zone. You can pass
// values for a complete maintenance schedule, or update policy, or both. Previous
// values will persist for whichever setting you choose not to modify. If an
// incomplete or invalid maintenance schedule is passed, the entire request will be
// rejected with an error and no changes will occur.
//
// A complete maintenance schedule must include values for both MinuteOfHour and
// HourOfDay , and either DayOfMonth or DayOfWeek .
//
// We recommend keeping maintenance updates turned on, except in specific use
// cases where the brief disruptions caused by updating the gateway could
// critically impact your deployment.
func storagegateway_UpdateMaintenanceStartTime(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.UpdateMaintenanceStartTimeInput{
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewayDayOfMonth) > 0 {
		if err := assignInputField(input, "DayOfMonth", _storagegatewayDayOfMonth); err != nil {
			log.Errorf("invalid --day-of-month: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayDayOfWeek) > 0 {
		if err := assignInputField(input, "DayOfWeek", _storagegatewayDayOfWeek); err != nil {
			log.Errorf("invalid --day-of-week: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayHourOfDay) > 0 {
		if err := assignInputField(input, "HourOfDay", _storagegatewayHourOfDay); err != nil {
			log.Errorf("invalid --hour-of-day: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayMinuteOfHour) > 0 {
		if err := assignInputField(input, "MinuteOfHour", _storagegatewayMinuteOfHour); err != nil {
			log.Errorf("invalid --minute-of-hour: %s", err.Error())
			return
		}
	}
	if len(_storagegatewaySoftwareUpdatePreferences) > 0 {
		if err := assignInputField(input, "SoftwareUpdatePreferences", _storagegatewaySoftwareUpdatePreferences); err != nil {
			log.Errorf("invalid --software-update-preferences: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMaintenanceStartTime(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Network File System (NFS) file share. This operation is only
// supported in S3 File Gateways.
//
// To leave a file share field unchanged, set the corresponding input field to
// null.
//
// Updates the following file share settings:
//
// - Default storage class for your S3 bucket
//
// - Metadata defaults for your S3 bucket
//
// - Allowed NFS clients for your file share
//
// - Squash settings
//
// - Write status of your file share
func storagegateway_UpdateNFSFileShare(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.UpdateNFSFileShareInput{
		// FileShareARN: *string, // Required
	}

	if len(_storagegatewayFileShareARN) > 0 {
		input.FileShareARN = aws.String(_storagegatewayFileShareARN)
	}
	if len(_storagegatewayAuditDestinationARN) > 0 {
		input.AuditDestinationARN = aws.String(_storagegatewayAuditDestinationARN)
	}
	if len(_storagegatewayCacheAttributes) > 0 {
		if err := assignInputField(input, "CacheAttributes", _storagegatewayCacheAttributes); err != nil {
			log.Errorf("invalid --cache-attributes: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayClientList) > 0 {
		input.ClientList = append([]string(nil), _storagegatewayClientList...)
	}
	if len(_storagegatewayDefaultStorageClass) > 0 {
		input.DefaultStorageClass = aws.String(_storagegatewayDefaultStorageClass)
	}
	if len(_storagegatewayEncryptionType) > 0 {
		if err := assignInputField(input, "EncryptionType", _storagegatewayEncryptionType); err != nil {
			log.Errorf("invalid --encryption-type: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayFileShareName) > 0 {
		input.FileShareName = aws.String(_storagegatewayFileShareName)
	}
	if len(_storagegatewayGuessMIMETypeEnabled) > 0 {
		if err := assignInputField(input, "GuessMIMETypeEnabled", _storagegatewayGuessMIMETypeEnabled); err != nil {
			log.Errorf("invalid --guess-mime-type-enabled: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayKMSEncrypted) > 0 {
		if err := assignInputField(input, "KMSEncrypted", _storagegatewayKMSEncrypted); err != nil {
			log.Errorf("invalid --kms-encrypted: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayKMSKey) > 0 {
		input.KMSKey = aws.String(_storagegatewayKMSKey)
	}
	if len(_storagegatewayNFSFileShareDefaults) > 0 {
		if err := assignInputField(input, "NFSFileShareDefaults", _storagegatewayNFSFileShareDefaults); err != nil {
			log.Errorf("invalid --nfs-file-share-defaults: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayNotificationPolicy) > 0 {
		input.NotificationPolicy = aws.String(_storagegatewayNotificationPolicy)
	}
	if len(_storagegatewayObjectACL) > 0 {
		if err := assignInputField(input, "ObjectACL", _storagegatewayObjectACL); err != nil {
			log.Errorf("invalid --object-acl: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayReadOnly) > 0 {
		if err := assignInputField(input, "ReadOnly", _storagegatewayReadOnly); err != nil {
			log.Errorf("invalid --read-only: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayRequesterPays) > 0 {
		if err := assignInputField(input, "RequesterPays", _storagegatewayRequesterPays); err != nil {
			log.Errorf("invalid --requester-pays: %s", err.Error())
			return
		}
	}
	if len(_storagegatewaySquash) > 0 {
		input.Squash = aws.String(_storagegatewaySquash)
	}

	if resp, err := client.UpdateNFSFileShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Server Message Block (SMB) file share. This operation is only
// supported for S3 File Gateways.
//
// To leave a file share field unchanged, set the corresponding input field to
// null.
//
// File gateways require Security Token Service (Amazon Web Services STS) to be
// activated to enable you to create a file share. Make sure that Amazon Web
// Services STS is activated in the Amazon Web Services Region you are creating
// your file gateway in. If Amazon Web Services STS is not activated in this Amazon
// Web Services Region, activate it. For information about how to activate Amazon
// Web Services STS, see [Activating and deactivating Amazon Web Services STS in an Amazon Web Services Region]in the Identity and Access Management User Guide.
//
// File gateways don't support creating hard or symbolic links on a file share.
//
// [Activating and deactivating Amazon Web Services STS in an Amazon Web Services Region]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html
func storagegateway_UpdateSMBFileShare(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.UpdateSMBFileShareInput{
		// FileShareARN: *string, // Required
	}

	if len(_storagegatewayFileShareARN) > 0 {
		input.FileShareARN = aws.String(_storagegatewayFileShareARN)
	}
	if len(_storagegatewayAccessBasedEnumeration) > 0 {
		if err := assignInputField(input, "AccessBasedEnumeration", _storagegatewayAccessBasedEnumeration); err != nil {
			log.Errorf("invalid --access-based-enumeration: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayAdminUserList) > 0 {
		input.AdminUserList = append([]string(nil), _storagegatewayAdminUserList...)
	}
	if len(_storagegatewayAuditDestinationARN) > 0 {
		input.AuditDestinationARN = aws.String(_storagegatewayAuditDestinationARN)
	}
	if len(_storagegatewayCacheAttributes) > 0 {
		if err := assignInputField(input, "CacheAttributes", _storagegatewayCacheAttributes); err != nil {
			log.Errorf("invalid --cache-attributes: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayCaseSensitivity) > 0 {
		if err := assignInputField(input, "CaseSensitivity", _storagegatewayCaseSensitivity); err != nil {
			log.Errorf("invalid --case-sensitivity: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayDefaultStorageClass) > 0 {
		input.DefaultStorageClass = aws.String(_storagegatewayDefaultStorageClass)
	}
	if len(_storagegatewayEncryptionType) > 0 {
		if err := assignInputField(input, "EncryptionType", _storagegatewayEncryptionType); err != nil {
			log.Errorf("invalid --encryption-type: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayFileShareName) > 0 {
		input.FileShareName = aws.String(_storagegatewayFileShareName)
	}
	if len(_storagegatewayGuessMIMETypeEnabled) > 0 {
		if err := assignInputField(input, "GuessMIMETypeEnabled", _storagegatewayGuessMIMETypeEnabled); err != nil {
			log.Errorf("invalid --guess-mime-type-enabled: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayInvalidUserList) > 0 {
		input.InvalidUserList = append([]string(nil), _storagegatewayInvalidUserList...)
	}
	if len(_storagegatewayKMSEncrypted) > 0 {
		if err := assignInputField(input, "KMSEncrypted", _storagegatewayKMSEncrypted); err != nil {
			log.Errorf("invalid --kms-encrypted: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayKMSKey) > 0 {
		input.KMSKey = aws.String(_storagegatewayKMSKey)
	}
	if len(_storagegatewayNotificationPolicy) > 0 {
		input.NotificationPolicy = aws.String(_storagegatewayNotificationPolicy)
	}
	if len(_storagegatewayObjectACL) > 0 {
		if err := assignInputField(input, "ObjectACL", _storagegatewayObjectACL); err != nil {
			log.Errorf("invalid --object-acl: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayOplocksEnabled) > 0 {
		if err := assignInputField(input, "OplocksEnabled", _storagegatewayOplocksEnabled); err != nil {
			log.Errorf("invalid --oplocks-enabled: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayReadOnly) > 0 {
		if err := assignInputField(input, "ReadOnly", _storagegatewayReadOnly); err != nil {
			log.Errorf("invalid --read-only: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayRequesterPays) > 0 {
		if err := assignInputField(input, "RequesterPays", _storagegatewayRequesterPays); err != nil {
			log.Errorf("invalid --requester-pays: %s", err.Error())
			return
		}
	}
	if len(_storagegatewaySMBACLEnabled) > 0 {
		if err := assignInputField(input, "SMBACLEnabled", _storagegatewaySMBACLEnabled); err != nil {
			log.Errorf("invalid --smbacl-enabled: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayValidUserList) > 0 {
		input.ValidUserList = append([]string(nil), _storagegatewayValidUserList...)
	}

	if resp, err := client.UpdateSMBFileShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Controls whether the shares on an S3 File Gateway are visible in a net view or
// browse list. The operation is only supported for S3 File Gateways.
func storagegateway_UpdateSMBFileShareVisibility(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.UpdateSMBFileShareVisibilityInput{
		// FileSharesVisible: *bool, // Required
		// GatewayARN: *string, // Required
	}

	if len(_storagegatewayFileSharesVisible) > 0 {
		if err := assignInputField(input, "FileSharesVisible", _storagegatewayFileSharesVisible); err != nil {
			log.Errorf("invalid --file-shares-visible: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}

	if resp, err := client.UpdateSMBFileShareVisibility(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the list of Active Directory users and groups that have special
// permissions for SMB file shares on the gateway.
func storagegateway_UpdateSMBLocalGroups(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.UpdateSMBLocalGroupsInput{
		// GatewayARN: *string, // Required
		// SMBLocalGroups: *types.SMBLocalGroups, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewaySMBLocalGroups) > 0 {
		if err := assignInputField(input, "SMBLocalGroups", _storagegatewaySMBLocalGroups); err != nil {
			log.Errorf("invalid --smb-local-groups: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSMBLocalGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the SMB security strategy level for an Amazon S3 file gateway. This
// action is only supported for Amazon S3 file gateways.
//
// For information about configuring this setting using the Amazon Web Services
// console, see [Setting a security level for your gateway]in the Amazon S3 File Gateway User Guide.
//
// A higher security strategy level can affect performance of the gateway.
//
// [Setting a security level for your gateway]: https://docs.aws.amazon.com/filegateway/latest/files3/security-strategy.html
func storagegateway_UpdateSMBSecurityStrategy(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.UpdateSMBSecurityStrategyInput{
		// GatewayARN: *string, // Required
		// SMBSecurityStrategy: types.SMBSecurityStrategy, // Required
	}

	if len(_storagegatewayGatewayARN) > 0 {
		input.GatewayARN = aws.String(_storagegatewayGatewayARN)
	}
	if len(_storagegatewaySMBSecurityStrategy) > 0 {
		if err := assignInputField(input, "SMBSecurityStrategy", _storagegatewaySMBSecurityStrategy); err != nil {
			log.Errorf("invalid --smb-security-strategy: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSMBSecurityStrategy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a snapshot schedule configured for a gateway volume. This operation is
// only supported in the cached volume and stored volume gateway types.
//
// The default snapshot schedule for volume is once every 24 hours, starting at
// the creation time of the volume. You can use this API to change the snapshot
// schedule configured for the volume.
//
// In the request you must identify the gateway volume whose snapshot schedule you
// want to update, and the schedule information, including when you want the
// snapshot to begin on a day and the frequency (in hours) of snapshots.
func storagegateway_UpdateSnapshotSchedule(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.UpdateSnapshotScheduleInput{
		// RecurrenceInHours: *int32, // Required
		// StartAt: *int32, // Required
		// VolumeARN: *string, // Required
	}

	if len(_storagegatewayRecurrenceInHours) > 0 {
		if err := assignInputField(input, "RecurrenceInHours", _storagegatewayRecurrenceInHours); err != nil {
			log.Errorf("invalid --recurrence-in-hours: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayStartAt) > 0 {
		if err := assignInputField(input, "StartAt", _storagegatewayStartAt); err != nil {
			log.Errorf("invalid --start-at: %s", err.Error())
			return
		}
	}
	if len(_storagegatewayVolumeARN) > 0 {
		input.VolumeARN = aws.String(_storagegatewayVolumeARN)
	}
	if len(_storagegatewayDescription) > 0 {
		input.Description = aws.String(_storagegatewayDescription)
	}
	if len(_storagegatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _storagegatewayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSnapshotSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the type of medium changer in a tape gateway. When you activate a tape
// gateway, you select a medium changer type for the tape gateway. This operation
// enables you to select a different type of medium changer after a tape gateway is
// activated. This operation is only supported in the tape gateway type.
func storagegateway_UpdateVTLDeviceType(cfg aws.Config, client *storagegateway.Client) {
	input := &storagegateway.UpdateVTLDeviceTypeInput{
		// DeviceType: *string, // Required
		// VTLDeviceARN: *string, // Required
	}

	if len(_storagegatewayDeviceType) > 0 {
		input.DeviceType = aws.String(_storagegatewayDeviceType)
	}
	if len(_storagegatewayVTLDeviceARN) > 0 {
		input.VTLDeviceARN = aws.String(_storagegatewayVTLDeviceARN)
	}

	if resp, err := client.UpdateVTLDeviceType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_storagegatewayCmd)
	_storagegatewayCmd.Flags().SortFlags = false

	_storagegatewayCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_storagegatewayCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_storagegatewayCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayAccessBasedEnumeration, "access-based-enumeration", "", "", "Access Based Enumeration")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayActivationKey, "activation-key", "", "", "Activation Key")
	_storagegatewayCmd.Flags().StringSliceVarP(&_storagegatewayAdminUserList, "admin-user-list", "", nil, "Admin User List")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayAuditDestinationARN, "audit-destination-arn", "", "", "Audit Destination ARN")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayAuthentication, "authentication", "", "", "Authentication")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayAutomaticTapeCreationRules, "automatic-tape-creation-rules", "", "", "Automatic Tape Creation Rules")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayAverageDownloadRateLimitInBitsPerSec, "average-download-rate-limit-in-bits-per-sec", "", "", "Average Download Rate Limit In Bits Per Sec")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayAverageUploadRateLimitInBitsPerSec, "average-upload-rate-limit-in-bits-per-sec", "", "", "Average Upload Rate Limit In Bits Per Sec")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayBandwidthRateLimitIntervals, "bandwidth-rate-limit-intervals", "", "", "Bandwidth Rate Limit Intervals")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayBandwidthType, "bandwidth-type", "", "", "Bandwidth Type")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayBucketRegion, "bucket-region", "", "", "Bucket Region")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayBypassGovernanceRetention, "bypass-governance-retention", "", "", "Bypass Governance Retention")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayCacheAttributes, "cache-attributes", "", "", "Cache Attributes")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayCacheReportARN, "cache-report-arn", "", "", "Cache Report ARN")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayCaseSensitivity, "case-sensitivity", "", "", "Case Sensitivity")
	_storagegatewayCmd.Flags().StringSliceVarP(&_storagegatewayClientList, "client-list", "", nil, "Client List")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayClientToken, "client-token", "", "", "Client Token")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayCloudWatchLogGroupARN, "cloud-watch-log-group-arn", "", "", "Cloud Watch Log Group ARN")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayDayOfMonth, "day-of-month", "", "", "Day Of Month")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayDayOfWeek, "day-of-week", "", "", "Day Of Week")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayDefaultStorageClass, "default-storage-class", "", "", "Default Storage Class")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayDescription, "description", "", "", "Description")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayDeviceType, "device-type", "", "", "Device Type")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayDiskId, "disk-id", "", "", "Disk ID")
	_storagegatewayCmd.Flags().StringSliceVarP(&_storagegatewayDiskIds, "disk-ids", "", nil, "Disk Ids")
	_storagegatewayCmd.Flags().StringSliceVarP(&_storagegatewayDomainControllers, "domain-controllers", "", nil, "Domain Controllers")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayDomainName, "domain-name", "", "", "Domain Name")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayEncryptionType, "encryption-type", "", "", "Encryption Type")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayEndpointNetworkConfiguration, "endpoint-network-configuration", "", "", "Endpoint Network Configuration")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayExclusionFilters, "exclusion-filters", "", "", "Exclusion Filters")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayFileShareARN, "file-share-arn", "", "", "File Share ARN")
	_storagegatewayCmd.Flags().StringSliceVarP(&_storagegatewayFileShareARNList, "file-share-arn-list", "", nil, "File Share ARN List")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayFileShareName, "file-share-name", "", "", "File Share Name")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayFileSharesVisible, "file-shares-visible", "", "", "File Shares Visible")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayFileSystemAssociationARN, "file-system-association-arn", "", "", "File System Association ARN")
	_storagegatewayCmd.Flags().StringSliceVarP(&_storagegatewayFileSystemAssociationARNList, "file-system-association-arn-list", "", nil, "File System Association ARN List")
	_storagegatewayCmd.Flags().StringSliceVarP(&_storagegatewayFolderList, "folder-list", "", nil, "Folder List")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayForceDelete, "force-delete", "", "", "Force Delete")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayForceDetach, "force-detach", "", "", "Force Detach")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayForceRemove, "force-remove", "", "", "Force Remove")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayGatewayARN, "gateway-arn", "", "", "Gateway ARN")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayGatewayCapacity, "gateway-capacity", "", "", "Gateway Capacity")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayGatewayName, "gateway-name", "", "", "Gateway Name")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayGatewayRegion, "gateway-region", "", "", "Gateway Region")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayGatewayTimezone, "gateway-timezone", "", "", "Gateway Timezone")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayGatewayType, "gateway-type", "", "", "Gateway Type")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayGuessMIMETypeEnabled, "guess-mime-type-enabled", "", "", "Guess Mime Type Enabled")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayHourOfDay, "hour-of-day", "", "", "Hour Of Day")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayInclusionFilters, "inclusion-filters", "", "", "Inclusion Filters")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayInitiatorName, "initiator-name", "", "", "Initiator Name")
	_storagegatewayCmd.Flags().StringSliceVarP(&_storagegatewayInvalidUserList, "invalid-user-list", "", nil, "Invalid User List")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayKMSEncrypted, "kms-encrypted", "", "", "KMS Encrypted")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayKMSKey, "kms-key", "", "", "KMS Key")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayLimit, "limit", "", "", "Limit")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayLocalConsolePassword, "local-console-password", "", "", "Local Console Password")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayLocationARN, "location-arn", "", "", "Location ARN")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayMarker, "marker", "", "", "Marker")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayMediumChangerType, "medium-changer-type", "", "", "Medium Changer Type")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayMinuteOfHour, "minute-of-hour", "", "", "Minute Of Hour")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayNetworkInterfaceId, "network-interface-id", "", "", "Network Interface ID")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayNFSFileShareDefaults, "nfs-file-share-defaults", "", "", "Nfs File Share Defaults")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayNotificationPolicy, "notification-policy", "", "", "Notification Policy")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayNumTapesToCreate, "num-tapes-to-create", "", "", "Num Tapes To Create")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayObjectACL, "object-acl", "", "", "Object ACL")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayOplocksEnabled, "oplocks-enabled", "", "", "Oplocks Enabled")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayOrganizationalUnit, "organizational-unit", "", "", "Organizational Unit")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayPassword, "password", "", "", "Password")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayPoolARN, "pool-arn", "", "", "Pool ARN")
	_storagegatewayCmd.Flags().StringSliceVarP(&_storagegatewayPoolARNs, "pool-arns", "", nil, "Pool Arns")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayPoolId, "pool-id", "", "", "Pool ID")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayPoolName, "pool-name", "", "", "Pool Name")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayPreserveExistingData, "preserve-existing-data", "", "", "Preserve Existing Data")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayReadOnly, "read-only", "", "", "Read Only")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayRecurrenceInHours, "recurrence-in-hours", "", "", "Recurrence In Hours")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayRecursive, "recursive", "", "", "Recursive")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayRequesterPays, "requester-pays", "", "", "Requester Pays")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayResourceARN, "resource-arn", "", "", "Resource ARN")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayRetentionLockTimeInDays, "retention-lock-time-in-days", "", "", "Retention Lock Time In Days")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayRetentionLockType, "retention-lock-type", "", "", "Retention Lock Type")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayRole, "role", "", "", "Role")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewaySecretToAuthenticateInitiator, "secret-to-authenticate-initiator", "", "", "Secret To Authenticate Initiator")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewaySecretToAuthenticateTarget, "secret-to-authenticate-target", "", "", "Secret To Authenticate Target")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewaySMBLocalGroups, "smb-local-groups", "", "", "Smb Local Groups")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewaySMBSecurityStrategy, "smb-security-strategy", "", "", "Smb Security Strategy")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewaySMBACLEnabled, "smbacl-enabled", "", "", "Smbacl Enabled")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewaySnapshotDescription, "snapshot-description", "", "", "Snapshot Description")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewaySnapshotId, "snapshot-id", "", "", "Snapshot ID")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewaySoftwareUpdatePreferences, "software-update-preferences", "", "", "Software Update Preferences")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewaySourceVolumeARN, "source-volume-arn", "", "", "Source Volume ARN")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewaySquash, "squash", "", "", "Squash")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayStartAt, "start-at", "", "", "Start At")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayStorageClass, "storage-class", "", "", "Storage Class")
	_storagegatewayCmd.Flags().StringSliceVarP(&_storagegatewayTagKeys, "tag-keys", "", nil, "Tag Keys")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayTags, "tags", "", "", "Tags")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayTapeARN, "tape-arn", "", "", "Tape ARN")
	_storagegatewayCmd.Flags().StringSliceVarP(&_storagegatewayTapeARNs, "tape-arns", "", nil, "Tape Arns")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayTapeBarcode, "tape-barcode", "", "", "Tape Barcode")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayTapeBarcodePrefix, "tape-barcode-prefix", "", "", "Tape Barcode Prefix")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayTapeDriveType, "tape-drive-type", "", "", "Tape Drive Type")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayTapeSizeInBytes, "tape-size-in-bytes", "", "", "Tape Size In Bytes")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayTargetARN, "target-arn", "", "", "Target ARN")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayTargetName, "target-name", "", "", "Target Name")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayTimeoutInSeconds, "timeout-in-seconds", "", "", "Timeout In Seconds")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayUserName, "user-name", "", "", "User Name")
	_storagegatewayCmd.Flags().StringSliceVarP(&_storagegatewayValidUserList, "valid-user-list", "", nil, "Valid User List")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayVolumeARN, "volume-arn", "", "", "Volume ARN")
	_storagegatewayCmd.Flags().StringSliceVarP(&_storagegatewayVolumeARNs, "volume-arns", "", nil, "Volume Arns")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayVolumeSizeInBytes, "volume-size-in-bytes", "", "", "Volume Size In Bytes")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayVPCEndpointDNSName, "vpc-endpoint-dns-name", "", "", "VPC Endpoint DNS Name")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayVTLDeviceARN, "vtl-device-arn", "", "", "VTL Device ARN")
	_storagegatewayCmd.Flags().StringSliceVarP(&_storagegatewayVTLDeviceARNs, "vtl-device-arns", "", nil, "VTL Device Arns")
	_storagegatewayCmd.Flags().StringVarP(&_storagegatewayWorm, "worm", "", "", "Worm")

	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayActivateGateway, "activate-gateway", "", false, "Activate Gateway")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayAddCache, "add-cache", "", false, "Add Cache")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayAddTagsToResource, "add-tags-to-resource", "", false, "Add Tags To Resource")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayAddUploadBuffer, "add-upload-buffer", "", false, "Add Upload Buffer")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayAddWorkingStorage, "add-working-storage", "", false, "Add Working Storage")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayAssignTapePool, "assign-tape-pool", "", false, "Assign Tape Pool")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayAssociateFileSystem, "associate-file-system", "", false, "Associate File System")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayAttachVolume, "attach-volume", "", false, "Attach Volume")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayCancelArchival, "cancel-archival", "", false, "Cancel Archival")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayCancelCacheReport, "cancel-cache-report", "", false, "Cancel Cache Report")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayCancelRetrieval, "cancel-retrieval", "", false, "Cancel Retrieval")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayCreateCachediSCSIVolume, "create-cachedi-scsi-volume", "", false, "Create Cachedi Scsi Volume")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayCreateNFSFileShare, "create-nfs-file-share", "", false, "Create Nfs File Share")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayCreateSMBFileShare, "create-smb-file-share", "", false, "Create Smb File Share")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayCreateSnapshot, "create-snapshot", "", false, "Create Snapshot")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayCreateSnapshotFromVolumeRecoveryPoint, "create-snapshot-from-volume-recovery-point", "", false, "Create Snapshot From Volume Recovery Point")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayCreateStorediSCSIVolume, "create-storedi-scsi-volume", "", false, "Create Storedi Scsi Volume")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayCreateTapePool, "create-tape-pool", "", false, "Create Tape Pool")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayCreateTapeWithBarcode, "create-tape-with-barcode", "", false, "Create Tape With Barcode")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayCreateTapes, "create-tapes", "", false, "Create Tapes")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDeleteAutomaticTapeCreationPolicy, "delete-automatic-tape-creation-policy", "", false, "Delete Automatic Tape Creation Policy")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDeleteBandwidthRateLimit, "delete-bandwidth-rate-limit", "", false, "Delete Bandwidth Rate Limit")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDeleteCacheReport, "delete-cache-report", "", false, "Delete Cache Report")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDeleteChapCredentials, "delete-chap-credentials", "", false, "Delete Chap Credentials")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDeleteFileShare, "delete-file-share", "", false, "Delete File Share")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDeleteGateway, "delete-gateway", "", false, "Delete Gateway")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDeleteSnapshotSchedule, "delete-snapshot-schedule", "", false, "Delete Snapshot Schedule")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDeleteTape, "delete-tape", "", false, "Delete Tape")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDeleteTapeArchive, "delete-tape-archive", "", false, "Delete Tape Archive")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDeleteTapePool, "delete-tape-pool", "", false, "Delete Tape Pool")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDeleteVolume, "delete-volume", "", false, "Delete Volume")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDescribeAvailabilityMonitorTest, "describe-availability-monitor-test", "", false, "Describe Availability Monitor Test")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDescribeBandwidthRateLimit, "describe-bandwidth-rate-limit", "", false, "Describe Bandwidth Rate Limit")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDescribeBandwidthRateLimitSchedule, "describe-bandwidth-rate-limit-schedule", "", false, "Describe Bandwidth Rate Limit Schedule")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDescribeCache, "describe-cache", "", false, "Describe Cache")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDescribeCacheReport, "describe-cache-report", "", false, "Describe Cache Report")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDescribeCachediSCSIVolumes, "describe-cachedi-scsi-volumes", "", false, "Describe Cachedi Scsi Volumes")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDescribeChapCredentials, "describe-chap-credentials", "", false, "Describe Chap Credentials")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDescribeFileSystemAssociations, "describe-file-system-associations", "", false, "Describe File System Associations")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDescribeGatewayInformation, "describe-gateway-information", "", false, "Describe Gateway Information")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDescribeMaintenanceStartTime, "describe-maintenance-start-time", "", false, "Describe Maintenance Start Time")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDescribeNFSFileShares, "describe-nfs-file-shares", "", false, "Describe Nfs File Shares")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDescribeSMBFileShares, "describe-smb-file-shares", "", false, "Describe Smb File Shares")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDescribeSMBSettings, "describe-smb-settings", "", false, "Describe Smb Settings")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDescribeSnapshotSchedule, "describe-snapshot-schedule", "", false, "Describe Snapshot Schedule")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDescribeStorediSCSIVolumes, "describe-storedi-scsi-volumes", "", false, "Describe Storedi Scsi Volumes")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDescribeTapeArchives, "describe-tape-archives", "", false, "Describe Tape Archives")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDescribeTapeRecoveryPoints, "describe-tape-recovery-points", "", false, "Describe Tape Recovery Points")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDescribeTapes, "describe-tapes", "", false, "Describe Tapes")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDescribeUploadBuffer, "describe-upload-buffer", "", false, "Describe Upload Buffer")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDescribeVTLDevices, "describe-vtl-devices", "", false, "Describe VTL Devices")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDescribeWorkingStorage, "describe-working-storage", "", false, "Describe Working Storage")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDetachVolume, "detach-volume", "", false, "Detach Volume")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDisableGateway, "disable-gateway", "", false, "Disable Gateway")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayDisassociateFileSystem, "disassociate-file-system", "", false, "Disassociate File System")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayEvictFilesFailingUpload, "evict-files-failing-upload", "", false, "Evict Files Failing Upload")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayJoinDomain, "join-domain", "", false, "Join Domain")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayListAutomaticTapeCreationPolicies, "list-automatic-tape-creation-policies", "", false, "List Automatic Tape Creation Policies")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayListCacheReports, "list-cache-reports", "", false, "List Cache Reports")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayListFileShares, "list-file-shares", "", false, "List File Shares")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayListFileSystemAssociations, "list-file-system-associations", "", false, "List File System Associations")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayListGateways, "list-gateways", "", false, "List Gateways")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayListLocalDisks, "list-local-disks", "", false, "List Local Disks")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayListTapePools, "list-tape-pools", "", false, "List Tape Pools")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayListTapes, "list-tapes", "", false, "List Tapes")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayListVolumeInitiators, "list-volume-initiators", "", false, "List Volume Initiators")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayListVolumeRecoveryPoints, "list-volume-recovery-points", "", false, "List Volume Recovery Points")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayListVolumes, "list-volumes", "", false, "List Volumes")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayNotifyWhenUploaded, "notify-when-uploaded", "", false, "Notify When Uploaded")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayRefreshCache, "refresh-cache", "", false, "Refresh Cache")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayRemoveTagsFromResource, "remove-tags-from-resource", "", false, "Remove Tags From Resource")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayResetCache, "reset-cache", "", false, "Reset Cache")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayRetrieveTapeArchive, "retrieve-tape-archive", "", false, "Retrieve Tape Archive")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayRetrieveTapeRecoveryPoint, "retrieve-tape-recovery-point", "", false, "Retrieve Tape Recovery Point")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewaySetLocalConsolePassword, "set-local-console-password", "", false, "Set Local Console Password")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewaySetSMBGuestPassword, "set-smb-guest-password", "", false, "Set Smb Guest Password")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayShutdownGateway, "shutdown-gateway", "", false, "Shutdown Gateway")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayStartAvailabilityMonitorTest, "start-availability-monitor-test", "", false, "Start Availability Monitor Test")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayStartCacheReport, "start-cache-report", "", false, "Start Cache Report")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayStartGateway, "start-gateway", "", false, "Start Gateway")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayUpdateAutomaticTapeCreationPolicy, "update-automatic-tape-creation-policy", "", false, "Update Automatic Tape Creation Policy")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayUpdateBandwidthRateLimit, "update-bandwidth-rate-limit", "", false, "Update Bandwidth Rate Limit")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayUpdateBandwidthRateLimitSchedule, "update-bandwidth-rate-limit-schedule", "", false, "Update Bandwidth Rate Limit Schedule")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayUpdateChapCredentials, "update-chap-credentials", "", false, "Update Chap Credentials")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayUpdateFileSystemAssociation, "update-file-system-association", "", false, "Update File System Association")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayUpdateGatewayInformation, "update-gateway-information", "", false, "Update Gateway Information")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayUpdateGatewaySoftwareNow, "update-gateway-software-now", "", false, "Update Gateway Software Now")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayUpdateMaintenanceStartTime, "update-maintenance-start-time", "", false, "Update Maintenance Start Time")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayUpdateNFSFileShare, "update-nfs-file-share", "", false, "Update Nfs File Share")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayUpdateSMBFileShare, "update-smb-file-share", "", false, "Update Smb File Share")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayUpdateSMBFileShareVisibility, "update-smb-file-share-visibility", "", false, "Update Smb File Share Visibility")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayUpdateSMBLocalGroups, "update-smb-local-groups", "", false, "Update Smb Local Groups")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayUpdateSMBSecurityStrategy, "update-smb-security-strategy", "", false, "Update Smb Security Strategy")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayUpdateSnapshotSchedule, "update-snapshot-schedule", "", false, "Update Snapshot Schedule")
	_storagegatewayCmd.Flags().BoolVarP(&_storagegatewayUpdateVTLDeviceType, "update-vtl-device-type", "", false, "Update VTL Device Type")

}
