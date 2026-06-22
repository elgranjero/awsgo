package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/amplifybackend"
)

var fields_clone_backend = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BackendEnvironmentName", Flag: "backend-environment-name", Type: "*string", Required: true},
	{Name: "TargetEnvironmentName", Flag: "target-environment-name", Type: "*string", Required: true},
}

var fields_create_backend = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "AppName", Flag: "app-name", Type: "*string", Required: true},
	{Name: "BackendEnvironmentName", Flag: "backend-environment-name", Type: "*string", Required: true},
	{Name: "ResourceConfig", Flag: "resource-config", Type: "*types.ResourceConfig", Required: false},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: false},
}

var fields_create_backend_api = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BackendEnvironmentName", Flag: "backend-environment-name", Type: "*string", Required: true},
	{Name: "ResourceConfig", Flag: "resource-config", Type: "*types.BackendAPIResourceConfig", Required: true},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_create_backend_auth = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BackendEnvironmentName", Flag: "backend-environment-name", Type: "*string", Required: true},
	{Name: "ResourceConfig", Flag: "resource-config", Type: "*types.CreateBackendAuthResourceConfig", Required: true},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_create_backend_config = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BackendManagerAppId", Flag: "backend-manager-app-id", Type: "*string", Required: false},
}

var fields_create_backend_storage = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BackendEnvironmentName", Flag: "backend-environment-name", Type: "*string", Required: true},
	{Name: "ResourceConfig", Flag: "resource-config", Type: "*types.CreateBackendStorageResourceConfig", Required: true},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_create_token = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
}

var fields_delete_backend = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BackendEnvironmentName", Flag: "backend-environment-name", Type: "*string", Required: true},
}

var fields_delete_backend_api = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BackendEnvironmentName", Flag: "backend-environment-name", Type: "*string", Required: true},
	{Name: "ResourceConfig", Flag: "resource-config", Type: "*types.BackendAPIResourceConfig", Required: false},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_delete_backend_auth = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BackendEnvironmentName", Flag: "backend-environment-name", Type: "*string", Required: true},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_delete_backend_storage = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BackendEnvironmentName", Flag: "backend-environment-name", Type: "*string", Required: true},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
	{Name: "ServiceName", Flag: "service-name", Type: "types.ServiceName", Required: true},
}

var fields_delete_token = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_generate_backend_api_models = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BackendEnvironmentName", Flag: "backend-environment-name", Type: "*string", Required: true},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_get_backend = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BackendEnvironmentName", Flag: "backend-environment-name", Type: "*string", Required: false},
}

var fields_get_backend_api = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BackendEnvironmentName", Flag: "backend-environment-name", Type: "*string", Required: true},
	{Name: "ResourceConfig", Flag: "resource-config", Type: "*types.BackendAPIResourceConfig", Required: false},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_get_backend_api_models = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BackendEnvironmentName", Flag: "backend-environment-name", Type: "*string", Required: true},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_get_backend_auth = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BackendEnvironmentName", Flag: "backend-environment-name", Type: "*string", Required: true},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_get_backend_job = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BackendEnvironmentName", Flag: "backend-environment-name", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_backend_storage = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BackendEnvironmentName", Flag: "backend-environment-name", Type: "*string", Required: true},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_get_token = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_import_backend_auth = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BackendEnvironmentName", Flag: "backend-environment-name", Type: "*string", Required: true},
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: false},
	{Name: "NativeClientId", Flag: "native-client-id", Type: "*string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "WebClientId", Flag: "web-client-id", Type: "*string", Required: true},
}

var fields_import_backend_storage = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BackendEnvironmentName", Flag: "backend-environment-name", Type: "*string", Required: true},
	{Name: "BucketName", Flag: "bucket-name", Type: "*string", Required: false},
	{Name: "ServiceName", Flag: "service-name", Type: "types.ServiceName", Required: true},
}

var fields_list_backend_jobs = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BackendEnvironmentName", Flag: "backend-environment-name", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Operation", Flag: "operation", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "*string", Required: false},
}

var fields_list_s3_buckets = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_remove_all_backends = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "CleanAmplifyApp", Flag: "clean-amplify-app", Type: "*bool", Required: false},
}

var fields_remove_backend_config = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
}

var fields_update_backend_api = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BackendEnvironmentName", Flag: "backend-environment-name", Type: "*string", Required: true},
	{Name: "ResourceConfig", Flag: "resource-config", Type: "*types.BackendAPIResourceConfig", Required: false},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_update_backend_auth = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BackendEnvironmentName", Flag: "backend-environment-name", Type: "*string", Required: true},
	{Name: "ResourceConfig", Flag: "resource-config", Type: "*types.UpdateBackendAuthResourceConfig", Required: true},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_update_backend_config = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "LoginAuthConfig", Flag: "login-auth-config", Type: "*types.LoginAuthConfigReqObj", Required: false},
}

var fields_update_backend_job = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BackendEnvironmentName", Flag: "backend-environment-name", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "Operation", Flag: "operation", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "*string", Required: false},
}

