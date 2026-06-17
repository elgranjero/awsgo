package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lookoutequipment"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// lookoutequipmentCmd represents the lookoutequipment command
var _lookoutequipmentCmd = &cobra.Command{
	Use:   "lookoutequipment",
	Short: "AWS lookoutequipment CLI",
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
		client := lookoutequipment.NewFromConfig(cfg)
		if _lookoutequipmentCreateDataset {
			lookoutequipment_CreateDataset(cfg, client)
			return
		}
		if _lookoutequipmentCreateInferenceScheduler {
			lookoutequipment_CreateInferenceScheduler(cfg, client)
			return
		}
		if _lookoutequipmentCreateLabel {
			lookoutequipment_CreateLabel(cfg, client)
			return
		}
		if _lookoutequipmentCreateLabelGroup {
			lookoutequipment_CreateLabelGroup(cfg, client)
			return
		}
		if _lookoutequipmentCreateModel {
			lookoutequipment_CreateModel(cfg, client)
			return
		}
		if _lookoutequipmentCreateRetrainingScheduler {
			lookoutequipment_CreateRetrainingScheduler(cfg, client)
			return
		}
		if _lookoutequipmentDeleteDataset {
			lookoutequipment_DeleteDataset(cfg, client)
			return
		}
		if _lookoutequipmentDeleteInferenceScheduler {
			lookoutequipment_DeleteInferenceScheduler(cfg, client)
			return
		}
		if _lookoutequipmentDeleteLabel {
			lookoutequipment_DeleteLabel(cfg, client)
			return
		}
		if _lookoutequipmentDeleteLabelGroup {
			lookoutequipment_DeleteLabelGroup(cfg, client)
			return
		}
		if _lookoutequipmentDeleteModel {
			lookoutequipment_DeleteModel(cfg, client)
			return
		}
		if _lookoutequipmentDeleteResourcePolicy {
			lookoutequipment_DeleteResourcePolicy(cfg, client)
			return
		}
		if _lookoutequipmentDeleteRetrainingScheduler {
			lookoutequipment_DeleteRetrainingScheduler(cfg, client)
			return
		}
		if _lookoutequipmentDescribeDataIngestionJob {
			lookoutequipment_DescribeDataIngestionJob(cfg, client)
			return
		}
		if _lookoutequipmentDescribeDataset {
			lookoutequipment_DescribeDataset(cfg, client)
			return
		}
		if _lookoutequipmentDescribeInferenceScheduler {
			lookoutequipment_DescribeInferenceScheduler(cfg, client)
			return
		}
		if _lookoutequipmentDescribeLabel {
			lookoutequipment_DescribeLabel(cfg, client)
			return
		}
		if _lookoutequipmentDescribeLabelGroup {
			lookoutequipment_DescribeLabelGroup(cfg, client)
			return
		}
		if _lookoutequipmentDescribeModel {
			lookoutequipment_DescribeModel(cfg, client)
			return
		}
		if _lookoutequipmentDescribeModelVersion {
			lookoutequipment_DescribeModelVersion(cfg, client)
			return
		}
		if _lookoutequipmentDescribeResourcePolicy {
			lookoutequipment_DescribeResourcePolicy(cfg, client)
			return
		}
		if _lookoutequipmentDescribeRetrainingScheduler {
			lookoutequipment_DescribeRetrainingScheduler(cfg, client)
			return
		}
		if _lookoutequipmentImportDataset {
			lookoutequipment_ImportDataset(cfg, client)
			return
		}
		if _lookoutequipmentImportModelVersion {
			lookoutequipment_ImportModelVersion(cfg, client)
			return
		}
		if _lookoutequipmentListDataIngestionJobs {
			lookoutequipment_ListDataIngestionJobs(cfg, client)
			return
		}
		if _lookoutequipmentListDatasets {
			lookoutequipment_ListDatasets(cfg, client)
			return
		}
		if _lookoutequipmentListInferenceEvents {
			lookoutequipment_ListInferenceEvents(cfg, client)
			return
		}
		if _lookoutequipmentListInferenceExecutions {
			lookoutequipment_ListInferenceExecutions(cfg, client)
			return
		}
		if _lookoutequipmentListInferenceSchedulers {
			lookoutequipment_ListInferenceSchedulers(cfg, client)
			return
		}
		if _lookoutequipmentListLabelGroups {
			lookoutequipment_ListLabelGroups(cfg, client)
			return
		}
		if _lookoutequipmentListLabels {
			lookoutequipment_ListLabels(cfg, client)
			return
		}
		if _lookoutequipmentListModelVersions {
			lookoutequipment_ListModelVersions(cfg, client)
			return
		}
		if _lookoutequipmentListModels {
			lookoutequipment_ListModels(cfg, client)
			return
		}
		if _lookoutequipmentListRetrainingSchedulers {
			lookoutequipment_ListRetrainingSchedulers(cfg, client)
			return
		}
		if _lookoutequipmentListSensorStatistics {
			lookoutequipment_ListSensorStatistics(cfg, client)
			return
		}
		if _lookoutequipmentListTagsForResource {
			lookoutequipment_ListTagsForResource(cfg, client)
			return
		}
		if _lookoutequipmentPutResourcePolicy {
			lookoutequipment_PutResourcePolicy(cfg, client)
			return
		}
		if _lookoutequipmentStartDataIngestionJob {
			lookoutequipment_StartDataIngestionJob(cfg, client)
			return
		}
		if _lookoutequipmentStartInferenceScheduler {
			lookoutequipment_StartInferenceScheduler(cfg, client)
			return
		}
		if _lookoutequipmentStartRetrainingScheduler {
			lookoutequipment_StartRetrainingScheduler(cfg, client)
			return
		}
		if _lookoutequipmentStopInferenceScheduler {
			lookoutequipment_StopInferenceScheduler(cfg, client)
			return
		}
		if _lookoutequipmentStopRetrainingScheduler {
			lookoutequipment_StopRetrainingScheduler(cfg, client)
			return
		}
		if _lookoutequipmentTagResource {
			lookoutequipment_TagResource(cfg, client)
			return
		}
		if _lookoutequipmentUntagResource {
			lookoutequipment_UntagResource(cfg, client)
			return
		}
		if _lookoutequipmentUpdateActiveModelVersion {
			lookoutequipment_UpdateActiveModelVersion(cfg, client)
			return
		}
		if _lookoutequipmentUpdateInferenceScheduler {
			lookoutequipment_UpdateInferenceScheduler(cfg, client)
			return
		}
		if _lookoutequipmentUpdateLabelGroup {
			lookoutequipment_UpdateLabelGroup(cfg, client)
			return
		}
		if _lookoutequipmentUpdateModel {
			lookoutequipment_UpdateModel(cfg, client)
			return
		}
		if _lookoutequipmentUpdateRetrainingScheduler {
			lookoutequipment_UpdateRetrainingScheduler(cfg, client)
			return
		}

	},
}

