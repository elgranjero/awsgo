package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kafkaconnect"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// kafkaconnectCmd represents the kafkaconnect command
var _kafkaconnectCmd = &cobra.Command{
	Use:   "kafkaconnect",
	Short: "AWS kafkaconnect CLI",
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
		client := kafkaconnect.NewFromConfig(cfg)
		if _kafkaconnectCreateConnector {
			kafkaconnect_CreateConnector(cfg, client)
			return
		}
		if _kafkaconnectCreateCustomPlugin {
			kafkaconnect_CreateCustomPlugin(cfg, client)
			return
		}
		if _kafkaconnectCreateWorkerConfiguration {
			kafkaconnect_CreateWorkerConfiguration(cfg, client)
			return
		}
		if _kafkaconnectDeleteConnector {
			kafkaconnect_DeleteConnector(cfg, client)
			return
		}
		if _kafkaconnectDeleteCustomPlugin {
			kafkaconnect_DeleteCustomPlugin(cfg, client)
			return
		}
		if _kafkaconnectDeleteWorkerConfiguration {
			kafkaconnect_DeleteWorkerConfiguration(cfg, client)
			return
		}
		if _kafkaconnectDescribeConnector {
			kafkaconnect_DescribeConnector(cfg, client)
			return
		}
		if _kafkaconnectDescribeConnectorOperation {
			kafkaconnect_DescribeConnectorOperation(cfg, client)
			return
		}
		if _kafkaconnectDescribeCustomPlugin {
			kafkaconnect_DescribeCustomPlugin(cfg, client)
			return
		}
		if _kafkaconnectDescribeWorkerConfiguration {
			kafkaconnect_DescribeWorkerConfiguration(cfg, client)
			return
		}
		if _kafkaconnectListConnectorOperations {
			kafkaconnect_ListConnectorOperations(cfg, client)
			return
		}
		if _kafkaconnectListConnectors {
			kafkaconnect_ListConnectors(cfg, client)
			return
		}
		if _kafkaconnectListCustomPlugins {
			kafkaconnect_ListCustomPlugins(cfg, client)
			return
		}
		if _kafkaconnectListTagsForResource {
			kafkaconnect_ListTagsForResource(cfg, client)
			return
		}
		if _kafkaconnectListWorkerConfigurations {
			kafkaconnect_ListWorkerConfigurations(cfg, client)
			return
		}
		if _kafkaconnectTagResource {
			kafkaconnect_TagResource(cfg, client)
			return
		}
		if _kafkaconnectUntagResource {
			kafkaconnect_UntagResource(cfg, client)
			return
		}
		if _kafkaconnectUpdateConnector {
			kafkaconnect_UpdateConnector(cfg, client)
			return
		}

	},
}

var (
	_kafkaconnectCreateConnector             bool
	_kafkaconnectCreateCustomPlugin          bool
	_kafkaconnectCreateWorkerConfiguration   bool
	_kafkaconnectDeleteConnector             bool
	_kafkaconnectDeleteCustomPlugin          bool
	_kafkaconnectDeleteWorkerConfiguration   bool
	_kafkaconnectDescribeConnector           bool
	_kafkaconnectDescribeConnectorOperation  bool
	_kafkaconnectDescribeCustomPlugin        bool
	_kafkaconnectDescribeWorkerConfiguration bool
	_kafkaconnectListConnectorOperations     bool
	_kafkaconnectListConnectors              bool
	_kafkaconnectListCustomPlugins           bool
	_kafkaconnectListTagsForResource         bool
	_kafkaconnectListWorkerConfigurations    bool
	_kafkaconnectTagResource                 bool
	_kafkaconnectUntagResource               bool
	_kafkaconnectUpdateConnector             bool

	_kafkaconnectCapacity                         string
	_kafkaconnectConnectorArn                     string
	_kafkaconnectConnectorConfiguration           string
	_kafkaconnectConnectorDescription             string
	_kafkaconnectConnectorName                    string
	_kafkaconnectConnectorNamePrefix              string
	_kafkaconnectConnectorOperationArn            string
	_kafkaconnectContentType                      string
	_kafkaconnectCurrentVersion                   string
	_kafkaconnectCustomPluginArn                  string
	_kafkaconnectDescription                      string
	_kafkaconnectKafkaCluster                     string
	_kafkaconnectKafkaClusterClientAuthentication string
	_kafkaconnectKafkaClusterEncryptionInTransit  string
	_kafkaconnectKafkaConnectVersion              string
	_kafkaconnectLocation                         string
	_kafkaconnectLogDelivery                      string
	_kafkaconnectMaxResults                       string
	_kafkaconnectName                             string
	_kafkaconnectNamePrefix                       string
	_kafkaconnectNetworkType                      string
	_kafkaconnectNextToken                        string
	_kafkaconnectPlugins                          string
	_kafkaconnectPropertiesFileContent            string
	_kafkaconnectResourceArn                      string
	_kafkaconnectServiceExecutionRoleArn          string
	_kafkaconnectTagKeys                          []string
	_kafkaconnectTags                             string
	_kafkaconnectWorkerConfiguration              string
	_kafkaconnectWorkerConfigurationArn           string
)

