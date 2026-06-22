package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/glacier"
)

var fields_abort_multipart_upload = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "UploadId", Flag: "upload-id", Type: "*string", Required: true},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_abort_vault_lock = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_add_tags_to_vault = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_complete_multipart_upload = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ArchiveSize", Flag: "archive-size", Type: "*string", Required: false},
	{Name: "Checksum", Flag: "checksum", Type: "*string", Required: false},
	{Name: "UploadId", Flag: "upload-id", Type: "*string", Required: true},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_complete_vault_lock = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "LockId", Flag: "lock-id", Type: "*string", Required: true},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_create_vault = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_delete_archive = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ArchiveId", Flag: "archive-id", Type: "*string", Required: true},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_delete_vault = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_delete_vault_access_policy = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_delete_vault_notifications = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_describe_job = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_describe_vault = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_get_data_retrieval_policy = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_get_job_output = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "Range", Flag: "range", Type: "*string", Required: false},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_get_vault_access_policy = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_get_vault_lock = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_get_vault_notifications = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_initiate_job = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "JobParameters", Flag: "job-parameters", Type: "*types.JobParameters", Required: false},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_initiate_multipart_upload = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ArchiveDescription", Flag: "archive-description", Type: "*string", Required: false},
	{Name: "PartSize", Flag: "part-size", Type: "*string", Required: false},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_initiate_vault_lock = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*types.VaultLockPolicy", Required: false},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_list_jobs = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Completed", Flag: "completed", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "Statuscode", Flag: "statuscode", Type: "*string", Required: false},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_list_multipart_uploads = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_list_parts = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "UploadId", Flag: "upload-id", Type: "*string", Required: true},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_list_provisioned_capacity = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_list_tags_for_vault = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_list_vaults = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_purchase_provisioned_capacity = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_remove_tags_from_vault = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: false},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_set_data_retrieval_policy = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*types.DataRetrievalPolicy", Required: false},
}

var fields_set_vault_access_policy = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*types.VaultAccessPolicy", Required: false},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_set_vault_notifications = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
	{Name: "VaultNotificationConfig", Flag: "vault-notification-config", Type: "*types.VaultNotificationConfig", Required: false},
}

var fields_upload_archive = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ArchiveDescription", Flag: "archive-description", Type: "*string", Required: false},
	{Name: "Body", Flag: "body", Type: "io.Reader", Required: false},
	{Name: "Checksum", Flag: "checksum", Type: "*string", Required: false},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

