package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// pinpointCmd represents the pinpoint command
var _pinpointCmd = &cobra.Command{
	Use:   "pinpoint",
	Short: "AWS pinpoint CLI",
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
		client := pinpoint.NewFromConfig(cfg)
		if _pinpointCreateApp {
			pinpoint_CreateApp(cfg, client)
			return
		}
		if _pinpointCreateCampaign {
			pinpoint_CreateCampaign(cfg, client)
			return
		}
		if _pinpointCreateEmailTemplate {
			pinpoint_CreateEmailTemplate(cfg, client)
			return
		}
		if _pinpointCreateExportJob {
			pinpoint_CreateExportJob(cfg, client)
			return
		}
		if _pinpointCreateImportJob {
			pinpoint_CreateImportJob(cfg, client)
			return
		}
		if _pinpointCreateInAppTemplate {
			pinpoint_CreateInAppTemplate(cfg, client)
			return
		}
		if _pinpointCreateJourney {
			pinpoint_CreateJourney(cfg, client)
			return
		}
		if _pinpointCreatePushTemplate {
			pinpoint_CreatePushTemplate(cfg, client)
			return
		}
		if _pinpointCreateRecommenderConfiguration {
			pinpoint_CreateRecommenderConfiguration(cfg, client)
			return
		}
		if _pinpointCreateSegment {
			pinpoint_CreateSegment(cfg, client)
			return
		}
		if _pinpointCreateSmsTemplate {
			pinpoint_CreateSmsTemplate(cfg, client)
			return
		}
		if _pinpointCreateVoiceTemplate {
			pinpoint_CreateVoiceTemplate(cfg, client)
			return
		}
		if _pinpointDeleteAdmChannel {
			pinpoint_DeleteAdmChannel(cfg, client)
			return
		}
		if _pinpointDeleteApnsChannel {
			pinpoint_DeleteApnsChannel(cfg, client)
			return
		}
		if _pinpointDeleteApnsSandboxChannel {
			pinpoint_DeleteApnsSandboxChannel(cfg, client)
			return
		}
		if _pinpointDeleteApnsVoipChannel {
			pinpoint_DeleteApnsVoipChannel(cfg, client)
			return
		}
		if _pinpointDeleteApnsVoipSandboxChannel {
			pinpoint_DeleteApnsVoipSandboxChannel(cfg, client)
			return
		}
		if _pinpointDeleteApp {
			pinpoint_DeleteApp(cfg, client)
			return
		}
		if _pinpointDeleteBaiduChannel {
			pinpoint_DeleteBaiduChannel(cfg, client)
			return
		}
		if _pinpointDeleteCampaign {
			pinpoint_DeleteCampaign(cfg, client)
			return
		}
		if _pinpointDeleteEmailChannel {
			pinpoint_DeleteEmailChannel(cfg, client)
			return
		}
		if _pinpointDeleteEmailTemplate {
			pinpoint_DeleteEmailTemplate(cfg, client)
			return
		}
		if _pinpointDeleteEndpoint {
			pinpoint_DeleteEndpoint(cfg, client)
			return
		}
		if _pinpointDeleteEventStream {
			pinpoint_DeleteEventStream(cfg, client)
			return
		}
		if _pinpointDeleteGcmChannel {
			pinpoint_DeleteGcmChannel(cfg, client)
			return
		}
		if _pinpointDeleteInAppTemplate {
			pinpoint_DeleteInAppTemplate(cfg, client)
			return
		}
		if _pinpointDeleteJourney {
			pinpoint_DeleteJourney(cfg, client)
			return
		}
		if _pinpointDeletePushTemplate {
			pinpoint_DeletePushTemplate(cfg, client)
			return
		}
		if _pinpointDeleteRecommenderConfiguration {
			pinpoint_DeleteRecommenderConfiguration(cfg, client)
			return
		}
		if _pinpointDeleteSegment {
			pinpoint_DeleteSegment(cfg, client)
			return
		}
		if _pinpointDeleteSmsChannel {
			pinpoint_DeleteSmsChannel(cfg, client)
			return
		}
		if _pinpointDeleteSmsTemplate {
			pinpoint_DeleteSmsTemplate(cfg, client)
			return
		}
		if _pinpointDeleteUserEndpoints {
			pinpoint_DeleteUserEndpoints(cfg, client)
			return
		}
		if _pinpointDeleteVoiceChannel {
			pinpoint_DeleteVoiceChannel(cfg, client)
			return
		}
		if _pinpointDeleteVoiceTemplate {
			pinpoint_DeleteVoiceTemplate(cfg, client)
			return
		}
		if _pinpointGetAdmChannel {
			pinpoint_GetAdmChannel(cfg, client)
			return
		}
		if _pinpointGetApnsChannel {
			pinpoint_GetApnsChannel(cfg, client)
			return
		}
		if _pinpointGetApnsSandboxChannel {
			pinpoint_GetApnsSandboxChannel(cfg, client)
			return
		}
		if _pinpointGetApnsVoipChannel {
			pinpoint_GetApnsVoipChannel(cfg, client)
			return
		}
		if _pinpointGetApnsVoipSandboxChannel {
			pinpoint_GetApnsVoipSandboxChannel(cfg, client)
			return
		}
		if _pinpointGetApp {
			pinpoint_GetApp(cfg, client)
			return
		}
		if _pinpointGetApplicationDateRangeKpi {
			pinpoint_GetApplicationDateRangeKpi(cfg, client)
			return
		}
		if _pinpointGetApplicationSettings {
			pinpoint_GetApplicationSettings(cfg, client)
			return
		}
		if _pinpointGetApps {
			pinpoint_GetApps(cfg, client)
			return
		}
		if _pinpointGetBaiduChannel {
			pinpoint_GetBaiduChannel(cfg, client)
			return
		}
		if _pinpointGetCampaign {
			pinpoint_GetCampaign(cfg, client)
			return
		}
		if _pinpointGetCampaignActivities {
			pinpoint_GetCampaignActivities(cfg, client)
			return
		}
		if _pinpointGetCampaignDateRangeKpi {
			pinpoint_GetCampaignDateRangeKpi(cfg, client)
			return
		}
		if _pinpointGetCampaignVersion {
			pinpoint_GetCampaignVersion(cfg, client)
			return
		}
		if _pinpointGetCampaignVersions {
			pinpoint_GetCampaignVersions(cfg, client)
			return
		}
		if _pinpointGetCampaigns {
			pinpoint_GetCampaigns(cfg, client)
			return
		}
		if _pinpointGetChannels {
			pinpoint_GetChannels(cfg, client)
			return
		}
		if _pinpointGetEmailChannel {
			pinpoint_GetEmailChannel(cfg, client)
			return
		}
		if _pinpointGetEmailTemplate {
			pinpoint_GetEmailTemplate(cfg, client)
			return
		}
		if _pinpointGetEndpoint {
			pinpoint_GetEndpoint(cfg, client)
			return
		}
		if _pinpointGetEventStream {
			pinpoint_GetEventStream(cfg, client)
			return
		}
		if _pinpointGetExportJob {
			pinpoint_GetExportJob(cfg, client)
			return
		}
		if _pinpointGetExportJobs {
			pinpoint_GetExportJobs(cfg, client)
			return
		}
		if _pinpointGetGcmChannel {
			pinpoint_GetGcmChannel(cfg, client)
			return
		}
		if _pinpointGetImportJob {
			pinpoint_GetImportJob(cfg, client)
			return
		}
		if _pinpointGetImportJobs {
			pinpoint_GetImportJobs(cfg, client)
			return
		}
		if _pinpointGetInAppMessages {
			pinpoint_GetInAppMessages(cfg, client)
			return
		}
		if _pinpointGetInAppTemplate {
			pinpoint_GetInAppTemplate(cfg, client)
			return
		}
		if _pinpointGetJourney {
			pinpoint_GetJourney(cfg, client)
			return
		}
		if _pinpointGetJourneyDateRangeKpi {
			pinpoint_GetJourneyDateRangeKpi(cfg, client)
			return
		}
		if _pinpointGetJourneyExecutionActivityMetrics {
			pinpoint_GetJourneyExecutionActivityMetrics(cfg, client)
			return
		}
		if _pinpointGetJourneyExecutionMetrics {
			pinpoint_GetJourneyExecutionMetrics(cfg, client)
			return
		}
		if _pinpointGetJourneyRunExecutionActivityMetrics {
			pinpoint_GetJourneyRunExecutionActivityMetrics(cfg, client)
			return
		}
		if _pinpointGetJourneyRunExecutionMetrics {
			pinpoint_GetJourneyRunExecutionMetrics(cfg, client)
			return
		}
		if _pinpointGetJourneyRuns {
			pinpoint_GetJourneyRuns(cfg, client)
			return
		}
		if _pinpointGetPushTemplate {
			pinpoint_GetPushTemplate(cfg, client)
			return
		}
		if _pinpointGetRecommenderConfiguration {
			pinpoint_GetRecommenderConfiguration(cfg, client)
			return
		}
		if _pinpointGetRecommenderConfigurations {
			pinpoint_GetRecommenderConfigurations(cfg, client)
			return
		}
		if _pinpointGetSegment {
			pinpoint_GetSegment(cfg, client)
			return
		}
		if _pinpointGetSegmentExportJobs {
			pinpoint_GetSegmentExportJobs(cfg, client)
			return
		}
		if _pinpointGetSegmentImportJobs {
			pinpoint_GetSegmentImportJobs(cfg, client)
			return
		}
		if _pinpointGetSegmentVersion {
			pinpoint_GetSegmentVersion(cfg, client)
			return
		}
		if _pinpointGetSegmentVersions {
			pinpoint_GetSegmentVersions(cfg, client)
			return
		}
		if _pinpointGetSegments {
			pinpoint_GetSegments(cfg, client)
			return
		}
		if _pinpointGetSmsChannel {
			pinpoint_GetSmsChannel(cfg, client)
			return
		}
		if _pinpointGetSmsTemplate {
			pinpoint_GetSmsTemplate(cfg, client)
			return
		}
		if _pinpointGetUserEndpoints {
			pinpoint_GetUserEndpoints(cfg, client)
			return
		}
		if _pinpointGetVoiceChannel {
			pinpoint_GetVoiceChannel(cfg, client)
			return
		}
		if _pinpointGetVoiceTemplate {
			pinpoint_GetVoiceTemplate(cfg, client)
			return
		}
		if _pinpointListJourneys {
			pinpoint_ListJourneys(cfg, client)
			return
		}
		if _pinpointListTagsForResource {
			pinpoint_ListTagsForResource(cfg, client)
			return
		}
		if _pinpointListTemplateVersions {
			pinpoint_ListTemplateVersions(cfg, client)
			return
		}
		if _pinpointListTemplates {
			pinpoint_ListTemplates(cfg, client)
			return
		}
		if _pinpointPhoneNumberValidate {
			pinpoint_PhoneNumberValidate(cfg, client)
			return
		}
		if _pinpointPutEventStream {
			pinpoint_PutEventStream(cfg, client)
			return
		}
		if _pinpointPutEvents {
			pinpoint_PutEvents(cfg, client)
			return
		}
		if _pinpointRemoveAttributes {
			pinpoint_RemoveAttributes(cfg, client)
			return
		}
		if _pinpointSendMessages {
			pinpoint_SendMessages(cfg, client)
			return
		}
		if _pinpointSendOTPMessage {
			pinpoint_SendOTPMessage(cfg, client)
			return
		}
		if _pinpointSendUsersMessages {
			pinpoint_SendUsersMessages(cfg, client)
			return
		}
		if _pinpointTagResource {
			pinpoint_TagResource(cfg, client)
			return
		}
		if _pinpointUntagResource {
			pinpoint_UntagResource(cfg, client)
			return
		}
		if _pinpointUpdateAdmChannel {
			pinpoint_UpdateAdmChannel(cfg, client)
			return
		}
		if _pinpointUpdateApnsChannel {
			pinpoint_UpdateApnsChannel(cfg, client)
			return
		}
		if _pinpointUpdateApnsSandboxChannel {
			pinpoint_UpdateApnsSandboxChannel(cfg, client)
			return
		}
		if _pinpointUpdateApnsVoipChannel {
			pinpoint_UpdateApnsVoipChannel(cfg, client)
			return
		}
		if _pinpointUpdateApnsVoipSandboxChannel {
			pinpoint_UpdateApnsVoipSandboxChannel(cfg, client)
			return
		}
		if _pinpointUpdateApplicationSettings {
			pinpoint_UpdateApplicationSettings(cfg, client)
			return
		}
		if _pinpointUpdateBaiduChannel {
			pinpoint_UpdateBaiduChannel(cfg, client)
			return
		}
		if _pinpointUpdateCampaign {
			pinpoint_UpdateCampaign(cfg, client)
			return
		}
		if _pinpointUpdateEmailChannel {
			pinpoint_UpdateEmailChannel(cfg, client)
			return
		}
		if _pinpointUpdateEmailTemplate {
			pinpoint_UpdateEmailTemplate(cfg, client)
			return
		}
		if _pinpointUpdateEndpoint {
			pinpoint_UpdateEndpoint(cfg, client)
			return
		}
		if _pinpointUpdateEndpointsBatch {
			pinpoint_UpdateEndpointsBatch(cfg, client)
			return
		}
		if _pinpointUpdateGcmChannel {
			pinpoint_UpdateGcmChannel(cfg, client)
			return
		}
		if _pinpointUpdateInAppTemplate {
			pinpoint_UpdateInAppTemplate(cfg, client)
			return
		}
		if _pinpointUpdateJourney {
			pinpoint_UpdateJourney(cfg, client)
			return
		}
		if _pinpointUpdateJourneyState {
			pinpoint_UpdateJourneyState(cfg, client)
			return
		}
		if _pinpointUpdatePushTemplate {
			pinpoint_UpdatePushTemplate(cfg, client)
			return
		}
		if _pinpointUpdateRecommenderConfiguration {
			pinpoint_UpdateRecommenderConfiguration(cfg, client)
			return
		}
		if _pinpointUpdateSegment {
			pinpoint_UpdateSegment(cfg, client)
			return
		}
		if _pinpointUpdateSmsChannel {
			pinpoint_UpdateSmsChannel(cfg, client)
			return
		}
		if _pinpointUpdateSmsTemplate {
			pinpoint_UpdateSmsTemplate(cfg, client)
			return
		}
		if _pinpointUpdateTemplateActiveVersion {
			pinpoint_UpdateTemplateActiveVersion(cfg, client)
			return
		}
		if _pinpointUpdateVoiceChannel {
			pinpoint_UpdateVoiceChannel(cfg, client)
			return
		}
		if _pinpointUpdateVoiceTemplate {
			pinpoint_UpdateVoiceTemplate(cfg, client)
			return
		}
		if _pinpointVerifyOTPMessage {
			pinpoint_VerifyOTPMessage(cfg, client)
			return
		}

	},
}

