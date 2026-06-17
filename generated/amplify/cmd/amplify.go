package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// amplifyCmd represents the amplify command
var _amplifyCmd = &cobra.Command{
	Use:   "amplify",
	Short: "AWS amplify CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := amplify.NewFromConfig(cfg)
		if _amplifyCreateApp {
			amplify_CreateApp(cfg, client)
			return
		}
		if _amplifyCreateBackendEnvironment {
			amplify_CreateBackendEnvironment(cfg, client)
			return
		}
		if _amplifyCreateBranch {
			amplify_CreateBranch(cfg, client)
			return
		}
		if _amplifyCreateDeployment {
			amplify_CreateDeployment(cfg, client)
			return
		}
		if _amplifyCreateDomainAssociation {
			amplify_CreateDomainAssociation(cfg, client)
			return
		}
		if _amplifyCreateWebhook {
			amplify_CreateWebhook(cfg, client)
			return
		}
		if _amplifyDeleteApp {
			amplify_DeleteApp(cfg, client)
			return
		}
		if _amplifyDeleteBackendEnvironment {
			amplify_DeleteBackendEnvironment(cfg, client)
			return
		}
		if _amplifyDeleteBranch {
			amplify_DeleteBranch(cfg, client)
			return
		}
		if _amplifyDeleteDomainAssociation {
			amplify_DeleteDomainAssociation(cfg, client)
			return
		}
		if _amplifyDeleteJob {
			amplify_DeleteJob(cfg, client)
			return
		}
		if _amplifyDeleteWebhook {
			amplify_DeleteWebhook(cfg, client)
			return
		}
		if _amplifyGenerateAccessLogs {
			amplify_GenerateAccessLogs(cfg, client)
			return
		}
		if _amplifyGetApp {
			amplify_GetApp(cfg, client)
			return
		}
		if _amplifyGetArtifactUrl {
			amplify_GetArtifactUrl(cfg, client)
			return
		}
		if _amplifyGetBackendEnvironment {
			amplify_GetBackendEnvironment(cfg, client)
			return
		}
		if _amplifyGetBranch {
			amplify_GetBranch(cfg, client)
			return
		}
		if _amplifyGetDomainAssociation {
			amplify_GetDomainAssociation(cfg, client)
			return
		}
		if _amplifyGetJob {
			amplify_GetJob(cfg, client)
			return
		}
		if _amplifyGetWebhook {
			amplify_GetWebhook(cfg, client)
			return
		}
		if _amplifyListApps {
			amplify_ListApps(cfg, client)
			return
		}
		if _amplifyListArtifacts {
			amplify_ListArtifacts(cfg, client)
			return
		}
		if _amplifyListBackendEnvironments {
			amplify_ListBackendEnvironments(cfg, client)
			return
		}
		if _amplifyListBranches {
			amplify_ListBranches(cfg, client)
			return
		}
		if _amplifyListDomainAssociations {
			amplify_ListDomainAssociations(cfg, client)
			return
		}
		if _amplifyListJobs {
			amplify_ListJobs(cfg, client)
			return
		}
		if _amplifyListTagsForResource {
			amplify_ListTagsForResource(cfg, client)
			return
		}
		if _amplifyListWebhooks {
			amplify_ListWebhooks(cfg, client)
			return
		}
		if _amplifyStartDeployment {
			amplify_StartDeployment(cfg, client)
			return
		}
		if _amplifyStartJob {
			amplify_StartJob(cfg, client)
			return
		}
		if _amplifyStopJob {
			amplify_StopJob(cfg, client)
			return
		}
		if _amplifyTagResource {
			amplify_TagResource(cfg, client)
			return
		}
		if _amplifyUntagResource {
			amplify_UntagResource(cfg, client)
			return
		}
		if _amplifyUpdateApp {
			amplify_UpdateApp(cfg, client)
			return
		}
		if _amplifyUpdateBranch {
			amplify_UpdateBranch(cfg, client)
			return
		}
		if _amplifyUpdateDomainAssociation {
			amplify_UpdateDomainAssociation(cfg, client)
			return
		}
		if _amplifyUpdateWebhook {
			amplify_UpdateWebhook(cfg, client)
			return
		}

	},
}

var (
	_amplifyCreateApp                bool
	_amplifyCreateBackendEnvironment bool
	_amplifyCreateBranch             bool
	_amplifyCreateDeployment         bool
	_amplifyCreateDomainAssociation  bool
	_amplifyCreateWebhook            bool
	_amplifyDeleteApp                bool
	_amplifyDeleteBackendEnvironment bool
	_amplifyDeleteBranch             bool
	_amplifyDeleteDomainAssociation  bool
	_amplifyDeleteJob                bool
	_amplifyDeleteWebhook            bool
	_amplifyGenerateAccessLogs       bool
	_amplifyGetApp                   bool
	_amplifyGetArtifactUrl           bool
	_amplifyGetBackendEnvironment    bool
	_amplifyGetBranch                bool
	_amplifyGetDomainAssociation     bool
	_amplifyGetJob                   bool
	_amplifyGetWebhook               bool
	_amplifyListApps                 bool
	_amplifyListArtifacts            bool
	_amplifyListBackendEnvironments  bool
	_amplifyListBranches             bool
	_amplifyListDomainAssociations   bool
	_amplifyListJobs                 bool
	_amplifyListTagsForResource      bool
	_amplifyListWebhooks             bool
	_amplifyStartDeployment          bool
	_amplifyStartJob                 bool
	_amplifyStopJob                  bool
	_amplifyTagResource              bool
	_amplifyUntagResource            bool
	_amplifyUpdateApp                bool
	_amplifyUpdateBranch             bool
	_amplifyUpdateDomainAssociation  bool
	_amplifyUpdateWebhook            bool

	_amplifyAccessToken                   string
	_amplifyAppId                         string
	_amplifyArtifactId                    string
	_amplifyAutoBranchCreationConfig      string
	_amplifyAutoBranchCreationPatterns    []string
	_amplifyAutoSubDomainCreationPatterns []string
	_amplifyAutoSubDomainIAMRole          string
	_amplifyBackend                       string
	_amplifyBackendEnvironmentArn         string
	_amplifyBasicAuthCredentials          string
	_amplifyBranchName                    string
	_amplifyBuildSpec                     string
	_amplifyCacheConfig                   string
	_amplifyCertificateSettings           string
	_amplifyCommitId                      string
	_amplifyCommitMessage                 string
	_amplifyCommitTime                    string
	_amplifyComputeRoleArn                string
	_amplifyCustomHeaders                 string
	_amplifyCustomRules                   string
	_amplifyDeploymentArtifacts           string
	_amplifyDescription                   string
	_amplifyDisplayName                   string
	_amplifyDomainName                    string
	_amplifyEnableAutoBranchCreation      string
	_amplifyEnableAutoBuild               string
	_amplifyEnableAutoSubDomain           string
	_amplifyEnableBasicAuth               string
	_amplifyEnableBranchAutoBuild         string
	_amplifyEnableBranchAutoDeletion      string
	_amplifyEnableNotification            string
	_amplifyEnablePerformanceMode         string
	_amplifyEnablePullRequestPreview      string
	_amplifyEnableSkewProtection          string
	_amplifyEndTime                       string
	_amplifyEnvironmentName               string
	_amplifyEnvironmentVariables          string
	_amplifyFileMap                       string
	_amplifyFramework                     string
	_amplifyIamServiceRoleArn             string
	_amplifyJobConfig                     string
	_amplifyJobId                         string
	_amplifyJobReason                     string
	_amplifyJobType                       string
	_amplifyMaxResults                    string
	_amplifyName                          string
	_amplifyNextToken                     string
	_amplifyOauthToken                    string
	_amplifyPlatform                      string
	_amplifyPullRequestEnvironmentName    string
	_amplifyRepository                    string
	_amplifyResourceArn                   string
	_amplifySourceUrl                     string
	_amplifySourceUrlType                 string
	_amplifyStackName                     string
	_amplifyStage                         string
	_amplifyStartTime                     string
	_amplifySubDomainSettings             string
	_amplifyTagKeys                       []string
	_amplifyTags                          string
	_amplifyTtl                           string
	_amplifyWebhookId                     string
)

