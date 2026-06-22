package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/lexmodelsv2"
)

var fields_batch_create_custom_vocabulary_item = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "CustomVocabularyItemList", Flag: "custom-vocabulary-item-list", Type: "[]types.NewCustomVocabularyItem", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
}

var fields_batch_delete_custom_vocabulary_item = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "CustomVocabularyItemList", Flag: "custom-vocabulary-item-list", Type: "[]types.CustomVocabularyEntryId", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
}

var fields_batch_update_custom_vocabulary_item = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "CustomVocabularyItemList", Flag: "custom-vocabulary-item-list", Type: "[]types.CustomVocabularyItem", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
}

var fields_build_bot_locale = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
}

var fields_create_bot = []leanruntime.Field{
	{Name: "BotMembers", Flag: "bot-members", Type: "[]types.BotMember", Required: false},
	{Name: "BotName", Flag: "bot-name", Type: "*string", Required: true},
	{Name: "BotTags", Flag: "bot-tags", Type: "map[string]string", Required: false},
	{Name: "BotType", Flag: "bot-type", Type: "types.BotType", Required: false},
	{Name: "DataPrivacy", Flag: "data-privacy", Type: "*types.DataPrivacy", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ErrorLogSettings", Flag: "error-log-settings", Type: "*types.ErrorLogSettings", Required: false},
	{Name: "IdleSessionTTLInSeconds", Flag: "idle-session-ttlin-seconds", Type: "*int32", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "TestBotAliasTags", Flag: "test-bot-alias-tags", Type: "map[string]string", Required: false},
}

var fields_create_bot_alias = []leanruntime.Field{
	{Name: "BotAliasLocaleSettings", Flag: "bot-alias-locale-settings", Type: "map[string]types.BotAliasLocaleSettings", Required: false},
	{Name: "BotAliasName", Flag: "bot-alias-name", Type: "*string", Required: true},
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: false},
	{Name: "ConversationLogSettings", Flag: "conversation-log-settings", Type: "*types.ConversationLogSettings", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "SentimentAnalysisSettings", Flag: "sentiment-analysis-settings", Type: "*types.SentimentAnalysisSettings", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_bot_locale = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GenerativeAISettings", Flag: "generative-ai-settings", Type: "*types.GenerativeAISettings", Required: false},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "NluIntentConfidenceThreshold", Flag: "nlu-intent-confidence-threshold", Type: "*float64", Required: true},
	{Name: "SpeechDetectionSensitivity", Flag: "speech-detection-sensitivity", Type: "types.SpeechDetectionSensitivity", Required: false},
	{Name: "SpeechRecognitionSettings", Flag: "speech-recognition-settings", Type: "*types.SpeechRecognitionSettings", Required: false},
	{Name: "UnifiedSpeechSettings", Flag: "unified-speech-settings", Type: "*types.UnifiedSpeechSettings", Required: false},
	{Name: "VoiceSettings", Flag: "voice-settings", Type: "*types.VoiceSettings", Required: false},
}

var fields_create_bot_replica = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "ReplicaRegion", Flag: "replica-region", Type: "*string", Required: true},
}

var fields_create_bot_version = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersionLocaleSpecification", Flag: "bot-version-locale-specification", Type: "map[string]types.BotVersionLocaleDetails", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
}

var fields_create_export = []leanruntime.Field{
	{Name: "FileFormat", Flag: "file-format", Type: "types.ImportExportFileFormat", Required: true},
	{Name: "FilePassword", Flag: "file-password", Type: "*string", Required: false},
	{Name: "ResourceSpecification", Flag: "resource-specification", Type: "*types.ExportResourceSpecification", Required: true},
}

var fields_create_intent = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DialogCodeHook", Flag: "dialog-code-hook", Type: "*types.DialogCodeHookSettings", Required: false},
	{Name: "FulfillmentCodeHook", Flag: "fulfillment-code-hook", Type: "*types.FulfillmentCodeHookSettings", Required: false},
	{Name: "InitialResponseSetting", Flag: "initial-response-setting", Type: "*types.InitialResponseSetting", Required: false},
	{Name: "InputContexts", Flag: "input-contexts", Type: "[]types.InputContext", Required: false},
	{Name: "IntentClosingSetting", Flag: "intent-closing-setting", Type: "*types.IntentClosingSetting", Required: false},
	{Name: "IntentConfirmationSetting", Flag: "intent-confirmation-setting", Type: "*types.IntentConfirmationSetting", Required: false},
	{Name: "IntentDisplayName", Flag: "intent-display-name", Type: "*string", Required: false},
	{Name: "IntentName", Flag: "intent-name", Type: "*string", Required: true},
	{Name: "KendraConfiguration", Flag: "kendra-configuration", Type: "*types.KendraConfiguration", Required: false},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "OutputContexts", Flag: "output-contexts", Type: "[]types.OutputContext", Required: false},
	{Name: "ParentIntentSignature", Flag: "parent-intent-signature", Type: "*string", Required: false},
	{Name: "QInConnectIntentConfiguration", Flag: "qin-connect-intent-configuration", Type: "*types.QInConnectIntentConfiguration", Required: false},
	{Name: "QnAIntentConfiguration", Flag: "qna-intent-configuration", Type: "*types.QnAIntentConfiguration", Required: false},
	{Name: "SampleUtterances", Flag: "sample-utterances", Type: "[]types.SampleUtterance", Required: false},
}

var fields_create_resource_policy = []leanruntime.Field{
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_create_resource_policy_statement = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "[]string", Required: true},
	{Name: "Condition", Flag: "condition", Type: "map[string]map[string]string", Required: false},
	{Name: "Effect", Flag: "effect", Type: "types.Effect", Required: true},
	{Name: "ExpectedRevisionId", Flag: "expected-revision-id", Type: "*string", Required: false},
	{Name: "Principal", Flag: "principal", Type: "[]types.Principal", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "StatementId", Flag: "statement-id", Type: "*string", Required: true},
}