// Creates a connector using the specified properties.
func kafkaconnect_CreateConnector(cfg aws.Config, client *kafkaconnect.Client) {
	input := &kafkaconnect.CreateConnectorInput{
		// Capacity: *types.Capacity, // Required
		// ConnectorConfiguration: map[string]string, // Required
		// ConnectorName: *string, // Required
		// KafkaCluster: *types.KafkaCluster, // Required
		// KafkaClusterClientAuthentication: *types.KafkaClusterClientAuthentication, // Required
		// KafkaClusterEncryptionInTransit: *types.KafkaClusterEncryptionInTransit, // Required
		// KafkaConnectVersion: *string, // Required
		// Plugins: []types.Plugin, // Required
		// ServiceExecutionRoleArn: *string, // Required
	}

	if len(_kafkaconnectCapacity) > 0 {
		if err := assignInputField(input, "Capacity", _kafkaconnectCapacity); err != nil {
			log.Errorf("invalid --capacity: %s", err.Error())
			return
		}
	}
	if len(_kafkaconnectConnectorConfiguration) > 0 {
		if err := assignInputField(input, "ConnectorConfiguration", _kafkaconnectConnectorConfiguration); err != nil {
			log.Errorf("invalid --connector-configuration: %s", err.Error())
			return
		}
	}
	if len(_kafkaconnectConnectorName) > 0 {
		input.ConnectorName = aws.String(_kafkaconnectConnectorName)
	}
	if len(_kafkaconnectKafkaCluster) > 0 {
		if err := assignInputField(input, "KafkaCluster", _kafkaconnectKafkaCluster); err != nil {
			log.Errorf("invalid --kafka-cluster: %s", err.Error())
			return
		}
	}
	if len(_kafkaconnectKafkaClusterClientAuthentication) > 0 {
		if err := assignInputField(input, "KafkaClusterClientAuthentication", _kafkaconnectKafkaClusterClientAuthentication); err != nil {
			log.Errorf("invalid --kafka-cluster-client-authentication: %s", err.Error())
			return
		}
	}
	if len(_kafkaconnectKafkaClusterEncryptionInTransit) > 0 {
		if err := assignInputField(input, "KafkaClusterEncryptionInTransit", _kafkaconnectKafkaClusterEncryptionInTransit); err != nil {
			log.Errorf("invalid --kafka-cluster-encryption-in-transit: %s", err.Error())
			return
		}
	}
	if len(_kafkaconnectKafkaConnectVersion) > 0 {
		input.KafkaConnectVersion = aws.String(_kafkaconnectKafkaConnectVersion)
	}
	if len(_kafkaconnectPlugins) > 0 {
		if err := assignInputField(input, "Plugins", _kafkaconnectPlugins); err != nil {
			log.Errorf("invalid --plugins: %s", err.Error())
			return
		}
	}
	if len(_kafkaconnectServiceExecutionRoleArn) > 0 {
		input.ServiceExecutionRoleArn = aws.String(_kafkaconnectServiceExecutionRoleArn)
	}
	if len(_kafkaconnectConnectorDescription) > 0 {
		input.ConnectorDescription = aws.String(_kafkaconnectConnectorDescription)
	}
	if len(_kafkaconnectLogDelivery) > 0 {
		if err := assignInputField(input, "LogDelivery", _kafkaconnectLogDelivery); err != nil {
			log.Errorf("invalid --log-delivery: %s", err.Error())
			return
		}
	}
	if len(_kafkaconnectNetworkType) > 0 {
		if err := assignInputField(input, "NetworkType", _kafkaconnectNetworkType); err != nil {
			log.Errorf("invalid --network-type: %s", err.Error())
			return
		}
	}
	if len(_kafkaconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _kafkaconnectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_kafkaconnectWorkerConfiguration) > 0 {
		if err := assignInputField(input, "WorkerConfiguration", _kafkaconnectWorkerConfiguration); err != nil {
			log.Errorf("invalid --worker-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a custom plugin using the specified properties.
func kafkaconnect_CreateCustomPlugin(cfg aws.Config, client *kafkaconnect.Client) {
	input := &kafkaconnect.CreateCustomPluginInput{
		// ContentType: types.CustomPluginContentType, // Required
		// Location: *types.CustomPluginLocation, // Required
		// Name: *string, // Required
	}

	if len(_kafkaconnectContentType) > 0 {
		if err := assignInputField(input, "ContentType", _kafkaconnectContentType); err != nil {
			log.Errorf("invalid --content-type: %s", err.Error())
			return
		}
	}
	if len(_kafkaconnectLocation) > 0 {
		if err := assignInputField(input, "Location", _kafkaconnectLocation); err != nil {
			log.Errorf("invalid --location: %s", err.Error())
			return
		}
	}
	if len(_kafkaconnectName) > 0 {
		input.Name = aws.String(_kafkaconnectName)
	}
	if len(_kafkaconnectDescription) > 0 {
		input.Description = aws.String(_kafkaconnectDescription)
	}
	if len(_kafkaconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _kafkaconnectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCustomPlugin(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a worker configuration using the specified properties.
func kafkaconnect_CreateWorkerConfiguration(cfg aws.Config, client *kafkaconnect.Client) {
	input := &kafkaconnect.CreateWorkerConfigurationInput{
		// Name: *string, // Required
		// PropertiesFileContent: *string, // Required
	}

	if len(_kafkaconnectName) > 0 {
		input.Name = aws.String(_kafkaconnectName)
	}
	if len(_kafkaconnectPropertiesFileContent) > 0 {
		input.PropertiesFileContent = aws.String(_kafkaconnectPropertiesFileContent)
	}
	if len(_kafkaconnectDescription) > 0 {
		input.Description = aws.String(_kafkaconnectDescription)
	}
	if len(_kafkaconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _kafkaconnectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWorkerConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified connector.
func kafkaconnect_DeleteConnector(cfg aws.Config, client *kafkaconnect.Client) {
	input := &kafkaconnect.DeleteConnectorInput{
		// ConnectorArn: *string, // Required
	}

	if len(_kafkaconnectConnectorArn) > 0 {
		input.ConnectorArn = aws.String(_kafkaconnectConnectorArn)
	}
	if len(_kafkaconnectCurrentVersion) > 0 {
		input.CurrentVersion = aws.String(_kafkaconnectCurrentVersion)
	}

	if resp, err := client.DeleteConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a custom plugin.
func kafkaconnect_DeleteCustomPlugin(cfg aws.Config, client *kafkaconnect.Client) {
	input := &kafkaconnect.DeleteCustomPluginInput{
		// CustomPluginArn: *string, // Required
	}

	if len(_kafkaconnectCustomPluginArn) > 0 {
		input.CustomPluginArn = aws.String(_kafkaconnectCustomPluginArn)
	}

	if resp, err := client.DeleteCustomPlugin(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified worker configuration.
func kafkaconnect_DeleteWorkerConfiguration(cfg aws.Config, client *kafkaconnect.Client) {
	input := &kafkaconnect.DeleteWorkerConfigurationInput{
		// WorkerConfigurationArn: *string, // Required
	}

	if len(_kafkaconnectWorkerConfigurationArn) > 0 {
		input.WorkerConfigurationArn = aws.String(_kafkaconnectWorkerConfigurationArn)
	}

	if resp, err := client.DeleteWorkerConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns summary information about the connector.
func kafkaconnect_DescribeConnector(cfg aws.Config, client *kafkaconnect.Client) {
	input := &kafkaconnect.DescribeConnectorInput{
		// ConnectorArn: *string, // Required
	}

	if len(_kafkaconnectConnectorArn) > 0 {
		input.ConnectorArn = aws.String(_kafkaconnectConnectorArn)
	}

	if resp, err := client.DescribeConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified connector's operations.
func kafkaconnect_DescribeConnectorOperation(cfg aws.Config, client *kafkaconnect.Client) {
	input := &kafkaconnect.DescribeConnectorOperationInput{
		// ConnectorOperationArn: *string, // Required
	}

	if len(_kafkaconnectConnectorOperationArn) > 0 {
		input.ConnectorOperationArn = aws.String(_kafkaconnectConnectorOperationArn)
	}

	if resp, err := client.DescribeConnectorOperation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A summary description of the custom plugin.
func kafkaconnect_DescribeCustomPlugin(cfg aws.Config, client *kafkaconnect.Client) {
	input := &kafkaconnect.DescribeCustomPluginInput{
		// CustomPluginArn: *string, // Required
	}

	if len(_kafkaconnectCustomPluginArn) > 0 {
		input.CustomPluginArn = aws.String(_kafkaconnectCustomPluginArn)
	}

	if resp, err := client.DescribeCustomPlugin(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a worker configuration.
func kafkaconnect_DescribeWorkerConfiguration(cfg aws.Config, client *kafkaconnect.Client) {
	input := &kafkaconnect.DescribeWorkerConfigurationInput{
		// WorkerConfigurationArn: *string, // Required
	}

	if len(_kafkaconnectWorkerConfigurationArn) > 0 {
		input.WorkerConfigurationArn = aws.String(_kafkaconnectWorkerConfigurationArn)
	}

	if resp, err := client.DescribeWorkerConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists information about a connector's operation(s).
func kafkaconnect_ListConnectorOperations(cfg aws.Config, client *kafkaconnect.Client) {
	input := &kafkaconnect.ListConnectorOperationsInput{
		// ConnectorArn: *string, // Required
	}

	if len(_kafkaconnectConnectorArn) > 0 {
		input.ConnectorArn = aws.String(_kafkaconnectConnectorArn)
	}
	if len(_kafkaconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kafkaconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kafkaconnectNextToken) > 0 {
		input.NextToken = aws.String(_kafkaconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConnectorOperations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kafkaconnect.ListConnectorOperationsOutput
	p := kafkaconnect.NewListConnectorOperationsPaginator(client, input)
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

// Returns a list of all the connectors in this account and Region. The list is
// limited to connectors whose name starts with the specified prefix. The response
// also includes a description of each of the listed connectors.
func kafkaconnect_ListConnectors(cfg aws.Config, client *kafkaconnect.Client) {
	input := &kafkaconnect.ListConnectorsInput{}

	if len(_kafkaconnectConnectorNamePrefix) > 0 {
		input.ConnectorNamePrefix = aws.String(_kafkaconnectConnectorNamePrefix)
	}
	if len(_kafkaconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kafkaconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kafkaconnectNextToken) > 0 {
		input.NextToken = aws.String(_kafkaconnectNextToken)
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

	var results []*kafkaconnect.ListConnectorsOutput
	p := kafkaconnect.NewListConnectorsPaginator(client, input)
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

// Returns a list of all of the custom plugins in this account and Region.
func kafkaconnect_ListCustomPlugins(cfg aws.Config, client *kafkaconnect.Client) {
	input := &kafkaconnect.ListCustomPluginsInput{}

	if len(_kafkaconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kafkaconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kafkaconnectNamePrefix) > 0 {
		input.NamePrefix = aws.String(_kafkaconnectNamePrefix)
	}
	if len(_kafkaconnectNextToken) > 0 {
		input.NextToken = aws.String(_kafkaconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCustomPlugins(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kafkaconnect.ListCustomPluginsOutput
	p := kafkaconnect.NewListCustomPluginsPaginator(client, input)
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

// Lists all the tags attached to the specified resource.
func kafkaconnect_ListTagsForResource(cfg aws.Config, client *kafkaconnect.Client) {
	input := &kafkaconnect.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_kafkaconnectResourceArn) > 0 {
		input.ResourceArn = aws.String(_kafkaconnectResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of all of the worker configurations in this account and Region.
func kafkaconnect_ListWorkerConfigurations(cfg aws.Config, client *kafkaconnect.Client) {
	input := &kafkaconnect.ListWorkerConfigurationsInput{}

	if len(_kafkaconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kafkaconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kafkaconnectNamePrefix) > 0 {
		input.NamePrefix = aws.String(_kafkaconnectNamePrefix)
	}
	if len(_kafkaconnectNextToken) > 0 {
		input.NextToken = aws.String(_kafkaconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkerConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kafkaconnect.ListWorkerConfigurationsOutput
	p := kafkaconnect.NewListWorkerConfigurationsPaginator(client, input)
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

// Attaches tags to the specified resource.
func kafkaconnect_TagResource(cfg aws.Config, client *kafkaconnect.Client) {
	input := &kafkaconnect.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_kafkaconnectResourceArn) > 0 {
		input.ResourceArn = aws.String(_kafkaconnectResourceArn)
	}
	if len(_kafkaconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _kafkaconnectTags); err != nil {
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
func kafkaconnect_UntagResource(cfg aws.Config, client *kafkaconnect.Client) {
	input := &kafkaconnect.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_kafkaconnectResourceArn) > 0 {
		input.ResourceArn = aws.String(_kafkaconnectResourceArn)
	}
	if len(_kafkaconnectTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _kafkaconnectTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified connector. For request body, specify only one parameter:
// either capacity or connectorConfiguration .
func kafkaconnect_UpdateConnector(cfg aws.Config, client *kafkaconnect.Client) {
	input := &kafkaconnect.UpdateConnectorInput{
		// ConnectorArn: *string, // Required
		// CurrentVersion: *string, // Required
	}

	if len(_kafkaconnectConnectorArn) > 0 {
		input.ConnectorArn = aws.String(_kafkaconnectConnectorArn)
	}
	if len(_kafkaconnectCurrentVersion) > 0 {
		input.CurrentVersion = aws.String(_kafkaconnectCurrentVersion)
	}
	if len(_kafkaconnectCapacity) > 0 {
		if err := assignInputField(input, "Capacity", _kafkaconnectCapacity); err != nil {
			log.Errorf("invalid --capacity: %s", err.Error())
			return
		}
	}
	if len(_kafkaconnectConnectorConfiguration) > 0 {
		if err := assignInputField(input, "ConnectorConfiguration", _kafkaconnectConnectorConfiguration); err != nil {
			log.Errorf("invalid --connector-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_kafkaconnectCmd)
	_kafkaconnectCmd.Flags().SortFlags = false

	_kafkaconnectCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_kafkaconnectCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_kafkaconnectCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectCapacity, "capacity", "", "", "Capacity")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectConnectorArn, "connector-arn", "", "", "Connector ARN")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectConnectorConfiguration, "connector-configuration", "", "", "Connector Configuration")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectConnectorDescription, "connector-description", "", "", "Connector Description")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectConnectorName, "connector-name", "", "", "Connector Name")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectConnectorNamePrefix, "connector-name-prefix", "", "", "Connector Name Prefix")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectConnectorOperationArn, "connector-operation-arn", "", "", "Connector Operation ARN")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectContentType, "content-type", "", "", "Content Type")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectCurrentVersion, "current-version", "", "", "Current Version")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectCustomPluginArn, "custom-plugin-arn", "", "", "Custom Plugin ARN")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectDescription, "description", "", "", "Description")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectKafkaCluster, "kafka-cluster", "", "", "Kafka Cluster")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectKafkaClusterClientAuthentication, "kafka-cluster-client-authentication", "", "", "Kafka Cluster Client Authentication")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectKafkaClusterEncryptionInTransit, "kafka-cluster-encryption-in-transit", "", "", "Kafka Cluster Encryption In Transit")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectKafkaConnectVersion, "kafka-connect-version", "", "", "Kafka Connect Version")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectLocation, "location", "", "", "Location")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectLogDelivery, "log-delivery", "", "", "Log Delivery")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectMaxResults, "max-results", "", "", "Max Results")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectName, "name", "", "", "Name")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectNamePrefix, "name-prefix", "", "", "Name Prefix")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectNetworkType, "network-type", "", "", "Network Type")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectNextToken, "next-token", "", "", "Next Token")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectPlugins, "plugins", "", "", "Plugins")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectPropertiesFileContent, "properties-file-content", "", "", "Properties File Content")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectResourceArn, "resource-arn", "", "", "Resource ARN")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectServiceExecutionRoleArn, "service-execution-role-arn", "", "", "Service Execution Role ARN")
	_kafkaconnectCmd.Flags().StringSliceVarP(&_kafkaconnectTagKeys, "tag-keys", "", nil, "Tag Keys")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectTags, "tags", "", "", "Tags")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectWorkerConfiguration, "worker-configuration", "", "", "Worker Configuration")
	_kafkaconnectCmd.Flags().StringVarP(&_kafkaconnectWorkerConfigurationArn, "worker-configuration-arn", "", "", "Worker Configuration ARN")

	_kafkaconnectCmd.Flags().BoolVarP(&_kafkaconnectCreateConnector, "create-connector", "", false, "Create Connector")
	_kafkaconnectCmd.Flags().BoolVarP(&_kafkaconnectCreateCustomPlugin, "create-custom-plugin", "", false, "Create Custom Plugin")
	_kafkaconnectCmd.Flags().BoolVarP(&_kafkaconnectCreateWorkerConfiguration, "create-worker-configuration", "", false, "Create Worker Configuration")
	_kafkaconnectCmd.Flags().BoolVarP(&_kafkaconnectDeleteConnector, "delete-connector", "", false, "Delete Connector")
	_kafkaconnectCmd.Flags().BoolVarP(&_kafkaconnectDeleteCustomPlugin, "delete-custom-plugin", "", false, "Delete Custom Plugin")
	_kafkaconnectCmd.Flags().BoolVarP(&_kafkaconnectDeleteWorkerConfiguration, "delete-worker-configuration", "", false, "Delete Worker Configuration")
	_kafkaconnectCmd.Flags().BoolVarP(&_kafkaconnectDescribeConnector, "describe-connector", "", false, "Describe Connector")
	_kafkaconnectCmd.Flags().BoolVarP(&_kafkaconnectDescribeConnectorOperation, "describe-connector-operation", "", false, "Describe Connector Operation")
	_kafkaconnectCmd.Flags().BoolVarP(&_kafkaconnectDescribeCustomPlugin, "describe-custom-plugin", "", false, "Describe Custom Plugin")
	_kafkaconnectCmd.Flags().BoolVarP(&_kafkaconnectDescribeWorkerConfiguration, "describe-worker-configuration", "", false, "Describe Worker Configuration")
	_kafkaconnectCmd.Flags().BoolVarP(&_kafkaconnectListConnectorOperations, "list-connector-operations", "", false, "List Connector Operations")
	_kafkaconnectCmd.Flags().BoolVarP(&_kafkaconnectListConnectors, "list-connectors", "", false, "List Connectors")
	_kafkaconnectCmd.Flags().BoolVarP(&_kafkaconnectListCustomPlugins, "list-custom-plugins", "", false, "List Custom Plugins")
	_kafkaconnectCmd.Flags().BoolVarP(&_kafkaconnectListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_kafkaconnectCmd.Flags().BoolVarP(&_kafkaconnectListWorkerConfigurations, "list-worker-configurations", "", false, "List Worker Configurations")
	_kafkaconnectCmd.Flags().BoolVarP(&_kafkaconnectTagResource, "tag-resource", "", false, "Tag Resource")
	_kafkaconnectCmd.Flags().BoolVarP(&_kafkaconnectUntagResource, "untag-resource", "", false, "Untag Resource")
	_kafkaconnectCmd.Flags().BoolVarP(&_kafkaconnectUpdateConnector, "update-connector", "", false, "Update Connector")

}
