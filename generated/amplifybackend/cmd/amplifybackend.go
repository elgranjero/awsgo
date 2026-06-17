package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amplifybackend"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// amplifybackendCmd represents the amplifybackend command
var _amplifybackendCmd = &cobra.Command{
	Use:   "amplifybackend",
	Short: "AWS amplifybackend CLI",
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
		client := amplifybackend.NewFromConfig(cfg)
		if _amplifybackendCloneBackend {
			amplifybackend_CloneBackend(cfg, client)
			return
		}
		if _amplifybackendCreateBackend {
			amplifybackend_CreateBackend(cfg, client)
			return
		}
		if _amplifybackendCreateBackendAPI {
			amplifybackend_CreateBackendAPI(cfg, client)
			return
		}
		if _amplifybackendCreateBackendAuth {
			amplifybackend_CreateBackendAuth(cfg, client)
			return
		}
		if _amplifybackendCreateBackendConfig {
			amplifybackend_CreateBackendConfig(cfg, client)
			return
		}
		if _amplifybackendCreateBackendStorage {
			amplifybackend_CreateBackendStorage(cfg, client)
			return
		}
		if _amplifybackendCreateToken {
			amplifybackend_CreateToken(cfg, client)
			return
		}
		if _amplifybackendDeleteBackend {
			amplifybackend_DeleteBackend(cfg, client)
			return
		}
		if _amplifybackendDeleteBackendAPI {
			amplifybackend_DeleteBackendAPI(cfg, client)
			return
		}
		if _amplifybackendDeleteBackendAuth {
			amplifybackend_DeleteBackendAuth(cfg, client)
			return
		}
		if _amplifybackendDeleteBackendStorage {
			amplifybackend_DeleteBackendStorage(cfg, client)
			return
		}
		if _amplifybackendDeleteToken {
			amplifybackend_DeleteToken(cfg, client)
			return
		}
		if _amplifybackendGenerateBackendAPIModels {
			amplifybackend_GenerateBackendAPIModels(cfg, client)
			return
		}
		if _amplifybackendGetBackend {
			amplifybackend_GetBackend(cfg, client)
			return
		}
		if _amplifybackendGetBackendAPI {
			amplifybackend_GetBackendAPI(cfg, client)
			return
		}
		if _amplifybackendGetBackendAPIModels {
			amplifybackend_GetBackendAPIModels(cfg, client)
			return
		}
		if _amplifybackendGetBackendAuth {
			amplifybackend_GetBackendAuth(cfg, client)
			return
		}
		if _amplifybackendGetBackendJob {
			amplifybackend_GetBackendJob(cfg, client)
			return
		}
		if _amplifybackendGetBackendStorage {
			amplifybackend_GetBackendStorage(cfg, client)
			return
		}
		if _amplifybackendGetToken {
			amplifybackend_GetToken(cfg, client)
			return
		}
		if _amplifybackendImportBackendAuth {
			amplifybackend_ImportBackendAuth(cfg, client)
			return
		}
		if _amplifybackendImportBackendStorage {
			amplifybackend_ImportBackendStorage(cfg, client)
			return
		}
		if _amplifybackendListBackendJobs {
			amplifybackend_ListBackendJobs(cfg, client)
			return
		}
		if _amplifybackendListS3Buckets {
			amplifybackend_ListS3Buckets(cfg, client)
			return
		}
		if _amplifybackendRemoveAllBackends {
			amplifybackend_RemoveAllBackends(cfg, client)
			return
		}
		if _amplifybackendRemoveBackendConfig {
			amplifybackend_RemoveBackendConfig(cfg, client)
			return
		}
		if _amplifybackendUpdateBackendAPI {
			amplifybackend_UpdateBackendAPI(cfg, client)
			return
		}
		if _amplifybackendUpdateBackendAuth {
			amplifybackend_UpdateBackendAuth(cfg, client)
			return
		}
		if _amplifybackendUpdateBackendConfig {
			amplifybackend_UpdateBackendConfig(cfg, client)
			return
		}
		if _amplifybackendUpdateBackendJob {
			amplifybackend_UpdateBackendJob(cfg, client)
			return
		}
		if _amplifybackendUpdateBackendStorage {
			amplifybackend_UpdateBackendStorage(cfg, client)
			return
		}

	},
}

