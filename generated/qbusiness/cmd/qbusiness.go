package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/qbusiness"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// qbusinessCmd represents the qbusiness command
var _qbusinessCmd = &cobra.Command{
	Use:   "qbusiness",
	Short: "AWS qbusiness CLI",
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
		client := qbusiness.NewFromConfig(cfg)
		if _qbusinessAssociatePermission {
			qbusiness_AssociatePermission(cfg, client)
			return
		}
		if _qbusinessBatchDeleteDocument {
			qbusiness_BatchDeleteDocument(cfg, client)
			return
		}
		if _qbusinessBatchPutDocument {
			qbusiness_BatchPutDocument(cfg, client)
			return
		}
		if _qbusinessCancelSubscription {
			qbusiness_CancelSubscription(cfg, client)
			return
		}
		if _qbusinessChat {
			qbusiness_Chat(cfg, client)
			return
		}
		if _qbusinessChatSync {
			qbusiness_ChatSync(cfg, client)
			return
		}
		if _qbusinessCheckDocumentAccess {
			qbusiness_CheckDocumentAccess(cfg, client)
			return
		}
		if _qbusinessCreateAnonymousWebExperienceUrl {
			qbusiness_CreateAnonymousWebExperienceUrl(cfg, client)
			return
		}
		if _qbusinessCreateApplication {
			qbusiness_CreateApplication(cfg, client)
			return
		}
		if _qbusinessCreateChatResponseConfiguration {
			qbusiness_CreateChatResponseConfiguration(cfg, client)
			return
		}
		if _qbusinessCreateDataAccessor {
			qbusiness_CreateDataAccessor(cfg, client)
			return
		}
		if _qbusinessCreateDataSource {
			qbusiness_CreateDataSource(cfg, client)
			return
		}
		if _qbusinessCreateIndex {
			qbusiness_CreateIndex(cfg, client)
			return
		}
		if _qbusinessCreatePlugin {
			qbusiness_CreatePlugin(cfg, client)
			return
		}
		if _qbusinessCreateRetriever {
			qbusiness_CreateRetriever(cfg, client)
			return
		}
		if _qbusinessCreateSubscription {
			qbusiness_CreateSubscription(cfg, client)
			return
		}
		if _qbusinessCreateUser {
			qbusiness_CreateUser(cfg, client)
			return
		}
		if _qbusinessCreateWebExperience {
			qbusiness_CreateWebExperience(cfg, client)
			return
		}
		if _qbusinessDeleteApplication {
			qbusiness_DeleteApplication(cfg, client)
			return
		}
		if _qbusinessDeleteAttachment {
			qbusiness_DeleteAttachment(cfg, client)
			return
		}
		if _qbusinessDeleteChatControlsConfiguration {
			qbusiness_DeleteChatControlsConfiguration(cfg, client)
			return
		}
		if _qbusinessDeleteChatResponseConfiguration {
			qbusiness_DeleteChatResponseConfiguration(cfg, client)
			return
		}
		if _qbusinessDeleteConversation {
			qbusiness_DeleteConversation(cfg, client)
			return
		}
		if _qbusinessDeleteDataAccessor {
			qbusiness_DeleteDataAccessor(cfg, client)
			return
		}
		if _qbusinessDeleteDataSource {
			qbusiness_DeleteDataSource(cfg, client)
			return
		}
		if _qbusinessDeleteGroup {
			qbusiness_DeleteGroup(cfg, client)
			return
		}
		if _qbusinessDeleteIndex {
			qbusiness_DeleteIndex(cfg, client)
			return
		}
		if _qbusinessDeletePlugin {
			qbusiness_DeletePlugin(cfg, client)
			return
		}
		if _qbusinessDeleteRetriever {
			qbusiness_DeleteRetriever(cfg, client)
			return
		}
		if _qbusinessDeleteUser {
			qbusiness_DeleteUser(cfg, client)
			return
		}
		if _qbusinessDeleteWebExperience {
			qbusiness_DeleteWebExperience(cfg, client)
			return
		}
		if _qbusinessDisassociatePermission {
			qbusiness_DisassociatePermission(cfg, client)
			return
		}
		if _qbusinessGetApplication {
			qbusiness_GetApplication(cfg, client)
			return
		}
		if _qbusinessGetChatControlsConfiguration {
			qbusiness_GetChatControlsConfiguration(cfg, client)
			return
		}
		if _qbusinessGetChatResponseConfiguration {
			qbusiness_GetChatResponseConfiguration(cfg, client)
			return
		}
		if _qbusinessGetDataAccessor {
			qbusiness_GetDataAccessor(cfg, client)
			return
		}
		if _qbusinessGetDataSource {
			qbusiness_GetDataSource(cfg, client)
			return
		}
		if _qbusinessGetDocumentContent {
			qbusiness_GetDocumentContent(cfg, client)
			return
		}
		if _qbusinessGetGroup {
			qbusiness_GetGroup(cfg, client)
			return
		}
		if _qbusinessGetIndex {
			qbusiness_GetIndex(cfg, client)
			return
		}
		if _qbusinessGetMedia {
			qbusiness_GetMedia(cfg, client)
			return
		}
		if _qbusinessGetPlugin {
			qbusiness_GetPlugin(cfg, client)
			return
		}
		if _qbusinessGetPolicy {
			qbusiness_GetPolicy(cfg, client)
			return
		}
		if _qbusinessGetRetriever {
			qbusiness_GetRetriever(cfg, client)
			return
		}
		if _qbusinessGetUser {
			qbusiness_GetUser(cfg, client)
			return
		}
		if _qbusinessGetWebExperience {
			qbusiness_GetWebExperience(cfg, client)
			return
		}
		if _qbusinessListApplications {
			qbusiness_ListApplications(cfg, client)
			return
		}
		if _qbusinessListAttachments {
			qbusiness_ListAttachments(cfg, client)
			return
		}
		if _qbusinessListChatResponseConfigurations {
			qbusiness_ListChatResponseConfigurations(cfg, client)
			return
		}
		if _qbusinessListConversations {
			qbusiness_ListConversations(cfg, client)
			return
		}
		if _qbusinessListDataAccessors {
			qbusiness_ListDataAccessors(cfg, client)
			return
		}
		if _qbusinessListDataSourceSyncJobs {
			qbusiness_ListDataSourceSyncJobs(cfg, client)
			return
		}
		if _qbusinessListDataSources {
			qbusiness_ListDataSources(cfg, client)
			return
		}
		if _qbusinessListDocuments {
			qbusiness_ListDocuments(cfg, client)
			return
		}
		if _qbusinessListGroups {
			qbusiness_ListGroups(cfg, client)
			return
		}
		if _qbusinessListIndices {
			qbusiness_ListIndices(cfg, client)
			return
		}
		if _qbusinessListMessages {
			qbusiness_ListMessages(cfg, client)
			return
		}
		if _qbusinessListPluginActions {
			qbusiness_ListPluginActions(cfg, client)
			return
		}
		if _qbusinessListPluginTypeActions {
			qbusiness_ListPluginTypeActions(cfg, client)
			return
		}
		if _qbusinessListPluginTypeMetadata {
			qbusiness_ListPluginTypeMetadata(cfg, client)
			return
		}
		if _qbusinessListPlugins {
			qbusiness_ListPlugins(cfg, client)
			return
		}
		if _qbusinessListRetrievers {
			qbusiness_ListRetrievers(cfg, client)
			return
		}
		if _qbusinessListSubscriptions {
			qbusiness_ListSubscriptions(cfg, client)
			return
		}
		if _qbusinessListTagsForResource {
			qbusiness_ListTagsForResource(cfg, client)
			return
		}
		if _qbusinessListWebExperiences {
			qbusiness_ListWebExperiences(cfg, client)
			return
		}
		if _qbusinessPutFeedback {
			qbusiness_PutFeedback(cfg, client)
			return
		}
		if _qbusinessPutGroup {
			qbusiness_PutGroup(cfg, client)
			return
		}
		if _qbusinessSearchRelevantContent {
			qbusiness_SearchRelevantContent(cfg, client)
			return
		}
		if _qbusinessStartDataSourceSyncJob {
			qbusiness_StartDataSourceSyncJob(cfg, client)
			return
		}
		if _qbusinessStopDataSourceSyncJob {
			qbusiness_StopDataSourceSyncJob(cfg, client)
			return
		}
		if _qbusinessTagResource {
			qbusiness_TagResource(cfg, client)
			return
		}
		if _qbusinessUntagResource {
			qbusiness_UntagResource(cfg, client)
			return
		}
		if _qbusinessUpdateApplication {
			qbusiness_UpdateApplication(cfg, client)
			return
		}
		if _qbusinessUpdateChatControlsConfiguration {
			qbusiness_UpdateChatControlsConfiguration(cfg, client)
			return
		}
		if _qbusinessUpdateChatResponseConfiguration {
			qbusiness_UpdateChatResponseConfiguration(cfg, client)
			return
		}
		if _qbusinessUpdateDataAccessor {
			qbusiness_UpdateDataAccessor(cfg, client)
			return
		}
		if _qbusinessUpdateDataSource {
			qbusiness_UpdateDataSource(cfg, client)
			return
		}
		if _qbusinessUpdateIndex {
			qbusiness_UpdateIndex(cfg, client)
			return
		}
		if _qbusinessUpdatePlugin {
			qbusiness_UpdatePlugin(cfg, client)
			return
		}
		if _qbusinessUpdateRetriever {
			qbusiness_UpdateRetriever(cfg, client)
			return
		}
		if _qbusinessUpdateSubscription {
			qbusiness_UpdateSubscription(cfg, client)
			return
		}
		if _qbusinessUpdateUser {
			qbusiness_UpdateUser(cfg, client)
			return
		}
		if _qbusinessUpdateWebExperience {
			qbusiness_UpdateWebExperience(cfg, client)
			return
		}

	},
}

