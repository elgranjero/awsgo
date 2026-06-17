package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/drs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// drsCmd represents the drs command
var _drsCmd = &cobra.Command{
	Use:   "drs",
	Short: "AWS drs CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := drs.NewFromConfig(cfg)
		if _drsAssociateSourceNetworkStack {
			drs_AssociateSourceNetworkStack(cfg, client)
			return
		}
		if _drsCreateExtendedSourceServer {
			drs_CreateExtendedSourceServer(cfg, client)
			return
		}
		if _drsCreateLaunchConfigurationTemplate {
			drs_CreateLaunchConfigurationTemplate(cfg, client)
			return
		}
		if _drsCreateReplicationConfigurationTemplate {
			drs_CreateReplicationConfigurationTemplate(cfg, client)
			return
		}
		if _drsCreateSourceNetwork {
			drs_CreateSourceNetwork(cfg, client)
			return
		}
		if _drsDeleteJob {
			drs_DeleteJob(cfg, client)
			return
		}
		if _drsDeleteLaunchAction {
			drs_DeleteLaunchAction(cfg, client)
			return
		}
		if _drsDeleteLaunchConfigurationTemplate {
			drs_DeleteLaunchConfigurationTemplate(cfg, client)
			return
		}
		if _drsDeleteRecoveryInstance {
			drs_DeleteRecoveryInstance(cfg, client)
			return
		}
		if _drsDeleteReplicationConfigurationTemplate {
			drs_DeleteReplicationConfigurationTemplate(cfg, client)
			return
		}
		if _drsDeleteSourceNetwork {
			drs_DeleteSourceNetwork(cfg, client)
			return
		}
		if _drsDeleteSourceServer {
			drs_DeleteSourceServer(cfg, client)
			return
		}
		if _drsDescribeJobLogItems {
			drs_DescribeJobLogItems(cfg, client)
			return
		}
		if _drsDescribeJobs {
			drs_DescribeJobs(cfg, client)
			return
		}
		if _drsDescribeLaunchConfigurationTemplates {
			drs_DescribeLaunchConfigurationTemplates(cfg, client)
			return
		}
		if _drsDescribeRecoveryInstances {
			drs_DescribeRecoveryInstances(cfg, client)
			return
		}
		if _drsDescribeRecoverySnapshots {
			drs_DescribeRecoverySnapshots(cfg, client)
			return
		}
		if _drsDescribeReplicationConfigurationTemplates {
			drs_DescribeReplicationConfigurationTemplates(cfg, client)
			return
		}
		if _drsDescribeSourceNetworks {
			drs_DescribeSourceNetworks(cfg, client)
			return
		}
		if _drsDescribeSourceServers {
			drs_DescribeSourceServers(cfg, client)
			return
		}
		if _drsDisconnectRecoveryInstance {
			drs_DisconnectRecoveryInstance(cfg, client)
			return
		}
		if _drsDisconnectSourceServer {
			drs_DisconnectSourceServer(cfg, client)
			return
		}
		if _drsExportSourceNetworkCfnTemplate {
			drs_ExportSourceNetworkCfnTemplate(cfg, client)
			return
		}
		if _drsGetFailbackReplicationConfiguration {
			drs_GetFailbackReplicationConfiguration(cfg, client)
			return
		}
		if _drsGetLaunchConfiguration {
			drs_GetLaunchConfiguration(cfg, client)
			return
		}
		if _drsGetReplicationConfiguration {
			drs_GetReplicationConfiguration(cfg, client)
			return
		}
		if _drsInitializeService {
			drs_InitializeService(cfg, client)
			return
		}
		if _drsListExtensibleSourceServers {
			drs_ListExtensibleSourceServers(cfg, client)
			return
		}
		if _drsListLaunchActions {
			drs_ListLaunchActions(cfg, client)
			return
		}
		if _drsListStagingAccounts {
			drs_ListStagingAccounts(cfg, client)
			return
		}
		if _drsListTagsForResource {
			drs_ListTagsForResource(cfg, client)
			return
		}
		if _drsPutLaunchAction {
			drs_PutLaunchAction(cfg, client)
			return
		}
		if _drsRetryDataReplication {
			drs_RetryDataReplication(cfg, client)
			return
		}
		if _drsReverseReplication {
			drs_ReverseReplication(cfg, client)
			return
		}
		if _drsStartFailbackLaunch {
			drs_StartFailbackLaunch(cfg, client)
			return
		}
		if _drsStartRecovery {
			drs_StartRecovery(cfg, client)
			return
		}
		if _drsStartReplication {
			drs_StartReplication(cfg, client)
			return
		}
		if _drsStartSourceNetworkRecovery {
			drs_StartSourceNetworkRecovery(cfg, client)
			return
		}
		if _drsStartSourceNetworkReplication {
			drs_StartSourceNetworkReplication(cfg, client)
			return
		}
		if _drsStopFailback {
			drs_StopFailback(cfg, client)
			return
		}
		if _drsStopReplication {
			drs_StopReplication(cfg, client)
			return
		}
		if _drsStopSourceNetworkReplication {
			drs_StopSourceNetworkReplication(cfg, client)
			return
		}
		if _drsTagResource {
			drs_TagResource(cfg, client)
			return
		}
		if _drsTerminateRecoveryInstances {
			drs_TerminateRecoveryInstances(cfg, client)
			return
		}
		if _drsUntagResource {
			drs_UntagResource(cfg, client)
			return
		}
		if _drsUpdateFailbackReplicationConfiguration {
			drs_UpdateFailbackReplicationConfiguration(cfg, client)
			return
		}
		if _drsUpdateLaunchConfiguration {
			drs_UpdateLaunchConfiguration(cfg, client)
			return
		}
		if _drsUpdateLaunchConfigurationTemplate {
			drs_UpdateLaunchConfigurationTemplate(cfg, client)
			return
		}
		if _drsUpdateReplicationConfiguration {
			drs_UpdateReplicationConfiguration(cfg, client)
			return
		}
		if _drsUpdateReplicationConfigurationTemplate {
			drs_UpdateReplicationConfigurationTemplate(cfg, client)
			return
		}

	},
}

var (
	_drsAssociateSourceNetworkStack               bool
	_drsCreateExtendedSourceServer                bool
	_drsCreateLaunchConfigurationTemplate         bool
	_drsCreateReplicationConfigurationTemplate    bool
	_drsCreateSourceNetwork                       bool
	_drsDeleteJob                                 bool
	_drsDeleteLaunchAction                        bool
	_drsDeleteLaunchConfigurationTemplate         bool
	_drsDeleteRecoveryInstance                    bool
	_drsDeleteReplicationConfigurationTemplate    bool
	_drsDeleteSourceNetwork                       bool
	_drsDeleteSourceServer                        bool
	_drsDescribeJobLogItems                       bool
	_drsDescribeJobs                              bool
	_drsDescribeLaunchConfigurationTemplates      bool
	_drsDescribeRecoveryInstances                 bool
	_drsDescribeRecoverySnapshots                 bool
	_drsDescribeReplicationConfigurationTemplates bool
	_drsDescribeSourceNetworks                    bool
	_drsDescribeSourceServers                     bool
	_drsDisconnectRecoveryInstance                bool
	_drsDisconnectSourceServer                    bool
	_drsExportSourceNetworkCfnTemplate            bool
	_drsGetFailbackReplicationConfiguration       bool
	_drsGetLaunchConfiguration                    bool
	_drsGetReplicationConfiguration               bool
	_drsInitializeService                         bool
	_drsListExtensibleSourceServers               bool
	_drsListLaunchActions                         bool
	_drsListStagingAccounts                       bool
	_drsListTagsForResource                       bool
	_drsPutLaunchAction                           bool
	_drsRetryDataReplication                      bool
	_drsReverseReplication                        bool
	_drsStartFailbackLaunch                       bool
	_drsStartRecovery                             bool
	_drsStartReplication                          bool
	_drsStartSourceNetworkRecovery                bool
	_drsStartSourceNetworkReplication             bool
	_drsStopFailback                              bool
	_drsStopReplication                           bool
	_drsStopSourceNetworkReplication              bool
	_drsTagResource                               bool
	_drsTerminateRecoveryInstances                bool
	_drsUntagResource                             bool
	_drsUpdateFailbackReplicationConfiguration    bool
	_drsUpdateLaunchConfiguration                 bool
	_drsUpdateLaunchConfigurationTemplate         bool
	_drsUpdateReplicationConfiguration            bool
	_drsUpdateReplicationConfigurationTemplate    bool

	_drsActionCode                          string
	_drsActionId                            string
	_drsActionVersion                       string
	_drsActive                              string
	_drsArn                                 string
	_drsAssociateDefaultSecurityGroup       string
	_drsAutoReplicateNewDisks               string
	_drsBandwidthThrottling                 string
	_drsCategory                            string
	_drsCfnStackName                        string
	_drsCopyPrivateIp                       string
	_drsCopyTags                            string
	_drsCreatePublicIP                      string
	_drsDataPlaneRouting                    string
	_drsDefaultLargeStagingDiskType         string
	_drsDeployAsNew                         string
	_drsDescription                         string
	_drsEbsEncryption                       string
	_drsEbsEncryptionKeyArn                 string
	_drsExportBucketArn                     string
	_drsFilters                             string
	_drsIsDrill                             string
	_drsJobID                               string
	_drsLaunchConfigurationTemplateID       string
	_drsLaunchConfigurationTemplateIDs      []string
	_drsLaunchDisposition                   string
	_drsLaunchIntoInstanceProperties        string
	_drsLaunchIntoSourceInstance            string
	_drsLicensing                           string
	_drsMaxResults                          string
	_drsName                                string
	_drsNextToken                           string
	_drsOptional                            string
	_drsOrder                               string
	_drsOriginAccountID                     string
	_drsOriginRegion                        string
	_drsParameters                          string
	_drsPitPolicy                           string
	_drsPostLaunchEnabled                   string
	_drsRecoveryInstanceID                  string
	_drsRecoveryInstanceIDs                 []string
	_drsReplicatedDisks                     string
	_drsReplicationConfigurationTemplateID  string
	_drsReplicationConfigurationTemplateIDs []string
	_drsReplicationServerInstanceType       string
	_drsReplicationServersSecurityGroupsIDs []string
	_drsResourceArn                         string
	_drsResourceId                          string
	_drsSourceNetworkID                     string
	_drsSourceNetworks                      string
	_drsSourceServerArn                     string
	_drsSourceServerID                      string
	_drsSourceServers                       string
	_drsStagingAccountID                    string
	_drsStagingAreaSubnetId                 string
	_drsStagingAreaTags                     string
	_drsTagKeys                             []string
	_drsTags                                string
	_drsTargetInstanceTypeRightSizingMethod string
	_drsUseDedicatedReplicationServer       string
	_drsUsePrivateIP                        string
	_drsVpcID                               string
)

