package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/grafana"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// grafanaCmd represents the grafana command
var _grafanaCmd = &cobra.Command{
	Use:   "grafana",
	Short: "AWS grafana CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := grafana.NewFromConfig(cfg)
		if _grafanaAssociateLicense {
			grafana_AssociateLicense(cfg, client)
			return
		}
		if _grafanaCreateWorkspace {
			grafana_CreateWorkspace(cfg, client)
			return
		}
		if _grafanaCreateWorkspaceApiKey {
			grafana_CreateWorkspaceApiKey(cfg, client)
			return
		}
		if _grafanaCreateWorkspaceServiceAccount {
			grafana_CreateWorkspaceServiceAccount(cfg, client)
			return
		}
		if _grafanaCreateWorkspaceServiceAccountToken {
			grafana_CreateWorkspaceServiceAccountToken(cfg, client)
			return
		}
		if _grafanaDeleteWorkspace {
			grafana_DeleteWorkspace(cfg, client)
			return
		}
		if _grafanaDeleteWorkspaceApiKey {
			grafana_DeleteWorkspaceApiKey(cfg, client)
			return
		}
		if _grafanaDeleteWorkspaceServiceAccount {
			grafana_DeleteWorkspaceServiceAccount(cfg, client)
			return
		}
		if _grafanaDeleteWorkspaceServiceAccountToken {
			grafana_DeleteWorkspaceServiceAccountToken(cfg, client)
			return
		}
		if _grafanaDescribeWorkspace {
			grafana_DescribeWorkspace(cfg, client)
			return
		}
		if _grafanaDescribeWorkspaceAuthentication {
			grafana_DescribeWorkspaceAuthentication(cfg, client)
			return
		}
		if _grafanaDescribeWorkspaceConfiguration {
			grafana_DescribeWorkspaceConfiguration(cfg, client)
			return
		}
		if _grafanaDisassociateLicense {
			grafana_DisassociateLicense(cfg, client)
			return
		}
		if _grafanaListPermissions {
			grafana_ListPermissions(cfg, client)
			return
		}
		if _grafanaListTagsForResource {
			grafana_ListTagsForResource(cfg, client)
			return
		}
		if _grafanaListVersions {
			grafana_ListVersions(cfg, client)
			return
		}
		if _grafanaListWorkspaceServiceAccountTokens {
			grafana_ListWorkspaceServiceAccountTokens(cfg, client)
			return
		}
		if _grafanaListWorkspaceServiceAccounts {
			grafana_ListWorkspaceServiceAccounts(cfg, client)
			return
		}
		if _grafanaListWorkspaces {
			grafana_ListWorkspaces(cfg, client)
			return
		}
		if _grafanaTagResource {
			grafana_TagResource(cfg, client)
			return
		}
		if _grafanaUntagResource {
			grafana_UntagResource(cfg, client)
			return
		}
		if _grafanaUpdatePermissions {
			grafana_UpdatePermissions(cfg, client)
			return
		}
		if _grafanaUpdateWorkspace {
			grafana_UpdateWorkspace(cfg, client)
			return
		}
		if _grafanaUpdateWorkspaceAuthentication {
			grafana_UpdateWorkspaceAuthentication(cfg, client)
			return
		}
		if _grafanaUpdateWorkspaceConfiguration {
			grafana_UpdateWorkspaceConfiguration(cfg, client)
			return
		}

	},
}

var (
	_grafanaAssociateLicense                   bool
	_grafanaCreateWorkspace                    bool
	_grafanaCreateWorkspaceApiKey              bool
	_grafanaCreateWorkspaceServiceAccount      bool
	_grafanaCreateWorkspaceServiceAccountToken bool
	_grafanaDeleteWorkspace                    bool
	_grafanaDeleteWorkspaceApiKey              bool
	_grafanaDeleteWorkspaceServiceAccount      bool
	_grafanaDeleteWorkspaceServiceAccountToken bool
	_grafanaDescribeWorkspace                  bool
	_grafanaDescribeWorkspaceAuthentication    bool
	_grafanaDescribeWorkspaceConfiguration     bool
	_grafanaDisassociateLicense                bool
	_grafanaListPermissions                    bool
	_grafanaListTagsForResource                bool
	_grafanaListVersions                       bool
	_grafanaListWorkspaceServiceAccountTokens  bool
	_grafanaListWorkspaceServiceAccounts       bool
	_grafanaListWorkspaces                     bool
	_grafanaTagResource                        bool
	_grafanaUntagResource                      bool
	_grafanaUpdatePermissions                  bool
	_grafanaUpdateWorkspace                    bool
	_grafanaUpdateWorkspaceAuthentication      bool
	_grafanaUpdateWorkspaceConfiguration       bool

	_grafanaAccountAccessType                 string
	_grafanaAuthenticationProviders           string
	_grafanaClientToken                       string
	_grafanaConfiguration                     string
	_grafanaGrafanaRole                       string
	_grafanaGrafanaToken                      string
	_grafanaGrafanaVersion                    string
	_grafanaGroupId                           string
	_grafanaKeyName                           string
	_grafanaKeyRole                           string
	_grafanaKmsKeyId                          string
	_grafanaLicenseType                       string
	_grafanaMaxResults                        string
	_grafanaName                              string
	_grafanaNetworkAccessControl              string
	_grafanaNextToken                         string
	_grafanaOrganizationRoleName              string
	_grafanaPermissionType                    string
	_grafanaRemoveNetworkAccessConfiguration  string
	_grafanaRemoveVpcConfiguration            string
	_grafanaResourceArn                       string
	_grafanaSamlConfiguration                 string
	_grafanaSecondsToLive                     string
	_grafanaServiceAccountId                  string
	_grafanaStackSetName                      string
	_grafanaTagKeys                           []string
	_grafanaTags                              string
	_grafanaTokenId                           string
	_grafanaUpdateInstructionBatch            string
	_grafanaUserId                            string
	_grafanaUserType                          string
	_grafanaVpcConfiguration                  string
	_grafanaWorkspaceDataSources              string
	_grafanaWorkspaceDescription              string
	_grafanaWorkspaceId                       string
	_grafanaWorkspaceName                     string
	_grafanaWorkspaceNotificationDestinations string
	_grafanaWorkspaceOrganizationalUnits      []string
	_grafanaWorkspaceRoleArn                  string
)