var fields_create_slot = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IntentId", Flag: "intent-id", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "MultipleValuesSetting", Flag: "multiple-values-setting", Type: "*types.MultipleValuesSetting", Required: false},
	{Name: "ObfuscationSetting", Flag: "obfuscation-setting", Type: "*types.ObfuscationSetting", Required: false},
	{Name: "SlotName", Flag: "slot-name", Type: "*string", Required: true},
	{Name: "SlotTypeId", Flag: "slot-type-id", Type: "*string", Required: false},
	{Name: "SubSlotSetting", Flag: "sub-slot-setting", Type: "*types.SubSlotSetting", Required: false},
	{Name: "ValueElicitationSetting", Flag: "value-elicitation-setting", Type: "*types.SlotValueElicitationSetting", Required: true},
}

var fields_create_slot_type = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "CompositeSlotTypeSetting", Flag: "composite-slot-type-setting", Type: "*types.CompositeSlotTypeSetting", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExternalSourceSetting", Flag: "external-source-setting", Type: "*types.ExternalSourceSetting", Required: false},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "ParentSlotTypeSignature", Flag: "parent-slot-type-signature", Type: "*string", Required: false},
	{Name: "SlotTypeName", Flag: "slot-type-name", Type: "*string", Required: true},
	{Name: "SlotTypeValues", Flag: "slot-type-values", Type: "[]types.SlotTypeValue", Required: false},
	{Name: "ValueSelectionSetting", Flag: "value-selection-setting", Type: "*types.SlotValueSelectionSetting", Required: false},
}

var fields_create_test_set_discrepancy_report = []leanruntime.Field{
	{Name: "Target", Flag: "target", Type: "*types.TestSetDiscrepancyReportResourceTarget", Required: true},
	{Name: "TestSetId", Flag: "test-set-id", Type: "*string", Required: true},
}

var fields_create_upload_url = []leanruntime.Field{}

var fields_delete_bot = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "SkipResourceInUseCheck", Flag: "skip-resource-in-use-check", Type: "bool", Required: false},
}

var fields_delete_bot_alias = []leanruntime.Field{
	{Name: "BotAliasId", Flag: "bot-alias-id", Type: "*string", Required: true},
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "SkipResourceInUseCheck", Flag: "skip-resource-in-use-check", Type: "bool", Required: false},
}

var fields_delete_bot_locale = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
}

var fields_delete_bot_replica = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "ReplicaRegion", Flag: "replica-region", Type: "*string", Required: true},
}

var fields_delete_bot_version = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "SkipResourceInUseCheck", Flag: "skip-resource-in-use-check", Type: "bool", Required: false},
}

var fields_delete_custom_vocabulary = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
}

var fields_delete_export = []leanruntime.Field{
	{Name: "ExportId", Flag: "export-id", Type: "*string", Required: true},
}

var fields_delete_import = []leanruntime.Field{
	{Name: "ImportId", Flag: "import-id", Type: "*string", Required: true},
}

var fields_delete_intent = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "IntentId", Flag: "intent-id", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "ExpectedRevisionId", Flag: "expected-revision-id", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_delete_resource_policy_statement = []leanruntime.Field{
	{Name: "ExpectedRevisionId", Flag: "expected-revision-id", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "StatementId", Flag: "statement-id", Type: "*string", Required: true},
}

var fields_delete_slot = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "IntentId", Flag: "intent-id", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "SlotId", Flag: "slot-id", Type: "*string", Required: true},
}

var fields_delete_slot_type = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "SkipResourceInUseCheck", Flag: "skip-resource-in-use-check", Type: "bool", Required: false},
	{Name: "SlotTypeId", Flag: "slot-type-id", Type: "*string", Required: true},
}

var fields_delete_test_set = []leanruntime.Field{
	{Name: "TestSetId", Flag: "test-set-id", Type: "*string", Required: true},
}

var fields_delete_utterances = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: false},
}

var fields_describe_bot = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
}

var fields_describe_bot_alias = []leanruntime.Field{
	{Name: "BotAliasId", Flag: "bot-alias-id", Type: "*string", Required: true},
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
}

var fields_describe_bot_locale = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
}

var fields_describe_bot_recommendation = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotRecommendationId", Flag: "bot-recommendation-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
}

var fields_describe_bot_replica = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "ReplicaRegion", Flag: "replica-region", Type: "*string", Required: true},
}

var fields_describe_bot_resource_generation = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "GenerationId", Flag: "generation-id", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
}

var fields_describe_bot_version = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
}

var fields_describe_custom_vocabulary_metadata = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
}

var fields_describe_export = []leanruntime.Field{
	{Name: "ExportId", Flag: "export-id", Type: "*string", Required: true},
}

var fields_describe_import = []leanruntime.Field{
	{Name: "ImportId", Flag: "import-id", Type: "*string", Required: true},
}

var fields_describe_intent = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "IntentId", Flag: "intent-id", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
}

var fields_describe_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_describe_slot = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "IntentId", Flag: "intent-id", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "SlotId", Flag: "slot-id", Type: "*string", Required: true},
}

var fields_describe_slot_type = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "SlotTypeId", Flag: "slot-type-id", Type: "*string", Required: true},
}

var fields_describe_test_execution = []leanruntime.Field{
	{Name: "TestExecutionId", Flag: "test-execution-id", Type: "*string", Required: true},
}

var fields_describe_test_set = []leanruntime.Field{
	{Name: "TestSetId", Flag: "test-set-id", Type: "*string", Required: true},
}

var fields_describe_test_set_discrepancy_report = []leanruntime.Field{
	{Name: "TestSetDiscrepancyReportId", Flag: "test-set-discrepancy-report-id", Type: "*string", Required: true},
}

var fields_describe_test_set_generation = []leanruntime.Field{
	{Name: "TestSetGenerationId", Flag: "test-set-generation-id", Type: "*string", Required: true},
}

var fields_generate_bot_element = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "IntentId", Flag: "intent-id", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
}

var fields_get_test_execution_artifacts_url = []leanruntime.Field{
	{Name: "TestExecutionId", Flag: "test-execution-id", Type: "*string", Required: true},
}

