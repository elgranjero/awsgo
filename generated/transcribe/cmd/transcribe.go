package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/transcribe"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// transcribeCmd represents the transcribe command
var _transcribeCmd = &cobra.Command{
	Use:   "transcribe",
	Short: "AWS transcribe CLI",
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
		client := transcribe.NewFromConfig(cfg)
		if _transcribeCreateCallAnalyticsCategory {
			transcribe_CreateCallAnalyticsCategory(cfg, client)
			return
		}
		if _transcribeCreateLanguageModel {
			transcribe_CreateLanguageModel(cfg, client)
			return
		}
		if _transcribeCreateMedicalVocabulary {
			transcribe_CreateMedicalVocabulary(cfg, client)
			return
		}
		if _transcribeCreateVocabulary {
			transcribe_CreateVocabulary(cfg, client)
			return
		}
		if _transcribeCreateVocabularyFilter {
			transcribe_CreateVocabularyFilter(cfg, client)
			return
		}
		if _transcribeDeleteCallAnalyticsCategory {
			transcribe_DeleteCallAnalyticsCategory(cfg, client)
			return
		}
		if _transcribeDeleteCallAnalyticsJob {
			transcribe_DeleteCallAnalyticsJob(cfg, client)
			return
		}
		if _transcribeDeleteLanguageModel {
			transcribe_DeleteLanguageModel(cfg, client)
			return
		}
		if _transcribeDeleteMedicalScribeJob {
			transcribe_DeleteMedicalScribeJob(cfg, client)
			return
		}
		if _transcribeDeleteMedicalTranscriptionJob {
			transcribe_DeleteMedicalTranscriptionJob(cfg, client)
			return
		}
		if _transcribeDeleteMedicalVocabulary {
			transcribe_DeleteMedicalVocabulary(cfg, client)
			return
		}
		if _transcribeDeleteTranscriptionJob {
			transcribe_DeleteTranscriptionJob(cfg, client)
			return
		}
		if _transcribeDeleteVocabulary {
			transcribe_DeleteVocabulary(cfg, client)
			return
		}
		if _transcribeDeleteVocabularyFilter {
			transcribe_DeleteVocabularyFilter(cfg, client)
			return
		}
		if _transcribeDescribeLanguageModel {
			transcribe_DescribeLanguageModel(cfg, client)
			return
		}
		if _transcribeGetCallAnalyticsCategory {
			transcribe_GetCallAnalyticsCategory(cfg, client)
			return
		}
		if _transcribeGetCallAnalyticsJob {
			transcribe_GetCallAnalyticsJob(cfg, client)
			return
		}
		if _transcribeGetMedicalScribeJob {
			transcribe_GetMedicalScribeJob(cfg, client)
			return
		}
		if _transcribeGetMedicalTranscriptionJob {
			transcribe_GetMedicalTranscriptionJob(cfg, client)
			return
		}
		if _transcribeGetMedicalVocabulary {
			transcribe_GetMedicalVocabulary(cfg, client)
			return
		}
		if _transcribeGetTranscriptionJob {
			transcribe_GetTranscriptionJob(cfg, client)
			return
		}
		if _transcribeGetVocabulary {
			transcribe_GetVocabulary(cfg, client)
			return
		}
		if _transcribeGetVocabularyFilter {
			transcribe_GetVocabularyFilter(cfg, client)
			return
		}
		if _transcribeListCallAnalyticsCategories {
			transcribe_ListCallAnalyticsCategories(cfg, client)
			return
		}
		if _transcribeListCallAnalyticsJobs {
			transcribe_ListCallAnalyticsJobs(cfg, client)
			return
		}
		if _transcribeListLanguageModels {
			transcribe_ListLanguageModels(cfg, client)
			return
		}
		if _transcribeListMedicalScribeJobs {
			transcribe_ListMedicalScribeJobs(cfg, client)
			return
		}
		if _transcribeListMedicalTranscriptionJobs {
			transcribe_ListMedicalTranscriptionJobs(cfg, client)
			return
		}
		if _transcribeListMedicalVocabularies {
			transcribe_ListMedicalVocabularies(cfg, client)
			return
		}
		if _transcribeListTagsForResource {
			transcribe_ListTagsForResource(cfg, client)
			return
		}
		if _transcribeListTranscriptionJobs {
			transcribe_ListTranscriptionJobs(cfg, client)
			return
		}
		if _transcribeListVocabularies {
			transcribe_ListVocabularies(cfg, client)
			return
		}
		if _transcribeListVocabularyFilters {
			transcribe_ListVocabularyFilters(cfg, client)
			return
		}
		if _transcribeStartCallAnalyticsJob {
			transcribe_StartCallAnalyticsJob(cfg, client)
			return
		}
		if _transcribeStartMedicalScribeJob {
			transcribe_StartMedicalScribeJob(cfg, client)
			return
		}
		if _transcribeStartMedicalTranscriptionJob {
			transcribe_StartMedicalTranscriptionJob(cfg, client)
			return
		}
		if _transcribeStartTranscriptionJob {
			transcribe_StartTranscriptionJob(cfg, client)
			return
		}
		if _transcribeTagResource {
			transcribe_TagResource(cfg, client)
			return
		}
		if _transcribeUntagResource {
			transcribe_UntagResource(cfg, client)
			return
		}
		if _transcribeUpdateCallAnalyticsCategory {
			transcribe_UpdateCallAnalyticsCategory(cfg, client)
			return
		}
		if _transcribeUpdateMedicalVocabulary {
			transcribe_UpdateMedicalVocabulary(cfg, client)
			return
		}
		if _transcribeUpdateVocabulary {
			transcribe_UpdateVocabulary(cfg, client)
			return
		}
		if _transcribeUpdateVocabularyFilter {
			transcribe_UpdateVocabularyFilter(cfg, client)
			return
		}

	},
}

var (
	_transcribeCreateCallAnalyticsCategory   bool
	_transcribeCreateLanguageModel           bool
	_transcribeCreateMedicalVocabulary       bool
	_transcribeCreateVocabulary              bool
	_transcribeCreateVocabularyFilter        bool
	_transcribeDeleteCallAnalyticsCategory   bool
	_transcribeDeleteCallAnalyticsJob        bool
	_transcribeDeleteLanguageModel           bool
	_transcribeDeleteMedicalScribeJob        bool
	_transcribeDeleteMedicalTranscriptionJob bool
	_transcribeDeleteMedicalVocabulary       bool
	_transcribeDeleteTranscriptionJob        bool
	_transcribeDeleteVocabulary              bool
	_transcribeDeleteVocabularyFilter        bool
	_transcribeDescribeLanguageModel         bool
	_transcribeGetCallAnalyticsCategory      bool
	_transcribeGetCallAnalyticsJob           bool
	_transcribeGetMedicalScribeJob           bool
	_transcribeGetMedicalTranscriptionJob    bool
	_transcribeGetMedicalVocabulary          bool
	_transcribeGetTranscriptionJob           bool
	_transcribeGetVocabulary                 bool
	_transcribeGetVocabularyFilter           bool
	_transcribeListCallAnalyticsCategories   bool
	_transcribeListCallAnalyticsJobs         bool
	_transcribeListLanguageModels            bool
	_transcribeListMedicalScribeJobs         bool
	_transcribeListMedicalTranscriptionJobs  bool
	_transcribeListMedicalVocabularies       bool
	_transcribeListTagsForResource           bool
	_transcribeListTranscriptionJobs         bool
	_transcribeListVocabularies              bool
	_transcribeListVocabularyFilters         bool
	_transcribeStartCallAnalyticsJob         bool
	_transcribeStartMedicalScribeJob         bool
	_transcribeStartMedicalTranscriptionJob  bool
	_transcribeStartTranscriptionJob         bool
	_transcribeTagResource                   bool
	_transcribeUntagResource                 bool
	_transcribeUpdateCallAnalyticsCategory   bool
	_transcribeUpdateMedicalVocabulary       bool
	_transcribeUpdateVocabulary              bool
	_transcribeUpdateVocabularyFilter        bool

	_transcribeBaseModelName               string
	_transcribeCallAnalyticsJobName        string
	_transcribeCategoryName                string
	_transcribeChannelDefinitions          string
	_transcribeContentIdentificationType   string
	_transcribeContentRedaction            string
	_transcribeDataAccessRoleArn           string
	_transcribeIdentifyLanguage            string
	_transcribeIdentifyMultipleLanguages   string
	_transcribeInputDataConfig             string
	_transcribeInputType                   string
	_transcribeJobExecutionSettings        string
	_transcribeJobNameContains             string
	_transcribeKMSEncryptionContext        string
	_transcribeLanguageCode                string
	_transcribeLanguageIdSettings          string
	_transcribeLanguageOptions             string
	_transcribeMaxResults                  string
	_transcribeMedia                       string
	_transcribeMediaFormat                 string
	_transcribeMediaSampleRateHertz        string
	_transcribeMedicalScribeContext        string
	_transcribeMedicalScribeJobName        string
	_transcribeMedicalTranscriptionJobName string
	_transcribeModelName                   string
	_transcribeModelSettings               string
	_transcribeNameContains                string
	_transcribeNextToken                   string
	_transcribeOutputBucketName            string
	_transcribeOutputEncryptionKMSKeyId    string
	_transcribeOutputKey                   string
	_transcribeOutputLocation              string
	_transcribePhrases                     []string
	_transcribeResourceArn                 string
	_transcribeRules                       string
	_transcribeSettings                    string
	_transcribeSpecialty                   string
	_transcribeStateEquals                 string
	_transcribeStatus                      string
	_transcribeStatusEquals                string
	_transcribeSubtitles                   string
	_transcribeTagKeys                     []string
	_transcribeTags                        string
	_transcribeToxicityDetection           string
	_transcribeTranscriptionJobName        string
	_transcribeType                        string
	_transcribeVocabularyFileUri           string
	_transcribeVocabularyFilterFileUri     string
	_transcribeVocabularyFilterName        string
	_transcribeVocabularyName              string
	_transcribeWords                       []string
)

