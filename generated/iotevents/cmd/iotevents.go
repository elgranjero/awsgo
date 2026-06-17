package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotevents"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ioteventsCmd represents the iotevents command
var _ioteventsCmd = &cobra.Command{
	Use:   "iotevents",
	Short: "AWS iotevents CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := iotevents.NewFromConfig(cfg)
		if _ioteventsCreateAlarmModel {
			iotevents_CreateAlarmModel(cfg, client)
			return
		}
		if _ioteventsCreateDetectorModel {
			iotevents_CreateDetectorModel(cfg, client)
			return
		}
		if _ioteventsCreateInput {
			iotevents_CreateInput(cfg, client)
			return
		}
		if _ioteventsDeleteAlarmModel {
			iotevents_DeleteAlarmModel(cfg, client)
			return
		}
		if _ioteventsDeleteDetectorModel {
			iotevents_DeleteDetectorModel(cfg, client)
			return
		}
		if _ioteventsDeleteInput {
			iotevents_DeleteInput(cfg, client)
			return
		}
		if _ioteventsDescribeAlarmModel {
			iotevents_DescribeAlarmModel(cfg, client)
			return
		}
		if _ioteventsDescribeDetectorModel {
			iotevents_DescribeDetectorModel(cfg, client)
			return
		}
		if _ioteventsDescribeDetectorModelAnalysis {
			iotevents_DescribeDetectorModelAnalysis(cfg, client)
			return
		}
		if _ioteventsDescribeInput {
			iotevents_DescribeInput(cfg, client)
			return
		}
		if _ioteventsDescribeLoggingOptions {
			iotevents_DescribeLoggingOptions(cfg, client)
			return
		}
		if _ioteventsGetDetectorModelAnalysisResults {
			iotevents_GetDetectorModelAnalysisResults(cfg, client)
			return
		}
		if _ioteventsListAlarmModelVersions {
			iotevents_ListAlarmModelVersions(cfg, client)
			return
		}
		if _ioteventsListAlarmModels {
			iotevents_ListAlarmModels(cfg, client)
			return
		}
		if _ioteventsListDetectorModelVersions {
			iotevents_ListDetectorModelVersions(cfg, client)
			return
		}
		if _ioteventsListDetectorModels {
			iotevents_ListDetectorModels(cfg, client)
			return
		}
		if _ioteventsListInputRoutings {
			iotevents_ListInputRoutings(cfg, client)
			return
		}
		if _ioteventsListInputs {
			iotevents_ListInputs(cfg, client)
			return
		}
		if _ioteventsListTagsForResource {
			iotevents_ListTagsForResource(cfg, client)
			return
		}
		if _ioteventsPutLoggingOptions {
			iotevents_PutLoggingOptions(cfg, client)
			return
		}
		if _ioteventsStartDetectorModelAnalysis {
			iotevents_StartDetectorModelAnalysis(cfg, client)
			return
		}
		if _ioteventsTagResource {
			iotevents_TagResource(cfg, client)
			return
		}
		if _ioteventsUntagResource {
			iotevents_UntagResource(cfg, client)
			return
		}
		if _ioteventsUpdateAlarmModel {
			iotevents_UpdateAlarmModel(cfg, client)
			return
		}
		if _ioteventsUpdateDetectorModel {
			iotevents_UpdateDetectorModel(cfg, client)
			return
		}
		if _ioteventsUpdateInput {
			iotevents_UpdateInput(cfg, client)
			return
		}

	},
}

