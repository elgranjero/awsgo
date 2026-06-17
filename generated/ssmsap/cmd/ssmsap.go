package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssmsap"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ssmsapCmd represents the ssmsap command
var _ssmsapCmd = &cobra.Command{
	Use:   "ssmsap",
	Short: "AWS ssmsap CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := ssmsap.NewFromConfig(cfg)
		if _ssmsapDeleteResourcePermission {
			ssmsap_DeleteResourcePermission(cfg, client)
			return
		}
		if _ssmsapDeregisterApplication {
			ssmsap_DeregisterApplication(cfg, client)
			return
		}
		if _ssmsapGetApplication {
			ssmsap_GetApplication(cfg, client)
			return
		}
		if _ssmsapGetComponent {
			ssmsap_GetComponent(cfg, client)
			return
		}
		if _ssmsapGetConfigurationCheckOperation {
			ssmsap_GetConfigurationCheckOperation(cfg, client)
			return
		}
		if _ssmsapGetDatabase {
			ssmsap_GetDatabase(cfg, client)
			return
		}
		if _ssmsapGetOperation {
			ssmsap_GetOperation(cfg, client)
			return
		}
		if _ssmsapGetResourcePermission {
			ssmsap_GetResourcePermission(cfg, client)
			return
		}
		if _ssmsapListApplications {
			ssmsap_ListApplications(cfg, client)
			return
		}
		if _ssmsapListComponents {
			ssmsap_ListComponents(cfg, client)
			return
		}
		if _ssmsapListConfigurationCheckDefinitions {
			ssmsap_ListConfigurationCheckDefinitions(cfg, client)
			return
		}
		if _ssmsapListConfigurationCheckOperations {
			ssmsap_ListConfigurationCheckOperations(cfg, client)
			return
		}
		if _ssmsapListDatabases {
			ssmsap_ListDatabases(cfg, client)
			return
		}
		if _ssmsapListOperationEvents {
			ssmsap_ListOperationEvents(cfg, client)
			return
		}
		if _ssmsapListOperations {
			ssmsap_ListOperations(cfg, client)
			return
		}
		if _ssmsapListSubCheckResults {
			ssmsap_ListSubCheckResults(cfg, client)
			return
		}
		if _ssmsapListSubCheckRuleResults {
			ssmsap_ListSubCheckRuleResults(cfg, client)
			return
		}
		if _ssmsapListTagsForResource {
			ssmsap_ListTagsForResource(cfg, client)
			return
		}
		if _ssmsapPutResourcePermission {
			ssmsap_PutResourcePermission(cfg, client)
			return
		}
		if _ssmsapRegisterApplication {
			ssmsap_RegisterApplication(cfg, client)
			return
		}
		if _ssmsapStartApplication {
			ssmsap_StartApplication(cfg, client)
			return
		}
		if _ssmsapStartApplicationRefresh {
			ssmsap_StartApplicationRefresh(cfg, client)
			return
		}
		if _ssmsapStartConfigurationChecks {
			ssmsap_StartConfigurationChecks(cfg, client)
			return
		}
		if _ssmsapStopApplication {
			ssmsap_StopApplication(cfg, client)
			return
		}
		if _ssmsapTagResource {
			ssmsap_TagResource(cfg, client)
			return
		}
		if _ssmsapUntagResource {
			ssmsap_UntagResource(cfg, client)
			return
		}
		if _ssmsapUpdateApplicationSettings {
			ssmsap_UpdateApplicationSettings(cfg, client)
			return
		}

	},
}