// Creates a new Call Analytics category.
// All categories are automatically applied to your Call Analytics transcriptions.
// Note that in order to apply categories to your transcriptions, you must create
// them before submitting your transcription request, as categories cannot be
// applied retroactively.
//
// When creating a new category, you can use the InputType parameter to label the
// category as a POST_CALL or a REAL_TIME category. POST_CALL categories can only
// be applied to post-call transcriptions and REAL_TIME categories can only be
// applied to real-time transcriptions. If you do not include InputType , your
// category is created as a POST_CALL category by default.
//
// Call Analytics categories are composed of rules. For each category, you must
// create between 1 and 20 rules. Rules can include these parameters: , , , and .
//
// To update an existing category, see .
//
// To learn more about Call Analytics categories, see [Creating categories for post-call transcriptions] and [Creating categories for real-time transcriptions].
//
// [Creating categories for post-call transcriptions]: https://docs.aws.amazon.com/transcribe/latest/dg/tca-categories-batch.html
// [Creating categories for real-time transcriptions]: https://docs.aws.amazon.com/transcribe/latest/dg/tca-categories-stream.html
func transcribe_CreateCallAnalyticsCategory(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.CreateCallAnalyticsCategoryInput{
		// CategoryName: *string, // Required
		// Rules: []types.Rule, // Required
	}

	if len(_transcribeCategoryName) > 0 {
		input.CategoryName = aws.String(_transcribeCategoryName)
	}
	if len(_transcribeRules) > 0 {
		if err := assignInputField(input, "Rules", _transcribeRules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_transcribeInputType) > 0 {
		if err := assignInputField(input, "InputType", _transcribeInputType); err != nil {
			log.Errorf("invalid --input-type: %s", err.Error())
			return
		}
	}
	if len(_transcribeTags) > 0 {
		if err := assignInputField(input, "Tags", _transcribeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCallAnalyticsCategory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new custom language model.
// When creating a new custom language model, you must specify:
//
// - If you want a Wideband (audio sample rates over 16,000 Hz) or Narrowband
// (audio sample rates under 16,000 Hz) base model
//
// - The location of your training and tuning files (this must be an Amazon S3
// URI)
//
// - The language of your model
//
// - A unique name for your model
func transcribe_CreateLanguageModel(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.CreateLanguageModelInput{
		// BaseModelName: types.BaseModelName, // Required
		// InputDataConfig: *types.InputDataConfig, // Required
		// LanguageCode: types.CLMLanguageCode, // Required
		// ModelName: *string, // Required
	}

	if len(_transcribeBaseModelName) > 0 {
		if err := assignInputField(input, "BaseModelName", _transcribeBaseModelName); err != nil {
			log.Errorf("invalid --base-model-name: %s", err.Error())
			return
		}
	}
	if len(_transcribeInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _transcribeInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_transcribeLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _transcribeLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_transcribeModelName) > 0 {
		input.ModelName = aws.String(_transcribeModelName)
	}
	if len(_transcribeTags) > 0 {
		if err := assignInputField(input, "Tags", _transcribeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLanguageModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new custom medical vocabulary.
// Before creating a new custom medical vocabulary, you must first upload a text
// file that contains your vocabulary table into an Amazon S3 bucket. Note that
// this differs from , where you can include a list of terms within your request
// using the Phrases flag; CreateMedicalVocabulary does not support the Phrases
// flag and only accepts vocabularies in table format.
//
// Each language has a character set that contains all allowed characters for that
// specific language. If you use unsupported characters, your custom vocabulary
// request fails. Refer to [Character Sets for Custom Vocabularies]to get the character set for your language.
//
// For more information, see [Custom vocabularies].
//
// [Custom vocabularies]: https://docs.aws.amazon.com/transcribe/latest/dg/custom-vocabulary.html
// [Character Sets for Custom Vocabularies]: https://docs.aws.amazon.com/transcribe/latest/dg/charsets.html
func transcribe_CreateMedicalVocabulary(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.CreateMedicalVocabularyInput{
		// LanguageCode: types.LanguageCode, // Required
		// VocabularyFileUri: *string, // Required
		// VocabularyName: *string, // Required
	}

	if len(_transcribeLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _transcribeLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_transcribeVocabularyFileUri) > 0 {
		input.VocabularyFileUri = aws.String(_transcribeVocabularyFileUri)
	}
	if len(_transcribeVocabularyName) > 0 {
		input.VocabularyName = aws.String(_transcribeVocabularyName)
	}
	if len(_transcribeTags) > 0 {
		if err := assignInputField(input, "Tags", _transcribeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMedicalVocabulary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new custom vocabulary.
// When creating a new custom vocabulary, you can either upload a text file that
// contains your new entries, phrases, and terms into an Amazon S3 bucket and
// include the URI in your request. Or you can include a list of terms directly in
// your request using the Phrases flag.
//
// Each language has a character set that contains all allowed characters for that
// specific language. If you use unsupported characters, your custom vocabulary
// request fails. Refer to [Character Sets for Custom Vocabularies]to get the character set for your language.
//
// For more information, see [Custom vocabularies].
//
// [Custom vocabularies]: https://docs.aws.amazon.com/transcribe/latest/dg/custom-vocabulary.html
// [Character Sets for Custom Vocabularies]: https://docs.aws.amazon.com/transcribe/latest/dg/charsets.html
func transcribe_CreateVocabulary(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.CreateVocabularyInput{
		// LanguageCode: types.LanguageCode, // Required
		// VocabularyName: *string, // Required
	}

	if len(_transcribeLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _transcribeLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_transcribeVocabularyName) > 0 {
		input.VocabularyName = aws.String(_transcribeVocabularyName)
	}
	if len(_transcribeDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_transcribeDataAccessRoleArn)
	}
	if len(_transcribePhrases) > 0 {
		input.Phrases = append([]string(nil), _transcribePhrases...)
	}
	if len(_transcribeTags) > 0 {
		if err := assignInputField(input, "Tags", _transcribeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_transcribeVocabularyFileUri) > 0 {
		input.VocabularyFileUri = aws.String(_transcribeVocabularyFileUri)
	}

	if resp, err := client.CreateVocabulary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new custom vocabulary filter.
// You can use custom vocabulary filters to mask, delete, or flag specific words
// from your transcript. Custom vocabulary filters are commonly used to mask
// profanity in transcripts.
//
// Each language has a character set that contains all allowed characters for that
// specific language. If you use unsupported characters, your custom vocabulary
// filter request fails. Refer to [Character Sets for Custom Vocabularies]to get the character set for your language.
//
// For more information, see [Vocabulary filtering].
//
// [Character Sets for Custom Vocabularies]: https://docs.aws.amazon.com/transcribe/latest/dg/charsets.html
// [Vocabulary filtering]: https://docs.aws.amazon.com/transcribe/latest/dg/vocabulary-filtering.html
func transcribe_CreateVocabularyFilter(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.CreateVocabularyFilterInput{
		// LanguageCode: types.LanguageCode, // Required
		// VocabularyFilterName: *string, // Required
	}

	if len(_transcribeLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _transcribeLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_transcribeVocabularyFilterName) > 0 {
		input.VocabularyFilterName = aws.String(_transcribeVocabularyFilterName)
	}
	if len(_transcribeDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_transcribeDataAccessRoleArn)
	}
	if len(_transcribeTags) > 0 {
		if err := assignInputField(input, "Tags", _transcribeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_transcribeVocabularyFilterFileUri) > 0 {
		input.VocabularyFilterFileUri = aws.String(_transcribeVocabularyFilterFileUri)
	}
	if len(_transcribeWords) > 0 {
		input.Words = append([]string(nil), _transcribeWords...)
	}

	if resp, err := client.CreateVocabularyFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Call Analytics category. To use this operation, specify the name of
// the category you want to delete using CategoryName . Category names are case
// sensitive.
func transcribe_DeleteCallAnalyticsCategory(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.DeleteCallAnalyticsCategoryInput{
		// CategoryName: *string, // Required
	}

	if len(_transcribeCategoryName) > 0 {
		input.CategoryName = aws.String(_transcribeCategoryName)
	}

	if resp, err := client.DeleteCallAnalyticsCategory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Call Analytics job. To use this operation, specify the name of the
// job you want to delete using CallAnalyticsJobName . Job names are case sensitive.
func transcribe_DeleteCallAnalyticsJob(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.DeleteCallAnalyticsJobInput{
		// CallAnalyticsJobName: *string, // Required
	}

	if len(_transcribeCallAnalyticsJobName) > 0 {
		input.CallAnalyticsJobName = aws.String(_transcribeCallAnalyticsJobName)
	}

	if resp, err := client.DeleteCallAnalyticsJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a custom language model. To use this operation, specify the name of the
// language model you want to delete using ModelName . custom language model names
// are case sensitive.
func transcribe_DeleteLanguageModel(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.DeleteLanguageModelInput{
		// ModelName: *string, // Required
	}

	if len(_transcribeModelName) > 0 {
		input.ModelName = aws.String(_transcribeModelName)
	}

	if resp, err := client.DeleteLanguageModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Medical Scribe job. To use this operation, specify the name of the
// job you want to delete using MedicalScribeJobName . Job names are case sensitive.
func transcribe_DeleteMedicalScribeJob(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.DeleteMedicalScribeJobInput{
		// MedicalScribeJobName: *string, // Required
	}

	if len(_transcribeMedicalScribeJobName) > 0 {
		input.MedicalScribeJobName = aws.String(_transcribeMedicalScribeJobName)
	}

	if resp, err := client.DeleteMedicalScribeJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a medical transcription job. To use this operation, specify the name of
// the job you want to delete using MedicalTranscriptionJobName . Job names are
// case sensitive.
func transcribe_DeleteMedicalTranscriptionJob(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.DeleteMedicalTranscriptionJobInput{
		// MedicalTranscriptionJobName: *string, // Required
	}

	if len(_transcribeMedicalTranscriptionJobName) > 0 {
		input.MedicalTranscriptionJobName = aws.String(_transcribeMedicalTranscriptionJobName)
	}

	if resp, err := client.DeleteMedicalTranscriptionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a custom medical vocabulary. To use this operation, specify the name of
// the custom vocabulary you want to delete using VocabularyName . Custom
// vocabulary names are case sensitive.
func transcribe_DeleteMedicalVocabulary(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.DeleteMedicalVocabularyInput{
		// VocabularyName: *string, // Required
	}

	if len(_transcribeVocabularyName) > 0 {
		input.VocabularyName = aws.String(_transcribeVocabularyName)
	}

	if resp, err := client.DeleteMedicalVocabulary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a transcription job. To use this operation, specify the name of the job
// you want to delete using TranscriptionJobName . Job names are case sensitive.
func transcribe_DeleteTranscriptionJob(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.DeleteTranscriptionJobInput{
		// TranscriptionJobName: *string, // Required
	}

	if len(_transcribeTranscriptionJobName) > 0 {
		input.TranscriptionJobName = aws.String(_transcribeTranscriptionJobName)
	}

	if resp, err := client.DeleteTranscriptionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a custom vocabulary. To use this operation, specify the name of the
// custom vocabulary you want to delete using VocabularyName . Custom vocabulary
// names are case sensitive.
func transcribe_DeleteVocabulary(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.DeleteVocabularyInput{
		// VocabularyName: *string, // Required
	}

	if len(_transcribeVocabularyName) > 0 {
		input.VocabularyName = aws.String(_transcribeVocabularyName)
	}

	if resp, err := client.DeleteVocabulary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a custom vocabulary filter. To use this operation, specify the name of
// the custom vocabulary filter you want to delete using VocabularyFilterName .
// Custom vocabulary filter names are case sensitive.
func transcribe_DeleteVocabularyFilter(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.DeleteVocabularyFilterInput{
		// VocabularyFilterName: *string, // Required
	}

	if len(_transcribeVocabularyFilterName) > 0 {
		input.VocabularyFilterName = aws.String(_transcribeVocabularyFilterName)
	}

	if resp, err := client.DeleteVocabularyFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about the specified custom language model.
// This operation also shows if the base language model that you used to create
// your custom language model has been updated. If Amazon Transcribe has updated
// the base model, you can create a new custom language model using the updated
// base model.
//
// If you tried to create a new custom language model and the request wasn't
// successful, you can use DescribeLanguageModel to help identify the reason for
// this failure.
func transcribe_DescribeLanguageModel(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.DescribeLanguageModelInput{
		// ModelName: *string, // Required
	}

	if len(_transcribeModelName) > 0 {
		input.ModelName = aws.String(_transcribeModelName)
	}

	if resp, err := client.DescribeLanguageModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about the specified Call Analytics category.
// To get a list of your Call Analytics categories, use the operation.
func transcribe_GetCallAnalyticsCategory(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.GetCallAnalyticsCategoryInput{
		// CategoryName: *string, // Required
	}

	if len(_transcribeCategoryName) > 0 {
		input.CategoryName = aws.String(_transcribeCategoryName)
	}

	if resp, err := client.GetCallAnalyticsCategory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about the specified Call Analytics job.
// To view the job's status, refer to CallAnalyticsJobStatus . If the status is
// COMPLETED , the job is finished. You can find your completed transcript at the
// URI specified in TranscriptFileUri . If the status is FAILED , FailureReason
// provides details on why your transcription job failed.
//
// If you enabled personally identifiable information (PII) redaction, the
// redacted transcript appears at the location specified in
// RedactedTranscriptFileUri .
//
// If you chose to redact the audio in your media file, you can find your redacted
// media file at the location specified in RedactedMediaFileUri .
//
// To get a list of your Call Analytics jobs, use the operation.
func transcribe_GetCallAnalyticsJob(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.GetCallAnalyticsJobInput{
		// CallAnalyticsJobName: *string, // Required
	}

	if len(_transcribeCallAnalyticsJobName) > 0 {
		input.CallAnalyticsJobName = aws.String(_transcribeCallAnalyticsJobName)
	}

	if resp, err := client.GetCallAnalyticsJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about the specified Medical Scribe job.
// To view the status of the specified medical transcription job, check the
// MedicalScribeJobStatus field. If the status is COMPLETED , the job is finished.
// You can find the results at the location specified in MedicalScribeOutput . If
// the status is FAILED , FailureReason provides details on why your Medical
// Scribe job failed.
//
// To get a list of your Medical Scribe jobs, use the operation.
func transcribe_GetMedicalScribeJob(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.GetMedicalScribeJobInput{
		// MedicalScribeJobName: *string, // Required
	}

	if len(_transcribeMedicalScribeJobName) > 0 {
		input.MedicalScribeJobName = aws.String(_transcribeMedicalScribeJobName)
	}

	if resp, err := client.GetMedicalScribeJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about the specified medical transcription job.
// To view the status of the specified medical transcription job, check the
// TranscriptionJobStatus field. If the status is COMPLETED , the job is finished.
// You can find the results at the location specified in TranscriptFileUri . If the
// status is FAILED , FailureReason provides details on why your transcription job
// failed.
//
// To get a list of your medical transcription jobs, use the operation.
func transcribe_GetMedicalTranscriptionJob(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.GetMedicalTranscriptionJobInput{
		// MedicalTranscriptionJobName: *string, // Required
	}

	if len(_transcribeMedicalTranscriptionJobName) > 0 {
		input.MedicalTranscriptionJobName = aws.String(_transcribeMedicalTranscriptionJobName)
	}

	if resp, err := client.GetMedicalTranscriptionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about the specified custom medical vocabulary.
// To view the status of the specified custom medical vocabulary, check the
// VocabularyState field. If the status is READY , your custom vocabulary is
// available to use. If the status is FAILED , FailureReason provides details on
// why your vocabulary failed.
//
// To get a list of your custom medical vocabularies, use the operation.
func transcribe_GetMedicalVocabulary(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.GetMedicalVocabularyInput{
		// VocabularyName: *string, // Required
	}

	if len(_transcribeVocabularyName) > 0 {
		input.VocabularyName = aws.String(_transcribeVocabularyName)
	}

	if resp, err := client.GetMedicalVocabulary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about the specified transcription job.
// To view the status of the specified transcription job, check the
// TranscriptionJobStatus field. If the status is COMPLETED , the job is finished.
// You can find the results at the location specified in TranscriptFileUri . If the
// status is FAILED , FailureReason provides details on why your transcription job
// failed.
//
// If you enabled content redaction, the redacted transcript can be found at the
// location specified in RedactedTranscriptFileUri .
//
// To get a list of your transcription jobs, use the operation.
func transcribe_GetTranscriptionJob(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.GetTranscriptionJobInput{
		// TranscriptionJobName: *string, // Required
	}

	if len(_transcribeTranscriptionJobName) > 0 {
		input.TranscriptionJobName = aws.String(_transcribeTranscriptionJobName)
	}

	if resp, err := client.GetTranscriptionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about the specified custom vocabulary.
// To view the status of the specified custom vocabulary, check the VocabularyState
// field. If the status is READY , your custom vocabulary is available to use. If
// the status is FAILED , FailureReason provides details on why your custom
// vocabulary failed.
//
// To get a list of your custom vocabularies, use the operation.
func transcribe_GetVocabulary(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.GetVocabularyInput{
		// VocabularyName: *string, // Required
	}

	if len(_transcribeVocabularyName) > 0 {
		input.VocabularyName = aws.String(_transcribeVocabularyName)
	}

	if resp, err := client.GetVocabulary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about the specified custom vocabulary filter.
// To get a list of your custom vocabulary filters, use the operation.
func transcribe_GetVocabularyFilter(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.GetVocabularyFilterInput{
		// VocabularyFilterName: *string, // Required
	}

	if len(_transcribeVocabularyFilterName) > 0 {
		input.VocabularyFilterName = aws.String(_transcribeVocabularyFilterName)
	}

	if resp, err := client.GetVocabularyFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a list of Call Analytics categories, including all rules that make up
// each category.
//
// To get detailed information about a specific Call Analytics category, use the
// operation.
func transcribe_ListCallAnalyticsCategories(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.ListCallAnalyticsCategoriesInput{}

	if len(_transcribeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _transcribeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_transcribeNextToken) > 0 {
		input.NextToken = aws.String(_transcribeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCallAnalyticsCategories(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*transcribe.ListCallAnalyticsCategoriesOutput
	p := transcribe.NewListCallAnalyticsCategoriesPaginator(client, input)
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

// Provides a list of Call Analytics jobs that match the specified criteria. If no
// criteria are specified, all Call Analytics jobs are returned.
//
// To get detailed information about a specific Call Analytics job, use the
// operation.
func transcribe_ListCallAnalyticsJobs(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.ListCallAnalyticsJobsInput{}

	if len(_transcribeJobNameContains) > 0 {
		input.JobNameContains = aws.String(_transcribeJobNameContains)
	}
	if len(_transcribeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _transcribeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_transcribeNextToken) > 0 {
		input.NextToken = aws.String(_transcribeNextToken)
	}
	if len(_transcribeStatus) > 0 {
		if err := assignInputField(input, "Status", _transcribeStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCallAnalyticsJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*transcribe.ListCallAnalyticsJobsOutput
	p := transcribe.NewListCallAnalyticsJobsPaginator(client, input)
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

// Provides a list of custom language models that match the specified criteria. If
// no criteria are specified, all custom language models are returned.
//
// To get detailed information about a specific custom language model, use the
// operation.
func transcribe_ListLanguageModels(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.ListLanguageModelsInput{}

	if len(_transcribeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _transcribeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_transcribeNameContains) > 0 {
		input.NameContains = aws.String(_transcribeNameContains)
	}
	if len(_transcribeNextToken) > 0 {
		input.NextToken = aws.String(_transcribeNextToken)
	}
	if len(_transcribeStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _transcribeStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListLanguageModels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*transcribe.ListLanguageModelsOutput
	p := transcribe.NewListLanguageModelsPaginator(client, input)
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

// Provides a list of Medical Scribe jobs that match the specified criteria. If no
// criteria are specified, all Medical Scribe jobs are returned.
//
// To get detailed information about a specific Medical Scribe job, use the
// operation.
func transcribe_ListMedicalScribeJobs(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.ListMedicalScribeJobsInput{}

	if len(_transcribeJobNameContains) > 0 {
		input.JobNameContains = aws.String(_transcribeJobNameContains)
	}
	if len(_transcribeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _transcribeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_transcribeNextToken) > 0 {
		input.NextToken = aws.String(_transcribeNextToken)
	}
	if len(_transcribeStatus) > 0 {
		if err := assignInputField(input, "Status", _transcribeStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListMedicalScribeJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*transcribe.ListMedicalScribeJobsOutput
	p := transcribe.NewListMedicalScribeJobsPaginator(client, input)
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

// Provides a list of medical transcription jobs that match the specified
// criteria. If no criteria are specified, all medical transcription jobs are
// returned.
//
// To get detailed information about a specific medical transcription job, use the
// operation.
func transcribe_ListMedicalTranscriptionJobs(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.ListMedicalTranscriptionJobsInput{}

	if len(_transcribeJobNameContains) > 0 {
		input.JobNameContains = aws.String(_transcribeJobNameContains)
	}
	if len(_transcribeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _transcribeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_transcribeNextToken) > 0 {
		input.NextToken = aws.String(_transcribeNextToken)
	}
	if len(_transcribeStatus) > 0 {
		if err := assignInputField(input, "Status", _transcribeStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListMedicalTranscriptionJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*transcribe.ListMedicalTranscriptionJobsOutput
	p := transcribe.NewListMedicalTranscriptionJobsPaginator(client, input)
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

// Provides a list of custom medical vocabularies that match the specified
// criteria. If no criteria are specified, all custom medical vocabularies are
// returned.
//
// To get detailed information about a specific custom medical vocabulary, use the
// operation.
func transcribe_ListMedicalVocabularies(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.ListMedicalVocabulariesInput{}

	if len(_transcribeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _transcribeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_transcribeNameContains) > 0 {
		input.NameContains = aws.String(_transcribeNameContains)
	}
	if len(_transcribeNextToken) > 0 {
		input.NextToken = aws.String(_transcribeNextToken)
	}
	if len(_transcribeStateEquals) > 0 {
		if err := assignInputField(input, "StateEquals", _transcribeStateEquals); err != nil {
			log.Errorf("invalid --state-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListMedicalVocabularies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*transcribe.ListMedicalVocabulariesOutput
	p := transcribe.NewListMedicalVocabulariesPaginator(client, input)
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

// Lists all tags associated with the specified transcription job, vocabulary,
// model, or resource.
//
// To learn more about using tags with Amazon Transcribe, refer to [Tagging resources].
//
// [Tagging resources]: https://docs.aws.amazon.com/transcribe/latest/dg/tagging.html
func transcribe_ListTagsForResource(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_transcribeResourceArn) > 0 {
		input.ResourceArn = aws.String(_transcribeResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a list of transcription jobs that match the specified criteria. If no
// criteria are specified, all transcription jobs are returned.
//
// To get detailed information about a specific transcription job, use the
// operation.
func transcribe_ListTranscriptionJobs(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.ListTranscriptionJobsInput{}

	if len(_transcribeJobNameContains) > 0 {
		input.JobNameContains = aws.String(_transcribeJobNameContains)
	}
	if len(_transcribeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _transcribeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_transcribeNextToken) > 0 {
		input.NextToken = aws.String(_transcribeNextToken)
	}
	if len(_transcribeStatus) > 0 {
		if err := assignInputField(input, "Status", _transcribeStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTranscriptionJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*transcribe.ListTranscriptionJobsOutput
	p := transcribe.NewListTranscriptionJobsPaginator(client, input)
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

// Provides a list of custom vocabularies that match the specified criteria. If no
// criteria are specified, all custom vocabularies are returned.
//
// To get detailed information about a specific custom vocabulary, use the
// operation.
func transcribe_ListVocabularies(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.ListVocabulariesInput{}

	if len(_transcribeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _transcribeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_transcribeNameContains) > 0 {
		input.NameContains = aws.String(_transcribeNameContains)
	}
	if len(_transcribeNextToken) > 0 {
		input.NextToken = aws.String(_transcribeNextToken)
	}
	if len(_transcribeStateEquals) > 0 {
		if err := assignInputField(input, "StateEquals", _transcribeStateEquals); err != nil {
			log.Errorf("invalid --state-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListVocabularies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*transcribe.ListVocabulariesOutput
	p := transcribe.NewListVocabulariesPaginator(client, input)
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

// Provides a list of custom vocabulary filters that match the specified criteria.
// If no criteria are specified, all custom vocabularies are returned.
//
// To get detailed information about a specific custom vocabulary filter, use the
// operation.
func transcribe_ListVocabularyFilters(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.ListVocabularyFiltersInput{}

	if len(_transcribeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _transcribeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_transcribeNameContains) > 0 {
		input.NameContains = aws.String(_transcribeNameContains)
	}
	if len(_transcribeNextToken) > 0 {
		input.NextToken = aws.String(_transcribeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListVocabularyFilters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*transcribe.ListVocabularyFiltersOutput
	p := transcribe.NewListVocabularyFiltersPaginator(client, input)
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

// Transcribes the audio from a customer service call and applies any additional
// Request Parameters you choose to include in your request.
//
// In addition to many standard transcription features, Call Analytics provides
// you with call characteristics, call summarization, speaker sentiment, and
// optional redaction of your text transcript and your audio file. You can also
// apply custom categories to flag specified conditions. To learn more about these
// features and insights, refer to [Analyzing call center audio with Call Analytics].
//
// If you want to apply categories to your Call Analytics job, you must create
// them before submitting your job request. Categories cannot be retroactively
// applied to a job. To create a new category, use the operation. To learn more
// about Call Analytics categories, see [Creating categories for post-call transcriptions]and [Creating categories for real-time transcriptions].
//
// To make a StartCallAnalyticsJob request, you must first upload your media file
// into an Amazon S3 bucket; you can then specify the Amazon S3 location of the
// file using the Media parameter.
//
// Job queuing is available for Call Analytics jobs. If you pass a
// DataAccessRoleArn in your request and you exceed your Concurrent Job Limit, your
// job will automatically be added to a queue to be processed once your concurrent
// job count is below the limit.
//
// You must include the following parameters in your StartCallAnalyticsJob request:
//
// - region : The Amazon Web Services Region where you are making your request.
// For a list of Amazon Web Services Regions supported with Amazon Transcribe,
// refer to [Amazon Transcribe endpoints and quotas].
//
// - CallAnalyticsJobName : A custom name that you create for your transcription
// job that's unique within your Amazon Web Services account.
//
// - Media ( MediaFileUri or RedactedMediaFileUri ): The Amazon S3 location of
// your media file.
//
// With Call Analytics, you can redact the audio contained in your media file by
// including RedactedMediaFileUri , instead of MediaFileUri , to specify the
// location of your input audio. If you choose to redact your audio, you can find
// your redacted media at the location specified in the RedactedMediaFileUri field
// of your response.
//
// [Amazon Transcribe endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/transcribe.html
// [Analyzing call center audio with Call Analytics]: https://docs.aws.amazon.com/transcribe/latest/dg/call-analytics.html
// [Creating categories for post-call transcriptions]: https://docs.aws.amazon.com/transcribe/latest/dg/tca-categories-batch.html
// [Creating categories for real-time transcriptions]: https://docs.aws.amazon.com/transcribe/latest/dg/tca-categories-stream.html
func transcribe_StartCallAnalyticsJob(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.StartCallAnalyticsJobInput{
		// CallAnalyticsJobName: *string, // Required
		// Media: *types.Media, // Required
	}

	if len(_transcribeCallAnalyticsJobName) > 0 {
		input.CallAnalyticsJobName = aws.String(_transcribeCallAnalyticsJobName)
	}
	if len(_transcribeMedia) > 0 {
		if err := assignInputField(input, "Media", _transcribeMedia); err != nil {
			log.Errorf("invalid --media: %s", err.Error())
			return
		}
	}
	if len(_transcribeChannelDefinitions) > 0 {
		if err := assignInputField(input, "ChannelDefinitions", _transcribeChannelDefinitions); err != nil {
			log.Errorf("invalid --channel-definitions: %s", err.Error())
			return
		}
	}
	if len(_transcribeDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_transcribeDataAccessRoleArn)
	}
	if len(_transcribeOutputEncryptionKMSKeyId) > 0 {
		input.OutputEncryptionKMSKeyId = aws.String(_transcribeOutputEncryptionKMSKeyId)
	}
	if len(_transcribeOutputLocation) > 0 {
		input.OutputLocation = aws.String(_transcribeOutputLocation)
	}
	if len(_transcribeSettings) > 0 {
		if err := assignInputField(input, "Settings", _transcribeSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}
	if len(_transcribeTags) > 0 {
		if err := assignInputField(input, "Tags", _transcribeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartCallAnalyticsJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Transcribes patient-clinician conversations and generates clinical notes.
// Amazon Web Services HealthScribe automatically provides rich conversation
// transcripts, identifies speaker roles, classifies dialogues, extracts medical
// terms, and generates preliminary clinical notes. To learn more about these
// features, refer to [Amazon Web Services HealthScribe].
//
// To make a StartMedicalScribeJob request, you must first upload your media file
// into an Amazon S3 bucket; you can then specify the Amazon S3 location of the
// file using the Media parameter.
//
// You must include the following parameters in your StartMedicalTranscriptionJob
// request:
//
// - DataAccessRoleArn : The ARN of an IAM role with the these minimum
// permissions: read permission on input file Amazon S3 bucket specified in Media
// , write permission on the Amazon S3 bucket specified in OutputBucketName , and
// full permissions on the KMS key specified in OutputEncryptionKMSKeyId (if
// set). The role should also allow transcribe.amazonaws.com to assume it.
//
// - Media ( MediaFileUri ): The Amazon S3 location of your media file.
//
// - MedicalScribeJobName : A custom name you create for your MedicalScribe job
// that is unique within your Amazon Web Services account.
//
// - OutputBucketName : The Amazon S3 bucket where you want your output files
// stored.
//
// - Settings : A MedicalScribeSettings object that must set exactly one of
// ShowSpeakerLabels or ChannelIdentification to true. If ShowSpeakerLabels is
// true, MaxSpeakerLabels must also be set.
//
// - ChannelDefinitions : A MedicalScribeChannelDefinitions array should be set
// if and only if the ChannelIdentification value of Settings is set to true.
//
// [Amazon Web Services HealthScribe]: https://docs.aws.amazon.com/transcribe/latest/dg/health-scribe.html
func transcribe_StartMedicalScribeJob(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.StartMedicalScribeJobInput{
		// DataAccessRoleArn: *string, // Required
		// Media: *types.Media, // Required
		// MedicalScribeJobName: *string, // Required
		// OutputBucketName: *string, // Required
		// Settings: *types.MedicalScribeSettings, // Required
	}

	if len(_transcribeDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_transcribeDataAccessRoleArn)
	}
	if len(_transcribeMedia) > 0 {
		if err := assignInputField(input, "Media", _transcribeMedia); err != nil {
			log.Errorf("invalid --media: %s", err.Error())
			return
		}
	}
	if len(_transcribeMedicalScribeJobName) > 0 {
		input.MedicalScribeJobName = aws.String(_transcribeMedicalScribeJobName)
	}
	if len(_transcribeOutputBucketName) > 0 {
		input.OutputBucketName = aws.String(_transcribeOutputBucketName)
	}
	if len(_transcribeSettings) > 0 {
		if err := assignInputField(input, "Settings", _transcribeSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}
	if len(_transcribeChannelDefinitions) > 0 {
		if err := assignInputField(input, "ChannelDefinitions", _transcribeChannelDefinitions); err != nil {
			log.Errorf("invalid --channel-definitions: %s", err.Error())
			return
		}
	}
	if len(_transcribeKMSEncryptionContext) > 0 {
		if err := assignInputField(input, "KMSEncryptionContext", _transcribeKMSEncryptionContext); err != nil {
			log.Errorf("invalid --kms-encryption-context: %s", err.Error())
			return
		}
	}
	if len(_transcribeMedicalScribeContext) > 0 {
		if err := assignInputField(input, "MedicalScribeContext", _transcribeMedicalScribeContext); err != nil {
			log.Errorf("invalid --medical-scribe-context: %s", err.Error())
			return
		}
	}
	if len(_transcribeOutputEncryptionKMSKeyId) > 0 {
		input.OutputEncryptionKMSKeyId = aws.String(_transcribeOutputEncryptionKMSKeyId)
	}
	if len(_transcribeTags) > 0 {
		if err := assignInputField(input, "Tags", _transcribeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartMedicalScribeJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Transcribes the audio from a medical dictation or conversation and applies any
// additional Request Parameters you choose to include in your request.
//
// In addition to many standard transcription features, Amazon Transcribe Medical
// provides you with a robust medical vocabulary and, optionally, content
// identification, which adds flags to personal health information (PHI). To learn
// more about these features, refer to [How Amazon Transcribe Medical works].
//
// To make a StartMedicalTranscriptionJob request, you must first upload your
// media file into an Amazon S3 bucket; you can then specify the Amazon S3 location
// of the file using the Media parameter.
//
// You must include the following parameters in your StartMedicalTranscriptionJob
// request:
//
// - region : The Amazon Web Services Region where you are making your request.
// For a list of Amazon Web Services Regions supported with Amazon Transcribe,
// refer to [Amazon Transcribe endpoints and quotas].
//
// - MedicalTranscriptionJobName : A custom name you create for your
// transcription job that is unique within your Amazon Web Services account.
//
// - Media ( MediaFileUri ): The Amazon S3 location of your media file.
//
// - LanguageCode : This must be en-US .
//
// - OutputBucketName : The Amazon S3 bucket where you want your transcript
// stored. If you want your output stored in a sub-folder of this bucket, you must
// also include OutputKey .
//
// - Specialty : This must be PRIMARYCARE .
//
// - Type : Choose whether your audio is a conversation or a dictation.
//
// [How Amazon Transcribe Medical works]: https://docs.aws.amazon.com/transcribe/latest/dg/how-it-works-med.html
// [Amazon Transcribe endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/transcribe.html
func transcribe_StartMedicalTranscriptionJob(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.StartMedicalTranscriptionJobInput{
		// LanguageCode: types.LanguageCode, // Required
		// Media: *types.Media, // Required
		// MedicalTranscriptionJobName: *string, // Required
		// OutputBucketName: *string, // Required
		// Specialty: types.Specialty, // Required
		// Type: types.Type, // Required
	}

	if len(_transcribeLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _transcribeLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_transcribeMedia) > 0 {
		if err := assignInputField(input, "Media", _transcribeMedia); err != nil {
			log.Errorf("invalid --media: %s", err.Error())
			return
		}
	}
	if len(_transcribeMedicalTranscriptionJobName) > 0 {
		input.MedicalTranscriptionJobName = aws.String(_transcribeMedicalTranscriptionJobName)
	}
	if len(_transcribeOutputBucketName) > 0 {
		input.OutputBucketName = aws.String(_transcribeOutputBucketName)
	}
	if len(_transcribeSpecialty) > 0 {
		if err := assignInputField(input, "Specialty", _transcribeSpecialty); err != nil {
			log.Errorf("invalid --specialty: %s", err.Error())
			return
		}
	}
	if len(_transcribeType) > 0 {
		if err := assignInputField(input, "Type", _transcribeType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_transcribeContentIdentificationType) > 0 {
		if err := assignInputField(input, "ContentIdentificationType", _transcribeContentIdentificationType); err != nil {
			log.Errorf("invalid --content-identification-type: %s", err.Error())
			return
		}
	}
	if len(_transcribeKMSEncryptionContext) > 0 {
		if err := assignInputField(input, "KMSEncryptionContext", _transcribeKMSEncryptionContext); err != nil {
			log.Errorf("invalid --kms-encryption-context: %s", err.Error())
			return
		}
	}
	if len(_transcribeMediaFormat) > 0 {
		if err := assignInputField(input, "MediaFormat", _transcribeMediaFormat); err != nil {
			log.Errorf("invalid --media-format: %s", err.Error())
			return
		}
	}
	if len(_transcribeMediaSampleRateHertz) > 0 {
		if err := assignInputField(input, "MediaSampleRateHertz", _transcribeMediaSampleRateHertz); err != nil {
			log.Errorf("invalid --media-sample-rate-hertz: %s", err.Error())
			return
		}
	}
	if len(_transcribeOutputEncryptionKMSKeyId) > 0 {
		input.OutputEncryptionKMSKeyId = aws.String(_transcribeOutputEncryptionKMSKeyId)
	}
	if len(_transcribeOutputKey) > 0 {
		input.OutputKey = aws.String(_transcribeOutputKey)
	}
	if len(_transcribeSettings) > 0 {
		if err := assignInputField(input, "Settings", _transcribeSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}
	if len(_transcribeTags) > 0 {
		if err := assignInputField(input, "Tags", _transcribeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartMedicalTranscriptionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Transcribes the audio from a media file and applies any additional Request
// Parameters you choose to include in your request.
//
// To make a StartTranscriptionJob request, you must first upload your media file
// into an Amazon S3 bucket; you can then specify the Amazon S3 location of the
// file using the Media parameter.
//
// You must include the following parameters in your StartTranscriptionJob request:
//
// - region : The Amazon Web Services Region where you are making your request.
// For a list of Amazon Web Services Regions supported with Amazon Transcribe,
// refer to [Amazon Transcribe endpoints and quotas].
//
// - TranscriptionJobName : A custom name you create for your transcription job
// that is unique within your Amazon Web Services account.
//
// - Media ( MediaFileUri ): The Amazon S3 location of your media file.
//
// - One of LanguageCode , IdentifyLanguage , or IdentifyMultipleLanguages : If
// you know the language of your media file, specify it using the LanguageCode
// parameter; you can find all valid language codes in the [Supported languages]table. If you do not
// know the languages spoken in your media, use either IdentifyLanguage or
// IdentifyMultipleLanguages and let Amazon Transcribe identify the languages for
// you.
//
// [Amazon Transcribe endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/transcribe.html
// [Supported languages]: https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html
func transcribe_StartTranscriptionJob(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.StartTranscriptionJobInput{
		// Media: *types.Media, // Required
		// TranscriptionJobName: *string, // Required
	}

	if len(_transcribeMedia) > 0 {
		if err := assignInputField(input, "Media", _transcribeMedia); err != nil {
			log.Errorf("invalid --media: %s", err.Error())
			return
		}
	}
	if len(_transcribeTranscriptionJobName) > 0 {
		input.TranscriptionJobName = aws.String(_transcribeTranscriptionJobName)
	}
	if len(_transcribeContentRedaction) > 0 {
		if err := assignInputField(input, "ContentRedaction", _transcribeContentRedaction); err != nil {
			log.Errorf("invalid --content-redaction: %s", err.Error())
			return
		}
	}
	if len(_transcribeIdentifyLanguage) > 0 {
		if err := assignInputField(input, "IdentifyLanguage", _transcribeIdentifyLanguage); err != nil {
			log.Errorf("invalid --identify-language: %s", err.Error())
			return
		}
	}
	if len(_transcribeIdentifyMultipleLanguages) > 0 {
		if err := assignInputField(input, "IdentifyMultipleLanguages", _transcribeIdentifyMultipleLanguages); err != nil {
			log.Errorf("invalid --identify-multiple-languages: %s", err.Error())
			return
		}
	}
	if len(_transcribeJobExecutionSettings) > 0 {
		if err := assignInputField(input, "JobExecutionSettings", _transcribeJobExecutionSettings); err != nil {
			log.Errorf("invalid --job-execution-settings: %s", err.Error())
			return
		}
	}
	if len(_transcribeKMSEncryptionContext) > 0 {
		if err := assignInputField(input, "KMSEncryptionContext", _transcribeKMSEncryptionContext); err != nil {
			log.Errorf("invalid --kms-encryption-context: %s", err.Error())
			return
		}
	}
	if len(_transcribeLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _transcribeLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_transcribeLanguageIdSettings) > 0 {
		if err := assignInputField(input, "LanguageIdSettings", _transcribeLanguageIdSettings); err != nil {
			log.Errorf("invalid --language-id-settings: %s", err.Error())
			return
		}
	}
	if len(_transcribeLanguageOptions) > 0 {
		if err := assignInputField(input, "LanguageOptions", _transcribeLanguageOptions); err != nil {
			log.Errorf("invalid --language-options: %s", err.Error())
			return
		}
	}
	if len(_transcribeMediaFormat) > 0 {
		if err := assignInputField(input, "MediaFormat", _transcribeMediaFormat); err != nil {
			log.Errorf("invalid --media-format: %s", err.Error())
			return
		}
	}
	if len(_transcribeMediaSampleRateHertz) > 0 {
		if err := assignInputField(input, "MediaSampleRateHertz", _transcribeMediaSampleRateHertz); err != nil {
			log.Errorf("invalid --media-sample-rate-hertz: %s", err.Error())
			return
		}
	}
	if len(_transcribeModelSettings) > 0 {
		if err := assignInputField(input, "ModelSettings", _transcribeModelSettings); err != nil {
			log.Errorf("invalid --model-settings: %s", err.Error())
			return
		}
	}
	if len(_transcribeOutputBucketName) > 0 {
		input.OutputBucketName = aws.String(_transcribeOutputBucketName)
	}
	if len(_transcribeOutputEncryptionKMSKeyId) > 0 {
		input.OutputEncryptionKMSKeyId = aws.String(_transcribeOutputEncryptionKMSKeyId)
	}
	if len(_transcribeOutputKey) > 0 {
		input.OutputKey = aws.String(_transcribeOutputKey)
	}
	if len(_transcribeSettings) > 0 {
		if err := assignInputField(input, "Settings", _transcribeSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}
	if len(_transcribeSubtitles) > 0 {
		if err := assignInputField(input, "Subtitles", _transcribeSubtitles); err != nil {
			log.Errorf("invalid --subtitles: %s", err.Error())
			return
		}
	}
	if len(_transcribeTags) > 0 {
		if err := assignInputField(input, "Tags", _transcribeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_transcribeToxicityDetection) > 0 {
		if err := assignInputField(input, "ToxicityDetection", _transcribeToxicityDetection); err != nil {
			log.Errorf("invalid --toxicity-detection: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartTranscriptionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more custom tags, each in the form of a key:value pair, to the
// specified resource.
//
// To learn more about using tags with Amazon Transcribe, refer to [Tagging resources].
//
// [Tagging resources]: https://docs.aws.amazon.com/transcribe/latest/dg/tagging.html
func transcribe_TagResource(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_transcribeResourceArn) > 0 {
		input.ResourceArn = aws.String(_transcribeResourceArn)
	}
	if len(_transcribeTags) > 0 {
		if err := assignInputField(input, "Tags", _transcribeTags); err != nil {
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

// Removes the specified tags from the specified Amazon Transcribe resource.
// If you include UntagResource in your request, you must also include ResourceArn
// and TagKeys .
func transcribe_UntagResource(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_transcribeResourceArn) > 0 {
		input.ResourceArn = aws.String(_transcribeResourceArn)
	}
	if len(_transcribeTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _transcribeTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified Call Analytics category with new rules. Note that the
// UpdateCallAnalyticsCategory operation overwrites all existing rules contained in
// the specified category. You cannot append additional rules onto an existing
// category.
//
// To create a new category, see .
func transcribe_UpdateCallAnalyticsCategory(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.UpdateCallAnalyticsCategoryInput{
		// CategoryName: *string, // Required
		// Rules: []types.Rule, // Required
	}

	if len(_transcribeCategoryName) > 0 {
		input.CategoryName = aws.String(_transcribeCategoryName)
	}
	if len(_transcribeRules) > 0 {
		if err := assignInputField(input, "Rules", _transcribeRules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_transcribeInputType) > 0 {
		if err := assignInputField(input, "InputType", _transcribeInputType); err != nil {
			log.Errorf("invalid --input-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCallAnalyticsCategory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing custom medical vocabulary with new values. This operation
// overwrites all existing information with your new values; you cannot append new
// terms onto an existing custom vocabulary.
func transcribe_UpdateMedicalVocabulary(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.UpdateMedicalVocabularyInput{
		// LanguageCode: types.LanguageCode, // Required
		// VocabularyFileUri: *string, // Required
		// VocabularyName: *string, // Required
	}

	if len(_transcribeLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _transcribeLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_transcribeVocabularyFileUri) > 0 {
		input.VocabularyFileUri = aws.String(_transcribeVocabularyFileUri)
	}
	if len(_transcribeVocabularyName) > 0 {
		input.VocabularyName = aws.String(_transcribeVocabularyName)
	}

	if resp, err := client.UpdateMedicalVocabulary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing custom vocabulary with new values. This operation
// overwrites all existing information with your new values; you cannot append new
// terms onto an existing custom vocabulary.
func transcribe_UpdateVocabulary(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.UpdateVocabularyInput{
		// LanguageCode: types.LanguageCode, // Required
		// VocabularyName: *string, // Required
	}

	if len(_transcribeLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _transcribeLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_transcribeVocabularyName) > 0 {
		input.VocabularyName = aws.String(_transcribeVocabularyName)
	}
	if len(_transcribeDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_transcribeDataAccessRoleArn)
	}
	if len(_transcribePhrases) > 0 {
		input.Phrases = append([]string(nil), _transcribePhrases...)
	}
	if len(_transcribeVocabularyFileUri) > 0 {
		input.VocabularyFileUri = aws.String(_transcribeVocabularyFileUri)
	}

	if resp, err := client.UpdateVocabulary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing custom vocabulary filter with a new list of words. The new
// list you provide overwrites all previous entries; you cannot append new terms
// onto an existing custom vocabulary filter.
func transcribe_UpdateVocabularyFilter(cfg aws.Config, client *transcribe.Client) {
	input := &transcribe.UpdateVocabularyFilterInput{
		// VocabularyFilterName: *string, // Required
	}

	if len(_transcribeVocabularyFilterName) > 0 {
		input.VocabularyFilterName = aws.String(_transcribeVocabularyFilterName)
	}
	if len(_transcribeDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_transcribeDataAccessRoleArn)
	}
	if len(_transcribeVocabularyFilterFileUri) > 0 {
		input.VocabularyFilterFileUri = aws.String(_transcribeVocabularyFilterFileUri)
	}
	if len(_transcribeWords) > 0 {
		input.Words = append([]string(nil), _transcribeWords...)
	}

	if resp, err := client.UpdateVocabularyFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_transcribeCmd)
	_transcribeCmd.Flags().SortFlags = false

	_transcribeCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_transcribeCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_transcribeCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_transcribeCmd.Flags().StringVarP(&_transcribeBaseModelName, "base-model-name", "", "", "Base Model Name")
	_transcribeCmd.Flags().StringVarP(&_transcribeCallAnalyticsJobName, "call-analytics-job-name", "", "", "Call Analytics Job Name")
	_transcribeCmd.Flags().StringVarP(&_transcribeCategoryName, "category-name", "", "", "Category Name")
	_transcribeCmd.Flags().StringVarP(&_transcribeChannelDefinitions, "channel-definitions", "", "", "Channel Definitions")
	_transcribeCmd.Flags().StringVarP(&_transcribeContentIdentificationType, "content-identification-type", "", "", "Content Identification Type")
	_transcribeCmd.Flags().StringVarP(&_transcribeContentRedaction, "content-redaction", "", "", "Content Redaction")
	_transcribeCmd.Flags().StringVarP(&_transcribeDataAccessRoleArn, "data-access-role-arn", "", "", "Data Access Role ARN")
	_transcribeCmd.Flags().StringVarP(&_transcribeIdentifyLanguage, "identify-language", "", "", "Identify Language")
	_transcribeCmd.Flags().StringVarP(&_transcribeIdentifyMultipleLanguages, "identify-multiple-languages", "", "", "Identify Multiple Languages")
	_transcribeCmd.Flags().StringVarP(&_transcribeInputDataConfig, "input-data-config", "", "", "Input Data Config")
	_transcribeCmd.Flags().StringVarP(&_transcribeInputType, "input-type", "", "", "Input Type")
	_transcribeCmd.Flags().StringVarP(&_transcribeJobExecutionSettings, "job-execution-settings", "", "", "Job Execution Settings")
	_transcribeCmd.Flags().StringVarP(&_transcribeJobNameContains, "job-name-contains", "", "", "Job Name Contains")
	_transcribeCmd.Flags().StringVarP(&_transcribeKMSEncryptionContext, "kms-encryption-context", "", "", "KMS Encryption Context")
	_transcribeCmd.Flags().StringVarP(&_transcribeLanguageCode, "language-code", "", "", "Language Code")
	_transcribeCmd.Flags().StringVarP(&_transcribeLanguageIdSettings, "language-id-settings", "", "", "Language ID Settings")
	_transcribeCmd.Flags().StringVarP(&_transcribeLanguageOptions, "language-options", "", "", "Language Options")
	_transcribeCmd.Flags().StringVarP(&_transcribeMaxResults, "max-results", "", "", "Max Results")
	_transcribeCmd.Flags().StringVarP(&_transcribeMedia, "media", "", "", "Media")
	_transcribeCmd.Flags().StringVarP(&_transcribeMediaFormat, "media-format", "", "", "Media Format")
	_transcribeCmd.Flags().StringVarP(&_transcribeMediaSampleRateHertz, "media-sample-rate-hertz", "", "", "Media Sample Rate Hertz")
	_transcribeCmd.Flags().StringVarP(&_transcribeMedicalScribeContext, "medical-scribe-context", "", "", "Medical Scribe Context")
	_transcribeCmd.Flags().StringVarP(&_transcribeMedicalScribeJobName, "medical-scribe-job-name", "", "", "Medical Scribe Job Name")
	_transcribeCmd.Flags().StringVarP(&_transcribeMedicalTranscriptionJobName, "medical-transcription-job-name", "", "", "Medical Transcription Job Name")
	_transcribeCmd.Flags().StringVarP(&_transcribeModelName, "model-name", "", "", "Model Name")
	_transcribeCmd.Flags().StringVarP(&_transcribeModelSettings, "model-settings", "", "", "Model Settings")
	_transcribeCmd.Flags().StringVarP(&_transcribeNameContains, "name-contains", "", "", "Name Contains")
	_transcribeCmd.Flags().StringVarP(&_transcribeNextToken, "next-token", "", "", "Next Token")
	_transcribeCmd.Flags().StringVarP(&_transcribeOutputBucketName, "output-bucket-name", "", "", "Output Bucket Name")
	_transcribeCmd.Flags().StringVarP(&_transcribeOutputEncryptionKMSKeyId, "output-encryption-kms-key-id", "", "", "Output Encryption KMS Key ID")
	_transcribeCmd.Flags().StringVarP(&_transcribeOutputKey, "output-key", "", "", "Output Key")
	_transcribeCmd.Flags().StringVarP(&_transcribeOutputLocation, "output-location", "", "", "Output Location")
	_transcribeCmd.Flags().StringSliceVarP(&_transcribePhrases, "phrases", "", nil, "Phrases")
	_transcribeCmd.Flags().StringVarP(&_transcribeResourceArn, "resource-arn", "", "", "Resource ARN")
	_transcribeCmd.Flags().StringVarP(&_transcribeRules, "rules", "", "", "Rules")
	_transcribeCmd.Flags().StringVarP(&_transcribeSettings, "settings", "", "", "Settings")
	_transcribeCmd.Flags().StringVarP(&_transcribeSpecialty, "specialty", "", "", "Specialty")
	_transcribeCmd.Flags().StringVarP(&_transcribeStateEquals, "state-equals", "", "", "State Equals")
	_transcribeCmd.Flags().StringVarP(&_transcribeStatus, "status", "", "", "Status")
	_transcribeCmd.Flags().StringVarP(&_transcribeStatusEquals, "status-equals", "", "", "Status Equals")
	_transcribeCmd.Flags().StringVarP(&_transcribeSubtitles, "subtitles", "", "", "Subtitles")
	_transcribeCmd.Flags().StringSliceVarP(&_transcribeTagKeys, "tag-keys", "", nil, "Tag Keys")
	_transcribeCmd.Flags().StringVarP(&_transcribeTags, "tags", "", "", "Tags")
	_transcribeCmd.Flags().StringVarP(&_transcribeToxicityDetection, "toxicity-detection", "", "", "Toxicity Detection")
	_transcribeCmd.Flags().StringVarP(&_transcribeTranscriptionJobName, "transcription-job-name", "", "", "Transcription Job Name")
	_transcribeCmd.Flags().StringVarP(&_transcribeType, "type", "", "", "Type")
	_transcribeCmd.Flags().StringVarP(&_transcribeVocabularyFileUri, "vocabulary-file-uri", "", "", "Vocabulary File URI")
	_transcribeCmd.Flags().StringVarP(&_transcribeVocabularyFilterFileUri, "vocabulary-filter-file-uri", "", "", "Vocabulary Filter File URI")
	_transcribeCmd.Flags().StringVarP(&_transcribeVocabularyFilterName, "vocabulary-filter-name", "", "", "Vocabulary Filter Name")
	_transcribeCmd.Flags().StringVarP(&_transcribeVocabularyName, "vocabulary-name", "", "", "Vocabulary Name")
	_transcribeCmd.Flags().StringSliceVarP(&_transcribeWords, "words", "", nil, "Words")

	_transcribeCmd.Flags().BoolVarP(&_transcribeCreateCallAnalyticsCategory, "create-call-analytics-category", "", false, "Create Call Analytics Category")
	_transcribeCmd.Flags().BoolVarP(&_transcribeCreateLanguageModel, "create-language-model", "", false, "Create Language Model")
	_transcribeCmd.Flags().BoolVarP(&_transcribeCreateMedicalVocabulary, "create-medical-vocabulary", "", false, "Create Medical Vocabulary")
	_transcribeCmd.Flags().BoolVarP(&_transcribeCreateVocabulary, "create-vocabulary", "", false, "Create Vocabulary")
	_transcribeCmd.Flags().BoolVarP(&_transcribeCreateVocabularyFilter, "create-vocabulary-filter", "", false, "Create Vocabulary Filter")
	_transcribeCmd.Flags().BoolVarP(&_transcribeDeleteCallAnalyticsCategory, "delete-call-analytics-category", "", false, "Delete Call Analytics Category")
	_transcribeCmd.Flags().BoolVarP(&_transcribeDeleteCallAnalyticsJob, "delete-call-analytics-job", "", false, "Delete Call Analytics Job")
	_transcribeCmd.Flags().BoolVarP(&_transcribeDeleteLanguageModel, "delete-language-model", "", false, "Delete Language Model")
	_transcribeCmd.Flags().BoolVarP(&_transcribeDeleteMedicalScribeJob, "delete-medical-scribe-job", "", false, "Delete Medical Scribe Job")
	_transcribeCmd.Flags().BoolVarP(&_transcribeDeleteMedicalTranscriptionJob, "delete-medical-transcription-job", "", false, "Delete Medical Transcription Job")
	_transcribeCmd.Flags().BoolVarP(&_transcribeDeleteMedicalVocabulary, "delete-medical-vocabulary", "", false, "Delete Medical Vocabulary")
	_transcribeCmd.Flags().BoolVarP(&_transcribeDeleteTranscriptionJob, "delete-transcription-job", "", false, "Delete Transcription Job")
	_transcribeCmd.Flags().BoolVarP(&_transcribeDeleteVocabulary, "delete-vocabulary", "", false, "Delete Vocabulary")
	_transcribeCmd.Flags().BoolVarP(&_transcribeDeleteVocabularyFilter, "delete-vocabulary-filter", "", false, "Delete Vocabulary Filter")
	_transcribeCmd.Flags().BoolVarP(&_transcribeDescribeLanguageModel, "describe-language-model", "", false, "Describe Language Model")
	_transcribeCmd.Flags().BoolVarP(&_transcribeGetCallAnalyticsCategory, "get-call-analytics-category", "", false, "Get Call Analytics Category")
	_transcribeCmd.Flags().BoolVarP(&_transcribeGetCallAnalyticsJob, "get-call-analytics-job", "", false, "Get Call Analytics Job")
	_transcribeCmd.Flags().BoolVarP(&_transcribeGetMedicalScribeJob, "get-medical-scribe-job", "", false, "Get Medical Scribe Job")
	_transcribeCmd.Flags().BoolVarP(&_transcribeGetMedicalTranscriptionJob, "get-medical-transcription-job", "", false, "Get Medical Transcription Job")
	_transcribeCmd.Flags().BoolVarP(&_transcribeGetMedicalVocabulary, "get-medical-vocabulary", "", false, "Get Medical Vocabulary")
	_transcribeCmd.Flags().BoolVarP(&_transcribeGetTranscriptionJob, "get-transcription-job", "", false, "Get Transcription Job")
	_transcribeCmd.Flags().BoolVarP(&_transcribeGetVocabulary, "get-vocabulary", "", false, "Get Vocabulary")
	_transcribeCmd.Flags().BoolVarP(&_transcribeGetVocabularyFilter, "get-vocabulary-filter", "", false, "Get Vocabulary Filter")
	_transcribeCmd.Flags().BoolVarP(&_transcribeListCallAnalyticsCategories, "list-call-analytics-categories", "", false, "List Call Analytics Categories")
	_transcribeCmd.Flags().BoolVarP(&_transcribeListCallAnalyticsJobs, "list-call-analytics-jobs", "", false, "List Call Analytics Jobs")
	_transcribeCmd.Flags().BoolVarP(&_transcribeListLanguageModels, "list-language-models", "", false, "List Language Models")
	_transcribeCmd.Flags().BoolVarP(&_transcribeListMedicalScribeJobs, "list-medical-scribe-jobs", "", false, "List Medical Scribe Jobs")
	_transcribeCmd.Flags().BoolVarP(&_transcribeListMedicalTranscriptionJobs, "list-medical-transcription-jobs", "", false, "List Medical Transcription Jobs")
	_transcribeCmd.Flags().BoolVarP(&_transcribeListMedicalVocabularies, "list-medical-vocabularies", "", false, "List Medical Vocabularies")
	_transcribeCmd.Flags().BoolVarP(&_transcribeListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_transcribeCmd.Flags().BoolVarP(&_transcribeListTranscriptionJobs, "list-transcription-jobs", "", false, "List Transcription Jobs")
	_transcribeCmd.Flags().BoolVarP(&_transcribeListVocabularies, "list-vocabularies", "", false, "List Vocabularies")
	_transcribeCmd.Flags().BoolVarP(&_transcribeListVocabularyFilters, "list-vocabulary-filters", "", false, "List Vocabulary Filters")
	_transcribeCmd.Flags().BoolVarP(&_transcribeStartCallAnalyticsJob, "start-call-analytics-job", "", false, "Start Call Analytics Job")
	_transcribeCmd.Flags().BoolVarP(&_transcribeStartMedicalScribeJob, "start-medical-scribe-job", "", false, "Start Medical Scribe Job")
	_transcribeCmd.Flags().BoolVarP(&_transcribeStartMedicalTranscriptionJob, "start-medical-transcription-job", "", false, "Start Medical Transcription Job")
	_transcribeCmd.Flags().BoolVarP(&_transcribeStartTranscriptionJob, "start-transcription-job", "", false, "Start Transcription Job")
	_transcribeCmd.Flags().BoolVarP(&_transcribeTagResource, "tag-resource", "", false, "Tag Resource")
	_transcribeCmd.Flags().BoolVarP(&_transcribeUntagResource, "untag-resource", "", false, "Untag Resource")
	_transcribeCmd.Flags().BoolVarP(&_transcribeUpdateCallAnalyticsCategory, "update-call-analytics-category", "", false, "Update Call Analytics Category")
	_transcribeCmd.Flags().BoolVarP(&_transcribeUpdateMedicalVocabulary, "update-medical-vocabulary", "", false, "Update Medical Vocabulary")
	_transcribeCmd.Flags().BoolVarP(&_transcribeUpdateVocabulary, "update-vocabulary", "", false, "Update Vocabulary")
	_transcribeCmd.Flags().BoolVarP(&_transcribeUpdateVocabularyFilter, "update-vocabulary-filter", "", false, "Update Vocabulary Filter")

}
