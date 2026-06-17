package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appconfig"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// appconfigCmd represents the appconfig command
var _appconfigCmd = &cobra.Command{
	Use:   "appconfig",
	Short: "AWS appconfig CLI",
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
		client := appconfig.NewFromConfig(cfg)
		if _appconfigCreateApplication {
			appconfig_CreateApplication(cfg, client)
			return
		}
		if _appconfigCreateConfigurationProfile {
			appconfig_CreateConfigurationProfile(cfg, client)
			return
		}
		if _appconfigCreateDeploymentStrategy {
			appconfig_CreateDeploymentStrategy(cfg, client)
			return
		}
		if _appconfigCreateEnvironment {
			appconfig_CreateEnvironment(cfg, client)
			return
		}
		if _appconfigCreateExtension {
			appconfig_CreateExtension(cfg, client)
			return
		}
		if _appconfigCreateExtensionAssociation {
			appconfig_CreateExtensionAssociation(cfg, client)
			return
		}
		if _appconfigCreateHostedConfigurationVersion {
			appconfig_CreateHostedConfigurationVersion(cfg, client)
			return
		}
		if _appconfigDeleteApplication {
			appconfig_DeleteApplication(cfg, client)
			return
		}
		if _appconfigDeleteConfigurationProfile {
			appconfig_DeleteConfigurationProfile(cfg, client)
			return
		}
		if _appconfigDeleteDeploymentStrategy {
			appconfig_DeleteDeploymentStrategy(cfg, client)
			return
		}
		if _appconfigDeleteEnvironment {
			appconfig_DeleteEnvironment(cfg, client)
			return
		}
		if _appconfigDeleteExtension {
			appconfig_DeleteExtension(cfg, client)
			return
		}
		if _appconfigDeleteExtensionAssociation {
			appconfig_DeleteExtensionAssociation(cfg, client)
			return
		}
		if _appconfigDeleteHostedConfigurationVersion {
			appconfig_DeleteHostedConfigurationVersion(cfg, client)
			return
		}
		if _appconfigGetAccountSettings {
			appconfig_GetAccountSettings(cfg, client)
			return
		}
		if _appconfigGetApplication {
			appconfig_GetApplication(cfg, client)
			return
		}
		if _appconfigGetConfiguration {
			appconfig_GetConfiguration(cfg, client)
			return
		}
		if _appconfigGetConfigurationProfile {
			appconfig_GetConfigurationProfile(cfg, client)
			return
		}
		if _appconfigGetDeployment {
			appconfig_GetDeployment(cfg, client)
			return
		}
		if _appconfigGetDeploymentStrategy {
			appconfig_GetDeploymentStrategy(cfg, client)
			return
		}
		if _appconfigGetEnvironment {
			appconfig_GetEnvironment(cfg, client)
			return
		}
		if _appconfigGetExtension {
			appconfig_GetExtension(cfg, client)
			return
		}
		if _appconfigGetExtensionAssociation {
			appconfig_GetExtensionAssociation(cfg, client)
			return
		}
		if _appconfigGetHostedConfigurationVersion {
			appconfig_GetHostedConfigurationVersion(cfg, client)
			return
		}
		if _appconfigListApplications {
			appconfig_ListApplications(cfg, client)
			return
		}
		if _appconfigListConfigurationProfiles {
			appconfig_ListConfigurationProfiles(cfg, client)
			return
		}
		if _appconfigListDeploymentStrategies {
			appconfig_ListDeploymentStrategies(cfg, client)
			return
		}
		if _appconfigListDeployments {
			appconfig_ListDeployments(cfg, client)
			return
		}
		if _appconfigListEnvironments {
			appconfig_ListEnvironments(cfg, client)
			return
		}
		if _appconfigListExtensionAssociations {
			appconfig_ListExtensionAssociations(cfg, client)
			return
		}
		if _appconfigListExtensions {
			appconfig_ListExtensions(cfg, client)
			return
		}
		if _appconfigListHostedConfigurationVersions {
			appconfig_ListHostedConfigurationVersions(cfg, client)
			return
		}
		if _appconfigListTagsForResource {
			appconfig_ListTagsForResource(cfg, client)
			return
		}
		if _appconfigStartDeployment {
			appconfig_StartDeployment(cfg, client)
			return
		}
		if _appconfigStopDeployment {
			appconfig_StopDeployment(cfg, client)
			return
		}
		if _appconfigTagResource {
			appconfig_TagResource(cfg, client)
			return
		}
		if _appconfigUntagResource {
			appconfig_UntagResource(cfg, client)
			return
		}
		if _appconfigUpdateAccountSettings {
			appconfig_UpdateAccountSettings(cfg, client)
			return
		}
		if _appconfigUpdateApplication {
			appconfig_UpdateApplication(cfg, client)
			return
		}
		if _appconfigUpdateConfigurationProfile {
			appconfig_UpdateConfigurationProfile(cfg, client)
			return
		}
		if _appconfigUpdateDeploymentStrategy {
			appconfig_UpdateDeploymentStrategy(cfg, client)
			return
		}
		if _appconfigUpdateEnvironment {
			appconfig_UpdateEnvironment(cfg, client)
			return
		}
		if _appconfigUpdateExtension {
			appconfig_UpdateExtension(cfg, client)
			return
		}
		if _appconfigUpdateExtensionAssociation {
			appconfig_UpdateExtensionAssociation(cfg, client)
			return
		}
		if _appconfigValidateConfiguration {
			appconfig_ValidateConfiguration(cfg, client)
			return
		}

	},
}

