package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// timestreamwriteCmd represents the timestreamwrite command
var _timestreamwriteCmd = &cobra.Command{
	Use:   "timestreamwrite",
	Short: "AWS timestreamwrite CLI",
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
		client := timestreamwrite.NewFromConfig(cfg)
		if _timestreamwriteCreateBatchLoadTask {
			timestreamwrite_CreateBatchLoadTask(cfg, client)
			return
		}
		if _timestreamwriteCreateDatabase {
			timestreamwrite_CreateDatabase(cfg, client)
			return
		}
		if _timestreamwriteCreateTable {
			timestreamwrite_CreateTable(cfg, client)
			return
		}
		if _timestreamwriteDeleteDatabase {
			timestreamwrite_DeleteDatabase(cfg, client)
			return
		}
		if _timestreamwriteDeleteTable {
			timestreamwrite_DeleteTable(cfg, client)
			return
		}
		if _timestreamwriteDescribeBatchLoadTask {
			timestreamwrite_DescribeBatchLoadTask(cfg, client)
			return
		}
		if _timestreamwriteDescribeDatabase {
			timestreamwrite_DescribeDatabase(cfg, client)
			return
		}
		if _timestreamwriteDescribeEndpoints {
			timestreamwrite_DescribeEndpoints(cfg, client)
			return
		}
		if _timestreamwriteDescribeTable {
			timestreamwrite_DescribeTable(cfg, client)
			return
		}
		if _timestreamwriteListBatchLoadTasks {
			timestreamwrite_ListBatchLoadTasks(cfg, client)
			return
		}
		if _timestreamwriteListDatabases {
			timestreamwrite_ListDatabases(cfg, client)
			return
		}
		if _timestreamwriteListTables {
			timestreamwrite_ListTables(cfg, client)
			return
		}
		if _timestreamwriteListTagsForResource {
			timestreamwrite_ListTagsForResource(cfg, client)
			return
		}
		if _timestreamwriteResumeBatchLoadTask {
			timestreamwrite_ResumeBatchLoadTask(cfg, client)
			return
		}
		if _timestreamwriteTagResource {
			timestreamwrite_TagResource(cfg, client)
			return
		}
		if _timestreamwriteUntagResource {
			timestreamwrite_UntagResource(cfg, client)
			return
		}
		if _timestreamwriteUpdateDatabase {
			timestreamwrite_UpdateDatabase(cfg, client)
			return
		}
		if _timestreamwriteUpdateTable {
			timestreamwrite_UpdateTable(cfg, client)
			return
		}
		if _timestreamwriteWriteRecords {
			timestreamwrite_WriteRecords(cfg, client)
			return
		}

	},
}

var (
	_timestreamwriteCreateBatchLoadTask   bool
	_timestreamwriteCreateDatabase        bool
	_timestreamwriteCreateTable           bool
	_timestreamwriteDeleteDatabase        bool
	_timestreamwriteDeleteTable           bool
	_timestreamwriteDescribeBatchLoadTask bool
	_timestreamwriteDescribeDatabase      bool
	_timestreamwriteDescribeEndpoints     bool
	_timestreamwriteDescribeTable         bool
	_timestreamwriteListBatchLoadTasks    bool
	_timestreamwriteListDatabases         bool
	_timestreamwriteListTables            bool
	_timestreamwriteListTagsForResource   bool
	_timestreamwriteResumeBatchLoadTask   bool
	_timestreamwriteTagResource           bool
	_timestreamwriteUntagResource         bool
	_timestreamwriteUpdateDatabase        bool
	_timestreamwriteUpdateTable           bool
	_timestreamwriteWriteRecords          bool

	_timestreamwriteClientToken                  string
	_timestreamwriteCommonAttributes             string
	_timestreamwriteDataModelConfiguration       string
	_timestreamwriteDataSourceConfiguration      string
	_timestreamwriteDatabaseName                 string
	_timestreamwriteKmsKeyId                     string
	_timestreamwriteMagneticStoreWriteProperties string
	_timestreamwriteMaxResults                   string
	_timestreamwriteNextToken                    string
	_timestreamwriteRecordVersion                string
	_timestreamwriteRecords                      string
	_timestreamwriteReportConfiguration          string
	_timestreamwriteResourceARN                  string
	_timestreamwriteRetentionProperties          string
	_timestreamwriteSchema                       string
	_timestreamwriteTableName                    string
	_timestreamwriteTagKeys                      []string
	_timestreamwriteTags                         string
	_timestreamwriteTargetDatabaseName           string
	_timestreamwriteTargetTableName              string
	_timestreamwriteTaskId                       string
	_timestreamwriteTaskStatus                   string
)

