package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// lexmodelsv2Cmd represents the lexmodelsv2 command
var _lexmodelsv2Cmd = &cobra.Command{
	Use:   "lexmodelsv2",
	Short: "AWS lexmodelsv2 CLI",
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
		client := lexmodelsv2.NewFromConfig(cfg)
		if _lexmodelsv2BatchCreateCustomVocabularyItem {
			lexmodelsv2_BatchCreateCustomVocabularyItem(cfg, client)
			return
		}
		if _lexmodelsv2BatchDeleteCustomVocabularyItem {
			lexmodelsv2_BatchDeleteCustomVocabularyItem(cfg, client)
			return
		}
		if _lexmodelsv2BatchUpdateCustomVocabularyItem {
			lexmodelsv2_BatchUpdateCustomVocabularyItem(cfg, client)
			return
		}
		if _lexmodelsv2BuildBotLocale {
			lexmodelsv2_BuildBotLocale(cfg, client)
			return
		}
		if _lexmodelsv2CreateBot {
			lexmodelsv2_CreateBot(cfg, client)
			return
		}
		if _lexmodelsv2CreateBotAlias {
			lexmodelsv2_CreateBotAlias(cfg, client)
			return
		}
		if _lexmodelsv2CreateBotLocale {
			lexmodelsv2_CreateBotLocale(cfg, client)
			return
		}
		if _lexmodelsv2CreateBotReplica {
			lexmodelsv2_CreateBotReplica(cfg, client)
			return
		}
		if _lexmodelsv2CreateBotVersion {
			lexmodelsv2_CreateBotVersion(cfg, client)
			return
		}
		if _lexmodelsv2CreateExport {
			lexmodelsv2_CreateExport(cfg, client)
			return
		}
		if _lexmodelsv2CreateIntent {
			lexmodelsv2_CreateIntent(cfg, client)
			return
		}
		if _lexmodelsv2CreateResourcePolicy {
			lexmodelsv2_CreateResourcePolicy(cfg, client)
			return
		}
		if _lexmodelsv2CreateResourcePolicyStatement {
			lexmodelsv2_CreateResourcePolicyStatement(cfg, client)
			return
		}
		if _lexmodelsv2CreateSlot {
			lexmodelsv2_CreateSlot(cfg, client)
			return
		}
		if _lexmodelsv2CreateSlotType {
			lexmodelsv2_CreateSlotType(cfg, client)
			return
		}
		if _lexmodelsv2CreateTestSetDiscrepancyReport {
			lexmodelsv2_CreateTestSetDiscrepancyReport(cfg, client)
			return
		}
		if _lexmodelsv2CreateUploadUrl {
			lexmodelsv2_CreateUploadUrl(cfg, client)
			return
		}
		if _lexmodelsv2DeleteBot {
			lexmodelsv2_DeleteBot(cfg, client)
			return
		}
		if _lexmodelsv2DeleteBotAlias {
			lexmodelsv2_DeleteBotAlias(cfg, client)
			return
		}
		if _lexmodelsv2DeleteBotLocale {
			lexmodelsv2_DeleteBotLocale(cfg, client)
			return
		}
		if _lexmodelsv2DeleteBotReplica {
			lexmodelsv2_DeleteBotReplica(cfg, client)
			return
		}
		if _lexmodelsv2DeleteBotVersion {
			lexmodelsv2_DeleteBotVersion(cfg, client)
			return
		}
		if _lexmodelsv2DeleteCustomVocabulary {
			lexmodelsv2_DeleteCustomVocabulary(cfg, client)
			return
		}
		if _lexmodelsv2DeleteExport {
			lexmodelsv2_DeleteExport(cfg, client)
			return
		}
		if _lexmodelsv2DeleteImport {
			lexmodelsv2_DeleteImport(cfg, client)
			return
		}
		if _lexmodelsv2DeleteIntent {
			lexmodelsv2_DeleteIntent(cfg, client)
			return
		}
		if _lexmodelsv2DeleteResourcePolicy {
			lexmodelsv2_DeleteResourcePolicy(cfg, client)
			return
		}
		if _lexmodelsv2DeleteResourcePolicyStatement {
			lexmodelsv2_DeleteResourcePolicyStatement(cfg, client)
			return
		}
		if _lexmodelsv2DeleteSlot {
			lexmodelsv2_DeleteSlot(cfg, client)
			return
		}
		if _lexmodelsv2DeleteSlotType {
			lexmodelsv2_DeleteSlotType(cfg, client)
			return
		}
		if _lexmodelsv2DeleteTestSet {
			lexmodelsv2_DeleteTestSet(cfg, client)
			return
		}
		if _lexmodelsv2DeleteUtterances {
			lexmodelsv2_DeleteUtterances(cfg, client)
			return
		}
		if _lexmodelsv2DescribeBot {
			lexmodelsv2_DescribeBot(cfg, client)
			return
		}
		if _lexmodelsv2DescribeBotAlias {
			lexmodelsv2_DescribeBotAlias(cfg, client)
			return
		}
		if _lexmodelsv2DescribeBotLocale {
			lexmodelsv2_DescribeBotLocale(cfg, client)
			return
		}
		if _lexmodelsv2DescribeBotRecommendation {
			lexmodelsv2_DescribeBotRecommendation(cfg, client)
			return
		}
		if _lexmodelsv2DescribeBotReplica {
			lexmodelsv2_DescribeBotReplica(cfg, client)
			return
		}
		if _lexmodelsv2DescribeBotResourceGeneration {
			lexmodelsv2_DescribeBotResourceGeneration(cfg, client)
			return
		}
		if _lexmodelsv2DescribeBotVersion {
			lexmodelsv2_DescribeBotVersion(cfg, client)
			return
		}
		if _lexmodelsv2DescribeCustomVocabularyMetadata {
			lexmodelsv2_DescribeCustomVocabularyMetadata(cfg, client)
			return
		}
		if _lexmodelsv2DescribeExport {
			lexmodelsv2_DescribeExport(cfg, client)
			return
		}
		if _lexmodelsv2DescribeImport {
			lexmodelsv2_DescribeImport(cfg, client)
			return
		}
		if _lexmodelsv2DescribeIntent {
			lexmodelsv2_DescribeIntent(cfg, client)
			return
		}
		if _lexmodelsv2DescribeResourcePolicy {
			lexmodelsv2_DescribeResourcePolicy(cfg, client)
			return
		}
		if _lexmodelsv2DescribeSlot {
			lexmodelsv2_DescribeSlot(cfg, client)
			return
		}
		if _lexmodelsv2DescribeSlotType {
			lexmodelsv2_DescribeSlotType(cfg, client)
			return
		}
		if _lexmodelsv2DescribeTestExecution {
			lexmodelsv2_DescribeTestExecution(cfg, client)
			return
		}
		if _lexmodelsv2DescribeTestSet {
			lexmodelsv2_DescribeTestSet(cfg, client)
			return
		}
		if _lexmodelsv2DescribeTestSetDiscrepancyReport {
			lexmodelsv2_DescribeTestSetDiscrepancyReport(cfg, client)
			return
		}
		if _lexmodelsv2DescribeTestSetGeneration {
			lexmodelsv2_DescribeTestSetGeneration(cfg, client)
			return
		}
		if _lexmodelsv2GenerateBotElement {
			lexmodelsv2_GenerateBotElement(cfg, client)
			return
		}
		if _lexmodelsv2GetTestExecutionArtifactsUrl {
			lexmodelsv2_GetTestExecutionArtifactsUrl(cfg, client)
			return
		}
		if _lexmodelsv2ListAggregatedUtterances {
			lexmodelsv2_ListAggregatedUtterances(cfg, client)
			return
		}
		if _lexmodelsv2ListBotAliasReplicas {
			lexmodelsv2_ListBotAliasReplicas(cfg, client)
			return
		}
		if _lexmodelsv2ListBotAliases {
			lexmodelsv2_ListBotAliases(cfg, client)
			return
		}
		if _lexmodelsv2ListBotLocales {
			lexmodelsv2_ListBotLocales(cfg, client)
			return
		}
		if _lexmodelsv2ListBotRecommendations {
			lexmodelsv2_ListBotRecommendations(cfg, client)
			return
		}
		if _lexmodelsv2ListBotReplicas {
			lexmodelsv2_ListBotReplicas(cfg, client)
			return
		}
		if _lexmodelsv2ListBotResourceGenerations {
			lexmodelsv2_ListBotResourceGenerations(cfg, client)
			return
		}
		if _lexmodelsv2ListBotVersionReplicas {
			lexmodelsv2_ListBotVersionReplicas(cfg, client)
			return
		}
		if _lexmodelsv2ListBotVersions {
			lexmodelsv2_ListBotVersions(cfg, client)
			return
		}
		if _lexmodelsv2ListBots {
			lexmodelsv2_ListBots(cfg, client)
			return
		}
		if _lexmodelsv2ListBuiltInIntents {
			lexmodelsv2_ListBuiltInIntents(cfg, client)
			return
		}
		if _lexmodelsv2ListBuiltInSlotTypes {
			lexmodelsv2_ListBuiltInSlotTypes(cfg, client)
			return
		}
		if _lexmodelsv2ListCustomVocabularyItems {
			lexmodelsv2_ListCustomVocabularyItems(cfg, client)
			return
		}
		if _lexmodelsv2ListExports {
			lexmodelsv2_ListExports(cfg, client)
			return
		}
		if _lexmodelsv2ListImports {
			lexmodelsv2_ListImports(cfg, client)
			return
		}
		if _lexmodelsv2ListIntentMetrics {
			lexmodelsv2_ListIntentMetrics(cfg, client)
			return
		}
		if _lexmodelsv2ListIntentPaths {
			lexmodelsv2_ListIntentPaths(cfg, client)
			return
		}
		if _lexmodelsv2ListIntentStageMetrics {
			lexmodelsv2_ListIntentStageMetrics(cfg, client)
			return
		}
		if _lexmodelsv2ListIntents {
			lexmodelsv2_ListIntents(cfg, client)
			return
		}
		if _lexmodelsv2ListRecommendedIntents {
			lexmodelsv2_ListRecommendedIntents(cfg, client)
			return
		}
		if _lexmodelsv2ListSessionAnalyticsData {
			lexmodelsv2_ListSessionAnalyticsData(cfg, client)
			return
		}
		if _lexmodelsv2ListSessionMetrics {
			lexmodelsv2_ListSessionMetrics(cfg, client)
			return
		}
		if _lexmodelsv2ListSlotTypes {
			lexmodelsv2_ListSlotTypes(cfg, client)
			return
		}
		if _lexmodelsv2ListSlots {
			lexmodelsv2_ListSlots(cfg, client)
			return
		}
		if _lexmodelsv2ListTagsForResource {
			lexmodelsv2_ListTagsForResource(cfg, client)
			return
		}
		if _lexmodelsv2ListTestExecutionResultItems {
			lexmodelsv2_ListTestExecutionResultItems(cfg, client)
			return
		}
		if _lexmodelsv2ListTestExecutions {
			lexmodelsv2_ListTestExecutions(cfg, client)
			return
		}
		if _lexmodelsv2ListTestSetRecords {
			lexmodelsv2_ListTestSetRecords(cfg, client)
			return
		}
		if _lexmodelsv2ListTestSets {
			lexmodelsv2_ListTestSets(cfg, client)
			return
		}
		if _lexmodelsv2ListUtteranceAnalyticsData {
			lexmodelsv2_ListUtteranceAnalyticsData(cfg, client)
			return
		}
		if _lexmodelsv2ListUtteranceMetrics {
			lexmodelsv2_ListUtteranceMetrics(cfg, client)
			return
		}
		if _lexmodelsv2SearchAssociatedTranscripts {
			lexmodelsv2_SearchAssociatedTranscripts(cfg, client)
			return
		}
		if _lexmodelsv2StartBotRecommendation {
			lexmodelsv2_StartBotRecommendation(cfg, client)
			return
		}
		if _lexmodelsv2StartBotResourceGeneration {
			lexmodelsv2_StartBotResourceGeneration(cfg, client)
			return
		}
		if _lexmodelsv2StartImport {
			lexmodelsv2_StartImport(cfg, client)
			return
		}
		if _lexmodelsv2StartTestExecution {
			lexmodelsv2_StartTestExecution(cfg, client)
			return
		}
		if _lexmodelsv2StartTestSetGeneration {
			lexmodelsv2_StartTestSetGeneration(cfg, client)
			return
		}
		if _lexmodelsv2StopBotRecommendation {
			lexmodelsv2_StopBotRecommendation(cfg, client)
			return
		}
		if _lexmodelsv2TagResource {
			lexmodelsv2_TagResource(cfg, client)
			return
		}
		if _lexmodelsv2UntagResource {
			lexmodelsv2_UntagResource(cfg, client)
			return
		}
		if _lexmodelsv2UpdateBot {
			lexmodelsv2_UpdateBot(cfg, client)
			return
		}
		if _lexmodelsv2UpdateBotAlias {
			lexmodelsv2_UpdateBotAlias(cfg, client)
			return
		}
		if _lexmodelsv2UpdateBotLocale {
			lexmodelsv2_UpdateBotLocale(cfg, client)
			return
		}
		if _lexmodelsv2UpdateBotRecommendation {
			lexmodelsv2_UpdateBotRecommendation(cfg, client)
			return
		}
		if _lexmodelsv2UpdateExport {
			lexmodelsv2_UpdateExport(cfg, client)
			return
		}
		if _lexmodelsv2UpdateIntent {
			lexmodelsv2_UpdateIntent(cfg, client)
			return
		}
		if _lexmodelsv2UpdateResourcePolicy {
			lexmodelsv2_UpdateResourcePolicy(cfg, client)
			return
		}
		if _lexmodelsv2UpdateSlot {
			lexmodelsv2_UpdateSlot(cfg, client)
			return
		}
		if _lexmodelsv2UpdateSlotType {
			lexmodelsv2_UpdateSlotType(cfg, client)
			return
		}
		if _lexmodelsv2UpdateTestSet {
			lexmodelsv2_UpdateTestSet(cfg, client)
			return
		}

	},
}

