package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdbelastic"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// docdbelasticCmd represents the docdbelastic command
var _docdbelasticCmd = &cobra.Command{
	Use:   "docdbelastic",
	Short: "AWS docdbelastic CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := docdbelastic.NewFromConfig(cfg)
		if _docdbelasticApplyPendingMaintenanceAction {
			docdbelastic_ApplyPendingMaintenanceAction(cfg, client)
			return
		}
		if _docdbelasticCopyClusterSnapshot {
			docdbelastic_CopyClusterSnapshot(cfg, client)
			return
		}
		if _docdbelasticCreateCluster {
			docdbelastic_CreateCluster(cfg, client)
			return
		}
		if _docdbelasticCreateClusterSnapshot {
			docdbelastic_CreateClusterSnapshot(cfg, client)
			return
		}
		if _docdbelasticDeleteCluster {
			docdbelastic_DeleteCluster(cfg, client)
			return
		}
		if _docdbelasticDeleteClusterSnapshot {
			docdbelastic_DeleteClusterSnapshot(cfg, client)
			return
		}
		if _docdbelasticGetCluster {
			docdbelastic_GetCluster(cfg, client)
			return
		}
		if _docdbelasticGetClusterSnapshot {
			docdbelastic_GetClusterSnapshot(cfg, client)
			return
		}
		if _docdbelasticGetPendingMaintenanceAction {
			docdbelastic_GetPendingMaintenanceAction(cfg, client)
			return
		}
		if _docdbelasticListClusterSnapshots {
			docdbelastic_ListClusterSnapshots(cfg, client)
			return
		}
		if _docdbelasticListClusters {
			docdbelastic_ListClusters(cfg, client)
			return
		}
		if _docdbelasticListPendingMaintenanceActions {
			docdbelastic_ListPendingMaintenanceActions(cfg, client)
			return
		}
		if _docdbelasticListTagsForResource {
			docdbelastic_ListTagsForResource(cfg, client)
			return
		}
		if _docdbelasticRestoreClusterFromSnapshot {
			docdbelastic_RestoreClusterFromSnapshot(cfg, client)
			return
		}
		if _docdbelasticStartCluster {
			docdbelastic_StartCluster(cfg, client)
			return
		}
		if _docdbelasticStopCluster {
			docdbelastic_StopCluster(cfg, client)
			return
		}
		if _docdbelasticTagResource {
			docdbelastic_TagResource(cfg, client)
			return
		}
		if _docdbelasticUntagResource {
			docdbelastic_UntagResource(cfg, client)
			return
		}
		if _docdbelasticUpdateCluster {
			docdbelastic_UpdateCluster(cfg, client)
			return
		}

	},
}

var (
	_docdbelasticApplyPendingMaintenanceAction bool
	_docdbelasticCopyClusterSnapshot           bool
	_docdbelasticCreateCluster                 bool
	_docdbelasticCreateClusterSnapshot         bool
	_docdbelasticDeleteCluster                 bool
	_docdbelasticDeleteClusterSnapshot         bool
	_docdbelasticGetCluster                    bool
	_docdbelasticGetClusterSnapshot            bool
	_docdbelasticGetPendingMaintenanceAction   bool
	_docdbelasticListClusterSnapshots          bool
	_docdbelasticListClusters                  bool
	_docdbelasticListPendingMaintenanceActions bool
	_docdbelasticListTagsForResource           bool
	_docdbelasticRestoreClusterFromSnapshot    bool
	_docdbelasticStartCluster                  bool
	_docdbelasticStopCluster                   bool
	_docdbelasticTagResource                   bool
	_docdbelasticUntagResource                 bool
	_docdbelasticUpdateCluster                 bool

	_docdbelasticAdminUserName              string
	_docdbelasticAdminUserPassword          string
	_docdbelasticApplyAction                string
	_docdbelasticApplyOn                    string
	_docdbelasticAuthType                   string
	_docdbelasticBackupRetentionPeriod      string
	_docdbelasticClientToken                string
	_docdbelasticClusterArn                 string
	_docdbelasticClusterName                string
	_docdbelasticCopyTags                   string
	_docdbelasticKmsKeyId                   string
	_docdbelasticMaxResults                 string
	_docdbelasticNextToken                  string
	_docdbelasticOptInType                  string
	_docdbelasticPreferredBackupWindow      string
	_docdbelasticPreferredMaintenanceWindow string
	_docdbelasticResourceArn                string
	_docdbelasticShardCapacity              string
	_docdbelasticShardCount                 string
	_docdbelasticShardInstanceCount         string
	_docdbelasticSnapshotArn                string
	_docdbelasticSnapshotName               string
	_docdbelasticSnapshotType               string
	_docdbelasticSubnetIds                  []string
	_docdbelasticTagKeys                    []string
	_docdbelasticTags                       string
	_docdbelasticTargetSnapshotName         string
	_docdbelasticVpcSecurityGroupIds        []string
)