var fields_upload_multipart_part = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Body", Flag: "body", Type: "io.Reader", Required: false},
	{Name: "Checksum", Flag: "checksum", Type: "*string", Required: false},
	{Name: "Range", Flag: "range", Type: "*string", Required: false},
	{Name: "UploadId", Flag: "upload-id", Type: "*string", Required: true},
	{Name: "VaultName", Flag: "vault-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"abort-multipart-upload": {
			Name:   "abort-multipart-upload",
			Fields: fields_abort_multipart_upload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AbortMultipartUploadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_abort_multipart_upload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AbortMultipartUpload(ctx, input)
			},
		},
		"abort-vault-lock": {
			Name:   "abort-vault-lock",
			Fields: fields_abort_vault_lock,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AbortVaultLockInput{}
				if _, err := leanruntime.ApplyInput(input, fields_abort_vault_lock, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AbortVaultLock(ctx, input)
			},
		},
		"add-tags-to-vault": {
			Name:   "add-tags-to-vault",
			Fields: fields_add_tags_to_vault,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddTagsToVaultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_tags_to_vault, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddTagsToVault(ctx, input)
			},
		},
		"complete-multipart-upload": {
			Name:   "complete-multipart-upload",
			Fields: fields_complete_multipart_upload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CompleteMultipartUploadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_complete_multipart_upload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CompleteMultipartUpload(ctx, input)
			},
		},
		"complete-vault-lock": {
			Name:   "complete-vault-lock",
			Fields: fields_complete_vault_lock,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CompleteVaultLockInput{}
				if _, err := leanruntime.ApplyInput(input, fields_complete_vault_lock, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CompleteVaultLock(ctx, input)
			},
		},
		"create-vault": {
			Name:   "create-vault",
			Fields: fields_create_vault,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVaultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vault, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVault(ctx, input)
			},
		},
		"delete-archive": {
			Name:   "delete-archive",
			Fields: fields_delete_archive,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteArchiveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_archive, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteArchive(ctx, input)
			},
		},
		"delete-vault": {
			Name:   "delete-vault",
			Fields: fields_delete_vault,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVaultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vault, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVault(ctx, input)
			},
		},
		"delete-vault-access-policy": {
			Name:   "delete-vault-access-policy",
			Fields: fields_delete_vault_access_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVaultAccessPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vault_access_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVaultAccessPolicy(ctx, input)
			},
		},
		"delete-vault-notifications": {
			Name:   "delete-vault-notifications",
			Fields: fields_delete_vault_notifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVaultNotificationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vault_notifications, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVaultNotifications(ctx, input)
			},
		},
		"describe-job": {
			Name:   "describe-job",
			Fields: fields_describe_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeJob(ctx, input)
			},
		},
		"describe-vault": {
			Name:   "describe-vault",
			Fields: fields_describe_vault,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVaultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_vault, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVault(ctx, input)
			},
		},
		"get-data-retrieval-policy": {
			Name:   "get-data-retrieval-policy",
			Fields: fields_get_data_retrieval_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataRetrievalPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_retrieval_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataRetrievalPolicy(ctx, input)
			},
		},
		"get-job-output": {
			Name:   "get-job-output",
			Fields: fields_get_job_output,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJobOutputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_job_output, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJobOutput(ctx, input)
			},
		},
		"get-vault-access-policy": {
			Name:   "get-vault-access-policy",
			Fields: fields_get_vault_access_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVaultAccessPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_vault_access_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVaultAccessPolicy(ctx, input)
			},
		},
		"get-vault-lock": {
			Name:   "get-vault-lock",
			Fields: fields_get_vault_lock,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVaultLockInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_vault_lock, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVaultLock(ctx, input)
			},
		},
		"get-vault-notifications": {
			Name:   "get-vault-notifications",
			Fields: fields_get_vault_notifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVaultNotificationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_vault_notifications, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVaultNotifications(ctx, input)
			},
		},
		"initiate-job": {
			Name:   "initiate-job",
			Fields: fields_initiate_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InitiateJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_initiate_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InitiateJob(ctx, input)
			},
		},
		"initiate-multipart-upload": {
			Name:   "initiate-multipart-upload",
			Fields: fields_initiate_multipart_upload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InitiateMultipartUploadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_initiate_multipart_upload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InitiateMultipartUpload(ctx, input)
			},
		},
		"initiate-vault-lock": {
			Name:   "initiate-vault-lock",
			Fields: fields_initiate_vault_lock,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InitiateVaultLockInput{}
				if _, err := leanruntime.ApplyInput(input, fields_initiate_vault_lock, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InitiateVaultLock(ctx, input)
			},
		},
		"list-jobs": {
			Name:   "list-jobs",
			Fields: fields_list_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListJobs(ctx, input)
				}
				var results []*svc.ListJobsOutput
				p := svc.NewListJobsPaginator(client, input)
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
		"list-multipart-uploads": {
			Name:   "list-multipart-uploads",
			Fields: fields_list_multipart_uploads,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMultipartUploadsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_multipart_uploads, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMultipartUploads(ctx, input)
				}
				var results []*svc.ListMultipartUploadsOutput
				p := svc.NewListMultipartUploadsPaginator(client, input)
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
		"list-parts": {
			Name:   "list-parts",
			Fields: fields_list_parts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPartsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_parts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListParts(ctx, input)
				}
				var results []*svc.ListPartsOutput
				p := svc.NewListPartsPaginator(client, input)
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
		"list-provisioned-capacity": {
			Name:   "list-provisioned-capacity",
			Fields: fields_list_provisioned_capacity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProvisionedCapacityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_provisioned_capacity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListProvisionedCapacity(ctx, input)
			},
		},
		"list-tags-for-vault": {
			Name:   "list-tags-for-vault",
			Fields: fields_list_tags_for_vault,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForVaultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_vault, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForVault(ctx, input)
			},
		},
		"list-vaults": {
			Name:   "list-vaults",
			Fields: fields_list_vaults,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVaultsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_vaults, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVaults(ctx, input)
				}
				var results []*svc.ListVaultsOutput
				p := svc.NewListVaultsPaginator(client, input)
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
		"purchase-provisioned-capacity": {
			Name:   "purchase-provisioned-capacity",
			Fields: fields_purchase_provisioned_capacity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PurchaseProvisionedCapacityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_purchase_provisioned_capacity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PurchaseProvisionedCapacity(ctx, input)
			},
		},
		"remove-tags-from-vault": {
			Name:   "remove-tags-from-vault",
			Fields: fields_remove_tags_from_vault,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveTagsFromVaultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_tags_from_vault, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveTagsFromVault(ctx, input)
			},
		},
		"set-data-retrieval-policy": {
			Name:   "set-data-retrieval-policy",
			Fields: fields_set_data_retrieval_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetDataRetrievalPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_data_retrieval_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetDataRetrievalPolicy(ctx, input)
			},
		},
		"set-vault-access-policy": {
			Name:   "set-vault-access-policy",
			Fields: fields_set_vault_access_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetVaultAccessPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_vault_access_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetVaultAccessPolicy(ctx, input)
			},
		},
		"set-vault-notifications": {
			Name:   "set-vault-notifications",
			Fields: fields_set_vault_notifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetVaultNotificationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_vault_notifications, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetVaultNotifications(ctx, input)
			},
		},
		"upload-archive": {
			Name:   "upload-archive",
			Fields: fields_upload_archive,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UploadArchiveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_upload_archive, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UploadArchive(ctx, input)
			},
		},
		"upload-multipart-part": {
			Name:   "upload-multipart-part",
			Fields: fields_upload_multipart_part,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UploadMultipartPartInput{}
				if _, err := leanruntime.ApplyInput(input, fields_upload_multipart_part, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UploadMultipartPart(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("glacier", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
