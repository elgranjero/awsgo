package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/migrationhuborchestrator"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// migrationhuborchestratorCmd represents the migrationhuborchestrator command
var _migrationhuborchestratorCmd = &cobra.Command{
	Use:   "migrationhuborchestrator",
	Short: "AWS migrationhuborchestrator CLI",
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
		client := migrationhuborchestrator.NewFromConfig(cfg)
		if _migrationhuborchestratorCreateTemplate {
			migrationhuborchestrator_CreateTemplate(cfg, client)
			return
		}
		if _migrationhuborchestratorCreateWorkflow {
			migrationhuborchestrator_CreateWorkflow(cfg, client)
			return
		}
		if _migrationhuborchestratorCreateWorkflowStep {
			migrationhuborchestrator_CreateWorkflowStep(cfg, client)
			return
		}
		if _migrationhuborchestratorCreateWorkflowStepGroup {
			migrationhuborchestrator_CreateWorkflowStepGroup(cfg, client)
			return
		}
		if _migrationhuborchestratorDeleteTemplate {
			migrationhuborchestrator_DeleteTemplate(cfg, client)
			return
		}
		if _migrationhuborchestratorDeleteWorkflow {
			migrationhuborchestrator_DeleteWorkflow(cfg, client)
			return
		}
		if _migrationhuborchestratorDeleteWorkflowStep {
			migrationhuborchestrator_DeleteWorkflowStep(cfg, client)
			return
		}
		if _migrationhuborchestratorDeleteWorkflowStepGroup {
			migrationhuborchestrator_DeleteWorkflowStepGroup(cfg, client)
			return
		}
		if _migrationhuborchestratorGetTemplate {
			migrationhuborchestrator_GetTemplate(cfg, client)
			return
		}
		if _migrationhuborchestratorGetTemplateStep {
			migrationhuborchestrator_GetTemplateStep(cfg, client)
			return
		}
		if _migrationhuborchestratorGetTemplateStepGroup {
			migrationhuborchestrator_GetTemplateStepGroup(cfg, client)
			return
		}
		if _migrationhuborchestratorGetWorkflow {
			migrationhuborchestrator_GetWorkflow(cfg, client)
			return
		}
		if _migrationhuborchestratorGetWorkflowStep {
			migrationhuborchestrator_GetWorkflowStep(cfg, client)
			return
		}
		if _migrationhuborchestratorGetWorkflowStepGroup {
			migrationhuborchestrator_GetWorkflowStepGroup(cfg, client)
			return
		}
		if _migrationhuborchestratorListPlugins {
			migrationhuborchestrator_ListPlugins(cfg, client)
			return
		}
		if _migrationhuborchestratorListTagsForResource {
			migrationhuborchestrator_ListTagsForResource(cfg, client)
			return
		}
		if _migrationhuborchestratorListTemplateStepGroups {
			migrationhuborchestrator_ListTemplateStepGroups(cfg, client)
			return
		}
		if _migrationhuborchestratorListTemplateSteps {
			migrationhuborchestrator_ListTemplateSteps(cfg, client)
			return
		}
		if _migrationhuborchestratorListTemplates {
			migrationhuborchestrator_ListTemplates(cfg, client)
			return
		}
		if _migrationhuborchestratorListWorkflowStepGroups {
			migrationhuborchestrator_ListWorkflowStepGroups(cfg, client)
			return
		}
		if _migrationhuborchestratorListWorkflowSteps {
			migrationhuborchestrator_ListWorkflowSteps(cfg, client)
			return
		}
		if _migrationhuborchestratorListWorkflows {
			migrationhuborchestrator_ListWorkflows(cfg, client)
			return
		}
		if _migrationhuborchestratorRetryWorkflowStep {
			migrationhuborchestrator_RetryWorkflowStep(cfg, client)
			return
		}
		if _migrationhuborchestratorStartWorkflow {
			migrationhuborchestrator_StartWorkflow(cfg, client)
			return
		}
		if _migrationhuborchestratorStopWorkflow {
			migrationhuborchestrator_StopWorkflow(cfg, client)
			return
		}
		if _migrationhuborchestratorTagResource {
			migrationhuborchestrator_TagResource(cfg, client)
			return
		}
		if _migrationhuborchestratorUntagResource {
			migrationhuborchestrator_UntagResource(cfg, client)
			return
		}
		if _migrationhuborchestratorUpdateTemplate {
			migrationhuborchestrator_UpdateTemplate(cfg, client)
			return
		}
		if _migrationhuborchestratorUpdateWorkflow {
			migrationhuborchestrator_UpdateWorkflow(cfg, client)
			return
		}
		if _migrationhuborchestratorUpdateWorkflowStep {
			migrationhuborchestrator_UpdateWorkflowStep(cfg, client)
			return
		}
		if _migrationhuborchestratorUpdateWorkflowStepGroup {
			migrationhuborchestrator_UpdateWorkflowStepGroup(cfg, client)
			return
		}

	},
}

