package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/rdsdata/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"batch-execute-statement", "begin-transaction", "commit-transaction", "execute-sql", "execute-statement", "rollback-transaction"},
		OperationSet: map[string]bool{"batch-execute-statement": true, "begin-transaction": true, "commit-transaction": true, "execute-sql": true, "execute-statement": true, "rollback-transaction": true},
		OperationInputs: map[string][]string{
			"batch-execute-statement": {"Database", "ParameterSets", "ResourceArn", "Schema", "SecretArn", "Sql", "TransactionId"},
			"begin-transaction":       {"Database", "ResourceArn", "Schema", "SecretArn"},
			"commit-transaction":      {"ResourceArn", "SecretArn", "TransactionId"},
			"execute-sql":             {"AwsSecretStoreArn", "Database", "DbClusterOrInstanceArn", "Schema", "SqlStatements"},
			"execute-statement":       {"ContinueAfterTimeout", "Database", "FormatRecordsAs", "IncludeResultMetadata", "Parameters", "ResourceArn", "ResultSetOptions", "Schema", "SecretArn", "Sql", "TransactionId"},
			"rollback-transaction":    {"ResourceArn", "SecretArn", "TransactionId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"batch-execute-statement": {"Database": "*string", "ParameterSets": "[][]types.SqlParameter", "ResourceArn": "*string", "Schema": "*string", "SecretArn": "*string", "Sql": "*string", "TransactionId": "*string"},
			"begin-transaction":       {"Database": "*string", "ResourceArn": "*string", "Schema": "*string", "SecretArn": "*string"},
			"commit-transaction":      {"ResourceArn": "*string", "SecretArn": "*string", "TransactionId": "*string"},
			"execute-sql":             {"AwsSecretStoreArn": "*string", "Database": "*string", "DbClusterOrInstanceArn": "*string", "Schema": "*string", "SqlStatements": "*string"},
			"execute-statement":       {"ContinueAfterTimeout": "bool", "Database": "*string", "FormatRecordsAs": "types.RecordsFormatType", "IncludeResultMetadata": "bool", "Parameters": "[]types.SqlParameter", "ResourceArn": "*string", "ResultSetOptions": "*types.ResultSetOptions", "Schema": "*string", "SecretArn": "*string", "Sql": "*string", "TransactionId": "*string"},
			"rollback-transaction":    {"ResourceArn": "*string", "SecretArn": "*string", "TransactionId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"batch-execute-statement": {"ResourceArn", "SecretArn", "Sql"},
			"begin-transaction":       {"ResourceArn", "SecretArn"},
			"commit-transaction":      {"ResourceArn", "SecretArn", "TransactionId"},
			"execute-sql":             {"AwsSecretStoreArn", "DbClusterOrInstanceArn", "SqlStatements"},
			"execute-statement":       {"ResourceArn", "SecretArn", "Sql"},
			"rollback-transaction":    {"ResourceArn", "SecretArn", "TransactionId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("rdsdata", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