var fields_list_aggregated_utterances = []leanruntime.Field{
	{Name: "AggregationDuration", Flag: "aggregation-duration", Type: "*types.UtteranceAggregationDuration", Required: true},
	{Name: "BotAliasId", Flag: "bot-alias-id", Type: "*string", Required: false},
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.AggregatedUtterancesFilter", Required: false},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*types.AggregatedUtterancesSortBy", Required: false},
}

var fields_list_bot_alias_replicas = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReplicaRegion", Flag: "replica-region", Type: "*string", Required: true},
}

var fields_list_bot_aliases = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_bot_locales = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.BotLocaleFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*types.BotLocaleSortBy", Required: false},
}

var fields_list_bot_recommendations = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_bot_replicas = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
}

var fields_list_bot_resource_generations = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*types.GenerationSortBy", Required: false},
}

var fields_list_bot_version_replicas = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReplicaRegion", Flag: "replica-region", Type: "*string", Required: true},
	{Name: "SortBy", Flag: "sort-by", Type: "*types.BotVersionReplicaSortBy", Required: false},
}

var fields_list_bot_versions = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*types.BotVersionSortBy", Required: false},
}

var fields_list_bots = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.BotFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*types.BotSortBy", Required: false},
}

var fields_list_built_in_intents = []leanruntime.Field{
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*types.BuiltInIntentSortBy", Required: false},
}

var fields_list_built_in_slot_types = []leanruntime.Field{
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*types.BuiltInSlotTypeSortBy", Required: false},
}

var fields_list_custom_vocabulary_items = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_exports = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: false},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.ExportFilter", Required: false},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*types.ExportSortBy", Required: false},
}

var fields_list_imports = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: false},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.ImportFilter", Required: false},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*types.ImportSortBy", Required: false},
}

var fields_list_intent_metrics = []leanruntime.Field{
	{Name: "BinBy", Flag: "bin-by", Type: "[]types.AnalyticsBinBySpecification", Required: false},
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "EndDateTime", Flag: "end-date-time", Type: "*time.Time", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.AnalyticsIntentFilter", Required: false},
	{Name: "GroupBy", Flag: "group-by", Type: "[]types.AnalyticsIntentGroupBySpecification", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Metrics", Flag: "metrics", Type: "[]types.AnalyticsIntentMetric", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartDateTime", Flag: "start-date-time", Type: "*time.Time", Required: true},
}

var fields_list_intent_paths = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "EndDateTime", Flag: "end-date-time", Type: "*time.Time", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.AnalyticsPathFilter", Required: false},
	{Name: "IntentPath", Flag: "intent-path", Type: "*string", Required: true},
	{Name: "StartDateTime", Flag: "start-date-time", Type: "*time.Time", Required: true},
}

var fields_list_intent_stage_metrics = []leanruntime.Field{
	{Name: "BinBy", Flag: "bin-by", Type: "[]types.AnalyticsBinBySpecification", Required: false},
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "EndDateTime", Flag: "end-date-time", Type: "*time.Time", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.AnalyticsIntentStageFilter", Required: false},
	{Name: "GroupBy", Flag: "group-by", Type: "[]types.AnalyticsIntentStageGroupBySpecification", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Metrics", Flag: "metrics", Type: "[]types.AnalyticsIntentStageMetric", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartDateTime", Flag: "start-date-time", Type: "*time.Time", Required: true},
}

var fields_list_intents = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.IntentFilter", Required: false},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*types.IntentSortBy", Required: false},
}

var fields_list_recommended_intents = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotRecommendationId", Flag: "bot-recommendation-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_session_analytics_data = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "EndDateTime", Flag: "end-date-time", Type: "*time.Time", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.AnalyticsSessionFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*types.SessionDataSortBy", Required: false},
	{Name: "StartDateTime", Flag: "start-date-time", Type: "*time.Time", Required: true},
}

var fields_list_session_metrics = []leanruntime.Field{
	{Name: "BinBy", Flag: "bin-by", Type: "[]types.AnalyticsBinBySpecification", Required: false},
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "EndDateTime", Flag: "end-date-time", Type: "*time.Time", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.AnalyticsSessionFilter", Required: false},
	{Name: "GroupBy", Flag: "group-by", Type: "[]types.AnalyticsSessionGroupBySpecification", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Metrics", Flag: "metrics", Type: "[]types.AnalyticsSessionMetric", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartDateTime", Flag: "start-date-time", Type: "*time.Time", Required: true},
}

var fields_list_slot_types = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.SlotTypeFilter", Required: false},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*types.SlotTypeSortBy", Required: false},
}

var fields_list_slots = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.SlotFilter", Required: false},
	{Name: "IntentId", Flag: "intent-id", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*types.SlotSortBy", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_test_execution_result_items = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResultFilterBy", Flag: "result-filter-by", Type: "*types.TestExecutionResultFilterBy", Required: true},
	{Name: "TestExecutionId", Flag: "test-execution-id", Type: "*string", Required: true},
}

var fields_list_test_executions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*types.TestExecutionSortBy", Required: false},
}

var fields_list_test_set_records = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TestSetId", Flag: "test-set-id", Type: "*string", Required: true},
}

var fields_list_test_sets = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*types.TestSetSortBy", Required: false},
}

var fields_list_utterance_analytics_data = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "EndDateTime", Flag: "end-date-time", Type: "*time.Time", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.AnalyticsUtteranceFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*types.UtteranceDataSortBy", Required: false},
	{Name: "StartDateTime", Flag: "start-date-time", Type: "*time.Time", Required: true},
}

var fields_list_utterance_metrics = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "[]types.AnalyticsUtteranceAttribute", Required: false},
	{Name: "BinBy", Flag: "bin-by", Type: "[]types.AnalyticsBinBySpecification", Required: false},
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "EndDateTime", Flag: "end-date-time", Type: "*time.Time", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.AnalyticsUtteranceFilter", Required: false},
	{Name: "GroupBy", Flag: "group-by", Type: "[]types.AnalyticsUtteranceGroupBySpecification", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Metrics", Flag: "metrics", Type: "[]types.AnalyticsUtteranceMetric", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartDateTime", Flag: "start-date-time", Type: "*time.Time", Required: true},
}

