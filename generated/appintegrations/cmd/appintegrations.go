package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appintegrations"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// appintegrationsCmd represents the appintegrations command
var _appintegrationsCmd = &cobra.Command{
	Use:   "appintegrations",
	Short: "AWS appintegrations CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := appintegrations.NewFromConfig(cfg)
		if _appintegrationsCreateApplication {
			appintegrations_CreateApplication(cfg, client)
			return
		}
		if _appintegrationsCreateDataIntegration {
			appintegrations_CreateDataIntegration(cfg, client)
			return
		}
		if _appintegrationsCreateDataIntegrationAssociation {
			appintegrations_CreateDataIntegrationAssociation(cfg, client)
			return
		}
		if _appintegrationsCreateEventIntegration {
			appintegrations_CreateEventIntegration(cfg, client)
			return
		}
		if _appintegrationsDeleteApplication {
			appintegrations_DeleteApplication(cfg, client)
			return
		}
		if _appintegrationsDeleteDataIntegration {
			appintegrations_DeleteDataIntegration(cfg, client)
			return
		}
		if _appintegrationsDeleteEventIntegration {
			appintegrations_DeleteEventIntegration(cfg, client)
			return
		}
		if _appintegrationsGetApplication {
			appintegrations_GetApplication(cfg, client)
			return
		}
		if _appintegrationsGetDataIntegration {
			appintegrations_GetDataIntegration(cfg, client)
			return
		}
		if _appintegrationsGetEventIntegration {
			appintegrations_GetEventIntegration(cfg, client)
			return
		}
		if _appintegrationsListApplicationAssociations {
			appintegrations_ListApplicationAssociations(cfg, client)
			return
		}
		if _appintegrationsListApplications {
			appintegrations_ListApplications(cfg, client)
			return
		}
		if _appintegrationsListDataIntegrationAssociations {
			appintegrations_ListDataIntegrationAssociations(cfg, client)
			return
		}
		if _appintegrationsListDataIntegrations {
			appintegrations_ListDataIntegrations(cfg, client)
			return
		}
		if _appintegrationsListEventIntegrationAssociations {
			appintegrations_ListEventIntegrationAssociations(cfg, client)
			return
		}
		if _appintegrationsListEventIntegrations {
			appintegrations_ListEventIntegrations(cfg, client)
			return
		}
		if _appintegrationsListTagsForResource {
			appintegrations_ListTagsForResource(cfg, client)
			return
		}
		if _appintegrationsTagResource {
			appintegrations_TagResource(cfg, client)
			return
		}
		if _appintegrationsUntagResource {
			appintegrations_UntagResource(cfg, client)
			return
		}
		if _appintegrationsUpdateApplication {
			appintegrations_UpdateApplication(cfg, client)
			return
		}
		if _appintegrationsUpdateDataIntegration {
			appintegrations_UpdateDataIntegration(cfg, client)
			return
		}
		if _appintegrationsUpdateDataIntegrationAssociation {
			appintegrations_UpdateDataIntegrationAssociation(cfg, client)
			return
		}
		if _appintegrationsUpdateEventIntegration {
			appintegrations_UpdateEventIntegration(cfg, client)
			return
		}

	},
}

