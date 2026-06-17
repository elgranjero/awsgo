package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// serverlessapplicationrepositoryCmd represents the serverlessapplicationrepository command
var _serverlessapplicationrepositoryCmd = &cobra.Command{
	Use:   "serverlessapplicationrepository",
	Short: "AWS serverlessapplicationrepository CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := serverlessapplicationrepository.NewFromConfig(cfg)
		if _serverlessapplicationrepositoryCreateApplication {
			serverlessapplicationrepository_CreateApplication(cfg, client)
			return
		}
		if _serverlessapplicationrepositoryCreateApplicationVersion {
			serverlessapplicationrepository_CreateApplicationVersion(cfg, client)
			return
		}
		if _serverlessapplicationrepositoryCreateCloudFormationChangeSet {
			serverlessapplicationrepository_CreateCloudFormationChangeSet(cfg, client)
			return
		}
		if _serverlessapplicationrepositoryCreateCloudFormationTemplate {
			serverlessapplicationrepository_CreateCloudFormationTemplate(cfg, client)
			return
		}
		if _serverlessapplicationrepositoryDeleteApplication {
			serverlessapplicationrepository_DeleteApplication(cfg, client)
			return
		}
		if _serverlessapplicationrepositoryGetApplication {
			serverlessapplicationrepository_GetApplication(cfg, client)
			return
		}
		if _serverlessapplicationrepositoryGetApplicationPolicy {
			serverlessapplicationrepository_GetApplicationPolicy(cfg, client)
			return
		}
		if _serverlessapplicationrepositoryGetCloudFormationTemplate {
			serverlessapplicationrepository_GetCloudFormationTemplate(cfg, client)
			return
		}
		if _serverlessapplicationrepositoryListApplicationDependencies {
			serverlessapplicationrepository_ListApplicationDependencies(cfg, client)
			return
		}
		if _serverlessapplicationrepositoryListApplicationVersions {
			serverlessapplicationrepository_ListApplicationVersions(cfg, client)
			return
		}
		if _serverlessapplicationrepositoryListApplications {
			serverlessapplicationrepository_ListApplications(cfg, client)
			return
		}
		if _serverlessapplicationrepositoryPutApplicationPolicy {
			serverlessapplicationrepository_PutApplicationPolicy(cfg, client)
			return
		}
		if _serverlessapplicationrepositoryUnshareApplication {
			serverlessapplicationrepository_UnshareApplication(cfg, client)
			return
		}
		if _serverlessapplicationrepositoryUpdateApplication {
			serverlessapplicationrepository_UpdateApplication(cfg, client)
			return
		}

	},
}

