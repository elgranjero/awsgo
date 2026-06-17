package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// elasticbeanstalkCmd represents the elasticbeanstalk command
var _elasticbeanstalkCmd = &cobra.Command{
	Use:   "elasticbeanstalk",
	Short: "AWS elasticbeanstalk CLI",
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
		client := elasticbeanstalk.NewFromConfig(cfg)
		if _elasticbeanstalkAbortEnvironmentUpdate {
			elasticbeanstalk_AbortEnvironmentUpdate(cfg, client)
			return
		}
		if _elasticbeanstalkApplyEnvironmentManagedAction {
			elasticbeanstalk_ApplyEnvironmentManagedAction(cfg, client)
			return
		}
		if _elasticbeanstalkAssociateEnvironmentOperationsRole {
			elasticbeanstalk_AssociateEnvironmentOperationsRole(cfg, client)
			return
		}
		if _elasticbeanstalkCheckDNSAvailability {
			elasticbeanstalk_CheckDNSAvailability(cfg, client)
			return
		}
		if _elasticbeanstalkComposeEnvironments {
			elasticbeanstalk_ComposeEnvironments(cfg, client)
			return
		}
		if _elasticbeanstalkCreateApplication {
			elasticbeanstalk_CreateApplication(cfg, client)
			return
		}
		if _elasticbeanstalkCreateApplicationVersion {
			elasticbeanstalk_CreateApplicationVersion(cfg, client)
			return
		}
		if _elasticbeanstalkCreateConfigurationTemplate {
			elasticbeanstalk_CreateConfigurationTemplate(cfg, client)
			return
		}
		if _elasticbeanstalkCreateEnvironment {
			elasticbeanstalk_CreateEnvironment(cfg, client)
			return
		}
		if _elasticbeanstalkCreatePlatformVersion {
			elasticbeanstalk_CreatePlatformVersion(cfg, client)
			return
		}
		if _elasticbeanstalkCreateStorageLocation {
			elasticbeanstalk_CreateStorageLocation(cfg, client)
			return
		}
		if _elasticbeanstalkDeleteApplication {
			elasticbeanstalk_DeleteApplication(cfg, client)
			return
		}
		if _elasticbeanstalkDeleteApplicationVersion {
			elasticbeanstalk_DeleteApplicationVersion(cfg, client)
			return
		}
		if _elasticbeanstalkDeleteConfigurationTemplate {
			elasticbeanstalk_DeleteConfigurationTemplate(cfg, client)
			return
		}
		if _elasticbeanstalkDeleteEnvironmentConfiguration {
			elasticbeanstalk_DeleteEnvironmentConfiguration(cfg, client)
			return
		}
		if _elasticbeanstalkDeletePlatformVersion {
			elasticbeanstalk_DeletePlatformVersion(cfg, client)
			return
		}
		if _elasticbeanstalkDescribeAccountAttributes {
			elasticbeanstalk_DescribeAccountAttributes(cfg, client)
			return
		}
		if _elasticbeanstalkDescribeApplicationVersions {
			elasticbeanstalk_DescribeApplicationVersions(cfg, client)
			return
		}
		if _elasticbeanstalkDescribeApplications {
			elasticbeanstalk_DescribeApplications(cfg, client)
			return
		}
		if _elasticbeanstalkDescribeConfigurationOptions {
			elasticbeanstalk_DescribeConfigurationOptions(cfg, client)
			return
		}
		if _elasticbeanstalkDescribeConfigurationSettings {
			elasticbeanstalk_DescribeConfigurationSettings(cfg, client)
			return
		}
		if _elasticbeanstalkDescribeEnvironmentHealth {
			elasticbeanstalk_DescribeEnvironmentHealth(cfg, client)
			return
		}
		if _elasticbeanstalkDescribeEnvironmentManagedActionHistory {
			elasticbeanstalk_DescribeEnvironmentManagedActionHistory(cfg, client)
			return
		}
		if _elasticbeanstalkDescribeEnvironmentManagedActions {
			elasticbeanstalk_DescribeEnvironmentManagedActions(cfg, client)
			return
		}
		if _elasticbeanstalkDescribeEnvironmentResources {
			elasticbeanstalk_DescribeEnvironmentResources(cfg, client)
			return
		}
		if _elasticbeanstalkDescribeEnvironments {
			elasticbeanstalk_DescribeEnvironments(cfg, client)
			return
		}
		if _elasticbeanstalkDescribeEvents {
			elasticbeanstalk_DescribeEvents(cfg, client)
			return
		}
		if _elasticbeanstalkDescribeInstancesHealth {
			elasticbeanstalk_DescribeInstancesHealth(cfg, client)
			return
		}
		if _elasticbeanstalkDescribePlatformVersion {
			elasticbeanstalk_DescribePlatformVersion(cfg, client)
			return
		}
		if _elasticbeanstalkDisassociateEnvironmentOperationsRole {
			elasticbeanstalk_DisassociateEnvironmentOperationsRole(cfg, client)
			return
		}
		if _elasticbeanstalkListAvailableSolutionStacks {
			elasticbeanstalk_ListAvailableSolutionStacks(cfg, client)
			return
		}
		if _elasticbeanstalkListPlatformBranches {
			elasticbeanstalk_ListPlatformBranches(cfg, client)
			return
		}
		if _elasticbeanstalkListPlatformVersions {
			elasticbeanstalk_ListPlatformVersions(cfg, client)
			return
		}
		if _elasticbeanstalkListTagsForResource {
			elasticbeanstalk_ListTagsForResource(cfg, client)
			return
		}
		if _elasticbeanstalkRebuildEnvironment {
			elasticbeanstalk_RebuildEnvironment(cfg, client)
			return
		}
		if _elasticbeanstalkRequestEnvironmentInfo {
			elasticbeanstalk_RequestEnvironmentInfo(cfg, client)
			return
		}
		if _elasticbeanstalkRestartAppServer {
			elasticbeanstalk_RestartAppServer(cfg, client)
			return
		}
		if _elasticbeanstalkRetrieveEnvironmentInfo {
			elasticbeanstalk_RetrieveEnvironmentInfo(cfg, client)
			return
		}
		if _elasticbeanstalkSwapEnvironmentCNAMEs {
			elasticbeanstalk_SwapEnvironmentCNAMEs(cfg, client)
			return
		}
		if _elasticbeanstalkTerminateEnvironment {
			elasticbeanstalk_TerminateEnvironment(cfg, client)
			return
		}
		if _elasticbeanstalkUpdateApplication {
			elasticbeanstalk_UpdateApplication(cfg, client)
			return
		}
		if _elasticbeanstalkUpdateApplicationResourceLifecycle {
			elasticbeanstalk_UpdateApplicationResourceLifecycle(cfg, client)
			return
		}
		if _elasticbeanstalkUpdateApplicationVersion {
			elasticbeanstalk_UpdateApplicationVersion(cfg, client)
			return
		}
		if _elasticbeanstalkUpdateConfigurationTemplate {
			elasticbeanstalk_UpdateConfigurationTemplate(cfg, client)
			return
		}
		if _elasticbeanstalkUpdateEnvironment {
			elasticbeanstalk_UpdateEnvironment(cfg, client)
			return
		}
		if _elasticbeanstalkUpdateTagsForResource {
			elasticbeanstalk_UpdateTagsForResource(cfg, client)
			return
		}
		if _elasticbeanstalkValidateConfigurationSettings {
			elasticbeanstalk_ValidateConfigurationSettings(cfg, client)
			return
		}

	},
}

