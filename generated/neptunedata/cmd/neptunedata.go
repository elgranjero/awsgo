package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptunedata"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// neptunedataCmd represents the neptunedata command
var _neptunedataCmd = &cobra.Command{
	Use:   "neptunedata",
	Short: "AWS neptunedata CLI",
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
		client := neptunedata.NewFromConfig(cfg)
		if _neptunedataCancelGremlinQuery {
			neptunedata_CancelGremlinQuery(cfg, client)
			return
		}
		if _neptunedataCancelLoaderJob {
			neptunedata_CancelLoaderJob(cfg, client)
			return
		}
		if _neptunedataCancelMLDataProcessingJob {
			neptunedata_CancelMLDataProcessingJob(cfg, client)
			return
		}
		if _neptunedataCancelMLModelTrainingJob {
			neptunedata_CancelMLModelTrainingJob(cfg, client)
			return
		}
		if _neptunedataCancelMLModelTransformJob {
			neptunedata_CancelMLModelTransformJob(cfg, client)
			return
		}
		if _neptunedataCancelOpenCypherQuery {
			neptunedata_CancelOpenCypherQuery(cfg, client)
			return
		}
		if _neptunedataCreateMLEndpoint {
			neptunedata_CreateMLEndpoint(cfg, client)
			return
		}
		if _neptunedataDeleteMLEndpoint {
			neptunedata_DeleteMLEndpoint(cfg, client)
			return
		}
		if _neptunedataDeletePropertygraphStatistics {
			neptunedata_DeletePropertygraphStatistics(cfg, client)
			return
		}
		if _neptunedataDeleteSparqlStatistics {
			neptunedata_DeleteSparqlStatistics(cfg, client)
			return
		}
		if _neptunedataExecuteFastReset {
			neptunedata_ExecuteFastReset(cfg, client)
			return
		}
		if _neptunedataExecuteGremlinExplainQuery {
			neptunedata_ExecuteGremlinExplainQuery(cfg, client)
			return
		}
		if _neptunedataExecuteGremlinProfileQuery {
			neptunedata_ExecuteGremlinProfileQuery(cfg, client)
			return
		}
		if _neptunedataExecuteGremlinQuery {
			neptunedata_ExecuteGremlinQuery(cfg, client)
			return
		}
		if _neptunedataExecuteOpenCypherExplainQuery {
			neptunedata_ExecuteOpenCypherExplainQuery(cfg, client)
			return
		}
		if _neptunedataExecuteOpenCypherQuery {
			neptunedata_ExecuteOpenCypherQuery(cfg, client)
			return
		}
		if _neptunedataGetEngineStatus {
			neptunedata_GetEngineStatus(cfg, client)
			return
		}
		if _neptunedataGetGremlinQueryStatus {
			neptunedata_GetGremlinQueryStatus(cfg, client)
			return
		}
		if _neptunedataGetLoaderJobStatus {
			neptunedata_GetLoaderJobStatus(cfg, client)
			return
		}
		if _neptunedataGetMLDataProcessingJob {
			neptunedata_GetMLDataProcessingJob(cfg, client)
			return
		}
		if _neptunedataGetMLEndpoint {
			neptunedata_GetMLEndpoint(cfg, client)
			return
		}
		if _neptunedataGetMLModelTrainingJob {
			neptunedata_GetMLModelTrainingJob(cfg, client)
			return
		}
		if _neptunedataGetMLModelTransformJob {
			neptunedata_GetMLModelTransformJob(cfg, client)
			return
		}
		if _neptunedataGetOpenCypherQueryStatus {
			neptunedata_GetOpenCypherQueryStatus(cfg, client)
			return
		}
		if _neptunedataGetPropertygraphStatistics {
			neptunedata_GetPropertygraphStatistics(cfg, client)
			return
		}
		if _neptunedataGetPropertygraphStream {
			neptunedata_GetPropertygraphStream(cfg, client)
			return
		}
		if _neptunedataGetPropertygraphSummary {
			neptunedata_GetPropertygraphSummary(cfg, client)
			return
		}
		if _neptunedataGetRDFGraphSummary {
			neptunedata_GetRDFGraphSummary(cfg, client)
			return
		}
		if _neptunedataGetSparqlStatistics {
			neptunedata_GetSparqlStatistics(cfg, client)
			return
		}
		if _neptunedataGetSparqlStream {
			neptunedata_GetSparqlStream(cfg, client)
			return
		}
		if _neptunedataListGremlinQueries {
			neptunedata_ListGremlinQueries(cfg, client)
			return
		}
		if _neptunedataListLoaderJobs {
			neptunedata_ListLoaderJobs(cfg, client)
			return
		}
		if _neptunedataListMLDataProcessingJobs {
			neptunedata_ListMLDataProcessingJobs(cfg, client)
			return
		}
		if _neptunedataListMLEndpoints {
			neptunedata_ListMLEndpoints(cfg, client)
			return
		}
		if _neptunedataListMLModelTrainingJobs {
			neptunedata_ListMLModelTrainingJobs(cfg, client)
			return
		}
		if _neptunedataListMLModelTransformJobs {
			neptunedata_ListMLModelTransformJobs(cfg, client)
			return
		}
		if _neptunedataListOpenCypherQueries {
			neptunedata_ListOpenCypherQueries(cfg, client)
			return
		}
		if _neptunedataManagePropertygraphStatistics {
			neptunedata_ManagePropertygraphStatistics(cfg, client)
			return
		}
		if _neptunedataManageSparqlStatistics {
			neptunedata_ManageSparqlStatistics(cfg, client)
			return
		}
		if _neptunedataStartLoaderJob {
			neptunedata_StartLoaderJob(cfg, client)
			return
		}
		if _neptunedataStartMLDataProcessingJob {
			neptunedata_StartMLDataProcessingJob(cfg, client)
			return
		}
		if _neptunedataStartMLModelTrainingJob {
			neptunedata_StartMLModelTrainingJob(cfg, client)
			return
		}
		if _neptunedataStartMLModelTransformJob {
			neptunedata_StartMLModelTransformJob(cfg, client)
			return
		}

	},
}

var (
	_neptunedataCancelGremlinQuery            bool
	_neptunedataCancelLoaderJob               bool
	_neptunedataCancelMLDataProcessingJob     bool
	_neptunedataCancelMLModelTrainingJob      bool
	_neptunedataCancelMLModelTransformJob     bool
	_neptunedataCancelOpenCypherQuery         bool
	_neptunedataCreateMLEndpoint              bool
	_neptunedataDeleteMLEndpoint              bool
	_neptunedataDeletePropertygraphStatistics bool
	_neptunedataDeleteSparqlStatistics        bool
	_neptunedataExecuteFastReset              bool
	_neptunedataExecuteGremlinExplainQuery    bool
	_neptunedataExecuteGremlinProfileQuery    bool
	_neptunedataExecuteGremlinQuery           bool
	_neptunedataExecuteOpenCypherExplainQuery bool
	_neptunedataExecuteOpenCypherQuery        bool
	_neptunedataGetEngineStatus               bool
	_neptunedataGetGremlinQueryStatus         bool
	_neptunedataGetLoaderJobStatus            bool
	_neptunedataGetMLDataProcessingJob        bool
	_neptunedataGetMLEndpoint                 bool
	_neptunedataGetMLModelTrainingJob         bool
	_neptunedataGetMLModelTransformJob        bool
	_neptunedataGetOpenCypherQueryStatus      bool
	_neptunedataGetPropertygraphStatistics    bool
	_neptunedataGetPropertygraphStream        bool
	_neptunedataGetPropertygraphSummary       bool
	_neptunedataGetRDFGraphSummary            bool
	_neptunedataGetSparqlStatistics           bool
	_neptunedataGetSparqlStream               bool
	_neptunedataListGremlinQueries            bool
	_neptunedataListLoaderJobs                bool
	_neptunedataListMLDataProcessingJobs      bool
	_neptunedataListMLEndpoints               bool
	_neptunedataListMLModelTrainingJobs       bool
	_neptunedataListMLModelTransformJobs      bool
	_neptunedataListOpenCypherQueries         bool
	_neptunedataManagePropertygraphStatistics bool
	_neptunedataManageSparqlStatistics        bool
	_neptunedataStartLoaderJob                bool
	_neptunedataStartMLDataProcessingJob      bool
	_neptunedataStartMLModelTrainingJob       bool
	_neptunedataStartMLModelTransformJob      bool

	_neptunedataAction                               string
	_neptunedataBaseProcessingInstanceType           string
	_neptunedataBaseProcessingInstanceVolumeSizeInGB string
	_neptunedataChop                                 string
	_neptunedataClean                                string
	_neptunedataCommitNum                            string
	_neptunedataConfigFileName                       string
	_neptunedataCustomModelTrainingParameters        string
	_neptunedataCustomModelTransformParameters       string
	_neptunedataDataProcessingJobId                  string
	_neptunedataDependencies                         []string
	_neptunedataDetails                              string
	_neptunedataEdgeOnlyLoad                         string
	_neptunedataEnableManagedSpotTraining            string
	_neptunedataEncoding                             string
	_neptunedataErrors                               string
	_neptunedataErrorsPerPage                        string
	_neptunedataExplainMode                          string
	_neptunedataFailOnError                          string
	_neptunedataFormat                               string
	_neptunedataGremlinQuery                         string
	_neptunedataIamRoleArn                           string
	_neptunedataId                                   string
	_neptunedataIncludeQueuedLoads                   string
	_neptunedataIncludeWaiting                       string
	_neptunedataIndexOps                             string
	_neptunedataInputDataS3Location                  string
	_neptunedataInstanceCount                        string
	_neptunedataInstanceType                         string
	_neptunedataIteratorType                         string
	_neptunedataLimit                                string
	_neptunedataLoadId                               string
	_neptunedataMaxHPONumberOfTrainingJobs           string
	_neptunedataMaxHPOParallelTrainingJobs           string
	_neptunedataMaxItems                             string
	_neptunedataMlModelTrainingJobId                 string
	_neptunedataMlModelTransformJobId                string
	_neptunedataMode                                 string
	_neptunedataModelName                            string
	_neptunedataModelTransformOutputS3Location       string
	_neptunedataModelType                            string
	_neptunedataNeptuneIamRoleArn                    string
	_neptunedataOpNum                                string
	_neptunedataOpenCypherQuery                      string
	_neptunedataPage                                 string
	_neptunedataParallelism                          string
	_neptunedataParameters                           string
	_neptunedataParserConfiguration                  string
	_neptunedataPreviousDataProcessingJobId          string
	_neptunedataPreviousModelTrainingJobId           string
	_neptunedataProcessedDataS3Location              string
	_neptunedataProcessingInstanceType               string
	_neptunedataProcessingInstanceVolumeSizeInGB     string
	_neptunedataProcessingTimeOutInSeconds           string
	_neptunedataQueryId                              string
	_neptunedataQueueRequest                         string
	_neptunedataResults                              string
	_neptunedataS3BucketRegion                       string
	_neptunedataS3OutputEncryptionKMSKey             string
	_neptunedataSagemakerIamRoleArn                  string
	_neptunedataSecurityGroupIds                     []string
	_neptunedataSerializer                           string
	_neptunedataSilent                               string
	_neptunedataSource                               string
	_neptunedataSubnets                              []string
	_neptunedataToken                                string
	_neptunedataTrainModelS3Location                 string
	_neptunedataTrainingInstanceType                 string
	_neptunedataTrainingInstanceVolumeSizeInGB       string
	_neptunedataTrainingJobName                      string
	_neptunedataTrainingTimeOutInSeconds             string
	_neptunedataUpdate                               string
	_neptunedataUpdateSingleCardinalityProperties    string
	_neptunedataUserProvidedEdgeIds                  string
	_neptunedataVolumeEncryptionKMSKey               string
)

