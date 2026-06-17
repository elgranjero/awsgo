package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshiftdata"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// redshiftdataCmd represents the redshiftdata command
var _redshiftdataCmd = &cobra.Command{
	Use:   "redshiftdata",
	Short: "AWS redshiftdata CLI",
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
		client := redshiftdata.NewFromConfig(cfg)
		if _redshiftdataBatchExecuteStatement {
			redshiftdata_BatchExecuteStatement(cfg, client)
			return
		}
		if _redshiftdataCancelStatement {
			redshiftdata_CancelStatement(cfg, client)
			return
		}
		if _redshiftdataDescribeStatement {
			redshiftdata_DescribeStatement(cfg, client)
			return
		}
		if _redshiftdataDescribeTable {
			redshiftdata_DescribeTable(cfg, client)
			return
		}
		if _redshiftdataExecuteStatement {
			redshiftdata_ExecuteStatement(cfg, client)
			return
		}
		if _redshiftdataGetStatementResult {
			redshiftdata_GetStatementResult(cfg, client)
			return
		}
		if _redshiftdataGetStatementResultV2 {
			redshiftdata_GetStatementResultV2(cfg, client)
			return
		}
		if _redshiftdataListDatabases {
			redshiftdata_ListDatabases(cfg, client)
			return
		}
		if _redshiftdataListSchemas {
			redshiftdata_ListSchemas(cfg, client)
			return
		}
		if _redshiftdataListStatements {
			redshiftdata_ListStatements(cfg, client)
			return
		}
		if _redshiftdataListTables {
			redshiftdata_ListTables(cfg, client)
			return
		}

	},
}

var (
	_redshiftdataBatchExecuteStatement bool
	_redshiftdataCancelStatement       bool
	_redshiftdataDescribeStatement     bool
	_redshiftdataDescribeTable         bool
	_redshiftdataExecuteStatement      bool
	_redshiftdataGetStatementResult    bool
	_redshiftdataGetStatementResultV2  bool
	_redshiftdataListDatabases         bool
	_redshiftdataListSchemas           bool
	_redshiftdataListStatements        bool
	_redshiftdataListTables            bool

	_redshiftdataClientToken             string
	_redshiftdataClusterIdentifier       string
	_redshiftdataConnectedDatabase       string
	_redshiftdataDatabase                string
	_redshiftdataDbUser                  string
	_redshiftdataId                      string
	_redshiftdataMaxResults              string
	_redshiftdataNextToken               string
	_redshiftdataParameters              string
	_redshiftdataResultFormat            string
	_redshiftdataRoleLevel               string
	_redshiftdataSchema                  string
	_redshiftdataSchemaPattern           string
	_redshiftdataSecretArn               string
	_redshiftdataSessionId               string
	_redshiftdataSessionKeepAliveSeconds string
	_redshiftdataSql                     string
	_redshiftdataSqls                    []string
	_redshiftdataStatementName           string
	_redshiftdataStatus                  string
	_redshiftdataTable                   string
	_redshiftdataTablePattern            string
	_redshiftdataWithEvent               string
	_redshiftdataWorkgroupName           string
)

