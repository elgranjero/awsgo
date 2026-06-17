package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/connectcampaignsv2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// connectcampaignsv2Cmd represents the connectcampaignsv2 command
var _connectcampaignsv2Cmd = &cobra.Command{
	Use:   "connectcampaignsv2",
	Short: "AWS connectcampaignsv2 CLI",
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
		client := connectcampaignsv2.NewFromConfig(cfg)
		if _connectcampaignsv2CreateCampaign {
			connectcampaignsv2_CreateCampaign(cfg, client)
			return
		}
		if _connectcampaignsv2DeleteCampaign {
			connectcampaignsv2_DeleteCampaign(cfg, client)
			return
		}
		if _connectcampaignsv2DeleteCampaignChannelSubtypeConfig {
			connectcampaignsv2_DeleteCampaignChannelSubtypeConfig(cfg, client)
			return
		}
		if _connectcampaignsv2DeleteCampaignCommunicationLimits {
			connectcampaignsv2_DeleteCampaignCommunicationLimits(cfg, client)
			return
		}
		if _connectcampaignsv2DeleteCampaignCommunicationTime {
			connectcampaignsv2_DeleteCampaignCommunicationTime(cfg, client)
			return
		}
		if _connectcampaignsv2DeleteConnectInstanceConfig {
			connectcampaignsv2_DeleteConnectInstanceConfig(cfg, client)
			return
		}
		if _connectcampaignsv2DeleteConnectInstanceIntegration {
			connectcampaignsv2_DeleteConnectInstanceIntegration(cfg, client)
			return
		}
		if _connectcampaignsv2DeleteInstanceOnboardingJob {
			connectcampaignsv2_DeleteInstanceOnboardingJob(cfg, client)
			return
		}
		if _connectcampaignsv2DescribeCampaign {
			connectcampaignsv2_DescribeCampaign(cfg, client)
			return
		}
		if _connectcampaignsv2GetCampaignState {
			connectcampaignsv2_GetCampaignState(cfg, client)
			return
		}
		if _connectcampaignsv2GetCampaignStateBatch {
			connectcampaignsv2_GetCampaignStateBatch(cfg, client)
			return
		}
		if _connectcampaignsv2GetConnectInstanceConfig {
			connectcampaignsv2_GetConnectInstanceConfig(cfg, client)
			return
		}
		if _connectcampaignsv2GetInstanceCommunicationLimits {
			connectcampaignsv2_GetInstanceCommunicationLimits(cfg, client)
			return
		}
		if _connectcampaignsv2GetInstanceOnboardingJobStatus {
			connectcampaignsv2_GetInstanceOnboardingJobStatus(cfg, client)
			return
		}
		if _connectcampaignsv2ListCampaigns {
			connectcampaignsv2_ListCampaigns(cfg, client)
			return
		}
		if _connectcampaignsv2ListConnectInstanceIntegrations {
			connectcampaignsv2_ListConnectInstanceIntegrations(cfg, client)
			return
		}
		if _connectcampaignsv2ListTagsForResource {
			connectcampaignsv2_ListTagsForResource(cfg, client)
			return
		}
		if _connectcampaignsv2PauseCampaign {
			connectcampaignsv2_PauseCampaign(cfg, client)
			return
		}
		if _connectcampaignsv2PutConnectInstanceIntegration {
			connectcampaignsv2_PutConnectInstanceIntegration(cfg, client)
			return
		}
		if _connectcampaignsv2PutInstanceCommunicationLimits {
			connectcampaignsv2_PutInstanceCommunicationLimits(cfg, client)
			return
		}
		if _connectcampaignsv2PutOutboundRequestBatch {
			connectcampaignsv2_PutOutboundRequestBatch(cfg, client)
			return
		}
		if _connectcampaignsv2PutProfileOutboundRequestBatch {
			connectcampaignsv2_PutProfileOutboundRequestBatch(cfg, client)
			return
		}
		if _connectcampaignsv2ResumeCampaign {
			connectcampaignsv2_ResumeCampaign(cfg, client)
			return
		}
		if _connectcampaignsv2StartCampaign {
			connectcampaignsv2_StartCampaign(cfg, client)
			return
		}
		if _connectcampaignsv2StartInstanceOnboardingJob {
			connectcampaignsv2_StartInstanceOnboardingJob(cfg, client)
			return
		}
		if _connectcampaignsv2StopCampaign {
			connectcampaignsv2_StopCampaign(cfg, client)
			return
		}
		if _connectcampaignsv2TagResource {
			connectcampaignsv2_TagResource(cfg, client)
			return
		}
		if _connectcampaignsv2UntagResource {
			connectcampaignsv2_UntagResource(cfg, client)
			return
		}
		if _connectcampaignsv2UpdateCampaignChannelSubtypeConfig {
			connectcampaignsv2_UpdateCampaignChannelSubtypeConfig(cfg, client)
			return
		}
		if _connectcampaignsv2UpdateCampaignCommunicationLimits {
			connectcampaignsv2_UpdateCampaignCommunicationLimits(cfg, client)
			return
		}
		if _connectcampaignsv2UpdateCampaignCommunicationTime {
			connectcampaignsv2_UpdateCampaignCommunicationTime(cfg, client)
			return
		}
		if _connectcampaignsv2UpdateCampaignFlowAssociation {
			connectcampaignsv2_UpdateCampaignFlowAssociation(cfg, client)
			return
		}
		if _connectcampaignsv2UpdateCampaignName {
			connectcampaignsv2_UpdateCampaignName(cfg, client)
			return
		}
		if _connectcampaignsv2UpdateCampaignSchedule {
			connectcampaignsv2_UpdateCampaignSchedule(cfg, client)
			return
		}
		if _connectcampaignsv2UpdateCampaignSource {
			connectcampaignsv2_UpdateCampaignSource(cfg, client)
			return
		}

	},
}