var (
	_serverlessapplicationrepositoryCreateApplication             bool
	_serverlessapplicationrepositoryCreateApplicationVersion      bool
	_serverlessapplicationrepositoryCreateCloudFormationChangeSet bool
	_serverlessapplicationrepositoryCreateCloudFormationTemplate  bool
	_serverlessapplicationrepositoryDeleteApplication             bool
	_serverlessapplicationrepositoryGetApplication                bool
	_serverlessapplicationrepositoryGetApplicationPolicy          bool
	_serverlessapplicationrepositoryGetCloudFormationTemplate     bool
	_serverlessapplicationrepositoryListApplicationDependencies   bool
	_serverlessapplicationrepositoryListApplicationVersions       bool
	_serverlessapplicationrepositoryListApplications              bool
	_serverlessapplicationrepositoryPutApplicationPolicy          bool
	_serverlessapplicationrepositoryUnshareApplication            bool
	_serverlessapplicationrepositoryUpdateApplication             bool

	_serverlessapplicationrepositoryApplicationId         string
	_serverlessapplicationrepositoryAuthor                string
	_serverlessapplicationrepositoryCapabilities          []string
	_serverlessapplicationrepositoryChangeSetName         string
	_serverlessapplicationrepositoryClientToken           string
	_serverlessapplicationrepositoryDescription           string
	_serverlessapplicationrepositoryHomePageUrl           string
	_serverlessapplicationrepositoryLabels                []string
	_serverlessapplicationrepositoryLicenseBody           string
	_serverlessapplicationrepositoryLicenseUrl            string
	_serverlessapplicationrepositoryMaxItems              string
	_serverlessapplicationrepositoryName                  string
	_serverlessapplicationrepositoryNextToken             string
	_serverlessapplicationrepositoryNotificationArns      []string
	_serverlessapplicationrepositoryOrganizationId        string
	_serverlessapplicationrepositoryParameterOverrides    string
	_serverlessapplicationrepositoryReadmeBody            string
	_serverlessapplicationrepositoryReadmeUrl             string
	_serverlessapplicationrepositoryResourceTypes         []string
	_serverlessapplicationrepositoryRollbackConfiguration string
	_serverlessapplicationrepositorySemanticVersion       string
	_serverlessapplicationrepositorySourceCodeArchiveUrl  string
	_serverlessapplicationrepositorySourceCodeUrl         string
	_serverlessapplicationrepositorySpdxLicenseId         string
	_serverlessapplicationrepositoryStackName             string
	_serverlessapplicationrepositoryStatements            string
	_serverlessapplicationrepositoryTags                  string
	_serverlessapplicationrepositoryTemplateBody          string
	_serverlessapplicationrepositoryTemplateId            string
	_serverlessapplicationrepositoryTemplateUrl           string
)

