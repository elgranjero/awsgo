package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/supplychain"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// supplychainCmd represents the supplychain command
var _supplychainCmd = &cobra.Command{
	Use:   "supplychain",
	Short: "AWS supplychain CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := supplychain.NewFromConfig(cfg)
		if _supplychainCreateBillOfMaterialsImportJob {
			supplychain_CreateBillOfMaterialsImportJob(cfg, client)
			return
		}
		if _supplychainCreateDataIntegrationFlow {
			supplychain_CreateDataIntegrationFlow(cfg, client)
			return
		}
		if _supplychainCreateDataLakeDataset {
			supplychain_CreateDataLakeDataset(cfg, client)
			return
		}
		if _supplychainCreateDataLakeNamespace {
			supplychain_CreateDataLakeNamespace(cfg, client)
			return
		}
		if _supplychainCreateInstance {
			supplychain_CreateInstance(cfg, client)
			return
		}
		if _supplychainDeleteDataIntegrationFlow {
			supplychain_DeleteDataIntegrationFlow(cfg, client)
			return
		}
		if _supplychainDeleteDataLakeDataset {
			supplychain_DeleteDataLakeDataset(cfg, client)
			return
		}
		if _supplychainDeleteDataLakeNamespace {
			supplychain_DeleteDataLakeNamespace(cfg, client)
			return
		}
		if _supplychainDeleteInstance {
			supplychain_DeleteInstance(cfg, client)
			return
		}
		if _supplychainGetBillOfMaterialsImportJob {
			supplychain_GetBillOfMaterialsImportJob(cfg, client)
			return
		}
		if _supplychainGetDataIntegrationEvent {
			supplychain_GetDataIntegrationEvent(cfg, client)
			return
		}
		if _supplychainGetDataIntegrationFlow {
			supplychain_GetDataIntegrationFlow(cfg, client)
			return
		}
		if _supplychainGetDataIntegrationFlowExecution {
			supplychain_GetDataIntegrationFlowExecution(cfg, client)
			return
		}
		if _supplychainGetDataLakeDataset {
			supplychain_GetDataLakeDataset(cfg, client)
			return
		}
		if _supplychainGetDataLakeNamespace {
			supplychain_GetDataLakeNamespace(cfg, client)
			return
		}
		if _supplychainGetInstance {
			supplychain_GetInstance(cfg, client)
			return
		}
		if _supplychainListDataIntegrationEvents {
			supplychain_ListDataIntegrationEvents(cfg, client)
			return
		}
		if _supplychainListDataIntegrationFlowExecutions {
			supplychain_ListDataIntegrationFlowExecutions(cfg, client)
			return
		}
		if _supplychainListDataIntegrationFlows {
			supplychain_ListDataIntegrationFlows(cfg, client)
			return
		}
		if _supplychainListDataLakeDatasets {
			supplychain_ListDataLakeDatasets(cfg, client)
			return
		}
		if _supplychainListDataLakeNamespaces {
			supplychain_ListDataLakeNamespaces(cfg, client)
			return
		}
		if _supplychainListInstances {
			supplychain_ListInstances(cfg, client)
			return
		}
		if _supplychainListTagsForResource {
			supplychain_ListTagsForResource(cfg, client)
			return
		}
		if _supplychainSendDataIntegrationEvent {
			supplychain_SendDataIntegrationEvent(cfg, client)
			return
		}
		if _supplychainTagResource {
			supplychain_TagResource(cfg, client)
			return
		}
		if _supplychainUntagResource {
			supplychain_UntagResource(cfg, client)
			return
		}
		if _supplychainUpdateDataIntegrationFlow {
			supplychain_UpdateDataIntegrationFlow(cfg, client)
			return
		}
		if _supplychainUpdateDataLakeDataset {
			supplychain_UpdateDataLakeDataset(cfg, client)
			return
		}
		if _supplychainUpdateDataLakeNamespace {
			supplychain_UpdateDataLakeNamespace(cfg, client)
			return
		}
		if _supplychainUpdateInstance {
			supplychain_UpdateInstance(cfg, client)
			return
		}

	},
}

