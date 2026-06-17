package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// frauddetectorCmd represents the frauddetector command
var _frauddetectorCmd = &cobra.Command{
	Use:   "frauddetector",
	Short: "AWS frauddetector CLI",
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
		client := frauddetector.NewFromConfig(cfg)
		if _frauddetectorBatchCreateVariable {
			frauddetector_BatchCreateVariable(cfg, client)
			return
		}
		if _frauddetectorBatchGetVariable {
			frauddetector_BatchGetVariable(cfg, client)
			return
		}
		if _frauddetectorCancelBatchImportJob {
			frauddetector_CancelBatchImportJob(cfg, client)
			return
		}
		if _frauddetectorCancelBatchPredictionJob {
			frauddetector_CancelBatchPredictionJob(cfg, client)
			return
		}
		if _frauddetectorCreateBatchImportJob {
			frauddetector_CreateBatchImportJob(cfg, client)
			return
		}
		if _frauddetectorCreateBatchPredictionJob {
			frauddetector_CreateBatchPredictionJob(cfg, client)
			return
		}
		if _frauddetectorCreateDetectorVersion {
			frauddetector_CreateDetectorVersion(cfg, client)
			return
		}
		if _frauddetectorCreateList {
			frauddetector_CreateList(cfg, client)
			return
		}
		if _frauddetectorCreateModel {
			frauddetector_CreateModel(cfg, client)
			return
		}
		if _frauddetectorCreateModelVersion {
			frauddetector_CreateModelVersion(cfg, client)
			return
		}
		if _frauddetectorCreateRule {
			frauddetector_CreateRule(cfg, client)
			return
		}
		if _frauddetectorCreateVariable {
			frauddetector_CreateVariable(cfg, client)
			return
		}
		if _frauddetectorDeleteBatchImportJob {
			frauddetector_DeleteBatchImportJob(cfg, client)
			return
		}
		if _frauddetectorDeleteBatchPredictionJob {
			frauddetector_DeleteBatchPredictionJob(cfg, client)
			return
		}
		if _frauddetectorDeleteDetector {
			frauddetector_DeleteDetector(cfg, client)
			return
		}
		if _frauddetectorDeleteDetectorVersion {
			frauddetector_DeleteDetectorVersion(cfg, client)
			return
		}
		if _frauddetectorDeleteEntityType {
			frauddetector_DeleteEntityType(cfg, client)
			return
		}
		if _frauddetectorDeleteEvent {
			frauddetector_DeleteEvent(cfg, client)
			return
		}
		if _frauddetectorDeleteEventType {
			frauddetector_DeleteEventType(cfg, client)
			return
		}
		if _frauddetectorDeleteEventsByEventType {
			frauddetector_DeleteEventsByEventType(cfg, client)
			return
		}
		if _frauddetectorDeleteExternalModel {
			frauddetector_DeleteExternalModel(cfg, client)
			return
		}
		if _frauddetectorDeleteLabel {
			frauddetector_DeleteLabel(cfg, client)
			return
		}
		if _frauddetectorDeleteList {
			frauddetector_DeleteList(cfg, client)
			return
		}
		if _frauddetectorDeleteModel {
			frauddetector_DeleteModel(cfg, client)
			return
		}
		if _frauddetectorDeleteModelVersion {
			frauddetector_DeleteModelVersion(cfg, client)
			return
		}
		if _frauddetectorDeleteOutcome {
			frauddetector_DeleteOutcome(cfg, client)
			return
		}
		if _frauddetectorDeleteRule {
			frauddetector_DeleteRule(cfg, client)
			return
		}
		if _frauddetectorDeleteVariable {
			frauddetector_DeleteVariable(cfg, client)
			return
		}
		if _frauddetectorDescribeDetector {
			frauddetector_DescribeDetector(cfg, client)
			return
		}
		if _frauddetectorDescribeModelVersions {
			frauddetector_DescribeModelVersions(cfg, client)
			return
		}
		if _frauddetectorGetBatchImportJobs {
			frauddetector_GetBatchImportJobs(cfg, client)
			return
		}
		if _frauddetectorGetBatchPredictionJobs {
			frauddetector_GetBatchPredictionJobs(cfg, client)
			return
		}
		if _frauddetectorGetDeleteEventsByEventTypeStatus {
			frauddetector_GetDeleteEventsByEventTypeStatus(cfg, client)
			return
		}
		if _frauddetectorGetDetectorVersion {
			frauddetector_GetDetectorVersion(cfg, client)
			return
		}
		if _frauddetectorGetDetectors {
			frauddetector_GetDetectors(cfg, client)
			return
		}
		if _frauddetectorGetEntityTypes {
			frauddetector_GetEntityTypes(cfg, client)
			return
		}
		if _frauddetectorGetEvent {
			frauddetector_GetEvent(cfg, client)
			return
		}
		if _frauddetectorGetEventPrediction {
			frauddetector_GetEventPrediction(cfg, client)
			return
		}
		if _frauddetectorGetEventPredictionMetadata {
			frauddetector_GetEventPredictionMetadata(cfg, client)
			return
		}
		if _frauddetectorGetEventTypes {
			frauddetector_GetEventTypes(cfg, client)
			return
		}
		if _frauddetectorGetExternalModels {
			frauddetector_GetExternalModels(cfg, client)
			return
		}
		if _frauddetectorGetKMSEncryptionKey {
			frauddetector_GetKMSEncryptionKey(cfg, client)
			return
		}
		if _frauddetectorGetLabels {
			frauddetector_GetLabels(cfg, client)
			return
		}
		if _frauddetectorGetListElements {
			frauddetector_GetListElements(cfg, client)
			return
		}
		if _frauddetectorGetListsMetadata {
			frauddetector_GetListsMetadata(cfg, client)
			return
		}
		if _frauddetectorGetModelVersion {
			frauddetector_GetModelVersion(cfg, client)
			return
		}
		if _frauddetectorGetModels {
			frauddetector_GetModels(cfg, client)
			return
		}
		if _frauddetectorGetOutcomes {
			frauddetector_GetOutcomes(cfg, client)
			return
		}
		if _frauddetectorGetRules {
			frauddetector_GetRules(cfg, client)
			return
		}
		if _frauddetectorGetVariables {
			frauddetector_GetVariables(cfg, client)
			return
		}
		if _frauddetectorListEventPredictions {
			frauddetector_ListEventPredictions(cfg, client)
			return
		}
		if _frauddetectorListTagsForResource {
			frauddetector_ListTagsForResource(cfg, client)
			return
		}
		if _frauddetectorPutDetector {
			frauddetector_PutDetector(cfg, client)
			return
		}
		if _frauddetectorPutEntityType {
			frauddetector_PutEntityType(cfg, client)
			return
		}
		if _frauddetectorPutEventType {
			frauddetector_PutEventType(cfg, client)
			return
		}
		if _frauddetectorPutExternalModel {
			frauddetector_PutExternalModel(cfg, client)
			return
		}
		if _frauddetectorPutKMSEncryptionKey {
			frauddetector_PutKMSEncryptionKey(cfg, client)
			return
		}
		if _frauddetectorPutLabel {
			frauddetector_PutLabel(cfg, client)
			return
		}
		if _frauddetectorPutOutcome {
			frauddetector_PutOutcome(cfg, client)
			return
		}
		if _frauddetectorSendEvent {
			frauddetector_SendEvent(cfg, client)
			return
		}
		if _frauddetectorTagResource {
			frauddetector_TagResource(cfg, client)
			return
		}
		if _frauddetectorUntagResource {
			frauddetector_UntagResource(cfg, client)
			return
		}
		if _frauddetectorUpdateDetectorVersion {
			frauddetector_UpdateDetectorVersion(cfg, client)
			return
		}
		if _frauddetectorUpdateDetectorVersionMetadata {
			frauddetector_UpdateDetectorVersionMetadata(cfg, client)
			return
		}
		if _frauddetectorUpdateDetectorVersionStatus {
			frauddetector_UpdateDetectorVersionStatus(cfg, client)
			return
		}
		if _frauddetectorUpdateEventLabel {
			frauddetector_UpdateEventLabel(cfg, client)
			return
		}
		if _frauddetectorUpdateList {
			frauddetector_UpdateList(cfg, client)
			return
		}
		if _frauddetectorUpdateModel {
			frauddetector_UpdateModel(cfg, client)
			return
		}
		if _frauddetectorUpdateModelVersion {
			frauddetector_UpdateModelVersion(cfg, client)
			return
		}
		if _frauddetectorUpdateModelVersionStatus {
			frauddetector_UpdateModelVersionStatus(cfg, client)
			return
		}
		if _frauddetectorUpdateRuleMetadata {
			frauddetector_UpdateRuleMetadata(cfg, client)
			return
		}
		if _frauddetectorUpdateRuleVersion {
			frauddetector_UpdateRuleVersion(cfg, client)
			return
		}
		if _frauddetectorUpdateVariable {
			frauddetector_UpdateVariable(cfg, client)
			return
		}

	},
}