var (
	_ioteventsCreateAlarmModel                bool
	_ioteventsCreateDetectorModel             bool
	_ioteventsCreateInput                     bool
	_ioteventsDeleteAlarmModel                bool
	_ioteventsDeleteDetectorModel             bool
	_ioteventsDeleteInput                     bool
	_ioteventsDescribeAlarmModel              bool
	_ioteventsDescribeDetectorModel           bool
	_ioteventsDescribeDetectorModelAnalysis   bool
	_ioteventsDescribeInput                   bool
	_ioteventsDescribeLoggingOptions          bool
	_ioteventsGetDetectorModelAnalysisResults bool
	_ioteventsListAlarmModelVersions          bool
	_ioteventsListAlarmModels                 bool
	_ioteventsListDetectorModelVersions       bool
	_ioteventsListDetectorModels              bool
	_ioteventsListInputRoutings               bool
	_ioteventsListInputs                      bool
	_ioteventsListTagsForResource             bool
	_ioteventsPutLoggingOptions               bool
	_ioteventsStartDetectorModelAnalysis      bool
	_ioteventsTagResource                     bool
	_ioteventsUntagResource                   bool
	_ioteventsUpdateAlarmModel                bool
	_ioteventsUpdateDetectorModel             bool
	_ioteventsUpdateInput                     bool

	_ioteventsAlarmCapabilities        string
	_ioteventsAlarmEventActions        string
	_ioteventsAlarmModelDescription    string
	_ioteventsAlarmModelName           string
	_ioteventsAlarmModelVersion        string
	_ioteventsAlarmNotification        string
	_ioteventsAlarmRule                string
	_ioteventsAnalysisId               string
	_ioteventsDetectorModelDefinition  string
	_ioteventsDetectorModelDescription string
	_ioteventsDetectorModelName        string
	_ioteventsDetectorModelVersion     string
	_ioteventsEvaluationMethod         string
	_ioteventsInputDefinition          string
	_ioteventsInputDescription         string
	_ioteventsInputIdentifier          string
	_ioteventsInputName                string
	_ioteventsKey                      string
	_ioteventsLoggingOptions           string
	_ioteventsMaxResults               string
	_ioteventsNextToken                string
	_ioteventsResourceArn              string
	_ioteventsRoleArn                  string
	_ioteventsSeverity                 string
	_ioteventsTagKeys                  []string
	_ioteventsTags                     string
)