var (
	_ssmsapDeleteResourcePermission          bool
	_ssmsapDeregisterApplication             bool
	_ssmsapGetApplication                    bool
	_ssmsapGetComponent                      bool
	_ssmsapGetConfigurationCheckOperation    bool
	_ssmsapGetDatabase                       bool
	_ssmsapGetOperation                      bool
	_ssmsapGetResourcePermission             bool
	_ssmsapListApplications                  bool
	_ssmsapListComponents                    bool
	_ssmsapListConfigurationCheckDefinitions bool
	_ssmsapListConfigurationCheckOperations  bool
	_ssmsapListDatabases                     bool
	_ssmsapListOperationEvents               bool
	_ssmsapListOperations                    bool
	_ssmsapListSubCheckResults               bool
	_ssmsapListSubCheckRuleResults           bool
	_ssmsapListTagsForResource               bool
	_ssmsapPutResourcePermission             bool
	_ssmsapRegisterApplication               bool
	_ssmsapStartApplication                  bool
	_ssmsapStartApplicationRefresh           bool
	_ssmsapStartConfigurationChecks          bool
	_ssmsapStopApplication                   bool
	_ssmsapTagResource                       bool
	_ssmsapUntagResource                     bool
	_ssmsapUpdateApplicationSettings         bool

	_ssmsapActionType                 string
	_ssmsapAppRegistryArn             string
	_ssmsapApplicationArn             string
	_ssmsapApplicationId              string
	_ssmsapApplicationType            string
	_ssmsapBackint                    string
	_ssmsapComponentId                string
	_ssmsapComponentsInfo             string
	_ssmsapConfigurationCheckIds      string
	_ssmsapCredentials                string
	_ssmsapCredentialsToAddOrUpdate   string
	_ssmsapCredentialsToRemove        string
	_ssmsapDatabaseArn                string
	_ssmsapDatabaseId                 string
	_ssmsapFilters                    string
	_ssmsapIncludeEc2InstanceShutdown string
	_ssmsapInstances                  []string
	_ssmsapListMode                   string
	_ssmsapMaxResults                 string
	_ssmsapNextToken                  string
	_ssmsapOperationId                string
	_ssmsapResourceArn                string
	_ssmsapSapInstanceNumber          string
	_ssmsapSid                        string
	_ssmsapSourceResourceArn          string
	_ssmsapStopConnectedEntity        string
	_ssmsapSubCheckResultId           string
	_ssmsapTagKeys                    []string
	_ssmsapTags                       string
)