// Runs one or more SQL statements, which can be data manipulation language (DML)
// or data definition language (DDL). Depending on the authorization method, use
// one of the following combinations of request parameters:
//
// - Secrets Manager - when connecting to a cluster, provide the secret-arn of a
// secret stored in Secrets Manager which has username and password . The
// specified secret contains credentials to connect to the database you specify.
// When you are connecting to a cluster, you also supply the database name, If you
// provide a cluster identifier ( dbClusterIdentifier ), it must match the
// cluster identifier stored in the secret. When you are connecting to a serverless
// workgroup, you also supply the database name.
//
// - Temporary credentials - when connecting to your data warehouse, choose one
// of the following options:
//
// - When connecting to a serverless workgroup, specify the workgroup name and
// database name. The database user name is derived from the IAM identity. For
// example, arn:iam::123456789012:user:foo has the database user name IAM:foo .
// Also, permission to call the redshift-serverless:GetCredentials operation is
// required.
//
// - When connecting to a cluster as an IAM identity, specify the cluster
// identifier and the database name. The database user name is derived from the IAM
// identity. For example, arn:iam::123456789012:user:foo has the database user
// name IAM:foo . Also, permission to call the
// redshift:GetClusterCredentialsWithIAM operation is required.
//
// - When connecting to a cluster as a database user, specify the cluster
// identifier, the database name, and the database user name. Also, permission to
// call the redshift:GetClusterCredentials operation is required.
//
// For more information about the Amazon Redshift Data API and CLI usage examples,
// see [Using the Amazon Redshift Data API]in the Amazon Redshift Management Guide.
//
// [Using the Amazon Redshift Data API]: https://docs.aws.amazon.com/redshift/latest/mgmt/data-api.html
func redshiftdata_BatchExecuteStatement(cfg aws.Config, client *redshiftdata.Client) {
	input := &redshiftdata.BatchExecuteStatementInput{
		// Sqls: []string, // Required
	}

	if len(_redshiftdataSqls) > 0 {
		input.Sqls = append([]string(nil), _redshiftdataSqls...)
	}
	if len(_redshiftdataClientToken) > 0 {
		input.ClientToken = aws.String(_redshiftdataClientToken)
	}
	if len(_redshiftdataClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftdataClusterIdentifier)
	}
	if len(_redshiftdataDatabase) > 0 {
		input.Database = aws.String(_redshiftdataDatabase)
	}
	if len(_redshiftdataDbUser) > 0 {
		input.DbUser = aws.String(_redshiftdataDbUser)
	}
	if len(_redshiftdataResultFormat) > 0 {
		if err := assignInputField(input, "ResultFormat", _redshiftdataResultFormat); err != nil {
			log.Errorf("invalid --result-format: %s", err.Error())
			return
		}
	}
	if len(_redshiftdataSecretArn) > 0 {
		input.SecretArn = aws.String(_redshiftdataSecretArn)
	}
	if len(_redshiftdataSessionId) > 0 {
		input.SessionId = aws.String(_redshiftdataSessionId)
	}
	if len(_redshiftdataSessionKeepAliveSeconds) > 0 {
		if err := assignInputField(input, "SessionKeepAliveSeconds", _redshiftdataSessionKeepAliveSeconds); err != nil {
			log.Errorf("invalid --session-keep-alive-seconds: %s", err.Error())
			return
		}
	}
	if len(_redshiftdataStatementName) > 0 {
		input.StatementName = aws.String(_redshiftdataStatementName)
	}
	if len(_redshiftdataWithEvent) > 0 {
		if err := assignInputField(input, "WithEvent", _redshiftdataWithEvent); err != nil {
			log.Errorf("invalid --with-event: %s", err.Error())
			return
		}
	}
	if len(_redshiftdataWorkgroupName) > 0 {
		input.WorkgroupName = aws.String(_redshiftdataWorkgroupName)
	}

	if resp, err := client.BatchExecuteStatement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a running query. To be canceled, a query must be running.
