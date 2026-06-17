package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptunegraph"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// neptunegraphCmd represents the neptunegraph command
var _neptunegraphCmd = &cobra.Command{
	Use:   "neptunegraph",
	Short: "AWS neptunegraph CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := neptunegraph.NewFromConfig(cfg)
		if _neptunegraphCancelExportTask {
			neptunegraph_CancelExportTask(cfg, client)
			return
		}
		if _neptunegraphCancelImportTask {
			neptunegraph_CancelImportTask(cfg, client)
			return
		}
		if _neptunegraphCancelQuery {
			neptunegraph_CancelQuery(cfg, client)
			return
		}
		if _neptunegraphCreateGraph {
			neptunegraph_CreateGraph(cfg, client)
			return
		}
		if _neptunegraphCreateGraphSnapshot {
			neptunegraph_CreateGraphSnapshot(cfg, client)
			return
		}
		if _neptunegraphCreateGraphUsingImportTask {
			neptunegraph_CreateGraphUsingImportTask(cfg, client)
			return
		}
		if _neptunegraphCreatePrivateGraphEndpoint {
			neptunegraph_CreatePrivateGraphEndpoint(cfg, client)
			return
		}
		if _neptunegraphDeleteGraph {
			neptunegraph_DeleteGraph(cfg, client)
			return
		}
		if _neptunegraphDeleteGraphSnapshot {
			neptunegraph_DeleteGraphSnapshot(cfg, client)
			return
		}
		if _neptunegraphDeletePrivateGraphEndpoint {
			neptunegraph_DeletePrivateGraphEndpoint(cfg, client)
			return
		}
		if _neptunegraphExecuteQuery {
			neptunegraph_ExecuteQuery(cfg, client)
			return
		}
		if _neptunegraphGetExportTask {
			neptunegraph_GetExportTask(cfg, client)
			return
		}
		if _neptunegraphGetGraph {
			neptunegraph_GetGraph(cfg, client)
			return
		}
		if _neptunegraphGetGraphSnapshot {
			neptunegraph_GetGraphSnapshot(cfg, client)
			return
		}
		if _neptunegraphGetGraphSummary {
			neptunegraph_GetGraphSummary(cfg, client)
			return
		}
		if _neptunegraphGetImportTask {
			neptunegraph_GetImportTask(cfg, client)
			return
		}
		if _neptunegraphGetPrivateGraphEndpoint {
			neptunegraph_GetPrivateGraphEndpoint(cfg, client)
			return
		}
		if _neptunegraphGetQuery {
			neptunegraph_GetQuery(cfg, client)
			return
		}
		if _neptunegraphListExportTasks {
			neptunegraph_ListExportTasks(cfg, client)
			return
		}
		if _neptunegraphListGraphSnapshots {
			neptunegraph_ListGraphSnapshots(cfg, client)
			return
		}
		if _neptunegraphListGraphs {
			neptunegraph_ListGraphs(cfg, client)
			return
		}
		if _neptunegraphListImportTasks {
			neptunegraph_ListImportTasks(cfg, client)
			return
		}
		if _neptunegraphListPrivateGraphEndpoints {
			neptunegraph_ListPrivateGraphEndpoints(cfg, client)
			return
		}
		if _neptunegraphListQueries {
			neptunegraph_ListQueries(cfg, client)
			return
		}
		if _neptunegraphListTagsForResource {
			neptunegraph_ListTagsForResource(cfg, client)
			return
		}
		if _neptunegraphResetGraph {
			neptunegraph_ResetGraph(cfg, client)
			return
		}
		if _neptunegraphRestoreGraphFromSnapshot {
			neptunegraph_RestoreGraphFromSnapshot(cfg, client)
			return
		}
		if _neptunegraphStartExportTask {
			neptunegraph_StartExportTask(cfg, client)
			return
		}
		if _neptunegraphStartGraph {
			neptunegraph_StartGraph(cfg, client)
			return
		}
		if _neptunegraphStartImportTask {
			neptunegraph_StartImportTask(cfg, client)
			return
		}
		if _neptunegraphStopGraph {
			neptunegraph_StopGraph(cfg, client)
			return
		}
		if _neptunegraphTagResource {
			neptunegraph_TagResource(cfg, client)
			return
		}
		if _neptunegraphUntagResource {
			neptunegraph_UntagResource(cfg, client)
			return
		}
		if _neptunegraphUpdateGraph {
			neptunegraph_UpdateGraph(cfg, client)
			return
		}

	},
}

