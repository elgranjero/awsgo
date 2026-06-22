package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/rdsdata"
)

var fields_batch_execute_statement = []leanruntime.Field{
	{Name: "Database", Flag: "database", Type: "*string", Required: false},
	{Name: "ParameterSets", Flag: "parameter-sets", Type: "[][]types.SqlParameter", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Schema", Flag: "schema", Type: "*string", Required: false},
	{Name: "SecretArn", Flag: "secret-arn", Type: "*string", Required: true},
	{Name: "Sql", Flag: "sql", Type: "*string", Required: true},
	{Name: "TransactionId", Flag: "transaction-id", Type: "*string", Required: false},
}

var fields_begin_transaction = []leanruntime.Field{
	{Name: "Database", Flag: "database", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Schema", Flag: "schema", Type: "*string", Required: false},
	{Name: "SecretArn", Flag: "secret-arn", Type: "*string", Required: true},
}

var fields_commit_transaction = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "SecretArn", Flag: "secret-arn", Type: "*string", Required: true},
	{Name: "TransactionId", Flag: "transaction-id", Type: "*string", Required: true},
}

var fields_execute_sql = []leanruntime.Field{
	{Name: "AwsSecretStoreArn", Flag: "aws-secret-store-arn", Type: "*string", Required: true},
	{Name: "Database", Flag: "database", Type: "*string", Required: false},
	{Name: "DbClusterOrInstanceArn", Flag: "db-cluster-or-instance-arn", Type: "*string", Required: true},
	{Name: "Schema", Flag: "schema", Type: "*string", Required: false},
	{Name: "SqlStatements", Flag: "sql-statements", Type: "*string", Required: true},
}

var fields_execute_statement = []leanruntime.Field{
	{Name: "ContinueAfterTimeout", Flag: "continue-after-timeout", Type: "bool", Required: false},
	{Name: "Database", Flag: "database", Type: "*string", Required: false},
	{Name: "FormatRecordsAs", Flag: "format-records-as", Type: "types.RecordsFormatType", Required: false},
	{Name: "IncludeResultMetadata", Flag: "include-result-metadata", Type: "bool", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "[]types.SqlParameter", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "ResultSetOptions", Flag: "result-set-options", Type: "*types.ResultSetOptions", Required: false},
	{Name: "Schema", Flag: "schema", Type: "*string", Required: false},
	{Name: "SecretArn", Flag: "secret-arn", Type: "*string", Required: true},
	{Name: "Sql", Flag: "sql", Type: "*string", Required: true},
	{Name: "TransactionId", Flag: "transaction-id", Type: "*string", Required: false},
}

var fields_rollback_transaction = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "SecretArn", Flag: "secret-arn", Type: "*string", Required: true},
	{Name: "TransactionId", Flag: "transaction-id", Type: "*string", Required: true},
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
		"begin-transaction": {
			Name:   "begin-transaction",
			Fields: fields_begin_transaction,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BeginTransactionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_begin_transaction, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BeginTransaction(ctx, input)
			},
		},
		"commit-transaction": {
			Name:   "commit-transaction",
			Fields: fields_commit_transaction,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CommitTransactionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_commit_transaction, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CommitTransaction(ctx, input)
			},
		},
		"execute-sql": {
			Name:   "execute-sql",
			Fields: fields_execute_sql,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExecuteSqlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_execute_sql, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExecuteSql(ctx, input)
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
		"rollback-transaction": {
			Name:   "rollback-transaction",
			Fields: fields_rollback_transaction,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RollbackTransactionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_rollback_transaction, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RollbackTransaction(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("rdsdata", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
