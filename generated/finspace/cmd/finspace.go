package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/finspace"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// finspaceCmd represents the finspace command
var _finspaceCmd = &cobra.Command{
	Use:   "finspace",
	Short: "AWS finspace CLI",
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
		client := finspace.NewFromConfig(cfg)
		if _finspaceCreateEnvironment {
			finspace_CreateEnvironment(cfg, client)
			return
		}
		if _finspaceCreateKxChangeset {
			finspace_CreateKxChangeset(cfg, client)
			return
		}
		if _finspaceCreateKxCluster {
			finspace_CreateKxCluster(cfg, client)
			return
		}
		if _finspaceCreateKxDatabase {
			finspace_CreateKxDatabase(cfg, client)
			return
		}
		if _finspaceCreateKxDataview {
			finspace_CreateKxDataview(cfg, client)
			return
		}
		if _finspaceCreateKxEnvironment {
			finspace_CreateKxEnvironment(cfg, client)
			return
		}
		if _finspaceCreateKxScalingGroup {
			finspace_CreateKxScalingGroup(cfg, client)
			return
		}
		if _finspaceCreateKxUser {
			finspace_CreateKxUser(cfg, client)
			return
		}
		if _finspaceCreateKxVolume {
			finspace_CreateKxVolume(cfg, client)
			return
		}
		if _finspaceDeleteEnvironment {
			finspace_DeleteEnvironment(cfg, client)
			return
		}
		if _finspaceDeleteKxCluster {
			finspace_DeleteKxCluster(cfg, client)
			return
		}
		if _finspaceDeleteKxClusterNode {
			finspace_DeleteKxClusterNode(cfg, client)
			return
		}
		if _finspaceDeleteKxDatabase {
			finspace_DeleteKxDatabase(cfg, client)
			return
		}
		if _finspaceDeleteKxDataview {
			finspace_DeleteKxDataview(cfg, client)
			return
		}
		if _finspaceDeleteKxEnvironment {
			finspace_DeleteKxEnvironment(cfg, client)
			return
		}
		if _finspaceDeleteKxScalingGroup {
			finspace_DeleteKxScalingGroup(cfg, client)
			return
		}
		if _finspaceDeleteKxUser {
			finspace_DeleteKxUser(cfg, client)
			return
		}
		if _finspaceDeleteKxVolume {
			finspace_DeleteKxVolume(cfg, client)
			return
		}
		if _finspaceGetEnvironment {
			finspace_GetEnvironment(cfg, client)
			return
		}
		if _finspaceGetKxChangeset {
			finspace_GetKxChangeset(cfg, client)
			return
		}
		if _finspaceGetKxCluster {
			finspace_GetKxCluster(cfg, client)
			return
		}
		if _finspaceGetKxConnectionString {
			finspace_GetKxConnectionString(cfg, client)
			return
		}
		if _finspaceGetKxDatabase {
			finspace_GetKxDatabase(cfg, client)
			return
		}
		if _finspaceGetKxDataview {
			finspace_GetKxDataview(cfg, client)
			return
		}
		if _finspaceGetKxEnvironment {
			finspace_GetKxEnvironment(cfg, client)
			return
		}
		if _finspaceGetKxScalingGroup {
			finspace_GetKxScalingGroup(cfg, client)
			return
		}
		if _finspaceGetKxUser {
			finspace_GetKxUser(cfg, client)
			return
		}
		if _finspaceGetKxVolume {
			finspace_GetKxVolume(cfg, client)
			return
		}
		if _finspaceListEnvironments {
			finspace_ListEnvironments(cfg, client)
			return
		}
		if _finspaceListKxChangesets {
			finspace_ListKxChangesets(cfg, client)
			return
		}
		if _finspaceListKxClusterNodes {
			finspace_ListKxClusterNodes(cfg, client)
			return
		}
		if _finspaceListKxClusters {
			finspace_ListKxClusters(cfg, client)
			return
		}
		if _finspaceListKxDatabases {
			finspace_ListKxDatabases(cfg, client)
			return
		}
		if _finspaceListKxDataviews {
			finspace_ListKxDataviews(cfg, client)
			return
		}
		if _finspaceListKxEnvironments {
			finspace_ListKxEnvironments(cfg, client)
			return
		}
		if _finspaceListKxScalingGroups {
			finspace_ListKxScalingGroups(cfg, client)
			return
		}
		if _finspaceListKxUsers {
			finspace_ListKxUsers(cfg, client)
			return
		}
		if _finspaceListKxVolumes {
			finspace_ListKxVolumes(cfg, client)
			return
		}
		if _finspaceListTagsForResource {
			finspace_ListTagsForResource(cfg, client)
			return
		}
		if _finspaceTagResource {
			finspace_TagResource(cfg, client)
			return
		}
		if _finspaceUntagResource {
			finspace_UntagResource(cfg, client)
			return
		}
		if _finspaceUpdateEnvironment {
			finspace_UpdateEnvironment(cfg, client)
			return
		}
		if _finspaceUpdateKxClusterCodeConfiguration {
			finspace_UpdateKxClusterCodeConfiguration(cfg, client)
			return
		}
		if _finspaceUpdateKxClusterDatabases {
			finspace_UpdateKxClusterDatabases(cfg, client)
			return
		}
		if _finspaceUpdateKxDatabase {
			finspace_UpdateKxDatabase(cfg, client)
			return
		}
		if _finspaceUpdateKxDataview {
			finspace_UpdateKxDataview(cfg, client)
			return
		}
		if _finspaceUpdateKxEnvironment {
			finspace_UpdateKxEnvironment(cfg, client)
			return
		}
		if _finspaceUpdateKxEnvironmentNetwork {
			finspace_UpdateKxEnvironmentNetwork(cfg, client)
			return
		}
		if _finspaceUpdateKxUser {
			finspace_UpdateKxUser(cfg, client)
			return
		}
		if _finspaceUpdateKxVolume {
			finspace_UpdateKxVolume(cfg, client)
			return
		}

	},
}

