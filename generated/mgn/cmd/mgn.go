package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mgn"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// mgnCmd represents the mgn command
var _mgnCmd = &cobra.Command{
	Use:   "mgn",
	Short: "AWS mgn CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := mgn.NewFromConfig(cfg)
		if _mgnArchiveApplication {
			mgn_ArchiveApplication(cfg, client)
			return
		}
		if _mgnArchiveWave {
			mgn_ArchiveWave(cfg, client)
			return
		}
		if _mgnAssociateApplications {
			mgn_AssociateApplications(cfg, client)
			return
		}
		if _mgnAssociateSourceServers {
			mgn_AssociateSourceServers(cfg, client)
			return
		}
		if _mgnChangeServerLifeCycleState {
			mgn_ChangeServerLifeCycleState(cfg, client)
			return
		}
		if _mgnCreateApplication {
			mgn_CreateApplication(cfg, client)
			return
		}
		if _mgnCreateConnector {
			mgn_CreateConnector(cfg, client)
			return
		}
		if _mgnCreateLaunchConfigurationTemplate {
			mgn_CreateLaunchConfigurationTemplate(cfg, client)
			return
		}
		if _mgnCreateReplicationConfigurationTemplate {
			mgn_CreateReplicationConfigurationTemplate(cfg, client)
			return
		}
		if _mgnCreateWave {
			mgn_CreateWave(cfg, client)
			return
		}
		if _mgnDeleteApplication {
			mgn_DeleteApplication(cfg, client)
			return
		}
		if _mgnDeleteConnector {
			mgn_DeleteConnector(cfg, client)
			return
		}
		if _mgnDeleteJob {
			mgn_DeleteJob(cfg, client)
			return
		}
		if _mgnDeleteLaunchConfigurationTemplate {
			mgn_DeleteLaunchConfigurationTemplate(cfg, client)
			return
		}
		if _mgnDeleteReplicationConfigurationTemplate {
			mgn_DeleteReplicationConfigurationTemplate(cfg, client)
			return
		}
		if _mgnDeleteSourceServer {
			mgn_DeleteSourceServer(cfg, client)
			return
		}
		if _mgnDeleteVcenterClient {
			mgn_DeleteVcenterClient(cfg, client)
			return
		}
		if _mgnDeleteWave {
			mgn_DeleteWave(cfg, client)
			return
		}
		if _mgnDescribeJobLogItems {
			mgn_DescribeJobLogItems(cfg, client)
			return
		}
		if _mgnDescribeJobs {
			mgn_DescribeJobs(cfg, client)
			return
		}
		if _mgnDescribeLaunchConfigurationTemplates {
			mgn_DescribeLaunchConfigurationTemplates(cfg, client)
			return
		}
		if _mgnDescribeReplicationConfigurationTemplates {
			mgn_DescribeReplicationConfigurationTemplates(cfg, client)
			return
		}
		if _mgnDescribeSourceServers {
			mgn_DescribeSourceServers(cfg, client)
			return
		}
		if _mgnDescribeVcenterClients {
			mgn_DescribeVcenterClients(cfg, client)
			return
		}
		if _mgnDisassociateApplications {
			mgn_DisassociateApplications(cfg, client)
			return
		}
		if _mgnDisassociateSourceServers {
			mgn_DisassociateSourceServers(cfg, client)
			return
		}
		if _mgnDisconnectFromService {
			mgn_DisconnectFromService(cfg, client)
			return
		}
		if _mgnFinalizeCutover {
			mgn_FinalizeCutover(cfg, client)
			return
		}
		if _mgnGetLaunchConfiguration {
			mgn_GetLaunchConfiguration(cfg, client)
			return
		}
		if _mgnGetReplicationConfiguration {
			mgn_GetReplicationConfiguration(cfg, client)
			return
		}
		if _mgnInitializeService {
			mgn_InitializeService(cfg, client)
			return
		}
		if _mgnListApplications {
			mgn_ListApplications(cfg, client)
			return
		}
		if _mgnListConnectors {
			mgn_ListConnectors(cfg, client)
			return
		}
		if _mgnListExportErrors {
			mgn_ListExportErrors(cfg, client)
			return
		}
		if _mgnListExports {
			mgn_ListExports(cfg, client)
			return
		}
		if _mgnListImportErrors {
			mgn_ListImportErrors(cfg, client)
			return
		}
		if _mgnListImports {
			mgn_ListImports(cfg, client)
			return
		}
		if _mgnListManagedAccounts {
			mgn_ListManagedAccounts(cfg, client)
			return
		}
		if _mgnListSourceServerActions {
			mgn_ListSourceServerActions(cfg, client)
			return
		}
		if _mgnListTagsForResource {
			mgn_ListTagsForResource(cfg, client)
			return
		}
		if _mgnListTemplateActions {
			mgn_ListTemplateActions(cfg, client)
			return
		}
		if _mgnListWaves {
			mgn_ListWaves(cfg, client)
			return
		}
		if _mgnMarkAsArchived {
			mgn_MarkAsArchived(cfg, client)
			return
		}
		if _mgnPauseReplication {
			mgn_PauseReplication(cfg, client)
			return
		}
		if _mgnPutSourceServerAction {
			mgn_PutSourceServerAction(cfg, client)
			return
		}
		if _mgnPutTemplateAction {
			mgn_PutTemplateAction(cfg, client)
			return
		}
		if _mgnRemoveSourceServerAction {
			mgn_RemoveSourceServerAction(cfg, client)
			return
		}
		if _mgnRemoveTemplateAction {
			mgn_RemoveTemplateAction(cfg, client)
			return
		}
		if _mgnResumeReplication {
			mgn_ResumeReplication(cfg, client)
			return
		}
		if _mgnRetryDataReplication {
			mgn_RetryDataReplication(cfg, client)
			return
		}
		if _mgnStartCutover {
			mgn_StartCutover(cfg, client)
			return
		}
		if _mgnStartExport {
			mgn_StartExport(cfg, client)
			return
		}
		if _mgnStartImport {
			mgn_StartImport(cfg, client)
			return
		}
		if _mgnStartReplication {
			mgn_StartReplication(cfg, client)
			return
		}
		if _mgnStartTest {
			mgn_StartTest(cfg, client)
			return
		}
		if _mgnStopReplication {
			mgn_StopReplication(cfg, client)
			return
		}
		if _mgnTagResource {
			mgn_TagResource(cfg, client)
			return
		}
		if _mgnTerminateTargetInstances {
			mgn_TerminateTargetInstances(cfg, client)
			return
		}
		if _mgnUnarchiveApplication {
			mgn_UnarchiveApplication(cfg, client)
			return
		}
		if _mgnUnarchiveWave {
			mgn_UnarchiveWave(cfg, client)
			return
		}
		if _mgnUntagResource {
			mgn_UntagResource(cfg, client)
			return
		}
		if _mgnUpdateApplication {
			mgn_UpdateApplication(cfg, client)
			return
		}
		if _mgnUpdateConnector {
			mgn_UpdateConnector(cfg, client)
			return
		}
		if _mgnUpdateLaunchConfiguration {
			mgn_UpdateLaunchConfiguration(cfg, client)
			return
		}
		if _mgnUpdateLaunchConfigurationTemplate {
			mgn_UpdateLaunchConfigurationTemplate(cfg, client)
			return
		}
		if _mgnUpdateReplicationConfiguration {
			mgn_UpdateReplicationConfiguration(cfg, client)
			return
		}
		if _mgnUpdateReplicationConfigurationTemplate {
			mgn_UpdateReplicationConfigurationTemplate(cfg, client)
			return
		}
		if _mgnUpdateSourceServer {
			mgn_UpdateSourceServer(cfg, client)
			return
		}
		if _mgnUpdateSourceServerReplicationType {
			mgn_UpdateSourceServerReplicationType(cfg, client)
			return
		}
		if _mgnUpdateWave {
			mgn_UpdateWave(cfg, client)
			return
		}

	},
}

var (
	_mgnArchiveApplication                        bool
	_mgnArchiveWave                               bool
	_mgnAssociateApplications                     bool
	_mgnAssociateSourceServers                    bool
	_mgnChangeServerLifeCycleState                bool
	_mgnCreateApplication                         bool
	_mgnCreateConnector                           bool
	_mgnCreateLaunchConfigurationTemplate         bool
	_mgnCreateReplicationConfigurationTemplate    bool
	_mgnCreateWave                                bool
	_mgnDeleteApplication                         bool
	_mgnDeleteConnector                           bool
	_mgnDeleteJob                                 bool
	_mgnDeleteLaunchConfigurationTemplate         bool
	_mgnDeleteReplicationConfigurationTemplate    bool
	_mgnDeleteSourceServer                        bool
	_mgnDeleteVcenterClient                       bool
	_mgnDeleteWave                                bool
	_mgnDescribeJobLogItems                       bool
	_mgnDescribeJobs                              bool
	_mgnDescribeLaunchConfigurationTemplates      bool
	_mgnDescribeReplicationConfigurationTemplates bool
	_mgnDescribeSourceServers                     bool
	_mgnDescribeVcenterClients                    bool
	_mgnDisassociateApplications                  bool
	_mgnDisassociateSourceServers                 bool
	_mgnDisconnectFromService                     bool
	_mgnFinalizeCutover                           bool
	_mgnGetLaunchConfiguration                    bool
	_mgnGetReplicationConfiguration               bool
	_mgnInitializeService                         bool
	_mgnListApplications                          bool
	_mgnListConnectors                            bool
	_mgnListExportErrors                          bool
	_mgnListExports                               bool
	_mgnListImportErrors                          bool
	_mgnListImports                               bool
	_mgnListManagedAccounts                       bool
	_mgnListSourceServerActions                   bool
	_mgnListTagsForResource                       bool
	_mgnListTemplateActions                       bool
	_mgnListWaves                                 bool
	_mgnMarkAsArchived                            bool
	_mgnPauseReplication                          bool
	_mgnPutSourceServerAction                     bool
	_mgnPutTemplateAction                         bool
	_mgnRemoveSourceServerAction                  bool
	_mgnRemoveTemplateAction                      bool
	_mgnResumeReplication                         bool
	_mgnRetryDataReplication                      bool
	_mgnStartCutover                              bool
	_mgnStartExport                               bool
	_mgnStartImport                               bool
	_mgnStartReplication                          bool
	_mgnStartTest                                 bool
	_mgnStopReplication                           bool
	_mgnTagResource                               bool
	_mgnTerminateTargetInstances                  bool
	_mgnUnarchiveApplication                      bool
	_mgnUnarchiveWave                             bool
	_mgnUntagResource                             bool
	_mgnUpdateApplication                         bool
	_mgnUpdateConnector                           bool
	_mgnUpdateLaunchConfiguration                 bool
	_mgnUpdateLaunchConfigurationTemplate         bool
	_mgnUpdateReplicationConfiguration            bool
	_mgnUpdateReplicationConfigurationTemplate    bool
	_mgnUpdateSourceServer                        bool
	_mgnUpdateSourceServerReplicationType         bool
	_mgnUpdateWave                                bool

	_mgnAccountID                           string
	_mgnActionID                            string
	_mgnActionName                          string
	_mgnActive                              string
	_mgnApplicationID                       string
	_mgnApplicationIDs                      []string
	_mgnArn                                 string
	_mgnAssociateDefaultSecurityGroup       string
	_mgnAssociatePublicIpAddress            string
	_mgnBandwidthThrottling                 string
	_mgnBootMode                            string
	_mgnCategory                            string
	_mgnClientToken                         string
	_mgnConnectorAction                     string
	_mgnConnectorID                         string
	_mgnCopyPrivateIp                       string
	_mgnCopyTags                            string
	_mgnCreatePublicIP                      string
	_mgnDataPlaneRouting                    string
	_mgnDefaultLargeStagingDiskType         string
	_mgnDescription                         string
	_mgnDocumentIdentifier                  string
	_mgnDocumentVersion                     string
	_mgnEbsEncryption                       string
	_mgnEbsEncryptionKeyArn                 string
	_mgnEnableMapAutoTagging                string
	_mgnEnableParametersEncryption          string
	_mgnExportID                            string
	_mgnExternalParameters                  string
	_mgnFilters                             string
	_mgnImportID                            string
	_mgnInternetProtocol                    string
	_mgnJobID                               string
	_mgnLargeVolumeConf                     string
	_mgnLaunchConfigurationTemplateID       string
	_mgnLaunchConfigurationTemplateIDs      []string
	_mgnLaunchDisposition                   string
	_mgnLicensing                           string
	_mgnLifeCycle                           string
	_mgnMapAutoTaggingMpeID                 string
	_mgnMaxResults                          string
	_mgnMustSucceedForCutover               string
	_mgnName                                string
	_mgnNextToken                           string
	_mgnOperatingSystem                     string
	_mgnOrder                               string
	_mgnParameters                          string
	_mgnParametersEncryptionKey             string
	_mgnPostLaunchActions                   string
	_mgnReplicatedDisks                     string
	_mgnReplicationConfigurationTemplateID  string
	_mgnReplicationConfigurationTemplateIDs []string
	_mgnReplicationServerInstanceType       string
	_mgnReplicationServersSecurityGroupsIDs []string
	_mgnReplicationType                     string
	_mgnResourceArn                         string
	_mgnS3Bucket                            string
	_mgnS3BucketOwner                       string
	_mgnS3BucketSource                      string
	_mgnS3Key                               string
	_mgnSmallVolumeConf                     string
	_mgnSmallVolumeMaxSize                  string
	_mgnSourceServerID                      string
	_mgnSourceServerIDs                     []string
	_mgnSsmCommandConfig                    string
	_mgnSsmInstanceID                       string
	_mgnStagingAreaSubnetId                 string
	_mgnStagingAreaTags                     string
	_mgnTagKeys                             []string
	_mgnTags                                string
	_mgnTargetInstanceTypeRightSizingMethod string
	_mgnTimeoutSeconds                      string
	_mgnUseDedicatedReplicationServer       string
	_mgnUseFipsEndpoint                     string
	_mgnVcenterClientID                     string
	_mgnWaveID                              string
)