var (
	_neptunegraphCancelExportTask           bool
	_neptunegraphCancelImportTask           bool
	_neptunegraphCancelQuery                bool
	_neptunegraphCreateGraph                bool
	_neptunegraphCreateGraphSnapshot        bool
	_neptunegraphCreateGraphUsingImportTask bool
	_neptunegraphCreatePrivateGraphEndpoint bool
	_neptunegraphDeleteGraph                bool
	_neptunegraphDeleteGraphSnapshot        bool
	_neptunegraphDeletePrivateGraphEndpoint bool
	_neptunegraphExecuteQuery               bool
	_neptunegraphGetExportTask              bool
	_neptunegraphGetGraph                   bool
	_neptunegraphGetGraphSnapshot           bool
	_neptunegraphGetGraphSummary            bool
	_neptunegraphGetImportTask              bool
	_neptunegraphGetPrivateGraphEndpoint    bool
	_neptunegraphGetQuery                   bool
	_neptunegraphListExportTasks            bool
	_neptunegraphListGraphSnapshots         bool
	_neptunegraphListGraphs                 bool
	_neptunegraphListImportTasks            bool
	_neptunegraphListPrivateGraphEndpoints  bool
	_neptunegraphListQueries                bool
	_neptunegraphListTagsForResource        bool
	_neptunegraphResetGraph                 bool
	_neptunegraphRestoreGraphFromSnapshot   bool
	_neptunegraphStartExportTask            bool
	_neptunegraphStartGraph                 bool
	_neptunegraphStartImportTask            bool
	_neptunegraphStopGraph                  bool
	_neptunegraphTagResource                bool
	_neptunegraphUntagResource              bool
	_neptunegraphUpdateGraph                bool

	_neptunegraphBlankNodeHandling         string
	_neptunegraphDeletionProtection        string
	_neptunegraphDestination               string
	_neptunegraphExplainMode               string
	_neptunegraphExportFilter              string
	_neptunegraphFailOnError               string
	_neptunegraphFormat                    string
	_neptunegraphGraphIdentifier           string
	_neptunegraphGraphName                 string
	_neptunegraphImportOptions             string
	_neptunegraphKmsKeyIdentifier          string
	_neptunegraphLanguage                  string
	_neptunegraphMaxProvisionedMemory      string
	_neptunegraphMaxResults                string
	_neptunegraphMinProvisionedMemory      string
	_neptunegraphMode                      string
	_neptunegraphNextToken                 string
	_neptunegraphParameters                string
	_neptunegraphParquetType               string
	_neptunegraphPlanCache                 string
	_neptunegraphProvisionedMemory         string
	_neptunegraphPublicConnectivity        string
	_neptunegraphQueryId                   string
	_neptunegraphQueryString               string
	_neptunegraphQueryTimeoutMilliseconds  string
	_neptunegraphReplicaCount              string
	_neptunegraphResourceArn               string
	_neptunegraphRoleArn                   string
	_neptunegraphSkipSnapshot              string
	_neptunegraphSnapshotIdentifier        string
	_neptunegraphSnapshotName              string
	_neptunegraphSource                    string
	_neptunegraphState                     string
	_neptunegraphSubnetIds                 []string
	_neptunegraphTagKeys                   []string
	_neptunegraphTags                      string
	_neptunegraphTaskIdentifier            string
	_neptunegraphVectorSearchConfiguration string
	_neptunegraphVpcId                     string
	_neptunegraphVpcSecurityGroupIds       []string
)

