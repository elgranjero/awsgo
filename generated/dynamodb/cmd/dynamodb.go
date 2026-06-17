package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// dynamodbCmd represents the dynamodb command
var _dynamodbCmd = &cobra.Command{
	Use:   "dynamodb",
	Short: "AWS dynamodb CLI",
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
		client := dynamodb.NewFromConfig(cfg)
		if _dynamodbBatchExecuteStatement {
			dynamodb_BatchExecuteStatement(cfg, client)
			return
		}
		if _dynamodbBatchGetItem {
			dynamodb_BatchGetItem(cfg, client)
			return
		}
		if _dynamodbBatchWriteItem {
			dynamodb_BatchWriteItem(cfg, client)
			return
		}
		if _dynamodbCreateBackup {
			dynamodb_CreateBackup(cfg, client)
			return
		}
		if _dynamodbCreateGlobalTable {
			dynamodb_CreateGlobalTable(cfg, client)
			return
		}
		if _dynamodbCreateTable {
			dynamodb_CreateTable(cfg, client)
			return
		}
		if _dynamodbDeleteBackup {
			dynamodb_DeleteBackup(cfg, client)
			return
		}
		if _dynamodbDeleteItem {
			dynamodb_DeleteItem(cfg, client)
			return
		}
		if _dynamodbDeleteResourcePolicy {
			dynamodb_DeleteResourcePolicy(cfg, client)
			return
		}
		if _dynamodbDeleteTable {
			dynamodb_DeleteTable(cfg, client)
			return
		}
		if _dynamodbDescribeBackup {
			dynamodb_DescribeBackup(cfg, client)
			return
		}
		if _dynamodbDescribeContinuousBackups {
			dynamodb_DescribeContinuousBackups(cfg, client)
			return
		}
		if _dynamodbDescribeContributorInsights {
			dynamodb_DescribeContributorInsights(cfg, client)
			return
		}
		if _dynamodbDescribeEndpoints {
			dynamodb_DescribeEndpoints(cfg, client)
			return
		}
		if _dynamodbDescribeExport {
			dynamodb_DescribeExport(cfg, client)
			return
		}
		if _dynamodbDescribeGlobalTable {
			dynamodb_DescribeGlobalTable(cfg, client)
			return
		}
		if _dynamodbDescribeGlobalTableSettings {
			dynamodb_DescribeGlobalTableSettings(cfg, client)
			return
		}
		if _dynamodbDescribeImport {
			dynamodb_DescribeImport(cfg, client)
			return
		}
		if _dynamodbDescribeKinesisStreamingDestination {
			dynamodb_DescribeKinesisStreamingDestination(cfg, client)
			return
		}
		if _dynamodbDescribeLimits {
			dynamodb_DescribeLimits(cfg, client)
			return
		}
		if _dynamodbDescribeTable {
			dynamodb_DescribeTable(cfg, client)
			return
		}
		if _dynamodbDescribeTableReplicaAutoScaling {
			dynamodb_DescribeTableReplicaAutoScaling(cfg, client)
			return
		}
		if _dynamodbDescribeTimeToLive {
			dynamodb_DescribeTimeToLive(cfg, client)
			return
		}
		if _dynamodbDisableKinesisStreamingDestination {
			dynamodb_DisableKinesisStreamingDestination(cfg, client)
			return
		}
		if _dynamodbEnableKinesisStreamingDestination {
			dynamodb_EnableKinesisStreamingDestination(cfg, client)
			return
		}
		if _dynamodbExecuteStatement {
			dynamodb_ExecuteStatement(cfg, client)
			return
		}
		if _dynamodbExecuteTransaction {
			dynamodb_ExecuteTransaction(cfg, client)
			return
		}
		if _dynamodbExportTableToPointInTime {
			dynamodb_ExportTableToPointInTime(cfg, client)
			return
		}
		if _dynamodbGetItem {
			dynamodb_GetItem(cfg, client)
			return
		}
		if _dynamodbGetResourcePolicy {
			dynamodb_GetResourcePolicy(cfg, client)
			return
		}
		if _dynamodbImportTable {
			dynamodb_ImportTable(cfg, client)
			return
		}
		if _dynamodbListBackups {
			dynamodb_ListBackups(cfg, client)
			return
		}
		if _dynamodbListContributorInsights {
			dynamodb_ListContributorInsights(cfg, client)
			return
		}
		if _dynamodbListExports {
			dynamodb_ListExports(cfg, client)
			return
		}
		if _dynamodbListGlobalTables {
			dynamodb_ListGlobalTables(cfg, client)
			return
		}
		if _dynamodbListImports {
			dynamodb_ListImports(cfg, client)
			return
		}
		if _dynamodbListTables {
			dynamodb_ListTables(cfg, client)
			return
		}
		if _dynamodbListTagsOfResource {
			dynamodb_ListTagsOfResource(cfg, client)
			return
		}
		if _dynamodbPutItem {
			dynamodb_PutItem(cfg, client)
			return
		}
		if _dynamodbPutResourcePolicy {
			dynamodb_PutResourcePolicy(cfg, client)
			return
		}
		if _dynamodbQuery {
			dynamodb_Query(cfg, client)
			return
		}
		if _dynamodbRestoreTableFromBackup {
			dynamodb_RestoreTableFromBackup(cfg, client)
			return
		}
		if _dynamodbRestoreTableToPointInTime {
			dynamodb_RestoreTableToPointInTime(cfg, client)
			return
		}
		if _dynamodbScan {
			dynamodb_Scan(cfg, client)
			return
		}
		if _dynamodbTagResource {
			dynamodb_TagResource(cfg, client)
			return
		}
		if _dynamodbTransactGetItems {
			dynamodb_TransactGetItems(cfg, client)
			return
		}
		if _dynamodbTransactWriteItems {
			dynamodb_TransactWriteItems(cfg, client)
			return
		}
		if _dynamodbUntagResource {
			dynamodb_UntagResource(cfg, client)
			return
		}
		if _dynamodbUpdateContinuousBackups {
			dynamodb_UpdateContinuousBackups(cfg, client)
			return
		}
		if _dynamodbUpdateContributorInsights {
			dynamodb_UpdateContributorInsights(cfg, client)
			return
		}
		if _dynamodbUpdateGlobalTable {
			dynamodb_UpdateGlobalTable(cfg, client)
			return
		}
		if _dynamodbUpdateGlobalTableSettings {
			dynamodb_UpdateGlobalTableSettings(cfg, client)
			return
		}
		if _dynamodbUpdateItem {
			dynamodb_UpdateItem(cfg, client)
			return
		}
		if _dynamodbUpdateKinesisStreamingDestination {
			dynamodb_UpdateKinesisStreamingDestination(cfg, client)
			return
		}
		if _dynamodbUpdateTable {
			dynamodb_UpdateTable(cfg, client)
			return
		}
		if _dynamodbUpdateTableReplicaAutoScaling {
			dynamodb_UpdateTableReplicaAutoScaling(cfg, client)
			return
		}
		if _dynamodbUpdateTimeToLive {
			dynamodb_UpdateTimeToLive(cfg, client)
			return
		}

	},
}

var (
	_dynamodbBatchExecuteStatement               bool
	_dynamodbBatchGetItem                        bool
	_dynamodbBatchWriteItem                      bool
	_dynamodbCreateBackup                        bool
	_dynamodbCreateGlobalTable                   bool
	_dynamodbCreateTable                         bool
	_dynamodbDeleteBackup                        bool
	_dynamodbDeleteItem                          bool
	_dynamodbDeleteResourcePolicy                bool
	_dynamodbDeleteTable                         bool
	_dynamodbDescribeBackup                      bool
	_dynamodbDescribeContinuousBackups           bool
	_dynamodbDescribeContributorInsights         bool
	_dynamodbDescribeEndpoints                   bool
	_dynamodbDescribeExport                      bool
	_dynamodbDescribeGlobalTable                 bool
	_dynamodbDescribeGlobalTableSettings         bool
	_dynamodbDescribeImport                      bool
	_dynamodbDescribeKinesisStreamingDestination bool
	_dynamodbDescribeLimits                      bool
	_dynamodbDescribeTable                       bool
	_dynamodbDescribeTableReplicaAutoScaling     bool
	_dynamodbDescribeTimeToLive                  bool
	_dynamodbDisableKinesisStreamingDestination  bool
	_dynamodbEnableKinesisStreamingDestination   bool
	_dynamodbExecuteStatement                    bool
	_dynamodbExecuteTransaction                  bool
	_dynamodbExportTableToPointInTime            bool
	_dynamodbGetItem                             bool
	_dynamodbGetResourcePolicy                   bool
	_dynamodbImportTable                         bool
	_dynamodbListBackups                         bool
	_dynamodbListContributorInsights             bool
	_dynamodbListExports                         bool
	_dynamodbListGlobalTables                    bool
	_dynamodbListImports                         bool
	_dynamodbListTables                          bool
	_dynamodbListTagsOfResource                  bool
	_dynamodbPutItem                             bool
	_dynamodbPutResourcePolicy                   bool
	_dynamodbQuery                               bool
	_dynamodbRestoreTableFromBackup              bool
	_dynamodbRestoreTableToPointInTime           bool
	_dynamodbScan                                bool
	_dynamodbTagResource                         bool
	_dynamodbTransactGetItems                    bool
	_dynamodbTransactWriteItems                  bool
	_dynamodbUntagResource                       bool
	_dynamodbUpdateContinuousBackups             bool
	_dynamodbUpdateContributorInsights           bool
	_dynamodbUpdateGlobalTable                   bool
	_dynamodbUpdateGlobalTableSettings           bool
	_dynamodbUpdateItem                          bool
	_dynamodbUpdateKinesisStreamingDestination   bool
	_dynamodbUpdateTable                         bool
	_dynamodbUpdateTableReplicaAutoScaling       bool
	_dynamodbUpdateTimeToLive                    bool

	_dynamodbAttributeDefinitions                                         string
	_dynamodbAttributeUpdates                                             string
	_dynamodbAttributesToGet                                              []string
	_dynamodbBackupArn                                                    string
	_dynamodbBackupName                                                   string
	_dynamodbBackupType                                                   string
	_dynamodbBillingMode                                                  string
	_dynamodbBillingModeOverride                                          string
	_dynamodbClientRequestToken                                           string
	_dynamodbClientToken                                                  string
	_dynamodbConditionExpression                                          string
	_dynamodbConditionalOperator                                          string
	_dynamodbConfirmRemoveSelfResourceAccess                              string
	_dynamodbConsistentRead                                               string
	_dynamodbContributorInsightsAction                                    string
	_dynamodbContributorInsightsMode                                      string
	_dynamodbDeletionProtectionEnabled                                    string
	_dynamodbEnableKinesisStreamingConfiguration                          string
	_dynamodbExclusiveStartBackupArn                                      string
	_dynamodbExclusiveStartGlobalTableName                                string
	_dynamodbExclusiveStartKey                                            string
	_dynamodbExclusiveStartTableName                                      string
	_dynamodbExpected                                                     string
	_dynamodbExpectedRevisionId                                           string
	_dynamodbExportArn                                                    string
	_dynamodbExportFormat                                                 string
	_dynamodbExportTime                                                   string
	_dynamodbExportType                                                   string
	_dynamodbExpressionAttributeNames                                     string
	_dynamodbExpressionAttributeValues                                    string
	_dynamodbFilterExpression                                             string
	_dynamodbGlobalSecondaryIndexOverride                                 string
	_dynamodbGlobalSecondaryIndexUpdates                                  string
	_dynamodbGlobalSecondaryIndexes                                       string
	_dynamodbGlobalTableBillingMode                                       string
	_dynamodbGlobalTableGlobalSecondaryIndexSettingsUpdate                string
	_dynamodbGlobalTableName                                              string
	_dynamodbGlobalTableProvisionedWriteCapacityAutoScalingSettingsUpdate string
	_dynamodbGlobalTableProvisionedWriteCapacityUnits                     string
	_dynamodbGlobalTableSettingsReplicationMode                           string
	_dynamodbGlobalTableSourceArn                                         string
	_dynamodbGlobalTableWitnessUpdates                                    string
	_dynamodbImportArn                                                    string
	_dynamodbIncrementalExportSpecification                               string
	_dynamodbIndexName                                                    string
	_dynamodbInputCompressionType                                         string
	_dynamodbInputFormat                                                  string
	_dynamodbInputFormatOptions                                           string
	_dynamodbItem                                                         string
	_dynamodbKey                                                          string
	_dynamodbKeyConditionExpression                                       string
	_dynamodbKeyConditions                                                string
	_dynamodbKeySchema                                                    string
	_dynamodbLimit                                                        string
	_dynamodbLocalSecondaryIndexOverride                                  string
	_dynamodbLocalSecondaryIndexes                                        string
	_dynamodbMaxResults                                                   string
	_dynamodbMultiRegionConsistency                                       string
	_dynamodbNextToken                                                    string
	_dynamodbOnDemandThroughput                                           string
	_dynamodbOnDemandThroughputOverride                                   string
	_dynamodbPageSize                                                     string
	_dynamodbParameters                                                   string
	_dynamodbPointInTimeRecoverySpecification                             string
	_dynamodbPolicy                                                       string
	_dynamodbProjectionExpression                                         string
	_dynamodbProvisionedThroughput                                        string
	_dynamodbProvisionedThroughputOverride                                string
	_dynamodbProvisionedWriteCapacityAutoScalingUpdate                    string
	_dynamodbQueryFilter                                                  string
	_dynamodbRegionName                                                   string
	_dynamodbReplicaSettingsUpdate                                        string
	_dynamodbReplicaUpdates                                               string
	_dynamodbReplicationGroup                                             string
	_dynamodbRequestItems                                                 string
	_dynamodbResourceArn                                                  string
	_dynamodbResourcePolicy                                               string
	_dynamodbRestoreDateTime                                              string
	_dynamodbReturnConsumedCapacity                                       string
	_dynamodbReturnItemCollectionMetrics                                  string
	_dynamodbReturnValues                                                 string
	_dynamodbReturnValuesOnConditionCheckFailure                          string
	_dynamodbS3Bucket                                                     string
	_dynamodbS3BucketOwner                                                string
	_dynamodbS3BucketSource                                               string
	_dynamodbS3Prefix                                                     string
	_dynamodbS3SseAlgorithm                                               string
	_dynamodbS3SseKmsKeyId                                                string
	_dynamodbScanFilter                                                   string
	_dynamodbScanIndexForward                                             string
	_dynamodbSegment                                                      string
	_dynamodbSelect                                                       string
	_dynamodbSourceTableArn                                               string
	_dynamodbSourceTableName                                              string
	_dynamodbSSESpecification                                             string
	_dynamodbSSESpecificationOverride                                     string
	_dynamodbStatement                                                    string
	_dynamodbStatements                                                   string
	_dynamodbStreamArn                                                    string
	_dynamodbStreamSpecification                                          string
	_dynamodbTableArn                                                     string
	_dynamodbTableClass                                                   string
	_dynamodbTableCreationParameters                                      string
	_dynamodbTableName                                                    string
	_dynamodbTagKeys                                                      []string
	_dynamodbTags                                                         string
	_dynamodbTargetTableName                                              string
	_dynamodbTimeRangeLowerBound                                          string
	_dynamodbTimeRangeUpperBound                                          string
	_dynamodbTimeToLiveSpecification                                      string
	_dynamodbTotalSegments                                                string
	_dynamodbTransactItems                                                string
	_dynamodbTransactStatements                                           string
	_dynamodbUpdateExpression                                             string
	_dynamodbUpdateKinesisStreamingConfiguration                          string
	_dynamodbUseLatestRestorableTime                                      string
	_dynamodbWarmThroughput                                               string
)