// Assigns a Grafana Enterprise license to a workspace. To upgrade, you must use
// ENTERPRISE for the licenseType , and pass in a valid Grafana Labs token for the
// grafanaToken . Upgrading to Grafana Enterprise incurs additional fees. For more
// information, see [Upgrade a workspace to Grafana Enterprise].
//
// [Upgrade a workspace to Grafana Enterprise]: https://docs.aws.amazon.com/grafana/latest/userguide/upgrade-to-Grafana-Enterprise.html
func grafana_AssociateLicense(cfg aws.Config, client *grafana.Client) {
	input := &grafana.AssociateLicenseInput{
		// LicenseType: types.LicenseType, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_grafanaLicenseType) > 0 {
		if err := assignInputField(input, "LicenseType", _grafanaLicenseType); err != nil {
			log.Errorf("invalid --license-type: %s", err.Error())
			return
		}
	}
	if len(_grafanaWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_grafanaWorkspaceId)
	}
	if len(_grafanaGrafanaToken) > 0 {
		input.GrafanaToken = aws.String(_grafanaGrafanaToken)
	}

	if resp, err := client.AssociateLicense(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a workspace. In a workspace, you can create Grafana dashboards and
// visualizations to analyze your metrics, logs, and traces. You don't have to
// build, package, or deploy any hardware to run the Grafana server.
//
// Don't use CreateWorkspace to modify an existing workspace. Instead, use [UpdateWorkspace].
//
// [UpdateWorkspace]: https://docs.aws.amazon.com/grafana/latest/APIReference/API_UpdateWorkspace.html
func grafana_CreateWorkspace(cfg aws.Config, client *grafana.Client) {
	input := &grafana.CreateWorkspaceInput{
		// AccountAccessType: types.AccountAccessType, // Required
		// AuthenticationProviders: []types.AuthenticationProviderTypes, // Required
		// PermissionType: types.PermissionType, // Required
	}

	if len(_grafanaAccountAccessType) > 0 {
		if err := assignInputField(input, "AccountAccessType", _grafanaAccountAccessType); err != nil {
			log.Errorf("invalid --account-access-type: %s", err.Error())
			return
		}
	}
	if len(_grafanaAuthenticationProviders) > 0 {
		if err := assignInputField(input, "AuthenticationProviders", _grafanaAuthenticationProviders); err != nil {
			log.Errorf("invalid --authentication-providers: %s", err.Error())
			return
		}
	}
	if len(_grafanaPermissionType) > 0 {
		if err := assignInputField(input, "PermissionType", _grafanaPermissionType); err != nil {
			log.Errorf("invalid --permission-type: %s", err.Error())
			return
		}
	}
	if len(_grafanaClientToken) > 0 {
		input.ClientToken = aws.String(_grafanaClientToken)
	}
	if len(_grafanaConfiguration) > 0 {
		input.Configuration = aws.String(_grafanaConfiguration)
	}
	if len(_grafanaGrafanaVersion) > 0 {
		input.GrafanaVersion = aws.String(_grafanaGrafanaVersion)
	}
	if len(_grafanaKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_grafanaKmsKeyId)
	}
	if len(_grafanaNetworkAccessControl) > 0 {
		if err := assignInputField(input, "NetworkAccessControl", _grafanaNetworkAccessControl); err != nil {
			log.Errorf("invalid --network-access-control: %s", err.Error())
			return
		}
	}
	if len(_grafanaOrganizationRoleName) > 0 {
		input.OrganizationRoleName = aws.String(_grafanaOrganizationRoleName)
	}
	if len(_grafanaStackSetName) > 0 {
		input.StackSetName = aws.String(_grafanaStackSetName)
	}
	if len(_grafanaTags) > 0 {
		if err := assignInputField(input, "Tags", _grafanaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_grafanaVpcConfiguration) > 0 {
		if err := assignInputField(input, "VpcConfiguration", _grafanaVpcConfiguration); err != nil {
			log.Errorf("invalid --vpc-configuration: %s", err.Error())
			return
		}
	}
	if len(_grafanaWorkspaceDataSources) > 0 {
		if err := assignInputField(input, "WorkspaceDataSources", _grafanaWorkspaceDataSources); err != nil {
			log.Errorf("invalid --workspace-data-sources: %s", err.Error())
			return
		}
	}
	if len(_grafanaWorkspaceDescription) > 0 {
		input.WorkspaceDescription = aws.String(_grafanaWorkspaceDescription)
	}
	if len(_grafanaWorkspaceName) > 0 {
		input.WorkspaceName = aws.String(_grafanaWorkspaceName)
	}
	if len(_grafanaWorkspaceNotificationDestinations) > 0 {
		if err := assignInputField(input, "WorkspaceNotificationDestinations", _grafanaWorkspaceNotificationDestinations); err != nil {
			log.Errorf("invalid --workspace-notification-destinations: %s", err.Error())
			return
		}
	}
	if len(_grafanaWorkspaceOrganizationalUnits) > 0 {
		input.WorkspaceOrganizationalUnits = append([]string(nil), _grafanaWorkspaceOrganizationalUnits...)
	}
	if len(_grafanaWorkspaceRoleArn) > 0 {
		input.WorkspaceRoleArn = aws.String(_grafanaWorkspaceRoleArn)
	}

	if resp, err := client.CreateWorkspace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Grafana API key for the workspace. This key can be used to
// authenticate requests sent to the workspace's HTTP API. See [https://docs.aws.amazon.com/grafana/latest/userguide/Using-Grafana-APIs.html]for available APIs
// and example requests.
//
// In workspaces compatible with Grafana version 9 or above, use workspace service
// accounts instead of API keys. API keys will be removed in a future release.
//
// [https://docs.aws.amazon.com/grafana/latest/userguide/Using-Grafana-APIs.html]: https://docs.aws.amazon.com/grafana/latest/userguide/Using-Grafana-APIs.html
func grafana_CreateWorkspaceApiKey(cfg aws.Config, client *grafana.Client) {
	input := &grafana.CreateWorkspaceApiKeyInput{
		// KeyName: *string, // Required
		// KeyRole: *string, // Required
		// SecondsToLive: *int32, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_grafanaKeyName) > 0 {
		input.KeyName = aws.String(_grafanaKeyName)
	}
	if len(_grafanaKeyRole) > 0 {
		input.KeyRole = aws.String(_grafanaKeyRole)
	}
	if len(_grafanaSecondsToLive) > 0 {
		if err := assignInputField(input, "SecondsToLive", _grafanaSecondsToLive); err != nil {
			log.Errorf("invalid --seconds-to-live: %s", err.Error())
			return
		}
	}
	if len(_grafanaWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_grafanaWorkspaceId)
	}

	if resp, err := client.CreateWorkspaceApiKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a service account for the workspace. A service account can be used to
// call Grafana HTTP APIs, and run automated workloads. After creating the service
// account with the correct GrafanaRole for your use case, use
// CreateWorkspaceServiceAccountToken to create a token that can be used to
// authenticate and authorize Grafana HTTP API calls.
//
// You can only create service accounts for workspaces that are compatible with
// Grafana version 9 and above.
//
// For more information about service accounts, see [Service accounts] in the Amazon Managed Grafana
// User Guide.
//
// For more information about the Grafana HTTP APIs, see [Using Grafana HTTP APIs] in the Amazon Managed
// Grafana User Guide.
//
// [Service accounts]: https://docs.aws.amazon.com/grafana/latest/userguide/service-accounts.html
// [Using Grafana HTTP APIs]: https://docs.aws.amazon.com/grafana/latest/userguide/Using-Grafana-APIs.html
func grafana_CreateWorkspaceServiceAccount(cfg aws.Config, client *grafana.Client) {
	input := &grafana.CreateWorkspaceServiceAccountInput{
		// GrafanaRole: types.Role, // Required
		// Name: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_grafanaGrafanaRole) > 0 {
		if err := assignInputField(input, "GrafanaRole", _grafanaGrafanaRole); err != nil {
			log.Errorf("invalid --grafana-role: %s", err.Error())
			return
		}
	}
	if len(_grafanaName) > 0 {
		input.Name = aws.String(_grafanaName)
	}
	if len(_grafanaWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_grafanaWorkspaceId)
	}

	if resp, err := client.CreateWorkspaceServiceAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a token that can be used to authenticate and authorize Grafana HTTP API
// operations for the given [workspace service account]. The service account acts as a user for the API
// operations, and defines the permissions that are used by the API.
//
// When you create the service account token, you will receive a key that is used
// when calling Grafana APIs. Do not lose this key, as it will not be retrievable
// again.
//
// If you do lose the key, you can delete the token and recreate it to receive a
// new key. This will disable the initial key.
//
// Service accounts are only available for workspaces that are compatible with
// Grafana version 9 and above.
//
// [workspace service account]: https://docs.aws.amazon.com/grafana/latest/userguide/service-accounts.html
func grafana_CreateWorkspaceServiceAccountToken(cfg aws.Config, client *grafana.Client) {
	input := &grafana.CreateWorkspaceServiceAccountTokenInput{
		// Name: *string, // Required
		// SecondsToLive: *int32, // Required
		// ServiceAccountId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_grafanaName) > 0 {
		input.Name = aws.String(_grafanaName)
	}
	if len(_grafanaSecondsToLive) > 0 {
		if err := assignInputField(input, "SecondsToLive", _grafanaSecondsToLive); err != nil {
			log.Errorf("invalid --seconds-to-live: %s", err.Error())
			return
		}
	}
	if len(_grafanaServiceAccountId) > 0 {
		input.ServiceAccountId = aws.String(_grafanaServiceAccountId)
	}
	if len(_grafanaWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_grafanaWorkspaceId)
	}

	if resp, err := client.CreateWorkspaceServiceAccountToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Managed Grafana workspace.
func grafana_DeleteWorkspace(cfg aws.Config, client *grafana.Client) {
	input := &grafana.DeleteWorkspaceInput{
		// WorkspaceId: *string, // Required
	}

	if len(_grafanaWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_grafanaWorkspaceId)
	}

	if resp, err := client.DeleteWorkspace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Grafana API key for the workspace.
// In workspaces compatible with Grafana version 9 or above, use workspace service
// accounts instead of API keys. API keys will be removed in a future release.
func grafana_DeleteWorkspaceApiKey(cfg aws.Config, client *grafana.Client) {
	input := &grafana.DeleteWorkspaceApiKeyInput{
		// KeyName: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_grafanaKeyName) > 0 {
		input.KeyName = aws.String(_grafanaKeyName)
	}
	if len(_grafanaWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_grafanaWorkspaceId)
	}

	if resp, err := client.DeleteWorkspaceApiKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a workspace service account from the workspace.
// This will delete any tokens created for the service account, as well. If the
// tokens are currently in use, the will fail to authenticate / authorize after
// they are deleted.
//
// Service accounts are only available for workspaces that are compatible with
// Grafana version 9 and above.
func grafana_DeleteWorkspaceServiceAccount(cfg aws.Config, client *grafana.Client) {
	input := &grafana.DeleteWorkspaceServiceAccountInput{
		// ServiceAccountId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_grafanaServiceAccountId) > 0 {
		input.ServiceAccountId = aws.String(_grafanaServiceAccountId)
	}
	if len(_grafanaWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_grafanaWorkspaceId)
	}

	if resp, err := client.DeleteWorkspaceServiceAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a token for the workspace service account.
// This will disable the key associated with the token. If any automation is
// currently using the key, it will no longer be authenticated or authorized to
// perform actions with the Grafana HTTP APIs.
//
// Service accounts are only available for workspaces that are compatible with
// Grafana version 9 and above.
func grafana_DeleteWorkspaceServiceAccountToken(cfg aws.Config, client *grafana.Client) {
	input := &grafana.DeleteWorkspaceServiceAccountTokenInput{
		// ServiceAccountId: *string, // Required
		// TokenId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_grafanaServiceAccountId) > 0 {
		input.ServiceAccountId = aws.String(_grafanaServiceAccountId)
	}
	if len(_grafanaTokenId) > 0 {
		input.TokenId = aws.String(_grafanaTokenId)
	}
	if len(_grafanaWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_grafanaWorkspaceId)
	}

	if resp, err := client.DeleteWorkspaceServiceAccountToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays information about one Amazon Managed Grafana workspace.
func grafana_DescribeWorkspace(cfg aws.Config, client *grafana.Client) {
	input := &grafana.DescribeWorkspaceInput{
		// WorkspaceId: *string, // Required
	}

	if len(_grafanaWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_grafanaWorkspaceId)
	}

	if resp, err := client.DescribeWorkspace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays information about the authentication methods used in one Amazon
// Managed Grafana workspace.
func grafana_DescribeWorkspaceAuthentication(cfg aws.Config, client *grafana.Client) {
	input := &grafana.DescribeWorkspaceAuthenticationInput{
		// WorkspaceId: *string, // Required
	}

	if len(_grafanaWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_grafanaWorkspaceId)
	}

	if resp, err := client.DescribeWorkspaceAuthentication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the current configuration string for the given workspace.
func grafana_DescribeWorkspaceConfiguration(cfg aws.Config, client *grafana.Client) {
	input := &grafana.DescribeWorkspaceConfigurationInput{
		// WorkspaceId: *string, // Required
	}

	if len(_grafanaWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_grafanaWorkspaceId)
	}

	if resp, err := client.DescribeWorkspaceConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the Grafana Enterprise license from a workspace.
func grafana_DisassociateLicense(cfg aws.Config, client *grafana.Client) {
	input := &grafana.DisassociateLicenseInput{
		// LicenseType: types.LicenseType, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_grafanaLicenseType) > 0 {
		if err := assignInputField(input, "LicenseType", _grafanaLicenseType); err != nil {
			log.Errorf("invalid --license-type: %s", err.Error())
			return
		}
	}
	if len(_grafanaWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_grafanaWorkspaceId)
	}

	if resp, err := client.DisassociateLicense(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the users and groups who have the Grafana Admin and Editor roles in this
// workspace. If you use this operation without specifying userId or groupId , the
// operation returns the roles of all users and groups. If you specify a userId or
// a groupId , only the roles for that user or group are returned. If you do this,
// you can specify only one userId or one groupId .
func grafana_ListPermissions(cfg aws.Config, client *grafana.Client) {
	input := &grafana.ListPermissionsInput{
		// WorkspaceId: *string, // Required
	}

	if len(_grafanaWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_grafanaWorkspaceId)
	}
	if len(_grafanaGroupId) > 0 {
		input.GroupId = aws.String(_grafanaGroupId)
	}
	if len(_grafanaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _grafanaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_grafanaNextToken) > 0 {
		input.NextToken = aws.String(_grafanaNextToken)
	}
	if len(_grafanaUserId) > 0 {
		input.UserId = aws.String(_grafanaUserId)
	}
	if len(_grafanaUserType) > 0 {
		if err := assignInputField(input, "UserType", _grafanaUserType); err != nil {
			log.Errorf("invalid --user-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPermissions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*grafana.ListPermissionsOutput
	p := grafana.NewListPermissionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// The ListTagsForResource operation returns the tags that are associated with the
// Amazon Managed Service for Grafana resource specified by the resourceArn .
// Currently, the only resource that can be tagged is a workspace.
func grafana_ListTagsForResource(cfg aws.Config, client *grafana.Client) {
	input := &grafana.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_grafanaResourceArn) > 0 {
		input.ResourceArn = aws.String(_grafanaResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists available versions of Grafana. These are available when calling
// CreateWorkspace . Optionally, include a workspace to list the versions to which
// it can be upgraded.
func grafana_ListVersions(cfg aws.Config, client *grafana.Client) {
	input := &grafana.ListVersionsInput{}

	if len(_grafanaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _grafanaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_grafanaNextToken) > 0 {
		input.NextToken = aws.String(_grafanaNextToken)
	}
	if len(_grafanaWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_grafanaWorkspaceId)
	}

	if disablePaginator() {
		if resp, err := client.ListVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*grafana.ListVersionsOutput
	p := grafana.NewListVersionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of tokens for a workspace service account.
// This does not return the key for each token. You cannot access keys after they
// are created. To create a new key, delete the token and recreate it.
//
// Service accounts are only available for workspaces that are compatible with
// Grafana version 9 and above.
func grafana_ListWorkspaceServiceAccountTokens(cfg aws.Config, client *grafana.Client) {
	input := &grafana.ListWorkspaceServiceAccountTokensInput{
		// ServiceAccountId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_grafanaServiceAccountId) > 0 {
		input.ServiceAccountId = aws.String(_grafanaServiceAccountId)
	}
	if len(_grafanaWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_grafanaWorkspaceId)
	}
	if len(_grafanaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _grafanaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_grafanaNextToken) > 0 {
		input.NextToken = aws.String(_grafanaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkspaceServiceAccountTokens(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*grafana.ListWorkspaceServiceAccountTokensOutput
	p := grafana.NewListWorkspaceServiceAccountTokensPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of service accounts for a workspace.
// Service accounts are only available for workspaces that are compatible with
// Grafana version 9 and above.
func grafana_ListWorkspaceServiceAccounts(cfg aws.Config, client *grafana.Client) {
	input := &grafana.ListWorkspaceServiceAccountsInput{
		// WorkspaceId: *string, // Required
	}

	if len(_grafanaWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_grafanaWorkspaceId)
	}
	if len(_grafanaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _grafanaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_grafanaNextToken) > 0 {
		input.NextToken = aws.String(_grafanaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkspaceServiceAccounts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*grafana.ListWorkspaceServiceAccountsOutput
	p := grafana.NewListWorkspaceServiceAccountsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of Amazon Managed Grafana workspaces in the account, with some
// information about each workspace. For more complete information about one
// workspace, use [DescribeWorkspace].
//
// [DescribeWorkspace]: https://docs.aws.amazon.com/grafana/latest/APIReference/API_DescribeWorkspace.html
func grafana_ListWorkspaces(cfg aws.Config, client *grafana.Client) {
	input := &grafana.ListWorkspacesInput{}

	if len(_grafanaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _grafanaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_grafanaNextToken) > 0 {
		input.NextToken = aws.String(_grafanaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkspaces(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*grafana.ListWorkspacesOutput
	p := grafana.NewListWorkspacesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// The TagResource operation associates tags with an Amazon Managed Grafana
// resource. Currently, the only resource that can be tagged is workspaces.
//
// If you specify a new tag key for the resource, this tag is appended to the list
// of tags associated with the resource. If you specify a tag key that is already
// associated with the resource, the new tag value that you specify replaces the
// previous value for that tag.
func grafana_TagResource(cfg aws.Config, client *grafana.Client) {
	input := &grafana.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_grafanaResourceArn) > 0 {
		input.ResourceArn = aws.String(_grafanaResourceArn)
	}
	if len(_grafanaTags) > 0 {
		if err := assignInputField(input, "Tags", _grafanaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The UntagResource operation removes the association of the tag with the Amazon
// Managed Grafana resource.
func grafana_UntagResource(cfg aws.Config, client *grafana.Client) {
	input := &grafana.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_grafanaResourceArn) > 0 {
		input.ResourceArn = aws.String(_grafanaResourceArn)
	}
	if len(_grafanaTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _grafanaTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates which users in a workspace have the Grafana Admin or Editor roles.
func grafana_UpdatePermissions(cfg aws.Config, client *grafana.Client) {
	input := &grafana.UpdatePermissionsInput{
		// UpdateInstructionBatch: []types.UpdateInstruction, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_grafanaUpdateInstructionBatch) > 0 {
		if err := assignInputField(input, "UpdateInstructionBatch", _grafanaUpdateInstructionBatch); err != nil {
			log.Errorf("invalid --update-instruction-batch: %s", err.Error())
			return
		}
	}
	if len(_grafanaWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_grafanaWorkspaceId)
	}

	if resp, err := client.UpdatePermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an existing Amazon Managed Grafana workspace. If you use this
// operation and omit any optional parameters, the existing values of those
// parameters are not changed.
//
// To modify the user authentication methods that the workspace uses, such as SAML
// or IAM Identity Center, use [UpdateWorkspaceAuthentication].
//
// To modify which users in the workspace have the Admin and Editor Grafana roles,
// use [UpdatePermissions].
//
// [UpdatePermissions]: https://docs.aws.amazon.com/grafana/latest/APIReference/API_UpdatePermissions.html
// [UpdateWorkspaceAuthentication]: https://docs.aws.amazon.com/grafana/latest/APIReference/API_UpdateWorkspaceAuthentication.html
func grafana_UpdateWorkspace(cfg aws.Config, client *grafana.Client) {
	input := &grafana.UpdateWorkspaceInput{
		// WorkspaceId: *string, // Required
	}

	if len(_grafanaWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_grafanaWorkspaceId)
	}
	if len(_grafanaAccountAccessType) > 0 {
		if err := assignInputField(input, "AccountAccessType", _grafanaAccountAccessType); err != nil {
			log.Errorf("invalid --account-access-type: %s", err.Error())
			return
		}
	}
	if len(_grafanaNetworkAccessControl) > 0 {
		if err := assignInputField(input, "NetworkAccessControl", _grafanaNetworkAccessControl); err != nil {
			log.Errorf("invalid --network-access-control: %s", err.Error())
			return
		}
	}
	if len(_grafanaOrganizationRoleName) > 0 {
		input.OrganizationRoleName = aws.String(_grafanaOrganizationRoleName)
	}
	if len(_grafanaPermissionType) > 0 {
		if err := assignInputField(input, "PermissionType", _grafanaPermissionType); err != nil {
			log.Errorf("invalid --permission-type: %s", err.Error())
			return
		}
	}
	if len(_grafanaRemoveNetworkAccessConfiguration) > 0 {
		if err := assignInputField(input, "RemoveNetworkAccessConfiguration", _grafanaRemoveNetworkAccessConfiguration); err != nil {
			log.Errorf("invalid --remove-network-access-configuration: %s", err.Error())
			return
		}
	}
	if len(_grafanaRemoveVpcConfiguration) > 0 {
		if err := assignInputField(input, "RemoveVpcConfiguration", _grafanaRemoveVpcConfiguration); err != nil {
			log.Errorf("invalid --remove-vpc-configuration: %s", err.Error())
			return
		}
	}
	if len(_grafanaStackSetName) > 0 {
		input.StackSetName = aws.String(_grafanaStackSetName)
	}
	if len(_grafanaVpcConfiguration) > 0 {
		if err := assignInputField(input, "VpcConfiguration", _grafanaVpcConfiguration); err != nil {
			log.Errorf("invalid --vpc-configuration: %s", err.Error())
			return
		}
	}
	if len(_grafanaWorkspaceDataSources) > 0 {
		if err := assignInputField(input, "WorkspaceDataSources", _grafanaWorkspaceDataSources); err != nil {
			log.Errorf("invalid --workspace-data-sources: %s", err.Error())
			return
		}
	}
	if len(_grafanaWorkspaceDescription) > 0 {
		input.WorkspaceDescription = aws.String(_grafanaWorkspaceDescription)
	}
	if len(_grafanaWorkspaceName) > 0 {
		input.WorkspaceName = aws.String(_grafanaWorkspaceName)
	}
	if len(_grafanaWorkspaceNotificationDestinations) > 0 {
		if err := assignInputField(input, "WorkspaceNotificationDestinations", _grafanaWorkspaceNotificationDestinations); err != nil {
			log.Errorf("invalid --workspace-notification-destinations: %s", err.Error())
			return
		}
	}
	if len(_grafanaWorkspaceOrganizationalUnits) > 0 {
		input.WorkspaceOrganizationalUnits = append([]string(nil), _grafanaWorkspaceOrganizationalUnits...)
	}
	if len(_grafanaWorkspaceRoleArn) > 0 {
		input.WorkspaceRoleArn = aws.String(_grafanaWorkspaceRoleArn)
	}

	if resp, err := client.UpdateWorkspace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to define the identity provider (IdP) that this workspace
// authenticates users from, using SAML. You can also map SAML assertion attributes
// to workspace user information and define which groups in the assertion attribute
// are to have the Admin and Editor roles in the workspace.
//
// Changes to the authentication method for a workspace may take a few minutes to
// take effect.
func grafana_UpdateWorkspaceAuthentication(cfg aws.Config, client *grafana.Client) {
	input := &grafana.UpdateWorkspaceAuthenticationInput{
		// AuthenticationProviders: []types.AuthenticationProviderTypes, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_grafanaAuthenticationProviders) > 0 {
		if err := assignInputField(input, "AuthenticationProviders", _grafanaAuthenticationProviders); err != nil {
			log.Errorf("invalid --authentication-providers: %s", err.Error())
			return
		}
	}
	if len(_grafanaWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_grafanaWorkspaceId)
	}
	if len(_grafanaSamlConfiguration) > 0 {
		if err := assignInputField(input, "SamlConfiguration", _grafanaSamlConfiguration); err != nil {
			log.Errorf("invalid --saml-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateWorkspaceAuthentication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration string for the given workspace
func grafana_UpdateWorkspaceConfiguration(cfg aws.Config, client *grafana.Client) {
	input := &grafana.UpdateWorkspaceConfigurationInput{
		// Configuration: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_grafanaConfiguration) > 0 {
		input.Configuration = aws.String(_grafanaConfiguration)
	}
	if len(_grafanaWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_grafanaWorkspaceId)
	}
	if len(_grafanaGrafanaVersion) > 0 {
		input.GrafanaVersion = aws.String(_grafanaGrafanaVersion)
	}

	if resp, err := client.UpdateWorkspaceConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_grafanaCmd)
	_grafanaCmd.Flags().SortFlags = false

	_grafanaCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_grafanaCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_grafanaCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_grafanaCmd.Flags().StringVarP(&_grafanaAccountAccessType, "account-access-type", "", "", "Account Access Type")
	_grafanaCmd.Flags().StringVarP(&_grafanaAuthenticationProviders, "authentication-providers", "", "", "Authentication Providers")
	_grafanaCmd.Flags().StringVarP(&_grafanaClientToken, "client-token", "", "", "Client Token")
	_grafanaCmd.Flags().StringVarP(&_grafanaConfiguration, "configuration", "", "", "Configuration")
	_grafanaCmd.Flags().StringVarP(&_grafanaGrafanaRole, "grafana-role", "", "", "Grafana Role")
	_grafanaCmd.Flags().StringVarP(&_grafanaGrafanaToken, "grafana-token", "", "", "Grafana Token")
	_grafanaCmd.Flags().StringVarP(&_grafanaGrafanaVersion, "grafana-version", "", "", "Grafana Version")
	_grafanaCmd.Flags().StringVarP(&_grafanaGroupId, "group-id", "", "", "Group ID")
	_grafanaCmd.Flags().StringVarP(&_grafanaKeyName, "key-name", "", "", "Key Name")
	_grafanaCmd.Flags().StringVarP(&_grafanaKeyRole, "key-role", "", "", "Key Role")
	_grafanaCmd.Flags().StringVarP(&_grafanaKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_grafanaCmd.Flags().StringVarP(&_grafanaLicenseType, "license-type", "", "", "License Type")
	_grafanaCmd.Flags().StringVarP(&_grafanaMaxResults, "max-results", "", "", "Max Results")
	_grafanaCmd.Flags().StringVarP(&_grafanaName, "name", "", "", "Name")
	_grafanaCmd.Flags().StringVarP(&_grafanaNetworkAccessControl, "network-access-control", "", "", "Network Access Control")
	_grafanaCmd.Flags().StringVarP(&_grafanaNextToken, "next-token", "", "", "Next Token")
	_grafanaCmd.Flags().StringVarP(&_grafanaOrganizationRoleName, "organization-role-name", "", "", "Organization Role Name")
	_grafanaCmd.Flags().StringVarP(&_grafanaPermissionType, "permission-type", "", "", "Permission Type")
	_grafanaCmd.Flags().StringVarP(&_grafanaRemoveNetworkAccessConfiguration, "remove-network-access-configuration", "", "", "Remove Network Access Configuration")
	_grafanaCmd.Flags().StringVarP(&_grafanaRemoveVpcConfiguration, "remove-vpc-configuration", "", "", "Remove VPC Configuration")
	_grafanaCmd.Flags().StringVarP(&_grafanaResourceArn, "resource-arn", "", "", "Resource ARN")
	_grafanaCmd.Flags().StringVarP(&_grafanaSamlConfiguration, "saml-configuration", "", "", "Saml Configuration")
	_grafanaCmd.Flags().StringVarP(&_grafanaSecondsToLive, "seconds-to-live", "", "", "Seconds To Live")
	_grafanaCmd.Flags().StringVarP(&_grafanaServiceAccountId, "service-account-id", "", "", "Service Account ID")
	_grafanaCmd.Flags().StringVarP(&_grafanaStackSetName, "stack-set-name", "", "", "Stack Set Name")
	_grafanaCmd.Flags().StringSliceVarP(&_grafanaTagKeys, "tag-keys", "", nil, "Tag Keys")
	_grafanaCmd.Flags().StringVarP(&_grafanaTags, "tags", "", "", "Tags")
	_grafanaCmd.Flags().StringVarP(&_grafanaTokenId, "token-id", "", "", "Token ID")
	_grafanaCmd.Flags().StringVarP(&_grafanaUpdateInstructionBatch, "update-instruction-batch", "", "", "Update Instruction Batch")
	_grafanaCmd.Flags().StringVarP(&_grafanaUserId, "user-id", "", "", "User ID")
	_grafanaCmd.Flags().StringVarP(&_grafanaUserType, "user-type", "", "", "User Type")
	_grafanaCmd.Flags().StringVarP(&_grafanaVpcConfiguration, "vpc-configuration", "", "", "VPC Configuration")
	_grafanaCmd.Flags().StringVarP(&_grafanaWorkspaceDataSources, "workspace-data-sources", "", "", "Workspace Data Sources")
	_grafanaCmd.Flags().StringVarP(&_grafanaWorkspaceDescription, "workspace-description", "", "", "Workspace Description")
	_grafanaCmd.Flags().StringVarP(&_grafanaWorkspaceId, "workspace-id", "", "", "Workspace ID")
	_grafanaCmd.Flags().StringVarP(&_grafanaWorkspaceName, "workspace-name", "", "", "Workspace Name")
	_grafanaCmd.Flags().StringVarP(&_grafanaWorkspaceNotificationDestinations, "workspace-notification-destinations", "", "", "Workspace Notification Destinations")
	_grafanaCmd.Flags().StringSliceVarP(&_grafanaWorkspaceOrganizationalUnits, "workspace-organizational-units", "", nil, "Workspace Organizational Units")
	_grafanaCmd.Flags().StringVarP(&_grafanaWorkspaceRoleArn, "workspace-role-arn", "", "", "Workspace Role ARN")

	_grafanaCmd.Flags().BoolVarP(&_grafanaAssociateLicense, "associate-license", "", false, "Associate License")
	_grafanaCmd.Flags().BoolVarP(&_grafanaCreateWorkspace, "create-workspace", "", false, "Create Workspace")
	_grafanaCmd.Flags().BoolVarP(&_grafanaCreateWorkspaceApiKey, "create-workspace-api-key", "", false, "Create Workspace API Key")
	_grafanaCmd.Flags().BoolVarP(&_grafanaCreateWorkspaceServiceAccount, "create-workspace-service-account", "", false, "Create Workspace Service Account")
	_grafanaCmd.Flags().BoolVarP(&_grafanaCreateWorkspaceServiceAccountToken, "create-workspace-service-account-token", "", false, "Create Workspace Service Account Token")
	_grafanaCmd.Flags().BoolVarP(&_grafanaDeleteWorkspace, "delete-workspace", "", false, "Delete Workspace")
	_grafanaCmd.Flags().BoolVarP(&_grafanaDeleteWorkspaceApiKey, "delete-workspace-api-key", "", false, "Delete Workspace API Key")
	_grafanaCmd.Flags().BoolVarP(&_grafanaDeleteWorkspaceServiceAccount, "delete-workspace-service-account", "", false, "Delete Workspace Service Account")
	_grafanaCmd.Flags().BoolVarP(&_grafanaDeleteWorkspaceServiceAccountToken, "delete-workspace-service-account-token", "", false, "Delete Workspace Service Account Token")
	_grafanaCmd.Flags().BoolVarP(&_grafanaDescribeWorkspace, "describe-workspace", "", false, "Describe Workspace")
	_grafanaCmd.Flags().BoolVarP(&_grafanaDescribeWorkspaceAuthentication, "describe-workspace-authentication", "", false, "Describe Workspace Authentication")
	_grafanaCmd.Flags().BoolVarP(&_grafanaDescribeWorkspaceConfiguration, "describe-workspace-configuration", "", false, "Describe Workspace Configuration")
	_grafanaCmd.Flags().BoolVarP(&_grafanaDisassociateLicense, "disassociate-license", "", false, "Disassociate License")
	_grafanaCmd.Flags().BoolVarP(&_grafanaListPermissions, "list-permissions", "", false, "List Permissions")
	_grafanaCmd.Flags().BoolVarP(&_grafanaListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_grafanaCmd.Flags().BoolVarP(&_grafanaListVersions, "list-versions", "", false, "List Versions")
	_grafanaCmd.Flags().BoolVarP(&_grafanaListWorkspaceServiceAccountTokens, "list-workspace-service-account-tokens", "", false, "List Workspace Service Account Tokens")
	_grafanaCmd.Flags().BoolVarP(&_grafanaListWorkspaceServiceAccounts, "list-workspace-service-accounts", "", false, "List Workspace Service Accounts")
	_grafanaCmd.Flags().BoolVarP(&_grafanaListWorkspaces, "list-workspaces", "", false, "List Workspaces")
	_grafanaCmd.Flags().BoolVarP(&_grafanaTagResource, "tag-resource", "", false, "Tag Resource")
	_grafanaCmd.Flags().BoolVarP(&_grafanaUntagResource, "untag-resource", "", false, "Untag Resource")
	_grafanaCmd.Flags().BoolVarP(&_grafanaUpdatePermissions, "update-permissions", "", false, "Update Permissions")
	_grafanaCmd.Flags().BoolVarP(&_grafanaUpdateWorkspace, "update-workspace", "", false, "Update Workspace")
	_grafanaCmd.Flags().BoolVarP(&_grafanaUpdateWorkspaceAuthentication, "update-workspace-authentication", "", false, "Update Workspace Authentication")
	_grafanaCmd.Flags().BoolVarP(&_grafanaUpdateWorkspaceConfiguration, "update-workspace-configuration", "", false, "Update Workspace Configuration")

}