var fields_search_associated_transcripts = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotRecommendationId", Flag: "bot-recommendation-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.AssociatedTranscriptFilter", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextIndex", Flag: "next-index", Type: "*int32", Required: false},
	{Name: "SearchOrder", Flag: "search-order", Type: "types.SearchOrder", Required: false},
}

var fields_start_bot_recommendation = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "EncryptionSetting", Flag: "encryption-setting", Type: "*types.EncryptionSetting", Required: false},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "TranscriptSourceSetting", Flag: "transcript-source-setting", Type: "*types.TranscriptSourceSetting", Required: true},
}

var fields_start_bot_resource_generation = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "GenerationInputPrompt", Flag: "generation-input-prompt", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
}

var fields_start_import = []leanruntime.Field{
	{Name: "FilePassword", Flag: "file-password", Type: "*string", Required: false},
	{Name: "ImportId", Flag: "import-id", Type: "*string", Required: true},
	{Name: "MergeStrategy", Flag: "merge-strategy", Type: "types.MergeStrategy", Required: true},
	{Name: "ResourceSpecification", Flag: "resource-specification", Type: "*types.ImportResourceSpecification", Required: true},
}

var fields_start_test_execution = []leanruntime.Field{
	{Name: "ApiMode", Flag: "api-mode", Type: "types.TestExecutionApiMode", Required: true},
	{Name: "Target", Flag: "target", Type: "*types.TestExecutionTarget", Required: true},
	{Name: "TestExecutionModality", Flag: "test-execution-modality", Type: "types.TestExecutionModality", Required: false},
	{Name: "TestSetId", Flag: "test-set-id", Type: "*string", Required: true},
}

var fields_start_test_set_generation = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GenerationDataSource", Flag: "generation-data-source", Type: "*types.TestSetGenerationDataSource", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "StorageLocation", Flag: "storage-location", Type: "*types.TestSetStorageLocation", Required: true},
	{Name: "TestSetName", Flag: "test-set-name", Type: "*string", Required: true},
	{Name: "TestSetTags", Flag: "test-set-tags", Type: "map[string]string", Required: false},
}

var fields_stop_bot_recommendation = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotRecommendationId", Flag: "bot-recommendation-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_bot = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotMembers", Flag: "bot-members", Type: "[]types.BotMember", Required: false},
	{Name: "BotName", Flag: "bot-name", Type: "*string", Required: true},
	{Name: "BotType", Flag: "bot-type", Type: "types.BotType", Required: false},
	{Name: "DataPrivacy", Flag: "data-privacy", Type: "*types.DataPrivacy", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ErrorLogSettings", Flag: "error-log-settings", Type: "*types.ErrorLogSettings", Required: false},
	{Name: "IdleSessionTTLInSeconds", Flag: "idle-session-ttlin-seconds", Type: "*int32", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
}

var fields_update_bot_alias = []leanruntime.Field{
	{Name: "BotAliasId", Flag: "bot-alias-id", Type: "*string", Required: true},
	{Name: "BotAliasLocaleSettings", Flag: "bot-alias-locale-settings", Type: "map[string]types.BotAliasLocaleSettings", Required: false},
	{Name: "BotAliasName", Flag: "bot-alias-name", Type: "*string", Required: true},
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: false},
	{Name: "ConversationLogSettings", Flag: "conversation-log-settings", Type: "*types.ConversationLogSettings", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "SentimentAnalysisSettings", Flag: "sentiment-analysis-settings", Type: "*types.SentimentAnalysisSettings", Required: false},
}

var fields_update_bot_locale = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GenerativeAISettings", Flag: "generative-ai-settings", Type: "*types.GenerativeAISettings", Required: false},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "NluIntentConfidenceThreshold", Flag: "nlu-intent-confidence-threshold", Type: "*float64", Required: true},
	{Name: "SpeechDetectionSensitivity", Flag: "speech-detection-sensitivity", Type: "types.SpeechDetectionSensitivity", Required: false},
	{Name: "SpeechRecognitionSettings", Flag: "speech-recognition-settings", Type: "*types.SpeechRecognitionSettings", Required: false},
	{Name: "UnifiedSpeechSettings", Flag: "unified-speech-settings", Type: "*types.UnifiedSpeechSettings", Required: false},
	{Name: "VoiceSettings", Flag: "voice-settings", Type: "*types.VoiceSettings", Required: false},
}

var fields_update_bot_recommendation = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotRecommendationId", Flag: "bot-recommendation-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "EncryptionSetting", Flag: "encryption-setting", Type: "*types.EncryptionSetting", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
}

var fields_update_export = []leanruntime.Field{
	{Name: "ExportId", Flag: "export-id", Type: "*string", Required: true},
	{Name: "FilePassword", Flag: "file-password", Type: "*string", Required: false},
}

var fields_update_intent = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DialogCodeHook", Flag: "dialog-code-hook", Type: "*types.DialogCodeHookSettings", Required: false},
	{Name: "FulfillmentCodeHook", Flag: "fulfillment-code-hook", Type: "*types.FulfillmentCodeHookSettings", Required: false},
	{Name: "InitialResponseSetting", Flag: "initial-response-setting", Type: "*types.InitialResponseSetting", Required: false},
	{Name: "InputContexts", Flag: "input-contexts", Type: "[]types.InputContext", Required: false},
	{Name: "IntentClosingSetting", Flag: "intent-closing-setting", Type: "*types.IntentClosingSetting", Required: false},
	{Name: "IntentConfirmationSetting", Flag: "intent-confirmation-setting", Type: "*types.IntentConfirmationSetting", Required: false},
	{Name: "IntentDisplayName", Flag: "intent-display-name", Type: "*string", Required: false},
	{Name: "IntentId", Flag: "intent-id", Type: "*string", Required: true},
	{Name: "IntentName", Flag: "intent-name", Type: "*string", Required: true},
	{Name: "KendraConfiguration", Flag: "kendra-configuration", Type: "*types.KendraConfiguration", Required: false},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "OutputContexts", Flag: "output-contexts", Type: "[]types.OutputContext", Required: false},
	{Name: "ParentIntentSignature", Flag: "parent-intent-signature", Type: "*string", Required: false},
	{Name: "QInConnectIntentConfiguration", Flag: "qin-connect-intent-configuration", Type: "*types.QInConnectIntentConfiguration", Required: false},
	{Name: "QnAIntentConfiguration", Flag: "qna-intent-configuration", Type: "*types.QnAIntentConfiguration", Required: false},
	{Name: "SampleUtterances", Flag: "sample-utterances", Type: "[]types.SampleUtterance", Required: false},
	{Name: "SlotPriorities", Flag: "slot-priorities", Type: "[]types.SlotPriority", Required: false},
}

