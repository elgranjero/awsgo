package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fis"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// fisCmd represents the fis command
var _fisCmd = &cobra.Command{
	Use:   "fis",
	Short: "AWS fis CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := fis.NewFromConfig(cfg)
		if _fisCreateExperimentTemplate {
			fis_CreateExperimentTemplate(cfg, client)
			return
		}
		if _fisCreateTargetAccountConfiguration {
			fis_CreateTargetAccountConfiguration(cfg, client)
			return
		}
		if _fisDeleteExperimentTemplate {
			fis_DeleteExperimentTemplate(cfg, client)
			return
		}
		if _fisDeleteTargetAccountConfiguration {
			fis_DeleteTargetAccountConfiguration(cfg, client)
			return
		}
		if _fisGetAction {
			fis_GetAction(cfg, client)
			return
		}
		if _fisGetExperiment {
			fis_GetExperiment(cfg, client)
			return
		}
		if _fisGetExperimentTargetAccountConfiguration {
			fis_GetExperimentTargetAccountConfiguration(cfg, client)
			return
		}
		if _fisGetExperimentTemplate {
			fis_GetExperimentTemplate(cfg, client)
			return
		}
		if _fisGetSafetyLever {
			fis_GetSafetyLever(cfg, client)
			return
		}
		if _fisGetTargetAccountConfiguration {
			fis_GetTargetAccountConfiguration(cfg, client)
			return
		}
		if _fisGetTargetResourceType {
			fis_GetTargetResourceType(cfg, client)
			return
		}
		if _fisListActions {
			fis_ListActions(cfg, client)
			return
		}
		if _fisListExperimentResolvedTargets {
			fis_ListExperimentResolvedTargets(cfg, client)
			return
		}
		if _fisListExperimentTargetAccountConfigurations {
			fis_ListExperimentTargetAccountConfigurations(cfg, client)
			return
		}
		if _fisListExperimentTemplates {
			fis_ListExperimentTemplates(cfg, client)
			return
		}
		if _fisListExperiments {
			fis_ListExperiments(cfg, client)
			return
		}
		if _fisListTagsForResource {
			fis_ListTagsForResource(cfg, client)
			return
		}
		if _fisListTargetAccountConfigurations {
			fis_ListTargetAccountConfigurations(cfg, client)
			return
		}
		if _fisListTargetResourceTypes {
			fis_ListTargetResourceTypes(cfg, client)
			return
		}
		if _fisStartExperiment {
			fis_StartExperiment(cfg, client)
			return
		}
		if _fisStopExperiment {
			fis_StopExperiment(cfg, client)
			return
		}
		if _fisTagResource {
			fis_TagResource(cfg, client)
			return
		}
		if _fisUntagResource {
			fis_UntagResource(cfg, client)
			return
		}
		if _fisUpdateExperimentTemplate {
			fis_UpdateExperimentTemplate(cfg, client)
			return
		}
		if _fisUpdateSafetyLeverState {
			fis_UpdateSafetyLeverState(cfg, client)
			return
		}
		if _fisUpdateTargetAccountConfiguration {
			fis_UpdateTargetAccountConfiguration(cfg, client)
			return
		}

	},
}

