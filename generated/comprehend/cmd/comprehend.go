package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/comprehend"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// comprehendCmd represents the comprehend command
var _comprehendCmd = &cobra.Command{
	Use:   "comprehend",
	Short: "AWS comprehend CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := comprehend.NewFromConfig(cfg)
		if _comprehendBatchDetectDominantLanguage {
			comprehend_BatchDetectDominantLanguage(cfg, client)
			return
		}
		if _comprehendBatchDetectEntities {
			comprehend_BatchDetectEntities(cfg, client)
			return
		}
		if _comprehendBatchDetectKeyPhrases {
			comprehend_BatchDetectKeyPhrases(cfg, client)
			return
		}
		if _comprehendBatchDetectSentiment {
			comprehend_BatchDetectSentiment(cfg, client)
			return
		}
		if _comprehendBatchDetectSyntax {
			comprehend_BatchDetectSyntax(cfg, client)
			return
		}
		if _comprehendBatchDetectTargetedSentiment {
			comprehend_BatchDetectTargetedSentiment(cfg, client)
			return
		}
		if _comprehendClassifyDocument {
			comprehend_ClassifyDocument(cfg, client)
			return
		}
		if _comprehendContainsPiiEntities {
			comprehend_ContainsPiiEntities(cfg, client)
			return
		}
		if _comprehendCreateDataset {
			comprehend_CreateDataset(cfg, client)
			return
		}
		if _comprehendCreateDocumentClassifier {
			comprehend_CreateDocumentClassifier(cfg, client)
			return
		}
		if _comprehendCreateEndpoint {
			comprehend_CreateEndpoint(cfg, client)
			return
		}
		if _comprehendCreateEntityRecognizer {
			comprehend_CreateEntityRecognizer(cfg, client)
			return
		}
		if _comprehendCreateFlywheel {
			comprehend_CreateFlywheel(cfg, client)
			return
		}
		if _comprehendDeleteDocumentClassifier {
			comprehend_DeleteDocumentClassifier(cfg, client)
			return
		}
		if _comprehendDeleteEndpoint {
			comprehend_DeleteEndpoint(cfg, client)
			return
		}
		if _comprehendDeleteEntityRecognizer {
			comprehend_DeleteEntityRecognizer(cfg, client)
			return
		}
		if _comprehendDeleteFlywheel {
			comprehend_DeleteFlywheel(cfg, client)
			return
		}
		if _comprehendDeleteResourcePolicy {
			comprehend_DeleteResourcePolicy(cfg, client)
			return
		}
		if _comprehendDescribeDataset {
			comprehend_DescribeDataset(cfg, client)
			return
		}
		if _comprehendDescribeDocumentClassificationJob {
			comprehend_DescribeDocumentClassificationJob(cfg, client)
			return
		}
		if _comprehendDescribeDocumentClassifier {
			comprehend_DescribeDocumentClassifier(cfg, client)
			return
		}
		if _comprehendDescribeDominantLanguageDetectionJob {
			comprehend_DescribeDominantLanguageDetectionJob(cfg, client)
			return
		}
		if _comprehendDescribeEndpoint {
			comprehend_DescribeEndpoint(cfg, client)
			return
		}
		if _comprehendDescribeEntitiesDetectionJob {
			comprehend_DescribeEntitiesDetectionJob(cfg, client)
			return
		}
		if _comprehendDescribeEntityRecognizer {
			comprehend_DescribeEntityRecognizer(cfg, client)
			return
		}
		if _comprehendDescribeEventsDetectionJob {
			comprehend_DescribeEventsDetectionJob(cfg, client)
			return
		}
		if _comprehendDescribeFlywheel {
			comprehend_DescribeFlywheel(cfg, client)
			return
		}
		if _comprehendDescribeFlywheelIteration {
			comprehend_DescribeFlywheelIteration(cfg, client)
			return
		}
		if _comprehendDescribeKeyPhrasesDetectionJob {
			comprehend_DescribeKeyPhrasesDetectionJob(cfg, client)
			return
		}
		if _comprehendDescribePiiEntitiesDetectionJob {
			comprehend_DescribePiiEntitiesDetectionJob(cfg, client)
			return
		}
		if _comprehendDescribeResourcePolicy {
			comprehend_DescribeResourcePolicy(cfg, client)
			return
		}
		if _comprehendDescribeSentimentDetectionJob {
			comprehend_DescribeSentimentDetectionJob(cfg, client)
			return
		}
		if _comprehendDescribeTargetedSentimentDetectionJob {
			comprehend_DescribeTargetedSentimentDetectionJob(cfg, client)
			return
		}
		if _comprehendDescribeTopicsDetectionJob {
			comprehend_DescribeTopicsDetectionJob(cfg, client)
			return
		}
		if _comprehendDetectDominantLanguage {
			comprehend_DetectDominantLanguage(cfg, client)
			return
		}
		if _comprehendDetectEntities {
			comprehend_DetectEntities(cfg, client)
			return
		}
		if _comprehendDetectKeyPhrases {
			comprehend_DetectKeyPhrases(cfg, client)
			return
		}
		if _comprehendDetectPiiEntities {
			comprehend_DetectPiiEntities(cfg, client)
			return
		}
		if _comprehendDetectSentiment {
			comprehend_DetectSentiment(cfg, client)
			return
		}
		if _comprehendDetectSyntax {
			comprehend_DetectSyntax(cfg, client)
			return
		}
		if _comprehendDetectTargetedSentiment {
			comprehend_DetectTargetedSentiment(cfg, client)
			return
		}
		if _comprehendDetectToxicContent {
			comprehend_DetectToxicContent(cfg, client)
			return
		}
		if _comprehendImportModel {
			comprehend_ImportModel(cfg, client)
			return
		}
		if _comprehendListDatasets {
			comprehend_ListDatasets(cfg, client)
			return
		}
		if _comprehendListDocumentClassificationJobs {
			comprehend_ListDocumentClassificationJobs(cfg, client)
			return
		}
		if _comprehendListDocumentClassifierSummaries {
			comprehend_ListDocumentClassifierSummaries(cfg, client)
			return
		}
		if _comprehendListDocumentClassifiers {
			comprehend_ListDocumentClassifiers(cfg, client)
			return
		}
		if _comprehendListDominantLanguageDetectionJobs {
			comprehend_ListDominantLanguageDetectionJobs(cfg, client)
			return
		}
		if _comprehendListEndpoints {
			comprehend_ListEndpoints(cfg, client)
			return
		}
		if _comprehendListEntitiesDetectionJobs {
			comprehend_ListEntitiesDetectionJobs(cfg, client)
			return
		}
		if _comprehendListEntityRecognizerSummaries {
			comprehend_ListEntityRecognizerSummaries(cfg, client)
			return
		}
		if _comprehendListEntityRecognizers {
			comprehend_ListEntityRecognizers(cfg, client)
			return
		}
		if _comprehendListEventsDetectionJobs {
			comprehend_ListEventsDetectionJobs(cfg, client)
			return
		}
		if _comprehendListFlywheelIterationHistory {
			comprehend_ListFlywheelIterationHistory(cfg, client)
			return
		}
		if _comprehendListFlywheels {
			comprehend_ListFlywheels(cfg, client)
			return
		}
		if _comprehendListKeyPhrasesDetectionJobs {
			comprehend_ListKeyPhrasesDetectionJobs(cfg, client)
			return
		}
		if _comprehendListPiiEntitiesDetectionJobs {
			comprehend_ListPiiEntitiesDetectionJobs(cfg, client)
			return
		}
		if _comprehendListSentimentDetectionJobs {
			comprehend_ListSentimentDetectionJobs(cfg, client)
			return
		}
		if _comprehendListTagsForResource {
			comprehend_ListTagsForResource(cfg, client)
			return
		}
		if _comprehendListTargetedSentimentDetectionJobs {
			comprehend_ListTargetedSentimentDetectionJobs(cfg, client)
			return
		}
		if _comprehendListTopicsDetectionJobs {
			comprehend_ListTopicsDetectionJobs(cfg, client)
			return
		}
		if _comprehendPutResourcePolicy {
			comprehend_PutResourcePolicy(cfg, client)
			return
		}
		if _comprehendStartDocumentClassificationJob {
			comprehend_StartDocumentClassificationJob(cfg, client)
			return
		}
		if _comprehendStartDominantLanguageDetectionJob {
			comprehend_StartDominantLanguageDetectionJob(cfg, client)
			return
		}
		if _comprehendStartEntitiesDetectionJob {
			comprehend_StartEntitiesDetectionJob(cfg, client)
			return
		}
		if _comprehendStartEventsDetectionJob {
			comprehend_StartEventsDetectionJob(cfg, client)
			return
		}
		if _comprehendStartFlywheelIteration {
			comprehend_StartFlywheelIteration(cfg, client)
			return
		}
		if _comprehendStartKeyPhrasesDetectionJob {
			comprehend_StartKeyPhrasesDetectionJob(cfg, client)
			return
		}
		if _comprehendStartPiiEntitiesDetectionJob {
			comprehend_StartPiiEntitiesDetectionJob(cfg, client)
			return
		}
		if _comprehendStartSentimentDetectionJob {
			comprehend_StartSentimentDetectionJob(cfg, client)
			return
		}
		if _comprehendStartTargetedSentimentDetectionJob {
			comprehend_StartTargetedSentimentDetectionJob(cfg, client)
			return
		}
		if _comprehendStartTopicsDetectionJob {
			comprehend_StartTopicsDetectionJob(cfg, client)
			return
		}
		if _comprehendStopDominantLanguageDetectionJob {
			comprehend_StopDominantLanguageDetectionJob(cfg, client)
			return
		}
		if _comprehendStopEntitiesDetectionJob {
			comprehend_StopEntitiesDetectionJob(cfg, client)
			return
		}
		if _comprehendStopEventsDetectionJob {
			comprehend_StopEventsDetectionJob(cfg, client)
			return
		}
		if _comprehendStopKeyPhrasesDetectionJob {
			comprehend_StopKeyPhrasesDetectionJob(cfg, client)
			return
		}
		if _comprehendStopPiiEntitiesDetectionJob {
			comprehend_StopPiiEntitiesDetectionJob(cfg, client)
			return
		}
		if _comprehendStopSentimentDetectionJob {
			comprehend_StopSentimentDetectionJob(cfg, client)
			return
		}
		if _comprehendStopTargetedSentimentDetectionJob {
			comprehend_StopTargetedSentimentDetectionJob(cfg, client)
			return
		}
		if _comprehendStopTrainingDocumentClassifier {
			comprehend_StopTrainingDocumentClassifier(cfg, client)
			return
		}
		if _comprehendStopTrainingEntityRecognizer {
			comprehend_StopTrainingEntityRecognizer(cfg, client)
			return
		}
		if _comprehendTagResource {
			comprehend_TagResource(cfg, client)
			return
		}
		if _comprehendUntagResource {
			comprehend_UntagResource(cfg, client)
			return
		}
		if _comprehendUpdateEndpoint {
			comprehend_UpdateEndpoint(cfg, client)
			return
		}
		if _comprehendUpdateFlywheel {
			comprehend_UpdateFlywheel(cfg, client)
			return
		}

	},
}

