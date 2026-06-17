package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rdsdata"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rdsdataCmd represents the rdsdata command
var _rdsdataCmd = &cobra.Command{
	Use:   "rdsdata",
	Short: "AWS rdsdata CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := rdsdata.NewFromConfig(cfg)
		if _rdsdataBatchExecuteStatement {
			rdsdata_BatchExecuteStatement(cfg, client)
			return
		}
		if _rdsdataBeginTransaction {
			rdsdata_BeginTransaction(cfg, client)
			return
		}
		if _rdsdataCommitTransaction {
			rdsdata_CommitTransaction(cfg, client)
			return
		}
		if _rdsdataExecuteSql {
			rdsdata_ExecuteSql(cfg, client)
			return
		}
		if _rdsdataExecuteStatement {
			rdsdata_ExecuteStatement(cfg, client)
			return
		}
		if _rdsdataRollbackTransaction {
			rdsdata_RollbackTransaction(cfg, client)
			return
		}

	},
}

var (
	_rdsdataBatchExecuteStatement bool
	_rdsdataBeginTransaction      bool
	_rdsdataCommitTransaction     bool
	_rdsdataExecuteSql            bool
	_rdsdataExecuteStatement      bool
	_rdsdataRollbackTransaction   bool

	_rdsdataAwsSecretStoreArn      string
	_rdsdataContinueAfterTimeout   string
	_rdsdataDatabase               string
	_rdsdataDbClusterOrInstanceArn string
	_rdsdataFormatRecordsAs        string
	_rdsdataIncludeResultMetadata  string
	_rdsdataParameterSets          string
	_rdsdataParameters             string
	_rdsdataResourceArn            string
	_rdsdataResultSetOptions       string
	_rdsdataSchema                 string
	_rdsdataSecretArn              string
	_rdsdataSql                    string
	_rdsdataSqlStatements          string
	_rdsdataTransactionId          string
)