var (
	_fisCreateExperimentTemplate                  bool
	_fisCreateTargetAccountConfiguration          bool
	_fisDeleteExperimentTemplate                  bool
	_fisDeleteTargetAccountConfiguration          bool
	_fisGetAction                                 bool
	_fisGetExperiment                             bool
	_fisGetExperimentTargetAccountConfiguration   bool
	_fisGetExperimentTemplate                     bool
	_fisGetSafetyLever                            bool
	_fisGetTargetAccountConfiguration             bool
	_fisGetTargetResourceType                     bool
	_fisListActions                               bool
	_fisListExperimentResolvedTargets             bool
	_fisListExperimentTargetAccountConfigurations bool
	_fisListExperimentTemplates                   bool
	_fisListExperiments                           bool
	_fisListTagsForResource                       bool
	_fisListTargetAccountConfigurations           bool
	_fisListTargetResourceTypes                   bool
	_fisStartExperiment                           bool
	_fisStopExperiment                            bool
	_fisTagResource                               bool
	_fisUntagResource                             bool
	_fisUpdateExperimentTemplate                  bool
	_fisUpdateSafetyLeverState                    bool
	_fisUpdateTargetAccountConfiguration          bool

	_fisAccountId                     string
	_fisActions                       string
	_fisClientToken                   string
	_fisDescription                   string
	_fisExperimentId                  string
	_fisExperimentOptions             string
	_fisExperimentReportConfiguration string
	_fisExperimentTemplateId          string
	_fisId                            string
	_fisLogConfiguration              string
	_fisMaxResults                    string
	_fisNextToken                     string
	_fisResourceArn                   string
	_fisResourceType                  string
	_fisRoleArn                       string
	_fisState                         string
	_fisStopConditions                string
	_fisTagKeys                       []string
	_fisTags                          string
	_fisTargetName                    string
	_fisTargets                       string
)

