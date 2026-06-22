package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
)

var fields_create_batch_load_task = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DataModelConfiguration", Flag: "data-model-configuration", Type: "*types.DataModelConfiguration", Required: false},
	{Name: "DataSourceConfiguration", Flag: "data-source-configuration", Type: "*types.DataSourceConfiguration", Required: true},
	{Name: "RecordVersion", Flag: "record-version", Type: "*int64", Required: false},
	{Name: "ReportConfiguration", Flag: "report-configuration", Type: "*types.ReportConfiguration", Required: true},
	{Name: "TargetDatabaseName", Flag: "target-database-name", Type: "*string", Required: true},
	{Name: "TargetTableName", Flag: "target-table-name", Type: "*string", Required: true},
}

var fields_create_database = []leanruntime.Field{
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_table = []leanruntime.Field{
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "MagneticStoreWriteProperties", Flag: "magnetic-store-write-properties", Type: "*types.MagneticStoreWriteProperties", Required: false},
	{Name: "RetentionProperties", Flag: "retention-properties", Type: "*types.RetentionProperties", Required: false},
	{Name: "Schema", Flag: "schema", Type: "*types.Schema", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_database = []leanruntime.Field{
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
}

var fields_delete_table = []leanruntime.Field{
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_describe_batch_load_task = []leanruntime.Field{
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_describe_database = []leanruntime.Field{
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
}

var fields_describe_endpoints = []leanruntime.Field{}

var fields_describe_table = []leanruntime.Field{
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_list_batch_load_tasks = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TaskStatus", Flag: "task-status", Type: "types.BatchLoadStatus", Required: false},
}

var fields_list_databases = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tables = []leanruntime.Field{
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_resume_batch_load_task = []leanruntime.Field{
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_database = []leanruntime.Field{
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: true},
}

var fields_update_table = []leanruntime.Field{
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "MagneticStoreWriteProperties", Flag: "magnetic-store-write-properties", Type: "*types.MagneticStoreWriteProperties", Required: false},
	{Name: "RetentionProperties", Flag: "retention-properties", Type: "*types.RetentionProperties", Required: false},
	{Name: "Schema", Flag: "schema", Type: "*types.Schema", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_write_records = []leanruntime.Field{
	{Name: "CommonAttributes", Flag: "common-attributes", Type: "*types.Record", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "Records", Flag: "records", Type: "[]types.Record", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-batch-load-task": {
			Name:   "create-batch-load-task",
			Fields: fields_create_batch_load_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBatchLoadTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_batch_load_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBatchLoadTask(ctx, input)
			},
		},
		"create-database": {
			Name:   "create-database",
			Fields: fields_create_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDatabase(ctx, input)
			},
		},
		"create-table": {
			Name:   "create-table",
			Fields: fields_create_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTable(ctx, input)
			},
		},
		"delete-database": {
			Name:   "delete-database",
			Fields: fields_delete_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDatabase(ctx, input)
			},
		},
		"delete-table": {
			Name:   "delete-table",
			Fields: fields_delete_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTable(ctx, input)
			},
		},
		"describe-batch-load-task": {
			Name:   "describe-batch-load-task",
			Fields: fields_describe_batch_load_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBatchLoadTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_batch_load_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBatchLoadTask(ctx, input)
			},
		},
		"describe-database": {
			Name:   "describe-database",
			Fields: fields_describe_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDatabase(ctx, input)
			},
		},
		"describe-endpoints": {
			Name:   "describe-endpoints",
			Fields: fields_describe_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEndpointsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_endpoints, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEndpoints(ctx, input)
			},
		},
		"describe-table": {
			Name:   "describe-table",
			Fields: fields_describe_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTable(ctx, input)
			},
		},
		"list-batch-load-tasks": {
			Name:   "list-batch-load-tasks",
			Fields: fields_list_batch_load_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBatchLoadTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_batch_load_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBatchLoadTasks(ctx, input)
				}
				var results []*svc.ListBatchLoadTasksOutput
				p := svc.NewListBatchLoadTasksPaginator(client, input)
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
		"list-tags-for-resource": {
			Name:   "list-tags-for-resource",
			Fields: fields_list_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForResource(ctx, input)
			},
		},
		"resume-batch-load-task": {
			Name:   "resume-batch-load-task",
			Fields: fields_resume_batch_load_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResumeBatchLoadTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_resume_batch_load_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResumeBatchLoadTask(ctx, input)
			},
		},
		"tag-resource": {
			Name:   "tag-resource",
			Fields: fields_tag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagResource(ctx, input)
			},
		},
		"untag-resource": {
			Name:   "untag-resource",
			Fields: fields_untag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagResource(ctx, input)
			},
		},
		"update-database": {
			Name:   "update-database",
			Fields: fields_update_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDatabase(ctx, input)
			},
		},
		"update-table": {
			Name:   "update-table",
			Fields: fields_update_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTable(ctx, input)
			},
		},
		"write-records": {
			Name:   "write-records",
			Fields: fields_write_records,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.WriteRecordsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_write_records, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.WriteRecords(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("timestreamwrite", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
