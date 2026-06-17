package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/redshiftdata/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"batch-execute-statement", "cancel-statement", "describe-statement", "describe-table", "execute-statement", "get-statement-result", "get-statement-result-v2", "list-databases", "list-schemas", "list-statements", "list-tables"},
		OperationSet: map[string]bool{"batch-execute-statement": true, "cancel-statement": true, "describe-statement": true, "describe-table": true, "execute-statement": true, "get-statement-result": true, "get-statement-result-v2": true, "list-databases": true, "list-schemas": true, "list-statements": true, "list-tables": true},
		OperationInputs: map[string][]string{
			"batch-execute-statement": {"ClientToken", "ClusterIdentifier", "Database", "DbUser", "ResultFormat", "SecretArn", "SessionId", "SessionKeepAliveSeconds", "Sqls", "StatementName", "WithEvent", "WorkgroupName"},
			"cancel-statement":        {"Id"},
			"describe-statement":      {"Id"},
			"describe-table":          {"ClusterIdentifier", "ConnectedDatabase", "Database", "DbUser", "MaxResults", "NextToken", "Schema", "SecretArn", "Table", "WorkgroupName"},
			"execute-statement":       {"ClientToken", "ClusterIdentifier", "Database", "DbUser", "Parameters", "ResultFormat", "SecretArn", "SessionId", "SessionKeepAliveSeconds", "Sql", "StatementName", "WithEvent", "WorkgroupName"},
			"get-statement-result":    {"Id", "NextToken"},
			"get-statement-result-v2": {"Id", "NextToken"},
			"list-databases":          {"ClusterIdentifier", "Database", "DbUser", "MaxResults", "NextToken", "SecretArn", "WorkgroupName"},
			"list-schemas":            {"ClusterIdentifier", "ConnectedDatabase", "Database", "DbUser", "MaxResults", "NextToken", "SchemaPattern", "SecretArn", "WorkgroupName"},
			"list-statements":         {"ClusterIdentifier", "Database", "MaxResults", "NextToken", "RoleLevel", "StatementName", "Status", "WorkgroupName"},
			"list-tables":             {"ClusterIdentifier", "ConnectedDatabase", "Database", "DbUser", "MaxResults", "NextToken", "SchemaPattern", "SecretArn", "TablePattern", "WorkgroupName"},
		},
		OperationInputTypes: map[string]map[string]string{
			"batch-execute-statement": {"ClientToken": "*string", "ClusterIdentifier": "*string", "Database": "*string", "DbUser": "*string", "ResultFormat": "types.ResultFormatString", "SecretArn": "*string", "SessionId": "*string", "SessionKeepAliveSeconds": "*int32", "Sqls": "[]string", "StatementName": "*string", "WithEvent": "*bool", "WorkgroupName": "*string"},
			"cancel-statement":        {"Id": "*string"},
			"describe-statement":      {"Id": "*string"},
			"describe-table":          {"ClusterIdentifier": "*string", "ConnectedDatabase": "*string", "Database": "*string", "DbUser": "*string", "MaxResults": "int32", "NextToken": "*string", "Schema": "*string", "SecretArn": "*string", "Table": "*string", "WorkgroupName": "*string"},
			"execute-statement":       {"ClientToken": "*string", "ClusterIdentifier": "*string", "Database": "*string", "DbUser": "*string", "Parameters": "[]types.SqlParameter", "ResultFormat": "types.ResultFormatString", "SecretArn": "*string", "SessionId": "*string", "SessionKeepAliveSeconds": "*int32", "Sql": "*string", "StatementName": "*string", "WithEvent": "*bool", "WorkgroupName": "*string"},
			"get-statement-result":    {"Id": "*string", "NextToken": "*string"},
			"get-statement-result-v2": {"Id": "*string", "NextToken": "*string"},
			"list-databases":          {"ClusterIdentifier": "*string", "Database": "*string", "DbUser": "*string", "MaxResults": "int32", "NextToken": "*string", "SecretArn": "*string", "WorkgroupName": "*string"},
			"list-schemas":            {"ClusterIdentifier": "*string", "ConnectedDatabase": "*string", "Database": "*string", "DbUser": "*string", "MaxResults": "int32", "NextToken": "*string", "SchemaPattern": "*string", "SecretArn": "*string", "WorkgroupName": "*string"},
			"list-statements":         {"ClusterIdentifier": "*string", "Database": "*string", "MaxResults": "int32", "NextToken": "*string", "RoleLevel": "*bool", "StatementName": "*string", "Status": "types.StatusString", "WorkgroupName": "*string"},
			"list-tables":             {"ClusterIdentifier": "*string", "ConnectedDatabase": "*string", "Database": "*string", "DbUser": "*string", "MaxResults": "int32", "NextToken": "*string", "SchemaPattern": "*string", "SecretArn": "*string", "TablePattern": "*string", "WorkgroupName": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"batch-execute-statement": {"Sqls"},
			"cancel-statement":        {"Id"},
			"describe-statement":      {"Id"},
			"describe-table":          {"Database"},
			"execute-statement":       {"Sql"},
			"get-statement-result":    {"Id"},
			"get-statement-result-v2": {"Id"},
			"list-databases":          {"Database"},
			"list-schemas":            {"Database"},
			"list-statements":         {},
			"list-tables":             {"Database"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("redshiftdata", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
