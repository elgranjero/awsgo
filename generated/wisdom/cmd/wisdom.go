package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wisdom"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// wisdomCmd represents the wisdom command
var _wisdomCmd = &cobra.Command{
	Use:   "wisdom",
	Short: "AWS wisdom CLI",
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
		client := wisdom.NewFromConfig(cfg)
		if _wisdomCreateAssistant {
			wisdom_CreateAssistant(cfg, client)
			return
		}
		if _wisdomCreateAssistantAssociation {
			wisdom_CreateAssistantAssociation(cfg, client)
			return
		}
		if _wisdomCreateContent {
			wisdom_CreateContent(cfg, client)
			return
		}
		if _wisdomCreateKnowledgeBase {
			wisdom_CreateKnowledgeBase(cfg, client)
			return
		}
		if _wisdomCreateQuickResponse {
			wisdom_CreateQuickResponse(cfg, client)
			return
		}
		if _wisdomCreateSession {
			wisdom_CreateSession(cfg, client)
			return
		}
		if _wisdomDeleteAssistant {
			wisdom_DeleteAssistant(cfg, client)
			return
		}
		if _wisdomDeleteAssistantAssociation {
			wisdom_DeleteAssistantAssociation(cfg, client)
			return
		}
		if _wisdomDeleteContent {
			wisdom_DeleteContent(cfg, client)
			return
		}
		if _wisdomDeleteImportJob {
			wisdom_DeleteImportJob(cfg, client)
			return
		}
		if _wisdomDeleteKnowledgeBase {
			wisdom_DeleteKnowledgeBase(cfg, client)
			return
		}
		if _wisdomDeleteQuickResponse {
			wisdom_DeleteQuickResponse(cfg, client)
			return
		}
		if _wisdomGetAssistant {
			wisdom_GetAssistant(cfg, client)
			return
		}
		if _wisdomGetAssistantAssociation {
			wisdom_GetAssistantAssociation(cfg, client)
			return
		}
		if _wisdomGetContent {
			wisdom_GetContent(cfg, client)
			return
		}
		if _wisdomGetContentSummary {
			wisdom_GetContentSummary(cfg, client)
			return
		}
		if _wisdomGetImportJob {
			wisdom_GetImportJob(cfg, client)
			return
		}
		if _wisdomGetKnowledgeBase {
			wisdom_GetKnowledgeBase(cfg, client)
			return
		}
		if _wisdomGetQuickResponse {
			wisdom_GetQuickResponse(cfg, client)
			return
		}
		if _wisdomGetRecommendations {
			wisdom_GetRecommendations(cfg, client)
			return
		}
		if _wisdomGetSession {
			wisdom_GetSession(cfg, client)
			return
		}
		if _wisdomListAssistantAssociations {
			wisdom_ListAssistantAssociations(cfg, client)
			return
		}
		if _wisdomListAssistants {
			wisdom_ListAssistants(cfg, client)
			return
		}
		if _wisdomListContents {
			wisdom_ListContents(cfg, client)
			return
		}
		if _wisdomListImportJobs {
			wisdom_ListImportJobs(cfg, client)
			return
		}
		if _wisdomListKnowledgeBases {
			wisdom_ListKnowledgeBases(cfg, client)
			return
		}
		if _wisdomListQuickResponses {
			wisdom_ListQuickResponses(cfg, client)
			return
		}
		if _wisdomListTagsForResource {
			wisdom_ListTagsForResource(cfg, client)
			return
		}
		if _wisdomNotifyRecommendationsReceived {
			wisdom_NotifyRecommendationsReceived(cfg, client)
			return
		}
		if _wisdomQueryAssistant {
			wisdom_QueryAssistant(cfg, client)
			return
		}
		if _wisdomRemoveKnowledgeBaseTemplateUri {
			wisdom_RemoveKnowledgeBaseTemplateUri(cfg, client)
			return
		}
		if _wisdomSearchContent {
			wisdom_SearchContent(cfg, client)
			return
		}
		if _wisdomSearchQuickResponses {
			wisdom_SearchQuickResponses(cfg, client)
			return
		}
		if _wisdomSearchSessions {
			wisdom_SearchSessions(cfg, client)
			return
		}
		if _wisdomStartContentUpload {
			wisdom_StartContentUpload(cfg, client)
			return
		}
		if _wisdomStartImportJob {
			wisdom_StartImportJob(cfg, client)
			return
		}
		if _wisdomTagResource {
			wisdom_TagResource(cfg, client)
			return
		}
		if _wisdomUntagResource {
			wisdom_UntagResource(cfg, client)
			return
		}
		if _wisdomUpdateContent {
			wisdom_UpdateContent(cfg, client)
			return
		}
		if _wisdomUpdateKnowledgeBaseTemplateUri {
			wisdom_UpdateKnowledgeBaseTemplateUri(cfg, client)
			return
		}
		if _wisdomUpdateQuickResponse {
			wisdom_UpdateQuickResponse(cfg, client)
			return
		}

	},
}