var (
	_elasticbeanstalkAbortEnvironmentUpdate                  bool
	_elasticbeanstalkApplyEnvironmentManagedAction           bool
	_elasticbeanstalkAssociateEnvironmentOperationsRole      bool
	_elasticbeanstalkCheckDNSAvailability                    bool
	_elasticbeanstalkComposeEnvironments                     bool
	_elasticbeanstalkCreateApplication                       bool
	_elasticbeanstalkCreateApplicationVersion                bool
	_elasticbeanstalkCreateConfigurationTemplate             bool
	_elasticbeanstalkCreateEnvironment                       bool
	_elasticbeanstalkCreatePlatformVersion                   bool
	_elasticbeanstalkCreateStorageLocation                   bool
	_elasticbeanstalkDeleteApplication                       bool
	_elasticbeanstalkDeleteApplicationVersion                bool
	_elasticbeanstalkDeleteConfigurationTemplate             bool
	_elasticbeanstalkDeleteEnvironmentConfiguration          bool
	_elasticbeanstalkDeletePlatformVersion                   bool
	_elasticbeanstalkDescribeAccountAttributes               bool
	_elasticbeanstalkDescribeApplicationVersions             bool
	_elasticbeanstalkDescribeApplications                    bool
	_elasticbeanstalkDescribeConfigurationOptions            bool
	_elasticbeanstalkDescribeConfigurationSettings           bool
	_elasticbeanstalkDescribeEnvironmentHealth               bool
	_elasticbeanstalkDescribeEnvironmentManagedActionHistory bool
	_elasticbeanstalkDescribeEnvironmentManagedActions       bool
	_elasticbeanstalkDescribeEnvironmentResources            bool
	_elasticbeanstalkDescribeEnvironments                    bool
	_elasticbeanstalkDescribeEvents                          bool
	_elasticbeanstalkDescribeInstancesHealth                 bool
	_elasticbeanstalkDescribePlatformVersion                 bool
	_elasticbeanstalkDisassociateEnvironmentOperationsRole   bool
	_elasticbeanstalkListAvailableSolutionStacks             bool
	_elasticbeanstalkListPlatformBranches                    bool
	_elasticbeanstalkListPlatformVersions                    bool
	_elasticbeanstalkListTagsForResource                     bool
	_elasticbeanstalkRebuildEnvironment                      bool
	_elasticbeanstalkRequestEnvironmentInfo                  bool
	_elasticbeanstalkRestartAppServer                        bool
	_elasticbeanstalkRetrieveEnvironmentInfo                 bool
	_elasticbeanstalkSwapEnvironmentCNAMEs                   bool
	_elasticbeanstalkTerminateEnvironment                    bool
	_elasticbeanstalkUpdateApplication                       bool
	_elasticbeanstalkUpdateApplicationResourceLifecycle      bool
	_elasticbeanstalkUpdateApplicationVersion                bool
	_elasticbeanstalkUpdateConfigurationTemplate             bool
	_elasticbeanstalkUpdateEnvironment                       bool
	_elasticbeanstalkUpdateTagsForResource                   bool
	_elasticbeanstalkValidateConfigurationSettings           bool

	_elasticbeanstalkActionId                   string
	_elasticbeanstalkApplicationName            string
	_elasticbeanstalkApplicationNames           []string
	_elasticbeanstalkAttributeNames             string
	_elasticbeanstalkAutoCreateApplication      string
	_elasticbeanstalkBuildConfiguration         string
	_elasticbeanstalkCNAMEPrefix                string
	_elasticbeanstalkDeleteSourceBundle         string
	_elasticbeanstalkDescription                string
	_elasticbeanstalkDestinationEnvironmentId   string
	_elasticbeanstalkDestinationEnvironmentName string
	_elasticbeanstalkEndTime                    string
	_elasticbeanstalkEnvironmentId              string
	_elasticbeanstalkEnvironmentIds             []string
	_elasticbeanstalkEnvironmentName            string
	_elasticbeanstalkEnvironmentNames           []string
	_elasticbeanstalkFilters                    string
	_elasticbeanstalkForceTerminate             string
	_elasticbeanstalkGroupName                  string
	_elasticbeanstalkIncludeDeleted             string
	_elasticbeanstalkIncludedDeletedBackTo      string
	_elasticbeanstalkInfoType                   string
	_elasticbeanstalkMaxItems                   string
	_elasticbeanstalkMaxRecords                 string
	_elasticbeanstalkNextToken                  string
	_elasticbeanstalkOperationsRole             string
	_elasticbeanstalkOptionSettings             string
	_elasticbeanstalkOptions                    string
	_elasticbeanstalkOptionsToRemove            string
	_elasticbeanstalkPlatformArn                string
	_elasticbeanstalkPlatformDefinitionBundle   string
	_elasticbeanstalkPlatformName               string
	_elasticbeanstalkPlatformVersion            string
	_elasticbeanstalkProcess                    string
	_elasticbeanstalkRequestId                  string
	_elasticbeanstalkResourceArn                string
	_elasticbeanstalkResourceLifecycleConfig    string
	_elasticbeanstalkSeverity                   string
	_elasticbeanstalkSolutionStackName          string
	_elasticbeanstalkSourceBuildInformation     string
	_elasticbeanstalkSourceBundle               string
	_elasticbeanstalkSourceConfiguration        string
	_elasticbeanstalkSourceEnvironmentId        string
	_elasticbeanstalkSourceEnvironmentName      string
	_elasticbeanstalkStartTime                  string
	_elasticbeanstalkStatus                     string
	_elasticbeanstalkTags                       string
	_elasticbeanstalkTagsToAdd                  string
	_elasticbeanstalkTagsToRemove               []string
	_elasticbeanstalkTemplateName               string
	_elasticbeanstalkTerminateEnvByForce        string
	_elasticbeanstalkTerminateResources         string
	_elasticbeanstalkTier                       string
	_elasticbeanstalkVersionLabel               string
	_elasticbeanstalkVersionLabels              []string
)

// Cancels in-progress environment configuration update or application version
// deployment.
func elasticbeanstalk_AbortEnvironmentUpdate(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.AbortEnvironmentUpdateInput{}

	if len(_elasticbeanstalkEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_elasticbeanstalkEnvironmentId)
	}
	if len(_elasticbeanstalkEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_elasticbeanstalkEnvironmentName)
	}

	if resp, err := client.AbortEnvironmentUpdate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies a scheduled managed action immediately. A managed action can be applied
// only if its status is Scheduled . Get the status and action ID of a managed
// action with DescribeEnvironmentManagedActions.
func elasticbeanstalk_ApplyEnvironmentManagedAction(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.ApplyEnvironmentManagedActionInput{
		// ActionId: *string, // Required
	}

	if len(_elasticbeanstalkActionId) > 0 {
		input.ActionId = aws.String(_elasticbeanstalkActionId)
	}
	if len(_elasticbeanstalkEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_elasticbeanstalkEnvironmentId)
	}
	if len(_elasticbeanstalkEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_elasticbeanstalkEnvironmentName)
	}

	if resp, err := client.ApplyEnvironmentManagedAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add or change the operations role used by an environment. After this call is
// made, Elastic Beanstalk uses the associated operations role for permissions to
// downstream services during subsequent calls acting on this environment. For more
// information, see [Operations roles]in the AWS Elastic Beanstalk Developer Guide.
//
// [Operations roles]: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/iam-operationsrole.html
func elasticbeanstalk_AssociateEnvironmentOperationsRole(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.AssociateEnvironmentOperationsRoleInput{
		// EnvironmentName: *string, // Required
		// OperationsRole: *string, // Required
	}

	if len(_elasticbeanstalkEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_elasticbeanstalkEnvironmentName)
	}
	if len(_elasticbeanstalkOperationsRole) > 0 {
		input.OperationsRole = aws.String(_elasticbeanstalkOperationsRole)
	}

	if resp, err := client.AssociateEnvironmentOperationsRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Checks if the specified CNAME is available.