// Associate a Source Network to an existing CloudFormation Stack and modify
// launch templates to use this network. Can be used for reverting to previously
// deployed CloudFormation stacks.
func drs_AssociateSourceNetworkStack(cfg aws.Config, client *drs.Client) {
	input := &drs.AssociateSourceNetworkStackInput{
		// CfnStackName: *string, // Required
		// SourceNetworkID: *string, // Required
	}

	if len(_drsCfnStackName) > 0 {
		input.CfnStackName = aws.String(_drsCfnStackName)
	}
	if len(_drsSourceNetworkID) > 0 {
		input.SourceNetworkID = aws.String(_drsSourceNetworkID)
	}

	if resp, err := client.AssociateSourceNetworkStack(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create an extended source server in the target Account based on the source
// server in staging account.
func drs_CreateExtendedSourceServer(cfg aws.Config, client *drs.Client) {
	input := &drs.CreateExtendedSourceServerInput{
		// SourceServerArn: *string, // Required
	}

	if len(_drsSourceServerArn) > 0 {
		input.SourceServerArn = aws.String(_drsSourceServerArn)
	}
	if len(_drsTags) > 0 {
		if err := assignInputField(input, "Tags", _drsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateExtendedSourceServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Launch Configuration Template.
func drs_CreateLaunchConfigurationTemplate(cfg aws.Config, client *drs.Client) {
	input := &drs.CreateLaunchConfigurationTemplateInput{}

	if len(_drsCopyPrivateIp) > 0 {
		if err := assignInputField(input, "CopyPrivateIp", _drsCopyPrivateIp); err != nil {
			log.Errorf("invalid --copy-private-ip: %s", err.Error())
			return
		}
	}
	if len(_drsCopyTags) > 0 {
		if err := assignInputField(input, "CopyTags", _drsCopyTags); err != nil {
			log.Errorf("invalid --copy-tags: %s", err.Error())
			return
		}
	}
	if len(_drsExportBucketArn) > 0 {
		input.ExportBucketArn = aws.String(_drsExportBucketArn)
	}
	if len(_drsLaunchDisposition) > 0 {
		if err := assignInputField(input, "LaunchDisposition", _drsLaunchDisposition); err != nil {
			log.Errorf("invalid --launch-disposition: %s", err.Error())
			return
		}
	}
	if len(_drsLaunchIntoSourceInstance) > 0 {
		if err := assignInputField(input, "LaunchIntoSourceInstance", _drsLaunchIntoSourceInstance); err != nil {
			log.Errorf("invalid --launch-into-source-instance: %s", err.Error())
			return
		}
	}
	if len(_drsLicensing) > 0 {
		if err := assignInputField(input, "Licensing", _drsLicensing); err != nil {
			log.Errorf("invalid --licensing: %s", err.Error())
			return
		}
	}
	if len(_drsPostLaunchEnabled) > 0 {
		if err := assignInputField(input, "PostLaunchEnabled", _drsPostLaunchEnabled); err != nil {
			log.Errorf("invalid --post-launch-enabled: %s", err.Error())
			return
		}
	}
	if len(_drsTags) > 0 {
		if err := assignInputField(input, "Tags", _drsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_drsTargetInstanceTypeRightSizingMethod) > 0 {
		if err := assignInputField(input, "TargetInstanceTypeRightSizingMethod", _drsTargetInstanceTypeRightSizingMethod); err != nil {
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
func drs_CreateReplicationConfigurationTemplate(cfg aws.Config, client *drs.Client) {
	input := &drs.CreateReplicationConfigurationTemplateInput{
		// AssociateDefaultSecurityGroup: *bool, // Required
		// BandwidthThrottling: int64, // Required
		// CreatePublicIP: *bool, // Required
		// DataPlaneRouting: types.ReplicationConfigurationDataPlaneRouting, // Required
		// DefaultLargeStagingDiskType: types.ReplicationConfigurationDefaultLargeStagingDiskType, // Required
		// EbsEncryption: types.ReplicationConfigurationEbsEncryption, // Required
		// PitPolicy: []types.PITPolicyRule, // Required
		// ReplicationServerInstanceType: *string, // Required
		// ReplicationServersSecurityGroupsIDs: []string, // Required
		// StagingAreaSubnetId: *string, // Required
		// StagingAreaTags: map[string]string, // Required
		// UseDedicatedReplicationServer: *bool, // Required
	}

	if len(_drsAssociateDefaultSecurityGroup) > 0 {
		if err := assignInputField(input, "AssociateDefaultSecurityGroup", _drsAssociateDefaultSecurityGroup); err != nil {
			log.Errorf("invalid --associate-default-security-group: %s", err.Error())
			return
		}
	}
	if len(_drsBandwidthThrottling) > 0 {
		if err := assignInputField(input, "BandwidthThrottling", _drsBandwidthThrottling); err != nil {
			log.Errorf("invalid --bandwidth-throttling: %s", err.Error())
			return
		}
	}
	if len(_drsCreatePublicIP) > 0 {
		if err := assignInputField(input, "CreatePublicIP", _drsCreatePublicIP); err != nil {
			log.Errorf("invalid --create-public-ip: %s", err.Error())
			return
		}
	}
	if len(_drsDataPlaneRouting) > 0 {
		if err := assignInputField(input, "DataPlaneRouting", _drsDataPlaneRouting); err != nil {
			log.Errorf("invalid --data-plane-routing: %s", err.Error())
			return
		}
	}
	if len(_drsDefaultLargeStagingDiskType) > 0 {
		if err := assignInputField(input, "DefaultLargeStagingDiskType", _drsDefaultLargeStagingDiskType); err != nil {
			log.Errorf("invalid --default-large-staging-disk-type: %s", err.Error())
			return
		}
	}
	if len(_drsEbsEncryption) > 0 {
		if err := assignInputField(input, "EbsEncryption", _drsEbsEncryption); err != nil {
			log.Errorf("invalid --ebs-encryption: %s", err.Error())
			return
		}
	}
	if len(_drsPitPolicy) > 0 {
		if err := assignInputField(input, "PitPolicy", _drsPitPolicy); err != nil {
			log.Errorf("invalid --pit-policy: %s", err.Error())
			return
		}
	}
	if len(_drsReplicationServerInstanceType) > 0 {
		input.ReplicationServerInstanceType = aws.String(_drsReplicationServerInstanceType)
	}
	if len(_drsReplicationServersSecurityGroupsIDs) > 0 {
		input.ReplicationServersSecurityGroupsIDs = append([]string(nil), _drsReplicationServersSecurityGroupsIDs...)
	}
	if len(_drsStagingAreaSubnetId) > 0 {
		input.StagingAreaSubnetId = aws.String(_drsStagingAreaSubnetId)
	}
	if len(_drsStagingAreaTags) > 0 {
		if err := assignInputField(input, "StagingAreaTags", _drsStagingAreaTags); err != nil {
			log.Errorf("invalid --staging-area-tags: %s", err.Error())
			return
		}
	}
	if len(_drsUseDedicatedReplicationServer) > 0 {
		if err := assignInputField(input, "UseDedicatedReplicationServer", _drsUseDedicatedReplicationServer); err != nil {
			log.Errorf("invalid --use-dedicated-replication-server: %s", err.Error())
			return
		}
	}
	if len(_drsAutoReplicateNewDisks) > 0 {
		if err := assignInputField(input, "AutoReplicateNewDisks", _drsAutoReplicateNewDisks); err != nil {
			log.Errorf("invalid --auto-replicate-new-disks: %s", err.Error())
			return
		}
	}
	if len(_drsEbsEncryptionKeyArn) > 0 {
		input.EbsEncryptionKeyArn = aws.String(_drsEbsEncryptionKeyArn)
	}
	if len(_drsTags) > 0 {
		if err := assignInputField(input, "Tags", _drsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
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

// Create a new Source Network resource for a provided VPC ID.
func drs_CreateSourceNetwork(cfg aws.Config, client *drs.Client) {
	input := &drs.CreateSourceNetworkInput{
		// OriginAccountID: *string, // Required
		// OriginRegion: *string, // Required
		// VpcID: *string, // Required
	}

	if len(_drsOriginAccountID) > 0 {
		input.OriginAccountID = aws.String(_drsOriginAccountID)
	}
	if len(_drsOriginRegion) > 0 {
		input.OriginRegion = aws.String(_drsOriginRegion)
	}
	if len(_drsVpcID) > 0 {
		input.VpcID = aws.String(_drsVpcID)
	}
	if len(_drsTags) > 0 {
		if err := assignInputField(input, "Tags", _drsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSourceNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a single Job by ID.
func drs_DeleteJob(cfg aws.Config, client *drs.Client) {
	input := &drs.DeleteJobInput{
		// JobID: *string, // Required
	}

	if len(_drsJobID) > 0 {
		input.JobID = aws.String(_drsJobID)
	}

	if resp, err := client.DeleteJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a resource launch action.
func drs_DeleteLaunchAction(cfg aws.Config, client *drs.Client) {
	input := &drs.DeleteLaunchActionInput{
		// ActionId: *string, // Required
		// ResourceId: *string, // Required
	}

	if len(_drsActionId) > 0 {
		input.ActionId = aws.String(_drsActionId)
	}
	if len(_drsResourceId) > 0 {
		input.ResourceId = aws.String(_drsResourceId)
	}

	if resp, err := client.DeleteLaunchAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a single Launch Configuration Template by ID.
func drs_DeleteLaunchConfigurationTemplate(cfg aws.Config, client *drs.Client) {
	input := &drs.DeleteLaunchConfigurationTemplateInput{
		// LaunchConfigurationTemplateID: *string, // Required
	}

	if len(_drsLaunchConfigurationTemplateID) > 0 {
		input.LaunchConfigurationTemplateID = aws.String(_drsLaunchConfigurationTemplateID)
	}

	if resp, err := client.DeleteLaunchConfigurationTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a single Recovery Instance by ID. This deletes the Recovery Instance
// resource from Elastic Disaster Recovery. The Recovery Instance must be
// disconnected first in order to delete it.
func drs_DeleteRecoveryInstance(cfg aws.Config, client *drs.Client) {
	input := &drs.DeleteRecoveryInstanceInput{
		// RecoveryInstanceID: *string, // Required
	}

	if len(_drsRecoveryInstanceID) > 0 {
		input.RecoveryInstanceID = aws.String(_drsRecoveryInstanceID)
	}

	if resp, err := client.DeleteRecoveryInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a single Replication Configuration Template by ID
func drs_DeleteReplicationConfigurationTemplate(cfg aws.Config, client *drs.Client) {
	input := &drs.DeleteReplicationConfigurationTemplateInput{
		// ReplicationConfigurationTemplateID: *string, // Required
	}

	if len(_drsReplicationConfigurationTemplateID) > 0 {
		input.ReplicationConfigurationTemplateID = aws.String(_drsReplicationConfigurationTemplateID)
	}

	if resp, err := client.DeleteReplicationConfigurationTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete Source Network resource.
func drs_DeleteSourceNetwork(cfg aws.Config, client *drs.Client) {
	input := &drs.DeleteSourceNetworkInput{
		// SourceNetworkID: *string, // Required
	}

	if len(_drsSourceNetworkID) > 0 {
		input.SourceNetworkID = aws.String(_drsSourceNetworkID)
	}

	if resp, err := client.DeleteSourceNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a single Source Server by ID. The Source Server must be disconnected
// first.
func drs_DeleteSourceServer(cfg aws.Config, client *drs.Client) {
	input := &drs.DeleteSourceServerInput{
		// SourceServerID: *string, // Required
	}

	if len(_drsSourceServerID) > 0 {
		input.SourceServerID = aws.String(_drsSourceServerID)
	}

	if resp, err := client.DeleteSourceServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a detailed Job log with pagination.
func drs_DescribeJobLogItems(cfg aws.Config, client *drs.Client) {
	input := &drs.DescribeJobLogItemsInput{
		// JobID: *string, // Required
	}

	if len(_drsJobID) > 0 {
		input.JobID = aws.String(_drsJobID)
	}
	if len(_drsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _drsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_drsNextToken) > 0 {
		input.NextToken = aws.String(_drsNextToken)
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

	var results []*drs.DescribeJobLogItemsOutput
	p := drs.NewDescribeJobLogItemsPaginator(client, input)
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

// Returns a list of Jobs. Use the JobsID and fromDate and toDate filters to limit
// which jobs are returned. The response is sorted by creationDataTime - latest
// date first. Jobs are created by the StartRecovery, TerminateRecoveryInstances
// and StartFailbackLaunch APIs. Jobs are also created by DiagnosticLaunch and
// TerminateDiagnosticInstances, which are APIs available only to *Support* and
// only used in response to relevant support tickets.
func drs_DescribeJobs(cfg aws.Config, client *drs.Client) {
	input := &drs.DescribeJobsInput{}

	if len(_drsFilters) > 0 {
		if err := assignInputField(input, "Filters", _drsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_drsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _drsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_drsNextToken) > 0 {
		input.NextToken = aws.String(_drsNextToken)
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

	var results []*drs.DescribeJobsOutput
	p := drs.NewDescribeJobsPaginator(client, input)
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
func drs_DescribeLaunchConfigurationTemplates(cfg aws.Config, client *drs.Client) {
	input := &drs.DescribeLaunchConfigurationTemplatesInput{}

	if len(_drsLaunchConfigurationTemplateIDs) > 0 {
		input.LaunchConfigurationTemplateIDs = append([]string(nil), _drsLaunchConfigurationTemplateIDs...)
	}
	if len(_drsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _drsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_drsNextToken) > 0 {
		input.NextToken = aws.String(_drsNextToken)
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

	var results []*drs.DescribeLaunchConfigurationTemplatesOutput
	p := drs.NewDescribeLaunchConfigurationTemplatesPaginator(client, input)
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

// Lists all Recovery Instances or multiple Recovery Instances by ID.
func drs_DescribeRecoveryInstances(cfg aws.Config, client *drs.Client) {
	input := &drs.DescribeRecoveryInstancesInput{}

	if len(_drsFilters) > 0 {
		if err := assignInputField(input, "Filters", _drsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_drsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _drsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_drsNextToken) > 0 {
		input.NextToken = aws.String(_drsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeRecoveryInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*drs.DescribeRecoveryInstancesOutput
	p := drs.NewDescribeRecoveryInstancesPaginator(client, input)
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

// Lists all Recovery Snapshots for a single Source Server.
func drs_DescribeRecoverySnapshots(cfg aws.Config, client *drs.Client) {
	input := &drs.DescribeRecoverySnapshotsInput{
		// SourceServerID: *string, // Required
	}

	if len(_drsSourceServerID) > 0 {
		input.SourceServerID = aws.String(_drsSourceServerID)
	}
	if len(_drsFilters) > 0 {
		if err := assignInputField(input, "Filters", _drsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_drsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _drsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_drsNextToken) > 0 {
		input.NextToken = aws.String(_drsNextToken)
	}
	if len(_drsOrder) > 0 {
		if err := assignInputField(input, "Order", _drsOrder); err != nil {
			log.Errorf("invalid --order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeRecoverySnapshots(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*drs.DescribeRecoverySnapshotsOutput
	p := drs.NewDescribeRecoverySnapshotsPaginator(client, input)
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
func drs_DescribeReplicationConfigurationTemplates(cfg aws.Config, client *drs.Client) {
	input := &drs.DescribeReplicationConfigurationTemplatesInput{}

	if len(_drsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _drsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_drsNextToken) > 0 {
		input.NextToken = aws.String(_drsNextToken)
	}
	if len(_drsReplicationConfigurationTemplateIDs) > 0 {
		input.ReplicationConfigurationTemplateIDs = append([]string(nil), _drsReplicationConfigurationTemplateIDs...)
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

	var results []*drs.DescribeReplicationConfigurationTemplatesOutput
	p := drs.NewDescribeReplicationConfigurationTemplatesPaginator(client, input)
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

// Lists all Source Networks or multiple Source Networks filtered by ID.
func drs_DescribeSourceNetworks(cfg aws.Config, client *drs.Client) {
	input := &drs.DescribeSourceNetworksInput{}

	if len(_drsFilters) > 0 {
		if err := assignInputField(input, "Filters", _drsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_drsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _drsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_drsNextToken) > 0 {
		input.NextToken = aws.String(_drsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeSourceNetworks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*drs.DescribeSourceNetworksOutput
	p := drs.NewDescribeSourceNetworksPaginator(client, input)
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

// Lists all Source Servers or multiple Source Servers filtered by ID.
func drs_DescribeSourceServers(cfg aws.Config, client *drs.Client) {
	input := &drs.DescribeSourceServersInput{}

	if len(_drsFilters) > 0 {
		if err := assignInputField(input, "Filters", _drsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_drsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _drsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_drsNextToken) > 0 {
		input.NextToken = aws.String(_drsNextToken)
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

	var results []*drs.DescribeSourceServersOutput
	p := drs.NewDescribeSourceServersPaginator(client, input)
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

// Disconnect a Recovery Instance from Elastic Disaster Recovery. Data replication
// is stopped immediately. All AWS resources created by Elastic Disaster Recovery
// for enabling the replication of the Recovery Instance will be terminated /
// deleted within 90 minutes. If the agent on the Recovery Instance has not been
// prevented from communicating with the Elastic Disaster Recovery service, then it
// will receive a command to uninstall itself (within approximately 10 minutes).
// The following properties of the Recovery Instance will be changed immediately:
// dataReplicationInfo.dataReplicationState will be set to DISCONNECTED; The
// totalStorageBytes property for each of dataReplicationInfo.replicatedDisks will
// be set to zero; dataReplicationInfo.lagDuration and
// dataReplicationInfo.lagDuration will be nullified.
func drs_DisconnectRecoveryInstance(cfg aws.Config, client *drs.Client) {
	input := &drs.DisconnectRecoveryInstanceInput{
		// RecoveryInstanceID: *string, // Required
	}

	if len(_drsRecoveryInstanceID) > 0 {
		input.RecoveryInstanceID = aws.String(_drsRecoveryInstanceID)
	}

	if resp, err := client.DisconnectRecoveryInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disconnects a specific Source Server from Elastic Disaster Recovery. Data
// replication is stopped immediately. All AWS resources created by Elastic
// Disaster Recovery for enabling the replication of the Source Server will be
// terminated / deleted within 90 minutes. You cannot disconnect a Source Server if
// it has a Recovery Instance. If the agent on the Source Server has not been
// prevented from communicating with the Elastic Disaster Recovery service, then it
// will receive a command to uninstall itself (within approximately 10 minutes).
// The following properties of the SourceServer will be changed immediately:
// dataReplicationInfo.dataReplicationState will be set to DISCONNECTED; The
// totalStorageBytes property for each of dataReplicationInfo.replicatedDisks will
// be set to zero; dataReplicationInfo.lagDuration and
// dataReplicationInfo.lagDuration will be nullified.
func drs_DisconnectSourceServer(cfg aws.Config, client *drs.Client) {
	input := &drs.DisconnectSourceServerInput{
		// SourceServerID: *string, // Required
	}

	if len(_drsSourceServerID) > 0 {
		input.SourceServerID = aws.String(_drsSourceServerID)
	}

	if resp, err := client.DisconnectSourceServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Export the Source Network CloudFormation template to an S3 bucket.
func drs_ExportSourceNetworkCfnTemplate(cfg aws.Config, client *drs.Client) {
	input := &drs.ExportSourceNetworkCfnTemplateInput{
		// SourceNetworkID: *string, // Required
	}

	if len(_drsSourceNetworkID) > 0 {
		input.SourceNetworkID = aws.String(_drsSourceNetworkID)
	}

	if resp, err := client.ExportSourceNetworkCfnTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all Failback ReplicationConfigurations, filtered by Recovery Instance ID.
func drs_GetFailbackReplicationConfiguration(cfg aws.Config, client *drs.Client) {
	input := &drs.GetFailbackReplicationConfigurationInput{
		// RecoveryInstanceID: *string, // Required
	}

	if len(_drsRecoveryInstanceID) > 0 {
		input.RecoveryInstanceID = aws.String(_drsRecoveryInstanceID)
	}

	if resp, err := client.GetFailbackReplicationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a LaunchConfiguration, filtered by Source Server IDs.
func drs_GetLaunchConfiguration(cfg aws.Config, client *drs.Client) {
	input := &drs.GetLaunchConfigurationInput{
		// SourceServerID: *string, // Required
	}

	if len(_drsSourceServerID) > 0 {
		input.SourceServerID = aws.String(_drsSourceServerID)
	}

	if resp, err := client.GetLaunchConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a ReplicationConfiguration, filtered by Source Server ID.
func drs_GetReplicationConfiguration(cfg aws.Config, client *drs.Client) {
	input := &drs.GetReplicationConfigurationInput{
		// SourceServerID: *string, // Required
	}

	if len(_drsSourceServerID) > 0 {
		input.SourceServerID = aws.String(_drsSourceServerID)
	}

	if resp, err := client.GetReplicationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initialize Elastic Disaster Recovery.
func drs_InitializeService(cfg aws.Config, client *drs.Client) {
	input := &drs.InitializeServiceInput{}

	if resp, err := client.InitializeService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of source servers on a staging account that are extensible,
// which means that: a. The source server is not already extended into this
// Account. b. The source server on the Account we’re reading from is not an
// extension of another source server.
func drs_ListExtensibleSourceServers(cfg aws.Config, client *drs.Client) {
	input := &drs.ListExtensibleSourceServersInput{
		// StagingAccountID: *string, // Required
	}

	if len(_drsStagingAccountID) > 0 {
		input.StagingAccountID = aws.String(_drsStagingAccountID)
	}
	if len(_drsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _drsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_drsNextToken) > 0 {
		input.NextToken = aws.String(_drsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListExtensibleSourceServers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*drs.ListExtensibleSourceServersOutput
	p := drs.NewListExtensibleSourceServersPaginator(client, input)
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

// Lists resource launch actions.
func drs_ListLaunchActions(cfg aws.Config, client *drs.Client) {
	input := &drs.ListLaunchActionsInput{
		// ResourceId: *string, // Required
	}

	if len(_drsResourceId) > 0 {
		input.ResourceId = aws.String(_drsResourceId)
	}
	if len(_drsFilters) > 0 {
		if err := assignInputField(input, "Filters", _drsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_drsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _drsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_drsNextToken) > 0 {
		input.NextToken = aws.String(_drsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLaunchActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*drs.ListLaunchActionsOutput
	p := drs.NewListLaunchActionsPaginator(client, input)
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

// Returns an array of staging accounts for existing extended source servers.
func drs_ListStagingAccounts(cfg aws.Config, client *drs.Client) {
	input := &drs.ListStagingAccountsInput{}

	if len(_drsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _drsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_drsNextToken) > 0 {
		input.NextToken = aws.String(_drsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStagingAccounts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*drs.ListStagingAccountsOutput
	p := drs.NewListStagingAccountsPaginator(client, input)
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

// List all tags for your Elastic Disaster Recovery resources.
func drs_ListTagsForResource(cfg aws.Config, client *drs.Client) {
	input := &drs.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_drsResourceArn) > 0 {
		input.ResourceArn = aws.String(_drsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Puts a resource launch action.
func drs_PutLaunchAction(cfg aws.Config, client *drs.Client) {
	input := &drs.PutLaunchActionInput{
		// ActionCode: *string, // Required
		// ActionId: *string, // Required
		// ActionVersion: *string, // Required
		// Active: *bool, // Required
		// Category: types.LaunchActionCategory, // Required
		// Description: *string, // Required
		// Name: *string, // Required
		// Optional: *bool, // Required
		// Order: *int32, // Required
		// ResourceId: *string, // Required
	}

	if len(_drsActionCode) > 0 {
		input.ActionCode = aws.String(_drsActionCode)
	}
	if len(_drsActionId) > 0 {
		input.ActionId = aws.String(_drsActionId)
	}
	if len(_drsActionVersion) > 0 {
		input.ActionVersion = aws.String(_drsActionVersion)
	}
	if len(_drsActive) > 0 {
		if err := assignInputField(input, "Active", _drsActive); err != nil {
			log.Errorf("invalid --active: %s", err.Error())
			return
		}
	}
	if len(_drsCategory) > 0 {
		if err := assignInputField(input, "Category", _drsCategory); err != nil {
			log.Errorf("invalid --category: %s", err.Error())
			return
		}
	}
	if len(_drsDescription) > 0 {
		input.Description = aws.String(_drsDescription)
	}
	if len(_drsName) > 0 {
		input.Name = aws.String(_drsName)
	}
	if len(_drsOptional) > 0 {
		if err := assignInputField(input, "Optional", _drsOptional); err != nil {
			log.Errorf("invalid --optional: %s", err.Error())
			return
		}
	}
	if len(_drsOrder) > 0 {
		if err := assignInputField(input, "Order", _drsOrder); err != nil {
			log.Errorf("invalid --order: %s", err.Error())
			return
		}
	}
	if len(_drsResourceId) > 0 {
		input.ResourceId = aws.String(_drsResourceId)
	}
	if len(_drsParameters) > 0 {
		if err := assignInputField(input, "Parameters", _drsParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutLaunchAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// WARNING: RetryDataReplication is deprecated. Causes the data replication
// initiation sequence to begin immediately upon next Handshake for the specified
// Source Server ID, regardless of when the previous initiation started. This
// command will work only if the Source Server is stalled or is in a DISCONNECTED
// or STOPPED state.
//
// Deprecated: WARNING: RetryDataReplication is deprecated
func drs_RetryDataReplication(cfg aws.Config, client *drs.Client) {
	input := &drs.RetryDataReplicationInput{
		// SourceServerID: *string, // Required
	}

	if len(_drsSourceServerID) > 0 {
		input.SourceServerID = aws.String(_drsSourceServerID)
	}

	if resp, err := client.RetryDataReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Start replication to origin / target region - applies only to protected
// instances that originated in EC2. For recovery instances on target region -
// starts replication back to origin region. For failback instances on origin
// region - starts replication to target region to re-protect them.
func drs_ReverseReplication(cfg aws.Config, client *drs.Client) {
	input := &drs.ReverseReplicationInput{
		// RecoveryInstanceID: *string, // Required
	}

	if len(_drsRecoveryInstanceID) > 0 {
		input.RecoveryInstanceID = aws.String(_drsRecoveryInstanceID)
	}

	if resp, err := client.ReverseReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a Job for launching the machine that is being failed back to from the
// specified Recovery Instance. This will run conversion on the failback client and
// will reboot your machine, thus completing the failback process.
func drs_StartFailbackLaunch(cfg aws.Config, client *drs.Client) {
	input := &drs.StartFailbackLaunchInput{
		// RecoveryInstanceIDs: []string, // Required
	}

	if len(_drsRecoveryInstanceIDs) > 0 {
		input.RecoveryInstanceIDs = append([]string(nil), _drsRecoveryInstanceIDs...)
	}
	if len(_drsTags) > 0 {
		if err := assignInputField(input, "Tags", _drsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartFailbackLaunch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Launches Recovery Instances for the specified Source Servers. For each Source
// Server you may choose a point in time snapshot to launch from, or use an on
// demand snapshot.
func drs_StartRecovery(cfg aws.Config, client *drs.Client) {
	input := &drs.StartRecoveryInput{
		// SourceServers: []types.StartRecoveryRequestSourceServer, // Required
	}

	if len(_drsSourceServers) > 0 {
		if err := assignInputField(input, "SourceServers", _drsSourceServers); err != nil {
			log.Errorf("invalid --source-servers: %s", err.Error())
			return
		}
	}
	if len(_drsIsDrill) > 0 {
		if err := assignInputField(input, "IsDrill", _drsIsDrill); err != nil {
			log.Errorf("invalid --is-drill: %s", err.Error())
			return
		}
	}
	if len(_drsTags) > 0 {
		if err := assignInputField(input, "Tags", _drsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartRecovery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts replication for a stopped Source Server. This action would make the
// Source Server protected again and restart billing for it.
func drs_StartReplication(cfg aws.Config, client *drs.Client) {
	input := &drs.StartReplicationInput{
		// SourceServerID: *string, // Required
	}

	if len(_drsSourceServerID) > 0 {
		input.SourceServerID = aws.String(_drsSourceServerID)
	}

	if resp, err := client.StartReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deploy VPC for the specified Source Network and modify launch templates to use
// this network. The VPC will be deployed using a dedicated CloudFormation stack.
func drs_StartSourceNetworkRecovery(cfg aws.Config, client *drs.Client) {
	input := &drs.StartSourceNetworkRecoveryInput{
		// SourceNetworks: []types.StartSourceNetworkRecoveryRequestNetworkEntry, // Required
	}

	if len(_drsSourceNetworks) > 0 {
		if err := assignInputField(input, "SourceNetworks", _drsSourceNetworks); err != nil {
			log.Errorf("invalid --source-networks: %s", err.Error())
			return
		}
	}
	if len(_drsDeployAsNew) > 0 {
		if err := assignInputField(input, "DeployAsNew", _drsDeployAsNew); err != nil {
			log.Errorf("invalid --deploy-as-new: %s", err.Error())
			return
		}
	}
	if len(_drsTags) > 0 {
		if err := assignInputField(input, "Tags", _drsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartSourceNetworkRecovery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts replication for a Source Network. This action would make the Source
// Network protected.
func drs_StartSourceNetworkReplication(cfg aws.Config, client *drs.Client) {
	input := &drs.StartSourceNetworkReplicationInput{
		// SourceNetworkID: *string, // Required
	}

	if len(_drsSourceNetworkID) > 0 {
		input.SourceNetworkID = aws.String(_drsSourceNetworkID)
	}

	if resp, err := client.StartSourceNetworkReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the failback process for a specified Recovery Instance. This changes the
// Failback State of the Recovery Instance back to FAILBACK_NOT_STARTED.
func drs_StopFailback(cfg aws.Config, client *drs.Client) {
	input := &drs.StopFailbackInput{
		// RecoveryInstanceID: *string, // Required
	}

	if len(_drsRecoveryInstanceID) > 0 {
		input.RecoveryInstanceID = aws.String(_drsRecoveryInstanceID)
	}

	if resp, err := client.StopFailback(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops replication for a Source Server. This action would make the Source Server
// unprotected, delete its existing snapshots and stop billing for it.
func drs_StopReplication(cfg aws.Config, client *drs.Client) {
	input := &drs.StopReplicationInput{
		// SourceServerID: *string, // Required
	}

	if len(_drsSourceServerID) > 0 {
		input.SourceServerID = aws.String(_drsSourceServerID)
	}

	if resp, err := client.StopReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops replication for a Source Network. This action would make the Source
// Network unprotected.
func drs_StopSourceNetworkReplication(cfg aws.Config, client *drs.Client) {
	input := &drs.StopSourceNetworkReplicationInput{
		// SourceNetworkID: *string, // Required
	}

	if len(_drsSourceNetworkID) > 0 {
		input.SourceNetworkID = aws.String(_drsSourceNetworkID)
	}

	if resp, err := client.StopSourceNetworkReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or overwrites only the specified tags for the specified Elastic Disaster
// Recovery resource or resources. When you specify an existing tag key, the value
// is overwritten with the new value. Each resource can have a maximum of 50 tags.
// Each tag consists of a key and optional value.
func drs_TagResource(cfg aws.Config, client *drs.Client) {
	input := &drs.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_drsResourceArn) > 0 {
		input.ResourceArn = aws.String(_drsResourceArn)
	}
	if len(_drsTags) > 0 {
		if err := assignInputField(input, "Tags", _drsTags); err != nil {
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

// Initiates a Job for terminating the EC2 resources associated with the specified
// Recovery Instances, and then will delete the Recovery Instances from the Elastic
// Disaster Recovery service.
func drs_TerminateRecoveryInstances(cfg aws.Config, client *drs.Client) {
	input := &drs.TerminateRecoveryInstancesInput{
		// RecoveryInstanceIDs: []string, // Required
	}

	if len(_drsRecoveryInstanceIDs) > 0 {
		input.RecoveryInstanceIDs = append([]string(nil), _drsRecoveryInstanceIDs...)
	}

	if resp, err := client.TerminateRecoveryInstances(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified set of tags from the specified set of Elastic Disaster
// Recovery resources.
func drs_UntagResource(cfg aws.Config, client *drs.Client) {
	input := &drs.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_drsResourceArn) > 0 {
		input.ResourceArn = aws.String(_drsResourceArn)
	}
	if len(_drsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _drsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows you to update the failback replication configuration of a Recovery
// Instance by ID.
func drs_UpdateFailbackReplicationConfiguration(cfg aws.Config, client *drs.Client) {
	input := &drs.UpdateFailbackReplicationConfigurationInput{
		// RecoveryInstanceID: *string, // Required
	}

	if len(_drsRecoveryInstanceID) > 0 {
		input.RecoveryInstanceID = aws.String(_drsRecoveryInstanceID)
	}
	if len(_drsBandwidthThrottling) > 0 {
		if err := assignInputField(input, "BandwidthThrottling", _drsBandwidthThrottling); err != nil {
			log.Errorf("invalid --bandwidth-throttling: %s", err.Error())
			return
		}
	}
	if len(_drsName) > 0 {
		input.Name = aws.String(_drsName)
	}
	if len(_drsUsePrivateIP) > 0 {
		if err := assignInputField(input, "UsePrivateIP", _drsUsePrivateIP); err != nil {
			log.Errorf("invalid --use-private-ip: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFailbackReplicationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a LaunchConfiguration by Source Server ID.
func drs_UpdateLaunchConfiguration(cfg aws.Config, client *drs.Client) {
	input := &drs.UpdateLaunchConfigurationInput{
		// SourceServerID: *string, // Required
	}

	if len(_drsSourceServerID) > 0 {
		input.SourceServerID = aws.String(_drsSourceServerID)
	}
	if len(_drsCopyPrivateIp) > 0 {
		if err := assignInputField(input, "CopyPrivateIp", _drsCopyPrivateIp); err != nil {
			log.Errorf("invalid --copy-private-ip: %s", err.Error())
			return
		}
	}
	if len(_drsCopyTags) > 0 {
		if err := assignInputField(input, "CopyTags", _drsCopyTags); err != nil {
			log.Errorf("invalid --copy-tags: %s", err.Error())
			return
		}
	}
	if len(_drsLaunchDisposition) > 0 {
		if err := assignInputField(input, "LaunchDisposition", _drsLaunchDisposition); err != nil {
			log.Errorf("invalid --launch-disposition: %s", err.Error())
			return
		}
	}
	if len(_drsLaunchIntoInstanceProperties) > 0 {
		if err := assignInputField(input, "LaunchIntoInstanceProperties", _drsLaunchIntoInstanceProperties); err != nil {
			log.Errorf("invalid --launch-into-instance-properties: %s", err.Error())
			return
		}
	}
	if len(_drsLicensing) > 0 {
		if err := assignInputField(input, "Licensing", _drsLicensing); err != nil {
			log.Errorf("invalid --licensing: %s", err.Error())
			return
		}
	}
	if len(_drsName) > 0 {
		input.Name = aws.String(_drsName)
	}
	if len(_drsPostLaunchEnabled) > 0 {
		if err := assignInputField(input, "PostLaunchEnabled", _drsPostLaunchEnabled); err != nil {
			log.Errorf("invalid --post-launch-enabled: %s", err.Error())
			return
		}
	}
	if len(_drsTargetInstanceTypeRightSizingMethod) > 0 {
		if err := assignInputField(input, "TargetInstanceTypeRightSizingMethod", _drsTargetInstanceTypeRightSizingMethod); err != nil {
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
func drs_UpdateLaunchConfigurationTemplate(cfg aws.Config, client *drs.Client) {
	input := &drs.UpdateLaunchConfigurationTemplateInput{
		// LaunchConfigurationTemplateID: *string, // Required
	}

	if len(_drsLaunchConfigurationTemplateID) > 0 {
		input.LaunchConfigurationTemplateID = aws.String(_drsLaunchConfigurationTemplateID)
	}
	if len(_drsCopyPrivateIp) > 0 {
		if err := assignInputField(input, "CopyPrivateIp", _drsCopyPrivateIp); err != nil {
			log.Errorf("invalid --copy-private-ip: %s", err.Error())
			return
		}
	}
	if len(_drsCopyTags) > 0 {
		if err := assignInputField(input, "CopyTags", _drsCopyTags); err != nil {
			log.Errorf("invalid --copy-tags: %s", err.Error())
			return
		}
	}
	if len(_drsExportBucketArn) > 0 {
		input.ExportBucketArn = aws.String(_drsExportBucketArn)
	}
	if len(_drsLaunchDisposition) > 0 {
		if err := assignInputField(input, "LaunchDisposition", _drsLaunchDisposition); err != nil {
			log.Errorf("invalid --launch-disposition: %s", err.Error())
			return
		}
	}
	if len(_drsLaunchIntoSourceInstance) > 0 {
		if err := assignInputField(input, "LaunchIntoSourceInstance", _drsLaunchIntoSourceInstance); err != nil {
			log.Errorf("invalid --launch-into-source-instance: %s", err.Error())
			return
		}
	}
	if len(_drsLicensing) > 0 {
		if err := assignInputField(input, "Licensing", _drsLicensing); err != nil {
			log.Errorf("invalid --licensing: %s", err.Error())
			return
		}
	}
	if len(_drsPostLaunchEnabled) > 0 {
		if err := assignInputField(input, "PostLaunchEnabled", _drsPostLaunchEnabled); err != nil {
			log.Errorf("invalid --post-launch-enabled: %s", err.Error())
			return
		}
	}
	if len(_drsTargetInstanceTypeRightSizingMethod) > 0 {
		if err := assignInputField(input, "TargetInstanceTypeRightSizingMethod", _drsTargetInstanceTypeRightSizingMethod); err != nil {
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

// Allows you to update a ReplicationConfiguration by Source Server ID.
func drs_UpdateReplicationConfiguration(cfg aws.Config, client *drs.Client) {
	input := &drs.UpdateReplicationConfigurationInput{
		// SourceServerID: *string, // Required
	}

	if len(_drsSourceServerID) > 0 {
		input.SourceServerID = aws.String(_drsSourceServerID)
	}
	if len(_drsAssociateDefaultSecurityGroup) > 0 {
		if err := assignInputField(input, "AssociateDefaultSecurityGroup", _drsAssociateDefaultSecurityGroup); err != nil {
			log.Errorf("invalid --associate-default-security-group: %s", err.Error())
			return
		}
	}
	if len(_drsAutoReplicateNewDisks) > 0 {
		if err := assignInputField(input, "AutoReplicateNewDisks", _drsAutoReplicateNewDisks); err != nil {
			log.Errorf("invalid --auto-replicate-new-disks: %s", err.Error())
			return
		}
	}
	if len(_drsBandwidthThrottling) > 0 {
		if err := assignInputField(input, "BandwidthThrottling", _drsBandwidthThrottling); err != nil {
			log.Errorf("invalid --bandwidth-throttling: %s", err.Error())
			return
		}
	}
	if len(_drsCreatePublicIP) > 0 {
		if err := assignInputField(input, "CreatePublicIP", _drsCreatePublicIP); err != nil {
			log.Errorf("invalid --create-public-ip: %s", err.Error())
			return
		}
	}
	if len(_drsDataPlaneRouting) > 0 {
		if err := assignInputField(input, "DataPlaneRouting", _drsDataPlaneRouting); err != nil {
			log.Errorf("invalid --data-plane-routing: %s", err.Error())
			return
		}
	}
	if len(_drsDefaultLargeStagingDiskType) > 0 {
		if err := assignInputField(input, "DefaultLargeStagingDiskType", _drsDefaultLargeStagingDiskType); err != nil {
			log.Errorf("invalid --default-large-staging-disk-type: %s", err.Error())
			return
		}
	}
	if len(_drsEbsEncryption) > 0 {
		if err := assignInputField(input, "EbsEncryption", _drsEbsEncryption); err != nil {
			log.Errorf("invalid --ebs-encryption: %s", err.Error())
			return
		}
	}
	if len(_drsEbsEncryptionKeyArn) > 0 {
		input.EbsEncryptionKeyArn = aws.String(_drsEbsEncryptionKeyArn)
	}
	if len(_drsName) > 0 {
		input.Name = aws.String(_drsName)
	}
	if len(_drsPitPolicy) > 0 {
		if err := assignInputField(input, "PitPolicy", _drsPitPolicy); err != nil {
			log.Errorf("invalid --pit-policy: %s", err.Error())
			return
		}
	}
	if len(_drsReplicatedDisks) > 0 {
		if err := assignInputField(input, "ReplicatedDisks", _drsReplicatedDisks); err != nil {
			log.Errorf("invalid --replicated-disks: %s", err.Error())
			return
		}
	}
	if len(_drsReplicationServerInstanceType) > 0 {
		input.ReplicationServerInstanceType = aws.String(_drsReplicationServerInstanceType)
	}
	if len(_drsReplicationServersSecurityGroupsIDs) > 0 {
		input.ReplicationServersSecurityGroupsIDs = append([]string(nil), _drsReplicationServersSecurityGroupsIDs...)
	}
	if len(_drsStagingAreaSubnetId) > 0 {
		input.StagingAreaSubnetId = aws.String(_drsStagingAreaSubnetId)
	}
	if len(_drsStagingAreaTags) > 0 {
		if err := assignInputField(input, "StagingAreaTags", _drsStagingAreaTags); err != nil {
			log.Errorf("invalid --staging-area-tags: %s", err.Error())
			return
		}
	}
	if len(_drsUseDedicatedReplicationServer) > 0 {
		if err := assignInputField(input, "UseDedicatedReplicationServer", _drsUseDedicatedReplicationServer); err != nil {
			log.Errorf("invalid --use-dedicated-replication-server: %s", err.Error())
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

// Updates a ReplicationConfigurationTemplate by ID.
func drs_UpdateReplicationConfigurationTemplate(cfg aws.Config, client *drs.Client) {
	input := &drs.UpdateReplicationConfigurationTemplateInput{
		// ReplicationConfigurationTemplateID: *string, // Required
	}

	if len(_drsReplicationConfigurationTemplateID) > 0 {
		input.ReplicationConfigurationTemplateID = aws.String(_drsReplicationConfigurationTemplateID)
	}
	if len(_drsArn) > 0 {
		input.Arn = aws.String(_drsArn)
	}
	if len(_drsAssociateDefaultSecurityGroup) > 0 {
		if err := assignInputField(input, "AssociateDefaultSecurityGroup", _drsAssociateDefaultSecurityGroup); err != nil {
			log.Errorf("invalid --associate-default-security-group: %s", err.Error())
			return
		}
	}
	if len(_drsAutoReplicateNewDisks) > 0 {
		if err := assignInputField(input, "AutoReplicateNewDisks", _drsAutoReplicateNewDisks); err != nil {
			log.Errorf("invalid --auto-replicate-new-disks: %s", err.Error())
			return
		}
	}
	if len(_drsBandwidthThrottling) > 0 {
		if err := assignInputField(input, "BandwidthThrottling", _drsBandwidthThrottling); err != nil {
			log.Errorf("invalid --bandwidth-throttling: %s", err.Error())
			return
		}
	}
	if len(_drsCreatePublicIP) > 0 {
		if err := assignInputField(input, "CreatePublicIP", _drsCreatePublicIP); err != nil {
			log.Errorf("invalid --create-public-ip: %s", err.Error())
			return
		}
	}
	if len(_drsDataPlaneRouting) > 0 {
		if err := assignInputField(input, "DataPlaneRouting", _drsDataPlaneRouting); err != nil {
			log.Errorf("invalid --data-plane-routing: %s", err.Error())
			return
		}
	}
	if len(_drsDefaultLargeStagingDiskType) > 0 {
		if err := assignInputField(input, "DefaultLargeStagingDiskType", _drsDefaultLargeStagingDiskType); err != nil {
			log.Errorf("invalid --default-large-staging-disk-type: %s", err.Error())
			return
		}
	}
	if len(_drsEbsEncryption) > 0 {
		if err := assignInputField(input, "EbsEncryption", _drsEbsEncryption); err != nil {
			log.Errorf("invalid --ebs-encryption: %s", err.Error())
			return
		}
	}
	if len(_drsEbsEncryptionKeyArn) > 0 {
		input.EbsEncryptionKeyArn = aws.String(_drsEbsEncryptionKeyArn)
	}
	if len(_drsPitPolicy) > 0 {
		if err := assignInputField(input, "PitPolicy", _drsPitPolicy); err != nil {
			log.Errorf("invalid --pit-policy: %s", err.Error())
			return
		}
	}
	if len(_drsReplicationServerInstanceType) > 0 {
		input.ReplicationServerInstanceType = aws.String(_drsReplicationServerInstanceType)
	}
	if len(_drsReplicationServersSecurityGroupsIDs) > 0 {
		input.ReplicationServersSecurityGroupsIDs = append([]string(nil), _drsReplicationServersSecurityGroupsIDs...)
	}
	if len(_drsStagingAreaSubnetId) > 0 {
		input.StagingAreaSubnetId = aws.String(_drsStagingAreaSubnetId)
	}
	if len(_drsStagingAreaTags) > 0 {
		if err := assignInputField(input, "StagingAreaTags", _drsStagingAreaTags); err != nil {
			log.Errorf("invalid --staging-area-tags: %s", err.Error())
			return
		}
	}
	if len(_drsUseDedicatedReplicationServer) > 0 {
		if err := assignInputField(input, "UseDedicatedReplicationServer", _drsUseDedicatedReplicationServer); err != nil {
			log.Errorf("invalid --use-dedicated-replication-server: %s", err.Error())
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

func init() {
	_rootCmd.AddCommand(_drsCmd)
	_drsCmd.Flags().SortFlags = false

	_drsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_drsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_drsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_drsCmd.Flags().StringVarP(&_drsActionCode, "action-code", "", "", "Action Code")
	_drsCmd.Flags().StringVarP(&_drsActionId, "action-id", "", "", "Action ID")
	_drsCmd.Flags().StringVarP(&_drsActionVersion, "action-version", "", "", "Action Version")
	_drsCmd.Flags().StringVarP(&_drsActive, "active", "", "", "Active")
	_drsCmd.Flags().StringVarP(&_drsArn, "arn", "", "", "ARN")
	_drsCmd.Flags().StringVarP(&_drsAssociateDefaultSecurityGroup, "associate-default-security-group", "", "", "Associate Default Security Group")
	_drsCmd.Flags().StringVarP(&_drsAutoReplicateNewDisks, "auto-replicate-new-disks", "", "", "Auto Replicate New Disks")
	_drsCmd.Flags().StringVarP(&_drsBandwidthThrottling, "bandwidth-throttling", "", "", "Bandwidth Throttling")
	_drsCmd.Flags().StringVarP(&_drsCategory, "category", "", "", "Category")
	_drsCmd.Flags().StringVarP(&_drsCfnStackName, "cfn-stack-name", "", "", "Cfn Stack Name")
	_drsCmd.Flags().StringVarP(&_drsCopyPrivateIp, "copy-private-ip", "", "", "Copy Private IP")
	_drsCmd.Flags().StringVarP(&_drsCopyTags, "copy-tags", "", "", "Copy Tags")
	_drsCmd.Flags().StringVarP(&_drsCreatePublicIP, "create-public-ip", "", "", "Create Public IP")
	_drsCmd.Flags().StringVarP(&_drsDataPlaneRouting, "data-plane-routing", "", "", "Data Plane Routing")
	_drsCmd.Flags().StringVarP(&_drsDefaultLargeStagingDiskType, "default-large-staging-disk-type", "", "", "Default Large Staging Disk Type")
	_drsCmd.Flags().StringVarP(&_drsDeployAsNew, "deploy-as-new", "", "", "Deploy As New")
	_drsCmd.Flags().StringVarP(&_drsDescription, "description", "", "", "Description")
	_drsCmd.Flags().StringVarP(&_drsEbsEncryption, "ebs-encryption", "", "", "Ebs Encryption")
	_drsCmd.Flags().StringVarP(&_drsEbsEncryptionKeyArn, "ebs-encryption-key-arn", "", "", "Ebs Encryption Key ARN")
	_drsCmd.Flags().StringVarP(&_drsExportBucketArn, "export-bucket-arn", "", "", "Export Bucket ARN")
	_drsCmd.Flags().StringVarP(&_drsFilters, "filters", "", "", "Filters")
	_drsCmd.Flags().StringVarP(&_drsIsDrill, "is-drill", "", "", "Is Drill")
	_drsCmd.Flags().StringVarP(&_drsJobID, "job-id", "", "", "Job ID")
	_drsCmd.Flags().StringVarP(&_drsLaunchConfigurationTemplateID, "launch-configuration-template-id", "", "", "Launch Configuration Template ID")
	_drsCmd.Flags().StringSliceVarP(&_drsLaunchConfigurationTemplateIDs, "launch-configuration-template-ids", "", nil, "Launch Configuration Template Ids")
	_drsCmd.Flags().StringVarP(&_drsLaunchDisposition, "launch-disposition", "", "", "Launch Disposition")
	_drsCmd.Flags().StringVarP(&_drsLaunchIntoInstanceProperties, "launch-into-instance-properties", "", "", "Launch Into Instance Properties")
	_drsCmd.Flags().StringVarP(&_drsLaunchIntoSourceInstance, "launch-into-source-instance", "", "", "Launch Into Source Instance")
	_drsCmd.Flags().StringVarP(&_drsLicensing, "licensing", "", "", "Licensing")
	_drsCmd.Flags().StringVarP(&_drsMaxResults, "max-results", "", "", "Max Results")
	_drsCmd.Flags().StringVarP(&_drsName, "name", "", "", "Name")
	_drsCmd.Flags().StringVarP(&_drsNextToken, "next-token", "", "", "Next Token")
	_drsCmd.Flags().StringVarP(&_drsOptional, "optional", "", "", "Optional")
	_drsCmd.Flags().StringVarP(&_drsOrder, "order", "", "", "Order")
	_drsCmd.Flags().StringVarP(&_drsOriginAccountID, "origin-account-id", "", "", "Origin Account ID")
	_drsCmd.Flags().StringVarP(&_drsOriginRegion, "origin-region", "", "", "Origin Region")
	_drsCmd.Flags().StringVarP(&_drsParameters, "parameters", "", "", "Parameters")
	_drsCmd.Flags().StringVarP(&_drsPitPolicy, "pit-policy", "", "", "Pit Policy")
	_drsCmd.Flags().StringVarP(&_drsPostLaunchEnabled, "post-launch-enabled", "", "", "Post Launch Enabled")
	_drsCmd.Flags().StringVarP(&_drsRecoveryInstanceID, "recovery-instance-id", "", "", "Recovery Instance ID")
	_drsCmd.Flags().StringSliceVarP(&_drsRecoveryInstanceIDs, "recovery-instance-ids", "", nil, "Recovery Instance Ids")
	_drsCmd.Flags().StringVarP(&_drsReplicatedDisks, "replicated-disks", "", "", "Replicated Disks")
	_drsCmd.Flags().StringVarP(&_drsReplicationConfigurationTemplateID, "replication-configuration-template-id", "", "", "Replication Configuration Template ID")
	_drsCmd.Flags().StringSliceVarP(&_drsReplicationConfigurationTemplateIDs, "replication-configuration-template-ids", "", nil, "Replication Configuration Template Ids")
	_drsCmd.Flags().StringVarP(&_drsReplicationServerInstanceType, "replication-server-instance-type", "", "", "Replication Server Instance Type")
	_drsCmd.Flags().StringSliceVarP(&_drsReplicationServersSecurityGroupsIDs, "replication-servers-security-groups-ids", "", nil, "Replication Servers Security Groups Ids")
	_drsCmd.Flags().StringVarP(&_drsResourceArn, "resource-arn", "", "", "Resource ARN")
	_drsCmd.Flags().StringVarP(&_drsResourceId, "resource-id", "", "", "Resource ID")
	_drsCmd.Flags().StringVarP(&_drsSourceNetworkID, "source-network-id", "", "", "Source Network ID")
	_drsCmd.Flags().StringVarP(&_drsSourceNetworks, "source-networks", "", "", "Source Networks")
	_drsCmd.Flags().StringVarP(&_drsSourceServerArn, "source-server-arn", "", "", "Source Server ARN")
	_drsCmd.Flags().StringVarP(&_drsSourceServerID, "source-server-id", "", "", "Source Server ID")
	_drsCmd.Flags().StringVarP(&_drsSourceServers, "source-servers", "", "", "Source Servers")
	_drsCmd.Flags().StringVarP(&_drsStagingAccountID, "staging-account-id", "", "", "Staging Account ID")
	_drsCmd.Flags().StringVarP(&_drsStagingAreaSubnetId, "staging-area-subnet-id", "", "", "Staging Area Subnet ID")
	_drsCmd.Flags().StringVarP(&_drsStagingAreaTags, "staging-area-tags", "", "", "Staging Area Tags")
	_drsCmd.Flags().StringSliceVarP(&_drsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_drsCmd.Flags().StringVarP(&_drsTags, "tags", "", "", "Tags")
	_drsCmd.Flags().StringVarP(&_drsTargetInstanceTypeRightSizingMethod, "target-instance-type-right-sizing-method", "", "", "Target Instance Type Right Sizing Method")
	_drsCmd.Flags().StringVarP(&_drsUseDedicatedReplicationServer, "use-dedicated-replication-server", "", "", "Use Dedicated Replication Server")
	_drsCmd.Flags().StringVarP(&_drsUsePrivateIP, "use-private-ip", "", "", "Use Private IP")
	_drsCmd.Flags().StringVarP(&_drsVpcID, "vpc-id", "", "", "VPC ID")

	_drsCmd.Flags().BoolVarP(&_drsAssociateSourceNetworkStack, "associate-source-network-stack", "", false, "Associate Source Network Stack")
	_drsCmd.Flags().BoolVarP(&_drsCreateExtendedSourceServer, "create-extended-source-server", "", false, "Create Extended Source Server")
	_drsCmd.Flags().BoolVarP(&_drsCreateLaunchConfigurationTemplate, "create-launch-configuration-template", "", false, "Create Launch Configuration Template")
	_drsCmd.Flags().BoolVarP(&_drsCreateReplicationConfigurationTemplate, "create-replication-configuration-template", "", false, "Create Replication Configuration Template")
	_drsCmd.Flags().BoolVarP(&_drsCreateSourceNetwork, "create-source-network", "", false, "Create Source Network")
	_drsCmd.Flags().BoolVarP(&_drsDeleteJob, "delete-job", "", false, "Delete Job")
	_drsCmd.Flags().BoolVarP(&_drsDeleteLaunchAction, "delete-launch-action", "", false, "Delete Launch Action")
	_drsCmd.Flags().BoolVarP(&_drsDeleteLaunchConfigurationTemplate, "delete-launch-configuration-template", "", false, "Delete Launch Configuration Template")
	_drsCmd.Flags().BoolVarP(&_drsDeleteRecoveryInstance, "delete-recovery-instance", "", false, "Delete Recovery Instance")
	_drsCmd.Flags().BoolVarP(&_drsDeleteReplicationConfigurationTemplate, "delete-replication-configuration-template", "", false, "Delete Replication Configuration Template")
	_drsCmd.Flags().BoolVarP(&_drsDeleteSourceNetwork, "delete-source-network", "", false, "Delete Source Network")
	_drsCmd.Flags().BoolVarP(&_drsDeleteSourceServer, "delete-source-server", "", false, "Delete Source Server")
	_drsCmd.Flags().BoolVarP(&_drsDescribeJobLogItems, "describe-job-log-items", "", false, "Describe Job Log Items")
	_drsCmd.Flags().BoolVarP(&_drsDescribeJobs, "describe-jobs", "", false, "Describe Jobs")
	_drsCmd.Flags().BoolVarP(&_drsDescribeLaunchConfigurationTemplates, "describe-launch-configuration-templates", "", false, "Describe Launch Configuration Templates")
	_drsCmd.Flags().BoolVarP(&_drsDescribeRecoveryInstances, "describe-recovery-instances", "", false, "Describe Recovery Instances")
	_drsCmd.Flags().BoolVarP(&_drsDescribeRecoverySnapshots, "describe-recovery-snapshots", "", false, "Describe Recovery Snapshots")
	_drsCmd.Flags().BoolVarP(&_drsDescribeReplicationConfigurationTemplates, "describe-replication-configuration-templates", "", false, "Describe Replication Configuration Templates")
	_drsCmd.Flags().BoolVarP(&_drsDescribeSourceNetworks, "describe-source-networks", "", false, "Describe Source Networks")
	_drsCmd.Flags().BoolVarP(&_drsDescribeSourceServers, "describe-source-servers", "", false, "Describe Source Servers")
	_drsCmd.Flags().BoolVarP(&_drsDisconnectRecoveryInstance, "disconnect-recovery-instance", "", false, "Disconnect Recovery Instance")
	_drsCmd.Flags().BoolVarP(&_drsDisconnectSourceServer, "disconnect-source-server", "", false, "Disconnect Source Server")
	_drsCmd.Flags().BoolVarP(&_drsExportSourceNetworkCfnTemplate, "export-source-network-cfn-template", "", false, "Export Source Network Cfn Template")
	_drsCmd.Flags().BoolVarP(&_drsGetFailbackReplicationConfiguration, "get-failback-replication-configuration", "", false, "Get Failback Replication Configuration")
	_drsCmd.Flags().BoolVarP(&_drsGetLaunchConfiguration, "get-launch-configuration", "", false, "Get Launch Configuration")
	_drsCmd.Flags().BoolVarP(&_drsGetReplicationConfiguration, "get-replication-configuration", "", false, "Get Replication Configuration")
	_drsCmd.Flags().BoolVarP(&_drsInitializeService, "initialize-service", "", false, "Initialize Service")
	_drsCmd.Flags().BoolVarP(&_drsListExtensibleSourceServers, "list-extensible-source-servers", "", false, "List Extensible Source Servers")
	_drsCmd.Flags().BoolVarP(&_drsListLaunchActions, "list-launch-actions", "", false, "List Launch Actions")
	_drsCmd.Flags().BoolVarP(&_drsListStagingAccounts, "list-staging-accounts", "", false, "List Staging Accounts")
	_drsCmd.Flags().BoolVarP(&_drsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_drsCmd.Flags().BoolVarP(&_drsPutLaunchAction, "put-launch-action", "", false, "Put Launch Action")
	_drsCmd.Flags().BoolVarP(&_drsRetryDataReplication, "retry-data-replication", "", false, "Retry Data Replication")
	_drsCmd.Flags().BoolVarP(&_drsReverseReplication, "reverse-replication", "", false, "Reverse Replication")
	_drsCmd.Flags().BoolVarP(&_drsStartFailbackLaunch, "start-failback-launch", "", false, "Start Failback Launch")
	_drsCmd.Flags().BoolVarP(&_drsStartRecovery, "start-recovery", "", false, "Start Recovery")
	_drsCmd.Flags().BoolVarP(&_drsStartReplication, "start-replication", "", false, "Start Replication")
	_drsCmd.Flags().BoolVarP(&_drsStartSourceNetworkRecovery, "start-source-network-recovery", "", false, "Start Source Network Recovery")
	_drsCmd.Flags().BoolVarP(&_drsStartSourceNetworkReplication, "start-source-network-replication", "", false, "Start Source Network Replication")
	_drsCmd.Flags().BoolVarP(&_drsStopFailback, "stop-failback", "", false, "Stop Failback")
	_drsCmd.Flags().BoolVarP(&_drsStopReplication, "stop-replication", "", false, "Stop Replication")
	_drsCmd.Flags().BoolVarP(&_drsStopSourceNetworkReplication, "stop-source-network-replication", "", false, "Stop Source Network Replication")
	_drsCmd.Flags().BoolVarP(&_drsTagResource, "tag-resource", "", false, "Tag Resource")
	_drsCmd.Flags().BoolVarP(&_drsTerminateRecoveryInstances, "terminate-recovery-instances", "", false, "Terminate Recovery Instances")
	_drsCmd.Flags().BoolVarP(&_drsUntagResource, "untag-resource", "", false, "Untag Resource")
	_drsCmd.Flags().BoolVarP(&_drsUpdateFailbackReplicationConfiguration, "update-failback-replication-configuration", "", false, "Update Failback Replication Configuration")
	_drsCmd.Flags().BoolVarP(&_drsUpdateLaunchConfiguration, "update-launch-configuration", "", false, "Update Launch Configuration")
	_drsCmd.Flags().BoolVarP(&_drsUpdateLaunchConfigurationTemplate, "update-launch-configuration-template", "", false, "Update Launch Configuration Template")
	_drsCmd.Flags().BoolVarP(&_drsUpdateReplicationConfiguration, "update-replication-configuration", "", false, "Update Replication Configuration")
	_drsCmd.Flags().BoolVarP(&_drsUpdateReplicationConfigurationTemplate, "update-replication-configuration-template", "", false, "Update Replication Configuration Template")

}
