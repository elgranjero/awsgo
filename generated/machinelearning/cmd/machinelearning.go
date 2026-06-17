package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// machinelearningCmd represents the machinelearning command
var _machinelearningCmd = &cobra.Command{
	Use:   "machinelearning",
	Short: "AWS machinelearning CLI",
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
		client := machinelearning.NewFromConfig(cfg)
		if _machinelearningAddTags {
			machinelearning_AddTags(cfg, client)
			return
		}
		if _machinelearningCreateBatchPrediction {
			machinelearning_CreateBatchPrediction(cfg, client)
			return
		}
		if _machinelearningCreateDataSourceFromRDS {
			machinelearning_CreateDataSourceFromRDS(cfg, client)
			return
		}
		if _machinelearningCreateDataSourceFromRedshift {
			machinelearning_CreateDataSourceFromRedshift(cfg, client)
			return
		}
		if _machinelearningCreateDataSourceFromS3 {
			machinelearning_CreateDataSourceFromS3(cfg, client)
			return
		}
		if _machinelearningCreateEvaluation {
			machinelearning_CreateEvaluation(cfg, client)
			return
		}
		if _machinelearningCreateMLModel {
			machinelearning_CreateMLModel(cfg, client)
			return
		}
		if _machinelearningCreateRealtimeEndpoint {
			machinelearning_CreateRealtimeEndpoint(cfg, client)
			return
		}
		if _machinelearningDeleteBatchPrediction {
			machinelearning_DeleteBatchPrediction(cfg, client)
			return
		}
		if _machinelearningDeleteDataSource {
			machinelearning_DeleteDataSource(cfg, client)
			return
		}
		if _machinelearningDeleteEvaluation {
			machinelearning_DeleteEvaluation(cfg, client)
			return
		}
		if _machinelearningDeleteMLModel {
			machinelearning_DeleteMLModel(cfg, client)
			return
		}
		if _machinelearningDeleteRealtimeEndpoint {
			machinelearning_DeleteRealtimeEndpoint(cfg, client)
			return
		}
		if _machinelearningDeleteTags {
			machinelearning_DeleteTags(cfg, client)
			return
		}
		if _machinelearningDescribeBatchPredictions {
			machinelearning_DescribeBatchPredictions(cfg, client)
			return
		}
		if _machinelearningDescribeDataSources {
			machinelearning_DescribeDataSources(cfg, client)
			return
		}
		if _machinelearningDescribeEvaluations {
			machinelearning_DescribeEvaluations(cfg, client)
			return
		}
		if _machinelearningDescribeMLModels {
			machinelearning_DescribeMLModels(cfg, client)
			return
		}
		if _machinelearningDescribeTags {
			machinelearning_DescribeTags(cfg, client)
			return
		}
		if _machinelearningGetBatchPrediction {
			machinelearning_GetBatchPrediction(cfg, client)
			return
		}
		if _machinelearningGetDataSource {
			machinelearning_GetDataSource(cfg, client)
			return
		}
		if _machinelearningGetEvaluation {
			machinelearning_GetEvaluation(cfg, client)
			return
		}
		if _machinelearningGetMLModel {
			machinelearning_GetMLModel(cfg, client)
			return
		}
		if _machinelearningPredict {
			machinelearning_Predict(cfg, client)
			return
		}
		if _machinelearningUpdateBatchPrediction {
			machinelearning_UpdateBatchPrediction(cfg, client)
			return
		}
		if _machinelearningUpdateDataSource {
			machinelearning_UpdateDataSource(cfg, client)
			return
		}
		if _machinelearningUpdateEvaluation {
			machinelearning_UpdateEvaluation(cfg, client)
			return
		}
		if _machinelearningUpdateMLModel {
			machinelearning_UpdateMLModel(cfg, client)
			return
		}

	},
}

var (
	_machinelearningAddTags                      bool
	_machinelearningCreateBatchPrediction        bool
	_machinelearningCreateDataSourceFromRDS      bool
	_machinelearningCreateDataSourceFromRedshift bool
	_machinelearningCreateDataSourceFromS3       bool
	_machinelearningCreateEvaluation             bool
	_machinelearningCreateMLModel                bool
	_machinelearningCreateRealtimeEndpoint       bool
	_machinelearningDeleteBatchPrediction        bool
	_machinelearningDeleteDataSource             bool
	_machinelearningDeleteEvaluation             bool
	_machinelearningDeleteMLModel                bool
	_machinelearningDeleteRealtimeEndpoint       bool
	_machinelearningDeleteTags                   bool
	_machinelearningDescribeBatchPredictions     bool
	_machinelearningDescribeDataSources          bool
	_machinelearningDescribeEvaluations          bool
	_machinelearningDescribeMLModels             bool
	_machinelearningDescribeTags                 bool
	_machinelearningGetBatchPrediction           bool
	_machinelearningGetDataSource                bool
	_machinelearningGetEvaluation                bool
	_machinelearningGetMLModel                   bool
	_machinelearningPredict                      bool
	_machinelearningUpdateBatchPrediction        bool
	_machinelearningUpdateDataSource             bool
	_machinelearningUpdateEvaluation             bool
	_machinelearningUpdateMLModel                bool

	_machinelearningBatchPredictionDataSourceId string
	_machinelearningBatchPredictionId           string
	_machinelearningBatchPredictionName         string
	_machinelearningComputeStatistics           string
	_machinelearningDataSourceId                string
	_machinelearningDataSourceName              string
	_machinelearningDataSpec                    string
	_machinelearningEQ                          string
	_machinelearningEvaluationDataSourceId      string
	_machinelearningEvaluationId                string
	_machinelearningEvaluationName              string
	_machinelearningFilterVariable              string
	_machinelearningGE                          string
	_machinelearningGT                          string
	_machinelearningLE                          string
	_machinelearningLimit                       string
	_machinelearningLT                          string
	_machinelearningMLModelId                   string
	_machinelearningMLModelName                 string
	_machinelearningMLModelType                 string
	_machinelearningNE                          string
	_machinelearningNextToken                   string
	_machinelearningOutputUri                   string
	_machinelearningParameters                  string
	_machinelearningPredictEndpoint             string
	_machinelearningPrefix                      string
	_machinelearningRDSData                     string
	_machinelearningRecipe                      string
	_machinelearningRecipeUri                   string
	_machinelearningRecord                      string
	_machinelearningResourceId                  string
	_machinelearningResourceType                string
	_machinelearningRoleARN                     string
	_machinelearningScoreThreshold              string
	_machinelearningSortOrder                   string
	_machinelearningTagKeys                     []string
	_machinelearningTags                        string
	_machinelearningTrainingDataSourceId        string
	_machinelearningVerbose                     string
)

