package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/qconnect"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// qconnectCmd represents the qconnect command
var _qconnectCmd = &cobra.Command{
	Use:   "qconnect",
	Short: "AWS qconnect CLI",
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
		client := qconnect.NewFromConfig(cfg)
		if _qconnectActivateMessageTemplate {
			qconnect_ActivateMessageTemplate(cfg, client)
			return
		}
		if _qconnectCreateAIAgent {
			qconnect_CreateAIAgent(cfg, client)
			return
		}
		if _qconnectCreateAIAgentVersion {
			qconnect_CreateAIAgentVersion(cfg, client)
			return
		}
		if _qconnectCreateAIGuardrail {
			qconnect_CreateAIGuardrail(cfg, client)
			return
		}
		if _qconnectCreateAIGuardrailVersion {
			qconnect_CreateAIGuardrailVersion(cfg, client)
			return
		}
		if _qconnectCreateAIPrompt {
			qconnect_CreateAIPrompt(cfg, client)
			return
		}
		if _qconnectCreateAIPromptVersion {
			qconnect_CreateAIPromptVersion(cfg, client)
			return
		}
		if _qconnectCreateAssistant {
			qconnect_CreateAssistant(cfg, client)
			return
		}
		if _qconnectCreateAssistantAssociation {
			qconnect_CreateAssistantAssociation(cfg, client)
			return
		}
		if _qconnectCreateContent {
			qconnect_CreateContent(cfg, client)
			return
		}
		if _qconnectCreateContentAssociation {
			qconnect_CreateContentAssociation(cfg, client)
			return
		}
		if _qconnectCreateKnowledgeBase {
			qconnect_CreateKnowledgeBase(cfg, client)
			return
		}
		if _qconnectCreateMessageTemplate {
			qconnect_CreateMessageTemplate(cfg, client)
			return
		}
		if _qconnectCreateMessageTemplateAttachment {
			qconnect_CreateMessageTemplateAttachment(cfg, client)
			return
		}
		if _qconnectCreateMessageTemplateVersion {
			qconnect_CreateMessageTemplateVersion(cfg, client)
			return
		}
		if _qconnectCreateQuickResponse {
			qconnect_CreateQuickResponse(cfg, client)
			return
		}
		if _qconnectCreateSession {
			qconnect_CreateSession(cfg, client)
			return
		}
		if _qconnectDeactivateMessageTemplate {
			qconnect_DeactivateMessageTemplate(cfg, client)
			return
		}
		if _qconnectDeleteAIAgent {
			qconnect_DeleteAIAgent(cfg, client)
			return
		}
		if _qconnectDeleteAIAgentVersion {
			qconnect_DeleteAIAgentVersion(cfg, client)
			return
		}
		if _qconnectDeleteAIGuardrail {
			qconnect_DeleteAIGuardrail(cfg, client)
			return
		}
		if _qconnectDeleteAIGuardrailVersion {
			qconnect_DeleteAIGuardrailVersion(cfg, client)
			return
		}
		if _qconnectDeleteAIPrompt {
			qconnect_DeleteAIPrompt(cfg, client)
			return
		}
		if _qconnectDeleteAIPromptVersion {
			qconnect_DeleteAIPromptVersion(cfg, client)
			return
		}
		if _qconnectDeleteAssistant {
			qconnect_DeleteAssistant(cfg, client)
			return
		}
		if _qconnectDeleteAssistantAssociation {
			qconnect_DeleteAssistantAssociation(cfg, client)
			return
		}
		if _qconnectDeleteContent {
			qconnect_DeleteContent(cfg, client)
			return
		}
		if _qconnectDeleteContentAssociation {
			qconnect_DeleteContentAssociation(cfg, client)
			return
		}
		if _qconnectDeleteImportJob {
			qconnect_DeleteImportJob(cfg, client)
			return
		}
		if _qconnectDeleteKnowledgeBase {
			qconnect_DeleteKnowledgeBase(cfg, client)
			return
		}
		if _qconnectDeleteMessageTemplate {
			qconnect_DeleteMessageTemplate(cfg, client)
			return
		}
		if _qconnectDeleteMessageTemplateAttachment {
			qconnect_DeleteMessageTemplateAttachment(cfg, client)
			return
		}
		if _qconnectDeleteQuickResponse {
			qconnect_DeleteQuickResponse(cfg, client)
			return
		}
		if _qconnectGetAIAgent {
			qconnect_GetAIAgent(cfg, client)
			return
		}
		if _qconnectGetAIGuardrail {
			qconnect_GetAIGuardrail(cfg, client)
			return
		}
		if _qconnectGetAIPrompt {
			qconnect_GetAIPrompt(cfg, client)
			return
		}
		if _qconnectGetAssistant {
			qconnect_GetAssistant(cfg, client)
			return
		}
		if _qconnectGetAssistantAssociation {
			qconnect_GetAssistantAssociation(cfg, client)
			return
		}
		if _qconnectGetContent {
			qconnect_GetContent(cfg, client)
			return
		}
		if _qconnectGetContentAssociation {
			qconnect_GetContentAssociation(cfg, client)
			return
		}
		if _qconnectGetContentSummary {
			qconnect_GetContentSummary(cfg, client)
			return
		}
		if _qconnectGetImportJob {
			qconnect_GetImportJob(cfg, client)
			return
		}
		if _qconnectGetKnowledgeBase {
			qconnect_GetKnowledgeBase(cfg, client)
			return
		}
		if _qconnectGetMessageTemplate {
			qconnect_GetMessageTemplate(cfg, client)
			return
		}
		if _qconnectGetNextMessage {
			qconnect_GetNextMessage(cfg, client)
			return
		}
		if _qconnectGetQuickResponse {
			qconnect_GetQuickResponse(cfg, client)
			return
		}
		if _qconnectGetRecommendations {
			qconnect_GetRecommendations(cfg, client)
			return
		}
		if _qconnectGetSession {
			qconnect_GetSession(cfg, client)
			return
		}
		if _qconnectListAIAgentVersions {
			qconnect_ListAIAgentVersions(cfg, client)
			return
		}
		if _qconnectListAIAgents {
			qconnect_ListAIAgents(cfg, client)
			return
		}
		if _qconnectListAIGuardrailVersions {
			qconnect_ListAIGuardrailVersions(cfg, client)
			return
		}
		if _qconnectListAIGuardrails {
			qconnect_ListAIGuardrails(cfg, client)
			return
		}
		if _qconnectListAIPromptVersions {
			qconnect_ListAIPromptVersions(cfg, client)
			return
		}
		if _qconnectListAIPrompts {
			qconnect_ListAIPrompts(cfg, client)
			return
		}
		if _qconnectListAssistantAssociations {
			qconnect_ListAssistantAssociations(cfg, client)
			return
		}
		if _qconnectListAssistants {
			qconnect_ListAssistants(cfg, client)
			return
		}
		if _qconnectListContentAssociations {
			qconnect_ListContentAssociations(cfg, client)
			return
		}
		if _qconnectListContents {
			qconnect_ListContents(cfg, client)
			return
		}
		if _qconnectListImportJobs {
			qconnect_ListImportJobs(cfg, client)
			return
		}
		if _qconnectListKnowledgeBases {
			qconnect_ListKnowledgeBases(cfg, client)
			return
		}
		if _qconnectListMessageTemplateVersions {
			qconnect_ListMessageTemplateVersions(cfg, client)
			return
		}
		if _qconnectListMessageTemplates {
			qconnect_ListMessageTemplates(cfg, client)
			return
		}
		if _qconnectListMessages {
			qconnect_ListMessages(cfg, client)
			return
		}
		if _qconnectListQuickResponses {
			qconnect_ListQuickResponses(cfg, client)
			return
		}
		if _qconnectListSpans {
			qconnect_ListSpans(cfg, client)
			return
		}
		if _qconnectListTagsForResource {
			qconnect_ListTagsForResource(cfg, client)
			return
		}
		if _qconnectNotifyRecommendationsReceived {
			qconnect_NotifyRecommendationsReceived(cfg, client)
			return
		}
		if _qconnectPutFeedback {
			qconnect_PutFeedback(cfg, client)
			return
		}
		if _qconnectQueryAssistant {
			qconnect_QueryAssistant(cfg, client)
			return
		}
		if _qconnectRemoveAssistantAIAgent {
			qconnect_RemoveAssistantAIAgent(cfg, client)
			return
		}
		if _qconnectRemoveKnowledgeBaseTemplateUri {
			qconnect_RemoveKnowledgeBaseTemplateUri(cfg, client)
			return
		}
		if _qconnectRenderMessageTemplate {
			qconnect_RenderMessageTemplate(cfg, client)
			return
		}
		if _qconnectRetrieve {
			qconnect_Retrieve(cfg, client)
			return
		}
		if _qconnectSearchContent {
			qconnect_SearchContent(cfg, client)
			return
		}
		if _qconnectSearchMessageTemplates {
			qconnect_SearchMessageTemplates(cfg, client)
			return
		}
		if _qconnectSearchQuickResponses {
			qconnect_SearchQuickResponses(cfg, client)
			return
		}
		if _qconnectSearchSessions {
			qconnect_SearchSessions(cfg, client)
			return
		}
		if _qconnectSendMessage {
			qconnect_SendMessage(cfg, client)
			return
		}
		if _qconnectStartContentUpload {
			qconnect_StartContentUpload(cfg, client)
			return
		}
		if _qconnectStartImportJob {
			qconnect_StartImportJob(cfg, client)
			return
		}
		if _qconnectTagResource {
			qconnect_TagResource(cfg, client)
			return
		}
		if _qconnectUntagResource {
			qconnect_UntagResource(cfg, client)
			return
		}
		if _qconnectUpdateAIAgent {
			qconnect_UpdateAIAgent(cfg, client)
			return
		}
		if _qconnectUpdateAIGuardrail {
			qconnect_UpdateAIGuardrail(cfg, client)
			return
		}
		if _qconnectUpdateAIPrompt {
			qconnect_UpdateAIPrompt(cfg, client)
			return
		}
		if _qconnectUpdateAssistantAIAgent {
			qconnect_UpdateAssistantAIAgent(cfg, client)
			return
		}
		if _qconnectUpdateContent {
			qconnect_UpdateContent(cfg, client)
			return
		}
		if _qconnectUpdateKnowledgeBaseTemplateUri {
			qconnect_UpdateKnowledgeBaseTemplateUri(cfg, client)
			return
		}
		if _qconnectUpdateMessageTemplate {
			qconnect_UpdateMessageTemplate(cfg, client)
			return
		}
		if _qconnectUpdateMessageTemplateMetadata {
			qconnect_UpdateMessageTemplateMetadata(cfg, client)
			return
		}
		if _qconnectUpdateQuickResponse {
			qconnect_UpdateQuickResponse(cfg, client)
			return
		}
		if _qconnectUpdateSession {
			qconnect_UpdateSession(cfg, client)
			return
		}
		if _qconnectUpdateSessionData {
			qconnect_UpdateSessionData(cfg, client)
			return
		}

	},
}

var (
	_qconnectActivateMessageTemplate         bool
	_qconnectCreateAIAgent                   bool
	_qconnectCreateAIAgentVersion            bool
	_qconnectCreateAIGuardrail               bool
	_qconnectCreateAIGuardrailVersion        bool
	_qconnectCreateAIPrompt                  bool
	_qconnectCreateAIPromptVersion           bool
	_qconnectCreateAssistant                 bool
	_qconnectCreateAssistantAssociation      bool
	_qconnectCreateContent                   bool
	_qconnectCreateContentAssociation        bool
	_qconnectCreateKnowledgeBase             bool
	_qconnectCreateMessageTemplate           bool
	_qconnectCreateMessageTemplateAttachment bool
	_qconnectCreateMessageTemplateVersion    bool
	_qconnectCreateQuickResponse             bool
	_qconnectCreateSession                   bool
	_qconnectDeactivateMessageTemplate       bool
	_qconnectDeleteAIAgent                   bool
	_qconnectDeleteAIAgentVersion            bool
	_qconnectDeleteAIGuardrail               bool
	_qconnectDeleteAIGuardrailVersion        bool
	_qconnectDeleteAIPrompt                  bool
	_qconnectDeleteAIPromptVersion           bool
	_qconnectDeleteAssistant                 bool
	_qconnectDeleteAssistantAssociation      bool
	_qconnectDeleteContent                   bool
	_qconnectDeleteContentAssociation        bool
	_qconnectDeleteImportJob                 bool
	_qconnectDeleteKnowledgeBase             bool
	_qconnectDeleteMessageTemplate           bool
	_qconnectDeleteMessageTemplateAttachment bool
	_qconnectDeleteQuickResponse             bool
	_qconnectGetAIAgent                      bool
	_qconnectGetAIGuardrail                  bool
	_qconnectGetAIPrompt                     bool
	_qconnectGetAssistant                    bool
	_qconnectGetAssistantAssociation         bool
	_qconnectGetContent                      bool
	_qconnectGetContentAssociation           bool
	_qconnectGetContentSummary               bool
	_qconnectGetImportJob                    bool
	_qconnectGetKnowledgeBase                bool
	_qconnectGetMessageTemplate              bool
	_qconnectGetNextMessage                  bool
	_qconnectGetQuickResponse                bool
	_qconnectGetRecommendations              bool
	_qconnectGetSession                      bool
	_qconnectListAIAgentVersions             bool
	_qconnectListAIAgents                    bool
	_qconnectListAIGuardrailVersions         bool
	_qconnectListAIGuardrails                bool
	_qconnectListAIPromptVersions            bool
	_qconnectListAIPrompts                   bool
	_qconnectListAssistantAssociations       bool
	_qconnectListAssistants                  bool
	_qconnectListContentAssociations         bool
	_qconnectListContents                    bool
	_qconnectListImportJobs                  bool
	_qconnectListKnowledgeBases              bool
	_qconnectListMessageTemplateVersions     bool
	_qconnectListMessageTemplates            bool
	_qconnectListMessages                    bool
	_qconnectListQuickResponses              bool
	_qconnectListSpans                       bool
	_qconnectListTagsForResource             bool
	_qconnectNotifyRecommendationsReceived   bool
	_qconnectPutFeedback                     bool
	_qconnectQueryAssistant                  bool
	_qconnectRemoveAssistantAIAgent          bool
	_qconnectRemoveKnowledgeBaseTemplateUri  bool
	_qconnectRenderMessageTemplate           bool
	_qconnectRetrieve                        bool
	_qconnectSearchContent                   bool
	_qconnectSearchMessageTemplates          bool
	_qconnectSearchQuickResponses            bool
	_qconnectSearchSessions                  bool
	_qconnectSendMessage                     bool
	_qconnectStartContentUpload              bool
	_qconnectStartImportJob                  bool
	_qconnectTagResource                     bool
	_qconnectUntagResource                   bool
	_qconnectUpdateAIAgent                   bool
	_qconnectUpdateAIGuardrail               bool
	_qconnectUpdateAIPrompt                  bool
	_qconnectUpdateAssistantAIAgent          bool
	_qconnectUpdateContent                   bool
	_qconnectUpdateKnowledgeBaseTemplateUri  bool
	_qconnectUpdateMessageTemplate           bool
	_qconnectUpdateMessageTemplateMetadata   bool
	_qconnectUpdateQuickResponse             bool
	_qconnectUpdateSession                   bool
	_qconnectUpdateSessionData               bool

	_qconnectAiAgentConfiguration                string
	_qconnectAiAgentId                           string
	_qconnectAiAgentType                         string
	_qconnectAiGuardrailId                       string
	_qconnectAiPromptId                          string
	_qconnectApiFormat                           string
	_qconnectAssistantAssociationId              string
	_qconnectAssistantId                         string
	_qconnectAssociation                         string
	_qconnectAssociationType                     string
	_qconnectAttachmentId                        string
	_qconnectAttributes                          string
	_qconnectBlockedInputMessaging               string
	_qconnectBlockedOutputsMessaging             string
	_qconnectBody                                string
	_qconnectChannelSubtype                      string
	_qconnectChannels                            []string
	_qconnectClientToken                         string
	_qconnectConfiguration                       string
	_qconnectContactArn                          string
	_qconnectContent                             string
	_qconnectContentAssociationId                string
	_qconnectContentDisposition                  string
	_qconnectContentFeedback                     string
	_qconnectContentId                           string
	_qconnectContentPolicyConfig                 string
	_qconnectContentType                         string
	_qconnectContextualGroundingPolicyConfig     string
	_qconnectConversationContext                 string
	_qconnectData                                string
	_qconnectDefaultAttributes                   string
	_qconnectDescription                         string
	_qconnectExternalSourceConfiguration         string
	_qconnectFilter                              string
	_qconnectGroupingConfiguration               string
	_qconnectImportJobId                         string
	_qconnectImportJobType                       string
	_qconnectInferenceConfiguration              string
	_qconnectIsActive                            string
	_qconnectKnowledgeBaseId                     string
	_qconnectKnowledgeBaseType                   string
	_qconnectLanguage                            string
	_qconnectMaxResults                          string
	_qconnectMessage                             string
	_qconnectMessageTemplateContentSha256        string
	_qconnectMessageTemplateId                   string
	_qconnectMetadata                            string
	_qconnectModelId                             string
	_qconnectModifiedTime                        string
	_qconnectName                                string
	_qconnectNamespace                           string
	_qconnectNextChunkToken                      string
	_qconnectNextMessageToken                    string
	_qconnectNextToken                           string
	_qconnectOrchestratorConfigurationList       string
	_qconnectOrchestratorUseCase                 string
	_qconnectOrigin                              string
	_qconnectOverrideKnowledgeBaseSearchType     string
	_qconnectOverrideLinkOutUri                  string
	_qconnectPresignedUrlTimeToLive              string
	_qconnectQueryCondition                      string
	_qconnectQueryInputData                      string
	_qconnectQueryText                           string
	_qconnectQuickResponseId                     string
	_qconnectRecommendationIds                   []string
	_qconnectRecommendationType                  string
	_qconnectRemoveDescription                   string
	_qconnectRemoveGroupingConfiguration         string
	_qconnectRemoveOrchestratorConfigurationList string
	_qconnectRemoveOverrideLinkOutUri            string
	_qconnectRemoveShortcutKey                   string
	_qconnectRenderingConfiguration              string
	_qconnectResourceArn                         string
	_qconnectRetrievalConfiguration              string
	_qconnectRetrievalQuery                      string
	_qconnectRevisionId                          string
	_qconnectSearchExpression                    string
	_qconnectSensitiveInformationPolicyConfig    string
	_qconnectServerSideEncryptionConfiguration   string
	_qconnectSessionId                           string
	_qconnectShortcutKey                         string
	_qconnectSourceConfiguration                 string
	_qconnectTagFilter                           string
	_qconnectTagKeys                             []string
	_qconnectTags                                string
	_qconnectTargetId                            string
	_qconnectTargetType                          string
	_qconnectTemplateConfiguration               string
	_qconnectTemplateType                        string
	_qconnectTemplateUri                         string
	_qconnectTitle                               string
	_qconnectTopicPolicyConfig                   string
	_qconnectType                                string
	_qconnectUploadId                            string
	_qconnectVectorIngestionConfiguration        string
	_qconnectVersionNumber                       string
	_qconnectVisibilityStatus                    string
	_qconnectWaitTimeSeconds                     string
	_qconnectWordPolicyConfig                    string
)

