package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/connectcampaigns"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// connectcampaignsCmd represents the connectcampaigns command
var _connectcampaignsCmd = &cobra.Command{
	Use:   "connectcampaigns",
	Short: "AWS connectcampaigns CLI",
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
		client := connectcampaigns.NewFromConfig(cfg)
		if _connectcampaignsCreateCampaign {
			connectcampaigns_CreateCampaign(cfg, client)
			return
		}
		if _connectcampaignsDeleteCampaign {
			connectcampaigns_DeleteCampaign(cfg, client)
			return
		}
		if _connectcampaignsDeleteConnectInstanceConfig {
			connectcampaigns_DeleteConnectInstanceConfig(cfg, client)
			return
		}
		if _connectcampaignsDeleteInstanceOnboardingJob {
			connectcampaigns_DeleteInstanceOnboardingJob(cfg, client)
			return
		}
		if _connectcampaignsDescribeCampaign {
			connectcampaigns_DescribeCampaign(cfg, client)
			return
		}
		if _connectcampaignsGetCampaignState {
			connectcampaigns_GetCampaignState(cfg, client)
			return
		}
		if _connectcampaignsGetCampaignStateBatch {
			connectcampaigns_GetCampaignStateBatch(cfg, client)
			return
		}
		if _connectcampaignsGetConnectInstanceConfig {
			connectcampaigns_GetConnectInstanceConfig(cfg, client)
			return
		}
		if _connectcampaignsGetInstanceOnboardingJobStatus {
			connectcampaigns_GetInstanceOnboardingJobStatus(cfg, client)
			return
		}
		if _connectcampaignsListCampaigns {
			connectcampaigns_ListCampaigns(cfg, client)
			return
		}
		if _connectcampaignsListTagsForResource {
			connectcampaigns_ListTagsForResource(cfg, client)
			return
		}
		if _connectcampaignsPauseCampaign {
			connectcampaigns_PauseCampaign(cfg, client)
			return
		}
		if _connectcampaignsPutDialRequestBatch {
			connectcampaigns_PutDialRequestBatch(cfg, client)
			return
		}
		if _connectcampaignsResumeCampaign {
			connectcampaigns_ResumeCampaign(cfg, client)
			return
		}
		if _connectcampaignsStartCampaign {
			connectcampaigns_StartCampaign(cfg, client)
			return
		}
		if _connectcampaignsStartInstanceOnboardingJob {
			connectcampaigns_StartInstanceOnboardingJob(cfg, client)
			return
		}
		if _connectcampaignsStopCampaign {
			connectcampaigns_StopCampaign(cfg, client)
			return
		}
		if _connectcampaignsTagResource {
			connectcampaigns_TagResource(cfg, client)
			return
		}
		if _connectcampaignsUntagResource {
			connectcampaigns_UntagResource(cfg, client)
			return
		}
		if _connectcampaignsUpdateCampaignDialerConfig {
			connectcampaigns_UpdateCampaignDialerConfig(cfg, client)
			return
		}
		if _connectcampaignsUpdateCampaignName {
			connectcampaigns_UpdateCampaignName(cfg, client)
			return
		}
		if _connectcampaignsUpdateCampaignOutboundCallConfig {
			connectcampaigns_UpdateCampaignOutboundCallConfig(cfg, client)
			return
		}

	},
}