var (
	_appintegrationsCreateApplication                bool
	_appintegrationsCreateDataIntegration            bool
	_appintegrationsCreateDataIntegrationAssociation bool
	_appintegrationsCreateEventIntegration           bool
	_appintegrationsDeleteApplication                bool
	_appintegrationsDeleteDataIntegration            bool
	_appintegrationsDeleteEventIntegration           bool
	_appintegrationsGetApplication                   bool
	_appintegrationsGetDataIntegration               bool
	_appintegrationsGetEventIntegration              bool
	_appintegrationsListApplicationAssociations      bool
	_appintegrationsListApplications                 bool
	_appintegrationsListDataIntegrationAssociations  bool
	_appintegrationsListDataIntegrations             bool
	_appintegrationsListEventIntegrationAssociations bool
	_appintegrationsListEventIntegrations            bool
	_appintegrationsListTagsForResource              bool
	_appintegrationsTagResource                      bool
	_appintegrationsUntagResource                    bool
	_appintegrationsUpdateApplication                bool
	_appintegrationsUpdateDataIntegration            bool
	_appintegrationsUpdateDataIntegrationAssociation bool
	_appintegrationsUpdateEventIntegration           bool

	_appintegrationsApplicationConfig                    string
	_appintegrationsApplicationId                        string
	_appintegrationsApplicationSourceConfig              string
	_appintegrationsApplicationType                      string
	_appintegrationsArn                                  string
	_appintegrationsClientAssociationMetadata            string
	_appintegrationsClientId                             string
	_appintegrationsClientToken                          string
	_appintegrationsDataIntegrationAssociationIdentifier string
	_appintegrationsDataIntegrationIdentifier            string
	_appintegrationsDescription                          string
	_appintegrationsDestinationURI                       string
	_appintegrationsEventBridgeBus                       string
	_appintegrationsEventFilter                          string
	_appintegrationsEventIntegrationName                 string
	_appintegrationsExecutionConfiguration               string
	_appintegrationsFileConfiguration                    string
	_appintegrationsIdentifier                           string
	_appintegrationsIframeConfig                         string
	_appintegrationsInitializationTimeout                string
	_appintegrationsIsService                            string
	_appintegrationsKmsKey                               string
	_appintegrationsMaxResults                           string
	_appintegrationsName                                 string
	_appintegrationsNamespace                            string
	_appintegrationsNextToken                            string
	_appintegrationsObjectConfiguration                  string
	_appintegrationsPermissions                          []string
	_appintegrationsPublications                         string
	_appintegrationsResourceArn                          string
	_appintegrationsScheduleConfig                       string
	_appintegrationsSourceURI                            string
	_appintegrationsSubscriptions                        string
	_appintegrationsTagKeys                              []string
	_appintegrationsTags                                 string
)