var (
	_lexmodelsv2BatchCreateCustomVocabularyItem  bool
	_lexmodelsv2BatchDeleteCustomVocabularyItem  bool
	_lexmodelsv2BatchUpdateCustomVocabularyItem  bool
	_lexmodelsv2BuildBotLocale                   bool
	_lexmodelsv2CreateBot                        bool
	_lexmodelsv2CreateBotAlias                   bool
	_lexmodelsv2CreateBotLocale                  bool
	_lexmodelsv2CreateBotReplica                 bool
	_lexmodelsv2CreateBotVersion                 bool
	_lexmodelsv2CreateExport                     bool
	_lexmodelsv2CreateIntent                     bool
	_lexmodelsv2CreateResourcePolicy             bool
	_lexmodelsv2CreateResourcePolicyStatement    bool
	_lexmodelsv2CreateSlot                       bool
	_lexmodelsv2CreateSlotType                   bool
	_lexmodelsv2CreateTestSetDiscrepancyReport   bool
	_lexmodelsv2CreateUploadUrl                  bool
	_lexmodelsv2DeleteBot                        bool
	_lexmodelsv2DeleteBotAlias                   bool
	_lexmodelsv2DeleteBotLocale                  bool
	_lexmodelsv2DeleteBotReplica                 bool
	_lexmodelsv2DeleteBotVersion                 bool
	_lexmodelsv2DeleteCustomVocabulary           bool
	_lexmodelsv2DeleteExport                     bool
	_lexmodelsv2DeleteImport                     bool
	_lexmodelsv2DeleteIntent                     bool
	_lexmodelsv2DeleteResourcePolicy             bool
	_lexmodelsv2DeleteResourcePolicyStatement    bool
	_lexmodelsv2DeleteSlot                       bool
	_lexmodelsv2DeleteSlotType                   bool
	_lexmodelsv2DeleteTestSet                    bool
	_lexmodelsv2DeleteUtterances                 bool
	_lexmodelsv2DescribeBot                      bool
	_lexmodelsv2DescribeBotAlias                 bool
	_lexmodelsv2DescribeBotLocale                bool
	_lexmodelsv2DescribeBotRecommendation        bool
	_lexmodelsv2DescribeBotReplica               bool
	_lexmodelsv2DescribeBotResourceGeneration    bool
	_lexmodelsv2DescribeBotVersion               bool
	_lexmodelsv2DescribeCustomVocabularyMetadata bool
	_lexmodelsv2DescribeExport                   bool
	_lexmodelsv2DescribeImport                   bool
	_lexmodelsv2DescribeIntent                   bool
	_lexmodelsv2DescribeResourcePolicy           bool
	_lexmodelsv2DescribeSlot                     bool
	_lexmodelsv2DescribeSlotType                 bool
	_lexmodelsv2DescribeTestExecution            bool
	_lexmodelsv2DescribeTestSet                  bool
	_lexmodelsv2DescribeTestSetDiscrepancyReport bool
	_lexmodelsv2DescribeTestSetGeneration        bool
	_lexmodelsv2GenerateBotElement               bool
	_lexmodelsv2GetTestExecutionArtifactsUrl     bool
	_lexmodelsv2ListAggregatedUtterances         bool
	_lexmodelsv2ListBotAliasReplicas             bool
	_lexmodelsv2ListBotAliases                   bool
	_lexmodelsv2ListBotLocales                   bool
	_lexmodelsv2ListBotRecommendations           bool
	_lexmodelsv2ListBotReplicas                  bool
	_lexmodelsv2ListBotResourceGenerations       bool
	_lexmodelsv2ListBotVersionReplicas           bool
	_lexmodelsv2ListBotVersions                  bool
	_lexmodelsv2ListBots                         bool
	_lexmodelsv2ListBuiltInIntents               bool
	_lexmodelsv2ListBuiltInSlotTypes             bool
	_lexmodelsv2ListCustomVocabularyItems        bool
	_lexmodelsv2ListExports                      bool
	_lexmodelsv2ListImports                      bool
	_lexmodelsv2ListIntentMetrics                bool
	_lexmodelsv2ListIntentPaths                  bool
	_lexmodelsv2ListIntentStageMetrics           bool
	_lexmodelsv2ListIntents                      bool
	_lexmodelsv2ListRecommendedIntents           bool
	_lexmodelsv2ListSessionAnalyticsData         bool
	_lexmodelsv2ListSessionMetrics               bool
	_lexmodelsv2ListSlotTypes                    bool
	_lexmodelsv2ListSlots                        bool
	_lexmodelsv2ListTagsForResource              bool
	_lexmodelsv2ListTestExecutionResultItems     bool
	_lexmodelsv2ListTestExecutions               bool
	_lexmodelsv2ListTestSetRecords               bool
	_lexmodelsv2ListTestSets                     bool
	_lexmodelsv2ListUtteranceAnalyticsData       bool
	_lexmodelsv2ListUtteranceMetrics             bool
	_lexmodelsv2SearchAssociatedTranscripts      bool
	_lexmodelsv2StartBotRecommendation           bool
	_lexmodelsv2StartBotResourceGeneration       bool
	_lexmodelsv2StartImport                      bool
	_lexmodelsv2StartTestExecution               bool
	_lexmodelsv2StartTestSetGeneration           bool
	_lexmodelsv2StopBotRecommendation            bool
	_lexmodelsv2TagResource                      bool
	_lexmodelsv2UntagResource                    bool
	_lexmodelsv2UpdateBot                        bool
	_lexmodelsv2UpdateBotAlias                   bool
	_lexmodelsv2UpdateBotLocale                  bool
	_lexmodelsv2UpdateBotRecommendation          bool
	_lexmodelsv2UpdateExport                     bool
	_lexmodelsv2UpdateIntent                     bool
	_lexmodelsv2UpdateResourcePolicy             bool
	_lexmodelsv2UpdateSlot                       bool
	_lexmodelsv2UpdateSlotType                   bool
	_lexmodelsv2UpdateTestSet                    bool

	_lexmodelsv2Action                        []string
	_lexmodelsv2AggregationDuration           string
	_lexmodelsv2ApiMode                       string
	_lexmodelsv2Attributes                    string
	_lexmodelsv2BinBy                         string
	_lexmodelsv2BotAliasId                    string
	_lexmodelsv2BotAliasLocaleSettings        string
	_lexmodelsv2BotAliasName                  string
	_lexmodelsv2BotId                         string
	_lexmodelsv2BotMembers                    string
	_lexmodelsv2BotName                       string
	_lexmodelsv2BotRecommendationId           string
	_lexmodelsv2BotTags                       string
	_lexmodelsv2BotType                       string
	_lexmodelsv2BotVersion                    string
	_lexmodelsv2BotVersionLocaleSpecification string
	_lexmodelsv2CompositeSlotTypeSetting      string
	_lexmodelsv2Condition                     string
	_lexmodelsv2ConversationLogSettings       string
	_lexmodelsv2CustomVocabularyItemList      string
	_lexmodelsv2DataPrivacy                   string
	_lexmodelsv2Description                   string
	_lexmodelsv2DialogCodeHook                string
	_lexmodelsv2Effect                        string
	_lexmodelsv2EncryptionSetting             string
	_lexmodelsv2EndDateTime                   string
	_lexmodelsv2ErrorLogSettings              string
	_lexmodelsv2ExpectedRevisionId            string
	_lexmodelsv2ExportId                      string
	_lexmodelsv2ExternalSourceSetting         string
	_lexmodelsv2FileFormat                    string
	_lexmodelsv2FilePassword                  string
	_lexmodelsv2Filters                       string
	_lexmodelsv2FulfillmentCodeHook           string
	_lexmodelsv2GenerationDataSource          string
	_lexmodelsv2GenerationId                  string
	_lexmodelsv2GenerationInputPrompt         string
	_lexmodelsv2GenerativeAISettings          string
	_lexmodelsv2GroupBy                       string
	_lexmodelsv2IdleSessionTTLInSeconds       string
	_lexmodelsv2ImportId                      string
	_lexmodelsv2InitialResponseSetting        string
	_lexmodelsv2InputContexts                 string
	_lexmodelsv2IntentClosingSetting          string
	_lexmodelsv2IntentConfirmationSetting     string
	_lexmodelsv2IntentDisplayName             string
	_lexmodelsv2IntentId                      string
	_lexmodelsv2IntentName                    string
	_lexmodelsv2IntentPath                    string
	_lexmodelsv2KendraConfiguration           string
	_lexmodelsv2LocaleId                      string
	_lexmodelsv2MaxResults                    string
	_lexmodelsv2MergeStrategy                 string
	_lexmodelsv2Metrics                       string
	_lexmodelsv2MultipleValuesSetting         string
	_lexmodelsv2NextIndex                     string
	_lexmodelsv2NextToken                     string
	_lexmodelsv2NluIntentConfidenceThreshold  string
	_lexmodelsv2ObfuscationSetting            string
	_lexmodelsv2OutputContexts                string
	_lexmodelsv2ParentIntentSignature         string
	_lexmodelsv2ParentSlotTypeSignature       string
	_lexmodelsv2Policy                        string
	_lexmodelsv2Principal                     string
	_lexmodelsv2QInConnectIntentConfiguration string
	_lexmodelsv2QnAIntentConfiguration        string
	_lexmodelsv2ReplicaRegion                 string
	_lexmodelsv2ResourceARN                   string
	_lexmodelsv2ResourceSpecification         string
	_lexmodelsv2ResultFilterBy                string
	_lexmodelsv2RoleArn                       string
	_lexmodelsv2SampleUtterances              string
	_lexmodelsv2SearchOrder                   string
	_lexmodelsv2SentimentAnalysisSettings     string
	_lexmodelsv2SessionId                     string
	_lexmodelsv2SkipResourceInUseCheck        string
	_lexmodelsv2SlotId                        string
	_lexmodelsv2SlotName                      string
	_lexmodelsv2SlotPriorities                string
	_lexmodelsv2SlotTypeId                    string
	_lexmodelsv2SlotTypeName                  string
	_lexmodelsv2SlotTypeValues                string
	_lexmodelsv2SortBy                        string
	_lexmodelsv2SpeechDetectionSensitivity    string
	_lexmodelsv2SpeechRecognitionSettings     string
	_lexmodelsv2StartDateTime                 string
	_lexmodelsv2StatementId                   string
	_lexmodelsv2StorageLocation               string
	_lexmodelsv2SubSlotSetting                string
	_lexmodelsv2TagKeys                       []string
	_lexmodelsv2Tags                          string
	_lexmodelsv2Target                        string
	_lexmodelsv2TestBotAliasTags              string
	_lexmodelsv2TestExecutionId               string
	_lexmodelsv2TestExecutionModality         string
	_lexmodelsv2TestSetDiscrepancyReportId    string
	_lexmodelsv2TestSetGenerationId           string
	_lexmodelsv2TestSetId                     string
	_lexmodelsv2TestSetName                   string
	_lexmodelsv2TestSetTags                   string
	_lexmodelsv2TranscriptSourceSetting       string
	_lexmodelsv2UnifiedSpeechSettings         string
	_lexmodelsv2ValueElicitationSetting       string
	_lexmodelsv2ValueSelectionSetting         string
	_lexmodelsv2VoiceSettings                 string
)