// The type of pending maintenance action to be applied to the resource.
func docdbelastic_ApplyPendingMaintenanceAction(cfg aws.Config, client *docdbelastic.Client) {
	input := &docdbelastic.ApplyPendingMaintenanceActionInput{
		// ApplyAction: *string, // Required
		// OptInType: types.OptInType, // Required
		// ResourceArn: *string, // Required
	}

	if len(_docdbelasticApplyAction) > 0 {
		input.ApplyAction = aws.String(_docdbelasticApplyAction)
	}
	if len(_docdbelasticOptInType) > 0 {
		if err := assignInputField(input, "OptInType", _docdbelasticOptInType); err != nil {
			log.Errorf("invalid --opt-in-type: %s", err.Error())
			return
		}
	}
	if len(_docdbelasticResourceArn) > 0 {
		input.ResourceArn = aws.String(_docdbelasticResourceArn)
	}
	if len(_docdbelasticApplyOn) > 0 {
		input.ApplyOn = aws.String(_docdbelasticApplyOn)
	}

	if resp, err := client.ApplyPendingMaintenanceAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Copies a snapshot of an elastic cluster.
func docdbelastic_CopyClusterSnapshot(cfg aws.Config, client *docdbelastic.Client) {
	input := &docdbelastic.CopyClusterSnapshotInput{
		// SnapshotArn: *string, // Required
		// TargetSnapshotName: *string, // Required
	}

	if len(_docdbelasticSnapshotArn) > 0 {
		input.SnapshotArn = aws.String(_docdbelasticSnapshotArn)
	}
	if len(_docdbelasticTargetSnapshotName) > 0 {
		input.TargetSnapshotName = aws.String(_docdbelasticTargetSnapshotName)
	}
	if len(_docdbelasticCopyTags) > 0 {
		if err := assignInputField(input, "CopyTags", _docdbelasticCopyTags); err != nil {
			log.Errorf("invalid --copy-tags: %s", err.Error())
			return
		}
	}
	if len(_docdbelasticKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_docdbelasticKmsKeyId)
	}
	if len(_docdbelasticTags) > 0 {
		if err := assignInputField(input, "Tags", _docdbelasticTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CopyClusterSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Amazon DocumentDB elastic cluster and returns its cluster
// structure.
func docdbelastic_CreateCluster(cfg aws.Config, client *docdbelastic.Client) {
	input := &docdbelastic.CreateClusterInput{
		// AdminUserName: *string, // Required
		// AdminUserPassword: *string, // Required
		// AuthType: types.Auth, // Required
		// ClusterName: *string, // Required
		// ShardCapacity: *int32, // Required
		// ShardCount: *int32, // Required
	}

	if len(_docdbelasticAdminUserName) > 0 {
		input.AdminUserName = aws.String(_docdbelasticAdminUserName)
	}
	if len(_docdbelasticAdminUserPassword) > 0 {
		input.AdminUserPassword = aws.String(_docdbelasticAdminUserPassword)
	}
	if len(_docdbelasticAuthType) > 0 {
		if err := assignInputField(input, "AuthType", _docdbelasticAuthType); err != nil {
			log.Errorf("invalid --auth-type: %s", err.Error())
			return
		}
	}
	if len(_docdbelasticClusterName) > 0 {
		input.ClusterName = aws.String(_docdbelasticClusterName)
	}
	if len(_docdbelasticShardCapacity) > 0 {
		if err := assignInputField(input, "ShardCapacity", _docdbelasticShardCapacity); err != nil {
			log.Errorf("invalid --shard-capacity: %s", err.Error())
			return
		}
	}
	if len(_docdbelasticShardCount) > 0 {
		if err := assignInputField(input, "ShardCount", _docdbelasticShardCount); err != nil {
			log.Errorf("invalid --shard-count: %s", err.Error())
			return
		}
	}
	if len(_docdbelasticBackupRetentionPeriod) > 0 {
		if err := assignInputField(input, "BackupRetentionPeriod", _docdbelasticBackupRetentionPeriod); err != nil {
			log.Errorf("invalid --backup-retention-period: %s", err.Error())
			return
		}
	}
	if len(_docdbelasticClientToken) > 0 {
		input.ClientToken = aws.String(_docdbelasticClientToken)
	}
	if len(_docdbelasticKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_docdbelasticKmsKeyId)
	}
	if len(_docdbelasticPreferredBackupWindow) > 0 {
		input.PreferredBackupWindow = aws.String(_docdbelasticPreferredBackupWindow)
	}
	if len(_docdbelasticPreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_docdbelasticPreferredMaintenanceWindow)
	}
	if len(_docdbelasticShardInstanceCount) > 0 {
		if err := assignInputField(input, "ShardInstanceCount", _docdbelasticShardInstanceCount); err != nil {
			log.Errorf("invalid --shard-instance-count: %s", err.Error())
			return
		}
	}
	if len(_docdbelasticSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _docdbelasticSubnetIds...)
	}
	if len(_docdbelasticTags) > 0 {
		if err := assignInputField(input, "Tags", _docdbelasticTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_docdbelasticVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _docdbelasticVpcSecurityGroupIds...)
	}

	if resp, err := client.CreateCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a snapshot of an elastic cluster.
func docdbelastic_CreateClusterSnapshot(cfg aws.Config, client *docdbelastic.Client) {
	input := &docdbelastic.CreateClusterSnapshotInput{
		// ClusterArn: *string, // Required
		// SnapshotName: *string, // Required
	}

	if len(_docdbelasticClusterArn) > 0 {
		input.ClusterArn = aws.String(_docdbelasticClusterArn)
	}
	if len(_docdbelasticSnapshotName) > 0 {
		input.SnapshotName = aws.String(_docdbelasticSnapshotName)
	}
	if len(_docdbelasticTags) > 0 {
		if err := assignInputField(input, "Tags", _docdbelasticTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateClusterSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an elastic cluster.
func docdbelastic_DeleteCluster(cfg aws.Config, client *docdbelastic.Client) {
	input := &docdbelastic.DeleteClusterInput{
		// ClusterArn: *string, // Required
	}

	if len(_docdbelasticClusterArn) > 0 {
		input.ClusterArn = aws.String(_docdbelasticClusterArn)
	}

	if resp, err := client.DeleteCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an elastic cluster snapshot.
func docdbelastic_DeleteClusterSnapshot(cfg aws.Config, client *docdbelastic.Client) {
	input := &docdbelastic.DeleteClusterSnapshotInput{
		// SnapshotArn: *string, // Required
	}

	if len(_docdbelasticSnapshotArn) > 0 {
		input.SnapshotArn = aws.String(_docdbelasticSnapshotArn)
	}

	if resp, err := client.DeleteClusterSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific elastic cluster.
func docdbelastic_GetCluster(cfg aws.Config, client *docdbelastic.Client) {
	input := &docdbelastic.GetClusterInput{
		// ClusterArn: *string, // Required
	}

	if len(_docdbelasticClusterArn) > 0 {
		input.ClusterArn = aws.String(_docdbelasticClusterArn)
	}

	if resp, err := client.GetCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific elastic cluster snapshot
func docdbelastic_GetClusterSnapshot(cfg aws.Config, client *docdbelastic.Client) {
	input := &docdbelastic.GetClusterSnapshotInput{
		// SnapshotArn: *string, // Required
	}

	if len(_docdbelasticSnapshotArn) > 0 {
		input.SnapshotArn = aws.String(_docdbelasticSnapshotArn)
	}

	if resp, err := client.GetClusterSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves all maintenance actions that are pending.
func docdbelastic_GetPendingMaintenanceAction(cfg aws.Config, client *docdbelastic.Client) {
	input := &docdbelastic.GetPendingMaintenanceActionInput{
		// ResourceArn: *string, // Required
	}

	if len(_docdbelasticResourceArn) > 0 {
		input.ResourceArn = aws.String(_docdbelasticResourceArn)
	}

	if resp, err := client.GetPendingMaintenanceAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about snapshots for a specified elastic cluster.
func docdbelastic_ListClusterSnapshots(cfg aws.Config, client *docdbelastic.Client) {
	input := &docdbelastic.ListClusterSnapshotsInput{}

	if len(_docdbelasticClusterArn) > 0 {
		input.ClusterArn = aws.String(_docdbelasticClusterArn)
	}
	if len(_docdbelasticMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _docdbelasticMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_docdbelasticNextToken) > 0 {
		input.NextToken = aws.String(_docdbelasticNextToken)
	}
	if len(_docdbelasticSnapshotType) > 0 {
		input.SnapshotType = aws.String(_docdbelasticSnapshotType)
	}

	if disablePaginator() {
		if resp, err := client.ListClusterSnapshots(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*docdbelastic.ListClusterSnapshotsOutput
	p := docdbelastic.NewListClusterSnapshotsPaginator(client, input)
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

// Returns information about provisioned Amazon DocumentDB elastic clusters.
func docdbelastic_ListClusters(cfg aws.Config, client *docdbelastic.Client) {
	input := &docdbelastic.ListClustersInput{}

	if len(_docdbelasticMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _docdbelasticMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_docdbelasticNextToken) > 0 {
		input.NextToken = aws.String(_docdbelasticNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListClusters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*docdbelastic.ListClustersOutput
	p := docdbelastic.NewListClustersPaginator(client, input)
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

// Retrieves a list of all maintenance actions that are pending.
func docdbelastic_ListPendingMaintenanceActions(cfg aws.Config, client *docdbelastic.Client) {
	input := &docdbelastic.ListPendingMaintenanceActionsInput{}

	if len(_docdbelasticMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _docdbelasticMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_docdbelasticNextToken) > 0 {
		input.NextToken = aws.String(_docdbelasticNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPendingMaintenanceActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*docdbelastic.ListPendingMaintenanceActionsOutput
	p := docdbelastic.NewListPendingMaintenanceActionsPaginator(client, input)
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

// Lists all tags on a elastic cluster resource
func docdbelastic_ListTagsForResource(cfg aws.Config, client *docdbelastic.Client) {
	input := &docdbelastic.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_docdbelasticResourceArn) > 0 {
		input.ResourceArn = aws.String(_docdbelasticResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restores an elastic cluster from a snapshot.
func docdbelastic_RestoreClusterFromSnapshot(cfg aws.Config, client *docdbelastic.Client) {
	input := &docdbelastic.RestoreClusterFromSnapshotInput{
		// ClusterName: *string, // Required
		// SnapshotArn: *string, // Required
	}

	if len(_docdbelasticClusterName) > 0 {
		input.ClusterName = aws.String(_docdbelasticClusterName)
	}
	if len(_docdbelasticSnapshotArn) > 0 {
		input.SnapshotArn = aws.String(_docdbelasticSnapshotArn)
	}
	if len(_docdbelasticKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_docdbelasticKmsKeyId)
	}
	if len(_docdbelasticShardCapacity) > 0 {
		if err := assignInputField(input, "ShardCapacity", _docdbelasticShardCapacity); err != nil {
			log.Errorf("invalid --shard-capacity: %s", err.Error())
			return
		}
	}
	if len(_docdbelasticShardInstanceCount) > 0 {
		if err := assignInputField(input, "ShardInstanceCount", _docdbelasticShardInstanceCount); err != nil {
			log.Errorf("invalid --shard-instance-count: %s", err.Error())
			return
		}
	}
	if len(_docdbelasticSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _docdbelasticSubnetIds...)
	}
	if len(_docdbelasticTags) > 0 {
		if err := assignInputField(input, "Tags", _docdbelasticTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_docdbelasticVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _docdbelasticVpcSecurityGroupIds...)
	}

	if resp, err := client.RestoreClusterFromSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restarts the stopped elastic cluster that is specified by clusterARN .
func docdbelastic_StartCluster(cfg aws.Config, client *docdbelastic.Client) {
	input := &docdbelastic.StartClusterInput{
		// ClusterArn: *string, // Required
	}

	if len(_docdbelasticClusterArn) > 0 {
		input.ClusterArn = aws.String(_docdbelasticClusterArn)
	}

	if resp, err := client.StartCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the running elastic cluster that is specified by clusterArn . The elastic
// cluster must be in the available state.
func docdbelastic_StopCluster(cfg aws.Config, client *docdbelastic.Client) {
	input := &docdbelastic.StopClusterInput{
		// ClusterArn: *string, // Required
	}

	if len(_docdbelasticClusterArn) > 0 {
		input.ClusterArn = aws.String(_docdbelasticClusterArn)
	}

	if resp, err := client.StopCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds metadata tags to an elastic cluster resource
func docdbelastic_TagResource(cfg aws.Config, client *docdbelastic.Client) {
	input := &docdbelastic.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_docdbelasticResourceArn) > 0 {
		input.ResourceArn = aws.String(_docdbelasticResourceArn)
	}
	if len(_docdbelasticTags) > 0 {
		if err := assignInputField(input, "Tags", _docdbelasticTags); err != nil {
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

// Removes metadata tags from an elastic cluster resource
func docdbelastic_UntagResource(cfg aws.Config, client *docdbelastic.Client) {
	input := &docdbelastic.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_docdbelasticResourceArn) > 0 {
		input.ResourceArn = aws.String(_docdbelasticResourceArn)
	}
	if len(_docdbelasticTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _docdbelasticTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an elastic cluster. This includes updating admin-username/password,
// upgrading the API version, and setting up a backup window and maintenance window
func docdbelastic_UpdateCluster(cfg aws.Config, client *docdbelastic.Client) {
	input := &docdbelastic.UpdateClusterInput{
		// ClusterArn: *string, // Required
	}

	if len(_docdbelasticClusterArn) > 0 {
		input.ClusterArn = aws.String(_docdbelasticClusterArn)
	}
	if len(_docdbelasticAdminUserPassword) > 0 {
		input.AdminUserPassword = aws.String(_docdbelasticAdminUserPassword)
	}
	if len(_docdbelasticAuthType) > 0 {
		if err := assignInputField(input, "AuthType", _docdbelasticAuthType); err != nil {
			log.Errorf("invalid --auth-type: %s", err.Error())
			return
		}
	}
	if len(_docdbelasticBackupRetentionPeriod) > 0 {
		if err := assignInputField(input, "BackupRetentionPeriod", _docdbelasticBackupRetentionPeriod); err != nil {
			log.Errorf("invalid --backup-retention-period: %s", err.Error())
			return
		}
	}
	if len(_docdbelasticClientToken) > 0 {
		input.ClientToken = aws.String(_docdbelasticClientToken)
	}
	if len(_docdbelasticPreferredBackupWindow) > 0 {
		input.PreferredBackupWindow = aws.String(_docdbelasticPreferredBackupWindow)
	}
	if len(_docdbelasticPreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_docdbelasticPreferredMaintenanceWindow)
	}
	if len(_docdbelasticShardCapacity) > 0 {
		if err := assignInputField(input, "ShardCapacity", _docdbelasticShardCapacity); err != nil {
			log.Errorf("invalid --shard-capacity: %s", err.Error())
			return
		}
	}
	if len(_docdbelasticShardCount) > 0 {
		if err := assignInputField(input, "ShardCount", _docdbelasticShardCount); err != nil {
			log.Errorf("invalid --shard-count: %s", err.Error())
			return
		}
	}
	if len(_docdbelasticShardInstanceCount) > 0 {
		if err := assignInputField(input, "ShardInstanceCount", _docdbelasticShardInstanceCount); err != nil {
			log.Errorf("invalid --shard-instance-count: %s", err.Error())
			return
		}
	}
	if len(_docdbelasticSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _docdbelasticSubnetIds...)
	}
	if len(_docdbelasticVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _docdbelasticVpcSecurityGroupIds...)
	}

	if resp, err := client.UpdateCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_docdbelasticCmd)
	_docdbelasticCmd.Flags().SortFlags = false

	_docdbelasticCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_docdbelasticCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_docdbelasticCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticAdminUserName, "admin-user-name", "", "", "Admin User Name")
	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticAdminUserPassword, "admin-user-password", "", "", "Admin User Password")
	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticApplyAction, "apply-action", "", "", "Apply Action")
	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticApplyOn, "apply-on", "", "", "Apply On")
	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticAuthType, "auth-type", "", "", "Auth Type")
	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticBackupRetentionPeriod, "backup-retention-period", "", "", "Backup Retention Period")
	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticClientToken, "client-token", "", "", "Client Token")
	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticClusterArn, "cluster-arn", "", "", "Cluster ARN")
	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticClusterName, "cluster-name", "", "", "Cluster Name")
	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticCopyTags, "copy-tags", "", "", "Copy Tags")
	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticMaxResults, "max-results", "", "", "Max Results")
	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticNextToken, "next-token", "", "", "Next Token")
	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticOptInType, "opt-in-type", "", "", "Opt In Type")
	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticPreferredBackupWindow, "preferred-backup-window", "", "", "Preferred Backup Window")
	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticPreferredMaintenanceWindow, "preferred-maintenance-window", "", "", "Preferred Maintenance Window")
	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticResourceArn, "resource-arn", "", "", "Resource ARN")
	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticShardCapacity, "shard-capacity", "", "", "Shard Capacity")
	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticShardCount, "shard-count", "", "", "Shard Count")
	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticShardInstanceCount, "shard-instance-count", "", "", "Shard Instance Count")
	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticSnapshotArn, "snapshot-arn", "", "", "Snapshot ARN")
	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticSnapshotName, "snapshot-name", "", "", "Snapshot Name")
	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticSnapshotType, "snapshot-type", "", "", "Snapshot Type")
	_docdbelasticCmd.Flags().StringSliceVarP(&_docdbelasticSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_docdbelasticCmd.Flags().StringSliceVarP(&_docdbelasticTagKeys, "tag-keys", "", nil, "Tag Keys")
	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticTags, "tags", "", "", "Tags")
	_docdbelasticCmd.Flags().StringVarP(&_docdbelasticTargetSnapshotName, "target-snapshot-name", "", "", "Target Snapshot Name")
	_docdbelasticCmd.Flags().StringSliceVarP(&_docdbelasticVpcSecurityGroupIds, "vpc-security-group-ids", "", nil, "VPC Security Group Ids")

	_docdbelasticCmd.Flags().BoolVarP(&_docdbelasticApplyPendingMaintenanceAction, "apply-pending-maintenance-action", "", false, "Apply Pending Maintenance Action")
	_docdbelasticCmd.Flags().BoolVarP(&_docdbelasticCopyClusterSnapshot, "copy-cluster-snapshot", "", false, "Copy Cluster Snapshot")
	_docdbelasticCmd.Flags().BoolVarP(&_docdbelasticCreateCluster, "create-cluster", "", false, "Create Cluster")
	_docdbelasticCmd.Flags().BoolVarP(&_docdbelasticCreateClusterSnapshot, "create-cluster-snapshot", "", false, "Create Cluster Snapshot")
	_docdbelasticCmd.Flags().BoolVarP(&_docdbelasticDeleteCluster, "delete-cluster", "", false, "Delete Cluster")
	_docdbelasticCmd.Flags().BoolVarP(&_docdbelasticDeleteClusterSnapshot, "delete-cluster-snapshot", "", false, "Delete Cluster Snapshot")
	_docdbelasticCmd.Flags().BoolVarP(&_docdbelasticGetCluster, "get-cluster", "", false, "Get Cluster")
	_docdbelasticCmd.Flags().BoolVarP(&_docdbelasticGetClusterSnapshot, "get-cluster-snapshot", "", false, "Get Cluster Snapshot")
	_docdbelasticCmd.Flags().BoolVarP(&_docdbelasticGetPendingMaintenanceAction, "get-pending-maintenance-action", "", false, "Get Pending Maintenance Action")
	_docdbelasticCmd.Flags().BoolVarP(&_docdbelasticListClusterSnapshots, "list-cluster-snapshots", "", false, "List Cluster Snapshots")
	_docdbelasticCmd.Flags().BoolVarP(&_docdbelasticListClusters, "list-clusters", "", false, "List Clusters")
	_docdbelasticCmd.Flags().BoolVarP(&_docdbelasticListPendingMaintenanceActions, "list-pending-maintenance-actions", "", false, "List Pending Maintenance Actions")
	_docdbelasticCmd.Flags().BoolVarP(&_docdbelasticListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_docdbelasticCmd.Flags().BoolVarP(&_docdbelasticRestoreClusterFromSnapshot, "restore-cluster-from-snapshot", "", false, "Restore Cluster From Snapshot")
	_docdbelasticCmd.Flags().BoolVarP(&_docdbelasticStartCluster, "start-cluster", "", false, "Start Cluster")
	_docdbelasticCmd.Flags().BoolVarP(&_docdbelasticStopCluster, "stop-cluster", "", false, "Stop Cluster")
	_docdbelasticCmd.Flags().BoolVarP(&_docdbelasticTagResource, "tag-resource", "", false, "Tag Resource")
	_docdbelasticCmd.Flags().BoolVarP(&_docdbelasticUntagResource, "untag-resource", "", false, "Untag Resource")
	_docdbelasticCmd.Flags().BoolVarP(&_docdbelasticUpdateCluster, "update-cluster", "", false, "Update Cluster")

}
