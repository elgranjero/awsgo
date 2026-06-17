package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/databrew"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// databrewCmd represents the databrew command
var _databrewCmd = &cobra.Command{
	Use:   "databrew",
	Short: "AWS databrew CLI",
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
		client := databrew.NewFromConfig(cfg)
		if _databrewBatchDeleteRecipeVersion {
			databrew_BatchDeleteRecipeVersion(cfg, client)
			return
		}
		if _databrewCreateDataset {
			databrew_CreateDataset(cfg, client)
			return
		}
		if _databrewCreateProfileJob {
			databrew_CreateProfileJob(cfg, client)
			return
		}
		if _databrewCreateProject {
			databrew_CreateProject(cfg, client)
			return
		}
		if _databrewCreateRecipe {
			databrew_CreateRecipe(cfg, client)
			return
		}
		if _databrewCreateRecipeJob {
			databrew_CreateRecipeJob(cfg, client)
			return
		}
		if _databrewCreateRuleset {
			databrew_CreateRuleset(cfg, client)
			return
		}
		if _databrewCreateSchedule {
			databrew_CreateSchedule(cfg, client)
			return
		}
		if _databrewDeleteDataset {
			databrew_DeleteDataset(cfg, client)
			return
		}
		if _databrewDeleteJob {
			databrew_DeleteJob(cfg, client)
			return
		}
		if _databrewDeleteProject {
			databrew_DeleteProject(cfg, client)
			return
		}
		if _databrewDeleteRecipeVersion {
			databrew_DeleteRecipeVersion(cfg, client)
			return
		}
		if _databrewDeleteRuleset {
			databrew_DeleteRuleset(cfg, client)
			return
		}
		if _databrewDeleteSchedule {
			databrew_DeleteSchedule(cfg, client)
			return
		}
		if _databrewDescribeDataset {
			databrew_DescribeDataset(cfg, client)
			return
		}
		if _databrewDescribeJob {
			databrew_DescribeJob(cfg, client)
			return
		}
		if _databrewDescribeJobRun {
			databrew_DescribeJobRun(cfg, client)
			return
		}
		if _databrewDescribeProject {
			databrew_DescribeProject(cfg, client)
			return
		}
		if _databrewDescribeRecipe {
			databrew_DescribeRecipe(cfg, client)
			return
		}
		if _databrewDescribeRuleset {
			databrew_DescribeRuleset(cfg, client)
			return
		}
		if _databrewDescribeSchedule {
			databrew_DescribeSchedule(cfg, client)
			return
		}
		if _databrewListDatasets {
			databrew_ListDatasets(cfg, client)
			return
		}
		if _databrewListJobRuns {
			databrew_ListJobRuns(cfg, client)
			return
		}
		if _databrewListJobs {
			databrew_ListJobs(cfg, client)
			return
		}
		if _databrewListProjects {
			databrew_ListProjects(cfg, client)
			return
		}
		if _databrewListRecipeVersions {
			databrew_ListRecipeVersions(cfg, client)
			return
		}
		if _databrewListRecipes {
			databrew_ListRecipes(cfg, client)
			return
		}
		if _databrewListRulesets {
			databrew_ListRulesets(cfg, client)
			return
		}
		if _databrewListSchedules {
			databrew_ListSchedules(cfg, client)
			return
		}
		if _databrewListTagsForResource {
			databrew_ListTagsForResource(cfg, client)
			return
		}
		if _databrewPublishRecipe {
			databrew_PublishRecipe(cfg, client)
			return
		}
		if _databrewSendProjectSessionAction {
			databrew_SendProjectSessionAction(cfg, client)
			return
		}
		if _databrewStartJobRun {
			databrew_StartJobRun(cfg, client)
			return
		}
		if _databrewStartProjectSession {
			databrew_StartProjectSession(cfg, client)
			return
		}
		if _databrewStopJobRun {
			databrew_StopJobRun(cfg, client)
			return
		}
		if _databrewTagResource {
			databrew_TagResource(cfg, client)
			return
		}
		if _databrewUntagResource {
			databrew_UntagResource(cfg, client)
			return
		}
		if _databrewUpdateDataset {
			databrew_UpdateDataset(cfg, client)
			return
		}
		if _databrewUpdateProfileJob {
			databrew_UpdateProfileJob(cfg, client)
			return
		}
		if _databrewUpdateProject {
			databrew_UpdateProject(cfg, client)
			return
		}
		if _databrewUpdateRecipe {
			databrew_UpdateRecipe(cfg, client)
			return
		}
		if _databrewUpdateRecipeJob {
			databrew_UpdateRecipeJob(cfg, client)
			return
		}
		if _databrewUpdateRuleset {
			databrew_UpdateRuleset(cfg, client)
			return
		}
		if _databrewUpdateSchedule {
			databrew_UpdateSchedule(cfg, client)
			return
		}

	},
}