// Create a batch of custom vocabulary items for a given bot locale's custom
// vocabulary.
func lexmodelsv2_BatchCreateCustomVocabularyItem(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.BatchCreateCustomVocabularyItemInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// CustomVocabularyItemList: []types.NewCustomVocabularyItem, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2CustomVocabularyItemList) > 0 {
		if err := assignInputField(input, "CustomVocabularyItemList", _lexmodelsv2CustomVocabularyItemList); err != nil {
			log.Errorf("invalid --custom-vocabulary-item-list: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}

	if resp, err := client.BatchCreateCustomVocabularyItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a batch of custom vocabulary items for a given bot locale's custom
// vocabulary.
func lexmodelsv2_BatchDeleteCustomVocabularyItem(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.BatchDeleteCustomVocabularyItemInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// CustomVocabularyItemList: []types.CustomVocabularyEntryId, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2CustomVocabularyItemList) > 0 {
		if err := assignInputField(input, "CustomVocabularyItemList", _lexmodelsv2CustomVocabularyItemList); err != nil {
			log.Errorf("invalid --custom-vocabulary-item-list: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}

	if resp, err := client.BatchDeleteCustomVocabularyItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a batch of custom vocabulary items for a given bot locale's custom
// vocabulary.
func lexmodelsv2_BatchUpdateCustomVocabularyItem(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.BatchUpdateCustomVocabularyItemInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// CustomVocabularyItemList: []types.CustomVocabularyItem, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2CustomVocabularyItemList) > 0 {
		if err := assignInputField(input, "CustomVocabularyItemList", _lexmodelsv2CustomVocabularyItemList); err != nil {
			log.Errorf("invalid --custom-vocabulary-item-list: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}

	if resp, err := client.BatchUpdateCustomVocabularyItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Builds a bot, its intents, and its slot types into a specific locale. A bot can
// be built into multiple locales. At runtime the locale is used to choose a
// specific build of the bot.
func lexmodelsv2_BuildBotLocale(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.BuildBotLocaleInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}

	if resp, err := client.BuildBotLocale(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Lex conversational bot.
func lexmodelsv2_CreateBot(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.CreateBotInput{
		// BotName: *string, // Required
		// DataPrivacy: *types.DataPrivacy, // Required
		// IdleSessionTTLInSeconds: *int32, // Required
		// RoleArn: *string, // Required
	}

	if len(_lexmodelsv2BotName) > 0 {
		input.BotName = aws.String(_lexmodelsv2BotName)
	}
	if len(_lexmodelsv2DataPrivacy) > 0 {
		if err := assignInputField(input, "DataPrivacy", _lexmodelsv2DataPrivacy); err != nil {
			log.Errorf("invalid --data-privacy: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2IdleSessionTTLInSeconds) > 0 {
		if err := assignInputField(input, "IdleSessionTTLInSeconds", _lexmodelsv2IdleSessionTTLInSeconds); err != nil {
			log.Errorf("invalid --idle-session-ttlin-seconds: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2RoleArn) > 0 {
		input.RoleArn = aws.String(_lexmodelsv2RoleArn)
	}
	if len(_lexmodelsv2BotMembers) > 0 {
		if err := assignInputField(input, "BotMembers", _lexmodelsv2BotMembers); err != nil {
			log.Errorf("invalid --bot-members: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2BotTags) > 0 {
		if err := assignInputField(input, "BotTags", _lexmodelsv2BotTags); err != nil {
			log.Errorf("invalid --bot-tags: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2BotType) > 0 {
		if err := assignInputField(input, "BotType", _lexmodelsv2BotType); err != nil {
			log.Errorf("invalid --bot-type: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Description) > 0 {
		input.Description = aws.String(_lexmodelsv2Description)
	}
	if len(_lexmodelsv2ErrorLogSettings) > 0 {
		if err := assignInputField(input, "ErrorLogSettings", _lexmodelsv2ErrorLogSettings); err != nil {
			log.Errorf("invalid --error-log-settings: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2TestBotAliasTags) > 0 {
		if err := assignInputField(input, "TestBotAliasTags", _lexmodelsv2TestBotAliasTags); err != nil {
			log.Errorf("invalid --test-bot-alias-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an alias for the specified version of a bot. Use an alias to enable you
// to change the version of a bot without updating applications that use the bot.
//
// For example, you can create an alias called "PROD" that your applications use
// to call the Amazon Lex bot.
func lexmodelsv2_CreateBotAlias(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.CreateBotAliasInput{
		// BotAliasName: *string, // Required
		// BotId: *string, // Required
	}

	if len(_lexmodelsv2BotAliasName) > 0 {
		input.BotAliasName = aws.String(_lexmodelsv2BotAliasName)
	}
	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotAliasLocaleSettings) > 0 {
		if err := assignInputField(input, "BotAliasLocaleSettings", _lexmodelsv2BotAliasLocaleSettings); err != nil {
			log.Errorf("invalid --bot-alias-locale-settings: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2ConversationLogSettings) > 0 {
		if err := assignInputField(input, "ConversationLogSettings", _lexmodelsv2ConversationLogSettings); err != nil {
			log.Errorf("invalid --conversation-log-settings: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Description) > 0 {
		input.Description = aws.String(_lexmodelsv2Description)
	}
	if len(_lexmodelsv2SentimentAnalysisSettings) > 0 {
		if err := assignInputField(input, "SentimentAnalysisSettings", _lexmodelsv2SentimentAnalysisSettings); err != nil {
			log.Errorf("invalid --sentiment-analysis-settings: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _lexmodelsv2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBotAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a locale in the bot. The locale contains the intents and slot types
// that the bot uses in conversations with users in the specified language and
// locale. You must add a locale to a bot before you can add intents and slot types
// to the bot.
func lexmodelsv2_CreateBotLocale(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.CreateBotLocaleInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// LocaleId: *string, // Required
		// NluIntentConfidenceThreshold: *float64, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2NluIntentConfidenceThreshold) > 0 {
		if err := assignInputField(input, "NluIntentConfidenceThreshold", _lexmodelsv2NluIntentConfidenceThreshold); err != nil {
			log.Errorf("invalid --nlu-intent-confidence-threshold: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Description) > 0 {
		input.Description = aws.String(_lexmodelsv2Description)
	}
	if len(_lexmodelsv2GenerativeAISettings) > 0 {
		if err := assignInputField(input, "GenerativeAISettings", _lexmodelsv2GenerativeAISettings); err != nil {
			log.Errorf("invalid --generative-ai-settings: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2SpeechDetectionSensitivity) > 0 {
		if err := assignInputField(input, "SpeechDetectionSensitivity", _lexmodelsv2SpeechDetectionSensitivity); err != nil {
			log.Errorf("invalid --speech-detection-sensitivity: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2SpeechRecognitionSettings) > 0 {
		if err := assignInputField(input, "SpeechRecognitionSettings", _lexmodelsv2SpeechRecognitionSettings); err != nil {
			log.Errorf("invalid --speech-recognition-settings: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2UnifiedSpeechSettings) > 0 {
		if err := assignInputField(input, "UnifiedSpeechSettings", _lexmodelsv2UnifiedSpeechSettings); err != nil {
			log.Errorf("invalid --unified-speech-settings: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2VoiceSettings) > 0 {
		if err := assignInputField(input, "VoiceSettings", _lexmodelsv2VoiceSettings); err != nil {
			log.Errorf("invalid --voice-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBotLocale(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Action to create a replication of the source bot in the secondary region.
func lexmodelsv2_CreateBotReplica(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.CreateBotReplicaInput{
		// BotId: *string, // Required
		// ReplicaRegion: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2ReplicaRegion) > 0 {
		input.ReplicaRegion = aws.String(_lexmodelsv2ReplicaRegion)
	}

	if resp, err := client.CreateBotReplica(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an immutable version of the bot. When you create the first version of a
// bot, Amazon Lex sets the version number to 1. Subsequent bot versions increase
// in an increment of 1. The version number will always represent the total number
// of versions created of the bot, not the current number of versions. If a bot
// version is deleted, that bot version number will not be reused.
func lexmodelsv2_CreateBotVersion(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.CreateBotVersionInput{
		// BotId: *string, // Required
		// BotVersionLocaleSpecification: map[string]types.BotVersionLocaleDetails, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersionLocaleSpecification) > 0 {
		if err := assignInputField(input, "BotVersionLocaleSpecification", _lexmodelsv2BotVersionLocaleSpecification); err != nil {
			log.Errorf("invalid --bot-version-locale-specification: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Description) > 0 {
		input.Description = aws.String(_lexmodelsv2Description)
	}

	if resp, err := client.CreateBotVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a zip archive containing the contents of a bot or a bot locale. The
// archive contains a directory structure that contains JSON files that define the
// bot.
//
// You can create an archive that contains the complete definition of a bot, or
// you can specify that the archive contain only the definition of a single bot
// locale.
//
// For more information about exporting bots, and about the structure of the
// export archive, see [Importing and exporting bots]
//
// [Importing and exporting bots]: https://docs.aws.amazon.com/lexv2/latest/dg/importing-exporting.html
func lexmodelsv2_CreateExport(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.CreateExportInput{
		// FileFormat: types.ImportExportFileFormat, // Required
		// ResourceSpecification: *types.ExportResourceSpecification, // Required
	}

	if len(_lexmodelsv2FileFormat) > 0 {
		if err := assignInputField(input, "FileFormat", _lexmodelsv2FileFormat); err != nil {
			log.Errorf("invalid --file-format: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2ResourceSpecification) > 0 {
		if err := assignInputField(input, "ResourceSpecification", _lexmodelsv2ResourceSpecification); err != nil {
			log.Errorf("invalid --resource-specification: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2FilePassword) > 0 {
		input.FilePassword = aws.String(_lexmodelsv2FilePassword)
	}

	if resp, err := client.CreateExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an intent.
// To define the interaction between the user and your bot, you define one or more
// intents. For example, for a pizza ordering bot you would create an OrderPizza
// intent.
//
// When you create an intent, you must provide a name. You can optionally provide
// the following:
//
// - Sample utterances. For example, "I want to order a pizza" and "Can I order
// a pizza." You can't provide utterances for built-in intents.
//
// - Information to be gathered. You specify slots for the information that you
// bot requests from the user. You can specify standard slot types, such as date
// and time, or custom slot types for your application.
//
// - How the intent is fulfilled. You can provide a Lambda function or configure
// the intent to return the intent information to your client application. If you
// use a Lambda function, Amazon Lex invokes the function when all of the intent
// information is available.
//
// - A confirmation prompt to send to the user to confirm an intent. For
// example, "Shall I order your pizza?"
//
// - A conclusion statement to send to the user after the intent is fulfilled.
// For example, "I ordered your pizza."
//
// - A follow-up prompt that asks the user for additional activity. For example,
// "Do you want a drink with your pizza?"
func lexmodelsv2_CreateIntent(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.CreateIntentInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// IntentName: *string, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2IntentName) > 0 {
		input.IntentName = aws.String(_lexmodelsv2IntentName)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2Description) > 0 {
		input.Description = aws.String(_lexmodelsv2Description)
	}
	if len(_lexmodelsv2DialogCodeHook) > 0 {
		if err := assignInputField(input, "DialogCodeHook", _lexmodelsv2DialogCodeHook); err != nil {
			log.Errorf("invalid --dialog-code-hook: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2FulfillmentCodeHook) > 0 {
		if err := assignInputField(input, "FulfillmentCodeHook", _lexmodelsv2FulfillmentCodeHook); err != nil {
			log.Errorf("invalid --fulfillment-code-hook: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2InitialResponseSetting) > 0 {
		if err := assignInputField(input, "InitialResponseSetting", _lexmodelsv2InitialResponseSetting); err != nil {
			log.Errorf("invalid --initial-response-setting: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2InputContexts) > 0 {
		if err := assignInputField(input, "InputContexts", _lexmodelsv2InputContexts); err != nil {
			log.Errorf("invalid --input-contexts: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2IntentClosingSetting) > 0 {
		if err := assignInputField(input, "IntentClosingSetting", _lexmodelsv2IntentClosingSetting); err != nil {
			log.Errorf("invalid --intent-closing-setting: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2IntentConfirmationSetting) > 0 {
		if err := assignInputField(input, "IntentConfirmationSetting", _lexmodelsv2IntentConfirmationSetting); err != nil {
			log.Errorf("invalid --intent-confirmation-setting: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2IntentDisplayName) > 0 {
		input.IntentDisplayName = aws.String(_lexmodelsv2IntentDisplayName)
	}
	if len(_lexmodelsv2KendraConfiguration) > 0 {
		if err := assignInputField(input, "KendraConfiguration", _lexmodelsv2KendraConfiguration); err != nil {
			log.Errorf("invalid --kendra-configuration: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2OutputContexts) > 0 {
		if err := assignInputField(input, "OutputContexts", _lexmodelsv2OutputContexts); err != nil {
			log.Errorf("invalid --output-contexts: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2ParentIntentSignature) > 0 {
		input.ParentIntentSignature = aws.String(_lexmodelsv2ParentIntentSignature)
	}
	if len(_lexmodelsv2QInConnectIntentConfiguration) > 0 {
		if err := assignInputField(input, "QInConnectIntentConfiguration", _lexmodelsv2QInConnectIntentConfiguration); err != nil {
			log.Errorf("invalid --qin-connect-intent-configuration: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2QnAIntentConfiguration) > 0 {
		if err := assignInputField(input, "QnAIntentConfiguration", _lexmodelsv2QnAIntentConfiguration); err != nil {
			log.Errorf("invalid --qna-intent-configuration: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2SampleUtterances) > 0 {
		if err := assignInputField(input, "SampleUtterances", _lexmodelsv2SampleUtterances); err != nil {
			log.Errorf("invalid --sample-utterances: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIntent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new resource policy with the specified policy statements.
func lexmodelsv2_CreateResourcePolicy(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.CreateResourcePolicyInput{
		// Policy: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_lexmodelsv2Policy) > 0 {
		input.Policy = aws.String(_lexmodelsv2Policy)
	}
	if len(_lexmodelsv2ResourceARN) > 0 {
		input.ResourceArn = aws.String(_lexmodelsv2ResourceARN)
	}

	if resp, err := client.CreateResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a new resource policy statement to a bot or bot alias. If a resource
// policy exists, the statement is added to the current resource policy. If a
// policy doesn't exist, a new policy is created.
//
// You can't create a resource policy statement that allows cross-account access.
//
// You need to add the CreateResourcePolicy or UpdateResourcePolicy action to the
// bot role in order to call the API.
func lexmodelsv2_CreateResourcePolicyStatement(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.CreateResourcePolicyStatementInput{
		// Action: []string, // Required
		// Effect: types.Effect, // Required
		// Principal: []types.Principal, // Required
		// ResourceArn: *string, // Required
		// StatementId: *string, // Required
	}

	if len(_lexmodelsv2Action) > 0 {
		input.Action = append([]string(nil), _lexmodelsv2Action...)
	}
	if len(_lexmodelsv2Effect) > 0 {
		if err := assignInputField(input, "Effect", _lexmodelsv2Effect); err != nil {
			log.Errorf("invalid --effect: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Principal) > 0 {
		if err := assignInputField(input, "Principal", _lexmodelsv2Principal); err != nil {
			log.Errorf("invalid --principal: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2ResourceARN) > 0 {
		input.ResourceArn = aws.String(_lexmodelsv2ResourceARN)
	}
	if len(_lexmodelsv2StatementId) > 0 {
		input.StatementId = aws.String(_lexmodelsv2StatementId)
	}
	if len(_lexmodelsv2Condition) > 0 {
		if err := assignInputField(input, "Condition", _lexmodelsv2Condition); err != nil {
			log.Errorf("invalid --condition: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2ExpectedRevisionId) > 0 {
		input.ExpectedRevisionId = aws.String(_lexmodelsv2ExpectedRevisionId)
	}

	if resp, err := client.CreateResourcePolicyStatement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a slot in an intent. A slot is a variable needed to fulfill an intent.
// For example, an OrderPizza intent might need slots for size, crust, and number
// of pizzas. For each slot, you define one or more utterances that Amazon Lex uses
// to elicit a response from the user.
func lexmodelsv2_CreateSlot(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.CreateSlotInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// IntentId: *string, // Required
		// LocaleId: *string, // Required
		// SlotName: *string, // Required
		// ValueElicitationSetting: *types.SlotValueElicitationSetting, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2IntentId) > 0 {
		input.IntentId = aws.String(_lexmodelsv2IntentId)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2SlotName) > 0 {
		input.SlotName = aws.String(_lexmodelsv2SlotName)
	}
	if len(_lexmodelsv2ValueElicitationSetting) > 0 {
		if err := assignInputField(input, "ValueElicitationSetting", _lexmodelsv2ValueElicitationSetting); err != nil {
			log.Errorf("invalid --value-elicitation-setting: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Description) > 0 {
		input.Description = aws.String(_lexmodelsv2Description)
	}
	if len(_lexmodelsv2MultipleValuesSetting) > 0 {
		if err := assignInputField(input, "MultipleValuesSetting", _lexmodelsv2MultipleValuesSetting); err != nil {
			log.Errorf("invalid --multiple-values-setting: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2ObfuscationSetting) > 0 {
		if err := assignInputField(input, "ObfuscationSetting", _lexmodelsv2ObfuscationSetting); err != nil {
			log.Errorf("invalid --obfuscation-setting: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2SlotTypeId) > 0 {
		input.SlotTypeId = aws.String(_lexmodelsv2SlotTypeId)
	}
	if len(_lexmodelsv2SubSlotSetting) > 0 {
		if err := assignInputField(input, "SubSlotSetting", _lexmodelsv2SubSlotSetting); err != nil {
			log.Errorf("invalid --sub-slot-setting: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSlot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a custom slot type
// To create a custom slot type, specify a name for the slot type and a set of
// enumeration values, the values that a slot of this type can assume.
func lexmodelsv2_CreateSlotType(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.CreateSlotTypeInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// LocaleId: *string, // Required
		// SlotTypeName: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2SlotTypeName) > 0 {
		input.SlotTypeName = aws.String(_lexmodelsv2SlotTypeName)
	}
	if len(_lexmodelsv2CompositeSlotTypeSetting) > 0 {
		if err := assignInputField(input, "CompositeSlotTypeSetting", _lexmodelsv2CompositeSlotTypeSetting); err != nil {
			log.Errorf("invalid --composite-slot-type-setting: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Description) > 0 {
		input.Description = aws.String(_lexmodelsv2Description)
	}
	if len(_lexmodelsv2ExternalSourceSetting) > 0 {
		if err := assignInputField(input, "ExternalSourceSetting", _lexmodelsv2ExternalSourceSetting); err != nil {
			log.Errorf("invalid --external-source-setting: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2ParentSlotTypeSignature) > 0 {
		input.ParentSlotTypeSignature = aws.String(_lexmodelsv2ParentSlotTypeSignature)
	}
	if len(_lexmodelsv2SlotTypeValues) > 0 {
		if err := assignInputField(input, "SlotTypeValues", _lexmodelsv2SlotTypeValues); err != nil {
			log.Errorf("invalid --slot-type-values: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2ValueSelectionSetting) > 0 {
		if err := assignInputField(input, "ValueSelectionSetting", _lexmodelsv2ValueSelectionSetting); err != nil {
			log.Errorf("invalid --value-selection-setting: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSlotType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a report that describes the differences between the bot and the test set.
func lexmodelsv2_CreateTestSetDiscrepancyReport(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.CreateTestSetDiscrepancyReportInput{
		// Target: *types.TestSetDiscrepancyReportResourceTarget, // Required
		// TestSetId: *string, // Required
	}

	if len(_lexmodelsv2Target) > 0 {
		if err := assignInputField(input, "Target", _lexmodelsv2Target); err != nil {
			log.Errorf("invalid --target: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2TestSetId) > 0 {
		input.TestSetId = aws.String(_lexmodelsv2TestSetId)
	}

	if resp, err := client.CreateTestSetDiscrepancyReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a pre-signed S3 write URL that you use to upload the zip archive when
// importing a bot or a bot locale.
func lexmodelsv2_CreateUploadUrl(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.CreateUploadUrlInput{}

	if resp, err := client.CreateUploadUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes all versions of a bot, including the Draft version. To delete a
// specific version, use the DeleteBotVersion operation.
//
// When you delete a bot, all of the resources contained in the bot are also
// deleted. Deleting a bot removes all locales, intents, slot, and slot types
// defined for the bot.
//
// If a bot has an alias, the DeleteBot operation returns a ResourceInUseException
// exception. If you want to delete the bot and the alias, set the
// skipResourceInUseCheck parameter to true .
func lexmodelsv2_DeleteBot(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DeleteBotInput{
		// BotId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2SkipResourceInUseCheck) > 0 {
		if err := assignInputField(input, "SkipResourceInUseCheck", _lexmodelsv2SkipResourceInUseCheck); err != nil {
			log.Errorf("invalid --skip-resource-in-use-check: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified bot alias.
func lexmodelsv2_DeleteBotAlias(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DeleteBotAliasInput{
		// BotAliasId: *string, // Required
		// BotId: *string, // Required
	}

	if len(_lexmodelsv2BotAliasId) > 0 {
		input.BotAliasId = aws.String(_lexmodelsv2BotAliasId)
	}
	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2SkipResourceInUseCheck) > 0 {
		if err := assignInputField(input, "SkipResourceInUseCheck", _lexmodelsv2SkipResourceInUseCheck); err != nil {
			log.Errorf("invalid --skip-resource-in-use-check: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteBotAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a locale from a bot.
// When you delete a locale, all intents, slots, and slot types defined for the
// locale are also deleted.
func lexmodelsv2_DeleteBotLocale(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DeleteBotLocaleInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}

	if resp, err := client.DeleteBotLocale(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The action to delete the replicated bot in the secondary region.
func lexmodelsv2_DeleteBotReplica(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DeleteBotReplicaInput{
		// BotId: *string, // Required
		// ReplicaRegion: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2ReplicaRegion) > 0 {
		input.ReplicaRegion = aws.String(_lexmodelsv2ReplicaRegion)
	}

	if resp, err := client.DeleteBotReplica(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specific version of a bot. To delete all versions of a bot, use the [DeleteBot]
// operation.
//
// [DeleteBot]: https://docs.aws.amazon.com/lexv2/latest/APIReference/API_DeleteBot.html
func lexmodelsv2_DeleteBotVersion(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DeleteBotVersionInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2SkipResourceInUseCheck) > 0 {
		if err := assignInputField(input, "SkipResourceInUseCheck", _lexmodelsv2SkipResourceInUseCheck); err != nil {
			log.Errorf("invalid --skip-resource-in-use-check: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteBotVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a custom vocabulary from the specified locale in the specified bot.
func lexmodelsv2_DeleteCustomVocabulary(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DeleteCustomVocabularyInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}

	if resp, err := client.DeleteCustomVocabulary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a previous export and the associated files stored in an S3 bucket.
func lexmodelsv2_DeleteExport(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DeleteExportInput{
		// ExportId: *string, // Required
	}

	if len(_lexmodelsv2ExportId) > 0 {
		input.ExportId = aws.String(_lexmodelsv2ExportId)
	}

	if resp, err := client.DeleteExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a previous import and the associated file stored in an S3 bucket.
func lexmodelsv2_DeleteImport(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DeleteImportInput{
		// ImportId: *string, // Required
	}

	if len(_lexmodelsv2ImportId) > 0 {
		input.ImportId = aws.String(_lexmodelsv2ImportId)
	}

	if resp, err := client.DeleteImport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified intent.
// Deleting an intent also deletes the slots associated with the intent.
func lexmodelsv2_DeleteIntent(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DeleteIntentInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// IntentId: *string, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2IntentId) > 0 {
		input.IntentId = aws.String(_lexmodelsv2IntentId)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}

	if resp, err := client.DeleteIntent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes an existing policy from a bot or bot alias. If the resource doesn't
// have a policy attached, Amazon Lex returns an exception.
func lexmodelsv2_DeleteResourcePolicy(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DeleteResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_lexmodelsv2ResourceARN) > 0 {
		input.ResourceArn = aws.String(_lexmodelsv2ResourceARN)
	}
	if len(_lexmodelsv2ExpectedRevisionId) > 0 {
		input.ExpectedRevisionId = aws.String(_lexmodelsv2ExpectedRevisionId)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a policy statement from a resource policy. If you delete the last
// statement from a policy, the policy is deleted. If you specify a statement ID
// that doesn't exist in the policy, or if the bot or bot alias doesn't have a
// policy attached, Amazon Lex returns an exception.
//
// You need to add the DeleteResourcePolicy or UpdateResourcePolicy action to the
// bot role in order to call the API.
func lexmodelsv2_DeleteResourcePolicyStatement(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DeleteResourcePolicyStatementInput{
		// ResourceArn: *string, // Required
		// StatementId: *string, // Required
	}

	if len(_lexmodelsv2ResourceARN) > 0 {
		input.ResourceArn = aws.String(_lexmodelsv2ResourceARN)
	}
	if len(_lexmodelsv2StatementId) > 0 {
		input.StatementId = aws.String(_lexmodelsv2StatementId)
	}
	if len(_lexmodelsv2ExpectedRevisionId) > 0 {
		input.ExpectedRevisionId = aws.String(_lexmodelsv2ExpectedRevisionId)
	}

	if resp, err := client.DeleteResourcePolicyStatement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified slot from an intent.
func lexmodelsv2_DeleteSlot(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DeleteSlotInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// IntentId: *string, // Required
		// LocaleId: *string, // Required
		// SlotId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2IntentId) > 0 {
		input.IntentId = aws.String(_lexmodelsv2IntentId)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2SlotId) > 0 {
		input.SlotId = aws.String(_lexmodelsv2SlotId)
	}

	if resp, err := client.DeleteSlot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a slot type from a bot locale.
// If a slot is using the slot type, Amazon Lex throws a ResourceInUseException
// exception. To avoid the exception, set the skipResourceInUseCheck parameter to
// true .
func lexmodelsv2_DeleteSlotType(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DeleteSlotTypeInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// LocaleId: *string, // Required
		// SlotTypeId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2SlotTypeId) > 0 {
		input.SlotTypeId = aws.String(_lexmodelsv2SlotTypeId)
	}
	if len(_lexmodelsv2SkipResourceInUseCheck) > 0 {
		if err := assignInputField(input, "SkipResourceInUseCheck", _lexmodelsv2SkipResourceInUseCheck); err != nil {
			log.Errorf("invalid --skip-resource-in-use-check: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteSlotType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The action to delete the selected test set.
func lexmodelsv2_DeleteTestSet(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DeleteTestSetInput{
		// TestSetId: *string, // Required
	}

	if len(_lexmodelsv2TestSetId) > 0 {
		input.TestSetId = aws.String(_lexmodelsv2TestSetId)
	}

	if resp, err := client.DeleteTestSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes stored utterances.
// Amazon Lex stores the utterances that users send to your bot. Utterances are
// stored for 15 days for use with the [ListAggregatedUtterances]operation, and then stored indefinitely for
// use in improving the ability of your bot to respond to user input..
//
// Use the DeleteUtterances operation to manually delete utterances for a specific
// session. When you use the DeleteUtterances operation, utterances stored for
// improving your bot's ability to respond to user input are deleted immediately.
// Utterances stored for use with the ListAggregatedUtterances operation are
// deleted after 15 days.
//
// [ListAggregatedUtterances]: https://docs.aws.amazon.com/lexv2/latest/APIReference/API_ListAggregatedUtterances.html
func lexmodelsv2_DeleteUtterances(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DeleteUtterancesInput{
		// BotId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2SessionId) > 0 {
		input.SessionId = aws.String(_lexmodelsv2SessionId)
	}

	if resp, err := client.DeleteUtterances(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides metadata information about a bot.
func lexmodelsv2_DescribeBot(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DescribeBotInput{
		// BotId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}

	if resp, err := client.DescribeBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get information about a specific bot alias.
func lexmodelsv2_DescribeBotAlias(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DescribeBotAliasInput{
		// BotAliasId: *string, // Required
		// BotId: *string, // Required
	}

	if len(_lexmodelsv2BotAliasId) > 0 {
		input.BotAliasId = aws.String(_lexmodelsv2BotAliasId)
	}
	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}

	if resp, err := client.DescribeBotAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the settings that a bot has for a specific locale.
func lexmodelsv2_DescribeBotLocale(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DescribeBotLocaleInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}

	if resp, err := client.DescribeBotLocale(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides metadata information about a bot recommendation. This information will
// enable you to get a description on the request inputs, to download associated
// transcripts after processing is complete, and to download intents and slot-types
// generated by the bot recommendation.
func lexmodelsv2_DescribeBotRecommendation(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DescribeBotRecommendationInput{
		// BotId: *string, // Required
		// BotRecommendationId: *string, // Required
		// BotVersion: *string, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotRecommendationId) > 0 {
		input.BotRecommendationId = aws.String(_lexmodelsv2BotRecommendationId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}

	if resp, err := client.DescribeBotRecommendation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Monitors the bot replication status through the UI console.
func lexmodelsv2_DescribeBotReplica(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DescribeBotReplicaInput{
		// BotId: *string, // Required
		// ReplicaRegion: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2ReplicaRegion) > 0 {
		input.ReplicaRegion = aws.String(_lexmodelsv2ReplicaRegion)
	}

	if resp, err := client.DescribeBotReplica(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a request to generate a bot through natural language
// description, made through the StartBotResource API. Use the
// generatedBotLocaleUrl to retrieve the Amazon S3 object containing the bot locale
// configuration. You can then modify and import this configuration.
func lexmodelsv2_DescribeBotResourceGeneration(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DescribeBotResourceGenerationInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// GenerationId: *string, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2GenerationId) > 0 {
		input.GenerationId = aws.String(_lexmodelsv2GenerationId)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}

	if resp, err := client.DescribeBotResourceGeneration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides metadata about a version of a bot.
func lexmodelsv2_DescribeBotVersion(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DescribeBotVersionInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}

	if resp, err := client.DescribeBotVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides metadata information about a custom vocabulary.
func lexmodelsv2_DescribeCustomVocabularyMetadata(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DescribeCustomVocabularyMetadataInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}

	if resp, err := client.DescribeCustomVocabularyMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specific export.
func lexmodelsv2_DescribeExport(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DescribeExportInput{
		// ExportId: *string, // Required
	}

	if len(_lexmodelsv2ExportId) > 0 {
		input.ExportId = aws.String(_lexmodelsv2ExportId)
	}

	if resp, err := client.DescribeExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specific import.
func lexmodelsv2_DescribeImport(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DescribeImportInput{
		// ImportId: *string, // Required
	}

	if len(_lexmodelsv2ImportId) > 0 {
		input.ImportId = aws.String(_lexmodelsv2ImportId)
	}

	if resp, err := client.DescribeImport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns metadata about an intent.
func lexmodelsv2_DescribeIntent(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DescribeIntentInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// IntentId: *string, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2IntentId) > 0 {
		input.IntentId = aws.String(_lexmodelsv2IntentId)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}

	if resp, err := client.DescribeIntent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the resource policy and policy revision for a bot or bot alias.
func lexmodelsv2_DescribeResourcePolicy(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DescribeResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_lexmodelsv2ResourceARN) > 0 {
		input.ResourceArn = aws.String(_lexmodelsv2ResourceARN)
	}

	if resp, err := client.DescribeResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets metadata information about a slot.
func lexmodelsv2_DescribeSlot(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DescribeSlotInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// IntentId: *string, // Required
		// LocaleId: *string, // Required
		// SlotId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2IntentId) > 0 {
		input.IntentId = aws.String(_lexmodelsv2IntentId)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2SlotId) > 0 {
		input.SlotId = aws.String(_lexmodelsv2SlotId)
	}

	if resp, err := client.DescribeSlot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets metadata information about a slot type.
func lexmodelsv2_DescribeSlotType(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DescribeSlotTypeInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// LocaleId: *string, // Required
		// SlotTypeId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2SlotTypeId) > 0 {
		input.SlotTypeId = aws.String(_lexmodelsv2SlotTypeId)
	}

	if resp, err := client.DescribeSlotType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets metadata information about the test execution.
func lexmodelsv2_DescribeTestExecution(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DescribeTestExecutionInput{
		// TestExecutionId: *string, // Required
	}

	if len(_lexmodelsv2TestExecutionId) > 0 {
		input.TestExecutionId = aws.String(_lexmodelsv2TestExecutionId)
	}

	if resp, err := client.DescribeTestExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets metadata information about the test set.
func lexmodelsv2_DescribeTestSet(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DescribeTestSetInput{
		// TestSetId: *string, // Required
	}

	if len(_lexmodelsv2TestSetId) > 0 {
		input.TestSetId = aws.String(_lexmodelsv2TestSetId)
	}

	if resp, err := client.DescribeTestSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets metadata information about the test set discrepancy report.
func lexmodelsv2_DescribeTestSetDiscrepancyReport(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DescribeTestSetDiscrepancyReportInput{
		// TestSetDiscrepancyReportId: *string, // Required
	}

	if len(_lexmodelsv2TestSetDiscrepancyReportId) > 0 {
		input.TestSetDiscrepancyReportId = aws.String(_lexmodelsv2TestSetDiscrepancyReportId)
	}

	if resp, err := client.DescribeTestSetDiscrepancyReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets metadata information about the test set generation.
func lexmodelsv2_DescribeTestSetGeneration(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.DescribeTestSetGenerationInput{
		// TestSetGenerationId: *string, // Required
	}

	if len(_lexmodelsv2TestSetGenerationId) > 0 {
		input.TestSetGenerationId = aws.String(_lexmodelsv2TestSetGenerationId)
	}

	if resp, err := client.DescribeTestSetGeneration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates sample utterances for an intent.
func lexmodelsv2_GenerateBotElement(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.GenerateBotElementInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// IntentId: *string, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2IntentId) > 0 {
		input.IntentId = aws.String(_lexmodelsv2IntentId)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}

	if resp, err := client.GenerateBotElement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The pre-signed Amazon S3 URL to download the test execution result artifacts.
func lexmodelsv2_GetTestExecutionArtifactsUrl(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.GetTestExecutionArtifactsUrlInput{
		// TestExecutionId: *string, // Required
	}

	if len(_lexmodelsv2TestExecutionId) > 0 {
		input.TestExecutionId = aws.String(_lexmodelsv2TestExecutionId)
	}

	if resp, err := client.GetTestExecutionArtifactsUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a list of utterances that users have sent to the bot.
// Utterances are aggregated by the text of the utterance. For example, all
// instances where customers used the phrase "I want to order pizza" are aggregated
// into the same line in the response.
//
// You can see both detected utterances and missed utterances. A detected
// utterance is where the bot properly recognized the utterance and activated the
// associated intent. A missed utterance was not recognized by the bot and didn't
// activate an intent.
//
// Utterances can be aggregated for a bot alias or for a bot version, but not both
// at the same time.
//
// Utterances statistics are not generated under the following conditions:
//
// - The childDirected field was set to true when the bot was created.
//
// - You are using slot obfuscation with one or more slots.
//
// - You opted out of participating in improving Amazon Lex.
func lexmodelsv2_ListAggregatedUtterances(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListAggregatedUtterancesInput{
		// AggregationDuration: *types.UtteranceAggregationDuration, // Required
		// BotId: *string, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2AggregationDuration) > 0 {
		if err := assignInputField(input, "AggregationDuration", _lexmodelsv2AggregationDuration); err != nil {
			log.Errorf("invalid --aggregation-duration: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2BotAliasId) > 0 {
		input.BotAliasId = aws.String(_lexmodelsv2BotAliasId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2Filters) > 0 {
		if err := assignInputField(input, "Filters", _lexmodelsv2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}
	if len(_lexmodelsv2SortBy) > 0 {
		if err := assignInputField(input, "SortBy", _lexmodelsv2SortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAggregatedUtterances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListAggregatedUtterancesOutput
	p := lexmodelsv2.NewListAggregatedUtterancesPaginator(client, input)
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

// The action to list the replicated bots created from the source bot alias.
func lexmodelsv2_ListBotAliasReplicas(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListBotAliasReplicasInput{
		// BotId: *string, // Required
		// ReplicaRegion: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2ReplicaRegion) > 0 {
		input.ReplicaRegion = aws.String(_lexmodelsv2ReplicaRegion)
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBotAliasReplicas(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListBotAliasReplicasOutput
	p := lexmodelsv2.NewListBotAliasReplicasPaginator(client, input)
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

// Gets a list of aliases for the specified bot.
func lexmodelsv2_ListBotAliases(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListBotAliasesInput{
		// BotId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBotAliases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListBotAliasesOutput
	p := lexmodelsv2.NewListBotAliasesPaginator(client, input)
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

// Gets a list of locales for the specified bot.
func lexmodelsv2_ListBotLocales(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListBotLocalesInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2Filters) > 0 {
		if err := assignInputField(input, "Filters", _lexmodelsv2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}
	if len(_lexmodelsv2SortBy) > 0 {
		if err := assignInputField(input, "SortBy", _lexmodelsv2SortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListBotLocales(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListBotLocalesOutput
	p := lexmodelsv2.NewListBotLocalesPaginator(client, input)
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

// Get a list of bot recommendations that meet the specified criteria.
func lexmodelsv2_ListBotRecommendations(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListBotRecommendationsInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBotRecommendations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListBotRecommendationsOutput
	p := lexmodelsv2.NewListBotRecommendationsPaginator(client, input)
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

// The action to list the replicated bots.
func lexmodelsv2_ListBotReplicas(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListBotReplicasInput{
		// BotId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}

	if resp, err := client.ListBotReplicas(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the generation requests made for a bot locale.
func lexmodelsv2_ListBotResourceGenerations(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListBotResourceGenerationsInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}
	if len(_lexmodelsv2SortBy) > 0 {
		if err := assignInputField(input, "SortBy", _lexmodelsv2SortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListBotResourceGenerations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListBotResourceGenerationsOutput
	p := lexmodelsv2.NewListBotResourceGenerationsPaginator(client, input)
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

// Contains information about all the versions replication statuses applicable for
// Global Resiliency.
func lexmodelsv2_ListBotVersionReplicas(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListBotVersionReplicasInput{
		// BotId: *string, // Required
		// ReplicaRegion: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2ReplicaRegion) > 0 {
		input.ReplicaRegion = aws.String(_lexmodelsv2ReplicaRegion)
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}
	if len(_lexmodelsv2SortBy) > 0 {
		if err := assignInputField(input, "SortBy", _lexmodelsv2SortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListBotVersionReplicas(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListBotVersionReplicasOutput
	p := lexmodelsv2.NewListBotVersionReplicasPaginator(client, input)
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

// Gets information about all of the versions of a bot.
// The ListBotVersions operation returns a summary of each version of a bot. For
// example, if a bot has three numbered versions, the ListBotVersions operation
// returns for summaries, one for each numbered version and one for the DRAFT
// version.
//
// The ListBotVersions operation always returns at least one version, the DRAFT
// version.
func lexmodelsv2_ListBotVersions(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListBotVersionsInput{
		// BotId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}
	if len(_lexmodelsv2SortBy) > 0 {
		if err := assignInputField(input, "SortBy", _lexmodelsv2SortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListBotVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListBotVersionsOutput
	p := lexmodelsv2.NewListBotVersionsPaginator(client, input)
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

// Gets a list of available bots.
func lexmodelsv2_ListBots(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListBotsInput{}

	if len(_lexmodelsv2Filters) > 0 {
		if err := assignInputField(input, "Filters", _lexmodelsv2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}
	if len(_lexmodelsv2SortBy) > 0 {
		if err := assignInputField(input, "SortBy", _lexmodelsv2SortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListBots(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListBotsOutput
	p := lexmodelsv2.NewListBotsPaginator(client, input)
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

// Gets a list of built-in intents provided by Amazon Lex that you can use in your
// bot.
//
// To use a built-in intent as a the base for your own intent, include the
// built-in intent signature in the parentIntentSignature parameter when you call
// the CreateIntent operation. For more information, see [CreateIntent].
//
// [CreateIntent]: https://docs.aws.amazon.com/lexv2/latest/APIReference/API_CreateIntent.html
func lexmodelsv2_ListBuiltInIntents(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListBuiltInIntentsInput{
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}
	if len(_lexmodelsv2SortBy) > 0 {
		if err := assignInputField(input, "SortBy", _lexmodelsv2SortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListBuiltInIntents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListBuiltInIntentsOutput
	p := lexmodelsv2.NewListBuiltInIntentsPaginator(client, input)
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

// Gets a list of built-in slot types that meet the specified criteria.
func lexmodelsv2_ListBuiltInSlotTypes(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListBuiltInSlotTypesInput{
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}
	if len(_lexmodelsv2SortBy) > 0 {
		if err := assignInputField(input, "SortBy", _lexmodelsv2SortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListBuiltInSlotTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListBuiltInSlotTypesOutput
	p := lexmodelsv2.NewListBuiltInSlotTypesPaginator(client, input)
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

// Paginated list of custom vocabulary items for a given bot locale's custom
// vocabulary.
func lexmodelsv2_ListCustomVocabularyItems(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListCustomVocabularyItemsInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCustomVocabularyItems(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListCustomVocabularyItemsOutput
	p := lexmodelsv2.NewListCustomVocabularyItemsPaginator(client, input)
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

// Lists the exports for a bot, bot locale, or custom vocabulary. Exports are kept
// in the list for 7 days.
func lexmodelsv2_ListExports(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListExportsInput{}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2Filters) > 0 {
		if err := assignInputField(input, "Filters", _lexmodelsv2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}
	if len(_lexmodelsv2SortBy) > 0 {
		if err := assignInputField(input, "SortBy", _lexmodelsv2SortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
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

	var results []*lexmodelsv2.ListExportsOutput
	p := lexmodelsv2.NewListExportsPaginator(client, input)
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

// Lists the imports for a bot, bot locale, or custom vocabulary. Imports are kept
// in the list for 7 days.
func lexmodelsv2_ListImports(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListImportsInput{}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2Filters) > 0 {
		if err := assignInputField(input, "Filters", _lexmodelsv2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}
	if len(_lexmodelsv2SortBy) > 0 {
		if err := assignInputField(input, "SortBy", _lexmodelsv2SortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
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

	var results []*lexmodelsv2.ListImportsOutput
	p := lexmodelsv2.NewListImportsPaginator(client, input)
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

// Retrieves summary metrics for the intents in your bot. The following fields are
// required:
//
// - metrics – A list of [AnalyticsIntentMetric]objects. In each object, use the name field to specify
// the metric to calculate, the statistic field to specify whether to calculate
// the Sum , Average , or Max number, and the order field to specify whether to
// sort the results in Ascending or Descending order.
//
// - startDateTime and endDateTime – Define a time range for which you want to
// retrieve results.
//
// Of the optional fields, you can organize the results in the following ways:
//
// - Use the filters field to filter the results, the groupBy field to specify
// categories by which to group the results, and the binBy field to specify time
// intervals by which to group the results.
//
// - Use the maxResults field to limit the number of results to return in a
// single response and the nextToken field to return the next batch of results if
// the response does not return the full set of results.
//
// Note that an order field exists in both binBy and metrics . You can specify only
// one order in a given request.
//
// [AnalyticsIntentMetric]: https://docs.aws.amazon.com/lexv2/latest/APIReference/API_AnalyticsIntentMetric.html
func lexmodelsv2_ListIntentMetrics(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListIntentMetricsInput{
		// BotId: *string, // Required
		// EndDateTime: *time.Time, // Required
		// Metrics: []types.AnalyticsIntentMetric, // Required
		// StartDateTime: *time.Time, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2EndDateTime) > 0 {
		if err := assignInputField(input, "EndDateTime", _lexmodelsv2EndDateTime); err != nil {
			log.Errorf("invalid --end-date-time: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Metrics) > 0 {
		if err := assignInputField(input, "Metrics", _lexmodelsv2Metrics); err != nil {
			log.Errorf("invalid --metrics: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2StartDateTime) > 0 {
		if err := assignInputField(input, "StartDateTime", _lexmodelsv2StartDateTime); err != nil {
			log.Errorf("invalid --start-date-time: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2BinBy) > 0 {
		if err := assignInputField(input, "BinBy", _lexmodelsv2BinBy); err != nil {
			log.Errorf("invalid --bin-by: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Filters) > 0 {
		if err := assignInputField(input, "Filters", _lexmodelsv2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2GroupBy) > 0 {
		if err := assignInputField(input, "GroupBy", _lexmodelsv2GroupBy); err != nil {
			log.Errorf("invalid --group-by: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIntentMetrics(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListIntentMetricsOutput
	p := lexmodelsv2.NewListIntentMetricsPaginator(client, input)
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

// Retrieves summary statistics for a path of intents that users take over
// sessions with your bot. The following fields are required:
//
// - startDateTime and endDateTime – Define a time range for which you want to
// retrieve results.
//
// - intentPath – Define an order of intents for which you want to retrieve
// metrics. Separate intents in the path with a forward slash. For example,
// populate the intentPath field with /BookCar/BookHotel to see details about how
// many times users invoked the BookCar and BookHotel intents in that order.
//
// Use the optional filters field to filter the results.
func lexmodelsv2_ListIntentPaths(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListIntentPathsInput{
		// BotId: *string, // Required
		// EndDateTime: *time.Time, // Required
		// IntentPath: *string, // Required
		// StartDateTime: *time.Time, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2EndDateTime) > 0 {
		if err := assignInputField(input, "EndDateTime", _lexmodelsv2EndDateTime); err != nil {
			log.Errorf("invalid --end-date-time: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2IntentPath) > 0 {
		input.IntentPath = aws.String(_lexmodelsv2IntentPath)
	}
	if len(_lexmodelsv2StartDateTime) > 0 {
		if err := assignInputField(input, "StartDateTime", _lexmodelsv2StartDateTime); err != nil {
			log.Errorf("invalid --start-date-time: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Filters) > 0 {
		if err := assignInputField(input, "Filters", _lexmodelsv2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListIntentPaths(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves summary metrics for the stages within intents in your bot. The
// following fields are required:
//
// - metrics – A list of [AnalyticsIntentStageMetric]objects. In each object, use the name field to specify
// the metric to calculate, the statistic field to specify whether to calculate
// the Sum , Average , or Max number, and the order field to specify whether to
// sort the results in Ascending or Descending order.
//
// - startDateTime and endDateTime – Define a time range for which you want to
// retrieve results.
//
// Of the optional fields, you can organize the results in the following ways:
//
// - Use the filters field to filter the results, the groupBy field to specify
// categories by which to group the results, and the binBy field to specify time
// intervals by which to group the results.
//
// - Use the maxResults field to limit the number of results to return in a
// single response and the nextToken field to return the next batch of results if
// the response does not return the full set of results.
//
// Note that an order field exists in both binBy and metrics . You can only specify
// one order in a given request.
//
// [AnalyticsIntentStageMetric]: https://docs.aws.amazon.com/lexv2/latest/APIReference/API_AnalyticsIntentStageMetric.html
func lexmodelsv2_ListIntentStageMetrics(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListIntentStageMetricsInput{
		// BotId: *string, // Required
		// EndDateTime: *time.Time, // Required
		// Metrics: []types.AnalyticsIntentStageMetric, // Required
		// StartDateTime: *time.Time, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2EndDateTime) > 0 {
		if err := assignInputField(input, "EndDateTime", _lexmodelsv2EndDateTime); err != nil {
			log.Errorf("invalid --end-date-time: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Metrics) > 0 {
		if err := assignInputField(input, "Metrics", _lexmodelsv2Metrics); err != nil {
			log.Errorf("invalid --metrics: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2StartDateTime) > 0 {
		if err := assignInputField(input, "StartDateTime", _lexmodelsv2StartDateTime); err != nil {
			log.Errorf("invalid --start-date-time: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2BinBy) > 0 {
		if err := assignInputField(input, "BinBy", _lexmodelsv2BinBy); err != nil {
			log.Errorf("invalid --bin-by: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Filters) > 0 {
		if err := assignInputField(input, "Filters", _lexmodelsv2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2GroupBy) > 0 {
		if err := assignInputField(input, "GroupBy", _lexmodelsv2GroupBy); err != nil {
			log.Errorf("invalid --group-by: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIntentStageMetrics(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListIntentStageMetricsOutput
	p := lexmodelsv2.NewListIntentStageMetricsPaginator(client, input)
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

// Get a list of intents that meet the specified criteria.
func lexmodelsv2_ListIntents(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListIntentsInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2Filters) > 0 {
		if err := assignInputField(input, "Filters", _lexmodelsv2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}
	if len(_lexmodelsv2SortBy) > 0 {
		if err := assignInputField(input, "SortBy", _lexmodelsv2SortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListIntents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListIntentsOutput
	p := lexmodelsv2.NewListIntentsPaginator(client, input)
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

// Gets a list of recommended intents provided by the bot recommendation that you
// can use in your bot. Intents in the response are ordered by relevance.
func lexmodelsv2_ListRecommendedIntents(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListRecommendedIntentsInput{
		// BotId: *string, // Required
		// BotRecommendationId: *string, // Required
		// BotVersion: *string, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotRecommendationId) > 0 {
		input.BotRecommendationId = aws.String(_lexmodelsv2BotRecommendationId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRecommendedIntents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListRecommendedIntentsOutput
	p := lexmodelsv2.NewListRecommendedIntentsPaginator(client, input)
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

// Retrieves a list of metadata for individual user sessions with your bot. The
// startDateTime and endDateTime fields are required. These fields define a time
// range for which you want to retrieve results. Of the optional fields, you can
// organize the results in the following ways:
//
// - Use the filters field to filter the results and the sortBy field to specify
// the values by which to sort the results.
//
// - Use the maxResults field to limit the number of results to return in a
// single response and the nextToken field to return the next batch of results if
// the response does not return the full set of results.
func lexmodelsv2_ListSessionAnalyticsData(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListSessionAnalyticsDataInput{
		// BotId: *string, // Required
		// EndDateTime: *time.Time, // Required
		// StartDateTime: *time.Time, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2EndDateTime) > 0 {
		if err := assignInputField(input, "EndDateTime", _lexmodelsv2EndDateTime); err != nil {
			log.Errorf("invalid --end-date-time: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2StartDateTime) > 0 {
		if err := assignInputField(input, "StartDateTime", _lexmodelsv2StartDateTime); err != nil {
			log.Errorf("invalid --start-date-time: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Filters) > 0 {
		if err := assignInputField(input, "Filters", _lexmodelsv2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}
	if len(_lexmodelsv2SortBy) > 0 {
		if err := assignInputField(input, "SortBy", _lexmodelsv2SortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSessionAnalyticsData(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListSessionAnalyticsDataOutput
	p := lexmodelsv2.NewListSessionAnalyticsDataPaginator(client, input)
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

// Retrieves summary metrics for the user sessions with your bot. The following
// fields are required:
//
// - metrics – A list of [AnalyticsSessionMetric]objects. In each object, use the name field to specify
// the metric to calculate, the statistic field to specify whether to calculate
// the Sum , Average , or Max number, and the order field to specify whether to
// sort the results in Ascending or Descending order.
//
// - startDateTime and endDateTime – Define a time range for which you want to
// retrieve results.
//
// Of the optional fields, you can organize the results in the following ways:
//
// - Use the filters field to filter the results, the groupBy field to specify
// categories by which to group the results, and the binBy field to specify time
// intervals by which to group the results.
//
// - Use the maxResults field to limit the number of results to return in a
// single response and the nextToken field to return the next batch of results if
// the response does not return the full set of results.
//
// Note that an order field exists in both binBy and metrics . Currently, you can
// specify it in either field, but not in both.
//
// [AnalyticsSessionMetric]: https://docs.aws.amazon.com/lexv2/latest/APIReference/API_AnalyticsSessionMetric.html
func lexmodelsv2_ListSessionMetrics(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListSessionMetricsInput{
		// BotId: *string, // Required
		// EndDateTime: *time.Time, // Required
		// Metrics: []types.AnalyticsSessionMetric, // Required
		// StartDateTime: *time.Time, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2EndDateTime) > 0 {
		if err := assignInputField(input, "EndDateTime", _lexmodelsv2EndDateTime); err != nil {
			log.Errorf("invalid --end-date-time: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Metrics) > 0 {
		if err := assignInputField(input, "Metrics", _lexmodelsv2Metrics); err != nil {
			log.Errorf("invalid --metrics: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2StartDateTime) > 0 {
		if err := assignInputField(input, "StartDateTime", _lexmodelsv2StartDateTime); err != nil {
			log.Errorf("invalid --start-date-time: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2BinBy) > 0 {
		if err := assignInputField(input, "BinBy", _lexmodelsv2BinBy); err != nil {
			log.Errorf("invalid --bin-by: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Filters) > 0 {
		if err := assignInputField(input, "Filters", _lexmodelsv2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2GroupBy) > 0 {
		if err := assignInputField(input, "GroupBy", _lexmodelsv2GroupBy); err != nil {
			log.Errorf("invalid --group-by: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSessionMetrics(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListSessionMetricsOutput
	p := lexmodelsv2.NewListSessionMetricsPaginator(client, input)
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

// Gets a list of slot types that match the specified criteria.
func lexmodelsv2_ListSlotTypes(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListSlotTypesInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2Filters) > 0 {
		if err := assignInputField(input, "Filters", _lexmodelsv2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}
	if len(_lexmodelsv2SortBy) > 0 {
		if err := assignInputField(input, "SortBy", _lexmodelsv2SortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSlotTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListSlotTypesOutput
	p := lexmodelsv2.NewListSlotTypesPaginator(client, input)
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

// Gets a list of slots that match the specified criteria.
func lexmodelsv2_ListSlots(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListSlotsInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// IntentId: *string, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2IntentId) > 0 {
		input.IntentId = aws.String(_lexmodelsv2IntentId)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2Filters) > 0 {
		if err := assignInputField(input, "Filters", _lexmodelsv2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}
	if len(_lexmodelsv2SortBy) > 0 {
		if err := assignInputField(input, "SortBy", _lexmodelsv2SortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSlots(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListSlotsOutput
	p := lexmodelsv2.NewListSlotsPaginator(client, input)
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

// Gets a list of tags associated with a resource. Only bots, bot aliases, and bot
// channels can have tags associated with them.
func lexmodelsv2_ListTagsForResource(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_lexmodelsv2ResourceARN) > 0 {
		input.ResourceARN = aws.String(_lexmodelsv2ResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of test execution result items.
func lexmodelsv2_ListTestExecutionResultItems(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListTestExecutionResultItemsInput{
		// ResultFilterBy: *types.TestExecutionResultFilterBy, // Required
		// TestExecutionId: *string, // Required
	}

	if len(_lexmodelsv2ResultFilterBy) > 0 {
		if err := assignInputField(input, "ResultFilterBy", _lexmodelsv2ResultFilterBy); err != nil {
			log.Errorf("invalid --result-filter-by: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2TestExecutionId) > 0 {
		input.TestExecutionId = aws.String(_lexmodelsv2TestExecutionId)
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTestExecutionResultItems(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListTestExecutionResultItemsOutput
	p := lexmodelsv2.NewListTestExecutionResultItemsPaginator(client, input)
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

// The list of test set executions.
func lexmodelsv2_ListTestExecutions(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListTestExecutionsInput{}

	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}
	if len(_lexmodelsv2SortBy) > 0 {
		if err := assignInputField(input, "SortBy", _lexmodelsv2SortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTestExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListTestExecutionsOutput
	p := lexmodelsv2.NewListTestExecutionsPaginator(client, input)
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

// The list of test set records.
func lexmodelsv2_ListTestSetRecords(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListTestSetRecordsInput{
		// TestSetId: *string, // Required
	}

	if len(_lexmodelsv2TestSetId) > 0 {
		input.TestSetId = aws.String(_lexmodelsv2TestSetId)
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTestSetRecords(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListTestSetRecordsOutput
	p := lexmodelsv2.NewListTestSetRecordsPaginator(client, input)
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

// The list of the test sets
func lexmodelsv2_ListTestSets(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListTestSetsInput{}

	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}
	if len(_lexmodelsv2SortBy) > 0 {
		if err := assignInputField(input, "SortBy", _lexmodelsv2SortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTestSets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListTestSetsOutput
	p := lexmodelsv2.NewListTestSetsPaginator(client, input)
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

// To use this API operation, your IAM role must have permissions to perform the [ListAggregatedUtterances]
// operation, which provides access to utterance-related analytics. See [Viewing utterance statistics]for the
// IAM policy to apply to the IAM role.
//
// Retrieves a list of metadata for individual user utterances to your bot. The
// following fields are required:
//
// - startDateTime and endDateTime – Define a time range for which you want to
// retrieve results.
//
// Of the optional fields, you can organize the results in the following ways:
//
// - Use the filters field to filter the results and the sortBy field to specify
// the values by which to sort the results.
//
// - Use the maxResults field to limit the number of results to return in a
// single response and the nextToken field to return the next batch of results if
// the response does not return the full set of results.
//
// [Viewing utterance statistics]: https://docs.aws.amazon.com/lexv2/latest/dg/monitoring-utterances.html
// [ListAggregatedUtterances]: https://docs.aws.amazon.com/lexv2/latest/APIReference/API_ListAggregatedUtterances.html
func lexmodelsv2_ListUtteranceAnalyticsData(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListUtteranceAnalyticsDataInput{
		// BotId: *string, // Required
		// EndDateTime: *time.Time, // Required
		// StartDateTime: *time.Time, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2EndDateTime) > 0 {
		if err := assignInputField(input, "EndDateTime", _lexmodelsv2EndDateTime); err != nil {
			log.Errorf("invalid --end-date-time: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2StartDateTime) > 0 {
		if err := assignInputField(input, "StartDateTime", _lexmodelsv2StartDateTime); err != nil {
			log.Errorf("invalid --start-date-time: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Filters) > 0 {
		if err := assignInputField(input, "Filters", _lexmodelsv2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}
	if len(_lexmodelsv2SortBy) > 0 {
		if err := assignInputField(input, "SortBy", _lexmodelsv2SortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListUtteranceAnalyticsData(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListUtteranceAnalyticsDataOutput
	p := lexmodelsv2.NewListUtteranceAnalyticsDataPaginator(client, input)
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

// To use this API operation, your IAM role must have permissions to perform the [ListAggregatedUtterances]
// operation, which provides access to utterance-related analytics. See [Viewing utterance statistics]for the
// IAM policy to apply to the IAM role.
//
// Retrieves summary metrics for the utterances in your bot. The following fields
// are required:
//
// - metrics – A list of [AnalyticsUtteranceMetric]objects. In each object, use the name field to specify
// the metric to calculate, the statistic field to specify whether to calculate
// the Sum , Average , or Max number, and the order field to specify whether to
// sort the results in Ascending or Descending order.
//
// - startDateTime and endDateTime – Define a time range for which you want to
// retrieve results.
//
// Of the optional fields, you can organize the results in the following ways:
//
// - Use the filters field to filter the results, the groupBy field to specify
// categories by which to group the results, and the binBy field to specify time
// intervals by which to group the results.
//
// - Use the maxResults field to limit the number of results to return in a
// single response and the nextToken field to return the next batch of results if
// the response does not return the full set of results.
//
// Note that an order field exists in both binBy and metrics . Currently, you can
// specify it in either field, but not in both.
//
// [AnalyticsUtteranceMetric]: https://docs.aws.amazon.com/lexv2/latest/APIReference/API_AnalyticsUtteranceMetric.html
// [Viewing utterance statistics]: https://docs.aws.amazon.com/lexv2/latest/dg/monitoring-utterances.html
// [ListAggregatedUtterances]: https://docs.aws.amazon.com/lexv2/latest/APIReference/API_ListAggregatedUtterances.html
func lexmodelsv2_ListUtteranceMetrics(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.ListUtteranceMetricsInput{
		// BotId: *string, // Required
		// EndDateTime: *time.Time, // Required
		// Metrics: []types.AnalyticsUtteranceMetric, // Required
		// StartDateTime: *time.Time, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2EndDateTime) > 0 {
		if err := assignInputField(input, "EndDateTime", _lexmodelsv2EndDateTime); err != nil {
			log.Errorf("invalid --end-date-time: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Metrics) > 0 {
		if err := assignInputField(input, "Metrics", _lexmodelsv2Metrics); err != nil {
			log.Errorf("invalid --metrics: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2StartDateTime) > 0 {
		if err := assignInputField(input, "StartDateTime", _lexmodelsv2StartDateTime); err != nil {
			log.Errorf("invalid --start-date-time: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Attributes) > 0 {
		if err := assignInputField(input, "Attributes", _lexmodelsv2Attributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2BinBy) > 0 {
		if err := assignInputField(input, "BinBy", _lexmodelsv2BinBy); err != nil {
			log.Errorf("invalid --bin-by: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Filters) > 0 {
		if err := assignInputField(input, "Filters", _lexmodelsv2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2GroupBy) > 0 {
		if err := assignInputField(input, "GroupBy", _lexmodelsv2GroupBy); err != nil {
			log.Errorf("invalid --group-by: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextToken) > 0 {
		input.NextToken = aws.String(_lexmodelsv2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListUtteranceMetrics(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelsv2.ListUtteranceMetricsOutput
	p := lexmodelsv2.NewListUtteranceMetricsPaginator(client, input)
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

// Search for associated transcripts that meet the specified criteria.
func lexmodelsv2_SearchAssociatedTranscripts(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.SearchAssociatedTranscriptsInput{
		// BotId: *string, // Required
		// BotRecommendationId: *string, // Required
		// BotVersion: *string, // Required
		// Filters: []types.AssociatedTranscriptFilter, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotRecommendationId) > 0 {
		input.BotRecommendationId = aws.String(_lexmodelsv2BotRecommendationId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2Filters) > 0 {
		if err := assignInputField(input, "Filters", _lexmodelsv2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelsv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2NextIndex) > 0 {
		if err := assignInputField(input, "NextIndex", _lexmodelsv2NextIndex); err != nil {
			log.Errorf("invalid --next-index: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2SearchOrder) > 0 {
		if err := assignInputField(input, "SearchOrder", _lexmodelsv2SearchOrder); err != nil {
			log.Errorf("invalid --search-order: %s", err.Error())
			return
		}
	}

	if resp, err := client.SearchAssociatedTranscripts(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this to provide your transcript data, and to start the bot recommendation
// process.
func lexmodelsv2_StartBotRecommendation(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.StartBotRecommendationInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// LocaleId: *string, // Required
		// TranscriptSourceSetting: *types.TranscriptSourceSetting, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2TranscriptSourceSetting) > 0 {
		if err := assignInputField(input, "TranscriptSourceSetting", _lexmodelsv2TranscriptSourceSetting); err != nil {
			log.Errorf("invalid --transcript-source-setting: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2EncryptionSetting) > 0 {
		if err := assignInputField(input, "EncryptionSetting", _lexmodelsv2EncryptionSetting); err != nil {
			log.Errorf("invalid --encryption-setting: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartBotRecommendation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a request for the descriptive bot builder to generate a bot locale
// configuration based on the prompt you provide it. After you make this call, use
// the DescribeBotResourceGeneration operation to check on the status of the
// generation and for the generatedBotLocaleUrl when the generation is complete.
// Use that value to retrieve the Amazon S3 object containing the bot locale
// configuration. You can then modify and import this configuration.
func lexmodelsv2_StartBotResourceGeneration(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.StartBotResourceGenerationInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// GenerationInputPrompt: *string, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2GenerationInputPrompt) > 0 {
		input.GenerationInputPrompt = aws.String(_lexmodelsv2GenerationInputPrompt)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}

	if resp, err := client.StartBotResourceGeneration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts importing a bot, bot locale, or custom vocabulary from a zip archive
// that you uploaded to an S3 bucket.
func lexmodelsv2_StartImport(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.StartImportInput{
		// ImportId: *string, // Required
		// MergeStrategy: types.MergeStrategy, // Required
		// ResourceSpecification: *types.ImportResourceSpecification, // Required
	}

	if len(_lexmodelsv2ImportId) > 0 {
		input.ImportId = aws.String(_lexmodelsv2ImportId)
	}
	if len(_lexmodelsv2MergeStrategy) > 0 {
		if err := assignInputField(input, "MergeStrategy", _lexmodelsv2MergeStrategy); err != nil {
			log.Errorf("invalid --merge-strategy: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2ResourceSpecification) > 0 {
		if err := assignInputField(input, "ResourceSpecification", _lexmodelsv2ResourceSpecification); err != nil {
			log.Errorf("invalid --resource-specification: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2FilePassword) > 0 {
		input.FilePassword = aws.String(_lexmodelsv2FilePassword)
	}

	if resp, err := client.StartImport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The action to start test set execution.
func lexmodelsv2_StartTestExecution(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.StartTestExecutionInput{
		// ApiMode: types.TestExecutionApiMode, // Required
		// Target: *types.TestExecutionTarget, // Required
		// TestSetId: *string, // Required
	}

	if len(_lexmodelsv2ApiMode) > 0 {
		if err := assignInputField(input, "ApiMode", _lexmodelsv2ApiMode); err != nil {
			log.Errorf("invalid --api-mode: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Target) > 0 {
		if err := assignInputField(input, "Target", _lexmodelsv2Target); err != nil {
			log.Errorf("invalid --target: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2TestSetId) > 0 {
		input.TestSetId = aws.String(_lexmodelsv2TestSetId)
	}
	if len(_lexmodelsv2TestExecutionModality) > 0 {
		if err := assignInputField(input, "TestExecutionModality", _lexmodelsv2TestExecutionModality); err != nil {
			log.Errorf("invalid --test-execution-modality: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartTestExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The action to start the generation of test set.
func lexmodelsv2_StartTestSetGeneration(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.StartTestSetGenerationInput{
		// GenerationDataSource: *types.TestSetGenerationDataSource, // Required
		// RoleArn: *string, // Required
		// StorageLocation: *types.TestSetStorageLocation, // Required
		// TestSetName: *string, // Required
	}

	if len(_lexmodelsv2GenerationDataSource) > 0 {
		if err := assignInputField(input, "GenerationDataSource", _lexmodelsv2GenerationDataSource); err != nil {
			log.Errorf("invalid --generation-data-source: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2RoleArn) > 0 {
		input.RoleArn = aws.String(_lexmodelsv2RoleArn)
	}
	if len(_lexmodelsv2StorageLocation) > 0 {
		if err := assignInputField(input, "StorageLocation", _lexmodelsv2StorageLocation); err != nil {
			log.Errorf("invalid --storage-location: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2TestSetName) > 0 {
		input.TestSetName = aws.String(_lexmodelsv2TestSetName)
	}
	if len(_lexmodelsv2Description) > 0 {
		input.Description = aws.String(_lexmodelsv2Description)
	}
	if len(_lexmodelsv2TestSetTags) > 0 {
		if err := assignInputField(input, "TestSetTags", _lexmodelsv2TestSetTags); err != nil {
			log.Errorf("invalid --test-set-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartTestSetGeneration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stop an already running Bot Recommendation request.
func lexmodelsv2_StopBotRecommendation(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.StopBotRecommendationInput{
		// BotId: *string, // Required
		// BotRecommendationId: *string, // Required
		// BotVersion: *string, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotRecommendationId) > 0 {
		input.BotRecommendationId = aws.String(_lexmodelsv2BotRecommendationId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}

	if resp, err := client.StopBotRecommendation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the specified tags to the specified resource. If a tag key already exists,
// the existing value is replaced with the new value.
func lexmodelsv2_TagResource(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_lexmodelsv2ResourceARN) > 0 {
		input.ResourceARN = aws.String(_lexmodelsv2ResourceARN)
	}
	if len(_lexmodelsv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _lexmodelsv2Tags); err != nil {
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

// Removes tags from a bot, bot alias, or bot channel.
func lexmodelsv2_UntagResource(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_lexmodelsv2ResourceARN) > 0 {
		input.ResourceARN = aws.String(_lexmodelsv2ResourceARN)
	}
	if len(_lexmodelsv2TagKeys) > 0 {
		input.TagKeys = append([]string(nil), _lexmodelsv2TagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an existing bot.
func lexmodelsv2_UpdateBot(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.UpdateBotInput{
		// BotId: *string, // Required
		// BotName: *string, // Required
		// DataPrivacy: *types.DataPrivacy, // Required
		// IdleSessionTTLInSeconds: *int32, // Required
		// RoleArn: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotName) > 0 {
		input.BotName = aws.String(_lexmodelsv2BotName)
	}
	if len(_lexmodelsv2DataPrivacy) > 0 {
		if err := assignInputField(input, "DataPrivacy", _lexmodelsv2DataPrivacy); err != nil {
			log.Errorf("invalid --data-privacy: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2IdleSessionTTLInSeconds) > 0 {
		if err := assignInputField(input, "IdleSessionTTLInSeconds", _lexmodelsv2IdleSessionTTLInSeconds); err != nil {
			log.Errorf("invalid --idle-session-ttlin-seconds: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2RoleArn) > 0 {
		input.RoleArn = aws.String(_lexmodelsv2RoleArn)
	}
	if len(_lexmodelsv2BotMembers) > 0 {
		if err := assignInputField(input, "BotMembers", _lexmodelsv2BotMembers); err != nil {
			log.Errorf("invalid --bot-members: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2BotType) > 0 {
		if err := assignInputField(input, "BotType", _lexmodelsv2BotType); err != nil {
			log.Errorf("invalid --bot-type: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Description) > 0 {
		input.Description = aws.String(_lexmodelsv2Description)
	}
	if len(_lexmodelsv2ErrorLogSettings) > 0 {
		if err := assignInputField(input, "ErrorLogSettings", _lexmodelsv2ErrorLogSettings); err != nil {
			log.Errorf("invalid --error-log-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an existing bot alias.
func lexmodelsv2_UpdateBotAlias(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.UpdateBotAliasInput{
		// BotAliasId: *string, // Required
		// BotAliasName: *string, // Required
		// BotId: *string, // Required
	}

	if len(_lexmodelsv2BotAliasId) > 0 {
		input.BotAliasId = aws.String(_lexmodelsv2BotAliasId)
	}
	if len(_lexmodelsv2BotAliasName) > 0 {
		input.BotAliasName = aws.String(_lexmodelsv2BotAliasName)
	}
	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotAliasLocaleSettings) > 0 {
		if err := assignInputField(input, "BotAliasLocaleSettings", _lexmodelsv2BotAliasLocaleSettings); err != nil {
			log.Errorf("invalid --bot-alias-locale-settings: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2ConversationLogSettings) > 0 {
		if err := assignInputField(input, "ConversationLogSettings", _lexmodelsv2ConversationLogSettings); err != nil {
			log.Errorf("invalid --conversation-log-settings: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Description) > 0 {
		input.Description = aws.String(_lexmodelsv2Description)
	}
	if len(_lexmodelsv2SentimentAnalysisSettings) > 0 {
		if err := assignInputField(input, "SentimentAnalysisSettings", _lexmodelsv2SentimentAnalysisSettings); err != nil {
			log.Errorf("invalid --sentiment-analysis-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBotAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the settings that a bot has for a specific locale.
func lexmodelsv2_UpdateBotLocale(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.UpdateBotLocaleInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// LocaleId: *string, // Required
		// NluIntentConfidenceThreshold: *float64, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2NluIntentConfidenceThreshold) > 0 {
		if err := assignInputField(input, "NluIntentConfidenceThreshold", _lexmodelsv2NluIntentConfidenceThreshold); err != nil {
			log.Errorf("invalid --nlu-intent-confidence-threshold: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Description) > 0 {
		input.Description = aws.String(_lexmodelsv2Description)
	}
	if len(_lexmodelsv2GenerativeAISettings) > 0 {
		if err := assignInputField(input, "GenerativeAISettings", _lexmodelsv2GenerativeAISettings); err != nil {
			log.Errorf("invalid --generative-ai-settings: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2SpeechDetectionSensitivity) > 0 {
		if err := assignInputField(input, "SpeechDetectionSensitivity", _lexmodelsv2SpeechDetectionSensitivity); err != nil {
			log.Errorf("invalid --speech-detection-sensitivity: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2SpeechRecognitionSettings) > 0 {
		if err := assignInputField(input, "SpeechRecognitionSettings", _lexmodelsv2SpeechRecognitionSettings); err != nil {
			log.Errorf("invalid --speech-recognition-settings: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2UnifiedSpeechSettings) > 0 {
		if err := assignInputField(input, "UnifiedSpeechSettings", _lexmodelsv2UnifiedSpeechSettings); err != nil {
			log.Errorf("invalid --unified-speech-settings: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2VoiceSettings) > 0 {
		if err := assignInputField(input, "VoiceSettings", _lexmodelsv2VoiceSettings); err != nil {
			log.Errorf("invalid --voice-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBotLocale(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing bot recommendation request.
func lexmodelsv2_UpdateBotRecommendation(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.UpdateBotRecommendationInput{
		// BotId: *string, // Required
		// BotRecommendationId: *string, // Required
		// BotVersion: *string, // Required
		// EncryptionSetting: *types.EncryptionSetting, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotRecommendationId) > 0 {
		input.BotRecommendationId = aws.String(_lexmodelsv2BotRecommendationId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2EncryptionSetting) > 0 {
		if err := assignInputField(input, "EncryptionSetting", _lexmodelsv2EncryptionSetting); err != nil {
			log.Errorf("invalid --encryption-setting: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}

	if resp, err := client.UpdateBotRecommendation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the password used to protect an export zip archive.
// The password is not required. If you don't supply a password, Amazon Lex
// generates a zip file that is not protected by a password. This is the archive
// that is available at the pre-signed S3 URL provided by the [DescribeExport]operation.
//
// [DescribeExport]: https://docs.aws.amazon.com/lexv2/latest/APIReference/API_DescribeExport.html
func lexmodelsv2_UpdateExport(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.UpdateExportInput{
		// ExportId: *string, // Required
	}

	if len(_lexmodelsv2ExportId) > 0 {
		input.ExportId = aws.String(_lexmodelsv2ExportId)
	}
	if len(_lexmodelsv2FilePassword) > 0 {
		input.FilePassword = aws.String(_lexmodelsv2FilePassword)
	}

	if resp, err := client.UpdateExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the settings for an intent.
func lexmodelsv2_UpdateIntent(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.UpdateIntentInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// IntentId: *string, // Required
		// IntentName: *string, // Required
		// LocaleId: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2IntentId) > 0 {
		input.IntentId = aws.String(_lexmodelsv2IntentId)
	}
	if len(_lexmodelsv2IntentName) > 0 {
		input.IntentName = aws.String(_lexmodelsv2IntentName)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2Description) > 0 {
		input.Description = aws.String(_lexmodelsv2Description)
	}
	if len(_lexmodelsv2DialogCodeHook) > 0 {
		if err := assignInputField(input, "DialogCodeHook", _lexmodelsv2DialogCodeHook); err != nil {
			log.Errorf("invalid --dialog-code-hook: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2FulfillmentCodeHook) > 0 {
		if err := assignInputField(input, "FulfillmentCodeHook", _lexmodelsv2FulfillmentCodeHook); err != nil {
			log.Errorf("invalid --fulfillment-code-hook: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2InitialResponseSetting) > 0 {
		if err := assignInputField(input, "InitialResponseSetting", _lexmodelsv2InitialResponseSetting); err != nil {
			log.Errorf("invalid --initial-response-setting: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2InputContexts) > 0 {
		if err := assignInputField(input, "InputContexts", _lexmodelsv2InputContexts); err != nil {
			log.Errorf("invalid --input-contexts: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2IntentClosingSetting) > 0 {
		if err := assignInputField(input, "IntentClosingSetting", _lexmodelsv2IntentClosingSetting); err != nil {
			log.Errorf("invalid --intent-closing-setting: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2IntentConfirmationSetting) > 0 {
		if err := assignInputField(input, "IntentConfirmationSetting", _lexmodelsv2IntentConfirmationSetting); err != nil {
			log.Errorf("invalid --intent-confirmation-setting: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2IntentDisplayName) > 0 {
		input.IntentDisplayName = aws.String(_lexmodelsv2IntentDisplayName)
	}
	if len(_lexmodelsv2KendraConfiguration) > 0 {
		if err := assignInputField(input, "KendraConfiguration", _lexmodelsv2KendraConfiguration); err != nil {
			log.Errorf("invalid --kendra-configuration: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2OutputContexts) > 0 {
		if err := assignInputField(input, "OutputContexts", _lexmodelsv2OutputContexts); err != nil {
			log.Errorf("invalid --output-contexts: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2ParentIntentSignature) > 0 {
		input.ParentIntentSignature = aws.String(_lexmodelsv2ParentIntentSignature)
	}
	if len(_lexmodelsv2QInConnectIntentConfiguration) > 0 {
		if err := assignInputField(input, "QInConnectIntentConfiguration", _lexmodelsv2QInConnectIntentConfiguration); err != nil {
			log.Errorf("invalid --qin-connect-intent-configuration: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2QnAIntentConfiguration) > 0 {
		if err := assignInputField(input, "QnAIntentConfiguration", _lexmodelsv2QnAIntentConfiguration); err != nil {
			log.Errorf("invalid --qna-intent-configuration: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2SampleUtterances) > 0 {
		if err := assignInputField(input, "SampleUtterances", _lexmodelsv2SampleUtterances); err != nil {
			log.Errorf("invalid --sample-utterances: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2SlotPriorities) > 0 {
		if err := assignInputField(input, "SlotPriorities", _lexmodelsv2SlotPriorities); err != nil {
			log.Errorf("invalid --slot-priorities: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateIntent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Replaces the existing resource policy for a bot or bot alias with a new one. If
// the policy doesn't exist, Amazon Lex returns an exception.
func lexmodelsv2_UpdateResourcePolicy(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.UpdateResourcePolicyInput{
		// Policy: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_lexmodelsv2Policy) > 0 {
		input.Policy = aws.String(_lexmodelsv2Policy)
	}
	if len(_lexmodelsv2ResourceARN) > 0 {
		input.ResourceArn = aws.String(_lexmodelsv2ResourceARN)
	}
	if len(_lexmodelsv2ExpectedRevisionId) > 0 {
		input.ExpectedRevisionId = aws.String(_lexmodelsv2ExpectedRevisionId)
	}

	if resp, err := client.UpdateResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the settings for a slot.
func lexmodelsv2_UpdateSlot(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.UpdateSlotInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// IntentId: *string, // Required
		// LocaleId: *string, // Required
		// SlotId: *string, // Required
		// SlotName: *string, // Required
		// ValueElicitationSetting: *types.SlotValueElicitationSetting, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2IntentId) > 0 {
		input.IntentId = aws.String(_lexmodelsv2IntentId)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2SlotId) > 0 {
		input.SlotId = aws.String(_lexmodelsv2SlotId)
	}
	if len(_lexmodelsv2SlotName) > 0 {
		input.SlotName = aws.String(_lexmodelsv2SlotName)
	}
	if len(_lexmodelsv2ValueElicitationSetting) > 0 {
		if err := assignInputField(input, "ValueElicitationSetting", _lexmodelsv2ValueElicitationSetting); err != nil {
			log.Errorf("invalid --value-elicitation-setting: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Description) > 0 {
		input.Description = aws.String(_lexmodelsv2Description)
	}
	if len(_lexmodelsv2MultipleValuesSetting) > 0 {
		if err := assignInputField(input, "MultipleValuesSetting", _lexmodelsv2MultipleValuesSetting); err != nil {
			log.Errorf("invalid --multiple-values-setting: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2ObfuscationSetting) > 0 {
		if err := assignInputField(input, "ObfuscationSetting", _lexmodelsv2ObfuscationSetting); err != nil {
			log.Errorf("invalid --obfuscation-setting: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2SlotTypeId) > 0 {
		input.SlotTypeId = aws.String(_lexmodelsv2SlotTypeId)
	}
	if len(_lexmodelsv2SubSlotSetting) > 0 {
		if err := assignInputField(input, "SubSlotSetting", _lexmodelsv2SubSlotSetting); err != nil {
			log.Errorf("invalid --sub-slot-setting: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSlot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an existing slot type.
func lexmodelsv2_UpdateSlotType(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.UpdateSlotTypeInput{
		// BotId: *string, // Required
		// BotVersion: *string, // Required
		// LocaleId: *string, // Required
		// SlotTypeId: *string, // Required
		// SlotTypeName: *string, // Required
	}

	if len(_lexmodelsv2BotId) > 0 {
		input.BotId = aws.String(_lexmodelsv2BotId)
	}
	if len(_lexmodelsv2BotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelsv2BotVersion)
	}
	if len(_lexmodelsv2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexmodelsv2LocaleId)
	}
	if len(_lexmodelsv2SlotTypeId) > 0 {
		input.SlotTypeId = aws.String(_lexmodelsv2SlotTypeId)
	}
	if len(_lexmodelsv2SlotTypeName) > 0 {
		input.SlotTypeName = aws.String(_lexmodelsv2SlotTypeName)
	}
	if len(_lexmodelsv2CompositeSlotTypeSetting) > 0 {
		if err := assignInputField(input, "CompositeSlotTypeSetting", _lexmodelsv2CompositeSlotTypeSetting); err != nil {
			log.Errorf("invalid --composite-slot-type-setting: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2Description) > 0 {
		input.Description = aws.String(_lexmodelsv2Description)
	}
	if len(_lexmodelsv2ExternalSourceSetting) > 0 {
		if err := assignInputField(input, "ExternalSourceSetting", _lexmodelsv2ExternalSourceSetting); err != nil {
			log.Errorf("invalid --external-source-setting: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2ParentSlotTypeSignature) > 0 {
		input.ParentSlotTypeSignature = aws.String(_lexmodelsv2ParentSlotTypeSignature)
	}
	if len(_lexmodelsv2SlotTypeValues) > 0 {
		if err := assignInputField(input, "SlotTypeValues", _lexmodelsv2SlotTypeValues); err != nil {
			log.Errorf("invalid --slot-type-values: %s", err.Error())
			return
		}
	}
	if len(_lexmodelsv2ValueSelectionSetting) > 0 {
		if err := assignInputField(input, "ValueSelectionSetting", _lexmodelsv2ValueSelectionSetting); err != nil {
			log.Errorf("invalid --value-selection-setting: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSlotType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The action to update the test set.
func lexmodelsv2_UpdateTestSet(cfg aws.Config, client *lexmodelsv2.Client) {
	input := &lexmodelsv2.UpdateTestSetInput{
		// TestSetId: *string, // Required
		// TestSetName: *string, // Required
	}

	if len(_lexmodelsv2TestSetId) > 0 {
		input.TestSetId = aws.String(_lexmodelsv2TestSetId)
	}
	if len(_lexmodelsv2TestSetName) > 0 {
		input.TestSetName = aws.String(_lexmodelsv2TestSetName)
	}
	if len(_lexmodelsv2Description) > 0 {
		input.Description = aws.String(_lexmodelsv2Description)
	}

	if resp, err := client.UpdateTestSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_lexmodelsv2Cmd)
	_lexmodelsv2Cmd.Flags().SortFlags = false

	_lexmodelsv2Cmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_lexmodelsv2Cmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_lexmodelsv2Cmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_lexmodelsv2Cmd.Flags().StringSliceVarP(&_lexmodelsv2Action, "action", "", nil, "Action")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2AggregationDuration, "aggregation-duration", "", "", "Aggregation Duration")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2ApiMode, "api-mode", "", "", "API Mode")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2Attributes, "attributes", "", "", "Attributes")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2BinBy, "bin-by", "", "", "Bin By")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2BotAliasId, "bot-alias-id", "", "", "Bot Alias ID")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2BotAliasLocaleSettings, "bot-alias-locale-settings", "", "", "Bot Alias Locale Settings")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2BotAliasName, "bot-alias-name", "", "", "Bot Alias Name")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2BotId, "bot-id", "", "", "Bot ID")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2BotMembers, "bot-members", "", "", "Bot Members")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2BotName, "bot-name", "", "", "Bot Name")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2BotRecommendationId, "bot-recommendation-id", "", "", "Bot Recommendation ID")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2BotTags, "bot-tags", "", "", "Bot Tags")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2BotType, "bot-type", "", "", "Bot Type")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2BotVersion, "bot-version", "", "", "Bot Version")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2BotVersionLocaleSpecification, "bot-version-locale-specification", "", "", "Bot Version Locale Specification")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2CompositeSlotTypeSetting, "composite-slot-type-setting", "", "", "Composite Slot Type Setting")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2Condition, "condition", "", "", "Condition")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2ConversationLogSettings, "conversation-log-settings", "", "", "Conversation Log Settings")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2CustomVocabularyItemList, "custom-vocabulary-item-list", "", "", "Custom Vocabulary Item List")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2DataPrivacy, "data-privacy", "", "", "Data Privacy")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2Description, "description", "", "", "Description")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2DialogCodeHook, "dialog-code-hook", "", "", "Dialog Code Hook")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2Effect, "effect", "", "", "Effect")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2EncryptionSetting, "encryption-setting", "", "", "Encryption Setting")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2EndDateTime, "end-date-time", "", "", "End Date Time")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2ErrorLogSettings, "error-log-settings", "", "", "Error Log Settings")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2ExpectedRevisionId, "expected-revision-id", "", "", "Expected Revision ID")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2ExportId, "export-id", "", "", "Export ID")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2ExternalSourceSetting, "external-source-setting", "", "", "External Source Setting")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2FileFormat, "file-format", "", "", "File Format")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2FilePassword, "file-password", "", "", "File Password")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2Filters, "filters", "", "", "Filters")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2FulfillmentCodeHook, "fulfillment-code-hook", "", "", "Fulfillment Code Hook")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2GenerationDataSource, "generation-data-source", "", "", "Generation Data Source")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2GenerationId, "generation-id", "", "", "Generation ID")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2GenerationInputPrompt, "generation-input-prompt", "", "", "Generation Input Prompt")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2GenerativeAISettings, "generative-ai-settings", "", "", "Generative Ai Settings")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2GroupBy, "group-by", "", "", "Group By")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2IdleSessionTTLInSeconds, "idle-session-ttlin-seconds", "", "", "Idle Session Ttlin Seconds")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2ImportId, "import-id", "", "", "Import ID")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2InitialResponseSetting, "initial-response-setting", "", "", "Initial Response Setting")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2InputContexts, "input-contexts", "", "", "Input Contexts")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2IntentClosingSetting, "intent-closing-setting", "", "", "Intent Closing Setting")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2IntentConfirmationSetting, "intent-confirmation-setting", "", "", "Intent Confirmation Setting")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2IntentDisplayName, "intent-display-name", "", "", "Intent Display Name")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2IntentId, "intent-id", "", "", "Intent ID")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2IntentName, "intent-name", "", "", "Intent Name")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2IntentPath, "intent-path", "", "", "Intent Path")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2KendraConfiguration, "kendra-configuration", "", "", "Kendra Configuration")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2LocaleId, "locale-id", "", "", "Locale ID")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2MaxResults, "max-results", "", "", "Max Results")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2MergeStrategy, "merge-strategy", "", "", "Merge Strategy")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2Metrics, "metrics", "", "", "Metrics")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2MultipleValuesSetting, "multiple-values-setting", "", "", "Multiple Values Setting")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2NextIndex, "next-index", "", "", "Next Index")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2NextToken, "next-token", "", "", "Next Token")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2NluIntentConfidenceThreshold, "nlu-intent-confidence-threshold", "", "", "Nlu Intent Confidence Threshold")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2ObfuscationSetting, "obfuscation-setting", "", "", "Obfuscation Setting")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2OutputContexts, "output-contexts", "", "", "Output Contexts")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2ParentIntentSignature, "parent-intent-signature", "", "", "Parent Intent Signature")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2ParentSlotTypeSignature, "parent-slot-type-signature", "", "", "Parent Slot Type Signature")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2Policy, "policy", "", "", "Policy")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2Principal, "principal", "", "", "Principal")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2QInConnectIntentConfiguration, "qin-connect-intent-configuration", "", "", "Qin Connect Intent Configuration")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2QnAIntentConfiguration, "qna-intent-configuration", "", "", "Qna Intent Configuration")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2ReplicaRegion, "replica-region", "", "", "Replica Region")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2ResourceARN, "resource-arn", "", "", "Resource ARN")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2ResourceSpecification, "resource-specification", "", "", "Resource Specification")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2ResultFilterBy, "result-filter-by", "", "", "Result Filter By")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2RoleArn, "role-arn", "", "", "Role ARN")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2SampleUtterances, "sample-utterances", "", "", "Sample Utterances")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2SearchOrder, "search-order", "", "", "Search Order")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2SentimentAnalysisSettings, "sentiment-analysis-settings", "", "", "Sentiment Analysis Settings")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2SessionId, "session-id", "", "", "Session ID")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2SkipResourceInUseCheck, "skip-resource-in-use-check", "", "", "Skip Resource In Use Check")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2SlotId, "slot-id", "", "", "Slot ID")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2SlotName, "slot-name", "", "", "Slot Name")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2SlotPriorities, "slot-priorities", "", "", "Slot Priorities")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2SlotTypeId, "slot-type-id", "", "", "Slot Type ID")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2SlotTypeName, "slot-type-name", "", "", "Slot Type Name")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2SlotTypeValues, "slot-type-values", "", "", "Slot Type Values")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2SortBy, "sort-by", "", "", "Sort By")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2SpeechDetectionSensitivity, "speech-detection-sensitivity", "", "", "Speech Detection Sensitivity")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2SpeechRecognitionSettings, "speech-recognition-settings", "", "", "Speech Recognition Settings")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2StartDateTime, "start-date-time", "", "", "Start Date Time")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2StatementId, "statement-id", "", "", "Statement ID")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2StorageLocation, "storage-location", "", "", "Storage Location")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2SubSlotSetting, "sub-slot-setting", "", "", "Sub Slot Setting")
	_lexmodelsv2Cmd.Flags().StringSliceVarP(&_lexmodelsv2TagKeys, "tag-keys", "", nil, "Tag Keys")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2Tags, "tags", "", "", "Tags")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2Target, "target", "", "", "Target")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2TestBotAliasTags, "test-bot-alias-tags", "", "", "Test Bot Alias Tags")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2TestExecutionId, "test-execution-id", "", "", "Test Execution ID")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2TestExecutionModality, "test-execution-modality", "", "", "Test Execution Modality")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2TestSetDiscrepancyReportId, "test-set-discrepancy-report-id", "", "", "Test Set Discrepancy Report ID")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2TestSetGenerationId, "test-set-generation-id", "", "", "Test Set Generation ID")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2TestSetId, "test-set-id", "", "", "Test Set ID")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2TestSetName, "test-set-name", "", "", "Test Set Name")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2TestSetTags, "test-set-tags", "", "", "Test Set Tags")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2TranscriptSourceSetting, "transcript-source-setting", "", "", "Transcript Source Setting")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2UnifiedSpeechSettings, "unified-speech-settings", "", "", "Unified Speech Settings")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2ValueElicitationSetting, "value-elicitation-setting", "", "", "Value Elicitation Setting")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2ValueSelectionSetting, "value-selection-setting", "", "", "Value Selection Setting")
	_lexmodelsv2Cmd.Flags().StringVarP(&_lexmodelsv2VoiceSettings, "voice-settings", "", "", "Voice Settings")

	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2BatchCreateCustomVocabularyItem, "batch-create-custom-vocabulary-item", "", false, "Batch Create Custom Vocabulary Item")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2BatchDeleteCustomVocabularyItem, "batch-delete-custom-vocabulary-item", "", false, "Batch Delete Custom Vocabulary Item")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2BatchUpdateCustomVocabularyItem, "batch-update-custom-vocabulary-item", "", false, "Batch Update Custom Vocabulary Item")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2BuildBotLocale, "build-bot-locale", "", false, "Build Bot Locale")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2CreateBot, "create-bot", "", false, "Create Bot")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2CreateBotAlias, "create-bot-alias", "", false, "Create Bot Alias")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2CreateBotLocale, "create-bot-locale", "", false, "Create Bot Locale")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2CreateBotReplica, "create-bot-replica", "", false, "Create Bot Replica")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2CreateBotVersion, "create-bot-version", "", false, "Create Bot Version")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2CreateExport, "create-export", "", false, "Create Export")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2CreateIntent, "create-intent", "", false, "Create Intent")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2CreateResourcePolicy, "create-resource-policy", "", false, "Create Resource Policy")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2CreateResourcePolicyStatement, "create-resource-policy-statement", "", false, "Create Resource Policy Statement")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2CreateSlot, "create-slot", "", false, "Create Slot")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2CreateSlotType, "create-slot-type", "", false, "Create Slot Type")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2CreateTestSetDiscrepancyReport, "create-test-set-discrepancy-report", "", false, "Create Test Set Discrepancy Report")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2CreateUploadUrl, "create-upload-url", "", false, "Create Upload URL")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DeleteBot, "delete-bot", "", false, "Delete Bot")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DeleteBotAlias, "delete-bot-alias", "", false, "Delete Bot Alias")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DeleteBotLocale, "delete-bot-locale", "", false, "Delete Bot Locale")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DeleteBotReplica, "delete-bot-replica", "", false, "Delete Bot Replica")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DeleteBotVersion, "delete-bot-version", "", false, "Delete Bot Version")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DeleteCustomVocabulary, "delete-custom-vocabulary", "", false, "Delete Custom Vocabulary")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DeleteExport, "delete-export", "", false, "Delete Export")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DeleteImport, "delete-import", "", false, "Delete Import")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DeleteIntent, "delete-intent", "", false, "Delete Intent")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DeleteResourcePolicyStatement, "delete-resource-policy-statement", "", false, "Delete Resource Policy Statement")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DeleteSlot, "delete-slot", "", false, "Delete Slot")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DeleteSlotType, "delete-slot-type", "", false, "Delete Slot Type")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DeleteTestSet, "delete-test-set", "", false, "Delete Test Set")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DeleteUtterances, "delete-utterances", "", false, "Delete Utterances")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DescribeBot, "describe-bot", "", false, "Describe Bot")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DescribeBotAlias, "describe-bot-alias", "", false, "Describe Bot Alias")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DescribeBotLocale, "describe-bot-locale", "", false, "Describe Bot Locale")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DescribeBotRecommendation, "describe-bot-recommendation", "", false, "Describe Bot Recommendation")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DescribeBotReplica, "describe-bot-replica", "", false, "Describe Bot Replica")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DescribeBotResourceGeneration, "describe-bot-resource-generation", "", false, "Describe Bot Resource Generation")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DescribeBotVersion, "describe-bot-version", "", false, "Describe Bot Version")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DescribeCustomVocabularyMetadata, "describe-custom-vocabulary-metadata", "", false, "Describe Custom Vocabulary Metadata")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DescribeExport, "describe-export", "", false, "Describe Export")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DescribeImport, "describe-import", "", false, "Describe Import")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DescribeIntent, "describe-intent", "", false, "Describe Intent")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DescribeResourcePolicy, "describe-resource-policy", "", false, "Describe Resource Policy")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DescribeSlot, "describe-slot", "", false, "Describe Slot")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DescribeSlotType, "describe-slot-type", "", false, "Describe Slot Type")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DescribeTestExecution, "describe-test-execution", "", false, "Describe Test Execution")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DescribeTestSet, "describe-test-set", "", false, "Describe Test Set")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DescribeTestSetDiscrepancyReport, "describe-test-set-discrepancy-report", "", false, "Describe Test Set Discrepancy Report")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2DescribeTestSetGeneration, "describe-test-set-generation", "", false, "Describe Test Set Generation")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2GenerateBotElement, "generate-bot-element", "", false, "Generate Bot Element")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2GetTestExecutionArtifactsUrl, "get-test-execution-artifacts-url", "", false, "Get Test Execution Artifacts URL")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListAggregatedUtterances, "list-aggregated-utterances", "", false, "List Aggregated Utterances")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListBotAliasReplicas, "list-bot-alias-replicas", "", false, "List Bot Alias Replicas")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListBotAliases, "list-bot-aliases", "", false, "List Bot Aliases")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListBotLocales, "list-bot-locales", "", false, "List Bot Locales")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListBotRecommendations, "list-bot-recommendations", "", false, "List Bot Recommendations")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListBotReplicas, "list-bot-replicas", "", false, "List Bot Replicas")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListBotResourceGenerations, "list-bot-resource-generations", "", false, "List Bot Resource Generations")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListBotVersionReplicas, "list-bot-version-replicas", "", false, "List Bot Version Replicas")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListBotVersions, "list-bot-versions", "", false, "List Bot Versions")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListBots, "list-bots", "", false, "List Bots")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListBuiltInIntents, "list-built-in-intents", "", false, "List Built In Intents")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListBuiltInSlotTypes, "list-built-in-slot-types", "", false, "List Built In Slot Types")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListCustomVocabularyItems, "list-custom-vocabulary-items", "", false, "List Custom Vocabulary Items")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListExports, "list-exports", "", false, "List Exports")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListImports, "list-imports", "", false, "List Imports")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListIntentMetrics, "list-intent-metrics", "", false, "List Intent Metrics")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListIntentPaths, "list-intent-paths", "", false, "List Intent Paths")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListIntentStageMetrics, "list-intent-stage-metrics", "", false, "List Intent Stage Metrics")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListIntents, "list-intents", "", false, "List Intents")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListRecommendedIntents, "list-recommended-intents", "", false, "List Recommended Intents")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListSessionAnalyticsData, "list-session-analytics-data", "", false, "List Session Analytics Data")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListSessionMetrics, "list-session-metrics", "", false, "List Session Metrics")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListSlotTypes, "list-slot-types", "", false, "List Slot Types")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListSlots, "list-slots", "", false, "List Slots")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListTestExecutionResultItems, "list-test-execution-result-items", "", false, "List Test Execution Result Items")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListTestExecutions, "list-test-executions", "", false, "List Test Executions")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListTestSetRecords, "list-test-set-records", "", false, "List Test Set Records")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListTestSets, "list-test-sets", "", false, "List Test Sets")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListUtteranceAnalyticsData, "list-utterance-analytics-data", "", false, "List Utterance Analytics Data")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2ListUtteranceMetrics, "list-utterance-metrics", "", false, "List Utterance Metrics")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2SearchAssociatedTranscripts, "search-associated-transcripts", "", false, "Search Associated Transcripts")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2StartBotRecommendation, "start-bot-recommendation", "", false, "Start Bot Recommendation")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2StartBotResourceGeneration, "start-bot-resource-generation", "", false, "Start Bot Resource Generation")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2StartImport, "start-import", "", false, "Start Import")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2StartTestExecution, "start-test-execution", "", false, "Start Test Execution")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2StartTestSetGeneration, "start-test-set-generation", "", false, "Start Test Set Generation")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2StopBotRecommendation, "stop-bot-recommendation", "", false, "Stop Bot Recommendation")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2TagResource, "tag-resource", "", false, "Tag Resource")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2UntagResource, "untag-resource", "", false, "Untag Resource")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2UpdateBot, "update-bot", "", false, "Update Bot")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2UpdateBotAlias, "update-bot-alias", "", false, "Update Bot Alias")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2UpdateBotLocale, "update-bot-locale", "", false, "Update Bot Locale")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2UpdateBotRecommendation, "update-bot-recommendation", "", false, "Update Bot Recommendation")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2UpdateExport, "update-export", "", false, "Update Export")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2UpdateIntent, "update-intent", "", false, "Update Intent")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2UpdateResourcePolicy, "update-resource-policy", "", false, "Update Resource Policy")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2UpdateSlot, "update-slot", "", false, "Update Slot")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2UpdateSlotType, "update-slot-type", "", false, "Update Slot Type")
	_lexmodelsv2Cmd.Flags().BoolVarP(&_lexmodelsv2UpdateTestSet, "update-test-set", "", false, "Update Test Set")

}
