package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// applicationdiscoveryserviceCmd represents the applicationdiscoveryservice command
var _applicationdiscoveryserviceCmd = &cobra.Command{
	Use:   "applicationdiscoveryservice",
	Short: "AWS applicationdiscoveryservice CLI",
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
		client := applicationdiscoveryservice.NewFromConfig(cfg)
		if _applicationdiscoveryserviceAssociateConfigurationItemsToApplication {
			applicationdiscoveryservice_AssociateConfigurationItemsToApplication(cfg, client)
			return
		}
		if _applicationdiscoveryserviceBatchDeleteAgents {
			applicationdiscoveryservice_BatchDeleteAgents(cfg, client)
			return
		}
		if _applicationdiscoveryserviceBatchDeleteImportData {
			applicationdiscoveryservice_BatchDeleteImportData(cfg, client)
			return
		}
		if _applicationdiscoveryserviceCreateApplication {
			applicationdiscoveryservice_CreateApplication(cfg, client)
			return
		}
		if _applicationdiscoveryserviceCreateTags {
			applicationdiscoveryservice_CreateTags(cfg, client)
			return
		}
		if _applicationdiscoveryserviceDeleteApplications {
			applicationdiscoveryservice_DeleteApplications(cfg, client)
			return
		}
		if _applicationdiscoveryserviceDeleteTags {
			applicationdiscoveryservice_DeleteTags(cfg, client)
			return
		}
		if _applicationdiscoveryserviceDescribeAgents {
			applicationdiscoveryservice_DescribeAgents(cfg, client)
			return
		}
		if _applicationdiscoveryserviceDescribeBatchDeleteConfigurationTask {
			applicationdiscoveryservice_DescribeBatchDeleteConfigurationTask(cfg, client)
			return
		}
		if _applicationdiscoveryserviceDescribeConfigurations {
			applicationdiscoveryservice_DescribeConfigurations(cfg, client)
			return
		}
		if _applicationdiscoveryserviceDescribeContinuousExports {
			applicationdiscoveryservice_DescribeContinuousExports(cfg, client)
			return
		}
		if _applicationdiscoveryserviceDescribeExportConfigurations {
			applicationdiscoveryservice_DescribeExportConfigurations(cfg, client)
			return
		}
		if _applicationdiscoveryserviceDescribeExportTasks {
			applicationdiscoveryservice_DescribeExportTasks(cfg, client)
			return
		}
		if _applicationdiscoveryserviceDescribeImportTasks {
			applicationdiscoveryservice_DescribeImportTasks(cfg, client)
			return
		}
		if _applicationdiscoveryserviceDescribeTags {
			applicationdiscoveryservice_DescribeTags(cfg, client)
			return
		}
		if _applicationdiscoveryserviceDisassociateConfigurationItemsFromApplication {
			applicationdiscoveryservice_DisassociateConfigurationItemsFromApplication(cfg, client)
			return
		}
		if _applicationdiscoveryserviceExportConfigurations {
			applicationdiscoveryservice_ExportConfigurations(cfg, client)
			return
		}
		if _applicationdiscoveryserviceGetDiscoverySummary {
			applicationdiscoveryservice_GetDiscoverySummary(cfg, client)
			return
		}
		if _applicationdiscoveryserviceListConfigurations {
			applicationdiscoveryservice_ListConfigurations(cfg, client)
			return
		}
		if _applicationdiscoveryserviceListServerNeighbors {
			applicationdiscoveryservice_ListServerNeighbors(cfg, client)
			return
		}
		if _applicationdiscoveryserviceStartBatchDeleteConfigurationTask {
			applicationdiscoveryservice_StartBatchDeleteConfigurationTask(cfg, client)
			return
		}
		if _applicationdiscoveryserviceStartContinuousExport {
			applicationdiscoveryservice_StartContinuousExport(cfg, client)
			return
		}
		if _applicationdiscoveryserviceStartDataCollectionByAgentIds {
			applicationdiscoveryservice_StartDataCollectionByAgentIds(cfg, client)
			return
		}
		if _applicationdiscoveryserviceStartExportTask {
			applicationdiscoveryservice_StartExportTask(cfg, client)
			return
		}
		if _applicationdiscoveryserviceStartImportTask {
			applicationdiscoveryservice_StartImportTask(cfg, client)
			return
		}
		if _applicationdiscoveryserviceStopContinuousExport {
			applicationdiscoveryservice_StopContinuousExport(cfg, client)
			return
		}
		if _applicationdiscoveryserviceStopDataCollectionByAgentIds {
			applicationdiscoveryservice_StopDataCollectionByAgentIds(cfg, client)
			return
		}
		if _applicationdiscoveryserviceUpdateApplication {
			applicationdiscoveryservice_UpdateApplication(cfg, client)
			return
		}

	},
}