var fields_update_backend_storage = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BackendEnvironmentName", Flag: "backend-environment-name", Type: "*string", Required: true},
	{Name: "ResourceConfig", Flag: "resource-config", Type: "*types.UpdateBackendStorageResourceConfig", Required: true},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"clone-backend": {
			Name:   "clone-backend",
			Fields: fields_clone_backend,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CloneBackendInput{}
				if _, err := leanruntime.ApplyInput(input, fields_clone_backend, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CloneBackend(ctx, input)
			},
		},
		"create-backend": {
			Name:   "create-backend",
			Fields: fields_create_backend,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBackendInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_backend, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBackend(ctx, input)
			},
		},
		"create-backend-api": {
			Name:   "create-backend-api",
			Fields: fields_create_backend_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBackendAPIInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_backend_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBackendAPI(ctx, input)
			},
		},
		"create-backend-auth": {
			Name:   "create-backend-auth",
			Fields: fields_create_backend_auth,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBackendAuthInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_backend_auth, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBackendAuth(ctx, input)
			},
		},
		"create-backend-config": {
			Name:   "create-backend-config",
			Fields: fields_create_backend_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBackendConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_backend_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBackendConfig(ctx, input)
			},
		},
		"create-backend-storage": {
			Name:   "create-backend-storage",
			Fields: fields_create_backend_storage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBackendStorageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_backend_storage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBackendStorage(ctx, input)
			},
		},
		"create-token": {
			Name:   "create-token",
			Fields: fields_create_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateToken(ctx, input)
			},
		},
		"delete-backend": {
			Name:   "delete-backend",
			Fields: fields_delete_backend,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBackendInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_backend, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBackend(ctx, input)
			},
		},
		"delete-backend-api": {
			Name:   "delete-backend-api",
			Fields: fields_delete_backend_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBackendAPIInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_backend_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBackendAPI(ctx, input)
			},
		},
		"delete-backend-auth": {
			Name:   "delete-backend-auth",
			Fields: fields_delete_backend_auth,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBackendAuthInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_backend_auth, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBackendAuth(ctx, input)
			},
		},
		"delete-backend-storage": {
			Name:   "delete-backend-storage",
			Fields: fields_delete_backend_storage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBackendStorageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_backend_storage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBackendStorage(ctx, input)
			},
		},
		"delete-token": {
			Name:   "delete-token",
			Fields: fields_delete_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteToken(ctx, input)
			},
		},
		"generate-backend-api-models": {
			Name:   "generate-backend-api-models",
			Fields: fields_generate_backend_api_models,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateBackendAPIModelsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_backend_api_models, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateBackendAPIModels(ctx, input)
			},
		},
		"get-backend": {
			Name:   "get-backend",
			Fields: fields_get_backend,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBackendInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_backend, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBackend(ctx, input)
			},
		},
		"get-backend-api": {
			Name:   "get-backend-api",
			Fields: fields_get_backend_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBackendAPIInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_backend_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBackendAPI(ctx, input)
			},
		},
		"get-backend-api-models": {
			Name:   "get-backend-api-models",
			Fields: fields_get_backend_api_models,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBackendAPIModelsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_backend_api_models, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBackendAPIModels(ctx, input)
			},
		},
		"get-backend-auth": {
			Name:   "get-backend-auth",
			Fields: fields_get_backend_auth,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBackendAuthInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_backend_auth, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBackendAuth(ctx, input)
			},
		},
		"get-backend-job": {
			Name:   "get-backend-job",
			Fields: fields_get_backend_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBackendJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_backend_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBackendJob(ctx, input)
			},
		},
		"get-backend-storage": {
			Name:   "get-backend-storage",
			Fields: fields_get_backend_storage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBackendStorageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_backend_storage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBackendStorage(ctx, input)
			},
		},
		"get-token": {
			Name:   "get-token",
			Fields: fields_get_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetToken(ctx, input)
			},
		},
		"import-backend-auth": {
			Name:   "import-backend-auth",
			Fields: fields_import_backend_auth,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportBackendAuthInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_backend_auth, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportBackendAuth(ctx, input)
			},
		},
		"import-backend-storage": {
			Name:   "import-backend-storage",
			Fields: fields_import_backend_storage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportBackendStorageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_backend_storage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportBackendStorage(ctx, input)
			},
		},
		"list-backend-jobs": {
			Name:   "list-backend-jobs",
			Fields: fields_list_backend_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBackendJobsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_backend_jobs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListBackendJobs(ctx, input)
			},
		},
		"list-s3-buckets": {
			Name:   "list-s3-buckets",
			Fields: fields_list_s3_buckets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListS3BucketsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_s3_buckets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListS3Buckets(ctx, input)
			},
		},
		"remove-all-backends": {
			Name:   "remove-all-backends",
			Fields: fields_remove_all_backends,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveAllBackendsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_all_backends, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveAllBackends(ctx, input)
			},
		},
		"remove-backend-config": {
			Name:   "remove-backend-config",
			Fields: fields_remove_backend_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveBackendConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_backend_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveBackendConfig(ctx, input)
			},
		},
		"update-backend-api": {
			Name:   "update-backend-api",
			Fields: fields_update_backend_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBackendAPIInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_backend_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBackendAPI(ctx, input)
			},
		},
		"update-backend-auth": {
			Name:   "update-backend-auth",
			Fields: fields_update_backend_auth,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBackendAuthInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_backend_auth, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBackendAuth(ctx, input)
			},
		},
		"update-backend-config": {
			Name:   "update-backend-config",
			Fields: fields_update_backend_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBackendConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_backend_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBackendConfig(ctx, input)
			},
		},
		"update-backend-job": {
			Name:   "update-backend-job",
			Fields: fields_update_backend_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBackendJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_backend_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBackendJob(ctx, input)
			},
		},
		"update-backend-storage": {
			Name:   "update-backend-storage",
			Fields: fields_update_backend_storage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBackendStorageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_backend_storage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBackendStorage(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("amplifybackend", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