var fields_update_resource_policy = []leanruntime.Field{
	{Name: "ExpectedRevisionId", Flag: "expected-revision-id", Type: "*string", Required: false},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_update_slot = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IntentId", Flag: "intent-id", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "MultipleValuesSetting", Flag: "multiple-values-setting", Type: "*types.MultipleValuesSetting", Required: false},
	{Name: "ObfuscationSetting", Flag: "obfuscation-setting", Type: "*types.ObfuscationSetting", Required: false},
	{Name: "SlotId", Flag: "slot-id", Type: "*string", Required: true},
	{Name: "SlotName", Flag: "slot-name", Type: "*string", Required: true},
	{Name: "SlotTypeId", Flag: "slot-type-id", Type: "*string", Required: false},
	{Name: "SubSlotSetting", Flag: "sub-slot-setting", Type: "*types.SubSlotSetting", Required: false},
	{Name: "ValueElicitationSetting", Flag: "value-elicitation-setting", Type: "*types.SlotValueElicitationSetting", Required: true},
}

var fields_update_slot_type = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "CompositeSlotTypeSetting", Flag: "composite-slot-type-setting", Type: "*types.CompositeSlotTypeSetting", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExternalSourceSetting", Flag: "external-source-setting", Type: "*types.ExternalSourceSetting", Required: false},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "ParentSlotTypeSignature", Flag: "parent-slot-type-signature", Type: "*string", Required: false},
	{Name: "SlotTypeId", Flag: "slot-type-id", Type: "*string", Required: true},
	{Name: "SlotTypeName", Flag: "slot-type-name", Type: "*string", Required: true},
	{Name: "SlotTypeValues", Flag: "slot-type-values", Type: "[]types.SlotTypeValue", Required: false},
	{Name: "ValueSelectionSetting", Flag: "value-selection-setting", Type: "*types.SlotValueSelectionSetting", Required: false},
}