// Creates a new Timestream batch load task. A batch load task processes data from
// a CSV source in an S3 location and writes to a Timestream table. A mapping from
// source to target is defined in a batch load task. Errors and events are written
// to a report at an S3 location. For the report, if the KMS key is not specified,
// the report will be encrypted with an S3 managed key when SSE_S3 is the option.
// Otherwise an error is thrown. For more information, see [Amazon Web Services managed keys]. [Service quotas apply]. For details, see [code sample].
//
// [Service quotas apply]: https://docs.aws.amazon.com/timestream/latest/developerguide/ts-limits.html
// [code sample]: https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.create-batch-load.html
// [Amazon Web Services managed keys]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk
func timestreamwrite_CreateBatchLoadTask(cfg aws.Config, client *timestreamwrite.Client) {
	input := &timestreamwrite.CreateBatchLoadTaskInput{
		// DataSourceConfiguration: *types.DataSourceConfiguration, // Required
		// ReportConfiguration: *types.ReportConfiguration, // Required
		// TargetDatabaseName: *string, // Required
		// TargetTableName: *string, // Required
	}

	if len(_timestreamwriteDataSourceConfiguration) > 0 {
		if err := assignInputField(input, "DataSourceConfiguration", _timestreamwriteDataSourceConfiguration); err != nil {
			log.Errorf("invalid --data-source-configuration: %s", err.Error())
			return
		}
	}
	if len(_timestreamwriteReportConfiguration) > 0 {
		if err := assignInputField(input, "ReportConfiguration", _timestreamwriteReportConfiguration); err != nil {
			log.Errorf("invalid --report-configuration: %s", err.Error())
			return
		}
	}
	if len(_timestreamwriteTargetDatabaseName) > 0 {
		input.TargetDatabaseName = aws.String(_timestreamwriteTargetDatabaseName)
	}
	if len(_timestreamwriteTargetTableName) > 0 {
		input.TargetTableName = aws.String(_timestreamwriteTargetTableName)
	}
	if len(_timestreamwriteClientToken) > 0 {
		input.ClientToken = aws.String(_timestreamwriteClientToken)
	}
	if len(_timestreamwriteDataModelConfiguration) > 0 {
		if err := assignInputField(input, "DataModelConfiguration", _timestreamwriteDataModelConfiguration); err != nil {
			log.Errorf("invalid --data-model-configuration: %s", err.Error())
			return
		}
	}
	if len(_timestreamwriteRecordVersion) > 0 {
		if err := assignInputField(input, "RecordVersion", _timestreamwriteRecordVersion); err != nil {
			log.Errorf("invalid --record-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBatchLoadTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Timestream database. If the KMS key is not specified, the
// database will be encrypted with a Timestream managed KMS key located in your
// account. For more information, see [Amazon Web Services managed keys]. [Service quotas apply]. For details, see [code sample].
//
// [Service quotas apply]: https://docs.aws.amazon.com/timestream/latest/developerguide/ts-limits.html
// [code sample]: https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.create-db.html
// [Amazon Web Services managed keys]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk
func timestreamwrite_CreateDatabase(cfg aws.Config, client *timestreamwrite.Client) {
	input := &timestreamwrite.CreateDatabaseInput{
		// DatabaseName: *string, // Required
	}

	if len(_timestreamwriteDatabaseName) > 0 {
		input.DatabaseName = aws.String(_timestreamwriteDatabaseName)
	}
	if len(_timestreamwriteKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_timestreamwriteKmsKeyId)
	}
	if len(_timestreamwriteTags) > 0 {
		if err := assignInputField(input, "Tags", _timestreamwriteTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDatabase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a new table to an existing database in your account. In an Amazon Web
// Services account, table names must be at least unique within each Region if they
// are in the same database. You might have identical table names in the same
// Region if the tables are in separate databases. While creating the table, you
// must specify the table name, database name, and the retention properties. [Service quotas apply]. See [code sample]
// for details.
//
// [Service quotas apply]: https://docs.aws.amazon.com/timestream/latest/developerguide/ts-limits.html
// [code sample]: https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.create-table.html
func timestreamwrite_CreateTable(cfg aws.Config, client *timestreamwrite.Client) {
	input := &timestreamwrite.CreateTableInput{
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_timestreamwriteDatabaseName) > 0 {
		input.DatabaseName = aws.String(_timestreamwriteDatabaseName)
	}
	if len(_timestreamwriteTableName) > 0 {
		input.TableName = aws.String(_timestreamwriteTableName)
	}
	if len(_timestreamwriteMagneticStoreWriteProperties) > 0 {
		if err := assignInputField(input, "MagneticStoreWriteProperties", _timestreamwriteMagneticStoreWriteProperties); err != nil {
			log.Errorf("invalid --magnetic-store-write-properties: %s", err.Error())
			return
		}
	}
	if len(_timestreamwriteRetentionProperties) > 0 {
		if err := assignInputField(input, "RetentionProperties", _timestreamwriteRetentionProperties); err != nil {
			log.Errorf("invalid --retention-properties: %s", err.Error())
			return
		}
	}
	if len(_timestreamwriteSchema) > 0 {
		if err := assignInputField(input, "Schema", _timestreamwriteSchema); err != nil {
			log.Errorf("invalid --schema: %s", err.Error())
			return
		}
	}
	if len(_timestreamwriteTags) > 0 {
		if err := assignInputField(input, "Tags", _timestreamwriteTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
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

// Deletes a given Timestream database. This is an irreversible operation. After a
// database is deleted, the time-series data from its tables cannot be recovered.
//
// All tables in the database must be deleted first, or a ValidationException
// error will be thrown.
//
// Due to the nature of distributed retries, the operation can return either
// success or a ResourceNotFoundException. Clients should consider them equivalent.
//
// See [code sample] for details.
//
// [code sample]: https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.delete-db.html
func timestreamwrite_DeleteDatabase(cfg aws.Config, client *timestreamwrite.Client) {
	input := &timestreamwrite.DeleteDatabaseInput{
		// DatabaseName: *string, // Required
	}

	if len(_timestreamwriteDatabaseName) > 0 {
		input.DatabaseName = aws.String(_timestreamwriteDatabaseName)
	}

	if resp, err := client.DeleteDatabase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a given Timestream table. This is an irreversible operation. After a
// Timestream database table is deleted, the time-series data stored in the table
// cannot be recovered.
//
// Due to the nature of distributed retries, the operation can return either
// success or a ResourceNotFoundException. Clients should consider them equivalent.
//
// See [code sample] for details.
//
// [code sample]: https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.delete-table.html
func timestreamwrite_DeleteTable(cfg aws.Config, client *timestreamwrite.Client) {
	input := &timestreamwrite.DeleteTableInput{
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_timestreamwriteDatabaseName) > 0 {
		input.DatabaseName = aws.String(_timestreamwriteDatabaseName)
	}
	if len(_timestreamwriteTableName) > 0 {
		input.TableName = aws.String(_timestreamwriteTableName)
	}

	if resp, err := client.DeleteTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the batch load task, including configurations,
// mappings, progress, and other details. [Service quotas apply]. See [code sample] for details.
//
// [Service quotas apply]: https://docs.aws.amazon.com/timestream/latest/developerguide/ts-limits.html
// [code sample]: https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.describe-batch-load.html
func timestreamwrite_DescribeBatchLoadTask(cfg aws.Config, client *timestreamwrite.Client) {
	input := &timestreamwrite.DescribeBatchLoadTaskInput{
		// TaskId: *string, // Required
	}

	if len(_timestreamwriteTaskId) > 0 {
		input.TaskId = aws.String(_timestreamwriteTaskId)
	}

	if resp, err := client.DescribeBatchLoadTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the database, including the database name, time that
// the database was created, and the total number of tables found within the
// database. [Service quotas apply]. See [code sample] for details.
//
// [Service quotas apply]: https://docs.aws.amazon.com/timestream/latest/developerguide/ts-limits.html
// [code sample]: https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.describe-db.html
func timestreamwrite_DescribeDatabase(cfg aws.Config, client *timestreamwrite.Client) {
	input := &timestreamwrite.DescribeDatabaseInput{
		// DatabaseName: *string, // Required
	}

	if len(_timestreamwriteDatabaseName) > 0 {
		input.DatabaseName = aws.String(_timestreamwriteDatabaseName)
	}

	if resp, err := client.DescribeDatabase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of available endpoints to make Timestream API calls against.
// This API operation is available through both the Write and Query APIs.
//
// Because the Timestream SDKs are designed to transparently work with the
// service’s architecture, including the management and mapping of the service
// endpoints, we don't recommend that you use this API operation unless:
//
// - You are using [VPC endpoints (Amazon Web Services PrivateLink) with Timestream]
//
// - Your application uses a programming language that does not yet have SDK
// support
//
// - You require better control over the client-side implementation
//
// For detailed information on how and when to use and implement
// DescribeEndpoints, see [The Endpoint Discovery Pattern].
//
// [The Endpoint Discovery Pattern]: https://docs.aws.amazon.com/timestream/latest/developerguide/Using.API.html#Using-API.endpoint-discovery
// [VPC endpoints (Amazon Web Services PrivateLink) with Timestream]: https://docs.aws.amazon.com/timestream/latest/developerguide/VPCEndpoints
func timestreamwrite_DescribeEndpoints(cfg aws.Config, client *timestreamwrite.Client) {
	input := &timestreamwrite.DescribeEndpointsInput{}

	if resp, err := client.DescribeEndpoints(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the table, including the table name, database name,
// retention duration of the memory store and the magnetic store. [Service quotas apply]. See [code sample] for
// details.
//
// [Service quotas apply]: https://docs.aws.amazon.com/timestream/latest/developerguide/ts-limits.html
// [code sample]: https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.describe-table.html
func timestreamwrite_DescribeTable(cfg aws.Config, client *timestreamwrite.Client) {
	input := &timestreamwrite.DescribeTableInput{
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_timestreamwriteDatabaseName) > 0 {
		input.DatabaseName = aws.String(_timestreamwriteDatabaseName)
	}
	if len(_timestreamwriteTableName) > 0 {
		input.TableName = aws.String(_timestreamwriteTableName)
	}

	if resp, err := client.DescribeTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a list of batch load tasks, along with the name, status, when the task
// is resumable until, and other details. See [code sample]for details.
//
// [code sample]: https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.list-batch-load-tasks.html
func timestreamwrite_ListBatchLoadTasks(cfg aws.Config, client *timestreamwrite.Client) {
	input := &timestreamwrite.ListBatchLoadTasksInput{}

	if len(_timestreamwriteMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _timestreamwriteMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_timestreamwriteNextToken) > 0 {
		input.NextToken = aws.String(_timestreamwriteNextToken)
	}
	if len(_timestreamwriteTaskStatus) > 0 {
		if err := assignInputField(input, "TaskStatus", _timestreamwriteTaskStatus); err != nil {
			log.Errorf("invalid --task-status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListBatchLoadTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*timestreamwrite.ListBatchLoadTasksOutput
	p := timestreamwrite.NewListBatchLoadTasksPaginator(client, input)
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

// Returns a list of your Timestream databases. [Service quotas apply]. See [code sample] for details.
//
// [Service quotas apply]: https://docs.aws.amazon.com/timestream/latest/developerguide/ts-limits.html
// [code sample]: https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.list-db.html
func timestreamwrite_ListDatabases(cfg aws.Config, client *timestreamwrite.Client) {
	input := &timestreamwrite.ListDatabasesInput{}

	if len(_timestreamwriteMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _timestreamwriteMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_timestreamwriteNextToken) > 0 {
		input.NextToken = aws.String(_timestreamwriteNextToken)
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

	var results []*timestreamwrite.ListDatabasesOutput
	p := timestreamwrite.NewListDatabasesPaginator(client, input)
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

// Provides a list of tables, along with the name, status, and retention
// properties of each table. See [code sample]for details.
//
// [code sample]: https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.list-table.html
func timestreamwrite_ListTables(cfg aws.Config, client *timestreamwrite.Client) {
	input := &timestreamwrite.ListTablesInput{}

	if len(_timestreamwriteDatabaseName) > 0 {
		input.DatabaseName = aws.String(_timestreamwriteDatabaseName)
	}
	if len(_timestreamwriteMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _timestreamwriteMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_timestreamwriteNextToken) > 0 {
		input.NextToken = aws.String(_timestreamwriteNextToken)
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

	var results []*timestreamwrite.ListTablesOutput
	p := timestreamwrite.NewListTablesPaginator(client, input)
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

// Lists all tags on a Timestream resource.
func timestreamwrite_ListTagsForResource(cfg aws.Config, client *timestreamwrite.Client) {
	input := &timestreamwrite.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_timestreamwriteResourceARN) > 0 {
		input.ResourceARN = aws.String(_timestreamwriteResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func timestreamwrite_ResumeBatchLoadTask(cfg aws.Config, client *timestreamwrite.Client) {
	input := &timestreamwrite.ResumeBatchLoadTaskInput{
		// TaskId: *string, // Required
	}

	if len(_timestreamwriteTaskId) > 0 {
		input.TaskId = aws.String(_timestreamwriteTaskId)
	}

	if resp, err := client.ResumeBatchLoadTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a set of tags with a Timestream resource. You can then activate
// these user-defined tags so that they appear on the Billing and Cost Management
// console for cost allocation tracking.
func timestreamwrite_TagResource(cfg aws.Config, client *timestreamwrite.Client) {
	input := &timestreamwrite.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_timestreamwriteResourceARN) > 0 {
		input.ResourceARN = aws.String(_timestreamwriteResourceARN)
	}
	if len(_timestreamwriteTags) > 0 {
		if err := assignInputField(input, "Tags", _timestreamwriteTags); err != nil {
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

// Removes the association of tags from a Timestream resource.
func timestreamwrite_UntagResource(cfg aws.Config, client *timestreamwrite.Client) {
	input := &timestreamwrite.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_timestreamwriteResourceARN) > 0 {
		input.ResourceARN = aws.String(_timestreamwriteResourceARN)
	}
	if len(_timestreamwriteTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _timestreamwriteTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the KMS key for an existing database. While updating the database,
// you must specify the database name and the identifier of the new KMS key to be
// used ( KmsKeyId ). If there are any concurrent UpdateDatabase requests, first
// writer wins.
//
// See [code sample] for details.
//
// [code sample]: https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.update-db.html
func timestreamwrite_UpdateDatabase(cfg aws.Config, client *timestreamwrite.Client) {
	input := &timestreamwrite.UpdateDatabaseInput{
		// DatabaseName: *string, // Required
		// KmsKeyId: *string, // Required
	}

	if len(_timestreamwriteDatabaseName) > 0 {
		input.DatabaseName = aws.String(_timestreamwriteDatabaseName)
	}
	if len(_timestreamwriteKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_timestreamwriteKmsKeyId)
	}

	if resp, err := client.UpdateDatabase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the retention duration of the memory store and magnetic store for your
// Timestream table. Note that the change in retention duration takes effect
// immediately. For example, if the retention period of the memory store was
// initially set to 2 hours and then changed to 24 hours, the memory store will be
// capable of holding 24 hours of data, but will be populated with 24 hours of data
// 22 hours after this change was made. Timestream does not retrieve data from the
// magnetic store to populate the memory store.
//
// See [code sample] for details.
//
// [code sample]: https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.update-table.html
func timestreamwrite_UpdateTable(cfg aws.Config, client *timestreamwrite.Client) {
	input := &timestreamwrite.UpdateTableInput{
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_timestreamwriteDatabaseName) > 0 {
		input.DatabaseName = aws.String(_timestreamwriteDatabaseName)
	}
	if len(_timestreamwriteTableName) > 0 {
		input.TableName = aws.String(_timestreamwriteTableName)
	}
	if len(_timestreamwriteMagneticStoreWriteProperties) > 0 {
		if err := assignInputField(input, "MagneticStoreWriteProperties", _timestreamwriteMagneticStoreWriteProperties); err != nil {
			log.Errorf("invalid --magnetic-store-write-properties: %s", err.Error())
			return
		}
	}
	if len(_timestreamwriteRetentionProperties) > 0 {
		if err := assignInputField(input, "RetentionProperties", _timestreamwriteRetentionProperties); err != nil {
			log.Errorf("invalid --retention-properties: %s", err.Error())
			return
		}
	}
	if len(_timestreamwriteSchema) > 0 {
		if err := assignInputField(input, "Schema", _timestreamwriteSchema); err != nil {
			log.Errorf("invalid --schema: %s", err.Error())
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

// Enables you to write your time-series data into Timestream. You can specify a
// single data point or a batch of data points to be inserted into the system.
// Timestream offers you a flexible schema that auto detects the column names and
// data types for your Timestream tables based on the dimension names and data
// types of the data points you specify when invoking writes into the database.
//
// Timestream supports eventual consistency read semantics. This means that when
// you query data immediately after writing a batch of data into Timestream, the
// query results might not reflect the results of a recently completed write
// operation. The results may also include some stale data. If you repeat the query
// request after a short time, the results should return the latest data. [Service quotas apply].
//
// See [code sample] for details.
//
// # Upserts
//
// You can use the Version parameter in a WriteRecords request to update data
// points. Timestream tracks a version number with each record. Version defaults
// to 1 when it's not specified for the record in the request. Timestream updates
// an existing record’s measure value along with its Version when it receives a
// write request with a higher Version number for that record. When it receives an
// update request where the measure value is the same as that of the existing
// record, Timestream still updates Version , if it is greater than the existing
// value of Version . You can update a data point as many times as desired, as long
// as the value of Version continuously increases.
//
// For example, suppose you write a new record without indicating Version in the
// request. Timestream stores this record, and set Version to 1 . Now, suppose you
// try to update this record with a WriteRecords request of the same record with a
// different measure value but, like before, do not provide Version . In this case,
// Timestream will reject this update with a RejectedRecordsException since the
// updated record’s version is not greater than the existing value of Version.
//
// However, if you were to resend the update request with Version set to 2 ,
// Timestream would then succeed in updating the record’s value, and the Version
// would be set to 2 . Next, suppose you sent a WriteRecords request with this
// same record and an identical measure value, but with Version set to 3 . In this
// case, Timestream would only update Version to 3 . Any further updates would need
// to send a version number greater than 3 , or the update requests would receive a
// RejectedRecordsException .
//
// [Service quotas apply]: https://docs.aws.amazon.com/timestream/latest/developerguide/ts-limits.html
// [code sample]: https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.write.html
func timestreamwrite_WriteRecords(cfg aws.Config, client *timestreamwrite.Client) {
	input := &timestreamwrite.WriteRecordsInput{
		// DatabaseName: *string, // Required
		// Records: []types.Record, // Required
		// TableName: *string, // Required
	}

	if len(_timestreamwriteDatabaseName) > 0 {
		input.DatabaseName = aws.String(_timestreamwriteDatabaseName)
	}
	if len(_timestreamwriteRecords) > 0 {
		if err := assignInputField(input, "Records", _timestreamwriteRecords); err != nil {
			log.Errorf("invalid --records: %s", err.Error())
			return
		}
	}
	if len(_timestreamwriteTableName) > 0 {
		input.TableName = aws.String(_timestreamwriteTableName)
	}
	if len(_timestreamwriteCommonAttributes) > 0 {
		if err := assignInputField(input, "CommonAttributes", _timestreamwriteCommonAttributes); err != nil {
			log.Errorf("invalid --common-attributes: %s", err.Error())
			return
		}
	}

	if resp, err := client.WriteRecords(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_timestreamwriteCmd)
	_timestreamwriteCmd.Flags().SortFlags = false

	_timestreamwriteCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_timestreamwriteCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_timestreamwriteCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_timestreamwriteCmd.Flags().StringVarP(&_timestreamwriteClientToken, "client-token", "", "", "Client Token")
	_timestreamwriteCmd.Flags().StringVarP(&_timestreamwriteCommonAttributes, "common-attributes", "", "", "Common Attributes")
	_timestreamwriteCmd.Flags().StringVarP(&_timestreamwriteDataModelConfiguration, "data-model-configuration", "", "", "Data Model Configuration")
	_timestreamwriteCmd.Flags().StringVarP(&_timestreamwriteDataSourceConfiguration, "data-source-configuration", "", "", "Data Source Configuration")
	_timestreamwriteCmd.Flags().StringVarP(&_timestreamwriteDatabaseName, "database-name", "", "", "Database Name")
	_timestreamwriteCmd.Flags().StringVarP(&_timestreamwriteKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_timestreamwriteCmd.Flags().StringVarP(&_timestreamwriteMagneticStoreWriteProperties, "magnetic-store-write-properties", "", "", "Magnetic Store Write Properties")
	_timestreamwriteCmd.Flags().StringVarP(&_timestreamwriteMaxResults, "max-results", "", "", "Max Results")
	_timestreamwriteCmd.Flags().StringVarP(&_timestreamwriteNextToken, "next-token", "", "", "Next Token")
	_timestreamwriteCmd.Flags().StringVarP(&_timestreamwriteRecordVersion, "record-version", "", "", "Record Version")
	_timestreamwriteCmd.Flags().StringVarP(&_timestreamwriteRecords, "records", "", "", "Records")
	_timestreamwriteCmd.Flags().StringVarP(&_timestreamwriteReportConfiguration, "report-configuration", "", "", "Report Configuration")
	_timestreamwriteCmd.Flags().StringVarP(&_timestreamwriteResourceARN, "resource-arn", "", "", "Resource ARN")
	_timestreamwriteCmd.Flags().StringVarP(&_timestreamwriteRetentionProperties, "retention-properties", "", "", "Retention Properties")
	_timestreamwriteCmd.Flags().StringVarP(&_timestreamwriteSchema, "schema", "", "", "Schema")
	_timestreamwriteCmd.Flags().StringVarP(&_timestreamwriteTableName, "table-name", "", "", "Table Name")
	_timestreamwriteCmd.Flags().StringSliceVarP(&_timestreamwriteTagKeys, "tag-keys", "", nil, "Tag Keys")
	_timestreamwriteCmd.Flags().StringVarP(&_timestreamwriteTags, "tags", "", "", "Tags")
	_timestreamwriteCmd.Flags().StringVarP(&_timestreamwriteTargetDatabaseName, "target-database-name", "", "", "Target Database Name")
	_timestreamwriteCmd.Flags().StringVarP(&_timestreamwriteTargetTableName, "target-table-name", "", "", "Target Table Name")
	_timestreamwriteCmd.Flags().StringVarP(&_timestreamwriteTaskId, "task-id", "", "", "Task ID")
	_timestreamwriteCmd.Flags().StringVarP(&_timestreamwriteTaskStatus, "task-status", "", "", "Task Status")

	_timestreamwriteCmd.Flags().BoolVarP(&_timestreamwriteCreateBatchLoadTask, "create-batch-load-task", "", false, "Create Batch Load Task")
	_timestreamwriteCmd.Flags().BoolVarP(&_timestreamwriteCreateDatabase, "create-database", "", false, "Create Database")
	_timestreamwriteCmd.Flags().BoolVarP(&_timestreamwriteCreateTable, "create-table", "", false, "Create Table")
	_timestreamwriteCmd.Flags().BoolVarP(&_timestreamwriteDeleteDatabase, "delete-database", "", false, "Delete Database")
	_timestreamwriteCmd.Flags().BoolVarP(&_timestreamwriteDeleteTable, "delete-table", "", false, "Delete Table")
	_timestreamwriteCmd.Flags().BoolVarP(&_timestreamwriteDescribeBatchLoadTask, "describe-batch-load-task", "", false, "Describe Batch Load Task")
	_timestreamwriteCmd.Flags().BoolVarP(&_timestreamwriteDescribeDatabase, "describe-database", "", false, "Describe Database")
	_timestreamwriteCmd.Flags().BoolVarP(&_timestreamwriteDescribeEndpoints, "describe-endpoints", "", false, "Describe Endpoints")
	_timestreamwriteCmd.Flags().BoolVarP(&_timestreamwriteDescribeTable, "describe-table", "", false, "Describe Table")
	_timestreamwriteCmd.Flags().BoolVarP(&_timestreamwriteListBatchLoadTasks, "list-batch-load-tasks", "", false, "List Batch Load Tasks")
	_timestreamwriteCmd.Flags().BoolVarP(&_timestreamwriteListDatabases, "list-databases", "", false, "List Databases")
	_timestreamwriteCmd.Flags().BoolVarP(&_timestreamwriteListTables, "list-tables", "", false, "List Tables")
	_timestreamwriteCmd.Flags().BoolVarP(&_timestreamwriteListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_timestreamwriteCmd.Flags().BoolVarP(&_timestreamwriteResumeBatchLoadTask, "resume-batch-load-task", "", false, "Resume Batch Load Task")
	_timestreamwriteCmd.Flags().BoolVarP(&_timestreamwriteTagResource, "tag-resource", "", false, "Tag Resource")
	_timestreamwriteCmd.Flags().BoolVarP(&_timestreamwriteUntagResource, "untag-resource", "", false, "Untag Resource")
	_timestreamwriteCmd.Flags().BoolVarP(&_timestreamwriteUpdateDatabase, "update-database", "", false, "Update Database")
	_timestreamwriteCmd.Flags().BoolVarP(&_timestreamwriteUpdateTable, "update-table", "", false, "Update Table")
	_timestreamwriteCmd.Flags().BoolVarP(&_timestreamwriteWriteRecords, "write-records", "", false, "Write Records")

}
