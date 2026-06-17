package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appflow"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// appflowCmd represents the appflow command
var _appflowCmd = &cobra.Command{
	Use:   "appflow",
	Short: "AWS appflow CLI",
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
		client := appflow.NewFromConfig(cfg)
		if _appflowCancelFlowExecutions {
			appflow_CancelFlowExecutions(cfg, client)
			return
		}
		if _appflowCreateConnectorProfile {
			appflow_CreateConnectorProfile(cfg, client)
			return
		}
		if _appflowCreateFlow {
			appflow_CreateFlow(cfg, client)
			return
		}
		if _appflowDeleteConnectorProfile {
			appflow_DeleteConnectorProfile(cfg, client)
			return
		}
		if _appflowDeleteFlow {
			appflow_DeleteFlow(cfg, client)
			return
		}
		if _appflowDescribeConnector {
			appflow_DescribeConnector(cfg, client)
			return
		}
		if _appflowDescribeConnectorEntity {
			appflow_DescribeConnectorEntity(cfg, client)
			return
		}
		if _appflowDescribeConnectorProfiles {
			appflow_DescribeConnectorProfiles(cfg, client)
			return
		}
		if _appflowDescribeConnectors {
			appflow_DescribeConnectors(cfg, client)
			return
		}
		if _appflowDescribeFlow {
			appflow_DescribeFlow(cfg, client)
			return
		}
		if _appflowDescribeFlowExecutionRecords {
			appflow_DescribeFlowExecutionRecords(cfg, client)
			return
		}
		if _appflowListConnectorEntities {
			appflow_ListConnectorEntities(cfg, client)
			return
		}
		if _appflowListConnectors {
			appflow_ListConnectors(cfg, client)
			return
		}
		if _appflowListFlows {
			appflow_ListFlows(cfg, client)
			return
		}
		if _appflowListTagsForResource {
			appflow_ListTagsForResource(cfg, client)
			return
		}
		if _appflowRegisterConnector {
			appflow_RegisterConnector(cfg, client)
			return
		}
		if _appflowResetConnectorMetadataCache {
			appflow_ResetConnectorMetadataCache(cfg, client)
			return
		}
		if _appflowStartFlow {
			appflow_StartFlow(cfg, client)
			return
		}
		if _appflowStopFlow {
			appflow_StopFlow(cfg, client)
			return
		}
		if _appflowTagResource {
			appflow_TagResource(cfg, client)
			return
		}
		if _appflowUnregisterConnector {
			appflow_UnregisterConnector(cfg, client)
			return
		}
		if _appflowUntagResource {
			appflow_UntagResource(cfg, client)
			return
		}
		if _appflowUpdateConnectorProfile {
			appflow_UpdateConnectorProfile(cfg, client)
			return
		}
		if _appflowUpdateConnectorRegistration {
			appflow_UpdateConnectorRegistration(cfg, client)
			return
		}
		if _appflowUpdateFlow {
			appflow_UpdateFlow(cfg, client)
			return
		}

	},
}

var (
	_appflowCancelFlowExecutions         bool
	_appflowCreateConnectorProfile       bool
	_appflowCreateFlow                   bool
	_appflowDeleteConnectorProfile       bool
	_appflowDeleteFlow                   bool
	_appflowDescribeConnector            bool
	_appflowDescribeConnectorEntity      bool
	_appflowDescribeConnectorProfiles    bool
	_appflowDescribeConnectors           bool
	_appflowDescribeFlow                 bool
	_appflowDescribeFlowExecutionRecords bool
	_appflowListConnectorEntities        bool
	_appflowListConnectors               bool
	_appflowListFlows                    bool
	_appflowListTagsForResource          bool
	_appflowRegisterConnector            bool
	_appflowResetConnectorMetadataCache  bool
	_appflowStartFlow                    bool
	_appflowStopFlow                     bool
	_appflowTagResource                  bool
	_appflowUnregisterConnector          bool
	_appflowUntagResource                bool
	_appflowUpdateConnectorProfile       bool
	_appflowUpdateConnectorRegistration  bool
	_appflowUpdateFlow                   bool

	_appflowApiVersion                  string
	_appflowClientToken                 string
	_appflowConnectionMode              string
	_appflowConnectorEntityName         string
	_appflowConnectorLabel              string
	_appflowConnectorProfileConfig      string
	_appflowConnectorProfileName        string
	_appflowConnectorProfileNames       []string
	_appflowConnectorProvisioningConfig string
	_appflowConnectorProvisioningType   string
	_appflowConnectorType               string
	_appflowConnectorTypes              string
	_appflowDescription                 string
	_appflowDestinationFlowConfigList   string
	_appflowEntitiesPath                string
	_appflowExecutionIds                []string
	_appflowFlowName                    string
	_appflowForceDelete                 string
	_appflowKmsArn                      string
	_appflowMaxResults                  string
	_appflowMetadataCatalogConfig       string
	_appflowNextToken                   string
	_appflowResourceArn                 string
	_appflowSourceFlowConfig            string
	_appflowTagKeys                     []string
	_appflowTags                        string
	_appflowTasks                       string
	_appflowTriggerConfig               string
)

