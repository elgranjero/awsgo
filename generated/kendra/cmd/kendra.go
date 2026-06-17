package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kendra"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// kendraCmd represents the kendra command
var _kendraCmd = &cobra.Command{
	Use:   "kendra",
	Short: "AWS kendra CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := kendra.NewFromConfig(cfg)
		if _kendraAssociateEntitiesToExperience {
			kendra_AssociateEntitiesToExperience(cfg, client)
			return
		}
		if _kendraAssociatePersonasToEntities {
			kendra_AssociatePersonasToEntities(cfg, client)
			return
		}
		if _kendraBatchDeleteDocument {
			kendra_BatchDeleteDocument(cfg, client)
			return
		}
		if _kendraBatchDeleteFeaturedResultsSet {
			kendra_BatchDeleteFeaturedResultsSet(cfg, client)
			return
		}
		if _kendraBatchGetDocumentStatus {
			kendra_BatchGetDocumentStatus(cfg, client)
			return
		}
		if _kendraBatchPutDocument {
			kendra_BatchPutDocument(cfg, client)
			return
		}
		if _kendraClearQuerySuggestions {
			kendra_ClearQuerySuggestions(cfg, client)
			return
		}
		if _kendraCreateAccessControlConfiguration {
			kendra_CreateAccessControlConfiguration(cfg, client)
			return
		}
		if _kendraCreateDataSource {
			kendra_CreateDataSource(cfg, client)
			return
		}
		if _kendraCreateExperience {
			kendra_CreateExperience(cfg, client)
			return
		}
		if _kendraCreateFaq {
			kendra_CreateFaq(cfg, client)
			return
		}
		if _kendraCreateFeaturedResultsSet {
			kendra_CreateFeaturedResultsSet(cfg, client)
			return
		}
		if _kendraCreateIndex {
			kendra_CreateIndex(cfg, client)
			return
		}
		if _kendraCreateQuerySuggestionsBlockList {
			kendra_CreateQuerySuggestionsBlockList(cfg, client)
			return
		}
		if _kendraCreateThesaurus {
			kendra_CreateThesaurus(cfg, client)
			return
		}
		if _kendraDeleteAccessControlConfiguration {
			kendra_DeleteAccessControlConfiguration(cfg, client)
			return
		}
		if _kendraDeleteDataSource {
			kendra_DeleteDataSource(cfg, client)
			return
		}
		if _kendraDeleteExperience {
			kendra_DeleteExperience(cfg, client)
			return
		}
		if _kendraDeleteFaq {
			kendra_DeleteFaq(cfg, client)
			return
		}
		if _kendraDeleteIndex {
			kendra_DeleteIndex(cfg, client)
			return
		}
		if _kendraDeletePrincipalMapping {
			kendra_DeletePrincipalMapping(cfg, client)
			return
		}
		if _kendraDeleteQuerySuggestionsBlockList {
			kendra_DeleteQuerySuggestionsBlockList(cfg, client)
			return
		}
		if _kendraDeleteThesaurus {
			kendra_DeleteThesaurus(cfg, client)
			return
		}
		if _kendraDescribeAccessControlConfiguration {
			kendra_DescribeAccessControlConfiguration(cfg, client)
			return
		}
		if _kendraDescribeDataSource {
			kendra_DescribeDataSource(cfg, client)
			return
		}
		if _kendraDescribeExperience {
			kendra_DescribeExperience(cfg, client)
			return
		}
		if _kendraDescribeFaq {
			kendra_DescribeFaq(cfg, client)
			return
		}
		if _kendraDescribeFeaturedResultsSet {
			kendra_DescribeFeaturedResultsSet(cfg, client)
			return
		}
		if _kendraDescribeIndex {
			kendra_DescribeIndex(cfg, client)
			return
		}
		if _kendraDescribePrincipalMapping {
			kendra_DescribePrincipalMapping(cfg, client)
			return
		}
		if _kendraDescribeQuerySuggestionsBlockList {
			kendra_DescribeQuerySuggestionsBlockList(cfg, client)
			return
		}
		if _kendraDescribeQuerySuggestionsConfig {
			kendra_DescribeQuerySuggestionsConfig(cfg, client)
			return
		}
		if _kendraDescribeThesaurus {
			kendra_DescribeThesaurus(cfg, client)
			return
		}
		if _kendraDisassociateEntitiesFromExperience {
			kendra_DisassociateEntitiesFromExperience(cfg, client)
			return
		}
		if _kendraDisassociatePersonasFromEntities {
			kendra_DisassociatePersonasFromEntities(cfg, client)
			return
		}
		if _kendraGetQuerySuggestions {
			kendra_GetQuerySuggestions(cfg, client)
			return
		}
		if _kendraGetSnapshots {
			kendra_GetSnapshots(cfg, client)
			return
		}
		if _kendraListAccessControlConfigurations {
			kendra_ListAccessControlConfigurations(cfg, client)
			return
		}
		if _kendraListDataSourceSyncJobs {
			kendra_ListDataSourceSyncJobs(cfg, client)
			return
		}
		if _kendraListDataSources {
			kendra_ListDataSources(cfg, client)
			return
		}
		if _kendraListEntityPersonas {
			kendra_ListEntityPersonas(cfg, client)
			return
		}
		if _kendraListExperienceEntities {
			kendra_ListExperienceEntities(cfg, client)
			return
		}
		if _kendraListExperiences {
			kendra_ListExperiences(cfg, client)
			return
		}
		if _kendraListFaqs {
			kendra_ListFaqs(cfg, client)
			return
		}
		if _kendraListFeaturedResultsSets {
			kendra_ListFeaturedResultsSets(cfg, client)
			return
		}
		if _kendraListGroupsOlderThanOrderingId {
			kendra_ListGroupsOlderThanOrderingId(cfg, client)
			return
		}
		if _kendraListIndices {
			kendra_ListIndices(cfg, client)
			return
		}
		if _kendraListQuerySuggestionsBlockLists {
			kendra_ListQuerySuggestionsBlockLists(cfg, client)
			return
		}
		if _kendraListTagsForResource {
			kendra_ListTagsForResource(cfg, client)
			return
		}
		if _kendraListThesauri {
			kendra_ListThesauri(cfg, client)
			return
		}
		if _kendraPutPrincipalMapping {
			kendra_PutPrincipalMapping(cfg, client)
			return
		}
		if _kendraQuery {
			kendra_Query(cfg, client)
			return
		}
		if _kendraRetrieve {
			kendra_Retrieve(cfg, client)
			return
		}
		if _kendraStartDataSourceSyncJob {
			kendra_StartDataSourceSyncJob(cfg, client)
			return
		}
		if _kendraStopDataSourceSyncJob {
			kendra_StopDataSourceSyncJob(cfg, client)
			return
		}
		if _kendraSubmitFeedback {
			kendra_SubmitFeedback(cfg, client)
			return
		}
		if _kendraTagResource {
			kendra_TagResource(cfg, client)
			return
		}
		if _kendraUntagResource {
			kendra_UntagResource(cfg, client)
			return
		}
		if _kendraUpdateAccessControlConfiguration {
			kendra_UpdateAccessControlConfiguration(cfg, client)
			return
		}
		if _kendraUpdateDataSource {
			kendra_UpdateDataSource(cfg, client)
			return
		}
		if _kendraUpdateExperience {
			kendra_UpdateExperience(cfg, client)
			return
		}
		if _kendraUpdateFeaturedResultsSet {
			kendra_UpdateFeaturedResultsSet(cfg, client)
			return
		}
		if _kendraUpdateIndex {
			kendra_UpdateIndex(cfg, client)
			return
		}
		if _kendraUpdateQuerySuggestionsBlockList {
			kendra_UpdateQuerySuggestionsBlockList(cfg, client)
			return
		}
		if _kendraUpdateQuerySuggestionsConfig {
			kendra_UpdateQuerySuggestionsConfig(cfg, client)
			return
		}
		if _kendraUpdateThesaurus {
			kendra_UpdateThesaurus(cfg, client)
			return
		}

	},
}

var (
	_kendraAssociateEntitiesToExperience      bool
	_kendraAssociatePersonasToEntities        bool
	_kendraBatchDeleteDocument                bool
	_kendraBatchDeleteFeaturedResultsSet      bool
	_kendraBatchGetDocumentStatus             bool
	_kendraBatchPutDocument                   bool
	_kendraClearQuerySuggestions              bool
	_kendraCreateAccessControlConfiguration   bool
	_kendraCreateDataSource                   bool
	_kendraCreateExperience                   bool
	_kendraCreateFaq                          bool
	_kendraCreateFeaturedResultsSet           bool
	_kendraCreateIndex                        bool
	_kendraCreateQuerySuggestionsBlockList    bool
	_kendraCreateThesaurus                    bool
	_kendraDeleteAccessControlConfiguration   bool
	_kendraDeleteDataSource                   bool
	_kendraDeleteExperience                   bool
	_kendraDeleteFaq                          bool
	_kendraDeleteIndex                        bool
	_kendraDeletePrincipalMapping             bool
	_kendraDeleteQuerySuggestionsBlockList    bool
	_kendraDeleteThesaurus                    bool
	_kendraDescribeAccessControlConfiguration bool
	_kendraDescribeDataSource                 bool
	_kendraDescribeExperience                 bool
	_kendraDescribeFaq                        bool
	_kendraDescribeFeaturedResultsSet         bool
	_kendraDescribeIndex                      bool
	_kendraDescribePrincipalMapping           bool
	_kendraDescribeQuerySuggestionsBlockList  bool
	_kendraDescribeQuerySuggestionsConfig     bool
	_kendraDescribeThesaurus                  bool
	_kendraDisassociateEntitiesFromExperience bool
	_kendraDisassociatePersonasFromEntities   bool
	_kendraGetQuerySuggestions                bool
	_kendraGetSnapshots                       bool
	_kendraListAccessControlConfigurations    bool
	_kendraListDataSourceSyncJobs             bool
	_kendraListDataSources                    bool
	_kendraListEntityPersonas                 bool
	_kendraListExperienceEntities             bool
	_kendraListExperiences                    bool
	_kendraListFaqs                           bool
	_kendraListFeaturedResultsSets            bool
	_kendraListGroupsOlderThanOrderingId      bool
	_kendraListIndices                        bool
	_kendraListQuerySuggestionsBlockLists     bool
	_kendraListTagsForResource                bool
	_kendraListThesauri                       bool
	_kendraPutPrincipalMapping                bool
	_kendraQuery                              bool
	_kendraRetrieve                           bool
	_kendraStartDataSourceSyncJob             bool
	_kendraStopDataSourceSyncJob              bool
	_kendraSubmitFeedback                     bool
	_kendraTagResource                        bool
	_kendraUntagResource                      bool
	_kendraUpdateAccessControlConfiguration   bool
	_kendraUpdateDataSource                   bool
	_kendraUpdateExperience                   bool
	_kendraUpdateFeaturedResultsSet           bool
	_kendraUpdateIndex                        bool
	_kendraUpdateQuerySuggestionsBlockList    bool
	_kendraUpdateQuerySuggestionsConfig       bool
	_kendraUpdateThesaurus                    bool

	_kendraAccessControlList                       string
	_kendraAttributeFilter                         string
	_kendraAttributeSuggestionsConfig              string
	_kendraCapacityUnits                           string
	_kendraClickFeedbackItems                      string
	_kendraClientToken                             string
	_kendraCollapseConfiguration                   string
	_kendraConfiguration                           string
	_kendraCustomDocumentEnrichmentConfiguration   string
	_kendraDataSourceId                            string
	_kendraDataSourceSyncJobMetricTarget           string
	_kendraDescription                             string
	_kendraDocumentIdList                          []string
	_kendraDocumentInfoList                        string
	_kendraDocumentMetadataConfigurationUpdates    string
	_kendraDocumentRelevanceOverrideConfigurations string
	_kendraDocuments                               string
	_kendraEdition                                 string
	_kendraEntityIds                               []string
	_kendraEntityList                              string
	_kendraFacets                                  string
	_kendraFeaturedDocuments                       string
	_kendraFeaturedResultsSetId                    string
	_kendraFeaturedResultsSetIds                   []string
	_kendraFeaturedResultsSetName                  string
	_kendraFileFormat                              string
	_kendraGroupId                                 string
	_kendraGroupMembers                            string
	_kendraHierarchicalAccessControlList           string
	_kendraId                                      string
	_kendraIncludeQueriesWithoutUserInformation    string
	_kendraIndexId                                 string
	_kendraInterval                                string
	_kendraLanguageCode                            string
	_kendraMaxResults                              string
	_kendraMaxSuggestionsCount                     string
	_kendraMetricType                              string
	_kendraMinimumNumberOfQueryingUsers            string
	_kendraMinimumQueryCount                       string
	_kendraMode                                    string
	_kendraName                                    string
	_kendraNextToken                               string
	_kendraOrderingId                              string
	_kendraPageNumber                              string
	_kendraPageSize                                string
	_kendraPersonas                                string
	_kendraQueryId                                 string
	_kendraQueryLogLookBackWindowInDays            string
	_kendraQueryResultTypeFilter                   string
	_kendraQueryText                               string
	_kendraQueryTexts                              []string
	_kendraRelevanceFeedbackItems                  string
	_kendraRequestedDocumentAttributes             []string
	_kendraResourceARN                             string
	_kendraRoleArn                                 string
	_kendraS3Path                                  string
	_kendraSchedule                                string
	_kendraServerSideEncryptionConfiguration       string
	_kendraSortingConfiguration                    string
	_kendraSortingConfigurations                   string
	_kendraSourceS3Path                            string
	_kendraSpellCorrectionConfiguration            string
	_kendraStartTimeFilter                         string
	_kendraStatus                                  string
	_kendraStatusFilter                            string
	_kendraSuggestionTypes                         string
	_kendraTagKeys                                 []string
	_kendraTags                                    string
	_kendraType                                    string
	_kendraUserContext                             string
	_kendraUserContextPolicy                       string
	_kendraUserGroupResolutionConfiguration        string
	_kendraUserTokenConfigurations                 string
	_kendraVisitorId                               string
	_kendraVpcConfiguration                        string
)

