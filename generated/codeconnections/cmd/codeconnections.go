package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codeconnections"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// codeconnectionsCmd represents the codeconnections command
var _codeconnectionsCmd = &cobra.Command{
	Use:   "codeconnections",
	Short: "AWS codeconnections CLI",
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
		client := codeconnections.NewFromConfig(cfg)
		if _codeconnectionsCreateConnection {
			codeconnections_CreateConnection(cfg, client)
			return
		}
		if _codeconnectionsCreateHost {
			codeconnections_CreateHost(cfg, client)
			return
		}
		if _codeconnectionsCreateRepositoryLink {
			codeconnections_CreateRepositoryLink(cfg, client)
			return
		}
		if _codeconnectionsCreateSyncConfiguration {
			codeconnections_CreateSyncConfiguration(cfg, client)
			return
		}
		if _codeconnectionsDeleteConnection {
			codeconnections_DeleteConnection(cfg, client)
			return
		}
		if _codeconnectionsDeleteHost {
			codeconnections_DeleteHost(cfg, client)
			return
		}
		if _codeconnectionsDeleteRepositoryLink {
			codeconnections_DeleteRepositoryLink(cfg, client)
			return
		}
		if _codeconnectionsDeleteSyncConfiguration {
			codeconnections_DeleteSyncConfiguration(cfg, client)
			return
		}
		if _codeconnectionsGetConnection {
			codeconnections_GetConnection(cfg, client)
			return
		}
		if _codeconnectionsGetHost {
			codeconnections_GetHost(cfg, client)
			return
		}
		if _codeconnectionsGetRepositoryLink {
			codeconnections_GetRepositoryLink(cfg, client)
			return
		}
		if _codeconnectionsGetRepositorySyncStatus {
			codeconnections_GetRepositorySyncStatus(cfg, client)
			return
		}
		if _codeconnectionsGetResourceSyncStatus {
			codeconnections_GetResourceSyncStatus(cfg, client)
			return
		}
		if _codeconnectionsGetSyncBlockerSummary {
			codeconnections_GetSyncBlockerSummary(cfg, client)
			return
		}
		if _codeconnectionsGetSyncConfiguration {
			codeconnections_GetSyncConfiguration(cfg, client)
			return
		}
		if _codeconnectionsListConnections {
			codeconnections_ListConnections(cfg, client)
			return
		}
		if _codeconnectionsListHosts {
			codeconnections_ListHosts(cfg, client)
			return
		}
		if _codeconnectionsListRepositoryLinks {
			codeconnections_ListRepositoryLinks(cfg, client)
			return
		}
		if _codeconnectionsListRepositorySyncDefinitions {
			codeconnections_ListRepositorySyncDefinitions(cfg, client)
			return
		}
		if _codeconnectionsListSyncConfigurations {
			codeconnections_ListSyncConfigurations(cfg, client)
			return
		}
		if _codeconnectionsListTagsForResource {
			codeconnections_ListTagsForResource(cfg, client)
			return
		}
		if _codeconnectionsTagResource {
			codeconnections_TagResource(cfg, client)
			return
		}
		if _codeconnectionsUntagResource {
			codeconnections_UntagResource(cfg, client)
			return
		}
		if _codeconnectionsUpdateHost {
			codeconnections_UpdateHost(cfg, client)
			return
		}
		if _codeconnectionsUpdateRepositoryLink {
			codeconnections_UpdateRepositoryLink(cfg, client)
			return
		}
		if _codeconnectionsUpdateSyncBlocker {
			codeconnections_UpdateSyncBlocker(cfg, client)
			return
		}
		if _codeconnectionsUpdateSyncConfiguration {
			codeconnections_UpdateSyncConfiguration(cfg, client)
			return
		}

	},
}

