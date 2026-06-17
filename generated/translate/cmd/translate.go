package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/translate"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// translateCmd represents the translate command
var _translateCmd = &cobra.Command{
	Use:   "translate",
	Short: "AWS translate CLI",
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
		client := translate.NewFromConfig(cfg)
		if _translateCreateParallelData {
			translate_CreateParallelData(cfg, client)
			return
		}
		if _translateDeleteParallelData {
			translate_DeleteParallelData(cfg, client)
			return
		}
		if _translateDeleteTerminology {
			translate_DeleteTerminology(cfg, client)
			return
		}
		if _translateDescribeTextTranslationJob {
			translate_DescribeTextTranslationJob(cfg, client)
			return
		}
		if _translateGetParallelData {
			translate_GetParallelData(cfg, client)
			return
		}
		if _translateGetTerminology {
			translate_GetTerminology(cfg, client)
			return
		}
		if _translateImportTerminology {
			translate_ImportTerminology(cfg, client)
			return
		}
		if _translateListLanguages {
			translate_ListLanguages(cfg, client)
			return
		}
		if _translateListParallelData {
			translate_ListParallelData(cfg, client)
			return
		}
		if _translateListTagsForResource {
			translate_ListTagsForResource(cfg, client)
			return
		}
		if _translateListTerminologies {
			translate_ListTerminologies(cfg, client)
			return
		}
		if _translateListTextTranslationJobs {
			translate_ListTextTranslationJobs(cfg, client)
			return
		}
		if _translateStartTextTranslationJob {
			translate_StartTextTranslationJob(cfg, client)
			return
		}
		if _translateStopTextTranslationJob {
			translate_StopTextTranslationJob(cfg, client)
			return
		}
		if _translateTagResource {
			translate_TagResource(cfg, client)
			return
		}
		if _translateTranslateDocument {
			translate_TranslateDocument(cfg, client)
			return
		}
		if _translateTranslateText {
			translate_TranslateText(cfg, client)
			return
		}
		if _translateUntagResource {
			translate_UntagResource(cfg, client)
			return
		}
		if _translateUpdateParallelData {
			translate_UpdateParallelData(cfg, client)
			return
		}

	},
}

var (
	_translateCreateParallelData         bool
	_translateDeleteParallelData         bool
	_translateDeleteTerminology          bool
	_translateDescribeTextTranslationJob bool
	_translateGetParallelData            bool
	_translateGetTerminology             bool
	_translateImportTerminology          bool
	_translateListLanguages              bool
	_translateListParallelData           bool
	_translateListTagsForResource        bool
	_translateListTerminologies          bool
	_translateListTextTranslationJobs    bool
	_translateStartTextTranslationJob    bool
	_translateStopTextTranslationJob     bool
	_translateTagResource                bool
	_translateTranslateDocument          bool
	_translateTranslateText              bool
	_translateUntagResource              bool
	_translateUpdateParallelData         bool

	_translateClientToken           string
	_translateDataAccessRoleArn     string
	_translateDescription           string
	_translateDisplayLanguageCode   string
	_translateDocument              string
	_translateEncryptionKey         string
	_translateFilter                string
	_translateInputDataConfig       string
	_translateJobId                 string
	_translateJobName               string
	_translateMaxResults            string
	_translateMergeStrategy         string
	_translateName                  string
	_translateNextToken             string
	_translateOutputDataConfig      string
	_translateParallelDataConfig    string
	_translateParallelDataNames     []string
	_translateResourceArn           string
	_translateSettings              string
	_translateSourceLanguageCode    string
	_translateTagKeys               []string
	_translateTags                  string
	_translateTargetLanguageCode    string
	_translateTargetLanguageCodes   []string
	_translateTerminologyData       string
	_translateTerminologyDataFormat string
	_translateTerminologyNames      []string
	_translateText                  string
)