var (
	_applicationdiscoveryserviceAssociateConfigurationItemsToApplication      bool
	_applicationdiscoveryserviceBatchDeleteAgents                             bool
	_applicationdiscoveryserviceBatchDeleteImportData                         bool
	_applicationdiscoveryserviceCreateApplication                             bool
	_applicationdiscoveryserviceCreateTags                                    bool
	_applicationdiscoveryserviceDeleteApplications                            bool
	_applicationdiscoveryserviceDeleteTags                                    bool
	_applicationdiscoveryserviceDescribeAgents                                bool
	_applicationdiscoveryserviceDescribeBatchDeleteConfigurationTask          bool
	_applicationdiscoveryserviceDescribeConfigurations                        bool
	_applicationdiscoveryserviceDescribeContinuousExports                     bool
	_applicationdiscoveryserviceDescribeExportConfigurations                  bool
	_applicationdiscoveryserviceDescribeExportTasks                           bool
	_applicationdiscoveryserviceDescribeImportTasks                           bool
	_applicationdiscoveryserviceDescribeTags                                  bool
	_applicationdiscoveryserviceDisassociateConfigurationItemsFromApplication bool
	_applicationdiscoveryserviceExportConfigurations                          bool
	_applicationdiscoveryserviceGetDiscoverySummary                           bool
	_applicationdiscoveryserviceListConfigurations                            bool
	_applicationdiscoveryserviceListServerNeighbors                           bool
	_applicationdiscoveryserviceStartBatchDeleteConfigurationTask             bool
	_applicationdiscoveryserviceStartContinuousExport                         bool
	_applicationdiscoveryserviceStartDataCollectionByAgentIds                 bool
	_applicationdiscoveryserviceStartExportTask                               bool
	_applicationdiscoveryserviceStartImportTask                               bool
	_applicationdiscoveryserviceStopContinuousExport                          bool
	_applicationdiscoveryserviceStopDataCollectionByAgentIds                  bool
	_applicationdiscoveryserviceUpdateApplication                             bool

	_applicationdiscoveryserviceAgentIds                   []string
	_applicationdiscoveryserviceApplicationConfigurationId string
	_applicationdiscoveryserviceClientRequestToken         string
	_applicationdiscoveryserviceConfigurationId            string
	_applicationdiscoveryserviceConfigurationIds           []string
	_applicationdiscoveryserviceConfigurationType          string
	_applicationdiscoveryserviceDeleteAgents               string
	_applicationdiscoveryserviceDeleteHistory              string
	_applicationdiscoveryserviceDescription                string
	_applicationdiscoveryserviceEndTime                    string
	_applicationdiscoveryserviceExportDataFormat           string
	_applicationdiscoveryserviceExportId                   string
	_applicationdiscoveryserviceExportIds                  []string
	_applicationdiscoveryserviceFilters                    string
	_applicationdiscoveryserviceImportTaskIds              []string
	_applicationdiscoveryserviceImportUrl                  string
	_applicationdiscoveryserviceMaxResults                 string
	_applicationdiscoveryserviceName                       string
	_applicationdiscoveryserviceNeighborConfigurationIds   []string
	_applicationdiscoveryserviceNextToken                  string
	_applicationdiscoveryserviceOrderBy                    string
	_applicationdiscoveryservicePortInformationNeeded      string
	_applicationdiscoveryservicePreferences                string
	_applicationdiscoveryserviceStartTime                  string
	_applicationdiscoveryserviceTags                       string
	_applicationdiscoveryserviceTaskId                     string
	_applicationdiscoveryserviceWave                       string
)