var (
	_comprehendBatchDetectDominantLanguage           bool
	_comprehendBatchDetectEntities                   bool
	_comprehendBatchDetectKeyPhrases                 bool
	_comprehendBatchDetectSentiment                  bool
	_comprehendBatchDetectSyntax                     bool
	_comprehendBatchDetectTargetedSentiment          bool
	_comprehendClassifyDocument                      bool
	_comprehendContainsPiiEntities                   bool
	_comprehendCreateDataset                         bool
	_comprehendCreateDocumentClassifier              bool
	_comprehendCreateEndpoint                        bool
	_comprehendCreateEntityRecognizer                bool
	_comprehendCreateFlywheel                        bool
	_comprehendDeleteDocumentClassifier              bool
	_comprehendDeleteEndpoint                        bool
	_comprehendDeleteEntityRecognizer                bool
	_comprehendDeleteFlywheel                        bool
	_comprehendDeleteResourcePolicy                  bool
	_comprehendDescribeDataset                       bool
	_comprehendDescribeDocumentClassificationJob     bool
	_comprehendDescribeDocumentClassifier            bool
	_comprehendDescribeDominantLanguageDetectionJob  bool
	_comprehendDescribeEndpoint                      bool
	_comprehendDescribeEntitiesDetectionJob          bool
	_comprehendDescribeEntityRecognizer              bool
	_comprehendDescribeEventsDetectionJob            bool
	_comprehendDescribeFlywheel                      bool
	_comprehendDescribeFlywheelIteration             bool
	_comprehendDescribeKeyPhrasesDetectionJob        bool
	_comprehendDescribePiiEntitiesDetectionJob       bool
	_comprehendDescribeResourcePolicy                bool
	_comprehendDescribeSentimentDetectionJob         bool
	_comprehendDescribeTargetedSentimentDetectionJob bool
	_comprehendDescribeTopicsDetectionJob            bool
	_comprehendDetectDominantLanguage                bool
	_comprehendDetectEntities                        bool
	_comprehendDetectKeyPhrases                      bool
	_comprehendDetectPiiEntities                     bool
	_comprehendDetectSentiment                       bool
	_comprehendDetectSyntax                          bool
	_comprehendDetectTargetedSentiment               bool
	_comprehendDetectToxicContent                    bool
	_comprehendImportModel                           bool
	_comprehendListDatasets                          bool
	_comprehendListDocumentClassificationJobs        bool
	_comprehendListDocumentClassifierSummaries       bool
	_comprehendListDocumentClassifiers               bool
	_comprehendListDominantLanguageDetectionJobs     bool
	_comprehendListEndpoints                         bool
	_comprehendListEntitiesDetectionJobs             bool
	_comprehendListEntityRecognizerSummaries         bool
	_comprehendListEntityRecognizers                 bool
	_comprehendListEventsDetectionJobs               bool
	_comprehendListFlywheelIterationHistory          bool
	_comprehendListFlywheels                         bool
	_comprehendListKeyPhrasesDetectionJobs           bool
	_comprehendListPiiEntitiesDetectionJobs          bool
	_comprehendListSentimentDetectionJobs            bool
	_comprehendListTagsForResource                   bool
	_comprehendListTargetedSentimentDetectionJobs    bool
	_comprehendListTopicsDetectionJobs               bool
	_comprehendPutResourcePolicy                     bool
	_comprehendStartDocumentClassificationJob        bool
	_comprehendStartDominantLanguageDetectionJob     bool
	_comprehendStartEntitiesDetectionJob             bool
	_comprehendStartEventsDetectionJob               bool
	_comprehendStartFlywheelIteration                bool
	_comprehendStartKeyPhrasesDetectionJob           bool
	_comprehendStartPiiEntitiesDetectionJob          bool
	_comprehendStartSentimentDetectionJob            bool
	_comprehendStartTargetedSentimentDetectionJob    bool
	_comprehendStartTopicsDetectionJob               bool
	_comprehendStopDominantLanguageDetectionJob      bool
	_comprehendStopEntitiesDetectionJob              bool
	_comprehendStopEventsDetectionJob                bool
	_comprehendStopKeyPhrasesDetectionJob            bool
	_comprehendStopPiiEntitiesDetectionJob           bool
	_comprehendStopSentimentDetectionJob             bool
	_comprehendStopTargetedSentimentDetectionJob     bool
	_comprehendStopTrainingDocumentClassifier        bool
	_comprehendStopTrainingEntityRecognizer          bool
	_comprehendTagResource                           bool
	_comprehendUntagResource                         bool
	_comprehendUpdateEndpoint                        bool
	_comprehendUpdateFlywheel                        bool

	_comprehendActiveModelArn           string
	_comprehendBytes                    string
	_comprehendClientRequestToken       string
	_comprehendDataAccessRoleArn        string
	_comprehendDataLakeS3Uri            string
	_comprehendDataSecurityConfig       string
	_comprehendDatasetArn               string
	_comprehendDatasetName              string
	_comprehendDatasetType              string
	_comprehendDescription              string
	_comprehendDesiredDataAccessRoleArn string
	_comprehendDesiredInferenceUnits    string
	_comprehendDesiredModelArn          string
	_comprehendDocumentClassifierArn    string
	_comprehendDocumentClassifierName   string
	_comprehendDocumentReaderConfig     string
	_comprehendEndpointArn              string
	_comprehendEndpointName             string
	_comprehendEntityRecognizerArn      string
	_comprehendFilter                   string
	_comprehendFlywheelArn              string
	_comprehendFlywheelIterationId      string
	_comprehendFlywheelName             string
	_comprehendInputDataConfig          string
	_comprehendJobId                    string
	_comprehendJobName                  string
	_comprehendLanguageCode             string
	_comprehendMaxResults               string
	_comprehendMode                     string
	_comprehendModelArn                 string
	_comprehendModelKmsKeyId            string
	_comprehendModelName                string
	_comprehendModelPolicy              string
	_comprehendModelType                string
	_comprehendNextToken                string
	_comprehendNumberOfTopics           string
	_comprehendOutputDataConfig         string
	_comprehendPolicyRevisionId         string
	_comprehendRecognizerName           string
	_comprehendRedactionConfig          string
	_comprehendResourceArn              string
	_comprehendResourcePolicy           string
	_comprehendSourceModelArn           string
	_comprehendTagKeys                  []string
	_comprehendTags                     string
	_comprehendTargetEventTypes         []string
	_comprehendTaskConfig               string
	_comprehendText                     string
	_comprehendTextList                 []string
	_comprehendTextSegments             string
	_comprehendVersionName              string
	_comprehendVolumeKmsKeyId           string
	_comprehendVpcConfig                string
)