// Creates an application, optionally including an AWS SAM file to create the
// first application version in the same call.
func serverlessapplicationrepository_CreateApplication(cfg aws.Config, client *serverlessapplicationrepository.Client) {
	input := &serverlessapplicationrepository.CreateApplicationInput{
		// Author: *string, // Required
		// Description: *string, // Required
		// Name: *string, // Required
	}

	if len(_serverlessapplicationrepositoryAuthor) > 0 {
		input.Author = aws.String(_serverlessapplicationrepositoryAuthor)
	}
	if len(_serverlessapplicationrepositoryDescription) > 0 {
		input.Description = aws.String(_serverlessapplicationrepositoryDescription)
	}
	if len(_serverlessapplicationrepositoryName) > 0 {
		input.Name = aws.String(_serverlessapplicationrepositoryName)
	}
	if len(_serverlessapplicationrepositoryHomePageUrl) > 0 {
		input.HomePageUrl = aws.String(_serverlessapplicationrepositoryHomePageUrl)
	}
	if len(_serverlessapplicationrepositoryLabels) > 0 {
		input.Labels = append([]string(nil), _serverlessapplicationrepositoryLabels...)
	}
	if len(_serverlessapplicationrepositoryLicenseBody) > 0 {
		input.LicenseBody = aws.String(_serverlessapplicationrepositoryLicenseBody)
	}
	if len(_serverlessapplicationrepositoryLicenseUrl) > 0 {
		input.LicenseUrl = aws.String(_serverlessapplicationrepositoryLicenseUrl)
	}
	if len(_serverlessapplicationrepositoryReadmeBody) > 0 {
		input.ReadmeBody = aws.String(_serverlessapplicationrepositoryReadmeBody)
	}
	if len(_serverlessapplicationrepositoryReadmeUrl) > 0 {
		input.ReadmeUrl = aws.String(_serverlessapplicationrepositoryReadmeUrl)
	}
	if len(_serverlessapplicationrepositorySemanticVersion) > 0 {
		input.SemanticVersion = aws.String(_serverlessapplicationrepositorySemanticVersion)
	}
	if len(_serverlessapplicationrepositorySourceCodeArchiveUrl) > 0 {
		input.SourceCodeArchiveUrl = aws.String(_serverlessapplicationrepositorySourceCodeArchiveUrl)
	}
	if len(_serverlessapplicationrepositorySourceCodeUrl) > 0 {
		input.SourceCodeUrl = aws.String(_serverlessapplicationrepositorySourceCodeUrl)
	}
	if len(_serverlessapplicationrepositorySpdxLicenseId) > 0 {
		input.SpdxLicenseId = aws.String(_serverlessapplicationrepositorySpdxLicenseId)
	}
	if len(_serverlessapplicationrepositoryTemplateBody) > 0 {
		input.TemplateBody = aws.String(_serverlessapplicationrepositoryTemplateBody)
	}
	if len(_serverlessapplicationrepositoryTemplateUrl) > 0 {
		input.TemplateUrl = aws.String(_serverlessapplicationrepositoryTemplateUrl)
	}

	if resp, err := client.CreateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an application version.
func serverlessapplicationrepository_CreateApplicationVersion(cfg aws.Config, client *serverlessapplicationrepository.Client) {
	input := &serverlessapplicationrepository.CreateApplicationVersionInput{
		// ApplicationId: *string, // Required
		// SemanticVersion: *string, // Required
	}

	if len(_serverlessapplicationrepositoryApplicationId) > 0 {
		input.ApplicationId = aws.String(_serverlessapplicationrepositoryApplicationId)
	}
	if len(_serverlessapplicationrepositorySemanticVersion) > 0 {
		input.SemanticVersion = aws.String(_serverlessapplicationrepositorySemanticVersion)
	}
	if len(_serverlessapplicationrepositorySourceCodeArchiveUrl) > 0 {
		input.SourceCodeArchiveUrl = aws.String(_serverlessapplicationrepositorySourceCodeArchiveUrl)
	}
	if len(_serverlessapplicationrepositorySourceCodeUrl) > 0 {
		input.SourceCodeUrl = aws.String(_serverlessapplicationrepositorySourceCodeUrl)
	}
	if len(_serverlessapplicationrepositoryTemplateBody) > 0 {
		input.TemplateBody = aws.String(_serverlessapplicationrepositoryTemplateBody)
	}
	if len(_serverlessapplicationrepositoryTemplateUrl) > 0 {
		input.TemplateUrl = aws.String(_serverlessapplicationrepositoryTemplateUrl)
	}

	if resp, err := client.CreateApplicationVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an AWS CloudFormation change set for the given application.
func serverlessapplicationrepository_CreateCloudFormationChangeSet(cfg aws.Config, client *serverlessapplicationrepository.Client) {
	input := &serverlessapplicationrepository.CreateCloudFormationChangeSetInput{
		// ApplicationId: *string, // Required
		// StackName: *string, // Required
	}

	if len(_serverlessapplicationrepositoryApplicationId) > 0 {
		input.ApplicationId = aws.String(_serverlessapplicationrepositoryApplicationId)
	}
	if len(_serverlessapplicationrepositoryStackName) > 0 {
		input.StackName = aws.String(_serverlessapplicationrepositoryStackName)
	}
	if len(_serverlessapplicationrepositoryCapabilities) > 0 {
		input.Capabilities = append([]string(nil), _serverlessapplicationrepositoryCapabilities...)
	}
	if len(_serverlessapplicationrepositoryChangeSetName) > 0 {
		input.ChangeSetName = aws.String(_serverlessapplicationrepositoryChangeSetName)
	}
	if len(_serverlessapplicationrepositoryClientToken) > 0 {
		input.ClientToken = aws.String(_serverlessapplicationrepositoryClientToken)
	}
	if len(_serverlessapplicationrepositoryDescription) > 0 {
		input.Description = aws.String(_serverlessapplicationrepositoryDescription)
	}
	if len(_serverlessapplicationrepositoryNotificationArns) > 0 {
		input.NotificationArns = append([]string(nil), _serverlessapplicationrepositoryNotificationArns...)
	}
	if len(_serverlessapplicationrepositoryParameterOverrides) > 0 {
		if err := assignInputField(input, "ParameterOverrides", _serverlessapplicationrepositoryParameterOverrides); err != nil {
			log.Errorf("invalid --parameter-overrides: %s", err.Error())
			return
		}
	}
	if len(_serverlessapplicationrepositoryResourceTypes) > 0 {
		input.ResourceTypes = append([]string(nil), _serverlessapplicationrepositoryResourceTypes...)
	}
	if len(_serverlessapplicationrepositoryRollbackConfiguration) > 0 {
		if err := assignInputField(input, "RollbackConfiguration", _serverlessapplicationrepositoryRollbackConfiguration); err != nil {
			log.Errorf("invalid --rollback-configuration: %s", err.Error())
			return
		}
	}
	if len(_serverlessapplicationrepositorySemanticVersion) > 0 {
		input.SemanticVersion = aws.String(_serverlessapplicationrepositorySemanticVersion)
	}
	if len(_serverlessapplicationrepositoryTags) > 0 {
		if err := assignInputField(input, "Tags", _serverlessapplicationrepositoryTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_serverlessapplicationrepositoryTemplateId) > 0 {
		input.TemplateId = aws.String(_serverlessapplicationrepositoryTemplateId)
	}

	if resp, err := client.CreateCloudFormationChangeSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an AWS CloudFormation template.
func serverlessapplicationrepository_CreateCloudFormationTemplate(cfg aws.Config, client *serverlessapplicationrepository.Client) {
	input := &serverlessapplicationrepository.CreateCloudFormationTemplateInput{
		// ApplicationId: *string, // Required
	}

	if len(_serverlessapplicationrepositoryApplicationId) > 0 {
		input.ApplicationId = aws.String(_serverlessapplicationrepositoryApplicationId)
	}
	if len(_serverlessapplicationrepositorySemanticVersion) > 0 {
		input.SemanticVersion = aws.String(_serverlessapplicationrepositorySemanticVersion)
	}

	if resp, err := client.CreateCloudFormationTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified application.
func serverlessapplicationrepository_DeleteApplication(cfg aws.Config, client *serverlessapplicationrepository.Client) {
	input := &serverlessapplicationrepository.DeleteApplicationInput{
		// ApplicationId: *string, // Required
	}

	if len(_serverlessapplicationrepositoryApplicationId) > 0 {
		input.ApplicationId = aws.String(_serverlessapplicationrepositoryApplicationId)
	}

	if resp, err := client.DeleteApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the specified application.
func serverlessapplicationrepository_GetApplication(cfg aws.Config, client *serverlessapplicationrepository.Client) {
	input := &serverlessapplicationrepository.GetApplicationInput{
		// ApplicationId: *string, // Required
	}

	if len(_serverlessapplicationrepositoryApplicationId) > 0 {
		input.ApplicationId = aws.String(_serverlessapplicationrepositoryApplicationId)
	}
	if len(_serverlessapplicationrepositorySemanticVersion) > 0 {
		input.SemanticVersion = aws.String(_serverlessapplicationrepositorySemanticVersion)
	}

	if resp, err := client.GetApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the policy for the application.
func serverlessapplicationrepository_GetApplicationPolicy(cfg aws.Config, client *serverlessapplicationrepository.Client) {
	input := &serverlessapplicationrepository.GetApplicationPolicyInput{
		// ApplicationId: *string, // Required
	}

	if len(_serverlessapplicationrepositoryApplicationId) > 0 {
		input.ApplicationId = aws.String(_serverlessapplicationrepositoryApplicationId)
	}

	if resp, err := client.GetApplicationPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the specified AWS CloudFormation template.
func serverlessapplicationrepository_GetCloudFormationTemplate(cfg aws.Config, client *serverlessapplicationrepository.Client) {
	input := &serverlessapplicationrepository.GetCloudFormationTemplateInput{
		// ApplicationId: *string, // Required
		// TemplateId: *string, // Required
	}

	if len(_serverlessapplicationrepositoryApplicationId) > 0 {
		input.ApplicationId = aws.String(_serverlessapplicationrepositoryApplicationId)
	}
	if len(_serverlessapplicationrepositoryTemplateId) > 0 {
		input.TemplateId = aws.String(_serverlessapplicationrepositoryTemplateId)
	}

	if resp, err := client.GetCloudFormationTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the list of applications nested in the containing application.
func serverlessapplicationrepository_ListApplicationDependencies(cfg aws.Config, client *serverlessapplicationrepository.Client) {
	input := &serverlessapplicationrepository.ListApplicationDependenciesInput{
		// ApplicationId: *string, // Required
	}

	if len(_serverlessapplicationrepositoryApplicationId) > 0 {
		input.ApplicationId = aws.String(_serverlessapplicationrepositoryApplicationId)
	}
	if len(_serverlessapplicationrepositoryMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _serverlessapplicationrepositoryMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_serverlessapplicationrepositoryNextToken) > 0 {
		input.NextToken = aws.String(_serverlessapplicationrepositoryNextToken)
	}
	if len(_serverlessapplicationrepositorySemanticVersion) > 0 {
		input.SemanticVersion = aws.String(_serverlessapplicationrepositorySemanticVersion)
	}

	if disablePaginator() {
		if resp, err := client.ListApplicationDependencies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*serverlessapplicationrepository.ListApplicationDependenciesOutput
	p := serverlessapplicationrepository.NewListApplicationDependenciesPaginator(client, input)
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

// Lists versions for the specified application.
func serverlessapplicationrepository_ListApplicationVersions(cfg aws.Config, client *serverlessapplicationrepository.Client) {
	input := &serverlessapplicationrepository.ListApplicationVersionsInput{
		// ApplicationId: *string, // Required
	}

	if len(_serverlessapplicationrepositoryApplicationId) > 0 {
		input.ApplicationId = aws.String(_serverlessapplicationrepositoryApplicationId)
	}
	if len(_serverlessapplicationrepositoryMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _serverlessapplicationrepositoryMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_serverlessapplicationrepositoryNextToken) > 0 {
		input.NextToken = aws.String(_serverlessapplicationrepositoryNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApplicationVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*serverlessapplicationrepository.ListApplicationVersionsOutput
	p := serverlessapplicationrepository.NewListApplicationVersionsPaginator(client, input)
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

// Lists applications owned by the requester.
func serverlessapplicationrepository_ListApplications(cfg aws.Config, client *serverlessapplicationrepository.Client) {
	input := &serverlessapplicationrepository.ListApplicationsInput{}

	if len(_serverlessapplicationrepositoryMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _serverlessapplicationrepositoryMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_serverlessapplicationrepositoryNextToken) > 0 {
		input.NextToken = aws.String(_serverlessapplicationrepositoryNextToken)
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

	var results []*serverlessapplicationrepository.ListApplicationsOutput
	p := serverlessapplicationrepository.NewListApplicationsPaginator(client, input)
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

// Sets the permission policy for an application. For the list of actions
// supported for this operation, see [Application Permissions].
//
// [Application Permissions]: https://docs.aws.amazon.com/serverlessrepo/latest/devguide/access-control-resource-based.html#application-permissions
func serverlessapplicationrepository_PutApplicationPolicy(cfg aws.Config, client *serverlessapplicationrepository.Client) {
	input := &serverlessapplicationrepository.PutApplicationPolicyInput{
		// ApplicationId: *string, // Required
		// Statements: []types.ApplicationPolicyStatement, // Required
	}

	if len(_serverlessapplicationrepositoryApplicationId) > 0 {
		input.ApplicationId = aws.String(_serverlessapplicationrepositoryApplicationId)
	}
	if len(_serverlessapplicationrepositoryStatements) > 0 {
		if err := assignInputField(input, "Statements", _serverlessapplicationrepositoryStatements); err != nil {
			log.Errorf("invalid --statements: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutApplicationPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Unshares an application from an AWS Organization.
// This operation can be called only from the organization's master account.
func serverlessapplicationrepository_UnshareApplication(cfg aws.Config, client *serverlessapplicationrepository.Client) {
	input := &serverlessapplicationrepository.UnshareApplicationInput{
		// ApplicationId: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_serverlessapplicationrepositoryApplicationId) > 0 {
		input.ApplicationId = aws.String(_serverlessapplicationrepositoryApplicationId)
	}
	if len(_serverlessapplicationrepositoryOrganizationId) > 0 {
		input.OrganizationId = aws.String(_serverlessapplicationrepositoryOrganizationId)
	}

	if resp, err := client.UnshareApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified application.
func serverlessapplicationrepository_UpdateApplication(cfg aws.Config, client *serverlessapplicationrepository.Client) {
	input := &serverlessapplicationrepository.UpdateApplicationInput{
		// ApplicationId: *string, // Required
	}

	if len(_serverlessapplicationrepositoryApplicationId) > 0 {
		input.ApplicationId = aws.String(_serverlessapplicationrepositoryApplicationId)
	}
	if len(_serverlessapplicationrepositoryAuthor) > 0 {
		input.Author = aws.String(_serverlessapplicationrepositoryAuthor)
	}
	if len(_serverlessapplicationrepositoryDescription) > 0 {
		input.Description = aws.String(_serverlessapplicationrepositoryDescription)
	}
	if len(_serverlessapplicationrepositoryHomePageUrl) > 0 {
		input.HomePageUrl = aws.String(_serverlessapplicationrepositoryHomePageUrl)
	}
	if len(_serverlessapplicationrepositoryLabels) > 0 {
		input.Labels = append([]string(nil), _serverlessapplicationrepositoryLabels...)
	}
	if len(_serverlessapplicationrepositoryReadmeBody) > 0 {
		input.ReadmeBody = aws.String(_serverlessapplicationrepositoryReadmeBody)
	}
	if len(_serverlessapplicationrepositoryReadmeUrl) > 0 {
		input.ReadmeUrl = aws.String(_serverlessapplicationrepositoryReadmeUrl)
	}

	if resp, err := client.UpdateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_serverlessapplicationrepositoryCmd)
	_serverlessapplicationrepositoryCmd.Flags().SortFlags = false

	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositoryApplicationId, "application-id", "", "", "Application ID")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositoryAuthor, "author", "", "", "Author")
	_serverlessapplicationrepositoryCmd.Flags().StringSliceVarP(&_serverlessapplicationrepositoryCapabilities, "capabilities", "", nil, "Capabilities")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositoryChangeSetName, "change-set-name", "", "", "Change Set Name")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositoryClientToken, "client-token", "", "", "Client Token")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositoryDescription, "description", "", "", "Description")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositoryHomePageUrl, "home-page-url", "", "", "Home Page URL")
	_serverlessapplicationrepositoryCmd.Flags().StringSliceVarP(&_serverlessapplicationrepositoryLabels, "labels", "", nil, "Labels")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositoryLicenseBody, "license-body", "", "", "License Body")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositoryLicenseUrl, "license-url", "", "", "License URL")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositoryMaxItems, "max-items", "", "", "Max Items")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositoryName, "name", "", "", "Name")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositoryNextToken, "next-token", "", "", "Next Token")
	_serverlessapplicationrepositoryCmd.Flags().StringSliceVarP(&_serverlessapplicationrepositoryNotificationArns, "notification-arns", "", nil, "Notification Arns")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositoryOrganizationId, "organization-id", "", "", "Organization ID")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositoryParameterOverrides, "parameter-overrides", "", "", "Parameter Overrides")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositoryReadmeBody, "readme-body", "", "", "Readme Body")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositoryReadmeUrl, "readme-url", "", "", "Readme URL")
	_serverlessapplicationrepositoryCmd.Flags().StringSliceVarP(&_serverlessapplicationrepositoryResourceTypes, "resource-types", "", nil, "Resource Types")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositoryRollbackConfiguration, "rollback-configuration", "", "", "Rollback Configuration")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositorySemanticVersion, "semantic-version", "", "", "Semantic Version")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositorySourceCodeArchiveUrl, "source-code-archive-url", "", "", "Source Code Archive URL")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositorySourceCodeUrl, "source-code-url", "", "", "Source Code URL")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositorySpdxLicenseId, "spdx-license-id", "", "", "Spdx License ID")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositoryStackName, "stack-name", "", "", "Stack Name")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositoryStatements, "statements", "", "", "Statements")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositoryTags, "tags", "", "", "Tags")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositoryTemplateBody, "template-body", "", "", "Template Body")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositoryTemplateId, "template-id", "", "", "Template ID")
	_serverlessapplicationrepositoryCmd.Flags().StringVarP(&_serverlessapplicationrepositoryTemplateUrl, "template-url", "", "", "Template URL")

	_serverlessapplicationrepositoryCmd.Flags().BoolVarP(&_serverlessapplicationrepositoryCreateApplication, "create-application", "", false, "Create Application")
	_serverlessapplicationrepositoryCmd.Flags().BoolVarP(&_serverlessapplicationrepositoryCreateApplicationVersion, "create-application-version", "", false, "Create Application Version")
	_serverlessapplicationrepositoryCmd.Flags().BoolVarP(&_serverlessapplicationrepositoryCreateCloudFormationChangeSet, "create-cloud-formation-change-set", "", false, "Create Cloud Formation Change Set")
	_serverlessapplicationrepositoryCmd.Flags().BoolVarP(&_serverlessapplicationrepositoryCreateCloudFormationTemplate, "create-cloud-formation-template", "", false, "Create Cloud Formation Template")
	_serverlessapplicationrepositoryCmd.Flags().BoolVarP(&_serverlessapplicationrepositoryDeleteApplication, "delete-application", "", false, "Delete Application")
	_serverlessapplicationrepositoryCmd.Flags().BoolVarP(&_serverlessapplicationrepositoryGetApplication, "get-application", "", false, "Get Application")
	_serverlessapplicationrepositoryCmd.Flags().BoolVarP(&_serverlessapplicationrepositoryGetApplicationPolicy, "get-application-policy", "", false, "Get Application Policy")
	_serverlessapplicationrepositoryCmd.Flags().BoolVarP(&_serverlessapplicationrepositoryGetCloudFormationTemplate, "get-cloud-formation-template", "", false, "Get Cloud Formation Template")
	_serverlessapplicationrepositoryCmd.Flags().BoolVarP(&_serverlessapplicationrepositoryListApplicationDependencies, "list-application-dependencies", "", false, "List Application Dependencies")
	_serverlessapplicationrepositoryCmd.Flags().BoolVarP(&_serverlessapplicationrepositoryListApplicationVersions, "list-application-versions", "", false, "List Application Versions")
	_serverlessapplicationrepositoryCmd.Flags().BoolVarP(&_serverlessapplicationrepositoryListApplications, "list-applications", "", false, "List Applications")
	_serverlessapplicationrepositoryCmd.Flags().BoolVarP(&_serverlessapplicationrepositoryPutApplicationPolicy, "put-application-policy", "", false, "Put Application Policy")
	_serverlessapplicationrepositoryCmd.Flags().BoolVarP(&_serverlessapplicationrepositoryUnshareApplication, "unshare-application", "", false, "Unshare Application")
	_serverlessapplicationrepositoryCmd.Flags().BoolVarP(&_serverlessapplicationrepositoryUpdateApplication, "update-application", "", false, "Update Application")

}
