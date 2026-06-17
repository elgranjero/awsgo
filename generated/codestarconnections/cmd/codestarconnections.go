package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codestarconnections"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// codestarconnectionsCmd represents the codestarconnections command
var _codestarconnectionsCmd = &cobra.Command{
	Use:   "codestarconnections",
	Short: "AWS codestarconnections CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := codestarconnections.NewFromConfig(cfg)
		if _codestarconnectionsCreateConnection {
			codestarconnections_CreateConnection(cfg, client)
			return
		}
		if _codestarconnectionsCreateHost {
			codestarconnections_CreateHost(cfg, client)
			return
		}
		if _codestarconnectionsCreateRepositoryLink {
			codestarconnections_CreateRepositoryLink(cfg, client)
			return
		}
		if _codestarconnectionsCreateSyncConfiguration {
			codestarconnections_CreateSyncConfiguration(cfg, client)
			return
		}
		if _codestarconnectionsDeleteConnection {
			codestarconnections_DeleteConnection(cfg, client)
			return
		}
		if _codestarconnectionsDeleteHost {
			codestarconnections_DeleteHost(cfg, client)
			return
		}
		if _codestarconnectionsDeleteRepositoryLink {
			codestarconnections_DeleteRepositoryLink(cfg, client)
			return
		}
		if _codestarconnectionsDeleteSyncConfiguration {
			codestarconnections_DeleteSyncConfiguration(cfg, client)
			return
		}
		if _codestarconnectionsGetConnection {
			codestarconnections_GetConnection(cfg, client)
			return
		}
		if _codestarconnectionsGetHost {
			codestarconnections_GetHost(cfg, client)
			return
		}
		if _codestarconnectionsGetRepositoryLink {
			codestarconnections_GetRepositoryLink(cfg, client)
			return
		}
		if _codestarconnectionsGetRepositorySyncStatus {
			codestarconnections_GetRepositorySyncStatus(cfg, client)
			return
		}
		if _codestarconnectionsGetResourceSyncStatus {
			codestarconnections_GetResourceSyncStatus(cfg, client)
			return
		}
		if _codestarconnectionsGetSyncBlockerSummary {
			codestarconnections_GetSyncBlockerSummary(cfg, client)
			return
		}
		if _codestarconnectionsGetSyncConfiguration {
			codestarconnections_GetSyncConfiguration(cfg, client)
			return
		}
		if _codestarconnectionsListConnections {
			codestarconnections_ListConnections(cfg, client)
			return
		}
		if _codestarconnectionsListHosts {
			codestarconnections_ListHosts(cfg, client)
			return
		}
		if _codestarconnectionsListRepositoryLinks {
			codestarconnections_ListRepositoryLinks(cfg, client)
			return
		}
		if _codestarconnectionsListRepositorySyncDefinitions {
			codestarconnections_ListRepositorySyncDefinitions(cfg, client)
			return
		}
		if _codestarconnectionsListSyncConfigurations {
			codestarconnections_ListSyncConfigurations(cfg, client)
			return
		}
		if _codestarconnectionsListTagsForResource {
			codestarconnections_ListTagsForResource(cfg, client)
			return
		}
		if _codestarconnectionsTagResource {
			codestarconnections_TagResource(cfg, client)
			return
		}
		if _codestarconnectionsUntagResource {
			codestarconnections_UntagResource(cfg, client)
			return
		}
		if _codestarconnectionsUpdateHost {
			codestarconnections_UpdateHost(cfg, client)
			return
		}
		if _codestarconnectionsUpdateRepositoryLink {
			codestarconnections_UpdateRepositoryLink(cfg, client)
			return
		}
		if _codestarconnectionsUpdateSyncBlocker {
			codestarconnections_UpdateSyncBlocker(cfg, client)
			return
		}
		if _codestarconnectionsUpdateSyncConfiguration {
			codestarconnections_UpdateSyncConfiguration(cfg, client)
			return
		}

	},
}

