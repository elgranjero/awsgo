package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/forecast"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// forecastCmd represents the forecast command
var _forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "AWS forecast CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := forecast.NewFromConfig(cfg)
		if _forecastCreateAutoPredictor {
			forecast_CreateAutoPredictor(cfg, client)
			return
		}
		if _forecastCreateDataset {
			forecast_CreateDataset(cfg, client)
			return
		}
		if _forecastCreateDatasetGroup {
			forecast_CreateDatasetGroup(cfg, client)
			return
		}
		if _forecastCreateDatasetImportJob {
			forecast_CreateDatasetImportJob(cfg, client)
			return
		}
		if _forecastCreateExplainability {
			forecast_CreateExplainability(cfg, client)
			return
		}
		if _forecastCreateExplainabilityExport {
			forecast_CreateExplainabilityExport(cfg, client)
			return
		}
		if _forecastCreateForecast {
			forecast_CreateForecast(cfg, client)
			return
		}
		if _forecastCreateForecastExportJob {
			forecast_CreateForecastExportJob(cfg, client)
			return
		}
		if _forecastCreateMonitor {
			forecast_CreateMonitor(cfg, client)
			return
		}
		if _forecastCreatePredictor {
			forecast_CreatePredictor(cfg, client)
			return
		}
		if _forecastCreatePredictorBacktestExportJob {
			forecast_CreatePredictorBacktestExportJob(cfg, client)
			return
		}
		if _forecastCreateWhatIfAnalysis {
			forecast_CreateWhatIfAnalysis(cfg, client)
			return
		}
		if _forecastCreateWhatIfForecast {
			forecast_CreateWhatIfForecast(cfg, client)
			return
		}
		if _forecastCreateWhatIfForecastExport {
			forecast_CreateWhatIfForecastExport(cfg, client)
			return
		}
		if _forecastDeleteDataset {
			forecast_DeleteDataset(cfg, client)
			return
		}
		if _forecastDeleteDatasetGroup {
			forecast_DeleteDatasetGroup(cfg, client)
			return
		}
		if _forecastDeleteDatasetImportJob {
			forecast_DeleteDatasetImportJob(cfg, client)
			return
		}
		if _forecastDeleteExplainability {
			forecast_DeleteExplainability(cfg, client)
			return
		}
		if _forecastDeleteExplainabilityExport {
			forecast_DeleteExplainabilityExport(cfg, client)
			return
		}
		if _forecastDeleteForecast {
			forecast_DeleteForecast(cfg, client)
			return
		}
		if _forecastDeleteForecastExportJob {
			forecast_DeleteForecastExportJob(cfg, client)
			return
		}
		if _forecastDeleteMonitor {
			forecast_DeleteMonitor(cfg, client)
			return
		}
		if _forecastDeletePredictor {
			forecast_DeletePredictor(cfg, client)
			return
		}
		if _forecastDeletePredictorBacktestExportJob {
			forecast_DeletePredictorBacktestExportJob(cfg, client)
			return
		}
		if _forecastDeleteResourceTree {
			forecast_DeleteResourceTree(cfg, client)
			return
		}
		if _forecastDeleteWhatIfAnalysis {
			forecast_DeleteWhatIfAnalysis(cfg, client)
			return
		}
		if _forecastDeleteWhatIfForecast {
			forecast_DeleteWhatIfForecast(cfg, client)
			return
		}
		if _forecastDeleteWhatIfForecastExport {
			forecast_DeleteWhatIfForecastExport(cfg, client)
			return
		}
		if _forecastDescribeAutoPredictor {
			forecast_DescribeAutoPredictor(cfg, client)
			return
		}
		if _forecastDescribeDataset {
			forecast_DescribeDataset(cfg, client)
			return
		}
		if _forecastDescribeDatasetGroup {
			forecast_DescribeDatasetGroup(cfg, client)
			return
		}
		if _forecastDescribeDatasetImportJob {
			forecast_DescribeDatasetImportJob(cfg, client)
			return
		}
		if _forecastDescribeExplainability {
			forecast_DescribeExplainability(cfg, client)
			return
		}
		if _forecastDescribeExplainabilityExport {
			forecast_DescribeExplainabilityExport(cfg, client)
			return
		}
		if _forecastDescribeForecast {
			forecast_DescribeForecast(cfg, client)
			return
		}
		if _forecastDescribeForecastExportJob {
			forecast_DescribeForecastExportJob(cfg, client)
			return
		}
		if _forecastDescribeMonitor {
			forecast_DescribeMonitor(cfg, client)
			return
		}
		if _forecastDescribePredictor {
			forecast_DescribePredictor(cfg, client)
			return
		}
		if _forecastDescribePredictorBacktestExportJob {
			forecast_DescribePredictorBacktestExportJob(cfg, client)
			return
		}
		if _forecastDescribeWhatIfAnalysis {
			forecast_DescribeWhatIfAnalysis(cfg, client)
			return
		}
		if _forecastDescribeWhatIfForecast {
			forecast_DescribeWhatIfForecast(cfg, client)
			return
		}
		if _forecastDescribeWhatIfForecastExport {
			forecast_DescribeWhatIfForecastExport(cfg, client)
			return
		}
		if _forecastGetAccuracyMetrics {
			forecast_GetAccuracyMetrics(cfg, client)
			return
		}
		if _forecastListDatasetGroups {
			forecast_ListDatasetGroups(cfg, client)
			return
		}
		if _forecastListDatasetImportJobs {
			forecast_ListDatasetImportJobs(cfg, client)
			return
		}
		if _forecastListDatasets {
			forecast_ListDatasets(cfg, client)
			return
		}
		if _forecastListExplainabilities {
			forecast_ListExplainabilities(cfg, client)
			return
		}
		if _forecastListExplainabilityExports {
			forecast_ListExplainabilityExports(cfg, client)
			return
		}
		if _forecastListForecastExportJobs {
			forecast_ListForecastExportJobs(cfg, client)
			return
		}
		if _forecastListForecasts {
			forecast_ListForecasts(cfg, client)
			return
		}
		if _forecastListMonitorEvaluations {
			forecast_ListMonitorEvaluations(cfg, client)
			return
		}
		if _forecastListMonitors {
			forecast_ListMonitors(cfg, client)
			return
		}
		if _forecastListPredictorBacktestExportJobs {
			forecast_ListPredictorBacktestExportJobs(cfg, client)
			return
		}
		if _forecastListPredictors {
			forecast_ListPredictors(cfg, client)
			return
		}
		if _forecastListTagsForResource {
			forecast_ListTagsForResource(cfg, client)
			return
		}
		if _forecastListWhatIfAnalyses {
			forecast_ListWhatIfAnalyses(cfg, client)
			return
		}
		if _forecastListWhatIfForecastExports {
			forecast_ListWhatIfForecastExports(cfg, client)
			return
		}
		if _forecastListWhatIfForecasts {
			forecast_ListWhatIfForecasts(cfg, client)
			return
		}
		if _forecastResumeResource {
			forecast_ResumeResource(cfg, client)
			return
		}
		if _forecastStopResource {
			forecast_StopResource(cfg, client)
			return
		}
		if _forecastTagResource {
			forecast_TagResource(cfg, client)
			return
		}
		if _forecastUntagResource {
			forecast_UntagResource(cfg, client)
			return
		}
		if _forecastUpdateDatasetGroup {
			forecast_UpdateDatasetGroup(cfg, client)
			return
		}

	},
}

var (
	_forecastCreateAutoPredictor                bool
	_forecastCreateDataset                      bool
	_forecastCreateDatasetGroup                 bool
	_forecastCreateDatasetImportJob             bool
	_forecastCreateExplainability               bool
	_forecastCreateExplainabilityExport         bool
	_forecastCreateForecast                     bool
	_forecastCreateForecastExportJob            bool
	_forecastCreateMonitor                      bool
	_forecastCreatePredictor                    bool
	_forecastCreatePredictorBacktestExportJob   bool
	_forecastCreateWhatIfAnalysis               bool
	_forecastCreateWhatIfForecast               bool
	_forecastCreateWhatIfForecastExport         bool
	_forecastDeleteDataset                      bool
	_forecastDeleteDatasetGroup                 bool
	_forecastDeleteDatasetImportJob             bool
	_forecastDeleteExplainability               bool
	_forecastDeleteExplainabilityExport         bool
	_forecastDeleteForecast                     bool
	_forecastDeleteForecastExportJob            bool
	_forecastDeleteMonitor                      bool
	_forecastDeletePredictor                    bool
	_forecastDeletePredictorBacktestExportJob   bool
	_forecastDeleteResourceTree                 bool
	_forecastDeleteWhatIfAnalysis               bool
	_forecastDeleteWhatIfForecast               bool
	_forecastDeleteWhatIfForecastExport         bool
	_forecastDescribeAutoPredictor              bool
	_forecastDescribeDataset                    bool
	_forecastDescribeDatasetGroup               bool
	_forecastDescribeDatasetImportJob           bool
	_forecastDescribeExplainability             bool
	_forecastDescribeExplainabilityExport       bool
	_forecastDescribeForecast                   bool
	_forecastDescribeForecastExportJob          bool
	_forecastDescribeMonitor                    bool
	_forecastDescribePredictor                  bool
	_forecastDescribePredictorBacktestExportJob bool
	_forecastDescribeWhatIfAnalysis             bool
	_forecastDescribeWhatIfForecast             bool
	_forecastDescribeWhatIfForecastExport       bool
	_forecastGetAccuracyMetrics                 bool
	_forecastListDatasetGroups                  bool
	_forecastListDatasetImportJobs              bool
	_forecastListDatasets                       bool
	_forecastListExplainabilities               bool
	_forecastListExplainabilityExports          bool
	_forecastListForecastExportJobs             bool
	_forecastListForecasts                      bool
	_forecastListMonitorEvaluations             bool
	_forecastListMonitors                       bool
	_forecastListPredictorBacktestExportJobs    bool
	_forecastListPredictors                     bool
	_forecastListTagsForResource                bool
	_forecastListWhatIfAnalyses                 bool
	_forecastListWhatIfForecastExports          bool
	_forecastListWhatIfForecasts                bool
	_forecastResumeResource                     bool
	_forecastStopResource                       bool
	_forecastTagResource                        bool
	_forecastUntagResource                      bool
	_forecastUpdateDatasetGroup                 bool

	_forecastAlgorithmArn                     string
	_forecastAutoMLOverrideStrategy           string
	_forecastDataConfig                       string
	_forecastDataFrequency                    string
	_forecastDataSource                       string
	_forecastDatasetArn                       string
	_forecastDatasetArns                      []string
	_forecastDatasetGroupArn                  string
	_forecastDatasetGroupName                 string
	_forecastDatasetImportJobArn              string
	_forecastDatasetImportJobName             string
	_forecastDatasetName                      string
	_forecastDatasetType                      string
	_forecastDestination                      string
	_forecastDomain                           string
	_forecastEnableVisualization              string
	_forecastEncryptionConfig                 string
	_forecastEndDateTime                      string
	_forecastEvaluationParameters             string
	_forecastExplainPredictor                 string
	_forecastExplainabilityArn                string
	_forecastExplainabilityConfig             string
	_forecastExplainabilityExportArn          string
	_forecastExplainabilityExportName         string
	_forecastExplainabilityName               string
	_forecastFeaturizationConfig              string
	_forecastFilters                          string
	_forecastForecastArn                      string
	_forecastForecastDimensions               []string
	_forecastForecastExportJobArn             string
	_forecastForecastExportJobName            string
	_forecastForecastFrequency                string
	_forecastForecastHorizon                  string
	_forecastForecastName                     string
	_forecastForecastTypes                    []string
	_forecastFormat                           string
	_forecastGeolocationFormat                string
	_forecastHPOConfig                        string
	_forecastImportMode                       string
	_forecastInputDataConfig                  string
	_forecastMaxResults                       string
	_forecastMonitorArn                       string
	_forecastMonitorConfig                    string
	_forecastMonitorName                      string
	_forecastNextToken                        string
	_forecastOptimizationMetric               string
	_forecastPerformAutoML                    string
	_forecastPerformHPO                       string
	_forecastPredictorArn                     string
	_forecastPredictorBacktestExportJobArn    string
	_forecastPredictorBacktestExportJobName   string
	_forecastPredictorName                    string
	_forecastReferencePredictorArn            string
	_forecastResourceArn                      string
	_forecastSchema                           string
	_forecastStartDateTime                    string
	_forecastTagKeys                          []string
	_forecastTags                             string
	_forecastTimeAlignmentBoundary            string
	_forecastTimeSeriesReplacementsDataSource string
	_forecastTimeSeriesSelector               string
	_forecastTimeSeriesTransformations        string
	_forecastTimeZone                         string
	_forecastTimestampFormat                  string
	_forecastTrainingParameters               string
	_forecastUseGeolocationForTimeZone        string
	_forecastWhatIfAnalysisArn                string
	_forecastWhatIfAnalysisName               string
	_forecastWhatIfForecastArn                string
	_forecastWhatIfForecastArns               []string
	_forecastWhatIfForecastExportArn          string
	_forecastWhatIfForecastExportName         string
	_forecastWhatIfForecastName               string
)