var (
	_pinpointCreateApp                             bool
	_pinpointCreateCampaign                        bool
	_pinpointCreateEmailTemplate                   bool
	_pinpointCreateExportJob                       bool
	_pinpointCreateImportJob                       bool
	_pinpointCreateInAppTemplate                   bool
	_pinpointCreateJourney                         bool
	_pinpointCreatePushTemplate                    bool
	_pinpointCreateRecommenderConfiguration        bool
	_pinpointCreateSegment                         bool
	_pinpointCreateSmsTemplate                     bool
	_pinpointCreateVoiceTemplate                   bool
	_pinpointDeleteAdmChannel                      bool
	_pinpointDeleteApnsChannel                     bool
	_pinpointDeleteApnsSandboxChannel              bool
	_pinpointDeleteApnsVoipChannel                 bool
	_pinpointDeleteApnsVoipSandboxChannel          bool
	_pinpointDeleteApp                             bool
	_pinpointDeleteBaiduChannel                    bool
	_pinpointDeleteCampaign                        bool
	_pinpointDeleteEmailChannel                    bool
	_pinpointDeleteEmailTemplate                   bool
	_pinpointDeleteEndpoint                        bool
	_pinpointDeleteEventStream                     bool
	_pinpointDeleteGcmChannel                      bool
	_pinpointDeleteInAppTemplate                   bool
	_pinpointDeleteJourney                         bool
	_pinpointDeletePushTemplate                    bool
	_pinpointDeleteRecommenderConfiguration        bool
	_pinpointDeleteSegment                         bool
	_pinpointDeleteSmsChannel                      bool
	_pinpointDeleteSmsTemplate                     bool
	_pinpointDeleteUserEndpoints                   bool
	_pinpointDeleteVoiceChannel                    bool
	_pinpointDeleteVoiceTemplate                   bool
	_pinpointGetAdmChannel                         bool
	_pinpointGetApnsChannel                        bool
	_pinpointGetApnsSandboxChannel                 bool
	_pinpointGetApnsVoipChannel                    bool
	_pinpointGetApnsVoipSandboxChannel             bool
	_pinpointGetApp                                bool
	_pinpointGetApplicationDateRangeKpi            bool
	_pinpointGetApplicationSettings                bool
	_pinpointGetApps                               bool
	_pinpointGetBaiduChannel                       bool
	_pinpointGetCampaign                           bool
	_pinpointGetCampaignActivities                 bool
	_pinpointGetCampaignDateRangeKpi               bool
	_pinpointGetCampaignVersion                    bool
	_pinpointGetCampaignVersions                   bool
	_pinpointGetCampaigns                          bool
	_pinpointGetChannels                           bool
	_pinpointGetEmailChannel                       bool
	_pinpointGetEmailTemplate                      bool
	_pinpointGetEndpoint                           bool
	_pinpointGetEventStream                        bool
	_pinpointGetExportJob                          bool
	_pinpointGetExportJobs                         bool
	_pinpointGetGcmChannel                         bool
	_pinpointGetImportJob                          bool
	_pinpointGetImportJobs                         bool
	_pinpointGetInAppMessages                      bool
	_pinpointGetInAppTemplate                      bool
	_pinpointGetJourney                            bool
	_pinpointGetJourneyDateRangeKpi                bool
	_pinpointGetJourneyExecutionActivityMetrics    bool
	_pinpointGetJourneyExecutionMetrics            bool
	_pinpointGetJourneyRunExecutionActivityMetrics bool
	_pinpointGetJourneyRunExecutionMetrics         bool
	_pinpointGetJourneyRuns                        bool
	_pinpointGetPushTemplate                       bool
	_pinpointGetRecommenderConfiguration           bool
	_pinpointGetRecommenderConfigurations          bool
	_pinpointGetSegment                            bool
	_pinpointGetSegmentExportJobs                  bool
	_pinpointGetSegmentImportJobs                  bool
	_pinpointGetSegmentVersion                     bool
	_pinpointGetSegmentVersions                    bool
	_pinpointGetSegments                           bool
	_pinpointGetSmsChannel                         bool
	_pinpointGetSmsTemplate                        bool
	_pinpointGetUserEndpoints                      bool
	_pinpointGetVoiceChannel                       bool
	_pinpointGetVoiceTemplate                      bool
	_pinpointListJourneys                          bool
	_pinpointListTagsForResource                   bool
	_pinpointListTemplateVersions                  bool
	_pinpointListTemplates                         bool
	_pinpointPhoneNumberValidate                   bool
	_pinpointPutEventStream                        bool
	_pinpointPutEvents                             bool
	_pinpointRemoveAttributes                      bool
	_pinpointSendMessages                          bool
	_pinpointSendOTPMessage                        bool
	_pinpointSendUsersMessages                     bool
	_pinpointTagResource                           bool
	_pinpointUntagResource                         bool
	_pinpointUpdateAdmChannel                      bool
	_pinpointUpdateApnsChannel                     bool
	_pinpointUpdateApnsSandboxChannel              bool
	_pinpointUpdateApnsVoipChannel                 bool
	_pinpointUpdateApnsVoipSandboxChannel          bool
	_pinpointUpdateApplicationSettings             bool
	_pinpointUpdateBaiduChannel                    bool
	_pinpointUpdateCampaign                        bool
	_pinpointUpdateEmailChannel                    bool
	_pinpointUpdateEmailTemplate                   bool
	_pinpointUpdateEndpoint                        bool
	_pinpointUpdateEndpointsBatch                  bool
	_pinpointUpdateGcmChannel                      bool
	_pinpointUpdateInAppTemplate                   bool
	_pinpointUpdateJourney                         bool
	_pinpointUpdateJourneyState                    bool
	_pinpointUpdatePushTemplate                    bool
	_pinpointUpdateRecommenderConfiguration        bool
	_pinpointUpdateSegment                         bool
	_pinpointUpdateSmsChannel                      bool
	_pinpointUpdateSmsTemplate                     bool
	_pinpointUpdateTemplateActiveVersion           bool
	_pinpointUpdateVoiceChannel                    bool
	_pinpointUpdateVoiceTemplate                   bool
	_pinpointVerifyOTPMessage                      bool

	_pinpointADMChannelRequest                 string
	_pinpointAPNSChannelRequest                string
	_pinpointAPNSSandboxChannelRequest         string
	_pinpointAPNSVoipChannelRequest            string
	_pinpointAPNSVoipSandboxChannelRequest     string
	_pinpointApplicationId                     string
	_pinpointAttributeType                     string
	_pinpointBaiduChannelRequest               string
	_pinpointCampaignId                        string
	_pinpointCreateApplicationRequest          string
	_pinpointCreateNewVersion                  string
	_pinpointEmailChannelRequest               string
	_pinpointEmailTemplateRequest              string
	_pinpointEndTime                           string
	_pinpointEndpointBatchRequest              string
	_pinpointEndpointId                        string
	_pinpointEndpointRequest                   string
	_pinpointEventsRequest                     string
	_pinpointExportJobRequest                  string
	_pinpointGCMChannelRequest                 string
	_pinpointImportJobRequest                  string
	_pinpointInAppTemplateRequest              string
	_pinpointJobId                             string
	_pinpointJourneyActivityId                 string
	_pinpointJourneyId                         string
	_pinpointJourneyStateRequest               string
	_pinpointKpiName                           string
	_pinpointMessageRequest                    string
	_pinpointNextToken                         string
	_pinpointNumberValidateRequest             string
	_pinpointPageSize                          string
	_pinpointPrefix                            string
	_pinpointPushNotificationTemplateRequest   string
	_pinpointRecommenderId                     string
	_pinpointResourceArn                       string
	_pinpointRunId                             string
	_pinpointSegmentId                         string
	_pinpointSendOTPMessageRequestParameters   string
	_pinpointSendUsersMessageRequest           string
	_pinpointSMSChannelRequest                 string
	_pinpointSMSTemplateRequest                string
	_pinpointStartTime                         string
	_pinpointTagKeys                           []string
	_pinpointTagsModel                         string
	_pinpointTemplateActiveVersionRequest      string
	_pinpointTemplateName                      string
	_pinpointTemplateType                      string
	_pinpointToken                             string
	_pinpointUpdateAttributesRequest           string
	_pinpointUserId                            string
	_pinpointVerifyOTPMessageRequestParameters string
	_pinpointVersion                           string
	_pinpointVoiceChannelRequest               string
	_pinpointVoiceTemplateRequest              string
	_pinpointWriteApplicationSettingsRequest   string
	_pinpointWriteCampaignRequest              string
	_pinpointWriteEventStream                  string
	_pinpointWriteJourneyRequest               string
	_pinpointWriteSegmentRequest               string
)