// Creates and persists an Application resource.
func appintegrations_CreateApplication(cfg aws.Config, client *appintegrations.Client) {
	input := &appintegrations.CreateApplicationInput{
		// ApplicationSourceConfig: *types.ApplicationSourceConfig, // Required
		// Name: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_appintegrationsApplicationSourceConfig) > 0 {
		if err := assignInputField(input, "ApplicationSourceConfig", _appintegrationsApplicationSourceConfig); err != nil {
			log.Errorf("invalid --application-source-config: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsName) > 0 {
		input.Name = aws.String(_appintegrationsName)
	}
	if len(_appintegrationsNamespace) > 0 {
		input.Namespace = aws.String(_appintegrationsNamespace)
	}
	if len(_appintegrationsApplicationConfig) > 0 {
		if err := assignInputField(input, "ApplicationConfig", _appintegrationsApplicationConfig); err != nil {
			log.Errorf("invalid --application-config: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsApplicationType) > 0 {
		if err := assignInputField(input, "ApplicationType", _appintegrationsApplicationType); err != nil {
			log.Errorf("invalid --application-type: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsClientToken) > 0 {
		input.ClientToken = aws.String(_appintegrationsClientToken)
	}
	if len(_appintegrationsDescription) > 0 {
		input.Description = aws.String(_appintegrationsDescription)
	}
	if len(_appintegrationsIframeConfig) > 0 {
		if err := assignInputField(input, "IframeConfig", _appintegrationsIframeConfig); err != nil {
			log.Errorf("invalid --iframe-config: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsInitializationTimeout) > 0 {
		if err := assignInputField(input, "InitializationTimeout", _appintegrationsInitializationTimeout); err != nil {
			log.Errorf("invalid --initialization-timeout: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsIsService) > 0 {
		if err := assignInputField(input, "IsService", _appintegrationsIsService); err != nil {
			log.Errorf("invalid --is-service: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsPermissions) > 0 {
		input.Permissions = append([]string(nil), _appintegrationsPermissions...)
	}
	if len(_appintegrationsPublications) > 0 {
		if err := assignInputField(input, "Publications", _appintegrationsPublications); err != nil {
			log.Errorf("invalid --publications: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsSubscriptions) > 0 {
		if err := assignInputField(input, "Subscriptions", _appintegrationsSubscriptions); err != nil {
			log.Errorf("invalid --subscriptions: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsTags) > 0 {
		if err := assignInputField(input, "Tags", _appintegrationsTags); err != nil {
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

// Creates and persists a DataIntegration resource.
// You cannot create a DataIntegration association for a DataIntegration that has
// been previously associated. Use a different DataIntegration, or recreate the
// DataIntegration using the CreateDataIntegration API.
func appintegrations_CreateDataIntegration(cfg aws.Config, client *appintegrations.Client) {
	input := &appintegrations.CreateDataIntegrationInput{
		// KmsKey: *string, // Required
		// Name: *string, // Required
	}

	if len(_appintegrationsKmsKey) > 0 {
		input.KmsKey = aws.String(_appintegrationsKmsKey)
	}
	if len(_appintegrationsName) > 0 {
		input.Name = aws.String(_appintegrationsName)
	}
	if len(_appintegrationsClientToken) > 0 {
		input.ClientToken = aws.String(_appintegrationsClientToken)
	}
	if len(_appintegrationsDescription) > 0 {
		input.Description = aws.String(_appintegrationsDescription)
	}
	if len(_appintegrationsFileConfiguration) > 0 {
		if err := assignInputField(input, "FileConfiguration", _appintegrationsFileConfiguration); err != nil {
			log.Errorf("invalid --file-configuration: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsObjectConfiguration) > 0 {
		if err := assignInputField(input, "ObjectConfiguration", _appintegrationsObjectConfiguration); err != nil {
			log.Errorf("invalid --object-configuration: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsScheduleConfig) > 0 {
		if err := assignInputField(input, "ScheduleConfig", _appintegrationsScheduleConfig); err != nil {
			log.Errorf("invalid --schedule-config: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsSourceURI) > 0 {
		input.SourceURI = aws.String(_appintegrationsSourceURI)
	}
	if len(_appintegrationsTags) > 0 {
		if err := assignInputField(input, "Tags", _appintegrationsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates and persists a DataIntegrationAssociation resource.
func appintegrations_CreateDataIntegrationAssociation(cfg aws.Config, client *appintegrations.Client) {
	input := &appintegrations.CreateDataIntegrationAssociationInput{
		// DataIntegrationIdentifier: *string, // Required
	}

	if len(_appintegrationsDataIntegrationIdentifier) > 0 {
		input.DataIntegrationIdentifier = aws.String(_appintegrationsDataIntegrationIdentifier)
	}
	if len(_appintegrationsClientAssociationMetadata) > 0 {
		if err := assignInputField(input, "ClientAssociationMetadata", _appintegrationsClientAssociationMetadata); err != nil {
			log.Errorf("invalid --client-association-metadata: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsClientId) > 0 {
		input.ClientId = aws.String(_appintegrationsClientId)
	}
	if len(_appintegrationsClientToken) > 0 {
		input.ClientToken = aws.String(_appintegrationsClientToken)
	}
	if len(_appintegrationsDestinationURI) > 0 {
		input.DestinationURI = aws.String(_appintegrationsDestinationURI)
	}
	if len(_appintegrationsExecutionConfiguration) > 0 {
		if err := assignInputField(input, "ExecutionConfiguration", _appintegrationsExecutionConfiguration); err != nil {
			log.Errorf("invalid --execution-configuration: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsObjectConfiguration) > 0 {
		if err := assignInputField(input, "ObjectConfiguration", _appintegrationsObjectConfiguration); err != nil {
			log.Errorf("invalid --object-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataIntegrationAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an EventIntegration, given a specified name, description, and a
// reference to an Amazon EventBridge bus in your account and a partner event
// source that pushes events to that bus. No objects are created in the your
// account, only metadata that is persisted on the EventIntegration control plane.
func appintegrations_CreateEventIntegration(cfg aws.Config, client *appintegrations.Client) {
	input := &appintegrations.CreateEventIntegrationInput{
		// EventBridgeBus: *string, // Required
		// EventFilter: *types.EventFilter, // Required
		// Name: *string, // Required
	}

	if len(_appintegrationsEventBridgeBus) > 0 {
		input.EventBridgeBus = aws.String(_appintegrationsEventBridgeBus)
	}
	if len(_appintegrationsEventFilter) > 0 {
		if err := assignInputField(input, "EventFilter", _appintegrationsEventFilter); err != nil {
			log.Errorf("invalid --event-filter: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsName) > 0 {
		input.Name = aws.String(_appintegrationsName)
	}
	if len(_appintegrationsClientToken) > 0 {
		input.ClientToken = aws.String(_appintegrationsClientToken)
	}
	if len(_appintegrationsDescription) > 0 {
		input.Description = aws.String(_appintegrationsDescription)
	}
	if len(_appintegrationsTags) > 0 {
		if err := assignInputField(input, "Tags", _appintegrationsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEventIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the Application. Only Applications that don't have any Application
// Associations can be deleted.
func appintegrations_DeleteApplication(cfg aws.Config, client *appintegrations.Client) {
	input := &appintegrations.DeleteApplicationInput{
		// Arn: *string, // Required
	}

	if len(_appintegrationsArn) > 0 {
		input.Arn = aws.String(_appintegrationsArn)
	}

	if resp, err := client.DeleteApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the DataIntegration. Only DataIntegrations that don't have any
// DataIntegrationAssociations can be deleted. Deleting a DataIntegration also
// deletes the underlying Amazon AppFlow flow and service linked role.
//
// You cannot create a DataIntegration association for a DataIntegration that has
// been previously associated. Use a different DataIntegration, or recreate the
// DataIntegration using the [CreateDataIntegration]API.
//
// [CreateDataIntegration]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_CreateDataIntegration.html
func appintegrations_DeleteDataIntegration(cfg aws.Config, client *appintegrations.Client) {
	input := &appintegrations.DeleteDataIntegrationInput{
		// DataIntegrationIdentifier: *string, // Required
	}

	if len(_appintegrationsDataIntegrationIdentifier) > 0 {
		input.DataIntegrationIdentifier = aws.String(_appintegrationsDataIntegrationIdentifier)
	}

	if resp, err := client.DeleteDataIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified existing event integration. If the event integration is
// associated with clients, the request is rejected.
func appintegrations_DeleteEventIntegration(cfg aws.Config, client *appintegrations.Client) {
	input := &appintegrations.DeleteEventIntegrationInput{
		// Name: *string, // Required
	}

	if len(_appintegrationsName) > 0 {
		input.Name = aws.String(_appintegrationsName)
	}

	if resp, err := client.DeleteEventIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get an Application resource.
func appintegrations_GetApplication(cfg aws.Config, client *appintegrations.Client) {
	input := &appintegrations.GetApplicationInput{
		// Arn: *string, // Required
	}

	if len(_appintegrationsArn) > 0 {
		input.Arn = aws.String(_appintegrationsArn)
	}

	if resp, err := client.GetApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the DataIntegration.
// You cannot create a DataIntegration association for a DataIntegration that has
// been previously associated. Use a different DataIntegration, or recreate the
// DataIntegration using the [CreateDataIntegration]API.
//
// [CreateDataIntegration]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_CreateDataIntegration.html
func appintegrations_GetDataIntegration(cfg aws.Config, client *appintegrations.Client) {
	input := &appintegrations.GetDataIntegrationInput{
		// Identifier: *string, // Required
	}

	if len(_appintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_appintegrationsIdentifier)
	}

	if resp, err := client.GetDataIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the event integration.
func appintegrations_GetEventIntegration(cfg aws.Config, client *appintegrations.Client) {
	input := &appintegrations.GetEventIntegrationInput{
		// Name: *string, // Required
	}

	if len(_appintegrationsName) > 0 {
		input.Name = aws.String(_appintegrationsName)
	}

	if resp, err := client.GetEventIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a paginated list of application associations for an application.
func appintegrations_ListApplicationAssociations(cfg aws.Config, client *appintegrations.Client) {
	input := &appintegrations.ListApplicationAssociationsInput{
		// ApplicationId: *string, // Required
	}

	if len(_appintegrationsApplicationId) > 0 {
		input.ApplicationId = aws.String(_appintegrationsApplicationId)
	}
	if len(_appintegrationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appintegrationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsNextToken) > 0 {
		input.NextToken = aws.String(_appintegrationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApplicationAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appintegrations.ListApplicationAssociationsOutput
	p := appintegrations.NewListApplicationAssociationsPaginator(client, input)
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

// Lists applications in the account.
func appintegrations_ListApplications(cfg aws.Config, client *appintegrations.Client) {
	input := &appintegrations.ListApplicationsInput{}

	if len(_appintegrationsApplicationType) > 0 {
		if err := assignInputField(input, "ApplicationType", _appintegrationsApplicationType); err != nil {
			log.Errorf("invalid --application-type: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appintegrationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsNextToken) > 0 {
		input.NextToken = aws.String(_appintegrationsNextToken)
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

	var results []*appintegrations.ListApplicationsOutput
	p := appintegrations.NewListApplicationsPaginator(client, input)
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

// Returns a paginated list of DataIntegration associations in the account.
// You cannot create a DataIntegration association for a DataIntegration that has
// been previously associated. Use a different DataIntegration, or recreate the
// DataIntegration using the [CreateDataIntegration]API.
//
// [CreateDataIntegration]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_CreateDataIntegration.html
func appintegrations_ListDataIntegrationAssociations(cfg aws.Config, client *appintegrations.Client) {
	input := &appintegrations.ListDataIntegrationAssociationsInput{
		// DataIntegrationIdentifier: *string, // Required
	}

	if len(_appintegrationsDataIntegrationIdentifier) > 0 {
		input.DataIntegrationIdentifier = aws.String(_appintegrationsDataIntegrationIdentifier)
	}
	if len(_appintegrationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appintegrationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsNextToken) > 0 {
		input.NextToken = aws.String(_appintegrationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataIntegrationAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appintegrations.ListDataIntegrationAssociationsOutput
	p := appintegrations.NewListDataIntegrationAssociationsPaginator(client, input)
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

// Returns a paginated list of DataIntegrations in the account.
// You cannot create a DataIntegration association for a DataIntegration that has
// been previously associated. Use a different DataIntegration, or recreate the
// DataIntegration using the [CreateDataIntegration]API.
//
// [CreateDataIntegration]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_CreateDataIntegration.html
func appintegrations_ListDataIntegrations(cfg aws.Config, client *appintegrations.Client) {
	input := &appintegrations.ListDataIntegrationsInput{}

	if len(_appintegrationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appintegrationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsNextToken) > 0 {
		input.NextToken = aws.String(_appintegrationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataIntegrations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appintegrations.ListDataIntegrationsOutput
	p := appintegrations.NewListDataIntegrationsPaginator(client, input)
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

// Returns a paginated list of event integration associations in the account.
func appintegrations_ListEventIntegrationAssociations(cfg aws.Config, client *appintegrations.Client) {
	input := &appintegrations.ListEventIntegrationAssociationsInput{
		// EventIntegrationName: *string, // Required
	}

	if len(_appintegrationsEventIntegrationName) > 0 {
		input.EventIntegrationName = aws.String(_appintegrationsEventIntegrationName)
	}
	if len(_appintegrationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appintegrationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsNextToken) > 0 {
		input.NextToken = aws.String(_appintegrationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEventIntegrationAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appintegrations.ListEventIntegrationAssociationsOutput
	p := appintegrations.NewListEventIntegrationAssociationsPaginator(client, input)
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

// Returns a paginated list of event integrations in the account.
func appintegrations_ListEventIntegrations(cfg aws.Config, client *appintegrations.Client) {
	input := &appintegrations.ListEventIntegrationsInput{}

	if len(_appintegrationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appintegrationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsNextToken) > 0 {
		input.NextToken = aws.String(_appintegrationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEventIntegrations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appintegrations.ListEventIntegrationsOutput
	p := appintegrations.NewListEventIntegrationsPaginator(client, input)
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
func appintegrations_ListTagsForResource(cfg aws.Config, client *appintegrations.Client) {
	input := &appintegrations.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_appintegrationsResourceArn) > 0 {
		input.ResourceArn = aws.String(_appintegrationsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the specified tags to the specified resource.
func appintegrations_TagResource(cfg aws.Config, client *appintegrations.Client) {
	input := &appintegrations.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_appintegrationsResourceArn) > 0 {
		input.ResourceArn = aws.String(_appintegrationsResourceArn)
	}
	if len(_appintegrationsTags) > 0 {
		if err := assignInputField(input, "Tags", _appintegrationsTags); err != nil {
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
func appintegrations_UntagResource(cfg aws.Config, client *appintegrations.Client) {
	input := &appintegrations.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_appintegrationsResourceArn) > 0 {
		input.ResourceArn = aws.String(_appintegrationsResourceArn)
	}
	if len(_appintegrationsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _appintegrationsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates and persists an Application resource.
func appintegrations_UpdateApplication(cfg aws.Config, client *appintegrations.Client) {
	input := &appintegrations.UpdateApplicationInput{
		// Arn: *string, // Required
	}

	if len(_appintegrationsArn) > 0 {
		input.Arn = aws.String(_appintegrationsArn)
	}
	if len(_appintegrationsApplicationConfig) > 0 {
		if err := assignInputField(input, "ApplicationConfig", _appintegrationsApplicationConfig); err != nil {
			log.Errorf("invalid --application-config: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsApplicationSourceConfig) > 0 {
		if err := assignInputField(input, "ApplicationSourceConfig", _appintegrationsApplicationSourceConfig); err != nil {
			log.Errorf("invalid --application-source-config: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsApplicationType) > 0 {
		if err := assignInputField(input, "ApplicationType", _appintegrationsApplicationType); err != nil {
			log.Errorf("invalid --application-type: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsDescription) > 0 {
		input.Description = aws.String(_appintegrationsDescription)
	}
	if len(_appintegrationsIframeConfig) > 0 {
		if err := assignInputField(input, "IframeConfig", _appintegrationsIframeConfig); err != nil {
			log.Errorf("invalid --iframe-config: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsInitializationTimeout) > 0 {
		if err := assignInputField(input, "InitializationTimeout", _appintegrationsInitializationTimeout); err != nil {
			log.Errorf("invalid --initialization-timeout: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsIsService) > 0 {
		if err := assignInputField(input, "IsService", _appintegrationsIsService); err != nil {
			log.Errorf("invalid --is-service: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsName) > 0 {
		input.Name = aws.String(_appintegrationsName)
	}
	if len(_appintegrationsPermissions) > 0 {
		input.Permissions = append([]string(nil), _appintegrationsPermissions...)
	}
	if len(_appintegrationsPublications) > 0 {
		if err := assignInputField(input, "Publications", _appintegrationsPublications); err != nil {
			log.Errorf("invalid --publications: %s", err.Error())
			return
		}
	}
	if len(_appintegrationsSubscriptions) > 0 {
		if err := assignInputField(input, "Subscriptions", _appintegrationsSubscriptions); err != nil {
			log.Errorf("invalid --subscriptions: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the description of a DataIntegration.
// You cannot create a DataIntegration association for a DataIntegration that has
// been previously associated. Use a different DataIntegration, or recreate the
// DataIntegration using the [CreateDataIntegration]API.
//
// [CreateDataIntegration]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_CreateDataIntegration.html
func appintegrations_UpdateDataIntegration(cfg aws.Config, client *appintegrations.Client) {
	input := &appintegrations.UpdateDataIntegrationInput{
		// Identifier: *string, // Required
	}

	if len(_appintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_appintegrationsIdentifier)
	}
	if len(_appintegrationsDescription) > 0 {
		input.Description = aws.String(_appintegrationsDescription)
	}
	if len(_appintegrationsName) > 0 {
		input.Name = aws.String(_appintegrationsName)
	}

	if resp, err := client.UpdateDataIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates and persists a DataIntegrationAssociation resource.
// Updating a DataIntegrationAssociation with ExecutionConfiguration will rerun
// the on-demand job.
func appintegrations_UpdateDataIntegrationAssociation(cfg aws.Config, client *appintegrations.Client) {
	input := &appintegrations.UpdateDataIntegrationAssociationInput{
		// DataIntegrationAssociationIdentifier: *string, // Required
		// DataIntegrationIdentifier: *string, // Required
		// ExecutionConfiguration: *types.ExecutionConfiguration, // Required
	}

	if len(_appintegrationsDataIntegrationAssociationIdentifier) > 0 {
		input.DataIntegrationAssociationIdentifier = aws.String(_appintegrationsDataIntegrationAssociationIdentifier)
	}
	if len(_appintegrationsDataIntegrationIdentifier) > 0 {
		input.DataIntegrationIdentifier = aws.String(_appintegrationsDataIntegrationIdentifier)
	}
	if len(_appintegrationsExecutionConfiguration) > 0 {
		if err := assignInputField(input, "ExecutionConfiguration", _appintegrationsExecutionConfiguration); err != nil {
			log.Errorf("invalid --execution-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDataIntegrationAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the description of an event integration.
func appintegrations_UpdateEventIntegration(cfg aws.Config, client *appintegrations.Client) {
	input := &appintegrations.UpdateEventIntegrationInput{
		// Name: *string, // Required
	}

	if len(_appintegrationsName) > 0 {
		input.Name = aws.String(_appintegrationsName)
	}
	if len(_appintegrationsDescription) > 0 {
		input.Description = aws.String(_appintegrationsDescription)
	}

	if resp, err := client.UpdateEventIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_appintegrationsCmd)
	_appintegrationsCmd.Flags().SortFlags = false

	_appintegrationsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_appintegrationsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_appintegrationsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsApplicationConfig, "application-config", "", "", "Application Config")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsApplicationId, "application-id", "", "", "Application ID")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsApplicationSourceConfig, "application-source-config", "", "", "Application Source Config")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsApplicationType, "application-type", "", "", "Application Type")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsArn, "arn", "", "", "ARN")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsClientAssociationMetadata, "client-association-metadata", "", "", "Client Association Metadata")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsClientId, "client-id", "", "", "Client ID")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsClientToken, "client-token", "", "", "Client Token")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsDataIntegrationAssociationIdentifier, "data-integration-association-identifier", "", "", "Data Integration Association Identifier")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsDataIntegrationIdentifier, "data-integration-identifier", "", "", "Data Integration Identifier")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsDescription, "description", "", "", "Description")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsDestinationURI, "destination-uri", "", "", "Destination URI")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsEventBridgeBus, "event-bridge-bus", "", "", "Event Bridge Bus")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsEventFilter, "event-filter", "", "", "Event Filter")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsEventIntegrationName, "event-integration-name", "", "", "Event Integration Name")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsExecutionConfiguration, "execution-configuration", "", "", "Execution Configuration")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsFileConfiguration, "file-configuration", "", "", "File Configuration")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsIdentifier, "identifier", "", "", "Identifier")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsIframeConfig, "iframe-config", "", "", "Iframe Config")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsInitializationTimeout, "initialization-timeout", "", "", "Initialization Timeout")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsIsService, "is-service", "", "", "Is Service")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsKmsKey, "kms-key", "", "", "KMS Key")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsMaxResults, "max-results", "", "", "Max Results")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsName, "name", "", "", "Name")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsNamespace, "namespace", "", "", "Namespace")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsNextToken, "next-token", "", "", "Next Token")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsObjectConfiguration, "object-configuration", "", "", "Object Configuration")
	_appintegrationsCmd.Flags().StringSliceVarP(&_appintegrationsPermissions, "permissions", "", nil, "Permissions")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsPublications, "publications", "", "", "Publications")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsResourceArn, "resource-arn", "", "", "Resource ARN")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsScheduleConfig, "schedule-config", "", "", "Schedule Config")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsSourceURI, "source-uri", "", "", "Source URI")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsSubscriptions, "subscriptions", "", "", "Subscriptions")
	_appintegrationsCmd.Flags().StringSliceVarP(&_appintegrationsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_appintegrationsCmd.Flags().StringVarP(&_appintegrationsTags, "tags", "", "", "Tags")

	_appintegrationsCmd.Flags().BoolVarP(&_appintegrationsCreateApplication, "create-application", "", false, "Create Application")
	_appintegrationsCmd.Flags().BoolVarP(&_appintegrationsCreateDataIntegration, "create-data-integration", "", false, "Create Data Integration")
	_appintegrationsCmd.Flags().BoolVarP(&_appintegrationsCreateDataIntegrationAssociation, "create-data-integration-association", "", false, "Create Data Integration Association")
	_appintegrationsCmd.Flags().BoolVarP(&_appintegrationsCreateEventIntegration, "create-event-integration", "", false, "Create Event Integration")
	_appintegrationsCmd.Flags().BoolVarP(&_appintegrationsDeleteApplication, "delete-application", "", false, "Delete Application")
	_appintegrationsCmd.Flags().BoolVarP(&_appintegrationsDeleteDataIntegration, "delete-data-integration", "", false, "Delete Data Integration")
	_appintegrationsCmd.Flags().BoolVarP(&_appintegrationsDeleteEventIntegration, "delete-event-integration", "", false, "Delete Event Integration")
	_appintegrationsCmd.Flags().BoolVarP(&_appintegrationsGetApplication, "get-application", "", false, "Get Application")
	_appintegrationsCmd.Flags().BoolVarP(&_appintegrationsGetDataIntegration, "get-data-integration", "", false, "Get Data Integration")
	_appintegrationsCmd.Flags().BoolVarP(&_appintegrationsGetEventIntegration, "get-event-integration", "", false, "Get Event Integration")
	_appintegrationsCmd.Flags().BoolVarP(&_appintegrationsListApplicationAssociations, "list-application-associations", "", false, "List Application Associations")
	_appintegrationsCmd.Flags().BoolVarP(&_appintegrationsListApplications, "list-applications", "", false, "List Applications")
	_appintegrationsCmd.Flags().BoolVarP(&_appintegrationsListDataIntegrationAssociations, "list-data-integration-associations", "", false, "List Data Integration Associations")
	_appintegrationsCmd.Flags().BoolVarP(&_appintegrationsListDataIntegrations, "list-data-integrations", "", false, "List Data Integrations")
	_appintegrationsCmd.Flags().BoolVarP(&_appintegrationsListEventIntegrationAssociations, "list-event-integration-associations", "", false, "List Event Integration Associations")
	_appintegrationsCmd.Flags().BoolVarP(&_appintegrationsListEventIntegrations, "list-event-integrations", "", false, "List Event Integrations")
	_appintegrationsCmd.Flags().BoolVarP(&_appintegrationsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_appintegrationsCmd.Flags().BoolVarP(&_appintegrationsTagResource, "tag-resource", "", false, "Tag Resource")
	_appintegrationsCmd.Flags().BoolVarP(&_appintegrationsUntagResource, "untag-resource", "", false, "Untag Resource")
	_appintegrationsCmd.Flags().BoolVarP(&_appintegrationsUpdateApplication, "update-application", "", false, "Update Application")
	_appintegrationsCmd.Flags().BoolVarP(&_appintegrationsUpdateDataIntegration, "update-data-integration", "", false, "Update Data Integration")
	_appintegrationsCmd.Flags().BoolVarP(&_appintegrationsUpdateDataIntegrationAssociation, "update-data-integration-association", "", false, "Update Data Integration Association")
	_appintegrationsCmd.Flags().BoolVarP(&_appintegrationsUpdateEventIntegration, "update-event-integration", "", false, "Update Event Integration")

}
