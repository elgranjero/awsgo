package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pcs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// pcsCmd represents the pcs command
var _pcsCmd = &cobra.Command{
	Use:   "pcs",
	Short: "AWS pcs CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := pcs.NewFromConfig(cfg)
		if _pcsCreateCluster {
			pcs_CreateCluster(cfg, client)
			return
		}
		if _pcsCreateComputeNodeGroup {
			pcs_CreateComputeNodeGroup(cfg, client)
			return
		}
		if _pcsCreateQueue {
			pcs_CreateQueue(cfg, client)
			return
		}
		if _pcsDeleteCluster {
			pcs_DeleteCluster(cfg, client)
			return
		}
		if _pcsDeleteComputeNodeGroup {
			pcs_DeleteComputeNodeGroup(cfg, client)
			return
		}
		if _pcsDeleteQueue {
			pcs_DeleteQueue(cfg, client)
			return
		}
		if _pcsGetCluster {
			pcs_GetCluster(cfg, client)
			return
		}
		if _pcsGetComputeNodeGroup {
			pcs_GetComputeNodeGroup(cfg, client)
			return
		}
		if _pcsGetQueue {
			pcs_GetQueue(cfg, client)
			return
		}
		if _pcsListClusters {
			pcs_ListClusters(cfg, client)
			return
		}
		if _pcsListComputeNodeGroups {
			pcs_ListComputeNodeGroups(cfg, client)
			return
		}
		if _pcsListQueues {
			pcs_ListQueues(cfg, client)
			return
		}
		if _pcsListTagsForResource {
			pcs_ListTagsForResource(cfg, client)
			return
		}
		if _pcsRegisterComputeNodeGroupInstance {
			pcs_RegisterComputeNodeGroupInstance(cfg, client)
			return
		}
		if _pcsTagResource {
			pcs_TagResource(cfg, client)
			return
		}
		if _pcsUntagResource {
			pcs_UntagResource(cfg, client)
			return
		}
		if _pcsUpdateCluster {
			pcs_UpdateCluster(cfg, client)
			return
		}
		if _pcsUpdateComputeNodeGroup {
			pcs_UpdateComputeNodeGroup(cfg, client)
			return
		}
		if _pcsUpdateQueue {
			pcs_UpdateQueue(cfg, client)
			return
		}

	},
}

var (
	_pcsCreateCluster                    bool
	_pcsCreateComputeNodeGroup           bool
	_pcsCreateQueue                      bool
	_pcsDeleteCluster                    bool
	_pcsDeleteComputeNodeGroup           bool
	_pcsDeleteQueue                      bool
	_pcsGetCluster                       bool
	_pcsGetComputeNodeGroup              bool
	_pcsGetQueue                         bool
	_pcsListClusters                     bool
	_pcsListComputeNodeGroups            bool
	_pcsListQueues                       bool
	_pcsListTagsForResource              bool
	_pcsRegisterComputeNodeGroupInstance bool
	_pcsTagResource                      bool
	_pcsUntagResource                    bool
	_pcsUpdateCluster                    bool
	_pcsUpdateComputeNodeGroup           bool
	_pcsUpdateQueue                      bool

	_pcsAmiId                          string
	_pcsBootstrapId                    string
	_pcsClientToken                    string
	_pcsClusterIdentifier              string
	_pcsClusterName                    string
	_pcsComputeNodeGroupConfigurations string
	_pcsComputeNodeGroupIdentifier     string
	_pcsComputeNodeGroupName           string
	_pcsCustomLaunchTemplate           string
	_pcsIamInstanceProfileArn          string
	_pcsInstanceConfigs                string
	_pcsMaxResults                     string
	_pcsNetworking                     string
	_pcsNextToken                      string
	_pcsPurchaseOption                 string
	_pcsQueueIdentifier                string
	_pcsQueueName                      string
	_pcsResourceArn                    string
	_pcsScalingConfiguration           string
	_pcsScheduler                      string
	_pcsSize                           string
	_pcsSlurmConfiguration             string
	_pcsSpotOptions                    string
	_pcsSubnetIds                      []string
	_pcsTagKeys                        []string
	_pcsTags                           string
)

