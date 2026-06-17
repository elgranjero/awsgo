package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// autoscalingCmd represents the autoscaling command
var _autoscalingCmd = &cobra.Command{
	Use:   "autoscaling",
	Short: "AWS autoscaling CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := autoscaling.NewFromConfig(cfg)
		if _autoscalingAttachInstances {
			autoscaling_AttachInstances(cfg, client)
			return
		}
		if _autoscalingAttachLoadBalancerTargetGroups {
			autoscaling_AttachLoadBalancerTargetGroups(cfg, client)
			return
		}
		if _autoscalingAttachLoadBalancers {
			autoscaling_AttachLoadBalancers(cfg, client)
			return
		}
		if _autoscalingAttachTrafficSources {
			autoscaling_AttachTrafficSources(cfg, client)
			return
		}
		if _autoscalingBatchDeleteScheduledAction {
			autoscaling_BatchDeleteScheduledAction(cfg, client)
			return
		}
		if _autoscalingBatchPutScheduledUpdateGroupAction {
			autoscaling_BatchPutScheduledUpdateGroupAction(cfg, client)
			return
		}
		if _autoscalingCancelInstanceRefresh {
			autoscaling_CancelInstanceRefresh(cfg, client)
			return
		}
		if _autoscalingCompleteLifecycleAction {
			autoscaling_CompleteLifecycleAction(cfg, client)
			return
		}
		if _autoscalingCreateAutoScalingGroup {
			autoscaling_CreateAutoScalingGroup(cfg, client)
			return
		}
		if _autoscalingCreateLaunchConfiguration {
			autoscaling_CreateLaunchConfiguration(cfg, client)
			return
		}
		if _autoscalingCreateOrUpdateTags {
			autoscaling_CreateOrUpdateTags(cfg, client)
			return
		}
		if _autoscalingDeleteAutoScalingGroup {
			autoscaling_DeleteAutoScalingGroup(cfg, client)
			return
		}
		if _autoscalingDeleteLaunchConfiguration {
			autoscaling_DeleteLaunchConfiguration(cfg, client)
			return
		}
		if _autoscalingDeleteLifecycleHook {
			autoscaling_DeleteLifecycleHook(cfg, client)
			return
		}
		if _autoscalingDeleteNotificationConfiguration {
			autoscaling_DeleteNotificationConfiguration(cfg, client)
			return
		}
		if _autoscalingDeletePolicy {
			autoscaling_DeletePolicy(cfg, client)
			return
		}
		if _autoscalingDeleteScheduledAction {
			autoscaling_DeleteScheduledAction(cfg, client)
			return
		}
		if _autoscalingDeleteTags {
			autoscaling_DeleteTags(cfg, client)
			return
		}
		if _autoscalingDeleteWarmPool {
			autoscaling_DeleteWarmPool(cfg, client)
			return
		}
		if _autoscalingDescribeAccountLimits {
			autoscaling_DescribeAccountLimits(cfg, client)
			return
		}
		if _autoscalingDescribeAdjustmentTypes {
			autoscaling_DescribeAdjustmentTypes(cfg, client)
			return
		}
		if _autoscalingDescribeAutoScalingGroups {
			autoscaling_DescribeAutoScalingGroups(cfg, client)
			return
		}
		if _autoscalingDescribeAutoScalingInstances {
			autoscaling_DescribeAutoScalingInstances(cfg, client)
			return
		}
		if _autoscalingDescribeAutoScalingNotificationTypes {
			autoscaling_DescribeAutoScalingNotificationTypes(cfg, client)
			return
		}
		if _autoscalingDescribeInstanceRefreshes {
			autoscaling_DescribeInstanceRefreshes(cfg, client)
			return
		}
		if _autoscalingDescribeLaunchConfigurations {
			autoscaling_DescribeLaunchConfigurations(cfg, client)
			return
		}
		if _autoscalingDescribeLifecycleHookTypes {
			autoscaling_DescribeLifecycleHookTypes(cfg, client)
			return
		}
		if _autoscalingDescribeLifecycleHooks {
			autoscaling_DescribeLifecycleHooks(cfg, client)
			return
		}
		if _autoscalingDescribeLoadBalancerTargetGroups {
			autoscaling_DescribeLoadBalancerTargetGroups(cfg, client)
			return
		}
		if _autoscalingDescribeLoadBalancers {
			autoscaling_DescribeLoadBalancers(cfg, client)
			return
		}
		if _autoscalingDescribeMetricCollectionTypes {
			autoscaling_DescribeMetricCollectionTypes(cfg, client)
			return
		}
		if _autoscalingDescribeNotificationConfigurations {
			autoscaling_DescribeNotificationConfigurations(cfg, client)
			return
		}
		if _autoscalingDescribePolicies {
			autoscaling_DescribePolicies(cfg, client)
			return
		}
		if _autoscalingDescribeScalingActivities {
			autoscaling_DescribeScalingActivities(cfg, client)
			return
		}
		if _autoscalingDescribeScalingProcessTypes {
			autoscaling_DescribeScalingProcessTypes(cfg, client)
			return
		}
		if _autoscalingDescribeScheduledActions {
			autoscaling_DescribeScheduledActions(cfg, client)
			return
		}
		if _autoscalingDescribeTags {
			autoscaling_DescribeTags(cfg, client)
			return
		}
		if _autoscalingDescribeTerminationPolicyTypes {
			autoscaling_DescribeTerminationPolicyTypes(cfg, client)
			return
		}
		if _autoscalingDescribeTrafficSources {
			autoscaling_DescribeTrafficSources(cfg, client)
			return
		}
		if _autoscalingDescribeWarmPool {
			autoscaling_DescribeWarmPool(cfg, client)
			return
		}
		if _autoscalingDetachInstances {
			autoscaling_DetachInstances(cfg, client)
			return
		}
		if _autoscalingDetachLoadBalancerTargetGroups {
			autoscaling_DetachLoadBalancerTargetGroups(cfg, client)
			return
		}
		if _autoscalingDetachLoadBalancers {
			autoscaling_DetachLoadBalancers(cfg, client)
			return
		}
		if _autoscalingDetachTrafficSources {
			autoscaling_DetachTrafficSources(cfg, client)
			return
		}
		if _autoscalingDisableMetricsCollection {
			autoscaling_DisableMetricsCollection(cfg, client)
			return
		}
		if _autoscalingEnableMetricsCollection {
			autoscaling_EnableMetricsCollection(cfg, client)
			return
		}
		if _autoscalingEnterStandby {
			autoscaling_EnterStandby(cfg, client)
			return
		}
		if _autoscalingExecutePolicy {
			autoscaling_ExecutePolicy(cfg, client)
			return
		}
		if _autoscalingExitStandby {
			autoscaling_ExitStandby(cfg, client)
			return
		}
		if _autoscalingGetPredictiveScalingForecast {
			autoscaling_GetPredictiveScalingForecast(cfg, client)
			return
		}
		if _autoscalingLaunchInstances {
			autoscaling_LaunchInstances(cfg, client)
			return
		}
		if _autoscalingPutLifecycleHook {
			autoscaling_PutLifecycleHook(cfg, client)
			return
		}
		if _autoscalingPutNotificationConfiguration {
			autoscaling_PutNotificationConfiguration(cfg, client)
			return
		}
		if _autoscalingPutScalingPolicy {
			autoscaling_PutScalingPolicy(cfg, client)
			return
		}
		if _autoscalingPutScheduledUpdateGroupAction {
			autoscaling_PutScheduledUpdateGroupAction(cfg, client)
			return
		}
		if _autoscalingPutWarmPool {
			autoscaling_PutWarmPool(cfg, client)
			return
		}
		if _autoscalingRecordLifecycleActionHeartbeat {
			autoscaling_RecordLifecycleActionHeartbeat(cfg, client)
			return
		}
		if _autoscalingResumeProcesses {
			autoscaling_ResumeProcesses(cfg, client)
			return
		}
		if _autoscalingRollbackInstanceRefresh {
			autoscaling_RollbackInstanceRefresh(cfg, client)
			return
		}
		if _autoscalingSetDesiredCapacity {
			autoscaling_SetDesiredCapacity(cfg, client)
			return
		}
		if _autoscalingSetInstanceHealth {
			autoscaling_SetInstanceHealth(cfg, client)
			return
		}
		if _autoscalingSetInstanceProtection {
			autoscaling_SetInstanceProtection(cfg, client)
			return
		}
		if _autoscalingStartInstanceRefresh {
			autoscaling_StartInstanceRefresh(cfg, client)
			return
		}
		if _autoscalingSuspendProcesses {
			autoscaling_SuspendProcesses(cfg, client)
			return
		}
		if _autoscalingTerminateInstanceInAutoScalingGroup {
			autoscaling_TerminateInstanceInAutoScalingGroup(cfg, client)
			return
		}
		if _autoscalingUpdateAutoScalingGroup {
			autoscaling_UpdateAutoScalingGroup(cfg, client)
			return
		}

	},
}

var (
	_autoscalingAttachInstances                      bool
	_autoscalingAttachLoadBalancerTargetGroups       bool
	_autoscalingAttachLoadBalancers                  bool
	_autoscalingAttachTrafficSources                 bool
	_autoscalingBatchDeleteScheduledAction           bool
	_autoscalingBatchPutScheduledUpdateGroupAction   bool
	_autoscalingCancelInstanceRefresh                bool
	_autoscalingCompleteLifecycleAction              bool
	_autoscalingCreateAutoScalingGroup               bool
	_autoscalingCreateLaunchConfiguration            bool
	_autoscalingCreateOrUpdateTags                   bool
	_autoscalingDeleteAutoScalingGroup               bool
	_autoscalingDeleteLaunchConfiguration            bool
	_autoscalingDeleteLifecycleHook                  bool
	_autoscalingDeleteNotificationConfiguration      bool
	_autoscalingDeletePolicy                         bool
	_autoscalingDeleteScheduledAction                bool
	_autoscalingDeleteTags                           bool
	_autoscalingDeleteWarmPool                       bool
	_autoscalingDescribeAccountLimits                bool
	_autoscalingDescribeAdjustmentTypes              bool
	_autoscalingDescribeAutoScalingGroups            bool
	_autoscalingDescribeAutoScalingInstances         bool
	_autoscalingDescribeAutoScalingNotificationTypes bool
	_autoscalingDescribeInstanceRefreshes            bool
	_autoscalingDescribeLaunchConfigurations         bool
	_autoscalingDescribeLifecycleHookTypes           bool
	_autoscalingDescribeLifecycleHooks               bool
	_autoscalingDescribeLoadBalancerTargetGroups     bool
	_autoscalingDescribeLoadBalancers                bool
	_autoscalingDescribeMetricCollectionTypes        bool
	_autoscalingDescribeNotificationConfigurations   bool
	_autoscalingDescribePolicies                     bool
	_autoscalingDescribeScalingActivities            bool
	_autoscalingDescribeScalingProcessTypes          bool
	_autoscalingDescribeScheduledActions             bool
	_autoscalingDescribeTags                         bool
	_autoscalingDescribeTerminationPolicyTypes       bool
	_autoscalingDescribeTrafficSources               bool
	_autoscalingDescribeWarmPool                     bool
	_autoscalingDetachInstances                      bool
	_autoscalingDetachLoadBalancerTargetGroups       bool
	_autoscalingDetachLoadBalancers                  bool
	_autoscalingDetachTrafficSources                 bool
	_autoscalingDisableMetricsCollection             bool
	_autoscalingEnableMetricsCollection              bool
	_autoscalingEnterStandby                         bool
	_autoscalingExecutePolicy                        bool
	_autoscalingExitStandby                          bool
	_autoscalingGetPredictiveScalingForecast         bool
	_autoscalingLaunchInstances                      bool
	_autoscalingPutLifecycleHook                     bool
	_autoscalingPutNotificationConfiguration         bool
	_autoscalingPutScalingPolicy                     bool
	_autoscalingPutScheduledUpdateGroupAction        bool
	_autoscalingPutWarmPool                          bool
	_autoscalingRecordLifecycleActionHeartbeat       bool
	_autoscalingResumeProcesses                      bool
	_autoscalingRollbackInstanceRefresh              bool
	_autoscalingSetDesiredCapacity                   bool
	_autoscalingSetInstanceHealth                    bool
	_autoscalingSetInstanceProtection                bool
	_autoscalingStartInstanceRefresh                 bool
	_autoscalingSuspendProcesses                     bool
	_autoscalingTerminateInstanceInAutoScalingGroup  bool
	_autoscalingUpdateAutoScalingGroup               bool

	_autoscalingActivityIds                      []string
	_autoscalingAdjustmentType                   string
	_autoscalingAssociatePublicIpAddress         string
	_autoscalingAutoScalingGroupName             string
	_autoscalingAutoScalingGroupNames            []string
	_autoscalingAvailabilityZoneDistribution     string
	_autoscalingAvailabilityZoneIds              []string
	_autoscalingAvailabilityZoneImpairmentPolicy string
	_autoscalingAvailabilityZones                []string
	_autoscalingBlockDeviceMappings              string
	_autoscalingBreachThreshold                  string
	_autoscalingCapacityRebalance                string
	_autoscalingCapacityReservationSpecification string
	_autoscalingClassicLinkVPCSecurityGroups     []string
	_autoscalingClassicLinkVPCId                 string
	_autoscalingClientToken                      string
	_autoscalingContext                          string
	_autoscalingCooldown                         string
	_autoscalingDefaultCooldown                  string
	_autoscalingDefaultInstanceWarmup            string
	_autoscalingDefaultResult                    string
	_autoscalingDeletionProtection               string
	_autoscalingDesiredCapacity                  string
	_autoscalingDesiredCapacityType              string
	_autoscalingDesiredConfiguration             string
	_autoscalingEbsOptimized                     string
	_autoscalingEnabled                          string
	_autoscalingEndTime                          string
	_autoscalingEstimatedInstanceWarmup          string
	_autoscalingFilters                          string
	_autoscalingForceDelete                      string
	_autoscalingGranularity                      string
	_autoscalingHealthCheckGracePeriod           string
	_autoscalingHealthCheckType                  string
	_autoscalingHealthStatus                     string
	_autoscalingHeartbeatTimeout                 string
	_autoscalingHonorCooldown                    string
	_autoscalingIamInstanceProfile               string
	_autoscalingImageId                          string
	_autoscalingIncludeDeletedGroups             string
	_autoscalingIncludeInstances                 string
	_autoscalingInstanceId                       string
	_autoscalingInstanceIds                      []string
	_autoscalingInstanceLifecyclePolicy          string
	_autoscalingInstanceMaintenancePolicy        string
	_autoscalingInstanceMonitoring               string
	_autoscalingInstanceRefreshIds               []string
	_autoscalingInstanceReusePolicy              string
	_autoscalingInstanceType                     string
	_autoscalingKernelId                         string
	_autoscalingKeyName                          string
	_autoscalingLaunchConfigurationName          string
	_autoscalingLaunchConfigurationNames         []string
	_autoscalingLaunchTemplate                   string
	_autoscalingLifecycleActionResult            string
	_autoscalingLifecycleActionToken             string
	_autoscalingLifecycleHookName                string
	_autoscalingLifecycleHookNames               []string
	_autoscalingLifecycleHookSpecificationList   string
	_autoscalingLifecycleTransition              string
	_autoscalingLoadBalancerNames                []string
	_autoscalingMaxGroupPreparedCapacity         string
	_autoscalingMaxInstanceLifetime              string
	_autoscalingMaxRecords                       string
	_autoscalingMaxSize                          string
	_autoscalingMetadataOptions                  string
	_autoscalingMetricAggregationType            string
	_autoscalingMetricValue                      string
	_autoscalingMetrics                          []string
	_autoscalingMinAdjustmentMagnitude           string
	_autoscalingMinAdjustmentStep                string
	_autoscalingMinSize                          string
	_autoscalingMixedInstancesPolicy             string
	_autoscalingNewInstancesProtectedFromScaleIn string
	_autoscalingNextToken                        string
	_autoscalingNotificationMetadata             string
	_autoscalingNotificationTargetARN            string
	_autoscalingNotificationTypes                []string
	_autoscalingPlacementGroup                   string
	_autoscalingPlacementTenancy                 string
	_autoscalingPolicyName                       string
	_autoscalingPolicyNames                      []string
	_autoscalingPolicyType                       string
	_autoscalingPolicyTypes                      []string
	_autoscalingPoolState                        string
	_autoscalingPredictiveScalingConfiguration   string
	_autoscalingPreferences                      string
	_autoscalingProtectedFromScaleIn             string
	_autoscalingRamdiskId                        string
	_autoscalingRecurrence                       string
	_autoscalingRequestedCapacity                string
	_autoscalingRetryStrategy                    string
	_autoscalingRoleARN                          string
	_autoscalingScalingAdjustment                string
	_autoscalingScalingProcesses                 []string
	_autoscalingScheduledActionName              string
	_autoscalingScheduledActionNames             []string
	_autoscalingScheduledUpdateGroupActions      string
	_autoscalingSecurityGroups                   []string
	_autoscalingServiceLinkedRoleARN             string
	_autoscalingShouldDecrementDesiredCapacity   string
	_autoscalingShouldRespectGracePeriod         string
	_autoscalingSkipZonalShiftValidation         string
	_autoscalingSpotPrice                        string
	_autoscalingStartTime                        string
	_autoscalingStepAdjustments                  string
	_autoscalingStrategy                         string
	_autoscalingSubnetIds                        []string
	_autoscalingTags                             string
	_autoscalingTargetGroupARNs                  []string
	_autoscalingTargetTrackingConfiguration      string
	_autoscalingTerminationPolicies              []string
	_autoscalingTime                             string
	_autoscalingTimeZone                         string
	_autoscalingTopicARN                         string
	_autoscalingTrafficSourceType                string
	_autoscalingTrafficSources                   string
	_autoscalingUserData                         string
	_autoscalingVPCZoneIdentifier                string
	_autoscalingWaitForTransitioningInstances    string
)