// Creates an experiment template.
// An experiment template includes the following components:
//
// - Targets: A target can be a specific resource in your Amazon Web Services
// environment, or one or more resources that match criteria that you specify, for
// example, resources that have specific tags.
//
// - Actions: The actions to carry out on the target. You can specify multiple
// actions, the duration of each action, and when to start each action during an
// experiment.
//
// - Stop conditions: If a stop condition is triggered while an experiment is
// running, the experiment is automatically stopped. You can define a stop
// condition as a CloudWatch alarm.
//
// For more information, see [experiment templates] in the Fault Injection Service User Guide.
//
// [experiment templates]: https://docs.aws.amazon.com/fis/latest/userguide/experiment-templates.html
func fis_CreateExperimentTemplate(cfg aws.Config, client *fis.Client) {
	input := &fis.CreateExperimentTemplateInput{
		// Actions: map[string]types.CreateExperimentTemplateActionInput, // Required
		// ClientToken: *string, // Required
		// Description: *string, // Required
		// RoleArn: *string, // Required
		// StopConditions: []types.CreateExperimentTemplateStopConditionInput, // Required
	}

	if len(_fisActions) > 0 {
		if err := assignInputField(input, "Actions", _fisActions); err != nil {
			log.Errorf("invalid --actions: %s", err.Error())
			return
		}
	}
	if len(_fisClientToken) > 0 {
		input.ClientToken = aws.String(_fisClientToken)
	}
	if len(_fisDescription) > 0 {
		input.Description = aws.String(_fisDescription)
	}
	if len(_fisRoleArn) > 0 {
		input.RoleArn = aws.String(_fisRoleArn)
	}
	if len(_fisStopConditions) > 0 {
		if err := assignInputField(input, "StopConditions", _fisStopConditions); err != nil {
			log.Errorf("invalid --stop-conditions: %s", err.Error())
			return
		}
	}
	if len(_fisExperimentOptions) > 0 {
		if err := assignInputField(input, "ExperimentOptions", _fisExperimentOptions); err != nil {
			log.Errorf("invalid --experiment-options: %s", err.Error())
			return
		}
	}
	if len(_fisExperimentReportConfiguration) > 0 {
		if err := assignInputField(input, "ExperimentReportConfiguration", _fisExperimentReportConfiguration); err != nil {
			log.Errorf("invalid --experiment-report-configuration: %s", err.Error())
			return
		}
	}
	if len(_fisLogConfiguration) > 0 {
		if err := assignInputField(input, "LogConfiguration", _fisLogConfiguration); err != nil {
			log.Errorf("invalid --log-configuration: %s", err.Error())
			return
		}
	}
	if len(_fisTags) > 0 {
		if err := assignInputField(input, "Tags", _fisTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_fisTargets) > 0 {
		if err := assignInputField(input, "Targets", _fisTargets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateExperimentTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a target account configuration for the experiment template. A target
// account configuration is required when accountTargeting of experimentOptions is
// set to multi-account . For more information, see [experiment options] in the Fault Injection
// Service User Guide.
//
// [experiment options]: https://docs.aws.amazon.com/fis/latest/userguide/experiment-options.html
func fis_CreateTargetAccountConfiguration(cfg aws.Config, client *fis.Client) {
	input := &fis.CreateTargetAccountConfigurationInput{
		// AccountId: *string, // Required
		// ExperimentTemplateId: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_fisAccountId) > 0 {
		input.AccountId = aws.String(_fisAccountId)
	}
	if len(_fisExperimentTemplateId) > 0 {
		input.ExperimentTemplateId = aws.String(_fisExperimentTemplateId)
	}
	if len(_fisRoleArn) > 0 {
		input.RoleArn = aws.String(_fisRoleArn)
	}
	if len(_fisClientToken) > 0 {
		input.ClientToken = aws.String(_fisClientToken)
	}
	if len(_fisDescription) > 0 {
		input.Description = aws.String(_fisDescription)
	}

	if resp, err := client.CreateTargetAccountConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified experiment template.
func fis_DeleteExperimentTemplate(cfg aws.Config, client *fis.Client) {
	input := &fis.DeleteExperimentTemplateInput{
		// Id: *string, // Required
	}

	if len(_fisId) > 0 {
		input.Id = aws.String(_fisId)
	}

	if resp, err := client.DeleteExperimentTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified target account configuration of the experiment template.
func fis_DeleteTargetAccountConfiguration(cfg aws.Config, client *fis.Client) {
	input := &fis.DeleteTargetAccountConfigurationInput{
		// AccountId: *string, // Required
		// ExperimentTemplateId: *string, // Required
	}

	if len(_fisAccountId) > 0 {
		input.AccountId = aws.String(_fisAccountId)
	}
	if len(_fisExperimentTemplateId) > 0 {
		input.ExperimentTemplateId = aws.String(_fisExperimentTemplateId)
	}

	if resp, err := client.DeleteTargetAccountConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified FIS action.
func fis_GetAction(cfg aws.Config, client *fis.Client) {
	input := &fis.GetActionInput{
		// Id: *string, // Required
	}

	if len(_fisId) > 0 {
		input.Id = aws.String(_fisId)
	}

	if resp, err := client.GetAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified experiment.
func fis_GetExperiment(cfg aws.Config, client *fis.Client) {
	input := &fis.GetExperimentInput{
		// Id: *string, // Required
	}

	if len(_fisId) > 0 {
		input.Id = aws.String(_fisId)
	}

	if resp, err := client.GetExperiment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified target account configuration of the
// experiment.
func fis_GetExperimentTargetAccountConfiguration(cfg aws.Config, client *fis.Client) {
	input := &fis.GetExperimentTargetAccountConfigurationInput{
		// AccountId: *string, // Required
		// ExperimentId: *string, // Required
	}

	if len(_fisAccountId) > 0 {
		input.AccountId = aws.String(_fisAccountId)
	}
	if len(_fisExperimentId) > 0 {
		input.ExperimentId = aws.String(_fisExperimentId)
	}

	if resp, err := client.GetExperimentTargetAccountConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified experiment template.
func fis_GetExperimentTemplate(cfg aws.Config, client *fis.Client) {
	input := &fis.GetExperimentTemplateInput{
		// Id: *string, // Required
	}

	if len(_fisId) > 0 {
		input.Id = aws.String(_fisId)
	}

	if resp, err := client.GetExperimentTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified safety lever.
func fis_GetSafetyLever(cfg aws.Config, client *fis.Client) {
	input := &fis.GetSafetyLeverInput{
		// Id: *string, // Required
	}

	if len(_fisId) > 0 {
		input.Id = aws.String(_fisId)
	}

	if resp, err := client.GetSafetyLever(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified target account configuration of the
// experiment template.
func fis_GetTargetAccountConfiguration(cfg aws.Config, client *fis.Client) {
	input := &fis.GetTargetAccountConfigurationInput{
		// AccountId: *string, // Required
		// ExperimentTemplateId: *string, // Required
	}

	if len(_fisAccountId) > 0 {
		input.AccountId = aws.String(_fisAccountId)
	}
	if len(_fisExperimentTemplateId) > 0 {
		input.ExperimentTemplateId = aws.String(_fisExperimentTemplateId)
	}

	if resp, err := client.GetTargetAccountConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified resource type.
func fis_GetTargetResourceType(cfg aws.Config, client *fis.Client) {
	input := &fis.GetTargetResourceTypeInput{
		// ResourceType: *string, // Required
	}

	if len(_fisResourceType) > 0 {
		input.ResourceType = aws.String(_fisResourceType)
	}

	if resp, err := client.GetTargetResourceType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the available FIS actions.
func fis_ListActions(cfg aws.Config, client *fis.Client) {
	input := &fis.ListActionsInput{}

	if len(_fisMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fisMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fisNextToken) > 0 {
		input.NextToken = aws.String(_fisNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*fis.ListActionsOutput
	p := fis.NewListActionsPaginator(client, input)
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

// Lists the resolved targets information of the specified experiment.
func fis_ListExperimentResolvedTargets(cfg aws.Config, client *fis.Client) {
	input := &fis.ListExperimentResolvedTargetsInput{
		// ExperimentId: *string, // Required
	}

	if len(_fisExperimentId) > 0 {
		input.ExperimentId = aws.String(_fisExperimentId)
	}
	if len(_fisMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fisMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fisNextToken) > 0 {
		input.NextToken = aws.String(_fisNextToken)
	}
	if len(_fisTargetName) > 0 {
		input.TargetName = aws.String(_fisTargetName)
	}

	if disablePaginator() {
		if resp, err := client.ListExperimentResolvedTargets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*fis.ListExperimentResolvedTargetsOutput
	p := fis.NewListExperimentResolvedTargetsPaginator(client, input)
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

// Lists the target account configurations of the specified experiment.
func fis_ListExperimentTargetAccountConfigurations(cfg aws.Config, client *fis.Client) {
	input := &fis.ListExperimentTargetAccountConfigurationsInput{
		// ExperimentId: *string, // Required
	}

	if len(_fisExperimentId) > 0 {
		input.ExperimentId = aws.String(_fisExperimentId)
	}
	if len(_fisNextToken) > 0 {
		input.NextToken = aws.String(_fisNextToken)
	}

	if resp, err := client.ListExperimentTargetAccountConfigurations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists your experiment templates.
func fis_ListExperimentTemplates(cfg aws.Config, client *fis.Client) {
	input := &fis.ListExperimentTemplatesInput{}

	if len(_fisMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fisMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fisNextToken) > 0 {
		input.NextToken = aws.String(_fisNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListExperimentTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*fis.ListExperimentTemplatesOutput
	p := fis.NewListExperimentTemplatesPaginator(client, input)
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

// Lists your experiments.
func fis_ListExperiments(cfg aws.Config, client *fis.Client) {
	input := &fis.ListExperimentsInput{}

	if len(_fisExperimentTemplateId) > 0 {
		input.ExperimentTemplateId = aws.String(_fisExperimentTemplateId)
	}
	if len(_fisMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fisMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fisNextToken) > 0 {
		input.NextToken = aws.String(_fisNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListExperiments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*fis.ListExperimentsOutput
	p := fis.NewListExperimentsPaginator(client, input)
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

// Lists the tags for the specified resource.
func fis_ListTagsForResource(cfg aws.Config, client *fis.Client) {
	input := &fis.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_fisResourceArn) > 0 {
		input.ResourceArn = aws.String(_fisResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the target account configurations of the specified experiment template.
func fis_ListTargetAccountConfigurations(cfg aws.Config, client *fis.Client) {
	input := &fis.ListTargetAccountConfigurationsInput{
		// ExperimentTemplateId: *string, // Required
	}

	if len(_fisExperimentTemplateId) > 0 {
		input.ExperimentTemplateId = aws.String(_fisExperimentTemplateId)
	}
	if len(_fisMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fisMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fisNextToken) > 0 {
		input.NextToken = aws.String(_fisNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTargetAccountConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*fis.ListTargetAccountConfigurationsOutput
	p := fis.NewListTargetAccountConfigurationsPaginator(client, input)
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

// Lists the target resource types.
func fis_ListTargetResourceTypes(cfg aws.Config, client *fis.Client) {
	input := &fis.ListTargetResourceTypesInput{}

	if len(_fisMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fisMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fisNextToken) > 0 {
		input.NextToken = aws.String(_fisNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTargetResourceTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*fis.ListTargetResourceTypesOutput
	p := fis.NewListTargetResourceTypesPaginator(client, input)
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

// Starts running an experiment from the specified experiment template.
func fis_StartExperiment(cfg aws.Config, client *fis.Client) {
	input := &fis.StartExperimentInput{
		// ClientToken: *string, // Required
		// ExperimentTemplateId: *string, // Required
	}

	if len(_fisClientToken) > 0 {
		input.ClientToken = aws.String(_fisClientToken)
	}
	if len(_fisExperimentTemplateId) > 0 {
		input.ExperimentTemplateId = aws.String(_fisExperimentTemplateId)
	}
	if len(_fisExperimentOptions) > 0 {
		if err := assignInputField(input, "ExperimentOptions", _fisExperimentOptions); err != nil {
			log.Errorf("invalid --experiment-options: %s", err.Error())
			return
		}
	}
	if len(_fisTags) > 0 {
		if err := assignInputField(input, "Tags", _fisTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartExperiment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the specified experiment.
func fis_StopExperiment(cfg aws.Config, client *fis.Client) {
	input := &fis.StopExperimentInput{
		// Id: *string, // Required
	}

	if len(_fisId) > 0 {
		input.Id = aws.String(_fisId)
	}

	if resp, err := client.StopExperiment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies the specified tags to the specified resource.
func fis_TagResource(cfg aws.Config, client *fis.Client) {
	input := &fis.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_fisResourceArn) > 0 {
		input.ResourceArn = aws.String(_fisResourceArn)
	}
	if len(_fisTags) > 0 {
		if err := assignInputField(input, "Tags", _fisTags); err != nil {
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

// Removes the specified tags from the specified resource.
func fis_UntagResource(cfg aws.Config, client *fis.Client) {
	input := &fis.UntagResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_fisResourceArn) > 0 {
		input.ResourceArn = aws.String(_fisResourceArn)
	}
	if len(_fisTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _fisTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified experiment template.
func fis_UpdateExperimentTemplate(cfg aws.Config, client *fis.Client) {
	input := &fis.UpdateExperimentTemplateInput{
		// Id: *string, // Required
	}

	if len(_fisId) > 0 {
		input.Id = aws.String(_fisId)
	}
	if len(_fisActions) > 0 {
		if err := assignInputField(input, "Actions", _fisActions); err != nil {
			log.Errorf("invalid --actions: %s", err.Error())
			return
		}
	}
	if len(_fisDescription) > 0 {
		input.Description = aws.String(_fisDescription)
	}
	if len(_fisExperimentOptions) > 0 {
		if err := assignInputField(input, "ExperimentOptions", _fisExperimentOptions); err != nil {
			log.Errorf("invalid --experiment-options: %s", err.Error())
			return
		}
	}
	if len(_fisExperimentReportConfiguration) > 0 {
		if err := assignInputField(input, "ExperimentReportConfiguration", _fisExperimentReportConfiguration); err != nil {
			log.Errorf("invalid --experiment-report-configuration: %s", err.Error())
			return
		}
	}
	if len(_fisLogConfiguration) > 0 {
		if err := assignInputField(input, "LogConfiguration", _fisLogConfiguration); err != nil {
			log.Errorf("invalid --log-configuration: %s", err.Error())
			return
		}
	}
	if len(_fisRoleArn) > 0 {
		input.RoleArn = aws.String(_fisRoleArn)
	}
	if len(_fisStopConditions) > 0 {
		if err := assignInputField(input, "StopConditions", _fisStopConditions); err != nil {
			log.Errorf("invalid --stop-conditions: %s", err.Error())
			return
		}
	}
	if len(_fisTargets) > 0 {
		if err := assignInputField(input, "Targets", _fisTargets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateExperimentTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified safety lever state.
func fis_UpdateSafetyLeverState(cfg aws.Config, client *fis.Client) {
	input := &fis.UpdateSafetyLeverStateInput{
		// Id: *string, // Required
		// State: *types.UpdateSafetyLeverStateInput, // Required
	}

	if len(_fisId) > 0 {
		input.Id = aws.String(_fisId)
	}
	if len(_fisState) > 0 {
		if err := assignInputField(input, "State", _fisState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSafetyLeverState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the target account configuration for the specified experiment template.
func fis_UpdateTargetAccountConfiguration(cfg aws.Config, client *fis.Client) {
	input := &fis.UpdateTargetAccountConfigurationInput{
		// AccountId: *string, // Required
		// ExperimentTemplateId: *string, // Required
	}

	if len(_fisAccountId) > 0 {
		input.AccountId = aws.String(_fisAccountId)
	}
	if len(_fisExperimentTemplateId) > 0 {
		input.ExperimentTemplateId = aws.String(_fisExperimentTemplateId)
	}
	if len(_fisDescription) > 0 {
		input.Description = aws.String(_fisDescription)
	}
	if len(_fisRoleArn) > 0 {
		input.RoleArn = aws.String(_fisRoleArn)
	}

	if resp, err := client.UpdateTargetAccountConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_fisCmd)
	_fisCmd.Flags().SortFlags = false

	_fisCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_fisCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_fisCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_fisCmd.Flags().StringVarP(&_fisAccountId, "account-id", "", "", "Account ID")
	_fisCmd.Flags().StringVarP(&_fisActions, "actions", "", "", "Actions")
	_fisCmd.Flags().StringVarP(&_fisClientToken, "client-token", "", "", "Client Token")
	_fisCmd.Flags().StringVarP(&_fisDescription, "description", "", "", "Description")
	_fisCmd.Flags().StringVarP(&_fisExperimentId, "experiment-id", "", "", "Experiment ID")
	_fisCmd.Flags().StringVarP(&_fisExperimentOptions, "experiment-options", "", "", "Experiment Options")
	_fisCmd.Flags().StringVarP(&_fisExperimentReportConfiguration, "experiment-report-configuration", "", "", "Experiment Report Configuration")
	_fisCmd.Flags().StringVarP(&_fisExperimentTemplateId, "experiment-template-id", "", "", "Experiment Template ID")
	_fisCmd.Flags().StringVarP(&_fisId, "id", "", "", "ID")
	_fisCmd.Flags().StringVarP(&_fisLogConfiguration, "log-configuration", "", "", "Log Configuration")
	_fisCmd.Flags().StringVarP(&_fisMaxResults, "max-results", "", "", "Max Results")
	_fisCmd.Flags().StringVarP(&_fisNextToken, "next-token", "", "", "Next Token")
	_fisCmd.Flags().StringVarP(&_fisResourceArn, "resource-arn", "", "", "Resource ARN")
	_fisCmd.Flags().StringVarP(&_fisResourceType, "resource-type", "", "", "Resource Type")
	_fisCmd.Flags().StringVarP(&_fisRoleArn, "role-arn", "", "", "Role ARN")
	_fisCmd.Flags().StringVarP(&_fisState, "state", "", "", "State")
	_fisCmd.Flags().StringVarP(&_fisStopConditions, "stop-conditions", "", "", "Stop Conditions")
	_fisCmd.Flags().StringSliceVarP(&_fisTagKeys, "tag-keys", "", nil, "Tag Keys")
	_fisCmd.Flags().StringVarP(&_fisTags, "tags", "", "", "Tags")
	_fisCmd.Flags().StringVarP(&_fisTargetName, "target-name", "", "", "Target Name")
	_fisCmd.Flags().StringVarP(&_fisTargets, "targets", "", "", "Targets")

	_fisCmd.Flags().BoolVarP(&_fisCreateExperimentTemplate, "create-experiment-template", "", false, "Create Experiment Template")
	_fisCmd.Flags().BoolVarP(&_fisCreateTargetAccountConfiguration, "create-target-account-configuration", "", false, "Create Target Account Configuration")
	_fisCmd.Flags().BoolVarP(&_fisDeleteExperimentTemplate, "delete-experiment-template", "", false, "Delete Experiment Template")
	_fisCmd.Flags().BoolVarP(&_fisDeleteTargetAccountConfiguration, "delete-target-account-configuration", "", false, "Delete Target Account Configuration")
	_fisCmd.Flags().BoolVarP(&_fisGetAction, "get-action", "", false, "Get Action")
	_fisCmd.Flags().BoolVarP(&_fisGetExperiment, "get-experiment", "", false, "Get Experiment")
	_fisCmd.Flags().BoolVarP(&_fisGetExperimentTargetAccountConfiguration, "get-experiment-target-account-configuration", "", false, "Get Experiment Target Account Configuration")
	_fisCmd.Flags().BoolVarP(&_fisGetExperimentTemplate, "get-experiment-template", "", false, "Get Experiment Template")
	_fisCmd.Flags().BoolVarP(&_fisGetSafetyLever, "get-safety-lever", "", false, "Get Safety Lever")
	_fisCmd.Flags().BoolVarP(&_fisGetTargetAccountConfiguration, "get-target-account-configuration", "", false, "Get Target Account Configuration")
	_fisCmd.Flags().BoolVarP(&_fisGetTargetResourceType, "get-target-resource-type", "", false, "Get Target Resource Type")
	_fisCmd.Flags().BoolVarP(&_fisListActions, "list-actions", "", false, "List Actions")
	_fisCmd.Flags().BoolVarP(&_fisListExperimentResolvedTargets, "list-experiment-resolved-targets", "", false, "List Experiment Resolved Targets")
	_fisCmd.Flags().BoolVarP(&_fisListExperimentTargetAccountConfigurations, "list-experiment-target-account-configurations", "", false, "List Experiment Target Account Configurations")
	_fisCmd.Flags().BoolVarP(&_fisListExperimentTemplates, "list-experiment-templates", "", false, "List Experiment Templates")
	_fisCmd.Flags().BoolVarP(&_fisListExperiments, "list-experiments", "", false, "List Experiments")
	_fisCmd.Flags().BoolVarP(&_fisListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_fisCmd.Flags().BoolVarP(&_fisListTargetAccountConfigurations, "list-target-account-configurations", "", false, "List Target Account Configurations")
	_fisCmd.Flags().BoolVarP(&_fisListTargetResourceTypes, "list-target-resource-types", "", false, "List Target Resource Types")
	_fisCmd.Flags().BoolVarP(&_fisStartExperiment, "start-experiment", "", false, "Start Experiment")
	_fisCmd.Flags().BoolVarP(&_fisStopExperiment, "stop-experiment", "", false, "Stop Experiment")
	_fisCmd.Flags().BoolVarP(&_fisTagResource, "tag-resource", "", false, "Tag Resource")
	_fisCmd.Flags().BoolVarP(&_fisUntagResource, "untag-resource", "", false, "Untag Resource")
	_fisCmd.Flags().BoolVarP(&_fisUpdateExperimentTemplate, "update-experiment-template", "", false, "Update Experiment Template")
	_fisCmd.Flags().BoolVarP(&_fisUpdateSafetyLeverState, "update-safety-lever-state", "", false, "Update Safety Lever State")
	_fisCmd.Flags().BoolVarP(&_fisUpdateTargetAccountConfiguration, "update-target-account-configuration", "", false, "Update Target Account Configuration")

}