// For more information about the Amazon Redshift Data API and CLI usage examples,
// see [Using the Amazon Redshift Data API]in the Amazon Redshift Management Guide.
//
// [Using the Amazon Redshift Data API]: https://docs.aws.amazon.com/redshift/latest/mgmt/data-api.html
func redshiftdata_CancelStatement(cfg aws.Config, client *redshiftdata.Client) {
	input := &redshiftdata.CancelStatementInput{
		// Id: *string, // Required
	}

	if len(_redshiftdataId) > 0 {
		input.Id = aws.String(_redshiftdataId)
	}

	if resp, err := client.CancelStatement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the details about a specific instance when a query was run by the
// Amazon Redshift Data API. The information includes when the query started, when
// it finished, the query status, the number of rows returned, and the SQL
// statement.
//
// For more information about the Amazon Redshift Data API and CLI usage examples,
// see [Using the Amazon Redshift Data API]in the Amazon Redshift Management Guide.
//
// [Using the Amazon Redshift Data API]: https://docs.aws.amazon.com/redshift/latest/mgmt/data-api.html
func redshiftdata_DescribeStatement(cfg aws.Config, client *redshiftdata.Client) {
	input := &redshiftdata.DescribeStatementInput{
		// Id: *string, // Required
	}

	if len(_redshiftdataId) > 0 {
		input.Id = aws.String(_redshiftdataId)
	}

	if resp, err := client.DescribeStatement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the detailed information about a table from metadata in the cluster.
// The information includes its columns. A token is returned to page through the
// column list. Depending on the authorization method, use one of the following
// combinations of request parameters:
//
// - Secrets Manager - when connecting to a cluster, provide the secret-arn of a
// secret stored in Secrets Manager which has username and password . The
// specified secret contains credentials to connect to the database you specify.
// When you are connecting to a cluster, you also supply the database name, If you
// provide a cluster identifier ( dbClusterIdentifier ), it must match the
// cluster identifier stored in the secret. When you are connecting to a serverless
// workgroup, you also supply the database name.
//
// - Temporary credentials - when connecting to your data warehouse, choose one
// of the following options:
//
// - When connecting to a serverless workgroup, specify the workgroup name and
// database name. The database user name is derived from the IAM identity. For
// example, arn:iam::123456789012:user:foo has the database user name IAM:foo .
// Also, permission to call the redshift-serverless:GetCredentials operation is
// required.
//
// - When connecting to a cluster as an IAM identity, specify the cluster
// identifier and the database name. The database user name is derived from the IAM
// identity. For example, arn:iam::123456789012:user:foo has the database user
// name IAM:foo . Also, permission to call the
// redshift:GetClusterCredentialsWithIAM operation is required.
//
// - When connecting to a cluster as a database user, specify the cluster
// identifier, the database name, and the database user name. Also, permission to
// call the redshift:GetClusterCredentials operation is required.
//
// For more information about the Amazon Redshift Data API and CLI usage examples,
// see [Using the Amazon Redshift Data API]in the Amazon Redshift Management Guide.
//
// [Using the Amazon Redshift Data API]: https://docs.aws.amazon.com/redshift/latest/mgmt/data-api.html
func redshiftdata_DescribeTable(cfg aws.Config, client *redshiftdata.Client) {
	input := &redshiftdata.DescribeTableInput{
		// Database: *string, // Required
	}

	if len(_redshiftdataDatabase) > 0 {
		input.Database = aws.String(_redshiftdataDatabase)
	}
	if len(_redshiftdataClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftdataClusterIdentifier)
	}
	if len(_redshiftdataConnectedDatabase) > 0 {
		input.ConnectedDatabase = aws.String(_redshiftdataConnectedDatabase)
	}
	if len(_redshiftdataDbUser) > 0 {
		input.DbUser = aws.String(_redshiftdataDbUser)
	}
	if len(_redshiftdataMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _redshiftdataMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_redshiftdataNextToken) > 0 {
		input.NextToken = aws.String(_redshiftdataNextToken)
	}
	if len(_redshiftdataSchema) > 0 {
		input.Schema = aws.String(_redshiftdataSchema)
	}
	if len(_redshiftdataSecretArn) > 0 {
		input.SecretArn = aws.String(_redshiftdataSecretArn)
	}
	if len(_redshiftdataTable) > 0 {
		input.Table = aws.String(_redshiftdataTable)
	}
	if len(_redshiftdataWorkgroupName) > 0 {
		input.WorkgroupName = aws.String(_redshiftdataWorkgroupName)
	}

	if disablePaginator() {
		if resp, err := client.DescribeTable(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshiftdata.DescribeTableOutput
	p := redshiftdata.NewDescribeTablePaginator(client, input)
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

// Runs an SQL statement, which can be data manipulation language (DML) or data
// definition language (DDL). This statement must be a single SQL statement.
// Depending on the authorization method, use one of the following combinations of
// request parameters:
//
// - Secrets Manager - when connecting to a cluster, provide the secret-arn of a
// secret stored in Secrets Manager which has username and password . The
// specified secret contains credentials to connect to the database you specify.
// When you are connecting to a cluster, you also supply the database name, If you
// provide a cluster identifier ( dbClusterIdentifier ), it must match the
// cluster identifier stored in the secret. When you are connecting to a serverless
// workgroup, you also supply the database name.
//
// - Temporary credentials - when connecting to your data warehouse, choose one
// of the following options:
//
// - When connecting to a serverless workgroup, specify the workgroup name and
// database name. The database user name is derived from the IAM identity. For
// example, arn:iam::123456789012:user:foo has the database user name IAM:foo .
// Also, permission to call the redshift-serverless:GetCredentials operation is
// required.
//
// - When connecting to a cluster as an IAM identity, specify the cluster
// identifier and the database name. The database user name is derived from the IAM
// identity. For example, arn:iam::123456789012:user:foo has the database user
// name IAM:foo . Also, permission to call the
// redshift:GetClusterCredentialsWithIAM operation is required.
//
// - When connecting to a cluster as a database user, specify the cluster
// identifier, the database name, and the database user name. Also, permission to
// call the redshift:GetClusterCredentials operation is required.
//
// For more information about the Amazon Redshift Data API and CLI usage examples,
// see [Using the Amazon Redshift Data API]in the Amazon Redshift Management Guide.
//
// [Using the Amazon Redshift Data API]: https://docs.aws.amazon.com/redshift/latest/mgmt/data-api.html
func redshiftdata_ExecuteStatement(cfg aws.Config, client *redshiftdata.Client) {
	input := &redshiftdata.ExecuteStatementInput{
		// Sql: *string, // Required
	}

	if len(_redshiftdataSql) > 0 {
		input.Sql = aws.String(_redshiftdataSql)
	}
	if len(_redshiftdataClientToken) > 0 {
		input.ClientToken = aws.String(_redshiftdataClientToken)
	}
	if len(_redshiftdataClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftdataClusterIdentifier)
	}
	if len(_redshiftdataDatabase) > 0 {
		input.Database = aws.String(_redshiftdataDatabase)
	}
	if len(_redshiftdataDbUser) > 0 {
		input.DbUser = aws.String(_redshiftdataDbUser)
	}
	if len(_redshiftdataParameters) > 0 {
		if err := assignInputField(input, "Parameters", _redshiftdataParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_redshiftdataResultFormat) > 0 {
		if err := assignInputField(input, "ResultFormat", _redshiftdataResultFormat); err != nil {
			log.Errorf("invalid --result-format: %s", err.Error())
			return
		}
	}
	if len(_redshiftdataSecretArn) > 0 {
		input.SecretArn = aws.String(_redshiftdataSecretArn)
	}
	if len(_redshiftdataSessionId) > 0 {
		input.SessionId = aws.String(_redshiftdataSessionId)
	}
	if len(_redshiftdataSessionKeepAliveSeconds) > 0 {
		if err := assignInputField(input, "SessionKeepAliveSeconds", _redshiftdataSessionKeepAliveSeconds); err != nil {
			log.Errorf("invalid --session-keep-alive-seconds: %s", err.Error())
			return
		}
	}
	if len(_redshiftdataStatementName) > 0 {
		input.StatementName = aws.String(_redshiftdataStatementName)
	}
	if len(_redshiftdataWithEvent) > 0 {
		if err := assignInputField(input, "WithEvent", _redshiftdataWithEvent); err != nil {
			log.Errorf("invalid --with-event: %s", err.Error())
			return
		}
	}
	if len(_redshiftdataWorkgroupName) > 0 {
		input.WorkgroupName = aws.String(_redshiftdataWorkgroupName)
	}

	if resp, err := client.ExecuteStatement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Fetches the temporarily cached result of an SQL statement in JSON format. The
// ExecuteStatement or BatchExecuteStatement operation that ran the SQL statement
// must have specified ResultFormat as JSON , or let the format default to JSON. A
// token is returned to page through the statement results.
//
// For more information about the Amazon Redshift Data API and CLI usage examples,
// see [Using the Amazon Redshift Data API]in the Amazon Redshift Management Guide.
//
// [Using the Amazon Redshift Data API]: https://docs.aws.amazon.com/redshift/latest/mgmt/data-api.html
func redshiftdata_GetStatementResult(cfg aws.Config, client *redshiftdata.Client) {
	input := &redshiftdata.GetStatementResultInput{
		// Id: *string, // Required
	}

	if len(_redshiftdataId) > 0 {
		input.Id = aws.String(_redshiftdataId)
	}
	if len(_redshiftdataNextToken) > 0 {
		input.NextToken = aws.String(_redshiftdataNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetStatementResult(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshiftdata.GetStatementResultOutput
	p := redshiftdata.NewGetStatementResultPaginator(client, input)
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

// Fetches the temporarily cached result of an SQL statement in CSV format. The
// ExecuteStatement or BatchExecuteStatement operation that ran the SQL statement
// must have specified ResultFormat as CSV . A token is returned to page through
// the statement results.
//
// For more information about the Amazon Redshift Data API and CLI usage examples,
// see [Using the Amazon Redshift Data API]in the Amazon Redshift Management Guide.
//
// [Using the Amazon Redshift Data API]: https://docs.aws.amazon.com/redshift/latest/mgmt/data-api.html
func redshiftdata_GetStatementResultV2(cfg aws.Config, client *redshiftdata.Client) {
	input := &redshiftdata.GetStatementResultV2Input{
		// Id: *string, // Required
	}

	if len(_redshiftdataId) > 0 {
		input.Id = aws.String(_redshiftdataId)
	}
	if len(_redshiftdataNextToken) > 0 {
		input.NextToken = aws.String(_redshiftdataNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetStatementResultV2(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshiftdata.GetStatementResultV2Output
	p := redshiftdata.NewGetStatementResultV2Paginator(client, input)
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

// List the databases in a cluster. A token is returned to page through the
// database list. Depending on the authorization method, use one of the following
// combinations of request parameters:
//
// - Secrets Manager - when connecting to a cluster, provide the secret-arn of a
// secret stored in Secrets Manager which has username and password . The
// specified secret contains credentials to connect to the database you specify.
// When you are connecting to a cluster, you also supply the database name, If you
// provide a cluster identifier ( dbClusterIdentifier ), it must match the
// cluster identifier stored in the secret. When you are connecting to a serverless
// workgroup, you also supply the database name.
//
// - Temporary credentials - when connecting to your data warehouse, choose one
// of the following options:
//
// - When connecting to a serverless workgroup, specify the workgroup name and
// database name. The database user name is derived from the IAM identity. For
// example, arn:iam::123456789012:user:foo has the database user name IAM:foo .
// Also, permission to call the redshift-serverless:GetCredentials operation is
// required.
//
// - When connecting to a cluster as an IAM identity, specify the cluster
// identifier and the database name. The database user name is derived from the IAM
// identity. For example, arn:iam::123456789012:user:foo has the database user
// name IAM:foo . Also, permission to call the
// redshift:GetClusterCredentialsWithIAM operation is required.
//
// - When connecting to a cluster as a database user, specify the cluster
// identifier, the database name, and the database user name. Also, permission to
// call the redshift:GetClusterCredentials operation is required.
//
// For more information about the Amazon Redshift Data API and CLI usage examples,
// see [Using the Amazon Redshift Data API]in the Amazon Redshift Management Guide.
//
// [Using the Amazon Redshift Data API]: https://docs.aws.amazon.com/redshift/latest/mgmt/data-api.html
func redshiftdata_ListDatabases(cfg aws.Config, client *redshiftdata.Client) {
	input := &redshiftdata.ListDatabasesInput{
		// Database: *string, // Required
	}

	if len(_redshiftdataDatabase) > 0 {
		input.Database = aws.String(_redshiftdataDatabase)
	}
	if len(_redshiftdataClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftdataClusterIdentifier)
	}
	if len(_redshiftdataDbUser) > 0 {
		input.DbUser = aws.String(_redshiftdataDbUser)
	}
	if len(_redshiftdataMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _redshiftdataMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_redshiftdataNextToken) > 0 {
		input.NextToken = aws.String(_redshiftdataNextToken)
	}
	if len(_redshiftdataSecretArn) > 0 {
		input.SecretArn = aws.String(_redshiftdataSecretArn)
	}
	if len(_redshiftdataWorkgroupName) > 0 {
		input.WorkgroupName = aws.String(_redshiftdataWorkgroupName)
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

	var results []*redshiftdata.ListDatabasesOutput
	p := redshiftdata.NewListDatabasesPaginator(client, input)
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

// Lists the schemas in a database. A token is returned to page through the schema
// list. Depending on the authorization method, use one of the following
// combinations of request parameters:
//
// - Secrets Manager - when connecting to a cluster, provide the secret-arn of a
// secret stored in Secrets Manager which has username and password . The
// specified secret contains credentials to connect to the database you specify.
// When you are connecting to a cluster, you also supply the database name, If you
// provide a cluster identifier ( dbClusterIdentifier ), it must match the
// cluster identifier stored in the secret. When you are connecting to a serverless
// workgroup, you also supply the database name.
//
// - Temporary credentials - when connecting to your data warehouse, choose one
// of the following options:
//
// - When connecting to a serverless workgroup, specify the workgroup name and
// database name. The database user name is derived from the IAM identity. For
// example, arn:iam::123456789012:user:foo has the database user name IAM:foo .
// Also, permission to call the redshift-serverless:GetCredentials operation is
// required.
//
// - When connecting to a cluster as an IAM identity, specify the cluster
// identifier and the database name. The database user name is derived from the IAM
// identity. For example, arn:iam::123456789012:user:foo has the database user
// name IAM:foo . Also, permission to call the
// redshift:GetClusterCredentialsWithIAM operation is required.
//
// - When connecting to a cluster as a database user, specify the cluster
// identifier, the database name, and the database user name. Also, permission to
// call the redshift:GetClusterCredentials operation is required.
//
// For more information about the Amazon Redshift Data API and CLI usage examples,
// see [Using the Amazon Redshift Data API]in the Amazon Redshift Management Guide.
//
// [Using the Amazon Redshift Data API]: https://docs.aws.amazon.com/redshift/latest/mgmt/data-api.html
func redshiftdata_ListSchemas(cfg aws.Config, client *redshiftdata.Client) {
	input := &redshiftdata.ListSchemasInput{
		// Database: *string, // Required
	}

	if len(_redshiftdataDatabase) > 0 {
		input.Database = aws.String(_redshiftdataDatabase)
	}
	if len(_redshiftdataClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftdataClusterIdentifier)
	}
	if len(_redshiftdataConnectedDatabase) > 0 {
		input.ConnectedDatabase = aws.String(_redshiftdataConnectedDatabase)
	}
	if len(_redshiftdataDbUser) > 0 {
		input.DbUser = aws.String(_redshiftdataDbUser)
	}
	if len(_redshiftdataMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _redshiftdataMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_redshiftdataNextToken) > 0 {
		input.NextToken = aws.String(_redshiftdataNextToken)
	}
	if len(_redshiftdataSchemaPattern) > 0 {
		input.SchemaPattern = aws.String(_redshiftdataSchemaPattern)
	}
	if len(_redshiftdataSecretArn) > 0 {
		input.SecretArn = aws.String(_redshiftdataSecretArn)
	}
	if len(_redshiftdataWorkgroupName) > 0 {
		input.WorkgroupName = aws.String(_redshiftdataWorkgroupName)
	}

	if disablePaginator() {
		if resp, err := client.ListSchemas(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshiftdata.ListSchemasOutput
	p := redshiftdata.NewListSchemasPaginator(client, input)
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

// List of SQL statements. By default, only finished statements are shown. A token
// is returned to page through the statement list.
//
// When you use identity-enhanced role sessions to list statements, you must
// provide either the cluster-identifier or workgroup-name parameter. This ensures
// that the IdC user can only access the Amazon Redshift IdC applications they are
// assigned. For more information, see [Trusted identity propagation overview].
//
// For more information about the Amazon Redshift Data API and CLI usage examples,
// see [Using the Amazon Redshift Data API]in the Amazon Redshift Management Guide.
//
// [Using the Amazon Redshift Data API]: https://docs.aws.amazon.com/redshift/latest/mgmt/data-api.html
// [Trusted identity propagation overview]: https://docs.aws.amazon.com/singlesignon/latest/userguide/trustedidentitypropagation-overview.html
func redshiftdata_ListStatements(cfg aws.Config, client *redshiftdata.Client) {
	input := &redshiftdata.ListStatementsInput{}

	if len(_redshiftdataClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftdataClusterIdentifier)
	}
	if len(_redshiftdataDatabase) > 0 {
		input.Database = aws.String(_redshiftdataDatabase)
	}
	if len(_redshiftdataMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _redshiftdataMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_redshiftdataNextToken) > 0 {
		input.NextToken = aws.String(_redshiftdataNextToken)
	}
	if len(_redshiftdataRoleLevel) > 0 {
		if err := assignInputField(input, "RoleLevel", _redshiftdataRoleLevel); err != nil {
			log.Errorf("invalid --role-level: %s", err.Error())
			return
		}
	}
	if len(_redshiftdataStatementName) > 0 {
		input.StatementName = aws.String(_redshiftdataStatementName)
	}
	if len(_redshiftdataStatus) > 0 {
		if err := assignInputField(input, "Status", _redshiftdataStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_redshiftdataWorkgroupName) > 0 {
		input.WorkgroupName = aws.String(_redshiftdataWorkgroupName)
	}

	if disablePaginator() {
		if resp, err := client.ListStatements(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshiftdata.ListStatementsOutput
	p := redshiftdata.NewListStatementsPaginator(client, input)
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

// List the tables in a database. If neither SchemaPattern nor TablePattern are
// specified, then all tables in the database are returned. A token is returned to
// page through the table list. Depending on the authorization method, use one of
// the following combinations of request parameters:
//
// - Secrets Manager - when connecting to a cluster, provide the secret-arn of a
// secret stored in Secrets Manager which has username and password . The
// specified secret contains credentials to connect to the database you specify.
// When you are connecting to a cluster, you also supply the database name, If you
// provide a cluster identifier ( dbClusterIdentifier ), it must match the
// cluster identifier stored in the secret. When you are connecting to a serverless
// workgroup, you also supply the database name.
//
// - Temporary credentials - when connecting to your data warehouse, choose one
// of the following options:
//
// - When connecting to a serverless workgroup, specify the workgroup name and
// database name. The database user name is derived from the IAM identity. For
// example, arn:iam::123456789012:user:foo has the database user name IAM:foo .
// Also, permission to call the redshift-serverless:GetCredentials operation is
// required.
//
// - When connecting to a cluster as an IAM identity, specify the cluster
// identifier and the database name. The database user name is derived from the IAM
// identity. For example, arn:iam::123456789012:user:foo has the database user
// name IAM:foo . Also, permission to call the
// redshift:GetClusterCredentialsWithIAM operation is required.
//
// - When connecting to a cluster as a database user, specify the cluster
// identifier, the database name, and the database user name. Also, permission to
// call the redshift:GetClusterCredentials operation is required.
//
// For more information about the Amazon Redshift Data API and CLI usage examples,
// see [Using the Amazon Redshift Data API]in the Amazon Redshift Management Guide.
//
// [Using the Amazon Redshift Data API]: https://docs.aws.amazon.com/redshift/latest/mgmt/data-api.html
func redshiftdata_ListTables(cfg aws.Config, client *redshiftdata.Client) {
	input := &redshiftdata.ListTablesInput{
		// Database: *string, // Required
	}

	if len(_redshiftdataDatabase) > 0 {
		input.Database = aws.String(_redshiftdataDatabase)
	}
	if len(_redshiftdataClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftdataClusterIdentifier)
	}
	if len(_redshiftdataConnectedDatabase) > 0 {
		input.ConnectedDatabase = aws.String(_redshiftdataConnectedDatabase)
	}
	if len(_redshiftdataDbUser) > 0 {
		input.DbUser = aws.String(_redshiftdataDbUser)
	}
	if len(_redshiftdataMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _redshiftdataMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_redshiftdataNextToken) > 0 {
		input.NextToken = aws.String(_redshiftdataNextToken)
	}
	if len(_redshiftdataSchemaPattern) > 0 {
		input.SchemaPattern = aws.String(_redshiftdataSchemaPattern)
	}
	if len(_redshiftdataSecretArn) > 0 {
		input.SecretArn = aws.String(_redshiftdataSecretArn)
	}
	if len(_redshiftdataTablePattern) > 0 {
		input.TablePattern = aws.String(_redshiftdataTablePattern)
	}
	if len(_redshiftdataWorkgroupName) > 0 {
		input.WorkgroupName = aws.String(_redshiftdataWorkgroupName)
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

	var results []*redshiftdata.ListTablesOutput
	p := redshiftdata.NewListTablesPaginator(client, input)
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

func init() {
	_rootCmd.AddCommand(_redshiftdataCmd)
	_redshiftdataCmd.Flags().SortFlags = false

	_redshiftdataCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_redshiftdataCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_redshiftdataCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_redshiftdataCmd.Flags().StringVarP(&_redshiftdataClientToken, "client-token", "", "", "Client Token")
	_redshiftdataCmd.Flags().StringVarP(&_redshiftdataClusterIdentifier, "cluster-identifier", "", "", "Cluster Identifier")
	_redshiftdataCmd.Flags().StringVarP(&_redshiftdataConnectedDatabase, "connected-database", "", "", "Connected Database")
	_redshiftdataCmd.Flags().StringVarP(&_redshiftdataDatabase, "database", "", "", "Database")
	_redshiftdataCmd.Flags().StringVarP(&_redshiftdataDbUser, "db-user", "", "", "DB User")
	_redshiftdataCmd.Flags().StringVarP(&_redshiftdataId, "id", "", "", "ID")
	_redshiftdataCmd.Flags().StringVarP(&_redshiftdataMaxResults, "max-results", "", "", "Max Results")
	_redshiftdataCmd.Flags().StringVarP(&_redshiftdataNextToken, "next-token", "", "", "Next Token")
	_redshiftdataCmd.Flags().StringVarP(&_redshiftdataParameters, "parameters", "", "", "Parameters")
	_redshiftdataCmd.Flags().StringVarP(&_redshiftdataResultFormat, "result-format", "", "", "Result Format")
	_redshiftdataCmd.Flags().StringVarP(&_redshiftdataRoleLevel, "role-level", "", "", "Role Level")
	_redshiftdataCmd.Flags().StringVarP(&_redshiftdataSchema, "schema", "", "", "Schema")
	_redshiftdataCmd.Flags().StringVarP(&_redshiftdataSchemaPattern, "schema-pattern", "", "", "Schema Pattern")
	_redshiftdataCmd.Flags().StringVarP(&_redshiftdataSecretArn, "secret-arn", "", "", "Secret ARN")
	_redshiftdataCmd.Flags().StringVarP(&_redshiftdataSessionId, "session-id", "", "", "Session ID")
	_redshiftdataCmd.Flags().StringVarP(&_redshiftdataSessionKeepAliveSeconds, "session-keep-alive-seconds", "", "", "Session Keep Alive Seconds")
	_redshiftdataCmd.Flags().StringVarP(&_redshiftdataSql, "sql", "", "", "Sql")
	_redshiftdataCmd.Flags().StringSliceVarP(&_redshiftdataSqls, "sqls", "", nil, "Sqls")
	_redshiftdataCmd.Flags().StringVarP(&_redshiftdataStatementName, "statement-name", "", "", "Statement Name")
	_redshiftdataCmd.Flags().StringVarP(&_redshiftdataStatus, "status", "", "", "Status")
	_redshiftdataCmd.Flags().StringVarP(&_redshiftdataTable, "table", "", "", "Table")
	_redshiftdataCmd.Flags().StringVarP(&_redshiftdataTablePattern, "table-pattern", "", "", "Table Pattern")
	_redshiftdataCmd.Flags().StringVarP(&_redshiftdataWithEvent, "with-event", "", "", "With Event")
	_redshiftdataCmd.Flags().StringVarP(&_redshiftdataWorkgroupName, "workgroup-name", "", "", "Workgroup Name")

	_redshiftdataCmd.Flags().BoolVarP(&_redshiftdataBatchExecuteStatement, "batch-execute-statement", "", false, "Batch Execute Statement")
	_redshiftdataCmd.Flags().BoolVarP(&_redshiftdataCancelStatement, "cancel-statement", "", false, "Cancel Statement")
	_redshiftdataCmd.Flags().BoolVarP(&_redshiftdataDescribeStatement, "describe-statement", "", false, "Describe Statement")
	_redshiftdataCmd.Flags().BoolVarP(&_redshiftdataDescribeTable, "describe-table", "", false, "Describe Table")
	_redshiftdataCmd.Flags().BoolVarP(&_redshiftdataExecuteStatement, "execute-statement", "", false, "Execute Statement")
	_redshiftdataCmd.Flags().BoolVarP(&_redshiftdataGetStatementResult, "get-statement-result", "", false, "Get Statement Result")
	_redshiftdataCmd.Flags().BoolVarP(&_redshiftdataGetStatementResultV2, "get-statement-result-v2", "", false, "Get Statement Result V2")
	_redshiftdataCmd.Flags().BoolVarP(&_redshiftdataListDatabases, "list-databases", "", false, "List Databases")
	_redshiftdataCmd.Flags().BoolVarP(&_redshiftdataListSchemas, "list-schemas", "", false, "List Schemas")
	_redshiftdataCmd.Flags().BoolVarP(&_redshiftdataListStatements, "list-statements", "", false, "List Statements")
	_redshiftdataCmd.Flags().BoolVarP(&_redshiftdataListTables, "list-tables", "", false, "List Tables")

}
