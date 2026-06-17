package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssmquicksetup"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ssmquicksetupCmd represents the ssmquicksetup command
var _ssmquicksetupCmd = &cobra.Command{
	Use:   "ssmquicksetup",
	Short: "AWS ssmquicksetup CLI",
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
		client := ssmquicksetup.NewFromConfig(cfg)
		if _ssmquicksetupCreateConfigurationManager {
			ssmquicksetup_CreateConfigurationManager(cfg, client)
			return
		}
		if _ssmquicksetupDeleteConfigurationManager {
			ssmquicksetup_DeleteConfigurationManager(cfg, client)
			return
		}
		if _ssmquicksetupGetConfiguration {
			ssmquicksetup_GetConfiguration(cfg, client)
			return
		}
		if _ssmquicksetupGetConfigurationManager {
			ssmquicksetup_GetConfigurationManager(cfg, client)
			return
		}
		if _ssmquicksetupGetServiceSettings {
			ssmquicksetup_GetServiceSettings(cfg, client)
			return
		}
		if _ssmquicksetupListConfigurationManagers {
			ssmquicksetup_ListConfigurationManagers(cfg, client)
			return
		}
		if _ssmquicksetupListConfigurations {
			ssmquicksetup_ListConfigurations(cfg, client)
			return
		}
		if _ssmquicksetupListQuickSetupTypes {
			ssmquicksetup_ListQuickSetupTypes(cfg, client)
			return
		}
		if _ssmquicksetupListTagsForResource {
			ssmquicksetup_ListTagsForResource(cfg, client)
			return
		}
		if _ssmquicksetupTagResource {
			ssmquicksetup_TagResource(cfg, client)
			return
		}
		if _ssmquicksetupUntagResource {
			ssmquicksetup_UntagResource(cfg, client)
			return
		}
		if _ssmquicksetupUpdateConfigurationDefinition {
			ssmquicksetup_UpdateConfigurationDefinition(cfg, client)
			return
		}
		if _ssmquicksetupUpdateConfigurationManager {
			ssmquicksetup_UpdateConfigurationManager(cfg, client)
			return
		}
		if _ssmquicksetupUpdateServiceSettings {
			ssmquicksetup_UpdateServiceSettings(cfg, client)
			return
		}

	},
}

var (
	_ssmquicksetupCreateConfigurationManager    bool
	_ssmquicksetupDeleteConfigurationManager    bool
	_ssmquicksetupGetConfiguration              bool
	_ssmquicksetupGetConfigurationManager       bool
	_ssmquicksetupGetServiceSettings            bool
	_ssmquicksetupListConfigurationManagers     bool
	_ssmquicksetupListConfigurations            bool
	_ssmquicksetupListQuickSetupTypes           bool
	_ssmquicksetupListTagsForResource           bool
	_ssmquicksetupTagResource                   bool
	_ssmquicksetupUntagResource                 bool
	_ssmquicksetupUpdateConfigurationDefinition bool
	_ssmquicksetupUpdateConfigurationManager    bool
	_ssmquicksetupUpdateServiceSettings         bool

	_ssmquicksetupConfigurationDefinitionId            string
	_ssmquicksetupConfigurationDefinitions             string
	_ssmquicksetupConfigurationId                      string
	_ssmquicksetupDescription                          string
	_ssmquicksetupExplorerEnablingRoleArn              string
	_ssmquicksetupFilters                              string
	_ssmquicksetupId                                   string
	_ssmquicksetupLocalDeploymentAdministrationRoleArn string
	_ssmquicksetupLocalDeploymentExecutionRoleName     string
	_ssmquicksetupManagerArn                           string
	_ssmquicksetupMaxItems                             string
	_ssmquicksetupName                                 string
	_ssmquicksetupParameters                           string
	_ssmquicksetupResourceArn                          string
	_ssmquicksetupStartingToken                        string
	_ssmquicksetupTagKeys                              []string
	_ssmquicksetupTags                                 string
	_ssmquicksetupTypeVersion                          string
)

