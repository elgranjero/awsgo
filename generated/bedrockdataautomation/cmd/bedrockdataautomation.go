package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockdataautomation"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// bedrockdataautomationCmd represents the bedrockdataautomation command
var _bedrockdataautomationCmd = &cobra.Command{
	Use:   "bedrockdataautomation",
	Short: "AWS bedrockdataautomation CLI",
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
		client := bedrockdataautomation.NewFromConfig(cfg)
		if _bedrockdataautomationCopyBlueprintStage {
			bedrockdataautomation_CopyBlueprintStage(cfg, client)
			return
		}
		if _bedrockdataautomationCreateBlueprint {
			bedrockdataautomation_CreateBlueprint(cfg, client)
			return
		}
		if _bedrockdataautomationCreateBlueprintVersion {
			bedrockdataautomation_CreateBlueprintVersion(cfg, client)
			return
		}
		if _bedrockdataautomationCreateDataAutomationProject {
			bedrockdataautomation_CreateDataAutomationProject(cfg, client)
			return
		}
		if _bedrockdataautomationDeleteBlueprint {
			bedrockdataautomation_DeleteBlueprint(cfg, client)
			return
		}
		if _bedrockdataautomationDeleteDataAutomationProject {
			bedrockdataautomation_DeleteDataAutomationProject(cfg, client)
			return
		}
		if _bedrockdataautomationGetBlueprint {
			bedrockdataautomation_GetBlueprint(cfg, client)
			return
		}
		if _bedrockdataautomationGetBlueprintOptimizationStatus {
			bedrockdataautomation_GetBlueprintOptimizationStatus(cfg, client)
			return
		}
		if _bedrockdataautomationGetDataAutomationProject {
			bedrockdataautomation_GetDataAutomationProject(cfg, client)
			return
		}
		if _bedrockdataautomationInvokeBlueprintOptimizationAsync {
			bedrockdataautomation_InvokeBlueprintOptimizationAsync(cfg, client)
			return
		}
		if _bedrockdataautomationListBlueprints {
			bedrockdataautomation_ListBlueprints(cfg, client)
			return
		}
		if _bedrockdataautomationListDataAutomationProjects {
			bedrockdataautomation_ListDataAutomationProjects(cfg, client)
			return
		}
		if _bedrockdataautomationListTagsForResource {
			bedrockdataautomation_ListTagsForResource(cfg, client)
			return
		}
		if _bedrockdataautomationTagResource {
			bedrockdataautomation_TagResource(cfg, client)
			return
		}
		if _bedrockdataautomationUntagResource {
			bedrockdataautomation_UntagResource(cfg, client)
			return
		}
		if _bedrockdataautomationUpdateBlueprint {
			bedrockdataautomation_UpdateBlueprint(cfg, client)
			return
		}
		if _bedrockdataautomationUpdateDataAutomationProject {
			bedrockdataautomation_UpdateDataAutomationProject(cfg, client)
			return
		}

	},
}

