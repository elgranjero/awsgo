package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/redshiftdata"
)

var fields_batch_execute_statement = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: false},
	{Name: "Database", Flag: "database", Type: "*string", Required: false},
	{Name: "DbUser", Flag: "db-user", Type: "*string", Required: false},
	{Name: "ResultFormat", Flag: "result-format", Type: "types.ResultFormatString", Required: false},
	{Name: "SecretArn", Flag: "secret-arn", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: false},
	{Name: "SessionKeepAliveSeconds", Flag: "session-keep-alive-seconds", Type: "*int32", Required: false},
	{Name: "Sqls", Flag: "sqls", Type: "[]string", Required: true},
	{Name: "StatementName", Flag: "statement-name", Type: "*string", Required: false},
	{Name: "WithEvent", Flag: "with-event", Type: "*bool", Required: false},
	{Name: "WorkgroupName", Flag: "workgroup-name", Type: "*string", Required: false},
}

var fields_cancel_statement = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_statement = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_table = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: false},
	{Name: "ConnectedDatabase", Flag: "connected-database", Type: "*string", Required: false},
	{Name: "Database", Flag: "database", Type: "*string", Required: true},
	{Name: "DbUser", Flag: "db-user", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Schema", Flag: "schema", Type: "*string", Required: false},
	{Name: "SecretArn", Flag: "secret-arn", Type: "*string", Required: false},
	{Name: "Table", Flag: "table", Type: "*string", Required: false},
	{Name: "WorkgroupName", Flag: "workgroup-name", Type: "*string", Required: false},
}

var fields_execute_statement = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: false},
	{Name: "Database", Flag: "database", Type: "*string", Required: false},
	{Name: "DbUser", Flag: "db-user", Type: "*string", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "[]types.SqlParameter", Required: false},
	{Name: "ResultFormat", Flag: "result-format", Type: "types.ResultFormatString", Required: false},
	{Name: "SecretArn", Flag: "secret-arn", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: false},
	{Name: "SessionKeepAliveSeconds", Flag: "session-keep-alive-seconds", Type: "*int32", Required: false},
	{Name: "Sql", Flag: "sql", Type: "*string", Required: true},
	{Name: "StatementName", Flag: "statement-name", Type: "*string", Required: false},
	{Name: "WithEvent", Flag: "with-event", Type: "*bool", Required: false},
	{Name: "WorkgroupName", Flag: "workgroup-name", Type: "*string", Required: false},
}

var fields_get_statement_result = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_statement_result_v2 = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_databases = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: false},
	{Name: "Database", Flag: "database", Type: "*string", Required: true},
	{Name: "DbUser", Flag: "db-user", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SecretArn", Flag: "secret-arn", Type: "*string", Required: false},
	{Name: "WorkgroupName", Flag: "workgroup-name", Type: "*string", Required: false},
}

var fields_list_schemas = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: false},
	{Name: "ConnectedDatabase", Flag: "connected-database", Type: "*string", Required: false},
	{Name: "Database", Flag: "database", Type: "*string", Required: true},
	{Name: "DbUser", Flag: "db-user", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SchemaPattern", Flag: "schema-pattern", Type: "*string", Required: false},
	{Name: "SecretArn", Flag: "secret-arn", Type: "*string", Required: false},
	{Name: "WorkgroupName", Flag: "workgroup-name", Type: "*string", Required: false},
}

var fields_list_statements = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: false},
	{Name: "Database", Flag: "database", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RoleLevel", Flag: "role-level", Type: "*bool", Required: false},
	{Name: "StatementName", Flag: "statement-name", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.StatusString", Required: false},
	{Name: "WorkgroupName", Flag: "workgroup-name", Type: "*string", Required: false},
}

var fields_list_tables = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: false},
	{Name: "ConnectedDatabase", Flag: "connected-database", Type: "*string", Required: false},
	{Name: "Database", Flag: "database", Type: "*string", Required: true},
	{Name: "DbUser", Flag: "db-user", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SchemaPattern", Flag: "schema-pattern", Type: "*string", Required: false},
	{Name: "SecretArn", Flag: "secret-arn", Type: "*string", Required: false},
	{Name: "TablePattern", Flag: "table-pattern", Type: "*string", Required: false},
	{Name: "WorkgroupName", Flag: "workgroup-name", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-execute-statement": {
			Name:   "batch-execute-statement",
			Fields: fields_batch_execute_statement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchExecuteStatementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_execute_statement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchExecuteStatement(ctx, input)
			},
		},
		"cancel-statement": {
			Name:   "cancel-statement",
			Fields: fields_cancel_statement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelStatementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_statement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelStatement(ctx, input)
			},
		},
		"describe-statement": {
			Name:   "describe-statement",
			Fields: fields_describe_statement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStatementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_statement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStatement(ctx, input)
			},
		},
		"describe-table": {
			Name:   "describe-table",
			Fields: fields_describe_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTableInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_table, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTable(ctx, input)
				}
				var results []*svc.DescribeTableOutput
				p := svc.NewDescribeTablePaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"execute-statement": {
			Name:   "execute-statement",
			Fields: fields_execute_statement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExecuteStatementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_execute_statement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExecuteStatement(ctx, input)
			},
		},
		"get-statement-result": {
			Name:   "get-statement-result",
			Fields: fields_get_statement_result,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStatementResultInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_statement_result, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetStatementResult(ctx, input)
				}
				var results []*svc.GetStatementResultOutput
				p := svc.NewGetStatementResultPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-statement-result-v2": {
			Name:   "get-statement-result-v2",
			Fields: fields_get_statement_result_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStatementResultV2Input{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_statement_result_v2, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetStatementResultV2(ctx, input)
				}
				var results []*svc.GetStatementResultV2Output
				p := svc.NewGetStatementResultV2Paginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-databases": {
			Name:   "list-databases",
			Fields: fields_list_databases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDatabasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_databases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDatabases(ctx, input)
				}
				var results []*svc.ListDatabasesOutput
				p := svc.NewListDatabasesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-schemas": {
			Name:   "list-schemas",
			Fields: fields_list_schemas,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSchemasInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_schemas, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSchemas(ctx, input)
				}
				var results []*svc.ListSchemasOutput
				p := svc.NewListSchemasPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-statements": {
			Name:   "list-statements",
			Fields: fields_list_statements,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStatementsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_statements, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStatements(ctx, input)
				}
				var results []*svc.ListStatementsOutput
				p := svc.NewListStatementsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-tables": {
			Name:   "list-tables",
			Fields: fields_list_tables,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTablesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tables, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTables(ctx, input)
				}
				var results []*svc.ListTablesOutput
				p := svc.NewListTablesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
	}
	if err := leanruntime.Execute("redshiftdata", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