func elasticbeanstalk_CheckDNSAvailability(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.CheckDNSAvailabilityInput{
		// CNAMEPrefix: *string, // Required
	}

	if len(_elasticbeanstalkCNAMEPrefix) > 0 {
		input.CNAMEPrefix = aws.String(_elasticbeanstalkCNAMEPrefix)
	}

	if resp, err := client.CheckDNSAvailability(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create or update a group of environments that each run a separate component of
// a single application. Takes a list of version labels that specify application
// source bundles for each of the environments to create or update. The name of
// each environment and other required information must be included in the source
// bundles in an environment manifest named env.yaml . See [Compose Environments] for details.
//
// [Compose Environments]: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/environment-mgmt-compose.html
func elasticbeanstalk_ComposeEnvironments(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.ComposeEnvironmentsInput{}

	if len(_elasticbeanstalkApplicationName) > 0 {
		input.ApplicationName = aws.String(_elasticbeanstalkApplicationName)
	}
	if len(_elasticbeanstalkGroupName) > 0 {
		input.GroupName = aws.String(_elasticbeanstalkGroupName)
	}
	if len(_elasticbeanstalkVersionLabels) > 0 {
		input.VersionLabels = append([]string(nil), _elasticbeanstalkVersionLabels...)
	}

	if resp, err := client.ComposeEnvironments(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an application that has one configuration template named default and no
// application versions.
func elasticbeanstalk_CreateApplication(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.CreateApplicationInput{
		// ApplicationName: *string, // Required
	}

	if len(_elasticbeanstalkApplicationName) > 0 {
		input.ApplicationName = aws.String(_elasticbeanstalkApplicationName)
	}
	if len(_elasticbeanstalkDescription) > 0 {
		input.Description = aws.String(_elasticbeanstalkDescription)
	}
	if len(_elasticbeanstalkResourceLifecycleConfig) > 0 {
		if err := assignInputField(input, "ResourceLifecycleConfig", _elasticbeanstalkResourceLifecycleConfig); err != nil {
			log.Errorf("invalid --resource-lifecycle-config: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkTags) > 0 {
		if err := assignInputField(input, "Tags", _elasticbeanstalkTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an application version for the specified application. You can create an
// application version from a source bundle in Amazon S3, a commit in AWS
// CodeCommit, or the output of an AWS CodeBuild build as follows:
//
// Specify a commit in an AWS CodeCommit repository with SourceBuildInformation .
//
// Specify a build in an AWS CodeBuild with SourceBuildInformation and
// BuildConfiguration .
//
// # Specify a source bundle in S3 with SourceBundle
//
// Omit both SourceBuildInformation and SourceBundle to use the default sample
// application.
//
// After you create an application version with a specified Amazon S3 bucket and
// key location, you can't change that Amazon S3 location. If you change the Amazon
// S3 location, you receive an exception when you attempt to launch an environment
// from the application version.
func elasticbeanstalk_CreateApplicationVersion(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.CreateApplicationVersionInput{
		// ApplicationName: *string, // Required
		// VersionLabel: *string, // Required
	}

	if len(_elasticbeanstalkApplicationName) > 0 {
		input.ApplicationName = aws.String(_elasticbeanstalkApplicationName)
	}
	if len(_elasticbeanstalkVersionLabel) > 0 {
		input.VersionLabel = aws.String(_elasticbeanstalkVersionLabel)
	}
	if len(_elasticbeanstalkAutoCreateApplication) > 0 {
		if err := assignInputField(input, "AutoCreateApplication", _elasticbeanstalkAutoCreateApplication); err != nil {
			log.Errorf("invalid --auto-create-application: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkBuildConfiguration) > 0 {
		if err := assignInputField(input, "BuildConfiguration", _elasticbeanstalkBuildConfiguration); err != nil {
			log.Errorf("invalid --build-configuration: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkDescription) > 0 {
		input.Description = aws.String(_elasticbeanstalkDescription)
	}
	if len(_elasticbeanstalkProcess) > 0 {
		if err := assignInputField(input, "Process", _elasticbeanstalkProcess); err != nil {
			log.Errorf("invalid --process: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkSourceBuildInformation) > 0 {
		if err := assignInputField(input, "SourceBuildInformation", _elasticbeanstalkSourceBuildInformation); err != nil {
			log.Errorf("invalid --source-build-information: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkSourceBundle) > 0 {
		if err := assignInputField(input, "SourceBundle", _elasticbeanstalkSourceBundle); err != nil {
			log.Errorf("invalid --source-bundle: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkTags) > 0 {
		if err := assignInputField(input, "Tags", _elasticbeanstalkTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateApplicationVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an AWS Elastic Beanstalk configuration template, associated with a
// specific Elastic Beanstalk application. You define application configuration
// settings in a configuration template. You can then use the configuration
// template to deploy different versions of the application with the same
// configuration settings.
//
// Templates aren't associated with any environment. The EnvironmentName response
// element is always null .
//
// # Related Topics
//
// # DescribeConfigurationOptions
//
// # DescribeConfigurationSettings
//
// ListAvailableSolutionStacks
func elasticbeanstalk_CreateConfigurationTemplate(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.CreateConfigurationTemplateInput{
		// ApplicationName: *string, // Required
		// TemplateName: *string, // Required
	}

	if len(_elasticbeanstalkApplicationName) > 0 {
		input.ApplicationName = aws.String(_elasticbeanstalkApplicationName)
	}
	if len(_elasticbeanstalkTemplateName) > 0 {
		input.TemplateName = aws.String(_elasticbeanstalkTemplateName)
	}
	if len(_elasticbeanstalkDescription) > 0 {
		input.Description = aws.String(_elasticbeanstalkDescription)
	}
	if len(_elasticbeanstalkEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_elasticbeanstalkEnvironmentId)
	}
	if len(_elasticbeanstalkOptionSettings) > 0 {
		if err := assignInputField(input, "OptionSettings", _elasticbeanstalkOptionSettings); err != nil {
			log.Errorf("invalid --option-settings: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkPlatformArn) > 0 {
		input.PlatformArn = aws.String(_elasticbeanstalkPlatformArn)
	}
	if len(_elasticbeanstalkSolutionStackName) > 0 {
		input.SolutionStackName = aws.String(_elasticbeanstalkSolutionStackName)
	}
	if len(_elasticbeanstalkSourceConfiguration) > 0 {
		if err := assignInputField(input, "SourceConfiguration", _elasticbeanstalkSourceConfiguration); err != nil {
			log.Errorf("invalid --source-configuration: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkTags) > 0 {
		if err := assignInputField(input, "Tags", _elasticbeanstalkTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConfigurationTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Launches an AWS Elastic Beanstalk environment for the specified application
// using the specified configuration.
func elasticbeanstalk_CreateEnvironment(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.CreateEnvironmentInput{
		// ApplicationName: *string, // Required
	}

	if len(_elasticbeanstalkApplicationName) > 0 {
		input.ApplicationName = aws.String(_elasticbeanstalkApplicationName)
	}
	if len(_elasticbeanstalkCNAMEPrefix) > 0 {
		input.CNAMEPrefix = aws.String(_elasticbeanstalkCNAMEPrefix)
	}
	if len(_elasticbeanstalkDescription) > 0 {
		input.Description = aws.String(_elasticbeanstalkDescription)
	}
	if len(_elasticbeanstalkEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_elasticbeanstalkEnvironmentName)
	}
	if len(_elasticbeanstalkGroupName) > 0 {
		input.GroupName = aws.String(_elasticbeanstalkGroupName)
	}
	if len(_elasticbeanstalkOperationsRole) > 0 {
		input.OperationsRole = aws.String(_elasticbeanstalkOperationsRole)
	}
	if len(_elasticbeanstalkOptionSettings) > 0 {
		if err := assignInputField(input, "OptionSettings", _elasticbeanstalkOptionSettings); err != nil {
			log.Errorf("invalid --option-settings: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkOptionsToRemove) > 0 {
		if err := assignInputField(input, "OptionsToRemove", _elasticbeanstalkOptionsToRemove); err != nil {
			log.Errorf("invalid --options-to-remove: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkPlatformArn) > 0 {
		input.PlatformArn = aws.String(_elasticbeanstalkPlatformArn)
	}
	if len(_elasticbeanstalkSolutionStackName) > 0 {
		input.SolutionStackName = aws.String(_elasticbeanstalkSolutionStackName)
	}
	if len(_elasticbeanstalkTags) > 0 {
		if err := assignInputField(input, "Tags", _elasticbeanstalkTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkTemplateName) > 0 {
		input.TemplateName = aws.String(_elasticbeanstalkTemplateName)
	}
	if len(_elasticbeanstalkTier) > 0 {
		if err := assignInputField(input, "Tier", _elasticbeanstalkTier); err != nil {
			log.Errorf("invalid --tier: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkVersionLabel) > 0 {
		input.VersionLabel = aws.String(_elasticbeanstalkVersionLabel)
	}

	if resp, err := client.CreateEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new version of your custom platform.
func elasticbeanstalk_CreatePlatformVersion(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.CreatePlatformVersionInput{
		// PlatformDefinitionBundle: *types.S3Location, // Required
		// PlatformName: *string, // Required
		// PlatformVersion: *string, // Required
	}

	if len(_elasticbeanstalkPlatformDefinitionBundle) > 0 {
		if err := assignInputField(input, "PlatformDefinitionBundle", _elasticbeanstalkPlatformDefinitionBundle); err != nil {
			log.Errorf("invalid --platform-definition-bundle: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkPlatformName) > 0 {
		input.PlatformName = aws.String(_elasticbeanstalkPlatformName)
	}
	if len(_elasticbeanstalkPlatformVersion) > 0 {
		input.PlatformVersion = aws.String(_elasticbeanstalkPlatformVersion)
	}
	if len(_elasticbeanstalkEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_elasticbeanstalkEnvironmentName)
	}
	if len(_elasticbeanstalkOptionSettings) > 0 {
		if err := assignInputField(input, "OptionSettings", _elasticbeanstalkOptionSettings); err != nil {
			log.Errorf("invalid --option-settings: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkTags) > 0 {
		if err := assignInputField(input, "Tags", _elasticbeanstalkTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePlatformVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a bucket in Amazon S3 to store application versions, logs, and other
// files used by Elastic Beanstalk environments. The Elastic Beanstalk console and
// EB CLI call this API the first time you create an environment in a region. If
// the storage location already exists, CreateStorageLocation still returns the
// bucket name but does not create a new bucket.
func elasticbeanstalk_CreateStorageLocation(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.CreateStorageLocationInput{}

	if resp, err := client.CreateStorageLocation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified application along with all associated versions and
// configurations. The application versions will not be deleted from your Amazon S3
// bucket.
//
// You cannot delete an application that has a running environment.
func elasticbeanstalk_DeleteApplication(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.DeleteApplicationInput{
		// ApplicationName: *string, // Required
	}

	if len(_elasticbeanstalkApplicationName) > 0 {
		input.ApplicationName = aws.String(_elasticbeanstalkApplicationName)
	}
	if len(_elasticbeanstalkTerminateEnvByForce) > 0 {
		if err := assignInputField(input, "TerminateEnvByForce", _elasticbeanstalkTerminateEnvByForce); err != nil {
			log.Errorf("invalid --terminate-env-by-force: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified version from the specified application.
// You cannot delete an application version that is associated with a running
// environment.
func elasticbeanstalk_DeleteApplicationVersion(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.DeleteApplicationVersionInput{
		// ApplicationName: *string, // Required
		// VersionLabel: *string, // Required
	}

	if len(_elasticbeanstalkApplicationName) > 0 {
		input.ApplicationName = aws.String(_elasticbeanstalkApplicationName)
	}
	if len(_elasticbeanstalkVersionLabel) > 0 {
		input.VersionLabel = aws.String(_elasticbeanstalkVersionLabel)
	}
	if len(_elasticbeanstalkDeleteSourceBundle) > 0 {
		if err := assignInputField(input, "DeleteSourceBundle", _elasticbeanstalkDeleteSourceBundle); err != nil {
			log.Errorf("invalid --delete-source-bundle: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteApplicationVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified configuration template.
// When you launch an environment using a configuration template, the environment
// gets a copy of the template. You can delete or modify the environment's copy of
// the template without affecting the running environment.
func elasticbeanstalk_DeleteConfigurationTemplate(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.DeleteConfigurationTemplateInput{
		// ApplicationName: *string, // Required
		// TemplateName: *string, // Required
	}

	if len(_elasticbeanstalkApplicationName) > 0 {
		input.ApplicationName = aws.String(_elasticbeanstalkApplicationName)
	}
	if len(_elasticbeanstalkTemplateName) > 0 {
		input.TemplateName = aws.String(_elasticbeanstalkTemplateName)
	}

	if resp, err := client.DeleteConfigurationTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the draft configuration associated with the running environment.
// Updating a running environment with any configuration changes creates a draft
// configuration set. You can get the draft configuration using DescribeConfigurationSettingswhile the update
// is in progress or if the update fails. The DeploymentStatus for the draft
// configuration indicates whether the deployment is in process or has failed. The
// draft configuration remains in existence until it is deleted with this action.
func elasticbeanstalk_DeleteEnvironmentConfiguration(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.DeleteEnvironmentConfigurationInput{
		// ApplicationName: *string, // Required
		// EnvironmentName: *string, // Required
	}

	if len(_elasticbeanstalkApplicationName) > 0 {
		input.ApplicationName = aws.String(_elasticbeanstalkApplicationName)
	}
	if len(_elasticbeanstalkEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_elasticbeanstalkEnvironmentName)
	}

	if resp, err := client.DeleteEnvironmentConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified version of a custom platform.
func elasticbeanstalk_DeletePlatformVersion(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.DeletePlatformVersionInput{}

	if len(_elasticbeanstalkPlatformArn) > 0 {
		input.PlatformArn = aws.String(_elasticbeanstalkPlatformArn)
	}

	if resp, err := client.DeletePlatformVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns attributes related to AWS Elastic Beanstalk that are associated with
// the calling AWS account.
//
// The result currently has one set of attributes—resource quotas.
func elasticbeanstalk_DescribeAccountAttributes(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.DescribeAccountAttributesInput{}

	if resp, err := client.DescribeAccountAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve a list of application versions.
func elasticbeanstalk_DescribeApplicationVersions(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.DescribeApplicationVersionsInput{}

	if len(_elasticbeanstalkApplicationName) > 0 {
		input.ApplicationName = aws.String(_elasticbeanstalkApplicationName)
	}
	if len(_elasticbeanstalkMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _elasticbeanstalkMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkNextToken) > 0 {
		input.NextToken = aws.String(_elasticbeanstalkNextToken)
	}
	if len(_elasticbeanstalkVersionLabels) > 0 {
		input.VersionLabels = append([]string(nil), _elasticbeanstalkVersionLabels...)
	}

	if resp, err := client.DescribeApplicationVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the descriptions of existing applications.
func elasticbeanstalk_DescribeApplications(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.DescribeApplicationsInput{}

	if len(_elasticbeanstalkApplicationNames) > 0 {
		input.ApplicationNames = append([]string(nil), _elasticbeanstalkApplicationNames...)
	}

	if resp, err := client.DescribeApplications(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the configuration options that are used in a particular configuration
// template or environment, or that a specified solution stack defines. The
// description includes the values the options, their default values, and an
// indication of the required action on a running environment if an option value is
// changed.
func elasticbeanstalk_DescribeConfigurationOptions(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.DescribeConfigurationOptionsInput{}

	if len(_elasticbeanstalkApplicationName) > 0 {
		input.ApplicationName = aws.String(_elasticbeanstalkApplicationName)
	}
	if len(_elasticbeanstalkEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_elasticbeanstalkEnvironmentName)
	}
	if len(_elasticbeanstalkOptions) > 0 {
		if err := assignInputField(input, "Options", _elasticbeanstalkOptions); err != nil {
			log.Errorf("invalid --options: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkPlatformArn) > 0 {
		input.PlatformArn = aws.String(_elasticbeanstalkPlatformArn)
	}
	if len(_elasticbeanstalkSolutionStackName) > 0 {
		input.SolutionStackName = aws.String(_elasticbeanstalkSolutionStackName)
	}
	if len(_elasticbeanstalkTemplateName) > 0 {
		input.TemplateName = aws.String(_elasticbeanstalkTemplateName)
	}

	if resp, err := client.DescribeConfigurationOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a description of the settings for the specified configuration set, that
// is, either a configuration template or the configuration set associated with a
// running environment.
//
// When describing the settings for the configuration set associated with a
// running environment, it is possible to receive two sets of setting descriptions.
// One is the deployed configuration set, and the other is a draft configuration of
// an environment that is either in the process of deployment or that failed to
// deploy.
//
// # Related Topics
//
// DeleteEnvironmentConfiguration
func elasticbeanstalk_DescribeConfigurationSettings(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.DescribeConfigurationSettingsInput{
		// ApplicationName: *string, // Required
	}

	if len(_elasticbeanstalkApplicationName) > 0 {
		input.ApplicationName = aws.String(_elasticbeanstalkApplicationName)
	}
	if len(_elasticbeanstalkEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_elasticbeanstalkEnvironmentName)
	}
	if len(_elasticbeanstalkTemplateName) > 0 {
		input.TemplateName = aws.String(_elasticbeanstalkTemplateName)
	}

	if resp, err := client.DescribeConfigurationSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the overall health of the specified environment. The
// DescribeEnvironmentHealth operation is only available with AWS Elastic Beanstalk
// Enhanced Health.
func elasticbeanstalk_DescribeEnvironmentHealth(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.DescribeEnvironmentHealthInput{}

	if len(_elasticbeanstalkAttributeNames) > 0 {
		if err := assignInputField(input, "AttributeNames", _elasticbeanstalkAttributeNames); err != nil {
			log.Errorf("invalid --attribute-names: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_elasticbeanstalkEnvironmentId)
	}
	if len(_elasticbeanstalkEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_elasticbeanstalkEnvironmentName)
	}

	if resp, err := client.DescribeEnvironmentHealth(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists an environment's completed and failed managed actions.
func elasticbeanstalk_DescribeEnvironmentManagedActionHistory(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput{}

	if len(_elasticbeanstalkEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_elasticbeanstalkEnvironmentId)
	}
	if len(_elasticbeanstalkEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_elasticbeanstalkEnvironmentName)
	}
	if len(_elasticbeanstalkMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _elasticbeanstalkMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkNextToken) > 0 {
		input.NextToken = aws.String(_elasticbeanstalkNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeEnvironmentManagedActionHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput
	p := elasticbeanstalk.NewDescribeEnvironmentManagedActionHistoryPaginator(client, input)
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

// Lists an environment's upcoming and in-progress managed actions.
func elasticbeanstalk_DescribeEnvironmentManagedActions(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.DescribeEnvironmentManagedActionsInput{}

	if len(_elasticbeanstalkEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_elasticbeanstalkEnvironmentId)
	}
	if len(_elasticbeanstalkEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_elasticbeanstalkEnvironmentName)
	}
	if len(_elasticbeanstalkStatus) > 0 {
		if err := assignInputField(input, "Status", _elasticbeanstalkStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeEnvironmentManagedActions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns AWS resources for this environment.
func elasticbeanstalk_DescribeEnvironmentResources(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.DescribeEnvironmentResourcesInput{}

	if len(_elasticbeanstalkEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_elasticbeanstalkEnvironmentId)
	}
	if len(_elasticbeanstalkEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_elasticbeanstalkEnvironmentName)
	}

	if resp, err := client.DescribeEnvironmentResources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns descriptions for existing environments.
func elasticbeanstalk_DescribeEnvironments(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.DescribeEnvironmentsInput{}

	if len(_elasticbeanstalkApplicationName) > 0 {
		input.ApplicationName = aws.String(_elasticbeanstalkApplicationName)
	}
	if len(_elasticbeanstalkEnvironmentIds) > 0 {
		input.EnvironmentIds = append([]string(nil), _elasticbeanstalkEnvironmentIds...)
	}
	if len(_elasticbeanstalkEnvironmentNames) > 0 {
		input.EnvironmentNames = append([]string(nil), _elasticbeanstalkEnvironmentNames...)
	}
	if len(_elasticbeanstalkIncludeDeleted) > 0 {
		if err := assignInputField(input, "IncludeDeleted", _elasticbeanstalkIncludeDeleted); err != nil {
			log.Errorf("invalid --include-deleted: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkIncludedDeletedBackTo) > 0 {
		if err := assignInputField(input, "IncludedDeletedBackTo", _elasticbeanstalkIncludedDeletedBackTo); err != nil {
			log.Errorf("invalid --included-deleted-back-to: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _elasticbeanstalkMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkNextToken) > 0 {
		input.NextToken = aws.String(_elasticbeanstalkNextToken)
	}
	if len(_elasticbeanstalkVersionLabel) > 0 {
		input.VersionLabel = aws.String(_elasticbeanstalkVersionLabel)
	}

	if resp, err := client.DescribeEnvironments(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns list of event descriptions matching criteria up to the last 6 weeks.
// This action returns the most recent 1,000 events from the specified NextToken .
func elasticbeanstalk_DescribeEvents(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.DescribeEventsInput{}

	if len(_elasticbeanstalkApplicationName) > 0 {
		input.ApplicationName = aws.String(_elasticbeanstalkApplicationName)
	}
	if len(_elasticbeanstalkEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _elasticbeanstalkEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_elasticbeanstalkEnvironmentId)
	}
	if len(_elasticbeanstalkEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_elasticbeanstalkEnvironmentName)
	}
	if len(_elasticbeanstalkMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _elasticbeanstalkMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkNextToken) > 0 {
		input.NextToken = aws.String(_elasticbeanstalkNextToken)
	}
	if len(_elasticbeanstalkPlatformArn) > 0 {
		input.PlatformArn = aws.String(_elasticbeanstalkPlatformArn)
	}
	if len(_elasticbeanstalkRequestId) > 0 {
		input.RequestId = aws.String(_elasticbeanstalkRequestId)
	}
	if len(_elasticbeanstalkSeverity) > 0 {
		if err := assignInputField(input, "Severity", _elasticbeanstalkSeverity); err != nil {
			log.Errorf("invalid --severity: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _elasticbeanstalkStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkTemplateName) > 0 {
		input.TemplateName = aws.String(_elasticbeanstalkTemplateName)
	}
	if len(_elasticbeanstalkVersionLabel) > 0 {
		input.VersionLabel = aws.String(_elasticbeanstalkVersionLabel)
	}

	if disablePaginator() {
		if resp, err := client.DescribeEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticbeanstalk.DescribeEventsOutput
	p := elasticbeanstalk.NewDescribeEventsPaginator(client, input)
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

// Retrieves detailed information about the health of instances in your AWS
// Elastic Beanstalk. This operation requires [enhanced health reporting].
//
// [enhanced health reporting]: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/health-enhanced.html
func elasticbeanstalk_DescribeInstancesHealth(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.DescribeInstancesHealthInput{}

	if len(_elasticbeanstalkAttributeNames) > 0 {
		if err := assignInputField(input, "AttributeNames", _elasticbeanstalkAttributeNames); err != nil {
			log.Errorf("invalid --attribute-names: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_elasticbeanstalkEnvironmentId)
	}
	if len(_elasticbeanstalkEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_elasticbeanstalkEnvironmentName)
	}
	if len(_elasticbeanstalkNextToken) > 0 {
		input.NextToken = aws.String(_elasticbeanstalkNextToken)
	}

	if resp, err := client.DescribeInstancesHealth(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a platform version. Provides full details. Compare to ListPlatformVersions, which
// provides summary information about a list of platform versions.
//
// For definitions of platform version and other platform-related terms, see [AWS Elastic Beanstalk Platforms Glossary].
//
// [AWS Elastic Beanstalk Platforms Glossary]: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/platforms-glossary.html
func elasticbeanstalk_DescribePlatformVersion(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.DescribePlatformVersionInput{}

	if len(_elasticbeanstalkPlatformArn) > 0 {
		input.PlatformArn = aws.String(_elasticbeanstalkPlatformArn)
	}

	if resp, err := client.DescribePlatformVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociate the operations role from an environment. After this call is made,
// Elastic Beanstalk uses the caller's permissions for permissions to downstream
// services during subsequent calls acting on this environment. For more
// information, see [Operations roles]in the AWS Elastic Beanstalk Developer Guide.
//
// [Operations roles]: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/iam-operationsrole.html
func elasticbeanstalk_DisassociateEnvironmentOperationsRole(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.DisassociateEnvironmentOperationsRoleInput{
		// EnvironmentName: *string, // Required
	}

	if len(_elasticbeanstalkEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_elasticbeanstalkEnvironmentName)
	}

	if resp, err := client.DisassociateEnvironmentOperationsRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of the available solution stack names, with the public version
// first and then in reverse chronological order.
func elasticbeanstalk_ListAvailableSolutionStacks(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.ListAvailableSolutionStacksInput{}

	if resp, err := client.ListAvailableSolutionStacks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the platform branches available for your account in an AWS Region.
// Provides summary information about each platform branch.
//
// For definitions of platform branch and other platform-related terms, see [AWS Elastic Beanstalk Platforms Glossary].
//
// [AWS Elastic Beanstalk Platforms Glossary]: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/platforms-glossary.html
func elasticbeanstalk_ListPlatformBranches(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.ListPlatformBranchesInput{}

	if len(_elasticbeanstalkFilters) > 0 {
		if err := assignInputField(input, "Filters", _elasticbeanstalkFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _elasticbeanstalkMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkNextToken) > 0 {
		input.NextToken = aws.String(_elasticbeanstalkNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPlatformBranches(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticbeanstalk.ListPlatformBranchesOutput
	p := elasticbeanstalk.NewListPlatformBranchesPaginator(client, input)
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

// Lists the platform versions available for your account in an AWS Region.
// Provides summary information about each platform version. Compare to DescribePlatformVersion, which
// provides full details about a single platform version.
//
// For definitions of platform version and other platform-related terms, see [AWS Elastic Beanstalk Platforms Glossary].
//
// [AWS Elastic Beanstalk Platforms Glossary]: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/platforms-glossary.html
func elasticbeanstalk_ListPlatformVersions(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.ListPlatformVersionsInput{}

	if len(_elasticbeanstalkFilters) > 0 {
		if err := assignInputField(input, "Filters", _elasticbeanstalkFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _elasticbeanstalkMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkNextToken) > 0 {
		input.NextToken = aws.String(_elasticbeanstalkNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPlatformVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticbeanstalk.ListPlatformVersionsOutput
	p := elasticbeanstalk.NewListPlatformVersionsPaginator(client, input)
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

// Return the tags applied to an AWS Elastic Beanstalk resource. The response
// contains a list of tag key-value pairs.
//
// Elastic Beanstalk supports tagging of all of its resources. For details about
// resource tagging, see [Tagging Application Resources].
//
// [Tagging Application Resources]: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/applications-tagging-resources.html
func elasticbeanstalk_ListTagsForResource(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_elasticbeanstalkResourceArn) > 0 {
		input.ResourceArn = aws.String(_elasticbeanstalkResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes and recreates all of the AWS resources (for example: the Auto Scaling
// group, load balancer, etc.) for a specified environment and forces a restart.
func elasticbeanstalk_RebuildEnvironment(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.RebuildEnvironmentInput{}

	if len(_elasticbeanstalkEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_elasticbeanstalkEnvironmentId)
	}
	if len(_elasticbeanstalkEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_elasticbeanstalkEnvironmentName)
	}

	if resp, err := client.RebuildEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a request to compile the specified type of information of the
// deployed environment.
//
// Setting the InfoType to tail compiles the last lines from the application
// server log files of every Amazon EC2 instance in your environment.
//
// Setting the InfoType to bundle compresses the application server log files for
// every Amazon EC2 instance into a .zip file. Legacy and .NET containers do not
// support bundle logs.
//
// Use RetrieveEnvironmentInfo to obtain the set of logs.
//
// # Related Topics
//
// RetrieveEnvironmentInfo
func elasticbeanstalk_RequestEnvironmentInfo(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.RequestEnvironmentInfoInput{
		// InfoType: types.EnvironmentInfoType, // Required
	}

	if len(_elasticbeanstalkInfoType) > 0 {
		if err := assignInputField(input, "InfoType", _elasticbeanstalkInfoType); err != nil {
			log.Errorf("invalid --info-type: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_elasticbeanstalkEnvironmentId)
	}
	if len(_elasticbeanstalkEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_elasticbeanstalkEnvironmentName)
	}

	if resp, err := client.RequestEnvironmentInfo(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Causes the environment to restart the application container server running on
// each Amazon EC2 instance.
func elasticbeanstalk_RestartAppServer(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.RestartAppServerInput{}

	if len(_elasticbeanstalkEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_elasticbeanstalkEnvironmentId)
	}
	if len(_elasticbeanstalkEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_elasticbeanstalkEnvironmentName)
	}

	if resp, err := client.RestartAppServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the compiled information from a RequestEnvironmentInfo request.
// # Related Topics
//
// RequestEnvironmentInfo
func elasticbeanstalk_RetrieveEnvironmentInfo(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.RetrieveEnvironmentInfoInput{
		// InfoType: types.EnvironmentInfoType, // Required
	}

	if len(_elasticbeanstalkInfoType) > 0 {
		if err := assignInputField(input, "InfoType", _elasticbeanstalkInfoType); err != nil {
			log.Errorf("invalid --info-type: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_elasticbeanstalkEnvironmentId)
	}
	if len(_elasticbeanstalkEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_elasticbeanstalkEnvironmentName)
	}

	if resp, err := client.RetrieveEnvironmentInfo(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Swaps the CNAMEs of two environments.
func elasticbeanstalk_SwapEnvironmentCNAMEs(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.SwapEnvironmentCNAMEsInput{}

	if len(_elasticbeanstalkDestinationEnvironmentId) > 0 {
		input.DestinationEnvironmentId = aws.String(_elasticbeanstalkDestinationEnvironmentId)
	}
	if len(_elasticbeanstalkDestinationEnvironmentName) > 0 {
		input.DestinationEnvironmentName = aws.String(_elasticbeanstalkDestinationEnvironmentName)
	}
	if len(_elasticbeanstalkSourceEnvironmentId) > 0 {
		input.SourceEnvironmentId = aws.String(_elasticbeanstalkSourceEnvironmentId)
	}
	if len(_elasticbeanstalkSourceEnvironmentName) > 0 {
		input.SourceEnvironmentName = aws.String(_elasticbeanstalkSourceEnvironmentName)
	}

	if resp, err := client.SwapEnvironmentCNAMEs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Terminates the specified environment.
func elasticbeanstalk_TerminateEnvironment(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.TerminateEnvironmentInput{}

	if len(_elasticbeanstalkEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_elasticbeanstalkEnvironmentId)
	}
	if len(_elasticbeanstalkEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_elasticbeanstalkEnvironmentName)
	}
	if len(_elasticbeanstalkForceTerminate) > 0 {
		if err := assignInputField(input, "ForceTerminate", _elasticbeanstalkForceTerminate); err != nil {
			log.Errorf("invalid --force-terminate: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkTerminateResources) > 0 {
		if err := assignInputField(input, "TerminateResources", _elasticbeanstalkTerminateResources); err != nil {
			log.Errorf("invalid --terminate-resources: %s", err.Error())
			return
		}
	}

	if resp, err := client.TerminateEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified application to have the specified properties.
// If a property (for example, description ) is not provided, the value remains
// unchanged. To clear these properties, specify an empty string.
func elasticbeanstalk_UpdateApplication(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.UpdateApplicationInput{
		// ApplicationName: *string, // Required
	}

	if len(_elasticbeanstalkApplicationName) > 0 {
		input.ApplicationName = aws.String(_elasticbeanstalkApplicationName)
	}
	if len(_elasticbeanstalkDescription) > 0 {
		input.Description = aws.String(_elasticbeanstalkDescription)
	}

	if resp, err := client.UpdateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies lifecycle settings for an application.
func elasticbeanstalk_UpdateApplicationResourceLifecycle(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.UpdateApplicationResourceLifecycleInput{
		// ApplicationName: *string, // Required
		// ResourceLifecycleConfig: *types.ApplicationResourceLifecycleConfig, // Required
	}

	if len(_elasticbeanstalkApplicationName) > 0 {
		input.ApplicationName = aws.String(_elasticbeanstalkApplicationName)
	}
	if len(_elasticbeanstalkResourceLifecycleConfig) > 0 {
		if err := assignInputField(input, "ResourceLifecycleConfig", _elasticbeanstalkResourceLifecycleConfig); err != nil {
			log.Errorf("invalid --resource-lifecycle-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateApplicationResourceLifecycle(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified application version to have the specified properties.
// If a property (for example, description ) is not provided, the value remains
// unchanged. To clear properties, specify an empty string.
func elasticbeanstalk_UpdateApplicationVersion(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.UpdateApplicationVersionInput{
		// ApplicationName: *string, // Required
		// VersionLabel: *string, // Required
	}

	if len(_elasticbeanstalkApplicationName) > 0 {
		input.ApplicationName = aws.String(_elasticbeanstalkApplicationName)
	}
	if len(_elasticbeanstalkVersionLabel) > 0 {
		input.VersionLabel = aws.String(_elasticbeanstalkVersionLabel)
	}
	if len(_elasticbeanstalkDescription) > 0 {
		input.Description = aws.String(_elasticbeanstalkDescription)
	}

	if resp, err := client.UpdateApplicationVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified configuration template to have the specified properties
// or configuration option values.
//
// If a property (for example, ApplicationName ) is not provided, its value remains
// unchanged. To clear such properties, specify an empty string.
//
// # Related Topics
//
// DescribeConfigurationOptions
func elasticbeanstalk_UpdateConfigurationTemplate(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.UpdateConfigurationTemplateInput{
		// ApplicationName: *string, // Required
		// TemplateName: *string, // Required
	}

	if len(_elasticbeanstalkApplicationName) > 0 {
		input.ApplicationName = aws.String(_elasticbeanstalkApplicationName)
	}
	if len(_elasticbeanstalkTemplateName) > 0 {
		input.TemplateName = aws.String(_elasticbeanstalkTemplateName)
	}
	if len(_elasticbeanstalkDescription) > 0 {
		input.Description = aws.String(_elasticbeanstalkDescription)
	}
	if len(_elasticbeanstalkOptionSettings) > 0 {
		if err := assignInputField(input, "OptionSettings", _elasticbeanstalkOptionSettings); err != nil {
			log.Errorf("invalid --option-settings: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkOptionsToRemove) > 0 {
		if err := assignInputField(input, "OptionsToRemove", _elasticbeanstalkOptionsToRemove); err != nil {
			log.Errorf("invalid --options-to-remove: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateConfigurationTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the environment description, deploys a new application version, updates
// the configuration settings to an entirely new configuration template, or updates
// select configuration option values in the running environment.
//
// Attempting to update both the release and configuration is not allowed and AWS
// Elastic Beanstalk returns an InvalidParameterCombination error.
//
// When updating the configuration settings to a new template or individual
// settings, a draft configuration is created and DescribeConfigurationSettingsfor this environment returns two
// setting descriptions with different DeploymentStatus values.
func elasticbeanstalk_UpdateEnvironment(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.UpdateEnvironmentInput{}

	if len(_elasticbeanstalkApplicationName) > 0 {
		input.ApplicationName = aws.String(_elasticbeanstalkApplicationName)
	}
	if len(_elasticbeanstalkDescription) > 0 {
		input.Description = aws.String(_elasticbeanstalkDescription)
	}
	if len(_elasticbeanstalkEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_elasticbeanstalkEnvironmentId)
	}
	if len(_elasticbeanstalkEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_elasticbeanstalkEnvironmentName)
	}
	if len(_elasticbeanstalkGroupName) > 0 {
		input.GroupName = aws.String(_elasticbeanstalkGroupName)
	}
	if len(_elasticbeanstalkOptionSettings) > 0 {
		if err := assignInputField(input, "OptionSettings", _elasticbeanstalkOptionSettings); err != nil {
			log.Errorf("invalid --option-settings: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkOptionsToRemove) > 0 {
		if err := assignInputField(input, "OptionsToRemove", _elasticbeanstalkOptionsToRemove); err != nil {
			log.Errorf("invalid --options-to-remove: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkPlatformArn) > 0 {
		input.PlatformArn = aws.String(_elasticbeanstalkPlatformArn)
	}
	if len(_elasticbeanstalkSolutionStackName) > 0 {
		input.SolutionStackName = aws.String(_elasticbeanstalkSolutionStackName)
	}
	if len(_elasticbeanstalkTemplateName) > 0 {
		input.TemplateName = aws.String(_elasticbeanstalkTemplateName)
	}
	if len(_elasticbeanstalkTier) > 0 {
		if err := assignInputField(input, "Tier", _elasticbeanstalkTier); err != nil {
			log.Errorf("invalid --tier: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkVersionLabel) > 0 {
		input.VersionLabel = aws.String(_elasticbeanstalkVersionLabel)
	}

	if resp, err := client.UpdateEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the list of tags applied to an AWS Elastic Beanstalk resource. Two lists
// can be passed: TagsToAdd for tags to add or update, and TagsToRemove .
//
// Elastic Beanstalk supports tagging of all of its resources. For details about
// resource tagging, see [Tagging Application Resources].
//
// If you create a custom IAM user policy to control permission to this operation,
// specify one of the following two virtual actions (or both) instead of the API
// operation name:
//
// elasticbeanstalk:AddTags Controls permission to call UpdateTagsForResource and
// pass a list of tags to add in the TagsToAdd parameter.
//
// elasticbeanstalk:RemoveTags Controls permission to call UpdateTagsForResource
// and pass a list of tag keys to remove in the TagsToRemove parameter.
//
// For details about creating a custom user policy, see [Creating a Custom User Policy].
//
// [Creating a Custom User Policy]: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/AWSHowTo.iam.managed-policies.html#AWSHowTo.iam.policies
// [Tagging Application Resources]: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/applications-tagging-resources.html
func elasticbeanstalk_UpdateTagsForResource(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.UpdateTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_elasticbeanstalkResourceArn) > 0 {
		input.ResourceArn = aws.String(_elasticbeanstalkResourceArn)
	}
	if len(_elasticbeanstalkTagsToAdd) > 0 {
		if err := assignInputField(input, "TagsToAdd", _elasticbeanstalkTagsToAdd); err != nil {
			log.Errorf("invalid --tags-to-add: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkTagsToRemove) > 0 {
		input.TagsToRemove = append([]string(nil), _elasticbeanstalkTagsToRemove...)
	}

	if resp, err := client.UpdateTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Takes a set of configuration settings and either a configuration template or
// environment, and determines whether those values are valid.
//
// This action returns a list of messages indicating any errors or warnings
// associated with the selection of option values.
func elasticbeanstalk_ValidateConfigurationSettings(cfg aws.Config, client *elasticbeanstalk.Client) {
	input := &elasticbeanstalk.ValidateConfigurationSettingsInput{
		// ApplicationName: *string, // Required
		// OptionSettings: []types.ConfigurationOptionSetting, // Required
	}

	if len(_elasticbeanstalkApplicationName) > 0 {
		input.ApplicationName = aws.String(_elasticbeanstalkApplicationName)
	}
	if len(_elasticbeanstalkOptionSettings) > 0 {
		if err := assignInputField(input, "OptionSettings", _elasticbeanstalkOptionSettings); err != nil {
			log.Errorf("invalid --option-settings: %s", err.Error())
			return
		}
	}
	if len(_elasticbeanstalkEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_elasticbeanstalkEnvironmentName)
	}
	if len(_elasticbeanstalkTemplateName) > 0 {
		input.TemplateName = aws.String(_elasticbeanstalkTemplateName)
	}

	if resp, err := client.ValidateConfigurationSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_elasticbeanstalkCmd)
	_elasticbeanstalkCmd.Flags().SortFlags = false

	_elasticbeanstalkCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_elasticbeanstalkCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_elasticbeanstalkCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkActionId, "action-id", "", "", "Action ID")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkApplicationName, "application-name", "", "", "Application Name")
	_elasticbeanstalkCmd.Flags().StringSliceVarP(&_elasticbeanstalkApplicationNames, "application-names", "", nil, "Application Names")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkAttributeNames, "attribute-names", "", "", "Attribute Names")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkAutoCreateApplication, "auto-create-application", "", "", "Auto Create Application")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkBuildConfiguration, "build-configuration", "", "", "Build Configuration")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkCNAMEPrefix, "cname-prefix", "", "", "Cname Prefix")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkDeleteSourceBundle, "delete-source-bundle", "", "", "Delete Source Bundle")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkDescription, "description", "", "", "Description")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkDestinationEnvironmentId, "destination-environment-id", "", "", "Destination Environment ID")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkDestinationEnvironmentName, "destination-environment-name", "", "", "Destination Environment Name")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkEndTime, "end-time", "", "", "End Time")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkEnvironmentId, "environment-id", "", "", "Environment ID")
	_elasticbeanstalkCmd.Flags().StringSliceVarP(&_elasticbeanstalkEnvironmentIds, "environment-ids", "", nil, "Environment Ids")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkEnvironmentName, "environment-name", "", "", "Environment Name")
	_elasticbeanstalkCmd.Flags().StringSliceVarP(&_elasticbeanstalkEnvironmentNames, "environment-names", "", nil, "Environment Names")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkFilters, "filters", "", "", "Filters")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkForceTerminate, "force-terminate", "", "", "Force Terminate")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkGroupName, "group-name", "", "", "Group Name")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkIncludeDeleted, "include-deleted", "", "", "Include Deleted")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkIncludedDeletedBackTo, "included-deleted-back-to", "", "", "Included Deleted Back To")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkInfoType, "info-type", "", "", "Info Type")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkMaxItems, "max-items", "", "", "Max Items")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkMaxRecords, "max-records", "", "", "Max Records")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkNextToken, "next-token", "", "", "Next Token")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkOperationsRole, "operations-role", "", "", "Operations Role")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkOptionSettings, "option-settings", "", "", "Option Settings")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkOptions, "options", "", "", "Options")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkOptionsToRemove, "options-to-remove", "", "", "Options To Remove")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkPlatformArn, "platform-arn", "", "", "Platform ARN")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkPlatformDefinitionBundle, "platform-definition-bundle", "", "", "Platform Definition Bundle")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkPlatformName, "platform-name", "", "", "Platform Name")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkPlatformVersion, "platform-version", "", "", "Platform Version")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkProcess, "process", "", "", "Process")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkRequestId, "request-id", "", "", "Request ID")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkResourceArn, "resource-arn", "", "", "Resource ARN")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkResourceLifecycleConfig, "resource-lifecycle-config", "", "", "Resource Lifecycle Config")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkSeverity, "severity", "", "", "Severity")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkSolutionStackName, "solution-stack-name", "", "", "Solution Stack Name")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkSourceBuildInformation, "source-build-information", "", "", "Source Build Information")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkSourceBundle, "source-bundle", "", "", "Source Bundle")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkSourceConfiguration, "source-configuration", "", "", "Source Configuration")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkSourceEnvironmentId, "source-environment-id", "", "", "Source Environment ID")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkSourceEnvironmentName, "source-environment-name", "", "", "Source Environment Name")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkStartTime, "start-time", "", "", "Start Time")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkStatus, "status", "", "", "Status")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkTags, "tags", "", "", "Tags")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkTagsToAdd, "tags-to-add", "", "", "Tags To Add")
	_elasticbeanstalkCmd.Flags().StringSliceVarP(&_elasticbeanstalkTagsToRemove, "tags-to-remove", "", nil, "Tags To Remove")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkTemplateName, "template-name", "", "", "Template Name")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkTerminateEnvByForce, "terminate-env-by-force", "", "", "Terminate Env By Force")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkTerminateResources, "terminate-resources", "", "", "Terminate Resources")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkTier, "tier", "", "", "Tier")
	_elasticbeanstalkCmd.Flags().StringVarP(&_elasticbeanstalkVersionLabel, "version-label", "", "", "Version Label")
	_elasticbeanstalkCmd.Flags().StringSliceVarP(&_elasticbeanstalkVersionLabels, "version-labels", "", nil, "Version Labels")

	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkAbortEnvironmentUpdate, "abort-environment-update", "", false, "Abort Environment Update")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkApplyEnvironmentManagedAction, "apply-environment-managed-action", "", false, "Apply Environment Managed Action")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkAssociateEnvironmentOperationsRole, "associate-environment-operations-role", "", false, "Associate Environment Operations Role")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkCheckDNSAvailability, "check-dns-availability", "", false, "Check DNS Availability")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkComposeEnvironments, "compose-environments", "", false, "Compose Environments")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkCreateApplication, "create-application", "", false, "Create Application")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkCreateApplicationVersion, "create-application-version", "", false, "Create Application Version")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkCreateConfigurationTemplate, "create-configuration-template", "", false, "Create Configuration Template")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkCreateEnvironment, "create-environment", "", false, "Create Environment")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkCreatePlatformVersion, "create-platform-version", "", false, "Create Platform Version")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkCreateStorageLocation, "create-storage-location", "", false, "Create Storage Location")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkDeleteApplication, "delete-application", "", false, "Delete Application")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkDeleteApplicationVersion, "delete-application-version", "", false, "Delete Application Version")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkDeleteConfigurationTemplate, "delete-configuration-template", "", false, "Delete Configuration Template")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkDeleteEnvironmentConfiguration, "delete-environment-configuration", "", false, "Delete Environment Configuration")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkDeletePlatformVersion, "delete-platform-version", "", false, "Delete Platform Version")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkDescribeAccountAttributes, "describe-account-attributes", "", false, "Describe Account Attributes")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkDescribeApplicationVersions, "describe-application-versions", "", false, "Describe Application Versions")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkDescribeApplications, "describe-applications", "", false, "Describe Applications")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkDescribeConfigurationOptions, "describe-configuration-options", "", false, "Describe Configuration Options")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkDescribeConfigurationSettings, "describe-configuration-settings", "", false, "Describe Configuration Settings")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkDescribeEnvironmentHealth, "describe-environment-health", "", false, "Describe Environment Health")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkDescribeEnvironmentManagedActionHistory, "describe-environment-managed-action-history", "", false, "Describe Environment Managed Action History")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkDescribeEnvironmentManagedActions, "describe-environment-managed-actions", "", false, "Describe Environment Managed Actions")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkDescribeEnvironmentResources, "describe-environment-resources", "", false, "Describe Environment Resources")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkDescribeEnvironments, "describe-environments", "", false, "Describe Environments")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkDescribeEvents, "describe-events", "", false, "Describe Events")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkDescribeInstancesHealth, "describe-instances-health", "", false, "Describe Instances Health")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkDescribePlatformVersion, "describe-platform-version", "", false, "Describe Platform Version")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkDisassociateEnvironmentOperationsRole, "disassociate-environment-operations-role", "", false, "Disassociate Environment Operations Role")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkListAvailableSolutionStacks, "list-available-solution-stacks", "", false, "List Available Solution Stacks")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkListPlatformBranches, "list-platform-branches", "", false, "List Platform Branches")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkListPlatformVersions, "list-platform-versions", "", false, "List Platform Versions")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkRebuildEnvironment, "rebuild-environment", "", false, "Rebuild Environment")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkRequestEnvironmentInfo, "request-environment-info", "", false, "Request Environment Info")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkRestartAppServer, "restart-app-server", "", false, "Restart App Server")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkRetrieveEnvironmentInfo, "retrieve-environment-info", "", false, "Retrieve Environment Info")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkSwapEnvironmentCNAMEs, "swap-environment-cnames", "", false, "Swap Environment Cnames")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkTerminateEnvironment, "terminate-environment", "", false, "Terminate Environment")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkUpdateApplication, "update-application", "", false, "Update Application")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkUpdateApplicationResourceLifecycle, "update-application-resource-lifecycle", "", false, "Update Application Resource Lifecycle")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkUpdateApplicationVersion, "update-application-version", "", false, "Update Application Version")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkUpdateConfigurationTemplate, "update-configuration-template", "", false, "Update Configuration Template")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkUpdateEnvironment, "update-environment", "", false, "Update Environment")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkUpdateTagsForResource, "update-tags-for-resource", "", false, "Update Tags For Resource")
	_elasticbeanstalkCmd.Flags().BoolVarP(&_elasticbeanstalkValidateConfigurationSettings, "validate-configuration-settings", "", false, "Validate Configuration Settings")

}
