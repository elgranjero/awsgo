package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/personalize"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// personalizeCmd represents the personalize command
var _personalizeCmd = &cobra.Command{
	Use:   "personalize",
	Short: "AWS personalize CLI",
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
		client := personalize.NewFromConfig(cfg)
		if _personalizeCreateBatchInferenceJob {
			personalize_CreateBatchInferenceJob(cfg, client)
			return
		}
		if _personalizeCreateBatchSegmentJob {
			personalize_CreateBatchSegmentJob(cfg, client)
			return
		}
		if _personalizeCreateCampaign {
			personalize_CreateCampaign(cfg, client)
			return
		}
		if _personalizeCreateDataDeletionJob {
			personalize_CreateDataDeletionJob(cfg, client)
			return
		}
		if _personalizeCreateDataset {
			personalize_CreateDataset(cfg, client)
			return
		}
		if _personalizeCreateDatasetExportJob {
			personalize_CreateDatasetExportJob(cfg, client)
			return
		}
		if _personalizeCreateDatasetGroup {
			personalize_CreateDatasetGroup(cfg, client)
			return
		}
		if _personalizeCreateDatasetImportJob {
			personalize_CreateDatasetImportJob(cfg, client)
			return
		}
		if _personalizeCreateEventTracker {
			personalize_CreateEventTracker(cfg, client)
			return
		}
		if _personalizeCreateFilter {
			personalize_CreateFilter(cfg, client)
			return
		}
		if _personalizeCreateMetricAttribution {
			personalize_CreateMetricAttribution(cfg, client)
			return
		}
		if _personalizeCreateRecommender {
			personalize_CreateRecommender(cfg, client)
			return
		}
		if _personalizeCreateSchema {
			personalize_CreateSchema(cfg, client)
			return
		}
		if _personalizeCreateSolution {
			personalize_CreateSolution(cfg, client)
			return
		}
		if _personalizeCreateSolutionVersion {
			personalize_CreateSolutionVersion(cfg, client)
			return
		}
		if _personalizeDeleteCampaign {
			personalize_DeleteCampaign(cfg, client)
			return
		}
		if _personalizeDeleteDataset {
			personalize_DeleteDataset(cfg, client)
			return
		}
		if _personalizeDeleteDatasetGroup {
			personalize_DeleteDatasetGroup(cfg, client)
			return
		}
		if _personalizeDeleteEventTracker {
			personalize_DeleteEventTracker(cfg, client)
			return
		}
		if _personalizeDeleteFilter {
			personalize_DeleteFilter(cfg, client)
			return
		}
		if _personalizeDeleteMetricAttribution {
			personalize_DeleteMetricAttribution(cfg, client)
			return
		}
		if _personalizeDeleteRecommender {
			personalize_DeleteRecommender(cfg, client)
			return
		}
		if _personalizeDeleteSchema {
			personalize_DeleteSchema(cfg, client)
			return
		}
		if _personalizeDeleteSolution {
			personalize_DeleteSolution(cfg, client)
			return
		}
		if _personalizeDescribeAlgorithm {
			personalize_DescribeAlgorithm(cfg, client)
			return
		}
		if _personalizeDescribeBatchInferenceJob {
			personalize_DescribeBatchInferenceJob(cfg, client)
			return
		}
		if _personalizeDescribeBatchSegmentJob {
			personalize_DescribeBatchSegmentJob(cfg, client)
			return
		}
		if _personalizeDescribeCampaign {
			personalize_DescribeCampaign(cfg, client)
			return
		}
		if _personalizeDescribeDataDeletionJob {
			personalize_DescribeDataDeletionJob(cfg, client)
			return
		}
		if _personalizeDescribeDataset {
			personalize_DescribeDataset(cfg, client)
			return
		}
		if _personalizeDescribeDatasetExportJob {
			personalize_DescribeDatasetExportJob(cfg, client)
			return
		}
		if _personalizeDescribeDatasetGroup {
			personalize_DescribeDatasetGroup(cfg, client)
			return
		}
		if _personalizeDescribeDatasetImportJob {
			personalize_DescribeDatasetImportJob(cfg, client)
			return
		}
		if _personalizeDescribeEventTracker {
			personalize_DescribeEventTracker(cfg, client)
			return
		}
		if _personalizeDescribeFeatureTransformation {
			personalize_DescribeFeatureTransformation(cfg, client)
			return
		}
		if _personalizeDescribeFilter {
			personalize_DescribeFilter(cfg, client)
			return
		}
		if _personalizeDescribeMetricAttribution {
			personalize_DescribeMetricAttribution(cfg, client)
			return
		}
		if _personalizeDescribeRecipe {
			personalize_DescribeRecipe(cfg, client)
			return
		}
		if _personalizeDescribeRecommender {
			personalize_DescribeRecommender(cfg, client)
			return
		}
		if _personalizeDescribeSchema {
			personalize_DescribeSchema(cfg, client)
			return
		}
		if _personalizeDescribeSolution {
			personalize_DescribeSolution(cfg, client)
			return
		}
		if _personalizeDescribeSolutionVersion {
			personalize_DescribeSolutionVersion(cfg, client)
			return
		}
		if _personalizeGetSolutionMetrics {
			personalize_GetSolutionMetrics(cfg, client)
			return
		}
		if _personalizeListBatchInferenceJobs {
			personalize_ListBatchInferenceJobs(cfg, client)
			return
		}
		if _personalizeListBatchSegmentJobs {
			personalize_ListBatchSegmentJobs(cfg, client)
			return
		}
		if _personalizeListCampaigns {
			personalize_ListCampaigns(cfg, client)
			return
		}
		if _personalizeListDataDeletionJobs {
			personalize_ListDataDeletionJobs(cfg, client)
			return
		}
		if _personalizeListDatasetExportJobs {
			personalize_ListDatasetExportJobs(cfg, client)
			return
		}
		if _personalizeListDatasetGroups {
			personalize_ListDatasetGroups(cfg, client)
			return
		}
		if _personalizeListDatasetImportJobs {
			personalize_ListDatasetImportJobs(cfg, client)
			return
		}
		if _personalizeListDatasets {
			personalize_ListDatasets(cfg, client)
			return
		}
		if _personalizeListEventTrackers {
			personalize_ListEventTrackers(cfg, client)
			return
		}
		if _personalizeListFilters {
			personalize_ListFilters(cfg, client)
			return
		}
		if _personalizeListMetricAttributionMetrics {
			personalize_ListMetricAttributionMetrics(cfg, client)
			return
		}
		if _personalizeListMetricAttributions {
			personalize_ListMetricAttributions(cfg, client)
			return
		}
		if _personalizeListRecipes {
			personalize_ListRecipes(cfg, client)
			return
		}
		if _personalizeListRecommenders {
			personalize_ListRecommenders(cfg, client)
			return
		}
		if _personalizeListSchemas {
			personalize_ListSchemas(cfg, client)
			return
		}
		if _personalizeListSolutionVersions {
			personalize_ListSolutionVersions(cfg, client)
			return
		}
		if _personalizeListSolutions {
			personalize_ListSolutions(cfg, client)
			return
		}
		if _personalizeListTagsForResource {
			personalize_ListTagsForResource(cfg, client)
			return
		}
		if _personalizeStartRecommender {
			personalize_StartRecommender(cfg, client)
			return
		}
		if _personalizeStopRecommender {
			personalize_StopRecommender(cfg, client)
			return
		}
		if _personalizeStopSolutionVersionCreation {
			personalize_StopSolutionVersionCreation(cfg, client)
			return
		}
		if _personalizeTagResource {
			personalize_TagResource(cfg, client)
			return
		}
		if _personalizeUntagResource {
			personalize_UntagResource(cfg, client)
			return
		}
		if _personalizeUpdateCampaign {
			personalize_UpdateCampaign(cfg, client)
			return
		}
		if _personalizeUpdateDataset {
			personalize_UpdateDataset(cfg, client)
			return
		}
		if _personalizeUpdateMetricAttribution {
			personalize_UpdateMetricAttribution(cfg, client)
			return
		}
		if _personalizeUpdateRecommender {
			personalize_UpdateRecommender(cfg, client)
			return
		}
		if _personalizeUpdateSolution {
			personalize_UpdateSolution(cfg, client)
			return
		}

	},
}

var (
	_personalizeCreateBatchInferenceJob       bool
	_personalizeCreateBatchSegmentJob         bool
	_personalizeCreateCampaign                bool
	_personalizeCreateDataDeletionJob         bool
	_personalizeCreateDataset                 bool
	_personalizeCreateDatasetExportJob        bool
	_personalizeCreateDatasetGroup            bool
	_personalizeCreateDatasetImportJob        bool
	_personalizeCreateEventTracker            bool
	_personalizeCreateFilter                  bool
	_personalizeCreateMetricAttribution       bool
	_personalizeCreateRecommender             bool
	_personalizeCreateSchema                  bool
	_personalizeCreateSolution                bool
	_personalizeCreateSolutionVersion         bool
	_personalizeDeleteCampaign                bool
	_personalizeDeleteDataset                 bool
	_personalizeDeleteDatasetGroup            bool
	_personalizeDeleteEventTracker            bool
	_personalizeDeleteFilter                  bool
	_personalizeDeleteMetricAttribution       bool
	_personalizeDeleteRecommender             bool
	_personalizeDeleteSchema                  bool
	_personalizeDeleteSolution                bool
	_personalizeDescribeAlgorithm             bool
	_personalizeDescribeBatchInferenceJob     bool
	_personalizeDescribeBatchSegmentJob       bool
	_personalizeDescribeCampaign              bool
	_personalizeDescribeDataDeletionJob       bool
	_personalizeDescribeDataset               bool
	_personalizeDescribeDatasetExportJob      bool
	_personalizeDescribeDatasetGroup          bool
	_personalizeDescribeDatasetImportJob      bool
	_personalizeDescribeEventTracker          bool
	_personalizeDescribeFeatureTransformation bool
	_personalizeDescribeFilter                bool
	_personalizeDescribeMetricAttribution     bool
	_personalizeDescribeRecipe                bool
	_personalizeDescribeRecommender           bool
	_personalizeDescribeSchema                bool
	_personalizeDescribeSolution              bool
	_personalizeDescribeSolutionVersion       bool
	_personalizeGetSolutionMetrics            bool
	_personalizeListBatchInferenceJobs        bool
	_personalizeListBatchSegmentJobs          bool
	_personalizeListCampaigns                 bool
	_personalizeListDataDeletionJobs          bool
	_personalizeListDatasetExportJobs         bool
	_personalizeListDatasetGroups             bool
	_personalizeListDatasetImportJobs         bool
	_personalizeListDatasets                  bool
	_personalizeListEventTrackers             bool
	_personalizeListFilters                   bool
	_personalizeListMetricAttributionMetrics  bool
	_personalizeListMetricAttributions        bool
	_personalizeListRecipes                   bool
	_personalizeListRecommenders              bool
	_personalizeListSchemas                   bool
	_personalizeListSolutionVersions          bool
	_personalizeListSolutions                 bool
	_personalizeListTagsForResource           bool
	_personalizeStartRecommender              bool
	_personalizeStopRecommender               bool
	_personalizeStopSolutionVersionCreation   bool
	_personalizeTagResource                   bool
	_personalizeUntagResource                 bool
	_personalizeUpdateCampaign                bool
	_personalizeUpdateDataset                 bool
	_personalizeUpdateMetricAttribution       bool
	_personalizeUpdateRecommender             bool
	_personalizeUpdateSolution                bool

	_personalizeAddMetrics                    string
	_personalizeAlgorithmArn                  string
	_personalizeBatchInferenceJobArn          string
	_personalizeBatchInferenceJobConfig       string
	_personalizeBatchInferenceJobMode         string
	_personalizeBatchSegmentJobArn            string
	_personalizeCampaignArn                   string
	_personalizeCampaignConfig                string
	_personalizeDataDeletionJobArn            string
	_personalizeDataSource                    string
	_personalizeDatasetArn                    string
	_personalizeDatasetExportJobArn           string
	_personalizeDatasetGroupArn               string
	_personalizeDatasetImportJobArn           string
	_personalizeDatasetType                   string
	_personalizeDomain                        string
	_personalizeEventTrackerArn               string
	_personalizeEventType                     string
	_personalizeFeatureTransformationArn      string
	_personalizeFilterArn                     string
	_personalizeFilterExpression              string
	_personalizeImportMode                    string
	_personalizeIngestionMode                 string
	_personalizeJobInput                      string
	_personalizeJobName                       string
	_personalizeJobOutput                     string
	_personalizeKmsKeyArn                     string
	_personalizeMaxResults                    string
	_personalizeMetricAttributionArn          string
	_personalizeMetrics                       string
	_personalizeMetricsOutputConfig           string
	_personalizeMinProvisionedTPS             string
	_personalizeName                          string
	_personalizeNextToken                     string
	_personalizeNumResults                    string
	_personalizePerformAutoML                 string
	_personalizePerformAutoTraining           string
	_personalizePerformHPO                    string
	_personalizePerformIncrementalUpdate      string
	_personalizePublishAttributionMetricsToS3 string
	_personalizeRecipeArn                     string
	_personalizeRecipeProvider                string
	_personalizeRecommenderArn                string
	_personalizeRecommenderConfig             string
	_personalizeRemoveMetrics                 []string
	_personalizeResourceArn                   string
	_personalizeRoleArn                       string
	_personalizeSchema                        string
	_personalizeSchemaArn                     string
	_personalizeSolutionArn                   string
	_personalizeSolutionConfig                string
	_personalizeSolutionUpdateConfig          string
	_personalizeSolutionVersionArn            string
	_personalizeTagKeys                       []string
	_personalizeTags                          string
	_personalizeThemeGenerationConfig         string
	_personalizeTrainingMode                  string
)