var (
	_wisdomCreateAssistant                bool
	_wisdomCreateAssistantAssociation     bool
	_wisdomCreateContent                  bool
	_wisdomCreateKnowledgeBase            bool
	_wisdomCreateQuickResponse            bool
	_wisdomCreateSession                  bool
	_wisdomDeleteAssistant                bool
	_wisdomDeleteAssistantAssociation     bool
	_wisdomDeleteContent                  bool
	_wisdomDeleteImportJob                bool
	_wisdomDeleteKnowledgeBase            bool
	_wisdomDeleteQuickResponse            bool
	_wisdomGetAssistant                   bool
	_wisdomGetAssistantAssociation        bool
	_wisdomGetContent                     bool
	_wisdomGetContentSummary              bool
	_wisdomGetImportJob                   bool
	_wisdomGetKnowledgeBase               bool
	_wisdomGetQuickResponse               bool
	_wisdomGetRecommendations             bool
	_wisdomGetSession                     bool
	_wisdomListAssistantAssociations      bool
	_wisdomListAssistants                 bool
	_wisdomListContents                   bool
	_wisdomListImportJobs                 bool
	_wisdomListKnowledgeBases             bool
	_wisdomListQuickResponses             bool
	_wisdomListTagsForResource            bool
	_wisdomNotifyRecommendationsReceived  bool
	_wisdomQueryAssistant                 bool
	_wisdomRemoveKnowledgeBaseTemplateUri bool
	_wisdomSearchContent                  bool
	_wisdomSearchQuickResponses           bool
	_wisdomSearchSessions                 bool
	_wisdomStartContentUpload             bool
	_wisdomStartImportJob                 bool
	_wisdomTagResource                    bool
	_wisdomUntagResource                  bool
	_wisdomUpdateContent                  bool
	_wisdomUpdateKnowledgeBaseTemplateUri bool
	_wisdomUpdateQuickResponse            bool

	_wisdomAssistantAssociationId            string
	_wisdomAssistantId                       string
	_wisdomAssociation                       string
	_wisdomAssociationType                   string
	_wisdomAttributes                        string
	_wisdomChannels                          []string
	_wisdomClientToken                       string
	_wisdomContent                           string
	_wisdomContentId                         string
	_wisdomContentType                       string
	_wisdomDescription                       string
	_wisdomExternalSourceConfiguration       string
	_wisdomGroupingConfiguration             string
	_wisdomImportJobId                       string
	_wisdomImportJobType                     string
	_wisdomIsActive                          string
	_wisdomKnowledgeBaseId                   string
	_wisdomKnowledgeBaseType                 string
	_wisdomLanguage                          string
	_wisdomMaxResults                        string
	_wisdomMetadata                          string
	_wisdomName                              string
	_wisdomNextToken                         string
	_wisdomOverrideLinkOutUri                string
	_wisdomPresignedUrlTimeToLive            string
	_wisdomQueryText                         string
	_wisdomQuickResponseId                   string
	_wisdomRecommendationIds                 []string
	_wisdomRemoveDescription                 string
	_wisdomRemoveGroupingConfiguration       string
	_wisdomRemoveOverrideLinkOutUri          string
	_wisdomRemoveShortcutKey                 string
	_wisdomRenderingConfiguration            string
	_wisdomResourceArn                       string
	_wisdomRevisionId                        string
	_wisdomSearchExpression                  string
	_wisdomServerSideEncryptionConfiguration string
	_wisdomSessionId                         string
	_wisdomShortcutKey                       string
	_wisdomSourceConfiguration               string
	_wisdomTagKeys                           []string
	_wisdomTags                              string
	_wisdomTemplateUri                       string
	_wisdomTitle                             string
	_wisdomType                              string
	_wisdomUploadId                          string
	_wisdomWaitTimeSeconds                   string
)