var (
	_finspaceCreateEnvironment                bool
	_finspaceCreateKxChangeset                bool
	_finspaceCreateKxCluster                  bool
	_finspaceCreateKxDatabase                 bool
	_finspaceCreateKxDataview                 bool
	_finspaceCreateKxEnvironment              bool
	_finspaceCreateKxScalingGroup             bool
	_finspaceCreateKxUser                     bool
	_finspaceCreateKxVolume                   bool
	_finspaceDeleteEnvironment                bool
	_finspaceDeleteKxCluster                  bool
	_finspaceDeleteKxClusterNode              bool
	_finspaceDeleteKxDatabase                 bool
	_finspaceDeleteKxDataview                 bool
	_finspaceDeleteKxEnvironment              bool
	_finspaceDeleteKxScalingGroup             bool
	_finspaceDeleteKxUser                     bool
	_finspaceDeleteKxVolume                   bool
	_finspaceGetEnvironment                   bool
	_finspaceGetKxChangeset                   bool
	_finspaceGetKxCluster                     bool
	_finspaceGetKxConnectionString            bool
	_finspaceGetKxDatabase                    bool
	_finspaceGetKxDataview                    bool
	_finspaceGetKxEnvironment                 bool
	_finspaceGetKxScalingGroup                bool
	_finspaceGetKxUser                        bool
	_finspaceGetKxVolume                      bool
	_finspaceListEnvironments                 bool
	_finspaceListKxChangesets                 bool
	_finspaceListKxClusterNodes               bool
	_finspaceListKxClusters                   bool
	_finspaceListKxDatabases                  bool
	_finspaceListKxDataviews                  bool
	_finspaceListKxEnvironments               bool
	_finspaceListKxScalingGroups              bool
	_finspaceListKxUsers                      bool
	_finspaceListKxVolumes                    bool
	_finspaceListTagsForResource              bool
	_finspaceTagResource                      bool
	_finspaceUntagResource                    bool
	_finspaceUpdateEnvironment                bool
	_finspaceUpdateKxClusterCodeConfiguration bool
	_finspaceUpdateKxClusterDatabases         bool
	_finspaceUpdateKxDatabase                 bool
	_finspaceUpdateKxDataview                 bool
	_finspaceUpdateKxEnvironment              bool
	_finspaceUpdateKxEnvironmentNetwork       bool
	_finspaceUpdateKxUser                     bool
	_finspaceUpdateKxVolume                   bool

	_finspaceAutoScalingConfiguration     string
	_finspaceAutoUpdate                   string
	_finspaceAvailabilityZoneId           string
	_finspaceAvailabilityZoneIds          []string
	_finspaceAzMode                       string
	_finspaceCacheStorageConfigurations   string
	_finspaceCapacityConfiguration        string
	_finspaceChangeRequests               string
	_finspaceChangesetId                  string
	_finspaceClientToken                  string
	_finspaceClusterDescription           string
	_finspaceClusterName                  string
	_finspaceClusterType                  string
	_finspaceCode                         string
	_finspaceCommandLineArguments         string
	_finspaceCustomDNSConfiguration       string
	_finspaceDataBundles                  []string
	_finspaceDatabaseName                 string
	_finspaceDatabases                    string
	_finspaceDataviewName                 string
	_finspaceDeploymentConfiguration      string
	_finspaceDescription                  string
	_finspaceEnvironmentId                string
	_finspaceExecutionRole                string
	_finspaceFederationMode               string
	_finspaceFederationParameters         string
	_finspaceHostType                     string
	_finspaceIamRole                      string
	_finspaceInitializationScript         string
	_finspaceKmsKeyId                     string
	_finspaceMaxResults                   string
	_finspaceName                         string
	_finspaceNas1Configuration            string
	_finspaceNextToken                    string
	_finspaceNodeId                       string
	_finspaceReadWrite                    string
	_finspaceReleaseLabel                 string
	_finspaceResourceArn                  string
	_finspaceSavedownStorageConfiguration string
	_finspaceScalingGroupConfiguration    string
	_finspaceScalingGroupName             string
	_finspaceSegmentConfigurations        string
	_finspaceSuperuserParameters          string
	_finspaceTagKeys                      []string
	_finspaceTags                         string
	_finspaceTickerplantLogConfiguration  string
	_finspaceTransitGatewayConfiguration  string
	_finspaceUserArn                      string
	_finspaceUserName                     string
	_finspaceVolumeName                   string
	_finspaceVolumeType                   string
	_finspaceVpcConfiguration             string
)