// Adds one or more tags to an object, up to a limit of 10. Each tag consists of a
// key and an optional value. If you add a tag using a key that is already
// associated with the ML object, AddTags updates the tag's value.
func machinelearning_AddTags(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.AddTagsInput{
		// ResourceId: *string, // Required
		// ResourceType: types.TaggableResourceType, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_machinelearningResourceId) > 0 {
		input.ResourceId = aws.String(_machinelearningResourceId)
	}
	if len(_machinelearningResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _machinelearningResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_machinelearningTags) > 0 {
		if err := assignInputField(input, "Tags", _machinelearningTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates predictions for a group of observations. The observations to process
// exist in one or more data files referenced by a DataSource . This operation
// creates a new BatchPrediction , and uses an MLModel and the data files
// referenced by the DataSource as information sources.
//
// CreateBatchPrediction is an asynchronous operation. In response to
// CreateBatchPrediction , Amazon Machine Learning (Amazon ML) immediately returns
// and sets the BatchPrediction status to PENDING . After the BatchPrediction
// completes, Amazon ML sets the status to COMPLETED .
//
// You can poll for status updates by using the GetBatchPrediction operation and checking the Status
// parameter of the result. After the COMPLETED status appears, the results are
// available in the location specified by the OutputUri parameter.
func machinelearning_CreateBatchPrediction(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.CreateBatchPredictionInput{
		// BatchPredictionDataSourceId: *string, // Required
		// BatchPredictionId: *string, // Required
		// MLModelId: *string, // Required
		// OutputUri: *string, // Required
	}

	if len(_machinelearningBatchPredictionDataSourceId) > 0 {
		input.BatchPredictionDataSourceId = aws.String(_machinelearningBatchPredictionDataSourceId)
	}
	if len(_machinelearningBatchPredictionId) > 0 {
		input.BatchPredictionId = aws.String(_machinelearningBatchPredictionId)
	}
	if len(_machinelearningMLModelId) > 0 {
		input.MLModelId = aws.String(_machinelearningMLModelId)
	}
	if len(_machinelearningOutputUri) > 0 {
		input.OutputUri = aws.String(_machinelearningOutputUri)
	}
	if len(_machinelearningBatchPredictionName) > 0 {
		input.BatchPredictionName = aws.String(_machinelearningBatchPredictionName)
	}

	if resp, err := client.CreateBatchPrediction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a DataSource object from an [Amazon Relational Database Service] (Amazon RDS). A DataSource references data
// that can be used to perform CreateMLModel , CreateEvaluation , or
// CreateBatchPrediction operations.
//
// CreateDataSourceFromRDS is an asynchronous operation. In response to
// CreateDataSourceFromRDS , Amazon Machine Learning (Amazon ML) immediately
// returns and sets the DataSource status to PENDING . After the DataSource is
// created and ready for use, Amazon ML sets the Status parameter to COMPLETED .
// DataSource in the COMPLETED or PENDING state can be used only to perform
// >CreateMLModel >, CreateEvaluation , or CreateBatchPrediction operations.
//
// If Amazon ML cannot accept the input source, it sets the Status parameter to
// FAILED and includes an error message in the Message attribute of the
// GetDataSource operation response.
//
// [Amazon Relational Database Service]: http://aws.amazon.com/rds/
func machinelearning_CreateDataSourceFromRDS(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.CreateDataSourceFromRDSInput{
		// DataSourceId: *string, // Required
		// RDSData: *types.RDSDataSpec, // Required
		// RoleARN: *string, // Required
	}

	if len(_machinelearningDataSourceId) > 0 {
		input.DataSourceId = aws.String(_machinelearningDataSourceId)
	}
	if len(_machinelearningRDSData) > 0 {
		if err := assignInputField(input, "RDSData", _machinelearningRDSData); err != nil {
			log.Errorf("invalid --rds-data: %s", err.Error())
			return
		}
	}
	if len(_machinelearningRoleARN) > 0 {
		input.RoleARN = aws.String(_machinelearningRoleARN)
	}
	if len(_machinelearningComputeStatistics) > 0 {
		if err := assignInputField(input, "ComputeStatistics", _machinelearningComputeStatistics); err != nil {
			log.Errorf("invalid --compute-statistics: %s", err.Error())
			return
		}
	}
	if len(_machinelearningDataSourceName) > 0 {
		input.DataSourceName = aws.String(_machinelearningDataSourceName)
	}

	if resp, err := client.CreateDataSourceFromRDS(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a DataSource from a database hosted on an Amazon Redshift cluster. A
// DataSource references data that can be used to perform either CreateMLModel ,
// CreateEvaluation , or CreateBatchPrediction operations.
//
// CreateDataSourceFromRedshift is an asynchronous operation. In response to
// CreateDataSourceFromRedshift , Amazon Machine Learning (Amazon ML) immediately
// returns and sets the DataSource status to PENDING . After the DataSource is
// created and ready for use, Amazon ML sets the Status parameter to COMPLETED .
// DataSource in COMPLETED or PENDING states can be used to perform only
// CreateMLModel , CreateEvaluation , or CreateBatchPrediction operations.
//
// If Amazon ML can't accept the input source, it sets the Status parameter to
// FAILED and includes an error message in the Message attribute of the
// GetDataSource operation response.
//
// The observations should be contained in the database hosted on an Amazon
// Redshift cluster and should be specified by a SelectSqlQuery query. Amazon ML
// executes an Unload command in Amazon Redshift to transfer the result set of the
// SelectSqlQuery query to S3StagingLocation .
//
// After the DataSource has been created, it's ready for use in evaluations and
// batch predictions. If you plan to use the DataSource to train an MLModel , the
// DataSource also requires a recipe. A recipe describes how each input variable
// will be used in training an MLModel . Will the variable be included or excluded
// from training? Will the variable be manipulated; for example, will it be
// combined with another variable or will it be split apart into word combinations?
// The recipe provides answers to these questions.
//
// You can't change an existing datasource, but you can copy and modify the
// settings from an existing Amazon Redshift datasource to create a new datasource.
// To do so, call GetDataSource for an existing datasource and copy the values to
// a CreateDataSource call. Change the settings that you want to change and make
// sure that all required fields have the appropriate values.
func machinelearning_CreateDataSourceFromRedshift(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.CreateDataSourceFromRedshiftInput{
		// DataSourceId: *string, // Required
		// DataSpec: *types.RedshiftDataSpec, // Required
		// RoleARN: *string, // Required
	}

	if len(_machinelearningDataSourceId) > 0 {
		input.DataSourceId = aws.String(_machinelearningDataSourceId)
	}
	if len(_machinelearningDataSpec) > 0 {
		if err := assignInputField(input, "DataSpec", _machinelearningDataSpec); err != nil {
			log.Errorf("invalid --data-spec: %s", err.Error())
			return
		}
	}
	if len(_machinelearningRoleARN) > 0 {
		input.RoleARN = aws.String(_machinelearningRoleARN)
	}
	if len(_machinelearningComputeStatistics) > 0 {
		if err := assignInputField(input, "ComputeStatistics", _machinelearningComputeStatistics); err != nil {
			log.Errorf("invalid --compute-statistics: %s", err.Error())
			return
		}
	}
	if len(_machinelearningDataSourceName) > 0 {
		input.DataSourceName = aws.String(_machinelearningDataSourceName)
	}

	if resp, err := client.CreateDataSourceFromRedshift(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a DataSource object. A DataSource references data that can be used to
// perform CreateMLModel , CreateEvaluation , or CreateBatchPrediction operations.
//
// CreateDataSourceFromS3 is an asynchronous operation. In response to
// CreateDataSourceFromS3 , Amazon Machine Learning (Amazon ML) immediately returns
// and sets the DataSource status to PENDING . After the DataSource has been
// created and is ready for use, Amazon ML sets the Status parameter to COMPLETED .
// DataSource in the COMPLETED or PENDING state can be used to perform only
// CreateMLModel , CreateEvaluation or CreateBatchPrediction operations.
//
// If Amazon ML can't accept the input source, it sets the Status parameter to
// FAILED and includes an error message in the Message attribute of the
// GetDataSource operation response.
//
// The observation data used in a DataSource should be ready to use; that is, it
// should have a consistent structure, and missing data values should be kept to a
// minimum. The observation data must reside in one or more .csv files in an Amazon
// Simple Storage Service (Amazon S3) location, along with a schema that describes
// the data items by name and type. The same schema must be used for all of the
// data files referenced by the DataSource .
//
// After the DataSource has been created, it's ready to use in evaluations and
// batch predictions. If you plan to use the DataSource to train an MLModel , the
// DataSource also needs a recipe. A recipe describes how each input variable will
// be used in training an MLModel . Will the variable be included or excluded from
// training? Will the variable be manipulated; for example, will it be combined
// with another variable or will it be split apart into word combinations? The
// recipe provides answers to these questions.
func machinelearning_CreateDataSourceFromS3(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.CreateDataSourceFromS3Input{
		// DataSourceId: *string, // Required
		// DataSpec: *types.S3DataSpec, // Required
	}

	if len(_machinelearningDataSourceId) > 0 {
		input.DataSourceId = aws.String(_machinelearningDataSourceId)
	}
	if len(_machinelearningDataSpec) > 0 {
		if err := assignInputField(input, "DataSpec", _machinelearningDataSpec); err != nil {
			log.Errorf("invalid --data-spec: %s", err.Error())
			return
		}
	}
	if len(_machinelearningComputeStatistics) > 0 {
		if err := assignInputField(input, "ComputeStatistics", _machinelearningComputeStatistics); err != nil {
			log.Errorf("invalid --compute-statistics: %s", err.Error())
			return
		}
	}
	if len(_machinelearningDataSourceName) > 0 {
		input.DataSourceName = aws.String(_machinelearningDataSourceName)
	}

	if resp, err := client.CreateDataSourceFromS3(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Evaluation of an MLModel . An MLModel is evaluated on a set of
// observations associated to a DataSource . Like a DataSource for an MLModel , the
// DataSource for an Evaluation contains values for the Target Variable . The
// Evaluation compares the predicted result for each observation to the actual
// outcome and provides a summary so that you know how effective the MLModel
// functions on the test data. Evaluation generates a relevant performance metric,
// such as BinaryAUC, RegressionRMSE or MulticlassAvgFScore based on the
// corresponding MLModelType : BINARY , REGRESSION or MULTICLASS .
//
// CreateEvaluation is an asynchronous operation. In response to CreateEvaluation ,
// Amazon Machine Learning (Amazon ML) immediately returns and sets the evaluation
// status to PENDING . After the Evaluation is created and ready for use, Amazon
// ML sets the status to COMPLETED .
//
// You can use the GetEvaluation operation to check progress of the evaluation
// during the creation operation.
func machinelearning_CreateEvaluation(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.CreateEvaluationInput{
		// EvaluationDataSourceId: *string, // Required
		// EvaluationId: *string, // Required
		// MLModelId: *string, // Required
	}

	if len(_machinelearningEvaluationDataSourceId) > 0 {
		input.EvaluationDataSourceId = aws.String(_machinelearningEvaluationDataSourceId)
	}
	if len(_machinelearningEvaluationId) > 0 {
		input.EvaluationId = aws.String(_machinelearningEvaluationId)
	}
	if len(_machinelearningMLModelId) > 0 {
		input.MLModelId = aws.String(_machinelearningMLModelId)
	}
	if len(_machinelearningEvaluationName) > 0 {
		input.EvaluationName = aws.String(_machinelearningEvaluationName)
	}

	if resp, err := client.CreateEvaluation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new MLModel using the DataSource and the recipe as information
// sources.
//
// An MLModel is nearly immutable. Users can update only the MLModelName and the
// ScoreThreshold in an MLModel without creating a new MLModel .
//
// CreateMLModel is an asynchronous operation. In response to CreateMLModel ,
// Amazon Machine Learning (Amazon ML) immediately returns and sets the MLModel
// status to PENDING . After the MLModel has been created and ready is for use,
// Amazon ML sets the status to COMPLETED .
//
// You can use the GetMLModel operation to check the progress of the MLModel
// during the creation operation.
//
// CreateMLModel requires a DataSource with computed statistics, which can be
// created by setting ComputeStatistics to true in CreateDataSourceFromRDS ,
// CreateDataSourceFromS3 , or CreateDataSourceFromRedshift operations.
func machinelearning_CreateMLModel(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.CreateMLModelInput{
		// MLModelId: *string, // Required
		// MLModelType: types.MLModelType, // Required
		// TrainingDataSourceId: *string, // Required
	}

	if len(_machinelearningMLModelId) > 0 {
		input.MLModelId = aws.String(_machinelearningMLModelId)
	}
	if len(_machinelearningMLModelType) > 0 {
		if err := assignInputField(input, "MLModelType", _machinelearningMLModelType); err != nil {
			log.Errorf("invalid --ml-model-type: %s", err.Error())
			return
		}
	}
	if len(_machinelearningTrainingDataSourceId) > 0 {
		input.TrainingDataSourceId = aws.String(_machinelearningTrainingDataSourceId)
	}
	if len(_machinelearningMLModelName) > 0 {
		input.MLModelName = aws.String(_machinelearningMLModelName)
	}
	if len(_machinelearningParameters) > 0 {
		if err := assignInputField(input, "Parameters", _machinelearningParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_machinelearningRecipe) > 0 {
		input.Recipe = aws.String(_machinelearningRecipe)
	}
	if len(_machinelearningRecipeUri) > 0 {
		input.RecipeUri = aws.String(_machinelearningRecipeUri)
	}

	if resp, err := client.CreateMLModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a real-time endpoint for the MLModel . The endpoint contains the URI of
// the MLModel ; that is, the location to send real-time prediction requests for
// the specified MLModel .
func machinelearning_CreateRealtimeEndpoint(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.CreateRealtimeEndpointInput{
		// MLModelId: *string, // Required
	}

	if len(_machinelearningMLModelId) > 0 {
		input.MLModelId = aws.String(_machinelearningMLModelId)
	}

	if resp, err := client.CreateRealtimeEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns the DELETED status to a BatchPrediction , rendering it unusable.
// After using the DeleteBatchPrediction operation, you can use the GetBatchPrediction operation to
// verify that the status of the BatchPrediction changed to DELETED.
//
// Caution: The result of the DeleteBatchPrediction operation is irreversible.
func machinelearning_DeleteBatchPrediction(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.DeleteBatchPredictionInput{
		// BatchPredictionId: *string, // Required
	}

	if len(_machinelearningBatchPredictionId) > 0 {
		input.BatchPredictionId = aws.String(_machinelearningBatchPredictionId)
	}

	if resp, err := client.DeleteBatchPrediction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns the DELETED status to a DataSource , rendering it unusable.
// After using the DeleteDataSource operation, you can use the GetDataSource operation to
// verify that the status of the DataSource changed to DELETED.
//
// Caution: The results of the DeleteDataSource operation are irreversible.
func machinelearning_DeleteDataSource(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.DeleteDataSourceInput{
		// DataSourceId: *string, // Required
	}

	if len(_machinelearningDataSourceId) > 0 {
		input.DataSourceId = aws.String(_machinelearningDataSourceId)
	}

	if resp, err := client.DeleteDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns the DELETED status to an Evaluation , rendering it unusable.
// After invoking the DeleteEvaluation operation, you can use the GetEvaluation
// operation to verify that the status of the Evaluation changed to DELETED .
//
// Caution: The results of the DeleteEvaluation operation are irreversible.
func machinelearning_DeleteEvaluation(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.DeleteEvaluationInput{
		// EvaluationId: *string, // Required
	}

	if len(_machinelearningEvaluationId) > 0 {
		input.EvaluationId = aws.String(_machinelearningEvaluationId)
	}

	if resp, err := client.DeleteEvaluation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns the DELETED status to an MLModel , rendering it unusable.
// After using the DeleteMLModel operation, you can use the GetMLModel operation
// to verify that the status of the MLModel changed to DELETED.
//
// Caution: The result of the DeleteMLModel operation is irreversible.
func machinelearning_DeleteMLModel(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.DeleteMLModelInput{
		// MLModelId: *string, // Required
	}

	if len(_machinelearningMLModelId) > 0 {
		input.MLModelId = aws.String(_machinelearningMLModelId)
	}

	if resp, err := client.DeleteMLModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a real time endpoint of an MLModel .
func machinelearning_DeleteRealtimeEndpoint(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.DeleteRealtimeEndpointInput{
		// MLModelId: *string, // Required
	}

	if len(_machinelearningMLModelId) > 0 {
		input.MLModelId = aws.String(_machinelearningMLModelId)
	}

	if resp, err := client.DeleteRealtimeEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified tags associated with an ML object. After this operation
// is complete, you can't recover deleted tags.
//
// If you specify a tag that doesn't exist, Amazon ML ignores it.
func machinelearning_DeleteTags(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.DeleteTagsInput{
		// ResourceId: *string, // Required
		// ResourceType: types.TaggableResourceType, // Required
		// TagKeys: []string, // Required
	}

	if len(_machinelearningResourceId) > 0 {
		input.ResourceId = aws.String(_machinelearningResourceId)
	}
	if len(_machinelearningResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _machinelearningResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_machinelearningTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _machinelearningTagKeys...)
	}

	if resp, err := client.DeleteTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of BatchPrediction operations that match the search criteria in
// the request.
func machinelearning_DescribeBatchPredictions(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.DescribeBatchPredictionsInput{}

	if len(_machinelearningEQ) > 0 {
		input.EQ = aws.String(_machinelearningEQ)
	}
	if len(_machinelearningFilterVariable) > 0 {
		if err := assignInputField(input, "FilterVariable", _machinelearningFilterVariable); err != nil {
			log.Errorf("invalid --filter-variable: %s", err.Error())
			return
		}
	}
	if len(_machinelearningGE) > 0 {
		input.GE = aws.String(_machinelearningGE)
	}
	if len(_machinelearningGT) > 0 {
		input.GT = aws.String(_machinelearningGT)
	}
	if len(_machinelearningLE) > 0 {
		input.LE = aws.String(_machinelearningLE)
	}
	if len(_machinelearningLT) > 0 {
		input.LT = aws.String(_machinelearningLT)
	}
	if len(_machinelearningLimit) > 0 {
		if err := assignInputField(input, "Limit", _machinelearningLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_machinelearningNE) > 0 {
		input.NE = aws.String(_machinelearningNE)
	}
	if len(_machinelearningNextToken) > 0 {
		input.NextToken = aws.String(_machinelearningNextToken)
	}
	if len(_machinelearningPrefix) > 0 {
		input.Prefix = aws.String(_machinelearningPrefix)
	}
	if len(_machinelearningSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _machinelearningSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeBatchPredictions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*machinelearning.DescribeBatchPredictionsOutput
	p := machinelearning.NewDescribeBatchPredictionsPaginator(client, input)
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

// Returns a list of DataSource that match the search criteria in the request.
func machinelearning_DescribeDataSources(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.DescribeDataSourcesInput{}

	if len(_machinelearningEQ) > 0 {
		input.EQ = aws.String(_machinelearningEQ)
	}
	if len(_machinelearningFilterVariable) > 0 {
		if err := assignInputField(input, "FilterVariable", _machinelearningFilterVariable); err != nil {
			log.Errorf("invalid --filter-variable: %s", err.Error())
			return
		}
	}
	if len(_machinelearningGE) > 0 {
		input.GE = aws.String(_machinelearningGE)
	}
	if len(_machinelearningGT) > 0 {
		input.GT = aws.String(_machinelearningGT)
	}
	if len(_machinelearningLE) > 0 {
		input.LE = aws.String(_machinelearningLE)
	}
	if len(_machinelearningLT) > 0 {
		input.LT = aws.String(_machinelearningLT)
	}
	if len(_machinelearningLimit) > 0 {
		if err := assignInputField(input, "Limit", _machinelearningLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_machinelearningNE) > 0 {
		input.NE = aws.String(_machinelearningNE)
	}
	if len(_machinelearningNextToken) > 0 {
		input.NextToken = aws.String(_machinelearningNextToken)
	}
	if len(_machinelearningPrefix) > 0 {
		input.Prefix = aws.String(_machinelearningPrefix)
	}
	if len(_machinelearningSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _machinelearningSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDataSources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*machinelearning.DescribeDataSourcesOutput
	p := machinelearning.NewDescribeDataSourcesPaginator(client, input)
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

// Returns a list of DescribeEvaluations that match the search criteria in the
// request.
func machinelearning_DescribeEvaluations(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.DescribeEvaluationsInput{}

	if len(_machinelearningEQ) > 0 {
		input.EQ = aws.String(_machinelearningEQ)
	}
	if len(_machinelearningFilterVariable) > 0 {
		if err := assignInputField(input, "FilterVariable", _machinelearningFilterVariable); err != nil {
			log.Errorf("invalid --filter-variable: %s", err.Error())
			return
		}
	}
	if len(_machinelearningGE) > 0 {
		input.GE = aws.String(_machinelearningGE)
	}
	if len(_machinelearningGT) > 0 {
		input.GT = aws.String(_machinelearningGT)
	}
	if len(_machinelearningLE) > 0 {
		input.LE = aws.String(_machinelearningLE)
	}
	if len(_machinelearningLT) > 0 {
		input.LT = aws.String(_machinelearningLT)
	}
	if len(_machinelearningLimit) > 0 {
		if err := assignInputField(input, "Limit", _machinelearningLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_machinelearningNE) > 0 {
		input.NE = aws.String(_machinelearningNE)
	}
	if len(_machinelearningNextToken) > 0 {
		input.NextToken = aws.String(_machinelearningNextToken)
	}
	if len(_machinelearningPrefix) > 0 {
		input.Prefix = aws.String(_machinelearningPrefix)
	}
	if len(_machinelearningSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _machinelearningSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeEvaluations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*machinelearning.DescribeEvaluationsOutput
	p := machinelearning.NewDescribeEvaluationsPaginator(client, input)
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

// Returns a list of MLModel that match the search criteria in the request.
func machinelearning_DescribeMLModels(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.DescribeMLModelsInput{}

	if len(_machinelearningEQ) > 0 {
		input.EQ = aws.String(_machinelearningEQ)
	}
	if len(_machinelearningFilterVariable) > 0 {
		if err := assignInputField(input, "FilterVariable", _machinelearningFilterVariable); err != nil {
			log.Errorf("invalid --filter-variable: %s", err.Error())
			return
		}
	}
	if len(_machinelearningGE) > 0 {
		input.GE = aws.String(_machinelearningGE)
	}
	if len(_machinelearningGT) > 0 {
		input.GT = aws.String(_machinelearningGT)
	}
	if len(_machinelearningLE) > 0 {
		input.LE = aws.String(_machinelearningLE)
	}
	if len(_machinelearningLT) > 0 {
		input.LT = aws.String(_machinelearningLT)
	}
	if len(_machinelearningLimit) > 0 {
		if err := assignInputField(input, "Limit", _machinelearningLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_machinelearningNE) > 0 {
		input.NE = aws.String(_machinelearningNE)
	}
	if len(_machinelearningNextToken) > 0 {
		input.NextToken = aws.String(_machinelearningNextToken)
	}
	if len(_machinelearningPrefix) > 0 {
		input.Prefix = aws.String(_machinelearningPrefix)
	}
	if len(_machinelearningSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _machinelearningSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeMLModels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*machinelearning.DescribeMLModelsOutput
	p := machinelearning.NewDescribeMLModelsPaginator(client, input)
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

// Describes one or more of the tags for your Amazon ML object.
func machinelearning_DescribeTags(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.DescribeTagsInput{
		// ResourceId: *string, // Required
		// ResourceType: types.TaggableResourceType, // Required
	}

	if len(_machinelearningResourceId) > 0 {
		input.ResourceId = aws.String(_machinelearningResourceId)
	}
	if len(_machinelearningResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _machinelearningResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a BatchPrediction that includes detailed metadata, status, and data
// file information for a Batch Prediction request.
func machinelearning_GetBatchPrediction(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.GetBatchPredictionInput{
		// BatchPredictionId: *string, // Required
	}

	if len(_machinelearningBatchPredictionId) > 0 {
		input.BatchPredictionId = aws.String(_machinelearningBatchPredictionId)
	}

	if resp, err := client.GetBatchPrediction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a DataSource that includes metadata and data file information, as well
// as the current status of the DataSource .
//
// GetDataSource provides results in normal or verbose format. The verbose format
// adds the schema description and the list of files pointed to by the DataSource
// to the normal format.
func machinelearning_GetDataSource(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.GetDataSourceInput{
		// DataSourceId: *string, // Required
	}

	if len(_machinelearningDataSourceId) > 0 {
		input.DataSourceId = aws.String(_machinelearningDataSourceId)
	}
	if len(_machinelearningVerbose) > 0 {
		if err := assignInputField(input, "Verbose", _machinelearningVerbose); err != nil {
			log.Errorf("invalid --verbose: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an Evaluation that includes metadata as well as the current status of
// the Evaluation .
func machinelearning_GetEvaluation(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.GetEvaluationInput{
		// EvaluationId: *string, // Required
	}

	if len(_machinelearningEvaluationId) > 0 {
		input.EvaluationId = aws.String(_machinelearningEvaluationId)
	}

	if resp, err := client.GetEvaluation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an MLModel that includes detailed metadata, data source information,
// and the current status of the MLModel .
//
// GetMLModel provides results in normal or verbose format.
func machinelearning_GetMLModel(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.GetMLModelInput{
		// MLModelId: *string, // Required
	}

	if len(_machinelearningMLModelId) > 0 {
		input.MLModelId = aws.String(_machinelearningMLModelId)
	}
	if len(_machinelearningVerbose) > 0 {
		if err := assignInputField(input, "Verbose", _machinelearningVerbose); err != nil {
			log.Errorf("invalid --verbose: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetMLModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates a prediction for the observation using the specified ML Model .
// Note: Not all response parameters will be populated. Whether a response
// parameter is populated depends on the type of model requested.
func machinelearning_Predict(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.PredictInput{
		// MLModelId: *string, // Required
		// PredictEndpoint: *string, // Required
		// Record: map[string]string, // Required
	}

	if len(_machinelearningMLModelId) > 0 {
		input.MLModelId = aws.String(_machinelearningMLModelId)
	}
	if len(_machinelearningPredictEndpoint) > 0 {
		input.PredictEndpoint = aws.String(_machinelearningPredictEndpoint)
	}
	if len(_machinelearningRecord) > 0 {
		if err := assignInputField(input, "Record", _machinelearningRecord); err != nil {
			log.Errorf("invalid --record: %s", err.Error())
			return
		}
	}

	if resp, err := client.Predict(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the BatchPredictionName of a BatchPrediction .
// You can use the GetBatchPrediction operation to view the contents of the
// updated data element.
func machinelearning_UpdateBatchPrediction(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.UpdateBatchPredictionInput{
		// BatchPredictionId: *string, // Required
		// BatchPredictionName: *string, // Required
	}

	if len(_machinelearningBatchPredictionId) > 0 {
		input.BatchPredictionId = aws.String(_machinelearningBatchPredictionId)
	}
	if len(_machinelearningBatchPredictionName) > 0 {
		input.BatchPredictionName = aws.String(_machinelearningBatchPredictionName)
	}

	if resp, err := client.UpdateBatchPrediction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the DataSourceName of a DataSource .
// You can use the GetDataSource operation to view the contents of the updated
// data element.
func machinelearning_UpdateDataSource(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.UpdateDataSourceInput{
		// DataSourceId: *string, // Required
		// DataSourceName: *string, // Required
	}

	if len(_machinelearningDataSourceId) > 0 {
		input.DataSourceId = aws.String(_machinelearningDataSourceId)
	}
	if len(_machinelearningDataSourceName) > 0 {
		input.DataSourceName = aws.String(_machinelearningDataSourceName)
	}

	if resp, err := client.UpdateDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the EvaluationName of an Evaluation .
// You can use the GetEvaluation operation to view the contents of the updated
// data element.
func machinelearning_UpdateEvaluation(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.UpdateEvaluationInput{
		// EvaluationId: *string, // Required
		// EvaluationName: *string, // Required
	}

	if len(_machinelearningEvaluationId) > 0 {
		input.EvaluationId = aws.String(_machinelearningEvaluationId)
	}
	if len(_machinelearningEvaluationName) > 0 {
		input.EvaluationName = aws.String(_machinelearningEvaluationName)
	}

	if resp, err := client.UpdateEvaluation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the MLModelName and the ScoreThreshold of an MLModel .
// You can use the GetMLModel operation to view the contents of the updated data
// element.
func machinelearning_UpdateMLModel(cfg aws.Config, client *machinelearning.Client) {
	input := &machinelearning.UpdateMLModelInput{
		// MLModelId: *string, // Required
	}

	if len(_machinelearningMLModelId) > 0 {
		input.MLModelId = aws.String(_machinelearningMLModelId)
	}
	if len(_machinelearningMLModelName) > 0 {
		input.MLModelName = aws.String(_machinelearningMLModelName)
	}
	if len(_machinelearningScoreThreshold) > 0 {
		if err := assignInputField(input, "ScoreThreshold", _machinelearningScoreThreshold); err != nil {
			log.Errorf("invalid --score-threshold: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMLModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_machinelearningCmd)
	_machinelearningCmd.Flags().SortFlags = false

	_machinelearningCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_machinelearningCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_machinelearningCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_machinelearningCmd.Flags().StringVarP(&_machinelearningBatchPredictionDataSourceId, "batch-prediction-data-source-id", "", "", "Batch Prediction Data Source ID")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningBatchPredictionId, "batch-prediction-id", "", "", "Batch Prediction ID")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningBatchPredictionName, "batch-prediction-name", "", "", "Batch Prediction Name")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningComputeStatistics, "compute-statistics", "", "", "Compute Statistics")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningDataSourceId, "data-source-id", "", "", "Data Source ID")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningDataSourceName, "data-source-name", "", "", "Data Source Name")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningDataSpec, "data-spec", "", "", "Data Spec")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningEQ, "eq", "", "", "Eq")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningEvaluationDataSourceId, "evaluation-data-source-id", "", "", "Evaluation Data Source ID")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningEvaluationId, "evaluation-id", "", "", "Evaluation ID")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningEvaluationName, "evaluation-name", "", "", "Evaluation Name")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningFilterVariable, "filter-variable", "", "", "Filter Variable")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningGE, "ge", "", "", "Ge")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningGT, "gt", "", "", "Gt")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningLE, "le", "", "", "Le")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningLimit, "limit", "", "", "Limit")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningLT, "lt", "", "", "Lt")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningMLModelId, "ml-model-id", "", "", "Ml Model ID")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningMLModelName, "ml-model-name", "", "", "Ml Model Name")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningMLModelType, "ml-model-type", "", "", "Ml Model Type")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningNE, "ne", "", "", "Ne")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningNextToken, "next-token", "", "", "Next Token")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningOutputUri, "output-uri", "", "", "Output URI")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningParameters, "parameters", "", "", "Parameters")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningPredictEndpoint, "predict-endpoint", "", "", "Predict Endpoint")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningPrefix, "prefix", "", "", "Prefix")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningRDSData, "rds-data", "", "", "RDS Data")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningRecipe, "recipe", "", "", "Recipe")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningRecipeUri, "recipe-uri", "", "", "Recipe URI")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningRecord, "record", "", "", "Record")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningResourceId, "resource-id", "", "", "Resource ID")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningResourceType, "resource-type", "", "", "Resource Type")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningRoleARN, "role-arn", "", "", "Role ARN")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningScoreThreshold, "score-threshold", "", "", "Score Threshold")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningSortOrder, "sort-order", "", "", "Sort Order")
	_machinelearningCmd.Flags().StringSliceVarP(&_machinelearningTagKeys, "tag-keys", "", nil, "Tag Keys")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningTags, "tags", "", "", "Tags")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningTrainingDataSourceId, "training-data-source-id", "", "", "Training Data Source ID")
	_machinelearningCmd.Flags().StringVarP(&_machinelearningVerbose, "verbose", "", "", "Verbose")

	_machinelearningCmd.Flags().BoolVarP(&_machinelearningAddTags, "add-tags", "", false, "Add Tags")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningCreateBatchPrediction, "create-batch-prediction", "", false, "Create Batch Prediction")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningCreateDataSourceFromRDS, "create-data-source-from-rds", "", false, "Create Data Source From RDS")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningCreateDataSourceFromRedshift, "create-data-source-from-redshift", "", false, "Create Data Source From Redshift")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningCreateDataSourceFromS3, "create-data-source-from-s3", "", false, "Create Data Source From S3")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningCreateEvaluation, "create-evaluation", "", false, "Create Evaluation")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningCreateMLModel, "create-ml-model", "", false, "Create Ml Model")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningCreateRealtimeEndpoint, "create-realtime-endpoint", "", false, "Create Realtime Endpoint")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningDeleteBatchPrediction, "delete-batch-prediction", "", false, "Delete Batch Prediction")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningDeleteDataSource, "delete-data-source", "", false, "Delete Data Source")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningDeleteEvaluation, "delete-evaluation", "", false, "Delete Evaluation")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningDeleteMLModel, "delete-ml-model", "", false, "Delete Ml Model")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningDeleteRealtimeEndpoint, "delete-realtime-endpoint", "", false, "Delete Realtime Endpoint")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningDeleteTags, "delete-tags", "", false, "Delete Tags")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningDescribeBatchPredictions, "describe-batch-predictions", "", false, "Describe Batch Predictions")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningDescribeDataSources, "describe-data-sources", "", false, "Describe Data Sources")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningDescribeEvaluations, "describe-evaluations", "", false, "Describe Evaluations")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningDescribeMLModels, "describe-ml-models", "", false, "Describe Ml Models")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningDescribeTags, "describe-tags", "", false, "Describe Tags")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningGetBatchPrediction, "get-batch-prediction", "", false, "Get Batch Prediction")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningGetDataSource, "get-data-source", "", false, "Get Data Source")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningGetEvaluation, "get-evaluation", "", false, "Get Evaluation")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningGetMLModel, "get-ml-model", "", false, "Get Ml Model")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningPredict, "predict", "", false, "Predict")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningUpdateBatchPrediction, "update-batch-prediction", "", false, "Update Batch Prediction")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningUpdateDataSource, "update-data-source", "", false, "Update Data Source")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningUpdateEvaluation, "update-evaluation", "", false, "Update Evaluation")
	_machinelearningCmd.Flags().BoolVarP(&_machinelearningUpdateMLModel, "update-ml-model", "", false, "Update Ml Model")

}