// Creates an Amazon Connect Wisdom assistant.
func wisdom_CreateAssistant(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.CreateAssistantInput{
		// Name: *string, // Required
		// Type: types.AssistantType, // Required
	}

	if len(_wisdomName) > 0 {
		input.Name = aws.String(_wisdomName)
	}
	if len(_wisdomType) > 0 {
		if err := assignInputField(input, "Type", _wisdomType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_wisdomClientToken) > 0 {
		input.ClientToken = aws.String(_wisdomClientToken)
	}
	if len(_wisdomDescription) > 0 {
		input.Description = aws.String(_wisdomDescription)
	}
	if len(_wisdomServerSideEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "ServerSideEncryptionConfiguration", _wisdomServerSideEncryptionConfiguration); err != nil {
			log.Errorf("invalid --server-side-encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_wisdomTags) > 0 {
		if err := assignInputField(input, "Tags", _wisdomTags); err != nil {
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

// Creates an association between an Amazon Connect Wisdom assistant and another
// resource. Currently, the only supported association is with a knowledge base. An
// assistant can have only a single association.
func wisdom_CreateAssistantAssociation(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.CreateAssistantAssociationInput{
		// AssistantId: *string, // Required
		// Association: types.AssistantAssociationInputData, // Required
		// AssociationType: types.AssociationType, // Required
	}

	if len(_wisdomAssistantId) > 0 {
		input.AssistantId = aws.String(_wisdomAssistantId)
	}
	if len(_wisdomAssociation) > 0 {
		if err := assignInputField(input, "Association", _wisdomAssociation); err != nil {
			log.Errorf("invalid --association: %s", err.Error())
			return
		}
	}
	if len(_wisdomAssociationType) > 0 {
		if err := assignInputField(input, "AssociationType", _wisdomAssociationType); err != nil {
			log.Errorf("invalid --association-type: %s", err.Error())
			return
		}
	}
	if len(_wisdomClientToken) > 0 {
		input.ClientToken = aws.String(_wisdomClientToken)
	}
	if len(_wisdomTags) > 0 {
		if err := assignInputField(input, "Tags", _wisdomTags); err != nil {
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

// Creates Wisdom content. Before to calling this API, use [StartContentUpload] to upload an asset.
//
// [StartContentUpload]: https://docs.aws.amazon.com/wisdom/latest/APIReference/API_StartContentUpload.html
func wisdom_CreateContent(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.CreateContentInput{
		// KnowledgeBaseId: *string, // Required
		// Name: *string, // Required
		// UploadId: *string, // Required
	}

	if len(_wisdomKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_wisdomKnowledgeBaseId)
	}
	if len(_wisdomName) > 0 {
		input.Name = aws.String(_wisdomName)
	}
	if len(_wisdomUploadId) > 0 {
		input.UploadId = aws.String(_wisdomUploadId)
	}
	if len(_wisdomClientToken) > 0 {
		input.ClientToken = aws.String(_wisdomClientToken)
	}
	if len(_wisdomMetadata) > 0 {
		if err := assignInputField(input, "Metadata", _wisdomMetadata); err != nil {
			log.Errorf("invalid --metadata: %s", err.Error())
			return
		}
	}
	if len(_wisdomOverrideLinkOutUri) > 0 {
		input.OverrideLinkOutUri = aws.String(_wisdomOverrideLinkOutUri)
	}
	if len(_wisdomTags) > 0 {
		if err := assignInputField(input, "Tags", _wisdomTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_wisdomTitle) > 0 {
		input.Title = aws.String(_wisdomTitle)
	}

	if resp, err := client.CreateContent(context.TODO(), input); err != nil {
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
// [DeleteKnowledgeBase]: https://docs.aws.amazon.com/wisdom/latest/APIReference/API_DeleteKnowledgeBase.html
// [DeleteDataIntegration]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_DeleteDataIntegration.html
// [CreateDataIntegration]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_CreateDataIntegration.html
func wisdom_CreateKnowledgeBase(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.CreateKnowledgeBaseInput{
		// KnowledgeBaseType: types.KnowledgeBaseType, // Required
		// Name: *string, // Required
	}

	if len(_wisdomKnowledgeBaseType) > 0 {
		if err := assignInputField(input, "KnowledgeBaseType", _wisdomKnowledgeBaseType); err != nil {
			log.Errorf("invalid --knowledge-base-type: %s", err.Error())
			return
		}
	}
	if len(_wisdomName) > 0 {
		input.Name = aws.String(_wisdomName)
	}
	if len(_wisdomClientToken) > 0 {
		input.ClientToken = aws.String(_wisdomClientToken)
	}
	if len(_wisdomDescription) > 0 {
		input.Description = aws.String(_wisdomDescription)
	}
	if len(_wisdomRenderingConfiguration) > 0 {
		if err := assignInputField(input, "RenderingConfiguration", _wisdomRenderingConfiguration); err != nil {
			log.Errorf("invalid --rendering-configuration: %s", err.Error())
			return
		}
	}
	if len(_wisdomServerSideEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "ServerSideEncryptionConfiguration", _wisdomServerSideEncryptionConfiguration); err != nil {
			log.Errorf("invalid --server-side-encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_wisdomSourceConfiguration) > 0 {
		if err := assignInputField(input, "SourceConfiguration", _wisdomSourceConfiguration); err != nil {
			log.Errorf("invalid --source-configuration: %s", err.Error())
			return
		}
	}
	if len(_wisdomTags) > 0 {
		if err := assignInputField(input, "Tags", _wisdomTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
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

// Creates a Wisdom quick response.
func wisdom_CreateQuickResponse(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.CreateQuickResponseInput{
		// Content: types.QuickResponseDataProvider, // Required
		// KnowledgeBaseId: *string, // Required
		// Name: *string, // Required
	}

	if len(_wisdomContent) > 0 {
		if err := assignInputField(input, "Content", _wisdomContent); err != nil {
			log.Errorf("invalid --content: %s", err.Error())
			return
		}
	}
	if len(_wisdomKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_wisdomKnowledgeBaseId)
	}
	if len(_wisdomName) > 0 {
		input.Name = aws.String(_wisdomName)
	}
	if len(_wisdomChannels) > 0 {
		input.Channels = append([]string(nil), _wisdomChannels...)
	}
	if len(_wisdomClientToken) > 0 {
		input.ClientToken = aws.String(_wisdomClientToken)
	}
	if len(_wisdomContentType) > 0 {
		input.ContentType = aws.String(_wisdomContentType)
	}
	if len(_wisdomDescription) > 0 {
		input.Description = aws.String(_wisdomDescription)
	}
	if len(_wisdomGroupingConfiguration) > 0 {
		if err := assignInputField(input, "GroupingConfiguration", _wisdomGroupingConfiguration); err != nil {
			log.Errorf("invalid --grouping-configuration: %s", err.Error())
			return
		}
	}
	if len(_wisdomIsActive) > 0 {
		if err := assignInputField(input, "IsActive", _wisdomIsActive); err != nil {
			log.Errorf("invalid --is-active: %s", err.Error())
			return
		}
	}
	if len(_wisdomLanguage) > 0 {
		input.Language = aws.String(_wisdomLanguage)
	}
	if len(_wisdomShortcutKey) > 0 {
		input.ShortcutKey = aws.String(_wisdomShortcutKey)
	}
	if len(_wisdomTags) > 0 {
		if err := assignInputField(input, "Tags", _wisdomTags); err != nil {
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
// recommendations. Amazon Connect creates a new Wisdom session for each contact on
// which Wisdom is enabled.
func wisdom_CreateSession(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.CreateSessionInput{
		// AssistantId: *string, // Required
		// Name: *string, // Required
	}

	if len(_wisdomAssistantId) > 0 {
		input.AssistantId = aws.String(_wisdomAssistantId)
	}
	if len(_wisdomName) > 0 {
		input.Name = aws.String(_wisdomName)
	}
	if len(_wisdomClientToken) > 0 {
		input.ClientToken = aws.String(_wisdomClientToken)
	}
	if len(_wisdomDescription) > 0 {
		input.Description = aws.String(_wisdomDescription)
	}
	if len(_wisdomTags) > 0 {
		if err := assignInputField(input, "Tags", _wisdomTags); err != nil {
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

// Deletes an assistant.
func wisdom_DeleteAssistant(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.DeleteAssistantInput{
		// AssistantId: *string, // Required
	}

	if len(_wisdomAssistantId) > 0 {
		input.AssistantId = aws.String(_wisdomAssistantId)
	}

	if resp, err := client.DeleteAssistant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an assistant association.
func wisdom_DeleteAssistantAssociation(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.DeleteAssistantAssociationInput{
		// AssistantAssociationId: *string, // Required
		// AssistantId: *string, // Required
	}

	if len(_wisdomAssistantAssociationId) > 0 {
		input.AssistantAssociationId = aws.String(_wisdomAssistantAssociationId)
	}
	if len(_wisdomAssistantId) > 0 {
		input.AssistantId = aws.String(_wisdomAssistantId)
	}

	if resp, err := client.DeleteAssistantAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the content.
func wisdom_DeleteContent(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.DeleteContentInput{
		// ContentId: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_wisdomContentId) > 0 {
		input.ContentId = aws.String(_wisdomContentId)
	}
	if len(_wisdomKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_wisdomKnowledgeBaseId)
	}

	if resp, err := client.DeleteContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the quick response import job.
func wisdom_DeleteImportJob(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.DeleteImportJobInput{
		// ImportJobId: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_wisdomImportJobId) > 0 {
		input.ImportJobId = aws.String(_wisdomImportJobId)
	}
	if len(_wisdomKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_wisdomKnowledgeBaseId)
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
func wisdom_DeleteKnowledgeBase(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.DeleteKnowledgeBaseInput{
		// KnowledgeBaseId: *string, // Required
	}

	if len(_wisdomKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_wisdomKnowledgeBaseId)
	}

	if resp, err := client.DeleteKnowledgeBase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a quick response.
func wisdom_DeleteQuickResponse(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.DeleteQuickResponseInput{
		// KnowledgeBaseId: *string, // Required
		// QuickResponseId: *string, // Required
	}

	if len(_wisdomKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_wisdomKnowledgeBaseId)
	}
	if len(_wisdomQuickResponseId) > 0 {
		input.QuickResponseId = aws.String(_wisdomQuickResponseId)
	}

	if resp, err := client.DeleteQuickResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an assistant.
func wisdom_GetAssistant(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.GetAssistantInput{
		// AssistantId: *string, // Required
	}

	if len(_wisdomAssistantId) > 0 {
		input.AssistantId = aws.String(_wisdomAssistantId)
	}

	if resp, err := client.GetAssistant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an assistant association.
func wisdom_GetAssistantAssociation(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.GetAssistantAssociationInput{
		// AssistantAssociationId: *string, // Required
		// AssistantId: *string, // Required
	}

	if len(_wisdomAssistantAssociationId) > 0 {
		input.AssistantAssociationId = aws.String(_wisdomAssistantAssociationId)
	}
	if len(_wisdomAssistantId) > 0 {
		input.AssistantId = aws.String(_wisdomAssistantId)
	}

	if resp, err := client.GetAssistantAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves content, including a pre-signed URL to download the content.
func wisdom_GetContent(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.GetContentInput{
		// ContentId: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_wisdomContentId) > 0 {
		input.ContentId = aws.String(_wisdomContentId)
	}
	if len(_wisdomKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_wisdomKnowledgeBaseId)
	}

	if resp, err := client.GetContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves summary information about the content.
func wisdom_GetContentSummary(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.GetContentSummaryInput{
		// ContentId: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_wisdomContentId) > 0 {
		input.ContentId = aws.String(_wisdomContentId)
	}
	if len(_wisdomKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_wisdomKnowledgeBaseId)
	}

	if resp, err := client.GetContentSummary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the started import job.
func wisdom_GetImportJob(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.GetImportJobInput{
		// ImportJobId: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_wisdomImportJobId) > 0 {
		input.ImportJobId = aws.String(_wisdomImportJobId)
	}
	if len(_wisdomKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_wisdomKnowledgeBaseId)
	}

	if resp, err := client.GetImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the knowledge base.
func wisdom_GetKnowledgeBase(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.GetKnowledgeBaseInput{
		// KnowledgeBaseId: *string, // Required
	}

	if len(_wisdomKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_wisdomKnowledgeBaseId)
	}

	if resp, err := client.GetKnowledgeBase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the quick response.
func wisdom_GetQuickResponse(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.GetQuickResponseInput{
		// KnowledgeBaseId: *string, // Required
		// QuickResponseId: *string, // Required
	}

	if len(_wisdomKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_wisdomKnowledgeBaseId)
	}
	if len(_wisdomQuickResponseId) > 0 {
		input.QuickResponseId = aws.String(_wisdomQuickResponseId)
	}

	if resp, err := client.GetQuickResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

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
// [QueryAssistant]: https://docs.aws.amazon.com/wisdom/latest/APIReference/API_QueryAssistant.html
// [NotifyRecommendationsReceived]: https://docs.aws.amazon.com/wisdom/latest/APIReference/API_NotifyRecommendationsReceived.html
func wisdom_GetRecommendations(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.GetRecommendationsInput{
		// AssistantId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_wisdomAssistantId) > 0 {
		input.AssistantId = aws.String(_wisdomAssistantId)
	}
	if len(_wisdomSessionId) > 0 {
		input.SessionId = aws.String(_wisdomSessionId)
	}
	if len(_wisdomMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wisdomMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wisdomWaitTimeSeconds) > 0 {
		if err := assignInputField(input, "WaitTimeSeconds", _wisdomWaitTimeSeconds); err != nil {
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
func wisdom_GetSession(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.GetSessionInput{
		// AssistantId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_wisdomAssistantId) > 0 {
		input.AssistantId = aws.String(_wisdomAssistantId)
	}
	if len(_wisdomSessionId) > 0 {
		input.SessionId = aws.String(_wisdomSessionId)
	}

	if resp, err := client.GetSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists information about assistant associations.
func wisdom_ListAssistantAssociations(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.ListAssistantAssociationsInput{
		// AssistantId: *string, // Required
	}

	if len(_wisdomAssistantId) > 0 {
		input.AssistantId = aws.String(_wisdomAssistantId)
	}
	if len(_wisdomMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wisdomMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wisdomNextToken) > 0 {
		input.NextToken = aws.String(_wisdomNextToken)
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

	var results []*wisdom.ListAssistantAssociationsOutput
	p := wisdom.NewListAssistantAssociationsPaginator(client, input)
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
func wisdom_ListAssistants(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.ListAssistantsInput{}

	if len(_wisdomMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wisdomMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wisdomNextToken) > 0 {
		input.NextToken = aws.String(_wisdomNextToken)
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

	var results []*wisdom.ListAssistantsOutput
	p := wisdom.NewListAssistantsPaginator(client, input)
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
func wisdom_ListContents(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.ListContentsInput{
		// KnowledgeBaseId: *string, // Required
	}

	if len(_wisdomKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_wisdomKnowledgeBaseId)
	}
	if len(_wisdomMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wisdomMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wisdomNextToken) > 0 {
		input.NextToken = aws.String(_wisdomNextToken)
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

	var results []*wisdom.ListContentsOutput
	p := wisdom.NewListContentsPaginator(client, input)
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
func wisdom_ListImportJobs(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.ListImportJobsInput{
		// KnowledgeBaseId: *string, // Required
	}

	if len(_wisdomKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_wisdomKnowledgeBaseId)
	}
	if len(_wisdomMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wisdomMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wisdomNextToken) > 0 {
		input.NextToken = aws.String(_wisdomNextToken)
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

	var results []*wisdom.ListImportJobsOutput
	p := wisdom.NewListImportJobsPaginator(client, input)
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
func wisdom_ListKnowledgeBases(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.ListKnowledgeBasesInput{}

	if len(_wisdomMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wisdomMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wisdomNextToken) > 0 {
		input.NextToken = aws.String(_wisdomNextToken)
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

	var results []*wisdom.ListKnowledgeBasesOutput
	p := wisdom.NewListKnowledgeBasesPaginator(client, input)
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
func wisdom_ListQuickResponses(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.ListQuickResponsesInput{
		// KnowledgeBaseId: *string, // Required
	}

	if len(_wisdomKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_wisdomKnowledgeBaseId)
	}
	if len(_wisdomMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wisdomMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wisdomNextToken) > 0 {
		input.NextToken = aws.String(_wisdomNextToken)
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

	var results []*wisdom.ListQuickResponsesOutput
	p := wisdom.NewListQuickResponsesPaginator(client, input)
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
func wisdom_ListTagsForResource(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_wisdomResourceArn) > 0 {
		input.ResourceArn = aws.String(_wisdomResourceArn)
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
// [GetRecommendations]: https://docs.aws.amazon.com/wisdom/latest/APIReference/API_GetRecommendations.html
func wisdom_NotifyRecommendationsReceived(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.NotifyRecommendationsReceivedInput{
		// AssistantId: *string, // Required
		// RecommendationIds: []string, // Required
		// SessionId: *string, // Required
	}

	if len(_wisdomAssistantId) > 0 {
		input.AssistantId = aws.String(_wisdomAssistantId)
	}
	if len(_wisdomRecommendationIds) > 0 {
		input.RecommendationIds = append([]string(nil), _wisdomRecommendationIds...)
	}
	if len(_wisdomSessionId) > 0 {
		input.SessionId = aws.String(_wisdomSessionId)
	}

	if resp, err := client.NotifyRecommendationsReceived(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Performs a manual search against the specified assistant. To retrieve
// recommendations for an assistant, use [GetRecommendations].
//
// Deprecated: QueryAssistant API will be discontinued starting June 1, 2024. To
// receive generative responses after March 1, 2024 you will need to create a new
// Assistant in the Connect console and integrate the Amazon Q in Connect
// JavaScript library (amazon-q-connectjs) into your applications.
//
// [GetRecommendations]: https://docs.aws.amazon.com/wisdom/latest/APIReference/API_GetRecommendations.html
func wisdom_QueryAssistant(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.QueryAssistantInput{
		// AssistantId: *string, // Required
		// QueryText: *string, // Required
	}

	if len(_wisdomAssistantId) > 0 {
		input.AssistantId = aws.String(_wisdomAssistantId)
	}
	if len(_wisdomQueryText) > 0 {
		input.QueryText = aws.String(_wisdomQueryText)
	}
	if len(_wisdomMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wisdomMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wisdomNextToken) > 0 {
		input.NextToken = aws.String(_wisdomNextToken)
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

	var results []*wisdom.QueryAssistantOutput
	p := wisdom.NewQueryAssistantPaginator(client, input)
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

// Removes a URI template from a knowledge base.
func wisdom_RemoveKnowledgeBaseTemplateUri(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.RemoveKnowledgeBaseTemplateUriInput{
		// KnowledgeBaseId: *string, // Required
	}

	if len(_wisdomKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_wisdomKnowledgeBaseId)
	}

	if resp, err := client.RemoveKnowledgeBaseTemplateUri(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches for content in a specified knowledge base. Can be used to get a
// specific content resource by its name.
func wisdom_SearchContent(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.SearchContentInput{
		// KnowledgeBaseId: *string, // Required
		// SearchExpression: *types.SearchExpression, // Required
	}

	if len(_wisdomKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_wisdomKnowledgeBaseId)
	}
	if len(_wisdomSearchExpression) > 0 {
		if err := assignInputField(input, "SearchExpression", _wisdomSearchExpression); err != nil {
			log.Errorf("invalid --search-expression: %s", err.Error())
			return
		}
	}
	if len(_wisdomMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wisdomMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wisdomNextToken) > 0 {
		input.NextToken = aws.String(_wisdomNextToken)
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

	var results []*wisdom.SearchContentOutput
	p := wisdom.NewSearchContentPaginator(client, input)
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

// Searches existing Wisdom quick responses in a Wisdom knowledge base.
func wisdom_SearchQuickResponses(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.SearchQuickResponsesInput{
		// KnowledgeBaseId: *string, // Required
		// SearchExpression: *types.QuickResponseSearchExpression, // Required
	}

	if len(_wisdomKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_wisdomKnowledgeBaseId)
	}
	if len(_wisdomSearchExpression) > 0 {
		if err := assignInputField(input, "SearchExpression", _wisdomSearchExpression); err != nil {
			log.Errorf("invalid --search-expression: %s", err.Error())
			return
		}
	}
	if len(_wisdomAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _wisdomAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_wisdomMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wisdomMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wisdomNextToken) > 0 {
		input.NextToken = aws.String(_wisdomNextToken)
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

	var results []*wisdom.SearchQuickResponsesOutput
	p := wisdom.NewSearchQuickResponsesPaginator(client, input)
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
func wisdom_SearchSessions(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.SearchSessionsInput{
		// AssistantId: *string, // Required
		// SearchExpression: *types.SearchExpression, // Required
	}

	if len(_wisdomAssistantId) > 0 {
		input.AssistantId = aws.String(_wisdomAssistantId)
	}
	if len(_wisdomSearchExpression) > 0 {
		if err := assignInputField(input, "SearchExpression", _wisdomSearchExpression); err != nil {
			log.Errorf("invalid --search-expression: %s", err.Error())
			return
		}
	}
	if len(_wisdomMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wisdomMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wisdomNextToken) > 0 {
		input.NextToken = aws.String(_wisdomNextToken)
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

	var results []*wisdom.SearchSessionsOutput
	p := wisdom.NewSearchSessionsPaginator(client, input)
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

// Get a URL to upload content to a knowledge base. To upload content, first make
// a PUT request to the returned URL with your file, making sure to include the
// required headers. Then use [CreateContent]to finalize the content creation process or [UpdateContent] to
// modify an existing resource. You can only upload content to a knowledge base of
// type CUSTOM.
//
// [CreateContent]: https://docs.aws.amazon.com/wisdom/latest/APIReference/API_CreateContent.html
// [UpdateContent]: https://docs.aws.amazon.com/wisdom/latest/APIReference/API_UpdateContent.html
func wisdom_StartContentUpload(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.StartContentUploadInput{
		// ContentType: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_wisdomContentType) > 0 {
		input.ContentType = aws.String(_wisdomContentType)
	}
	if len(_wisdomKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_wisdomKnowledgeBaseId)
	}
	if len(_wisdomPresignedUrlTimeToLive) > 0 {
		if err := assignInputField(input, "PresignedUrlTimeToLive", _wisdomPresignedUrlTimeToLive); err != nil {
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

// Start an asynchronous job to import Wisdom resources from an uploaded source
// file. Before calling this API, use [StartContentUpload]to upload an asset that contains the
// resource data.
//
// - For importing Wisdom quick responses, you need to upload a csv file
// including the quick responses. For information about how to format the csv file
// for importing quick responses, see [Import quick responses].
//
// [StartContentUpload]: https://docs.aws.amazon.com/wisdom/latest/APIReference/API_StartContentUpload.html
// [Import quick responses]: https://docs.aws.amazon.com/console/connect/quick-responses/add-data
func wisdom_StartImportJob(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.StartImportJobInput{
		// ImportJobType: types.ImportJobType, // Required
		// KnowledgeBaseId: *string, // Required
		// UploadId: *string, // Required
	}

	if len(_wisdomImportJobType) > 0 {
		if err := assignInputField(input, "ImportJobType", _wisdomImportJobType); err != nil {
			log.Errorf("invalid --import-job-type: %s", err.Error())
			return
		}
	}
	if len(_wisdomKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_wisdomKnowledgeBaseId)
	}
	if len(_wisdomUploadId) > 0 {
		input.UploadId = aws.String(_wisdomUploadId)
	}
	if len(_wisdomClientToken) > 0 {
		input.ClientToken = aws.String(_wisdomClientToken)
	}
	if len(_wisdomExternalSourceConfiguration) > 0 {
		if err := assignInputField(input, "ExternalSourceConfiguration", _wisdomExternalSourceConfiguration); err != nil {
			log.Errorf("invalid --external-source-configuration: %s", err.Error())
			return
		}
	}
	if len(_wisdomMetadata) > 0 {
		if err := assignInputField(input, "Metadata", _wisdomMetadata); err != nil {
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
func wisdom_TagResource(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_wisdomResourceArn) > 0 {
		input.ResourceArn = aws.String(_wisdomResourceArn)
	}
	if len(_wisdomTags) > 0 {
		if err := assignInputField(input, "Tags", _wisdomTags); err != nil {
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
func wisdom_UntagResource(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_wisdomResourceArn) > 0 {
		input.ResourceArn = aws.String(_wisdomResourceArn)
	}
	if len(_wisdomTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _wisdomTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates information about the content.
func wisdom_UpdateContent(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.UpdateContentInput{
		// ContentId: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_wisdomContentId) > 0 {
		input.ContentId = aws.String(_wisdomContentId)
	}
	if len(_wisdomKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_wisdomKnowledgeBaseId)
	}
	if len(_wisdomMetadata) > 0 {
		if err := assignInputField(input, "Metadata", _wisdomMetadata); err != nil {
			log.Errorf("invalid --metadata: %s", err.Error())
			return
		}
	}
	if len(_wisdomOverrideLinkOutUri) > 0 {
		input.OverrideLinkOutUri = aws.String(_wisdomOverrideLinkOutUri)
	}
	if len(_wisdomRemoveOverrideLinkOutUri) > 0 {
		if err := assignInputField(input, "RemoveOverrideLinkOutUri", _wisdomRemoveOverrideLinkOutUri); err != nil {
			log.Errorf("invalid --remove-override-link-out-uri: %s", err.Error())
			return
		}
	}
	if len(_wisdomRevisionId) > 0 {
		input.RevisionId = aws.String(_wisdomRevisionId)
	}
	if len(_wisdomTitle) > 0 {
		input.Title = aws.String(_wisdomTitle)
	}
	if len(_wisdomUploadId) > 0 {
		input.UploadId = aws.String(_wisdomUploadId)
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
// format; this interpolated by Wisdom using ingested content. For example, if you
// ingest a Salesforce article, it has an Id value, and you can set the template
// URI to
// https://myInstanceName.lightning.force.com/lightning/r/Knowledge__kav/*${Id}*/view
// .
func wisdom_UpdateKnowledgeBaseTemplateUri(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.UpdateKnowledgeBaseTemplateUriInput{
		// KnowledgeBaseId: *string, // Required
		// TemplateUri: *string, // Required
	}

	if len(_wisdomKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_wisdomKnowledgeBaseId)
	}
	if len(_wisdomTemplateUri) > 0 {
		input.TemplateUri = aws.String(_wisdomTemplateUri)
	}

	if resp, err := client.UpdateKnowledgeBaseTemplateUri(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing Wisdom quick response.
func wisdom_UpdateQuickResponse(cfg aws.Config, client *wisdom.Client) {
	input := &wisdom.UpdateQuickResponseInput{
		// KnowledgeBaseId: *string, // Required
		// QuickResponseId: *string, // Required
	}

	if len(_wisdomKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_wisdomKnowledgeBaseId)
	}
	if len(_wisdomQuickResponseId) > 0 {
		input.QuickResponseId = aws.String(_wisdomQuickResponseId)
	}
	if len(_wisdomChannels) > 0 {
		input.Channels = append([]string(nil), _wisdomChannels...)
	}
	if len(_wisdomContent) > 0 {
		if err := assignInputField(input, "Content", _wisdomContent); err != nil {
			log.Errorf("invalid --content: %s", err.Error())
			return
		}
	}
	if len(_wisdomContentType) > 0 {
		input.ContentType = aws.String(_wisdomContentType)
	}
	if len(_wisdomDescription) > 0 {
		input.Description = aws.String(_wisdomDescription)
	}
	if len(_wisdomGroupingConfiguration) > 0 {
		if err := assignInputField(input, "GroupingConfiguration", _wisdomGroupingConfiguration); err != nil {
			log.Errorf("invalid --grouping-configuration: %s", err.Error())
			return
		}
	}
	if len(_wisdomIsActive) > 0 {
		if err := assignInputField(input, "IsActive", _wisdomIsActive); err != nil {
			log.Errorf("invalid --is-active: %s", err.Error())
			return
		}
	}
	if len(_wisdomLanguage) > 0 {
		input.Language = aws.String(_wisdomLanguage)
	}
	if len(_wisdomName) > 0 {
		input.Name = aws.String(_wisdomName)
	}
	if len(_wisdomRemoveDescription) > 0 {
		if err := assignInputField(input, "RemoveDescription", _wisdomRemoveDescription); err != nil {
			log.Errorf("invalid --remove-description: %s", err.Error())
			return
		}
	}
	if len(_wisdomRemoveGroupingConfiguration) > 0 {
		if err := assignInputField(input, "RemoveGroupingConfiguration", _wisdomRemoveGroupingConfiguration); err != nil {
			log.Errorf("invalid --remove-grouping-configuration: %s", err.Error())
			return
		}
	}
	if len(_wisdomRemoveShortcutKey) > 0 {
		if err := assignInputField(input, "RemoveShortcutKey", _wisdomRemoveShortcutKey); err != nil {
			log.Errorf("invalid --remove-shortcut-key: %s", err.Error())
			return
		}
	}
	if len(_wisdomShortcutKey) > 0 {
		input.ShortcutKey = aws.String(_wisdomShortcutKey)
	}

	if resp, err := client.UpdateQuickResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_wisdomCmd)
	_wisdomCmd.Flags().SortFlags = false

	_wisdomCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_wisdomCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_wisdomCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_wisdomCmd.Flags().StringVarP(&_wisdomAssistantAssociationId, "assistant-association-id", "", "", "Assistant Association ID")
	_wisdomCmd.Flags().StringVarP(&_wisdomAssistantId, "assistant-id", "", "", "Assistant ID")
	_wisdomCmd.Flags().StringVarP(&_wisdomAssociation, "association", "", "", "Association")
	_wisdomCmd.Flags().StringVarP(&_wisdomAssociationType, "association-type", "", "", "Association Type")
	_wisdomCmd.Flags().StringVarP(&_wisdomAttributes, "attributes", "", "", "Attributes")
	_wisdomCmd.Flags().StringSliceVarP(&_wisdomChannels, "channels", "", nil, "Channels")
	_wisdomCmd.Flags().StringVarP(&_wisdomClientToken, "client-token", "", "", "Client Token")
	_wisdomCmd.Flags().StringVarP(&_wisdomContent, "content", "", "", "Content")
	_wisdomCmd.Flags().StringVarP(&_wisdomContentId, "content-id", "", "", "Content ID")
	_wisdomCmd.Flags().StringVarP(&_wisdomContentType, "content-type", "", "", "Content Type")
	_wisdomCmd.Flags().StringVarP(&_wisdomDescription, "description", "", "", "Description")
	_wisdomCmd.Flags().StringVarP(&_wisdomExternalSourceConfiguration, "external-source-configuration", "", "", "External Source Configuration")
	_wisdomCmd.Flags().StringVarP(&_wisdomGroupingConfiguration, "grouping-configuration", "", "", "Grouping Configuration")
	_wisdomCmd.Flags().StringVarP(&_wisdomImportJobId, "import-job-id", "", "", "Import Job ID")
	_wisdomCmd.Flags().StringVarP(&_wisdomImportJobType, "import-job-type", "", "", "Import Job Type")
	_wisdomCmd.Flags().StringVarP(&_wisdomIsActive, "is-active", "", "", "Is Active")
	_wisdomCmd.Flags().StringVarP(&_wisdomKnowledgeBaseId, "knowledge-base-id", "", "", "Knowledge Base ID")
	_wisdomCmd.Flags().StringVarP(&_wisdomKnowledgeBaseType, "knowledge-base-type", "", "", "Knowledge Base Type")
	_wisdomCmd.Flags().StringVarP(&_wisdomLanguage, "language", "", "", "Language")
	_wisdomCmd.Flags().StringVarP(&_wisdomMaxResults, "max-results", "", "", "Max Results")
	_wisdomCmd.Flags().StringVarP(&_wisdomMetadata, "metadata", "", "", "Metadata")
	_wisdomCmd.Flags().StringVarP(&_wisdomName, "name", "", "", "Name")
	_wisdomCmd.Flags().StringVarP(&_wisdomNextToken, "next-token", "", "", "Next Token")
	_wisdomCmd.Flags().StringVarP(&_wisdomOverrideLinkOutUri, "override-link-out-uri", "", "", "Override Link Out URI")
	_wisdomCmd.Flags().StringVarP(&_wisdomPresignedUrlTimeToLive, "presigned-url-time-to-live", "", "", "Presigned URL Time To Live")
	_wisdomCmd.Flags().StringVarP(&_wisdomQueryText, "query-text", "", "", "Query Text")
	_wisdomCmd.Flags().StringVarP(&_wisdomQuickResponseId, "quick-response-id", "", "", "Quick Response ID")
	_wisdomCmd.Flags().StringSliceVarP(&_wisdomRecommendationIds, "recommendation-ids", "", nil, "Recommendation Ids")
	_wisdomCmd.Flags().StringVarP(&_wisdomRemoveDescription, "remove-description", "", "", "Remove Description")
	_wisdomCmd.Flags().StringVarP(&_wisdomRemoveGroupingConfiguration, "remove-grouping-configuration", "", "", "Remove Grouping Configuration")
	_wisdomCmd.Flags().StringVarP(&_wisdomRemoveOverrideLinkOutUri, "remove-override-link-out-uri", "", "", "Remove Override Link Out URI")
	_wisdomCmd.Flags().StringVarP(&_wisdomRemoveShortcutKey, "remove-shortcut-key", "", "", "Remove Shortcut Key")
	_wisdomCmd.Flags().StringVarP(&_wisdomRenderingConfiguration, "rendering-configuration", "", "", "Rendering Configuration")
	_wisdomCmd.Flags().StringVarP(&_wisdomResourceArn, "resource-arn", "", "", "Resource ARN")
	_wisdomCmd.Flags().StringVarP(&_wisdomRevisionId, "revision-id", "", "", "Revision ID")
	_wisdomCmd.Flags().StringVarP(&_wisdomSearchExpression, "search-expression", "", "", "Search Expression")
	_wisdomCmd.Flags().StringVarP(&_wisdomServerSideEncryptionConfiguration, "server-side-encryption-configuration", "", "", "Server Side Encryption Configuration")
	_wisdomCmd.Flags().StringVarP(&_wisdomSessionId, "session-id", "", "", "Session ID")
	_wisdomCmd.Flags().StringVarP(&_wisdomShortcutKey, "shortcut-key", "", "", "Shortcut Key")
	_wisdomCmd.Flags().StringVarP(&_wisdomSourceConfiguration, "source-configuration", "", "", "Source Configuration")
	_wisdomCmd.Flags().StringSliceVarP(&_wisdomTagKeys, "tag-keys", "", nil, "Tag Keys")
	_wisdomCmd.Flags().StringVarP(&_wisdomTags, "tags", "", "", "Tags")
	_wisdomCmd.Flags().StringVarP(&_wisdomTemplateUri, "template-uri", "", "", "Template URI")
	_wisdomCmd.Flags().StringVarP(&_wisdomTitle, "title", "", "", "Title")
	_wisdomCmd.Flags().StringVarP(&_wisdomType, "type", "", "", "Type")
	_wisdomCmd.Flags().StringVarP(&_wisdomUploadId, "upload-id", "", "", "Upload ID")
	_wisdomCmd.Flags().StringVarP(&_wisdomWaitTimeSeconds, "wait-time-seconds", "", "", "Wait Time Seconds")

	_wisdomCmd.Flags().BoolVarP(&_wisdomCreateAssistant, "create-assistant", "", false, "Create Assistant")
	_wisdomCmd.Flags().BoolVarP(&_wisdomCreateAssistantAssociation, "create-assistant-association", "", false, "Create Assistant Association")
	_wisdomCmd.Flags().BoolVarP(&_wisdomCreateContent, "create-content", "", false, "Create Content")
	_wisdomCmd.Flags().BoolVarP(&_wisdomCreateKnowledgeBase, "create-knowledge-base", "", false, "Create Knowledge Base")
	_wisdomCmd.Flags().BoolVarP(&_wisdomCreateQuickResponse, "create-quick-response", "", false, "Create Quick Response")
	_wisdomCmd.Flags().BoolVarP(&_wisdomCreateSession, "create-session", "", false, "Create Session")
	_wisdomCmd.Flags().BoolVarP(&_wisdomDeleteAssistant, "delete-assistant", "", false, "Delete Assistant")
	_wisdomCmd.Flags().BoolVarP(&_wisdomDeleteAssistantAssociation, "delete-assistant-association", "", false, "Delete Assistant Association")
	_wisdomCmd.Flags().BoolVarP(&_wisdomDeleteContent, "delete-content", "", false, "Delete Content")
	_wisdomCmd.Flags().BoolVarP(&_wisdomDeleteImportJob, "delete-import-job", "", false, "Delete Import Job")
	_wisdomCmd.Flags().BoolVarP(&_wisdomDeleteKnowledgeBase, "delete-knowledge-base", "", false, "Delete Knowledge Base")
	_wisdomCmd.Flags().BoolVarP(&_wisdomDeleteQuickResponse, "delete-quick-response", "", false, "Delete Quick Response")
	_wisdomCmd.Flags().BoolVarP(&_wisdomGetAssistant, "get-assistant", "", false, "Get Assistant")
	_wisdomCmd.Flags().BoolVarP(&_wisdomGetAssistantAssociation, "get-assistant-association", "", false, "Get Assistant Association")
	_wisdomCmd.Flags().BoolVarP(&_wisdomGetContent, "get-content", "", false, "Get Content")
	_wisdomCmd.Flags().BoolVarP(&_wisdomGetContentSummary, "get-content-summary", "", false, "Get Content Summary")
	_wisdomCmd.Flags().BoolVarP(&_wisdomGetImportJob, "get-import-job", "", false, "Get Import Job")
	_wisdomCmd.Flags().BoolVarP(&_wisdomGetKnowledgeBase, "get-knowledge-base", "", false, "Get Knowledge Base")
	_wisdomCmd.Flags().BoolVarP(&_wisdomGetQuickResponse, "get-quick-response", "", false, "Get Quick Response")
	_wisdomCmd.Flags().BoolVarP(&_wisdomGetRecommendations, "get-recommendations", "", false, "Get Recommendations")
	_wisdomCmd.Flags().BoolVarP(&_wisdomGetSession, "get-session", "", false, "Get Session")
	_wisdomCmd.Flags().BoolVarP(&_wisdomListAssistantAssociations, "list-assistant-associations", "", false, "List Assistant Associations")
	_wisdomCmd.Flags().BoolVarP(&_wisdomListAssistants, "list-assistants", "", false, "List Assistants")
	_wisdomCmd.Flags().BoolVarP(&_wisdomListContents, "list-contents", "", false, "List Contents")
	_wisdomCmd.Flags().BoolVarP(&_wisdomListImportJobs, "list-import-jobs", "", false, "List Import Jobs")
	_wisdomCmd.Flags().BoolVarP(&_wisdomListKnowledgeBases, "list-knowledge-bases", "", false, "List Knowledge Bases")
	_wisdomCmd.Flags().BoolVarP(&_wisdomListQuickResponses, "list-quick-responses", "", false, "List Quick Responses")
	_wisdomCmd.Flags().BoolVarP(&_wisdomListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_wisdomCmd.Flags().BoolVarP(&_wisdomNotifyRecommendationsReceived, "notify-recommendations-received", "", false, "Notify Recommendations Received")
	_wisdomCmd.Flags().BoolVarP(&_wisdomQueryAssistant, "query-assistant", "", false, "Query Assistant")
	_wisdomCmd.Flags().BoolVarP(&_wisdomRemoveKnowledgeBaseTemplateUri, "remove-knowledge-base-template-uri", "", false, "Remove Knowledge Base Template URI")
	_wisdomCmd.Flags().BoolVarP(&_wisdomSearchContent, "search-content", "", false, "Search Content")
	_wisdomCmd.Flags().BoolVarP(&_wisdomSearchQuickResponses, "search-quick-responses", "", false, "Search Quick Responses")
	_wisdomCmd.Flags().BoolVarP(&_wisdomSearchSessions, "search-sessions", "", false, "Search Sessions")
	_wisdomCmd.Flags().BoolVarP(&_wisdomStartContentUpload, "start-content-upload", "", false, "Start Content Upload")
	_wisdomCmd.Flags().BoolVarP(&_wisdomStartImportJob, "start-import-job", "", false, "Start Import Job")
	_wisdomCmd.Flags().BoolVarP(&_wisdomTagResource, "tag-resource", "", false, "Tag Resource")
	_wisdomCmd.Flags().BoolVarP(&_wisdomUntagResource, "untag-resource", "", false, "Untag Resource")
	_wisdomCmd.Flags().BoolVarP(&_wisdomUpdateContent, "update-content", "", false, "Update Content")
	_wisdomCmd.Flags().BoolVarP(&_wisdomUpdateKnowledgeBaseTemplateUri, "update-knowledge-base-template-uri", "", false, "Update Knowledge Base Template URI")
	_wisdomCmd.Flags().BoolVarP(&_wisdomUpdateQuickResponse, "update-quick-response", "", false, "Update Quick Response")

}