// Runs a batch SQL statement over an array of data.
// You can run bulk update and insert operations for multiple records using a DML
// statement with different parameter sets. Bulk operations can provide a
// significant performance improvement over individual insert and update
// operations.
//
// If a call isn't part of a transaction because it doesn't include the
// transactionID parameter, changes that result from the call are committed
// automatically.
//
// There isn't a fixed upper limit on the number of parameter sets. However, the
// maximum size of the HTTP request submitted through the Data API is 4 MiB. If the
// request exceeds this limit, the Data API returns an error and doesn't process
// the request. This 4-MiB limit includes the size of the HTTP headers and the JSON
// notation in the request. Thus, the number of parameter sets that you can include
// depends on a combination of factors, such as the size of the SQL statement and
// the size of each parameter set.
//
// The response size limit is 1 MiB. If the call returns more than 1 MiB of
// response data, the call is terminated.
func rdsdata_BatchExecuteStatement(cfg aws.Config, client *rdsdata.Client) {
	input := &rdsdata.BatchExecuteStatementInput{
		// ResourceArn: *string, // Required
		// SecretArn: *string, // Required
		// Sql: *string, // Required
	}

	if len(_rdsdataResourceArn) > 0 {
		input.ResourceArn = aws.String(_rdsdataResourceArn)
	}
	if len(_rdsdataSecretArn) > 0 {
		input.SecretArn = aws.String(_rdsdataSecretArn)
	}
	if len(_rdsdataSql) > 0 {
		input.Sql = aws.String(_rdsdataSql)
	}
	if len(_rdsdataDatabase) > 0 {
		input.Database = aws.String(_rdsdataDatabase)
	}
	if len(_rdsdataParameterSets) > 0 {
		if err := assignInputField(input, "ParameterSets", _rdsdataParameterSets); err != nil {
			log.Errorf("invalid --parameter-sets: %s", err.Error())
			return
		}
	}
	if len(_rdsdataSchema) > 0 {
		input.Schema = aws.String(_rdsdataSchema)
	}
	if len(_rdsdataTransactionId) > 0 {
		input.TransactionId = aws.String(_rdsdataTransactionId)
	}

	if resp, err := client.BatchExecuteStatement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a SQL transaction.
// A transaction can run for a maximum of 24 hours. A transaction is terminated
// and rolled back automatically after 24 hours.
//
// A transaction times out if no calls use its transaction ID in three minutes. If
// a transaction times out before it's committed, it's rolled back automatically.
//
// For Aurora MySQL, DDL statements inside a transaction cause an implicit commit.
// We recommend that you run each MySQL DDL statement in a separate
// ExecuteStatement call with continueAfterTimeout enabled.
func rdsdata_BeginTransaction(cfg aws.Config, client *rdsdata.Client) {
	input := &rdsdata.BeginTransactionInput{
		// ResourceArn: *string, // Required
		// SecretArn: *string, // Required
	}

	if len(_rdsdataResourceArn) > 0 {
		input.ResourceArn = aws.String(_rdsdataResourceArn)
	}
	if len(_rdsdataSecretArn) > 0 {
		input.SecretArn = aws.String(_rdsdataSecretArn)
	}
	if len(_rdsdataDatabase) > 0 {
		input.Database = aws.String(_rdsdataDatabase)
	}
	if len(_rdsdataSchema) > 0 {
		input.Schema = aws.String(_rdsdataSchema)
	}

	if resp, err := client.BeginTransaction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Ends a SQL transaction started with the BeginTransaction operation and commits
// the changes.
func rdsdata_CommitTransaction(cfg aws.Config, client *rdsdata.Client) {
	input := &rdsdata.CommitTransactionInput{
		// ResourceArn: *string, // Required
		// SecretArn: *string, // Required
		// TransactionId: *string, // Required
	}

	if len(_rdsdataResourceArn) > 0 {
		input.ResourceArn = aws.String(_rdsdataResourceArn)
	}
	if len(_rdsdataSecretArn) > 0 {
		input.SecretArn = aws.String(_rdsdataSecretArn)
	}
	if len(_rdsdataTransactionId) > 0 {
		input.TransactionId = aws.String(_rdsdataTransactionId)
	}

	if resp, err := client.CommitTransaction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Runs one or more SQL statements.
// This operation isn't supported for Aurora Serverless v2 and provisioned DB
// clusters. For Aurora Serverless v1 DB clusters, the operation is deprecated. Use
// the BatchExecuteStatement or ExecuteStatement operation.
//
// Deprecated: The ExecuteSql API is deprecated, please use the ExecuteStatement
// API.
func rdsdata_ExecuteSql(cfg aws.Config, client *rdsdata.Client) {
	input := &rdsdata.ExecuteSqlInput{
		// AwsSecretStoreArn: *string, // Required
		// DbClusterOrInstanceArn: *string, // Required
		// SqlStatements: *string, // Required
	}

	if len(_rdsdataAwsSecretStoreArn) > 0 {
		input.AwsSecretStoreArn = aws.String(_rdsdataAwsSecretStoreArn)
	}
	if len(_rdsdataDbClusterOrInstanceArn) > 0 {
		input.DbClusterOrInstanceArn = aws.String(_rdsdataDbClusterOrInstanceArn)
	}
	if len(_rdsdataSqlStatements) > 0 {
		input.SqlStatements = aws.String(_rdsdataSqlStatements)
	}
	if len(_rdsdataDatabase) > 0 {
		input.Database = aws.String(_rdsdataDatabase)
	}
	if len(_rdsdataSchema) > 0 {
		input.Schema = aws.String(_rdsdataSchema)
	}

	if resp, err := client.ExecuteSql(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Runs a SQL statement against a database.
// If a call isn't part of a transaction because it doesn't include the
// transactionID parameter, changes that result from the call are committed
// automatically.
//
// If the binary response data from the database is more than 1 MB, the call is
// terminated.
func rdsdata_ExecuteStatement(cfg aws.Config, client *rdsdata.Client) {
	input := &rdsdata.ExecuteStatementInput{
		// ResourceArn: *string, // Required
		// SecretArn: *string, // Required
		// Sql: *string, // Required
	}

	if len(_rdsdataResourceArn) > 0 {
		input.ResourceArn = aws.String(_rdsdataResourceArn)
	}
	if len(_rdsdataSecretArn) > 0 {
		input.SecretArn = aws.String(_rdsdataSecretArn)
	}
	if len(_rdsdataSql) > 0 {
		input.Sql = aws.String(_rdsdataSql)
	}
	if len(_rdsdataContinueAfterTimeout) > 0 {
		if err := assignInputField(input, "ContinueAfterTimeout", _rdsdataContinueAfterTimeout); err != nil {
			log.Errorf("invalid --continue-after-timeout: %s", err.Error())
			return
		}
	}
	if len(_rdsdataDatabase) > 0 {
		input.Database = aws.String(_rdsdataDatabase)
	}
	if len(_rdsdataFormatRecordsAs) > 0 {
		if err := assignInputField(input, "FormatRecordsAs", _rdsdataFormatRecordsAs); err != nil {
			log.Errorf("invalid --format-records-as: %s", err.Error())
			return
		}
	}
	if len(_rdsdataIncludeResultMetadata) > 0 {
		if err := assignInputField(input, "IncludeResultMetadata", _rdsdataIncludeResultMetadata); err != nil {
			log.Errorf("invalid --include-result-metadata: %s", err.Error())
			return
		}
	}
	if len(_rdsdataParameters) > 0 {
		if err := assignInputField(input, "Parameters", _rdsdataParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_rdsdataResultSetOptions) > 0 {
		if err := assignInputField(input, "ResultSetOptions", _rdsdataResultSetOptions); err != nil {
			log.Errorf("invalid --result-set-options: %s", err.Error())
			return
		}
	}
	if len(_rdsdataSchema) > 0 {
		input.Schema = aws.String(_rdsdataSchema)
	}
	if len(_rdsdataTransactionId) > 0 {
		input.TransactionId = aws.String(_rdsdataTransactionId)
	}

	if resp, err := client.ExecuteStatement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Performs a rollback of a transaction. Rolling back a transaction cancels its
// changes.
func rdsdata_RollbackTransaction(cfg aws.Config, client *rdsdata.Client) {
	input := &rdsdata.RollbackTransactionInput{
		// ResourceArn: *string, // Required
		// SecretArn: *string, // Required
		// TransactionId: *string, // Required
	}

	if len(_rdsdataResourceArn) > 0 {
		input.ResourceArn = aws.String(_rdsdataResourceArn)
	}
	if len(_rdsdataSecretArn) > 0 {
		input.SecretArn = aws.String(_rdsdataSecretArn)
	}
	if len(_rdsdataTransactionId) > 0 {
		input.TransactionId = aws.String(_rdsdataTransactionId)
	}

	if resp, err := client.RollbackTransaction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_rdsdataCmd)
	_rdsdataCmd.Flags().SortFlags = false

	_rdsdataCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_rdsdataCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_rdsdataCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_rdsdataCmd.Flags().StringVarP(&_rdsdataAwsSecretStoreArn, "aws-secret-store-arn", "", "", "AWS Secret Store ARN")
	_rdsdataCmd.Flags().StringVarP(&_rdsdataContinueAfterTimeout, "continue-after-timeout", "", "", "Continue After Timeout")
	_rdsdataCmd.Flags().StringVarP(&_rdsdataDatabase, "database", "", "", "Database")
	_rdsdataCmd.Flags().StringVarP(&_rdsdataDbClusterOrInstanceArn, "db-cluster-or-instance-arn", "", "", "DB Cluster Or Instance ARN")
	_rdsdataCmd.Flags().StringVarP(&_rdsdataFormatRecordsAs, "format-records-as", "", "", "Format Records As")
	_rdsdataCmd.Flags().StringVarP(&_rdsdataIncludeResultMetadata, "include-result-metadata", "", "", "Include Result Metadata")
	_rdsdataCmd.Flags().StringVarP(&_rdsdataParameterSets, "parameter-sets", "", "", "Parameter Sets")
	_rdsdataCmd.Flags().StringVarP(&_rdsdataParameters, "parameters", "", "", "Parameters")
	_rdsdataCmd.Flags().StringVarP(&_rdsdataResourceArn, "resource-arn", "", "", "Resource ARN")
	_rdsdataCmd.Flags().StringVarP(&_rdsdataResultSetOptions, "result-set-options", "", "", "Result Set Options")
	_rdsdataCmd.Flags().StringVarP(&_rdsdataSchema, "schema", "", "", "Schema")
	_rdsdataCmd.Flags().StringVarP(&_rdsdataSecretArn, "secret-arn", "", "", "Secret ARN")
	_rdsdataCmd.Flags().StringVarP(&_rdsdataSql, "sql", "", "", "Sql")
	_rdsdataCmd.Flags().StringVarP(&_rdsdataSqlStatements, "sql-statements", "", "", "Sql Statements")
	_rdsdataCmd.Flags().StringVarP(&_rdsdataTransactionId, "transaction-id", "", "", "Transaction ID")

	_rdsdataCmd.Flags().BoolVarP(&_rdsdataBatchExecuteStatement, "batch-execute-statement", "", false, "Batch Execute Statement")
	_rdsdataCmd.Flags().BoolVarP(&_rdsdataBeginTransaction, "begin-transaction", "", false, "Begin Transaction")
	_rdsdataCmd.Flags().BoolVarP(&_rdsdataCommitTransaction, "commit-transaction", "", false, "Commit Transaction")
	_rdsdataCmd.Flags().BoolVarP(&_rdsdataExecuteSql, "execute-sql", "", false, "Execute Sql")
	_rdsdataCmd.Flags().BoolVarP(&_rdsdataExecuteStatement, "execute-statement", "", false, "Execute Statement")
	_rdsdataCmd.Flags().BoolVarP(&_rdsdataRollbackTransaction, "rollback-transaction", "", false, "Rollback Transaction")

}