var (
	_qbusinessAssociatePermission             bool
	_qbusinessBatchDeleteDocument             bool
	_qbusinessBatchPutDocument                bool
	_qbusinessCancelSubscription              bool
	_qbusinessChat                            bool
	_qbusinessChatSync                        bool
	_qbusinessCheckDocumentAccess             bool
	_qbusinessCreateAnonymousWebExperienceUrl bool
	_qbusinessCreateApplication               bool
	_qbusinessCreateChatResponseConfiguration bool
	_qbusinessCreateDataAccessor              bool
	_qbusinessCreateDataSource                bool
	_qbusinessCreateIndex                     bool
	_qbusinessCreatePlugin                    bool
	_qbusinessCreateRetriever                 bool
	_qbusinessCreateSubscription              bool
	_qbusinessCreateUser                      bool
	_qbusinessCreateWebExperience             bool
	_qbusinessDeleteApplication               bool
	_qbusinessDeleteAttachment                bool
	_qbusinessDeleteChatControlsConfiguration bool
	_qbusinessDeleteChatResponseConfiguration bool
	_qbusinessDeleteConversation              bool
	_qbusinessDeleteDataAccessor              bool
	_qbusinessDeleteDataSource                bool
	_qbusinessDeleteGroup                     bool
	_qbusinessDeleteIndex                     bool
	_qbusinessDeletePlugin                    bool
	_qbusinessDeleteRetriever                 bool
	_qbusinessDeleteUser                      bool
	_qbusinessDeleteWebExperience             bool
	_qbusinessDisassociatePermission          bool
	_qbusinessGetApplication                  bool
	_qbusinessGetChatControlsConfiguration    bool
	_qbusinessGetChatResponseConfiguration    bool
	_qbusinessGetDataAccessor                 bool
	_qbusinessGetDataSource                   bool
	_qbusinessGetDocumentContent              bool
	_qbusinessGetGroup                        bool
	_qbusinessGetIndex                        bool
	_qbusinessGetMedia                        bool
	_qbusinessGetPlugin                       bool
	_qbusinessGetPolicy                       bool
	_qbusinessGetRetriever                    bool
	_qbusinessGetUser                         bool
	_qbusinessGetWebExperience                bool
	_qbusinessListApplications                bool
	_qbusinessListAttachments                 bool
	_qbusinessListChatResponseConfigurations  bool
	_qbusinessListConversations               bool
	_qbusinessListDataAccessors               bool
	_qbusinessListDataSourceSyncJobs          bool
	_qbusinessListDataSources                 bool
	_qbusinessListDocuments                   bool
	_qbusinessListGroups                      bool
	_qbusinessListIndices                     bool
	_qbusinessListMessages                    bool
	_qbusinessListPluginActions               bool
	_qbusinessListPluginTypeActions           bool
	_qbusinessListPluginTypeMetadata          bool
	_qbusinessListPlugins                     bool
	_qbusinessListRetrievers                  bool
	_qbusinessListSubscriptions               bool
	_qbusinessListTagsForResource             bool
	_qbusinessListWebExperiences              bool
	_qbusinessPutFeedback                     bool
	_qbusinessPutGroup                        bool
	_qbusinessSearchRelevantContent           bool
	_qbusinessStartDataSourceSyncJob          bool
	_qbusinessStopDataSourceSyncJob           bool
	_qbusinessTagResource                     bool
	_qbusinessUntagResource                   bool
	_qbusinessUpdateApplication               bool
	_qbusinessUpdateChatControlsConfiguration bool
	_qbusinessUpdateChatResponseConfiguration bool
	_qbusinessUpdateDataAccessor              bool
	_qbusinessUpdateDataSource                bool
	_qbusinessUpdateIndex                     bool
	_qbusinessUpdatePlugin                    bool
	_qbusinessUpdateRetriever                 bool
	_qbusinessUpdateSubscription              bool
	_qbusinessUpdateUser                      bool
	_qbusinessUpdateWebExperience             bool

	_qbusinessActionConfigurations                string
	_qbusinessActionExecution                     string
	_qbusinessActions                             []string
	_qbusinessApplicationId                       string
	_qbusinessAttachmentId                        string
	_qbusinessAttachments                         string
	_qbusinessAttachmentsConfiguration            string
	_qbusinessAttributeFilter                     string
	_qbusinessAuthChallengeResponse               string
	_qbusinessAuthConfiguration                   string
	_qbusinessAuthenticationConfiguration         string
	_qbusinessAuthenticationDetail                string
	_qbusinessAutoSubscriptionConfiguration       string
	_qbusinessBlockedPhrasesConfigurationUpdate   string
	_qbusinessBrowserExtensionConfiguration       string
	_qbusinessCapacityConfiguration               string
	_qbusinessChatMode                            string
	_qbusinessChatModeConfiguration               string
	_qbusinessChatResponseConfigurationId         string
	_qbusinessClientIdsForOIDC                    []string
	_qbusinessClientToken                         string
	_qbusinessConditions                          string
	_qbusinessConfiguration                       string
	_qbusinessContentSource                       string
	_qbusinessConversationId                      string
	_qbusinessCreatorModeConfiguration            string
	_qbusinessCustomPluginConfiguration           string
	_qbusinessCustomizationConfiguration          string
	_qbusinessDataAccessorId                      string
	_qbusinessDataSourceId                        string
	_qbusinessDataSourceIds                       []string
	_qbusinessDataSourceSyncId                    string
	_qbusinessDescription                         string
	_qbusinessDisplayName                         string
	_qbusinessDocumentAttributeConfigurations     string
	_qbusinessDocumentEnrichmentConfiguration     string
	_qbusinessDocumentId                          string
	_qbusinessDocuments                           string
	_qbusinessEncryptionConfiguration             string
	_qbusinessEndTime                             string
	_qbusinessGroupMembers                        string
	_qbusinessGroupName                           string
	_qbusinessHallucinationReductionConfiguration string
	_qbusinessIamIdentityProviderArn              string
	_qbusinessIdentityCenterInstanceArn           string
	_qbusinessIdentityProviderConfiguration       string
	_qbusinessIdentityType                        string
	_qbusinessIndexId                             string
	_qbusinessMaxResults                          string
	_qbusinessMediaExtractionConfiguration        string
	_qbusinessMediaId                             string
	_qbusinessMessageCopiedAt                     string
	_qbusinessMessageId                           string
	_qbusinessMessageUsefulness                   string
	_qbusinessNextToken                           string
	_qbusinessOrchestrationConfiguration          string
	_qbusinessOrigins                             []string
	_qbusinessOutputFormat                        string
	_qbusinessParentMessageId                     string
	_qbusinessPersonalizationConfiguration        string
	_qbusinessPluginId                            string
	_qbusinessPluginType                          string
	_qbusinessPrincipal                           string
	_qbusinessQAppsConfiguration                  string
	_qbusinessQueryText                           string
	_qbusinessQuickSightConfiguration             string
	_qbusinessResourceARN                         string
	_qbusinessResponseConfigurations              string
	_qbusinessResponseScope                       string
	_qbusinessRetrieverId                         string
	_qbusinessRoleArn                             string
	_qbusinessSamplePromptsControlMode            string
	_qbusinessServerUrl                           string
	_qbusinessSessionDurationInMinutes            string
	_qbusinessStartTime                           string
	_qbusinessState                               string
	_qbusinessStatementId                         string
	_qbusinessStatusFilter                        string
	_qbusinessSubscriptionId                      string
	_qbusinessSubtitle                            string
	_qbusinessSyncSchedule                        string
	_qbusinessTagKeys                             []string
	_qbusinessTags                                string
	_qbusinessTitle                               string
	_qbusinessTopicConfigurationsToCreateOrUpdate string
	_qbusinessTopicConfigurationsToDelete         string
	_qbusinessType                                string
	_qbusinessUpdatedEarlierThan                  string
	_qbusinessUserAliases                         string
	_qbusinessUserAliasesToDelete                 string
	_qbusinessUserAliasesToUpdate                 string
	_qbusinessUserGroups                          []string
	_qbusinessUserId                              string
	_qbusinessUserMessage                         string
	_qbusinessVpcConfiguration                    string
	_qbusinessWebExperienceId                     string
	_qbusinessWelcomeMessage                      string
)