var (
	_appconfigCreateApplication                bool
	_appconfigCreateConfigurationProfile       bool
	_appconfigCreateDeploymentStrategy         bool
	_appconfigCreateEnvironment                bool
	_appconfigCreateExtension                  bool
	_appconfigCreateExtensionAssociation       bool
	_appconfigCreateHostedConfigurationVersion bool
	_appconfigDeleteApplication                bool
	_appconfigDeleteConfigurationProfile       bool
	_appconfigDeleteDeploymentStrategy         bool
	_appconfigDeleteEnvironment                bool
	_appconfigDeleteExtension                  bool
	_appconfigDeleteExtensionAssociation       bool
	_appconfigDeleteHostedConfigurationVersion bool
	_appconfigGetAccountSettings               bool
	_appconfigGetApplication                   bool
	_appconfigGetConfiguration                 bool
	_appconfigGetConfigurationProfile          bool
	_appconfigGetDeployment                    bool
	_appconfigGetDeploymentStrategy            bool
	_appconfigGetEnvironment                   bool
	_appconfigGetExtension                     bool
	_appconfigGetExtensionAssociation          bool
	_appconfigGetHostedConfigurationVersion    bool
	_appconfigListApplications                 bool
	_appconfigListConfigurationProfiles        bool
	_appconfigListDeploymentStrategies         bool
	_appconfigListDeployments                  bool
	_appconfigListEnvironments                 bool
	_appconfigListExtensionAssociations        bool
	_appconfigListExtensions                   bool
	_appconfigListHostedConfigurationVersions  bool
	_appconfigListTagsForResource              bool
	_appconfigStartDeployment                  bool
	_appconfigStopDeployment                   bool
	_appconfigTagResource                      bool
	_appconfigUntagResource                    bool
	_appconfigUpdateAccountSettings            bool
	_appconfigUpdateApplication                bool
	_appconfigUpdateConfigurationProfile       bool
	_appconfigUpdateDeploymentStrategy         bool
	_appconfigUpdateEnvironment                bool
	_appconfigUpdateExtension                  bool
	_appconfigUpdateExtensionAssociation       bool
	_appconfigValidateConfiguration            bool

	_appconfigActions                     string
	_appconfigAllowRevert                 string
	_appconfigApplication                 string
	_appconfigApplicationId               string
	_appconfigClientConfigurationVersion  string
	_appconfigClientId                    string
	_appconfigConfiguration               string
	_appconfigConfigurationProfileId      string
	_appconfigConfigurationVersion        string
	_appconfigContent                     string
	_appconfigContentType                 string
	_appconfigDeletionProtection          string
	_appconfigDeletionProtectionCheck     string
	_appconfigDeploymentDurationInMinutes string
	_appconfigDeploymentNumber            string
	_appconfigDeploymentStrategyId        string
	_appconfigDescription                 string
	_appconfigDynamicExtensionParameters  string
	_appconfigEnvironment                 string
	_appconfigEnvironmentId               string
	_appconfigExtensionAssociationId      string
	_appconfigExtensionIdentifier         string
	_appconfigExtensionVersionNumber      string
	_appconfigFinalBakeTimeInMinutes      string
	_appconfigGrowthFactor                string
	_appconfigGrowthType                  string
	_appconfigKmsKeyIdentifier            string
	_appconfigLatestVersionNumber         string
	_appconfigLocationUri                 string
	_appconfigMaxResults                  string
	_appconfigMonitors                    string
	_appconfigName                        string
	_appconfigNextToken                   string
	_appconfigParameters                  string
	_appconfigReplicateTo                 string
	_appconfigResourceArn                 string
	_appconfigResourceIdentifier          string
	_appconfigRetrievalRoleArn            string
	_appconfigTagKeys                     []string
	_appconfigTags                        string
	_appconfigType                        string
	_appconfigValidators                  string
	_appconfigVersionLabel                string
	_appconfigVersionNumber               string
)

