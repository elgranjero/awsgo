package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/textract"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// textractCmd represents the textract command
var _textractCmd = &cobra.Command{
	Use:   "textract",
	Short: "AWS textract CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := textract.NewFromConfig(cfg)
		if _textractAnalyzeDocument {
			textract_AnalyzeDocument(cfg, client)
			return
		}
		if _textractAnalyzeExpense {
			textract_AnalyzeExpense(cfg, client)
			return
		}
		if _textractAnalyzeID {
			textract_AnalyzeID(cfg, client)
			return
		}
		if _textractCreateAdapter {
			textract_CreateAdapter(cfg, client)
			return
		}
		if _textractCreateAdapterVersion {
			textract_CreateAdapterVersion(cfg, client)
			return
		}
		if _textractDeleteAdapter {
			textract_DeleteAdapter(cfg, client)
			return
		}
		if _textractDeleteAdapterVersion {
			textract_DeleteAdapterVersion(cfg, client)
			return
		}
		if _textractDetectDocumentText {
			textract_DetectDocumentText(cfg, client)
			return
		}
		if _textractGetAdapter {
			textract_GetAdapter(cfg, client)
			return
		}
		if _textractGetAdapterVersion {
			textract_GetAdapterVersion(cfg, client)
			return
		}
		if _textractGetDocumentAnalysis {
			textract_GetDocumentAnalysis(cfg, client)
			return
		}
		if _textractGetDocumentTextDetection {
			textract_GetDocumentTextDetection(cfg, client)
			return
		}
		if _textractGetExpenseAnalysis {
			textract_GetExpenseAnalysis(cfg, client)
			return
		}
		if _textractGetLendingAnalysis {
			textract_GetLendingAnalysis(cfg, client)
			return
		}
		if _textractGetLendingAnalysisSummary {
			textract_GetLendingAnalysisSummary(cfg, client)
			return
		}
		if _textractListAdapterVersions {
			textract_ListAdapterVersions(cfg, client)
			return
		}
		if _textractListAdapters {
			textract_ListAdapters(cfg, client)
			return
		}
		if _textractListTagsForResource {
			textract_ListTagsForResource(cfg, client)
			return
		}
		if _textractStartDocumentAnalysis {
			textract_StartDocumentAnalysis(cfg, client)
			return
		}
		if _textractStartDocumentTextDetection {
			textract_StartDocumentTextDetection(cfg, client)
			return
		}
		if _textractStartExpenseAnalysis {
			textract_StartExpenseAnalysis(cfg, client)
			return
		}
		if _textractStartLendingAnalysis {
			textract_StartLendingAnalysis(cfg, client)
			return
		}
		if _textractTagResource {
			textract_TagResource(cfg, client)
			return
		}
		if _textractUntagResource {
			textract_UntagResource(cfg, client)
			return
		}
		if _textractUpdateAdapter {
			textract_UpdateAdapter(cfg, client)
			return
		}

	},
}

var (
	_textractAnalyzeDocument            bool
	_textractAnalyzeExpense             bool
	_textractAnalyzeID                  bool
	_textractCreateAdapter              bool
	_textractCreateAdapterVersion       bool
	_textractDeleteAdapter              bool
	_textractDeleteAdapterVersion       bool
	_textractDetectDocumentText         bool
	_textractGetAdapter                 bool
	_textractGetAdapterVersion          bool
	_textractGetDocumentAnalysis        bool
	_textractGetDocumentTextDetection   bool
	_textractGetExpenseAnalysis         bool
	_textractGetLendingAnalysis         bool
	_textractGetLendingAnalysisSummary  bool
	_textractListAdapterVersions        bool
	_textractListAdapters               bool
	_textractListTagsForResource        bool
	_textractStartDocumentAnalysis      bool
	_textractStartDocumentTextDetection bool
	_textractStartExpenseAnalysis       bool
	_textractStartLendingAnalysis       bool
	_textractTagResource                bool
	_textractUntagResource              bool
	_textractUpdateAdapter              bool

	_textractAdapterId           string
	_textractAdapterName         string
	_textractAdapterVersion      string
	_textractAdaptersConfig      string
	_textractAfterCreationTime   string
	_textractAutoUpdate          string
	_textractBeforeCreationTime  string
	_textractClientRequestToken  string
	_textractDatasetConfig       string
	_textractDescription         string
	_textractDocument            string
	_textractDocumentLocation    string
	_textractDocumentPages       string
	_textractFeatureTypes        string
	_textractHumanLoopConfig     string
	_textractJobId               string
	_textractJobTag              string
	_textractKMSKeyId            string
	_textractMaxResults          string
	_textractNextToken           string
	_textractNotificationChannel string
	_textractOutputConfig        string
	_textractQueriesConfig       string
	_textractResourceARN         string
	_textractTagKeys             []string
	_textractTags                string
)

