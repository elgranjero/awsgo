package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/batch"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// batchCmd represents the batch command
var _batchCmd = &cobra.Command{
	Use:   "batch",
	Short: "AWS batch CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := batch.NewFromConfig(cfg)
		if _batchCancelJob {
			batch_CancelJob(cfg, client)
			return
		}
		if _batchCreateComputeEnvironment {
			batch_CreateComputeEnvironment(cfg, client)
			return
		}
		if _batchCreateConsumableResource {
			batch_CreateConsumableResource(cfg, client)
			return
		}
		if _batchCreateJobQueue {
			batch_CreateJobQueue(cfg, client)
			return
		}
		if _batchCreateSchedulingPolicy {
			batch_CreateSchedulingPolicy(cfg, client)
			return
		}
		if _batchCreateServiceEnvironment {
			batch_CreateServiceEnvironment(cfg, client)
			return
		}
		if _batchDeleteComputeEnvironment {
			batch_DeleteComputeEnvironment(cfg, client)
			return
		}
		if _batchDeleteConsumableResource {
			batch_DeleteConsumableResource(cfg, client)
			return
		}
		if _batchDeleteJobQueue {
			batch_DeleteJobQueue(cfg, client)
			return
		}
		if _batchDeleteSchedulingPolicy {
			batch_DeleteSchedulingPolicy(cfg, client)
			return
		}
		if _batchDeleteServiceEnvironment {
			batch_DeleteServiceEnvironment(cfg, client)
			return
		}
		if _batchDeregisterJobDefinition {
			batch_DeregisterJobDefinition(cfg, client)
			return
		}
		if _batchDescribeComputeEnvironments {
			batch_DescribeComputeEnvironments(cfg, client)
			return
		}
		if _batchDescribeConsumableResource {
			batch_DescribeConsumableResource(cfg, client)
			return
		}
		if _batchDescribeJobDefinitions {
			batch_DescribeJobDefinitions(cfg, client)
			return
		}
		if _batchDescribeJobQueues {
			batch_DescribeJobQueues(cfg, client)
			return
		}
		if _batchDescribeJobs {
			batch_DescribeJobs(cfg, client)
			return
		}
		if _batchDescribeSchedulingPolicies {
			batch_DescribeSchedulingPolicies(cfg, client)
			return
		}
		if _batchDescribeServiceEnvironments {
			batch_DescribeServiceEnvironments(cfg, client)
			return
		}
		if _batchDescribeServiceJob {
			batch_DescribeServiceJob(cfg, client)
			return
		}
		if _batchGetJobQueueSnapshot {
			batch_GetJobQueueSnapshot(cfg, client)
			return
		}
		if _batchListConsumableResources {
			batch_ListConsumableResources(cfg, client)
			return
		}
		if _batchListJobs {
			batch_ListJobs(cfg, client)
			return
		}
		if _batchListJobsByConsumableResource {
			batch_ListJobsByConsumableResource(cfg, client)
			return
		}
		if _batchListSchedulingPolicies {
			batch_ListSchedulingPolicies(cfg, client)
			return
		}
		if _batchListServiceJobs {
			batch_ListServiceJobs(cfg, client)
			return
		}
		if _batchListTagsForResource {
			batch_ListTagsForResource(cfg, client)
			return
		}
		if _batchRegisterJobDefinition {
			batch_RegisterJobDefinition(cfg, client)
			return
		}
		if _batchSubmitJob {
			batch_SubmitJob(cfg, client)
			return
		}
		if _batchSubmitServiceJob {
			batch_SubmitServiceJob(cfg, client)
			return
		}
		if _batchTagResource {
			batch_TagResource(cfg, client)
			return
		}
		if _batchTerminateJob {
			batch_TerminateJob(cfg, client)
			return
		}
		if _batchTerminateServiceJob {
			batch_TerminateServiceJob(cfg, client)
			return
		}
		if _batchUntagResource {
			batch_UntagResource(cfg, client)
			return
		}
		if _batchUpdateComputeEnvironment {
			batch_UpdateComputeEnvironment(cfg, client)
			return
		}
		if _batchUpdateConsumableResource {
			batch_UpdateConsumableResource(cfg, client)
			return
		}
		if _batchUpdateJobQueue {
			batch_UpdateJobQueue(cfg, client)
			return
		}
		if _batchUpdateSchedulingPolicy {
			batch_UpdateSchedulingPolicy(cfg, client)
			return
		}
		if _batchUpdateServiceEnvironment {
			batch_UpdateServiceEnvironment(cfg, client)
			return
		}

	},
}

var (
	_batchCancelJob                    bool
	_batchCreateComputeEnvironment     bool
	_batchCreateConsumableResource     bool
	_batchCreateJobQueue               bool
	_batchCreateSchedulingPolicy       bool
	_batchCreateServiceEnvironment     bool
	_batchDeleteComputeEnvironment     bool
	_batchDeleteConsumableResource     bool
	_batchDeleteJobQueue               bool
	_batchDeleteSchedulingPolicy       bool
	_batchDeleteServiceEnvironment     bool
	_batchDeregisterJobDefinition      bool
	_batchDescribeComputeEnvironments  bool
	_batchDescribeConsumableResource   bool
	_batchDescribeJobDefinitions       bool
	_batchDescribeJobQueues            bool
	_batchDescribeJobs                 bool
	_batchDescribeSchedulingPolicies   bool
	_batchDescribeServiceEnvironments  bool
	_batchDescribeServiceJob           bool
	_batchGetJobQueueSnapshot          bool
	_batchListConsumableResources      bool
	_batchListJobs                     bool
	_batchListJobsByConsumableResource bool
	_batchListSchedulingPolicies       bool
	_batchListServiceJobs              bool
	_batchListTagsForResource          bool
	_batchRegisterJobDefinition        bool
	_batchSubmitJob                    bool
	_batchSubmitServiceJob             bool
	_batchTagResource                  bool
	_batchTerminateJob                 bool
	_batchTerminateServiceJob          bool
	_batchUntagResource                bool
	_batchUpdateComputeEnvironment     bool
	_batchUpdateConsumableResource     bool
	_batchUpdateJobQueue               bool
	_batchUpdateSchedulingPolicy       bool
	_batchUpdateServiceEnvironment     bool

	_batchArn                                  string
	_batchArns                                 []string
	_batchArrayJobId                           string
	_batchArrayProperties                      string
	_batchCapacityLimits                       string
	_batchClientToken                          string
	_batchComputeEnvironment                   string
	_batchComputeEnvironmentName               string
	_batchComputeEnvironmentOrder              string
	_batchComputeEnvironments                  []string
	_batchComputeResources                     string
	_batchConsumableResource                   string
	_batchConsumableResourceName               string
	_batchConsumableResourceProperties         string
	_batchConsumableResourcePropertiesOverride string
	_batchContainerOverrides                   string
	_batchContainerProperties                  string
	_batchContext                              string
	_batchDependsOn                            string
	_batchEcsProperties                        string
	_batchEcsPropertiesOverride                string
	_batchEksConfiguration                     string
	_batchEksProperties                        string
	_batchEksPropertiesOverride                string
	_batchFairsharePolicy                      string
	_batchFilters                              string
	_batchJobDefinition                        string
	_batchJobDefinitionName                    string
	_batchJobDefinitions                       []string
	_batchJobId                                string
	_batchJobName                              string
	_batchJobQueue                             string
	_batchJobQueueName                         string
	_batchJobQueueType                         string
	_batchJobQueues                            []string
	_batchJobStateTimeLimitActions             string
	_batchJobStatus                            string
	_batchJobs                                 []string
	_batchMaxResults                           string
	_batchMultiNodeJobId                       string
	_batchName                                 string
	_batchNextToken                            string
	_batchNodeOverrides                        string
	_batchNodeProperties                       string
	_batchOperation                            string
	_batchParameters                           string
	_batchPlatformCapabilities                 string
	_batchPriority                             string
	_batchPropagateTags                        string
	_batchQuantity                             string
	_batchReason                               string
	_batchResourceArn                          string
	_batchResourceType                         string
	_batchRetryStrategy                        string
	_batchSchedulingPolicyArn                  string
	_batchSchedulingPriority                   string
	_batchSchedulingPriorityOverride           string
	_batchServiceEnvironment                   string
	_batchServiceEnvironmentName               string
	_batchServiceEnvironmentOrder              string
	_batchServiceEnvironmentType               string
	_batchServiceEnvironments                  []string
	_batchServiceJobType                       string
	_batchServiceRequestPayload                string
	_batchServiceRole                          string
	_batchShareIdentifier                      string
	_batchState                                string
	_batchStatus                               string
	_batchTagKeys                              []string
	_batchTags                                 string
	_batchTimeout                              string
	_batchTimeoutConfig                        string
	_batchTotalQuantity                        string
	_batchType                                 string
	_batchUnmanagedvCpus                       string
	_batchUpdatePolicy                         string
)