var (
	_codeconnectionsCreateConnection              bool
	_codeconnectionsCreateHost                    bool
	_codeconnectionsCreateRepositoryLink          bool
	_codeconnectionsCreateSyncConfiguration       bool
	_codeconnectionsDeleteConnection              bool
	_codeconnectionsDeleteHost                    bool
	_codeconnectionsDeleteRepositoryLink          bool
	_codeconnectionsDeleteSyncConfiguration       bool
	_codeconnectionsGetConnection                 bool
	_codeconnectionsGetHost                       bool
	_codeconnectionsGetRepositoryLink             bool
	_codeconnectionsGetRepositorySyncStatus       bool
	_codeconnectionsGetResourceSyncStatus         bool
	_codeconnectionsGetSyncBlockerSummary         bool
	_codeconnectionsGetSyncConfiguration          bool
	_codeconnectionsListConnections               bool
	_codeconnectionsListHosts                     bool
	_codeconnectionsListRepositoryLinks           bool
	_codeconnectionsListRepositorySyncDefinitions bool
	_codeconnectionsListSyncConfigurations        bool
	_codeconnectionsListTagsForResource           bool
	_codeconnectionsTagResource                   bool
	_codeconnectionsUntagResource                 bool
	_codeconnectionsUpdateHost                    bool
	_codeconnectionsUpdateRepositoryLink          bool
	_codeconnectionsUpdateSyncBlocker             bool
	_codeconnectionsUpdateSyncConfiguration       bool

	_codeconnectionsBranch                  string
	_codeconnectionsConfigFile              string
	_codeconnectionsConnectionArn           string
	_codeconnectionsConnectionName          string
	_codeconnectionsEncryptionKeyArn        string
	_codeconnectionsHostArn                 string
	_codeconnectionsHostArnFilter           string
	_codeconnectionsId                      string
	_codeconnectionsMaxResults              string
	_codeconnectionsName                    string
	_codeconnectionsNextToken               string
	_codeconnectionsOwnerId                 string
	_codeconnectionsProviderEndpoint        string
	_codeconnectionsProviderType            string
	_codeconnectionsProviderTypeFilter      string
	_codeconnectionsPublishDeploymentStatus string
	_codeconnectionsPullRequestComment      string
	_codeconnectionsRepositoryLinkId        string
	_codeconnectionsRepositoryName          string
	_codeconnectionsResolvedReason          string
	_codeconnectionsResourceArn             string
	_codeconnectionsResourceName            string
	_codeconnectionsRoleArn                 string
	_codeconnectionsSyncType                string
	_codeconnectionsTagKeys                 []string
	_codeconnectionsTags                    string
	_codeconnectionsTriggerResourceUpdateOn string
	_codeconnectionsVpcConfiguration        string
)

