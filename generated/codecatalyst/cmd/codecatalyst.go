package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codecatalyst"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// codecatalystCmd represents the codecatalyst command
var _codecatalystCmd = &cobra.Command{
	Use:   "codecatalyst",
	Short: "AWS codecatalyst CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := codecatalyst.NewFromConfig(cfg)
		if _codecatalystCreateAccessToken {
			codecatalyst_CreateAccessToken(cfg, client)
			return
		}
		if _codecatalystCreateDevEnvironment {
			codecatalyst_CreateDevEnvironment(cfg, client)
			return
		}
		if _codecatalystCreateProject {
			codecatalyst_CreateProject(cfg, client)
			return
		}
		if _codecatalystCreateSourceRepository {
			codecatalyst_CreateSourceRepository(cfg, client)
			return
		}
		if _codecatalystCreateSourceRepositoryBranch {
			codecatalyst_CreateSourceRepositoryBranch(cfg, client)
			return
		}
		if _codecatalystDeleteAccessToken {
			codecatalyst_DeleteAccessToken(cfg, client)
			return
		}
		if _codecatalystDeleteDevEnvironment {
			codecatalyst_DeleteDevEnvironment(cfg, client)
			return
		}
		if _codecatalystDeleteProject {
			codecatalyst_DeleteProject(cfg, client)
			return
		}
		if _codecatalystDeleteSourceRepository {
			codecatalyst_DeleteSourceRepository(cfg, client)
			return
		}
		if _codecatalystDeleteSpace {
			codecatalyst_DeleteSpace(cfg, client)
			return
		}
		if _codecatalystGetDevEnvironment {
			codecatalyst_GetDevEnvironment(cfg, client)
			return
		}
		if _codecatalystGetProject {
			codecatalyst_GetProject(cfg, client)
			return
		}
		if _codecatalystGetSourceRepository {
			codecatalyst_GetSourceRepository(cfg, client)
			return
		}
		if _codecatalystGetSourceRepositoryCloneUrls {
			codecatalyst_GetSourceRepositoryCloneUrls(cfg, client)
			return
		}
		if _codecatalystGetSpace {
			codecatalyst_GetSpace(cfg, client)
			return
		}
		if _codecatalystGetSubscription {
			codecatalyst_GetSubscription(cfg, client)
			return
		}
		if _codecatalystGetUserDetails {
			codecatalyst_GetUserDetails(cfg, client)
			return
		}
		if _codecatalystGetWorkflow {
			codecatalyst_GetWorkflow(cfg, client)
			return
		}
		if _codecatalystGetWorkflowRun {
			codecatalyst_GetWorkflowRun(cfg, client)
			return
		}
		if _codecatalystListAccessTokens {
			codecatalyst_ListAccessTokens(cfg, client)
			return
		}
		if _codecatalystListDevEnvironmentSessions {
			codecatalyst_ListDevEnvironmentSessions(cfg, client)
			return
		}
		if _codecatalystListDevEnvironments {
			codecatalyst_ListDevEnvironments(cfg, client)
			return
		}
		if _codecatalystListEventLogs {
			codecatalyst_ListEventLogs(cfg, client)
			return
		}
		if _codecatalystListProjects {
			codecatalyst_ListProjects(cfg, client)
			return
		}
		if _codecatalystListSourceRepositories {
			codecatalyst_ListSourceRepositories(cfg, client)
			return
		}
		if _codecatalystListSourceRepositoryBranches {
			codecatalyst_ListSourceRepositoryBranches(cfg, client)
			return
		}
		if _codecatalystListSpaces {
			codecatalyst_ListSpaces(cfg, client)
			return
		}
		if _codecatalystListWorkflowRuns {
			codecatalyst_ListWorkflowRuns(cfg, client)
			return
		}
		if _codecatalystListWorkflows {
			codecatalyst_ListWorkflows(cfg, client)
			return
		}
		if _codecatalystStartDevEnvironment {
			codecatalyst_StartDevEnvironment(cfg, client)
			return
		}
		if _codecatalystStartDevEnvironmentSession {
			codecatalyst_StartDevEnvironmentSession(cfg, client)
			return
		}
		if _codecatalystStartWorkflowRun {
			codecatalyst_StartWorkflowRun(cfg, client)
			return
		}
		if _codecatalystStopDevEnvironment {
			codecatalyst_StopDevEnvironment(cfg, client)
			return
		}
		if _codecatalystStopDevEnvironmentSession {
			codecatalyst_StopDevEnvironmentSession(cfg, client)
			return
		}
		if _codecatalystUpdateDevEnvironment {
			codecatalyst_UpdateDevEnvironment(cfg, client)
			return
		}
		if _codecatalystUpdateProject {
			codecatalyst_UpdateProject(cfg, client)
			return
		}
		if _codecatalystUpdateSpace {
			codecatalyst_UpdateSpace(cfg, client)
			return
		}
		if _codecatalystVerifySession {
			codecatalyst_VerifySession(cfg, client)
			return
		}

	},
}

