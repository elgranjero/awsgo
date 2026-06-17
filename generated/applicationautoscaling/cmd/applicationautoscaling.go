package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// applicationautoscalingCmd represents the applicationautoscaling command
var _applicationautoscalingCmd = &cobra.Command{
	Use:   "applicationautoscaling",
	Short: "AWS applicationautoscaling CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := applicationautoscaling.NewFromConfig(cfg)
		if _applicationautoscalingDeleteScalingPolicy {
			applicationautoscaling_DeleteScalingPolicy(cfg, client)
			return
		}
		if _applicationautoscalingDeleteScheduledAction {
			applicationautoscaling_DeleteScheduledAction(cfg, client)
			return
		}
		if _applicationautoscalingDeregisterScalableTarget {
			applicationautoscaling_DeregisterScalableTarget(cfg, client)
			return
		}
		if _applicationautoscalingDescribeScalableTargets {
			applicationautoscaling_DescribeScalableTargets(cfg, client)
			return
		}
		if _applicationautoscalingDescribeScalingActivities {
			applicationautoscaling_DescribeScalingActivities(cfg, client)
			return
		}
		if _applicationautoscalingDescribeScalingPolicies {
			applicationautoscaling_DescribeScalingPolicies(cfg, client)
			return
		}
		if _applicationautoscalingDescribeScheduledActions {
			applicationautoscaling_DescribeScheduledActions(cfg, client)
			return
		}
		if _applicationautoscalingGetPredictiveScalingForecast {
			applicationautoscaling_GetPredictiveScalingForecast(cfg, client)
			return
		}
		if _applicationautoscalingListTagsForResource {
			applicationautoscaling_ListTagsForResource(cfg, client)
			return
		}
		if _applicationautoscalingPutScalingPolicy {
			applicationautoscaling_PutScalingPolicy(cfg, client)
			return
		}
		if _applicationautoscalingPutScheduledAction {
			applicationautoscaling_PutScheduledAction(cfg, client)
			return
		}
		if _applicationautoscalingRegisterScalableTarget {
			applicationautoscaling_RegisterScalableTarget(cfg, client)
			return
		}
		if _applicationautoscalingTagResource {
			applicationautoscaling_TagResource(cfg, client)
			return
		}
		if _applicationautoscalingUntagResource {
			applicationautoscaling_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_applicationautoscalingDeleteScalingPolicy          bool
	_applicationautoscalingDeleteScheduledAction        bool
	_applicationautoscalingDeregisterScalableTarget     bool
	_applicationautoscalingDescribeScalableTargets      bool
	_applicationautoscalingDescribeScalingActivities    bool
	_applicationautoscalingDescribeScalingPolicies      bool
	_applicationautoscalingDescribeScheduledActions     bool
	_applicationautoscalingGetPredictiveScalingForecast bool
	_applicationautoscalingListTagsForResource          bool
	_applicationautoscalingPutScalingPolicy             bool
	_applicationautoscalingPutScheduledAction           bool
	_applicationautoscalingRegisterScalableTarget       bool
	_applicationautoscalingTagResource                  bool
	_applicationautoscalingUntagResource                bool

	_applicationautoscalingEndTime                                  string
	_applicationautoscalingIncludeNotScaledActivities               string
	_applicationautoscalingMaxCapacity                              string
	_applicationautoscalingMaxResults                               string
	_applicationautoscalingMinCapacity                              string
	_applicationautoscalingNextToken                                string
	_applicationautoscalingPolicyName                               string
	_applicationautoscalingPolicyNames                              []string
	_applicationautoscalingPolicyType                               string
	_applicationautoscalingPredictiveScalingPolicyConfiguration     string
	_applicationautoscalingResourceARN                              string
	_applicationautoscalingResourceId                               string
	_applicationautoscalingResourceIds                              []string
	_applicationautoscalingRoleARN                                  string
	_applicationautoscalingScalableDimension                        string
	_applicationautoscalingScalableTargetAction                     string
	_applicationautoscalingSchedule                                 string
	_applicationautoscalingScheduledActionName                      string
	_applicationautoscalingScheduledActionNames                     []string
	_applicationautoscalingServiceNamespace                         string
	_applicationautoscalingStartTime                                string
	_applicationautoscalingStepScalingPolicyConfiguration           string
	_applicationautoscalingSuspendedState                           string
	_applicationautoscalingTagKeys                                  []string
	_applicationautoscalingTags                                     string
	_applicationautoscalingTargetTrackingScalingPolicyConfiguration string
	_applicationautoscalingTimezone                                 string
)

// Deletes the specified scaling policy for an Application Auto Scaling scalable
// target.
//
// Deleting a step scaling policy deletes the underlying alarm action, but does
// not delete the CloudWatch alarm associated with the scaling policy, even if it
// no longer has an associated action.
//
// For more information, see [Delete a step scaling policy] and [Delete a target tracking scaling policy] in the Application Auto Scaling User Guide.
//
// [Delete a target tracking scaling policy]: https://docs.aws.amazon.com/autoscaling/application/userguide/create-target-tracking-policy-cli.html#delete-target-tracking-policy
// [Delete a step scaling policy]: https://docs.aws.amazon.com/autoscaling/application/userguide/create-step-scaling-policy-cli.html#delete-step-scaling-policy
func applicationautoscaling_DeleteScalingPolicy(cfg aws.Config, client *applicationautoscaling.Client) {
	input := &applicationautoscaling.DeleteScalingPolicyInput{
		// PolicyName: *string, // Required
		// ResourceId: *string, // Required
		// ScalableDimension: types.ScalableDimension, // Required
		// ServiceNamespace: types.ServiceNamespace, // Required
	}

	if len(_applicationautoscalingPolicyName) > 0 {
		input.PolicyName = aws.String(_applicationautoscalingPolicyName)
	}
	if len(_applicationautoscalingResourceId) > 0 {
		input.ResourceId = aws.String(_applicationautoscalingResourceId)
	}
	if len(_applicationautoscalingScalableDimension) > 0 {
		if err := assignInputField(input, "ScalableDimension", _applicationautoscalingScalableDimension); err != nil {
			log.Errorf("invalid --scalable-dimension: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingServiceNamespace) > 0 {
		if err := assignInputField(input, "ServiceNamespace", _applicationautoscalingServiceNamespace); err != nil {
			log.Errorf("invalid --service-namespace: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteScalingPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified scheduled action for an Application Auto Scaling scalable
// target.
//
// For more information, see [Delete a scheduled action] in the Application Auto Scaling User Guide.
//
// [Delete a scheduled action]: https://docs.aws.amazon.com/autoscaling/application/userguide/scheduled-scaling-additional-cli-commands.html#delete-scheduled-action
func applicationautoscaling_DeleteScheduledAction(cfg aws.Config, client *applicationautoscaling.Client) {
	input := &applicationautoscaling.DeleteScheduledActionInput{
		// ResourceId: *string, // Required
		// ScalableDimension: types.ScalableDimension, // Required
		// ScheduledActionName: *string, // Required
		// ServiceNamespace: types.ServiceNamespace, // Required
	}

	if len(_applicationautoscalingResourceId) > 0 {
		input.ResourceId = aws.String(_applicationautoscalingResourceId)
	}
	if len(_applicationautoscalingScalableDimension) > 0 {
		if err := assignInputField(input, "ScalableDimension", _applicationautoscalingScalableDimension); err != nil {
			log.Errorf("invalid --scalable-dimension: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingScheduledActionName) > 0 {
		input.ScheduledActionName = aws.String(_applicationautoscalingScheduledActionName)
	}
	if len(_applicationautoscalingServiceNamespace) > 0 {
		if err := assignInputField(input, "ServiceNamespace", _applicationautoscalingServiceNamespace); err != nil {
			log.Errorf("invalid --service-namespace: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteScheduledAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregisters an Application Auto Scaling scalable target when you have finished
// using it. To see which resources have been registered, use [DescribeScalableTargets].
//
// Deregistering a scalable target deletes the scaling policies and the scheduled
// actions that are associated with it.
//
// [DescribeScalableTargets]: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_DescribeScalableTargets.html
func applicationautoscaling_DeregisterScalableTarget(cfg aws.Config, client *applicationautoscaling.Client) {
	input := &applicationautoscaling.DeregisterScalableTargetInput{
		// ResourceId: *string, // Required
		// ScalableDimension: types.ScalableDimension, // Required
		// ServiceNamespace: types.ServiceNamespace, // Required
	}

	if len(_applicationautoscalingResourceId) > 0 {
		input.ResourceId = aws.String(_applicationautoscalingResourceId)
	}
	if len(_applicationautoscalingScalableDimension) > 0 {
		if err := assignInputField(input, "ScalableDimension", _applicationautoscalingScalableDimension); err != nil {
			log.Errorf("invalid --scalable-dimension: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingServiceNamespace) > 0 {
		if err := assignInputField(input, "ServiceNamespace", _applicationautoscalingServiceNamespace); err != nil {
			log.Errorf("invalid --service-namespace: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeregisterScalableTarget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the scalable targets in the specified namespace.
// You can filter the results using ResourceIds and ScalableDimension .
func applicationautoscaling_DescribeScalableTargets(cfg aws.Config, client *applicationautoscaling.Client) {
	input := &applicationautoscaling.DescribeScalableTargetsInput{
		// ServiceNamespace: types.ServiceNamespace, // Required
	}

	if len(_applicationautoscalingServiceNamespace) > 0 {
		if err := assignInputField(input, "ServiceNamespace", _applicationautoscalingServiceNamespace); err != nil {
			log.Errorf("invalid --service-namespace: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationautoscalingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingNextToken) > 0 {
		input.NextToken = aws.String(_applicationautoscalingNextToken)
	}
	if len(_applicationautoscalingResourceIds) > 0 {
		input.ResourceIds = append([]string(nil), _applicationautoscalingResourceIds...)
	}
	if len(_applicationautoscalingScalableDimension) > 0 {
		if err := assignInputField(input, "ScalableDimension", _applicationautoscalingScalableDimension); err != nil {
			log.Errorf("invalid --scalable-dimension: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeScalableTargets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*applicationautoscaling.DescribeScalableTargetsOutput
	p := applicationautoscaling.NewDescribeScalableTargetsPaginator(client, input)
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

// Provides descriptive information about the scaling activities in the specified
// namespace from the previous six weeks.
//
// You can filter the results using ResourceId and ScalableDimension .
//
// For information about viewing scaling activities using the Amazon Web Services
// CLI, see [Scaling activities for Application Auto Scaling].
//
// [Scaling activities for Application Auto Scaling]: https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-scaling-activities.html
func applicationautoscaling_DescribeScalingActivities(cfg aws.Config, client *applicationautoscaling.Client) {
	input := &applicationautoscaling.DescribeScalingActivitiesInput{
		// ServiceNamespace: types.ServiceNamespace, // Required
	}

	if len(_applicationautoscalingServiceNamespace) > 0 {
		if err := assignInputField(input, "ServiceNamespace", _applicationautoscalingServiceNamespace); err != nil {
			log.Errorf("invalid --service-namespace: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingIncludeNotScaledActivities) > 0 {
		if err := assignInputField(input, "IncludeNotScaledActivities", _applicationautoscalingIncludeNotScaledActivities); err != nil {
			log.Errorf("invalid --include-not-scaled-activities: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationautoscalingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingNextToken) > 0 {
		input.NextToken = aws.String(_applicationautoscalingNextToken)
	}
	if len(_applicationautoscalingResourceId) > 0 {
		input.ResourceId = aws.String(_applicationautoscalingResourceId)
	}
	if len(_applicationautoscalingScalableDimension) > 0 {
		if err := assignInputField(input, "ScalableDimension", _applicationautoscalingScalableDimension); err != nil {
			log.Errorf("invalid --scalable-dimension: %s", err.Error())
			return
		}
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

	var results []*applicationautoscaling.DescribeScalingActivitiesOutput
	p := applicationautoscaling.NewDescribeScalingActivitiesPaginator(client, input)
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

// Describes the Application Auto Scaling scaling policies for the specified
// service namespace.
//
// You can filter the results using ResourceId , ScalableDimension , and
// PolicyNames .
//
// For more information, see [Target tracking scaling policies] and [Step scaling policies] in the Application Auto Scaling User Guide.
//
// [Step scaling policies]: https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html
// [Target tracking scaling policies]: https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html
func applicationautoscaling_DescribeScalingPolicies(cfg aws.Config, client *applicationautoscaling.Client) {
	input := &applicationautoscaling.DescribeScalingPoliciesInput{
		// ServiceNamespace: types.ServiceNamespace, // Required
	}

	if len(_applicationautoscalingServiceNamespace) > 0 {
		if err := assignInputField(input, "ServiceNamespace", _applicationautoscalingServiceNamespace); err != nil {
			log.Errorf("invalid --service-namespace: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationautoscalingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingNextToken) > 0 {
		input.NextToken = aws.String(_applicationautoscalingNextToken)
	}
	if len(_applicationautoscalingPolicyNames) > 0 {
		input.PolicyNames = append([]string(nil), _applicationautoscalingPolicyNames...)
	}
	if len(_applicationautoscalingResourceId) > 0 {
		input.ResourceId = aws.String(_applicationautoscalingResourceId)
	}
	if len(_applicationautoscalingScalableDimension) > 0 {
		if err := assignInputField(input, "ScalableDimension", _applicationautoscalingScalableDimension); err != nil {
			log.Errorf("invalid --scalable-dimension: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeScalingPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*applicationautoscaling.DescribeScalingPoliciesOutput
	p := applicationautoscaling.NewDescribeScalingPoliciesPaginator(client, input)
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

// Describes the Application Auto Scaling scheduled actions for the specified
// service namespace.
//
// You can filter the results using the ResourceId , ScalableDimension , and
// ScheduledActionNames parameters.
//
// For more information, see [Scheduled scaling] in the Application Auto Scaling User Guide.
//
// [Scheduled scaling]: https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-scheduled-scaling.html
func applicationautoscaling_DescribeScheduledActions(cfg aws.Config, client *applicationautoscaling.Client) {
	input := &applicationautoscaling.DescribeScheduledActionsInput{
		// ServiceNamespace: types.ServiceNamespace, // Required
	}

	if len(_applicationautoscalingServiceNamespace) > 0 {
		if err := assignInputField(input, "ServiceNamespace", _applicationautoscalingServiceNamespace); err != nil {
			log.Errorf("invalid --service-namespace: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationautoscalingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingNextToken) > 0 {
		input.NextToken = aws.String(_applicationautoscalingNextToken)
	}
	if len(_applicationautoscalingResourceId) > 0 {
		input.ResourceId = aws.String(_applicationautoscalingResourceId)
	}
	if len(_applicationautoscalingScalableDimension) > 0 {
		if err := assignInputField(input, "ScalableDimension", _applicationautoscalingScalableDimension); err != nil {
			log.Errorf("invalid --scalable-dimension: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingScheduledActionNames) > 0 {
		input.ScheduledActionNames = append([]string(nil), _applicationautoscalingScheduledActionNames...)
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

	var results []*applicationautoscaling.DescribeScheduledActionsOutput
	p := applicationautoscaling.NewDescribeScheduledActionsPaginator(client, input)
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

// Retrieves the forecast data for a predictive scaling policy.
// Load forecasts are predictions of the hourly load values using historical load
// data from CloudWatch and an analysis of historical trends. Capacity forecasts
// are represented as predicted values for the minimum capacity that is needed on
// an hourly basis, based on the hourly load forecast.
//
// A minimum of 24 hours of data is required to create the initial forecasts.
// However, having a full 14 days of historical data results in more accurate
// forecasts.
func applicationautoscaling_GetPredictiveScalingForecast(cfg aws.Config, client *applicationautoscaling.Client) {
	input := &applicationautoscaling.GetPredictiveScalingForecastInput{
		// EndTime: *time.Time, // Required
		// PolicyName: *string, // Required
		// ResourceId: *string, // Required
		// ScalableDimension: types.ScalableDimension, // Required
		// ServiceNamespace: types.ServiceNamespace, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_applicationautoscalingEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _applicationautoscalingEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingPolicyName) > 0 {
		input.PolicyName = aws.String(_applicationautoscalingPolicyName)
	}
	if len(_applicationautoscalingResourceId) > 0 {
		input.ResourceId = aws.String(_applicationautoscalingResourceId)
	}
	if len(_applicationautoscalingScalableDimension) > 0 {
		if err := assignInputField(input, "ScalableDimension", _applicationautoscalingScalableDimension); err != nil {
			log.Errorf("invalid --scalable-dimension: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingServiceNamespace) > 0 {
		if err := assignInputField(input, "ServiceNamespace", _applicationautoscalingServiceNamespace); err != nil {
			log.Errorf("invalid --service-namespace: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _applicationautoscalingStartTime); err != nil {
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

// Returns all the tags on the specified Application Auto Scaling scalable target.
// For general information about tags, including the format and syntax, see [Tagging your Amazon Web Services resources] in
// the Amazon Web Services General Reference.
//
// [Tagging your Amazon Web Services resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
func applicationautoscaling_ListTagsForResource(cfg aws.Config, client *applicationautoscaling.Client) {
	input := &applicationautoscaling.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_applicationautoscalingResourceARN) > 0 {
		input.ResourceARN = aws.String(_applicationautoscalingResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a scaling policy for an Application Auto Scaling scalable
// target.
//
// Each scalable target is identified by a service namespace, resource ID, and
// scalable dimension. A scaling policy applies to the scalable target identified
// by those three attributes. You cannot create a scaling policy until you have
// registered the resource as a scalable target.
//
// Multiple scaling policies can be in force at the same time for the same
// scalable target. You can have one or more target tracking scaling policies, one
// or more step scaling policies, or both. However, there is a chance that multiple
// policies could conflict, instructing the scalable target to scale out or in at
// the same time. Application Auto Scaling gives precedence to the policy that
// provides the largest capacity for both scale out and scale in. For example, if
// one policy increases capacity by 3, another policy increases capacity by 200
// percent, and the current capacity is 10, Application Auto Scaling uses the
// policy with the highest calculated capacity (200% of 10 = 20) and scales out to
// 30.
//
// We recommend caution, however, when using target tracking scaling policies with
// step scaling policies because conflicts between these policies can cause
// undesirable behavior. For example, if the step scaling policy initiates a
// scale-in activity before the target tracking policy is ready to scale in, the
// scale-in activity will not be blocked. After the scale-in activity completes,
// the target tracking policy could instruct the scalable target to scale out
// again.
//
// For more information, see [Target tracking scaling policies], [Step scaling policies], and [Predictive scaling policies] in the Application Auto Scaling User Guide.
//
// If a scalable target is deregistered, the scalable target is no longer
// available to use scaling policies. Any scaling policies that were specified for
// the scalable target are deleted.
//
// [Step scaling policies]: https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html
// [Predictive scaling policies]: https://docs.aws.amazon.com/autoscaling/application/userguide/aas-create-predictive-scaling-policy.html
// [Target tracking scaling policies]: https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html
func applicationautoscaling_PutScalingPolicy(cfg aws.Config, client *applicationautoscaling.Client) {
	input := &applicationautoscaling.PutScalingPolicyInput{
		// PolicyName: *string, // Required
		// ResourceId: *string, // Required
		// ScalableDimension: types.ScalableDimension, // Required
		// ServiceNamespace: types.ServiceNamespace, // Required
	}

	if len(_applicationautoscalingPolicyName) > 0 {
		input.PolicyName = aws.String(_applicationautoscalingPolicyName)
	}
	if len(_applicationautoscalingResourceId) > 0 {
		input.ResourceId = aws.String(_applicationautoscalingResourceId)
	}
	if len(_applicationautoscalingScalableDimension) > 0 {
		if err := assignInputField(input, "ScalableDimension", _applicationautoscalingScalableDimension); err != nil {
			log.Errorf("invalid --scalable-dimension: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingServiceNamespace) > 0 {
		if err := assignInputField(input, "ServiceNamespace", _applicationautoscalingServiceNamespace); err != nil {
			log.Errorf("invalid --service-namespace: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingPolicyType) > 0 {
		if err := assignInputField(input, "PolicyType", _applicationautoscalingPolicyType); err != nil {
			log.Errorf("invalid --policy-type: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingPredictiveScalingPolicyConfiguration) > 0 {
		if err := assignInputField(input, "PredictiveScalingPolicyConfiguration", _applicationautoscalingPredictiveScalingPolicyConfiguration); err != nil {
			log.Errorf("invalid --predictive-scaling-policy-configuration: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingStepScalingPolicyConfiguration) > 0 {
		if err := assignInputField(input, "StepScalingPolicyConfiguration", _applicationautoscalingStepScalingPolicyConfiguration); err != nil {
			log.Errorf("invalid --step-scaling-policy-configuration: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingTargetTrackingScalingPolicyConfiguration) > 0 {
		if err := assignInputField(input, "TargetTrackingScalingPolicyConfiguration", _applicationautoscalingTargetTrackingScalingPolicyConfiguration); err != nil {
			log.Errorf("invalid --target-tracking-scaling-policy-configuration: %s", err.Error())
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

// Creates or updates a scheduled action for an Application Auto Scaling scalable
// target.
//
// Each scalable target is identified by a service namespace, resource ID, and
// scalable dimension. A scheduled action applies to the scalable target identified
// by those three attributes. You cannot create a scheduled action until you have
// registered the resource as a scalable target.
//
// When you specify start and end times with a recurring schedule using a cron
// expression or rates, they form the boundaries for when the recurring action
// starts and stops.
//
// To update a scheduled action, specify the parameters that you want to change.
// If you don't specify start and end times, the old values are deleted.
//
// For more information, see [Scheduled scaling] in the Application Auto Scaling User Guide.
//
// If a scalable target is deregistered, the scalable target is no longer
// available to run scheduled actions. Any scheduled actions that were specified
// for the scalable target are deleted.
//
// [Scheduled scaling]: https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-scheduled-scaling.html
func applicationautoscaling_PutScheduledAction(cfg aws.Config, client *applicationautoscaling.Client) {
	input := &applicationautoscaling.PutScheduledActionInput{
		// ResourceId: *string, // Required
		// ScalableDimension: types.ScalableDimension, // Required
		// ScheduledActionName: *string, // Required
		// ServiceNamespace: types.ServiceNamespace, // Required
	}

	if len(_applicationautoscalingResourceId) > 0 {
		input.ResourceId = aws.String(_applicationautoscalingResourceId)
	}
	if len(_applicationautoscalingScalableDimension) > 0 {
		if err := assignInputField(input, "ScalableDimension", _applicationautoscalingScalableDimension); err != nil {
			log.Errorf("invalid --scalable-dimension: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingScheduledActionName) > 0 {
		input.ScheduledActionName = aws.String(_applicationautoscalingScheduledActionName)
	}
	if len(_applicationautoscalingServiceNamespace) > 0 {
		if err := assignInputField(input, "ServiceNamespace", _applicationautoscalingServiceNamespace); err != nil {
			log.Errorf("invalid --service-namespace: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _applicationautoscalingEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingScalableTargetAction) > 0 {
		if err := assignInputField(input, "ScalableTargetAction", _applicationautoscalingScalableTargetAction); err != nil {
			log.Errorf("invalid --scalable-target-action: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingSchedule) > 0 {
		input.Schedule = aws.String(_applicationautoscalingSchedule)
	}
	if len(_applicationautoscalingStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _applicationautoscalingStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingTimezone) > 0 {
		input.Timezone = aws.String(_applicationautoscalingTimezone)
	}

	if resp, err := client.PutScheduledAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers or updates a scalable target, which is the resource that you want to
// scale.
//
// Scalable targets are uniquely identified by the combination of resource ID,
// scalable dimension, and namespace, which represents some capacity dimension of
// the underlying service.
//
// When you register a new scalable target, you must specify values for the
// minimum and maximum capacity. If the specified resource is not active in the
// target service, this operation does not change the resource's current capacity.
// Otherwise, it changes the resource's current capacity to a value that is inside
// of this range.
//
// If you add a scaling policy, current capacity is adjustable within the
// specified range when scaling starts. Application Auto Scaling scaling policies
// will not scale capacity to values that are outside of the minimum and maximum
// range.
//
// After you register a scalable target, you do not need to register it again to
// use other Application Auto Scaling operations. To see which resources have been
// registered, use [DescribeScalableTargets]. You can also view the scaling policies for a service
// namespace by using [DescribeScalableTargets]. If you no longer need a scalable target, you can
// deregister it by using [DeregisterScalableTarget].
//
// To update a scalable target, specify the parameters that you want to change.
// Include the parameters that identify the scalable target: resource ID, scalable
// dimension, and namespace. Any parameters that you don't specify are not changed
// by this update request.
//
// If you call the RegisterScalableTarget API operation to create a scalable
// target, there might be a brief delay until the operation achieves [eventual consistency]. You might
// become aware of this brief delay if you get unexpected errors when performing
// sequential operations. The typical strategy is to retry the request, and some
// Amazon Web Services SDKs include automatic backoff and retry logic.
//
// If you call the RegisterScalableTarget API operation to update an existing
// scalable target, Application Auto Scaling retrieves the current capacity of the
// resource. If it's below the minimum capacity or above the maximum capacity,
// Application Auto Scaling adjusts the capacity of the scalable target to place it
// within these bounds, even if you don't include the MinCapacity or MaxCapacity
// request parameters.
//
// [DescribeScalableTargets]: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_DescribeScalableTargets.html
// [eventual consistency]: https://en.wikipedia.org/wiki/Eventual_consistency
// [DeregisterScalableTarget]: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_DeregisterScalableTarget.html
func applicationautoscaling_RegisterScalableTarget(cfg aws.Config, client *applicationautoscaling.Client) {
	input := &applicationautoscaling.RegisterScalableTargetInput{
		// ResourceId: *string, // Required
		// ScalableDimension: types.ScalableDimension, // Required
		// ServiceNamespace: types.ServiceNamespace, // Required
	}

	if len(_applicationautoscalingResourceId) > 0 {
		input.ResourceId = aws.String(_applicationautoscalingResourceId)
	}
	if len(_applicationautoscalingScalableDimension) > 0 {
		if err := assignInputField(input, "ScalableDimension", _applicationautoscalingScalableDimension); err != nil {
			log.Errorf("invalid --scalable-dimension: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingServiceNamespace) > 0 {
		if err := assignInputField(input, "ServiceNamespace", _applicationautoscalingServiceNamespace); err != nil {
			log.Errorf("invalid --service-namespace: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingMaxCapacity) > 0 {
		if err := assignInputField(input, "MaxCapacity", _applicationautoscalingMaxCapacity); err != nil {
			log.Errorf("invalid --max-capacity: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingMinCapacity) > 0 {
		if err := assignInputField(input, "MinCapacity", _applicationautoscalingMinCapacity); err != nil {
			log.Errorf("invalid --min-capacity: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingRoleARN) > 0 {
		input.RoleARN = aws.String(_applicationautoscalingRoleARN)
	}
	if len(_applicationautoscalingSuspendedState) > 0 {
		if err := assignInputField(input, "SuspendedState", _applicationautoscalingSuspendedState); err != nil {
			log.Errorf("invalid --suspended-state: %s", err.Error())
			return
		}
	}
	if len(_applicationautoscalingTags) > 0 {
		if err := assignInputField(input, "Tags", _applicationautoscalingTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterScalableTarget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or edits tags on an Application Auto Scaling scalable target.
// Each tag consists of a tag key and a tag value, which are both case-sensitive
// strings. To add a tag, specify a new tag key and a tag value. To edit a tag,
// specify an existing tag key and a new tag value.
//
// You can use this operation to tag an Application Auto Scaling scalable target,
// but you cannot tag a scaling policy or scheduled action.
//
// You can also add tags to an Application Auto Scaling scalable target while
// creating it ( RegisterScalableTarget ).
//
// For general information about tags, including the format and syntax, see [Tagging your Amazon Web Services resources] in
// the Amazon Web Services General Reference.
//
// Use tags to control access to a scalable target. For more information, see [Tagging support for Application Auto Scaling] in
// the Application Auto Scaling User Guide.
//
// [Tagging your Amazon Web Services resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
// [Tagging support for Application Auto Scaling]: https://docs.aws.amazon.com/autoscaling/application/userguide/resource-tagging-support.html
func applicationautoscaling_TagResource(cfg aws.Config, client *applicationautoscaling.Client) {
	input := &applicationautoscaling.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_applicationautoscalingResourceARN) > 0 {
		input.ResourceARN = aws.String(_applicationautoscalingResourceARN)
	}
	if len(_applicationautoscalingTags) > 0 {
		if err := assignInputField(input, "Tags", _applicationautoscalingTags); err != nil {
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

// Deletes tags from an Application Auto Scaling scalable target. To delete a tag,
// specify the tag key and the Application Auto Scaling scalable target.
func applicationautoscaling_UntagResource(cfg aws.Config, client *applicationautoscaling.Client) {
	input := &applicationautoscaling.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_applicationautoscalingResourceARN) > 0 {
		input.ResourceARN = aws.String(_applicationautoscalingResourceARN)
	}
	if len(_applicationautoscalingTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _applicationautoscalingTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_applicationautoscalingCmd)
	_applicationautoscalingCmd.Flags().SortFlags = false

	_applicationautoscalingCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_applicationautoscalingCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_applicationautoscalingCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_applicationautoscalingCmd.Flags().StringVarP(&_applicationautoscalingEndTime, "end-time", "", "", "End Time")
	_applicationautoscalingCmd.Flags().StringVarP(&_applicationautoscalingIncludeNotScaledActivities, "include-not-scaled-activities", "", "", "Include Not Scaled Activities")
	_applicationautoscalingCmd.Flags().StringVarP(&_applicationautoscalingMaxCapacity, "max-capacity", "", "", "Max Capacity")
	_applicationautoscalingCmd.Flags().StringVarP(&_applicationautoscalingMaxResults, "max-results", "", "", "Max Results")
	_applicationautoscalingCmd.Flags().StringVarP(&_applicationautoscalingMinCapacity, "min-capacity", "", "", "Min Capacity")
	_applicationautoscalingCmd.Flags().StringVarP(&_applicationautoscalingNextToken, "next-token", "", "", "Next Token")
	_applicationautoscalingCmd.Flags().StringVarP(&_applicationautoscalingPolicyName, "policy-name", "", "", "Policy Name")
	_applicationautoscalingCmd.Flags().StringSliceVarP(&_applicationautoscalingPolicyNames, "policy-names", "", nil, "Policy Names")
	_applicationautoscalingCmd.Flags().StringVarP(&_applicationautoscalingPolicyType, "policy-type", "", "", "Policy Type")
	_applicationautoscalingCmd.Flags().StringVarP(&_applicationautoscalingPredictiveScalingPolicyConfiguration, "predictive-scaling-policy-configuration", "", "", "Predictive Scaling Policy Configuration")
	_applicationautoscalingCmd.Flags().StringVarP(&_applicationautoscalingResourceARN, "resource-arn", "", "", "Resource ARN")
	_applicationautoscalingCmd.Flags().StringVarP(&_applicationautoscalingResourceId, "resource-id", "", "", "Resource ID")
	_applicationautoscalingCmd.Flags().StringSliceVarP(&_applicationautoscalingResourceIds, "resource-ids", "", nil, "Resource Ids")
	_applicationautoscalingCmd.Flags().StringVarP(&_applicationautoscalingRoleARN, "role-arn", "", "", "Role ARN")
	_applicationautoscalingCmd.Flags().StringVarP(&_applicationautoscalingScalableDimension, "scalable-dimension", "", "", "Scalable Dimension")
	_applicationautoscalingCmd.Flags().StringVarP(&_applicationautoscalingScalableTargetAction, "scalable-target-action", "", "", "Scalable Target Action")
	_applicationautoscalingCmd.Flags().StringVarP(&_applicationautoscalingSchedule, "schedule", "", "", "Schedule")
	_applicationautoscalingCmd.Flags().StringVarP(&_applicationautoscalingScheduledActionName, "scheduled-action-name", "", "", "Scheduled Action Name")
	_applicationautoscalingCmd.Flags().StringSliceVarP(&_applicationautoscalingScheduledActionNames, "scheduled-action-names", "", nil, "Scheduled Action Names")
	_applicationautoscalingCmd.Flags().StringVarP(&_applicationautoscalingServiceNamespace, "service-namespace", "", "", "Service Namespace")
	_applicationautoscalingCmd.Flags().StringVarP(&_applicationautoscalingStartTime, "start-time", "", "", "Start Time")
	_applicationautoscalingCmd.Flags().StringVarP(&_applicationautoscalingStepScalingPolicyConfiguration, "step-scaling-policy-configuration", "", "", "Step Scaling Policy Configuration")
	_applicationautoscalingCmd.Flags().StringVarP(&_applicationautoscalingSuspendedState, "suspended-state", "", "", "Suspended State")
	_applicationautoscalingCmd.Flags().StringSliceVarP(&_applicationautoscalingTagKeys, "tag-keys", "", nil, "Tag Keys")
	_applicationautoscalingCmd.Flags().StringVarP(&_applicationautoscalingTags, "tags", "", "", "Tags")
	_applicationautoscalingCmd.Flags().StringVarP(&_applicationautoscalingTargetTrackingScalingPolicyConfiguration, "target-tracking-scaling-policy-configuration", "", "", "Target Tracking Scaling Policy Configuration")
	_applicationautoscalingCmd.Flags().StringVarP(&_applicationautoscalingTimezone, "timezone", "", "", "Timezone")

	_applicationautoscalingCmd.Flags().BoolVarP(&_applicationautoscalingDeleteScalingPolicy, "delete-scaling-policy", "", false, "Delete Scaling Policy")
	_applicationautoscalingCmd.Flags().BoolVarP(&_applicationautoscalingDeleteScheduledAction, "delete-scheduled-action", "", false, "Delete Scheduled Action")
	_applicationautoscalingCmd.Flags().BoolVarP(&_applicationautoscalingDeregisterScalableTarget, "deregister-scalable-target", "", false, "Deregister Scalable Target")
	_applicationautoscalingCmd.Flags().BoolVarP(&_applicationautoscalingDescribeScalableTargets, "describe-scalable-targets", "", false, "Describe Scalable Targets")
	_applicationautoscalingCmd.Flags().BoolVarP(&_applicationautoscalingDescribeScalingActivities, "describe-scaling-activities", "", false, "Describe Scaling Activities")
	_applicationautoscalingCmd.Flags().BoolVarP(&_applicationautoscalingDescribeScalingPolicies, "describe-scaling-policies", "", false, "Describe Scaling Policies")
	_applicationautoscalingCmd.Flags().BoolVarP(&_applicationautoscalingDescribeScheduledActions, "describe-scheduled-actions", "", false, "Describe Scheduled Actions")
	_applicationautoscalingCmd.Flags().BoolVarP(&_applicationautoscalingGetPredictiveScalingForecast, "get-predictive-scaling-forecast", "", false, "Get Predictive Scaling Forecast")
	_applicationautoscalingCmd.Flags().BoolVarP(&_applicationautoscalingListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_applicationautoscalingCmd.Flags().BoolVarP(&_applicationautoscalingPutScalingPolicy, "put-scaling-policy", "", false, "Put Scaling Policy")
	_applicationautoscalingCmd.Flags().BoolVarP(&_applicationautoscalingPutScheduledAction, "put-scheduled-action", "", false, "Put Scheduled Action")
	_applicationautoscalingCmd.Flags().BoolVarP(&_applicationautoscalingRegisterScalableTarget, "register-scalable-target", "", false, "Register Scalable Target")
	_applicationautoscalingCmd.Flags().BoolVarP(&_applicationautoscalingTagResource, "tag-resource", "", false, "Tag Resource")
	_applicationautoscalingCmd.Flags().BoolVarP(&_applicationautoscalingUntagResource, "untag-resource", "", false, "Untag Resource")

}