// Generates batch recommendations based on a list of items or users stored in
// Amazon S3 and exports the recommendations to an Amazon S3 bucket.
//
// To generate batch recommendations, specify the ARN of a solution version and an
// Amazon S3 URI for the input and output data. For user personalization, popular
// items, and personalized ranking solutions, the batch inference job generates a
// list of recommended items for each user ID in the input file. For related items
// solutions, the job generates a list of recommended items for each item ID in the
// input file.
//
// For more information, see [Creating a batch inference job].
//
// If you use the Similar-Items recipe, Amazon Personalize can add descriptive
// themes to batch recommendations. To generate themes, set the job's mode to
// THEME_GENERATION and specify the name of the field that contains item names in
// the input data.
//
// For more information about generating themes, see [Batch recommendations with themes from Content Generator].
//
// You can't get batch recommendations with the Trending-Now or Next-Best-Action
// recipes.
//
// [Creating a batch inference job]: https://docs.aws.amazon.com/personalize/latest/dg/getting-batch-recommendations.html
// [Batch recommendations with themes from Content Generator]: https://docs.aws.amazon.com/personalize/latest/dg/themed-batch-recommendations.html
func personalize_CreateBatchInferenceJob(cfg aws.Config, client *personalize.Client) {
	input := &personalize.CreateBatchInferenceJobInput{
		// JobInput: *types.BatchInferenceJobInput, // Required
		// JobName: *string, // Required
		// JobOutput: *types.BatchInferenceJobOutput, // Required
		// RoleArn: *string, // Required
		// SolutionVersionArn: *string, // Required
	}

	if len(_personalizeJobInput) > 0 {
		if err := assignInputField(input, "JobInput", _personalizeJobInput); err != nil {
			log.Errorf("invalid --job-input: %s", err.Error())
			return
		}
	}
	if len(_personalizeJobName) > 0 {
		input.JobName = aws.String(_personalizeJobName)
	}
	if len(_personalizeJobOutput) > 0 {
		if err := assignInputField(input, "JobOutput", _personalizeJobOutput); err != nil {
			log.Errorf("invalid --job-output: %s", err.Error())
			return
		}
	}
	if len(_personalizeRoleArn) > 0 {
		input.RoleArn = aws.String(_personalizeRoleArn)
	}
	if len(_personalizeSolutionVersionArn) > 0 {
		input.SolutionVersionArn = aws.String(_personalizeSolutionVersionArn)
	}
	if len(_personalizeBatchInferenceJobConfig) > 0 {
		if err := assignInputField(input, "BatchInferenceJobConfig", _personalizeBatchInferenceJobConfig); err != nil {
			log.Errorf("invalid --batch-inference-job-config: %s", err.Error())
			return
		}
	}
	if len(_personalizeBatchInferenceJobMode) > 0 {
		if err := assignInputField(input, "BatchInferenceJobMode", _personalizeBatchInferenceJobMode); err != nil {
			log.Errorf("invalid --batch-inference-job-mode: %s", err.Error())
			return
		}
	}
	if len(_personalizeFilterArn) > 0 {
		input.FilterArn = aws.String(_personalizeFilterArn)
	}
	if len(_personalizeNumResults) > 0 {
		if err := assignInputField(input, "NumResults", _personalizeNumResults); err != nil {
			log.Errorf("invalid --num-results: %s", err.Error())
			return
		}
	}
	if len(_personalizeTags) > 0 {
		if err := assignInputField(input, "Tags", _personalizeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_personalizeThemeGenerationConfig) > 0 {
		if err := assignInputField(input, "ThemeGenerationConfig", _personalizeThemeGenerationConfig); err != nil {
			log.Errorf("invalid --theme-generation-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBatchInferenceJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a batch segment job. The operation can handle up to 50 million records
// and the input file must be in JSON format. For more information, see [Getting batch recommendations and user segments].
//
// [Getting batch recommendations and user segments]: https://docs.aws.amazon.com/personalize/latest/dg/recommendations-batch.html
func personalize_CreateBatchSegmentJob(cfg aws.Config, client *personalize.Client) {
	input := &personalize.CreateBatchSegmentJobInput{
		// JobInput: *types.BatchSegmentJobInput, // Required
		// JobName: *string, // Required
		// JobOutput: *types.BatchSegmentJobOutput, // Required
		// RoleArn: *string, // Required
		// SolutionVersionArn: *string, // Required
	}

	if len(_personalizeJobInput) > 0 {
		if err := assignInputField(input, "JobInput", _personalizeJobInput); err != nil {
			log.Errorf("invalid --job-input: %s", err.Error())
			return
		}
	}
	if len(_personalizeJobName) > 0 {
		input.JobName = aws.String(_personalizeJobName)
	}
	if len(_personalizeJobOutput) > 0 {
		if err := assignInputField(input, "JobOutput", _personalizeJobOutput); err != nil {
			log.Errorf("invalid --job-output: %s", err.Error())
			return
		}
	}
	if len(_personalizeRoleArn) > 0 {
		input.RoleArn = aws.String(_personalizeRoleArn)
	}
	if len(_personalizeSolutionVersionArn) > 0 {
		input.SolutionVersionArn = aws.String(_personalizeSolutionVersionArn)
	}
	if len(_personalizeFilterArn) > 0 {
		input.FilterArn = aws.String(_personalizeFilterArn)
	}
	if len(_personalizeNumResults) > 0 {
		if err := assignInputField(input, "NumResults", _personalizeNumResults); err != nil {
			log.Errorf("invalid --num-results: %s", err.Error())
			return
		}
	}
	if len(_personalizeTags) > 0 {
		if err := assignInputField(input, "Tags", _personalizeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBatchSegmentJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// You incur campaign costs while it is active. To avoid unnecessary costs, make
// sure to delete the campaign when you are finished. For information about
// campaign costs, see [Amazon Personalize pricing].
//
// Creates a campaign that deploys a solution version. When a client calls the [GetRecommendations]
// and [GetPersonalizedRanking]APIs, a campaign is specified in the request.
//
// # Minimum Provisioned TPS and Auto-Scaling
//
// A high minProvisionedTPS will increase your cost. We recommend starting with 1
// for minProvisionedTPS (the default). Track your usage using Amazon CloudWatch
// metrics, and increase the minProvisionedTPS as necessary.
//
// When you create an Amazon Personalize campaign, you can specify the minimum
// provisioned transactions per second ( minProvisionedTPS ) for the campaign. This
// is the baseline transaction throughput for the campaign provisioned by Amazon
// Personalize. It sets the minimum billing charge for the campaign while it is
// active. A transaction is a single GetRecommendations or GetPersonalizedRanking
// request. The default minProvisionedTPS is 1.
//
// If your TPS increases beyond the minProvisionedTPS , Amazon Personalize
// auto-scales the provisioned capacity up and down, but never below
// minProvisionedTPS . There's a short time delay while the capacity is increased
// that might cause loss of transactions. When your traffic reduces, capacity
// returns to the minProvisionedTPS .
//
// You are charged for the the minimum provisioned TPS or, if your requests exceed
// the minProvisionedTPS , the actual TPS. The actual TPS is the total number of
// recommendation requests you make. We recommend starting with a low
// minProvisionedTPS , track your usage using Amazon CloudWatch metrics, and then
// increase the minProvisionedTPS as necessary.
//
// For more information about campaign costs, see [Amazon Personalize pricing].
//
// # Status
//
// A campaign can be in one of the following states:
//
// - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// - DELETE PENDING > DELETE IN_PROGRESS
//
// To get the campaign status, call [DescribeCampaign].
//
// Wait until the status of the campaign is ACTIVE before asking the campaign for
// recommendations.
//
// # Related APIs
//
// [ListCampaigns]
//
// [DescribeCampaign]
//
// [UpdateCampaign]
//
// [DeleteCampaign]
//
// [UpdateCampaign]: https://docs.aws.amazon.com/personalize/latest/dg/API_UpdateCampaign.html
// [GetRecommendations]: https://docs.aws.amazon.com/personalize/latest/dg/API_RS_GetRecommendations.html
// [ListCampaigns]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListCampaigns.html
// [DeleteCampaign]: https://docs.aws.amazon.com/personalize/latest/dg/API_DeleteCampaign.html
// [GetPersonalizedRanking]: https://docs.aws.amazon.com/personalize/latest/dg/API_RS_GetPersonalizedRanking.html
// [Amazon Personalize pricing]: https://aws.amazon.com/personalize/pricing/
// [DescribeCampaign]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeCampaign.html
func personalize_CreateCampaign(cfg aws.Config, client *personalize.Client) {
	input := &personalize.CreateCampaignInput{
		// Name: *string, // Required
		// SolutionVersionArn: *string, // Required
	}

	if len(_personalizeName) > 0 {
		input.Name = aws.String(_personalizeName)
	}
	if len(_personalizeSolutionVersionArn) > 0 {
		input.SolutionVersionArn = aws.String(_personalizeSolutionVersionArn)
	}
	if len(_personalizeCampaignConfig) > 0 {
		if err := assignInputField(input, "CampaignConfig", _personalizeCampaignConfig); err != nil {
			log.Errorf("invalid --campaign-config: %s", err.Error())
			return
		}
	}
	if len(_personalizeMinProvisionedTPS) > 0 {
		if err := assignInputField(input, "MinProvisionedTPS", _personalizeMinProvisionedTPS); err != nil {
			log.Errorf("invalid --min-provisioned-tps: %s", err.Error())
			return
		}
	}
	if len(_personalizeTags) > 0 {
		if err := assignInputField(input, "Tags", _personalizeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a batch job that deletes all references to specific users from an
// Amazon Personalize dataset group in batches. You specify the users to delete in
// a CSV file of userIds in an Amazon S3 bucket. After a job completes, Amazon
// Personalize no longer trains on the users’ data and no longer considers the
// users when generating user segments. For more information about creating a data
// deletion job, see [Deleting users].
//
// - Your input file must be a CSV file with a single USER_ID column that lists
// the users IDs. For more information about preparing the CSV file, see [Preparing your data deletion file and uploading it to Amazon S3].
//
// - To give Amazon Personalize permission to access your input CSV file of
// userIds, you must specify an IAM service role that has permission to read from
// the data source. This role needs GetObject and ListBucket permissions for the
// bucket and its content. These permissions are the same as importing data. For
// information on granting access to your Amazon S3 bucket, see [Giving Amazon Personalize Access to Amazon S3 Resources].
//
// After you create a job, it can take up to a day to delete all references to the
// users from datasets and models. Until the job completes, Amazon Personalize
// continues to use the data when training. And if you use a User Segmentation
// recipe, the users might appear in user segments.
//
// # Status
//
// A data deletion job can have one of the following statuses:
//
// - PENDING > IN_PROGRESS > COMPLETED -or- FAILED
//
// To get the status of the data deletion job, call [DescribeDataDeletionJob] API operation and specify the
// Amazon Resource Name (ARN) of the job. If the status is FAILED, the response
// includes a failureReason key, which describes why the job failed.
//
// # Related APIs
//
// [ListDataDeletionJobs]
//
// [DescribeDataDeletionJob]
//
// [ListDataDeletionJobs]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListDataDeletionJobs.html
// [Giving Amazon Personalize Access to Amazon S3 Resources]: https://docs.aws.amazon.com/personalize/latest/dg/granting-personalize-s3-access.html
// [Deleting users]: https://docs.aws.amazon.com/personalize/latest/dg/delete-records.html
// [Preparing your data deletion file and uploading it to Amazon S3]: https://docs.aws.amazon.com/personalize/latest/dg/prepare-deletion-input-file.html
// [DescribeDataDeletionJob]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeDataDeletionJob.html
func personalize_CreateDataDeletionJob(cfg aws.Config, client *personalize.Client) {
	input := &personalize.CreateDataDeletionJobInput{
		// DataSource: *types.DataSource, // Required
		// DatasetGroupArn: *string, // Required
		// JobName: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_personalizeDataSource) > 0 {
		if err := assignInputField(input, "DataSource", _personalizeDataSource); err != nil {
			log.Errorf("invalid --data-source: %s", err.Error())
			return
		}
	}
	if len(_personalizeDatasetGroupArn) > 0 {
		input.DatasetGroupArn = aws.String(_personalizeDatasetGroupArn)
	}
	if len(_personalizeJobName) > 0 {
		input.JobName = aws.String(_personalizeJobName)
	}
	if len(_personalizeRoleArn) > 0 {
		input.RoleArn = aws.String(_personalizeRoleArn)
	}
	if len(_personalizeTags) > 0 {
		if err := assignInputField(input, "Tags", _personalizeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataDeletionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an empty dataset and adds it to the specified dataset group. Use [CreateDatasetImportJob] to
// import your training data to a dataset.
//
// There are 5 types of datasets:
//
// - Item interactions
//
// - Items
//
// - Users
//
// - Action interactions
//
// - Actions
//
// Each dataset type has an associated schema with required field types. Only the
// Item interactions dataset is required in order to train a model (also referred
// to as creating a solution).
//
// A dataset can be in one of the following states:
//
// - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// - DELETE PENDING > DELETE IN_PROGRESS
//
// To get the status of the dataset, call [DescribeDataset].
//
// # Related APIs
//
// [CreateDatasetGroup]
//
// [ListDatasets]
//
// [DescribeDataset]
//
// [DeleteDataset]
//
// [DescribeDataset]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeDataset.html
// [ListDatasets]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListDatasets.html
// [DeleteDataset]: https://docs.aws.amazon.com/personalize/latest/dg/API_DeleteDataset.html
// [CreateDatasetGroup]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDatasetGroup.html
// [CreateDatasetImportJob]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDatasetImportJob.html
func personalize_CreateDataset(cfg aws.Config, client *personalize.Client) {
	input := &personalize.CreateDatasetInput{
		// DatasetGroupArn: *string, // Required
		// DatasetType: *string, // Required
		// Name: *string, // Required
		// SchemaArn: *string, // Required
	}

	if len(_personalizeDatasetGroupArn) > 0 {
		input.DatasetGroupArn = aws.String(_personalizeDatasetGroupArn)
	}
	if len(_personalizeDatasetType) > 0 {
		input.DatasetType = aws.String(_personalizeDatasetType)
	}
	if len(_personalizeName) > 0 {
		input.Name = aws.String(_personalizeName)
	}
	if len(_personalizeSchemaArn) > 0 {
		input.SchemaArn = aws.String(_personalizeSchemaArn)
	}
	if len(_personalizeTags) > 0 {
		if err := assignInputField(input, "Tags", _personalizeTags); err != nil {
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

// Creates a job that exports data from your dataset to an Amazon S3 bucket. To
// allow Amazon Personalize to export the training data, you must specify an
// service-linked IAM role that gives Amazon Personalize PutObject permissions for
// your Amazon S3 bucket. For information, see [Exporting a dataset]in the Amazon Personalize developer
// guide.
//
// # Status
//
// A dataset export job can be in one of the following states:
//
// - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// To get the status of the export job, call [DescribeDatasetExportJob], and specify the Amazon Resource
// Name (ARN) of the dataset export job. The dataset export is complete when the
// status shows as ACTIVE. If the status shows as CREATE FAILED, the response
// includes a failureReason key, which describes why the job failed.
//
// [Exporting a dataset]: https://docs.aws.amazon.com/personalize/latest/dg/export-data.html
// [DescribeDatasetExportJob]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeDatasetExportJob.html
func personalize_CreateDatasetExportJob(cfg aws.Config, client *personalize.Client) {
	input := &personalize.CreateDatasetExportJobInput{
		// DatasetArn: *string, // Required
		// JobName: *string, // Required
		// JobOutput: *types.DatasetExportJobOutput, // Required
		// RoleArn: *string, // Required
	}

	if len(_personalizeDatasetArn) > 0 {
		input.DatasetArn = aws.String(_personalizeDatasetArn)
	}
	if len(_personalizeJobName) > 0 {
		input.JobName = aws.String(_personalizeJobName)
	}
	if len(_personalizeJobOutput) > 0 {
		if err := assignInputField(input, "JobOutput", _personalizeJobOutput); err != nil {
			log.Errorf("invalid --job-output: %s", err.Error())
			return
		}
	}
	if len(_personalizeRoleArn) > 0 {
		input.RoleArn = aws.String(_personalizeRoleArn)
	}
	if len(_personalizeIngestionMode) > 0 {
		if err := assignInputField(input, "IngestionMode", _personalizeIngestionMode); err != nil {
			log.Errorf("invalid --ingestion-mode: %s", err.Error())
			return
		}
	}
	if len(_personalizeTags) > 0 {
		if err := assignInputField(input, "Tags", _personalizeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDatasetExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an empty dataset group. A dataset group is a container for Amazon
// Personalize resources. A dataset group can contain at most three datasets, one
// for each type of dataset:
//
// - Item interactions
//
// - Items
//
// - Users
//
// - Actions
//
// - Action interactions
//
// A dataset group can be a Domain dataset group, where you specify a domain and
// use pre-configured resources like recommenders, or a Custom dataset group, where
// you use custom resources, such as a solution with a solution version, that you
// deploy with a campaign. If you start with a Domain dataset group, you can still
// add custom resources such as solutions and solution versions trained with
// recipes for custom use cases and deployed with campaigns.
//
// A dataset group can be in one of the following states:
//
// - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// - DELETE PENDING
//
// To get the status of the dataset group, call [DescribeDatasetGroup]. If the status shows as CREATE
// FAILED, the response includes a failureReason key, which describes why the
// creation failed.
//
// You must wait until the status of the dataset group is ACTIVE before adding a
// dataset to the group.
//
// You can specify an Key Management Service (KMS) key to encrypt the datasets in
// the group. If you specify a KMS key, you must also include an Identity and
// Access Management (IAM) role that has permission to access the key.
//
// # APIs that require a dataset group ARN in the request
//
// [CreateDataset]
//
// [CreateEventTracker]
//
// [CreateSolution]
//
// # Related APIs
//
// [ListDatasetGroups]
//
// [DescribeDatasetGroup]
//
// [DeleteDatasetGroup]
//
// [CreateDataset]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDataset.html
// [ListDatasetGroups]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListDatasetGroups.html
// [CreateSolution]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateSolution.html
// [DescribeDatasetGroup]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeDatasetGroup.html
// [CreateEventTracker]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateEventTracker.html
// [DeleteDatasetGroup]: https://docs.aws.amazon.com/personalize/latest/dg/API_DeleteDatasetGroup.html
func personalize_CreateDatasetGroup(cfg aws.Config, client *personalize.Client) {
	input := &personalize.CreateDatasetGroupInput{
		// Name: *string, // Required
	}

	if len(_personalizeName) > 0 {
		input.Name = aws.String(_personalizeName)
	}
	if len(_personalizeDomain) > 0 {
		if err := assignInputField(input, "Domain", _personalizeDomain); err != nil {
			log.Errorf("invalid --domain: %s", err.Error())
			return
		}
	}
	if len(_personalizeKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_personalizeKmsKeyArn)
	}
	if len(_personalizeRoleArn) > 0 {
		input.RoleArn = aws.String(_personalizeRoleArn)
	}
	if len(_personalizeTags) > 0 {
		if err := assignInputField(input, "Tags", _personalizeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDatasetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a job that imports training data from your data source (an Amazon S3
// bucket) to an Amazon Personalize dataset. To allow Amazon Personalize to import
// the training data, you must specify an IAM service role that has permission to
// read from the data source, as Amazon Personalize makes a copy of your data and
// processes it internally. For information on granting access to your Amazon S3
// bucket, see [Giving Amazon Personalize Access to Amazon S3 Resources].
//
// If you already created a recommender or deployed a custom solution version with
// a campaign, how new bulk records influence recommendations depends on the domain
// use case or recipe that you use. For more information, see [How new data influences real-time recommendations].
//
// By default, a dataset import job replaces any existing data in the dataset that
// you imported in bulk. To add new records without replacing existing data,
// specify INCREMENTAL for the import mode in the CreateDatasetImportJob operation.
//
// # Status
//
// A dataset import job can be in one of the following states:
//
// - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// To get the status of the import job, call [DescribeDatasetImportJob], providing the Amazon Resource Name
// (ARN) of the dataset import job. The dataset import is complete when the status
// shows as ACTIVE. If the status shows as CREATE FAILED, the response includes a
// failureReason key, which describes why the job failed.
//
// Importing takes time. You must wait until the status shows as ACTIVE before
// training a model using the dataset.
//
// # Related APIs
//
// [ListDatasetImportJobs]
//
// [DescribeDatasetImportJob]
//
// [ListDatasetImportJobs]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListDatasetImportJobs.html
// [Giving Amazon Personalize Access to Amazon S3 Resources]: https://docs.aws.amazon.com/personalize/latest/dg/granting-personalize-s3-access.html
// [DescribeDatasetImportJob]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeDatasetImportJob.html
// [How new data influences real-time recommendations]: https://docs.aws.amazon.com/personalize/latest/dg/how-new-data-influences-recommendations.html
func personalize_CreateDatasetImportJob(cfg aws.Config, client *personalize.Client) {
	input := &personalize.CreateDatasetImportJobInput{
		// DataSource: *types.DataSource, // Required
		// DatasetArn: *string, // Required
		// JobName: *string, // Required
	}

	if len(_personalizeDataSource) > 0 {
		if err := assignInputField(input, "DataSource", _personalizeDataSource); err != nil {
			log.Errorf("invalid --data-source: %s", err.Error())
			return
		}
	}
	if len(_personalizeDatasetArn) > 0 {
		input.DatasetArn = aws.String(_personalizeDatasetArn)
	}
	if len(_personalizeJobName) > 0 {
		input.JobName = aws.String(_personalizeJobName)
	}
	if len(_personalizeImportMode) > 0 {
		if err := assignInputField(input, "ImportMode", _personalizeImportMode); err != nil {
			log.Errorf("invalid --import-mode: %s", err.Error())
			return
		}
	}
	if len(_personalizePublishAttributionMetricsToS3) > 0 {
		if err := assignInputField(input, "PublishAttributionMetricsToS3", _personalizePublishAttributionMetricsToS3); err != nil {
			log.Errorf("invalid --publish-attribution-metrics-to-s3: %s", err.Error())
			return
		}
	}
	if len(_personalizeRoleArn) > 0 {
		input.RoleArn = aws.String(_personalizeRoleArn)
	}
	if len(_personalizeTags) > 0 {
		if err := assignInputField(input, "Tags", _personalizeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDatasetImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an event tracker that you use when adding event data to a specified
// dataset group using the [PutEvents]API.
//
// Only one event tracker can be associated with a dataset group. You will get an
// error if you call CreateEventTracker using the same dataset group as an
// existing event tracker.
//
// When you create an event tracker, the response includes a tracking ID, which
// you pass as a parameter when you use the [PutEvents]operation. Amazon Personalize then
// appends the event data to the Item interactions dataset of the dataset group you
// specify in your event tracker.
//
// The event tracker can be in one of the following states:
//
// - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// - DELETE PENDING > DELETE IN_PROGRESS
//
// To get the status of the event tracker, call [DescribeEventTracker].
//
// The event tracker must be in the ACTIVE state before using the tracking ID.
//
// # Related APIs
//
// [ListEventTrackers]
//
// [DescribeEventTracker]
//
// [DeleteEventTracker]
//
// [PutEvents]: https://docs.aws.amazon.com/personalize/latest/dg/API_UBS_PutEvents.html
// [DescribeEventTracker]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeEventTracker.html
// [ListEventTrackers]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListEventTrackers.html
// [DeleteEventTracker]: https://docs.aws.amazon.com/personalize/latest/dg/API_DeleteEventTracker.html
func personalize_CreateEventTracker(cfg aws.Config, client *personalize.Client) {
	input := &personalize.CreateEventTrackerInput{
		// DatasetGroupArn: *string, // Required
		// Name: *string, // Required
	}

	if len(_personalizeDatasetGroupArn) > 0 {
		input.DatasetGroupArn = aws.String(_personalizeDatasetGroupArn)
	}
	if len(_personalizeName) > 0 {
		input.Name = aws.String(_personalizeName)
	}
	if len(_personalizeTags) > 0 {
		if err := assignInputField(input, "Tags", _personalizeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEventTracker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a recommendation filter. For more information, see [Filtering recommendations and user segments].
//
// [Filtering recommendations and user segments]: https://docs.aws.amazon.com/personalize/latest/dg/filter.html
func personalize_CreateFilter(cfg aws.Config, client *personalize.Client) {
	input := &personalize.CreateFilterInput{
		// DatasetGroupArn: *string, // Required
		// FilterExpression: *string, // Required
		// Name: *string, // Required
	}

	if len(_personalizeDatasetGroupArn) > 0 {
		input.DatasetGroupArn = aws.String(_personalizeDatasetGroupArn)
	}
	if len(_personalizeFilterExpression) > 0 {
		input.FilterExpression = aws.String(_personalizeFilterExpression)
	}
	if len(_personalizeName) > 0 {
		input.Name = aws.String(_personalizeName)
	}
	if len(_personalizeTags) > 0 {
		if err := assignInputField(input, "Tags", _personalizeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a metric attribution. A metric attribution creates reports on the data
// that you import into Amazon Personalize. Depending on how you imported the data,
// you can view reports in Amazon CloudWatch or Amazon S3. For more information,
// see [Measuring impact of recommendations].
//
// [Measuring impact of recommendations]: https://docs.aws.amazon.com/personalize/latest/dg/measuring-recommendation-impact.html
func personalize_CreateMetricAttribution(cfg aws.Config, client *personalize.Client) {
	input := &personalize.CreateMetricAttributionInput{
		// DatasetGroupArn: *string, // Required
		// Metrics: []types.MetricAttribute, // Required
		// MetricsOutputConfig: *types.MetricAttributionOutput, // Required
		// Name: *string, // Required
	}

	if len(_personalizeDatasetGroupArn) > 0 {
		input.DatasetGroupArn = aws.String(_personalizeDatasetGroupArn)
	}
	if len(_personalizeMetrics) > 0 {
		if err := assignInputField(input, "Metrics", _personalizeMetrics); err != nil {
			log.Errorf("invalid --metrics: %s", err.Error())
			return
		}
	}
	if len(_personalizeMetricsOutputConfig) > 0 {
		if err := assignInputField(input, "MetricsOutputConfig", _personalizeMetricsOutputConfig); err != nil {
			log.Errorf("invalid --metrics-output-config: %s", err.Error())
			return
		}
	}
	if len(_personalizeName) > 0 {
		input.Name = aws.String(_personalizeName)
	}

	if resp, err := client.CreateMetricAttribution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a recommender with the recipe (a Domain dataset group use case) you
// specify. You create recommenders for a Domain dataset group and specify the
// recommender's Amazon Resource Name (ARN) when you make a [GetRecommendations]request.
//
// # Minimum recommendation requests per second
//
// A high minRecommendationRequestsPerSecond will increase your bill. We recommend
// starting with 1 for minRecommendationRequestsPerSecond (the default). Track
// your usage using Amazon CloudWatch metrics, and increase the
// minRecommendationRequestsPerSecond as necessary.
//
// When you create a recommender, you can configure the recommender's minimum
// recommendation requests per second. The minimum recommendation requests per
// second ( minRecommendationRequestsPerSecond ) specifies the baseline
// recommendation request throughput provisioned by Amazon Personalize. The default
// minRecommendationRequestsPerSecond is 1 . A recommendation request is a single
// GetRecommendations operation. Request throughput is measured in requests per
// second and Amazon Personalize uses your requests per second to derive your
// requests per hour and the price of your recommender usage.
//
// If your requests per second increases beyond minRecommendationRequestsPerSecond
// , Amazon Personalize auto-scales the provisioned capacity up and down, but never
// below minRecommendationRequestsPerSecond . There's a short time delay while the
// capacity is increased that might cause loss of requests.
//
// Your bill is the greater of either the minimum requests per hour (based on
// minRecommendationRequestsPerSecond) or the actual number of requests. The actual
// request throughput used is calculated as the average requests/second within a
// one-hour window.
//
// We recommend starting with the default minRecommendationRequestsPerSecond ,
// track your usage using Amazon CloudWatch metrics, and then increase the
// minRecommendationRequestsPerSecond as necessary.
//
// # Status
//
// A recommender can be in one of the following states:
//
// - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// - STOP PENDING > STOP IN_PROGRESS > INACTIVE > START PENDING > START
// IN_PROGRESS > ACTIVE
//
// - DELETE PENDING > DELETE IN_PROGRESS
//
// To get the recommender status, call [DescribeRecommender].
//
// Wait until the status of the recommender is ACTIVE before asking the
// recommender for recommendations.
//
// # Related APIs
//
// [ListRecommenders]
//
// [DescribeRecommender]
//
// [UpdateRecommender]
//
// [DeleteRecommender]
//
// [ListRecommenders]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListRecommenders.html
// [GetRecommendations]: https://docs.aws.amazon.com/personalize/latest/dg/API_RS_GetRecommendations.html
// [UpdateRecommender]: https://docs.aws.amazon.com/personalize/latest/dg/API_UpdateRecommender.html
// [DeleteRecommender]: https://docs.aws.amazon.com/personalize/latest/dg/API_DeleteRecommender.html
// [DescribeRecommender]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeRecommender.html
func personalize_CreateRecommender(cfg aws.Config, client *personalize.Client) {
	input := &personalize.CreateRecommenderInput{
		// DatasetGroupArn: *string, // Required
		// Name: *string, // Required
		// RecipeArn: *string, // Required
	}

	if len(_personalizeDatasetGroupArn) > 0 {
		input.DatasetGroupArn = aws.String(_personalizeDatasetGroupArn)
	}
	if len(_personalizeName) > 0 {
		input.Name = aws.String(_personalizeName)
	}
	if len(_personalizeRecipeArn) > 0 {
		input.RecipeArn = aws.String(_personalizeRecipeArn)
	}
	if len(_personalizeRecommenderConfig) > 0 {
		if err := assignInputField(input, "RecommenderConfig", _personalizeRecommenderConfig); err != nil {
			log.Errorf("invalid --recommender-config: %s", err.Error())
			return
		}
	}
	if len(_personalizeTags) > 0 {
		if err := assignInputField(input, "Tags", _personalizeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRecommender(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Personalize schema from the specified schema string. The
// schema you create must be in Avro JSON format.
//
// Amazon Personalize recognizes three schema variants. Each schema is associated
// with a dataset type and has a set of required field and keywords. If you are
// creating a schema for a dataset in a Domain dataset group, you provide the
// domain of the Domain dataset group. You specify a schema when you call [CreateDataset].
//
// # Related APIs
//
// [ListSchemas]
//
// [DescribeSchema]
//
// [DeleteSchema]
//
// [CreateDataset]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDataset.html
// [DescribeSchema]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeSchema.html
// [DeleteSchema]: https://docs.aws.amazon.com/personalize/latest/dg/API_DeleteSchema.html
// [ListSchemas]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListSchemas.html
func personalize_CreateSchema(cfg aws.Config, client *personalize.Client) {
	input := &personalize.CreateSchemaInput{
		// Name: *string, // Required
		// Schema: *string, // Required
	}

	if len(_personalizeName) > 0 {
		input.Name = aws.String(_personalizeName)
	}
	if len(_personalizeSchema) > 0 {
		input.Schema = aws.String(_personalizeSchema)
	}
	if len(_personalizeDomain) > 0 {
		if err := assignInputField(input, "Domain", _personalizeDomain); err != nil {
			log.Errorf("invalid --domain: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// By default, all new solutions use automatic training. With automatic training,
// you incur training costs while your solution is active. To avoid unnecessary
// costs, when you are finished you can [update the solution]to turn off automatic training. For
// information about training costs, see [Amazon Personalize pricing].
//
// Creates the configuration for training a model (creating a solution version).
// This configuration includes the recipe to use for model training and optional
// training configuration, such as columns to use in training and feature
// transformation parameters. For more information about configuring a solution,
// see [Creating and configuring a solution].
//
// By default, new solutions use automatic training to create solution versions
// every 7 days. You can change the training frequency. Automatic solution version
// creation starts within one hour after the solution is ACTIVE. If you manually
// create a solution version within the hour, the solution skips the first
// automatic training. For more information, see [Configuring automatic training].
//
// To turn off automatic training, set performAutoTraining to false. If you turn
// off automatic training, you must manually create a solution version by calling
// the [CreateSolutionVersion]operation.
//
// After training starts, you can get the solution version's Amazon Resource Name
// (ARN) with the [ListSolutionVersions]API operation. To get its status, use the [DescribeSolutionVersion].
//
// After training completes you can evaluate model accuracy by calling [GetSolutionMetrics]. When you
// are satisfied with the solution version, you deploy it using [CreateCampaign]. The campaign
// provides recommendations to a client through the [GetRecommendations]API.
//
// Amazon Personalize doesn't support configuring the hpoObjective for solution
// hyperparameter optimization at this time.
//
// # Status
//
// A solution can be in one of the following states:
//
// - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// - DELETE PENDING > DELETE IN_PROGRESS
//
// To get the status of the solution, call [DescribeSolution]. If you use manual training, the
// status must be ACTIVE before you call CreateSolutionVersion .
//
// # Related APIs
//
// [UpdateSolution]
//
// [ListSolutions]
//
// [CreateSolutionVersion]
//
// [DescribeSolution]
//
// [DeleteSolution]
//
// [ListSolutionVersions]
//
// [DescribeSolutionVersion]
//
// [CreateCampaign]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateCampaign.html
// [GetSolutionMetrics]: https://docs.aws.amazon.com/personalize/latest/dg/API_GetSolutionMetrics.html
// [update the solution]: https://docs.aws.amazon.com/personalize/latest/dg/API_UpdateSolution.html
// [ListSolutions]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListSolutions.html
// [Amazon Personalize pricing]: https://aws.amazon.com/personalize/pricing/
// [DescribeSolution]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeSolution.html
// [DescribeSolutionVersion]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeSolutionVersion.html
// [DeleteSolution]: https://docs.aws.amazon.com/personalize/latest/dg/API_DeleteSolution.html
// [UpdateSolution]: https://docs.aws.amazon.com/personalize/latest/dg/API_UpdateSolution.html
// [ListSolutionVersions]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListSolutionVersions.html
// [Creating and configuring a solution]: https://docs.aws.amazon.com/personalize/latest/dg/customizing-solution-config.html
// [GetRecommendations]: https://docs.aws.amazon.com/personalize/latest/dg/API_RS_GetRecommendations.html
// [Configuring automatic training]: https://docs.aws.amazon.com/personalize/latest/dg/solution-config-auto-training.html
// [CreateSolutionVersion]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateSolutionVersion.html
func personalize_CreateSolution(cfg aws.Config, client *personalize.Client) {
	input := &personalize.CreateSolutionInput{
		// DatasetGroupArn: *string, // Required
		// Name: *string, // Required
	}

	if len(_personalizeDatasetGroupArn) > 0 {
		input.DatasetGroupArn = aws.String(_personalizeDatasetGroupArn)
	}
	if len(_personalizeName) > 0 {
		input.Name = aws.String(_personalizeName)
	}
	if len(_personalizeEventType) > 0 {
		input.EventType = aws.String(_personalizeEventType)
	}
	if len(_personalizePerformAutoML) > 0 {
		if err := assignInputField(input, "PerformAutoML", _personalizePerformAutoML); err != nil {
			log.Errorf("invalid --perform-auto-ml: %s", err.Error())
			return
		}
	}
	if len(_personalizePerformAutoTraining) > 0 {
		if err := assignInputField(input, "PerformAutoTraining", _personalizePerformAutoTraining); err != nil {
			log.Errorf("invalid --perform-auto-training: %s", err.Error())
			return
		}
	}
	if len(_personalizePerformHPO) > 0 {
		if err := assignInputField(input, "PerformHPO", _personalizePerformHPO); err != nil {
			log.Errorf("invalid --perform-hpo: %s", err.Error())
			return
		}
	}
	if len(_personalizePerformIncrementalUpdate) > 0 {
		if err := assignInputField(input, "PerformIncrementalUpdate", _personalizePerformIncrementalUpdate); err != nil {
			log.Errorf("invalid --perform-incremental-update: %s", err.Error())
			return
		}
	}
	if len(_personalizeRecipeArn) > 0 {
		input.RecipeArn = aws.String(_personalizeRecipeArn)
	}
	if len(_personalizeSolutionConfig) > 0 {
		if err := assignInputField(input, "SolutionConfig", _personalizeSolutionConfig); err != nil {
			log.Errorf("invalid --solution-config: %s", err.Error())
			return
		}
	}
	if len(_personalizeTags) > 0 {
		if err := assignInputField(input, "Tags", _personalizeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSolution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Trains or retrains an active solution in a Custom dataset group. A solution is
// created using the [CreateSolution]operation and must be in the ACTIVE state before calling
// CreateSolutionVersion . A new version of the solution is created every time you
// call this operation.
//
// # Status
//
// A solution version can be in one of the following states:
//
// - CREATE PENDING
//
// - CREATE IN_PROGRESS
//
// - ACTIVE
//
// - CREATE FAILED
//
// - CREATE STOPPING
//
// - CREATE STOPPED
//
// To get the status of the version, call [DescribeSolutionVersion]. Wait until the status shows as ACTIVE
// before calling CreateCampaign .
//
// If the status shows as CREATE FAILED, the response includes a failureReason
// key, which describes why the job failed.
//
// # Related APIs
//
// [ListSolutionVersions]
//
// [DescribeSolutionVersion]
//
// [ListSolutions]
//
// [CreateSolution]
//
// [DescribeSolution]
//
// [DeleteSolution]
//
// [DescribeSolutionVersion]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeSolutionVersion.html
// [DeleteSolution]: https://docs.aws.amazon.com/personalize/latest/dg/API_DeleteSolution.html
// [CreateSolution]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateSolution.html
// [ListSolutionVersions]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListSolutionVersions.html
// [ListSolutions]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListSolutions.html
// [DescribeSolution]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeSolution.html
func personalize_CreateSolutionVersion(cfg aws.Config, client *personalize.Client) {
	input := &personalize.CreateSolutionVersionInput{
		// SolutionArn: *string, // Required
	}

	if len(_personalizeSolutionArn) > 0 {
		input.SolutionArn = aws.String(_personalizeSolutionArn)
	}
	if len(_personalizeName) > 0 {
		input.Name = aws.String(_personalizeName)
	}
	if len(_personalizeTags) > 0 {
		if err := assignInputField(input, "Tags", _personalizeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_personalizeTrainingMode) > 0 {
		if err := assignInputField(input, "TrainingMode", _personalizeTrainingMode); err != nil {
			log.Errorf("invalid --training-mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSolutionVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a campaign by deleting the solution deployment. The solution that the
// campaign is based on is not deleted and can be redeployed when needed. A deleted
// campaign can no longer be specified in a [GetRecommendations]request. For information on creating
// campaigns, see [CreateCampaign].
//
// [CreateCampaign]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateCampaign.html
// [GetRecommendations]: https://docs.aws.amazon.com/personalize/latest/dg/API_RS_GetRecommendations.html
func personalize_DeleteCampaign(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DeleteCampaignInput{
		// CampaignArn: *string, // Required
	}

	if len(_personalizeCampaignArn) > 0 {
		input.CampaignArn = aws.String(_personalizeCampaignArn)
	}

	if resp, err := client.DeleteCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a dataset. You can't delete a dataset if an associated DatasetImportJob
// or SolutionVersion is in the CREATE PENDING or IN PROGRESS state. For more
// information about deleting datasets, see [Deleting a dataset].
//
// [Deleting a dataset]: https://docs.aws.amazon.com/personalize/latest/dg/delete-dataset.html
func personalize_DeleteDataset(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DeleteDatasetInput{
		// DatasetArn: *string, // Required
	}

	if len(_personalizeDatasetArn) > 0 {
		input.DatasetArn = aws.String(_personalizeDatasetArn)
	}

	if resp, err := client.DeleteDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a dataset group. Before you delete a dataset group, you must delete the
// following:
//
// - All associated event trackers.
//
// - All associated solutions.
//
// - All datasets in the dataset group.
func personalize_DeleteDatasetGroup(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DeleteDatasetGroupInput{
		// DatasetGroupArn: *string, // Required
	}

	if len(_personalizeDatasetGroupArn) > 0 {
		input.DatasetGroupArn = aws.String(_personalizeDatasetGroupArn)
	}

	if resp, err := client.DeleteDatasetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the event tracker. Does not delete the dataset from the dataset group.
// For more information on event trackers, see [CreateEventTracker].
//
// [CreateEventTracker]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateEventTracker.html
func personalize_DeleteEventTracker(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DeleteEventTrackerInput{
		// EventTrackerArn: *string, // Required
	}

	if len(_personalizeEventTrackerArn) > 0 {
		input.EventTrackerArn = aws.String(_personalizeEventTrackerArn)
	}

	if resp, err := client.DeleteEventTracker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a filter.
func personalize_DeleteFilter(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DeleteFilterInput{
		// FilterArn: *string, // Required
	}

	if len(_personalizeFilterArn) > 0 {
		input.FilterArn = aws.String(_personalizeFilterArn)
	}

	if resp, err := client.DeleteFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a metric attribution.
func personalize_DeleteMetricAttribution(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DeleteMetricAttributionInput{
		// MetricAttributionArn: *string, // Required
	}

	if len(_personalizeMetricAttributionArn) > 0 {
		input.MetricAttributionArn = aws.String(_personalizeMetricAttributionArn)
	}

	if resp, err := client.DeleteMetricAttribution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deactivates and removes a recommender. A deleted recommender can no longer be
// specified in a [GetRecommendations]request.
//
// [GetRecommendations]: https://docs.aws.amazon.com/personalize/latest/dg/API_RS_GetRecommendations.html
func personalize_DeleteRecommender(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DeleteRecommenderInput{
		// RecommenderArn: *string, // Required
	}

	if len(_personalizeRecommenderArn) > 0 {
		input.RecommenderArn = aws.String(_personalizeRecommenderArn)
	}

	if resp, err := client.DeleteRecommender(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a schema. Before deleting a schema, you must delete all datasets
// referencing the schema. For more information on schemas, see [CreateSchema].
//
// [CreateSchema]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateSchema.html
func personalize_DeleteSchema(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DeleteSchemaInput{
		// SchemaArn: *string, // Required
	}

	if len(_personalizeSchemaArn) > 0 {
		input.SchemaArn = aws.String(_personalizeSchemaArn)
	}

	if resp, err := client.DeleteSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes all versions of a solution and the Solution object itself. Before
// deleting a solution, you must delete all campaigns based on the solution. To
// determine what campaigns are using the solution, call [ListCampaigns]and supply the Amazon
// Resource Name (ARN) of the solution. You can't delete a solution if an
// associated SolutionVersion is in the CREATE PENDING or IN PROGRESS state. For
// more information on solutions, see [CreateSolution].
//
// [CreateSolution]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateSolution.html
// [ListCampaigns]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListCampaigns.html
func personalize_DeleteSolution(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DeleteSolutionInput{
		// SolutionArn: *string, // Required
	}

	if len(_personalizeSolutionArn) > 0 {
		input.SolutionArn = aws.String(_personalizeSolutionArn)
	}

	if resp, err := client.DeleteSolution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the given algorithm.
func personalize_DescribeAlgorithm(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DescribeAlgorithmInput{
		// AlgorithmArn: *string, // Required
	}

	if len(_personalizeAlgorithmArn) > 0 {
		input.AlgorithmArn = aws.String(_personalizeAlgorithmArn)
	}

	if resp, err := client.DescribeAlgorithm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the properties of a batch inference job including name, Amazon Resource
// Name (ARN), status, input and output configurations, and the ARN of the solution
// version used to generate the recommendations.
func personalize_DescribeBatchInferenceJob(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DescribeBatchInferenceJobInput{
		// BatchInferenceJobArn: *string, // Required
	}

	if len(_personalizeBatchInferenceJobArn) > 0 {
		input.BatchInferenceJobArn = aws.String(_personalizeBatchInferenceJobArn)
	}

	if resp, err := client.DescribeBatchInferenceJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the properties of a batch segment job including name, Amazon Resource Name
// (ARN), status, input and output configurations, and the ARN of the solution
// version used to generate segments.
func personalize_DescribeBatchSegmentJob(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DescribeBatchSegmentJobInput{
		// BatchSegmentJobArn: *string, // Required
	}

	if len(_personalizeBatchSegmentJobArn) > 0 {
		input.BatchSegmentJobArn = aws.String(_personalizeBatchSegmentJobArn)
	}

	if resp, err := client.DescribeBatchSegmentJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the given campaign, including its status.
// A campaign can be in one of the following states:
//
// - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// - DELETE PENDING > DELETE IN_PROGRESS
//
// When the status is CREATE FAILED , the response includes the failureReason key,
// which describes why.
//
// For more information on campaigns, see [CreateCampaign].
//
// [CreateCampaign]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateCampaign.html
func personalize_DescribeCampaign(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DescribeCampaignInput{
		// CampaignArn: *string, // Required
	}

	if len(_personalizeCampaignArn) > 0 {
		input.CampaignArn = aws.String(_personalizeCampaignArn)
	}

	if resp, err := client.DescribeCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the data deletion job created by [CreateDataDeletionJob], including the job status.
//
// [CreateDataDeletionJob]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDataDeletionJob.html
func personalize_DescribeDataDeletionJob(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DescribeDataDeletionJobInput{
		// DataDeletionJobArn: *string, // Required
	}

	if len(_personalizeDataDeletionJobArn) > 0 {
		input.DataDeletionJobArn = aws.String(_personalizeDataDeletionJobArn)
	}

	if resp, err := client.DescribeDataDeletionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the given dataset. For more information on datasets, see [CreateDataset].
//
// [CreateDataset]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDataset.html
func personalize_DescribeDataset(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DescribeDatasetInput{
		// DatasetArn: *string, // Required
	}

	if len(_personalizeDatasetArn) > 0 {
		input.DatasetArn = aws.String(_personalizeDatasetArn)
	}

	if resp, err := client.DescribeDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the dataset export job created by [CreateDatasetExportJob], including the export job status.
//
// [CreateDatasetExportJob]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDatasetExportJob.html
func personalize_DescribeDatasetExportJob(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DescribeDatasetExportJobInput{
		// DatasetExportJobArn: *string, // Required
	}

	if len(_personalizeDatasetExportJobArn) > 0 {
		input.DatasetExportJobArn = aws.String(_personalizeDatasetExportJobArn)
	}

	if resp, err := client.DescribeDatasetExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the given dataset group. For more information on dataset groups, see [CreateDatasetGroup].
//
// [CreateDatasetGroup]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDatasetGroup.html
func personalize_DescribeDatasetGroup(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DescribeDatasetGroupInput{
		// DatasetGroupArn: *string, // Required
	}

	if len(_personalizeDatasetGroupArn) > 0 {
		input.DatasetGroupArn = aws.String(_personalizeDatasetGroupArn)
	}

	if resp, err := client.DescribeDatasetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the dataset import job created by [CreateDatasetImportJob], including the import job status.
//
// [CreateDatasetImportJob]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDatasetImportJob.html
func personalize_DescribeDatasetImportJob(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DescribeDatasetImportJobInput{
		// DatasetImportJobArn: *string, // Required
	}

	if len(_personalizeDatasetImportJobArn) > 0 {
		input.DatasetImportJobArn = aws.String(_personalizeDatasetImportJobArn)
	}

	if resp, err := client.DescribeDatasetImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an event tracker. The response includes the trackingId and status of
// the event tracker. For more information on event trackers, see [CreateEventTracker].
//
// [CreateEventTracker]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateEventTracker.html
func personalize_DescribeEventTracker(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DescribeEventTrackerInput{
		// EventTrackerArn: *string, // Required
	}

	if len(_personalizeEventTrackerArn) > 0 {
		input.EventTrackerArn = aws.String(_personalizeEventTrackerArn)
	}

	if resp, err := client.DescribeEventTracker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the given feature transformation.
func personalize_DescribeFeatureTransformation(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DescribeFeatureTransformationInput{
		// FeatureTransformationArn: *string, // Required
	}

	if len(_personalizeFeatureTransformationArn) > 0 {
		input.FeatureTransformationArn = aws.String(_personalizeFeatureTransformationArn)
	}

	if resp, err := client.DescribeFeatureTransformation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a filter's properties.
func personalize_DescribeFilter(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DescribeFilterInput{
		// FilterArn: *string, // Required
	}

	if len(_personalizeFilterArn) > 0 {
		input.FilterArn = aws.String(_personalizeFilterArn)
	}

	if resp, err := client.DescribeFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a metric attribution.
func personalize_DescribeMetricAttribution(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DescribeMetricAttributionInput{
		// MetricAttributionArn: *string, // Required
	}

	if len(_personalizeMetricAttributionArn) > 0 {
		input.MetricAttributionArn = aws.String(_personalizeMetricAttributionArn)
	}

	if resp, err := client.DescribeMetricAttribution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a recipe.
// A recipe contains three items:
//
// - An algorithm that trains a model.
//
// - Hyperparameters that govern the training.
//
// - Feature transformation information for modifying the input data before
// training.
//
// Amazon Personalize provides a set of predefined recipes. You specify a recipe
// when you create a solution with the [CreateSolution]API. CreateSolution trains a model by using
// the algorithm in the specified recipe and a training dataset. The solution, when
// deployed as a campaign, can provide recommendations using the [GetRecommendations]API.
//
// [CreateSolution]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateSolution.html
// [GetRecommendations]: https://docs.aws.amazon.com/personalize/latest/dg/API_RS_GetRecommendations.html
func personalize_DescribeRecipe(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DescribeRecipeInput{
		// RecipeArn: *string, // Required
	}

	if len(_personalizeRecipeArn) > 0 {
		input.RecipeArn = aws.String(_personalizeRecipeArn)
	}

	if resp, err := client.DescribeRecipe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the given recommender, including its status.
// A recommender can be in one of the following states:
//
// - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// - STOP PENDING > STOP IN_PROGRESS > INACTIVE > START PENDING > START
// IN_PROGRESS > ACTIVE
//
// - DELETE PENDING > DELETE IN_PROGRESS
//
// When the status is CREATE FAILED , the response includes the failureReason key,
// which describes why.
//
// The modelMetrics key is null when the recommender is being created or deleted.
//
// For more information on recommenders, see [CreateRecommender].
//
// [CreateRecommender]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateRecommender.html
func personalize_DescribeRecommender(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DescribeRecommenderInput{
		// RecommenderArn: *string, // Required
	}

	if len(_personalizeRecommenderArn) > 0 {
		input.RecommenderArn = aws.String(_personalizeRecommenderArn)
	}

	if resp, err := client.DescribeRecommender(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a schema. For more information on schemas, see [CreateSchema].
//
// [CreateSchema]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateSchema.html
func personalize_DescribeSchema(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DescribeSchemaInput{
		// SchemaArn: *string, // Required
	}

	if len(_personalizeSchemaArn) > 0 {
		input.SchemaArn = aws.String(_personalizeSchemaArn)
	}

	if resp, err := client.DescribeSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a solution. For more information on solutions, see [CreateSolution].
//
// [CreateSolution]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateSolution.html
func personalize_DescribeSolution(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DescribeSolutionInput{
		// SolutionArn: *string, // Required
	}

	if len(_personalizeSolutionArn) > 0 {
		input.SolutionArn = aws.String(_personalizeSolutionArn)
	}

	if resp, err := client.DescribeSolution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a specific version of a solution. For more information on solutions,
// see [CreateSolution]
//
// [CreateSolution]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateSolution.html
func personalize_DescribeSolutionVersion(cfg aws.Config, client *personalize.Client) {
	input := &personalize.DescribeSolutionVersionInput{
		// SolutionVersionArn: *string, // Required
	}

	if len(_personalizeSolutionVersionArn) > 0 {
		input.SolutionVersionArn = aws.String(_personalizeSolutionVersionArn)
	}

	if resp, err := client.DescribeSolutionVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the metrics for the specified solution version.
func personalize_GetSolutionMetrics(cfg aws.Config, client *personalize.Client) {
	input := &personalize.GetSolutionMetricsInput{
		// SolutionVersionArn: *string, // Required
	}

	if len(_personalizeSolutionVersionArn) > 0 {
		input.SolutionVersionArn = aws.String(_personalizeSolutionVersionArn)
	}

	if resp, err := client.GetSolutionMetrics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of the batch inference jobs that have been performed off of a
// solution version.
func personalize_ListBatchInferenceJobs(cfg aws.Config, client *personalize.Client) {
	input := &personalize.ListBatchInferenceJobsInput{}

	if len(_personalizeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _personalizeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_personalizeNextToken) > 0 {
		input.NextToken = aws.String(_personalizeNextToken)
	}
	if len(_personalizeSolutionVersionArn) > 0 {
		input.SolutionVersionArn = aws.String(_personalizeSolutionVersionArn)
	}

	if disablePaginator() {
		if resp, err := client.ListBatchInferenceJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*personalize.ListBatchInferenceJobsOutput
	p := personalize.NewListBatchInferenceJobsPaginator(client, input)
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

// Gets a list of the batch segment jobs that have been performed off of a
// solution version that you specify.
func personalize_ListBatchSegmentJobs(cfg aws.Config, client *personalize.Client) {
	input := &personalize.ListBatchSegmentJobsInput{}

	if len(_personalizeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _personalizeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_personalizeNextToken) > 0 {
		input.NextToken = aws.String(_personalizeNextToken)
	}
	if len(_personalizeSolutionVersionArn) > 0 {
		input.SolutionVersionArn = aws.String(_personalizeSolutionVersionArn)
	}

	if disablePaginator() {
		if resp, err := client.ListBatchSegmentJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*personalize.ListBatchSegmentJobsOutput
	p := personalize.NewListBatchSegmentJobsPaginator(client, input)
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

// Returns a list of campaigns that use the given solution. When a solution is not
// specified, all the campaigns associated with the account are listed. The
// response provides the properties for each campaign, including the Amazon
// Resource Name (ARN). For more information on campaigns, see [CreateCampaign].
//
// [CreateCampaign]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateCampaign.html
func personalize_ListCampaigns(cfg aws.Config, client *personalize.Client) {
	input := &personalize.ListCampaignsInput{}

	if len(_personalizeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _personalizeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_personalizeNextToken) > 0 {
		input.NextToken = aws.String(_personalizeNextToken)
	}
	if len(_personalizeSolutionArn) > 0 {
		input.SolutionArn = aws.String(_personalizeSolutionArn)
	}

	if disablePaginator() {
		if resp, err := client.ListCampaigns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*personalize.ListCampaignsOutput
	p := personalize.NewListCampaignsPaginator(client, input)
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

// Returns a list of data deletion jobs for a dataset group ordered by creation
// time, with the most recent first. When a dataset group is not specified, all the
// data deletion jobs associated with the account are listed. The response provides
// the properties for each job, including the Amazon Resource Name (ARN). For more
// information on data deletion jobs, see [Deleting users].
//
// [Deleting users]: https://docs.aws.amazon.com/personalize/latest/dg/delete-records.html
func personalize_ListDataDeletionJobs(cfg aws.Config, client *personalize.Client) {
	input := &personalize.ListDataDeletionJobsInput{}

	if len(_personalizeDatasetGroupArn) > 0 {
		input.DatasetGroupArn = aws.String(_personalizeDatasetGroupArn)
	}
	if len(_personalizeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _personalizeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_personalizeNextToken) > 0 {
		input.NextToken = aws.String(_personalizeNextToken)
	}

	if resp, err := client.ListDataDeletionJobs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of dataset export jobs that use the given dataset. When a
// dataset is not specified, all the dataset export jobs associated with the
// account are listed. The response provides the properties for each dataset export
// job, including the Amazon Resource Name (ARN). For more information on dataset
// export jobs, see [CreateDatasetExportJob]. For more information on datasets, see [CreateDataset].
//
// [CreateDataset]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDataset.html
// [CreateDatasetExportJob]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDatasetExportJob.html
func personalize_ListDatasetExportJobs(cfg aws.Config, client *personalize.Client) {
	input := &personalize.ListDatasetExportJobsInput{}

	if len(_personalizeDatasetArn) > 0 {
		input.DatasetArn = aws.String(_personalizeDatasetArn)
	}
	if len(_personalizeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _personalizeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_personalizeNextToken) > 0 {
		input.NextToken = aws.String(_personalizeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDatasetExportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*personalize.ListDatasetExportJobsOutput
	p := personalize.NewListDatasetExportJobsPaginator(client, input)
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

// Returns a list of dataset groups. The response provides the properties for each
// dataset group, including the Amazon Resource Name (ARN). For more information on
// dataset groups, see [CreateDatasetGroup].
//
// [CreateDatasetGroup]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDatasetGroup.html
func personalize_ListDatasetGroups(cfg aws.Config, client *personalize.Client) {
	input := &personalize.ListDatasetGroupsInput{}

	if len(_personalizeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _personalizeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_personalizeNextToken) > 0 {
		input.NextToken = aws.String(_personalizeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDatasetGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*personalize.ListDatasetGroupsOutput
	p := personalize.NewListDatasetGroupsPaginator(client, input)
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

// Returns a list of dataset import jobs that use the given dataset. When a
// dataset is not specified, all the dataset import jobs associated with the
// account are listed. The response provides the properties for each dataset import
// job, including the Amazon Resource Name (ARN). For more information on dataset
// import jobs, see [CreateDatasetImportJob]. For more information on datasets, see [CreateDataset].
//
// [CreateDataset]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDataset.html
// [CreateDatasetImportJob]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDatasetImportJob.html
func personalize_ListDatasetImportJobs(cfg aws.Config, client *personalize.Client) {
	input := &personalize.ListDatasetImportJobsInput{}

	if len(_personalizeDatasetArn) > 0 {
		input.DatasetArn = aws.String(_personalizeDatasetArn)
	}
	if len(_personalizeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _personalizeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_personalizeNextToken) > 0 {
		input.NextToken = aws.String(_personalizeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDatasetImportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*personalize.ListDatasetImportJobsOutput
	p := personalize.NewListDatasetImportJobsPaginator(client, input)
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

// Returns the list of datasets contained in the given dataset group. The response
// provides the properties for each dataset, including the Amazon Resource Name
// (ARN). For more information on datasets, see [CreateDataset].
//
// [CreateDataset]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDataset.html
func personalize_ListDatasets(cfg aws.Config, client *personalize.Client) {
	input := &personalize.ListDatasetsInput{}

	if len(_personalizeDatasetGroupArn) > 0 {
		input.DatasetGroupArn = aws.String(_personalizeDatasetGroupArn)
	}
	if len(_personalizeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _personalizeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_personalizeNextToken) > 0 {
		input.NextToken = aws.String(_personalizeNextToken)
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

	var results []*personalize.ListDatasetsOutput
	p := personalize.NewListDatasetsPaginator(client, input)
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

// Returns the list of event trackers associated with the account. The response
// provides the properties for each event tracker, including the Amazon Resource
// Name (ARN) and tracking ID. For more information on event trackers, see [CreateEventTracker].
//
// [CreateEventTracker]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateEventTracker.html
func personalize_ListEventTrackers(cfg aws.Config, client *personalize.Client) {
	input := &personalize.ListEventTrackersInput{}

	if len(_personalizeDatasetGroupArn) > 0 {
		input.DatasetGroupArn = aws.String(_personalizeDatasetGroupArn)
	}
	if len(_personalizeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _personalizeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_personalizeNextToken) > 0 {
		input.NextToken = aws.String(_personalizeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEventTrackers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*personalize.ListEventTrackersOutput
	p := personalize.NewListEventTrackersPaginator(client, input)
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

// Lists all filters that belong to a given dataset group.
func personalize_ListFilters(cfg aws.Config, client *personalize.Client) {
	input := &personalize.ListFiltersInput{}

	if len(_personalizeDatasetGroupArn) > 0 {
		input.DatasetGroupArn = aws.String(_personalizeDatasetGroupArn)
	}
	if len(_personalizeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _personalizeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_personalizeNextToken) > 0 {
		input.NextToken = aws.String(_personalizeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFilters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*personalize.ListFiltersOutput
	p := personalize.NewListFiltersPaginator(client, input)
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

// Lists the metrics for the metric attribution.
func personalize_ListMetricAttributionMetrics(cfg aws.Config, client *personalize.Client) {
	input := &personalize.ListMetricAttributionMetricsInput{}

	if len(_personalizeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _personalizeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_personalizeMetricAttributionArn) > 0 {
		input.MetricAttributionArn = aws.String(_personalizeMetricAttributionArn)
	}
	if len(_personalizeNextToken) > 0 {
		input.NextToken = aws.String(_personalizeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMetricAttributionMetrics(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*personalize.ListMetricAttributionMetricsOutput
	p := personalize.NewListMetricAttributionMetricsPaginator(client, input)
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

// Lists metric attributions.
func personalize_ListMetricAttributions(cfg aws.Config, client *personalize.Client) {
	input := &personalize.ListMetricAttributionsInput{}

	if len(_personalizeDatasetGroupArn) > 0 {
		input.DatasetGroupArn = aws.String(_personalizeDatasetGroupArn)
	}
	if len(_personalizeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _personalizeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_personalizeNextToken) > 0 {
		input.NextToken = aws.String(_personalizeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMetricAttributions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*personalize.ListMetricAttributionsOutput
	p := personalize.NewListMetricAttributionsPaginator(client, input)
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

// Returns a list of available recipes. The response provides the properties for
// each recipe, including the recipe's Amazon Resource Name (ARN).
func personalize_ListRecipes(cfg aws.Config, client *personalize.Client) {
	input := &personalize.ListRecipesInput{}

	if len(_personalizeDomain) > 0 {
		if err := assignInputField(input, "Domain", _personalizeDomain); err != nil {
			log.Errorf("invalid --domain: %s", err.Error())
			return
		}
	}
	if len(_personalizeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _personalizeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_personalizeNextToken) > 0 {
		input.NextToken = aws.String(_personalizeNextToken)
	}
	if len(_personalizeRecipeProvider) > 0 {
		if err := assignInputField(input, "RecipeProvider", _personalizeRecipeProvider); err != nil {
			log.Errorf("invalid --recipe-provider: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRecipes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*personalize.ListRecipesOutput
	p := personalize.NewListRecipesPaginator(client, input)
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

// Returns a list of recommenders in a given Domain dataset group. When a Domain
// dataset group is not specified, all the recommenders associated with the account
// are listed. The response provides the properties for each recommender, including
// the Amazon Resource Name (ARN). For more information on recommenders, see [CreateRecommender].
//
// [CreateRecommender]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateRecommender.html
func personalize_ListRecommenders(cfg aws.Config, client *personalize.Client) {
	input := &personalize.ListRecommendersInput{}

	if len(_personalizeDatasetGroupArn) > 0 {
		input.DatasetGroupArn = aws.String(_personalizeDatasetGroupArn)
	}
	if len(_personalizeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _personalizeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_personalizeNextToken) > 0 {
		input.NextToken = aws.String(_personalizeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRecommenders(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*personalize.ListRecommendersOutput
	p := personalize.NewListRecommendersPaginator(client, input)
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

// Returns the list of schemas associated with the account. The response provides
// the properties for each schema, including the Amazon Resource Name (ARN). For
// more information on schemas, see [CreateSchema].
//
// [CreateSchema]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateSchema.html
func personalize_ListSchemas(cfg aws.Config, client *personalize.Client) {
	input := &personalize.ListSchemasInput{}

	if len(_personalizeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _personalizeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_personalizeNextToken) > 0 {
		input.NextToken = aws.String(_personalizeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSchemas(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*personalize.ListSchemasOutput
	p := personalize.NewListSchemasPaginator(client, input)
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

// Returns a list of solution versions for the given solution. When a solution is
// not specified, all the solution versions associated with the account are listed.
// The response provides the properties for each solution version, including the
// Amazon Resource Name (ARN).
func personalize_ListSolutionVersions(cfg aws.Config, client *personalize.Client) {
	input := &personalize.ListSolutionVersionsInput{}

	if len(_personalizeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _personalizeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_personalizeNextToken) > 0 {
		input.NextToken = aws.String(_personalizeNextToken)
	}
	if len(_personalizeSolutionArn) > 0 {
		input.SolutionArn = aws.String(_personalizeSolutionArn)
	}

	if disablePaginator() {
		if resp, err := client.ListSolutionVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*personalize.ListSolutionVersionsOutput
	p := personalize.NewListSolutionVersionsPaginator(client, input)
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

// Returns a list of solutions in a given dataset group. When a dataset group is
// not specified, all the solutions associated with the account are listed. The
// response provides the properties for each solution, including the Amazon
// Resource Name (ARN). For more information on solutions, see [CreateSolution].
//
// [CreateSolution]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateSolution.html
func personalize_ListSolutions(cfg aws.Config, client *personalize.Client) {
	input := &personalize.ListSolutionsInput{}

	if len(_personalizeDatasetGroupArn) > 0 {
		input.DatasetGroupArn = aws.String(_personalizeDatasetGroupArn)
	}
	if len(_personalizeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _personalizeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_personalizeNextToken) > 0 {
		input.NextToken = aws.String(_personalizeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSolutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*personalize.ListSolutionsOutput
	p := personalize.NewListSolutionsPaginator(client, input)
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

// Get a list of [tags] attached to a resource.
//
// [tags]: https://docs.aws.amazon.com/personalize/latest/dg/tagging-resources.html
func personalize_ListTagsForResource(cfg aws.Config, client *personalize.Client) {
	input := &personalize.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_personalizeResourceArn) > 0 {
		input.ResourceArn = aws.String(_personalizeResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a recommender that is INACTIVE. Starting a recommender does not create
// any new models, but resumes billing and automatic retraining for the
// recommender.
func personalize_StartRecommender(cfg aws.Config, client *personalize.Client) {
	input := &personalize.StartRecommenderInput{
		// RecommenderArn: *string, // Required
	}

	if len(_personalizeRecommenderArn) > 0 {
		input.RecommenderArn = aws.String(_personalizeRecommenderArn)
	}

	if resp, err := client.StartRecommender(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a recommender that is ACTIVE. Stopping a recommender halts billing and
// automatic retraining for the recommender.
func personalize_StopRecommender(cfg aws.Config, client *personalize.Client) {
	input := &personalize.StopRecommenderInput{
		// RecommenderArn: *string, // Required
	}

	if len(_personalizeRecommenderArn) > 0 {
		input.RecommenderArn = aws.String(_personalizeRecommenderArn)
	}

	if resp, err := client.StopRecommender(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops creating a solution version that is in a state of CREATE_PENDING or
// CREATE IN_PROGRESS.
//
// Depending on the current state of the solution version, the solution version
// state changes as follows:
//
// - CREATE_PENDING > CREATE_STOPPED
//
// or
//
// - CREATE_IN_PROGRESS > CREATE_STOPPING > CREATE_STOPPED
//
// You are billed for all of the training completed up until you stop the solution
// version creation. You cannot resume creating a solution version once it has been
// stopped.
func personalize_StopSolutionVersionCreation(cfg aws.Config, client *personalize.Client) {
	input := &personalize.StopSolutionVersionCreationInput{
		// SolutionVersionArn: *string, // Required
	}

	if len(_personalizeSolutionVersionArn) > 0 {
		input.SolutionVersionArn = aws.String(_personalizeSolutionVersionArn)
	}

	if resp, err := client.StopSolutionVersionCreation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add a list of tags to a resource.
func personalize_TagResource(cfg aws.Config, client *personalize.Client) {
	input := &personalize.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_personalizeResourceArn) > 0 {
		input.ResourceArn = aws.String(_personalizeResourceArn)
	}
	if len(_personalizeTags) > 0 {
		if err := assignInputField(input, "Tags", _personalizeTags); err != nil {
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

// Removes the specified tags that are attached to a resource. For more
// information, see [Removing tags from Amazon Personalize resources].
//
// [Removing tags from Amazon Personalize resources]: https://docs.aws.amazon.com/personalize/latest/dg/tags-remove.html
func personalize_UntagResource(cfg aws.Config, client *personalize.Client) {
	input := &personalize.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_personalizeResourceArn) > 0 {
		input.ResourceArn = aws.String(_personalizeResourceArn)
	}
	if len(_personalizeTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _personalizeTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a campaign to deploy a retrained solution version with an existing
// campaign, change your campaign's minProvisionedTPS , or modify your campaign's
// configuration. For example, you can set enableMetadataWithRecommendations to
// true for an existing campaign.
//
// To update a campaign to start automatically using the latest solution version,
// specify the following:
//
// - For the SolutionVersionArn parameter, specify the Amazon Resource Name (ARN)
// of your solution in SolutionArn/$LATEST format.
//
// - In the campaignConfig , set syncWithLatestSolutionVersion to true .
//
// To update a campaign, the campaign status must be ACTIVE or CREATE FAILED.
// Check the campaign status using the [DescribeCampaign]operation.
//
// You can still get recommendations from a campaign while an update is in
// progress. The campaign will use the previous solution version and campaign
// configuration to generate recommendations until the latest campaign update
// status is Active .
//
// For more information about updating a campaign, including code samples, see [Updating a campaign].
// For more information about campaigns, see [Creating a campaign].
//
// [Creating a campaign]: https://docs.aws.amazon.com/personalize/latest/dg/campaigns.html
// [Updating a campaign]: https://docs.aws.amazon.com/personalize/latest/dg/update-campaigns.html
// [DescribeCampaign]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeCampaign.html
func personalize_UpdateCampaign(cfg aws.Config, client *personalize.Client) {
	input := &personalize.UpdateCampaignInput{
		// CampaignArn: *string, // Required
	}

	if len(_personalizeCampaignArn) > 0 {
		input.CampaignArn = aws.String(_personalizeCampaignArn)
	}
	if len(_personalizeCampaignConfig) > 0 {
		if err := assignInputField(input, "CampaignConfig", _personalizeCampaignConfig); err != nil {
			log.Errorf("invalid --campaign-config: %s", err.Error())
			return
		}
	}
	if len(_personalizeMinProvisionedTPS) > 0 {
		if err := assignInputField(input, "MinProvisionedTPS", _personalizeMinProvisionedTPS); err != nil {
			log.Errorf("invalid --min-provisioned-tps: %s", err.Error())
			return
		}
	}
	if len(_personalizeSolutionVersionArn) > 0 {
		input.SolutionVersionArn = aws.String(_personalizeSolutionVersionArn)
	}

	if resp, err := client.UpdateCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a dataset to replace its schema with a new or existing one. For more
// information, see [Replacing a dataset's schema].
//
// [Replacing a dataset's schema]: https://docs.aws.amazon.com/personalize/latest/dg/updating-dataset-schema.html
func personalize_UpdateDataset(cfg aws.Config, client *personalize.Client) {
	input := &personalize.UpdateDatasetInput{
		// DatasetArn: *string, // Required
		// SchemaArn: *string, // Required
	}

	if len(_personalizeDatasetArn) > 0 {
		input.DatasetArn = aws.String(_personalizeDatasetArn)
	}
	if len(_personalizeSchemaArn) > 0 {
		input.SchemaArn = aws.String(_personalizeSchemaArn)
	}

	if resp, err := client.UpdateDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a metric attribution.
func personalize_UpdateMetricAttribution(cfg aws.Config, client *personalize.Client) {
	input := &personalize.UpdateMetricAttributionInput{}

	if len(_personalizeAddMetrics) > 0 {
		if err := assignInputField(input, "AddMetrics", _personalizeAddMetrics); err != nil {
			log.Errorf("invalid --add-metrics: %s", err.Error())
			return
		}
	}
	if len(_personalizeMetricAttributionArn) > 0 {
		input.MetricAttributionArn = aws.String(_personalizeMetricAttributionArn)
	}
	if len(_personalizeMetricsOutputConfig) > 0 {
		if err := assignInputField(input, "MetricsOutputConfig", _personalizeMetricsOutputConfig); err != nil {
			log.Errorf("invalid --metrics-output-config: %s", err.Error())
			return
		}
	}
	if len(_personalizeRemoveMetrics) > 0 {
		input.RemoveMetrics = append([]string(nil), _personalizeRemoveMetrics...)
	}

	if resp, err := client.UpdateMetricAttribution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the recommender to modify the recommender configuration. If you update
// the recommender to modify the columns used in training, Amazon Personalize
// automatically starts a full retraining of the models backing your recommender.
// While the update completes, you can still get recommendations from the
// recommender. The recommender uses the previous configuration until the update
// completes. To track the status of this update, use the latestRecommenderUpdate
// returned in the [DescribeRecommender]operation.
//
// [DescribeRecommender]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeRecommender.html
func personalize_UpdateRecommender(cfg aws.Config, client *personalize.Client) {
	input := &personalize.UpdateRecommenderInput{
		// RecommenderArn: *string, // Required
		// RecommenderConfig: *types.RecommenderConfig, // Required
	}

	if len(_personalizeRecommenderArn) > 0 {
		input.RecommenderArn = aws.String(_personalizeRecommenderArn)
	}
	if len(_personalizeRecommenderConfig) > 0 {
		if err := assignInputField(input, "RecommenderConfig", _personalizeRecommenderConfig); err != nil {
			log.Errorf("invalid --recommender-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRecommender(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Amazon Personalize solution to use a different automatic training
// configuration. When you update a solution, you can change whether the solution
// uses automatic training, and you can change the training frequency. For more
// information about updating a solution, see [Updating a solution].
//
// A solution update can be in one of the following states:
//
// CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// To get the status of a solution update, call the [DescribeSolution] API operation and find the
// status in the latestSolutionUpdate .
//
// [Updating a solution]: https://docs.aws.amazon.com/personalize/latest/dg/updating-solution.html
// [DescribeSolution]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeSolution.html
func personalize_UpdateSolution(cfg aws.Config, client *personalize.Client) {
	input := &personalize.UpdateSolutionInput{
		// SolutionArn: *string, // Required
	}

	if len(_personalizeSolutionArn) > 0 {
		input.SolutionArn = aws.String(_personalizeSolutionArn)
	}
	if len(_personalizePerformAutoTraining) > 0 {
		if err := assignInputField(input, "PerformAutoTraining", _personalizePerformAutoTraining); err != nil {
			log.Errorf("invalid --perform-auto-training: %s", err.Error())
			return
		}
	}
	if len(_personalizePerformIncrementalUpdate) > 0 {
		if err := assignInputField(input, "PerformIncrementalUpdate", _personalizePerformIncrementalUpdate); err != nil {
			log.Errorf("invalid --perform-incremental-update: %s", err.Error())
			return
		}
	}
	if len(_personalizeSolutionUpdateConfig) > 0 {
		if err := assignInputField(input, "SolutionUpdateConfig", _personalizeSolutionUpdateConfig); err != nil {
			log.Errorf("invalid --solution-update-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSolution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_personalizeCmd)
	_personalizeCmd.Flags().SortFlags = false

	_personalizeCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_personalizeCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_personalizeCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_personalizeCmd.Flags().StringVarP(&_personalizeAddMetrics, "add-metrics", "", "", "Add Metrics")
	_personalizeCmd.Flags().StringVarP(&_personalizeAlgorithmArn, "algorithm-arn", "", "", "Algorithm ARN")
	_personalizeCmd.Flags().StringVarP(&_personalizeBatchInferenceJobArn, "batch-inference-job-arn", "", "", "Batch Inference Job ARN")
	_personalizeCmd.Flags().StringVarP(&_personalizeBatchInferenceJobConfig, "batch-inference-job-config", "", "", "Batch Inference Job Config")
	_personalizeCmd.Flags().StringVarP(&_personalizeBatchInferenceJobMode, "batch-inference-job-mode", "", "", "Batch Inference Job Mode")
	_personalizeCmd.Flags().StringVarP(&_personalizeBatchSegmentJobArn, "batch-segment-job-arn", "", "", "Batch Segment Job ARN")
	_personalizeCmd.Flags().StringVarP(&_personalizeCampaignArn, "campaign-arn", "", "", "Campaign ARN")
	_personalizeCmd.Flags().StringVarP(&_personalizeCampaignConfig, "campaign-config", "", "", "Campaign Config")
	_personalizeCmd.Flags().StringVarP(&_personalizeDataDeletionJobArn, "data-deletion-job-arn", "", "", "Data Deletion Job ARN")
	_personalizeCmd.Flags().StringVarP(&_personalizeDataSource, "data-source", "", "", "Data Source")
	_personalizeCmd.Flags().StringVarP(&_personalizeDatasetArn, "dataset-arn", "", "", "Dataset ARN")
	_personalizeCmd.Flags().StringVarP(&_personalizeDatasetExportJobArn, "dataset-export-job-arn", "", "", "Dataset Export Job ARN")
	_personalizeCmd.Flags().StringVarP(&_personalizeDatasetGroupArn, "dataset-group-arn", "", "", "Dataset Group ARN")
	_personalizeCmd.Flags().StringVarP(&_personalizeDatasetImportJobArn, "dataset-import-job-arn", "", "", "Dataset Import Job ARN")
	_personalizeCmd.Flags().StringVarP(&_personalizeDatasetType, "dataset-type", "", "", "Dataset Type")
	_personalizeCmd.Flags().StringVarP(&_personalizeDomain, "domain", "", "", "Domain")
	_personalizeCmd.Flags().StringVarP(&_personalizeEventTrackerArn, "event-tracker-arn", "", "", "Event Tracker ARN")
	_personalizeCmd.Flags().StringVarP(&_personalizeEventType, "event-type", "", "", "Event Type")
	_personalizeCmd.Flags().StringVarP(&_personalizeFeatureTransformationArn, "feature-transformation-arn", "", "", "Feature Transformation ARN")
	_personalizeCmd.Flags().StringVarP(&_personalizeFilterArn, "filter-arn", "", "", "Filter ARN")
	_personalizeCmd.Flags().StringVarP(&_personalizeFilterExpression, "filter-expression", "", "", "Filter Expression")
	_personalizeCmd.Flags().StringVarP(&_personalizeImportMode, "import-mode", "", "", "Import Mode")
	_personalizeCmd.Flags().StringVarP(&_personalizeIngestionMode, "ingestion-mode", "", "", "Ingestion Mode")
	_personalizeCmd.Flags().StringVarP(&_personalizeJobInput, "job-input", "", "", "Job Input")
	_personalizeCmd.Flags().StringVarP(&_personalizeJobName, "job-name", "", "", "Job Name")
	_personalizeCmd.Flags().StringVarP(&_personalizeJobOutput, "job-output", "", "", "Job Output")
	_personalizeCmd.Flags().StringVarP(&_personalizeKmsKeyArn, "kms-key-arn", "", "", "KMS Key ARN")
	_personalizeCmd.Flags().StringVarP(&_personalizeMaxResults, "max-results", "", "", "Max Results")
	_personalizeCmd.Flags().StringVarP(&_personalizeMetricAttributionArn, "metric-attribution-arn", "", "", "Metric Attribution ARN")
	_personalizeCmd.Flags().StringVarP(&_personalizeMetrics, "metrics", "", "", "Metrics")
	_personalizeCmd.Flags().StringVarP(&_personalizeMetricsOutputConfig, "metrics-output-config", "", "", "Metrics Output Config")
	_personalizeCmd.Flags().StringVarP(&_personalizeMinProvisionedTPS, "min-provisioned-tps", "", "", "Min Provisioned Tps")
	_personalizeCmd.Flags().StringVarP(&_personalizeName, "name", "", "", "Name")
	_personalizeCmd.Flags().StringVarP(&_personalizeNextToken, "next-token", "", "", "Next Token")
	_personalizeCmd.Flags().StringVarP(&_personalizeNumResults, "num-results", "", "", "Num Results")
	_personalizeCmd.Flags().StringVarP(&_personalizePerformAutoML, "perform-auto-ml", "", "", "Perform Auto Ml")
	_personalizeCmd.Flags().StringVarP(&_personalizePerformAutoTraining, "perform-auto-training", "", "", "Perform Auto Training")
	_personalizeCmd.Flags().StringVarP(&_personalizePerformHPO, "perform-hpo", "", "", "Perform Hpo")
	_personalizeCmd.Flags().StringVarP(&_personalizePerformIncrementalUpdate, "perform-incremental-update", "", "", "Perform Incremental Update")
	_personalizeCmd.Flags().StringVarP(&_personalizePublishAttributionMetricsToS3, "publish-attribution-metrics-to-s3", "", "", "Publish Attribution Metrics To S3")
	_personalizeCmd.Flags().StringVarP(&_personalizeRecipeArn, "recipe-arn", "", "", "Recipe ARN")
	_personalizeCmd.Flags().StringVarP(&_personalizeRecipeProvider, "recipe-provider", "", "", "Recipe Provider")
	_personalizeCmd.Flags().StringVarP(&_personalizeRecommenderArn, "recommender-arn", "", "", "Recommender ARN")
	_personalizeCmd.Flags().StringVarP(&_personalizeRecommenderConfig, "recommender-config", "", "", "Recommender Config")
	_personalizeCmd.Flags().StringSliceVarP(&_personalizeRemoveMetrics, "remove-metrics", "", nil, "Remove Metrics")
	_personalizeCmd.Flags().StringVarP(&_personalizeResourceArn, "resource-arn", "", "", "Resource ARN")
	_personalizeCmd.Flags().StringVarP(&_personalizeRoleArn, "role-arn", "", "", "Role ARN")
	_personalizeCmd.Flags().StringVarP(&_personalizeSchema, "schema", "", "", "Schema")
	_personalizeCmd.Flags().StringVarP(&_personalizeSchemaArn, "schema-arn", "", "", "Schema ARN")
	_personalizeCmd.Flags().StringVarP(&_personalizeSolutionArn, "solution-arn", "", "", "Solution ARN")
	_personalizeCmd.Flags().StringVarP(&_personalizeSolutionConfig, "solution-config", "", "", "Solution Config")
	_personalizeCmd.Flags().StringVarP(&_personalizeSolutionUpdateConfig, "solution-update-config", "", "", "Solution Update Config")
	_personalizeCmd.Flags().StringVarP(&_personalizeSolutionVersionArn, "solution-version-arn", "", "", "Solution Version ARN")
	_personalizeCmd.Flags().StringSliceVarP(&_personalizeTagKeys, "tag-keys", "", nil, "Tag Keys")
	_personalizeCmd.Flags().StringVarP(&_personalizeTags, "tags", "", "", "Tags")
	_personalizeCmd.Flags().StringVarP(&_personalizeThemeGenerationConfig, "theme-generation-config", "", "", "Theme Generation Config")
	_personalizeCmd.Flags().StringVarP(&_personalizeTrainingMode, "training-mode", "", "", "Training Mode")

	_personalizeCmd.Flags().BoolVarP(&_personalizeCreateBatchInferenceJob, "create-batch-inference-job", "", false, "Create Batch Inference Job")
	_personalizeCmd.Flags().BoolVarP(&_personalizeCreateBatchSegmentJob, "create-batch-segment-job", "", false, "Create Batch Segment Job")
	_personalizeCmd.Flags().BoolVarP(&_personalizeCreateCampaign, "create-campaign", "", false, "Create Campaign")
	_personalizeCmd.Flags().BoolVarP(&_personalizeCreateDataDeletionJob, "create-data-deletion-job", "", false, "Create Data Deletion Job")
	_personalizeCmd.Flags().BoolVarP(&_personalizeCreateDataset, "create-dataset", "", false, "Create Dataset")
	_personalizeCmd.Flags().BoolVarP(&_personalizeCreateDatasetExportJob, "create-dataset-export-job", "", false, "Create Dataset Export Job")
	_personalizeCmd.Flags().BoolVarP(&_personalizeCreateDatasetGroup, "create-dataset-group", "", false, "Create Dataset Group")
	_personalizeCmd.Flags().BoolVarP(&_personalizeCreateDatasetImportJob, "create-dataset-import-job", "", false, "Create Dataset Import Job")
	_personalizeCmd.Flags().BoolVarP(&_personalizeCreateEventTracker, "create-event-tracker", "", false, "Create Event Tracker")
	_personalizeCmd.Flags().BoolVarP(&_personalizeCreateFilter, "create-filter", "", false, "Create Filter")
	_personalizeCmd.Flags().BoolVarP(&_personalizeCreateMetricAttribution, "create-metric-attribution", "", false, "Create Metric Attribution")
	_personalizeCmd.Flags().BoolVarP(&_personalizeCreateRecommender, "create-recommender", "", false, "Create Recommender")
	_personalizeCmd.Flags().BoolVarP(&_personalizeCreateSchema, "create-schema", "", false, "Create Schema")
	_personalizeCmd.Flags().BoolVarP(&_personalizeCreateSolution, "create-solution", "", false, "Create Solution")
	_personalizeCmd.Flags().BoolVarP(&_personalizeCreateSolutionVersion, "create-solution-version", "", false, "Create Solution Version")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDeleteCampaign, "delete-campaign", "", false, "Delete Campaign")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDeleteDataset, "delete-dataset", "", false, "Delete Dataset")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDeleteDatasetGroup, "delete-dataset-group", "", false, "Delete Dataset Group")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDeleteEventTracker, "delete-event-tracker", "", false, "Delete Event Tracker")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDeleteFilter, "delete-filter", "", false, "Delete Filter")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDeleteMetricAttribution, "delete-metric-attribution", "", false, "Delete Metric Attribution")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDeleteRecommender, "delete-recommender", "", false, "Delete Recommender")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDeleteSchema, "delete-schema", "", false, "Delete Schema")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDeleteSolution, "delete-solution", "", false, "Delete Solution")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDescribeAlgorithm, "describe-algorithm", "", false, "Describe Algorithm")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDescribeBatchInferenceJob, "describe-batch-inference-job", "", false, "Describe Batch Inference Job")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDescribeBatchSegmentJob, "describe-batch-segment-job", "", false, "Describe Batch Segment Job")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDescribeCampaign, "describe-campaign", "", false, "Describe Campaign")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDescribeDataDeletionJob, "describe-data-deletion-job", "", false, "Describe Data Deletion Job")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDescribeDataset, "describe-dataset", "", false, "Describe Dataset")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDescribeDatasetExportJob, "describe-dataset-export-job", "", false, "Describe Dataset Export Job")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDescribeDatasetGroup, "describe-dataset-group", "", false, "Describe Dataset Group")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDescribeDatasetImportJob, "describe-dataset-import-job", "", false, "Describe Dataset Import Job")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDescribeEventTracker, "describe-event-tracker", "", false, "Describe Event Tracker")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDescribeFeatureTransformation, "describe-feature-transformation", "", false, "Describe Feature Transformation")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDescribeFilter, "describe-filter", "", false, "Describe Filter")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDescribeMetricAttribution, "describe-metric-attribution", "", false, "Describe Metric Attribution")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDescribeRecipe, "describe-recipe", "", false, "Describe Recipe")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDescribeRecommender, "describe-recommender", "", false, "Describe Recommender")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDescribeSchema, "describe-schema", "", false, "Describe Schema")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDescribeSolution, "describe-solution", "", false, "Describe Solution")
	_personalizeCmd.Flags().BoolVarP(&_personalizeDescribeSolutionVersion, "describe-solution-version", "", false, "Describe Solution Version")
	_personalizeCmd.Flags().BoolVarP(&_personalizeGetSolutionMetrics, "get-solution-metrics", "", false, "Get Solution Metrics")
	_personalizeCmd.Flags().BoolVarP(&_personalizeListBatchInferenceJobs, "list-batch-inference-jobs", "", false, "List Batch Inference Jobs")
	_personalizeCmd.Flags().BoolVarP(&_personalizeListBatchSegmentJobs, "list-batch-segment-jobs", "", false, "List Batch Segment Jobs")
	_personalizeCmd.Flags().BoolVarP(&_personalizeListCampaigns, "list-campaigns", "", false, "List Campaigns")
	_personalizeCmd.Flags().BoolVarP(&_personalizeListDataDeletionJobs, "list-data-deletion-jobs", "", false, "List Data Deletion Jobs")
	_personalizeCmd.Flags().BoolVarP(&_personalizeListDatasetExportJobs, "list-dataset-export-jobs", "", false, "List Dataset Export Jobs")
	_personalizeCmd.Flags().BoolVarP(&_personalizeListDatasetGroups, "list-dataset-groups", "", false, "List Dataset Groups")
	_personalizeCmd.Flags().BoolVarP(&_personalizeListDatasetImportJobs, "list-dataset-import-jobs", "", false, "List Dataset Import Jobs")
	_personalizeCmd.Flags().BoolVarP(&_personalizeListDatasets, "list-datasets", "", false, "List Datasets")
	_personalizeCmd.Flags().BoolVarP(&_personalizeListEventTrackers, "list-event-trackers", "", false, "List Event Trackers")
	_personalizeCmd.Flags().BoolVarP(&_personalizeListFilters, "list-filters", "", false, "List Filters")
	_personalizeCmd.Flags().BoolVarP(&_personalizeListMetricAttributionMetrics, "list-metric-attribution-metrics", "", false, "List Metric Attribution Metrics")
	_personalizeCmd.Flags().BoolVarP(&_personalizeListMetricAttributions, "list-metric-attributions", "", false, "List Metric Attributions")
	_personalizeCmd.Flags().BoolVarP(&_personalizeListRecipes, "list-recipes", "", false, "List Recipes")
	_personalizeCmd.Flags().BoolVarP(&_personalizeListRecommenders, "list-recommenders", "", false, "List Recommenders")
	_personalizeCmd.Flags().BoolVarP(&_personalizeListSchemas, "list-schemas", "", false, "List Schemas")
	_personalizeCmd.Flags().BoolVarP(&_personalizeListSolutionVersions, "list-solution-versions", "", false, "List Solution Versions")
	_personalizeCmd.Flags().BoolVarP(&_personalizeListSolutions, "list-solutions", "", false, "List Solutions")
	_personalizeCmd.Flags().BoolVarP(&_personalizeListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_personalizeCmd.Flags().BoolVarP(&_personalizeStartRecommender, "start-recommender", "", false, "Start Recommender")
	_personalizeCmd.Flags().BoolVarP(&_personalizeStopRecommender, "stop-recommender", "", false, "Stop Recommender")
	_personalizeCmd.Flags().BoolVarP(&_personalizeStopSolutionVersionCreation, "stop-solution-version-creation", "", false, "Stop Solution Version Creation")
	_personalizeCmd.Flags().BoolVarP(&_personalizeTagResource, "tag-resource", "", false, "Tag Resource")
	_personalizeCmd.Flags().BoolVarP(&_personalizeUntagResource, "untag-resource", "", false, "Untag Resource")
	_personalizeCmd.Flags().BoolVarP(&_personalizeUpdateCampaign, "update-campaign", "", false, "Update Campaign")
	_personalizeCmd.Flags().BoolVarP(&_personalizeUpdateDataset, "update-dataset", "", false, "Update Dataset")
	_personalizeCmd.Flags().BoolVarP(&_personalizeUpdateMetricAttribution, "update-metric-attribution", "", false, "Update Metric Attribution")
	_personalizeCmd.Flags().BoolVarP(&_personalizeUpdateRecommender, "update-recommender", "", false, "Update Recommender")
	_personalizeCmd.Flags().BoolVarP(&_personalizeUpdateSolution, "update-solution", "", false, "Update Solution")

}