// Create a new FinSpace environment.
// Deprecated: This method will be discontinued.
func finspace_CreateEnvironment(cfg aws.Config, client *finspace.Client) {
	input := &finspace.CreateEnvironmentInput{
		// Name: *string, // Required
	}

	if len(_finspaceName) > 0 {
		input.Name = aws.String(_finspaceName)
	}
	if len(_finspaceDataBundles) > 0 {
		input.DataBundles = append([]string(nil), _finspaceDataBundles...)
	}
	if len(_finspaceDescription) > 0 {
		input.Description = aws.String(_finspaceDescription)
	}
	if len(_finspaceFederationMode) > 0 {
		if err := assignInputField(input, "FederationMode", _finspaceFederationMode); err != nil {
			log.Errorf("invalid --federation-mode: %s", err.Error())
			return
		}
	}
	if len(_finspaceFederationParameters) > 0 {
		if err := assignInputField(input, "FederationParameters", _finspaceFederationParameters); err != nil {
			log.Errorf("invalid --federation-parameters: %s", err.Error())
			return
		}
	}
	if len(_finspaceKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_finspaceKmsKeyId)
	}
	if len(_finspaceSuperuserParameters) > 0 {
		if err := assignInputField(input, "SuperuserParameters", _finspaceSuperuserParameters); err != nil {
			log.Errorf("invalid --superuser-parameters: %s", err.Error())
			return
		}
	}
	if len(_finspaceTags) > 0 {
		if err := assignInputField(input, "Tags", _finspaceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a changeset for a kdb database. A changeset allows you to add and
// delete existing files by using an ordered list of change requests.
func finspace_CreateKxChangeset(cfg aws.Config, client *finspace.Client) {
	input := &finspace.CreateKxChangesetInput{
		// ChangeRequests: []types.ChangeRequest, // Required
		// ClientToken: *string, // Required
		// DatabaseName: *string, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceChangeRequests) > 0 {
		if err := assignInputField(input, "ChangeRequests", _finspaceChangeRequests); err != nil {
			log.Errorf("invalid --change-requests: %s", err.Error())
			return
		}
	}
	if len(_finspaceClientToken) > 0 {
		input.ClientToken = aws.String(_finspaceClientToken)
	}
	if len(_finspaceDatabaseName) > 0 {
		input.DatabaseName = aws.String(_finspaceDatabaseName)
	}
	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}

	if resp, err := client.CreateKxChangeset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new kdb cluster.
func finspace_CreateKxCluster(cfg aws.Config, client *finspace.Client) {
	input := &finspace.CreateKxClusterInput{
		// AzMode: types.KxAzMode, // Required
		// ClusterName: *string, // Required
		// ClusterType: types.KxClusterType, // Required
		// EnvironmentId: *string, // Required
		// ReleaseLabel: *string, // Required
		// VpcConfiguration: *types.VpcConfiguration, // Required
	}

	if len(_finspaceAzMode) > 0 {
		if err := assignInputField(input, "AzMode", _finspaceAzMode); err != nil {
			log.Errorf("invalid --az-mode: %s", err.Error())
			return
		}
	}
	if len(_finspaceClusterName) > 0 {
		input.ClusterName = aws.String(_finspaceClusterName)
	}
	if len(_finspaceClusterType) > 0 {
		if err := assignInputField(input, "ClusterType", _finspaceClusterType); err != nil {
			log.Errorf("invalid --cluster-type: %s", err.Error())
			return
		}
	}
	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceReleaseLabel) > 0 {
		input.ReleaseLabel = aws.String(_finspaceReleaseLabel)
	}
	if len(_finspaceVpcConfiguration) > 0 {
		if err := assignInputField(input, "VpcConfiguration", _finspaceVpcConfiguration); err != nil {
			log.Errorf("invalid --vpc-configuration: %s", err.Error())
			return
		}
	}
	if len(_finspaceAutoScalingConfiguration) > 0 {
		if err := assignInputField(input, "AutoScalingConfiguration", _finspaceAutoScalingConfiguration); err != nil {
			log.Errorf("invalid --auto-scaling-configuration: %s", err.Error())
			return
		}
	}
	if len(_finspaceAvailabilityZoneId) > 0 {
		input.AvailabilityZoneId = aws.String(_finspaceAvailabilityZoneId)
	}
	if len(_finspaceCacheStorageConfigurations) > 0 {
		if err := assignInputField(input, "CacheStorageConfigurations", _finspaceCacheStorageConfigurations); err != nil {
			log.Errorf("invalid --cache-storage-configurations: %s", err.Error())
			return
		}
	}
	if len(_finspaceCapacityConfiguration) > 0 {
		if err := assignInputField(input, "CapacityConfiguration", _finspaceCapacityConfiguration); err != nil {
			log.Errorf("invalid --capacity-configuration: %s", err.Error())
			return
		}
	}
	if len(_finspaceClientToken) > 0 {
		input.ClientToken = aws.String(_finspaceClientToken)
	}
	if len(_finspaceClusterDescription) > 0 {
		input.ClusterDescription = aws.String(_finspaceClusterDescription)
	}
	if len(_finspaceCode) > 0 {
		if err := assignInputField(input, "Code", _finspaceCode); err != nil {
			log.Errorf("invalid --code: %s", err.Error())
			return
		}
	}
	if len(_finspaceCommandLineArguments) > 0 {
		if err := assignInputField(input, "CommandLineArguments", _finspaceCommandLineArguments); err != nil {
			log.Errorf("invalid --command-line-arguments: %s", err.Error())
			return
		}
	}
	if len(_finspaceDatabases) > 0 {
		if err := assignInputField(input, "Databases", _finspaceDatabases); err != nil {
			log.Errorf("invalid --databases: %s", err.Error())
			return
		}
	}
	if len(_finspaceExecutionRole) > 0 {
		input.ExecutionRole = aws.String(_finspaceExecutionRole)
	}
	if len(_finspaceInitializationScript) > 0 {
		input.InitializationScript = aws.String(_finspaceInitializationScript)
	}
	if len(_finspaceSavedownStorageConfiguration) > 0 {
		if err := assignInputField(input, "SavedownStorageConfiguration", _finspaceSavedownStorageConfiguration); err != nil {
			log.Errorf("invalid --savedown-storage-configuration: %s", err.Error())
			return
		}
	}
	if len(_finspaceScalingGroupConfiguration) > 0 {
		if err := assignInputField(input, "ScalingGroupConfiguration", _finspaceScalingGroupConfiguration); err != nil {
			log.Errorf("invalid --scaling-group-configuration: %s", err.Error())
			return
		}
	}
	if len(_finspaceTags) > 0 {
		if err := assignInputField(input, "Tags", _finspaceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_finspaceTickerplantLogConfiguration) > 0 {
		if err := assignInputField(input, "TickerplantLogConfiguration", _finspaceTickerplantLogConfiguration); err != nil {
			log.Errorf("invalid --tickerplant-log-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateKxCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new kdb database in the environment.
func finspace_CreateKxDatabase(cfg aws.Config, client *finspace.Client) {
	input := &finspace.CreateKxDatabaseInput{
		// ClientToken: *string, // Required
		// DatabaseName: *string, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceClientToken) > 0 {
		input.ClientToken = aws.String(_finspaceClientToken)
	}
	if len(_finspaceDatabaseName) > 0 {
		input.DatabaseName = aws.String(_finspaceDatabaseName)
	}
	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceDescription) > 0 {
		input.Description = aws.String(_finspaceDescription)
	}
	if len(_finspaceTags) > 0 {
		if err := assignInputField(input, "Tags", _finspaceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateKxDatabase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a snapshot of kdb database with tiered storage capabilities and a
// pre-warmed cache, ready for mounting on kdb clusters. Dataviews are only
// available for clusters running on a scaling group. They are not supported on
// dedicated clusters.
func finspace_CreateKxDataview(cfg aws.Config, client *finspace.Client) {
	input := &finspace.CreateKxDataviewInput{
		// AzMode: types.KxAzMode, // Required
		// ClientToken: *string, // Required
		// DatabaseName: *string, // Required
		// DataviewName: *string, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceAzMode) > 0 {
		if err := assignInputField(input, "AzMode", _finspaceAzMode); err != nil {
			log.Errorf("invalid --az-mode: %s", err.Error())
			return
		}
	}
	if len(_finspaceClientToken) > 0 {
		input.ClientToken = aws.String(_finspaceClientToken)
	}
	if len(_finspaceDatabaseName) > 0 {
		input.DatabaseName = aws.String(_finspaceDatabaseName)
	}
	if len(_finspaceDataviewName) > 0 {
		input.DataviewName = aws.String(_finspaceDataviewName)
	}
	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceAutoUpdate) > 0 {
		if err := assignInputField(input, "AutoUpdate", _finspaceAutoUpdate); err != nil {
			log.Errorf("invalid --auto-update: %s", err.Error())
			return
		}
	}
	if len(_finspaceAvailabilityZoneId) > 0 {
		input.AvailabilityZoneId = aws.String(_finspaceAvailabilityZoneId)
	}
	if len(_finspaceChangesetId) > 0 {
		input.ChangesetId = aws.String(_finspaceChangesetId)
	}
	if len(_finspaceDescription) > 0 {
		input.Description = aws.String(_finspaceDescription)
	}
	if len(_finspaceReadWrite) > 0 {
		if err := assignInputField(input, "ReadWrite", _finspaceReadWrite); err != nil {
			log.Errorf("invalid --read-write: %s", err.Error())
			return
		}
	}
	if len(_finspaceSegmentConfigurations) > 0 {
		if err := assignInputField(input, "SegmentConfigurations", _finspaceSegmentConfigurations); err != nil {
			log.Errorf("invalid --segment-configurations: %s", err.Error())
			return
		}
	}
	if len(_finspaceTags) > 0 {
		if err := assignInputField(input, "Tags", _finspaceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateKxDataview(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a managed kdb environment for the account.
func finspace_CreateKxEnvironment(cfg aws.Config, client *finspace.Client) {
	input := &finspace.CreateKxEnvironmentInput{
		// KmsKeyId: *string, // Required
		// Name: *string, // Required
	}

	if len(_finspaceKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_finspaceKmsKeyId)
	}
	if len(_finspaceName) > 0 {
		input.Name = aws.String(_finspaceName)
	}
	if len(_finspaceClientToken) > 0 {
		input.ClientToken = aws.String(_finspaceClientToken)
	}
	if len(_finspaceDescription) > 0 {
		input.Description = aws.String(_finspaceDescription)
	}
	if len(_finspaceTags) > 0 {
		if err := assignInputField(input, "Tags", _finspaceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateKxEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new scaling group.
func finspace_CreateKxScalingGroup(cfg aws.Config, client *finspace.Client) {
	input := &finspace.CreateKxScalingGroupInput{
		// AvailabilityZoneId: *string, // Required
		// ClientToken: *string, // Required
		// EnvironmentId: *string, // Required
		// HostType: *string, // Required
		// ScalingGroupName: *string, // Required
	}

	if len(_finspaceAvailabilityZoneId) > 0 {
		input.AvailabilityZoneId = aws.String(_finspaceAvailabilityZoneId)
	}
	if len(_finspaceClientToken) > 0 {
		input.ClientToken = aws.String(_finspaceClientToken)
	}
	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceHostType) > 0 {
		input.HostType = aws.String(_finspaceHostType)
	}
	if len(_finspaceScalingGroupName) > 0 {
		input.ScalingGroupName = aws.String(_finspaceScalingGroupName)
	}
	if len(_finspaceTags) > 0 {
		if err := assignInputField(input, "Tags", _finspaceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateKxScalingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a user in FinSpace kdb environment with an associated IAM role.
func finspace_CreateKxUser(cfg aws.Config, client *finspace.Client) {
	input := &finspace.CreateKxUserInput{
		// EnvironmentId: *string, // Required
		// IamRole: *string, // Required
		// UserName: *string, // Required
	}

	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceIamRole) > 0 {
		input.IamRole = aws.String(_finspaceIamRole)
	}
	if len(_finspaceUserName) > 0 {
		input.UserName = aws.String(_finspaceUserName)
	}
	if len(_finspaceClientToken) > 0 {
		input.ClientToken = aws.String(_finspaceClientToken)
	}
	if len(_finspaceTags) > 0 {
		if err := assignInputField(input, "Tags", _finspaceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateKxUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new volume with a specific amount of throughput and storage
// capacity.
func finspace_CreateKxVolume(cfg aws.Config, client *finspace.Client) {
	input := &finspace.CreateKxVolumeInput{
		// AvailabilityZoneIds: []string, // Required
		// AzMode: types.KxAzMode, // Required
		// EnvironmentId: *string, // Required
		// VolumeName: *string, // Required
		// VolumeType: types.KxVolumeType, // Required
	}

	if len(_finspaceAvailabilityZoneIds) > 0 {
		input.AvailabilityZoneIds = append([]string(nil), _finspaceAvailabilityZoneIds...)
	}
	if len(_finspaceAzMode) > 0 {
		if err := assignInputField(input, "AzMode", _finspaceAzMode); err != nil {
			log.Errorf("invalid --az-mode: %s", err.Error())
			return
		}
	}
	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceVolumeName) > 0 {
		input.VolumeName = aws.String(_finspaceVolumeName)
	}
	if len(_finspaceVolumeType) > 0 {
		if err := assignInputField(input, "VolumeType", _finspaceVolumeType); err != nil {
			log.Errorf("invalid --volume-type: %s", err.Error())
			return
		}
	}
	if len(_finspaceClientToken) > 0 {
		input.ClientToken = aws.String(_finspaceClientToken)
	}
	if len(_finspaceDescription) > 0 {
		input.Description = aws.String(_finspaceDescription)
	}
	if len(_finspaceNas1Configuration) > 0 {
		if err := assignInputField(input, "Nas1Configuration", _finspaceNas1Configuration); err != nil {
			log.Errorf("invalid --nas1-configuration: %s", err.Error())
			return
		}
	}
	if len(_finspaceTags) > 0 {
		if err := assignInputField(input, "Tags", _finspaceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateKxVolume(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an FinSpace environment.
// Deprecated: This method will be discontinued.
func finspace_DeleteEnvironment(cfg aws.Config, client *finspace.Client) {
	input := &finspace.DeleteEnvironmentInput{
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}

	if resp, err := client.DeleteEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a kdb cluster.
func finspace_DeleteKxCluster(cfg aws.Config, client *finspace.Client) {
	input := &finspace.DeleteKxClusterInput{
		// ClusterName: *string, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceClusterName) > 0 {
		input.ClusterName = aws.String(_finspaceClusterName)
	}
	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceClientToken) > 0 {
		input.ClientToken = aws.String(_finspaceClientToken)
	}

	if resp, err := client.DeleteKxCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified nodes from a cluster.
func finspace_DeleteKxClusterNode(cfg aws.Config, client *finspace.Client) {
	input := &finspace.DeleteKxClusterNodeInput{
		// ClusterName: *string, // Required
		// EnvironmentId: *string, // Required
		// NodeId: *string, // Required
	}

	if len(_finspaceClusterName) > 0 {
		input.ClusterName = aws.String(_finspaceClusterName)
	}
	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceNodeId) > 0 {
		input.NodeId = aws.String(_finspaceNodeId)
	}

	if resp, err := client.DeleteKxClusterNode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified database and all of its associated data. This action is
// irreversible. You must copy any data out of the database before deleting it if
// the data is to be retained.
func finspace_DeleteKxDatabase(cfg aws.Config, client *finspace.Client) {
	input := &finspace.DeleteKxDatabaseInput{
		// ClientToken: *string, // Required
		// DatabaseName: *string, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceClientToken) > 0 {
		input.ClientToken = aws.String(_finspaceClientToken)
	}
	if len(_finspaceDatabaseName) > 0 {
		input.DatabaseName = aws.String(_finspaceDatabaseName)
	}
	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}

	if resp, err := client.DeleteKxDatabase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified dataview. Before deleting a dataview, make sure that it
// is not in use by any cluster.
func finspace_DeleteKxDataview(cfg aws.Config, client *finspace.Client) {
	input := &finspace.DeleteKxDataviewInput{
		// ClientToken: *string, // Required
		// DatabaseName: *string, // Required
		// DataviewName: *string, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceClientToken) > 0 {
		input.ClientToken = aws.String(_finspaceClientToken)
	}
	if len(_finspaceDatabaseName) > 0 {
		input.DatabaseName = aws.String(_finspaceDatabaseName)
	}
	if len(_finspaceDataviewName) > 0 {
		input.DataviewName = aws.String(_finspaceDataviewName)
	}
	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}

	if resp, err := client.DeleteKxDataview(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the kdb environment. This action is irreversible. Deleting a kdb
// environment will remove all the associated data and any services running in it.
func finspace_DeleteKxEnvironment(cfg aws.Config, client *finspace.Client) {
	input := &finspace.DeleteKxEnvironmentInput{
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceClientToken) > 0 {
		input.ClientToken = aws.String(_finspaceClientToken)
	}

	if resp, err := client.DeleteKxEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified scaling group. This action is irreversible. You cannot
// delete a scaling group until all the clusters running on it have been deleted.
func finspace_DeleteKxScalingGroup(cfg aws.Config, client *finspace.Client) {
	input := &finspace.DeleteKxScalingGroupInput{
		// EnvironmentId: *string, // Required
		// ScalingGroupName: *string, // Required
	}

	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceScalingGroupName) > 0 {
		input.ScalingGroupName = aws.String(_finspaceScalingGroupName)
	}
	if len(_finspaceClientToken) > 0 {
		input.ClientToken = aws.String(_finspaceClientToken)
	}

	if resp, err := client.DeleteKxScalingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a user in the specified kdb environment.
func finspace_DeleteKxUser(cfg aws.Config, client *finspace.Client) {
	input := &finspace.DeleteKxUserInput{
		// EnvironmentId: *string, // Required
		// UserName: *string, // Required
	}

	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceUserName) > 0 {
		input.UserName = aws.String(_finspaceUserName)
	}
	if len(_finspaceClientToken) > 0 {
		input.ClientToken = aws.String(_finspaceClientToken)
	}

	if resp, err := client.DeleteKxUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a volume. You can only delete a volume if it's not attached to a
// cluster or a dataview. When a volume is deleted, any data on the volume is lost.
// This action is irreversible.
func finspace_DeleteKxVolume(cfg aws.Config, client *finspace.Client) {
	input := &finspace.DeleteKxVolumeInput{
		// EnvironmentId: *string, // Required
		// VolumeName: *string, // Required
	}

	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceVolumeName) > 0 {
		input.VolumeName = aws.String(_finspaceVolumeName)
	}
	if len(_finspaceClientToken) > 0 {
		input.ClientToken = aws.String(_finspaceClientToken)
	}

	if resp, err := client.DeleteKxVolume(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the FinSpace environment object.
// Deprecated: This method will be discontinued.
func finspace_GetEnvironment(cfg aws.Config, client *finspace.Client) {
	input := &finspace.GetEnvironmentInput{
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}

	if resp, err := client.GetEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a kdb changeset.
func finspace_GetKxChangeset(cfg aws.Config, client *finspace.Client) {
	input := &finspace.GetKxChangesetInput{
		// ChangesetId: *string, // Required
		// DatabaseName: *string, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceChangesetId) > 0 {
		input.ChangesetId = aws.String(_finspaceChangesetId)
	}
	if len(_finspaceDatabaseName) > 0 {
		input.DatabaseName = aws.String(_finspaceDatabaseName)
	}
	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}

	if resp, err := client.GetKxChangeset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a kdb cluster.
func finspace_GetKxCluster(cfg aws.Config, client *finspace.Client) {
	input := &finspace.GetKxClusterInput{
		// ClusterName: *string, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceClusterName) > 0 {
		input.ClusterName = aws.String(_finspaceClusterName)
	}
	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}

	if resp, err := client.GetKxCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a connection string for a user to connect to a kdb cluster. You must
// call this API using the same role that you have defined while creating a user.
func finspace_GetKxConnectionString(cfg aws.Config, client *finspace.Client) {
	input := &finspace.GetKxConnectionStringInput{
		// ClusterName: *string, // Required
		// EnvironmentId: *string, // Required
		// UserArn: *string, // Required
	}

	if len(_finspaceClusterName) > 0 {
		input.ClusterName = aws.String(_finspaceClusterName)
	}
	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceUserArn) > 0 {
		input.UserArn = aws.String(_finspaceUserArn)
	}

	if resp, err := client.GetKxConnectionString(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns database information for the specified environment ID.
func finspace_GetKxDatabase(cfg aws.Config, client *finspace.Client) {
	input := &finspace.GetKxDatabaseInput{
		// DatabaseName: *string, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceDatabaseName) > 0 {
		input.DatabaseName = aws.String(_finspaceDatabaseName)
	}
	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}

	if resp, err := client.GetKxDatabase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details of the dataview.
func finspace_GetKxDataview(cfg aws.Config, client *finspace.Client) {
	input := &finspace.GetKxDataviewInput{
		// DatabaseName: *string, // Required
		// DataviewName: *string, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceDatabaseName) > 0 {
		input.DatabaseName = aws.String(_finspaceDatabaseName)
	}
	if len(_finspaceDataviewName) > 0 {
		input.DataviewName = aws.String(_finspaceDataviewName)
	}
	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}

	if resp, err := client.GetKxDataview(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves all the information for the specified kdb environment.
func finspace_GetKxEnvironment(cfg aws.Config, client *finspace.Client) {
	input := &finspace.GetKxEnvironmentInput{
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}

	if resp, err := client.GetKxEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details of a scaling group.
func finspace_GetKxScalingGroup(cfg aws.Config, client *finspace.Client) {
	input := &finspace.GetKxScalingGroupInput{
		// EnvironmentId: *string, // Required
		// ScalingGroupName: *string, // Required
	}

	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceScalingGroupName) > 0 {
		input.ScalingGroupName = aws.String(_finspaceScalingGroupName)
	}

	if resp, err := client.GetKxScalingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified kdb user.
func finspace_GetKxUser(cfg aws.Config, client *finspace.Client) {
	input := &finspace.GetKxUserInput{
		// EnvironmentId: *string, // Required
		// UserName: *string, // Required
	}

	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceUserName) > 0 {
		input.UserName = aws.String(_finspaceUserName)
	}

	if resp, err := client.GetKxUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the information about the volume.
func finspace_GetKxVolume(cfg aws.Config, client *finspace.Client) {
	input := &finspace.GetKxVolumeInput{
		// EnvironmentId: *string, // Required
		// VolumeName: *string, // Required
	}

	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceVolumeName) > 0 {
		input.VolumeName = aws.String(_finspaceVolumeName)
	}

	if resp, err := client.GetKxVolume(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A list of all of your FinSpace environments.
// Deprecated: This method will be discontinued.
func finspace_ListEnvironments(cfg aws.Config, client *finspace.Client) {
	input := &finspace.ListEnvironmentsInput{}

	if len(_finspaceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _finspaceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_finspaceNextToken) > 0 {
		input.NextToken = aws.String(_finspaceNextToken)
	}

	if resp, err := client.ListEnvironments(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of all the changesets for a database.
func finspace_ListKxChangesets(cfg aws.Config, client *finspace.Client) {
	input := &finspace.ListKxChangesetsInput{
		// DatabaseName: *string, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceDatabaseName) > 0 {
		input.DatabaseName = aws.String(_finspaceDatabaseName)
	}
	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _finspaceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_finspaceNextToken) > 0 {
		input.NextToken = aws.String(_finspaceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListKxChangesets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*finspace.ListKxChangesetsOutput
	p := finspace.NewListKxChangesetsPaginator(client, input)
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

// Lists all the nodes in a kdb cluster.
func finspace_ListKxClusterNodes(cfg aws.Config, client *finspace.Client) {
	input := &finspace.ListKxClusterNodesInput{
		// ClusterName: *string, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceClusterName) > 0 {
		input.ClusterName = aws.String(_finspaceClusterName)
	}
	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _finspaceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_finspaceNextToken) > 0 {
		input.NextToken = aws.String(_finspaceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListKxClusterNodes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*finspace.ListKxClusterNodesOutput
	p := finspace.NewListKxClusterNodesPaginator(client, input)
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

// Returns a list of clusters.
func finspace_ListKxClusters(cfg aws.Config, client *finspace.Client) {
	input := &finspace.ListKxClustersInput{
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceClusterType) > 0 {
		if err := assignInputField(input, "ClusterType", _finspaceClusterType); err != nil {
			log.Errorf("invalid --cluster-type: %s", err.Error())
			return
		}
	}
	if len(_finspaceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _finspaceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_finspaceNextToken) > 0 {
		input.NextToken = aws.String(_finspaceNextToken)
	}

	if resp, err := client.ListKxClusters(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of all the databases in the kdb environment.
func finspace_ListKxDatabases(cfg aws.Config, client *finspace.Client) {
	input := &finspace.ListKxDatabasesInput{
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _finspaceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_finspaceNextToken) > 0 {
		input.NextToken = aws.String(_finspaceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListKxDatabases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*finspace.ListKxDatabasesOutput
	p := finspace.NewListKxDatabasesPaginator(client, input)
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

// Returns a list of all the dataviews in the database.
func finspace_ListKxDataviews(cfg aws.Config, client *finspace.Client) {
	input := &finspace.ListKxDataviewsInput{
		// DatabaseName: *string, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceDatabaseName) > 0 {
		input.DatabaseName = aws.String(_finspaceDatabaseName)
	}
	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _finspaceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_finspaceNextToken) > 0 {
		input.NextToken = aws.String(_finspaceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListKxDataviews(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*finspace.ListKxDataviewsOutput
	p := finspace.NewListKxDataviewsPaginator(client, input)
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

// Returns a list of kdb environments created in an account.
func finspace_ListKxEnvironments(cfg aws.Config, client *finspace.Client) {
	input := &finspace.ListKxEnvironmentsInput{}

	if len(_finspaceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _finspaceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_finspaceNextToken) > 0 {
		input.NextToken = aws.String(_finspaceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListKxEnvironments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*finspace.ListKxEnvironmentsOutput
	p := finspace.NewListKxEnvironmentsPaginator(client, input)
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

// Returns a list of scaling groups in a kdb environment.
func finspace_ListKxScalingGroups(cfg aws.Config, client *finspace.Client) {
	input := &finspace.ListKxScalingGroupsInput{
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _finspaceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_finspaceNextToken) > 0 {
		input.NextToken = aws.String(_finspaceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListKxScalingGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*finspace.ListKxScalingGroupsOutput
	p := finspace.NewListKxScalingGroupsPaginator(client, input)
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

// Lists all the users in a kdb environment.
func finspace_ListKxUsers(cfg aws.Config, client *finspace.Client) {
	input := &finspace.ListKxUsersInput{
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _finspaceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_finspaceNextToken) > 0 {
		input.NextToken = aws.String(_finspaceNextToken)
	}

	if resp, err := client.ListKxUsers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the volumes in a kdb environment.
func finspace_ListKxVolumes(cfg aws.Config, client *finspace.Client) {
	input := &finspace.ListKxVolumesInput{
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _finspaceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_finspaceNextToken) > 0 {
		input.NextToken = aws.String(_finspaceNextToken)
	}
	if len(_finspaceVolumeType) > 0 {
		if err := assignInputField(input, "VolumeType", _finspaceVolumeType); err != nil {
			log.Errorf("invalid --volume-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListKxVolumes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A list of all tags for a resource.
func finspace_ListTagsForResource(cfg aws.Config, client *finspace.Client) {
	input := &finspace.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_finspaceResourceArn) > 0 {
		input.ResourceArn = aws.String(_finspaceResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds metadata tags to a FinSpace resource.
func finspace_TagResource(cfg aws.Config, client *finspace.Client) {
	input := &finspace.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_finspaceResourceArn) > 0 {
		input.ResourceArn = aws.String(_finspaceResourceArn)
	}
	if len(_finspaceTags) > 0 {
		if err := assignInputField(input, "Tags", _finspaceTags); err != nil {
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

// Removes metadata tags from a FinSpace resource.
func finspace_UntagResource(cfg aws.Config, client *finspace.Client) {
	input := &finspace.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_finspaceResourceArn) > 0 {
		input.ResourceArn = aws.String(_finspaceResourceArn)
	}
	if len(_finspaceTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _finspaceTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update your FinSpace environment.
// Deprecated: This method will be discontinued.
func finspace_UpdateEnvironment(cfg aws.Config, client *finspace.Client) {
	input := &finspace.UpdateEnvironmentInput{
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceDescription) > 0 {
		input.Description = aws.String(_finspaceDescription)
	}
	if len(_finspaceFederationMode) > 0 {
		if err := assignInputField(input, "FederationMode", _finspaceFederationMode); err != nil {
			log.Errorf("invalid --federation-mode: %s", err.Error())
			return
		}
	}
	if len(_finspaceFederationParameters) > 0 {
		if err := assignInputField(input, "FederationParameters", _finspaceFederationParameters); err != nil {
			log.Errorf("invalid --federation-parameters: %s", err.Error())
			return
		}
	}
	if len(_finspaceName) > 0 {
		input.Name = aws.String(_finspaceName)
	}

	if resp, err := client.UpdateEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows you to update code configuration on a running cluster. By using this
// API you can update the code, the initialization script path, and the command
// line arguments for a specific cluster. The configuration that you want to update
// will override any existing configurations on the cluster.
func finspace_UpdateKxClusterCodeConfiguration(cfg aws.Config, client *finspace.Client) {
	input := &finspace.UpdateKxClusterCodeConfigurationInput{
		// ClusterName: *string, // Required
		// Code: *types.CodeConfiguration, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceClusterName) > 0 {
		input.ClusterName = aws.String(_finspaceClusterName)
	}
	if len(_finspaceCode) > 0 {
		if err := assignInputField(input, "Code", _finspaceCode); err != nil {
			log.Errorf("invalid --code: %s", err.Error())
			return
		}
	}
	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceClientToken) > 0 {
		input.ClientToken = aws.String(_finspaceClientToken)
	}
	if len(_finspaceCommandLineArguments) > 0 {
		if err := assignInputField(input, "CommandLineArguments", _finspaceCommandLineArguments); err != nil {
			log.Errorf("invalid --command-line-arguments: %s", err.Error())
			return
		}
	}
	if len(_finspaceDeploymentConfiguration) > 0 {
		if err := assignInputField(input, "DeploymentConfiguration", _finspaceDeploymentConfiguration); err != nil {
			log.Errorf("invalid --deployment-configuration: %s", err.Error())
			return
		}
	}
	if len(_finspaceInitializationScript) > 0 {
		input.InitializationScript = aws.String(_finspaceInitializationScript)
	}

	if resp, err := client.UpdateKxClusterCodeConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the databases mounted on a kdb cluster, which includes the changesetId
// and all the dbPaths to be cached. This API does not allow you to change a
// database name or add a database if you created a cluster without one.
//
// Using this API you can point a cluster to a different changeset and modify a
// list of partitions being cached.
func finspace_UpdateKxClusterDatabases(cfg aws.Config, client *finspace.Client) {
	input := &finspace.UpdateKxClusterDatabasesInput{
		// ClusterName: *string, // Required
		// Databases: []types.KxDatabaseConfiguration, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceClusterName) > 0 {
		input.ClusterName = aws.String(_finspaceClusterName)
	}
	if len(_finspaceDatabases) > 0 {
		if err := assignInputField(input, "Databases", _finspaceDatabases); err != nil {
			log.Errorf("invalid --databases: %s", err.Error())
			return
		}
	}
	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceClientToken) > 0 {
		input.ClientToken = aws.String(_finspaceClientToken)
	}
	if len(_finspaceDeploymentConfiguration) > 0 {
		if err := assignInputField(input, "DeploymentConfiguration", _finspaceDeploymentConfiguration); err != nil {
			log.Errorf("invalid --deployment-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateKxClusterDatabases(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates information for the given kdb database.
func finspace_UpdateKxDatabase(cfg aws.Config, client *finspace.Client) {
	input := &finspace.UpdateKxDatabaseInput{
		// ClientToken: *string, // Required
		// DatabaseName: *string, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceClientToken) > 0 {
		input.ClientToken = aws.String(_finspaceClientToken)
	}
	if len(_finspaceDatabaseName) > 0 {
		input.DatabaseName = aws.String(_finspaceDatabaseName)
	}
	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceDescription) > 0 {
		input.Description = aws.String(_finspaceDescription)
	}

	if resp, err := client.UpdateKxDatabase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified dataview. The dataviews get automatically updated when
// any new changesets are ingested. Each update of the dataview creates a new
// version, including changeset details and cache configurations
func finspace_UpdateKxDataview(cfg aws.Config, client *finspace.Client) {
	input := &finspace.UpdateKxDataviewInput{
		// ClientToken: *string, // Required
		// DatabaseName: *string, // Required
		// DataviewName: *string, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceClientToken) > 0 {
		input.ClientToken = aws.String(_finspaceClientToken)
	}
	if len(_finspaceDatabaseName) > 0 {
		input.DatabaseName = aws.String(_finspaceDatabaseName)
	}
	if len(_finspaceDataviewName) > 0 {
		input.DataviewName = aws.String(_finspaceDataviewName)
	}
	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceChangesetId) > 0 {
		input.ChangesetId = aws.String(_finspaceChangesetId)
	}
	if len(_finspaceDescription) > 0 {
		input.Description = aws.String(_finspaceDescription)
	}
	if len(_finspaceSegmentConfigurations) > 0 {
		if err := assignInputField(input, "SegmentConfigurations", _finspaceSegmentConfigurations); err != nil {
			log.Errorf("invalid --segment-configurations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateKxDataview(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates information for the given kdb environment.
func finspace_UpdateKxEnvironment(cfg aws.Config, client *finspace.Client) {
	input := &finspace.UpdateKxEnvironmentInput{
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceClientToken) > 0 {
		input.ClientToken = aws.String(_finspaceClientToken)
	}
	if len(_finspaceDescription) > 0 {
		input.Description = aws.String(_finspaceDescription)
	}
	if len(_finspaceName) > 0 {
		input.Name = aws.String(_finspaceName)
	}

	if resp, err := client.UpdateKxEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates environment network to connect to your internal network by using a
// transit gateway. This API supports request to create a transit gateway
// attachment from FinSpace VPC to your transit gateway ID and create a custom
// Route-53 outbound resolvers.
//
// Once you send a request to update a network, you cannot change it again.
// Network update might require termination of any clusters that are running in the
// existing network.
func finspace_UpdateKxEnvironmentNetwork(cfg aws.Config, client *finspace.Client) {
	input := &finspace.UpdateKxEnvironmentNetworkInput{
		// EnvironmentId: *string, // Required
	}

	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceClientToken) > 0 {
		input.ClientToken = aws.String(_finspaceClientToken)
	}
	if len(_finspaceCustomDNSConfiguration) > 0 {
		if err := assignInputField(input, "CustomDNSConfiguration", _finspaceCustomDNSConfiguration); err != nil {
			log.Errorf("invalid --custom-dns-configuration: %s", err.Error())
			return
		}
	}
	if len(_finspaceTransitGatewayConfiguration) > 0 {
		if err := assignInputField(input, "TransitGatewayConfiguration", _finspaceTransitGatewayConfiguration); err != nil {
			log.Errorf("invalid --transit-gateway-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateKxEnvironmentNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the user details. You can only update the IAM role associated with a
// user.
func finspace_UpdateKxUser(cfg aws.Config, client *finspace.Client) {
	input := &finspace.UpdateKxUserInput{
		// EnvironmentId: *string, // Required
		// IamRole: *string, // Required
		// UserName: *string, // Required
	}

	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceIamRole) > 0 {
		input.IamRole = aws.String(_finspaceIamRole)
	}
	if len(_finspaceUserName) > 0 {
		input.UserName = aws.String(_finspaceUserName)
	}
	if len(_finspaceClientToken) > 0 {
		input.ClientToken = aws.String(_finspaceClientToken)
	}

	if resp, err := client.UpdateKxUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the throughput or capacity of a volume. During the update process, the
// filesystem might be unavailable for a few minutes. You can retry any operations
// after the update is complete.
func finspace_UpdateKxVolume(cfg aws.Config, client *finspace.Client) {
	input := &finspace.UpdateKxVolumeInput{
		// EnvironmentId: *string, // Required
		// VolumeName: *string, // Required
	}

	if len(_finspaceEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspaceEnvironmentId)
	}
	if len(_finspaceVolumeName) > 0 {
		input.VolumeName = aws.String(_finspaceVolumeName)
	}
	if len(_finspaceClientToken) > 0 {
		input.ClientToken = aws.String(_finspaceClientToken)
	}
	if len(_finspaceDescription) > 0 {
		input.Description = aws.String(_finspaceDescription)
	}
	if len(_finspaceNas1Configuration) > 0 {
		if err := assignInputField(input, "Nas1Configuration", _finspaceNas1Configuration); err != nil {
			log.Errorf("invalid --nas1-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateKxVolume(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_finspaceCmd)
	_finspaceCmd.Flags().SortFlags = false

	_finspaceCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_finspaceCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_finspaceCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_finspaceCmd.Flags().StringVarP(&_finspaceAutoScalingConfiguration, "auto-scaling-configuration", "", "", "Auto Scaling Configuration")
	_finspaceCmd.Flags().StringVarP(&_finspaceAutoUpdate, "auto-update", "", "", "Auto Update")
	_finspaceCmd.Flags().StringVarP(&_finspaceAvailabilityZoneId, "availability-zone-id", "", "", "Availability Zone ID")
	_finspaceCmd.Flags().StringSliceVarP(&_finspaceAvailabilityZoneIds, "availability-zone-ids", "", nil, "Availability Zone Ids")
	_finspaceCmd.Flags().StringVarP(&_finspaceAzMode, "az-mode", "", "", "AZ Mode")
	_finspaceCmd.Flags().StringVarP(&_finspaceCacheStorageConfigurations, "cache-storage-configurations", "", "", "Cache Storage Configurations")
	_finspaceCmd.Flags().StringVarP(&_finspaceCapacityConfiguration, "capacity-configuration", "", "", "Capacity Configuration")
	_finspaceCmd.Flags().StringVarP(&_finspaceChangeRequests, "change-requests", "", "", "Change Requests")
	_finspaceCmd.Flags().StringVarP(&_finspaceChangesetId, "changeset-id", "", "", "Changeset ID")
	_finspaceCmd.Flags().StringVarP(&_finspaceClientToken, "client-token", "", "", "Client Token")
	_finspaceCmd.Flags().StringVarP(&_finspaceClusterDescription, "cluster-description", "", "", "Cluster Description")
	_finspaceCmd.Flags().StringVarP(&_finspaceClusterName, "cluster-name", "", "", "Cluster Name")
	_finspaceCmd.Flags().StringVarP(&_finspaceClusterType, "cluster-type", "", "", "Cluster Type")
	_finspaceCmd.Flags().StringVarP(&_finspaceCode, "code", "", "", "Code")
	_finspaceCmd.Flags().StringVarP(&_finspaceCommandLineArguments, "command-line-arguments", "", "", "Command Line Arguments")
	_finspaceCmd.Flags().StringVarP(&_finspaceCustomDNSConfiguration, "custom-dns-configuration", "", "", "Custom DNS Configuration")
	_finspaceCmd.Flags().StringSliceVarP(&_finspaceDataBundles, "data-bundles", "", nil, "Data Bundles")
	_finspaceCmd.Flags().StringVarP(&_finspaceDatabaseName, "database-name", "", "", "Database Name")
	_finspaceCmd.Flags().StringVarP(&_finspaceDatabases, "databases", "", "", "Databases")
	_finspaceCmd.Flags().StringVarP(&_finspaceDataviewName, "dataview-name", "", "", "Dataview Name")
	_finspaceCmd.Flags().StringVarP(&_finspaceDeploymentConfiguration, "deployment-configuration", "", "", "Deployment Configuration")
	_finspaceCmd.Flags().StringVarP(&_finspaceDescription, "description", "", "", "Description")
	_finspaceCmd.Flags().StringVarP(&_finspaceEnvironmentId, "environment-id", "", "", "Environment ID")
	_finspaceCmd.Flags().StringVarP(&_finspaceExecutionRole, "execution-role", "", "", "Execution Role")
	_finspaceCmd.Flags().StringVarP(&_finspaceFederationMode, "federation-mode", "", "", "Federation Mode")
	_finspaceCmd.Flags().StringVarP(&_finspaceFederationParameters, "federation-parameters", "", "", "Federation Parameters")
	_finspaceCmd.Flags().StringVarP(&_finspaceHostType, "host-type", "", "", "Host Type")
	_finspaceCmd.Flags().StringVarP(&_finspaceIamRole, "iam-role", "", "", "IAM Role")
	_finspaceCmd.Flags().StringVarP(&_finspaceInitializationScript, "initialization-script", "", "", "Initialization Script")
	_finspaceCmd.Flags().StringVarP(&_finspaceKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_finspaceCmd.Flags().StringVarP(&_finspaceMaxResults, "max-results", "", "", "Max Results")
	_finspaceCmd.Flags().StringVarP(&_finspaceName, "name", "", "", "Name")
	_finspaceCmd.Flags().StringVarP(&_finspaceNas1Configuration, "nas1-configuration", "", "", "Nas1 Configuration")
	_finspaceCmd.Flags().StringVarP(&_finspaceNextToken, "next-token", "", "", "Next Token")
	_finspaceCmd.Flags().StringVarP(&_finspaceNodeId, "node-id", "", "", "Node ID")
	_finspaceCmd.Flags().StringVarP(&_finspaceReadWrite, "read-write", "", "", "Read Write")
	_finspaceCmd.Flags().StringVarP(&_finspaceReleaseLabel, "release-label", "", "", "Release Label")
	_finspaceCmd.Flags().StringVarP(&_finspaceResourceArn, "resource-arn", "", "", "Resource ARN")
	_finspaceCmd.Flags().StringVarP(&_finspaceSavedownStorageConfiguration, "savedown-storage-configuration", "", "", "Savedown Storage Configuration")
	_finspaceCmd.Flags().StringVarP(&_finspaceScalingGroupConfiguration, "scaling-group-configuration", "", "", "Scaling Group Configuration")
	_finspaceCmd.Flags().StringVarP(&_finspaceScalingGroupName, "scaling-group-name", "", "", "Scaling Group Name")
	_finspaceCmd.Flags().StringVarP(&_finspaceSegmentConfigurations, "segment-configurations", "", "", "Segment Configurations")
	_finspaceCmd.Flags().StringVarP(&_finspaceSuperuserParameters, "superuser-parameters", "", "", "Superuser Parameters")
	_finspaceCmd.Flags().StringSliceVarP(&_finspaceTagKeys, "tag-keys", "", nil, "Tag Keys")
	_finspaceCmd.Flags().StringVarP(&_finspaceTags, "tags", "", "", "Tags")
	_finspaceCmd.Flags().StringVarP(&_finspaceTickerplantLogConfiguration, "tickerplant-log-configuration", "", "", "Tickerplant Log Configuration")
	_finspaceCmd.Flags().StringVarP(&_finspaceTransitGatewayConfiguration, "transit-gateway-configuration", "", "", "Transit Gateway Configuration")
	_finspaceCmd.Flags().StringVarP(&_finspaceUserArn, "user-arn", "", "", "User ARN")
	_finspaceCmd.Flags().StringVarP(&_finspaceUserName, "user-name", "", "", "User Name")
	_finspaceCmd.Flags().StringVarP(&_finspaceVolumeName, "volume-name", "", "", "Volume Name")
	_finspaceCmd.Flags().StringVarP(&_finspaceVolumeType, "volume-type", "", "", "Volume Type")
	_finspaceCmd.Flags().StringVarP(&_finspaceVpcConfiguration, "vpc-configuration", "", "", "VPC Configuration")

	_finspaceCmd.Flags().BoolVarP(&_finspaceCreateEnvironment, "create-environment", "", false, "Create Environment")
	_finspaceCmd.Flags().BoolVarP(&_finspaceCreateKxChangeset, "create-kx-changeset", "", false, "Create Kx Changeset")
	_finspaceCmd.Flags().BoolVarP(&_finspaceCreateKxCluster, "create-kx-cluster", "", false, "Create Kx Cluster")
	_finspaceCmd.Flags().BoolVarP(&_finspaceCreateKxDatabase, "create-kx-database", "", false, "Create Kx Database")
	_finspaceCmd.Flags().BoolVarP(&_finspaceCreateKxDataview, "create-kx-dataview", "", false, "Create Kx Dataview")
	_finspaceCmd.Flags().BoolVarP(&_finspaceCreateKxEnvironment, "create-kx-environment", "", false, "Create Kx Environment")
	_finspaceCmd.Flags().BoolVarP(&_finspaceCreateKxScalingGroup, "create-kx-scaling-group", "", false, "Create Kx Scaling Group")
	_finspaceCmd.Flags().BoolVarP(&_finspaceCreateKxUser, "create-kx-user", "", false, "Create Kx User")
	_finspaceCmd.Flags().BoolVarP(&_finspaceCreateKxVolume, "create-kx-volume", "", false, "Create Kx Volume")
	_finspaceCmd.Flags().BoolVarP(&_finspaceDeleteEnvironment, "delete-environment", "", false, "Delete Environment")
	_finspaceCmd.Flags().BoolVarP(&_finspaceDeleteKxCluster, "delete-kx-cluster", "", false, "Delete Kx Cluster")
	_finspaceCmd.Flags().BoolVarP(&_finspaceDeleteKxClusterNode, "delete-kx-cluster-node", "", false, "Delete Kx Cluster Node")
	_finspaceCmd.Flags().BoolVarP(&_finspaceDeleteKxDatabase, "delete-kx-database", "", false, "Delete Kx Database")
	_finspaceCmd.Flags().BoolVarP(&_finspaceDeleteKxDataview, "delete-kx-dataview", "", false, "Delete Kx Dataview")
	_finspaceCmd.Flags().BoolVarP(&_finspaceDeleteKxEnvironment, "delete-kx-environment", "", false, "Delete Kx Environment")
	_finspaceCmd.Flags().BoolVarP(&_finspaceDeleteKxScalingGroup, "delete-kx-scaling-group", "", false, "Delete Kx Scaling Group")
	_finspaceCmd.Flags().BoolVarP(&_finspaceDeleteKxUser, "delete-kx-user", "", false, "Delete Kx User")
	_finspaceCmd.Flags().BoolVarP(&_finspaceDeleteKxVolume, "delete-kx-volume", "", false, "Delete Kx Volume")
	_finspaceCmd.Flags().BoolVarP(&_finspaceGetEnvironment, "get-environment", "", false, "Get Environment")
	_finspaceCmd.Flags().BoolVarP(&_finspaceGetKxChangeset, "get-kx-changeset", "", false, "Get Kx Changeset")
	_finspaceCmd.Flags().BoolVarP(&_finspaceGetKxCluster, "get-kx-cluster", "", false, "Get Kx Cluster")
	_finspaceCmd.Flags().BoolVarP(&_finspaceGetKxConnectionString, "get-kx-connection-string", "", false, "Get Kx Connection String")
	_finspaceCmd.Flags().BoolVarP(&_finspaceGetKxDatabase, "get-kx-database", "", false, "Get Kx Database")
	_finspaceCmd.Flags().BoolVarP(&_finspaceGetKxDataview, "get-kx-dataview", "", false, "Get Kx Dataview")
	_finspaceCmd.Flags().BoolVarP(&_finspaceGetKxEnvironment, "get-kx-environment", "", false, "Get Kx Environment")
	_finspaceCmd.Flags().BoolVarP(&_finspaceGetKxScalingGroup, "get-kx-scaling-group", "", false, "Get Kx Scaling Group")
	_finspaceCmd.Flags().BoolVarP(&_finspaceGetKxUser, "get-kx-user", "", false, "Get Kx User")
	_finspaceCmd.Flags().BoolVarP(&_finspaceGetKxVolume, "get-kx-volume", "", false, "Get Kx Volume")
	_finspaceCmd.Flags().BoolVarP(&_finspaceListEnvironments, "list-environments", "", false, "List Environments")
	_finspaceCmd.Flags().BoolVarP(&_finspaceListKxChangesets, "list-kx-changesets", "", false, "List Kx Changesets")
	_finspaceCmd.Flags().BoolVarP(&_finspaceListKxClusterNodes, "list-kx-cluster-nodes", "", false, "List Kx Cluster Nodes")
	_finspaceCmd.Flags().BoolVarP(&_finspaceListKxClusters, "list-kx-clusters", "", false, "List Kx Clusters")
	_finspaceCmd.Flags().BoolVarP(&_finspaceListKxDatabases, "list-kx-databases", "", false, "List Kx Databases")
	_finspaceCmd.Flags().BoolVarP(&_finspaceListKxDataviews, "list-kx-dataviews", "", false, "List Kx Dataviews")
	_finspaceCmd.Flags().BoolVarP(&_finspaceListKxEnvironments, "list-kx-environments", "", false, "List Kx Environments")
	_finspaceCmd.Flags().BoolVarP(&_finspaceListKxScalingGroups, "list-kx-scaling-groups", "", false, "List Kx Scaling Groups")
	_finspaceCmd.Flags().BoolVarP(&_finspaceListKxUsers, "list-kx-users", "", false, "List Kx Users")
	_finspaceCmd.Flags().BoolVarP(&_finspaceListKxVolumes, "list-kx-volumes", "", false, "List Kx Volumes")
	_finspaceCmd.Flags().BoolVarP(&_finspaceListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_finspaceCmd.Flags().BoolVarP(&_finspaceTagResource, "tag-resource", "", false, "Tag Resource")
	_finspaceCmd.Flags().BoolVarP(&_finspaceUntagResource, "untag-resource", "", false, "Untag Resource")
	_finspaceCmd.Flags().BoolVarP(&_finspaceUpdateEnvironment, "update-environment", "", false, "Update Environment")
	_finspaceCmd.Flags().BoolVarP(&_finspaceUpdateKxClusterCodeConfiguration, "update-kx-cluster-code-configuration", "", false, "Update Kx Cluster Code Configuration")
	_finspaceCmd.Flags().BoolVarP(&_finspaceUpdateKxClusterDatabases, "update-kx-cluster-databases", "", false, "Update Kx Cluster Databases")
	_finspaceCmd.Flags().BoolVarP(&_finspaceUpdateKxDatabase, "update-kx-database", "", false, "Update Kx Database")
	_finspaceCmd.Flags().BoolVarP(&_finspaceUpdateKxDataview, "update-kx-dataview", "", false, "Update Kx Dataview")
	_finspaceCmd.Flags().BoolVarP(&_finspaceUpdateKxEnvironment, "update-kx-environment", "", false, "Update Kx Environment")
	_finspaceCmd.Flags().BoolVarP(&_finspaceUpdateKxEnvironmentNetwork, "update-kx-environment-network", "", false, "Update Kx Environment Network")
	_finspaceCmd.Flags().BoolVarP(&_finspaceUpdateKxUser, "update-kx-user", "", false, "Update Kx User")
	_finspaceCmd.Flags().BoolVarP(&_finspaceUpdateKxVolume, "update-kx-volume", "", false, "Update Kx Volume")

}