var (
	_connectcampaignsCreateCampaign                   bool
	_connectcampaignsDeleteCampaign                   bool
	_connectcampaignsDeleteConnectInstanceConfig      bool
	_connectcampaignsDeleteInstanceOnboardingJob      bool
	_connectcampaignsDescribeCampaign                 bool
	_connectcampaignsGetCampaignState                 bool
	_connectcampaignsGetCampaignStateBatch            bool
	_connectcampaignsGetConnectInstanceConfig         bool
	_connectcampaignsGetInstanceOnboardingJobStatus   bool
	_connectcampaignsListCampaigns                    bool
	_connectcampaignsListTagsForResource              bool
	_connectcampaignsPauseCampaign                    bool
	_connectcampaignsPutDialRequestBatch              bool
	_connectcampaignsResumeCampaign                   bool
	_connectcampaignsStartCampaign                    bool
	_connectcampaignsStartInstanceOnboardingJob       bool
	_connectcampaignsStopCampaign                     bool
	_connectcampaignsTagResource                      bool
	_connectcampaignsUntagResource                    bool
	_connectcampaignsUpdateCampaignDialerConfig       bool
	_connectcampaignsUpdateCampaignName               bool
	_connectcampaignsUpdateCampaignOutboundCallConfig bool

	_connectcampaignsAnswerMachineDetectionConfig string
	_connectcampaignsArn                          string
	_connectcampaignsCampaignIds                  []string
	_connectcampaignsConnectContactFlowId         string
	_connectcampaignsConnectInstanceId            string
	_connectcampaignsConnectSourcePhoneNumber     string
	_connectcampaignsDialRequests                 string
	_connectcampaignsDialerConfig                 string
	_connectcampaignsEncryptionConfig             string
	_connectcampaignsFilters                      string
	_connectcampaignsId                           string
	_connectcampaignsMaxResults                   string
	_connectcampaignsName                         string
	_connectcampaignsNextToken                    string
	_connectcampaignsOutboundCallConfig           string
	_connectcampaignsTagKeys                      []string
	_connectcampaignsTags                         string
)