// Creates an application.
func pinpoint_CreateApp(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.CreateAppInput{
		// CreateApplicationRequest: *types.CreateApplicationRequest, // Required
	}

	if len(_pinpointCreateApplicationRequest) > 0 {
		if err := assignInputField(input, "CreateApplicationRequest", _pinpointCreateApplicationRequest); err != nil {
			log.Errorf("invalid --create-application-request: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new campaign for an application or updates the settings of an
// existing campaign for an application.
func pinpoint_CreateCampaign(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.CreateCampaignInput{
		// ApplicationId: *string, // Required
		// WriteCampaignRequest: *types.WriteCampaignRequest, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointWriteCampaignRequest) > 0 {
		if err := assignInputField(input, "WriteCampaignRequest", _pinpointWriteCampaignRequest); err != nil {
			log.Errorf("invalid --write-campaign-request: %s", err.Error())
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

// Creates a message template for messages that are sent through the email channel.
func pinpoint_CreateEmailTemplate(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.CreateEmailTemplateInput{
		// EmailTemplateRequest: *types.EmailTemplateRequest, // Required
		// TemplateName: *string, // Required
	}

	if len(_pinpointEmailTemplateRequest) > 0 {
		if err := assignInputField(input, "EmailTemplateRequest", _pinpointEmailTemplateRequest); err != nil {
			log.Errorf("invalid --email-template-request: %s", err.Error())
			return
		}
	}
	if len(_pinpointTemplateName) > 0 {
		input.TemplateName = aws.String(_pinpointTemplateName)
	}

	if resp, err := client.CreateEmailTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an export job for an application.
func pinpoint_CreateExportJob(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.CreateExportJobInput{
		// ApplicationId: *string, // Required
		// ExportJobRequest: *types.ExportJobRequest, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointExportJobRequest) > 0 {
		if err := assignInputField(input, "ExportJobRequest", _pinpointExportJobRequest); err != nil {
			log.Errorf("invalid --export-job-request: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an import job for an application.
func pinpoint_CreateImportJob(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.CreateImportJobInput{
		// ApplicationId: *string, // Required
		// ImportJobRequest: *types.ImportJobRequest, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointImportJobRequest) > 0 {
		if err := assignInputField(input, "ImportJobRequest", _pinpointImportJobRequest); err != nil {
			log.Errorf("invalid --import-job-request: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new message template for messages using the in-app message channel.
func pinpoint_CreateInAppTemplate(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.CreateInAppTemplateInput{
		// InAppTemplateRequest: *types.InAppTemplateRequest, // Required
		// TemplateName: *string, // Required
	}

	if len(_pinpointInAppTemplateRequest) > 0 {
		if err := assignInputField(input, "InAppTemplateRequest", _pinpointInAppTemplateRequest); err != nil {
			log.Errorf("invalid --in-app-template-request: %s", err.Error())
			return
		}
	}
	if len(_pinpointTemplateName) > 0 {
		input.TemplateName = aws.String(_pinpointTemplateName)
	}

	if resp, err := client.CreateInAppTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a journey for an application.
func pinpoint_CreateJourney(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.CreateJourneyInput{
		// ApplicationId: *string, // Required
		// WriteJourneyRequest: *types.WriteJourneyRequest, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointWriteJourneyRequest) > 0 {
		if err := assignInputField(input, "WriteJourneyRequest", _pinpointWriteJourneyRequest); err != nil {
			log.Errorf("invalid --write-journey-request: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateJourney(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a message template for messages that are sent through a push
// notification channel.
func pinpoint_CreatePushTemplate(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.CreatePushTemplateInput{
		// PushNotificationTemplateRequest: *types.PushNotificationTemplateRequest, // Required
		// TemplateName: *string, // Required
	}

	if len(_pinpointPushNotificationTemplateRequest) > 0 {
		if err := assignInputField(input, "PushNotificationTemplateRequest", _pinpointPushNotificationTemplateRequest); err != nil {
			log.Errorf("invalid --push-notification-template-request: %s", err.Error())
			return
		}
	}
	if len(_pinpointTemplateName) > 0 {
		input.TemplateName = aws.String(_pinpointTemplateName)
	}

	if resp, err := client.CreatePushTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Pinpoint configuration for a recommender model.
func pinpoint_CreateRecommenderConfiguration(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.CreateRecommenderConfigurationInput{
		// CreateRecommenderConfiguration: *types.CreateRecommenderConfigurationShape, // Required
	}

	if resp, err := client.CreateRecommenderConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new segment for an application or updates the configuration,
// dimension, and other settings for an existing segment that's associated with an
// application.
func pinpoint_CreateSegment(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.CreateSegmentInput{
		// ApplicationId: *string, // Required
		// WriteSegmentRequest: *types.WriteSegmentRequest, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointWriteSegmentRequest) > 0 {
		if err := assignInputField(input, "WriteSegmentRequest", _pinpointWriteSegmentRequest); err != nil {
			log.Errorf("invalid --write-segment-request: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSegment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a message template for messages that are sent through the SMS channel.
func pinpoint_CreateSmsTemplate(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.CreateSmsTemplateInput{
		// SMSTemplateRequest: *types.SMSTemplateRequest, // Required
		// TemplateName: *string, // Required
	}

	if len(_pinpointSMSTemplateRequest) > 0 {
		if err := assignInputField(input, "SMSTemplateRequest", _pinpointSMSTemplateRequest); err != nil {
			log.Errorf("invalid --sms-template-request: %s", err.Error())
			return
		}
	}
	if len(_pinpointTemplateName) > 0 {
		input.TemplateName = aws.String(_pinpointTemplateName)
	}

	if resp, err := client.CreateSmsTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a message template for messages that are sent through the voice channel.
func pinpoint_CreateVoiceTemplate(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.CreateVoiceTemplateInput{
		// TemplateName: *string, // Required
		// VoiceTemplateRequest: *types.VoiceTemplateRequest, // Required
	}

	if len(_pinpointTemplateName) > 0 {
		input.TemplateName = aws.String(_pinpointTemplateName)
	}
	if len(_pinpointVoiceTemplateRequest) > 0 {
		if err := assignInputField(input, "VoiceTemplateRequest", _pinpointVoiceTemplateRequest); err != nil {
			log.Errorf("invalid --voice-template-request: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVoiceTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the ADM channel for an application and deletes any existing settings
// for the channel.
func pinpoint_DeleteAdmChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.DeleteAdmChannelInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.DeleteAdmChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the APNs channel for an application and deletes any existing settings
// for the channel.
func pinpoint_DeleteApnsChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.DeleteApnsChannelInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.DeleteApnsChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the APNs sandbox channel for an application and deletes any existing
// settings for the channel.
func pinpoint_DeleteApnsSandboxChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.DeleteApnsSandboxChannelInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.DeleteApnsSandboxChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the APNs VoIP channel for an application and deletes any existing
// settings for the channel.
func pinpoint_DeleteApnsVoipChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.DeleteApnsVoipChannelInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.DeleteApnsVoipChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the APNs VoIP sandbox channel for an application and deletes any
// existing settings for the channel.
func pinpoint_DeleteApnsVoipSandboxChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.DeleteApnsVoipSandboxChannelInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.DeleteApnsVoipSandboxChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an application.
func pinpoint_DeleteApp(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.DeleteAppInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.DeleteApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the Baidu channel for an application and deletes any existing settings
// for the channel.
func pinpoint_DeleteBaiduChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.DeleteBaiduChannelInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.DeleteBaiduChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a campaign from an application.
func pinpoint_DeleteCampaign(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.DeleteCampaignInput{
		// ApplicationId: *string, // Required
		// CampaignId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointCampaignId) > 0 {
		input.CampaignId = aws.String(_pinpointCampaignId)
	}

	if resp, err := client.DeleteCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the email channel for an application and deletes any existing settings
// for the channel.
func pinpoint_DeleteEmailChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.DeleteEmailChannelInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.DeleteEmailChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a message template for messages that were sent through the email
// channel.
func pinpoint_DeleteEmailTemplate(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.DeleteEmailTemplateInput{
		// TemplateName: *string, // Required
	}

	if len(_pinpointTemplateName) > 0 {
		input.TemplateName = aws.String(_pinpointTemplateName)
	}
	if len(_pinpointVersion) > 0 {
		input.Version = aws.String(_pinpointVersion)
	}

	if resp, err := client.DeleteEmailTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an endpoint from an application.
func pinpoint_DeleteEndpoint(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.DeleteEndpointInput{
		// ApplicationId: *string, // Required
		// EndpointId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointEndpointId) > 0 {
		input.EndpointId = aws.String(_pinpointEndpointId)
	}

	if resp, err := client.DeleteEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the event stream for an application.
func pinpoint_DeleteEventStream(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.DeleteEventStreamInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.DeleteEventStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the GCM channel for an application and deletes any existing settings
// for the channel.
func pinpoint_DeleteGcmChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.DeleteGcmChannelInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.DeleteGcmChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a message template for messages sent using the in-app message channel.
func pinpoint_DeleteInAppTemplate(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.DeleteInAppTemplateInput{
		// TemplateName: *string, // Required
	}

	if len(_pinpointTemplateName) > 0 {
		input.TemplateName = aws.String(_pinpointTemplateName)
	}
	if len(_pinpointVersion) > 0 {
		input.Version = aws.String(_pinpointVersion)
	}

	if resp, err := client.DeleteInAppTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a journey from an application.
func pinpoint_DeleteJourney(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.DeleteJourneyInput{
		// ApplicationId: *string, // Required
		// JourneyId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointJourneyId) > 0 {
		input.JourneyId = aws.String(_pinpointJourneyId)
	}

	if resp, err := client.DeleteJourney(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a message template for messages that were sent through a push
// notification channel.
func pinpoint_DeletePushTemplate(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.DeletePushTemplateInput{
		// TemplateName: *string, // Required
	}

	if len(_pinpointTemplateName) > 0 {
		input.TemplateName = aws.String(_pinpointTemplateName)
	}
	if len(_pinpointVersion) > 0 {
		input.Version = aws.String(_pinpointVersion)
	}

	if resp, err := client.DeletePushTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Pinpoint configuration for a recommender model.
func pinpoint_DeleteRecommenderConfiguration(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.DeleteRecommenderConfigurationInput{
		// RecommenderId: *string, // Required
	}

	if len(_pinpointRecommenderId) > 0 {
		input.RecommenderId = aws.String(_pinpointRecommenderId)
	}

	if resp, err := client.DeleteRecommenderConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a segment from an application.
func pinpoint_DeleteSegment(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.DeleteSegmentInput{
		// ApplicationId: *string, // Required
		// SegmentId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointSegmentId) > 0 {
		input.SegmentId = aws.String(_pinpointSegmentId)
	}

	if resp, err := client.DeleteSegment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the SMS channel for an application and deletes any existing settings
// for the channel.
func pinpoint_DeleteSmsChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.DeleteSmsChannelInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.DeleteSmsChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a message template for messages that were sent through the SMS channel.
func pinpoint_DeleteSmsTemplate(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.DeleteSmsTemplateInput{
		// TemplateName: *string, // Required
	}

	if len(_pinpointTemplateName) > 0 {
		input.TemplateName = aws.String(_pinpointTemplateName)
	}
	if len(_pinpointVersion) > 0 {
		input.Version = aws.String(_pinpointVersion)
	}

	if resp, err := client.DeleteSmsTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes all the endpoints that are associated with a specific user ID.
func pinpoint_DeleteUserEndpoints(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.DeleteUserEndpointsInput{
		// ApplicationId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointUserId) > 0 {
		input.UserId = aws.String(_pinpointUserId)
	}

	if resp, err := client.DeleteUserEndpoints(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the voice channel for an application and deletes any existing settings
// for the channel.
func pinpoint_DeleteVoiceChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.DeleteVoiceChannelInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.DeleteVoiceChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a message template for messages that were sent through the voice
// channel.
func pinpoint_DeleteVoiceTemplate(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.DeleteVoiceTemplateInput{
		// TemplateName: *string, // Required
	}

	if len(_pinpointTemplateName) > 0 {
		input.TemplateName = aws.String(_pinpointTemplateName)
	}
	if len(_pinpointVersion) > 0 {
		input.Version = aws.String(_pinpointVersion)
	}

	if resp, err := client.DeleteVoiceTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the status and settings of the ADM channel for an
// application.
func pinpoint_GetAdmChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetAdmChannelInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.GetAdmChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the status and settings of the APNs channel for an
// application.
func pinpoint_GetApnsChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetApnsChannelInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.GetApnsChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the status and settings of the APNs sandbox channel
// for an application.
func pinpoint_GetApnsSandboxChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetApnsSandboxChannelInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.GetApnsSandboxChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the status and settings of the APNs VoIP channel
// for an application.
func pinpoint_GetApnsVoipChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetApnsVoipChannelInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.GetApnsVoipChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the status and settings of the APNs VoIP sandbox
// channel for an application.
func pinpoint_GetApnsVoipSandboxChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetApnsVoipSandboxChannelInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.GetApnsVoipSandboxChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an application.
func pinpoint_GetApp(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetAppInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.GetApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves (queries) pre-aggregated data for a standard metric that applies to
// an application.
func pinpoint_GetApplicationDateRangeKpi(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetApplicationDateRangeKpiInput{
		// ApplicationId: *string, // Required
		// KpiName: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointKpiName) > 0 {
		input.KpiName = aws.String(_pinpointKpiName)
	}
	if len(_pinpointEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _pinpointEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_pinpointNextToken) > 0 {
		input.NextToken = aws.String(_pinpointNextToken)
	}
	if len(_pinpointPageSize) > 0 {
		input.PageSize = aws.String(_pinpointPageSize)
	}
	if len(_pinpointStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _pinpointStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetApplicationDateRangeKpi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the settings for an application.
func pinpoint_GetApplicationSettings(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetApplicationSettingsInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.GetApplicationSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about all the applications that are associated with your
// Amazon Pinpoint account.
func pinpoint_GetApps(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetAppsInput{}

	if len(_pinpointPageSize) > 0 {
		input.PageSize = aws.String(_pinpointPageSize)
	}
	if len(_pinpointToken) > 0 {
		input.Token = aws.String(_pinpointToken)
	}

	if resp, err := client.GetApps(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the status and settings of the Baidu channel for an
// application.
func pinpoint_GetBaiduChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetBaiduChannelInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.GetBaiduChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the status, configuration, and other settings for a
// campaign.
func pinpoint_GetCampaign(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetCampaignInput{
		// ApplicationId: *string, // Required
		// CampaignId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointCampaignId) > 0 {
		input.CampaignId = aws.String(_pinpointCampaignId)
	}

	if resp, err := client.GetCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about all the activities for a campaign.
func pinpoint_GetCampaignActivities(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetCampaignActivitiesInput{
		// ApplicationId: *string, // Required
		// CampaignId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointCampaignId) > 0 {
		input.CampaignId = aws.String(_pinpointCampaignId)
	}
	if len(_pinpointPageSize) > 0 {
		input.PageSize = aws.String(_pinpointPageSize)
	}
	if len(_pinpointToken) > 0 {
		input.Token = aws.String(_pinpointToken)
	}

	if resp, err := client.GetCampaignActivities(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves (queries) pre-aggregated data for a standard metric that applies to a
// campaign.
func pinpoint_GetCampaignDateRangeKpi(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetCampaignDateRangeKpiInput{
		// ApplicationId: *string, // Required
		// CampaignId: *string, // Required
		// KpiName: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointCampaignId) > 0 {
		input.CampaignId = aws.String(_pinpointCampaignId)
	}
	if len(_pinpointKpiName) > 0 {
		input.KpiName = aws.String(_pinpointKpiName)
	}
	if len(_pinpointEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _pinpointEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_pinpointNextToken) > 0 {
		input.NextToken = aws.String(_pinpointNextToken)
	}
	if len(_pinpointPageSize) > 0 {
		input.PageSize = aws.String(_pinpointPageSize)
	}
	if len(_pinpointStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _pinpointStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetCampaignDateRangeKpi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the status, configuration, and other settings for a
// specific version of a campaign.
func pinpoint_GetCampaignVersion(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetCampaignVersionInput{
		// ApplicationId: *string, // Required
		// CampaignId: *string, // Required
		// Version: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointCampaignId) > 0 {
		input.CampaignId = aws.String(_pinpointCampaignId)
	}
	if len(_pinpointVersion) > 0 {
		input.Version = aws.String(_pinpointVersion)
	}

	if resp, err := client.GetCampaignVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the status, configuration, and other settings for
// all versions of a campaign.
func pinpoint_GetCampaignVersions(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetCampaignVersionsInput{
		// ApplicationId: *string, // Required
		// CampaignId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointCampaignId) > 0 {
		input.CampaignId = aws.String(_pinpointCampaignId)
	}
	if len(_pinpointPageSize) > 0 {
		input.PageSize = aws.String(_pinpointPageSize)
	}
	if len(_pinpointToken) > 0 {
		input.Token = aws.String(_pinpointToken)
	}

	if resp, err := client.GetCampaignVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the status, configuration, and other settings for
// all the campaigns that are associated with an application.
func pinpoint_GetCampaigns(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetCampaignsInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointPageSize) > 0 {
		input.PageSize = aws.String(_pinpointPageSize)
	}
	if len(_pinpointToken) > 0 {
		input.Token = aws.String(_pinpointToken)
	}

	if resp, err := client.GetCampaigns(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the history and status of each channel for an
// application.
func pinpoint_GetChannels(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetChannelsInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.GetChannels(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the status and settings of the email channel for an
// application.
func pinpoint_GetEmailChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetEmailChannelInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.GetEmailChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the content and settings of a message template for messages that are
// sent through the email channel.
func pinpoint_GetEmailTemplate(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetEmailTemplateInput{
		// TemplateName: *string, // Required
	}

	if len(_pinpointTemplateName) > 0 {
		input.TemplateName = aws.String(_pinpointTemplateName)
	}
	if len(_pinpointVersion) > 0 {
		input.Version = aws.String(_pinpointVersion)
	}

	if resp, err := client.GetEmailTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the settings and attributes of a specific endpoint
// for an application.
func pinpoint_GetEndpoint(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetEndpointInput{
		// ApplicationId: *string, // Required
		// EndpointId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointEndpointId) > 0 {
		input.EndpointId = aws.String(_pinpointEndpointId)
	}

	if resp, err := client.GetEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the event stream settings for an application.
func pinpoint_GetEventStream(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetEventStreamInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.GetEventStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the status and settings of a specific export job
// for an application.
func pinpoint_GetExportJob(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetExportJobInput{
		// ApplicationId: *string, // Required
		// JobId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointJobId) > 0 {
		input.JobId = aws.String(_pinpointJobId)
	}

	if resp, err := client.GetExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the status and settings of all the export jobs for
// an application.
func pinpoint_GetExportJobs(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetExportJobsInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointPageSize) > 0 {
		input.PageSize = aws.String(_pinpointPageSize)
	}
	if len(_pinpointToken) > 0 {
		input.Token = aws.String(_pinpointToken)
	}

	if resp, err := client.GetExportJobs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the status and settings of the GCM channel for an
// application.
func pinpoint_GetGcmChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetGcmChannelInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.GetGcmChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the status and settings of a specific import job
// for an application.
func pinpoint_GetImportJob(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetImportJobInput{
		// ApplicationId: *string, // Required
		// JobId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointJobId) > 0 {
		input.JobId = aws.String(_pinpointJobId)
	}

	if resp, err := client.GetImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the status and settings of all the import jobs for
// an application.
func pinpoint_GetImportJobs(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetImportJobsInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointPageSize) > 0 {
		input.PageSize = aws.String(_pinpointPageSize)
	}
	if len(_pinpointToken) > 0 {
		input.Token = aws.String(_pinpointToken)
	}

	if resp, err := client.GetImportJobs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the in-app messages targeted for the provided endpoint ID.
func pinpoint_GetInAppMessages(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetInAppMessagesInput{
		// ApplicationId: *string, // Required
		// EndpointId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointEndpointId) > 0 {
		input.EndpointId = aws.String(_pinpointEndpointId)
	}

	if resp, err := client.GetInAppMessages(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the content and settings of a message template for messages sent
// through the in-app channel.
func pinpoint_GetInAppTemplate(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetInAppTemplateInput{
		// TemplateName: *string, // Required
	}

	if len(_pinpointTemplateName) > 0 {
		input.TemplateName = aws.String(_pinpointTemplateName)
	}
	if len(_pinpointVersion) > 0 {
		input.Version = aws.String(_pinpointVersion)
	}

	if resp, err := client.GetInAppTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the status, configuration, and other settings for a
// journey.
func pinpoint_GetJourney(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetJourneyInput{
		// ApplicationId: *string, // Required
		// JourneyId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointJourneyId) > 0 {
		input.JourneyId = aws.String(_pinpointJourneyId)
	}

	if resp, err := client.GetJourney(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves (queries) pre-aggregated data for a standard engagement metric that
// applies to a journey.
func pinpoint_GetJourneyDateRangeKpi(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetJourneyDateRangeKpiInput{
		// ApplicationId: *string, // Required
		// JourneyId: *string, // Required
		// KpiName: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointJourneyId) > 0 {
		input.JourneyId = aws.String(_pinpointJourneyId)
	}
	if len(_pinpointKpiName) > 0 {
		input.KpiName = aws.String(_pinpointKpiName)
	}
	if len(_pinpointEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _pinpointEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_pinpointNextToken) > 0 {
		input.NextToken = aws.String(_pinpointNextToken)
	}
	if len(_pinpointPageSize) > 0 {
		input.PageSize = aws.String(_pinpointPageSize)
	}
	if len(_pinpointStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _pinpointStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetJourneyDateRangeKpi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves (queries) pre-aggregated data for a standard execution metric that
// applies to a journey activity.
func pinpoint_GetJourneyExecutionActivityMetrics(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetJourneyExecutionActivityMetricsInput{
		// ApplicationId: *string, // Required
		// JourneyActivityId: *string, // Required
		// JourneyId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointJourneyActivityId) > 0 {
		input.JourneyActivityId = aws.String(_pinpointJourneyActivityId)
	}
	if len(_pinpointJourneyId) > 0 {
		input.JourneyId = aws.String(_pinpointJourneyId)
	}
	if len(_pinpointNextToken) > 0 {
		input.NextToken = aws.String(_pinpointNextToken)
	}
	if len(_pinpointPageSize) > 0 {
		input.PageSize = aws.String(_pinpointPageSize)
	}

	if resp, err := client.GetJourneyExecutionActivityMetrics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves (queries) pre-aggregated data for a standard execution metric that
// applies to a journey.
func pinpoint_GetJourneyExecutionMetrics(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetJourneyExecutionMetricsInput{
		// ApplicationId: *string, // Required
		// JourneyId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointJourneyId) > 0 {
		input.JourneyId = aws.String(_pinpointJourneyId)
	}
	if len(_pinpointNextToken) > 0 {
		input.NextToken = aws.String(_pinpointNextToken)
	}
	if len(_pinpointPageSize) > 0 {
		input.PageSize = aws.String(_pinpointPageSize)
	}

	if resp, err := client.GetJourneyExecutionMetrics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves (queries) pre-aggregated data for a standard run execution metric
// that applies to a journey activity.
func pinpoint_GetJourneyRunExecutionActivityMetrics(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetJourneyRunExecutionActivityMetricsInput{
		// ApplicationId: *string, // Required
		// JourneyActivityId: *string, // Required
		// JourneyId: *string, // Required
		// RunId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointJourneyActivityId) > 0 {
		input.JourneyActivityId = aws.String(_pinpointJourneyActivityId)
	}
	if len(_pinpointJourneyId) > 0 {
		input.JourneyId = aws.String(_pinpointJourneyId)
	}
	if len(_pinpointRunId) > 0 {
		input.RunId = aws.String(_pinpointRunId)
	}
	if len(_pinpointNextToken) > 0 {
		input.NextToken = aws.String(_pinpointNextToken)
	}
	if len(_pinpointPageSize) > 0 {
		input.PageSize = aws.String(_pinpointPageSize)
	}

	if resp, err := client.GetJourneyRunExecutionActivityMetrics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves (queries) pre-aggregated data for a standard run execution metric
// that applies to a journey.
func pinpoint_GetJourneyRunExecutionMetrics(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetJourneyRunExecutionMetricsInput{
		// ApplicationId: *string, // Required
		// JourneyId: *string, // Required
		// RunId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointJourneyId) > 0 {
		input.JourneyId = aws.String(_pinpointJourneyId)
	}
	if len(_pinpointRunId) > 0 {
		input.RunId = aws.String(_pinpointRunId)
	}
	if len(_pinpointNextToken) > 0 {
		input.NextToken = aws.String(_pinpointNextToken)
	}
	if len(_pinpointPageSize) > 0 {
		input.PageSize = aws.String(_pinpointPageSize)
	}

	if resp, err := client.GetJourneyRunExecutionMetrics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about the runs of a journey.
func pinpoint_GetJourneyRuns(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetJourneyRunsInput{
		// ApplicationId: *string, // Required
		// JourneyId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointJourneyId) > 0 {
		input.JourneyId = aws.String(_pinpointJourneyId)
	}
	if len(_pinpointPageSize) > 0 {
		input.PageSize = aws.String(_pinpointPageSize)
	}
	if len(_pinpointToken) > 0 {
		input.Token = aws.String(_pinpointToken)
	}

	if resp, err := client.GetJourneyRuns(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the content and settings of a message template for messages that are
// sent through a push notification channel.
func pinpoint_GetPushTemplate(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetPushTemplateInput{
		// TemplateName: *string, // Required
	}

	if len(_pinpointTemplateName) > 0 {
		input.TemplateName = aws.String(_pinpointTemplateName)
	}
	if len(_pinpointVersion) > 0 {
		input.Version = aws.String(_pinpointVersion)
	}

	if resp, err := client.GetPushTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an Amazon Pinpoint configuration for a recommender
// model.
func pinpoint_GetRecommenderConfiguration(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetRecommenderConfigurationInput{
		// RecommenderId: *string, // Required
	}

	if len(_pinpointRecommenderId) > 0 {
		input.RecommenderId = aws.String(_pinpointRecommenderId)
	}

	if resp, err := client.GetRecommenderConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about all the recommender model configurations that are
// associated with your Amazon Pinpoint account.
func pinpoint_GetRecommenderConfigurations(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetRecommenderConfigurationsInput{}

	if len(_pinpointPageSize) > 0 {
		input.PageSize = aws.String(_pinpointPageSize)
	}
	if len(_pinpointToken) > 0 {
		input.Token = aws.String(_pinpointToken)
	}

	if resp, err := client.GetRecommenderConfigurations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the configuration, dimension, and other settings
// for a specific segment that's associated with an application.
func pinpoint_GetSegment(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetSegmentInput{
		// ApplicationId: *string, // Required
		// SegmentId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointSegmentId) > 0 {
		input.SegmentId = aws.String(_pinpointSegmentId)
	}

	if resp, err := client.GetSegment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the status and settings of the export jobs for a
// segment.
func pinpoint_GetSegmentExportJobs(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetSegmentExportJobsInput{
		// ApplicationId: *string, // Required
		// SegmentId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointSegmentId) > 0 {
		input.SegmentId = aws.String(_pinpointSegmentId)
	}
	if len(_pinpointPageSize) > 0 {
		input.PageSize = aws.String(_pinpointPageSize)
	}
	if len(_pinpointToken) > 0 {
		input.Token = aws.String(_pinpointToken)
	}

	if resp, err := client.GetSegmentExportJobs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the status and settings of the import jobs for a
// segment.
func pinpoint_GetSegmentImportJobs(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetSegmentImportJobsInput{
		// ApplicationId: *string, // Required
		// SegmentId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointSegmentId) > 0 {
		input.SegmentId = aws.String(_pinpointSegmentId)
	}
	if len(_pinpointPageSize) > 0 {
		input.PageSize = aws.String(_pinpointPageSize)
	}
	if len(_pinpointToken) > 0 {
		input.Token = aws.String(_pinpointToken)
	}

	if resp, err := client.GetSegmentImportJobs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the configuration, dimension, and other settings
// for a specific version of a segment that's associated with an application.
func pinpoint_GetSegmentVersion(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetSegmentVersionInput{
		// ApplicationId: *string, // Required
		// SegmentId: *string, // Required
		// Version: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointSegmentId) > 0 {
		input.SegmentId = aws.String(_pinpointSegmentId)
	}
	if len(_pinpointVersion) > 0 {
		input.Version = aws.String(_pinpointVersion)
	}

	if resp, err := client.GetSegmentVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the configuration, dimension, and other settings
// for all the versions of a specific segment that's associated with an
// application.
func pinpoint_GetSegmentVersions(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetSegmentVersionsInput{
		// ApplicationId: *string, // Required
		// SegmentId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointSegmentId) > 0 {
		input.SegmentId = aws.String(_pinpointSegmentId)
	}
	if len(_pinpointPageSize) > 0 {
		input.PageSize = aws.String(_pinpointPageSize)
	}
	if len(_pinpointToken) > 0 {
		input.Token = aws.String(_pinpointToken)
	}

	if resp, err := client.GetSegmentVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the configuration, dimension, and other settings
// for all the segments that are associated with an application.
func pinpoint_GetSegments(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetSegmentsInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointPageSize) > 0 {
		input.PageSize = aws.String(_pinpointPageSize)
	}
	if len(_pinpointToken) > 0 {
		input.Token = aws.String(_pinpointToken)
	}

	if resp, err := client.GetSegments(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the status and settings of the SMS channel for an
// application.
func pinpoint_GetSmsChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetSmsChannelInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.GetSmsChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the content and settings of a message template for messages that are
// sent through the SMS channel.
func pinpoint_GetSmsTemplate(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetSmsTemplateInput{
		// TemplateName: *string, // Required
	}

	if len(_pinpointTemplateName) > 0 {
		input.TemplateName = aws.String(_pinpointTemplateName)
	}
	if len(_pinpointVersion) > 0 {
		input.Version = aws.String(_pinpointVersion)
	}

	if resp, err := client.GetSmsTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about all the endpoints that are associated with a
// specific user ID.
func pinpoint_GetUserEndpoints(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetUserEndpointsInput{
		// ApplicationId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointUserId) > 0 {
		input.UserId = aws.String(_pinpointUserId)
	}

	if resp, err := client.GetUserEndpoints(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the status and settings of the voice channel for an
// application.
func pinpoint_GetVoiceChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetVoiceChannelInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.GetVoiceChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the content and settings of a message template for messages that are
// sent through the voice channel.
func pinpoint_GetVoiceTemplate(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.GetVoiceTemplateInput{
		// TemplateName: *string, // Required
	}

	if len(_pinpointTemplateName) > 0 {
		input.TemplateName = aws.String(_pinpointTemplateName)
	}
	if len(_pinpointVersion) > 0 {
		input.Version = aws.String(_pinpointVersion)
	}

	if resp, err := client.GetVoiceTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the status, configuration, and other settings for
// all the journeys that are associated with an application.
func pinpoint_ListJourneys(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.ListJourneysInput{
		// ApplicationId: *string, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointPageSize) > 0 {
		input.PageSize = aws.String(_pinpointPageSize)
	}
	if len(_pinpointToken) > 0 {
		input.Token = aws.String(_pinpointToken)
	}

	if resp, err := client.ListJourneys(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves all the tags (keys and values) that are associated with an
// application, campaign, message template, or segment.
func pinpoint_ListTagsForResource(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_pinpointResourceArn) > 0 {
		input.ResourceArn = aws.String(_pinpointResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about all the versions of a specific message template.
func pinpoint_ListTemplateVersions(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.ListTemplateVersionsInput{
		// TemplateName: *string, // Required
		// TemplateType: *string, // Required
	}

	if len(_pinpointTemplateName) > 0 {
		input.TemplateName = aws.String(_pinpointTemplateName)
	}
	if len(_pinpointTemplateType) > 0 {
		input.TemplateType = aws.String(_pinpointTemplateType)
	}
	if len(_pinpointNextToken) > 0 {
		input.NextToken = aws.String(_pinpointNextToken)
	}
	if len(_pinpointPageSize) > 0 {
		input.PageSize = aws.String(_pinpointPageSize)
	}

	if resp, err := client.ListTemplateVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about all the message templates that are associated with
// your Amazon Pinpoint account.
func pinpoint_ListTemplates(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.ListTemplatesInput{}

	if len(_pinpointNextToken) > 0 {
		input.NextToken = aws.String(_pinpointNextToken)
	}
	if len(_pinpointPageSize) > 0 {
		input.PageSize = aws.String(_pinpointPageSize)
	}
	if len(_pinpointPrefix) > 0 {
		input.Prefix = aws.String(_pinpointPrefix)
	}
	if len(_pinpointTemplateType) > 0 {
		input.TemplateType = aws.String(_pinpointTemplateType)
	}

	if resp, err := client.ListTemplates(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a phone number.
func pinpoint_PhoneNumberValidate(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.PhoneNumberValidateInput{
		// NumberValidateRequest: *types.NumberValidateRequest, // Required
	}

	if len(_pinpointNumberValidateRequest) > 0 {
		if err := assignInputField(input, "NumberValidateRequest", _pinpointNumberValidateRequest); err != nil {
			log.Errorf("invalid --number-validate-request: %s", err.Error())
			return
		}
	}

	if resp, err := client.PhoneNumberValidate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new event stream for an application or updates the settings of an
// existing event stream for an application.
func pinpoint_PutEventStream(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.PutEventStreamInput{
		// ApplicationId: *string, // Required
		// WriteEventStream: *types.WriteEventStream, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointWriteEventStream) > 0 {
		if err := assignInputField(input, "WriteEventStream", _pinpointWriteEventStream); err != nil {
			log.Errorf("invalid --write-event-stream: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutEventStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new event to record for endpoints, or creates or updates endpoint
// data that existing events are associated with.
func pinpoint_PutEvents(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.PutEventsInput{
		// ApplicationId: *string, // Required
		// EventsRequest: *types.EventsRequest, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointEventsRequest) > 0 {
		if err := assignInputField(input, "EventsRequest", _pinpointEventsRequest); err != nil {
			log.Errorf("invalid --events-request: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutEvents(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes one or more custom attributes, of the same attribute type, from the
// application. Existing endpoints still have the attributes but Amazon Pinpoint
// will stop capturing new or changed values for these attributes.
func pinpoint_RemoveAttributes(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.RemoveAttributesInput{
		// ApplicationId: *string, // Required
		// AttributeType: *string, // Required
		// UpdateAttributesRequest: *types.UpdateAttributesRequest, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointAttributeType) > 0 {
		input.AttributeType = aws.String(_pinpointAttributeType)
	}
	if len(_pinpointUpdateAttributesRequest) > 0 {
		if err := assignInputField(input, "UpdateAttributesRequest", _pinpointUpdateAttributesRequest); err != nil {
			log.Errorf("invalid --update-attributes-request: %s", err.Error())
			return
		}
	}

	if resp, err := client.RemoveAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates and sends a direct message.
func pinpoint_SendMessages(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.SendMessagesInput{
		// ApplicationId: *string, // Required
		// MessageRequest: *types.MessageRequest, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointMessageRequest) > 0 {
		if err := assignInputField(input, "MessageRequest", _pinpointMessageRequest); err != nil {
			log.Errorf("invalid --message-request: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendMessages(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Send an OTP message
func pinpoint_SendOTPMessage(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.SendOTPMessageInput{
		// ApplicationId: *string, // Required
		// SendOTPMessageRequestParameters: *types.SendOTPMessageRequestParameters, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointSendOTPMessageRequestParameters) > 0 {
		if err := assignInputField(input, "SendOTPMessageRequestParameters", _pinpointSendOTPMessageRequestParameters); err != nil {
			log.Errorf("invalid --send-otp-message-request-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendOTPMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates and sends a message to a list of users.
func pinpoint_SendUsersMessages(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.SendUsersMessagesInput{
		// ApplicationId: *string, // Required
		// SendUsersMessageRequest: *types.SendUsersMessageRequest, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointSendUsersMessageRequest) > 0 {
		if err := assignInputField(input, "SendUsersMessageRequest", _pinpointSendUsersMessageRequest); err != nil {
			log.Errorf("invalid --send-users-message-request: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendUsersMessages(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more tags (keys and values) to an application, campaign, message
// template, or segment.
func pinpoint_TagResource(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.TagResourceInput{
		// ResourceArn: *string, // Required
		// TagsModel: *types.TagsModel, // Required
	}

	if len(_pinpointResourceArn) > 0 {
		input.ResourceArn = aws.String(_pinpointResourceArn)
	}
	if len(_pinpointTagsModel) > 0 {
		if err := assignInputField(input, "TagsModel", _pinpointTagsModel); err != nil {
			log.Errorf("invalid --tags-model: %s", err.Error())
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

// Removes one or more tags (keys and values) from an application, campaign,
// message template, or segment.
func pinpoint_UntagResource(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_pinpointResourceArn) > 0 {
		input.ResourceArn = aws.String(_pinpointResourceArn)
	}
	if len(_pinpointTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _pinpointTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the ADM channel for an application or updates the status and settings
// of the ADM channel for an application.
func pinpoint_UpdateAdmChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UpdateAdmChannelInput{
		// ADMChannelRequest: *types.ADMChannelRequest, // Required
		// ApplicationId: *string, // Required
	}

	if len(_pinpointADMChannelRequest) > 0 {
		if err := assignInputField(input, "ADMChannelRequest", _pinpointADMChannelRequest); err != nil {
			log.Errorf("invalid --adm-channel-request: %s", err.Error())
			return
		}
	}
	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.UpdateAdmChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the APNs channel for an application or updates the status and settings
// of the APNs channel for an application.
func pinpoint_UpdateApnsChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UpdateApnsChannelInput{
		// APNSChannelRequest: *types.APNSChannelRequest, // Required
		// ApplicationId: *string, // Required
	}

	if len(_pinpointAPNSChannelRequest) > 0 {
		if err := assignInputField(input, "APNSChannelRequest", _pinpointAPNSChannelRequest); err != nil {
			log.Errorf("invalid --apns-channel-request: %s", err.Error())
			return
		}
	}
	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.UpdateApnsChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the APNs sandbox channel for an application or updates the status and
// settings of the APNs sandbox channel for an application.
func pinpoint_UpdateApnsSandboxChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UpdateApnsSandboxChannelInput{
		// APNSSandboxChannelRequest: *types.APNSSandboxChannelRequest, // Required
		// ApplicationId: *string, // Required
	}

	if len(_pinpointAPNSSandboxChannelRequest) > 0 {
		if err := assignInputField(input, "APNSSandboxChannelRequest", _pinpointAPNSSandboxChannelRequest); err != nil {
			log.Errorf("invalid --apns-sandbox-channel-request: %s", err.Error())
			return
		}
	}
	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.UpdateApnsSandboxChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the APNs VoIP channel for an application or updates the status and
// settings of the APNs VoIP channel for an application.
func pinpoint_UpdateApnsVoipChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UpdateApnsVoipChannelInput{
		// APNSVoipChannelRequest: *types.APNSVoipChannelRequest, // Required
		// ApplicationId: *string, // Required
	}

	if len(_pinpointAPNSVoipChannelRequest) > 0 {
		if err := assignInputField(input, "APNSVoipChannelRequest", _pinpointAPNSVoipChannelRequest); err != nil {
			log.Errorf("invalid --apns-voip-channel-request: %s", err.Error())
			return
		}
	}
	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.UpdateApnsVoipChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the APNs VoIP sandbox channel for an application or updates the status
// and settings of the APNs VoIP sandbox channel for an application.
func pinpoint_UpdateApnsVoipSandboxChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UpdateApnsVoipSandboxChannelInput{
		// APNSVoipSandboxChannelRequest: *types.APNSVoipSandboxChannelRequest, // Required
		// ApplicationId: *string, // Required
	}

	if len(_pinpointAPNSVoipSandboxChannelRequest) > 0 {
		if err := assignInputField(input, "APNSVoipSandboxChannelRequest", _pinpointAPNSVoipSandboxChannelRequest); err != nil {
			log.Errorf("invalid --apns-voip-sandbox-channel-request: %s", err.Error())
			return
		}
	}
	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}

	if resp, err := client.UpdateApnsVoipSandboxChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the settings for an application.
func pinpoint_UpdateApplicationSettings(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UpdateApplicationSettingsInput{
		// ApplicationId: *string, // Required
		// WriteApplicationSettingsRequest: *types.WriteApplicationSettingsRequest, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointWriteApplicationSettingsRequest) > 0 {
		if err := assignInputField(input, "WriteApplicationSettingsRequest", _pinpointWriteApplicationSettingsRequest); err != nil {
			log.Errorf("invalid --write-application-settings-request: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateApplicationSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the Baidu channel for an application or updates the status and settings
// of the Baidu channel for an application.
func pinpoint_UpdateBaiduChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UpdateBaiduChannelInput{
		// ApplicationId: *string, // Required
		// BaiduChannelRequest: *types.BaiduChannelRequest, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointBaiduChannelRequest) > 0 {
		if err := assignInputField(input, "BaiduChannelRequest", _pinpointBaiduChannelRequest); err != nil {
			log.Errorf("invalid --baidu-channel-request: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBaiduChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration and other settings for a campaign.
func pinpoint_UpdateCampaign(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UpdateCampaignInput{
		// ApplicationId: *string, // Required
		// CampaignId: *string, // Required
		// WriteCampaignRequest: *types.WriteCampaignRequest, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointCampaignId) > 0 {
		input.CampaignId = aws.String(_pinpointCampaignId)
	}
	if len(_pinpointWriteCampaignRequest) > 0 {
		if err := assignInputField(input, "WriteCampaignRequest", _pinpointWriteCampaignRequest); err != nil {
			log.Errorf("invalid --write-campaign-request: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the email channel for an application or updates the status and settings
// of the email channel for an application.
func pinpoint_UpdateEmailChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UpdateEmailChannelInput{
		// ApplicationId: *string, // Required
		// EmailChannelRequest: *types.EmailChannelRequest, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointEmailChannelRequest) > 0 {
		if err := assignInputField(input, "EmailChannelRequest", _pinpointEmailChannelRequest); err != nil {
			log.Errorf("invalid --email-channel-request: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEmailChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing message template for messages that are sent through the
// email channel.
func pinpoint_UpdateEmailTemplate(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UpdateEmailTemplateInput{
		// EmailTemplateRequest: *types.EmailTemplateRequest, // Required
		// TemplateName: *string, // Required
	}

	if len(_pinpointEmailTemplateRequest) > 0 {
		if err := assignInputField(input, "EmailTemplateRequest", _pinpointEmailTemplateRequest); err != nil {
			log.Errorf("invalid --email-template-request: %s", err.Error())
			return
		}
	}
	if len(_pinpointTemplateName) > 0 {
		input.TemplateName = aws.String(_pinpointTemplateName)
	}
	if len(_pinpointCreateNewVersion) > 0 {
		if err := assignInputField(input, "CreateNewVersion", _pinpointCreateNewVersion); err != nil {
			log.Errorf("invalid --create-new-version: %s", err.Error())
			return
		}
	}
	if len(_pinpointVersion) > 0 {
		input.Version = aws.String(_pinpointVersion)
	}

	if resp, err := client.UpdateEmailTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new endpoint for an application or updates the settings and
// attributes of an existing endpoint for an application. You can also use this
// operation to define custom attributes for an endpoint. If an update includes one
// or more values for a custom attribute, Amazon Pinpoint replaces (overwrites) any
// existing values with the new values.
func pinpoint_UpdateEndpoint(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UpdateEndpointInput{
		// ApplicationId: *string, // Required
		// EndpointId: *string, // Required
		// EndpointRequest: *types.EndpointRequest, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointEndpointId) > 0 {
		input.EndpointId = aws.String(_pinpointEndpointId)
	}
	if len(_pinpointEndpointRequest) > 0 {
		if err := assignInputField(input, "EndpointRequest", _pinpointEndpointRequest); err != nil {
			log.Errorf("invalid --endpoint-request: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new batch of endpoints for an application or updates the settings and
// attributes of a batch of existing endpoints for an application. You can also use
// this operation to define custom attributes for a batch of endpoints. If an
// update includes one or more values for a custom attribute, Amazon Pinpoint
// replaces (overwrites) any existing values with the new values.
func pinpoint_UpdateEndpointsBatch(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UpdateEndpointsBatchInput{
		// ApplicationId: *string, // Required
		// EndpointBatchRequest: *types.EndpointBatchRequest, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointEndpointBatchRequest) > 0 {
		if err := assignInputField(input, "EndpointBatchRequest", _pinpointEndpointBatchRequest); err != nil {
			log.Errorf("invalid --endpoint-batch-request: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEndpointsBatch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the GCM channel for an application or updates the status and settings
// of the GCM channel for an application.
func pinpoint_UpdateGcmChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UpdateGcmChannelInput{
		// ApplicationId: *string, // Required
		// GCMChannelRequest: *types.GCMChannelRequest, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointGCMChannelRequest) > 0 {
		if err := assignInputField(input, "GCMChannelRequest", _pinpointGCMChannelRequest); err != nil {
			log.Errorf("invalid --gcm-channel-request: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateGcmChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing message template for messages sent through the in-app
// message channel.
func pinpoint_UpdateInAppTemplate(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UpdateInAppTemplateInput{
		// InAppTemplateRequest: *types.InAppTemplateRequest, // Required
		// TemplateName: *string, // Required
	}

	if len(_pinpointInAppTemplateRequest) > 0 {
		if err := assignInputField(input, "InAppTemplateRequest", _pinpointInAppTemplateRequest); err != nil {
			log.Errorf("invalid --in-app-template-request: %s", err.Error())
			return
		}
	}
	if len(_pinpointTemplateName) > 0 {
		input.TemplateName = aws.String(_pinpointTemplateName)
	}
	if len(_pinpointCreateNewVersion) > 0 {
		if err := assignInputField(input, "CreateNewVersion", _pinpointCreateNewVersion); err != nil {
			log.Errorf("invalid --create-new-version: %s", err.Error())
			return
		}
	}
	if len(_pinpointVersion) > 0 {
		input.Version = aws.String(_pinpointVersion)
	}

	if resp, err := client.UpdateInAppTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration and other settings for a journey.
func pinpoint_UpdateJourney(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UpdateJourneyInput{
		// ApplicationId: *string, // Required
		// JourneyId: *string, // Required
		// WriteJourneyRequest: *types.WriteJourneyRequest, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointJourneyId) > 0 {
		input.JourneyId = aws.String(_pinpointJourneyId)
	}
	if len(_pinpointWriteJourneyRequest) > 0 {
		if err := assignInputField(input, "WriteJourneyRequest", _pinpointWriteJourneyRequest); err != nil {
			log.Errorf("invalid --write-journey-request: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateJourney(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels (stops) an active journey.
func pinpoint_UpdateJourneyState(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UpdateJourneyStateInput{
		// ApplicationId: *string, // Required
		// JourneyId: *string, // Required
		// JourneyStateRequest: *types.JourneyStateRequest, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointJourneyId) > 0 {
		input.JourneyId = aws.String(_pinpointJourneyId)
	}
	if len(_pinpointJourneyStateRequest) > 0 {
		if err := assignInputField(input, "JourneyStateRequest", _pinpointJourneyStateRequest); err != nil {
			log.Errorf("invalid --journey-state-request: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateJourneyState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing message template for messages that are sent through a push
// notification channel.
func pinpoint_UpdatePushTemplate(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UpdatePushTemplateInput{
		// PushNotificationTemplateRequest: *types.PushNotificationTemplateRequest, // Required
		// TemplateName: *string, // Required
	}

	if len(_pinpointPushNotificationTemplateRequest) > 0 {
		if err := assignInputField(input, "PushNotificationTemplateRequest", _pinpointPushNotificationTemplateRequest); err != nil {
			log.Errorf("invalid --push-notification-template-request: %s", err.Error())
			return
		}
	}
	if len(_pinpointTemplateName) > 0 {
		input.TemplateName = aws.String(_pinpointTemplateName)
	}
	if len(_pinpointCreateNewVersion) > 0 {
		if err := assignInputField(input, "CreateNewVersion", _pinpointCreateNewVersion); err != nil {
			log.Errorf("invalid --create-new-version: %s", err.Error())
			return
		}
	}
	if len(_pinpointVersion) > 0 {
		input.Version = aws.String(_pinpointVersion)
	}

	if resp, err := client.UpdatePushTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Amazon Pinpoint configuration for a recommender model.
func pinpoint_UpdateRecommenderConfiguration(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UpdateRecommenderConfigurationInput{
		// RecommenderId: *string, // Required
		// UpdateRecommenderConfiguration: *types.UpdateRecommenderConfigurationShape, // Required
	}

	if len(_pinpointRecommenderId) > 0 {
		input.RecommenderId = aws.String(_pinpointRecommenderId)
	}

	if resp, err := client.UpdateRecommenderConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new segment for an application or updates the configuration,
// dimension, and other settings for an existing segment that's associated with an
// application.
func pinpoint_UpdateSegment(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UpdateSegmentInput{
		// ApplicationId: *string, // Required
		// SegmentId: *string, // Required
		// WriteSegmentRequest: *types.WriteSegmentRequest, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointSegmentId) > 0 {
		input.SegmentId = aws.String(_pinpointSegmentId)
	}
	if len(_pinpointWriteSegmentRequest) > 0 {
		if err := assignInputField(input, "WriteSegmentRequest", _pinpointWriteSegmentRequest); err != nil {
			log.Errorf("invalid --write-segment-request: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSegment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the SMS channel for an application or updates the status and settings
// of the SMS channel for an application.
func pinpoint_UpdateSmsChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UpdateSmsChannelInput{
		// ApplicationId: *string, // Required
		// SMSChannelRequest: *types.SMSChannelRequest, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointSMSChannelRequest) > 0 {
		if err := assignInputField(input, "SMSChannelRequest", _pinpointSMSChannelRequest); err != nil {
			log.Errorf("invalid --sms-channel-request: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSmsChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing message template for messages that are sent through the SMS
// channel.
func pinpoint_UpdateSmsTemplate(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UpdateSmsTemplateInput{
		// SMSTemplateRequest: *types.SMSTemplateRequest, // Required
		// TemplateName: *string, // Required
	}

	if len(_pinpointSMSTemplateRequest) > 0 {
		if err := assignInputField(input, "SMSTemplateRequest", _pinpointSMSTemplateRequest); err != nil {
			log.Errorf("invalid --sms-template-request: %s", err.Error())
			return
		}
	}
	if len(_pinpointTemplateName) > 0 {
		input.TemplateName = aws.String(_pinpointTemplateName)
	}
	if len(_pinpointCreateNewVersion) > 0 {
		if err := assignInputField(input, "CreateNewVersion", _pinpointCreateNewVersion); err != nil {
			log.Errorf("invalid --create-new-version: %s", err.Error())
			return
		}
	}
	if len(_pinpointVersion) > 0 {
		input.Version = aws.String(_pinpointVersion)
	}

	if resp, err := client.UpdateSmsTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the status of a specific version of a message template to active.
func pinpoint_UpdateTemplateActiveVersion(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UpdateTemplateActiveVersionInput{
		// TemplateActiveVersionRequest: *types.TemplateActiveVersionRequest, // Required
		// TemplateName: *string, // Required
		// TemplateType: *string, // Required
	}

	if len(_pinpointTemplateActiveVersionRequest) > 0 {
		if err := assignInputField(input, "TemplateActiveVersionRequest", _pinpointTemplateActiveVersionRequest); err != nil {
			log.Errorf("invalid --template-active-version-request: %s", err.Error())
			return
		}
	}
	if len(_pinpointTemplateName) > 0 {
		input.TemplateName = aws.String(_pinpointTemplateName)
	}
	if len(_pinpointTemplateType) > 0 {
		input.TemplateType = aws.String(_pinpointTemplateType)
	}

	if resp, err := client.UpdateTemplateActiveVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the voice channel for an application or updates the status and settings
// of the voice channel for an application.
func pinpoint_UpdateVoiceChannel(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UpdateVoiceChannelInput{
		// ApplicationId: *string, // Required
		// VoiceChannelRequest: *types.VoiceChannelRequest, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointVoiceChannelRequest) > 0 {
		if err := assignInputField(input, "VoiceChannelRequest", _pinpointVoiceChannelRequest); err != nil {
			log.Errorf("invalid --voice-channel-request: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateVoiceChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing message template for messages that are sent through the
// voice channel.
func pinpoint_UpdateVoiceTemplate(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.UpdateVoiceTemplateInput{
		// TemplateName: *string, // Required
		// VoiceTemplateRequest: *types.VoiceTemplateRequest, // Required
	}

	if len(_pinpointTemplateName) > 0 {
		input.TemplateName = aws.String(_pinpointTemplateName)
	}
	if len(_pinpointVoiceTemplateRequest) > 0 {
		if err := assignInputField(input, "VoiceTemplateRequest", _pinpointVoiceTemplateRequest); err != nil {
			log.Errorf("invalid --voice-template-request: %s", err.Error())
			return
		}
	}
	if len(_pinpointCreateNewVersion) > 0 {
		if err := assignInputField(input, "CreateNewVersion", _pinpointCreateNewVersion); err != nil {
			log.Errorf("invalid --create-new-version: %s", err.Error())
			return
		}
	}
	if len(_pinpointVersion) > 0 {
		input.Version = aws.String(_pinpointVersion)
	}

	if resp, err := client.UpdateVoiceTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Verify an OTP
func pinpoint_VerifyOTPMessage(cfg aws.Config, client *pinpoint.Client) {
	input := &pinpoint.VerifyOTPMessageInput{
		// ApplicationId: *string, // Required
		// VerifyOTPMessageRequestParameters: *types.VerifyOTPMessageRequestParameters, // Required
	}

	if len(_pinpointApplicationId) > 0 {
		input.ApplicationId = aws.String(_pinpointApplicationId)
	}
	if len(_pinpointVerifyOTPMessageRequestParameters) > 0 {
		if err := assignInputField(input, "VerifyOTPMessageRequestParameters", _pinpointVerifyOTPMessageRequestParameters); err != nil {
			log.Errorf("invalid --verify-otp-message-request-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.VerifyOTPMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_pinpointCmd)
	_pinpointCmd.Flags().SortFlags = false

	_pinpointCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_pinpointCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_pinpointCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_pinpointCmd.Flags().StringVarP(&_pinpointADMChannelRequest, "adm-channel-request", "", "", "Adm Channel Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointAPNSChannelRequest, "apns-channel-request", "", "", "Apns Channel Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointAPNSSandboxChannelRequest, "apns-sandbox-channel-request", "", "", "Apns Sandbox Channel Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointAPNSVoipChannelRequest, "apns-voip-channel-request", "", "", "Apns Voip Channel Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointAPNSVoipSandboxChannelRequest, "apns-voip-sandbox-channel-request", "", "", "Apns Voip Sandbox Channel Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointApplicationId, "application-id", "", "", "Application ID")
	_pinpointCmd.Flags().StringVarP(&_pinpointAttributeType, "attribute-type", "", "", "Attribute Type")
	_pinpointCmd.Flags().StringVarP(&_pinpointBaiduChannelRequest, "baidu-channel-request", "", "", "Baidu Channel Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointCampaignId, "campaign-id", "", "", "Campaign ID")
	_pinpointCmd.Flags().StringVarP(&_pinpointCreateApplicationRequest, "create-application-request", "", "", "Create Application Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointCreateNewVersion, "create-new-version", "", "", "Create New Version")
	_pinpointCmd.Flags().StringVarP(&_pinpointEmailChannelRequest, "email-channel-request", "", "", "Email Channel Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointEmailTemplateRequest, "email-template-request", "", "", "Email Template Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointEndTime, "end-time", "", "", "End Time")
	_pinpointCmd.Flags().StringVarP(&_pinpointEndpointBatchRequest, "endpoint-batch-request", "", "", "Endpoint Batch Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointEndpointId, "endpoint-id", "", "", "Endpoint ID")
	_pinpointCmd.Flags().StringVarP(&_pinpointEndpointRequest, "endpoint-request", "", "", "Endpoint Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointEventsRequest, "events-request", "", "", "Events Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointExportJobRequest, "export-job-request", "", "", "Export Job Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointGCMChannelRequest, "gcm-channel-request", "", "", "Gcm Channel Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointImportJobRequest, "import-job-request", "", "", "Import Job Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointInAppTemplateRequest, "in-app-template-request", "", "", "In App Template Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointJobId, "job-id", "", "", "Job ID")
	_pinpointCmd.Flags().StringVarP(&_pinpointJourneyActivityId, "journey-activity-id", "", "", "Journey Activity ID")
	_pinpointCmd.Flags().StringVarP(&_pinpointJourneyId, "journey-id", "", "", "Journey ID")
	_pinpointCmd.Flags().StringVarP(&_pinpointJourneyStateRequest, "journey-state-request", "", "", "Journey State Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointKpiName, "kpi-name", "", "", "Kpi Name")
	_pinpointCmd.Flags().StringVarP(&_pinpointMessageRequest, "message-request", "", "", "Message Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointNextToken, "next-token", "", "", "Next Token")
	_pinpointCmd.Flags().StringVarP(&_pinpointNumberValidateRequest, "number-validate-request", "", "", "Number Validate Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointPageSize, "page-size", "", "", "Page Size")
	_pinpointCmd.Flags().StringVarP(&_pinpointPrefix, "prefix", "", "", "Prefix")
	_pinpointCmd.Flags().StringVarP(&_pinpointPushNotificationTemplateRequest, "push-notification-template-request", "", "", "Push Notification Template Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointRecommenderId, "recommender-id", "", "", "Recommender ID")
	_pinpointCmd.Flags().StringVarP(&_pinpointResourceArn, "resource-arn", "", "", "Resource ARN")
	_pinpointCmd.Flags().StringVarP(&_pinpointRunId, "run-id", "", "", "Run ID")
	_pinpointCmd.Flags().StringVarP(&_pinpointSegmentId, "segment-id", "", "", "Segment ID")
	_pinpointCmd.Flags().StringVarP(&_pinpointSendOTPMessageRequestParameters, "send-otp-message-request-parameters", "", "", "Send Otp Message Request Parameters")
	_pinpointCmd.Flags().StringVarP(&_pinpointSendUsersMessageRequest, "send-users-message-request", "", "", "Send Users Message Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointSMSChannelRequest, "sms-channel-request", "", "", "Sms Channel Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointSMSTemplateRequest, "sms-template-request", "", "", "Sms Template Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointStartTime, "start-time", "", "", "Start Time")
	_pinpointCmd.Flags().StringSliceVarP(&_pinpointTagKeys, "tag-keys", "", nil, "Tag Keys")
	_pinpointCmd.Flags().StringVarP(&_pinpointTagsModel, "tags-model", "", "", "Tags Model")
	_pinpointCmd.Flags().StringVarP(&_pinpointTemplateActiveVersionRequest, "template-active-version-request", "", "", "Template Active Version Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointTemplateName, "template-name", "", "", "Template Name")
	_pinpointCmd.Flags().StringVarP(&_pinpointTemplateType, "template-type", "", "", "Template Type")
	_pinpointCmd.Flags().StringVarP(&_pinpointToken, "token", "", "", "Token")
	_pinpointCmd.Flags().StringVarP(&_pinpointUpdateAttributesRequest, "update-attributes-request", "", "", "Update Attributes Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointUserId, "user-id", "", "", "User ID")
	_pinpointCmd.Flags().StringVarP(&_pinpointVerifyOTPMessageRequestParameters, "verify-otp-message-request-parameters", "", "", "Verify Otp Message Request Parameters")
	_pinpointCmd.Flags().StringVarP(&_pinpointVersion, "version", "", "", "Version")
	_pinpointCmd.Flags().StringVarP(&_pinpointVoiceChannelRequest, "voice-channel-request", "", "", "Voice Channel Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointVoiceTemplateRequest, "voice-template-request", "", "", "Voice Template Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointWriteApplicationSettingsRequest, "write-application-settings-request", "", "", "Write Application Settings Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointWriteCampaignRequest, "write-campaign-request", "", "", "Write Campaign Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointWriteEventStream, "write-event-stream", "", "", "Write Event Stream")
	_pinpointCmd.Flags().StringVarP(&_pinpointWriteJourneyRequest, "write-journey-request", "", "", "Write Journey Request")
	_pinpointCmd.Flags().StringVarP(&_pinpointWriteSegmentRequest, "write-segment-request", "", "", "Write Segment Request")

	_pinpointCmd.Flags().BoolVarP(&_pinpointCreateApp, "create-app", "", false, "Create App")
	_pinpointCmd.Flags().BoolVarP(&_pinpointCreateCampaign, "create-campaign", "", false, "Create Campaign")
	_pinpointCmd.Flags().BoolVarP(&_pinpointCreateEmailTemplate, "create-email-template", "", false, "Create Email Template")
	_pinpointCmd.Flags().BoolVarP(&_pinpointCreateExportJob, "create-export-job", "", false, "Create Export Job")
	_pinpointCmd.Flags().BoolVarP(&_pinpointCreateImportJob, "create-import-job", "", false, "Create Import Job")
	_pinpointCmd.Flags().BoolVarP(&_pinpointCreateInAppTemplate, "create-in-app-template", "", false, "Create In App Template")
	_pinpointCmd.Flags().BoolVarP(&_pinpointCreateJourney, "create-journey", "", false, "Create Journey")
	_pinpointCmd.Flags().BoolVarP(&_pinpointCreatePushTemplate, "create-push-template", "", false, "Create Push Template")
	_pinpointCmd.Flags().BoolVarP(&_pinpointCreateRecommenderConfiguration, "create-recommender-configuration", "", false, "Create Recommender Configuration")
	_pinpointCmd.Flags().BoolVarP(&_pinpointCreateSegment, "create-segment", "", false, "Create Segment")
	_pinpointCmd.Flags().BoolVarP(&_pinpointCreateSmsTemplate, "create-sms-template", "", false, "Create Sms Template")
	_pinpointCmd.Flags().BoolVarP(&_pinpointCreateVoiceTemplate, "create-voice-template", "", false, "Create Voice Template")
	_pinpointCmd.Flags().BoolVarP(&_pinpointDeleteAdmChannel, "delete-adm-channel", "", false, "Delete Adm Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointDeleteApnsChannel, "delete-apns-channel", "", false, "Delete Apns Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointDeleteApnsSandboxChannel, "delete-apns-sandbox-channel", "", false, "Delete Apns Sandbox Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointDeleteApnsVoipChannel, "delete-apns-voip-channel", "", false, "Delete Apns Voip Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointDeleteApnsVoipSandboxChannel, "delete-apns-voip-sandbox-channel", "", false, "Delete Apns Voip Sandbox Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointDeleteApp, "delete-app", "", false, "Delete App")
	_pinpointCmd.Flags().BoolVarP(&_pinpointDeleteBaiduChannel, "delete-baidu-channel", "", false, "Delete Baidu Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointDeleteCampaign, "delete-campaign", "", false, "Delete Campaign")
	_pinpointCmd.Flags().BoolVarP(&_pinpointDeleteEmailChannel, "delete-email-channel", "", false, "Delete Email Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointDeleteEmailTemplate, "delete-email-template", "", false, "Delete Email Template")
	_pinpointCmd.Flags().BoolVarP(&_pinpointDeleteEndpoint, "delete-endpoint", "", false, "Delete Endpoint")
	_pinpointCmd.Flags().BoolVarP(&_pinpointDeleteEventStream, "delete-event-stream", "", false, "Delete Event Stream")
	_pinpointCmd.Flags().BoolVarP(&_pinpointDeleteGcmChannel, "delete-gcm-channel", "", false, "Delete Gcm Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointDeleteInAppTemplate, "delete-in-app-template", "", false, "Delete In App Template")
	_pinpointCmd.Flags().BoolVarP(&_pinpointDeleteJourney, "delete-journey", "", false, "Delete Journey")
	_pinpointCmd.Flags().BoolVarP(&_pinpointDeletePushTemplate, "delete-push-template", "", false, "Delete Push Template")
	_pinpointCmd.Flags().BoolVarP(&_pinpointDeleteRecommenderConfiguration, "delete-recommender-configuration", "", false, "Delete Recommender Configuration")
	_pinpointCmd.Flags().BoolVarP(&_pinpointDeleteSegment, "delete-segment", "", false, "Delete Segment")
	_pinpointCmd.Flags().BoolVarP(&_pinpointDeleteSmsChannel, "delete-sms-channel", "", false, "Delete Sms Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointDeleteSmsTemplate, "delete-sms-template", "", false, "Delete Sms Template")
	_pinpointCmd.Flags().BoolVarP(&_pinpointDeleteUserEndpoints, "delete-user-endpoints", "", false, "Delete User Endpoints")
	_pinpointCmd.Flags().BoolVarP(&_pinpointDeleteVoiceChannel, "delete-voice-channel", "", false, "Delete Voice Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointDeleteVoiceTemplate, "delete-voice-template", "", false, "Delete Voice Template")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetAdmChannel, "get-adm-channel", "", false, "Get Adm Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetApnsChannel, "get-apns-channel", "", false, "Get Apns Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetApnsSandboxChannel, "get-apns-sandbox-channel", "", false, "Get Apns Sandbox Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetApnsVoipChannel, "get-apns-voip-channel", "", false, "Get Apns Voip Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetApnsVoipSandboxChannel, "get-apns-voip-sandbox-channel", "", false, "Get Apns Voip Sandbox Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetApp, "get-app", "", false, "Get App")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetApplicationDateRangeKpi, "get-application-date-range-kpi", "", false, "Get Application Date Range Kpi")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetApplicationSettings, "get-application-settings", "", false, "Get Application Settings")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetApps, "get-apps", "", false, "Get Apps")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetBaiduChannel, "get-baidu-channel", "", false, "Get Baidu Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetCampaign, "get-campaign", "", false, "Get Campaign")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetCampaignActivities, "get-campaign-activities", "", false, "Get Campaign Activities")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetCampaignDateRangeKpi, "get-campaign-date-range-kpi", "", false, "Get Campaign Date Range Kpi")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetCampaignVersion, "get-campaign-version", "", false, "Get Campaign Version")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetCampaignVersions, "get-campaign-versions", "", false, "Get Campaign Versions")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetCampaigns, "get-campaigns", "", false, "Get Campaigns")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetChannels, "get-channels", "", false, "Get Channels")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetEmailChannel, "get-email-channel", "", false, "Get Email Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetEmailTemplate, "get-email-template", "", false, "Get Email Template")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetEndpoint, "get-endpoint", "", false, "Get Endpoint")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetEventStream, "get-event-stream", "", false, "Get Event Stream")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetExportJob, "get-export-job", "", false, "Get Export Job")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetExportJobs, "get-export-jobs", "", false, "Get Export Jobs")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetGcmChannel, "get-gcm-channel", "", false, "Get Gcm Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetImportJob, "get-import-job", "", false, "Get Import Job")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetImportJobs, "get-import-jobs", "", false, "Get Import Jobs")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetInAppMessages, "get-in-app-messages", "", false, "Get In App Messages")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetInAppTemplate, "get-in-app-template", "", false, "Get In App Template")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetJourney, "get-journey", "", false, "Get Journey")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetJourneyDateRangeKpi, "get-journey-date-range-kpi", "", false, "Get Journey Date Range Kpi")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetJourneyExecutionActivityMetrics, "get-journey-execution-activity-metrics", "", false, "Get Journey Execution Activity Metrics")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetJourneyExecutionMetrics, "get-journey-execution-metrics", "", false, "Get Journey Execution Metrics")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetJourneyRunExecutionActivityMetrics, "get-journey-run-execution-activity-metrics", "", false, "Get Journey Run Execution Activity Metrics")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetJourneyRunExecutionMetrics, "get-journey-run-execution-metrics", "", false, "Get Journey Run Execution Metrics")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetJourneyRuns, "get-journey-runs", "", false, "Get Journey Runs")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetPushTemplate, "get-push-template", "", false, "Get Push Template")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetRecommenderConfiguration, "get-recommender-configuration", "", false, "Get Recommender Configuration")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetRecommenderConfigurations, "get-recommender-configurations", "", false, "Get Recommender Configurations")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetSegment, "get-segment", "", false, "Get Segment")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetSegmentExportJobs, "get-segment-export-jobs", "", false, "Get Segment Export Jobs")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetSegmentImportJobs, "get-segment-import-jobs", "", false, "Get Segment Import Jobs")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetSegmentVersion, "get-segment-version", "", false, "Get Segment Version")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetSegmentVersions, "get-segment-versions", "", false, "Get Segment Versions")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetSegments, "get-segments", "", false, "Get Segments")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetSmsChannel, "get-sms-channel", "", false, "Get Sms Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetSmsTemplate, "get-sms-template", "", false, "Get Sms Template")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetUserEndpoints, "get-user-endpoints", "", false, "Get User Endpoints")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetVoiceChannel, "get-voice-channel", "", false, "Get Voice Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointGetVoiceTemplate, "get-voice-template", "", false, "Get Voice Template")
	_pinpointCmd.Flags().BoolVarP(&_pinpointListJourneys, "list-journeys", "", false, "List Journeys")
	_pinpointCmd.Flags().BoolVarP(&_pinpointListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_pinpointCmd.Flags().BoolVarP(&_pinpointListTemplateVersions, "list-template-versions", "", false, "List Template Versions")
	_pinpointCmd.Flags().BoolVarP(&_pinpointListTemplates, "list-templates", "", false, "List Templates")
	_pinpointCmd.Flags().BoolVarP(&_pinpointPhoneNumberValidate, "phone-number-validate", "", false, "Phone Number Validate")
	_pinpointCmd.Flags().BoolVarP(&_pinpointPutEventStream, "put-event-stream", "", false, "Put Event Stream")
	_pinpointCmd.Flags().BoolVarP(&_pinpointPutEvents, "put-events", "", false, "Put Events")
	_pinpointCmd.Flags().BoolVarP(&_pinpointRemoveAttributes, "remove-attributes", "", false, "Remove Attributes")
	_pinpointCmd.Flags().BoolVarP(&_pinpointSendMessages, "send-messages", "", false, "Send Messages")
	_pinpointCmd.Flags().BoolVarP(&_pinpointSendOTPMessage, "send-otp-message", "", false, "Send Otp Message")
	_pinpointCmd.Flags().BoolVarP(&_pinpointSendUsersMessages, "send-users-messages", "", false, "Send Users Messages")
	_pinpointCmd.Flags().BoolVarP(&_pinpointTagResource, "tag-resource", "", false, "Tag Resource")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUntagResource, "untag-resource", "", false, "Untag Resource")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUpdateAdmChannel, "update-adm-channel", "", false, "Update Adm Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUpdateApnsChannel, "update-apns-channel", "", false, "Update Apns Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUpdateApnsSandboxChannel, "update-apns-sandbox-channel", "", false, "Update Apns Sandbox Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUpdateApnsVoipChannel, "update-apns-voip-channel", "", false, "Update Apns Voip Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUpdateApnsVoipSandboxChannel, "update-apns-voip-sandbox-channel", "", false, "Update Apns Voip Sandbox Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUpdateApplicationSettings, "update-application-settings", "", false, "Update Application Settings")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUpdateBaiduChannel, "update-baidu-channel", "", false, "Update Baidu Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUpdateCampaign, "update-campaign", "", false, "Update Campaign")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUpdateEmailChannel, "update-email-channel", "", false, "Update Email Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUpdateEmailTemplate, "update-email-template", "", false, "Update Email Template")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUpdateEndpoint, "update-endpoint", "", false, "Update Endpoint")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUpdateEndpointsBatch, "update-endpoints-batch", "", false, "Update Endpoints Batch")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUpdateGcmChannel, "update-gcm-channel", "", false, "Update Gcm Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUpdateInAppTemplate, "update-in-app-template", "", false, "Update In App Template")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUpdateJourney, "update-journey", "", false, "Update Journey")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUpdateJourneyState, "update-journey-state", "", false, "Update Journey State")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUpdatePushTemplate, "update-push-template", "", false, "Update Push Template")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUpdateRecommenderConfiguration, "update-recommender-configuration", "", false, "Update Recommender Configuration")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUpdateSegment, "update-segment", "", false, "Update Segment")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUpdateSmsChannel, "update-sms-channel", "", false, "Update Sms Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUpdateSmsTemplate, "update-sms-template", "", false, "Update Sms Template")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUpdateTemplateActiveVersion, "update-template-active-version", "", false, "Update Template Active Version")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUpdateVoiceChannel, "update-voice-channel", "", false, "Update Voice Channel")
	_pinpointCmd.Flags().BoolVarP(&_pinpointUpdateVoiceTemplate, "update-voice-template", "", false, "Update Voice Template")
	_pinpointCmd.Flags().BoolVarP(&_pinpointVerifyOTPMessage, "verify-otp-message", "", false, "Verify Otp Message")

}