// Activates a specific version of the Amazon Q in Connect message template. After
// the version is activated, the previous active version will be deactivated
// automatically. You can use the $ACTIVE_VERSION qualifier later to reference the
// version that is in active status.
func qconnect_ActivateMessageTemplate(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.ActivateMessageTemplateInput{
		// KnowledgeBaseId: *string, // Required
		// MessageTemplateId: *string, // Required
		// VersionNumber: *int64, // Required
	}

	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectMessageTemplateId) > 0 {
		input.MessageTemplateId = aws.String(_qconnectMessageTemplateId)
	}
	if len(_qconnectVersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _qconnectVersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.ActivateMessageTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Q in Connect AI Agent.
func qconnect_CreateAIAgent(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.CreateAIAgentInput{
		// AssistantId: *string, // Required
		// Configuration: types.AIAgentConfiguration, // Required
		// Name: *string, // Required
		// Type: types.AIAgentType, // Required
		// VisibilityStatus: types.VisibilityStatus, // Required
	}

	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _qconnectConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_qconnectName) > 0 {
		input.Name = aws.String(_qconnectName)
	}
	if len(_qconnectType) > 0 {
		if err := assignInputField(input, "Type", _qconnectType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_qconnectVisibilityStatus) > 0 {
		if err := assignInputField(input, "VisibilityStatus", _qconnectVisibilityStatus); err != nil {
			log.Errorf("invalid --visibility-status: %s", err.Error())
			return
		}
	}
	if len(_qconnectClientToken) > 0 {
		input.ClientToken = aws.String(_qconnectClientToken)
	}
	if len(_qconnectDescription) > 0 {
		input.Description = aws.String(_qconnectDescription)
	}
	if len(_qconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _qconnectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAIAgent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates and Amazon Q in Connect AI Agent version.
func qconnect_CreateAIAgentVersion(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.CreateAIAgentVersionInput{
		// AiAgentId: *string, // Required
		// AssistantId: *string, // Required
	}

	if len(_qconnectAiAgentId) > 0 {
		input.AiAgentId = aws.String(_qconnectAiAgentId)
	}
	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectClientToken) > 0 {
		input.ClientToken = aws.String(_qconnectClientToken)
	}
	if len(_qconnectModifiedTime) > 0 {
		if err := assignInputField(input, "ModifiedTime", _qconnectModifiedTime); err != nil {
			log.Errorf("invalid --modified-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAIAgentVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Q in Connect AI Guardrail.
func qconnect_CreateAIGuardrail(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.CreateAIGuardrailInput{
		// AssistantId: *string, // Required
		// BlockedInputMessaging: *string, // Required
		// BlockedOutputsMessaging: *string, // Required
		// Name: *string, // Required
		// VisibilityStatus: types.VisibilityStatus, // Required
	}

	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectBlockedInputMessaging) > 0 {
		input.BlockedInputMessaging = aws.String(_qconnectBlockedInputMessaging)
	}
	if len(_qconnectBlockedOutputsMessaging) > 0 {
		input.BlockedOutputsMessaging = aws.String(_qconnectBlockedOutputsMessaging)
	}
	if len(_qconnectName) > 0 {
		input.Name = aws.String(_qconnectName)
	}
	if len(_qconnectVisibilityStatus) > 0 {
		if err := assignInputField(input, "VisibilityStatus", _qconnectVisibilityStatus); err != nil {
			log.Errorf("invalid --visibility-status: %s", err.Error())
			return
		}
	}
	if len(_qconnectClientToken) > 0 {
		input.ClientToken = aws.String(_qconnectClientToken)
	}
	if len(_qconnectContentPolicyConfig) > 0 {
		if err := assignInputField(input, "ContentPolicyConfig", _qconnectContentPolicyConfig); err != nil {
			log.Errorf("invalid --content-policy-config: %s", err.Error())
			return
		}
	}
	if len(_qconnectContextualGroundingPolicyConfig) > 0 {
		if err := assignInputField(input, "ContextualGroundingPolicyConfig", _qconnectContextualGroundingPolicyConfig); err != nil {
			log.Errorf("invalid --contextual-grounding-policy-config: %s", err.Error())
			return
		}
	}
	if len(_qconnectDescription) > 0 {
		input.Description = aws.String(_qconnectDescription)
	}
	if len(_qconnectSensitiveInformationPolicyConfig) > 0 {
		if err := assignInputField(input, "SensitiveInformationPolicyConfig", _qconnectSensitiveInformationPolicyConfig); err != nil {
			log.Errorf("invalid --sensitive-information-policy-config: %s", err.Error())
			return
		}
	}
	if len(_qconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _qconnectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_qconnectTopicPolicyConfig) > 0 {
		if err := assignInputField(input, "TopicPolicyConfig", _qconnectTopicPolicyConfig); err != nil {
			log.Errorf("invalid --topic-policy-config: %s", err.Error())
			return
		}
	}
	if len(_qconnectWordPolicyConfig) > 0 {
		if err := assignInputField(input, "WordPolicyConfig", _qconnectWordPolicyConfig); err != nil {
			log.Errorf("invalid --word-policy-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAIGuardrail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Q in Connect AI Guardrail version.
func qconnect_CreateAIGuardrailVersion(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.CreateAIGuardrailVersionInput{
		// AiGuardrailId: *string, // Required
		// AssistantId: *string, // Required
	}

	if len(_qconnectAiGuardrailId) > 0 {
		input.AiGuardrailId = aws.String(_qconnectAiGuardrailId)
	}
	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectClientToken) > 0 {
		input.ClientToken = aws.String(_qconnectClientToken)
	}
	if len(_qconnectModifiedTime) > 0 {
		if err := assignInputField(input, "ModifiedTime", _qconnectModifiedTime); err != nil {
			log.Errorf("invalid --modified-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAIGuardrailVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Q in Connect AI Prompt.
func qconnect_CreateAIPrompt(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.CreateAIPromptInput{
		// ApiFormat: types.AIPromptAPIFormat, // Required
		// AssistantId: *string, // Required
		// ModelId: *string, // Required
		// Name: *string, // Required
		// TemplateConfiguration: types.AIPromptTemplateConfiguration, // Required
		// TemplateType: types.AIPromptTemplateType, // Required
		// Type: types.AIPromptType, // Required
		// VisibilityStatus: types.VisibilityStatus, // Required
	}

	if len(_qconnectApiFormat) > 0 {
		if err := assignInputField(input, "ApiFormat", _qconnectApiFormat); err != nil {
			log.Errorf("invalid --api-format: %s", err.Error())
			return
		}
	}
	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectModelId) > 0 {
		input.ModelId = aws.String(_qconnectModelId)
	}
	if len(_qconnectName) > 0 {
		input.Name = aws.String(_qconnectName)
	}
	if len(_qconnectTemplateConfiguration) > 0 {
		if err := assignInputField(input, "TemplateConfiguration", _qconnectTemplateConfiguration); err != nil {
			log.Errorf("invalid --template-configuration: %s", err.Error())
			return
		}
	}
	if len(_qconnectTemplateType) > 0 {
		if err := assignInputField(input, "TemplateType", _qconnectTemplateType); err != nil {
			log.Errorf("invalid --template-type: %s", err.Error())
			return
		}
	}
	if len(_qconnectType) > 0 {
		if err := assignInputField(input, "Type", _qconnectType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_qconnectVisibilityStatus) > 0 {
		if err := assignInputField(input, "VisibilityStatus", _qconnectVisibilityStatus); err != nil {
			log.Errorf("invalid --visibility-status: %s", err.Error())
			return
		}
	}
	if len(_qconnectClientToken) > 0 {
		input.ClientToken = aws.String(_qconnectClientToken)
	}
	if len(_qconnectDescription) > 0 {
		input.Description = aws.String(_qconnectDescription)
	}
	if len(_qconnectInferenceConfiguration) > 0 {
		if err := assignInputField(input, "InferenceConfiguration", _qconnectInferenceConfiguration); err != nil {
			log.Errorf("invalid --inference-configuration: %s", err.Error())
			return
		}
	}
	if len(_qconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _qconnectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAIPrompt(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Q in Connect AI Prompt version.
func qconnect_CreateAIPromptVersion(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.CreateAIPromptVersionInput{
		// AiPromptId: *string, // Required
		// AssistantId: *string, // Required
	}

	if len(_qconnectAiPromptId) > 0 {
		input.AiPromptId = aws.String(_qconnectAiPromptId)
	}
	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectClientToken) > 0 {
		input.ClientToken = aws.String(_qconnectClientToken)
	}
	if len(_qconnectModifiedTime) > 0 {
		if err := assignInputField(input, "ModifiedTime", _qconnectModifiedTime); err != nil {
			log.Errorf("invalid --modified-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAIPromptVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Q in Connect assistant.
func qconnect_CreateAssistant(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.CreateAssistantInput{
		// Name: *string, // Required
		// Type: types.AssistantType, // Required
	}

	if len(_qconnectName) > 0 {
		input.Name = aws.String(_qconnectName)
	}
	if len(_qconnectType) > 0 {
		if err := assignInputField(input, "Type", _qconnectType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_qconnectClientToken) > 0 {
		input.ClientToken = aws.String(_qconnectClientToken)
	}
	if len(_qconnectDescription) > 0 {
		input.Description = aws.String(_qconnectDescription)
	}
	if len(_qconnectServerSideEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "ServerSideEncryptionConfiguration", _qconnectServerSideEncryptionConfiguration); err != nil {
			log.Errorf("invalid --server-side-encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_qconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _qconnectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAssistant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an association between an Amazon Q in Connect assistant and another
// resource. Currently, the only supported association is with a knowledge base. An
// assistant can have only a single association.
func qconnect_CreateAssistantAssociation(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.CreateAssistantAssociationInput{
		// AssistantId: *string, // Required
		// Association: types.AssistantAssociationInputData, // Required
		// AssociationType: types.AssociationType, // Required
	}

	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectAssociation) > 0 {
		if err := assignInputField(input, "Association", _qconnectAssociation); err != nil {
			log.Errorf("invalid --association: %s", err.Error())
			return
		}
	}
	if len(_qconnectAssociationType) > 0 {
		if err := assignInputField(input, "AssociationType", _qconnectAssociationType); err != nil {
			log.Errorf("invalid --association-type: %s", err.Error())
			return
		}
	}
	if len(_qconnectClientToken) > 0 {
		input.ClientToken = aws.String(_qconnectClientToken)
	}
	if len(_qconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _qconnectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAssistantAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates Amazon Q in Connect content. Before to calling this API, use [StartContentUpload] to upload
// an asset.
//
// [StartContentUpload]: https://docs.aws.amazon.com/amazon-q-connect/latest/APIReference/API_StartContentUpload.html
func qconnect_CreateContent(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.CreateContentInput{
		// KnowledgeBaseId: *string, // Required
		// Name: *string, // Required
		// UploadId: *string, // Required
	}

	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectName) > 0 {
		input.Name = aws.String(_qconnectName)
	}
	if len(_qconnectUploadId) > 0 {
		input.UploadId = aws.String(_qconnectUploadId)
	}
	if len(_qconnectClientToken) > 0 {
		input.ClientToken = aws.String(_qconnectClientToken)
	}
	if len(_qconnectMetadata) > 0 {
		if err := assignInputField(input, "Metadata", _qconnectMetadata); err != nil {
			log.Errorf("invalid --metadata: %s", err.Error())
			return
		}
	}
	if len(_qconnectOverrideLinkOutUri) > 0 {
		input.OverrideLinkOutUri = aws.String(_qconnectOverrideLinkOutUri)
	}
	if len(_qconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _qconnectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_qconnectTitle) > 0 {
		input.Title = aws.String(_qconnectTitle)
	}

	if resp, err := client.CreateContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an association between a content resource in a knowledge base and [step-by-step guides].
// Step-by-step guides offer instructions to agents for resolving common customer
// issues. You create a content association to integrate Amazon Q in Connect and
// step-by-step guides.
//
// After you integrate Amazon Q and step-by-step guides, when Amazon Q provides a
// recommendation to an agent based on the intent that it's detected, it also
// provides them with the option to start the step-by-step guide that you have
// associated with the content.
//
// Note the following limitations:
//
// - You can create only one content association for each content resource in a
// knowledge base.
//
// - You can associate a step-by-step guide with multiple content resources.
//
// For more information, see [Integrate Amazon Q in Connect with step-by-step guides] in the Amazon Connect Administrator Guide.
//
// [Integrate Amazon Q in Connect with step-by-step guides]: https://docs.aws.amazon.com/connect/latest/adminguide/integrate-q-with-guides.html
// [step-by-step guides]: https://docs.aws.amazon.com/connect/latest/adminguide/step-by-step-guided-experiences.html
func qconnect_CreateContentAssociation(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.CreateContentAssociationInput{
		// Association: types.ContentAssociationContents, // Required
		// AssociationType: types.ContentAssociationType, // Required
		// ContentId: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_qconnectAssociation) > 0 {
		if err := assignInputField(input, "Association", _qconnectAssociation); err != nil {
			log.Errorf("invalid --association: %s", err.Error())
			return
		}
	}
	if len(_qconnectAssociationType) > 0 {
		if err := assignInputField(input, "AssociationType", _qconnectAssociationType); err != nil {
			log.Errorf("invalid --association-type: %s", err.Error())
			return
		}
	}
	if len(_qconnectContentId) > 0 {
		input.ContentId = aws.String(_qconnectContentId)
	}
	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectClientToken) > 0 {
		input.ClientToken = aws.String(_qconnectClientToken)
	}
	if len(_qconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _qconnectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateContentAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a knowledge base.
// When using this API, you cannot reuse [Amazon AppIntegrations] DataIntegrations with external knowledge
// bases such as Salesforce and ServiceNow. If you do, you'll get an
// InvalidRequestException error.
//
// For example, you're programmatically managing your external knowledge base, and
// you want to add or remove one of the fields that is being ingested from
// Salesforce. Do the following:
//
// - Call [DeleteKnowledgeBase].
//
// - Call [DeleteDataIntegration].
//
// - Call [CreateDataIntegration]to recreate the DataIntegration or a create different one.
//
// - Call CreateKnowledgeBase.
//
// [Amazon AppIntegrations]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/Welcome.html
// [DeleteKnowledgeBase]: https://docs.aws.amazon.com/amazon-q-connect/latest/APIReference/API_DeleteKnowledgeBase.html
// [DeleteDataIntegration]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_DeleteDataIntegration.html
// [CreateDataIntegration]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_CreateDataIntegration.html
func qconnect_CreateKnowledgeBase(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.CreateKnowledgeBaseInput{
		// KnowledgeBaseType: types.KnowledgeBaseType, // Required
		// Name: *string, // Required
	}

	if len(_qconnectKnowledgeBaseType) > 0 {
		if err := assignInputField(input, "KnowledgeBaseType", _qconnectKnowledgeBaseType); err != nil {
			log.Errorf("invalid --knowledge-base-type: %s", err.Error())
			return
		}
	}
	if len(_qconnectName) > 0 {
		input.Name = aws.String(_qconnectName)
	}
	if len(_qconnectClientToken) > 0 {
		input.ClientToken = aws.String(_qconnectClientToken)
	}
	if len(_qconnectDescription) > 0 {
		input.Description = aws.String(_qconnectDescription)
	}
	if len(_qconnectRenderingConfiguration) > 0 {
		if err := assignInputField(input, "RenderingConfiguration", _qconnectRenderingConfiguration); err != nil {
			log.Errorf("invalid --rendering-configuration: %s", err.Error())
			return
		}
	}
	if len(_qconnectServerSideEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "ServerSideEncryptionConfiguration", _qconnectServerSideEncryptionConfiguration); err != nil {
			log.Errorf("invalid --server-side-encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_qconnectSourceConfiguration) > 0 {
		if err := assignInputField(input, "SourceConfiguration", _qconnectSourceConfiguration); err != nil {
			log.Errorf("invalid --source-configuration: %s", err.Error())
			return
		}
	}
	if len(_qconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _qconnectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_qconnectVectorIngestionConfiguration) > 0 {
		if err := assignInputField(input, "VectorIngestionConfiguration", _qconnectVectorIngestionConfiguration); err != nil {
			log.Errorf("invalid --vector-ingestion-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateKnowledgeBase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Q in Connect message template. The name of the message
// template has to be unique for each knowledge base. The channel subtype of the
// message template is immutable and cannot be modified after creation. After the
// message template is created, you can use the $LATEST qualifier to reference the
// created message template.
func qconnect_CreateMessageTemplate(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.CreateMessageTemplateInput{
		// ChannelSubtype: types.ChannelSubtype, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_qconnectChannelSubtype) > 0 {
		if err := assignInputField(input, "ChannelSubtype", _qconnectChannelSubtype); err != nil {
			log.Errorf("invalid --channel-subtype: %s", err.Error())
			return
		}
	}
	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectClientToken) > 0 {
		input.ClientToken = aws.String(_qconnectClientToken)
	}
	if len(_qconnectContent) > 0 {
		if err := assignInputField(input, "Content", _qconnectContent); err != nil {
			log.Errorf("invalid --content: %s", err.Error())
			return
		}
	}
	if len(_qconnectDefaultAttributes) > 0 {
		if err := assignInputField(input, "DefaultAttributes", _qconnectDefaultAttributes); err != nil {
			log.Errorf("invalid --default-attributes: %s", err.Error())
			return
		}
	}
	if len(_qconnectDescription) > 0 {
		input.Description = aws.String(_qconnectDescription)
	}
	if len(_qconnectGroupingConfiguration) > 0 {
		if err := assignInputField(input, "GroupingConfiguration", _qconnectGroupingConfiguration); err != nil {
			log.Errorf("invalid --grouping-configuration: %s", err.Error())
			return
		}
	}
	if len(_qconnectLanguage) > 0 {
		input.Language = aws.String(_qconnectLanguage)
	}
	if len(_qconnectName) > 0 {
		input.Name = aws.String(_qconnectName)
	}
	if len(_qconnectSourceConfiguration) > 0 {
		if err := assignInputField(input, "SourceConfiguration", _qconnectSourceConfiguration); err != nil {
			log.Errorf("invalid --source-configuration: %s", err.Error())
			return
		}
	}
	if len(_qconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _qconnectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMessageTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Uploads an attachment file to the specified Amazon Q in Connect message
// template. The name of the message template attachment has to be unique for each
// message template referenced by the $LATEST qualifier. The body of the
// attachment file should be encoded using base64 encoding. After the file is
// uploaded, you can use the pre-signed Amazon S3 URL returned in response to
// download the uploaded file.
func qconnect_CreateMessageTemplateAttachment(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.CreateMessageTemplateAttachmentInput{
		// Body: *string, // Required
		// ContentDisposition: types.ContentDisposition, // Required
		// KnowledgeBaseId: *string, // Required
		// MessageTemplateId: *string, // Required
		// Name: *string, // Required
	}

	if len(_qconnectBody) > 0 {
		input.Body = aws.String(_qconnectBody)
	}
	if len(_qconnectContentDisposition) > 0 {
		if err := assignInputField(input, "ContentDisposition", _qconnectContentDisposition); err != nil {
			log.Errorf("invalid --content-disposition: %s", err.Error())
			return
		}
	}
	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectMessageTemplateId) > 0 {
		input.MessageTemplateId = aws.String(_qconnectMessageTemplateId)
	}
	if len(_qconnectName) > 0 {
		input.Name = aws.String(_qconnectName)
	}
	if len(_qconnectClientToken) > 0 {
		input.ClientToken = aws.String(_qconnectClientToken)
	}

	if resp, err := client.CreateMessageTemplateAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Amazon Q in Connect message template version from the current
// content and configuration of a message template. Versions are immutable and
// monotonically increasing. Once a version is created, you can reference a
// specific version of the message template by passing in
// <message-template-id>:<versionNumber> as the message template identifier. An
// error is displayed if the supplied messageTemplateContentSha256 is different
// from the messageTemplateContentSha256 of the message template with $LATEST
// qualifier. If multiple CreateMessageTemplateVersion requests are made while the
// message template remains the same, only the first invocation creates a new
// version and the succeeding requests will return the same response as the first
// invocation.
func qconnect_CreateMessageTemplateVersion(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.CreateMessageTemplateVersionInput{
		// KnowledgeBaseId: *string, // Required
		// MessageTemplateId: *string, // Required
	}

	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectMessageTemplateId) > 0 {
		input.MessageTemplateId = aws.String(_qconnectMessageTemplateId)
	}
	if len(_qconnectMessageTemplateContentSha256) > 0 {
		input.MessageTemplateContentSha256 = aws.String(_qconnectMessageTemplateContentSha256)
	}

	if resp, err := client.CreateMessageTemplateVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Q in Connect quick response.
func qconnect_CreateQuickResponse(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.CreateQuickResponseInput{
		// Content: types.QuickResponseDataProvider, // Required
		// KnowledgeBaseId: *string, // Required
		// Name: *string, // Required
	}

	if len(_qconnectContent) > 0 {
		if err := assignInputField(input, "Content", _qconnectContent); err != nil {
			log.Errorf("invalid --content: %s", err.Error())
			return
		}
	}
	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectName) > 0 {
		input.Name = aws.String(_qconnectName)
	}
	if len(_qconnectChannels) > 0 {
		input.Channels = append([]string(nil), _qconnectChannels...)
	}
	if len(_qconnectClientToken) > 0 {
		input.ClientToken = aws.String(_qconnectClientToken)
	}
	if len(_qconnectContentType) > 0 {
		input.ContentType = aws.String(_qconnectContentType)
	}
	if len(_qconnectDescription) > 0 {
		input.Description = aws.String(_qconnectDescription)
	}
	if len(_qconnectGroupingConfiguration) > 0 {
		if err := assignInputField(input, "GroupingConfiguration", _qconnectGroupingConfiguration); err != nil {
			log.Errorf("invalid --grouping-configuration: %s", err.Error())
			return
		}
	}
	if len(_qconnectIsActive) > 0 {
		if err := assignInputField(input, "IsActive", _qconnectIsActive); err != nil {
			log.Errorf("invalid --is-active: %s", err.Error())
			return
		}
	}
	if len(_qconnectLanguage) > 0 {
		input.Language = aws.String(_qconnectLanguage)
	}
	if len(_qconnectShortcutKey) > 0 {
		input.ShortcutKey = aws.String(_qconnectShortcutKey)
	}
	if len(_qconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _qconnectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateQuickResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a session. A session is a contextual container used for generating
// recommendations. Amazon Connect creates a new Amazon Q in Connect session for
// each contact on which Amazon Q in Connect is enabled.
func qconnect_CreateSession(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.CreateSessionInput{
		// AssistantId: *string, // Required
		// Name: *string, // Required
	}

	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectName) > 0 {
		input.Name = aws.String(_qconnectName)
	}
	if len(_qconnectAiAgentConfiguration) > 0 {
		if err := assignInputField(input, "AiAgentConfiguration", _qconnectAiAgentConfiguration); err != nil {
			log.Errorf("invalid --ai-agent-configuration: %s", err.Error())
			return
		}
	}
	if len(_qconnectClientToken) > 0 {
		input.ClientToken = aws.String(_qconnectClientToken)
	}
	if len(_qconnectContactArn) > 0 {
		input.ContactArn = aws.String(_qconnectContactArn)
	}
	if len(_qconnectDescription) > 0 {
		input.Description = aws.String(_qconnectDescription)
	}
	if len(_qconnectOrchestratorConfigurationList) > 0 {
		if err := assignInputField(input, "OrchestratorConfigurationList", _qconnectOrchestratorConfigurationList); err != nil {
			log.Errorf("invalid --orchestrator-configuration-list: %s", err.Error())
			return
		}
	}
	if len(_qconnectRemoveOrchestratorConfigurationList) > 0 {
		if err := assignInputField(input, "RemoveOrchestratorConfigurationList", _qconnectRemoveOrchestratorConfigurationList); err != nil {
			log.Errorf("invalid --remove-orchestrator-configuration-list: %s", err.Error())
			return
		}
	}
	if len(_qconnectTagFilter) > 0 {
		if err := assignInputField(input, "TagFilter", _qconnectTagFilter); err != nil {
			log.Errorf("invalid --tag-filter: %s", err.Error())
			return
		}
	}
	if len(_qconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _qconnectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deactivates a specific version of the Amazon Q in Connect message template .
// After the version is deactivated, you can no longer use the $ACTIVE_VERSION
// qualifier to reference the version in active status.
func qconnect_DeactivateMessageTemplate(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.DeactivateMessageTemplateInput{
		// KnowledgeBaseId: *string, // Required
		// MessageTemplateId: *string, // Required
		// VersionNumber: *int64, // Required
	}

	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectMessageTemplateId) > 0 {
		input.MessageTemplateId = aws.String(_qconnectMessageTemplateId)
	}
	if len(_qconnectVersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _qconnectVersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeactivateMessageTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Q in Connect AI Agent.
func qconnect_DeleteAIAgent(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.DeleteAIAgentInput{
		// AiAgentId: *string, // Required
		// AssistantId: *string, // Required
	}

	if len(_qconnectAiAgentId) > 0 {
		input.AiAgentId = aws.String(_qconnectAiAgentId)
	}
	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}

	if resp, err := client.DeleteAIAgent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Q in Connect AI Agent Version.
func qconnect_DeleteAIAgentVersion(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.DeleteAIAgentVersionInput{
		// AiAgentId: *string, // Required
		// AssistantId: *string, // Required
		// VersionNumber: *int64, // Required
	}

	if len(_qconnectAiAgentId) > 0 {
		input.AiAgentId = aws.String(_qconnectAiAgentId)
	}
	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectVersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _qconnectVersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteAIAgentVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Q in Connect AI Guardrail.
func qconnect_DeleteAIGuardrail(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.DeleteAIGuardrailInput{
		// AiGuardrailId: *string, // Required
		// AssistantId: *string, // Required
	}

	if len(_qconnectAiGuardrailId) > 0 {
		input.AiGuardrailId = aws.String(_qconnectAiGuardrailId)
	}
	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}

	if resp, err := client.DeleteAIGuardrail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete and Amazon Q in Connect AI Guardrail version.
func qconnect_DeleteAIGuardrailVersion(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.DeleteAIGuardrailVersionInput{
		// AiGuardrailId: *string, // Required
		// AssistantId: *string, // Required
		// VersionNumber: *int64, // Required
	}

	if len(_qconnectAiGuardrailId) > 0 {
		input.AiGuardrailId = aws.String(_qconnectAiGuardrailId)
	}
	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectVersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _qconnectVersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteAIGuardrailVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Q in Connect AI Prompt.
func qconnect_DeleteAIPrompt(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.DeleteAIPromptInput{
		// AiPromptId: *string, // Required
		// AssistantId: *string, // Required
	}

	if len(_qconnectAiPromptId) > 0 {
		input.AiPromptId = aws.String(_qconnectAiPromptId)
	}
	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}

	if resp, err := client.DeleteAIPrompt(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete and Amazon Q in Connect AI Prompt version.
func qconnect_DeleteAIPromptVersion(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.DeleteAIPromptVersionInput{
		// AiPromptId: *string, // Required
		// AssistantId: *string, // Required
		// VersionNumber: *int64, // Required
	}

	if len(_qconnectAiPromptId) > 0 {
		input.AiPromptId = aws.String(_qconnectAiPromptId)
	}
	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectVersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _qconnectVersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteAIPromptVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an assistant.
func qconnect_DeleteAssistant(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.DeleteAssistantInput{
		// AssistantId: *string, // Required
	}

	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}

	if resp, err := client.DeleteAssistant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an assistant association.
func qconnect_DeleteAssistantAssociation(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.DeleteAssistantAssociationInput{
		// AssistantAssociationId: *string, // Required
		// AssistantId: *string, // Required
	}

	if len(_qconnectAssistantAssociationId) > 0 {
		input.AssistantAssociationId = aws.String(_qconnectAssistantAssociationId)
	}
	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}

	if resp, err := client.DeleteAssistantAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the content.
func qconnect_DeleteContent(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.DeleteContentInput{
		// ContentId: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_qconnectContentId) > 0 {
		input.ContentId = aws.String(_qconnectContentId)
	}
	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}

	if resp, err := client.DeleteContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the content association.
// For more information about content associations--what they are and when they
// are used--see [Integrate Amazon Q in Connect with step-by-step guides]in the Amazon Connect Administrator Guide.
//
// [Integrate Amazon Q in Connect with step-by-step guides]: https://docs.aws.amazon.com/connect/latest/adminguide/integrate-q-with-guides.html
func qconnect_DeleteContentAssociation(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.DeleteContentAssociationInput{
		// ContentAssociationId: *string, // Required
		// ContentId: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_qconnectContentAssociationId) > 0 {
		input.ContentAssociationId = aws.String(_qconnectContentAssociationId)
	}
	if len(_qconnectContentId) > 0 {
		input.ContentId = aws.String(_qconnectContentId)
	}
	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}

	if resp, err := client.DeleteContentAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the quick response import job.
func qconnect_DeleteImportJob(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.DeleteImportJobInput{
		// ImportJobId: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_qconnectImportJobId) > 0 {
		input.ImportJobId = aws.String(_qconnectImportJobId)
	}
	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}

	if resp, err := client.DeleteImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the knowledge base.
// When you use this API to delete an external knowledge base such as Salesforce
// or ServiceNow, you must also delete the [Amazon AppIntegrations]DataIntegration. This is because you
// can't reuse the DataIntegration after it's been associated with an external
// knowledge base. However, you can delete and recreate it. See [DeleteDataIntegration]and [CreateDataIntegration] in the Amazon
// AppIntegrations API Reference.
//
// [Amazon AppIntegrations]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/Welcome.html
// [DeleteDataIntegration]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_DeleteDataIntegration.html
// [CreateDataIntegration]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_CreateDataIntegration.html
func qconnect_DeleteKnowledgeBase(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.DeleteKnowledgeBaseInput{
		// KnowledgeBaseId: *string, // Required
	}

	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}

	if resp, err := client.DeleteKnowledgeBase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Q in Connect message template entirely or a specific version
// of the message template if version is supplied in the request. You can provide
// the message template identifier as <message-template-id>:<versionNumber> to
// delete a specific version of the message template. If it is not supplied, the
// message template and all available versions will be deleted.
func qconnect_DeleteMessageTemplate(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.DeleteMessageTemplateInput{
		// KnowledgeBaseId: *string, // Required
		// MessageTemplateId: *string, // Required
	}

	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectMessageTemplateId) > 0 {
		input.MessageTemplateId = aws.String(_qconnectMessageTemplateId)
	}

	if resp, err := client.DeleteMessageTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the attachment file from the Amazon Q in Connect message template that
// is referenced by $LATEST qualifier. Attachments on available message template
// versions will remain unchanged.
func qconnect_DeleteMessageTemplateAttachment(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.DeleteMessageTemplateAttachmentInput{
		// AttachmentId: *string, // Required
		// KnowledgeBaseId: *string, // Required
		// MessageTemplateId: *string, // Required
	}

	if len(_qconnectAttachmentId) > 0 {
		input.AttachmentId = aws.String(_qconnectAttachmentId)
	}
	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectMessageTemplateId) > 0 {
		input.MessageTemplateId = aws.String(_qconnectMessageTemplateId)
	}

	if resp, err := client.DeleteMessageTemplateAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a quick response.
func qconnect_DeleteQuickResponse(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.DeleteQuickResponseInput{
		// KnowledgeBaseId: *string, // Required
		// QuickResponseId: *string, // Required
	}

	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectQuickResponseId) > 0 {
		input.QuickResponseId = aws.String(_qconnectQuickResponseId)
	}

	if resp, err := client.DeleteQuickResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an Amazon Q in Connect AI Agent.
func qconnect_GetAIAgent(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.GetAIAgentInput{
		// AiAgentId: *string, // Required
		// AssistantId: *string, // Required
	}

	if len(_qconnectAiAgentId) > 0 {
		input.AiAgentId = aws.String(_qconnectAiAgentId)
	}
	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}

	if resp, err := client.GetAIAgent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the Amazon Q in Connect AI Guardrail.
func qconnect_GetAIGuardrail(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.GetAIGuardrailInput{
		// AiGuardrailId: *string, // Required
		// AssistantId: *string, // Required
	}

	if len(_qconnectAiGuardrailId) > 0 {
		input.AiGuardrailId = aws.String(_qconnectAiGuardrailId)
	}
	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}

	if resp, err := client.GetAIGuardrail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets and Amazon Q in Connect AI Prompt.
func qconnect_GetAIPrompt(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.GetAIPromptInput{
		// AiPromptId: *string, // Required
		// AssistantId: *string, // Required
	}

	if len(_qconnectAiPromptId) > 0 {
		input.AiPromptId = aws.String(_qconnectAiPromptId)
	}
	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}

	if resp, err := client.GetAIPrompt(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an assistant.
func qconnect_GetAssistant(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.GetAssistantInput{
		// AssistantId: *string, // Required
	}

	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}

	if resp, err := client.GetAssistant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an assistant association.
func qconnect_GetAssistantAssociation(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.GetAssistantAssociationInput{
		// AssistantAssociationId: *string, // Required
		// AssistantId: *string, // Required
	}

	if len(_qconnectAssistantAssociationId) > 0 {
		input.AssistantAssociationId = aws.String(_qconnectAssistantAssociationId)
	}
	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}

	if resp, err := client.GetAssistantAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves content, including a pre-signed URL to download the content.
func qconnect_GetContent(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.GetContentInput{
		// ContentId: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_qconnectContentId) > 0 {
		input.ContentId = aws.String(_qconnectContentId)
	}
	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}

	if resp, err := client.GetContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the content association.
// For more information about content associations--what they are and when they
// are used--see [Integrate Amazon Q in Connect with step-by-step guides]in the Amazon Connect Administrator Guide.
//
// [Integrate Amazon Q in Connect with step-by-step guides]: https://docs.aws.amazon.com/connect/latest/adminguide/integrate-q-with-guides.html
func qconnect_GetContentAssociation(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.GetContentAssociationInput{
		// ContentAssociationId: *string, // Required
		// ContentId: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_qconnectContentAssociationId) > 0 {
		input.ContentAssociationId = aws.String(_qconnectContentAssociationId)
	}
	if len(_qconnectContentId) > 0 {
		input.ContentId = aws.String(_qconnectContentId)
	}
	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}

	if resp, err := client.GetContentAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves summary information about the content.
func qconnect_GetContentSummary(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.GetContentSummaryInput{
		// ContentId: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_qconnectContentId) > 0 {
		input.ContentId = aws.String(_qconnectContentId)
	}
	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}

	if resp, err := client.GetContentSummary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the started import job.
func qconnect_GetImportJob(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.GetImportJobInput{
		// ImportJobId: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_qconnectImportJobId) > 0 {
		input.ImportJobId = aws.String(_qconnectImportJobId)
	}
	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}

	if resp, err := client.GetImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the knowledge base.
func qconnect_GetKnowledgeBase(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.GetKnowledgeBaseInput{
		// KnowledgeBaseId: *string, // Required
	}

	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}

	if resp, err := client.GetKnowledgeBase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the Amazon Q in Connect message template. The message template
// identifier can contain an optional qualifier, for example,
// <message-template-id>:<qualifier> , which is either an actual version number or
// an Amazon Q Connect managed qualifier $ACTIVE_VERSION | $LATEST . If it is not
// supplied, then $LATEST is assumed implicitly.
func qconnect_GetMessageTemplate(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.GetMessageTemplateInput{
		// KnowledgeBaseId: *string, // Required
		// MessageTemplateId: *string, // Required
	}

	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectMessageTemplateId) > 0 {
		input.MessageTemplateId = aws.String(_qconnectMessageTemplateId)
	}

	if resp, err := client.GetMessageTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves next message on an Amazon Q in Connect session.
func qconnect_GetNextMessage(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.GetNextMessageInput{
		// AssistantId: *string, // Required
		// NextMessageToken: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectNextMessageToken) > 0 {
		input.NextMessageToken = aws.String(_qconnectNextMessageToken)
	}
	if len(_qconnectSessionId) > 0 {
		input.SessionId = aws.String(_qconnectSessionId)
	}

	if resp, err := client.GetNextMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the quick response.
func qconnect_GetQuickResponse(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.GetQuickResponseInput{
		// KnowledgeBaseId: *string, // Required
		// QuickResponseId: *string, // Required
	}

	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectQuickResponseId) > 0 {
		input.QuickResponseId = aws.String(_qconnectQuickResponseId)
	}

	if resp, err := client.GetQuickResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API will be discontinued starting June 1, 2024. To receive generative
// responses after March 1, 2024, you will need to create a new Assistant in the
// Amazon Connect console and integrate the Amazon Q in Connect JavaScript library
// (amazon-q-connectjs) into your applications.
//
// Retrieves recommendations for the specified session. To avoid retrieving the
// same recommendations in subsequent calls, use [NotifyRecommendationsReceived]. This API supports long-polling
// behavior with the waitTimeSeconds parameter. Short poll is the default behavior
// and only returns recommendations already available. To perform a manual query
// against an assistant, use [QueryAssistant].
//
// Deprecated: GetRecommendations API will be discontinued starting June 1, 2024.
// To receive generative responses after March 1, 2024 you will need to create a
// new Assistant in the Connect console and integrate the Amazon Q in Connect
// JavaScript library (amazon-q-connectjs) into your applications.
//
// [QueryAssistant]: https://docs.aws.amazon.com/amazon-q-connect/latest/APIReference/API_QueryAssistant.html
// [NotifyRecommendationsReceived]: https://docs.aws.amazon.com/amazon-q-connect/latest/APIReference/API_NotifyRecommendationsReceived.html
func qconnect_GetRecommendations(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.GetRecommendationsInput{
		// AssistantId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectSessionId) > 0 {
		input.SessionId = aws.String(_qconnectSessionId)
	}
	if len(_qconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qconnectNextChunkToken) > 0 {
		input.NextChunkToken = aws.String(_qconnectNextChunkToken)
	}
	if len(_qconnectRecommendationType) > 0 {
		if err := assignInputField(input, "RecommendationType", _qconnectRecommendationType); err != nil {
			log.Errorf("invalid --recommendation-type: %s", err.Error())
			return
		}
	}
	if len(_qconnectWaitTimeSeconds) > 0 {
		if err := assignInputField(input, "WaitTimeSeconds", _qconnectWaitTimeSeconds); err != nil {
			log.Errorf("invalid --wait-time-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information for a specified session.
func qconnect_GetSession(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.GetSessionInput{
		// AssistantId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectSessionId) > 0 {
		input.SessionId = aws.String(_qconnectSessionId)
	}

	if resp, err := client.GetSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List AI Agent versions.
func qconnect_ListAIAgentVersions(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.ListAIAgentVersionsInput{
		// AiAgentId: *string, // Required
		// AssistantId: *string, // Required
	}

	if len(_qconnectAiAgentId) > 0 {
		input.AiAgentId = aws.String(_qconnectAiAgentId)
	}
	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qconnectNextToken) > 0 {
		input.NextToken = aws.String(_qconnectNextToken)
	}
	if len(_qconnectOrigin) > 0 {
		if err := assignInputField(input, "Origin", _qconnectOrigin); err != nil {
			log.Errorf("invalid --origin: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAIAgentVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qconnect.ListAIAgentVersionsOutput
	p := qconnect.NewListAIAgentVersionsPaginator(client, input)
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

// Lists AI Agents.
func qconnect_ListAIAgents(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.ListAIAgentsInput{
		// AssistantId: *string, // Required
	}

	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qconnectNextToken) > 0 {
		input.NextToken = aws.String(_qconnectNextToken)
	}
	if len(_qconnectOrigin) > 0 {
		if err := assignInputField(input, "Origin", _qconnectOrigin); err != nil {
			log.Errorf("invalid --origin: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAIAgents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qconnect.ListAIAgentsOutput
	p := qconnect.NewListAIAgentsPaginator(client, input)
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

// Lists AI Guardrail versions.
func qconnect_ListAIGuardrailVersions(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.ListAIGuardrailVersionsInput{
		// AiGuardrailId: *string, // Required
		// AssistantId: *string, // Required
	}

	if len(_qconnectAiGuardrailId) > 0 {
		input.AiGuardrailId = aws.String(_qconnectAiGuardrailId)
	}
	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qconnectNextToken) > 0 {
		input.NextToken = aws.String(_qconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAIGuardrailVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qconnect.ListAIGuardrailVersionsOutput
	p := qconnect.NewListAIGuardrailVersionsPaginator(client, input)
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

// Lists the AI Guardrails available on the Amazon Q in Connect assistant.
func qconnect_ListAIGuardrails(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.ListAIGuardrailsInput{
		// AssistantId: *string, // Required
	}

	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qconnectNextToken) > 0 {
		input.NextToken = aws.String(_qconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAIGuardrails(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qconnect.ListAIGuardrailsOutput
	p := qconnect.NewListAIGuardrailsPaginator(client, input)
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

// Lists AI Prompt versions.
func qconnect_ListAIPromptVersions(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.ListAIPromptVersionsInput{
		// AiPromptId: *string, // Required
		// AssistantId: *string, // Required
	}

	if len(_qconnectAiPromptId) > 0 {
		input.AiPromptId = aws.String(_qconnectAiPromptId)
	}
	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qconnectNextToken) > 0 {
		input.NextToken = aws.String(_qconnectNextToken)
	}
	if len(_qconnectOrigin) > 0 {
		if err := assignInputField(input, "Origin", _qconnectOrigin); err != nil {
			log.Errorf("invalid --origin: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAIPromptVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qconnect.ListAIPromptVersionsOutput
	p := qconnect.NewListAIPromptVersionsPaginator(client, input)
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

// Lists the AI Prompts available on the Amazon Q in Connect assistant.
func qconnect_ListAIPrompts(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.ListAIPromptsInput{
		// AssistantId: *string, // Required
	}

	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qconnectNextToken) > 0 {
		input.NextToken = aws.String(_qconnectNextToken)
	}
	if len(_qconnectOrigin) > 0 {
		if err := assignInputField(input, "Origin", _qconnectOrigin); err != nil {
			log.Errorf("invalid --origin: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAIPrompts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qconnect.ListAIPromptsOutput
	p := qconnect.NewListAIPromptsPaginator(client, input)
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

// Lists information about assistant associations.
func qconnect_ListAssistantAssociations(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.ListAssistantAssociationsInput{
		// AssistantId: *string, // Required
	}

	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qconnectNextToken) > 0 {
		input.NextToken = aws.String(_qconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssistantAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qconnect.ListAssistantAssociationsOutput
	p := qconnect.NewListAssistantAssociationsPaginator(client, input)
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

// Lists information about assistants.
func qconnect_ListAssistants(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.ListAssistantsInput{}

	if len(_qconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qconnectNextToken) > 0 {
		input.NextToken = aws.String(_qconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssistants(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qconnect.ListAssistantsOutput
	p := qconnect.NewListAssistantsPaginator(client, input)
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

// Lists the content associations.
// For more information about content associations--what they are and when they
// are used--see [Integrate Amazon Q in Connect with step-by-step guides]in the Amazon Connect Administrator Guide.
//
// [Integrate Amazon Q in Connect with step-by-step guides]: https://docs.aws.amazon.com/connect/latest/adminguide/integrate-q-with-guides.html
func qconnect_ListContentAssociations(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.ListContentAssociationsInput{
		// ContentId: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_qconnectContentId) > 0 {
		input.ContentId = aws.String(_qconnectContentId)
	}
	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qconnectNextToken) > 0 {
		input.NextToken = aws.String(_qconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListContentAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qconnect.ListContentAssociationsOutput
	p := qconnect.NewListContentAssociationsPaginator(client, input)
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

// Lists the content.
func qconnect_ListContents(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.ListContentsInput{
		// KnowledgeBaseId: *string, // Required
	}

	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qconnectNextToken) > 0 {
		input.NextToken = aws.String(_qconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListContents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qconnect.ListContentsOutput
	p := qconnect.NewListContentsPaginator(client, input)
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

// Lists information about import jobs.
func qconnect_ListImportJobs(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.ListImportJobsInput{
		// KnowledgeBaseId: *string, // Required
	}

	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qconnectNextToken) > 0 {
		input.NextToken = aws.String(_qconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListImportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qconnect.ListImportJobsOutput
	p := qconnect.NewListImportJobsPaginator(client, input)
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

// Lists the knowledge bases.
func qconnect_ListKnowledgeBases(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.ListKnowledgeBasesInput{}

	if len(_qconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qconnectNextToken) > 0 {
		input.NextToken = aws.String(_qconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListKnowledgeBases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qconnect.ListKnowledgeBasesOutput
	p := qconnect.NewListKnowledgeBasesPaginator(client, input)
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

// Lists all the available versions for the specified Amazon Q in Connect message
// template.
func qconnect_ListMessageTemplateVersions(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.ListMessageTemplateVersionsInput{
		// KnowledgeBaseId: *string, // Required
		// MessageTemplateId: *string, // Required
	}

	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectMessageTemplateId) > 0 {
		input.MessageTemplateId = aws.String(_qconnectMessageTemplateId)
	}
	if len(_qconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qconnectNextToken) > 0 {
		input.NextToken = aws.String(_qconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMessageTemplateVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qconnect.ListMessageTemplateVersionsOutput
	p := qconnect.NewListMessageTemplateVersionsPaginator(client, input)
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

// Lists all the available Amazon Q in Connect message templates for the specified
// knowledge base.
func qconnect_ListMessageTemplates(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.ListMessageTemplatesInput{
		// KnowledgeBaseId: *string, // Required
	}

	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qconnectNextToken) > 0 {
		input.NextToken = aws.String(_qconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMessageTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qconnect.ListMessageTemplatesOutput
	p := qconnect.NewListMessageTemplatesPaginator(client, input)
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

// Lists messages on an Amazon Q in Connect session.
func qconnect_ListMessages(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.ListMessagesInput{
		// AssistantId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectSessionId) > 0 {
		input.SessionId = aws.String(_qconnectSessionId)
	}
	if len(_qconnectFilter) > 0 {
		if err := assignInputField(input, "Filter", _qconnectFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_qconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qconnectNextToken) > 0 {
		input.NextToken = aws.String(_qconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMessages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qconnect.ListMessagesOutput
	p := qconnect.NewListMessagesPaginator(client, input)
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

// Lists information about quick response.
func qconnect_ListQuickResponses(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.ListQuickResponsesInput{
		// KnowledgeBaseId: *string, // Required
	}

	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qconnectNextToken) > 0 {
		input.NextToken = aws.String(_qconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListQuickResponses(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qconnect.ListQuickResponsesOutput
	p := qconnect.NewListQuickResponsesPaginator(client, input)
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

// Retrieves AI agent execution traces for a session, providing granular
// visibility into agent orchestration flows, LLM interactions, and tool
// invocations.
func qconnect_ListSpans(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.ListSpansInput{
		// AssistantId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectSessionId) > 0 {
		input.SessionId = aws.String(_qconnectSessionId)
	}
	if len(_qconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qconnectNextToken) > 0 {
		input.NextToken = aws.String(_qconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSpans(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qconnect.ListSpansOutput
	p := qconnect.NewListSpansPaginator(client, input)
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

// Lists the tags for the specified resource.
func qconnect_ListTagsForResource(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_qconnectResourceArn) > 0 {
		input.ResourceArn = aws.String(_qconnectResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified recommendations from the specified assistant's queue of
// newly available recommendations. You can use this API in conjunction with [GetRecommendations]and a
// waitTimeSeconds input for long-polling behavior and avoiding duplicate
// recommendations.
//
// [GetRecommendations]: https://docs.aws.amazon.com/amazon-q-connect/latest/APIReference/API_GetRecommendations.html
func qconnect_NotifyRecommendationsReceived(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.NotifyRecommendationsReceivedInput{
		// AssistantId: *string, // Required
		// RecommendationIds: []string, // Required
		// SessionId: *string, // Required
	}

	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectRecommendationIds) > 0 {
		input.RecommendationIds = append([]string(nil), _qconnectRecommendationIds...)
	}
	if len(_qconnectSessionId) > 0 {
		input.SessionId = aws.String(_qconnectSessionId)
	}

	if resp, err := client.NotifyRecommendationsReceived(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides feedback against the specified assistant for the specified target.
// This API only supports generative targets.
func qconnect_PutFeedback(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.PutFeedbackInput{
		// AssistantId: *string, // Required
		// ContentFeedback: types.ContentFeedbackData, // Required
		// TargetId: *string, // Required
		// TargetType: types.TargetType, // Required
	}

	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectContentFeedback) > 0 {
		if err := assignInputField(input, "ContentFeedback", _qconnectContentFeedback); err != nil {
			log.Errorf("invalid --content-feedback: %s", err.Error())
			return
		}
	}
	if len(_qconnectTargetId) > 0 {
		input.TargetId = aws.String(_qconnectTargetId)
	}
	if len(_qconnectTargetType) > 0 {
		if err := assignInputField(input, "TargetType", _qconnectTargetType); err != nil {
			log.Errorf("invalid --target-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutFeedback(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API will be discontinued starting June 1, 2024. To receive generative
// responses after March 1, 2024, you will need to create a new Assistant in the
// Amazon Connect console and integrate the Amazon Q in Connect JavaScript library
// (amazon-q-connectjs) into your applications.
//
// Performs a manual search against the specified assistant. To retrieve
// recommendations for an assistant, use [GetRecommendations].
//
// Deprecated: QueryAssistant API will be discontinued starting June 1, 2024. To
// receive generative responses after March 1, 2024 you will need to create a new
// Assistant in the Connect console and integrate the Amazon Q in Connect
// JavaScript library (amazon-q-connectjs) into your applications.
//
// [GetRecommendations]: https://docs.aws.amazon.com/amazon-q-connect/latest/APIReference/API_GetRecommendations.html
func qconnect_QueryAssistant(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.QueryAssistantInput{
		// AssistantId: *string, // Required
	}

	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qconnectNextToken) > 0 {
		input.NextToken = aws.String(_qconnectNextToken)
	}
	if len(_qconnectOverrideKnowledgeBaseSearchType) > 0 {
		if err := assignInputField(input, "OverrideKnowledgeBaseSearchType", _qconnectOverrideKnowledgeBaseSearchType); err != nil {
			log.Errorf("invalid --override-knowledge-base-search-type: %s", err.Error())
			return
		}
	}
	if len(_qconnectQueryCondition) > 0 {
		if err := assignInputField(input, "QueryCondition", _qconnectQueryCondition); err != nil {
			log.Errorf("invalid --query-condition: %s", err.Error())
			return
		}
	}
	if len(_qconnectQueryInputData) > 0 {
		if err := assignInputField(input, "QueryInputData", _qconnectQueryInputData); err != nil {
			log.Errorf("invalid --query-input-data: %s", err.Error())
			return
		}
	}
	if len(_qconnectQueryText) > 0 {
		input.QueryText = aws.String(_qconnectQueryText)
	}
	if len(_qconnectSessionId) > 0 {
		input.SessionId = aws.String(_qconnectSessionId)
	}

	if disablePaginator() {
		if resp, err := client.QueryAssistant(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qconnect.QueryAssistantOutput
	p := qconnect.NewQueryAssistantPaginator(client, input)
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

// Removes the AI Agent that is set for use by default on an Amazon Q in Connect
// Assistant.
func qconnect_RemoveAssistantAIAgent(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.RemoveAssistantAIAgentInput{
		// AiAgentType: types.AIAgentType, // Required
		// AssistantId: *string, // Required
	}

	if len(_qconnectAiAgentType) > 0 {
		if err := assignInputField(input, "AiAgentType", _qconnectAiAgentType); err != nil {
			log.Errorf("invalid --ai-agent-type: %s", err.Error())
			return
		}
	}
	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectOrchestratorUseCase) > 0 {
		input.OrchestratorUseCase = aws.String(_qconnectOrchestratorUseCase)
	}

	if resp, err := client.RemoveAssistantAIAgent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a URI template from a knowledge base.
func qconnect_RemoveKnowledgeBaseTemplateUri(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.RemoveKnowledgeBaseTemplateUriInput{
		// KnowledgeBaseId: *string, // Required
	}

	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}

	if resp, err := client.RemoveKnowledgeBaseTemplateUri(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Renders the Amazon Q in Connect message template based on the attribute values
// provided and generates the message content. For any variable present in the
// message template, if the attribute value is neither provided in the attribute
// request parameter nor the default attribute of the message template, the
// rendered message content will keep the variable placeholder as it is and return
// the attribute keys that are missing.
func qconnect_RenderMessageTemplate(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.RenderMessageTemplateInput{
		// Attributes: *types.MessageTemplateAttributes, // Required
		// KnowledgeBaseId: *string, // Required
		// MessageTemplateId: *string, // Required
	}

	if len(_qconnectAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _qconnectAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectMessageTemplateId) > 0 {
		input.MessageTemplateId = aws.String(_qconnectMessageTemplateId)
	}

	if resp, err := client.RenderMessageTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves content from knowledge sources based on a query.
func qconnect_Retrieve(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.RetrieveInput{
		// AssistantId: *string, // Required
		// RetrievalConfiguration: *types.RetrievalConfiguration, // Required
		// RetrievalQuery: *string, // Required
	}

	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectRetrievalConfiguration) > 0 {
		if err := assignInputField(input, "RetrievalConfiguration", _qconnectRetrievalConfiguration); err != nil {
			log.Errorf("invalid --retrieval-configuration: %s", err.Error())
			return
		}
	}
	if len(_qconnectRetrievalQuery) > 0 {
		input.RetrievalQuery = aws.String(_qconnectRetrievalQuery)
	}

	if resp, err := client.Retrieve(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches for content in a specified knowledge base. Can be used to get a
// specific content resource by its name.
func qconnect_SearchContent(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.SearchContentInput{
		// KnowledgeBaseId: *string, // Required
		// SearchExpression: *types.SearchExpression, // Required
	}

	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectSearchExpression) > 0 {
		if err := assignInputField(input, "SearchExpression", _qconnectSearchExpression); err != nil {
			log.Errorf("invalid --search-expression: %s", err.Error())
			return
		}
	}
	if len(_qconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qconnectNextToken) > 0 {
		input.NextToken = aws.String(_qconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchContent(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qconnect.SearchContentOutput
	p := qconnect.NewSearchContentPaginator(client, input)
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

// Searches for Amazon Q in Connect message templates in the specified knowledge
// base.
func qconnect_SearchMessageTemplates(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.SearchMessageTemplatesInput{
		// KnowledgeBaseId: *string, // Required
		// SearchExpression: *types.MessageTemplateSearchExpression, // Required
	}

	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectSearchExpression) > 0 {
		if err := assignInputField(input, "SearchExpression", _qconnectSearchExpression); err != nil {
			log.Errorf("invalid --search-expression: %s", err.Error())
			return
		}
	}
	if len(_qconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qconnectNextToken) > 0 {
		input.NextToken = aws.String(_qconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchMessageTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qconnect.SearchMessageTemplatesOutput
	p := qconnect.NewSearchMessageTemplatesPaginator(client, input)
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

// Searches existing Amazon Q in Connect quick responses in an Amazon Q in Connect
// knowledge base.
func qconnect_SearchQuickResponses(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.SearchQuickResponsesInput{
		// KnowledgeBaseId: *string, // Required
		// SearchExpression: *types.QuickResponseSearchExpression, // Required
	}

	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectSearchExpression) > 0 {
		if err := assignInputField(input, "SearchExpression", _qconnectSearchExpression); err != nil {
			log.Errorf("invalid --search-expression: %s", err.Error())
			return
		}
	}
	if len(_qconnectAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _qconnectAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_qconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qconnectNextToken) > 0 {
		input.NextToken = aws.String(_qconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchQuickResponses(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qconnect.SearchQuickResponsesOutput
	p := qconnect.NewSearchQuickResponsesPaginator(client, input)
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

// Searches for sessions.
func qconnect_SearchSessions(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.SearchSessionsInput{
		// AssistantId: *string, // Required
		// SearchExpression: *types.SearchExpression, // Required
	}

	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectSearchExpression) > 0 {
		if err := assignInputField(input, "SearchExpression", _qconnectSearchExpression); err != nil {
			log.Errorf("invalid --search-expression: %s", err.Error())
			return
		}
	}
	if len(_qconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qconnectNextToken) > 0 {
		input.NextToken = aws.String(_qconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchSessions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qconnect.SearchSessionsOutput
	p := qconnect.NewSearchSessionsPaginator(client, input)
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

// Submits a message to the Amazon Q in Connect session.
func qconnect_SendMessage(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.SendMessageInput{
		// AssistantId: *string, // Required
		// Message: *types.MessageInput, // Required
		// SessionId: *string, // Required
		// Type: types.MessageType, // Required
	}

	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectMessage) > 0 {
		if err := assignInputField(input, "Message", _qconnectMessage); err != nil {
			log.Errorf("invalid --message: %s", err.Error())
			return
		}
	}
	if len(_qconnectSessionId) > 0 {
		input.SessionId = aws.String(_qconnectSessionId)
	}
	if len(_qconnectType) > 0 {
		if err := assignInputField(input, "Type", _qconnectType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_qconnectAiAgentId) > 0 {
		input.AiAgentId = aws.String(_qconnectAiAgentId)
	}
	if len(_qconnectClientToken) > 0 {
		input.ClientToken = aws.String(_qconnectClientToken)
	}
	if len(_qconnectConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _qconnectConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_qconnectConversationContext) > 0 {
		if err := assignInputField(input, "ConversationContext", _qconnectConversationContext); err != nil {
			log.Errorf("invalid --conversation-context: %s", err.Error())
			return
		}
	}
	if len(_qconnectMetadata) > 0 {
		if err := assignInputField(input, "Metadata", _qconnectMetadata); err != nil {
			log.Errorf("invalid --metadata: %s", err.Error())
			return
		}
	}
	if len(_qconnectOrchestratorUseCase) > 0 {
		input.OrchestratorUseCase = aws.String(_qconnectOrchestratorUseCase)
	}

	if resp, err := client.SendMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get a URL to upload content to a knowledge base. To upload content, first make
// a PUT request to the returned URL with your file, making sure to include the
// required headers. Then use [CreateContent]to finalize the content creation process or [UpdateContent] to
// modify an existing resource. You can only upload content to a knowledge base of
// type CUSTOM.
//
// [CreateContent]: https://docs.aws.amazon.com/amazon-q-connect/latest/APIReference/API_CreateContent.html
// [UpdateContent]: https://docs.aws.amazon.com/amazon-q-connect/latest/APIReference/API_UpdateContent.html
func qconnect_StartContentUpload(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.StartContentUploadInput{
		// ContentType: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_qconnectContentType) > 0 {
		input.ContentType = aws.String(_qconnectContentType)
	}
	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectPresignedUrlTimeToLive) > 0 {
		if err := assignInputField(input, "PresignedUrlTimeToLive", _qconnectPresignedUrlTimeToLive); err != nil {
			log.Errorf("invalid --presigned-url-time-to-live: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartContentUpload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Start an asynchronous job to import Amazon Q in Connect resources from an
// uploaded source file. Before calling this API, use [StartContentUpload]to upload an asset that
// contains the resource data.
//
// - For importing Amazon Q in Connect quick responses, you need to upload a csv
// file including the quick responses. For information about how to format the csv
// file for importing quick responses, see [Import quick responses].
//
// [StartContentUpload]: https://docs.aws.amazon.com/wisdom/latest/APIReference/API_StartContentUpload.html
// [Import quick responses]: https://docs.aws.amazon.com/console/connect/quick-responses/add-data
func qconnect_StartImportJob(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.StartImportJobInput{
		// ImportJobType: types.ImportJobType, // Required
		// KnowledgeBaseId: *string, // Required
		// UploadId: *string, // Required
	}

	if len(_qconnectImportJobType) > 0 {
		if err := assignInputField(input, "ImportJobType", _qconnectImportJobType); err != nil {
			log.Errorf("invalid --import-job-type: %s", err.Error())
			return
		}
	}
	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectUploadId) > 0 {
		input.UploadId = aws.String(_qconnectUploadId)
	}
	if len(_qconnectClientToken) > 0 {
		input.ClientToken = aws.String(_qconnectClientToken)
	}
	if len(_qconnectExternalSourceConfiguration) > 0 {
		if err := assignInputField(input, "ExternalSourceConfiguration", _qconnectExternalSourceConfiguration); err != nil {
			log.Errorf("invalid --external-source-configuration: %s", err.Error())
			return
		}
	}
	if len(_qconnectMetadata) > 0 {
		if err := assignInputField(input, "Metadata", _qconnectMetadata); err != nil {
			log.Errorf("invalid --metadata: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the specified tags to the specified resource.
func qconnect_TagResource(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_qconnectResourceArn) > 0 {
		input.ResourceArn = aws.String(_qconnectResourceArn)
	}
	if len(_qconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _qconnectTags); err != nil {
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

// Removes the specified tags from the specified resource.
func qconnect_UntagResource(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_qconnectResourceArn) > 0 {
		input.ResourceArn = aws.String(_qconnectResourceArn)
	}
	if len(_qconnectTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _qconnectTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an AI Agent.
func qconnect_UpdateAIAgent(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.UpdateAIAgentInput{
		// AiAgentId: *string, // Required
		// AssistantId: *string, // Required
		// VisibilityStatus: types.VisibilityStatus, // Required
	}

	if len(_qconnectAiAgentId) > 0 {
		input.AiAgentId = aws.String(_qconnectAiAgentId)
	}
	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectVisibilityStatus) > 0 {
		if err := assignInputField(input, "VisibilityStatus", _qconnectVisibilityStatus); err != nil {
			log.Errorf("invalid --visibility-status: %s", err.Error())
			return
		}
	}
	if len(_qconnectClientToken) > 0 {
		input.ClientToken = aws.String(_qconnectClientToken)
	}
	if len(_qconnectConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _qconnectConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_qconnectDescription) > 0 {
		input.Description = aws.String(_qconnectDescription)
	}

	if resp, err := client.UpdateAIAgent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an AI Guardrail.
func qconnect_UpdateAIGuardrail(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.UpdateAIGuardrailInput{
		// AiGuardrailId: *string, // Required
		// AssistantId: *string, // Required
		// BlockedInputMessaging: *string, // Required
		// BlockedOutputsMessaging: *string, // Required
		// VisibilityStatus: types.VisibilityStatus, // Required
	}

	if len(_qconnectAiGuardrailId) > 0 {
		input.AiGuardrailId = aws.String(_qconnectAiGuardrailId)
	}
	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectBlockedInputMessaging) > 0 {
		input.BlockedInputMessaging = aws.String(_qconnectBlockedInputMessaging)
	}
	if len(_qconnectBlockedOutputsMessaging) > 0 {
		input.BlockedOutputsMessaging = aws.String(_qconnectBlockedOutputsMessaging)
	}
	if len(_qconnectVisibilityStatus) > 0 {
		if err := assignInputField(input, "VisibilityStatus", _qconnectVisibilityStatus); err != nil {
			log.Errorf("invalid --visibility-status: %s", err.Error())
			return
		}
	}
	if len(_qconnectClientToken) > 0 {
		input.ClientToken = aws.String(_qconnectClientToken)
	}
	if len(_qconnectContentPolicyConfig) > 0 {
		if err := assignInputField(input, "ContentPolicyConfig", _qconnectContentPolicyConfig); err != nil {
			log.Errorf("invalid --content-policy-config: %s", err.Error())
			return
		}
	}
	if len(_qconnectContextualGroundingPolicyConfig) > 0 {
		if err := assignInputField(input, "ContextualGroundingPolicyConfig", _qconnectContextualGroundingPolicyConfig); err != nil {
			log.Errorf("invalid --contextual-grounding-policy-config: %s", err.Error())
			return
		}
	}
	if len(_qconnectDescription) > 0 {
		input.Description = aws.String(_qconnectDescription)
	}
	if len(_qconnectSensitiveInformationPolicyConfig) > 0 {
		if err := assignInputField(input, "SensitiveInformationPolicyConfig", _qconnectSensitiveInformationPolicyConfig); err != nil {
			log.Errorf("invalid --sensitive-information-policy-config: %s", err.Error())
			return
		}
	}
	if len(_qconnectTopicPolicyConfig) > 0 {
		if err := assignInputField(input, "TopicPolicyConfig", _qconnectTopicPolicyConfig); err != nil {
			log.Errorf("invalid --topic-policy-config: %s", err.Error())
			return
		}
	}
	if len(_qconnectWordPolicyConfig) > 0 {
		if err := assignInputField(input, "WordPolicyConfig", _qconnectWordPolicyConfig); err != nil {
			log.Errorf("invalid --word-policy-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAIGuardrail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an AI Prompt.
func qconnect_UpdateAIPrompt(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.UpdateAIPromptInput{
		// AiPromptId: *string, // Required
		// AssistantId: *string, // Required
		// VisibilityStatus: types.VisibilityStatus, // Required
	}

	if len(_qconnectAiPromptId) > 0 {
		input.AiPromptId = aws.String(_qconnectAiPromptId)
	}
	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectVisibilityStatus) > 0 {
		if err := assignInputField(input, "VisibilityStatus", _qconnectVisibilityStatus); err != nil {
			log.Errorf("invalid --visibility-status: %s", err.Error())
			return
		}
	}
	if len(_qconnectClientToken) > 0 {
		input.ClientToken = aws.String(_qconnectClientToken)
	}
	if len(_qconnectDescription) > 0 {
		input.Description = aws.String(_qconnectDescription)
	}
	if len(_qconnectInferenceConfiguration) > 0 {
		if err := assignInputField(input, "InferenceConfiguration", _qconnectInferenceConfiguration); err != nil {
			log.Errorf("invalid --inference-configuration: %s", err.Error())
			return
		}
	}
	if len(_qconnectModelId) > 0 {
		input.ModelId = aws.String(_qconnectModelId)
	}
	if len(_qconnectTemplateConfiguration) > 0 {
		if err := assignInputField(input, "TemplateConfiguration", _qconnectTemplateConfiguration); err != nil {
			log.Errorf("invalid --template-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAIPrompt(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the AI Agent that is set for use by default on an Amazon Q in Connect
// Assistant.
func qconnect_UpdateAssistantAIAgent(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.UpdateAssistantAIAgentInput{
		// AiAgentType: types.AIAgentType, // Required
		// AssistantId: *string, // Required
		// Configuration: *types.AIAgentConfigurationData, // Required
	}

	if len(_qconnectAiAgentType) > 0 {
		if err := assignInputField(input, "AiAgentType", _qconnectAiAgentType); err != nil {
			log.Errorf("invalid --ai-agent-type: %s", err.Error())
			return
		}
	}
	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _qconnectConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_qconnectOrchestratorUseCase) > 0 {
		input.OrchestratorUseCase = aws.String(_qconnectOrchestratorUseCase)
	}

	if resp, err := client.UpdateAssistantAIAgent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates information about the content.
func qconnect_UpdateContent(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.UpdateContentInput{
		// ContentId: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_qconnectContentId) > 0 {
		input.ContentId = aws.String(_qconnectContentId)
	}
	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectMetadata) > 0 {
		if err := assignInputField(input, "Metadata", _qconnectMetadata); err != nil {
			log.Errorf("invalid --metadata: %s", err.Error())
			return
		}
	}
	if len(_qconnectOverrideLinkOutUri) > 0 {
		input.OverrideLinkOutUri = aws.String(_qconnectOverrideLinkOutUri)
	}
	if len(_qconnectRemoveOverrideLinkOutUri) > 0 {
		if err := assignInputField(input, "RemoveOverrideLinkOutUri", _qconnectRemoveOverrideLinkOutUri); err != nil {
			log.Errorf("invalid --remove-override-link-out-uri: %s", err.Error())
			return
		}
	}
	if len(_qconnectRevisionId) > 0 {
		input.RevisionId = aws.String(_qconnectRevisionId)
	}
	if len(_qconnectTitle) > 0 {
		input.Title = aws.String(_qconnectTitle)
	}
	if len(_qconnectUploadId) > 0 {
		input.UploadId = aws.String(_qconnectUploadId)
	}

	if resp, err := client.UpdateContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the template URI of a knowledge base. This is only supported for
// knowledge bases of type EXTERNAL. Include a single variable in ${variable}
// format; this interpolated by Amazon Q in Connect using ingested content. For
// example, if you ingest a Salesforce article, it has an Id value, and you can
// set the template URI to
// https://myInstanceName.lightning.force.com/lightning/r/Knowledge__kav/*${Id}*/view
// .
func qconnect_UpdateKnowledgeBaseTemplateUri(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.UpdateKnowledgeBaseTemplateUriInput{
		// KnowledgeBaseId: *string, // Required
		// TemplateUri: *string, // Required
	}

	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectTemplateUri) > 0 {
		input.TemplateUri = aws.String(_qconnectTemplateUri)
	}

	if resp, err := client.UpdateKnowledgeBaseTemplateUri(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the Amazon Q in Connect message template. Partial update is supported.
// If any field is not supplied, it will remain unchanged for the message template
// that is referenced by the $LATEST qualifier. Any modification will only apply
// to the message template that is referenced by the $LATEST qualifier. The fields
// for all available versions will remain unchanged.
func qconnect_UpdateMessageTemplate(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.UpdateMessageTemplateInput{
		// KnowledgeBaseId: *string, // Required
		// MessageTemplateId: *string, // Required
	}

	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectMessageTemplateId) > 0 {
		input.MessageTemplateId = aws.String(_qconnectMessageTemplateId)
	}
	if len(_qconnectContent) > 0 {
		if err := assignInputField(input, "Content", _qconnectContent); err != nil {
			log.Errorf("invalid --content: %s", err.Error())
			return
		}
	}
	if len(_qconnectDefaultAttributes) > 0 {
		if err := assignInputField(input, "DefaultAttributes", _qconnectDefaultAttributes); err != nil {
			log.Errorf("invalid --default-attributes: %s", err.Error())
			return
		}
	}
	if len(_qconnectLanguage) > 0 {
		input.Language = aws.String(_qconnectLanguage)
	}
	if len(_qconnectSourceConfiguration) > 0 {
		if err := assignInputField(input, "SourceConfiguration", _qconnectSourceConfiguration); err != nil {
			log.Errorf("invalid --source-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMessageTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the Amazon Q in Connect message template metadata. Note that any
// modification to the message template’s name, description and grouping
// configuration will applied to the message template pointed by the $LATEST
// qualifier and all available versions. Partial update is supported. If any field
// is not supplied, it will remain unchanged for the message template.
func qconnect_UpdateMessageTemplateMetadata(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.UpdateMessageTemplateMetadataInput{
		// KnowledgeBaseId: *string, // Required
		// MessageTemplateId: *string, // Required
	}

	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectMessageTemplateId) > 0 {
		input.MessageTemplateId = aws.String(_qconnectMessageTemplateId)
	}
	if len(_qconnectDescription) > 0 {
		input.Description = aws.String(_qconnectDescription)
	}
	if len(_qconnectGroupingConfiguration) > 0 {
		if err := assignInputField(input, "GroupingConfiguration", _qconnectGroupingConfiguration); err != nil {
			log.Errorf("invalid --grouping-configuration: %s", err.Error())
			return
		}
	}
	if len(_qconnectName) > 0 {
		input.Name = aws.String(_qconnectName)
	}

	if resp, err := client.UpdateMessageTemplateMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing Amazon Q in Connect quick response.
func qconnect_UpdateQuickResponse(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.UpdateQuickResponseInput{
		// KnowledgeBaseId: *string, // Required
		// QuickResponseId: *string, // Required
	}

	if len(_qconnectKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_qconnectKnowledgeBaseId)
	}
	if len(_qconnectQuickResponseId) > 0 {
		input.QuickResponseId = aws.String(_qconnectQuickResponseId)
	}
	if len(_qconnectChannels) > 0 {
		input.Channels = append([]string(nil), _qconnectChannels...)
	}
	if len(_qconnectContent) > 0 {
		if err := assignInputField(input, "Content", _qconnectContent); err != nil {
			log.Errorf("invalid --content: %s", err.Error())
			return
		}
	}
	if len(_qconnectContentType) > 0 {
		input.ContentType = aws.String(_qconnectContentType)
	}
	if len(_qconnectDescription) > 0 {
		input.Description = aws.String(_qconnectDescription)
	}
	if len(_qconnectGroupingConfiguration) > 0 {
		if err := assignInputField(input, "GroupingConfiguration", _qconnectGroupingConfiguration); err != nil {
			log.Errorf("invalid --grouping-configuration: %s", err.Error())
			return
		}
	}
	if len(_qconnectIsActive) > 0 {
		if err := assignInputField(input, "IsActive", _qconnectIsActive); err != nil {
			log.Errorf("invalid --is-active: %s", err.Error())
			return
		}
	}
	if len(_qconnectLanguage) > 0 {
		input.Language = aws.String(_qconnectLanguage)
	}
	if len(_qconnectName) > 0 {
		input.Name = aws.String(_qconnectName)
	}
	if len(_qconnectRemoveDescription) > 0 {
		if err := assignInputField(input, "RemoveDescription", _qconnectRemoveDescription); err != nil {
			log.Errorf("invalid --remove-description: %s", err.Error())
			return
		}
	}
	if len(_qconnectRemoveGroupingConfiguration) > 0 {
		if err := assignInputField(input, "RemoveGroupingConfiguration", _qconnectRemoveGroupingConfiguration); err != nil {
			log.Errorf("invalid --remove-grouping-configuration: %s", err.Error())
			return
		}
	}
	if len(_qconnectRemoveShortcutKey) > 0 {
		if err := assignInputField(input, "RemoveShortcutKey", _qconnectRemoveShortcutKey); err != nil {
			log.Errorf("invalid --remove-shortcut-key: %s", err.Error())
			return
		}
	}
	if len(_qconnectShortcutKey) > 0 {
		input.ShortcutKey = aws.String(_qconnectShortcutKey)
	}

	if resp, err := client.UpdateQuickResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a session. A session is a contextual container used for generating
// recommendations. Amazon Connect updates the existing Amazon Q in Connect session
// for each contact on which Amazon Q in Connect is enabled.
func qconnect_UpdateSession(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.UpdateSessionInput{
		// AssistantId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectSessionId) > 0 {
		input.SessionId = aws.String(_qconnectSessionId)
	}
	if len(_qconnectAiAgentConfiguration) > 0 {
		if err := assignInputField(input, "AiAgentConfiguration", _qconnectAiAgentConfiguration); err != nil {
			log.Errorf("invalid --ai-agent-configuration: %s", err.Error())
			return
		}
	}
	if len(_qconnectDescription) > 0 {
		input.Description = aws.String(_qconnectDescription)
	}
	if len(_qconnectOrchestratorConfigurationList) > 0 {
		if err := assignInputField(input, "OrchestratorConfigurationList", _qconnectOrchestratorConfigurationList); err != nil {
			log.Errorf("invalid --orchestrator-configuration-list: %s", err.Error())
			return
		}
	}
	if len(_qconnectRemoveOrchestratorConfigurationList) > 0 {
		if err := assignInputField(input, "RemoveOrchestratorConfigurationList", _qconnectRemoveOrchestratorConfigurationList); err != nil {
			log.Errorf("invalid --remove-orchestrator-configuration-list: %s", err.Error())
			return
		}
	}
	if len(_qconnectTagFilter) > 0 {
		if err := assignInputField(input, "TagFilter", _qconnectTagFilter); err != nil {
			log.Errorf("invalid --tag-filter: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the data stored on an Amazon Q in Connect Session.
func qconnect_UpdateSessionData(cfg aws.Config, client *qconnect.Client) {
	input := &qconnect.UpdateSessionDataInput{
		// AssistantId: *string, // Required
		// Data: []types.RuntimeSessionData, // Required
		// SessionId: *string, // Required
	}

	if len(_qconnectAssistantId) > 0 {
		input.AssistantId = aws.String(_qconnectAssistantId)
	}
	if len(_qconnectData) > 0 {
		if err := assignInputField(input, "Data", _qconnectData); err != nil {
			log.Errorf("invalid --data: %s", err.Error())
			return
		}
	}
	if len(_qconnectSessionId) > 0 {
		input.SessionId = aws.String(_qconnectSessionId)
	}
	if len(_qconnectNamespace) > 0 {
		if err := assignInputField(input, "Namespace", _qconnectNamespace); err != nil {
			log.Errorf("invalid --namespace: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSessionData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_qconnectCmd)
	_qconnectCmd.Flags().SortFlags = false

	_qconnectCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_qconnectCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_qconnectCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_qconnectCmd.Flags().StringVarP(&_qconnectAiAgentConfiguration, "ai-agent-configuration", "", "", "Ai Agent Configuration")
	_qconnectCmd.Flags().StringVarP(&_qconnectAiAgentId, "ai-agent-id", "", "", "Ai Agent ID")
	_qconnectCmd.Flags().StringVarP(&_qconnectAiAgentType, "ai-agent-type", "", "", "Ai Agent Type")
	_qconnectCmd.Flags().StringVarP(&_qconnectAiGuardrailId, "ai-guardrail-id", "", "", "Ai Guardrail ID")
	_qconnectCmd.Flags().StringVarP(&_qconnectAiPromptId, "ai-prompt-id", "", "", "Ai Prompt ID")
	_qconnectCmd.Flags().StringVarP(&_qconnectApiFormat, "api-format", "", "", "API Format")
	_qconnectCmd.Flags().StringVarP(&_qconnectAssistantAssociationId, "assistant-association-id", "", "", "Assistant Association ID")
	_qconnectCmd.Flags().StringVarP(&_qconnectAssistantId, "assistant-id", "", "", "Assistant ID")
	_qconnectCmd.Flags().StringVarP(&_qconnectAssociation, "association", "", "", "Association")
	_qconnectCmd.Flags().StringVarP(&_qconnectAssociationType, "association-type", "", "", "Association Type")
	_qconnectCmd.Flags().StringVarP(&_qconnectAttachmentId, "attachment-id", "", "", "Attachment ID")
	_qconnectCmd.Flags().StringVarP(&_qconnectAttributes, "attributes", "", "", "Attributes")
	_qconnectCmd.Flags().StringVarP(&_qconnectBlockedInputMessaging, "blocked-input-messaging", "", "", "Blocked Input Messaging")
	_qconnectCmd.Flags().StringVarP(&_qconnectBlockedOutputsMessaging, "blocked-outputs-messaging", "", "", "Blocked Outputs Messaging")
	_qconnectCmd.Flags().StringVarP(&_qconnectBody, "body", "", "", "Body")
	_qconnectCmd.Flags().StringVarP(&_qconnectChannelSubtype, "channel-subtype", "", "", "Channel Subtype")
	_qconnectCmd.Flags().StringSliceVarP(&_qconnectChannels, "channels", "", nil, "Channels")
	_qconnectCmd.Flags().StringVarP(&_qconnectClientToken, "client-token", "", "", "Client Token")
	_qconnectCmd.Flags().StringVarP(&_qconnectConfiguration, "configuration", "", "", "Configuration")
	_qconnectCmd.Flags().StringVarP(&_qconnectContactArn, "contact-arn", "", "", "Contact ARN")
	_qconnectCmd.Flags().StringVarP(&_qconnectContent, "content", "", "", "Content")
	_qconnectCmd.Flags().StringVarP(&_qconnectContentAssociationId, "content-association-id", "", "", "Content Association ID")
	_qconnectCmd.Flags().StringVarP(&_qconnectContentDisposition, "content-disposition", "", "", "Content Disposition")
	_qconnectCmd.Flags().StringVarP(&_qconnectContentFeedback, "content-feedback", "", "", "Content Feedback")
	_qconnectCmd.Flags().StringVarP(&_qconnectContentId, "content-id", "", "", "Content ID")
	_qconnectCmd.Flags().StringVarP(&_qconnectContentPolicyConfig, "content-policy-config", "", "", "Content Policy Config")
	_qconnectCmd.Flags().StringVarP(&_qconnectContentType, "content-type", "", "", "Content Type")
	_qconnectCmd.Flags().StringVarP(&_qconnectContextualGroundingPolicyConfig, "contextual-grounding-policy-config", "", "", "Contextual Grounding Policy Config")
	_qconnectCmd.Flags().StringVarP(&_qconnectConversationContext, "conversation-context", "", "", "Conversation Context")
	_qconnectCmd.Flags().StringVarP(&_qconnectData, "data", "", "", "Data")
	_qconnectCmd.Flags().StringVarP(&_qconnectDefaultAttributes, "default-attributes", "", "", "Default Attributes")
	_qconnectCmd.Flags().StringVarP(&_qconnectDescription, "description", "", "", "Description")
	_qconnectCmd.Flags().StringVarP(&_qconnectExternalSourceConfiguration, "external-source-configuration", "", "", "External Source Configuration")
	_qconnectCmd.Flags().StringVarP(&_qconnectFilter, "filter", "", "", "Filter")
	_qconnectCmd.Flags().StringVarP(&_qconnectGroupingConfiguration, "grouping-configuration", "", "", "Grouping Configuration")
	_qconnectCmd.Flags().StringVarP(&_qconnectImportJobId, "import-job-id", "", "", "Import Job ID")
	_qconnectCmd.Flags().StringVarP(&_qconnectImportJobType, "import-job-type", "", "", "Import Job Type")
	_qconnectCmd.Flags().StringVarP(&_qconnectInferenceConfiguration, "inference-configuration", "", "", "Inference Configuration")
	_qconnectCmd.Flags().StringVarP(&_qconnectIsActive, "is-active", "", "", "Is Active")
	_qconnectCmd.Flags().StringVarP(&_qconnectKnowledgeBaseId, "knowledge-base-id", "", "", "Knowledge Base ID")
	_qconnectCmd.Flags().StringVarP(&_qconnectKnowledgeBaseType, "knowledge-base-type", "", "", "Knowledge Base Type")
	_qconnectCmd.Flags().StringVarP(&_qconnectLanguage, "language", "", "", "Language")
	_qconnectCmd.Flags().StringVarP(&_qconnectMaxResults, "max-results", "", "", "Max Results")
	_qconnectCmd.Flags().StringVarP(&_qconnectMessage, "message", "", "", "Message")
	_qconnectCmd.Flags().StringVarP(&_qconnectMessageTemplateContentSha256, "message-template-content-sha256", "", "", "Message Template Content SHA256")
	_qconnectCmd.Flags().StringVarP(&_qconnectMessageTemplateId, "message-template-id", "", "", "Message Template ID")
	_qconnectCmd.Flags().StringVarP(&_qconnectMetadata, "metadata", "", "", "Metadata")
	_qconnectCmd.Flags().StringVarP(&_qconnectModelId, "model-id", "", "", "Model ID")
	_qconnectCmd.Flags().StringVarP(&_qconnectModifiedTime, "modified-time", "", "", "Modified Time")
	_qconnectCmd.Flags().StringVarP(&_qconnectName, "name", "", "", "Name")
	_qconnectCmd.Flags().StringVarP(&_qconnectNamespace, "namespace", "", "", "Namespace")
	_qconnectCmd.Flags().StringVarP(&_qconnectNextChunkToken, "next-chunk-token", "", "", "Next Chunk Token")
	_qconnectCmd.Flags().StringVarP(&_qconnectNextMessageToken, "next-message-token", "", "", "Next Message Token")
	_qconnectCmd.Flags().StringVarP(&_qconnectNextToken, "next-token", "", "", "Next Token")
	_qconnectCmd.Flags().StringVarP(&_qconnectOrchestratorConfigurationList, "orchestrator-configuration-list", "", "", "Orchestrator Configuration List")
	_qconnectCmd.Flags().StringVarP(&_qconnectOrchestratorUseCase, "orchestrator-use-case", "", "", "Orchestrator Use Case")
	_qconnectCmd.Flags().StringVarP(&_qconnectOrigin, "origin", "", "", "Origin")
	_qconnectCmd.Flags().StringVarP(&_qconnectOverrideKnowledgeBaseSearchType, "override-knowledge-base-search-type", "", "", "Override Knowledge Base Search Type")
	_qconnectCmd.Flags().StringVarP(&_qconnectOverrideLinkOutUri, "override-link-out-uri", "", "", "Override Link Out URI")
	_qconnectCmd.Flags().StringVarP(&_qconnectPresignedUrlTimeToLive, "presigned-url-time-to-live", "", "", "Presigned URL Time To Live")
	_qconnectCmd.Flags().StringVarP(&_qconnectQueryCondition, "query-condition", "", "", "Query Condition")
	_qconnectCmd.Flags().StringVarP(&_qconnectQueryInputData, "query-input-data", "", "", "Query Input Data")
	_qconnectCmd.Flags().StringVarP(&_qconnectQueryText, "query-text", "", "", "Query Text")
	_qconnectCmd.Flags().StringVarP(&_qconnectQuickResponseId, "quick-response-id", "", "", "Quick Response ID")
	_qconnectCmd.Flags().StringSliceVarP(&_qconnectRecommendationIds, "recommendation-ids", "", nil, "Recommendation Ids")
	_qconnectCmd.Flags().StringVarP(&_qconnectRecommendationType, "recommendation-type", "", "", "Recommendation Type")
	_qconnectCmd.Flags().StringVarP(&_qconnectRemoveDescription, "remove-description", "", "", "Remove Description")
	_qconnectCmd.Flags().StringVarP(&_qconnectRemoveGroupingConfiguration, "remove-grouping-configuration", "", "", "Remove Grouping Configuration")
	_qconnectCmd.Flags().StringVarP(&_qconnectRemoveOrchestratorConfigurationList, "remove-orchestrator-configuration-list", "", "", "Remove Orchestrator Configuration List")
	_qconnectCmd.Flags().StringVarP(&_qconnectRemoveOverrideLinkOutUri, "remove-override-link-out-uri", "", "", "Remove Override Link Out URI")
	_qconnectCmd.Flags().StringVarP(&_qconnectRemoveShortcutKey, "remove-shortcut-key", "", "", "Remove Shortcut Key")
	_qconnectCmd.Flags().StringVarP(&_qconnectRenderingConfiguration, "rendering-configuration", "", "", "Rendering Configuration")
	_qconnectCmd.Flags().StringVarP(&_qconnectResourceArn, "resource-arn", "", "", "Resource ARN")
	_qconnectCmd.Flags().StringVarP(&_qconnectRetrievalConfiguration, "retrieval-configuration", "", "", "Retrieval Configuration")
	_qconnectCmd.Flags().StringVarP(&_qconnectRetrievalQuery, "retrieval-query", "", "", "Retrieval Query")
	_qconnectCmd.Flags().StringVarP(&_qconnectRevisionId, "revision-id", "", "", "Revision ID")
	_qconnectCmd.Flags().StringVarP(&_qconnectSearchExpression, "search-expression", "", "", "Search Expression")
	_qconnectCmd.Flags().StringVarP(&_qconnectSensitiveInformationPolicyConfig, "sensitive-information-policy-config", "", "", "Sensitive Information Policy Config")
	_qconnectCmd.Flags().StringVarP(&_qconnectServerSideEncryptionConfiguration, "server-side-encryption-configuration", "", "", "Server Side Encryption Configuration")
	_qconnectCmd.Flags().StringVarP(&_qconnectSessionId, "session-id", "", "", "Session ID")
	_qconnectCmd.Flags().StringVarP(&_qconnectShortcutKey, "shortcut-key", "", "", "Shortcut Key")
	_qconnectCmd.Flags().StringVarP(&_qconnectSourceConfiguration, "source-configuration", "", "", "Source Configuration")
	_qconnectCmd.Flags().StringVarP(&_qconnectTagFilter, "tag-filter", "", "", "Tag Filter")
	_qconnectCmd.Flags().StringSliceVarP(&_qconnectTagKeys, "tag-keys", "", nil, "Tag Keys")
	_qconnectCmd.Flags().StringVarP(&_qconnectTags, "tags", "", "", "Tags")
	_qconnectCmd.Flags().StringVarP(&_qconnectTargetId, "target-id", "", "", "Target ID")
	_qconnectCmd.Flags().StringVarP(&_qconnectTargetType, "target-type", "", "", "Target Type")
	_qconnectCmd.Flags().StringVarP(&_qconnectTemplateConfiguration, "template-configuration", "", "", "Template Configuration")
	_qconnectCmd.Flags().StringVarP(&_qconnectTemplateType, "template-type", "", "", "Template Type")
	_qconnectCmd.Flags().StringVarP(&_qconnectTemplateUri, "template-uri", "", "", "Template URI")
	_qconnectCmd.Flags().StringVarP(&_qconnectTitle, "title", "", "", "Title")
	_qconnectCmd.Flags().StringVarP(&_qconnectTopicPolicyConfig, "topic-policy-config", "", "", "Topic Policy Config")
	_qconnectCmd.Flags().StringVarP(&_qconnectType, "type", "", "", "Type")
	_qconnectCmd.Flags().StringVarP(&_qconnectUploadId, "upload-id", "", "", "Upload ID")
	_qconnectCmd.Flags().StringVarP(&_qconnectVectorIngestionConfiguration, "vector-ingestion-configuration", "", "", "Vector Ingestion Configuration")
	_qconnectCmd.Flags().StringVarP(&_qconnectVersionNumber, "version-number", "", "", "Version Number")
	_qconnectCmd.Flags().StringVarP(&_qconnectVisibilityStatus, "visibility-status", "", "", "Visibility Status")
	_qconnectCmd.Flags().StringVarP(&_qconnectWaitTimeSeconds, "wait-time-seconds", "", "", "Wait Time Seconds")
	_qconnectCmd.Flags().StringVarP(&_qconnectWordPolicyConfig, "word-policy-config", "", "", "Word Policy Config")

	_qconnectCmd.Flags().BoolVarP(&_qconnectActivateMessageTemplate, "activate-message-template", "", false, "Activate Message Template")
	_qconnectCmd.Flags().BoolVarP(&_qconnectCreateAIAgent, "create-ai-agent", "", false, "Create Ai Agent")
	_qconnectCmd.Flags().BoolVarP(&_qconnectCreateAIAgentVersion, "create-ai-agent-version", "", false, "Create Ai Agent Version")
	_qconnectCmd.Flags().BoolVarP(&_qconnectCreateAIGuardrail, "create-ai-guardrail", "", false, "Create Ai Guardrail")
	_qconnectCmd.Flags().BoolVarP(&_qconnectCreateAIGuardrailVersion, "create-ai-guardrail-version", "", false, "Create Ai Guardrail Version")
	_qconnectCmd.Flags().BoolVarP(&_qconnectCreateAIPrompt, "create-ai-prompt", "", false, "Create Ai Prompt")
	_qconnectCmd.Flags().BoolVarP(&_qconnectCreateAIPromptVersion, "create-ai-prompt-version", "", false, "Create Ai Prompt Version")
	_qconnectCmd.Flags().BoolVarP(&_qconnectCreateAssistant, "create-assistant", "", false, "Create Assistant")
	_qconnectCmd.Flags().BoolVarP(&_qconnectCreateAssistantAssociation, "create-assistant-association", "", false, "Create Assistant Association")
	_qconnectCmd.Flags().BoolVarP(&_qconnectCreateContent, "create-content", "", false, "Create Content")
	_qconnectCmd.Flags().BoolVarP(&_qconnectCreateContentAssociation, "create-content-association", "", false, "Create Content Association")
	_qconnectCmd.Flags().BoolVarP(&_qconnectCreateKnowledgeBase, "create-knowledge-base", "", false, "Create Knowledge Base")
	_qconnectCmd.Flags().BoolVarP(&_qconnectCreateMessageTemplate, "create-message-template", "", false, "Create Message Template")
	_qconnectCmd.Flags().BoolVarP(&_qconnectCreateMessageTemplateAttachment, "create-message-template-attachment", "", false, "Create Message Template Attachment")
	_qconnectCmd.Flags().BoolVarP(&_qconnectCreateMessageTemplateVersion, "create-message-template-version", "", false, "Create Message Template Version")
	_qconnectCmd.Flags().BoolVarP(&_qconnectCreateQuickResponse, "create-quick-response", "", false, "Create Quick Response")
	_qconnectCmd.Flags().BoolVarP(&_qconnectCreateSession, "create-session", "", false, "Create Session")
	_qconnectCmd.Flags().BoolVarP(&_qconnectDeactivateMessageTemplate, "deactivate-message-template", "", false, "Deactivate Message Template")
	_qconnectCmd.Flags().BoolVarP(&_qconnectDeleteAIAgent, "delete-ai-agent", "", false, "Delete Ai Agent")
	_qconnectCmd.Flags().BoolVarP(&_qconnectDeleteAIAgentVersion, "delete-ai-agent-version", "", false, "Delete Ai Agent Version")
	_qconnectCmd.Flags().BoolVarP(&_qconnectDeleteAIGuardrail, "delete-ai-guardrail", "", false, "Delete Ai Guardrail")
	_qconnectCmd.Flags().BoolVarP(&_qconnectDeleteAIGuardrailVersion, "delete-ai-guardrail-version", "", false, "Delete Ai Guardrail Version")
	_qconnectCmd.Flags().BoolVarP(&_qconnectDeleteAIPrompt, "delete-ai-prompt", "", false, "Delete Ai Prompt")
	_qconnectCmd.Flags().BoolVarP(&_qconnectDeleteAIPromptVersion, "delete-ai-prompt-version", "", false, "Delete Ai Prompt Version")
	_qconnectCmd.Flags().BoolVarP(&_qconnectDeleteAssistant, "delete-assistant", "", false, "Delete Assistant")
	_qconnectCmd.Flags().BoolVarP(&_qconnectDeleteAssistantAssociation, "delete-assistant-association", "", false, "Delete Assistant Association")
	_qconnectCmd.Flags().BoolVarP(&_qconnectDeleteContent, "delete-content", "", false, "Delete Content")
	_qconnectCmd.Flags().BoolVarP(&_qconnectDeleteContentAssociation, "delete-content-association", "", false, "Delete Content Association")
	_qconnectCmd.Flags().BoolVarP(&_qconnectDeleteImportJob, "delete-import-job", "", false, "Delete Import Job")
	_qconnectCmd.Flags().BoolVarP(&_qconnectDeleteKnowledgeBase, "delete-knowledge-base", "", false, "Delete Knowledge Base")
	_qconnectCmd.Flags().BoolVarP(&_qconnectDeleteMessageTemplate, "delete-message-template", "", false, "Delete Message Template")
	_qconnectCmd.Flags().BoolVarP(&_qconnectDeleteMessageTemplateAttachment, "delete-message-template-attachment", "", false, "Delete Message Template Attachment")
	_qconnectCmd.Flags().BoolVarP(&_qconnectDeleteQuickResponse, "delete-quick-response", "", false, "Delete Quick Response")
	_qconnectCmd.Flags().BoolVarP(&_qconnectGetAIAgent, "get-ai-agent", "", false, "Get Ai Agent")
	_qconnectCmd.Flags().BoolVarP(&_qconnectGetAIGuardrail, "get-ai-guardrail", "", false, "Get Ai Guardrail")
	_qconnectCmd.Flags().BoolVarP(&_qconnectGetAIPrompt, "get-ai-prompt", "", false, "Get Ai Prompt")
	_qconnectCmd.Flags().BoolVarP(&_qconnectGetAssistant, "get-assistant", "", false, "Get Assistant")
	_qconnectCmd.Flags().BoolVarP(&_qconnectGetAssistantAssociation, "get-assistant-association", "", false, "Get Assistant Association")
	_qconnectCmd.Flags().BoolVarP(&_qconnectGetContent, "get-content", "", false, "Get Content")
	_qconnectCmd.Flags().BoolVarP(&_qconnectGetContentAssociation, "get-content-association", "", false, "Get Content Association")
	_qconnectCmd.Flags().BoolVarP(&_qconnectGetContentSummary, "get-content-summary", "", false, "Get Content Summary")
	_qconnectCmd.Flags().BoolVarP(&_qconnectGetImportJob, "get-import-job", "", false, "Get Import Job")
	_qconnectCmd.Flags().BoolVarP(&_qconnectGetKnowledgeBase, "get-knowledge-base", "", false, "Get Knowledge Base")
	_qconnectCmd.Flags().BoolVarP(&_qconnectGetMessageTemplate, "get-message-template", "", false, "Get Message Template")
	_qconnectCmd.Flags().BoolVarP(&_qconnectGetNextMessage, "get-next-message", "", false, "Get Next Message")
	_qconnectCmd.Flags().BoolVarP(&_qconnectGetQuickResponse, "get-quick-response", "", false, "Get Quick Response")
	_qconnectCmd.Flags().BoolVarP(&_qconnectGetRecommendations, "get-recommendations", "", false, "Get Recommendations")
	_qconnectCmd.Flags().BoolVarP(&_qconnectGetSession, "get-session", "", false, "Get Session")
	_qconnectCmd.Flags().BoolVarP(&_qconnectListAIAgentVersions, "list-ai-agent-versions", "", false, "List Ai Agent Versions")
	_qconnectCmd.Flags().BoolVarP(&_qconnectListAIAgents, "list-ai-agents", "", false, "List Ai Agents")
	_qconnectCmd.Flags().BoolVarP(&_qconnectListAIGuardrailVersions, "list-ai-guardrail-versions", "", false, "List Ai Guardrail Versions")
	_qconnectCmd.Flags().BoolVarP(&_qconnectListAIGuardrails, "list-ai-guardrails", "", false, "List Ai Guardrails")
	_qconnectCmd.Flags().BoolVarP(&_qconnectListAIPromptVersions, "list-ai-prompt-versions", "", false, "List Ai Prompt Versions")
	_qconnectCmd.Flags().BoolVarP(&_qconnectListAIPrompts, "list-ai-prompts", "", false, "List Ai Prompts")
	_qconnectCmd.Flags().BoolVarP(&_qconnectListAssistantAssociations, "list-assistant-associations", "", false, "List Assistant Associations")
	_qconnectCmd.Flags().BoolVarP(&_qconnectListAssistants, "list-assistants", "", false, "List Assistants")
	_qconnectCmd.Flags().BoolVarP(&_qconnectListContentAssociations, "list-content-associations", "", false, "List Content Associations")
	_qconnectCmd.Flags().BoolVarP(&_qconnectListContents, "list-contents", "", false, "List Contents")
	_qconnectCmd.Flags().BoolVarP(&_qconnectListImportJobs, "list-import-jobs", "", false, "List Import Jobs")
	_qconnectCmd.Flags().BoolVarP(&_qconnectListKnowledgeBases, "list-knowledge-bases", "", false, "List Knowledge Bases")
	_qconnectCmd.Flags().BoolVarP(&_qconnectListMessageTemplateVersions, "list-message-template-versions", "", false, "List Message Template Versions")
	_qconnectCmd.Flags().BoolVarP(&_qconnectListMessageTemplates, "list-message-templates", "", false, "List Message Templates")
	_qconnectCmd.Flags().BoolVarP(&_qconnectListMessages, "list-messages", "", false, "List Messages")
	_qconnectCmd.Flags().BoolVarP(&_qconnectListQuickResponses, "list-quick-responses", "", false, "List Quick Responses")
	_qconnectCmd.Flags().BoolVarP(&_qconnectListSpans, "list-spans", "", false, "List Spans")
	_qconnectCmd.Flags().BoolVarP(&_qconnectListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_qconnectCmd.Flags().BoolVarP(&_qconnectNotifyRecommendationsReceived, "notify-recommendations-received", "", false, "Notify Recommendations Received")
	_qconnectCmd.Flags().BoolVarP(&_qconnectPutFeedback, "put-feedback", "", false, "Put Feedback")
	_qconnectCmd.Flags().BoolVarP(&_qconnectQueryAssistant, "query-assistant", "", false, "Query Assistant")
	_qconnectCmd.Flags().BoolVarP(&_qconnectRemoveAssistantAIAgent, "remove-assistant-ai-agent", "", false, "Remove Assistant Ai Agent")
	_qconnectCmd.Flags().BoolVarP(&_qconnectRemoveKnowledgeBaseTemplateUri, "remove-knowledge-base-template-uri", "", false, "Remove Knowledge Base Template URI")
	_qconnectCmd.Flags().BoolVarP(&_qconnectRenderMessageTemplate, "render-message-template", "", false, "Render Message Template")
	_qconnectCmd.Flags().BoolVarP(&_qconnectRetrieve, "retrieve", "", false, "Retrieve")
	_qconnectCmd.Flags().BoolVarP(&_qconnectSearchContent, "search-content", "", false, "Search Content")
	_qconnectCmd.Flags().BoolVarP(&_qconnectSearchMessageTemplates, "search-message-templates", "", false, "Search Message Templates")
	_qconnectCmd.Flags().BoolVarP(&_qconnectSearchQuickResponses, "search-quick-responses", "", false, "Search Quick Responses")
	_qconnectCmd.Flags().BoolVarP(&_qconnectSearchSessions, "search-sessions", "", false, "Search Sessions")
	_qconnectCmd.Flags().BoolVarP(&_qconnectSendMessage, "send-message", "", false, "Send Message")
	_qconnectCmd.Flags().BoolVarP(&_qconnectStartContentUpload, "start-content-upload", "", false, "Start Content Upload")
	_qconnectCmd.Flags().BoolVarP(&_qconnectStartImportJob, "start-import-job", "", false, "Start Import Job")
	_qconnectCmd.Flags().BoolVarP(&_qconnectTagResource, "tag-resource", "", false, "Tag Resource")
	_qconnectCmd.Flags().BoolVarP(&_qconnectUntagResource, "untag-resource", "", false, "Untag Resource")
	_qconnectCmd.Flags().BoolVarP(&_qconnectUpdateAIAgent, "update-ai-agent", "", false, "Update Ai Agent")
	_qconnectCmd.Flags().BoolVarP(&_qconnectUpdateAIGuardrail, "update-ai-guardrail", "", false, "Update Ai Guardrail")
	_qconnectCmd.Flags().BoolVarP(&_qconnectUpdateAIPrompt, "update-ai-prompt", "", false, "Update Ai Prompt")
	_qconnectCmd.Flags().BoolVarP(&_qconnectUpdateAssistantAIAgent, "update-assistant-ai-agent", "", false, "Update Assistant Ai Agent")
	_qconnectCmd.Flags().BoolVarP(&_qconnectUpdateContent, "update-content", "", false, "Update Content")
	_qconnectCmd.Flags().BoolVarP(&_qconnectUpdateKnowledgeBaseTemplateUri, "update-knowledge-base-template-uri", "", false, "Update Knowledge Base Template URI")
	_qconnectCmd.Flags().BoolVarP(&_qconnectUpdateMessageTemplate, "update-message-template", "", false, "Update Message Template")
	_qconnectCmd.Flags().BoolVarP(&_qconnectUpdateMessageTemplateMetadata, "update-message-template-metadata", "", false, "Update Message Template Metadata")
	_qconnectCmd.Flags().BoolVarP(&_qconnectUpdateQuickResponse, "update-quick-response", "", false, "Update Quick Response")
	_qconnectCmd.Flags().BoolVarP(&_qconnectUpdateSession, "update-session", "", false, "Update Session")
	_qconnectCmd.Flags().BoolVarP(&_qconnectUpdateSessionData, "update-session-data", "", false, "Update Session Data")

}