// Archive application.
func mgn_ArchiveApplication(cfg aws.Config, client *mgn.Client) {
	input := &mgn.ArchiveApplicationInput{
		// ApplicationID: *string, // Required
	}

	if len(_mgnApplicationID) > 0 {
		input.ApplicationID = aws.String(_mgnApplicationID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.ArchiveApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Archive wave.
func mgn_ArchiveWave(cfg aws.Config, client *mgn.Client) {
	input := &mgn.ArchiveWaveInput{
		// WaveID: *string, // Required
	}

	if len(_mgnWaveID) > 0 {
		input.WaveID = aws.String(_mgnWaveID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.ArchiveWave(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associate applications to wave.
func mgn_AssociateApplications(cfg aws.Config, client *mgn.Client) {
	input := &mgn.AssociateApplicationsInput{
		// ApplicationIDs: []string, // Required
		// WaveID: *string, // Required
	}

	if len(_mgnApplicationIDs) > 0 {
		input.ApplicationIDs = append([]string(nil), _mgnApplicationIDs...)
	}
	if len(_mgnWaveID) > 0 {
		input.WaveID = aws.String(_mgnWaveID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.AssociateApplications(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associate source servers to application.
func mgn_AssociateSourceServers(cfg aws.Config, client *mgn.Client) {
	input := &mgn.AssociateSourceServersInput{
		// ApplicationID: *string, // Required
		// SourceServerIDs: []string, // Required
	}

	if len(_mgnApplicationID) > 0 {
		input.ApplicationID = aws.String(_mgnApplicationID)
	}
	if len(_mgnSourceServerIDs) > 0 {
		input.SourceServerIDs = append([]string(nil), _mgnSourceServerIDs...)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.AssociateSourceServers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows the user to set the SourceServer.LifeCycle.state property for specific
// Source Server IDs to one of the following: READY_FOR_TEST or READY_FOR_CUTOVER.
// This command only works if the Source Server is already launchable
// (dataReplicationInfo.lagDuration is not null.)
func mgn_ChangeServerLifeCycleState(cfg aws.Config, client *mgn.Client) {
	input := &mgn.ChangeServerLifeCycleStateInput{
		// LifeCycle: *types.ChangeServerLifeCycleStateSourceServerLifecycle, // Required
		// SourceServerID: *string, // Required
	}

	if len(_mgnLifeCycle) > 0 {
		if err := assignInputField(input, "LifeCycle", _mgnLifeCycle); err != nil {
			log.Errorf("invalid --life-cycle: %s", err.Error())
			return
		}
	}
	if len(_mgnSourceServerID) > 0 {
		input.SourceServerID = aws.String(_mgnSourceServerID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.ChangeServerLifeCycleState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create application.
func mgn_CreateApplication(cfg aws.Config, client *mgn.Client) {
	input := &mgn.CreateApplicationInput{
		// Name: *string, // Required
	}

	if len(_mgnName) > 0 {
		input.Name = aws.String(_mgnName)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}
	if len(_mgnDescription) > 0 {
		input.Description = aws.String(_mgnDescription)
	}
	if len(_mgnTags) > 0 {
		if err := assignInputField(input, "Tags", _mgnTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create Connector.
func mgn_CreateConnector(cfg aws.Config, client *mgn.Client) {
	input := &mgn.CreateConnectorInput{
		// Name: *string, // Required
		// SsmInstanceID: *string, // Required
	}

	if len(_mgnName) > 0 {
		input.Name = aws.String(_mgnName)
	}
	if len(_mgnSsmInstanceID) > 0 {
		input.SsmInstanceID = aws.String(_mgnSsmInstanceID)
	}
	if len(_mgnSsmCommandConfig) > 0 {
		if err := assignInputField(input, "SsmCommandConfig", _mgnSsmCommandConfig); err != nil {
			log.Errorf("invalid --ssm-command-config: %s", err.Error())
			return
		}
	}
	if len(_mgnTags) > 0 {
		if err := assignInputField(input, "Tags", _mgnTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Launch Configuration Template.
func mgn_CreateLaunchConfigurationTemplate(cfg aws.Config, client *mgn.Client) {
	input := &mgn.CreateLaunchConfigurationTemplateInput{}

	if len(_mgnAssociatePublicIpAddress) > 0 {
		if err := assignInputField(input, "AssociatePublicIpAddress", _mgnAssociatePublicIpAddress); err != nil {
			log.Errorf("invalid --associate-public-ip-address: %s", err.Error())
			return
		}
	}
	if len(_mgnBootMode) > 0 {
		if err := assignInputField(input, "BootMode", _mgnBootMode); err != nil {
			log.Errorf("invalid --boot-mode: %s", err.Error())
			return
		}
	}
	if len(_mgnCopyPrivateIp) > 0 {
		if err := assignInputField(input, "CopyPrivateIp", _mgnCopyPrivateIp); err != nil {
			log.Errorf("invalid --copy-private-ip: %s", err.Error())
			return
		}
	}
	if len(_mgnCopyTags) > 0 {
		if err := assignInputField(input, "CopyTags", _mgnCopyTags); err != nil {
			log.Errorf("invalid --copy-tags: %s", err.Error())
			return
		}
	}
	if len(_mgnEnableMapAutoTagging) > 0 {
		if err := assignInputField(input, "EnableMapAutoTagging", _mgnEnableMapAutoTagging); err != nil {
			log.Errorf("invalid --enable-map-auto-tagging: %s", err.Error())
			return
		}
	}
	if len(_mgnEnableParametersEncryption) > 0 {
		if err := assignInputField(input, "EnableParametersEncryption", _mgnEnableParametersEncryption); err != nil {
			log.Errorf("invalid --enable-parameters-encryption: %s", err.Error())
			return
		}
	}
	if len(_mgnLargeVolumeConf) > 0 {
		if err := assignInputField(input, "LargeVolumeConf", _mgnLargeVolumeConf); err != nil {
			log.Errorf("invalid --large-volume-conf: %s", err.Error())
			return
		}
	}
	if len(_mgnLaunchDisposition) > 0 {
		if err := assignInputField(input, "LaunchDisposition", _mgnLaunchDisposition); err != nil {
			log.Errorf("invalid --launch-disposition: %s", err.Error())
			return
		}
	}
	if len(_mgnLicensing) > 0 {
		if err := assignInputField(input, "Licensing", _mgnLicensing); err != nil {
			log.Errorf("invalid --licensing: %s", err.Error())
			return
		}
	}
	if len(_mgnMapAutoTaggingMpeID) > 0 {
		input.MapAutoTaggingMpeID = aws.String(_mgnMapAutoTaggingMpeID)
	}
	if len(_mgnParametersEncryptionKey) > 0 {
		input.ParametersEncryptionKey = aws.String(_mgnParametersEncryptionKey)
	}
	if len(_mgnPostLaunchActions) > 0 {
		if err := assignInputField(input, "PostLaunchActions", _mgnPostLaunchActions); err != nil {
			log.Errorf("invalid --post-launch-actions: %s", err.Error())
			return
		}
	}
	if len(_mgnSmallVolumeConf) > 0 {
		if err := assignInputField(input, "SmallVolumeConf", _mgnSmallVolumeConf); err != nil {
			log.Errorf("invalid --small-volume-conf: %s", err.Error())
			return
		}
	}
	if len(_mgnSmallVolumeMaxSize) > 0 {
		if err := assignInputField(input, "SmallVolumeMaxSize", _mgnSmallVolumeMaxSize); err != nil {
			log.Errorf("invalid --small-volume-max-size: %s", err.Error())
			return
		}
	}
	if len(_mgnTags) > 0 {
		if err := assignInputField(input, "Tags", _mgnTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_mgnTargetInstanceTypeRightSizingMethod) > 0 {
		if err := assignInputField(input, "TargetInstanceTypeRightSizingMethod", _mgnTargetInstanceTypeRightSizingMethod); err != nil {
			log.Errorf("invalid --target-instance-type-right-sizing-method: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLaunchConfigurationTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new ReplicationConfigurationTemplate.
func mgn_CreateReplicationConfigurationTemplate(cfg aws.Config, client *mgn.Client) {
	input := &mgn.CreateReplicationConfigurationTemplateInput{
		// AssociateDefaultSecurityGroup: *bool, // Required
		// BandwidthThrottling: int64, // Required
		// CreatePublicIP: *bool, // Required
		// DataPlaneRouting: types.ReplicationConfigurationDataPlaneRouting, // Required
		// DefaultLargeStagingDiskType: types.ReplicationConfigurationDefaultLargeStagingDiskType, // Required
		// EbsEncryption: types.ReplicationConfigurationEbsEncryption, // Required
		// ReplicationServerInstanceType: *string, // Required
		// ReplicationServersSecurityGroupsIDs: []string, // Required
		// StagingAreaSubnetId: *string, // Required
		// StagingAreaTags: map[string]string, // Required
		// UseDedicatedReplicationServer: *bool, // Required
	}

	if len(_mgnAssociateDefaultSecurityGroup) > 0 {
		if err := assignInputField(input, "AssociateDefaultSecurityGroup", _mgnAssociateDefaultSecurityGroup); err != nil {
			log.Errorf("invalid --associate-default-security-group: %s", err.Error())
			return
		}
	}
	if len(_mgnBandwidthThrottling) > 0 {
		if err := assignInputField(input, "BandwidthThrottling", _mgnBandwidthThrottling); err != nil {
			log.Errorf("invalid --bandwidth-throttling: %s", err.Error())
			return
		}
	}
	if len(_mgnCreatePublicIP) > 0 {
		if err := assignInputField(input, "CreatePublicIP", _mgnCreatePublicIP); err != nil {
			log.Errorf("invalid --create-public-ip: %s", err.Error())
			return
		}
	}
	if len(_mgnDataPlaneRouting) > 0 {
		if err := assignInputField(input, "DataPlaneRouting", _mgnDataPlaneRouting); err != nil {
			log.Errorf("invalid --data-plane-routing: %s", err.Error())
			return
		}
	}
	if len(_mgnDefaultLargeStagingDiskType) > 0 {
		if err := assignInputField(input, "DefaultLargeStagingDiskType", _mgnDefaultLargeStagingDiskType); err != nil {
			log.Errorf("invalid --default-large-staging-disk-type: %s", err.Error())
			return
		}
	}
	if len(_mgnEbsEncryption) > 0 {
		if err := assignInputField(input, "EbsEncryption", _mgnEbsEncryption); err != nil {
			log.Errorf("invalid --ebs-encryption: %s", err.Error())
			return
		}
	}
	if len(_mgnReplicationServerInstanceType) > 0 {
		input.ReplicationServerInstanceType = aws.String(_mgnReplicationServerInstanceType)
	}
	if len(_mgnReplicationServersSecurityGroupsIDs) > 0 {
		input.ReplicationServersSecurityGroupsIDs = append([]string(nil), _mgnReplicationServersSecurityGroupsIDs...)
	}
	if len(_mgnStagingAreaSubnetId) > 0 {
		input.StagingAreaSubnetId = aws.String(_mgnStagingAreaSubnetId)
	}
	if len(_mgnStagingAreaTags) > 0 {
		if err := assignInputField(input, "StagingAreaTags", _mgnStagingAreaTags); err != nil {
			log.Errorf("invalid --staging-area-tags: %s", err.Error())
			return
		}
	}
	if len(_mgnUseDedicatedReplicationServer) > 0 {
		if err := assignInputField(input, "UseDedicatedReplicationServer", _mgnUseDedicatedReplicationServer); err != nil {
			log.Errorf("invalid --use-dedicated-replication-server: %s", err.Error())
			return
		}
	}
	if len(_mgnEbsEncryptionKeyArn) > 0 {
		input.EbsEncryptionKeyArn = aws.String(_mgnEbsEncryptionKeyArn)
	}
	if len(_mgnInternetProtocol) > 0 {
		if err := assignInputField(input, "InternetProtocol", _mgnInternetProtocol); err != nil {
			log.Errorf("invalid --internet-protocol: %s", err.Error())
			return
		}
	}
	if len(_mgnTags) > 0 {
		if err := assignInputField(input, "Tags", _mgnTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_mgnUseFipsEndpoint) > 0 {
		if err := assignInputField(input, "UseFipsEndpoint", _mgnUseFipsEndpoint); err != nil {
			log.Errorf("invalid --use-fips-endpoint: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateReplicationConfigurationTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create wave.
func mgn_CreateWave(cfg aws.Config, client *mgn.Client) {
	input := &mgn.CreateWaveInput{
		// Name: *string, // Required
	}

	if len(_mgnName) > 0 {
		input.Name = aws.String(_mgnName)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}
	if len(_mgnDescription) > 0 {
		input.Description = aws.String(_mgnDescription)
	}
	if len(_mgnTags) > 0 {
		if err := assignInputField(input, "Tags", _mgnTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWave(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete application.
func mgn_DeleteApplication(cfg aws.Config, client *mgn.Client) {
	input := &mgn.DeleteApplicationInput{
		// ApplicationID: *string, // Required
	}

	if len(_mgnApplicationID) > 0 {
		input.ApplicationID = aws.String(_mgnApplicationID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.DeleteApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete Connector.
func mgn_DeleteConnector(cfg aws.Config, client *mgn.Client) {
	input := &mgn.DeleteConnectorInput{
		// ConnectorID: *string, // Required
	}

	if len(_mgnConnectorID) > 0 {
		input.ConnectorID = aws.String(_mgnConnectorID)
	}

	if resp, err := client.DeleteConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a single Job by ID.
func mgn_DeleteJob(cfg aws.Config, client *mgn.Client) {
	input := &mgn.DeleteJobInput{
		// JobID: *string, // Required
	}

	if len(_mgnJobID) > 0 {
		input.JobID = aws.String(_mgnJobID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.DeleteJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a single Launch Configuration Template by ID.
func mgn_DeleteLaunchConfigurationTemplate(cfg aws.Config, client *mgn.Client) {
	input := &mgn.DeleteLaunchConfigurationTemplateInput{
		// LaunchConfigurationTemplateID: *string, // Required
	}

	if len(_mgnLaunchConfigurationTemplateID) > 0 {
		input.LaunchConfigurationTemplateID = aws.String(_mgnLaunchConfigurationTemplateID)
	}

	if resp, err := client.DeleteLaunchConfigurationTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a single Replication Configuration Template by ID
func mgn_DeleteReplicationConfigurationTemplate(cfg aws.Config, client *mgn.Client) {
	input := &mgn.DeleteReplicationConfigurationTemplateInput{
		// ReplicationConfigurationTemplateID: *string, // Required
	}

	if len(_mgnReplicationConfigurationTemplateID) > 0 {
		input.ReplicationConfigurationTemplateID = aws.String(_mgnReplicationConfigurationTemplateID)
	}

	if resp, err := client.DeleteReplicationConfigurationTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a single source server by ID.
func mgn_DeleteSourceServer(cfg aws.Config, client *mgn.Client) {
	input := &mgn.DeleteSourceServerInput{
		// SourceServerID: *string, // Required
	}

	if len(_mgnSourceServerID) > 0 {
		input.SourceServerID = aws.String(_mgnSourceServerID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.DeleteSourceServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a given vCenter client by ID.
func mgn_DeleteVcenterClient(cfg aws.Config, client *mgn.Client) {
	input := &mgn.DeleteVcenterClientInput{
		// VcenterClientID: *string, // Required
	}

	if len(_mgnVcenterClientID) > 0 {
		input.VcenterClientID = aws.String(_mgnVcenterClientID)
	}

	if resp, err := client.DeleteVcenterClient(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete wave.
func mgn_DeleteWave(cfg aws.Config, client *mgn.Client) {
	input := &mgn.DeleteWaveInput{
		// WaveID: *string, // Required
	}

	if len(_mgnWaveID) > 0 {
		input.WaveID = aws.String(_mgnWaveID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.DeleteWave(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed job log items with paging.
func mgn_DescribeJobLogItems(cfg aws.Config, client *mgn.Client) {
	input := &mgn.DescribeJobLogItemsInput{
		// JobID: *string, // Required
	}

	if len(_mgnJobID) > 0 {
		input.JobID = aws.String(_mgnJobID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}
	if len(_mgnMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mgnMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mgnNextToken) > 0 {
		input.NextToken = aws.String(_mgnNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeJobLogItems(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mgn.DescribeJobLogItemsOutput
	p := mgn.NewDescribeJobLogItemsPaginator(client, input)
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

// Returns a list of Jobs. Use the JobsID and fromDate and toData filters to limit
// which jobs are returned. The response is sorted by creationDataTime - latest
// date first. Jobs are normally created by the StartTest, StartCutover, and
// TerminateTargetInstances APIs. Jobs are also created by DiagnosticLaunch and
// TerminateDiagnosticInstances, which are APIs available only to *Support* and
// only used in response to relevant support tickets.
func mgn_DescribeJobs(cfg aws.Config, client *mgn.Client) {
	input := &mgn.DescribeJobsInput{}

	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}
	if len(_mgnFilters) > 0 {
		if err := assignInputField(input, "Filters", _mgnFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_mgnMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mgnMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mgnNextToken) > 0 {
		input.NextToken = aws.String(_mgnNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mgn.DescribeJobsOutput
	p := mgn.NewDescribeJobsPaginator(client, input)
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

// Lists all Launch Configuration Templates, filtered by Launch Configuration
// Template IDs
func mgn_DescribeLaunchConfigurationTemplates(cfg aws.Config, client *mgn.Client) {
	input := &mgn.DescribeLaunchConfigurationTemplatesInput{}

	if len(_mgnLaunchConfigurationTemplateIDs) > 0 {
		input.LaunchConfigurationTemplateIDs = append([]string(nil), _mgnLaunchConfigurationTemplateIDs...)
	}
	if len(_mgnMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mgnMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mgnNextToken) > 0 {
		input.NextToken = aws.String(_mgnNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeLaunchConfigurationTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mgn.DescribeLaunchConfigurationTemplatesOutput
	p := mgn.NewDescribeLaunchConfigurationTemplatesPaginator(client, input)
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

// Lists all ReplicationConfigurationTemplates, filtered by Source Server IDs.
func mgn_DescribeReplicationConfigurationTemplates(cfg aws.Config, client *mgn.Client) {
	input := &mgn.DescribeReplicationConfigurationTemplatesInput{}

	if len(_mgnMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mgnMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mgnNextToken) > 0 {
		input.NextToken = aws.String(_mgnNextToken)
	}
	if len(_mgnReplicationConfigurationTemplateIDs) > 0 {
		input.ReplicationConfigurationTemplateIDs = append([]string(nil), _mgnReplicationConfigurationTemplateIDs...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeReplicationConfigurationTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mgn.DescribeReplicationConfigurationTemplatesOutput
	p := mgn.NewDescribeReplicationConfigurationTemplatesPaginator(client, input)
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

// Retrieves all SourceServers or multiple SourceServers by ID.
func mgn_DescribeSourceServers(cfg aws.Config, client *mgn.Client) {
	input := &mgn.DescribeSourceServersInput{}

	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}
	if len(_mgnFilters) > 0 {
		if err := assignInputField(input, "Filters", _mgnFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_mgnMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mgnMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mgnNextToken) > 0 {
		input.NextToken = aws.String(_mgnNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeSourceServers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mgn.DescribeSourceServersOutput
	p := mgn.NewDescribeSourceServersPaginator(client, input)
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

// Returns a list of the installed vCenter clients.
func mgn_DescribeVcenterClients(cfg aws.Config, client *mgn.Client) {
	input := &mgn.DescribeVcenterClientsInput{}

	if len(_mgnMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mgnMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mgnNextToken) > 0 {
		input.NextToken = aws.String(_mgnNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeVcenterClients(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mgn.DescribeVcenterClientsOutput
	p := mgn.NewDescribeVcenterClientsPaginator(client, input)
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

// Disassociate applications from wave.
func mgn_DisassociateApplications(cfg aws.Config, client *mgn.Client) {
	input := &mgn.DisassociateApplicationsInput{
		// ApplicationIDs: []string, // Required
		// WaveID: *string, // Required
	}

	if len(_mgnApplicationIDs) > 0 {
		input.ApplicationIDs = append([]string(nil), _mgnApplicationIDs...)
	}
	if len(_mgnWaveID) > 0 {
		input.WaveID = aws.String(_mgnWaveID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.DisassociateApplications(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociate source servers from application.
func mgn_DisassociateSourceServers(cfg aws.Config, client *mgn.Client) {
	input := &mgn.DisassociateSourceServersInput{
		// ApplicationID: *string, // Required
		// SourceServerIDs: []string, // Required
	}

	if len(_mgnApplicationID) > 0 {
		input.ApplicationID = aws.String(_mgnApplicationID)
	}
	if len(_mgnSourceServerIDs) > 0 {
		input.SourceServerIDs = append([]string(nil), _mgnSourceServerIDs...)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.DisassociateSourceServers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disconnects specific Source Servers from Application Migration Service. Data
// replication is stopped immediately. All AWS resources created by Application
// Migration Service for enabling the replication of these source servers will be
// terminated / deleted within 90 minutes. Launched Test or Cutover instances will
// NOT be terminated. If the agent on the source server has not been prevented from
// communicating with the Application Migration Service service, then it will
// receive a command to uninstall itself (within approximately 10 minutes). The
// following properties of the SourceServer will be changed immediately:
// dataReplicationInfo.dataReplicationState will be set to DISCONNECTED; The
// totalStorageBytes property for each of dataReplicationInfo.replicatedDisks will
// be set to zero; dataReplicationInfo.lagDuration and
// dataReplicationInfo.lagDuration will be nullified.
func mgn_DisconnectFromService(cfg aws.Config, client *mgn.Client) {
	input := &mgn.DisconnectFromServiceInput{
		// SourceServerID: *string, // Required
	}

	if len(_mgnSourceServerID) > 0 {
		input.SourceServerID = aws.String(_mgnSourceServerID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.DisconnectFromService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Finalizes the cutover immediately for specific Source Servers. All AWS
// resources created by Application Migration Service for enabling the replication
// of these source servers will be terminated / deleted within 90 minutes. Launched
// Test or Cutover instances will NOT be terminated. The AWS Replication Agent will
// receive a command to uninstall itself (within 10 minutes). The following
// properties of the SourceServer will be changed immediately:
// dataReplicationInfo.dataReplicationState will be changed to DISCONNECTED; The
// SourceServer.lifeCycle.state will be changed to CUTOVER; The totalStorageBytes
// property fo each of dataReplicationInfo.replicatedDisks will be set to zero;
// dataReplicationInfo.lagDuration and dataReplicationInfo.lagDuration will be
// nullified.
func mgn_FinalizeCutover(cfg aws.Config, client *mgn.Client) {
	input := &mgn.FinalizeCutoverInput{
		// SourceServerID: *string, // Required
	}

	if len(_mgnSourceServerID) > 0 {
		input.SourceServerID = aws.String(_mgnSourceServerID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.FinalizeCutover(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all LaunchConfigurations available, filtered by Source Server IDs.
func mgn_GetLaunchConfiguration(cfg aws.Config, client *mgn.Client) {
	input := &mgn.GetLaunchConfigurationInput{
		// SourceServerID: *string, // Required
	}

	if len(_mgnSourceServerID) > 0 {
		input.SourceServerID = aws.String(_mgnSourceServerID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.GetLaunchConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all ReplicationConfigurations, filtered by Source Server ID.
func mgn_GetReplicationConfiguration(cfg aws.Config, client *mgn.Client) {
	input := &mgn.GetReplicationConfigurationInput{
		// SourceServerID: *string, // Required
	}

	if len(_mgnSourceServerID) > 0 {
		input.SourceServerID = aws.String(_mgnSourceServerID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.GetReplicationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initialize Application Migration Service.
func mgn_InitializeService(cfg aws.Config, client *mgn.Client) {
	input := &mgn.InitializeServiceInput{}

	if resp, err := client.InitializeService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves all applications or multiple applications by ID.
func mgn_ListApplications(cfg aws.Config, client *mgn.Client) {
	input := &mgn.ListApplicationsInput{}

	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}
	if len(_mgnFilters) > 0 {
		if err := assignInputField(input, "Filters", _mgnFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_mgnMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mgnMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mgnNextToken) > 0 {
		input.NextToken = aws.String(_mgnNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApplications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mgn.ListApplicationsOutput
	p := mgn.NewListApplicationsPaginator(client, input)
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

// List Connectors.
func mgn_ListConnectors(cfg aws.Config, client *mgn.Client) {
	input := &mgn.ListConnectorsInput{}

	if len(_mgnFilters) > 0 {
		if err := assignInputField(input, "Filters", _mgnFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_mgnMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mgnMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mgnNextToken) > 0 {
		input.NextToken = aws.String(_mgnNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConnectors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mgn.ListConnectorsOutput
	p := mgn.NewListConnectorsPaginator(client, input)
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

// List export errors.
func mgn_ListExportErrors(cfg aws.Config, client *mgn.Client) {
	input := &mgn.ListExportErrorsInput{
		// ExportID: *string, // Required
	}

	if len(_mgnExportID) > 0 {
		input.ExportID = aws.String(_mgnExportID)
	}
	if len(_mgnMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mgnMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mgnNextToken) > 0 {
		input.NextToken = aws.String(_mgnNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListExportErrors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mgn.ListExportErrorsOutput
	p := mgn.NewListExportErrorsPaginator(client, input)
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

// List exports.
func mgn_ListExports(cfg aws.Config, client *mgn.Client) {
	input := &mgn.ListExportsInput{}

	if len(_mgnFilters) > 0 {
		if err := assignInputField(input, "Filters", _mgnFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_mgnMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mgnMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mgnNextToken) > 0 {
		input.NextToken = aws.String(_mgnNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListExports(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mgn.ListExportsOutput
	p := mgn.NewListExportsPaginator(client, input)
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

// List import errors.
func mgn_ListImportErrors(cfg aws.Config, client *mgn.Client) {
	input := &mgn.ListImportErrorsInput{
		// ImportID: *string, // Required
	}

	if len(_mgnImportID) > 0 {
		input.ImportID = aws.String(_mgnImportID)
	}
	if len(_mgnMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mgnMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mgnNextToken) > 0 {
		input.NextToken = aws.String(_mgnNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListImportErrors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mgn.ListImportErrorsOutput
	p := mgn.NewListImportErrorsPaginator(client, input)
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

// List imports.
func mgn_ListImports(cfg aws.Config, client *mgn.Client) {
	input := &mgn.ListImportsInput{}

	if len(_mgnFilters) > 0 {
		if err := assignInputField(input, "Filters", _mgnFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_mgnMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mgnMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mgnNextToken) > 0 {
		input.NextToken = aws.String(_mgnNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListImports(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mgn.ListImportsOutput
	p := mgn.NewListImportsPaginator(client, input)
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

// List Managed Accounts.
func mgn_ListManagedAccounts(cfg aws.Config, client *mgn.Client) {
	input := &mgn.ListManagedAccountsInput{}

	if len(_mgnMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mgnMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mgnNextToken) > 0 {
		input.NextToken = aws.String(_mgnNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListManagedAccounts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mgn.ListManagedAccountsOutput
	p := mgn.NewListManagedAccountsPaginator(client, input)
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

// List source server post migration custom actions.
func mgn_ListSourceServerActions(cfg aws.Config, client *mgn.Client) {
	input := &mgn.ListSourceServerActionsInput{
		// SourceServerID: *string, // Required
	}

	if len(_mgnSourceServerID) > 0 {
		input.SourceServerID = aws.String(_mgnSourceServerID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}
	if len(_mgnFilters) > 0 {
		if err := assignInputField(input, "Filters", _mgnFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_mgnMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mgnMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mgnNextToken) > 0 {
		input.NextToken = aws.String(_mgnNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSourceServerActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mgn.ListSourceServerActionsOutput
	p := mgn.NewListSourceServerActionsPaginator(client, input)
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

// List all tags for your Application Migration Service resources.
func mgn_ListTagsForResource(cfg aws.Config, client *mgn.Client) {
	input := &mgn.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_mgnResourceArn) > 0 {
		input.ResourceArn = aws.String(_mgnResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List template post migration custom actions.
func mgn_ListTemplateActions(cfg aws.Config, client *mgn.Client) {
	input := &mgn.ListTemplateActionsInput{
		// LaunchConfigurationTemplateID: *string, // Required
	}

	if len(_mgnLaunchConfigurationTemplateID) > 0 {
		input.LaunchConfigurationTemplateID = aws.String(_mgnLaunchConfigurationTemplateID)
	}
	if len(_mgnFilters) > 0 {
		if err := assignInputField(input, "Filters", _mgnFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_mgnMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mgnMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mgnNextToken) > 0 {
		input.NextToken = aws.String(_mgnNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTemplateActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mgn.ListTemplateActionsOutput
	p := mgn.NewListTemplateActionsPaginator(client, input)
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

// Retrieves all waves or multiple waves by ID.
func mgn_ListWaves(cfg aws.Config, client *mgn.Client) {
	input := &mgn.ListWavesInput{}

	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}
	if len(_mgnFilters) > 0 {
		if err := assignInputField(input, "Filters", _mgnFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_mgnMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mgnMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mgnNextToken) > 0 {
		input.NextToken = aws.String(_mgnNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWaves(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mgn.ListWavesOutput
	p := mgn.NewListWavesPaginator(client, input)
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

// Archives specific Source Servers by setting the SourceServer.isArchived
// property to true for specified SourceServers by ID. This command only works for
// SourceServers with a lifecycle. state which equals DISCONNECTED or CUTOVER.
func mgn_MarkAsArchived(cfg aws.Config, client *mgn.Client) {
	input := &mgn.MarkAsArchivedInput{
		// SourceServerID: *string, // Required
	}

	if len(_mgnSourceServerID) > 0 {
		input.SourceServerID = aws.String(_mgnSourceServerID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.MarkAsArchived(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Pause Replication.
func mgn_PauseReplication(cfg aws.Config, client *mgn.Client) {
	input := &mgn.PauseReplicationInput{
		// SourceServerID: *string, // Required
	}

	if len(_mgnSourceServerID) > 0 {
		input.SourceServerID = aws.String(_mgnSourceServerID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.PauseReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Put source server post migration custom action.
func mgn_PutSourceServerAction(cfg aws.Config, client *mgn.Client) {
	input := &mgn.PutSourceServerActionInput{
		// ActionID: *string, // Required
		// ActionName: *string, // Required
		// DocumentIdentifier: *string, // Required
		// Order: *int32, // Required
		// SourceServerID: *string, // Required
	}

	if len(_mgnActionID) > 0 {
		input.ActionID = aws.String(_mgnActionID)
	}
	if len(_mgnActionName) > 0 {
		input.ActionName = aws.String(_mgnActionName)
	}
	if len(_mgnDocumentIdentifier) > 0 {
		input.DocumentIdentifier = aws.String(_mgnDocumentIdentifier)
	}
	if len(_mgnOrder) > 0 {
		if err := assignInputField(input, "Order", _mgnOrder); err != nil {
			log.Errorf("invalid --order: %s", err.Error())
			return
		}
	}
	if len(_mgnSourceServerID) > 0 {
		input.SourceServerID = aws.String(_mgnSourceServerID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}
	if len(_mgnActive) > 0 {
		if err := assignInputField(input, "Active", _mgnActive); err != nil {
			log.Errorf("invalid --active: %s", err.Error())
			return
		}
	}
	if len(_mgnCategory) > 0 {
		if err := assignInputField(input, "Category", _mgnCategory); err != nil {
			log.Errorf("invalid --category: %s", err.Error())
			return
		}
	}
	if len(_mgnDescription) > 0 {
		input.Description = aws.String(_mgnDescription)
	}
	if len(_mgnDocumentVersion) > 0 {
		input.DocumentVersion = aws.String(_mgnDocumentVersion)
	}
	if len(_mgnExternalParameters) > 0 {
		if err := assignInputField(input, "ExternalParameters", _mgnExternalParameters); err != nil {
			log.Errorf("invalid --external-parameters: %s", err.Error())
			return
		}
	}
	if len(_mgnMustSucceedForCutover) > 0 {
		if err := assignInputField(input, "MustSucceedForCutover", _mgnMustSucceedForCutover); err != nil {
			log.Errorf("invalid --must-succeed-for-cutover: %s", err.Error())
			return
		}
	}
	if len(_mgnParameters) > 0 {
		if err := assignInputField(input, "Parameters", _mgnParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_mgnTimeoutSeconds) > 0 {
		if err := assignInputField(input, "TimeoutSeconds", _mgnTimeoutSeconds); err != nil {
			log.Errorf("invalid --timeout-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutSourceServerAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Put template post migration custom action.
func mgn_PutTemplateAction(cfg aws.Config, client *mgn.Client) {
	input := &mgn.PutTemplateActionInput{
		// ActionID: *string, // Required
		// ActionName: *string, // Required
		// DocumentIdentifier: *string, // Required
		// LaunchConfigurationTemplateID: *string, // Required
		// Order: *int32, // Required
	}

	if len(_mgnActionID) > 0 {
		input.ActionID = aws.String(_mgnActionID)
	}
	if len(_mgnActionName) > 0 {
		input.ActionName = aws.String(_mgnActionName)
	}
	if len(_mgnDocumentIdentifier) > 0 {
		input.DocumentIdentifier = aws.String(_mgnDocumentIdentifier)
	}
	if len(_mgnLaunchConfigurationTemplateID) > 0 {
		input.LaunchConfigurationTemplateID = aws.String(_mgnLaunchConfigurationTemplateID)
	}
	if len(_mgnOrder) > 0 {
		if err := assignInputField(input, "Order", _mgnOrder); err != nil {
			log.Errorf("invalid --order: %s", err.Error())
			return
		}
	}
	if len(_mgnActive) > 0 {
		if err := assignInputField(input, "Active", _mgnActive); err != nil {
			log.Errorf("invalid --active: %s", err.Error())
			return
		}
	}
	if len(_mgnCategory) > 0 {
		if err := assignInputField(input, "Category", _mgnCategory); err != nil {
			log.Errorf("invalid --category: %s", err.Error())
			return
		}
	}
	if len(_mgnDescription) > 0 {
		input.Description = aws.String(_mgnDescription)
	}
	if len(_mgnDocumentVersion) > 0 {
		input.DocumentVersion = aws.String(_mgnDocumentVersion)
	}
	if len(_mgnExternalParameters) > 0 {
		if err := assignInputField(input, "ExternalParameters", _mgnExternalParameters); err != nil {
			log.Errorf("invalid --external-parameters: %s", err.Error())
			return
		}
	}
	if len(_mgnMustSucceedForCutover) > 0 {
		if err := assignInputField(input, "MustSucceedForCutover", _mgnMustSucceedForCutover); err != nil {
			log.Errorf("invalid --must-succeed-for-cutover: %s", err.Error())
			return
		}
	}
	if len(_mgnOperatingSystem) > 0 {
		input.OperatingSystem = aws.String(_mgnOperatingSystem)
	}
	if len(_mgnParameters) > 0 {
		if err := assignInputField(input, "Parameters", _mgnParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_mgnTimeoutSeconds) > 0 {
		if err := assignInputField(input, "TimeoutSeconds", _mgnTimeoutSeconds); err != nil {
			log.Errorf("invalid --timeout-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutTemplateAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove source server post migration custom action.
func mgn_RemoveSourceServerAction(cfg aws.Config, client *mgn.Client) {
	input := &mgn.RemoveSourceServerActionInput{
		// ActionID: *string, // Required
		// SourceServerID: *string, // Required
	}

	if len(_mgnActionID) > 0 {
		input.ActionID = aws.String(_mgnActionID)
	}
	if len(_mgnSourceServerID) > 0 {
		input.SourceServerID = aws.String(_mgnSourceServerID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.RemoveSourceServerAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove template post migration custom action.
func mgn_RemoveTemplateAction(cfg aws.Config, client *mgn.Client) {
	input := &mgn.RemoveTemplateActionInput{
		// ActionID: *string, // Required
		// LaunchConfigurationTemplateID: *string, // Required
	}

	if len(_mgnActionID) > 0 {
		input.ActionID = aws.String(_mgnActionID)
	}
	if len(_mgnLaunchConfigurationTemplateID) > 0 {
		input.LaunchConfigurationTemplateID = aws.String(_mgnLaunchConfigurationTemplateID)
	}

	if resp, err := client.RemoveTemplateAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resume Replication.
func mgn_ResumeReplication(cfg aws.Config, client *mgn.Client) {
	input := &mgn.ResumeReplicationInput{
		// SourceServerID: *string, // Required
	}

	if len(_mgnSourceServerID) > 0 {
		input.SourceServerID = aws.String(_mgnSourceServerID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.ResumeReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Causes the data replication initiation sequence to begin immediately upon next
// Handshake for specified SourceServer IDs, regardless of when the previous
// initiation started. This command will not work if the SourceServer is not
// stalled or is in a DISCONNECTED or STOPPED state.
func mgn_RetryDataReplication(cfg aws.Config, client *mgn.Client) {
	input := &mgn.RetryDataReplicationInput{
		// SourceServerID: *string, // Required
	}

	if len(_mgnSourceServerID) > 0 {
		input.SourceServerID = aws.String(_mgnSourceServerID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.RetryDataReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Launches a Cutover Instance for specific Source Servers. This command starts a
// LAUNCH job whose initiatedBy property is StartCutover and changes the
// SourceServer.lifeCycle.state property to CUTTING_OVER.
func mgn_StartCutover(cfg aws.Config, client *mgn.Client) {
	input := &mgn.StartCutoverInput{
		// SourceServerIDs: []string, // Required
	}

	if len(_mgnSourceServerIDs) > 0 {
		input.SourceServerIDs = append([]string(nil), _mgnSourceServerIDs...)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}
	if len(_mgnTags) > 0 {
		if err := assignInputField(input, "Tags", _mgnTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartCutover(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Start export.
func mgn_StartExport(cfg aws.Config, client *mgn.Client) {
	input := &mgn.StartExportInput{
		// S3Bucket: *string, // Required
		// S3Key: *string, // Required
	}

	if len(_mgnS3Bucket) > 0 {
		input.S3Bucket = aws.String(_mgnS3Bucket)
	}
	if len(_mgnS3Key) > 0 {
		input.S3Key = aws.String(_mgnS3Key)
	}
	if len(_mgnS3BucketOwner) > 0 {
		input.S3BucketOwner = aws.String(_mgnS3BucketOwner)
	}
	if len(_mgnTags) > 0 {
		if err := assignInputField(input, "Tags", _mgnTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Start import.
func mgn_StartImport(cfg aws.Config, client *mgn.Client) {
	input := &mgn.StartImportInput{
		// S3BucketSource: *types.S3BucketSource, // Required
	}

	if len(_mgnS3BucketSource) > 0 {
		if err := assignInputField(input, "S3BucketSource", _mgnS3BucketSource); err != nil {
			log.Errorf("invalid --s3-bucket-source: %s", err.Error())
			return
		}
	}
	if len(_mgnClientToken) > 0 {
		input.ClientToken = aws.String(_mgnClientToken)
	}
	if len(_mgnTags) > 0 {
		if err := assignInputField(input, "Tags", _mgnTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartImport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Start replication for source server irrespective of its replication type.
func mgn_StartReplication(cfg aws.Config, client *mgn.Client) {
	input := &mgn.StartReplicationInput{
		// SourceServerID: *string, // Required
	}

	if len(_mgnSourceServerID) > 0 {
		input.SourceServerID = aws.String(_mgnSourceServerID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.StartReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Launches a Test Instance for specific Source Servers. This command starts a
// LAUNCH job whose initiatedBy property is StartTest and changes the
// SourceServer.lifeCycle.state property to TESTING.
func mgn_StartTest(cfg aws.Config, client *mgn.Client) {
	input := &mgn.StartTestInput{
		// SourceServerIDs: []string, // Required
	}

	if len(_mgnSourceServerIDs) > 0 {
		input.SourceServerIDs = append([]string(nil), _mgnSourceServerIDs...)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}
	if len(_mgnTags) > 0 {
		if err := assignInputField(input, "Tags", _mgnTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartTest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stop Replication.
func mgn_StopReplication(cfg aws.Config, client *mgn.Client) {
	input := &mgn.StopReplicationInput{
		// SourceServerID: *string, // Required
	}

	if len(_mgnSourceServerID) > 0 {
		input.SourceServerID = aws.String(_mgnSourceServerID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.StopReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or overwrites only the specified tags for the specified Application
// Migration Service resource or resources. When you specify an existing tag key,
// the value is overwritten with the new value. Each resource can have a maximum of
// 50 tags. Each tag consists of a key and optional value.
func mgn_TagResource(cfg aws.Config, client *mgn.Client) {
	input := &mgn.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_mgnResourceArn) > 0 {
		input.ResourceArn = aws.String(_mgnResourceArn)
	}
	if len(_mgnTags) > 0 {
		if err := assignInputField(input, "Tags", _mgnTags); err != nil {
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

// Starts a job that terminates specific launched EC2 Test and Cutover instances.
// This command will not work for any Source Server with a lifecycle.state of
// TESTING, CUTTING_OVER, or CUTOVER.
func mgn_TerminateTargetInstances(cfg aws.Config, client *mgn.Client) {
	input := &mgn.TerminateTargetInstancesInput{
		// SourceServerIDs: []string, // Required
	}

	if len(_mgnSourceServerIDs) > 0 {
		input.SourceServerIDs = append([]string(nil), _mgnSourceServerIDs...)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}
	if len(_mgnTags) > 0 {
		if err := assignInputField(input, "Tags", _mgnTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TerminateTargetInstances(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Unarchive application.
func mgn_UnarchiveApplication(cfg aws.Config, client *mgn.Client) {
	input := &mgn.UnarchiveApplicationInput{
		// ApplicationID: *string, // Required
	}

	if len(_mgnApplicationID) > 0 {
		input.ApplicationID = aws.String(_mgnApplicationID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.UnarchiveApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Unarchive wave.
func mgn_UnarchiveWave(cfg aws.Config, client *mgn.Client) {
	input := &mgn.UnarchiveWaveInput{
		// WaveID: *string, // Required
	}

	if len(_mgnWaveID) > 0 {
		input.WaveID = aws.String(_mgnWaveID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.UnarchiveWave(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified set of tags from the specified set of Application
// Migration Service resources.
func mgn_UntagResource(cfg aws.Config, client *mgn.Client) {
	input := &mgn.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_mgnResourceArn) > 0 {
		input.ResourceArn = aws.String(_mgnResourceArn)
	}
	if len(_mgnTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _mgnTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update application.
func mgn_UpdateApplication(cfg aws.Config, client *mgn.Client) {
	input := &mgn.UpdateApplicationInput{
		// ApplicationID: *string, // Required
	}

	if len(_mgnApplicationID) > 0 {
		input.ApplicationID = aws.String(_mgnApplicationID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}
	if len(_mgnDescription) > 0 {
		input.Description = aws.String(_mgnDescription)
	}
	if len(_mgnName) > 0 {
		input.Name = aws.String(_mgnName)
	}

	if resp, err := client.UpdateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update Connector.
func mgn_UpdateConnector(cfg aws.Config, client *mgn.Client) {
	input := &mgn.UpdateConnectorInput{
		// ConnectorID: *string, // Required
	}

	if len(_mgnConnectorID) > 0 {
		input.ConnectorID = aws.String(_mgnConnectorID)
	}
	if len(_mgnName) > 0 {
		input.Name = aws.String(_mgnName)
	}
	if len(_mgnSsmCommandConfig) > 0 {
		if err := assignInputField(input, "SsmCommandConfig", _mgnSsmCommandConfig); err != nil {
			log.Errorf("invalid --ssm-command-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates multiple LaunchConfigurations by Source Server ID.
// bootMode valid values are LEGACY_BIOS | UEFI
func mgn_UpdateLaunchConfiguration(cfg aws.Config, client *mgn.Client) {
	input := &mgn.UpdateLaunchConfigurationInput{
		// SourceServerID: *string, // Required
	}

	if len(_mgnSourceServerID) > 0 {
		input.SourceServerID = aws.String(_mgnSourceServerID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}
	if len(_mgnBootMode) > 0 {
		if err := assignInputField(input, "BootMode", _mgnBootMode); err != nil {
			log.Errorf("invalid --boot-mode: %s", err.Error())
			return
		}
	}
	if len(_mgnCopyPrivateIp) > 0 {
		if err := assignInputField(input, "CopyPrivateIp", _mgnCopyPrivateIp); err != nil {
			log.Errorf("invalid --copy-private-ip: %s", err.Error())
			return
		}
	}
	if len(_mgnCopyTags) > 0 {
		if err := assignInputField(input, "CopyTags", _mgnCopyTags); err != nil {
			log.Errorf("invalid --copy-tags: %s", err.Error())
			return
		}
	}
	if len(_mgnEnableMapAutoTagging) > 0 {
		if err := assignInputField(input, "EnableMapAutoTagging", _mgnEnableMapAutoTagging); err != nil {
			log.Errorf("invalid --enable-map-auto-tagging: %s", err.Error())
			return
		}
	}
	if len(_mgnLaunchDisposition) > 0 {
		if err := assignInputField(input, "LaunchDisposition", _mgnLaunchDisposition); err != nil {
			log.Errorf("invalid --launch-disposition: %s", err.Error())
			return
		}
	}
	if len(_mgnLicensing) > 0 {
		if err := assignInputField(input, "Licensing", _mgnLicensing); err != nil {
			log.Errorf("invalid --licensing: %s", err.Error())
			return
		}
	}
	if len(_mgnMapAutoTaggingMpeID) > 0 {
		input.MapAutoTaggingMpeID = aws.String(_mgnMapAutoTaggingMpeID)
	}
	if len(_mgnName) > 0 {
		input.Name = aws.String(_mgnName)
	}
	if len(_mgnPostLaunchActions) > 0 {
		if err := assignInputField(input, "PostLaunchActions", _mgnPostLaunchActions); err != nil {
			log.Errorf("invalid --post-launch-actions: %s", err.Error())
			return
		}
	}
	if len(_mgnTargetInstanceTypeRightSizingMethod) > 0 {
		if err := assignInputField(input, "TargetInstanceTypeRightSizingMethod", _mgnTargetInstanceTypeRightSizingMethod); err != nil {
			log.Errorf("invalid --target-instance-type-right-sizing-method: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateLaunchConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing Launch Configuration Template by ID.
func mgn_UpdateLaunchConfigurationTemplate(cfg aws.Config, client *mgn.Client) {
	input := &mgn.UpdateLaunchConfigurationTemplateInput{
		// LaunchConfigurationTemplateID: *string, // Required
	}

	if len(_mgnLaunchConfigurationTemplateID) > 0 {
		input.LaunchConfigurationTemplateID = aws.String(_mgnLaunchConfigurationTemplateID)
	}
	if len(_mgnAssociatePublicIpAddress) > 0 {
		if err := assignInputField(input, "AssociatePublicIpAddress", _mgnAssociatePublicIpAddress); err != nil {
			log.Errorf("invalid --associate-public-ip-address: %s", err.Error())
			return
		}
	}
	if len(_mgnBootMode) > 0 {
		if err := assignInputField(input, "BootMode", _mgnBootMode); err != nil {
			log.Errorf("invalid --boot-mode: %s", err.Error())
			return
		}
	}
	if len(_mgnCopyPrivateIp) > 0 {
		if err := assignInputField(input, "CopyPrivateIp", _mgnCopyPrivateIp); err != nil {
			log.Errorf("invalid --copy-private-ip: %s", err.Error())
			return
		}
	}
	if len(_mgnCopyTags) > 0 {
		if err := assignInputField(input, "CopyTags", _mgnCopyTags); err != nil {
			log.Errorf("invalid --copy-tags: %s", err.Error())
			return
		}
	}
	if len(_mgnEnableMapAutoTagging) > 0 {
		if err := assignInputField(input, "EnableMapAutoTagging", _mgnEnableMapAutoTagging); err != nil {
			log.Errorf("invalid --enable-map-auto-tagging: %s", err.Error())
			return
		}
	}
	if len(_mgnEnableParametersEncryption) > 0 {
		if err := assignInputField(input, "EnableParametersEncryption", _mgnEnableParametersEncryption); err != nil {
			log.Errorf("invalid --enable-parameters-encryption: %s", err.Error())
			return
		}
	}
	if len(_mgnLargeVolumeConf) > 0 {
		if err := assignInputField(input, "LargeVolumeConf", _mgnLargeVolumeConf); err != nil {
			log.Errorf("invalid --large-volume-conf: %s", err.Error())
			return
		}
	}
	if len(_mgnLaunchDisposition) > 0 {
		if err := assignInputField(input, "LaunchDisposition", _mgnLaunchDisposition); err != nil {
			log.Errorf("invalid --launch-disposition: %s", err.Error())
			return
		}
	}
	if len(_mgnLicensing) > 0 {
		if err := assignInputField(input, "Licensing", _mgnLicensing); err != nil {
			log.Errorf("invalid --licensing: %s", err.Error())
			return
		}
	}
	if len(_mgnMapAutoTaggingMpeID) > 0 {
		input.MapAutoTaggingMpeID = aws.String(_mgnMapAutoTaggingMpeID)
	}
	if len(_mgnParametersEncryptionKey) > 0 {
		input.ParametersEncryptionKey = aws.String(_mgnParametersEncryptionKey)
	}
	if len(_mgnPostLaunchActions) > 0 {
		if err := assignInputField(input, "PostLaunchActions", _mgnPostLaunchActions); err != nil {
			log.Errorf("invalid --post-launch-actions: %s", err.Error())
			return
		}
	}
	if len(_mgnSmallVolumeConf) > 0 {
		if err := assignInputField(input, "SmallVolumeConf", _mgnSmallVolumeConf); err != nil {
			log.Errorf("invalid --small-volume-conf: %s", err.Error())
			return
		}
	}
	if len(_mgnSmallVolumeMaxSize) > 0 {
		if err := assignInputField(input, "SmallVolumeMaxSize", _mgnSmallVolumeMaxSize); err != nil {
			log.Errorf("invalid --small-volume-max-size: %s", err.Error())
			return
		}
	}
	if len(_mgnTargetInstanceTypeRightSizingMethod) > 0 {
		if err := assignInputField(input, "TargetInstanceTypeRightSizingMethod", _mgnTargetInstanceTypeRightSizingMethod); err != nil {
			log.Errorf("invalid --target-instance-type-right-sizing-method: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateLaunchConfigurationTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows you to update multiple ReplicationConfigurations by Source Server ID.
func mgn_UpdateReplicationConfiguration(cfg aws.Config, client *mgn.Client) {
	input := &mgn.UpdateReplicationConfigurationInput{
		// SourceServerID: *string, // Required
	}

	if len(_mgnSourceServerID) > 0 {
		input.SourceServerID = aws.String(_mgnSourceServerID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}
	if len(_mgnAssociateDefaultSecurityGroup) > 0 {
		if err := assignInputField(input, "AssociateDefaultSecurityGroup", _mgnAssociateDefaultSecurityGroup); err != nil {
			log.Errorf("invalid --associate-default-security-group: %s", err.Error())
			return
		}
	}
	if len(_mgnBandwidthThrottling) > 0 {
		if err := assignInputField(input, "BandwidthThrottling", _mgnBandwidthThrottling); err != nil {
			log.Errorf("invalid --bandwidth-throttling: %s", err.Error())
			return
		}
	}
	if len(_mgnCreatePublicIP) > 0 {
		if err := assignInputField(input, "CreatePublicIP", _mgnCreatePublicIP); err != nil {
			log.Errorf("invalid --create-public-ip: %s", err.Error())
			return
		}
	}
	if len(_mgnDataPlaneRouting) > 0 {
		if err := assignInputField(input, "DataPlaneRouting", _mgnDataPlaneRouting); err != nil {
			log.Errorf("invalid --data-plane-routing: %s", err.Error())
			return
		}
	}
	if len(_mgnDefaultLargeStagingDiskType) > 0 {
		if err := assignInputField(input, "DefaultLargeStagingDiskType", _mgnDefaultLargeStagingDiskType); err != nil {
			log.Errorf("invalid --default-large-staging-disk-type: %s", err.Error())
			return
		}
	}
	if len(_mgnEbsEncryption) > 0 {
		if err := assignInputField(input, "EbsEncryption", _mgnEbsEncryption); err != nil {
			log.Errorf("invalid --ebs-encryption: %s", err.Error())
			return
		}
	}
	if len(_mgnEbsEncryptionKeyArn) > 0 {
		input.EbsEncryptionKeyArn = aws.String(_mgnEbsEncryptionKeyArn)
	}
	if len(_mgnInternetProtocol) > 0 {
		if err := assignInputField(input, "InternetProtocol", _mgnInternetProtocol); err != nil {
			log.Errorf("invalid --internet-protocol: %s", err.Error())
			return
		}
	}
	if len(_mgnName) > 0 {
		input.Name = aws.String(_mgnName)
	}
	if len(_mgnReplicatedDisks) > 0 {
		if err := assignInputField(input, "ReplicatedDisks", _mgnReplicatedDisks); err != nil {
			log.Errorf("invalid --replicated-disks: %s", err.Error())
			return
		}
	}
	if len(_mgnReplicationServerInstanceType) > 0 {
		input.ReplicationServerInstanceType = aws.String(_mgnReplicationServerInstanceType)
	}
	if len(_mgnReplicationServersSecurityGroupsIDs) > 0 {
		input.ReplicationServersSecurityGroupsIDs = append([]string(nil), _mgnReplicationServersSecurityGroupsIDs...)
	}
	if len(_mgnStagingAreaSubnetId) > 0 {
		input.StagingAreaSubnetId = aws.String(_mgnStagingAreaSubnetId)
	}
	if len(_mgnStagingAreaTags) > 0 {
		if err := assignInputField(input, "StagingAreaTags", _mgnStagingAreaTags); err != nil {
			log.Errorf("invalid --staging-area-tags: %s", err.Error())
			return
		}
	}
	if len(_mgnUseDedicatedReplicationServer) > 0 {
		if err := assignInputField(input, "UseDedicatedReplicationServer", _mgnUseDedicatedReplicationServer); err != nil {
			log.Errorf("invalid --use-dedicated-replication-server: %s", err.Error())
			return
		}
	}
	if len(_mgnUseFipsEndpoint) > 0 {
		if err := assignInputField(input, "UseFipsEndpoint", _mgnUseFipsEndpoint); err != nil {
			log.Errorf("invalid --use-fips-endpoint: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateReplicationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates multiple ReplicationConfigurationTemplates by ID.
func mgn_UpdateReplicationConfigurationTemplate(cfg aws.Config, client *mgn.Client) {
	input := &mgn.UpdateReplicationConfigurationTemplateInput{
		// ReplicationConfigurationTemplateID: *string, // Required
	}

	if len(_mgnReplicationConfigurationTemplateID) > 0 {
		input.ReplicationConfigurationTemplateID = aws.String(_mgnReplicationConfigurationTemplateID)
	}
	if len(_mgnArn) > 0 {
		input.Arn = aws.String(_mgnArn)
	}
	if len(_mgnAssociateDefaultSecurityGroup) > 0 {
		if err := assignInputField(input, "AssociateDefaultSecurityGroup", _mgnAssociateDefaultSecurityGroup); err != nil {
			log.Errorf("invalid --associate-default-security-group: %s", err.Error())
			return
		}
	}
	if len(_mgnBandwidthThrottling) > 0 {
		if err := assignInputField(input, "BandwidthThrottling", _mgnBandwidthThrottling); err != nil {
			log.Errorf("invalid --bandwidth-throttling: %s", err.Error())
			return
		}
	}
	if len(_mgnCreatePublicIP) > 0 {
		if err := assignInputField(input, "CreatePublicIP", _mgnCreatePublicIP); err != nil {
			log.Errorf("invalid --create-public-ip: %s", err.Error())
			return
		}
	}
	if len(_mgnDataPlaneRouting) > 0 {
		if err := assignInputField(input, "DataPlaneRouting", _mgnDataPlaneRouting); err != nil {
			log.Errorf("invalid --data-plane-routing: %s", err.Error())
			return
		}
	}
	if len(_mgnDefaultLargeStagingDiskType) > 0 {
		if err := assignInputField(input, "DefaultLargeStagingDiskType", _mgnDefaultLargeStagingDiskType); err != nil {
			log.Errorf("invalid --default-large-staging-disk-type: %s", err.Error())
			return
		}
	}
	if len(_mgnEbsEncryption) > 0 {
		if err := assignInputField(input, "EbsEncryption", _mgnEbsEncryption); err != nil {
			log.Errorf("invalid --ebs-encryption: %s", err.Error())
			return
		}
	}
	if len(_mgnEbsEncryptionKeyArn) > 0 {
		input.EbsEncryptionKeyArn = aws.String(_mgnEbsEncryptionKeyArn)
	}
	if len(_mgnInternetProtocol) > 0 {
		if err := assignInputField(input, "InternetProtocol", _mgnInternetProtocol); err != nil {
			log.Errorf("invalid --internet-protocol: %s", err.Error())
			return
		}
	}
	if len(_mgnReplicationServerInstanceType) > 0 {
		input.ReplicationServerInstanceType = aws.String(_mgnReplicationServerInstanceType)
	}
	if len(_mgnReplicationServersSecurityGroupsIDs) > 0 {
		input.ReplicationServersSecurityGroupsIDs = append([]string(nil), _mgnReplicationServersSecurityGroupsIDs...)
	}
	if len(_mgnStagingAreaSubnetId) > 0 {
		input.StagingAreaSubnetId = aws.String(_mgnStagingAreaSubnetId)
	}
	if len(_mgnStagingAreaTags) > 0 {
		if err := assignInputField(input, "StagingAreaTags", _mgnStagingAreaTags); err != nil {
			log.Errorf("invalid --staging-area-tags: %s", err.Error())
			return
		}
	}
	if len(_mgnUseDedicatedReplicationServer) > 0 {
		if err := assignInputField(input, "UseDedicatedReplicationServer", _mgnUseDedicatedReplicationServer); err != nil {
			log.Errorf("invalid --use-dedicated-replication-server: %s", err.Error())
			return
		}
	}
	if len(_mgnUseFipsEndpoint) > 0 {
		if err := assignInputField(input, "UseFipsEndpoint", _mgnUseFipsEndpoint); err != nil {
			log.Errorf("invalid --use-fips-endpoint: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateReplicationConfigurationTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update Source Server.
func mgn_UpdateSourceServer(cfg aws.Config, client *mgn.Client) {
	input := &mgn.UpdateSourceServerInput{
		// SourceServerID: *string, // Required
	}

	if len(_mgnSourceServerID) > 0 {
		input.SourceServerID = aws.String(_mgnSourceServerID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}
	if len(_mgnConnectorAction) > 0 {
		if err := assignInputField(input, "ConnectorAction", _mgnConnectorAction); err != nil {
			log.Errorf("invalid --connector-action: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSourceServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows you to change between the AGENT_BASED replication type and the
// SNAPSHOT_SHIPPING replication type.
//
// SNAPSHOT_SHIPPING should be used for agentless replication.
func mgn_UpdateSourceServerReplicationType(cfg aws.Config, client *mgn.Client) {
	input := &mgn.UpdateSourceServerReplicationTypeInput{
		// ReplicationType: types.ReplicationType, // Required
		// SourceServerID: *string, // Required
	}

	if len(_mgnReplicationType) > 0 {
		if err := assignInputField(input, "ReplicationType", _mgnReplicationType); err != nil {
			log.Errorf("invalid --replication-type: %s", err.Error())
			return
		}
	}
	if len(_mgnSourceServerID) > 0 {
		input.SourceServerID = aws.String(_mgnSourceServerID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}

	if resp, err := client.UpdateSourceServerReplicationType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update wave.
func mgn_UpdateWave(cfg aws.Config, client *mgn.Client) {
	input := &mgn.UpdateWaveInput{
		// WaveID: *string, // Required
	}

	if len(_mgnWaveID) > 0 {
		input.WaveID = aws.String(_mgnWaveID)
	}
	if len(_mgnAccountID) > 0 {
		input.AccountID = aws.String(_mgnAccountID)
	}
	if len(_mgnDescription) > 0 {
		input.Description = aws.String(_mgnDescription)
	}
	if len(_mgnName) > 0 {
		input.Name = aws.String(_mgnName)
	}

	if resp, err := client.UpdateWave(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_mgnCmd)
	_mgnCmd.Flags().SortFlags = false

	_mgnCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_mgnCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_mgnCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_mgnCmd.Flags().StringVarP(&_mgnAccountID, "account-id", "", "", "Account ID")
	_mgnCmd.Flags().StringVarP(&_mgnActionID, "action-id", "", "", "Action ID")
	_mgnCmd.Flags().StringVarP(&_mgnActionName, "action-name", "", "", "Action Name")
	_mgnCmd.Flags().StringVarP(&_mgnActive, "active", "", "", "Active")
	_mgnCmd.Flags().StringVarP(&_mgnApplicationID, "application-id", "", "", "Application ID")
	_mgnCmd.Flags().StringSliceVarP(&_mgnApplicationIDs, "application-ids", "", nil, "Application Ids")
	_mgnCmd.Flags().StringVarP(&_mgnArn, "arn", "", "", "ARN")
	_mgnCmd.Flags().StringVarP(&_mgnAssociateDefaultSecurityGroup, "associate-default-security-group", "", "", "Associate Default Security Group")
	_mgnCmd.Flags().StringVarP(&_mgnAssociatePublicIpAddress, "associate-public-ip-address", "", "", "Associate Public IP Address")
	_mgnCmd.Flags().StringVarP(&_mgnBandwidthThrottling, "bandwidth-throttling", "", "", "Bandwidth Throttling")
	_mgnCmd.Flags().StringVarP(&_mgnBootMode, "boot-mode", "", "", "Boot Mode")
	_mgnCmd.Flags().StringVarP(&_mgnCategory, "category", "", "", "Category")
	_mgnCmd.Flags().StringVarP(&_mgnClientToken, "client-token", "", "", "Client Token")
	_mgnCmd.Flags().StringVarP(&_mgnConnectorAction, "connector-action", "", "", "Connector Action")
	_mgnCmd.Flags().StringVarP(&_mgnConnectorID, "connector-id", "", "", "Connector ID")
	_mgnCmd.Flags().StringVarP(&_mgnCopyPrivateIp, "copy-private-ip", "", "", "Copy Private IP")
	_mgnCmd.Flags().StringVarP(&_mgnCopyTags, "copy-tags", "", "", "Copy Tags")
	_mgnCmd.Flags().StringVarP(&_mgnCreatePublicIP, "create-public-ip", "", "", "Create Public IP")
	_mgnCmd.Flags().StringVarP(&_mgnDataPlaneRouting, "data-plane-routing", "", "", "Data Plane Routing")
	_mgnCmd.Flags().StringVarP(&_mgnDefaultLargeStagingDiskType, "default-large-staging-disk-type", "", "", "Default Large Staging Disk Type")
	_mgnCmd.Flags().StringVarP(&_mgnDescription, "description", "", "", "Description")
	_mgnCmd.Flags().StringVarP(&_mgnDocumentIdentifier, "document-identifier", "", "", "Document Identifier")
	_mgnCmd.Flags().StringVarP(&_mgnDocumentVersion, "document-version", "", "", "Document Version")
	_mgnCmd.Flags().StringVarP(&_mgnEbsEncryption, "ebs-encryption", "", "", "Ebs Encryption")
	_mgnCmd.Flags().StringVarP(&_mgnEbsEncryptionKeyArn, "ebs-encryption-key-arn", "", "", "Ebs Encryption Key ARN")
	_mgnCmd.Flags().StringVarP(&_mgnEnableMapAutoTagging, "enable-map-auto-tagging", "", "", "Enable Map Auto Tagging")
	_mgnCmd.Flags().StringVarP(&_mgnEnableParametersEncryption, "enable-parameters-encryption", "", "", "Enable Parameters Encryption")
	_mgnCmd.Flags().StringVarP(&_mgnExportID, "export-id", "", "", "Export ID")
	_mgnCmd.Flags().StringVarP(&_mgnExternalParameters, "external-parameters", "", "", "External Parameters")
	_mgnCmd.Flags().StringVarP(&_mgnFilters, "filters", "", "", "Filters")
	_mgnCmd.Flags().StringVarP(&_mgnImportID, "import-id", "", "", "Import ID")
	_mgnCmd.Flags().StringVarP(&_mgnInternetProtocol, "internet-protocol", "", "", "Internet Protocol")
	_mgnCmd.Flags().StringVarP(&_mgnJobID, "job-id", "", "", "Job ID")
	_mgnCmd.Flags().StringVarP(&_mgnLargeVolumeConf, "large-volume-conf", "", "", "Large Volume Conf")
	_mgnCmd.Flags().StringVarP(&_mgnLaunchConfigurationTemplateID, "launch-configuration-template-id", "", "", "Launch Configuration Template ID")
	_mgnCmd.Flags().StringSliceVarP(&_mgnLaunchConfigurationTemplateIDs, "launch-configuration-template-ids", "", nil, "Launch Configuration Template Ids")
	_mgnCmd.Flags().StringVarP(&_mgnLaunchDisposition, "launch-disposition", "", "", "Launch Disposition")
	_mgnCmd.Flags().StringVarP(&_mgnLicensing, "licensing", "", "", "Licensing")
	_mgnCmd.Flags().StringVarP(&_mgnLifeCycle, "life-cycle", "", "", "Life Cycle")
	_mgnCmd.Flags().StringVarP(&_mgnMapAutoTaggingMpeID, "map-auto-tagging-mpe-id", "", "", "Map Auto Tagging Mpe ID")
	_mgnCmd.Flags().StringVarP(&_mgnMaxResults, "max-results", "", "", "Max Results")
	_mgnCmd.Flags().StringVarP(&_mgnMustSucceedForCutover, "must-succeed-for-cutover", "", "", "Must Succeed For Cutover")
	_mgnCmd.Flags().StringVarP(&_mgnName, "name", "", "", "Name")
	_mgnCmd.Flags().StringVarP(&_mgnNextToken, "next-token", "", "", "Next Token")
	_mgnCmd.Flags().StringVarP(&_mgnOperatingSystem, "operating-system", "", "", "Operating System")
	_mgnCmd.Flags().StringVarP(&_mgnOrder, "order", "", "", "Order")
	_mgnCmd.Flags().StringVarP(&_mgnParameters, "parameters", "", "", "Parameters")
	_mgnCmd.Flags().StringVarP(&_mgnParametersEncryptionKey, "parameters-encryption-key", "", "", "Parameters Encryption Key")
	_mgnCmd.Flags().StringVarP(&_mgnPostLaunchActions, "post-launch-actions", "", "", "Post Launch Actions")
	_mgnCmd.Flags().StringVarP(&_mgnReplicatedDisks, "replicated-disks", "", "", "Replicated Disks")
	_mgnCmd.Flags().StringVarP(&_mgnReplicationConfigurationTemplateID, "replication-configuration-template-id", "", "", "Replication Configuration Template ID")
	_mgnCmd.Flags().StringSliceVarP(&_mgnReplicationConfigurationTemplateIDs, "replication-configuration-template-ids", "", nil, "Replication Configuration Template Ids")
	_mgnCmd.Flags().StringVarP(&_mgnReplicationServerInstanceType, "replication-server-instance-type", "", "", "Replication Server Instance Type")
	_mgnCmd.Flags().StringSliceVarP(&_mgnReplicationServersSecurityGroupsIDs, "replication-servers-security-groups-ids", "", nil, "Replication Servers Security Groups Ids")
	_mgnCmd.Flags().StringVarP(&_mgnReplicationType, "replication-type", "", "", "Replication Type")
	_mgnCmd.Flags().StringVarP(&_mgnResourceArn, "resource-arn", "", "", "Resource ARN")
	_mgnCmd.Flags().StringVarP(&_mgnS3Bucket, "s3-bucket", "", "", "S3 Bucket")
	_mgnCmd.Flags().StringVarP(&_mgnS3BucketOwner, "s3-bucket-owner", "", "", "S3 Bucket Owner")
	_mgnCmd.Flags().StringVarP(&_mgnS3BucketSource, "s3-bucket-source", "", "", "S3 Bucket Source")
	_mgnCmd.Flags().StringVarP(&_mgnS3Key, "s3-key", "", "", "S3 Key")
	_mgnCmd.Flags().StringVarP(&_mgnSmallVolumeConf, "small-volume-conf", "", "", "Small Volume Conf")
	_mgnCmd.Flags().StringVarP(&_mgnSmallVolumeMaxSize, "small-volume-max-size", "", "", "Small Volume Max Size")
	_mgnCmd.Flags().StringVarP(&_mgnSourceServerID, "source-server-id", "", "", "Source Server ID")
	_mgnCmd.Flags().StringSliceVarP(&_mgnSourceServerIDs, "source-server-ids", "", nil, "Source Server Ids")
	_mgnCmd.Flags().StringVarP(&_mgnSsmCommandConfig, "ssm-command-config", "", "", "Ssm Command Config")
	_mgnCmd.Flags().StringVarP(&_mgnSsmInstanceID, "ssm-instance-id", "", "", "Ssm Instance ID")
	_mgnCmd.Flags().StringVarP(&_mgnStagingAreaSubnetId, "staging-area-subnet-id", "", "", "Staging Area Subnet ID")
	_mgnCmd.Flags().StringVarP(&_mgnStagingAreaTags, "staging-area-tags", "", "", "Staging Area Tags")
	_mgnCmd.Flags().StringSliceVarP(&_mgnTagKeys, "tag-keys", "", nil, "Tag Keys")
	_mgnCmd.Flags().StringVarP(&_mgnTags, "tags", "", "", "Tags")
	_mgnCmd.Flags().StringVarP(&_mgnTargetInstanceTypeRightSizingMethod, "target-instance-type-right-sizing-method", "", "", "Target Instance Type Right Sizing Method")
	_mgnCmd.Flags().StringVarP(&_mgnTimeoutSeconds, "timeout-seconds", "", "", "Timeout Seconds")
	_mgnCmd.Flags().StringVarP(&_mgnUseDedicatedReplicationServer, "use-dedicated-replication-server", "", "", "Use Dedicated Replication Server")
	_mgnCmd.Flags().StringVarP(&_mgnUseFipsEndpoint, "use-fips-endpoint", "", "", "Use Fips Endpoint")
	_mgnCmd.Flags().StringVarP(&_mgnVcenterClientID, "vcenter-client-id", "", "", "Vcenter Client ID")
	_mgnCmd.Flags().StringVarP(&_mgnWaveID, "wave-id", "", "", "Wave ID")

	_mgnCmd.Flags().BoolVarP(&_mgnArchiveApplication, "archive-application", "", false, "Archive Application")
	_mgnCmd.Flags().BoolVarP(&_mgnArchiveWave, "archive-wave", "", false, "Archive Wave")
	_mgnCmd.Flags().BoolVarP(&_mgnAssociateApplications, "associate-applications", "", false, "Associate Applications")
	_mgnCmd.Flags().BoolVarP(&_mgnAssociateSourceServers, "associate-source-servers", "", false, "Associate Source Servers")
	_mgnCmd.Flags().BoolVarP(&_mgnChangeServerLifeCycleState, "change-server-life-cycle-state", "", false, "Change Server Life Cycle State")
	_mgnCmd.Flags().BoolVarP(&_mgnCreateApplication, "create-application", "", false, "Create Application")
	_mgnCmd.Flags().BoolVarP(&_mgnCreateConnector, "create-connector", "", false, "Create Connector")
	_mgnCmd.Flags().BoolVarP(&_mgnCreateLaunchConfigurationTemplate, "create-launch-configuration-template", "", false, "Create Launch Configuration Template")
	_mgnCmd.Flags().BoolVarP(&_mgnCreateReplicationConfigurationTemplate, "create-replication-configuration-template", "", false, "Create Replication Configuration Template")
	_mgnCmd.Flags().BoolVarP(&_mgnCreateWave, "create-wave", "", false, "Create Wave")
	_mgnCmd.Flags().BoolVarP(&_mgnDeleteApplication, "delete-application", "", false, "Delete Application")
	_mgnCmd.Flags().BoolVarP(&_mgnDeleteConnector, "delete-connector", "", false, "Delete Connector")
	_mgnCmd.Flags().BoolVarP(&_mgnDeleteJob, "delete-job", "", false, "Delete Job")
	_mgnCmd.Flags().BoolVarP(&_mgnDeleteLaunchConfigurationTemplate, "delete-launch-configuration-template", "", false, "Delete Launch Configuration Template")
	_mgnCmd.Flags().BoolVarP(&_mgnDeleteReplicationConfigurationTemplate, "delete-replication-configuration-template", "", false, "Delete Replication Configuration Template")
	_mgnCmd.Flags().BoolVarP(&_mgnDeleteSourceServer, "delete-source-server", "", false, "Delete Source Server")
	_mgnCmd.Flags().BoolVarP(&_mgnDeleteVcenterClient, "delete-vcenter-client", "", false, "Delete Vcenter Client")
	_mgnCmd.Flags().BoolVarP(&_mgnDeleteWave, "delete-wave", "", false, "Delete Wave")
	_mgnCmd.Flags().BoolVarP(&_mgnDescribeJobLogItems, "describe-job-log-items", "", false, "Describe Job Log Items")
	_mgnCmd.Flags().BoolVarP(&_mgnDescribeJobs, "describe-jobs", "", false, "Describe Jobs")
	_mgnCmd.Flags().BoolVarP(&_mgnDescribeLaunchConfigurationTemplates, "describe-launch-configuration-templates", "", false, "Describe Launch Configuration Templates")
	_mgnCmd.Flags().BoolVarP(&_mgnDescribeReplicationConfigurationTemplates, "describe-replication-configuration-templates", "", false, "Describe Replication Configuration Templates")
	_mgnCmd.Flags().BoolVarP(&_mgnDescribeSourceServers, "describe-source-servers", "", false, "Describe Source Servers")
	_mgnCmd.Flags().BoolVarP(&_mgnDescribeVcenterClients, "describe-vcenter-clients", "", false, "Describe Vcenter Clients")
	_mgnCmd.Flags().BoolVarP(&_mgnDisassociateApplications, "disassociate-applications", "", false, "Disassociate Applications")
	_mgnCmd.Flags().BoolVarP(&_mgnDisassociateSourceServers, "disassociate-source-servers", "", false, "Disassociate Source Servers")
	_mgnCmd.Flags().BoolVarP(&_mgnDisconnectFromService, "disconnect-from-service", "", false, "Disconnect From Service")
	_mgnCmd.Flags().BoolVarP(&_mgnFinalizeCutover, "finalize-cutover", "", false, "Finalize Cutover")
	_mgnCmd.Flags().BoolVarP(&_mgnGetLaunchConfiguration, "get-launch-configuration", "", false, "Get Launch Configuration")
	_mgnCmd.Flags().BoolVarP(&_mgnGetReplicationConfiguration, "get-replication-configuration", "", false, "Get Replication Configuration")
	_mgnCmd.Flags().BoolVarP(&_mgnInitializeService, "initialize-service", "", false, "Initialize Service")
	_mgnCmd.Flags().BoolVarP(&_mgnListApplications, "list-applications", "", false, "List Applications")
	_mgnCmd.Flags().BoolVarP(&_mgnListConnectors, "list-connectors", "", false, "List Connectors")
	_mgnCmd.Flags().BoolVarP(&_mgnListExportErrors, "list-export-errors", "", false, "List Export Errors")
	_mgnCmd.Flags().BoolVarP(&_mgnListExports, "list-exports", "", false, "List Exports")
	_mgnCmd.Flags().BoolVarP(&_mgnListImportErrors, "list-import-errors", "", false, "List Import Errors")
	_mgnCmd.Flags().BoolVarP(&_mgnListImports, "list-imports", "", false, "List Imports")
	_mgnCmd.Flags().BoolVarP(&_mgnListManagedAccounts, "list-managed-accounts", "", false, "List Managed Accounts")
	_mgnCmd.Flags().BoolVarP(&_mgnListSourceServerActions, "list-source-server-actions", "", false, "List Source Server Actions")
	_mgnCmd.Flags().BoolVarP(&_mgnListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_mgnCmd.Flags().BoolVarP(&_mgnListTemplateActions, "list-template-actions", "", false, "List Template Actions")
	_mgnCmd.Flags().BoolVarP(&_mgnListWaves, "list-waves", "", false, "List Waves")
	_mgnCmd.Flags().BoolVarP(&_mgnMarkAsArchived, "mark-as-archived", "", false, "Mark As Archived")
	_mgnCmd.Flags().BoolVarP(&_mgnPauseReplication, "pause-replication", "", false, "Pause Replication")
	_mgnCmd.Flags().BoolVarP(&_mgnPutSourceServerAction, "put-source-server-action", "", false, "Put Source Server Action")
	_mgnCmd.Flags().BoolVarP(&_mgnPutTemplateAction, "put-template-action", "", false, "Put Template Action")
	_mgnCmd.Flags().BoolVarP(&_mgnRemoveSourceServerAction, "remove-source-server-action", "", false, "Remove Source Server Action")
	_mgnCmd.Flags().BoolVarP(&_mgnRemoveTemplateAction, "remove-template-action", "", false, "Remove Template Action")
	_mgnCmd.Flags().BoolVarP(&_mgnResumeReplication, "resume-replication", "", false, "Resume Replication")
	_mgnCmd.Flags().BoolVarP(&_mgnRetryDataReplication, "retry-data-replication", "", false, "Retry Data Replication")
	_mgnCmd.Flags().BoolVarP(&_mgnStartCutover, "start-cutover", "", false, "Start Cutover")
	_mgnCmd.Flags().BoolVarP(&_mgnStartExport, "start-export", "", false, "Start Export")
	_mgnCmd.Flags().BoolVarP(&_mgnStartImport, "start-import", "", false, "Start Import")
	_mgnCmd.Flags().BoolVarP(&_mgnStartReplication, "start-replication", "", false, "Start Replication")
	_mgnCmd.Flags().BoolVarP(&_mgnStartTest, "start-test", "", false, "Start Test")
	_mgnCmd.Flags().BoolVarP(&_mgnStopReplication, "stop-replication", "", false, "Stop Replication")
	_mgnCmd.Flags().BoolVarP(&_mgnTagResource, "tag-resource", "", false, "Tag Resource")
	_mgnCmd.Flags().BoolVarP(&_mgnTerminateTargetInstances, "terminate-target-instances", "", false, "Terminate Target Instances")
	_mgnCmd.Flags().BoolVarP(&_mgnUnarchiveApplication, "unarchive-application", "", false, "Unarchive Application")
	_mgnCmd.Flags().BoolVarP(&_mgnUnarchiveWave, "unarchive-wave", "", false, "Unarchive Wave")
	_mgnCmd.Flags().BoolVarP(&_mgnUntagResource, "untag-resource", "", false, "Untag Resource")
	_mgnCmd.Flags().BoolVarP(&_mgnUpdateApplication, "update-application", "", false, "Update Application")
	_mgnCmd.Flags().BoolVarP(&_mgnUpdateConnector, "update-connector", "", false, "Update Connector")
	_mgnCmd.Flags().BoolVarP(&_mgnUpdateLaunchConfiguration, "update-launch-configuration", "", false, "Update Launch Configuration")
	_mgnCmd.Flags().BoolVarP(&_mgnUpdateLaunchConfigurationTemplate, "update-launch-configuration-template", "", false, "Update Launch Configuration Template")
	_mgnCmd.Flags().BoolVarP(&_mgnUpdateReplicationConfiguration, "update-replication-configuration", "", false, "Update Replication Configuration")
	_mgnCmd.Flags().BoolVarP(&_mgnUpdateReplicationConfigurationTemplate, "update-replication-configuration-template", "", false, "Update Replication Configuration Template")
	_mgnCmd.Flags().BoolVarP(&_mgnUpdateSourceServer, "update-source-server", "", false, "Update Source Server")
	_mgnCmd.Flags().BoolVarP(&_mgnUpdateSourceServerReplicationType, "update-source-server-replication-type", "", false, "Update Source Server Replication Type")
	_mgnCmd.Flags().BoolVarP(&_mgnUpdateWave, "update-wave", "", false, "Update Wave")

}