// Analyzes an input document for relationships between detected items.
// The types of information returned are as follows:
//
// - Form data (key-value pairs). The related information is returned in two Block
// objects, each of type KEY_VALUE_SET : a KEY Block object and a VALUE Block
// object. For example, Name: Ana Silva Carolina contains a key and value. Name: is
// the key. Ana Silva Carolina is the value.
//
// - Table and table cell data. A TABLE Block object contains information about a
// detected table. A CELL Block object is returned for each cell in a table.
//
// - Lines and words of text. A LINE Block object contains one or more WORD Block
// objects. All lines and words that are detected in the document are returned
// (including text that doesn't have a relationship with the value of
// FeatureTypes ).
//
// - Signatures. A SIGNATURE Block object contains the location information of a
// signature in a document. If used in conjunction with forms or tables, a
// signature can be given a Key-Value pairing or be detected in the cell of a
// table.
//
// - Query. A QUERY Block object contains the query text, alias and link to the
// associated Query results block object.
//
// - Query Result. A QUERY_RESULT Block object contains the answer to the query
// and an ID that connects it to the query asked. This Block also contains a
// confidence score.
//
// Selection elements such as check boxes and option buttons (radio buttons) can
// be detected in form data and in tables. A SELECTION_ELEMENT Block object
// contains information about a selection element, including the selection status.
//
// You can choose which type of analysis to perform by specifying the FeatureTypes
// list.
//
// The output is returned in a list of Block objects.
//
// AnalyzeDocument is a synchronous operation. To analyze documents
// asynchronously, use StartDocumentAnalysis.
//
// For more information, see [Document Text Analysis].
//
// [Document Text Analysis]: https://docs.aws.amazon.com/textract/latest/dg/how-it-works-analyzing.html
func textract_AnalyzeDocument(cfg aws.Config, client *textract.Client) {
	input := &textract.AnalyzeDocumentInput{
		// Document: *types.Document, // Required
		// FeatureTypes: []types.FeatureType, // Required
	}

	if len(_textractDocument) > 0 {
		if err := assignInputField(input, "Document", _textractDocument); err != nil {
			log.Errorf("invalid --document: %s", err.Error())
			return
		}
	}
	if len(_textractFeatureTypes) > 0 {
		if err := assignInputField(input, "FeatureTypes", _textractFeatureTypes); err != nil {
			log.Errorf("invalid --feature-types: %s", err.Error())
			return
		}
	}
	if len(_textractAdaptersConfig) > 0 {
		if err := assignInputField(input, "AdaptersConfig", _textractAdaptersConfig); err != nil {
			log.Errorf("invalid --adapters-config: %s", err.Error())
			return
		}
	}
	if len(_textractHumanLoopConfig) > 0 {
		if err := assignInputField(input, "HumanLoopConfig", _textractHumanLoopConfig); err != nil {
			log.Errorf("invalid --human-loop-config: %s", err.Error())
			return
		}
	}
	if len(_textractQueriesConfig) > 0 {
		if err := assignInputField(input, "QueriesConfig", _textractQueriesConfig); err != nil {
			log.Errorf("invalid --queries-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.AnalyzeDocument(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// AnalyzeExpense synchronously analyzes an input document for financially related
// relationships between text.
//
// Information is returned as ExpenseDocuments and seperated as follows:
//
// - LineItemGroups - A data set containing LineItems which store information
// about the lines of text, such as an item purchased and its price on a receipt.
//
// - SummaryFields - Contains all other information a receipt, such as header
// information or the vendors name.
func textract_AnalyzeExpense(cfg aws.Config, client *textract.Client) {
	input := &textract.AnalyzeExpenseInput{
		// Document: *types.Document, // Required
	}

	if len(_textractDocument) > 0 {
		if err := assignInputField(input, "Document", _textractDocument); err != nil {
			log.Errorf("invalid --document: %s", err.Error())
			return
		}
	}

	if resp, err := client.AnalyzeExpense(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Analyzes identity documents for relevant information. This information is
// extracted and returned as IdentityDocumentFields , which records both the
// normalized field and value of the extracted text. Unlike other Amazon Textract
// operations, AnalyzeID doesn't return any Geometry data.
func textract_AnalyzeID(cfg aws.Config, client *textract.Client) {
	input := &textract.AnalyzeIDInput{
		// DocumentPages: []types.Document, // Required
	}

	if len(_textractDocumentPages) > 0 {
		if err := assignInputField(input, "DocumentPages", _textractDocumentPages); err != nil {
			log.Errorf("invalid --document-pages: %s", err.Error())
			return
		}
	}

	if resp, err := client.AnalyzeID(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an adapter, which can be fine-tuned for enhanced performance on user
// provided documents. Takes an AdapterName and FeatureType. Currently the only
// supported feature type is QUERIES . You can also provide a Description, Tags,
// and a ClientRequestToken. You can choose whether or not the adapter should be
// AutoUpdated with the AutoUpdate argument. By default, AutoUpdate is set to
// DISABLED.
func textract_CreateAdapter(cfg aws.Config, client *textract.Client) {
	input := &textract.CreateAdapterInput{
		// AdapterName: *string, // Required
		// FeatureTypes: []types.FeatureType, // Required
	}

	if len(_textractAdapterName) > 0 {
		input.AdapterName = aws.String(_textractAdapterName)
	}
	if len(_textractFeatureTypes) > 0 {
		if err := assignInputField(input, "FeatureTypes", _textractFeatureTypes); err != nil {
			log.Errorf("invalid --feature-types: %s", err.Error())
			return
		}
	}
	if len(_textractAutoUpdate) > 0 {
		if err := assignInputField(input, "AutoUpdate", _textractAutoUpdate); err != nil {
			log.Errorf("invalid --auto-update: %s", err.Error())
			return
		}
	}
	if len(_textractClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_textractClientRequestToken)
	}
	if len(_textractDescription) > 0 {
		input.Description = aws.String(_textractDescription)
	}
	if len(_textractTags) > 0 {
		if err := assignInputField(input, "Tags", _textractTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAdapter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new version of an adapter. Operates on a provided AdapterId and a
// specified dataset provided via the DatasetConfig argument. Requires that you
// specify an Amazon S3 bucket with the OutputConfig argument. You can provide an
// optional KMSKeyId, an optional ClientRequestToken, and optional tags.
func textract_CreateAdapterVersion(cfg aws.Config, client *textract.Client) {
	input := &textract.CreateAdapterVersionInput{
		// AdapterId: *string, // Required
		// DatasetConfig: *types.AdapterVersionDatasetConfig, // Required
		// OutputConfig: *types.OutputConfig, // Required
	}

	if len(_textractAdapterId) > 0 {
		input.AdapterId = aws.String(_textractAdapterId)
	}
	if len(_textractDatasetConfig) > 0 {
		if err := assignInputField(input, "DatasetConfig", _textractDatasetConfig); err != nil {
			log.Errorf("invalid --dataset-config: %s", err.Error())
			return
		}
	}
	if len(_textractOutputConfig) > 0 {
		if err := assignInputField(input, "OutputConfig", _textractOutputConfig); err != nil {
			log.Errorf("invalid --output-config: %s", err.Error())
			return
		}
	}
	if len(_textractClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_textractClientRequestToken)
	}
	if len(_textractKMSKeyId) > 0 {
		input.KMSKeyId = aws.String(_textractKMSKeyId)
	}
	if len(_textractTags) > 0 {
		if err := assignInputField(input, "Tags", _textractTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAdapterVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Textract adapter. Takes an AdapterId and deletes the adapter
// specified by the ID.
func textract_DeleteAdapter(cfg aws.Config, client *textract.Client) {
	input := &textract.DeleteAdapterInput{
		// AdapterId: *string, // Required
	}

	if len(_textractAdapterId) > 0 {
		input.AdapterId = aws.String(_textractAdapterId)
	}

	if resp, err := client.DeleteAdapter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Textract adapter version. Requires that you specify both an
// AdapterId and a AdapterVersion. Deletes the adapter version specified by the
// AdapterId and the AdapterVersion.
func textract_DeleteAdapterVersion(cfg aws.Config, client *textract.Client) {
	input := &textract.DeleteAdapterVersionInput{
		// AdapterId: *string, // Required
		// AdapterVersion: *string, // Required
	}

	if len(_textractAdapterId) > 0 {
		input.AdapterId = aws.String(_textractAdapterId)
	}
	if len(_textractAdapterVersion) > 0 {
		input.AdapterVersion = aws.String(_textractAdapterVersion)
	}

	if resp, err := client.DeleteAdapterVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detects text in the input document. Amazon Textract can detect lines of text
// and the words that make up a line of text. The input document must be in one of
// the following image formats: JPEG, PNG, PDF, or TIFF. DetectDocumentText
// returns the detected text in an array of Blockobjects.
//
// Each document page has as an associated Block of type PAGE. Each PAGE Block
// object is the parent of LINE Block objects that represent the lines of detected
// text on a page. A LINE Block object is a parent for each word that makes up the
// line. Words are represented by Block objects of type WORD.
//
// DetectDocumentText is a synchronous operation. To analyze documents
// asynchronously, use StartDocumentTextDetection.
//
// For more information, see [Document Text Detection].
//
// [Document Text Detection]: https://docs.aws.amazon.com/textract/latest/dg/how-it-works-detecting.html
func textract_DetectDocumentText(cfg aws.Config, client *textract.Client) {
	input := &textract.DetectDocumentTextInput{
		// Document: *types.Document, // Required
	}

	if len(_textractDocument) > 0 {
		if err := assignInputField(input, "Document", _textractDocument); err != nil {
			log.Errorf("invalid --document: %s", err.Error())
			return
		}
	}

	if resp, err := client.DetectDocumentText(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets configuration information for an adapter specified by an AdapterId,
// returning information on AdapterName, Description, CreationTime, AutoUpdate
// status, and FeatureTypes.
func textract_GetAdapter(cfg aws.Config, client *textract.Client) {
	input := &textract.GetAdapterInput{
		// AdapterId: *string, // Required
	}

	if len(_textractAdapterId) > 0 {
		input.AdapterId = aws.String(_textractAdapterId)
	}

	if resp, err := client.GetAdapter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets configuration information for the specified adapter version, including:
// AdapterId, AdapterVersion, FeatureTypes, Status, StatusMessage, DatasetConfig,
// KMSKeyId, OutputConfig, Tags and EvaluationMetrics.
func textract_GetAdapterVersion(cfg aws.Config, client *textract.Client) {
	input := &textract.GetAdapterVersionInput{
		// AdapterId: *string, // Required
		// AdapterVersion: *string, // Required
	}

	if len(_textractAdapterId) > 0 {
		input.AdapterId = aws.String(_textractAdapterId)
	}
	if len(_textractAdapterVersion) > 0 {
		input.AdapterVersion = aws.String(_textractAdapterVersion)
	}

	if resp, err := client.GetAdapterVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the results for an Amazon Textract asynchronous operation that analyzes
// text in a document.
//
// You start asynchronous text analysis by calling StartDocumentAnalysis, which returns a job
// identifier ( JobId ). When the text analysis operation finishes, Amazon Textract
// publishes a completion status to the Amazon Simple Notification Service (Amazon
// SNS) topic that's registered in the initial call to StartDocumentAnalysis . To
// get the results of the text-detection operation, first check that the status
// value published to the Amazon SNS topic is SUCCEEDED . If so, call
// GetDocumentAnalysis , and pass the job identifier ( JobId ) from the initial
// call to StartDocumentAnalysis .
//
// GetDocumentAnalysis returns an array of Block objects. The following types of
// information are returned:
//
// - Form data (key-value pairs). The related information is returned in two Block
// objects, each of type KEY_VALUE_SET : a KEY Block object and a VALUE Block
// object. For example, Name: Ana Silva Carolina contains a key and value. Name: is
// the key. Ana Silva Carolina is the value.
//
// - Table and table cell data. A TABLE Block object contains information about a
// detected table. A CELL Block object is returned for each cell in a table.
//
// - Lines and words of text. A LINE Block object contains one or more WORD Block
// objects. All lines and words that are detected in the document are returned
// (including text that doesn't have a relationship with the value of the
// StartDocumentAnalysis FeatureTypes input parameter).
//
// - Query. A QUERY Block object contains the query text, alias and link to the
// associated Query results block object.
//
// - Query Results. A QUERY_RESULT Block object contains the answer to the query
// and an ID that connects it to the query asked. This Block also contains a
// confidence score.
//
// While processing a document with queries, look out for
// INVALID_REQUEST_PARAMETERS output. This indicates that either the per page query
// limit has been exceeded or that the operation is trying to query a page in the
// document which doesn’t exist.
//
// Selection elements such as check boxes and option buttons (radio buttons) can
// be detected in form data and in tables. A SELECTION_ELEMENT Block object
// contains information about a selection element, including the selection status.
//
// Use the MaxResults parameter to limit the number of blocks that are returned.
// If there are more results than specified in MaxResults , the value of NextToken
// in the operation response contains a pagination token for getting the next set
// of results. To get the next page of results, call GetDocumentAnalysis , and
// populate the NextToken request parameter with the token value that's returned
// from the previous call to GetDocumentAnalysis .
//
// For more information, see [Document Text Analysis].
//
// [Document Text Analysis]: https://docs.aws.amazon.com/textract/latest/dg/how-it-works-analyzing.html
func textract_GetDocumentAnalysis(cfg aws.Config, client *textract.Client) {
	input := &textract.GetDocumentAnalysisInput{
		// JobId: *string, // Required
	}

	if len(_textractJobId) > 0 {
		input.JobId = aws.String(_textractJobId)
	}
	if len(_textractMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _textractMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_textractNextToken) > 0 {
		input.NextToken = aws.String(_textractNextToken)
	}

	if resp, err := client.GetDocumentAnalysis(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the results for an Amazon Textract asynchronous operation that detects
// text in a document. Amazon Textract can detect lines of text and the words that
// make up a line of text.
//
// You start asynchronous text detection by calling StartDocumentTextDetection, which returns a job
// identifier ( JobId ). When the text detection operation finishes, Amazon
// Textract publishes a completion status to the Amazon Simple Notification Service
// (Amazon SNS) topic that's registered in the initial call to
// StartDocumentTextDetection . To get the results of the text-detection operation,
// first check that the status value published to the Amazon SNS topic is SUCCEEDED
// . If so, call GetDocumentTextDetection , and pass the job identifier ( JobId )
// from the initial call to StartDocumentTextDetection .
//
// GetDocumentTextDetection returns an array of Block objects.
//
// Each document page has as an associated Block of type PAGE. Each PAGE Block
// object is the parent of LINE Block objects that represent the lines of detected
// text on a page. A LINE Block object is a parent for each word that makes up the
// line. Words are represented by Block objects of type WORD.
//
// Use the MaxResults parameter to limit the number of blocks that are returned.
// If there are more results than specified in MaxResults , the value of NextToken
// in the operation response contains a pagination token for getting the next set
// of results. To get the next page of results, call GetDocumentTextDetection , and
// populate the NextToken request parameter with the token value that's returned
// from the previous call to GetDocumentTextDetection .
//
// For more information, see [Document Text Detection].
//
// [Document Text Detection]: https://docs.aws.amazon.com/textract/latest/dg/how-it-works-detecting.html
func textract_GetDocumentTextDetection(cfg aws.Config, client *textract.Client) {
	input := &textract.GetDocumentTextDetectionInput{
		// JobId: *string, // Required
	}

	if len(_textractJobId) > 0 {
		input.JobId = aws.String(_textractJobId)
	}
	if len(_textractMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _textractMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_textractNextToken) > 0 {
		input.NextToken = aws.String(_textractNextToken)
	}

	if resp, err := client.GetDocumentTextDetection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the results for an Amazon Textract asynchronous operation that analyzes
// invoices and receipts. Amazon Textract finds contact information, items
// purchased, and vendor name, from input invoices and receipts.
//
// You start asynchronous invoice/receipt analysis by calling StartExpenseAnalysis, which returns a
// job identifier ( JobId ). Upon completion of the invoice/receipt analysis,
// Amazon Textract publishes the completion status to the Amazon Simple
// Notification Service (Amazon SNS) topic. This topic must be registered in the
// initial call to StartExpenseAnalysis . To get the results of the invoice/receipt
// analysis operation, first ensure that the status value published to the Amazon
// SNS topic is SUCCEEDED . If so, call GetExpenseAnalysis , and pass the job
// identifier ( JobId ) from the initial call to StartExpenseAnalysis .
//
// Use the MaxResults parameter to limit the number of blocks that are returned.
// If there are more results than specified in MaxResults , the value of NextToken
// in the operation response contains a pagination token for getting the next set
// of results. To get the next page of results, call GetExpenseAnalysis , and
// populate the NextToken request parameter with the token value that's returned
// from the previous call to GetExpenseAnalysis .
//
// For more information, see [Analyzing Invoices and Receipts].
//
// [Analyzing Invoices and Receipts]: https://docs.aws.amazon.com/textract/latest/dg/invoices-receipts.html
func textract_GetExpenseAnalysis(cfg aws.Config, client *textract.Client) {
	input := &textract.GetExpenseAnalysisInput{
		// JobId: *string, // Required
	}

	if len(_textractJobId) > 0 {
		input.JobId = aws.String(_textractJobId)
	}
	if len(_textractMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _textractMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_textractNextToken) > 0 {
		input.NextToken = aws.String(_textractNextToken)
	}

	if resp, err := client.GetExpenseAnalysis(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the results for an Amazon Textract asynchronous operation that analyzes
// text in a lending document.
//
// You start asynchronous text analysis by calling StartLendingAnalysis , which
// returns a job identifier ( JobId ). When the text analysis operation finishes,
// Amazon Textract publishes a completion status to the Amazon Simple Notification
// Service (Amazon SNS) topic that's registered in the initial call to
// StartLendingAnalysis .
//
// To get the results of the text analysis operation, first check that the status
// value published to the Amazon SNS topic is SUCCEEDED. If so, call
// GetLendingAnalysis, and pass the job identifier ( JobId ) from the initial call
// to StartLendingAnalysis .
func textract_GetLendingAnalysis(cfg aws.Config, client *textract.Client) {
	input := &textract.GetLendingAnalysisInput{
		// JobId: *string, // Required
	}

	if len(_textractJobId) > 0 {
		input.JobId = aws.String(_textractJobId)
	}
	if len(_textractMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _textractMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_textractNextToken) > 0 {
		input.NextToken = aws.String(_textractNextToken)
	}

	if resp, err := client.GetLendingAnalysis(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets summarized results for the StartLendingAnalysis operation, which analyzes
// text in a lending document. The returned summary consists of information about
// documents grouped together by a common document type. Information like detected
// signatures, page numbers, and split documents is returned with respect to the
// type of grouped document.
//
// You start asynchronous text analysis by calling StartLendingAnalysis , which
// returns a job identifier ( JobId ). When the text analysis operation finishes,
// Amazon Textract publishes a completion status to the Amazon Simple Notification
// Service (Amazon SNS) topic that's registered in the initial call to
// StartLendingAnalysis .
//
// To get the results of the text analysis operation, first check that the status
// value published to the Amazon SNS topic is SUCCEEDED. If so, call
// GetLendingAnalysisSummary , and pass the job identifier ( JobId ) from the
// initial call to StartLendingAnalysis .
func textract_GetLendingAnalysisSummary(cfg aws.Config, client *textract.Client) {
	input := &textract.GetLendingAnalysisSummaryInput{
		// JobId: *string, // Required
	}

	if len(_textractJobId) > 0 {
		input.JobId = aws.String(_textractJobId)
	}

	if resp, err := client.GetLendingAnalysisSummary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List all version of an adapter that meet the specified filtration criteria.
func textract_ListAdapterVersions(cfg aws.Config, client *textract.Client) {
	input := &textract.ListAdapterVersionsInput{}

	if len(_textractAdapterId) > 0 {
		input.AdapterId = aws.String(_textractAdapterId)
	}
	if len(_textractAfterCreationTime) > 0 {
		if err := assignInputField(input, "AfterCreationTime", _textractAfterCreationTime); err != nil {
			log.Errorf("invalid --after-creation-time: %s", err.Error())
			return
		}
	}
	if len(_textractBeforeCreationTime) > 0 {
		if err := assignInputField(input, "BeforeCreationTime", _textractBeforeCreationTime); err != nil {
			log.Errorf("invalid --before-creation-time: %s", err.Error())
			return
		}
	}
	if len(_textractMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _textractMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_textractNextToken) > 0 {
		input.NextToken = aws.String(_textractNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAdapterVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*textract.ListAdapterVersionsOutput
	p := textract.NewListAdapterVersionsPaginator(client, input)
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

// Lists all adapters that match the specified filtration criteria.
func textract_ListAdapters(cfg aws.Config, client *textract.Client) {
	input := &textract.ListAdaptersInput{}

	if len(_textractAfterCreationTime) > 0 {
		if err := assignInputField(input, "AfterCreationTime", _textractAfterCreationTime); err != nil {
			log.Errorf("invalid --after-creation-time: %s", err.Error())
			return
		}
	}
	if len(_textractBeforeCreationTime) > 0 {
		if err := assignInputField(input, "BeforeCreationTime", _textractBeforeCreationTime); err != nil {
			log.Errorf("invalid --before-creation-time: %s", err.Error())
			return
		}
	}
	if len(_textractMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _textractMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_textractNextToken) > 0 {
		input.NextToken = aws.String(_textractNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAdapters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*textract.ListAdaptersOutput
	p := textract.NewListAdaptersPaginator(client, input)
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

// Lists all tags for an Amazon Textract resource.
func textract_ListTagsForResource(cfg aws.Config, client *textract.Client) {
	input := &textract.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_textractResourceARN) > 0 {
		input.ResourceARN = aws.String(_textractResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the asynchronous analysis of an input document for relationships between
// detected items such as key-value pairs, tables, and selection elements.
//
// StartDocumentAnalysis can analyze text in documents that are in JPEG, PNG,
// TIFF, and PDF format. The documents are stored in an Amazon S3 bucket. Use DocumentLocationto
// specify the bucket name and file name of the document.
//
// StartDocumentAnalysis returns a job identifier ( JobId ) that you use to get the
// results of the operation. When text analysis is finished, Amazon Textract
// publishes a completion status to the Amazon Simple Notification Service (Amazon
// SNS) topic that you specify in NotificationChannel . To get the results of the
// text analysis operation, first check that the status value published to the
// Amazon SNS topic is SUCCEEDED . If so, call GetDocumentAnalysis, and pass the job identifier ( JobId
// ) from the initial call to StartDocumentAnalysis .
//
// For more information, see [Document Text Analysis].
//
// [Document Text Analysis]: https://docs.aws.amazon.com/textract/latest/dg/how-it-works-analyzing.html
func textract_StartDocumentAnalysis(cfg aws.Config, client *textract.Client) {
	input := &textract.StartDocumentAnalysisInput{
		// DocumentLocation: *types.DocumentLocation, // Required
		// FeatureTypes: []types.FeatureType, // Required
	}

	if len(_textractDocumentLocation) > 0 {
		if err := assignInputField(input, "DocumentLocation", _textractDocumentLocation); err != nil {
			log.Errorf("invalid --document-location: %s", err.Error())
			return
		}
	}
	if len(_textractFeatureTypes) > 0 {
		if err := assignInputField(input, "FeatureTypes", _textractFeatureTypes); err != nil {
			log.Errorf("invalid --feature-types: %s", err.Error())
			return
		}
	}
	if len(_textractAdaptersConfig) > 0 {
		if err := assignInputField(input, "AdaptersConfig", _textractAdaptersConfig); err != nil {
			log.Errorf("invalid --adapters-config: %s", err.Error())
			return
		}
	}
	if len(_textractClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_textractClientRequestToken)
	}
	if len(_textractJobTag) > 0 {
		input.JobTag = aws.String(_textractJobTag)
	}
	if len(_textractKMSKeyId) > 0 {
		input.KMSKeyId = aws.String(_textractKMSKeyId)
	}
	if len(_textractNotificationChannel) > 0 {
		if err := assignInputField(input, "NotificationChannel", _textractNotificationChannel); err != nil {
			log.Errorf("invalid --notification-channel: %s", err.Error())
			return
		}
	}
	if len(_textractOutputConfig) > 0 {
		if err := assignInputField(input, "OutputConfig", _textractOutputConfig); err != nil {
			log.Errorf("invalid --output-config: %s", err.Error())
			return
		}
	}
	if len(_textractQueriesConfig) > 0 {
		if err := assignInputField(input, "QueriesConfig", _textractQueriesConfig); err != nil {
			log.Errorf("invalid --queries-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartDocumentAnalysis(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the asynchronous detection of text in a document. Amazon Textract can
// detect lines of text and the words that make up a line of text.
//
// StartDocumentTextDetection can analyze text in documents that are in JPEG, PNG,
// TIFF, and PDF format. The documents are stored in an Amazon S3 bucket. Use DocumentLocationto
// specify the bucket name and file name of the document.
//
// StartDocumentTextDetection returns a job identifier ( JobId ) that you use to
// get the results of the operation. When text detection is finished, Amazon
// Textract publishes a completion status to the Amazon Simple Notification Service
// (Amazon SNS) topic that you specify in NotificationChannel . To get the results
// of the text detection operation, first check that the status value published to
// the Amazon SNS topic is SUCCEEDED . If so, call GetDocumentTextDetection, and pass the job identifier (
// JobId ) from the initial call to StartDocumentTextDetection .
//
// For more information, see [Document Text Detection].
//
// [Document Text Detection]: https://docs.aws.amazon.com/textract/latest/dg/how-it-works-detecting.html
func textract_StartDocumentTextDetection(cfg aws.Config, client *textract.Client) {
	input := &textract.StartDocumentTextDetectionInput{
		// DocumentLocation: *types.DocumentLocation, // Required
	}

	if len(_textractDocumentLocation) > 0 {
		if err := assignInputField(input, "DocumentLocation", _textractDocumentLocation); err != nil {
			log.Errorf("invalid --document-location: %s", err.Error())
			return
		}
	}
	if len(_textractClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_textractClientRequestToken)
	}
	if len(_textractJobTag) > 0 {
		input.JobTag = aws.String(_textractJobTag)
	}
	if len(_textractKMSKeyId) > 0 {
		input.KMSKeyId = aws.String(_textractKMSKeyId)
	}
	if len(_textractNotificationChannel) > 0 {
		if err := assignInputField(input, "NotificationChannel", _textractNotificationChannel); err != nil {
			log.Errorf("invalid --notification-channel: %s", err.Error())
			return
		}
	}
	if len(_textractOutputConfig) > 0 {
		if err := assignInputField(input, "OutputConfig", _textractOutputConfig); err != nil {
			log.Errorf("invalid --output-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartDocumentTextDetection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the asynchronous analysis of invoices or receipts for data like contact
// information, items purchased, and vendor names.
//
// StartExpenseAnalysis can analyze text in documents that are in JPEG, PNG, and
// PDF format. The documents must be stored in an Amazon S3 bucket. Use the DocumentLocation
// parameter to specify the name of your S3 bucket and the name of the document in
// that bucket.
//
// StartExpenseAnalysis returns a job identifier ( JobId ) that you will provide to
// GetExpenseAnalysis to retrieve the results of the operation. When the analysis
// of the input invoices/receipts is finished, Amazon Textract publishes a
// completion status to the Amazon Simple Notification Service (Amazon SNS) topic
// that you provide to the NotificationChannel . To obtain the results of the
// invoice and receipt analysis operation, ensure that the status value published
// to the Amazon SNS topic is SUCCEEDED . If so, call GetExpenseAnalysis, and pass the job
// identifier ( JobId ) that was returned by your call to StartExpenseAnalysis .
//
// For more information, see [Analyzing Invoices and Receipts].
//
// [Analyzing Invoices and Receipts]: https://docs.aws.amazon.com/textract/latest/dg/invoice-receipts.html
func textract_StartExpenseAnalysis(cfg aws.Config, client *textract.Client) {
	input := &textract.StartExpenseAnalysisInput{
		// DocumentLocation: *types.DocumentLocation, // Required
	}

	if len(_textractDocumentLocation) > 0 {
		if err := assignInputField(input, "DocumentLocation", _textractDocumentLocation); err != nil {
			log.Errorf("invalid --document-location: %s", err.Error())
			return
		}
	}
	if len(_textractClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_textractClientRequestToken)
	}
	if len(_textractJobTag) > 0 {
		input.JobTag = aws.String(_textractJobTag)
	}
	if len(_textractKMSKeyId) > 0 {
		input.KMSKeyId = aws.String(_textractKMSKeyId)
	}
	if len(_textractNotificationChannel) > 0 {
		if err := assignInputField(input, "NotificationChannel", _textractNotificationChannel); err != nil {
			log.Errorf("invalid --notification-channel: %s", err.Error())
			return
		}
	}
	if len(_textractOutputConfig) > 0 {
		if err := assignInputField(input, "OutputConfig", _textractOutputConfig); err != nil {
			log.Errorf("invalid --output-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartExpenseAnalysis(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the classification and analysis of an input document.
// StartLendingAnalysis initiates the classification and analysis of a packet of
// lending documents. StartLendingAnalysis operates on a document file located in
// an Amazon S3 bucket.
//
// StartLendingAnalysis can analyze text in documents that are in one of the
// following formats: JPEG, PNG, TIFF, PDF. Use DocumentLocation to specify the
// bucket name and the file name of the document.
//
// StartLendingAnalysis returns a job identifier ( JobId ) that you use to get the
// results of the operation. When the text analysis is finished, Amazon Textract
// publishes a completion status to the Amazon Simple Notification Service (Amazon
// SNS) topic that you specify in NotificationChannel . To get the results of the
// text analysis operation, first check that the status value published to the
// Amazon SNS topic is SUCCEEDED. If the status is SUCCEEDED you can call either
// GetLendingAnalysis or GetLendingAnalysisSummary and provide the JobId to obtain
// the results of the analysis.
//
// If using OutputConfig to specify an Amazon S3 bucket, the output will be
// contained within the specified prefix in a directory labeled with the job-id. In
// the directory there are 3 sub-directories:
//
// - detailedResponse (contains the GetLendingAnalysis response)
//
// - summaryResponse (for the GetLendingAnalysisSummary response)
//
// - splitDocuments (documents split across logical boundaries)
func textract_StartLendingAnalysis(cfg aws.Config, client *textract.Client) {
	input := &textract.StartLendingAnalysisInput{
		// DocumentLocation: *types.DocumentLocation, // Required
	}

	if len(_textractDocumentLocation) > 0 {
		if err := assignInputField(input, "DocumentLocation", _textractDocumentLocation); err != nil {
			log.Errorf("invalid --document-location: %s", err.Error())
			return
		}
	}
	if len(_textractClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_textractClientRequestToken)
	}
	if len(_textractJobTag) > 0 {
		input.JobTag = aws.String(_textractJobTag)
	}
	if len(_textractKMSKeyId) > 0 {
		input.KMSKeyId = aws.String(_textractKMSKeyId)
	}
	if len(_textractNotificationChannel) > 0 {
		if err := assignInputField(input, "NotificationChannel", _textractNotificationChannel); err != nil {
			log.Errorf("invalid --notification-channel: %s", err.Error())
			return
		}
	}
	if len(_textractOutputConfig) > 0 {
		if err := assignInputField(input, "OutputConfig", _textractOutputConfig); err != nil {
			log.Errorf("invalid --output-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartLendingAnalysis(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more tags to the specified resource.
func textract_TagResource(cfg aws.Config, client *textract.Client) {
	input := &textract.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_textractResourceARN) > 0 {
		input.ResourceARN = aws.String(_textractResourceARN)
	}
	if len(_textractTags) > 0 {
		if err := assignInputField(input, "Tags", _textractTags); err != nil {
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

// Removes any tags with the specified keys from the specified resource.
func textract_UntagResource(cfg aws.Config, client *textract.Client) {
	input := &textract.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_textractResourceARN) > 0 {
		input.ResourceARN = aws.String(_textractResourceARN)
	}
	if len(_textractTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _textractTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the configuration for an adapter. FeatureTypes configurations cannot be
// updated. At least one new parameter must be specified as an argument.
func textract_UpdateAdapter(cfg aws.Config, client *textract.Client) {
	input := &textract.UpdateAdapterInput{
		// AdapterId: *string, // Required
	}

	if len(_textractAdapterId) > 0 {
		input.AdapterId = aws.String(_textractAdapterId)
	}
	if len(_textractAdapterName) > 0 {
		input.AdapterName = aws.String(_textractAdapterName)
	}
	if len(_textractAutoUpdate) > 0 {
		if err := assignInputField(input, "AutoUpdate", _textractAutoUpdate); err != nil {
			log.Errorf("invalid --auto-update: %s", err.Error())
			return
		}
	}
	if len(_textractDescription) > 0 {
		input.Description = aws.String(_textractDescription)
	}

	if resp, err := client.UpdateAdapter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_textractCmd)
	_textractCmd.Flags().SortFlags = false

	_textractCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_textractCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_textractCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_textractCmd.Flags().StringVarP(&_textractAdapterId, "adapter-id", "", "", "Adapter ID")
	_textractCmd.Flags().StringVarP(&_textractAdapterName, "adapter-name", "", "", "Adapter Name")
	_textractCmd.Flags().StringVarP(&_textractAdapterVersion, "adapter-version", "", "", "Adapter Version")
	_textractCmd.Flags().StringVarP(&_textractAdaptersConfig, "adapters-config", "", "", "Adapters Config")
	_textractCmd.Flags().StringVarP(&_textractAfterCreationTime, "after-creation-time", "", "", "After Creation Time")
	_textractCmd.Flags().StringVarP(&_textractAutoUpdate, "auto-update", "", "", "Auto Update")
	_textractCmd.Flags().StringVarP(&_textractBeforeCreationTime, "before-creation-time", "", "", "Before Creation Time")
	_textractCmd.Flags().StringVarP(&_textractClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_textractCmd.Flags().StringVarP(&_textractDatasetConfig, "dataset-config", "", "", "Dataset Config")
	_textractCmd.Flags().StringVarP(&_textractDescription, "description", "", "", "Description")
	_textractCmd.Flags().StringVarP(&_textractDocument, "document", "", "", "Document")
	_textractCmd.Flags().StringVarP(&_textractDocumentLocation, "document-location", "", "", "Document Location")
	_textractCmd.Flags().StringVarP(&_textractDocumentPages, "document-pages", "", "", "Document Pages")
	_textractCmd.Flags().StringVarP(&_textractFeatureTypes, "feature-types", "", "", "Feature Types")
	_textractCmd.Flags().StringVarP(&_textractHumanLoopConfig, "human-loop-config", "", "", "Human Loop Config")
	_textractCmd.Flags().StringVarP(&_textractJobId, "job-id", "", "", "Job ID")
	_textractCmd.Flags().StringVarP(&_textractJobTag, "job-tag", "", "", "Job Tag")
	_textractCmd.Flags().StringVarP(&_textractKMSKeyId, "kms-key-id", "", "", "KMS Key ID")
	_textractCmd.Flags().StringVarP(&_textractMaxResults, "max-results", "", "", "Max Results")
	_textractCmd.Flags().StringVarP(&_textractNextToken, "next-token", "", "", "Next Token")
	_textractCmd.Flags().StringVarP(&_textractNotificationChannel, "notification-channel", "", "", "Notification Channel")
	_textractCmd.Flags().StringVarP(&_textractOutputConfig, "output-config", "", "", "Output Config")
	_textractCmd.Flags().StringVarP(&_textractQueriesConfig, "queries-config", "", "", "Queries Config")
	_textractCmd.Flags().StringVarP(&_textractResourceARN, "resource-arn", "", "", "Resource ARN")
	_textractCmd.Flags().StringSliceVarP(&_textractTagKeys, "tag-keys", "", nil, "Tag Keys")
	_textractCmd.Flags().StringVarP(&_textractTags, "tags", "", "", "Tags")

	_textractCmd.Flags().BoolVarP(&_textractAnalyzeDocument, "analyze-document", "", false, "Analyze Document")
	_textractCmd.Flags().BoolVarP(&_textractAnalyzeExpense, "analyze-expense", "", false, "Analyze Expense")
	_textractCmd.Flags().BoolVarP(&_textractAnalyzeID, "analyze-id", "", false, "Analyze ID")
	_textractCmd.Flags().BoolVarP(&_textractCreateAdapter, "create-adapter", "", false, "Create Adapter")
	_textractCmd.Flags().BoolVarP(&_textractCreateAdapterVersion, "create-adapter-version", "", false, "Create Adapter Version")
	_textractCmd.Flags().BoolVarP(&_textractDeleteAdapter, "delete-adapter", "", false, "Delete Adapter")
	_textractCmd.Flags().BoolVarP(&_textractDeleteAdapterVersion, "delete-adapter-version", "", false, "Delete Adapter Version")
	_textractCmd.Flags().BoolVarP(&_textractDetectDocumentText, "detect-document-text", "", false, "Detect Document Text")
	_textractCmd.Flags().BoolVarP(&_textractGetAdapter, "get-adapter", "", false, "Get Adapter")
	_textractCmd.Flags().BoolVarP(&_textractGetAdapterVersion, "get-adapter-version", "", false, "Get Adapter Version")
	_textractCmd.Flags().BoolVarP(&_textractGetDocumentAnalysis, "get-document-analysis", "", false, "Get Document Analysis")
	_textractCmd.Flags().BoolVarP(&_textractGetDocumentTextDetection, "get-document-text-detection", "", false, "Get Document Text Detection")
	_textractCmd.Flags().BoolVarP(&_textractGetExpenseAnalysis, "get-expense-analysis", "", false, "Get Expense Analysis")
	_textractCmd.Flags().BoolVarP(&_textractGetLendingAnalysis, "get-lending-analysis", "", false, "Get Lending Analysis")
	_textractCmd.Flags().BoolVarP(&_textractGetLendingAnalysisSummary, "get-lending-analysis-summary", "", false, "Get Lending Analysis Summary")
	_textractCmd.Flags().BoolVarP(&_textractListAdapterVersions, "list-adapter-versions", "", false, "List Adapter Versions")
	_textractCmd.Flags().BoolVarP(&_textractListAdapters, "list-adapters", "", false, "List Adapters")
	_textractCmd.Flags().BoolVarP(&_textractListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_textractCmd.Flags().BoolVarP(&_textractStartDocumentAnalysis, "start-document-analysis", "", false, "Start Document Analysis")
	_textractCmd.Flags().BoolVarP(&_textractStartDocumentTextDetection, "start-document-text-detection", "", false, "Start Document Text Detection")
	_textractCmd.Flags().BoolVarP(&_textractStartExpenseAnalysis, "start-expense-analysis", "", false, "Start Expense Analysis")
	_textractCmd.Flags().BoolVarP(&_textractStartLendingAnalysis, "start-lending-analysis", "", false, "Start Lending Analysis")
	_textractCmd.Flags().BoolVarP(&_textractTagResource, "tag-resource", "", false, "Tag Resource")
	_textractCmd.Flags().BoolVarP(&_textractUntagResource, "untag-resource", "", false, "Untag Resource")
	_textractCmd.Flags().BoolVarP(&_textractUpdateAdapter, "update-adapter", "", false, "Update Adapter")

}