// Creates a Quick Setup configuration manager resource. This object is a
// collection of desired state configurations for multiple configuration
// definitions and summaries describing the deployments of those definitions.
func ssmquicksetup_CreateConfigurationManager(cfg aws.Config, client *ssmquicksetup.Client) {
	input := &ssmquicksetup.CreateConfigurationManagerInput{
		// ConfigurationDefinitions: []types.ConfigurationDefinitionInput, // Required
	}

	if len(_ssmquicksetupConfigurationDefinitions) > 0 {
		if err := assignInputField(input, "ConfigurationDefinitions", _ssmquicksetupConfigurationDefinitions); err != nil {
			log.Errorf("invalid --configuration-definitions: %s", err.Error())
			return
		}
	}
	if len(_ssmquicksetupDescription) > 0 {
		input.Description = aws.String(_ssmquicksetupDescription)
	}
	if len(_ssmquicksetupName) > 0 {
		input.Name = aws.String(_ssmquicksetupName)
	}
	if len(_ssmquicksetupTags) > 0 {
		if err := assignInputField(input, "Tags", _ssmquicksetupTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConfigurationManager(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a configuration manager.
func ssmquicksetup_DeleteConfigurationManager(cfg aws.Config, client *ssmquicksetup.Client) {
	input := &ssmquicksetup.DeleteConfigurationManagerInput{
		// ManagerArn: *string, // Required
	}

	if len(_ssmquicksetupManagerArn) > 0 {
		input.ManagerArn = aws.String(_ssmquicksetupManagerArn)
	}

	if resp, err := client.DeleteConfigurationManager(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details about the specified configuration.
func ssmquicksetup_GetConfiguration(cfg aws.Config, client *ssmquicksetup.Client) {
	input := &ssmquicksetup.GetConfigurationInput{
		// ConfigurationId: *string, // Required
	}

	if len(_ssmquicksetupConfigurationId) > 0 {
		input.ConfigurationId = aws.String(_ssmquicksetupConfigurationId)
	}

	if resp, err := client.GetConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a configuration manager.
func ssmquicksetup_GetConfigurationManager(cfg aws.Config, client *ssmquicksetup.Client) {
	input := &ssmquicksetup.GetConfigurationManagerInput{
		// ManagerArn: *string, // Required
	}

	if len(_ssmquicksetupManagerArn) > 0 {
		input.ManagerArn = aws.String(_ssmquicksetupManagerArn)
	}

	if resp, err := client.GetConfigurationManager(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns settings configured for Quick Setup in the requesting Amazon Web
// Services account and Amazon Web Services Region.
func ssmquicksetup_GetServiceSettings(cfg aws.Config, client *ssmquicksetup.Client) {
	input := &ssmquicksetup.GetServiceSettingsInput{}

	if resp, err := client.GetServiceSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns Quick Setup configuration managers.
func ssmquicksetup_ListConfigurationManagers(cfg aws.Config, client *ssmquicksetup.Client) {
	input := &ssmquicksetup.ListConfigurationManagersInput{}

	if len(_ssmquicksetupFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmquicksetupFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmquicksetupMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _ssmquicksetupMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_ssmquicksetupStartingToken) > 0 {
		input.StartingToken = aws.String(_ssmquicksetupStartingToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConfigurationManagers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmquicksetup.ListConfigurationManagersOutput
	p := ssmquicksetup.NewListConfigurationManagersPaginator(client, input)
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

// Returns configurations deployed by Quick Setup in the requesting Amazon Web
// Services account and Amazon Web Services Region.
func ssmquicksetup_ListConfigurations(cfg aws.Config, client *ssmquicksetup.Client) {
	input := &ssmquicksetup.ListConfigurationsInput{}

	if len(_ssmquicksetupConfigurationDefinitionId) > 0 {
		input.ConfigurationDefinitionId = aws.String(_ssmquicksetupConfigurationDefinitionId)
	}
	if len(_ssmquicksetupFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmquicksetupFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmquicksetupManagerArn) > 0 {
		input.ManagerArn = aws.String(_ssmquicksetupManagerArn)
	}
	if len(_ssmquicksetupMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _ssmquicksetupMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_ssmquicksetupStartingToken) > 0 {
		input.StartingToken = aws.String(_ssmquicksetupStartingToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmquicksetup.ListConfigurationsOutput
	p := ssmquicksetup.NewListConfigurationsPaginator(client, input)
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

// Returns the available Quick Setup types.
func ssmquicksetup_ListQuickSetupTypes(cfg aws.Config, client *ssmquicksetup.Client) {
	input := &ssmquicksetup.ListQuickSetupTypesInput{}

	if resp, err := client.ListQuickSetupTypes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns tags assigned to the resource.
func ssmquicksetup_ListTagsForResource(cfg aws.Config, client *ssmquicksetup.Client) {
	input := &ssmquicksetup.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_ssmquicksetupResourceArn) > 0 {
		input.ResourceArn = aws.String(_ssmquicksetupResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns key-value pairs of metadata to Amazon Web Services resources.
func ssmquicksetup_TagResource(cfg aws.Config, client *ssmquicksetup.Client) {
	input := &ssmquicksetup.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_ssmquicksetupResourceArn) > 0 {
		input.ResourceArn = aws.String(_ssmquicksetupResourceArn)
	}
	if len(_ssmquicksetupTags) > 0 {
		if err := assignInputField(input, "Tags", _ssmquicksetupTags); err != nil {
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

// Removes tags from the specified resource.
func ssmquicksetup_UntagResource(cfg aws.Config, client *ssmquicksetup.Client) {
	input := &ssmquicksetup.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_ssmquicksetupResourceArn) > 0 {
		input.ResourceArn = aws.String(_ssmquicksetupResourceArn)
	}
	if len(_ssmquicksetupTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _ssmquicksetupTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Quick Setup configuration definition.
func ssmquicksetup_UpdateConfigurationDefinition(cfg aws.Config, client *ssmquicksetup.Client) {
	input := &ssmquicksetup.UpdateConfigurationDefinitionInput{
		// Id: *string, // Required
		// ManagerArn: *string, // Required
	}

	if len(_ssmquicksetupId) > 0 {
		input.Id = aws.String(_ssmquicksetupId)
	}
	if len(_ssmquicksetupManagerArn) > 0 {
		input.ManagerArn = aws.String(_ssmquicksetupManagerArn)
	}
	if len(_ssmquicksetupLocalDeploymentAdministrationRoleArn) > 0 {
		input.LocalDeploymentAdministrationRoleArn = aws.String(_ssmquicksetupLocalDeploymentAdministrationRoleArn)
	}
	if len(_ssmquicksetupLocalDeploymentExecutionRoleName) > 0 {
		input.LocalDeploymentExecutionRoleName = aws.String(_ssmquicksetupLocalDeploymentExecutionRoleName)
	}
	if len(_ssmquicksetupParameters) > 0 {
		if err := assignInputField(input, "Parameters", _ssmquicksetupParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_ssmquicksetupTypeVersion) > 0 {
		input.TypeVersion = aws.String(_ssmquicksetupTypeVersion)
	}

	if resp, err := client.UpdateConfigurationDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Quick Setup configuration manager.
func ssmquicksetup_UpdateConfigurationManager(cfg aws.Config, client *ssmquicksetup.Client) {
	input := &ssmquicksetup.UpdateConfigurationManagerInput{
		// ManagerArn: *string, // Required
	}

	if len(_ssmquicksetupManagerArn) > 0 {
		input.ManagerArn = aws.String(_ssmquicksetupManagerArn)
	}
	if len(_ssmquicksetupDescription) > 0 {
		input.Description = aws.String(_ssmquicksetupDescription)
	}
	if len(_ssmquicksetupName) > 0 {
		input.Name = aws.String(_ssmquicksetupName)
	}

	if resp, err := client.UpdateConfigurationManager(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates settings configured for Quick Setup.
func ssmquicksetup_UpdateServiceSettings(cfg aws.Config, client *ssmquicksetup.Client) {
	input := &ssmquicksetup.UpdateServiceSettingsInput{}

	if len(_ssmquicksetupExplorerEnablingRoleArn) > 0 {
		input.ExplorerEnablingRoleArn = aws.String(_ssmquicksetupExplorerEnablingRoleArn)
	}

	if resp, err := client.UpdateServiceSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_ssmquicksetupCmd)
	_ssmquicksetupCmd.Flags().SortFlags = false

	_ssmquicksetupCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_ssmquicksetupCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_ssmquicksetupCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_ssmquicksetupCmd.Flags().StringVarP(&_ssmquicksetupConfigurationDefinitionId, "configuration-definition-id", "", "", "Configuration Definition ID")
	_ssmquicksetupCmd.Flags().StringVarP(&_ssmquicksetupConfigurationDefinitions, "configuration-definitions", "", "", "Configuration Definitions")
	_ssmquicksetupCmd.Flags().StringVarP(&_ssmquicksetupConfigurationId, "configuration-id", "", "", "Configuration ID")
	_ssmquicksetupCmd.Flags().StringVarP(&_ssmquicksetupDescription, "description", "", "", "Description")
	_ssmquicksetupCmd.Flags().StringVarP(&_ssmquicksetupExplorerEnablingRoleArn, "explorer-enabling-role-arn", "", "", "Explorer Enabling Role ARN")
	_ssmquicksetupCmd.Flags().StringVarP(&_ssmquicksetupFilters, "filters", "", "", "Filters")
	_ssmquicksetupCmd.Flags().StringVarP(&_ssmquicksetupId, "id", "", "", "ID")
	_ssmquicksetupCmd.Flags().StringVarP(&_ssmquicksetupLocalDeploymentAdministrationRoleArn, "local-deployment-administration-role-arn", "", "", "Local Deployment Administration Role ARN")
	_ssmquicksetupCmd.Flags().StringVarP(&_ssmquicksetupLocalDeploymentExecutionRoleName, "local-deployment-execution-role-name", "", "", "Local Deployment Execution Role Name")
	_ssmquicksetupCmd.Flags().StringVarP(&_ssmquicksetupManagerArn, "manager-arn", "", "", "Manager ARN")
	_ssmquicksetupCmd.Flags().StringVarP(&_ssmquicksetupMaxItems, "max-items", "", "", "Max Items")
	_ssmquicksetupCmd.Flags().StringVarP(&_ssmquicksetupName, "name", "", "", "Name")
	_ssmquicksetupCmd.Flags().StringVarP(&_ssmquicksetupParameters, "parameters", "", "", "Parameters")
	_ssmquicksetupCmd.Flags().StringVarP(&_ssmquicksetupResourceArn, "resource-arn", "", "", "Resource ARN")
	_ssmquicksetupCmd.Flags().StringVarP(&_ssmquicksetupStartingToken, "starting-token", "", "", "Starting Token")
	_ssmquicksetupCmd.Flags().StringSliceVarP(&_ssmquicksetupTagKeys, "tag-keys", "", nil, "Tag Keys")
	_ssmquicksetupCmd.Flags().StringVarP(&_ssmquicksetupTags, "tags", "", "", "Tags")
	_ssmquicksetupCmd.Flags().StringVarP(&_ssmquicksetupTypeVersion, "type-version", "", "", "Type Version")

	_ssmquicksetupCmd.Flags().BoolVarP(&_ssmquicksetupCreateConfigurationManager, "create-configuration-manager", "", false, "Create Configuration Manager")
	_ssmquicksetupCmd.Flags().BoolVarP(&_ssmquicksetupDeleteConfigurationManager, "delete-configuration-manager", "", false, "Delete Configuration Manager")
	_ssmquicksetupCmd.Flags().BoolVarP(&_ssmquicksetupGetConfiguration, "get-configuration", "", false, "Get Configuration")
	_ssmquicksetupCmd.Flags().BoolVarP(&_ssmquicksetupGetConfigurationManager, "get-configuration-manager", "", false, "Get Configuration Manager")
	_ssmquicksetupCmd.Flags().BoolVarP(&_ssmquicksetupGetServiceSettings, "get-service-settings", "", false, "Get Service Settings")
	_ssmquicksetupCmd.Flags().BoolVarP(&_ssmquicksetupListConfigurationManagers, "list-configuration-managers", "", false, "List Configuration Managers")
	_ssmquicksetupCmd.Flags().BoolVarP(&_ssmquicksetupListConfigurations, "list-configurations", "", false, "List Configurations")
	_ssmquicksetupCmd.Flags().BoolVarP(&_ssmquicksetupListQuickSetupTypes, "list-quick-setup-types", "", false, "List Quick Setup Types")
	_ssmquicksetupCmd.Flags().BoolVarP(&_ssmquicksetupListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_ssmquicksetupCmd.Flags().BoolVarP(&_ssmquicksetupTagResource, "tag-resource", "", false, "Tag Resource")
	_ssmquicksetupCmd.Flags().BoolVarP(&_ssmquicksetupUntagResource, "untag-resource", "", false, "Untag Resource")
	_ssmquicksetupCmd.Flags().BoolVarP(&_ssmquicksetupUpdateConfigurationDefinition, "update-configuration-definition", "", false, "Update Configuration Definition")
	_ssmquicksetupCmd.Flags().BoolVarP(&_ssmquicksetupUpdateConfigurationManager, "update-configuration-manager", "", false, "Update Configuration Manager")
	_ssmquicksetupCmd.Flags().BoolVarP(&_ssmquicksetupUpdateServiceSettings, "update-service-settings", "", false, "Update Service Settings")

}
