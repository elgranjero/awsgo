package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dax"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// daxCmd represents the dax command
var _daxCmd = &cobra.Command{
	Use:   "dax",
	Short: "AWS dax CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := dax.NewFromConfig(cfg)
		if _daxCreateCluster {
			dax_CreateCluster(cfg, client)
			return
		}
		if _daxCreateParameterGroup {
			dax_CreateParameterGroup(cfg, client)
			return
		}
		if _daxCreateSubnetGroup {
			dax_CreateSubnetGroup(cfg, client)
			return
		}
		if _daxDecreaseReplicationFactor {
			dax_DecreaseReplicationFactor(cfg, client)
			return
		}
		if _daxDeleteCluster {
			dax_DeleteCluster(cfg, client)
			return
		}
		if _daxDeleteParameterGroup {
			dax_DeleteParameterGroup(cfg, client)
			return
		}
		if _daxDeleteSubnetGroup {
			dax_DeleteSubnetGroup(cfg, client)
			return
		}
		if _daxDescribeClusters {
			dax_DescribeClusters(cfg, client)
			return
		}
		if _daxDescribeDefaultParameters {
			dax_DescribeDefaultParameters(cfg, client)
			return
		}
		if _daxDescribeEvents {
			dax_DescribeEvents(cfg, client)
			return
		}
		if _daxDescribeParameterGroups {
			dax_DescribeParameterGroups(cfg, client)
			return
		}
		if _daxDescribeParameters {
			dax_DescribeParameters(cfg, client)
			return
		}
		if _daxDescribeSubnetGroups {
			dax_DescribeSubnetGroups(cfg, client)
			return
		}
		if _daxIncreaseReplicationFactor {
			dax_IncreaseReplicationFactor(cfg, client)
			return
		}
		if _daxListTags {
			dax_ListTags(cfg, client)
			return
		}
		if _daxRebootNode {
			dax_RebootNode(cfg, client)
			return
		}
		if _daxTagResource {
			dax_TagResource(cfg, client)
			return
		}
		if _daxUntagResource {
			dax_UntagResource(cfg, client)
			return
		}
		if _daxUpdateCluster {
			dax_UpdateCluster(cfg, client)
			return
		}
		if _daxUpdateParameterGroup {
			dax_UpdateParameterGroup(cfg, client)
			return
		}
		if _daxUpdateSubnetGroup {
			dax_UpdateSubnetGroup(cfg, client)
			return
		}

	},
}

var (
	_daxCreateCluster             bool
	_daxCreateParameterGroup      bool
	_daxCreateSubnetGroup         bool
	_daxDecreaseReplicationFactor bool
	_daxDeleteCluster             bool
	_daxDeleteParameterGroup      bool
	_daxDeleteSubnetGroup         bool
	_daxDescribeClusters          bool
	_daxDescribeDefaultParameters bool
	_daxDescribeEvents            bool
	_daxDescribeParameterGroups   bool
	_daxDescribeParameters        bool
	_daxDescribeSubnetGroups      bool
	_daxIncreaseReplicationFactor bool
	_daxListTags                  bool
	_daxRebootNode                bool
	_daxTagResource               bool
	_daxUntagResource             bool
	_daxUpdateCluster             bool
	_daxUpdateParameterGroup      bool
	_daxUpdateSubnetGroup         bool

	_daxAvailabilityZones             []string
	_daxClusterEndpointEncryptionType string
	_daxClusterName                   string
	_daxClusterNames                  []string
	_daxDescription                   string
	_daxDuration                      string
	_daxEndTime                       string
	_daxIamRoleArn                    string
	_daxMaxResults                    string
	_daxNetworkType                   string
	_daxNewReplicationFactor          string
	_daxNextToken                     string
	_daxNodeId                        string
	_daxNodeIdsToRemove               []string
	_daxNodeType                      string
	_daxNotificationTopicArn          string
	_daxNotificationTopicStatus       string
	_daxParameterGroupName            string
	_daxParameterGroupNames           []string
	_daxParameterNameValues           string
	_daxPreferredMaintenanceWindow    string
	_daxReplicationFactor             string
	_daxResourceName                  string
	_daxSecurityGroupIds              []string
	_daxSource                        string
	_daxSourceName                    string
	_daxSourceType                    string
	_daxSSESpecification              string
	_daxStartTime                     string
	_daxSubnetGroupName               string
	_daxSubnetGroupNames              []string
	_daxSubnetIds                     []string
	_daxTagKeys                       []string
	_daxTags                          string
)