// Creates an application. In AppConfig, an application is simply an
// organizational construct like a folder. This organizational construct has a
// relationship with some unit of executable code. For example, you could create an
// application called MyMobileApp to organize and manage configuration data for a
// mobile application installed by your users.
func appconfig_CreateApplication(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.CreateApplicationInput{
		// Name: *string, // Required
	}

	if len(_appconfigName) > 0 {
		input.Name = aws.String(_appconfigName)
	}
	if len(_appconfigDescription) > 0 {
		input.Description = aws.String(_appconfigDescription)
	}
	if len(_appconfigTags) > 0 {
		if err := assignInputField(input, "Tags", _appconfigTags); err != nil {
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

// Creates a configuration profile, which is information that enables AppConfig to
// access the configuration source. Valid configuration sources include the
// following:
//
// - Configuration data in YAML, JSON, and other formats stored in the AppConfig
// hosted configuration store
//
// - Configuration data stored as objects in an Amazon Simple Storage Service
// (Amazon S3) bucket
//
// - Pipelines stored in CodePipeline
//
// - Secrets stored in Secrets Manager
//
// - Standard and secure string parameters stored in Amazon Web Services Systems
// Manager Parameter Store
//
// - Configuration data in SSM documents stored in the Systems Manager document
// store
//
// A configuration profile includes the following information:
//
// - The URI location of the configuration data.
//
// - The Identity and Access Management (IAM) role that provides access to the
// configuration data.
//
// - A validator for the configuration data. Available validators include either
// a JSON Schema or an Amazon Web Services Lambda function.
//
// For more information, see [Create a Configuration and a Configuration Profile] in the AppConfig User Guide.
//
// [Create a Configuration and a Configuration Profile]: http://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-configuration-and-profile.html
func appconfig_CreateConfigurationProfile(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.CreateConfigurationProfileInput{
		// ApplicationId: *string, // Required
		// LocationUri: *string, // Required
		// Name: *string, // Required
	}

	if len(_appconfigApplicationId) > 0 {
		input.ApplicationId = aws.String(_appconfigApplicationId)
	}
	if len(_appconfigLocationUri) > 0 {
		input.LocationUri = aws.String(_appconfigLocationUri)
	}
	if len(_appconfigName) > 0 {
		input.Name = aws.String(_appconfigName)
	}
	if len(_appconfigDescription) > 0 {
		input.Description = aws.String(_appconfigDescription)
	}
	if len(_appconfigKmsKeyIdentifier) > 0 {
		input.KmsKeyIdentifier = aws.String(_appconfigKmsKeyIdentifier)
	}
	if len(_appconfigRetrievalRoleArn) > 0 {
		input.RetrievalRoleArn = aws.String(_appconfigRetrievalRoleArn)
	}
	if len(_appconfigTags) > 0 {
		if err := assignInputField(input, "Tags", _appconfigTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_appconfigType) > 0 {
		input.Type = aws.String(_appconfigType)
	}
	if len(_appconfigValidators) > 0 {
		if err := assignInputField(input, "Validators", _appconfigValidators); err != nil {
			log.Errorf("invalid --validators: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConfigurationProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a deployment strategy that defines important criteria for rolling out
// your configuration to the designated targets. A deployment strategy includes the
// overall duration required, a percentage of targets to receive the deployment
// during each interval, an algorithm that defines how percentage grows, and bake
// time.
func appconfig_CreateDeploymentStrategy(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.CreateDeploymentStrategyInput{
		// DeploymentDurationInMinutes: *int32, // Required
		// GrowthFactor: *float32, // Required
		// Name: *string, // Required
	}

	if len(_appconfigDeploymentDurationInMinutes) > 0 {
		if err := assignInputField(input, "DeploymentDurationInMinutes", _appconfigDeploymentDurationInMinutes); err != nil {
			log.Errorf("invalid --deployment-duration-in-minutes: %s", err.Error())
			return
		}
	}
	if len(_appconfigGrowthFactor) > 0 {
		if err := assignInputField(input, "GrowthFactor", _appconfigGrowthFactor); err != nil {
			log.Errorf("invalid --growth-factor: %s", err.Error())
			return
		}
	}
	if len(_appconfigName) > 0 {
		input.Name = aws.String(_appconfigName)
	}
	if len(_appconfigDescription) > 0 {
		input.Description = aws.String(_appconfigDescription)
	}
	if len(_appconfigFinalBakeTimeInMinutes) > 0 {
		if err := assignInputField(input, "FinalBakeTimeInMinutes", _appconfigFinalBakeTimeInMinutes); err != nil {
			log.Errorf("invalid --final-bake-time-in-minutes: %s", err.Error())
			return
		}
	}
	if len(_appconfigGrowthType) > 0 {
		if err := assignInputField(input, "GrowthType", _appconfigGrowthType); err != nil {
			log.Errorf("invalid --growth-type: %s", err.Error())
			return
		}
	}
	if len(_appconfigReplicateTo) > 0 {
		if err := assignInputField(input, "ReplicateTo", _appconfigReplicateTo); err != nil {
			log.Errorf("invalid --replicate-to: %s", err.Error())
			return
		}
	}
	if len(_appconfigTags) > 0 {
		if err := assignInputField(input, "Tags", _appconfigTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDeploymentStrategy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an environment. For each application, you define one or more
// environments. An environment is a deployment group of AppConfig targets, such as
// applications in a Beta or Production environment. You can also define
// environments for application subcomponents such as the Web , Mobile and Back-end
// components for your application. You can configure Amazon CloudWatch alarms for
// each environment. The system monitors alarms during a configuration deployment.
// If an alarm is triggered, the system rolls back the configuration.
func appconfig_CreateEnvironment(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.CreateEnvironmentInput{
		// ApplicationId: *string, // Required
		// Name: *string, // Required
	}

	if len(_appconfigApplicationId) > 0 {
		input.ApplicationId = aws.String(_appconfigApplicationId)
	}
	if len(_appconfigName) > 0 {
		input.Name = aws.String(_appconfigName)
	}
	if len(_appconfigDescription) > 0 {
		input.Description = aws.String(_appconfigDescription)
	}
	if len(_appconfigMonitors) > 0 {
		if err := assignInputField(input, "Monitors", _appconfigMonitors); err != nil {
			log.Errorf("invalid --monitors: %s", err.Error())
			return
		}
	}
	if len(_appconfigTags) > 0 {
		if err := assignInputField(input, "Tags", _appconfigTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an AppConfig extension. An extension augments your ability to inject
// logic or behavior at different points during the AppConfig workflow of creating
// or deploying a configuration.
//
// You can create your own extensions or use the Amazon Web Services authored
// extensions provided by AppConfig. For an AppConfig extension that uses Lambda,
// you must create a Lambda function to perform any computation and processing
// defined in the extension. If you plan to create custom versions of the Amazon
// Web Services authored notification extensions, you only need to specify an
// Amazon Resource Name (ARN) in the Uri field for the new extension version.
//
// - For a custom EventBridge notification extension, enter the ARN of the
// EventBridge default events in the Uri field.
//
// - For a custom Amazon SNS notification extension, enter the ARN of an Amazon
// SNS topic in the Uri field.
//
// - For a custom Amazon SQS notification extension, enter the ARN of an Amazon
// SQS message queue in the Uri field.
//
// For more information about extensions, see [Extending workflows] in the AppConfig User Guide.
//
// [Extending workflows]: https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html
func appconfig_CreateExtension(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.CreateExtensionInput{
		// Actions: map[string][]types.Action, // Required
		// Name: *string, // Required
	}

	if len(_appconfigActions) > 0 {
		if err := assignInputField(input, "Actions", _appconfigActions); err != nil {
			log.Errorf("invalid --actions: %s", err.Error())
			return
		}
	}
	if len(_appconfigName) > 0 {
		input.Name = aws.String(_appconfigName)
	}
	if len(_appconfigDescription) > 0 {
		input.Description = aws.String(_appconfigDescription)
	}
	if len(_appconfigLatestVersionNumber) > 0 {
		if err := assignInputField(input, "LatestVersionNumber", _appconfigLatestVersionNumber); err != nil {
			log.Errorf("invalid --latest-version-number: %s", err.Error())
			return
		}
	}
	if len(_appconfigParameters) > 0 {
		if err := assignInputField(input, "Parameters", _appconfigParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_appconfigTags) > 0 {
		if err := assignInputField(input, "Tags", _appconfigTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateExtension(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// When you create an extension or configure an Amazon Web Services authored
// extension, you associate the extension with an AppConfig application,
// environment, or configuration profile. For example, you can choose to run the
// AppConfig deployment events to Amazon SNS Amazon Web Services authored extension
// and receive notifications on an Amazon SNS topic anytime a configuration
// deployment is started for a specific application. Defining which extension to
// associate with an AppConfig resource is called an extension association. An
// extension association is a specified relationship between an extension and an
// AppConfig resource, such as an application or a configuration profile. For more
// information about extensions and associations, see [Extending workflows]in the AppConfig User Guide.
//
// [Extending workflows]: https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html
func appconfig_CreateExtensionAssociation(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.CreateExtensionAssociationInput{
		// ExtensionIdentifier: *string, // Required
		// ResourceIdentifier: *string, // Required
	}

	if len(_appconfigExtensionIdentifier) > 0 {
		input.ExtensionIdentifier = aws.String(_appconfigExtensionIdentifier)
	}
	if len(_appconfigResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_appconfigResourceIdentifier)
	}
	if len(_appconfigExtensionVersionNumber) > 0 {
		if err := assignInputField(input, "ExtensionVersionNumber", _appconfigExtensionVersionNumber); err != nil {
			log.Errorf("invalid --extension-version-number: %s", err.Error())
			return
		}
	}
	if len(_appconfigParameters) > 0 {
		if err := assignInputField(input, "Parameters", _appconfigParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_appconfigTags) > 0 {
		if err := assignInputField(input, "Tags", _appconfigTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateExtensionAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new configuration in the AppConfig hosted configuration store. If
// you're creating a feature flag, we recommend you familiarize yourself with the
// JSON schema for feature flag data. For more information, see [Type reference for AWS.AppConfig.FeatureFlags]in the AppConfig
// User Guide.
//
// [Type reference for AWS.AppConfig.FeatureFlags]: https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-configuration-and-profile-feature-flags.html#appconfig-type-reference-feature-flags
func appconfig_CreateHostedConfigurationVersion(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.CreateHostedConfigurationVersionInput{
		// ApplicationId: *string, // Required
		// ConfigurationProfileId: *string, // Required
		// Content: []byte, // Required
		// ContentType: *string, // Required
	}

	if len(_appconfigApplicationId) > 0 {
		input.ApplicationId = aws.String(_appconfigApplicationId)
	}
	if len(_appconfigConfigurationProfileId) > 0 {
		input.ConfigurationProfileId = aws.String(_appconfigConfigurationProfileId)
	}
	if len(_appconfigContent) > 0 {
		if err := assignInputField(input, "Content", _appconfigContent); err != nil {
			log.Errorf("invalid --content: %s", err.Error())
			return
		}
	}
	if len(_appconfigContentType) > 0 {
		input.ContentType = aws.String(_appconfigContentType)
	}
	if len(_appconfigDescription) > 0 {
		input.Description = aws.String(_appconfigDescription)
	}
	if len(_appconfigLatestVersionNumber) > 0 {
		if err := assignInputField(input, "LatestVersionNumber", _appconfigLatestVersionNumber); err != nil {
			log.Errorf("invalid --latest-version-number: %s", err.Error())
			return
		}
	}
	if len(_appconfigVersionLabel) > 0 {
		input.VersionLabel = aws.String(_appconfigVersionLabel)
	}

	if resp, err := client.CreateHostedConfigurationVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an application.
func appconfig_DeleteApplication(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.DeleteApplicationInput{
		// ApplicationId: *string, // Required
	}

	if len(_appconfigApplicationId) > 0 {
		input.ApplicationId = aws.String(_appconfigApplicationId)
	}

	if resp, err := client.DeleteApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a configuration profile.
// To prevent users from unintentionally deleting actively-used configuration
// profiles, enable [deletion protection].
//
// [deletion protection]: https://docs.aws.amazon.com/appconfig/latest/userguide/deletion-protection.html
func appconfig_DeleteConfigurationProfile(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.DeleteConfigurationProfileInput{
		// ApplicationId: *string, // Required
		// ConfigurationProfileId: *string, // Required
	}

	if len(_appconfigApplicationId) > 0 {
		input.ApplicationId = aws.String(_appconfigApplicationId)
	}
	if len(_appconfigConfigurationProfileId) > 0 {
		input.ConfigurationProfileId = aws.String(_appconfigConfigurationProfileId)
	}
	if len(_appconfigDeletionProtectionCheck) > 0 {
		if err := assignInputField(input, "DeletionProtectionCheck", _appconfigDeletionProtectionCheck); err != nil {
			log.Errorf("invalid --deletion-protection-check: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteConfigurationProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a deployment strategy.
func appconfig_DeleteDeploymentStrategy(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.DeleteDeploymentStrategyInput{
		// DeploymentStrategyId: *string, // Required
	}

	if len(_appconfigDeploymentStrategyId) > 0 {
		input.DeploymentStrategyId = aws.String(_appconfigDeploymentStrategyId)
	}

	if resp, err := client.DeleteDeploymentStrategy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an environment.
// To prevent users from unintentionally deleting actively-used environments,
// enable [deletion protection].
//
// [deletion protection]: https://docs.aws.amazon.com/appconfig/latest/userguide/deletion-protection.html
func appconfig_DeleteEnvironment(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.DeleteEnvironmentInput{
		// ApplicationId: *string, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_appconfigApplicationId) > 0 {
		input.ApplicationId = aws.String(_appconfigApplicationId)
	}
	if len(_appconfigEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_appconfigEnvironmentId)
	}
	if len(_appconfigDeletionProtectionCheck) > 0 {
		if err := assignInputField(input, "DeletionProtectionCheck", _appconfigDeletionProtectionCheck); err != nil {
			log.Errorf("invalid --deletion-protection-check: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an AppConfig extension. You must delete all associations to an
// extension before you delete the extension.
func appconfig_DeleteExtension(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.DeleteExtensionInput{
		// ExtensionIdentifier: *string, // Required
	}

	if len(_appconfigExtensionIdentifier) > 0 {
		input.ExtensionIdentifier = aws.String(_appconfigExtensionIdentifier)
	}
	if len(_appconfigVersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _appconfigVersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteExtension(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an extension association. This action doesn't delete extensions defined
// in the association.
func appconfig_DeleteExtensionAssociation(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.DeleteExtensionAssociationInput{
		// ExtensionAssociationId: *string, // Required
	}

	if len(_appconfigExtensionAssociationId) > 0 {
		input.ExtensionAssociationId = aws.String(_appconfigExtensionAssociationId)
	}

	if resp, err := client.DeleteExtensionAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a version of a configuration from the AppConfig hosted configuration
// store.
func appconfig_DeleteHostedConfigurationVersion(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.DeleteHostedConfigurationVersionInput{
		// ApplicationId: *string, // Required
		// ConfigurationProfileId: *string, // Required
		// VersionNumber: *int32, // Required
	}

	if len(_appconfigApplicationId) > 0 {
		input.ApplicationId = aws.String(_appconfigApplicationId)
	}
	if len(_appconfigConfigurationProfileId) > 0 {
		input.ConfigurationProfileId = aws.String(_appconfigConfigurationProfileId)
	}
	if len(_appconfigVersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _appconfigVersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteHostedConfigurationVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the status of the DeletionProtection parameter.
func appconfig_GetAccountSettings(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.GetAccountSettingsInput{}

	if resp, err := client.GetAccountSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an application.
func appconfig_GetApplication(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.GetApplicationInput{
		// ApplicationId: *string, // Required
	}

	if len(_appconfigApplicationId) > 0 {
		input.ApplicationId = aws.String(_appconfigApplicationId)
	}

	if resp, err := client.GetApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// (Deprecated) Retrieves the latest deployed configuration.
// Note the following important information.
//
// - This API action is deprecated. Calls to receive configuration data should
// use the [StartConfigurationSession]and [GetLatestConfiguration]APIs instead.
//
// GetConfiguration
// - is a priced call. For more information, see [Pricing].
//
// Deprecated: This API has been deprecated in favor of the GetLatestConfiguration
// API used in conjunction with StartConfigurationSession.
//
// [GetLatestConfiguration]: https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_appconfigdata_GetLatestConfiguration.html
// [StartConfigurationSession]: https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_appconfigdata_StartConfigurationSession.html
// [Pricing]: https://aws.amazon.com/systems-manager/pricing/
func appconfig_GetConfiguration(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.GetConfigurationInput{
		// Application: *string, // Required
		// ClientId: *string, // Required
		// Configuration: *string, // Required
		// Environment: *string, // Required
	}

	if len(_appconfigApplication) > 0 {
		input.Application = aws.String(_appconfigApplication)
	}
	if len(_appconfigClientId) > 0 {
		input.ClientId = aws.String(_appconfigClientId)
	}
	if len(_appconfigConfiguration) > 0 {
		input.Configuration = aws.String(_appconfigConfiguration)
	}
	if len(_appconfigEnvironment) > 0 {
		input.Environment = aws.String(_appconfigEnvironment)
	}
	if len(_appconfigClientConfigurationVersion) > 0 {
		input.ClientConfigurationVersion = aws.String(_appconfigClientConfigurationVersion)
	}

	if resp, err := client.GetConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a configuration profile.
func appconfig_GetConfigurationProfile(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.GetConfigurationProfileInput{
		// ApplicationId: *string, // Required
		// ConfigurationProfileId: *string, // Required
	}

	if len(_appconfigApplicationId) > 0 {
		input.ApplicationId = aws.String(_appconfigApplicationId)
	}
	if len(_appconfigConfigurationProfileId) > 0 {
		input.ConfigurationProfileId = aws.String(_appconfigConfigurationProfileId)
	}

	if resp, err := client.GetConfigurationProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a configuration deployment.
func appconfig_GetDeployment(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.GetDeploymentInput{
		// ApplicationId: *string, // Required
		// DeploymentNumber: *int32, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_appconfigApplicationId) > 0 {
		input.ApplicationId = aws.String(_appconfigApplicationId)
	}
	if len(_appconfigDeploymentNumber) > 0 {
		if err := assignInputField(input, "DeploymentNumber", _appconfigDeploymentNumber); err != nil {
			log.Errorf("invalid --deployment-number: %s", err.Error())
			return
		}
	}
	if len(_appconfigEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_appconfigEnvironmentId)
	}

	if resp, err := client.GetDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a deployment strategy. A deployment strategy
// defines important criteria for rolling out your configuration to the designated
// targets. A deployment strategy includes the overall duration required, a
// percentage of targets to receive the deployment during each interval, an
// algorithm that defines how percentage grows, and bake time.
func appconfig_GetDeploymentStrategy(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.GetDeploymentStrategyInput{
		// DeploymentStrategyId: *string, // Required
	}

	if len(_appconfigDeploymentStrategyId) > 0 {
		input.DeploymentStrategyId = aws.String(_appconfigDeploymentStrategyId)
	}

	if resp, err := client.GetDeploymentStrategy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an environment. An environment is a deployment
// group of AppConfig applications, such as applications in a Production
// environment or in an EU_Region environment. Each configuration deployment
// targets an environment. You can enable one or more Amazon CloudWatch alarms for
// an environment. If an alarm is triggered during a deployment, AppConfig roles
// back the configuration.
func appconfig_GetEnvironment(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.GetEnvironmentInput{
		// ApplicationId: *string, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_appconfigApplicationId) > 0 {
		input.ApplicationId = aws.String(_appconfigApplicationId)
	}
	if len(_appconfigEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_appconfigEnvironmentId)
	}

	if resp, err := client.GetEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an AppConfig extension.
func appconfig_GetExtension(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.GetExtensionInput{
		// ExtensionIdentifier: *string, // Required
	}

	if len(_appconfigExtensionIdentifier) > 0 {
		input.ExtensionIdentifier = aws.String(_appconfigExtensionIdentifier)
	}
	if len(_appconfigVersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _appconfigVersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetExtension(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an AppConfig extension association. For more
// information about extensions and associations, see [Extending workflows]in the AppConfig User Guide.
//
// [Extending workflows]: https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html
func appconfig_GetExtensionAssociation(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.GetExtensionAssociationInput{
		// ExtensionAssociationId: *string, // Required
	}

	if len(_appconfigExtensionAssociationId) > 0 {
		input.ExtensionAssociationId = aws.String(_appconfigExtensionAssociationId)
	}

	if resp, err := client.GetExtensionAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a specific configuration version.
func appconfig_GetHostedConfigurationVersion(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.GetHostedConfigurationVersionInput{
		// ApplicationId: *string, // Required
		// ConfigurationProfileId: *string, // Required
		// VersionNumber: *int32, // Required
	}

	if len(_appconfigApplicationId) > 0 {
		input.ApplicationId = aws.String(_appconfigApplicationId)
	}
	if len(_appconfigConfigurationProfileId) > 0 {
		input.ConfigurationProfileId = aws.String(_appconfigConfigurationProfileId)
	}
	if len(_appconfigVersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _appconfigVersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetHostedConfigurationVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all applications in your Amazon Web Services account.
func appconfig_ListApplications(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.ListApplicationsInput{}

	if len(_appconfigMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appconfigMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appconfigNextToken) > 0 {
		input.NextToken = aws.String(_appconfigNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApplications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appconfig.ListApplicationsOutput
	p := appconfig.NewListApplicationsPaginator(client, input)
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

// Lists the configuration profiles for an application.
func appconfig_ListConfigurationProfiles(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.ListConfigurationProfilesInput{
		// ApplicationId: *string, // Required
	}

	if len(_appconfigApplicationId) > 0 {
		input.ApplicationId = aws.String(_appconfigApplicationId)
	}
	if len(_appconfigMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appconfigMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appconfigNextToken) > 0 {
		input.NextToken = aws.String(_appconfigNextToken)
	}
	if len(_appconfigType) > 0 {
		input.Type = aws.String(_appconfigType)
	}

	if disablePaginator() {
		if resp, err := client.ListConfigurationProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appconfig.ListConfigurationProfilesOutput
	p := appconfig.NewListConfigurationProfilesPaginator(client, input)
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

// Lists deployment strategies.
func appconfig_ListDeploymentStrategies(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.ListDeploymentStrategiesInput{}

	if len(_appconfigMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appconfigMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appconfigNextToken) > 0 {
		input.NextToken = aws.String(_appconfigNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDeploymentStrategies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appconfig.ListDeploymentStrategiesOutput
	p := appconfig.NewListDeploymentStrategiesPaginator(client, input)
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

// Lists the deployments for an environment in descending deployment number order.
func appconfig_ListDeployments(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.ListDeploymentsInput{
		// ApplicationId: *string, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_appconfigApplicationId) > 0 {
		input.ApplicationId = aws.String(_appconfigApplicationId)
	}
	if len(_appconfigEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_appconfigEnvironmentId)
	}
	if len(_appconfigMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appconfigMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appconfigNextToken) > 0 {
		input.NextToken = aws.String(_appconfigNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDeployments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appconfig.ListDeploymentsOutput
	p := appconfig.NewListDeploymentsPaginator(client, input)
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

// Lists the environments for an application.
func appconfig_ListEnvironments(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.ListEnvironmentsInput{
		// ApplicationId: *string, // Required
	}

	if len(_appconfigApplicationId) > 0 {
		input.ApplicationId = aws.String(_appconfigApplicationId)
	}
	if len(_appconfigMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appconfigMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appconfigNextToken) > 0 {
		input.NextToken = aws.String(_appconfigNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEnvironments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appconfig.ListEnvironmentsOutput
	p := appconfig.NewListEnvironmentsPaginator(client, input)
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

// Lists all AppConfig extension associations in the account. For more information
// about extensions and associations, see [Extending workflows]in the AppConfig User Guide.
//
// [Extending workflows]: https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html
func appconfig_ListExtensionAssociations(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.ListExtensionAssociationsInput{}

	if len(_appconfigExtensionIdentifier) > 0 {
		input.ExtensionIdentifier = aws.String(_appconfigExtensionIdentifier)
	}
	if len(_appconfigExtensionVersionNumber) > 0 {
		if err := assignInputField(input, "ExtensionVersionNumber", _appconfigExtensionVersionNumber); err != nil {
			log.Errorf("invalid --extension-version-number: %s", err.Error())
			return
		}
	}
	if len(_appconfigMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appconfigMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appconfigNextToken) > 0 {
		input.NextToken = aws.String(_appconfigNextToken)
	}
	if len(_appconfigResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_appconfigResourceIdentifier)
	}

	if disablePaginator() {
		if resp, err := client.ListExtensionAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appconfig.ListExtensionAssociationsOutput
	p := appconfig.NewListExtensionAssociationsPaginator(client, input)
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

// Lists all custom and Amazon Web Services authored AppConfig extensions in the
// account. For more information about extensions, see [Extending workflows]in the AppConfig User Guide.
//
// [Extending workflows]: https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html
func appconfig_ListExtensions(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.ListExtensionsInput{}

	if len(_appconfigMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appconfigMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appconfigName) > 0 {
		input.Name = aws.String(_appconfigName)
	}
	if len(_appconfigNextToken) > 0 {
		input.NextToken = aws.String(_appconfigNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListExtensions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appconfig.ListExtensionsOutput
	p := appconfig.NewListExtensionsPaginator(client, input)
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

// Lists configurations stored in the AppConfig hosted configuration store by
// version.
func appconfig_ListHostedConfigurationVersions(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.ListHostedConfigurationVersionsInput{
		// ApplicationId: *string, // Required
		// ConfigurationProfileId: *string, // Required
	}

	if len(_appconfigApplicationId) > 0 {
		input.ApplicationId = aws.String(_appconfigApplicationId)
	}
	if len(_appconfigConfigurationProfileId) > 0 {
		input.ConfigurationProfileId = aws.String(_appconfigConfigurationProfileId)
	}
	if len(_appconfigMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appconfigMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appconfigNextToken) > 0 {
		input.NextToken = aws.String(_appconfigNextToken)
	}
	if len(_appconfigVersionLabel) > 0 {
		input.VersionLabel = aws.String(_appconfigVersionLabel)
	}

	if disablePaginator() {
		if resp, err := client.ListHostedConfigurationVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appconfig.ListHostedConfigurationVersionsOutput
	p := appconfig.NewListHostedConfigurationVersionsPaginator(client, input)
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

// Retrieves the list of key-value tags assigned to the resource.
func appconfig_ListTagsForResource(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_appconfigResourceArn) > 0 {
		input.ResourceArn = aws.String(_appconfigResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a deployment.
func appconfig_StartDeployment(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.StartDeploymentInput{
		// ApplicationId: *string, // Required
		// ConfigurationProfileId: *string, // Required
		// ConfigurationVersion: *string, // Required
		// DeploymentStrategyId: *string, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_appconfigApplicationId) > 0 {
		input.ApplicationId = aws.String(_appconfigApplicationId)
	}
	if len(_appconfigConfigurationProfileId) > 0 {
		input.ConfigurationProfileId = aws.String(_appconfigConfigurationProfileId)
	}
	if len(_appconfigConfigurationVersion) > 0 {
		input.ConfigurationVersion = aws.String(_appconfigConfigurationVersion)
	}
	if len(_appconfigDeploymentStrategyId) > 0 {
		input.DeploymentStrategyId = aws.String(_appconfigDeploymentStrategyId)
	}
	if len(_appconfigEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_appconfigEnvironmentId)
	}
	if len(_appconfigDescription) > 0 {
		input.Description = aws.String(_appconfigDescription)
	}
	if len(_appconfigDynamicExtensionParameters) > 0 {
		if err := assignInputField(input, "DynamicExtensionParameters", _appconfigDynamicExtensionParameters); err != nil {
			log.Errorf("invalid --dynamic-extension-parameters: %s", err.Error())
			return
		}
	}
	if len(_appconfigKmsKeyIdentifier) > 0 {
		input.KmsKeyIdentifier = aws.String(_appconfigKmsKeyIdentifier)
	}
	if len(_appconfigTags) > 0 {
		if err := assignInputField(input, "Tags", _appconfigTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
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

// Stops a deployment. This API action works only on deployments that have a
// status of DEPLOYING , unless an AllowRevert parameter is supplied. If the
// AllowRevert parameter is supplied, the status of an in-progress deployment will
// be ROLLED_BACK . The status of a completed deployment will be REVERTED .
// AppConfig only allows a revert within 72 hours of deployment completion.
func appconfig_StopDeployment(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.StopDeploymentInput{
		// ApplicationId: *string, // Required
		// DeploymentNumber: *int32, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_appconfigApplicationId) > 0 {
		input.ApplicationId = aws.String(_appconfigApplicationId)
	}
	if len(_appconfigDeploymentNumber) > 0 {
		if err := assignInputField(input, "DeploymentNumber", _appconfigDeploymentNumber); err != nil {
			log.Errorf("invalid --deployment-number: %s", err.Error())
			return
		}
	}
	if len(_appconfigEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_appconfigEnvironmentId)
	}
	if len(_appconfigAllowRevert) > 0 {
		if err := assignInputField(input, "AllowRevert", _appconfigAllowRevert); err != nil {
			log.Errorf("invalid --allow-revert: %s", err.Error())
			return
		}
	}

	if resp, err := client.StopDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns metadata to an AppConfig resource. Tags help organize and categorize
// your AppConfig resources. Each tag consists of a key and an optional value, both
// of which you define. You can specify a maximum of 50 tags for a resource.
func appconfig_TagResource(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_appconfigResourceArn) > 0 {
		input.ResourceArn = aws.String(_appconfigResourceArn)
	}
	if len(_appconfigTags) > 0 {
		if err := assignInputField(input, "Tags", _appconfigTags); err != nil {
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

// Deletes a tag key and value from an AppConfig resource.
func appconfig_UntagResource(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_appconfigResourceArn) > 0 {
		input.ResourceArn = aws.String(_appconfigResourceArn)
	}
	if len(_appconfigTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _appconfigTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the value of the DeletionProtection parameter.
func appconfig_UpdateAccountSettings(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.UpdateAccountSettingsInput{}

	if len(_appconfigDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _appconfigDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAccountSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an application.
func appconfig_UpdateApplication(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.UpdateApplicationInput{
		// ApplicationId: *string, // Required
	}

	if len(_appconfigApplicationId) > 0 {
		input.ApplicationId = aws.String(_appconfigApplicationId)
	}
	if len(_appconfigDescription) > 0 {
		input.Description = aws.String(_appconfigDescription)
	}
	if len(_appconfigName) > 0 {
		input.Name = aws.String(_appconfigName)
	}

	if resp, err := client.UpdateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a configuration profile.
func appconfig_UpdateConfigurationProfile(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.UpdateConfigurationProfileInput{
		// ApplicationId: *string, // Required
		// ConfigurationProfileId: *string, // Required
	}

	if len(_appconfigApplicationId) > 0 {
		input.ApplicationId = aws.String(_appconfigApplicationId)
	}
	if len(_appconfigConfigurationProfileId) > 0 {
		input.ConfigurationProfileId = aws.String(_appconfigConfigurationProfileId)
	}
	if len(_appconfigDescription) > 0 {
		input.Description = aws.String(_appconfigDescription)
	}
	if len(_appconfigKmsKeyIdentifier) > 0 {
		input.KmsKeyIdentifier = aws.String(_appconfigKmsKeyIdentifier)
	}
	if len(_appconfigName) > 0 {
		input.Name = aws.String(_appconfigName)
	}
	if len(_appconfigRetrievalRoleArn) > 0 {
		input.RetrievalRoleArn = aws.String(_appconfigRetrievalRoleArn)
	}
	if len(_appconfigValidators) > 0 {
		if err := assignInputField(input, "Validators", _appconfigValidators); err != nil {
			log.Errorf("invalid --validators: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateConfigurationProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a deployment strategy.
func appconfig_UpdateDeploymentStrategy(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.UpdateDeploymentStrategyInput{
		// DeploymentStrategyId: *string, // Required
	}

	if len(_appconfigDeploymentStrategyId) > 0 {
		input.DeploymentStrategyId = aws.String(_appconfigDeploymentStrategyId)
	}
	if len(_appconfigDeploymentDurationInMinutes) > 0 {
		if err := assignInputField(input, "DeploymentDurationInMinutes", _appconfigDeploymentDurationInMinutes); err != nil {
			log.Errorf("invalid --deployment-duration-in-minutes: %s", err.Error())
			return
		}
	}
	if len(_appconfigDescription) > 0 {
		input.Description = aws.String(_appconfigDescription)
	}
	if len(_appconfigFinalBakeTimeInMinutes) > 0 {
		if err := assignInputField(input, "FinalBakeTimeInMinutes", _appconfigFinalBakeTimeInMinutes); err != nil {
			log.Errorf("invalid --final-bake-time-in-minutes: %s", err.Error())
			return
		}
	}
	if len(_appconfigGrowthFactor) > 0 {
		if err := assignInputField(input, "GrowthFactor", _appconfigGrowthFactor); err != nil {
			log.Errorf("invalid --growth-factor: %s", err.Error())
			return
		}
	}
	if len(_appconfigGrowthType) > 0 {
		if err := assignInputField(input, "GrowthType", _appconfigGrowthType); err != nil {
			log.Errorf("invalid --growth-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDeploymentStrategy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an environment.
func appconfig_UpdateEnvironment(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.UpdateEnvironmentInput{
		// ApplicationId: *string, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_appconfigApplicationId) > 0 {
		input.ApplicationId = aws.String(_appconfigApplicationId)
	}
	if len(_appconfigEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_appconfigEnvironmentId)
	}
	if len(_appconfigDescription) > 0 {
		input.Description = aws.String(_appconfigDescription)
	}
	if len(_appconfigMonitors) > 0 {
		if err := assignInputField(input, "Monitors", _appconfigMonitors); err != nil {
			log.Errorf("invalid --monitors: %s", err.Error())
			return
		}
	}
	if len(_appconfigName) > 0 {
		input.Name = aws.String(_appconfigName)
	}

	if resp, err := client.UpdateEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an AppConfig extension. For more information about extensions, see [Extending workflows] in
// the AppConfig User Guide.
//
// [Extending workflows]: https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html
func appconfig_UpdateExtension(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.UpdateExtensionInput{
		// ExtensionIdentifier: *string, // Required
	}

	if len(_appconfigExtensionIdentifier) > 0 {
		input.ExtensionIdentifier = aws.String(_appconfigExtensionIdentifier)
	}
	if len(_appconfigActions) > 0 {
		if err := assignInputField(input, "Actions", _appconfigActions); err != nil {
			log.Errorf("invalid --actions: %s", err.Error())
			return
		}
	}
	if len(_appconfigDescription) > 0 {
		input.Description = aws.String(_appconfigDescription)
	}
	if len(_appconfigParameters) > 0 {
		if err := assignInputField(input, "Parameters", _appconfigParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_appconfigVersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _appconfigVersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateExtension(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an association. For more information about extensions and associations,
// see [Extending workflows]in the AppConfig User Guide.
//
// [Extending workflows]: https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html
func appconfig_UpdateExtensionAssociation(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.UpdateExtensionAssociationInput{
		// ExtensionAssociationId: *string, // Required
	}

	if len(_appconfigExtensionAssociationId) > 0 {
		input.ExtensionAssociationId = aws.String(_appconfigExtensionAssociationId)
	}
	if len(_appconfigParameters) > 0 {
		if err := assignInputField(input, "Parameters", _appconfigParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateExtensionAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Uses the validators in a configuration profile to validate a configuration.
func appconfig_ValidateConfiguration(cfg aws.Config, client *appconfig.Client) {
	input := &appconfig.ValidateConfigurationInput{
		// ApplicationId: *string, // Required
		// ConfigurationProfileId: *string, // Required
		// ConfigurationVersion: *string, // Required
	}

	if len(_appconfigApplicationId) > 0 {
		input.ApplicationId = aws.String(_appconfigApplicationId)
	}
	if len(_appconfigConfigurationProfileId) > 0 {
		input.ConfigurationProfileId = aws.String(_appconfigConfigurationProfileId)
	}
	if len(_appconfigConfigurationVersion) > 0 {
		input.ConfigurationVersion = aws.String(_appconfigConfigurationVersion)
	}

	if resp, err := client.ValidateConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_appconfigCmd)
	_appconfigCmd.Flags().SortFlags = false

	_appconfigCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_appconfigCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_appconfigCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_appconfigCmd.Flags().StringVarP(&_appconfigActions, "actions", "", "", "Actions")
	_appconfigCmd.Flags().StringVarP(&_appconfigAllowRevert, "allow-revert", "", "", "Allow Revert")
	_appconfigCmd.Flags().StringVarP(&_appconfigApplication, "application", "", "", "Application")
	_appconfigCmd.Flags().StringVarP(&_appconfigApplicationId, "application-id", "", "", "Application ID")
	_appconfigCmd.Flags().StringVarP(&_appconfigClientConfigurationVersion, "client-configuration-version", "", "", "Client Configuration Version")
	_appconfigCmd.Flags().StringVarP(&_appconfigClientId, "client-id", "", "", "Client ID")
	_appconfigCmd.Flags().StringVarP(&_appconfigConfiguration, "configuration", "", "", "Configuration")
	_appconfigCmd.Flags().StringVarP(&_appconfigConfigurationProfileId, "configuration-profile-id", "", "", "Configuration Profile ID")
	_appconfigCmd.Flags().StringVarP(&_appconfigConfigurationVersion, "configuration-version", "", "", "Configuration Version")
	_appconfigCmd.Flags().StringVarP(&_appconfigContent, "content", "", "", "Content")
	_appconfigCmd.Flags().StringVarP(&_appconfigContentType, "content-type", "", "", "Content Type")
	_appconfigCmd.Flags().StringVarP(&_appconfigDeletionProtection, "deletion-protection", "", "", "Deletion Protection")
	_appconfigCmd.Flags().StringVarP(&_appconfigDeletionProtectionCheck, "deletion-protection-check", "", "", "Deletion Protection Check")
	_appconfigCmd.Flags().StringVarP(&_appconfigDeploymentDurationInMinutes, "deployment-duration-in-minutes", "", "", "Deployment Duration In Minutes")
	_appconfigCmd.Flags().StringVarP(&_appconfigDeploymentNumber, "deployment-number", "", "", "Deployment Number")
	_appconfigCmd.Flags().StringVarP(&_appconfigDeploymentStrategyId, "deployment-strategy-id", "", "", "Deployment Strategy ID")
	_appconfigCmd.Flags().StringVarP(&_appconfigDescription, "description", "", "", "Description")
	_appconfigCmd.Flags().StringVarP(&_appconfigDynamicExtensionParameters, "dynamic-extension-parameters", "", "", "Dynamic Extension Parameters")
	_appconfigCmd.Flags().StringVarP(&_appconfigEnvironment, "environment", "", "", "Environment")
	_appconfigCmd.Flags().StringVarP(&_appconfigEnvironmentId, "environment-id", "", "", "Environment ID")
	_appconfigCmd.Flags().StringVarP(&_appconfigExtensionAssociationId, "extension-association-id", "", "", "Extension Association ID")
	_appconfigCmd.Flags().StringVarP(&_appconfigExtensionIdentifier, "extension-identifier", "", "", "Extension Identifier")
	_appconfigCmd.Flags().StringVarP(&_appconfigExtensionVersionNumber, "extension-version-number", "", "", "Extension Version Number")
	_appconfigCmd.Flags().StringVarP(&_appconfigFinalBakeTimeInMinutes, "final-bake-time-in-minutes", "", "", "Final Bake Time In Minutes")
	_appconfigCmd.Flags().StringVarP(&_appconfigGrowthFactor, "growth-factor", "", "", "Growth Factor")
	_appconfigCmd.Flags().StringVarP(&_appconfigGrowthType, "growth-type", "", "", "Growth Type")
	_appconfigCmd.Flags().StringVarP(&_appconfigKmsKeyIdentifier, "kms-key-identifier", "", "", "KMS Key Identifier")
	_appconfigCmd.Flags().StringVarP(&_appconfigLatestVersionNumber, "latest-version-number", "", "", "Latest Version Number")
	_appconfigCmd.Flags().StringVarP(&_appconfigLocationUri, "location-uri", "", "", "Location URI")
	_appconfigCmd.Flags().StringVarP(&_appconfigMaxResults, "max-results", "", "", "Max Results")
	_appconfigCmd.Flags().StringVarP(&_appconfigMonitors, "monitors", "", "", "Monitors")
	_appconfigCmd.Flags().StringVarP(&_appconfigName, "name", "", "", "Name")
	_appconfigCmd.Flags().StringVarP(&_appconfigNextToken, "next-token", "", "", "Next Token")
	_appconfigCmd.Flags().StringVarP(&_appconfigParameters, "parameters", "", "", "Parameters")
	_appconfigCmd.Flags().StringVarP(&_appconfigReplicateTo, "replicate-to", "", "", "Replicate To")
	_appconfigCmd.Flags().StringVarP(&_appconfigResourceArn, "resource-arn", "", "", "Resource ARN")
	_appconfigCmd.Flags().StringVarP(&_appconfigResourceIdentifier, "resource-identifier", "", "", "Resource Identifier")
	_appconfigCmd.Flags().StringVarP(&_appconfigRetrievalRoleArn, "retrieval-role-arn", "", "", "Retrieval Role ARN")
	_appconfigCmd.Flags().StringSliceVarP(&_appconfigTagKeys, "tag-keys", "", nil, "Tag Keys")
	_appconfigCmd.Flags().StringVarP(&_appconfigTags, "tags", "", "", "Tags")
	_appconfigCmd.Flags().StringVarP(&_appconfigType, "type", "", "", "Type")
	_appconfigCmd.Flags().StringVarP(&_appconfigValidators, "validators", "", "", "Validators")
	_appconfigCmd.Flags().StringVarP(&_appconfigVersionLabel, "version-label", "", "", "Version Label")
	_appconfigCmd.Flags().StringVarP(&_appconfigVersionNumber, "version-number", "", "", "Version Number")

	_appconfigCmd.Flags().BoolVarP(&_appconfigCreateApplication, "create-application", "", false, "Create Application")
	_appconfigCmd.Flags().BoolVarP(&_appconfigCreateConfigurationProfile, "create-configuration-profile", "", false, "Create Configuration Profile")
	_appconfigCmd.Flags().BoolVarP(&_appconfigCreateDeploymentStrategy, "create-deployment-strategy", "", false, "Create Deployment Strategy")
	_appconfigCmd.Flags().BoolVarP(&_appconfigCreateEnvironment, "create-environment", "", false, "Create Environment")
	_appconfigCmd.Flags().BoolVarP(&_appconfigCreateExtension, "create-extension", "", false, "Create Extension")
	_appconfigCmd.Flags().BoolVarP(&_appconfigCreateExtensionAssociation, "create-extension-association", "", false, "Create Extension Association")
	_appconfigCmd.Flags().BoolVarP(&_appconfigCreateHostedConfigurationVersion, "create-hosted-configuration-version", "", false, "Create Hosted Configuration Version")
	_appconfigCmd.Flags().BoolVarP(&_appconfigDeleteApplication, "delete-application", "", false, "Delete Application")
	_appconfigCmd.Flags().BoolVarP(&_appconfigDeleteConfigurationProfile, "delete-configuration-profile", "", false, "Delete Configuration Profile")
	_appconfigCmd.Flags().BoolVarP(&_appconfigDeleteDeploymentStrategy, "delete-deployment-strategy", "", false, "Delete Deployment Strategy")
	_appconfigCmd.Flags().BoolVarP(&_appconfigDeleteEnvironment, "delete-environment", "", false, "Delete Environment")
	_appconfigCmd.Flags().BoolVarP(&_appconfigDeleteExtension, "delete-extension", "", false, "Delete Extension")
	_appconfigCmd.Flags().BoolVarP(&_appconfigDeleteExtensionAssociation, "delete-extension-association", "", false, "Delete Extension Association")
	_appconfigCmd.Flags().BoolVarP(&_appconfigDeleteHostedConfigurationVersion, "delete-hosted-configuration-version", "", false, "Delete Hosted Configuration Version")
	_appconfigCmd.Flags().BoolVarP(&_appconfigGetAccountSettings, "get-account-settings", "", false, "Get Account Settings")
	_appconfigCmd.Flags().BoolVarP(&_appconfigGetApplication, "get-application", "", false, "Get Application")
	_appconfigCmd.Flags().BoolVarP(&_appconfigGetConfiguration, "get-configuration", "", false, "Get Configuration")
	_appconfigCmd.Flags().BoolVarP(&_appconfigGetConfigurationProfile, "get-configuration-profile", "", false, "Get Configuration Profile")
	_appconfigCmd.Flags().BoolVarP(&_appconfigGetDeployment, "get-deployment", "", false, "Get Deployment")
	_appconfigCmd.Flags().BoolVarP(&_appconfigGetDeploymentStrategy, "get-deployment-strategy", "", false, "Get Deployment Strategy")
	_appconfigCmd.Flags().BoolVarP(&_appconfigGetEnvironment, "get-environment", "", false, "Get Environment")
	_appconfigCmd.Flags().BoolVarP(&_appconfigGetExtension, "get-extension", "", false, "Get Extension")
	_appconfigCmd.Flags().BoolVarP(&_appconfigGetExtensionAssociation, "get-extension-association", "", false, "Get Extension Association")
	_appconfigCmd.Flags().BoolVarP(&_appconfigGetHostedConfigurationVersion, "get-hosted-configuration-version", "", false, "Get Hosted Configuration Version")
	_appconfigCmd.Flags().BoolVarP(&_appconfigListApplications, "list-applications", "", false, "List Applications")
	_appconfigCmd.Flags().BoolVarP(&_appconfigListConfigurationProfiles, "list-configuration-profiles", "", false, "List Configuration Profiles")
	_appconfigCmd.Flags().BoolVarP(&_appconfigListDeploymentStrategies, "list-deployment-strategies", "", false, "List Deployment Strategies")
	_appconfigCmd.Flags().BoolVarP(&_appconfigListDeployments, "list-deployments", "", false, "List Deployments")
	_appconfigCmd.Flags().BoolVarP(&_appconfigListEnvironments, "list-environments", "", false, "List Environments")
	_appconfigCmd.Flags().BoolVarP(&_appconfigListExtensionAssociations, "list-extension-associations", "", false, "List Extension Associations")
	_appconfigCmd.Flags().BoolVarP(&_appconfigListExtensions, "list-extensions", "", false, "List Extensions")
	_appconfigCmd.Flags().BoolVarP(&_appconfigListHostedConfigurationVersions, "list-hosted-configuration-versions", "", false, "List Hosted Configuration Versions")
	_appconfigCmd.Flags().BoolVarP(&_appconfigListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_appconfigCmd.Flags().BoolVarP(&_appconfigStartDeployment, "start-deployment", "", false, "Start Deployment")
	_appconfigCmd.Flags().BoolVarP(&_appconfigStopDeployment, "stop-deployment", "", false, "Stop Deployment")
	_appconfigCmd.Flags().BoolVarP(&_appconfigTagResource, "tag-resource", "", false, "Tag Resource")
	_appconfigCmd.Flags().BoolVarP(&_appconfigUntagResource, "untag-resource", "", false, "Untag Resource")
	_appconfigCmd.Flags().BoolVarP(&_appconfigUpdateAccountSettings, "update-account-settings", "", false, "Update Account Settings")
	_appconfigCmd.Flags().BoolVarP(&_appconfigUpdateApplication, "update-application", "", false, "Update Application")
	_appconfigCmd.Flags().BoolVarP(&_appconfigUpdateConfigurationProfile, "update-configuration-profile", "", false, "Update Configuration Profile")
	_appconfigCmd.Flags().BoolVarP(&_appconfigUpdateDeploymentStrategy, "update-deployment-strategy", "", false, "Update Deployment Strategy")
	_appconfigCmd.Flags().BoolVarP(&_appconfigUpdateEnvironment, "update-environment", "", false, "Update Environment")
	_appconfigCmd.Flags().BoolVarP(&_appconfigUpdateExtension, "update-extension", "", false, "Update Extension")
	_appconfigCmd.Flags().BoolVarP(&_appconfigUpdateExtensionAssociation, "update-extension-association", "", false, "Update Extension Association")
	_appconfigCmd.Flags().BoolVarP(&_appconfigValidateConfiguration, "validate-configuration", "", false, "Validate Configuration")

}