var (
	_lookoutequipmentCreateDataset               bool
	_lookoutequipmentCreateInferenceScheduler    bool
	_lookoutequipmentCreateLabel                 bool
	_lookoutequipmentCreateLabelGroup            bool
	_lookoutequipmentCreateModel                 bool
	_lookoutequipmentCreateRetrainingScheduler   bool
	_lookoutequipmentDeleteDataset               bool
	_lookoutequipmentDeleteInferenceScheduler    bool
	_lookoutequipmentDeleteLabel                 bool
	_lookoutequipmentDeleteLabelGroup            bool
	_lookoutequipmentDeleteModel                 bool
	_lookoutequipmentDeleteResourcePolicy        bool
	_lookoutequipmentDeleteRetrainingScheduler   bool
	_lookoutequipmentDescribeDataIngestionJob    bool
	_lookoutequipmentDescribeDataset             bool
	_lookoutequipmentDescribeInferenceScheduler  bool
	_lookoutequipmentDescribeLabel               bool
	_lookoutequipmentDescribeLabelGroup          bool
	_lookoutequipmentDescribeModel               bool
	_lookoutequipmentDescribeModelVersion        bool
	_lookoutequipmentDescribeResourcePolicy      bool
	_lookoutequipmentDescribeRetrainingScheduler bool
	_lookoutequipmentImportDataset               bool
	_lookoutequipmentImportModelVersion          bool
	_lookoutequipmentListDataIngestionJobs       bool
	_lookoutequipmentListDatasets                bool
	_lookoutequipmentListInferenceEvents         bool
	_lookoutequipmentListInferenceExecutions     bool
	_lookoutequipmentListInferenceSchedulers     bool
	_lookoutequipmentListLabelGroups             bool
	_lookoutequipmentListLabels                  bool
	_lookoutequipmentListModelVersions           bool
	_lookoutequipmentListModels                  bool
	_lookoutequipmentListRetrainingSchedulers    bool
	_lookoutequipmentListSensorStatistics        bool
	_lookoutequipmentListTagsForResource         bool
	_lookoutequipmentPutResourcePolicy           bool
	_lookoutequipmentStartDataIngestionJob       bool
	_lookoutequipmentStartInferenceScheduler     bool
	_lookoutequipmentStartRetrainingScheduler    bool
	_lookoutequipmentStopInferenceScheduler      bool
	_lookoutequipmentStopRetrainingScheduler     bool
	_lookoutequipmentTagResource                 bool
	_lookoutequipmentUntagResource               bool
	_lookoutequipmentUpdateActiveModelVersion    bool
	_lookoutequipmentUpdateInferenceScheduler    bool
	_lookoutequipmentUpdateLabelGroup            bool
	_lookoutequipmentUpdateModel                 bool
	_lookoutequipmentUpdateRetrainingScheduler   bool

	_lookoutequipmentClientToken                         string
	_lookoutequipmentCreatedAtEndTime                    string
	_lookoutequipmentCreatedAtStartTime                  string
	_lookoutequipmentDataDelayOffsetInMinutes            string
	_lookoutequipmentDataEndTimeBefore                   string
	_lookoutequipmentDataInputConfiguration              string
	_lookoutequipmentDataOutputConfiguration             string
	_lookoutequipmentDataPreProcessingConfiguration      string
	_lookoutequipmentDataStartTimeAfter                  string
	_lookoutequipmentDataUploadFrequency                 string
	_lookoutequipmentDatasetName                         string
	_lookoutequipmentDatasetNameBeginsWith               string
	_lookoutequipmentDatasetSchema                       string
	_lookoutequipmentEndTime                             string
	_lookoutequipmentEquipment                           string
	_lookoutequipmentEvaluationDataEndTime               string
	_lookoutequipmentEvaluationDataStartTime             string
	_lookoutequipmentFaultCode                           string
	_lookoutequipmentFaultCodes                          []string
	_lookoutequipmentInferenceDataImportStrategy         string
	_lookoutequipmentInferenceSchedulerName              string
	_lookoutequipmentInferenceSchedulerNameBeginsWith    string
	_lookoutequipmentIngestionInputConfiguration         string
	_lookoutequipmentIngestionJobId                      string
	_lookoutequipmentIntervalEndTime                     string
	_lookoutequipmentIntervalStartTime                   string
	_lookoutequipmentJobId                               string
	_lookoutequipmentLabelGroupName                      string
	_lookoutequipmentLabelGroupNameBeginsWith            string
	_lookoutequipmentLabelId                             string
	_lookoutequipmentLabelsInputConfiguration            string
	_lookoutequipmentLookbackWindow                      string
	_lookoutequipmentMaxModelVersion                     string
	_lookoutequipmentMaxResults                          string
	_lookoutequipmentMinModelVersion                     string
	_lookoutequipmentModelDiagnosticsOutputConfiguration string
	_lookoutequipmentModelName                           string
	_lookoutequipmentModelNameBeginsWith                 string
	_lookoutequipmentModelVersion                        string
	_lookoutequipmentNextToken                           string
	_lookoutequipmentNotes                               string
	_lookoutequipmentOffCondition                        string
	_lookoutequipmentPolicyRevisionId                    string
	_lookoutequipmentPromoteMode                         string
	_lookoutequipmentRating                              string
	_lookoutequipmentResourceArn                         string
	_lookoutequipmentResourcePolicy                      string
	_lookoutequipmentRetrainingFrequency                 string
	_lookoutequipmentRetrainingStartDate                 string
	_lookoutequipmentRoleArn                             string
	_lookoutequipmentServerSideKmsKeyId                  string
	_lookoutequipmentSourceDatasetArn                    string
	_lookoutequipmentSourceModelVersionArn               string
	_lookoutequipmentSourceType                          string
	_lookoutequipmentStartTime                           string
	_lookoutequipmentStatus                              string
	_lookoutequipmentTagKeys                             []string
	_lookoutequipmentTags                                string
	_lookoutequipmentTrainingDataEndTime                 string
	_lookoutequipmentTrainingDataStartTime               string
)