var (
	_codecatalystCreateAccessToken            bool
	_codecatalystCreateDevEnvironment         bool
	_codecatalystCreateProject                bool
	_codecatalystCreateSourceRepository       bool
	_codecatalystCreateSourceRepositoryBranch bool
	_codecatalystDeleteAccessToken            bool
	_codecatalystDeleteDevEnvironment         bool
	_codecatalystDeleteProject                bool
	_codecatalystDeleteSourceRepository       bool
	_codecatalystDeleteSpace                  bool
	_codecatalystGetDevEnvironment            bool
	_codecatalystGetProject                   bool
	_codecatalystGetSourceRepository          bool
	_codecatalystGetSourceRepositoryCloneUrls bool
	_codecatalystGetSpace                     bool
	_codecatalystGetSubscription              bool
	_codecatalystGetUserDetails               bool
	_codecatalystGetWorkflow                  bool
	_codecatalystGetWorkflowRun               bool
	_codecatalystListAccessTokens             bool
	_codecatalystListDevEnvironmentSessions   bool
	_codecatalystListDevEnvironments          bool
	_codecatalystListEventLogs                bool
	_codecatalystListProjects                 bool
	_codecatalystListSourceRepositories       bool
	_codecatalystListSourceRepositoryBranches bool
	_codecatalystListSpaces                   bool
	_codecatalystListWorkflowRuns             bool
	_codecatalystListWorkflows                bool
	_codecatalystStartDevEnvironment          bool
	_codecatalystStartDevEnvironmentSession   bool
	_codecatalystStartWorkflowRun             bool
	_codecatalystStopDevEnvironment           bool
	_codecatalystStopDevEnvironmentSession    bool
	_codecatalystUpdateDevEnvironment         bool
	_codecatalystUpdateProject                bool
	_codecatalystUpdateSpace                  bool
	_codecatalystVerifySession                bool

	_codecatalystAlias                    string
	_codecatalystClientToken              string
	_codecatalystDescription              string
	_codecatalystDevEnvironmentId         string
	_codecatalystDisplayName              string
	_codecatalystEndTime                  string
	_codecatalystEventName                string
	_codecatalystExpiresTime              string
	_codecatalystFilters                  string
	_codecatalystHeadCommitId             string
	_codecatalystId                       string
	_codecatalystIdes                     string
	_codecatalystInactivityTimeoutMinutes string
	_codecatalystInstanceType             string
	_codecatalystMaxResults               string
	_codecatalystName                     string
	_codecatalystNextToken                string
	_codecatalystPersistentStorage        string
	_codecatalystProjectName              string
	_codecatalystRepositories             string
	_codecatalystSessionConfiguration     string
	_codecatalystSessionId                string
	_codecatalystSortBy                   string
	_codecatalystSourceRepositoryName     string
	_codecatalystSpaceName                string
	_codecatalystStartTime                string
	_codecatalystUserName                 string
	_codecatalystVpcConnectionName        string
	_codecatalystWorkflowId               string
)