var (
	_amplifybackendCloneBackend             bool
	_amplifybackendCreateBackend            bool
	_amplifybackendCreateBackendAPI         bool
	_amplifybackendCreateBackendAuth        bool
	_amplifybackendCreateBackendConfig      bool
	_amplifybackendCreateBackendStorage     bool
	_amplifybackendCreateToken              bool
	_amplifybackendDeleteBackend            bool
	_amplifybackendDeleteBackendAPI         bool
	_amplifybackendDeleteBackendAuth        bool
	_amplifybackendDeleteBackendStorage     bool
	_amplifybackendDeleteToken              bool
	_amplifybackendGenerateBackendAPIModels bool
	_amplifybackendGetBackend               bool
	_amplifybackendGetBackendAPI            bool
	_amplifybackendGetBackendAPIModels      bool
	_amplifybackendGetBackendAuth           bool
	_amplifybackendGetBackendJob            bool
	_amplifybackendGetBackendStorage        bool
	_amplifybackendGetToken                 bool
	_amplifybackendImportBackendAuth        bool
	_amplifybackendImportBackendStorage     bool
	_amplifybackendListBackendJobs          bool
	_amplifybackendListS3Buckets            bool
	_amplifybackendRemoveAllBackends        bool
	_amplifybackendRemoveBackendConfig      bool
	_amplifybackendUpdateBackendAPI         bool
	_amplifybackendUpdateBackendAuth        bool
	_amplifybackendUpdateBackendConfig      bool
	_amplifybackendUpdateBackendJob         bool
	_amplifybackendUpdateBackendStorage     bool

	_amplifybackendAppId                  string
	_amplifybackendAppName                string
	_amplifybackendBackendEnvironmentName string
	_amplifybackendBackendManagerAppId    string
	_amplifybackendBucketName             string
	_amplifybackendCleanAmplifyApp        string
	_amplifybackendIdentityPoolId         string
	_amplifybackendJobId                  string
	_amplifybackendLoginAuthConfig        string
	_amplifybackendMaxResults             string
	_amplifybackendNativeClientId         string
	_amplifybackendNextToken              string
	_amplifybackendOperation              string
	_amplifybackendResourceConfig         string
	_amplifybackendResourceName           string
	_amplifybackendServiceName            string
	_amplifybackendSessionId              string
	_amplifybackendStatus                 string
	_amplifybackendTargetEnvironmentName  string
	_amplifybackendUserPoolId             string
	_amplifybackendWebClientId            string
)