// Attaches one or more EC2 instances to the specified Auto Scaling group.
// When you attach instances, Amazon EC2 Auto Scaling increases the desired
// capacity of the group by the number of instances being attached. If the number
// of instances being attached plus the desired capacity of the group exceeds the
// maximum size of the group, the operation fails.
//
// If there is a Classic Load Balancer attached to your Auto Scaling group, the
// instances are also registered with the load balancer. If there are target groups
// attached to your Auto Scaling group, the instances are also registered with the
// target groups.
//
// For more information, see [Detach or attach instances] in the Amazon EC2 Auto Scaling User Guide.
//
// [Detach or attach instances]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-detach-attach-instances.html
func autoscaling_AttachInstances(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.AttachInstancesInput{
		// AutoScalingGroupName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingInstanceIds) > 0 {
		input.InstanceIds = append([]string(nil), _autoscalingInstanceIds...)
	}

	if resp, err := client.AttachInstances(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API operation is superseded by [AttachTrafficSources], which can attach multiple traffic sources
// types. We recommend using AttachTrafficSources to simplify how you manage
// traffic sources. However, we continue to support AttachLoadBalancerTargetGroups
// . You can use both the original AttachLoadBalancerTargetGroups API operation
// and AttachTrafficSources on the same Auto Scaling group.
//
// Attaches one or more target groups to the specified Auto Scaling group.
//
// This operation is used with the following load balancer types:
//
// - Application Load Balancer - Operates at the application layer (layer 7) and
// supports HTTP and HTTPS.
//
// - Network Load Balancer - Operates at the transport layer (layer 4) and
// supports TCP, TLS, and UDP.
//
// - Gateway Load Balancer - Operates at the network layer (layer 3).
//
// To describe the target groups for an Auto Scaling group, call the [DescribeLoadBalancerTargetGroups] API. To
// detach the target group from the Auto Scaling group, call the [DetachLoadBalancerTargetGroups]API.
//
// This operation is additive and does not detach existing target groups or
// Classic Load Balancers from the Auto Scaling group.
//
// For more information, see [Use Elastic Load Balancing to distribute traffic across the instances in your Auto Scaling group] in the Amazon EC2 Auto Scaling User Guide.
//
// [DescribeLoadBalancerTargetGroups]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeLoadBalancerTargetGroups.html
// [DetachLoadBalancerTargetGroups]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DetachLoadBalancerTargetGroups.html
// [Use Elastic Load Balancing to distribute traffic across the instances in your Auto Scaling group]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-load-balancer.html
// [AttachTrafficSources]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_AttachTrafficSources.html
func autoscaling_AttachLoadBalancerTargetGroups(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.AttachLoadBalancerTargetGroupsInput{
		// AutoScalingGroupName: *string, // Required
		// TargetGroupARNs: []string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingTargetGroupARNs) > 0 {
		input.TargetGroupARNs = append([]string(nil), _autoscalingTargetGroupARNs...)
	}

	if resp, err := client.AttachLoadBalancerTargetGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API operation is superseded by [AttachTrafficSources], which can attach multiple traffic sources
// types. We recommend using AttachTrafficSources to simplify how you manage
// traffic sources. However, we continue to support AttachLoadBalancers . You can
// use both the original AttachLoadBalancers API operation and AttachTrafficSources
// on the same Auto Scaling group.
//
// Attaches one or more Classic Load Balancers to the specified Auto Scaling
// group. Amazon EC2 Auto Scaling registers the running instances with these
// Classic Load Balancers.
//
// To describe the load balancers for an Auto Scaling group, call the [DescribeLoadBalancers] API. To
// detach a load balancer from the Auto Scaling group, call the [DetachLoadBalancers]API.
//
// This operation is additive and does not detach existing Classic Load Balancers
// or target groups from the Auto Scaling group.
//
// For more information, see [Use Elastic Load Balancing to distribute traffic across the instances in your Auto Scaling group] in the Amazon EC2 Auto Scaling User Guide.
//
// [DetachLoadBalancers]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DetachLoadBalancers.html
// [DescribeLoadBalancers]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeLoadBalancers.html
// [Use Elastic Load Balancing to distribute traffic across the instances in your Auto Scaling group]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-load-balancer.html
// [AttachTrafficSources]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_AttachTrafficSources.html
func autoscaling_AttachLoadBalancers(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.AttachLoadBalancersInput{
		// AutoScalingGroupName: *string, // Required
		// LoadBalancerNames: []string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingLoadBalancerNames) > 0 {
		input.LoadBalancerNames = append([]string(nil), _autoscalingLoadBalancerNames...)
	}

	if resp, err := client.AttachLoadBalancers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches one or more traffic sources to the specified Auto Scaling group.
// You can use any of the following as traffic sources for an Auto Scaling group:
//
// - Application Load Balancer
//
// - Classic Load Balancer
//
// - Gateway Load Balancer
//
// - Network Load Balancer
//
// - VPC Lattice
//
// This operation is additive and does not detach existing traffic sources from
// the Auto Scaling group.
//
// After the operation completes, use the [DescribeTrafficSources] API to return details about the state
// of the attachments between traffic sources and your Auto Scaling group. To
// detach a traffic source from the Auto Scaling group, call the [DetachTrafficSources]API.
//
// [DescribeTrafficSources]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeTrafficSources.html
// [DetachTrafficSources]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DetachTrafficSources.html
func autoscaling_AttachTrafficSources(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.AttachTrafficSourcesInput{
		// AutoScalingGroupName: *string, // Required
		// TrafficSources: []types.TrafficSourceIdentifier, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingTrafficSources) > 0 {
		if err := assignInputField(input, "TrafficSources", _autoscalingTrafficSources); err != nil {
			log.Errorf("invalid --traffic-sources: %s", err.Error())
			return
		}
	}
	if len(_autoscalingSkipZonalShiftValidation) > 0 {
		if err := assignInputField(input, "SkipZonalShiftValidation", _autoscalingSkipZonalShiftValidation); err != nil {
			log.Errorf("invalid --skip-zonal-shift-validation: %s", err.Error())
			return
		}
	}

	if resp, err := client.AttachTrafficSources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes one or more scheduled actions for the specified Auto Scaling group.
func autoscaling_BatchDeleteScheduledAction(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.BatchDeleteScheduledActionInput{
		// AutoScalingGroupName: *string, // Required
		// ScheduledActionNames: []string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingScheduledActionNames) > 0 {
		input.ScheduledActionNames = append([]string(nil), _autoscalingScheduledActionNames...)
	}

	if resp, err := client.BatchDeleteScheduledAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates one or more scheduled scaling actions for an Auto Scaling
// group.
func autoscaling_BatchPutScheduledUpdateGroupAction(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.BatchPutScheduledUpdateGroupActionInput{
		// AutoScalingGroupName: *string, // Required
		// ScheduledUpdateGroupActions: []types.ScheduledUpdateGroupActionRequest, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingScheduledUpdateGroupActions) > 0 {
		if err := assignInputField(input, "ScheduledUpdateGroupActions", _autoscalingScheduledUpdateGroupActions); err != nil {
			log.Errorf("invalid --scheduled-update-group-actions: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchPutScheduledUpdateGroupAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels an instance refresh or rollback that is in progress. If an instance
// refresh or rollback is not in progress, an ActiveInstanceRefreshNotFound error
// occurs.
//
// This operation is part of the [instance refresh feature] in Amazon EC2 Auto Scaling, which helps you
// update instances in your Auto Scaling group after you make configuration
// changes.
//
// When you cancel an instance refresh, this does not roll back any changes that
// it made. Use the [RollbackInstanceRefresh]API to roll back instead.
//
// [instance refresh feature]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-refresh.html
// [RollbackInstanceRefresh]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_RollbackInstanceRefresh.html
func autoscaling_CancelInstanceRefresh(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.CancelInstanceRefreshInput{
		// AutoScalingGroupName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingWaitForTransitioningInstances) > 0 {
		if err := assignInputField(input, "WaitForTransitioningInstances", _autoscalingWaitForTransitioningInstances); err != nil {
			log.Errorf("invalid --wait-for-transitioning-instances: %s", err.Error())
			return
		}
	}

	if resp, err := client.CancelInstanceRefresh(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Completes the lifecycle action for the specified token or instance with the
// specified result.
//
// This step is a part of the procedure for adding a lifecycle hook to an Auto
// Scaling group:
//
// - (Optional) Create a launch template or launch configuration with a user
// data script that runs while an instance is in a wait state due to a lifecycle
// hook.
//
// - (Optional) Create a Lambda function and a rule that allows Amazon
// EventBridge to invoke your Lambda function when an instance is put into a wait
// state due to a lifecycle hook.
//
// - (Optional) Create a notification target and an IAM role. The target can be
// either an Amazon SQS queue or an Amazon SNS topic. The role allows Amazon EC2
// Auto Scaling to publish lifecycle notifications to the target.
//
// - Create the lifecycle hook. Specify whether the hook is used when the
// instances launch or terminate.
//
// - If you need more time, record the lifecycle action heartbeat to keep the
// instance in a wait state.
//
// - If you finish before the timeout period ends, send a callback by using the [CompleteLifecycleAction]
// API call.
//
// For more information, see [Complete a lifecycle action] in the Amazon EC2 Auto Scaling User Guide.
//
// [CompleteLifecycleAction]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_CompleteLifecycleAction.html
// [Complete a lifecycle action]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/completing-lifecycle-hooks.html
func autoscaling_CompleteLifecycleAction(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.CompleteLifecycleActionInput{
		// AutoScalingGroupName: *string, // Required
		// LifecycleActionResult: *string, // Required
		// LifecycleHookName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingLifecycleActionResult) > 0 {
		input.LifecycleActionResult = aws.String(_autoscalingLifecycleActionResult)
	}
	if len(_autoscalingLifecycleHookName) > 0 {
		input.LifecycleHookName = aws.String(_autoscalingLifecycleHookName)
	}
	if len(_autoscalingInstanceId) > 0 {
		input.InstanceId = aws.String(_autoscalingInstanceId)
	}
	if len(_autoscalingLifecycleActionToken) > 0 {
		input.LifecycleActionToken = aws.String(_autoscalingLifecycleActionToken)
	}

	if resp, err := client.CompleteLifecycleAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// We strongly recommend using a launch template when calling this operation to
// ensure full functionality for Amazon EC2 Auto Scaling and Amazon EC2.
//
// Creates an Auto Scaling group with the specified name and attributes.
//
// If you exceed your maximum limit of Auto Scaling groups, the call fails. To
// query this limit, call the [DescribeAccountLimits]API. For information about updating this limit, see [Quotas for Amazon EC2 Auto Scaling]
// in the Amazon EC2 Auto Scaling User Guide.
//
// If you're new to Amazon EC2 Auto Scaling, see the introductory tutorials in [Get started with Amazon EC2 Auto Scaling] in
// the Amazon EC2 Auto Scaling User Guide.
//
// Every Auto Scaling group has three size properties ( DesiredCapacity , MaxSize ,
// and MinSize ). Usually, you set these sizes based on a specific number of
// instances. However, if you configure a mixed instances policy that defines
// weights for the instance types, you must specify these sizes with the same units
// that you use for weighting instances.
//
// [DescribeAccountLimits]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeAccountLimits.html
// [Get started with Amazon EC2 Auto Scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/get-started-with-ec2-auto-scaling.html
// [Quotas for Amazon EC2 Auto Scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-quotas.html
func autoscaling_CreateAutoScalingGroup(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.CreateAutoScalingGroupInput{
		// AutoScalingGroupName: *string, // Required
		// MaxSize: *int32, // Required
		// MinSize: *int32, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingMaxSize) > 0 {
		if err := assignInputField(input, "MaxSize", _autoscalingMaxSize); err != nil {
			log.Errorf("invalid --max-size: %s", err.Error())
			return
		}
	}
	if len(_autoscalingMinSize) > 0 {
		if err := assignInputField(input, "MinSize", _autoscalingMinSize); err != nil {
			log.Errorf("invalid --min-size: %s", err.Error())
			return
		}
	}
	if len(_autoscalingAvailabilityZoneDistribution) > 0 {
		if err := assignInputField(input, "AvailabilityZoneDistribution", _autoscalingAvailabilityZoneDistribution); err != nil {
			log.Errorf("invalid --availability-zone-distribution: %s", err.Error())
			return
		}
	}
	if len(_autoscalingAvailabilityZoneImpairmentPolicy) > 0 {
		if err := assignInputField(input, "AvailabilityZoneImpairmentPolicy", _autoscalingAvailabilityZoneImpairmentPolicy); err != nil {
			log.Errorf("invalid --availability-zone-impairment-policy: %s", err.Error())
			return
		}
	}
	if len(_autoscalingAvailabilityZones) > 0 {
		input.AvailabilityZones = append([]string(nil), _autoscalingAvailabilityZones...)
	}
	if len(_autoscalingCapacityRebalance) > 0 {
		if err := assignInputField(input, "CapacityRebalance", _autoscalingCapacityRebalance); err != nil {
			log.Errorf("invalid --capacity-rebalance: %s", err.Error())
			return
		}
	}
	if len(_autoscalingCapacityReservationSpecification) > 0 {
		if err := assignInputField(input, "CapacityReservationSpecification", _autoscalingCapacityReservationSpecification); err != nil {
			log.Errorf("invalid --capacity-reservation-specification: %s", err.Error())
			return
		}
	}
	if len(_autoscalingContext) > 0 {
		input.Context = aws.String(_autoscalingContext)
	}
	if len(_autoscalingDefaultCooldown) > 0 {
		if err := assignInputField(input, "DefaultCooldown", _autoscalingDefaultCooldown); err != nil {
			log.Errorf("invalid --default-cooldown: %s", err.Error())
			return
		}
	}
	if len(_autoscalingDefaultInstanceWarmup) > 0 {
		if err := assignInputField(input, "DefaultInstanceWarmup", _autoscalingDefaultInstanceWarmup); err != nil {
			log.Errorf("invalid --default-instance-warmup: %s", err.Error())
			return
		}
	}
	if len(_autoscalingDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _autoscalingDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_autoscalingDesiredCapacity) > 0 {
		if err := assignInputField(input, "DesiredCapacity", _autoscalingDesiredCapacity); err != nil {
			log.Errorf("invalid --desired-capacity: %s", err.Error())
			return
		}
	}
	if len(_autoscalingDesiredCapacityType) > 0 {
		input.DesiredCapacityType = aws.String(_autoscalingDesiredCapacityType)
	}
	if len(_autoscalingHealthCheckGracePeriod) > 0 {
		if err := assignInputField(input, "HealthCheckGracePeriod", _autoscalingHealthCheckGracePeriod); err != nil {
			log.Errorf("invalid --health-check-grace-period: %s", err.Error())
			return
		}
	}
	if len(_autoscalingHealthCheckType) > 0 {
		input.HealthCheckType = aws.String(_autoscalingHealthCheckType)
	}
	if len(_autoscalingInstanceId) > 0 {
		input.InstanceId = aws.String(_autoscalingInstanceId)
	}
	if len(_autoscalingInstanceLifecyclePolicy) > 0 {
		if err := assignInputField(input, "InstanceLifecyclePolicy", _autoscalingInstanceLifecyclePolicy); err != nil {
			log.Errorf("invalid --instance-lifecycle-policy: %s", err.Error())
			return
		}
	}
	if len(_autoscalingInstanceMaintenancePolicy) > 0 {
		if err := assignInputField(input, "InstanceMaintenancePolicy", _autoscalingInstanceMaintenancePolicy); err != nil {
			log.Errorf("invalid --instance-maintenance-policy: %s", err.Error())
			return
		}
	}
	if len(_autoscalingLaunchConfigurationName) > 0 {
		input.LaunchConfigurationName = aws.String(_autoscalingLaunchConfigurationName)
	}
	if len(_autoscalingLaunchTemplate) > 0 {
		if err := assignInputField(input, "LaunchTemplate", _autoscalingLaunchTemplate); err != nil {
			log.Errorf("invalid --launch-template: %s", err.Error())
			return
		}
	}
	if len(_autoscalingLifecycleHookSpecificationList) > 0 {
		if err := assignInputField(input, "LifecycleHookSpecificationList", _autoscalingLifecycleHookSpecificationList); err != nil {
			log.Errorf("invalid --lifecycle-hook-specification-list: %s", err.Error())
			return
		}
	}
	if len(_autoscalingLoadBalancerNames) > 0 {
		input.LoadBalancerNames = append([]string(nil), _autoscalingLoadBalancerNames...)
	}
	if len(_autoscalingMaxInstanceLifetime) > 0 {
		if err := assignInputField(input, "MaxInstanceLifetime", _autoscalingMaxInstanceLifetime); err != nil {
			log.Errorf("invalid --max-instance-lifetime: %s", err.Error())
			return
		}
	}
	if len(_autoscalingMixedInstancesPolicy) > 0 {
		if err := assignInputField(input, "MixedInstancesPolicy", _autoscalingMixedInstancesPolicy); err != nil {
			log.Errorf("invalid --mixed-instances-policy: %s", err.Error())
			return
		}
	}
	if len(_autoscalingNewInstancesProtectedFromScaleIn) > 0 {
		if err := assignInputField(input, "NewInstancesProtectedFromScaleIn", _autoscalingNewInstancesProtectedFromScaleIn); err != nil {
			log.Errorf("invalid --new-instances-protected-from-scale-in: %s", err.Error())
			return
		}
	}
	if len(_autoscalingPlacementGroup) > 0 {
		input.PlacementGroup = aws.String(_autoscalingPlacementGroup)
	}
	if len(_autoscalingServiceLinkedRoleARN) > 0 {
		input.ServiceLinkedRoleARN = aws.String(_autoscalingServiceLinkedRoleARN)
	}
	if len(_autoscalingSkipZonalShiftValidation) > 0 {
		if err := assignInputField(input, "SkipZonalShiftValidation", _autoscalingSkipZonalShiftValidation); err != nil {
			log.Errorf("invalid --skip-zonal-shift-validation: %s", err.Error())
			return
		}
	}
	if len(_autoscalingTags) > 0 {
		if err := assignInputField(input, "Tags", _autoscalingTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_autoscalingTargetGroupARNs) > 0 {
		input.TargetGroupARNs = append([]string(nil), _autoscalingTargetGroupARNs...)
	}
	if len(_autoscalingTerminationPolicies) > 0 {
		input.TerminationPolicies = append([]string(nil), _autoscalingTerminationPolicies...)
	}
	if len(_autoscalingTrafficSources) > 0 {
		if err := assignInputField(input, "TrafficSources", _autoscalingTrafficSources); err != nil {
			log.Errorf("invalid --traffic-sources: %s", err.Error())
			return
		}
	}
	if len(_autoscalingVPCZoneIdentifier) > 0 {
		input.VPCZoneIdentifier = aws.String(_autoscalingVPCZoneIdentifier)
	}

	if resp, err := client.CreateAutoScalingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a launch configuration.
// If you exceed your maximum limit of launch configurations, the call fails. To
// query this limit, call the [DescribeAccountLimits]API. For information about updating this limit, see [Quotas for Amazon EC2 Auto Scaling]
// in the Amazon EC2 Auto Scaling User Guide.
//
// For more information, see [Launch configurations] in the Amazon EC2 Auto Scaling User Guide.
//
// Amazon EC2 Auto Scaling configures instances launched as part of an Auto
// Scaling group using either a launch template or a launch configuration. We
// strongly recommend that you do not use launch configurations. They do not
// provide full functionality for Amazon EC2 Auto Scaling or Amazon EC2. For
// information about using launch templates, see [Launch templates]in the Amazon EC2 Auto Scaling
// User Guide.
//
// [DescribeAccountLimits]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeAccountLimits.html
// [Quotas for Amazon EC2 Auto Scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-quotas.html
// [Launch configurations]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/launch-configurations.html
// [Launch templates]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/launch-templates.html
func autoscaling_CreateLaunchConfiguration(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.CreateLaunchConfigurationInput{
		// LaunchConfigurationName: *string, // Required
	}

	if len(_autoscalingLaunchConfigurationName) > 0 {
		input.LaunchConfigurationName = aws.String(_autoscalingLaunchConfigurationName)
	}
	if len(_autoscalingAssociatePublicIpAddress) > 0 {
		if err := assignInputField(input, "AssociatePublicIpAddress", _autoscalingAssociatePublicIpAddress); err != nil {
			log.Errorf("invalid --associate-public-ip-address: %s", err.Error())
			return
		}
	}
	if len(_autoscalingBlockDeviceMappings) > 0 {
		if err := assignInputField(input, "BlockDeviceMappings", _autoscalingBlockDeviceMappings); err != nil {
			log.Errorf("invalid --block-device-mappings: %s", err.Error())
			return
		}
	}
	if len(_autoscalingClassicLinkVPCId) > 0 {
		input.ClassicLinkVPCId = aws.String(_autoscalingClassicLinkVPCId)
	}
	if len(_autoscalingClassicLinkVPCSecurityGroups) > 0 {
		input.ClassicLinkVPCSecurityGroups = append([]string(nil), _autoscalingClassicLinkVPCSecurityGroups...)
	}
	if len(_autoscalingEbsOptimized) > 0 {
		if err := assignInputField(input, "EbsOptimized", _autoscalingEbsOptimized); err != nil {
			log.Errorf("invalid --ebs-optimized: %s", err.Error())
			return
		}
	}
	if len(_autoscalingIamInstanceProfile) > 0 {
		input.IamInstanceProfile = aws.String(_autoscalingIamInstanceProfile)
	}
	if len(_autoscalingImageId) > 0 {
		input.ImageId = aws.String(_autoscalingImageId)
	}
	if len(_autoscalingInstanceId) > 0 {
		input.InstanceId = aws.String(_autoscalingInstanceId)
	}
	if len(_autoscalingInstanceMonitoring) > 0 {
		if err := assignInputField(input, "InstanceMonitoring", _autoscalingInstanceMonitoring); err != nil {
			log.Errorf("invalid --instance-monitoring: %s", err.Error())
			return
		}
	}
	if len(_autoscalingInstanceType) > 0 {
		input.InstanceType = aws.String(_autoscalingInstanceType)
	}
	if len(_autoscalingKernelId) > 0 {
		input.KernelId = aws.String(_autoscalingKernelId)
	}
	if len(_autoscalingKeyName) > 0 {
		input.KeyName = aws.String(_autoscalingKeyName)
	}
	if len(_autoscalingMetadataOptions) > 0 {
		if err := assignInputField(input, "MetadataOptions", _autoscalingMetadataOptions); err != nil {
			log.Errorf("invalid --metadata-options: %s", err.Error())
			return
		}
	}
	if len(_autoscalingPlacementTenancy) > 0 {
		input.PlacementTenancy = aws.String(_autoscalingPlacementTenancy)
	}
	if len(_autoscalingRamdiskId) > 0 {
		input.RamdiskId = aws.String(_autoscalingRamdiskId)
	}
	if len(_autoscalingSecurityGroups) > 0 {
		input.SecurityGroups = append([]string(nil), _autoscalingSecurityGroups...)
	}
	if len(_autoscalingSpotPrice) > 0 {
		input.SpotPrice = aws.String(_autoscalingSpotPrice)
	}
	if len(_autoscalingUserData) > 0 {
		input.UserData = aws.String(_autoscalingUserData)
	}

	if resp, err := client.CreateLaunchConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates tags for the specified Auto Scaling group.
// When you specify a tag with a key that already exists, the operation overwrites
// the previous tag definition, and you do not get an error message.
//
// For more information, see [Tag Auto Scaling groups and instances] in the Amazon EC2 Auto Scaling User Guide.
//
// [Tag Auto Scaling groups and instances]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-tagging.html
func autoscaling_CreateOrUpdateTags(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.CreateOrUpdateTagsInput{
		// Tags: []types.Tag, // Required
	}

	if len(_autoscalingTags) > 0 {
		if err := assignInputField(input, "Tags", _autoscalingTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOrUpdateTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified Auto Scaling group.
// If the group has instances or scaling activities in progress, you must specify
// the option to force the deletion in order for it to succeed. The force delete
// operation will also terminate the EC2 instances. If the group has a warm pool,
// the force delete option also deletes the warm pool.
//
// To remove instances from the Auto Scaling group before deleting it, call the [DetachInstances]
// API with the list of instances and the option to decrement the desired capacity.
// This ensures that Amazon EC2 Auto Scaling does not launch replacement instances.
//
// To terminate all instances before deleting the Auto Scaling group, call the [UpdateAutoScalingGroup]
// API and set the minimum size and desired capacity of the Auto Scaling group to
// zero.
//
// If the group has scaling policies, deleting the group deletes the policies, the
// underlying alarm actions, and any alarm that no longer has an associated action.
//
// For more information, see [Delete your Auto Scaling infrastructure] in the Amazon EC2 Auto Scaling User Guide.
//
// [Delete your Auto Scaling infrastructure]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-process-shutdown.html
// [DetachInstances]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DetachInstances.html
// [UpdateAutoScalingGroup]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_UpdateAutoScalingGroup.html
func autoscaling_DeleteAutoScalingGroup(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DeleteAutoScalingGroupInput{
		// AutoScalingGroupName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingForceDelete) > 0 {
		if err := assignInputField(input, "ForceDelete", _autoscalingForceDelete); err != nil {
			log.Errorf("invalid --force-delete: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteAutoScalingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified launch configuration.
// The launch configuration must not be attached to an Auto Scaling group. When
// this call completes, the launch configuration is no longer available for use.
func autoscaling_DeleteLaunchConfiguration(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DeleteLaunchConfigurationInput{
		// LaunchConfigurationName: *string, // Required
	}

	if len(_autoscalingLaunchConfigurationName) > 0 {
		input.LaunchConfigurationName = aws.String(_autoscalingLaunchConfigurationName)
	}

	if resp, err := client.DeleteLaunchConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified lifecycle hook.
// If there are any outstanding lifecycle actions, they are completed first (
// ABANDON for launching instances, CONTINUE for terminating instances).
func autoscaling_DeleteLifecycleHook(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DeleteLifecycleHookInput{
		// AutoScalingGroupName: *string, // Required
		// LifecycleHookName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingLifecycleHookName) > 0 {
		input.LifecycleHookName = aws.String(_autoscalingLifecycleHookName)
	}

	if resp, err := client.DeleteLifecycleHook(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified notification.
func autoscaling_DeleteNotificationConfiguration(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DeleteNotificationConfigurationInput{
		// AutoScalingGroupName: *string, // Required
		// TopicARN: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingTopicARN) > 0 {
		input.TopicARN = aws.String(_autoscalingTopicARN)
	}

	if resp, err := client.DeleteNotificationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified scaling policy.
// Deleting either a step scaling policy or a simple scaling policy deletes the
// underlying alarm action, but does not delete the alarm, even if it no longer has
// an associated action.
//
// For more information, see [Delete a scaling policy] in the Amazon EC2 Auto Scaling User Guide.
//
// [Delete a scaling policy]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/deleting-scaling-policy.html
func autoscaling_DeletePolicy(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DeletePolicyInput{
		// PolicyName: *string, // Required
	}

	if len(_autoscalingPolicyName) > 0 {
		input.PolicyName = aws.String(_autoscalingPolicyName)
	}
	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}

	if resp, err := client.DeletePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified scheduled action.
func autoscaling_DeleteScheduledAction(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DeleteScheduledActionInput{
		// AutoScalingGroupName: *string, // Required
		// ScheduledActionName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingScheduledActionName) > 0 {
		input.ScheduledActionName = aws.String(_autoscalingScheduledActionName)
	}

	if resp, err := client.DeleteScheduledAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified tags.
func autoscaling_DeleteTags(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DeleteTagsInput{
		// Tags: []types.Tag, // Required
	}

	if len(_autoscalingTags) > 0 {
		if err := assignInputField(input, "Tags", _autoscalingTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the warm pool for the specified Auto Scaling group.
// For more information, see [Warm pools for Amazon EC2 Auto Scaling] in the Amazon EC2 Auto Scaling User Guide.
//
// [Warm pools for Amazon EC2 Auto Scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-warm-pools.html
func autoscaling_DeleteWarmPool(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DeleteWarmPoolInput{
		// AutoScalingGroupName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingForceDelete) > 0 {
		if err := assignInputField(input, "ForceDelete", _autoscalingForceDelete); err != nil {
			log.Errorf("invalid --force-delete: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteWarmPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the current Amazon EC2 Auto Scaling resource quotas for your account.
// When you establish an Amazon Web Services account, the account has initial
// quotas on the maximum number of Auto Scaling groups and launch configurations
// that you can create in a given Region. For more information, see [Quotas for Amazon EC2 Auto Scaling]in the Amazon
// EC2 Auto Scaling User Guide.
//
// [Quotas for Amazon EC2 Auto Scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-quotas.html
func autoscaling_DescribeAccountLimits(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DescribeAccountLimitsInput{}

	if resp, err := client.DescribeAccountLimits(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the available adjustment types for step scaling and simple scaling
// policies.
//
// The following adjustment types are supported:
//
// - ChangeInCapacity
//
// - ExactCapacity
//
// - PercentChangeInCapacity
func autoscaling_DescribeAdjustmentTypes(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DescribeAdjustmentTypesInput{}

	if resp, err := client.DescribeAdjustmentTypes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the Auto Scaling groups in the account and Region.
// If you specify Auto Scaling group names, the output includes information for
// only the specified Auto Scaling groups. If you specify filters, the output
// includes information for only those Auto Scaling groups that meet the filter
// criteria. If you do not specify group names or filters, the output includes
// information for all Auto Scaling groups.
//
// This operation also returns information about instances in Auto Scaling groups.
// To retrieve information about the instances in a warm pool, you must call the [DescribeWarmPool]
// API.
//
// [DescribeWarmPool]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeWarmPool.html
func autoscaling_DescribeAutoScalingGroups(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DescribeAutoScalingGroupsInput{}

	if len(_autoscalingAutoScalingGroupNames) > 0 {
		input.AutoScalingGroupNames = append([]string(nil), _autoscalingAutoScalingGroupNames...)
	}
	if len(_autoscalingFilters) > 0 {
		if err := assignInputField(input, "Filters", _autoscalingFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_autoscalingIncludeInstances) > 0 {
		if err := assignInputField(input, "IncludeInstances", _autoscalingIncludeInstances); err != nil {
			log.Errorf("invalid --include-instances: %s", err.Error())
			return
		}
	}
	if len(_autoscalingMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _autoscalingMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_autoscalingNextToken) > 0 {
		input.NextToken = aws.String(_autoscalingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeAutoScalingGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*autoscaling.DescribeAutoScalingGroupsOutput
	p := autoscaling.NewDescribeAutoScalingGroupsPaginator(client, input)
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

// Gets information about the Auto Scaling instances in the account and Region.
func autoscaling_DescribeAutoScalingInstances(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DescribeAutoScalingInstancesInput{}

	if len(_autoscalingInstanceIds) > 0 {
		input.InstanceIds = append([]string(nil), _autoscalingInstanceIds...)
	}
	if len(_autoscalingMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _autoscalingMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_autoscalingNextToken) > 0 {
		input.NextToken = aws.String(_autoscalingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeAutoScalingInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*autoscaling.DescribeAutoScalingInstancesOutput
	p := autoscaling.NewDescribeAutoScalingInstancesPaginator(client, input)
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

// Describes the notification types that are supported by Amazon EC2 Auto Scaling.
func autoscaling_DescribeAutoScalingNotificationTypes(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DescribeAutoScalingNotificationTypesInput{}

	if resp, err := client.DescribeAutoScalingNotificationTypes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the instance refreshes for the specified Auto Scaling
// group from the previous six weeks.
//
// This operation is part of the [instance refresh feature] in Amazon EC2 Auto Scaling, which helps you
// update instances in your Auto Scaling group after you make configuration
// changes.
//
// To help you determine the status of an instance refresh, Amazon EC2 Auto
// Scaling returns information about the instance refreshes you previously
// initiated, including their status, start time, end time, the percentage of the
// instance refresh that is complete, and the number of instances remaining to
// update before the instance refresh is complete. If a rollback is initiated while
// an instance refresh is in progress, Amazon EC2 Auto Scaling also returns
// information about the rollback of the instance refresh.
//
// [instance refresh feature]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-refresh.html
func autoscaling_DescribeInstanceRefreshes(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DescribeInstanceRefreshesInput{
		// AutoScalingGroupName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingInstanceRefreshIds) > 0 {
		input.InstanceRefreshIds = append([]string(nil), _autoscalingInstanceRefreshIds...)
	}
	if len(_autoscalingMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _autoscalingMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_autoscalingNextToken) > 0 {
		input.NextToken = aws.String(_autoscalingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeInstanceRefreshes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*autoscaling.DescribeInstanceRefreshesOutput
	p := autoscaling.NewDescribeInstanceRefreshesPaginator(client, input)
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

// Gets information about the launch configurations in the account and Region.
func autoscaling_DescribeLaunchConfigurations(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DescribeLaunchConfigurationsInput{}

	if len(_autoscalingLaunchConfigurationNames) > 0 {
		input.LaunchConfigurationNames = append([]string(nil), _autoscalingLaunchConfigurationNames...)
	}
	if len(_autoscalingMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _autoscalingMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_autoscalingNextToken) > 0 {
		input.NextToken = aws.String(_autoscalingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeLaunchConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*autoscaling.DescribeLaunchConfigurationsOutput
	p := autoscaling.NewDescribeLaunchConfigurationsPaginator(client, input)
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

// Describes the available types of lifecycle hooks.
// The following hook types are supported:
//
// - autoscaling:EC2_INSTANCE_LAUNCHING
//
// - autoscaling:EC2_INSTANCE_TERMINATING
func autoscaling_DescribeLifecycleHookTypes(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DescribeLifecycleHookTypesInput{}

	if resp, err := client.DescribeLifecycleHookTypes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the lifecycle hooks for the specified Auto Scaling group.
func autoscaling_DescribeLifecycleHooks(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DescribeLifecycleHooksInput{
		// AutoScalingGroupName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingLifecycleHookNames) > 0 {
		input.LifecycleHookNames = append([]string(nil), _autoscalingLifecycleHookNames...)
	}

	if resp, err := client.DescribeLifecycleHooks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API operation is superseded by [DescribeTrafficSources], which can describe multiple traffic
// sources types. We recommend using DetachTrafficSources to simplify how you
// manage traffic sources. However, we continue to support
// DescribeLoadBalancerTargetGroups . You can use both the original
// DescribeLoadBalancerTargetGroups API operation and DescribeTrafficSources on
// the same Auto Scaling group.
//
// Gets information about the Elastic Load Balancing target groups for the
// specified Auto Scaling group.
//
// To determine the attachment status of the target group, use the State element
// in the response. When you attach a target group to an Auto Scaling group, the
// initial State value is Adding . The state transitions to Added after all Auto
// Scaling instances are registered with the target group. If Elastic Load
// Balancing health checks are enabled for the Auto Scaling group, the state
// transitions to InService after at least one Auto Scaling instance passes the
// health check. When the target group is in the InService state, Amazon EC2 Auto
// Scaling can terminate and replace any instances that are reported as unhealthy.
// If no registered instances pass the health checks, the target group doesn't
// enter the InService state.
//
// Target groups also have an InService state if you attach them in the [CreateAutoScalingGroup] API call.
// If your target group state is InService , but it is not working properly, check
// the scaling activities by calling [DescribeScalingActivities]and take any corrective actions necessary.
//
// For help with failed health checks, see [Troubleshooting Amazon EC2 Auto Scaling: Health checks] in the Amazon EC2 Auto Scaling User
// Guide. For more information, see [Use Elastic Load Balancing to distribute traffic across the instances in your Auto Scaling group]in the Amazon EC2 Auto Scaling User Guide.
//
// You can use this operation to describe target groups that were attached by
// using [AttachLoadBalancerTargetGroups], but not for target groups that were attached by using [AttachTrafficSources].
//
// [Troubleshooting Amazon EC2 Auto Scaling: Health checks]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ts-as-healthchecks.html
// [AttachLoadBalancerTargetGroups]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_AttachLoadBalancerTargetGroups.html
// [DescribeScalingActivities]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeScalingActivities.html
// [CreateAutoScalingGroup]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_CreateAutoScalingGroup.html
// [AttachTrafficSources]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_AttachTrafficSources.html
// [Use Elastic Load Balancing to distribute traffic across the instances in your Auto Scaling group]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-load-balancer.html
// [DescribeTrafficSources]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeTrafficSources.html
func autoscaling_DescribeLoadBalancerTargetGroups(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DescribeLoadBalancerTargetGroupsInput{
		// AutoScalingGroupName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _autoscalingMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_autoscalingNextToken) > 0 {
		input.NextToken = aws.String(_autoscalingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeLoadBalancerTargetGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*autoscaling.DescribeLoadBalancerTargetGroupsOutput
	p := autoscaling.NewDescribeLoadBalancerTargetGroupsPaginator(client, input)
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

// This API operation is superseded by [DescribeTrafficSources], which can describe multiple traffic
// sources types. We recommend using DescribeTrafficSources to simplify how you
// manage traffic sources. However, we continue to support DescribeLoadBalancers .
// You can use both the original DescribeLoadBalancers API operation and
// DescribeTrafficSources on the same Auto Scaling group.
//
// Gets information about the load balancers for the specified Auto Scaling group.
//
// This operation describes only Classic Load Balancers. If you have Application
// Load Balancers, Network Load Balancers, or Gateway Load Balancers, use the [DescribeLoadBalancerTargetGroups]API
// instead.
//
// To determine the attachment status of the load balancer, use the State element
// in the response. When you attach a load balancer to an Auto Scaling group, the
// initial State value is Adding . The state transitions to Added after all Auto
// Scaling instances are registered with the load balancer. If Elastic Load
// Balancing health checks are enabled for the Auto Scaling group, the state
// transitions to InService after at least one Auto Scaling instance passes the
// health check. When the load balancer is in the InService state, Amazon EC2 Auto
// Scaling can terminate and replace any instances that are reported as unhealthy.
// If no registered instances pass the health checks, the load balancer doesn't
// enter the InService state.
//
// Load balancers also have an InService state if you attach them in the [CreateAutoScalingGroup] API
// call. If your load balancer state is InService , but it is not working properly,
// check the scaling activities by calling [DescribeScalingActivities]and take any corrective actions
// necessary.
//
// For help with failed health checks, see [Troubleshooting Amazon EC2 Auto Scaling: Health checks] in the Amazon EC2 Auto Scaling User
// Guide. For more information, see [Use Elastic Load Balancing to distribute traffic across the instances in your Auto Scaling group]in the Amazon EC2 Auto Scaling User Guide.
//
// [Troubleshooting Amazon EC2 Auto Scaling: Health checks]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ts-as-healthchecks.html
// [DescribeScalingActivities]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeScalingActivities.html
// [DescribeLoadBalancerTargetGroups]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeLoadBalancerTargetGroups.html
// [CreateAutoScalingGroup]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_CreateAutoScalingGroup.html
// [Use Elastic Load Balancing to distribute traffic across the instances in your Auto Scaling group]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-load-balancer.html
// [DescribeTrafficSources]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeTrafficSources.html
func autoscaling_DescribeLoadBalancers(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DescribeLoadBalancersInput{
		// AutoScalingGroupName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _autoscalingMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_autoscalingNextToken) > 0 {
		input.NextToken = aws.String(_autoscalingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeLoadBalancers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*autoscaling.DescribeLoadBalancersOutput
	p := autoscaling.NewDescribeLoadBalancersPaginator(client, input)
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

// Describes the available CloudWatch metrics for Amazon EC2 Auto Scaling.
func autoscaling_DescribeMetricCollectionTypes(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DescribeMetricCollectionTypesInput{}

	if resp, err := client.DescribeMetricCollectionTypes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the Amazon SNS notifications that are configured for one
// or more Auto Scaling groups.
func autoscaling_DescribeNotificationConfigurations(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DescribeNotificationConfigurationsInput{}

	if len(_autoscalingAutoScalingGroupNames) > 0 {
		input.AutoScalingGroupNames = append([]string(nil), _autoscalingAutoScalingGroupNames...)
	}
	if len(_autoscalingMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _autoscalingMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_autoscalingNextToken) > 0 {
		input.NextToken = aws.String(_autoscalingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeNotificationConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*autoscaling.DescribeNotificationConfigurationsOutput
	p := autoscaling.NewDescribeNotificationConfigurationsPaginator(client, input)
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

// Gets information about the scaling policies in the account and Region.
func autoscaling_DescribePolicies(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DescribePoliciesInput{}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _autoscalingMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_autoscalingNextToken) > 0 {
		input.NextToken = aws.String(_autoscalingNextToken)
	}
	if len(_autoscalingPolicyNames) > 0 {
		input.PolicyNames = append([]string(nil), _autoscalingPolicyNames...)
	}
	if len(_autoscalingPolicyTypes) > 0 {
		input.PolicyTypes = append([]string(nil), _autoscalingPolicyTypes...)
	}

	if disablePaginator() {
		if resp, err := client.DescribePolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*autoscaling.DescribePoliciesOutput
	p := autoscaling.NewDescribePoliciesPaginator(client, input)
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

// Gets information about the scaling activities in the account and Region.
// When scaling events occur, you see a record of the scaling activity in the
// scaling activities. For more information, see [Verify a scaling activity for an Auto Scaling group]in the Amazon EC2 Auto Scaling
// User Guide.
//
// If the scaling event succeeds, the value of the StatusCode element in the
// response is Successful . If an attempt to launch instances failed, the
// StatusCode value is Failed or Cancelled and the StatusMessage element in the
// response indicates the cause of the failure. For help interpreting the
// StatusMessage , see [Troubleshooting Amazon EC2 Auto Scaling] in the Amazon EC2 Auto Scaling User Guide.
//
// [Troubleshooting Amazon EC2 Auto Scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/CHAP_Troubleshooting.html
// [Verify a scaling activity for an Auto Scaling group]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-verify-scaling-activity.html
func autoscaling_DescribeScalingActivities(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DescribeScalingActivitiesInput{}

	if len(_autoscalingActivityIds) > 0 {
		input.ActivityIds = append([]string(nil), _autoscalingActivityIds...)
	}
	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingFilters) > 0 {
		if err := assignInputField(input, "Filters", _autoscalingFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_autoscalingIncludeDeletedGroups) > 0 {
		if err := assignInputField(input, "IncludeDeletedGroups", _autoscalingIncludeDeletedGroups); err != nil {
			log.Errorf("invalid --include-deleted-groups: %s", err.Error())
			return
		}
	}
	if len(_autoscalingMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _autoscalingMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_autoscalingNextToken) > 0 {
		input.NextToken = aws.String(_autoscalingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeScalingActivities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*autoscaling.DescribeScalingActivitiesOutput
	p := autoscaling.NewDescribeScalingActivitiesPaginator(client, input)
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

// Describes the scaling process types for use with the [ResumeProcesses] and [SuspendProcesses] APIs.
//
// [ResumeProcesses]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_ResumeProcesses.html
// [SuspendProcesses]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_SuspendProcesses.html
func autoscaling_DescribeScalingProcessTypes(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DescribeScalingProcessTypesInput{}

	if resp, err := client.DescribeScalingProcessTypes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the scheduled actions that haven't run or that have not
// reached their end time.
//
// To describe the scaling activities for scheduled actions that have already run,
// call the [DescribeScalingActivities]API.
//
// [DescribeScalingActivities]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeScalingActivities.html
func autoscaling_DescribeScheduledActions(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DescribeScheduledActionsInput{}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _autoscalingEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_autoscalingMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _autoscalingMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_autoscalingNextToken) > 0 {
		input.NextToken = aws.String(_autoscalingNextToken)
	}
	if len(_autoscalingScheduledActionNames) > 0 {
		input.ScheduledActionNames = append([]string(nil), _autoscalingScheduledActionNames...)
	}
	if len(_autoscalingStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _autoscalingStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeScheduledActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*autoscaling.DescribeScheduledActionsOutput
	p := autoscaling.NewDescribeScheduledActionsPaginator(client, input)
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

// Describes the specified tags.
// You can use filters to limit the results. For example, you can query for the
// tags for a specific Auto Scaling group. You can specify multiple values for a
// filter. A tag must match at least one of the specified values for it to be
// included in the results.
//
// You can also specify multiple filters. The result includes information for a
// particular tag only if it matches all the filters. If there's no match, no
// special message is returned.
//
// For more information, see [Tag Auto Scaling groups and instances] in the Amazon EC2 Auto Scaling User Guide.
//
// [Tag Auto Scaling groups and instances]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-tagging.html
func autoscaling_DescribeTags(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DescribeTagsInput{}

	if len(_autoscalingFilters) > 0 {
		if err := assignInputField(input, "Filters", _autoscalingFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_autoscalingMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _autoscalingMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_autoscalingNextToken) > 0 {
		input.NextToken = aws.String(_autoscalingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*autoscaling.DescribeTagsOutput
	p := autoscaling.NewDescribeTagsPaginator(client, input)
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

// Describes the termination policies supported by Amazon EC2 Auto Scaling.
// For more information, see [Configure termination policies for Amazon EC2 Auto Scaling] in the Amazon EC2 Auto Scaling User Guide.
//
// [Configure termination policies for Amazon EC2 Auto Scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-termination-policies.html
func autoscaling_DescribeTerminationPolicyTypes(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DescribeTerminationPolicyTypesInput{}

	if resp, err := client.DescribeTerminationPolicyTypes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the traffic sources for the specified Auto Scaling group.
// You can optionally provide a traffic source type. If you provide a traffic
// source type, then the results only include that traffic source type.
//
// If you do not provide a traffic source type, then the results include all the
// traffic sources for the specified Auto Scaling group.
func autoscaling_DescribeTrafficSources(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DescribeTrafficSourcesInput{
		// AutoScalingGroupName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _autoscalingMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_autoscalingNextToken) > 0 {
		input.NextToken = aws.String(_autoscalingNextToken)
	}
	if len(_autoscalingTrafficSourceType) > 0 {
		input.TrafficSourceType = aws.String(_autoscalingTrafficSourceType)
	}

	if disablePaginator() {
		if resp, err := client.DescribeTrafficSources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*autoscaling.DescribeTrafficSourcesOutput
	p := autoscaling.NewDescribeTrafficSourcesPaginator(client, input)
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

// Gets information about a warm pool and its instances.
// For more information, see [Warm pools for Amazon EC2 Auto Scaling] in the Amazon EC2 Auto Scaling User Guide.
//
// [Warm pools for Amazon EC2 Auto Scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-warm-pools.html
func autoscaling_DescribeWarmPool(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DescribeWarmPoolInput{
		// AutoScalingGroupName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _autoscalingMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_autoscalingNextToken) > 0 {
		input.NextToken = aws.String(_autoscalingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeWarmPool(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*autoscaling.DescribeWarmPoolOutput
	p := autoscaling.NewDescribeWarmPoolPaginator(client, input)
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

// Removes one or more instances from the specified Auto Scaling group.
// After the instances are detached, you can manage them independent of the Auto
// Scaling group.
//
// If you do not specify the option to decrement the desired capacity, Amazon EC2
// Auto Scaling launches instances to replace the ones that are detached.
//
// If there is a Classic Load Balancer attached to the Auto Scaling group, the
// instances are deregistered from the load balancer. If there are target groups
// attached to the Auto Scaling group, the instances are deregistered from the
// target groups.
//
// For more information, see [Detach or attach instances] in the Amazon EC2 Auto Scaling User Guide.
//
// [Detach or attach instances]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-detach-attach-instances.html
func autoscaling_DetachInstances(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DetachInstancesInput{
		// AutoScalingGroupName: *string, // Required
		// ShouldDecrementDesiredCapacity: *bool, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingShouldDecrementDesiredCapacity) > 0 {
		if err := assignInputField(input, "ShouldDecrementDesiredCapacity", _autoscalingShouldDecrementDesiredCapacity); err != nil {
			log.Errorf("invalid --should-decrement-desired-capacity: %s", err.Error())
			return
		}
	}
	if len(_autoscalingInstanceIds) > 0 {
		input.InstanceIds = append([]string(nil), _autoscalingInstanceIds...)
	}

	if resp, err := client.DetachInstances(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API operation is superseded by [DetachTrafficSources], which can detach multiple traffic sources
// types. We recommend using DetachTrafficSources to simplify how you manage
// traffic sources. However, we continue to support DetachLoadBalancerTargetGroups
// . You can use both the original DetachLoadBalancerTargetGroups API operation
// and DetachTrafficSources on the same Auto Scaling group.
//
// Detaches one or more target groups from the specified Auto Scaling group.
//
// When you detach a target group, it enters the Removing state while
// deregistering the instances in the group. When all instances are deregistered,
// then you can no longer describe the target group using the [DescribeLoadBalancerTargetGroups]API call. The
// instances remain running.
//
// You can use this operation to detach target groups that were attached by using [AttachLoadBalancerTargetGroups]
// , but not for target groups that were attached by using [AttachTrafficSources].
//
// [AttachLoadBalancerTargetGroups]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_AttachLoadBalancerTargetGroups.html
// [DescribeLoadBalancerTargetGroups]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeLoadBalancerTargetGroups.html
// [AttachTrafficSources]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_AttachTrafficSources.html
// [DetachTrafficSources]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DetachTrafficSources.html
func autoscaling_DetachLoadBalancerTargetGroups(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DetachLoadBalancerTargetGroupsInput{
		// AutoScalingGroupName: *string, // Required
		// TargetGroupARNs: []string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingTargetGroupARNs) > 0 {
		input.TargetGroupARNs = append([]string(nil), _autoscalingTargetGroupARNs...)
	}

	if resp, err := client.DetachLoadBalancerTargetGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API operation is superseded by [DetachTrafficSources], which can detach multiple traffic sources
// types. We recommend using DetachTrafficSources to simplify how you manage
// traffic sources. However, we continue to support DetachLoadBalancers . You can
// use both the original DetachLoadBalancers API operation and DetachTrafficSources
// on the same Auto Scaling group.
//
// Detaches one or more Classic Load Balancers from the specified Auto Scaling
// group.
//
// This operation detaches only Classic Load Balancers. If you have Application
// Load Balancers, Network Load Balancers, or Gateway Load Balancers, use the [DetachLoadBalancerTargetGroups]API
// instead.
//
// When you detach a load balancer, it enters the Removing state while
// deregistering the instances in the group. When all instances are deregistered,
// then you can no longer describe the load balancer using the [DescribeLoadBalancers]API call. The
// instances remain running.
//
// [DetachLoadBalancerTargetGroups]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DetachLoadBalancerTargetGroups.html
// [DescribeLoadBalancers]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeLoadBalancers.html
// [DetachTrafficSources]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DetachTrafficSources.html
func autoscaling_DetachLoadBalancers(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DetachLoadBalancersInput{
		// AutoScalingGroupName: *string, // Required
		// LoadBalancerNames: []string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingLoadBalancerNames) > 0 {
		input.LoadBalancerNames = append([]string(nil), _autoscalingLoadBalancerNames...)
	}

	if resp, err := client.DetachLoadBalancers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detaches one or more traffic sources from the specified Auto Scaling group.
// When you detach a traffic source, it enters the Removing state while
// deregistering the instances in the group. When all instances are deregistered,
// then you can no longer describe the traffic source using the [DescribeTrafficSources]API call. The
// instances continue to run.
//
// [DescribeTrafficSources]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeTrafficSources.html
func autoscaling_DetachTrafficSources(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DetachTrafficSourcesInput{
		// AutoScalingGroupName: *string, // Required
		// TrafficSources: []types.TrafficSourceIdentifier, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingTrafficSources) > 0 {
		if err := assignInputField(input, "TrafficSources", _autoscalingTrafficSources); err != nil {
			log.Errorf("invalid --traffic-sources: %s", err.Error())
			return
		}
	}

	if resp, err := client.DetachTrafficSources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables group metrics collection for the specified Auto Scaling group.
func autoscaling_DisableMetricsCollection(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.DisableMetricsCollectionInput{
		// AutoScalingGroupName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingMetrics) > 0 {
		input.Metrics = append([]string(nil), _autoscalingMetrics...)
	}

	if resp, err := client.DisableMetricsCollection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables group metrics collection for the specified Auto Scaling group.
// You can use these metrics to track changes in an Auto Scaling group and to set
// alarms on threshold values. You can view group metrics using the Amazon EC2 Auto
// Scaling console or the CloudWatch console. For more information, see [Monitor CloudWatch metrics for your Auto Scaling groups and instances]in the
// Amazon EC2 Auto Scaling User Guide.
//
// [Monitor CloudWatch metrics for your Auto Scaling groups and instances]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-cloudwatch-monitoring.html
func autoscaling_EnableMetricsCollection(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.EnableMetricsCollectionInput{
		// AutoScalingGroupName: *string, // Required
		// Granularity: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingGranularity) > 0 {
		input.Granularity = aws.String(_autoscalingGranularity)
	}
	if len(_autoscalingMetrics) > 0 {
		input.Metrics = append([]string(nil), _autoscalingMetrics...)
	}

	if resp, err := client.EnableMetricsCollection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Moves the specified instances into the standby state.
// If you choose to decrement the desired capacity of the Auto Scaling group, the
// instances can enter standby as long as the desired capacity of the Auto Scaling
// group after the instances are placed into standby is equal to or greater than
// the minimum capacity of the group.
//
// If you choose not to decrement the desired capacity of the Auto Scaling group,
// the Auto Scaling group launches new instances to replace the instances on
// standby.
//
// For more information, see [Temporarily removing instances from your Auto Scaling group] in the Amazon EC2 Auto Scaling User Guide.
//
// [Temporarily removing instances from your Auto Scaling group]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-enter-exit-standby.html
func autoscaling_EnterStandby(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.EnterStandbyInput{
		// AutoScalingGroupName: *string, // Required
		// ShouldDecrementDesiredCapacity: *bool, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingShouldDecrementDesiredCapacity) > 0 {
		if err := assignInputField(input, "ShouldDecrementDesiredCapacity", _autoscalingShouldDecrementDesiredCapacity); err != nil {
			log.Errorf("invalid --should-decrement-desired-capacity: %s", err.Error())
			return
		}
	}
	if len(_autoscalingInstanceIds) > 0 {
		input.InstanceIds = append([]string(nil), _autoscalingInstanceIds...)
	}

	if resp, err := client.EnterStandby(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Executes the specified policy. This can be useful for testing the design of
// your scaling policy.
func autoscaling_ExecutePolicy(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.ExecutePolicyInput{
		// PolicyName: *string, // Required
	}

	if len(_autoscalingPolicyName) > 0 {
		input.PolicyName = aws.String(_autoscalingPolicyName)
	}
	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingBreachThreshold) > 0 {
		if err := assignInputField(input, "BreachThreshold", _autoscalingBreachThreshold); err != nil {
			log.Errorf("invalid --breach-threshold: %s", err.Error())
			return
		}
	}
	if len(_autoscalingHonorCooldown) > 0 {
		if err := assignInputField(input, "HonorCooldown", _autoscalingHonorCooldown); err != nil {
			log.Errorf("invalid --honor-cooldown: %s", err.Error())
			return
		}
	}
	if len(_autoscalingMetricValue) > 0 {
		if err := assignInputField(input, "MetricValue", _autoscalingMetricValue); err != nil {
			log.Errorf("invalid --metric-value: %s", err.Error())
			return
		}
	}

	if resp, err := client.ExecutePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Moves the specified instances out of the standby state.
// After you put the instances back in service, the desired capacity is
// incremented.
//
// For more information, see [Temporarily removing instances from your Auto Scaling group] in the Amazon EC2 Auto Scaling User Guide.
//
// [Temporarily removing instances from your Auto Scaling group]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-enter-exit-standby.html
func autoscaling_ExitStandby(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.ExitStandbyInput{
		// AutoScalingGroupName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingInstanceIds) > 0 {
		input.InstanceIds = append([]string(nil), _autoscalingInstanceIds...)
	}

	if resp, err := client.ExitStandby(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the forecast data for a predictive scaling policy.
// Load forecasts are predictions of the hourly load values using historical load
// data from CloudWatch and an analysis of historical trends. Capacity forecasts
// are represented as predicted values for the minimum capacity that is needed on
// an hourly basis, based on the hourly load forecast.
//
// A minimum of 24 hours of data is required to create the initial forecasts.
// However, having a full 14 days of historical data results in more accurate
// forecasts.
//
// For more information, see [Predictive scaling for Amazon EC2 Auto Scaling] in the Amazon EC2 Auto Scaling User Guide.
//
// [Predictive scaling for Amazon EC2 Auto Scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-predictive-scaling.html
func autoscaling_GetPredictiveScalingForecast(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.GetPredictiveScalingForecastInput{
		// AutoScalingGroupName: *string, // Required
		// EndTime: *time.Time, // Required
		// PolicyName: *string, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _autoscalingEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_autoscalingPolicyName) > 0 {
		input.PolicyName = aws.String(_autoscalingPolicyName)
	}
	if len(_autoscalingStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _autoscalingStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetPredictiveScalingForecast(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Launches a specified number of instances in an Auto Scaling group. Returns
// instance IDs and other details if launch is successful or error details if
// launch is unsuccessful.
func autoscaling_LaunchInstances(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.LaunchInstancesInput{
		// AutoScalingGroupName: *string, // Required
		// ClientToken: *string, // Required
		// RequestedCapacity: *int32, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingClientToken) > 0 {
		input.ClientToken = aws.String(_autoscalingClientToken)
	}
	if len(_autoscalingRequestedCapacity) > 0 {
		if err := assignInputField(input, "RequestedCapacity", _autoscalingRequestedCapacity); err != nil {
			log.Errorf("invalid --requested-capacity: %s", err.Error())
			return
		}
	}
	if len(_autoscalingAvailabilityZoneIds) > 0 {
		input.AvailabilityZoneIds = append([]string(nil), _autoscalingAvailabilityZoneIds...)
	}
	if len(_autoscalingAvailabilityZones) > 0 {
		input.AvailabilityZones = append([]string(nil), _autoscalingAvailabilityZones...)
	}
	if len(_autoscalingRetryStrategy) > 0 {
		if err := assignInputField(input, "RetryStrategy", _autoscalingRetryStrategy); err != nil {
			log.Errorf("invalid --retry-strategy: %s", err.Error())
			return
		}
	}
	if len(_autoscalingSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _autoscalingSubnetIds...)
	}

	if resp, err := client.LaunchInstances(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a lifecycle hook for the specified Auto Scaling group.
// Lifecycle hooks let you create solutions that are aware of events in the Auto
// Scaling instance lifecycle, and then perform a custom action on instances when
// the corresponding lifecycle event occurs.
//
// This step is a part of the procedure for adding a lifecycle hook to an Auto
// Scaling group:
//
// - (Optional) Create a launch template or launch configuration with a user
// data script that runs while an instance is in a wait state due to a lifecycle
// hook.
//
// - (Optional) Create a Lambda function and a rule that allows Amazon
// EventBridge to invoke your Lambda function when an instance is put into a wait
// state due to a lifecycle hook.
//
// - (Optional) Create a notification target and an IAM role. The target can be
// either an Amazon SQS queue or an Amazon SNS topic. The role allows Amazon EC2
// Auto Scaling to publish lifecycle notifications to the target.
//
// - Create the lifecycle hook. Specify whether the hook is used when the
// instances launch or terminate.
//
// - If you need more time, record the lifecycle action heartbeat to keep the
// instance in a wait state using the [RecordLifecycleActionHeartbeat]API call.
//
// - If you finish before the timeout period ends, send a callback by using the [CompleteLifecycleAction]
// API call.
//
// For more information, see [Amazon EC2 Auto Scaling lifecycle hooks] in the Amazon EC2 Auto Scaling User Guide.
//
// If you exceed your maximum limit of lifecycle hooks, which by default is 50 per
// Auto Scaling group, the call fails.
//
// You can view the lifecycle hooks for an Auto Scaling group using the [DescribeLifecycleHooks] API call.
// If you are no longer using a lifecycle hook, you can delete it by calling the [DeleteLifecycleHook]
// API.
//
// [RecordLifecycleActionHeartbeat]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_RecordLifecycleActionHeartbeat.html
// [CompleteLifecycleAction]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_CompleteLifecycleAction.html
// [Amazon EC2 Auto Scaling lifecycle hooks]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/lifecycle-hooks.html
// [DescribeLifecycleHooks]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeLifecycleHooks.html
// [DeleteLifecycleHook]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DeleteLifecycleHook.html
func autoscaling_PutLifecycleHook(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.PutLifecycleHookInput{
		// AutoScalingGroupName: *string, // Required
		// LifecycleHookName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingLifecycleHookName) > 0 {
		input.LifecycleHookName = aws.String(_autoscalingLifecycleHookName)
	}
	if len(_autoscalingDefaultResult) > 0 {
		input.DefaultResult = aws.String(_autoscalingDefaultResult)
	}
	if len(_autoscalingHeartbeatTimeout) > 0 {
		if err := assignInputField(input, "HeartbeatTimeout", _autoscalingHeartbeatTimeout); err != nil {
			log.Errorf("invalid --heartbeat-timeout: %s", err.Error())
			return
		}
	}
	if len(_autoscalingLifecycleTransition) > 0 {
		input.LifecycleTransition = aws.String(_autoscalingLifecycleTransition)
	}
	if len(_autoscalingNotificationMetadata) > 0 {
		input.NotificationMetadata = aws.String(_autoscalingNotificationMetadata)
	}
	if len(_autoscalingNotificationTargetARN) > 0 {
		input.NotificationTargetARN = aws.String(_autoscalingNotificationTargetARN)
	}
	if len(_autoscalingRoleARN) > 0 {
		input.RoleARN = aws.String(_autoscalingRoleARN)
	}

	if resp, err := client.PutLifecycleHook(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configures an Auto Scaling group to send notifications when specified events
// take place. Subscribers to the specified topic can have messages delivered to an
// endpoint such as a web server or an email address.
//
// This configuration overwrites any existing configuration.
//
// For more information, see [Amazon SNS notification options for Amazon EC2 Auto Scaling] in the Amazon EC2 Auto Scaling User Guide.
//
// If you exceed your maximum limit of SNS topics, which is 10 per Auto Scaling
// group, the call fails.
//
// [Amazon SNS notification options for Amazon EC2 Auto Scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-sns-notifications.html
func autoscaling_PutNotificationConfiguration(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.PutNotificationConfigurationInput{
		// AutoScalingGroupName: *string, // Required
		// NotificationTypes: []string, // Required
		// TopicARN: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingNotificationTypes) > 0 {
		input.NotificationTypes = append([]string(nil), _autoscalingNotificationTypes...)
	}
	if len(_autoscalingTopicARN) > 0 {
		input.TopicARN = aws.String(_autoscalingTopicARN)
	}

	if resp, err := client.PutNotificationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a scaling policy for an Auto Scaling group. Scaling policies
// are used to scale an Auto Scaling group based on configurable metrics. If no
// policies are defined, the dynamic scaling and predictive scaling features are
// not used.
//
// For more information about using dynamic scaling, see [Target tracking scaling policies] and [Step and simple scaling policies] in the Amazon EC2
// Auto Scaling User Guide.
//
// For more information about using predictive scaling, see [Predictive scaling for Amazon EC2 Auto Scaling] in the Amazon EC2
// Auto Scaling User Guide.
//
// You can view the scaling policies for an Auto Scaling group using the [DescribePolicies] API
// call. If you are no longer using a scaling policy, you can delete it by calling
// the [DeletePolicy]API.
//
// [Step and simple scaling policies]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html
// [DeletePolicy]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DeletePolicy.html
// [Target tracking scaling policies]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-target-tracking.html
// [DescribePolicies]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribePolicies.html
// [Predictive scaling for Amazon EC2 Auto Scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-predictive-scaling.html
func autoscaling_PutScalingPolicy(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.PutScalingPolicyInput{
		// AutoScalingGroupName: *string, // Required
		// PolicyName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingPolicyName) > 0 {
		input.PolicyName = aws.String(_autoscalingPolicyName)
	}
	if len(_autoscalingAdjustmentType) > 0 {
		input.AdjustmentType = aws.String(_autoscalingAdjustmentType)
	}
	if len(_autoscalingCooldown) > 0 {
		if err := assignInputField(input, "Cooldown", _autoscalingCooldown); err != nil {
			log.Errorf("invalid --cooldown: %s", err.Error())
			return
		}
	}
	if len(_autoscalingEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _autoscalingEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_autoscalingEstimatedInstanceWarmup) > 0 {
		if err := assignInputField(input, "EstimatedInstanceWarmup", _autoscalingEstimatedInstanceWarmup); err != nil {
			log.Errorf("invalid --estimated-instance-warmup: %s", err.Error())
			return
		}
	}
	if len(_autoscalingMetricAggregationType) > 0 {
		input.MetricAggregationType = aws.String(_autoscalingMetricAggregationType)
	}
	if len(_autoscalingMinAdjustmentMagnitude) > 0 {
		if err := assignInputField(input, "MinAdjustmentMagnitude", _autoscalingMinAdjustmentMagnitude); err != nil {
			log.Errorf("invalid --min-adjustment-magnitude: %s", err.Error())
			return
		}
	}
	if len(_autoscalingMinAdjustmentStep) > 0 {
		if err := assignInputField(input, "MinAdjustmentStep", _autoscalingMinAdjustmentStep); err != nil {
			log.Errorf("invalid --min-adjustment-step: %s", err.Error())
			return
		}
	}
	if len(_autoscalingPolicyType) > 0 {
		input.PolicyType = aws.String(_autoscalingPolicyType)
	}
	if len(_autoscalingPredictiveScalingConfiguration) > 0 {
		if err := assignInputField(input, "PredictiveScalingConfiguration", _autoscalingPredictiveScalingConfiguration); err != nil {
			log.Errorf("invalid --predictive-scaling-configuration: %s", err.Error())
			return
		}
	}
	if len(_autoscalingScalingAdjustment) > 0 {
		if err := assignInputField(input, "ScalingAdjustment", _autoscalingScalingAdjustment); err != nil {
			log.Errorf("invalid --scaling-adjustment: %s", err.Error())
			return
		}
	}
	if len(_autoscalingStepAdjustments) > 0 {
		if err := assignInputField(input, "StepAdjustments", _autoscalingStepAdjustments); err != nil {
			log.Errorf("invalid --step-adjustments: %s", err.Error())
			return
		}
	}
	if len(_autoscalingTargetTrackingConfiguration) > 0 {
		if err := assignInputField(input, "TargetTrackingConfiguration", _autoscalingTargetTrackingConfiguration); err != nil {
			log.Errorf("invalid --target-tracking-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutScalingPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a scheduled scaling action for an Auto Scaling group.
// For more information, see [Scheduled scaling] in the Amazon EC2 Auto Scaling User Guide.
//
// You can view the scheduled actions for an Auto Scaling group using the [DescribeScheduledActions] API
// call. If you are no longer using a scheduled action, you can delete it by
// calling the [DeleteScheduledAction]API.
//
// If you try to schedule your action in the past, Amazon EC2 Auto Scaling returns
// an error message.
//
// [DeleteScheduledAction]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DeleteScheduledAction.html
// [DescribeScheduledActions]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeScheduledActions.html
// [Scheduled scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-scheduled-scaling.html
func autoscaling_PutScheduledUpdateGroupAction(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.PutScheduledUpdateGroupActionInput{
		// AutoScalingGroupName: *string, // Required
		// ScheduledActionName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingScheduledActionName) > 0 {
		input.ScheduledActionName = aws.String(_autoscalingScheduledActionName)
	}
	if len(_autoscalingDesiredCapacity) > 0 {
		if err := assignInputField(input, "DesiredCapacity", _autoscalingDesiredCapacity); err != nil {
			log.Errorf("invalid --desired-capacity: %s", err.Error())
			return
		}
	}
	if len(_autoscalingEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _autoscalingEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_autoscalingMaxSize) > 0 {
		if err := assignInputField(input, "MaxSize", _autoscalingMaxSize); err != nil {
			log.Errorf("invalid --max-size: %s", err.Error())
			return
		}
	}
	if len(_autoscalingMinSize) > 0 {
		if err := assignInputField(input, "MinSize", _autoscalingMinSize); err != nil {
			log.Errorf("invalid --min-size: %s", err.Error())
			return
		}
	}
	if len(_autoscalingRecurrence) > 0 {
		input.Recurrence = aws.String(_autoscalingRecurrence)
	}
	if len(_autoscalingStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _autoscalingStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_autoscalingTime) > 0 {
		if err := assignInputField(input, "Time", _autoscalingTime); err != nil {
			log.Errorf("invalid --time: %s", err.Error())
			return
		}
	}
	if len(_autoscalingTimeZone) > 0 {
		input.TimeZone = aws.String(_autoscalingTimeZone)
	}

	if resp, err := client.PutScheduledUpdateGroupAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a warm pool for the specified Auto Scaling group. A warm
// pool is a pool of pre-initialized EC2 instances that sits alongside the Auto
// Scaling group. Whenever your application needs to scale out, the Auto Scaling
// group can draw on the warm pool to meet its new desired capacity.
//
// This operation must be called from the Region in which the Auto Scaling group
// was created.
//
// You can view the instances in the warm pool using the [DescribeWarmPool] API call. If you are no
// longer using a warm pool, you can delete it by calling the [DeleteWarmPool]API.
//
// For more information, see [Warm pools for Amazon EC2 Auto Scaling] in the Amazon EC2 Auto Scaling User Guide.
//
// [DeleteWarmPool]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DeleteWarmPool.html
// [DescribeWarmPool]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeWarmPool.html
// [Warm pools for Amazon EC2 Auto Scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-warm-pools.html
func autoscaling_PutWarmPool(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.PutWarmPoolInput{
		// AutoScalingGroupName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingInstanceReusePolicy) > 0 {
		if err := assignInputField(input, "InstanceReusePolicy", _autoscalingInstanceReusePolicy); err != nil {
			log.Errorf("invalid --instance-reuse-policy: %s", err.Error())
			return
		}
	}
	if len(_autoscalingMaxGroupPreparedCapacity) > 0 {
		if err := assignInputField(input, "MaxGroupPreparedCapacity", _autoscalingMaxGroupPreparedCapacity); err != nil {
			log.Errorf("invalid --max-group-prepared-capacity: %s", err.Error())
			return
		}
	}
	if len(_autoscalingMinSize) > 0 {
		if err := assignInputField(input, "MinSize", _autoscalingMinSize); err != nil {
			log.Errorf("invalid --min-size: %s", err.Error())
			return
		}
	}
	if len(_autoscalingPoolState) > 0 {
		if err := assignInputField(input, "PoolState", _autoscalingPoolState); err != nil {
			log.Errorf("invalid --pool-state: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutWarmPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Records a heartbeat for the lifecycle action associated with the specified
// token or instance. This extends the timeout by the length of time defined using
// the [PutLifecycleHook]API call.
//
// This step is a part of the procedure for adding a lifecycle hook to an Auto
// Scaling group:
//
// - (Optional) Create a launch template or launch configuration with a user
// data script that runs while an instance is in a wait state due to a lifecycle
// hook.
//
// - (Optional) Create a Lambda function and a rule that allows Amazon
// EventBridge to invoke your Lambda function when an instance is put into a wait
// state due to a lifecycle hook.
//
// - (Optional) Create a notification target and an IAM role. The target can be
// either an Amazon SQS queue or an Amazon SNS topic. The role allows Amazon EC2
// Auto Scaling to publish lifecycle notifications to the target.
//
// - Create the lifecycle hook. Specify whether the hook is used when the
// instances launch or terminate.
//
// - If you need more time, record the lifecycle action heartbeat to keep the
// instance in a wait state.
//
// - If you finish before the timeout period ends, send a callback by using the [CompleteLifecycleAction]
// API call.
//
// For more information, see [Amazon EC2 Auto Scaling lifecycle hooks] in the Amazon EC2 Auto Scaling User Guide.
//
// [CompleteLifecycleAction]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_CompleteLifecycleAction.html
// [Amazon EC2 Auto Scaling lifecycle hooks]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/lifecycle-hooks.html
// [PutLifecycleHook]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_PutLifecycleHook.html
func autoscaling_RecordLifecycleActionHeartbeat(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.RecordLifecycleActionHeartbeatInput{
		// AutoScalingGroupName: *string, // Required
		// LifecycleHookName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingLifecycleHookName) > 0 {
		input.LifecycleHookName = aws.String(_autoscalingLifecycleHookName)
	}
	if len(_autoscalingInstanceId) > 0 {
		input.InstanceId = aws.String(_autoscalingInstanceId)
	}
	if len(_autoscalingLifecycleActionToken) > 0 {
		input.LifecycleActionToken = aws.String(_autoscalingLifecycleActionToken)
	}

	if resp, err := client.RecordLifecycleActionHeartbeat(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resumes the specified suspended auto scaling processes, or all suspended
// process, for the specified Auto Scaling group.
//
// For more information, see [Suspend and resume Amazon EC2 Auto Scaling processes] in the Amazon EC2 Auto Scaling User Guide.
//
// [Suspend and resume Amazon EC2 Auto Scaling processes]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-suspend-resume-processes.html
func autoscaling_ResumeProcesses(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.ResumeProcessesInput{
		// AutoScalingGroupName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingScalingProcesses) > 0 {
		input.ScalingProcesses = append([]string(nil), _autoscalingScalingProcesses...)
	}

	if resp, err := client.ResumeProcesses(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels an instance refresh that is in progress and rolls back any changes that
// it made. Amazon EC2 Auto Scaling replaces any instances that were replaced
// during the instance refresh. This restores your Auto Scaling group to the
// configuration that it was using before the start of the instance refresh.
//
// This operation is part of the [instance refresh feature] in Amazon EC2 Auto Scaling, which helps you
// update instances in your Auto Scaling group after you make configuration
// changes.
//
// A rollback is not supported in the following situations:
//
// - There is no desired configuration specified for the instance refresh.
//
// - The Auto Scaling group has a launch template that uses an Amazon Web
// Services Systems Manager parameter instead of an AMI ID for the ImageId
// property.
//
// - The Auto Scaling group uses the launch template's $Latest or $Default
// version.
//
// When you receive a successful response from this operation, Amazon EC2 Auto
// Scaling immediately begins replacing instances. You can check the status of this
// operation through the [DescribeInstanceRefreshes]API operation.
//
// [instance refresh feature]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-refresh.html
// [DescribeInstanceRefreshes]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeInstanceRefreshes.html
func autoscaling_RollbackInstanceRefresh(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.RollbackInstanceRefreshInput{
		// AutoScalingGroupName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}

	if resp, err := client.RollbackInstanceRefresh(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the size of the specified Auto Scaling group.
// If a scale-in activity occurs as a result of a new DesiredCapacity value that
// is lower than the current size of the group, the Auto Scaling group uses its
// termination policy to determine which instances to terminate.
//
// For more information, see [Manual scaling] in the Amazon EC2 Auto Scaling User Guide.
//
// [Manual scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-scaling-manually.html
func autoscaling_SetDesiredCapacity(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.SetDesiredCapacityInput{
		// AutoScalingGroupName: *string, // Required
		// DesiredCapacity: *int32, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingDesiredCapacity) > 0 {
		if err := assignInputField(input, "DesiredCapacity", _autoscalingDesiredCapacity); err != nil {
			log.Errorf("invalid --desired-capacity: %s", err.Error())
			return
		}
	}
	if len(_autoscalingHonorCooldown) > 0 {
		if err := assignInputField(input, "HonorCooldown", _autoscalingHonorCooldown); err != nil {
			log.Errorf("invalid --honor-cooldown: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetDesiredCapacity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the health status of the specified instance.
// For more information, see [Set up a custom health check for your Auto Scaling group] in the Amazon EC2 Auto Scaling User Guide.
//
// [Set up a custom health check for your Auto Scaling group]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/set-up-a-custom-health-check.html
func autoscaling_SetInstanceHealth(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.SetInstanceHealthInput{
		// HealthStatus: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_autoscalingHealthStatus) > 0 {
		input.HealthStatus = aws.String(_autoscalingHealthStatus)
	}
	if len(_autoscalingInstanceId) > 0 {
		input.InstanceId = aws.String(_autoscalingInstanceId)
	}
	if len(_autoscalingShouldRespectGracePeriod) > 0 {
		if err := assignInputField(input, "ShouldRespectGracePeriod", _autoscalingShouldRespectGracePeriod); err != nil {
			log.Errorf("invalid --should-respect-grace-period: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetInstanceHealth(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the instance protection settings of the specified instances. This
// operation cannot be called on instances in a warm pool.
//
// For more information, see [Use instance scale-in protection] in the Amazon EC2 Auto Scaling User Guide.
//
// If you exceed your maximum limit of instance IDs, which is 50 per Auto Scaling
// group, the call fails.
//
// [Use instance scale-in protection]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-instance-protection.html
func autoscaling_SetInstanceProtection(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.SetInstanceProtectionInput{
		// AutoScalingGroupName: *string, // Required
		// InstanceIds: []string, // Required
		// ProtectedFromScaleIn: *bool, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingInstanceIds) > 0 {
		input.InstanceIds = append([]string(nil), _autoscalingInstanceIds...)
	}
	if len(_autoscalingProtectedFromScaleIn) > 0 {
		if err := assignInputField(input, "ProtectedFromScaleIn", _autoscalingProtectedFromScaleIn); err != nil {
			log.Errorf("invalid --protected-from-scale-in: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetInstanceProtection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an instance refresh.
// This operation is part of the [instance refresh feature] in Amazon EC2 Auto Scaling, which helps you
// update instances in your Auto Scaling group. This feature is helpful, for
// example, when you have a new AMI or a new user data script. You just need to
// create a new launch template that specifies the new AMI or user data script.
// Then start an instance refresh to immediately begin the process of updating
// instances in the group.
//
// If successful, the request's response contains a unique ID that you can use to
// track the progress of the instance refresh. To query its status, call the [DescribeInstanceRefreshes]API.
// To describe the instance refreshes that have already run, call the [DescribeInstanceRefreshes]API. To
// cancel an instance refresh that is in progress, use the [CancelInstanceRefresh]API.
//
// An instance refresh might fail for several reasons, such as EC2 launch
// failures, misconfigured health checks, or not ignoring or allowing the
// termination of instances that are in Standby state or protected from scale in.
// You can monitor for failed EC2 launches using the scaling activities. To find
// the scaling activities, call the [DescribeScalingActivities]API.
//
// If you enable auto rollback, your Auto Scaling group will be rolled back
// automatically when the instance refresh fails. You can enable this feature
// before starting an instance refresh by specifying the AutoRollback property in
// the instance refresh preferences. Otherwise, to roll back an instance refresh
// before it finishes, use the [RollbackInstanceRefresh]API.
//
// [DescribeScalingActivities]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeScalingActivities.html
// [instance refresh feature]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-refresh.html
// [DescribeInstanceRefreshes]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeInstanceRefreshes.html
// [CancelInstanceRefresh]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_CancelInstanceRefresh.html
// [RollbackInstanceRefresh]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_RollbackInstanceRefresh.html
func autoscaling_StartInstanceRefresh(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.StartInstanceRefreshInput{
		// AutoScalingGroupName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingDesiredConfiguration) > 0 {
		if err := assignInputField(input, "DesiredConfiguration", _autoscalingDesiredConfiguration); err != nil {
			log.Errorf("invalid --desired-configuration: %s", err.Error())
			return
		}
	}
	if len(_autoscalingPreferences) > 0 {
		if err := assignInputField(input, "Preferences", _autoscalingPreferences); err != nil {
			log.Errorf("invalid --preferences: %s", err.Error())
			return
		}
	}
	if len(_autoscalingStrategy) > 0 {
		if err := assignInputField(input, "Strategy", _autoscalingStrategy); err != nil {
			log.Errorf("invalid --strategy: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartInstanceRefresh(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Suspends the specified auto scaling processes, or all processes, for the
// specified Auto Scaling group.
//
// If you suspend either the Launch or Terminate process types, it can prevent
// other process types from functioning properly. For more information, see [Suspend and resume Amazon EC2 Auto Scaling processes]in the
// Amazon EC2 Auto Scaling User Guide.
//
// To resume processes that have been suspended, call the [ResumeProcesses] API.
//
// [ResumeProcesses]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_ResumeProcesses.html
// [Suspend and resume Amazon EC2 Auto Scaling processes]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-suspend-resume-processes.html
func autoscaling_SuspendProcesses(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.SuspendProcessesInput{
		// AutoScalingGroupName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingScalingProcesses) > 0 {
		input.ScalingProcesses = append([]string(nil), _autoscalingScalingProcesses...)
	}

	if resp, err := client.SuspendProcesses(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Terminates the specified instance and optionally adjusts the desired group
// size. This operation cannot be called on instances in a warm pool.
//
// This call simply makes a termination request. The instance is not terminated
// immediately. When an instance is terminated, the instance status changes to
// terminated . You can't connect to or start an instance after you've terminated
// it.
//
// If you do not specify the option to decrement the desired capacity, Amazon EC2
// Auto Scaling launches instances to replace the ones that are terminated.
//
// By default, Amazon EC2 Auto Scaling balances instances across all Availability
// Zones. If you decrement the desired capacity, your Auto Scaling group can become
// unbalanced between Availability Zones. Amazon EC2 Auto Scaling tries to
// rebalance the group, and rebalancing might terminate instances in other zones.
// For more information, see [Manual scaling]in the Amazon EC2 Auto Scaling User Guide.
//
// [Manual scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-scaling-manually.html
func autoscaling_TerminateInstanceInAutoScalingGroup(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.TerminateInstanceInAutoScalingGroupInput{
		// InstanceId: *string, // Required
		// ShouldDecrementDesiredCapacity: *bool, // Required
	}

	if len(_autoscalingInstanceId) > 0 {
		input.InstanceId = aws.String(_autoscalingInstanceId)
	}
	if len(_autoscalingShouldDecrementDesiredCapacity) > 0 {
		if err := assignInputField(input, "ShouldDecrementDesiredCapacity", _autoscalingShouldDecrementDesiredCapacity); err != nil {
			log.Errorf("invalid --should-decrement-desired-capacity: %s", err.Error())
			return
		}
	}

	if resp, err := client.TerminateInstanceInAutoScalingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// We strongly recommend that all Auto Scaling groups use launch templates to
// ensure full functionality for Amazon EC2 Auto Scaling and Amazon EC2.
//
// Updates the configuration for the specified Auto Scaling group.
//
// To update an Auto Scaling group, specify the name of the group and the property
// that you want to change. Any properties that you don't specify are not changed
// by this update request. The new settings take effect on any scaling activities
// after this call returns.
//
// If you associate a new launch configuration or template with an Auto Scaling
// group, all new instances will get the updated configuration. Existing instances
// continue to run with the configuration that they were originally launched with.
// When you update a group to specify a mixed instances policy instead of a launch
// configuration or template, existing instances may be replaced to match the new
// purchasing options that you specified in the policy. For example, if the group
// currently has 100% On-Demand capacity and the policy specifies 50% Spot
// capacity, this means that half of your instances will be gradually terminated
// and relaunched as Spot Instances. When replacing instances, Amazon EC2 Auto
// Scaling launches new instances before terminating the old ones, so that updating
// your group does not compromise the performance or availability of your
// application.
//
// Note the following about changing DesiredCapacity , MaxSize , or MinSize :
//
// - If a scale-in activity occurs as a result of a new DesiredCapacity value
// that is lower than the current size of the group, the Auto Scaling group uses
// its termination policy to determine which instances to terminate.
//
// - If you specify a new value for MinSize without specifying a value for
// DesiredCapacity , and the new MinSize is larger than the current size of the
// group, this sets the group's DesiredCapacity to the new MinSize value.
//
// - If you specify a new value for MaxSize without specifying a value for
// DesiredCapacity , and the new MaxSize is smaller than the current size of the
// group, this sets the group's DesiredCapacity to the new MaxSize value.
//
// To see which properties have been set, call the [DescribeAutoScalingGroups] API. To view the scaling
// policies for an Auto Scaling group, call the [DescribePolicies]API. If the group has scaling
// policies, you can update them by calling the [PutScalingPolicy]API.
//
// [DescribeAutoScalingGroups]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeAutoScalingGroups.html
// [DescribePolicies]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribePolicies.html
// [PutScalingPolicy]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_PutScalingPolicy.html
func autoscaling_UpdateAutoScalingGroup(cfg aws.Config, client *autoscaling.Client) {
	input := &autoscaling.UpdateAutoScalingGroupInput{
		// AutoScalingGroupName: *string, // Required
	}

	if len(_autoscalingAutoScalingGroupName) > 0 {
		input.AutoScalingGroupName = aws.String(_autoscalingAutoScalingGroupName)
	}
	if len(_autoscalingAvailabilityZoneDistribution) > 0 {
		if err := assignInputField(input, "AvailabilityZoneDistribution", _autoscalingAvailabilityZoneDistribution); err != nil {
			log.Errorf("invalid --availability-zone-distribution: %s", err.Error())
			return
		}
	}
	if len(_autoscalingAvailabilityZoneImpairmentPolicy) > 0 {
		if err := assignInputField(input, "AvailabilityZoneImpairmentPolicy", _autoscalingAvailabilityZoneImpairmentPolicy); err != nil {
			log.Errorf("invalid --availability-zone-impairment-policy: %s", err.Error())
			return
		}
	}
	if len(_autoscalingAvailabilityZones) > 0 {
		input.AvailabilityZones = append([]string(nil), _autoscalingAvailabilityZones...)
	}
	if len(_autoscalingCapacityRebalance) > 0 {
		if err := assignInputField(input, "CapacityRebalance", _autoscalingCapacityRebalance); err != nil {
			log.Errorf("invalid --capacity-rebalance: %s", err.Error())
			return
		}
	}
	if len(_autoscalingCapacityReservationSpecification) > 0 {
		if err := assignInputField(input, "CapacityReservationSpecification", _autoscalingCapacityReservationSpecification); err != nil {
			log.Errorf("invalid --capacity-reservation-specification: %s", err.Error())
			return
		}
	}
	if len(_autoscalingContext) > 0 {
		input.Context = aws.String(_autoscalingContext)
	}
	if len(_autoscalingDefaultCooldown) > 0 {
		if err := assignInputField(input, "DefaultCooldown", _autoscalingDefaultCooldown); err != nil {
			log.Errorf("invalid --default-cooldown: %s", err.Error())
			return
		}
	}
	if len(_autoscalingDefaultInstanceWarmup) > 0 {
		if err := assignInputField(input, "DefaultInstanceWarmup", _autoscalingDefaultInstanceWarmup); err != nil {
			log.Errorf("invalid --default-instance-warmup: %s", err.Error())
			return
		}
	}
	if len(_autoscalingDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _autoscalingDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_autoscalingDesiredCapacity) > 0 {
		if err := assignInputField(input, "DesiredCapacity", _autoscalingDesiredCapacity); err != nil {
			log.Errorf("invalid --desired-capacity: %s", err.Error())
			return
		}
	}
	if len(_autoscalingDesiredCapacityType) > 0 {
		input.DesiredCapacityType = aws.String(_autoscalingDesiredCapacityType)
	}
	if len(_autoscalingHealthCheckGracePeriod) > 0 {
		if err := assignInputField(input, "HealthCheckGracePeriod", _autoscalingHealthCheckGracePeriod); err != nil {
			log.Errorf("invalid --health-check-grace-period: %s", err.Error())
			return
		}
	}
	if len(_autoscalingHealthCheckType) > 0 {
		input.HealthCheckType = aws.String(_autoscalingHealthCheckType)
	}
	if len(_autoscalingInstanceLifecyclePolicy) > 0 {
		if err := assignInputField(input, "InstanceLifecyclePolicy", _autoscalingInstanceLifecyclePolicy); err != nil {
			log.Errorf("invalid --instance-lifecycle-policy: %s", err.Error())
			return
		}
	}
	if len(_autoscalingInstanceMaintenancePolicy) > 0 {
		if err := assignInputField(input, "InstanceMaintenancePolicy", _autoscalingInstanceMaintenancePolicy); err != nil {
			log.Errorf("invalid --instance-maintenance-policy: %s", err.Error())
			return
		}
	}
	if len(_autoscalingLaunchConfigurationName) > 0 {
		input.LaunchConfigurationName = aws.String(_autoscalingLaunchConfigurationName)
	}
	if len(_autoscalingLaunchTemplate) > 0 {
		if err := assignInputField(input, "LaunchTemplate", _autoscalingLaunchTemplate); err != nil {
			log.Errorf("invalid --launch-template: %s", err.Error())
			return
		}
	}
	if len(_autoscalingMaxInstanceLifetime) > 0 {
		if err := assignInputField(input, "MaxInstanceLifetime", _autoscalingMaxInstanceLifetime); err != nil {
			log.Errorf("invalid --max-instance-lifetime: %s", err.Error())
			return
		}
	}
	if len(_autoscalingMaxSize) > 0 {
		if err := assignInputField(input, "MaxSize", _autoscalingMaxSize); err != nil {
			log.Errorf("invalid --max-size: %s", err.Error())
			return
		}
	}
	if len(_autoscalingMinSize) > 0 {
		if err := assignInputField(input, "MinSize", _autoscalingMinSize); err != nil {
			log.Errorf("invalid --min-size: %s", err.Error())
			return
		}
	}
	if len(_autoscalingMixedInstancesPolicy) > 0 {
		if err := assignInputField(input, "MixedInstancesPolicy", _autoscalingMixedInstancesPolicy); err != nil {
			log.Errorf("invalid --mixed-instances-policy: %s", err.Error())
			return
		}
	}
	if len(_autoscalingNewInstancesProtectedFromScaleIn) > 0 {
		if err := assignInputField(input, "NewInstancesProtectedFromScaleIn", _autoscalingNewInstancesProtectedFromScaleIn); err != nil {
			log.Errorf("invalid --new-instances-protected-from-scale-in: %s", err.Error())
			return
		}
	}
	if len(_autoscalingPlacementGroup) > 0 {
		input.PlacementGroup = aws.String(_autoscalingPlacementGroup)
	}
	if len(_autoscalingServiceLinkedRoleARN) > 0 {
		input.ServiceLinkedRoleARN = aws.String(_autoscalingServiceLinkedRoleARN)
	}
	if len(_autoscalingSkipZonalShiftValidation) > 0 {
		if err := assignInputField(input, "SkipZonalShiftValidation", _autoscalingSkipZonalShiftValidation); err != nil {
			log.Errorf("invalid --skip-zonal-shift-validation: %s", err.Error())
			return
		}
	}
	if len(_autoscalingTerminationPolicies) > 0 {
		input.TerminationPolicies = append([]string(nil), _autoscalingTerminationPolicies...)
	}
	if len(_autoscalingVPCZoneIdentifier) > 0 {
		input.VPCZoneIdentifier = aws.String(_autoscalingVPCZoneIdentifier)
	}

	if resp, err := client.UpdateAutoScalingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_autoscalingCmd)
	_autoscalingCmd.Flags().SortFlags = false

	_autoscalingCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_autoscalingCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_autoscalingCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_autoscalingCmd.Flags().StringSliceVarP(&_autoscalingActivityIds, "activity-ids", "", nil, "Activity Ids")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingAdjustmentType, "adjustment-type", "", "", "Adjustment Type")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingAssociatePublicIpAddress, "associate-public-ip-address", "", "", "Associate Public IP Address")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingAutoScalingGroupName, "auto-scaling-group-name", "", "", "Auto Scaling Group Name")
	_autoscalingCmd.Flags().StringSliceVarP(&_autoscalingAutoScalingGroupNames, "auto-scaling-group-names", "", nil, "Auto Scaling Group Names")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingAvailabilityZoneDistribution, "availability-zone-distribution", "", "", "Availability Zone Distribution")
	_autoscalingCmd.Flags().StringSliceVarP(&_autoscalingAvailabilityZoneIds, "availability-zone-ids", "", nil, "Availability Zone Ids")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingAvailabilityZoneImpairmentPolicy, "availability-zone-impairment-policy", "", "", "Availability Zone Impairment Policy")
	_autoscalingCmd.Flags().StringSliceVarP(&_autoscalingAvailabilityZones, "availability-zones", "", nil, "Availability Zones")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingBlockDeviceMappings, "block-device-mappings", "", "", "Block Device Mappings")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingBreachThreshold, "breach-threshold", "", "", "Breach Threshold")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingCapacityRebalance, "capacity-rebalance", "", "", "Capacity Rebalance")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingCapacityReservationSpecification, "capacity-reservation-specification", "", "", "Capacity Reservation Specification")
	_autoscalingCmd.Flags().StringSliceVarP(&_autoscalingClassicLinkVPCSecurityGroups, "classic-link-vpc-security-groups", "", nil, "Classic Link VPC Security Groups")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingClassicLinkVPCId, "classic-link-vpcid", "", "", "Classic Link Vpcid")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingClientToken, "client-token", "", "", "Client Token")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingContext, "context", "", "", "Context")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingCooldown, "cooldown", "", "", "Cooldown")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingDefaultCooldown, "default-cooldown", "", "", "Default Cooldown")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingDefaultInstanceWarmup, "default-instance-warmup", "", "", "Default Instance Warmup")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingDefaultResult, "default-result", "", "", "Default Result")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingDeletionProtection, "deletion-protection", "", "", "Deletion Protection")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingDesiredCapacity, "desired-capacity", "", "", "Desired Capacity")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingDesiredCapacityType, "desired-capacity-type", "", "", "Desired Capacity Type")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingDesiredConfiguration, "desired-configuration", "", "", "Desired Configuration")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingEbsOptimized, "ebs-optimized", "", "", "Ebs Optimized")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingEnabled, "enabled", "", "", "Enabled")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingEndTime, "end-time", "", "", "End Time")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingEstimatedInstanceWarmup, "estimated-instance-warmup", "", "", "Estimated Instance Warmup")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingFilters, "filters", "", "", "Filters")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingForceDelete, "force-delete", "", "", "Force Delete")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingGranularity, "granularity", "", "", "Granularity")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingHealthCheckGracePeriod, "health-check-grace-period", "", "", "Health Check Grace Period")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingHealthCheckType, "health-check-type", "", "", "Health Check Type")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingHealthStatus, "health-status", "", "", "Health Status")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingHeartbeatTimeout, "heartbeat-timeout", "", "", "Heartbeat Timeout")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingHonorCooldown, "honor-cooldown", "", "", "Honor Cooldown")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingIamInstanceProfile, "iam-instance-profile", "", "", "IAM Instance Profile")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingImageId, "image-id", "", "", "Image ID")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingIncludeDeletedGroups, "include-deleted-groups", "", "", "Include Deleted Groups")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingIncludeInstances, "include-instances", "", "", "Include Instances")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingInstanceId, "instance-id", "", "", "Instance ID")
	_autoscalingCmd.Flags().StringSliceVarP(&_autoscalingInstanceIds, "instance-ids", "", nil, "Instance Ids")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingInstanceLifecyclePolicy, "instance-lifecycle-policy", "", "", "Instance Lifecycle Policy")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingInstanceMaintenancePolicy, "instance-maintenance-policy", "", "", "Instance Maintenance Policy")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingInstanceMonitoring, "instance-monitoring", "", "", "Instance Monitoring")
	_autoscalingCmd.Flags().StringSliceVarP(&_autoscalingInstanceRefreshIds, "instance-refresh-ids", "", nil, "Instance Refresh Ids")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingInstanceReusePolicy, "instance-reuse-policy", "", "", "Instance Reuse Policy")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingInstanceType, "instance-type", "", "", "Instance Type")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingKernelId, "kernel-id", "", "", "Kernel ID")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingKeyName, "key-name", "", "", "Key Name")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingLaunchConfigurationName, "launch-configuration-name", "", "", "Launch Configuration Name")
	_autoscalingCmd.Flags().StringSliceVarP(&_autoscalingLaunchConfigurationNames, "launch-configuration-names", "", nil, "Launch Configuration Names")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingLaunchTemplate, "launch-template", "", "", "Launch Template")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingLifecycleActionResult, "lifecycle-action-result", "", "", "Lifecycle Action Result")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingLifecycleActionToken, "lifecycle-action-token", "", "", "Lifecycle Action Token")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingLifecycleHookName, "lifecycle-hook-name", "", "", "Lifecycle Hook Name")
	_autoscalingCmd.Flags().StringSliceVarP(&_autoscalingLifecycleHookNames, "lifecycle-hook-names", "", nil, "Lifecycle Hook Names")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingLifecycleHookSpecificationList, "lifecycle-hook-specification-list", "", "", "Lifecycle Hook Specification List")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingLifecycleTransition, "lifecycle-transition", "", "", "Lifecycle Transition")
	_autoscalingCmd.Flags().StringSliceVarP(&_autoscalingLoadBalancerNames, "load-balancer-names", "", nil, "Load Balancer Names")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingMaxGroupPreparedCapacity, "max-group-prepared-capacity", "", "", "Max Group Prepared Capacity")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingMaxInstanceLifetime, "max-instance-lifetime", "", "", "Max Instance Lifetime")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingMaxRecords, "max-records", "", "", "Max Records")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingMaxSize, "max-size", "", "", "Max Size")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingMetadataOptions, "metadata-options", "", "", "Metadata Options")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingMetricAggregationType, "metric-aggregation-type", "", "", "Metric Aggregation Type")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingMetricValue, "metric-value", "", "", "Metric Value")
	_autoscalingCmd.Flags().StringSliceVarP(&_autoscalingMetrics, "metrics", "", nil, "Metrics")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingMinAdjustmentMagnitude, "min-adjustment-magnitude", "", "", "Min Adjustment Magnitude")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingMinAdjustmentStep, "min-adjustment-step", "", "", "Min Adjustment Step")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingMinSize, "min-size", "", "", "Min Size")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingMixedInstancesPolicy, "mixed-instances-policy", "", "", "Mixed Instances Policy")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingNewInstancesProtectedFromScaleIn, "new-instances-protected-from-scale-in", "", "", "New Instances Protected From Scale In")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingNextToken, "next-token", "", "", "Next Token")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingNotificationMetadata, "notification-metadata", "", "", "Notification Metadata")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingNotificationTargetARN, "notification-target-arn", "", "", "Notification Target ARN")
	_autoscalingCmd.Flags().StringSliceVarP(&_autoscalingNotificationTypes, "notification-types", "", nil, "Notification Types")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingPlacementGroup, "placement-group", "", "", "Placement Group")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingPlacementTenancy, "placement-tenancy", "", "", "Placement Tenancy")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingPolicyName, "policy-name", "", "", "Policy Name")
	_autoscalingCmd.Flags().StringSliceVarP(&_autoscalingPolicyNames, "policy-names", "", nil, "Policy Names")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingPolicyType, "policy-type", "", "", "Policy Type")
	_autoscalingCmd.Flags().StringSliceVarP(&_autoscalingPolicyTypes, "policy-types", "", nil, "Policy Types")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingPoolState, "pool-state", "", "", "Pool State")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingPredictiveScalingConfiguration, "predictive-scaling-configuration", "", "", "Predictive Scaling Configuration")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingPreferences, "preferences", "", "", "Preferences")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingProtectedFromScaleIn, "protected-from-scale-in", "", "", "Protected From Scale In")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingRamdiskId, "ramdisk-id", "", "", "Ramdisk ID")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingRecurrence, "recurrence", "", "", "Recurrence")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingRequestedCapacity, "requested-capacity", "", "", "Requested Capacity")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingRetryStrategy, "retry-strategy", "", "", "Retry Strategy")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingRoleARN, "role-arn", "", "", "Role ARN")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingScalingAdjustment, "scaling-adjustment", "", "", "Scaling Adjustment")
	_autoscalingCmd.Flags().StringSliceVarP(&_autoscalingScalingProcesses, "scaling-processes", "", nil, "Scaling Processes")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingScheduledActionName, "scheduled-action-name", "", "", "Scheduled Action Name")
	_autoscalingCmd.Flags().StringSliceVarP(&_autoscalingScheduledActionNames, "scheduled-action-names", "", nil, "Scheduled Action Names")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingScheduledUpdateGroupActions, "scheduled-update-group-actions", "", "", "Scheduled Update Group Actions")
	_autoscalingCmd.Flags().StringSliceVarP(&_autoscalingSecurityGroups, "security-groups", "", nil, "Security Groups")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingServiceLinkedRoleARN, "service-linked-role-arn", "", "", "Service Linked Role ARN")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingShouldDecrementDesiredCapacity, "should-decrement-desired-capacity", "", "", "Should Decrement Desired Capacity")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingShouldRespectGracePeriod, "should-respect-grace-period", "", "", "Should Respect Grace Period")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingSkipZonalShiftValidation, "skip-zonal-shift-validation", "", "", "Skip Zonal Shift Validation")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingSpotPrice, "spot-price", "", "", "Spot Price")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingStartTime, "start-time", "", "", "Start Time")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingStepAdjustments, "step-adjustments", "", "", "Step Adjustments")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingStrategy, "strategy", "", "", "Strategy")
	_autoscalingCmd.Flags().StringSliceVarP(&_autoscalingSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingTags, "tags", "", "", "Tags")
	_autoscalingCmd.Flags().StringSliceVarP(&_autoscalingTargetGroupARNs, "target-group-arns", "", nil, "Target Group Arns")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingTargetTrackingConfiguration, "target-tracking-configuration", "", "", "Target Tracking Configuration")
	_autoscalingCmd.Flags().StringSliceVarP(&_autoscalingTerminationPolicies, "termination-policies", "", nil, "Termination Policies")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingTime, "time", "", "", "Time")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingTimeZone, "time-zone", "", "", "Time Zone")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingTopicARN, "topic-arn", "", "", "Topic ARN")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingTrafficSourceType, "traffic-source-type", "", "", "Traffic Source Type")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingTrafficSources, "traffic-sources", "", "", "Traffic Sources")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingUserData, "user-data", "", "", "User Data")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingVPCZoneIdentifier, "vpc-zone-identifier", "", "", "VPC Zone Identifier")
	_autoscalingCmd.Flags().StringVarP(&_autoscalingWaitForTransitioningInstances, "wait-for-transitioning-instances", "", "", "Wait For Transitioning Instances")

	_autoscalingCmd.Flags().BoolVarP(&_autoscalingAttachInstances, "attach-instances", "", false, "Attach Instances")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingAttachLoadBalancerTargetGroups, "attach-load-balancer-target-groups", "", false, "Attach Load Balancer Target Groups")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingAttachLoadBalancers, "attach-load-balancers", "", false, "Attach Load Balancers")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingAttachTrafficSources, "attach-traffic-sources", "", false, "Attach Traffic Sources")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingBatchDeleteScheduledAction, "batch-delete-scheduled-action", "", false, "Batch Delete Scheduled Action")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingBatchPutScheduledUpdateGroupAction, "batch-put-scheduled-update-group-action", "", false, "Batch Put Scheduled Update Group Action")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingCancelInstanceRefresh, "cancel-instance-refresh", "", false, "Cancel Instance Refresh")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingCompleteLifecycleAction, "complete-lifecycle-action", "", false, "Complete Lifecycle Action")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingCreateAutoScalingGroup, "create-auto-scaling-group", "", false, "Create Auto Scaling Group")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingCreateLaunchConfiguration, "create-launch-configuration", "", false, "Create Launch Configuration")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingCreateOrUpdateTags, "create-or-update-tags", "", false, "Create Or Update Tags")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDeleteAutoScalingGroup, "delete-auto-scaling-group", "", false, "Delete Auto Scaling Group")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDeleteLaunchConfiguration, "delete-launch-configuration", "", false, "Delete Launch Configuration")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDeleteLifecycleHook, "delete-lifecycle-hook", "", false, "Delete Lifecycle Hook")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDeleteNotificationConfiguration, "delete-notification-configuration", "", false, "Delete Notification Configuration")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDeletePolicy, "delete-policy", "", false, "Delete Policy")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDeleteScheduledAction, "delete-scheduled-action", "", false, "Delete Scheduled Action")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDeleteTags, "delete-tags", "", false, "Delete Tags")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDeleteWarmPool, "delete-warm-pool", "", false, "Delete Warm Pool")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDescribeAccountLimits, "describe-account-limits", "", false, "Describe Account Limits")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDescribeAdjustmentTypes, "describe-adjustment-types", "", false, "Describe Adjustment Types")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDescribeAutoScalingGroups, "describe-auto-scaling-groups", "", false, "Describe Auto Scaling Groups")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDescribeAutoScalingInstances, "describe-auto-scaling-instances", "", false, "Describe Auto Scaling Instances")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDescribeAutoScalingNotificationTypes, "describe-auto-scaling-notification-types", "", false, "Describe Auto Scaling Notification Types")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDescribeInstanceRefreshes, "describe-instance-refreshes", "", false, "Describe Instance Refreshes")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDescribeLaunchConfigurations, "describe-launch-configurations", "", false, "Describe Launch Configurations")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDescribeLifecycleHookTypes, "describe-lifecycle-hook-types", "", false, "Describe Lifecycle Hook Types")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDescribeLifecycleHooks, "describe-lifecycle-hooks", "", false, "Describe Lifecycle Hooks")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDescribeLoadBalancerTargetGroups, "describe-load-balancer-target-groups", "", false, "Describe Load Balancer Target Groups")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDescribeLoadBalancers, "describe-load-balancers", "", false, "Describe Load Balancers")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDescribeMetricCollectionTypes, "describe-metric-collection-types", "", false, "Describe Metric Collection Types")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDescribeNotificationConfigurations, "describe-notification-configurations", "", false, "Describe Notification Configurations")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDescribePolicies, "describe-policies", "", false, "Describe Policies")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDescribeScalingActivities, "describe-scaling-activities", "", false, "Describe Scaling Activities")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDescribeScalingProcessTypes, "describe-scaling-process-types", "", false, "Describe Scaling Process Types")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDescribeScheduledActions, "describe-scheduled-actions", "", false, "Describe Scheduled Actions")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDescribeTags, "describe-tags", "", false, "Describe Tags")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDescribeTerminationPolicyTypes, "describe-termination-policy-types", "", false, "Describe Termination Policy Types")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDescribeTrafficSources, "describe-traffic-sources", "", false, "Describe Traffic Sources")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDescribeWarmPool, "describe-warm-pool", "", false, "Describe Warm Pool")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDetachInstances, "detach-instances", "", false, "Detach Instances")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDetachLoadBalancerTargetGroups, "detach-load-balancer-target-groups", "", false, "Detach Load Balancer Target Groups")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDetachLoadBalancers, "detach-load-balancers", "", false, "Detach Load Balancers")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDetachTrafficSources, "detach-traffic-sources", "", false, "Detach Traffic Sources")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingDisableMetricsCollection, "disable-metrics-collection", "", false, "Disable Metrics Collection")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingEnableMetricsCollection, "enable-metrics-collection", "", false, "Enable Metrics Collection")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingEnterStandby, "enter-standby", "", false, "Enter Standby")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingExecutePolicy, "execute-policy", "", false, "Execute Policy")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingExitStandby, "exit-standby", "", false, "Exit Standby")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingGetPredictiveScalingForecast, "get-predictive-scaling-forecast", "", false, "Get Predictive Scaling Forecast")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingLaunchInstances, "launch-instances", "", false, "Launch Instances")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingPutLifecycleHook, "put-lifecycle-hook", "", false, "Put Lifecycle Hook")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingPutNotificationConfiguration, "put-notification-configuration", "", false, "Put Notification Configuration")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingPutScalingPolicy, "put-scaling-policy", "", false, "Put Scaling Policy")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingPutScheduledUpdateGroupAction, "put-scheduled-update-group-action", "", false, "Put Scheduled Update Group Action")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingPutWarmPool, "put-warm-pool", "", false, "Put Warm Pool")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingRecordLifecycleActionHeartbeat, "record-lifecycle-action-heartbeat", "", false, "Record Lifecycle Action Heartbeat")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingResumeProcesses, "resume-processes", "", false, "Resume Processes")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingRollbackInstanceRefresh, "rollback-instance-refresh", "", false, "Rollback Instance Refresh")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingSetDesiredCapacity, "set-desired-capacity", "", false, "Set Desired Capacity")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingSetInstanceHealth, "set-instance-health", "", false, "Set Instance Health")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingSetInstanceProtection, "set-instance-protection", "", false, "Set Instance Protection")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingStartInstanceRefresh, "start-instance-refresh", "", false, "Start Instance Refresh")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingSuspendProcesses, "suspend-processes", "", false, "Suspend Processes")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingTerminateInstanceInAutoScalingGroup, "terminate-instance-in-auto-scaling-group", "", false, "Terminate Instance In Auto Scaling Group")
	_autoscalingCmd.Flags().BoolVarP(&_autoscalingUpdateAutoScalingGroup, "update-auto-scaling-group", "", false, "Update Auto Scaling Group")

}