var (
	_codestarconnectionsCreateConnection              bool
	_codestarconnectionsCreateHost                    bool
	_codestarconnectionsCreateRepositoryLink          bool
	_codestarconnectionsCreateSyncConfiguration       bool
	_codestarconnectionsDeleteConnection              bool
	_codestarconnectionsDeleteHost                    bool
	_codestarconnectionsDeleteRepositoryLink          bool
	_codestarconnectionsDeleteSyncConfiguration       bool
	_codestarconnectionsGetConnection                 bool
	_codestarconnectionsGetHost                       bool
	_codestarconnectionsGetRepositoryLink             bool
	_codestarconnectionsGetRepositorySyncStatus       bool
	_codestarconnectionsGetResourceSyncStatus         bool
	_codestarconnectionsGetSyncBlockerSummary         bool
	_codestarconnectionsGetSyncConfiguration          bool
	_codestarconnectionsListConnections               bool
	_codestarconnectionsListHosts                     bool
	_codestarconnectionsListRepositoryLinks           bool
	_codestarconnectionsListRepositorySyncDefinitions bool
	_codestarconnectionsListSyncConfigurations        bool
	_codestarconnectionsListTagsForResource           bool
	_codestarconnectionsTagResource                   bool
	_codestarconnectionsUntagResource                 bool
	_codestarconnectionsUpdateHost                    bool
	_codestarconnectionsUpdateRepositoryLink          bool
	_codestarconnectionsUpdateSyncBlocker             bool
	_codestarconnectionsUpdateSyncConfiguration       bool

	_codestarconnectionsBranch                  string
	_codestarconnectionsConfigFile              string
	_codestarconnectionsConnectionArn           string
	_codestarconnectionsConnectionName          string
	_codestarconnectionsEncryptionKeyArn        string
	_codestarconnectionsHostArn                 string
	_codestarconnectionsHostArnFilter           string
	_codestarconnectionsId                      string
	_codestarconnectionsMaxResults              string
	_codestarconnectionsName                    string
	_codestarconnectionsNextToken               string
	_codestarconnectionsOwnerId                 string
	_codestarconnectionsProviderEndpoint        string
	_codestarconnectionsProviderType            string
	_codestarconnectionsProviderTypeFilter      string
	_codestarconnectionsPublishDeploymentStatus string
	_codestarconnectionsRepositoryLinkId        string
	_codestarconnectionsRepositoryName          string
	_codestarconnectionsResolvedReason          string
	_codestarconnectionsResourceArn             string
	_codestarconnectionsResourceName            string
	_codestarconnectionsRoleArn                 string
	_codestarconnectionsSyncType                string
	_codestarconnectionsTagKeys                 []string
	_codestarconnectionsTags                    string
	_codestarconnectionsTriggerResourceUpdateOn string
	_codestarconnectionsVpcConfiguration        string
)