// Grants users or groups in your IAM Identity Center identity source access to
// your Amazon Kendra experience. You can create an Amazon Kendra experience such
// as a search application. For more information on creating a search application
// experience, see [Building a search experience with no code].
//
// [Building a search experience with no code]: https://docs.aws.amazon.com/kendra/latest/dg/deploying-search-experience-no-code.html
func kendra_AssociateEntitiesToExperience(cfg aws.Config, client *kendra.Client) {
	input := &kendra.AssociateEntitiesToExperienceInput{
		// EntityList: []types.EntityConfiguration, // Required
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraEntityList) > 0 {
		if err := assignInputField(input, "EntityList", _kendraEntityList); err != nil {
			log.Errorf("invalid --entity-list: %s", err.Error())
			return
		}
	}
	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}

	if resp, err := client.AssociateEntitiesToExperience(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Defines the specific permissions of users or groups in your IAM Identity Center
// identity source with access to your Amazon Kendra experience. You can create an
// Amazon Kendra experience such as a search application. For more information on
// creating a search application experience, see [Building a search experience with no code].
//
// [Building a search experience with no code]: https://docs.aws.amazon.com/kendra/latest/dg/deploying-search-experience-no-code.html
func kendra_AssociatePersonasToEntities(cfg aws.Config, client *kendra.Client) {
	input := &kendra.AssociatePersonasToEntitiesInput{
		// Id: *string, // Required
		// IndexId: *string, // Required
		// Personas: []types.EntityPersonaConfiguration, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraPersonas) > 0 {
		if err := assignInputField(input, "Personas", _kendraPersonas); err != nil {
			log.Errorf("invalid --personas: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociatePersonasToEntities(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes one or more documents from an index. The documents must have been added
// with the BatchPutDocument API.
//
// The documents are deleted asynchronously. You can see the progress of the
// deletion by using Amazon Web Services CloudWatch. Any error messages related to
// the processing of the batch are sent to your Amazon Web Services CloudWatch log.
// You can also use the BatchGetDocumentStatus API to monitor the progress of
// deleting your documents.
//
// Deleting documents from an index using BatchDeleteDocument could take up to an
// hour or more, depending on the number of documents you want to delete.
func kendra_BatchDeleteDocument(cfg aws.Config, client *kendra.Client) {
	input := &kendra.BatchDeleteDocumentInput{
		// DocumentIdList: []string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraDocumentIdList) > 0 {
		input.DocumentIdList = append([]string(nil), _kendraDocumentIdList...)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraDataSourceSyncJobMetricTarget) > 0 {
		if err := assignInputField(input, "DataSourceSyncJobMetricTarget", _kendraDataSourceSyncJobMetricTarget); err != nil {
			log.Errorf("invalid --data-source-sync-job-metric-target: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchDeleteDocument(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes one or more sets of featured results. Features results are placed above
// all other results for certain queries. If there's an exact match of a query,
// then one or more specific documents are featured in the search results.
func kendra_BatchDeleteFeaturedResultsSet(cfg aws.Config, client *kendra.Client) {
	input := &kendra.BatchDeleteFeaturedResultsSetInput{
		// FeaturedResultsSetIds: []string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraFeaturedResultsSetIds) > 0 {
		input.FeaturedResultsSetIds = append([]string(nil), _kendraFeaturedResultsSetIds...)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}

	if resp, err := client.BatchDeleteFeaturedResultsSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the indexing status for one or more documents submitted with the [BatchPutDocument] API.
// When you use the BatchPutDocument API, documents are indexed asynchronously.
// You can use the BatchGetDocumentStatus API to get the current status of a list
// of documents so that you can determine if they have been successfully indexed.
//
// You can also use the BatchGetDocumentStatus API to check the status of the [BatchDeleteDocument]
// API. When a document is deleted from the index, Amazon Kendra returns NOT_FOUND
// as the status.
//
// [BatchDeleteDocument]: https://docs.aws.amazon.com/kendra/latest/dg/API_BatchDeleteDocument.html
// [BatchPutDocument]: https://docs.aws.amazon.com/kendra/latest/dg/API_BatchPutDocument.html
func kendra_BatchGetDocumentStatus(cfg aws.Config, client *kendra.Client) {
	input := &kendra.BatchGetDocumentStatusInput{
		// DocumentInfoList: []types.DocumentInfo, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraDocumentInfoList) > 0 {
		if err := assignInputField(input, "DocumentInfoList", _kendraDocumentInfoList); err != nil {
			log.Errorf("invalid --document-info-list: %s", err.Error())
			return
		}
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}

	if resp, err := client.BatchGetDocumentStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more documents to an index.
// The BatchPutDocument API enables you to ingest inline documents or a set of
// documents stored in an Amazon S3 bucket. Use this API to ingest your text and
// unstructured text into an index, add custom attributes to the documents, and to
// attach an access control list to the documents added to the index.
//
// The documents are indexed asynchronously. You can see the progress of the batch
// using Amazon Web Services CloudWatch. Any error messages related to processing
// the batch are sent to your Amazon Web Services CloudWatch log. You can also use
// the BatchGetDocumentStatus API to monitor the progress of indexing your
// documents.
//
// For an example of ingesting inline documents using Python and Java SDKs, see [Adding files directly to an index].
//
// [Adding files directly to an index]: https://docs.aws.amazon.com/kendra/latest/dg/in-adding-binary-doc.html
func kendra_BatchPutDocument(cfg aws.Config, client *kendra.Client) {
	input := &kendra.BatchPutDocumentInput{
		// Documents: []types.Document, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraDocuments) > 0 {
		if err := assignInputField(input, "Documents", _kendraDocuments); err != nil {
			log.Errorf("invalid --documents: %s", err.Error())
			return
		}
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraCustomDocumentEnrichmentConfiguration) > 0 {
		if err := assignInputField(input, "CustomDocumentEnrichmentConfiguration", _kendraCustomDocumentEnrichmentConfiguration); err != nil {
			log.Errorf("invalid --custom-document-enrichment-configuration: %s", err.Error())
			return
		}
	}
	if len(_kendraRoleArn) > 0 {
		input.RoleArn = aws.String(_kendraRoleArn)
	}

	if resp, err := client.BatchPutDocument(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Clears existing query suggestions from an index.
// This deletes existing suggestions only, not the queries in the query log. After
// you clear suggestions, Amazon Kendra learns new suggestions based on new queries
// added to the query log from the time you cleared suggestions. If you do not see
// any new suggestions, then please allow Amazon Kendra to collect enough queries
// to learn new suggestions.
//
// ClearQuerySuggestions is currently not supported in the Amazon Web Services
// GovCloud (US-West) region.
func kendra_ClearQuerySuggestions(cfg aws.Config, client *kendra.Client) {
	input := &kendra.ClearQuerySuggestionsInput{
		// IndexId: *string, // Required
	}

	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}

	if resp, err := client.ClearQuerySuggestions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an access configuration for your documents. This includes user and
// group access information for your documents. This is useful for user context
// filtering, where search results are filtered based on the user or their group
// access to documents.
//
// You can use this to re-configure your existing document level access control
// without indexing all of your documents again. For example, your index contains
// top-secret company documents that only certain employees or users should access.
// One of these users leaves the company or switches to a team that should be
// blocked from accessing top-secret documents. The user still has access to
// top-secret documents because the user had access when your documents were
// previously indexed. You can create a specific access control configuration for
// the user with deny access. You can later update the access control configuration
// to allow access if the user returns to the company and re-joins the 'top-secret'
// team. You can re-configure access control for your documents as circumstances
// change.
//
// To apply your access control configuration to certain documents, you call the [BatchPutDocument]
// API with the AccessControlConfigurationId included in the [Document] object. If you use
// an S3 bucket as a data source, you update the .metadata.json with the
// AccessControlConfigurationId and synchronize your data source. Amazon Kendra
// currently only supports access control configuration for S3 data sources and
// documents indexed using the BatchPutDocument API.
//
// You can't configure access control using CreateAccessControlConfiguration for
// an Amazon Kendra Gen AI Enterprise Edition index. Amazon Kendra will return a
// ValidationException error for a Gen_AI_ENTERPRISE_EDITION index.
//
// [BatchPutDocument]: https://docs.aws.amazon.com/kendra/latest/dg/API_BatchPutDocument.html
// [Document]: https://docs.aws.amazon.com/kendra/latest/dg/API_Document.html
func kendra_CreateAccessControlConfiguration(cfg aws.Config, client *kendra.Client) {
	input := &kendra.CreateAccessControlConfigurationInput{
		// IndexId: *string, // Required
		// Name: *string, // Required
	}

	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraName) > 0 {
		input.Name = aws.String(_kendraName)
	}
	if len(_kendraAccessControlList) > 0 {
		if err := assignInputField(input, "AccessControlList", _kendraAccessControlList); err != nil {
			log.Errorf("invalid --access-control-list: %s", err.Error())
			return
		}
	}
	if len(_kendraClientToken) > 0 {
		input.ClientToken = aws.String(_kendraClientToken)
	}
	if len(_kendraDescription) > 0 {
		input.Description = aws.String(_kendraDescription)
	}
	if len(_kendraHierarchicalAccessControlList) > 0 {
		if err := assignInputField(input, "HierarchicalAccessControlList", _kendraHierarchicalAccessControlList); err != nil {
			log.Errorf("invalid --hierarchical-access-control-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAccessControlConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a data source connector that you want to use with an Amazon Kendra
// index.
//
// You specify a name, data source connector type and description for your data
// source. You also specify configuration information for the data source
// connector.
//
// CreateDataSource is a synchronous operation. The operation returns 200 if the
// data source was successfully created. Otherwise, an exception is raised.
//
// For an example of creating an index and data source using the Python SDK, see [Getting started with Python SDK].
// For an example of creating an index and data source using the Java SDK, see [Getting started with Java SDK].
//
// [Getting started with Python SDK]: https://docs.aws.amazon.com/kendra/latest/dg/gs-python.html
// [Getting started with Java SDK]: https://docs.aws.amazon.com/kendra/latest/dg/gs-java.html
func kendra_CreateDataSource(cfg aws.Config, client *kendra.Client) {
	input := &kendra.CreateDataSourceInput{
		// IndexId: *string, // Required
		// Name: *string, // Required
		// Type: types.DataSourceType, // Required
	}

	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraName) > 0 {
		input.Name = aws.String(_kendraName)
	}
	if len(_kendraType) > 0 {
		if err := assignInputField(input, "Type", _kendraType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_kendraClientToken) > 0 {
		input.ClientToken = aws.String(_kendraClientToken)
	}
	if len(_kendraConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _kendraConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_kendraCustomDocumentEnrichmentConfiguration) > 0 {
		if err := assignInputField(input, "CustomDocumentEnrichmentConfiguration", _kendraCustomDocumentEnrichmentConfiguration); err != nil {
			log.Errorf("invalid --custom-document-enrichment-configuration: %s", err.Error())
			return
		}
	}
	if len(_kendraDescription) > 0 {
		input.Description = aws.String(_kendraDescription)
	}
	if len(_kendraLanguageCode) > 0 {
		input.LanguageCode = aws.String(_kendraLanguageCode)
	}
	if len(_kendraRoleArn) > 0 {
		input.RoleArn = aws.String(_kendraRoleArn)
	}
	if len(_kendraSchedule) > 0 {
		input.Schedule = aws.String(_kendraSchedule)
	}
	if len(_kendraTags) > 0 {
		if err := assignInputField(input, "Tags", _kendraTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_kendraVpcConfiguration) > 0 {
		if err := assignInputField(input, "VpcConfiguration", _kendraVpcConfiguration); err != nil {
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

// Creates an Amazon Kendra experience such as a search application. For more
// information on creating a search application experience, including using the
// Python and Java SDKs, see [Building a search experience with no code].
//
// [Building a search experience with no code]: https://docs.aws.amazon.com/kendra/latest/dg/deploying-search-experience-no-code.html
func kendra_CreateExperience(cfg aws.Config, client *kendra.Client) {
	input := &kendra.CreateExperienceInput{
		// IndexId: *string, // Required
		// Name: *string, // Required
	}

	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraName) > 0 {
		input.Name = aws.String(_kendraName)
	}
	if len(_kendraClientToken) > 0 {
		input.ClientToken = aws.String(_kendraClientToken)
	}
	if len(_kendraConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _kendraConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_kendraDescription) > 0 {
		input.Description = aws.String(_kendraDescription)
	}
	if len(_kendraRoleArn) > 0 {
		input.RoleArn = aws.String(_kendraRoleArn)
	}

	if resp, err := client.CreateExperience(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a set of frequently ask questions (FAQs) using a specified FAQ file
// stored in an Amazon S3 bucket.
//
// Adding FAQs to an index is an asynchronous operation.
//
// For an example of adding an FAQ to an index using Python and Java SDKs, see [Using your FAQ file].
//
// [Using your FAQ file]: https://docs.aws.amazon.com/kendra/latest/dg/in-creating-faq.html#using-faq-file
func kendra_CreateFaq(cfg aws.Config, client *kendra.Client) {
	input := &kendra.CreateFaqInput{
		// IndexId: *string, // Required
		// Name: *string, // Required
		// RoleArn: *string, // Required
		// S3Path: *types.S3Path, // Required
	}

	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraName) > 0 {
		input.Name = aws.String(_kendraName)
	}
	if len(_kendraRoleArn) > 0 {
		input.RoleArn = aws.String(_kendraRoleArn)
	}
	if len(_kendraS3Path) > 0 {
		if err := assignInputField(input, "S3Path", _kendraS3Path); err != nil {
			log.Errorf("invalid --s3-path: %s", err.Error())
			return
		}
	}
	if len(_kendraClientToken) > 0 {
		input.ClientToken = aws.String(_kendraClientToken)
	}
	if len(_kendraDescription) > 0 {
		input.Description = aws.String(_kendraDescription)
	}
	if len(_kendraFileFormat) > 0 {
		if err := assignInputField(input, "FileFormat", _kendraFileFormat); err != nil {
			log.Errorf("invalid --file-format: %s", err.Error())
			return
		}
	}
	if len(_kendraLanguageCode) > 0 {
		input.LanguageCode = aws.String(_kendraLanguageCode)
	}
	if len(_kendraTags) > 0 {
		if err := assignInputField(input, "Tags", _kendraTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFaq(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a set of featured results to display at the top of the search results
// page. Featured results are placed above all other results for certain queries.
// You map specific queries to specific documents for featuring in the results. If
// a query contains an exact match, then one or more specific documents are
// featured in the search results.
//
// You can create up to 50 sets of featured results per index. You can request to
// increase this limit by contacting [Support].
//
// [Support]: http://aws.amazon.com/contact-us/
func kendra_CreateFeaturedResultsSet(cfg aws.Config, client *kendra.Client) {
	input := &kendra.CreateFeaturedResultsSetInput{
		// FeaturedResultsSetName: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraFeaturedResultsSetName) > 0 {
		input.FeaturedResultsSetName = aws.String(_kendraFeaturedResultsSetName)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraClientToken) > 0 {
		input.ClientToken = aws.String(_kendraClientToken)
	}
	if len(_kendraDescription) > 0 {
		input.Description = aws.String(_kendraDescription)
	}
	if len(_kendraFeaturedDocuments) > 0 {
		if err := assignInputField(input, "FeaturedDocuments", _kendraFeaturedDocuments); err != nil {
			log.Errorf("invalid --featured-documents: %s", err.Error())
			return
		}
	}
	if len(_kendraQueryTexts) > 0 {
		input.QueryTexts = append([]string(nil), _kendraQueryTexts...)
	}
	if len(_kendraStatus) > 0 {
		if err := assignInputField(input, "Status", _kendraStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_kendraTags) > 0 {
		if err := assignInputField(input, "Tags", _kendraTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFeaturedResultsSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Kendra index. Index creation is an asynchronous API. To
// determine if index creation has completed, check the Status field returned from
// a call to DescribeIndex . The Status field is set to ACTIVE when the index is
// ready to use.
//
// Once the index is active, you can index your documents using the
// BatchPutDocument API or using one of the supported [data sources].
//
// For an example of creating an index and data source using the Python SDK, see [Getting started with Python SDK].
// For an example of creating an index and data source using the Java SDK, see [Getting started with Java SDK].
//
// [data sources]: https://docs.aws.amazon.com/kendra/latest/dg/data-sources.html
// [Getting started with Python SDK]: https://docs.aws.amazon.com/kendra/latest/dg/gs-python.html
// [Getting started with Java SDK]: https://docs.aws.amazon.com/kendra/latest/dg/gs-java.html
func kendra_CreateIndex(cfg aws.Config, client *kendra.Client) {
	input := &kendra.CreateIndexInput{
		// Name: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_kendraName) > 0 {
		input.Name = aws.String(_kendraName)
	}
	if len(_kendraRoleArn) > 0 {
		input.RoleArn = aws.String(_kendraRoleArn)
	}
	if len(_kendraClientToken) > 0 {
		input.ClientToken = aws.String(_kendraClientToken)
	}
	if len(_kendraDescription) > 0 {
		input.Description = aws.String(_kendraDescription)
	}
	if len(_kendraEdition) > 0 {
		if err := assignInputField(input, "Edition", _kendraEdition); err != nil {
			log.Errorf("invalid --edition: %s", err.Error())
			return
		}
	}
	if len(_kendraServerSideEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "ServerSideEncryptionConfiguration", _kendraServerSideEncryptionConfiguration); err != nil {
			log.Errorf("invalid --server-side-encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_kendraTags) > 0 {
		if err := assignInputField(input, "Tags", _kendraTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_kendraUserContextPolicy) > 0 {
		if err := assignInputField(input, "UserContextPolicy", _kendraUserContextPolicy); err != nil {
			log.Errorf("invalid --user-context-policy: %s", err.Error())
			return
		}
	}
	if len(_kendraUserGroupResolutionConfiguration) > 0 {
		if err := assignInputField(input, "UserGroupResolutionConfiguration", _kendraUserGroupResolutionConfiguration); err != nil {
			log.Errorf("invalid --user-group-resolution-configuration: %s", err.Error())
			return
		}
	}
	if len(_kendraUserTokenConfigurations) > 0 {
		if err := assignInputField(input, "UserTokenConfigurations", _kendraUserTokenConfigurations); err != nil {
			log.Errorf("invalid --user-token-configurations: %s", err.Error())
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

// Creates a block list to exlcude certain queries from suggestions.
// Any query that contains words or phrases specified in the block list is blocked
// or filtered out from being shown as a suggestion.
//
// You need to provide the file location of your block list text file in your S3
// bucket. In your text file, enter each block word or phrase on a separate line.
//
// For information on the current quota limits for block lists, see [Quotas for Amazon Kendra].
//
// CreateQuerySuggestionsBlockList is currently not supported in the Amazon Web
// Services GovCloud (US-West) region.
//
// For an example of creating a block list for query suggestions using the Python
// SDK, see [Query suggestions block list].
//
// [Quotas for Amazon Kendra]: https://docs.aws.amazon.com/kendra/latest/dg/quotas.html
// [Query suggestions block list]: https://docs.aws.amazon.com/kendra/latest/dg/query-suggestions.html#query-suggestions-blocklist
func kendra_CreateQuerySuggestionsBlockList(cfg aws.Config, client *kendra.Client) {
	input := &kendra.CreateQuerySuggestionsBlockListInput{
		// IndexId: *string, // Required
		// Name: *string, // Required
		// RoleArn: *string, // Required
		// SourceS3Path: *types.S3Path, // Required
	}

	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraName) > 0 {
		input.Name = aws.String(_kendraName)
	}
	if len(_kendraRoleArn) > 0 {
		input.RoleArn = aws.String(_kendraRoleArn)
	}
	if len(_kendraSourceS3Path) > 0 {
		if err := assignInputField(input, "SourceS3Path", _kendraSourceS3Path); err != nil {
			log.Errorf("invalid --source-s3-path: %s", err.Error())
			return
		}
	}
	if len(_kendraClientToken) > 0 {
		input.ClientToken = aws.String(_kendraClientToken)
	}
	if len(_kendraDescription) > 0 {
		input.Description = aws.String(_kendraDescription)
	}
	if len(_kendraTags) > 0 {
		if err := assignInputField(input, "Tags", _kendraTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateQuerySuggestionsBlockList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a thesaurus for an index. The thesaurus contains a list of synonyms in
// Solr format.
//
// For an example of adding a thesaurus file to an index, see [Adding custom synonyms to an index].
//
// [Adding custom synonyms to an index]: https://docs.aws.amazon.com/kendra/latest/dg/index-synonyms-adding-thesaurus-file.html
func kendra_CreateThesaurus(cfg aws.Config, client *kendra.Client) {
	input := &kendra.CreateThesaurusInput{
		// IndexId: *string, // Required
		// Name: *string, // Required
		// RoleArn: *string, // Required
		// SourceS3Path: *types.S3Path, // Required
	}

	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraName) > 0 {
		input.Name = aws.String(_kendraName)
	}
	if len(_kendraRoleArn) > 0 {
		input.RoleArn = aws.String(_kendraRoleArn)
	}
	if len(_kendraSourceS3Path) > 0 {
		if err := assignInputField(input, "SourceS3Path", _kendraSourceS3Path); err != nil {
			log.Errorf("invalid --source-s3-path: %s", err.Error())
			return
		}
	}
	if len(_kendraClientToken) > 0 {
		input.ClientToken = aws.String(_kendraClientToken)
	}
	if len(_kendraDescription) > 0 {
		input.Description = aws.String(_kendraDescription)
	}
	if len(_kendraTags) > 0 {
		if err := assignInputField(input, "Tags", _kendraTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateThesaurus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an access control configuration that you created for your documents in
// an index. This includes user and group access information for your documents.
// This is useful for user context filtering, where search results are filtered
// based on the user or their group access to documents.
func kendra_DeleteAccessControlConfiguration(cfg aws.Config, client *kendra.Client) {
	input := &kendra.DeleteAccessControlConfigurationInput{
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}

	if resp, err := client.DeleteAccessControlConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Kendra data source connector. An exception is not thrown if
// the data source is already being deleted. While the data source is being
// deleted, the Status field returned by a call to the DescribeDataSource API is
// set to DELETING . For more information, see [Deleting Data Sources].
//
// Deleting an entire data source or re-syncing your index after deleting specific
// documents from a data source could take up to an hour or more, depending on the
// number of documents you want to delete.
//
// [Deleting Data Sources]: https://docs.aws.amazon.com/kendra/latest/dg/delete-data-source.html
func kendra_DeleteDataSource(cfg aws.Config, client *kendra.Client) {
	input := &kendra.DeleteDataSourceInput{
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}

	if resp, err := client.DeleteDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes your Amazon Kendra experience such as a search application. For more
// information on creating a search application experience, see [Building a search experience with no code].
//
// [Building a search experience with no code]: https://docs.aws.amazon.com/kendra/latest/dg/deploying-search-experience-no-code.html
func kendra_DeleteExperience(cfg aws.Config, client *kendra.Client) {
	input := &kendra.DeleteExperienceInput{
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}

	if resp, err := client.DeleteExperience(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a FAQ from an index.
func kendra_DeleteFaq(cfg aws.Config, client *kendra.Client) {
	input := &kendra.DeleteFaqInput{
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}

	if resp, err := client.DeleteFaq(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Kendra index. An exception is not thrown if the index is
// already being deleted. While the index is being deleted, the Status field
// returned by a call to the DescribeIndex API is set to DELETING .
func kendra_DeleteIndex(cfg aws.Config, client *kendra.Client) {
	input := &kendra.DeleteIndexInput{
		// Id: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}

	if resp, err := client.DeleteIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a group so that all users that belong to the group can no longer access
// documents only available to that group.
//
// For example, after deleting the group "Summer Interns", all interns who
// belonged to that group no longer see intern-only documents in their search
// results.
//
// If you want to delete or replace users or sub groups of a group, you need to
// use the PutPrincipalMapping operation. For example, if a user in the group
// "Engineering" leaves the engineering team and another user takes their place,
// you provide an updated list of users or sub groups that belong to the
// "Engineering" group when calling PutPrincipalMapping . You can update your
// internal list of users or sub groups and input this list when calling
// PutPrincipalMapping .
//
// DeletePrincipalMapping is currently not supported in the Amazon Web Services
// GovCloud (US-West) region.
func kendra_DeletePrincipalMapping(cfg aws.Config, client *kendra.Client) {
	input := &kendra.DeletePrincipalMappingInput{
		// GroupId: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraGroupId) > 0 {
		input.GroupId = aws.String(_kendraGroupId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraDataSourceId) > 0 {
		input.DataSourceId = aws.String(_kendraDataSourceId)
	}
	if len(_kendraOrderingId) > 0 {
		if err := assignInputField(input, "OrderingId", _kendraOrderingId); err != nil {
			log.Errorf("invalid --ordering-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeletePrincipalMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a block list used for query suggestions for an index.
// A deleted block list might not take effect right away. Amazon Kendra needs to
// refresh the entire suggestions list to add back the queries that were previously
// blocked.
//
// DeleteQuerySuggestionsBlockList is currently not supported in the Amazon Web
// Services GovCloud (US-West) region.
func kendra_DeleteQuerySuggestionsBlockList(cfg aws.Config, client *kendra.Client) {
	input := &kendra.DeleteQuerySuggestionsBlockListInput{
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}

	if resp, err := client.DeleteQuerySuggestionsBlockList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Kendra thesaurus.
func kendra_DeleteThesaurus(cfg aws.Config, client *kendra.Client) {
	input := &kendra.DeleteThesaurusInput{
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}

	if resp, err := client.DeleteThesaurus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an access control configuration that you created for
// your documents in an index. This includes user and group access information for
// your documents. This is useful for user context filtering, where search results
// are filtered based on the user or their group access to documents.
func kendra_DescribeAccessControlConfiguration(cfg aws.Config, client *kendra.Client) {
	input := &kendra.DescribeAccessControlConfigurationInput{
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}

	if resp, err := client.DescribeAccessControlConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an Amazon Kendra data source connector.
func kendra_DescribeDataSource(cfg aws.Config, client *kendra.Client) {
	input := &kendra.DescribeDataSourceInput{
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}

	if resp, err := client.DescribeDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about your Amazon Kendra experience such as a search
// application. For more information on creating a search application experience,
// see [Building a search experience with no code].
//
// [Building a search experience with no code]: https://docs.aws.amazon.com/kendra/latest/dg/deploying-search-experience-no-code.html
func kendra_DescribeExperience(cfg aws.Config, client *kendra.Client) {
	input := &kendra.DescribeExperienceInput{
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}

	if resp, err := client.DescribeExperience(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a FAQ.
func kendra_DescribeFaq(cfg aws.Config, client *kendra.Client) {
	input := &kendra.DescribeFaqInput{
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}

	if resp, err := client.DescribeFaq(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a set of featured results. Features results are placed
// above all other results for certain queries. If there's an exact match of a
// query, then one or more specific documents are featured in the search results.
func kendra_DescribeFeaturedResultsSet(cfg aws.Config, client *kendra.Client) {
	input := &kendra.DescribeFeaturedResultsSetInput{
		// FeaturedResultsSetId: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraFeaturedResultsSetId) > 0 {
		input.FeaturedResultsSetId = aws.String(_kendraFeaturedResultsSetId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}

	if resp, err := client.DescribeFeaturedResultsSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an Amazon Kendra index.
func kendra_DescribeIndex(cfg aws.Config, client *kendra.Client) {
	input := &kendra.DescribeIndexInput{
		// Id: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}

	if resp, err := client.DescribeIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the processing of PUT and DELETE actions for mapping users to their
// groups. This includes information on the status of actions currently processing
// or yet to be processed, when actions were last updated, when actions were
// received by Amazon Kendra, the latest action that should process and apply after
// other actions, and useful error messages if an action could not be processed.
//
// DescribePrincipalMapping is currently not supported in the Amazon Web Services
// GovCloud (US-West) region.
func kendra_DescribePrincipalMapping(cfg aws.Config, client *kendra.Client) {
	input := &kendra.DescribePrincipalMappingInput{
		// GroupId: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraGroupId) > 0 {
		input.GroupId = aws.String(_kendraGroupId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraDataSourceId) > 0 {
		input.DataSourceId = aws.String(_kendraDataSourceId)
	}

	if resp, err := client.DescribePrincipalMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a block list used for query suggestions for an index.
// This is used to check the current settings that are applied to a block list.
//
// DescribeQuerySuggestionsBlockList is currently not supported in the Amazon Web
// Services GovCloud (US-West) region.
func kendra_DescribeQuerySuggestionsBlockList(cfg aws.Config, client *kendra.Client) {
	input := &kendra.DescribeQuerySuggestionsBlockListInput{
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}

	if resp, err := client.DescribeQuerySuggestionsBlockList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information on the settings of query suggestions for an index.
// This is used to check the current settings applied to query suggestions.
//
// DescribeQuerySuggestionsConfig is currently not supported in the Amazon Web
// Services GovCloud (US-West) region.
func kendra_DescribeQuerySuggestionsConfig(cfg aws.Config, client *kendra.Client) {
	input := &kendra.DescribeQuerySuggestionsConfigInput{
		// IndexId: *string, // Required
	}

	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}

	if resp, err := client.DescribeQuerySuggestionsConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an Amazon Kendra thesaurus.
func kendra_DescribeThesaurus(cfg aws.Config, client *kendra.Client) {
	input := &kendra.DescribeThesaurusInput{
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}

	if resp, err := client.DescribeThesaurus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Prevents users or groups in your IAM Identity Center identity source from
// accessing your Amazon Kendra experience. You can create an Amazon Kendra
// experience such as a search application. For more information on creating a
// search application experience, see [Building a search experience with no code].
//
// [Building a search experience with no code]: https://docs.aws.amazon.com/kendra/latest/dg/deploying-search-experience-no-code.html
func kendra_DisassociateEntitiesFromExperience(cfg aws.Config, client *kendra.Client) {
	input := &kendra.DisassociateEntitiesFromExperienceInput{
		// EntityList: []types.EntityConfiguration, // Required
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraEntityList) > 0 {
		if err := assignInputField(input, "EntityList", _kendraEntityList); err != nil {
			log.Errorf("invalid --entity-list: %s", err.Error())
			return
		}
	}
	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}

	if resp, err := client.DisassociateEntitiesFromExperience(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specific permissions of users or groups in your IAM Identity Center
// identity source with access to your Amazon Kendra experience. You can create an
// Amazon Kendra experience such as a search application. For more information on
// creating a search application experience, see [Building a search experience with no code].
//
// [Building a search experience with no code]: https://docs.aws.amazon.com/kendra/latest/dg/deploying-search-experience-no-code.html
func kendra_DisassociatePersonasFromEntities(cfg aws.Config, client *kendra.Client) {
	input := &kendra.DisassociatePersonasFromEntitiesInput{
		// EntityIds: []string, // Required
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraEntityIds) > 0 {
		input.EntityIds = append([]string(nil), _kendraEntityIds...)
	}
	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}

	if resp, err := client.DisassociatePersonasFromEntities(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Fetches the queries that are suggested to your users.
// GetQuerySuggestions is currently not supported in the Amazon Web Services
// GovCloud (US-West) region.
func kendra_GetQuerySuggestions(cfg aws.Config, client *kendra.Client) {
	input := &kendra.GetQuerySuggestionsInput{
		// IndexId: *string, // Required
		// QueryText: *string, // Required
	}

	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraQueryText) > 0 {
		input.QueryText = aws.String(_kendraQueryText)
	}
	if len(_kendraAttributeSuggestionsConfig) > 0 {
		if err := assignInputField(input, "AttributeSuggestionsConfig", _kendraAttributeSuggestionsConfig); err != nil {
			log.Errorf("invalid --attribute-suggestions-config: %s", err.Error())
			return
		}
	}
	if len(_kendraMaxSuggestionsCount) > 0 {
		if err := assignInputField(input, "MaxSuggestionsCount", _kendraMaxSuggestionsCount); err != nil {
			log.Errorf("invalid --max-suggestions-count: %s", err.Error())
			return
		}
	}
	if len(_kendraSuggestionTypes) > 0 {
		if err := assignInputField(input, "SuggestionTypes", _kendraSuggestionTypes); err != nil {
			log.Errorf("invalid --suggestion-types: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetQuerySuggestions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves search metrics data. The data provides a snapshot of how your users
// interact with your search application and how effective the application is.
func kendra_GetSnapshots(cfg aws.Config, client *kendra.Client) {
	input := &kendra.GetSnapshotsInput{
		// IndexId: *string, // Required
		// Interval: types.Interval, // Required
		// MetricType: types.MetricType, // Required
	}

	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraInterval) > 0 {
		if err := assignInputField(input, "Interval", _kendraInterval); err != nil {
			log.Errorf("invalid --interval: %s", err.Error())
			return
		}
	}
	if len(_kendraMetricType) > 0 {
		if err := assignInputField(input, "MetricType", _kendraMetricType); err != nil {
			log.Errorf("invalid --metric-type: %s", err.Error())
			return
		}
	}
	if len(_kendraMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kendraMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kendraNextToken) > 0 {
		input.NextToken = aws.String(_kendraNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetSnapshots(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kendra.GetSnapshotsOutput
	p := kendra.NewGetSnapshotsPaginator(client, input)
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

// Lists one or more access control configurations for an index. This includes
// user and group access information for your documents. This is useful for user
// context filtering, where search results are filtered based on the user or their
// group access to documents.
func kendra_ListAccessControlConfigurations(cfg aws.Config, client *kendra.Client) {
	input := &kendra.ListAccessControlConfigurationsInput{
		// IndexId: *string, // Required
	}

	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kendraMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kendraNextToken) > 0 {
		input.NextToken = aws.String(_kendraNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccessControlConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kendra.ListAccessControlConfigurationsOutput
	p := kendra.NewListAccessControlConfigurationsPaginator(client, input)
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

// Gets statistics about synchronizing a data source connector.
func kendra_ListDataSourceSyncJobs(cfg aws.Config, client *kendra.Client) {
	input := &kendra.ListDataSourceSyncJobsInput{
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kendraMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kendraNextToken) > 0 {
		input.NextToken = aws.String(_kendraNextToken)
	}
	if len(_kendraStartTimeFilter) > 0 {
		if err := assignInputField(input, "StartTimeFilter", _kendraStartTimeFilter); err != nil {
			log.Errorf("invalid --start-time-filter: %s", err.Error())
			return
		}
	}
	if len(_kendraStatusFilter) > 0 {
		if err := assignInputField(input, "StatusFilter", _kendraStatusFilter); err != nil {
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

	var results []*kendra.ListDataSourceSyncJobsOutput
	p := kendra.NewListDataSourceSyncJobsPaginator(client, input)
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

// Lists the data source connectors that you have created.
func kendra_ListDataSources(cfg aws.Config, client *kendra.Client) {
	input := &kendra.ListDataSourcesInput{
		// IndexId: *string, // Required
	}

	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kendraMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kendraNextToken) > 0 {
		input.NextToken = aws.String(_kendraNextToken)
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

	var results []*kendra.ListDataSourcesOutput
	p := kendra.NewListDataSourcesPaginator(client, input)
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

// Lists specific permissions of users and groups with access to your Amazon
// Kendra experience.
func kendra_ListEntityPersonas(cfg aws.Config, client *kendra.Client) {
	input := &kendra.ListEntityPersonasInput{
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kendraMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kendraNextToken) > 0 {
		input.NextToken = aws.String(_kendraNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEntityPersonas(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kendra.ListEntityPersonasOutput
	p := kendra.NewListEntityPersonasPaginator(client, input)
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

// Lists users or groups in your IAM Identity Center identity source that are
// granted access to your Amazon Kendra experience. You can create an Amazon Kendra
// experience such as a search application. For more information on creating a
// search application experience, see [Building a search experience with no code].
//
// [Building a search experience with no code]: https://docs.aws.amazon.com/kendra/latest/dg/deploying-search-experience-no-code.html
func kendra_ListExperienceEntities(cfg aws.Config, client *kendra.Client) {
	input := &kendra.ListExperienceEntitiesInput{
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraNextToken) > 0 {
		input.NextToken = aws.String(_kendraNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListExperienceEntities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kendra.ListExperienceEntitiesOutput
	p := kendra.NewListExperienceEntitiesPaginator(client, input)
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

// Lists one or more Amazon Kendra experiences. You can create an Amazon Kendra
// experience such as a search application. For more information on creating a
// search application experience, see [Building a search experience with no code].
//
// [Building a search experience with no code]: https://docs.aws.amazon.com/kendra/latest/dg/deploying-search-experience-no-code.html
func kendra_ListExperiences(cfg aws.Config, client *kendra.Client) {
	input := &kendra.ListExperiencesInput{
		// IndexId: *string, // Required
	}

	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kendraMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kendraNextToken) > 0 {
		input.NextToken = aws.String(_kendraNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListExperiences(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kendra.ListExperiencesOutput
	p := kendra.NewListExperiencesPaginator(client, input)
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

// Gets a list of FAQs associated with an index.
func kendra_ListFaqs(cfg aws.Config, client *kendra.Client) {
	input := &kendra.ListFaqsInput{
		// IndexId: *string, // Required
	}

	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kendraMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kendraNextToken) > 0 {
		input.NextToken = aws.String(_kendraNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFaqs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kendra.ListFaqsOutput
	p := kendra.NewListFaqsPaginator(client, input)
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

// Lists all your sets of featured results for a given index. Features results are
// placed above all other results for certain queries. If there's an exact match of
// a query, then one or more specific documents are featured in the search results.
func kendra_ListFeaturedResultsSets(cfg aws.Config, client *kendra.Client) {
	input := &kendra.ListFeaturedResultsSetsInput{
		// IndexId: *string, // Required
	}

	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kendraMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kendraNextToken) > 0 {
		input.NextToken = aws.String(_kendraNextToken)
	}

	if resp, err := client.ListFeaturedResultsSets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a list of groups that are mapped to users before a given ordering or
// timestamp identifier.
//
// ListGroupsOlderThanOrderingId is currently not supported in the Amazon Web
// Services GovCloud (US-West) region.
func kendra_ListGroupsOlderThanOrderingId(cfg aws.Config, client *kendra.Client) {
	input := &kendra.ListGroupsOlderThanOrderingIdInput{
		// IndexId: *string, // Required
		// OrderingId: *int64, // Required
	}

	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraOrderingId) > 0 {
		if err := assignInputField(input, "OrderingId", _kendraOrderingId); err != nil {
			log.Errorf("invalid --ordering-id: %s", err.Error())
			return
		}
	}
	if len(_kendraDataSourceId) > 0 {
		input.DataSourceId = aws.String(_kendraDataSourceId)
	}
	if len(_kendraMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kendraMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kendraNextToken) > 0 {
		input.NextToken = aws.String(_kendraNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGroupsOlderThanOrderingId(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kendra.ListGroupsOlderThanOrderingIdOutput
	p := kendra.NewListGroupsOlderThanOrderingIdPaginator(client, input)
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

// Lists the Amazon Kendra indexes that you created.
func kendra_ListIndices(cfg aws.Config, client *kendra.Client) {
	input := &kendra.ListIndicesInput{}

	if len(_kendraMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kendraMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kendraNextToken) > 0 {
		input.NextToken = aws.String(_kendraNextToken)
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

	var results []*kendra.ListIndicesOutput
	p := kendra.NewListIndicesPaginator(client, input)
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

// Lists the block lists used for query suggestions for an index.
// For information on the current quota limits for block lists, see [Quotas for Amazon Kendra].
//
// ListQuerySuggestionsBlockLists is currently not supported in the Amazon Web
// Services GovCloud (US-West) region.
//
// [Quotas for Amazon Kendra]: https://docs.aws.amazon.com/kendra/latest/dg/quotas.html
func kendra_ListQuerySuggestionsBlockLists(cfg aws.Config, client *kendra.Client) {
	input := &kendra.ListQuerySuggestionsBlockListsInput{
		// IndexId: *string, // Required
	}

	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kendraMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kendraNextToken) > 0 {
		input.NextToken = aws.String(_kendraNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListQuerySuggestionsBlockLists(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kendra.ListQuerySuggestionsBlockListsOutput
	p := kendra.NewListQuerySuggestionsBlockListsPaginator(client, input)
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

// Gets a list of tags associated with a resource. Indexes, FAQs, data sources,
// and other resources can have tags associated with them.
func kendra_ListTagsForResource(cfg aws.Config, client *kendra.Client) {
	input := &kendra.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_kendraResourceARN) > 0 {
		input.ResourceARN = aws.String(_kendraResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the thesauri for an index.
func kendra_ListThesauri(cfg aws.Config, client *kendra.Client) {
	input := &kendra.ListThesauriInput{
		// IndexId: *string, // Required
	}

	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kendraMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kendraNextToken) > 0 {
		input.NextToken = aws.String(_kendraNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListThesauri(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kendra.ListThesauriOutput
	p := kendra.NewListThesauriPaginator(client, input)
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

// Maps users to their groups so that you only need to provide the user ID when
// you issue the query.
//
// You can also map sub groups to groups. For example, the group "Company
// Intellectual Property Teams" includes sub groups "Research" and "Engineering".
// These sub groups include their own list of users or people who work in these
// teams. Only users who work in research and engineering, and therefore belong in
// the intellectual property group, can see top-secret company documents in their
// search results.
//
// This is useful for user context filtering, where search results are filtered
// based on the user or their group access to documents. For more information, see [Filtering on user context]
// .
//
// If more than five PUT actions for a group are currently processing, a
// validation exception is thrown.
//
// [Filtering on user context]: https://docs.aws.amazon.com/kendra/latest/dg/user-context-filter.html
func kendra_PutPrincipalMapping(cfg aws.Config, client *kendra.Client) {
	input := &kendra.PutPrincipalMappingInput{
		// GroupId: *string, // Required
		// GroupMembers: *types.GroupMembers, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraGroupId) > 0 {
		input.GroupId = aws.String(_kendraGroupId)
	}
	if len(_kendraGroupMembers) > 0 {
		if err := assignInputField(input, "GroupMembers", _kendraGroupMembers); err != nil {
			log.Errorf("invalid --group-members: %s", err.Error())
			return
		}
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraDataSourceId) > 0 {
		input.DataSourceId = aws.String(_kendraDataSourceId)
	}
	if len(_kendraOrderingId) > 0 {
		if err := assignInputField(input, "OrderingId", _kendraOrderingId); err != nil {
			log.Errorf("invalid --ordering-id: %s", err.Error())
			return
		}
	}
	if len(_kendraRoleArn) > 0 {
		input.RoleArn = aws.String(_kendraRoleArn)
	}

	if resp, err := client.PutPrincipalMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches an index given an input query.
// If you are working with large language models (LLMs) or implementing retrieval
// augmented generation (RAG) systems, you can use Amazon Kendra's [Retrieve]API, which can
// return longer semantically relevant passages. We recommend using the Retrieve
// API instead of filing a service limit increase to increase the Query API
// document excerpt length.
//
// You can configure boosting or relevance tuning at the query level to override
// boosting at the index level, filter based on document fields/attributes and
// faceted search, and filter based on the user or their group access to documents.
// You can also include certain fields in the response that might provide useful
// additional information.
//
// A query response contains three types of results.
//
// - Relevant suggested answers. The answers can be either a text excerpt or
// table excerpt. The answer can be highlighted in the excerpt.
//
// - Matching FAQs or questions-answer from your FAQ file.
//
// - Relevant documents. This result type includes an excerpt of the document
// with the document title. The searched terms can be highlighted in the excerpt.
//
// You can specify that the query return only one type of result using the
// QueryResultTypeFilter parameter. Each query returns the 100 most relevant
// results. If you filter result type to only question-answers, a maximum of four
// results are returned. If you filter result type to only answers, a maximum of
// three results are returned.
//
// If you're using an Amazon Kendra Gen AI Enterprise Edition index, you can only
// use ATTRIBUTE_FILTER to filter search results by user context. If you're using
// an Amazon Kendra Gen AI Enterprise Edition index and you try to use USER_TOKEN
// to configure user context policy, Amazon Kendra returns a ValidationException
// error.
//
// [Retrieve]: https://docs.aws.amazon.com/kendra/latest/APIReference/API_Retrieve.html
func kendra_Query(cfg aws.Config, client *kendra.Client) {
	input := &kendra.QueryInput{
		// IndexId: *string, // Required
	}

	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraAttributeFilter) > 0 {
		if err := assignInputField(input, "AttributeFilter", _kendraAttributeFilter); err != nil {
			log.Errorf("invalid --attribute-filter: %s", err.Error())
			return
		}
	}
	if len(_kendraCollapseConfiguration) > 0 {
		if err := assignInputField(input, "CollapseConfiguration", _kendraCollapseConfiguration); err != nil {
			log.Errorf("invalid --collapse-configuration: %s", err.Error())
			return
		}
	}
	if len(_kendraDocumentRelevanceOverrideConfigurations) > 0 {
		if err := assignInputField(input, "DocumentRelevanceOverrideConfigurations", _kendraDocumentRelevanceOverrideConfigurations); err != nil {
			log.Errorf("invalid --document-relevance-override-configurations: %s", err.Error())
			return
		}
	}
	if len(_kendraFacets) > 0 {
		if err := assignInputField(input, "Facets", _kendraFacets); err != nil {
			log.Errorf("invalid --facets: %s", err.Error())
			return
		}
	}
	if len(_kendraPageNumber) > 0 {
		if err := assignInputField(input, "PageNumber", _kendraPageNumber); err != nil {
			log.Errorf("invalid --page-number: %s", err.Error())
			return
		}
	}
	if len(_kendraPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _kendraPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_kendraQueryResultTypeFilter) > 0 {
		if err := assignInputField(input, "QueryResultTypeFilter", _kendraQueryResultTypeFilter); err != nil {
			log.Errorf("invalid --query-result-type-filter: %s", err.Error())
			return
		}
	}
	if len(_kendraQueryText) > 0 {
		input.QueryText = aws.String(_kendraQueryText)
	}
	if len(_kendraRequestedDocumentAttributes) > 0 {
		input.RequestedDocumentAttributes = append([]string(nil), _kendraRequestedDocumentAttributes...)
	}
	if len(_kendraSortingConfiguration) > 0 {
		if err := assignInputField(input, "SortingConfiguration", _kendraSortingConfiguration); err != nil {
			log.Errorf("invalid --sorting-configuration: %s", err.Error())
			return
		}
	}
	if len(_kendraSortingConfigurations) > 0 {
		if err := assignInputField(input, "SortingConfigurations", _kendraSortingConfigurations); err != nil {
			log.Errorf("invalid --sorting-configurations: %s", err.Error())
			return
		}
	}
	if len(_kendraSpellCorrectionConfiguration) > 0 {
		if err := assignInputField(input, "SpellCorrectionConfiguration", _kendraSpellCorrectionConfiguration); err != nil {
			log.Errorf("invalid --spell-correction-configuration: %s", err.Error())
			return
		}
	}
	if len(_kendraUserContext) > 0 {
		if err := assignInputField(input, "UserContext", _kendraUserContext); err != nil {
			log.Errorf("invalid --user-context: %s", err.Error())
			return
		}
	}
	if len(_kendraVisitorId) > 0 {
		input.VisitorId = aws.String(_kendraVisitorId)
	}

	if resp, err := client.Query(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves relevant passages or text excerpts given an input query.
// This API is similar to the [Query] API. However, by default, the Query API only
// returns excerpt passages of up to 100 token words. With the Retrieve API, you
// can retrieve longer passages of up to 200 token words and up to 100 semantically
// relevant passages. This doesn't include question-answer or FAQ type responses
// from your index. The passages are text excerpts that can be semantically
// extracted from multiple documents and multiple parts of the same document. If in
// extreme cases your documents produce zero passages using the Retrieve API, you
// can alternatively use the Query API and its types of responses.
//
// You can also do the following:
//
// - Override boosting at the index level
//
// - Filter based on document fields or attributes
//
// - Filter based on the user or their group access to documents
//
// - View the confidence score bucket for a retrieved passage result. The
// confidence bucket provides a relative ranking that indicates how confident
// Amazon Kendra is that the response is relevant to the query.
//
// Confidence score buckets are currently available only for English.
//
// You can also include certain fields in the response that might provide useful
// additional information.
//
// The Retrieve API shares the number of [query capacity units] that you set for your index. For more
// information on what's included in a single capacity unit and the default base
// capacity for an index, see [Adjusting capacity].
//
// If you're using an Amazon Kendra Gen AI Enterprise Edition index, you can only
// use ATTRIBUTE_FILTER to filter search results by user context. If you're using
// an Amazon Kendra Gen AI Enterprise Edition index and you try to use USER_TOKEN
// to configure user context policy, Amazon Kendra returns a ValidationException
// error.
//
// [Adjusting capacity]: https://docs.aws.amazon.com/kendra/latest/dg/adjusting-capacity.html
// [Query]: https://docs.aws.amazon.com/kendra/latest/APIReference/API_Query.html
// [query capacity units]: https://docs.aws.amazon.com/kendra/latest/APIReference/API_CapacityUnitsConfiguration.html
func kendra_Retrieve(cfg aws.Config, client *kendra.Client) {
	input := &kendra.RetrieveInput{
		// IndexId: *string, // Required
		// QueryText: *string, // Required
	}

	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraQueryText) > 0 {
		input.QueryText = aws.String(_kendraQueryText)
	}
	if len(_kendraAttributeFilter) > 0 {
		if err := assignInputField(input, "AttributeFilter", _kendraAttributeFilter); err != nil {
			log.Errorf("invalid --attribute-filter: %s", err.Error())
			return
		}
	}
	if len(_kendraDocumentRelevanceOverrideConfigurations) > 0 {
		if err := assignInputField(input, "DocumentRelevanceOverrideConfigurations", _kendraDocumentRelevanceOverrideConfigurations); err != nil {
			log.Errorf("invalid --document-relevance-override-configurations: %s", err.Error())
			return
		}
	}
	if len(_kendraPageNumber) > 0 {
		if err := assignInputField(input, "PageNumber", _kendraPageNumber); err != nil {
			log.Errorf("invalid --page-number: %s", err.Error())
			return
		}
	}
	if len(_kendraPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _kendraPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_kendraRequestedDocumentAttributes) > 0 {
		input.RequestedDocumentAttributes = append([]string(nil), _kendraRequestedDocumentAttributes...)
	}
	if len(_kendraUserContext) > 0 {
		if err := assignInputField(input, "UserContext", _kendraUserContext); err != nil {
			log.Errorf("invalid --user-context: %s", err.Error())
			return
		}
	}

	if resp, err := client.Retrieve(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a synchronization job for a data source connector. If a synchronization
// job is already in progress, Amazon Kendra returns a ResourceInUseException
// exception.
//
// Re-syncing your data source with your index after modifying, adding, or
// deleting documents from your data source respository could take up to an hour or
// more, depending on the number of documents to sync.
func kendra_StartDataSourceSyncJob(cfg aws.Config, client *kendra.Client) {
	input := &kendra.StartDataSourceSyncJobInput{
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}

	if resp, err := client.StartDataSourceSyncJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a synchronization job that is currently running. You can't stop a
// scheduled synchronization job.
func kendra_StopDataSourceSyncJob(cfg aws.Config, client *kendra.Client) {
	input := &kendra.StopDataSourceSyncJobInput{
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}

	if resp, err := client.StopDataSourceSyncJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables you to provide feedback to Amazon Kendra to improve the performance of
// your index.
//
// SubmitFeedback is currently not supported in the Amazon Web Services GovCloud
// (US-West) region.
func kendra_SubmitFeedback(cfg aws.Config, client *kendra.Client) {
	input := &kendra.SubmitFeedbackInput{
		// IndexId: *string, // Required
		// QueryId: *string, // Required
	}

	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraQueryId) > 0 {
		input.QueryId = aws.String(_kendraQueryId)
	}
	if len(_kendraClickFeedbackItems) > 0 {
		if err := assignInputField(input, "ClickFeedbackItems", _kendraClickFeedbackItems); err != nil {
			log.Errorf("invalid --click-feedback-items: %s", err.Error())
			return
		}
	}
	if len(_kendraRelevanceFeedbackItems) > 0 {
		if err := assignInputField(input, "RelevanceFeedbackItems", _kendraRelevanceFeedbackItems); err != nil {
			log.Errorf("invalid --relevance-feedback-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.SubmitFeedback(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the specified tag to the specified index, FAQ, data source, or other
// resource. If the tag already exists, the existing value is replaced with the new
// value.
func kendra_TagResource(cfg aws.Config, client *kendra.Client) {
	input := &kendra.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_kendraResourceARN) > 0 {
		input.ResourceARN = aws.String(_kendraResourceARN)
	}
	if len(_kendraTags) > 0 {
		if err := assignInputField(input, "Tags", _kendraTags); err != nil {
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

// Removes a tag from an index, FAQ, data source, or other resource.
func kendra_UntagResource(cfg aws.Config, client *kendra.Client) {
	input := &kendra.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_kendraResourceARN) > 0 {
		input.ResourceARN = aws.String(_kendraResourceARN)
	}
	if len(_kendraTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _kendraTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an access control configuration for your documents in an index. This
// includes user and group access information for your documents. This is useful
// for user context filtering, where search results are filtered based on the user
// or their group access to documents.
//
// You can update an access control configuration you created without indexing all
// of your documents again. For example, your index contains top-secret company
// documents that only certain employees or users should access. You created an
// 'allow' access control configuration for one user who recently joined the
// 'top-secret' team, switching from a team with 'deny' access to top-secret
// documents. However, the user suddenly returns to their previous team and should
// no longer have access to top secret documents. You can update the access control
// configuration to re-configure access control for your documents as circumstances
// change.
//
// You call the [BatchPutDocument] API to apply the updated access control configuration, with the
// AccessControlConfigurationId included in the [Document] object. If you use an S3 bucket
// as a data source, you synchronize your data source to apply the
// AccessControlConfigurationId in the .metadata.json file. Amazon Kendra
// currently only supports access control configuration for S3 data sources and
// documents indexed using the BatchPutDocument API.
//
// You can't configure access control using CreateAccessControlConfiguration for
// an Amazon Kendra Gen AI Enterprise Edition index. Amazon Kendra will return a
// ValidationException error for a Gen_AI_ENTERPRISE_EDITION index.
//
// [BatchPutDocument]: https://docs.aws.amazon.com/kendra/latest/dg/API_BatchPutDocument.html
// [Document]: https://docs.aws.amazon.com/kendra/latest/dg/API_Document.html
func kendra_UpdateAccessControlConfiguration(cfg aws.Config, client *kendra.Client) {
	input := &kendra.UpdateAccessControlConfigurationInput{
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraAccessControlList) > 0 {
		if err := assignInputField(input, "AccessControlList", _kendraAccessControlList); err != nil {
			log.Errorf("invalid --access-control-list: %s", err.Error())
			return
		}
	}
	if len(_kendraDescription) > 0 {
		input.Description = aws.String(_kendraDescription)
	}
	if len(_kendraHierarchicalAccessControlList) > 0 {
		if err := assignInputField(input, "HierarchicalAccessControlList", _kendraHierarchicalAccessControlList); err != nil {
			log.Errorf("invalid --hierarchical-access-control-list: %s", err.Error())
			return
		}
	}
	if len(_kendraName) > 0 {
		input.Name = aws.String(_kendraName)
	}

	if resp, err := client.UpdateAccessControlConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Amazon Kendra data source connector.
func kendra_UpdateDataSource(cfg aws.Config, client *kendra.Client) {
	input := &kendra.UpdateDataSourceInput{
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _kendraConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_kendraCustomDocumentEnrichmentConfiguration) > 0 {
		if err := assignInputField(input, "CustomDocumentEnrichmentConfiguration", _kendraCustomDocumentEnrichmentConfiguration); err != nil {
			log.Errorf("invalid --custom-document-enrichment-configuration: %s", err.Error())
			return
		}
	}
	if len(_kendraDescription) > 0 {
		input.Description = aws.String(_kendraDescription)
	}
	if len(_kendraLanguageCode) > 0 {
		input.LanguageCode = aws.String(_kendraLanguageCode)
	}
	if len(_kendraName) > 0 {
		input.Name = aws.String(_kendraName)
	}
	if len(_kendraRoleArn) > 0 {
		input.RoleArn = aws.String(_kendraRoleArn)
	}
	if len(_kendraSchedule) > 0 {
		input.Schedule = aws.String(_kendraSchedule)
	}
	if len(_kendraVpcConfiguration) > 0 {
		if err := assignInputField(input, "VpcConfiguration", _kendraVpcConfiguration); err != nil {
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

// Updates your Amazon Kendra experience such as a search application. For more
// information on creating a search application experience, see [Building a search experience with no code].
//
// [Building a search experience with no code]: https://docs.aws.amazon.com/kendra/latest/dg/deploying-search-experience-no-code.html
func kendra_UpdateExperience(cfg aws.Config, client *kendra.Client) {
	input := &kendra.UpdateExperienceInput{
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _kendraConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_kendraDescription) > 0 {
		input.Description = aws.String(_kendraDescription)
	}
	if len(_kendraName) > 0 {
		input.Name = aws.String(_kendraName)
	}
	if len(_kendraRoleArn) > 0 {
		input.RoleArn = aws.String(_kendraRoleArn)
	}

	if resp, err := client.UpdateExperience(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a set of featured results. Features results are placed above all other
// results for certain queries. You map specific queries to specific documents for
// featuring in the results. If a query contains an exact match of a query, then
// one or more specific documents are featured in the search results.
func kendra_UpdateFeaturedResultsSet(cfg aws.Config, client *kendra.Client) {
	input := &kendra.UpdateFeaturedResultsSetInput{
		// FeaturedResultsSetId: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraFeaturedResultsSetId) > 0 {
		input.FeaturedResultsSetId = aws.String(_kendraFeaturedResultsSetId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraDescription) > 0 {
		input.Description = aws.String(_kendraDescription)
	}
	if len(_kendraFeaturedDocuments) > 0 {
		if err := assignInputField(input, "FeaturedDocuments", _kendraFeaturedDocuments); err != nil {
			log.Errorf("invalid --featured-documents: %s", err.Error())
			return
		}
	}
	if len(_kendraFeaturedResultsSetName) > 0 {
		input.FeaturedResultsSetName = aws.String(_kendraFeaturedResultsSetName)
	}
	if len(_kendraQueryTexts) > 0 {
		input.QueryTexts = append([]string(nil), _kendraQueryTexts...)
	}
	if len(_kendraStatus) > 0 {
		if err := assignInputField(input, "Status", _kendraStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFeaturedResultsSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Amazon Kendra index.
func kendra_UpdateIndex(cfg aws.Config, client *kendra.Client) {
	input := &kendra.UpdateIndexInput{
		// Id: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraCapacityUnits) > 0 {
		if err := assignInputField(input, "CapacityUnits", _kendraCapacityUnits); err != nil {
			log.Errorf("invalid --capacity-units: %s", err.Error())
			return
		}
	}
	if len(_kendraDescription) > 0 {
		input.Description = aws.String(_kendraDescription)
	}
	if len(_kendraDocumentMetadataConfigurationUpdates) > 0 {
		if err := assignInputField(input, "DocumentMetadataConfigurationUpdates", _kendraDocumentMetadataConfigurationUpdates); err != nil {
			log.Errorf("invalid --document-metadata-configuration-updates: %s", err.Error())
			return
		}
	}
	if len(_kendraName) > 0 {
		input.Name = aws.String(_kendraName)
	}
	if len(_kendraRoleArn) > 0 {
		input.RoleArn = aws.String(_kendraRoleArn)
	}
	if len(_kendraUserContextPolicy) > 0 {
		if err := assignInputField(input, "UserContextPolicy", _kendraUserContextPolicy); err != nil {
			log.Errorf("invalid --user-context-policy: %s", err.Error())
			return
		}
	}
	if len(_kendraUserGroupResolutionConfiguration) > 0 {
		if err := assignInputField(input, "UserGroupResolutionConfiguration", _kendraUserGroupResolutionConfiguration); err != nil {
			log.Errorf("invalid --user-group-resolution-configuration: %s", err.Error())
			return
		}
	}
	if len(_kendraUserTokenConfigurations) > 0 {
		if err := assignInputField(input, "UserTokenConfigurations", _kendraUserTokenConfigurations); err != nil {
			log.Errorf("invalid --user-token-configurations: %s", err.Error())
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

// Updates a block list used for query suggestions for an index.
// Updates to a block list might not take effect right away. Amazon Kendra needs
// to refresh the entire suggestions list to apply any updates to the block list.
// Other changes not related to the block list apply immediately.
//
// If a block list is updating, then you need to wait for the first update to
// finish before submitting another update.
//
// Amazon Kendra supports partial updates, so you only need to provide the fields
// you want to update.
//
// UpdateQuerySuggestionsBlockList is currently not supported in the Amazon Web
// Services GovCloud (US-West) region.
func kendra_UpdateQuerySuggestionsBlockList(cfg aws.Config, client *kendra.Client) {
	input := &kendra.UpdateQuerySuggestionsBlockListInput{
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraDescription) > 0 {
		input.Description = aws.String(_kendraDescription)
	}
	if len(_kendraName) > 0 {
		input.Name = aws.String(_kendraName)
	}
	if len(_kendraRoleArn) > 0 {
		input.RoleArn = aws.String(_kendraRoleArn)
	}
	if len(_kendraSourceS3Path) > 0 {
		if err := assignInputField(input, "SourceS3Path", _kendraSourceS3Path); err != nil {
			log.Errorf("invalid --source-s3-path: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateQuerySuggestionsBlockList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the settings of query suggestions for an index.
// Amazon Kendra supports partial updates, so you only need to provide the fields
// you want to update.
//
// If an update is currently processing, you need to wait for the update to finish
// before making another update.
//
// Updates to query suggestions settings might not take effect right away. The
// time for your updated settings to take effect depends on the updates made and
// the number of search queries in your index.
//
// You can still enable/disable query suggestions at any time.
//
// UpdateQuerySuggestionsConfig is currently not supported in the Amazon Web
// Services GovCloud (US-West) region.
func kendra_UpdateQuerySuggestionsConfig(cfg aws.Config, client *kendra.Client) {
	input := &kendra.UpdateQuerySuggestionsConfigInput{
		// IndexId: *string, // Required
	}

	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraAttributeSuggestionsConfig) > 0 {
		if err := assignInputField(input, "AttributeSuggestionsConfig", _kendraAttributeSuggestionsConfig); err != nil {
			log.Errorf("invalid --attribute-suggestions-config: %s", err.Error())
			return
		}
	}
	if len(_kendraIncludeQueriesWithoutUserInformation) > 0 {
		if err := assignInputField(input, "IncludeQueriesWithoutUserInformation", _kendraIncludeQueriesWithoutUserInformation); err != nil {
			log.Errorf("invalid --include-queries-without-user-information: %s", err.Error())
			return
		}
	}
	if len(_kendraMinimumNumberOfQueryingUsers) > 0 {
		if err := assignInputField(input, "MinimumNumberOfQueryingUsers", _kendraMinimumNumberOfQueryingUsers); err != nil {
			log.Errorf("invalid --minimum-number-of-querying-users: %s", err.Error())
			return
		}
	}
	if len(_kendraMinimumQueryCount) > 0 {
		if err := assignInputField(input, "MinimumQueryCount", _kendraMinimumQueryCount); err != nil {
			log.Errorf("invalid --minimum-query-count: %s", err.Error())
			return
		}
	}
	if len(_kendraMode) > 0 {
		if err := assignInputField(input, "Mode", _kendraMode); err != nil {
			log.Errorf("invalid --mode: %s", err.Error())
			return
		}
	}
	if len(_kendraQueryLogLookBackWindowInDays) > 0 {
		if err := assignInputField(input, "QueryLogLookBackWindowInDays", _kendraQueryLogLookBackWindowInDays); err != nil {
			log.Errorf("invalid --query-log-look-back-window-in-days: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateQuerySuggestionsConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a thesaurus for an index.
func kendra_UpdateThesaurus(cfg aws.Config, client *kendra.Client) {
	input := &kendra.UpdateThesaurusInput{
		// Id: *string, // Required
		// IndexId: *string, // Required
	}

	if len(_kendraId) > 0 {
		input.Id = aws.String(_kendraId)
	}
	if len(_kendraIndexId) > 0 {
		input.IndexId = aws.String(_kendraIndexId)
	}
	if len(_kendraDescription) > 0 {
		input.Description = aws.String(_kendraDescription)
	}
	if len(_kendraName) > 0 {
		input.Name = aws.String(_kendraName)
	}
	if len(_kendraRoleArn) > 0 {
		input.RoleArn = aws.String(_kendraRoleArn)
	}
	if len(_kendraSourceS3Path) > 0 {
		if err := assignInputField(input, "SourceS3Path", _kendraSourceS3Path); err != nil {
			log.Errorf("invalid --source-s3-path: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateThesaurus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_kendraCmd)
	_kendraCmd.Flags().SortFlags = false

	_kendraCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_kendraCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_kendraCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_kendraCmd.Flags().StringVarP(&_kendraAccessControlList, "access-control-list", "", "", "Access Control List")
	_kendraCmd.Flags().StringVarP(&_kendraAttributeFilter, "attribute-filter", "", "", "Attribute Filter")
	_kendraCmd.Flags().StringVarP(&_kendraAttributeSuggestionsConfig, "attribute-suggestions-config", "", "", "Attribute Suggestions Config")
	_kendraCmd.Flags().StringVarP(&_kendraCapacityUnits, "capacity-units", "", "", "Capacity Units")
	_kendraCmd.Flags().StringVarP(&_kendraClickFeedbackItems, "click-feedback-items", "", "", "Click Feedback Items")
	_kendraCmd.Flags().StringVarP(&_kendraClientToken, "client-token", "", "", "Client Token")
	_kendraCmd.Flags().StringVarP(&_kendraCollapseConfiguration, "collapse-configuration", "", "", "Collapse Configuration")
	_kendraCmd.Flags().StringVarP(&_kendraConfiguration, "configuration", "", "", "Configuration")
	_kendraCmd.Flags().StringVarP(&_kendraCustomDocumentEnrichmentConfiguration, "custom-document-enrichment-configuration", "", "", "Custom Document Enrichment Configuration")
	_kendraCmd.Flags().StringVarP(&_kendraDataSourceId, "data-source-id", "", "", "Data Source ID")
	_kendraCmd.Flags().StringVarP(&_kendraDataSourceSyncJobMetricTarget, "data-source-sync-job-metric-target", "", "", "Data Source Sync Job Metric Target")
	_kendraCmd.Flags().StringVarP(&_kendraDescription, "description", "", "", "Description")
	_kendraCmd.Flags().StringSliceVarP(&_kendraDocumentIdList, "document-id-list", "", nil, "Document ID List")
	_kendraCmd.Flags().StringVarP(&_kendraDocumentInfoList, "document-info-list", "", "", "Document Info List")
	_kendraCmd.Flags().StringVarP(&_kendraDocumentMetadataConfigurationUpdates, "document-metadata-configuration-updates", "", "", "Document Metadata Configuration Updates")
	_kendraCmd.Flags().StringVarP(&_kendraDocumentRelevanceOverrideConfigurations, "document-relevance-override-configurations", "", "", "Document Relevance Override Configurations")
	_kendraCmd.Flags().StringVarP(&_kendraDocuments, "documents", "", "", "Documents")
	_kendraCmd.Flags().StringVarP(&_kendraEdition, "edition", "", "", "Edition")
	_kendraCmd.Flags().StringSliceVarP(&_kendraEntityIds, "entity-ids", "", nil, "Entity Ids")
	_kendraCmd.Flags().StringVarP(&_kendraEntityList, "entity-list", "", "", "Entity List")
	_kendraCmd.Flags().StringVarP(&_kendraFacets, "facets", "", "", "Facets")
	_kendraCmd.Flags().StringVarP(&_kendraFeaturedDocuments, "featured-documents", "", "", "Featured Documents")
	_kendraCmd.Flags().StringVarP(&_kendraFeaturedResultsSetId, "featured-results-set-id", "", "", "Featured Results Set ID")
	_kendraCmd.Flags().StringSliceVarP(&_kendraFeaturedResultsSetIds, "featured-results-set-ids", "", nil, "Featured Results Set Ids")
	_kendraCmd.Flags().StringVarP(&_kendraFeaturedResultsSetName, "featured-results-set-name", "", "", "Featured Results Set Name")
	_kendraCmd.Flags().StringVarP(&_kendraFileFormat, "file-format", "", "", "File Format")
	_kendraCmd.Flags().StringVarP(&_kendraGroupId, "group-id", "", "", "Group ID")
	_kendraCmd.Flags().StringVarP(&_kendraGroupMembers, "group-members", "", "", "Group Members")
	_kendraCmd.Flags().StringVarP(&_kendraHierarchicalAccessControlList, "hierarchical-access-control-list", "", "", "Hierarchical Access Control List")
	_kendraCmd.Flags().StringVarP(&_kendraId, "id", "", "", "ID")
	_kendraCmd.Flags().StringVarP(&_kendraIncludeQueriesWithoutUserInformation, "include-queries-without-user-information", "", "", "Include Queries Without User Information")
	_kendraCmd.Flags().StringVarP(&_kendraIndexId, "index-id", "", "", "Index ID")
	_kendraCmd.Flags().StringVarP(&_kendraInterval, "interval", "", "", "Interval")
	_kendraCmd.Flags().StringVarP(&_kendraLanguageCode, "language-code", "", "", "Language Code")
	_kendraCmd.Flags().StringVarP(&_kendraMaxResults, "max-results", "", "", "Max Results")
	_kendraCmd.Flags().StringVarP(&_kendraMaxSuggestionsCount, "max-suggestions-count", "", "", "Max Suggestions Count")
	_kendraCmd.Flags().StringVarP(&_kendraMetricType, "metric-type", "", "", "Metric Type")
	_kendraCmd.Flags().StringVarP(&_kendraMinimumNumberOfQueryingUsers, "minimum-number-of-querying-users", "", "", "Minimum Number Of Querying Users")
	_kendraCmd.Flags().StringVarP(&_kendraMinimumQueryCount, "minimum-query-count", "", "", "Minimum Query Count")
	_kendraCmd.Flags().StringVarP(&_kendraMode, "mode", "", "", "Mode")
	_kendraCmd.Flags().StringVarP(&_kendraName, "name", "", "", "Name")
	_kendraCmd.Flags().StringVarP(&_kendraNextToken, "next-token", "", "", "Next Token")
	_kendraCmd.Flags().StringVarP(&_kendraOrderingId, "ordering-id", "", "", "Ordering ID")
	_kendraCmd.Flags().StringVarP(&_kendraPageNumber, "page-number", "", "", "Page Number")
	_kendraCmd.Flags().StringVarP(&_kendraPageSize, "page-size", "", "", "Page Size")
	_kendraCmd.Flags().StringVarP(&_kendraPersonas, "personas", "", "", "Personas")
	_kendraCmd.Flags().StringVarP(&_kendraQueryId, "query-id", "", "", "Query ID")
	_kendraCmd.Flags().StringVarP(&_kendraQueryLogLookBackWindowInDays, "query-log-look-back-window-in-days", "", "", "Query Log Look Back Window In Days")
	_kendraCmd.Flags().StringVarP(&_kendraQueryResultTypeFilter, "query-result-type-filter", "", "", "Query Result Type Filter")
	_kendraCmd.Flags().StringVarP(&_kendraQueryText, "query-text", "", "", "Query Text")
	_kendraCmd.Flags().StringSliceVarP(&_kendraQueryTexts, "query-texts", "", nil, "Query Texts")
	_kendraCmd.Flags().StringVarP(&_kendraRelevanceFeedbackItems, "relevance-feedback-items", "", "", "Relevance Feedback Items")
	_kendraCmd.Flags().StringSliceVarP(&_kendraRequestedDocumentAttributes, "requested-document-attributes", "", nil, "Requested Document Attributes")
	_kendraCmd.Flags().StringVarP(&_kendraResourceARN, "resource-arn", "", "", "Resource ARN")
	_kendraCmd.Flags().StringVarP(&_kendraRoleArn, "role-arn", "", "", "Role ARN")
	_kendraCmd.Flags().StringVarP(&_kendraS3Path, "s3-path", "", "", "S3 Path")
	_kendraCmd.Flags().StringVarP(&_kendraSchedule, "schedule", "", "", "Schedule")
	_kendraCmd.Flags().StringVarP(&_kendraServerSideEncryptionConfiguration, "server-side-encryption-configuration", "", "", "Server Side Encryption Configuration")
	_kendraCmd.Flags().StringVarP(&_kendraSortingConfiguration, "sorting-configuration", "", "", "Sorting Configuration")
	_kendraCmd.Flags().StringVarP(&_kendraSortingConfigurations, "sorting-configurations", "", "", "Sorting Configurations")
	_kendraCmd.Flags().StringVarP(&_kendraSourceS3Path, "source-s3-path", "", "", "Source S3 Path")
	_kendraCmd.Flags().StringVarP(&_kendraSpellCorrectionConfiguration, "spell-correction-configuration", "", "", "Spell Correction Configuration")
	_kendraCmd.Flags().StringVarP(&_kendraStartTimeFilter, "start-time-filter", "", "", "Start Time Filter")
	_kendraCmd.Flags().StringVarP(&_kendraStatus, "status", "", "", "Status")
	_kendraCmd.Flags().StringVarP(&_kendraStatusFilter, "status-filter", "", "", "Status Filter")
	_kendraCmd.Flags().StringVarP(&_kendraSuggestionTypes, "suggestion-types", "", "", "Suggestion Types")
	_kendraCmd.Flags().StringSliceVarP(&_kendraTagKeys, "tag-keys", "", nil, "Tag Keys")
	_kendraCmd.Flags().StringVarP(&_kendraTags, "tags", "", "", "Tags")
	_kendraCmd.Flags().StringVarP(&_kendraType, "type", "", "", "Type")
	_kendraCmd.Flags().StringVarP(&_kendraUserContext, "user-context", "", "", "User Context")
	_kendraCmd.Flags().StringVarP(&_kendraUserContextPolicy, "user-context-policy", "", "", "User Context Policy")
	_kendraCmd.Flags().StringVarP(&_kendraUserGroupResolutionConfiguration, "user-group-resolution-configuration", "", "", "User Group Resolution Configuration")
	_kendraCmd.Flags().StringVarP(&_kendraUserTokenConfigurations, "user-token-configurations", "", "", "User Token Configurations")
	_kendraCmd.Flags().StringVarP(&_kendraVisitorId, "visitor-id", "", "", "Visitor ID")
	_kendraCmd.Flags().StringVarP(&_kendraVpcConfiguration, "vpc-configuration", "", "", "VPC Configuration")

	_kendraCmd.Flags().BoolVarP(&_kendraAssociateEntitiesToExperience, "associate-entities-to-experience", "", false, "Associate Entities To Experience")
	_kendraCmd.Flags().BoolVarP(&_kendraAssociatePersonasToEntities, "associate-personas-to-entities", "", false, "Associate Personas To Entities")
	_kendraCmd.Flags().BoolVarP(&_kendraBatchDeleteDocument, "batch-delete-document", "", false, "Batch Delete Document")
	_kendraCmd.Flags().BoolVarP(&_kendraBatchDeleteFeaturedResultsSet, "batch-delete-featured-results-set", "", false, "Batch Delete Featured Results Set")
	_kendraCmd.Flags().BoolVarP(&_kendraBatchGetDocumentStatus, "batch-get-document-status", "", false, "Batch Get Document Status")
	_kendraCmd.Flags().BoolVarP(&_kendraBatchPutDocument, "batch-put-document", "", false, "Batch Put Document")
	_kendraCmd.Flags().BoolVarP(&_kendraClearQuerySuggestions, "clear-query-suggestions", "", false, "Clear Query Suggestions")
	_kendraCmd.Flags().BoolVarP(&_kendraCreateAccessControlConfiguration, "create-access-control-configuration", "", false, "Create Access Control Configuration")
	_kendraCmd.Flags().BoolVarP(&_kendraCreateDataSource, "create-data-source", "", false, "Create Data Source")
	_kendraCmd.Flags().BoolVarP(&_kendraCreateExperience, "create-experience", "", false, "Create Experience")
	_kendraCmd.Flags().BoolVarP(&_kendraCreateFaq, "create-faq", "", false, "Create Faq")
	_kendraCmd.Flags().BoolVarP(&_kendraCreateFeaturedResultsSet, "create-featured-results-set", "", false, "Create Featured Results Set")
	_kendraCmd.Flags().BoolVarP(&_kendraCreateIndex, "create-index", "", false, "Create Index")
	_kendraCmd.Flags().BoolVarP(&_kendraCreateQuerySuggestionsBlockList, "create-query-suggestions-block-list", "", false, "Create Query Suggestions Block List")
	_kendraCmd.Flags().BoolVarP(&_kendraCreateThesaurus, "create-thesaurus", "", false, "Create Thesaurus")
	_kendraCmd.Flags().BoolVarP(&_kendraDeleteAccessControlConfiguration, "delete-access-control-configuration", "", false, "Delete Access Control Configuration")
	_kendraCmd.Flags().BoolVarP(&_kendraDeleteDataSource, "delete-data-source", "", false, "Delete Data Source")
	_kendraCmd.Flags().BoolVarP(&_kendraDeleteExperience, "delete-experience", "", false, "Delete Experience")
	_kendraCmd.Flags().BoolVarP(&_kendraDeleteFaq, "delete-faq", "", false, "Delete Faq")
	_kendraCmd.Flags().BoolVarP(&_kendraDeleteIndex, "delete-index", "", false, "Delete Index")
	_kendraCmd.Flags().BoolVarP(&_kendraDeletePrincipalMapping, "delete-principal-mapping", "", false, "Delete Principal Mapping")
	_kendraCmd.Flags().BoolVarP(&_kendraDeleteQuerySuggestionsBlockList, "delete-query-suggestions-block-list", "", false, "Delete Query Suggestions Block List")
	_kendraCmd.Flags().BoolVarP(&_kendraDeleteThesaurus, "delete-thesaurus", "", false, "Delete Thesaurus")
	_kendraCmd.Flags().BoolVarP(&_kendraDescribeAccessControlConfiguration, "describe-access-control-configuration", "", false, "Describe Access Control Configuration")
	_kendraCmd.Flags().BoolVarP(&_kendraDescribeDataSource, "describe-data-source", "", false, "Describe Data Source")
	_kendraCmd.Flags().BoolVarP(&_kendraDescribeExperience, "describe-experience", "", false, "Describe Experience")
	_kendraCmd.Flags().BoolVarP(&_kendraDescribeFaq, "describe-faq", "", false, "Describe Faq")
	_kendraCmd.Flags().BoolVarP(&_kendraDescribeFeaturedResultsSet, "describe-featured-results-set", "", false, "Describe Featured Results Set")
	_kendraCmd.Flags().BoolVarP(&_kendraDescribeIndex, "describe-index", "", false, "Describe Index")
	_kendraCmd.Flags().BoolVarP(&_kendraDescribePrincipalMapping, "describe-principal-mapping", "", false, "Describe Principal Mapping")
	_kendraCmd.Flags().BoolVarP(&_kendraDescribeQuerySuggestionsBlockList, "describe-query-suggestions-block-list", "", false, "Describe Query Suggestions Block List")
	_kendraCmd.Flags().BoolVarP(&_kendraDescribeQuerySuggestionsConfig, "describe-query-suggestions-config", "", false, "Describe Query Suggestions Config")
	_kendraCmd.Flags().BoolVarP(&_kendraDescribeThesaurus, "describe-thesaurus", "", false, "Describe Thesaurus")
	_kendraCmd.Flags().BoolVarP(&_kendraDisassociateEntitiesFromExperience, "disassociate-entities-from-experience", "", false, "Disassociate Entities From Experience")
	_kendraCmd.Flags().BoolVarP(&_kendraDisassociatePersonasFromEntities, "disassociate-personas-from-entities", "", false, "Disassociate Personas From Entities")
	_kendraCmd.Flags().BoolVarP(&_kendraGetQuerySuggestions, "get-query-suggestions", "", false, "Get Query Suggestions")
	_kendraCmd.Flags().BoolVarP(&_kendraGetSnapshots, "get-snapshots", "", false, "Get Snapshots")
	_kendraCmd.Flags().BoolVarP(&_kendraListAccessControlConfigurations, "list-access-control-configurations", "", false, "List Access Control Configurations")
	_kendraCmd.Flags().BoolVarP(&_kendraListDataSourceSyncJobs, "list-data-source-sync-jobs", "", false, "List Data Source Sync Jobs")
	_kendraCmd.Flags().BoolVarP(&_kendraListDataSources, "list-data-sources", "", false, "List Data Sources")
	_kendraCmd.Flags().BoolVarP(&_kendraListEntityPersonas, "list-entity-personas", "", false, "List Entity Personas")
	_kendraCmd.Flags().BoolVarP(&_kendraListExperienceEntities, "list-experience-entities", "", false, "List Experience Entities")
	_kendraCmd.Flags().BoolVarP(&_kendraListExperiences, "list-experiences", "", false, "List Experiences")
	_kendraCmd.Flags().BoolVarP(&_kendraListFaqs, "list-faqs", "", false, "List Faqs")
	_kendraCmd.Flags().BoolVarP(&_kendraListFeaturedResultsSets, "list-featured-results-sets", "", false, "List Featured Results Sets")
	_kendraCmd.Flags().BoolVarP(&_kendraListGroupsOlderThanOrderingId, "list-groups-older-than-ordering-id", "", false, "List Groups Older Than Ordering ID")
	_kendraCmd.Flags().BoolVarP(&_kendraListIndices, "list-indices", "", false, "List Indices")
	_kendraCmd.Flags().BoolVarP(&_kendraListQuerySuggestionsBlockLists, "list-query-suggestions-block-lists", "", false, "List Query Suggestions Block Lists")
	_kendraCmd.Flags().BoolVarP(&_kendraListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_kendraCmd.Flags().BoolVarP(&_kendraListThesauri, "list-thesauri", "", false, "List Thesauri")
	_kendraCmd.Flags().BoolVarP(&_kendraPutPrincipalMapping, "put-principal-mapping", "", false, "Put Principal Mapping")
	_kendraCmd.Flags().BoolVarP(&_kendraQuery, "query", "", false, "Query")
	_kendraCmd.Flags().BoolVarP(&_kendraRetrieve, "retrieve", "", false, "Retrieve")
	_kendraCmd.Flags().BoolVarP(&_kendraStartDataSourceSyncJob, "start-data-source-sync-job", "", false, "Start Data Source Sync Job")
	_kendraCmd.Flags().BoolVarP(&_kendraStopDataSourceSyncJob, "stop-data-source-sync-job", "", false, "Stop Data Source Sync Job")
	_kendraCmd.Flags().BoolVarP(&_kendraSubmitFeedback, "submit-feedback", "", false, "Submit Feedback")
	_kendraCmd.Flags().BoolVarP(&_kendraTagResource, "tag-resource", "", false, "Tag Resource")
	_kendraCmd.Flags().BoolVarP(&_kendraUntagResource, "untag-resource", "", false, "Untag Resource")
	_kendraCmd.Flags().BoolVarP(&_kendraUpdateAccessControlConfiguration, "update-access-control-configuration", "", false, "Update Access Control Configuration")
	_kendraCmd.Flags().BoolVarP(&_kendraUpdateDataSource, "update-data-source", "", false, "Update Data Source")
	_kendraCmd.Flags().BoolVarP(&_kendraUpdateExperience, "update-experience", "", false, "Update Experience")
	_kendraCmd.Flags().BoolVarP(&_kendraUpdateFeaturedResultsSet, "update-featured-results-set", "", false, "Update Featured Results Set")
	_kendraCmd.Flags().BoolVarP(&_kendraUpdateIndex, "update-index", "", false, "Update Index")
	_kendraCmd.Flags().BoolVarP(&_kendraUpdateQuerySuggestionsBlockList, "update-query-suggestions-block-list", "", false, "Update Query Suggestions Block List")
	_kendraCmd.Flags().BoolVarP(&_kendraUpdateQuerySuggestionsConfig, "update-query-suggestions-config", "", false, "Update Query Suggestions Config")
	_kendraCmd.Flags().BoolVarP(&_kendraUpdateThesaurus, "update-thesaurus", "", false, "Update Thesaurus")

}