// Associates one or more configuration items with an application.
func applicationdiscoveryservice_AssociateConfigurationItemsToApplication(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.AssociateConfigurationItemsToApplicationInput{
		// ApplicationConfigurationId: *string, // Required
		// ConfigurationIds: []string, // Required
	}

	if len(_applicationdiscoveryserviceApplicationConfigurationId) > 0 {
		input.ApplicationConfigurationId = aws.String(_applicationdiscoveryserviceApplicationConfigurationId)
	}
	if len(_applicationdiscoveryserviceConfigurationIds) > 0 {
		input.ConfigurationIds = append([]string(nil), _applicationdiscoveryserviceConfigurationIds...)
	}

	if resp, err := client.AssociateConfigurationItemsToApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes one or more agents or collectors as specified by ID. Deleting an agent
// or collector does not delete the previously discovered data. To delete the data
// collected, use StartBatchDeleteConfigurationTask .
func applicationdiscoveryservice_BatchDeleteAgents(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.BatchDeleteAgentsInput{
		// DeleteAgents: []types.DeleteAgent, // Required
	}

	if len(_applicationdiscoveryserviceDeleteAgents) > 0 {
		if err := assignInputField(input, "DeleteAgents", _applicationdiscoveryserviceDeleteAgents); err != nil {
			log.Errorf("invalid --delete-agents: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchDeleteAgents(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes one or more import tasks, each identified by their import ID. Each
// import task has a number of records that can identify servers or applications.
//
// Amazon Web Services Application Discovery Service has built-in matching logic
// that will identify when discovered servers match existing entries that you've
// previously discovered, the information for the already-existing discovered
// server is updated. When you delete an import task that contains records that
// were used to match, the information in those matched records that comes from the
// deleted records will also be deleted.
func applicationdiscoveryservice_BatchDeleteImportData(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.BatchDeleteImportDataInput{
		// ImportTaskIds: []string, // Required
	}

	if len(_applicationdiscoveryserviceImportTaskIds) > 0 {
		input.ImportTaskIds = append([]string(nil), _applicationdiscoveryserviceImportTaskIds...)
	}
	if len(_applicationdiscoveryserviceDeleteHistory) > 0 {
		if err := assignInputField(input, "DeleteHistory", _applicationdiscoveryserviceDeleteHistory); err != nil {
			log.Errorf("invalid --delete-history: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchDeleteImportData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an application with the given name and description.
func applicationdiscoveryservice_CreateApplication(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.CreateApplicationInput{
		// Name: *string, // Required
	}

	if len(_applicationdiscoveryserviceName) > 0 {
		input.Name = aws.String(_applicationdiscoveryserviceName)
	}
	if len(_applicationdiscoveryserviceDescription) > 0 {
		input.Description = aws.String(_applicationdiscoveryserviceDescription)
	}
	if len(_applicationdiscoveryserviceWave) > 0 {
		input.Wave = aws.String(_applicationdiscoveryserviceWave)
	}

	if resp, err := client.CreateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates one or more tags for configuration items. Tags are metadata that help
// you categorize IT assets. This API accepts a list of multiple configuration
// items.
//
// Do not store sensitive information (like personal data) in tags.
func applicationdiscoveryservice_CreateTags(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.CreateTagsInput{
		// ConfigurationIds: []string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_applicationdiscoveryserviceConfigurationIds) > 0 {
		input.ConfigurationIds = append([]string(nil), _applicationdiscoveryserviceConfigurationIds...)
	}
	if len(_applicationdiscoveryserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _applicationdiscoveryserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a list of applications and their associations with configuration items.
func applicationdiscoveryservice_DeleteApplications(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.DeleteApplicationsInput{
		// ConfigurationIds: []string, // Required
	}

	if len(_applicationdiscoveryserviceConfigurationIds) > 0 {
		input.ConfigurationIds = append([]string(nil), _applicationdiscoveryserviceConfigurationIds...)
	}

	if resp, err := client.DeleteApplications(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the association between configuration items and one or more tags. This
// API accepts a list of multiple configuration items.
func applicationdiscoveryservice_DeleteTags(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.DeleteTagsInput{
		// ConfigurationIds: []string, // Required
	}

	if len(_applicationdiscoveryserviceConfigurationIds) > 0 {
		input.ConfigurationIds = append([]string(nil), _applicationdiscoveryserviceConfigurationIds...)
	}
	if len(_applicationdiscoveryserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _applicationdiscoveryserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists agents or collectors as specified by ID or other filters. All
// agents/collectors associated with your user can be listed if you call
// DescribeAgents as is without passing any parameters.
func applicationdiscoveryservice_DescribeAgents(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.DescribeAgentsInput{}

	if len(_applicationdiscoveryserviceAgentIds) > 0 {
		input.AgentIds = append([]string(nil), _applicationdiscoveryserviceAgentIds...)
	}
	if len(_applicationdiscoveryserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _applicationdiscoveryserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_applicationdiscoveryserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationdiscoveryserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationdiscoveryserviceNextToken) > 0 {
		input.NextToken = aws.String(_applicationdiscoveryserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeAgents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*applicationdiscoveryservice.DescribeAgentsOutput
	p := applicationdiscoveryservice.NewDescribeAgentsPaginator(client, input)
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

// Takes a unique deletion task identifier as input and returns metadata about a
// configuration deletion task.
func applicationdiscoveryservice_DescribeBatchDeleteConfigurationTask(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.DescribeBatchDeleteConfigurationTaskInput{
		// TaskId: *string, // Required
	}

	if len(_applicationdiscoveryserviceTaskId) > 0 {
		input.TaskId = aws.String(_applicationdiscoveryserviceTaskId)
	}

	if resp, err := client.DescribeBatchDeleteConfigurationTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves attributes for a list of configuration item IDs.
// All of the supplied IDs must be for the same asset type from one of the
// following:
//
// - server
//
// - application
//
// - process
//
// - connection
//
// Output fields are specific to the asset type specified. For example, the output
// for a server configuration item includes a list of attributes about the server,
// such as host name, operating system, number of network cards, etc.
//
// For a complete list of outputs for each asset type, see [Using the DescribeConfigurations Action] in the Amazon Web
// Services Application Discovery Service User Guide.
//
// [Using the DescribeConfigurations Action]: https://docs.aws.amazon.com/application-discovery/latest/userguide/discovery-api-queries.html#DescribeConfigurations
func applicationdiscoveryservice_DescribeConfigurations(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.DescribeConfigurationsInput{
		// ConfigurationIds: []string, // Required
	}

	if len(_applicationdiscoveryserviceConfigurationIds) > 0 {
		input.ConfigurationIds = append([]string(nil), _applicationdiscoveryserviceConfigurationIds...)
	}

	if resp, err := client.DescribeConfigurations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists exports as specified by ID. All continuous exports associated with your
// user can be listed if you call DescribeContinuousExports as is without passing
// any parameters.
func applicationdiscoveryservice_DescribeContinuousExports(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.DescribeContinuousExportsInput{}

	if len(_applicationdiscoveryserviceExportIds) > 0 {
		input.ExportIds = append([]string(nil), _applicationdiscoveryserviceExportIds...)
	}
	if len(_applicationdiscoveryserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationdiscoveryserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationdiscoveryserviceNextToken) > 0 {
		input.NextToken = aws.String(_applicationdiscoveryserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeContinuousExports(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*applicationdiscoveryservice.DescribeContinuousExportsOutput
	p := applicationdiscoveryservice.NewDescribeContinuousExportsPaginator(client, input)
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

// DescribeExportConfigurations is deprecated. Use [DescribeExportTasks], instead.
// Deprecated: This operation has been deprecated.
//
// [DescribeExportTasks]: https://docs.aws.amazon.com/application-discovery/latest/APIReference/API_DescribeExportTasks.html
func applicationdiscoveryservice_DescribeExportConfigurations(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.DescribeExportConfigurationsInput{}

	if len(_applicationdiscoveryserviceExportIds) > 0 {
		input.ExportIds = append([]string(nil), _applicationdiscoveryserviceExportIds...)
	}
	if len(_applicationdiscoveryserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationdiscoveryserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationdiscoveryserviceNextToken) > 0 {
		input.NextToken = aws.String(_applicationdiscoveryserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeExportConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*applicationdiscoveryservice.DescribeExportConfigurationsOutput
	p := applicationdiscoveryservice.NewDescribeExportConfigurationsPaginator(client, input)
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

// Retrieve status of one or more export tasks. You can retrieve the status of up
// to 100 export tasks.
func applicationdiscoveryservice_DescribeExportTasks(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.DescribeExportTasksInput{}

	if len(_applicationdiscoveryserviceExportIds) > 0 {
		input.ExportIds = append([]string(nil), _applicationdiscoveryserviceExportIds...)
	}
	if len(_applicationdiscoveryserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _applicationdiscoveryserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_applicationdiscoveryserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationdiscoveryserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationdiscoveryserviceNextToken) > 0 {
		input.NextToken = aws.String(_applicationdiscoveryserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeExportTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*applicationdiscoveryservice.DescribeExportTasksOutput
	p := applicationdiscoveryservice.NewDescribeExportTasksPaginator(client, input)
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

// Returns an array of import tasks for your account, including status
// information, times, IDs, the Amazon S3 Object URL for the import file, and more.
func applicationdiscoveryservice_DescribeImportTasks(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.DescribeImportTasksInput{}

	if len(_applicationdiscoveryserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _applicationdiscoveryserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_applicationdiscoveryserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationdiscoveryserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationdiscoveryserviceNextToken) > 0 {
		input.NextToken = aws.String(_applicationdiscoveryserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeImportTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*applicationdiscoveryservice.DescribeImportTasksOutput
	p := applicationdiscoveryservice.NewDescribeImportTasksPaginator(client, input)
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

// Retrieves a list of configuration items that have tags as specified by the
// key-value pairs, name and value, passed to the optional parameter filters .
//
// There are three valid tag filter names:
//
// - tagKey
//
// - tagValue
//
// - configurationId
//
// Also, all configuration items associated with your user that have tags can be
// listed if you call DescribeTags as is without passing any parameters.
func applicationdiscoveryservice_DescribeTags(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.DescribeTagsInput{}

	if len(_applicationdiscoveryserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _applicationdiscoveryserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_applicationdiscoveryserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationdiscoveryserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationdiscoveryserviceNextToken) > 0 {
		input.NextToken = aws.String(_applicationdiscoveryserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*applicationdiscoveryservice.DescribeTagsOutput
	p := applicationdiscoveryservice.NewDescribeTagsPaginator(client, input)
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

// Disassociates one or more configuration items from an application.
func applicationdiscoveryservice_DisassociateConfigurationItemsFromApplication(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationInput{
		// ApplicationConfigurationId: *string, // Required
		// ConfigurationIds: []string, // Required
	}

	if len(_applicationdiscoveryserviceApplicationConfigurationId) > 0 {
		input.ApplicationConfigurationId = aws.String(_applicationdiscoveryserviceApplicationConfigurationId)
	}
	if len(_applicationdiscoveryserviceConfigurationIds) > 0 {
		input.ConfigurationIds = append([]string(nil), _applicationdiscoveryserviceConfigurationIds...)
	}

	if resp, err := client.DisassociateConfigurationItemsFromApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deprecated. Use StartExportTask instead.
// Exports all discovered configuration data to an Amazon S3 bucket or an
// application that enables you to view and evaluate the data. Data includes tags
// and tag associations, processes, connections, servers, and system performance.
// This API returns an export ID that you can query using the
// DescribeExportConfigurations API. The system imposes a limit of two
// configuration exports in six hours.
//
// Deprecated: This operation has been deprecated.
func applicationdiscoveryservice_ExportConfigurations(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.ExportConfigurationsInput{}

	if resp, err := client.ExportConfigurations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a short summary of discovered assets.
// This API operation takes no request parameters and is called as is at the
// command prompt as shown in the example.
func applicationdiscoveryservice_GetDiscoverySummary(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.GetDiscoverySummaryInput{}

	if resp, err := client.GetDiscoverySummary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of configuration items as specified by the value passed to the
// required parameter configurationType . Optional filtering may be applied to
// refine search results.
func applicationdiscoveryservice_ListConfigurations(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.ListConfigurationsInput{
		// ConfigurationType: types.ConfigurationItemType, // Required
	}

	if len(_applicationdiscoveryserviceConfigurationType) > 0 {
		if err := assignInputField(input, "ConfigurationType", _applicationdiscoveryserviceConfigurationType); err != nil {
			log.Errorf("invalid --configuration-type: %s", err.Error())
			return
		}
	}
	if len(_applicationdiscoveryserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _applicationdiscoveryserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_applicationdiscoveryserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationdiscoveryserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationdiscoveryserviceNextToken) > 0 {
		input.NextToken = aws.String(_applicationdiscoveryserviceNextToken)
	}
	if len(_applicationdiscoveryserviceOrderBy) > 0 {
		if err := assignInputField(input, "OrderBy", _applicationdiscoveryserviceOrderBy); err != nil {
			log.Errorf("invalid --order-by: %s", err.Error())
			return
		}
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

	var results []*applicationdiscoveryservice.ListConfigurationsOutput
	p := applicationdiscoveryservice.NewListConfigurationsPaginator(client, input)
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

// Retrieves a list of servers that are one network hop away from a specified
// server.
func applicationdiscoveryservice_ListServerNeighbors(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.ListServerNeighborsInput{
		// ConfigurationId: *string, // Required
	}

	if len(_applicationdiscoveryserviceConfigurationId) > 0 {
		input.ConfigurationId = aws.String(_applicationdiscoveryserviceConfigurationId)
	}
	if len(_applicationdiscoveryserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationdiscoveryserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationdiscoveryserviceNeighborConfigurationIds) > 0 {
		input.NeighborConfigurationIds = append([]string(nil), _applicationdiscoveryserviceNeighborConfigurationIds...)
	}
	if len(_applicationdiscoveryserviceNextToken) > 0 {
		input.NextToken = aws.String(_applicationdiscoveryserviceNextToken)
	}
	if len(_applicationdiscoveryservicePortInformationNeeded) > 0 {
		if err := assignInputField(input, "PortInformationNeeded", _applicationdiscoveryservicePortInformationNeeded); err != nil {
			log.Errorf("invalid --port-information-needed: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListServerNeighbors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Takes a list of configurationId as input and starts an asynchronous deletion
// task to remove the configurationItems. Returns a unique deletion task
// identifier.
func applicationdiscoveryservice_StartBatchDeleteConfigurationTask(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.StartBatchDeleteConfigurationTaskInput{
		// ConfigurationIds: []string, // Required
		// ConfigurationType: types.DeletionConfigurationItemType, // Required
	}

	if len(_applicationdiscoveryserviceConfigurationIds) > 0 {
		input.ConfigurationIds = append([]string(nil), _applicationdiscoveryserviceConfigurationIds...)
	}
	if len(_applicationdiscoveryserviceConfigurationType) > 0 {
		if err := assignInputField(input, "ConfigurationType", _applicationdiscoveryserviceConfigurationType); err != nil {
			log.Errorf("invalid --configuration-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartBatchDeleteConfigurationTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Start the continuous flow of agent's discovered data into Amazon Athena.
func applicationdiscoveryservice_StartContinuousExport(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.StartContinuousExportInput{}

	if resp, err := client.StartContinuousExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Instructs the specified agents to start collecting data.
func applicationdiscoveryservice_StartDataCollectionByAgentIds(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.StartDataCollectionByAgentIdsInput{
		// AgentIds: []string, // Required
	}

	if len(_applicationdiscoveryserviceAgentIds) > 0 {
		input.AgentIds = append([]string(nil), _applicationdiscoveryserviceAgentIds...)
	}

	if resp, err := client.StartDataCollectionByAgentIds(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Begins the export of a discovered data report to an Amazon S3 bucket managed by
// Amazon Web Services.
//
// Exports might provide an estimate of fees and savings based on certain
// information that you provide. Fee estimates do not include any taxes that might
// apply. Your actual fees and savings depend on a variety of factors, including
// your actual usage of Amazon Web Services services, which might vary from the
// estimates provided in this report.
//
// If you do not specify preferences or agentIds in the filter, a summary of all
// servers, applications, tags, and performance is generated. This data is an
// aggregation of all server data collected through on-premises tooling, file
// import, application grouping and applying tags.
//
// If you specify agentIds in a filter, the task exports up to 72 hours of
// detailed data collected by the identified Application Discovery Agent, including
// network, process, and performance details. A time range for exported agent data
// may be set by using startTime and endTime . Export of detailed agent data is
// limited to five concurrently running exports. Export of detailed agent data is
// limited to two exports per day.
//
// If you enable ec2RecommendationsPreferences in preferences , an Amazon EC2
// instance matching the characteristics of each server in Application Discovery
// Service is generated. Changing the attributes of the
// ec2RecommendationsPreferences changes the criteria of the recommendation.
func applicationdiscoveryservice_StartExportTask(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.StartExportTaskInput{}

	if len(_applicationdiscoveryserviceEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _applicationdiscoveryserviceEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_applicationdiscoveryserviceExportDataFormat) > 0 {
		if err := assignInputField(input, "ExportDataFormat", _applicationdiscoveryserviceExportDataFormat); err != nil {
			log.Errorf("invalid --export-data-format: %s", err.Error())
			return
		}
	}
	if len(_applicationdiscoveryserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _applicationdiscoveryserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_applicationdiscoveryservicePreferences) > 0 {
		if err := assignInputField(input, "Preferences", _applicationdiscoveryservicePreferences); err != nil {
			log.Errorf("invalid --preferences: %s", err.Error())
			return
		}
	}
	if len(_applicationdiscoveryserviceStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _applicationdiscoveryserviceStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartExportTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an import task, which allows you to import details of your on-premises
// environment directly into Amazon Web Services Migration Hub without having to
// use the Amazon Web Services Application Discovery Service (Application Discovery
// Service) tools such as the Amazon Web Services Application Discovery Service
// Agentless Collector or Application Discovery Agent. This gives you the option to
// perform migration assessment and planning directly from your imported data,
// including the ability to group your devices as applications and track their
// migration status.
//
// To start an import request, do this:
//
// - Download the specially formatted comma separated value (CSV) import
// template, which you can find here: [https://s3.us-west-2.amazonaws.com/templates-7cffcf56-bd96-4b1c-b45b-a5b42f282e46/import_template.csv].
//
// - Fill out the template with your server and application data.
//
// - Upload your import file to an Amazon S3 bucket, and make a note of it's
// Object URL. Your import file must be in the CSV format.
//
// - Use the console or the StartImportTask command with the Amazon Web Services
// CLI or one of the Amazon Web Services SDKs to import the records from your file.
//
// For more information, including step-by-step procedures, see [Migration Hub Import] in the Amazon Web
// Services Application Discovery Service User Guide.
//
// There are limits to the number of import tasks you can create (and delete) in
// an Amazon Web Services account. For more information, see [Amazon Web Services Application Discovery Service Limits]in the Amazon Web
// Services Application Discovery Service User Guide.
//
// [Amazon Web Services Application Discovery Service Limits]: https://docs.aws.amazon.com/application-discovery/latest/userguide/ads_service_limits.html
// [https://s3.us-west-2.amazonaws.com/templates-7cffcf56-bd96-4b1c-b45b-a5b42f282e46/import_template.csv]: https://s3.us-west-2.amazonaws.com/templates-7cffcf56-bd96-4b1c-b45b-a5b42f282e46/import_template.csv
// [Migration Hub Import]: https://docs.aws.amazon.com/application-discovery/latest/userguide/discovery-import.html
func applicationdiscoveryservice_StartImportTask(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.StartImportTaskInput{
		// ImportUrl: *string, // Required
		// Name: *string, // Required
	}

	if len(_applicationdiscoveryserviceImportUrl) > 0 {
		input.ImportUrl = aws.String(_applicationdiscoveryserviceImportUrl)
	}
	if len(_applicationdiscoveryserviceName) > 0 {
		input.Name = aws.String(_applicationdiscoveryserviceName)
	}
	if len(_applicationdiscoveryserviceClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_applicationdiscoveryserviceClientRequestToken)
	}

	if resp, err := client.StartImportTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stop the continuous flow of agent's discovered data into Amazon Athena.
func applicationdiscoveryservice_StopContinuousExport(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.StopContinuousExportInput{
		// ExportId: *string, // Required
	}

	if len(_applicationdiscoveryserviceExportId) > 0 {
		input.ExportId = aws.String(_applicationdiscoveryserviceExportId)
	}

	if resp, err := client.StopContinuousExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Instructs the specified agents to stop collecting data.
func applicationdiscoveryservice_StopDataCollectionByAgentIds(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.StopDataCollectionByAgentIdsInput{
		// AgentIds: []string, // Required
	}

	if len(_applicationdiscoveryserviceAgentIds) > 0 {
		input.AgentIds = append([]string(nil), _applicationdiscoveryserviceAgentIds...)
	}

	if resp, err := client.StopDataCollectionByAgentIds(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates metadata about an application.
func applicationdiscoveryservice_UpdateApplication(cfg aws.Config, client *applicationdiscoveryservice.Client) {
	input := &applicationdiscoveryservice.UpdateApplicationInput{
		// ConfigurationId: *string, // Required
	}

	if len(_applicationdiscoveryserviceConfigurationId) > 0 {
		input.ConfigurationId = aws.String(_applicationdiscoveryserviceConfigurationId)
	}
	if len(_applicationdiscoveryserviceDescription) > 0 {
		input.Description = aws.String(_applicationdiscoveryserviceDescription)
	}
	if len(_applicationdiscoveryserviceName) > 0 {
		input.Name = aws.String(_applicationdiscoveryserviceName)
	}
	if len(_applicationdiscoveryserviceWave) > 0 {
		input.Wave = aws.String(_applicationdiscoveryserviceWave)
	}

	if resp, err := client.UpdateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_applicationdiscoveryserviceCmd)
	_applicationdiscoveryserviceCmd.Flags().SortFlags = false

	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_applicationdiscoveryserviceCmd.Flags().StringSliceVarP(&_applicationdiscoveryserviceAgentIds, "agent-ids", "", nil, "Agent Ids")
	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_applicationdiscoveryserviceApplicationConfigurationId, "application-configuration-id", "", "", "Application Configuration ID")
	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_applicationdiscoveryserviceClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_applicationdiscoveryserviceConfigurationId, "configuration-id", "", "", "Configuration ID")
	_applicationdiscoveryserviceCmd.Flags().StringSliceVarP(&_applicationdiscoveryserviceConfigurationIds, "configuration-ids", "", nil, "Configuration Ids")
	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_applicationdiscoveryserviceConfigurationType, "configuration-type", "", "", "Configuration Type")
	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_applicationdiscoveryserviceDeleteAgents, "delete-agents", "", "", "Delete Agents")
	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_applicationdiscoveryserviceDeleteHistory, "delete-history", "", "", "Delete History")
	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_applicationdiscoveryserviceDescription, "description", "", "", "Description")
	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_applicationdiscoveryserviceEndTime, "end-time", "", "", "End Time")
	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_applicationdiscoveryserviceExportDataFormat, "export-data-format", "", "", "Export Data Format")
	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_applicationdiscoveryserviceExportId, "export-id", "", "", "Export ID")
	_applicationdiscoveryserviceCmd.Flags().StringSliceVarP(&_applicationdiscoveryserviceExportIds, "export-ids", "", nil, "Export Ids")
	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_applicationdiscoveryserviceFilters, "filters", "", "", "Filters")
	_applicationdiscoveryserviceCmd.Flags().StringSliceVarP(&_applicationdiscoveryserviceImportTaskIds, "import-task-ids", "", nil, "Import Task Ids")
	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_applicationdiscoveryserviceImportUrl, "import-url", "", "", "Import URL")
	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_applicationdiscoveryserviceMaxResults, "max-results", "", "", "Max Results")
	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_applicationdiscoveryserviceName, "name", "", "", "Name")
	_applicationdiscoveryserviceCmd.Flags().StringSliceVarP(&_applicationdiscoveryserviceNeighborConfigurationIds, "neighbor-configuration-ids", "", nil, "Neighbor Configuration Ids")
	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_applicationdiscoveryserviceNextToken, "next-token", "", "", "Next Token")
	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_applicationdiscoveryserviceOrderBy, "order-by", "", "", "Order By")
	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_applicationdiscoveryservicePortInformationNeeded, "port-information-needed", "", "", "Port Information Needed")
	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_applicationdiscoveryservicePreferences, "preferences", "", "", "Preferences")
	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_applicationdiscoveryserviceStartTime, "start-time", "", "", "Start Time")
	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_applicationdiscoveryserviceTags, "tags", "", "", "Tags")
	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_applicationdiscoveryserviceTaskId, "task-id", "", "", "Task ID")
	_applicationdiscoveryserviceCmd.Flags().StringVarP(&_applicationdiscoveryserviceWave, "wave", "", "", "Wave")

	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceAssociateConfigurationItemsToApplication, "associate-configuration-items-to-application", "", false, "Associate Configuration Items To Application")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceBatchDeleteAgents, "batch-delete-agents", "", false, "Batch Delete Agents")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceBatchDeleteImportData, "batch-delete-import-data", "", false, "Batch Delete Import Data")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceCreateApplication, "create-application", "", false, "Create Application")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceCreateTags, "create-tags", "", false, "Create Tags")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceDeleteApplications, "delete-applications", "", false, "Delete Applications")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceDeleteTags, "delete-tags", "", false, "Delete Tags")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceDescribeAgents, "describe-agents", "", false, "Describe Agents")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceDescribeBatchDeleteConfigurationTask, "describe-batch-delete-configuration-task", "", false, "Describe Batch Delete Configuration Task")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceDescribeConfigurations, "describe-configurations", "", false, "Describe Configurations")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceDescribeContinuousExports, "describe-continuous-exports", "", false, "Describe Continuous Exports")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceDescribeExportConfigurations, "describe-export-configurations", "", false, "Describe Export Configurations")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceDescribeExportTasks, "describe-export-tasks", "", false, "Describe Export Tasks")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceDescribeImportTasks, "describe-import-tasks", "", false, "Describe Import Tasks")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceDescribeTags, "describe-tags", "", false, "Describe Tags")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceDisassociateConfigurationItemsFromApplication, "disassociate-configuration-items-from-application", "", false, "Disassociate Configuration Items From Application")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceExportConfigurations, "export-configurations", "", false, "Export Configurations")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceGetDiscoverySummary, "get-discovery-summary", "", false, "Get Discovery Summary")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceListConfigurations, "list-configurations", "", false, "List Configurations")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceListServerNeighbors, "list-server-neighbors", "", false, "List Server Neighbors")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceStartBatchDeleteConfigurationTask, "start-batch-delete-configuration-task", "", false, "Start Batch Delete Configuration Task")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceStartContinuousExport, "start-continuous-export", "", false, "Start Continuous Export")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceStartDataCollectionByAgentIds, "start-data-collection-by-agent-ids", "", false, "Start Data Collection By Agent Ids")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceStartExportTask, "start-export-task", "", false, "Start Export Task")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceStartImportTask, "start-import-task", "", false, "Start Import Task")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceStopContinuousExport, "stop-continuous-export", "", false, "Stop Continuous Export")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceStopDataCollectionByAgentIds, "stop-data-collection-by-agent-ids", "", false, "Stop Data Collection By Agent Ids")
	_applicationdiscoveryserviceCmd.Flags().BoolVarP(&_applicationdiscoveryserviceUpdateApplication, "update-application", "", false, "Update Application")

}