// Creates an alarm model to monitor an AWS IoT Events input attribute. You can
// use the alarm to get notified when the value is outside a specified range. For
// more information, see [Create an alarm model]in the AWS IoT Events Developer Guide.
//
// [Create an alarm model]: https://docs.aws.amazon.com/iotevents/latest/developerguide/create-alarms.html
func iotevents_CreateAlarmModel(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.CreateAlarmModelInput{
		// AlarmModelName: *string, // Required
		// AlarmRule: *types.AlarmRule, // Required
		// RoleArn: *string, // Required
	}

	if len(_ioteventsAlarmModelName) > 0 {
		input.AlarmModelName = aws.String(_ioteventsAlarmModelName)
	}
	if len(_ioteventsAlarmRule) > 0 {
		if err := assignInputField(input, "AlarmRule", _ioteventsAlarmRule); err != nil {
			log.Errorf("invalid --alarm-rule: %s", err.Error())
			return
		}
	}
	if len(_ioteventsRoleArn) > 0 {
		input.RoleArn = aws.String(_ioteventsRoleArn)
	}
	if len(_ioteventsAlarmCapabilities) > 0 {
		if err := assignInputField(input, "AlarmCapabilities", _ioteventsAlarmCapabilities); err != nil {
			log.Errorf("invalid --alarm-capabilities: %s", err.Error())
			return
		}
	}
	if len(_ioteventsAlarmEventActions) > 0 {
		if err := assignInputField(input, "AlarmEventActions", _ioteventsAlarmEventActions); err != nil {
			log.Errorf("invalid --alarm-event-actions: %s", err.Error())
			return
		}
	}
	if len(_ioteventsAlarmModelDescription) > 0 {
		input.AlarmModelDescription = aws.String(_ioteventsAlarmModelDescription)
	}
	if len(_ioteventsAlarmNotification) > 0 {
		if err := assignInputField(input, "AlarmNotification", _ioteventsAlarmNotification); err != nil {
			log.Errorf("invalid --alarm-notification: %s", err.Error())
			return
		}
	}
	if len(_ioteventsKey) > 0 {
		input.Key = aws.String(_ioteventsKey)
	}
	if len(_ioteventsSeverity) > 0 {
		if err := assignInputField(input, "Severity", _ioteventsSeverity); err != nil {
			log.Errorf("invalid --severity: %s", err.Error())
			return
		}
	}
	if len(_ioteventsTags) > 0 {
		if err := assignInputField(input, "Tags", _ioteventsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAlarmModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a detector model.
func iotevents_CreateDetectorModel(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.CreateDetectorModelInput{
		// DetectorModelDefinition: *types.DetectorModelDefinition, // Required
		// DetectorModelName: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_ioteventsDetectorModelDefinition) > 0 {
		if err := assignInputField(input, "DetectorModelDefinition", _ioteventsDetectorModelDefinition); err != nil {
			log.Errorf("invalid --detector-model-definition: %s", err.Error())
			return
		}
	}
	if len(_ioteventsDetectorModelName) > 0 {
		input.DetectorModelName = aws.String(_ioteventsDetectorModelName)
	}
	if len(_ioteventsRoleArn) > 0 {
		input.RoleArn = aws.String(_ioteventsRoleArn)
	}
	if len(_ioteventsDetectorModelDescription) > 0 {
		input.DetectorModelDescription = aws.String(_ioteventsDetectorModelDescription)
	}
	if len(_ioteventsEvaluationMethod) > 0 {
		if err := assignInputField(input, "EvaluationMethod", _ioteventsEvaluationMethod); err != nil {
			log.Errorf("invalid --evaluation-method: %s", err.Error())
			return
		}
	}
	if len(_ioteventsKey) > 0 {
		input.Key = aws.String(_ioteventsKey)
	}
	if len(_ioteventsTags) > 0 {
		if err := assignInputField(input, "Tags", _ioteventsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDetectorModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an input.
func iotevents_CreateInput(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.CreateInputInput{
		// InputDefinition: *types.InputDefinition, // Required
		// InputName: *string, // Required
	}

	if len(_ioteventsInputDefinition) > 0 {
		if err := assignInputField(input, "InputDefinition", _ioteventsInputDefinition); err != nil {
			log.Errorf("invalid --input-definition: %s", err.Error())
			return
		}
	}
	if len(_ioteventsInputName) > 0 {
		input.InputName = aws.String(_ioteventsInputName)
	}
	if len(_ioteventsInputDescription) > 0 {
		input.InputDescription = aws.String(_ioteventsInputDescription)
	}
	if len(_ioteventsTags) > 0 {
		if err := assignInputField(input, "Tags", _ioteventsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateInput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an alarm model. Any alarm instances that were created based on this
// alarm model are also deleted. This action can't be undone.
func iotevents_DeleteAlarmModel(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.DeleteAlarmModelInput{
		// AlarmModelName: *string, // Required
	}

	if len(_ioteventsAlarmModelName) > 0 {
		input.AlarmModelName = aws.String(_ioteventsAlarmModelName)
	}

	if resp, err := client.DeleteAlarmModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a detector model. Any active instances of the detector model are also
// deleted.
func iotevents_DeleteDetectorModel(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.DeleteDetectorModelInput{
		// DetectorModelName: *string, // Required
	}

	if len(_ioteventsDetectorModelName) > 0 {
		input.DetectorModelName = aws.String(_ioteventsDetectorModelName)
	}

	if resp, err := client.DeleteDetectorModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an input.
func iotevents_DeleteInput(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.DeleteInputInput{
		// InputName: *string, // Required
	}

	if len(_ioteventsInputName) > 0 {
		input.InputName = aws.String(_ioteventsInputName)
	}

	if resp, err := client.DeleteInput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an alarm model. If you don't specify a value for
// the alarmModelVersion parameter, the latest version is returned.
func iotevents_DescribeAlarmModel(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.DescribeAlarmModelInput{
		// AlarmModelName: *string, // Required
	}

	if len(_ioteventsAlarmModelName) > 0 {
		input.AlarmModelName = aws.String(_ioteventsAlarmModelName)
	}
	if len(_ioteventsAlarmModelVersion) > 0 {
		input.AlarmModelVersion = aws.String(_ioteventsAlarmModelVersion)
	}

	if resp, err := client.DescribeAlarmModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a detector model. If the version parameter is not specified,
// information about the latest version is returned.
func iotevents_DescribeDetectorModel(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.DescribeDetectorModelInput{
		// DetectorModelName: *string, // Required
	}

	if len(_ioteventsDetectorModelName) > 0 {
		input.DetectorModelName = aws.String(_ioteventsDetectorModelName)
	}
	if len(_ioteventsDetectorModelVersion) > 0 {
		input.DetectorModelVersion = aws.String(_ioteventsDetectorModelVersion)
	}

	if resp, err := client.DescribeDetectorModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves runtime information about a detector model analysis.
// After AWS IoT Events starts analyzing your detector model, you have up to 24
// hours to retrieve the analysis results.
func iotevents_DescribeDetectorModelAnalysis(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.DescribeDetectorModelAnalysisInput{
		// AnalysisId: *string, // Required
	}

	if len(_ioteventsAnalysisId) > 0 {
		input.AnalysisId = aws.String(_ioteventsAnalysisId)
	}

	if resp, err := client.DescribeDetectorModelAnalysis(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an input.
func iotevents_DescribeInput(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.DescribeInputInput{
		// InputName: *string, // Required
	}

	if len(_ioteventsInputName) > 0 {
		input.InputName = aws.String(_ioteventsInputName)
	}

	if resp, err := client.DescribeInput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the current settings of the AWS IoT Events logging options.
func iotevents_DescribeLoggingOptions(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.DescribeLoggingOptionsInput{}

	if resp, err := client.DescribeLoggingOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves one or more analysis results of the detector model.
// After AWS IoT Events starts analyzing your detector model, you have up to 24
// hours to retrieve the analysis results.
func iotevents_GetDetectorModelAnalysisResults(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.GetDetectorModelAnalysisResultsInput{
		// AnalysisId: *string, // Required
	}

	if len(_ioteventsAnalysisId) > 0 {
		input.AnalysisId = aws.String(_ioteventsAnalysisId)
	}
	if len(_ioteventsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ioteventsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ioteventsNextToken) > 0 {
		input.NextToken = aws.String(_ioteventsNextToken)
	}

	if resp, err := client.GetDetectorModelAnalysisResults(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the versions of an alarm model. The operation returns only the
// metadata associated with each alarm model version.
func iotevents_ListAlarmModelVersions(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.ListAlarmModelVersionsInput{
		// AlarmModelName: *string, // Required
	}

	if len(_ioteventsAlarmModelName) > 0 {
		input.AlarmModelName = aws.String(_ioteventsAlarmModelName)
	}
	if len(_ioteventsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ioteventsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ioteventsNextToken) > 0 {
		input.NextToken = aws.String(_ioteventsNextToken)
	}

	if resp, err := client.ListAlarmModelVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the alarm models that you created. The operation returns only the
// metadata associated with each alarm model.
func iotevents_ListAlarmModels(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.ListAlarmModelsInput{}

	if len(_ioteventsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ioteventsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ioteventsNextToken) > 0 {
		input.NextToken = aws.String(_ioteventsNextToken)
	}

	if resp, err := client.ListAlarmModels(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the versions of a detector model. Only the metadata associated with
// each detector model version is returned.
func iotevents_ListDetectorModelVersions(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.ListDetectorModelVersionsInput{
		// DetectorModelName: *string, // Required
	}

	if len(_ioteventsDetectorModelName) > 0 {
		input.DetectorModelName = aws.String(_ioteventsDetectorModelName)
	}
	if len(_ioteventsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ioteventsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ioteventsNextToken) > 0 {
		input.NextToken = aws.String(_ioteventsNextToken)
	}

	if resp, err := client.ListDetectorModelVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the detector models you have created. Only the metadata associated with
// each detector model is returned.
func iotevents_ListDetectorModels(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.ListDetectorModelsInput{}

	if len(_ioteventsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ioteventsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ioteventsNextToken) > 0 {
		input.NextToken = aws.String(_ioteventsNextToken)
	}

	if resp, err := client.ListDetectorModels(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists one or more input routings.
func iotevents_ListInputRoutings(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.ListInputRoutingsInput{
		// InputIdentifier: *types.InputIdentifier, // Required
	}

	if len(_ioteventsInputIdentifier) > 0 {
		if err := assignInputField(input, "InputIdentifier", _ioteventsInputIdentifier); err != nil {
			log.Errorf("invalid --input-identifier: %s", err.Error())
			return
		}
	}
	if len(_ioteventsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ioteventsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ioteventsNextToken) > 0 {
		input.NextToken = aws.String(_ioteventsNextToken)
	}

	if resp, err := client.ListInputRoutings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the inputs you have created.
func iotevents_ListInputs(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.ListInputsInput{}

	if len(_ioteventsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ioteventsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ioteventsNextToken) > 0 {
		input.NextToken = aws.String(_ioteventsNextToken)
	}

	if resp, err := client.ListInputs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the tags (metadata) you have assigned to the resource.
func iotevents_ListTagsForResource(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_ioteventsResourceArn) > 0 {
		input.ResourceArn = aws.String(_ioteventsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets or updates the AWS IoT Events logging options.
// If you update the value of any loggingOptions field, it takes up to one minute
// for the change to take effect. If you change the policy attached to the role you
// specified in the roleArn field (for example, to correct an invalid policy), it
// takes up to five minutes for that change to take effect.
func iotevents_PutLoggingOptions(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.PutLoggingOptionsInput{
		// LoggingOptions: *types.LoggingOptions, // Required
	}

	if len(_ioteventsLoggingOptions) > 0 {
		if err := assignInputField(input, "LoggingOptions", _ioteventsLoggingOptions); err != nil {
			log.Errorf("invalid --logging-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutLoggingOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Performs an analysis of your detector model. For more information, see [Troubleshooting a detector model] in the
// AWS IoT Events Developer Guide.
//
// [Troubleshooting a detector model]: https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-analyze-api.html
func iotevents_StartDetectorModelAnalysis(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.StartDetectorModelAnalysisInput{
		// DetectorModelDefinition: *types.DetectorModelDefinition, // Required
	}

	if len(_ioteventsDetectorModelDefinition) > 0 {
		if err := assignInputField(input, "DetectorModelDefinition", _ioteventsDetectorModelDefinition); err != nil {
			log.Errorf("invalid --detector-model-definition: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartDetectorModelAnalysis(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds to or modifies the tags of the given resource. Tags are metadata that can
// be used to manage a resource.
func iotevents_TagResource(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_ioteventsResourceArn) > 0 {
		input.ResourceArn = aws.String(_ioteventsResourceArn)
	}
	if len(_ioteventsTags) > 0 {
		if err := assignInputField(input, "Tags", _ioteventsTags); err != nil {
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

// Removes the given tags (metadata) from the resource.
func iotevents_UntagResource(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_ioteventsResourceArn) > 0 {
		input.ResourceArn = aws.String(_ioteventsResourceArn)
	}
	if len(_ioteventsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _ioteventsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an alarm model. Any alarms that were created based on the previous
// version are deleted and then created again as new data arrives.
func iotevents_UpdateAlarmModel(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.UpdateAlarmModelInput{
		// AlarmModelName: *string, // Required
		// AlarmRule: *types.AlarmRule, // Required
		// RoleArn: *string, // Required
	}

	if len(_ioteventsAlarmModelName) > 0 {
		input.AlarmModelName = aws.String(_ioteventsAlarmModelName)
	}
	if len(_ioteventsAlarmRule) > 0 {
		if err := assignInputField(input, "AlarmRule", _ioteventsAlarmRule); err != nil {
			log.Errorf("invalid --alarm-rule: %s", err.Error())
			return
		}
	}
	if len(_ioteventsRoleArn) > 0 {
		input.RoleArn = aws.String(_ioteventsRoleArn)
	}
	if len(_ioteventsAlarmCapabilities) > 0 {
		if err := assignInputField(input, "AlarmCapabilities", _ioteventsAlarmCapabilities); err != nil {
			log.Errorf("invalid --alarm-capabilities: %s", err.Error())
			return
		}
	}
	if len(_ioteventsAlarmEventActions) > 0 {
		if err := assignInputField(input, "AlarmEventActions", _ioteventsAlarmEventActions); err != nil {
			log.Errorf("invalid --alarm-event-actions: %s", err.Error())
			return
		}
	}
	if len(_ioteventsAlarmModelDescription) > 0 {
		input.AlarmModelDescription = aws.String(_ioteventsAlarmModelDescription)
	}
	if len(_ioteventsAlarmNotification) > 0 {
		if err := assignInputField(input, "AlarmNotification", _ioteventsAlarmNotification); err != nil {
			log.Errorf("invalid --alarm-notification: %s", err.Error())
			return
		}
	}
	if len(_ioteventsSeverity) > 0 {
		if err := assignInputField(input, "Severity", _ioteventsSeverity); err != nil {
			log.Errorf("invalid --severity: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAlarmModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a detector model. Detectors (instances) spawned by the previous version
// are deleted and then re-created as new inputs arrive.
func iotevents_UpdateDetectorModel(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.UpdateDetectorModelInput{
		// DetectorModelDefinition: *types.DetectorModelDefinition, // Required
		// DetectorModelName: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_ioteventsDetectorModelDefinition) > 0 {
		if err := assignInputField(input, "DetectorModelDefinition", _ioteventsDetectorModelDefinition); err != nil {
			log.Errorf("invalid --detector-model-definition: %s", err.Error())
			return
		}
	}
	if len(_ioteventsDetectorModelName) > 0 {
		input.DetectorModelName = aws.String(_ioteventsDetectorModelName)
	}
	if len(_ioteventsRoleArn) > 0 {
		input.RoleArn = aws.String(_ioteventsRoleArn)
	}
	if len(_ioteventsDetectorModelDescription) > 0 {
		input.DetectorModelDescription = aws.String(_ioteventsDetectorModelDescription)
	}
	if len(_ioteventsEvaluationMethod) > 0 {
		if err := assignInputField(input, "EvaluationMethod", _ioteventsEvaluationMethod); err != nil {
			log.Errorf("invalid --evaluation-method: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDetectorModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an input.
func iotevents_UpdateInput(cfg aws.Config, client *iotevents.Client) {
	input := &iotevents.UpdateInputInput{
		// InputDefinition: *types.InputDefinition, // Required
		// InputName: *string, // Required
	}

	if len(_ioteventsInputDefinition) > 0 {
		if err := assignInputField(input, "InputDefinition", _ioteventsInputDefinition); err != nil {
			log.Errorf("invalid --input-definition: %s", err.Error())
			return
		}
	}
	if len(_ioteventsInputName) > 0 {
		input.InputName = aws.String(_ioteventsInputName)
	}
	if len(_ioteventsInputDescription) > 0 {
		input.InputDescription = aws.String(_ioteventsInputDescription)
	}

	if resp, err := client.UpdateInput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_ioteventsCmd)
	_ioteventsCmd.Flags().SortFlags = false

	_ioteventsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_ioteventsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_ioteventsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_ioteventsCmd.Flags().StringVarP(&_ioteventsAlarmCapabilities, "alarm-capabilities", "", "", "Alarm Capabilities")
	_ioteventsCmd.Flags().StringVarP(&_ioteventsAlarmEventActions, "alarm-event-actions", "", "", "Alarm Event Actions")
	_ioteventsCmd.Flags().StringVarP(&_ioteventsAlarmModelDescription, "alarm-model-description", "", "", "Alarm Model Description")
	_ioteventsCmd.Flags().StringVarP(&_ioteventsAlarmModelName, "alarm-model-name", "", "", "Alarm Model Name")
	_ioteventsCmd.Flags().StringVarP(&_ioteventsAlarmModelVersion, "alarm-model-version", "", "", "Alarm Model Version")
	_ioteventsCmd.Flags().StringVarP(&_ioteventsAlarmNotification, "alarm-notification", "", "", "Alarm Notification")
	_ioteventsCmd.Flags().StringVarP(&_ioteventsAlarmRule, "alarm-rule", "", "", "Alarm Rule")
	_ioteventsCmd.Flags().StringVarP(&_ioteventsAnalysisId, "analysis-id", "", "", "Analysis ID")
	_ioteventsCmd.Flags().StringVarP(&_ioteventsDetectorModelDefinition, "detector-model-definition", "", "", "Detector Model Definition")
	_ioteventsCmd.Flags().StringVarP(&_ioteventsDetectorModelDescription, "detector-model-description", "", "", "Detector Model Description")
	_ioteventsCmd.Flags().StringVarP(&_ioteventsDetectorModelName, "detector-model-name", "", "", "Detector Model Name")
	_ioteventsCmd.Flags().StringVarP(&_ioteventsDetectorModelVersion, "detector-model-version", "", "", "Detector Model Version")
	_ioteventsCmd.Flags().StringVarP(&_ioteventsEvaluationMethod, "evaluation-method", "", "", "Evaluation Method")
	_ioteventsCmd.Flags().StringVarP(&_ioteventsInputDefinition, "input-definition", "", "", "Input Definition")
	_ioteventsCmd.Flags().StringVarP(&_ioteventsInputDescription, "input-description", "", "", "Input Description")
	_ioteventsCmd.Flags().StringVarP(&_ioteventsInputIdentifier, "input-identifier", "", "", "Input Identifier")
	_ioteventsCmd.Flags().StringVarP(&_ioteventsInputName, "input-name", "", "", "Input Name")
	_ioteventsCmd.Flags().StringVarP(&_ioteventsKey, "key", "", "", "Key")
	_ioteventsCmd.Flags().StringVarP(&_ioteventsLoggingOptions, "logging-options", "", "", "Logging Options")
	_ioteventsCmd.Flags().StringVarP(&_ioteventsMaxResults, "max-results", "", "", "Max Results")
	_ioteventsCmd.Flags().StringVarP(&_ioteventsNextToken, "next-token", "", "", "Next Token")
	_ioteventsCmd.Flags().StringVarP(&_ioteventsResourceArn, "resource-arn", "", "", "Resource ARN")
	_ioteventsCmd.Flags().StringVarP(&_ioteventsRoleArn, "role-arn", "", "", "Role ARN")
	_ioteventsCmd.Flags().StringVarP(&_ioteventsSeverity, "severity", "", "", "Severity")
	_ioteventsCmd.Flags().StringSliceVarP(&_ioteventsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_ioteventsCmd.Flags().StringVarP(&_ioteventsTags, "tags", "", "", "Tags")

	_ioteventsCmd.Flags().BoolVarP(&_ioteventsCreateAlarmModel, "create-alarm-model", "", false, "Create Alarm Model")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsCreateDetectorModel, "create-detector-model", "", false, "Create Detector Model")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsCreateInput, "create-input", "", false, "Create Input")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsDeleteAlarmModel, "delete-alarm-model", "", false, "Delete Alarm Model")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsDeleteDetectorModel, "delete-detector-model", "", false, "Delete Detector Model")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsDeleteInput, "delete-input", "", false, "Delete Input")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsDescribeAlarmModel, "describe-alarm-model", "", false, "Describe Alarm Model")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsDescribeDetectorModel, "describe-detector-model", "", false, "Describe Detector Model")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsDescribeDetectorModelAnalysis, "describe-detector-model-analysis", "", false, "Describe Detector Model Analysis")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsDescribeInput, "describe-input", "", false, "Describe Input")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsDescribeLoggingOptions, "describe-logging-options", "", false, "Describe Logging Options")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsGetDetectorModelAnalysisResults, "get-detector-model-analysis-results", "", false, "Get Detector Model Analysis Results")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsListAlarmModelVersions, "list-alarm-model-versions", "", false, "List Alarm Model Versions")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsListAlarmModels, "list-alarm-models", "", false, "List Alarm Models")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsListDetectorModelVersions, "list-detector-model-versions", "", false, "List Detector Model Versions")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsListDetectorModels, "list-detector-models", "", false, "List Detector Models")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsListInputRoutings, "list-input-routings", "", false, "List Input Routings")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsListInputs, "list-inputs", "", false, "List Inputs")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsPutLoggingOptions, "put-logging-options", "", false, "Put Logging Options")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsStartDetectorModelAnalysis, "start-detector-model-analysis", "", false, "Start Detector Model Analysis")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsTagResource, "tag-resource", "", false, "Tag Resource")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsUntagResource, "untag-resource", "", false, "Untag Resource")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsUpdateAlarmModel, "update-alarm-model", "", false, "Update Alarm Model")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsUpdateDetectorModel, "update-detector-model", "", false, "Update Detector Model")
	_ioteventsCmd.Flags().BoolVarP(&_ioteventsUpdateInput, "update-input", "", false, "Update Input")

}