// Creates a connection that can then be given to other Amazon Web Services
// services like CodePipeline so that it can access third-party code repositories.
// The connection is in pending status until the third-party connection handshake
// is completed from the console.
func codestarconnections_CreateConnection(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.CreateConnectionInput{
		// ConnectionName: *string, // Required
	}

	if len(_codestarconnectionsConnectionName) > 0 {
		input.ConnectionName = aws.String(_codestarconnectionsConnectionName)
	}
	if len(_codestarconnectionsHostArn) > 0 {
		input.HostArn = aws.String(_codestarconnectionsHostArn)
	}
	if len(_codestarconnectionsProviderType) > 0 {
		if err := assignInputField(input, "ProviderType", _codestarconnectionsProviderType); err != nil {
			log.Errorf("invalid --provider-type: %s", err.Error())
			return
		}
	}
	if len(_codestarconnectionsTags) > 0 {
		if err := assignInputField(input, "Tags", _codestarconnectionsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a resource that represents the infrastructure where a third-party
// provider is installed. The host is used when you create connections to an
// installed third-party provider type, such as GitHub Enterprise Server. You
// create one host for all connections to that provider.
//
// A host created through the CLI or the SDK is in `PENDING` status by default.
// You can make its status `AVAILABLE` by setting up the host in the console.
func codestarconnections_CreateHost(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.CreateHostInput{
		// Name: *string, // Required
		// ProviderEndpoint: *string, // Required
		// ProviderType: types.ProviderType, // Required
	}

	if len(_codestarconnectionsName) > 0 {
		input.Name = aws.String(_codestarconnectionsName)
	}
	if len(_codestarconnectionsProviderEndpoint) > 0 {
		input.ProviderEndpoint = aws.String(_codestarconnectionsProviderEndpoint)
	}
	if len(_codestarconnectionsProviderType) > 0 {
		if err := assignInputField(input, "ProviderType", _codestarconnectionsProviderType); err != nil {
			log.Errorf("invalid --provider-type: %s", err.Error())
			return
		}
	}
	if len(_codestarconnectionsTags) > 0 {
		if err := assignInputField(input, "Tags", _codestarconnectionsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_codestarconnectionsVpcConfiguration) > 0 {
		if err := assignInputField(input, "VpcConfiguration", _codestarconnectionsVpcConfiguration); err != nil {
			log.Errorf("invalid --vpc-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateHost(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a link to a specified external Git repository. A repository link allows
// Git sync to monitor and sync changes to files in a specified Git repository.
func codestarconnections_CreateRepositoryLink(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.CreateRepositoryLinkInput{
		// ConnectionArn: *string, // Required
		// OwnerId: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_codestarconnectionsConnectionArn) > 0 {
		input.ConnectionArn = aws.String(_codestarconnectionsConnectionArn)
	}
	if len(_codestarconnectionsOwnerId) > 0 {
		input.OwnerId = aws.String(_codestarconnectionsOwnerId)
	}
	if len(_codestarconnectionsRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codestarconnectionsRepositoryName)
	}
	if len(_codestarconnectionsEncryptionKeyArn) > 0 {
		input.EncryptionKeyArn = aws.String(_codestarconnectionsEncryptionKeyArn)
	}
	if len(_codestarconnectionsTags) > 0 {
		if err := assignInputField(input, "Tags", _codestarconnectionsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRepositoryLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a sync configuration which allows Amazon Web Services to sync content
// from a Git repository to update a specified Amazon Web Services resource.
// Parameters for the sync configuration are determined by the sync type.
func codestarconnections_CreateSyncConfiguration(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.CreateSyncConfigurationInput{
		// Branch: *string, // Required
		// ConfigFile: *string, // Required
		// RepositoryLinkId: *string, // Required
		// ResourceName: *string, // Required
		// RoleArn: *string, // Required
		// SyncType: types.SyncConfigurationType, // Required
	}

	if len(_codestarconnectionsBranch) > 0 {
		input.Branch = aws.String(_codestarconnectionsBranch)
	}
	if len(_codestarconnectionsConfigFile) > 0 {
		input.ConfigFile = aws.String(_codestarconnectionsConfigFile)
	}
	if len(_codestarconnectionsRepositoryLinkId) > 0 {
		input.RepositoryLinkId = aws.String(_codestarconnectionsRepositoryLinkId)
	}
	if len(_codestarconnectionsResourceName) > 0 {
		input.ResourceName = aws.String(_codestarconnectionsResourceName)
	}
	if len(_codestarconnectionsRoleArn) > 0 {
		input.RoleArn = aws.String(_codestarconnectionsRoleArn)
	}
	if len(_codestarconnectionsSyncType) > 0 {
		if err := assignInputField(input, "SyncType", _codestarconnectionsSyncType); err != nil {
			log.Errorf("invalid --sync-type: %s", err.Error())
			return
		}
	}
	if len(_codestarconnectionsPublishDeploymentStatus) > 0 {
		if err := assignInputField(input, "PublishDeploymentStatus", _codestarconnectionsPublishDeploymentStatus); err != nil {
			log.Errorf("invalid --publish-deployment-status: %s", err.Error())
			return
		}
	}
	if len(_codestarconnectionsTriggerResourceUpdateOn) > 0 {
		if err := assignInputField(input, "TriggerResourceUpdateOn", _codestarconnectionsTriggerResourceUpdateOn); err != nil {
			log.Errorf("invalid --trigger-resource-update-on: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSyncConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The connection to be deleted.
func codestarconnections_DeleteConnection(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.DeleteConnectionInput{
		// ConnectionArn: *string, // Required
	}

	if len(_codestarconnectionsConnectionArn) > 0 {
		input.ConnectionArn = aws.String(_codestarconnectionsConnectionArn)
	}

	if resp, err := client.DeleteConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The host to be deleted. Before you delete a host, all connections associated to
// the host must be deleted.
//
// A host cannot be deleted if it is in the VPC_CONFIG_INITIALIZING or
// VPC_CONFIG_DELETING state.
func codestarconnections_DeleteHost(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.DeleteHostInput{
		// HostArn: *string, // Required
	}

	if len(_codestarconnectionsHostArn) > 0 {
		input.HostArn = aws.String(_codestarconnectionsHostArn)
	}

	if resp, err := client.DeleteHost(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the association between your connection and a specified external Git
// repository.
func codestarconnections_DeleteRepositoryLink(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.DeleteRepositoryLinkInput{
		// RepositoryLinkId: *string, // Required
	}

	if len(_codestarconnectionsRepositoryLinkId) > 0 {
		input.RepositoryLinkId = aws.String(_codestarconnectionsRepositoryLinkId)
	}

	if resp, err := client.DeleteRepositoryLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the sync configuration for a specified repository and connection.
func codestarconnections_DeleteSyncConfiguration(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.DeleteSyncConfigurationInput{
		// ResourceName: *string, // Required
		// SyncType: types.SyncConfigurationType, // Required
	}

	if len(_codestarconnectionsResourceName) > 0 {
		input.ResourceName = aws.String(_codestarconnectionsResourceName)
	}
	if len(_codestarconnectionsSyncType) > 0 {
		if err := assignInputField(input, "SyncType", _codestarconnectionsSyncType); err != nil {
			log.Errorf("invalid --sync-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteSyncConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the connection ARN and details such as status, owner, and provider type.
func codestarconnections_GetConnection(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.GetConnectionInput{
		// ConnectionArn: *string, // Required
	}

	if len(_codestarconnectionsConnectionArn) > 0 {
		input.ConnectionArn = aws.String(_codestarconnectionsConnectionArn)
	}

	if resp, err := client.GetConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the host ARN and details such as status, provider type, endpoint, and,
// if applicable, the VPC configuration.
func codestarconnections_GetHost(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.GetHostInput{
		// HostArn: *string, // Required
	}

	if len(_codestarconnectionsHostArn) > 0 {
		input.HostArn = aws.String(_codestarconnectionsHostArn)
	}

	if resp, err := client.GetHost(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details about a repository link. A repository link allows Git sync to
// monitor and sync changes from files in a specified Git repository.
func codestarconnections_GetRepositoryLink(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.GetRepositoryLinkInput{
		// RepositoryLinkId: *string, // Required
	}

	if len(_codestarconnectionsRepositoryLinkId) > 0 {
		input.RepositoryLinkId = aws.String(_codestarconnectionsRepositoryLinkId)
	}

	if resp, err := client.GetRepositoryLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details about the sync status for a repository. A repository sync uses
// Git sync to push and pull changes from your remote repository.
func codestarconnections_GetRepositorySyncStatus(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.GetRepositorySyncStatusInput{
		// Branch: *string, // Required
		// RepositoryLinkId: *string, // Required
		// SyncType: types.SyncConfigurationType, // Required
	}

	if len(_codestarconnectionsBranch) > 0 {
		input.Branch = aws.String(_codestarconnectionsBranch)
	}
	if len(_codestarconnectionsRepositoryLinkId) > 0 {
		input.RepositoryLinkId = aws.String(_codestarconnectionsRepositoryLinkId)
	}
	if len(_codestarconnectionsSyncType) > 0 {
		if err := assignInputField(input, "SyncType", _codestarconnectionsSyncType); err != nil {
			log.Errorf("invalid --sync-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetRepositorySyncStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the status of the sync with the Git repository for a specific Amazon
// Web Services resource.
func codestarconnections_GetResourceSyncStatus(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.GetResourceSyncStatusInput{
		// ResourceName: *string, // Required
		// SyncType: types.SyncConfigurationType, // Required
	}

	if len(_codestarconnectionsResourceName) > 0 {
		input.ResourceName = aws.String(_codestarconnectionsResourceName)
	}
	if len(_codestarconnectionsSyncType) > 0 {
		if err := assignInputField(input, "SyncType", _codestarconnectionsSyncType); err != nil {
			log.Errorf("invalid --sync-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetResourceSyncStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of the most recent sync blockers.
func codestarconnections_GetSyncBlockerSummary(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.GetSyncBlockerSummaryInput{
		// ResourceName: *string, // Required
		// SyncType: types.SyncConfigurationType, // Required
	}

	if len(_codestarconnectionsResourceName) > 0 {
		input.ResourceName = aws.String(_codestarconnectionsResourceName)
	}
	if len(_codestarconnectionsSyncType) > 0 {
		if err := assignInputField(input, "SyncType", _codestarconnectionsSyncType); err != nil {
			log.Errorf("invalid --sync-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetSyncBlockerSummary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details about a sync configuration, including the sync type and
// resource name. A sync configuration allows the configuration to sync (push and
// pull) changes from the remote repository for a specified branch in a Git
// repository.
func codestarconnections_GetSyncConfiguration(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.GetSyncConfigurationInput{
		// ResourceName: *string, // Required
		// SyncType: types.SyncConfigurationType, // Required
	}

	if len(_codestarconnectionsResourceName) > 0 {
		input.ResourceName = aws.String(_codestarconnectionsResourceName)
	}
	if len(_codestarconnectionsSyncType) > 0 {
		if err := assignInputField(input, "SyncType", _codestarconnectionsSyncType); err != nil {
			log.Errorf("invalid --sync-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetSyncConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the connections associated with your account.
func codestarconnections_ListConnections(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.ListConnectionsInput{}

	if len(_codestarconnectionsHostArnFilter) > 0 {
		input.HostArnFilter = aws.String(_codestarconnectionsHostArnFilter)
	}
	if len(_codestarconnectionsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codestarconnectionsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codestarconnectionsNextToken) > 0 {
		input.NextToken = aws.String(_codestarconnectionsNextToken)
	}
	if len(_codestarconnectionsProviderTypeFilter) > 0 {
		if err := assignInputField(input, "ProviderTypeFilter", _codestarconnectionsProviderTypeFilter); err != nil {
			log.Errorf("invalid --provider-type-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListConnections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codestarconnections.ListConnectionsOutput
	p := codestarconnections.NewListConnectionsPaginator(client, input)
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

// Lists the hosts associated with your account.
func codestarconnections_ListHosts(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.ListHostsInput{}

	if len(_codestarconnectionsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codestarconnectionsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codestarconnectionsNextToken) > 0 {
		input.NextToken = aws.String(_codestarconnectionsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListHosts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codestarconnections.ListHostsOutput
	p := codestarconnections.NewListHostsPaginator(client, input)
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

// Lists the repository links created for connections in your account.
func codestarconnections_ListRepositoryLinks(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.ListRepositoryLinksInput{}

	if len(_codestarconnectionsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codestarconnectionsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codestarconnectionsNextToken) > 0 {
		input.NextToken = aws.String(_codestarconnectionsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRepositoryLinks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codestarconnections.ListRepositoryLinksOutput
	p := codestarconnections.NewListRepositoryLinksPaginator(client, input)
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

// Lists the repository sync definitions for repository links in your account.
func codestarconnections_ListRepositorySyncDefinitions(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.ListRepositorySyncDefinitionsInput{
		// RepositoryLinkId: *string, // Required
		// SyncType: types.SyncConfigurationType, // Required
	}

	if len(_codestarconnectionsRepositoryLinkId) > 0 {
		input.RepositoryLinkId = aws.String(_codestarconnectionsRepositoryLinkId)
	}
	if len(_codestarconnectionsSyncType) > 0 {
		if err := assignInputField(input, "SyncType", _codestarconnectionsSyncType); err != nil {
			log.Errorf("invalid --sync-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListRepositorySyncDefinitions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of sync configurations for a specified repository.
func codestarconnections_ListSyncConfigurations(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.ListSyncConfigurationsInput{
		// RepositoryLinkId: *string, // Required
		// SyncType: types.SyncConfigurationType, // Required
	}

	if len(_codestarconnectionsRepositoryLinkId) > 0 {
		input.RepositoryLinkId = aws.String(_codestarconnectionsRepositoryLinkId)
	}
	if len(_codestarconnectionsSyncType) > 0 {
		if err := assignInputField(input, "SyncType", _codestarconnectionsSyncType); err != nil {
			log.Errorf("invalid --sync-type: %s", err.Error())
			return
		}
	}
	if len(_codestarconnectionsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codestarconnectionsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codestarconnectionsNextToken) > 0 {
		input.NextToken = aws.String(_codestarconnectionsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSyncConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codestarconnections.ListSyncConfigurationsOutput
	p := codestarconnections.NewListSyncConfigurationsPaginator(client, input)
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

// Gets the set of key-value pairs (metadata) that are used to manage the resource.
func codestarconnections_ListTagsForResource(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_codestarconnectionsResourceArn) > 0 {
		input.ResourceArn = aws.String(_codestarconnectionsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds to or modifies the tags of the given resource. Tags are metadata that can
// be used to manage a resource.
func codestarconnections_TagResource(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_codestarconnectionsResourceArn) > 0 {
		input.ResourceArn = aws.String(_codestarconnectionsResourceArn)
	}
	if len(_codestarconnectionsTags) > 0 {
		if err := assignInputField(input, "Tags", _codestarconnectionsTags); err != nil {
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

// Removes tags from an Amazon Web Services resource.
func codestarconnections_UntagResource(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_codestarconnectionsResourceArn) > 0 {
		input.ResourceArn = aws.String(_codestarconnectionsResourceArn)
	}
	if len(_codestarconnectionsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _codestarconnectionsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a specified host with the provided configurations.
func codestarconnections_UpdateHost(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.UpdateHostInput{
		// HostArn: *string, // Required
	}

	if len(_codestarconnectionsHostArn) > 0 {
		input.HostArn = aws.String(_codestarconnectionsHostArn)
	}
	if len(_codestarconnectionsProviderEndpoint) > 0 {
		input.ProviderEndpoint = aws.String(_codestarconnectionsProviderEndpoint)
	}
	if len(_codestarconnectionsVpcConfiguration) > 0 {
		if err := assignInputField(input, "VpcConfiguration", _codestarconnectionsVpcConfiguration); err != nil {
			log.Errorf("invalid --vpc-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateHost(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the association between your connection and a specified external Git
// repository. A repository link allows Git sync to monitor and sync changes to
// files in a specified Git repository.
func codestarconnections_UpdateRepositoryLink(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.UpdateRepositoryLinkInput{
		// RepositoryLinkId: *string, // Required
	}

	if len(_codestarconnectionsRepositoryLinkId) > 0 {
		input.RepositoryLinkId = aws.String(_codestarconnectionsRepositoryLinkId)
	}
	if len(_codestarconnectionsConnectionArn) > 0 {
		input.ConnectionArn = aws.String(_codestarconnectionsConnectionArn)
	}
	if len(_codestarconnectionsEncryptionKeyArn) > 0 {
		input.EncryptionKeyArn = aws.String(_codestarconnectionsEncryptionKeyArn)
	}

	if resp, err := client.UpdateRepositoryLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows you to update the status of a sync blocker, resolving the blocker and
// allowing syncing to continue.
func codestarconnections_UpdateSyncBlocker(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.UpdateSyncBlockerInput{
		// Id: *string, // Required
		// ResolvedReason: *string, // Required
		// ResourceName: *string, // Required
		// SyncType: types.SyncConfigurationType, // Required
	}

	if len(_codestarconnectionsId) > 0 {
		input.Id = aws.String(_codestarconnectionsId)
	}
	if len(_codestarconnectionsResolvedReason) > 0 {
		input.ResolvedReason = aws.String(_codestarconnectionsResolvedReason)
	}
	if len(_codestarconnectionsResourceName) > 0 {
		input.ResourceName = aws.String(_codestarconnectionsResourceName)
	}
	if len(_codestarconnectionsSyncType) > 0 {
		if err := assignInputField(input, "SyncType", _codestarconnectionsSyncType); err != nil {
			log.Errorf("invalid --sync-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSyncBlocker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the sync configuration for your connection and a specified external Git
// repository.
func codestarconnections_UpdateSyncConfiguration(cfg aws.Config, client *codestarconnections.Client) {
	input := &codestarconnections.UpdateSyncConfigurationInput{
		// ResourceName: *string, // Required
		// SyncType: types.SyncConfigurationType, // Required
	}

	if len(_codestarconnectionsResourceName) > 0 {
		input.ResourceName = aws.String(_codestarconnectionsResourceName)
	}
	if len(_codestarconnectionsSyncType) > 0 {
		if err := assignInputField(input, "SyncType", _codestarconnectionsSyncType); err != nil {
			log.Errorf("invalid --sync-type: %s", err.Error())
			return
		}
	}
	if len(_codestarconnectionsBranch) > 0 {
		input.Branch = aws.String(_codestarconnectionsBranch)
	}
	if len(_codestarconnectionsConfigFile) > 0 {
		input.ConfigFile = aws.String(_codestarconnectionsConfigFile)
	}
	if len(_codestarconnectionsPublishDeploymentStatus) > 0 {
		if err := assignInputField(input, "PublishDeploymentStatus", _codestarconnectionsPublishDeploymentStatus); err != nil {
			log.Errorf("invalid --publish-deployment-status: %s", err.Error())
			return
		}
	}
	if len(_codestarconnectionsRepositoryLinkId) > 0 {
		input.RepositoryLinkId = aws.String(_codestarconnectionsRepositoryLinkId)
	}
	if len(_codestarconnectionsRoleArn) > 0 {
		input.RoleArn = aws.String(_codestarconnectionsRoleArn)
	}
	if len(_codestarconnectionsTriggerResourceUpdateOn) > 0 {
		if err := assignInputField(input, "TriggerResourceUpdateOn", _codestarconnectionsTriggerResourceUpdateOn); err != nil {
			log.Errorf("invalid --trigger-resource-update-on: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSyncConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_codestarconnectionsCmd)
	_codestarconnectionsCmd.Flags().SortFlags = false

	_codestarconnectionsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_codestarconnectionsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_codestarconnectionsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsBranch, "branch", "", "", "Branch")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsConfigFile, "config-file", "", "", "Config File")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsConnectionArn, "connection-arn", "", "", "Connection ARN")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsConnectionName, "connection-name", "", "", "Connection Name")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsEncryptionKeyArn, "encryption-key-arn", "", "", "Encryption Key ARN")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsHostArn, "host-arn", "", "", "Host ARN")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsHostArnFilter, "host-arn-filter", "", "", "Host ARN Filter")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsId, "id", "", "", "ID")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsMaxResults, "max-results", "", "", "Max Results")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsName, "name", "", "", "Name")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsNextToken, "next-token", "", "", "Next Token")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsOwnerId, "owner-id", "", "", "Owner ID")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsProviderEndpoint, "provider-endpoint", "", "", "Provider Endpoint")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsProviderType, "provider-type", "", "", "Provider Type")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsProviderTypeFilter, "provider-type-filter", "", "", "Provider Type Filter")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsPublishDeploymentStatus, "publish-deployment-status", "", "", "Publish Deployment Status")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsRepositoryLinkId, "repository-link-id", "", "", "Repository Link ID")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsRepositoryName, "repository-name", "", "", "Repository Name")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsResolvedReason, "resolved-reason", "", "", "Resolved Reason")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsResourceArn, "resource-arn", "", "", "Resource ARN")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsResourceName, "resource-name", "", "", "Resource Name")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsRoleArn, "role-arn", "", "", "Role ARN")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsSyncType, "sync-type", "", "", "Sync Type")
	_codestarconnectionsCmd.Flags().StringSliceVarP(&_codestarconnectionsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsTags, "tags", "", "", "Tags")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsTriggerResourceUpdateOn, "trigger-resource-update-on", "", "", "Trigger Resource Update On")
	_codestarconnectionsCmd.Flags().StringVarP(&_codestarconnectionsVpcConfiguration, "vpc-configuration", "", "", "VPC Configuration")

	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsCreateConnection, "create-connection", "", false, "Create Connection")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsCreateHost, "create-host", "", false, "Create Host")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsCreateRepositoryLink, "create-repository-link", "", false, "Create Repository Link")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsCreateSyncConfiguration, "create-sync-configuration", "", false, "Create Sync Configuration")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsDeleteConnection, "delete-connection", "", false, "Delete Connection")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsDeleteHost, "delete-host", "", false, "Delete Host")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsDeleteRepositoryLink, "delete-repository-link", "", false, "Delete Repository Link")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsDeleteSyncConfiguration, "delete-sync-configuration", "", false, "Delete Sync Configuration")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsGetConnection, "get-connection", "", false, "Get Connection")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsGetHost, "get-host", "", false, "Get Host")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsGetRepositoryLink, "get-repository-link", "", false, "Get Repository Link")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsGetRepositorySyncStatus, "get-repository-sync-status", "", false, "Get Repository Sync Status")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsGetResourceSyncStatus, "get-resource-sync-status", "", false, "Get Resource Sync Status")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsGetSyncBlockerSummary, "get-sync-blocker-summary", "", false, "Get Sync Blocker Summary")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsGetSyncConfiguration, "get-sync-configuration", "", false, "Get Sync Configuration")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsListConnections, "list-connections", "", false, "List Connections")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsListHosts, "list-hosts", "", false, "List Hosts")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsListRepositoryLinks, "list-repository-links", "", false, "List Repository Links")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsListRepositorySyncDefinitions, "list-repository-sync-definitions", "", false, "List Repository Sync Definitions")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsListSyncConfigurations, "list-sync-configurations", "", false, "List Sync Configurations")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsTagResource, "tag-resource", "", false, "Tag Resource")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsUntagResource, "untag-resource", "", false, "Untag Resource")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsUpdateHost, "update-host", "", false, "Update Host")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsUpdateRepositoryLink, "update-repository-link", "", false, "Update Repository Link")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsUpdateSyncBlocker, "update-sync-blocker", "", false, "Update Sync Blocker")
	_codestarconnectionsCmd.Flags().BoolVarP(&_codestarconnectionsUpdateSyncConfiguration, "update-sync-configuration", "", false, "Update Sync Configuration")

}