var (
	_databrewBatchDeleteRecipeVersion bool
	_databrewCreateDataset            bool
	_databrewCreateProfileJob         bool
	_databrewCreateProject            bool
	_databrewCreateRecipe             bool
	_databrewCreateRecipeJob          bool
	_databrewCreateRuleset            bool
	_databrewCreateSchedule           bool
	_databrewDeleteDataset            bool
	_databrewDeleteJob                bool
	_databrewDeleteProject            bool
	_databrewDeleteRecipeVersion      bool
	_databrewDeleteRuleset            bool
	_databrewDeleteSchedule           bool
	_databrewDescribeDataset          bool
	_databrewDescribeJob              bool
	_databrewDescribeJobRun           bool
	_databrewDescribeProject          bool
	_databrewDescribeRecipe           bool
	_databrewDescribeRuleset          bool
	_databrewDescribeSchedule         bool
	_databrewListDatasets             bool
	_databrewListJobRuns              bool
	_databrewListJobs                 bool
	_databrewListProjects             bool
	_databrewListRecipeVersions       bool
	_databrewListRecipes              bool
	_databrewListRulesets             bool
	_databrewListSchedules            bool
	_databrewListTagsForResource      bool
	_databrewPublishRecipe            bool
	_databrewSendProjectSessionAction bool
	_databrewStartJobRun              bool
	_databrewStartProjectSession      bool
	_databrewStopJobRun               bool
	_databrewTagResource              bool
	_databrewUntagResource            bool
	_databrewUpdateDataset            bool
	_databrewUpdateProfileJob         bool
	_databrewUpdateProject            bool
	_databrewUpdateRecipe             bool
	_databrewUpdateRecipeJob          bool
	_databrewUpdateRuleset            bool
	_databrewUpdateSchedule           bool

	_databrewAssumeControl            string
	_databrewClientSessionId          string
	_databrewConfiguration            string
	_databrewCronExpression           string
	_databrewDataCatalogOutputs       string
	_databrewDatabaseOutputs          string
	_databrewDatasetName              string
	_databrewDescription              string
	_databrewEncryptionKeyArn         string
	_databrewEncryptionMode           string
	_databrewFormat                   string
	_databrewFormatOptions            string
	_databrewInput                    string
	_databrewJobName                  string
	_databrewJobNames                 []string
	_databrewJobSample                string
	_databrewLogSubscription          string
	_databrewMaxCapacity              string
	_databrewMaxResults               string
	_databrewMaxRetries               string
	_databrewName                     string
	_databrewNextToken                string
	_databrewOutputLocation           string
	_databrewOutputs                  string
	_databrewPathOptions              string
	_databrewPreview                  string
	_databrewProjectName              string
	_databrewRecipeName               string
	_databrewRecipeReference          string
	_databrewRecipeStep               string
	_databrewRecipeVersion            string
	_databrewRecipeVersions           []string
	_databrewResourceArn              string
	_databrewRoleArn                  string
	_databrewRules                    string
	_databrewRunId                    string
	_databrewSample                   string
	_databrewStepIndex                string
	_databrewSteps                    string
	_databrewTagKeys                  []string
	_databrewTags                     string
	_databrewTargetArn                string
	_databrewTimeout                  string
	_databrewValidationConfigurations string
	_databrewViewFrame                string
)