// Creates a DAX cluster. All nodes in the cluster run the same DAX caching
// software.
func dax_CreateCluster(cfg aws.Config, client *dax.Client) {
	input := &dax.CreateClusterInput{
		// ClusterName: *string, // Required
		// IamRoleArn: *string, // Required
		// NodeType: *string, // Required
		// ReplicationFactor: int32, // Required
	}

	if len(_daxClusterName) > 0 {
		input.ClusterName = aws.String(_daxClusterName)
	}
	if len(_daxIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_daxIamRoleArn)
	}
	if len(_daxNodeType) > 0 {
		input.NodeType = aws.String(_daxNodeType)
	}
	if len(_daxReplicationFactor) > 0 {
		if err := assignInputField(input, "ReplicationFactor", _daxReplicationFactor); err != nil {
			log.Errorf("invalid --replication-factor: %s", err.Error())
			return
		}
	}
	if len(_daxAvailabilityZones) > 0 {
		input.AvailabilityZones = append([]string(nil), _daxAvailabilityZones...)
	}
	if len(_daxClusterEndpointEncryptionType) > 0 {
		if err := assignInputField(input, "ClusterEndpointEncryptionType", _daxClusterEndpointEncryptionType); err != nil {
			log.Errorf("invalid --cluster-endpoint-encryption-type: %s", err.Error())
			return
		}
	}
	if len(_daxDescription) > 0 {
		input.Description = aws.String(_daxDescription)
	}
	if len(_daxNetworkType) > 0 {
		if err := assignInputField(input, "NetworkType", _daxNetworkType); err != nil {
			log.Errorf("invalid --network-type: %s", err.Error())
			return
		}
	}
	if len(_daxNotificationTopicArn) > 0 {
		input.NotificationTopicArn = aws.String(_daxNotificationTopicArn)
	}
	if len(_daxParameterGroupName) > 0 {
		input.ParameterGroupName = aws.String(_daxParameterGroupName)
	}
	if len(_daxPreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_daxPreferredMaintenanceWindow)
	}
	if len(_daxSSESpecification) > 0 {
		if err := assignInputField(input, "SSESpecification", _daxSSESpecification); err != nil {
			log.Errorf("invalid --sse-specification: %s", err.Error())
			return
		}
	}
	if len(_daxSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _daxSecurityGroupIds...)
	}
	if len(_daxSubnetGroupName) > 0 {
		input.SubnetGroupName = aws.String(_daxSubnetGroupName)
	}
	if len(_daxTags) > 0 {
		if err := assignInputField(input, "Tags", _daxTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new parameter group. A parameter group is a collection of parameters
// that you apply to all of the nodes in a DAX cluster.
func dax_CreateParameterGroup(cfg aws.Config, client *dax.Client) {
	input := &dax.CreateParameterGroupInput{
		// ParameterGroupName: *string, // Required
	}

	if len(_daxParameterGroupName) > 0 {
		input.ParameterGroupName = aws.String(_daxParameterGroupName)
	}
	if len(_daxDescription) > 0 {
		input.Description = aws.String(_daxDescription)
	}

	if resp, err := client.CreateParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new subnet group.
func dax_CreateSubnetGroup(cfg aws.Config, client *dax.Client) {
	input := &dax.CreateSubnetGroupInput{
		// SubnetGroupName: *string, // Required
		// SubnetIds: []string, // Required
	}

	if len(_daxSubnetGroupName) > 0 {
		input.SubnetGroupName = aws.String(_daxSubnetGroupName)
	}
	if len(_daxSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _daxSubnetIds...)
	}
	if len(_daxDescription) > 0 {
		input.Description = aws.String(_daxDescription)
	}

	if resp, err := client.CreateSubnetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes one or more nodes from a DAX cluster.
// You cannot use DecreaseReplicationFactor to remove the last node in a DAX
// cluster. If you need to do this, use DeleteCluster instead.
func dax_DecreaseReplicationFactor(cfg aws.Config, client *dax.Client) {
	input := &dax.DecreaseReplicationFactorInput{
		// ClusterName: *string, // Required
		// NewReplicationFactor: int32, // Required
	}

	if len(_daxClusterName) > 0 {
		input.ClusterName = aws.String(_daxClusterName)
	}
	if len(_daxNewReplicationFactor) > 0 {
		if err := assignInputField(input, "NewReplicationFactor", _daxNewReplicationFactor); err != nil {
			log.Errorf("invalid --new-replication-factor: %s", err.Error())
			return
		}
	}
	if len(_daxAvailabilityZones) > 0 {
		input.AvailabilityZones = append([]string(nil), _daxAvailabilityZones...)
	}
	if len(_daxNodeIdsToRemove) > 0 {
		input.NodeIdsToRemove = append([]string(nil), _daxNodeIdsToRemove...)
	}

	if resp, err := client.DecreaseReplicationFactor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a previously provisioned DAX cluster. DeleteCluster deletes all
// associated nodes, node endpoints and the DAX cluster itself. When you receive a
// successful response from this action, DAX immediately begins deleting the
// cluster; you cannot cancel or revert this action.
func dax_DeleteCluster(cfg aws.Config, client *dax.Client) {
	input := &dax.DeleteClusterInput{
		// ClusterName: *string, // Required
	}

	if len(_daxClusterName) > 0 {
		input.ClusterName = aws.String(_daxClusterName)
	}

	if resp, err := client.DeleteCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified parameter group. You cannot delete a parameter group if
// it is associated with any DAX clusters.
func dax_DeleteParameterGroup(cfg aws.Config, client *dax.Client) {
	input := &dax.DeleteParameterGroupInput{
		// ParameterGroupName: *string, // Required
	}

	if len(_daxParameterGroupName) > 0 {
		input.ParameterGroupName = aws.String(_daxParameterGroupName)
	}

	if resp, err := client.DeleteParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a subnet group.
// You cannot delete a subnet group if it is associated with any DAX clusters.
func dax_DeleteSubnetGroup(cfg aws.Config, client *dax.Client) {
	input := &dax.DeleteSubnetGroupInput{
		// SubnetGroupName: *string, // Required
	}

	if len(_daxSubnetGroupName) > 0 {
		input.SubnetGroupName = aws.String(_daxSubnetGroupName)
	}

	if resp, err := client.DeleteSubnetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about all provisioned DAX clusters if no cluster identifier
// is specified, or about a specific DAX cluster if a cluster identifier is
// supplied.
//
// If the cluster is in the CREATING state, only cluster level information will be
// displayed until all of the nodes are successfully provisioned.
//
// If the cluster is in the DELETING state, only cluster level information will be
// displayed.
//
// If nodes are currently being added to the DAX cluster, node endpoint
// information and creation time for the additional nodes will not be displayed
// until they are completely provisioned. When the DAX cluster state is available,
// the cluster is ready for use.
//
// If nodes are currently being removed from the DAX cluster, no endpoint
// information for the removed nodes is displayed.
func dax_DescribeClusters(cfg aws.Config, client *dax.Client) {
	input := &dax.DescribeClustersInput{}

	if len(_daxClusterNames) > 0 {
		input.ClusterNames = append([]string(nil), _daxClusterNames...)
	}
	if len(_daxMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _daxMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_daxNextToken) > 0 {
		input.NextToken = aws.String(_daxNextToken)
	}

	if resp, err := client.DescribeClusters(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the default system parameter information for the DAX caching software.
func dax_DescribeDefaultParameters(cfg aws.Config, client *dax.Client) {
	input := &dax.DescribeDefaultParametersInput{}

	if len(_daxMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _daxMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_daxNextToken) > 0 {
		input.NextToken = aws.String(_daxNextToken)
	}

	if resp, err := client.DescribeDefaultParameters(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns events related to DAX clusters and parameter groups. You can obtain
// events specific to a particular DAX cluster or parameter group by providing the
// name as a parameter.
//
// By default, only the events occurring within the last 24 hours are returned;
// however, you can retrieve up to 14 days' worth of events if necessary.
func dax_DescribeEvents(cfg aws.Config, client *dax.Client) {
	input := &dax.DescribeEventsInput{}

	if len(_daxDuration) > 0 {
		if err := assignInputField(input, "Duration", _daxDuration); err != nil {
			log.Errorf("invalid --duration: %s", err.Error())
			return
		}
	}
	if len(_daxEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _daxEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_daxMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _daxMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_daxNextToken) > 0 {
		input.NextToken = aws.String(_daxNextToken)
	}
	if len(_daxSourceName) > 0 {
		input.SourceName = aws.String(_daxSourceName)
	}
	if len(_daxSourceType) > 0 {
		if err := assignInputField(input, "SourceType", _daxSourceType); err != nil {
			log.Errorf("invalid --source-type: %s", err.Error())
			return
		}
	}
	if len(_daxStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _daxStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeEvents(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of parameter group descriptions. If a parameter group name is
// specified, the list will contain only the descriptions for that group.
func dax_DescribeParameterGroups(cfg aws.Config, client *dax.Client) {
	input := &dax.DescribeParameterGroupsInput{}

	if len(_daxMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _daxMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_daxNextToken) > 0 {
		input.NextToken = aws.String(_daxNextToken)
	}
	if len(_daxParameterGroupNames) > 0 {
		input.ParameterGroupNames = append([]string(nil), _daxParameterGroupNames...)
	}

	if resp, err := client.DescribeParameterGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the detailed parameter list for a particular parameter group.
func dax_DescribeParameters(cfg aws.Config, client *dax.Client) {
	input := &dax.DescribeParametersInput{
		// ParameterGroupName: *string, // Required
	}

	if len(_daxParameterGroupName) > 0 {
		input.ParameterGroupName = aws.String(_daxParameterGroupName)
	}
	if len(_daxMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _daxMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_daxNextToken) > 0 {
		input.NextToken = aws.String(_daxNextToken)
	}
	if len(_daxSource) > 0 {
		input.Source = aws.String(_daxSource)
	}

	if resp, err := client.DescribeParameters(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of subnet group descriptions. If a subnet group name is
// specified, the list will contain only the description of that group.
func dax_DescribeSubnetGroups(cfg aws.Config, client *dax.Client) {
	input := &dax.DescribeSubnetGroupsInput{}

	if len(_daxMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _daxMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_daxNextToken) > 0 {
		input.NextToken = aws.String(_daxNextToken)
	}
	if len(_daxSubnetGroupNames) > 0 {
		input.SubnetGroupNames = append([]string(nil), _daxSubnetGroupNames...)
	}

	if resp, err := client.DescribeSubnetGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more nodes to a DAX cluster.
func dax_IncreaseReplicationFactor(cfg aws.Config, client *dax.Client) {
	input := &dax.IncreaseReplicationFactorInput{
		// ClusterName: *string, // Required
		// NewReplicationFactor: int32, // Required
	}

	if len(_daxClusterName) > 0 {
		input.ClusterName = aws.String(_daxClusterName)
	}
	if len(_daxNewReplicationFactor) > 0 {
		if err := assignInputField(input, "NewReplicationFactor", _daxNewReplicationFactor); err != nil {
			log.Errorf("invalid --new-replication-factor: %s", err.Error())
			return
		}
	}
	if len(_daxAvailabilityZones) > 0 {
		input.AvailabilityZones = append([]string(nil), _daxAvailabilityZones...)
	}

	if resp, err := client.IncreaseReplicationFactor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List all of the tags for a DAX cluster. You can call ListTags up to 10 times
// per second, per account.
func dax_ListTags(cfg aws.Config, client *dax.Client) {
	input := &dax.ListTagsInput{
		// ResourceName: *string, // Required
	}

	if len(_daxResourceName) > 0 {
		input.ResourceName = aws.String(_daxResourceName)
	}
	if len(_daxNextToken) > 0 {
		input.NextToken = aws.String(_daxNextToken)
	}

	if resp, err := client.ListTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Reboots a single node of a DAX cluster. The reboot action takes place as soon
// as possible. During the reboot, the node status is set to REBOOTING.
//
// RebootNode restarts the DAX engine process and does not remove the contents of
// the cache.
func dax_RebootNode(cfg aws.Config, client *dax.Client) {
	input := &dax.RebootNodeInput{
		// ClusterName: *string, // Required
		// NodeId: *string, // Required
	}

	if len(_daxClusterName) > 0 {
		input.ClusterName = aws.String(_daxClusterName)
	}
	if len(_daxNodeId) > 0 {
		input.NodeId = aws.String(_daxNodeId)
	}

	if resp, err := client.RebootNode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a set of tags with a DAX resource. You can call TagResource up to 5
// times per second, per account.
func dax_TagResource(cfg aws.Config, client *dax.Client) {
	input := &dax.TagResourceInput{
		// ResourceName: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_daxResourceName) > 0 {
		input.ResourceName = aws.String(_daxResourceName)
	}
	if len(_daxTags) > 0 {
		if err := assignInputField(input, "Tags", _daxTags); err != nil {
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

// Removes the association of tags from a DAX resource. You can call UntagResource
// up to 5 times per second, per account.
func dax_UntagResource(cfg aws.Config, client *dax.Client) {
	input := &dax.UntagResourceInput{
		// ResourceName: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_daxResourceName) > 0 {
		input.ResourceName = aws.String(_daxResourceName)
	}
	if len(_daxTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _daxTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the settings for a DAX cluster. You can use this action to change one
// or more cluster configuration parameters by specifying the parameters and the
// new values.
func dax_UpdateCluster(cfg aws.Config, client *dax.Client) {
	input := &dax.UpdateClusterInput{
		// ClusterName: *string, // Required
	}

	if len(_daxClusterName) > 0 {
		input.ClusterName = aws.String(_daxClusterName)
	}
	if len(_daxDescription) > 0 {
		input.Description = aws.String(_daxDescription)
	}
	if len(_daxNotificationTopicArn) > 0 {
		input.NotificationTopicArn = aws.String(_daxNotificationTopicArn)
	}
	if len(_daxNotificationTopicStatus) > 0 {
		input.NotificationTopicStatus = aws.String(_daxNotificationTopicStatus)
	}
	if len(_daxParameterGroupName) > 0 {
		input.ParameterGroupName = aws.String(_daxParameterGroupName)
	}
	if len(_daxPreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_daxPreferredMaintenanceWindow)
	}
	if len(_daxSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _daxSecurityGroupIds...)
	}

	if resp, err := client.UpdateCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the parameters of a parameter group. You can modify up to 20
// parameters in a single request by submitting a list parameter name and value
// pairs.
func dax_UpdateParameterGroup(cfg aws.Config, client *dax.Client) {
	input := &dax.UpdateParameterGroupInput{
		// ParameterGroupName: *string, // Required
		// ParameterNameValues: []types.ParameterNameValue, // Required
	}

	if len(_daxParameterGroupName) > 0 {
		input.ParameterGroupName = aws.String(_daxParameterGroupName)
	}
	if len(_daxParameterNameValues) > 0 {
		if err := assignInputField(input, "ParameterNameValues", _daxParameterNameValues); err != nil {
			log.Errorf("invalid --parameter-name-values: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an existing subnet group.
func dax_UpdateSubnetGroup(cfg aws.Config, client *dax.Client) {
	input := &dax.UpdateSubnetGroupInput{
		// SubnetGroupName: *string, // Required
	}

	if len(_daxSubnetGroupName) > 0 {
		input.SubnetGroupName = aws.String(_daxSubnetGroupName)
	}
	if len(_daxDescription) > 0 {
		input.Description = aws.String(_daxDescription)
	}
	if len(_daxSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _daxSubnetIds...)
	}

	if resp, err := client.UpdateSubnetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_daxCmd)
	_daxCmd.Flags().SortFlags = false

	_daxCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_daxCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_daxCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_daxCmd.Flags().StringSliceVarP(&_daxAvailabilityZones, "availability-zones", "", nil, "Availability Zones")
	_daxCmd.Flags().StringVarP(&_daxClusterEndpointEncryptionType, "cluster-endpoint-encryption-type", "", "", "Cluster Endpoint Encryption Type")
	_daxCmd.Flags().StringVarP(&_daxClusterName, "cluster-name", "", "", "Cluster Name")
	_daxCmd.Flags().StringSliceVarP(&_daxClusterNames, "cluster-names", "", nil, "Cluster Names")
	_daxCmd.Flags().StringVarP(&_daxDescription, "description", "", "", "Description")
	_daxCmd.Flags().StringVarP(&_daxDuration, "duration", "", "", "Duration")
	_daxCmd.Flags().StringVarP(&_daxEndTime, "end-time", "", "", "End Time")
	_daxCmd.Flags().StringVarP(&_daxIamRoleArn, "iam-role-arn", "", "", "IAM Role ARN")
	_daxCmd.Flags().StringVarP(&_daxMaxResults, "max-results", "", "", "Max Results")
	_daxCmd.Flags().StringVarP(&_daxNetworkType, "network-type", "", "", "Network Type")
	_daxCmd.Flags().StringVarP(&_daxNewReplicationFactor, "new-replication-factor", "", "", "New Replication Factor")
	_daxCmd.Flags().StringVarP(&_daxNextToken, "next-token", "", "", "Next Token")
	_daxCmd.Flags().StringVarP(&_daxNodeId, "node-id", "", "", "Node ID")
	_daxCmd.Flags().StringSliceVarP(&_daxNodeIdsToRemove, "node-ids-to-remove", "", nil, "Node Ids To Remove")
	_daxCmd.Flags().StringVarP(&_daxNodeType, "node-type", "", "", "Node Type")
	_daxCmd.Flags().StringVarP(&_daxNotificationTopicArn, "notification-topic-arn", "", "", "Notification Topic ARN")
	_daxCmd.Flags().StringVarP(&_daxNotificationTopicStatus, "notification-topic-status", "", "", "Notification Topic Status")
	_daxCmd.Flags().StringVarP(&_daxParameterGroupName, "parameter-group-name", "", "", "Parameter Group Name")
	_daxCmd.Flags().StringSliceVarP(&_daxParameterGroupNames, "parameter-group-names", "", nil, "Parameter Group Names")
	_daxCmd.Flags().StringVarP(&_daxParameterNameValues, "parameter-name-values", "", "", "Parameter Name Values")
	_daxCmd.Flags().StringVarP(&_daxPreferredMaintenanceWindow, "preferred-maintenance-window", "", "", "Preferred Maintenance Window")
	_daxCmd.Flags().StringVarP(&_daxReplicationFactor, "replication-factor", "", "", "Replication Factor")
	_daxCmd.Flags().StringVarP(&_daxResourceName, "resource-name", "", "", "Resource Name")
	_daxCmd.Flags().StringSliceVarP(&_daxSecurityGroupIds, "security-group-ids", "", nil, "Security Group Ids")
	_daxCmd.Flags().StringVarP(&_daxSource, "source", "", "", "Source")
	_daxCmd.Flags().StringVarP(&_daxSourceName, "source-name", "", "", "Source Name")
	_daxCmd.Flags().StringVarP(&_daxSourceType, "source-type", "", "", "Source Type")
	_daxCmd.Flags().StringVarP(&_daxSSESpecification, "sse-specification", "", "", "SSE Specification")
	_daxCmd.Flags().StringVarP(&_daxStartTime, "start-time", "", "", "Start Time")
	_daxCmd.Flags().StringVarP(&_daxSubnetGroupName, "subnet-group-name", "", "", "Subnet Group Name")
	_daxCmd.Flags().StringSliceVarP(&_daxSubnetGroupNames, "subnet-group-names", "", nil, "Subnet Group Names")
	_daxCmd.Flags().StringSliceVarP(&_daxSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_daxCmd.Flags().StringSliceVarP(&_daxTagKeys, "tag-keys", "", nil, "Tag Keys")
	_daxCmd.Flags().StringVarP(&_daxTags, "tags", "", "", "Tags")

	_daxCmd.Flags().BoolVarP(&_daxCreateCluster, "create-cluster", "", false, "Create Cluster")
	_daxCmd.Flags().BoolVarP(&_daxCreateParameterGroup, "create-parameter-group", "", false, "Create Parameter Group")
	_daxCmd.Flags().BoolVarP(&_daxCreateSubnetGroup, "create-subnet-group", "", false, "Create Subnet Group")
	_daxCmd.Flags().BoolVarP(&_daxDecreaseReplicationFactor, "decrease-replication-factor", "", false, "Decrease Replication Factor")
	_daxCmd.Flags().BoolVarP(&_daxDeleteCluster, "delete-cluster", "", false, "Delete Cluster")
	_daxCmd.Flags().BoolVarP(&_daxDeleteParameterGroup, "delete-parameter-group", "", false, "Delete Parameter Group")
	_daxCmd.Flags().BoolVarP(&_daxDeleteSubnetGroup, "delete-subnet-group", "", false, "Delete Subnet Group")
	_daxCmd.Flags().BoolVarP(&_daxDescribeClusters, "describe-clusters", "", false, "Describe Clusters")
	_daxCmd.Flags().BoolVarP(&_daxDescribeDefaultParameters, "describe-default-parameters", "", false, "Describe Default Parameters")
	_daxCmd.Flags().BoolVarP(&_daxDescribeEvents, "describe-events", "", false, "Describe Events")
	_daxCmd.Flags().BoolVarP(&_daxDescribeParameterGroups, "describe-parameter-groups", "", false, "Describe Parameter Groups")
	_daxCmd.Flags().BoolVarP(&_daxDescribeParameters, "describe-parameters", "", false, "Describe Parameters")
	_daxCmd.Flags().BoolVarP(&_daxDescribeSubnetGroups, "describe-subnet-groups", "", false, "Describe Subnet Groups")
	_daxCmd.Flags().BoolVarP(&_daxIncreaseReplicationFactor, "increase-replication-factor", "", false, "Increase Replication Factor")
	_daxCmd.Flags().BoolVarP(&_daxListTags, "list-tags", "", false, "List Tags")
	_daxCmd.Flags().BoolVarP(&_daxRebootNode, "reboot-node", "", false, "Reboot Node")
	_daxCmd.Flags().BoolVarP(&_daxTagResource, "tag-resource", "", false, "Tag Resource")
	_daxCmd.Flags().BoolVarP(&_daxUntagResource, "untag-resource", "", false, "Untag Resource")
	_daxCmd.Flags().BoolVarP(&_daxUpdateCluster, "update-cluster", "", false, "Update Cluster")
	_daxCmd.Flags().BoolVarP(&_daxUpdateParameterGroup, "update-parameter-group", "", false, "Update Parameter Group")
	_daxCmd.Flags().BoolVarP(&_daxUpdateSubnetGroup, "update-subnet-group", "", false, "Update Subnet Group")

}