var (
	_supplychainCreateBillOfMaterialsImportJob    bool
	_supplychainCreateDataIntegrationFlow         bool
	_supplychainCreateDataLakeDataset             bool
	_supplychainCreateDataLakeNamespace           bool
	_supplychainCreateInstance                    bool
	_supplychainDeleteDataIntegrationFlow         bool
	_supplychainDeleteDataLakeDataset             bool
	_supplychainDeleteDataLakeNamespace           bool
	_supplychainDeleteInstance                    bool
	_supplychainGetBillOfMaterialsImportJob       bool
	_supplychainGetDataIntegrationEvent           bool
	_supplychainGetDataIntegrationFlow            bool
	_supplychainGetDataIntegrationFlowExecution   bool
	_supplychainGetDataLakeDataset                bool
	_supplychainGetDataLakeNamespace              bool
	_supplychainGetInstance                       bool
	_supplychainListDataIntegrationEvents         bool
	_supplychainListDataIntegrationFlowExecutions bool
	_supplychainListDataIntegrationFlows          bool
	_supplychainListDataLakeDatasets              bool
	_supplychainListDataLakeNamespaces            bool
	_supplychainListInstances                     bool
	_supplychainListTagsForResource               bool
	_supplychainSendDataIntegrationEvent          bool
	_supplychainTagResource                       bool
	_supplychainUntagResource                     bool
	_supplychainUpdateDataIntegrationFlow         bool
	_supplychainUpdateDataLakeDataset             bool
	_supplychainUpdateDataLakeNamespace           bool
	_supplychainUpdateInstance                    bool

	_supplychainClientToken         string
	_supplychainData                string
	_supplychainDatasetTarget       string
	_supplychainDescription         string
	_supplychainEventGroupId        string
	_supplychainEventId             string
	_supplychainEventTimestamp      string
	_supplychainEventType           string
	_supplychainExecutionId         string
	_supplychainFlowName            string
	_supplychainInstanceDescription string
	_supplychainInstanceId          string
	_supplychainInstanceName        string
	_supplychainInstanceNameFilter  []string
	_supplychainInstanceStateFilter string
	_supplychainJobId               string
	_supplychainKmsKeyArn           string
	_supplychainMaxResults          string
	_supplychainName                string
	_supplychainNamespace           string
	_supplychainNextToken           string
	_supplychainPartitionSpec       string
	_supplychainResourceArn         string
	_supplychainS3uri               string
	_supplychainSchema              string
	_supplychainSources             string
	_supplychainTagKeys             []string
	_supplychainTags                string
	_supplychainTarget              string
	_supplychainTransformation      string
	_supplychainWebAppDnsDomain     string
)

// CreateBillOfMaterialsImportJob creates an import job for the Product Bill Of
// Materials (BOM) entity. For information on the product_bom entity, see the AWS
// Supply Chain User Guide.
//
// The CSV file must be located in an Amazon S3 location accessible to AWS Supply
// Chain. It is recommended to use the same Amazon S3 bucket created during your
// AWS Supply Chain instance creation.
func supplychain_CreateBillOfMaterialsImportJob(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.CreateBillOfMaterialsImportJobInput{
		// InstanceId: *string, // Required
		// S3uri: *string, // Required
	}

	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}
	if len(_supplychainS3uri) > 0 {
		input.S3uri = aws.String(_supplychainS3uri)
	}
	if len(_supplychainClientToken) > 0 {
		input.ClientToken = aws.String(_supplychainClientToken)
	}

	if resp, err := client.CreateBillOfMaterialsImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables you to programmatically create a data pipeline to ingest data from