var (
	_connectcampaignsv2CreateCampaign                     bool
	_connectcampaignsv2DeleteCampaign                     bool
	_connectcampaignsv2DeleteCampaignChannelSubtypeConfig bool
	_connectcampaignsv2DeleteCampaignCommunicationLimits  bool
	_connectcampaignsv2DeleteCampaignCommunicationTime    bool
	_connectcampaignsv2DeleteConnectInstanceConfig        bool
	_connectcampaignsv2DeleteConnectInstanceIntegration   bool
	_connectcampaignsv2DeleteInstanceOnboardingJob        bool
	_connectcampaignsv2DescribeCampaign                   bool
	_connectcampaignsv2GetCampaignState                   bool
	_connectcampaignsv2GetCampaignStateBatch              bool
	_connectcampaignsv2GetConnectInstanceConfig           bool
	_connectcampaignsv2GetInstanceCommunicationLimits     bool
	_connectcampaignsv2GetInstanceOnboardingJobStatus     bool
	_connectcampaignsv2ListCampaigns                      bool
	_connectcampaignsv2ListConnectInstanceIntegrations    bool
	_connectcampaignsv2ListTagsForResource                bool
	_connectcampaignsv2PauseCampaign                      bool
	_connectcampaignsv2PutConnectInstanceIntegration      bool
	_connectcampaignsv2PutInstanceCommunicationLimits     bool
	_connectcampaignsv2PutOutboundRequestBatch            bool
	_connectcampaignsv2PutProfileOutboundRequestBatch     bool
	_connectcampaignsv2ResumeCampaign                     bool
	_connectcampaignsv2StartCampaign                      bool
	_connectcampaignsv2StartInstanceOnboardingJob         bool
	_connectcampaignsv2StopCampaign                       bool
	_connectcampaignsv2TagResource                        bool
	_connectcampaignsv2UntagResource                      bool
	_connectcampaignsv2UpdateCampaignChannelSubtypeConfig bool
	_connectcampaignsv2UpdateCampaignCommunicationLimits  bool
	_connectcampaignsv2UpdateCampaignCommunicationTime    bool
	_connectcampaignsv2UpdateCampaignFlowAssociation      bool
	_connectcampaignsv2UpdateCampaignName                 bool
	_connectcampaignsv2UpdateCampaignSchedule             bool
	_connectcampaignsv2UpdateCampaignSource               bool

	_connectcampaignsv2Arn                         string
	_connectcampaignsv2CampaignDeletionPolicy      string
	_connectcampaignsv2CampaignIds                 []string
	_connectcampaignsv2ChannelSubtype              string
	_connectcampaignsv2ChannelSubtypeConfig        string
	_connectcampaignsv2CommunicationLimitsConfig   string
	_connectcampaignsv2CommunicationLimitsOverride string
	_connectcampaignsv2CommunicationTimeConfig     string
	_connectcampaignsv2Config                      string
	_connectcampaignsv2ConnectCampaignFlowArn      string
	_connectcampaignsv2ConnectInstanceId           string
	_connectcampaignsv2EncryptionConfig            string
	_connectcampaignsv2Filters                     string
	_connectcampaignsv2Id                          string
	_connectcampaignsv2IntegrationConfig           string
	_connectcampaignsv2IntegrationIdentifier       string
	_connectcampaignsv2MaxResults                  string
	_connectcampaignsv2Name                        string
	_connectcampaignsv2NextToken                   string
	_connectcampaignsv2OutboundRequests            string
	_connectcampaignsv2ProfileOutboundRequests     string
	_connectcampaignsv2Schedule                    string
	_connectcampaignsv2Source                      string
	_connectcampaignsv2TagKeys                     []string
	_connectcampaignsv2Tags                        string
	_connectcampaignsv2Type                        string
)