// Cancels a Gremlin query. See [Gremlin query cancellation] for more information.
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:CancelQuery]IAM action in that cluster.
//
// [neptune-db:CancelQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#cancelquery
// [Gremlin query cancellation]: https://docs.aws.amazon.com/neptune/latest/userguide/gremlin-api-status-cancel.html
func neptunedata_CancelGremlinQuery(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.CancelGremlinQueryInput{
		// QueryId: *string, // Required
	}

	if len(_neptunedataQueryId) > 0 {
		input.QueryId = aws.String(_neptunedataQueryId)
	}

	if resp, err := client.CancelGremlinQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a specified load job. This is an HTTP DELETE request. See [Neptune Loader Get-Status API] for more
// information.
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:CancelLoaderJob]IAM action in that cluster..
//
// [neptune-db:CancelLoaderJob]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#cancelloaderjob
// [Neptune Loader Get-Status API]: https://docs.aws.amazon.com/neptune/latest/userguide/load-api-reference-status.htm
func neptunedata_CancelLoaderJob(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.CancelLoaderJobInput{
		// LoadId: *string, // Required
	}

	if len(_neptunedataLoadId) > 0 {
		input.LoadId = aws.String(_neptunedataLoadId)
	}

	if resp, err := client.CancelLoaderJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a Neptune ML data processing job. See [The dataprocessing command]dataprocessing .
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:CancelMLDataProcessingJob]IAM action in that cluster.
//
// [neptune-db:CancelMLDataProcessingJob]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#cancelmldataprocessingjob
// [The dataprocessing command]: https://docs.aws.amazon.com/neptune/latest/userguide/machine-learning-api-dataprocessing.html
func neptunedata_CancelMLDataProcessingJob(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.CancelMLDataProcessingJobInput{
		// Id: *string, // Required
	}

	if len(_neptunedataId) > 0 {
		input.Id = aws.String(_neptunedataId)
	}
	if len(_neptunedataClean) > 0 {
		if err := assignInputField(input, "Clean", _neptunedataClean); err != nil {
			log.Errorf("invalid --clean: %s", err.Error())
			return
		}
	}
	if len(_neptunedataNeptuneIamRoleArn) > 0 {
		input.NeptuneIamRoleArn = aws.String(_neptunedataNeptuneIamRoleArn)
	}

	if resp, err := client.CancelMLDataProcessingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a Neptune ML model training job. See [Model training using the modeltraining command]modeltraining .
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:CancelMLModelTrainingJob]IAM action in that cluster.
//
// [neptune-db:CancelMLModelTrainingJob]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#cancelmlmodeltrainingjob
// [Model training using the modeltraining command]: https://docs.aws.amazon.com/neptune/latest/userguide/machine-learning-api-modeltraining.html
func neptunedata_CancelMLModelTrainingJob(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.CancelMLModelTrainingJobInput{
		// Id: *string, // Required
	}

	if len(_neptunedataId) > 0 {
		input.Id = aws.String(_neptunedataId)
	}
	if len(_neptunedataClean) > 0 {
		if err := assignInputField(input, "Clean", _neptunedataClean); err != nil {
			log.Errorf("invalid --clean: %s", err.Error())
			return
		}
	}
	if len(_neptunedataNeptuneIamRoleArn) > 0 {
		input.NeptuneIamRoleArn = aws.String(_neptunedataNeptuneIamRoleArn)
	}

	if resp, err := client.CancelMLModelTrainingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a specified model transform job. See [Use a trained model to generate new model artifacts].
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:CancelMLModelTransformJob]IAM action in that cluster.
//
// [neptune-db:CancelMLModelTransformJob]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#cancelmlmodeltransformjob
// [Use a trained model to generate new model artifacts]: https://docs.aws.amazon.com/neptune/latest/userguide/machine-learning-model-transform.html
func neptunedata_CancelMLModelTransformJob(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.CancelMLModelTransformJobInput{
		// Id: *string, // Required
	}

	if len(_neptunedataId) > 0 {
		input.Id = aws.String(_neptunedataId)
	}
	if len(_neptunedataClean) > 0 {
		if err := assignInputField(input, "Clean", _neptunedataClean); err != nil {
			log.Errorf("invalid --clean: %s", err.Error())
			return
		}
	}
	if len(_neptunedataNeptuneIamRoleArn) > 0 {
		input.NeptuneIamRoleArn = aws.String(_neptunedataNeptuneIamRoleArn)
	}

	if resp, err := client.CancelMLModelTransformJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a specified openCypher query. See [Neptune openCypher status endpoint] for more information.
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:CancelQuery]IAM action in that cluster.
//
// [neptune-db:CancelQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#cancelquery
// [Neptune openCypher status endpoint]: https://docs.aws.amazon.com/neptune/latest/userguide/access-graph-opencypher-status.html
func neptunedata_CancelOpenCypherQuery(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.CancelOpenCypherQueryInput{
		// QueryId: *string, // Required
	}

	if len(_neptunedataQueryId) > 0 {
		input.QueryId = aws.String(_neptunedataQueryId)
	}
	if len(_neptunedataSilent) > 0 {
		if err := assignInputField(input, "Silent", _neptunedataSilent); err != nil {
			log.Errorf("invalid --silent: %s", err.Error())
			return
		}
	}

	if resp, err := client.CancelOpenCypherQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Neptune ML inference endpoint that lets you query one specific
// model that the model-training process constructed. See [Managing inference endpoints using the endpoints command].
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:CreateMLEndpoint]IAM action in that cluster.
//
// [Managing inference endpoints using the endpoints command]: https://docs.aws.amazon.com/neptune/latest/userguide/machine-learning-api-endpoints.html
// [neptune-db:CreateMLEndpoint]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#createmlendpoint
func neptunedata_CreateMLEndpoint(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.CreateMLEndpointInput{}

	if len(_neptunedataId) > 0 {
		input.Id = aws.String(_neptunedataId)
	}
	if len(_neptunedataInstanceCount) > 0 {
		if err := assignInputField(input, "InstanceCount", _neptunedataInstanceCount); err != nil {
			log.Errorf("invalid --instance-count: %s", err.Error())
			return
		}
	}
	if len(_neptunedataInstanceType) > 0 {
		input.InstanceType = aws.String(_neptunedataInstanceType)
	}
	if len(_neptunedataMlModelTrainingJobId) > 0 {
		input.MlModelTrainingJobId = aws.String(_neptunedataMlModelTrainingJobId)
	}
	if len(_neptunedataMlModelTransformJobId) > 0 {
		input.MlModelTransformJobId = aws.String(_neptunedataMlModelTransformJobId)
	}
	if len(_neptunedataModelName) > 0 {
		input.ModelName = aws.String(_neptunedataModelName)
	}
	if len(_neptunedataNeptuneIamRoleArn) > 0 {
		input.NeptuneIamRoleArn = aws.String(_neptunedataNeptuneIamRoleArn)
	}
	if len(_neptunedataUpdate) > 0 {
		if err := assignInputField(input, "Update", _neptunedataUpdate); err != nil {
			log.Errorf("invalid --update: %s", err.Error())
			return
		}
	}
	if len(_neptunedataVolumeEncryptionKMSKey) > 0 {
		input.VolumeEncryptionKMSKey = aws.String(_neptunedataVolumeEncryptionKMSKey)
	}

	if resp, err := client.CreateMLEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels the creation of a Neptune ML inference endpoint. See [Managing inference endpoints using the endpoints command].
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:DeleteMLEndpoint]IAM action in that cluster.
//
// [neptune-db:DeleteMLEndpoint]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#deletemlendpoint
// [Managing inference endpoints using the endpoints command]: https://docs.aws.amazon.com/neptune/latest/userguide/machine-learning-api-endpoints.html
func neptunedata_DeleteMLEndpoint(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.DeleteMLEndpointInput{
		// Id: *string, // Required
	}

	if len(_neptunedataId) > 0 {
		input.Id = aws.String(_neptunedataId)
	}
	if len(_neptunedataClean) > 0 {
		if err := assignInputField(input, "Clean", _neptunedataClean); err != nil {
			log.Errorf("invalid --clean: %s", err.Error())
			return
		}
	}
	if len(_neptunedataNeptuneIamRoleArn) > 0 {
		input.NeptuneIamRoleArn = aws.String(_neptunedataNeptuneIamRoleArn)
	}

	if resp, err := client.DeleteMLEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes statistics for Gremlin and openCypher (property graph) data.
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:DeleteStatistics]IAM action in that cluster.
//
// [neptune-db:DeleteStatistics]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#deletestatistics
func neptunedata_DeletePropertygraphStatistics(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.DeletePropertygraphStatisticsInput{}

	if resp, err := client.DeletePropertygraphStatistics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes SPARQL statistics
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:DeleteStatistics]IAM action in that cluster.
//
// [neptune-db:DeleteStatistics]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#deletestatistics
func neptunedata_DeleteSparqlStatistics(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.DeleteSparqlStatisticsInput{}

	if resp, err := client.DeleteSparqlStatistics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The fast reset REST API lets you reset a Neptune graph quicky and easily,
// removing all of its data.
//
// Neptune fast reset is a two-step process. First you call ExecuteFastReset with
// action set to initiateDatabaseReset . This returns a UUID token which you then
// include when calling ExecuteFastReset again with action set to
// performDatabaseReset . See [Empty an Amazon Neptune DB cluster using the fast reset API].
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:ResetDatabase]IAM action in that cluster.
//
// [Empty an Amazon Neptune DB cluster using the fast reset API]: https://docs.aws.amazon.com/neptune/latest/userguide/manage-console-fast-reset.html
// [neptune-db:ResetDatabase]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#resetdatabase
func neptunedata_ExecuteFastReset(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.ExecuteFastResetInput{
		// Action: types.Action, // Required
	}

	if len(_neptunedataAction) > 0 {
		if err := assignInputField(input, "Action", _neptunedataAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_neptunedataToken) > 0 {
		input.Token = aws.String(_neptunedataToken)
	}

	if resp, err := client.ExecuteFastReset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Executes a Gremlin Explain query.
// Amazon Neptune has added a Gremlin feature named explain that provides is a
// self-service tool for understanding the execution approach being taken by the
// Neptune engine for the query. You invoke it by adding an explain parameter to
// an HTTP call that submits a Gremlin query.
//
// The explain feature provides information about the logical structure of query
// execution plans. You can use this information to identify potential evaluation
// and execution bottlenecks and to tune your query, as explained in [Tuning Gremlin queries]. You can
// also use query hints to improve query execution plans.
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows one of the following IAM actions in that cluster, depending on the
// query:
//
// [neptune-db:ReadDataViaQuery]
//
// [neptune-db:WriteDataViaQuery]
//
// [neptune-db:DeleteDataViaQuery]
//
// Note that the [neptune-db:QueryLanguage:Gremlin] IAM condition key can be used in the policy document to restrict
// the use of Gremlin queries (see [Condition keys available in Neptune IAM data-access policy statements]).
//
// [neptune-db:DeleteDataViaQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#deletedataviaquery
// [Condition keys available in Neptune IAM data-access policy statements]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html
// [neptune-db:ReadDataViaQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#readdataviaquery
// [Tuning Gremlin queries]: https://docs.aws.amazon.com/neptune/latest/userguide/gremlin-traversal-tuning.html
// [neptune-db:WriteDataViaQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#writedataviaquery
// [neptune-db:QueryLanguage:Gremlin]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys
func neptunedata_ExecuteGremlinExplainQuery(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.ExecuteGremlinExplainQueryInput{
		// GremlinQuery: *string, // Required
	}

	if len(_neptunedataGremlinQuery) > 0 {
		input.GremlinQuery = aws.String(_neptunedataGremlinQuery)
	}

	if resp, err := client.ExecuteGremlinExplainQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Executes a Gremlin Profile query, which runs a specified traversal, collects
// various metrics about the run, and produces a profile report as output. See [Gremlin profile API in Neptune]for
// details.
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:ReadDataViaQuery]IAM action in that cluster.
//
// Note that the [neptune-db:QueryLanguage:Gremlin] IAM condition key can be used in the policy document to restrict
// the use of Gremlin queries (see [Condition keys available in Neptune IAM data-access policy statements]).
//
// [Gremlin profile API in Neptune]: https://docs.aws.amazon.com/neptune/latest/userguide/gremlin-profile-api.html
// [Condition keys available in Neptune IAM data-access policy statements]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html
// [neptune-db:ReadDataViaQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#readdataviaquery
// [neptune-db:QueryLanguage:Gremlin]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys
func neptunedata_ExecuteGremlinProfileQuery(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.ExecuteGremlinProfileQueryInput{
		// GremlinQuery: *string, // Required
	}

	if len(_neptunedataGremlinQuery) > 0 {
		input.GremlinQuery = aws.String(_neptunedataGremlinQuery)
	}
	if len(_neptunedataChop) > 0 {
		if err := assignInputField(input, "Chop", _neptunedataChop); err != nil {
			log.Errorf("invalid --chop: %s", err.Error())
			return
		}
	}
	if len(_neptunedataIndexOps) > 0 {
		if err := assignInputField(input, "IndexOps", _neptunedataIndexOps); err != nil {
			log.Errorf("invalid --index-ops: %s", err.Error())
			return
		}
	}
	if len(_neptunedataResults) > 0 {
		if err := assignInputField(input, "Results", _neptunedataResults); err != nil {
			log.Errorf("invalid --results: %s", err.Error())
			return
		}
	}
	if len(_neptunedataSerializer) > 0 {
		input.Serializer = aws.String(_neptunedataSerializer)
	}

	if resp, err := client.ExecuteGremlinProfileQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This commands executes a Gremlin query. Amazon Neptune is compatible with
// Apache TinkerPop3 and Gremlin, so you can use the Gremlin traversal language to
// query the graph, as described under [The Graph]in the Apache TinkerPop3 documentation.
// More details can also be found in [Accessing a Neptune graph with Gremlin].
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that enables one of the following IAM actions in that cluster, depending on the
// query:
//
// [neptune-db:ReadDataViaQuery]
//
// [neptune-db:WriteDataViaQuery]
//
// [neptune-db:DeleteDataViaQuery]
//
// Note that the [neptune-db:QueryLanguage:Gremlin] IAM condition key can be used in the policy document to restrict
// the use of Gremlin queries (see [Condition keys available in Neptune IAM data-access policy statements]).
//
// [neptune-db:DeleteDataViaQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#deletedataviaquery
// [Condition keys available in Neptune IAM data-access policy statements]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html
// [The Graph]: https://tinkerpop.apache.org/docs/current/reference/#graph
// [Accessing a Neptune graph with Gremlin]: https://docs.aws.amazon.com/neptune/latest/userguide/access-graph-gremlin.html
// [neptune-db:ReadDataViaQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#readdataviaquery
// [neptune-db:WriteDataViaQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#writedataviaquery
// [neptune-db:QueryLanguage:Gremlin]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys
func neptunedata_ExecuteGremlinQuery(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.ExecuteGremlinQueryInput{
		// GremlinQuery: *string, // Required
	}

	if len(_neptunedataGremlinQuery) > 0 {
		input.GremlinQuery = aws.String(_neptunedataGremlinQuery)
	}
	if len(_neptunedataSerializer) > 0 {
		input.Serializer = aws.String(_neptunedataSerializer)
	}

	if resp, err := client.ExecuteGremlinQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Executes an openCypher explain request. See [The openCypher explain feature] for more information.
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:ReadDataViaQuery]IAM action in that cluster.
//
// Note that the [neptune-db:QueryLanguage:OpenCypher] IAM condition key can be used in the policy document to restrict
// the use of openCypher queries (see [Condition keys available in Neptune IAM data-access policy statements]).
//
// [Condition keys available in Neptune IAM data-access policy statements]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html
// [neptune-db:ReadDataViaQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#readdataviaquery
// [neptune-db:QueryLanguage:OpenCypher]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys
// [The openCypher explain feature]: https://docs.aws.amazon.com/neptune/latest/userguide/access-graph-opencypher-explain.html
func neptunedata_ExecuteOpenCypherExplainQuery(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.ExecuteOpenCypherExplainQueryInput{
		// ExplainMode: types.OpenCypherExplainMode, // Required
		// OpenCypherQuery: *string, // Required
	}

	if len(_neptunedataExplainMode) > 0 {
		if err := assignInputField(input, "ExplainMode", _neptunedataExplainMode); err != nil {
			log.Errorf("invalid --explain-mode: %s", err.Error())
			return
		}
	}
	if len(_neptunedataOpenCypherQuery) > 0 {
		input.OpenCypherQuery = aws.String(_neptunedataOpenCypherQuery)
	}
	if len(_neptunedataParameters) > 0 {
		input.Parameters = aws.String(_neptunedataParameters)
	}

	if resp, err := client.ExecuteOpenCypherExplainQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Executes an openCypher query. See [Accessing the Neptune Graph with openCypher] for more information.
// Neptune supports building graph applications using openCypher, which is
// currently one of the most popular query languages among developers working with
// graph databases. Developers, business analysts, and data scientists like
// openCypher's declarative, SQL-inspired syntax because it provides a familiar
// structure in which to querying property graphs.
//
// The openCypher language was originally developed by Neo4j, then open-sourced in
// 2015 and contributed to the [openCypher project]under an Apache 2 open-source license.
//
// Note that when invoking this operation in a Neptune cluster that has IAM
// authentication enabled, the IAM user or role making the request must have a
// policy attached that allows one of the following IAM actions in that cluster,
// depending on the query:
//
// [neptune-db:ReadDataViaQuery]
//
// [neptune-db:WriteDataViaQuery]
//
// [neptune-db:DeleteDataViaQuery]
//
// Note also that the [neptune-db:QueryLanguage:OpenCypher] IAM condition key can be used in the policy document to
// restrict the use of openCypher queries (see [Condition keys available in Neptune IAM data-access policy statements]).
//
// [neptune-db:DeleteDataViaQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#deletedataviaquery
// [Condition keys available in Neptune IAM data-access policy statements]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html
// [openCypher project]: https://opencypher.org/
// [neptune-db:ReadDataViaQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#readdataviaquery
// [neptune-db:WriteDataViaQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#writedataviaquery
// [neptune-db:QueryLanguage:OpenCypher]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys
// [Accessing the Neptune Graph with openCypher]: https://docs.aws.amazon.com/neptune/latest/userguide/access-graph-opencypher.html
func neptunedata_ExecuteOpenCypherQuery(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.ExecuteOpenCypherQueryInput{
		// OpenCypherQuery: *string, // Required
	}

	if len(_neptunedataOpenCypherQuery) > 0 {
		input.OpenCypherQuery = aws.String(_neptunedataOpenCypherQuery)
	}
	if len(_neptunedataParameters) > 0 {
		input.Parameters = aws.String(_neptunedataParameters)
	}

	if resp, err := client.ExecuteOpenCypherQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the status of the graph database on the host.
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:GetEngineStatus]IAM action in that cluster.
//
// [neptune-db:GetEngineStatus]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getenginestatus
func neptunedata_GetEngineStatus(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.GetEngineStatusInput{}

	if resp, err := client.GetEngineStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the status of a specified Gremlin query.
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:GetQueryStatus]IAM action in that cluster.
//
// Note that the [neptune-db:QueryLanguage:Gremlin] IAM condition key can be used in the policy document to restrict
// the use of Gremlin queries (see [Condition keys available in Neptune IAM data-access policy statements]).
//
// [Condition keys available in Neptune IAM data-access policy statements]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html
// [neptune-db:GetQueryStatus]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getquerystatus
// [neptune-db:QueryLanguage:Gremlin]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys
func neptunedata_GetGremlinQueryStatus(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.GetGremlinQueryStatusInput{
		// QueryId: *string, // Required
	}

	if len(_neptunedataQueryId) > 0 {
		input.QueryId = aws.String(_neptunedataQueryId)
	}

	if resp, err := client.GetGremlinQueryStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets status information about a specified load job. Neptune keeps track of the
// most recent 1,024 bulk load jobs, and stores the last 10,000 error details per
// job.
//
// See [Neptune Loader Get-Status API] for more information.
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:GetLoaderJobStatus]IAM action in that cluster..
//
// [neptune-db:GetLoaderJobStatus]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getloaderjobstatus
// [Neptune Loader Get-Status API]: https://docs.aws.amazon.com/neptune/latest/userguide/load-api-reference-status.htm
func neptunedata_GetLoaderJobStatus(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.GetLoaderJobStatusInput{
		// LoadId: *string, // Required
	}

	if len(_neptunedataLoadId) > 0 {
		input.LoadId = aws.String(_neptunedataLoadId)
	}
	if len(_neptunedataDetails) > 0 {
		if err := assignInputField(input, "Details", _neptunedataDetails); err != nil {
			log.Errorf("invalid --details: %s", err.Error())
			return
		}
	}
	if len(_neptunedataErrors) > 0 {
		if err := assignInputField(input, "Errors", _neptunedataErrors); err != nil {
			log.Errorf("invalid --errors: %s", err.Error())
			return
		}
	}
	if len(_neptunedataErrorsPerPage) > 0 {
		if err := assignInputField(input, "ErrorsPerPage", _neptunedataErrorsPerPage); err != nil {
			log.Errorf("invalid --errors-per-page: %s", err.Error())
			return
		}
	}
	if len(_neptunedataPage) > 0 {
		if err := assignInputField(input, "Page", _neptunedataPage); err != nil {
			log.Errorf("invalid --page: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetLoaderJobStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a specified data processing job. See [The dataprocessing command]dataprocessing .
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:neptune-db:GetMLDataProcessingJobStatus]IAM action in that cluster.
//
// [neptune-db:neptune-db:GetMLDataProcessingJobStatus]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getmldataprocessingjobstatus
// [The dataprocessing command]: https://docs.aws.amazon.com/neptune/latest/userguide/machine-learning-api-dataprocessing.html
func neptunedata_GetMLDataProcessingJob(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.GetMLDataProcessingJobInput{
		// Id: *string, // Required
	}

	if len(_neptunedataId) > 0 {
		input.Id = aws.String(_neptunedataId)
	}
	if len(_neptunedataNeptuneIamRoleArn) > 0 {
		input.NeptuneIamRoleArn = aws.String(_neptunedataNeptuneIamRoleArn)
	}

	if resp, err := client.GetMLDataProcessingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about an inference endpoint. See [Managing inference endpoints using the endpoints command].
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:GetMLEndpointStatus]IAM action in that cluster.
//
// [neptune-db:GetMLEndpointStatus]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getmlendpointstatus
// [Managing inference endpoints using the endpoints command]: https://docs.aws.amazon.com/neptune/latest/userguide/machine-learning-api-endpoints.html
func neptunedata_GetMLEndpoint(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.GetMLEndpointInput{
		// Id: *string, // Required
	}

	if len(_neptunedataId) > 0 {
		input.Id = aws.String(_neptunedataId)
	}
	if len(_neptunedataNeptuneIamRoleArn) > 0 {
		input.NeptuneIamRoleArn = aws.String(_neptunedataNeptuneIamRoleArn)
	}

	if resp, err := client.GetMLEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a Neptune ML model training job. See [Model training using the modeltraining command]modeltraining .
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:GetMLModelTrainingJobStatus]IAM action in that cluster.
//
// [neptune-db:GetMLModelTrainingJobStatus]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getmlmodeltrainingjobstatus
// [Model training using the modeltraining command]: https://docs.aws.amazon.com/neptune/latest/userguide/machine-learning-api-modeltraining.html
func neptunedata_GetMLModelTrainingJob(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.GetMLModelTrainingJobInput{
		// Id: *string, // Required
	}

	if len(_neptunedataId) > 0 {
		input.Id = aws.String(_neptunedataId)
	}
	if len(_neptunedataNeptuneIamRoleArn) > 0 {
		input.NeptuneIamRoleArn = aws.String(_neptunedataNeptuneIamRoleArn)
	}

	if resp, err := client.GetMLModelTrainingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specified model transform job. See [Use a trained model to generate new model artifacts].
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:GetMLModelTransformJobStatus]IAM action in that cluster.
//
// [neptune-db:GetMLModelTransformJobStatus]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getmlmodeltransformjobstatus
// [Use a trained model to generate new model artifacts]: https://docs.aws.amazon.com/neptune/latest/userguide/machine-learning-model-transform.html
func neptunedata_GetMLModelTransformJob(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.GetMLModelTransformJobInput{
		// Id: *string, // Required
	}

	if len(_neptunedataId) > 0 {
		input.Id = aws.String(_neptunedataId)
	}
	if len(_neptunedataNeptuneIamRoleArn) > 0 {
		input.NeptuneIamRoleArn = aws.String(_neptunedataNeptuneIamRoleArn)
	}

	if resp, err := client.GetMLModelTransformJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the status of a specified openCypher query.
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:GetQueryStatus]IAM action in that cluster.
//
// Note that the [neptune-db:QueryLanguage:OpenCypher] IAM condition key can be used in the policy document to restrict
// the use of openCypher queries (see [Condition keys available in Neptune IAM data-access policy statements]).
//
// [Condition keys available in Neptune IAM data-access policy statements]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html
// [neptune-db:GetQueryStatus]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getquerystatus
// [neptune-db:QueryLanguage:OpenCypher]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys
func neptunedata_GetOpenCypherQueryStatus(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.GetOpenCypherQueryStatusInput{
		// QueryId: *string, // Required
	}

	if len(_neptunedataQueryId) > 0 {
		input.QueryId = aws.String(_neptunedataQueryId)
	}

	if resp, err := client.GetOpenCypherQueryStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets property graph statistics (Gremlin and openCypher).
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:GetStatisticsStatus]IAM action in that cluster.
//
// [neptune-db:GetStatisticsStatus]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getstatisticsstatus
func neptunedata_GetPropertygraphStatistics(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.GetPropertygraphStatisticsInput{}

	if resp, err := client.GetPropertygraphStatistics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a stream for a property graph.
// With the Neptune Streams feature, you can generate a complete sequence of
// change-log entries that record every change made to your graph data as it
// happens. GetPropertygraphStream lets you collect these change-log entries for a
// property graph.
//
// The Neptune streams feature needs to be enabled on your Neptune DBcluster. To
// enable streams, set the [neptune_streams]DB cluster parameter to 1 .
//
// See [Capturing graph changes in real time using Neptune streams].
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:GetStreamRecords]IAM action in that cluster.
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that enables one of the following IAM actions, depending on the query:
//
// Note that you can restrict property-graph queries using the following IAM
// context keys:
//
// [neptune-db:QueryLanguage:Gremlin]
//
// [neptune-db:QueryLanguage:OpenCypher]
//
// See [Condition keys available in Neptune IAM data-access policy statements]).
//
// [Condition keys available in Neptune IAM data-access policy statements]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html
// [neptune_streams]: https://docs.aws.amazon.com/neptune/latest/userguide/parameters.html#parameters-db-cluster-parameters-neptune_streams
// [Capturing graph changes in real time using Neptune streams]: https://docs.aws.amazon.com/neptune/latest/userguide/streams.html
// [neptune-db:GetStreamRecords]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getstreamrecords
// [neptune-db:QueryLanguage:Gremlin]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys
// [neptune-db:QueryLanguage:OpenCypher]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys
func neptunedata_GetPropertygraphStream(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.GetPropertygraphStreamInput{}

	if len(_neptunedataCommitNum) > 0 {
		if err := assignInputField(input, "CommitNum", _neptunedataCommitNum); err != nil {
			log.Errorf("invalid --commit-num: %s", err.Error())
			return
		}
	}
	if len(_neptunedataEncoding) > 0 {
		if err := assignInputField(input, "Encoding", _neptunedataEncoding); err != nil {
			log.Errorf("invalid --encoding: %s", err.Error())
			return
		}
	}
	if len(_neptunedataIteratorType) > 0 {
		if err := assignInputField(input, "IteratorType", _neptunedataIteratorType); err != nil {
			log.Errorf("invalid --iterator-type: %s", err.Error())
			return
		}
	}
	if len(_neptunedataLimit) > 0 {
		if err := assignInputField(input, "Limit", _neptunedataLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_neptunedataOpNum) > 0 {
		if err := assignInputField(input, "OpNum", _neptunedataOpNum); err != nil {
			log.Errorf("invalid --op-num: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetPropertygraphStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a graph summary for a property graph.
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:GetGraphSummary]IAM action in that cluster.
//
// [neptune-db:GetGraphSummary]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getgraphsummary
func neptunedata_GetPropertygraphSummary(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.GetPropertygraphSummaryInput{}

	if len(_neptunedataMode) > 0 {
		if err := assignInputField(input, "Mode", _neptunedataMode); err != nil {
			log.Errorf("invalid --mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetPropertygraphSummary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a graph summary for an RDF graph.
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:GetGraphSummary]IAM action in that cluster.
//
// [neptune-db:GetGraphSummary]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getgraphsummary
func neptunedata_GetRDFGraphSummary(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.GetRDFGraphSummaryInput{}

	if len(_neptunedataMode) > 0 {
		if err := assignInputField(input, "Mode", _neptunedataMode); err != nil {
			log.Errorf("invalid --mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetRDFGraphSummary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets RDF statistics (SPARQL).
func neptunedata_GetSparqlStatistics(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.GetSparqlStatisticsInput{}

	if resp, err := client.GetSparqlStatistics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a stream for an RDF graph.
// With the Neptune Streams feature, you can generate a complete sequence of
// change-log entries that record every change made to your graph data as it
// happens. GetSparqlStream lets you collect these change-log entries for an RDF
// graph.
//
// The Neptune streams feature needs to be enabled on your Neptune DBcluster. To
// enable streams, set the [neptune_streams]DB cluster parameter to 1 .
//
// See [Capturing graph changes in real time using Neptune streams].
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:GetStreamRecords]IAM action in that cluster.
//
// Note that the [neptune-db:QueryLanguage:Sparql] IAM condition key can be used in the policy document to restrict
// the use of SPARQL queries (see [Condition keys available in Neptune IAM data-access policy statements]).
//
// [Condition keys available in Neptune IAM data-access policy statements]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html
// [neptune-db:QueryLanguage:Sparql]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys
// [neptune_streams]: https://docs.aws.amazon.com/neptune/latest/userguide/parameters.html#parameters-db-cluster-parameters-neptune_streams
// [Capturing graph changes in real time using Neptune streams]: https://docs.aws.amazon.com/neptune/latest/userguide/streams.html
// [neptune-db:GetStreamRecords]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getstreamrecords
func neptunedata_GetSparqlStream(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.GetSparqlStreamInput{}

	if len(_neptunedataCommitNum) > 0 {
		if err := assignInputField(input, "CommitNum", _neptunedataCommitNum); err != nil {
			log.Errorf("invalid --commit-num: %s", err.Error())
			return
		}
	}
	if len(_neptunedataEncoding) > 0 {
		if err := assignInputField(input, "Encoding", _neptunedataEncoding); err != nil {
			log.Errorf("invalid --encoding: %s", err.Error())
			return
		}
	}
	if len(_neptunedataIteratorType) > 0 {
		if err := assignInputField(input, "IteratorType", _neptunedataIteratorType); err != nil {
			log.Errorf("invalid --iterator-type: %s", err.Error())
			return
		}
	}
	if len(_neptunedataLimit) > 0 {
		if err := assignInputField(input, "Limit", _neptunedataLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_neptunedataOpNum) > 0 {
		if err := assignInputField(input, "OpNum", _neptunedataOpNum); err != nil {
			log.Errorf("invalid --op-num: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetSparqlStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists active Gremlin queries. See [Gremlin query status API] for details about the output.
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:GetQueryStatus]IAM action in that cluster.
//
// Note that the [neptune-db:QueryLanguage:Gremlin] IAM condition key can be used in the policy document to restrict
// the use of Gremlin queries (see [Condition keys available in Neptune IAM data-access policy statements]).
//
// [Condition keys available in Neptune IAM data-access policy statements]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html
// [neptune-db:GetQueryStatus]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getquerystatus
// [neptune-db:QueryLanguage:Gremlin]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys
// [Gremlin query status API]: https://docs.aws.amazon.com/neptune/latest/userguide/gremlin-api-status.html
func neptunedata_ListGremlinQueries(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.ListGremlinQueriesInput{}

	if len(_neptunedataIncludeWaiting) > 0 {
		if err := assignInputField(input, "IncludeWaiting", _neptunedataIncludeWaiting); err != nil {
			log.Errorf("invalid --include-waiting: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListGremlinQueries(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of the loadIds for all active loader jobs.
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:ListLoaderJobs]IAM action in that cluster..
//
// [neptune-db:ListLoaderJobs]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#listloaderjobs
func neptunedata_ListLoaderJobs(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.ListLoaderJobsInput{}

	if len(_neptunedataIncludeQueuedLoads) > 0 {
		if err := assignInputField(input, "IncludeQueuedLoads", _neptunedataIncludeQueuedLoads); err != nil {
			log.Errorf("invalid --include-queued-loads: %s", err.Error())
			return
		}
	}
	if len(_neptunedataLimit) > 0 {
		if err := assignInputField(input, "Limit", _neptunedataLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListLoaderJobs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of Neptune ML data processing jobs. See [Listing active data-processing jobs using the Neptune ML dataprocessing command].
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:ListMLDataProcessingJobs]IAM action in that cluster.
//
// [neptune-db:ListMLDataProcessingJobs]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#listmldataprocessingjobs
// [Listing active data-processing jobs using the Neptune ML dataprocessing command]: https://docs.aws.amazon.com/neptune/latest/userguide/machine-learning-api-dataprocessing.html#machine-learning-api-dataprocessing-list-jobs
func neptunedata_ListMLDataProcessingJobs(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.ListMLDataProcessingJobsInput{}

	if len(_neptunedataMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _neptunedataMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_neptunedataNeptuneIamRoleArn) > 0 {
		input.NeptuneIamRoleArn = aws.String(_neptunedataNeptuneIamRoleArn)
	}

	if resp, err := client.ListMLDataProcessingJobs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists existing inference endpoints. See [Managing inference endpoints using the endpoints command].
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:ListMLEndpoints]IAM action in that cluster.
//
// [neptune-db:ListMLEndpoints]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#listmlendpoints
// [Managing inference endpoints using the endpoints command]: https://docs.aws.amazon.com/neptune/latest/userguide/machine-learning-api-endpoints.html
func neptunedata_ListMLEndpoints(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.ListMLEndpointsInput{}

	if len(_neptunedataMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _neptunedataMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_neptunedataNeptuneIamRoleArn) > 0 {
		input.NeptuneIamRoleArn = aws.String(_neptunedataNeptuneIamRoleArn)
	}

	if resp, err := client.ListMLEndpoints(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists Neptune ML model-training jobs. See [Model training using the modeltraining command]modeltraining .
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:neptune-db:ListMLModelTrainingJobs]IAM action in that cluster.
//
// [neptune-db:neptune-db:ListMLModelTrainingJobs]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#neptune-db:listmlmodeltrainingjobs
// [Model training using the modeltraining command]: https://docs.aws.amazon.com/neptune/latest/userguide/machine-learning-api-modeltraining.html
func neptunedata_ListMLModelTrainingJobs(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.ListMLModelTrainingJobsInput{}

	if len(_neptunedataMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _neptunedataMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_neptunedataNeptuneIamRoleArn) > 0 {
		input.NeptuneIamRoleArn = aws.String(_neptunedataNeptuneIamRoleArn)
	}

	if resp, err := client.ListMLModelTrainingJobs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of model transform job IDs. See [Use a trained model to generate new model artifacts].
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:ListMLModelTransformJobs]IAM action in that cluster.
//
// [neptune-db:ListMLModelTransformJobs]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#listmlmodeltransformjobs
// [Use a trained model to generate new model artifacts]: https://docs.aws.amazon.com/neptune/latest/userguide/machine-learning-model-transform.html
func neptunedata_ListMLModelTransformJobs(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.ListMLModelTransformJobsInput{}

	if len(_neptunedataMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _neptunedataMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_neptunedataNeptuneIamRoleArn) > 0 {
		input.NeptuneIamRoleArn = aws.String(_neptunedataNeptuneIamRoleArn)
	}

	if resp, err := client.ListMLModelTransformJobs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists active openCypher queries. See [Neptune openCypher status endpoint] for more information.
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:GetQueryStatus]IAM action in that cluster.
//
// Note that the [neptune-db:QueryLanguage:OpenCypher] IAM condition key can be used in the policy document to restrict
// the use of openCypher queries (see [Condition keys available in Neptune IAM data-access policy statements]).
//
// [Condition keys available in Neptune IAM data-access policy statements]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html
// [neptune-db:GetQueryStatus]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getquerystatus
// [neptune-db:QueryLanguage:OpenCypher]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys
// [Neptune openCypher status endpoint]: https://docs.aws.amazon.com/neptune/latest/userguide/access-graph-opencypher-status.html
func neptunedata_ListOpenCypherQueries(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.ListOpenCypherQueriesInput{}

	if len(_neptunedataIncludeWaiting) > 0 {
		if err := assignInputField(input, "IncludeWaiting", _neptunedataIncludeWaiting); err != nil {
			log.Errorf("invalid --include-waiting: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListOpenCypherQueries(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Manages the generation and use of property graph statistics.
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:ManageStatistics]IAM action in that cluster.
//
// [neptune-db:ManageStatistics]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#managestatistics
func neptunedata_ManagePropertygraphStatistics(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.ManagePropertygraphStatisticsInput{}

	if len(_neptunedataMode) > 0 {
		if err := assignInputField(input, "Mode", _neptunedataMode); err != nil {
			log.Errorf("invalid --mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.ManagePropertygraphStatistics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Manages the generation and use of RDF graph statistics.
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:ManageStatistics]IAM action in that cluster.
//
// [neptune-db:ManageStatistics]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#managestatistics
func neptunedata_ManageSparqlStatistics(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.ManageSparqlStatisticsInput{}

	if len(_neptunedataMode) > 0 {
		if err := assignInputField(input, "Mode", _neptunedataMode); err != nil {
			log.Errorf("invalid --mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.ManageSparqlStatistics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a Neptune bulk loader job to load data from an Amazon S3 bucket into a
// Neptune DB instance. See [Using the Amazon Neptune Bulk Loader to Ingest Data].
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:StartLoaderJob]IAM action in that cluster.
//
// [neptune-db:StartLoaderJob]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#startloaderjob
// [Using the Amazon Neptune Bulk Loader to Ingest Data]: https://docs.aws.amazon.com/neptune/latest/userguide/bulk-load.html
func neptunedata_StartLoaderJob(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.StartLoaderJobInput{
		// Format: types.Format, // Required
		// IamRoleArn: *string, // Required
		// S3BucketRegion: types.S3BucketRegion, // Required
		// Source: *string, // Required
	}

	if len(_neptunedataFormat) > 0 {
		if err := assignInputField(input, "Format", _neptunedataFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_neptunedataIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_neptunedataIamRoleArn)
	}
	if len(_neptunedataS3BucketRegion) > 0 {
		if err := assignInputField(input, "S3BucketRegion", _neptunedataS3BucketRegion); err != nil {
			log.Errorf("invalid --s3-bucket-region: %s", err.Error())
			return
		}
	}
	if len(_neptunedataSource) > 0 {
		input.Source = aws.String(_neptunedataSource)
	}
	if len(_neptunedataDependencies) > 0 {
		input.Dependencies = append([]string(nil), _neptunedataDependencies...)
	}
	if len(_neptunedataEdgeOnlyLoad) > 0 {
		if err := assignInputField(input, "EdgeOnlyLoad", _neptunedataEdgeOnlyLoad); err != nil {
			log.Errorf("invalid --edge-only-load: %s", err.Error())
			return
		}
	}
	if len(_neptunedataFailOnError) > 0 {
		if err := assignInputField(input, "FailOnError", _neptunedataFailOnError); err != nil {
			log.Errorf("invalid --fail-on-error: %s", err.Error())
			return
		}
	}
	if len(_neptunedataMode) > 0 {
		if err := assignInputField(input, "Mode", _neptunedataMode); err != nil {
			log.Errorf("invalid --mode: %s", err.Error())
			return
		}
	}
	if len(_neptunedataParallelism) > 0 {
		if err := assignInputField(input, "Parallelism", _neptunedataParallelism); err != nil {
			log.Errorf("invalid --parallelism: %s", err.Error())
			return
		}
	}
	if len(_neptunedataParserConfiguration) > 0 {
		if err := assignInputField(input, "ParserConfiguration", _neptunedataParserConfiguration); err != nil {
			log.Errorf("invalid --parser-configuration: %s", err.Error())
			return
		}
	}
	if len(_neptunedataQueueRequest) > 0 {
		if err := assignInputField(input, "QueueRequest", _neptunedataQueueRequest); err != nil {
			log.Errorf("invalid --queue-request: %s", err.Error())
			return
		}
	}
	if len(_neptunedataUpdateSingleCardinalityProperties) > 0 {
		if err := assignInputField(input, "UpdateSingleCardinalityProperties", _neptunedataUpdateSingleCardinalityProperties); err != nil {
			log.Errorf("invalid --update-single-cardinality-properties: %s", err.Error())
			return
		}
	}
	if len(_neptunedataUserProvidedEdgeIds) > 0 {
		if err := assignInputField(input, "UserProvidedEdgeIds", _neptunedataUserProvidedEdgeIds); err != nil {
			log.Errorf("invalid --user-provided-edge-ids: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartLoaderJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Neptune ML data processing job for processing the graph data
// exported from Neptune for training. See [The dataprocessing command]dataprocessing .
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:StartMLModelDataProcessingJob]IAM action in that cluster.
//
// [neptune-db:StartMLModelDataProcessingJob]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#startmlmodeldataprocessingjob
// [The dataprocessing command]: https://docs.aws.amazon.com/neptune/latest/userguide/machine-learning-api-dataprocessing.html
func neptunedata_StartMLDataProcessingJob(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.StartMLDataProcessingJobInput{
		// InputDataS3Location: *string, // Required
		// ProcessedDataS3Location: *string, // Required
	}

	if len(_neptunedataInputDataS3Location) > 0 {
		input.InputDataS3Location = aws.String(_neptunedataInputDataS3Location)
	}
	if len(_neptunedataProcessedDataS3Location) > 0 {
		input.ProcessedDataS3Location = aws.String(_neptunedataProcessedDataS3Location)
	}
	if len(_neptunedataConfigFileName) > 0 {
		input.ConfigFileName = aws.String(_neptunedataConfigFileName)
	}
	if len(_neptunedataId) > 0 {
		input.Id = aws.String(_neptunedataId)
	}
	if len(_neptunedataModelType) > 0 {
		input.ModelType = aws.String(_neptunedataModelType)
	}
	if len(_neptunedataNeptuneIamRoleArn) > 0 {
		input.NeptuneIamRoleArn = aws.String(_neptunedataNeptuneIamRoleArn)
	}
	if len(_neptunedataPreviousDataProcessingJobId) > 0 {
		input.PreviousDataProcessingJobId = aws.String(_neptunedataPreviousDataProcessingJobId)
	}
	if len(_neptunedataProcessingInstanceType) > 0 {
		input.ProcessingInstanceType = aws.String(_neptunedataProcessingInstanceType)
	}
	if len(_neptunedataProcessingInstanceVolumeSizeInGB) > 0 {
		if err := assignInputField(input, "ProcessingInstanceVolumeSizeInGB", _neptunedataProcessingInstanceVolumeSizeInGB); err != nil {
			log.Errorf("invalid --processing-instance-volume-size-in-gb: %s", err.Error())
			return
		}
	}
	if len(_neptunedataProcessingTimeOutInSeconds) > 0 {
		if err := assignInputField(input, "ProcessingTimeOutInSeconds", _neptunedataProcessingTimeOutInSeconds); err != nil {
			log.Errorf("invalid --processing-time-out-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_neptunedataS3OutputEncryptionKMSKey) > 0 {
		input.S3OutputEncryptionKMSKey = aws.String(_neptunedataS3OutputEncryptionKMSKey)
	}
	if len(_neptunedataSagemakerIamRoleArn) > 0 {
		input.SagemakerIamRoleArn = aws.String(_neptunedataSagemakerIamRoleArn)
	}
	if len(_neptunedataSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _neptunedataSecurityGroupIds...)
	}
	if len(_neptunedataSubnets) > 0 {
		input.Subnets = append([]string(nil), _neptunedataSubnets...)
	}
	if len(_neptunedataVolumeEncryptionKMSKey) > 0 {
		input.VolumeEncryptionKMSKey = aws.String(_neptunedataVolumeEncryptionKMSKey)
	}

	if resp, err := client.StartMLDataProcessingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Neptune ML model training job. See [Model training using the modeltraining command]modeltraining .
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:StartMLModelTrainingJob]IAM action in that cluster.
//
// [neptune-db:StartMLModelTrainingJob]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#startmlmodeltrainingjob
// [Model training using the modeltraining command]: https://docs.aws.amazon.com/neptune/latest/userguide/machine-learning-api-modeltraining.html
func neptunedata_StartMLModelTrainingJob(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.StartMLModelTrainingJobInput{
		// DataProcessingJobId: *string, // Required
		// TrainModelS3Location: *string, // Required
	}

	if len(_neptunedataDataProcessingJobId) > 0 {
		input.DataProcessingJobId = aws.String(_neptunedataDataProcessingJobId)
	}
	if len(_neptunedataTrainModelS3Location) > 0 {
		input.TrainModelS3Location = aws.String(_neptunedataTrainModelS3Location)
	}
	if len(_neptunedataBaseProcessingInstanceType) > 0 {
		input.BaseProcessingInstanceType = aws.String(_neptunedataBaseProcessingInstanceType)
	}
	if len(_neptunedataCustomModelTrainingParameters) > 0 {
		if err := assignInputField(input, "CustomModelTrainingParameters", _neptunedataCustomModelTrainingParameters); err != nil {
			log.Errorf("invalid --custom-model-training-parameters: %s", err.Error())
			return
		}
	}
	if len(_neptunedataEnableManagedSpotTraining) > 0 {
		if err := assignInputField(input, "EnableManagedSpotTraining", _neptunedataEnableManagedSpotTraining); err != nil {
			log.Errorf("invalid --enable-managed-spot-training: %s", err.Error())
			return
		}
	}
	if len(_neptunedataId) > 0 {
		input.Id = aws.String(_neptunedataId)
	}
	if len(_neptunedataMaxHPONumberOfTrainingJobs) > 0 {
		if err := assignInputField(input, "MaxHPONumberOfTrainingJobs", _neptunedataMaxHPONumberOfTrainingJobs); err != nil {
			log.Errorf("invalid --max-hpo-number-of-training-jobs: %s", err.Error())
			return
		}
	}
	if len(_neptunedataMaxHPOParallelTrainingJobs) > 0 {
		if err := assignInputField(input, "MaxHPOParallelTrainingJobs", _neptunedataMaxHPOParallelTrainingJobs); err != nil {
			log.Errorf("invalid --max-hpo-parallel-training-jobs: %s", err.Error())
			return
		}
	}
	if len(_neptunedataNeptuneIamRoleArn) > 0 {
		input.NeptuneIamRoleArn = aws.String(_neptunedataNeptuneIamRoleArn)
	}
	if len(_neptunedataPreviousModelTrainingJobId) > 0 {
		input.PreviousModelTrainingJobId = aws.String(_neptunedataPreviousModelTrainingJobId)
	}
	if len(_neptunedataS3OutputEncryptionKMSKey) > 0 {
		input.S3OutputEncryptionKMSKey = aws.String(_neptunedataS3OutputEncryptionKMSKey)
	}
	if len(_neptunedataSagemakerIamRoleArn) > 0 {
		input.SagemakerIamRoleArn = aws.String(_neptunedataSagemakerIamRoleArn)
	}
	if len(_neptunedataSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _neptunedataSecurityGroupIds...)
	}
	if len(_neptunedataSubnets) > 0 {
		input.Subnets = append([]string(nil), _neptunedataSubnets...)
	}
	if len(_neptunedataTrainingInstanceType) > 0 {
		input.TrainingInstanceType = aws.String(_neptunedataTrainingInstanceType)
	}
	if len(_neptunedataTrainingInstanceVolumeSizeInGB) > 0 {
		if err := assignInputField(input, "TrainingInstanceVolumeSizeInGB", _neptunedataTrainingInstanceVolumeSizeInGB); err != nil {
			log.Errorf("invalid --training-instance-volume-size-in-gb: %s", err.Error())
			return
		}
	}
	if len(_neptunedataTrainingTimeOutInSeconds) > 0 {
		if err := assignInputField(input, "TrainingTimeOutInSeconds", _neptunedataTrainingTimeOutInSeconds); err != nil {
			log.Errorf("invalid --training-time-out-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_neptunedataVolumeEncryptionKMSKey) > 0 {
		input.VolumeEncryptionKMSKey = aws.String(_neptunedataVolumeEncryptionKMSKey)
	}

	if resp, err := client.StartMLModelTrainingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new model transform job. See [Use a trained model to generate new model artifacts].
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:StartMLModelTransformJob]IAM action in that cluster.
//
// [neptune-db:StartMLModelTransformJob]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#startmlmodeltransformjob
// [Use a trained model to generate new model artifacts]: https://docs.aws.amazon.com/neptune/latest/userguide/machine-learning-model-transform.html
func neptunedata_StartMLModelTransformJob(cfg aws.Config, client *neptunedata.Client) {
	input := &neptunedata.StartMLModelTransformJobInput{
		// ModelTransformOutputS3Location: *string, // Required
	}

	if len(_neptunedataModelTransformOutputS3Location) > 0 {
		input.ModelTransformOutputS3Location = aws.String(_neptunedataModelTransformOutputS3Location)
	}
	if len(_neptunedataBaseProcessingInstanceType) > 0 {
		input.BaseProcessingInstanceType = aws.String(_neptunedataBaseProcessingInstanceType)
	}
	if len(_neptunedataBaseProcessingInstanceVolumeSizeInGB) > 0 {
		if err := assignInputField(input, "BaseProcessingInstanceVolumeSizeInGB", _neptunedataBaseProcessingInstanceVolumeSizeInGB); err != nil {
			log.Errorf("invalid --base-processing-instance-volume-size-in-gb: %s", err.Error())
			return
		}
	}
	if len(_neptunedataCustomModelTransformParameters) > 0 {
		if err := assignInputField(input, "CustomModelTransformParameters", _neptunedataCustomModelTransformParameters); err != nil {
			log.Errorf("invalid --custom-model-transform-parameters: %s", err.Error())
			return
		}
	}
	if len(_neptunedataDataProcessingJobId) > 0 {
		input.DataProcessingJobId = aws.String(_neptunedataDataProcessingJobId)
	}
	if len(_neptunedataId) > 0 {
		input.Id = aws.String(_neptunedataId)
	}
	if len(_neptunedataMlModelTrainingJobId) > 0 {
		input.MlModelTrainingJobId = aws.String(_neptunedataMlModelTrainingJobId)
	}
	if len(_neptunedataNeptuneIamRoleArn) > 0 {
		input.NeptuneIamRoleArn = aws.String(_neptunedataNeptuneIamRoleArn)
	}
	if len(_neptunedataS3OutputEncryptionKMSKey) > 0 {
		input.S3OutputEncryptionKMSKey = aws.String(_neptunedataS3OutputEncryptionKMSKey)
	}
	if len(_neptunedataSagemakerIamRoleArn) > 0 {
		input.SagemakerIamRoleArn = aws.String(_neptunedataSagemakerIamRoleArn)
	}
	if len(_neptunedataSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _neptunedataSecurityGroupIds...)
	}
	if len(_neptunedataSubnets) > 0 {
		input.Subnets = append([]string(nil), _neptunedataSubnets...)
	}
	if len(_neptunedataTrainingJobName) > 0 {
		input.TrainingJobName = aws.String(_neptunedataTrainingJobName)
	}
	if len(_neptunedataVolumeEncryptionKMSKey) > 0 {
		input.VolumeEncryptionKMSKey = aws.String(_neptunedataVolumeEncryptionKMSKey)
	}

	if resp, err := client.StartMLModelTransformJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_neptunedataCmd)
	_neptunedataCmd.Flags().SortFlags = false

	_neptunedataCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_neptunedataCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_neptunedataCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_neptunedataCmd.Flags().StringVarP(&_neptunedataAction, "action", "", "", "Action")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataBaseProcessingInstanceType, "base-processing-instance-type", "", "", "Base Processing Instance Type")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataBaseProcessingInstanceVolumeSizeInGB, "base-processing-instance-volume-size-in-gb", "", "", "Base Processing Instance Volume Size In Gb")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataChop, "chop", "", "", "Chop")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataClean, "clean", "", "", "Clean")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataCommitNum, "commit-num", "", "", "Commit Num")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataConfigFileName, "config-file-name", "", "", "Config File Name")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataCustomModelTrainingParameters, "custom-model-training-parameters", "", "", "Custom Model Training Parameters")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataCustomModelTransformParameters, "custom-model-transform-parameters", "", "", "Custom Model Transform Parameters")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataDataProcessingJobId, "data-processing-job-id", "", "", "Data Processing Job ID")
	_neptunedataCmd.Flags().StringSliceVarP(&_neptunedataDependencies, "dependencies", "", nil, "Dependencies")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataDetails, "details", "", "", "Details")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataEdgeOnlyLoad, "edge-only-load", "", "", "Edge Only Load")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataEnableManagedSpotTraining, "enable-managed-spot-training", "", "", "Enable Managed Spot Training")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataEncoding, "encoding", "", "", "Encoding")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataErrors, "errors", "", "", "Errors")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataErrorsPerPage, "errors-per-page", "", "", "Errors Per Page")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataExplainMode, "explain-mode", "", "", "Explain Mode")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataFailOnError, "fail-on-error", "", "", "Fail On Error")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataFormat, "format", "", "", "Format")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataGremlinQuery, "gremlin-query", "", "", "Gremlin Query")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataIamRoleArn, "iam-role-arn", "", "", "IAM Role ARN")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataId, "id", "", "", "ID")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataIncludeQueuedLoads, "include-queued-loads", "", "", "Include Queued Loads")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataIncludeWaiting, "include-waiting", "", "", "Include Waiting")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataIndexOps, "index-ops", "", "", "Index Ops")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataInputDataS3Location, "input-data-s3-location", "", "", "Input Data S3 Location")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataInstanceCount, "instance-count", "", "", "Instance Count")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataInstanceType, "instance-type", "", "", "Instance Type")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataIteratorType, "iterator-type", "", "", "Iterator Type")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataLimit, "limit", "", "", "Limit")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataLoadId, "load-id", "", "", "Load ID")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataMaxHPONumberOfTrainingJobs, "max-hpo-number-of-training-jobs", "", "", "Max Hpo Number Of Training Jobs")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataMaxHPOParallelTrainingJobs, "max-hpo-parallel-training-jobs", "", "", "Max Hpo Parallel Training Jobs")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataMaxItems, "max-items", "", "", "Max Items")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataMlModelTrainingJobId, "ml-model-training-job-id", "", "", "Ml Model Training Job ID")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataMlModelTransformJobId, "ml-model-transform-job-id", "", "", "Ml Model Transform Job ID")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataMode, "mode", "", "", "Mode")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataModelName, "model-name", "", "", "Model Name")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataModelTransformOutputS3Location, "model-transform-output-s3-location", "", "", "Model Transform Output S3 Location")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataModelType, "model-type", "", "", "Model Type")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataNeptuneIamRoleArn, "neptune-iam-role-arn", "", "", "Neptune IAM Role ARN")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataOpNum, "op-num", "", "", "Op Num")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataOpenCypherQuery, "open-cypher-query", "", "", "Open Cypher Query")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataPage, "page", "", "", "Page")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataParallelism, "parallelism", "", "", "Parallelism")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataParameters, "parameters", "", "", "Parameters")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataParserConfiguration, "parser-configuration", "", "", "Parser Configuration")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataPreviousDataProcessingJobId, "previous-data-processing-job-id", "", "", "Previous Data Processing Job ID")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataPreviousModelTrainingJobId, "previous-model-training-job-id", "", "", "Previous Model Training Job ID")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataProcessedDataS3Location, "processed-data-s3-location", "", "", "Processed Data S3 Location")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataProcessingInstanceType, "processing-instance-type", "", "", "Processing Instance Type")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataProcessingInstanceVolumeSizeInGB, "processing-instance-volume-size-in-gb", "", "", "Processing Instance Volume Size In Gb")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataProcessingTimeOutInSeconds, "processing-time-out-in-seconds", "", "", "Processing Time Out In Seconds")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataQueryId, "query-id", "", "", "Query ID")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataQueueRequest, "queue-request", "", "", "Queue Request")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataResults, "results", "", "", "Results")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataS3BucketRegion, "s3-bucket-region", "", "", "S3 Bucket Region")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataS3OutputEncryptionKMSKey, "s3-output-encryption-kms-key", "", "", "S3 Output Encryption KMS Key")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataSagemakerIamRoleArn, "sagemaker-iam-role-arn", "", "", "Sagemaker IAM Role ARN")
	_neptunedataCmd.Flags().StringSliceVarP(&_neptunedataSecurityGroupIds, "security-group-ids", "", nil, "Security Group Ids")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataSerializer, "serializer", "", "", "Serializer")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataSilent, "silent", "", "", "Silent")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataSource, "source", "", "", "Source")
	_neptunedataCmd.Flags().StringSliceVarP(&_neptunedataSubnets, "subnets", "", nil, "Subnets")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataToken, "token", "", "", "Token")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataTrainModelS3Location, "train-model-s3-location", "", "", "Train Model S3 Location")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataTrainingInstanceType, "training-instance-type", "", "", "Training Instance Type")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataTrainingInstanceVolumeSizeInGB, "training-instance-volume-size-in-gb", "", "", "Training Instance Volume Size In Gb")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataTrainingJobName, "training-job-name", "", "", "Training Job Name")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataTrainingTimeOutInSeconds, "training-time-out-in-seconds", "", "", "Training Time Out In Seconds")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataUpdate, "update", "", "", "Update")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataUpdateSingleCardinalityProperties, "update-single-cardinality-properties", "", "", "Update Single Cardinality Properties")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataUserProvidedEdgeIds, "user-provided-edge-ids", "", "", "User Provided Edge Ids")
	_neptunedataCmd.Flags().StringVarP(&_neptunedataVolumeEncryptionKMSKey, "volume-encryption-kms-key", "", "", "Volume Encryption KMS Key")

	_neptunedataCmd.Flags().BoolVarP(&_neptunedataCancelGremlinQuery, "cancel-gremlin-query", "", false, "Cancel Gremlin Query")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataCancelLoaderJob, "cancel-loader-job", "", false, "Cancel Loader Job")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataCancelMLDataProcessingJob, "cancel-ml-data-processing-job", "", false, "Cancel Ml Data Processing Job")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataCancelMLModelTrainingJob, "cancel-ml-model-training-job", "", false, "Cancel Ml Model Training Job")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataCancelMLModelTransformJob, "cancel-ml-model-transform-job", "", false, "Cancel Ml Model Transform Job")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataCancelOpenCypherQuery, "cancel-open-cypher-query", "", false, "Cancel Open Cypher Query")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataCreateMLEndpoint, "create-ml-endpoint", "", false, "Create Ml Endpoint")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataDeleteMLEndpoint, "delete-ml-endpoint", "", false, "Delete Ml Endpoint")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataDeletePropertygraphStatistics, "delete-propertygraph-statistics", "", false, "Delete Propertygraph Statistics")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataDeleteSparqlStatistics, "delete-sparql-statistics", "", false, "Delete Sparql Statistics")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataExecuteFastReset, "execute-fast-reset", "", false, "Execute Fast Reset")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataExecuteGremlinExplainQuery, "execute-gremlin-explain-query", "", false, "Execute Gremlin Explain Query")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataExecuteGremlinProfileQuery, "execute-gremlin-profile-query", "", false, "Execute Gremlin Profile Query")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataExecuteGremlinQuery, "execute-gremlin-query", "", false, "Execute Gremlin Query")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataExecuteOpenCypherExplainQuery, "execute-open-cypher-explain-query", "", false, "Execute Open Cypher Explain Query")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataExecuteOpenCypherQuery, "execute-open-cypher-query", "", false, "Execute Open Cypher Query")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataGetEngineStatus, "get-engine-status", "", false, "Get Engine Status")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataGetGremlinQueryStatus, "get-gremlin-query-status", "", false, "Get Gremlin Query Status")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataGetLoaderJobStatus, "get-loader-job-status", "", false, "Get Loader Job Status")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataGetMLDataProcessingJob, "get-ml-data-processing-job", "", false, "Get Ml Data Processing Job")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataGetMLEndpoint, "get-ml-endpoint", "", false, "Get Ml Endpoint")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataGetMLModelTrainingJob, "get-ml-model-training-job", "", false, "Get Ml Model Training Job")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataGetMLModelTransformJob, "get-ml-model-transform-job", "", false, "Get Ml Model Transform Job")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataGetOpenCypherQueryStatus, "get-open-cypher-query-status", "", false, "Get Open Cypher Query Status")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataGetPropertygraphStatistics, "get-propertygraph-statistics", "", false, "Get Propertygraph Statistics")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataGetPropertygraphStream, "get-propertygraph-stream", "", false, "Get Propertygraph Stream")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataGetPropertygraphSummary, "get-propertygraph-summary", "", false, "Get Propertygraph Summary")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataGetRDFGraphSummary, "get-rdf-graph-summary", "", false, "Get Rdf Graph Summary")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataGetSparqlStatistics, "get-sparql-statistics", "", false, "Get Sparql Statistics")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataGetSparqlStream, "get-sparql-stream", "", false, "Get Sparql Stream")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataListGremlinQueries, "list-gremlin-queries", "", false, "List Gremlin Queries")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataListLoaderJobs, "list-loader-jobs", "", false, "List Loader Jobs")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataListMLDataProcessingJobs, "list-ml-data-processing-jobs", "", false, "List Ml Data Processing Jobs")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataListMLEndpoints, "list-ml-endpoints", "", false, "List Ml Endpoints")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataListMLModelTrainingJobs, "list-ml-model-training-jobs", "", false, "List Ml Model Training Jobs")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataListMLModelTransformJobs, "list-ml-model-transform-jobs", "", false, "List Ml Model Transform Jobs")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataListOpenCypherQueries, "list-open-cypher-queries", "", false, "List Open Cypher Queries")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataManagePropertygraphStatistics, "manage-propertygraph-statistics", "", false, "Manage Propertygraph Statistics")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataManageSparqlStatistics, "manage-sparql-statistics", "", false, "Manage Sparql Statistics")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataStartLoaderJob, "start-loader-job", "", false, "Start Loader Job")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataStartMLDataProcessingJob, "start-ml-data-processing-job", "", false, "Start Ml Data Processing Job")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataStartMLModelTrainingJob, "start-ml-model-training-job", "", false, "Start Ml Model Training Job")
	_neptunedataCmd.Flags().BoolVarP(&_neptunedataStartMLModelTransformJob, "start-ml-model-transform-job", "", false, "Start Ml Model Transform Job")

}