var fields_update_test_set = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "TestSetId", Flag: "test-set-id", Type: "*string", Required: true},
	{Name: "TestSetName", Flag: "test-set-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-create-custom-vocabulary-item": {
			Name:   "batch-create-custom-vocabulary-item",
			Fields: fields_batch_create_custom_vocabulary_item,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchCreateCustomVocabularyItemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_create_custom_vocabulary_item, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchCreateCustomVocabularyItem(ctx, input)
			},
		},
		"batch-delete-custom-vocabulary-item": {
			Name:   "batch-delete-custom-vocabulary-item",
			Fields: fields_batch_delete_custom_vocabulary_item,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteCustomVocabularyItemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_custom_vocabulary_item, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteCustomVocabularyItem(ctx, input)
			},
		},
		"batch-update-custom-vocabulary-item": {
			Name:   "batch-update-custom-vocabulary-item",
			Fields: fields_batch_update_custom_vocabulary_item,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdateCustomVocabularyItemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_custom_vocabulary_item, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdateCustomVocabularyItem(ctx, input)
			},
		},
		"build-bot-locale": {
			Name:   "build-bot-locale",
			Fields: fields_build_bot_locale,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BuildBotLocaleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_build_bot_locale, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BuildBotLocale(ctx, input)
			},
		},
		"create-bot": {
			Name:   "create-bot",
			Fields: fields_create_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBot(ctx, input)
			},
		},
		"create-bot-alias": {
			Name:   "create-bot-alias",
			Fields: fields_create_bot_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBotAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_bot_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBotAlias(ctx, input)
			},
		},
		"create-bot-locale": {
			Name:   "create-bot-locale",
			Fields: fields_create_bot_locale,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBotLocaleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_bot_locale, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBotLocale(ctx, input)
			},
		},
		"create-bot-replica": {
			Name:   "create-bot-replica",
			Fields: fields_create_bot_replica,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBotReplicaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_bot_replica, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBotReplica(ctx, input)
			},
		},
		"create-bot-version": {
			Name:   "create-bot-version",
			Fields: fields_create_bot_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBotVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_bot_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBotVersion(ctx, input)
			},
		},
		"create-export": {
			Name:   "create-export",
			Fields: fields_create_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateExport(ctx, input)
			},
		},
		"create-intent": {
			Name:   "create-intent",
			Fields: fields_create_intent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIntentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_intent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIntent(ctx, input)
			},
		},
		"create-resource-policy": {
			Name:   "create-resource-policy",
			Fields: fields_create_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResourcePolicy(ctx, input)
			},
		},
		"create-resource-policy-statement": {
			Name:   "create-resource-policy-statement",
			Fields: fields_create_resource_policy_statement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResourcePolicyStatementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_resource_policy_statement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResourcePolicyStatement(ctx, input)
			},
		},
		"create-slot": {
			Name:   "create-slot",
			Fields: fields_create_slot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSlotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_slot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSlot(ctx, input)
			},
		},
		"create-slot-type": {
			Name:   "create-slot-type",
			Fields: fields_create_slot_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSlotTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_slot_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSlotType(ctx, input)
			},
		},
		"create-test-set-discrepancy-report": {
			Name:   "create-test-set-discrepancy-report",
			Fields: fields_create_test_set_discrepancy_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTestSetDiscrepancyReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_test_set_discrepancy_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTestSetDiscrepancyReport(ctx, input)
			},
		},
		"create-upload-url": {
			Name:   "create-upload-url",
			Fields: fields_create_upload_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUploadUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_upload_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUploadUrl(ctx, input)
			},
		},
		"delete-bot": {
			Name:   "delete-bot",
			Fields: fields_delete_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBot(ctx, input)
			},
		},
		"delete-bot-alias": {
			Name:   "delete-bot-alias",
			Fields: fields_delete_bot_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBotAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bot_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBotAlias(ctx, input)
			},
		},
		"delete-bot-locale": {
			Name:   "delete-bot-locale",
			Fields: fields_delete_bot_locale,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBotLocaleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bot_locale, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBotLocale(ctx, input)
			},
		},
		"delete-bot-replica": {
			Name:   "delete-bot-replica",
			Fields: fields_delete_bot_replica,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBotReplicaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bot_replica, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBotReplica(ctx, input)
			},
		},
		"delete-bot-version": {
			Name:   "delete-bot-version",
			Fields: fields_delete_bot_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBotVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bot_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBotVersion(ctx, input)
			},
		},
		"delete-custom-vocabulary": {
			Name:   "delete-custom-vocabulary",
			Fields: fields_delete_custom_vocabulary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCustomVocabularyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_custom_vocabulary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCustomVocabulary(ctx, input)
			},
		},
		"delete-export": {
			Name:   "delete-export",
			Fields: fields_delete_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteExport(ctx, input)
			},
		},
		"delete-import": {
			Name:   "delete-import",
			Fields: fields_delete_import,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteImportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_import, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteImport(ctx, input)
			},
		},
		"delete-intent": {
			Name:   "delete-intent",
			Fields: fields_delete_intent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIntentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_intent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIntent(ctx, input)
			},
		},
		"delete-resource-policy": {
			Name:   "delete-resource-policy",
			Fields: fields_delete_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourcePolicy(ctx, input)
			},
		},
		"delete-resource-policy-statement": {
			Name:   "delete-resource-policy-statement",
			Fields: fields_delete_resource_policy_statement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourcePolicyStatementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_policy_statement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourcePolicyStatement(ctx, input)
			},
		},
		"delete-slot": {
			Name:   "delete-slot",
			Fields: fields_delete_slot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSlotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_slot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSlot(ctx, input)
			},
		},
		"delete-slot-type": {
			Name:   "delete-slot-type",
			Fields: fields_delete_slot_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSlotTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_slot_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSlotType(ctx, input)
			},
		},
		"delete-test-set": {
			Name:   "delete-test-set",
			Fields: fields_delete_test_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTestSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_test_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTestSet(ctx, input)
			},
		},
		"delete-utterances": {
			Name:   "delete-utterances",
			Fields: fields_delete_utterances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUtterancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_utterances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUtterances(ctx, input)
			},
		},
		"describe-bot": {
			Name:   "describe-bot",
			Fields: fields_describe_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBot(ctx, input)
			},
		},
		"describe-bot-alias": {
			Name:   "describe-bot-alias",
			Fields: fields_describe_bot_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBotAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_bot_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBotAlias(ctx, input)
			},
		},
		"describe-bot-locale": {
			Name:   "describe-bot-locale",
			Fields: fields_describe_bot_locale,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBotLocaleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_bot_locale, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBotLocale(ctx, input)
			},
		},
		"describe-bot-recommendation": {
			Name:   "describe-bot-recommendation",
			Fields: fields_describe_bot_recommendation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBotRecommendationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_bot_recommendation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBotRecommendation(ctx, input)
			},
		},
		"describe-bot-replica": {
			Name:   "describe-bot-replica",
			Fields: fields_describe_bot_replica,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBotReplicaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_bot_replica, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBotReplica(ctx, input)
			},
		},
		"describe-bot-resource-generation": {
			Name:   "describe-bot-resource-generation",
			Fields: fields_describe_bot_resource_generation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBotResourceGenerationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_bot_resource_generation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBotResourceGeneration(ctx, input)
			},
		},
		"describe-bot-version": {
			Name:   "describe-bot-version",
			Fields: fields_describe_bot_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBotVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_bot_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBotVersion(ctx, input)
			},
		},
		"describe-custom-vocabulary-metadata": {
			Name:   "describe-custom-vocabulary-metadata",
			Fields: fields_describe_custom_vocabulary_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCustomVocabularyMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_custom_vocabulary_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCustomVocabularyMetadata(ctx, input)
			},
		},
		"describe-export": {
			Name:   "describe-export",
			Fields: fields_describe_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeExport(ctx, input)
			},
		},
		"describe-import": {
			Name:   "describe-import",
			Fields: fields_describe_import,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeImportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_import, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeImport(ctx, input)
			},
		},
		"describe-intent": {
			Name:   "describe-intent",
			Fields: fields_describe_intent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIntentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_intent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeIntent(ctx, input)
			},
		},
		"describe-resource-policy": {
			Name:   "describe-resource-policy",
			Fields: fields_describe_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeResourcePolicy(ctx, input)
			},
		},
		"describe-slot": {
			Name:   "describe-slot",
			Fields: fields_describe_slot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSlotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_slot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSlot(ctx, input)
			},
		},
		"describe-slot-type": {
			Name:   "describe-slot-type",
			Fields: fields_describe_slot_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSlotTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_slot_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSlotType(ctx, input)
			},
		},
		"describe-test-execution": {
			Name:   "describe-test-execution",
			Fields: fields_describe_test_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTestExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_test_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTestExecution(ctx, input)
			},
		},
		"describe-test-set": {
			Name:   "describe-test-set",
			Fields: fields_describe_test_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTestSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_test_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTestSet(ctx, input)
			},
		},
		"describe-test-set-discrepancy-report": {
			Name:   "describe-test-set-discrepancy-report",
			Fields: fields_describe_test_set_discrepancy_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTestSetDiscrepancyReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_test_set_discrepancy_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTestSetDiscrepancyReport(ctx, input)
			},
		},
		"describe-test-set-generation": {
			Name:   "describe-test-set-generation",
			Fields: fields_describe_test_set_generation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTestSetGenerationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_test_set_generation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTestSetGeneration(ctx, input)
			},
		},
		"generate-bot-element": {
			Name:   "generate-bot-element",
			Fields: fields_generate_bot_element,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateBotElementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_bot_element, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateBotElement(ctx, input)
			},
		},
		"get-test-execution-artifacts-url": {
			Name:   "get-test-execution-artifacts-url",
			Fields: fields_get_test_execution_artifacts_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTestExecutionArtifactsUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_test_execution_artifacts_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTestExecutionArtifactsUrl(ctx, input)
			},
		},
		"list-aggregated-utterances": {
			Name:   "list-aggregated-utterances",
			Fields: fields_list_aggregated_utterances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAggregatedUtterancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_aggregated_utterances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAggregatedUtterances(ctx, input)
				}
				var results []*svc.ListAggregatedUtterancesOutput
				p := svc.NewListAggregatedUtterancesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-bot-alias-replicas": {
			Name:   "list-bot-alias-replicas",
			Fields: fields_list_bot_alias_replicas,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBotAliasReplicasInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_bot_alias_replicas, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBotAliasReplicas(ctx, input)
				}
				var results []*svc.ListBotAliasReplicasOutput
				p := svc.NewListBotAliasReplicasPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-bot-aliases": {
			Name:   "list-bot-aliases",
			Fields: fields_list_bot_aliases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBotAliasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_bot_aliases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBotAliases(ctx, input)
				}
				var results []*svc.ListBotAliasesOutput
				p := svc.NewListBotAliasesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-bot-locales": {
			Name:   "list-bot-locales",
			Fields: fields_list_bot_locales,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBotLocalesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_bot_locales, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBotLocales(ctx, input)
				}
				var results []*svc.ListBotLocalesOutput
				p := svc.NewListBotLocalesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-bot-recommendations": {
			Name:   "list-bot-recommendations",
			Fields: fields_list_bot_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBotRecommendationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_bot_recommendations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBotRecommendations(ctx, input)
				}
				var results []*svc.ListBotRecommendationsOutput
				p := svc.NewListBotRecommendationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-bot-replicas": {
			Name:   "list-bot-replicas",
			Fields: fields_list_bot_replicas,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBotReplicasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_bot_replicas, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListBotReplicas(ctx, input)
			},
		},
		"list-bot-resource-generations": {
			Name:   "list-bot-resource-generations",
			Fields: fields_list_bot_resource_generations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBotResourceGenerationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_bot_resource_generations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBotResourceGenerations(ctx, input)
				}
				var results []*svc.ListBotResourceGenerationsOutput
				p := svc.NewListBotResourceGenerationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-bot-version-replicas": {
			Name:   "list-bot-version-replicas",
			Fields: fields_list_bot_version_replicas,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBotVersionReplicasInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_bot_version_replicas, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBotVersionReplicas(ctx, input)
				}
				var results []*svc.ListBotVersionReplicasOutput
				p := svc.NewListBotVersionReplicasPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-bot-versions": {
			Name:   "list-bot-versions",
			Fields: fields_list_bot_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBotVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_bot_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBotVersions(ctx, input)
				}
				var results []*svc.ListBotVersionsOutput
				p := svc.NewListBotVersionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-bots": {
			Name:   "list-bots",
			Fields: fields_list_bots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBotsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_bots, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBots(ctx, input)
				}
				var results []*svc.ListBotsOutput
				p := svc.NewListBotsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-built-in-intents": {
			Name:   "list-built-in-intents",
			Fields: fields_list_built_in_intents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBuiltInIntentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_built_in_intents, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBuiltInIntents(ctx, input)
				}
				var results []*svc.ListBuiltInIntentsOutput
				p := svc.NewListBuiltInIntentsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-built-in-slot-types": {
			Name:   "list-built-in-slot-types",
			Fields: fields_list_built_in_slot_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBuiltInSlotTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_built_in_slot_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBuiltInSlotTypes(ctx, input)
				}
				var results []*svc.ListBuiltInSlotTypesOutput
				p := svc.NewListBuiltInSlotTypesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-custom-vocabulary-items": {
			Name:   "list-custom-vocabulary-items",
			Fields: fields_list_custom_vocabulary_items,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCustomVocabularyItemsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_custom_vocabulary_items, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCustomVocabularyItems(ctx, input)
				}
				var results []*svc.ListCustomVocabularyItemsOutput
				p := svc.NewListCustomVocabularyItemsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-exports": {
			Name:   "list-exports",
			Fields: fields_list_exports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_exports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExports(ctx, input)
				}
				var results []*svc.ListExportsOutput
				p := svc.NewListExportsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-imports": {
			Name:   "list-imports",
			Fields: fields_list_imports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_imports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListImports(ctx, input)
				}
				var results []*svc.ListImportsOutput
				p := svc.NewListImportsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-intent-metrics": {
			Name:   "list-intent-metrics",
			Fields: fields_list_intent_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIntentMetricsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_intent_metrics, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIntentMetrics(ctx, input)
				}
				var results []*svc.ListIntentMetricsOutput
				p := svc.NewListIntentMetricsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-intent-paths": {
			Name:   "list-intent-paths",
			Fields: fields_list_intent_paths,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIntentPathsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_intent_paths, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListIntentPaths(ctx, input)
			},
		},
		"list-intent-stage-metrics": {
			Name:   "list-intent-stage-metrics",
			Fields: fields_list_intent_stage_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIntentStageMetricsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_intent_stage_metrics, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIntentStageMetrics(ctx, input)
				}
				var results []*svc.ListIntentStageMetricsOutput
				p := svc.NewListIntentStageMetricsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-intents": {
			Name:   "list-intents",
			Fields: fields_list_intents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIntentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_intents, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIntents(ctx, input)
				}
				var results []*svc.ListIntentsOutput
				p := svc.NewListIntentsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-recommended-intents": {
			Name:   "list-recommended-intents",
			Fields: fields_list_recommended_intents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecommendedIntentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_recommended_intents, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRecommendedIntents(ctx, input)
				}
				var results []*svc.ListRecommendedIntentsOutput
				p := svc.NewListRecommendedIntentsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-session-analytics-data": {
			Name:   "list-session-analytics-data",
			Fields: fields_list_session_analytics_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSessionAnalyticsDataInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_session_analytics_data, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSessionAnalyticsData(ctx, input)
				}
				var results []*svc.ListSessionAnalyticsDataOutput
				p := svc.NewListSessionAnalyticsDataPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-session-metrics": {
			Name:   "list-session-metrics",
			Fields: fields_list_session_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSessionMetricsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_session_metrics, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSessionMetrics(ctx, input)
				}
				var results []*svc.ListSessionMetricsOutput
				p := svc.NewListSessionMetricsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-slot-types": {
			Name:   "list-slot-types",
			Fields: fields_list_slot_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSlotTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_slot_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSlotTypes(ctx, input)
				}
				var results []*svc.ListSlotTypesOutput
				p := svc.NewListSlotTypesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-slots": {
			Name:   "list-slots",
			Fields: fields_list_slots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSlotsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_slots, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSlots(ctx, input)
				}
				var results []*svc.ListSlotsOutput
				p := svc.NewListSlotsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-tags-for-resource": {
			Name:   "list-tags-for-resource",
			Fields: fields_list_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForResource(ctx, input)
			},
		},
		"list-test-execution-result-items": {
			Name:   "list-test-execution-result-items",
			Fields: fields_list_test_execution_result_items,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTestExecutionResultItemsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_test_execution_result_items, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTestExecutionResultItems(ctx, input)
				}
				var results []*svc.ListTestExecutionResultItemsOutput
				p := svc.NewListTestExecutionResultItemsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-test-executions": {
			Name:   "list-test-executions",
			Fields: fields_list_test_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTestExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_test_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTestExecutions(ctx, input)
				}
				var results []*svc.ListTestExecutionsOutput
				p := svc.NewListTestExecutionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-test-set-records": {
			Name:   "list-test-set-records",
			Fields: fields_list_test_set_records,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTestSetRecordsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_test_set_records, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTestSetRecords(ctx, input)
				}
				var results []*svc.ListTestSetRecordsOutput
				p := svc.NewListTestSetRecordsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-test-sets": {
			Name:   "list-test-sets",
			Fields: fields_list_test_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTestSetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_test_sets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTestSets(ctx, input)
				}
				var results []*svc.ListTestSetsOutput
				p := svc.NewListTestSetsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-utterance-analytics-data": {
			Name:   "list-utterance-analytics-data",
			Fields: fields_list_utterance_analytics_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUtteranceAnalyticsDataInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_utterance_analytics_data, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUtteranceAnalyticsData(ctx, input)
				}
				var results []*svc.ListUtteranceAnalyticsDataOutput
				p := svc.NewListUtteranceAnalyticsDataPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-utterance-metrics": {
			Name:   "list-utterance-metrics",
			Fields: fields_list_utterance_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUtteranceMetricsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_utterance_metrics, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUtteranceMetrics(ctx, input)
				}
				var results []*svc.ListUtteranceMetricsOutput
				p := svc.NewListUtteranceMetricsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"search-associated-transcripts": {
			Name:   "search-associated-transcripts",
			Fields: fields_search_associated_transcripts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchAssociatedTranscriptsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_search_associated_transcripts, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SearchAssociatedTranscripts(ctx, input)
			},
		},
		"start-bot-recommendation": {
			Name:   "start-bot-recommendation",
			Fields: fields_start_bot_recommendation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartBotRecommendationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_bot_recommendation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartBotRecommendation(ctx, input)
			},
		},
		"start-bot-resource-generation": {
			Name:   "start-bot-resource-generation",
			Fields: fields_start_bot_resource_generation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartBotResourceGenerationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_bot_resource_generation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartBotResourceGeneration(ctx, input)
			},
		},
		"start-import": {
			Name:   "start-import",
			Fields: fields_start_import,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartImportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_import, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartImport(ctx, input)
			},
		},
		"start-test-execution": {
			Name:   "start-test-execution",
			Fields: fields_start_test_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartTestExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_test_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartTestExecution(ctx, input)
			},
		},
		"start-test-set-generation": {
			Name:   "start-test-set-generation",
			Fields: fields_start_test_set_generation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartTestSetGenerationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_test_set_generation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartTestSetGeneration(ctx, input)
			},
		},
		"stop-bot-recommendation": {
			Name:   "stop-bot-recommendation",
			Fields: fields_stop_bot_recommendation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopBotRecommendationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_bot_recommendation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopBotRecommendation(ctx, input)
			},
		},
		"tag-resource": {
			Name:   "tag-resource",
			Fields: fields_tag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagResource(ctx, input)
			},
		},
		"untag-resource": {
			Name:   "untag-resource",
			Fields: fields_untag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagResource(ctx, input)
			},
		},
		"update-bot": {
			Name:   "update-bot",
			Fields: fields_update_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBot(ctx, input)
			},
		},
		"update-bot-alias": {
			Name:   "update-bot-alias",
			Fields: fields_update_bot_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBotAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_bot_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBotAlias(ctx, input)
			},
		},
		"update-bot-locale": {
			Name:   "update-bot-locale",
			Fields: fields_update_bot_locale,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBotLocaleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_bot_locale, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBotLocale(ctx, input)
			},
		},
		"update-bot-recommendation": {
			Name:   "update-bot-recommendation",
			Fields: fields_update_bot_recommendation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBotRecommendationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_bot_recommendation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBotRecommendation(ctx, input)
			},
		},
		"update-export": {
			Name:   "update-export",
			Fields: fields_update_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateExport(ctx, input)
			},
		},
		"update-intent": {
			Name:   "update-intent",
			Fields: fields_update_intent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIntentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_intent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIntent(ctx, input)
			},
		},
		"update-resource-policy": {
			Name:   "update-resource-policy",
			Fields: fields_update_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResourcePolicy(ctx, input)
			},
		},
		"update-slot": {
			Name:   "update-slot",
			Fields: fields_update_slot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSlotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_slot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSlot(ctx, input)
			},
		},
		"update-slot-type": {
			Name:   "update-slot-type",
			Fields: fields_update_slot_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSlotTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_slot_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSlotType(ctx, input)
			},
		},
		"update-test-set": {
			Name:   "update-test-set",
			Fields: fields_update_test_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTestSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_test_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTestSet(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("lexmodelsv2", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