// Creates a campaign for the specified Amazon Connect account. This API is
// idempotent.
func connectcampaigns_CreateCampaign(cfg aws.Config, client *connectcampaigns.Client) {
	input := &connectcampaigns.CreateCampaignInput{
		// ConnectInstanceId: *string, // Required
		// DialerConfig: types.DialerConfig, // Required
		// Name: *string, // Required
		// OutboundCallConfig: *types.OutboundCallConfig, // Required
	}

	if len(_connectcampaignsConnectInstanceId) > 0 {
		input.ConnectInstanceId = aws.String(_connectcampaignsConnectInstanceId)
	}
	if len(_connectcampaignsDialerConfig) > 0 {
		if err := assignInputField(input, "DialerConfig", _connectcampaignsDialerConfig); err != nil {
			log.Errorf("invalid --dialer-config: %s", err.Error())
			return
		}
	}
	if len(_connectcampaignsName) > 0 {
		input.Name = aws.String(_connectcampaignsName)
	}
	if len(_connectcampaignsOutboundCallConfig) > 0 {
		if err := assignInputField(input, "OutboundCallConfig", _connectcampaignsOutboundCallConfig); err != nil {
			log.Errorf("invalid --outbound-call-config: %s", err.Error())
			return
		}
	}
	if len(_connectcampaignsTags) > 0 {
		if err := assignInputField(input, "Tags", _connectcampaignsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
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
func connectcampaigns_DeleteCampaign(cfg aws.Config, client *connectcampaigns.Client) {
	input := &connectcampaigns.DeleteCampaignInput{
		// Id: *string, // Required
	}

	if len(_connectcampaignsId) > 0 {
		input.Id = aws.String(_connectcampaignsId)
	}

	if resp, err := client.DeleteCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a connect instance config from the specified AWS account.
func connectcampaigns_DeleteConnectInstanceConfig(cfg aws.Config, client *connectcampaigns.Client) {
	input := &connectcampaigns.DeleteConnectInstanceConfigInput{
		// ConnectInstanceId: *string, // Required
	}

	if len(_connectcampaignsConnectInstanceId) > 0 {
		input.ConnectInstanceId = aws.String(_connectcampaignsConnectInstanceId)
	}

	if resp, err := client.DeleteConnectInstanceConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete the Connect Campaigns onboarding job for the specified Amazon Connect
// instance.
func connectcampaigns_DeleteInstanceOnboardingJob(cfg aws.Config, client *connectcampaigns.Client) {
	input := &connectcampaigns.DeleteInstanceOnboardingJobInput{
		// ConnectInstanceId: *string, // Required
	}

	if len(_connectcampaignsConnectInstanceId) > 0 {
		input.ConnectInstanceId = aws.String(_connectcampaignsConnectInstanceId)
	}

	if resp, err := client.DeleteInstanceOnboardingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specific campaign.
func connectcampaigns_DescribeCampaign(cfg aws.Config, client *connectcampaigns.Client) {
	input := &connectcampaigns.DescribeCampaignInput{
		// Id: *string, // Required
	}

	if len(_connectcampaignsId) > 0 {
		input.Id = aws.String(_connectcampaignsId)
	}

	if resp, err := client.DescribeCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get state of a campaign for the specified Amazon Connect account.
func connectcampaigns_GetCampaignState(cfg aws.Config, client *connectcampaigns.Client) {
	input := &connectcampaigns.GetCampaignStateInput{
		// Id: *string, // Required
	}

	if len(_connectcampaignsId) > 0 {
		input.Id = aws.String(_connectcampaignsId)
	}

	if resp, err := client.GetCampaignState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get state of campaigns for the specified Amazon Connect account.
func connectcampaigns_GetCampaignStateBatch(cfg aws.Config, client *connectcampaigns.Client) {
	input := &connectcampaigns.GetCampaignStateBatchInput{
		// CampaignIds: []string, // Required
	}

	if len(_connectcampaignsCampaignIds) > 0 {
		input.CampaignIds = append([]string(nil), _connectcampaignsCampaignIds...)
	}

	if resp, err := client.GetCampaignStateBatch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the specific Connect instance config.
func connectcampaigns_GetConnectInstanceConfig(cfg aws.Config, client *connectcampaigns.Client) {
	input := &connectcampaigns.GetConnectInstanceConfigInput{
		// ConnectInstanceId: *string, // Required
	}

	if len(_connectcampaignsConnectInstanceId) > 0 {
		input.ConnectInstanceId = aws.String(_connectcampaignsConnectInstanceId)
	}

	if resp, err := client.GetConnectInstanceConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the specific instance onboarding job status.
func connectcampaigns_GetInstanceOnboardingJobStatus(cfg aws.Config, client *connectcampaigns.Client) {
	input := &connectcampaigns.GetInstanceOnboardingJobStatusInput{
		// ConnectInstanceId: *string, // Required
	}

	if len(_connectcampaignsConnectInstanceId) > 0 {
		input.ConnectInstanceId = aws.String(_connectcampaignsConnectInstanceId)
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
func connectcampaigns_ListCampaigns(cfg aws.Config, client *connectcampaigns.Client) {
	input := &connectcampaigns.ListCampaignsInput{}

	if len(_connectcampaignsFilters) > 0 {
		if err := assignInputField(input, "Filters", _connectcampaignsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_connectcampaignsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectcampaignsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectcampaignsNextToken) > 0 {
		input.NextToken = aws.String(_connectcampaignsNextToken)
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

	var results []*connectcampaigns.ListCampaignsOutput
	p := connectcampaigns.NewListCampaignsPaginator(client, input)
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
func connectcampaigns_ListTagsForResource(cfg aws.Config, client *connectcampaigns.Client) {
	input := &connectcampaigns.ListTagsForResourceInput{
		// Arn: *string, // Required
	}

	if len(_connectcampaignsArn) > 0 {
		input.Arn = aws.String(_connectcampaignsArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Pauses a campaign for the specified Amazon Connect account.
func connectcampaigns_PauseCampaign(cfg aws.Config, client *connectcampaigns.Client) {
	input := &connectcampaigns.PauseCampaignInput{
		// Id: *string, // Required
	}

	if len(_connectcampaignsId) > 0 {
		input.Id = aws.String(_connectcampaignsId)
	}

	if resp, err := client.PauseCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates dials requests for the specified campaign Amazon Connect account. This
// API is idempotent.
func connectcampaigns_PutDialRequestBatch(cfg aws.Config, client *connectcampaigns.Client) {
	input := &connectcampaigns.PutDialRequestBatchInput{
		// DialRequests: []types.DialRequest, // Required
		// Id: *string, // Required
	}

	if len(_connectcampaignsDialRequests) > 0 {
		if err := assignInputField(input, "DialRequests", _connectcampaignsDialRequests); err != nil {
			log.Errorf("invalid --dial-requests: %s", err.Error())
			return
		}
	}
	if len(_connectcampaignsId) > 0 {
		input.Id = aws.String(_connectcampaignsId)
	}

	if resp, err := client.PutDialRequestBatch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a campaign for the specified Amazon Connect account.
func connectcampaigns_ResumeCampaign(cfg aws.Config, client *connectcampaigns.Client) {
	input := &connectcampaigns.ResumeCampaignInput{
		// Id: *string, // Required
	}

	if len(_connectcampaignsId) > 0 {
		input.Id = aws.String(_connectcampaignsId)
	}

	if resp, err := client.ResumeCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a campaign for the specified Amazon Connect account.
func connectcampaigns_StartCampaign(cfg aws.Config, client *connectcampaigns.Client) {
	input := &connectcampaigns.StartCampaignInput{
		// Id: *string, // Required
	}

	if len(_connectcampaignsId) > 0 {
		input.Id = aws.String(_connectcampaignsId)
	}

	if resp, err := client.StartCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Onboard the specific Amazon Connect instance to Connect Campaigns.
func connectcampaigns_StartInstanceOnboardingJob(cfg aws.Config, client *connectcampaigns.Client) {
	input := &connectcampaigns.StartInstanceOnboardingJobInput{
		// ConnectInstanceId: *string, // Required
		// EncryptionConfig: *types.EncryptionConfig, // Required
	}

	if len(_connectcampaignsConnectInstanceId) > 0 {
		input.ConnectInstanceId = aws.String(_connectcampaignsConnectInstanceId)
	}
	if len(_connectcampaignsEncryptionConfig) > 0 {
		if err := assignInputField(input, "EncryptionConfig", _connectcampaignsEncryptionConfig); err != nil {
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
func connectcampaigns_StopCampaign(cfg aws.Config, client *connectcampaigns.Client) {
	input := &connectcampaigns.StopCampaignInput{
		// Id: *string, // Required
	}

	if len(_connectcampaignsId) > 0 {
		input.Id = aws.String(_connectcampaignsId)
	}

	if resp, err := client.StopCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tag a resource.
func connectcampaigns_TagResource(cfg aws.Config, client *connectcampaigns.Client) {
	input := &connectcampaigns.TagResourceInput{
		// Arn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_connectcampaignsArn) > 0 {
		input.Arn = aws.String(_connectcampaignsArn)
	}
	if len(_connectcampaignsTags) > 0 {
		if err := assignInputField(input, "Tags", _connectcampaignsTags); err != nil {
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
func connectcampaigns_UntagResource(cfg aws.Config, client *connectcampaigns.Client) {
	input := &connectcampaigns.UntagResourceInput{
		// Arn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_connectcampaignsArn) > 0 {
		input.Arn = aws.String(_connectcampaignsArn)
	}
	if len(_connectcampaignsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _connectcampaignsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the dialer config of a campaign. This API is idempotent.
func connectcampaigns_UpdateCampaignDialerConfig(cfg aws.Config, client *connectcampaigns.Client) {
	input := &connectcampaigns.UpdateCampaignDialerConfigInput{
		// DialerConfig: types.DialerConfig, // Required
		// Id: *string, // Required
	}

	if len(_connectcampaignsDialerConfig) > 0 {
		if err := assignInputField(input, "DialerConfig", _connectcampaignsDialerConfig); err != nil {
			log.Errorf("invalid --dialer-config: %s", err.Error())
			return
		}
	}
	if len(_connectcampaignsId) > 0 {
		input.Id = aws.String(_connectcampaignsId)
	}

	if resp, err := client.UpdateCampaignDialerConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the name of a campaign. This API is idempotent.
func connectcampaigns_UpdateCampaignName(cfg aws.Config, client *connectcampaigns.Client) {
	input := &connectcampaigns.UpdateCampaignNameInput{
		// Id: *string, // Required
		// Name: *string, // Required
	}

	if len(_connectcampaignsId) > 0 {
		input.Id = aws.String(_connectcampaignsId)
	}
	if len(_connectcampaignsName) > 0 {
		input.Name = aws.String(_connectcampaignsName)
	}

	if resp, err := client.UpdateCampaignName(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the outbound call config of a campaign. This API is idempotent.
func connectcampaigns_UpdateCampaignOutboundCallConfig(cfg aws.Config, client *connectcampaigns.Client) {
	input := &connectcampaigns.UpdateCampaignOutboundCallConfigInput{
		// Id: *string, // Required
	}

	if len(_connectcampaignsId) > 0 {
		input.Id = aws.String(_connectcampaignsId)
	}
	if len(_connectcampaignsAnswerMachineDetectionConfig) > 0 {
		if err := assignInputField(input, "AnswerMachineDetectionConfig", _connectcampaignsAnswerMachineDetectionConfig); err != nil {
			log.Errorf("invalid --answer-machine-detection-config: %s", err.Error())
			return
		}
	}
	if len(_connectcampaignsConnectContactFlowId) > 0 {
		input.ConnectContactFlowId = aws.String(_connectcampaignsConnectContactFlowId)
	}
	if len(_connectcampaignsConnectSourcePhoneNumber) > 0 {
		input.ConnectSourcePhoneNumber = aws.String(_connectcampaignsConnectSourcePhoneNumber)
	}

	if resp, err := client.UpdateCampaignOutboundCallConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_connectcampaignsCmd)
	_connectcampaignsCmd.Flags().SortFlags = false

	_connectcampaignsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_connectcampaignsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_connectcampaignsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_connectcampaignsCmd.Flags().StringVarP(&_connectcampaignsAnswerMachineDetectionConfig, "answer-machine-detection-config", "", "", "Answer Machine Detection Config")
	_connectcampaignsCmd.Flags().StringVarP(&_connectcampaignsArn, "arn", "", "", "ARN")
	_connectcampaignsCmd.Flags().StringSliceVarP(&_connectcampaignsCampaignIds, "campaign-ids", "", nil, "Campaign Ids")
	_connectcampaignsCmd.Flags().StringVarP(&_connectcampaignsConnectContactFlowId, "connect-contact-flow-id", "", "", "Connect Contact Flow ID")
	_connectcampaignsCmd.Flags().StringVarP(&_connectcampaignsConnectInstanceId, "connect-instance-id", "", "", "Connect Instance ID")
	_connectcampaignsCmd.Flags().StringVarP(&_connectcampaignsConnectSourcePhoneNumber, "connect-source-phone-number", "", "", "Connect Source Phone Number")
	_connectcampaignsCmd.Flags().StringVarP(&_connectcampaignsDialRequests, "dial-requests", "", "", "Dial Requests")
	_connectcampaignsCmd.Flags().StringVarP(&_connectcampaignsDialerConfig, "dialer-config", "", "", "Dialer Config")
	_connectcampaignsCmd.Flags().StringVarP(&_connectcampaignsEncryptionConfig, "encryption-config", "", "", "Encryption Config")
	_connectcampaignsCmd.Flags().StringVarP(&_connectcampaignsFilters, "filters", "", "", "Filters")
	_connectcampaignsCmd.Flags().StringVarP(&_connectcampaignsId, "id", "", "", "ID")
	_connectcampaignsCmd.Flags().StringVarP(&_connectcampaignsMaxResults, "max-results", "", "", "Max Results")
	_connectcampaignsCmd.Flags().StringVarP(&_connectcampaignsName, "name", "", "", "Name")
	_connectcampaignsCmd.Flags().StringVarP(&_connectcampaignsNextToken, "next-token", "", "", "Next Token")
	_connectcampaignsCmd.Flags().StringVarP(&_connectcampaignsOutboundCallConfig, "outbound-call-config", "", "", "Outbound Call Config")
	_connectcampaignsCmd.Flags().StringSliceVarP(&_connectcampaignsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_connectcampaignsCmd.Flags().StringVarP(&_connectcampaignsTags, "tags", "", "", "Tags")

	_connectcampaignsCmd.Flags().BoolVarP(&_connectcampaignsCreateCampaign, "create-campaign", "", false, "Create Campaign")
	_connectcampaignsCmd.Flags().BoolVarP(&_connectcampaignsDeleteCampaign, "delete-campaign", "", false, "Delete Campaign")
	_connectcampaignsCmd.Flags().BoolVarP(&_connectcampaignsDeleteConnectInstanceConfig, "delete-connect-instance-config", "", false, "Delete Connect Instance Config")
	_connectcampaignsCmd.Flags().BoolVarP(&_connectcampaignsDeleteInstanceOnboardingJob, "delete-instance-onboarding-job", "", false, "Delete Instance Onboarding Job")
	_connectcampaignsCmd.Flags().BoolVarP(&_connectcampaignsDescribeCampaign, "describe-campaign", "", false, "Describe Campaign")
	_connectcampaignsCmd.Flags().BoolVarP(&_connectcampaignsGetCampaignState, "get-campaign-state", "", false, "Get Campaign State")
	_connectcampaignsCmd.Flags().BoolVarP(&_connectcampaignsGetCampaignStateBatch, "get-campaign-state-batch", "", false, "Get Campaign State Batch")
	_connectcampaignsCmd.Flags().BoolVarP(&_connectcampaignsGetConnectInstanceConfig, "get-connect-instance-config", "", false, "Get Connect Instance Config")
	_connectcampaignsCmd.Flags().BoolVarP(&_connectcampaignsGetInstanceOnboardingJobStatus, "get-instance-onboarding-job-status", "", false, "Get Instance Onboarding Job Status")
	_connectcampaignsCmd.Flags().BoolVarP(&_connectcampaignsListCampaigns, "list-campaigns", "", false, "List Campaigns")
	_connectcampaignsCmd.Flags().BoolVarP(&_connectcampaignsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_connectcampaignsCmd.Flags().BoolVarP(&_connectcampaignsPauseCampaign, "pause-campaign", "", false, "Pause Campaign")
	_connectcampaignsCmd.Flags().BoolVarP(&_connectcampaignsPutDialRequestBatch, "put-dial-request-batch", "", false, "Put Dial Request Batch")
	_connectcampaignsCmd.Flags().BoolVarP(&_connectcampaignsResumeCampaign, "resume-campaign", "", false, "Resume Campaign")
	_connectcampaignsCmd.Flags().BoolVarP(&_connectcampaignsStartCampaign, "start-campaign", "", false, "Start Campaign")
	_connectcampaignsCmd.Flags().BoolVarP(&_connectcampaignsStartInstanceOnboardingJob, "start-instance-onboarding-job", "", false, "Start Instance Onboarding Job")
	_connectcampaignsCmd.Flags().BoolVarP(&_connectcampaignsStopCampaign, "stop-campaign", "", false, "Stop Campaign")
	_connectcampaignsCmd.Flags().BoolVarP(&_connectcampaignsTagResource, "tag-resource", "", false, "Tag Resource")
	_connectcampaignsCmd.Flags().BoolVarP(&_connectcampaignsUntagResource, "untag-resource", "", false, "Untag Resource")
	_connectcampaignsCmd.Flags().BoolVarP(&_connectcampaignsUpdateCampaignDialerConfig, "update-campaign-dialer-config", "", false, "Update Campaign Dialer Config")
	_connectcampaignsCmd.Flags().BoolVarP(&_connectcampaignsUpdateCampaignName, "update-campaign-name", "", false, "Update Campaign Name")
	_connectcampaignsCmd.Flags().BoolVarP(&_connectcampaignsUpdateCampaignOutboundCallConfig, "update-campaign-outbound-call-config", "", false, "Update Campaign Outbound Call Config")

}
