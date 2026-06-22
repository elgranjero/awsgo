package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/entityresolution"
)

var fields_add_policy_statement = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "[]string", Required: true},
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Condition", Flag: "condition", Type: "*string", Required: false},
	{Name: "Effect", Flag: "effect", Type: "types.StatementEffect", Required: true},
	{Name: "Principal", Flag: "principal", Type: "[]string", Required: true},
	{Name: "StatementId", Flag: "statement-id", Type: "*string", Required: true},
}

var fields_batch_delete_unique_id = []leanruntime.Field{
	{Name: "InputSource", Flag: "input-source", Type: "*string", Required: false},
	{Name: "UniqueIds", Flag: "unique-ids", Type: "[]string", Required: true},
	{Name: "WorkflowName", Flag: "workflow-name", Type: "*string", Required: true},
}

var fields_create_id_mapping_workflow = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IdMappingTechniques", Flag: "id-mapping-techniques", Type: "*types.IdMappingTechniques", Required: true},
	{Name: "IncrementalRunConfig", Flag: "incremental-run-config", Type: "*types.IdMappingIncrementalRunConfig", Required: false},
	{Name: "InputSourceConfig", Flag: "input-source-config", Type: "[]types.IdMappingWorkflowInputSource", Required: true},
	{Name: "OutputSourceConfig", Flag: "output-source-config", Type: "[]types.IdMappingWorkflowOutputSource", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "WorkflowName", Flag: "workflow-name", Type: "*string", Required: true},
}

var fields_create_id_namespace = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IdMappingWorkflowProperties", Flag: "id-mapping-workflow-properties", Type: "[]types.IdNamespaceIdMappingWorkflowProperties", Required: false},
	{Name: "IdNamespaceName", Flag: "id-namespace-name", Type: "*string", Required: true},
	{Name: "InputSourceConfig", Flag: "input-source-config", Type: "[]types.IdNamespaceInputSource", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.IdNamespaceType", Required: true},
}

var fields_create_matching_workflow = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IncrementalRunConfig", Flag: "incremental-run-config", Type: "*types.IncrementalRunConfig", Required: false},
	{Name: "InputSourceConfig", Flag: "input-source-config", Type: "[]types.InputSource", Required: true},
	{Name: "OutputSourceConfig", Flag: "output-source-config", Type: "[]types.OutputSource", Required: true},
	{Name: "ResolutionTechniques", Flag: "resolution-techniques", Type: "*types.ResolutionTechniques", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "WorkflowName", Flag: "workflow-name", Type: "*string", Required: true},
}

var fields_create_schema_mapping = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MappedInputFields", Flag: "mapped-input-fields", Type: "[]types.SchemaInputAttribute", Required: true},
	{Name: "SchemaName", Flag: "schema-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_id_mapping_workflow = []leanruntime.Field{
	{Name: "WorkflowName", Flag: "workflow-name", Type: "*string", Required: true},
}

var fields_delete_id_namespace = []leanruntime.Field{
	{Name: "IdNamespaceName", Flag: "id-namespace-name", Type: "*string", Required: true},
}

var fields_delete_matching_workflow = []leanruntime.Field{
	{Name: "WorkflowName", Flag: "workflow-name", Type: "*string", Required: true},
}

var fields_delete_policy_statement = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "StatementId", Flag: "statement-id", Type: "*string", Required: true},
}

var fields_delete_schema_mapping = []leanruntime.Field{
	{Name: "SchemaName", Flag: "schema-name", Type: "*string", Required: true},
}

var fields_generate_match_id = []leanruntime.Field{
	{Name: "ProcessingType", Flag: "processing-type", Type: "types.ProcessingType", Required: false},
	{Name: "Records", Flag: "records", Type: "[]types.Record", Required: true},
	{Name: "WorkflowName", Flag: "workflow-name", Type: "*string", Required: true},
}

var fields_get_id_mapping_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "WorkflowName", Flag: "workflow-name", Type: "*string", Required: true},
}