// Adds or updates a permission policy for a Amazon Q Business application,
// allowing cross-account access for an ISV. This operation creates a new policy
// statement for the specified Amazon Q Business application. The policy statement
// defines the IAM actions that the ISV is allowed to perform on the Amazon Q
// Business application's resources.
func qbusiness_AssociatePermission(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.AssociatePermissionInput{
		// Actions: []string, // Required
		// ApplicationId: *string, // Required
		// Principal: *string, // Required
		// StatementId: *string, // Required
	}

	if len(_qbusinessActions) > 0 {
		input.Actions = append([]string(nil), _qbusinessActions...)
	}
	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessPrincipal) > 0 {
		input.Principal = aws.String(_qbusinessPrincipal)
	}
	if len(_qbusinessStatementId) > 0 {
		input.StatementId = aws.String(_qbusinessStatementId)
	}
	if len(_qbusinessConditions) > 0 {
		if err := assignInputField(input, "Conditions", _qbusinessConditions); err != nil {
			log.Errorf("invalid --conditions: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociatePermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Asynchronously deletes one or more documents added using the BatchPutDocument
// API from an Amazon Q Business index.
//
// You can see the progress of the deletion, and any error messages related to the
// process, by using CloudWatch.
func qbusiness_BatchDeleteDocument(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.BatchDeleteDocumentInput{
		// ApplicationId: *string, // Required
		// Documents: []types.DeleteDocument, // Required
		// IndexId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessDocuments) > 0 {
		if err := assignInputField(input, "Documents", _qbusinessDocuments); err != nil {
			log.Errorf("invalid --documents: %s", err.Error())
			return
		}
	}
	if len(_qbusinessIndexId) > 0 {
		input.IndexId = aws.String(_qbusinessIndexId)
	}
	if len(_qbusinessDataSourceSyncId) > 0 {
		input.DataSourceSyncId = aws.String(_qbusinessDataSourceSyncId)
	}

	if resp, err := client.BatchDeleteDocument(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more documents to an Amazon Q Business index.
// You use this API to:
//
// - ingest your structured and unstructured documents and documents stored in
// an Amazon S3 bucket into an Amazon Q Business index.
//
// - add custom attributes to documents in an Amazon Q Business index.
//
// - attach an access control list to the documents added to an Amazon Q
// Business index.
//
// You can see the progress of the deletion, and any error messages related to the
// process, by using CloudWatch.
func qbusiness_BatchPutDocument(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.BatchPutDocumentInput{
		// ApplicationId: *string, // Required
		// Documents: []types.Document, // Required
		// IndexId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessDocuments) > 0 {
		if err := assignInputField(input, "Documents", _qbusinessDocuments); err != nil {
			log.Errorf("invalid --documents: %s", err.Error())
			return
		}
	}
	if len(_qbusinessIndexId) > 0 {
		input.IndexId = aws.String(_qbusinessIndexId)
	}
	if len(_qbusinessDataSourceSyncId) > 0 {
		input.DataSourceSyncId = aws.String(_qbusinessDataSourceSyncId)
	}
	if len(_qbusinessRoleArn) > 0 {
		input.RoleArn = aws.String(_qbusinessRoleArn)
	}

	if resp, err := client.BatchPutDocument(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Unsubscribes a user or a group from their pricing tier in an Amazon Q Business
// application. An unsubscribed user or group loses all Amazon Q Business feature
// access at the start of next month.
func qbusiness_CancelSubscription(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.CancelSubscriptionInput{
		// ApplicationId: *string, // Required
		// SubscriptionId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessSubscriptionId) > 0 {
		input.SubscriptionId = aws.String(_qbusinessSubscriptionId)
	}

	if resp, err := client.CancelSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts or continues a streaming Amazon Q Business conversation.
func qbusiness_Chat(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.ChatInput{
		// ApplicationId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessClientToken) > 0 {
		input.ClientToken = aws.String(_qbusinessClientToken)
	}
	if len(_qbusinessConversationId) > 0 {
		input.ConversationId = aws.String(_qbusinessConversationId)
	}
	if len(_qbusinessParentMessageId) > 0 {
		input.ParentMessageId = aws.String(_qbusinessParentMessageId)
	}
	if len(_qbusinessUserGroups) > 0 {
		input.UserGroups = append([]string(nil), _qbusinessUserGroups...)
	}
	if len(_qbusinessUserId) > 0 {
		input.UserId = aws.String(_qbusinessUserId)
	}

	if resp, err := client.Chat(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts or continues a non-streaming Amazon Q Business conversation.
func qbusiness_ChatSync(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.ChatSyncInput{
		// ApplicationId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessActionExecution) > 0 {
		if err := assignInputField(input, "ActionExecution", _qbusinessActionExecution); err != nil {
			log.Errorf("invalid --action-execution: %s", err.Error())
			return
		}
	}
	if len(_qbusinessAttachments) > 0 {
		if err := assignInputField(input, "Attachments", _qbusinessAttachments); err != nil {
			log.Errorf("invalid --attachments: %s", err.Error())
			return
		}
	}
	if len(_qbusinessAttributeFilter) > 0 {
		if err := assignInputField(input, "AttributeFilter", _qbusinessAttributeFilter); err != nil {
			log.Errorf("invalid --attribute-filter: %s", err.Error())
			return
		}
	}
	if len(_qbusinessAuthChallengeResponse) > 0 {
		if err := assignInputField(input, "AuthChallengeResponse", _qbusinessAuthChallengeResponse); err != nil {
			log.Errorf("invalid --auth-challenge-response: %s", err.Error())
			return
		}
	}
	if len(_qbusinessChatMode) > 0 {
		if err := assignInputField(input, "ChatMode", _qbusinessChatMode); err != nil {
			log.Errorf("invalid --chat-mode: %s", err.Error())
			return
		}
	}
	if len(_qbusinessChatModeConfiguration) > 0 {
		if err := assignInputField(input, "ChatModeConfiguration", _qbusinessChatModeConfiguration); err != nil {
			log.Errorf("invalid --chat-mode-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessClientToken) > 0 {
		input.ClientToken = aws.String(_qbusinessClientToken)
	}
	if len(_qbusinessConversationId) > 0 {
		input.ConversationId = aws.String(_qbusinessConversationId)
	}
	if len(_qbusinessParentMessageId) > 0 {
		input.ParentMessageId = aws.String(_qbusinessParentMessageId)
	}
	if len(_qbusinessUserGroups) > 0 {
		input.UserGroups = append([]string(nil), _qbusinessUserGroups...)
	}
	if len(_qbusinessUserId) > 0 {
		input.UserId = aws.String(_qbusinessUserId)
	}
	if len(_qbusinessUserMessage) > 0 {
		input.UserMessage = aws.String(_qbusinessUserMessage)
	}

	if resp, err := client.ChatSync(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Verifies if a user has access permissions for a specified document and returns
// the actual ACL attached to the document. Resolves user access on the document
// via user aliases and groups when verifying user access.
func qbusiness_CheckDocumentAccess(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.CheckDocumentAccessInput{
		// ApplicationId: *string, // Required
		// DocumentId: *string, // Required
		// IndexId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessDocumentId) > 0 {
		input.DocumentId = aws.String(_qbusinessDocumentId)
	}
	if len(_qbusinessIndexId) > 0 {
		input.IndexId = aws.String(_qbusinessIndexId)
	}
	if len(_qbusinessUserId) > 0 {
		input.UserId = aws.String(_qbusinessUserId)
	}
	if len(_qbusinessDataSourceId) > 0 {
		input.DataSourceId = aws.String(_qbusinessDataSourceId)
	}

	if resp, err := client.CheckDocumentAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a unique URL for anonymous Amazon Q Business web experience. This URL
// can only be used once and must be used within 5 minutes after it's generated.
func qbusiness_CreateAnonymousWebExperienceUrl(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.CreateAnonymousWebExperienceUrlInput{
		// ApplicationId: *string, // Required
		// WebExperienceId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessWebExperienceId) > 0 {
		input.WebExperienceId = aws.String(_qbusinessWebExperienceId)
	}
	if len(_qbusinessSessionDurationInMinutes) > 0 {
		if err := assignInputField(input, "SessionDurationInMinutes", _qbusinessSessionDurationInMinutes); err != nil {
			log.Errorf("invalid --session-duration-in-minutes: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAnonymousWebExperienceUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Q Business application.
// There are new tiers for Amazon Q Business. Not all features in Amazon Q
// Business Pro are also available in Amazon Q Business Lite. For information on
// what's included in Amazon Q Business Lite and what's included in Amazon Q
// Business Pro, see [Amazon Q Business tiers]. You must use the Amazon Q Business console to assign
// subscription tiers to users.
//
// An Amazon Q Apps service linked role will be created if it's absent in the
// Amazon Web Services account when QAppsConfiguration is enabled in the request.
// For more information, see [Using service-linked roles for Q Apps].
//
// When you create an application, Amazon Q Business may securely transmit data
// for processing from your selected Amazon Web Services region, but within your
// geography. For more information, see [Cross region inference in Amazon Q Business].
//
// [Amazon Q Business tiers]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/tiers.html#user-sub-tiers
// [Using service-linked roles for Q Apps]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/using-service-linked-roles-qapps.html
// [Cross region inference in Amazon Q Business]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/cross-region-inference.html
func qbusiness_CreateApplication(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.CreateApplicationInput{
		// DisplayName: *string, // Required
	}

	if len(_qbusinessDisplayName) > 0 {
		input.DisplayName = aws.String(_qbusinessDisplayName)
	}
	if len(_qbusinessAttachmentsConfiguration) > 0 {
		if err := assignInputField(input, "AttachmentsConfiguration", _qbusinessAttachmentsConfiguration); err != nil {
			log.Errorf("invalid --attachments-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessClientIdsForOIDC) > 0 {
		input.ClientIdsForOIDC = append([]string(nil), _qbusinessClientIdsForOIDC...)
	}
	if len(_qbusinessClientToken) > 0 {
		input.ClientToken = aws.String(_qbusinessClientToken)
	}
	if len(_qbusinessDescription) > 0 {
		input.Description = aws.String(_qbusinessDescription)
	}
	if len(_qbusinessEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _qbusinessEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessIamIdentityProviderArn) > 0 {
		input.IamIdentityProviderArn = aws.String(_qbusinessIamIdentityProviderArn)
	}
	if len(_qbusinessIdentityCenterInstanceArn) > 0 {
		input.IdentityCenterInstanceArn = aws.String(_qbusinessIdentityCenterInstanceArn)
	}
	if len(_qbusinessIdentityType) > 0 {
		if err := assignInputField(input, "IdentityType", _qbusinessIdentityType); err != nil {
			log.Errorf("invalid --identity-type: %s", err.Error())
			return
		}
	}
	if len(_qbusinessPersonalizationConfiguration) > 0 {
		if err := assignInputField(input, "PersonalizationConfiguration", _qbusinessPersonalizationConfiguration); err != nil {
			log.Errorf("invalid --personalization-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessQAppsConfiguration) > 0 {
		if err := assignInputField(input, "QAppsConfiguration", _qbusinessQAppsConfiguration); err != nil {
			log.Errorf("invalid --qapps-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessQuickSightConfiguration) > 0 {
		if err := assignInputField(input, "QuickSightConfiguration", _qbusinessQuickSightConfiguration); err != nil {
			log.Errorf("invalid --quicksight-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessRoleArn) > 0 {
		input.RoleArn = aws.String(_qbusinessRoleArn)
	}
	if len(_qbusinessTags) > 0 {
		if err := assignInputField(input, "Tags", _qbusinessTags); err != nil {
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

// Creates a new chat response configuration for an Amazon Q Business application.
// This operation establishes a set of parameters that define how the system
// generates and formats responses to user queries in chat interactions.
func qbusiness_CreateChatResponseConfiguration(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.CreateChatResponseConfigurationInput{
		// ApplicationId: *string, // Required
		// DisplayName: *string, // Required
		// ResponseConfigurations: map[string]types.ResponseConfiguration, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessDisplayName) > 0 {
		input.DisplayName = aws.String(_qbusinessDisplayName)
	}
	if len(_qbusinessResponseConfigurations) > 0 {
		if err := assignInputField(input, "ResponseConfigurations", _qbusinessResponseConfigurations); err != nil {
			log.Errorf("invalid --response-configurations: %s", err.Error())
			return
		}
	}
	if len(_qbusinessClientToken) > 0 {
		input.ClientToken = aws.String(_qbusinessClientToken)
	}
	if len(_qbusinessTags) > 0 {
		if err := assignInputField(input, "Tags", _qbusinessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateChatResponseConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new data accessor for an ISV to access data from a Amazon Q Business
// application. The data accessor is an entity that represents the ISV's access to
// the Amazon Q Business application's data. It includes the IAM role ARN for the
// ISV, a friendly name, and a set of action configurations that define the
// specific actions the ISV is allowed to perform and any associated data filters.
// When the data accessor is created, an IAM Identity Center application is also
// created to manage the ISV's identity and authentication for accessing the Amazon
// Q Business application.
func qbusiness_CreateDataAccessor(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.CreateDataAccessorInput{
		// ActionConfigurations: []types.ActionConfiguration, // Required
		// ApplicationId: *string, // Required
		// DisplayName: *string, // Required
		// Principal: *string, // Required
	}

	if len(_qbusinessActionConfigurations) > 0 {
		if err := assignInputField(input, "ActionConfigurations", _qbusinessActionConfigurations); err != nil {
			log.Errorf("invalid --action-configurations: %s", err.Error())
			return
		}
	}
	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessDisplayName) > 0 {
		input.DisplayName = aws.String(_qbusinessDisplayName)
	}
	if len(_qbusinessPrincipal) > 0 {
		input.Principal = aws.String(_qbusinessPrincipal)
	}
	if len(_qbusinessAuthenticationDetail) > 0 {
		if err := assignInputField(input, "AuthenticationDetail", _qbusinessAuthenticationDetail); err != nil {
			log.Errorf("invalid --authentication-detail: %s", err.Error())
			return
		}
	}
	if len(_qbusinessClientToken) > 0 {
		input.ClientToken = aws.String(_qbusinessClientToken)
	}
	if len(_qbusinessTags) > 0 {
		if err := assignInputField(input, "Tags", _qbusinessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataAccessor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a data source connector for an Amazon Q Business application.
// CreateDataSource is a synchronous operation. The operation returns 200 if the
// data source was successfully created. Otherwise, an exception is raised.
func qbusiness_CreateDataSource(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.CreateDataSourceInput{
		// ApplicationId: *string, // Required
		// Configuration: document.Interface, // Required
		// DisplayName: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _qbusinessConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessDisplayName) > 0 {
		input.DisplayName = aws.String(_qbusinessDisplayName)
	}
	if len(_qbusinessIndexId) > 0 {
		input.IndexId = aws.String(_qbusinessIndexId)
	}
	if len(_qbusinessClientToken) > 0 {
		input.ClientToken = aws.String(_qbusinessClientToken)
	}
	if len(_qbusinessDescription) > 0 {
		input.Description = aws.String(_qbusinessDescription)
	}
	if len(_qbusinessDocumentEnrichmentConfiguration) > 0 {
		if err := assignInputField(input, "DocumentEnrichmentConfiguration", _qbusinessDocumentEnrichmentConfiguration); err != nil {
			log.Errorf("invalid --document-enrichment-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessMediaExtractionConfiguration) > 0 {
		if err := assignInputField(input, "MediaExtractionConfiguration", _qbusinessMediaExtractionConfiguration); err != nil {
			log.Errorf("invalid --media-extraction-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessRoleArn) > 0 {
		input.RoleArn = aws.String(_qbusinessRoleArn)
	}
	if len(_qbusinessSyncSchedule) > 0 {
		input.SyncSchedule = aws.String(_qbusinessSyncSchedule)
	}
	if len(_qbusinessTags) > 0 {
		if err := assignInputField(input, "Tags", _qbusinessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_qbusinessVpcConfiguration) > 0 {
		if err := assignInputField(input, "VpcConfiguration", _qbusinessVpcConfiguration); err != nil {
			log.Errorf("invalid --vpc-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Q Business index.
// To determine if index creation has completed, check the Status field returned
// from a call to DescribeIndex . The Status field is set to ACTIVE when the index
// is ready to use.
//
// Once the index is active, you can index your documents using the [BatchPutDocument]
// BatchPutDocument API or the [CreateDataSource]CreateDataSource API.
//
// [BatchPutDocument]: https://docs.aws.amazon.com/amazonq/latest/api-reference/API_BatchPutDocument.html
// [CreateDataSource]: https://docs.aws.amazon.com/amazonq/latest/api-reference/API_CreateDataSource.html
func qbusiness_CreateIndex(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.CreateIndexInput{
		// ApplicationId: *string, // Required
		// DisplayName: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessDisplayName) > 0 {
		input.DisplayName = aws.String(_qbusinessDisplayName)
	}
	if len(_qbusinessCapacityConfiguration) > 0 {
		if err := assignInputField(input, "CapacityConfiguration", _qbusinessCapacityConfiguration); err != nil {
			log.Errorf("invalid --capacity-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessClientToken) > 0 {
		input.ClientToken = aws.String(_qbusinessClientToken)
	}
	if len(_qbusinessDescription) > 0 {
		input.Description = aws.String(_qbusinessDescription)
	}
	if len(_qbusinessTags) > 0 {
		if err := assignInputField(input, "Tags", _qbusinessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_qbusinessType) > 0 {
		if err := assignInputField(input, "Type", _qbusinessType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Q Business plugin.
func qbusiness_CreatePlugin(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.CreatePluginInput{
		// ApplicationId: *string, // Required
		// AuthConfiguration: types.PluginAuthConfiguration, // Required
		// DisplayName: *string, // Required
		// Type: types.PluginType, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessAuthConfiguration) > 0 {
		if err := assignInputField(input, "AuthConfiguration", _qbusinessAuthConfiguration); err != nil {
			log.Errorf("invalid --auth-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessDisplayName) > 0 {
		input.DisplayName = aws.String(_qbusinessDisplayName)
	}
	if len(_qbusinessType) > 0 {
		if err := assignInputField(input, "Type", _qbusinessType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_qbusinessClientToken) > 0 {
		input.ClientToken = aws.String(_qbusinessClientToken)
	}
	if len(_qbusinessCustomPluginConfiguration) > 0 {
		if err := assignInputField(input, "CustomPluginConfiguration", _qbusinessCustomPluginConfiguration); err != nil {
			log.Errorf("invalid --custom-plugin-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessServerUrl) > 0 {
		input.ServerUrl = aws.String(_qbusinessServerUrl)
	}
	if len(_qbusinessTags) > 0 {
		if err := assignInputField(input, "Tags", _qbusinessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePlugin(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a retriever to your Amazon Q Business application.
func qbusiness_CreateRetriever(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.CreateRetrieverInput{
		// ApplicationId: *string, // Required
		// Configuration: types.RetrieverConfiguration, // Required
		// DisplayName: *string, // Required
		// Type: types.RetrieverType, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _qbusinessConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessDisplayName) > 0 {
		input.DisplayName = aws.String(_qbusinessDisplayName)
	}
	if len(_qbusinessType) > 0 {
		if err := assignInputField(input, "Type", _qbusinessType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_qbusinessClientToken) > 0 {
		input.ClientToken = aws.String(_qbusinessClientToken)
	}
	if len(_qbusinessRoleArn) > 0 {
		input.RoleArn = aws.String(_qbusinessRoleArn)
	}
	if len(_qbusinessTags) > 0 {
		if err := assignInputField(input, "Tags", _qbusinessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRetriever(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Subscribes an IAM Identity Center user or a group to a pricing tier for an
// Amazon Q Business application.
//
// Amazon Q Business offers two subscription tiers: Q_LITE and Q_BUSINESS .
// Subscription tier determines feature access for the user. For more information
// on subscriptions and pricing tiers, see [Amazon Q Business pricing].
//
// For an example IAM role policy for assigning subscriptions, see [Set up required permissions] in the Amazon
// Q Business User Guide.
//
// [Amazon Q Business pricing]: https://aws.amazon.com/q/business/pricing/
// [Set up required permissions]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/setting-up.html#permissions
func qbusiness_CreateSubscription(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.CreateSubscriptionInput{
		// ApplicationId: *string, // Required
		// Principal: types.SubscriptionPrincipal, // Required
		// Type: types.SubscriptionType, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessPrincipal) > 0 {
		if err := assignInputField(input, "Principal", _qbusinessPrincipal); err != nil {
			log.Errorf("invalid --principal: %s", err.Error())
			return
		}
	}
	if len(_qbusinessType) > 0 {
		if err := assignInputField(input, "Type", _qbusinessType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_qbusinessClientToken) > 0 {
		input.ClientToken = aws.String(_qbusinessClientToken)
	}

	if resp, err := client.CreateSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a universally unique identifier (UUID) mapped to a list of local user
// ids within an application.
func qbusiness_CreateUser(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.CreateUserInput{
		// ApplicationId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessUserId) > 0 {
		input.UserId = aws.String(_qbusinessUserId)
	}
	if len(_qbusinessClientToken) > 0 {
		input.ClientToken = aws.String(_qbusinessClientToken)
	}
	if len(_qbusinessUserAliases) > 0 {
		if err := assignInputField(input, "UserAliases", _qbusinessUserAliases); err != nil {
			log.Errorf("invalid --user-aliases: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Q Business web experience.
func qbusiness_CreateWebExperience(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.CreateWebExperienceInput{
		// ApplicationId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessBrowserExtensionConfiguration) > 0 {
		if err := assignInputField(input, "BrowserExtensionConfiguration", _qbusinessBrowserExtensionConfiguration); err != nil {
			log.Errorf("invalid --browser-extension-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessClientToken) > 0 {
		input.ClientToken = aws.String(_qbusinessClientToken)
	}
	if len(_qbusinessCustomizationConfiguration) > 0 {
		if err := assignInputField(input, "CustomizationConfiguration", _qbusinessCustomizationConfiguration); err != nil {
			log.Errorf("invalid --customization-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessIdentityProviderConfiguration) > 0 {
		if err := assignInputField(input, "IdentityProviderConfiguration", _qbusinessIdentityProviderConfiguration); err != nil {
			log.Errorf("invalid --identity-provider-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessOrigins) > 0 {
		input.Origins = append([]string(nil), _qbusinessOrigins...)
	}
	if len(_qbusinessRoleArn) > 0 {
		input.RoleArn = aws.String(_qbusinessRoleArn)
	}
	if len(_qbusinessSamplePromptsControlMode) > 0 {
		if err := assignInputField(input, "SamplePromptsControlMode", _qbusinessSamplePromptsControlMode); err != nil {
			log.Errorf("invalid --sample-prompts-control-mode: %s", err.Error())
			return
		}
	}
	if len(_qbusinessSubtitle) > 0 {
		input.Subtitle = aws.String(_qbusinessSubtitle)
	}
	if len(_qbusinessTags) > 0 {
		if err := assignInputField(input, "Tags", _qbusinessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_qbusinessTitle) > 0 {
		input.Title = aws.String(_qbusinessTitle)
	}
	if len(_qbusinessWelcomeMessage) > 0 {
		input.WelcomeMessage = aws.String(_qbusinessWelcomeMessage)
	}

	if resp, err := client.CreateWebExperience(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Q Business application.
func qbusiness_DeleteApplication(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.DeleteApplicationInput{
		// ApplicationId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}

	if resp, err := client.DeleteApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an attachment associated with a specific Amazon Q Business conversation.
func qbusiness_DeleteAttachment(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.DeleteAttachmentInput{
		// ApplicationId: *string, // Required
		// AttachmentId: *string, // Required
		// ConversationId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessAttachmentId) > 0 {
		input.AttachmentId = aws.String(_qbusinessAttachmentId)
	}
	if len(_qbusinessConversationId) > 0 {
		input.ConversationId = aws.String(_qbusinessConversationId)
	}
	if len(_qbusinessUserId) > 0 {
		input.UserId = aws.String(_qbusinessUserId)
	}

	if resp, err := client.DeleteAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes chat controls configured for an existing Amazon Q Business application.
func qbusiness_DeleteChatControlsConfiguration(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.DeleteChatControlsConfigurationInput{
		// ApplicationId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}

	if resp, err := client.DeleteChatControlsConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified chat response configuration from an Amazon Q Business
// application.
func qbusiness_DeleteChatResponseConfiguration(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.DeleteChatResponseConfigurationInput{
		// ApplicationId: *string, // Required
		// ChatResponseConfigurationId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessChatResponseConfigurationId) > 0 {
		input.ChatResponseConfigurationId = aws.String(_qbusinessChatResponseConfigurationId)
	}

	if resp, err := client.DeleteChatResponseConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Q Business web experience conversation.
func qbusiness_DeleteConversation(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.DeleteConversationInput{
		// ApplicationId: *string, // Required
		// ConversationId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessConversationId) > 0 {
		input.ConversationId = aws.String(_qbusinessConversationId)
	}
	if len(_qbusinessUserId) > 0 {
		input.UserId = aws.String(_qbusinessUserId)
	}

	if resp, err := client.DeleteConversation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified data accessor. This operation permanently removes the data
// accessor and its associated IAM Identity Center application. Any access granted
// to the ISV through this data accessor will be revoked.
func qbusiness_DeleteDataAccessor(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.DeleteDataAccessorInput{
		// ApplicationId: *string, // Required
		// DataAccessorId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessDataAccessorId) > 0 {
		input.DataAccessorId = aws.String(_qbusinessDataAccessorId)
	}

	if resp, err := client.DeleteDataAccessor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Q Business data source connector. While the data source is
// being deleted, the Status field returned by a call to the DescribeDataSource
// API is set to DELETING .
func qbusiness_DeleteDataSource(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.DeleteDataSourceInput{
		// ApplicationId: *string, // Required
		// DataSourceId: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessDataSourceId) > 0 {
		input.DataSourceId = aws.String(_qbusinessDataSourceId)
	}
	if len(_qbusinessIndexId) > 0 {
		input.IndexId = aws.String(_qbusinessIndexId)
	}

	if resp, err := client.DeleteDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a group so that all users and sub groups that belong to the group can
// no longer access documents only available to that group. For example, after
// deleting the group "Summer Interns", all interns who belonged to that group no
// longer see intern-only documents in their chat results.
//
// If you want to delete, update, or replace users or sub groups of a group, you
// need to use the PutGroup operation. For example, if a user in the group
// "Engineering" leaves the engineering team and another user takes their place,
// you provide an updated list of users or sub groups that belong to the
// "Engineering" group when calling PutGroup .
func qbusiness_DeleteGroup(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.DeleteGroupInput{
		// ApplicationId: *string, // Required
		// GroupName: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessGroupName) > 0 {
		input.GroupName = aws.String(_qbusinessGroupName)
	}
	if len(_qbusinessIndexId) > 0 {
		input.IndexId = aws.String(_qbusinessIndexId)
	}
	if len(_qbusinessDataSourceId) > 0 {
		input.DataSourceId = aws.String(_qbusinessDataSourceId)
	}

	if resp, err := client.DeleteGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Q Business index.
func qbusiness_DeleteIndex(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.DeleteIndexInput{
		// ApplicationId: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessIndexId) > 0 {
		input.IndexId = aws.String(_qbusinessIndexId)
	}

	if resp, err := client.DeleteIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Q Business plugin.
func qbusiness_DeletePlugin(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.DeletePluginInput{
		// ApplicationId: *string, // Required
		// PluginId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessPluginId) > 0 {
		input.PluginId = aws.String(_qbusinessPluginId)
	}

	if resp, err := client.DeletePlugin(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the retriever used by an Amazon Q Business application.
func qbusiness_DeleteRetriever(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.DeleteRetrieverInput{
		// ApplicationId: *string, // Required
		// RetrieverId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessRetrieverId) > 0 {
		input.RetrieverId = aws.String(_qbusinessRetrieverId)
	}

	if resp, err := client.DeleteRetriever(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a user by email id.
func qbusiness_DeleteUser(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.DeleteUserInput{
		// ApplicationId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessUserId) > 0 {
		input.UserId = aws.String(_qbusinessUserId)
	}

	if resp, err := client.DeleteUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Q Business web experience.
func qbusiness_DeleteWebExperience(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.DeleteWebExperienceInput{
		// ApplicationId: *string, // Required
		// WebExperienceId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessWebExperienceId) > 0 {
		input.WebExperienceId = aws.String(_qbusinessWebExperienceId)
	}

	if resp, err := client.DeleteWebExperience(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a permission policy from a Amazon Q Business application, revoking the
// cross-account access that was previously granted to an ISV. This operation
// deletes the specified policy statement from the application's permission policy.
func qbusiness_DisassociatePermission(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.DisassociatePermissionInput{
		// ApplicationId: *string, // Required
		// StatementId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessStatementId) > 0 {
		input.StatementId = aws.String(_qbusinessStatementId)
	}

	if resp, err := client.DisassociatePermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an existing Amazon Q Business application.
func qbusiness_GetApplication(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.GetApplicationInput{
		// ApplicationId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}

	if resp, err := client.GetApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about chat controls configured for an existing Amazon Q
// Business application.
func qbusiness_GetChatControlsConfiguration(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.GetChatControlsConfigurationInput{
		// ApplicationId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qbusinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qbusinessNextToken) > 0 {
		input.NextToken = aws.String(_qbusinessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetChatControlsConfiguration(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qbusiness.GetChatControlsConfigurationOutput
	p := qbusiness.NewGetChatControlsConfigurationPaginator(client, input)
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

// Retrieves detailed information about a specific chat response configuration
// from an Amazon Q Business application. This operation returns the complete
// configuration settings and metadata.
func qbusiness_GetChatResponseConfiguration(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.GetChatResponseConfigurationInput{
		// ApplicationId: *string, // Required
		// ChatResponseConfigurationId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessChatResponseConfigurationId) > 0 {
		input.ChatResponseConfigurationId = aws.String(_qbusinessChatResponseConfigurationId)
	}

	if resp, err := client.GetChatResponseConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a specified data accessor. This operation returns
// details about the data accessor, including its display name, unique identifier,
// Amazon Resource Name (ARN), the associated Amazon Q Business application and IAM
// Identity Center application, the IAM role for the ISV, the action
// configurations, and the timestamps for when the data accessor was created and
// last updated.
func qbusiness_GetDataAccessor(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.GetDataAccessorInput{
		// ApplicationId: *string, // Required
		// DataAccessorId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessDataAccessorId) > 0 {
		input.DataAccessorId = aws.String(_qbusinessDataAccessorId)
	}

	if resp, err := client.GetDataAccessor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an existing Amazon Q Business data source connector.
func qbusiness_GetDataSource(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.GetDataSourceInput{
		// ApplicationId: *string, // Required
		// DataSourceId: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessDataSourceId) > 0 {
		input.DataSourceId = aws.String(_qbusinessDataSourceId)
	}
	if len(_qbusinessIndexId) > 0 {
		input.IndexId = aws.String(_qbusinessIndexId)
	}

	if resp, err := client.GetDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the content of a document that was ingested into Amazon Q Business.
// This API validates user authorization against document ACLs before returning a
// pre-signed URL for secure document access. You can download or view source
// documents referenced in chat responses through the URL.
func qbusiness_GetDocumentContent(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.GetDocumentContentInput{
		// ApplicationId: *string, // Required
		// DocumentId: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessDocumentId) > 0 {
		input.DocumentId = aws.String(_qbusinessDocumentId)
	}
	if len(_qbusinessIndexId) > 0 {
		input.IndexId = aws.String(_qbusinessIndexId)
	}
	if len(_qbusinessDataSourceId) > 0 {
		input.DataSourceId = aws.String(_qbusinessDataSourceId)
	}
	if len(_qbusinessOutputFormat) > 0 {
		if err := assignInputField(input, "OutputFormat", _qbusinessOutputFormat); err != nil {
			log.Errorf("invalid --output-format: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetDocumentContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a group by group name.
func qbusiness_GetGroup(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.GetGroupInput{
		// ApplicationId: *string, // Required
		// GroupName: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessGroupName) > 0 {
		input.GroupName = aws.String(_qbusinessGroupName)
	}
	if len(_qbusinessIndexId) > 0 {
		input.IndexId = aws.String(_qbusinessIndexId)
	}
	if len(_qbusinessDataSourceId) > 0 {
		input.DataSourceId = aws.String(_qbusinessDataSourceId)
	}

	if resp, err := client.GetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an existing Amazon Q Business index.
func qbusiness_GetIndex(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.GetIndexInput{
		// ApplicationId: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessIndexId) > 0 {
		input.IndexId = aws.String(_qbusinessIndexId)
	}

	if resp, err := client.GetIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the image bytes corresponding to a media object. If you have
// implemented your own application with the Chat and ChatSync APIs, and have
// enabled content extraction from visual data in Amazon Q Business, you use the
// GetMedia API operation to download the images so you can show them in your UI
// with responses.
//
// For more information, see [Extracting semantic meaning from images and visuals].
//
// [Extracting semantic meaning from images and visuals]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/extracting-meaning-from-images.html
func qbusiness_GetMedia(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.GetMediaInput{
		// ApplicationId: *string, // Required
		// ConversationId: *string, // Required
		// MediaId: *string, // Required
		// MessageId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessConversationId) > 0 {
		input.ConversationId = aws.String(_qbusinessConversationId)
	}
	if len(_qbusinessMediaId) > 0 {
		input.MediaId = aws.String(_qbusinessMediaId)
	}
	if len(_qbusinessMessageId) > 0 {
		input.MessageId = aws.String(_qbusinessMessageId)
	}

	if resp, err := client.GetMedia(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an existing Amazon Q Business plugin.
func qbusiness_GetPlugin(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.GetPluginInput{
		// ApplicationId: *string, // Required
		// PluginId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessPluginId) > 0 {
		input.PluginId = aws.String(_qbusinessPluginId)
	}

	if resp, err := client.GetPlugin(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the current permission policy for a Amazon Q Business application.
// The policy is returned as a JSON-formatted string and defines the IAM actions
// that are allowed or denied for the application's resources.
func qbusiness_GetPolicy(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.GetPolicyInput{
		// ApplicationId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}

	if resp, err := client.GetPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an existing retriever used by an Amazon Q Business
// application.
func qbusiness_GetRetriever(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.GetRetrieverInput{
		// ApplicationId: *string, // Required
		// RetrieverId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessRetrieverId) > 0 {
		input.RetrieverId = aws.String(_qbusinessRetrieverId)
	}

	if resp, err := client.GetRetriever(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the universally unique identifier (UUID) associated with a local user
// in a data source.
func qbusiness_GetUser(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.GetUserInput{
		// ApplicationId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessUserId) > 0 {
		input.UserId = aws.String(_qbusinessUserId)
	}

	if resp, err := client.GetUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an existing Amazon Q Business web experience.
func qbusiness_GetWebExperience(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.GetWebExperienceInput{
		// ApplicationId: *string, // Required
		// WebExperienceId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessWebExperienceId) > 0 {
		input.WebExperienceId = aws.String(_qbusinessWebExperienceId)
	}

	if resp, err := client.GetWebExperience(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists Amazon Q Business applications.
// Amazon Q Business applications may securely transmit data for processing across
// Amazon Web Services Regions within your geography. For more information, see [Cross region inference in Amazon Q Business].
//
// [Cross region inference in Amazon Q Business]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/cross-region-inference.html
func qbusiness_ListApplications(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.ListApplicationsInput{}

	if len(_qbusinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qbusinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qbusinessNextToken) > 0 {
		input.NextToken = aws.String(_qbusinessNextToken)
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

	var results []*qbusiness.ListApplicationsOutput
	p := qbusiness.NewListApplicationsPaginator(client, input)
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

// Gets a list of attachments associated with an Amazon Q Business web experience
// or a list of attachements associated with a specific Amazon Q Business
// conversation.
func qbusiness_ListAttachments(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.ListAttachmentsInput{
		// ApplicationId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessConversationId) > 0 {
		input.ConversationId = aws.String(_qbusinessConversationId)
	}
	if len(_qbusinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qbusinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qbusinessNextToken) > 0 {
		input.NextToken = aws.String(_qbusinessNextToken)
	}
	if len(_qbusinessUserId) > 0 {
		input.UserId = aws.String(_qbusinessUserId)
	}

	if disablePaginator() {
		if resp, err := client.ListAttachments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qbusiness.ListAttachmentsOutput
	p := qbusiness.NewListAttachmentsPaginator(client, input)
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

// Retrieves a list of all chat response configurations available in a specified
// Amazon Q Business application. This operation returns summary information about
// each configuration to help administrators manage and select appropriate response
// settings.
func qbusiness_ListChatResponseConfigurations(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.ListChatResponseConfigurationsInput{
		// ApplicationId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qbusinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qbusinessNextToken) > 0 {
		input.NextToken = aws.String(_qbusinessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListChatResponseConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qbusiness.ListChatResponseConfigurationsOutput
	p := qbusiness.NewListChatResponseConfigurationsPaginator(client, input)
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

// Lists one or more Amazon Q Business conversations.
func qbusiness_ListConversations(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.ListConversationsInput{
		// ApplicationId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qbusinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qbusinessNextToken) > 0 {
		input.NextToken = aws.String(_qbusinessNextToken)
	}
	if len(_qbusinessUserId) > 0 {
		input.UserId = aws.String(_qbusinessUserId)
	}

	if disablePaginator() {
		if resp, err := client.ListConversations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qbusiness.ListConversationsOutput
	p := qbusiness.NewListConversationsPaginator(client, input)
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

// Lists the data accessors for a Amazon Q Business application. This operation
// returns a paginated list of data accessor summaries, including the friendly
// name, unique identifier, ARN, associated IAM role, and creation/update
// timestamps for each data accessor.
func qbusiness_ListDataAccessors(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.ListDataAccessorsInput{
		// ApplicationId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qbusinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qbusinessNextToken) > 0 {
		input.NextToken = aws.String(_qbusinessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataAccessors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qbusiness.ListDataAccessorsOutput
	p := qbusiness.NewListDataAccessorsPaginator(client, input)
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

// Get information about an Amazon Q Business data source connector
// synchronization.
func qbusiness_ListDataSourceSyncJobs(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.ListDataSourceSyncJobsInput{
		// ApplicationId: *string, // Required
		// DataSourceId: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessDataSourceId) > 0 {
		input.DataSourceId = aws.String(_qbusinessDataSourceId)
	}
	if len(_qbusinessIndexId) > 0 {
		input.IndexId = aws.String(_qbusinessIndexId)
	}
	if len(_qbusinessEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _qbusinessEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_qbusinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qbusinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qbusinessNextToken) > 0 {
		input.NextToken = aws.String(_qbusinessNextToken)
	}
	if len(_qbusinessStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _qbusinessStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_qbusinessStatusFilter) > 0 {
		if err := assignInputField(input, "StatusFilter", _qbusinessStatusFilter); err != nil {
			log.Errorf("invalid --status-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDataSourceSyncJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qbusiness.ListDataSourceSyncJobsOutput
	p := qbusiness.NewListDataSourceSyncJobsPaginator(client, input)
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

// Lists the Amazon Q Business data source connectors that you have created.
func qbusiness_ListDataSources(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.ListDataSourcesInput{
		// ApplicationId: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessIndexId) > 0 {
		input.IndexId = aws.String(_qbusinessIndexId)
	}
	if len(_qbusinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qbusinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qbusinessNextToken) > 0 {
		input.NextToken = aws.String(_qbusinessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataSources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qbusiness.ListDataSourcesOutput
	p := qbusiness.NewListDataSourcesPaginator(client, input)
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

// A list of documents attached to an index.
func qbusiness_ListDocuments(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.ListDocumentsInput{
		// ApplicationId: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessIndexId) > 0 {
		input.IndexId = aws.String(_qbusinessIndexId)
	}
	if len(_qbusinessDataSourceIds) > 0 {
		input.DataSourceIds = append([]string(nil), _qbusinessDataSourceIds...)
	}
	if len(_qbusinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qbusinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qbusinessNextToken) > 0 {
		input.NextToken = aws.String(_qbusinessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDocuments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qbusiness.ListDocumentsOutput
	p := qbusiness.NewListDocumentsPaginator(client, input)
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

// Provides a list of groups that are mapped to users.
func qbusiness_ListGroups(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.ListGroupsInput{
		// ApplicationId: *string, // Required
		// IndexId: *string, // Required
		// UpdatedEarlierThan: *time.Time, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessIndexId) > 0 {
		input.IndexId = aws.String(_qbusinessIndexId)
	}
	if len(_qbusinessUpdatedEarlierThan) > 0 {
		if err := assignInputField(input, "UpdatedEarlierThan", _qbusinessUpdatedEarlierThan); err != nil {
			log.Errorf("invalid --updated-earlier-than: %s", err.Error())
			return
		}
	}
	if len(_qbusinessDataSourceId) > 0 {
		input.DataSourceId = aws.String(_qbusinessDataSourceId)
	}
	if len(_qbusinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qbusinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qbusinessNextToken) > 0 {
		input.NextToken = aws.String(_qbusinessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qbusiness.ListGroupsOutput
	p := qbusiness.NewListGroupsPaginator(client, input)
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

// Lists the Amazon Q Business indices you have created.
func qbusiness_ListIndices(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.ListIndicesInput{
		// ApplicationId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qbusinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qbusinessNextToken) > 0 {
		input.NextToken = aws.String(_qbusinessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIndices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qbusiness.ListIndicesOutput
	p := qbusiness.NewListIndicesPaginator(client, input)
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

// Gets a list of messages associated with an Amazon Q Business web experience.
func qbusiness_ListMessages(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.ListMessagesInput{
		// ApplicationId: *string, // Required
		// ConversationId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessConversationId) > 0 {
		input.ConversationId = aws.String(_qbusinessConversationId)
	}
	if len(_qbusinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qbusinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qbusinessNextToken) > 0 {
		input.NextToken = aws.String(_qbusinessNextToken)
	}
	if len(_qbusinessUserId) > 0 {
		input.UserId = aws.String(_qbusinessUserId)
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

	var results []*qbusiness.ListMessagesOutput
	p := qbusiness.NewListMessagesPaginator(client, input)
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

// Lists configured Amazon Q Business actions for a specific plugin in an Amazon Q
// Business application.
func qbusiness_ListPluginActions(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.ListPluginActionsInput{
		// ApplicationId: *string, // Required
		// PluginId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessPluginId) > 0 {
		input.PluginId = aws.String(_qbusinessPluginId)
	}
	if len(_qbusinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qbusinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qbusinessNextToken) > 0 {
		input.NextToken = aws.String(_qbusinessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPluginActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qbusiness.ListPluginActionsOutput
	p := qbusiness.NewListPluginActionsPaginator(client, input)
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

// Lists configured Amazon Q Business actions for any plugin type—both built-in
// and custom.
func qbusiness_ListPluginTypeActions(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.ListPluginTypeActionsInput{
		// PluginType: types.PluginType, // Required
	}

	if len(_qbusinessPluginType) > 0 {
		if err := assignInputField(input, "PluginType", _qbusinessPluginType); err != nil {
			log.Errorf("invalid --plugin-type: %s", err.Error())
			return
		}
	}
	if len(_qbusinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qbusinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qbusinessNextToken) > 0 {
		input.NextToken = aws.String(_qbusinessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPluginTypeActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qbusiness.ListPluginTypeActionsOutput
	p := qbusiness.NewListPluginTypeActionsPaginator(client, input)
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

// Lists metadata for all Amazon Q Business plugin types.
func qbusiness_ListPluginTypeMetadata(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.ListPluginTypeMetadataInput{}

	if len(_qbusinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qbusinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qbusinessNextToken) > 0 {
		input.NextToken = aws.String(_qbusinessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPluginTypeMetadata(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qbusiness.ListPluginTypeMetadataOutput
	p := qbusiness.NewListPluginTypeMetadataPaginator(client, input)
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

// Lists configured Amazon Q Business plugins.
func qbusiness_ListPlugins(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.ListPluginsInput{
		// ApplicationId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qbusinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qbusinessNextToken) > 0 {
		input.NextToken = aws.String(_qbusinessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPlugins(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qbusiness.ListPluginsOutput
	p := qbusiness.NewListPluginsPaginator(client, input)
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

// Lists the retriever used by an Amazon Q Business application.
func qbusiness_ListRetrievers(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.ListRetrieversInput{
		// ApplicationId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qbusinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qbusinessNextToken) > 0 {
		input.NextToken = aws.String(_qbusinessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRetrievers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qbusiness.ListRetrieversOutput
	p := qbusiness.NewListRetrieversPaginator(client, input)
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

// Lists all subscriptions created in an Amazon Q Business application.
func qbusiness_ListSubscriptions(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.ListSubscriptionsInput{
		// ApplicationId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qbusinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qbusinessNextToken) > 0 {
		input.NextToken = aws.String(_qbusinessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSubscriptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qbusiness.ListSubscriptionsOutput
	p := qbusiness.NewListSubscriptionsPaginator(client, input)
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

// Gets a list of tags associated with a specified resource. Amazon Q Business
// applications and data sources can have tags associated with them.
func qbusiness_ListTagsForResource(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_qbusinessResourceARN) > 0 {
		input.ResourceARN = aws.String(_qbusinessResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists one or more Amazon Q Business Web Experiences.
func qbusiness_ListWebExperiences(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.ListWebExperiencesInput{
		// ApplicationId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qbusinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qbusinessNextToken) > 0 {
		input.NextToken = aws.String(_qbusinessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWebExperiences(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qbusiness.ListWebExperiencesOutput
	p := qbusiness.NewListWebExperiencesPaginator(client, input)
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

// Enables your end user to provide feedback on their Amazon Q Business generated
// chat responses.
func qbusiness_PutFeedback(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.PutFeedbackInput{
		// ApplicationId: *string, // Required
		// ConversationId: *string, // Required
		// MessageId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessConversationId) > 0 {
		input.ConversationId = aws.String(_qbusinessConversationId)
	}
	if len(_qbusinessMessageId) > 0 {
		input.MessageId = aws.String(_qbusinessMessageId)
	}
	if len(_qbusinessMessageCopiedAt) > 0 {
		if err := assignInputField(input, "MessageCopiedAt", _qbusinessMessageCopiedAt); err != nil {
			log.Errorf("invalid --message-copied-at: %s", err.Error())
			return
		}
	}
	if len(_qbusinessMessageUsefulness) > 0 {
		if err := assignInputField(input, "MessageUsefulness", _qbusinessMessageUsefulness); err != nil {
			log.Errorf("invalid --message-usefulness: %s", err.Error())
			return
		}
	}
	if len(_qbusinessUserId) > 0 {
		input.UserId = aws.String(_qbusinessUserId)
	}

	if resp, err := client.PutFeedback(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create, or updates, a mapping of users—who have access to a document—to groups.
// You can also map sub groups to groups. For example, the group "Company
// Intellectual Property Teams" includes sub groups "Research" and "Engineering".
// These sub groups include their own list of users or people who work in these
// teams. Only users who work in research and engineering, and therefore belong in
// the intellectual property group, can see top-secret company documents in their
// Amazon Q Business chat results.
//
// There are two options for creating groups, either passing group members inline
// or using an S3 file via the S3PathForGroupMembers field. For inline groups,
// there is a limit of 1000 members per group and for provided S3 files there is a
// limit of 100 thousand members. When creating a group using an S3 file, you
// provide both an S3 file and a RoleArn for Amazon Q Buisness to access the file.
func qbusiness_PutGroup(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.PutGroupInput{
		// ApplicationId: *string, // Required
		// GroupMembers: *types.GroupMembers, // Required
		// GroupName: *string, // Required
		// IndexId: *string, // Required
		// Type: types.MembershipType, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessGroupMembers) > 0 {
		if err := assignInputField(input, "GroupMembers", _qbusinessGroupMembers); err != nil {
			log.Errorf("invalid --group-members: %s", err.Error())
			return
		}
	}
	if len(_qbusinessGroupName) > 0 {
		input.GroupName = aws.String(_qbusinessGroupName)
	}
	if len(_qbusinessIndexId) > 0 {
		input.IndexId = aws.String(_qbusinessIndexId)
	}
	if len(_qbusinessType) > 0 {
		if err := assignInputField(input, "Type", _qbusinessType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_qbusinessDataSourceId) > 0 {
		input.DataSourceId = aws.String(_qbusinessDataSourceId)
	}
	if len(_qbusinessRoleArn) > 0 {
		input.RoleArn = aws.String(_qbusinessRoleArn)
	}

	if resp, err := client.PutGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches for relevant content in a Amazon Q Business application based on a
// query. This operation takes a search query text, the Amazon Q Business
// application identifier, and optional filters (such as content source and maximum
// results) as input. It returns a list of relevant content items, where each item
// includes the content text, the unique document identifier, the document title,
// the document URI, any relevant document attributes, and score attributes
// indicating the confidence level of the relevance.
func qbusiness_SearchRelevantContent(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.SearchRelevantContentInput{
		// ApplicationId: *string, // Required
		// ContentSource: types.ContentSource, // Required
		// QueryText: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessContentSource) > 0 {
		if err := assignInputField(input, "ContentSource", _qbusinessContentSource); err != nil {
			log.Errorf("invalid --content-source: %s", err.Error())
			return
		}
	}
	if len(_qbusinessQueryText) > 0 {
		input.QueryText = aws.String(_qbusinessQueryText)
	}
	if len(_qbusinessAttributeFilter) > 0 {
		if err := assignInputField(input, "AttributeFilter", _qbusinessAttributeFilter); err != nil {
			log.Errorf("invalid --attribute-filter: %s", err.Error())
			return
		}
	}
	if len(_qbusinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _qbusinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_qbusinessNextToken) > 0 {
		input.NextToken = aws.String(_qbusinessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchRelevantContent(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qbusiness.SearchRelevantContentOutput
	p := qbusiness.NewSearchRelevantContentPaginator(client, input)
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

// Starts a data source connector synchronization job. If a synchronization job is
// already in progress, Amazon Q Business returns a ConflictException .
func qbusiness_StartDataSourceSyncJob(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.StartDataSourceSyncJobInput{
		// ApplicationId: *string, // Required
		// DataSourceId: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessDataSourceId) > 0 {
		input.DataSourceId = aws.String(_qbusinessDataSourceId)
	}
	if len(_qbusinessIndexId) > 0 {
		input.IndexId = aws.String(_qbusinessIndexId)
	}

	if resp, err := client.StartDataSourceSyncJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an Amazon Q Business data source connector synchronization job already in
// progress.
func qbusiness_StopDataSourceSyncJob(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.StopDataSourceSyncJobInput{
		// ApplicationId: *string, // Required
		// DataSourceId: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessDataSourceId) > 0 {
		input.DataSourceId = aws.String(_qbusinessDataSourceId)
	}
	if len(_qbusinessIndexId) > 0 {
		input.IndexId = aws.String(_qbusinessIndexId)
	}

	if resp, err := client.StopDataSourceSyncJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the specified tag to the specified Amazon Q Business application or data
// source resource. If the tag already exists, the existing value is replaced with
// the new value.
func qbusiness_TagResource(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_qbusinessResourceARN) > 0 {
		input.ResourceARN = aws.String(_qbusinessResourceARN)
	}
	if len(_qbusinessTags) > 0 {
		if err := assignInputField(input, "Tags", _qbusinessTags); err != nil {
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

// Removes a tag from an Amazon Q Business application or a data source.
func qbusiness_UntagResource(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_qbusinessResourceARN) > 0 {
		input.ResourceARN = aws.String(_qbusinessResourceARN)
	}
	if len(_qbusinessTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _qbusinessTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing Amazon Q Business application.
// Amazon Q Business applications may securely transmit data for processing across
// Amazon Web Services Regions within your geography. For more information, see [Cross region inference in Amazon Q Business].
//
// An Amazon Q Apps service-linked role will be created if it's absent in the
// Amazon Web Services account when QAppsConfiguration is enabled in the request.
// For more information, see [Using service-linked roles for Q Apps].
//
// [Using service-linked roles for Q Apps]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/using-service-linked-roles-qapps.html
// [Cross region inference in Amazon Q Business]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/cross-region-inference.html
func qbusiness_UpdateApplication(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.UpdateApplicationInput{
		// ApplicationId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessAttachmentsConfiguration) > 0 {
		if err := assignInputField(input, "AttachmentsConfiguration", _qbusinessAttachmentsConfiguration); err != nil {
			log.Errorf("invalid --attachments-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessAutoSubscriptionConfiguration) > 0 {
		if err := assignInputField(input, "AutoSubscriptionConfiguration", _qbusinessAutoSubscriptionConfiguration); err != nil {
			log.Errorf("invalid --auto-subscription-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessDescription) > 0 {
		input.Description = aws.String(_qbusinessDescription)
	}
	if len(_qbusinessDisplayName) > 0 {
		input.DisplayName = aws.String(_qbusinessDisplayName)
	}
	if len(_qbusinessIdentityCenterInstanceArn) > 0 {
		input.IdentityCenterInstanceArn = aws.String(_qbusinessIdentityCenterInstanceArn)
	}
	if len(_qbusinessPersonalizationConfiguration) > 0 {
		if err := assignInputField(input, "PersonalizationConfiguration", _qbusinessPersonalizationConfiguration); err != nil {
			log.Errorf("invalid --personalization-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessQAppsConfiguration) > 0 {
		if err := assignInputField(input, "QAppsConfiguration", _qbusinessQAppsConfiguration); err != nil {
			log.Errorf("invalid --qapps-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessRoleArn) > 0 {
		input.RoleArn = aws.String(_qbusinessRoleArn)
	}

	if resp, err := client.UpdateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a set of chat controls configured for an existing Amazon Q Business
// application.
func qbusiness_UpdateChatControlsConfiguration(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.UpdateChatControlsConfigurationInput{
		// ApplicationId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessBlockedPhrasesConfigurationUpdate) > 0 {
		if err := assignInputField(input, "BlockedPhrasesConfigurationUpdate", _qbusinessBlockedPhrasesConfigurationUpdate); err != nil {
			log.Errorf("invalid --blocked-phrases-configuration-update: %s", err.Error())
			return
		}
	}
	if len(_qbusinessClientToken) > 0 {
		input.ClientToken = aws.String(_qbusinessClientToken)
	}
	if len(_qbusinessCreatorModeConfiguration) > 0 {
		if err := assignInputField(input, "CreatorModeConfiguration", _qbusinessCreatorModeConfiguration); err != nil {
			log.Errorf("invalid --creator-mode-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessHallucinationReductionConfiguration) > 0 {
		if err := assignInputField(input, "HallucinationReductionConfiguration", _qbusinessHallucinationReductionConfiguration); err != nil {
			log.Errorf("invalid --hallucination-reduction-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessOrchestrationConfiguration) > 0 {
		if err := assignInputField(input, "OrchestrationConfiguration", _qbusinessOrchestrationConfiguration); err != nil {
			log.Errorf("invalid --orchestration-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessResponseScope) > 0 {
		if err := assignInputField(input, "ResponseScope", _qbusinessResponseScope); err != nil {
			log.Errorf("invalid --response-scope: %s", err.Error())
			return
		}
	}
	if len(_qbusinessTopicConfigurationsToCreateOrUpdate) > 0 {
		if err := assignInputField(input, "TopicConfigurationsToCreateOrUpdate", _qbusinessTopicConfigurationsToCreateOrUpdate); err != nil {
			log.Errorf("invalid --topic-configurations-to-create-or-update: %s", err.Error())
			return
		}
	}
	if len(_qbusinessTopicConfigurationsToDelete) > 0 {
		if err := assignInputField(input, "TopicConfigurationsToDelete", _qbusinessTopicConfigurationsToDelete); err != nil {
			log.Errorf("invalid --topic-configurations-to-delete: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateChatControlsConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing chat response configuration in an Amazon Q Business
// application. This operation allows administrators to modify configuration
// settings, display name, and response parameters to refine how the system
// generates responses.
func qbusiness_UpdateChatResponseConfiguration(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.UpdateChatResponseConfigurationInput{
		// ApplicationId: *string, // Required
		// ChatResponseConfigurationId: *string, // Required
		// ResponseConfigurations: map[string]types.ResponseConfiguration, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessChatResponseConfigurationId) > 0 {
		input.ChatResponseConfigurationId = aws.String(_qbusinessChatResponseConfigurationId)
	}
	if len(_qbusinessResponseConfigurations) > 0 {
		if err := assignInputField(input, "ResponseConfigurations", _qbusinessResponseConfigurations); err != nil {
			log.Errorf("invalid --response-configurations: %s", err.Error())
			return
		}
	}
	if len(_qbusinessClientToken) > 0 {
		input.ClientToken = aws.String(_qbusinessClientToken)
	}
	if len(_qbusinessDisplayName) > 0 {
		input.DisplayName = aws.String(_qbusinessDisplayName)
	}

	if resp, err := client.UpdateChatResponseConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing data accessor. This operation allows modifying the action
// configurations (the allowed actions and associated filters) and the display name
// of the data accessor. It does not allow changing the IAM role associated with
// the data accessor or other core properties of the data accessor.
func qbusiness_UpdateDataAccessor(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.UpdateDataAccessorInput{
		// ActionConfigurations: []types.ActionConfiguration, // Required
		// ApplicationId: *string, // Required
		// DataAccessorId: *string, // Required
	}

	if len(_qbusinessActionConfigurations) > 0 {
		if err := assignInputField(input, "ActionConfigurations", _qbusinessActionConfigurations); err != nil {
			log.Errorf("invalid --action-configurations: %s", err.Error())
			return
		}
	}
	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessDataAccessorId) > 0 {
		input.DataAccessorId = aws.String(_qbusinessDataAccessorId)
	}
	if len(_qbusinessAuthenticationDetail) > 0 {
		if err := assignInputField(input, "AuthenticationDetail", _qbusinessAuthenticationDetail); err != nil {
			log.Errorf("invalid --authentication-detail: %s", err.Error())
			return
		}
	}
	if len(_qbusinessDisplayName) > 0 {
		input.DisplayName = aws.String(_qbusinessDisplayName)
	}

	if resp, err := client.UpdateDataAccessor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing Amazon Q Business data source connector.
func qbusiness_UpdateDataSource(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.UpdateDataSourceInput{
		// ApplicationId: *string, // Required
		// DataSourceId: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessDataSourceId) > 0 {
		input.DataSourceId = aws.String(_qbusinessDataSourceId)
	}
	if len(_qbusinessIndexId) > 0 {
		input.IndexId = aws.String(_qbusinessIndexId)
	}
	if len(_qbusinessConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _qbusinessConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessDescription) > 0 {
		input.Description = aws.String(_qbusinessDescription)
	}
	if len(_qbusinessDisplayName) > 0 {
		input.DisplayName = aws.String(_qbusinessDisplayName)
	}
	if len(_qbusinessDocumentEnrichmentConfiguration) > 0 {
		if err := assignInputField(input, "DocumentEnrichmentConfiguration", _qbusinessDocumentEnrichmentConfiguration); err != nil {
			log.Errorf("invalid --document-enrichment-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessMediaExtractionConfiguration) > 0 {
		if err := assignInputField(input, "MediaExtractionConfiguration", _qbusinessMediaExtractionConfiguration); err != nil {
			log.Errorf("invalid --media-extraction-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessRoleArn) > 0 {
		input.RoleArn = aws.String(_qbusinessRoleArn)
	}
	if len(_qbusinessSyncSchedule) > 0 {
		input.SyncSchedule = aws.String(_qbusinessSyncSchedule)
	}
	if len(_qbusinessVpcConfiguration) > 0 {
		if err := assignInputField(input, "VpcConfiguration", _qbusinessVpcConfiguration); err != nil {
			log.Errorf("invalid --vpc-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Amazon Q Business index.
func qbusiness_UpdateIndex(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.UpdateIndexInput{
		// ApplicationId: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessIndexId) > 0 {
		input.IndexId = aws.String(_qbusinessIndexId)
	}
	if len(_qbusinessCapacityConfiguration) > 0 {
		if err := assignInputField(input, "CapacityConfiguration", _qbusinessCapacityConfiguration); err != nil {
			log.Errorf("invalid --capacity-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessDescription) > 0 {
		input.Description = aws.String(_qbusinessDescription)
	}
	if len(_qbusinessDisplayName) > 0 {
		input.DisplayName = aws.String(_qbusinessDisplayName)
	}
	if len(_qbusinessDocumentAttributeConfigurations) > 0 {
		if err := assignInputField(input, "DocumentAttributeConfigurations", _qbusinessDocumentAttributeConfigurations); err != nil {
			log.Errorf("invalid --document-attribute-configurations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Amazon Q Business plugin.
func qbusiness_UpdatePlugin(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.UpdatePluginInput{
		// ApplicationId: *string, // Required
		// PluginId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessPluginId) > 0 {
		input.PluginId = aws.String(_qbusinessPluginId)
	}
	if len(_qbusinessAuthConfiguration) > 0 {
		if err := assignInputField(input, "AuthConfiguration", _qbusinessAuthConfiguration); err != nil {
			log.Errorf("invalid --auth-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessCustomPluginConfiguration) > 0 {
		if err := assignInputField(input, "CustomPluginConfiguration", _qbusinessCustomPluginConfiguration); err != nil {
			log.Errorf("invalid --custom-plugin-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessDisplayName) > 0 {
		input.DisplayName = aws.String(_qbusinessDisplayName)
	}
	if len(_qbusinessServerUrl) > 0 {
		input.ServerUrl = aws.String(_qbusinessServerUrl)
	}
	if len(_qbusinessState) > 0 {
		if err := assignInputField(input, "State", _qbusinessState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePlugin(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the retriever used for your Amazon Q Business application.
func qbusiness_UpdateRetriever(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.UpdateRetrieverInput{
		// ApplicationId: *string, // Required
		// RetrieverId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessRetrieverId) > 0 {
		input.RetrieverId = aws.String(_qbusinessRetrieverId)
	}
	if len(_qbusinessConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _qbusinessConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessDisplayName) > 0 {
		input.DisplayName = aws.String(_qbusinessDisplayName)
	}
	if len(_qbusinessRoleArn) > 0 {
		input.RoleArn = aws.String(_qbusinessRoleArn)
	}

	if resp, err := client.UpdateRetriever(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the pricing tier for an Amazon Q Business subscription. Upgrades are
// instant. Downgrades apply at the start of the next month. Subscription tier
// determines feature access for the user. For more information on subscriptions
// and pricing tiers, see [Amazon Q Business pricing].
//
// [Amazon Q Business pricing]: https://aws.amazon.com/q/business/pricing/
func qbusiness_UpdateSubscription(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.UpdateSubscriptionInput{
		// ApplicationId: *string, // Required
		// SubscriptionId: *string, // Required
		// Type: types.SubscriptionType, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessSubscriptionId) > 0 {
		input.SubscriptionId = aws.String(_qbusinessSubscriptionId)
	}
	if len(_qbusinessType) > 0 {
		if err := assignInputField(input, "Type", _qbusinessType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a information associated with a user id.
func qbusiness_UpdateUser(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.UpdateUserInput{
		// ApplicationId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessUserId) > 0 {
		input.UserId = aws.String(_qbusinessUserId)
	}
	if len(_qbusinessUserAliasesToDelete) > 0 {
		if err := assignInputField(input, "UserAliasesToDelete", _qbusinessUserAliasesToDelete); err != nil {
			log.Errorf("invalid --user-aliases-to-delete: %s", err.Error())
			return
		}
	}
	if len(_qbusinessUserAliasesToUpdate) > 0 {
		if err := assignInputField(input, "UserAliasesToUpdate", _qbusinessUserAliasesToUpdate); err != nil {
			log.Errorf("invalid --user-aliases-to-update: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Amazon Q Business web experience.
func qbusiness_UpdateWebExperience(cfg aws.Config, client *qbusiness.Client) {
	input := &qbusiness.UpdateWebExperienceInput{
		// ApplicationId: *string, // Required
		// WebExperienceId: *string, // Required
	}

	if len(_qbusinessApplicationId) > 0 {
		input.ApplicationId = aws.String(_qbusinessApplicationId)
	}
	if len(_qbusinessWebExperienceId) > 0 {
		input.WebExperienceId = aws.String(_qbusinessWebExperienceId)
	}
	if len(_qbusinessAuthenticationConfiguration) > 0 {
		if err := assignInputField(input, "AuthenticationConfiguration", _qbusinessAuthenticationConfiguration); err != nil {
			log.Errorf("invalid --authentication-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessBrowserExtensionConfiguration) > 0 {
		if err := assignInputField(input, "BrowserExtensionConfiguration", _qbusinessBrowserExtensionConfiguration); err != nil {
			log.Errorf("invalid --browser-extension-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessCustomizationConfiguration) > 0 {
		if err := assignInputField(input, "CustomizationConfiguration", _qbusinessCustomizationConfiguration); err != nil {
			log.Errorf("invalid --customization-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessIdentityProviderConfiguration) > 0 {
		if err := assignInputField(input, "IdentityProviderConfiguration", _qbusinessIdentityProviderConfiguration); err != nil {
			log.Errorf("invalid --identity-provider-configuration: %s", err.Error())
			return
		}
	}
	if len(_qbusinessOrigins) > 0 {
		input.Origins = append([]string(nil), _qbusinessOrigins...)
	}
	if len(_qbusinessRoleArn) > 0 {
		input.RoleArn = aws.String(_qbusinessRoleArn)
	}
	if len(_qbusinessSamplePromptsControlMode) > 0 {
		if err := assignInputField(input, "SamplePromptsControlMode", _qbusinessSamplePromptsControlMode); err != nil {
			log.Errorf("invalid --sample-prompts-control-mode: %s", err.Error())
			return
		}
	}
	if len(_qbusinessSubtitle) > 0 {
		input.Subtitle = aws.String(_qbusinessSubtitle)
	}
	if len(_qbusinessTitle) > 0 {
		input.Title = aws.String(_qbusinessTitle)
	}
	if len(_qbusinessWelcomeMessage) > 0 {
		input.WelcomeMessage = aws.String(_qbusinessWelcomeMessage)
	}

	if resp, err := client.UpdateWebExperience(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_qbusinessCmd)
	_qbusinessCmd.Flags().SortFlags = false

	_qbusinessCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_qbusinessCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_qbusinessCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_qbusinessCmd.Flags().StringVarP(&_qbusinessActionConfigurations, "action-configurations", "", "", "Action Configurations")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessActionExecution, "action-execution", "", "", "Action Execution")
	_qbusinessCmd.Flags().StringSliceVarP(&_qbusinessActions, "actions", "", nil, "Actions")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessApplicationId, "application-id", "", "", "Application ID")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessAttachmentId, "attachment-id", "", "", "Attachment ID")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessAttachments, "attachments", "", "", "Attachments")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessAttachmentsConfiguration, "attachments-configuration", "", "", "Attachments Configuration")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessAttributeFilter, "attribute-filter", "", "", "Attribute Filter")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessAuthChallengeResponse, "auth-challenge-response", "", "", "Auth Challenge Response")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessAuthConfiguration, "auth-configuration", "", "", "Auth Configuration")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessAuthenticationConfiguration, "authentication-configuration", "", "", "Authentication Configuration")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessAuthenticationDetail, "authentication-detail", "", "", "Authentication Detail")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessAutoSubscriptionConfiguration, "auto-subscription-configuration", "", "", "Auto Subscription Configuration")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessBlockedPhrasesConfigurationUpdate, "blocked-phrases-configuration-update", "", "", "Blocked Phrases Configuration Update")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessBrowserExtensionConfiguration, "browser-extension-configuration", "", "", "Browser Extension Configuration")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessCapacityConfiguration, "capacity-configuration", "", "", "Capacity Configuration")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessChatMode, "chat-mode", "", "", "Chat Mode")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessChatModeConfiguration, "chat-mode-configuration", "", "", "Chat Mode Configuration")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessChatResponseConfigurationId, "chat-response-configuration-id", "", "", "Chat Response Configuration ID")
	_qbusinessCmd.Flags().StringSliceVarP(&_qbusinessClientIdsForOIDC, "client-ids-for-oidc", "", nil, "Client Ids For OIDC")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessClientToken, "client-token", "", "", "Client Token")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessConditions, "conditions", "", "", "Conditions")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessConfiguration, "configuration", "", "", "Configuration")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessContentSource, "content-source", "", "", "Content Source")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessConversationId, "conversation-id", "", "", "Conversation ID")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessCreatorModeConfiguration, "creator-mode-configuration", "", "", "Creator Mode Configuration")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessCustomPluginConfiguration, "custom-plugin-configuration", "", "", "Custom Plugin Configuration")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessCustomizationConfiguration, "customization-configuration", "", "", "Customization Configuration")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessDataAccessorId, "data-accessor-id", "", "", "Data Accessor ID")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessDataSourceId, "data-source-id", "", "", "Data Source ID")
	_qbusinessCmd.Flags().StringSliceVarP(&_qbusinessDataSourceIds, "data-source-ids", "", nil, "Data Source Ids")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessDataSourceSyncId, "data-source-sync-id", "", "", "Data Source Sync ID")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessDescription, "description", "", "", "Description")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessDisplayName, "display-name", "", "", "Display Name")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessDocumentAttributeConfigurations, "document-attribute-configurations", "", "", "Document Attribute Configurations")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessDocumentEnrichmentConfiguration, "document-enrichment-configuration", "", "", "Document Enrichment Configuration")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessDocumentId, "document-id", "", "", "Document ID")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessDocuments, "documents", "", "", "Documents")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessEncryptionConfiguration, "encryption-configuration", "", "", "Encryption Configuration")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessEndTime, "end-time", "", "", "End Time")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessGroupMembers, "group-members", "", "", "Group Members")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessGroupName, "group-name", "", "", "Group Name")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessHallucinationReductionConfiguration, "hallucination-reduction-configuration", "", "", "Hallucination Reduction Configuration")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessIamIdentityProviderArn, "iam-identity-provider-arn", "", "", "IAM Identity Provider ARN")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessIdentityCenterInstanceArn, "identity-center-instance-arn", "", "", "Identity Center Instance ARN")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessIdentityProviderConfiguration, "identity-provider-configuration", "", "", "Identity Provider Configuration")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessIdentityType, "identity-type", "", "", "Identity Type")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessIndexId, "index-id", "", "", "Index ID")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessMaxResults, "max-results", "", "", "Max Results")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessMediaExtractionConfiguration, "media-extraction-configuration", "", "", "Media Extraction Configuration")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessMediaId, "media-id", "", "", "Media ID")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessMessageCopiedAt, "message-copied-at", "", "", "Message Copied At")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessMessageId, "message-id", "", "", "Message ID")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessMessageUsefulness, "message-usefulness", "", "", "Message Usefulness")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessNextToken, "next-token", "", "", "Next Token")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessOrchestrationConfiguration, "orchestration-configuration", "", "", "Orchestration Configuration")
	_qbusinessCmd.Flags().StringSliceVarP(&_qbusinessOrigins, "origins", "", nil, "Origins")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessOutputFormat, "output-format", "", "", "Output Format")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessParentMessageId, "parent-message-id", "", "", "Parent Message ID")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessPersonalizationConfiguration, "personalization-configuration", "", "", "Personalization Configuration")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessPluginId, "plugin-id", "", "", "Plugin ID")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessPluginType, "plugin-type", "", "", "Plugin Type")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessPrincipal, "principal", "", "", "Principal")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessQAppsConfiguration, "qapps-configuration", "", "", "Qapps Configuration")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessQueryText, "query-text", "", "", "Query Text")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessQuickSightConfiguration, "quicksight-configuration", "", "", "Quicksight Configuration")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessResourceARN, "resource-arn", "", "", "Resource ARN")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessResponseConfigurations, "response-configurations", "", "", "Response Configurations")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessResponseScope, "response-scope", "", "", "Response Scope")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessRetrieverId, "retriever-id", "", "", "Retriever ID")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessRoleArn, "role-arn", "", "", "Role ARN")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessSamplePromptsControlMode, "sample-prompts-control-mode", "", "", "Sample Prompts Control Mode")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessServerUrl, "server-url", "", "", "Server URL")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessSessionDurationInMinutes, "session-duration-in-minutes", "", "", "Session Duration In Minutes")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessStartTime, "start-time", "", "", "Start Time")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessState, "state", "", "", "State")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessStatementId, "statement-id", "", "", "Statement ID")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessStatusFilter, "status-filter", "", "", "Status Filter")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessSubscriptionId, "subscription-id", "", "", "Subscription ID")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessSubtitle, "subtitle", "", "", "Subtitle")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessSyncSchedule, "sync-schedule", "", "", "Sync Schedule")
	_qbusinessCmd.Flags().StringSliceVarP(&_qbusinessTagKeys, "tag-keys", "", nil, "Tag Keys")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessTags, "tags", "", "", "Tags")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessTitle, "title", "", "", "Title")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessTopicConfigurationsToCreateOrUpdate, "topic-configurations-to-create-or-update", "", "", "Topic Configurations To Create Or Update")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessTopicConfigurationsToDelete, "topic-configurations-to-delete", "", "", "Topic Configurations To Delete")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessType, "type", "", "", "Type")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessUpdatedEarlierThan, "updated-earlier-than", "", "", "Updated Earlier Than")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessUserAliases, "user-aliases", "", "", "User Aliases")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessUserAliasesToDelete, "user-aliases-to-delete", "", "", "User Aliases To Delete")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessUserAliasesToUpdate, "user-aliases-to-update", "", "", "User Aliases To Update")
	_qbusinessCmd.Flags().StringSliceVarP(&_qbusinessUserGroups, "user-groups", "", nil, "User Groups")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessUserId, "user-id", "", "", "User ID")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessUserMessage, "user-message", "", "", "User Message")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessVpcConfiguration, "vpc-configuration", "", "", "VPC Configuration")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessWebExperienceId, "web-experience-id", "", "", "Web Experience ID")
	_qbusinessCmd.Flags().StringVarP(&_qbusinessWelcomeMessage, "welcome-message", "", "", "Welcome Message")

	_qbusinessCmd.Flags().BoolVarP(&_qbusinessAssociatePermission, "associate-permission", "", false, "Associate Permission")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessBatchDeleteDocument, "batch-delete-document", "", false, "Batch Delete Document")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessBatchPutDocument, "batch-put-document", "", false, "Batch Put Document")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessCancelSubscription, "cancel-subscription", "", false, "Cancel Subscription")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessChat, "chat", "", false, "Chat")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessChatSync, "chat-sync", "", false, "Chat Sync")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessCheckDocumentAccess, "check-document-access", "", false, "Check Document Access")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessCreateAnonymousWebExperienceUrl, "create-anonymous-web-experience-url", "", false, "Create Anonymous Web Experience URL")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessCreateApplication, "create-application", "", false, "Create Application")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessCreateChatResponseConfiguration, "create-chat-response-configuration", "", false, "Create Chat Response Configuration")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessCreateDataAccessor, "create-data-accessor", "", false, "Create Data Accessor")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessCreateDataSource, "create-data-source", "", false, "Create Data Source")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessCreateIndex, "create-index", "", false, "Create Index")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessCreatePlugin, "create-plugin", "", false, "Create Plugin")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessCreateRetriever, "create-retriever", "", false, "Create Retriever")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessCreateSubscription, "create-subscription", "", false, "Create Subscription")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessCreateUser, "create-user", "", false, "Create User")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessCreateWebExperience, "create-web-experience", "", false, "Create Web Experience")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessDeleteApplication, "delete-application", "", false, "Delete Application")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessDeleteAttachment, "delete-attachment", "", false, "Delete Attachment")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessDeleteChatControlsConfiguration, "delete-chat-controls-configuration", "", false, "Delete Chat Controls Configuration")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessDeleteChatResponseConfiguration, "delete-chat-response-configuration", "", false, "Delete Chat Response Configuration")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessDeleteConversation, "delete-conversation", "", false, "Delete Conversation")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessDeleteDataAccessor, "delete-data-accessor", "", false, "Delete Data Accessor")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessDeleteDataSource, "delete-data-source", "", false, "Delete Data Source")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessDeleteGroup, "delete-group", "", false, "Delete Group")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessDeleteIndex, "delete-index", "", false, "Delete Index")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessDeletePlugin, "delete-plugin", "", false, "Delete Plugin")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessDeleteRetriever, "delete-retriever", "", false, "Delete Retriever")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessDeleteUser, "delete-user", "", false, "Delete User")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessDeleteWebExperience, "delete-web-experience", "", false, "Delete Web Experience")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessDisassociatePermission, "disassociate-permission", "", false, "Disassociate Permission")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessGetApplication, "get-application", "", false, "Get Application")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessGetChatControlsConfiguration, "get-chat-controls-configuration", "", false, "Get Chat Controls Configuration")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessGetChatResponseConfiguration, "get-chat-response-configuration", "", false, "Get Chat Response Configuration")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessGetDataAccessor, "get-data-accessor", "", false, "Get Data Accessor")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessGetDataSource, "get-data-source", "", false, "Get Data Source")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessGetDocumentContent, "get-document-content", "", false, "Get Document Content")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessGetGroup, "get-group", "", false, "Get Group")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessGetIndex, "get-index", "", false, "Get Index")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessGetMedia, "get-media", "", false, "Get Media")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessGetPlugin, "get-plugin", "", false, "Get Plugin")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessGetPolicy, "get-policy", "", false, "Get Policy")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessGetRetriever, "get-retriever", "", false, "Get Retriever")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessGetUser, "get-user", "", false, "Get User")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessGetWebExperience, "get-web-experience", "", false, "Get Web Experience")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessListApplications, "list-applications", "", false, "List Applications")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessListAttachments, "list-attachments", "", false, "List Attachments")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessListChatResponseConfigurations, "list-chat-response-configurations", "", false, "List Chat Response Configurations")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessListConversations, "list-conversations", "", false, "List Conversations")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessListDataAccessors, "list-data-accessors", "", false, "List Data Accessors")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessListDataSourceSyncJobs, "list-data-source-sync-jobs", "", false, "List Data Source Sync Jobs")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessListDataSources, "list-data-sources", "", false, "List Data Sources")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessListDocuments, "list-documents", "", false, "List Documents")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessListGroups, "list-groups", "", false, "List Groups")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessListIndices, "list-indices", "", false, "List Indices")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessListMessages, "list-messages", "", false, "List Messages")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessListPluginActions, "list-plugin-actions", "", false, "List Plugin Actions")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessListPluginTypeActions, "list-plugin-type-actions", "", false, "List Plugin Type Actions")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessListPluginTypeMetadata, "list-plugin-type-metadata", "", false, "List Plugin Type Metadata")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessListPlugins, "list-plugins", "", false, "List Plugins")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessListRetrievers, "list-retrievers", "", false, "List Retrievers")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessListSubscriptions, "list-subscriptions", "", false, "List Subscriptions")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessListWebExperiences, "list-web-experiences", "", false, "List Web Experiences")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessPutFeedback, "put-feedback", "", false, "Put Feedback")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessPutGroup, "put-group", "", false, "Put Group")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessSearchRelevantContent, "search-relevant-content", "", false, "Search Relevant Content")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessStartDataSourceSyncJob, "start-data-source-sync-job", "", false, "Start Data Source Sync Job")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessStopDataSourceSyncJob, "stop-data-source-sync-job", "", false, "Stop Data Source Sync Job")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessTagResource, "tag-resource", "", false, "Tag Resource")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessUntagResource, "untag-resource", "", false, "Untag Resource")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessUpdateApplication, "update-application", "", false, "Update Application")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessUpdateChatControlsConfiguration, "update-chat-controls-configuration", "", false, "Update Chat Controls Configuration")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessUpdateChatResponseConfiguration, "update-chat-response-configuration", "", false, "Update Chat Response Configuration")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessUpdateDataAccessor, "update-data-accessor", "", false, "Update Data Accessor")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessUpdateDataSource, "update-data-source", "", false, "Update Data Source")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessUpdateIndex, "update-index", "", false, "Update Index")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessUpdatePlugin, "update-plugin", "", false, "Update Plugin")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessUpdateRetriever, "update-retriever", "", false, "Update Retriever")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessUpdateSubscription, "update-subscription", "", false, "Update Subscription")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessUpdateUser, "update-user", "", false, "Update User")
	_qbusinessCmd.Flags().BoolVarP(&_qbusinessUpdateWebExperience, "update-web-experience", "", false, "Update Web Experience")

}