// Creates an Amazon Forecast predictor.
// Amazon Forecast creates predictors with AutoPredictor, which involves applying
// the optimal combination of algorithms to each time series in your datasets. You
// can use CreateAutoPredictorto create new predictors or upgrade/retrain existing predictors.
//
// # Creating new predictors
//
// The following parameters are required when creating a new predictor:
//
// - PredictorName - A unique name for the predictor.
//
// - DatasetGroupArn - The ARN of the dataset group used to train the predictor.
//
// - ForecastFrequency - The granularity of your forecasts (hourly, daily,
// weekly, etc).
//
// - ForecastHorizon - The number of time-steps that the model predicts. The
// forecast horizon is also called the prediction length.
//
// When creating a new predictor, do not specify a value for ReferencePredictorArn .
//
// # Upgrading and retraining predictors
//
// The following parameters are required when retraining or upgrading a predictor:
//
// - PredictorName - A unique name for the predictor.
//
// - ReferencePredictorArn - The ARN of the predictor to retrain or upgrade.
//
// When upgrading or retraining a predictor, only specify values for the
// ReferencePredictorArn and PredictorName .
func forecast_CreateAutoPredictor(cfg aws.Config, client *forecast.Client) {
	input := &forecast.CreateAutoPredictorInput{
		// PredictorName: *string, // Required
	}

	if len(_forecastPredictorName) > 0 {
		input.PredictorName = aws.String(_forecastPredictorName)
	}
	if len(_forecastDataConfig) > 0 {
		if err := assignInputField(input, "DataConfig", _forecastDataConfig); err != nil {
			log.Errorf("invalid --data-config: %s", err.Error())
			return
		}
	}
	if len(_forecastEncryptionConfig) > 0 {
		if err := assignInputField(input, "EncryptionConfig", _forecastEncryptionConfig); err != nil {
			log.Errorf("invalid --encryption-config: %s", err.Error())
			return
		}
	}
	if len(_forecastExplainPredictor) > 0 {
		if err := assignInputField(input, "ExplainPredictor", _forecastExplainPredictor); err != nil {
			log.Errorf("invalid --explain-predictor: %s", err.Error())
			return
		}
	}
	if len(_forecastForecastDimensions) > 0 {
		input.ForecastDimensions = append([]string(nil), _forecastForecastDimensions...)
	}
	if len(_forecastForecastFrequency) > 0 {
		input.ForecastFrequency = aws.String(_forecastForecastFrequency)
	}
	if len(_forecastForecastHorizon) > 0 {
		if err := assignInputField(input, "ForecastHorizon", _forecastForecastHorizon); err != nil {
			log.Errorf("invalid --forecast-horizon: %s", err.Error())
			return
		}
	}
	if len(_forecastForecastTypes) > 0 {
		input.ForecastTypes = append([]string(nil), _forecastForecastTypes...)
	}
	if len(_forecastMonitorConfig) > 0 {
		if err := assignInputField(input, "MonitorConfig", _forecastMonitorConfig); err != nil {
			log.Errorf("invalid --monitor-config: %s", err.Error())
			return
		}
	}
	if len(_forecastOptimizationMetric) > 0 {
		if err := assignInputField(input, "OptimizationMetric", _forecastOptimizationMetric); err != nil {
			log.Errorf("invalid --optimization-metric: %s", err.Error())
			return
		}
	}
	if len(_forecastReferencePredictorArn) > 0 {
		input.ReferencePredictorArn = aws.String(_forecastReferencePredictorArn)
	}
	if len(_forecastTags) > 0 {
		if err := assignInputField(input, "Tags", _forecastTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_forecastTimeAlignmentBoundary) > 0 {
		if err := assignInputField(input, "TimeAlignmentBoundary", _forecastTimeAlignmentBoundary); err != nil {
			log.Errorf("invalid --time-alignment-boundary: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAutoPredictor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Forecast dataset. The information about the dataset that you
// provide helps Forecast understand how to consume the data for model training.
// This includes the following:
//
// - DataFrequency - How frequently your historical time-series data is collected.
//
// - Domain and DatasetType - Each dataset has an associated dataset domain and a
// type within the domain. Amazon Forecast provides a list of predefined domains
// and types within each domain. For each unique dataset domain and type within the
// domain, Amazon Forecast requires your data to include a minimum set of
// predefined fields.
//
// - Schema - A schema specifies the fields in the dataset, including the field
// name and data type.
//
// After creating a dataset, you import your training data into it and add the
// dataset to a dataset group. You use the dataset group to create a predictor. For
// more information, see [Importing datasets].
//
// To get a list of all your datasets, use the [ListDatasets] operation.
//
// For example Forecast datasets, see the [Amazon Forecast Sample GitHub repository].
//
// The Status of a dataset must be ACTIVE before you can import training data. Use
// the [DescribeDataset]operation to get the status.
//
// [Amazon Forecast Sample GitHub repository]: https://github.com/aws-samples/amazon-forecast-samples
// [DescribeDataset]: https://docs.aws.amazon.com/forecast/latest/dg/API_DescribeDataset.html
// [ListDatasets]: https://docs.aws.amazon.com/forecast/latest/dg/API_ListDatasets.html
// [Importing datasets]: https://docs.aws.amazon.com/forecast/latest/dg/howitworks-datasets-groups.html
func forecast_CreateDataset(cfg aws.Config, client *forecast.Client) {
	input := &forecast.CreateDatasetInput{
		// DatasetName: *string, // Required
		// DatasetType: types.DatasetType, // Required
		// Domain: types.Domain, // Required
		// Schema: *types.Schema, // Required
	}

	if len(_forecastDatasetName) > 0 {
		input.DatasetName = aws.String(_forecastDatasetName)
	}
	if len(_forecastDatasetType) > 0 {
		if err := assignInputField(input, "DatasetType", _forecastDatasetType); err != nil {
			log.Errorf("invalid --dataset-type: %s", err.Error())
			return
		}
	}
	if len(_forecastDomain) > 0 {
		if err := assignInputField(input, "Domain", _forecastDomain); err != nil {
			log.Errorf("invalid --domain: %s", err.Error())
			return
		}
	}
	if len(_forecastSchema) > 0 {
		if err := assignInputField(input, "Schema", _forecastSchema); err != nil {
			log.Errorf("invalid --schema: %s", err.Error())
			return
		}
	}
	if len(_forecastDataFrequency) > 0 {
		input.DataFrequency = aws.String(_forecastDataFrequency)
	}
	if len(_forecastEncryptionConfig) > 0 {
		if err := assignInputField(input, "EncryptionConfig", _forecastEncryptionConfig); err != nil {
			log.Errorf("invalid --encryption-config: %s", err.Error())
			return
		}
	}
	if len(_forecastTags) > 0 {
		if err := assignInputField(input, "Tags", _forecastTags); err != nil {
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

// Creates a dataset group, which holds a collection of related datasets. You can
// add datasets to the dataset group when you create the dataset group, or later by
// using the [UpdateDatasetGroup]operation.
//
// After creating a dataset group and adding datasets, you use the dataset group
// when you create a predictor. For more information, see [Dataset groups].
//
// To get a list of all your datasets groups, use the [ListDatasetGroups] operation.
//
// The Status of a dataset group must be ACTIVE before you can use the dataset
// group to create a predictor. To get the status, use the [DescribeDatasetGroup]operation.
//
// [ListDatasetGroups]: https://docs.aws.amazon.com/forecast/latest/dg/API_ListDatasetGroups.html
// [DescribeDatasetGroup]: https://docs.aws.amazon.com/forecast/latest/dg/API_DescribeDatasetGroup.html
// [UpdateDatasetGroup]: https://docs.aws.amazon.com/forecast/latest/dg/API_UpdateDatasetGroup.html
// [Dataset groups]: https://docs.aws.amazon.com/forecast/latest/dg/howitworks-datasets-groups.html
func forecast_CreateDatasetGroup(cfg aws.Config, client *forecast.Client) {
	input := &forecast.CreateDatasetGroupInput{
		// DatasetGroupName: *string, // Required
		// Domain: types.Domain, // Required
	}

	if len(_forecastDatasetGroupName) > 0 {
		input.DatasetGroupName = aws.String(_forecastDatasetGroupName)
	}
	if len(_forecastDomain) > 0 {
		if err := assignInputField(input, "Domain", _forecastDomain); err != nil {
			log.Errorf("invalid --domain: %s", err.Error())
			return
		}
	}
	if len(_forecastDatasetArns) > 0 {
		input.DatasetArns = append([]string(nil), _forecastDatasetArns...)
	}
	if len(_forecastTags) > 0 {
		if err := assignInputField(input, "Tags", _forecastTags); err != nil {
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

// Imports your training data to an Amazon Forecast dataset. You provide the
// location of your training data in an Amazon Simple Storage Service (Amazon S3)
// bucket and the Amazon Resource Name (ARN) of the dataset that you want to import
// the data to.
//
// You must specify a [DataSource] object that includes an Identity and Access Management
// (IAM) role that Amazon Forecast can assume to access the data, as Amazon
// Forecast makes a copy of your data and processes it in an internal Amazon Web
// Services system. For more information, see [Set up permissions].
//
// The training data must be in CSV or Parquet format. The delimiter must be a
// comma (,).
//
// You can specify the path to a specific file, the S3 bucket, or to a folder in
// the S3 bucket. For the latter two cases, Amazon Forecast imports all files up to
// the limit of 10,000 files.
//
// Because dataset imports are not aggregated, your most recent dataset import is
// the one that is used when training a predictor or generating a forecast. Make
// sure that your most recent dataset import contains all of the data you want to
// model off of, and not just the new data collected since the previous import.
//
// To get a list of all your dataset import jobs, filtered by specified criteria,
// use the [ListDatasetImportJobs]operation.
//
// [ListDatasetImportJobs]: https://docs.aws.amazon.com/forecast/latest/dg/API_ListDatasetImportJobs.html
// [Set up permissions]: https://docs.aws.amazon.com/forecast/latest/dg/aws-forecast-iam-roles.html
// [DataSource]: https://docs.aws.amazon.com/forecast/latest/dg/API_DataSource.html
func forecast_CreateDatasetImportJob(cfg aws.Config, client *forecast.Client) {
	input := &forecast.CreateDatasetImportJobInput{
		// DataSource: *types.DataSource, // Required
		// DatasetArn: *string, // Required
		// DatasetImportJobName: *string, // Required
	}

	if len(_forecastDataSource) > 0 {
		if err := assignInputField(input, "DataSource", _forecastDataSource); err != nil {
			log.Errorf("invalid --data-source: %s", err.Error())
			return
		}
	}
	if len(_forecastDatasetArn) > 0 {
		input.DatasetArn = aws.String(_forecastDatasetArn)
	}
	if len(_forecastDatasetImportJobName) > 0 {
		input.DatasetImportJobName = aws.String(_forecastDatasetImportJobName)
	}
	if len(_forecastFormat) > 0 {
		input.Format = aws.String(_forecastFormat)
	}
	if len(_forecastGeolocationFormat) > 0 {
		input.GeolocationFormat = aws.String(_forecastGeolocationFormat)
	}
	if len(_forecastImportMode) > 0 {
		if err := assignInputField(input, "ImportMode", _forecastImportMode); err != nil {
			log.Errorf("invalid --import-mode: %s", err.Error())
			return
		}
	}
	if len(_forecastTags) > 0 {
		if err := assignInputField(input, "Tags", _forecastTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_forecastTimeZone) > 0 {
		input.TimeZone = aws.String(_forecastTimeZone)
	}
	if len(_forecastTimestampFormat) > 0 {
		input.TimestampFormat = aws.String(_forecastTimestampFormat)
	}
	if len(_forecastUseGeolocationForTimeZone) > 0 {
		if err := assignInputField(input, "UseGeolocationForTimeZone", _forecastUseGeolocationForTimeZone); err != nil {
			log.Errorf("invalid --use-geolocation-for-time-zone: %s", err.Error())
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

// Explainability is only available for Forecasts and Predictors generated from an
// AutoPredictor (CreateAutoPredictor )
//
// Creates an Amazon Forecast Explainability.
//
// Explainability helps you better understand how the attributes in your datasets
// impact forecast. Amazon Forecast uses a metric called Impact scores to quantify
// the relative impact of each attribute and determine whether they increase or
// decrease forecast values.
//
// To enable Forecast Explainability, your predictor must include at least one of
// the following: related time series, item metadata, or additional datasets like
// Holidays and the Weather Index.
//
// CreateExplainability accepts either a Predictor ARN or Forecast ARN. To receive
// aggregated Impact scores for all time series and time points in your datasets,
// provide a Predictor ARN. To receive Impact scores for specific time series and
// time points, provide a Forecast ARN.
//
// # CreateExplainability with a Predictor ARN
//
// You can only have one Explainability resource per predictor. If you already
// enabled ExplainPredictor in CreateAutoPredictor, that predictor already has an Explainability
// resource.
//
// The following parameters are required when providing a Predictor ARN:
//
// - ExplainabilityName - A unique name for the Explainability.
//
// - ResourceArn - The Arn of the predictor.
//
// - TimePointGranularity - Must be set to “ALL”.
//
// - TimeSeriesGranularity - Must be set to “ALL”.
//
// Do not specify a value for the following parameters:
//
// - DataSource - Only valid when TimeSeriesGranularity is “SPECIFIC”.
//
// - Schema - Only valid when TimeSeriesGranularity is “SPECIFIC”.
//
// - StartDateTime - Only valid when TimePointGranularity is “SPECIFIC”.
//
// - EndDateTime - Only valid when TimePointGranularity is “SPECIFIC”.
//
// # CreateExplainability with a Forecast ARN
//
// You can specify a maximum of 50 time series and 500 time points.
//
// The following parameters are required when providing a Predictor ARN:
//
// - ExplainabilityName - A unique name for the Explainability.
//
// - ResourceArn - The Arn of the forecast.
//
// - TimePointGranularity - Either “ALL” or “SPECIFIC”.
//
// - TimeSeriesGranularity - Either “ALL” or “SPECIFIC”.
//
// If you set TimeSeriesGranularity to “SPECIFIC”, you must also provide the
// following:
//
// - DataSource - The S3 location of the CSV file specifying your time series.
//
// - Schema - The Schema defines the attributes and attribute types listed in the
// Data Source.
//
// If you set TimePointGranularity to “SPECIFIC”, you must also provide the
// following:
//
// - StartDateTime - The first timestamp in the range of time points.
//
// - EndDateTime - The last timestamp in the range of time points.
func forecast_CreateExplainability(cfg aws.Config, client *forecast.Client) {
	input := &forecast.CreateExplainabilityInput{
		// ExplainabilityConfig: *types.ExplainabilityConfig, // Required
		// ExplainabilityName: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_forecastExplainabilityConfig) > 0 {
		if err := assignInputField(input, "ExplainabilityConfig", _forecastExplainabilityConfig); err != nil {
			log.Errorf("invalid --explainability-config: %s", err.Error())
			return
		}
	}
	if len(_forecastExplainabilityName) > 0 {
		input.ExplainabilityName = aws.String(_forecastExplainabilityName)
	}
	if len(_forecastResourceArn) > 0 {
		input.ResourceArn = aws.String(_forecastResourceArn)
	}
	if len(_forecastDataSource) > 0 {
		if err := assignInputField(input, "DataSource", _forecastDataSource); err != nil {
			log.Errorf("invalid --data-source: %s", err.Error())
			return
		}
	}
	if len(_forecastEnableVisualization) > 0 {
		if err := assignInputField(input, "EnableVisualization", _forecastEnableVisualization); err != nil {
			log.Errorf("invalid --enable-visualization: %s", err.Error())
			return
		}
	}
	if len(_forecastEndDateTime) > 0 {
		input.EndDateTime = aws.String(_forecastEndDateTime)
	}
	if len(_forecastSchema) > 0 {
		if err := assignInputField(input, "Schema", _forecastSchema); err != nil {
			log.Errorf("invalid --schema: %s", err.Error())
			return
		}
	}
	if len(_forecastStartDateTime) > 0 {
		input.StartDateTime = aws.String(_forecastStartDateTime)
	}
	if len(_forecastTags) > 0 {
		if err := assignInputField(input, "Tags", _forecastTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateExplainability(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Exports an Explainability resource created by the CreateExplainability operation. Exported files
// are exported to an Amazon Simple Storage Service (Amazon S3) bucket.
//
// You must specify a DataDestination object that includes an Amazon S3 bucket and an Identity
// and Access Management (IAM) role that Amazon Forecast can assume to access the
// Amazon S3 bucket. For more information, see aws-forecast-iam-roles.
//
// The Status of the export job must be ACTIVE before you can access the export in
// your Amazon S3 bucket. To get the status, use the DescribeExplainabilityExportoperation.
func forecast_CreateExplainabilityExport(cfg aws.Config, client *forecast.Client) {
	input := &forecast.CreateExplainabilityExportInput{
		// Destination: *types.DataDestination, // Required
		// ExplainabilityArn: *string, // Required
		// ExplainabilityExportName: *string, // Required
	}

	if len(_forecastDestination) > 0 {
		if err := assignInputField(input, "Destination", _forecastDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_forecastExplainabilityArn) > 0 {
		input.ExplainabilityArn = aws.String(_forecastExplainabilityArn)
	}
	if len(_forecastExplainabilityExportName) > 0 {
		input.ExplainabilityExportName = aws.String(_forecastExplainabilityExportName)
	}
	if len(_forecastFormat) > 0 {
		input.Format = aws.String(_forecastFormat)
	}
	if len(_forecastTags) > 0 {
		if err := assignInputField(input, "Tags", _forecastTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateExplainabilityExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a forecast for each item in the TARGET_TIME_SERIES dataset that was
// used to train the predictor. This is known as inference. To retrieve the
// forecast for a single item at low latency, use the operation. To export the
// complete forecast into your Amazon Simple Storage Service (Amazon S3) bucket,
// use the CreateForecastExportJoboperation.
//
// The range of the forecast is determined by the ForecastHorizon value, which you
// specify in the CreatePredictorrequest. When you query a forecast, you can request a specific
// date range within the forecast.
//
// To get a list of all your forecasts, use the ListForecasts operation.
//
// The forecasts generated by Amazon Forecast are in the same time zone as the
// dataset that was used to create the predictor.
//
// For more information, see howitworks-forecast.
//
// The Status of the forecast must be ACTIVE before you can query or export the
// forecast. Use the DescribeForecastoperation to get the status.
//
// By default, a forecast includes predictions for every item ( item_id ) in the
// dataset group that was used to train the predictor. However, you can use the
// TimeSeriesSelector object to generate a forecast on a subset of time series.
// Forecast creation is skipped for any time series that you specify that are not
// in the input dataset. The forecast export file will not contain these time
// series or their forecasted values.
func forecast_CreateForecast(cfg aws.Config, client *forecast.Client) {
	input := &forecast.CreateForecastInput{
		// ForecastName: *string, // Required
		// PredictorArn: *string, // Required
	}

	if len(_forecastForecastName) > 0 {
		input.ForecastName = aws.String(_forecastForecastName)
	}
	if len(_forecastPredictorArn) > 0 {
		input.PredictorArn = aws.String(_forecastPredictorArn)
	}
	if len(_forecastForecastTypes) > 0 {
		input.ForecastTypes = append([]string(nil), _forecastForecastTypes...)
	}
	if len(_forecastTags) > 0 {
		if err := assignInputField(input, "Tags", _forecastTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_forecastTimeSeriesSelector) > 0 {
		if err := assignInputField(input, "TimeSeriesSelector", _forecastTimeSeriesSelector); err != nil {
			log.Errorf("invalid --time-series-selector: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateForecast(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Exports a forecast created by the CreateForecast operation to your Amazon Simple Storage
// Service (Amazon S3) bucket. The forecast file name will match the following
// conventions:
//
// __
//
// where the component is in Java SimpleDateFormat (yyyy-MM-ddTHH-mm-ssZ).
//
// You must specify a DataDestination object that includes an Identity and Access Management
// (IAM) role that Amazon Forecast can assume to access the Amazon S3 bucket. For
// more information, see aws-forecast-iam-roles.
//
// For more information, see howitworks-forecast.
//
// To get a list of all your forecast export jobs, use the ListForecastExportJobs operation.
//
// The Status of the forecast export job must be ACTIVE before you can access the
// forecast in your Amazon S3 bucket. To get the status, use the DescribeForecastExportJoboperation.
func forecast_CreateForecastExportJob(cfg aws.Config, client *forecast.Client) {
	input := &forecast.CreateForecastExportJobInput{
		// Destination: *types.DataDestination, // Required
		// ForecastArn: *string, // Required
		// ForecastExportJobName: *string, // Required
	}

	if len(_forecastDestination) > 0 {
		if err := assignInputField(input, "Destination", _forecastDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_forecastForecastArn) > 0 {
		input.ForecastArn = aws.String(_forecastForecastArn)
	}
	if len(_forecastForecastExportJobName) > 0 {
		input.ForecastExportJobName = aws.String(_forecastForecastExportJobName)
	}
	if len(_forecastFormat) > 0 {
		input.Format = aws.String(_forecastFormat)
	}
	if len(_forecastTags) > 0 {
		if err := assignInputField(input, "Tags", _forecastTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateForecastExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a predictor monitor resource for an existing auto predictor. Predictor
// monitoring allows you to see how your predictor's performance changes over time.
// For more information, see [Predictor Monitoring].
//
// [Predictor Monitoring]: https://docs.aws.amazon.com/forecast/latest/dg/predictor-monitoring.html
func forecast_CreateMonitor(cfg aws.Config, client *forecast.Client) {
	input := &forecast.CreateMonitorInput{
		// MonitorName: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_forecastMonitorName) > 0 {
		input.MonitorName = aws.String(_forecastMonitorName)
	}
	if len(_forecastResourceArn) > 0 {
		input.ResourceArn = aws.String(_forecastResourceArn)
	}
	if len(_forecastTags) > 0 {
		if err := assignInputField(input, "Tags", _forecastTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMonitor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation creates a legacy predictor that does not include all the
// predictor functionalities provided by Amazon Forecast. To create a predictor
// that is compatible with all aspects of Forecast, use CreateAutoPredictor.
//
// Creates an Amazon Forecast predictor.
//
// In the request, provide a dataset group and either specify an algorithm or let
// Amazon Forecast choose an algorithm for you using AutoML. If you specify an
// algorithm, you also can override algorithm-specific hyperparameters.
//
// Amazon Forecast uses the algorithm to train a predictor using the latest
// version of the datasets in the specified dataset group. You can then generate a
// forecast using the CreateForecastoperation.
//
// To see the evaluation metrics, use the GetAccuracyMetrics operation.
//
// You can specify a featurization configuration to fill and aggregate the data
// fields in the TARGET_TIME_SERIES dataset to improve model training. For more
// information, see FeaturizationConfig.
//
// For RELATED_TIME_SERIES datasets, CreatePredictor verifies that the
// DataFrequency specified when the dataset was created matches the
// ForecastFrequency . TARGET_TIME_SERIES datasets don't have this restriction.
// Amazon Forecast also verifies the delimiter and timestamp format. For more
// information, see howitworks-datasets-groups.
//
// By default, predictors are trained and evaluated at the 0.1 (P10), 0.5 (P50),
// and 0.9 (P90) quantiles. You can choose custom forecast types to train and
// evaluate your predictor by setting the ForecastTypes .
//
// # AutoML
//
// If you want Amazon Forecast to evaluate each algorithm and choose the one that
// minimizes the objective function , set PerformAutoML to true . The objective
// function is defined as the mean of the weighted losses over the forecast types.
// By default, these are the p10, p50, and p90 quantile losses. For more
// information, see EvaluationResult.
//
// When AutoML is enabled, the following properties are disallowed:
//
// - AlgorithmArn
//
// - HPOConfig
//
// - PerformHPO
//
// - TrainingParameters
//
// To get a list of all of your predictors, use the ListPredictors operation.
//
// Before you can use the predictor to create a forecast, the Status of the
// predictor must be ACTIVE , signifying that training has completed. To get the
// status, use the DescribePredictoroperation.
func forecast_CreatePredictor(cfg aws.Config, client *forecast.Client) {
	input := &forecast.CreatePredictorInput{
		// FeaturizationConfig: *types.FeaturizationConfig, // Required
		// ForecastHorizon: *int32, // Required
		// InputDataConfig: *types.InputDataConfig, // Required
		// PredictorName: *string, // Required
	}

	if len(_forecastFeaturizationConfig) > 0 {
		if err := assignInputField(input, "FeaturizationConfig", _forecastFeaturizationConfig); err != nil {
			log.Errorf("invalid --featurization-config: %s", err.Error())
			return
		}
	}
	if len(_forecastForecastHorizon) > 0 {
		if err := assignInputField(input, "ForecastHorizon", _forecastForecastHorizon); err != nil {
			log.Errorf("invalid --forecast-horizon: %s", err.Error())
			return
		}
	}
	if len(_forecastInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _forecastInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_forecastPredictorName) > 0 {
		input.PredictorName = aws.String(_forecastPredictorName)
	}
	if len(_forecastAlgorithmArn) > 0 {
		input.AlgorithmArn = aws.String(_forecastAlgorithmArn)
	}
	if len(_forecastAutoMLOverrideStrategy) > 0 {
		if err := assignInputField(input, "AutoMLOverrideStrategy", _forecastAutoMLOverrideStrategy); err != nil {
			log.Errorf("invalid --auto-ml-override-strategy: %s", err.Error())
			return
		}
	}
	if len(_forecastEncryptionConfig) > 0 {
		if err := assignInputField(input, "EncryptionConfig", _forecastEncryptionConfig); err != nil {
			log.Errorf("invalid --encryption-config: %s", err.Error())
			return
		}
	}
	if len(_forecastEvaluationParameters) > 0 {
		if err := assignInputField(input, "EvaluationParameters", _forecastEvaluationParameters); err != nil {
			log.Errorf("invalid --evaluation-parameters: %s", err.Error())
			return
		}
	}
	if len(_forecastForecastTypes) > 0 {
		input.ForecastTypes = append([]string(nil), _forecastForecastTypes...)
	}
	if len(_forecastHPOConfig) > 0 {
		if err := assignInputField(input, "HPOConfig", _forecastHPOConfig); err != nil {
			log.Errorf("invalid --hpo-config: %s", err.Error())
			return
		}
	}
	if len(_forecastOptimizationMetric) > 0 {
		if err := assignInputField(input, "OptimizationMetric", _forecastOptimizationMetric); err != nil {
			log.Errorf("invalid --optimization-metric: %s", err.Error())
			return
		}
	}
	if len(_forecastPerformAutoML) > 0 {
		if err := assignInputField(input, "PerformAutoML", _forecastPerformAutoML); err != nil {
			log.Errorf("invalid --perform-auto-ml: %s", err.Error())
			return
		}
	}
	if len(_forecastPerformHPO) > 0 {
		if err := assignInputField(input, "PerformHPO", _forecastPerformHPO); err != nil {
			log.Errorf("invalid --perform-hpo: %s", err.Error())
			return
		}
	}
	if len(_forecastTags) > 0 {
		if err := assignInputField(input, "Tags", _forecastTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_forecastTrainingParameters) > 0 {
		if err := assignInputField(input, "TrainingParameters", _forecastTrainingParameters); err != nil {
			log.Errorf("invalid --training-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePredictor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Exports backtest forecasts and accuracy metrics generated by the CreateAutoPredictor or CreatePredictor
// operations. Two folders containing CSV or Parquet files are exported to your
// specified S3 bucket.
//
// The export file names will match the following conventions:
//
// __.csv
//
// The component is in Java SimpleDate format (yyyy-MM-ddTHH-mm-ssZ).
//
// You must specify a DataDestination object that includes an Amazon S3 bucket and an Identity
// and Access Management (IAM) role that Amazon Forecast can assume to access the
// Amazon S3 bucket. For more information, see aws-forecast-iam-roles.
//
// The Status of the export job must be ACTIVE before you can access the export in
// your Amazon S3 bucket. To get the status, use the DescribePredictorBacktestExportJoboperation.
func forecast_CreatePredictorBacktestExportJob(cfg aws.Config, client *forecast.Client) {
	input := &forecast.CreatePredictorBacktestExportJobInput{
		// Destination: *types.DataDestination, // Required
		// PredictorArn: *string, // Required
		// PredictorBacktestExportJobName: *string, // Required
	}

	if len(_forecastDestination) > 0 {
		if err := assignInputField(input, "Destination", _forecastDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_forecastPredictorArn) > 0 {
		input.PredictorArn = aws.String(_forecastPredictorArn)
	}
	if len(_forecastPredictorBacktestExportJobName) > 0 {
		input.PredictorBacktestExportJobName = aws.String(_forecastPredictorBacktestExportJobName)
	}
	if len(_forecastFormat) > 0 {
		input.Format = aws.String(_forecastFormat)
	}
	if len(_forecastTags) > 0 {
		if err := assignInputField(input, "Tags", _forecastTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePredictorBacktestExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// What-if analysis is a scenario modeling technique where you make a hypothetical
// change to a time series and compare the forecasts generated by these changes
// against the baseline, unchanged time series. It is important to remember that
// the purpose of a what-if analysis is to understand how a forecast can change
// given different modifications to the baseline time series.
//
// For example, imagine you are a clothing retailer who is considering an end of
// season sale to clear space for new styles. After creating a baseline forecast,
// you can use a what-if analysis to investigate how different sales tactics might
// affect your goals.
//
// You could create a scenario where everything is given a 25% markdown, and
// another where everything is given a fixed dollar markdown. You could create a
// scenario where the sale lasts for one week and another where the sale lasts for
// one month. With a what-if analysis, you can compare many different scenarios
// against each other.
//
// Note that a what-if analysis is meant to display what the forecasting model has
// learned and how it will behave in the scenarios that you are evaluating. Do not
// blindly use the results of the what-if analysis to make business decisions. For
// instance, forecasts might not be accurate for novel scenarios where there is no
// reference available to determine whether a forecast is good.
//
// The TimeSeriesSelector object defines the items that you want in the what-if analysis.
func forecast_CreateWhatIfAnalysis(cfg aws.Config, client *forecast.Client) {
	input := &forecast.CreateWhatIfAnalysisInput{
		// ForecastArn: *string, // Required
		// WhatIfAnalysisName: *string, // Required
	}

	if len(_forecastForecastArn) > 0 {
		input.ForecastArn = aws.String(_forecastForecastArn)
	}
	if len(_forecastWhatIfAnalysisName) > 0 {
		input.WhatIfAnalysisName = aws.String(_forecastWhatIfAnalysisName)
	}
	if len(_forecastTags) > 0 {
		if err := assignInputField(input, "Tags", _forecastTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_forecastTimeSeriesSelector) > 0 {
		if err := assignInputField(input, "TimeSeriesSelector", _forecastTimeSeriesSelector); err != nil {
			log.Errorf("invalid --time-series-selector: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWhatIfAnalysis(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A what-if forecast is a forecast that is created from a modified version of the
// baseline forecast. Each what-if forecast incorporates either a replacement
// dataset or a set of transformations to the original dataset.
func forecast_CreateWhatIfForecast(cfg aws.Config, client *forecast.Client) {
	input := &forecast.CreateWhatIfForecastInput{
		// WhatIfAnalysisArn: *string, // Required
		// WhatIfForecastName: *string, // Required
	}

	if len(_forecastWhatIfAnalysisArn) > 0 {
		input.WhatIfAnalysisArn = aws.String(_forecastWhatIfAnalysisArn)
	}
	if len(_forecastWhatIfForecastName) > 0 {
		input.WhatIfForecastName = aws.String(_forecastWhatIfForecastName)
	}
	if len(_forecastTags) > 0 {
		if err := assignInputField(input, "Tags", _forecastTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_forecastTimeSeriesReplacementsDataSource) > 0 {
		if err := assignInputField(input, "TimeSeriesReplacementsDataSource", _forecastTimeSeriesReplacementsDataSource); err != nil {
			log.Errorf("invalid --time-series-replacements-data-source: %s", err.Error())
			return
		}
	}
	if len(_forecastTimeSeriesTransformations) > 0 {
		if err := assignInputField(input, "TimeSeriesTransformations", _forecastTimeSeriesTransformations); err != nil {
			log.Errorf("invalid --time-series-transformations: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWhatIfForecast(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Exports a forecast created by the CreateWhatIfForecast operation to your Amazon Simple Storage
// Service (Amazon S3) bucket. The forecast file name will match the following
// conventions:
//
// ≈__
//
// The component is in Java SimpleDateFormat (yyyy-MM-ddTHH-mm-ssZ).
//
// You must specify a DataDestination object that includes an Identity and Access Management
// (IAM) role that Amazon Forecast can assume to access the Amazon S3 bucket. For
// more information, see aws-forecast-iam-roles.
//
// For more information, see howitworks-forecast.
//
// To get a list of all your what-if forecast export jobs, use the ListWhatIfForecastExports operation.
//
// The Status of the forecast export job must be ACTIVE before you can access the
// forecast in your Amazon S3 bucket. To get the status, use the DescribeWhatIfForecastExportoperation.
func forecast_CreateWhatIfForecastExport(cfg aws.Config, client *forecast.Client) {
	input := &forecast.CreateWhatIfForecastExportInput{
		// Destination: *types.DataDestination, // Required
		// WhatIfForecastArns: []string, // Required
		// WhatIfForecastExportName: *string, // Required
	}

	if len(_forecastDestination) > 0 {
		if err := assignInputField(input, "Destination", _forecastDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_forecastWhatIfForecastArns) > 0 {
		input.WhatIfForecastArns = append([]string(nil), _forecastWhatIfForecastArns...)
	}
	if len(_forecastWhatIfForecastExportName) > 0 {
		input.WhatIfForecastExportName = aws.String(_forecastWhatIfForecastExportName)
	}
	if len(_forecastFormat) > 0 {
		input.Format = aws.String(_forecastFormat)
	}
	if len(_forecastTags) > 0 {
		if err := assignInputField(input, "Tags", _forecastTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWhatIfForecastExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Forecast dataset that was created using the [CreateDataset] operation. You
// can only delete datasets that have a status of ACTIVE or CREATE_FAILED . To get
// the status use the [DescribeDataset]operation.
//
// Forecast does not automatically update any dataset groups that contain the
// deleted dataset. In order to update the dataset group, use the [UpdateDatasetGroup]operation,
// omitting the deleted dataset's ARN.
//
// [UpdateDatasetGroup]: https://docs.aws.amazon.com/forecast/latest/dg/API_UpdateDatasetGroup.html
// [DescribeDataset]: https://docs.aws.amazon.com/forecast/latest/dg/API_DescribeDataset.html
// [CreateDataset]: https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDataset.html
func forecast_DeleteDataset(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DeleteDatasetInput{
		// DatasetArn: *string, // Required
	}

	if len(_forecastDatasetArn) > 0 {
		input.DatasetArn = aws.String(_forecastDatasetArn)
	}

	if resp, err := client.DeleteDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a dataset group created using the [CreateDatasetGroup] operation. You can only delete
// dataset groups that have a status of ACTIVE , CREATE_FAILED , or UPDATE_FAILED .
// To get the status, use the [DescribeDatasetGroup]operation.
//
// This operation deletes only the dataset group, not the datasets in the group.
//
// [DescribeDatasetGroup]: https://docs.aws.amazon.com/forecast/latest/dg/API_DescribeDatasetGroup.html
// [CreateDatasetGroup]: https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDatasetGroup.html
func forecast_DeleteDatasetGroup(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DeleteDatasetGroupInput{
		// DatasetGroupArn: *string, // Required
	}

	if len(_forecastDatasetGroupArn) > 0 {
		input.DatasetGroupArn = aws.String(_forecastDatasetGroupArn)
	}

	if resp, err := client.DeleteDatasetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a dataset import job created using the [CreateDatasetImportJob] operation. You can delete only
// dataset import jobs that have a status of ACTIVE or CREATE_FAILED . To get the
// status, use the [DescribeDatasetImportJob]operation.
//
// [DescribeDatasetImportJob]: https://docs.aws.amazon.com/forecast/latest/dg/API_DescribeDatasetImportJob.html
// [CreateDatasetImportJob]: https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDatasetImportJob.html
func forecast_DeleteDatasetImportJob(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DeleteDatasetImportJobInput{
		// DatasetImportJobArn: *string, // Required
	}

	if len(_forecastDatasetImportJobArn) > 0 {
		input.DatasetImportJobArn = aws.String(_forecastDatasetImportJobArn)
	}

	if resp, err := client.DeleteDatasetImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Explainability resource.
// You can delete only predictor that have a status of ACTIVE or CREATE_FAILED . To
// get the status, use the DescribeExplainabilityoperation.
func forecast_DeleteExplainability(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DeleteExplainabilityInput{
		// ExplainabilityArn: *string, // Required
	}

	if len(_forecastExplainabilityArn) > 0 {
		input.ExplainabilityArn = aws.String(_forecastExplainabilityArn)
	}

	if resp, err := client.DeleteExplainability(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Explainability export.
func forecast_DeleteExplainabilityExport(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DeleteExplainabilityExportInput{
		// ExplainabilityExportArn: *string, // Required
	}

	if len(_forecastExplainabilityExportArn) > 0 {
		input.ExplainabilityExportArn = aws.String(_forecastExplainabilityExportArn)
	}

	if resp, err := client.DeleteExplainabilityExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a forecast created using the CreateForecast operation. You can delete only forecasts
// that have a status of ACTIVE or CREATE_FAILED . To get the status, use the DescribeForecast
// operation.
//
// You can't delete a forecast while it is being exported. After a forecast is
// deleted, you can no longer query the forecast.
func forecast_DeleteForecast(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DeleteForecastInput{
		// ForecastArn: *string, // Required
	}

	if len(_forecastForecastArn) > 0 {
		input.ForecastArn = aws.String(_forecastForecastArn)
	}

	if resp, err := client.DeleteForecast(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a forecast export job created using the CreateForecastExportJob operation. You can delete only
// export jobs that have a status of ACTIVE or CREATE_FAILED . To get the status,
// use the DescribeForecastExportJoboperation.
func forecast_DeleteForecastExportJob(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DeleteForecastExportJobInput{
		// ForecastExportJobArn: *string, // Required
	}

	if len(_forecastForecastExportJobArn) > 0 {
		input.ForecastExportJobArn = aws.String(_forecastForecastExportJobArn)
	}

	if resp, err := client.DeleteForecastExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a monitor resource. You can only delete a monitor resource with a
// status of ACTIVE , ACTIVE_STOPPED , CREATE_FAILED , or CREATE_STOPPED .
func forecast_DeleteMonitor(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DeleteMonitorInput{
		// MonitorArn: *string, // Required
	}

	if len(_forecastMonitorArn) > 0 {
		input.MonitorArn = aws.String(_forecastMonitorArn)
	}

	if resp, err := client.DeleteMonitor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a predictor created using the DescribePredictor or CreatePredictor operations. You can delete only
// predictor that have a status of ACTIVE or CREATE_FAILED . To get the status, use
// the DescribePredictoroperation.
func forecast_DeletePredictor(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DeletePredictorInput{
		// PredictorArn: *string, // Required
	}

	if len(_forecastPredictorArn) > 0 {
		input.PredictorArn = aws.String(_forecastPredictorArn)
	}

	if resp, err := client.DeletePredictor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a predictor backtest export job.
func forecast_DeletePredictorBacktestExportJob(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DeletePredictorBacktestExportJobInput{
		// PredictorBacktestExportJobArn: *string, // Required
	}

	if len(_forecastPredictorBacktestExportJobArn) > 0 {
		input.PredictorBacktestExportJobArn = aws.String(_forecastPredictorBacktestExportJobArn)
	}

	if resp, err := client.DeletePredictorBacktestExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an entire resource tree. This operation will delete the parent resource
// and its child resources.
//
// Child resources are resources that were created from another resource. For
// example, when a forecast is generated from a predictor, the forecast is the
// child resource and the predictor is the parent resource.
//
// Amazon Forecast resources possess the following parent-child resource
// hierarchies:
//
// - Dataset: dataset import jobs
//
// - Dataset Group: predictors, predictor backtest export jobs, forecasts,
// forecast export jobs
//
// - Predictor: predictor backtest export jobs, forecasts, forecast export jobs
//
// - Forecast: forecast export jobs
//
// DeleteResourceTree will only delete Amazon Forecast resources, and will not
// delete datasets or exported files stored in Amazon S3.
func forecast_DeleteResourceTree(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DeleteResourceTreeInput{
		// ResourceArn: *string, // Required
	}

	if len(_forecastResourceArn) > 0 {
		input.ResourceArn = aws.String(_forecastResourceArn)
	}

	if resp, err := client.DeleteResourceTree(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a what-if analysis created using the CreateWhatIfAnalysis operation. You can delete only
// what-if analyses that have a status of ACTIVE or CREATE_FAILED . To get the
// status, use the DescribeWhatIfAnalysisoperation.
//
// You can't delete a what-if analysis while any of its forecasts are being
// exported.
func forecast_DeleteWhatIfAnalysis(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DeleteWhatIfAnalysisInput{
		// WhatIfAnalysisArn: *string, // Required
	}

	if len(_forecastWhatIfAnalysisArn) > 0 {
		input.WhatIfAnalysisArn = aws.String(_forecastWhatIfAnalysisArn)
	}

	if resp, err := client.DeleteWhatIfAnalysis(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a what-if forecast created using the CreateWhatIfForecast operation. You can delete only
// what-if forecasts that have a status of ACTIVE or CREATE_FAILED . To get the
// status, use the DescribeWhatIfForecastoperation.
//
// You can't delete a what-if forecast while it is being exported. After a what-if
// forecast is deleted, you can no longer query the what-if analysis.
func forecast_DeleteWhatIfForecast(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DeleteWhatIfForecastInput{
		// WhatIfForecastArn: *string, // Required
	}

	if len(_forecastWhatIfForecastArn) > 0 {
		input.WhatIfForecastArn = aws.String(_forecastWhatIfForecastArn)
	}

	if resp, err := client.DeleteWhatIfForecast(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a what-if forecast export created using the CreateWhatIfForecastExport operation. You can delete
// only what-if forecast exports that have a status of ACTIVE or CREATE_FAILED . To
// get the status, use the DescribeWhatIfForecastExportoperation.
func forecast_DeleteWhatIfForecastExport(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DeleteWhatIfForecastExportInput{
		// WhatIfForecastExportArn: *string, // Required
	}

	if len(_forecastWhatIfForecastExportArn) > 0 {
		input.WhatIfForecastExportArn = aws.String(_forecastWhatIfForecastExportArn)
	}

	if resp, err := client.DeleteWhatIfForecastExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a predictor created using the CreateAutoPredictor operation.
func forecast_DescribeAutoPredictor(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DescribeAutoPredictorInput{
		// PredictorArn: *string, // Required
	}

	if len(_forecastPredictorArn) > 0 {
		input.PredictorArn = aws.String(_forecastPredictorArn)
	}

	if resp, err := client.DescribeAutoPredictor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an Amazon Forecast dataset created using the [CreateDataset] operation.
// In addition to listing the parameters specified in the CreateDataset request,
// this operation includes the following dataset properties:
//
// - CreationTime
//
// - LastModificationTime
//
// - Status
//
// [CreateDataset]: https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDataset.html
func forecast_DescribeDataset(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DescribeDatasetInput{
		// DatasetArn: *string, // Required
	}

	if len(_forecastDatasetArn) > 0 {
		input.DatasetArn = aws.String(_forecastDatasetArn)
	}

	if resp, err := client.DescribeDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a dataset group created using the [CreateDatasetGroup] operation.
// In addition to listing the parameters provided in the CreateDatasetGroup
// request, this operation includes the following properties:
//
// - DatasetArns - The datasets belonging to the group.
//
// - CreationTime
//
// - LastModificationTime
//
// - Status
//
// [CreateDatasetGroup]: https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDatasetGroup.html
func forecast_DescribeDatasetGroup(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DescribeDatasetGroupInput{
		// DatasetGroupArn: *string, // Required
	}

	if len(_forecastDatasetGroupArn) > 0 {
		input.DatasetGroupArn = aws.String(_forecastDatasetGroupArn)
	}

	if resp, err := client.DescribeDatasetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a dataset import job created using the [CreateDatasetImportJob] operation.
// In addition to listing the parameters provided in the CreateDatasetImportJob
// request, this operation includes the following properties:
//
// - CreationTime
//
// - LastModificationTime
//
// - DataSize
//
// - FieldStatistics
//
// - Status
//
// - Message - If an error occurred, information about the error.
//
// [CreateDatasetImportJob]: https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDatasetImportJob.html
func forecast_DescribeDatasetImportJob(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DescribeDatasetImportJobInput{
		// DatasetImportJobArn: *string, // Required
	}

	if len(_forecastDatasetImportJobArn) > 0 {
		input.DatasetImportJobArn = aws.String(_forecastDatasetImportJobArn)
	}

	if resp, err := client.DescribeDatasetImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an Explainability resource created using the CreateExplainability operation.
func forecast_DescribeExplainability(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DescribeExplainabilityInput{
		// ExplainabilityArn: *string, // Required
	}

	if len(_forecastExplainabilityArn) > 0 {
		input.ExplainabilityArn = aws.String(_forecastExplainabilityArn)
	}

	if resp, err := client.DescribeExplainability(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an Explainability export created using the CreateExplainabilityExport operation.
func forecast_DescribeExplainabilityExport(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DescribeExplainabilityExportInput{
		// ExplainabilityExportArn: *string, // Required
	}

	if len(_forecastExplainabilityExportArn) > 0 {
		input.ExplainabilityExportArn = aws.String(_forecastExplainabilityExportArn)
	}

	if resp, err := client.DescribeExplainabilityExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a forecast created using the CreateForecast operation.
// In addition to listing the properties provided in the CreateForecast request,
// this operation lists the following properties:
//
// - DatasetGroupArn - The dataset group that provided the training data.
//
// - CreationTime
//
// - LastModificationTime
//
// - Status
//
// - Message - If an error occurred, information about the error.
func forecast_DescribeForecast(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DescribeForecastInput{
		// ForecastArn: *string, // Required
	}

	if len(_forecastForecastArn) > 0 {
		input.ForecastArn = aws.String(_forecastForecastArn)
	}

	if resp, err := client.DescribeForecast(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a forecast export job created using the CreateForecastExportJob operation.
// In addition to listing the properties provided by the user in the
// CreateForecastExportJob request, this operation lists the following properties:
//
// - CreationTime
//
// - LastModificationTime
//
// - Status
//
// - Message - If an error occurred, information about the error.
func forecast_DescribeForecastExportJob(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DescribeForecastExportJobInput{
		// ForecastExportJobArn: *string, // Required
	}

	if len(_forecastForecastExportJobArn) > 0 {
		input.ForecastExportJobArn = aws.String(_forecastForecastExportJobArn)
	}

	if resp, err := client.DescribeForecastExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a monitor resource. In addition to listing the properties provided in
// the CreateMonitorrequest, this operation lists the following properties:
//
// - Baseline
//
// - CreationTime
//
// - LastEvaluationTime
//
// - LastEvaluationState
//
// - LastModificationTime
//
// - Message
//
// - Status
func forecast_DescribeMonitor(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DescribeMonitorInput{
		// MonitorArn: *string, // Required
	}

	if len(_forecastMonitorArn) > 0 {
		input.MonitorArn = aws.String(_forecastMonitorArn)
	}

	if resp, err := client.DescribeMonitor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is only valid for legacy predictors created with
// CreatePredictor. If you are not using a legacy predictor, use DescribeAutoPredictor.
//
// Describes a predictor created using the CreatePredictor operation.
//
// In addition to listing the properties provided in the CreatePredictor request,
// this operation lists the following properties:
//
// - DatasetImportJobArns - The dataset import jobs used to import training data.
//
// - AutoMLAlgorithmArns - If AutoML is performed, the algorithms that were
// evaluated.
//
// - CreationTime
//
// - LastModificationTime
//
// - Status
//
// - Message - If an error occurred, information about the error.
func forecast_DescribePredictor(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DescribePredictorInput{
		// PredictorArn: *string, // Required
	}

	if len(_forecastPredictorArn) > 0 {
		input.PredictorArn = aws.String(_forecastPredictorArn)
	}

	if resp, err := client.DescribePredictor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a predictor backtest export job created using the CreatePredictorBacktestExportJob operation.
// In addition to listing the properties provided by the user in the
// CreatePredictorBacktestExportJob request, this operation lists the following
// properties:
//
// - CreationTime
//
// - LastModificationTime
//
// - Status
//
// - Message (if an error occurred)
func forecast_DescribePredictorBacktestExportJob(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DescribePredictorBacktestExportJobInput{
		// PredictorBacktestExportJobArn: *string, // Required
	}

	if len(_forecastPredictorBacktestExportJobArn) > 0 {
		input.PredictorBacktestExportJobArn = aws.String(_forecastPredictorBacktestExportJobArn)
	}

	if resp, err := client.DescribePredictorBacktestExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the what-if analysis created using the CreateWhatIfAnalysis operation.
// In addition to listing the properties provided in the CreateWhatIfAnalysis
// request, this operation lists the following properties:
//
// - CreationTime
//
// - LastModificationTime
//
// - Message - If an error occurred, information about the error.
//
// - Status
func forecast_DescribeWhatIfAnalysis(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DescribeWhatIfAnalysisInput{
		// WhatIfAnalysisArn: *string, // Required
	}

	if len(_forecastWhatIfAnalysisArn) > 0 {
		input.WhatIfAnalysisArn = aws.String(_forecastWhatIfAnalysisArn)
	}

	if resp, err := client.DescribeWhatIfAnalysis(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the what-if forecast created using the CreateWhatIfForecast operation.
// In addition to listing the properties provided in the CreateWhatIfForecast
// request, this operation lists the following properties:
//
// - CreationTime
//
// - LastModificationTime
//
// - Message - If an error occurred, information about the error.
//
// - Status
func forecast_DescribeWhatIfForecast(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DescribeWhatIfForecastInput{
		// WhatIfForecastArn: *string, // Required
	}

	if len(_forecastWhatIfForecastArn) > 0 {
		input.WhatIfForecastArn = aws.String(_forecastWhatIfForecastArn)
	}

	if resp, err := client.DescribeWhatIfForecast(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the what-if forecast export created using the CreateWhatIfForecastExport operation.
// In addition to listing the properties provided in the CreateWhatIfForecastExport
// request, this operation lists the following properties:
//
// - CreationTime
//
// - LastModificationTime
//
// - Message - If an error occurred, information about the error.
//
// - Status
func forecast_DescribeWhatIfForecastExport(cfg aws.Config, client *forecast.Client) {
	input := &forecast.DescribeWhatIfForecastExportInput{
		// WhatIfForecastExportArn: *string, // Required
	}

	if len(_forecastWhatIfForecastExportArn) > 0 {
		input.WhatIfForecastExportArn = aws.String(_forecastWhatIfForecastExportArn)
	}

	if resp, err := client.DescribeWhatIfForecastExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides metrics on the accuracy of the models that were trained by the CreatePredictor
// operation. Use metrics to see how well the model performed and to decide whether
// to use the predictor to generate a forecast. For more information, see [Predictor Metrics].
//
// This operation generates metrics for each backtest window that was evaluated.
// The number of backtest windows ( NumberOfBacktestWindows ) is specified using
// the EvaluationParametersobject, which is optionally included in the CreatePredictor request. If
// NumberOfBacktestWindows isn't specified, the number defaults to one.
//
// The parameters of the filling method determine which items contribute to the
// metrics. If you want all items to contribute, specify zero . If you want only
// those items that have complete data in the range being evaluated to contribute,
// specify nan . For more information, see FeaturizationMethod.
//
// Before you can get accuracy metrics, the Status of the predictor must be ACTIVE
// , signifying that training has completed. To get the status, use the DescribePredictoroperation.
//
// [Predictor Metrics]: https://docs.aws.amazon.com/forecast/latest/dg/metrics.html
func forecast_GetAccuracyMetrics(cfg aws.Config, client *forecast.Client) {
	input := &forecast.GetAccuracyMetricsInput{
		// PredictorArn: *string, // Required
	}

	if len(_forecastPredictorArn) > 0 {
		input.PredictorArn = aws.String(_forecastPredictorArn)
	}

	if resp, err := client.GetAccuracyMetrics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of dataset groups created using the [CreateDatasetGroup] operation. For each dataset
// group, this operation returns a summary of its properties, including its Amazon
// Resource Name (ARN). You can retrieve the complete set of properties by using
// the dataset group ARN with the [DescribeDatasetGroup]operation.
//
// [DescribeDatasetGroup]: https://docs.aws.amazon.com/forecast/latest/dg/API_DescribeDatasetGroup.html
// [CreateDatasetGroup]: https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDatasetGroup.html
func forecast_ListDatasetGroups(cfg aws.Config, client *forecast.Client) {
	input := &forecast.ListDatasetGroupsInput{}

	if len(_forecastMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _forecastMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_forecastNextToken) > 0 {
		input.NextToken = aws.String(_forecastNextToken)
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

	var results []*forecast.ListDatasetGroupsOutput
	p := forecast.NewListDatasetGroupsPaginator(client, input)
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

// Returns a list of dataset import jobs created using the [CreateDatasetImportJob] operation. For each
// import job, this operation returns a summary of its properties, including its
// Amazon Resource Name (ARN). You can retrieve the complete set of properties by
// using the ARN with the [DescribeDatasetImportJob]operation. You can filter the list by providing an array
// of [Filter]objects.
//
// [DescribeDatasetImportJob]: https://docs.aws.amazon.com/forecast/latest/dg/API_DescribeDatasetImportJob.html
// [Filter]: https://docs.aws.amazon.com/forecast/latest/dg/API_Filter.html
// [CreateDatasetImportJob]: https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDatasetImportJob.html
func forecast_ListDatasetImportJobs(cfg aws.Config, client *forecast.Client) {
	input := &forecast.ListDatasetImportJobsInput{}

	if len(_forecastFilters) > 0 {
		if err := assignInputField(input, "Filters", _forecastFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_forecastMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _forecastMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_forecastNextToken) > 0 {
		input.NextToken = aws.String(_forecastNextToken)
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

	var results []*forecast.ListDatasetImportJobsOutput
	p := forecast.NewListDatasetImportJobsPaginator(client, input)
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

// Returns a list of datasets created using the [CreateDataset] operation. For each dataset, a
// summary of its properties, including its Amazon Resource Name (ARN), is
// returned. To retrieve the complete set of properties, use the ARN with the [DescribeDataset]
// operation.
//
// [DescribeDataset]: https://docs.aws.amazon.com/forecast/latest/dg/API_DescribeDataset.html
// [CreateDataset]: https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDataset.html
func forecast_ListDatasets(cfg aws.Config, client *forecast.Client) {
	input := &forecast.ListDatasetsInput{}

	if len(_forecastMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _forecastMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_forecastNextToken) > 0 {
		input.NextToken = aws.String(_forecastNextToken)
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

	var results []*forecast.ListDatasetsOutput
	p := forecast.NewListDatasetsPaginator(client, input)
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

// Returns a list of Explainability resources created using the CreateExplainability operation. This
// operation returns a summary for each Explainability. You can filter the list
// using an array of Filterobjects.
//
// To retrieve the complete set of properties for a particular Explainability
// resource, use the ARN with the DescribeExplainabilityoperation.
func forecast_ListExplainabilities(cfg aws.Config, client *forecast.Client) {
	input := &forecast.ListExplainabilitiesInput{}

	if len(_forecastFilters) > 0 {
		if err := assignInputField(input, "Filters", _forecastFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_forecastMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _forecastMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_forecastNextToken) > 0 {
		input.NextToken = aws.String(_forecastNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListExplainabilities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*forecast.ListExplainabilitiesOutput
	p := forecast.NewListExplainabilitiesPaginator(client, input)
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

// Returns a list of Explainability exports created using the CreateExplainabilityExport operation. This
// operation returns a summary for each Explainability export. You can filter the
// list using an array of Filterobjects.
//
// To retrieve the complete set of properties for a particular Explainability
// export, use the ARN with the DescribeExplainabilityoperation.
func forecast_ListExplainabilityExports(cfg aws.Config, client *forecast.Client) {
	input := &forecast.ListExplainabilityExportsInput{}

	if len(_forecastFilters) > 0 {
		if err := assignInputField(input, "Filters", _forecastFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_forecastMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _forecastMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_forecastNextToken) > 0 {
		input.NextToken = aws.String(_forecastNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListExplainabilityExports(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*forecast.ListExplainabilityExportsOutput
	p := forecast.NewListExplainabilityExportsPaginator(client, input)
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

// Returns a list of forecast export jobs created using the CreateForecastExportJob operation. For each
// forecast export job, this operation returns a summary of its properties,
// including its Amazon Resource Name (ARN). To retrieve the complete set of
// properties, use the ARN with the DescribeForecastExportJoboperation. You can filter the list using an
// array of Filterobjects.
func forecast_ListForecastExportJobs(cfg aws.Config, client *forecast.Client) {
	input := &forecast.ListForecastExportJobsInput{}

	if len(_forecastFilters) > 0 {
		if err := assignInputField(input, "Filters", _forecastFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_forecastMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _forecastMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_forecastNextToken) > 0 {
		input.NextToken = aws.String(_forecastNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListForecastExportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*forecast.ListForecastExportJobsOutput
	p := forecast.NewListForecastExportJobsPaginator(client, input)
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

// Returns a list of forecasts created using the CreateForecast operation. For each forecast,
// this operation returns a summary of its properties, including its Amazon
// Resource Name (ARN). To retrieve the complete set of properties, specify the ARN
// with the DescribeForecastoperation. You can filter the list using an array of Filter objects.
func forecast_ListForecasts(cfg aws.Config, client *forecast.Client) {
	input := &forecast.ListForecastsInput{}

	if len(_forecastFilters) > 0 {
		if err := assignInputField(input, "Filters", _forecastFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_forecastMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _forecastMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_forecastNextToken) > 0 {
		input.NextToken = aws.String(_forecastNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListForecasts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*forecast.ListForecastsOutput
	p := forecast.NewListForecastsPaginator(client, input)
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

// Returns a list of the monitoring evaluation results and predictor events
// collected by the monitor resource during different windows of time.
//
// For information about monitoring see predictor-monitoring. For more information about retrieving
// monitoring results see [Viewing Monitoring Results].
//
// [Viewing Monitoring Results]: https://docs.aws.amazon.com/forecast/latest/dg/predictor-monitoring-results.html
func forecast_ListMonitorEvaluations(cfg aws.Config, client *forecast.Client) {
	input := &forecast.ListMonitorEvaluationsInput{
		// MonitorArn: *string, // Required
	}

	if len(_forecastMonitorArn) > 0 {
		input.MonitorArn = aws.String(_forecastMonitorArn)
	}
	if len(_forecastFilters) > 0 {
		if err := assignInputField(input, "Filters", _forecastFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_forecastMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _forecastMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_forecastNextToken) > 0 {
		input.NextToken = aws.String(_forecastNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMonitorEvaluations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*forecast.ListMonitorEvaluationsOutput
	p := forecast.NewListMonitorEvaluationsPaginator(client, input)
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

// Returns a list of monitors created with the CreateMonitor operation and CreateAutoPredictor operation. For each
// monitor resource, this operation returns of a summary of its properties,
// including its Amazon Resource Name (ARN). You can retrieve a complete set of
// properties of a monitor resource by specify the monitor's ARN in the DescribeMonitoroperation.
func forecast_ListMonitors(cfg aws.Config, client *forecast.Client) {
	input := &forecast.ListMonitorsInput{}

	if len(_forecastFilters) > 0 {
		if err := assignInputField(input, "Filters", _forecastFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_forecastMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _forecastMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_forecastNextToken) > 0 {
		input.NextToken = aws.String(_forecastNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMonitors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*forecast.ListMonitorsOutput
	p := forecast.NewListMonitorsPaginator(client, input)
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

// Returns a list of predictor backtest export jobs created using the CreatePredictorBacktestExportJob operation.
// This operation returns a summary for each backtest export job. You can filter
// the list using an array of Filterobjects.
//
// To retrieve the complete set of properties for a particular backtest export
// job, use the ARN with the DescribePredictorBacktestExportJoboperation.
func forecast_ListPredictorBacktestExportJobs(cfg aws.Config, client *forecast.Client) {
	input := &forecast.ListPredictorBacktestExportJobsInput{}

	if len(_forecastFilters) > 0 {
		if err := assignInputField(input, "Filters", _forecastFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_forecastMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _forecastMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_forecastNextToken) > 0 {
		input.NextToken = aws.String(_forecastNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPredictorBacktestExportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*forecast.ListPredictorBacktestExportJobsOutput
	p := forecast.NewListPredictorBacktestExportJobsPaginator(client, input)
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

// Returns a list of predictors created using the CreateAutoPredictor or CreatePredictor operations. For each
// predictor, this operation returns a summary of its properties, including its
// Amazon Resource Name (ARN).
//
// You can retrieve the complete set of properties by using the ARN with the DescribeAutoPredictor and DescribePredictor
// operations. You can filter the list using an array of Filterobjects.
func forecast_ListPredictors(cfg aws.Config, client *forecast.Client) {
	input := &forecast.ListPredictorsInput{}

	if len(_forecastFilters) > 0 {
		if err := assignInputField(input, "Filters", _forecastFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_forecastMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _forecastMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_forecastNextToken) > 0 {
		input.NextToken = aws.String(_forecastNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPredictors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*forecast.ListPredictorsOutput
	p := forecast.NewListPredictorsPaginator(client, input)
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

// Lists the tags for an Amazon Forecast resource.
func forecast_ListTagsForResource(cfg aws.Config, client *forecast.Client) {
	input := &forecast.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_forecastResourceArn) > 0 {
		input.ResourceArn = aws.String(_forecastResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of what-if analyses created using the CreateWhatIfAnalysis operation. For each
// what-if analysis, this operation returns a summary of its properties, including
// its Amazon Resource Name (ARN). You can retrieve the complete set of properties
// by using the what-if analysis ARN with the DescribeWhatIfAnalysisoperation.
func forecast_ListWhatIfAnalyses(cfg aws.Config, client *forecast.Client) {
	input := &forecast.ListWhatIfAnalysesInput{}

	if len(_forecastFilters) > 0 {
		if err := assignInputField(input, "Filters", _forecastFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_forecastMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _forecastMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_forecastNextToken) > 0 {
		input.NextToken = aws.String(_forecastNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWhatIfAnalyses(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*forecast.ListWhatIfAnalysesOutput
	p := forecast.NewListWhatIfAnalysesPaginator(client, input)
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

// Returns a list of what-if forecast exports created using the CreateWhatIfForecastExport operation. For
// each what-if forecast export, this operation returns a summary of its
// properties, including its Amazon Resource Name (ARN). You can retrieve the
// complete set of properties by using the what-if forecast export ARN with the DescribeWhatIfForecastExport
// operation.
func forecast_ListWhatIfForecastExports(cfg aws.Config, client *forecast.Client) {
	input := &forecast.ListWhatIfForecastExportsInput{}

	if len(_forecastFilters) > 0 {
		if err := assignInputField(input, "Filters", _forecastFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_forecastMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _forecastMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_forecastNextToken) > 0 {
		input.NextToken = aws.String(_forecastNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWhatIfForecastExports(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*forecast.ListWhatIfForecastExportsOutput
	p := forecast.NewListWhatIfForecastExportsPaginator(client, input)
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

// Returns a list of what-if forecasts created using the CreateWhatIfForecast operation. For each
// what-if forecast, this operation returns a summary of its properties, including
// its Amazon Resource Name (ARN). You can retrieve the complete set of properties
// by using the what-if forecast ARN with the DescribeWhatIfForecastoperation.
func forecast_ListWhatIfForecasts(cfg aws.Config, client *forecast.Client) {
	input := &forecast.ListWhatIfForecastsInput{}

	if len(_forecastFilters) > 0 {
		if err := assignInputField(input, "Filters", _forecastFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_forecastMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _forecastMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_forecastNextToken) > 0 {
		input.NextToken = aws.String(_forecastNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWhatIfForecasts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*forecast.ListWhatIfForecastsOutput
	p := forecast.NewListWhatIfForecastsPaginator(client, input)
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

// Resumes a stopped monitor resource.
func forecast_ResumeResource(cfg aws.Config, client *forecast.Client) {
	input := &forecast.ResumeResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_forecastResourceArn) > 0 {
		input.ResourceArn = aws.String(_forecastResourceArn)
	}

	if resp, err := client.ResumeResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a resource.
// The resource undergoes the following states: CREATE_STOPPING and CREATE_STOPPED
// . You cannot resume a resource once it has been stopped.
//
// This operation can be applied to the following resources (and their
// corresponding child resources):
//
// - Dataset Import Job
//
// - Predictor Job
//
// - Forecast Job
//
// - Forecast Export Job
//
// - Predictor Backtest Export Job
//
// - Explainability Job
//
// - Explainability Export Job
func forecast_StopResource(cfg aws.Config, client *forecast.Client) {
	input := &forecast.StopResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_forecastResourceArn) > 0 {
		input.ResourceArn = aws.String(_forecastResourceArn)
	}

	if resp, err := client.StopResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified tags to a resource with the specified resourceArn . If
// existing tags on a resource are not specified in the request parameters, they
// are not changed. When a resource is deleted, the tags associated with that
// resource are also deleted.
func forecast_TagResource(cfg aws.Config, client *forecast.Client) {
	input := &forecast.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_forecastResourceArn) > 0 {
		input.ResourceArn = aws.String(_forecastResourceArn)
	}
	if len(_forecastTags) > 0 {
		if err := assignInputField(input, "Tags", _forecastTags); err != nil {
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

// Deletes the specified tags from a resource.
func forecast_UntagResource(cfg aws.Config, client *forecast.Client) {
	input := &forecast.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_forecastResourceArn) > 0 {
		input.ResourceArn = aws.String(_forecastResourceArn)
	}
	if len(_forecastTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _forecastTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Replaces the datasets in a dataset group with the specified datasets.
// The Status of the dataset group must be ACTIVE before you can use the dataset
// group to create a predictor. Use the [DescribeDatasetGroup]operation to get the status.
//
// [DescribeDatasetGroup]: https://docs.aws.amazon.com/forecast/latest/dg/API_DescribeDatasetGroup.html
func forecast_UpdateDatasetGroup(cfg aws.Config, client *forecast.Client) {
	input := &forecast.UpdateDatasetGroupInput{
		// DatasetArns: []string, // Required
		// DatasetGroupArn: *string, // Required
	}

	if len(_forecastDatasetArns) > 0 {
		input.DatasetArns = append([]string(nil), _forecastDatasetArns...)
	}
	if len(_forecastDatasetGroupArn) > 0 {
		input.DatasetGroupArn = aws.String(_forecastDatasetGroupArn)
	}

	if resp, err := client.UpdateDatasetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_forecastCmd)
	_forecastCmd.Flags().SortFlags = false

	_forecastCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_forecastCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_forecastCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_forecastCmd.Flags().StringVarP(&_forecastAlgorithmArn, "algorithm-arn", "", "", "Algorithm ARN")
	_forecastCmd.Flags().StringVarP(&_forecastAutoMLOverrideStrategy, "auto-ml-override-strategy", "", "", "Auto Ml Override Strategy")
	_forecastCmd.Flags().StringVarP(&_forecastDataConfig, "data-config", "", "", "Data Config")
	_forecastCmd.Flags().StringVarP(&_forecastDataFrequency, "data-frequency", "", "", "Data Frequency")
	_forecastCmd.Flags().StringVarP(&_forecastDataSource, "data-source", "", "", "Data Source")
	_forecastCmd.Flags().StringVarP(&_forecastDatasetArn, "dataset-arn", "", "", "Dataset ARN")
	_forecastCmd.Flags().StringSliceVarP(&_forecastDatasetArns, "dataset-arns", "", nil, "Dataset Arns")
	_forecastCmd.Flags().StringVarP(&_forecastDatasetGroupArn, "dataset-group-arn", "", "", "Dataset Group ARN")
	_forecastCmd.Flags().StringVarP(&_forecastDatasetGroupName, "dataset-group-name", "", "", "Dataset Group Name")
	_forecastCmd.Flags().StringVarP(&_forecastDatasetImportJobArn, "dataset-import-job-arn", "", "", "Dataset Import Job ARN")
	_forecastCmd.Flags().StringVarP(&_forecastDatasetImportJobName, "dataset-import-job-name", "", "", "Dataset Import Job Name")
	_forecastCmd.Flags().StringVarP(&_forecastDatasetName, "dataset-name", "", "", "Dataset Name")
	_forecastCmd.Flags().StringVarP(&_forecastDatasetType, "dataset-type", "", "", "Dataset Type")
	_forecastCmd.Flags().StringVarP(&_forecastDestination, "destination", "", "", "Destination")
	_forecastCmd.Flags().StringVarP(&_forecastDomain, "domain", "", "", "Domain")
	_forecastCmd.Flags().StringVarP(&_forecastEnableVisualization, "enable-visualization", "", "", "Enable Visualization")
	_forecastCmd.Flags().StringVarP(&_forecastEncryptionConfig, "encryption-config", "", "", "Encryption Config")
	_forecastCmd.Flags().StringVarP(&_forecastEndDateTime, "end-date-time", "", "", "End Date Time")
	_forecastCmd.Flags().StringVarP(&_forecastEvaluationParameters, "evaluation-parameters", "", "", "Evaluation Parameters")
	_forecastCmd.Flags().StringVarP(&_forecastExplainPredictor, "explain-predictor", "", "", "Explain Predictor")
	_forecastCmd.Flags().StringVarP(&_forecastExplainabilityArn, "explainability-arn", "", "", "Explainability ARN")
	_forecastCmd.Flags().StringVarP(&_forecastExplainabilityConfig, "explainability-config", "", "", "Explainability Config")
	_forecastCmd.Flags().StringVarP(&_forecastExplainabilityExportArn, "explainability-export-arn", "", "", "Explainability Export ARN")
	_forecastCmd.Flags().StringVarP(&_forecastExplainabilityExportName, "explainability-export-name", "", "", "Explainability Export Name")
	_forecastCmd.Flags().StringVarP(&_forecastExplainabilityName, "explainability-name", "", "", "Explainability Name")
	_forecastCmd.Flags().StringVarP(&_forecastFeaturizationConfig, "featurization-config", "", "", "Featurization Config")
	_forecastCmd.Flags().StringVarP(&_forecastFilters, "filters", "", "", "Filters")
	_forecastCmd.Flags().StringVarP(&_forecastForecastArn, "forecast-arn", "", "", "Forecast ARN")
	_forecastCmd.Flags().StringSliceVarP(&_forecastForecastDimensions, "forecast-dimensions", "", nil, "Forecast Dimensions")
	_forecastCmd.Flags().StringVarP(&_forecastForecastExportJobArn, "forecast-export-job-arn", "", "", "Forecast Export Job ARN")
	_forecastCmd.Flags().StringVarP(&_forecastForecastExportJobName, "forecast-export-job-name", "", "", "Forecast Export Job Name")
	_forecastCmd.Flags().StringVarP(&_forecastForecastFrequency, "forecast-frequency", "", "", "Forecast Frequency")
	_forecastCmd.Flags().StringVarP(&_forecastForecastHorizon, "forecast-horizon", "", "", "Forecast Horizon")
	_forecastCmd.Flags().StringVarP(&_forecastForecastName, "forecast-name", "", "", "Forecast Name")
	_forecastCmd.Flags().StringSliceVarP(&_forecastForecastTypes, "forecast-types", "", nil, "Forecast Types")
	_forecastCmd.Flags().StringVarP(&_forecastFormat, "format", "", "", "Format")
	_forecastCmd.Flags().StringVarP(&_forecastGeolocationFormat, "geolocation-format", "", "", "Geolocation Format")
	_forecastCmd.Flags().StringVarP(&_forecastHPOConfig, "hpo-config", "", "", "Hpo Config")
	_forecastCmd.Flags().StringVarP(&_forecastImportMode, "import-mode", "", "", "Import Mode")
	_forecastCmd.Flags().StringVarP(&_forecastInputDataConfig, "input-data-config", "", "", "Input Data Config")
	_forecastCmd.Flags().StringVarP(&_forecastMaxResults, "max-results", "", "", "Max Results")
	_forecastCmd.Flags().StringVarP(&_forecastMonitorArn, "monitor-arn", "", "", "Monitor ARN")
	_forecastCmd.Flags().StringVarP(&_forecastMonitorConfig, "monitor-config", "", "", "Monitor Config")
	_forecastCmd.Flags().StringVarP(&_forecastMonitorName, "monitor-name", "", "", "Monitor Name")
	_forecastCmd.Flags().StringVarP(&_forecastNextToken, "next-token", "", "", "Next Token")
	_forecastCmd.Flags().StringVarP(&_forecastOptimizationMetric, "optimization-metric", "", "", "Optimization Metric")
	_forecastCmd.Flags().StringVarP(&_forecastPerformAutoML, "perform-auto-ml", "", "", "Perform Auto Ml")
	_forecastCmd.Flags().StringVarP(&_forecastPerformHPO, "perform-hpo", "", "", "Perform Hpo")
	_forecastCmd.Flags().StringVarP(&_forecastPredictorArn, "predictor-arn", "", "", "Predictor ARN")
	_forecastCmd.Flags().StringVarP(&_forecastPredictorBacktestExportJobArn, "predictor-backtest-export-job-arn", "", "", "Predictor Backtest Export Job ARN")
	_forecastCmd.Flags().StringVarP(&_forecastPredictorBacktestExportJobName, "predictor-backtest-export-job-name", "", "", "Predictor Backtest Export Job Name")
	_forecastCmd.Flags().StringVarP(&_forecastPredictorName, "predictor-name", "", "", "Predictor Name")
	_forecastCmd.Flags().StringVarP(&_forecastReferencePredictorArn, "reference-predictor-arn", "", "", "Reference Predictor ARN")
	_forecastCmd.Flags().StringVarP(&_forecastResourceArn, "resource-arn", "", "", "Resource ARN")
	_forecastCmd.Flags().StringVarP(&_forecastSchema, "schema", "", "", "Schema")
	_forecastCmd.Flags().StringVarP(&_forecastStartDateTime, "start-date-time", "", "", "Start Date Time")
	_forecastCmd.Flags().StringSliceVarP(&_forecastTagKeys, "tag-keys", "", nil, "Tag Keys")
	_forecastCmd.Flags().StringVarP(&_forecastTags, "tags", "", "", "Tags")
	_forecastCmd.Flags().StringVarP(&_forecastTimeAlignmentBoundary, "time-alignment-boundary", "", "", "Time Alignment Boundary")
	_forecastCmd.Flags().StringVarP(&_forecastTimeSeriesReplacementsDataSource, "time-series-replacements-data-source", "", "", "Time Series Replacements Data Source")
	_forecastCmd.Flags().StringVarP(&_forecastTimeSeriesSelector, "time-series-selector", "", "", "Time Series Selector")
	_forecastCmd.Flags().StringVarP(&_forecastTimeSeriesTransformations, "time-series-transformations", "", "", "Time Series Transformations")
	_forecastCmd.Flags().StringVarP(&_forecastTimeZone, "time-zone", "", "", "Time Zone")
	_forecastCmd.Flags().StringVarP(&_forecastTimestampFormat, "timestamp-format", "", "", "Timestamp Format")
	_forecastCmd.Flags().StringVarP(&_forecastTrainingParameters, "training-parameters", "", "", "Training Parameters")
	_forecastCmd.Flags().StringVarP(&_forecastUseGeolocationForTimeZone, "use-geolocation-for-time-zone", "", "", "Use Geolocation For Time Zone")
	_forecastCmd.Flags().StringVarP(&_forecastWhatIfAnalysisArn, "what-if-analysis-arn", "", "", "What If Analysis ARN")
	_forecastCmd.Flags().StringVarP(&_forecastWhatIfAnalysisName, "what-if-analysis-name", "", "", "What If Analysis Name")
	_forecastCmd.Flags().StringVarP(&_forecastWhatIfForecastArn, "what-if-forecast-arn", "", "", "What If Forecast ARN")
	_forecastCmd.Flags().StringSliceVarP(&_forecastWhatIfForecastArns, "what-if-forecast-arns", "", nil, "What If Forecast Arns")
	_forecastCmd.Flags().StringVarP(&_forecastWhatIfForecastExportArn, "what-if-forecast-export-arn", "", "", "What If Forecast Export ARN")
	_forecastCmd.Flags().StringVarP(&_forecastWhatIfForecastExportName, "what-if-forecast-export-name", "", "", "What If Forecast Export Name")
	_forecastCmd.Flags().StringVarP(&_forecastWhatIfForecastName, "what-if-forecast-name", "", "", "What If Forecast Name")

	_forecastCmd.Flags().BoolVarP(&_forecastCreateAutoPredictor, "create-auto-predictor", "", false, "Create Auto Predictor")
	_forecastCmd.Flags().BoolVarP(&_forecastCreateDataset, "create-dataset", "", false, "Create Dataset")
	_forecastCmd.Flags().BoolVarP(&_forecastCreateDatasetGroup, "create-dataset-group", "", false, "Create Dataset Group")
	_forecastCmd.Flags().BoolVarP(&_forecastCreateDatasetImportJob, "create-dataset-import-job", "", false, "Create Dataset Import Job")
	_forecastCmd.Flags().BoolVarP(&_forecastCreateExplainability, "create-explainability", "", false, "Create Explainability")
	_forecastCmd.Flags().BoolVarP(&_forecastCreateExplainabilityExport, "create-explainability-export", "", false, "Create Explainability Export")
	_forecastCmd.Flags().BoolVarP(&_forecastCreateForecast, "create-forecast", "", false, "Create Forecast")
	_forecastCmd.Flags().BoolVarP(&_forecastCreateForecastExportJob, "create-forecast-export-job", "", false, "Create Forecast Export Job")
	_forecastCmd.Flags().BoolVarP(&_forecastCreateMonitor, "create-monitor", "", false, "Create Monitor")
	_forecastCmd.Flags().BoolVarP(&_forecastCreatePredictor, "create-predictor", "", false, "Create Predictor")
	_forecastCmd.Flags().BoolVarP(&_forecastCreatePredictorBacktestExportJob, "create-predictor-backtest-export-job", "", false, "Create Predictor Backtest Export Job")
	_forecastCmd.Flags().BoolVarP(&_forecastCreateWhatIfAnalysis, "create-what-if-analysis", "", false, "Create What If Analysis")
	_forecastCmd.Flags().BoolVarP(&_forecastCreateWhatIfForecast, "create-what-if-forecast", "", false, "Create What If Forecast")
	_forecastCmd.Flags().BoolVarP(&_forecastCreateWhatIfForecastExport, "create-what-if-forecast-export", "", false, "Create What If Forecast Export")
	_forecastCmd.Flags().BoolVarP(&_forecastDeleteDataset, "delete-dataset", "", false, "Delete Dataset")
	_forecastCmd.Flags().BoolVarP(&_forecastDeleteDatasetGroup, "delete-dataset-group", "", false, "Delete Dataset Group")
	_forecastCmd.Flags().BoolVarP(&_forecastDeleteDatasetImportJob, "delete-dataset-import-job", "", false, "Delete Dataset Import Job")
	_forecastCmd.Flags().BoolVarP(&_forecastDeleteExplainability, "delete-explainability", "", false, "Delete Explainability")
	_forecastCmd.Flags().BoolVarP(&_forecastDeleteExplainabilityExport, "delete-explainability-export", "", false, "Delete Explainability Export")
	_forecastCmd.Flags().BoolVarP(&_forecastDeleteForecast, "delete-forecast", "", false, "Delete Forecast")
	_forecastCmd.Flags().BoolVarP(&_forecastDeleteForecastExportJob, "delete-forecast-export-job", "", false, "Delete Forecast Export Job")
	_forecastCmd.Flags().BoolVarP(&_forecastDeleteMonitor, "delete-monitor", "", false, "Delete Monitor")
	_forecastCmd.Flags().BoolVarP(&_forecastDeletePredictor, "delete-predictor", "", false, "Delete Predictor")
	_forecastCmd.Flags().BoolVarP(&_forecastDeletePredictorBacktestExportJob, "delete-predictor-backtest-export-job", "", false, "Delete Predictor Backtest Export Job")
	_forecastCmd.Flags().BoolVarP(&_forecastDeleteResourceTree, "delete-resource-tree", "", false, "Delete Resource Tree")
	_forecastCmd.Flags().BoolVarP(&_forecastDeleteWhatIfAnalysis, "delete-what-if-analysis", "", false, "Delete What If Analysis")
	_forecastCmd.Flags().BoolVarP(&_forecastDeleteWhatIfForecast, "delete-what-if-forecast", "", false, "Delete What If Forecast")
	_forecastCmd.Flags().BoolVarP(&_forecastDeleteWhatIfForecastExport, "delete-what-if-forecast-export", "", false, "Delete What If Forecast Export")
	_forecastCmd.Flags().BoolVarP(&_forecastDescribeAutoPredictor, "describe-auto-predictor", "", false, "Describe Auto Predictor")
	_forecastCmd.Flags().BoolVarP(&_forecastDescribeDataset, "describe-dataset", "", false, "Describe Dataset")
	_forecastCmd.Flags().BoolVarP(&_forecastDescribeDatasetGroup, "describe-dataset-group", "", false, "Describe Dataset Group")
	_forecastCmd.Flags().BoolVarP(&_forecastDescribeDatasetImportJob, "describe-dataset-import-job", "", false, "Describe Dataset Import Job")
	_forecastCmd.Flags().BoolVarP(&_forecastDescribeExplainability, "describe-explainability", "", false, "Describe Explainability")
	_forecastCmd.Flags().BoolVarP(&_forecastDescribeExplainabilityExport, "describe-explainability-export", "", false, "Describe Explainability Export")
	_forecastCmd.Flags().BoolVarP(&_forecastDescribeForecast, "describe-forecast", "", false, "Describe Forecast")
	_forecastCmd.Flags().BoolVarP(&_forecastDescribeForecastExportJob, "describe-forecast-export-job", "", false, "Describe Forecast Export Job")
	_forecastCmd.Flags().BoolVarP(&_forecastDescribeMonitor, "describe-monitor", "", false, "Describe Monitor")
	_forecastCmd.Flags().BoolVarP(&_forecastDescribePredictor, "describe-predictor", "", false, "Describe Predictor")
	_forecastCmd.Flags().BoolVarP(&_forecastDescribePredictorBacktestExportJob, "describe-predictor-backtest-export-job", "", false, "Describe Predictor Backtest Export Job")
	_forecastCmd.Flags().BoolVarP(&_forecastDescribeWhatIfAnalysis, "describe-what-if-analysis", "", false, "Describe What If Analysis")
	_forecastCmd.Flags().BoolVarP(&_forecastDescribeWhatIfForecast, "describe-what-if-forecast", "", false, "Describe What If Forecast")
	_forecastCmd.Flags().BoolVarP(&_forecastDescribeWhatIfForecastExport, "describe-what-if-forecast-export", "", false, "Describe What If Forecast Export")
	_forecastCmd.Flags().BoolVarP(&_forecastGetAccuracyMetrics, "get-accuracy-metrics", "", false, "Get Accuracy Metrics")
	_forecastCmd.Flags().BoolVarP(&_forecastListDatasetGroups, "list-dataset-groups", "", false, "List Dataset Groups")
	_forecastCmd.Flags().BoolVarP(&_forecastListDatasetImportJobs, "list-dataset-import-jobs", "", false, "List Dataset Import Jobs")
	_forecastCmd.Flags().BoolVarP(&_forecastListDatasets, "list-datasets", "", false, "List Datasets")
	_forecastCmd.Flags().BoolVarP(&_forecastListExplainabilities, "list-explainabilities", "", false, "List Explainabilities")
	_forecastCmd.Flags().BoolVarP(&_forecastListExplainabilityExports, "list-explainability-exports", "", false, "List Explainability Exports")
	_forecastCmd.Flags().BoolVarP(&_forecastListForecastExportJobs, "list-forecast-export-jobs", "", false, "List Forecast Export Jobs")
	_forecastCmd.Flags().BoolVarP(&_forecastListForecasts, "list-forecasts", "", false, "List Forecasts")
	_forecastCmd.Flags().BoolVarP(&_forecastListMonitorEvaluations, "list-monitor-evaluations", "", false, "List Monitor Evaluations")
	_forecastCmd.Flags().BoolVarP(&_forecastListMonitors, "list-monitors", "", false, "List Monitors")
	_forecastCmd.Flags().BoolVarP(&_forecastListPredictorBacktestExportJobs, "list-predictor-backtest-export-jobs", "", false, "List Predictor Backtest Export Jobs")
	_forecastCmd.Flags().BoolVarP(&_forecastListPredictors, "list-predictors", "", false, "List Predictors")
	_forecastCmd.Flags().BoolVarP(&_forecastListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_forecastCmd.Flags().BoolVarP(&_forecastListWhatIfAnalyses, "list-what-if-analyses", "", false, "List What If Analyses")
	_forecastCmd.Flags().BoolVarP(&_forecastListWhatIfForecastExports, "list-what-if-forecast-exports", "", false, "List What If Forecast Exports")
	_forecastCmd.Flags().BoolVarP(&_forecastListWhatIfForecasts, "list-what-if-forecasts", "", false, "List What If Forecasts")
	_forecastCmd.Flags().BoolVarP(&_forecastResumeResource, "resume-resource", "", false, "Resume Resource")
	_forecastCmd.Flags().BoolVarP(&_forecastStopResource, "stop-resource", "", false, "Stop Resource")
	_forecastCmd.Flags().BoolVarP(&_forecastTagResource, "tag-resource", "", false, "Tag Resource")
	_forecastCmd.Flags().BoolVarP(&_forecastUntagResource, "untag-resource", "", false, "Untag Resource")
	_forecastCmd.Flags().BoolVarP(&_forecastUpdateDatasetGroup, "update-dataset-group", "", false, "Update Dataset Group")

}