var fields_get_id_mapping_workflow = []leanruntime.Field{
	{Name: "WorkflowName", Flag: "workflow-name", Type: "*string", Required: true},
}

var fields_get_id_namespace = []leanruntime.Field{
	{Name: "IdNamespaceName", Flag: "id-namespace-name", Type: "*string", Required: true},
}

var fields_get_match_id = []leanruntime.Field{
	{Name: "ApplyNormalization", Flag: "apply-normalization", Type: "*bool", Required: false},
	{Name: "Record", Flag: "record", Type: "map[string]string", Required: true},
	{Name: "WorkflowName", Flag: "workflow-name", Type: "*string", Required: true},
}

var fields_get_matching_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "WorkflowName", Flag: "workflow-name", Type: "*string", Required: true},
}

var fields_get_matching_workflow = []leanruntime.Field{
	{Name: "WorkflowName", Flag: "workflow-name", Type: "*string", Required: true},
}

var fields_get_policy = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_provider_service = []leanruntime.Field{
	{Name: "ProviderName", Flag: "provider-name", Type: "*string", Required: true},
	{Name: "ProviderServiceName", Flag: "provider-service-name", Type: "*string", Required: true},
}

var fields_get_schema_mapping = []leanruntime.Field{
	{Name: "SchemaName", Flag: "schema-name", Type: "*string", Required: true},
}

var fields_list_id_mapping_jobs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkflowName", Flag: "workflow-name", Type: "*string", Required: true},
}

var fields_list_id_mapping_workflows = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_id_namespaces = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_matching_jobs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkflowName", Flag: "workflow-name", Type: "*string", Required: true},
}

var fields_list_matching_workflows = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_provider_services = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProviderName", Flag: "provider-name", Type: "*string", Required: false},
}

var fields_list_schema_mappings = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_policy = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "Token", Flag: "token", Type: "*string", Required: false},
}

var fields_start_id_mapping_job = []leanruntime.Field{
	{Name: "JobType", Flag: "job-type", Type: "types.JobType", Required: false},
	{Name: "OutputSourceConfig", Flag: "output-source-config", Type: "[]types.IdMappingJobOutputSource", Required: false},
	{Name: "WorkflowName", Flag: "workflow-name", Type: "*string", Required: true},
}