// Determines the dominant language of the input text for a batch of documents.
// For a list of languages that Amazon Comprehend can detect, see [Amazon Comprehend Supported Languages].
//
// [Amazon Comprehend Supported Languages]: https://docs.aws.amazon.com/comprehend/latest/dg/how-languages.html
func comprehend_BatchDetectDominantLanguage(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.BatchDetectDominantLanguageInput{
		// TextList: []string, // Required
	}

	if len(_comprehendTextList) > 0 {
		input.TextList = append([]string(nil), _comprehendTextList...)
	}

	if resp, err := client.BatchDetectDominantLanguage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Inspects the text of a batch of documents for named entities and returns
// information about them. For more information about named entities, see [Entities]in the
// Comprehend Developer Guide.
//
// [Entities]: https://docs.aws.amazon.com/comprehend/latest/dg/how-entities.html
func comprehend_BatchDetectEntities(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.BatchDetectEntitiesInput{
		// LanguageCode: types.LanguageCode, // Required
		// TextList: []string, // Required
	}

	if len(_comprehendLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendTextList) > 0 {
		input.TextList = append([]string(nil), _comprehendTextList...)
	}

	if resp, err := client.BatchDetectEntities(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detects the key noun phrases found in a batch of documents.
func comprehend_BatchDetectKeyPhrases(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.BatchDetectKeyPhrasesInput{
		// LanguageCode: types.LanguageCode, // Required
		// TextList: []string, // Required
	}

	if len(_comprehendLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendTextList) > 0 {
		input.TextList = append([]string(nil), _comprehendTextList...)
	}

	if resp, err := client.BatchDetectKeyPhrases(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Inspects a batch of documents and returns an inference of the prevailing
// sentiment, POSITIVE , NEUTRAL , MIXED , or NEGATIVE , in each one.
func comprehend_BatchDetectSentiment(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.BatchDetectSentimentInput{
		// LanguageCode: types.LanguageCode, // Required
		// TextList: []string, // Required
	}

	if len(_comprehendLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendTextList) > 0 {
		input.TextList = append([]string(nil), _comprehendTextList...)
	}

	if resp, err := client.BatchDetectSentiment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Inspects the text of a batch of documents for the syntax and part of speech of
// the words in the document and returns information about them. For more
// information, see [Syntax]in the Comprehend Developer Guide.
//
// [Syntax]: https://docs.aws.amazon.com/comprehend/latest/dg/how-syntax.html
func comprehend_BatchDetectSyntax(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.BatchDetectSyntaxInput{
		// LanguageCode: types.SyntaxLanguageCode, // Required
		// TextList: []string, // Required
	}

	if len(_comprehendLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendTextList) > 0 {
		input.TextList = append([]string(nil), _comprehendTextList...)
	}

	if resp, err := client.BatchDetectSyntax(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Inspects a batch of documents and returns a sentiment analysis for each entity
// identified in the documents.
//
// For more information about targeted sentiment, see [Targeted sentiment] in the Amazon Comprehend
// Developer Guide.
//
// [Targeted sentiment]: https://docs.aws.amazon.com/comprehend/latest/dg/how-targeted-sentiment.html
func comprehend_BatchDetectTargetedSentiment(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.BatchDetectTargetedSentimentInput{
		// LanguageCode: types.LanguageCode, // Required
		// TextList: []string, // Required
	}

	if len(_comprehendLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendTextList) > 0 {
		input.TextList = append([]string(nil), _comprehendTextList...)
	}

	if resp, err := client.BatchDetectTargetedSentiment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a classification request to analyze a single document in real-time.
// ClassifyDocument supports the following model types:
//
// - Custom classifier - a custom model that you have created and trained. For
// input, you can provide plain text, a single-page document (PDF, Word, or image),
// or Amazon Textract API output. For more information, see [Custom classification]in the Amazon
// Comprehend Developer Guide.
//
// - Prompt safety classifier - Amazon Comprehend provides a pre-trained model
// for classifying input prompts for generative AI applications. For input, you
// provide English plain text input. For prompt safety classification, the response
// includes only the Classes field. For more information about prompt safety
// classifiers, see [Prompt safety classification]in the Amazon Comprehend Developer Guide.
//
// If the system detects errors while processing a page in the input document, the
// API response includes an Errors field that describes the errors.
//
// If the system detects a document-level error in your input document, the API
// returns an InvalidRequestException error response. For details about this
// exception, see [Errors in semi-structured documents]in the Comprehend Developer Guide.
//
// [Custom classification]: https://docs.aws.amazon.com/comprehend/latest/dg/how-document-classification.html
// [Prompt safety classification]: https://docs.aws.amazon.com/comprehend/latest/dg/trust-safety.html#prompt-classification
// [Errors in semi-structured documents]: https://docs.aws.amazon.com/comprehend/latest/dg/idp-inputs-sync-err.html
func comprehend_ClassifyDocument(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.ClassifyDocumentInput{
		// EndpointArn: *string, // Required
	}

	if len(_comprehendEndpointArn) > 0 {
		input.EndpointArn = aws.String(_comprehendEndpointArn)
	}
	if len(_comprehendBytes) > 0 {
		if err := assignInputField(input, "Bytes", _comprehendBytes); err != nil {
			log.Errorf("invalid --bytes: %s", err.Error())
			return
		}
	}
	if len(_comprehendDocumentReaderConfig) > 0 {
		if err := assignInputField(input, "DocumentReaderConfig", _comprehendDocumentReaderConfig); err != nil {
			log.Errorf("invalid --document-reader-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendText) > 0 {
		input.Text = aws.String(_comprehendText)
	}

	if resp, err := client.ClassifyDocument(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Analyzes input text for the presence of personally identifiable information
// (PII) and returns the labels of identified PII entity types such as name,
// address, bank account number, or phone number.
func comprehend_ContainsPiiEntities(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.ContainsPiiEntitiesInput{
		// LanguageCode: types.LanguageCode, // Required
		// Text: *string, // Required
	}

	if len(_comprehendLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendText) > 0 {
		input.Text = aws.String(_comprehendText)
	}

	if resp, err := client.ContainsPiiEntities(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a dataset to upload training or test data for a model associated with a
// flywheel. For more information about datasets, see [Flywheel overview]in the Amazon Comprehend
// Developer Guide.
//
// [Flywheel overview]: https://docs.aws.amazon.com/comprehend/latest/dg/flywheels-about.html
func comprehend_CreateDataset(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.CreateDatasetInput{
		// DatasetName: *string, // Required
		// FlywheelArn: *string, // Required
		// InputDataConfig: *types.DatasetInputDataConfig, // Required
	}

	if len(_comprehendDatasetName) > 0 {
		input.DatasetName = aws.String(_comprehendDatasetName)
	}
	if len(_comprehendFlywheelArn) > 0 {
		input.FlywheelArn = aws.String(_comprehendFlywheelArn)
	}
	if len(_comprehendInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _comprehendInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_comprehendClientRequestToken)
	}
	if len(_comprehendDatasetType) > 0 {
		if err := assignInputField(input, "DatasetType", _comprehendDatasetType); err != nil {
			log.Errorf("invalid --dataset-type: %s", err.Error())
			return
		}
	}
	if len(_comprehendDescription) > 0 {
		input.Description = aws.String(_comprehendDescription)
	}
	if len(_comprehendTags) > 0 {
		if err := assignInputField(input, "Tags", _comprehendTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new document classifier that you can use to categorize documents. To
// create a classifier, you provide a set of training documents that are labeled
// with the categories that you want to use. For more information, see [Training classifier models]in the
// Comprehend Developer Guide.
//
// [Training classifier models]: https://docs.aws.amazon.com/comprehend/latest/dg/training-classifier-model.html
func comprehend_CreateDocumentClassifier(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.CreateDocumentClassifierInput{
		// DataAccessRoleArn: *string, // Required
		// DocumentClassifierName: *string, // Required
		// InputDataConfig: *types.DocumentClassifierInputDataConfig, // Required
		// LanguageCode: types.LanguageCode, // Required
	}

	if len(_comprehendDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_comprehendDataAccessRoleArn)
	}
	if len(_comprehendDocumentClassifierName) > 0 {
		input.DocumentClassifierName = aws.String(_comprehendDocumentClassifierName)
	}
	if len(_comprehendInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _comprehendInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_comprehendClientRequestToken)
	}
	if len(_comprehendMode) > 0 {
		if err := assignInputField(input, "Mode", _comprehendMode); err != nil {
			log.Errorf("invalid --mode: %s", err.Error())
			return
		}
	}
	if len(_comprehendModelKmsKeyId) > 0 {
		input.ModelKmsKeyId = aws.String(_comprehendModelKmsKeyId)
	}
	if len(_comprehendModelPolicy) > 0 {
		input.ModelPolicy = aws.String(_comprehendModelPolicy)
	}
	if len(_comprehendOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _comprehendOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendTags) > 0 {
		if err := assignInputField(input, "Tags", _comprehendTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_comprehendVersionName) > 0 {
		input.VersionName = aws.String(_comprehendVersionName)
	}
	if len(_comprehendVolumeKmsKeyId) > 0 {
		input.VolumeKmsKeyId = aws.String(_comprehendVolumeKmsKeyId)
	}
	if len(_comprehendVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _comprehendVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDocumentClassifier(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a model-specific endpoint for synchronous inference for a previously
// trained custom model For information about endpoints, see [Managing endpoints].
//
// [Managing endpoints]: https://docs.aws.amazon.com/comprehend/latest/dg/manage-endpoints.html
func comprehend_CreateEndpoint(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.CreateEndpointInput{
		// DesiredInferenceUnits: *int32, // Required
		// EndpointName: *string, // Required
	}

	if len(_comprehendDesiredInferenceUnits) > 0 {
		if err := assignInputField(input, "DesiredInferenceUnits", _comprehendDesiredInferenceUnits); err != nil {
			log.Errorf("invalid --desired-inference-units: %s", err.Error())
			return
		}
	}
	if len(_comprehendEndpointName) > 0 {
		input.EndpointName = aws.String(_comprehendEndpointName)
	}
	if len(_comprehendClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_comprehendClientRequestToken)
	}
	if len(_comprehendDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_comprehendDataAccessRoleArn)
	}
	if len(_comprehendFlywheelArn) > 0 {
		input.FlywheelArn = aws.String(_comprehendFlywheelArn)
	}
	if len(_comprehendModelArn) > 0 {
		input.ModelArn = aws.String(_comprehendModelArn)
	}
	if len(_comprehendTags) > 0 {
		if err := assignInputField(input, "Tags", _comprehendTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an entity recognizer using submitted files. After your
// CreateEntityRecognizer request is submitted, you can check job status using the
// DescribeEntityRecognizer API.
func comprehend_CreateEntityRecognizer(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.CreateEntityRecognizerInput{
		// DataAccessRoleArn: *string, // Required
		// InputDataConfig: *types.EntityRecognizerInputDataConfig, // Required
		// LanguageCode: types.LanguageCode, // Required
		// RecognizerName: *string, // Required
	}

	if len(_comprehendDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_comprehendDataAccessRoleArn)
	}
	if len(_comprehendInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _comprehendInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendRecognizerName) > 0 {
		input.RecognizerName = aws.String(_comprehendRecognizerName)
	}
	if len(_comprehendClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_comprehendClientRequestToken)
	}
	if len(_comprehendModelKmsKeyId) > 0 {
		input.ModelKmsKeyId = aws.String(_comprehendModelKmsKeyId)
	}
	if len(_comprehendModelPolicy) > 0 {
		input.ModelPolicy = aws.String(_comprehendModelPolicy)
	}
	if len(_comprehendTags) > 0 {
		if err := assignInputField(input, "Tags", _comprehendTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_comprehendVersionName) > 0 {
		input.VersionName = aws.String(_comprehendVersionName)
	}
	if len(_comprehendVolumeKmsKeyId) > 0 {
		input.VolumeKmsKeyId = aws.String(_comprehendVolumeKmsKeyId)
	}
	if len(_comprehendVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _comprehendVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEntityRecognizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A flywheel is an Amazon Web Services resource that orchestrates the ongoing
// training of a model for custom classification or custom entity recognition. You
// can create a flywheel to start with an existing trained model, or Comprehend can
// create and train a new model.
//
// When you create the flywheel, Comprehend creates a data lake in your account.
// The data lake holds the training data and test data for all versions of the
// model.
//
// To use a flywheel with an existing trained model, you specify the active model
// version. Comprehend copies the model's training data and test data into the
// flywheel's data lake.
//
// To use the flywheel with a new model, you need to provide a dataset for
// training data (and optional test data) when you create the flywheel.
//
// For more information about flywheels, see [Flywheel overview] in the Amazon Comprehend Developer
// Guide.
//
// [Flywheel overview]: https://docs.aws.amazon.com/comprehend/latest/dg/flywheels-about.html
func comprehend_CreateFlywheel(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.CreateFlywheelInput{
		// DataAccessRoleArn: *string, // Required
		// DataLakeS3Uri: *string, // Required
		// FlywheelName: *string, // Required
	}

	if len(_comprehendDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_comprehendDataAccessRoleArn)
	}
	if len(_comprehendDataLakeS3Uri) > 0 {
		input.DataLakeS3Uri = aws.String(_comprehendDataLakeS3Uri)
	}
	if len(_comprehendFlywheelName) > 0 {
		input.FlywheelName = aws.String(_comprehendFlywheelName)
	}
	if len(_comprehendActiveModelArn) > 0 {
		input.ActiveModelArn = aws.String(_comprehendActiveModelArn)
	}
	if len(_comprehendClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_comprehendClientRequestToken)
	}
	if len(_comprehendDataSecurityConfig) > 0 {
		if err := assignInputField(input, "DataSecurityConfig", _comprehendDataSecurityConfig); err != nil {
			log.Errorf("invalid --data-security-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendModelType) > 0 {
		if err := assignInputField(input, "ModelType", _comprehendModelType); err != nil {
			log.Errorf("invalid --model-type: %s", err.Error())
			return
		}
	}
	if len(_comprehendTags) > 0 {
		if err := assignInputField(input, "Tags", _comprehendTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_comprehendTaskConfig) > 0 {
		if err := assignInputField(input, "TaskConfig", _comprehendTaskConfig); err != nil {
			log.Errorf("invalid --task-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFlywheel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a previously created document classifier
// Only those classifiers that are in terminated states (IN_ERROR, TRAINED) will
// be deleted. If an active inference job is using the model, a
// ResourceInUseException will be returned.
//
// This is an asynchronous action that puts the classifier into a DELETING state,
// and it is then removed by a background job. Once removed, the classifier
// disappears from your account and is no longer available for use.
func comprehend_DeleteDocumentClassifier(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DeleteDocumentClassifierInput{
		// DocumentClassifierArn: *string, // Required
	}

	if len(_comprehendDocumentClassifierArn) > 0 {
		input.DocumentClassifierArn = aws.String(_comprehendDocumentClassifierArn)
	}

	if resp, err := client.DeleteDocumentClassifier(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a model-specific endpoint for a previously-trained custom model. All
// endpoints must be deleted in order for the model to be deleted. For information
// about endpoints, see [Managing endpoints].
//
// [Managing endpoints]: https://docs.aws.amazon.com/comprehend/latest/dg/manage-endpoints.html
func comprehend_DeleteEndpoint(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DeleteEndpointInput{
		// EndpointArn: *string, // Required
	}

	if len(_comprehendEndpointArn) > 0 {
		input.EndpointArn = aws.String(_comprehendEndpointArn)
	}

	if resp, err := client.DeleteEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an entity recognizer.
// Only those recognizers that are in terminated states (IN_ERROR, TRAINED) will
// be deleted. If an active inference job is using the model, a
// ResourceInUseException will be returned.
//
// This is an asynchronous action that puts the recognizer into a DELETING state,
// and it is then removed by a background job. Once removed, the recognizer
// disappears from your account and is no longer available for use.
func comprehend_DeleteEntityRecognizer(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DeleteEntityRecognizerInput{
		// EntityRecognizerArn: *string, // Required
	}

	if len(_comprehendEntityRecognizerArn) > 0 {
		input.EntityRecognizerArn = aws.String(_comprehendEntityRecognizerArn)
	}

	if resp, err := client.DeleteEntityRecognizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a flywheel. When you delete the flywheel, Amazon Comprehend does not
// delete the data lake or the model associated with the flywheel.
//
// For more information about flywheels, see [Flywheel overview] in the Amazon Comprehend Developer
// Guide.
//
// [Flywheel overview]: https://docs.aws.amazon.com/comprehend/latest/dg/flywheels-about.html
func comprehend_DeleteFlywheel(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DeleteFlywheelInput{
		// FlywheelArn: *string, // Required
	}

	if len(_comprehendFlywheelArn) > 0 {
		input.FlywheelArn = aws.String(_comprehendFlywheelArn)
	}

	if resp, err := client.DeleteFlywheel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a resource-based policy that is attached to a custom model.
func comprehend_DeleteResourcePolicy(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DeleteResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_comprehendResourceArn) > 0 {
		input.ResourceArn = aws.String(_comprehendResourceArn)
	}
	if len(_comprehendPolicyRevisionId) > 0 {
		input.PolicyRevisionId = aws.String(_comprehendPolicyRevisionId)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the dataset that you specify. For more information
// about datasets, see [Flywheel overview]in the Amazon Comprehend Developer Guide.
//
// [Flywheel overview]: https://docs.aws.amazon.com/comprehend/latest/dg/flywheels-about.html
func comprehend_DescribeDataset(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DescribeDatasetInput{
		// DatasetArn: *string, // Required
	}

	if len(_comprehendDatasetArn) > 0 {
		input.DatasetArn = aws.String(_comprehendDatasetArn)
	}

	if resp, err := client.DescribeDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the properties associated with a document classification job. Use this
// operation to get the status of a classification job.
func comprehend_DescribeDocumentClassificationJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DescribeDocumentClassificationJobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendJobId) > 0 {
		input.JobId = aws.String(_comprehendJobId)
	}

	if resp, err := client.DescribeDocumentClassificationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the properties associated with a document classifier.
func comprehend_DescribeDocumentClassifier(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DescribeDocumentClassifierInput{
		// DocumentClassifierArn: *string, // Required
	}

	if len(_comprehendDocumentClassifierArn) > 0 {
		input.DocumentClassifierArn = aws.String(_comprehendDocumentClassifierArn)
	}

	if resp, err := client.DescribeDocumentClassifier(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the properties associated with a dominant language detection job. Use this
// operation to get the status of a detection job.
func comprehend_DescribeDominantLanguageDetectionJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DescribeDominantLanguageDetectionJobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendJobId) > 0 {
		input.JobId = aws.String(_comprehendJobId)
	}

	if resp, err := client.DescribeDominantLanguageDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the properties associated with a specific endpoint. Use this operation to
// get the status of an endpoint. For information about endpoints, see [Managing endpoints].
//
// [Managing endpoints]: https://docs.aws.amazon.com/comprehend/latest/dg/manage-endpoints.html
func comprehend_DescribeEndpoint(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DescribeEndpointInput{
		// EndpointArn: *string, // Required
	}

	if len(_comprehendEndpointArn) > 0 {
		input.EndpointArn = aws.String(_comprehendEndpointArn)
	}

	if resp, err := client.DescribeEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the properties associated with an entities detection job. Use this
// operation to get the status of a detection job.
func comprehend_DescribeEntitiesDetectionJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DescribeEntitiesDetectionJobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendJobId) > 0 {
		input.JobId = aws.String(_comprehendJobId)
	}

	if resp, err := client.DescribeEntitiesDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides details about an entity recognizer including status, S3 buckets
// containing training data, recognizer metadata, metrics, and so on.
func comprehend_DescribeEntityRecognizer(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DescribeEntityRecognizerInput{
		// EntityRecognizerArn: *string, // Required
	}

	if len(_comprehendEntityRecognizerArn) > 0 {
		input.EntityRecognizerArn = aws.String(_comprehendEntityRecognizerArn)
	}

	if resp, err := client.DescribeEntityRecognizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the status and details of an events detection job.
func comprehend_DescribeEventsDetectionJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DescribeEventsDetectionJobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendJobId) > 0 {
		input.JobId = aws.String(_comprehendJobId)
	}

	if resp, err := client.DescribeEventsDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides configuration information about the flywheel. For more information
// about flywheels, see [Flywheel overview]in the Amazon Comprehend Developer Guide.
//
// [Flywheel overview]: https://docs.aws.amazon.com/comprehend/latest/dg/flywheels-about.html
func comprehend_DescribeFlywheel(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DescribeFlywheelInput{
		// FlywheelArn: *string, // Required
	}

	if len(_comprehendFlywheelArn) > 0 {
		input.FlywheelArn = aws.String(_comprehendFlywheelArn)
	}

	if resp, err := client.DescribeFlywheel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve the configuration properties of a flywheel iteration. For more
// information about flywheels, see [Flywheel overview]in the Amazon Comprehend Developer Guide.
//
// [Flywheel overview]: https://docs.aws.amazon.com/comprehend/latest/dg/flywheels-about.html
func comprehend_DescribeFlywheelIteration(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DescribeFlywheelIterationInput{
		// FlywheelArn: *string, // Required
		// FlywheelIterationId: *string, // Required
	}

	if len(_comprehendFlywheelArn) > 0 {
		input.FlywheelArn = aws.String(_comprehendFlywheelArn)
	}
	if len(_comprehendFlywheelIterationId) > 0 {
		input.FlywheelIterationId = aws.String(_comprehendFlywheelIterationId)
	}

	if resp, err := client.DescribeFlywheelIteration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the properties associated with a key phrases detection job. Use this
// operation to get the status of a detection job.
func comprehend_DescribeKeyPhrasesDetectionJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DescribeKeyPhrasesDetectionJobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendJobId) > 0 {
		input.JobId = aws.String(_comprehendJobId)
	}

	if resp, err := client.DescribeKeyPhrasesDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the properties associated with a PII entities detection job. For example,
// you can use this operation to get the job status.
func comprehend_DescribePiiEntitiesDetectionJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DescribePiiEntitiesDetectionJobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendJobId) > 0 {
		input.JobId = aws.String(_comprehendJobId)
	}

	if resp, err := client.DescribePiiEntitiesDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the details of a resource-based policy that is attached to a custom model,
// including the JSON body of the policy.
func comprehend_DescribeResourcePolicy(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DescribeResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_comprehendResourceArn) > 0 {
		input.ResourceArn = aws.String(_comprehendResourceArn)
	}

	if resp, err := client.DescribeResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the properties associated with a sentiment detection job. Use this
// operation to get the status of a detection job.
func comprehend_DescribeSentimentDetectionJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DescribeSentimentDetectionJobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendJobId) > 0 {
		input.JobId = aws.String(_comprehendJobId)
	}

	if resp, err := client.DescribeSentimentDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the properties associated with a targeted sentiment detection job. Use
// this operation to get the status of the job.
func comprehend_DescribeTargetedSentimentDetectionJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DescribeTargetedSentimentDetectionJobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendJobId) > 0 {
		input.JobId = aws.String(_comprehendJobId)
	}

	if resp, err := client.DescribeTargetedSentimentDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the properties associated with a topic detection job. Use this operation
// to get the status of a detection job.
func comprehend_DescribeTopicsDetectionJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DescribeTopicsDetectionJobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendJobId) > 0 {
		input.JobId = aws.String(_comprehendJobId)
	}

	if resp, err := client.DescribeTopicsDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Determines the dominant language of the input text. For a list of languages
// that Amazon Comprehend can detect, see [Amazon Comprehend Supported Languages].
//
// [Amazon Comprehend Supported Languages]: https://docs.aws.amazon.com/comprehend/latest/dg/how-languages.html
func comprehend_DetectDominantLanguage(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DetectDominantLanguageInput{
		// Text: *string, // Required
	}

	if len(_comprehendText) > 0 {
		input.Text = aws.String(_comprehendText)
	}

	if resp, err := client.DetectDominantLanguage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detects named entities in input text when you use the pre-trained model.
// Detects custom entities if you have a custom entity recognition model.
//
// When detecting named entities using the pre-trained model, use plain text as
// the input. For more information about named entities, see [Entities]in the Comprehend
// Developer Guide.
//
// When you use a custom entity recognition model, you can input plain text or you
// can upload a single-page input document (text, PDF, Word, or image).
//
// If the system detects errors while processing a page in the input document, the
// API response includes an entry in Errors for each error.
//
// If the system detects a document-level error in your input document, the API
// returns an InvalidRequestException error response. For details about this
// exception, see [Errors in semi-structured documents]in the Comprehend Developer Guide.
//
// [Errors in semi-structured documents]: https://docs.aws.amazon.com/comprehend/latest/dg/idp-inputs-sync-err.html
// [Entities]: https://docs.aws.amazon.com/comprehend/latest/dg/how-entities.html
func comprehend_DetectEntities(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DetectEntitiesInput{}

	if len(_comprehendBytes) > 0 {
		if err := assignInputField(input, "Bytes", _comprehendBytes); err != nil {
			log.Errorf("invalid --bytes: %s", err.Error())
			return
		}
	}
	if len(_comprehendDocumentReaderConfig) > 0 {
		if err := assignInputField(input, "DocumentReaderConfig", _comprehendDocumentReaderConfig); err != nil {
			log.Errorf("invalid --document-reader-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendEndpointArn) > 0 {
		input.EndpointArn = aws.String(_comprehendEndpointArn)
	}
	if len(_comprehendLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendText) > 0 {
		input.Text = aws.String(_comprehendText)
	}

	if resp, err := client.DetectEntities(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detects the key noun phrases found in the text.
func comprehend_DetectKeyPhrases(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DetectKeyPhrasesInput{
		// LanguageCode: types.LanguageCode, // Required
		// Text: *string, // Required
	}

	if len(_comprehendLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendText) > 0 {
		input.Text = aws.String(_comprehendText)
	}

	if resp, err := client.DetectKeyPhrases(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Inspects the input text for entities that contain personally identifiable
// information (PII) and returns information about them.
func comprehend_DetectPiiEntities(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DetectPiiEntitiesInput{
		// LanguageCode: types.LanguageCode, // Required
		// Text: *string, // Required
	}

	if len(_comprehendLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendText) > 0 {
		input.Text = aws.String(_comprehendText)
	}

	if resp, err := client.DetectPiiEntities(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Inspects text and returns an inference of the prevailing sentiment ( POSITIVE ,
// NEUTRAL , MIXED , or NEGATIVE ).
func comprehend_DetectSentiment(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DetectSentimentInput{
		// LanguageCode: types.LanguageCode, // Required
		// Text: *string, // Required
	}

	if len(_comprehendLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendText) > 0 {
		input.Text = aws.String(_comprehendText)
	}

	if resp, err := client.DetectSentiment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Inspects text for syntax and the part of speech of words in the document. For
// more information, see [Syntax]in the Comprehend Developer Guide.
//
// [Syntax]: https://docs.aws.amazon.com/comprehend/latest/dg/how-syntax.html
func comprehend_DetectSyntax(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DetectSyntaxInput{
		// LanguageCode: types.SyntaxLanguageCode, // Required
		// Text: *string, // Required
	}

	if len(_comprehendLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendText) > 0 {
		input.Text = aws.String(_comprehendText)
	}

	if resp, err := client.DetectSyntax(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Inspects the input text and returns a sentiment analysis for each entity
// identified in the text.
//
// For more information about targeted sentiment, see [Targeted sentiment] in the Amazon Comprehend
// Developer Guide.
//
// [Targeted sentiment]: https://docs.aws.amazon.com/comprehend/latest/dg/how-targeted-sentiment.html
func comprehend_DetectTargetedSentiment(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DetectTargetedSentimentInput{
		// LanguageCode: types.LanguageCode, // Required
		// Text: *string, // Required
	}

	if len(_comprehendLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendText) > 0 {
		input.Text = aws.String(_comprehendText)
	}

	if resp, err := client.DetectTargetedSentiment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Performs toxicity analysis on the list of text strings that you provide as
// input. The API response contains a results list that matches the size of the
// input list. For more information about toxicity detection, see [Toxicity detection]in the Amazon
// Comprehend Developer Guide.
//
// [Toxicity detection]: https://docs.aws.amazon.com/comprehend/latest/dg/toxicity-detection.html
func comprehend_DetectToxicContent(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.DetectToxicContentInput{
		// LanguageCode: types.LanguageCode, // Required
		// TextSegments: []types.TextSegment, // Required
	}

	if len(_comprehendLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendTextSegments) > 0 {
		if err := assignInputField(input, "TextSegments", _comprehendTextSegments); err != nil {
			log.Errorf("invalid --text-segments: %s", err.Error())
			return
		}
	}

	if resp, err := client.DetectToxicContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new custom model that replicates a source custom model that you
// import. The source model can be in your Amazon Web Services account or another
// one.
//
// If the source model is in another Amazon Web Services account, then it must
// have a resource-based policy that authorizes you to import it.
//
// The source model must be in the same Amazon Web Services Region that you're
// using when you import. You can't import a model that's in a different Region.
func comprehend_ImportModel(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.ImportModelInput{
		// SourceModelArn: *string, // Required
	}

	if len(_comprehendSourceModelArn) > 0 {
		input.SourceModelArn = aws.String(_comprehendSourceModelArn)
	}
	if len(_comprehendDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_comprehendDataAccessRoleArn)
	}
	if len(_comprehendModelKmsKeyId) > 0 {
		input.ModelKmsKeyId = aws.String(_comprehendModelKmsKeyId)
	}
	if len(_comprehendModelName) > 0 {
		input.ModelName = aws.String(_comprehendModelName)
	}
	if len(_comprehendTags) > 0 {
		if err := assignInputField(input, "Tags", _comprehendTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_comprehendVersionName) > 0 {
		input.VersionName = aws.String(_comprehendVersionName)
	}

	if resp, err := client.ImportModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List the datasets that you have configured in this Region. For more information
// about datasets, see [Flywheel overview]in the Amazon Comprehend Developer Guide.
//
// [Flywheel overview]: https://docs.aws.amazon.com/comprehend/latest/dg/flywheels-about.html
func comprehend_ListDatasets(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.ListDatasetsInput{}

	if len(_comprehendFilter) > 0 {
		if err := assignInputField(input, "Filter", _comprehendFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_comprehendFlywheelArn) > 0 {
		input.FlywheelArn = aws.String(_comprehendFlywheelArn)
	}
	if len(_comprehendMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _comprehendMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_comprehendNextToken) > 0 {
		input.NextToken = aws.String(_comprehendNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDatasets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*comprehend.ListDatasetsOutput
	p := comprehend.NewListDatasetsPaginator(client, input)
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

// Gets a list of the documentation classification jobs that you have submitted.
func comprehend_ListDocumentClassificationJobs(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.ListDocumentClassificationJobsInput{}

	if len(_comprehendFilter) > 0 {
		if err := assignInputField(input, "Filter", _comprehendFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_comprehendMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _comprehendMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_comprehendNextToken) > 0 {
		input.NextToken = aws.String(_comprehendNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDocumentClassificationJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*comprehend.ListDocumentClassificationJobsOutput
	p := comprehend.NewListDocumentClassificationJobsPaginator(client, input)
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

// Gets a list of summaries of the document classifiers that you have created
func comprehend_ListDocumentClassifierSummaries(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.ListDocumentClassifierSummariesInput{}

	if len(_comprehendMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _comprehendMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_comprehendNextToken) > 0 {
		input.NextToken = aws.String(_comprehendNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDocumentClassifierSummaries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*comprehend.ListDocumentClassifierSummariesOutput
	p := comprehend.NewListDocumentClassifierSummariesPaginator(client, input)
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

// Gets a list of the document classifiers that you have created.
func comprehend_ListDocumentClassifiers(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.ListDocumentClassifiersInput{}

	if len(_comprehendFilter) > 0 {
		if err := assignInputField(input, "Filter", _comprehendFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_comprehendMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _comprehendMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_comprehendNextToken) > 0 {
		input.NextToken = aws.String(_comprehendNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDocumentClassifiers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*comprehend.ListDocumentClassifiersOutput
	p := comprehend.NewListDocumentClassifiersPaginator(client, input)
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

// Gets a list of the dominant language detection jobs that you have submitted.
func comprehend_ListDominantLanguageDetectionJobs(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.ListDominantLanguageDetectionJobsInput{}

	if len(_comprehendFilter) > 0 {
		if err := assignInputField(input, "Filter", _comprehendFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_comprehendMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _comprehendMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_comprehendNextToken) > 0 {
		input.NextToken = aws.String(_comprehendNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDominantLanguageDetectionJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*comprehend.ListDominantLanguageDetectionJobsOutput
	p := comprehend.NewListDominantLanguageDetectionJobsPaginator(client, input)
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

// Gets a list of all existing endpoints that you've created. For information
// about endpoints, see [Managing endpoints].
//
// [Managing endpoints]: https://docs.aws.amazon.com/comprehend/latest/dg/manage-endpoints.html
func comprehend_ListEndpoints(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.ListEndpointsInput{}

	if len(_comprehendFilter) > 0 {
		if err := assignInputField(input, "Filter", _comprehendFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_comprehendMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _comprehendMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_comprehendNextToken) > 0 {
		input.NextToken = aws.String(_comprehendNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEndpoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*comprehend.ListEndpointsOutput
	p := comprehend.NewListEndpointsPaginator(client, input)
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

// Gets a list of the entity detection jobs that you have submitted.
func comprehend_ListEntitiesDetectionJobs(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.ListEntitiesDetectionJobsInput{}

	if len(_comprehendFilter) > 0 {
		if err := assignInputField(input, "Filter", _comprehendFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_comprehendMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _comprehendMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_comprehendNextToken) > 0 {
		input.NextToken = aws.String(_comprehendNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEntitiesDetectionJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*comprehend.ListEntitiesDetectionJobsOutput
	p := comprehend.NewListEntitiesDetectionJobsPaginator(client, input)
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

// Gets a list of summaries for the entity recognizers that you have created.
func comprehend_ListEntityRecognizerSummaries(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.ListEntityRecognizerSummariesInput{}

	if len(_comprehendMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _comprehendMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_comprehendNextToken) > 0 {
		input.NextToken = aws.String(_comprehendNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEntityRecognizerSummaries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*comprehend.ListEntityRecognizerSummariesOutput
	p := comprehend.NewListEntityRecognizerSummariesPaginator(client, input)
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

// Gets a list of the properties of all entity recognizers that you created,
// including recognizers currently in training. Allows you to filter the list of
// recognizers based on criteria such as status and submission time. This call
// returns up to 500 entity recognizers in the list, with a default number of 100
// recognizers in the list.
//
// The results of this list are not in any particular order. Please get the list
// and sort locally if needed.
func comprehend_ListEntityRecognizers(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.ListEntityRecognizersInput{}

	if len(_comprehendFilter) > 0 {
		if err := assignInputField(input, "Filter", _comprehendFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_comprehendMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _comprehendMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_comprehendNextToken) > 0 {
		input.NextToken = aws.String(_comprehendNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEntityRecognizers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*comprehend.ListEntityRecognizersOutput
	p := comprehend.NewListEntityRecognizersPaginator(client, input)
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

// Gets a list of the events detection jobs that you have submitted.
func comprehend_ListEventsDetectionJobs(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.ListEventsDetectionJobsInput{}

	if len(_comprehendFilter) > 0 {
		if err := assignInputField(input, "Filter", _comprehendFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_comprehendMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _comprehendMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_comprehendNextToken) > 0 {
		input.NextToken = aws.String(_comprehendNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEventsDetectionJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*comprehend.ListEventsDetectionJobsOutput
	p := comprehend.NewListEventsDetectionJobsPaginator(client, input)
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

// Information about the history of a flywheel iteration. For more information
// about flywheels, see [Flywheel overview]in the Amazon Comprehend Developer Guide.
//
// [Flywheel overview]: https://docs.aws.amazon.com/comprehend/latest/dg/flywheels-about.html
func comprehend_ListFlywheelIterationHistory(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.ListFlywheelIterationHistoryInput{
		// FlywheelArn: *string, // Required
	}

	if len(_comprehendFlywheelArn) > 0 {
		input.FlywheelArn = aws.String(_comprehendFlywheelArn)
	}
	if len(_comprehendFilter) > 0 {
		if err := assignInputField(input, "Filter", _comprehendFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_comprehendMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _comprehendMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_comprehendNextToken) > 0 {
		input.NextToken = aws.String(_comprehendNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFlywheelIterationHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*comprehend.ListFlywheelIterationHistoryOutput
	p := comprehend.NewListFlywheelIterationHistoryPaginator(client, input)
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

// Gets a list of the flywheels that you have created.
func comprehend_ListFlywheels(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.ListFlywheelsInput{}

	if len(_comprehendFilter) > 0 {
		if err := assignInputField(input, "Filter", _comprehendFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_comprehendMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _comprehendMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_comprehendNextToken) > 0 {
		input.NextToken = aws.String(_comprehendNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFlywheels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*comprehend.ListFlywheelsOutput
	p := comprehend.NewListFlywheelsPaginator(client, input)
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

// Get a list of key phrase detection jobs that you have submitted.
func comprehend_ListKeyPhrasesDetectionJobs(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.ListKeyPhrasesDetectionJobsInput{}

	if len(_comprehendFilter) > 0 {
		if err := assignInputField(input, "Filter", _comprehendFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_comprehendMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _comprehendMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_comprehendNextToken) > 0 {
		input.NextToken = aws.String(_comprehendNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListKeyPhrasesDetectionJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*comprehend.ListKeyPhrasesDetectionJobsOutput
	p := comprehend.NewListKeyPhrasesDetectionJobsPaginator(client, input)
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

// Gets a list of the PII entity detection jobs that you have submitted.
func comprehend_ListPiiEntitiesDetectionJobs(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.ListPiiEntitiesDetectionJobsInput{}

	if len(_comprehendFilter) > 0 {
		if err := assignInputField(input, "Filter", _comprehendFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_comprehendMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _comprehendMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_comprehendNextToken) > 0 {
		input.NextToken = aws.String(_comprehendNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPiiEntitiesDetectionJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*comprehend.ListPiiEntitiesDetectionJobsOutput
	p := comprehend.NewListPiiEntitiesDetectionJobsPaginator(client, input)
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

// Gets a list of sentiment detection jobs that you have submitted.
func comprehend_ListSentimentDetectionJobs(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.ListSentimentDetectionJobsInput{}

	if len(_comprehendFilter) > 0 {
		if err := assignInputField(input, "Filter", _comprehendFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_comprehendMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _comprehendMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_comprehendNextToken) > 0 {
		input.NextToken = aws.String(_comprehendNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSentimentDetectionJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*comprehend.ListSentimentDetectionJobsOutput
	p := comprehend.NewListSentimentDetectionJobsPaginator(client, input)
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

// Lists all tags associated with a given Amazon Comprehend resource.
func comprehend_ListTagsForResource(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_comprehendResourceArn) > 0 {
		input.ResourceArn = aws.String(_comprehendResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of targeted sentiment detection jobs that you have submitted.
func comprehend_ListTargetedSentimentDetectionJobs(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.ListTargetedSentimentDetectionJobsInput{}

	if len(_comprehendFilter) > 0 {
		if err := assignInputField(input, "Filter", _comprehendFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_comprehendMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _comprehendMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_comprehendNextToken) > 0 {
		input.NextToken = aws.String(_comprehendNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTargetedSentimentDetectionJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*comprehend.ListTargetedSentimentDetectionJobsOutput
	p := comprehend.NewListTargetedSentimentDetectionJobsPaginator(client, input)
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

// Gets a list of the topic detection jobs that you have submitted.
func comprehend_ListTopicsDetectionJobs(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.ListTopicsDetectionJobsInput{}

	if len(_comprehendFilter) > 0 {
		if err := assignInputField(input, "Filter", _comprehendFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_comprehendMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _comprehendMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_comprehendNextToken) > 0 {
		input.NextToken = aws.String(_comprehendNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTopicsDetectionJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*comprehend.ListTopicsDetectionJobsOutput
	p := comprehend.NewListTopicsDetectionJobsPaginator(client, input)
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

// Attaches a resource-based policy to a custom model. You can use this policy to
// authorize an entity in another Amazon Web Services account to import the custom
// model, which replicates it in Amazon Comprehend in their account.
func comprehend_PutResourcePolicy(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.PutResourcePolicyInput{
		// ResourceArn: *string, // Required
		// ResourcePolicy: *string, // Required
	}

	if len(_comprehendResourceArn) > 0 {
		input.ResourceArn = aws.String(_comprehendResourceArn)
	}
	if len(_comprehendResourcePolicy) > 0 {
		input.ResourcePolicy = aws.String(_comprehendResourcePolicy)
	}
	if len(_comprehendPolicyRevisionId) > 0 {
		input.PolicyRevisionId = aws.String(_comprehendPolicyRevisionId)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an asynchronous document classification job using a custom
// classification model. Use the DescribeDocumentClassificationJob operation to
// track the progress of the job.
func comprehend_StartDocumentClassificationJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.StartDocumentClassificationJobInput{
		// DataAccessRoleArn: *string, // Required
		// InputDataConfig: *types.InputDataConfig, // Required
		// OutputDataConfig: *types.OutputDataConfig, // Required
	}

	if len(_comprehendDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_comprehendDataAccessRoleArn)
	}
	if len(_comprehendInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _comprehendInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _comprehendOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_comprehendClientRequestToken)
	}
	if len(_comprehendDocumentClassifierArn) > 0 {
		input.DocumentClassifierArn = aws.String(_comprehendDocumentClassifierArn)
	}
	if len(_comprehendFlywheelArn) > 0 {
		input.FlywheelArn = aws.String(_comprehendFlywheelArn)
	}
	if len(_comprehendJobName) > 0 {
		input.JobName = aws.String(_comprehendJobName)
	}
	if len(_comprehendTags) > 0 {
		if err := assignInputField(input, "Tags", _comprehendTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_comprehendVolumeKmsKeyId) > 0 {
		input.VolumeKmsKeyId = aws.String(_comprehendVolumeKmsKeyId)
	}
	if len(_comprehendVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _comprehendVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartDocumentClassificationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an asynchronous dominant language detection job for a collection of
// documents. Use the operation to track the status of a job.
func comprehend_StartDominantLanguageDetectionJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.StartDominantLanguageDetectionJobInput{
		// DataAccessRoleArn: *string, // Required
		// InputDataConfig: *types.InputDataConfig, // Required
		// OutputDataConfig: *types.OutputDataConfig, // Required
	}

	if len(_comprehendDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_comprehendDataAccessRoleArn)
	}
	if len(_comprehendInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _comprehendInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _comprehendOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_comprehendClientRequestToken)
	}
	if len(_comprehendJobName) > 0 {
		input.JobName = aws.String(_comprehendJobName)
	}
	if len(_comprehendTags) > 0 {
		if err := assignInputField(input, "Tags", _comprehendTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_comprehendVolumeKmsKeyId) > 0 {
		input.VolumeKmsKeyId = aws.String(_comprehendVolumeKmsKeyId)
	}
	if len(_comprehendVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _comprehendVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartDominantLanguageDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an asynchronous entity detection job for a collection of documents. Use
// the operation to track the status of a job.
//
// This API can be used for either standard entity detection or custom entity
// recognition. In order to be used for custom entity recognition, the optional
// EntityRecognizerArn must be used in order to provide access to the recognizer
// being used to detect the custom entity.
func comprehend_StartEntitiesDetectionJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.StartEntitiesDetectionJobInput{
		// DataAccessRoleArn: *string, // Required
		// InputDataConfig: *types.InputDataConfig, // Required
		// LanguageCode: types.LanguageCode, // Required
		// OutputDataConfig: *types.OutputDataConfig, // Required
	}

	if len(_comprehendDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_comprehendDataAccessRoleArn)
	}
	if len(_comprehendInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _comprehendInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _comprehendOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_comprehendClientRequestToken)
	}
	if len(_comprehendEntityRecognizerArn) > 0 {
		input.EntityRecognizerArn = aws.String(_comprehendEntityRecognizerArn)
	}
	if len(_comprehendFlywheelArn) > 0 {
		input.FlywheelArn = aws.String(_comprehendFlywheelArn)
	}
	if len(_comprehendJobName) > 0 {
		input.JobName = aws.String(_comprehendJobName)
	}
	if len(_comprehendTags) > 0 {
		if err := assignInputField(input, "Tags", _comprehendTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_comprehendVolumeKmsKeyId) > 0 {
		input.VolumeKmsKeyId = aws.String(_comprehendVolumeKmsKeyId)
	}
	if len(_comprehendVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _comprehendVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartEntitiesDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an asynchronous event detection job for a collection of documents.
func comprehend_StartEventsDetectionJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.StartEventsDetectionJobInput{
		// DataAccessRoleArn: *string, // Required
		// InputDataConfig: *types.InputDataConfig, // Required
		// LanguageCode: types.LanguageCode, // Required
		// OutputDataConfig: *types.OutputDataConfig, // Required
		// TargetEventTypes: []string, // Required
	}

	if len(_comprehendDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_comprehendDataAccessRoleArn)
	}
	if len(_comprehendInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _comprehendInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _comprehendOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendTargetEventTypes) > 0 {
		input.TargetEventTypes = append([]string(nil), _comprehendTargetEventTypes...)
	}
	if len(_comprehendClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_comprehendClientRequestToken)
	}
	if len(_comprehendJobName) > 0 {
		input.JobName = aws.String(_comprehendJobName)
	}
	if len(_comprehendTags) > 0 {
		if err := assignInputField(input, "Tags", _comprehendTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartEventsDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Start the flywheel iteration.This operation uses any new datasets to train a
// new model version. For more information about flywheels, see [Flywheel overview]in the Amazon
// Comprehend Developer Guide.
//
// [Flywheel overview]: https://docs.aws.amazon.com/comprehend/latest/dg/flywheels-about.html
func comprehend_StartFlywheelIteration(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.StartFlywheelIterationInput{
		// FlywheelArn: *string, // Required
	}

	if len(_comprehendFlywheelArn) > 0 {
		input.FlywheelArn = aws.String(_comprehendFlywheelArn)
	}
	if len(_comprehendClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_comprehendClientRequestToken)
	}

	if resp, err := client.StartFlywheelIteration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an asynchronous key phrase detection job for a collection of documents.
// Use the operation to track the status of a job.
func comprehend_StartKeyPhrasesDetectionJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.StartKeyPhrasesDetectionJobInput{
		// DataAccessRoleArn: *string, // Required
		// InputDataConfig: *types.InputDataConfig, // Required
		// LanguageCode: types.LanguageCode, // Required
		// OutputDataConfig: *types.OutputDataConfig, // Required
	}

	if len(_comprehendDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_comprehendDataAccessRoleArn)
	}
	if len(_comprehendInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _comprehendInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _comprehendOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_comprehendClientRequestToken)
	}
	if len(_comprehendJobName) > 0 {
		input.JobName = aws.String(_comprehendJobName)
	}
	if len(_comprehendTags) > 0 {
		if err := assignInputField(input, "Tags", _comprehendTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_comprehendVolumeKmsKeyId) > 0 {
		input.VolumeKmsKeyId = aws.String(_comprehendVolumeKmsKeyId)
	}
	if len(_comprehendVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _comprehendVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartKeyPhrasesDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an asynchronous PII entity detection job for a collection of documents.
func comprehend_StartPiiEntitiesDetectionJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.StartPiiEntitiesDetectionJobInput{
		// DataAccessRoleArn: *string, // Required
		// InputDataConfig: *types.InputDataConfig, // Required
		// LanguageCode: types.LanguageCode, // Required
		// Mode: types.PiiEntitiesDetectionMode, // Required
		// OutputDataConfig: *types.OutputDataConfig, // Required
	}

	if len(_comprehendDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_comprehendDataAccessRoleArn)
	}
	if len(_comprehendInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _comprehendInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendMode) > 0 {
		if err := assignInputField(input, "Mode", _comprehendMode); err != nil {
			log.Errorf("invalid --mode: %s", err.Error())
			return
		}
	}
	if len(_comprehendOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _comprehendOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_comprehendClientRequestToken)
	}
	if len(_comprehendJobName) > 0 {
		input.JobName = aws.String(_comprehendJobName)
	}
	if len(_comprehendRedactionConfig) > 0 {
		if err := assignInputField(input, "RedactionConfig", _comprehendRedactionConfig); err != nil {
			log.Errorf("invalid --redaction-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendTags) > 0 {
		if err := assignInputField(input, "Tags", _comprehendTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartPiiEntitiesDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an asynchronous sentiment detection job for a collection of documents.
// Use the operation to track the status of a job.
func comprehend_StartSentimentDetectionJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.StartSentimentDetectionJobInput{
		// DataAccessRoleArn: *string, // Required
		// InputDataConfig: *types.InputDataConfig, // Required
		// LanguageCode: types.LanguageCode, // Required
		// OutputDataConfig: *types.OutputDataConfig, // Required
	}

	if len(_comprehendDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_comprehendDataAccessRoleArn)
	}
	if len(_comprehendInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _comprehendInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _comprehendOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_comprehendClientRequestToken)
	}
	if len(_comprehendJobName) > 0 {
		input.JobName = aws.String(_comprehendJobName)
	}
	if len(_comprehendTags) > 0 {
		if err := assignInputField(input, "Tags", _comprehendTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_comprehendVolumeKmsKeyId) > 0 {
		input.VolumeKmsKeyId = aws.String(_comprehendVolumeKmsKeyId)
	}
	if len(_comprehendVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _comprehendVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartSentimentDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an asynchronous targeted sentiment detection job for a collection of
// documents. Use the DescribeTargetedSentimentDetectionJob operation to track the
// status of a job.
func comprehend_StartTargetedSentimentDetectionJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.StartTargetedSentimentDetectionJobInput{
		// DataAccessRoleArn: *string, // Required
		// InputDataConfig: *types.InputDataConfig, // Required
		// LanguageCode: types.LanguageCode, // Required
		// OutputDataConfig: *types.OutputDataConfig, // Required
	}

	if len(_comprehendDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_comprehendDataAccessRoleArn)
	}
	if len(_comprehendInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _comprehendInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _comprehendOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_comprehendClientRequestToken)
	}
	if len(_comprehendJobName) > 0 {
		input.JobName = aws.String(_comprehendJobName)
	}
	if len(_comprehendTags) > 0 {
		if err := assignInputField(input, "Tags", _comprehendTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_comprehendVolumeKmsKeyId) > 0 {
		input.VolumeKmsKeyId = aws.String(_comprehendVolumeKmsKeyId)
	}
	if len(_comprehendVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _comprehendVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartTargetedSentimentDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an asynchronous topic detection job. Use the DescribeTopicDetectionJob
// operation to track the status of a job.
func comprehend_StartTopicsDetectionJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.StartTopicsDetectionJobInput{
		// DataAccessRoleArn: *string, // Required
		// InputDataConfig: *types.InputDataConfig, // Required
		// OutputDataConfig: *types.OutputDataConfig, // Required
	}

	if len(_comprehendDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_comprehendDataAccessRoleArn)
	}
	if len(_comprehendInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _comprehendInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _comprehendOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_comprehendClientRequestToken)
	}
	if len(_comprehendJobName) > 0 {
		input.JobName = aws.String(_comprehendJobName)
	}
	if len(_comprehendNumberOfTopics) > 0 {
		if err := assignInputField(input, "NumberOfTopics", _comprehendNumberOfTopics); err != nil {
			log.Errorf("invalid --number-of-topics: %s", err.Error())
			return
		}
	}
	if len(_comprehendTags) > 0 {
		if err := assignInputField(input, "Tags", _comprehendTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_comprehendVolumeKmsKeyId) > 0 {
		input.VolumeKmsKeyId = aws.String(_comprehendVolumeKmsKeyId)
	}
	if len(_comprehendVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _comprehendVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartTopicsDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a dominant language detection job in progress.
// If the job state is IN_PROGRESS the job is marked for termination and put into
// the STOP_REQUESTED state. If the job completes before it can be stopped, it is
// put into the COMPLETED state; otherwise the job is stopped and put into the
// STOPPED state.
//
// If the job is in the COMPLETED or FAILED state when you call the
// StopDominantLanguageDetectionJob operation, the operation returns a 400 Internal
// Request Exception.
//
// When a job is stopped, any documents already processed are written to the
// output location.
func comprehend_StopDominantLanguageDetectionJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.StopDominantLanguageDetectionJobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendJobId) > 0 {
		input.JobId = aws.String(_comprehendJobId)
	}

	if resp, err := client.StopDominantLanguageDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an entities detection job in progress.
// If the job state is IN_PROGRESS the job is marked for termination and put into
// the STOP_REQUESTED state. If the job completes before it can be stopped, it is
// put into the COMPLETED state; otherwise the job is stopped and put into the
// STOPPED state.
//
// If the job is in the COMPLETED or FAILED state when you call the
// StopDominantLanguageDetectionJob operation, the operation returns a 400 Internal
// Request Exception.
//
// When a job is stopped, any documents already processed are written to the
// output location.
func comprehend_StopEntitiesDetectionJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.StopEntitiesDetectionJobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendJobId) > 0 {
		input.JobId = aws.String(_comprehendJobId)
	}

	if resp, err := client.StopEntitiesDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an events detection job in progress.
func comprehend_StopEventsDetectionJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.StopEventsDetectionJobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendJobId) > 0 {
		input.JobId = aws.String(_comprehendJobId)
	}

	if resp, err := client.StopEventsDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a key phrases detection job in progress.
// If the job state is IN_PROGRESS the job is marked for termination and put into
// the STOP_REQUESTED state. If the job completes before it can be stopped, it is
// put into the COMPLETED state; otherwise the job is stopped and put into the
// STOPPED state.
//
// If the job is in the COMPLETED or FAILED state when you call the
// StopDominantLanguageDetectionJob operation, the operation returns a 400 Internal
// Request Exception.
//
// When a job is stopped, any documents already processed are written to the
// output location.
func comprehend_StopKeyPhrasesDetectionJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.StopKeyPhrasesDetectionJobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendJobId) > 0 {
		input.JobId = aws.String(_comprehendJobId)
	}

	if resp, err := client.StopKeyPhrasesDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a PII entities detection job in progress.
func comprehend_StopPiiEntitiesDetectionJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.StopPiiEntitiesDetectionJobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendJobId) > 0 {
		input.JobId = aws.String(_comprehendJobId)
	}

	if resp, err := client.StopPiiEntitiesDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a sentiment detection job in progress.
// If the job state is IN_PROGRESS , the job is marked for termination and put into
// the STOP_REQUESTED state. If the job completes before it can be stopped, it is
// put into the COMPLETED state; otherwise the job is be stopped and put into the
// STOPPED state.
//
// If the job is in the COMPLETED or FAILED state when you call the
// StopDominantLanguageDetectionJob operation, the operation returns a 400 Internal
// Request Exception.
//
// When a job is stopped, any documents already processed are written to the
// output location.
func comprehend_StopSentimentDetectionJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.StopSentimentDetectionJobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendJobId) > 0 {
		input.JobId = aws.String(_comprehendJobId)
	}

	if resp, err := client.StopSentimentDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a targeted sentiment detection job in progress.
// If the job state is IN_PROGRESS , the job is marked for termination and put into
// the STOP_REQUESTED state. If the job completes before it can be stopped, it is
// put into the COMPLETED state; otherwise the job is be stopped and put into the
// STOPPED state.
//
// If the job is in the COMPLETED or FAILED state when you call the
// StopDominantLanguageDetectionJob operation, the operation returns a 400 Internal
// Request Exception.
//
// When a job is stopped, any documents already processed are written to the
// output location.
func comprehend_StopTargetedSentimentDetectionJob(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.StopTargetedSentimentDetectionJobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendJobId) > 0 {
		input.JobId = aws.String(_comprehendJobId)
	}

	if resp, err := client.StopTargetedSentimentDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a document classifier training job while in progress.
// If the training job state is TRAINING , the job is marked for termination and
// put into the STOP_REQUESTED state. If the training job completes before it can
// be stopped, it is put into the TRAINED ; otherwise the training job is stopped
// and put into the STOPPED state and the service sends back an HTTP 200 response
// with an empty HTTP body.
func comprehend_StopTrainingDocumentClassifier(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.StopTrainingDocumentClassifierInput{
		// DocumentClassifierArn: *string, // Required
	}

	if len(_comprehendDocumentClassifierArn) > 0 {
		input.DocumentClassifierArn = aws.String(_comprehendDocumentClassifierArn)
	}

	if resp, err := client.StopTrainingDocumentClassifier(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an entity recognizer training job while in progress.
// If the training job state is TRAINING , the job is marked for termination and
// put into the STOP_REQUESTED state. If the training job completes before it can
// be stopped, it is put into the TRAINED ; otherwise the training job is stopped
// and putted into the STOPPED state and the service sends back an HTTP 200
// response with an empty HTTP body.
func comprehend_StopTrainingEntityRecognizer(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.StopTrainingEntityRecognizerInput{
		// EntityRecognizerArn: *string, // Required
	}

	if len(_comprehendEntityRecognizerArn) > 0 {
		input.EntityRecognizerArn = aws.String(_comprehendEntityRecognizerArn)
	}

	if resp, err := client.StopTrainingEntityRecognizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a specific tag with an Amazon Comprehend resource. A tag is a
// key-value pair that adds as a metadata to a resource used by Amazon Comprehend.
// For example, a tag with "Sales" as the key might be added to a resource to
// indicate its use by the sales department.
func comprehend_TagResource(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_comprehendResourceArn) > 0 {
		input.ResourceArn = aws.String(_comprehendResourceArn)
	}
	if len(_comprehendTags) > 0 {
		if err := assignInputField(input, "Tags", _comprehendTags); err != nil {
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

// Removes a specific tag associated with an Amazon Comprehend resource.
func comprehend_UntagResource(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_comprehendResourceArn) > 0 {
		input.ResourceArn = aws.String(_comprehendResourceArn)
	}
	if len(_comprehendTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _comprehendTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates information about the specified endpoint. For information about
// endpoints, see [Managing endpoints].
//
// [Managing endpoints]: https://docs.aws.amazon.com/comprehend/latest/dg/manage-endpoints.html
func comprehend_UpdateEndpoint(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.UpdateEndpointInput{
		// EndpointArn: *string, // Required
	}

	if len(_comprehendEndpointArn) > 0 {
		input.EndpointArn = aws.String(_comprehendEndpointArn)
	}
	if len(_comprehendDesiredDataAccessRoleArn) > 0 {
		input.DesiredDataAccessRoleArn = aws.String(_comprehendDesiredDataAccessRoleArn)
	}
	if len(_comprehendDesiredInferenceUnits) > 0 {
		if err := assignInputField(input, "DesiredInferenceUnits", _comprehendDesiredInferenceUnits); err != nil {
			log.Errorf("invalid --desired-inference-units: %s", err.Error())
			return
		}
	}
	if len(_comprehendDesiredModelArn) > 0 {
		input.DesiredModelArn = aws.String(_comprehendDesiredModelArn)
	}
	if len(_comprehendFlywheelArn) > 0 {
		input.FlywheelArn = aws.String(_comprehendFlywheelArn)
	}

	if resp, err := client.UpdateEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the configuration information for an existing flywheel.
func comprehend_UpdateFlywheel(cfg aws.Config, client *comprehend.Client) {
	input := &comprehend.UpdateFlywheelInput{
		// FlywheelArn: *string, // Required
	}

	if len(_comprehendFlywheelArn) > 0 {
		input.FlywheelArn = aws.String(_comprehendFlywheelArn)
	}
	if len(_comprehendActiveModelArn) > 0 {
		input.ActiveModelArn = aws.String(_comprehendActiveModelArn)
	}
	if len(_comprehendDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_comprehendDataAccessRoleArn)
	}
	if len(_comprehendDataSecurityConfig) > 0 {
		if err := assignInputField(input, "DataSecurityConfig", _comprehendDataSecurityConfig); err != nil {
			log.Errorf("invalid --data-security-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFlywheel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_comprehendCmd)
	_comprehendCmd.Flags().SortFlags = false

	_comprehendCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_comprehendCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_comprehendCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_comprehendCmd.Flags().StringVarP(&_comprehendActiveModelArn, "active-model-arn", "", "", "Active Model ARN")
	_comprehendCmd.Flags().StringVarP(&_comprehendBytes, "bytes", "", "", "Bytes")
	_comprehendCmd.Flags().StringVarP(&_comprehendClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_comprehendCmd.Flags().StringVarP(&_comprehendDataAccessRoleArn, "data-access-role-arn", "", "", "Data Access Role ARN")
	_comprehendCmd.Flags().StringVarP(&_comprehendDataLakeS3Uri, "data-lake-s3-uri", "", "", "Data Lake S3 URI")
	_comprehendCmd.Flags().StringVarP(&_comprehendDataSecurityConfig, "data-security-config", "", "", "Data Security Config")
	_comprehendCmd.Flags().StringVarP(&_comprehendDatasetArn, "dataset-arn", "", "", "Dataset ARN")
	_comprehendCmd.Flags().StringVarP(&_comprehendDatasetName, "dataset-name", "", "", "Dataset Name")
	_comprehendCmd.Flags().StringVarP(&_comprehendDatasetType, "dataset-type", "", "", "Dataset Type")
	_comprehendCmd.Flags().StringVarP(&_comprehendDescription, "description", "", "", "Description")
	_comprehendCmd.Flags().StringVarP(&_comprehendDesiredDataAccessRoleArn, "desired-data-access-role-arn", "", "", "Desired Data Access Role ARN")
	_comprehendCmd.Flags().StringVarP(&_comprehendDesiredInferenceUnits, "desired-inference-units", "", "", "Desired Inference Units")
	_comprehendCmd.Flags().StringVarP(&_comprehendDesiredModelArn, "desired-model-arn", "", "", "Desired Model ARN")
	_comprehendCmd.Flags().StringVarP(&_comprehendDocumentClassifierArn, "document-classifier-arn", "", "", "Document Classifier ARN")
	_comprehendCmd.Flags().StringVarP(&_comprehendDocumentClassifierName, "document-classifier-name", "", "", "Document Classifier Name")
	_comprehendCmd.Flags().StringVarP(&_comprehendDocumentReaderConfig, "document-reader-config", "", "", "Document Reader Config")
	_comprehendCmd.Flags().StringVarP(&_comprehendEndpointArn, "endpoint-arn", "", "", "Endpoint ARN")
	_comprehendCmd.Flags().StringVarP(&_comprehendEndpointName, "endpoint-name", "", "", "Endpoint Name")
	_comprehendCmd.Flags().StringVarP(&_comprehendEntityRecognizerArn, "entity-recognizer-arn", "", "", "Entity Recognizer ARN")
	_comprehendCmd.Flags().StringVarP(&_comprehendFilter, "filter", "", "", "Filter")
	_comprehendCmd.Flags().StringVarP(&_comprehendFlywheelArn, "flywheel-arn", "", "", "Flywheel ARN")
	_comprehendCmd.Flags().StringVarP(&_comprehendFlywheelIterationId, "flywheel-iteration-id", "", "", "Flywheel Iteration ID")
	_comprehendCmd.Flags().StringVarP(&_comprehendFlywheelName, "flywheel-name", "", "", "Flywheel Name")
	_comprehendCmd.Flags().StringVarP(&_comprehendInputDataConfig, "input-data-config", "", "", "Input Data Config")
	_comprehendCmd.Flags().StringVarP(&_comprehendJobId, "job-id", "", "", "Job ID")
	_comprehendCmd.Flags().StringVarP(&_comprehendJobName, "job-name", "", "", "Job Name")
	_comprehendCmd.Flags().StringVarP(&_comprehendLanguageCode, "language-code", "", "", "Language Code")
	_comprehendCmd.Flags().StringVarP(&_comprehendMaxResults, "max-results", "", "", "Max Results")
	_comprehendCmd.Flags().StringVarP(&_comprehendMode, "mode", "", "", "Mode")
	_comprehendCmd.Flags().StringVarP(&_comprehendModelArn, "model-arn", "", "", "Model ARN")
	_comprehendCmd.Flags().StringVarP(&_comprehendModelKmsKeyId, "model-kms-key-id", "", "", "Model KMS Key ID")
	_comprehendCmd.Flags().StringVarP(&_comprehendModelName, "model-name", "", "", "Model Name")
	_comprehendCmd.Flags().StringVarP(&_comprehendModelPolicy, "model-policy", "", "", "Model Policy")
	_comprehendCmd.Flags().StringVarP(&_comprehendModelType, "model-type", "", "", "Model Type")
	_comprehendCmd.Flags().StringVarP(&_comprehendNextToken, "next-token", "", "", "Next Token")
	_comprehendCmd.Flags().StringVarP(&_comprehendNumberOfTopics, "number-of-topics", "", "", "Number Of Topics")
	_comprehendCmd.Flags().StringVarP(&_comprehendOutputDataConfig, "output-data-config", "", "", "Output Data Config")
	_comprehendCmd.Flags().StringVarP(&_comprehendPolicyRevisionId, "policy-revision-id", "", "", "Policy Revision ID")
	_comprehendCmd.Flags().StringVarP(&_comprehendRecognizerName, "recognizer-name", "", "", "Recognizer Name")
	_comprehendCmd.Flags().StringVarP(&_comprehendRedactionConfig, "redaction-config", "", "", "Redaction Config")
	_comprehendCmd.Flags().StringVarP(&_comprehendResourceArn, "resource-arn", "", "", "Resource ARN")
	_comprehendCmd.Flags().StringVarP(&_comprehendResourcePolicy, "resource-policy", "", "", "Resource Policy")
	_comprehendCmd.Flags().StringVarP(&_comprehendSourceModelArn, "source-model-arn", "", "", "Source Model ARN")
	_comprehendCmd.Flags().StringSliceVarP(&_comprehendTagKeys, "tag-keys", "", nil, "Tag Keys")
	_comprehendCmd.Flags().StringVarP(&_comprehendTags, "tags", "", "", "Tags")
	_comprehendCmd.Flags().StringSliceVarP(&_comprehendTargetEventTypes, "target-event-types", "", nil, "Target Event Types")
	_comprehendCmd.Flags().StringVarP(&_comprehendTaskConfig, "task-config", "", "", "Task Config")
	_comprehendCmd.Flags().StringVarP(&_comprehendText, "text", "", "", "Text")
	_comprehendCmd.Flags().StringSliceVarP(&_comprehendTextList, "text-list", "", nil, "Text List")
	_comprehendCmd.Flags().StringVarP(&_comprehendTextSegments, "text-segments", "", "", "Text Segments")
	_comprehendCmd.Flags().StringVarP(&_comprehendVersionName, "version-name", "", "", "Version Name")
	_comprehendCmd.Flags().StringVarP(&_comprehendVolumeKmsKeyId, "volume-kms-key-id", "", "", "Volume KMS Key ID")
	_comprehendCmd.Flags().StringVarP(&_comprehendVpcConfig, "vpc-config", "", "", "VPC Config")

	_comprehendCmd.Flags().BoolVarP(&_comprehendBatchDetectDominantLanguage, "batch-detect-dominant-language", "", false, "Batch Detect Dominant Language")
	_comprehendCmd.Flags().BoolVarP(&_comprehendBatchDetectEntities, "batch-detect-entities", "", false, "Batch Detect Entities")
	_comprehendCmd.Flags().BoolVarP(&_comprehendBatchDetectKeyPhrases, "batch-detect-key-phrases", "", false, "Batch Detect Key Phrases")
	_comprehendCmd.Flags().BoolVarP(&_comprehendBatchDetectSentiment, "batch-detect-sentiment", "", false, "Batch Detect Sentiment")
	_comprehendCmd.Flags().BoolVarP(&_comprehendBatchDetectSyntax, "batch-detect-syntax", "", false, "Batch Detect Syntax")
	_comprehendCmd.Flags().BoolVarP(&_comprehendBatchDetectTargetedSentiment, "batch-detect-targeted-sentiment", "", false, "Batch Detect Targeted Sentiment")
	_comprehendCmd.Flags().BoolVarP(&_comprehendClassifyDocument, "classify-document", "", false, "Classify Document")
	_comprehendCmd.Flags().BoolVarP(&_comprehendContainsPiiEntities, "contains-pii-entities", "", false, "Contains Pii Entities")
	_comprehendCmd.Flags().BoolVarP(&_comprehendCreateDataset, "create-dataset", "", false, "Create Dataset")
	_comprehendCmd.Flags().BoolVarP(&_comprehendCreateDocumentClassifier, "create-document-classifier", "", false, "Create Document Classifier")
	_comprehendCmd.Flags().BoolVarP(&_comprehendCreateEndpoint, "create-endpoint", "", false, "Create Endpoint")
	_comprehendCmd.Flags().BoolVarP(&_comprehendCreateEntityRecognizer, "create-entity-recognizer", "", false, "Create Entity Recognizer")
	_comprehendCmd.Flags().BoolVarP(&_comprehendCreateFlywheel, "create-flywheel", "", false, "Create Flywheel")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDeleteDocumentClassifier, "delete-document-classifier", "", false, "Delete Document Classifier")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDeleteEndpoint, "delete-endpoint", "", false, "Delete Endpoint")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDeleteEntityRecognizer, "delete-entity-recognizer", "", false, "Delete Entity Recognizer")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDeleteFlywheel, "delete-flywheel", "", false, "Delete Flywheel")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDescribeDataset, "describe-dataset", "", false, "Describe Dataset")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDescribeDocumentClassificationJob, "describe-document-classification-job", "", false, "Describe Document Classification Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDescribeDocumentClassifier, "describe-document-classifier", "", false, "Describe Document Classifier")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDescribeDominantLanguageDetectionJob, "describe-dominant-language-detection-job", "", false, "Describe Dominant Language Detection Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDescribeEndpoint, "describe-endpoint", "", false, "Describe Endpoint")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDescribeEntitiesDetectionJob, "describe-entities-detection-job", "", false, "Describe Entities Detection Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDescribeEntityRecognizer, "describe-entity-recognizer", "", false, "Describe Entity Recognizer")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDescribeEventsDetectionJob, "describe-events-detection-job", "", false, "Describe Events Detection Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDescribeFlywheel, "describe-flywheel", "", false, "Describe Flywheel")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDescribeFlywheelIteration, "describe-flywheel-iteration", "", false, "Describe Flywheel Iteration")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDescribeKeyPhrasesDetectionJob, "describe-key-phrases-detection-job", "", false, "Describe Key Phrases Detection Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDescribePiiEntitiesDetectionJob, "describe-pii-entities-detection-job", "", false, "Describe Pii Entities Detection Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDescribeResourcePolicy, "describe-resource-policy", "", false, "Describe Resource Policy")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDescribeSentimentDetectionJob, "describe-sentiment-detection-job", "", false, "Describe Sentiment Detection Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDescribeTargetedSentimentDetectionJob, "describe-targeted-sentiment-detection-job", "", false, "Describe Targeted Sentiment Detection Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDescribeTopicsDetectionJob, "describe-topics-detection-job", "", false, "Describe Topics Detection Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDetectDominantLanguage, "detect-dominant-language", "", false, "Detect Dominant Language")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDetectEntities, "detect-entities", "", false, "Detect Entities")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDetectKeyPhrases, "detect-key-phrases", "", false, "Detect Key Phrases")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDetectPiiEntities, "detect-pii-entities", "", false, "Detect Pii Entities")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDetectSentiment, "detect-sentiment", "", false, "Detect Sentiment")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDetectSyntax, "detect-syntax", "", false, "Detect Syntax")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDetectTargetedSentiment, "detect-targeted-sentiment", "", false, "Detect Targeted Sentiment")
	_comprehendCmd.Flags().BoolVarP(&_comprehendDetectToxicContent, "detect-toxic-content", "", false, "Detect Toxic Content")
	_comprehendCmd.Flags().BoolVarP(&_comprehendImportModel, "import-model", "", false, "Import Model")
	_comprehendCmd.Flags().BoolVarP(&_comprehendListDatasets, "list-datasets", "", false, "List Datasets")
	_comprehendCmd.Flags().BoolVarP(&_comprehendListDocumentClassificationJobs, "list-document-classification-jobs", "", false, "List Document Classification Jobs")
	_comprehendCmd.Flags().BoolVarP(&_comprehendListDocumentClassifierSummaries, "list-document-classifier-summaries", "", false, "List Document Classifier Summaries")
	_comprehendCmd.Flags().BoolVarP(&_comprehendListDocumentClassifiers, "list-document-classifiers", "", false, "List Document Classifiers")
	_comprehendCmd.Flags().BoolVarP(&_comprehendListDominantLanguageDetectionJobs, "list-dominant-language-detection-jobs", "", false, "List Dominant Language Detection Jobs")
	_comprehendCmd.Flags().BoolVarP(&_comprehendListEndpoints, "list-endpoints", "", false, "List Endpoints")
	_comprehendCmd.Flags().BoolVarP(&_comprehendListEntitiesDetectionJobs, "list-entities-detection-jobs", "", false, "List Entities Detection Jobs")
	_comprehendCmd.Flags().BoolVarP(&_comprehendListEntityRecognizerSummaries, "list-entity-recognizer-summaries", "", false, "List Entity Recognizer Summaries")
	_comprehendCmd.Flags().BoolVarP(&_comprehendListEntityRecognizers, "list-entity-recognizers", "", false, "List Entity Recognizers")
	_comprehendCmd.Flags().BoolVarP(&_comprehendListEventsDetectionJobs, "list-events-detection-jobs", "", false, "List Events Detection Jobs")
	_comprehendCmd.Flags().BoolVarP(&_comprehendListFlywheelIterationHistory, "list-flywheel-iteration-history", "", false, "List Flywheel Iteration History")
	_comprehendCmd.Flags().BoolVarP(&_comprehendListFlywheels, "list-flywheels", "", false, "List Flywheels")
	_comprehendCmd.Flags().BoolVarP(&_comprehendListKeyPhrasesDetectionJobs, "list-key-phrases-detection-jobs", "", false, "List Key Phrases Detection Jobs")
	_comprehendCmd.Flags().BoolVarP(&_comprehendListPiiEntitiesDetectionJobs, "list-pii-entities-detection-jobs", "", false, "List Pii Entities Detection Jobs")
	_comprehendCmd.Flags().BoolVarP(&_comprehendListSentimentDetectionJobs, "list-sentiment-detection-jobs", "", false, "List Sentiment Detection Jobs")
	_comprehendCmd.Flags().BoolVarP(&_comprehendListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_comprehendCmd.Flags().BoolVarP(&_comprehendListTargetedSentimentDetectionJobs, "list-targeted-sentiment-detection-jobs", "", false, "List Targeted Sentiment Detection Jobs")
	_comprehendCmd.Flags().BoolVarP(&_comprehendListTopicsDetectionJobs, "list-topics-detection-jobs", "", false, "List Topics Detection Jobs")
	_comprehendCmd.Flags().BoolVarP(&_comprehendPutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_comprehendCmd.Flags().BoolVarP(&_comprehendStartDocumentClassificationJob, "start-document-classification-job", "", false, "Start Document Classification Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendStartDominantLanguageDetectionJob, "start-dominant-language-detection-job", "", false, "Start Dominant Language Detection Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendStartEntitiesDetectionJob, "start-entities-detection-job", "", false, "Start Entities Detection Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendStartEventsDetectionJob, "start-events-detection-job", "", false, "Start Events Detection Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendStartFlywheelIteration, "start-flywheel-iteration", "", false, "Start Flywheel Iteration")
	_comprehendCmd.Flags().BoolVarP(&_comprehendStartKeyPhrasesDetectionJob, "start-key-phrases-detection-job", "", false, "Start Key Phrases Detection Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendStartPiiEntitiesDetectionJob, "start-pii-entities-detection-job", "", false, "Start Pii Entities Detection Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendStartSentimentDetectionJob, "start-sentiment-detection-job", "", false, "Start Sentiment Detection Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendStartTargetedSentimentDetectionJob, "start-targeted-sentiment-detection-job", "", false, "Start Targeted Sentiment Detection Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendStartTopicsDetectionJob, "start-topics-detection-job", "", false, "Start Topics Detection Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendStopDominantLanguageDetectionJob, "stop-dominant-language-detection-job", "", false, "Stop Dominant Language Detection Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendStopEntitiesDetectionJob, "stop-entities-detection-job", "", false, "Stop Entities Detection Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendStopEventsDetectionJob, "stop-events-detection-job", "", false, "Stop Events Detection Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendStopKeyPhrasesDetectionJob, "stop-key-phrases-detection-job", "", false, "Stop Key Phrases Detection Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendStopPiiEntitiesDetectionJob, "stop-pii-entities-detection-job", "", false, "Stop Pii Entities Detection Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendStopSentimentDetectionJob, "stop-sentiment-detection-job", "", false, "Stop Sentiment Detection Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendStopTargetedSentimentDetectionJob, "stop-targeted-sentiment-detection-job", "", false, "Stop Targeted Sentiment Detection Job")
	_comprehendCmd.Flags().BoolVarP(&_comprehendStopTrainingDocumentClassifier, "stop-training-document-classifier", "", false, "Stop Training Document Classifier")
	_comprehendCmd.Flags().BoolVarP(&_comprehendStopTrainingEntityRecognizer, "stop-training-entity-recognizer", "", false, "Stop Training Entity Recognizer")
	_comprehendCmd.Flags().BoolVarP(&_comprehendTagResource, "tag-resource", "", false, "Tag Resource")
	_comprehendCmd.Flags().BoolVarP(&_comprehendUntagResource, "untag-resource", "", false, "Untag Resource")
	_comprehendCmd.Flags().BoolVarP(&_comprehendUpdateEndpoint, "update-endpoint", "", false, "Update Endpoint")
	_comprehendCmd.Flags().BoolVarP(&_comprehendUpdateFlywheel, "update-flywheel", "", false, "Update Flywheel")

}