// Creates a cluster in your account. PCS creates the cluster controller in a
// service-owned account. The cluster controller communicates with the cluster
// resources in your account. The subnets and security groups for the cluster must
// already exist before you use this API action.
//
// It takes time for PCS to create the cluster. The cluster is in a Creating state
// until it is ready to use. There can only be 1 cluster in a Creating state per
// Amazon Web Services Region per Amazon Web Services account. CreateCluster fails
// with a ServiceQuotaExceededException if there is already a cluster in a Creating
// state.
func pcs_CreateCluster(cfg aws.Config, client *pcs.Client) {
	input := &pcs.CreateClusterInput{
		// ClusterName: *string, // Required
		// Networking: *types.NetworkingRequest, // Required
		// Scheduler: *types.SchedulerRequest, // Required
		// Size: types.Size, // Required
	}

	if len(_pcsClusterName) > 0 {
		input.ClusterName = aws.String(_pcsClusterName)
	}
	if len(_pcsNetworking) > 0 {
		if err := assignInputField(input, "Networking", _pcsNetworking); err != nil {
			log.Errorf("invalid --networking: %s", err.Error())
			return
		}
	}
	if len(_pcsScheduler) > 0 {
		if err := assignInputField(input, "Scheduler", _pcsScheduler); err != nil {
			log.Errorf("invalid --scheduler: %s", err.Error())
			return
		}
	}
	if len(_pcsSize) > 0 {
		if err := assignInputField(input, "Size", _pcsSize); err != nil {
			log.Errorf("invalid --size: %s", err.Error())
			return
		}
	}
	if len(_pcsClientToken) > 0 {
		input.ClientToken = aws.String(_pcsClientToken)
	}
	if len(_pcsSlurmConfiguration) > 0 {
		if err := assignInputField(input, "SlurmConfiguration", _pcsSlurmConfiguration); err != nil {
			log.Errorf("invalid --slurm-configuration: %s", err.Error())
			return
		}
	}
	if len(_pcsTags) > 0 {
		if err := assignInputField(input, "Tags", _pcsTags); err != nil {
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

// Creates a managed set of compute nodes. You associate a compute node group with
// a cluster through 1 or more PCS queues or as part of the login fleet. A compute
// node group includes the definition of the compute properties and lifecycle
// management. PCS uses the information you provide to this API action to launch
// compute nodes in your account. You can only specify subnets in the same Amazon
// VPC as your cluster. You receive billing charges for the compute nodes that PCS
// launches in your account. You must already have a launch template before you
// call this API. For more information, see [Launch an instance from a launch template]in the Amazon Elastic Compute Cloud
// User Guide for Linux Instances.
//
// [Launch an instance from a launch template]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html
func pcs_CreateComputeNodeGroup(cfg aws.Config, client *pcs.Client) {
	input := &pcs.CreateComputeNodeGroupInput{
		// ClusterIdentifier: *string, // Required
		// ComputeNodeGroupName: *string, // Required
		// CustomLaunchTemplate: *types.CustomLaunchTemplate, // Required
		// IamInstanceProfileArn: *string, // Required
		// InstanceConfigs: []types.InstanceConfig, // Required
		// ScalingConfiguration: *types.ScalingConfigurationRequest, // Required
		// SubnetIds: []string, // Required
	}

	if len(_pcsClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_pcsClusterIdentifier)
	}
	if len(_pcsComputeNodeGroupName) > 0 {
		input.ComputeNodeGroupName = aws.String(_pcsComputeNodeGroupName)
	}
	if len(_pcsCustomLaunchTemplate) > 0 {
		if err := assignInputField(input, "CustomLaunchTemplate", _pcsCustomLaunchTemplate); err != nil {
			log.Errorf("invalid --custom-launch-template: %s", err.Error())
			return
		}
	}
	if len(_pcsIamInstanceProfileArn) > 0 {
		input.IamInstanceProfileArn = aws.String(_pcsIamInstanceProfileArn)
	}
	if len(_pcsInstanceConfigs) > 0 {
		if err := assignInputField(input, "InstanceConfigs", _pcsInstanceConfigs); err != nil {
			log.Errorf("invalid --instance-configs: %s", err.Error())
			return
		}
	}
	if len(_pcsScalingConfiguration) > 0 {
		if err := assignInputField(input, "ScalingConfiguration", _pcsScalingConfiguration); err != nil {
			log.Errorf("invalid --scaling-configuration: %s", err.Error())
			return
		}
	}
	if len(_pcsSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _pcsSubnetIds...)
	}
	if len(_pcsAmiId) > 0 {
		input.AmiId = aws.String(_pcsAmiId)
	}
	if len(_pcsClientToken) > 0 {
		input.ClientToken = aws.String(_pcsClientToken)
	}
	if len(_pcsPurchaseOption) > 0 {
		if err := assignInputField(input, "PurchaseOption", _pcsPurchaseOption); err != nil {
			log.Errorf("invalid --purchase-option: %s", err.Error())
			return
		}
	}
	if len(_pcsSlurmConfiguration) > 0 {
		if err := assignInputField(input, "SlurmConfiguration", _pcsSlurmConfiguration); err != nil {
			log.Errorf("invalid --slurm-configuration: %s", err.Error())
			return
		}
	}
	if len(_pcsSpotOptions) > 0 {
		if err := assignInputField(input, "SpotOptions", _pcsSpotOptions); err != nil {
			log.Errorf("invalid --spot-options: %s", err.Error())
			return
		}
	}
	if len(_pcsTags) > 0 {
		if err := assignInputField(input, "Tags", _pcsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateComputeNodeGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a job queue. You must associate 1 or more compute node groups with the
// queue. You can associate 1 compute node group with multiple queues.
func pcs_CreateQueue(cfg aws.Config, client *pcs.Client) {
	input := &pcs.CreateQueueInput{
		// ClusterIdentifier: *string, // Required
		// QueueName: *string, // Required
	}

	if len(_pcsClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_pcsClusterIdentifier)
	}
	if len(_pcsQueueName) > 0 {
		input.QueueName = aws.String(_pcsQueueName)
	}
	if len(_pcsClientToken) > 0 {
		input.ClientToken = aws.String(_pcsClientToken)
	}
	if len(_pcsComputeNodeGroupConfigurations) > 0 {
		if err := assignInputField(input, "ComputeNodeGroupConfigurations", _pcsComputeNodeGroupConfigurations); err != nil {
			log.Errorf("invalid --compute-node-group-configurations: %s", err.Error())
			return
		}
	}
	if len(_pcsSlurmConfiguration) > 0 {
		if err := assignInputField(input, "SlurmConfiguration", _pcsSlurmConfiguration); err != nil {
			log.Errorf("invalid --slurm-configuration: %s", err.Error())
			return
		}
	}
	if len(_pcsTags) > 0 {
		if err := assignInputField(input, "Tags", _pcsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a cluster and all its linked resources. You must delete all queues and
// compute node groups associated with the cluster before you can delete the
// cluster.
func pcs_DeleteCluster(cfg aws.Config, client *pcs.Client) {
	input := &pcs.DeleteClusterInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_pcsClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_pcsClusterIdentifier)
	}
	if len(_pcsClientToken) > 0 {
		input.ClientToken = aws.String(_pcsClientToken)
	}

	if resp, err := client.DeleteCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a compute node group. You must delete all queues associated with the
// compute node group first.
func pcs_DeleteComputeNodeGroup(cfg aws.Config, client *pcs.Client) {
	input := &pcs.DeleteComputeNodeGroupInput{
		// ClusterIdentifier: *string, // Required
		// ComputeNodeGroupIdentifier: *string, // Required
	}

	if len(_pcsClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_pcsClusterIdentifier)
	}
	if len(_pcsComputeNodeGroupIdentifier) > 0 {
		input.ComputeNodeGroupIdentifier = aws.String(_pcsComputeNodeGroupIdentifier)
	}
	if len(_pcsClientToken) > 0 {
		input.ClientToken = aws.String(_pcsClientToken)
	}

	if resp, err := client.DeleteComputeNodeGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a job queue. If the compute node group associated with this queue isn't
// associated with any other queues, PCS terminates all the compute nodes for this
// queue.
func pcs_DeleteQueue(cfg aws.Config, client *pcs.Client) {
	input := &pcs.DeleteQueueInput{
		// ClusterIdentifier: *string, // Required
		// QueueIdentifier: *string, // Required
	}

	if len(_pcsClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_pcsClusterIdentifier)
	}
	if len(_pcsQueueIdentifier) > 0 {
		input.QueueIdentifier = aws.String(_pcsQueueIdentifier)
	}
	if len(_pcsClientToken) > 0 {
		input.ClientToken = aws.String(_pcsClientToken)
	}

	if resp, err := client.DeleteQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns detailed information about a running cluster in your account. This API
// action provides networking information, endpoint information for communication
// with the scheduler, and provisioning status.
func pcs_GetCluster(cfg aws.Config, client *pcs.Client) {
	input := &pcs.GetClusterInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_pcsClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_pcsClusterIdentifier)
	}

	if resp, err := client.GetCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns detailed information about a compute node group. This API action
// provides networking information, EC2 instance type, compute node group status,
// and scheduler (such as Slurm) configuration.
func pcs_GetComputeNodeGroup(cfg aws.Config, client *pcs.Client) {
	input := &pcs.GetComputeNodeGroupInput{
		// ClusterIdentifier: *string, // Required
		// ComputeNodeGroupIdentifier: *string, // Required
	}

	if len(_pcsClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_pcsClusterIdentifier)
	}
	if len(_pcsComputeNodeGroupIdentifier) > 0 {
		input.ComputeNodeGroupIdentifier = aws.String(_pcsComputeNodeGroupIdentifier)
	}

	if resp, err := client.GetComputeNodeGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns detailed information about a queue. The information includes the
// compute node groups that the queue uses to schedule jobs.
func pcs_GetQueue(cfg aws.Config, client *pcs.Client) {
	input := &pcs.GetQueueInput{
		// ClusterIdentifier: *string, // Required
		// QueueIdentifier: *string, // Required
	}

	if len(_pcsClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_pcsClusterIdentifier)
	}
	if len(_pcsQueueIdentifier) > 0 {
		input.QueueIdentifier = aws.String(_pcsQueueIdentifier)
	}

	if resp, err := client.GetQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of running clusters in your account.
func pcs_ListClusters(cfg aws.Config, client *pcs.Client) {
	input := &pcs.ListClustersInput{}

	if len(_pcsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pcsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pcsNextToken) > 0 {
		input.NextToken = aws.String(_pcsNextToken)
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

	var results []*pcs.ListClustersOutput
	p := pcs.NewListClustersPaginator(client, input)
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

// Returns a list of all compute node groups associated with a cluster.
func pcs_ListComputeNodeGroups(cfg aws.Config, client *pcs.Client) {
	input := &pcs.ListComputeNodeGroupsInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_pcsClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_pcsClusterIdentifier)
	}
	if len(_pcsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pcsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pcsNextToken) > 0 {
		input.NextToken = aws.String(_pcsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListComputeNodeGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pcs.ListComputeNodeGroupsOutput
	p := pcs.NewListComputeNodeGroupsPaginator(client, input)
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

// Returns a list of all queues associated with a cluster.
func pcs_ListQueues(cfg aws.Config, client *pcs.Client) {
	input := &pcs.ListQueuesInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_pcsClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_pcsClusterIdentifier)
	}
	if len(_pcsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pcsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pcsNextToken) > 0 {
		input.NextToken = aws.String(_pcsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListQueues(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pcs.ListQueuesOutput
	p := pcs.NewListQueuesPaginator(client, input)
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

// Returns a list of all tags on an PCS resource.
func pcs_ListTagsForResource(cfg aws.Config, client *pcs.Client) {
	input := &pcs.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_pcsResourceArn) > 0 {
		input.ResourceArn = aws.String(_pcsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API action isn't intended for you to use.
// PCS uses this API action to register the compute nodes it launches in your
// account.
func pcs_RegisterComputeNodeGroupInstance(cfg aws.Config, client *pcs.Client) {
	input := &pcs.RegisterComputeNodeGroupInstanceInput{
		// BootstrapId: *string, // Required
		// ClusterIdentifier: *string, // Required
	}

	if len(_pcsBootstrapId) > 0 {
		input.BootstrapId = aws.String(_pcsBootstrapId)
	}
	if len(_pcsClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_pcsClusterIdentifier)
	}

	if resp, err := client.RegisterComputeNodeGroupInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or edits tags on an PCS resource. Each tag consists of a tag key and a tag
// value. The tag key and tag value are case-sensitive strings. The tag value can
// be an empty (null) string. To add a tag, specify a new tag key and a tag value.
// To edit a tag, specify an existing tag key and a new tag value.
func pcs_TagResource(cfg aws.Config, client *pcs.Client) {
	input := &pcs.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_pcsResourceArn) > 0 {
		input.ResourceArn = aws.String(_pcsResourceArn)
	}
	if len(_pcsTags) > 0 {
		if err := assignInputField(input, "Tags", _pcsTags); err != nil {
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

// Deletes tags from an PCS resource. To delete a tag, specify the tag key and the
// Amazon Resource Name (ARN) of the PCS resource.
func pcs_UntagResource(cfg aws.Config, client *pcs.Client) {
	input := &pcs.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_pcsResourceArn) > 0 {
		input.ResourceArn = aws.String(_pcsResourceArn)
	}
	if len(_pcsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _pcsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a cluster configuration. You can modify Slurm scheduler settings,
// accounting configuration, and security groups for an existing cluster.
//
// You can only update clusters that are in ACTIVE , UPDATE_FAILED , or SUSPENDED
// state. All associated resources (queues and compute node groups) must be in
// ACTIVE state before you can update the cluster.
func pcs_UpdateCluster(cfg aws.Config, client *pcs.Client) {
	input := &pcs.UpdateClusterInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_pcsClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_pcsClusterIdentifier)
	}
	if len(_pcsClientToken) > 0 {
		input.ClientToken = aws.String(_pcsClientToken)
	}
	if len(_pcsSlurmConfiguration) > 0 {
		if err := assignInputField(input, "SlurmConfiguration", _pcsSlurmConfiguration); err != nil {
			log.Errorf("invalid --slurm-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a compute node group. You can update many of the fields related to your
// compute node group including the configurations for networking, compute nodes,
// and settings specific to your scheduler (such as Slurm).
func pcs_UpdateComputeNodeGroup(cfg aws.Config, client *pcs.Client) {
	input := &pcs.UpdateComputeNodeGroupInput{
		// ClusterIdentifier: *string, // Required
		// ComputeNodeGroupIdentifier: *string, // Required
	}

	if len(_pcsClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_pcsClusterIdentifier)
	}
	if len(_pcsComputeNodeGroupIdentifier) > 0 {
		input.ComputeNodeGroupIdentifier = aws.String(_pcsComputeNodeGroupIdentifier)
	}
	if len(_pcsAmiId) > 0 {
		input.AmiId = aws.String(_pcsAmiId)
	}
	if len(_pcsClientToken) > 0 {
		input.ClientToken = aws.String(_pcsClientToken)
	}
	if len(_pcsCustomLaunchTemplate) > 0 {
		if err := assignInputField(input, "CustomLaunchTemplate", _pcsCustomLaunchTemplate); err != nil {
			log.Errorf("invalid --custom-launch-template: %s", err.Error())
			return
		}
	}
	if len(_pcsIamInstanceProfileArn) > 0 {
		input.IamInstanceProfileArn = aws.String(_pcsIamInstanceProfileArn)
	}
	if len(_pcsPurchaseOption) > 0 {
		if err := assignInputField(input, "PurchaseOption", _pcsPurchaseOption); err != nil {
			log.Errorf("invalid --purchase-option: %s", err.Error())
			return
		}
	}
	if len(_pcsScalingConfiguration) > 0 {
		if err := assignInputField(input, "ScalingConfiguration", _pcsScalingConfiguration); err != nil {
			log.Errorf("invalid --scaling-configuration: %s", err.Error())
			return
		}
	}
	if len(_pcsSlurmConfiguration) > 0 {
		if err := assignInputField(input, "SlurmConfiguration", _pcsSlurmConfiguration); err != nil {
			log.Errorf("invalid --slurm-configuration: %s", err.Error())
			return
		}
	}
	if len(_pcsSpotOptions) > 0 {
		if err := assignInputField(input, "SpotOptions", _pcsSpotOptions); err != nil {
			log.Errorf("invalid --spot-options: %s", err.Error())
			return
		}
	}
	if len(_pcsSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _pcsSubnetIds...)
	}

	if resp, err := client.UpdateComputeNodeGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the compute node group configuration of a queue. Use this API to change
// the compute node groups that the queue can send jobs to.
func pcs_UpdateQueue(cfg aws.Config, client *pcs.Client) {
	input := &pcs.UpdateQueueInput{
		// ClusterIdentifier: *string, // Required
		// QueueIdentifier: *string, // Required
	}

	if len(_pcsClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_pcsClusterIdentifier)
	}
	if len(_pcsQueueIdentifier) > 0 {
		input.QueueIdentifier = aws.String(_pcsQueueIdentifier)
	}
	if len(_pcsClientToken) > 0 {
		input.ClientToken = aws.String(_pcsClientToken)
	}
	if len(_pcsComputeNodeGroupConfigurations) > 0 {
		if err := assignInputField(input, "ComputeNodeGroupConfigurations", _pcsComputeNodeGroupConfigurations); err != nil {
			log.Errorf("invalid --compute-node-group-configurations: %s", err.Error())
			return
		}
	}
	if len(_pcsSlurmConfiguration) > 0 {
		if err := assignInputField(input, "SlurmConfiguration", _pcsSlurmConfiguration); err != nil {
			log.Errorf("invalid --slurm-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_pcsCmd)
	_pcsCmd.Flags().SortFlags = false

	_pcsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_pcsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_pcsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_pcsCmd.Flags().StringVarP(&_pcsAmiId, "ami-id", "", "", "AMI ID")
	_pcsCmd.Flags().StringVarP(&_pcsBootstrapId, "bootstrap-id", "", "", "Bootstrap ID")
	_pcsCmd.Flags().StringVarP(&_pcsClientToken, "client-token", "", "", "Client Token")
	_pcsCmd.Flags().StringVarP(&_pcsClusterIdentifier, "cluster-identifier", "", "", "Cluster Identifier")
	_pcsCmd.Flags().StringVarP(&_pcsClusterName, "cluster-name", "", "", "Cluster Name")
	_pcsCmd.Flags().StringVarP(&_pcsComputeNodeGroupConfigurations, "compute-node-group-configurations", "", "", "Compute Node Group Configurations")
	_pcsCmd.Flags().StringVarP(&_pcsComputeNodeGroupIdentifier, "compute-node-group-identifier", "", "", "Compute Node Group Identifier")
	_pcsCmd.Flags().StringVarP(&_pcsComputeNodeGroupName, "compute-node-group-name", "", "", "Compute Node Group Name")
	_pcsCmd.Flags().StringVarP(&_pcsCustomLaunchTemplate, "custom-launch-template", "", "", "Custom Launch Template")
	_pcsCmd.Flags().StringVarP(&_pcsIamInstanceProfileArn, "iam-instance-profile-arn", "", "", "IAM Instance Profile ARN")
	_pcsCmd.Flags().StringVarP(&_pcsInstanceConfigs, "instance-configs", "", "", "Instance Configs")
	_pcsCmd.Flags().StringVarP(&_pcsMaxResults, "max-results", "", "", "Max Results")
	_pcsCmd.Flags().StringVarP(&_pcsNetworking, "networking", "", "", "Networking")
	_pcsCmd.Flags().StringVarP(&_pcsNextToken, "next-token", "", "", "Next Token")
	_pcsCmd.Flags().StringVarP(&_pcsPurchaseOption, "purchase-option", "", "", "Purchase Option")
	_pcsCmd.Flags().StringVarP(&_pcsQueueIdentifier, "queue-identifier", "", "", "Queue Identifier")
	_pcsCmd.Flags().StringVarP(&_pcsQueueName, "queue-name", "", "", "Queue Name")
	_pcsCmd.Flags().StringVarP(&_pcsResourceArn, "resource-arn", "", "", "Resource ARN")
	_pcsCmd.Flags().StringVarP(&_pcsScalingConfiguration, "scaling-configuration", "", "", "Scaling Configuration")
	_pcsCmd.Flags().StringVarP(&_pcsScheduler, "scheduler", "", "", "Scheduler")
	_pcsCmd.Flags().StringVarP(&_pcsSize, "size", "", "", "Size")
	_pcsCmd.Flags().StringVarP(&_pcsSlurmConfiguration, "slurm-configuration", "", "", "Slurm Configuration")
	_pcsCmd.Flags().StringVarP(&_pcsSpotOptions, "spot-options", "", "", "Spot Options")
	_pcsCmd.Flags().StringSliceVarP(&_pcsSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_pcsCmd.Flags().StringSliceVarP(&_pcsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_pcsCmd.Flags().StringVarP(&_pcsTags, "tags", "", "", "Tags")

	_pcsCmd.Flags().BoolVarP(&_pcsCreateCluster, "create-cluster", "", false, "Create Cluster")
	_pcsCmd.Flags().BoolVarP(&_pcsCreateComputeNodeGroup, "create-compute-node-group", "", false, "Create Compute Node Group")
	_pcsCmd.Flags().BoolVarP(&_pcsCreateQueue, "create-queue", "", false, "Create Queue")
	_pcsCmd.Flags().BoolVarP(&_pcsDeleteCluster, "delete-cluster", "", false, "Delete Cluster")
	_pcsCmd.Flags().BoolVarP(&_pcsDeleteComputeNodeGroup, "delete-compute-node-group", "", false, "Delete Compute Node Group")
	_pcsCmd.Flags().BoolVarP(&_pcsDeleteQueue, "delete-queue", "", false, "Delete Queue")
	_pcsCmd.Flags().BoolVarP(&_pcsGetCluster, "get-cluster", "", false, "Get Cluster")
	_pcsCmd.Flags().BoolVarP(&_pcsGetComputeNodeGroup, "get-compute-node-group", "", false, "Get Compute Node Group")
	_pcsCmd.Flags().BoolVarP(&_pcsGetQueue, "get-queue", "", false, "Get Queue")
	_pcsCmd.Flags().BoolVarP(&_pcsListClusters, "list-clusters", "", false, "List Clusters")
	_pcsCmd.Flags().BoolVarP(&_pcsListComputeNodeGroups, "list-compute-node-groups", "", false, "List Compute Node Groups")
	_pcsCmd.Flags().BoolVarP(&_pcsListQueues, "list-queues", "", false, "List Queues")
	_pcsCmd.Flags().BoolVarP(&_pcsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_pcsCmd.Flags().BoolVarP(&_pcsRegisterComputeNodeGroupInstance, "register-compute-node-group-instance", "", false, "Register Compute Node Group Instance")
	_pcsCmd.Flags().BoolVarP(&_pcsTagResource, "tag-resource", "", false, "Tag Resource")
	_pcsCmd.Flags().BoolVarP(&_pcsUntagResource, "untag-resource", "", false, "Untag Resource")
	_pcsCmd.Flags().BoolVarP(&_pcsUpdateCluster, "update-cluster", "", false, "Update Cluster")
	_pcsCmd.Flags().BoolVarP(&_pcsUpdateComputeNodeGroup, "update-compute-node-group", "", false, "Update Compute Node Group")
	_pcsCmd.Flags().BoolVarP(&_pcsUpdateQueue, "update-queue", "", false, "Update Queue")

}