// Creates a parallel data resource in Amazon Translate by importing an input file
// from Amazon S3. Parallel data files contain examples that show how you want
// segments of text to be translated. By adding parallel data, you can influence
// the style, tone, and word choice in your translation output.
func translate_CreateParallelData(cfg aws.Config, client *translate.Client) {
	input := &translate.CreateParallelDataInput{
		// ClientToken: *string, // Required
		// Name: *string, // Required
		// ParallelDataConfig: *types.ParallelDataConfig, // Required
	}

	if len(_translateClientToken) > 0 {
		input.ClientToken = aws.String(_translateClientToken)
	}
	if len(_translateName) > 0 {
		input.Name = aws.String(_translateName)
	}
	if len(_translateParallelDataConfig) > 0 {
		if err := assignInputField(input, "ParallelDataConfig", _translateParallelDataConfig); err != nil {
			log.Errorf("invalid --parallel-data-config: %s", err.Error())
			return
		}
	}
	if len(_translateDescription) > 0 {
		input.Description = aws.String(_translateDescription)
	}
	if len(_translateEncryptionKey) > 0 {
		if err := assignInputField(input, "EncryptionKey", _translateEncryptionKey); err != nil {
			log.Errorf("invalid --encryption-key: %s", err.Error())
			return
		}
	}
	if len(_translateTags) > 0 {
		if err := assignInputField(input, "Tags", _translateTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateParallelData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a parallel data resource in Amazon Translate.
func translate_DeleteParallelData(cfg aws.Config, client *translate.Client) {
	input := &translate.DeleteParallelDataInput{
		// Name: *string, // Required
	}

	if len(_translateName) > 0 {
		input.Name = aws.String(_translateName)
	}

	if resp, err := client.DeleteParallelData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A synchronous action that deletes a custom terminology.
func translate_DeleteTerminology(cfg aws.Config, client *translate.Client) {
	input := &translate.DeleteTerminologyInput{
		// Name: *string, // Required
	}

	if len(_translateName) > 0 {
		input.Name = aws.String(_translateName)
	}

	if resp, err := client.DeleteTerminology(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the properties associated with an asynchronous batch translation job
// including name, ID, status, source and target languages, input/output S3
// buckets, and so on.
func translate_DescribeTextTranslationJob(cfg aws.Config, client *translate.Client) {
	input := &translate.DescribeTextTranslationJobInput{
		// JobId: *string, // Required
	}

	if len(_translateJobId) > 0 {
		input.JobId = aws.String(_translateJobId)
	}

	if resp, err := client.DescribeTextTranslationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about a parallel data resource.
func translate_GetParallelData(cfg aws.Config, client *translate.Client) {
	input := &translate.GetParallelDataInput{
		// Name: *string, // Required
	}

	if len(_translateName) > 0 {
		input.Name = aws.String(_translateName)
	}

	if resp, err := client.GetParallelData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a custom terminology.
func translate_GetTerminology(cfg aws.Config, client *translate.Client) {
	input := &translate.GetTerminologyInput{
		// Name: *string, // Required
	}

	if len(_translateName) > 0 {
		input.Name = aws.String(_translateName)
	}
	if len(_translateTerminologyDataFormat) > 0 {
		if err := assignInputField(input, "TerminologyDataFormat", _translateTerminologyDataFormat); err != nil {
			log.Errorf("invalid --terminology-data-format: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetTerminology(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a custom terminology, depending on whether one already
// exists for the given terminology name. Importing a terminology with the same
// name as an existing one will merge the terminologies based on the chosen merge
// strategy. The only supported merge strategy is OVERWRITE, where the imported
// terminology overwrites the existing terminology of the same name.
//
// If you import a terminology that overwrites an existing one, the new
// terminology takes up to 10 minutes to fully propagate. After that, translations
// have access to the new terminology.
func translate_ImportTerminology(cfg aws.Config, client *translate.Client) {
	input := &translate.ImportTerminologyInput{
		// MergeStrategy: types.MergeStrategy, // Required
		// Name: *string, // Required
		// TerminologyData: *types.TerminologyData, // Required
	}

	if len(_translateMergeStrategy) > 0 {
		if err := assignInputField(input, "MergeStrategy", _translateMergeStrategy); err != nil {
			log.Errorf("invalid --merge-strategy: %s", err.Error())
			return
		}
	}
	if len(_translateName) > 0 {
		input.Name = aws.String(_translateName)
	}
	if len(_translateTerminologyData) > 0 {
		if err := assignInputField(input, "TerminologyData", _translateTerminologyData); err != nil {
			log.Errorf("invalid --terminology-data: %s", err.Error())
			return
		}
	}
	if len(_translateDescription) > 0 {
		input.Description = aws.String(_translateDescription)
	}
	if len(_translateEncryptionKey) > 0 {
		if err := assignInputField(input, "EncryptionKey", _translateEncryptionKey); err != nil {
			log.Errorf("invalid --encryption-key: %s", err.Error())
			return
		}
	}
	if len(_translateTags) > 0 {
		if err := assignInputField(input, "Tags", _translateTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportTerminology(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a list of languages (RFC-5646 codes and names) that Amazon Translate
// supports.
func translate_ListLanguages(cfg aws.Config, client *translate.Client) {
	input := &translate.ListLanguagesInput{}

	if len(_translateDisplayLanguageCode) > 0 {
		if err := assignInputField(input, "DisplayLanguageCode", _translateDisplayLanguageCode); err != nil {
			log.Errorf("invalid --display-language-code: %s", err.Error())
			return
		}
	}
	if len(_translateMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _translateMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_translateNextToken) > 0 {
		input.NextToken = aws.String(_translateNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLanguages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*translate.ListLanguagesOutput
	p := translate.NewListLanguagesPaginator(client, input)
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

// Provides a list of your parallel data resources in Amazon Translate.
func translate_ListParallelData(cfg aws.Config, client *translate.Client) {
	input := &translate.ListParallelDataInput{}

	if len(_translateMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _translateMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_translateNextToken) > 0 {
		input.NextToken = aws.String(_translateNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListParallelData(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*translate.ListParallelDataOutput
	p := translate.NewListParallelDataPaginator(client, input)
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

// Lists all tags associated with a given Amazon Translate resource. For more
// information, see [Tagging your resources].
//
// [Tagging your resources]: https://docs.aws.amazon.com/translate/latest/dg/tagging.html
func translate_ListTagsForResource(cfg aws.Config, client *translate.Client) {
	input := &translate.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_translateResourceArn) > 0 {
		input.ResourceArn = aws.String(_translateResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a list of custom terminologies associated with your account.
func translate_ListTerminologies(cfg aws.Config, client *translate.Client) {
	input := &translate.ListTerminologiesInput{}

	if len(_translateMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _translateMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_translateNextToken) > 0 {
		input.NextToken = aws.String(_translateNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTerminologies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*translate.ListTerminologiesOutput
	p := translate.NewListTerminologiesPaginator(client, input)
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

// Gets a list of the batch translation jobs that you have submitted.
func translate_ListTextTranslationJobs(cfg aws.Config, client *translate.Client) {
	input := &translate.ListTextTranslationJobsInput{}

	if len(_translateFilter) > 0 {
		if err := assignInputField(input, "Filter", _translateFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_translateMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _translateMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_translateNextToken) > 0 {
		input.NextToken = aws.String(_translateNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTextTranslationJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*translate.ListTextTranslationJobsOutput
	p := translate.NewListTextTranslationJobsPaginator(client, input)
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

// Starts an asynchronous batch translation job. Use batch translation jobs to
// translate large volumes of text across multiple documents at once. For batch
// translation, you can input documents with different source languages (specify
// auto as the source language). You can specify one or more target languages.
// Batch translation translates each input document into each of the target
// languages. For more information, see [Asynchronous batch processing].
//
// Batch translation jobs can be described with the DescribeTextTranslationJob operation, listed with the ListTextTranslationJobs
// operation, and stopped with the StopTextTranslationJoboperation.
//
// [Asynchronous batch processing]: https://docs.aws.amazon.com/translate/latest/dg/async.html
func translate_StartTextTranslationJob(cfg aws.Config, client *translate.Client) {
	input := &translate.StartTextTranslationJobInput{
		// ClientToken: *string, // Required
		// DataAccessRoleArn: *string, // Required
		// InputDataConfig: *types.InputDataConfig, // Required
		// OutputDataConfig: *types.OutputDataConfig, // Required
		// SourceLanguageCode: *string, // Required
		// TargetLanguageCodes: []string, // Required
	}

	if len(_translateClientToken) > 0 {
		input.ClientToken = aws.String(_translateClientToken)
	}
	if len(_translateDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_translateDataAccessRoleArn)
	}
	if len(_translateInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _translateInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_translateOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _translateOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_translateSourceLanguageCode) > 0 {
		input.SourceLanguageCode = aws.String(_translateSourceLanguageCode)
	}
	if len(_translateTargetLanguageCodes) > 0 {
		input.TargetLanguageCodes = append([]string(nil), _translateTargetLanguageCodes...)
	}
	if len(_translateJobName) > 0 {
		input.JobName = aws.String(_translateJobName)
	}
	if len(_translateParallelDataNames) > 0 {
		input.ParallelDataNames = append([]string(nil), _translateParallelDataNames...)
	}
	if len(_translateSettings) > 0 {
		if err := assignInputField(input, "Settings", _translateSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}
	if len(_translateTerminologyNames) > 0 {
		input.TerminologyNames = append([]string(nil), _translateTerminologyNames...)
	}

	if resp, err := client.StartTextTranslationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an asynchronous batch translation job that is in progress.
// If the job's state is IN_PROGRESS , the job will be marked for termination and
// put into the STOP_REQUESTED state. If the job completes before it can be
// stopped, it is put into the COMPLETED state. Otherwise, the job is put into the
// STOPPED state.
//
// Asynchronous batch translation jobs are started with the StartTextTranslationJob operation. You can
// use the DescribeTextTranslationJobor ListTextTranslationJobs operations to get a batch translation job's JobId .
func translate_StopTextTranslationJob(cfg aws.Config, client *translate.Client) {
	input := &translate.StopTextTranslationJobInput{
		// JobId: *string, // Required
	}

	if len(_translateJobId) > 0 {
		input.JobId = aws.String(_translateJobId)
	}

	if resp, err := client.StopTextTranslationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a specific tag with a resource. A tag is a key-value pair that adds
// as a metadata to a resource. For more information, see [Tagging your resources].
//
// [Tagging your resources]: https://docs.aws.amazon.com/translate/latest/dg/tagging.html
func translate_TagResource(cfg aws.Config, client *translate.Client) {
	input := &translate.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_translateResourceArn) > 0 {
		input.ResourceArn = aws.String(_translateResourceArn)
	}
	if len(_translateTags) > 0 {
		if err := assignInputField(input, "Tags", _translateTags); err != nil {
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

// Translates the input document from the source language to the target language.
// This synchronous operation supports text, HTML, or Word documents as the input
// document.
//
// TranslateDocument supports translations from English to any supported language,
// and from any supported language to English. Therefore, specify either the source
// language code or the target language code as “en” (English).
//
// If you set the Formality parameter, the request will fail if the target
// language does not support formality. For a list of target languages that support
// formality, see [Setting formality].
//
// [Setting formality]: https://docs.aws.amazon.com/translate/latest/dg/customizing-translations-formality.html
func translate_TranslateDocument(cfg aws.Config, client *translate.Client) {
	input := &translate.TranslateDocumentInput{
		// Document: *types.Document, // Required
		// SourceLanguageCode: *string, // Required
		// TargetLanguageCode: *string, // Required
	}

	if len(_translateDocument) > 0 {
		if err := assignInputField(input, "Document", _translateDocument); err != nil {
			log.Errorf("invalid --document: %s", err.Error())
			return
		}
	}
	if len(_translateSourceLanguageCode) > 0 {
		input.SourceLanguageCode = aws.String(_translateSourceLanguageCode)
	}
	if len(_translateTargetLanguageCode) > 0 {
		input.TargetLanguageCode = aws.String(_translateTargetLanguageCode)
	}
	if len(_translateSettings) > 0 {
		if err := assignInputField(input, "Settings", _translateSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}
	if len(_translateTerminologyNames) > 0 {
		input.TerminologyNames = append([]string(nil), _translateTerminologyNames...)
	}

	if resp, err := client.TranslateDocument(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Translates input text from the source language to the target language. For a
// list of available languages and language codes, see [Supported languages].
//
// [Supported languages]: https://docs.aws.amazon.com/translate/latest/dg/what-is-languages.html
func translate_TranslateText(cfg aws.Config, client *translate.Client) {
	input := &translate.TranslateTextInput{
		// SourceLanguageCode: *string, // Required
		// TargetLanguageCode: *string, // Required
		// Text: *string, // Required
	}

	if len(_translateSourceLanguageCode) > 0 {
		input.SourceLanguageCode = aws.String(_translateSourceLanguageCode)
	}
	if len(_translateTargetLanguageCode) > 0 {
		input.TargetLanguageCode = aws.String(_translateTargetLanguageCode)
	}
	if len(_translateText) > 0 {
		input.Text = aws.String(_translateText)
	}
	if len(_translateSettings) > 0 {
		if err := assignInputField(input, "Settings", _translateSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}
	if len(_translateTerminologyNames) > 0 {
		input.TerminologyNames = append([]string(nil), _translateTerminologyNames...)
	}

	if resp, err := client.TranslateText(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a specific tag associated with an Amazon Translate resource. For more
// information, see [Tagging your resources].
//
// [Tagging your resources]: https://docs.aws.amazon.com/translate/latest/dg/tagging.html
func translate_UntagResource(cfg aws.Config, client *translate.Client) {
	input := &translate.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_translateResourceArn) > 0 {
		input.ResourceArn = aws.String(_translateResourceArn)
	}
	if len(_translateTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _translateTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a previously created parallel data resource by importing a new input
// file from Amazon S3.
func translate_UpdateParallelData(cfg aws.Config, client *translate.Client) {
	input := &translate.UpdateParallelDataInput{
		// ClientToken: *string, // Required
		// Name: *string, // Required
		// ParallelDataConfig: *types.ParallelDataConfig, // Required
	}

	if len(_translateClientToken) > 0 {
		input.ClientToken = aws.String(_translateClientToken)
	}
	if len(_translateName) > 0 {
		input.Name = aws.String(_translateName)
	}
	if len(_translateParallelDataConfig) > 0 {
		if err := assignInputField(input, "ParallelDataConfig", _translateParallelDataConfig); err != nil {
			log.Errorf("invalid --parallel-data-config: %s", err.Error())
			return
		}
	}
	if len(_translateDescription) > 0 {
		input.Description = aws.String(_translateDescription)
	}

	if resp, err := client.UpdateParallelData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_translateCmd)
	_translateCmd.Flags().SortFlags = false

	_translateCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_translateCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_translateCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_translateCmd.Flags().StringVarP(&_translateClientToken, "client-token", "", "", "Client Token")
	_translateCmd.Flags().StringVarP(&_translateDataAccessRoleArn, "data-access-role-arn", "", "", "Data Access Role ARN")
	_translateCmd.Flags().StringVarP(&_translateDescription, "description", "", "", "Description")
	_translateCmd.Flags().StringVarP(&_translateDisplayLanguageCode, "display-language-code", "", "", "Display Language Code")
	_translateCmd.Flags().StringVarP(&_translateDocument, "document", "", "", "Document")
	_translateCmd.Flags().StringVarP(&_translateEncryptionKey, "encryption-key", "", "", "Encryption Key")
	_translateCmd.Flags().StringVarP(&_translateFilter, "filter", "", "", "Filter")
	_translateCmd.Flags().StringVarP(&_translateInputDataConfig, "input-data-config", "", "", "Input Data Config")
	_translateCmd.Flags().StringVarP(&_translateJobId, "job-id", "", "", "Job ID")
	_translateCmd.Flags().StringVarP(&_translateJobName, "job-name", "", "", "Job Name")
	_translateCmd.Flags().StringVarP(&_translateMaxResults, "max-results", "", "", "Max Results")
	_translateCmd.Flags().StringVarP(&_translateMergeStrategy, "merge-strategy", "", "", "Merge Strategy")
	_translateCmd.Flags().StringVarP(&_translateName, "name", "", "", "Name")
	_translateCmd.Flags().StringVarP(&_translateNextToken, "next-token", "", "", "Next Token")
	_translateCmd.Flags().StringVarP(&_translateOutputDataConfig, "output-data-config", "", "", "Output Data Config")
	_translateCmd.Flags().StringVarP(&_translateParallelDataConfig, "parallel-data-config", "", "", "Parallel Data Config")
	_translateCmd.Flags().StringSliceVarP(&_translateParallelDataNames, "parallel-data-names", "", nil, "Parallel Data Names")
	_translateCmd.Flags().StringVarP(&_translateResourceArn, "resource-arn", "", "", "Resource ARN")
	_translateCmd.Flags().StringVarP(&_translateSettings, "settings", "", "", "Settings")
	_translateCmd.Flags().StringVarP(&_translateSourceLanguageCode, "source-language-code", "", "", "Source Language Code")
	_translateCmd.Flags().StringSliceVarP(&_translateTagKeys, "tag-keys", "", nil, "Tag Keys")
	_translateCmd.Flags().StringVarP(&_translateTags, "tags", "", "", "Tags")
	_translateCmd.Flags().StringVarP(&_translateTargetLanguageCode, "target-language-code", "", "", "Target Language Code")
	_translateCmd.Flags().StringSliceVarP(&_translateTargetLanguageCodes, "target-language-codes", "", nil, "Target Language Codes")
	_translateCmd.Flags().StringVarP(&_translateTerminologyData, "terminology-data", "", "", "Terminology Data")
	_translateCmd.Flags().StringVarP(&_translateTerminologyDataFormat, "terminology-data-format", "", "", "Terminology Data Format")
	_translateCmd.Flags().StringSliceVarP(&_translateTerminologyNames, "terminology-names", "", nil, "Terminology Names")
	_translateCmd.Flags().StringVarP(&_translateText, "text", "", "", "Text")

	_translateCmd.Flags().BoolVarP(&_translateCreateParallelData, "create-parallel-data", "", false, "Create Parallel Data")
	_translateCmd.Flags().BoolVarP(&_translateDeleteParallelData, "delete-parallel-data", "", false, "Delete Parallel Data")
	_translateCmd.Flags().BoolVarP(&_translateDeleteTerminology, "delete-terminology", "", false, "Delete Terminology")
	_translateCmd.Flags().BoolVarP(&_translateDescribeTextTranslationJob, "describe-text-translation-job", "", false, "Describe Text Translation Job")
	_translateCmd.Flags().BoolVarP(&_translateGetParallelData, "get-parallel-data", "", false, "Get Parallel Data")
	_translateCmd.Flags().BoolVarP(&_translateGetTerminology, "get-terminology", "", false, "Get Terminology")
	_translateCmd.Flags().BoolVarP(&_translateImportTerminology, "import-terminology", "", false, "Import Terminology")
	_translateCmd.Flags().BoolVarP(&_translateListLanguages, "list-languages", "", false, "List Languages")
	_translateCmd.Flags().BoolVarP(&_translateListParallelData, "list-parallel-data", "", false, "List Parallel Data")
	_translateCmd.Flags().BoolVarP(&_translateListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_translateCmd.Flags().BoolVarP(&_translateListTerminologies, "list-terminologies", "", false, "List Terminologies")
	_translateCmd.Flags().BoolVarP(&_translateListTextTranslationJobs, "list-text-translation-jobs", "", false, "List Text Translation Jobs")
	_translateCmd.Flags().BoolVarP(&_translateStartTextTranslationJob, "start-text-translation-job", "", false, "Start Text Translation Job")
	_translateCmd.Flags().BoolVarP(&_translateStopTextTranslationJob, "stop-text-translation-job", "", false, "Stop Text Translation Job")
	_translateCmd.Flags().BoolVarP(&_translateTagResource, "tag-resource", "", false, "Tag Resource")
	_translateCmd.Flags().BoolVarP(&_translateTranslateDocument, "translate-document", "", false, "Translate Document")
	_translateCmd.Flags().BoolVarP(&_translateTranslateText, "translate-text", "", false, "Translate Text")
	_translateCmd.Flags().BoolVarP(&_translateUntagResource, "untag-resource", "", false, "Untag Resource")
	_translateCmd.Flags().BoolVarP(&_translateUpdateParallelData, "update-parallel-data", "", false, "Update Parallel Data")

}