// Cancel the specified export task.
func neptunegraph_CancelExportTask(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.CancelExportTaskInput{
		// TaskIdentifier: *string, // Required
	}

	if len(_neptunegraphTaskIdentifier) > 0 {
		input.TaskIdentifier = aws.String(_neptunegraphTaskIdentifier)
	}

	if resp, err := client.CancelExportTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified import task.
func neptunegraph_CancelImportTask(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.CancelImportTaskInput{
		// TaskIdentifier: *string, // Required
	}

	if len(_neptunegraphTaskIdentifier) > 0 {
		input.TaskIdentifier = aws.String(_neptunegraphTaskIdentifier)
	}

	if resp, err := client.CancelImportTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a specified query.
func neptunegraph_CancelQuery(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.CancelQueryInput{
		// GraphIdentifier: *string, // Required
		// QueryId: *string, // Required
	}

	if len(_neptunegraphGraphIdentifier) > 0 {
		input.GraphIdentifier = aws.String(_neptunegraphGraphIdentifier)
	}
	if len(_neptunegraphQueryId) > 0 {
		input.QueryId = aws.String(_neptunegraphQueryId)
	}

	if resp, err := client.CancelQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Neptune Analytics graph.
func neptunegraph_CreateGraph(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.CreateGraphInput{
		// GraphName: *string, // Required
		// ProvisionedMemory: *int32, // Required
	}

	if len(_neptunegraphGraphName) > 0 {
		input.GraphName = aws.String(_neptunegraphGraphName)
	}
	if len(_neptunegraphProvisionedMemory) > 0 {
		if err := assignInputField(input, "ProvisionedMemory", _neptunegraphProvisionedMemory); err != nil {
			log.Errorf("invalid --provisioned-memory: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _neptunegraphDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphKmsKeyIdentifier) > 0 {
		input.KmsKeyIdentifier = aws.String(_neptunegraphKmsKeyIdentifier)
	}
	if len(_neptunegraphPublicConnectivity) > 0 {
		if err := assignInputField(input, "PublicConnectivity", _neptunegraphPublicConnectivity); err != nil {
			log.Errorf("invalid --public-connectivity: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphReplicaCount) > 0 {
		if err := assignInputField(input, "ReplicaCount", _neptunegraphReplicaCount); err != nil {
			log.Errorf("invalid --replica-count: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphTags) > 0 {
		if err := assignInputField(input, "Tags", _neptunegraphTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphVectorSearchConfiguration) > 0 {
		if err := assignInputField(input, "VectorSearchConfiguration", _neptunegraphVectorSearchConfiguration); err != nil {
			log.Errorf("invalid --vector-search-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGraph(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a snapshot of the specific graph.
func neptunegraph_CreateGraphSnapshot(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.CreateGraphSnapshotInput{
		// GraphIdentifier: *string, // Required
		// SnapshotName: *string, // Required
	}

	if len(_neptunegraphGraphIdentifier) > 0 {
		input.GraphIdentifier = aws.String(_neptunegraphGraphIdentifier)
	}
	if len(_neptunegraphSnapshotName) > 0 {
		input.SnapshotName = aws.String(_neptunegraphSnapshotName)
	}
	if len(_neptunegraphTags) > 0 {
		if err := assignInputField(input, "Tags", _neptunegraphTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGraphSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Neptune Analytics graph and imports data into it, either from
// Amazon Simple Storage Service (S3) or from a Neptune database or a Neptune
// database snapshot.
//
// The data can be loaded from files in S3 that in either the [Gremlin CSV format] or the [openCypher load format].
//
// [Gremlin CSV format]: https://docs.aws.amazon.com/neptune/latest/userguide/bulk-load-tutorial-format-gremlin.html
// [openCypher load format]: https://docs.aws.amazon.com/neptune/latest/userguide/bulk-load-tutorial-format-opencypher.html
func neptunegraph_CreateGraphUsingImportTask(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.CreateGraphUsingImportTaskInput{
		// GraphName: *string, // Required
		// RoleArn: *string, // Required
		// Source: *string, // Required
	}

	if len(_neptunegraphGraphName) > 0 {
		input.GraphName = aws.String(_neptunegraphGraphName)
	}
	if len(_neptunegraphRoleArn) > 0 {
		input.RoleArn = aws.String(_neptunegraphRoleArn)
	}
	if len(_neptunegraphSource) > 0 {
		input.Source = aws.String(_neptunegraphSource)
	}
	if len(_neptunegraphBlankNodeHandling) > 0 {
		if err := assignInputField(input, "BlankNodeHandling", _neptunegraphBlankNodeHandling); err != nil {
			log.Errorf("invalid --blank-node-handling: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _neptunegraphDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphFailOnError) > 0 {
		if err := assignInputField(input, "FailOnError", _neptunegraphFailOnError); err != nil {
			log.Errorf("invalid --fail-on-error: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphFormat) > 0 {
		if err := assignInputField(input, "Format", _neptunegraphFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphImportOptions) > 0 {
		if err := assignInputField(input, "ImportOptions", _neptunegraphImportOptions); err != nil {
			log.Errorf("invalid --import-options: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphKmsKeyIdentifier) > 0 {
		input.KmsKeyIdentifier = aws.String(_neptunegraphKmsKeyIdentifier)
	}
	if len(_neptunegraphMaxProvisionedMemory) > 0 {
		if err := assignInputField(input, "MaxProvisionedMemory", _neptunegraphMaxProvisionedMemory); err != nil {
			log.Errorf("invalid --max-provisioned-memory: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphMinProvisionedMemory) > 0 {
		if err := assignInputField(input, "MinProvisionedMemory", _neptunegraphMinProvisionedMemory); err != nil {
			log.Errorf("invalid --min-provisioned-memory: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphParquetType) > 0 {
		if err := assignInputField(input, "ParquetType", _neptunegraphParquetType); err != nil {
			log.Errorf("invalid --parquet-type: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphPublicConnectivity) > 0 {
		if err := assignInputField(input, "PublicConnectivity", _neptunegraphPublicConnectivity); err != nil {
			log.Errorf("invalid --public-connectivity: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphReplicaCount) > 0 {
		if err := assignInputField(input, "ReplicaCount", _neptunegraphReplicaCount); err != nil {
			log.Errorf("invalid --replica-count: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphTags) > 0 {
		if err := assignInputField(input, "Tags", _neptunegraphTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphVectorSearchConfiguration) > 0 {
		if err := assignInputField(input, "VectorSearchConfiguration", _neptunegraphVectorSearchConfiguration); err != nil {
			log.Errorf("invalid --vector-search-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGraphUsingImportTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a private graph endpoint to allow private access to the graph from
// within a VPC. You can attach security groups to the private graph endpoint.
//
// VPC endpoint charges apply.
func neptunegraph_CreatePrivateGraphEndpoint(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.CreatePrivateGraphEndpointInput{
		// GraphIdentifier: *string, // Required
	}

	if len(_neptunegraphGraphIdentifier) > 0 {
		input.GraphIdentifier = aws.String(_neptunegraphGraphIdentifier)
	}
	if len(_neptunegraphSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _neptunegraphSubnetIds...)
	}
	if len(_neptunegraphVpcId) > 0 {
		input.VpcId = aws.String(_neptunegraphVpcId)
	}
	if len(_neptunegraphVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _neptunegraphVpcSecurityGroupIds...)
	}

	if resp, err := client.CreatePrivateGraphEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified graph. Graphs cannot be deleted if delete-protection is
// enabled.
func neptunegraph_DeleteGraph(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.DeleteGraphInput{
		// GraphIdentifier: *string, // Required
		// SkipSnapshot: *bool, // Required
	}

	if len(_neptunegraphGraphIdentifier) > 0 {
		input.GraphIdentifier = aws.String(_neptunegraphGraphIdentifier)
	}
	if len(_neptunegraphSkipSnapshot) > 0 {
		if err := assignInputField(input, "SkipSnapshot", _neptunegraphSkipSnapshot); err != nil {
			log.Errorf("invalid --skip-snapshot: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteGraph(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified graph snapshot.
func neptunegraph_DeleteGraphSnapshot(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.DeleteGraphSnapshotInput{
		// SnapshotIdentifier: *string, // Required
	}

	if len(_neptunegraphSnapshotIdentifier) > 0 {
		input.SnapshotIdentifier = aws.String(_neptunegraphSnapshotIdentifier)
	}

	if resp, err := client.DeleteGraphSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a private graph endpoint.
func neptunegraph_DeletePrivateGraphEndpoint(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.DeletePrivateGraphEndpointInput{
		// GraphIdentifier: *string, // Required
		// VpcId: *string, // Required
	}

	if len(_neptunegraphGraphIdentifier) > 0 {
		input.GraphIdentifier = aws.String(_neptunegraphGraphIdentifier)
	}
	if len(_neptunegraphVpcId) > 0 {
		input.VpcId = aws.String(_neptunegraphVpcId)
	}

	if resp, err := client.DeletePrivateGraphEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Execute an openCypher query.
// When invoking this operation in a Neptune Analytics cluster, the IAM user or
// role making the request must have a policy attached that allows one of the
// following IAM actions in that cluster, depending on the query:
//
// - neptune-graph:ReadDataViaQuery
//
// - neptune-graph:WriteDataViaQuery
//
// - neptune-graph:DeleteDataViaQuery
func neptunegraph_ExecuteQuery(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.ExecuteQueryInput{
		// GraphIdentifier: *string, // Required
		// Language: types.QueryLanguage, // Required
		// QueryString: *string, // Required
	}

	if len(_neptunegraphGraphIdentifier) > 0 {
		input.GraphIdentifier = aws.String(_neptunegraphGraphIdentifier)
	}
	if len(_neptunegraphLanguage) > 0 {
		if err := assignInputField(input, "Language", _neptunegraphLanguage); err != nil {
			log.Errorf("invalid --language: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphQueryString) > 0 {
		input.QueryString = aws.String(_neptunegraphQueryString)
	}
	if len(_neptunegraphExplainMode) > 0 {
		if err := assignInputField(input, "ExplainMode", _neptunegraphExplainMode); err != nil {
			log.Errorf("invalid --explain-mode: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphParameters) > 0 {
		if err := assignInputField(input, "Parameters", _neptunegraphParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphPlanCache) > 0 {
		if err := assignInputField(input, "PlanCache", _neptunegraphPlanCache); err != nil {
			log.Errorf("invalid --plan-cache: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphQueryTimeoutMilliseconds) > 0 {
		if err := assignInputField(input, "QueryTimeoutMilliseconds", _neptunegraphQueryTimeoutMilliseconds); err != nil {
			log.Errorf("invalid --query-timeout-milliseconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.ExecuteQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a specified export task.
func neptunegraph_GetExportTask(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.GetExportTaskInput{
		// TaskIdentifier: *string, // Required
	}

	if len(_neptunegraphTaskIdentifier) > 0 {
		input.TaskIdentifier = aws.String(_neptunegraphTaskIdentifier)
	}

	if resp, err := client.GetExportTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specified graph.
func neptunegraph_GetGraph(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.GetGraphInput{
		// GraphIdentifier: *string, // Required
	}

	if len(_neptunegraphGraphIdentifier) > 0 {
		input.GraphIdentifier = aws.String(_neptunegraphGraphIdentifier)
	}

	if resp, err := client.GetGraph(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a specified graph snapshot.
func neptunegraph_GetGraphSnapshot(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.GetGraphSnapshotInput{
		// SnapshotIdentifier: *string, // Required
	}

	if len(_neptunegraphSnapshotIdentifier) > 0 {
		input.SnapshotIdentifier = aws.String(_neptunegraphSnapshotIdentifier)
	}

	if resp, err := client.GetGraphSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a graph summary for a property graph.
func neptunegraph_GetGraphSummary(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.GetGraphSummaryInput{
		// GraphIdentifier: *string, // Required
	}

	if len(_neptunegraphGraphIdentifier) > 0 {
		input.GraphIdentifier = aws.String(_neptunegraphGraphIdentifier)
	}
	if len(_neptunegraphMode) > 0 {
		if err := assignInputField(input, "Mode", _neptunegraphMode); err != nil {
			log.Errorf("invalid --mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetGraphSummary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a specified import task.
func neptunegraph_GetImportTask(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.GetImportTaskInput{
		// TaskIdentifier: *string, // Required
	}

	if len(_neptunegraphTaskIdentifier) > 0 {
		input.TaskIdentifier = aws.String(_neptunegraphTaskIdentifier)
	}

	if resp, err := client.GetImportTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a specified private endpoint.
func neptunegraph_GetPrivateGraphEndpoint(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.GetPrivateGraphEndpointInput{
		// GraphIdentifier: *string, // Required
		// VpcId: *string, // Required
	}

	if len(_neptunegraphGraphIdentifier) > 0 {
		input.GraphIdentifier = aws.String(_neptunegraphGraphIdentifier)
	}
	if len(_neptunegraphVpcId) > 0 {
		input.VpcId = aws.String(_neptunegraphVpcId)
	}

	if resp, err := client.GetPrivateGraphEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the status of a specified query.
// When invoking this operation in a Neptune Analytics cluster, the IAM user or
// role making the request must have the neptune-graph:GetQueryStatus IAM action
// attached.
func neptunegraph_GetQuery(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.GetQueryInput{
		// GraphIdentifier: *string, // Required
		// QueryId: *string, // Required
	}

	if len(_neptunegraphGraphIdentifier) > 0 {
		input.GraphIdentifier = aws.String(_neptunegraphGraphIdentifier)
	}
	if len(_neptunegraphQueryId) > 0 {
		input.QueryId = aws.String(_neptunegraphQueryId)
	}

	if resp, err := client.GetQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of export tasks.
func neptunegraph_ListExportTasks(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.ListExportTasksInput{}

	if len(_neptunegraphGraphIdentifier) > 0 {
		input.GraphIdentifier = aws.String(_neptunegraphGraphIdentifier)
	}
	if len(_neptunegraphMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _neptunegraphMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphNextToken) > 0 {
		input.NextToken = aws.String(_neptunegraphNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListExportTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*neptunegraph.ListExportTasksOutput
	p := neptunegraph.NewListExportTasksPaginator(client, input)
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

// Lists available snapshots of a specified Neptune Analytics graph.
func neptunegraph_ListGraphSnapshots(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.ListGraphSnapshotsInput{}

	if len(_neptunegraphGraphIdentifier) > 0 {
		input.GraphIdentifier = aws.String(_neptunegraphGraphIdentifier)
	}
	if len(_neptunegraphMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _neptunegraphMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphNextToken) > 0 {
		input.NextToken = aws.String(_neptunegraphNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGraphSnapshots(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*neptunegraph.ListGraphSnapshotsOutput
	p := neptunegraph.NewListGraphSnapshotsPaginator(client, input)
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

// Lists available Neptune Analytics graphs.
func neptunegraph_ListGraphs(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.ListGraphsInput{}

	if len(_neptunegraphMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _neptunegraphMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphNextToken) > 0 {
		input.NextToken = aws.String(_neptunegraphNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGraphs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*neptunegraph.ListGraphsOutput
	p := neptunegraph.NewListGraphsPaginator(client, input)
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

// Lists import tasks.
func neptunegraph_ListImportTasks(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.ListImportTasksInput{}

	if len(_neptunegraphMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _neptunegraphMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphNextToken) > 0 {
		input.NextToken = aws.String(_neptunegraphNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListImportTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*neptunegraph.ListImportTasksOutput
	p := neptunegraph.NewListImportTasksPaginator(client, input)
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

// Lists private endpoints for a specified Neptune Analytics graph.
func neptunegraph_ListPrivateGraphEndpoints(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.ListPrivateGraphEndpointsInput{
		// GraphIdentifier: *string, // Required
	}

	if len(_neptunegraphGraphIdentifier) > 0 {
		input.GraphIdentifier = aws.String(_neptunegraphGraphIdentifier)
	}
	if len(_neptunegraphMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _neptunegraphMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphNextToken) > 0 {
		input.NextToken = aws.String(_neptunegraphNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPrivateGraphEndpoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*neptunegraph.ListPrivateGraphEndpointsOutput
	p := neptunegraph.NewListPrivateGraphEndpointsPaginator(client, input)
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

// Lists active openCypher queries.
func neptunegraph_ListQueries(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.ListQueriesInput{
		// GraphIdentifier: *string, // Required
		// MaxResults: *int32, // Required
	}

	if len(_neptunegraphGraphIdentifier) > 0 {
		input.GraphIdentifier = aws.String(_neptunegraphGraphIdentifier)
	}
	if len(_neptunegraphMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _neptunegraphMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphState) > 0 {
		if err := assignInputField(input, "State", _neptunegraphState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListQueries(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists tags associated with a specified resource.
func neptunegraph_ListTagsForResource(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_neptunegraphResourceArn) > 0 {
		input.ResourceArn = aws.String(_neptunegraphResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Empties the data from a specified Neptune Analytics graph.
func neptunegraph_ResetGraph(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.ResetGraphInput{
		// GraphIdentifier: *string, // Required
		// SkipSnapshot: *bool, // Required
	}

	if len(_neptunegraphGraphIdentifier) > 0 {
		input.GraphIdentifier = aws.String(_neptunegraphGraphIdentifier)
	}
	if len(_neptunegraphSkipSnapshot) > 0 {
		if err := assignInputField(input, "SkipSnapshot", _neptunegraphSkipSnapshot); err != nil {
			log.Errorf("invalid --skip-snapshot: %s", err.Error())
			return
		}
	}

	if resp, err := client.ResetGraph(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restores a graph from a snapshot.
func neptunegraph_RestoreGraphFromSnapshot(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.RestoreGraphFromSnapshotInput{
		// GraphName: *string, // Required
		// SnapshotIdentifier: *string, // Required
	}

	if len(_neptunegraphGraphName) > 0 {
		input.GraphName = aws.String(_neptunegraphGraphName)
	}
	if len(_neptunegraphSnapshotIdentifier) > 0 {
		input.SnapshotIdentifier = aws.String(_neptunegraphSnapshotIdentifier)
	}
	if len(_neptunegraphDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _neptunegraphDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphProvisionedMemory) > 0 {
		if err := assignInputField(input, "ProvisionedMemory", _neptunegraphProvisionedMemory); err != nil {
			log.Errorf("invalid --provisioned-memory: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphPublicConnectivity) > 0 {
		if err := assignInputField(input, "PublicConnectivity", _neptunegraphPublicConnectivity); err != nil {
			log.Errorf("invalid --public-connectivity: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphReplicaCount) > 0 {
		if err := assignInputField(input, "ReplicaCount", _neptunegraphReplicaCount); err != nil {
			log.Errorf("invalid --replica-count: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphTags) > 0 {
		if err := assignInputField(input, "Tags", _neptunegraphTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.RestoreGraphFromSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Export data from an existing Neptune Analytics graph to Amazon S3. The graph
// state should be AVAILABLE .
func neptunegraph_StartExportTask(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.StartExportTaskInput{
		// Destination: *string, // Required
		// Format: types.ExportFormat, // Required
		// GraphIdentifier: *string, // Required
		// KmsKeyIdentifier: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_neptunegraphDestination) > 0 {
		input.Destination = aws.String(_neptunegraphDestination)
	}
	if len(_neptunegraphFormat) > 0 {
		if err := assignInputField(input, "Format", _neptunegraphFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphGraphIdentifier) > 0 {
		input.GraphIdentifier = aws.String(_neptunegraphGraphIdentifier)
	}
	if len(_neptunegraphKmsKeyIdentifier) > 0 {
		input.KmsKeyIdentifier = aws.String(_neptunegraphKmsKeyIdentifier)
	}
	if len(_neptunegraphRoleArn) > 0 {
		input.RoleArn = aws.String(_neptunegraphRoleArn)
	}
	if len(_neptunegraphExportFilter) > 0 {
		if err := assignInputField(input, "ExportFilter", _neptunegraphExportFilter); err != nil {
			log.Errorf("invalid --export-filter: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphParquetType) > 0 {
		if err := assignInputField(input, "ParquetType", _neptunegraphParquetType); err != nil {
			log.Errorf("invalid --parquet-type: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphTags) > 0 {
		if err := assignInputField(input, "Tags", _neptunegraphTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
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

// Starts the specific graph.
func neptunegraph_StartGraph(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.StartGraphInput{
		// GraphIdentifier: *string, // Required
	}

	if len(_neptunegraphGraphIdentifier) > 0 {
		input.GraphIdentifier = aws.String(_neptunegraphGraphIdentifier)
	}

	if resp, err := client.StartGraph(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Import data into existing Neptune Analytics graph from Amazon Simple Storage
// Service (S3). The graph needs to be empty and in the AVAILABLE state.
func neptunegraph_StartImportTask(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.StartImportTaskInput{
		// GraphIdentifier: *string, // Required
		// RoleArn: *string, // Required
		// Source: *string, // Required
	}

	if len(_neptunegraphGraphIdentifier) > 0 {
		input.GraphIdentifier = aws.String(_neptunegraphGraphIdentifier)
	}
	if len(_neptunegraphRoleArn) > 0 {
		input.RoleArn = aws.String(_neptunegraphRoleArn)
	}
	if len(_neptunegraphSource) > 0 {
		input.Source = aws.String(_neptunegraphSource)
	}
	if len(_neptunegraphBlankNodeHandling) > 0 {
		if err := assignInputField(input, "BlankNodeHandling", _neptunegraphBlankNodeHandling); err != nil {
			log.Errorf("invalid --blank-node-handling: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphFailOnError) > 0 {
		if err := assignInputField(input, "FailOnError", _neptunegraphFailOnError); err != nil {
			log.Errorf("invalid --fail-on-error: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphFormat) > 0 {
		if err := assignInputField(input, "Format", _neptunegraphFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphImportOptions) > 0 {
		if err := assignInputField(input, "ImportOptions", _neptunegraphImportOptions); err != nil {
			log.Errorf("invalid --import-options: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphParquetType) > 0 {
		if err := assignInputField(input, "ParquetType", _neptunegraphParquetType); err != nil {
			log.Errorf("invalid --parquet-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartImportTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the specific graph.
func neptunegraph_StopGraph(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.StopGraphInput{
		// GraphIdentifier: *string, // Required
	}

	if len(_neptunegraphGraphIdentifier) > 0 {
		input.GraphIdentifier = aws.String(_neptunegraphGraphIdentifier)
	}

	if resp, err := client.StopGraph(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds tags to the specified resource.
func neptunegraph_TagResource(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_neptunegraphResourceArn) > 0 {
		input.ResourceArn = aws.String(_neptunegraphResourceArn)
	}
	if len(_neptunegraphTags) > 0 {
		if err := assignInputField(input, "Tags", _neptunegraphTags); err != nil {
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
func neptunegraph_UntagResource(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_neptunegraphResourceArn) > 0 {
		input.ResourceArn = aws.String(_neptunegraphResourceArn)
	}
	if len(_neptunegraphTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _neptunegraphTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of a specified Neptune Analytics graph
func neptunegraph_UpdateGraph(cfg aws.Config, client *neptunegraph.Client) {
	input := &neptunegraph.UpdateGraphInput{
		// GraphIdentifier: *string, // Required
	}

	if len(_neptunegraphGraphIdentifier) > 0 {
		input.GraphIdentifier = aws.String(_neptunegraphGraphIdentifier)
	}
	if len(_neptunegraphDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _neptunegraphDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphProvisionedMemory) > 0 {
		if err := assignInputField(input, "ProvisionedMemory", _neptunegraphProvisionedMemory); err != nil {
			log.Errorf("invalid --provisioned-memory: %s", err.Error())
			return
		}
	}
	if len(_neptunegraphPublicConnectivity) > 0 {
		if err := assignInputField(input, "PublicConnectivity", _neptunegraphPublicConnectivity); err != nil {
			log.Errorf("invalid --public-connectivity: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateGraph(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_neptunegraphCmd)
	_neptunegraphCmd.Flags().SortFlags = false

	_neptunegraphCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_neptunegraphCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_neptunegraphCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphBlankNodeHandling, "blank-node-handling", "", "", "Blank Node Handling")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphDeletionProtection, "deletion-protection", "", "", "Deletion Protection")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphDestination, "destination", "", "", "Destination")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphExplainMode, "explain-mode", "", "", "Explain Mode")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphExportFilter, "export-filter", "", "", "Export Filter")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphFailOnError, "fail-on-error", "", "", "Fail On Error")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphFormat, "format", "", "", "Format")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphGraphIdentifier, "graph-identifier", "", "", "Graph Identifier")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphGraphName, "graph-name", "", "", "Graph Name")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphImportOptions, "import-options", "", "", "Import Options")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphKmsKeyIdentifier, "kms-key-identifier", "", "", "KMS Key Identifier")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphLanguage, "language", "", "", "Language")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphMaxProvisionedMemory, "max-provisioned-memory", "", "", "Max Provisioned Memory")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphMaxResults, "max-results", "", "", "Max Results")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphMinProvisionedMemory, "min-provisioned-memory", "", "", "Min Provisioned Memory")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphMode, "mode", "", "", "Mode")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphNextToken, "next-token", "", "", "Next Token")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphParameters, "parameters", "", "", "Parameters")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphParquetType, "parquet-type", "", "", "Parquet Type")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphPlanCache, "plan-cache", "", "", "Plan Cache")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphProvisionedMemory, "provisioned-memory", "", "", "Provisioned Memory")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphPublicConnectivity, "public-connectivity", "", "", "Public Connectivity")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphQueryId, "query-id", "", "", "Query ID")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphQueryString, "query-string", "", "", "Query String")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphQueryTimeoutMilliseconds, "query-timeout-milliseconds", "", "", "Query Timeout Milliseconds")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphReplicaCount, "replica-count", "", "", "Replica Count")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphResourceArn, "resource-arn", "", "", "Resource ARN")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphRoleArn, "role-arn", "", "", "Role ARN")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphSkipSnapshot, "skip-snapshot", "", "", "Skip Snapshot")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphSnapshotIdentifier, "snapshot-identifier", "", "", "Snapshot Identifier")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphSnapshotName, "snapshot-name", "", "", "Snapshot Name")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphSource, "source", "", "", "Source")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphState, "state", "", "", "State")
	_neptunegraphCmd.Flags().StringSliceVarP(&_neptunegraphSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_neptunegraphCmd.Flags().StringSliceVarP(&_neptunegraphTagKeys, "tag-keys", "", nil, "Tag Keys")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphTags, "tags", "", "", "Tags")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphTaskIdentifier, "task-identifier", "", "", "Task Identifier")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphVectorSearchConfiguration, "vector-search-configuration", "", "", "Vector Search Configuration")
	_neptunegraphCmd.Flags().StringVarP(&_neptunegraphVpcId, "vpc-id", "", "", "VPC ID")
	_neptunegraphCmd.Flags().StringSliceVarP(&_neptunegraphVpcSecurityGroupIds, "vpc-security-group-ids", "", nil, "VPC Security Group Ids")

	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphCancelExportTask, "cancel-export-task", "", false, "Cancel Export Task")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphCancelImportTask, "cancel-import-task", "", false, "Cancel Import Task")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphCancelQuery, "cancel-query", "", false, "Cancel Query")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphCreateGraph, "create-graph", "", false, "Create Graph")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphCreateGraphSnapshot, "create-graph-snapshot", "", false, "Create Graph Snapshot")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphCreateGraphUsingImportTask, "create-graph-using-import-task", "", false, "Create Graph Using Import Task")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphCreatePrivateGraphEndpoint, "create-private-graph-endpoint", "", false, "Create Private Graph Endpoint")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphDeleteGraph, "delete-graph", "", false, "Delete Graph")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphDeleteGraphSnapshot, "delete-graph-snapshot", "", false, "Delete Graph Snapshot")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphDeletePrivateGraphEndpoint, "delete-private-graph-endpoint", "", false, "Delete Private Graph Endpoint")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphExecuteQuery, "execute-query", "", false, "Execute Query")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphGetExportTask, "get-export-task", "", false, "Get Export Task")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphGetGraph, "get-graph", "", false, "Get Graph")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphGetGraphSnapshot, "get-graph-snapshot", "", false, "Get Graph Snapshot")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphGetGraphSummary, "get-graph-summary", "", false, "Get Graph Summary")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphGetImportTask, "get-import-task", "", false, "Get Import Task")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphGetPrivateGraphEndpoint, "get-private-graph-endpoint", "", false, "Get Private Graph Endpoint")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphGetQuery, "get-query", "", false, "Get Query")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphListExportTasks, "list-export-tasks", "", false, "List Export Tasks")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphListGraphSnapshots, "list-graph-snapshots", "", false, "List Graph Snapshots")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphListGraphs, "list-graphs", "", false, "List Graphs")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphListImportTasks, "list-import-tasks", "", false, "List Import Tasks")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphListPrivateGraphEndpoints, "list-private-graph-endpoints", "", false, "List Private Graph Endpoints")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphListQueries, "list-queries", "", false, "List Queries")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphResetGraph, "reset-graph", "", false, "Reset Graph")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphRestoreGraphFromSnapshot, "restore-graph-from-snapshot", "", false, "Restore Graph From Snapshot")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphStartExportTask, "start-export-task", "", false, "Start Export Task")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphStartGraph, "start-graph", "", false, "Start Graph")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphStartImportTask, "start-import-task", "", false, "Start Import Task")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphStopGraph, "stop-graph", "", false, "Stop Graph")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphTagResource, "tag-resource", "", false, "Tag Resource")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphUntagResource, "untag-resource", "", false, "Untag Resource")
	_neptunegraphCmd.Flags().BoolVarP(&_neptunegraphUpdateGraph, "update-graph", "", false, "Update Graph")

}