var (
	_frauddetectorBatchCreateVariable              bool
	_frauddetectorBatchGetVariable                 bool
	_frauddetectorCancelBatchImportJob             bool
	_frauddetectorCancelBatchPredictionJob         bool
	_frauddetectorCreateBatchImportJob             bool
	_frauddetectorCreateBatchPredictionJob         bool
	_frauddetectorCreateDetectorVersion            bool
	_frauddetectorCreateList                       bool
	_frauddetectorCreateModel                      bool
	_frauddetectorCreateModelVersion               bool
	_frauddetectorCreateRule                       bool
	_frauddetectorCreateVariable                   bool
	_frauddetectorDeleteBatchImportJob             bool
	_frauddetectorDeleteBatchPredictionJob         bool
	_frauddetectorDeleteDetector                   bool
	_frauddetectorDeleteDetectorVersion            bool
	_frauddetectorDeleteEntityType                 bool
	_frauddetectorDeleteEvent                      bool
	_frauddetectorDeleteEventType                  bool
	_frauddetectorDeleteEventsByEventType          bool
	_frauddetectorDeleteExternalModel              bool
	_frauddetectorDeleteLabel                      bool
	_frauddetectorDeleteList                       bool
	_frauddetectorDeleteModel                      bool
	_frauddetectorDeleteModelVersion               bool
	_frauddetectorDeleteOutcome                    bool
	_frauddetectorDeleteRule                       bool
	_frauddetectorDeleteVariable                   bool
	_frauddetectorDescribeDetector                 bool
	_frauddetectorDescribeModelVersions            bool
	_frauddetectorGetBatchImportJobs               bool
	_frauddetectorGetBatchPredictionJobs           bool
	_frauddetectorGetDeleteEventsByEventTypeStatus bool
	_frauddetectorGetDetectorVersion               bool
	_frauddetectorGetDetectors                     bool
	_frauddetectorGetEntityTypes                   bool
	_frauddetectorGetEvent                         bool
	_frauddetectorGetEventPrediction               bool
	_frauddetectorGetEventPredictionMetadata       bool
	_frauddetectorGetEventTypes                    bool
	_frauddetectorGetExternalModels                bool
	_frauddetectorGetKMSEncryptionKey              bool
	_frauddetectorGetLabels                        bool
	_frauddetectorGetListElements                  bool
	_frauddetectorGetListsMetadata                 bool
	_frauddetectorGetModelVersion                  bool
	_frauddetectorGetModels                        bool
	_frauddetectorGetOutcomes                      bool
	_frauddetectorGetRules                         bool
	_frauddetectorGetVariables                     bool
	_frauddetectorListEventPredictions             bool
	_frauddetectorListTagsForResource              bool
	_frauddetectorPutDetector                      bool
	_frauddetectorPutEntityType                    bool
	_frauddetectorPutEventType                     bool
	_frauddetectorPutExternalModel                 bool
	_frauddetectorPutKMSEncryptionKey              bool
	_frauddetectorPutLabel                         bool
	_frauddetectorPutOutcome                       bool
	_frauddetectorSendEvent                        bool
	_frauddetectorTagResource                      bool
	_frauddetectorUntagResource                    bool
	_frauddetectorUpdateDetectorVersion            bool
	_frauddetectorUpdateDetectorVersionMetadata    bool
	_frauddetectorUpdateDetectorVersionStatus      bool
	_frauddetectorUpdateEventLabel                 bool
	_frauddetectorUpdateList                       bool
	_frauddetectorUpdateModel                      bool
	_frauddetectorUpdateModelVersion               bool
	_frauddetectorUpdateModelVersionStatus         bool
	_frauddetectorUpdateRuleMetadata               bool
	_frauddetectorUpdateRuleVersion                bool
	_frauddetectorUpdateVariable                   bool

	_frauddetectorAssignedLabel                  string
	_frauddetectorDataSource                     string
	_frauddetectorDataType                       string
	_frauddetectorDefaultValue                   string
	_frauddetectorDeleteAuditHistory             string
	_frauddetectorDescription                    string
	_frauddetectorDetectorId                     string
	_frauddetectorDetectorName                   string
	_frauddetectorDetectorVersion                string
	_frauddetectorDetectorVersionId              string
	_frauddetectorElements                       []string
	_frauddetectorEntities                       string
	_frauddetectorEntityTypes                    []string
	_frauddetectorEventId                        string
	_frauddetectorEventIngestion                 string
	_frauddetectorEventOrchestration             string
	_frauddetectorEventTimestamp                 string
	_frauddetectorEventType                      string
	_frauddetectorEventTypeName                  string
	_frauddetectorEventVariables                 string
	_frauddetectorExpression                     string
	_frauddetectorExternalEventsDetail           string
	_frauddetectorExternalModelEndpointDataBlobs string
	_frauddetectorExternalModelEndpoints         []string
	_frauddetectorIamRoleArn                     string
	_frauddetectorIngestedEventsDetail           string
	_frauddetectorInputConfiguration             string
	_frauddetectorInputPath                      string
	_frauddetectorInvokeModelEndpointRoleArn     string
	_frauddetectorJobId                          string
	_frauddetectorKmsEncryptionKeyArn            string
	_frauddetectorLabelTimestamp                 string
	_frauddetectorLabels                         []string
	_frauddetectorLanguage                       string
	_frauddetectorMajorVersionNumber             string
	_frauddetectorMaxResults                     string
	_frauddetectorModelEndpoint                  string
	_frauddetectorModelEndpointStatus            string
	_frauddetectorModelId                        string
	_frauddetectorModelSource                    string
	_frauddetectorModelType                      string
	_frauddetectorModelVersionNumber             string
	_frauddetectorModelVersions                  string
	_frauddetectorName                           string
	_frauddetectorNames                          []string
	_frauddetectorNextToken                      string
	_frauddetectorOutcomes                       []string
	_frauddetectorOutputConfiguration            string
	_frauddetectorOutputPath                     string
	_frauddetectorPredictionTimeRange            string
	_frauddetectorPredictionTimestamp            string
	_frauddetectorResourceARN                    string
	_frauddetectorRule                           string
	_frauddetectorRuleExecutionMode              string
	_frauddetectorRuleId                         string
	_frauddetectorRuleVersion                    string
	_frauddetectorRules                          string
	_frauddetectorStatus                         string
	_frauddetectorTagKeys                        []string
	_frauddetectorTags                           string
	_frauddetectorTrainingDataSchema             string
	_frauddetectorTrainingDataSource             string
	_frauddetectorUpdateMode                     string
	_frauddetectorVariableEntries                string
	_frauddetectorVariableType                   string
)