// Creates a campaign for the specified Amazon Connect account. This API is
// idempotent.
func connectcampaignsv2_CreateCampaign(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.CreateCampaignInput{
		// ConnectInstanceId: *string, // Required
		// Name: *string, // Required
	}

	if len(_connectcampaignsv2ConnectInstanceId) > 0 {
		input.ConnectInstanceId = aws.String(_connectcampaignsv2ConnectInstanceId)
	}
	if len(_connectcampaignsv2Name) > 0 {
		input.Name = aws.String(_connectcampaignsv2Name)
	}
	if len(_connectcampaignsv2ChannelSubtypeConfig) > 0 {
		if err := assignInputField(input, "ChannelSubtypeConfig", _connectcampaignsv2ChannelSubtypeConfig); err != nil {
			log.Errorf("invalid --channel-subtype-config: %s", err.Error())
			return
		}
	}
	if len(_connectcampaignsv2CommunicationLimitsOverride) > 0 {
		if err := assignInputField(input, "CommunicationLimitsOverride", _connectcampaignsv2CommunicationLimitsOverride); err != nil {
			log.Errorf("invalid --communication-limits-override: %s", err.Error())
			return
		}
	}
	if len(_connectcampaignsv2CommunicationTimeConfig) > 0 {
		if err := assignInputField(input, "CommunicationTimeConfig", _connectcampaignsv2CommunicationTimeConfig); err != nil {
			log.Errorf("invalid --communication-time-config: %s", err.Error())
			return
		}
	}
	if len(_connectcampaignsv2ConnectCampaignFlowArn) > 0 {
		input.ConnectCampaignFlowArn = aws.String(_connectcampaignsv2ConnectCampaignFlowArn)
	}
	if len(_connectcampaignsv2Schedule) > 0 {
		if err := assignInputField(input, "Schedule", _connectcampaignsv2Schedule); err != nil {
			log.Errorf("invalid --schedule: %s", err.Error())
			return
		}
	}
	if len(_connectcampaignsv2Source) > 0 {
		if err := assignInputField(input, "Source", _connectcampaignsv2Source); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}
	if len(_connectcampaignsv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _connectcampaignsv2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_connectcampaignsv2Type) > 0 {
		if err := assignInputField(input, "Type", _connectcampaignsv2Type); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a campaign from the specified Amazon Connect account.
func connectcampaignsv2_DeleteCampaign(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.DeleteCampaignInput{
		// Id: *string, // Required
	}

	if len(_connectcampaignsv2Id) > 0 {
		input.Id = aws.String(_connectcampaignsv2Id)
	}

	if resp, err := client.DeleteCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the channel subtype config of a campaign. This API is idempotent.
func connectcampaignsv2_DeleteCampaignChannelSubtypeConfig(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.DeleteCampaignChannelSubtypeConfigInput{
		// ChannelSubtype: types.ChannelSubtype, // Required
		// Id: *string, // Required
	}

	if len(_connectcampaignsv2ChannelSubtype) > 0 {
		if err := assignInputField(input, "ChannelSubtype", _connectcampaignsv2ChannelSubtype); err != nil {
			log.Errorf("invalid --channel-subtype: %s", err.Error())
			return
		}
	}
	if len(_connectcampaignsv2Id) > 0 {
		input.Id = aws.String(_connectcampaignsv2Id)
	}

	if resp, err := client.DeleteCampaignChannelSubtypeConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the communication limits config for a campaign. This API is idempotent.
func connectcampaignsv2_DeleteCampaignCommunicationLimits(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.DeleteCampaignCommunicationLimitsInput{
		// Config: types.CommunicationLimitsConfigType, // Required
		// Id: *string, // Required
	}

	if len(_connectcampaignsv2Config) > 0 {
		if err := assignInputField(input, "Config", _connectcampaignsv2Config); err != nil {
			log.Errorf("invalid --config: %s", err.Error())
			return
		}
	}
	if len(_connectcampaignsv2Id) > 0 {
		input.Id = aws.String(_connectcampaignsv2Id)
	}

	if resp, err := client.DeleteCampaignCommunicationLimits(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the communication time config for a campaign. This API is idempotent.
func connectcampaignsv2_DeleteCampaignCommunicationTime(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.DeleteCampaignCommunicationTimeInput{
		// Config: types.CommunicationTimeConfigType, // Required
		// Id: *string, // Required
	}

	if len(_connectcampaignsv2Config) > 0 {
		if err := assignInputField(input, "Config", _connectcampaignsv2Config); err != nil {
			log.Errorf("invalid --config: %s", err.Error())
			return
		}
	}
	if len(_connectcampaignsv2Id) > 0 {
		input.Id = aws.String(_connectcampaignsv2Id)
	}

	if resp, err := client.DeleteCampaignCommunicationTime(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a connect instance config from the specified AWS account.
func connectcampaignsv2_DeleteConnectInstanceConfig(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.DeleteConnectInstanceConfigInput{
		// ConnectInstanceId: *string, // Required
	}

	if len(_connectcampaignsv2ConnectInstanceId) > 0 {
		input.ConnectInstanceId = aws.String(_connectcampaignsv2ConnectInstanceId)
	}
	if len(_connectcampaignsv2CampaignDeletionPolicy) > 0 {
		if err := assignInputField(input, "CampaignDeletionPolicy", _connectcampaignsv2CampaignDeletionPolicy); err != nil {
			log.Errorf("invalid --campaign-deletion-policy: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteConnectInstanceConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete the integration for the specified Amazon Connect instance.
func connectcampaignsv2_DeleteConnectInstanceIntegration(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.DeleteConnectInstanceIntegrationInput{
		// ConnectInstanceId: *string, // Required
		// IntegrationIdentifier: types.IntegrationIdentifier, // Required
	}

	if len(_connectcampaignsv2ConnectInstanceId) > 0 {
		input.ConnectInstanceId = aws.String(_connectcampaignsv2ConnectInstanceId)
	}
	if len(_connectcampaignsv2IntegrationIdentifier) > 0 {
		if err := assignInputField(input, "IntegrationIdentifier", _connectcampaignsv2IntegrationIdentifier); err != nil {
			log.Errorf("invalid --integration-identifier: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteConnectInstanceIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete the Connect Campaigns onboarding job for the specified Amazon Connect
// instance.
func connectcampaignsv2_DeleteInstanceOnboardingJob(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.DeleteInstanceOnboardingJobInput{
		// ConnectInstanceId: *string, // Required
	}

	if len(_connectcampaignsv2ConnectInstanceId) > 0 {
		input.ConnectInstanceId = aws.String(_connectcampaignsv2ConnectInstanceId)
	}

	if resp, err := client.DeleteInstanceOnboardingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specific campaign.
func connectcampaignsv2_DescribeCampaign(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.DescribeCampaignInput{
		// Id: *string, // Required
	}

	if len(_connectcampaignsv2Id) > 0 {
		input.Id = aws.String(_connectcampaignsv2Id)
	}

	if resp, err := client.DescribeCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get state of a campaign for the specified Amazon Connect account.
func connectcampaignsv2_GetCampaignState(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.GetCampaignStateInput{
		// Id: *string, // Required
	}

	if len(_connectcampaignsv2Id) > 0 {
		input.Id = aws.String(_connectcampaignsv2Id)
	}

	if resp, err := client.GetCampaignState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get state of campaigns for the specified Amazon Connect account.
func connectcampaignsv2_GetCampaignStateBatch(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.GetCampaignStateBatchInput{
		// CampaignIds: []string, // Required
	}

	if len(_connectcampaignsv2CampaignIds) > 0 {
		input.CampaignIds = append([]string(nil), _connectcampaignsv2CampaignIds...)
	}

	if resp, err := client.GetCampaignStateBatch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the specific Connect instance config.
func connectcampaignsv2_GetConnectInstanceConfig(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.GetConnectInstanceConfigInput{
		// ConnectInstanceId: *string, // Required
	}

	if len(_connectcampaignsv2ConnectInstanceId) > 0 {
		input.ConnectInstanceId = aws.String(_connectcampaignsv2ConnectInstanceId)
	}

	if resp, err := client.GetConnectInstanceConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the instance communication limits.
func connectcampaignsv2_GetInstanceCommunicationLimits(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.GetInstanceCommunicationLimitsInput{
		// ConnectInstanceId: *string, // Required
	}

	if len(_connectcampaignsv2ConnectInstanceId) > 0 {
		input.ConnectInstanceId = aws.String(_connectcampaignsv2ConnectInstanceId)
	}

	if resp, err := client.GetInstanceCommunicationLimits(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the specific instance onboarding job status.
func connectcampaignsv2_GetInstanceOnboardingJobStatus(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.GetInstanceOnboardingJobStatusInput{
		// ConnectInstanceId: *string, // Required
	}

	if len(_connectcampaignsv2ConnectInstanceId) > 0 {
		input.ConnectInstanceId = aws.String(_connectcampaignsv2ConnectInstanceId)
	}

	if resp, err := client.GetInstanceOnboardingJobStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides summary information about the campaigns under the specified Amazon
// Connect account.
func connectcampaignsv2_ListCampaigns(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.ListCampaignsInput{}

	if len(_connectcampaignsv2Filters) > 0 {
		if err := assignInputField(input, "Filters", _connectcampaignsv2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_connectcampaignsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectcampaignsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectcampaignsv2NextToken) > 0 {
		input.NextToken = aws.String(_connectcampaignsv2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCampaigns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connectcampaignsv2.ListCampaignsOutput
	p := connectcampaignsv2.NewListCampaignsPaginator(client, input)
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

// Provides summary information about the integration under the specified Connect
// instance.
func connectcampaignsv2_ListConnectInstanceIntegrations(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.ListConnectInstanceIntegrationsInput{
		// ConnectInstanceId: *string, // Required
	}

	if len(_connectcampaignsv2ConnectInstanceId) > 0 {
		input.ConnectInstanceId = aws.String(_connectcampaignsv2ConnectInstanceId)
	}
	if len(_connectcampaignsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectcampaignsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectcampaignsv2NextToken) > 0 {
		input.NextToken = aws.String(_connectcampaignsv2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConnectInstanceIntegrations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connectcampaignsv2.ListConnectInstanceIntegrationsOutput
	p := connectcampaignsv2.NewListConnectInstanceIntegrationsPaginator(client, input)
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

// List tags for a resource.
func connectcampaignsv2_ListTagsForResource(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.ListTagsForResourceInput{
		// Arn: *string, // Required
	}

	if len(_connectcampaignsv2Arn) > 0 {
		input.Arn = aws.String(_connectcampaignsv2Arn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Pauses a campaign for the specified Amazon Connect account.
func connectcampaignsv2_PauseCampaign(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.PauseCampaignInput{
		// Id: *string, // Required
	}

	if len(_connectcampaignsv2Id) > 0 {
		input.Id = aws.String(_connectcampaignsv2Id)
	}

	if resp, err := client.PauseCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Put or update the integration for the specified Amazon Connect instance.
func connectcampaignsv2_PutConnectInstanceIntegration(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.PutConnectInstanceIntegrationInput{
		// ConnectInstanceId: *string, // Required
		// IntegrationConfig: types.IntegrationConfig, // Required
	}

	if len(_connectcampaignsv2ConnectInstanceId) > 0 {
		input.ConnectInstanceId = aws.String(_connectcampaignsv2ConnectInstanceId)
	}
	if len(_connectcampaignsv2IntegrationConfig) > 0 {
		if err := assignInputField(input, "IntegrationConfig", _connectcampaignsv2IntegrationConfig); err != nil {
			log.Errorf("invalid --integration-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutConnectInstanceIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Put the instance communication limits. This API is idempotent.
func connectcampaignsv2_PutInstanceCommunicationLimits(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.PutInstanceCommunicationLimitsInput{
		// CommunicationLimitsConfig: *types.InstanceCommunicationLimitsConfig, // Required
		// ConnectInstanceId: *string, // Required
	}

	if len(_connectcampaignsv2CommunicationLimitsConfig) > 0 {
		if err := assignInputField(input, "CommunicationLimitsConfig", _connectcampaignsv2CommunicationLimitsConfig); err != nil {
			log.Errorf("invalid --communication-limits-config: %s", err.Error())
			return
		}
	}
	if len(_connectcampaignsv2ConnectInstanceId) > 0 {
		input.ConnectInstanceId = aws.String(_connectcampaignsv2ConnectInstanceId)
	}

	if resp, err := client.PutInstanceCommunicationLimits(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates outbound requests for the specified campaign Amazon Connect account.
// This API is idempotent.
func connectcampaignsv2_PutOutboundRequestBatch(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.PutOutboundRequestBatchInput{
		// Id: *string, // Required
		// OutboundRequests: []types.OutboundRequest, // Required
	}

	if len(_connectcampaignsv2Id) > 0 {
		input.Id = aws.String(_connectcampaignsv2Id)
	}
	if len(_connectcampaignsv2OutboundRequests) > 0 {
		if err := assignInputField(input, "OutboundRequests", _connectcampaignsv2OutboundRequests); err != nil {
			log.Errorf("invalid --outbound-requests: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutOutboundRequestBatch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Takes in a list of profile outbound requests to be placed as part of an
// outbound campaign. This API is idempotent.
func connectcampaignsv2_PutProfileOutboundRequestBatch(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.PutProfileOutboundRequestBatchInput{
		// Id: *string, // Required
		// ProfileOutboundRequests: []types.ProfileOutboundRequest, // Required
	}

	if len(_connectcampaignsv2Id) > 0 {
		input.Id = aws.String(_connectcampaignsv2Id)
	}
	if len(_connectcampaignsv2ProfileOutboundRequests) > 0 {
		if err := assignInputField(input, "ProfileOutboundRequests", _connectcampaignsv2ProfileOutboundRequests); err != nil {
			log.Errorf("invalid --profile-outbound-requests: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutProfileOutboundRequestBatch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a campaign for the specified Amazon Connect account.
func connectcampaignsv2_ResumeCampaign(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.ResumeCampaignInput{
		// Id: *string, // Required
	}

	if len(_connectcampaignsv2Id) > 0 {
		input.Id = aws.String(_connectcampaignsv2Id)
	}

	if resp, err := client.ResumeCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a campaign for the specified Amazon Connect account.
func connectcampaignsv2_StartCampaign(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.StartCampaignInput{
		// Id: *string, // Required
	}

	if len(_connectcampaignsv2Id) > 0 {
		input.Id = aws.String(_connectcampaignsv2Id)
	}

	if resp, err := client.StartCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Onboard the specific Amazon Connect instance to Connect Campaigns.
func connectcampaignsv2_StartInstanceOnboardingJob(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.StartInstanceOnboardingJobInput{
		// ConnectInstanceId: *string, // Required
		// EncryptionConfig: *types.EncryptionConfig, // Required
	}

	if len(_connectcampaignsv2ConnectInstanceId) > 0 {
		input.ConnectInstanceId = aws.String(_connectcampaignsv2ConnectInstanceId)
	}
	if len(_connectcampaignsv2EncryptionConfig) > 0 {
		if err := assignInputField(input, "EncryptionConfig", _connectcampaignsv2EncryptionConfig); err != nil {
			log.Errorf("invalid --encryption-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartInstanceOnboardingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a campaign for the specified Amazon Connect account.
func connectcampaignsv2_StopCampaign(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.StopCampaignInput{
		// Id: *string, // Required
	}

	if len(_connectcampaignsv2Id) > 0 {
		input.Id = aws.String(_connectcampaignsv2Id)
	}

	if resp, err := client.StopCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tag a resource.
func connectcampaignsv2_TagResource(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.TagResourceInput{
		// Arn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_connectcampaignsv2Arn) > 0 {
		input.Arn = aws.String(_connectcampaignsv2Arn)
	}
	if len(_connectcampaignsv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _connectcampaignsv2Tags); err != nil {
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

// Untag a resource.
func connectcampaignsv2_UntagResource(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.UntagResourceInput{
		// Arn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_connectcampaignsv2Arn) > 0 {
		input.Arn = aws.String(_connectcampaignsv2Arn)
	}
	if len(_connectcampaignsv2TagKeys) > 0 {
		input.TagKeys = append([]string(nil), _connectcampaignsv2TagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the channel subtype config of a campaign. This API is idempotent.
func connectcampaignsv2_UpdateCampaignChannelSubtypeConfig(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.UpdateCampaignChannelSubtypeConfigInput{
		// ChannelSubtypeConfig: *types.ChannelSubtypeConfig, // Required
		// Id: *string, // Required
	}

	if len(_connectcampaignsv2ChannelSubtypeConfig) > 0 {
		if err := assignInputField(input, "ChannelSubtypeConfig", _connectcampaignsv2ChannelSubtypeConfig); err != nil {
			log.Errorf("invalid --channel-subtype-config: %s", err.Error())
			return
		}
	}
	if len(_connectcampaignsv2Id) > 0 {
		input.Id = aws.String(_connectcampaignsv2Id)
	}

	if resp, err := client.UpdateCampaignChannelSubtypeConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the communication limits config for a campaign. This API is idempotent.
func connectcampaignsv2_UpdateCampaignCommunicationLimits(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.UpdateCampaignCommunicationLimitsInput{
		// CommunicationLimitsOverride: *types.CommunicationLimitsConfig, // Required
		// Id: *string, // Required
	}

	if len(_connectcampaignsv2CommunicationLimitsOverride) > 0 {
		if err := assignInputField(input, "CommunicationLimitsOverride", _connectcampaignsv2CommunicationLimitsOverride); err != nil {
			log.Errorf("invalid --communication-limits-override: %s", err.Error())
			return
		}
	}
	if len(_connectcampaignsv2Id) > 0 {
		input.Id = aws.String(_connectcampaignsv2Id)
	}

	if resp, err := client.UpdateCampaignCommunicationLimits(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the communication time config for a campaign. This API is idempotent.
func connectcampaignsv2_UpdateCampaignCommunicationTime(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.UpdateCampaignCommunicationTimeInput{
		// CommunicationTimeConfig: *types.CommunicationTimeConfig, // Required
		// Id: *string, // Required
	}

	if len(_connectcampaignsv2CommunicationTimeConfig) > 0 {
		if err := assignInputField(input, "CommunicationTimeConfig", _connectcampaignsv2CommunicationTimeConfig); err != nil {
			log.Errorf("invalid --communication-time-config: %s", err.Error())
			return
		}
	}
	if len(_connectcampaignsv2Id) > 0 {
		input.Id = aws.String(_connectcampaignsv2Id)
	}

	if resp, err := client.UpdateCampaignCommunicationTime(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the campaign flow associated with a campaign. This API is idempotent.
func connectcampaignsv2_UpdateCampaignFlowAssociation(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.UpdateCampaignFlowAssociationInput{
		// ConnectCampaignFlowArn: *string, // Required
		// Id: *string, // Required
	}

	if len(_connectcampaignsv2ConnectCampaignFlowArn) > 0 {
		input.ConnectCampaignFlowArn = aws.String(_connectcampaignsv2ConnectCampaignFlowArn)
	}
	if len(_connectcampaignsv2Id) > 0 {
		input.Id = aws.String(_connectcampaignsv2Id)
	}

	if resp, err := client.UpdateCampaignFlowAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the name of a campaign. This API is idempotent.
func connectcampaignsv2_UpdateCampaignName(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.UpdateCampaignNameInput{
		// Id: *string, // Required
		// Name: *string, // Required
	}

	if len(_connectcampaignsv2Id) > 0 {
		input.Id = aws.String(_connectcampaignsv2Id)
	}
	if len(_connectcampaignsv2Name) > 0 {
		input.Name = aws.String(_connectcampaignsv2Name)
	}

	if resp, err := client.UpdateCampaignName(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the schedule for a campaign. This API is idempotent.
func connectcampaignsv2_UpdateCampaignSchedule(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.UpdateCampaignScheduleInput{
		// Id: *string, // Required
		// Schedule: *types.Schedule, // Required
	}

	if len(_connectcampaignsv2Id) > 0 {
		input.Id = aws.String(_connectcampaignsv2Id)
	}
	if len(_connectcampaignsv2Schedule) > 0 {
		if err := assignInputField(input, "Schedule", _connectcampaignsv2Schedule); err != nil {
			log.Errorf("invalid --schedule: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCampaignSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the campaign source with a campaign. This API is idempotent.
func connectcampaignsv2_UpdateCampaignSource(cfg aws.Config, client *connectcampaignsv2.Client) {
	input := &connectcampaignsv2.UpdateCampaignSourceInput{
		// Id: *string, // Required
		// Source: types.Source, // Required
	}

	if len(_connectcampaignsv2Id) > 0 {
		input.Id = aws.String(_connectcampaignsv2Id)
	}
	if len(_connectcampaignsv2Source) > 0 {
		if err := assignInputField(input, "Source", _connectcampaignsv2Source); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCampaignSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_connectcampaignsv2Cmd)
	_connectcampaignsv2Cmd.Flags().SortFlags = false

	_connectcampaignsv2Cmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_connectcampaignsv2Cmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_connectcampaignsv2Cmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_connectcampaignsv2Cmd.Flags().StringVarP(&_connectcampaignsv2Arn, "arn", "", "", "ARN")
	_connectcampaignsv2Cmd.Flags().StringVarP(&_connectcampaignsv2CampaignDeletionPolicy, "campaign-deletion-policy", "", "", "Campaign Deletion Policy")
	_connectcampaignsv2Cmd.Flags().StringSliceVarP(&_connectcampaignsv2CampaignIds, "campaign-ids", "", nil, "Campaign Ids")
	_connectcampaignsv2Cmd.Flags().StringVarP(&_connectcampaignsv2ChannelSubtype, "channel-subtype", "", "", "Channel Subtype")
	_connectcampaignsv2Cmd.Flags().StringVarP(&_connectcampaignsv2ChannelSubtypeConfig, "channel-subtype-config", "", "", "Channel Subtype Config")
	_connectcampaignsv2Cmd.Flags().StringVarP(&_connectcampaignsv2CommunicationLimitsConfig, "communication-limits-config", "", "", "Communication Limits Config")
	_connectcampaignsv2Cmd.Flags().StringVarP(&_connectcampaignsv2CommunicationLimitsOverride, "communication-limits-override", "", "", "Communication Limits Override")
	_connectcampaignsv2Cmd.Flags().StringVarP(&_connectcampaignsv2CommunicationTimeConfig, "communication-time-config", "", "", "Communication Time Config")
	_connectcampaignsv2Cmd.Flags().StringVarP(&_connectcampaignsv2Config, "config", "", "", "Config")
	_connectcampaignsv2Cmd.Flags().StringVarP(&_connectcampaignsv2ConnectCampaignFlowArn, "connect-campaign-flow-arn", "", "", "Connect Campaign Flow ARN")
	_connectcampaignsv2Cmd.Flags().StringVarP(&_connectcampaignsv2ConnectInstanceId, "connect-instance-id", "", "", "Connect Instance ID")
	_connectcampaignsv2Cmd.Flags().StringVarP(&_connectcampaignsv2EncryptionConfig, "encryption-config", "", "", "Encryption Config")
	_connectcampaignsv2Cmd.Flags().StringVarP(&_connectcampaignsv2Filters, "filters", "", "", "Filters")
	_connectcampaignsv2Cmd.Flags().StringVarP(&_connectcampaignsv2Id, "id", "", "", "ID")
	_connectcampaignsv2Cmd.Flags().StringVarP(&_connectcampaignsv2IntegrationConfig, "integration-config", "", "", "Integration Config")
	_connectcampaignsv2Cmd.Flags().StringVarP(&_connectcampaignsv2IntegrationIdentifier, "integration-identifier", "", "", "Integration Identifier")
	_connectcampaignsv2Cmd.Flags().StringVarP(&_connectcampaignsv2MaxResults, "max-results", "", "", "Max Results")
	_connectcampaignsv2Cmd.Flags().StringVarP(&_connectcampaignsv2Name, "name", "", "", "Name")
	_connectcampaignsv2Cmd.Flags().StringVarP(&_connectcampaignsv2NextToken, "next-token", "", "", "Next Token")
	_connectcampaignsv2Cmd.Flags().StringVarP(&_connectcampaignsv2OutboundRequests, "outbound-requests", "", "", "Outbound Requests")
	_connectcampaignsv2Cmd.Flags().StringVarP(&_connectcampaignsv2ProfileOutboundRequests, "profile-outbound-requests", "", "", "Profile Outbound Requests")
	_connectcampaignsv2Cmd.Flags().StringVarP(&_connectcampaignsv2Schedule, "schedule", "", "", "Schedule")
	_connectcampaignsv2Cmd.Flags().StringVarP(&_connectcampaignsv2Source, "source", "", "", "Source")
	_connectcampaignsv2Cmd.Flags().StringSliceVarP(&_connectcampaignsv2TagKeys, "tag-keys", "", nil, "Tag Keys")
	_connectcampaignsv2Cmd.Flags().StringVarP(&_connectcampaignsv2Tags, "tags", "", "", "Tags")
	_connectcampaignsv2Cmd.Flags().StringVarP(&_connectcampaignsv2Type, "type", "", "", "Type")

	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2CreateCampaign, "create-campaign", "", false, "Create Campaign")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2DeleteCampaign, "delete-campaign", "", false, "Delete Campaign")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2DeleteCampaignChannelSubtypeConfig, "delete-campaign-channel-subtype-config", "", false, "Delete Campaign Channel Subtype Config")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2DeleteCampaignCommunicationLimits, "delete-campaign-communication-limits", "", false, "Delete Campaign Communication Limits")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2DeleteCampaignCommunicationTime, "delete-campaign-communication-time", "", false, "Delete Campaign Communication Time")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2DeleteConnectInstanceConfig, "delete-connect-instance-config", "", false, "Delete Connect Instance Config")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2DeleteConnectInstanceIntegration, "delete-connect-instance-integration", "", false, "Delete Connect Instance Integration")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2DeleteInstanceOnboardingJob, "delete-instance-onboarding-job", "", false, "Delete Instance Onboarding Job")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2DescribeCampaign, "describe-campaign", "", false, "Describe Campaign")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2GetCampaignState, "get-campaign-state", "", false, "Get Campaign State")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2GetCampaignStateBatch, "get-campaign-state-batch", "", false, "Get Campaign State Batch")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2GetConnectInstanceConfig, "get-connect-instance-config", "", false, "Get Connect Instance Config")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2GetInstanceCommunicationLimits, "get-instance-communication-limits", "", false, "Get Instance Communication Limits")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2GetInstanceOnboardingJobStatus, "get-instance-onboarding-job-status", "", false, "Get Instance Onboarding Job Status")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2ListCampaigns, "list-campaigns", "", false, "List Campaigns")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2ListConnectInstanceIntegrations, "list-connect-instance-integrations", "", false, "List Connect Instance Integrations")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2ListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2PauseCampaign, "pause-campaign", "", false, "Pause Campaign")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2PutConnectInstanceIntegration, "put-connect-instance-integration", "", false, "Put Connect Instance Integration")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2PutInstanceCommunicationLimits, "put-instance-communication-limits", "", false, "Put Instance Communication Limits")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2PutOutboundRequestBatch, "put-outbound-request-batch", "", false, "Put Outbound Request Batch")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2PutProfileOutboundRequestBatch, "put-profile-outbound-request-batch", "", false, "Put Profile Outbound Request Batch")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2ResumeCampaign, "resume-campaign", "", false, "Resume Campaign")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2StartCampaign, "start-campaign", "", false, "Start Campaign")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2StartInstanceOnboardingJob, "start-instance-onboarding-job", "", false, "Start Instance Onboarding Job")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2StopCampaign, "stop-campaign", "", false, "Stop Campaign")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2TagResource, "tag-resource", "", false, "Tag Resource")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2UntagResource, "untag-resource", "", false, "Untag Resource")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2UpdateCampaignChannelSubtypeConfig, "update-campaign-channel-subtype-config", "", false, "Update Campaign Channel Subtype Config")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2UpdateCampaignCommunicationLimits, "update-campaign-communication-limits", "", false, "Update Campaign Communication Limits")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2UpdateCampaignCommunicationTime, "update-campaign-communication-time", "", false, "Update Campaign Communication Time")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2UpdateCampaignFlowAssociation, "update-campaign-flow-association", "", false, "Update Campaign Flow Association")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2UpdateCampaignName, "update-campaign-name", "", false, "Update Campaign Name")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2UpdateCampaignSchedule, "update-campaign-schedule", "", false, "Update Campaign Schedule")
	_connectcampaignsv2Cmd.Flags().BoolVarP(&_connectcampaignsv2UpdateCampaignSource, "update-campaign-source", "", false, "Update Campaign Source")

}