// Removes permissions associated with the target database.
func ssmsap_DeleteResourcePermission(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.DeleteResourcePermissionInput{
		// ResourceArn: *string, // Required
	}

	if len(_ssmsapResourceArn) > 0 {
		input.ResourceArn = aws.String(_ssmsapResourceArn)
	}
	if len(_ssmsapActionType) > 0 {
		if err := assignInputField(input, "ActionType", _ssmsapActionType); err != nil {
			log.Errorf("invalid --action-type: %s", err.Error())
			return
		}
	}
	if len(_ssmsapSourceResourceArn) > 0 {
		input.SourceResourceArn = aws.String(_ssmsapSourceResourceArn)
	}

	if resp, err := client.DeleteResourcePermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregister an SAP application with AWS Systems Manager for SAP. This action
// does not aﬀect the existing setup of your SAP workloads on Amazon EC2.
func ssmsap_DeregisterApplication(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.DeregisterApplicationInput{
		// ApplicationId: *string, // Required
	}

	if len(_ssmsapApplicationId) > 0 {
		input.ApplicationId = aws.String(_ssmsapApplicationId)
	}

	if resp, err := client.DeregisterApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an application registered with AWS Systems Manager for SAP. It also
// returns the components of the application.
func ssmsap_GetApplication(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.GetApplicationInput{}

	if len(_ssmsapAppRegistryArn) > 0 {
		input.AppRegistryArn = aws.String(_ssmsapAppRegistryArn)
	}
	if len(_ssmsapApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_ssmsapApplicationArn)
	}
	if len(_ssmsapApplicationId) > 0 {
		input.ApplicationId = aws.String(_ssmsapApplicationId)
	}

	if resp, err := client.GetApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the component of an application registered with AWS Systems Manager for
// SAP.
func ssmsap_GetComponent(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.GetComponentInput{
		// ApplicationId: *string, // Required
		// ComponentId: *string, // Required
	}

	if len(_ssmsapApplicationId) > 0 {
		input.ApplicationId = aws.String(_ssmsapApplicationId)
	}
	if len(_ssmsapComponentId) > 0 {
		input.ComponentId = aws.String(_ssmsapComponentId)
	}

	if resp, err := client.GetComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the details of a configuration check operation by specifying the operation
// ID.
func ssmsap_GetConfigurationCheckOperation(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.GetConfigurationCheckOperationInput{
		// OperationId: *string, // Required
	}

	if len(_ssmsapOperationId) > 0 {
		input.OperationId = aws.String(_ssmsapOperationId)
	}

	if resp, err := client.GetConfigurationCheckOperation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the SAP HANA database of an application registered with AWS Systems
// Manager for SAP.
func ssmsap_GetDatabase(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.GetDatabaseInput{}

	if len(_ssmsapApplicationId) > 0 {
		input.ApplicationId = aws.String(_ssmsapApplicationId)
	}
	if len(_ssmsapComponentId) > 0 {
		input.ComponentId = aws.String(_ssmsapComponentId)
	}
	if len(_ssmsapDatabaseArn) > 0 {
		input.DatabaseArn = aws.String(_ssmsapDatabaseArn)
	}
	if len(_ssmsapDatabaseId) > 0 {
		input.DatabaseId = aws.String(_ssmsapDatabaseId)
	}

	if resp, err := client.GetDatabase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the details of an operation by specifying the operation ID.
func ssmsap_GetOperation(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.GetOperationInput{
		// OperationId: *string, // Required
	}

	if len(_ssmsapOperationId) > 0 {
		input.OperationId = aws.String(_ssmsapOperationId)
	}

	if resp, err := client.GetOperation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets permissions associated with the target database.
func ssmsap_GetResourcePermission(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.GetResourcePermissionInput{
		// ResourceArn: *string, // Required
	}

	if len(_ssmsapResourceArn) > 0 {
		input.ResourceArn = aws.String(_ssmsapResourceArn)
	}
	if len(_ssmsapActionType) > 0 {
		if err := assignInputField(input, "ActionType", _ssmsapActionType); err != nil {
			log.Errorf("invalid --action-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetResourcePermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the applications registered with AWS Systems Manager for SAP.
func ssmsap_ListApplications(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.ListApplicationsInput{}

	if len(_ssmsapFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmsapFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmsapMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmsapMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmsapNextToken) > 0 {
		input.NextToken = aws.String(_ssmsapNextToken)
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

	var results []*ssmsap.ListApplicationsOutput
	p := ssmsap.NewListApplicationsPaginator(client, input)
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

// Lists all the components registered with AWS Systems Manager for SAP.
func ssmsap_ListComponents(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.ListComponentsInput{}

	if len(_ssmsapApplicationId) > 0 {
		input.ApplicationId = aws.String(_ssmsapApplicationId)
	}
	if len(_ssmsapMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmsapMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmsapNextToken) > 0 {
		input.NextToken = aws.String(_ssmsapNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListComponents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmsap.ListComponentsOutput
	p := ssmsap.NewListComponentsPaginator(client, input)
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

// Lists all configuration check types supported by AWS Systems Manager for SAP.
func ssmsap_ListConfigurationCheckDefinitions(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.ListConfigurationCheckDefinitionsInput{}

	if len(_ssmsapMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmsapMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmsapNextToken) > 0 {
		input.NextToken = aws.String(_ssmsapNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConfigurationCheckDefinitions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmsap.ListConfigurationCheckDefinitionsOutput
	p := ssmsap.NewListConfigurationCheckDefinitionsPaginator(client, input)
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

// Lists the configuration check operations performed by AWS Systems Manager for
// SAP.
func ssmsap_ListConfigurationCheckOperations(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.ListConfigurationCheckOperationsInput{
		// ApplicationId: *string, // Required
	}

	if len(_ssmsapApplicationId) > 0 {
		input.ApplicationId = aws.String(_ssmsapApplicationId)
	}
	if len(_ssmsapFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmsapFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmsapListMode) > 0 {
		if err := assignInputField(input, "ListMode", _ssmsapListMode); err != nil {
			log.Errorf("invalid --list-mode: %s", err.Error())
			return
		}
	}
	if len(_ssmsapMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmsapMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmsapNextToken) > 0 {
		input.NextToken = aws.String(_ssmsapNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConfigurationCheckOperations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmsap.ListConfigurationCheckOperationsOutput
	p := ssmsap.NewListConfigurationCheckOperationsPaginator(client, input)
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

// Lists the SAP HANA databases of an application registered with AWS Systems
// Manager for SAP.
func ssmsap_ListDatabases(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.ListDatabasesInput{}

	if len(_ssmsapApplicationId) > 0 {
		input.ApplicationId = aws.String(_ssmsapApplicationId)
	}
	if len(_ssmsapComponentId) > 0 {
		input.ComponentId = aws.String(_ssmsapComponentId)
	}
	if len(_ssmsapMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmsapMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmsapNextToken) > 0 {
		input.NextToken = aws.String(_ssmsapNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDatabases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmsap.ListDatabasesOutput
	p := ssmsap.NewListDatabasesPaginator(client, input)
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

// Returns a list of operations events.
// Available parameters include OperationID , as well as optional parameters
// MaxResults , NextToken , and Filters .
func ssmsap_ListOperationEvents(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.ListOperationEventsInput{
		// OperationId: *string, // Required
	}

	if len(_ssmsapOperationId) > 0 {
		input.OperationId = aws.String(_ssmsapOperationId)
	}
	if len(_ssmsapFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmsapFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmsapMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmsapMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmsapNextToken) > 0 {
		input.NextToken = aws.String(_ssmsapNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOperationEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmsap.ListOperationEventsOutput
	p := ssmsap.NewListOperationEventsPaginator(client, input)
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

// Lists the operations performed by AWS Systems Manager for SAP.
func ssmsap_ListOperations(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.ListOperationsInput{
		// ApplicationId: *string, // Required
	}

	if len(_ssmsapApplicationId) > 0 {
		input.ApplicationId = aws.String(_ssmsapApplicationId)
	}
	if len(_ssmsapFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmsapFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmsapMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmsapMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmsapNextToken) > 0 {
		input.NextToken = aws.String(_ssmsapNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOperations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmsap.ListOperationsOutput
	p := ssmsap.NewListOperationsPaginator(client, input)
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

// Lists the sub-check results of a specified configuration check operation.
func ssmsap_ListSubCheckResults(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.ListSubCheckResultsInput{
		// OperationId: *string, // Required
	}

	if len(_ssmsapOperationId) > 0 {
		input.OperationId = aws.String(_ssmsapOperationId)
	}
	if len(_ssmsapMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmsapMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmsapNextToken) > 0 {
		input.NextToken = aws.String(_ssmsapNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSubCheckResults(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmsap.ListSubCheckResultsOutput
	p := ssmsap.NewListSubCheckResultsPaginator(client, input)
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

// Lists the rules of a specified sub-check belonging to a configuration check
// operation.
func ssmsap_ListSubCheckRuleResults(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.ListSubCheckRuleResultsInput{
		// SubCheckResultId: *string, // Required
	}

	if len(_ssmsapSubCheckResultId) > 0 {
		input.SubCheckResultId = aws.String(_ssmsapSubCheckResultId)
	}
	if len(_ssmsapMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmsapMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmsapNextToken) > 0 {
		input.NextToken = aws.String(_ssmsapNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSubCheckRuleResults(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmsap.ListSubCheckRuleResultsOutput
	p := ssmsap.NewListSubCheckRuleResultsPaginator(client, input)
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

// Lists all tags on an SAP HANA application and/or database registered with AWS
// Systems Manager for SAP.
func ssmsap_ListTagsForResource(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_ssmsapResourceArn) > 0 {
		input.ResourceArn = aws.String(_ssmsapResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds permissions to the target database.
func ssmsap_PutResourcePermission(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.PutResourcePermissionInput{
		// ActionType: types.PermissionActionType, // Required
		// ResourceArn: *string, // Required
		// SourceResourceArn: *string, // Required
	}

	if len(_ssmsapActionType) > 0 {
		if err := assignInputField(input, "ActionType", _ssmsapActionType); err != nil {
			log.Errorf("invalid --action-type: %s", err.Error())
			return
		}
	}
	if len(_ssmsapResourceArn) > 0 {
		input.ResourceArn = aws.String(_ssmsapResourceArn)
	}
	if len(_ssmsapSourceResourceArn) > 0 {
		input.SourceResourceArn = aws.String(_ssmsapSourceResourceArn)
	}

	if resp, err := client.PutResourcePermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Register an SAP application with AWS Systems Manager for SAP. You must meet the
// following requirements before registering.
//
// The SAP application you want to register with AWS Systems Manager for SAP is
// running on Amazon EC2.
//
// AWS Systems Manager Agent must be setup on an Amazon EC2 instance along with
// the required IAM permissions.
//
// Amazon EC2 instance(s) must have access to the secrets created in AWS Secrets
// Manager to manage SAP applications and components.
func ssmsap_RegisterApplication(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.RegisterApplicationInput{
		// ApplicationId: *string, // Required
		// ApplicationType: types.ApplicationType, // Required
		// Instances: []string, // Required
	}

	if len(_ssmsapApplicationId) > 0 {
		input.ApplicationId = aws.String(_ssmsapApplicationId)
	}
	if len(_ssmsapApplicationType) > 0 {
		if err := assignInputField(input, "ApplicationType", _ssmsapApplicationType); err != nil {
			log.Errorf("invalid --application-type: %s", err.Error())
			return
		}
	}
	if len(_ssmsapInstances) > 0 {
		input.Instances = append([]string(nil), _ssmsapInstances...)
	}
	if len(_ssmsapComponentsInfo) > 0 {
		if err := assignInputField(input, "ComponentsInfo", _ssmsapComponentsInfo); err != nil {
			log.Errorf("invalid --components-info: %s", err.Error())
			return
		}
	}
	if len(_ssmsapCredentials) > 0 {
		if err := assignInputField(input, "Credentials", _ssmsapCredentials); err != nil {
			log.Errorf("invalid --credentials: %s", err.Error())
			return
		}
	}
	if len(_ssmsapDatabaseArn) > 0 {
		input.DatabaseArn = aws.String(_ssmsapDatabaseArn)
	}
	if len(_ssmsapSapInstanceNumber) > 0 {
		input.SapInstanceNumber = aws.String(_ssmsapSapInstanceNumber)
	}
	if len(_ssmsapSid) > 0 {
		input.Sid = aws.String(_ssmsapSid)
	}
	if len(_ssmsapTags) > 0 {
		if err := assignInputField(input, "Tags", _ssmsapTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Request is an operation which starts an application.
// Parameter ApplicationId is required.
func ssmsap_StartApplication(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.StartApplicationInput{
		// ApplicationId: *string, // Required
	}

	if len(_ssmsapApplicationId) > 0 {
		input.ApplicationId = aws.String(_ssmsapApplicationId)
	}

	if resp, err := client.StartApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Refreshes a registered application.
func ssmsap_StartApplicationRefresh(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.StartApplicationRefreshInput{
		// ApplicationId: *string, // Required
	}

	if len(_ssmsapApplicationId) > 0 {
		input.ApplicationId = aws.String(_ssmsapApplicationId)
	}

	if resp, err := client.StartApplicationRefresh(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates configuration check operations against a specified application.
func ssmsap_StartConfigurationChecks(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.StartConfigurationChecksInput{
		// ApplicationId: *string, // Required
	}

	if len(_ssmsapApplicationId) > 0 {
		input.ApplicationId = aws.String(_ssmsapApplicationId)
	}
	if len(_ssmsapConfigurationCheckIds) > 0 {
		if err := assignInputField(input, "ConfigurationCheckIds", _ssmsapConfigurationCheckIds); err != nil {
			log.Errorf("invalid --configuration-check-ids: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartConfigurationChecks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Request is an operation to stop an application.
// Parameter ApplicationId is required. Parameters StopConnectedEntity and
// IncludeEc2InstanceShutdown are optional.
func ssmsap_StopApplication(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.StopApplicationInput{
		// ApplicationId: *string, // Required
	}

	if len(_ssmsapApplicationId) > 0 {
		input.ApplicationId = aws.String(_ssmsapApplicationId)
	}
	if len(_ssmsapIncludeEc2InstanceShutdown) > 0 {
		if err := assignInputField(input, "IncludeEc2InstanceShutdown", _ssmsapIncludeEc2InstanceShutdown); err != nil {
			log.Errorf("invalid --include-ec2-instance-shutdown: %s", err.Error())
			return
		}
	}
	if len(_ssmsapStopConnectedEntity) > 0 {
		if err := assignInputField(input, "StopConnectedEntity", _ssmsapStopConnectedEntity); err != nil {
			log.Errorf("invalid --stop-connected-entity: %s", err.Error())
			return
		}
	}

	if resp, err := client.StopApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates tag for a resource by specifying the ARN.
func ssmsap_TagResource(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_ssmsapResourceArn) > 0 {
		input.ResourceArn = aws.String(_ssmsapResourceArn)
	}
	if len(_ssmsapTags) > 0 {
		if err := assignInputField(input, "Tags", _ssmsapTags); err != nil {
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

// Delete the tags for a resource.
func ssmsap_UntagResource(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_ssmsapResourceArn) > 0 {
		input.ResourceArn = aws.String(_ssmsapResourceArn)
	}
	if len(_ssmsapTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _ssmsapTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the settings of an application registered with AWS Systems Manager for
// SAP.
func ssmsap_UpdateApplicationSettings(cfg aws.Config, client *ssmsap.Client) {
	input := &ssmsap.UpdateApplicationSettingsInput{
		// ApplicationId: *string, // Required
	}

	if len(_ssmsapApplicationId) > 0 {
		input.ApplicationId = aws.String(_ssmsapApplicationId)
	}
	if len(_ssmsapBackint) > 0 {
		if err := assignInputField(input, "Backint", _ssmsapBackint); err != nil {
			log.Errorf("invalid --backint: %s", err.Error())
			return
		}
	}
	if len(_ssmsapCredentialsToAddOrUpdate) > 0 {
		if err := assignInputField(input, "CredentialsToAddOrUpdate", _ssmsapCredentialsToAddOrUpdate); err != nil {
			log.Errorf("invalid --credentials-to-add-or-update: %s", err.Error())
			return
		}
	}
	if len(_ssmsapCredentialsToRemove) > 0 {
		if err := assignInputField(input, "CredentialsToRemove", _ssmsapCredentialsToRemove); err != nil {
			log.Errorf("invalid --credentials-to-remove: %s", err.Error())
			return
		}
	}
	if len(_ssmsapDatabaseArn) > 0 {
		input.DatabaseArn = aws.String(_ssmsapDatabaseArn)
	}

	if resp, err := client.UpdateApplicationSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_ssmsapCmd)
	_ssmsapCmd.Flags().SortFlags = false

	_ssmsapCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_ssmsapCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_ssmsapCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_ssmsapCmd.Flags().StringVarP(&_ssmsapActionType, "action-type", "", "", "Action Type")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapAppRegistryArn, "app-registry-arn", "", "", "App Registry ARN")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapApplicationArn, "application-arn", "", "", "Application ARN")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapApplicationId, "application-id", "", "", "Application ID")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapApplicationType, "application-type", "", "", "Application Type")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapBackint, "backint", "", "", "Backint")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapComponentId, "component-id", "", "", "Component ID")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapComponentsInfo, "components-info", "", "", "Components Info")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapConfigurationCheckIds, "configuration-check-ids", "", "", "Configuration Check Ids")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapCredentials, "credentials", "", "", "Credentials")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapCredentialsToAddOrUpdate, "credentials-to-add-or-update", "", "", "Credentials To Add Or Update")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapCredentialsToRemove, "credentials-to-remove", "", "", "Credentials To Remove")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapDatabaseArn, "database-arn", "", "", "Database ARN")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapDatabaseId, "database-id", "", "", "Database ID")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapFilters, "filters", "", "", "Filters")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapIncludeEc2InstanceShutdown, "include-ec2-instance-shutdown", "", "", "Include EC2 Instance Shutdown")
	_ssmsapCmd.Flags().StringSliceVarP(&_ssmsapInstances, "instances", "", nil, "Instances")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapListMode, "list-mode", "", "", "List Mode")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapMaxResults, "max-results", "", "", "Max Results")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapNextToken, "next-token", "", "", "Next Token")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapOperationId, "operation-id", "", "", "Operation ID")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapResourceArn, "resource-arn", "", "", "Resource ARN")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapSapInstanceNumber, "sap-instance-number", "", "", "Sap Instance Number")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapSid, "sid", "", "", "Sid")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapSourceResourceArn, "source-resource-arn", "", "", "Source Resource ARN")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapStopConnectedEntity, "stop-connected-entity", "", "", "Stop Connected Entity")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapSubCheckResultId, "sub-check-result-id", "", "", "Sub Check Result ID")
	_ssmsapCmd.Flags().StringSliceVarP(&_ssmsapTagKeys, "tag-keys", "", nil, "Tag Keys")
	_ssmsapCmd.Flags().StringVarP(&_ssmsapTags, "tags", "", "", "Tags")

	_ssmsapCmd.Flags().BoolVarP(&_ssmsapDeleteResourcePermission, "delete-resource-permission", "", false, "Delete Resource Permission")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapDeregisterApplication, "deregister-application", "", false, "Deregister Application")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapGetApplication, "get-application", "", false, "Get Application")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapGetComponent, "get-component", "", false, "Get Component")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapGetConfigurationCheckOperation, "get-configuration-check-operation", "", false, "Get Configuration Check Operation")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapGetDatabase, "get-database", "", false, "Get Database")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapGetOperation, "get-operation", "", false, "Get Operation")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapGetResourcePermission, "get-resource-permission", "", false, "Get Resource Permission")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapListApplications, "list-applications", "", false, "List Applications")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapListComponents, "list-components", "", false, "List Components")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapListConfigurationCheckDefinitions, "list-configuration-check-definitions", "", false, "List Configuration Check Definitions")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapListConfigurationCheckOperations, "list-configuration-check-operations", "", false, "List Configuration Check Operations")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapListDatabases, "list-databases", "", false, "List Databases")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapListOperationEvents, "list-operation-events", "", false, "List Operation Events")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapListOperations, "list-operations", "", false, "List Operations")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapListSubCheckResults, "list-sub-check-results", "", false, "List Sub Check Results")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapListSubCheckRuleResults, "list-sub-check-rule-results", "", false, "List Sub Check Rule Results")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapPutResourcePermission, "put-resource-permission", "", false, "Put Resource Permission")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapRegisterApplication, "register-application", "", false, "Register Application")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapStartApplication, "start-application", "", false, "Start Application")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapStartApplicationRefresh, "start-application-refresh", "", false, "Start Application Refresh")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapStartConfigurationChecks, "start-configuration-checks", "", false, "Start Configuration Checks")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapStopApplication, "stop-application", "", false, "Stop Application")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapTagResource, "tag-resource", "", false, "Tag Resource")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapUntagResource, "untag-resource", "", false, "Untag Resource")
	_ssmsapCmd.Flags().BoolVarP(&_ssmsapUpdateApplicationSettings, "update-application-settings", "", false, "Update Application Settings")

}