// Creates a personal access token (PAT) for the current user. A personal access
// token (PAT) is similar to a password. It is associated with your user identity
// for use across all spaces and projects in Amazon CodeCatalyst. You use PATs to
// access CodeCatalyst from resources that include integrated development
// environments (IDEs) and Git-based source repositories. PATs represent you in
// Amazon CodeCatalyst and you can manage them in your user settings.For more
// information, see [Managing personal access tokens in Amazon CodeCatalyst].
//
// [Managing personal access tokens in Amazon CodeCatalyst]: https://docs.aws.amazon.com/codecatalyst/latest/userguide/ipa-tokens-keys.html
func codecatalyst_CreateAccessToken(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.CreateAccessTokenInput{
		// Name: *string, // Required
	}

	if len(_codecatalystName) > 0 {
		input.Name = aws.String(_codecatalystName)
	}
	if len(_codecatalystExpiresTime) > 0 {
		if err := assignInputField(input, "ExpiresTime", _codecatalystExpiresTime); err != nil {
			log.Errorf("invalid --expires-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAccessToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Dev Environment in Amazon CodeCatalyst, a cloud-based development
// environment that you can use to quickly work on the code stored in the source
// repositories of your project.
//
// When created in the Amazon CodeCatalyst console, by default a Dev Environment
// is configured to have a 2 core processor, 4GB of RAM, and 16GB of persistent
// storage. None of these defaults apply to a Dev Environment created
// programmatically.
func codecatalyst_CreateDevEnvironment(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.CreateDevEnvironmentInput{
		// InstanceType: types.InstanceType, // Required
		// PersistentStorage: *types.PersistentStorageConfiguration, // Required
		// ProjectName: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_codecatalystInstanceType) > 0 {
		if err := assignInputField(input, "InstanceType", _codecatalystInstanceType); err != nil {
			log.Errorf("invalid --instance-type: %s", err.Error())
			return
		}
	}
	if len(_codecatalystPersistentStorage) > 0 {
		if err := assignInputField(input, "PersistentStorage", _codecatalystPersistentStorage); err != nil {
			log.Errorf("invalid --persistent-storage: %s", err.Error())
			return
		}
	}
	if len(_codecatalystProjectName) > 0 {
		input.ProjectName = aws.String(_codecatalystProjectName)
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}
	if len(_codecatalystAlias) > 0 {
		input.Alias = aws.String(_codecatalystAlias)
	}
	if len(_codecatalystClientToken) > 0 {
		input.ClientToken = aws.String(_codecatalystClientToken)
	}
	if len(_codecatalystIdes) > 0 {
		if err := assignInputField(input, "Ides", _codecatalystIdes); err != nil {
			log.Errorf("invalid --ides: %s", err.Error())
			return
		}
	}
	if len(_codecatalystInactivityTimeoutMinutes) > 0 {
		if err := assignInputField(input, "InactivityTimeoutMinutes", _codecatalystInactivityTimeoutMinutes); err != nil {
			log.Errorf("invalid --inactivity-timeout-minutes: %s", err.Error())
			return
		}
	}
	if len(_codecatalystRepositories) > 0 {
		if err := assignInputField(input, "Repositories", _codecatalystRepositories); err != nil {
			log.Errorf("invalid --repositories: %s", err.Error())
			return
		}
	}
	if len(_codecatalystVpcConnectionName) > 0 {
		input.VpcConnectionName = aws.String(_codecatalystVpcConnectionName)
	}

	if resp, err := client.CreateDevEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a project in a specified space.
func codecatalyst_CreateProject(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.CreateProjectInput{
		// DisplayName: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_codecatalystDisplayName) > 0 {
		input.DisplayName = aws.String(_codecatalystDisplayName)
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}
	if len(_codecatalystDescription) > 0 {
		input.Description = aws.String(_codecatalystDescription)
	}

	if resp, err := client.CreateProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an empty Git-based source repository in a specified project. The
// repository is created with an initial empty commit with a default branch named
// main .
func codecatalyst_CreateSourceRepository(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.CreateSourceRepositoryInput{
		// Name: *string, // Required
		// ProjectName: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_codecatalystName) > 0 {
		input.Name = aws.String(_codecatalystName)
	}
	if len(_codecatalystProjectName) > 0 {
		input.ProjectName = aws.String(_codecatalystProjectName)
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}
	if len(_codecatalystDescription) > 0 {
		input.Description = aws.String(_codecatalystDescription)
	}

	if resp, err := client.CreateSourceRepository(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a branch in a specified source repository in Amazon CodeCatalyst.
// This API only creates a branch in a source repository hosted in Amazon
// CodeCatalyst. You cannot use this API to create a branch in a linked repository.
func codecatalyst_CreateSourceRepositoryBranch(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.CreateSourceRepositoryBranchInput{
		// Name: *string, // Required
		// ProjectName: *string, // Required
		// SourceRepositoryName: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_codecatalystName) > 0 {
		input.Name = aws.String(_codecatalystName)
	}
	if len(_codecatalystProjectName) > 0 {
		input.ProjectName = aws.String(_codecatalystProjectName)
	}
	if len(_codecatalystSourceRepositoryName) > 0 {
		input.SourceRepositoryName = aws.String(_codecatalystSourceRepositoryName)
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}
	if len(_codecatalystHeadCommitId) > 0 {
		input.HeadCommitId = aws.String(_codecatalystHeadCommitId)
	}

	if resp, err := client.CreateSourceRepositoryBranch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified personal access token (PAT). A personal access token can
// only be deleted by the user who created it.
func codecatalyst_DeleteAccessToken(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.DeleteAccessTokenInput{
		// Id: *string, // Required
	}

	if len(_codecatalystId) > 0 {
		input.Id = aws.String(_codecatalystId)
	}

	if resp, err := client.DeleteAccessToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Dev Environment.
func codecatalyst_DeleteDevEnvironment(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.DeleteDevEnvironmentInput{
		// Id: *string, // Required
		// ProjectName: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_codecatalystId) > 0 {
		input.Id = aws.String(_codecatalystId)
	}
	if len(_codecatalystProjectName) > 0 {
		input.ProjectName = aws.String(_codecatalystProjectName)
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}

	if resp, err := client.DeleteDevEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a project in a space.
func codecatalyst_DeleteProject(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.DeleteProjectInput{
		// Name: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_codecatalystName) > 0 {
		input.Name = aws.String(_codecatalystName)
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}

	if resp, err := client.DeleteProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a source repository in Amazon CodeCatalyst. You cannot use this API to
// delete a linked repository. It can only be used to delete a Amazon CodeCatalyst
// source repository.
func codecatalyst_DeleteSourceRepository(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.DeleteSourceRepositoryInput{
		// Name: *string, // Required
		// ProjectName: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_codecatalystName) > 0 {
		input.Name = aws.String(_codecatalystName)
	}
	if len(_codecatalystProjectName) > 0 {
		input.ProjectName = aws.String(_codecatalystProjectName)
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}

	if resp, err := client.DeleteSourceRepository(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a space.
// Deleting a space cannot be undone. Additionally, since space names must be
// unique across Amazon CodeCatalyst, you cannot reuse names of deleted spaces.
func codecatalyst_DeleteSpace(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.DeleteSpaceInput{
		// Name: *string, // Required
	}

	if len(_codecatalystName) > 0 {
		input.Name = aws.String(_codecatalystName)
	}

	if resp, err := client.DeleteSpace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a Dev Environment for a source repository in a
// project. Dev Environments are specific to the user who creates them.
func codecatalyst_GetDevEnvironment(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.GetDevEnvironmentInput{
		// Id: *string, // Required
		// ProjectName: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_codecatalystId) > 0 {
		input.Id = aws.String(_codecatalystId)
	}
	if len(_codecatalystProjectName) > 0 {
		input.ProjectName = aws.String(_codecatalystProjectName)
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}

	if resp, err := client.GetDevEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a project.
func codecatalyst_GetProject(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.GetProjectInput{
		// Name: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_codecatalystName) > 0 {
		input.Name = aws.String(_codecatalystName)
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}

	if resp, err := client.GetProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a source repository.
func codecatalyst_GetSourceRepository(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.GetSourceRepositoryInput{
		// Name: *string, // Required
		// ProjectName: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_codecatalystName) > 0 {
		input.Name = aws.String(_codecatalystName)
	}
	if len(_codecatalystProjectName) > 0 {
		input.ProjectName = aws.String(_codecatalystProjectName)
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}

	if resp, err := client.GetSourceRepository(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the URLs that can be used with a Git client to clone
// a source repository.
func codecatalyst_GetSourceRepositoryCloneUrls(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.GetSourceRepositoryCloneUrlsInput{
		// ProjectName: *string, // Required
		// SourceRepositoryName: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_codecatalystProjectName) > 0 {
		input.ProjectName = aws.String(_codecatalystProjectName)
	}
	if len(_codecatalystSourceRepositoryName) > 0 {
		input.SourceRepositoryName = aws.String(_codecatalystSourceRepositoryName)
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}

	if resp, err := client.GetSourceRepositoryCloneUrls(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an space.
func codecatalyst_GetSpace(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.GetSpaceInput{
		// Name: *string, // Required
	}

	if len(_codecatalystName) > 0 {
		input.Name = aws.String(_codecatalystName)
	}

	if resp, err := client.GetSpace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the Amazon Web Services account used for billing
// purposes and the billing plan for the space.
func codecatalyst_GetSubscription(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.GetSubscriptionInput{
		// SpaceName: *string, // Required
	}

	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}

	if resp, err := client.GetSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a user.
func codecatalyst_GetUserDetails(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.GetUserDetailsInput{}

	if len(_codecatalystId) > 0 {
		input.Id = aws.String(_codecatalystId)
	}
	if len(_codecatalystUserName) > 0 {
		input.UserName = aws.String(_codecatalystUserName)
	}

	if resp, err := client.GetUserDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a workflow.
func codecatalyst_GetWorkflow(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.GetWorkflowInput{
		// Id: *string, // Required
		// ProjectName: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_codecatalystId) > 0 {
		input.Id = aws.String(_codecatalystId)
	}
	if len(_codecatalystProjectName) > 0 {
		input.ProjectName = aws.String(_codecatalystProjectName)
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}

	if resp, err := client.GetWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specified run of a workflow.
func codecatalyst_GetWorkflowRun(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.GetWorkflowRunInput{
		// Id: *string, // Required
		// ProjectName: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_codecatalystId) > 0 {
		input.Id = aws.String(_codecatalystId)
	}
	if len(_codecatalystProjectName) > 0 {
		input.ProjectName = aws.String(_codecatalystProjectName)
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}

	if resp, err := client.GetWorkflowRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all personal access tokens (PATs) associated with the user who calls the
// API. You can only list PATs associated with your Amazon Web Services Builder ID.
func codecatalyst_ListAccessTokens(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.ListAccessTokensInput{}

	if len(_codecatalystMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codecatalystMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codecatalystNextToken) > 0 {
		input.NextToken = aws.String(_codecatalystNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccessTokens(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codecatalyst.ListAccessTokensOutput
	p := codecatalyst.NewListAccessTokensPaginator(client, input)
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

// Retrieves a list of active sessions for a Dev Environment in a project.
func codecatalyst_ListDevEnvironmentSessions(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.ListDevEnvironmentSessionsInput{
		// DevEnvironmentId: *string, // Required
		// ProjectName: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_codecatalystDevEnvironmentId) > 0 {
		input.DevEnvironmentId = aws.String(_codecatalystDevEnvironmentId)
	}
	if len(_codecatalystProjectName) > 0 {
		input.ProjectName = aws.String(_codecatalystProjectName)
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}
	if len(_codecatalystMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codecatalystMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codecatalystNextToken) > 0 {
		input.NextToken = aws.String(_codecatalystNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDevEnvironmentSessions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codecatalyst.ListDevEnvironmentSessionsOutput
	p := codecatalyst.NewListDevEnvironmentSessionsPaginator(client, input)
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

// Retrieves a list of Dev Environments in a project.
func codecatalyst_ListDevEnvironments(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.ListDevEnvironmentsInput{
		// SpaceName: *string, // Required
	}

	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}
	if len(_codecatalystFilters) > 0 {
		if err := assignInputField(input, "Filters", _codecatalystFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_codecatalystMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codecatalystMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codecatalystNextToken) > 0 {
		input.NextToken = aws.String(_codecatalystNextToken)
	}
	if len(_codecatalystProjectName) > 0 {
		input.ProjectName = aws.String(_codecatalystProjectName)
	}

	if disablePaginator() {
		if resp, err := client.ListDevEnvironments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codecatalyst.ListDevEnvironmentsOutput
	p := codecatalyst.NewListDevEnvironmentsPaginator(client, input)
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

// Retrieves a list of events that occurred during a specific time in a space. You
// can use these events to audit user and system activity in a space. For more
// information, see [Monitoring]in the Amazon CodeCatalyst User Guide.
//
// ListEventLogs guarantees events for the last 30 days in a given space. You can
// also view and retrieve a list of management events over the last 90 days for
// Amazon CodeCatalyst in the CloudTrail console by viewing Event history, or by
// creating a trail to create and maintain a record of events that extends past 90
// days. For more information, see [Working with CloudTrail Event History]and [Working with CloudTrail trails].
//
// [Working with CloudTrail Event History]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/view-cloudtrail-events.html
// [Working with CloudTrail trails]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-getting-started.html
// [Monitoring]: https://docs.aws.amazon.com/codecatalyst/latest/userguide/ipa-monitoring.html
func codecatalyst_ListEventLogs(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.ListEventLogsInput{
		// EndTime: *time.Time, // Required
		// SpaceName: *string, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_codecatalystEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _codecatalystEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}
	if len(_codecatalystStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _codecatalystStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_codecatalystEventName) > 0 {
		input.EventName = aws.String(_codecatalystEventName)
	}
	if len(_codecatalystMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codecatalystMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codecatalystNextToken) > 0 {
		input.NextToken = aws.String(_codecatalystNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEventLogs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codecatalyst.ListEventLogsOutput
	p := codecatalyst.NewListEventLogsPaginator(client, input)
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

// Retrieves a list of projects.
func codecatalyst_ListProjects(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.ListProjectsInput{
		// SpaceName: *string, // Required
	}

	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}
	if len(_codecatalystFilters) > 0 {
		if err := assignInputField(input, "Filters", _codecatalystFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_codecatalystMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codecatalystMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codecatalystNextToken) > 0 {
		input.NextToken = aws.String(_codecatalystNextToken)
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

	var results []*codecatalyst.ListProjectsOutput
	p := codecatalyst.NewListProjectsPaginator(client, input)
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

// Retrieves a list of source repositories in a project.
func codecatalyst_ListSourceRepositories(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.ListSourceRepositoriesInput{
		// ProjectName: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_codecatalystProjectName) > 0 {
		input.ProjectName = aws.String(_codecatalystProjectName)
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}
	if len(_codecatalystMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codecatalystMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codecatalystNextToken) > 0 {
		input.NextToken = aws.String(_codecatalystNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSourceRepositories(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codecatalyst.ListSourceRepositoriesOutput
	p := codecatalyst.NewListSourceRepositoriesPaginator(client, input)
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

// Retrieves a list of branches in a specified source repository.
func codecatalyst_ListSourceRepositoryBranches(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.ListSourceRepositoryBranchesInput{
		// ProjectName: *string, // Required
		// SourceRepositoryName: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_codecatalystProjectName) > 0 {
		input.ProjectName = aws.String(_codecatalystProjectName)
	}
	if len(_codecatalystSourceRepositoryName) > 0 {
		input.SourceRepositoryName = aws.String(_codecatalystSourceRepositoryName)
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}
	if len(_codecatalystMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codecatalystMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codecatalystNextToken) > 0 {
		input.NextToken = aws.String(_codecatalystNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSourceRepositoryBranches(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codecatalyst.ListSourceRepositoryBranchesOutput
	p := codecatalyst.NewListSourceRepositoryBranchesPaginator(client, input)
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

// Retrieves a list of spaces.
func codecatalyst_ListSpaces(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.ListSpacesInput{}

	if len(_codecatalystNextToken) > 0 {
		input.NextToken = aws.String(_codecatalystNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSpaces(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codecatalyst.ListSpacesOutput
	p := codecatalyst.NewListSpacesPaginator(client, input)
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

// Retrieves a list of workflow runs of a specified workflow.
func codecatalyst_ListWorkflowRuns(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.ListWorkflowRunsInput{
		// ProjectName: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_codecatalystProjectName) > 0 {
		input.ProjectName = aws.String(_codecatalystProjectName)
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}
	if len(_codecatalystMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codecatalystMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codecatalystNextToken) > 0 {
		input.NextToken = aws.String(_codecatalystNextToken)
	}
	if len(_codecatalystSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _codecatalystSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_codecatalystWorkflowId) > 0 {
		input.WorkflowId = aws.String(_codecatalystWorkflowId)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkflowRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codecatalyst.ListWorkflowRunsOutput
	p := codecatalyst.NewListWorkflowRunsPaginator(client, input)
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

// Retrieves a list of workflows in a specified project.
func codecatalyst_ListWorkflows(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.ListWorkflowsInput{
		// ProjectName: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_codecatalystProjectName) > 0 {
		input.ProjectName = aws.String(_codecatalystProjectName)
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}
	if len(_codecatalystMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codecatalystMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codecatalystNextToken) > 0 {
		input.NextToken = aws.String(_codecatalystNextToken)
	}
	if len(_codecatalystSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _codecatalystSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
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

	var results []*codecatalyst.ListWorkflowsOutput
	p := codecatalyst.NewListWorkflowsPaginator(client, input)
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

// Starts a specified Dev Environment and puts it into an active state.
func codecatalyst_StartDevEnvironment(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.StartDevEnvironmentInput{
		// Id: *string, // Required
		// ProjectName: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_codecatalystId) > 0 {
		input.Id = aws.String(_codecatalystId)
	}
	if len(_codecatalystProjectName) > 0 {
		input.ProjectName = aws.String(_codecatalystProjectName)
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}
	if len(_codecatalystIdes) > 0 {
		if err := assignInputField(input, "Ides", _codecatalystIdes); err != nil {
			log.Errorf("invalid --ides: %s", err.Error())
			return
		}
	}
	if len(_codecatalystInactivityTimeoutMinutes) > 0 {
		if err := assignInputField(input, "InactivityTimeoutMinutes", _codecatalystInactivityTimeoutMinutes); err != nil {
			log.Errorf("invalid --inactivity-timeout-minutes: %s", err.Error())
			return
		}
	}
	if len(_codecatalystInstanceType) > 0 {
		if err := assignInputField(input, "InstanceType", _codecatalystInstanceType); err != nil {
			log.Errorf("invalid --instance-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartDevEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a session for a specified Dev Environment.
func codecatalyst_StartDevEnvironmentSession(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.StartDevEnvironmentSessionInput{
		// Id: *string, // Required
		// ProjectName: *string, // Required
		// SessionConfiguration: *types.DevEnvironmentSessionConfiguration, // Required
		// SpaceName: *string, // Required
	}

	if len(_codecatalystId) > 0 {
		input.Id = aws.String(_codecatalystId)
	}
	if len(_codecatalystProjectName) > 0 {
		input.ProjectName = aws.String(_codecatalystProjectName)
	}
	if len(_codecatalystSessionConfiguration) > 0 {
		if err := assignInputField(input, "SessionConfiguration", _codecatalystSessionConfiguration); err != nil {
			log.Errorf("invalid --session-configuration: %s", err.Error())
			return
		}
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}

	if resp, err := client.StartDevEnvironmentSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Begins a run of a specified workflow.
func codecatalyst_StartWorkflowRun(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.StartWorkflowRunInput{
		// ProjectName: *string, // Required
		// SpaceName: *string, // Required
		// WorkflowId: *string, // Required
	}

	if len(_codecatalystProjectName) > 0 {
		input.ProjectName = aws.String(_codecatalystProjectName)
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}
	if len(_codecatalystWorkflowId) > 0 {
		input.WorkflowId = aws.String(_codecatalystWorkflowId)
	}
	if len(_codecatalystClientToken) > 0 {
		input.ClientToken = aws.String(_codecatalystClientToken)
	}

	if resp, err := client.StartWorkflowRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Pauses a specified Dev Environment and places it in a non-running state.
// Stopped Dev Environments do not consume compute minutes.
func codecatalyst_StopDevEnvironment(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.StopDevEnvironmentInput{
		// Id: *string, // Required
		// ProjectName: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_codecatalystId) > 0 {
		input.Id = aws.String(_codecatalystId)
	}
	if len(_codecatalystProjectName) > 0 {
		input.ProjectName = aws.String(_codecatalystProjectName)
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}

	if resp, err := client.StopDevEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a session for a specified Dev Environment.
func codecatalyst_StopDevEnvironmentSession(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.StopDevEnvironmentSessionInput{
		// Id: *string, // Required
		// ProjectName: *string, // Required
		// SessionId: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_codecatalystId) > 0 {
		input.Id = aws.String(_codecatalystId)
	}
	if len(_codecatalystProjectName) > 0 {
		input.ProjectName = aws.String(_codecatalystProjectName)
	}
	if len(_codecatalystSessionId) > 0 {
		input.SessionId = aws.String(_codecatalystSessionId)
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}

	if resp, err := client.StopDevEnvironmentSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes one or more values for a Dev Environment. Updating certain values of
// the Dev Environment will cause a restart.
func codecatalyst_UpdateDevEnvironment(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.UpdateDevEnvironmentInput{
		// Id: *string, // Required
		// ProjectName: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_codecatalystId) > 0 {
		input.Id = aws.String(_codecatalystId)
	}
	if len(_codecatalystProjectName) > 0 {
		input.ProjectName = aws.String(_codecatalystProjectName)
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}
	if len(_codecatalystAlias) > 0 {
		input.Alias = aws.String(_codecatalystAlias)
	}
	if len(_codecatalystClientToken) > 0 {
		input.ClientToken = aws.String(_codecatalystClientToken)
	}
	if len(_codecatalystIdes) > 0 {
		if err := assignInputField(input, "Ides", _codecatalystIdes); err != nil {
			log.Errorf("invalid --ides: %s", err.Error())
			return
		}
	}
	if len(_codecatalystInactivityTimeoutMinutes) > 0 {
		if err := assignInputField(input, "InactivityTimeoutMinutes", _codecatalystInactivityTimeoutMinutes); err != nil {
			log.Errorf("invalid --inactivity-timeout-minutes: %s", err.Error())
			return
		}
	}
	if len(_codecatalystInstanceType) > 0 {
		if err := assignInputField(input, "InstanceType", _codecatalystInstanceType); err != nil {
			log.Errorf("invalid --instance-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDevEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes one or more values for a project.
func codecatalyst_UpdateProject(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.UpdateProjectInput{
		// Name: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_codecatalystName) > 0 {
		input.Name = aws.String(_codecatalystName)
	}
	if len(_codecatalystSpaceName) > 0 {
		input.SpaceName = aws.String(_codecatalystSpaceName)
	}
	if len(_codecatalystDescription) > 0 {
		input.Description = aws.String(_codecatalystDescription)
	}

	if resp, err := client.UpdateProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes one or more values for a space.
func codecatalyst_UpdateSpace(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.UpdateSpaceInput{
		// Name: *string, // Required
	}

	if len(_codecatalystName) > 0 {
		input.Name = aws.String(_codecatalystName)
	}
	if len(_codecatalystDescription) > 0 {
		input.Description = aws.String(_codecatalystDescription)
	}

	if resp, err := client.UpdateSpace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Verifies whether the calling user has a valid Amazon CodeCatalyst login and
// session. If successful, this returns the ID of the user in Amazon CodeCatalyst.
func codecatalyst_VerifySession(cfg aws.Config, client *codecatalyst.Client) {
	input := &codecatalyst.VerifySessionInput{}

	if resp, err := client.VerifySession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_codecatalystCmd)
	_codecatalystCmd.Flags().SortFlags = false

	_codecatalystCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_codecatalystCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_codecatalystCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_codecatalystCmd.Flags().StringVarP(&_codecatalystAlias, "alias", "", "", "Alias")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystClientToken, "client-token", "", "", "Client Token")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystDescription, "description", "", "", "Description")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystDevEnvironmentId, "dev-environment-id", "", "", "Dev Environment ID")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystDisplayName, "display-name", "", "", "Display Name")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystEndTime, "end-time", "", "", "End Time")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystEventName, "event-name", "", "", "Event Name")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystExpiresTime, "expires-time", "", "", "Expires Time")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystFilters, "filters", "", "", "Filters")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystHeadCommitId, "head-commit-id", "", "", "Head Commit ID")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystId, "id", "", "", "ID")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystIdes, "ides", "", "", "Ides")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystInactivityTimeoutMinutes, "inactivity-timeout-minutes", "", "", "Inactivity Timeout Minutes")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystInstanceType, "instance-type", "", "", "Instance Type")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystMaxResults, "max-results", "", "", "Max Results")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystName, "name", "", "", "Name")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystNextToken, "next-token", "", "", "Next Token")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystPersistentStorage, "persistent-storage", "", "", "Persistent Storage")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystProjectName, "project-name", "", "", "Project Name")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystRepositories, "repositories", "", "", "Repositories")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystSessionConfiguration, "session-configuration", "", "", "Session Configuration")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystSessionId, "session-id", "", "", "Session ID")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystSortBy, "sort-by", "", "", "Sort By")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystSourceRepositoryName, "source-repository-name", "", "", "Source Repository Name")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystSpaceName, "space-name", "", "", "Space Name")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystStartTime, "start-time", "", "", "Start Time")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystUserName, "user-name", "", "", "User Name")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystVpcConnectionName, "vpc-connection-name", "", "", "VPC Connection Name")
	_codecatalystCmd.Flags().StringVarP(&_codecatalystWorkflowId, "workflow-id", "", "", "Workflow ID")

	_codecatalystCmd.Flags().BoolVarP(&_codecatalystCreateAccessToken, "create-access-token", "", false, "Create Access Token")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystCreateDevEnvironment, "create-dev-environment", "", false, "Create Dev Environment")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystCreateProject, "create-project", "", false, "Create Project")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystCreateSourceRepository, "create-source-repository", "", false, "Create Source Repository")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystCreateSourceRepositoryBranch, "create-source-repository-branch", "", false, "Create Source Repository Branch")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystDeleteAccessToken, "delete-access-token", "", false, "Delete Access Token")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystDeleteDevEnvironment, "delete-dev-environment", "", false, "Delete Dev Environment")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystDeleteProject, "delete-project", "", false, "Delete Project")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystDeleteSourceRepository, "delete-source-repository", "", false, "Delete Source Repository")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystDeleteSpace, "delete-space", "", false, "Delete Space")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystGetDevEnvironment, "get-dev-environment", "", false, "Get Dev Environment")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystGetProject, "get-project", "", false, "Get Project")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystGetSourceRepository, "get-source-repository", "", false, "Get Source Repository")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystGetSourceRepositoryCloneUrls, "get-source-repository-clone-urls", "", false, "Get Source Repository Clone Urls")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystGetSpace, "get-space", "", false, "Get Space")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystGetSubscription, "get-subscription", "", false, "Get Subscription")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystGetUserDetails, "get-user-details", "", false, "Get User Details")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystGetWorkflow, "get-workflow", "", false, "Get Workflow")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystGetWorkflowRun, "get-workflow-run", "", false, "Get Workflow Run")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystListAccessTokens, "list-access-tokens", "", false, "List Access Tokens")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystListDevEnvironmentSessions, "list-dev-environment-sessions", "", false, "List Dev Environment Sessions")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystListDevEnvironments, "list-dev-environments", "", false, "List Dev Environments")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystListEventLogs, "list-event-logs", "", false, "List Event Logs")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystListProjects, "list-projects", "", false, "List Projects")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystListSourceRepositories, "list-source-repositories", "", false, "List Source Repositories")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystListSourceRepositoryBranches, "list-source-repository-branches", "", false, "List Source Repository Branches")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystListSpaces, "list-spaces", "", false, "List Spaces")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystListWorkflowRuns, "list-workflow-runs", "", false, "List Workflow Runs")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystListWorkflows, "list-workflows", "", false, "List Workflows")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystStartDevEnvironment, "start-dev-environment", "", false, "Start Dev Environment")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystStartDevEnvironmentSession, "start-dev-environment-session", "", false, "Start Dev Environment Session")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystStartWorkflowRun, "start-workflow-run", "", false, "Start Workflow Run")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystStopDevEnvironment, "stop-dev-environment", "", false, "Stop Dev Environment")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystStopDevEnvironmentSession, "stop-dev-environment-session", "", false, "Stop Dev Environment Session")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystUpdateDevEnvironment, "update-dev-environment", "", false, "Update Dev Environment")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystUpdateProject, "update-project", "", false, "Update Project")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystUpdateSpace, "update-space", "", false, "Update Space")
	_codecatalystCmd.Flags().BoolVarP(&_codecatalystVerifySession, "verify-session", "", false, "Verify Session")

}