var (
	_bedrockdataautomationCopyBlueprintStage               bool
	_bedrockdataautomationCreateBlueprint                  bool
	_bedrockdataautomationCreateBlueprintVersion           bool
	_bedrockdataautomationCreateDataAutomationProject      bool
	_bedrockdataautomationDeleteBlueprint                  bool
	_bedrockdataautomationDeleteDataAutomationProject      bool
	_bedrockdataautomationGetBlueprint                     bool
	_bedrockdataautomationGetBlueprintOptimizationStatus   bool
	_bedrockdataautomationGetDataAutomationProject         bool
	_bedrockdataautomationInvokeBlueprintOptimizationAsync bool
	_bedrockdataautomationListBlueprints                   bool
	_bedrockdataautomationListDataAutomationProjects       bool
	_bedrockdataautomationListTagsForResource              bool
	_bedrockdataautomationTagResource                      bool
	_bedrockdataautomationUntagResource                    bool
	_bedrockdataautomationUpdateBlueprint                  bool
	_bedrockdataautomationUpdateDataAutomationProject      bool

	_bedrockdataautomationBlueprint                   string
	_bedrockdataautomationBlueprintArn                string
	_bedrockdataautomationBlueprintFilter             string
	_bedrockdataautomationBlueprintName               string
	_bedrockdataautomationBlueprintStage              string
	_bedrockdataautomationBlueprintStageFilter        string
	_bedrockdataautomationBlueprintVersion            string
	_bedrockdataautomationClientToken                 string
	_bedrockdataautomationCustomOutputConfiguration   string
	_bedrockdataautomationDataAutomationProfileArn    string
	_bedrockdataautomationEncryptionConfiguration     string
	_bedrockdataautomationInvocationArn               string
	_bedrockdataautomationMaxResults                  string
	_bedrockdataautomationNextToken                   string
	_bedrockdataautomationOutputConfiguration         string
	_bedrockdataautomationOverrideConfiguration       string
	_bedrockdataautomationProjectArn                  string
	_bedrockdataautomationProjectDescription          string
	_bedrockdataautomationProjectFilter               string
	_bedrockdataautomationProjectName                 string
	_bedrockdataautomationProjectStage                string
	_bedrockdataautomationProjectStageFilter          string
	_bedrockdataautomationProjectType                 string
	_bedrockdataautomationResourceARN                 string
	_bedrockdataautomationResourceOwner               string
	_bedrockdataautomationSamples                     string
	_bedrockdataautomationSchema                      string
	_bedrockdataautomationSourceStage                 string
	_bedrockdataautomationStandardOutputConfiguration string
	_bedrockdataautomationTagKeys                     []string
	_bedrockdataautomationTags                        string
	_bedrockdataautomationTargetStage                 string
	_bedrockdataautomationType                        string
)