// Cancels a job in an Batch job queue. Jobs that are in a SUBMITTED , PENDING , or
// RUNNABLE state are cancelled and the job status is updated to FAILED .
//
// A PENDING job is canceled after all dependency jobs are completed. Therefore,
// it may take longer than expected to cancel a job in PENDING status.
//
// When you try to cancel an array parent job in PENDING , Batch attempts to cancel
// all child jobs. The array parent job is canceled when all child jobs are
// completed.
//
// Jobs that progressed to the STARTING or RUNNING state aren't canceled. However,
// the API operation still succeeds, even if no job is canceled. These jobs must be
// terminated with the TerminateJoboperation.
func batch_CancelJob(cfg aws.Config, client *batch.Client) {
	input := &batch.CancelJobInput{
		// JobId: *string, // Required
		// Reason: *string, // Required
	}

	if len(_batchJobId) > 0 {
		input.JobId = aws.String(_batchJobId)
	}
	if len(_batchReason) > 0 {
		input.Reason = aws.String(_batchReason)
	}

	if resp, err := client.CancelJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Batch compute environment. You can create MANAGED or UNMANAGED
// compute environments. MANAGED compute environments can use Amazon EC2 or
// Fargate resources. UNMANAGED compute environments can only use EC2 resources.
//
// In a managed compute environment, Batch manages the capacity and instance types
// of the compute resources within the environment. This is based on the compute
// resource specification that you define or the [launch template]that you specify when you create
// the compute environment. Either, you can choose to use EC2 On-Demand Instances
// and EC2 Spot Instances. Or, you can use Fargate and Fargate Spot capacity in
// your managed compute environment. You can optionally set a maximum price so that
// Spot Instances only launch when the Spot Instance price is less than a specified
// percentage of the On-Demand price.
//
// In an unmanaged compute environment, you can manage your own EC2 compute
// resources and have flexibility with how you configure your compute resources.
// For example, you can use custom AMIs. However, you must verify that each of your
// AMIs meet the Amazon ECS container instance AMI specification. For more
// information, see [container instance AMIs]in the Amazon Elastic Container Service Developer Guide. After
// you created your unmanaged compute environment, you can use the DescribeComputeEnvironmentsoperation to
// find the Amazon ECS cluster that's associated with it. Then, launch your
// container instances into that Amazon ECS cluster. For more information, see [Launching an Amazon ECS container instance]in
// the Amazon Elastic Container Service Developer Guide.
//
// Batch doesn't automatically upgrade the AMIs in a compute environment after
// it's created. For more information on how to update a compute environment's AMI,
// see [Updating compute environments]in the Batch User Guide.
//
// [Updating compute environments]: https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html
// [launch template]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html
// [Launching an Amazon ECS container instance]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_container_instance.html
// [container instance AMIs]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container_instance_AMIs.html
func batch_CreateComputeEnvironment(cfg aws.Config, client *batch.Client) {
	input := &batch.CreateComputeEnvironmentInput{
		// ComputeEnvironmentName: *string, // Required
		// Type: types.CEType, // Required
	}

	if len(_batchComputeEnvironmentName) > 0 {
		input.ComputeEnvironmentName = aws.String(_batchComputeEnvironmentName)
	}
	if len(_batchType) > 0 {
		if err := assignInputField(input, "Type", _batchType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_batchComputeResources) > 0 {
		if err := assignInputField(input, "ComputeResources", _batchComputeResources); err != nil {
			log.Errorf("invalid --compute-resources: %s", err.Error())
			return
		}
	}
	if len(_batchContext) > 0 {
		input.Context = aws.String(_batchContext)
	}
	if len(_batchEksConfiguration) > 0 {
		if err := assignInputField(input, "EksConfiguration", _batchEksConfiguration); err != nil {
			log.Errorf("invalid --eks-configuration: %s", err.Error())
			return
		}
	}
	if len(_batchServiceRole) > 0 {
		input.ServiceRole = aws.String(_batchServiceRole)
	}
	if len(_batchState) > 0 {
		if err := assignInputField(input, "State", _batchState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}
	if len(_batchTags) > 0 {
		if err := assignInputField(input, "Tags", _batchTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_batchUnmanagedvCpus) > 0 {
		if err := assignInputField(input, "UnmanagedvCpus", _batchUnmanagedvCpus); err != nil {
			log.Errorf("invalid --unmanagedv-cpus: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateComputeEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Batch consumable resource.
func batch_CreateConsumableResource(cfg aws.Config, client *batch.Client) {
	input := &batch.CreateConsumableResourceInput{
		// ConsumableResourceName: *string, // Required
	}

	if len(_batchConsumableResourceName) > 0 {
		input.ConsumableResourceName = aws.String(_batchConsumableResourceName)
	}
	if len(_batchResourceType) > 0 {
		input.ResourceType = aws.String(_batchResourceType)
	}
	if len(_batchTags) > 0 {
		if err := assignInputField(input, "Tags", _batchTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_batchTotalQuantity) > 0 {
		if err := assignInputField(input, "TotalQuantity", _batchTotalQuantity); err != nil {
			log.Errorf("invalid --total-quantity: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConsumableResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Batch job queue. When you create a job queue, you associate one or
// more compute environments to the queue and assign an order of preference for the
// compute environments.
//
// You also set a priority to the job queue that determines the order that the
// Batch scheduler places jobs onto its associated compute environments. For
// example, if a compute environment is associated with more than one job queue,
// the job queue with a higher priority is given preference for scheduling jobs to
// that compute environment.
func batch_CreateJobQueue(cfg aws.Config, client *batch.Client) {
	input := &batch.CreateJobQueueInput{
		// JobQueueName: *string, // Required
		// Priority: *int32, // Required
	}

	if len(_batchJobQueueName) > 0 {
		input.JobQueueName = aws.String(_batchJobQueueName)
	}
	if len(_batchPriority) > 0 {
		if err := assignInputField(input, "Priority", _batchPriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_batchComputeEnvironmentOrder) > 0 {
		if err := assignInputField(input, "ComputeEnvironmentOrder", _batchComputeEnvironmentOrder); err != nil {
			log.Errorf("invalid --compute-environment-order: %s", err.Error())
			return
		}
	}
	if len(_batchJobQueueType) > 0 {
		if err := assignInputField(input, "JobQueueType", _batchJobQueueType); err != nil {
			log.Errorf("invalid --job-queue-type: %s", err.Error())
			return
		}
	}
	if len(_batchJobStateTimeLimitActions) > 0 {
		if err := assignInputField(input, "JobStateTimeLimitActions", _batchJobStateTimeLimitActions); err != nil {
			log.Errorf("invalid --job-state-time-limit-actions: %s", err.Error())
			return
		}
	}
	if len(_batchSchedulingPolicyArn) > 0 {
		input.SchedulingPolicyArn = aws.String(_batchSchedulingPolicyArn)
	}
	if len(_batchServiceEnvironmentOrder) > 0 {
		if err := assignInputField(input, "ServiceEnvironmentOrder", _batchServiceEnvironmentOrder); err != nil {
			log.Errorf("invalid --service-environment-order: %s", err.Error())
			return
		}
	}
	if len(_batchState) > 0 {
		if err := assignInputField(input, "State", _batchState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}
	if len(_batchTags) > 0 {
		if err := assignInputField(input, "Tags", _batchTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateJobQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Batch scheduling policy.
func batch_CreateSchedulingPolicy(cfg aws.Config, client *batch.Client) {
	input := &batch.CreateSchedulingPolicyInput{
		// Name: *string, // Required
	}

	if len(_batchName) > 0 {
		input.Name = aws.String(_batchName)
	}
	if len(_batchFairsharePolicy) > 0 {
		if err := assignInputField(input, "FairsharePolicy", _batchFairsharePolicy); err != nil {
			log.Errorf("invalid --fairshare-policy: %s", err.Error())
			return
		}
	}
	if len(_batchTags) > 0 {
		if err := assignInputField(input, "Tags", _batchTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSchedulingPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a service environment for running service jobs. Service environments
// define capacity limits for specific service types such as SageMaker Training
// jobs.
func batch_CreateServiceEnvironment(cfg aws.Config, client *batch.Client) {
	input := &batch.CreateServiceEnvironmentInput{
		// CapacityLimits: []types.CapacityLimit, // Required
		// ServiceEnvironmentName: *string, // Required
		// ServiceEnvironmentType: types.ServiceEnvironmentType, // Required
	}

	if len(_batchCapacityLimits) > 0 {
		if err := assignInputField(input, "CapacityLimits", _batchCapacityLimits); err != nil {
			log.Errorf("invalid --capacity-limits: %s", err.Error())
			return
		}
	}
	if len(_batchServiceEnvironmentName) > 0 {
		input.ServiceEnvironmentName = aws.String(_batchServiceEnvironmentName)
	}
	if len(_batchServiceEnvironmentType) > 0 {
		if err := assignInputField(input, "ServiceEnvironmentType", _batchServiceEnvironmentType); err != nil {
			log.Errorf("invalid --service-environment-type: %s", err.Error())
			return
		}
	}
	if len(_batchState) > 0 {
		if err := assignInputField(input, "State", _batchState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}
	if len(_batchTags) > 0 {
		if err := assignInputField(input, "Tags", _batchTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateServiceEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Batch compute environment.
// Before you can delete a compute environment, you must set its state to DISABLED
// with the UpdateComputeEnvironmentAPI operation and disassociate it from any job queues with the UpdateJobQueue API
// operation. Compute environments that use Fargate resources must terminate all
// active jobs on that compute environment before deleting the compute environment.
// If this isn't done, the compute environment enters an invalid state.
func batch_DeleteComputeEnvironment(cfg aws.Config, client *batch.Client) {
	input := &batch.DeleteComputeEnvironmentInput{
		// ComputeEnvironment: *string, // Required
	}

	if len(_batchComputeEnvironment) > 0 {
		input.ComputeEnvironment = aws.String(_batchComputeEnvironment)
	}

	if resp, err := client.DeleteComputeEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified consumable resource.
func batch_DeleteConsumableResource(cfg aws.Config, client *batch.Client) {
	input := &batch.DeleteConsumableResourceInput{
		// ConsumableResource: *string, // Required
	}

	if len(_batchConsumableResource) > 0 {
		input.ConsumableResource = aws.String(_batchConsumableResource)
	}

	if resp, err := client.DeleteConsumableResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified job queue. You must first disable submissions for a queue
// with the UpdateJobQueueoperation. All jobs in the queue are eventually terminated when you
// delete a job queue. The jobs are terminated at a rate of about 16 jobs each
// second.
//
// It's not necessary to disassociate compute environments from a queue before
// submitting a DeleteJobQueue request.
func batch_DeleteJobQueue(cfg aws.Config, client *batch.Client) {
	input := &batch.DeleteJobQueueInput{
		// JobQueue: *string, // Required
	}

	if len(_batchJobQueue) > 0 {
		input.JobQueue = aws.String(_batchJobQueue)
	}

	if resp, err := client.DeleteJobQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified scheduling policy.
// You can't delete a scheduling policy that's used in any job queues.
func batch_DeleteSchedulingPolicy(cfg aws.Config, client *batch.Client) {
	input := &batch.DeleteSchedulingPolicyInput{
		// Arn: *string, // Required
	}

	if len(_batchArn) > 0 {
		input.Arn = aws.String(_batchArn)
	}

	if resp, err := client.DeleteSchedulingPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Service environment. Before you can delete a service environment, you
// must first set its state to DISABLED with the UpdateServiceEnvironment API
// operation and disassociate it from any job queues with the UpdateJobQueue API
// operation.
func batch_DeleteServiceEnvironment(cfg aws.Config, client *batch.Client) {
	input := &batch.DeleteServiceEnvironmentInput{
		// ServiceEnvironment: *string, // Required
	}

	if len(_batchServiceEnvironment) > 0 {
		input.ServiceEnvironment = aws.String(_batchServiceEnvironment)
	}

	if resp, err := client.DeleteServiceEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregisters an Batch job definition. Job definitions are permanently deleted
// after 180 days.
func batch_DeregisterJobDefinition(cfg aws.Config, client *batch.Client) {
	input := &batch.DeregisterJobDefinitionInput{
		// JobDefinition: *string, // Required
	}

	if len(_batchJobDefinition) > 0 {
		input.JobDefinition = aws.String(_batchJobDefinition)
	}

	if resp, err := client.DeregisterJobDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes one or more of your compute environments.
// If you're using an unmanaged compute environment, you can use the
// DescribeComputeEnvironment operation to determine the ecsClusterArn that you
// launch your Amazon ECS container instances into.
func batch_DescribeComputeEnvironments(cfg aws.Config, client *batch.Client) {
	input := &batch.DescribeComputeEnvironmentsInput{}

	if len(_batchComputeEnvironments) > 0 {
		input.ComputeEnvironments = append([]string(nil), _batchComputeEnvironments...)
	}
	if len(_batchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _batchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_batchNextToken) > 0 {
		input.NextToken = aws.String(_batchNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeComputeEnvironments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*batch.DescribeComputeEnvironmentsOutput
	p := batch.NewDescribeComputeEnvironmentsPaginator(client, input)
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

// Returns a description of the specified consumable resource.
func batch_DescribeConsumableResource(cfg aws.Config, client *batch.Client) {
	input := &batch.DescribeConsumableResourceInput{
		// ConsumableResource: *string, // Required
	}

	if len(_batchConsumableResource) > 0 {
		input.ConsumableResource = aws.String(_batchConsumableResource)
	}

	if resp, err := client.DescribeConsumableResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a list of job definitions. You can specify a status (such as ACTIVE )
// to only return job definitions that match that status.
func batch_DescribeJobDefinitions(cfg aws.Config, client *batch.Client) {
	input := &batch.DescribeJobDefinitionsInput{}

	if len(_batchJobDefinitionName) > 0 {
		input.JobDefinitionName = aws.String(_batchJobDefinitionName)
	}
	if len(_batchJobDefinitions) > 0 {
		input.JobDefinitions = append([]string(nil), _batchJobDefinitions...)
	}
	if len(_batchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _batchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_batchNextToken) > 0 {
		input.NextToken = aws.String(_batchNextToken)
	}
	if len(_batchStatus) > 0 {
		input.Status = aws.String(_batchStatus)
	}

	if disablePaginator() {
		if resp, err := client.DescribeJobDefinitions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*batch.DescribeJobDefinitionsOutput
	p := batch.NewDescribeJobDefinitionsPaginator(client, input)
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

// Describes one or more of your job queues.
func batch_DescribeJobQueues(cfg aws.Config, client *batch.Client) {
	input := &batch.DescribeJobQueuesInput{}

	if len(_batchJobQueues) > 0 {
		input.JobQueues = append([]string(nil), _batchJobQueues...)
	}
	if len(_batchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _batchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_batchNextToken) > 0 {
		input.NextToken = aws.String(_batchNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeJobQueues(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*batch.DescribeJobQueuesOutput
	p := batch.NewDescribeJobQueuesPaginator(client, input)
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

// Describes a list of Batch jobs.
func batch_DescribeJobs(cfg aws.Config, client *batch.Client) {
	input := &batch.DescribeJobsInput{
		// Jobs: []string, // Required
	}

	if len(_batchJobs) > 0 {
		input.Jobs = append([]string(nil), _batchJobs...)
	}

	if resp, err := client.DescribeJobs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes one or more of your scheduling policies.
func batch_DescribeSchedulingPolicies(cfg aws.Config, client *batch.Client) {
	input := &batch.DescribeSchedulingPoliciesInput{
		// Arns: []string, // Required
	}

	if len(_batchArns) > 0 {
		input.Arns = append([]string(nil), _batchArns...)
	}

	if resp, err := client.DescribeSchedulingPolicies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes one or more of your service environments.
func batch_DescribeServiceEnvironments(cfg aws.Config, client *batch.Client) {
	input := &batch.DescribeServiceEnvironmentsInput{}

	if len(_batchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _batchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_batchNextToken) > 0 {
		input.NextToken = aws.String(_batchNextToken)
	}
	if len(_batchServiceEnvironments) > 0 {
		input.ServiceEnvironments = append([]string(nil), _batchServiceEnvironments...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeServiceEnvironments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*batch.DescribeServiceEnvironmentsOutput
	p := batch.NewDescribeServiceEnvironmentsPaginator(client, input)
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

// The details of a service job.
func batch_DescribeServiceJob(cfg aws.Config, client *batch.Client) {
	input := &batch.DescribeServiceJobInput{
		// JobId: *string, // Required
	}

	if len(_batchJobId) > 0 {
		input.JobId = aws.String(_batchJobId)
	}

	if resp, err := client.DescribeServiceJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a list of the first 100 RUNNABLE jobs associated to a single job queue
// and includes capacity utilization, including total usage and breakdown by share
// for fairshare scheduling job queues.
func batch_GetJobQueueSnapshot(cfg aws.Config, client *batch.Client) {
	input := &batch.GetJobQueueSnapshotInput{
		// JobQueue: *string, // Required
	}

	if len(_batchJobQueue) > 0 {
		input.JobQueue = aws.String(_batchJobQueue)
	}

	if resp, err := client.GetJobQueueSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of Batch consumable resources.
func batch_ListConsumableResources(cfg aws.Config, client *batch.Client) {
	input := &batch.ListConsumableResourcesInput{}

	if len(_batchFilters) > 0 {
		if err := assignInputField(input, "Filters", _batchFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_batchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _batchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_batchNextToken) > 0 {
		input.NextToken = aws.String(_batchNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConsumableResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*batch.ListConsumableResourcesOutput
	p := batch.NewListConsumableResourcesPaginator(client, input)
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

// Returns a list of Batch jobs.
// You must specify only one of the following items:
//
// - A job queue ID to return a list of jobs in that job queue
//
// - A multi-node parallel job ID to return a list of nodes for that job
//
// - An array job ID to return a list of the children for that job
func batch_ListJobs(cfg aws.Config, client *batch.Client) {
	input := &batch.ListJobsInput{}

	if len(_batchArrayJobId) > 0 {
		input.ArrayJobId = aws.String(_batchArrayJobId)
	}
	if len(_batchFilters) > 0 {
		if err := assignInputField(input, "Filters", _batchFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_batchJobQueue) > 0 {
		input.JobQueue = aws.String(_batchJobQueue)
	}
	if len(_batchJobStatus) > 0 {
		if err := assignInputField(input, "JobStatus", _batchJobStatus); err != nil {
			log.Errorf("invalid --job-status: %s", err.Error())
			return
		}
	}
	if len(_batchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _batchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_batchMultiNodeJobId) > 0 {
		input.MultiNodeJobId = aws.String(_batchMultiNodeJobId)
	}
	if len(_batchNextToken) > 0 {
		input.NextToken = aws.String(_batchNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*batch.ListJobsOutput
	p := batch.NewListJobsPaginator(client, input)
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

// Returns a list of Batch jobs that require a specific consumable resource.
func batch_ListJobsByConsumableResource(cfg aws.Config, client *batch.Client) {
	input := &batch.ListJobsByConsumableResourceInput{
		// ConsumableResource: *string, // Required
	}

	if len(_batchConsumableResource) > 0 {
		input.ConsumableResource = aws.String(_batchConsumableResource)
	}
	if len(_batchFilters) > 0 {
		if err := assignInputField(input, "Filters", _batchFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_batchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _batchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_batchNextToken) > 0 {
		input.NextToken = aws.String(_batchNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListJobsByConsumableResource(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*batch.ListJobsByConsumableResourceOutput
	p := batch.NewListJobsByConsumableResourcePaginator(client, input)
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

// Returns a list of Batch scheduling policies.
func batch_ListSchedulingPolicies(cfg aws.Config, client *batch.Client) {
	input := &batch.ListSchedulingPoliciesInput{}

	if len(_batchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _batchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_batchNextToken) > 0 {
		input.NextToken = aws.String(_batchNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSchedulingPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*batch.ListSchedulingPoliciesOutput
	p := batch.NewListSchedulingPoliciesPaginator(client, input)
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

// Returns a list of service jobs for a specified job queue.
func batch_ListServiceJobs(cfg aws.Config, client *batch.Client) {
	input := &batch.ListServiceJobsInput{}

	if len(_batchFilters) > 0 {
		if err := assignInputField(input, "Filters", _batchFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_batchJobQueue) > 0 {
		input.JobQueue = aws.String(_batchJobQueue)
	}
	if len(_batchJobStatus) > 0 {
		if err := assignInputField(input, "JobStatus", _batchJobStatus); err != nil {
			log.Errorf("invalid --job-status: %s", err.Error())
			return
		}
	}
	if len(_batchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _batchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_batchNextToken) > 0 {
		input.NextToken = aws.String(_batchNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServiceJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*batch.ListServiceJobsOutput
	p := batch.NewListServiceJobsPaginator(client, input)
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

// Lists the tags for an Batch resource. Batch resources that support tags are
// compute environments, jobs, job definitions, job queues, and scheduling
// policies. ARNs for child jobs of array and multi-node parallel (MNP) jobs aren't
// supported.
func batch_ListTagsForResource(cfg aws.Config, client *batch.Client) {
	input := &batch.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_batchResourceArn) > 0 {
		input.ResourceArn = aws.String(_batchResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers an Batch job definition.
func batch_RegisterJobDefinition(cfg aws.Config, client *batch.Client) {
	input := &batch.RegisterJobDefinitionInput{
		// JobDefinitionName: *string, // Required
		// Type: types.JobDefinitionType, // Required
	}

	if len(_batchJobDefinitionName) > 0 {
		input.JobDefinitionName = aws.String(_batchJobDefinitionName)
	}
	if len(_batchType) > 0 {
		if err := assignInputField(input, "Type", _batchType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_batchConsumableResourceProperties) > 0 {
		if err := assignInputField(input, "ConsumableResourceProperties", _batchConsumableResourceProperties); err != nil {
			log.Errorf("invalid --consumable-resource-properties: %s", err.Error())
			return
		}
	}
	if len(_batchContainerProperties) > 0 {
		if err := assignInputField(input, "ContainerProperties", _batchContainerProperties); err != nil {
			log.Errorf("invalid --container-properties: %s", err.Error())
			return
		}
	}
	if len(_batchEcsProperties) > 0 {
		if err := assignInputField(input, "EcsProperties", _batchEcsProperties); err != nil {
			log.Errorf("invalid --ecs-properties: %s", err.Error())
			return
		}
	}
	if len(_batchEksProperties) > 0 {
		if err := assignInputField(input, "EksProperties", _batchEksProperties); err != nil {
			log.Errorf("invalid --eks-properties: %s", err.Error())
			return
		}
	}
	if len(_batchNodeProperties) > 0 {
		if err := assignInputField(input, "NodeProperties", _batchNodeProperties); err != nil {
			log.Errorf("invalid --node-properties: %s", err.Error())
			return
		}
	}
	if len(_batchParameters) > 0 {
		if err := assignInputField(input, "Parameters", _batchParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_batchPlatformCapabilities) > 0 {
		if err := assignInputField(input, "PlatformCapabilities", _batchPlatformCapabilities); err != nil {
			log.Errorf("invalid --platform-capabilities: %s", err.Error())
			return
		}
	}
	if len(_batchPropagateTags) > 0 {
		if err := assignInputField(input, "PropagateTags", _batchPropagateTags); err != nil {
			log.Errorf("invalid --propagate-tags: %s", err.Error())
			return
		}
	}
	if len(_batchRetryStrategy) > 0 {
		if err := assignInputField(input, "RetryStrategy", _batchRetryStrategy); err != nil {
			log.Errorf("invalid --retry-strategy: %s", err.Error())
			return
		}
	}
	if len(_batchSchedulingPriority) > 0 {
		if err := assignInputField(input, "SchedulingPriority", _batchSchedulingPriority); err != nil {
			log.Errorf("invalid --scheduling-priority: %s", err.Error())
			return
		}
	}
	if len(_batchTags) > 0 {
		if err := assignInputField(input, "Tags", _batchTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_batchTimeout) > 0 {
		if err := assignInputField(input, "Timeout", _batchTimeout); err != nil {
			log.Errorf("invalid --timeout: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterJobDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Submits an Batch job from a job definition. Parameters that are specified
// during SubmitJoboverride parameters defined in the job definition. vCPU and memory
// requirements that are specified in the resourceRequirements objects in the job
// definition are the exception. They can't be overridden this way using the memory
// and vcpus parameters. Rather, you must specify updates to job definition
// parameters in a resourceRequirements object that's included in the
// containerOverrides parameter.
//
// Job queues with a scheduling policy are limited to 500 active share identifiers
// at a time.
//
// Jobs that run on Fargate resources can't be guaranteed to run for more than 14
// days. This is because, after 14 days, Fargate resources might become unavailable
// and job might be terminated.
func batch_SubmitJob(cfg aws.Config, client *batch.Client) {
	input := &batch.SubmitJobInput{
		// JobDefinition: *string, // Required
		// JobName: *string, // Required
		// JobQueue: *string, // Required
	}

	if len(_batchJobDefinition) > 0 {
		input.JobDefinition = aws.String(_batchJobDefinition)
	}
	if len(_batchJobName) > 0 {
		input.JobName = aws.String(_batchJobName)
	}
	if len(_batchJobQueue) > 0 {
		input.JobQueue = aws.String(_batchJobQueue)
	}
	if len(_batchArrayProperties) > 0 {
		if err := assignInputField(input, "ArrayProperties", _batchArrayProperties); err != nil {
			log.Errorf("invalid --array-properties: %s", err.Error())
			return
		}
	}
	if len(_batchConsumableResourcePropertiesOverride) > 0 {
		if err := assignInputField(input, "ConsumableResourcePropertiesOverride", _batchConsumableResourcePropertiesOverride); err != nil {
			log.Errorf("invalid --consumable-resource-properties-override: %s", err.Error())
			return
		}
	}
	if len(_batchContainerOverrides) > 0 {
		if err := assignInputField(input, "ContainerOverrides", _batchContainerOverrides); err != nil {
			log.Errorf("invalid --container-overrides: %s", err.Error())
			return
		}
	}
	if len(_batchDependsOn) > 0 {
		if err := assignInputField(input, "DependsOn", _batchDependsOn); err != nil {
			log.Errorf("invalid --depends-on: %s", err.Error())
			return
		}
	}
	if len(_batchEcsPropertiesOverride) > 0 {
		if err := assignInputField(input, "EcsPropertiesOverride", _batchEcsPropertiesOverride); err != nil {
			log.Errorf("invalid --ecs-properties-override: %s", err.Error())
			return
		}
	}
	if len(_batchEksPropertiesOverride) > 0 {
		if err := assignInputField(input, "EksPropertiesOverride", _batchEksPropertiesOverride); err != nil {
			log.Errorf("invalid --eks-properties-override: %s", err.Error())
			return
		}
	}
	if len(_batchNodeOverrides) > 0 {
		if err := assignInputField(input, "NodeOverrides", _batchNodeOverrides); err != nil {
			log.Errorf("invalid --node-overrides: %s", err.Error())
			return
		}
	}
	if len(_batchParameters) > 0 {
		if err := assignInputField(input, "Parameters", _batchParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_batchPropagateTags) > 0 {
		if err := assignInputField(input, "PropagateTags", _batchPropagateTags); err != nil {
			log.Errorf("invalid --propagate-tags: %s", err.Error())
			return
		}
	}
	if len(_batchRetryStrategy) > 0 {
		if err := assignInputField(input, "RetryStrategy", _batchRetryStrategy); err != nil {
			log.Errorf("invalid --retry-strategy: %s", err.Error())
			return
		}
	}
	if len(_batchSchedulingPriorityOverride) > 0 {
		if err := assignInputField(input, "SchedulingPriorityOverride", _batchSchedulingPriorityOverride); err != nil {
			log.Errorf("invalid --scheduling-priority-override: %s", err.Error())
			return
		}
	}
	if len(_batchShareIdentifier) > 0 {
		input.ShareIdentifier = aws.String(_batchShareIdentifier)
	}
	if len(_batchTags) > 0 {
		if err := assignInputField(input, "Tags", _batchTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_batchTimeout) > 0 {
		if err := assignInputField(input, "Timeout", _batchTimeout); err != nil {
			log.Errorf("invalid --timeout: %s", err.Error())
			return
		}
	}

	if resp, err := client.SubmitJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Submits a service job to a specified job queue to run on SageMaker AI. A
// service job is a unit of work that you submit to Batch for execution on
// SageMaker AI.
func batch_SubmitServiceJob(cfg aws.Config, client *batch.Client) {
	input := &batch.SubmitServiceJobInput{
		// JobName: *string, // Required
		// JobQueue: *string, // Required
		// ServiceJobType: types.ServiceJobType, // Required
		// ServiceRequestPayload: *string, // Required
	}

	if len(_batchJobName) > 0 {
		input.JobName = aws.String(_batchJobName)
	}
	if len(_batchJobQueue) > 0 {
		input.JobQueue = aws.String(_batchJobQueue)
	}
	if len(_batchServiceJobType) > 0 {
		if err := assignInputField(input, "ServiceJobType", _batchServiceJobType); err != nil {
			log.Errorf("invalid --service-job-type: %s", err.Error())
			return
		}
	}
	if len(_batchServiceRequestPayload) > 0 {
		input.ServiceRequestPayload = aws.String(_batchServiceRequestPayload)
	}
	if len(_batchClientToken) > 0 {
		input.ClientToken = aws.String(_batchClientToken)
	}
	if len(_batchRetryStrategy) > 0 {
		if err := assignInputField(input, "RetryStrategy", _batchRetryStrategy); err != nil {
			log.Errorf("invalid --retry-strategy: %s", err.Error())
			return
		}
	}
	if len(_batchSchedulingPriority) > 0 {
		if err := assignInputField(input, "SchedulingPriority", _batchSchedulingPriority); err != nil {
			log.Errorf("invalid --scheduling-priority: %s", err.Error())
			return
		}
	}
	if len(_batchShareIdentifier) > 0 {
		input.ShareIdentifier = aws.String(_batchShareIdentifier)
	}
	if len(_batchTags) > 0 {
		if err := assignInputField(input, "Tags", _batchTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_batchTimeoutConfig) > 0 {
		if err := assignInputField(input, "TimeoutConfig", _batchTimeoutConfig); err != nil {
			log.Errorf("invalid --timeout-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.SubmitServiceJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified tags to a resource with the specified resourceArn . If
// existing tags on a resource aren't specified in the request parameters, they
// aren't changed. When a resource is deleted, the tags that are associated with
// that resource are deleted as well. Batch resources that support tags are compute
// environments, jobs, job definitions, job queues, and scheduling policies. ARNs
// for child jobs of array and multi-node parallel (MNP) jobs aren't supported.
func batch_TagResource(cfg aws.Config, client *batch.Client) {
	input := &batch.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_batchResourceArn) > 0 {
		input.ResourceArn = aws.String(_batchResourceArn)
	}
	if len(_batchTags) > 0 {
		if err := assignInputField(input, "Tags", _batchTags); err != nil {
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

// Terminates a job in a job queue. Jobs that are in the STARTING or RUNNING state
// are terminated, which causes them to transition to FAILED . Jobs that have not
// progressed to the STARTING state are cancelled.
func batch_TerminateJob(cfg aws.Config, client *batch.Client) {
	input := &batch.TerminateJobInput{
		// JobId: *string, // Required
		// Reason: *string, // Required
	}

	if len(_batchJobId) > 0 {
		input.JobId = aws.String(_batchJobId)
	}
	if len(_batchReason) > 0 {
		input.Reason = aws.String(_batchReason)
	}

	if resp, err := client.TerminateJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Terminates a service job in a job queue.
func batch_TerminateServiceJob(cfg aws.Config, client *batch.Client) {
	input := &batch.TerminateServiceJobInput{
		// JobId: *string, // Required
		// Reason: *string, // Required
	}

	if len(_batchJobId) > 0 {
		input.JobId = aws.String(_batchJobId)
	}
	if len(_batchReason) > 0 {
		input.Reason = aws.String(_batchReason)
	}

	if resp, err := client.TerminateServiceJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes specified tags from an Batch resource.
func batch_UntagResource(cfg aws.Config, client *batch.Client) {
	input := &batch.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_batchResourceArn) > 0 {
		input.ResourceArn = aws.String(_batchResourceArn)
	}
	if len(_batchTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _batchTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Batch compute environment.
func batch_UpdateComputeEnvironment(cfg aws.Config, client *batch.Client) {
	input := &batch.UpdateComputeEnvironmentInput{
		// ComputeEnvironment: *string, // Required
	}

	if len(_batchComputeEnvironment) > 0 {
		input.ComputeEnvironment = aws.String(_batchComputeEnvironment)
	}
	if len(_batchComputeResources) > 0 {
		if err := assignInputField(input, "ComputeResources", _batchComputeResources); err != nil {
			log.Errorf("invalid --compute-resources: %s", err.Error())
			return
		}
	}
	if len(_batchContext) > 0 {
		input.Context = aws.String(_batchContext)
	}
	if len(_batchServiceRole) > 0 {
		input.ServiceRole = aws.String(_batchServiceRole)
	}
	if len(_batchState) > 0 {
		if err := assignInputField(input, "State", _batchState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}
	if len(_batchUnmanagedvCpus) > 0 {
		if err := assignInputField(input, "UnmanagedvCpus", _batchUnmanagedvCpus); err != nil {
			log.Errorf("invalid --unmanagedv-cpus: %s", err.Error())
			return
		}
	}
	if len(_batchUpdatePolicy) > 0 {
		if err := assignInputField(input, "UpdatePolicy", _batchUpdatePolicy); err != nil {
			log.Errorf("invalid --update-policy: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateComputeEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a consumable resource.
func batch_UpdateConsumableResource(cfg aws.Config, client *batch.Client) {
	input := &batch.UpdateConsumableResourceInput{
		// ConsumableResource: *string, // Required
	}

	if len(_batchConsumableResource) > 0 {
		input.ConsumableResource = aws.String(_batchConsumableResource)
	}
	if len(_batchClientToken) > 0 {
		input.ClientToken = aws.String(_batchClientToken)
	}
	if len(_batchOperation) > 0 {
		input.Operation = aws.String(_batchOperation)
	}
	if len(_batchQuantity) > 0 {
		if err := assignInputField(input, "Quantity", _batchQuantity); err != nil {
			log.Errorf("invalid --quantity: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateConsumableResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a job queue.
func batch_UpdateJobQueue(cfg aws.Config, client *batch.Client) {
	input := &batch.UpdateJobQueueInput{
		// JobQueue: *string, // Required
	}

	if len(_batchJobQueue) > 0 {
		input.JobQueue = aws.String(_batchJobQueue)
	}
	if len(_batchComputeEnvironmentOrder) > 0 {
		if err := assignInputField(input, "ComputeEnvironmentOrder", _batchComputeEnvironmentOrder); err != nil {
			log.Errorf("invalid --compute-environment-order: %s", err.Error())
			return
		}
	}
	if len(_batchJobStateTimeLimitActions) > 0 {
		if err := assignInputField(input, "JobStateTimeLimitActions", _batchJobStateTimeLimitActions); err != nil {
			log.Errorf("invalid --job-state-time-limit-actions: %s", err.Error())
			return
		}
	}
	if len(_batchPriority) > 0 {
		if err := assignInputField(input, "Priority", _batchPriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_batchSchedulingPolicyArn) > 0 {
		input.SchedulingPolicyArn = aws.String(_batchSchedulingPolicyArn)
	}
	if len(_batchServiceEnvironmentOrder) > 0 {
		if err := assignInputField(input, "ServiceEnvironmentOrder", _batchServiceEnvironmentOrder); err != nil {
			log.Errorf("invalid --service-environment-order: %s", err.Error())
			return
		}
	}
	if len(_batchState) > 0 {
		if err := assignInputField(input, "State", _batchState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateJobQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a scheduling policy.
func batch_UpdateSchedulingPolicy(cfg aws.Config, client *batch.Client) {
	input := &batch.UpdateSchedulingPolicyInput{
		// Arn: *string, // Required
	}

	if len(_batchArn) > 0 {
		input.Arn = aws.String(_batchArn)
	}
	if len(_batchFairsharePolicy) > 0 {
		if err := assignInputField(input, "FairsharePolicy", _batchFairsharePolicy); err != nil {
			log.Errorf("invalid --fairshare-policy: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSchedulingPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a service environment. You can update the state of a service
// environment from ENABLED to DISABLED to prevent new service jobs from being
// placed in the service environment.
func batch_UpdateServiceEnvironment(cfg aws.Config, client *batch.Client) {
	input := &batch.UpdateServiceEnvironmentInput{
		// ServiceEnvironment: *string, // Required
	}

	if len(_batchServiceEnvironment) > 0 {
		input.ServiceEnvironment = aws.String(_batchServiceEnvironment)
	}
	if len(_batchCapacityLimits) > 0 {
		if err := assignInputField(input, "CapacityLimits", _batchCapacityLimits); err != nil {
			log.Errorf("invalid --capacity-limits: %s", err.Error())
			return
		}
	}
	if len(_batchState) > 0 {
		if err := assignInputField(input, "State", _batchState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateServiceEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_batchCmd)
	_batchCmd.Flags().SortFlags = false

	_batchCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_batchCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_batchCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_batchCmd.Flags().StringVarP(&_batchArn, "arn", "", "", "ARN")
	_batchCmd.Flags().StringSliceVarP(&_batchArns, "arns", "", nil, "Arns")
	_batchCmd.Flags().StringVarP(&_batchArrayJobId, "array-job-id", "", "", "Array Job ID")
	_batchCmd.Flags().StringVarP(&_batchArrayProperties, "array-properties", "", "", "Array Properties")
	_batchCmd.Flags().StringVarP(&_batchCapacityLimits, "capacity-limits", "", "", "Capacity Limits")
	_batchCmd.Flags().StringVarP(&_batchClientToken, "client-token", "", "", "Client Token")
	_batchCmd.Flags().StringVarP(&_batchComputeEnvironment, "compute-environment", "", "", "Compute Environment")
	_batchCmd.Flags().StringVarP(&_batchComputeEnvironmentName, "compute-environment-name", "", "", "Compute Environment Name")
	_batchCmd.Flags().StringVarP(&_batchComputeEnvironmentOrder, "compute-environment-order", "", "", "Compute Environment Order")
	_batchCmd.Flags().StringSliceVarP(&_batchComputeEnvironments, "compute-environments", "", nil, "Compute Environments")
	_batchCmd.Flags().StringVarP(&_batchComputeResources, "compute-resources", "", "", "Compute Resources")
	_batchCmd.Flags().StringVarP(&_batchConsumableResource, "consumable-resource", "", "", "Consumable Resource")
	_batchCmd.Flags().StringVarP(&_batchConsumableResourceName, "consumable-resource-name", "", "", "Consumable Resource Name")
	_batchCmd.Flags().StringVarP(&_batchConsumableResourceProperties, "consumable-resource-properties", "", "", "Consumable Resource Properties")
	_batchCmd.Flags().StringVarP(&_batchConsumableResourcePropertiesOverride, "consumable-resource-properties-override", "", "", "Consumable Resource Properties Override")
	_batchCmd.Flags().StringVarP(&_batchContainerOverrides, "container-overrides", "", "", "Container Overrides")
	_batchCmd.Flags().StringVarP(&_batchContainerProperties, "container-properties", "", "", "Container Properties")
	_batchCmd.Flags().StringVarP(&_batchContext, "context", "", "", "Context")
	_batchCmd.Flags().StringVarP(&_batchDependsOn, "depends-on", "", "", "Depends On")
	_batchCmd.Flags().StringVarP(&_batchEcsProperties, "ecs-properties", "", "", "Ecs Properties")
	_batchCmd.Flags().StringVarP(&_batchEcsPropertiesOverride, "ecs-properties-override", "", "", "Ecs Properties Override")
	_batchCmd.Flags().StringVarP(&_batchEksConfiguration, "eks-configuration", "", "", "Eks Configuration")
	_batchCmd.Flags().StringVarP(&_batchEksProperties, "eks-properties", "", "", "Eks Properties")
	_batchCmd.Flags().StringVarP(&_batchEksPropertiesOverride, "eks-properties-override", "", "", "Eks Properties Override")
	_batchCmd.Flags().StringVarP(&_batchFairsharePolicy, "fairshare-policy", "", "", "Fairshare Policy")
	_batchCmd.Flags().StringVarP(&_batchFilters, "filters", "", "", "Filters")
	_batchCmd.Flags().StringVarP(&_batchJobDefinition, "job-definition", "", "", "Job Definition")
	_batchCmd.Flags().StringVarP(&_batchJobDefinitionName, "job-definition-name", "", "", "Job Definition Name")
	_batchCmd.Flags().StringSliceVarP(&_batchJobDefinitions, "job-definitions", "", nil, "Job Definitions")
	_batchCmd.Flags().StringVarP(&_batchJobId, "job-id", "", "", "Job ID")
	_batchCmd.Flags().StringVarP(&_batchJobName, "job-name", "", "", "Job Name")
	_batchCmd.Flags().StringVarP(&_batchJobQueue, "job-queue", "", "", "Job Queue")
	_batchCmd.Flags().StringVarP(&_batchJobQueueName, "job-queue-name", "", "", "Job Queue Name")
	_batchCmd.Flags().StringVarP(&_batchJobQueueType, "job-queue-type", "", "", "Job Queue Type")
	_batchCmd.Flags().StringSliceVarP(&_batchJobQueues, "job-queues", "", nil, "Job Queues")
	_batchCmd.Flags().StringVarP(&_batchJobStateTimeLimitActions, "job-state-time-limit-actions", "", "", "Job State Time Limit Actions")
	_batchCmd.Flags().StringVarP(&_batchJobStatus, "job-status", "", "", "Job Status")
	_batchCmd.Flags().StringSliceVarP(&_batchJobs, "jobs", "", nil, "Jobs")
	_batchCmd.Flags().StringVarP(&_batchMaxResults, "max-results", "", "", "Max Results")
	_batchCmd.Flags().StringVarP(&_batchMultiNodeJobId, "multi-node-job-id", "", "", "Multi Node Job ID")
	_batchCmd.Flags().StringVarP(&_batchName, "name", "", "", "Name")
	_batchCmd.Flags().StringVarP(&_batchNextToken, "next-token", "", "", "Next Token")
	_batchCmd.Flags().StringVarP(&_batchNodeOverrides, "node-overrides", "", "", "Node Overrides")
	_batchCmd.Flags().StringVarP(&_batchNodeProperties, "node-properties", "", "", "Node Properties")
	_batchCmd.Flags().StringVarP(&_batchOperation, "operation", "", "", "Operation")
	_batchCmd.Flags().StringVarP(&_batchParameters, "parameters", "", "", "Parameters")
	_batchCmd.Flags().StringVarP(&_batchPlatformCapabilities, "platform-capabilities", "", "", "Platform Capabilities")
	_batchCmd.Flags().StringVarP(&_batchPriority, "priority", "", "", "Priority")
	_batchCmd.Flags().StringVarP(&_batchPropagateTags, "propagate-tags", "", "", "Propagate Tags")
	_batchCmd.Flags().StringVarP(&_batchQuantity, "quantity", "", "", "Quantity")
	_batchCmd.Flags().StringVarP(&_batchReason, "reason", "", "", "Reason")
	_batchCmd.Flags().StringVarP(&_batchResourceArn, "resource-arn", "", "", "Resource ARN")
	_batchCmd.Flags().StringVarP(&_batchResourceType, "resource-type", "", "", "Resource Type")
	_batchCmd.Flags().StringVarP(&_batchRetryStrategy, "retry-strategy", "", "", "Retry Strategy")
	_batchCmd.Flags().StringVarP(&_batchSchedulingPolicyArn, "scheduling-policy-arn", "", "", "Scheduling Policy ARN")
	_batchCmd.Flags().StringVarP(&_batchSchedulingPriority, "scheduling-priority", "", "", "Scheduling Priority")
	_batchCmd.Flags().StringVarP(&_batchSchedulingPriorityOverride, "scheduling-priority-override", "", "", "Scheduling Priority Override")
	_batchCmd.Flags().StringVarP(&_batchServiceEnvironment, "service-environment", "", "", "Service Environment")
	_batchCmd.Flags().StringVarP(&_batchServiceEnvironmentName, "service-environment-name", "", "", "Service Environment Name")
	_batchCmd.Flags().StringVarP(&_batchServiceEnvironmentOrder, "service-environment-order", "", "", "Service Environment Order")
	_batchCmd.Flags().StringVarP(&_batchServiceEnvironmentType, "service-environment-type", "", "", "Service Environment Type")
	_batchCmd.Flags().StringSliceVarP(&_batchServiceEnvironments, "service-environments", "", nil, "Service Environments")
	_batchCmd.Flags().StringVarP(&_batchServiceJobType, "service-job-type", "", "", "Service Job Type")
	_batchCmd.Flags().StringVarP(&_batchServiceRequestPayload, "service-request-payload", "", "", "Service Request Payload")
	_batchCmd.Flags().StringVarP(&_batchServiceRole, "service-role", "", "", "Service Role")
	_batchCmd.Flags().StringVarP(&_batchShareIdentifier, "share-identifier", "", "", "Share Identifier")
	_batchCmd.Flags().StringVarP(&_batchState, "state", "", "", "State")
	_batchCmd.Flags().StringVarP(&_batchStatus, "status", "", "", "Status")
	_batchCmd.Flags().StringSliceVarP(&_batchTagKeys, "tag-keys", "", nil, "Tag Keys")
	_batchCmd.Flags().StringVarP(&_batchTags, "tags", "", "", "Tags")
	_batchCmd.Flags().StringVarP(&_batchTimeout, "timeout", "", "", "Timeout")
	_batchCmd.Flags().StringVarP(&_batchTimeoutConfig, "timeout-config", "", "", "Timeout Config")
	_batchCmd.Flags().StringVarP(&_batchTotalQuantity, "total-quantity", "", "", "Total Quantity")
	_batchCmd.Flags().StringVarP(&_batchType, "type", "", "", "Type")
	_batchCmd.Flags().StringVarP(&_batchUnmanagedvCpus, "unmanagedv-cpus", "", "", "Unmanagedv Cpus")
	_batchCmd.Flags().StringVarP(&_batchUpdatePolicy, "update-policy", "", "", "Update Policy")

	_batchCmd.Flags().BoolVarP(&_batchCancelJob, "cancel-job", "", false, "Cancel Job")
	_batchCmd.Flags().BoolVarP(&_batchCreateComputeEnvironment, "create-compute-environment", "", false, "Create Compute Environment")
	_batchCmd.Flags().BoolVarP(&_batchCreateConsumableResource, "create-consumable-resource", "", false, "Create Consumable Resource")
	_batchCmd.Flags().BoolVarP(&_batchCreateJobQueue, "create-job-queue", "", false, "Create Job Queue")
	_batchCmd.Flags().BoolVarP(&_batchCreateSchedulingPolicy, "create-scheduling-policy", "", false, "Create Scheduling Policy")
	_batchCmd.Flags().BoolVarP(&_batchCreateServiceEnvironment, "create-service-environment", "", false, "Create Service Environment")
	_batchCmd.Flags().BoolVarP(&_batchDeleteComputeEnvironment, "delete-compute-environment", "", false, "Delete Compute Environment")
	_batchCmd.Flags().BoolVarP(&_batchDeleteConsumableResource, "delete-consumable-resource", "", false, "Delete Consumable Resource")
	_batchCmd.Flags().BoolVarP(&_batchDeleteJobQueue, "delete-job-queue", "", false, "Delete Job Queue")
	_batchCmd.Flags().BoolVarP(&_batchDeleteSchedulingPolicy, "delete-scheduling-policy", "", false, "Delete Scheduling Policy")
	_batchCmd.Flags().BoolVarP(&_batchDeleteServiceEnvironment, "delete-service-environment", "", false, "Delete Service Environment")
	_batchCmd.Flags().BoolVarP(&_batchDeregisterJobDefinition, "deregister-job-definition", "", false, "Deregister Job Definition")
	_batchCmd.Flags().BoolVarP(&_batchDescribeComputeEnvironments, "describe-compute-environments", "", false, "Describe Compute Environments")
	_batchCmd.Flags().BoolVarP(&_batchDescribeConsumableResource, "describe-consumable-resource", "", false, "Describe Consumable Resource")
	_batchCmd.Flags().BoolVarP(&_batchDescribeJobDefinitions, "describe-job-definitions", "", false, "Describe Job Definitions")
	_batchCmd.Flags().BoolVarP(&_batchDescribeJobQueues, "describe-job-queues", "", false, "Describe Job Queues")
	_batchCmd.Flags().BoolVarP(&_batchDescribeJobs, "describe-jobs", "", false, "Describe Jobs")
	_batchCmd.Flags().BoolVarP(&_batchDescribeSchedulingPolicies, "describe-scheduling-policies", "", false, "Describe Scheduling Policies")
	_batchCmd.Flags().BoolVarP(&_batchDescribeServiceEnvironments, "describe-service-environments", "", false, "Describe Service Environments")
	_batchCmd.Flags().BoolVarP(&_batchDescribeServiceJob, "describe-service-job", "", false, "Describe Service Job")
	_batchCmd.Flags().BoolVarP(&_batchGetJobQueueSnapshot, "get-job-queue-snapshot", "", false, "Get Job Queue Snapshot")
	_batchCmd.Flags().BoolVarP(&_batchListConsumableResources, "list-consumable-resources", "", false, "List Consumable Resources")
	_batchCmd.Flags().BoolVarP(&_batchListJobs, "list-jobs", "", false, "List Jobs")
	_batchCmd.Flags().BoolVarP(&_batchListJobsByConsumableResource, "list-jobs-by-consumable-resource", "", false, "List Jobs By Consumable Resource")
	_batchCmd.Flags().BoolVarP(&_batchListSchedulingPolicies, "list-scheduling-policies", "", false, "List Scheduling Policies")
	_batchCmd.Flags().BoolVarP(&_batchListServiceJobs, "list-service-jobs", "", false, "List Service Jobs")
	_batchCmd.Flags().BoolVarP(&_batchListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_batchCmd.Flags().BoolVarP(&_batchRegisterJobDefinition, "register-job-definition", "", false, "Register Job Definition")
	_batchCmd.Flags().BoolVarP(&_batchSubmitJob, "submit-job", "", false, "Submit Job")
	_batchCmd.Flags().BoolVarP(&_batchSubmitServiceJob, "submit-service-job", "", false, "Submit Service Job")
	_batchCmd.Flags().BoolVarP(&_batchTagResource, "tag-resource", "", false, "Tag Resource")
	_batchCmd.Flags().BoolVarP(&_batchTerminateJob, "terminate-job", "", false, "Terminate Job")
	_batchCmd.Flags().BoolVarP(&_batchTerminateServiceJob, "terminate-service-job", "", false, "Terminate Service Job")
	_batchCmd.Flags().BoolVarP(&_batchUntagResource, "untag-resource", "", false, "Untag Resource")
	_batchCmd.Flags().BoolVarP(&_batchUpdateComputeEnvironment, "update-compute-environment", "", false, "Update Compute Environment")
	_batchCmd.Flags().BoolVarP(&_batchUpdateConsumableResource, "update-consumable-resource", "", false, "Update Consumable Resource")
	_batchCmd.Flags().BoolVarP(&_batchUpdateJobQueue, "update-job-queue", "", false, "Update Job Queue")
	_batchCmd.Flags().BoolVarP(&_batchUpdateSchedulingPolicy, "update-scheduling-policy", "", false, "Update Scheduling Policy")
	_batchCmd.Flags().BoolVarP(&_batchUpdateServiceEnvironment, "update-service-environment", "", false, "Update Service Environment")

}