// This operation clones an existing backend.
func amplifybackend_CloneBackend(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.CloneBackendInput{
		// AppId: *string, // Required
		// BackendEnvironmentName: *string, // Required
		// TargetEnvironmentName: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendBackendEnvironmentName) > 0 {
		input.BackendEnvironmentName = aws.String(_amplifybackendBackendEnvironmentName)
	}
	if len(_amplifybackendTargetEnvironmentName) > 0 {
		input.TargetEnvironmentName = aws.String(_amplifybackendTargetEnvironmentName)
	}

	if resp, err := client.CloneBackend(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation creates a backend for an Amplify app. Backends are automatically
// created at the time of app creation.
func amplifybackend_CreateBackend(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.CreateBackendInput{
		// AppId: *string, // Required
		// AppName: *string, // Required
		// BackendEnvironmentName: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendAppName) > 0 {
		input.AppName = aws.String(_amplifybackendAppName)
	}
	if len(_amplifybackendBackendEnvironmentName) > 0 {
		input.BackendEnvironmentName = aws.String(_amplifybackendBackendEnvironmentName)
	}
	if len(_amplifybackendResourceConfig) > 0 {
		if err := assignInputField(input, "ResourceConfig", _amplifybackendResourceConfig); err != nil {
			log.Errorf("invalid --resource-config: %s", err.Error())
			return
		}
	}
	if len(_amplifybackendResourceName) > 0 {
		input.ResourceName = aws.String(_amplifybackendResourceName)
	}

	if resp, err := client.CreateBackend(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new backend API resource.
func amplifybackend_CreateBackendAPI(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.CreateBackendAPIInput{
		// AppId: *string, // Required
		// BackendEnvironmentName: *string, // Required
		// ResourceConfig: *types.BackendAPIResourceConfig, // Required
		// ResourceName: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendBackendEnvironmentName) > 0 {
		input.BackendEnvironmentName = aws.String(_amplifybackendBackendEnvironmentName)
	}
	if len(_amplifybackendResourceConfig) > 0 {
		if err := assignInputField(input, "ResourceConfig", _amplifybackendResourceConfig); err != nil {
			log.Errorf("invalid --resource-config: %s", err.Error())
			return
		}
	}
	if len(_amplifybackendResourceName) > 0 {
		input.ResourceName = aws.String(_amplifybackendResourceName)
	}

	if resp, err := client.CreateBackendAPI(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new backend authentication resource.
func amplifybackend_CreateBackendAuth(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.CreateBackendAuthInput{
		// AppId: *string, // Required
		// BackendEnvironmentName: *string, // Required
		// ResourceConfig: *types.CreateBackendAuthResourceConfig, // Required
		// ResourceName: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendBackendEnvironmentName) > 0 {
		input.BackendEnvironmentName = aws.String(_amplifybackendBackendEnvironmentName)
	}
	if len(_amplifybackendResourceConfig) > 0 {
		if err := assignInputField(input, "ResourceConfig", _amplifybackendResourceConfig); err != nil {
			log.Errorf("invalid --resource-config: %s", err.Error())
			return
		}
	}
	if len(_amplifybackendResourceName) > 0 {
		input.ResourceName = aws.String(_amplifybackendResourceName)
	}

	if resp, err := client.CreateBackendAuth(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a config object for a backend.
func amplifybackend_CreateBackendConfig(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.CreateBackendConfigInput{
		// AppId: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendBackendManagerAppId) > 0 {
		input.BackendManagerAppId = aws.String(_amplifybackendBackendManagerAppId)
	}

	if resp, err := client.CreateBackendConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a backend storage resource.
func amplifybackend_CreateBackendStorage(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.CreateBackendStorageInput{
		// AppId: *string, // Required
		// BackendEnvironmentName: *string, // Required
		// ResourceConfig: *types.CreateBackendStorageResourceConfig, // Required
		// ResourceName: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendBackendEnvironmentName) > 0 {
		input.BackendEnvironmentName = aws.String(_amplifybackendBackendEnvironmentName)
	}
	if len(_amplifybackendResourceConfig) > 0 {
		if err := assignInputField(input, "ResourceConfig", _amplifybackendResourceConfig); err != nil {
			log.Errorf("invalid --resource-config: %s", err.Error())
			return
		}
	}
	if len(_amplifybackendResourceName) > 0 {
		input.ResourceName = aws.String(_amplifybackendResourceName)
	}

	if resp, err := client.CreateBackendStorage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates a one-time challenge code to authenticate a user into your Amplify
// Admin UI.
func amplifybackend_CreateToken(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.CreateTokenInput{
		// AppId: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}

	if resp, err := client.CreateToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes an existing environment from your Amplify project.
func amplifybackend_DeleteBackend(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.DeleteBackendInput{
		// AppId: *string, // Required
		// BackendEnvironmentName: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendBackendEnvironmentName) > 0 {
		input.BackendEnvironmentName = aws.String(_amplifybackendBackendEnvironmentName)
	}

	if resp, err := client.DeleteBackend(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing backend API resource.
func amplifybackend_DeleteBackendAPI(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.DeleteBackendAPIInput{
		// AppId: *string, // Required
		// BackendEnvironmentName: *string, // Required
		// ResourceName: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendBackendEnvironmentName) > 0 {
		input.BackendEnvironmentName = aws.String(_amplifybackendBackendEnvironmentName)
	}
	if len(_amplifybackendResourceName) > 0 {
		input.ResourceName = aws.String(_amplifybackendResourceName)
	}
	if len(_amplifybackendResourceConfig) > 0 {
		if err := assignInputField(input, "ResourceConfig", _amplifybackendResourceConfig); err != nil {
			log.Errorf("invalid --resource-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteBackendAPI(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing backend authentication resource.
func amplifybackend_DeleteBackendAuth(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.DeleteBackendAuthInput{
		// AppId: *string, // Required
		// BackendEnvironmentName: *string, // Required
		// ResourceName: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendBackendEnvironmentName) > 0 {
		input.BackendEnvironmentName = aws.String(_amplifybackendBackendEnvironmentName)
	}
	if len(_amplifybackendResourceName) > 0 {
		input.ResourceName = aws.String(_amplifybackendResourceName)
	}

	if resp, err := client.DeleteBackendAuth(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified backend storage resource.
func amplifybackend_DeleteBackendStorage(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.DeleteBackendStorageInput{
		// AppId: *string, // Required
		// BackendEnvironmentName: *string, // Required
		// ResourceName: *string, // Required
		// ServiceName: types.ServiceName, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendBackendEnvironmentName) > 0 {
		input.BackendEnvironmentName = aws.String(_amplifybackendBackendEnvironmentName)
	}
	if len(_amplifybackendResourceName) > 0 {
		input.ResourceName = aws.String(_amplifybackendResourceName)
	}
	if len(_amplifybackendServiceName) > 0 {
		if err := assignInputField(input, "ServiceName", _amplifybackendServiceName); err != nil {
			log.Errorf("invalid --service-name: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteBackendStorage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the challenge token based on the given appId and sessionId.
func amplifybackend_DeleteToken(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.DeleteTokenInput{
		// AppId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendSessionId) > 0 {
		input.SessionId = aws.String(_amplifybackendSessionId)
	}

	if resp, err := client.DeleteToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates a model schema for an existing backend API resource.
func amplifybackend_GenerateBackendAPIModels(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.GenerateBackendAPIModelsInput{
		// AppId: *string, // Required
		// BackendEnvironmentName: *string, // Required
		// ResourceName: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendBackendEnvironmentName) > 0 {
		input.BackendEnvironmentName = aws.String(_amplifybackendBackendEnvironmentName)
	}
	if len(_amplifybackendResourceName) > 0 {
		input.ResourceName = aws.String(_amplifybackendResourceName)
	}

	if resp, err := client.GenerateBackendAPIModels(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides project-level details for your Amplify UI project.
func amplifybackend_GetBackend(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.GetBackendInput{
		// AppId: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendBackendEnvironmentName) > 0 {
		input.BackendEnvironmentName = aws.String(_amplifybackendBackendEnvironmentName)
	}

	if resp, err := client.GetBackend(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the details for a backend API.
func amplifybackend_GetBackendAPI(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.GetBackendAPIInput{
		// AppId: *string, // Required
		// BackendEnvironmentName: *string, // Required
		// ResourceName: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendBackendEnvironmentName) > 0 {
		input.BackendEnvironmentName = aws.String(_amplifybackendBackendEnvironmentName)
	}
	if len(_amplifybackendResourceName) > 0 {
		input.ResourceName = aws.String(_amplifybackendResourceName)
	}
	if len(_amplifybackendResourceConfig) > 0 {
		if err := assignInputField(input, "ResourceConfig", _amplifybackendResourceConfig); err != nil {
			log.Errorf("invalid --resource-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetBackendAPI(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a model introspection schema for an existing backend API resource.
func amplifybackend_GetBackendAPIModels(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.GetBackendAPIModelsInput{
		// AppId: *string, // Required
		// BackendEnvironmentName: *string, // Required
		// ResourceName: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendBackendEnvironmentName) > 0 {
		input.BackendEnvironmentName = aws.String(_amplifybackendBackendEnvironmentName)
	}
	if len(_amplifybackendResourceName) > 0 {
		input.ResourceName = aws.String(_amplifybackendResourceName)
	}

	if resp, err := client.GetBackendAPIModels(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a backend auth details.
func amplifybackend_GetBackendAuth(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.GetBackendAuthInput{
		// AppId: *string, // Required
		// BackendEnvironmentName: *string, // Required
		// ResourceName: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendBackendEnvironmentName) > 0 {
		input.BackendEnvironmentName = aws.String(_amplifybackendBackendEnvironmentName)
	}
	if len(_amplifybackendResourceName) > 0 {
		input.ResourceName = aws.String(_amplifybackendResourceName)
	}

	if resp, err := client.GetBackendAuth(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific job.
func amplifybackend_GetBackendJob(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.GetBackendJobInput{
		// AppId: *string, // Required
		// BackendEnvironmentName: *string, // Required
		// JobId: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendBackendEnvironmentName) > 0 {
		input.BackendEnvironmentName = aws.String(_amplifybackendBackendEnvironmentName)
	}
	if len(_amplifybackendJobId) > 0 {
		input.JobId = aws.String(_amplifybackendJobId)
	}

	if resp, err := client.GetBackendJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details for a backend storage resource.
func amplifybackend_GetBackendStorage(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.GetBackendStorageInput{
		// AppId: *string, // Required
		// BackendEnvironmentName: *string, // Required
		// ResourceName: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendBackendEnvironmentName) > 0 {
		input.BackendEnvironmentName = aws.String(_amplifybackendBackendEnvironmentName)
	}
	if len(_amplifybackendResourceName) > 0 {
		input.ResourceName = aws.String(_amplifybackendResourceName)
	}

	if resp, err := client.GetBackendStorage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the challenge token based on the given appId and sessionId.
func amplifybackend_GetToken(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.GetTokenInput{
		// AppId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendSessionId) > 0 {
		input.SessionId = aws.String(_amplifybackendSessionId)
	}

	if resp, err := client.GetToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports an existing backend authentication resource.
func amplifybackend_ImportBackendAuth(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.ImportBackendAuthInput{
		// AppId: *string, // Required
		// BackendEnvironmentName: *string, // Required
		// NativeClientId: *string, // Required
		// UserPoolId: *string, // Required
		// WebClientId: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendBackendEnvironmentName) > 0 {
		input.BackendEnvironmentName = aws.String(_amplifybackendBackendEnvironmentName)
	}
	if len(_amplifybackendNativeClientId) > 0 {
		input.NativeClientId = aws.String(_amplifybackendNativeClientId)
	}
	if len(_amplifybackendUserPoolId) > 0 {
		input.UserPoolId = aws.String(_amplifybackendUserPoolId)
	}
	if len(_amplifybackendWebClientId) > 0 {
		input.WebClientId = aws.String(_amplifybackendWebClientId)
	}
	if len(_amplifybackendIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_amplifybackendIdentityPoolId)
	}

	if resp, err := client.ImportBackendAuth(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports an existing backend storage resource.
func amplifybackend_ImportBackendStorage(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.ImportBackendStorageInput{
		// AppId: *string, // Required
		// BackendEnvironmentName: *string, // Required
		// ServiceName: types.ServiceName, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendBackendEnvironmentName) > 0 {
		input.BackendEnvironmentName = aws.String(_amplifybackendBackendEnvironmentName)
	}
	if len(_amplifybackendServiceName) > 0 {
		if err := assignInputField(input, "ServiceName", _amplifybackendServiceName); err != nil {
			log.Errorf("invalid --service-name: %s", err.Error())
			return
		}
	}
	if len(_amplifybackendBucketName) > 0 {
		input.BucketName = aws.String(_amplifybackendBucketName)
	}

	if resp, err := client.ImportBackendStorage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the jobs for the backend of an Amplify app.
func amplifybackend_ListBackendJobs(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.ListBackendJobsInput{
		// AppId: *string, // Required
		// BackendEnvironmentName: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendBackendEnvironmentName) > 0 {
		input.BackendEnvironmentName = aws.String(_amplifybackendBackendEnvironmentName)
	}
	if len(_amplifybackendJobId) > 0 {
		input.JobId = aws.String(_amplifybackendJobId)
	}
	if len(_amplifybackendMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _amplifybackendMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_amplifybackendNextToken) > 0 {
		input.NextToken = aws.String(_amplifybackendNextToken)
	}
	if len(_amplifybackendOperation) > 0 {
		input.Operation = aws.String(_amplifybackendOperation)
	}
	if len(_amplifybackendStatus) > 0 {
		input.Status = aws.String(_amplifybackendStatus)
	}

	if resp, err := client.ListBackendJobs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The list of S3 buckets in your account.
func amplifybackend_ListS3Buckets(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.ListS3BucketsInput{}

	if len(_amplifybackendNextToken) > 0 {
		input.NextToken = aws.String(_amplifybackendNextToken)
	}

	if resp, err := client.ListS3Buckets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes all backend environments from your Amplify project.
func amplifybackend_RemoveAllBackends(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.RemoveAllBackendsInput{
		// AppId: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendCleanAmplifyApp) > 0 {
		if err := assignInputField(input, "CleanAmplifyApp", _amplifybackendCleanAmplifyApp); err != nil {
			log.Errorf("invalid --clean-amplify-app: %s", err.Error())
			return
		}
	}

	if resp, err := client.RemoveAllBackends(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the AWS resources required to access the Amplify Admin UI.
func amplifybackend_RemoveBackendConfig(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.RemoveBackendConfigInput{
		// AppId: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}

	if resp, err := client.RemoveBackendConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing backend API resource.
func amplifybackend_UpdateBackendAPI(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.UpdateBackendAPIInput{
		// AppId: *string, // Required
		// BackendEnvironmentName: *string, // Required
		// ResourceName: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendBackendEnvironmentName) > 0 {
		input.BackendEnvironmentName = aws.String(_amplifybackendBackendEnvironmentName)
	}
	if len(_amplifybackendResourceName) > 0 {
		input.ResourceName = aws.String(_amplifybackendResourceName)
	}
	if len(_amplifybackendResourceConfig) > 0 {
		if err := assignInputField(input, "ResourceConfig", _amplifybackendResourceConfig); err != nil {
			log.Errorf("invalid --resource-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBackendAPI(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing backend authentication resource.
func amplifybackend_UpdateBackendAuth(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.UpdateBackendAuthInput{
		// AppId: *string, // Required
		// BackendEnvironmentName: *string, // Required
		// ResourceConfig: *types.UpdateBackendAuthResourceConfig, // Required
		// ResourceName: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendBackendEnvironmentName) > 0 {
		input.BackendEnvironmentName = aws.String(_amplifybackendBackendEnvironmentName)
	}
	if len(_amplifybackendResourceConfig) > 0 {
		if err := assignInputField(input, "ResourceConfig", _amplifybackendResourceConfig); err != nil {
			log.Errorf("invalid --resource-config: %s", err.Error())
			return
		}
	}
	if len(_amplifybackendResourceName) > 0 {
		input.ResourceName = aws.String(_amplifybackendResourceName)
	}

	if resp, err := client.UpdateBackendAuth(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the AWS resources required to access the Amplify Admin UI.
func amplifybackend_UpdateBackendConfig(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.UpdateBackendConfigInput{
		// AppId: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendLoginAuthConfig) > 0 {
		if err := assignInputField(input, "LoginAuthConfig", _amplifybackendLoginAuthConfig); err != nil {
			log.Errorf("invalid --login-auth-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBackendConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a specific job.
func amplifybackend_UpdateBackendJob(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.UpdateBackendJobInput{
		// AppId: *string, // Required
		// BackendEnvironmentName: *string, // Required
		// JobId: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendBackendEnvironmentName) > 0 {
		input.BackendEnvironmentName = aws.String(_amplifybackendBackendEnvironmentName)
	}
	if len(_amplifybackendJobId) > 0 {
		input.JobId = aws.String(_amplifybackendJobId)
	}
	if len(_amplifybackendOperation) > 0 {
		input.Operation = aws.String(_amplifybackendOperation)
	}
	if len(_amplifybackendStatus) > 0 {
		input.Status = aws.String(_amplifybackendStatus)
	}

	if resp, err := client.UpdateBackendJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing backend storage resource.
func amplifybackend_UpdateBackendStorage(cfg aws.Config, client *amplifybackend.Client) {
	input := &amplifybackend.UpdateBackendStorageInput{
		// AppId: *string, // Required
		// BackendEnvironmentName: *string, // Required
		// ResourceConfig: *types.UpdateBackendStorageResourceConfig, // Required
		// ResourceName: *string, // Required
	}

	if len(_amplifybackendAppId) > 0 {
		input.AppId = aws.String(_amplifybackendAppId)
	}
	if len(_amplifybackendBackendEnvironmentName) > 0 {
		input.BackendEnvironmentName = aws.String(_amplifybackendBackendEnvironmentName)
	}
	if len(_amplifybackendResourceConfig) > 0 {
		if err := assignInputField(input, "ResourceConfig", _amplifybackendResourceConfig); err != nil {
			log.Errorf("invalid --resource-config: %s", err.Error())
			return
		}
	}
	if len(_amplifybackendResourceName) > 0 {
		input.ResourceName = aws.String(_amplifybackendResourceName)
	}

	if resp, err := client.UpdateBackendStorage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_amplifybackendCmd)
	_amplifybackendCmd.Flags().SortFlags = false

	_amplifybackendCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_amplifybackendCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_amplifybackendCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_amplifybackendCmd.Flags().StringVarP(&_amplifybackendAppId, "app-id", "", "", "App ID")
	_amplifybackendCmd.Flags().StringVarP(&_amplifybackendAppName, "app-name", "", "", "App Name")
	_amplifybackendCmd.Flags().StringVarP(&_amplifybackendBackendEnvironmentName, "backend-environment-name", "", "", "Backend Environment Name")
	_amplifybackendCmd.Flags().StringVarP(&_amplifybackendBackendManagerAppId, "backend-manager-app-id", "", "", "Backend Manager App ID")
	_amplifybackendCmd.Flags().StringVarP(&_amplifybackendBucketName, "bucket-name", "", "", "Bucket Name")
	_amplifybackendCmd.Flags().StringVarP(&_amplifybackendCleanAmplifyApp, "clean-amplify-app", "", "", "Clean Amplify App")
	_amplifybackendCmd.Flags().StringVarP(&_amplifybackendIdentityPoolId, "identity-pool-id", "", "", "Identity Pool ID")
	_amplifybackendCmd.Flags().StringVarP(&_amplifybackendJobId, "job-id", "", "", "Job ID")
	_amplifybackendCmd.Flags().StringVarP(&_amplifybackendLoginAuthConfig, "login-auth-config", "", "", "Login Auth Config")
	_amplifybackendCmd.Flags().StringVarP(&_amplifybackendMaxResults, "max-results", "", "", "Max Results")
	_amplifybackendCmd.Flags().StringVarP(&_amplifybackendNativeClientId, "native-client-id", "", "", "Native Client ID")
	_amplifybackendCmd.Flags().StringVarP(&_amplifybackendNextToken, "next-token", "", "", "Next Token")
	_amplifybackendCmd.Flags().StringVarP(&_amplifybackendOperation, "operation", "", "", "Operation")
	_amplifybackendCmd.Flags().StringVarP(&_amplifybackendResourceConfig, "resource-config", "", "", "Resource Config")
	_amplifybackendCmd.Flags().StringVarP(&_amplifybackendResourceName, "resource-name", "", "", "Resource Name")
	_amplifybackendCmd.Flags().StringVarP(&_amplifybackendServiceName, "service-name", "", "", "Service Name")
	_amplifybackendCmd.Flags().StringVarP(&_amplifybackendSessionId, "session-id", "", "", "Session ID")
	_amplifybackendCmd.Flags().StringVarP(&_amplifybackendStatus, "status", "", "", "Status")
	_amplifybackendCmd.Flags().StringVarP(&_amplifybackendTargetEnvironmentName, "target-environment-name", "", "", "Target Environment Name")
	_amplifybackendCmd.Flags().StringVarP(&_amplifybackendUserPoolId, "user-pool-id", "", "", "User Pool ID")
	_amplifybackendCmd.Flags().StringVarP(&_amplifybackendWebClientId, "web-client-id", "", "", "Web Client ID")

	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendCloneBackend, "clone-backend", "", false, "Clone Backend")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendCreateBackend, "create-backend", "", false, "Create Backend")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendCreateBackendAPI, "create-backend-api", "", false, "Create Backend API")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendCreateBackendAuth, "create-backend-auth", "", false, "Create Backend Auth")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendCreateBackendConfig, "create-backend-config", "", false, "Create Backend Config")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendCreateBackendStorage, "create-backend-storage", "", false, "Create Backend Storage")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendCreateToken, "create-token", "", false, "Create Token")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendDeleteBackend, "delete-backend", "", false, "Delete Backend")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendDeleteBackendAPI, "delete-backend-api", "", false, "Delete Backend API")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendDeleteBackendAuth, "delete-backend-auth", "", false, "Delete Backend Auth")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendDeleteBackendStorage, "delete-backend-storage", "", false, "Delete Backend Storage")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendDeleteToken, "delete-token", "", false, "Delete Token")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendGenerateBackendAPIModels, "generate-backend-api-models", "", false, "Generate Backend API Models")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendGetBackend, "get-backend", "", false, "Get Backend")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendGetBackendAPI, "get-backend-api", "", false, "Get Backend API")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendGetBackendAPIModels, "get-backend-api-models", "", false, "Get Backend API Models")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendGetBackendAuth, "get-backend-auth", "", false, "Get Backend Auth")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendGetBackendJob, "get-backend-job", "", false, "Get Backend Job")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendGetBackendStorage, "get-backend-storage", "", false, "Get Backend Storage")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendGetToken, "get-token", "", false, "Get Token")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendImportBackendAuth, "import-backend-auth", "", false, "Import Backend Auth")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendImportBackendStorage, "import-backend-storage", "", false, "Import Backend Storage")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendListBackendJobs, "list-backend-jobs", "", false, "List Backend Jobs")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendListS3Buckets, "list-s3-buckets", "", false, "List S3 Buckets")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendRemoveAllBackends, "remove-all-backends", "", false, "Remove All Backends")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendRemoveBackendConfig, "remove-backend-config", "", false, "Remove Backend Config")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendUpdateBackendAPI, "update-backend-api", "", false, "Update Backend API")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendUpdateBackendAuth, "update-backend-auth", "", false, "Update Backend Auth")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendUpdateBackendConfig, "update-backend-config", "", false, "Update Backend Config")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendUpdateBackendJob, "update-backend-job", "", false, "Update Backend Job")
	_amplifybackendCmd.Flags().BoolVarP(&_amplifybackendUpdateBackendStorage, "update-backend-storage", "", false, "Update Backend Storage")

}