// Copies a Blueprint from one stage to another
func bedrockdataautomation_CopyBlueprintStage(cfg aws.Config, client *bedrockdataautomation.Client) {
	input := &bedrockdataautomation.CopyBlueprintStageInput{
		// BlueprintArn: *string, // Required
		// SourceStage: types.BlueprintStage, // Required
		// TargetStage: types.BlueprintStage, // Required
	}

	if len(_bedrockdataautomationBlueprintArn) > 0 {
		input.BlueprintArn = aws.String(_bedrockdataautomationBlueprintArn)
	}
	if len(_bedrockdataautomationSourceStage) > 0 {
		if err := assignInputField(input, "SourceStage", _bedrockdataautomationSourceStage); err != nil {
			log.Errorf("invalid --source-stage: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationTargetStage) > 0 {
		if err := assignInputField(input, "TargetStage", _bedrockdataautomationTargetStage); err != nil {
			log.Errorf("invalid --target-stage: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockdataautomationClientToken)
	}

	if resp, err := client.CopyBlueprintStage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Bedrock Data Automation Blueprint
func bedrockdataautomation_CreateBlueprint(cfg aws.Config, client *bedrockdataautomation.Client) {
	input := &bedrockdataautomation.CreateBlueprintInput{
		// BlueprintName: *string, // Required
		// Schema: *string, // Required
		// Type: types.Type, // Required
	}

	if len(_bedrockdataautomationBlueprintName) > 0 {
		input.BlueprintName = aws.String(_bedrockdataautomationBlueprintName)
	}
	if len(_bedrockdataautomationSchema) > 0 {
		input.Schema = aws.String(_bedrockdataautomationSchema)
	}
	if len(_bedrockdataautomationType) > 0 {
		if err := assignInputField(input, "Type", _bedrockdataautomationType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationBlueprintStage) > 0 {
		if err := assignInputField(input, "BlueprintStage", _bedrockdataautomationBlueprintStage); err != nil {
			log.Errorf("invalid --blueprint-stage: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockdataautomationClientToken)
	}
	if len(_bedrockdataautomationEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _bedrockdataautomationEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockdataautomationTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBlueprint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new version of an existing Amazon Bedrock Data Automation Blueprint
func bedrockdataautomation_CreateBlueprintVersion(cfg aws.Config, client *bedrockdataautomation.Client) {
	input := &bedrockdataautomation.CreateBlueprintVersionInput{
		// BlueprintArn: *string, // Required
	}

	if len(_bedrockdataautomationBlueprintArn) > 0 {
		input.BlueprintArn = aws.String(_bedrockdataautomationBlueprintArn)
	}
	if len(_bedrockdataautomationClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockdataautomationClientToken)
	}

	if resp, err := client.CreateBlueprintVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Bedrock Data Automation Project
func bedrockdataautomation_CreateDataAutomationProject(cfg aws.Config, client *bedrockdataautomation.Client) {
	input := &bedrockdataautomation.CreateDataAutomationProjectInput{
		// ProjectName: *string, // Required
		// StandardOutputConfiguration: *types.StandardOutputConfiguration, // Required
	}

	if len(_bedrockdataautomationProjectName) > 0 {
		input.ProjectName = aws.String(_bedrockdataautomationProjectName)
	}
	if len(_bedrockdataautomationStandardOutputConfiguration) > 0 {
		if err := assignInputField(input, "StandardOutputConfiguration", _bedrockdataautomationStandardOutputConfiguration); err != nil {
			log.Errorf("invalid --standard-output-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockdataautomationClientToken)
	}
	if len(_bedrockdataautomationCustomOutputConfiguration) > 0 {
		if err := assignInputField(input, "CustomOutputConfiguration", _bedrockdataautomationCustomOutputConfiguration); err != nil {
			log.Errorf("invalid --custom-output-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _bedrockdataautomationEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationOverrideConfiguration) > 0 {
		if err := assignInputField(input, "OverrideConfiguration", _bedrockdataautomationOverrideConfiguration); err != nil {
			log.Errorf("invalid --override-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationProjectDescription) > 0 {
		input.ProjectDescription = aws.String(_bedrockdataautomationProjectDescription)
	}
	if len(_bedrockdataautomationProjectStage) > 0 {
		if err := assignInputField(input, "ProjectStage", _bedrockdataautomationProjectStage); err != nil {
			log.Errorf("invalid --project-stage: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationProjectType) > 0 {
		if err := assignInputField(input, "ProjectType", _bedrockdataautomationProjectType); err != nil {
			log.Errorf("invalid --project-type: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockdataautomationTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataAutomationProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing Amazon Bedrock Data Automation Blueprint
func bedrockdataautomation_DeleteBlueprint(cfg aws.Config, client *bedrockdataautomation.Client) {
	input := &bedrockdataautomation.DeleteBlueprintInput{
		// BlueprintArn: *string, // Required
	}

	if len(_bedrockdataautomationBlueprintArn) > 0 {
		input.BlueprintArn = aws.String(_bedrockdataautomationBlueprintArn)
	}
	if len(_bedrockdataautomationBlueprintVersion) > 0 {
		input.BlueprintVersion = aws.String(_bedrockdataautomationBlueprintVersion)
	}

	if resp, err := client.DeleteBlueprint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing Amazon Bedrock Data Automation Project
func bedrockdataautomation_DeleteDataAutomationProject(cfg aws.Config, client *bedrockdataautomation.Client) {
	input := &bedrockdataautomation.DeleteDataAutomationProjectInput{
		// ProjectArn: *string, // Required
	}

	if len(_bedrockdataautomationProjectArn) > 0 {
		input.ProjectArn = aws.String(_bedrockdataautomationProjectArn)
	}

	if resp, err := client.DeleteDataAutomationProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an existing Amazon Bedrock Data Automation Blueprint
func bedrockdataautomation_GetBlueprint(cfg aws.Config, client *bedrockdataautomation.Client) {
	input := &bedrockdataautomation.GetBlueprintInput{
		// BlueprintArn: *string, // Required
	}

	if len(_bedrockdataautomationBlueprintArn) > 0 {
		input.BlueprintArn = aws.String(_bedrockdataautomationBlueprintArn)
	}
	if len(_bedrockdataautomationBlueprintStage) > 0 {
		if err := assignInputField(input, "BlueprintStage", _bedrockdataautomationBlueprintStage); err != nil {
			log.Errorf("invalid --blueprint-stage: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationBlueprintVersion) > 0 {
		input.BlueprintVersion = aws.String(_bedrockdataautomationBlueprintVersion)
	}

	if resp, err := client.GetBlueprint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// API used to get blueprint optimization status.
func bedrockdataautomation_GetBlueprintOptimizationStatus(cfg aws.Config, client *bedrockdataautomation.Client) {
	input := &bedrockdataautomation.GetBlueprintOptimizationStatusInput{
		// InvocationArn: *string, // Required
	}

	if len(_bedrockdataautomationInvocationArn) > 0 {
		input.InvocationArn = aws.String(_bedrockdataautomationInvocationArn)
	}

	if resp, err := client.GetBlueprintOptimizationStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an existing Amazon Bedrock Data Automation Project
func bedrockdataautomation_GetDataAutomationProject(cfg aws.Config, client *bedrockdataautomation.Client) {
	input := &bedrockdataautomation.GetDataAutomationProjectInput{
		// ProjectArn: *string, // Required
	}

	if len(_bedrockdataautomationProjectArn) > 0 {
		input.ProjectArn = aws.String(_bedrockdataautomationProjectArn)
	}
	if len(_bedrockdataautomationProjectStage) > 0 {
		if err := assignInputField(input, "ProjectStage", _bedrockdataautomationProjectStage); err != nil {
			log.Errorf("invalid --project-stage: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetDataAutomationProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Invoke an async job to perform Blueprint Optimization
func bedrockdataautomation_InvokeBlueprintOptimizationAsync(cfg aws.Config, client *bedrockdataautomation.Client) {
	input := &bedrockdataautomation.InvokeBlueprintOptimizationAsyncInput{
		// Blueprint: *types.BlueprintOptimizationObject, // Required
		// DataAutomationProfileArn: *string, // Required
		// OutputConfiguration: *types.BlueprintOptimizationOutputConfiguration, // Required
		// Samples: []types.BlueprintOptimizationSample, // Required
	}

	if len(_bedrockdataautomationBlueprint) > 0 {
		if err := assignInputField(input, "Blueprint", _bedrockdataautomationBlueprint); err != nil {
			log.Errorf("invalid --blueprint: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationDataAutomationProfileArn) > 0 {
		input.DataAutomationProfileArn = aws.String(_bedrockdataautomationDataAutomationProfileArn)
	}
	if len(_bedrockdataautomationOutputConfiguration) > 0 {
		if err := assignInputField(input, "OutputConfiguration", _bedrockdataautomationOutputConfiguration); err != nil {
			log.Errorf("invalid --output-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationSamples) > 0 {
		if err := assignInputField(input, "Samples", _bedrockdataautomationSamples); err != nil {
			log.Errorf("invalid --samples: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _bedrockdataautomationEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockdataautomationTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.InvokeBlueprintOptimizationAsync(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all existing Amazon Bedrock Data Automation Blueprints
func bedrockdataautomation_ListBlueprints(cfg aws.Config, client *bedrockdataautomation.Client) {
	input := &bedrockdataautomation.ListBlueprintsInput{}

	if len(_bedrockdataautomationBlueprintArn) > 0 {
		input.BlueprintArn = aws.String(_bedrockdataautomationBlueprintArn)
	}
	if len(_bedrockdataautomationBlueprintStageFilter) > 0 {
		if err := assignInputField(input, "BlueprintStageFilter", _bedrockdataautomationBlueprintStageFilter); err != nil {
			log.Errorf("invalid --blueprint-stage-filter: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockdataautomationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationNextToken) > 0 {
		input.NextToken = aws.String(_bedrockdataautomationNextToken)
	}
	if len(_bedrockdataautomationProjectFilter) > 0 {
		if err := assignInputField(input, "ProjectFilter", _bedrockdataautomationProjectFilter); err != nil {
			log.Errorf("invalid --project-filter: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationResourceOwner) > 0 {
		if err := assignInputField(input, "ResourceOwner", _bedrockdataautomationResourceOwner); err != nil {
			log.Errorf("invalid --resource-owner: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListBlueprints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockdataautomation.ListBlueprintsOutput
	p := bedrockdataautomation.NewListBlueprintsPaginator(client, input)
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

// Lists all existing Amazon Bedrock Data Automation Projects
func bedrockdataautomation_ListDataAutomationProjects(cfg aws.Config, client *bedrockdataautomation.Client) {
	input := &bedrockdataautomation.ListDataAutomationProjectsInput{}

	if len(_bedrockdataautomationBlueprintFilter) > 0 {
		if err := assignInputField(input, "BlueprintFilter", _bedrockdataautomationBlueprintFilter); err != nil {
			log.Errorf("invalid --blueprint-filter: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockdataautomationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationNextToken) > 0 {
		input.NextToken = aws.String(_bedrockdataautomationNextToken)
	}
	if len(_bedrockdataautomationProjectStageFilter) > 0 {
		if err := assignInputField(input, "ProjectStageFilter", _bedrockdataautomationProjectStageFilter); err != nil {
			log.Errorf("invalid --project-stage-filter: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationResourceOwner) > 0 {
		if err := assignInputField(input, "ResourceOwner", _bedrockdataautomationResourceOwner); err != nil {
			log.Errorf("invalid --resource-owner: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDataAutomationProjects(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockdataautomation.ListDataAutomationProjectsOutput
	p := bedrockdataautomation.NewListDataAutomationProjectsPaginator(client, input)
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

// List tags for an Amazon Bedrock Data Automation resource
func bedrockdataautomation_ListTagsForResource(cfg aws.Config, client *bedrockdataautomation.Client) {
	input := &bedrockdataautomation.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_bedrockdataautomationResourceARN) > 0 {
		input.ResourceARN = aws.String(_bedrockdataautomationResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tag an Amazon Bedrock Data Automation resource
func bedrockdataautomation_TagResource(cfg aws.Config, client *bedrockdataautomation.Client) {
	input := &bedrockdataautomation.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_bedrockdataautomationResourceARN) > 0 {
		input.ResourceARN = aws.String(_bedrockdataautomationResourceARN)
	}
	if len(_bedrockdataautomationTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockdataautomationTags); err != nil {
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

// Untag an Amazon Bedrock Data Automation resource
func bedrockdataautomation_UntagResource(cfg aws.Config, client *bedrockdataautomation.Client) {
	input := &bedrockdataautomation.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_bedrockdataautomationResourceARN) > 0 {
		input.ResourceARN = aws.String(_bedrockdataautomationResourceARN)
	}
	if len(_bedrockdataautomationTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _bedrockdataautomationTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing Amazon Bedrock Data Automation Blueprint
func bedrockdataautomation_UpdateBlueprint(cfg aws.Config, client *bedrockdataautomation.Client) {
	input := &bedrockdataautomation.UpdateBlueprintInput{
		// BlueprintArn: *string, // Required
		// Schema: *string, // Required
	}

	if len(_bedrockdataautomationBlueprintArn) > 0 {
		input.BlueprintArn = aws.String(_bedrockdataautomationBlueprintArn)
	}
	if len(_bedrockdataautomationSchema) > 0 {
		input.Schema = aws.String(_bedrockdataautomationSchema)
	}
	if len(_bedrockdataautomationBlueprintStage) > 0 {
		if err := assignInputField(input, "BlueprintStage", _bedrockdataautomationBlueprintStage); err != nil {
			log.Errorf("invalid --blueprint-stage: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _bedrockdataautomationEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBlueprint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing Amazon Bedrock Data Automation Project
func bedrockdataautomation_UpdateDataAutomationProject(cfg aws.Config, client *bedrockdataautomation.Client) {
	input := &bedrockdataautomation.UpdateDataAutomationProjectInput{
		// ProjectArn: *string, // Required
		// StandardOutputConfiguration: *types.StandardOutputConfiguration, // Required
	}

	if len(_bedrockdataautomationProjectArn) > 0 {
		input.ProjectArn = aws.String(_bedrockdataautomationProjectArn)
	}
	if len(_bedrockdataautomationStandardOutputConfiguration) > 0 {
		if err := assignInputField(input, "StandardOutputConfiguration", _bedrockdataautomationStandardOutputConfiguration); err != nil {
			log.Errorf("invalid --standard-output-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationCustomOutputConfiguration) > 0 {
		if err := assignInputField(input, "CustomOutputConfiguration", _bedrockdataautomationCustomOutputConfiguration); err != nil {
			log.Errorf("invalid --custom-output-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _bedrockdataautomationEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationOverrideConfiguration) > 0 {
		if err := assignInputField(input, "OverrideConfiguration", _bedrockdataautomationOverrideConfiguration); err != nil {
			log.Errorf("invalid --override-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationProjectDescription) > 0 {
		input.ProjectDescription = aws.String(_bedrockdataautomationProjectDescription)
	}
	if len(_bedrockdataautomationProjectStage) > 0 {
		if err := assignInputField(input, "ProjectStage", _bedrockdataautomationProjectStage); err != nil {
			log.Errorf("invalid --project-stage: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDataAutomationProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_bedrockdataautomationCmd)
	_bedrockdataautomationCmd.Flags().SortFlags = false

	_bedrockdataautomationCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_bedrockdataautomationCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_bedrockdataautomationCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationBlueprint, "blueprint", "", "", "Blueprint")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationBlueprintArn, "blueprint-arn", "", "", "Blueprint ARN")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationBlueprintFilter, "blueprint-filter", "", "", "Blueprint Filter")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationBlueprintName, "blueprint-name", "", "", "Blueprint Name")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationBlueprintStage, "blueprint-stage", "", "", "Blueprint Stage")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationBlueprintStageFilter, "blueprint-stage-filter", "", "", "Blueprint Stage Filter")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationBlueprintVersion, "blueprint-version", "", "", "Blueprint Version")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationClientToken, "client-token", "", "", "Client Token")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationCustomOutputConfiguration, "custom-output-configuration", "", "", "Custom Output Configuration")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationDataAutomationProfileArn, "data-automation-profile-arn", "", "", "Data Automation Profile ARN")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationEncryptionConfiguration, "encryption-configuration", "", "", "Encryption Configuration")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationInvocationArn, "invocation-arn", "", "", "Invocation ARN")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationMaxResults, "max-results", "", "", "Max Results")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationNextToken, "next-token", "", "", "Next Token")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationOutputConfiguration, "output-configuration", "", "", "Output Configuration")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationOverrideConfiguration, "override-configuration", "", "", "Override Configuration")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationProjectArn, "project-arn", "", "", "Project ARN")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationProjectDescription, "project-description", "", "", "Project Description")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationProjectFilter, "project-filter", "", "", "Project Filter")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationProjectName, "project-name", "", "", "Project Name")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationProjectStage, "project-stage", "", "", "Project Stage")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationProjectStageFilter, "project-stage-filter", "", "", "Project Stage Filter")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationProjectType, "project-type", "", "", "Project Type")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationResourceARN, "resource-arn", "", "", "Resource ARN")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationResourceOwner, "resource-owner", "", "", "Resource Owner")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationSamples, "samples", "", "", "Samples")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationSchema, "schema", "", "", "Schema")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationSourceStage, "source-stage", "", "", "Source Stage")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationStandardOutputConfiguration, "standard-output-configuration", "", "", "Standard Output Configuration")
	_bedrockdataautomationCmd.Flags().StringSliceVarP(&_bedrockdataautomationTagKeys, "tag-keys", "", nil, "Tag Keys")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationTags, "tags", "", "", "Tags")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationTargetStage, "target-stage", "", "", "Target Stage")
	_bedrockdataautomationCmd.Flags().StringVarP(&_bedrockdataautomationType, "type", "", "", "Type")

	_bedrockdataautomationCmd.Flags().BoolVarP(&_bedrockdataautomationCopyBlueprintStage, "copy-blueprint-stage", "", false, "Copy Blueprint Stage")
	_bedrockdataautomationCmd.Flags().BoolVarP(&_bedrockdataautomationCreateBlueprint, "create-blueprint", "", false, "Create Blueprint")
	_bedrockdataautomationCmd.Flags().BoolVarP(&_bedrockdataautomationCreateBlueprintVersion, "create-blueprint-version", "", false, "Create Blueprint Version")
	_bedrockdataautomationCmd.Flags().BoolVarP(&_bedrockdataautomationCreateDataAutomationProject, "create-data-automation-project", "", false, "Create Data Automation Project")
	_bedrockdataautomationCmd.Flags().BoolVarP(&_bedrockdataautomationDeleteBlueprint, "delete-blueprint", "", false, "Delete Blueprint")
	_bedrockdataautomationCmd.Flags().BoolVarP(&_bedrockdataautomationDeleteDataAutomationProject, "delete-data-automation-project", "", false, "Delete Data Automation Project")
	_bedrockdataautomationCmd.Flags().BoolVarP(&_bedrockdataautomationGetBlueprint, "get-blueprint", "", false, "Get Blueprint")
	_bedrockdataautomationCmd.Flags().BoolVarP(&_bedrockdataautomationGetBlueprintOptimizationStatus, "get-blueprint-optimization-status", "", false, "Get Blueprint Optimization Status")
	_bedrockdataautomationCmd.Flags().BoolVarP(&_bedrockdataautomationGetDataAutomationProject, "get-data-automation-project", "", false, "Get Data Automation Project")
	_bedrockdataautomationCmd.Flags().BoolVarP(&_bedrockdataautomationInvokeBlueprintOptimizationAsync, "invoke-blueprint-optimization-async", "", false, "Invoke Blueprint Optimization Async")
	_bedrockdataautomationCmd.Flags().BoolVarP(&_bedrockdataautomationListBlueprints, "list-blueprints", "", false, "List Blueprints")
	_bedrockdataautomationCmd.Flags().BoolVarP(&_bedrockdataautomationListDataAutomationProjects, "list-data-automation-projects", "", false, "List Data Automation Projects")
	_bedrockdataautomationCmd.Flags().BoolVarP(&_bedrockdataautomationListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_bedrockdataautomationCmd.Flags().BoolVarP(&_bedrockdataautomationTagResource, "tag-resource", "", false, "Tag Resource")
	_bedrockdataautomationCmd.Flags().BoolVarP(&_bedrockdataautomationUntagResource, "untag-resource", "", false, "Untag Resource")
	_bedrockdataautomationCmd.Flags().BoolVarP(&_bedrockdataautomationUpdateBlueprint, "update-blueprint", "", false, "Update Blueprint")
	_bedrockdataautomationCmd.Flags().BoolVarP(&_bedrockdataautomationUpdateDataAutomationProject, "update-data-automation-project", "", false, "Update Data Automation Project")

}