// source systems such as Amazon S3 buckets, to a predefined Amazon Web Services
// Supply Chain dataset (product, inbound_order) or a temporary dataset along with
// the data transformation query provided with the API.
func supplychain_CreateDataIntegrationFlow(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.CreateDataIntegrationFlowInput{
		// InstanceId: *string, // Required
		// Name: *string, // Required
		// Sources: []types.DataIntegrationFlowSource, // Required
		// Target: *types.DataIntegrationFlowTarget, // Required
		// Transformation: *types.DataIntegrationFlowTransformation, // Required
	}

	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}
	if len(_supplychainName) > 0 {
		input.Name = aws.String(_supplychainName)
	}
	if len(_supplychainSources) > 0 {
		if err := assignInputField(input, "Sources", _supplychainSources); err != nil {
			log.Errorf("invalid --sources: %s", err.Error())
			return
		}
	}
	if len(_supplychainTarget) > 0 {
		if err := assignInputField(input, "Target", _supplychainTarget); err != nil {
			log.Errorf("invalid --target: %s", err.Error())
			return
		}
	}
	if len(_supplychainTransformation) > 0 {
		if err := assignInputField(input, "Transformation", _supplychainTransformation); err != nil {
			log.Errorf("invalid --transformation: %s", err.Error())
			return
		}
	}
	if len(_supplychainTags) > 0 {
		if err := assignInputField(input, "Tags", _supplychainTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataIntegrationFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables you to programmatically create an Amazon Web Services Supply Chain data
// lake dataset. Developers can create the datasets using their pre-defined or
// custom schema for a given instance ID, namespace, and dataset name.
func supplychain_CreateDataLakeDataset(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.CreateDataLakeDatasetInput{
		// InstanceId: *string, // Required
		// Name: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}
	if len(_supplychainName) > 0 {
		input.Name = aws.String(_supplychainName)
	}
	if len(_supplychainNamespace) > 0 {
		input.Namespace = aws.String(_supplychainNamespace)
	}
	if len(_supplychainDescription) > 0 {
		input.Description = aws.String(_supplychainDescription)
	}
	if len(_supplychainPartitionSpec) > 0 {
		if err := assignInputField(input, "PartitionSpec", _supplychainPartitionSpec); err != nil {
			log.Errorf("invalid --partition-spec: %s", err.Error())
			return
		}
	}
	if len(_supplychainSchema) > 0 {
		if err := assignInputField(input, "Schema", _supplychainSchema); err != nil {
			log.Errorf("invalid --schema: %s", err.Error())
			return
		}
	}
	if len(_supplychainTags) > 0 {
		if err := assignInputField(input, "Tags", _supplychainTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataLakeDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables you to programmatically create an Amazon Web Services Supply Chain data
// lake namespace. Developers can create the namespaces for a given instance ID.
func supplychain_CreateDataLakeNamespace(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.CreateDataLakeNamespaceInput{
		// InstanceId: *string, // Required
		// Name: *string, // Required
	}

	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}
	if len(_supplychainName) > 0 {
		input.Name = aws.String(_supplychainName)
	}
	if len(_supplychainDescription) > 0 {
		input.Description = aws.String(_supplychainDescription)
	}
	if len(_supplychainTags) > 0 {
		if err := assignInputField(input, "Tags", _supplychainTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataLakeNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables you to programmatically create an Amazon Web Services Supply Chain
// instance by applying KMS keys and relevant information associated with the API
// without using the Amazon Web Services console.
//
// This is an asynchronous operation. Upon receiving a CreateInstance request,
// Amazon Web Services Supply Chain immediately returns the instance resource,
// instance ID, and the initializing state while simultaneously creating all
// required Amazon Web Services resources for an instance creation. You can use
// GetInstance to check the status of the instance. If the instance results in an
// unhealthy state, you need to check the error message, delete the current
// instance, and recreate a new one based on the mitigation from the error message.
func supplychain_CreateInstance(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.CreateInstanceInput{}

	if len(_supplychainClientToken) > 0 {
		input.ClientToken = aws.String(_supplychainClientToken)
	}
	if len(_supplychainInstanceDescription) > 0 {
		input.InstanceDescription = aws.String(_supplychainInstanceDescription)
	}
	if len(_supplychainInstanceName) > 0 {
		input.InstanceName = aws.String(_supplychainInstanceName)
	}
	if len(_supplychainKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_supplychainKmsKeyArn)
	}
	if len(_supplychainTags) > 0 {
		if err := assignInputField(input, "Tags", _supplychainTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_supplychainWebAppDnsDomain) > 0 {
		input.WebAppDnsDomain = aws.String(_supplychainWebAppDnsDomain)
	}

	if resp, err := client.CreateInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enable you to programmatically delete an existing data pipeline for the
// provided Amazon Web Services Supply Chain instance and DataIntegrationFlow name.
func supplychain_DeleteDataIntegrationFlow(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.DeleteDataIntegrationFlowInput{
		// InstanceId: *string, // Required
		// Name: *string, // Required
	}

	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}
	if len(_supplychainName) > 0 {
		input.Name = aws.String(_supplychainName)
	}

	if resp, err := client.DeleteDataIntegrationFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables you to programmatically delete an Amazon Web Services Supply Chain data
// lake dataset. Developers can delete the existing datasets for a given instance
// ID, namespace, and instance name.
func supplychain_DeleteDataLakeDataset(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.DeleteDataLakeDatasetInput{
		// InstanceId: *string, // Required
		// Name: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}
	if len(_supplychainName) > 0 {
		input.Name = aws.String(_supplychainName)
	}
	if len(_supplychainNamespace) > 0 {
		input.Namespace = aws.String(_supplychainNamespace)
	}

	if resp, err := client.DeleteDataLakeDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables you to programmatically delete an Amazon Web Services Supply Chain data
// lake namespace and its underling datasets. Developers can delete the existing
// namespaces for a given instance ID and namespace name.
func supplychain_DeleteDataLakeNamespace(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.DeleteDataLakeNamespaceInput{
		// InstanceId: *string, // Required
		// Name: *string, // Required
	}

	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}
	if len(_supplychainName) > 0 {
		input.Name = aws.String(_supplychainName)
	}

	if resp, err := client.DeleteDataLakeNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables you to programmatically delete an Amazon Web Services Supply Chain
// instance by deleting the KMS keys and relevant information associated with the
// API without using the Amazon Web Services console.
//
// This is an asynchronous operation. Upon receiving a DeleteInstance request,
// Amazon Web Services Supply Chain immediately returns a response with the
// instance resource, delete state while cleaning up all Amazon Web Services
// resources created during the instance creation process. You can use the
// GetInstance action to check the instance status.
func supplychain_DeleteInstance(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.DeleteInstanceInput{
		// InstanceId: *string, // Required
	}

	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}

	if resp, err := client.DeleteInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get status and details of a BillOfMaterialsImportJob.
func supplychain_GetBillOfMaterialsImportJob(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.GetBillOfMaterialsImportJobInput{
		// InstanceId: *string, // Required
		// JobId: *string, // Required
	}

	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}
	if len(_supplychainJobId) > 0 {
		input.JobId = aws.String(_supplychainJobId)
	}

	if resp, err := client.GetBillOfMaterialsImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables you to programmatically view an Amazon Web Services Supply Chain Data
// Integration Event. Developers can view the eventType, eventGroupId,
// eventTimestamp, datasetTarget, datasetLoadExecution.
func supplychain_GetDataIntegrationEvent(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.GetDataIntegrationEventInput{
		// EventId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_supplychainEventId) > 0 {
		input.EventId = aws.String(_supplychainEventId)
	}
	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}

	if resp, err := client.GetDataIntegrationEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables you to programmatically view a specific data pipeline for the provided
// Amazon Web Services Supply Chain instance and DataIntegrationFlow name.
func supplychain_GetDataIntegrationFlow(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.GetDataIntegrationFlowInput{
		// InstanceId: *string, // Required
		// Name: *string, // Required
	}

	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}
	if len(_supplychainName) > 0 {
		input.Name = aws.String(_supplychainName)
	}

	if resp, err := client.GetDataIntegrationFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the flow execution.
func supplychain_GetDataIntegrationFlowExecution(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.GetDataIntegrationFlowExecutionInput{
		// ExecutionId: *string, // Required
		// FlowName: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_supplychainExecutionId) > 0 {
		input.ExecutionId = aws.String(_supplychainExecutionId)
	}
	if len(_supplychainFlowName) > 0 {
		input.FlowName = aws.String(_supplychainFlowName)
	}
	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}

	if resp, err := client.GetDataIntegrationFlowExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables you to programmatically view an Amazon Web Services Supply Chain data
// lake dataset. Developers can view the data lake dataset information such as
// namespace, schema, and so on for a given instance ID, namespace, and dataset
// name.
func supplychain_GetDataLakeDataset(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.GetDataLakeDatasetInput{
		// InstanceId: *string, // Required
		// Name: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}
	if len(_supplychainName) > 0 {
		input.Name = aws.String(_supplychainName)
	}
	if len(_supplychainNamespace) > 0 {
		input.Namespace = aws.String(_supplychainNamespace)
	}

	if resp, err := client.GetDataLakeDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables you to programmatically view an Amazon Web Services Supply Chain data
// lake namespace. Developers can view the data lake namespace information such as
// description for a given instance ID and namespace name.
func supplychain_GetDataLakeNamespace(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.GetDataLakeNamespaceInput{
		// InstanceId: *string, // Required
		// Name: *string, // Required
	}

	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}
	if len(_supplychainName) > 0 {
		input.Name = aws.String(_supplychainName)
	}

	if resp, err := client.GetDataLakeNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables you to programmatically retrieve the information related to an Amazon
// Web Services Supply Chain instance ID.
func supplychain_GetInstance(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.GetInstanceInput{
		// InstanceId: *string, // Required
	}

	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}

	if resp, err := client.GetInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables you to programmatically list all data integration events for the
// provided Amazon Web Services Supply Chain instance.
func supplychain_ListDataIntegrationEvents(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.ListDataIntegrationEventsInput{
		// InstanceId: *string, // Required
	}

	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}
	if len(_supplychainEventType) > 0 {
		if err := assignInputField(input, "EventType", _supplychainEventType); err != nil {
			log.Errorf("invalid --event-type: %s", err.Error())
			return
		}
	}
	if len(_supplychainMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _supplychainMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_supplychainNextToken) > 0 {
		input.NextToken = aws.String(_supplychainNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataIntegrationEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*supplychain.ListDataIntegrationEventsOutput
	p := supplychain.NewListDataIntegrationEventsPaginator(client, input)
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

// List flow executions.
func supplychain_ListDataIntegrationFlowExecutions(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.ListDataIntegrationFlowExecutionsInput{
		// FlowName: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_supplychainFlowName) > 0 {
		input.FlowName = aws.String(_supplychainFlowName)
	}
	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}
	if len(_supplychainMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _supplychainMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_supplychainNextToken) > 0 {
		input.NextToken = aws.String(_supplychainNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataIntegrationFlowExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*supplychain.ListDataIntegrationFlowExecutionsOutput
	p := supplychain.NewListDataIntegrationFlowExecutionsPaginator(client, input)
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

// Enables you to programmatically list all data pipelines for the provided Amazon
// Web Services Supply Chain instance.
func supplychain_ListDataIntegrationFlows(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.ListDataIntegrationFlowsInput{
		// InstanceId: *string, // Required
	}

	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}
	if len(_supplychainMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _supplychainMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_supplychainNextToken) > 0 {
		input.NextToken = aws.String(_supplychainNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataIntegrationFlows(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*supplychain.ListDataIntegrationFlowsOutput
	p := supplychain.NewListDataIntegrationFlowsPaginator(client, input)
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

// Enables you to programmatically view the list of Amazon Web Services Supply
// Chain data lake datasets. Developers can view the datasets and the corresponding
// information such as namespace, schema, and so on for a given instance ID and
// namespace.
func supplychain_ListDataLakeDatasets(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.ListDataLakeDatasetsInput{
		// InstanceId: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}
	if len(_supplychainNamespace) > 0 {
		input.Namespace = aws.String(_supplychainNamespace)
	}
	if len(_supplychainMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _supplychainMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_supplychainNextToken) > 0 {
		input.NextToken = aws.String(_supplychainNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataLakeDatasets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*supplychain.ListDataLakeDatasetsOutput
	p := supplychain.NewListDataLakeDatasetsPaginator(client, input)
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

// Enables you to programmatically view the list of Amazon Web Services Supply
// Chain data lake namespaces. Developers can view the namespaces and the
// corresponding information such as description for a given instance ID. Note that
// this API only return custom namespaces, instance pre-defined namespaces are not
// included.
func supplychain_ListDataLakeNamespaces(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.ListDataLakeNamespacesInput{
		// InstanceId: *string, // Required
	}

	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}
	if len(_supplychainMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _supplychainMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_supplychainNextToken) > 0 {
		input.NextToken = aws.String(_supplychainNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataLakeNamespaces(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*supplychain.ListDataLakeNamespacesOutput
	p := supplychain.NewListDataLakeNamespacesPaginator(client, input)
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

// List all Amazon Web Services Supply Chain instances for a specific account.
// Enables you to programmatically list all Amazon Web Services Supply Chain
// instances based on their account ID, instance name, and state of the instance
// (active or delete).
func supplychain_ListInstances(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.ListInstancesInput{}

	if len(_supplychainInstanceNameFilter) > 0 {
		input.InstanceNameFilter = append([]string(nil), _supplychainInstanceNameFilter...)
	}
	if len(_supplychainInstanceStateFilter) > 0 {
		if err := assignInputField(input, "InstanceStateFilter", _supplychainInstanceStateFilter); err != nil {
			log.Errorf("invalid --instance-state-filter: %s", err.Error())
			return
		}
	}
	if len(_supplychainMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _supplychainMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_supplychainNextToken) > 0 {
		input.NextToken = aws.String(_supplychainNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*supplychain.ListInstancesOutput
	p := supplychain.NewListInstancesPaginator(client, input)
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

// List all the tags for an Amazon Web ServicesSupply Chain resource. You can list
// all the tags added to a resource. By listing the tags, developers can view the
// tag level information on a resource and perform actions such as, deleting a
// resource associated with a particular tag.
func supplychain_ListTagsForResource(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_supplychainResourceArn) > 0 {
		input.ResourceArn = aws.String(_supplychainResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Send the data payload for the event with real-time data for analysis or
// monitoring. The real-time data events are stored in an Amazon Web Services
// service before being processed and stored in data lake.
func supplychain_SendDataIntegrationEvent(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.SendDataIntegrationEventInput{
		// Data: *string, // Required
		// EventGroupId: *string, // Required
		// EventType: types.DataIntegrationEventType, // Required
		// InstanceId: *string, // Required
	}

	if len(_supplychainData) > 0 {
		input.Data = aws.String(_supplychainData)
	}
	if len(_supplychainEventGroupId) > 0 {
		input.EventGroupId = aws.String(_supplychainEventGroupId)
	}
	if len(_supplychainEventType) > 0 {
		if err := assignInputField(input, "EventType", _supplychainEventType); err != nil {
			log.Errorf("invalid --event-type: %s", err.Error())
			return
		}
	}
	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}
	if len(_supplychainClientToken) > 0 {
		input.ClientToken = aws.String(_supplychainClientToken)
	}
	if len(_supplychainDatasetTarget) > 0 {
		if err := assignInputField(input, "DatasetTarget", _supplychainDatasetTarget); err != nil {
			log.Errorf("invalid --dataset-target: %s", err.Error())
			return
		}
	}
	if len(_supplychainEventTimestamp) > 0 {
		if err := assignInputField(input, "EventTimestamp", _supplychainEventTimestamp); err != nil {
			log.Errorf("invalid --event-timestamp: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendDataIntegrationEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// You can create tags during or after creating a resource such as instance, data
// flow, or dataset in AWS Supply chain. During the data ingestion process, you can
// add tags such as dev, test, or prod to data flows created during the data
// ingestion process in the AWS Supply Chain datasets. You can use these tags to
// identify a group of resources or a single resource used by the developer.
func supplychain_TagResource(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_supplychainResourceArn) > 0 {
		input.ResourceArn = aws.String(_supplychainResourceArn)
	}
	if len(_supplychainTags) > 0 {
		if err := assignInputField(input, "Tags", _supplychainTags); err != nil {
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

// You can delete tags for an Amazon Web Services Supply chain resource such as
// instance, data flow, or dataset in AWS Supply Chain. During the data ingestion
// process, you can delete tags such as dev, test, or prod to data flows created
// during the data ingestion process in the AWS Supply Chain datasets.
func supplychain_UntagResource(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_supplychainResourceArn) > 0 {
		input.ResourceArn = aws.String(_supplychainResourceArn)
	}
	if len(_supplychainTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _supplychainTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables you to programmatically update an existing data pipeline to ingest data
// from the source systems such as, Amazon S3 buckets, to a predefined Amazon Web
// Services Supply Chain dataset (product, inbound_order) or a temporary dataset
// along with the data transformation query provided with the API.
func supplychain_UpdateDataIntegrationFlow(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.UpdateDataIntegrationFlowInput{
		// InstanceId: *string, // Required
		// Name: *string, // Required
	}

	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}
	if len(_supplychainName) > 0 {
		input.Name = aws.String(_supplychainName)
	}
	if len(_supplychainSources) > 0 {
		if err := assignInputField(input, "Sources", _supplychainSources); err != nil {
			log.Errorf("invalid --sources: %s", err.Error())
			return
		}
	}
	if len(_supplychainTarget) > 0 {
		if err := assignInputField(input, "Target", _supplychainTarget); err != nil {
			log.Errorf("invalid --target: %s", err.Error())
			return
		}
	}
	if len(_supplychainTransformation) > 0 {
		if err := assignInputField(input, "Transformation", _supplychainTransformation); err != nil {
			log.Errorf("invalid --transformation: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDataIntegrationFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables you to programmatically update an Amazon Web Services Supply Chain data
// lake dataset. Developers can update the description of a data lake dataset for a
// given instance ID, namespace, and dataset name.
func supplychain_UpdateDataLakeDataset(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.UpdateDataLakeDatasetInput{
		// InstanceId: *string, // Required
		// Name: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}
	if len(_supplychainName) > 0 {
		input.Name = aws.String(_supplychainName)
	}
	if len(_supplychainNamespace) > 0 {
		input.Namespace = aws.String(_supplychainNamespace)
	}
	if len(_supplychainDescription) > 0 {
		input.Description = aws.String(_supplychainDescription)
	}

	if resp, err := client.UpdateDataLakeDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables you to programmatically update an Amazon Web Services Supply Chain data
// lake namespace. Developers can update the description of a data lake namespace
// for a given instance ID and namespace name.
func supplychain_UpdateDataLakeNamespace(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.UpdateDataLakeNamespaceInput{
		// InstanceId: *string, // Required
		// Name: *string, // Required
	}

	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}
	if len(_supplychainName) > 0 {
		input.Name = aws.String(_supplychainName)
	}
	if len(_supplychainDescription) > 0 {
		input.Description = aws.String(_supplychainDescription)
	}

	if resp, err := client.UpdateDataLakeNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables you to programmatically update an Amazon Web Services Supply Chain
// instance description by providing all the relevant information such as account
// ID, instance ID and so on without using the AWS console.
func supplychain_UpdateInstance(cfg aws.Config, client *supplychain.Client) {
	input := &supplychain.UpdateInstanceInput{
		// InstanceId: *string, // Required
	}

	if len(_supplychainInstanceId) > 0 {
		input.InstanceId = aws.String(_supplychainInstanceId)
	}
	if len(_supplychainInstanceDescription) > 0 {
		input.InstanceDescription = aws.String(_supplychainInstanceDescription)
	}
	if len(_supplychainInstanceName) > 0 {
		input.InstanceName = aws.String(_supplychainInstanceName)
	}

	if resp, err := client.UpdateInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_supplychainCmd)
	_supplychainCmd.Flags().SortFlags = false

	_supplychainCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_supplychainCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_supplychainCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_supplychainCmd.Flags().StringVarP(&_supplychainClientToken, "client-token", "", "", "Client Token")
	_supplychainCmd.Flags().StringVarP(&_supplychainData, "data", "", "", "Data")
	_supplychainCmd.Flags().StringVarP(&_supplychainDatasetTarget, "dataset-target", "", "", "Dataset Target")
	_supplychainCmd.Flags().StringVarP(&_supplychainDescription, "description", "", "", "Description")
	_supplychainCmd.Flags().StringVarP(&_supplychainEventGroupId, "event-group-id", "", "", "Event Group ID")
	_supplychainCmd.Flags().StringVarP(&_supplychainEventId, "event-id", "", "", "Event ID")
	_supplychainCmd.Flags().StringVarP(&_supplychainEventTimestamp, "event-timestamp", "", "", "Event Timestamp")
	_supplychainCmd.Flags().StringVarP(&_supplychainEventType, "event-type", "", "", "Event Type")
	_supplychainCmd.Flags().StringVarP(&_supplychainExecutionId, "execution-id", "", "", "Execution ID")
	_supplychainCmd.Flags().StringVarP(&_supplychainFlowName, "flow-name", "", "", "Flow Name")
	_supplychainCmd.Flags().StringVarP(&_supplychainInstanceDescription, "instance-description", "", "", "Instance Description")
	_supplychainCmd.Flags().StringVarP(&_supplychainInstanceId, "instance-id", "", "", "Instance ID")
	_supplychainCmd.Flags().StringVarP(&_supplychainInstanceName, "instance-name", "", "", "Instance Name")
	_supplychainCmd.Flags().StringSliceVarP(&_supplychainInstanceNameFilter, "instance-name-filter", "", nil, "Instance Name Filter")
	_supplychainCmd.Flags().StringVarP(&_supplychainInstanceStateFilter, "instance-state-filter", "", "", "Instance State Filter")
	_supplychainCmd.Flags().StringVarP(&_supplychainJobId, "job-id", "", "", "Job ID")
	_supplychainCmd.Flags().StringVarP(&_supplychainKmsKeyArn, "kms-key-arn", "", "", "KMS Key ARN")
	_supplychainCmd.Flags().StringVarP(&_supplychainMaxResults, "max-results", "", "", "Max Results")
	_supplychainCmd.Flags().StringVarP(&_supplychainName, "name", "", "", "Name")
	_supplychainCmd.Flags().StringVarP(&_supplychainNamespace, "namespace", "", "", "Namespace")
	_supplychainCmd.Flags().StringVarP(&_supplychainNextToken, "next-token", "", "", "Next Token")
	_supplychainCmd.Flags().StringVarP(&_supplychainPartitionSpec, "partition-spec", "", "", "Partition Spec")
	_supplychainCmd.Flags().StringVarP(&_supplychainResourceArn, "resource-arn", "", "", "Resource ARN")
	_supplychainCmd.Flags().StringVarP(&_supplychainS3uri, "s3uri", "", "", "S3uri")
	_supplychainCmd.Flags().StringVarP(&_supplychainSchema, "schema", "", "", "Schema")
	_supplychainCmd.Flags().StringVarP(&_supplychainSources, "sources", "", "", "Sources")
	_supplychainCmd.Flags().StringSliceVarP(&_supplychainTagKeys, "tag-keys", "", nil, "Tag Keys")
	_supplychainCmd.Flags().StringVarP(&_supplychainTags, "tags", "", "", "Tags")
	_supplychainCmd.Flags().StringVarP(&_supplychainTarget, "target", "", "", "Target")
	_supplychainCmd.Flags().StringVarP(&_supplychainTransformation, "transformation", "", "", "Transformation")
	_supplychainCmd.Flags().StringVarP(&_supplychainWebAppDnsDomain, "web-app-dns-domain", "", "", "Web App DNS Domain")

	_supplychainCmd.Flags().BoolVarP(&_supplychainCreateBillOfMaterialsImportJob, "create-bill-of-materials-import-job", "", false, "Create Bill Of Materials Import Job")
	_supplychainCmd.Flags().BoolVarP(&_supplychainCreateDataIntegrationFlow, "create-data-integration-flow", "", false, "Create Data Integration Flow")
	_supplychainCmd.Flags().BoolVarP(&_supplychainCreateDataLakeDataset, "create-data-lake-dataset", "", false, "Create Data Lake Dataset")
	_supplychainCmd.Flags().BoolVarP(&_supplychainCreateDataLakeNamespace, "create-data-lake-namespace", "", false, "Create Data Lake Namespace")
	_supplychainCmd.Flags().BoolVarP(&_supplychainCreateInstance, "create-instance", "", false, "Create Instance")
	_supplychainCmd.Flags().BoolVarP(&_supplychainDeleteDataIntegrationFlow, "delete-data-integration-flow", "", false, "Delete Data Integration Flow")
	_supplychainCmd.Flags().BoolVarP(&_supplychainDeleteDataLakeDataset, "delete-data-lake-dataset", "", false, "Delete Data Lake Dataset")
	_supplychainCmd.Flags().BoolVarP(&_supplychainDeleteDataLakeNamespace, "delete-data-lake-namespace", "", false, "Delete Data Lake Namespace")
	_supplychainCmd.Flags().BoolVarP(&_supplychainDeleteInstance, "delete-instance", "", false, "Delete Instance")
	_supplychainCmd.Flags().BoolVarP(&_supplychainGetBillOfMaterialsImportJob, "get-bill-of-materials-import-job", "", false, "Get Bill Of Materials Import Job")
	_supplychainCmd.Flags().BoolVarP(&_supplychainGetDataIntegrationEvent, "get-data-integration-event", "", false, "Get Data Integration Event")
	_supplychainCmd.Flags().BoolVarP(&_supplychainGetDataIntegrationFlow, "get-data-integration-flow", "", false, "Get Data Integration Flow")
	_supplychainCmd.Flags().BoolVarP(&_supplychainGetDataIntegrationFlowExecution, "get-data-integration-flow-execution", "", false, "Get Data Integration Flow Execution")
	_supplychainCmd.Flags().BoolVarP(&_supplychainGetDataLakeDataset, "get-data-lake-dataset", "", false, "Get Data Lake Dataset")
	_supplychainCmd.Flags().BoolVarP(&_supplychainGetDataLakeNamespace, "get-data-lake-namespace", "", false, "Get Data Lake Namespace")
	_supplychainCmd.Flags().BoolVarP(&_supplychainGetInstance, "get-instance", "", false, "Get Instance")
	_supplychainCmd.Flags().BoolVarP(&_supplychainListDataIntegrationEvents, "list-data-integration-events", "", false, "List Data Integration Events")
	_supplychainCmd.Flags().BoolVarP(&_supplychainListDataIntegrationFlowExecutions, "list-data-integration-flow-executions", "", false, "List Data Integration Flow Executions")
	_supplychainCmd.Flags().BoolVarP(&_supplychainListDataIntegrationFlows, "list-data-integration-flows", "", false, "List Data Integration Flows")
	_supplychainCmd.Flags().BoolVarP(&_supplychainListDataLakeDatasets, "list-data-lake-datasets", "", false, "List Data Lake Datasets")
	_supplychainCmd.Flags().BoolVarP(&_supplychainListDataLakeNamespaces, "list-data-lake-namespaces", "", false, "List Data Lake Namespaces")
	_supplychainCmd.Flags().BoolVarP(&_supplychainListInstances, "list-instances", "", false, "List Instances")
	_supplychainCmd.Flags().BoolVarP(&_supplychainListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_supplychainCmd.Flags().BoolVarP(&_supplychainSendDataIntegrationEvent, "send-data-integration-event", "", false, "Send Data Integration Event")
	_supplychainCmd.Flags().BoolVarP(&_supplychainTagResource, "tag-resource", "", false, "Tag Resource")
	_supplychainCmd.Flags().BoolVarP(&_supplychainUntagResource, "untag-resource", "", false, "Untag Resource")
	_supplychainCmd.Flags().BoolVarP(&_supplychainUpdateDataIntegrationFlow, "update-data-integration-flow", "", false, "Update Data Integration Flow")
	_supplychainCmd.Flags().BoolVarP(&_supplychainUpdateDataLakeDataset, "update-data-lake-dataset", "", false, "Update Data Lake Dataset")
	_supplychainCmd.Flags().BoolVarP(&_supplychainUpdateDataLakeNamespace, "update-data-lake-namespace", "", false, "Update Data Lake Namespace")
	_supplychainCmd.Flags().BoolVarP(&_supplychainUpdateInstance, "update-instance", "", false, "Update Instance")

}