// This operation allows you to perform batch reads or writes on data stored in
// DynamoDB, using PartiQL. Each read statement in a BatchExecuteStatement must
// specify an equality condition on all key attributes. This enforces that each
// SELECT statement in a batch returns at most a single item. For more information,
// see [Running batch operations with PartiQL for DynamoDB].
//
// The entire batch must consist of either read statements or write statements,
// you cannot mix both in one batch.
//
// A HTTP 200 response does not mean that all statements in the
// BatchExecuteStatement succeeded. Error details for individual statements can be
// found under the [Error]field of the BatchStatementResponse for each statement.
//
// [Error]: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_BatchStatementResponse.html#DDB-Type-BatchStatementResponse-Error
// [Running batch operations with PartiQL for DynamoDB]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ql-reference.multiplestatements.batching.html
func dynamodb_BatchExecuteStatement(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.BatchExecuteStatementInput{
		// Statements: []types.BatchStatementRequest, // Required
	}

	if len(_dynamodbStatements) > 0 {
		if err := assignInputField(input, "Statements", _dynamodbStatements); err != nil {
			log.Errorf("invalid --statements: %s", err.Error())
			return
		}
	}
	if len(_dynamodbReturnConsumedCapacity) > 0 {
		if err := assignInputField(input, "ReturnConsumedCapacity", _dynamodbReturnConsumedCapacity); err != nil {
			log.Errorf("invalid --return-consumed-capacity: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchExecuteStatement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The BatchGetItem operation returns the attributes of one or more items from one
// or more tables. You identify requested items by primary key.
//
// A single operation can retrieve up to 16 MB of data, which can contain as many
// as 100 items. BatchGetItem returns a partial result if the response size limit
// is exceeded, the table's provisioned throughput is exceeded, more than 1MB per
// partition is requested, or an internal processing failure occurs. If a partial
// result is returned, the operation returns a value for UnprocessedKeys . You can
// use this value to retry the operation starting with the next item to get.
//
// If you request more than 100 items, BatchGetItem returns a ValidationException
// with the message "Too many items requested for the BatchGetItem call."
//
// For example, if you ask to retrieve 100 items, but each individual item is 300
// KB in size, the system returns 52 items (so as not to exceed the 16 MB limit).
// It also returns an appropriate UnprocessedKeys value so you can get the next
// page of results. If desired, your application can include its own logic to
// assemble the pages of results into one dataset.
//
// If none of the items can be processed due to insufficient provisioned
// throughput on all of the tables in the request, then BatchGetItem returns a
// ProvisionedThroughputExceededException . If at least one of the items is
// successfully processed, then BatchGetItem completes successfully, while
// returning the keys of the unread items in UnprocessedKeys .
//
// If DynamoDB returns any unprocessed items, you should retry the batch operation
// on those items. However, we strongly recommend that you use an exponential
// backoff algorithm. If you retry the batch operation immediately, the underlying
// read or write requests can still fail due to throttling on the individual
// tables. If you delay the batch operation using exponential backoff, the
// individual requests in the batch are much more likely to succeed.
//
// For more information, see [Batch Operations and Error Handling] in the Amazon DynamoDB Developer Guide.
//
// By default, BatchGetItem performs eventually consistent reads on every table in
// the request. If you want strongly consistent reads instead, you can set
// ConsistentRead to true for any or all tables.
//
// In order to minimize response latency, BatchGetItem may retrieve items in
// parallel.
//
// When designing your application, keep in mind that DynamoDB does not return
// items in any particular order. To help parse the response by item, include the
// primary key values for the items in your request in the ProjectionExpression
// parameter.
//
// If a requested item does not exist, it is not returned in the result. Requests
// for nonexistent items consume the minimum read capacity units according to the
// type of read. For more information, see [Working with Tables]in the Amazon DynamoDB Developer Guide.
//
// BatchGetItem will result in a ValidationException if the same key is specified
// multiple times.
//
// [Batch Operations and Error Handling]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ErrorHandling.html#BatchOperations
// [Working with Tables]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#CapacityUnitCalculations
func dynamodb_BatchGetItem(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.BatchGetItemInput{
		// RequestItems: map[string]types.KeysAndAttributes, // Required
	}

	if len(_dynamodbRequestItems) > 0 {
		if err := assignInputField(input, "RequestItems", _dynamodbRequestItems); err != nil {
			log.Errorf("invalid --request-items: %s", err.Error())
			return
		}
	}
	if len(_dynamodbReturnConsumedCapacity) > 0 {
		if err := assignInputField(input, "ReturnConsumedCapacity", _dynamodbReturnConsumedCapacity); err != nil {
			log.Errorf("invalid --return-consumed-capacity: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchGetItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The BatchWriteItem operation puts or deletes multiple items in one or more
// tables. A single call to BatchWriteItem can transmit up to 16MB of data over
// the network, consisting of up to 25 item put or delete operations. While
// individual items can be up to 400 KB once stored, it's important to note that an
// item's representation might be greater than 400KB while being sent in DynamoDB's
// JSON format for the API call. For more details on this distinction, see [Naming Rules and Data Types].
//
// BatchWriteItem cannot update items. If you perform a BatchWriteItem operation
// on an existing item, that item's values will be overwritten by the operation and
// it will appear like it was updated. To update items, we recommend you use the
// UpdateItem action.
//
// The individual PutItem and DeleteItem operations specified in BatchWriteItem
// are atomic; however BatchWriteItem as a whole is not. If any requested
// operations fail because the table's provisioned throughput is exceeded or an
// internal processing failure occurs, the failed operations are returned in the
// UnprocessedItems response parameter. You can investigate and optionally resend
// the requests. Typically, you would call BatchWriteItem in a loop. Each
// iteration would check for unprocessed items and submit a new BatchWriteItem
// request with those unprocessed items until all items have been processed.
//
// For tables and indexes with provisioned capacity, if none of the items can be
// processed due to insufficient provisioned throughput on all of the tables in the
// request, then BatchWriteItem returns a ProvisionedThroughputExceededException .
// For all tables and indexes, if none of the items can be processed due to other
// throttling scenarios (such as exceeding partition level limits), then
// BatchWriteItem returns a ThrottlingException .
//
// If DynamoDB returns any unprocessed items, you should retry the batch operation
// on those items. However, we strongly recommend that you use an exponential
// backoff algorithm. If you retry the batch operation immediately, the underlying
// read or write requests can still fail due to throttling on the individual
// tables. If you delay the batch operation using exponential backoff, the
// individual requests in the batch are much more likely to succeed.
//
// For more information, see [Batch Operations and Error Handling] in the Amazon DynamoDB Developer Guide.
//
// With BatchWriteItem , you can efficiently write or delete large amounts of data,
// such as from Amazon EMR, or copy data from another database into DynamoDB. In
// order to improve performance with these large-scale operations, BatchWriteItem
// does not behave in the same way as individual PutItem and DeleteItem calls
// would. For example, you cannot specify conditions on individual put and delete
// requests, and BatchWriteItem does not return deleted items in the response.
//
// If you use a programming language that supports concurrency, you can use
// threads to write items in parallel. Your application must include the necessary
// logic to manage the threads. With languages that don't support threading, you
// must update or delete the specified items one at a time. In both situations,
// BatchWriteItem performs the specified put and delete operations in parallel,
// giving you the power of the thread pool approach without having to introduce
// complexity into your application.
//
// Parallel processing reduces latency, but each specified put and delete request
// consumes the same number of write capacity units whether it is processed in
// parallel or not. Delete operations on nonexistent items consume one write
// capacity unit.
//
// If one or more of the following is true, DynamoDB rejects the entire batch
// write operation:
//
// - One or more tables specified in the BatchWriteItem request does not exist.
//
// - Primary key attributes specified on an item in the request do not match
// those in the corresponding table's primary key schema.
//
// - You try to perform multiple operations on the same item in the same
// BatchWriteItem request. For example, you cannot put and delete the same item
// in the same BatchWriteItem request.
//
// - Your request contains at least two items with identical hash and range keys
// (which essentially is two put operations).
//
// - There are more than 25 requests in the batch.
//
// - Any individual item in a batch exceeds 400 KB.
//
// - The total request size exceeds 16 MB.
//
// - Any individual items with keys exceeding the key length limits. For a
// partition key, the limit is 2048 bytes and for a sort key, the limit is 1024
// bytes.
//
// [Batch Operations and Error Handling]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ErrorHandling.html#Programming.Errors.BatchOperations
// [Naming Rules and Data Types]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.NamingRulesDataTypes.html
func dynamodb_BatchWriteItem(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.BatchWriteItemInput{
		// RequestItems: map[string][]types.WriteRequest, // Required
	}

	if len(_dynamodbRequestItems) > 0 {
		if err := assignInputField(input, "RequestItems", _dynamodbRequestItems); err != nil {
			log.Errorf("invalid --request-items: %s", err.Error())
			return
		}
	}
	if len(_dynamodbReturnConsumedCapacity) > 0 {
		if err := assignInputField(input, "ReturnConsumedCapacity", _dynamodbReturnConsumedCapacity); err != nil {
			log.Errorf("invalid --return-consumed-capacity: %s", err.Error())
			return
		}
	}
	if len(_dynamodbReturnItemCollectionMetrics) > 0 {
		if err := assignInputField(input, "ReturnItemCollectionMetrics", _dynamodbReturnItemCollectionMetrics); err != nil {
			log.Errorf("invalid --return-item-collection-metrics: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchWriteItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a backup for an existing table.
// Each time you create an on-demand backup, the entire table data is backed up.
// There is no limit to the number of on-demand backups that can be taken.
//
// When you create an on-demand backup, a time marker of the request is cataloged,
// and the backup is created asynchronously, by applying all changes until the time
// of the request to the last full table snapshot. Backup requests are processed
// instantaneously and become available for restore within minutes.
//
// You can call CreateBackup at a maximum rate of 50 times per second.
//
// All backups in DynamoDB work without consuming any provisioned throughput on
// the table.
//
// If you submit a backup request on 2018-12-14 at 14:25:00, the backup is
// guaranteed to contain all data committed to the table up to 14:24:00, and data
// committed after 14:26:00 will not be. The backup might contain data
// modifications made between 14:24:00 and 14:26:00. On-demand backup does not
// support causal consistency.
//
// Along with data, the following are also included on the backups:
//
// - Global secondary indexes (GSIs)
//
// - Local secondary indexes (LSIs)
//
// - Streams
//
// - Provisioned read and write capacity
func dynamodb_CreateBackup(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.CreateBackupInput{
		// BackupName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_dynamodbBackupName) > 0 {
		input.BackupName = aws.String(_dynamodbBackupName)
	}
	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}

	if resp, err := client.CreateBackup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a global table from an existing table. A global table creates a
// replication relationship between two or more DynamoDB tables with the same table
// name in the provided Regions.
//
// This documentation is for version 2017.11.29 (Legacy) of global tables, which
// should be avoided for new global tables. Customers should use [Global Tables version 2019.11.21 (Current)]when possible,
// because it provides greater flexibility, higher efficiency, and consumes less
// write capacity than 2017.11.29 (Legacy).
//
// To determine which version you're using, see [Determining the global table version you are using]. To update existing global tables
// from version 2017.11.29 (Legacy) to version 2019.11.21 (Current), see [Upgrading global tables].
//
// If you want to add a new replica table to a global table, each of the following
// conditions must be true:
//
// - The table must have the same primary key as all of the other replicas.
//
// - The table must have the same name as all of the other replicas.
//
// - The table must have DynamoDB Streams enabled, with the stream containing
// both the new and the old images of the item.
//
// - None of the replica tables in the global table can contain any data.
//
// If global secondary indexes are specified, then the following conditions must
// also be met:
//
// - The global secondary indexes must have the same name.
//
// - The global secondary indexes must have the same hash key and sort key (if
// present).
//
// If local secondary indexes are specified, then the following conditions must
// also be met:
//
// - The local secondary indexes must have the same name.
//
// - The local secondary indexes must have the same hash key and sort key (if
// present).
//
// Write capacity settings should be set consistently across your replica tables
// and secondary indexes. DynamoDB strongly recommends enabling auto scaling to
// manage the write capacity settings for all of your global tables replicas and
// indexes.
//
// If you prefer to manage write capacity settings manually, you should provision
// equal replicated write capacity units to your replica tables. You should also
// provision equal replicated write capacity units to matching secondary indexes
// across your global table.
//
// [Global Tables version 2019.11.21 (Current)]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GlobalTables.html
// [Upgrading global tables]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/V2globaltables_upgrade.html
// [Determining the global table version you are using]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.DetermineVersion.html
func dynamodb_CreateGlobalTable(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.CreateGlobalTableInput{
		// GlobalTableName: *string, // Required
		// ReplicationGroup: []types.Replica, // Required
	}

	if len(_dynamodbGlobalTableName) > 0 {
		input.GlobalTableName = aws.String(_dynamodbGlobalTableName)
	}
	if len(_dynamodbReplicationGroup) > 0 {
		if err := assignInputField(input, "ReplicationGroup", _dynamodbReplicationGroup); err != nil {
			log.Errorf("invalid --replication-group: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGlobalTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The CreateTable operation adds a new table to your account. In an Amazon Web
// Services account, table names must be unique within each Region. That is, you
// can have two tables with same name if you create the tables in different
// Regions.
//
// CreateTable is an asynchronous operation. Upon receiving a CreateTable request,
// DynamoDB immediately returns a response with a TableStatus of CREATING . After
// the table is created, DynamoDB sets the TableStatus to ACTIVE . You can perform
// read and write operations only on an ACTIVE table.
//
// You can optionally define secondary indexes on the new table, as part of the
// CreateTable operation. If you want to create multiple tables with secondary
// indexes on them, you must create the tables sequentially. Only one table with
// secondary indexes can be in the CREATING state at any given time.
//
// You can use the DescribeTable action to check the table status.
func dynamodb_CreateTable(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.CreateTableInput{
		// TableName: *string, // Required
	}

	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}
	if len(_dynamodbAttributeDefinitions) > 0 {
		if err := assignInputField(input, "AttributeDefinitions", _dynamodbAttributeDefinitions); err != nil {
			log.Errorf("invalid --attribute-definitions: %s", err.Error())
			return
		}
	}
	if len(_dynamodbBillingMode) > 0 {
		if err := assignInputField(input, "BillingMode", _dynamodbBillingMode); err != nil {
			log.Errorf("invalid --billing-mode: %s", err.Error())
			return
		}
	}
	if len(_dynamodbDeletionProtectionEnabled) > 0 {
		if err := assignInputField(input, "DeletionProtectionEnabled", _dynamodbDeletionProtectionEnabled); err != nil {
			log.Errorf("invalid --deletion-protection-enabled: %s", err.Error())
			return
		}
	}
	if len(_dynamodbGlobalSecondaryIndexes) > 0 {
		if err := assignInputField(input, "GlobalSecondaryIndexes", _dynamodbGlobalSecondaryIndexes); err != nil {
			log.Errorf("invalid --global-secondary-indexes: %s", err.Error())
			return
		}
	}
	if len(_dynamodbGlobalTableSettingsReplicationMode) > 0 {
		if err := assignInputField(input, "GlobalTableSettingsReplicationMode", _dynamodbGlobalTableSettingsReplicationMode); err != nil {
			log.Errorf("invalid --global-table-settings-replication-mode: %s", err.Error())
			return
		}
	}
	if len(_dynamodbGlobalTableSourceArn) > 0 {
		input.GlobalTableSourceArn = aws.String(_dynamodbGlobalTableSourceArn)
	}
	if len(_dynamodbKeySchema) > 0 {
		if err := assignInputField(input, "KeySchema", _dynamodbKeySchema); err != nil {
			log.Errorf("invalid --key-schema: %s", err.Error())
			return
		}
	}
	if len(_dynamodbLocalSecondaryIndexes) > 0 {
		if err := assignInputField(input, "LocalSecondaryIndexes", _dynamodbLocalSecondaryIndexes); err != nil {
			log.Errorf("invalid --local-secondary-indexes: %s", err.Error())
			return
		}
	}
	if len(_dynamodbOnDemandThroughput) > 0 {
		if err := assignInputField(input, "OnDemandThroughput", _dynamodbOnDemandThroughput); err != nil {
			log.Errorf("invalid --on-demand-throughput: %s", err.Error())
			return
		}
	}
	if len(_dynamodbProvisionedThroughput) > 0 {
		if err := assignInputField(input, "ProvisionedThroughput", _dynamodbProvisionedThroughput); err != nil {
			log.Errorf("invalid --provisioned-throughput: %s", err.Error())
			return
		}
	}
	if len(_dynamodbResourcePolicy) > 0 {
		input.ResourcePolicy = aws.String(_dynamodbResourcePolicy)
	}
	if len(_dynamodbSSESpecification) > 0 {
		if err := assignInputField(input, "SSESpecification", _dynamodbSSESpecification); err != nil {
			log.Errorf("invalid --sse-specification: %s", err.Error())
			return
		}
	}
	if len(_dynamodbStreamSpecification) > 0 {
		if err := assignInputField(input, "StreamSpecification", _dynamodbStreamSpecification); err != nil {
			log.Errorf("invalid --stream-specification: %s", err.Error())
			return
		}
	}
	if len(_dynamodbTableClass) > 0 {
		if err := assignInputField(input, "TableClass", _dynamodbTableClass); err != nil {
			log.Errorf("invalid --table-class: %s", err.Error())
			return
		}
	}
	if len(_dynamodbTags) > 0 {
		if err := assignInputField(input, "Tags", _dynamodbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_dynamodbWarmThroughput) > 0 {
		if err := assignInputField(input, "WarmThroughput", _dynamodbWarmThroughput); err != nil {
			log.Errorf("invalid --warm-throughput: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing backup of a table.
// You can call DeleteBackup at a maximum rate of 10 times per second.
func dynamodb_DeleteBackup(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.DeleteBackupInput{
		// BackupArn: *string, // Required
	}

	if len(_dynamodbBackupArn) > 0 {
		input.BackupArn = aws.String(_dynamodbBackupArn)
	}

	if resp, err := client.DeleteBackup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a single item in a table by primary key. You can perform a conditional
// delete operation that deletes the item if it exists, or if it has an expected
// attribute value.
//
// In addition to deleting an item, you can also return the item's attribute
// values in the same operation, using the ReturnValues parameter.
//
// Unless you specify conditions, the DeleteItem is an idempotent operation;
// running it multiple times on the same item or attribute does not result in an
// error response.
//
// Conditional deletes are useful for deleting items only if specific conditions
// are met. If those conditions are met, DynamoDB performs the delete. Otherwise,
// the item is not deleted.
func dynamodb_DeleteItem(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.DeleteItemInput{
		// Key: map[string]types.AttributeValue, // Required
		// TableName: *string, // Required
	}

	if len(_dynamodbKey) > 0 {
		if err := assignInputField(input, "Key", _dynamodbKey); err != nil {
			log.Errorf("invalid --key: %s", err.Error())
			return
		}
	}
	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}
	if len(_dynamodbConditionExpression) > 0 {
		input.ConditionExpression = aws.String(_dynamodbConditionExpression)
	}
	if len(_dynamodbConditionalOperator) > 0 {
		if err := assignInputField(input, "ConditionalOperator", _dynamodbConditionalOperator); err != nil {
			log.Errorf("invalid --conditional-operator: %s", err.Error())
			return
		}
	}
	if len(_dynamodbExpected) > 0 {
		if err := assignInputField(input, "Expected", _dynamodbExpected); err != nil {
			log.Errorf("invalid --expected: %s", err.Error())
			return
		}
	}
	if len(_dynamodbExpressionAttributeNames) > 0 {
		if err := assignInputField(input, "ExpressionAttributeNames", _dynamodbExpressionAttributeNames); err != nil {
			log.Errorf("invalid --expression-attribute-names: %s", err.Error())
			return
		}
	}
	if len(_dynamodbExpressionAttributeValues) > 0 {
		if err := assignInputField(input, "ExpressionAttributeValues", _dynamodbExpressionAttributeValues); err != nil {
			log.Errorf("invalid --expression-attribute-values: %s", err.Error())
			return
		}
	}
	if len(_dynamodbReturnConsumedCapacity) > 0 {
		if err := assignInputField(input, "ReturnConsumedCapacity", _dynamodbReturnConsumedCapacity); err != nil {
			log.Errorf("invalid --return-consumed-capacity: %s", err.Error())
			return
		}
	}
	if len(_dynamodbReturnItemCollectionMetrics) > 0 {
		if err := assignInputField(input, "ReturnItemCollectionMetrics", _dynamodbReturnItemCollectionMetrics); err != nil {
			log.Errorf("invalid --return-item-collection-metrics: %s", err.Error())
			return
		}
	}
	if len(_dynamodbReturnValues) > 0 {
		if err := assignInputField(input, "ReturnValues", _dynamodbReturnValues); err != nil {
			log.Errorf("invalid --return-values: %s", err.Error())
			return
		}
	}
	if len(_dynamodbReturnValuesOnConditionCheckFailure) > 0 {
		if err := assignInputField(input, "ReturnValuesOnConditionCheckFailure", _dynamodbReturnValuesOnConditionCheckFailure); err != nil {
			log.Errorf("invalid --return-values-on-condition-check-failure: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the resource-based policy attached to the resource, which can be a
// table or stream.
//
// DeleteResourcePolicy is an idempotent operation; running it multiple times on
// the same resource doesn't result in an error response, unless you specify an
// ExpectedRevisionId , which will then return a PolicyNotFoundException .
//
// To make sure that you don't inadvertently lock yourself out of your own
// resources, the root principal in your Amazon Web Services account can perform
// DeleteResourcePolicy requests, even if your resource-based policy explicitly
// denies the root principal's access.
//
// DeleteResourcePolicy is an asynchronous operation. If you issue a
// GetResourcePolicy request immediately after running the DeleteResourcePolicy
// request, DynamoDB might still return the deleted policy. This is because the
// policy for your resource might not have been deleted yet. Wait for a few
// seconds, and then try the GetResourcePolicy request again.
func dynamodb_DeleteResourcePolicy(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.DeleteResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_dynamodbResourceArn) > 0 {
		input.ResourceArn = aws.String(_dynamodbResourceArn)
	}
	if len(_dynamodbExpectedRevisionId) > 0 {
		input.ExpectedRevisionId = aws.String(_dynamodbExpectedRevisionId)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The DeleteTable operation deletes a table and all of its items. After a
// DeleteTable request, the specified table is in the DELETING state until
// DynamoDB completes the deletion. If the table is in the ACTIVE state, you can
// delete it. If a table is in CREATING or UPDATING states, then DynamoDB returns
// a ResourceInUseException . If the specified table does not exist, DynamoDB
// returns a ResourceNotFoundException . If table is already in the DELETING
// state, no error is returned.
//
// DynamoDB might continue to accept data read and write operations, such as
// GetItem and PutItem , on a table in the DELETING state until the table deletion
// is complete. For the full list of table states, see [TableStatus].
//
// When you delete a table, any indexes on that table are also deleted.
//
// If you have DynamoDB Streams enabled on the table, then the corresponding
// stream on that table goes into the DISABLED state, and the stream is
// automatically deleted after 24 hours.
//
// Use the DescribeTable action to check the status of the table.
//
// [TableStatus]: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_TableDescription.html#DDB-Type-TableDescription-TableStatus
func dynamodb_DeleteTable(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.DeleteTableInput{
		// TableName: *string, // Required
	}

	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}

	if resp, err := client.DeleteTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an existing backup of a table.
// You can call DescribeBackup at a maximum rate of 10 times per second.
func dynamodb_DescribeBackup(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.DescribeBackupInput{
		// BackupArn: *string, // Required
	}

	if len(_dynamodbBackupArn) > 0 {
		input.BackupArn = aws.String(_dynamodbBackupArn)
	}

	if resp, err := client.DescribeBackup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Checks the status of continuous backups and point in time recovery on the
// specified table. Continuous backups are ENABLED on all tables at table
// creation. If point in time recovery is enabled, PointInTimeRecoveryStatus will
// be set to ENABLED.
//
// After continuous backups and point in time recovery are enabled, you can
// restore to any point in time within EarliestRestorableDateTime and
// LatestRestorableDateTime .
//
// LatestRestorableDateTime is typically 5 minutes before the current time. You
// can restore your table to any point in time in the last 35 days. You can set the
// recovery period to any value between 1 and 35 days.
//
// You can call DescribeContinuousBackups at a maximum rate of 10 times per second.
func dynamodb_DescribeContinuousBackups(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.DescribeContinuousBackupsInput{
		// TableName: *string, // Required
	}

	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}

	if resp, err := client.DescribeContinuousBackups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about contributor insights for a given table or global
// secondary index.
func dynamodb_DescribeContributorInsights(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.DescribeContributorInsightsInput{
		// TableName: *string, // Required
	}

	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}
	if len(_dynamodbIndexName) > 0 {
		input.IndexName = aws.String(_dynamodbIndexName)
	}

	if resp, err := client.DescribeContributorInsights(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the regional endpoint information. For more information on policy
// permissions, please see [Internetwork traffic privacy].
//
// [Internetwork traffic privacy]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/inter-network-traffic-privacy.html#inter-network-traffic-DescribeEndpoints
func dynamodb_DescribeEndpoints(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.DescribeEndpointsInput{}

	if resp, err := client.DescribeEndpoints(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an existing table export.
func dynamodb_DescribeExport(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.DescribeExportInput{
		// ExportArn: *string, // Required
	}

	if len(_dynamodbExportArn) > 0 {
		input.ExportArn = aws.String(_dynamodbExportArn)
	}

	if resp, err := client.DescribeExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified global table.
// This documentation is for version 2017.11.29 (Legacy) of global tables, which
// should be avoided for new global tables. Customers should use [Global Tables version 2019.11.21 (Current)]when possible,
// because it provides greater flexibility, higher efficiency, and consumes less
// write capacity than 2017.11.29 (Legacy).
//
// To determine which version you're using, see [Determining the global table version you are using]. To update existing global tables
// from version 2017.11.29 (Legacy) to version 2019.11.21 (Current), see [Upgrading global tables].
//
// [Global Tables version 2019.11.21 (Current)]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GlobalTables.html
// [Upgrading global tables]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/V2globaltables_upgrade.html
// [Determining the global table version you are using]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.DetermineVersion.html
func dynamodb_DescribeGlobalTable(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.DescribeGlobalTableInput{
		// GlobalTableName: *string, // Required
	}

	if len(_dynamodbGlobalTableName) > 0 {
		input.GlobalTableName = aws.String(_dynamodbGlobalTableName)
	}

	if resp, err := client.DescribeGlobalTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes Region-specific settings for a global table.
// This documentation is for version 2017.11.29 (Legacy) of global tables, which
// should be avoided for new global tables. Customers should use [Global Tables version 2019.11.21 (Current)]when possible,
// because it provides greater flexibility, higher efficiency, and consumes less
// write capacity than 2017.11.29 (Legacy).
//
// To determine which version you're using, see [Determining the global table version you are using]. To update existing global tables
// from version 2017.11.29 (Legacy) to version 2019.11.21 (Current), see [Upgrading global tables].
//
// [Global Tables version 2019.11.21 (Current)]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GlobalTables.html
// [Upgrading global tables]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/V2globaltables_upgrade.html
// [Determining the global table version you are using]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.DetermineVersion.html
func dynamodb_DescribeGlobalTableSettings(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.DescribeGlobalTableSettingsInput{
		// GlobalTableName: *string, // Required
	}

	if len(_dynamodbGlobalTableName) > 0 {
		input.GlobalTableName = aws.String(_dynamodbGlobalTableName)
	}

	if resp, err := client.DescribeGlobalTableSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Represents the properties of the import.
func dynamodb_DescribeImport(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.DescribeImportInput{
		// ImportArn: *string, // Required
	}

	if len(_dynamodbImportArn) > 0 {
		input.ImportArn = aws.String(_dynamodbImportArn)
	}

	if resp, err := client.DescribeImport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the status of Kinesis streaming.
func dynamodb_DescribeKinesisStreamingDestination(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.DescribeKinesisStreamingDestinationInput{
		// TableName: *string, // Required
	}

	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}

	if resp, err := client.DescribeKinesisStreamingDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the current provisioned-capacity quotas for your Amazon Web Services
// account in a Region, both for the Region as a whole and for any one DynamoDB
// table that you create there.
//
// When you establish an Amazon Web Services account, the account has initial
// quotas on the maximum read capacity units and write capacity units that you can
// provision across all of your DynamoDB tables in a given Region. Also, there are
// per-table quotas that apply when you create a table there. For more information,
// see [Service, Account, and Table Quotas]page in the Amazon DynamoDB Developer Guide.
//
// Although you can increase these quotas by filing a case at [Amazon Web Services Support Center], obtaining the
// increase is not instantaneous. The DescribeLimits action lets you write code to
// compare the capacity you are currently using to those quotas imposed by your
// account so that you have enough time to apply for an increase before you hit a
// quota.
//
// For example, you could use one of the Amazon Web Services SDKs to do the
// following:
//
// - Call DescribeLimits for a particular Region to obtain your current account
// quotas on provisioned capacity there.
//
// - Create a variable to hold the aggregate read capacity units provisioned for
// all your tables in that Region, and one to hold the aggregate write capacity
// units. Zero them both.
//
// - Call ListTables to obtain a list of all your DynamoDB tables.
//
// - For each table name listed by ListTables , do the following:
//
// - Call DescribeTable with the table name.
//
// - Use the data returned by DescribeTable to add the read capacity units and
// write capacity units provisioned for the table itself to your variables.
//
// - If the table has one or more global secondary indexes (GSIs), loop over
// these GSIs and add their provisioned capacity values to your variables as well.
//
// - Report the account quotas for that Region returned by DescribeLimits , along
// with the total current provisioned capacity levels you have calculated.
//
// This will let you see whether you are getting close to your account-level
// quotas.
//
// The per-table quotas apply only when you are creating a new table. They
// restrict the sum of the provisioned capacity of the new table itself and all its
// global secondary indexes.
//
// For existing tables and their GSIs, DynamoDB doesn't let you increase
// provisioned capacity extremely rapidly, but the only quota that applies is that
// the aggregate provisioned capacity over all your tables and GSIs cannot exceed
// either of the per-account quotas.
//
// DescribeLimits should only be called periodically. You can expect throttling
// errors if you call it more than once in a minute.
//
// The DescribeLimits Request element has no content.
//
// [Service, Account, and Table Quotas]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html
// [Amazon Web Services Support Center]: https://console.aws.amazon.com/support/home#/
func dynamodb_DescribeLimits(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.DescribeLimitsInput{}

	if resp, err := client.DescribeLimits(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the table, including the current status of the table,
// when it was created, the primary key schema, and any indexes on the table.
//
// If you issue a DescribeTable request immediately after a CreateTable request,
// DynamoDB might return a ResourceNotFoundException . This is because
// DescribeTable uses an eventually consistent query, and the metadata for your
// table might not be available at that moment. Wait for a few seconds, and then
// try the DescribeTable request again.
func dynamodb_DescribeTable(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.DescribeTableInput{
		// TableName: *string, // Required
	}

	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}

	if resp, err := client.DescribeTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes auto scaling settings across replicas of the global table at once.
func dynamodb_DescribeTableReplicaAutoScaling(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.DescribeTableReplicaAutoScalingInput{
		// TableName: *string, // Required
	}

	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}

	if resp, err := client.DescribeTableReplicaAutoScaling(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gives a description of the Time to Live (TTL) status on the specified table.
func dynamodb_DescribeTimeToLive(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.DescribeTimeToLiveInput{
		// TableName: *string, // Required
	}

	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}

	if resp, err := client.DescribeTimeToLive(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops replication from the DynamoDB table to the Kinesis data stream. This is
// done without deleting either of the resources.
func dynamodb_DisableKinesisStreamingDestination(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.DisableKinesisStreamingDestinationInput{
		// StreamArn: *string, // Required
		// TableName: *string, // Required
	}

	if len(_dynamodbStreamArn) > 0 {
		input.StreamArn = aws.String(_dynamodbStreamArn)
	}
	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}
	if len(_dynamodbEnableKinesisStreamingConfiguration) > 0 {
		if err := assignInputField(input, "EnableKinesisStreamingConfiguration", _dynamodbEnableKinesisStreamingConfiguration); err != nil {
			log.Errorf("invalid --enable-kinesis-streaming-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.DisableKinesisStreamingDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts table data replication to the specified Kinesis data stream at a
// timestamp chosen during the enable workflow. If this operation doesn't return
// results immediately, use DescribeKinesisStreamingDestination to check if
// streaming to the Kinesis data stream is ACTIVE.
func dynamodb_EnableKinesisStreamingDestination(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.EnableKinesisStreamingDestinationInput{
		// StreamArn: *string, // Required
		// TableName: *string, // Required
	}

	if len(_dynamodbStreamArn) > 0 {
		input.StreamArn = aws.String(_dynamodbStreamArn)
	}
	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}
	if len(_dynamodbEnableKinesisStreamingConfiguration) > 0 {
		if err := assignInputField(input, "EnableKinesisStreamingConfiguration", _dynamodbEnableKinesisStreamingConfiguration); err != nil {
			log.Errorf("invalid --enable-kinesis-streaming-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.EnableKinesisStreamingDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation allows you to perform reads and singleton writes on data stored
// in DynamoDB, using PartiQL.
//
// For PartiQL reads ( SELECT statement), if the total number of processed items
// exceeds the maximum dataset size limit of 1 MB, the read stops and results are
// returned to the user as a LastEvaluatedKey value to continue the read in a
// subsequent operation. If the filter criteria in WHERE clause does not match any
// data, the read will return an empty result set.
//
// A single SELECT statement response can return up to the maximum number of items
// (if using the Limit parameter) or a maximum of 1 MB of data (and then apply any
// filtering to the results using WHERE clause). If LastEvaluatedKey is present in
// the response, you need to paginate the result set. If NextToken is present, you
// need to paginate the result set and include NextToken .
func dynamodb_ExecuteStatement(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.ExecuteStatementInput{
		// Statement: *string, // Required
	}

	if len(_dynamodbStatement) > 0 {
		input.Statement = aws.String(_dynamodbStatement)
	}
	if len(_dynamodbConsistentRead) > 0 {
		if err := assignInputField(input, "ConsistentRead", _dynamodbConsistentRead); err != nil {
			log.Errorf("invalid --consistent-read: %s", err.Error())
			return
		}
	}
	if len(_dynamodbLimit) > 0 {
		if err := assignInputField(input, "Limit", _dynamodbLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_dynamodbNextToken) > 0 {
		input.NextToken = aws.String(_dynamodbNextToken)
	}
	if len(_dynamodbParameters) > 0 {
		if err := assignInputField(input, "Parameters", _dynamodbParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_dynamodbReturnConsumedCapacity) > 0 {
		if err := assignInputField(input, "ReturnConsumedCapacity", _dynamodbReturnConsumedCapacity); err != nil {
			log.Errorf("invalid --return-consumed-capacity: %s", err.Error())
			return
		}
	}
	if len(_dynamodbReturnValuesOnConditionCheckFailure) > 0 {
		if err := assignInputField(input, "ReturnValuesOnConditionCheckFailure", _dynamodbReturnValuesOnConditionCheckFailure); err != nil {
			log.Errorf("invalid --return-values-on-condition-check-failure: %s", err.Error())
			return
		}
	}

	if resp, err := client.ExecuteStatement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation allows you to perform transactional reads or writes on data
// stored in DynamoDB, using PartiQL.
//
// The entire transaction must consist of either read statements or write
// statements, you cannot mix both in one transaction. The EXISTS function is an
// exception and can be used to check the condition of specific attributes of the
// item in a similar manner to ConditionCheck in the [TransactWriteItems] API.
//
// [TransactWriteItems]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/transaction-apis.html#transaction-apis-txwriteitems
func dynamodb_ExecuteTransaction(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.ExecuteTransactionInput{
		// TransactStatements: []types.ParameterizedStatement, // Required
	}

	if len(_dynamodbTransactStatements) > 0 {
		if err := assignInputField(input, "TransactStatements", _dynamodbTransactStatements); err != nil {
			log.Errorf("invalid --transact-statements: %s", err.Error())
			return
		}
	}
	if len(_dynamodbClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_dynamodbClientRequestToken)
	}
	if len(_dynamodbReturnConsumedCapacity) > 0 {
		if err := assignInputField(input, "ReturnConsumedCapacity", _dynamodbReturnConsumedCapacity); err != nil {
			log.Errorf("invalid --return-consumed-capacity: %s", err.Error())
			return
		}
	}

	if resp, err := client.ExecuteTransaction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Exports table data to an S3 bucket. The table must have point in time recovery
// enabled, and you can export data from any time within the point in time recovery
// window.
func dynamodb_ExportTableToPointInTime(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.ExportTableToPointInTimeInput{
		// S3Bucket: *string, // Required
		// TableArn: *string, // Required
	}

	if len(_dynamodbS3Bucket) > 0 {
		input.S3Bucket = aws.String(_dynamodbS3Bucket)
	}
	if len(_dynamodbTableArn) > 0 {
		input.TableArn = aws.String(_dynamodbTableArn)
	}
	if len(_dynamodbClientToken) > 0 {
		input.ClientToken = aws.String(_dynamodbClientToken)
	}
	if len(_dynamodbExportFormat) > 0 {
		if err := assignInputField(input, "ExportFormat", _dynamodbExportFormat); err != nil {
			log.Errorf("invalid --export-format: %s", err.Error())
			return
		}
	}
	if len(_dynamodbExportTime) > 0 {
		if err := assignInputField(input, "ExportTime", _dynamodbExportTime); err != nil {
			log.Errorf("invalid --export-time: %s", err.Error())
			return
		}
	}
	if len(_dynamodbExportType) > 0 {
		if err := assignInputField(input, "ExportType", _dynamodbExportType); err != nil {
			log.Errorf("invalid --export-type: %s", err.Error())
			return
		}
	}
	if len(_dynamodbIncrementalExportSpecification) > 0 {
		if err := assignInputField(input, "IncrementalExportSpecification", _dynamodbIncrementalExportSpecification); err != nil {
			log.Errorf("invalid --incremental-export-specification: %s", err.Error())
			return
		}
	}
	if len(_dynamodbS3BucketOwner) > 0 {
		input.S3BucketOwner = aws.String(_dynamodbS3BucketOwner)
	}
	if len(_dynamodbS3Prefix) > 0 {
		input.S3Prefix = aws.String(_dynamodbS3Prefix)
	}
	if len(_dynamodbS3SseAlgorithm) > 0 {
		if err := assignInputField(input, "S3SseAlgorithm", _dynamodbS3SseAlgorithm); err != nil {
			log.Errorf("invalid --s3-sse-algorithm: %s", err.Error())
			return
		}
	}
	if len(_dynamodbS3SseKmsKeyId) > 0 {
		input.S3SseKmsKeyId = aws.String(_dynamodbS3SseKmsKeyId)
	}

	if resp, err := client.ExportTableToPointInTime(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The GetItem operation returns a set of attributes for the item with the given
// primary key. If there is no matching item, GetItem does not return any data and
// there will be no Item element in the response.
//
// GetItem provides an eventually consistent read by default. If your application
// requires a strongly consistent read, set ConsistentRead to true . Although a
// strongly consistent read might take more time than an eventually consistent
// read, it always returns the last updated value.
func dynamodb_GetItem(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.GetItemInput{
		// Key: map[string]types.AttributeValue, // Required
		// TableName: *string, // Required
	}

	if len(_dynamodbKey) > 0 {
		if err := assignInputField(input, "Key", _dynamodbKey); err != nil {
			log.Errorf("invalid --key: %s", err.Error())
			return
		}
	}
	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}
	if len(_dynamodbAttributesToGet) > 0 {
		input.AttributesToGet = append([]string(nil), _dynamodbAttributesToGet...)
	}
	if len(_dynamodbConsistentRead) > 0 {
		if err := assignInputField(input, "ConsistentRead", _dynamodbConsistentRead); err != nil {
			log.Errorf("invalid --consistent-read: %s", err.Error())
			return
		}
	}
	if len(_dynamodbExpressionAttributeNames) > 0 {
		if err := assignInputField(input, "ExpressionAttributeNames", _dynamodbExpressionAttributeNames); err != nil {
			log.Errorf("invalid --expression-attribute-names: %s", err.Error())
			return
		}
	}
	if len(_dynamodbProjectionExpression) > 0 {
		input.ProjectionExpression = aws.String(_dynamodbProjectionExpression)
	}
	if len(_dynamodbReturnConsumedCapacity) > 0 {
		if err := assignInputField(input, "ReturnConsumedCapacity", _dynamodbReturnConsumedCapacity); err != nil {
			log.Errorf("invalid --return-consumed-capacity: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the resource-based policy document attached to the resource, which can
// be a table or stream, in JSON format.
//
// GetResourcePolicy follows an [eventually consistent] model. The following list describes the outcomes
// when you issue the GetResourcePolicy request immediately after issuing another
// request:
//
// - If you issue a GetResourcePolicy request immediately after a
// PutResourcePolicy request, DynamoDB might return a PolicyNotFoundException .
//
// - If you issue a GetResourcePolicy request immediately after a
// DeleteResourcePolicy request, DynamoDB might return the policy that was
// present before the deletion request.
//
// - If you issue a GetResourcePolicy request immediately after a CreateTable
// request, which includes a resource-based policy, DynamoDB might return a
// ResourceNotFoundException or a PolicyNotFoundException .
//
// Because GetResourcePolicy uses an eventually consistent query, the metadata for
// your policy or table might not be available at that moment. Wait for a few
// seconds, and then retry the GetResourcePolicy request.
//
// After a GetResourcePolicy request returns a policy created using the
// PutResourcePolicy request, the policy will be applied in the authorization of
// requests to the resource. Because this process is eventually consistent, it will
// take some time to apply the policy to all requests to a resource. Policies that
// you attach while creating a table using the CreateTable request will always be
// applied to all requests for that table.
//
// [eventually consistent]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadConsistency.html
func dynamodb_GetResourcePolicy(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.GetResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_dynamodbResourceArn) > 0 {
		input.ResourceArn = aws.String(_dynamodbResourceArn)
	}

	if resp, err := client.GetResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports table data from an S3 bucket.
func dynamodb_ImportTable(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.ImportTableInput{
		// InputFormat: types.InputFormat, // Required
		// S3BucketSource: *types.S3BucketSource, // Required
		// TableCreationParameters: *types.TableCreationParameters, // Required
	}

	if len(_dynamodbInputFormat) > 0 {
		if err := assignInputField(input, "InputFormat", _dynamodbInputFormat); err != nil {
			log.Errorf("invalid --input-format: %s", err.Error())
			return
		}
	}
	if len(_dynamodbS3BucketSource) > 0 {
		if err := assignInputField(input, "S3BucketSource", _dynamodbS3BucketSource); err != nil {
			log.Errorf("invalid --s3-bucket-source: %s", err.Error())
			return
		}
	}
	if len(_dynamodbTableCreationParameters) > 0 {
		if err := assignInputField(input, "TableCreationParameters", _dynamodbTableCreationParameters); err != nil {
			log.Errorf("invalid --table-creation-parameters: %s", err.Error())
			return
		}
	}
	if len(_dynamodbClientToken) > 0 {
		input.ClientToken = aws.String(_dynamodbClientToken)
	}
	if len(_dynamodbInputCompressionType) > 0 {
		if err := assignInputField(input, "InputCompressionType", _dynamodbInputCompressionType); err != nil {
			log.Errorf("invalid --input-compression-type: %s", err.Error())
			return
		}
	}
	if len(_dynamodbInputFormatOptions) > 0 {
		if err := assignInputField(input, "InputFormatOptions", _dynamodbInputFormatOptions); err != nil {
			log.Errorf("invalid --input-format-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List DynamoDB backups that are associated with an Amazon Web Services account
// and weren't made with Amazon Web Services Backup. To list these backups for a
// given table, specify TableName . ListBackups returns a paginated list of
// results with at most 1 MB worth of items in a page. You can also specify a
// maximum number of entries to be returned in a page.
//
// In the request, start time is inclusive, but end time is exclusive. Note that
// these boundaries are for the time at which the original backup was requested.
//
// You can call ListBackups a maximum of five times per second.
//
// If you want to retrieve the complete list of backups made with Amazon Web
// Services Backup, use the [Amazon Web Services Backup list API.]
//
// [Amazon Web Services Backup list API.]: https://docs.aws.amazon.com/aws-backup/latest/devguide/API_ListBackupJobs.html
func dynamodb_ListBackups(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.ListBackupsInput{}

	if len(_dynamodbBackupType) > 0 {
		if err := assignInputField(input, "BackupType", _dynamodbBackupType); err != nil {
			log.Errorf("invalid --backup-type: %s", err.Error())
			return
		}
	}
	if len(_dynamodbExclusiveStartBackupArn) > 0 {
		input.ExclusiveStartBackupArn = aws.String(_dynamodbExclusiveStartBackupArn)
	}
	if len(_dynamodbLimit) > 0 {
		if err := assignInputField(input, "Limit", _dynamodbLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}
	if len(_dynamodbTimeRangeLowerBound) > 0 {
		if err := assignInputField(input, "TimeRangeLowerBound", _dynamodbTimeRangeLowerBound); err != nil {
			log.Errorf("invalid --time-range-lower-bound: %s", err.Error())
			return
		}
	}
	if len(_dynamodbTimeRangeUpperBound) > 0 {
		if err := assignInputField(input, "TimeRangeUpperBound", _dynamodbTimeRangeUpperBound); err != nil {
			log.Errorf("invalid --time-range-upper-bound: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListBackups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of ContributorInsightsSummary for a table and all its global
// secondary indexes.
func dynamodb_ListContributorInsights(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.ListContributorInsightsInput{}

	if len(_dynamodbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _dynamodbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_dynamodbNextToken) > 0 {
		input.NextToken = aws.String(_dynamodbNextToken)
	}
	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}

	if disablePaginator() {
		if resp, err := client.ListContributorInsights(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*dynamodb.ListContributorInsightsOutput
	p := dynamodb.NewListContributorInsightsPaginator(client, input)
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

// Lists completed exports within the past 90 days, in reverse alphanumeric order
// of ExportArn .
func dynamodb_ListExports(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.ListExportsInput{}

	if len(_dynamodbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _dynamodbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_dynamodbNextToken) > 0 {
		input.NextToken = aws.String(_dynamodbNextToken)
	}
	if len(_dynamodbTableArn) > 0 {
		input.TableArn = aws.String(_dynamodbTableArn)
	}

	if disablePaginator() {
		if resp, err := client.ListExports(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*dynamodb.ListExportsOutput
	p := dynamodb.NewListExportsPaginator(client, input)
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

// Lists all global tables that have a replica in the specified Region.
// This documentation is for version 2017.11.29 (Legacy) of global tables, which
// should be avoided for new global tables. Customers should use [Global Tables version 2019.11.21 (Current)]when possible,
// because it provides greater flexibility, higher efficiency, and consumes less
// write capacity than 2017.11.29 (Legacy).
//
// To determine which version you're using, see [Determining the global table version you are using]. To update existing global tables
// from version 2017.11.29 (Legacy) to version 2019.11.21 (Current), see [Upgrading global tables].
//
// [Global Tables version 2019.11.21 (Current)]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GlobalTables.html
// [Upgrading global tables]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/V2globaltables_upgrade.html
// [Determining the global table version you are using]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.DetermineVersion.html
func dynamodb_ListGlobalTables(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.ListGlobalTablesInput{}

	if len(_dynamodbExclusiveStartGlobalTableName) > 0 {
		input.ExclusiveStartGlobalTableName = aws.String(_dynamodbExclusiveStartGlobalTableName)
	}
	if len(_dynamodbLimit) > 0 {
		if err := assignInputField(input, "Limit", _dynamodbLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_dynamodbRegionName) > 0 {
		input.RegionName = aws.String(_dynamodbRegionName)
	}

	if resp, err := client.ListGlobalTables(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists completed imports within the past 90 days.
func dynamodb_ListImports(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.ListImportsInput{}

	if len(_dynamodbNextToken) > 0 {
		input.NextToken = aws.String(_dynamodbNextToken)
	}
	if len(_dynamodbPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _dynamodbPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_dynamodbTableArn) > 0 {
		input.TableArn = aws.String(_dynamodbTableArn)
	}

	if disablePaginator() {
		if resp, err := client.ListImports(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*dynamodb.ListImportsOutput
	p := dynamodb.NewListImportsPaginator(client, input)
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

// Returns an array of table names associated with the current account and
// endpoint. The output from ListTables is paginated, with each page returning a
// maximum of 100 table names.
func dynamodb_ListTables(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.ListTablesInput{}

	if len(_dynamodbExclusiveStartTableName) > 0 {
		input.ExclusiveStartTableName = aws.String(_dynamodbExclusiveStartTableName)
	}
	if len(_dynamodbLimit) > 0 {
		if err := assignInputField(input, "Limit", _dynamodbLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTables(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*dynamodb.ListTablesOutput
	p := dynamodb.NewListTablesPaginator(client, input)
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

// List all tags on an Amazon DynamoDB resource. You can call ListTagsOfResource
// up to 10 times per second, per account.
//
// For an overview on tagging DynamoDB resources, see [Tagging for DynamoDB] in the Amazon DynamoDB
// Developer Guide.
//
// [Tagging for DynamoDB]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Tagging.html
func dynamodb_ListTagsOfResource(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.ListTagsOfResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_dynamodbResourceArn) > 0 {
		input.ResourceArn = aws.String(_dynamodbResourceArn)
	}
	if len(_dynamodbNextToken) > 0 {
		input.NextToken = aws.String(_dynamodbNextToken)
	}

	if resp, err := client.ListTagsOfResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new item, or replaces an old item with a new item. If an item that
// has the same primary key as the new item already exists in the specified table,
// the new item completely replaces the existing item. You can perform a
// conditional put operation (add a new item if one with the specified primary key
// doesn't exist), or replace an existing item if it has certain attribute values.
// You can return the item's attribute values in the same operation, using the
// ReturnValues parameter.
//
// When you add an item, the primary key attributes are the only required
// attributes.
//
// Empty String and Binary attribute values are allowed. Attribute values of type
// String and Binary must have a length greater than zero if the attribute is used
// as a key attribute for a table or index. Set type attributes cannot be empty.
//
// Invalid Requests with empty values will be rejected with a ValidationException
// exception.
//
// To prevent a new item from replacing an existing item, use a conditional
// expression that contains the attribute_not_exists function with the name of the
// attribute being used as the partition key for the table. Since every record must
// contain that attribute, the attribute_not_exists function will only succeed if
// no matching item exists.
//
// To determine whether PutItem overwrote an existing item, use ReturnValues set
// to ALL_OLD . If the response includes the Attributes element, an existing item
// was overwritten.
//
// For more information about PutItem , see [Working with Items] in the Amazon DynamoDB Developer
// Guide.
//
// [Working with Items]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithItems.html
func dynamodb_PutItem(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.PutItemInput{
		// Item: map[string]types.AttributeValue, // Required
		// TableName: *string, // Required
	}

	if len(_dynamodbItem) > 0 {
		if err := assignInputField(input, "Item", _dynamodbItem); err != nil {
			log.Errorf("invalid --item: %s", err.Error())
			return
		}
	}
	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}
	if len(_dynamodbConditionExpression) > 0 {
		input.ConditionExpression = aws.String(_dynamodbConditionExpression)
	}
	if len(_dynamodbConditionalOperator) > 0 {
		if err := assignInputField(input, "ConditionalOperator", _dynamodbConditionalOperator); err != nil {
			log.Errorf("invalid --conditional-operator: %s", err.Error())
			return
		}
	}
	if len(_dynamodbExpected) > 0 {
		if err := assignInputField(input, "Expected", _dynamodbExpected); err != nil {
			log.Errorf("invalid --expected: %s", err.Error())
			return
		}
	}
	if len(_dynamodbExpressionAttributeNames) > 0 {
		if err := assignInputField(input, "ExpressionAttributeNames", _dynamodbExpressionAttributeNames); err != nil {
			log.Errorf("invalid --expression-attribute-names: %s", err.Error())
			return
		}
	}
	if len(_dynamodbExpressionAttributeValues) > 0 {
		if err := assignInputField(input, "ExpressionAttributeValues", _dynamodbExpressionAttributeValues); err != nil {
			log.Errorf("invalid --expression-attribute-values: %s", err.Error())
			return
		}
	}
	if len(_dynamodbReturnConsumedCapacity) > 0 {
		if err := assignInputField(input, "ReturnConsumedCapacity", _dynamodbReturnConsumedCapacity); err != nil {
			log.Errorf("invalid --return-consumed-capacity: %s", err.Error())
			return
		}
	}
	if len(_dynamodbReturnItemCollectionMetrics) > 0 {
		if err := assignInputField(input, "ReturnItemCollectionMetrics", _dynamodbReturnItemCollectionMetrics); err != nil {
			log.Errorf("invalid --return-item-collection-metrics: %s", err.Error())
			return
		}
	}
	if len(_dynamodbReturnValues) > 0 {
		if err := assignInputField(input, "ReturnValues", _dynamodbReturnValues); err != nil {
			log.Errorf("invalid --return-values: %s", err.Error())
			return
		}
	}
	if len(_dynamodbReturnValuesOnConditionCheckFailure) > 0 {
		if err := assignInputField(input, "ReturnValuesOnConditionCheckFailure", _dynamodbReturnValuesOnConditionCheckFailure); err != nil {
			log.Errorf("invalid --return-values-on-condition-check-failure: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches a resource-based policy document to the resource, which can be a table
// or stream. When you attach a resource-based policy using this API, the policy
// application is [eventually consistent].
//
// PutResourcePolicy is an idempotent operation; running it multiple times on the
// same resource using the same policy document will return the same revision ID.
// If you specify an ExpectedRevisionId that doesn't match the current policy's
// RevisionId , the PolicyNotFoundException will be returned.
//
// PutResourcePolicy is an asynchronous operation. If you issue a GetResourcePolicy
// request immediately after a PutResourcePolicy request, DynamoDB might return
// your previous policy, if there was one, or return the PolicyNotFoundException .
// This is because GetResourcePolicy uses an eventually consistent query, and the
// metadata for your policy or table might not be available at that moment. Wait
// for a few seconds, and then try the GetResourcePolicy request again.
//
// [eventually consistent]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadConsistency.html
func dynamodb_PutResourcePolicy(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.PutResourcePolicyInput{
		// Policy: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_dynamodbPolicy) > 0 {
		input.Policy = aws.String(_dynamodbPolicy)
	}
	if len(_dynamodbResourceArn) > 0 {
		input.ResourceArn = aws.String(_dynamodbResourceArn)
	}
	if len(_dynamodbConfirmRemoveSelfResourceAccess) > 0 {
		if err := assignInputField(input, "ConfirmRemoveSelfResourceAccess", _dynamodbConfirmRemoveSelfResourceAccess); err != nil {
			log.Errorf("invalid --confirm-remove-self-resource-access: %s", err.Error())
			return
		}
	}
	if len(_dynamodbExpectedRevisionId) > 0 {
		input.ExpectedRevisionId = aws.String(_dynamodbExpectedRevisionId)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// You must provide the name of the partition key attribute and a single value for
// that attribute. Query returns all items with that partition key value.
// Optionally, you can provide a sort key attribute and use a comparison operator
// to refine the search results.
//
// Use the KeyConditionExpression parameter to provide a specific value for the
// partition key. The Query operation will return all of the items from the table
// or index with that partition key value. You can optionally narrow the scope of
// the Query operation by specifying a sort key value and a comparison operator in
// KeyConditionExpression . To further refine the Query results, you can
// optionally provide a FilterExpression . A FilterExpression determines which
// items within the results should be returned to you. All of the other results are
// discarded.
//
// A Query operation always returns a result set. If no matching items are found,
// the result set will be empty. Queries that do not return results consume the
// minimum number of read capacity units for that type of read operation.
//
// DynamoDB calculates the number of read capacity units consumed based on item
// size, not on the amount of data that is returned to an application. The number
// of capacity units consumed will be the same whether you request all of the
// attributes (the default behavior) or just some of them (using a projection
// expression). The number will also be the same whether or not you use a
// FilterExpression .
//
// Query results are always sorted by the sort key value. If the data type of the
// sort key is Number, the results are returned in numeric order; otherwise, the
// results are returned in order of UTF-8 bytes. By default, the sort order is
// ascending. To reverse the order, set the ScanIndexForward parameter to false.
//
// A single Query operation will read up to the maximum number of items set (if
// using the Limit parameter) or a maximum of 1 MB of data and then apply any
// filtering to the results using FilterExpression . If LastEvaluatedKey is
// present in the response, you will need to paginate the result set. For more
// information, see [Paginating the Results]in the Amazon DynamoDB Developer Guide.
//
// FilterExpression is applied after a Query finishes, but before the results are
// returned. A FilterExpression cannot contain partition key or sort key
// attributes. You need to specify those attributes in the KeyConditionExpression .
//
// A Query operation can return an empty result set and a LastEvaluatedKey if all
// the items read for the page of results are filtered out.
//
// You can query a table, a local secondary index, or a global secondary index.
// For a query on a table or on a local secondary index, you can set the
// ConsistentRead parameter to true and obtain a strongly consistent result.
// Global secondary indexes support eventually consistent reads only, so do not
// specify ConsistentRead when querying a global secondary index.
//
// [Paginating the Results]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Query.html#Query.Pagination
func dynamodb_Query(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.QueryInput{
		// TableName: *string, // Required
	}

	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}
	if len(_dynamodbAttributesToGet) > 0 {
		input.AttributesToGet = append([]string(nil), _dynamodbAttributesToGet...)
	}
	if len(_dynamodbConditionalOperator) > 0 {
		if err := assignInputField(input, "ConditionalOperator", _dynamodbConditionalOperator); err != nil {
			log.Errorf("invalid --conditional-operator: %s", err.Error())
			return
		}
	}
	if len(_dynamodbConsistentRead) > 0 {
		if err := assignInputField(input, "ConsistentRead", _dynamodbConsistentRead); err != nil {
			log.Errorf("invalid --consistent-read: %s", err.Error())
			return
		}
	}
	if len(_dynamodbExclusiveStartKey) > 0 {
		if err := assignInputField(input, "ExclusiveStartKey", _dynamodbExclusiveStartKey); err != nil {
			log.Errorf("invalid --exclusive-start-key: %s", err.Error())
			return
		}
	}
	if len(_dynamodbExpressionAttributeNames) > 0 {
		if err := assignInputField(input, "ExpressionAttributeNames", _dynamodbExpressionAttributeNames); err != nil {
			log.Errorf("invalid --expression-attribute-names: %s", err.Error())
			return
		}
	}
	if len(_dynamodbExpressionAttributeValues) > 0 {
		if err := assignInputField(input, "ExpressionAttributeValues", _dynamodbExpressionAttributeValues); err != nil {
			log.Errorf("invalid --expression-attribute-values: %s", err.Error())
			return
		}
	}
	if len(_dynamodbFilterExpression) > 0 {
		input.FilterExpression = aws.String(_dynamodbFilterExpression)
	}
	if len(_dynamodbIndexName) > 0 {
		input.IndexName = aws.String(_dynamodbIndexName)
	}
	if len(_dynamodbKeyConditionExpression) > 0 {
		input.KeyConditionExpression = aws.String(_dynamodbKeyConditionExpression)
	}
	if len(_dynamodbKeyConditions) > 0 {
		if err := assignInputField(input, "KeyConditions", _dynamodbKeyConditions); err != nil {
			log.Errorf("invalid --key-conditions: %s", err.Error())
			return
		}
	}
	if len(_dynamodbLimit) > 0 {
		if err := assignInputField(input, "Limit", _dynamodbLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_dynamodbProjectionExpression) > 0 {
		input.ProjectionExpression = aws.String(_dynamodbProjectionExpression)
	}
	if len(_dynamodbQueryFilter) > 0 {
		if err := assignInputField(input, "QueryFilter", _dynamodbQueryFilter); err != nil {
			log.Errorf("invalid --query-filter: %s", err.Error())
			return
		}
	}
	if len(_dynamodbReturnConsumedCapacity) > 0 {
		if err := assignInputField(input, "ReturnConsumedCapacity", _dynamodbReturnConsumedCapacity); err != nil {
			log.Errorf("invalid --return-consumed-capacity: %s", err.Error())
			return
		}
	}
	if len(_dynamodbScanIndexForward) > 0 {
		if err := assignInputField(input, "ScanIndexForward", _dynamodbScanIndexForward); err != nil {
			log.Errorf("invalid --scan-index-forward: %s", err.Error())
			return
		}
	}
	if len(_dynamodbSelect) > 0 {
		if err := assignInputField(input, "Select", _dynamodbSelect); err != nil {
			log.Errorf("invalid --select: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.Query(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*dynamodb.QueryOutput
	p := dynamodb.NewQueryPaginator(client, input)
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

// Creates a new table from an existing backup. Any number of users can execute up
// to 50 concurrent restores (any type of restore) in a given account.
//
// You can call RestoreTableFromBackup at a maximum rate of 10 times per second.
//
// You must manually set up the following on the restored table:
//
// - Auto scaling policies
//
// - IAM policies
//
// - Amazon CloudWatch metrics and alarms
//
// - Tags
//
// - Stream settings
//
// - Time to Live (TTL) settings
func dynamodb_RestoreTableFromBackup(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.RestoreTableFromBackupInput{
		// BackupArn: *string, // Required
		// TargetTableName: *string, // Required
	}

	if len(_dynamodbBackupArn) > 0 {
		input.BackupArn = aws.String(_dynamodbBackupArn)
	}
	if len(_dynamodbTargetTableName) > 0 {
		input.TargetTableName = aws.String(_dynamodbTargetTableName)
	}
	if len(_dynamodbBillingModeOverride) > 0 {
		if err := assignInputField(input, "BillingModeOverride", _dynamodbBillingModeOverride); err != nil {
			log.Errorf("invalid --billing-mode-override: %s", err.Error())
			return
		}
	}
	if len(_dynamodbGlobalSecondaryIndexOverride) > 0 {
		if err := assignInputField(input, "GlobalSecondaryIndexOverride", _dynamodbGlobalSecondaryIndexOverride); err != nil {
			log.Errorf("invalid --global-secondary-index-override: %s", err.Error())
			return
		}
	}
	if len(_dynamodbLocalSecondaryIndexOverride) > 0 {
		if err := assignInputField(input, "LocalSecondaryIndexOverride", _dynamodbLocalSecondaryIndexOverride); err != nil {
			log.Errorf("invalid --local-secondary-index-override: %s", err.Error())
			return
		}
	}
	if len(_dynamodbOnDemandThroughputOverride) > 0 {
		if err := assignInputField(input, "OnDemandThroughputOverride", _dynamodbOnDemandThroughputOverride); err != nil {
			log.Errorf("invalid --on-demand-throughput-override: %s", err.Error())
			return
		}
	}
	if len(_dynamodbProvisionedThroughputOverride) > 0 {
		if err := assignInputField(input, "ProvisionedThroughputOverride", _dynamodbProvisionedThroughputOverride); err != nil {
			log.Errorf("invalid --provisioned-throughput-override: %s", err.Error())
			return
		}
	}
	if len(_dynamodbSSESpecificationOverride) > 0 {
		if err := assignInputField(input, "SSESpecificationOverride", _dynamodbSSESpecificationOverride); err != nil {
			log.Errorf("invalid --sse-specification-override: %s", err.Error())
			return
		}
	}

	if resp, err := client.RestoreTableFromBackup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restores the specified table to the specified point in time within
// EarliestRestorableDateTime and LatestRestorableDateTime . You can restore your
// table to any point in time in the last 35 days. You can set the recovery period
// to any value between 1 and 35 days. Any number of users can execute up to 50
// concurrent restores (any type of restore) in a given account.
//
// When you restore using point in time recovery, DynamoDB restores your table
// data to the state based on the selected date and time (day:hour:minute:second)
// to a new table.
//
// Along with data, the following are also included on the new restored table
// using point in time recovery:
//
// - Global secondary indexes (GSIs)
//
// - Local secondary indexes (LSIs)
//
// - Provisioned read and write capacity
//
// - Encryption settings
//
// # All these settings come from the current settings of the source table at the
//
// time of restore.
//
// You must manually set up the following on the restored table:
//
// - Auto scaling policies
//
// - IAM policies
//
// - Amazon CloudWatch metrics and alarms
//
// - Tags
//
// - Stream settings
//
// - Time to Live (TTL) settings
//
// - Point in time recovery settings
func dynamodb_RestoreTableToPointInTime(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.RestoreTableToPointInTimeInput{
		// TargetTableName: *string, // Required
	}

	if len(_dynamodbTargetTableName) > 0 {
		input.TargetTableName = aws.String(_dynamodbTargetTableName)
	}
	if len(_dynamodbBillingModeOverride) > 0 {
		if err := assignInputField(input, "BillingModeOverride", _dynamodbBillingModeOverride); err != nil {
			log.Errorf("invalid --billing-mode-override: %s", err.Error())
			return
		}
	}
	if len(_dynamodbGlobalSecondaryIndexOverride) > 0 {
		if err := assignInputField(input, "GlobalSecondaryIndexOverride", _dynamodbGlobalSecondaryIndexOverride); err != nil {
			log.Errorf("invalid --global-secondary-index-override: %s", err.Error())
			return
		}
	}
	if len(_dynamodbLocalSecondaryIndexOverride) > 0 {
		if err := assignInputField(input, "LocalSecondaryIndexOverride", _dynamodbLocalSecondaryIndexOverride); err != nil {
			log.Errorf("invalid --local-secondary-index-override: %s", err.Error())
			return
		}
	}
	if len(_dynamodbOnDemandThroughputOverride) > 0 {
		if err := assignInputField(input, "OnDemandThroughputOverride", _dynamodbOnDemandThroughputOverride); err != nil {
			log.Errorf("invalid --on-demand-throughput-override: %s", err.Error())
			return
		}
	}
	if len(_dynamodbProvisionedThroughputOverride) > 0 {
		if err := assignInputField(input, "ProvisionedThroughputOverride", _dynamodbProvisionedThroughputOverride); err != nil {
			log.Errorf("invalid --provisioned-throughput-override: %s", err.Error())
			return
		}
	}
	if len(_dynamodbRestoreDateTime) > 0 {
		if err := assignInputField(input, "RestoreDateTime", _dynamodbRestoreDateTime); err != nil {
			log.Errorf("invalid --restore-date-time: %s", err.Error())
			return
		}
	}
	if len(_dynamodbSSESpecificationOverride) > 0 {
		if err := assignInputField(input, "SSESpecificationOverride", _dynamodbSSESpecificationOverride); err != nil {
			log.Errorf("invalid --sse-specification-override: %s", err.Error())
			return
		}
	}
	if len(_dynamodbSourceTableArn) > 0 {
		input.SourceTableArn = aws.String(_dynamodbSourceTableArn)
	}
	if len(_dynamodbSourceTableName) > 0 {
		input.SourceTableName = aws.String(_dynamodbSourceTableName)
	}
	if len(_dynamodbUseLatestRestorableTime) > 0 {
		if err := assignInputField(input, "UseLatestRestorableTime", _dynamodbUseLatestRestorableTime); err != nil {
			log.Errorf("invalid --use-latest-restorable-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.RestoreTableToPointInTime(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The Scan operation returns one or more items and item attributes by accessing
// every item in a table or a secondary index. To have DynamoDB return fewer items,
// you can provide a FilterExpression operation.
//
// If the total size of scanned items exceeds the maximum dataset size limit of 1
// MB, the scan completes and results are returned to the user. The
// LastEvaluatedKey value is also returned and the requestor can use the
// LastEvaluatedKey to continue the scan in a subsequent operation. Each scan
// response also includes number of items that were scanned (ScannedCount) as part
// of the request. If using a FilterExpression , a scan result can result in no
// items meeting the criteria and the Count will result in zero. If you did not
// use a FilterExpression in the scan request, then Count is the same as
// ScannedCount .
//
// Count and ScannedCount only return the count of items specific to a single scan
// request and, unless the table is less than 1MB, do not represent the total
// number of items in the table.
//
// A single Scan operation first reads up to the maximum number of items set (if
// using the Limit parameter) or a maximum of 1 MB of data and then applies any
// filtering to the results if a FilterExpression is provided. If LastEvaluatedKey
// is present in the response, pagination is required to complete the full table
// scan. For more information, see [Paginating the Results]in the Amazon DynamoDB Developer Guide.
//
// Scan operations proceed sequentially; however, for faster performance on a
// large table or secondary index, applications can request a parallel Scan
// operation by providing the Segment and TotalSegments parameters. For more
// information, see [Parallel Scan]in the Amazon DynamoDB Developer Guide.
//
// By default, a Scan uses eventually consistent reads when accessing the items in
// a table. Therefore, the results from an eventually consistent Scan may not
// include the latest item changes at the time the scan iterates through each item
// in the table. If you require a strongly consistent read of each item as the scan
// iterates through the items in the table, you can set the ConsistentRead
// parameter to true. Strong consistency only relates to the consistency of the
// read at the item level.
//
// DynamoDB does not provide snapshot isolation for a scan operation when the
// ConsistentRead parameter is set to true. Thus, a DynamoDB scan operation does
// not guarantee that all reads in a scan see a consistent snapshot of the table
// when the scan operation was requested.
//
// [Paginating the Results]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Scan.html#Scan.Pagination
// [Parallel Scan]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Scan.html#Scan.ParallelScan
func dynamodb_Scan(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.ScanInput{
		// TableName: *string, // Required
	}

	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}
	if len(_dynamodbAttributesToGet) > 0 {
		input.AttributesToGet = append([]string(nil), _dynamodbAttributesToGet...)
	}
	if len(_dynamodbConditionalOperator) > 0 {
		if err := assignInputField(input, "ConditionalOperator", _dynamodbConditionalOperator); err != nil {
			log.Errorf("invalid --conditional-operator: %s", err.Error())
			return
		}
	}
	if len(_dynamodbConsistentRead) > 0 {
		if err := assignInputField(input, "ConsistentRead", _dynamodbConsistentRead); err != nil {
			log.Errorf("invalid --consistent-read: %s", err.Error())
			return
		}
	}
	if len(_dynamodbExclusiveStartKey) > 0 {
		if err := assignInputField(input, "ExclusiveStartKey", _dynamodbExclusiveStartKey); err != nil {
			log.Errorf("invalid --exclusive-start-key: %s", err.Error())
			return
		}
	}
	if len(_dynamodbExpressionAttributeNames) > 0 {
		if err := assignInputField(input, "ExpressionAttributeNames", _dynamodbExpressionAttributeNames); err != nil {
			log.Errorf("invalid --expression-attribute-names: %s", err.Error())
			return
		}
	}
	if len(_dynamodbExpressionAttributeValues) > 0 {
		if err := assignInputField(input, "ExpressionAttributeValues", _dynamodbExpressionAttributeValues); err != nil {
			log.Errorf("invalid --expression-attribute-values: %s", err.Error())
			return
		}
	}
	if len(_dynamodbFilterExpression) > 0 {
		input.FilterExpression = aws.String(_dynamodbFilterExpression)
	}
	if len(_dynamodbIndexName) > 0 {
		input.IndexName = aws.String(_dynamodbIndexName)
	}
	if len(_dynamodbLimit) > 0 {
		if err := assignInputField(input, "Limit", _dynamodbLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_dynamodbProjectionExpression) > 0 {
		input.ProjectionExpression = aws.String(_dynamodbProjectionExpression)
	}
	if len(_dynamodbReturnConsumedCapacity) > 0 {
		if err := assignInputField(input, "ReturnConsumedCapacity", _dynamodbReturnConsumedCapacity); err != nil {
			log.Errorf("invalid --return-consumed-capacity: %s", err.Error())
			return
		}
	}
	if len(_dynamodbScanFilter) > 0 {
		if err := assignInputField(input, "ScanFilter", _dynamodbScanFilter); err != nil {
			log.Errorf("invalid --scan-filter: %s", err.Error())
			return
		}
	}
	if len(_dynamodbSegment) > 0 {
		if err := assignInputField(input, "Segment", _dynamodbSegment); err != nil {
			log.Errorf("invalid --segment: %s", err.Error())
			return
		}
	}
	if len(_dynamodbSelect) > 0 {
		if err := assignInputField(input, "Select", _dynamodbSelect); err != nil {
			log.Errorf("invalid --select: %s", err.Error())
			return
		}
	}
	if len(_dynamodbTotalSegments) > 0 {
		if err := assignInputField(input, "TotalSegments", _dynamodbTotalSegments); err != nil {
			log.Errorf("invalid --total-segments: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.Scan(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*dynamodb.ScanOutput
	p := dynamodb.NewScanPaginator(client, input)
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

// Associate a set of tags with an Amazon DynamoDB resource. You can then activate
// these user-defined tags so that they appear on the Billing and Cost Management
// console for cost allocation tracking. You can call TagResource up to five times
// per second, per account.
//
// - TagResource is an asynchronous operation. If you issue a ListTagsOfResourcerequest
// immediately after a TagResource request, DynamoDB might return your previous
// tag set, if there was one, or an empty tag set. This is because
// ListTagsOfResource uses an eventually consistent query, and the metadata for
// your tags or table might not be available at that moment. Wait for a few
// seconds, and then try the ListTagsOfResource request again.
//
// - The application or removal of tags using TagResource and UntagResource APIs
// is eventually consistent. ListTagsOfResource API will only reflect the changes
// after a few seconds.
//
// For an overview on tagging DynamoDB resources, see [Tagging for DynamoDB] in the Amazon DynamoDB
// Developer Guide.
//
// [Tagging for DynamoDB]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Tagging.html
func dynamodb_TagResource(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_dynamodbResourceArn) > 0 {
		input.ResourceArn = aws.String(_dynamodbResourceArn)
	}
	if len(_dynamodbTags) > 0 {
		if err := assignInputField(input, "Tags", _dynamodbTags); err != nil {
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

// TransactGetItems is a synchronous operation that atomically retrieves multiple
// items from one or more tables (but not from indexes) in a single account and
// Region. A TransactGetItems call can contain up to 100 TransactGetItem objects,
// each of which contains a Get structure that specifies an item to retrieve from
// a table in the account and Region. A call to TransactGetItems cannot retrieve
// items from tables in more than one Amazon Web Services account or Region. The
// aggregate size of the items in the transaction cannot exceed 4 MB.
//
// DynamoDB rejects the entire TransactGetItems request if any of the following is
// true:
//
// - A conflicting operation is in the process of updating an item to be read.
//
// - There is insufficient provisioned capacity for the transaction to be
// completed.
//
// - There is a user error, such as an invalid data format.
//
// - The aggregate size of the items in the transaction exceeded 4 MB.
func dynamodb_TransactGetItems(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.TransactGetItemsInput{
		// TransactItems: []types.TransactGetItem, // Required
	}

	if len(_dynamodbTransactItems) > 0 {
		if err := assignInputField(input, "TransactItems", _dynamodbTransactItems); err != nil {
			log.Errorf("invalid --transact-items: %s", err.Error())
			return
		}
	}
	if len(_dynamodbReturnConsumedCapacity) > 0 {
		if err := assignInputField(input, "ReturnConsumedCapacity", _dynamodbReturnConsumedCapacity); err != nil {
			log.Errorf("invalid --return-consumed-capacity: %s", err.Error())
			return
		}
	}

	if resp, err := client.TransactGetItems(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// TransactWriteItems is a synchronous write operation that groups up to 100
// action requests. These actions can target items in different tables, but not in
// different Amazon Web Services accounts or Regions, and no two actions can target
// the same item. For example, you cannot both ConditionCheck and Update the same
// item. The aggregate size of the items in the transaction cannot exceed 4 MB.
//
// The actions are completed atomically so that either all of them succeed, or all
// of them fail. They are defined by the following objects:
//
// - Put — Initiates a PutItem operation to write a new item. This structure
// specifies the primary key of the item to be written, the name of the table to
// write it in, an optional condition expression that must be satisfied for the
// write to succeed, a list of the item's attributes, and a field indicating
// whether to retrieve the item's attributes if the condition is not met.
//
// - Update — Initiates an UpdateItem operation to update an existing item. This
// structure specifies the primary key of the item to be updated, the name of the
// table where it resides, an optional condition expression that must be satisfied
// for the update to succeed, an expression that defines one or more attributes to
// be updated, and a field indicating whether to retrieve the item's attributes if
// the condition is not met.
//
// - Delete — Initiates a DeleteItem operation to delete an existing item. This
// structure specifies the primary key of the item to be deleted, the name of the
// table where it resides, an optional condition expression that must be satisfied
// for the deletion to succeed, and a field indicating whether to retrieve the
// item's attributes if the condition is not met.
//
// - ConditionCheck — Applies a condition to an item that is not being modified
// by the transaction. This structure specifies the primary key of the item to be
// checked, the name of the table where it resides, a condition expression that
// must be satisfied for the transaction to succeed, and a field indicating whether
// to retrieve the item's attributes if the condition is not met.
//
// DynamoDB rejects the entire TransactWriteItems request if any of the following
// is true:
//
// - A condition in one of the condition expressions is not met.
//
// - An ongoing operation is in the process of updating the same item.
//
// - There is insufficient provisioned capacity for the transaction to be
// completed.
//
// - An item size becomes too large (bigger than 400 KB), a local secondary
// index (LSI) becomes too large, or a similar validation error occurs because of
// changes made by the transaction.
//
// - The aggregate size of the items in the transaction exceeds 4 MB.
//
// - There is a user error, such as an invalid data format.
func dynamodb_TransactWriteItems(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.TransactWriteItemsInput{
		// TransactItems: []types.TransactWriteItem, // Required
	}

	if len(_dynamodbTransactItems) > 0 {
		if err := assignInputField(input, "TransactItems", _dynamodbTransactItems); err != nil {
			log.Errorf("invalid --transact-items: %s", err.Error())
			return
		}
	}
	if len(_dynamodbClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_dynamodbClientRequestToken)
	}
	if len(_dynamodbReturnConsumedCapacity) > 0 {
		if err := assignInputField(input, "ReturnConsumedCapacity", _dynamodbReturnConsumedCapacity); err != nil {
			log.Errorf("invalid --return-consumed-capacity: %s", err.Error())
			return
		}
	}
	if len(_dynamodbReturnItemCollectionMetrics) > 0 {
		if err := assignInputField(input, "ReturnItemCollectionMetrics", _dynamodbReturnItemCollectionMetrics); err != nil {
			log.Errorf("invalid --return-item-collection-metrics: %s", err.Error())
			return
		}
	}

	if resp, err := client.TransactWriteItems(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the association of tags from an Amazon DynamoDB resource. You can call
// UntagResource up to five times per second, per account.
//
// - UntagResource is an asynchronous operation. If you issue a ListTagsOfResourcerequest
// immediately after an UntagResource request, DynamoDB might return your
// previous tag set, if there was one, or an empty tag set. This is because
// ListTagsOfResource uses an eventually consistent query, and the metadata for
// your tags or table might not be available at that moment. Wait for a few
// seconds, and then try the ListTagsOfResource request again.
//
// - The application or removal of tags using TagResource and UntagResource APIs
// is eventually consistent. ListTagsOfResource API will only reflect the changes
// after a few seconds.
//
// For an overview on tagging DynamoDB resources, see [Tagging for DynamoDB] in the Amazon DynamoDB
// Developer Guide.
//
// [Tagging for DynamoDB]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Tagging.html
func dynamodb_UntagResource(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_dynamodbResourceArn) > 0 {
		input.ResourceArn = aws.String(_dynamodbResourceArn)
	}
	if len(_dynamodbTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _dynamodbTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// UpdateContinuousBackups enables or disables point in time recovery for the
// specified table. A successful UpdateContinuousBackups call returns the current
// ContinuousBackupsDescription . Continuous backups are ENABLED on all tables at
// table creation. If point in time recovery is enabled, PointInTimeRecoveryStatus
// will be set to ENABLED.
//
// Once continuous backups and point in time recovery are enabled, you can restore
// to any point in time within EarliestRestorableDateTime and
// LatestRestorableDateTime .
//
// LatestRestorableDateTime is typically 5 minutes before the current time. You
// can restore your table to any point in time in the last 35 days. You can set the
// RecoveryPeriodInDays to any value between 1 and 35 days.
func dynamodb_UpdateContinuousBackups(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.UpdateContinuousBackupsInput{
		// PointInTimeRecoverySpecification: *types.PointInTimeRecoverySpecification, // Required
		// TableName: *string, // Required
	}

	if len(_dynamodbPointInTimeRecoverySpecification) > 0 {
		if err := assignInputField(input, "PointInTimeRecoverySpecification", _dynamodbPointInTimeRecoverySpecification); err != nil {
			log.Errorf("invalid --point-in-time-recovery-specification: %s", err.Error())
			return
		}
	}
	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}

	if resp, err := client.UpdateContinuousBackups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the status for contributor insights for a specific table or index.
// CloudWatch Contributor Insights for DynamoDB graphs display the partition key
// and (if applicable) sort key of frequently accessed items and frequently
// throttled items in plaintext. If you require the use of Amazon Web Services Key
// Management Service (KMS) to encrypt this table’s partition key and sort key data
// with an Amazon Web Services managed key or customer managed key, you should not
// enable CloudWatch Contributor Insights for DynamoDB for this table.
func dynamodb_UpdateContributorInsights(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.UpdateContributorInsightsInput{
		// ContributorInsightsAction: types.ContributorInsightsAction, // Required
		// TableName: *string, // Required
	}

	if len(_dynamodbContributorInsightsAction) > 0 {
		if err := assignInputField(input, "ContributorInsightsAction", _dynamodbContributorInsightsAction); err != nil {
			log.Errorf("invalid --contributor-insights-action: %s", err.Error())
			return
		}
	}
	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}
	if len(_dynamodbContributorInsightsMode) > 0 {
		if err := assignInputField(input, "ContributorInsightsMode", _dynamodbContributorInsightsMode); err != nil {
			log.Errorf("invalid --contributor-insights-mode: %s", err.Error())
			return
		}
	}
	if len(_dynamodbIndexName) > 0 {
		input.IndexName = aws.String(_dynamodbIndexName)
	}

	if resp, err := client.UpdateContributorInsights(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or removes replicas in the specified global table. The global table must
// already exist to be able to use this operation. Any replica to be added must be
// empty, have the same name as the global table, have the same key schema, have
// DynamoDB Streams enabled, and have the same provisioned and maximum write
// capacity units.
//
// This documentation is for version 2017.11.29 (Legacy) of global tables, which
// should be avoided for new global tables. Customers should use [Global Tables version 2019.11.21 (Current)]when possible,
// because it provides greater flexibility, higher efficiency, and consumes less
// write capacity than 2017.11.29 (Legacy).
//
// To determine which version you're using, see [Determining the global table version you are using]. To update existing global tables
// from version 2017.11.29 (Legacy) to version 2019.11.21 (Current), see [Upgrading global tables].
//
// If you are using global tables [Version 2019.11.21] (Current) you can use [UpdateTable] instead.
//
// Although you can use UpdateGlobalTable to add replicas and remove replicas in a
// single request, for simplicity we recommend that you issue separate requests for
// adding or removing replicas.
//
// If global secondary indexes are specified, then the following conditions must
// also be met:
//
// - The global secondary indexes must have the same name.
//
// - The global secondary indexes must have the same hash key and sort key (if
// present).
//
// - The global secondary indexes must have the same provisioned and maximum
// write capacity units.
//
// [UpdateTable]: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_UpdateTable.html
// [Global Tables version 2019.11.21 (Current)]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GlobalTables.html
// [Upgrading global tables]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/V2globaltables_upgrade.html
// [Version 2019.11.21]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GlobalTables.html
// [Determining the global table version you are using]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.DetermineVersion.html
func dynamodb_UpdateGlobalTable(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.UpdateGlobalTableInput{
		// GlobalTableName: *string, // Required
		// ReplicaUpdates: []types.ReplicaUpdate, // Required
	}

	if len(_dynamodbGlobalTableName) > 0 {
		input.GlobalTableName = aws.String(_dynamodbGlobalTableName)
	}
	if len(_dynamodbReplicaUpdates) > 0 {
		if err := assignInputField(input, "ReplicaUpdates", _dynamodbReplicaUpdates); err != nil {
			log.Errorf("invalid --replica-updates: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateGlobalTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates settings for a global table.
// This documentation is for version 2017.11.29 (Legacy) of global tables, which
// should be avoided for new global tables. Customers should use [Global Tables version 2019.11.21 (Current)]when possible,
// because it provides greater flexibility, higher efficiency, and consumes less
// write capacity than 2017.11.29 (Legacy).
//
// To determine which version you're using, see [Determining the global table version you are using]. To update existing global tables
// from version 2017.11.29 (Legacy) to version 2019.11.21 (Current), see [Upgrading global tables].
//
// [Global Tables version 2019.11.21 (Current)]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GlobalTables.html
// [Upgrading global tables]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/V2globaltables_upgrade.html
// [Determining the global table version you are using]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.DetermineVersion.html
func dynamodb_UpdateGlobalTableSettings(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.UpdateGlobalTableSettingsInput{
		// GlobalTableName: *string, // Required
	}

	if len(_dynamodbGlobalTableName) > 0 {
		input.GlobalTableName = aws.String(_dynamodbGlobalTableName)
	}
	if len(_dynamodbGlobalTableBillingMode) > 0 {
		if err := assignInputField(input, "GlobalTableBillingMode", _dynamodbGlobalTableBillingMode); err != nil {
			log.Errorf("invalid --global-table-billing-mode: %s", err.Error())
			return
		}
	}
	if len(_dynamodbGlobalTableGlobalSecondaryIndexSettingsUpdate) > 0 {
		if err := assignInputField(input, "GlobalTableGlobalSecondaryIndexSettingsUpdate", _dynamodbGlobalTableGlobalSecondaryIndexSettingsUpdate); err != nil {
			log.Errorf("invalid --global-table-global-secondary-index-settings-update: %s", err.Error())
			return
		}
	}
	if len(_dynamodbGlobalTableProvisionedWriteCapacityAutoScalingSettingsUpdate) > 0 {
		if err := assignInputField(input, "GlobalTableProvisionedWriteCapacityAutoScalingSettingsUpdate", _dynamodbGlobalTableProvisionedWriteCapacityAutoScalingSettingsUpdate); err != nil {
			log.Errorf("invalid --global-table-provisioned-write-capacity-auto-scaling-settings-update: %s", err.Error())
			return
		}
	}
	if len(_dynamodbGlobalTableProvisionedWriteCapacityUnits) > 0 {
		if err := assignInputField(input, "GlobalTableProvisionedWriteCapacityUnits", _dynamodbGlobalTableProvisionedWriteCapacityUnits); err != nil {
			log.Errorf("invalid --global-table-provisioned-write-capacity-units: %s", err.Error())
			return
		}
	}
	if len(_dynamodbReplicaSettingsUpdate) > 0 {
		if err := assignInputField(input, "ReplicaSettingsUpdate", _dynamodbReplicaSettingsUpdate); err != nil {
			log.Errorf("invalid --replica-settings-update: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateGlobalTableSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Edits an existing item's attributes, or adds a new item to the table if it does
// not already exist. You can put, delete, or add attribute values. You can also
// perform a conditional update on an existing item (insert a new attribute
// name-value pair if it doesn't exist, or replace an existing name-value pair if
// it has certain expected attribute values).
//
// You can also return the item's attribute values in the same UpdateItem
// operation using the ReturnValues parameter.
func dynamodb_UpdateItem(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.UpdateItemInput{
		// Key: map[string]types.AttributeValue, // Required
		// TableName: *string, // Required
	}

	if len(_dynamodbKey) > 0 {
		if err := assignInputField(input, "Key", _dynamodbKey); err != nil {
			log.Errorf("invalid --key: %s", err.Error())
			return
		}
	}
	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}
	if len(_dynamodbAttributeUpdates) > 0 {
		if err := assignInputField(input, "AttributeUpdates", _dynamodbAttributeUpdates); err != nil {
			log.Errorf("invalid --attribute-updates: %s", err.Error())
			return
		}
	}
	if len(_dynamodbConditionExpression) > 0 {
		input.ConditionExpression = aws.String(_dynamodbConditionExpression)
	}
	if len(_dynamodbConditionalOperator) > 0 {
		if err := assignInputField(input, "ConditionalOperator", _dynamodbConditionalOperator); err != nil {
			log.Errorf("invalid --conditional-operator: %s", err.Error())
			return
		}
	}
	if len(_dynamodbExpected) > 0 {
		if err := assignInputField(input, "Expected", _dynamodbExpected); err != nil {
			log.Errorf("invalid --expected: %s", err.Error())
			return
		}
	}
	if len(_dynamodbExpressionAttributeNames) > 0 {
		if err := assignInputField(input, "ExpressionAttributeNames", _dynamodbExpressionAttributeNames); err != nil {
			log.Errorf("invalid --expression-attribute-names: %s", err.Error())
			return
		}
	}
	if len(_dynamodbExpressionAttributeValues) > 0 {
		if err := assignInputField(input, "ExpressionAttributeValues", _dynamodbExpressionAttributeValues); err != nil {
			log.Errorf("invalid --expression-attribute-values: %s", err.Error())
			return
		}
	}
	if len(_dynamodbReturnConsumedCapacity) > 0 {
		if err := assignInputField(input, "ReturnConsumedCapacity", _dynamodbReturnConsumedCapacity); err != nil {
			log.Errorf("invalid --return-consumed-capacity: %s", err.Error())
			return
		}
	}
	if len(_dynamodbReturnItemCollectionMetrics) > 0 {
		if err := assignInputField(input, "ReturnItemCollectionMetrics", _dynamodbReturnItemCollectionMetrics); err != nil {
			log.Errorf("invalid --return-item-collection-metrics: %s", err.Error())
			return
		}
	}
	if len(_dynamodbReturnValues) > 0 {
		if err := assignInputField(input, "ReturnValues", _dynamodbReturnValues); err != nil {
			log.Errorf("invalid --return-values: %s", err.Error())
			return
		}
	}
	if len(_dynamodbReturnValuesOnConditionCheckFailure) > 0 {
		if err := assignInputField(input, "ReturnValuesOnConditionCheckFailure", _dynamodbReturnValuesOnConditionCheckFailure); err != nil {
			log.Errorf("invalid --return-values-on-condition-check-failure: %s", err.Error())
			return
		}
	}
	if len(_dynamodbUpdateExpression) > 0 {
		input.UpdateExpression = aws.String(_dynamodbUpdateExpression)
	}

	if resp, err := client.UpdateItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The command to update the Kinesis stream destination.
func dynamodb_UpdateKinesisStreamingDestination(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.UpdateKinesisStreamingDestinationInput{
		// StreamArn: *string, // Required
		// TableName: *string, // Required
	}

	if len(_dynamodbStreamArn) > 0 {
		input.StreamArn = aws.String(_dynamodbStreamArn)
	}
	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}
	if len(_dynamodbUpdateKinesisStreamingConfiguration) > 0 {
		if err := assignInputField(input, "UpdateKinesisStreamingConfiguration", _dynamodbUpdateKinesisStreamingConfiguration); err != nil {
			log.Errorf("invalid --update-kinesis-streaming-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateKinesisStreamingDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the provisioned throughput settings, global secondary indexes, or
// DynamoDB Streams settings for a given table.
//
// You can only perform one of the following operations at once:
//
// - Modify the provisioned throughput settings of the table.
//
// - Remove a global secondary index from the table.
//
// - Create a new global secondary index on the table. After the index begins
// backfilling, you can use UpdateTable to perform other operations.
//
// UpdateTable is an asynchronous operation; while it's executing, the table
// status changes from ACTIVE to UPDATING . While it's UPDATING , you can't issue
// another UpdateTable request. When the table returns to the ACTIVE state, the
// UpdateTable operation is complete.
func dynamodb_UpdateTable(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.UpdateTableInput{
		// TableName: *string, // Required
	}

	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}
	if len(_dynamodbAttributeDefinitions) > 0 {
		if err := assignInputField(input, "AttributeDefinitions", _dynamodbAttributeDefinitions); err != nil {
			log.Errorf("invalid --attribute-definitions: %s", err.Error())
			return
		}
	}
	if len(_dynamodbBillingMode) > 0 {
		if err := assignInputField(input, "BillingMode", _dynamodbBillingMode); err != nil {
			log.Errorf("invalid --billing-mode: %s", err.Error())
			return
		}
	}
	if len(_dynamodbDeletionProtectionEnabled) > 0 {
		if err := assignInputField(input, "DeletionProtectionEnabled", _dynamodbDeletionProtectionEnabled); err != nil {
			log.Errorf("invalid --deletion-protection-enabled: %s", err.Error())
			return
		}
	}
	if len(_dynamodbGlobalSecondaryIndexUpdates) > 0 {
		if err := assignInputField(input, "GlobalSecondaryIndexUpdates", _dynamodbGlobalSecondaryIndexUpdates); err != nil {
			log.Errorf("invalid --global-secondary-index-updates: %s", err.Error())
			return
		}
	}
	if len(_dynamodbGlobalTableSettingsReplicationMode) > 0 {
		if err := assignInputField(input, "GlobalTableSettingsReplicationMode", _dynamodbGlobalTableSettingsReplicationMode); err != nil {
			log.Errorf("invalid --global-table-settings-replication-mode: %s", err.Error())
			return
		}
	}
	if len(_dynamodbGlobalTableWitnessUpdates) > 0 {
		if err := assignInputField(input, "GlobalTableWitnessUpdates", _dynamodbGlobalTableWitnessUpdates); err != nil {
			log.Errorf("invalid --global-table-witness-updates: %s", err.Error())
			return
		}
	}
	if len(_dynamodbMultiRegionConsistency) > 0 {
		if err := assignInputField(input, "MultiRegionConsistency", _dynamodbMultiRegionConsistency); err != nil {
			log.Errorf("invalid --multi-region-consistency: %s", err.Error())
			return
		}
	}
	if len(_dynamodbOnDemandThroughput) > 0 {
		if err := assignInputField(input, "OnDemandThroughput", _dynamodbOnDemandThroughput); err != nil {
			log.Errorf("invalid --on-demand-throughput: %s", err.Error())
			return
		}
	}
	if len(_dynamodbProvisionedThroughput) > 0 {
		if err := assignInputField(input, "ProvisionedThroughput", _dynamodbProvisionedThroughput); err != nil {
			log.Errorf("invalid --provisioned-throughput: %s", err.Error())
			return
		}
	}
	if len(_dynamodbReplicaUpdates) > 0 {
		if err := assignInputField(input, "ReplicaUpdates", _dynamodbReplicaUpdates); err != nil {
			log.Errorf("invalid --replica-updates: %s", err.Error())
			return
		}
	}
	if len(_dynamodbSSESpecification) > 0 {
		if err := assignInputField(input, "SSESpecification", _dynamodbSSESpecification); err != nil {
			log.Errorf("invalid --sse-specification: %s", err.Error())
			return
		}
	}
	if len(_dynamodbStreamSpecification) > 0 {
		if err := assignInputField(input, "StreamSpecification", _dynamodbStreamSpecification); err != nil {
			log.Errorf("invalid --stream-specification: %s", err.Error())
			return
		}
	}
	if len(_dynamodbTableClass) > 0 {
		if err := assignInputField(input, "TableClass", _dynamodbTableClass); err != nil {
			log.Errorf("invalid --table-class: %s", err.Error())
			return
		}
	}
	if len(_dynamodbWarmThroughput) > 0 {
		if err := assignInputField(input, "WarmThroughput", _dynamodbWarmThroughput); err != nil {
			log.Errorf("invalid --warm-throughput: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates auto scaling settings on your global tables at once.
func dynamodb_UpdateTableReplicaAutoScaling(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.UpdateTableReplicaAutoScalingInput{
		// TableName: *string, // Required
	}

	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}
	if len(_dynamodbGlobalSecondaryIndexUpdates) > 0 {
		if err := assignInputField(input, "GlobalSecondaryIndexUpdates", _dynamodbGlobalSecondaryIndexUpdates); err != nil {
			log.Errorf("invalid --global-secondary-index-updates: %s", err.Error())
			return
		}
	}
	if len(_dynamodbProvisionedWriteCapacityAutoScalingUpdate) > 0 {
		if err := assignInputField(input, "ProvisionedWriteCapacityAutoScalingUpdate", _dynamodbProvisionedWriteCapacityAutoScalingUpdate); err != nil {
			log.Errorf("invalid --provisioned-write-capacity-auto-scaling-update: %s", err.Error())
			return
		}
	}
	if len(_dynamodbReplicaUpdates) > 0 {
		if err := assignInputField(input, "ReplicaUpdates", _dynamodbReplicaUpdates); err != nil {
			log.Errorf("invalid --replica-updates: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTableReplicaAutoScaling(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The UpdateTimeToLive method enables or disables Time to Live (TTL) for the
// specified table. A successful UpdateTimeToLive call returns the current
// TimeToLiveSpecification . It can take up to one hour for the change to fully
// process. Any additional UpdateTimeToLive calls for the same table during this
// one hour duration result in a ValidationException .
//
// TTL compares the current time in epoch time format to the time stored in the
// TTL attribute of an item. If the epoch time value stored in the attribute is
// less than the current time, the item is marked as expired and subsequently
// deleted.
//
// The epoch time format is the number of seconds elapsed since 12:00:00 AM
// January 1, 1970 UTC.
//
// DynamoDB deletes expired items on a best-effort basis to ensure availability of
// throughput for other data operations.
//
// DynamoDB typically deletes expired items within two days of expiration. The
// exact duration within which an item gets deleted after expiration is specific to
// the nature of the workload. Items that have expired and not been deleted will
// still show up in reads, queries, and scans.
//
// As items are deleted, they are removed from any local secondary index and
// global secondary index immediately in the same eventually consistent way as a
// standard delete operation.
//
// For more information, see [Time To Live] in the Amazon DynamoDB Developer Guide.
//
// [Time To Live]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/TTL.html
func dynamodb_UpdateTimeToLive(cfg aws.Config, client *dynamodb.Client) {
	input := &dynamodb.UpdateTimeToLiveInput{
		// TableName: *string, // Required
		// TimeToLiveSpecification: *types.TimeToLiveSpecification, // Required
	}

	if len(_dynamodbTableName) > 0 {
		input.TableName = aws.String(_dynamodbTableName)
	}
	if len(_dynamodbTimeToLiveSpecification) > 0 {
		if err := assignInputField(input, "TimeToLiveSpecification", _dynamodbTimeToLiveSpecification); err != nil {
			log.Errorf("invalid --time-to-live-specification: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTimeToLive(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_dynamodbCmd)
	_dynamodbCmd.Flags().SortFlags = false

	_dynamodbCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_dynamodbCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_dynamodbCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_dynamodbCmd.Flags().StringVarP(&_dynamodbAttributeDefinitions, "attribute-definitions", "", "", "Attribute Definitions")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbAttributeUpdates, "attribute-updates", "", "", "Attribute Updates")
	_dynamodbCmd.Flags().StringSliceVarP(&_dynamodbAttributesToGet, "attributes-to-get", "", nil, "Attributes To Get")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbBackupArn, "backup-arn", "", "", "Backup ARN")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbBackupName, "backup-name", "", "", "Backup Name")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbBackupType, "backup-type", "", "", "Backup Type")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbBillingMode, "billing-mode", "", "", "Billing Mode")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbBillingModeOverride, "billing-mode-override", "", "", "Billing Mode Override")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbClientToken, "client-token", "", "", "Client Token")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbConditionExpression, "condition-expression", "", "", "Condition Expression")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbConditionalOperator, "conditional-operator", "", "", "Conditional Operator")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbConfirmRemoveSelfResourceAccess, "confirm-remove-self-resource-access", "", "", "Confirm Remove Self Resource Access")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbConsistentRead, "consistent-read", "", "", "Consistent Read")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbContributorInsightsAction, "contributor-insights-action", "", "", "Contributor Insights Action")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbContributorInsightsMode, "contributor-insights-mode", "", "", "Contributor Insights Mode")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbDeletionProtectionEnabled, "deletion-protection-enabled", "", "", "Deletion Protection Enabled")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbEnableKinesisStreamingConfiguration, "enable-kinesis-streaming-configuration", "", "", "Enable Kinesis Streaming Configuration")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbExclusiveStartBackupArn, "exclusive-start-backup-arn", "", "", "Exclusive Start Backup ARN")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbExclusiveStartGlobalTableName, "exclusive-start-global-table-name", "", "", "Exclusive Start Global Table Name")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbExclusiveStartKey, "exclusive-start-key", "", "", "Exclusive Start Key")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbExclusiveStartTableName, "exclusive-start-table-name", "", "", "Exclusive Start Table Name")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbExpected, "expected", "", "", "Expected")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbExpectedRevisionId, "expected-revision-id", "", "", "Expected Revision ID")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbExportArn, "export-arn", "", "", "Export ARN")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbExportFormat, "export-format", "", "", "Export Format")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbExportTime, "export-time", "", "", "Export Time")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbExportType, "export-type", "", "", "Export Type")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbExpressionAttributeNames, "expression-attribute-names", "", "", "Expression Attribute Names")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbExpressionAttributeValues, "expression-attribute-values", "", "", "Expression Attribute Values")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbFilterExpression, "filter-expression", "", "", "Filter Expression")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbGlobalSecondaryIndexOverride, "global-secondary-index-override", "", "", "Global Secondary Index Override")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbGlobalSecondaryIndexUpdates, "global-secondary-index-updates", "", "", "Global Secondary Index Updates")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbGlobalSecondaryIndexes, "global-secondary-indexes", "", "", "Global Secondary Indexes")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbGlobalTableBillingMode, "global-table-billing-mode", "", "", "Global Table Billing Mode")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbGlobalTableGlobalSecondaryIndexSettingsUpdate, "global-table-global-secondary-index-settings-update", "", "", "Global Table Global Secondary Index Settings Update")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbGlobalTableName, "global-table-name", "", "", "Global Table Name")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbGlobalTableProvisionedWriteCapacityAutoScalingSettingsUpdate, "global-table-provisioned-write-capacity-auto-scaling-settings-update", "", "", "Global Table Provisioned Write Capacity Auto Scaling Settings Update")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbGlobalTableProvisionedWriteCapacityUnits, "global-table-provisioned-write-capacity-units", "", "", "Global Table Provisioned Write Capacity Units")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbGlobalTableSettingsReplicationMode, "global-table-settings-replication-mode", "", "", "Global Table Settings Replication Mode")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbGlobalTableSourceArn, "global-table-source-arn", "", "", "Global Table Source ARN")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbGlobalTableWitnessUpdates, "global-table-witness-updates", "", "", "Global Table Witness Updates")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbImportArn, "import-arn", "", "", "Import ARN")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbIncrementalExportSpecification, "incremental-export-specification", "", "", "Incremental Export Specification")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbIndexName, "index-name", "", "", "Index Name")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbInputCompressionType, "input-compression-type", "", "", "Input Compression Type")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbInputFormat, "input-format", "", "", "Input Format")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbInputFormatOptions, "input-format-options", "", "", "Input Format Options")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbItem, "item", "", "", "Item")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbKey, "key", "", "", "Key")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbKeyConditionExpression, "key-condition-expression", "", "", "Key Condition Expression")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbKeyConditions, "key-conditions", "", "", "Key Conditions")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbKeySchema, "key-schema", "", "", "Key Schema")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbLimit, "limit", "", "", "Limit")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbLocalSecondaryIndexOverride, "local-secondary-index-override", "", "", "Local Secondary Index Override")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbLocalSecondaryIndexes, "local-secondary-indexes", "", "", "Local Secondary Indexes")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbMaxResults, "max-results", "", "", "Max Results")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbMultiRegionConsistency, "multi-region-consistency", "", "", "Multi Region Consistency")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbNextToken, "next-token", "", "", "Next Token")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbOnDemandThroughput, "on-demand-throughput", "", "", "On Demand Throughput")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbOnDemandThroughputOverride, "on-demand-throughput-override", "", "", "On Demand Throughput Override")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbPageSize, "page-size", "", "", "Page Size")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbParameters, "parameters", "", "", "Parameters")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbPointInTimeRecoverySpecification, "point-in-time-recovery-specification", "", "", "Point In Time Recovery Specification")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbPolicy, "policy", "", "", "Policy")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbProjectionExpression, "projection-expression", "", "", "Projection Expression")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbProvisionedThroughput, "provisioned-throughput", "", "", "Provisioned Throughput")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbProvisionedThroughputOverride, "provisioned-throughput-override", "", "", "Provisioned Throughput Override")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbProvisionedWriteCapacityAutoScalingUpdate, "provisioned-write-capacity-auto-scaling-update", "", "", "Provisioned Write Capacity Auto Scaling Update")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbQueryFilter, "query-filter", "", "", "Query Filter")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbRegionName, "region-name", "", "", "Region Name")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbReplicaSettingsUpdate, "replica-settings-update", "", "", "Replica Settings Update")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbReplicaUpdates, "replica-updates", "", "", "Replica Updates")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbReplicationGroup, "replication-group", "", "", "Replication Group")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbRequestItems, "request-items", "", "", "Request Items")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbResourceArn, "resource-arn", "", "", "Resource ARN")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbResourcePolicy, "resource-policy", "", "", "Resource Policy")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbRestoreDateTime, "restore-date-time", "", "", "Restore Date Time")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbReturnConsumedCapacity, "return-consumed-capacity", "", "", "Return Consumed Capacity")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbReturnItemCollectionMetrics, "return-item-collection-metrics", "", "", "Return Item Collection Metrics")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbReturnValues, "return-values", "", "", "Return Values")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbReturnValuesOnConditionCheckFailure, "return-values-on-condition-check-failure", "", "", "Return Values On Condition Check Failure")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbS3Bucket, "s3-bucket", "", "", "S3 Bucket")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbS3BucketOwner, "s3-bucket-owner", "", "", "S3 Bucket Owner")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbS3BucketSource, "s3-bucket-source", "", "", "S3 Bucket Source")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbS3Prefix, "s3-prefix", "", "", "S3 Prefix")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbS3SseAlgorithm, "s3-sse-algorithm", "", "", "S3 SSE Algorithm")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbS3SseKmsKeyId, "s3-sse-kms-key-id", "", "", "S3 SSE KMS Key ID")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbScanFilter, "scan-filter", "", "", "Scan Filter")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbScanIndexForward, "scan-index-forward", "", "", "Scan Index Forward")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbSegment, "segment", "", "", "Segment")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbSelect, "select", "", "", "Select")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbSourceTableArn, "source-table-arn", "", "", "Source Table ARN")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbSourceTableName, "source-table-name", "", "", "Source Table Name")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbSSESpecification, "sse-specification", "", "", "SSE Specification")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbSSESpecificationOverride, "sse-specification-override", "", "", "SSE Specification Override")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbStatement, "statement", "", "", "Statement")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbStatements, "statements", "", "", "Statements")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbStreamArn, "stream-arn", "", "", "Stream ARN")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbStreamSpecification, "stream-specification", "", "", "Stream Specification")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbTableArn, "table-arn", "", "", "Table ARN")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbTableClass, "table-class", "", "", "Table Class")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbTableCreationParameters, "table-creation-parameters", "", "", "Table Creation Parameters")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbTableName, "table-name", "", "", "Table Name")
	_dynamodbCmd.Flags().StringSliceVarP(&_dynamodbTagKeys, "tag-keys", "", nil, "Tag Keys")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbTags, "tags", "", "", "Tags")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbTargetTableName, "target-table-name", "", "", "Target Table Name")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbTimeRangeLowerBound, "time-range-lower-bound", "", "", "Time Range Lower Bound")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbTimeRangeUpperBound, "time-range-upper-bound", "", "", "Time Range Upper Bound")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbTimeToLiveSpecification, "time-to-live-specification", "", "", "Time To Live Specification")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbTotalSegments, "total-segments", "", "", "Total Segments")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbTransactItems, "transact-items", "", "", "Transact Items")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbTransactStatements, "transact-statements", "", "", "Transact Statements")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbUpdateExpression, "update-expression", "", "", "Update Expression")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbUpdateKinesisStreamingConfiguration, "update-kinesis-streaming-configuration", "", "", "Update Kinesis Streaming Configuration")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbUseLatestRestorableTime, "use-latest-restorable-time", "", "", "Use Latest Restorable Time")
	_dynamodbCmd.Flags().StringVarP(&_dynamodbWarmThroughput, "warm-throughput", "", "", "Warm Throughput")

	_dynamodbCmd.Flags().BoolVarP(&_dynamodbBatchExecuteStatement, "batch-execute-statement", "", false, "Batch Execute Statement")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbBatchGetItem, "batch-get-item", "", false, "Batch Get Item")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbBatchWriteItem, "batch-write-item", "", false, "Batch Write Item")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbCreateBackup, "create-backup", "", false, "Create Backup")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbCreateGlobalTable, "create-global-table", "", false, "Create Global Table")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbCreateTable, "create-table", "", false, "Create Table")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbDeleteBackup, "delete-backup", "", false, "Delete Backup")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbDeleteItem, "delete-item", "", false, "Delete Item")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbDeleteTable, "delete-table", "", false, "Delete Table")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbDescribeBackup, "describe-backup", "", false, "Describe Backup")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbDescribeContinuousBackups, "describe-continuous-backups", "", false, "Describe Continuous Backups")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbDescribeContributorInsights, "describe-contributor-insights", "", false, "Describe Contributor Insights")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbDescribeEndpoints, "describe-endpoints", "", false, "Describe Endpoints")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbDescribeExport, "describe-export", "", false, "Describe Export")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbDescribeGlobalTable, "describe-global-table", "", false, "Describe Global Table")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbDescribeGlobalTableSettings, "describe-global-table-settings", "", false, "Describe Global Table Settings")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbDescribeImport, "describe-import", "", false, "Describe Import")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbDescribeKinesisStreamingDestination, "describe-kinesis-streaming-destination", "", false, "Describe Kinesis Streaming Destination")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbDescribeLimits, "describe-limits", "", false, "Describe Limits")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbDescribeTable, "describe-table", "", false, "Describe Table")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbDescribeTableReplicaAutoScaling, "describe-table-replica-auto-scaling", "", false, "Describe Table Replica Auto Scaling")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbDescribeTimeToLive, "describe-time-to-live", "", false, "Describe Time To Live")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbDisableKinesisStreamingDestination, "disable-kinesis-streaming-destination", "", false, "Disable Kinesis Streaming Destination")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbEnableKinesisStreamingDestination, "enable-kinesis-streaming-destination", "", false, "Enable Kinesis Streaming Destination")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbExecuteStatement, "execute-statement", "", false, "Execute Statement")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbExecuteTransaction, "execute-transaction", "", false, "Execute Transaction")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbExportTableToPointInTime, "export-table-to-point-in-time", "", false, "Export Table To Point In Time")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbGetItem, "get-item", "", false, "Get Item")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbGetResourcePolicy, "get-resource-policy", "", false, "Get Resource Policy")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbImportTable, "import-table", "", false, "Import Table")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbListBackups, "list-backups", "", false, "List Backups")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbListContributorInsights, "list-contributor-insights", "", false, "List Contributor Insights")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbListExports, "list-exports", "", false, "List Exports")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbListGlobalTables, "list-global-tables", "", false, "List Global Tables")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbListImports, "list-imports", "", false, "List Imports")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbListTables, "list-tables", "", false, "List Tables")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbListTagsOfResource, "list-tags-of-resource", "", false, "List Tags Of Resource")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbPutItem, "put-item", "", false, "Put Item")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbPutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbQuery, "query", "", false, "Query")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbRestoreTableFromBackup, "restore-table-from-backup", "", false, "Restore Table From Backup")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbRestoreTableToPointInTime, "restore-table-to-point-in-time", "", false, "Restore Table To Point In Time")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbScan, "scan", "", false, "Scan")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbTagResource, "tag-resource", "", false, "Tag Resource")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbTransactGetItems, "transact-get-items", "", false, "Transact Get Items")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbTransactWriteItems, "transact-write-items", "", false, "Transact Write Items")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbUntagResource, "untag-resource", "", false, "Untag Resource")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbUpdateContinuousBackups, "update-continuous-backups", "", false, "Update Continuous Backups")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbUpdateContributorInsights, "update-contributor-insights", "", false, "Update Contributor Insights")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbUpdateGlobalTable, "update-global-table", "", false, "Update Global Table")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbUpdateGlobalTableSettings, "update-global-table-settings", "", false, "Update Global Table Settings")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbUpdateItem, "update-item", "", false, "Update Item")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbUpdateKinesisStreamingDestination, "update-kinesis-streaming-destination", "", false, "Update Kinesis Streaming Destination")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbUpdateTable, "update-table", "", false, "Update Table")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbUpdateTableReplicaAutoScaling, "update-table-replica-auto-scaling", "", false, "Update Table Replica Auto Scaling")
	_dynamodbCmd.Flags().BoolVarP(&_dynamodbUpdateTimeToLive, "update-time-to-live", "", false, "Update Time To Live")

}