var fields_start_matching_job = []leanruntime.Field{
	{Name: "WorkflowName", Flag: "workflow-name", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_id_mapping_workflow = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IdMappingTechniques", Flag: "id-mapping-techniques", Type: "*types.IdMappingTechniques", Required: true},
	{Name: "IncrementalRunConfig", Flag: "incremental-run-config", Type: "*types.IdMappingIncrementalRunConfig", Required: false},
	{Name: "InputSourceConfig", Flag: "input-source-config", Type: "[]types.IdMappingWorkflowInputSource", Required: true},
	{Name: "OutputSourceConfig", Flag: "output-source-config", Type: "[]types.IdMappingWorkflowOutputSource", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "WorkflowName", Flag: "workflow-name", Type: "*string", Required: true},
}

var fields_update_id_namespace = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IdMappingWorkflowProperties", Flag: "id-mapping-workflow-properties", Type: "[]types.IdNamespaceIdMappingWorkflowProperties", Required: false},
	{Name: "IdNamespaceName", Flag: "id-namespace-name", Type: "*string", Required: true},
	{Name: "InputSourceConfig", Flag: "input-source-config", Type: "[]types.IdNamespaceInputSource", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_update_matching_workflow = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IncrementalRunConfig", Flag: "incremental-run-config", Type: "*types.IncrementalRunConfig", Required: false},
	{Name: "InputSourceConfig", Flag: "input-source-config", Type: "[]types.InputSource", Required: true},
	{Name: "OutputSourceConfig", Flag: "output-source-config", Type: "[]types.OutputSource", Required: true},
	{Name: "ResolutionTechniques", Flag: "resolution-techniques", Type: "*types.ResolutionTechniques", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "WorkflowName", Flag: "workflow-name", Type: "*string", Required: true},
}

var fields_update_schema_mapping = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MappedInputFields", Flag: "mapped-input-fields", Type: "[]types.SchemaInputAttribute", Required: true},
	{Name: "SchemaName", Flag: "schema-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-policy-statement": {
			Name:   "add-policy-statement",
			Fields: fields_add_policy_statement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddPolicyStatementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_policy_statement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddPolicyStatement(ctx, input)
			},
		},
		"batch-delete-unique-id": {
			Name:   "batch-delete-unique-id",
			Fields: fields_batch_delete_unique_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteUniqueIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_unique_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteUniqueId(ctx, input)
			},
		},
		"create-id-mapping-workflow": {
			Name:   "create-id-mapping-workflow",
			Fields: fields_create_id_mapping_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIdMappingWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_id_mapping_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIdMappingWorkflow(ctx, input)
			},
		},
		"create-id-namespace": {
			Name:   "create-id-namespace",
			Fields: fields_create_id_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIdNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_id_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIdNamespace(ctx, input)
			},
		},
		"create-matching-workflow": {
			Name:   "create-matching-workflow",
			Fields: fields_create_matching_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMatchingWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_matching_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMatchingWorkflow(ctx, input)
			},
		},
		"create-schema-mapping": {
			Name:   "create-schema-mapping",
			Fields: fields_create_schema_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSchemaMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_schema_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSchemaMapping(ctx, input)
			},
		},
		"delete-id-mapping-workflow": {
			Name:   "delete-id-mapping-workflow",
			Fields: fields_delete_id_mapping_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIdMappingWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_id_mapping_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIdMappingWorkflow(ctx, input)
			},
		},
		"delete-id-namespace": {
			Name:   "delete-id-namespace",
			Fields: fields_delete_id_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIdNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_id_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIdNamespace(ctx, input)
			},
		},
		"delete-matching-workflow": {
			Name:   "delete-matching-workflow",
			Fields: fields_delete_matching_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMatchingWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_matching_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMatchingWorkflow(ctx, input)
			},
		},
		"delete-policy-statement": {
			Name:   "delete-policy-statement",
			Fields: fields_delete_policy_statement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePolicyStatementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_policy_statement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePolicyStatement(ctx, input)
			},
		},
		"delete-schema-mapping": {
			Name:   "delete-schema-mapping",
			Fields: fields_delete_schema_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSchemaMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_schema_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSchemaMapping(ctx, input)
			},
		},
		"generate-match-id": {
			Name:   "generate-match-id",
			Fields: fields_generate_match_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateMatchIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_match_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateMatchId(ctx, input)
			},
		},
		"get-id-mapping-job": {
			Name:   "get-id-mapping-job",
			Fields: fields_get_id_mapping_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIdMappingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_id_mapping_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIdMappingJob(ctx, input)
			},
		},
		"get-id-mapping-workflow": {
			Name:   "get-id-mapping-workflow",
			Fields: fields_get_id_mapping_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIdMappingWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_id_mapping_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIdMappingWorkflow(ctx, input)
			},
		},
		"get-id-namespace": {
			Name:   "get-id-namespace",
			Fields: fields_get_id_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIdNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_id_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIdNamespace(ctx, input)
			},
		},
		"get-match-id": {
			Name:   "get-match-id",
			Fields: fields_get_match_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMatchIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_match_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMatchId(ctx, input)
			},
		},
		"get-matching-job": {
			Name:   "get-matching-job",
			Fields: fields_get_matching_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMatchingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_matching_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMatchingJob(ctx, input)
			},
		},
		"get-matching-workflow": {
			Name:   "get-matching-workflow",
			Fields: fields_get_matching_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMatchingWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_matching_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMatchingWorkflow(ctx, input)
			},
		},
		"get-policy": {
			Name:   "get-policy",
			Fields: fields_get_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPolicy(ctx, input)
			},
		},
		"get-provider-service": {
			Name:   "get-provider-service",
			Fields: fields_get_provider_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProviderServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_provider_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProviderService(ctx, input)
			},
		},
		"get-schema-mapping": {
			Name:   "get-schema-mapping",
			Fields: fields_get_schema_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSchemaMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_schema_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSchemaMapping(ctx, input)
			},
		},
		"list-id-mapping-jobs": {
			Name:   "list-id-mapping-jobs",
			Fields: fields_list_id_mapping_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIdMappingJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_id_mapping_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIdMappingJobs(ctx, input)
				}
				var results []*svc.ListIdMappingJobsOutput
				p := svc.NewListIdMappingJobsPaginator(client, input)
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
		"list-id-mapping-workflows": {
			Name:   "list-id-mapping-workflows",
			Fields: fields_list_id_mapping_workflows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIdMappingWorkflowsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_id_mapping_workflows, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIdMappingWorkflows(ctx, input)
				}
				var results []*svc.ListIdMappingWorkflowsOutput
				p := svc.NewListIdMappingWorkflowsPaginator(client, input)
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
		"list-id-namespaces": {
			Name:   "list-id-namespaces",
			Fields: fields_list_id_namespaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIdNamespacesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_id_namespaces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIdNamespaces(ctx, input)
				}
				var results []*svc.ListIdNamespacesOutput
				p := svc.NewListIdNamespacesPaginator(client, input)
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
		"list-matching-jobs": {
			Name:   "list-matching-jobs",
			Fields: fields_list_matching_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMatchingJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_matching_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMatchingJobs(ctx, input)
				}
				var results []*svc.ListMatchingJobsOutput
				p := svc.NewListMatchingJobsPaginator(client, input)
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
		"list-matching-workflows": {
			Name:   "list-matching-workflows",
			Fields: fields_list_matching_workflows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMatchingWorkflowsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_matching_workflows, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMatchingWorkflows(ctx, input)
				}
				var results []*svc.ListMatchingWorkflowsOutput
				p := svc.NewListMatchingWorkflowsPaginator(client, input)
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
		"list-provider-services": {
			Name:   "list-provider-services",
			Fields: fields_list_provider_services,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProviderServicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_provider_services, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProviderServices(ctx, input)
				}
				var results []*svc.ListProviderServicesOutput
				p := svc.NewListProviderServicesPaginator(client, input)
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
		"list-schema-mappings": {
			Name:   "list-schema-mappings",
			Fields: fields_list_schema_mappings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSchemaMappingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_schema_mappings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSchemaMappings(ctx, input)
				}
				var results []*svc.ListSchemaMappingsOutput
				p := svc.NewListSchemaMappingsPaginator(client, input)
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
		"put-policy": {
			Name:   "put-policy",
			Fields: fields_put_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutPolicy(ctx, input)
			},
		},
		"start-id-mapping-job": {
			Name:   "start-id-mapping-job",
			Fields: fields_start_id_mapping_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartIdMappingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_id_mapping_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartIdMappingJob(ctx, input)
			},
		},
		"start-matching-job": {
			Name:   "start-matching-job",
			Fields: fields_start_matching_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMatchingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_matching_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMatchingJob(ctx, input)
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
		"update-id-mapping-workflow": {
			Name:   "update-id-mapping-workflow",
			Fields: fields_update_id_mapping_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIdMappingWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_id_mapping_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIdMappingWorkflow(ctx, input)
			},
		},
		"update-id-namespace": {
			Name:   "update-id-namespace",
			Fields: fields_update_id_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIdNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_id_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIdNamespace(ctx, input)
			},
		},
		"update-matching-workflow": {
			Name:   "update-matching-workflow",
			Fields: fields_update_matching_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMatchingWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_matching_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMatchingWorkflow(ctx, input)
			},
		},
		"update-schema-mapping": {
			Name:   "update-schema-mapping",
			Fields: fields_update_schema_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSchemaMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_schema_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSchemaMapping(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("entityresolution", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