// Creates a container for a collection of data being ingested for analysis. The
// dataset contains the metadata describing where the data is and what the data
// actually looks like. For example, it contains the location of the data source,
// the data schema, and other information. A dataset also contains any tags
// associated with the ingested data.
func lookoutequipment_CreateDataset(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.CreateDatasetInput{
		// ClientToken: *string, // Required
		// DatasetName: *string, // Required
	}

	if len(_lookoutequipmentClientToken) > 0 {
		input.ClientToken = aws.String(_lookoutequipmentClientToken)
	}
	if len(_lookoutequipmentDatasetName) > 0 {
		input.DatasetName = aws.String(_lookoutequipmentDatasetName)
	}
	if len(_lookoutequipmentDatasetSchema) > 0 {
		if err := assignInputField(input, "DatasetSchema", _lookoutequipmentDatasetSchema); err != nil {
			log.Errorf("invalid --dataset-schema: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentServerSideKmsKeyId) > 0 {
		input.ServerSideKmsKeyId = aws.String(_lookoutequipmentServerSideKmsKeyId)
	}
	if len(_lookoutequipmentTags) > 0 {
		if err := assignInputField(input, "Tags", _lookoutequipmentTags); err != nil {
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

// Creates a scheduled inference. Scheduling an inference is setting up a
// continuous real-time inference plan to analyze new measurement data. When
// setting up the schedule, you provide an S3 bucket location for the input data,
// assign it a delimiter between separate entries in the data, set an offset delay
// if desired, and set the frequency of inferencing. You must also provide an S3
// bucket location for the output data.
func lookoutequipment_CreateInferenceScheduler(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.CreateInferenceSchedulerInput{
		// ClientToken: *string, // Required
		// DataInputConfiguration: *types.InferenceInputConfiguration, // Required
		// DataOutputConfiguration: *types.InferenceOutputConfiguration, // Required
		// DataUploadFrequency: types.DataUploadFrequency, // Required
		// InferenceSchedulerName: *string, // Required
		// ModelName: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_lookoutequipmentClientToken) > 0 {
		input.ClientToken = aws.String(_lookoutequipmentClientToken)
	}
	if len(_lookoutequipmentDataInputConfiguration) > 0 {
		if err := assignInputField(input, "DataInputConfiguration", _lookoutequipmentDataInputConfiguration); err != nil {
			log.Errorf("invalid --data-input-configuration: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentDataOutputConfiguration) > 0 {
		if err := assignInputField(input, "DataOutputConfiguration", _lookoutequipmentDataOutputConfiguration); err != nil {
			log.Errorf("invalid --data-output-configuration: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentDataUploadFrequency) > 0 {
		if err := assignInputField(input, "DataUploadFrequency", _lookoutequipmentDataUploadFrequency); err != nil {
			log.Errorf("invalid --data-upload-frequency: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentInferenceSchedulerName) > 0 {
		input.InferenceSchedulerName = aws.String(_lookoutequipmentInferenceSchedulerName)
	}
	if len(_lookoutequipmentModelName) > 0 {
		input.ModelName = aws.String(_lookoutequipmentModelName)
	}
	if len(_lookoutequipmentRoleArn) > 0 {
		input.RoleArn = aws.String(_lookoutequipmentRoleArn)
	}
	if len(_lookoutequipmentDataDelayOffsetInMinutes) > 0 {
		if err := assignInputField(input, "DataDelayOffsetInMinutes", _lookoutequipmentDataDelayOffsetInMinutes); err != nil {
			log.Errorf("invalid --data-delay-offset-in-minutes: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentServerSideKmsKeyId) > 0 {
		input.ServerSideKmsKeyId = aws.String(_lookoutequipmentServerSideKmsKeyId)
	}
	if len(_lookoutequipmentTags) > 0 {
		if err := assignInputField(input, "Tags", _lookoutequipmentTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateInferenceScheduler(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a label for an event.
func lookoutequipment_CreateLabel(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.CreateLabelInput{
		// ClientToken: *string, // Required
		// EndTime: *time.Time, // Required
		// LabelGroupName: *string, // Required
		// Rating: types.LabelRating, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_lookoutequipmentClientToken) > 0 {
		input.ClientToken = aws.String(_lookoutequipmentClientToken)
	}
	if len(_lookoutequipmentEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _lookoutequipmentEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentLabelGroupName) > 0 {
		input.LabelGroupName = aws.String(_lookoutequipmentLabelGroupName)
	}
	if len(_lookoutequipmentRating) > 0 {
		if err := assignInputField(input, "Rating", _lookoutequipmentRating); err != nil {
			log.Errorf("invalid --rating: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _lookoutequipmentStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentEquipment) > 0 {
		input.Equipment = aws.String(_lookoutequipmentEquipment)
	}
	if len(_lookoutequipmentFaultCode) > 0 {
		input.FaultCode = aws.String(_lookoutequipmentFaultCode)
	}
	if len(_lookoutequipmentNotes) > 0 {
		input.Notes = aws.String(_lookoutequipmentNotes)
	}

	if resp, err := client.CreateLabel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a group of labels.
func lookoutequipment_CreateLabelGroup(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.CreateLabelGroupInput{
		// ClientToken: *string, // Required
		// LabelGroupName: *string, // Required
	}

	if len(_lookoutequipmentClientToken) > 0 {
		input.ClientToken = aws.String(_lookoutequipmentClientToken)
	}
	if len(_lookoutequipmentLabelGroupName) > 0 {
		input.LabelGroupName = aws.String(_lookoutequipmentLabelGroupName)
	}
	if len(_lookoutequipmentFaultCodes) > 0 {
		input.FaultCodes = append([]string(nil), _lookoutequipmentFaultCodes...)
	}
	if len(_lookoutequipmentTags) > 0 {
		if err := assignInputField(input, "Tags", _lookoutequipmentTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLabelGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a machine learning model for data inference.
// A machine-learning (ML) model is a mathematical model that finds patterns in
// your data. In Amazon Lookout for Equipment, the model learns the patterns of
// normal behavior and detects abnormal behavior that could be potential equipment
// failure (or maintenance events). The models are made by analyzing normal data
// and abnormalities in machine behavior that have already occurred.
//
// Your model is trained using a portion of the data from your dataset and uses
// that data to learn patterns of normal behavior and abnormal patterns that lead
// to equipment failure. Another portion of the data is used to evaluate the
// model's accuracy.
func lookoutequipment_CreateModel(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.CreateModelInput{
		// ClientToken: *string, // Required
		// DatasetName: *string, // Required
		// ModelName: *string, // Required
	}

	if len(_lookoutequipmentClientToken) > 0 {
		input.ClientToken = aws.String(_lookoutequipmentClientToken)
	}
	if len(_lookoutequipmentDatasetName) > 0 {
		input.DatasetName = aws.String(_lookoutequipmentDatasetName)
	}
	if len(_lookoutequipmentModelName) > 0 {
		input.ModelName = aws.String(_lookoutequipmentModelName)
	}
	if len(_lookoutequipmentDataPreProcessingConfiguration) > 0 {
		if err := assignInputField(input, "DataPreProcessingConfiguration", _lookoutequipmentDataPreProcessingConfiguration); err != nil {
			log.Errorf("invalid --data-pre-processing-configuration: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentDatasetSchema) > 0 {
		if err := assignInputField(input, "DatasetSchema", _lookoutequipmentDatasetSchema); err != nil {
			log.Errorf("invalid --dataset-schema: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentEvaluationDataEndTime) > 0 {
		if err := assignInputField(input, "EvaluationDataEndTime", _lookoutequipmentEvaluationDataEndTime); err != nil {
			log.Errorf("invalid --evaluation-data-end-time: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentEvaluationDataStartTime) > 0 {
		if err := assignInputField(input, "EvaluationDataStartTime", _lookoutequipmentEvaluationDataStartTime); err != nil {
			log.Errorf("invalid --evaluation-data-start-time: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentLabelsInputConfiguration) > 0 {
		if err := assignInputField(input, "LabelsInputConfiguration", _lookoutequipmentLabelsInputConfiguration); err != nil {
			log.Errorf("invalid --labels-input-configuration: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentModelDiagnosticsOutputConfiguration) > 0 {
		if err := assignInputField(input, "ModelDiagnosticsOutputConfiguration", _lookoutequipmentModelDiagnosticsOutputConfiguration); err != nil {
			log.Errorf("invalid --model-diagnostics-output-configuration: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentOffCondition) > 0 {
		input.OffCondition = aws.String(_lookoutequipmentOffCondition)
	}
	if len(_lookoutequipmentRoleArn) > 0 {
		input.RoleArn = aws.String(_lookoutequipmentRoleArn)
	}
	if len(_lookoutequipmentServerSideKmsKeyId) > 0 {
		input.ServerSideKmsKeyId = aws.String(_lookoutequipmentServerSideKmsKeyId)
	}
	if len(_lookoutequipmentTags) > 0 {
		if err := assignInputField(input, "Tags", _lookoutequipmentTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentTrainingDataEndTime) > 0 {
		if err := assignInputField(input, "TrainingDataEndTime", _lookoutequipmentTrainingDataEndTime); err != nil {
			log.Errorf("invalid --training-data-end-time: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentTrainingDataStartTime) > 0 {
		if err := assignInputField(input, "TrainingDataStartTime", _lookoutequipmentTrainingDataStartTime); err != nil {
			log.Errorf("invalid --training-data-start-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a retraining scheduler on the specified model.
func lookoutequipment_CreateRetrainingScheduler(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.CreateRetrainingSchedulerInput{
		// ClientToken: *string, // Required
		// LookbackWindow: *string, // Required
		// ModelName: *string, // Required
		// RetrainingFrequency: *string, // Required
	}

	if len(_lookoutequipmentClientToken) > 0 {
		input.ClientToken = aws.String(_lookoutequipmentClientToken)
	}
	if len(_lookoutequipmentLookbackWindow) > 0 {
		input.LookbackWindow = aws.String(_lookoutequipmentLookbackWindow)
	}
	if len(_lookoutequipmentModelName) > 0 {
		input.ModelName = aws.String(_lookoutequipmentModelName)
	}
	if len(_lookoutequipmentRetrainingFrequency) > 0 {
		input.RetrainingFrequency = aws.String(_lookoutequipmentRetrainingFrequency)
	}
	if len(_lookoutequipmentPromoteMode) > 0 {
		if err := assignInputField(input, "PromoteMode", _lookoutequipmentPromoteMode); err != nil {
			log.Errorf("invalid --promote-mode: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentRetrainingStartDate) > 0 {
		if err := assignInputField(input, "RetrainingStartDate", _lookoutequipmentRetrainingStartDate); err != nil {
			log.Errorf("invalid --retraining-start-date: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRetrainingScheduler(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a dataset and associated artifacts. The operation will check to see if
// any inference scheduler or data ingestion job is currently using the dataset,
// and if there isn't, the dataset, its metadata, and any associated data stored in
// S3 will be deleted. This does not affect any models that used this dataset for
// training and evaluation, but does prevent it from being used in the future.
func lookoutequipment_DeleteDataset(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.DeleteDatasetInput{
		// DatasetName: *string, // Required
	}

	if len(_lookoutequipmentDatasetName) > 0 {
		input.DatasetName = aws.String(_lookoutequipmentDatasetName)
	}

	if resp, err := client.DeleteDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an inference scheduler that has been set up. Prior inference results
// will not be deleted.
func lookoutequipment_DeleteInferenceScheduler(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.DeleteInferenceSchedulerInput{
		// InferenceSchedulerName: *string, // Required
	}

	if len(_lookoutequipmentInferenceSchedulerName) > 0 {
		input.InferenceSchedulerName = aws.String(_lookoutequipmentInferenceSchedulerName)
	}

	if resp, err := client.DeleteInferenceScheduler(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a label.
func lookoutequipment_DeleteLabel(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.DeleteLabelInput{
		// LabelGroupName: *string, // Required
		// LabelId: *string, // Required
	}

	if len(_lookoutequipmentLabelGroupName) > 0 {
		input.LabelGroupName = aws.String(_lookoutequipmentLabelGroupName)
	}
	if len(_lookoutequipmentLabelId) > 0 {
		input.LabelId = aws.String(_lookoutequipmentLabelId)
	}

	if resp, err := client.DeleteLabel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a group of labels.
func lookoutequipment_DeleteLabelGroup(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.DeleteLabelGroupInput{
		// LabelGroupName: *string, // Required
	}

	if len(_lookoutequipmentLabelGroupName) > 0 {
		input.LabelGroupName = aws.String(_lookoutequipmentLabelGroupName)
	}

	if resp, err := client.DeleteLabelGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a machine learning model currently available for Amazon Lookout for
// Equipment. This will prevent it from being used with an inference scheduler,
// even one that is already set up.
func lookoutequipment_DeleteModel(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.DeleteModelInput{
		// ModelName: *string, // Required
	}

	if len(_lookoutequipmentModelName) > 0 {
		input.ModelName = aws.String(_lookoutequipmentModelName)
	}

	if resp, err := client.DeleteModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the resource policy attached to the resource.
func lookoutequipment_DeleteResourcePolicy(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.DeleteResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_lookoutequipmentResourceArn) > 0 {
		input.ResourceArn = aws.String(_lookoutequipmentResourceArn)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a retraining scheduler from a model. The retraining scheduler must be
// in the STOPPED status.
func lookoutequipment_DeleteRetrainingScheduler(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.DeleteRetrainingSchedulerInput{
		// ModelName: *string, // Required
	}

	if len(_lookoutequipmentModelName) > 0 {
		input.ModelName = aws.String(_lookoutequipmentModelName)
	}

	if resp, err := client.DeleteRetrainingScheduler(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information on a specific data ingestion job such as creation time,
// dataset ARN, and status.
func lookoutequipment_DescribeDataIngestionJob(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.DescribeDataIngestionJobInput{
		// JobId: *string, // Required
	}

	if len(_lookoutequipmentJobId) > 0 {
		input.JobId = aws.String(_lookoutequipmentJobId)
	}

	if resp, err := client.DescribeDataIngestionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a JSON description of the data in each time series dataset, including
// names, column names, and data types.
func lookoutequipment_DescribeDataset(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.DescribeDatasetInput{
		// DatasetName: *string, // Required
	}

	if len(_lookoutequipmentDatasetName) > 0 {
		input.DatasetName = aws.String(_lookoutequipmentDatasetName)
	}

	if resp, err := client.DescribeDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specifies information about the inference scheduler being used, including
// name, model, status, and associated metadata
func lookoutequipment_DescribeInferenceScheduler(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.DescribeInferenceSchedulerInput{
		// InferenceSchedulerName: *string, // Required
	}

	if len(_lookoutequipmentInferenceSchedulerName) > 0 {
		input.InferenceSchedulerName = aws.String(_lookoutequipmentInferenceSchedulerName)
	}

	if resp, err := client.DescribeInferenceScheduler(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the name of the label.
func lookoutequipment_DescribeLabel(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.DescribeLabelInput{
		// LabelGroupName: *string, // Required
		// LabelId: *string, // Required
	}

	if len(_lookoutequipmentLabelGroupName) > 0 {
		input.LabelGroupName = aws.String(_lookoutequipmentLabelGroupName)
	}
	if len(_lookoutequipmentLabelId) > 0 {
		input.LabelId = aws.String(_lookoutequipmentLabelId)
	}

	if resp, err := client.DescribeLabel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the label group.
func lookoutequipment_DescribeLabelGroup(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.DescribeLabelGroupInput{
		// LabelGroupName: *string, // Required
	}

	if len(_lookoutequipmentLabelGroupName) > 0 {
		input.LabelGroupName = aws.String(_lookoutequipmentLabelGroupName)
	}

	if resp, err := client.DescribeLabelGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a JSON containing the overall information about a specific machine
// learning model, including model name and ARN, dataset, training and evaluation
// information, status, and so on.
func lookoutequipment_DescribeModel(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.DescribeModelInput{
		// ModelName: *string, // Required
	}

	if len(_lookoutequipmentModelName) > 0 {
		input.ModelName = aws.String(_lookoutequipmentModelName)
	}

	if resp, err := client.DescribeModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a specific machine learning model version.
func lookoutequipment_DescribeModelVersion(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.DescribeModelVersionInput{
		// ModelName: *string, // Required
		// ModelVersion: *int64, // Required
	}

	if len(_lookoutequipmentModelName) > 0 {
		input.ModelName = aws.String(_lookoutequipmentModelName)
	}
	if len(_lookoutequipmentModelVersion) > 0 {
		if err := assignInputField(input, "ModelVersion", _lookoutequipmentModelVersion); err != nil {
			log.Errorf("invalid --model-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeModelVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides the details of a resource policy attached to a resource.
func lookoutequipment_DescribeResourcePolicy(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.DescribeResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_lookoutequipmentResourceArn) > 0 {
		input.ResourceArn = aws.String(_lookoutequipmentResourceArn)
	}

	if resp, err := client.DescribeResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a description of the retraining scheduler, including information such
// as the model name and retraining parameters.
func lookoutequipment_DescribeRetrainingScheduler(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.DescribeRetrainingSchedulerInput{
		// ModelName: *string, // Required
	}

	if len(_lookoutequipmentModelName) > 0 {
		input.ModelName = aws.String(_lookoutequipmentModelName)
	}

	if resp, err := client.DescribeRetrainingScheduler(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports a dataset.
func lookoutequipment_ImportDataset(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.ImportDatasetInput{
		// ClientToken: *string, // Required
		// SourceDatasetArn: *string, // Required
	}

	if len(_lookoutequipmentClientToken) > 0 {
		input.ClientToken = aws.String(_lookoutequipmentClientToken)
	}
	if len(_lookoutequipmentSourceDatasetArn) > 0 {
		input.SourceDatasetArn = aws.String(_lookoutequipmentSourceDatasetArn)
	}
	if len(_lookoutequipmentDatasetName) > 0 {
		input.DatasetName = aws.String(_lookoutequipmentDatasetName)
	}
	if len(_lookoutequipmentServerSideKmsKeyId) > 0 {
		input.ServerSideKmsKeyId = aws.String(_lookoutequipmentServerSideKmsKeyId)
	}
	if len(_lookoutequipmentTags) > 0 {
		if err := assignInputField(input, "Tags", _lookoutequipmentTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports a model that has been trained successfully.
func lookoutequipment_ImportModelVersion(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.ImportModelVersionInput{
		// ClientToken: *string, // Required
		// DatasetName: *string, // Required
		// SourceModelVersionArn: *string, // Required
	}

	if len(_lookoutequipmentClientToken) > 0 {
		input.ClientToken = aws.String(_lookoutequipmentClientToken)
	}
	if len(_lookoutequipmentDatasetName) > 0 {
		input.DatasetName = aws.String(_lookoutequipmentDatasetName)
	}
	if len(_lookoutequipmentSourceModelVersionArn) > 0 {
		input.SourceModelVersionArn = aws.String(_lookoutequipmentSourceModelVersionArn)
	}
	if len(_lookoutequipmentInferenceDataImportStrategy) > 0 {
		if err := assignInputField(input, "InferenceDataImportStrategy", _lookoutequipmentInferenceDataImportStrategy); err != nil {
			log.Errorf("invalid --inference-data-import-strategy: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentLabelsInputConfiguration) > 0 {
		if err := assignInputField(input, "LabelsInputConfiguration", _lookoutequipmentLabelsInputConfiguration); err != nil {
			log.Errorf("invalid --labels-input-configuration: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentModelName) > 0 {
		input.ModelName = aws.String(_lookoutequipmentModelName)
	}
	if len(_lookoutequipmentRoleArn) > 0 {
		input.RoleArn = aws.String(_lookoutequipmentRoleArn)
	}
	if len(_lookoutequipmentServerSideKmsKeyId) > 0 {
		input.ServerSideKmsKeyId = aws.String(_lookoutequipmentServerSideKmsKeyId)
	}
	if len(_lookoutequipmentTags) > 0 {
		if err := assignInputField(input, "Tags", _lookoutequipmentTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportModelVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a list of all data ingestion jobs, including dataset name and ARN, S3
// location of the input data, status, and so on.
func lookoutequipment_ListDataIngestionJobs(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.ListDataIngestionJobsInput{}

	if len(_lookoutequipmentDatasetName) > 0 {
		input.DatasetName = aws.String(_lookoutequipmentDatasetName)
	}
	if len(_lookoutequipmentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lookoutequipmentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentNextToken) > 0 {
		input.NextToken = aws.String(_lookoutequipmentNextToken)
	}
	if len(_lookoutequipmentStatus) > 0 {
		if err := assignInputField(input, "Status", _lookoutequipmentStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDataIngestionJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lookoutequipment.ListDataIngestionJobsOutput
	p := lookoutequipment.NewListDataIngestionJobsPaginator(client, input)
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

// Lists all datasets currently available in your account, filtering on the
// dataset name.
func lookoutequipment_ListDatasets(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.ListDatasetsInput{}

	if len(_lookoutequipmentDatasetNameBeginsWith) > 0 {
		input.DatasetNameBeginsWith = aws.String(_lookoutequipmentDatasetNameBeginsWith)
	}
	if len(_lookoutequipmentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lookoutequipmentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentNextToken) > 0 {
		input.NextToken = aws.String(_lookoutequipmentNextToken)
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

	var results []*lookoutequipment.ListDatasetsOutput
	p := lookoutequipment.NewListDatasetsPaginator(client, input)
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

// Lists all inference events that have been found for the specified inference
// scheduler.
func lookoutequipment_ListInferenceEvents(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.ListInferenceEventsInput{
		// InferenceSchedulerName: *string, // Required
		// IntervalEndTime: *time.Time, // Required
		// IntervalStartTime: *time.Time, // Required
	}

	if len(_lookoutequipmentInferenceSchedulerName) > 0 {
		input.InferenceSchedulerName = aws.String(_lookoutequipmentInferenceSchedulerName)
	}
	if len(_lookoutequipmentIntervalEndTime) > 0 {
		if err := assignInputField(input, "IntervalEndTime", _lookoutequipmentIntervalEndTime); err != nil {
			log.Errorf("invalid --interval-end-time: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentIntervalStartTime) > 0 {
		if err := assignInputField(input, "IntervalStartTime", _lookoutequipmentIntervalStartTime); err != nil {
			log.Errorf("invalid --interval-start-time: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lookoutequipmentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentNextToken) > 0 {
		input.NextToken = aws.String(_lookoutequipmentNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInferenceEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lookoutequipment.ListInferenceEventsOutput
	p := lookoutequipment.NewListInferenceEventsPaginator(client, input)
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

// Lists all inference executions that have been performed by the specified
// inference scheduler.
func lookoutequipment_ListInferenceExecutions(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.ListInferenceExecutionsInput{
		// InferenceSchedulerName: *string, // Required
	}

	if len(_lookoutequipmentInferenceSchedulerName) > 0 {
		input.InferenceSchedulerName = aws.String(_lookoutequipmentInferenceSchedulerName)
	}
	if len(_lookoutequipmentDataEndTimeBefore) > 0 {
		if err := assignInputField(input, "DataEndTimeBefore", _lookoutequipmentDataEndTimeBefore); err != nil {
			log.Errorf("invalid --data-end-time-before: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentDataStartTimeAfter) > 0 {
		if err := assignInputField(input, "DataStartTimeAfter", _lookoutequipmentDataStartTimeAfter); err != nil {
			log.Errorf("invalid --data-start-time-after: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lookoutequipmentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentNextToken) > 0 {
		input.NextToken = aws.String(_lookoutequipmentNextToken)
	}
	if len(_lookoutequipmentStatus) > 0 {
		if err := assignInputField(input, "Status", _lookoutequipmentStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListInferenceExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lookoutequipment.ListInferenceExecutionsOutput
	p := lookoutequipment.NewListInferenceExecutionsPaginator(client, input)
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

// Retrieves a list of all inference schedulers currently available for your
// account.
func lookoutequipment_ListInferenceSchedulers(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.ListInferenceSchedulersInput{}

	if len(_lookoutequipmentInferenceSchedulerNameBeginsWith) > 0 {
		input.InferenceSchedulerNameBeginsWith = aws.String(_lookoutequipmentInferenceSchedulerNameBeginsWith)
	}
	if len(_lookoutequipmentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lookoutequipmentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentModelName) > 0 {
		input.ModelName = aws.String(_lookoutequipmentModelName)
	}
	if len(_lookoutequipmentNextToken) > 0 {
		input.NextToken = aws.String(_lookoutequipmentNextToken)
	}
	if len(_lookoutequipmentStatus) > 0 {
		if err := assignInputField(input, "Status", _lookoutequipmentStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListInferenceSchedulers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lookoutequipment.ListInferenceSchedulersOutput
	p := lookoutequipment.NewListInferenceSchedulersPaginator(client, input)
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

// Returns a list of the label groups.
func lookoutequipment_ListLabelGroups(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.ListLabelGroupsInput{}

	if len(_lookoutequipmentLabelGroupNameBeginsWith) > 0 {
		input.LabelGroupNameBeginsWith = aws.String(_lookoutequipmentLabelGroupNameBeginsWith)
	}
	if len(_lookoutequipmentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lookoutequipmentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentNextToken) > 0 {
		input.NextToken = aws.String(_lookoutequipmentNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLabelGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lookoutequipment.ListLabelGroupsOutput
	p := lookoutequipment.NewListLabelGroupsPaginator(client, input)
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

// Provides a list of labels.
func lookoutequipment_ListLabels(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.ListLabelsInput{
		// LabelGroupName: *string, // Required
	}

	if len(_lookoutequipmentLabelGroupName) > 0 {
		input.LabelGroupName = aws.String(_lookoutequipmentLabelGroupName)
	}
	if len(_lookoutequipmentEquipment) > 0 {
		input.Equipment = aws.String(_lookoutequipmentEquipment)
	}
	if len(_lookoutequipmentFaultCode) > 0 {
		input.FaultCode = aws.String(_lookoutequipmentFaultCode)
	}
	if len(_lookoutequipmentIntervalEndTime) > 0 {
		if err := assignInputField(input, "IntervalEndTime", _lookoutequipmentIntervalEndTime); err != nil {
			log.Errorf("invalid --interval-end-time: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentIntervalStartTime) > 0 {
		if err := assignInputField(input, "IntervalStartTime", _lookoutequipmentIntervalStartTime); err != nil {
			log.Errorf("invalid --interval-start-time: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lookoutequipmentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentNextToken) > 0 {
		input.NextToken = aws.String(_lookoutequipmentNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLabels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lookoutequipment.ListLabelsOutput
	p := lookoutequipment.NewListLabelsPaginator(client, input)
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

// Generates a list of all model versions for a given model, including the model
// version, model version ARN, and status. To list a subset of versions, use the
// MaxModelVersion and MinModelVersion fields.
func lookoutequipment_ListModelVersions(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.ListModelVersionsInput{
		// ModelName: *string, // Required
	}

	if len(_lookoutequipmentModelName) > 0 {
		input.ModelName = aws.String(_lookoutequipmentModelName)
	}
	if len(_lookoutequipmentCreatedAtEndTime) > 0 {
		if err := assignInputField(input, "CreatedAtEndTime", _lookoutequipmentCreatedAtEndTime); err != nil {
			log.Errorf("invalid --created-at-end-time: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentCreatedAtStartTime) > 0 {
		if err := assignInputField(input, "CreatedAtStartTime", _lookoutequipmentCreatedAtStartTime); err != nil {
			log.Errorf("invalid --created-at-start-time: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentMaxModelVersion) > 0 {
		if err := assignInputField(input, "MaxModelVersion", _lookoutequipmentMaxModelVersion); err != nil {
			log.Errorf("invalid --max-model-version: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lookoutequipmentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentMinModelVersion) > 0 {
		if err := assignInputField(input, "MinModelVersion", _lookoutequipmentMinModelVersion); err != nil {
			log.Errorf("invalid --min-model-version: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentNextToken) > 0 {
		input.NextToken = aws.String(_lookoutequipmentNextToken)
	}
	if len(_lookoutequipmentSourceType) > 0 {
		if err := assignInputField(input, "SourceType", _lookoutequipmentSourceType); err != nil {
			log.Errorf("invalid --source-type: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentStatus) > 0 {
		if err := assignInputField(input, "Status", _lookoutequipmentStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListModelVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lookoutequipment.ListModelVersionsOutput
	p := lookoutequipment.NewListModelVersionsPaginator(client, input)
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

// Generates a list of all models in the account, including model name and ARN,
// dataset, and status.
func lookoutequipment_ListModels(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.ListModelsInput{}

	if len(_lookoutequipmentDatasetNameBeginsWith) > 0 {
		input.DatasetNameBeginsWith = aws.String(_lookoutequipmentDatasetNameBeginsWith)
	}
	if len(_lookoutequipmentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lookoutequipmentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentModelNameBeginsWith) > 0 {
		input.ModelNameBeginsWith = aws.String(_lookoutequipmentModelNameBeginsWith)
	}
	if len(_lookoutequipmentNextToken) > 0 {
		input.NextToken = aws.String(_lookoutequipmentNextToken)
	}
	if len(_lookoutequipmentStatus) > 0 {
		if err := assignInputField(input, "Status", _lookoutequipmentStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListModels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lookoutequipment.ListModelsOutput
	p := lookoutequipment.NewListModelsPaginator(client, input)
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

// Lists all retraining schedulers in your account, filtering by model name prefix
// and status.
func lookoutequipment_ListRetrainingSchedulers(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.ListRetrainingSchedulersInput{}

	if len(_lookoutequipmentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lookoutequipmentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentModelNameBeginsWith) > 0 {
		input.ModelNameBeginsWith = aws.String(_lookoutequipmentModelNameBeginsWith)
	}
	if len(_lookoutequipmentNextToken) > 0 {
		input.NextToken = aws.String(_lookoutequipmentNextToken)
	}
	if len(_lookoutequipmentStatus) > 0 {
		if err := assignInputField(input, "Status", _lookoutequipmentStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRetrainingSchedulers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lookoutequipment.ListRetrainingSchedulersOutput
	p := lookoutequipment.NewListRetrainingSchedulersPaginator(client, input)
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

// Lists statistics about the data collected for each of the sensors that have
// been successfully ingested in the particular dataset. Can also be used to
// retreive Sensor Statistics for a previous ingestion job.
func lookoutequipment_ListSensorStatistics(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.ListSensorStatisticsInput{
		// DatasetName: *string, // Required
	}

	if len(_lookoutequipmentDatasetName) > 0 {
		input.DatasetName = aws.String(_lookoutequipmentDatasetName)
	}
	if len(_lookoutequipmentIngestionJobId) > 0 {
		input.IngestionJobId = aws.String(_lookoutequipmentIngestionJobId)
	}
	if len(_lookoutequipmentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lookoutequipmentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentNextToken) > 0 {
		input.NextToken = aws.String(_lookoutequipmentNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSensorStatistics(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lookoutequipment.ListSensorStatisticsOutput
	p := lookoutequipment.NewListSensorStatisticsPaginator(client, input)
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

// Lists all the tags for a specified resource, including key and value.
func lookoutequipment_ListTagsForResource(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_lookoutequipmentResourceArn) > 0 {
		input.ResourceArn = aws.String(_lookoutequipmentResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a resource control policy for a given resource.
func lookoutequipment_PutResourcePolicy(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.PutResourcePolicyInput{
		// ClientToken: *string, // Required
		// ResourceArn: *string, // Required
		// ResourcePolicy: *string, // Required
	}

	if len(_lookoutequipmentClientToken) > 0 {
		input.ClientToken = aws.String(_lookoutequipmentClientToken)
	}
	if len(_lookoutequipmentResourceArn) > 0 {
		input.ResourceArn = aws.String(_lookoutequipmentResourceArn)
	}
	if len(_lookoutequipmentResourcePolicy) > 0 {
		input.ResourcePolicy = aws.String(_lookoutequipmentResourcePolicy)
	}
	if len(_lookoutequipmentPolicyRevisionId) > 0 {
		input.PolicyRevisionId = aws.String(_lookoutequipmentPolicyRevisionId)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a data ingestion job. Amazon Lookout for Equipment returns the job
// status.
func lookoutequipment_StartDataIngestionJob(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.StartDataIngestionJobInput{
		// ClientToken: *string, // Required
		// DatasetName: *string, // Required
		// IngestionInputConfiguration: *types.IngestionInputConfiguration, // Required
		// RoleArn: *string, // Required
	}

	if len(_lookoutequipmentClientToken) > 0 {
		input.ClientToken = aws.String(_lookoutequipmentClientToken)
	}
	if len(_lookoutequipmentDatasetName) > 0 {
		input.DatasetName = aws.String(_lookoutequipmentDatasetName)
	}
	if len(_lookoutequipmentIngestionInputConfiguration) > 0 {
		if err := assignInputField(input, "IngestionInputConfiguration", _lookoutequipmentIngestionInputConfiguration); err != nil {
			log.Errorf("invalid --ingestion-input-configuration: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentRoleArn) > 0 {
		input.RoleArn = aws.String(_lookoutequipmentRoleArn)
	}

	if resp, err := client.StartDataIngestionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an inference scheduler.
func lookoutequipment_StartInferenceScheduler(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.StartInferenceSchedulerInput{
		// InferenceSchedulerName: *string, // Required
	}

	if len(_lookoutequipmentInferenceSchedulerName) > 0 {
		input.InferenceSchedulerName = aws.String(_lookoutequipmentInferenceSchedulerName)
	}

	if resp, err := client.StartInferenceScheduler(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a retraining scheduler.
func lookoutequipment_StartRetrainingScheduler(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.StartRetrainingSchedulerInput{
		// ModelName: *string, // Required
	}

	if len(_lookoutequipmentModelName) > 0 {
		input.ModelName = aws.String(_lookoutequipmentModelName)
	}

	if resp, err := client.StartRetrainingScheduler(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an inference scheduler.
func lookoutequipment_StopInferenceScheduler(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.StopInferenceSchedulerInput{
		// InferenceSchedulerName: *string, // Required
	}

	if len(_lookoutequipmentInferenceSchedulerName) > 0 {
		input.InferenceSchedulerName = aws.String(_lookoutequipmentInferenceSchedulerName)
	}

	if resp, err := client.StopInferenceScheduler(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a retraining scheduler.
func lookoutequipment_StopRetrainingScheduler(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.StopRetrainingSchedulerInput{
		// ModelName: *string, // Required
	}

	if len(_lookoutequipmentModelName) > 0 {
		input.ModelName = aws.String(_lookoutequipmentModelName)
	}

	if resp, err := client.StopRetrainingScheduler(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a given tag to a resource in your account. A tag is a key-value pair
// which can be added to an Amazon Lookout for Equipment resource as metadata. Tags
// can be used for organizing your resources as well as helping you to search and
// filter by tag. Multiple tags can be added to a resource, either when you create
// it, or later. Up to 50 tags can be associated with each resource.
func lookoutequipment_TagResource(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_lookoutequipmentResourceArn) > 0 {
		input.ResourceArn = aws.String(_lookoutequipmentResourceArn)
	}
	if len(_lookoutequipmentTags) > 0 {
		if err := assignInputField(input, "Tags", _lookoutequipmentTags); err != nil {
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

// Removes a specific tag from a given resource. The tag is specified by its key.
func lookoutequipment_UntagResource(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_lookoutequipmentResourceArn) > 0 {
		input.ResourceArn = aws.String(_lookoutequipmentResourceArn)
	}
	if len(_lookoutequipmentTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _lookoutequipmentTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the active model version for a given machine learning model.
func lookoutequipment_UpdateActiveModelVersion(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.UpdateActiveModelVersionInput{
		// ModelName: *string, // Required
		// ModelVersion: *int64, // Required
	}

	if len(_lookoutequipmentModelName) > 0 {
		input.ModelName = aws.String(_lookoutequipmentModelName)
	}
	if len(_lookoutequipmentModelVersion) > 0 {
		if err := assignInputField(input, "ModelVersion", _lookoutequipmentModelVersion); err != nil {
			log.Errorf("invalid --model-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateActiveModelVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an inference scheduler.
func lookoutequipment_UpdateInferenceScheduler(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.UpdateInferenceSchedulerInput{
		// InferenceSchedulerName: *string, // Required
	}

	if len(_lookoutequipmentInferenceSchedulerName) > 0 {
		input.InferenceSchedulerName = aws.String(_lookoutequipmentInferenceSchedulerName)
	}
	if len(_lookoutequipmentDataDelayOffsetInMinutes) > 0 {
		if err := assignInputField(input, "DataDelayOffsetInMinutes", _lookoutequipmentDataDelayOffsetInMinutes); err != nil {
			log.Errorf("invalid --data-delay-offset-in-minutes: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentDataInputConfiguration) > 0 {
		if err := assignInputField(input, "DataInputConfiguration", _lookoutequipmentDataInputConfiguration); err != nil {
			log.Errorf("invalid --data-input-configuration: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentDataOutputConfiguration) > 0 {
		if err := assignInputField(input, "DataOutputConfiguration", _lookoutequipmentDataOutputConfiguration); err != nil {
			log.Errorf("invalid --data-output-configuration: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentDataUploadFrequency) > 0 {
		if err := assignInputField(input, "DataUploadFrequency", _lookoutequipmentDataUploadFrequency); err != nil {
			log.Errorf("invalid --data-upload-frequency: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentRoleArn) > 0 {
		input.RoleArn = aws.String(_lookoutequipmentRoleArn)
	}

	if resp, err := client.UpdateInferenceScheduler(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the label group.
func lookoutequipment_UpdateLabelGroup(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.UpdateLabelGroupInput{
		// LabelGroupName: *string, // Required
	}

	if len(_lookoutequipmentLabelGroupName) > 0 {
		input.LabelGroupName = aws.String(_lookoutequipmentLabelGroupName)
	}
	if len(_lookoutequipmentFaultCodes) > 0 {
		input.FaultCodes = append([]string(nil), _lookoutequipmentFaultCodes...)
	}

	if resp, err := client.UpdateLabelGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a model in the account.
func lookoutequipment_UpdateModel(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.UpdateModelInput{
		// ModelName: *string, // Required
	}

	if len(_lookoutequipmentModelName) > 0 {
		input.ModelName = aws.String(_lookoutequipmentModelName)
	}
	if len(_lookoutequipmentLabelsInputConfiguration) > 0 {
		if err := assignInputField(input, "LabelsInputConfiguration", _lookoutequipmentLabelsInputConfiguration); err != nil {
			log.Errorf("invalid --labels-input-configuration: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentModelDiagnosticsOutputConfiguration) > 0 {
		if err := assignInputField(input, "ModelDiagnosticsOutputConfiguration", _lookoutequipmentModelDiagnosticsOutputConfiguration); err != nil {
			log.Errorf("invalid --model-diagnostics-output-configuration: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentRoleArn) > 0 {
		input.RoleArn = aws.String(_lookoutequipmentRoleArn)
	}

	if resp, err := client.UpdateModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a retraining scheduler.
func lookoutequipment_UpdateRetrainingScheduler(cfg aws.Config, client *lookoutequipment.Client) {
	input := &lookoutequipment.UpdateRetrainingSchedulerInput{
		// ModelName: *string, // Required
	}

	if len(_lookoutequipmentModelName) > 0 {
		input.ModelName = aws.String(_lookoutequipmentModelName)
	}
	if len(_lookoutequipmentLookbackWindow) > 0 {
		input.LookbackWindow = aws.String(_lookoutequipmentLookbackWindow)
	}
	if len(_lookoutequipmentPromoteMode) > 0 {
		if err := assignInputField(input, "PromoteMode", _lookoutequipmentPromoteMode); err != nil {
			log.Errorf("invalid --promote-mode: %s", err.Error())
			return
		}
	}
	if len(_lookoutequipmentRetrainingFrequency) > 0 {
		input.RetrainingFrequency = aws.String(_lookoutequipmentRetrainingFrequency)
	}
	if len(_lookoutequipmentRetrainingStartDate) > 0 {
		if err := assignInputField(input, "RetrainingStartDate", _lookoutequipmentRetrainingStartDate); err != nil {
			log.Errorf("invalid --retraining-start-date: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRetrainingScheduler(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_lookoutequipmentCmd)
	_lookoutequipmentCmd.Flags().SortFlags = false

	_lookoutequipmentCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_lookoutequipmentCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_lookoutequipmentCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentClientToken, "client-token", "", "", "Client Token")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentCreatedAtEndTime, "created-at-end-time", "", "", "Created At End Time")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentCreatedAtStartTime, "created-at-start-time", "", "", "Created At Start Time")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentDataDelayOffsetInMinutes, "data-delay-offset-in-minutes", "", "", "Data Delay Offset In Minutes")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentDataEndTimeBefore, "data-end-time-before", "", "", "Data End Time Before")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentDataInputConfiguration, "data-input-configuration", "", "", "Data Input Configuration")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentDataOutputConfiguration, "data-output-configuration", "", "", "Data Output Configuration")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentDataPreProcessingConfiguration, "data-pre-processing-configuration", "", "", "Data Pre Processing Configuration")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentDataStartTimeAfter, "data-start-time-after", "", "", "Data Start Time After")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentDataUploadFrequency, "data-upload-frequency", "", "", "Data Upload Frequency")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentDatasetName, "dataset-name", "", "", "Dataset Name")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentDatasetNameBeginsWith, "dataset-name-begins-with", "", "", "Dataset Name Begins With")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentDatasetSchema, "dataset-schema", "", "", "Dataset Schema")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentEndTime, "end-time", "", "", "End Time")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentEquipment, "equipment", "", "", "Equipment")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentEvaluationDataEndTime, "evaluation-data-end-time", "", "", "Evaluation Data End Time")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentEvaluationDataStartTime, "evaluation-data-start-time", "", "", "Evaluation Data Start Time")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentFaultCode, "fault-code", "", "", "Fault Code")
	_lookoutequipmentCmd.Flags().StringSliceVarP(&_lookoutequipmentFaultCodes, "fault-codes", "", nil, "Fault Codes")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentInferenceDataImportStrategy, "inference-data-import-strategy", "", "", "Inference Data Import Strategy")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentInferenceSchedulerName, "inference-scheduler-name", "", "", "Inference Scheduler Name")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentInferenceSchedulerNameBeginsWith, "inference-scheduler-name-begins-with", "", "", "Inference Scheduler Name Begins With")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentIngestionInputConfiguration, "ingestion-input-configuration", "", "", "Ingestion Input Configuration")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentIngestionJobId, "ingestion-job-id", "", "", "Ingestion Job ID")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentIntervalEndTime, "interval-end-time", "", "", "Interval End Time")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentIntervalStartTime, "interval-start-time", "", "", "Interval Start Time")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentJobId, "job-id", "", "", "Job ID")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentLabelGroupName, "label-group-name", "", "", "Label Group Name")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentLabelGroupNameBeginsWith, "label-group-name-begins-with", "", "", "Label Group Name Begins With")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentLabelId, "label-id", "", "", "Label ID")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentLabelsInputConfiguration, "labels-input-configuration", "", "", "Labels Input Configuration")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentLookbackWindow, "lookback-window", "", "", "Lookback Window")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentMaxModelVersion, "max-model-version", "", "", "Max Model Version")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentMaxResults, "max-results", "", "", "Max Results")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentMinModelVersion, "min-model-version", "", "", "Min Model Version")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentModelDiagnosticsOutputConfiguration, "model-diagnostics-output-configuration", "", "", "Model Diagnostics Output Configuration")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentModelName, "model-name", "", "", "Model Name")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentModelNameBeginsWith, "model-name-begins-with", "", "", "Model Name Begins With")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentModelVersion, "model-version", "", "", "Model Version")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentNextToken, "next-token", "", "", "Next Token")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentNotes, "notes", "", "", "Notes")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentOffCondition, "off-condition", "", "", "Off Condition")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentPolicyRevisionId, "policy-revision-id", "", "", "Policy Revision ID")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentPromoteMode, "promote-mode", "", "", "Promote Mode")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentRating, "rating", "", "", "Rating")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentResourceArn, "resource-arn", "", "", "Resource ARN")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentResourcePolicy, "resource-policy", "", "", "Resource Policy")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentRetrainingFrequency, "retraining-frequency", "", "", "Retraining Frequency")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentRetrainingStartDate, "retraining-start-date", "", "", "Retraining Start Date")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentRoleArn, "role-arn", "", "", "Role ARN")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentServerSideKmsKeyId, "server-side-kms-key-id", "", "", "Server Side KMS Key ID")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentSourceDatasetArn, "source-dataset-arn", "", "", "Source Dataset ARN")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentSourceModelVersionArn, "source-model-version-arn", "", "", "Source Model Version ARN")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentSourceType, "source-type", "", "", "Source Type")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentStartTime, "start-time", "", "", "Start Time")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentStatus, "status", "", "", "Status")
	_lookoutequipmentCmd.Flags().StringSliceVarP(&_lookoutequipmentTagKeys, "tag-keys", "", nil, "Tag Keys")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentTags, "tags", "", "", "Tags")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentTrainingDataEndTime, "training-data-end-time", "", "", "Training Data End Time")
	_lookoutequipmentCmd.Flags().StringVarP(&_lookoutequipmentTrainingDataStartTime, "training-data-start-time", "", "", "Training Data Start Time")

	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentCreateDataset, "create-dataset", "", false, "Create Dataset")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentCreateInferenceScheduler, "create-inference-scheduler", "", false, "Create Inference Scheduler")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentCreateLabel, "create-label", "", false, "Create Label")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentCreateLabelGroup, "create-label-group", "", false, "Create Label Group")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentCreateModel, "create-model", "", false, "Create Model")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentCreateRetrainingScheduler, "create-retraining-scheduler", "", false, "Create Retraining Scheduler")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentDeleteDataset, "delete-dataset", "", false, "Delete Dataset")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentDeleteInferenceScheduler, "delete-inference-scheduler", "", false, "Delete Inference Scheduler")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentDeleteLabel, "delete-label", "", false, "Delete Label")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentDeleteLabelGroup, "delete-label-group", "", false, "Delete Label Group")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentDeleteModel, "delete-model", "", false, "Delete Model")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentDeleteRetrainingScheduler, "delete-retraining-scheduler", "", false, "Delete Retraining Scheduler")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentDescribeDataIngestionJob, "describe-data-ingestion-job", "", false, "Describe Data Ingestion Job")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentDescribeDataset, "describe-dataset", "", false, "Describe Dataset")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentDescribeInferenceScheduler, "describe-inference-scheduler", "", false, "Describe Inference Scheduler")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentDescribeLabel, "describe-label", "", false, "Describe Label")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentDescribeLabelGroup, "describe-label-group", "", false, "Describe Label Group")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentDescribeModel, "describe-model", "", false, "Describe Model")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentDescribeModelVersion, "describe-model-version", "", false, "Describe Model Version")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentDescribeResourcePolicy, "describe-resource-policy", "", false, "Describe Resource Policy")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentDescribeRetrainingScheduler, "describe-retraining-scheduler", "", false, "Describe Retraining Scheduler")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentImportDataset, "import-dataset", "", false, "Import Dataset")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentImportModelVersion, "import-model-version", "", false, "Import Model Version")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentListDataIngestionJobs, "list-data-ingestion-jobs", "", false, "List Data Ingestion Jobs")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentListDatasets, "list-datasets", "", false, "List Datasets")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentListInferenceEvents, "list-inference-events", "", false, "List Inference Events")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentListInferenceExecutions, "list-inference-executions", "", false, "List Inference Executions")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentListInferenceSchedulers, "list-inference-schedulers", "", false, "List Inference Schedulers")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentListLabelGroups, "list-label-groups", "", false, "List Label Groups")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentListLabels, "list-labels", "", false, "List Labels")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentListModelVersions, "list-model-versions", "", false, "List Model Versions")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentListModels, "list-models", "", false, "List Models")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentListRetrainingSchedulers, "list-retraining-schedulers", "", false, "List Retraining Schedulers")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentListSensorStatistics, "list-sensor-statistics", "", false, "List Sensor Statistics")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentPutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentStartDataIngestionJob, "start-data-ingestion-job", "", false, "Start Data Ingestion Job")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentStartInferenceScheduler, "start-inference-scheduler", "", false, "Start Inference Scheduler")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentStartRetrainingScheduler, "start-retraining-scheduler", "", false, "Start Retraining Scheduler")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentStopInferenceScheduler, "stop-inference-scheduler", "", false, "Stop Inference Scheduler")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentStopRetrainingScheduler, "stop-retraining-scheduler", "", false, "Stop Retraining Scheduler")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentTagResource, "tag-resource", "", false, "Tag Resource")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentUntagResource, "untag-resource", "", false, "Untag Resource")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentUpdateActiveModelVersion, "update-active-model-version", "", false, "Update Active Model Version")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentUpdateInferenceScheduler, "update-inference-scheduler", "", false, "Update Inference Scheduler")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentUpdateLabelGroup, "update-label-group", "", false, "Update Label Group")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentUpdateModel, "update-model", "", false, "Update Model")
	_lookoutequipmentCmd.Flags().BoolVarP(&_lookoutequipmentUpdateRetrainingScheduler, "update-retraining-scheduler", "", false, "Update Retraining Scheduler")

}