var (
	_migrationhuborchestratorCreateTemplate          bool
	_migrationhuborchestratorCreateWorkflow          bool
	_migrationhuborchestratorCreateWorkflowStep      bool
	_migrationhuborchestratorCreateWorkflowStepGroup bool
	_migrationhuborchestratorDeleteTemplate          bool
	_migrationhuborchestratorDeleteWorkflow          bool
	_migrationhuborchestratorDeleteWorkflowStep      bool
	_migrationhuborchestratorDeleteWorkflowStepGroup bool
	_migrationhuborchestratorGetTemplate             bool
	_migrationhuborchestratorGetTemplateStep         bool
	_migrationhuborchestratorGetTemplateStepGroup    bool
	_migrationhuborchestratorGetWorkflow             bool
	_migrationhuborchestratorGetWorkflowStep         bool
	_migrationhuborchestratorGetWorkflowStepGroup    bool
	_migrationhuborchestratorListPlugins             bool
	_migrationhuborchestratorListTagsForResource     bool
	_migrationhuborchestratorListTemplateStepGroups  bool
	_migrationhuborchestratorListTemplateSteps       bool
	_migrationhuborchestratorListTemplates           bool
	_migrationhuborchestratorListWorkflowStepGroups  bool
	_migrationhuborchestratorListWorkflowSteps       bool
	_migrationhuborchestratorListWorkflows           bool
	_migrationhuborchestratorRetryWorkflowStep       bool
	_migrationhuborchestratorStartWorkflow           bool
	_migrationhuborchestratorStopWorkflow            bool
	_migrationhuborchestratorTagResource             bool
	_migrationhuborchestratorUntagResource           bool
	_migrationhuborchestratorUpdateTemplate          bool
	_migrationhuborchestratorUpdateWorkflow          bool
	_migrationhuborchestratorUpdateWorkflowStep      bool
	_migrationhuborchestratorUpdateWorkflowStepGroup bool

	_migrationhuborchestratorAdsApplicationConfigurationName     string
	_migrationhuborchestratorApplicationConfigurationId          string
	_migrationhuborchestratorClientToken                         string
	_migrationhuborchestratorDescription                         string
	_migrationhuborchestratorId                                  string
	_migrationhuborchestratorInputParameters                     string
	_migrationhuborchestratorMaxResults                          string
	_migrationhuborchestratorName                                string
	_migrationhuborchestratorNext                                []string
	_migrationhuborchestratorNextToken                           string
	_migrationhuborchestratorOutputs                             string
	_migrationhuborchestratorPrevious                            []string
	_migrationhuborchestratorResourceArn                         string
	_migrationhuborchestratorStatus                              string
	_migrationhuborchestratorStepActionType                      string
	_migrationhuborchestratorStepGroupId                         string
	_migrationhuborchestratorStepTarget                          []string
	_migrationhuborchestratorStepTargets                         []string
	_migrationhuborchestratorTagKeys                             []string
	_migrationhuborchestratorTags                                string
	_migrationhuborchestratorTemplateDescription                 string
	_migrationhuborchestratorTemplateId                          string
	_migrationhuborchestratorTemplateName                        string
	_migrationhuborchestratorTemplateSource                      string
	_migrationhuborchestratorWorkflowId                          string
	_migrationhuborchestratorWorkflowStepAutomationConfiguration string
)