// Cancels active runs for a flow.
// You can cancel all of the active runs for a flow, or you can cancel specific
// runs by providing their IDs.
//
// You can cancel a flow run only when the run is in progress. You can't cancel a
// run that has already completed or failed. You also can't cancel a run that's
// scheduled to occur but hasn't started yet. To prevent a scheduled run, you can
// deactivate the flow with the StopFlow action.
//
// You cannot resume a run after you cancel it.
//
// When you send your request, the status for each run becomes CancelStarted . When
// the cancellation completes, the status becomes Canceled .
//
// When you cancel a run, you still incur charges for any data that the run
// already processed before the cancellation. If the run had already written some
// data to the flow destination, then that data remains in the destination. If you
// configured the flow to use a batch API (such as the Salesforce Bulk API 2.0),
// then the run will finish reading or writing its entire batch of data after the
// cancellation. For these operations, the data processing charges for Amazon
// AppFlow apply. For the pricing information, see [Amazon AppFlow pricing].
//
// [Amazon AppFlow pricing]: http://aws.amazon.com/appflow/pricing/
func appflow_CancelFlowExecutions(cfg aws.Config, client *appflow.Client) {
	input := &appflow.CancelFlowExecutionsInput{
		// FlowName: *string, // Required
	}

	if len(_appflowFlowName) > 0 {
		input.FlowName = aws.String(_appflowFlowName)
	}
	if len(_appflowExecutionIds) > 0 {
		input.ExecutionIds = append([]string(nil), _appflowExecutionIds...)
	}

	if resp, err := client.CancelFlowExecutions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new connector profile associated with your Amazon Web Services
// account. There is a soft quota of 100 connector profiles per Amazon Web Services
// account. If you need more connector profiles than this quota allows, you can
// submit a request to the Amazon AppFlow team through the Amazon AppFlow support
// channel. In each connector profile that you create, you can provide the
// credentials and properties for only one connector.
func appflow_CreateConnectorProfile(cfg aws.Config, client *appflow.Client) {
	input := &appflow.CreateConnectorProfileInput{
		// ConnectionMode: types.ConnectionMode, // Required
		// ConnectorProfileConfig: *types.ConnectorProfileConfig, // Required
		// ConnectorProfileName: *string, // Required
		// ConnectorType: types.ConnectorType, // Required
	}

	if len(_appflowConnectionMode) > 0 {
		if err := assignInputField(input, "ConnectionMode", _appflowConnectionMode); err != nil {
			log.Errorf("invalid --connection-mode: %s", err.Error())
			return
		}
	}
	if len(_appflowConnectorProfileConfig) > 0 {
		if err := assignInputField(input, "ConnectorProfileConfig", _appflowConnectorProfileConfig); err != nil {
			log.Errorf("invalid --connector-profile-config: %s", err.Error())
			return
		}
	}
	if len(_appflowConnectorProfileName) > 0 {
		input.ConnectorProfileName = aws.String(_appflowConnectorProfileName)
	}
	if len(_appflowConnectorType) > 0 {
		if err := assignInputField(input, "ConnectorType", _appflowConnectorType); err != nil {
			log.Errorf("invalid --connector-type: %s", err.Error())
			return
		}
	}
	if len(_appflowClientToken) > 0 {
		input.ClientToken = aws.String(_appflowClientToken)
	}
	if len(_appflowConnectorLabel) > 0 {
		input.ConnectorLabel = aws.String(_appflowConnectorLabel)
	}
	if len(_appflowKmsArn) > 0 {
		input.KmsArn = aws.String(_appflowKmsArn)
	}

	if resp, err := client.CreateConnectorProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables your application to create a new flow using Amazon AppFlow. You must
// create a connector profile before calling this API. Please note that the Request
// Syntax below shows syntax for multiple destinations, however, you can only
// transfer data to one item in this list at a time. Amazon AppFlow does not
// currently support flows to multiple destinations at once.
func appflow_CreateFlow(cfg aws.Config, client *appflow.Client) {
	input := &appflow.CreateFlowInput{
		// DestinationFlowConfigList: []types.DestinationFlowConfig, // Required
		// FlowName: *string, // Required
		// SourceFlowConfig: *types.SourceFlowConfig, // Required
		// Tasks: []types.Task, // Required
		// TriggerConfig: *types.TriggerConfig, // Required
	}

	if len(_appflowDestinationFlowConfigList) > 0 {
		if err := assignInputField(input, "DestinationFlowConfigList", _appflowDestinationFlowConfigList); err != nil {
			log.Errorf("invalid --destination-flow-config-list: %s", err.Error())
			return
		}
	}
	if len(_appflowFlowName) > 0 {
		input.FlowName = aws.String(_appflowFlowName)
	}
	if len(_appflowSourceFlowConfig) > 0 {
		if err := assignInputField(input, "SourceFlowConfig", _appflowSourceFlowConfig); err != nil {
			log.Errorf("invalid --source-flow-config: %s", err.Error())
			return
		}
	}
	if len(_appflowTasks) > 0 {
		if err := assignInputField(input, "Tasks", _appflowTasks); err != nil {
			log.Errorf("invalid --tasks: %s", err.Error())
			return
		}
	}
	if len(_appflowTriggerConfig) > 0 {
		if err := assignInputField(input, "TriggerConfig", _appflowTriggerConfig); err != nil {
			log.Errorf("invalid --trigger-config: %s", err.Error())
			return
		}
	}
	if len(_appflowClientToken) > 0 {
		input.ClientToken = aws.String(_appflowClientToken)
	}
	if len(_appflowDescription) > 0 {
		input.Description = aws.String(_appflowDescription)
	}
	if len(_appflowKmsArn) > 0 {
		input.KmsArn = aws.String(_appflowKmsArn)
	}
	if len(_appflowMetadataCatalogConfig) > 0 {
		if err := assignInputField(input, "MetadataCatalogConfig", _appflowMetadataCatalogConfig); err != nil {
			log.Errorf("invalid --metadata-catalog-config: %s", err.Error())
			return
		}
	}
	if len(_appflowTags) > 0 {
		if err := assignInputField(input, "Tags", _appflowTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables you to delete an existing connector profile.
func appflow_DeleteConnectorProfile(cfg aws.Config, client *appflow.Client) {
	input := &appflow.DeleteConnectorProfileInput{
		// ConnectorProfileName: *string, // Required
	}

	if len(_appflowConnectorProfileName) > 0 {
		input.ConnectorProfileName = aws.String(_appflowConnectorProfileName)
	}
	if len(_appflowForceDelete) > 0 {
		if err := assignInputField(input, "ForceDelete", _appflowForceDelete); err != nil {
			log.Errorf("invalid --force-delete: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteConnectorProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables your application to delete an existing flow. Before deleting the flow,
// Amazon AppFlow validates the request by checking the flow configuration and
// status. You can delete flows one at a time.
func appflow_DeleteFlow(cfg aws.Config, client *appflow.Client) {
	input := &appflow.DeleteFlowInput{
		// FlowName: *string, // Required
	}

	if len(_appflowFlowName) > 0 {
		input.FlowName = aws.String(_appflowFlowName)
	}
	if len(_appflowForceDelete) > 0 {
		if err := assignInputField(input, "ForceDelete", _appflowForceDelete); err != nil {
			log.Errorf("invalid --force-delete: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the given custom connector registered in your Amazon Web Services
// account. This API can be used for custom connectors that are registered in your
// account and also for Amazon authored connectors.
func appflow_DescribeConnector(cfg aws.Config, client *appflow.Client) {
	input := &appflow.DescribeConnectorInput{
		// ConnectorType: types.ConnectorType, // Required
	}

	if len(_appflowConnectorType) > 0 {
		if err := assignInputField(input, "ConnectorType", _appflowConnectorType); err != nil {
			log.Errorf("invalid --connector-type: %s", err.Error())
			return
		}
	}
	if len(_appflowConnectorLabel) > 0 {
		input.ConnectorLabel = aws.String(_appflowConnectorLabel)
	}

	if resp, err := client.DescribeConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides details regarding the entity used with the connector, with a
// description of the data model for each field in that entity.
func appflow_DescribeConnectorEntity(cfg aws.Config, client *appflow.Client) {
	input := &appflow.DescribeConnectorEntityInput{
		// ConnectorEntityName: *string, // Required
	}

	if len(_appflowConnectorEntityName) > 0 {
		input.ConnectorEntityName = aws.String(_appflowConnectorEntityName)
	}
	if len(_appflowApiVersion) > 0 {
		input.ApiVersion = aws.String(_appflowApiVersion)
	}
	if len(_appflowConnectorProfileName) > 0 {
		input.ConnectorProfileName = aws.String(_appflowConnectorProfileName)
	}
	if len(_appflowConnectorType) > 0 {
		if err := assignInputField(input, "ConnectorType", _appflowConnectorType); err != nil {
			log.Errorf("invalid --connector-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeConnectorEntity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of connector-profile details matching the provided
// connector-profile names and connector-types . Both input lists are optional, and
// you can use them to filter the result.
//
// If no names or connector-types are provided, returns all connector profiles in
// a paginated form. If there is no match, this operation returns an empty list.
func appflow_DescribeConnectorProfiles(cfg aws.Config, client *appflow.Client) {
	input := &appflow.DescribeConnectorProfilesInput{}

	if len(_appflowConnectorLabel) > 0 {
		input.ConnectorLabel = aws.String(_appflowConnectorLabel)
	}
	if len(_appflowConnectorProfileNames) > 0 {
		input.ConnectorProfileNames = append([]string(nil), _appflowConnectorProfileNames...)
	}
	if len(_appflowConnectorType) > 0 {
		if err := assignInputField(input, "ConnectorType", _appflowConnectorType); err != nil {
			log.Errorf("invalid --connector-type: %s", err.Error())
			return
		}
	}
	if len(_appflowMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appflowMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appflowNextToken) > 0 {
		input.NextToken = aws.String(_appflowNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeConnectorProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appflow.DescribeConnectorProfilesOutput
	p := appflow.NewDescribeConnectorProfilesPaginator(client, input)
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

// Describes the connectors vended by Amazon AppFlow for specified connector
// types. If you don't specify a connector type, this operation describes all
// connectors vended by Amazon AppFlow. If there are more connectors than can be
// returned in one page, the response contains a nextToken object, which can be be
// passed in to the next call to the DescribeConnectors API operation to retrieve
// the next page.
func appflow_DescribeConnectors(cfg aws.Config, client *appflow.Client) {
	input := &appflow.DescribeConnectorsInput{}

	if len(_appflowConnectorTypes) > 0 {
		if err := assignInputField(input, "ConnectorTypes", _appflowConnectorTypes); err != nil {
			log.Errorf("invalid --connector-types: %s", err.Error())
			return
		}
	}
	if len(_appflowMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appflowMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appflowNextToken) > 0 {
		input.NextToken = aws.String(_appflowNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeConnectors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appflow.DescribeConnectorsOutput
	p := appflow.NewDescribeConnectorsPaginator(client, input)
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

// Provides a description of the specified flow.
func appflow_DescribeFlow(cfg aws.Config, client *appflow.Client) {
	input := &appflow.DescribeFlowInput{
		// FlowName: *string, // Required
	}

	if len(_appflowFlowName) > 0 {
		input.FlowName = aws.String(_appflowFlowName)
	}

	if resp, err := client.DescribeFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Fetches the execution history of the flow.
func appflow_DescribeFlowExecutionRecords(cfg aws.Config, client *appflow.Client) {
	input := &appflow.DescribeFlowExecutionRecordsInput{
		// FlowName: *string, // Required
	}

	if len(_appflowFlowName) > 0 {
		input.FlowName = aws.String(_appflowFlowName)
	}
	if len(_appflowMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appflowMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appflowNextToken) > 0 {
		input.NextToken = aws.String(_appflowNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeFlowExecutionRecords(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appflow.DescribeFlowExecutionRecordsOutput
	p := appflow.NewDescribeFlowExecutionRecordsPaginator(client, input)
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

// Returns the list of available connector entities supported by Amazon AppFlow.
// For example, you can query Salesforce for Account and Opportunity entities, or
// query ServiceNow for the Incident entity.
func appflow_ListConnectorEntities(cfg aws.Config, client *appflow.Client) {
	input := &appflow.ListConnectorEntitiesInput{}

	if len(_appflowApiVersion) > 0 {
		input.ApiVersion = aws.String(_appflowApiVersion)
	}
	if len(_appflowConnectorProfileName) > 0 {
		input.ConnectorProfileName = aws.String(_appflowConnectorProfileName)
	}
	if len(_appflowConnectorType) > 0 {
		if err := assignInputField(input, "ConnectorType", _appflowConnectorType); err != nil {
			log.Errorf("invalid --connector-type: %s", err.Error())
			return
		}
	}
	if len(_appflowEntitiesPath) > 0 {
		input.EntitiesPath = aws.String(_appflowEntitiesPath)
	}
	if len(_appflowMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appflowMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appflowNextToken) > 0 {
		input.NextToken = aws.String(_appflowNextToken)
	}

	if resp, err := client.ListConnectorEntities(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the list of all registered custom connectors in your Amazon Web
// Services account. This API lists only custom connectors registered in this
// account, not the Amazon Web Services authored connectors.
func appflow_ListConnectors(cfg aws.Config, client *appflow.Client) {
	input := &appflow.ListConnectorsInput{}

	if len(_appflowMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appflowMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appflowNextToken) > 0 {
		input.NextToken = aws.String(_appflowNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConnectors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appflow.ListConnectorsOutput
	p := appflow.NewListConnectorsPaginator(client, input)
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

// Lists all of the flows associated with your account.
func appflow_ListFlows(cfg aws.Config, client *appflow.Client) {
	input := &appflow.ListFlowsInput{}

	if len(_appflowMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appflowMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appflowNextToken) > 0 {
		input.NextToken = aws.String(_appflowNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFlows(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appflow.ListFlowsOutput
	p := appflow.NewListFlowsPaginator(client, input)
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

// Retrieves the tags that are associated with a specified flow.
func appflow_ListTagsForResource(cfg aws.Config, client *appflow.Client) {
	input := &appflow.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_appflowResourceArn) > 0 {
		input.ResourceArn = aws.String(_appflowResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers a new custom connector with your Amazon Web Services account. Before
// you can register the connector, you must deploy the associated AWS lambda
// function in your account.
func appflow_RegisterConnector(cfg aws.Config, client *appflow.Client) {
	input := &appflow.RegisterConnectorInput{}

	if len(_appflowClientToken) > 0 {
		input.ClientToken = aws.String(_appflowClientToken)
	}
	if len(_appflowConnectorLabel) > 0 {
		input.ConnectorLabel = aws.String(_appflowConnectorLabel)
	}
	if len(_appflowConnectorProvisioningConfig) > 0 {
		if err := assignInputField(input, "ConnectorProvisioningConfig", _appflowConnectorProvisioningConfig); err != nil {
			log.Errorf("invalid --connector-provisioning-config: %s", err.Error())
			return
		}
	}
	if len(_appflowConnectorProvisioningType) > 0 {
		if err := assignInputField(input, "ConnectorProvisioningType", _appflowConnectorProvisioningType); err != nil {
			log.Errorf("invalid --connector-provisioning-type: %s", err.Error())
			return
		}
	}
	if len(_appflowDescription) > 0 {
		input.Description = aws.String(_appflowDescription)
	}

	if resp, err := client.RegisterConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resets metadata about your connector entities that Amazon AppFlow stored in its
// cache. Use this action when you want Amazon AppFlow to return the latest
// information about the data that you have in a source application.
//
// Amazon AppFlow returns metadata about your entities when you use the
// ListConnectorEntities or DescribeConnectorEntities actions. Following these
// actions, Amazon AppFlow caches the metadata to reduce the number of API requests
// that it must send to the source application. Amazon AppFlow automatically resets
// the cache once every hour, but you can use this action when you want to get the
// latest metadata right away.
func appflow_ResetConnectorMetadataCache(cfg aws.Config, client *appflow.Client) {
	input := &appflow.ResetConnectorMetadataCacheInput{}

	if len(_appflowApiVersion) > 0 {
		input.ApiVersion = aws.String(_appflowApiVersion)
	}
	if len(_appflowConnectorEntityName) > 0 {
		input.ConnectorEntityName = aws.String(_appflowConnectorEntityName)
	}
	if len(_appflowConnectorProfileName) > 0 {
		input.ConnectorProfileName = aws.String(_appflowConnectorProfileName)
	}
	if len(_appflowConnectorType) > 0 {
		if err := assignInputField(input, "ConnectorType", _appflowConnectorType); err != nil {
			log.Errorf("invalid --connector-type: %s", err.Error())
			return
		}
	}
	if len(_appflowEntitiesPath) > 0 {
		input.EntitiesPath = aws.String(_appflowEntitiesPath)
	}

	if resp, err := client.ResetConnectorMetadataCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Activates an existing flow. For on-demand flows, this operation runs the flow
// immediately. For schedule and event-triggered flows, this operation activates
// the flow.
func appflow_StartFlow(cfg aws.Config, client *appflow.Client) {
	input := &appflow.StartFlowInput{
		// FlowName: *string, // Required
	}

	if len(_appflowFlowName) > 0 {
		input.FlowName = aws.String(_appflowFlowName)
	}
	if len(_appflowClientToken) > 0 {
		input.ClientToken = aws.String(_appflowClientToken)
	}

	if resp, err := client.StartFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deactivates the existing flow. For on-demand flows, this operation returns an
// unsupportedOperationException error message. For schedule and event-triggered
// flows, this operation deactivates the flow.
func appflow_StopFlow(cfg aws.Config, client *appflow.Client) {
	input := &appflow.StopFlowInput{
		// FlowName: *string, // Required
	}

	if len(_appflowFlowName) > 0 {
		input.FlowName = aws.String(_appflowFlowName)
	}

	if resp, err := client.StopFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies a tag to the specified flow.
func appflow_TagResource(cfg aws.Config, client *appflow.Client) {
	input := &appflow.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_appflowResourceArn) > 0 {
		input.ResourceArn = aws.String(_appflowResourceArn)
	}
	if len(_appflowTags) > 0 {
		if err := assignInputField(input, "Tags", _appflowTags); err != nil {
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

// Unregisters the custom connector registered in your account that matches the
// connector label provided in the request.
func appflow_UnregisterConnector(cfg aws.Config, client *appflow.Client) {
	input := &appflow.UnregisterConnectorInput{
		// ConnectorLabel: *string, // Required
	}

	if len(_appflowConnectorLabel) > 0 {
		input.ConnectorLabel = aws.String(_appflowConnectorLabel)
	}
	if len(_appflowForceDelete) > 0 {
		if err := assignInputField(input, "ForceDelete", _appflowForceDelete); err != nil {
			log.Errorf("invalid --force-delete: %s", err.Error())
			return
		}
	}

	if resp, err := client.UnregisterConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a tag from the specified flow.
func appflow_UntagResource(cfg aws.Config, client *appflow.Client) {
	input := &appflow.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_appflowResourceArn) > 0 {
		input.ResourceArn = aws.String(_appflowResourceArn)
	}
	if len(_appflowTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _appflowTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a given connector profile associated with your account.
func appflow_UpdateConnectorProfile(cfg aws.Config, client *appflow.Client) {
	input := &appflow.UpdateConnectorProfileInput{
		// ConnectionMode: types.ConnectionMode, // Required
		// ConnectorProfileConfig: *types.ConnectorProfileConfig, // Required
		// ConnectorProfileName: *string, // Required
	}

	if len(_appflowConnectionMode) > 0 {
		if err := assignInputField(input, "ConnectionMode", _appflowConnectionMode); err != nil {
			log.Errorf("invalid --connection-mode: %s", err.Error())
			return
		}
	}
	if len(_appflowConnectorProfileConfig) > 0 {
		if err := assignInputField(input, "ConnectorProfileConfig", _appflowConnectorProfileConfig); err != nil {
			log.Errorf("invalid --connector-profile-config: %s", err.Error())
			return
		}
	}
	if len(_appflowConnectorProfileName) > 0 {
		input.ConnectorProfileName = aws.String(_appflowConnectorProfileName)
	}
	if len(_appflowClientToken) > 0 {
		input.ClientToken = aws.String(_appflowClientToken)
	}

	if resp, err := client.UpdateConnectorProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a custom connector that you've previously registered. This operation
// updates the connector with one of the following:
//
// - The latest version of the AWS Lambda function that's assigned to the
// connector
//
// - A new AWS Lambda function that you specify
func appflow_UpdateConnectorRegistration(cfg aws.Config, client *appflow.Client) {
	input := &appflow.UpdateConnectorRegistrationInput{
		// ConnectorLabel: *string, // Required
	}

	if len(_appflowConnectorLabel) > 0 {
		input.ConnectorLabel = aws.String(_appflowConnectorLabel)
	}
	if len(_appflowClientToken) > 0 {
		input.ClientToken = aws.String(_appflowClientToken)
	}
	if len(_appflowConnectorProvisioningConfig) > 0 {
		if err := assignInputField(input, "ConnectorProvisioningConfig", _appflowConnectorProvisioningConfig); err != nil {
			log.Errorf("invalid --connector-provisioning-config: %s", err.Error())
			return
		}
	}
	if len(_appflowDescription) > 0 {
		input.Description = aws.String(_appflowDescription)
	}

	if resp, err := client.UpdateConnectorRegistration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing flow.
func appflow_UpdateFlow(cfg aws.Config, client *appflow.Client) {
	input := &appflow.UpdateFlowInput{
		// DestinationFlowConfigList: []types.DestinationFlowConfig, // Required
		// FlowName: *string, // Required
		// SourceFlowConfig: *types.SourceFlowConfig, // Required
		// Tasks: []types.Task, // Required
		// TriggerConfig: *types.TriggerConfig, // Required
	}

	if len(_appflowDestinationFlowConfigList) > 0 {
		if err := assignInputField(input, "DestinationFlowConfigList", _appflowDestinationFlowConfigList); err != nil {
			log.Errorf("invalid --destination-flow-config-list: %s", err.Error())
			return
		}
	}
	if len(_appflowFlowName) > 0 {
		input.FlowName = aws.String(_appflowFlowName)
	}
	if len(_appflowSourceFlowConfig) > 0 {
		if err := assignInputField(input, "SourceFlowConfig", _appflowSourceFlowConfig); err != nil {
			log.Errorf("invalid --source-flow-config: %s", err.Error())
			return
		}
	}
	if len(_appflowTasks) > 0 {
		if err := assignInputField(input, "Tasks", _appflowTasks); err != nil {
			log.Errorf("invalid --tasks: %s", err.Error())
			return
		}
	}
	if len(_appflowTriggerConfig) > 0 {
		if err := assignInputField(input, "TriggerConfig", _appflowTriggerConfig); err != nil {
			log.Errorf("invalid --trigger-config: %s", err.Error())
			return
		}
	}
	if len(_appflowClientToken) > 0 {
		input.ClientToken = aws.String(_appflowClientToken)
	}
	if len(_appflowDescription) > 0 {
		input.Description = aws.String(_appflowDescription)
	}
	if len(_appflowMetadataCatalogConfig) > 0 {
		if err := assignInputField(input, "MetadataCatalogConfig", _appflowMetadataCatalogConfig); err != nil {
			log.Errorf("invalid --metadata-catalog-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_appflowCmd)
	_appflowCmd.Flags().SortFlags = false

	_appflowCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_appflowCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_appflowCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_appflowCmd.Flags().StringVarP(&_appflowApiVersion, "api-version", "", "", "API Version")
	_appflowCmd.Flags().StringVarP(&_appflowClientToken, "client-token", "", "", "Client Token")
	_appflowCmd.Flags().StringVarP(&_appflowConnectionMode, "connection-mode", "", "", "Connection Mode")
	_appflowCmd.Flags().StringVarP(&_appflowConnectorEntityName, "connector-entity-name", "", "", "Connector Entity Name")
	_appflowCmd.Flags().StringVarP(&_appflowConnectorLabel, "connector-label", "", "", "Connector Label")
	_appflowCmd.Flags().StringVarP(&_appflowConnectorProfileConfig, "connector-profile-config", "", "", "Connector Profile Config")
	_appflowCmd.Flags().StringVarP(&_appflowConnectorProfileName, "connector-profile-name", "", "", "Connector Profile Name")
	_appflowCmd.Flags().StringSliceVarP(&_appflowConnectorProfileNames, "connector-profile-names", "", nil, "Connector Profile Names")
	_appflowCmd.Flags().StringVarP(&_appflowConnectorProvisioningConfig, "connector-provisioning-config", "", "", "Connector Provisioning Config")
	_appflowCmd.Flags().StringVarP(&_appflowConnectorProvisioningType, "connector-provisioning-type", "", "", "Connector Provisioning Type")
	_appflowCmd.Flags().StringVarP(&_appflowConnectorType, "connector-type", "", "", "Connector Type")
	_appflowCmd.Flags().StringVarP(&_appflowConnectorTypes, "connector-types", "", "", "Connector Types")
	_appflowCmd.Flags().StringVarP(&_appflowDescription, "description", "", "", "Description")
	_appflowCmd.Flags().StringVarP(&_appflowDestinationFlowConfigList, "destination-flow-config-list", "", "", "Destination Flow Config List")
	_appflowCmd.Flags().StringVarP(&_appflowEntitiesPath, "entities-path", "", "", "Entities Path")
	_appflowCmd.Flags().StringSliceVarP(&_appflowExecutionIds, "execution-ids", "", nil, "Execution Ids")
	_appflowCmd.Flags().StringVarP(&_appflowFlowName, "flow-name", "", "", "Flow Name")
	_appflowCmd.Flags().StringVarP(&_appflowForceDelete, "force-delete", "", "", "Force Delete")
	_appflowCmd.Flags().StringVarP(&_appflowKmsArn, "kms-arn", "", "", "KMS ARN")
	_appflowCmd.Flags().StringVarP(&_appflowMaxResults, "max-results", "", "", "Max Results")
	_appflowCmd.Flags().StringVarP(&_appflowMetadataCatalogConfig, "metadata-catalog-config", "", "", "Metadata Catalog Config")
	_appflowCmd.Flags().StringVarP(&_appflowNextToken, "next-token", "", "", "Next Token")
	_appflowCmd.Flags().StringVarP(&_appflowResourceArn, "resource-arn", "", "", "Resource ARN")
	_appflowCmd.Flags().StringVarP(&_appflowSourceFlowConfig, "source-flow-config", "", "", "Source Flow Config")
	_appflowCmd.Flags().StringSliceVarP(&_appflowTagKeys, "tag-keys", "", nil, "Tag Keys")
	_appflowCmd.Flags().StringVarP(&_appflowTags, "tags", "", "", "Tags")
	_appflowCmd.Flags().StringVarP(&_appflowTasks, "tasks", "", "", "Tasks")
	_appflowCmd.Flags().StringVarP(&_appflowTriggerConfig, "trigger-config", "", "", "Trigger Config")

	_appflowCmd.Flags().BoolVarP(&_appflowCancelFlowExecutions, "cancel-flow-executions", "", false, "Cancel Flow Executions")
	_appflowCmd.Flags().BoolVarP(&_appflowCreateConnectorProfile, "create-connector-profile", "", false, "Create Connector Profile")
	_appflowCmd.Flags().BoolVarP(&_appflowCreateFlow, "create-flow", "", false, "Create Flow")
	_appflowCmd.Flags().BoolVarP(&_appflowDeleteConnectorProfile, "delete-connector-profile", "", false, "Delete Connector Profile")
	_appflowCmd.Flags().BoolVarP(&_appflowDeleteFlow, "delete-flow", "", false, "Delete Flow")
	_appflowCmd.Flags().BoolVarP(&_appflowDescribeConnector, "describe-connector", "", false, "Describe Connector")
	_appflowCmd.Flags().BoolVarP(&_appflowDescribeConnectorEntity, "describe-connector-entity", "", false, "Describe Connector Entity")
	_appflowCmd.Flags().BoolVarP(&_appflowDescribeConnectorProfiles, "describe-connector-profiles", "", false, "Describe Connector Profiles")
	_appflowCmd.Flags().BoolVarP(&_appflowDescribeConnectors, "describe-connectors", "", false, "Describe Connectors")
	_appflowCmd.Flags().BoolVarP(&_appflowDescribeFlow, "describe-flow", "", false, "Describe Flow")
	_appflowCmd.Flags().BoolVarP(&_appflowDescribeFlowExecutionRecords, "describe-flow-execution-records", "", false, "Describe Flow Execution Records")
	_appflowCmd.Flags().BoolVarP(&_appflowListConnectorEntities, "list-connector-entities", "", false, "List Connector Entities")
	_appflowCmd.Flags().BoolVarP(&_appflowListConnectors, "list-connectors", "", false, "List Connectors")
	_appflowCmd.Flags().BoolVarP(&_appflowListFlows, "list-flows", "", false, "List Flows")
	_appflowCmd.Flags().BoolVarP(&_appflowListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_appflowCmd.Flags().BoolVarP(&_appflowRegisterConnector, "register-connector", "", false, "Register Connector")
	_appflowCmd.Flags().BoolVarP(&_appflowResetConnectorMetadataCache, "reset-connector-metadata-cache", "", false, "Reset Connector Metadata Cache")
	_appflowCmd.Flags().BoolVarP(&_appflowStartFlow, "start-flow", "", false, "Start Flow")
	_appflowCmd.Flags().BoolVarP(&_appflowStopFlow, "stop-flow", "", false, "Stop Flow")
	_appflowCmd.Flags().BoolVarP(&_appflowTagResource, "tag-resource", "", false, "Tag Resource")
	_appflowCmd.Flags().BoolVarP(&_appflowUnregisterConnector, "unregister-connector", "", false, "Unregister Connector")
	_appflowCmd.Flags().BoolVarP(&_appflowUntagResource, "untag-resource", "", false, "Untag Resource")
	_appflowCmd.Flags().BoolVarP(&_appflowUpdateConnectorProfile, "update-connector-profile", "", false, "Update Connector Profile")
	_appflowCmd.Flags().BoolVarP(&_appflowUpdateConnectorRegistration, "update-connector-registration", "", false, "Update Connector Registration")
	_appflowCmd.Flags().BoolVarP(&_appflowUpdateFlow, "update-flow", "", false, "Update Flow")

}