// Creates a batch of variables.
func frauddetector_BatchCreateVariable(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.BatchCreateVariableInput{
		// VariableEntries: []types.VariableEntry, // Required
	}

	if len(_frauddetectorVariableEntries) > 0 {
		if err := assignInputField(input, "VariableEntries", _frauddetectorVariableEntries); err != nil {
			log.Errorf("invalid --variable-entries: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorTags) > 0 {
		if err := assignInputField(input, "Tags", _frauddetectorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchCreateVariable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a batch of variables.
func frauddetector_BatchGetVariable(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.BatchGetVariableInput{
		// Names: []string, // Required
	}

	if len(_frauddetectorNames) > 0 {
		input.Names = append([]string(nil), _frauddetectorNames...)
	}

	if resp, err := client.BatchGetVariable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels an in-progress batch import job.
func frauddetector_CancelBatchImportJob(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.CancelBatchImportJobInput{
		// JobId: *string, // Required
	}

	if len(_frauddetectorJobId) > 0 {
		input.JobId = aws.String(_frauddetectorJobId)
	}

	if resp, err := client.CancelBatchImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels the specified batch prediction job.
func frauddetector_CancelBatchPredictionJob(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.CancelBatchPredictionJobInput{
		// JobId: *string, // Required
	}

	if len(_frauddetectorJobId) > 0 {
		input.JobId = aws.String(_frauddetectorJobId)
	}

	if resp, err := client.CancelBatchPredictionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a batch import job.
func frauddetector_CreateBatchImportJob(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.CreateBatchImportJobInput{
		// EventTypeName: *string, // Required
		// IamRoleArn: *string, // Required
		// InputPath: *string, // Required
		// JobId: *string, // Required
		// OutputPath: *string, // Required
	}

	if len(_frauddetectorEventTypeName) > 0 {
		input.EventTypeName = aws.String(_frauddetectorEventTypeName)
	}
	if len(_frauddetectorIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_frauddetectorIamRoleArn)
	}
	if len(_frauddetectorInputPath) > 0 {
		input.InputPath = aws.String(_frauddetectorInputPath)
	}
	if len(_frauddetectorJobId) > 0 {
		input.JobId = aws.String(_frauddetectorJobId)
	}
	if len(_frauddetectorOutputPath) > 0 {
		input.OutputPath = aws.String(_frauddetectorOutputPath)
	}
	if len(_frauddetectorTags) > 0 {
		if err := assignInputField(input, "Tags", _frauddetectorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBatchImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a batch prediction job.
func frauddetector_CreateBatchPredictionJob(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.CreateBatchPredictionJobInput{
		// DetectorName: *string, // Required
		// EventTypeName: *string, // Required
		// IamRoleArn: *string, // Required
		// InputPath: *string, // Required
		// JobId: *string, // Required
		// OutputPath: *string, // Required
	}

	if len(_frauddetectorDetectorName) > 0 {
		input.DetectorName = aws.String(_frauddetectorDetectorName)
	}
	if len(_frauddetectorEventTypeName) > 0 {
		input.EventTypeName = aws.String(_frauddetectorEventTypeName)
	}
	if len(_frauddetectorIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_frauddetectorIamRoleArn)
	}
	if len(_frauddetectorInputPath) > 0 {
		input.InputPath = aws.String(_frauddetectorInputPath)
	}
	if len(_frauddetectorJobId) > 0 {
		input.JobId = aws.String(_frauddetectorJobId)
	}
	if len(_frauddetectorOutputPath) > 0 {
		input.OutputPath = aws.String(_frauddetectorOutputPath)
	}
	if len(_frauddetectorDetectorVersion) > 0 {
		input.DetectorVersion = aws.String(_frauddetectorDetectorVersion)
	}
	if len(_frauddetectorTags) > 0 {
		if err := assignInputField(input, "Tags", _frauddetectorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBatchPredictionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a detector version. The detector version starts in a DRAFT status.
func frauddetector_CreateDetectorVersion(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.CreateDetectorVersionInput{
		// DetectorId: *string, // Required
		// Rules: []types.Rule, // Required
	}

	if len(_frauddetectorDetectorId) > 0 {
		input.DetectorId = aws.String(_frauddetectorDetectorId)
	}
	if len(_frauddetectorRules) > 0 {
		if err := assignInputField(input, "Rules", _frauddetectorRules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorDescription) > 0 {
		input.Description = aws.String(_frauddetectorDescription)
	}
	if len(_frauddetectorExternalModelEndpoints) > 0 {
		input.ExternalModelEndpoints = append([]string(nil), _frauddetectorExternalModelEndpoints...)
	}
	if len(_frauddetectorModelVersions) > 0 {
		if err := assignInputField(input, "ModelVersions", _frauddetectorModelVersions); err != nil {
			log.Errorf("invalid --model-versions: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorRuleExecutionMode) > 0 {
		if err := assignInputField(input, "RuleExecutionMode", _frauddetectorRuleExecutionMode); err != nil {
			log.Errorf("invalid --rule-execution-mode: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorTags) > 0 {
		if err := assignInputField(input, "Tags", _frauddetectorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDetectorVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a list.
// List is a set of input data for a variable in your event dataset. You use the
// input data in a rule that's associated with your detector. For more information,
// see [Lists].
//
// [Lists]: https://docs.aws.amazon.com/frauddetector/latest/ug/lists.html
func frauddetector_CreateList(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.CreateListInput{
		// Name: *string, // Required
	}

	if len(_frauddetectorName) > 0 {
		input.Name = aws.String(_frauddetectorName)
	}
	if len(_frauddetectorDescription) > 0 {
		input.Description = aws.String(_frauddetectorDescription)
	}
	if len(_frauddetectorElements) > 0 {
		input.Elements = append([]string(nil), _frauddetectorElements...)
	}
	if len(_frauddetectorTags) > 0 {
		if err := assignInputField(input, "Tags", _frauddetectorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorVariableType) > 0 {
		input.VariableType = aws.String(_frauddetectorVariableType)
	}

	if resp, err := client.CreateList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a model using the specified model type.
func frauddetector_CreateModel(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.CreateModelInput{
		// EventTypeName: *string, // Required
		// ModelId: *string, // Required
		// ModelType: types.ModelTypeEnum, // Required
	}

	if len(_frauddetectorEventTypeName) > 0 {
		input.EventTypeName = aws.String(_frauddetectorEventTypeName)
	}
	if len(_frauddetectorModelId) > 0 {
		input.ModelId = aws.String(_frauddetectorModelId)
	}
	if len(_frauddetectorModelType) > 0 {
		if err := assignInputField(input, "ModelType", _frauddetectorModelType); err != nil {
			log.Errorf("invalid --model-type: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorDescription) > 0 {
		input.Description = aws.String(_frauddetectorDescription)
	}
	if len(_frauddetectorTags) > 0 {
		if err := assignInputField(input, "Tags", _frauddetectorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
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

// Creates a version of the model using the specified model type and model id.
func frauddetector_CreateModelVersion(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.CreateModelVersionInput{
		// ModelId: *string, // Required
		// ModelType: types.ModelTypeEnum, // Required
		// TrainingDataSchema: *types.TrainingDataSchema, // Required
		// TrainingDataSource: types.TrainingDataSourceEnum, // Required
	}

	if len(_frauddetectorModelId) > 0 {
		input.ModelId = aws.String(_frauddetectorModelId)
	}
	if len(_frauddetectorModelType) > 0 {
		if err := assignInputField(input, "ModelType", _frauddetectorModelType); err != nil {
			log.Errorf("invalid --model-type: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorTrainingDataSchema) > 0 {
		if err := assignInputField(input, "TrainingDataSchema", _frauddetectorTrainingDataSchema); err != nil {
			log.Errorf("invalid --training-data-schema: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorTrainingDataSource) > 0 {
		if err := assignInputField(input, "TrainingDataSource", _frauddetectorTrainingDataSource); err != nil {
			log.Errorf("invalid --training-data-source: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorExternalEventsDetail) > 0 {
		if err := assignInputField(input, "ExternalEventsDetail", _frauddetectorExternalEventsDetail); err != nil {
			log.Errorf("invalid --external-events-detail: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorIngestedEventsDetail) > 0 {
		if err := assignInputField(input, "IngestedEventsDetail", _frauddetectorIngestedEventsDetail); err != nil {
			log.Errorf("invalid --ingested-events-detail: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorTags) > 0 {
		if err := assignInputField(input, "Tags", _frauddetectorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateModelVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a rule for use with the specified detector.
func frauddetector_CreateRule(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.CreateRuleInput{
		// DetectorId: *string, // Required
		// Expression: *string, // Required
		// Language: types.Language, // Required
		// Outcomes: []string, // Required
		// RuleId: *string, // Required
	}

	if len(_frauddetectorDetectorId) > 0 {
		input.DetectorId = aws.String(_frauddetectorDetectorId)
	}
	if len(_frauddetectorExpression) > 0 {
		input.Expression = aws.String(_frauddetectorExpression)
	}
	if len(_frauddetectorLanguage) > 0 {
		if err := assignInputField(input, "Language", _frauddetectorLanguage); err != nil {
			log.Errorf("invalid --language: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorOutcomes) > 0 {
		input.Outcomes = append([]string(nil), _frauddetectorOutcomes...)
	}
	if len(_frauddetectorRuleId) > 0 {
		input.RuleId = aws.String(_frauddetectorRuleId)
	}
	if len(_frauddetectorDescription) > 0 {
		input.Description = aws.String(_frauddetectorDescription)
	}
	if len(_frauddetectorTags) > 0 {
		if err := assignInputField(input, "Tags", _frauddetectorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a variable.
func frauddetector_CreateVariable(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.CreateVariableInput{
		// DataSource: types.DataSource, // Required
		// DataType: types.DataType, // Required
		// DefaultValue: *string, // Required
		// Name: *string, // Required
	}

	if len(_frauddetectorDataSource) > 0 {
		if err := assignInputField(input, "DataSource", _frauddetectorDataSource); err != nil {
			log.Errorf("invalid --data-source: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorDataType) > 0 {
		if err := assignInputField(input, "DataType", _frauddetectorDataType); err != nil {
			log.Errorf("invalid --data-type: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorDefaultValue) > 0 {
		input.DefaultValue = aws.String(_frauddetectorDefaultValue)
	}
	if len(_frauddetectorName) > 0 {
		input.Name = aws.String(_frauddetectorName)
	}
	if len(_frauddetectorDescription) > 0 {
		input.Description = aws.String(_frauddetectorDescription)
	}
	if len(_frauddetectorTags) > 0 {
		if err := assignInputField(input, "Tags", _frauddetectorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorVariableType) > 0 {
		input.VariableType = aws.String(_frauddetectorVariableType)
	}

	if resp, err := client.CreateVariable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified batch import job ID record. This action does not delete
// the data that was batch imported.
func frauddetector_DeleteBatchImportJob(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.DeleteBatchImportJobInput{
		// JobId: *string, // Required
	}

	if len(_frauddetectorJobId) > 0 {
		input.JobId = aws.String(_frauddetectorJobId)
	}

	if resp, err := client.DeleteBatchImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a batch prediction job.
func frauddetector_DeleteBatchPredictionJob(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.DeleteBatchPredictionJobInput{
		// JobId: *string, // Required
	}

	if len(_frauddetectorJobId) > 0 {
		input.JobId = aws.String(_frauddetectorJobId)
	}

	if resp, err := client.DeleteBatchPredictionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the detector. Before deleting a detector, you must first delete all
// detector versions and rule versions associated with the detector.
//
// When you delete a detector, Amazon Fraud Detector permanently deletes the
// detector and the data is no longer stored in Amazon Fraud Detector.
func frauddetector_DeleteDetector(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.DeleteDetectorInput{
		// DetectorId: *string, // Required
	}

	if len(_frauddetectorDetectorId) > 0 {
		input.DetectorId = aws.String(_frauddetectorDetectorId)
	}

	if resp, err := client.DeleteDetector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the detector version. You cannot delete detector versions that are in
// ACTIVE status.
//
// When you delete a detector version, Amazon Fraud Detector permanently deletes
// the detector and the data is no longer stored in Amazon Fraud Detector.
func frauddetector_DeleteDetectorVersion(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.DeleteDetectorVersionInput{
		// DetectorId: *string, // Required
		// DetectorVersionId: *string, // Required
	}

	if len(_frauddetectorDetectorId) > 0 {
		input.DetectorId = aws.String(_frauddetectorDetectorId)
	}
	if len(_frauddetectorDetectorVersionId) > 0 {
		input.DetectorVersionId = aws.String(_frauddetectorDetectorVersionId)
	}

	if resp, err := client.DeleteDetectorVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an entity type.
// You cannot delete an entity type that is included in an event type.
//
// When you delete an entity type, Amazon Fraud Detector permanently deletes that
// entity type and the data is no longer stored in Amazon Fraud Detector.
func frauddetector_DeleteEntityType(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.DeleteEntityTypeInput{
		// Name: *string, // Required
	}

	if len(_frauddetectorName) > 0 {
		input.Name = aws.String(_frauddetectorName)
	}

	if resp, err := client.DeleteEntityType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified event.
// When you delete an event, Amazon Fraud Detector permanently deletes that event
// and the event data is no longer stored in Amazon Fraud Detector. If
// deleteAuditHistory is True , event data is available through search for up to 30
// seconds after the delete operation is completed.
func frauddetector_DeleteEvent(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.DeleteEventInput{
		// EventId: *string, // Required
		// EventTypeName: *string, // Required
	}

	if len(_frauddetectorEventId) > 0 {
		input.EventId = aws.String(_frauddetectorEventId)
	}
	if len(_frauddetectorEventTypeName) > 0 {
		input.EventTypeName = aws.String(_frauddetectorEventTypeName)
	}
	if len(_frauddetectorDeleteAuditHistory) > 0 {
		if err := assignInputField(input, "DeleteAuditHistory", _frauddetectorDeleteAuditHistory); err != nil {
			log.Errorf("invalid --delete-audit-history: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an event type.
// You cannot delete an event type that is used in a detector or a model.
//
// When you delete an event type, Amazon Fraud Detector permanently deletes that
// event type and the data is no longer stored in Amazon Fraud Detector.
func frauddetector_DeleteEventType(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.DeleteEventTypeInput{
		// Name: *string, // Required
	}

	if len(_frauddetectorName) > 0 {
		input.Name = aws.String(_frauddetectorName)
	}

	if resp, err := client.DeleteEventType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes all events of a particular event type.
func frauddetector_DeleteEventsByEventType(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.DeleteEventsByEventTypeInput{
		// EventTypeName: *string, // Required
	}

	if len(_frauddetectorEventTypeName) > 0 {
		input.EventTypeName = aws.String(_frauddetectorEventTypeName)
	}

	if resp, err := client.DeleteEventsByEventType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a SageMaker model from Amazon Fraud Detector.
// You can remove an Amazon SageMaker model if it is not associated with a
// detector version. Removing a SageMaker model disconnects it from Amazon Fraud
// Detector, but the model remains available in SageMaker.
func frauddetector_DeleteExternalModel(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.DeleteExternalModelInput{
		// ModelEndpoint: *string, // Required
	}

	if len(_frauddetectorModelEndpoint) > 0 {
		input.ModelEndpoint = aws.String(_frauddetectorModelEndpoint)
	}

	if resp, err := client.DeleteExternalModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a label.
// You cannot delete labels that are included in an event type in Amazon Fraud
// Detector.
//
// You cannot delete a label assigned to an event ID. You must first delete the
// relevant event ID.
//
// When you delete a label, Amazon Fraud Detector permanently deletes that label
// and the data is no longer stored in Amazon Fraud Detector.
func frauddetector_DeleteLabel(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.DeleteLabelInput{
		// Name: *string, // Required
	}

	if len(_frauddetectorName) > 0 {
		input.Name = aws.String(_frauddetectorName)
	}

	if resp, err := client.DeleteLabel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the list, provided it is not used in a rule.
// When you delete a list, Amazon Fraud Detector permanently deletes that list and
// the elements in the list.
func frauddetector_DeleteList(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.DeleteListInput{
		// Name: *string, // Required
	}

	if len(_frauddetectorName) > 0 {
		input.Name = aws.String(_frauddetectorName)
	}

	if resp, err := client.DeleteList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a model.
// You can delete models and model versions in Amazon Fraud Detector, provided
// that they are not associated with a detector version.
//
// When you delete a model, Amazon Fraud Detector permanently deletes that model
// and the data is no longer stored in Amazon Fraud Detector.
func frauddetector_DeleteModel(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.DeleteModelInput{
		// ModelId: *string, // Required
		// ModelType: types.ModelTypeEnum, // Required
	}

	if len(_frauddetectorModelId) > 0 {
		input.ModelId = aws.String(_frauddetectorModelId)
	}
	if len(_frauddetectorModelType) > 0 {
		if err := assignInputField(input, "ModelType", _frauddetectorModelType); err != nil {
			log.Errorf("invalid --model-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a model version.
// You can delete models and model versions in Amazon Fraud Detector, provided
// that they are not associated with a detector version.
//
// When you delete a model version, Amazon Fraud Detector permanently deletes that
// model version and the data is no longer stored in Amazon Fraud Detector.
func frauddetector_DeleteModelVersion(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.DeleteModelVersionInput{
		// ModelId: *string, // Required
		// ModelType: types.ModelTypeEnum, // Required
		// ModelVersionNumber: *string, // Required
	}

	if len(_frauddetectorModelId) > 0 {
		input.ModelId = aws.String(_frauddetectorModelId)
	}
	if len(_frauddetectorModelType) > 0 {
		if err := assignInputField(input, "ModelType", _frauddetectorModelType); err != nil {
			log.Errorf("invalid --model-type: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorModelVersionNumber) > 0 {
		input.ModelVersionNumber = aws.String(_frauddetectorModelVersionNumber)
	}

	if resp, err := client.DeleteModelVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an outcome.
// You cannot delete an outcome that is used in a rule version.
//
// When you delete an outcome, Amazon Fraud Detector permanently deletes that
// outcome and the data is no longer stored in Amazon Fraud Detector.
func frauddetector_DeleteOutcome(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.DeleteOutcomeInput{
		// Name: *string, // Required
	}

	if len(_frauddetectorName) > 0 {
		input.Name = aws.String(_frauddetectorName)
	}

	if resp, err := client.DeleteOutcome(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the rule. You cannot delete a rule if it is used by an ACTIVE or
// INACTIVE detector version.
//
// When you delete a rule, Amazon Fraud Detector permanently deletes that rule and
// the data is no longer stored in Amazon Fraud Detector.
func frauddetector_DeleteRule(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.DeleteRuleInput{
		// Rule: *types.Rule, // Required
	}

	if len(_frauddetectorRule) > 0 {
		if err := assignInputField(input, "Rule", _frauddetectorRule); err != nil {
			log.Errorf("invalid --rule: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a variable.
// You can't delete variables that are included in an event type in Amazon Fraud
// Detector.
//
// Amazon Fraud Detector automatically deletes model output variables and
// SageMaker model output variables when you delete the model. You can't delete
// these variables manually.
//
// When you delete a variable, Amazon Fraud Detector permanently deletes that
// variable and the data is no longer stored in Amazon Fraud Detector.
func frauddetector_DeleteVariable(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.DeleteVariableInput{
		// Name: *string, // Required
	}

	if len(_frauddetectorName) > 0 {
		input.Name = aws.String(_frauddetectorName)
	}

	if resp, err := client.DeleteVariable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets all versions for a specified detector.
func frauddetector_DescribeDetector(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.DescribeDetectorInput{
		// DetectorId: *string, // Required
	}

	if len(_frauddetectorDetectorId) > 0 {
		input.DetectorId = aws.String(_frauddetectorDetectorId)
	}
	if len(_frauddetectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _frauddetectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorNextToken) > 0 {
		input.NextToken = aws.String(_frauddetectorNextToken)
	}

	if resp, err := client.DescribeDetector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets all of the model versions for the specified model type or for the
// specified model type and model ID. You can also get details for a single,
// specified model version.
func frauddetector_DescribeModelVersions(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.DescribeModelVersionsInput{}

	if len(_frauddetectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _frauddetectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorModelId) > 0 {
		input.ModelId = aws.String(_frauddetectorModelId)
	}
	if len(_frauddetectorModelType) > 0 {
		if err := assignInputField(input, "ModelType", _frauddetectorModelType); err != nil {
			log.Errorf("invalid --model-type: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorModelVersionNumber) > 0 {
		input.ModelVersionNumber = aws.String(_frauddetectorModelVersionNumber)
	}
	if len(_frauddetectorNextToken) > 0 {
		input.NextToken = aws.String(_frauddetectorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeModelVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*frauddetector.DescribeModelVersionsOutput
	p := frauddetector.NewDescribeModelVersionsPaginator(client, input)
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

// Gets all batch import jobs or a specific job of the specified ID. This is a
// paginated API. If you provide a null maxResults , this action retrieves a
// maximum of 50 records per page. If you provide a maxResults , the value must be
// between 1 and 50. To get the next page results, provide the pagination token
// from the GetBatchImportJobsResponse as part of your request. A null pagination
// token fetches the records from the beginning.
func frauddetector_GetBatchImportJobs(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.GetBatchImportJobsInput{}

	if len(_frauddetectorJobId) > 0 {
		input.JobId = aws.String(_frauddetectorJobId)
	}
	if len(_frauddetectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _frauddetectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorNextToken) > 0 {
		input.NextToken = aws.String(_frauddetectorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetBatchImportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*frauddetector.GetBatchImportJobsOutput
	p := frauddetector.NewGetBatchImportJobsPaginator(client, input)
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

// Gets all batch prediction jobs or a specific job if you specify a job ID. This
// is a paginated API. If you provide a null maxResults, this action retrieves a
// maximum of 50 records per page. If you provide a maxResults, the value must be
// between 1 and 50. To get the next page results, provide the pagination token
// from the GetBatchPredictionJobsResponse as part of your request. A null
// pagination token fetches the records from the beginning.
func frauddetector_GetBatchPredictionJobs(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.GetBatchPredictionJobsInput{}

	if len(_frauddetectorJobId) > 0 {
		input.JobId = aws.String(_frauddetectorJobId)
	}
	if len(_frauddetectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _frauddetectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorNextToken) > 0 {
		input.NextToken = aws.String(_frauddetectorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetBatchPredictionJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*frauddetector.GetBatchPredictionJobsOutput
	p := frauddetector.NewGetBatchPredictionJobsPaginator(client, input)
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

// Retrieves the status of a DeleteEventsByEventType action.
func frauddetector_GetDeleteEventsByEventTypeStatus(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.GetDeleteEventsByEventTypeStatusInput{
		// EventTypeName: *string, // Required
	}

	if len(_frauddetectorEventTypeName) > 0 {
		input.EventTypeName = aws.String(_frauddetectorEventTypeName)
	}

	if resp, err := client.GetDeleteEventsByEventTypeStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a particular detector version.
func frauddetector_GetDetectorVersion(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.GetDetectorVersionInput{
		// DetectorId: *string, // Required
		// DetectorVersionId: *string, // Required
	}

	if len(_frauddetectorDetectorId) > 0 {
		input.DetectorId = aws.String(_frauddetectorDetectorId)
	}
	if len(_frauddetectorDetectorVersionId) > 0 {
		input.DetectorVersionId = aws.String(_frauddetectorDetectorVersionId)
	}

	if resp, err := client.GetDetectorVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets all detectors or a single detector if a detectorId is specified. This is a
// paginated API. If you provide a null maxResults , this action retrieves a
// maximum of 10 records per page. If you provide a maxResults , the value must be
// between 5 and 10. To get the next page results, provide the pagination token
// from the GetDetectorsResponse as part of your request. A null pagination token
// fetches the records from the beginning.
func frauddetector_GetDetectors(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.GetDetectorsInput{}

	if len(_frauddetectorDetectorId) > 0 {
		input.DetectorId = aws.String(_frauddetectorDetectorId)
	}
	if len(_frauddetectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _frauddetectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorNextToken) > 0 {
		input.NextToken = aws.String(_frauddetectorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetDetectors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*frauddetector.GetDetectorsOutput
	p := frauddetector.NewGetDetectorsPaginator(client, input)
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

// Gets all entity types or a specific entity type if a name is specified. This is
// a paginated API. If you provide a null maxResults , this action retrieves a
// maximum of 10 records per page. If you provide a maxResults , the value must be
// between 5 and 10. To get the next page results, provide the pagination token
// from the GetEntityTypesResponse as part of your request. A null pagination
// token fetches the records from the beginning.
func frauddetector_GetEntityTypes(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.GetEntityTypesInput{}

	if len(_frauddetectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _frauddetectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorName) > 0 {
		input.Name = aws.String(_frauddetectorName)
	}
	if len(_frauddetectorNextToken) > 0 {
		input.NextToken = aws.String(_frauddetectorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetEntityTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*frauddetector.GetEntityTypesOutput
	p := frauddetector.NewGetEntityTypesPaginator(client, input)
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

// Retrieves details of events stored with Amazon Fraud Detector. This action does
// not retrieve prediction results.
func frauddetector_GetEvent(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.GetEventInput{
		// EventId: *string, // Required
		// EventTypeName: *string, // Required
	}

	if len(_frauddetectorEventId) > 0 {
		input.EventId = aws.String(_frauddetectorEventId)
	}
	if len(_frauddetectorEventTypeName) > 0 {
		input.EventTypeName = aws.String(_frauddetectorEventTypeName)
	}

	if resp, err := client.GetEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Evaluates an event against a detector version. If a version ID is not provided,
// the detector’s ( ACTIVE ) version is used.
func frauddetector_GetEventPrediction(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.GetEventPredictionInput{
		// DetectorId: *string, // Required
		// Entities: []types.Entity, // Required
		// EventId: *string, // Required
		// EventTimestamp: *string, // Required
		// EventTypeName: *string, // Required
		// EventVariables: map[string]string, // Required
	}

	if len(_frauddetectorDetectorId) > 0 {
		input.DetectorId = aws.String(_frauddetectorDetectorId)
	}
	if len(_frauddetectorEntities) > 0 {
		if err := assignInputField(input, "Entities", _frauddetectorEntities); err != nil {
			log.Errorf("invalid --entities: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorEventId) > 0 {
		input.EventId = aws.String(_frauddetectorEventId)
	}
	if len(_frauddetectorEventTimestamp) > 0 {
		input.EventTimestamp = aws.String(_frauddetectorEventTimestamp)
	}
	if len(_frauddetectorEventTypeName) > 0 {
		input.EventTypeName = aws.String(_frauddetectorEventTypeName)
	}
	if len(_frauddetectorEventVariables) > 0 {
		if err := assignInputField(input, "EventVariables", _frauddetectorEventVariables); err != nil {
			log.Errorf("invalid --event-variables: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorDetectorVersionId) > 0 {
		input.DetectorVersionId = aws.String(_frauddetectorDetectorVersionId)
	}
	if len(_frauddetectorExternalModelEndpointDataBlobs) > 0 {
		if err := assignInputField(input, "ExternalModelEndpointDataBlobs", _frauddetectorExternalModelEndpointDataBlobs); err != nil {
			log.Errorf("invalid --external-model-endpoint-data-blobs: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetEventPrediction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details of the past fraud predictions for the specified event ID, event
// type, detector ID, and detector version ID that was generated in the specified
// time period.
func frauddetector_GetEventPredictionMetadata(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.GetEventPredictionMetadataInput{
		// DetectorId: *string, // Required
		// DetectorVersionId: *string, // Required
		// EventId: *string, // Required
		// EventTypeName: *string, // Required
		// PredictionTimestamp: *string, // Required
	}

	if len(_frauddetectorDetectorId) > 0 {
		input.DetectorId = aws.String(_frauddetectorDetectorId)
	}
	if len(_frauddetectorDetectorVersionId) > 0 {
		input.DetectorVersionId = aws.String(_frauddetectorDetectorVersionId)
	}
	if len(_frauddetectorEventId) > 0 {
		input.EventId = aws.String(_frauddetectorEventId)
	}
	if len(_frauddetectorEventTypeName) > 0 {
		input.EventTypeName = aws.String(_frauddetectorEventTypeName)
	}
	if len(_frauddetectorPredictionTimestamp) > 0 {
		input.PredictionTimestamp = aws.String(_frauddetectorPredictionTimestamp)
	}

	if resp, err := client.GetEventPredictionMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets all event types or a specific event type if name is provided. This is a
// paginated API. If you provide a null maxResults , this action retrieves a
// maximum of 10 records per page. If you provide a maxResults , the value must be
// between 5 and 10. To get the next page results, provide the pagination token
// from the GetEventTypesResponse as part of your request. A null pagination token
// fetches the records from the beginning.
func frauddetector_GetEventTypes(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.GetEventTypesInput{}

	if len(_frauddetectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _frauddetectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorName) > 0 {
		input.Name = aws.String(_frauddetectorName)
	}
	if len(_frauddetectorNextToken) > 0 {
		input.NextToken = aws.String(_frauddetectorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetEventTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*frauddetector.GetEventTypesOutput
	p := frauddetector.NewGetEventTypesPaginator(client, input)
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

// Gets the details for one or more Amazon SageMaker models that have been
// imported into the service. This is a paginated API. If you provide a null
// maxResults , this actions retrieves a maximum of 10 records per page. If you
// provide a maxResults , the value must be between 5 and 10. To get the next page
// results, provide the pagination token from the GetExternalModelsResult as part
// of your request. A null pagination token fetches the records from the beginning.
func frauddetector_GetExternalModels(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.GetExternalModelsInput{}

	if len(_frauddetectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _frauddetectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorModelEndpoint) > 0 {
		input.ModelEndpoint = aws.String(_frauddetectorModelEndpoint)
	}
	if len(_frauddetectorNextToken) > 0 {
		input.NextToken = aws.String(_frauddetectorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetExternalModels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*frauddetector.GetExternalModelsOutput
	p := frauddetector.NewGetExternalModelsPaginator(client, input)
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

// Gets the encryption key if a KMS key has been specified to be used to encrypt
// content in Amazon Fraud Detector.
func frauddetector_GetKMSEncryptionKey(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.GetKMSEncryptionKeyInput{}

	if resp, err := client.GetKMSEncryptionKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets all labels or a specific label if name is provided. This is a paginated
// API. If you provide a null maxResults , this action retrieves a maximum of 50
// records per page. If you provide a maxResults , the value must be between 10 and
// 50. To get the next page results, provide the pagination token from the
// GetGetLabelsResponse as part of your request. A null pagination token fetches
// the records from the beginning.
func frauddetector_GetLabels(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.GetLabelsInput{}

	if len(_frauddetectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _frauddetectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorName) > 0 {
		input.Name = aws.String(_frauddetectorName)
	}
	if len(_frauddetectorNextToken) > 0 {
		input.NextToken = aws.String(_frauddetectorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetLabels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*frauddetector.GetLabelsOutput
	p := frauddetector.NewGetLabelsPaginator(client, input)
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

// Gets all the elements in the specified list.
func frauddetector_GetListElements(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.GetListElementsInput{
		// Name: *string, // Required
	}

	if len(_frauddetectorName) > 0 {
		input.Name = aws.String(_frauddetectorName)
	}
	if len(_frauddetectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _frauddetectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorNextToken) > 0 {
		input.NextToken = aws.String(_frauddetectorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetListElements(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*frauddetector.GetListElementsOutput
	p := frauddetector.NewGetListElementsPaginator(client, input)
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

// Gets the metadata of either all the lists under the account or the specified
// list.
func frauddetector_GetListsMetadata(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.GetListsMetadataInput{}

	if len(_frauddetectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _frauddetectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorName) > 0 {
		input.Name = aws.String(_frauddetectorName)
	}
	if len(_frauddetectorNextToken) > 0 {
		input.NextToken = aws.String(_frauddetectorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetListsMetadata(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*frauddetector.GetListsMetadataOutput
	p := frauddetector.NewGetListsMetadataPaginator(client, input)
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

// Gets the details of the specified model version.
func frauddetector_GetModelVersion(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.GetModelVersionInput{
		// ModelId: *string, // Required
		// ModelType: types.ModelTypeEnum, // Required
		// ModelVersionNumber: *string, // Required
	}

	if len(_frauddetectorModelId) > 0 {
		input.ModelId = aws.String(_frauddetectorModelId)
	}
	if len(_frauddetectorModelType) > 0 {
		if err := assignInputField(input, "ModelType", _frauddetectorModelType); err != nil {
			log.Errorf("invalid --model-type: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorModelVersionNumber) > 0 {
		input.ModelVersionNumber = aws.String(_frauddetectorModelVersionNumber)
	}

	if resp, err := client.GetModelVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets one or more models. Gets all models for the Amazon Web Services account if
// no model type and no model id provided. Gets all models for the Amazon Web
// Services account and model type, if the model type is specified but model id is
// not provided. Gets a specific model if (model type, model id) tuple is
// specified.
//
// This is a paginated API. If you provide a null maxResults , this action
// retrieves a maximum of 10 records per page. If you provide a maxResults , the
// value must be between 1 and 10. To get the next page results, provide the
// pagination token from the response as part of your request. A null pagination
// token fetches the records from the beginning.
func frauddetector_GetModels(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.GetModelsInput{}

	if len(_frauddetectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _frauddetectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorModelId) > 0 {
		input.ModelId = aws.String(_frauddetectorModelId)
	}
	if len(_frauddetectorModelType) > 0 {
		if err := assignInputField(input, "ModelType", _frauddetectorModelType); err != nil {
			log.Errorf("invalid --model-type: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorNextToken) > 0 {
		input.NextToken = aws.String(_frauddetectorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetModels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*frauddetector.GetModelsOutput
	p := frauddetector.NewGetModelsPaginator(client, input)
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

// Gets one or more outcomes. This is a paginated API. If you provide a null
// maxResults , this actions retrieves a maximum of 100 records per page. If you
// provide a maxResults , the value must be between 50 and 100. To get the next
// page results, provide the pagination token from the GetOutcomesResult as part
// of your request. A null pagination token fetches the records from the beginning.
func frauddetector_GetOutcomes(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.GetOutcomesInput{}

	if len(_frauddetectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _frauddetectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorName) > 0 {
		input.Name = aws.String(_frauddetectorName)
	}
	if len(_frauddetectorNextToken) > 0 {
		input.NextToken = aws.String(_frauddetectorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetOutcomes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*frauddetector.GetOutcomesOutput
	p := frauddetector.NewGetOutcomesPaginator(client, input)
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

// Get all rules for a detector (paginated) if ruleId and ruleVersion are not
// specified. Gets all rules for the detector and the ruleId if present
// (paginated). Gets a specific rule if both the ruleId and the ruleVersion are
// specified.
//
// This is a paginated API. Providing null maxResults results in retrieving
// maximum of 100 records per page. If you provide maxResults the value must be
// between 50 and 100. To get the next page result, a provide a pagination token
// from GetRulesResult as part of your request. Null pagination token fetches the
// records from the beginning.
func frauddetector_GetRules(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.GetRulesInput{
		// DetectorId: *string, // Required
	}

	if len(_frauddetectorDetectorId) > 0 {
		input.DetectorId = aws.String(_frauddetectorDetectorId)
	}
	if len(_frauddetectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _frauddetectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorNextToken) > 0 {
		input.NextToken = aws.String(_frauddetectorNextToken)
	}
	if len(_frauddetectorRuleId) > 0 {
		input.RuleId = aws.String(_frauddetectorRuleId)
	}
	if len(_frauddetectorRuleVersion) > 0 {
		input.RuleVersion = aws.String(_frauddetectorRuleVersion)
	}

	if disablePaginator() {
		if resp, err := client.GetRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*frauddetector.GetRulesOutput
	p := frauddetector.NewGetRulesPaginator(client, input)
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

// Gets all of the variables or the specific variable. This is a paginated API.
// Providing null maxSizePerPage results in retrieving maximum of 100 records per
// page. If you provide maxSizePerPage the value must be between 50 and 100. To
// get the next page result, a provide a pagination token from GetVariablesResult
// as part of your request. Null pagination token fetches the records from the
// beginning.
func frauddetector_GetVariables(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.GetVariablesInput{}

	if len(_frauddetectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _frauddetectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorName) > 0 {
		input.Name = aws.String(_frauddetectorName)
	}
	if len(_frauddetectorNextToken) > 0 {
		input.NextToken = aws.String(_frauddetectorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetVariables(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*frauddetector.GetVariablesOutput
	p := frauddetector.NewGetVariablesPaginator(client, input)
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

// Gets a list of past predictions. The list can be filtered by detector ID,
// detector version ID, event ID, event type, or by specifying a time period. If
// filter is not specified, the most recent prediction is returned.
//
// For example, the following filter lists all past predictions for xyz event type
// - { "eventType":{ "value": "xyz" }” }
//
// This is a paginated API. If you provide a null maxResults , this action will
// retrieve a maximum of 10 records per page. If you provide a maxResults , the
// value must be between 50 and 100. To get the next page results, provide the
// nextToken from the response as part of your request. A null nextToken fetches
// the records from the beginning.
func frauddetector_ListEventPredictions(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.ListEventPredictionsInput{}

	if len(_frauddetectorDetectorId) > 0 {
		if err := assignInputField(input, "DetectorId", _frauddetectorDetectorId); err != nil {
			log.Errorf("invalid --detector-id: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorDetectorVersionId) > 0 {
		if err := assignInputField(input, "DetectorVersionId", _frauddetectorDetectorVersionId); err != nil {
			log.Errorf("invalid --detector-version-id: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorEventId) > 0 {
		if err := assignInputField(input, "EventId", _frauddetectorEventId); err != nil {
			log.Errorf("invalid --event-id: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorEventType) > 0 {
		if err := assignInputField(input, "EventType", _frauddetectorEventType); err != nil {
			log.Errorf("invalid --event-type: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _frauddetectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorNextToken) > 0 {
		input.NextToken = aws.String(_frauddetectorNextToken)
	}
	if len(_frauddetectorPredictionTimeRange) > 0 {
		if err := assignInputField(input, "PredictionTimeRange", _frauddetectorPredictionTimeRange); err != nil {
			log.Errorf("invalid --prediction-time-range: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListEventPredictions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*frauddetector.ListEventPredictionsOutput
	p := frauddetector.NewListEventPredictionsPaginator(client, input)
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

// Lists all tags associated with the resource. This is a paginated API. To get
// the next page results, provide the pagination token from the response as part of
// your request. A null pagination token fetches the records from the beginning.
func frauddetector_ListTagsForResource(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_frauddetectorResourceARN) > 0 {
		input.ResourceARN = aws.String(_frauddetectorResourceARN)
	}
	if len(_frauddetectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _frauddetectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorNextToken) > 0 {
		input.NextToken = aws.String(_frauddetectorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*frauddetector.ListTagsForResourceOutput
	p := frauddetector.NewListTagsForResourcePaginator(client, input)
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

// Creates or updates a detector.
func frauddetector_PutDetector(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.PutDetectorInput{
		// DetectorId: *string, // Required
		// EventTypeName: *string, // Required
	}

	if len(_frauddetectorDetectorId) > 0 {
		input.DetectorId = aws.String(_frauddetectorDetectorId)
	}
	if len(_frauddetectorEventTypeName) > 0 {
		input.EventTypeName = aws.String(_frauddetectorEventTypeName)
	}
	if len(_frauddetectorDescription) > 0 {
		input.Description = aws.String(_frauddetectorDescription)
	}
	if len(_frauddetectorTags) > 0 {
		if err := assignInputField(input, "Tags", _frauddetectorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutDetector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates an entity type. An entity represents who is performing the
// event. As part of a fraud prediction, you pass the entity ID to indicate the
// specific entity who performed the event. An entity type classifies the entity.
// Example classifications include customer, merchant, or account.
func frauddetector_PutEntityType(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.PutEntityTypeInput{
		// Name: *string, // Required
	}

	if len(_frauddetectorName) > 0 {
		input.Name = aws.String(_frauddetectorName)
	}
	if len(_frauddetectorDescription) > 0 {
		input.Description = aws.String(_frauddetectorDescription)
	}
	if len(_frauddetectorTags) > 0 {
		if err := assignInputField(input, "Tags", _frauddetectorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutEntityType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates an event type. An event is a business activity that is
// evaluated for fraud risk. With Amazon Fraud Detector, you generate fraud
// predictions for events. An event type defines the structure for an event sent to
// Amazon Fraud Detector. This includes the variables sent as part of the event,
// the entity performing the event (such as a customer), and the labels that
// classify the event. Example event types include online payment transactions,
// account registrations, and authentications.
func frauddetector_PutEventType(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.PutEventTypeInput{
		// EntityTypes: []string, // Required
		// EventVariables: []string, // Required
		// Name: *string, // Required
	}

	if len(_frauddetectorEntityTypes) > 0 {
		input.EntityTypes = append([]string(nil), _frauddetectorEntityTypes...)
	}
	if len(_frauddetectorEventVariables) > 0 {
		input.EventVariables = []string{_frauddetectorEventVariables}
	}
	if len(_frauddetectorName) > 0 {
		input.Name = aws.String(_frauddetectorName)
	}
	if len(_frauddetectorDescription) > 0 {
		input.Description = aws.String(_frauddetectorDescription)
	}
	if len(_frauddetectorEventIngestion) > 0 {
		if err := assignInputField(input, "EventIngestion", _frauddetectorEventIngestion); err != nil {
			log.Errorf("invalid --event-ingestion: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorEventOrchestration) > 0 {
		if err := assignInputField(input, "EventOrchestration", _frauddetectorEventOrchestration); err != nil {
			log.Errorf("invalid --event-orchestration: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorLabels) > 0 {
		input.Labels = append([]string(nil), _frauddetectorLabels...)
	}
	if len(_frauddetectorTags) > 0 {
		if err := assignInputField(input, "Tags", _frauddetectorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutEventType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates an Amazon SageMaker model endpoint. You can also use this
// action to update the configuration of the model endpoint, including the IAM role
// and/or the mapped variables.
func frauddetector_PutExternalModel(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.PutExternalModelInput{
		// InputConfiguration: *types.ModelInputConfiguration, // Required
		// InvokeModelEndpointRoleArn: *string, // Required
		// ModelEndpoint: *string, // Required
		// ModelEndpointStatus: types.ModelEndpointStatus, // Required
		// ModelSource: types.ModelSource, // Required
		// OutputConfiguration: *types.ModelOutputConfiguration, // Required
	}

	if len(_frauddetectorInputConfiguration) > 0 {
		if err := assignInputField(input, "InputConfiguration", _frauddetectorInputConfiguration); err != nil {
			log.Errorf("invalid --input-configuration: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorInvokeModelEndpointRoleArn) > 0 {
		input.InvokeModelEndpointRoleArn = aws.String(_frauddetectorInvokeModelEndpointRoleArn)
	}
	if len(_frauddetectorModelEndpoint) > 0 {
		input.ModelEndpoint = aws.String(_frauddetectorModelEndpoint)
	}
	if len(_frauddetectorModelEndpointStatus) > 0 {
		if err := assignInputField(input, "ModelEndpointStatus", _frauddetectorModelEndpointStatus); err != nil {
			log.Errorf("invalid --model-endpoint-status: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorModelSource) > 0 {
		if err := assignInputField(input, "ModelSource", _frauddetectorModelSource); err != nil {
			log.Errorf("invalid --model-source: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorOutputConfiguration) > 0 {
		if err := assignInputField(input, "OutputConfiguration", _frauddetectorOutputConfiguration); err != nil {
			log.Errorf("invalid --output-configuration: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorTags) > 0 {
		if err := assignInputField(input, "Tags", _frauddetectorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutExternalModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specifies the KMS key to be used to encrypt content in Amazon Fraud Detector.
func frauddetector_PutKMSEncryptionKey(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.PutKMSEncryptionKeyInput{
		// KmsEncryptionKeyArn: *string, // Required
	}

	if len(_frauddetectorKmsEncryptionKeyArn) > 0 {
		input.KmsEncryptionKeyArn = aws.String(_frauddetectorKmsEncryptionKeyArn)
	}

	if resp, err := client.PutKMSEncryptionKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates label. A label classifies an event as fraudulent or
// legitimate. Labels are associated with event types and used to train supervised
// machine learning models in Amazon Fraud Detector.
func frauddetector_PutLabel(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.PutLabelInput{
		// Name: *string, // Required
	}

	if len(_frauddetectorName) > 0 {
		input.Name = aws.String(_frauddetectorName)
	}
	if len(_frauddetectorDescription) > 0 {
		input.Description = aws.String(_frauddetectorDescription)
	}
	if len(_frauddetectorTags) > 0 {
		if err := assignInputField(input, "Tags", _frauddetectorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutLabel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates an outcome.
func frauddetector_PutOutcome(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.PutOutcomeInput{
		// Name: *string, // Required
	}

	if len(_frauddetectorName) > 0 {
		input.Name = aws.String(_frauddetectorName)
	}
	if len(_frauddetectorDescription) > 0 {
		input.Description = aws.String(_frauddetectorDescription)
	}
	if len(_frauddetectorTags) > 0 {
		if err := assignInputField(input, "Tags", _frauddetectorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutOutcome(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stores events in Amazon Fraud Detector without generating fraud predictions for
// those events. For example, you can use SendEvent to upload a historical
// dataset, which you can then later use to train a model.
func frauddetector_SendEvent(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.SendEventInput{
		// Entities: []types.Entity, // Required
		// EventId: *string, // Required
		// EventTimestamp: *string, // Required
		// EventTypeName: *string, // Required
		// EventVariables: map[string]string, // Required
	}

	if len(_frauddetectorEntities) > 0 {
		if err := assignInputField(input, "Entities", _frauddetectorEntities); err != nil {
			log.Errorf("invalid --entities: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorEventId) > 0 {
		input.EventId = aws.String(_frauddetectorEventId)
	}
	if len(_frauddetectorEventTimestamp) > 0 {
		input.EventTimestamp = aws.String(_frauddetectorEventTimestamp)
	}
	if len(_frauddetectorEventTypeName) > 0 {
		input.EventTypeName = aws.String(_frauddetectorEventTypeName)
	}
	if len(_frauddetectorEventVariables) > 0 {
		if err := assignInputField(input, "EventVariables", _frauddetectorEventVariables); err != nil {
			log.Errorf("invalid --event-variables: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorAssignedLabel) > 0 {
		input.AssignedLabel = aws.String(_frauddetectorAssignedLabel)
	}
	if len(_frauddetectorLabelTimestamp) > 0 {
		input.LabelTimestamp = aws.String(_frauddetectorLabelTimestamp)
	}

	if resp, err := client.SendEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns tags to a resource.
func frauddetector_TagResource(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_frauddetectorResourceARN) > 0 {
		input.ResourceARN = aws.String(_frauddetectorResourceARN)
	}
	if len(_frauddetectorTags) > 0 {
		if err := assignInputField(input, "Tags", _frauddetectorTags); err != nil {
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

// Removes tags from a resource.
func frauddetector_UntagResource(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_frauddetectorResourceARN) > 0 {
		input.ResourceARN = aws.String(_frauddetectorResourceARN)
	}
	if len(_frauddetectorTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _frauddetectorTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a detector version. The detector version attributes that you can
// update include models, external model endpoints, rules, rule execution mode, and
// description. You can only update a DRAFT detector version.
func frauddetector_UpdateDetectorVersion(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.UpdateDetectorVersionInput{
		// DetectorId: *string, // Required
		// DetectorVersionId: *string, // Required
		// ExternalModelEndpoints: []string, // Required
		// Rules: []types.Rule, // Required
	}

	if len(_frauddetectorDetectorId) > 0 {
		input.DetectorId = aws.String(_frauddetectorDetectorId)
	}
	if len(_frauddetectorDetectorVersionId) > 0 {
		input.DetectorVersionId = aws.String(_frauddetectorDetectorVersionId)
	}
	if len(_frauddetectorExternalModelEndpoints) > 0 {
		input.ExternalModelEndpoints = append([]string(nil), _frauddetectorExternalModelEndpoints...)
	}
	if len(_frauddetectorRules) > 0 {
		if err := assignInputField(input, "Rules", _frauddetectorRules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorDescription) > 0 {
		input.Description = aws.String(_frauddetectorDescription)
	}
	if len(_frauddetectorModelVersions) > 0 {
		if err := assignInputField(input, "ModelVersions", _frauddetectorModelVersions); err != nil {
			log.Errorf("invalid --model-versions: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorRuleExecutionMode) > 0 {
		if err := assignInputField(input, "RuleExecutionMode", _frauddetectorRuleExecutionMode); err != nil {
			log.Errorf("invalid --rule-execution-mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDetectorVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the detector version's description. You can update the metadata for any
// detector version ( DRAFT, ACTIVE, or INACTIVE ).
func frauddetector_UpdateDetectorVersionMetadata(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.UpdateDetectorVersionMetadataInput{
		// Description: *string, // Required
		// DetectorId: *string, // Required
		// DetectorVersionId: *string, // Required
	}

	if len(_frauddetectorDescription) > 0 {
		input.Description = aws.String(_frauddetectorDescription)
	}
	if len(_frauddetectorDetectorId) > 0 {
		input.DetectorId = aws.String(_frauddetectorDetectorId)
	}
	if len(_frauddetectorDetectorVersionId) > 0 {
		input.DetectorVersionId = aws.String(_frauddetectorDetectorVersionId)
	}

	if resp, err := client.UpdateDetectorVersionMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the detector version’s status. You can perform the following promotions
// or demotions using UpdateDetectorVersionStatus : DRAFT to ACTIVE , ACTIVE to
// INACTIVE , and INACTIVE to ACTIVE .
func frauddetector_UpdateDetectorVersionStatus(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.UpdateDetectorVersionStatusInput{
		// DetectorId: *string, // Required
		// DetectorVersionId: *string, // Required
		// Status: types.DetectorVersionStatus, // Required
	}

	if len(_frauddetectorDetectorId) > 0 {
		input.DetectorId = aws.String(_frauddetectorDetectorId)
	}
	if len(_frauddetectorDetectorVersionId) > 0 {
		input.DetectorVersionId = aws.String(_frauddetectorDetectorVersionId)
	}
	if len(_frauddetectorStatus) > 0 {
		if err := assignInputField(input, "Status", _frauddetectorStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDetectorVersionStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified event with a new label.
func frauddetector_UpdateEventLabel(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.UpdateEventLabelInput{
		// AssignedLabel: *string, // Required
		// EventId: *string, // Required
		// EventTypeName: *string, // Required
		// LabelTimestamp: *string, // Required
	}

	if len(_frauddetectorAssignedLabel) > 0 {
		input.AssignedLabel = aws.String(_frauddetectorAssignedLabel)
	}
	if len(_frauddetectorEventId) > 0 {
		input.EventId = aws.String(_frauddetectorEventId)
	}
	if len(_frauddetectorEventTypeName) > 0 {
		input.EventTypeName = aws.String(_frauddetectorEventTypeName)
	}
	if len(_frauddetectorLabelTimestamp) > 0 {
		input.LabelTimestamp = aws.String(_frauddetectorLabelTimestamp)
	}

	if resp, err := client.UpdateEventLabel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a list.
func frauddetector_UpdateList(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.UpdateListInput{
		// Name: *string, // Required
	}

	if len(_frauddetectorName) > 0 {
		input.Name = aws.String(_frauddetectorName)
	}
	if len(_frauddetectorDescription) > 0 {
		input.Description = aws.String(_frauddetectorDescription)
	}
	if len(_frauddetectorElements) > 0 {
		input.Elements = append([]string(nil), _frauddetectorElements...)
	}
	if len(_frauddetectorUpdateMode) > 0 {
		if err := assignInputField(input, "UpdateMode", _frauddetectorUpdateMode); err != nil {
			log.Errorf("invalid --update-mode: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorVariableType) > 0 {
		input.VariableType = aws.String(_frauddetectorVariableType)
	}

	if resp, err := client.UpdateList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates model description.
func frauddetector_UpdateModel(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.UpdateModelInput{
		// ModelId: *string, // Required
		// ModelType: types.ModelTypeEnum, // Required
	}

	if len(_frauddetectorModelId) > 0 {
		input.ModelId = aws.String(_frauddetectorModelId)
	}
	if len(_frauddetectorModelType) > 0 {
		if err := assignInputField(input, "ModelType", _frauddetectorModelType); err != nil {
			log.Errorf("invalid --model-type: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorDescription) > 0 {
		input.Description = aws.String(_frauddetectorDescription)
	}

	if resp, err := client.UpdateModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a model version. Updating a model version retrains an existing model
// version using updated training data and produces a new minor version of the
// model. You can update the training data set location and data access role
// attributes using this action. This action creates and trains a new minor version
// of the model, for example version 1.01, 1.02, 1.03.
func frauddetector_UpdateModelVersion(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.UpdateModelVersionInput{
		// MajorVersionNumber: *string, // Required
		// ModelId: *string, // Required
		// ModelType: types.ModelTypeEnum, // Required
	}

	if len(_frauddetectorMajorVersionNumber) > 0 {
		input.MajorVersionNumber = aws.String(_frauddetectorMajorVersionNumber)
	}
	if len(_frauddetectorModelId) > 0 {
		input.ModelId = aws.String(_frauddetectorModelId)
	}
	if len(_frauddetectorModelType) > 0 {
		if err := assignInputField(input, "ModelType", _frauddetectorModelType); err != nil {
			log.Errorf("invalid --model-type: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorExternalEventsDetail) > 0 {
		if err := assignInputField(input, "ExternalEventsDetail", _frauddetectorExternalEventsDetail); err != nil {
			log.Errorf("invalid --external-events-detail: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorIngestedEventsDetail) > 0 {
		if err := assignInputField(input, "IngestedEventsDetail", _frauddetectorIngestedEventsDetail); err != nil {
			log.Errorf("invalid --ingested-events-detail: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorTags) > 0 {
		if err := assignInputField(input, "Tags", _frauddetectorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateModelVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the status of a model version.
// You can perform the following status updates:
//
// - Change the TRAINING_IN_PROGRESS status to TRAINING_CANCELLED .
//
// - Change the TRAINING_COMPLETE status to ACTIVE .
//
// - Change ACTIVE to INACTIVE .
func frauddetector_UpdateModelVersionStatus(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.UpdateModelVersionStatusInput{
		// ModelId: *string, // Required
		// ModelType: types.ModelTypeEnum, // Required
		// ModelVersionNumber: *string, // Required
		// Status: types.ModelVersionStatus, // Required
	}

	if len(_frauddetectorModelId) > 0 {
		input.ModelId = aws.String(_frauddetectorModelId)
	}
	if len(_frauddetectorModelType) > 0 {
		if err := assignInputField(input, "ModelType", _frauddetectorModelType); err != nil {
			log.Errorf("invalid --model-type: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorModelVersionNumber) > 0 {
		input.ModelVersionNumber = aws.String(_frauddetectorModelVersionNumber)
	}
	if len(_frauddetectorStatus) > 0 {
		if err := assignInputField(input, "Status", _frauddetectorStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateModelVersionStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a rule's metadata. The description attribute can be updated.
func frauddetector_UpdateRuleMetadata(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.UpdateRuleMetadataInput{
		// Description: *string, // Required
		// Rule: *types.Rule, // Required
	}

	if len(_frauddetectorDescription) > 0 {
		input.Description = aws.String(_frauddetectorDescription)
	}
	if len(_frauddetectorRule) > 0 {
		if err := assignInputField(input, "Rule", _frauddetectorRule); err != nil {
			log.Errorf("invalid --rule: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRuleMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a rule version resulting in a new rule version. Updates a rule version
// resulting in a new rule version (version 1, 2, 3 ...).
func frauddetector_UpdateRuleVersion(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.UpdateRuleVersionInput{
		// Expression: *string, // Required
		// Language: types.Language, // Required
		// Outcomes: []string, // Required
		// Rule: *types.Rule, // Required
	}

	if len(_frauddetectorExpression) > 0 {
		input.Expression = aws.String(_frauddetectorExpression)
	}
	if len(_frauddetectorLanguage) > 0 {
		if err := assignInputField(input, "Language", _frauddetectorLanguage); err != nil {
			log.Errorf("invalid --language: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorOutcomes) > 0 {
		input.Outcomes = append([]string(nil), _frauddetectorOutcomes...)
	}
	if len(_frauddetectorRule) > 0 {
		if err := assignInputField(input, "Rule", _frauddetectorRule); err != nil {
			log.Errorf("invalid --rule: %s", err.Error())
			return
		}
	}
	if len(_frauddetectorDescription) > 0 {
		input.Description = aws.String(_frauddetectorDescription)
	}
	if len(_frauddetectorTags) > 0 {
		if err := assignInputField(input, "Tags", _frauddetectorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRuleVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a variable.
func frauddetector_UpdateVariable(cfg aws.Config, client *frauddetector.Client) {
	input := &frauddetector.UpdateVariableInput{
		// Name: *string, // Required
	}

	if len(_frauddetectorName) > 0 {
		input.Name = aws.String(_frauddetectorName)
	}
	if len(_frauddetectorDefaultValue) > 0 {
		input.DefaultValue = aws.String(_frauddetectorDefaultValue)
	}
	if len(_frauddetectorDescription) > 0 {
		input.Description = aws.String(_frauddetectorDescription)
	}
	if len(_frauddetectorVariableType) > 0 {
		input.VariableType = aws.String(_frauddetectorVariableType)
	}

	if resp, err := client.UpdateVariable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_frauddetectorCmd)
	_frauddetectorCmd.Flags().SortFlags = false

	_frauddetectorCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_frauddetectorCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_frauddetectorCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorAssignedLabel, "assigned-label", "", "", "Assigned Label")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorDataSource, "data-source", "", "", "Data Source")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorDataType, "data-type", "", "", "Data Type")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorDefaultValue, "default-value", "", "", "Default Value")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorDeleteAuditHistory, "delete-audit-history", "", "", "Delete Audit History")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorDescription, "description", "", "", "Description")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorDetectorId, "detector-id", "", "", "Detector ID")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorDetectorName, "detector-name", "", "", "Detector Name")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorDetectorVersion, "detector-version", "", "", "Detector Version")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorDetectorVersionId, "detector-version-id", "", "", "Detector Version ID")
	_frauddetectorCmd.Flags().StringSliceVarP(&_frauddetectorElements, "elements", "", nil, "Elements")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorEntities, "entities", "", "", "Entities")
	_frauddetectorCmd.Flags().StringSliceVarP(&_frauddetectorEntityTypes, "entity-types", "", nil, "Entity Types")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorEventId, "event-id", "", "", "Event ID")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorEventIngestion, "event-ingestion", "", "", "Event Ingestion")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorEventOrchestration, "event-orchestration", "", "", "Event Orchestration")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorEventTimestamp, "event-timestamp", "", "", "Event Timestamp")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorEventType, "event-type", "", "", "Event Type")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorEventTypeName, "event-type-name", "", "", "Event Type Name")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorEventVariables, "event-variables", "", "", "Event Variables")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorExpression, "expression", "", "", "Expression")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorExternalEventsDetail, "external-events-detail", "", "", "External Events Detail")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorExternalModelEndpointDataBlobs, "external-model-endpoint-data-blobs", "", "", "External Model Endpoint Data Blobs")
	_frauddetectorCmd.Flags().StringSliceVarP(&_frauddetectorExternalModelEndpoints, "external-model-endpoints", "", nil, "External Model Endpoints")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorIamRoleArn, "iam-role-arn", "", "", "IAM Role ARN")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorIngestedEventsDetail, "ingested-events-detail", "", "", "Ingested Events Detail")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorInputConfiguration, "input-configuration", "", "", "Input Configuration")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorInputPath, "input-path", "", "", "Input Path")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorInvokeModelEndpointRoleArn, "invoke-model-endpoint-role-arn", "", "", "Invoke Model Endpoint Role ARN")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorJobId, "job-id", "", "", "Job ID")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorKmsEncryptionKeyArn, "kms-encryption-key-arn", "", "", "KMS Encryption Key ARN")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorLabelTimestamp, "label-timestamp", "", "", "Label Timestamp")
	_frauddetectorCmd.Flags().StringSliceVarP(&_frauddetectorLabels, "labels", "", nil, "Labels")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorLanguage, "language", "", "", "Language")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorMajorVersionNumber, "major-version-number", "", "", "Major Version Number")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorMaxResults, "max-results", "", "", "Max Results")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorModelEndpoint, "model-endpoint", "", "", "Model Endpoint")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorModelEndpointStatus, "model-endpoint-status", "", "", "Model Endpoint Status")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorModelId, "model-id", "", "", "Model ID")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorModelSource, "model-source", "", "", "Model Source")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorModelType, "model-type", "", "", "Model Type")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorModelVersionNumber, "model-version-number", "", "", "Model Version Number")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorModelVersions, "model-versions", "", "", "Model Versions")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorName, "name", "", "", "Name")
	_frauddetectorCmd.Flags().StringSliceVarP(&_frauddetectorNames, "names", "", nil, "Names")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorNextToken, "next-token", "", "", "Next Token")
	_frauddetectorCmd.Flags().StringSliceVarP(&_frauddetectorOutcomes, "outcomes", "", nil, "Outcomes")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorOutputConfiguration, "output-configuration", "", "", "Output Configuration")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorOutputPath, "output-path", "", "", "Output Path")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorPredictionTimeRange, "prediction-time-range", "", "", "Prediction Time Range")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorPredictionTimestamp, "prediction-timestamp", "", "", "Prediction Timestamp")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorResourceARN, "resource-arn", "", "", "Resource ARN")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorRule, "rule", "", "", "Rule")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorRuleExecutionMode, "rule-execution-mode", "", "", "Rule Execution Mode")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorRuleId, "rule-id", "", "", "Rule ID")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorRuleVersion, "rule-version", "", "", "Rule Version")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorRules, "rules", "", "", "Rules")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorStatus, "status", "", "", "Status")
	_frauddetectorCmd.Flags().StringSliceVarP(&_frauddetectorTagKeys, "tag-keys", "", nil, "Tag Keys")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorTags, "tags", "", "", "Tags")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorTrainingDataSchema, "training-data-schema", "", "", "Training Data Schema")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorTrainingDataSource, "training-data-source", "", "", "Training Data Source")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorUpdateMode, "update-mode", "", "", "Update Mode")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorVariableEntries, "variable-entries", "", "", "Variable Entries")
	_frauddetectorCmd.Flags().StringVarP(&_frauddetectorVariableType, "variable-type", "", "", "Variable Type")

	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorBatchCreateVariable, "batch-create-variable", "", false, "Batch Create Variable")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorBatchGetVariable, "batch-get-variable", "", false, "Batch Get Variable")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorCancelBatchImportJob, "cancel-batch-import-job", "", false, "Cancel Batch Import Job")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorCancelBatchPredictionJob, "cancel-batch-prediction-job", "", false, "Cancel Batch Prediction Job")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorCreateBatchImportJob, "create-batch-import-job", "", false, "Create Batch Import Job")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorCreateBatchPredictionJob, "create-batch-prediction-job", "", false, "Create Batch Prediction Job")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorCreateDetectorVersion, "create-detector-version", "", false, "Create Detector Version")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorCreateList, "create-list", "", false, "Create List")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorCreateModel, "create-model", "", false, "Create Model")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorCreateModelVersion, "create-model-version", "", false, "Create Model Version")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorCreateRule, "create-rule", "", false, "Create Rule")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorCreateVariable, "create-variable", "", false, "Create Variable")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorDeleteBatchImportJob, "delete-batch-import-job", "", false, "Delete Batch Import Job")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorDeleteBatchPredictionJob, "delete-batch-prediction-job", "", false, "Delete Batch Prediction Job")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorDeleteDetector, "delete-detector", "", false, "Delete Detector")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorDeleteDetectorVersion, "delete-detector-version", "", false, "Delete Detector Version")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorDeleteEntityType, "delete-entity-type", "", false, "Delete Entity Type")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorDeleteEvent, "delete-event", "", false, "Delete Event")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorDeleteEventType, "delete-event-type", "", false, "Delete Event Type")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorDeleteEventsByEventType, "delete-events-by-event-type", "", false, "Delete Events By Event Type")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorDeleteExternalModel, "delete-external-model", "", false, "Delete External Model")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorDeleteLabel, "delete-label", "", false, "Delete Label")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorDeleteList, "delete-list", "", false, "Delete List")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorDeleteModel, "delete-model", "", false, "Delete Model")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorDeleteModelVersion, "delete-model-version", "", false, "Delete Model Version")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorDeleteOutcome, "delete-outcome", "", false, "Delete Outcome")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorDeleteRule, "delete-rule", "", false, "Delete Rule")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorDeleteVariable, "delete-variable", "", false, "Delete Variable")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorDescribeDetector, "describe-detector", "", false, "Describe Detector")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorDescribeModelVersions, "describe-model-versions", "", false, "Describe Model Versions")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorGetBatchImportJobs, "get-batch-import-jobs", "", false, "Get Batch Import Jobs")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorGetBatchPredictionJobs, "get-batch-prediction-jobs", "", false, "Get Batch Prediction Jobs")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorGetDeleteEventsByEventTypeStatus, "get-delete-events-by-event-type-status", "", false, "Get Delete Events By Event Type Status")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorGetDetectorVersion, "get-detector-version", "", false, "Get Detector Version")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorGetDetectors, "get-detectors", "", false, "Get Detectors")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorGetEntityTypes, "get-entity-types", "", false, "Get Entity Types")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorGetEvent, "get-event", "", false, "Get Event")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorGetEventPrediction, "get-event-prediction", "", false, "Get Event Prediction")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorGetEventPredictionMetadata, "get-event-prediction-metadata", "", false, "Get Event Prediction Metadata")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorGetEventTypes, "get-event-types", "", false, "Get Event Types")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorGetExternalModels, "get-external-models", "", false, "Get External Models")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorGetKMSEncryptionKey, "get-kms-encryption-key", "", false, "Get KMS Encryption Key")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorGetLabels, "get-labels", "", false, "Get Labels")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorGetListElements, "get-list-elements", "", false, "Get List Elements")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorGetListsMetadata, "get-lists-metadata", "", false, "Get Lists Metadata")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorGetModelVersion, "get-model-version", "", false, "Get Model Version")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorGetModels, "get-models", "", false, "Get Models")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorGetOutcomes, "get-outcomes", "", false, "Get Outcomes")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorGetRules, "get-rules", "", false, "Get Rules")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorGetVariables, "get-variables", "", false, "Get Variables")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorListEventPredictions, "list-event-predictions", "", false, "List Event Predictions")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorPutDetector, "put-detector", "", false, "Put Detector")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorPutEntityType, "put-entity-type", "", false, "Put Entity Type")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorPutEventType, "put-event-type", "", false, "Put Event Type")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorPutExternalModel, "put-external-model", "", false, "Put External Model")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorPutKMSEncryptionKey, "put-kms-encryption-key", "", false, "Put KMS Encryption Key")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorPutLabel, "put-label", "", false, "Put Label")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorPutOutcome, "put-outcome", "", false, "Put Outcome")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorSendEvent, "send-event", "", false, "Send Event")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorTagResource, "tag-resource", "", false, "Tag Resource")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorUntagResource, "untag-resource", "", false, "Untag Resource")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorUpdateDetectorVersion, "update-detector-version", "", false, "Update Detector Version")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorUpdateDetectorVersionMetadata, "update-detector-version-metadata", "", false, "Update Detector Version Metadata")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorUpdateDetectorVersionStatus, "update-detector-version-status", "", false, "Update Detector Version Status")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorUpdateEventLabel, "update-event-label", "", false, "Update Event Label")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorUpdateList, "update-list", "", false, "Update List")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorUpdateModel, "update-model", "", false, "Update Model")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorUpdateModelVersion, "update-model-version", "", false, "Update Model Version")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorUpdateModelVersionStatus, "update-model-version-status", "", false, "Update Model Version Status")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorUpdateRuleMetadata, "update-rule-metadata", "", false, "Update Rule Metadata")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorUpdateRuleVersion, "update-rule-version", "", false, "Update Rule Version")
	_frauddetectorCmd.Flags().BoolVarP(&_frauddetectorUpdateVariable, "update-variable", "", false, "Update Variable")

}