// Creates a new Amplify app.
func amplify_CreateApp(cfg aws.Config, client *amplify.Client) {
	input := &amplify.CreateAppInput{
		// Name: *string, // Required
	}

	if len(_amplifyName) > 0 {
		input.Name = aws.String(_amplifyName)
	}
	if len(_amplifyAccessToken) > 0 {
		input.AccessToken = aws.String(_amplifyAccessToken)
	}
	if len(_amplifyAutoBranchCreationConfig) > 0 {
		if err := assignInputField(input, "AutoBranchCreationConfig", _amplifyAutoBranchCreationConfig); err != nil {
			log.Errorf("invalid --auto-branch-creation-config: %s", err.Error())
			return
		}
	}
	if len(_amplifyAutoBranchCreationPatterns) > 0 {
		input.AutoBranchCreationPatterns = append([]string(nil), _amplifyAutoBranchCreationPatterns...)
	}
	if len(_amplifyBasicAuthCredentials) > 0 {
		input.BasicAuthCredentials = aws.String(_amplifyBasicAuthCredentials)
	}
	if len(_amplifyBuildSpec) > 0 {
		input.BuildSpec = aws.String(_amplifyBuildSpec)
	}
	if len(_amplifyCacheConfig) > 0 {
		if err := assignInputField(input, "CacheConfig", _amplifyCacheConfig); err != nil {
			log.Errorf("invalid --cache-config: %s", err.Error())
			return
		}
	}
	if len(_amplifyComputeRoleArn) > 0 {
		input.ComputeRoleArn = aws.String(_amplifyComputeRoleArn)
	}
	if len(_amplifyCustomHeaders) > 0 {
		input.CustomHeaders = aws.String(_amplifyCustomHeaders)
	}
	if len(_amplifyCustomRules) > 0 {
		if err := assignInputField(input, "CustomRules", _amplifyCustomRules); err != nil {
			log.Errorf("invalid --custom-rules: %s", err.Error())
			return
		}
	}
	if len(_amplifyDescription) > 0 {
		input.Description = aws.String(_amplifyDescription)
	}
	if len(_amplifyEnableAutoBranchCreation) > 0 {
		if err := assignInputField(input, "EnableAutoBranchCreation", _amplifyEnableAutoBranchCreation); err != nil {
			log.Errorf("invalid --enable-auto-branch-creation: %s", err.Error())
			return
		}
	}
	if len(_amplifyEnableBasicAuth) > 0 {
		if err := assignInputField(input, "EnableBasicAuth", _amplifyEnableBasicAuth); err != nil {
			log.Errorf("invalid --enable-basic-auth: %s", err.Error())
			return
		}
	}
	if len(_amplifyEnableBranchAutoBuild) > 0 {
		if err := assignInputField(input, "EnableBranchAutoBuild", _amplifyEnableBranchAutoBuild); err != nil {
			log.Errorf("invalid --enable-branch-auto-build: %s", err.Error())
			return
		}
	}
	if len(_amplifyEnableBranchAutoDeletion) > 0 {
		if err := assignInputField(input, "EnableBranchAutoDeletion", _amplifyEnableBranchAutoDeletion); err != nil {
			log.Errorf("invalid --enable-branch-auto-deletion: %s", err.Error())
			return
		}
	}
	if len(_amplifyEnvironmentVariables) > 0 {
		if err := assignInputField(input, "EnvironmentVariables", _amplifyEnvironmentVariables); err != nil {
			log.Errorf("invalid --environment-variables: %s", err.Error())
			return
		}
	}
	if len(_amplifyIamServiceRoleArn) > 0 {
		input.IamServiceRoleArn = aws.String(_amplifyIamServiceRoleArn)
	}
	if len(_amplifyJobConfig) > 0 {
		if err := assignInputField(input, "JobConfig", _amplifyJobConfig); err != nil {
			log.Errorf("invalid --job-config: %s", err.Error())
			return
		}
	}
	if len(_amplifyOauthToken) > 0 {
		input.OauthToken = aws.String(_amplifyOauthToken)
	}
	if len(_amplifyPlatform) > 0 {
		if err := assignInputField(input, "Platform", _amplifyPlatform); err != nil {
			log.Errorf("invalid --platform: %s", err.Error())
			return
		}
	}
	if len(_amplifyRepository) > 0 {
		input.Repository = aws.String(_amplifyRepository)
	}
	if len(_amplifyTags) > 0 {
		if err := assignInputField(input, "Tags", _amplifyTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new backend environment for an Amplify app.
// This API is available only to Amplify Gen 1 applications where the backend is
// created using Amplify Studio or the Amplify command line interface (CLI). This
// API isn’t available to Amplify Gen 2 applications. When you deploy an
// application with Amplify Gen 2, you provision the app's backend infrastructure
// using Typescript code.
func amplify_CreateBackendEnvironment(cfg aws.Config, client *amplify.Client) {
	input := &amplify.CreateBackendEnvironmentInput{
		// AppId: *string, // Required
		// EnvironmentName: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyEnvironmentName)
	}
	if len(_amplifyDeploymentArtifacts) > 0 {
		input.DeploymentArtifacts = aws.String(_amplifyDeploymentArtifacts)
	}
	if len(_amplifyStackName) > 0 {
		input.StackName = aws.String(_amplifyStackName)
	}

	if resp, err := client.CreateBackendEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new branch for an Amplify app.
func amplify_CreateBranch(cfg aws.Config, client *amplify.Client) {
	input := &amplify.CreateBranchInput{
		// AppId: *string, // Required
		// BranchName: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyBranchName) > 0 {
		input.BranchName = aws.String(_amplifyBranchName)
	}
	if len(_amplifyBackend) > 0 {
		if err := assignInputField(input, "Backend", _amplifyBackend); err != nil {
			log.Errorf("invalid --backend: %s", err.Error())
			return
		}
	}
	if len(_amplifyBackendEnvironmentArn) > 0 {
		input.BackendEnvironmentArn = aws.String(_amplifyBackendEnvironmentArn)
	}
	if len(_amplifyBasicAuthCredentials) > 0 {
		input.BasicAuthCredentials = aws.String(_amplifyBasicAuthCredentials)
	}
	if len(_amplifyBuildSpec) > 0 {
		input.BuildSpec = aws.String(_amplifyBuildSpec)
	}
	if len(_amplifyComputeRoleArn) > 0 {
		input.ComputeRoleArn = aws.String(_amplifyComputeRoleArn)
	}
	if len(_amplifyDescription) > 0 {
		input.Description = aws.String(_amplifyDescription)
	}
	if len(_amplifyDisplayName) > 0 {
		input.DisplayName = aws.String(_amplifyDisplayName)
	}
	if len(_amplifyEnableAutoBuild) > 0 {
		if err := assignInputField(input, "EnableAutoBuild", _amplifyEnableAutoBuild); err != nil {
			log.Errorf("invalid --enable-auto-build: %s", err.Error())
			return
		}
	}
	if len(_amplifyEnableBasicAuth) > 0 {
		if err := assignInputField(input, "EnableBasicAuth", _amplifyEnableBasicAuth); err != nil {
			log.Errorf("invalid --enable-basic-auth: %s", err.Error())
			return
		}
	}
	if len(_amplifyEnableNotification) > 0 {
		if err := assignInputField(input, "EnableNotification", _amplifyEnableNotification); err != nil {
			log.Errorf("invalid --enable-notification: %s", err.Error())
			return
		}
	}
	if len(_amplifyEnablePerformanceMode) > 0 {
		if err := assignInputField(input, "EnablePerformanceMode", _amplifyEnablePerformanceMode); err != nil {
			log.Errorf("invalid --enable-performance-mode: %s", err.Error())
			return
		}
	}
	if len(_amplifyEnablePullRequestPreview) > 0 {
		if err := assignInputField(input, "EnablePullRequestPreview", _amplifyEnablePullRequestPreview); err != nil {
			log.Errorf("invalid --enable-pull-request-preview: %s", err.Error())
			return
		}
	}
	if len(_amplifyEnableSkewProtection) > 0 {
		if err := assignInputField(input, "EnableSkewProtection", _amplifyEnableSkewProtection); err != nil {
			log.Errorf("invalid --enable-skew-protection: %s", err.Error())
			return
		}
	}
	if len(_amplifyEnvironmentVariables) > 0 {
		if err := assignInputField(input, "EnvironmentVariables", _amplifyEnvironmentVariables); err != nil {
			log.Errorf("invalid --environment-variables: %s", err.Error())
			return
		}
	}
	if len(_amplifyFramework) > 0 {
		input.Framework = aws.String(_amplifyFramework)
	}
	if len(_amplifyPullRequestEnvironmentName) > 0 {
		input.PullRequestEnvironmentName = aws.String(_amplifyPullRequestEnvironmentName)
	}
	if len(_amplifyStage) > 0 {
		if err := assignInputField(input, "Stage", _amplifyStage); err != nil {
			log.Errorf("invalid --stage: %s", err.Error())
			return
		}
	}
	if len(_amplifyTags) > 0 {
		if err := assignInputField(input, "Tags", _amplifyTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_amplifyTtl) > 0 {
		input.Ttl = aws.String(_amplifyTtl)
	}

	if resp, err := client.CreateBranch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a deployment for a manually deployed Amplify app. Manually deployed
// apps are not connected to a Git repository.
//
// The maximum duration between the CreateDeployment call and the StartDeployment
// call cannot exceed 8 hours. If the duration exceeds 8 hours, the StartDeployment
// call and the associated Job will fail.
func amplify_CreateDeployment(cfg aws.Config, client *amplify.Client) {
	input := &amplify.CreateDeploymentInput{
		// AppId: *string, // Required
		// BranchName: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyBranchName) > 0 {
		input.BranchName = aws.String(_amplifyBranchName)
	}
	if len(_amplifyFileMap) > 0 {
		if err := assignInputField(input, "FileMap", _amplifyFileMap); err != nil {
			log.Errorf("invalid --file-map: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new domain association for an Amplify app. This action associates a
// custom domain with the Amplify app
func amplify_CreateDomainAssociation(cfg aws.Config, client *amplify.Client) {
	input := &amplify.CreateDomainAssociationInput{
		// AppId: *string, // Required
		// DomainName: *string, // Required
		// SubDomainSettings: []types.SubDomainSetting, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyDomainName) > 0 {
		input.DomainName = aws.String(_amplifyDomainName)
	}
	if len(_amplifySubDomainSettings) > 0 {
		if err := assignInputField(input, "SubDomainSettings", _amplifySubDomainSettings); err != nil {
			log.Errorf("invalid --sub-domain-settings: %s", err.Error())
			return
		}
	}
	if len(_amplifyAutoSubDomainCreationPatterns) > 0 {
		input.AutoSubDomainCreationPatterns = append([]string(nil), _amplifyAutoSubDomainCreationPatterns...)
	}
	if len(_amplifyAutoSubDomainIAMRole) > 0 {
		input.AutoSubDomainIAMRole = aws.String(_amplifyAutoSubDomainIAMRole)
	}
	if len(_amplifyCertificateSettings) > 0 {
		if err := assignInputField(input, "CertificateSettings", _amplifyCertificateSettings); err != nil {
			log.Errorf("invalid --certificate-settings: %s", err.Error())
			return
		}
	}
	if len(_amplifyEnableAutoSubDomain) > 0 {
		if err := assignInputField(input, "EnableAutoSubDomain", _amplifyEnableAutoSubDomain); err != nil {
			log.Errorf("invalid --enable-auto-sub-domain: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDomainAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new webhook on an Amplify app.
func amplify_CreateWebhook(cfg aws.Config, client *amplify.Client) {
	input := &amplify.CreateWebhookInput{
		// AppId: *string, // Required
		// BranchName: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyBranchName) > 0 {
		input.BranchName = aws.String(_amplifyBranchName)
	}
	if len(_amplifyDescription) > 0 {
		input.Description = aws.String(_amplifyDescription)
	}

	if resp, err := client.CreateWebhook(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing Amplify app specified by an app ID.
func amplify_DeleteApp(cfg aws.Config, client *amplify.Client) {
	input := &amplify.DeleteAppInput{
		// AppId: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}

	if resp, err := client.DeleteApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a backend environment for an Amplify app.
// This API is available only to Amplify Gen 1 applications where the backend is
// created using Amplify Studio or the Amplify command line interface (CLI). This
// API isn’t available to Amplify Gen 2 applications. When you deploy an
// application with Amplify Gen 2, you provision the app's backend infrastructure
// using Typescript code.
func amplify_DeleteBackendEnvironment(cfg aws.Config, client *amplify.Client) {
	input := &amplify.DeleteBackendEnvironmentInput{
		// AppId: *string, // Required
		// EnvironmentName: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyEnvironmentName)
	}

	if resp, err := client.DeleteBackendEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a branch for an Amplify app.
func amplify_DeleteBranch(cfg aws.Config, client *amplify.Client) {
	input := &amplify.DeleteBranchInput{
		// AppId: *string, // Required
		// BranchName: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyBranchName) > 0 {
		input.BranchName = aws.String(_amplifyBranchName)
	}

	if resp, err := client.DeleteBranch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a domain association for an Amplify app.
func amplify_DeleteDomainAssociation(cfg aws.Config, client *amplify.Client) {
	input := &amplify.DeleteDomainAssociationInput{
		// AppId: *string, // Required
		// DomainName: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyDomainName) > 0 {
		input.DomainName = aws.String(_amplifyDomainName)
	}

	if resp, err := client.DeleteDomainAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a job for a branch of an Amplify app.
func amplify_DeleteJob(cfg aws.Config, client *amplify.Client) {
	input := &amplify.DeleteJobInput{
		// AppId: *string, // Required
		// BranchName: *string, // Required
		// JobId: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyBranchName) > 0 {
		input.BranchName = aws.String(_amplifyBranchName)
	}
	if len(_amplifyJobId) > 0 {
		input.JobId = aws.String(_amplifyJobId)
	}

	if resp, err := client.DeleteJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a webhook.
func amplify_DeleteWebhook(cfg aws.Config, client *amplify.Client) {
	input := &amplify.DeleteWebhookInput{
		// WebhookId: *string, // Required
	}

	if len(_amplifyWebhookId) > 0 {
		input.WebhookId = aws.String(_amplifyWebhookId)
	}

	if resp, err := client.DeleteWebhook(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the website access logs for a specific time range using a presigned
// URL.
func amplify_GenerateAccessLogs(cfg aws.Config, client *amplify.Client) {
	input := &amplify.GenerateAccessLogsInput{
		// AppId: *string, // Required
		// DomainName: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyDomainName) > 0 {
		input.DomainName = aws.String(_amplifyDomainName)
	}
	if len(_amplifyEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _amplifyEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_amplifyStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _amplifyStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.GenerateAccessLogs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an existing Amplify app specified by an app ID.
func amplify_GetApp(cfg aws.Config, client *amplify.Client) {
	input := &amplify.GetAppInput{
		// AppId: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}

	if resp, err := client.GetApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the artifact info that corresponds to an artifact id.
func amplify_GetArtifactUrl(cfg aws.Config, client *amplify.Client) {
	input := &amplify.GetArtifactUrlInput{
		// ArtifactId: *string, // Required
	}

	if len(_amplifyArtifactId) > 0 {
		input.ArtifactId = aws.String(_amplifyArtifactId)
	}

	if resp, err := client.GetArtifactUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a backend environment for an Amplify app.
// This API is available only to Amplify Gen 1 applications where the backend is
// created using Amplify Studio or the Amplify command line interface (CLI). This
// API isn’t available to Amplify Gen 2 applications. When you deploy an
// application with Amplify Gen 2, you provision the app's backend infrastructure
// using Typescript code.
func amplify_GetBackendEnvironment(cfg aws.Config, client *amplify.Client) {
	input := &amplify.GetBackendEnvironmentInput{
		// AppId: *string, // Required
		// EnvironmentName: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyEnvironmentName)
	}

	if resp, err := client.GetBackendEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a branch for an Amplify app.
func amplify_GetBranch(cfg aws.Config, client *amplify.Client) {
	input := &amplify.GetBranchInput{
		// AppId: *string, // Required
		// BranchName: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyBranchName) > 0 {
		input.BranchName = aws.String(_amplifyBranchName)
	}

	if resp, err := client.GetBranch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the domain information for an Amplify app.
func amplify_GetDomainAssociation(cfg aws.Config, client *amplify.Client) {
	input := &amplify.GetDomainAssociationInput{
		// AppId: *string, // Required
		// DomainName: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyDomainName) > 0 {
		input.DomainName = aws.String(_amplifyDomainName)
	}

	if resp, err := client.GetDomainAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a job for a branch of an Amplify app.
func amplify_GetJob(cfg aws.Config, client *amplify.Client) {
	input := &amplify.GetJobInput{
		// AppId: *string, // Required
		// BranchName: *string, // Required
		// JobId: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyBranchName) > 0 {
		input.BranchName = aws.String(_amplifyBranchName)
	}
	if len(_amplifyJobId) > 0 {
		input.JobId = aws.String(_amplifyJobId)
	}

	if resp, err := client.GetJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the webhook information that corresponds to a specified webhook ID.
func amplify_GetWebhook(cfg aws.Config, client *amplify.Client) {
	input := &amplify.GetWebhookInput{
		// WebhookId: *string, // Required
	}

	if len(_amplifyWebhookId) > 0 {
		input.WebhookId = aws.String(_amplifyWebhookId)
	}

	if resp, err := client.GetWebhook(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of the existing Amplify apps.
func amplify_ListApps(cfg aws.Config, client *amplify.Client) {
	input := &amplify.ListAppsInput{}

	if len(_amplifyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _amplifyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_amplifyNextToken) > 0 {
		input.NextToken = aws.String(_amplifyNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApps(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*amplify.ListAppsOutput
	p := amplify.NewListAppsPaginator(client, input)
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

// Returns a list of end-to-end testing artifacts for a specified app, branch, and
// job.
//
// To return the build artifacts, use the [GetJob] API.
//
// For more information about Amplify testing support, see [Setting up end-to-end Cypress tests for your Amplify application] in the Amplify Hosting
// User Guide.
//
// [GetJob]: https://docs.aws.amazon.com/amplify/latest/APIReference/API_GetJob.html
// [Setting up end-to-end Cypress tests for your Amplify application]: https://docs.aws.amazon.com/amplify/latest/userguide/running-tests.html
func amplify_ListArtifacts(cfg aws.Config, client *amplify.Client) {
	input := &amplify.ListArtifactsInput{
		// AppId: *string, // Required
		// BranchName: *string, // Required
		// JobId: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyBranchName) > 0 {
		input.BranchName = aws.String(_amplifyBranchName)
	}
	if len(_amplifyJobId) > 0 {
		input.JobId = aws.String(_amplifyJobId)
	}
	if len(_amplifyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _amplifyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_amplifyNextToken) > 0 {
		input.NextToken = aws.String(_amplifyNextToken)
	}

	if resp, err := client.ListArtifacts(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the backend environments for an Amplify app.
// This API is available only to Amplify Gen 1 applications where the backend is
// created using Amplify Studio or the Amplify command line interface (CLI). This
// API isn’t available to Amplify Gen 2 applications. When you deploy an
// application with Amplify Gen 2, you provision the app's backend infrastructure
// using Typescript code.
func amplify_ListBackendEnvironments(cfg aws.Config, client *amplify.Client) {
	input := &amplify.ListBackendEnvironmentsInput{
		// AppId: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyEnvironmentName)
	}
	if len(_amplifyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _amplifyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_amplifyNextToken) > 0 {
		input.NextToken = aws.String(_amplifyNextToken)
	}

	if resp, err := client.ListBackendEnvironments(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the branches of an Amplify app.
func amplify_ListBranches(cfg aws.Config, client *amplify.Client) {
	input := &amplify.ListBranchesInput{
		// AppId: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _amplifyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_amplifyNextToken) > 0 {
		input.NextToken = aws.String(_amplifyNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBranches(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*amplify.ListBranchesOutput
	p := amplify.NewListBranchesPaginator(client, input)
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

// Returns the domain associations for an Amplify app.
func amplify_ListDomainAssociations(cfg aws.Config, client *amplify.Client) {
	input := &amplify.ListDomainAssociationsInput{
		// AppId: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _amplifyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_amplifyNextToken) > 0 {
		input.NextToken = aws.String(_amplifyNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDomainAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*amplify.ListDomainAssociationsOutput
	p := amplify.NewListDomainAssociationsPaginator(client, input)
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

// Lists the jobs for a branch of an Amplify app.
func amplify_ListJobs(cfg aws.Config, client *amplify.Client) {
	input := &amplify.ListJobsInput{
		// AppId: *string, // Required
		// BranchName: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyBranchName) > 0 {
		input.BranchName = aws.String(_amplifyBranchName)
	}
	if len(_amplifyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _amplifyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_amplifyNextToken) > 0 {
		input.NextToken = aws.String(_amplifyNextToken)
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

	var results []*amplify.ListJobsOutput
	p := amplify.NewListJobsPaginator(client, input)
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

// Returns a list of tags for a specified Amazon Resource Name (ARN).
func amplify_ListTagsForResource(cfg aws.Config, client *amplify.Client) {
	input := &amplify.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_amplifyResourceArn) > 0 {
		input.ResourceArn = aws.String(_amplifyResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of webhooks for an Amplify app.
func amplify_ListWebhooks(cfg aws.Config, client *amplify.Client) {
	input := &amplify.ListWebhooksInput{
		// AppId: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _amplifyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_amplifyNextToken) > 0 {
		input.NextToken = aws.String(_amplifyNextToken)
	}

	if resp, err := client.ListWebhooks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a deployment for a manually deployed app. Manually deployed apps are not
// connected to a Git repository.
//
// The maximum duration between the CreateDeployment call and the StartDeployment
// call cannot exceed 8 hours. If the duration exceeds 8 hours, the StartDeployment
// call and the associated Job will fail.
func amplify_StartDeployment(cfg aws.Config, client *amplify.Client) {
	input := &amplify.StartDeploymentInput{
		// AppId: *string, // Required
		// BranchName: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyBranchName) > 0 {
		input.BranchName = aws.String(_amplifyBranchName)
	}
	if len(_amplifyJobId) > 0 {
		input.JobId = aws.String(_amplifyJobId)
	}
	if len(_amplifySourceUrl) > 0 {
		input.SourceUrl = aws.String(_amplifySourceUrl)
	}
	if len(_amplifySourceUrlType) > 0 {
		if err := assignInputField(input, "SourceUrlType", _amplifySourceUrlType); err != nil {
			log.Errorf("invalid --source-url-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a new job for a branch of an Amplify app.
func amplify_StartJob(cfg aws.Config, client *amplify.Client) {
	input := &amplify.StartJobInput{
		// AppId: *string, // Required
		// BranchName: *string, // Required
		// JobType: types.JobType, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyBranchName) > 0 {
		input.BranchName = aws.String(_amplifyBranchName)
	}
	if len(_amplifyJobType) > 0 {
		if err := assignInputField(input, "JobType", _amplifyJobType); err != nil {
			log.Errorf("invalid --job-type: %s", err.Error())
			return
		}
	}
	if len(_amplifyCommitId) > 0 {
		input.CommitId = aws.String(_amplifyCommitId)
	}
	if len(_amplifyCommitMessage) > 0 {
		input.CommitMessage = aws.String(_amplifyCommitMessage)
	}
	if len(_amplifyCommitTime) > 0 {
		if err := assignInputField(input, "CommitTime", _amplifyCommitTime); err != nil {
			log.Errorf("invalid --commit-time: %s", err.Error())
			return
		}
	}
	if len(_amplifyJobId) > 0 {
		input.JobId = aws.String(_amplifyJobId)
	}
	if len(_amplifyJobReason) > 0 {
		input.JobReason = aws.String(_amplifyJobReason)
	}

	if resp, err := client.StartJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a job that is in progress for a branch of an Amplify app.
func amplify_StopJob(cfg aws.Config, client *amplify.Client) {
	input := &amplify.StopJobInput{
		// AppId: *string, // Required
		// BranchName: *string, // Required
		// JobId: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyBranchName) > 0 {
		input.BranchName = aws.String(_amplifyBranchName)
	}
	if len(_amplifyJobId) > 0 {
		input.JobId = aws.String(_amplifyJobId)
	}

	if resp, err := client.StopJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tags the resource with a tag key and value.
func amplify_TagResource(cfg aws.Config, client *amplify.Client) {
	input := &amplify.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_amplifyResourceArn) > 0 {
		input.ResourceArn = aws.String(_amplifyResourceArn)
	}
	if len(_amplifyTags) > 0 {
		if err := assignInputField(input, "Tags", _amplifyTags); err != nil {
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

// Untags a resource with a specified Amazon Resource Name (ARN).
func amplify_UntagResource(cfg aws.Config, client *amplify.Client) {
	input := &amplify.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_amplifyResourceArn) > 0 {
		input.ResourceArn = aws.String(_amplifyResourceArn)
	}
	if len(_amplifyTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _amplifyTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing Amplify app.
func amplify_UpdateApp(cfg aws.Config, client *amplify.Client) {
	input := &amplify.UpdateAppInput{
		// AppId: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyAccessToken) > 0 {
		input.AccessToken = aws.String(_amplifyAccessToken)
	}
	if len(_amplifyAutoBranchCreationConfig) > 0 {
		if err := assignInputField(input, "AutoBranchCreationConfig", _amplifyAutoBranchCreationConfig); err != nil {
			log.Errorf("invalid --auto-branch-creation-config: %s", err.Error())
			return
		}
	}
	if len(_amplifyAutoBranchCreationPatterns) > 0 {
		input.AutoBranchCreationPatterns = append([]string(nil), _amplifyAutoBranchCreationPatterns...)
	}
	if len(_amplifyBasicAuthCredentials) > 0 {
		input.BasicAuthCredentials = aws.String(_amplifyBasicAuthCredentials)
	}
	if len(_amplifyBuildSpec) > 0 {
		input.BuildSpec = aws.String(_amplifyBuildSpec)
	}
	if len(_amplifyCacheConfig) > 0 {
		if err := assignInputField(input, "CacheConfig", _amplifyCacheConfig); err != nil {
			log.Errorf("invalid --cache-config: %s", err.Error())
			return
		}
	}
	if len(_amplifyComputeRoleArn) > 0 {
		input.ComputeRoleArn = aws.String(_amplifyComputeRoleArn)
	}
	if len(_amplifyCustomHeaders) > 0 {
		input.CustomHeaders = aws.String(_amplifyCustomHeaders)
	}
	if len(_amplifyCustomRules) > 0 {
		if err := assignInputField(input, "CustomRules", _amplifyCustomRules); err != nil {
			log.Errorf("invalid --custom-rules: %s", err.Error())
			return
		}
	}
	if len(_amplifyDescription) > 0 {
		input.Description = aws.String(_amplifyDescription)
	}
	if len(_amplifyEnableAutoBranchCreation) > 0 {
		if err := assignInputField(input, "EnableAutoBranchCreation", _amplifyEnableAutoBranchCreation); err != nil {
			log.Errorf("invalid --enable-auto-branch-creation: %s", err.Error())
			return
		}
	}
	if len(_amplifyEnableBasicAuth) > 0 {
		if err := assignInputField(input, "EnableBasicAuth", _amplifyEnableBasicAuth); err != nil {
			log.Errorf("invalid --enable-basic-auth: %s", err.Error())
			return
		}
	}
	if len(_amplifyEnableBranchAutoBuild) > 0 {
		if err := assignInputField(input, "EnableBranchAutoBuild", _amplifyEnableBranchAutoBuild); err != nil {
			log.Errorf("invalid --enable-branch-auto-build: %s", err.Error())
			return
		}
	}
	if len(_amplifyEnableBranchAutoDeletion) > 0 {
		if err := assignInputField(input, "EnableBranchAutoDeletion", _amplifyEnableBranchAutoDeletion); err != nil {
			log.Errorf("invalid --enable-branch-auto-deletion: %s", err.Error())
			return
		}
	}
	if len(_amplifyEnvironmentVariables) > 0 {
		if err := assignInputField(input, "EnvironmentVariables", _amplifyEnvironmentVariables); err != nil {
			log.Errorf("invalid --environment-variables: %s", err.Error())
			return
		}
	}
	if len(_amplifyIamServiceRoleArn) > 0 {
		input.IamServiceRoleArn = aws.String(_amplifyIamServiceRoleArn)
	}
	if len(_amplifyJobConfig) > 0 {
		if err := assignInputField(input, "JobConfig", _amplifyJobConfig); err != nil {
			log.Errorf("invalid --job-config: %s", err.Error())
			return
		}
	}
	if len(_amplifyName) > 0 {
		input.Name = aws.String(_amplifyName)
	}
	if len(_amplifyOauthToken) > 0 {
		input.OauthToken = aws.String(_amplifyOauthToken)
	}
	if len(_amplifyPlatform) > 0 {
		if err := assignInputField(input, "Platform", _amplifyPlatform); err != nil {
			log.Errorf("invalid --platform: %s", err.Error())
			return
		}
	}
	if len(_amplifyRepository) > 0 {
		input.Repository = aws.String(_amplifyRepository)
	}

	if resp, err := client.UpdateApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a branch for an Amplify app.
func amplify_UpdateBranch(cfg aws.Config, client *amplify.Client) {
	input := &amplify.UpdateBranchInput{
		// AppId: *string, // Required
		// BranchName: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyBranchName) > 0 {
		input.BranchName = aws.String(_amplifyBranchName)
	}
	if len(_amplifyBackend) > 0 {
		if err := assignInputField(input, "Backend", _amplifyBackend); err != nil {
			log.Errorf("invalid --backend: %s", err.Error())
			return
		}
	}
	if len(_amplifyBackendEnvironmentArn) > 0 {
		input.BackendEnvironmentArn = aws.String(_amplifyBackendEnvironmentArn)
	}
	if len(_amplifyBasicAuthCredentials) > 0 {
		input.BasicAuthCredentials = aws.String(_amplifyBasicAuthCredentials)
	}
	if len(_amplifyBuildSpec) > 0 {
		input.BuildSpec = aws.String(_amplifyBuildSpec)
	}
	if len(_amplifyComputeRoleArn) > 0 {
		input.ComputeRoleArn = aws.String(_amplifyComputeRoleArn)
	}
	if len(_amplifyDescription) > 0 {
		input.Description = aws.String(_amplifyDescription)
	}
	if len(_amplifyDisplayName) > 0 {
		input.DisplayName = aws.String(_amplifyDisplayName)
	}
	if len(_amplifyEnableAutoBuild) > 0 {
		if err := assignInputField(input, "EnableAutoBuild", _amplifyEnableAutoBuild); err != nil {
			log.Errorf("invalid --enable-auto-build: %s", err.Error())
			return
		}
	}
	if len(_amplifyEnableBasicAuth) > 0 {
		if err := assignInputField(input, "EnableBasicAuth", _amplifyEnableBasicAuth); err != nil {
			log.Errorf("invalid --enable-basic-auth: %s", err.Error())
			return
		}
	}
	if len(_amplifyEnableNotification) > 0 {
		if err := assignInputField(input, "EnableNotification", _amplifyEnableNotification); err != nil {
			log.Errorf("invalid --enable-notification: %s", err.Error())
			return
		}
	}
	if len(_amplifyEnablePerformanceMode) > 0 {
		if err := assignInputField(input, "EnablePerformanceMode", _amplifyEnablePerformanceMode); err != nil {
			log.Errorf("invalid --enable-performance-mode: %s", err.Error())
			return
		}
	}
	if len(_amplifyEnablePullRequestPreview) > 0 {
		if err := assignInputField(input, "EnablePullRequestPreview", _amplifyEnablePullRequestPreview); err != nil {
			log.Errorf("invalid --enable-pull-request-preview: %s", err.Error())
			return
		}
	}
	if len(_amplifyEnableSkewProtection) > 0 {
		if err := assignInputField(input, "EnableSkewProtection", _amplifyEnableSkewProtection); err != nil {
			log.Errorf("invalid --enable-skew-protection: %s", err.Error())
			return
		}
	}
	if len(_amplifyEnvironmentVariables) > 0 {
		if err := assignInputField(input, "EnvironmentVariables", _amplifyEnvironmentVariables); err != nil {
			log.Errorf("invalid --environment-variables: %s", err.Error())
			return
		}
	}
	if len(_amplifyFramework) > 0 {
		input.Framework = aws.String(_amplifyFramework)
	}
	if len(_amplifyPullRequestEnvironmentName) > 0 {
		input.PullRequestEnvironmentName = aws.String(_amplifyPullRequestEnvironmentName)
	}
	if len(_amplifyStage) > 0 {
		if err := assignInputField(input, "Stage", _amplifyStage); err != nil {
			log.Errorf("invalid --stage: %s", err.Error())
			return
		}
	}
	if len(_amplifyTtl) > 0 {
		input.Ttl = aws.String(_amplifyTtl)
	}

	if resp, err := client.UpdateBranch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new domain association for an Amplify app.
func amplify_UpdateDomainAssociation(cfg aws.Config, client *amplify.Client) {
	input := &amplify.UpdateDomainAssociationInput{
		// AppId: *string, // Required
		// DomainName: *string, // Required
	}

	if len(_amplifyAppId) > 0 {
		input.AppId = aws.String(_amplifyAppId)
	}
	if len(_amplifyDomainName) > 0 {
		input.DomainName = aws.String(_amplifyDomainName)
	}
	if len(_amplifyAutoSubDomainCreationPatterns) > 0 {
		input.AutoSubDomainCreationPatterns = append([]string(nil), _amplifyAutoSubDomainCreationPatterns...)
	}
	if len(_amplifyAutoSubDomainIAMRole) > 0 {
		input.AutoSubDomainIAMRole = aws.String(_amplifyAutoSubDomainIAMRole)
	}
	if len(_amplifyCertificateSettings) > 0 {
		if err := assignInputField(input, "CertificateSettings", _amplifyCertificateSettings); err != nil {
			log.Errorf("invalid --certificate-settings: %s", err.Error())
			return
		}
	}
	if len(_amplifyEnableAutoSubDomain) > 0 {
		if err := assignInputField(input, "EnableAutoSubDomain", _amplifyEnableAutoSubDomain); err != nil {
			log.Errorf("invalid --enable-auto-sub-domain: %s", err.Error())
			return
		}
	}
	if len(_amplifySubDomainSettings) > 0 {
		if err := assignInputField(input, "SubDomainSettings", _amplifySubDomainSettings); err != nil {
			log.Errorf("invalid --sub-domain-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDomainAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a webhook.
func amplify_UpdateWebhook(cfg aws.Config, client *amplify.Client) {
	input := &amplify.UpdateWebhookInput{
		// WebhookId: *string, // Required
	}

	if len(_amplifyWebhookId) > 0 {
		input.WebhookId = aws.String(_amplifyWebhookId)
	}
	if len(_amplifyBranchName) > 0 {
		input.BranchName = aws.String(_amplifyBranchName)
	}
	if len(_amplifyDescription) > 0 {
		input.Description = aws.String(_amplifyDescription)
	}

	if resp, err := client.UpdateWebhook(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_amplifyCmd)
	_amplifyCmd.Flags().SortFlags = false

	_amplifyCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_amplifyCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_amplifyCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_amplifyCmd.Flags().StringVarP(&_amplifyAccessToken, "access-token", "", "", "Access Token")
	_amplifyCmd.Flags().StringVarP(&_amplifyAppId, "app-id", "", "", "App ID")
	_amplifyCmd.Flags().StringVarP(&_amplifyArtifactId, "artifact-id", "", "", "Artifact ID")
	_amplifyCmd.Flags().StringVarP(&_amplifyAutoBranchCreationConfig, "auto-branch-creation-config", "", "", "Auto Branch Creation Config")
	_amplifyCmd.Flags().StringSliceVarP(&_amplifyAutoBranchCreationPatterns, "auto-branch-creation-patterns", "", nil, "Auto Branch Creation Patterns")
	_amplifyCmd.Flags().StringSliceVarP(&_amplifyAutoSubDomainCreationPatterns, "auto-sub-domain-creation-patterns", "", nil, "Auto Sub Domain Creation Patterns")
	_amplifyCmd.Flags().StringVarP(&_amplifyAutoSubDomainIAMRole, "auto-sub-domain-iam-role", "", "", "Auto Sub Domain IAM Role")
	_amplifyCmd.Flags().StringVarP(&_amplifyBackend, "backend", "", "", "Backend")
	_amplifyCmd.Flags().StringVarP(&_amplifyBackendEnvironmentArn, "backend-environment-arn", "", "", "Backend Environment ARN")
	_amplifyCmd.Flags().StringVarP(&_amplifyBasicAuthCredentials, "basic-auth-credentials", "", "", "Basic Auth Credentials")
	_amplifyCmd.Flags().StringVarP(&_amplifyBranchName, "branch-name", "", "", "Branch Name")
	_amplifyCmd.Flags().StringVarP(&_amplifyBuildSpec, "build-spec", "", "", "Build Spec")
	_amplifyCmd.Flags().StringVarP(&_amplifyCacheConfig, "cache-config", "", "", "Cache Config")
	_amplifyCmd.Flags().StringVarP(&_amplifyCertificateSettings, "certificate-settings", "", "", "Certificate Settings")
	_amplifyCmd.Flags().StringVarP(&_amplifyCommitId, "commit-id", "", "", "Commit ID")
	_amplifyCmd.Flags().StringVarP(&_amplifyCommitMessage, "commit-message", "", "", "Commit Message")
	_amplifyCmd.Flags().StringVarP(&_amplifyCommitTime, "commit-time", "", "", "Commit Time")
	_amplifyCmd.Flags().StringVarP(&_amplifyComputeRoleArn, "compute-role-arn", "", "", "Compute Role ARN")
	_amplifyCmd.Flags().StringVarP(&_amplifyCustomHeaders, "custom-headers", "", "", "Custom Headers")
	_amplifyCmd.Flags().StringVarP(&_amplifyCustomRules, "custom-rules", "", "", "Custom Rules")
	_amplifyCmd.Flags().StringVarP(&_amplifyDeploymentArtifacts, "deployment-artifacts", "", "", "Deployment Artifacts")
	_amplifyCmd.Flags().StringVarP(&_amplifyDescription, "description", "", "", "Description")
	_amplifyCmd.Flags().StringVarP(&_amplifyDisplayName, "display-name", "", "", "Display Name")
	_amplifyCmd.Flags().StringVarP(&_amplifyDomainName, "domain-name", "", "", "Domain Name")
	_amplifyCmd.Flags().StringVarP(&_amplifyEnableAutoBranchCreation, "enable-auto-branch-creation", "", "", "Enable Auto Branch Creation")
	_amplifyCmd.Flags().StringVarP(&_amplifyEnableAutoBuild, "enable-auto-build", "", "", "Enable Auto Build")
	_amplifyCmd.Flags().StringVarP(&_amplifyEnableAutoSubDomain, "enable-auto-sub-domain", "", "", "Enable Auto Sub Domain")
	_amplifyCmd.Flags().StringVarP(&_amplifyEnableBasicAuth, "enable-basic-auth", "", "", "Enable Basic Auth")
	_amplifyCmd.Flags().StringVarP(&_amplifyEnableBranchAutoBuild, "enable-branch-auto-build", "", "", "Enable Branch Auto Build")
	_amplifyCmd.Flags().StringVarP(&_amplifyEnableBranchAutoDeletion, "enable-branch-auto-deletion", "", "", "Enable Branch Auto Deletion")
	_amplifyCmd.Flags().StringVarP(&_amplifyEnableNotification, "enable-notification", "", "", "Enable Notification")
	_amplifyCmd.Flags().StringVarP(&_amplifyEnablePerformanceMode, "enable-performance-mode", "", "", "Enable Performance Mode")
	_amplifyCmd.Flags().StringVarP(&_amplifyEnablePullRequestPreview, "enable-pull-request-preview", "", "", "Enable Pull Request Preview")
	_amplifyCmd.Flags().StringVarP(&_amplifyEnableSkewProtection, "enable-skew-protection", "", "", "Enable Skew Protection")
	_amplifyCmd.Flags().StringVarP(&_amplifyEndTime, "end-time", "", "", "End Time")
	_amplifyCmd.Flags().StringVarP(&_amplifyEnvironmentName, "environment-name", "", "", "Environment Name")
	_amplifyCmd.Flags().StringVarP(&_amplifyEnvironmentVariables, "environment-variables", "", "", "Environment Variables")
	_amplifyCmd.Flags().StringVarP(&_amplifyFileMap, "file-map", "", "", "File Map")
	_amplifyCmd.Flags().StringVarP(&_amplifyFramework, "framework", "", "", "Framework")
	_amplifyCmd.Flags().StringVarP(&_amplifyIamServiceRoleArn, "iam-service-role-arn", "", "", "IAM Service Role ARN")
	_amplifyCmd.Flags().StringVarP(&_amplifyJobConfig, "job-config", "", "", "Job Config")
	_amplifyCmd.Flags().StringVarP(&_amplifyJobId, "job-id", "", "", "Job ID")
	_amplifyCmd.Flags().StringVarP(&_amplifyJobReason, "job-reason", "", "", "Job Reason")
	_amplifyCmd.Flags().StringVarP(&_amplifyJobType, "job-type", "", "", "Job Type")
	_amplifyCmd.Flags().StringVarP(&_amplifyMaxResults, "max-results", "", "", "Max Results")
	_amplifyCmd.Flags().StringVarP(&_amplifyName, "name", "", "", "Name")
	_amplifyCmd.Flags().StringVarP(&_amplifyNextToken, "next-token", "", "", "Next Token")
	_amplifyCmd.Flags().StringVarP(&_amplifyOauthToken, "oauth-token", "", "", "Oauth Token")
	_amplifyCmd.Flags().StringVarP(&_amplifyPlatform, "platform", "", "", "Platform")
	_amplifyCmd.Flags().StringVarP(&_amplifyPullRequestEnvironmentName, "pull-request-environment-name", "", "", "Pull Request Environment Name")
	_amplifyCmd.Flags().StringVarP(&_amplifyRepository, "repository", "", "", "Repository")
	_amplifyCmd.Flags().StringVarP(&_amplifyResourceArn, "resource-arn", "", "", "Resource ARN")
	_amplifyCmd.Flags().StringVarP(&_amplifySourceUrl, "source-url", "", "", "Source URL")
	_amplifyCmd.Flags().StringVarP(&_amplifySourceUrlType, "source-url-type", "", "", "Source URL Type")
	_amplifyCmd.Flags().StringVarP(&_amplifyStackName, "stack-name", "", "", "Stack Name")
	_amplifyCmd.Flags().StringVarP(&_amplifyStage, "stage", "", "", "Stage")
	_amplifyCmd.Flags().StringVarP(&_amplifyStartTime, "start-time", "", "", "Start Time")
	_amplifyCmd.Flags().StringVarP(&_amplifySubDomainSettings, "sub-domain-settings", "", "", "Sub Domain Settings")
	_amplifyCmd.Flags().StringSliceVarP(&_amplifyTagKeys, "tag-keys", "", nil, "Tag Keys")
	_amplifyCmd.Flags().StringVarP(&_amplifyTags, "tags", "", "", "Tags")
	_amplifyCmd.Flags().StringVarP(&_amplifyTtl, "ttl", "", "", "TTL")
	_amplifyCmd.Flags().StringVarP(&_amplifyWebhookId, "webhook-id", "", "", "Webhook ID")

	_amplifyCmd.Flags().BoolVarP(&_amplifyCreateApp, "create-app", "", false, "Create App")
	_amplifyCmd.Flags().BoolVarP(&_amplifyCreateBackendEnvironment, "create-backend-environment", "", false, "Create Backend Environment")
	_amplifyCmd.Flags().BoolVarP(&_amplifyCreateBranch, "create-branch", "", false, "Create Branch")
	_amplifyCmd.Flags().BoolVarP(&_amplifyCreateDeployment, "create-deployment", "", false, "Create Deployment")
	_amplifyCmd.Flags().BoolVarP(&_amplifyCreateDomainAssociation, "create-domain-association", "", false, "Create Domain Association")
	_amplifyCmd.Flags().BoolVarP(&_amplifyCreateWebhook, "create-webhook", "", false, "Create Webhook")
	_amplifyCmd.Flags().BoolVarP(&_amplifyDeleteApp, "delete-app", "", false, "Delete App")
	_amplifyCmd.Flags().BoolVarP(&_amplifyDeleteBackendEnvironment, "delete-backend-environment", "", false, "Delete Backend Environment")
	_amplifyCmd.Flags().BoolVarP(&_amplifyDeleteBranch, "delete-branch", "", false, "Delete Branch")
	_amplifyCmd.Flags().BoolVarP(&_amplifyDeleteDomainAssociation, "delete-domain-association", "", false, "Delete Domain Association")
	_amplifyCmd.Flags().BoolVarP(&_amplifyDeleteJob, "delete-job", "", false, "Delete Job")
	_amplifyCmd.Flags().BoolVarP(&_amplifyDeleteWebhook, "delete-webhook", "", false, "Delete Webhook")
	_amplifyCmd.Flags().BoolVarP(&_amplifyGenerateAccessLogs, "generate-access-logs", "", false, "Generate Access Logs")
	_amplifyCmd.Flags().BoolVarP(&_amplifyGetApp, "get-app", "", false, "Get App")
	_amplifyCmd.Flags().BoolVarP(&_amplifyGetArtifactUrl, "get-artifact-url", "", false, "Get Artifact URL")
	_amplifyCmd.Flags().BoolVarP(&_amplifyGetBackendEnvironment, "get-backend-environment", "", false, "Get Backend Environment")
	_amplifyCmd.Flags().BoolVarP(&_amplifyGetBranch, "get-branch", "", false, "Get Branch")
	_amplifyCmd.Flags().BoolVarP(&_amplifyGetDomainAssociation, "get-domain-association", "", false, "Get Domain Association")
	_amplifyCmd.Flags().BoolVarP(&_amplifyGetJob, "get-job", "", false, "Get Job")
	_amplifyCmd.Flags().BoolVarP(&_amplifyGetWebhook, "get-webhook", "", false, "Get Webhook")
	_amplifyCmd.Flags().BoolVarP(&_amplifyListApps, "list-apps", "", false, "List Apps")
	_amplifyCmd.Flags().BoolVarP(&_amplifyListArtifacts, "list-artifacts", "", false, "List Artifacts")
	_amplifyCmd.Flags().BoolVarP(&_amplifyListBackendEnvironments, "list-backend-environments", "", false, "List Backend Environments")
	_amplifyCmd.Flags().BoolVarP(&_amplifyListBranches, "list-branches", "", false, "List Branches")
	_amplifyCmd.Flags().BoolVarP(&_amplifyListDomainAssociations, "list-domain-associations", "", false, "List Domain Associations")
	_amplifyCmd.Flags().BoolVarP(&_amplifyListJobs, "list-jobs", "", false, "List Jobs")
	_amplifyCmd.Flags().BoolVarP(&_amplifyListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_amplifyCmd.Flags().BoolVarP(&_amplifyListWebhooks, "list-webhooks", "", false, "List Webhooks")
	_amplifyCmd.Flags().BoolVarP(&_amplifyStartDeployment, "start-deployment", "", false, "Start Deployment")
	_amplifyCmd.Flags().BoolVarP(&_amplifyStartJob, "start-job", "", false, "Start Job")
	_amplifyCmd.Flags().BoolVarP(&_amplifyStopJob, "stop-job", "", false, "Stop Job")
	_amplifyCmd.Flags().BoolVarP(&_amplifyTagResource, "tag-resource", "", false, "Tag Resource")
	_amplifyCmd.Flags().BoolVarP(&_amplifyUntagResource, "untag-resource", "", false, "Untag Resource")
	_amplifyCmd.Flags().BoolVarP(&_amplifyUpdateApp, "update-app", "", false, "Update App")
	_amplifyCmd.Flags().BoolVarP(&_amplifyUpdateBranch, "update-branch", "", false, "Update Branch")
	_amplifyCmd.Flags().BoolVarP(&_amplifyUpdateDomainAssociation, "update-domain-association", "", false, "Update Domain Association")
	_amplifyCmd.Flags().BoolVarP(&_amplifyUpdateWebhook, "update-webhook", "", false, "Update Webhook")

}