// Deletes one or more versions of a recipe at a time.
// The entire request will be rejected if:
//
// - The recipe does not exist.
//
// - There is an invalid version identifier in the list of versions.
//
// - The version list is empty.
//
// - The version list size exceeds 50.
//
// - The version list contains duplicate entries.
//
// The request will complete successfully, but with partial failures, if:
//
// - A version does not exist.
//
// - A version is being used by a job.
//
// - You specify LATEST_WORKING , but it's being used by a project.
//
// - The version fails to be deleted.
//
// The LATEST_WORKING version will only be deleted if the recipe has no other
// versions. If you try to delete LATEST_WORKING while other versions exist (or if
// they can't be deleted), then LATEST_WORKING will be listed as partial failure
// in the response.
func databrew_BatchDeleteRecipeVersion(cfg aws.Config, client *databrew.Client) {
	input := &databrew.BatchDeleteRecipeVersionInput{
		// Name: *string, // Required
		// RecipeVersions: []string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}
	if len(_databrewRecipeVersions) > 0 {
		input.RecipeVersions = append([]string(nil), _databrewRecipeVersions...)
	}

	if resp, err := client.BatchDeleteRecipeVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new DataBrew dataset.
func databrew_CreateDataset(cfg aws.Config, client *databrew.Client) {
	input := &databrew.CreateDatasetInput{
		// Input: *types.Input, // Required
		// Name: *string, // Required
	}

	if len(_databrewInput) > 0 {
		if err := assignInputField(input, "Input", _databrewInput); err != nil {
			log.Errorf("invalid --input: %s", err.Error())
			return
		}
	}
	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}
	if len(_databrewFormat) > 0 {
		if err := assignInputField(input, "Format", _databrewFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_databrewFormatOptions) > 0 {
		if err := assignInputField(input, "FormatOptions", _databrewFormatOptions); err != nil {
			log.Errorf("invalid --format-options: %s", err.Error())
			return
		}
	}
	if len(_databrewPathOptions) > 0 {
		if err := assignInputField(input, "PathOptions", _databrewPathOptions); err != nil {
			log.Errorf("invalid --path-options: %s", err.Error())
			return
		}
	}
	if len(_databrewTags) > 0 {
		if err := assignInputField(input, "Tags", _databrewTags); err != nil {
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

// Creates a new job to analyze a dataset and create its data profile.
func databrew_CreateProfileJob(cfg aws.Config, client *databrew.Client) {
	input := &databrew.CreateProfileJobInput{
		// DatasetName: *string, // Required
		// Name: *string, // Required
		// OutputLocation: *types.S3Location, // Required
		// RoleArn: *string, // Required
	}

	if len(_databrewDatasetName) > 0 {
		input.DatasetName = aws.String(_databrewDatasetName)
	}
	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}
	if len(_databrewOutputLocation) > 0 {
		if err := assignInputField(input, "OutputLocation", _databrewOutputLocation); err != nil {
			log.Errorf("invalid --output-location: %s", err.Error())
			return
		}
	}
	if len(_databrewRoleArn) > 0 {
		input.RoleArn = aws.String(_databrewRoleArn)
	}
	if len(_databrewConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _databrewConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_databrewEncryptionKeyArn) > 0 {
		input.EncryptionKeyArn = aws.String(_databrewEncryptionKeyArn)
	}
	if len(_databrewEncryptionMode) > 0 {
		if err := assignInputField(input, "EncryptionMode", _databrewEncryptionMode); err != nil {
			log.Errorf("invalid --encryption-mode: %s", err.Error())
			return
		}
	}
	if len(_databrewJobSample) > 0 {
		if err := assignInputField(input, "JobSample", _databrewJobSample); err != nil {
			log.Errorf("invalid --job-sample: %s", err.Error())
			return
		}
	}
	if len(_databrewLogSubscription) > 0 {
		if err := assignInputField(input, "LogSubscription", _databrewLogSubscription); err != nil {
			log.Errorf("invalid --log-subscription: %s", err.Error())
			return
		}
	}
	if len(_databrewMaxCapacity) > 0 {
		if err := assignInputField(input, "MaxCapacity", _databrewMaxCapacity); err != nil {
			log.Errorf("invalid --max-capacity: %s", err.Error())
			return
		}
	}
	if len(_databrewMaxRetries) > 0 {
		if err := assignInputField(input, "MaxRetries", _databrewMaxRetries); err != nil {
			log.Errorf("invalid --max-retries: %s", err.Error())
			return
		}
	}
	if len(_databrewTags) > 0 {
		if err := assignInputField(input, "Tags", _databrewTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_databrewTimeout) > 0 {
		if err := assignInputField(input, "Timeout", _databrewTimeout); err != nil {
			log.Errorf("invalid --timeout: %s", err.Error())
			return
		}
	}
	if len(_databrewValidationConfigurations) > 0 {
		if err := assignInputField(input, "ValidationConfigurations", _databrewValidationConfigurations); err != nil {
			log.Errorf("invalid --validation-configurations: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProfileJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new DataBrew project.
func databrew_CreateProject(cfg aws.Config, client *databrew.Client) {
	input := &databrew.CreateProjectInput{
		// DatasetName: *string, // Required
		// Name: *string, // Required
		// RecipeName: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_databrewDatasetName) > 0 {
		input.DatasetName = aws.String(_databrewDatasetName)
	}
	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}
	if len(_databrewRecipeName) > 0 {
		input.RecipeName = aws.String(_databrewRecipeName)
	}
	if len(_databrewRoleArn) > 0 {
		input.RoleArn = aws.String(_databrewRoleArn)
	}
	if len(_databrewSample) > 0 {
		if err := assignInputField(input, "Sample", _databrewSample); err != nil {
			log.Errorf("invalid --sample: %s", err.Error())
			return
		}
	}
	if len(_databrewTags) > 0 {
		if err := assignInputField(input, "Tags", _databrewTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new DataBrew recipe.
func databrew_CreateRecipe(cfg aws.Config, client *databrew.Client) {
	input := &databrew.CreateRecipeInput{
		// Name: *string, // Required
		// Steps: []types.RecipeStep, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}
	if len(_databrewSteps) > 0 {
		if err := assignInputField(input, "Steps", _databrewSteps); err != nil {
			log.Errorf("invalid --steps: %s", err.Error())
			return
		}
	}
	if len(_databrewDescription) > 0 {
		input.Description = aws.String(_databrewDescription)
	}
	if len(_databrewTags) > 0 {
		if err := assignInputField(input, "Tags", _databrewTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRecipe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new job to transform input data, using steps defined in an existing
// Glue DataBrew recipe
func databrew_CreateRecipeJob(cfg aws.Config, client *databrew.Client) {
	input := &databrew.CreateRecipeJobInput{
		// Name: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}
	if len(_databrewRoleArn) > 0 {
		input.RoleArn = aws.String(_databrewRoleArn)
	}
	if len(_databrewDataCatalogOutputs) > 0 {
		if err := assignInputField(input, "DataCatalogOutputs", _databrewDataCatalogOutputs); err != nil {
			log.Errorf("invalid --data-catalog-outputs: %s", err.Error())
			return
		}
	}
	if len(_databrewDatabaseOutputs) > 0 {
		if err := assignInputField(input, "DatabaseOutputs", _databrewDatabaseOutputs); err != nil {
			log.Errorf("invalid --database-outputs: %s", err.Error())
			return
		}
	}
	if len(_databrewDatasetName) > 0 {
		input.DatasetName = aws.String(_databrewDatasetName)
	}
	if len(_databrewEncryptionKeyArn) > 0 {
		input.EncryptionKeyArn = aws.String(_databrewEncryptionKeyArn)
	}
	if len(_databrewEncryptionMode) > 0 {
		if err := assignInputField(input, "EncryptionMode", _databrewEncryptionMode); err != nil {
			log.Errorf("invalid --encryption-mode: %s", err.Error())
			return
		}
	}
	if len(_databrewLogSubscription) > 0 {
		if err := assignInputField(input, "LogSubscription", _databrewLogSubscription); err != nil {
			log.Errorf("invalid --log-subscription: %s", err.Error())
			return
		}
	}
	if len(_databrewMaxCapacity) > 0 {
		if err := assignInputField(input, "MaxCapacity", _databrewMaxCapacity); err != nil {
			log.Errorf("invalid --max-capacity: %s", err.Error())
			return
		}
	}
	if len(_databrewMaxRetries) > 0 {
		if err := assignInputField(input, "MaxRetries", _databrewMaxRetries); err != nil {
			log.Errorf("invalid --max-retries: %s", err.Error())
			return
		}
	}
	if len(_databrewOutputs) > 0 {
		if err := assignInputField(input, "Outputs", _databrewOutputs); err != nil {
			log.Errorf("invalid --outputs: %s", err.Error())
			return
		}
	}
	if len(_databrewProjectName) > 0 {
		input.ProjectName = aws.String(_databrewProjectName)
	}
	if len(_databrewRecipeReference) > 0 {
		if err := assignInputField(input, "RecipeReference", _databrewRecipeReference); err != nil {
			log.Errorf("invalid --recipe-reference: %s", err.Error())
			return
		}
	}
	if len(_databrewTags) > 0 {
		if err := assignInputField(input, "Tags", _databrewTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_databrewTimeout) > 0 {
		if err := assignInputField(input, "Timeout", _databrewTimeout); err != nil {
			log.Errorf("invalid --timeout: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRecipeJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new ruleset that can be used in a profile job to validate the data
// quality of a dataset.
func databrew_CreateRuleset(cfg aws.Config, client *databrew.Client) {
	input := &databrew.CreateRulesetInput{
		// Name: *string, // Required
		// Rules: []types.Rule, // Required
		// TargetArn: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}
	if len(_databrewRules) > 0 {
		if err := assignInputField(input, "Rules", _databrewRules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_databrewTargetArn) > 0 {
		input.TargetArn = aws.String(_databrewTargetArn)
	}
	if len(_databrewDescription) > 0 {
		input.Description = aws.String(_databrewDescription)
	}
	if len(_databrewTags) > 0 {
		if err := assignInputField(input, "Tags", _databrewTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRuleset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new schedule for one or more DataBrew jobs. Jobs can be run at a
// specific date and time, or at regular intervals.
func databrew_CreateSchedule(cfg aws.Config, client *databrew.Client) {
	input := &databrew.CreateScheduleInput{
		// CronExpression: *string, // Required
		// Name: *string, // Required
	}

	if len(_databrewCronExpression) > 0 {
		input.CronExpression = aws.String(_databrewCronExpression)
	}
	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}
	if len(_databrewJobNames) > 0 {
		input.JobNames = append([]string(nil), _databrewJobNames...)
	}
	if len(_databrewTags) > 0 {
		if err := assignInputField(input, "Tags", _databrewTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a dataset from DataBrew.
func databrew_DeleteDataset(cfg aws.Config, client *databrew.Client) {
	input := &databrew.DeleteDatasetInput{
		// Name: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}

	if resp, err := client.DeleteDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified DataBrew job.
func databrew_DeleteJob(cfg aws.Config, client *databrew.Client) {
	input := &databrew.DeleteJobInput{
		// Name: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}

	if resp, err := client.DeleteJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing DataBrew project.
func databrew_DeleteProject(cfg aws.Config, client *databrew.Client) {
	input := &databrew.DeleteProjectInput{
		// Name: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}

	if resp, err := client.DeleteProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a single version of a DataBrew recipe.
func databrew_DeleteRecipeVersion(cfg aws.Config, client *databrew.Client) {
	input := &databrew.DeleteRecipeVersionInput{
		// Name: *string, // Required
		// RecipeVersion: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}
	if len(_databrewRecipeVersion) > 0 {
		input.RecipeVersion = aws.String(_databrewRecipeVersion)
	}

	if resp, err := client.DeleteRecipeVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a ruleset.
func databrew_DeleteRuleset(cfg aws.Config, client *databrew.Client) {
	input := &databrew.DeleteRulesetInput{
		// Name: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}

	if resp, err := client.DeleteRuleset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified DataBrew schedule.
func databrew_DeleteSchedule(cfg aws.Config, client *databrew.Client) {
	input := &databrew.DeleteScheduleInput{
		// Name: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}

	if resp, err := client.DeleteSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the definition of a specific DataBrew dataset.
func databrew_DescribeDataset(cfg aws.Config, client *databrew.Client) {
	input := &databrew.DescribeDatasetInput{
		// Name: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}

	if resp, err := client.DescribeDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the definition of a specific DataBrew job.
func databrew_DescribeJob(cfg aws.Config, client *databrew.Client) {
	input := &databrew.DescribeJobInput{
		// Name: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}

	if resp, err := client.DescribeJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Represents one run of a DataBrew job.
func databrew_DescribeJobRun(cfg aws.Config, client *databrew.Client) {
	input := &databrew.DescribeJobRunInput{
		// Name: *string, // Required
		// RunId: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}
	if len(_databrewRunId) > 0 {
		input.RunId = aws.String(_databrewRunId)
	}

	if resp, err := client.DescribeJobRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the definition of a specific DataBrew project.
func databrew_DescribeProject(cfg aws.Config, client *databrew.Client) {
	input := &databrew.DescribeProjectInput{
		// Name: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}

	if resp, err := client.DescribeProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the definition of a specific DataBrew recipe corresponding to a
// particular version.
func databrew_DescribeRecipe(cfg aws.Config, client *databrew.Client) {
	input := &databrew.DescribeRecipeInput{
		// Name: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}
	if len(_databrewRecipeVersion) > 0 {
		input.RecipeVersion = aws.String(_databrewRecipeVersion)
	}

	if resp, err := client.DescribeRecipe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about the ruleset.
func databrew_DescribeRuleset(cfg aws.Config, client *databrew.Client) {
	input := &databrew.DescribeRulesetInput{
		// Name: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}

	if resp, err := client.DescribeRuleset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the definition of a specific DataBrew schedule.
func databrew_DescribeSchedule(cfg aws.Config, client *databrew.Client) {
	input := &databrew.DescribeScheduleInput{
		// Name: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}

	if resp, err := client.DescribeSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all of the DataBrew datasets.
func databrew_ListDatasets(cfg aws.Config, client *databrew.Client) {
	input := &databrew.ListDatasetsInput{}

	if len(_databrewMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _databrewMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_databrewNextToken) > 0 {
		input.NextToken = aws.String(_databrewNextToken)
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

	var results []*databrew.ListDatasetsOutput
	p := databrew.NewListDatasetsPaginator(client, input)
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

// Lists all of the previous runs of a particular DataBrew job.
func databrew_ListJobRuns(cfg aws.Config, client *databrew.Client) {
	input := &databrew.ListJobRunsInput{
		// Name: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}
	if len(_databrewMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _databrewMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_databrewNextToken) > 0 {
		input.NextToken = aws.String(_databrewNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListJobRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databrew.ListJobRunsOutput
	p := databrew.NewListJobRunsPaginator(client, input)
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

// Lists all of the DataBrew jobs that are defined.
func databrew_ListJobs(cfg aws.Config, client *databrew.Client) {
	input := &databrew.ListJobsInput{}

	if len(_databrewDatasetName) > 0 {
		input.DatasetName = aws.String(_databrewDatasetName)
	}
	if len(_databrewMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _databrewMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_databrewNextToken) > 0 {
		input.NextToken = aws.String(_databrewNextToken)
	}
	if len(_databrewProjectName) > 0 {
		input.ProjectName = aws.String(_databrewProjectName)
	}

	if disablePaginator() {
		if resp, err := client.ListJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databrew.ListJobsOutput
	p := databrew.NewListJobsPaginator(client, input)
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

// Lists all of the DataBrew projects that are defined.
func databrew_ListProjects(cfg aws.Config, client *databrew.Client) {
	input := &databrew.ListProjectsInput{}

	if len(_databrewMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _databrewMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_databrewNextToken) > 0 {
		input.NextToken = aws.String(_databrewNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProjects(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databrew.ListProjectsOutput
	p := databrew.NewListProjectsPaginator(client, input)
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

// Lists the versions of a particular DataBrew recipe, except for LATEST_WORKING .
func databrew_ListRecipeVersions(cfg aws.Config, client *databrew.Client) {
	input := &databrew.ListRecipeVersionsInput{
		// Name: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}
	if len(_databrewMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _databrewMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_databrewNextToken) > 0 {
		input.NextToken = aws.String(_databrewNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRecipeVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databrew.ListRecipeVersionsOutput
	p := databrew.NewListRecipeVersionsPaginator(client, input)
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

// Lists all of the DataBrew recipes that are defined.
func databrew_ListRecipes(cfg aws.Config, client *databrew.Client) {
	input := &databrew.ListRecipesInput{}

	if len(_databrewMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _databrewMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_databrewNextToken) > 0 {
		input.NextToken = aws.String(_databrewNextToken)
	}
	if len(_databrewRecipeVersion) > 0 {
		input.RecipeVersion = aws.String(_databrewRecipeVersion)
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

	var results []*databrew.ListRecipesOutput
	p := databrew.NewListRecipesPaginator(client, input)
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

// List all rulesets available in the current account or rulesets associated with
// a specific resource (dataset).
func databrew_ListRulesets(cfg aws.Config, client *databrew.Client) {
	input := &databrew.ListRulesetsInput{}

	if len(_databrewMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _databrewMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_databrewNextToken) > 0 {
		input.NextToken = aws.String(_databrewNextToken)
	}
	if len(_databrewTargetArn) > 0 {
		input.TargetArn = aws.String(_databrewTargetArn)
	}

	if disablePaginator() {
		if resp, err := client.ListRulesets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databrew.ListRulesetsOutput
	p := databrew.NewListRulesetsPaginator(client, input)
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

// Lists the DataBrew schedules that are defined.
func databrew_ListSchedules(cfg aws.Config, client *databrew.Client) {
	input := &databrew.ListSchedulesInput{}

	if len(_databrewJobName) > 0 {
		input.JobName = aws.String(_databrewJobName)
	}
	if len(_databrewMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _databrewMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_databrewNextToken) > 0 {
		input.NextToken = aws.String(_databrewNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSchedules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databrew.ListSchedulesOutput
	p := databrew.NewListSchedulesPaginator(client, input)
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

// Lists all the tags for a DataBrew resource.
func databrew_ListTagsForResource(cfg aws.Config, client *databrew.Client) {
	input := &databrew.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_databrewResourceArn) > 0 {
		input.ResourceArn = aws.String(_databrewResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Publishes a new version of a DataBrew recipe.
func databrew_PublishRecipe(cfg aws.Config, client *databrew.Client) {
	input := &databrew.PublishRecipeInput{
		// Name: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}
	if len(_databrewDescription) > 0 {
		input.Description = aws.String(_databrewDescription)
	}

	if resp, err := client.PublishRecipe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Performs a recipe step within an interactive DataBrew session that's currently
// open.
func databrew_SendProjectSessionAction(cfg aws.Config, client *databrew.Client) {
	input := &databrew.SendProjectSessionActionInput{
		// Name: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}
	if len(_databrewClientSessionId) > 0 {
		input.ClientSessionId = aws.String(_databrewClientSessionId)
	}
	if len(_databrewPreview) > 0 {
		if err := assignInputField(input, "Preview", _databrewPreview); err != nil {
			log.Errorf("invalid --preview: %s", err.Error())
			return
		}
	}
	if len(_databrewRecipeStep) > 0 {
		if err := assignInputField(input, "RecipeStep", _databrewRecipeStep); err != nil {
			log.Errorf("invalid --recipe-step: %s", err.Error())
			return
		}
	}
	if len(_databrewStepIndex) > 0 {
		if err := assignInputField(input, "StepIndex", _databrewStepIndex); err != nil {
			log.Errorf("invalid --step-index: %s", err.Error())
			return
		}
	}
	if len(_databrewViewFrame) > 0 {
		if err := assignInputField(input, "ViewFrame", _databrewViewFrame); err != nil {
			log.Errorf("invalid --view-frame: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendProjectSessionAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Runs a DataBrew job.
func databrew_StartJobRun(cfg aws.Config, client *databrew.Client) {
	input := &databrew.StartJobRunInput{
		// Name: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}

	if resp, err := client.StartJobRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an interactive session, enabling you to manipulate data in a DataBrew
// project.
func databrew_StartProjectSession(cfg aws.Config, client *databrew.Client) {
	input := &databrew.StartProjectSessionInput{
		// Name: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}
	if len(_databrewAssumeControl) > 0 {
		if err := assignInputField(input, "AssumeControl", _databrewAssumeControl); err != nil {
			log.Errorf("invalid --assume-control: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartProjectSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a particular run of a job.
func databrew_StopJobRun(cfg aws.Config, client *databrew.Client) {
	input := &databrew.StopJobRunInput{
		// Name: *string, // Required
		// RunId: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}
	if len(_databrewRunId) > 0 {
		input.RunId = aws.String(_databrewRunId)
	}

	if resp, err := client.StopJobRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds metadata tags to a DataBrew resource, such as a dataset, project, recipe,
// job, or schedule.
func databrew_TagResource(cfg aws.Config, client *databrew.Client) {
	input := &databrew.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_databrewResourceArn) > 0 {
		input.ResourceArn = aws.String(_databrewResourceArn)
	}
	if len(_databrewTags) > 0 {
		if err := assignInputField(input, "Tags", _databrewTags); err != nil {
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

// Removes metadata tags from a DataBrew resource.
func databrew_UntagResource(cfg aws.Config, client *databrew.Client) {
	input := &databrew.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_databrewResourceArn) > 0 {
		input.ResourceArn = aws.String(_databrewResourceArn)
	}
	if len(_databrewTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _databrewTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the definition of an existing DataBrew dataset.
func databrew_UpdateDataset(cfg aws.Config, client *databrew.Client) {
	input := &databrew.UpdateDatasetInput{
		// Input: *types.Input, // Required
		// Name: *string, // Required
	}

	if len(_databrewInput) > 0 {
		if err := assignInputField(input, "Input", _databrewInput); err != nil {
			log.Errorf("invalid --input: %s", err.Error())
			return
		}
	}
	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}
	if len(_databrewFormat) > 0 {
		if err := assignInputField(input, "Format", _databrewFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_databrewFormatOptions) > 0 {
		if err := assignInputField(input, "FormatOptions", _databrewFormatOptions); err != nil {
			log.Errorf("invalid --format-options: %s", err.Error())
			return
		}
	}
	if len(_databrewPathOptions) > 0 {
		if err := assignInputField(input, "PathOptions", _databrewPathOptions); err != nil {
			log.Errorf("invalid --path-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the definition of an existing profile job.
func databrew_UpdateProfileJob(cfg aws.Config, client *databrew.Client) {
	input := &databrew.UpdateProfileJobInput{
		// Name: *string, // Required
		// OutputLocation: *types.S3Location, // Required
		// RoleArn: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}
	if len(_databrewOutputLocation) > 0 {
		if err := assignInputField(input, "OutputLocation", _databrewOutputLocation); err != nil {
			log.Errorf("invalid --output-location: %s", err.Error())
			return
		}
	}
	if len(_databrewRoleArn) > 0 {
		input.RoleArn = aws.String(_databrewRoleArn)
	}
	if len(_databrewConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _databrewConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_databrewEncryptionKeyArn) > 0 {
		input.EncryptionKeyArn = aws.String(_databrewEncryptionKeyArn)
	}
	if len(_databrewEncryptionMode) > 0 {
		if err := assignInputField(input, "EncryptionMode", _databrewEncryptionMode); err != nil {
			log.Errorf("invalid --encryption-mode: %s", err.Error())
			return
		}
	}
	if len(_databrewJobSample) > 0 {
		if err := assignInputField(input, "JobSample", _databrewJobSample); err != nil {
			log.Errorf("invalid --job-sample: %s", err.Error())
			return
		}
	}
	if len(_databrewLogSubscription) > 0 {
		if err := assignInputField(input, "LogSubscription", _databrewLogSubscription); err != nil {
			log.Errorf("invalid --log-subscription: %s", err.Error())
			return
		}
	}
	if len(_databrewMaxCapacity) > 0 {
		if err := assignInputField(input, "MaxCapacity", _databrewMaxCapacity); err != nil {
			log.Errorf("invalid --max-capacity: %s", err.Error())
			return
		}
	}
	if len(_databrewMaxRetries) > 0 {
		if err := assignInputField(input, "MaxRetries", _databrewMaxRetries); err != nil {
			log.Errorf("invalid --max-retries: %s", err.Error())
			return
		}
	}
	if len(_databrewTimeout) > 0 {
		if err := assignInputField(input, "Timeout", _databrewTimeout); err != nil {
			log.Errorf("invalid --timeout: %s", err.Error())
			return
		}
	}
	if len(_databrewValidationConfigurations) > 0 {
		if err := assignInputField(input, "ValidationConfigurations", _databrewValidationConfigurations); err != nil {
			log.Errorf("invalid --validation-configurations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateProfileJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the definition of an existing DataBrew project.
func databrew_UpdateProject(cfg aws.Config, client *databrew.Client) {
	input := &databrew.UpdateProjectInput{
		// Name: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}
	if len(_databrewRoleArn) > 0 {
		input.RoleArn = aws.String(_databrewRoleArn)
	}
	if len(_databrewSample) > 0 {
		if err := assignInputField(input, "Sample", _databrewSample); err != nil {
			log.Errorf("invalid --sample: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the definition of the LATEST_WORKING version of a DataBrew recipe.
func databrew_UpdateRecipe(cfg aws.Config, client *databrew.Client) {
	input := &databrew.UpdateRecipeInput{
		// Name: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}
	if len(_databrewDescription) > 0 {
		input.Description = aws.String(_databrewDescription)
	}
	if len(_databrewSteps) > 0 {
		if err := assignInputField(input, "Steps", _databrewSteps); err != nil {
			log.Errorf("invalid --steps: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRecipe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the definition of an existing DataBrew recipe job.
func databrew_UpdateRecipeJob(cfg aws.Config, client *databrew.Client) {
	input := &databrew.UpdateRecipeJobInput{
		// Name: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}
	if len(_databrewRoleArn) > 0 {
		input.RoleArn = aws.String(_databrewRoleArn)
	}
	if len(_databrewDataCatalogOutputs) > 0 {
		if err := assignInputField(input, "DataCatalogOutputs", _databrewDataCatalogOutputs); err != nil {
			log.Errorf("invalid --data-catalog-outputs: %s", err.Error())
			return
		}
	}
	if len(_databrewDatabaseOutputs) > 0 {
		if err := assignInputField(input, "DatabaseOutputs", _databrewDatabaseOutputs); err != nil {
			log.Errorf("invalid --database-outputs: %s", err.Error())
			return
		}
	}
	if len(_databrewEncryptionKeyArn) > 0 {
		input.EncryptionKeyArn = aws.String(_databrewEncryptionKeyArn)
	}
	if len(_databrewEncryptionMode) > 0 {
		if err := assignInputField(input, "EncryptionMode", _databrewEncryptionMode); err != nil {
			log.Errorf("invalid --encryption-mode: %s", err.Error())
			return
		}
	}
	if len(_databrewLogSubscription) > 0 {
		if err := assignInputField(input, "LogSubscription", _databrewLogSubscription); err != nil {
			log.Errorf("invalid --log-subscription: %s", err.Error())
			return
		}
	}
	if len(_databrewMaxCapacity) > 0 {
		if err := assignInputField(input, "MaxCapacity", _databrewMaxCapacity); err != nil {
			log.Errorf("invalid --max-capacity: %s", err.Error())
			return
		}
	}
	if len(_databrewMaxRetries) > 0 {
		if err := assignInputField(input, "MaxRetries", _databrewMaxRetries); err != nil {
			log.Errorf("invalid --max-retries: %s", err.Error())
			return
		}
	}
	if len(_databrewOutputs) > 0 {
		if err := assignInputField(input, "Outputs", _databrewOutputs); err != nil {
			log.Errorf("invalid --outputs: %s", err.Error())
			return
		}
	}
	if len(_databrewTimeout) > 0 {
		if err := assignInputField(input, "Timeout", _databrewTimeout); err != nil {
			log.Errorf("invalid --timeout: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRecipeJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates specified ruleset.
func databrew_UpdateRuleset(cfg aws.Config, client *databrew.Client) {
	input := &databrew.UpdateRulesetInput{
		// Name: *string, // Required
		// Rules: []types.Rule, // Required
	}

	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}
	if len(_databrewRules) > 0 {
		if err := assignInputField(input, "Rules", _databrewRules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_databrewDescription) > 0 {
		input.Description = aws.String(_databrewDescription)
	}

	if resp, err := client.UpdateRuleset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the definition of an existing DataBrew schedule.
func databrew_UpdateSchedule(cfg aws.Config, client *databrew.Client) {
	input := &databrew.UpdateScheduleInput{
		// CronExpression: *string, // Required
		// Name: *string, // Required
	}

	if len(_databrewCronExpression) > 0 {
		input.CronExpression = aws.String(_databrewCronExpression)
	}
	if len(_databrewName) > 0 {
		input.Name = aws.String(_databrewName)
	}
	if len(_databrewJobNames) > 0 {
		input.JobNames = append([]string(nil), _databrewJobNames...)
	}

	if resp, err := client.UpdateSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_databrewCmd)
	_databrewCmd.Flags().SortFlags = false

	_databrewCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_databrewCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_databrewCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_databrewCmd.Flags().StringVarP(&_databrewAssumeControl, "assume-control", "", "", "Assume Control")
	_databrewCmd.Flags().StringVarP(&_databrewClientSessionId, "client-session-id", "", "", "Client Session ID")
	_databrewCmd.Flags().StringVarP(&_databrewConfiguration, "configuration", "", "", "Configuration")
	_databrewCmd.Flags().StringVarP(&_databrewCronExpression, "cron-expression", "", "", "Cron Expression")
	_databrewCmd.Flags().StringVarP(&_databrewDataCatalogOutputs, "data-catalog-outputs", "", "", "Data Catalog Outputs")
	_databrewCmd.Flags().StringVarP(&_databrewDatabaseOutputs, "database-outputs", "", "", "Database Outputs")
	_databrewCmd.Flags().StringVarP(&_databrewDatasetName, "dataset-name", "", "", "Dataset Name")
	_databrewCmd.Flags().StringVarP(&_databrewDescription, "description", "", "", "Description")
	_databrewCmd.Flags().StringVarP(&_databrewEncryptionKeyArn, "encryption-key-arn", "", "", "Encryption Key ARN")
	_databrewCmd.Flags().StringVarP(&_databrewEncryptionMode, "encryption-mode", "", "", "Encryption Mode")
	_databrewCmd.Flags().StringVarP(&_databrewFormat, "format", "", "", "Format")
	_databrewCmd.Flags().StringVarP(&_databrewFormatOptions, "format-options", "", "", "Format Options")
	_databrewCmd.Flags().StringVarP(&_databrewInput, "input", "", "", "Input")
	_databrewCmd.Flags().StringVarP(&_databrewJobName, "job-name", "", "", "Job Name")
	_databrewCmd.Flags().StringSliceVarP(&_databrewJobNames, "job-names", "", nil, "Job Names")
	_databrewCmd.Flags().StringVarP(&_databrewJobSample, "job-sample", "", "", "Job Sample")
	_databrewCmd.Flags().StringVarP(&_databrewLogSubscription, "log-subscription", "", "", "Log Subscription")
	_databrewCmd.Flags().StringVarP(&_databrewMaxCapacity, "max-capacity", "", "", "Max Capacity")
	_databrewCmd.Flags().StringVarP(&_databrewMaxResults, "max-results", "", "", "Max Results")
	_databrewCmd.Flags().StringVarP(&_databrewMaxRetries, "max-retries", "", "", "Max Retries")
	_databrewCmd.Flags().StringVarP(&_databrewName, "name", "", "", "Name")
	_databrewCmd.Flags().StringVarP(&_databrewNextToken, "next-token", "", "", "Next Token")
	_databrewCmd.Flags().StringVarP(&_databrewOutputLocation, "output-location", "", "", "Output Location")
	_databrewCmd.Flags().StringVarP(&_databrewOutputs, "outputs", "", "", "Outputs")
	_databrewCmd.Flags().StringVarP(&_databrewPathOptions, "path-options", "", "", "Path Options")
	_databrewCmd.Flags().StringVarP(&_databrewPreview, "preview", "", "", "Preview")
	_databrewCmd.Flags().StringVarP(&_databrewProjectName, "project-name", "", "", "Project Name")
	_databrewCmd.Flags().StringVarP(&_databrewRecipeName, "recipe-name", "", "", "Recipe Name")
	_databrewCmd.Flags().StringVarP(&_databrewRecipeReference, "recipe-reference", "", "", "Recipe Reference")
	_databrewCmd.Flags().StringVarP(&_databrewRecipeStep, "recipe-step", "", "", "Recipe Step")
	_databrewCmd.Flags().StringVarP(&_databrewRecipeVersion, "recipe-version", "", "", "Recipe Version")
	_databrewCmd.Flags().StringSliceVarP(&_databrewRecipeVersions, "recipe-versions", "", nil, "Recipe Versions")
	_databrewCmd.Flags().StringVarP(&_databrewResourceArn, "resource-arn", "", "", "Resource ARN")
	_databrewCmd.Flags().StringVarP(&_databrewRoleArn, "role-arn", "", "", "Role ARN")
	_databrewCmd.Flags().StringVarP(&_databrewRules, "rules", "", "", "Rules")
	_databrewCmd.Flags().StringVarP(&_databrewRunId, "run-id", "", "", "Run ID")
	_databrewCmd.Flags().StringVarP(&_databrewSample, "sample", "", "", "Sample")
	_databrewCmd.Flags().StringVarP(&_databrewStepIndex, "step-index", "", "", "Step Index")
	_databrewCmd.Flags().StringVarP(&_databrewSteps, "steps", "", "", "Steps")
	_databrewCmd.Flags().StringSliceVarP(&_databrewTagKeys, "tag-keys", "", nil, "Tag Keys")
	_databrewCmd.Flags().StringVarP(&_databrewTags, "tags", "", "", "Tags")
	_databrewCmd.Flags().StringVarP(&_databrewTargetArn, "target-arn", "", "", "Target ARN")
	_databrewCmd.Flags().StringVarP(&_databrewTimeout, "timeout", "", "", "Timeout")
	_databrewCmd.Flags().StringVarP(&_databrewValidationConfigurations, "validation-configurations", "", "", "Validation Configurations")
	_databrewCmd.Flags().StringVarP(&_databrewViewFrame, "view-frame", "", "", "View Frame")

	_databrewCmd.Flags().BoolVarP(&_databrewBatchDeleteRecipeVersion, "batch-delete-recipe-version", "", false, "Batch Delete Recipe Version")
	_databrewCmd.Flags().BoolVarP(&_databrewCreateDataset, "create-dataset", "", false, "Create Dataset")
	_databrewCmd.Flags().BoolVarP(&_databrewCreateProfileJob, "create-profile-job", "", false, "Create Profile Job")
	_databrewCmd.Flags().BoolVarP(&_databrewCreateProject, "create-project", "", false, "Create Project")
	_databrewCmd.Flags().BoolVarP(&_databrewCreateRecipe, "create-recipe", "", false, "Create Recipe")
	_databrewCmd.Flags().BoolVarP(&_databrewCreateRecipeJob, "create-recipe-job", "", false, "Create Recipe Job")
	_databrewCmd.Flags().BoolVarP(&_databrewCreateRuleset, "create-ruleset", "", false, "Create Ruleset")
	_databrewCmd.Flags().BoolVarP(&_databrewCreateSchedule, "create-schedule", "", false, "Create Schedule")
	_databrewCmd.Flags().BoolVarP(&_databrewDeleteDataset, "delete-dataset", "", false, "Delete Dataset")
	_databrewCmd.Flags().BoolVarP(&_databrewDeleteJob, "delete-job", "", false, "Delete Job")
	_databrewCmd.Flags().BoolVarP(&_databrewDeleteProject, "delete-project", "", false, "Delete Project")
	_databrewCmd.Flags().BoolVarP(&_databrewDeleteRecipeVersion, "delete-recipe-version", "", false, "Delete Recipe Version")
	_databrewCmd.Flags().BoolVarP(&_databrewDeleteRuleset, "delete-ruleset", "", false, "Delete Ruleset")
	_databrewCmd.Flags().BoolVarP(&_databrewDeleteSchedule, "delete-schedule", "", false, "Delete Schedule")
	_databrewCmd.Flags().BoolVarP(&_databrewDescribeDataset, "describe-dataset", "", false, "Describe Dataset")
	_databrewCmd.Flags().BoolVarP(&_databrewDescribeJob, "describe-job", "", false, "Describe Job")
	_databrewCmd.Flags().BoolVarP(&_databrewDescribeJobRun, "describe-job-run", "", false, "Describe Job Run")
	_databrewCmd.Flags().BoolVarP(&_databrewDescribeProject, "describe-project", "", false, "Describe Project")
	_databrewCmd.Flags().BoolVarP(&_databrewDescribeRecipe, "describe-recipe", "", false, "Describe Recipe")
	_databrewCmd.Flags().BoolVarP(&_databrewDescribeRuleset, "describe-ruleset", "", false, "Describe Ruleset")
	_databrewCmd.Flags().BoolVarP(&_databrewDescribeSchedule, "describe-schedule", "", false, "Describe Schedule")
	_databrewCmd.Flags().BoolVarP(&_databrewListDatasets, "list-datasets", "", false, "List Datasets")
	_databrewCmd.Flags().BoolVarP(&_databrewListJobRuns, "list-job-runs", "", false, "List Job Runs")
	_databrewCmd.Flags().BoolVarP(&_databrewListJobs, "list-jobs", "", false, "List Jobs")
	_databrewCmd.Flags().BoolVarP(&_databrewListProjects, "list-projects", "", false, "List Projects")
	_databrewCmd.Flags().BoolVarP(&_databrewListRecipeVersions, "list-recipe-versions", "", false, "List Recipe Versions")
	_databrewCmd.Flags().BoolVarP(&_databrewListRecipes, "list-recipes", "", false, "List Recipes")
	_databrewCmd.Flags().BoolVarP(&_databrewListRulesets, "list-rulesets", "", false, "List Rulesets")
	_databrewCmd.Flags().BoolVarP(&_databrewListSchedules, "list-schedules", "", false, "List Schedules")
	_databrewCmd.Flags().BoolVarP(&_databrewListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_databrewCmd.Flags().BoolVarP(&_databrewPublishRecipe, "publish-recipe", "", false, "Publish Recipe")
	_databrewCmd.Flags().BoolVarP(&_databrewSendProjectSessionAction, "send-project-session-action", "", false, "Send Project Session Action")
	_databrewCmd.Flags().BoolVarP(&_databrewStartJobRun, "start-job-run", "", false, "Start Job Run")
	_databrewCmd.Flags().BoolVarP(&_databrewStartProjectSession, "start-project-session", "", false, "Start Project Session")
	_databrewCmd.Flags().BoolVarP(&_databrewStopJobRun, "stop-job-run", "", false, "Stop Job Run")
	_databrewCmd.Flags().BoolVarP(&_databrewTagResource, "tag-resource", "", false, "Tag Resource")
	_databrewCmd.Flags().BoolVarP(&_databrewUntagResource, "untag-resource", "", false, "Untag Resource")
	_databrewCmd.Flags().BoolVarP(&_databrewUpdateDataset, "update-dataset", "", false, "Update Dataset")
	_databrewCmd.Flags().BoolVarP(&_databrewUpdateProfileJob, "update-profile-job", "", false, "Update Profile Job")
	_databrewCmd.Flags().BoolVarP(&_databrewUpdateProject, "update-project", "", false, "Update Project")
	_databrewCmd.Flags().BoolVarP(&_databrewUpdateRecipe, "update-recipe", "", false, "Update Recipe")
	_databrewCmd.Flags().BoolVarP(&_databrewUpdateRecipeJob, "update-recipe-job", "", false, "Update Recipe Job")
	_databrewCmd.Flags().BoolVarP(&_databrewUpdateRuleset, "update-ruleset", "", false, "Update Ruleset")
	_databrewCmd.Flags().BoolVarP(&_databrewUpdateSchedule, "update-schedule", "", false, "Update Schedule")

}
