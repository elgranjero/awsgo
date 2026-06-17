package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/timestreamwrite/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-batch-load-task", "create-database", "create-table", "delete-database", "delete-table", "describe-batch-load-task", "describe-database", "describe-endpoints", "describe-table", "list-batch-load-tasks", "list-databases", "list-tables", "list-tags-for-resource", "resume-batch-load-task", "tag-resource", "untag-resource", "update-database", "update-table", "write-records"},
		OperationSet: map[string]bool{"create-batch-load-task": true, "create-database": true, "create-table": true, "delete-database": true, "delete-table": true, "describe-batch-load-task": true, "describe-database": true, "describe-endpoints": true, "describe-table": true, "list-batch-load-tasks": true, "list-databases": true, "list-tables": true, "list-tags-for-resource": true, "resume-batch-load-task": true, "tag-resource": true, "untag-resource": true, "update-database": true, "update-table": true, "write-records": true},
		OperationInputs: map[string][]string{
			"create-batch-load-task":   {"ClientToken", "DataModelConfiguration", "DataSourceConfiguration", "RecordVersion", "ReportConfiguration", "TargetDatabaseName", "TargetTableName"},
			"create-database":          {"DatabaseName", "KmsKeyId", "Tags"},
			"create-table":             {"DatabaseName", "MagneticStoreWriteProperties", "RetentionProperties", "Schema", "TableName", "Tags"},
			"delete-database":          {"DatabaseName"},
			"delete-table":             {"DatabaseName", "TableName"},
			"describe-batch-load-task": {"TaskId"},
			"describe-database":        {"DatabaseName"},
			"describe-endpoints":       {},
			"describe-table":           {"DatabaseName", "TableName"},
			"list-batch-load-tasks":    {"MaxResults", "NextToken", "TaskStatus"},
			"list-databases":           {"MaxResults", "NextToken"},
			"list-tables":              {"DatabaseName", "MaxResults", "NextToken"},
			"list-tags-for-resource":   {"ResourceARN"},
			"resume-batch-load-task":   {"TaskId"},
			"tag-resource":             {"ResourceARN", "Tags"},
			"untag-resource":           {"ResourceARN", "TagKeys"},
			"update-database":          {"DatabaseName", "KmsKeyId"},
			"update-table":             {"DatabaseName", "MagneticStoreWriteProperties", "RetentionProperties", "Schema", "TableName"},
			"write-records":            {"CommonAttributes", "DatabaseName", "Records", "TableName"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-batch-load-task":   {"ClientToken": "*string", "DataModelConfiguration": "*types.DataModelConfiguration", "DataSourceConfiguration": "*types.DataSourceConfiguration", "RecordVersion": "*int64", "ReportConfiguration": "*types.ReportConfiguration", "TargetDatabaseName": "*string", "TargetTableName": "*string"},
			"create-database":          {"DatabaseName": "*string", "KmsKeyId": "*string", "Tags": "[]types.Tag"},
			"create-table":             {"DatabaseName": "*string", "MagneticStoreWriteProperties": "*types.MagneticStoreWriteProperties", "RetentionProperties": "*types.RetentionProperties", "Schema": "*types.Schema", "TableName": "*string", "Tags": "[]types.Tag"},
			"delete-database":          {"DatabaseName": "*string"},
			"delete-table":             {"DatabaseName": "*string", "TableName": "*string"},
			"describe-batch-load-task": {"TaskId": "*string"},
			"describe-database":        {"DatabaseName": "*string"},
			"describe-endpoints":       {},
			"describe-table":           {"DatabaseName": "*string", "TableName": "*string"},
			"list-batch-load-tasks":    {"MaxResults": "*int32", "NextToken": "*string", "TaskStatus": "types.BatchLoadStatus"},
			"list-databases":           {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tables":              {"DatabaseName": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":   {"ResourceARN": "*string"},
			"resume-batch-load-task":   {"TaskId": "*string"},
			"tag-resource":             {"ResourceARN": "*string", "Tags": "[]types.Tag"},
			"untag-resource":           {"ResourceARN": "*string", "TagKeys": "[]string"},
			"update-database":          {"DatabaseName": "*string", "KmsKeyId": "*string"},
			"update-table":             {"DatabaseName": "*string", "MagneticStoreWriteProperties": "*types.MagneticStoreWriteProperties", "RetentionProperties": "*types.RetentionProperties", "Schema": "*types.Schema", "TableName": "*string"},
			"write-records":            {"CommonAttributes": "*types.Record", "DatabaseName": "*string", "Records": "[]types.Record", "TableName": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-batch-load-task":   {"DataSourceConfiguration", "ReportConfiguration", "TargetDatabaseName", "TargetTableName"},
			"create-database":          {"DatabaseName"},
			"create-table":             {"DatabaseName", "TableName"},
			"delete-database":          {"DatabaseName"},
			"delete-table":             {"DatabaseName", "TableName"},
			"describe-batch-load-task": {"TaskId"},
			"describe-database":        {"DatabaseName"},
			"describe-endpoints":       {},
			"describe-table":           {"DatabaseName", "TableName"},
			"list-batch-load-tasks":    {},
			"list-databases":           {},
			"list-tables":              {},
			"list-tags-for-resource":   {"ResourceARN"},
			"resume-batch-load-task":   {"TaskId"},
			"tag-resource":             {"ResourceARN", "Tags"},
			"untag-resource":           {"ResourceARN", "TagKeys"},
			"update-database":          {"DatabaseName", "KmsKeyId"},
			"update-table":             {"DatabaseName", "TableName"},
			"write-records":            {"DatabaseName", "Records", "TableName"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("timestreamwrite", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