// Creates a connection that can then be given to other Amazon Web Services
// services like CodePipeline so that it can access third-party code repositories.
// The connection is in pending status until the third-party connection handshake
// is completed from the console.
func codeconnections_CreateConnection(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.CreateConnectionInput{
		// ConnectionName: *string, // Required
	}

	if len(_codeconnectionsConnectionName) > 0 {
		input.ConnectionName = aws.String(_codeconnectionsConnectionName)
	}
	if len(_codeconnectionsHostArn) > 0 {
		input.HostArn = aws.String(_codeconnectionsHostArn)
	}
	if len(_codeconnectionsProviderType) > 0 {
		if err := assignInputField(input, "ProviderType", _codeconnectionsProviderType); err != nil {
			log.Errorf("invalid --provider-type: %s", err.Error())
			return
		}
	}
	if len(_codeconnectionsTags) > 0 {
		if err := assignInputField(input, "Tags", _codeconnectionsTags); err != nil {
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
func codeconnections_CreateHost(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.CreateHostInput{
		// Name: *string, // Required
		// ProviderEndpoint: *string, // Required
		// ProviderType: types.ProviderType, // Required
	}

	if len(_codeconnectionsName) > 0 {
		input.Name = aws.String(_codeconnectionsName)
	}
	if len(_codeconnectionsProviderEndpoint) > 0 {
		input.ProviderEndpoint = aws.String(_codeconnectionsProviderEndpoint)
	}
	if len(_codeconnectionsProviderType) > 0 {
		if err := assignInputField(input, "ProviderType", _codeconnectionsProviderType); err != nil {
			log.Errorf("invalid --provider-type: %s", err.Error())
			return
		}
	}
	if len(_codeconnectionsTags) > 0 {
		if err := assignInputField(input, "Tags", _codeconnectionsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_codeconnectionsVpcConfiguration) > 0 {
		if err := assignInputField(input, "VpcConfiguration", _codeconnectionsVpcConfiguration); err != nil {
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
func codeconnections_CreateRepositoryLink(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.CreateRepositoryLinkInput{
		// ConnectionArn: *string, // Required
		// OwnerId: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_codeconnectionsConnectionArn) > 0 {
		input.ConnectionArn = aws.String(_codeconnectionsConnectionArn)
	}
	if len(_codeconnectionsOwnerId) > 0 {
		input.OwnerId = aws.String(_codeconnectionsOwnerId)
	}
	if len(_codeconnectionsRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codeconnectionsRepositoryName)
	}
	if len(_codeconnectionsEncryptionKeyArn) > 0 {
		input.EncryptionKeyArn = aws.String(_codeconnectionsEncryptionKeyArn)
	}
	if len(_codeconnectionsTags) > 0 {
		if err := assignInputField(input, "Tags", _codeconnectionsTags); err != nil {
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
func codeconnections_CreateSyncConfiguration(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.CreateSyncConfigurationInput{
		// Branch: *string, // Required
		// ConfigFile: *string, // Required
		// RepositoryLinkId: *string, // Required
		// ResourceName: *string, // Required
		// RoleArn: *string, // Required
		// SyncType: types.SyncConfigurationType, // Required
	}

	if len(_codeconnectionsBranch) > 0 {
		input.Branch = aws.String(_codeconnectionsBranch)
	}
	if len(_codeconnectionsConfigFile) > 0 {
		input.ConfigFile = aws.String(_codeconnectionsConfigFile)
	}
	if len(_codeconnectionsRepositoryLinkId) > 0 {
		input.RepositoryLinkId = aws.String(_codeconnectionsRepositoryLinkId)
	}
	if len(_codeconnectionsResourceName) > 0 {
		input.ResourceName = aws.String(_codeconnectionsResourceName)
	}
	if len(_codeconnectionsRoleArn) > 0 {
		input.RoleArn = aws.String(_codeconnectionsRoleArn)
	}
	if len(_codeconnectionsSyncType) > 0 {
		if err := assignInputField(input, "SyncType", _codeconnectionsSyncType); err != nil {
			log.Errorf("invalid --sync-type: %s", err.Error())
			return
		}
	}
	if len(_codeconnectionsPublishDeploymentStatus) > 0 {
		if err := assignInputField(input, "PublishDeploymentStatus", _codeconnectionsPublishDeploymentStatus); err != nil {
			log.Errorf("invalid --publish-deployment-status: %s", err.Error())
			return
		}
	}
	if len(_codeconnectionsPullRequestComment) > 0 {
		if err := assignInputField(input, "PullRequestComment", _codeconnectionsPullRequestComment); err != nil {
			log.Errorf("invalid --pull-request-comment: %s", err.Error())
			return
		}
	}
	if len(_codeconnectionsTriggerResourceUpdateOn) > 0 {
		if err := assignInputField(input, "TriggerResourceUpdateOn", _codeconnectionsTriggerResourceUpdateOn); err != nil {
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
func codeconnections_DeleteConnection(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.DeleteConnectionInput{
		// ConnectionArn: *string, // Required
	}

	if len(_codeconnectionsConnectionArn) > 0 {
		input.ConnectionArn = aws.String(_codeconnectionsConnectionArn)
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
func codeconnections_DeleteHost(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.DeleteHostInput{
		// HostArn: *string, // Required
	}

	if len(_codeconnectionsHostArn) > 0 {
		input.HostArn = aws.String(_codeconnectionsHostArn)
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
func codeconnections_DeleteRepositoryLink(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.DeleteRepositoryLinkInput{
		// RepositoryLinkId: *string, // Required
	}

	if len(_codeconnectionsRepositoryLinkId) > 0 {
		input.RepositoryLinkId = aws.String(_codeconnectionsRepositoryLinkId)
	}

	if resp, err := client.DeleteRepositoryLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the sync configuration for a specified repository and connection.
func codeconnections_DeleteSyncConfiguration(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.DeleteSyncConfigurationInput{
		// ResourceName: *string, // Required
		// SyncType: types.SyncConfigurationType, // Required
	}

	if len(_codeconnectionsResourceName) > 0 {
		input.ResourceName = aws.String(_codeconnectionsResourceName)
	}
	if len(_codeconnectionsSyncType) > 0 {
		if err := assignInputField(input, "SyncType", _codeconnectionsSyncType); err != nil {
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
func codeconnections_GetConnection(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.GetConnectionInput{
		// ConnectionArn: *string, // Required
	}

	if len(_codeconnectionsConnectionArn) > 0 {
		input.ConnectionArn = aws.String(_codeconnectionsConnectionArn)
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
func codeconnections_GetHost(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.GetHostInput{
		// HostArn: *string, // Required
	}

	if len(_codeconnectionsHostArn) > 0 {
		input.HostArn = aws.String(_codeconnectionsHostArn)
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
func codeconnections_GetRepositoryLink(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.GetRepositoryLinkInput{
		// RepositoryLinkId: *string, // Required
	}

	if len(_codeconnectionsRepositoryLinkId) > 0 {
		input.RepositoryLinkId = aws.String(_codeconnectionsRepositoryLinkId)
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
func codeconnections_GetRepositorySyncStatus(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.GetRepositorySyncStatusInput{
		// Branch: *string, // Required
		// RepositoryLinkId: *string, // Required
		// SyncType: types.SyncConfigurationType, // Required
	}

	if len(_codeconnectionsBranch) > 0 {
		input.Branch = aws.String(_codeconnectionsBranch)
	}
	if len(_codeconnectionsRepositoryLinkId) > 0 {
		input.RepositoryLinkId = aws.String(_codeconnectionsRepositoryLinkId)
	}
	if len(_codeconnectionsSyncType) > 0 {
		if err := assignInputField(input, "SyncType", _codeconnectionsSyncType); err != nil {
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
func codeconnections_GetResourceSyncStatus(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.GetResourceSyncStatusInput{
		// ResourceName: *string, // Required
		// SyncType: types.SyncConfigurationType, // Required
	}

	if len(_codeconnectionsResourceName) > 0 {
		input.ResourceName = aws.String(_codeconnectionsResourceName)
	}
	if len(_codeconnectionsSyncType) > 0 {
		if err := assignInputField(input, "SyncType", _codeconnectionsSyncType); err != nil {
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
func codeconnections_GetSyncBlockerSummary(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.GetSyncBlockerSummaryInput{
		// ResourceName: *string, // Required
		// SyncType: types.SyncConfigurationType, // Required
	}

	if len(_codeconnectionsResourceName) > 0 {
		input.ResourceName = aws.String(_codeconnectionsResourceName)
	}
	if len(_codeconnectionsSyncType) > 0 {
		if err := assignInputField(input, "SyncType", _codeconnectionsSyncType); err != nil {
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
func codeconnections_GetSyncConfiguration(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.GetSyncConfigurationInput{
		// ResourceName: *string, // Required
		// SyncType: types.SyncConfigurationType, // Required
	}

	if len(_codeconnectionsResourceName) > 0 {
		input.ResourceName = aws.String(_codeconnectionsResourceName)
	}
	if len(_codeconnectionsSyncType) > 0 {
		if err := assignInputField(input, "SyncType", _codeconnectionsSyncType); err != nil {
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
func codeconnections_ListConnections(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.ListConnectionsInput{}

	if len(_codeconnectionsHostArnFilter) > 0 {
		input.HostArnFilter = aws.String(_codeconnectionsHostArnFilter)
	}
	if len(_codeconnectionsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codeconnectionsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codeconnectionsNextToken) > 0 {
		input.NextToken = aws.String(_codeconnectionsNextToken)
	}
	if len(_codeconnectionsProviderTypeFilter) > 0 {
		if err := assignInputField(input, "ProviderTypeFilter", _codeconnectionsProviderTypeFilter); err != nil {
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

	var results []*codeconnections.ListConnectionsOutput
	p := codeconnections.NewListConnectionsPaginator(client, input)
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
func codeconnections_ListHosts(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.ListHostsInput{}

	if len(_codeconnectionsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codeconnectionsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codeconnectionsNextToken) > 0 {
		input.NextToken = aws.String(_codeconnectionsNextToken)
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

	var results []*codeconnections.ListHostsOutput
	p := codeconnections.NewListHostsPaginator(client, input)
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
func codeconnections_ListRepositoryLinks(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.ListRepositoryLinksInput{}

	if len(_codeconnectionsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codeconnectionsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codeconnectionsNextToken) > 0 {
		input.NextToken = aws.String(_codeconnectionsNextToken)
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

	var results []*codeconnections.ListRepositoryLinksOutput
	p := codeconnections.NewListRepositoryLinksPaginator(client, input)
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
func codeconnections_ListRepositorySyncDefinitions(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.ListRepositorySyncDefinitionsInput{
		// RepositoryLinkId: *string, // Required
		// SyncType: types.SyncConfigurationType, // Required
	}

	if len(_codeconnectionsRepositoryLinkId) > 0 {
		input.RepositoryLinkId = aws.String(_codeconnectionsRepositoryLinkId)
	}
	if len(_codeconnectionsSyncType) > 0 {
		if err := assignInputField(input, "SyncType", _codeconnectionsSyncType); err != nil {
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
func codeconnections_ListSyncConfigurations(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.ListSyncConfigurationsInput{
		// RepositoryLinkId: *string, // Required
		// SyncType: types.SyncConfigurationType, // Required
	}

	if len(_codeconnectionsRepositoryLinkId) > 0 {
		input.RepositoryLinkId = aws.String(_codeconnectionsRepositoryLinkId)
	}
	if len(_codeconnectionsSyncType) > 0 {
		if err := assignInputField(input, "SyncType", _codeconnectionsSyncType); err != nil {
			log.Errorf("invalid --sync-type: %s", err.Error())
			return
		}
	}
	if len(_codeconnectionsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codeconnectionsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codeconnectionsNextToken) > 0 {
		input.NextToken = aws.String(_codeconnectionsNextToken)
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

	var results []*codeconnections.ListSyncConfigurationsOutput
	p := codeconnections.NewListSyncConfigurationsPaginator(client, input)
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
func codeconnections_ListTagsForResource(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_codeconnectionsResourceArn) > 0 {
		input.ResourceArn = aws.String(_codeconnectionsResourceArn)
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
func codeconnections_TagResource(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_codeconnectionsResourceArn) > 0 {
		input.ResourceArn = aws.String(_codeconnectionsResourceArn)
	}
	if len(_codeconnectionsTags) > 0 {
		if err := assignInputField(input, "Tags", _codeconnectionsTags); err != nil {
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
func codeconnections_UntagResource(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_codeconnectionsResourceArn) > 0 {
		input.ResourceArn = aws.String(_codeconnectionsResourceArn)
	}
	if len(_codeconnectionsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _codeconnectionsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a specified host with the provided configurations.
func codeconnections_UpdateHost(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.UpdateHostInput{
		// HostArn: *string, // Required
	}

	if len(_codeconnectionsHostArn) > 0 {
		input.HostArn = aws.String(_codeconnectionsHostArn)
	}
	if len(_codeconnectionsProviderEndpoint) > 0 {
		input.ProviderEndpoint = aws.String(_codeconnectionsProviderEndpoint)
	}
	if len(_codeconnectionsVpcConfiguration) > 0 {
		if err := assignInputField(input, "VpcConfiguration", _codeconnectionsVpcConfiguration); err != nil {
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
func codeconnections_UpdateRepositoryLink(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.UpdateRepositoryLinkInput{
		// RepositoryLinkId: *string, // Required
	}

	if len(_codeconnectionsRepositoryLinkId) > 0 {
		input.RepositoryLinkId = aws.String(_codeconnectionsRepositoryLinkId)
	}
	if len(_codeconnectionsConnectionArn) > 0 {
		input.ConnectionArn = aws.String(_codeconnectionsConnectionArn)
	}
	if len(_codeconnectionsEncryptionKeyArn) > 0 {
		input.EncryptionKeyArn = aws.String(_codeconnectionsEncryptionKeyArn)
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
func codeconnections_UpdateSyncBlocker(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.UpdateSyncBlockerInput{
		// Id: *string, // Required
		// ResolvedReason: *string, // Required
		// ResourceName: *string, // Required
		// SyncType: types.SyncConfigurationType, // Required
	}

	if len(_codeconnectionsId) > 0 {
		input.Id = aws.String(_codeconnectionsId)
	}
	if len(_codeconnectionsResolvedReason) > 0 {
		input.ResolvedReason = aws.String(_codeconnectionsResolvedReason)
	}
	if len(_codeconnectionsResourceName) > 0 {
		input.ResourceName = aws.String(_codeconnectionsResourceName)
	}
	if len(_codeconnectionsSyncType) > 0 {
		if err := assignInputField(input, "SyncType", _codeconnectionsSyncType); err != nil {
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
func codeconnections_UpdateSyncConfiguration(cfg aws.Config, client *codeconnections.Client) {
	input := &codeconnections.UpdateSyncConfigurationInput{
		// ResourceName: *string, // Required
		// SyncType: types.SyncConfigurationType, // Required
	}

	if len(_codeconnectionsResourceName) > 0 {
		input.ResourceName = aws.String(_codeconnectionsResourceName)
	}
	if len(_codeconnectionsSyncType) > 0 {
		if err := assignInputField(input, "SyncType", _codeconnectionsSyncType); err != nil {
			log.Errorf("invalid --sync-type: %s", err.Error())
			return
		}
	}
	if len(_codeconnectionsBranch) > 0 {
		input.Branch = aws.String(_codeconnectionsBranch)
	}
	if len(_codeconnectionsConfigFile) > 0 {
		input.ConfigFile = aws.String(_codeconnectionsConfigFile)
	}
	if len(_codeconnectionsPublishDeploymentStatus) > 0 {
		if err := assignInputField(input, "PublishDeploymentStatus", _codeconnectionsPublishDeploymentStatus); err != nil {
			log.Errorf("invalid --publish-deployment-status: %s", err.Error())
			return
		}
	}
	if len(_codeconnectionsPullRequestComment) > 0 {
		if err := assignInputField(input, "PullRequestComment", _codeconnectionsPullRequestComment); err != nil {
			log.Errorf("invalid --pull-request-comment: %s", err.Error())
			return
		}
	}
	if len(_codeconnectionsRepositoryLinkId) > 0 {
		input.RepositoryLinkId = aws.String(_codeconnectionsRepositoryLinkId)
	}
	if len(_codeconnectionsRoleArn) > 0 {
		input.RoleArn = aws.String(_codeconnectionsRoleArn)
	}
	if len(_codeconnectionsTriggerResourceUpdateOn) > 0 {
		if err := assignInputField(input, "TriggerResourceUpdateOn", _codeconnectionsTriggerResourceUpdateOn); err != nil {
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
	_rootCmd.AddCommand(_codeconnectionsCmd)
	_codeconnectionsCmd.Flags().SortFlags = false

	_codeconnectionsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_codeconnectionsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_codeconnectionsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsBranch, "branch", "", "", "Branch")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsConfigFile, "config-file", "", "", "Config File")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsConnectionArn, "connection-arn", "", "", "Connection ARN")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsConnectionName, "connection-name", "", "", "Connection Name")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsEncryptionKeyArn, "encryption-key-arn", "", "", "Encryption Key ARN")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsHostArn, "host-arn", "", "", "Host ARN")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsHostArnFilter, "host-arn-filter", "", "", "Host ARN Filter")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsId, "id", "", "", "ID")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsMaxResults, "max-results", "", "", "Max Results")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsName, "name", "", "", "Name")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsNextToken, "next-token", "", "", "Next Token")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsOwnerId, "owner-id", "", "", "Owner ID")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsProviderEndpoint, "provider-endpoint", "", "", "Provider Endpoint")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsProviderType, "provider-type", "", "", "Provider Type")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsProviderTypeFilter, "provider-type-filter", "", "", "Provider Type Filter")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsPublishDeploymentStatus, "publish-deployment-status", "", "", "Publish Deployment Status")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsPullRequestComment, "pull-request-comment", "", "", "Pull Request Comment")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsRepositoryLinkId, "repository-link-id", "", "", "Repository Link ID")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsRepositoryName, "repository-name", "", "", "Repository Name")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsResolvedReason, "resolved-reason", "", "", "Resolved Reason")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsResourceArn, "resource-arn", "", "", "Resource ARN")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsResourceName, "resource-name", "", "", "Resource Name")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsRoleArn, "role-arn", "", "", "Role ARN")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsSyncType, "sync-type", "", "", "Sync Type")
	_codeconnectionsCmd.Flags().StringSliceVarP(&_codeconnectionsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsTags, "tags", "", "", "Tags")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsTriggerResourceUpdateOn, "trigger-resource-update-on", "", "", "Trigger Resource Update On")
	_codeconnectionsCmd.Flags().StringVarP(&_codeconnectionsVpcConfiguration, "vpc-configuration", "", "", "VPC Configuration")

	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsCreateConnection, "create-connection", "", false, "Create Connection")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsCreateHost, "create-host", "", false, "Create Host")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsCreateRepositoryLink, "create-repository-link", "", false, "Create Repository Link")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsCreateSyncConfiguration, "create-sync-configuration", "", false, "Create Sync Configuration")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsDeleteConnection, "delete-connection", "", false, "Delete Connection")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsDeleteHost, "delete-host", "", false, "Delete Host")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsDeleteRepositoryLink, "delete-repository-link", "", false, "Delete Repository Link")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsDeleteSyncConfiguration, "delete-sync-configuration", "", false, "Delete Sync Configuration")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsGetConnection, "get-connection", "", false, "Get Connection")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsGetHost, "get-host", "", false, "Get Host")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsGetRepositoryLink, "get-repository-link", "", false, "Get Repository Link")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsGetRepositorySyncStatus, "get-repository-sync-status", "", false, "Get Repository Sync Status")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsGetResourceSyncStatus, "get-resource-sync-status", "", false, "Get Resource Sync Status")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsGetSyncBlockerSummary, "get-sync-blocker-summary", "", false, "Get Sync Blocker Summary")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsGetSyncConfiguration, "get-sync-configuration", "", false, "Get Sync Configuration")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsListConnections, "list-connections", "", false, "List Connections")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsListHosts, "list-hosts", "", false, "List Hosts")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsListRepositoryLinks, "list-repository-links", "", false, "List Repository Links")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsListRepositorySyncDefinitions, "list-repository-sync-definitions", "", false, "List Repository Sync Definitions")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsListSyncConfigurations, "list-sync-configurations", "", false, "List Sync Configurations")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsTagResource, "tag-resource", "", false, "Tag Resource")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsUntagResource, "untag-resource", "", false, "Untag Resource")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsUpdateHost, "update-host", "", false, "Update Host")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsUpdateRepositoryLink, "update-repository-link", "", false, "Update Repository Link")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsUpdateSyncBlocker, "update-sync-blocker", "", false, "Update Sync Blocker")
	_codeconnectionsCmd.Flags().BoolVarP(&_codeconnectionsUpdateSyncConfiguration, "update-sync-configuration", "", false, "Update Sync Configuration")

}