// Creates a migration workflow template.
func migrationhuborchestrator_CreateTemplate(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.CreateTemplateInput{
		// TemplateName: *string, // Required
		// TemplateSource: types.TemplateSource, // Required
	}

	if len(_migrationhuborchestratorTemplateName) > 0 {
		input.TemplateName = aws.String(_migrationhuborchestratorTemplateName)
	}
	if len(_migrationhuborchestratorTemplateSource) > 0 {
		if err := assignInputField(input, "TemplateSource", _migrationhuborchestratorTemplateSource); err != nil {
			log.Errorf("invalid --template-source: %s", err.Error())
			return
		}
	}
	if len(_migrationhuborchestratorClientToken) > 0 {
		input.ClientToken = aws.String(_migrationhuborchestratorClientToken)
	}
	if len(_migrationhuborchestratorTags) > 0 {
		if err := assignInputField(input, "Tags", _migrationhuborchestratorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_migrationhuborchestratorTemplateDescription) > 0 {
		input.TemplateDescription = aws.String(_migrationhuborchestratorTemplateDescription)
	}

	if resp, err := client.CreateTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a workflow to orchestrate your migrations.
func migrationhuborchestrator_CreateWorkflow(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.CreateWorkflowInput{
		// InputParameters: map[string]types.StepInput, // Required
		// Name: *string, // Required
		// TemplateId: *string, // Required
	}

	if len(_migrationhuborchestratorInputParameters) > 0 {
		if err := assignInputField(input, "InputParameters", _migrationhuborchestratorInputParameters); err != nil {
			log.Errorf("invalid --input-parameters: %s", err.Error())
			return
		}
	}
	if len(_migrationhuborchestratorName) > 0 {
		input.Name = aws.String(_migrationhuborchestratorName)
	}
	if len(_migrationhuborchestratorTemplateId) > 0 {
		input.TemplateId = aws.String(_migrationhuborchestratorTemplateId)
	}
	if len(_migrationhuborchestratorApplicationConfigurationId) > 0 {
		input.ApplicationConfigurationId = aws.String(_migrationhuborchestratorApplicationConfigurationId)
	}
	if len(_migrationhuborchestratorDescription) > 0 {
		input.Description = aws.String(_migrationhuborchestratorDescription)
	}
	if len(_migrationhuborchestratorStepTargets) > 0 {
		input.StepTargets = append([]string(nil), _migrationhuborchestratorStepTargets...)
	}
	if len(_migrationhuborchestratorTags) > 0 {
		if err := assignInputField(input, "Tags", _migrationhuborchestratorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a step in the migration workflow.
func migrationhuborchestrator_CreateWorkflowStep(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.CreateWorkflowStepInput{
		// Name: *string, // Required
		// StepActionType: types.StepActionType, // Required
		// StepGroupId: *string, // Required
		// WorkflowId: *string, // Required
	}

	if len(_migrationhuborchestratorName) > 0 {
		input.Name = aws.String(_migrationhuborchestratorName)
	}
	if len(_migrationhuborchestratorStepActionType) > 0 {
		if err := assignInputField(input, "StepActionType", _migrationhuborchestratorStepActionType); err != nil {
			log.Errorf("invalid --step-action-type: %s", err.Error())
			return
		}
	}
	if len(_migrationhuborchestratorStepGroupId) > 0 {
		input.StepGroupId = aws.String(_migrationhuborchestratorStepGroupId)
	}
	if len(_migrationhuborchestratorWorkflowId) > 0 {
		input.WorkflowId = aws.String(_migrationhuborchestratorWorkflowId)
	}
	if len(_migrationhuborchestratorDescription) > 0 {
		input.Description = aws.String(_migrationhuborchestratorDescription)
	}
	if len(_migrationhuborchestratorNext) > 0 {
		input.Next = append([]string(nil), _migrationhuborchestratorNext...)
	}
	if len(_migrationhuborchestratorOutputs) > 0 {
		if err := assignInputField(input, "Outputs", _migrationhuborchestratorOutputs); err != nil {
			log.Errorf("invalid --outputs: %s", err.Error())
			return
		}
	}
	if len(_migrationhuborchestratorPrevious) > 0 {
		input.Previous = append([]string(nil), _migrationhuborchestratorPrevious...)
	}
	if len(_migrationhuborchestratorStepTarget) > 0 {
		input.StepTarget = append([]string(nil), _migrationhuborchestratorStepTarget...)
	}
	if len(_migrationhuborchestratorWorkflowStepAutomationConfiguration) > 0 {
		if err := assignInputField(input, "WorkflowStepAutomationConfiguration", _migrationhuborchestratorWorkflowStepAutomationConfiguration); err != nil {
			log.Errorf("invalid --workflow-step-automation-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWorkflowStep(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a step group in a migration workflow.
func migrationhuborchestrator_CreateWorkflowStepGroup(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.CreateWorkflowStepGroupInput{
		// Name: *string, // Required
		// WorkflowId: *string, // Required
	}

	if len(_migrationhuborchestratorName) > 0 {
		input.Name = aws.String(_migrationhuborchestratorName)
	}
	if len(_migrationhuborchestratorWorkflowId) > 0 {
		input.WorkflowId = aws.String(_migrationhuborchestratorWorkflowId)
	}
	if len(_migrationhuborchestratorDescription) > 0 {
		input.Description = aws.String(_migrationhuborchestratorDescription)
	}
	if len(_migrationhuborchestratorNext) > 0 {
		input.Next = append([]string(nil), _migrationhuborchestratorNext...)
	}
	if len(_migrationhuborchestratorPrevious) > 0 {
		input.Previous = append([]string(nil), _migrationhuborchestratorPrevious...)
	}

	if resp, err := client.CreateWorkflowStepGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a migration workflow template.
func migrationhuborchestrator_DeleteTemplate(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.DeleteTemplateInput{
		// Id: *string, // Required
	}

	if len(_migrationhuborchestratorId) > 0 {
		input.Id = aws.String(_migrationhuborchestratorId)
	}

	if resp, err := client.DeleteTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a migration workflow. You must pause a running workflow in Migration Hub
// Orchestrator console to delete it.
func migrationhuborchestrator_DeleteWorkflow(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.DeleteWorkflowInput{
		// Id: *string, // Required
	}

	if len(_migrationhuborchestratorId) > 0 {
		input.Id = aws.String(_migrationhuborchestratorId)
	}

	if resp, err := client.DeleteWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a step in a migration workflow. Pause the workflow to delete a running
// step.
func migrationhuborchestrator_DeleteWorkflowStep(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.DeleteWorkflowStepInput{
		// Id: *string, // Required
		// StepGroupId: *string, // Required
		// WorkflowId: *string, // Required
	}

	if len(_migrationhuborchestratorId) > 0 {
		input.Id = aws.String(_migrationhuborchestratorId)
	}
	if len(_migrationhuborchestratorStepGroupId) > 0 {
		input.StepGroupId = aws.String(_migrationhuborchestratorStepGroupId)
	}
	if len(_migrationhuborchestratorWorkflowId) > 0 {
		input.WorkflowId = aws.String(_migrationhuborchestratorWorkflowId)
	}

	if resp, err := client.DeleteWorkflowStep(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a step group in a migration workflow.
func migrationhuborchestrator_DeleteWorkflowStepGroup(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.DeleteWorkflowStepGroupInput{
		// Id: *string, // Required
		// WorkflowId: *string, // Required
	}

	if len(_migrationhuborchestratorId) > 0 {
		input.Id = aws.String(_migrationhuborchestratorId)
	}
	if len(_migrationhuborchestratorWorkflowId) > 0 {
		input.WorkflowId = aws.String(_migrationhuborchestratorWorkflowId)
	}

	if resp, err := client.DeleteWorkflowStepGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the template you want to use for creating a migration workflow.
func migrationhuborchestrator_GetTemplate(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.GetTemplateInput{
		// Id: *string, // Required
	}

	if len(_migrationhuborchestratorId) > 0 {
		input.Id = aws.String(_migrationhuborchestratorId)
	}

	if resp, err := client.GetTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get a specific step in a template.
func migrationhuborchestrator_GetTemplateStep(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.GetTemplateStepInput{
		// Id: *string, // Required
		// StepGroupId: *string, // Required
		// TemplateId: *string, // Required
	}

	if len(_migrationhuborchestratorId) > 0 {
		input.Id = aws.String(_migrationhuborchestratorId)
	}
	if len(_migrationhuborchestratorStepGroupId) > 0 {
		input.StepGroupId = aws.String(_migrationhuborchestratorStepGroupId)
	}
	if len(_migrationhuborchestratorTemplateId) > 0 {
		input.TemplateId = aws.String(_migrationhuborchestratorTemplateId)
	}

	if resp, err := client.GetTemplateStep(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get a step group in a template.
func migrationhuborchestrator_GetTemplateStepGroup(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.GetTemplateStepGroupInput{
		// Id: *string, // Required
		// TemplateId: *string, // Required
	}

	if len(_migrationhuborchestratorId) > 0 {
		input.Id = aws.String(_migrationhuborchestratorId)
	}
	if len(_migrationhuborchestratorTemplateId) > 0 {
		input.TemplateId = aws.String(_migrationhuborchestratorTemplateId)
	}

	if resp, err := client.GetTemplateStepGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get migration workflow.
func migrationhuborchestrator_GetWorkflow(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.GetWorkflowInput{
		// Id: *string, // Required
	}

	if len(_migrationhuborchestratorId) > 0 {
		input.Id = aws.String(_migrationhuborchestratorId)
	}

	if resp, err := client.GetWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get a step in the migration workflow.
func migrationhuborchestrator_GetWorkflowStep(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.GetWorkflowStepInput{
		// Id: *string, // Required
		// StepGroupId: *string, // Required
		// WorkflowId: *string, // Required
	}

	if len(_migrationhuborchestratorId) > 0 {
		input.Id = aws.String(_migrationhuborchestratorId)
	}
	if len(_migrationhuborchestratorStepGroupId) > 0 {
		input.StepGroupId = aws.String(_migrationhuborchestratorStepGroupId)
	}
	if len(_migrationhuborchestratorWorkflowId) > 0 {
		input.WorkflowId = aws.String(_migrationhuborchestratorWorkflowId)
	}

	if resp, err := client.GetWorkflowStep(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the step group of a migration workflow.
func migrationhuborchestrator_GetWorkflowStepGroup(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.GetWorkflowStepGroupInput{
		// Id: *string, // Required
		// WorkflowId: *string, // Required
	}

	if len(_migrationhuborchestratorId) > 0 {
		input.Id = aws.String(_migrationhuborchestratorId)
	}
	if len(_migrationhuborchestratorWorkflowId) > 0 {
		input.WorkflowId = aws.String(_migrationhuborchestratorWorkflowId)
	}

	if resp, err := client.GetWorkflowStepGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List AWS Migration Hub Orchestrator plugins.
func migrationhuborchestrator_ListPlugins(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.ListPluginsInput{}

	if len(_migrationhuborchestratorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhuborchestratorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhuborchestratorNextToken) > 0 {
		input.NextToken = aws.String(_migrationhuborchestratorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPlugins(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhuborchestrator.ListPluginsOutput
	p := migrationhuborchestrator.NewListPluginsPaginator(client, input)
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

// List the tags added to a resource.
func migrationhuborchestrator_ListTagsForResource(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_migrationhuborchestratorResourceArn) > 0 {
		input.ResourceArn = aws.String(_migrationhuborchestratorResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List the step groups in a template.
func migrationhuborchestrator_ListTemplateStepGroups(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.ListTemplateStepGroupsInput{
		// TemplateId: *string, // Required
	}

	if len(_migrationhuborchestratorTemplateId) > 0 {
		input.TemplateId = aws.String(_migrationhuborchestratorTemplateId)
	}
	if len(_migrationhuborchestratorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhuborchestratorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhuborchestratorNextToken) > 0 {
		input.NextToken = aws.String(_migrationhuborchestratorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTemplateStepGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhuborchestrator.ListTemplateStepGroupsOutput
	p := migrationhuborchestrator.NewListTemplateStepGroupsPaginator(client, input)
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

// List the steps in a template.
func migrationhuborchestrator_ListTemplateSteps(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.ListTemplateStepsInput{
		// StepGroupId: *string, // Required
		// TemplateId: *string, // Required
	}

	if len(_migrationhuborchestratorStepGroupId) > 0 {
		input.StepGroupId = aws.String(_migrationhuborchestratorStepGroupId)
	}
	if len(_migrationhuborchestratorTemplateId) > 0 {
		input.TemplateId = aws.String(_migrationhuborchestratorTemplateId)
	}
	if len(_migrationhuborchestratorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhuborchestratorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhuborchestratorNextToken) > 0 {
		input.NextToken = aws.String(_migrationhuborchestratorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTemplateSteps(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhuborchestrator.ListTemplateStepsOutput
	p := migrationhuborchestrator.NewListTemplateStepsPaginator(client, input)
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

// List the templates available in Migration Hub Orchestrator to create a
// migration workflow.
func migrationhuborchestrator_ListTemplates(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.ListTemplatesInput{}

	if len(_migrationhuborchestratorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhuborchestratorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhuborchestratorName) > 0 {
		input.Name = aws.String(_migrationhuborchestratorName)
	}
	if len(_migrationhuborchestratorNextToken) > 0 {
		input.NextToken = aws.String(_migrationhuborchestratorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhuborchestrator.ListTemplatesOutput
	p := migrationhuborchestrator.NewListTemplatesPaginator(client, input)
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

// List the step groups in a migration workflow.
func migrationhuborchestrator_ListWorkflowStepGroups(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.ListWorkflowStepGroupsInput{
		// WorkflowId: *string, // Required
	}

	if len(_migrationhuborchestratorWorkflowId) > 0 {
		input.WorkflowId = aws.String(_migrationhuborchestratorWorkflowId)
	}
	if len(_migrationhuborchestratorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhuborchestratorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhuborchestratorNextToken) > 0 {
		input.NextToken = aws.String(_migrationhuborchestratorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkflowStepGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhuborchestrator.ListWorkflowStepGroupsOutput
	p := migrationhuborchestrator.NewListWorkflowStepGroupsPaginator(client, input)
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

// List the steps in a workflow.
func migrationhuborchestrator_ListWorkflowSteps(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.ListWorkflowStepsInput{
		// StepGroupId: *string, // Required
		// WorkflowId: *string, // Required
	}

	if len(_migrationhuborchestratorStepGroupId) > 0 {
		input.StepGroupId = aws.String(_migrationhuborchestratorStepGroupId)
	}
	if len(_migrationhuborchestratorWorkflowId) > 0 {
		input.WorkflowId = aws.String(_migrationhuborchestratorWorkflowId)
	}
	if len(_migrationhuborchestratorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhuborchestratorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhuborchestratorNextToken) > 0 {
		input.NextToken = aws.String(_migrationhuborchestratorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkflowSteps(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhuborchestrator.ListWorkflowStepsOutput
	p := migrationhuborchestrator.NewListWorkflowStepsPaginator(client, input)
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

// List the migration workflows.
func migrationhuborchestrator_ListWorkflows(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.ListWorkflowsInput{}

	if len(_migrationhuborchestratorAdsApplicationConfigurationName) > 0 {
		input.AdsApplicationConfigurationName = aws.String(_migrationhuborchestratorAdsApplicationConfigurationName)
	}
	if len(_migrationhuborchestratorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhuborchestratorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhuborchestratorName) > 0 {
		input.Name = aws.String(_migrationhuborchestratorName)
	}
	if len(_migrationhuborchestratorNextToken) > 0 {
		input.NextToken = aws.String(_migrationhuborchestratorNextToken)
	}
	if len(_migrationhuborchestratorStatus) > 0 {
		if err := assignInputField(input, "Status", _migrationhuborchestratorStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_migrationhuborchestratorTemplateId) > 0 {
		input.TemplateId = aws.String(_migrationhuborchestratorTemplateId)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkflows(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhuborchestrator.ListWorkflowsOutput
	p := migrationhuborchestrator.NewListWorkflowsPaginator(client, input)
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

// Retry a failed step in a migration workflow.
func migrationhuborchestrator_RetryWorkflowStep(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.RetryWorkflowStepInput{
		// Id: *string, // Required
		// StepGroupId: *string, // Required
		// WorkflowId: *string, // Required
	}

	if len(_migrationhuborchestratorId) > 0 {
		input.Id = aws.String(_migrationhuborchestratorId)
	}
	if len(_migrationhuborchestratorStepGroupId) > 0 {
		input.StepGroupId = aws.String(_migrationhuborchestratorStepGroupId)
	}
	if len(_migrationhuborchestratorWorkflowId) > 0 {
		input.WorkflowId = aws.String(_migrationhuborchestratorWorkflowId)
	}

	if resp, err := client.RetryWorkflowStep(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Start a migration workflow.
func migrationhuborchestrator_StartWorkflow(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.StartWorkflowInput{
		// Id: *string, // Required
	}

	if len(_migrationhuborchestratorId) > 0 {
		input.Id = aws.String(_migrationhuborchestratorId)
	}

	if resp, err := client.StartWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stop an ongoing migration workflow.
func migrationhuborchestrator_StopWorkflow(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.StopWorkflowInput{
		// Id: *string, // Required
	}

	if len(_migrationhuborchestratorId) > 0 {
		input.Id = aws.String(_migrationhuborchestratorId)
	}

	if resp, err := client.StopWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tag a resource by specifying its Amazon Resource Name (ARN).
func migrationhuborchestrator_TagResource(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_migrationhuborchestratorResourceArn) > 0 {
		input.ResourceArn = aws.String(_migrationhuborchestratorResourceArn)
	}
	if len(_migrationhuborchestratorTags) > 0 {
		if err := assignInputField(input, "Tags", _migrationhuborchestratorTags); err != nil {
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

// Deletes the tags for a resource.
func migrationhuborchestrator_UntagResource(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_migrationhuborchestratorResourceArn) > 0 {
		input.ResourceArn = aws.String(_migrationhuborchestratorResourceArn)
	}
	if len(_migrationhuborchestratorTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _migrationhuborchestratorTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a migration workflow template.
func migrationhuborchestrator_UpdateTemplate(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.UpdateTemplateInput{
		// Id: *string, // Required
	}

	if len(_migrationhuborchestratorId) > 0 {
		input.Id = aws.String(_migrationhuborchestratorId)
	}
	if len(_migrationhuborchestratorClientToken) > 0 {
		input.ClientToken = aws.String(_migrationhuborchestratorClientToken)
	}
	if len(_migrationhuborchestratorTemplateDescription) > 0 {
		input.TemplateDescription = aws.String(_migrationhuborchestratorTemplateDescription)
	}
	if len(_migrationhuborchestratorTemplateName) > 0 {
		input.TemplateName = aws.String(_migrationhuborchestratorTemplateName)
	}

	if resp, err := client.UpdateTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a migration workflow.
func migrationhuborchestrator_UpdateWorkflow(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.UpdateWorkflowInput{
		// Id: *string, // Required
	}

	if len(_migrationhuborchestratorId) > 0 {
		input.Id = aws.String(_migrationhuborchestratorId)
	}
	if len(_migrationhuborchestratorDescription) > 0 {
		input.Description = aws.String(_migrationhuborchestratorDescription)
	}
	if len(_migrationhuborchestratorInputParameters) > 0 {
		if err := assignInputField(input, "InputParameters", _migrationhuborchestratorInputParameters); err != nil {
			log.Errorf("invalid --input-parameters: %s", err.Error())
			return
		}
	}
	if len(_migrationhuborchestratorName) > 0 {
		input.Name = aws.String(_migrationhuborchestratorName)
	}
	if len(_migrationhuborchestratorStepTargets) > 0 {
		input.StepTargets = append([]string(nil), _migrationhuborchestratorStepTargets...)
	}

	if resp, err := client.UpdateWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a step in a migration workflow.
func migrationhuborchestrator_UpdateWorkflowStep(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.UpdateWorkflowStepInput{
		// Id: *string, // Required
		// StepGroupId: *string, // Required
		// WorkflowId: *string, // Required
	}

	if len(_migrationhuborchestratorId) > 0 {
		input.Id = aws.String(_migrationhuborchestratorId)
	}
	if len(_migrationhuborchestratorStepGroupId) > 0 {
		input.StepGroupId = aws.String(_migrationhuborchestratorStepGroupId)
	}
	if len(_migrationhuborchestratorWorkflowId) > 0 {
		input.WorkflowId = aws.String(_migrationhuborchestratorWorkflowId)
	}
	if len(_migrationhuborchestratorDescription) > 0 {
		input.Description = aws.String(_migrationhuborchestratorDescription)
	}
	if len(_migrationhuborchestratorName) > 0 {
		input.Name = aws.String(_migrationhuborchestratorName)
	}
	if len(_migrationhuborchestratorNext) > 0 {
		input.Next = append([]string(nil), _migrationhuborchestratorNext...)
	}
	if len(_migrationhuborchestratorOutputs) > 0 {
		if err := assignInputField(input, "Outputs", _migrationhuborchestratorOutputs); err != nil {
			log.Errorf("invalid --outputs: %s", err.Error())
			return
		}
	}
	if len(_migrationhuborchestratorPrevious) > 0 {
		input.Previous = append([]string(nil), _migrationhuborchestratorPrevious...)
	}
	if len(_migrationhuborchestratorStatus) > 0 {
		if err := assignInputField(input, "Status", _migrationhuborchestratorStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_migrationhuborchestratorStepActionType) > 0 {
		if err := assignInputField(input, "StepActionType", _migrationhuborchestratorStepActionType); err != nil {
			log.Errorf("invalid --step-action-type: %s", err.Error())
			return
		}
	}
	if len(_migrationhuborchestratorStepTarget) > 0 {
		input.StepTarget = append([]string(nil), _migrationhuborchestratorStepTarget...)
	}
	if len(_migrationhuborchestratorWorkflowStepAutomationConfiguration) > 0 {
		if err := assignInputField(input, "WorkflowStepAutomationConfiguration", _migrationhuborchestratorWorkflowStepAutomationConfiguration); err != nil {
			log.Errorf("invalid --workflow-step-automation-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateWorkflowStep(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the step group in a migration workflow.
func migrationhuborchestrator_UpdateWorkflowStepGroup(cfg aws.Config, client *migrationhuborchestrator.Client) {
	input := &migrationhuborchestrator.UpdateWorkflowStepGroupInput{
		// Id: *string, // Required
		// WorkflowId: *string, // Required
	}

	if len(_migrationhuborchestratorId) > 0 {
		input.Id = aws.String(_migrationhuborchestratorId)
	}
	if len(_migrationhuborchestratorWorkflowId) > 0 {
		input.WorkflowId = aws.String(_migrationhuborchestratorWorkflowId)
	}
	if len(_migrationhuborchestratorDescription) > 0 {
		input.Description = aws.String(_migrationhuborchestratorDescription)
	}
	if len(_migrationhuborchestratorName) > 0 {
		input.Name = aws.String(_migrationhuborchestratorName)
	}
	if len(_migrationhuborchestratorNext) > 0 {
		input.Next = append([]string(nil), _migrationhuborchestratorNext...)
	}
	if len(_migrationhuborchestratorPrevious) > 0 {
		input.Previous = append([]string(nil), _migrationhuborchestratorPrevious...)
	}

	if resp, err := client.UpdateWorkflowStepGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_migrationhuborchestratorCmd)
	_migrationhuborchestratorCmd.Flags().SortFlags = false

	_migrationhuborchestratorCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_migrationhuborchestratorCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_migrationhuborchestratorCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_migrationhuborchestratorCmd.Flags().StringVarP(&_migrationhuborchestratorAdsApplicationConfigurationName, "ads-application-configuration-name", "", "", "Ads Application Configuration Name")
	_migrationhuborchestratorCmd.Flags().StringVarP(&_migrationhuborchestratorApplicationConfigurationId, "application-configuration-id", "", "", "Application Configuration ID")
	_migrationhuborchestratorCmd.Flags().StringVarP(&_migrationhuborchestratorClientToken, "client-token", "", "", "Client Token")
	_migrationhuborchestratorCmd.Flags().StringVarP(&_migrationhuborchestratorDescription, "description", "", "", "Description")
	_migrationhuborchestratorCmd.Flags().StringVarP(&_migrationhuborchestratorId, "id", "", "", "ID")
	_migrationhuborchestratorCmd.Flags().StringVarP(&_migrationhuborchestratorInputParameters, "input-parameters", "", "", "Input Parameters")
	_migrationhuborchestratorCmd.Flags().StringVarP(&_migrationhuborchestratorMaxResults, "max-results", "", "", "Max Results")
	_migrationhuborchestratorCmd.Flags().StringVarP(&_migrationhuborchestratorName, "name", "", "", "Name")
	_migrationhuborchestratorCmd.Flags().StringSliceVarP(&_migrationhuborchestratorNext, "next", "", nil, "Next")
	_migrationhuborchestratorCmd.Flags().StringVarP(&_migrationhuborchestratorNextToken, "next-token", "", "", "Next Token")
	_migrationhuborchestratorCmd.Flags().StringVarP(&_migrationhuborchestratorOutputs, "outputs", "", "", "Outputs")
	_migrationhuborchestratorCmd.Flags().StringSliceVarP(&_migrationhuborchestratorPrevious, "previous", "", nil, "Previous")
	_migrationhuborchestratorCmd.Flags().StringVarP(&_migrationhuborchestratorResourceArn, "resource-arn", "", "", "Resource ARN")
	_migrationhuborchestratorCmd.Flags().StringVarP(&_migrationhuborchestratorStatus, "status", "", "", "Status")
	_migrationhuborchestratorCmd.Flags().StringVarP(&_migrationhuborchestratorStepActionType, "step-action-type", "", "", "Step Action Type")
	_migrationhuborchestratorCmd.Flags().StringVarP(&_migrationhuborchestratorStepGroupId, "step-group-id", "", "", "Step Group ID")
	_migrationhuborchestratorCmd.Flags().StringSliceVarP(&_migrationhuborchestratorStepTarget, "step-target", "", nil, "Step Target")
	_migrationhuborchestratorCmd.Flags().StringSliceVarP(&_migrationhuborchestratorStepTargets, "step-targets", "", nil, "Step Targets")
	_migrationhuborchestratorCmd.Flags().StringSliceVarP(&_migrationhuborchestratorTagKeys, "tag-keys", "", nil, "Tag Keys")
	_migrationhuborchestratorCmd.Flags().StringVarP(&_migrationhuborchestratorTags, "tags", "", "", "Tags")
	_migrationhuborchestratorCmd.Flags().StringVarP(&_migrationhuborchestratorTemplateDescription, "template-description", "", "", "Template Description")
	_migrationhuborchestratorCmd.Flags().StringVarP(&_migrationhuborchestratorTemplateId, "template-id", "", "", "Template ID")
	_migrationhuborchestratorCmd.Flags().StringVarP(&_migrationhuborchestratorTemplateName, "template-name", "", "", "Template Name")
	_migrationhuborchestratorCmd.Flags().StringVarP(&_migrationhuborchestratorTemplateSource, "template-source", "", "", "Template Source")
	_migrationhuborchestratorCmd.Flags().StringVarP(&_migrationhuborchestratorWorkflowId, "workflow-id", "", "", "Workflow ID")
	_migrationhuborchestratorCmd.Flags().StringVarP(&_migrationhuborchestratorWorkflowStepAutomationConfiguration, "workflow-step-automation-configuration", "", "", "Workflow Step Automation Configuration")

	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorCreateTemplate, "create-template", "", false, "Create Template")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorCreateWorkflow, "create-workflow", "", false, "Create Workflow")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorCreateWorkflowStep, "create-workflow-step", "", false, "Create Workflow Step")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorCreateWorkflowStepGroup, "create-workflow-step-group", "", false, "Create Workflow Step Group")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorDeleteTemplate, "delete-template", "", false, "Delete Template")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorDeleteWorkflow, "delete-workflow", "", false, "Delete Workflow")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorDeleteWorkflowStep, "delete-workflow-step", "", false, "Delete Workflow Step")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorDeleteWorkflowStepGroup, "delete-workflow-step-group", "", false, "Delete Workflow Step Group")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorGetTemplate, "get-template", "", false, "Get Template")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorGetTemplateStep, "get-template-step", "", false, "Get Template Step")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorGetTemplateStepGroup, "get-template-step-group", "", false, "Get Template Step Group")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorGetWorkflow, "get-workflow", "", false, "Get Workflow")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorGetWorkflowStep, "get-workflow-step", "", false, "Get Workflow Step")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorGetWorkflowStepGroup, "get-workflow-step-group", "", false, "Get Workflow Step Group")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorListPlugins, "list-plugins", "", false, "List Plugins")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorListTemplateStepGroups, "list-template-step-groups", "", false, "List Template Step Groups")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorListTemplateSteps, "list-template-steps", "", false, "List Template Steps")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorListTemplates, "list-templates", "", false, "List Templates")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorListWorkflowStepGroups, "list-workflow-step-groups", "", false, "List Workflow Step Groups")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorListWorkflowSteps, "list-workflow-steps", "", false, "List Workflow Steps")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorListWorkflows, "list-workflows", "", false, "List Workflows")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorRetryWorkflowStep, "retry-workflow-step", "", false, "Retry Workflow Step")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorStartWorkflow, "start-workflow", "", false, "Start Workflow")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorStopWorkflow, "stop-workflow", "", false, "Stop Workflow")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorTagResource, "tag-resource", "", false, "Tag Resource")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorUntagResource, "untag-resource", "", false, "Untag Resource")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorUpdateTemplate, "update-template", "", false, "Update Template")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorUpdateWorkflow, "update-workflow", "", false, "Update Workflow")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorUpdateWorkflowStep, "update-workflow-step", "", false, "Update Workflow Step")
	_migrationhuborchestratorCmd.Flags().BoolVarP(&_migrationhuborchestratorUpdateWorkflowStepGroup, "update-workflow-step-group", "", false, "Update Workflow Step Group")

}
