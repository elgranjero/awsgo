package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// codedeployCmd represents the codedeploy command
var _codedeployCmd = &cobra.Command{
	Use:   "codedeploy",
	Short: "AWS codedeploy CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := codedeploy.NewFromConfig(cfg)
		if _codedeployAddTagsToOnPremisesInstances {
			codedeploy_AddTagsToOnPremisesInstances(cfg, client)
			return
		}
		if _codedeployBatchGetApplicationRevisions {
			codedeploy_BatchGetApplicationRevisions(cfg, client)
			return
		}
		if _codedeployBatchGetApplications {
			codedeploy_BatchGetApplications(cfg, client)
			return
		}
		if _codedeployBatchGetDeploymentGroups {
			codedeploy_BatchGetDeploymentGroups(cfg, client)
			return
		}
		if _codedeployBatchGetDeploymentInstances {
			codedeploy_BatchGetDeploymentInstances(cfg, client)
			return
		}
		if _codedeployBatchGetDeploymentTargets {
			codedeploy_BatchGetDeploymentTargets(cfg, client)
			return
		}
		if _codedeployBatchGetDeployments {
			codedeploy_BatchGetDeployments(cfg, client)
			return
		}
		if _codedeployBatchGetOnPremisesInstances {
			codedeploy_BatchGetOnPremisesInstances(cfg, client)
			return
		}
		if _codedeployContinueDeployment {
			codedeploy_ContinueDeployment(cfg, client)
			return
		}
		if _codedeployCreateApplication {
			codedeploy_CreateApplication(cfg, client)
			return
		}
		if _codedeployCreateDeployment {
			codedeploy_CreateDeployment(cfg, client)
			return
		}
		if _codedeployCreateDeploymentConfig {
			codedeploy_CreateDeploymentConfig(cfg, client)
			return
		}
		if _codedeployCreateDeploymentGroup {
			codedeploy_CreateDeploymentGroup(cfg, client)
			return
		}
		if _codedeployDeleteApplication {
			codedeploy_DeleteApplication(cfg, client)
			return
		}
		if _codedeployDeleteDeploymentConfig {
			codedeploy_DeleteDeploymentConfig(cfg, client)
			return
		}
		if _codedeployDeleteDeploymentGroup {
			codedeploy_DeleteDeploymentGroup(cfg, client)
			return
		}
		if _codedeployDeleteGitHubAccountToken {
			codedeploy_DeleteGitHubAccountToken(cfg, client)
			return
		}
		if _codedeployDeleteResourcesByExternalId {
			codedeploy_DeleteResourcesByExternalId(cfg, client)
			return
		}
		if _codedeployDeregisterOnPremisesInstance {
			codedeploy_DeregisterOnPremisesInstance(cfg, client)
			return
		}
		if _codedeployGetApplication {
			codedeploy_GetApplication(cfg, client)
			return
		}
		if _codedeployGetApplicationRevision {
			codedeploy_GetApplicationRevision(cfg, client)
			return
		}
		if _codedeployGetDeployment {
			codedeploy_GetDeployment(cfg, client)
			return
		}
		if _codedeployGetDeploymentConfig {
			codedeploy_GetDeploymentConfig(cfg, client)
			return
		}
		if _codedeployGetDeploymentGroup {
			codedeploy_GetDeploymentGroup(cfg, client)
			return
		}
		if _codedeployGetDeploymentInstance {
			codedeploy_GetDeploymentInstance(cfg, client)
			return
		}
		if _codedeployGetDeploymentTarget {
			codedeploy_GetDeploymentTarget(cfg, client)
			return
		}
		if _codedeployGetOnPremisesInstance {
			codedeploy_GetOnPremisesInstance(cfg, client)
			return
		}
		if _codedeployListApplicationRevisions {
			codedeploy_ListApplicationRevisions(cfg, client)
			return
		}
		if _codedeployListApplications {
			codedeploy_ListApplications(cfg, client)
			return
		}
		if _codedeployListDeploymentConfigs {
			codedeploy_ListDeploymentConfigs(cfg, client)
			return
		}
		if _codedeployListDeploymentGroups {
			codedeploy_ListDeploymentGroups(cfg, client)
			return
		}
		if _codedeployListDeploymentInstances {
			codedeploy_ListDeploymentInstances(cfg, client)
			return
		}
		if _codedeployListDeploymentTargets {
			codedeploy_ListDeploymentTargets(cfg, client)
			return
		}
		if _codedeployListDeployments {
			codedeploy_ListDeployments(cfg, client)
			return
		}
		if _codedeployListGitHubAccountTokenNames {
			codedeploy_ListGitHubAccountTokenNames(cfg, client)
			return
		}
		if _codedeployListOnPremisesInstances {
			codedeploy_ListOnPremisesInstances(cfg, client)
			return
		}
		if _codedeployListTagsForResource {
			codedeploy_ListTagsForResource(cfg, client)
			return
		}
		if _codedeployPutLifecycleEventHookExecutionStatus {
			codedeploy_PutLifecycleEventHookExecutionStatus(cfg, client)
			return
		}
		if _codedeployRegisterApplicationRevision {
			codedeploy_RegisterApplicationRevision(cfg, client)
			return
		}
		if _codedeployRegisterOnPremisesInstance {
			codedeploy_RegisterOnPremisesInstance(cfg, client)
			return
		}
		if _codedeployRemoveTagsFromOnPremisesInstances {
			codedeploy_RemoveTagsFromOnPremisesInstances(cfg, client)
			return
		}
		if _codedeploySkipWaitTimeForInstanceTermination {
			codedeploy_SkipWaitTimeForInstanceTermination(cfg, client)
			return
		}
		if _codedeployStopDeployment {
			codedeploy_StopDeployment(cfg, client)
			return
		}
		if _codedeployTagResource {
			codedeploy_TagResource(cfg, client)
			return
		}
		if _codedeployUntagResource {
			codedeploy_UntagResource(cfg, client)
			return
		}
		if _codedeployUpdateApplication {
			codedeploy_UpdateApplication(cfg, client)
			return
		}
		if _codedeployUpdateDeploymentGroup {
			codedeploy_UpdateDeploymentGroup(cfg, client)
			return
		}

	},
}

var (
	_codedeployAddTagsToOnPremisesInstances         bool
	_codedeployBatchGetApplicationRevisions         bool
	_codedeployBatchGetApplications                 bool
	_codedeployBatchGetDeploymentGroups             bool
	_codedeployBatchGetDeploymentInstances          bool
	_codedeployBatchGetDeploymentTargets            bool
	_codedeployBatchGetDeployments                  bool
	_codedeployBatchGetOnPremisesInstances          bool
	_codedeployContinueDeployment                   bool
	_codedeployCreateApplication                    bool
	_codedeployCreateDeployment                     bool
	_codedeployCreateDeploymentConfig               bool
	_codedeployCreateDeploymentGroup                bool
	_codedeployDeleteApplication                    bool
	_codedeployDeleteDeploymentConfig               bool
	_codedeployDeleteDeploymentGroup                bool
	_codedeployDeleteGitHubAccountToken             bool
	_codedeployDeleteResourcesByExternalId          bool
	_codedeployDeregisterOnPremisesInstance         bool
	_codedeployGetApplication                       bool
	_codedeployGetApplicationRevision               bool
	_codedeployGetDeployment                        bool
	_codedeployGetDeploymentConfig                  bool
	_codedeployGetDeploymentGroup                   bool
	_codedeployGetDeploymentInstance                bool
	_codedeployGetDeploymentTarget                  bool
	_codedeployGetOnPremisesInstance                bool
	_codedeployListApplicationRevisions             bool
	_codedeployListApplications                     bool
	_codedeployListDeploymentConfigs                bool
	_codedeployListDeploymentGroups                 bool
	_codedeployListDeploymentInstances              bool
	_codedeployListDeploymentTargets                bool
	_codedeployListDeployments                      bool
	_codedeployListGitHubAccountTokenNames          bool
	_codedeployListOnPremisesInstances              bool
	_codedeployListTagsForResource                  bool
	_codedeployPutLifecycleEventHookExecutionStatus bool
	_codedeployRegisterApplicationRevision          bool
	_codedeployRegisterOnPremisesInstance           bool
	_codedeployRemoveTagsFromOnPremisesInstances    bool
	_codedeploySkipWaitTimeForInstanceTermination   bool
	_codedeployStopDeployment                       bool
	_codedeployTagResource                          bool
	_codedeployUntagResource                        bool
	_codedeployUpdateApplication                    bool
	_codedeployUpdateDeploymentGroup                bool

	_codedeployAlarmConfiguration               string
	_codedeployApplicationName                  string
	_codedeployApplicationNames                 []string
	_codedeployAutoRollbackConfiguration        string
	_codedeployAutoRollbackEnabled              string
	_codedeployAutoScalingGroups                []string
	_codedeployBlueGreenDeploymentConfiguration string
	_codedeployComputePlatform                  string
	_codedeployCreateTimeRange                  string
	_codedeployCurrentDeploymentGroupName       string
	_codedeployDeployed                         string
	_codedeployDeploymentConfigName             string
	_codedeployDeploymentGroupName              string
	_codedeployDeploymentGroupNames             []string
	_codedeployDeploymentId                     string
	_codedeployDeploymentIds                    []string
	_codedeployDeploymentStyle                  string
	_codedeployDeploymentWaitType               string
	_codedeployDescription                      string
	_codedeployEc2TagFilters                    string
	_codedeployEc2TagSet                        string
	_codedeployEcsServices                      string
	_codedeployExternalId                       string
	_codedeployFileExistsBehavior               string
	_codedeployIamSessionArn                    string
	_codedeployIamUserArn                       string
	_codedeployIgnoreApplicationStopFailures    string
	_codedeployIncludeOnlyStatuses              string
	_codedeployInstanceId                       string
	_codedeployInstanceIds                      []string
	_codedeployInstanceName                     string
	_codedeployInstanceNames                    []string
	_codedeployInstanceStatusFilter             string
	_codedeployInstanceTypeFilter               string
	_codedeployLifecycleEventHookExecutionId    string
	_codedeployLoadBalancerInfo                 string
	_codedeployMinimumHealthyHosts              string
	_codedeployNewApplicationName               string
	_codedeployNewDeploymentGroupName           string
	_codedeployNextToken                        string
	_codedeployOnPremisesInstanceTagFilters     string
	_codedeployOnPremisesTagSet                 string
	_codedeployOutdatedInstancesStrategy        string
	_codedeployOverrideAlarmConfiguration       string
	_codedeployRegistrationStatus               string
	_codedeployResourceArn                      string
	_codedeployRevision                         string
	_codedeployRevisions                        string
	_codedeployS3Bucket                         string
	_codedeployS3KeyPrefix                      string
	_codedeployServiceRoleArn                   string
	_codedeploySortBy                           string
	_codedeploySortOrder                        string
	_codedeployStatus                           string
	_codedeployTagFilters                       string
	_codedeployTagKeys                          []string
	_codedeployTags                             string
	_codedeployTargetFilters                    string
	_codedeployTargetId                         string
	_codedeployTargetIds                        []string
	_codedeployTargetInstances                  string
	_codedeployTerminationHookEnabled           string
	_codedeployTokenName                        string
	_codedeployTrafficRoutingConfig             string
	_codedeployTriggerConfigurations            string
	_codedeployUpdateOutdatedInstancesOnly      string
	_codedeployZonalConfig                      string
)

// Adds tags to on-premises instances.
func codedeploy_AddTagsToOnPremisesInstances(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.AddTagsToOnPremisesInstancesInput{
		// InstanceNames: []string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_codedeployInstanceNames) > 0 {
		input.InstanceNames = append([]string(nil), _codedeployInstanceNames...)
	}
	if len(_codedeployTags) > 0 {
		if err := assignInputField(input, "Tags", _codedeployTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddTagsToOnPremisesInstances(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about one or more application revisions. The maximum number of
// application revisions that can be returned is 25.
func codedeploy_BatchGetApplicationRevisions(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.BatchGetApplicationRevisionsInput{
		// ApplicationName: *string, // Required
		// Revisions: []types.RevisionLocation, // Required
	}

	if len(_codedeployApplicationName) > 0 {
		input.ApplicationName = aws.String(_codedeployApplicationName)
	}
	if len(_codedeployRevisions) > 0 {
		if err := assignInputField(input, "Revisions", _codedeployRevisions); err != nil {
			log.Errorf("invalid --revisions: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchGetApplicationRevisions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about one or more applications. The maximum number of
// applications that can be returned is 100.
func codedeploy_BatchGetApplications(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.BatchGetApplicationsInput{
		// ApplicationNames: []string, // Required
	}

	if len(_codedeployApplicationNames) > 0 {
		input.ApplicationNames = append([]string(nil), _codedeployApplicationNames...)
	}

	if resp, err := client.BatchGetApplications(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about one or more deployment groups.
func codedeploy_BatchGetDeploymentGroups(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.BatchGetDeploymentGroupsInput{
		// ApplicationName: *string, // Required
		// DeploymentGroupNames: []string, // Required
	}

	if len(_codedeployApplicationName) > 0 {
		input.ApplicationName = aws.String(_codedeployApplicationName)
	}
	if len(_codedeployDeploymentGroupNames) > 0 {
		input.DeploymentGroupNames = append([]string(nil), _codedeployDeploymentGroupNames...)
	}

	if resp, err := client.BatchGetDeploymentGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This method works, but is deprecated. Use BatchGetDeploymentTargets instead.
// Returns an array of one or more instances associated with a deployment. This
// method works with EC2/On-premises and Lambda compute platforms. The newer
// BatchGetDeploymentTargets works with all compute platforms. The maximum number
// of instances that can be returned is 25.
//
// Deprecated: This operation is deprecated, use BatchGetDeploymentTargets instead.
func codedeploy_BatchGetDeploymentInstances(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.BatchGetDeploymentInstancesInput{
		// DeploymentId: *string, // Required
		// InstanceIds: []string, // Required
	}

	if len(_codedeployDeploymentId) > 0 {
		input.DeploymentId = aws.String(_codedeployDeploymentId)
	}
	if len(_codedeployInstanceIds) > 0 {
		input.InstanceIds = append([]string(nil), _codedeployInstanceIds...)
	}

	if resp, err := client.BatchGetDeploymentInstances(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an array of one or more targets associated with a deployment. This
// method works with all compute types and should be used instead of the deprecated
// BatchGetDeploymentInstances . The maximum number of targets that can be returned
// is 25.
//
// The type of targets returned depends on the deployment's compute platform or
// deployment method:
//
// - EC2/On-premises: Information about Amazon EC2 instance targets.
//
// - Lambda: Information about Lambda functions targets.
//
// - Amazon ECS: Information about Amazon ECS service targets.
//
// - CloudFormation: Information about targets of blue/green deployments
// initiated by a CloudFormation stack update.
func codedeploy_BatchGetDeploymentTargets(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.BatchGetDeploymentTargetsInput{
		// DeploymentId: *string, // Required
		// TargetIds: []string, // Required
	}

	if len(_codedeployDeploymentId) > 0 {
		input.DeploymentId = aws.String(_codedeployDeploymentId)
	}
	if len(_codedeployTargetIds) > 0 {
		input.TargetIds = append([]string(nil), _codedeployTargetIds...)
	}

	if resp, err := client.BatchGetDeploymentTargets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about one or more deployments. The maximum number of
// deployments that can be returned is 25.
func codedeploy_BatchGetDeployments(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.BatchGetDeploymentsInput{
		// DeploymentIds: []string, // Required
	}

	if len(_codedeployDeploymentIds) > 0 {
		input.DeploymentIds = append([]string(nil), _codedeployDeploymentIds...)
	}

	if resp, err := client.BatchGetDeployments(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about one or more on-premises instances. The maximum number of
// on-premises instances that can be returned is 25.
func codedeploy_BatchGetOnPremisesInstances(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.BatchGetOnPremisesInstancesInput{
		// InstanceNames: []string, // Required
	}

	if len(_codedeployInstanceNames) > 0 {
		input.InstanceNames = append([]string(nil), _codedeployInstanceNames...)
	}

	if resp, err := client.BatchGetOnPremisesInstances(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For a blue/green deployment, starts the process of rerouting traffic from
// instances in the original environment to instances in the replacement
// environment without waiting for a specified wait time to elapse. (Traffic
// rerouting, which is achieved by registering instances in the replacement
// environment with the load balancer, can start as soon as all instances have a
// status of Ready.)
func codedeploy_ContinueDeployment(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.ContinueDeploymentInput{}

	if len(_codedeployDeploymentId) > 0 {
		input.DeploymentId = aws.String(_codedeployDeploymentId)
	}
	if len(_codedeployDeploymentWaitType) > 0 {
		if err := assignInputField(input, "DeploymentWaitType", _codedeployDeploymentWaitType); err != nil {
			log.Errorf("invalid --deployment-wait-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ContinueDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an application.
func codedeploy_CreateApplication(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.CreateApplicationInput{
		// ApplicationName: *string, // Required
	}

	if len(_codedeployApplicationName) > 0 {
		input.ApplicationName = aws.String(_codedeployApplicationName)
	}
	if len(_codedeployComputePlatform) > 0 {
		if err := assignInputField(input, "ComputePlatform", _codedeployComputePlatform); err != nil {
			log.Errorf("invalid --compute-platform: %s", err.Error())
			return
		}
	}
	if len(_codedeployTags) > 0 {
		if err := assignInputField(input, "Tags", _codedeployTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deploys an application revision through the specified deployment group.
func codedeploy_CreateDeployment(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.CreateDeploymentInput{
		// ApplicationName: *string, // Required
	}

	if len(_codedeployApplicationName) > 0 {
		input.ApplicationName = aws.String(_codedeployApplicationName)
	}
	if len(_codedeployAutoRollbackConfiguration) > 0 {
		if err := assignInputField(input, "AutoRollbackConfiguration", _codedeployAutoRollbackConfiguration); err != nil {
			log.Errorf("invalid --auto-rollback-configuration: %s", err.Error())
			return
		}
	}
	if len(_codedeployDeploymentConfigName) > 0 {
		input.DeploymentConfigName = aws.String(_codedeployDeploymentConfigName)
	}
	if len(_codedeployDeploymentGroupName) > 0 {
		input.DeploymentGroupName = aws.String(_codedeployDeploymentGroupName)
	}
	if len(_codedeployDescription) > 0 {
		input.Description = aws.String(_codedeployDescription)
	}
	if len(_codedeployFileExistsBehavior) > 0 {
		if err := assignInputField(input, "FileExistsBehavior", _codedeployFileExistsBehavior); err != nil {
			log.Errorf("invalid --file-exists-behavior: %s", err.Error())
			return
		}
	}
	if len(_codedeployIgnoreApplicationStopFailures) > 0 {
		if err := assignInputField(input, "IgnoreApplicationStopFailures", _codedeployIgnoreApplicationStopFailures); err != nil {
			log.Errorf("invalid --ignore-application-stop-failures: %s", err.Error())
			return
		}
	}
	if len(_codedeployOverrideAlarmConfiguration) > 0 {
		if err := assignInputField(input, "OverrideAlarmConfiguration", _codedeployOverrideAlarmConfiguration); err != nil {
			log.Errorf("invalid --override-alarm-configuration: %s", err.Error())
			return
		}
	}
	if len(_codedeployRevision) > 0 {
		if err := assignInputField(input, "Revision", _codedeployRevision); err != nil {
			log.Errorf("invalid --revision: %s", err.Error())
			return
		}
	}
	if len(_codedeployTargetInstances) > 0 {
		if err := assignInputField(input, "TargetInstances", _codedeployTargetInstances); err != nil {
			log.Errorf("invalid --target-instances: %s", err.Error())
			return
		}
	}
	if len(_codedeployUpdateOutdatedInstancesOnly) > 0 {
		if err := assignInputField(input, "UpdateOutdatedInstancesOnly", _codedeployUpdateOutdatedInstancesOnly); err != nil {
			log.Errorf("invalid --update-outdated-instances-only: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a deployment configuration.
func codedeploy_CreateDeploymentConfig(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.CreateDeploymentConfigInput{
		// DeploymentConfigName: *string, // Required
	}

	if len(_codedeployDeploymentConfigName) > 0 {
		input.DeploymentConfigName = aws.String(_codedeployDeploymentConfigName)
	}
	if len(_codedeployComputePlatform) > 0 {
		if err := assignInputField(input, "ComputePlatform", _codedeployComputePlatform); err != nil {
			log.Errorf("invalid --compute-platform: %s", err.Error())
			return
		}
	}
	if len(_codedeployMinimumHealthyHosts) > 0 {
		if err := assignInputField(input, "MinimumHealthyHosts", _codedeployMinimumHealthyHosts); err != nil {
			log.Errorf("invalid --minimum-healthy-hosts: %s", err.Error())
			return
		}
	}
	if len(_codedeployTrafficRoutingConfig) > 0 {
		if err := assignInputField(input, "TrafficRoutingConfig", _codedeployTrafficRoutingConfig); err != nil {
			log.Errorf("invalid --traffic-routing-config: %s", err.Error())
			return
		}
	}
	if len(_codedeployZonalConfig) > 0 {
		if err := assignInputField(input, "ZonalConfig", _codedeployZonalConfig); err != nil {
			log.Errorf("invalid --zonal-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDeploymentConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a deployment group to which application revisions are deployed.
func codedeploy_CreateDeploymentGroup(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.CreateDeploymentGroupInput{
		// ApplicationName: *string, // Required
		// DeploymentGroupName: *string, // Required
		// ServiceRoleArn: *string, // Required
	}

	if len(_codedeployApplicationName) > 0 {
		input.ApplicationName = aws.String(_codedeployApplicationName)
	}
	if len(_codedeployDeploymentGroupName) > 0 {
		input.DeploymentGroupName = aws.String(_codedeployDeploymentGroupName)
	}
	if len(_codedeployServiceRoleArn) > 0 {
		input.ServiceRoleArn = aws.String(_codedeployServiceRoleArn)
	}
	if len(_codedeployAlarmConfiguration) > 0 {
		if err := assignInputField(input, "AlarmConfiguration", _codedeployAlarmConfiguration); err != nil {
			log.Errorf("invalid --alarm-configuration: %s", err.Error())
			return
		}
	}
	if len(_codedeployAutoRollbackConfiguration) > 0 {
		if err := assignInputField(input, "AutoRollbackConfiguration", _codedeployAutoRollbackConfiguration); err != nil {
			log.Errorf("invalid --auto-rollback-configuration: %s", err.Error())
			return
		}
	}
	if len(_codedeployAutoScalingGroups) > 0 {
		input.AutoScalingGroups = append([]string(nil), _codedeployAutoScalingGroups...)
	}
	if len(_codedeployBlueGreenDeploymentConfiguration) > 0 {
		if err := assignInputField(input, "BlueGreenDeploymentConfiguration", _codedeployBlueGreenDeploymentConfiguration); err != nil {
			log.Errorf("invalid --blue-green-deployment-configuration: %s", err.Error())
			return
		}
	}
	if len(_codedeployDeploymentConfigName) > 0 {
		input.DeploymentConfigName = aws.String(_codedeployDeploymentConfigName)
	}
	if len(_codedeployDeploymentStyle) > 0 {
		if err := assignInputField(input, "DeploymentStyle", _codedeployDeploymentStyle); err != nil {
			log.Errorf("invalid --deployment-style: %s", err.Error())
			return
		}
	}
	if len(_codedeployEc2TagFilters) > 0 {
		if err := assignInputField(input, "Ec2TagFilters", _codedeployEc2TagFilters); err != nil {
			log.Errorf("invalid --ec2-tag-filters: %s", err.Error())
			return
		}
	}
	if len(_codedeployEc2TagSet) > 0 {
		if err := assignInputField(input, "Ec2TagSet", _codedeployEc2TagSet); err != nil {
			log.Errorf("invalid --ec2-tag-set: %s", err.Error())
			return
		}
	}
	if len(_codedeployEcsServices) > 0 {
		if err := assignInputField(input, "EcsServices", _codedeployEcsServices); err != nil {
			log.Errorf("invalid --ecs-services: %s", err.Error())
			return
		}
	}
	if len(_codedeployLoadBalancerInfo) > 0 {
		if err := assignInputField(input, "LoadBalancerInfo", _codedeployLoadBalancerInfo); err != nil {
			log.Errorf("invalid --load-balancer-info: %s", err.Error())
			return
		}
	}
	if len(_codedeployOnPremisesInstanceTagFilters) > 0 {
		if err := assignInputField(input, "OnPremisesInstanceTagFilters", _codedeployOnPremisesInstanceTagFilters); err != nil {
			log.Errorf("invalid --on-premises-instance-tag-filters: %s", err.Error())
			return
		}
	}
	if len(_codedeployOnPremisesTagSet) > 0 {
		if err := assignInputField(input, "OnPremisesTagSet", _codedeployOnPremisesTagSet); err != nil {
			log.Errorf("invalid --on-premises-tag-set: %s", err.Error())
			return
		}
	}
	if len(_codedeployOutdatedInstancesStrategy) > 0 {
		if err := assignInputField(input, "OutdatedInstancesStrategy", _codedeployOutdatedInstancesStrategy); err != nil {
			log.Errorf("invalid --outdated-instances-strategy: %s", err.Error())
			return
		}
	}
	if len(_codedeployTags) > 0 {
		if err := assignInputField(input, "Tags", _codedeployTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_codedeployTerminationHookEnabled) > 0 {
		if err := assignInputField(input, "TerminationHookEnabled", _codedeployTerminationHookEnabled); err != nil {
			log.Errorf("invalid --termination-hook-enabled: %s", err.Error())
			return
		}
	}
	if len(_codedeployTriggerConfigurations) > 0 {
		if err := assignInputField(input, "TriggerConfigurations", _codedeployTriggerConfigurations); err != nil {
			log.Errorf("invalid --trigger-configurations: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDeploymentGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an application.
func codedeploy_DeleteApplication(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.DeleteApplicationInput{
		// ApplicationName: *string, // Required
	}

	if len(_codedeployApplicationName) > 0 {
		input.ApplicationName = aws.String(_codedeployApplicationName)
	}

	if resp, err := client.DeleteApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a deployment configuration.
// A deployment configuration cannot be deleted if it is currently in use.
// Predefined configurations cannot be deleted.
func codedeploy_DeleteDeploymentConfig(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.DeleteDeploymentConfigInput{
		// DeploymentConfigName: *string, // Required
	}

	if len(_codedeployDeploymentConfigName) > 0 {
		input.DeploymentConfigName = aws.String(_codedeployDeploymentConfigName)
	}

	if resp, err := client.DeleteDeploymentConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a deployment group.
func codedeploy_DeleteDeploymentGroup(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.DeleteDeploymentGroupInput{
		// ApplicationName: *string, // Required
		// DeploymentGroupName: *string, // Required
	}

	if len(_codedeployApplicationName) > 0 {
		input.ApplicationName = aws.String(_codedeployApplicationName)
	}
	if len(_codedeployDeploymentGroupName) > 0 {
		input.DeploymentGroupName = aws.String(_codedeployDeploymentGroupName)
	}

	if resp, err := client.DeleteDeploymentGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a GitHub account connection.
func codedeploy_DeleteGitHubAccountToken(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.DeleteGitHubAccountTokenInput{}

	if len(_codedeployTokenName) > 0 {
		input.TokenName = aws.String(_codedeployTokenName)
	}

	if resp, err := client.DeleteGitHubAccountToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes resources linked to an external ID. This action only applies if you
// have configured blue/green deployments through CloudFormation.
//
// It is not necessary to call this action directly. CloudFormation calls it on
// your behalf when it needs to delete stack resources. This action is offered
// publicly in case you need to delete resources to comply with General Data
// Protection Regulation (GDPR) requirements.
func codedeploy_DeleteResourcesByExternalId(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.DeleteResourcesByExternalIdInput{}

	if len(_codedeployExternalId) > 0 {
		input.ExternalId = aws.String(_codedeployExternalId)
	}

	if resp, err := client.DeleteResourcesByExternalId(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregisters an on-premises instance.
func codedeploy_DeregisterOnPremisesInstance(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.DeregisterOnPremisesInstanceInput{
		// InstanceName: *string, // Required
	}

	if len(_codedeployInstanceName) > 0 {
		input.InstanceName = aws.String(_codedeployInstanceName)
	}

	if resp, err := client.DeregisterOnPremisesInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an application.
func codedeploy_GetApplication(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.GetApplicationInput{
		// ApplicationName: *string, // Required
	}

	if len(_codedeployApplicationName) > 0 {
		input.ApplicationName = aws.String(_codedeployApplicationName)
	}

	if resp, err := client.GetApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an application revision.
func codedeploy_GetApplicationRevision(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.GetApplicationRevisionInput{
		// ApplicationName: *string, // Required
		// Revision: *types.RevisionLocation, // Required
	}

	if len(_codedeployApplicationName) > 0 {
		input.ApplicationName = aws.String(_codedeployApplicationName)
	}
	if len(_codedeployRevision) > 0 {
		if err := assignInputField(input, "Revision", _codedeployRevision); err != nil {
			log.Errorf("invalid --revision: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetApplicationRevision(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a deployment.
// The content property of the appSpecContent object in the returned revision is
// always null. Use GetApplicationRevision and the sha256 property of the returned
// appSpecContent object to get the content of the deployment’s AppSpec file.
func codedeploy_GetDeployment(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.GetDeploymentInput{
		// DeploymentId: *string, // Required
	}

	if len(_codedeployDeploymentId) > 0 {
		input.DeploymentId = aws.String(_codedeployDeploymentId)
	}

	if resp, err := client.GetDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a deployment configuration.
func codedeploy_GetDeploymentConfig(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.GetDeploymentConfigInput{
		// DeploymentConfigName: *string, // Required
	}

	if len(_codedeployDeploymentConfigName) > 0 {
		input.DeploymentConfigName = aws.String(_codedeployDeploymentConfigName)
	}

	if resp, err := client.GetDeploymentConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a deployment group.
func codedeploy_GetDeploymentGroup(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.GetDeploymentGroupInput{
		// ApplicationName: *string, // Required
		// DeploymentGroupName: *string, // Required
	}

	if len(_codedeployApplicationName) > 0 {
		input.ApplicationName = aws.String(_codedeployApplicationName)
	}
	if len(_codedeployDeploymentGroupName) > 0 {
		input.DeploymentGroupName = aws.String(_codedeployDeploymentGroupName)
	}

	if resp, err := client.GetDeploymentGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an instance as part of a deployment.
// Deprecated: This operation is deprecated, use GetDeploymentTarget instead.
func codedeploy_GetDeploymentInstance(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.GetDeploymentInstanceInput{
		// DeploymentId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_codedeployDeploymentId) > 0 {
		input.DeploymentId = aws.String(_codedeployDeploymentId)
	}
	if len(_codedeployInstanceId) > 0 {
		input.InstanceId = aws.String(_codedeployInstanceId)
	}

	if resp, err := client.GetDeploymentInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a deployment target.
func codedeploy_GetDeploymentTarget(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.GetDeploymentTargetInput{
		// DeploymentId: *string, // Required
		// TargetId: *string, // Required
	}

	if len(_codedeployDeploymentId) > 0 {
		input.DeploymentId = aws.String(_codedeployDeploymentId)
	}
	if len(_codedeployTargetId) > 0 {
		input.TargetId = aws.String(_codedeployTargetId)
	}

	if resp, err := client.GetDeploymentTarget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an on-premises instance.
func codedeploy_GetOnPremisesInstance(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.GetOnPremisesInstanceInput{
		// InstanceName: *string, // Required
	}

	if len(_codedeployInstanceName) > 0 {
		input.InstanceName = aws.String(_codedeployInstanceName)
	}

	if resp, err := client.GetOnPremisesInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists information about revisions for an application.
func codedeploy_ListApplicationRevisions(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.ListApplicationRevisionsInput{
		// ApplicationName: *string, // Required
	}

	if len(_codedeployApplicationName) > 0 {
		input.ApplicationName = aws.String(_codedeployApplicationName)
	}
	if len(_codedeployDeployed) > 0 {
		if err := assignInputField(input, "Deployed", _codedeployDeployed); err != nil {
			log.Errorf("invalid --deployed: %s", err.Error())
			return
		}
	}
	if len(_codedeployNextToken) > 0 {
		input.NextToken = aws.String(_codedeployNextToken)
	}
	if len(_codedeployS3Bucket) > 0 {
		input.S3Bucket = aws.String(_codedeployS3Bucket)
	}
	if len(_codedeployS3KeyPrefix) > 0 {
		input.S3KeyPrefix = aws.String(_codedeployS3KeyPrefix)
	}
	if len(_codedeploySortBy) > 0 {
		if err := assignInputField(input, "SortBy", _codedeploySortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_codedeploySortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _codedeploySortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListApplicationRevisions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codedeploy.ListApplicationRevisionsOutput
	p := codedeploy.NewListApplicationRevisionsPaginator(client, input)
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

// Lists the applications registered with the user or Amazon Web Services account.
func codedeploy_ListApplications(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.ListApplicationsInput{}

	if len(_codedeployNextToken) > 0 {
		input.NextToken = aws.String(_codedeployNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApplications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codedeploy.ListApplicationsOutput
	p := codedeploy.NewListApplicationsPaginator(client, input)
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

// Lists the deployment configurations with the user or Amazon Web Services
// account.
func codedeploy_ListDeploymentConfigs(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.ListDeploymentConfigsInput{}

	if len(_codedeployNextToken) > 0 {
		input.NextToken = aws.String(_codedeployNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDeploymentConfigs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codedeploy.ListDeploymentConfigsOutput
	p := codedeploy.NewListDeploymentConfigsPaginator(client, input)
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

// Lists the deployment groups for an application registered with the Amazon Web
// Services user or Amazon Web Services account.
func codedeploy_ListDeploymentGroups(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.ListDeploymentGroupsInput{
		// ApplicationName: *string, // Required
	}

	if len(_codedeployApplicationName) > 0 {
		input.ApplicationName = aws.String(_codedeployApplicationName)
	}
	if len(_codedeployNextToken) > 0 {
		input.NextToken = aws.String(_codedeployNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDeploymentGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codedeploy.ListDeploymentGroupsOutput
	p := codedeploy.NewListDeploymentGroupsPaginator(client, input)
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

// The newer BatchGetDeploymentTargets should be used instead because it works
// with all compute types. ListDeploymentInstances throws an exception if it is
// used with a compute platform other than EC2/On-premises or Lambda.
//
// Lists the instance for a deployment associated with the user or Amazon Web
// Services account.
//
// Deprecated: This operation is deprecated, use ListDeploymentTargets instead.
func codedeploy_ListDeploymentInstances(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.ListDeploymentInstancesInput{
		// DeploymentId: *string, // Required
	}

	if len(_codedeployDeploymentId) > 0 {
		input.DeploymentId = aws.String(_codedeployDeploymentId)
	}
	if len(_codedeployInstanceStatusFilter) > 0 {
		if err := assignInputField(input, "InstanceStatusFilter", _codedeployInstanceStatusFilter); err != nil {
			log.Errorf("invalid --instance-status-filter: %s", err.Error())
			return
		}
	}
	if len(_codedeployInstanceTypeFilter) > 0 {
		if err := assignInputField(input, "InstanceTypeFilter", _codedeployInstanceTypeFilter); err != nil {
			log.Errorf("invalid --instance-type-filter: %s", err.Error())
			return
		}
	}
	if len(_codedeployNextToken) > 0 {
		input.NextToken = aws.String(_codedeployNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDeploymentInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codedeploy.ListDeploymentInstancesOutput
	p := codedeploy.NewListDeploymentInstancesPaginator(client, input)
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

// Returns an array of target IDs that are associated a deployment.
func codedeploy_ListDeploymentTargets(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.ListDeploymentTargetsInput{
		// DeploymentId: *string, // Required
	}

	if len(_codedeployDeploymentId) > 0 {
		input.DeploymentId = aws.String(_codedeployDeploymentId)
	}
	if len(_codedeployNextToken) > 0 {
		input.NextToken = aws.String(_codedeployNextToken)
	}
	if len(_codedeployTargetFilters) > 0 {
		if err := assignInputField(input, "TargetFilters", _codedeployTargetFilters); err != nil {
			log.Errorf("invalid --target-filters: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListDeploymentTargets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the deployments in a deployment group for an application registered with
// the user or Amazon Web Services account.
func codedeploy_ListDeployments(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.ListDeploymentsInput{}

	if len(_codedeployApplicationName) > 0 {
		input.ApplicationName = aws.String(_codedeployApplicationName)
	}
	if len(_codedeployCreateTimeRange) > 0 {
		if err := assignInputField(input, "CreateTimeRange", _codedeployCreateTimeRange); err != nil {
			log.Errorf("invalid --create-time-range: %s", err.Error())
			return
		}
	}
	if len(_codedeployDeploymentGroupName) > 0 {
		input.DeploymentGroupName = aws.String(_codedeployDeploymentGroupName)
	}
	if len(_codedeployExternalId) > 0 {
		input.ExternalId = aws.String(_codedeployExternalId)
	}
	if len(_codedeployIncludeOnlyStatuses) > 0 {
		if err := assignInputField(input, "IncludeOnlyStatuses", _codedeployIncludeOnlyStatuses); err != nil {
			log.Errorf("invalid --include-only-statuses: %s", err.Error())
			return
		}
	}
	if len(_codedeployNextToken) > 0 {
		input.NextToken = aws.String(_codedeployNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDeployments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codedeploy.ListDeploymentsOutput
	p := codedeploy.NewListDeploymentsPaginator(client, input)
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

// Lists the names of stored connections to GitHub accounts.
func codedeploy_ListGitHubAccountTokenNames(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.ListGitHubAccountTokenNamesInput{}

	if len(_codedeployNextToken) > 0 {
		input.NextToken = aws.String(_codedeployNextToken)
	}

	if resp, err := client.ListGitHubAccountTokenNames(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of names for one or more on-premises instances.
// Unless otherwise specified, both registered and deregistered on-premises
// instance names are listed. To list only registered or deregistered on-premises
// instance names, use the registration status parameter.
func codedeploy_ListOnPremisesInstances(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.ListOnPremisesInstancesInput{}

	if len(_codedeployNextToken) > 0 {
		input.NextToken = aws.String(_codedeployNextToken)
	}
	if len(_codedeployRegistrationStatus) > 0 {
		if err := assignInputField(input, "RegistrationStatus", _codedeployRegistrationStatus); err != nil {
			log.Errorf("invalid --registration-status: %s", err.Error())
			return
		}
	}
	if len(_codedeployTagFilters) > 0 {
		if err := assignInputField(input, "TagFilters", _codedeployTagFilters); err != nil {
			log.Errorf("invalid --tag-filters: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListOnPremisesInstances(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of tags for the resource identified by a specified Amazon
// Resource Name (ARN). Tags are used to organize and categorize your CodeDeploy
// resources.
func codedeploy_ListTagsForResource(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_codedeployResourceArn) > 0 {
		input.ResourceArn = aws.String(_codedeployResourceArn)
	}
	if len(_codedeployNextToken) > 0 {
		input.NextToken = aws.String(_codedeployNextToken)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the result of a Lambda validation function. The function validates
// lifecycle hooks during a deployment that uses the Lambda or Amazon ECS compute
// platform. For Lambda deployments, the available lifecycle hooks are
// BeforeAllowTraffic and AfterAllowTraffic . For Amazon ECS deployments, the
// available lifecycle hooks are BeforeInstall , AfterInstall ,
// AfterAllowTestTraffic , BeforeAllowTraffic , and AfterAllowTraffic . Lambda
// validation functions return Succeeded or Failed . For more information, see [AppSpec 'hooks' Section for an Lambda Deployment]
// and [AppSpec 'hooks' Section for an Amazon ECS Deployment].
//
// [AppSpec 'hooks' Section for an Amazon ECS Deployment]: https://docs.aws.amazon.com/codedeploy/latest/userguide/reference-appspec-file-structure-hooks.html#appspec-hooks-ecs
// [AppSpec 'hooks' Section for an Lambda Deployment]: https://docs.aws.amazon.com/codedeploy/latest/userguide/reference-appspec-file-structure-hooks.html#appspec-hooks-lambda
func codedeploy_PutLifecycleEventHookExecutionStatus(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.PutLifecycleEventHookExecutionStatusInput{}

	if len(_codedeployDeploymentId) > 0 {
		input.DeploymentId = aws.String(_codedeployDeploymentId)
	}
	if len(_codedeployLifecycleEventHookExecutionId) > 0 {
		input.LifecycleEventHookExecutionId = aws.String(_codedeployLifecycleEventHookExecutionId)
	}
	if len(_codedeployStatus) > 0 {
		if err := assignInputField(input, "Status", _codedeployStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutLifecycleEventHookExecutionStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers with CodeDeploy a revision for the specified application.
func codedeploy_RegisterApplicationRevision(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.RegisterApplicationRevisionInput{
		// ApplicationName: *string, // Required
		// Revision: *types.RevisionLocation, // Required
	}

	if len(_codedeployApplicationName) > 0 {
		input.ApplicationName = aws.String(_codedeployApplicationName)
	}
	if len(_codedeployRevision) > 0 {
		if err := assignInputField(input, "Revision", _codedeployRevision); err != nil {
			log.Errorf("invalid --revision: %s", err.Error())
			return
		}
	}
	if len(_codedeployDescription) > 0 {
		input.Description = aws.String(_codedeployDescription)
	}

	if resp, err := client.RegisterApplicationRevision(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers an on-premises instance.
// Only one IAM ARN (an IAM session ARN or IAM user ARN) is supported in the
// request. You cannot use both.
func codedeploy_RegisterOnPremisesInstance(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.RegisterOnPremisesInstanceInput{
		// InstanceName: *string, // Required
	}

	if len(_codedeployInstanceName) > 0 {
		input.InstanceName = aws.String(_codedeployInstanceName)
	}
	if len(_codedeployIamSessionArn) > 0 {
		input.IamSessionArn = aws.String(_codedeployIamSessionArn)
	}
	if len(_codedeployIamUserArn) > 0 {
		input.IamUserArn = aws.String(_codedeployIamUserArn)
	}

	if resp, err := client.RegisterOnPremisesInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes one or more tags from one or more on-premises instances.
func codedeploy_RemoveTagsFromOnPremisesInstances(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.RemoveTagsFromOnPremisesInstancesInput{
		// InstanceNames: []string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_codedeployInstanceNames) > 0 {
		input.InstanceNames = append([]string(nil), _codedeployInstanceNames...)
	}
	if len(_codedeployTags) > 0 {
		if err := assignInputField(input, "Tags", _codedeployTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.RemoveTagsFromOnPremisesInstances(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// In a blue/green deployment, overrides any specified wait time and starts
// terminating instances immediately after the traffic routing is complete.
//
// Deprecated: This operation is deprecated, use ContinueDeployment with
// DeploymentWaitType instead.
func codedeploy_SkipWaitTimeForInstanceTermination(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.SkipWaitTimeForInstanceTerminationInput{}

	if len(_codedeployDeploymentId) > 0 {
		input.DeploymentId = aws.String(_codedeployDeploymentId)
	}

	if resp, err := client.SkipWaitTimeForInstanceTermination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attempts to stop an ongoing deployment.
func codedeploy_StopDeployment(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.StopDeploymentInput{
		// DeploymentId: *string, // Required
	}

	if len(_codedeployDeploymentId) > 0 {
		input.DeploymentId = aws.String(_codedeployDeploymentId)
	}
	if len(_codedeployAutoRollbackEnabled) > 0 {
		if err := assignInputField(input, "AutoRollbackEnabled", _codedeployAutoRollbackEnabled); err != nil {
			log.Errorf("invalid --auto-rollback-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.StopDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the list of tags in the input Tags parameter with the resource
// identified by the ResourceArn input parameter.
func codedeploy_TagResource(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_codedeployResourceArn) > 0 {
		input.ResourceArn = aws.String(_codedeployResourceArn)
	}
	if len(_codedeployTags) > 0 {
		if err := assignInputField(input, "Tags", _codedeployTags); err != nil {
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

// Disassociates a resource from a list of tags. The resource is identified by
// the ResourceArn input parameter. The tags are identified by the list of keys in
// the TagKeys input parameter.
func codedeploy_UntagResource(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_codedeployResourceArn) > 0 {
		input.ResourceArn = aws.String(_codedeployResourceArn)
	}
	if len(_codedeployTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _codedeployTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the name of an application.
func codedeploy_UpdateApplication(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.UpdateApplicationInput{}

	if len(_codedeployApplicationName) > 0 {
		input.ApplicationName = aws.String(_codedeployApplicationName)
	}
	if len(_codedeployNewApplicationName) > 0 {
		input.NewApplicationName = aws.String(_codedeployNewApplicationName)
	}

	if resp, err := client.UpdateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes information about a deployment group.
func codedeploy_UpdateDeploymentGroup(cfg aws.Config, client *codedeploy.Client) {
	input := &codedeploy.UpdateDeploymentGroupInput{
		// ApplicationName: *string, // Required
		// CurrentDeploymentGroupName: *string, // Required
	}

	if len(_codedeployApplicationName) > 0 {
		input.ApplicationName = aws.String(_codedeployApplicationName)
	}
	if len(_codedeployCurrentDeploymentGroupName) > 0 {
		input.CurrentDeploymentGroupName = aws.String(_codedeployCurrentDeploymentGroupName)
	}
	if len(_codedeployAlarmConfiguration) > 0 {
		if err := assignInputField(input, "AlarmConfiguration", _codedeployAlarmConfiguration); err != nil {
			log.Errorf("invalid --alarm-configuration: %s", err.Error())
			return
		}
	}
	if len(_codedeployAutoRollbackConfiguration) > 0 {
		if err := assignInputField(input, "AutoRollbackConfiguration", _codedeployAutoRollbackConfiguration); err != nil {
			log.Errorf("invalid --auto-rollback-configuration: %s", err.Error())
			return
		}
	}
	if len(_codedeployAutoScalingGroups) > 0 {
		input.AutoScalingGroups = append([]string(nil), _codedeployAutoScalingGroups...)
	}
	if len(_codedeployBlueGreenDeploymentConfiguration) > 0 {
		if err := assignInputField(input, "BlueGreenDeploymentConfiguration", _codedeployBlueGreenDeploymentConfiguration); err != nil {
			log.Errorf("invalid --blue-green-deployment-configuration: %s", err.Error())
			return
		}
	}
	if len(_codedeployDeploymentConfigName) > 0 {
		input.DeploymentConfigName = aws.String(_codedeployDeploymentConfigName)
	}
	if len(_codedeployDeploymentStyle) > 0 {
		if err := assignInputField(input, "DeploymentStyle", _codedeployDeploymentStyle); err != nil {
			log.Errorf("invalid --deployment-style: %s", err.Error())
			return
		}
	}
	if len(_codedeployEc2TagFilters) > 0 {
		if err := assignInputField(input, "Ec2TagFilters", _codedeployEc2TagFilters); err != nil {
			log.Errorf("invalid --ec2-tag-filters: %s", err.Error())
			return
		}
	}
	if len(_codedeployEc2TagSet) > 0 {
		if err := assignInputField(input, "Ec2TagSet", _codedeployEc2TagSet); err != nil {
			log.Errorf("invalid --ec2-tag-set: %s", err.Error())
			return
		}
	}
	if len(_codedeployEcsServices) > 0 {
		if err := assignInputField(input, "EcsServices", _codedeployEcsServices); err != nil {
			log.Errorf("invalid --ecs-services: %s", err.Error())
			return
		}
	}
	if len(_codedeployLoadBalancerInfo) > 0 {
		if err := assignInputField(input, "LoadBalancerInfo", _codedeployLoadBalancerInfo); err != nil {
			log.Errorf("invalid --load-balancer-info: %s", err.Error())
			return
		}
	}
	if len(_codedeployNewDeploymentGroupName) > 0 {
		input.NewDeploymentGroupName = aws.String(_codedeployNewDeploymentGroupName)
	}
	if len(_codedeployOnPremisesInstanceTagFilters) > 0 {
		if err := assignInputField(input, "OnPremisesInstanceTagFilters", _codedeployOnPremisesInstanceTagFilters); err != nil {
			log.Errorf("invalid --on-premises-instance-tag-filters: %s", err.Error())
			return
		}
	}
	if len(_codedeployOnPremisesTagSet) > 0 {
		if err := assignInputField(input, "OnPremisesTagSet", _codedeployOnPremisesTagSet); err != nil {
			log.Errorf("invalid --on-premises-tag-set: %s", err.Error())
			return
		}
	}
	if len(_codedeployOutdatedInstancesStrategy) > 0 {
		if err := assignInputField(input, "OutdatedInstancesStrategy", _codedeployOutdatedInstancesStrategy); err != nil {
			log.Errorf("invalid --outdated-instances-strategy: %s", err.Error())
			return
		}
	}
	if len(_codedeployServiceRoleArn) > 0 {
		input.ServiceRoleArn = aws.String(_codedeployServiceRoleArn)
	}
	if len(_codedeployTerminationHookEnabled) > 0 {
		if err := assignInputField(input, "TerminationHookEnabled", _codedeployTerminationHookEnabled); err != nil {
			log.Errorf("invalid --termination-hook-enabled: %s", err.Error())
			return
		}
	}
	if len(_codedeployTriggerConfigurations) > 0 {
		if err := assignInputField(input, "TriggerConfigurations", _codedeployTriggerConfigurations); err != nil {
			log.Errorf("invalid --trigger-configurations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDeploymentGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_codedeployCmd)
	_codedeployCmd.Flags().SortFlags = false

	_codedeployCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_codedeployCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_codedeployCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_codedeployCmd.Flags().StringVarP(&_codedeployAlarmConfiguration, "alarm-configuration", "", "", "Alarm Configuration")
	_codedeployCmd.Flags().StringVarP(&_codedeployApplicationName, "application-name", "", "", "Application Name")
	_codedeployCmd.Flags().StringSliceVarP(&_codedeployApplicationNames, "application-names", "", nil, "Application Names")
	_codedeployCmd.Flags().StringVarP(&_codedeployAutoRollbackConfiguration, "auto-rollback-configuration", "", "", "Auto Rollback Configuration")
	_codedeployCmd.Flags().StringVarP(&_codedeployAutoRollbackEnabled, "auto-rollback-enabled", "", "", "Auto Rollback Enabled")
	_codedeployCmd.Flags().StringSliceVarP(&_codedeployAutoScalingGroups, "auto-scaling-groups", "", nil, "Auto Scaling Groups")
	_codedeployCmd.Flags().StringVarP(&_codedeployBlueGreenDeploymentConfiguration, "blue-green-deployment-configuration", "", "", "Blue Green Deployment Configuration")
	_codedeployCmd.Flags().StringVarP(&_codedeployComputePlatform, "compute-platform", "", "", "Compute Platform")
	_codedeployCmd.Flags().StringVarP(&_codedeployCreateTimeRange, "create-time-range", "", "", "Create Time Range")
	_codedeployCmd.Flags().StringVarP(&_codedeployCurrentDeploymentGroupName, "current-deployment-group-name", "", "", "Current Deployment Group Name")
	_codedeployCmd.Flags().StringVarP(&_codedeployDeployed, "deployed", "", "", "Deployed")
	_codedeployCmd.Flags().StringVarP(&_codedeployDeploymentConfigName, "deployment-config-name", "", "", "Deployment Config Name")
	_codedeployCmd.Flags().StringVarP(&_codedeployDeploymentGroupName, "deployment-group-name", "", "", "Deployment Group Name")
	_codedeployCmd.Flags().StringSliceVarP(&_codedeployDeploymentGroupNames, "deployment-group-names", "", nil, "Deployment Group Names")
	_codedeployCmd.Flags().StringVarP(&_codedeployDeploymentId, "deployment-id", "", "", "Deployment ID")
	_codedeployCmd.Flags().StringSliceVarP(&_codedeployDeploymentIds, "deployment-ids", "", nil, "Deployment Ids")
	_codedeployCmd.Flags().StringVarP(&_codedeployDeploymentStyle, "deployment-style", "", "", "Deployment Style")
	_codedeployCmd.Flags().StringVarP(&_codedeployDeploymentWaitType, "deployment-wait-type", "", "", "Deployment Wait Type")
	_codedeployCmd.Flags().StringVarP(&_codedeployDescription, "description", "", "", "Description")
	_codedeployCmd.Flags().StringVarP(&_codedeployEc2TagFilters, "ec2-tag-filters", "", "", "EC2 Tag Filters")
	_codedeployCmd.Flags().StringVarP(&_codedeployEc2TagSet, "ec2-tag-set", "", "", "EC2 Tag Set")
	_codedeployCmd.Flags().StringVarP(&_codedeployEcsServices, "ecs-services", "", "", "Ecs Services")
	_codedeployCmd.Flags().StringVarP(&_codedeployExternalId, "external-id", "", "", "External ID")
	_codedeployCmd.Flags().StringVarP(&_codedeployFileExistsBehavior, "file-exists-behavior", "", "", "File Exists Behavior")
	_codedeployCmd.Flags().StringVarP(&_codedeployIamSessionArn, "iam-session-arn", "", "", "IAM Session ARN")
	_codedeployCmd.Flags().StringVarP(&_codedeployIamUserArn, "iam-user-arn", "", "", "IAM User ARN")
	_codedeployCmd.Flags().StringVarP(&_codedeployIgnoreApplicationStopFailures, "ignore-application-stop-failures", "", "", "Ignore Application Stop Failures")
	_codedeployCmd.Flags().StringVarP(&_codedeployIncludeOnlyStatuses, "include-only-statuses", "", "", "Include Only Statuses")
	_codedeployCmd.Flags().StringVarP(&_codedeployInstanceId, "instance-id", "", "", "Instance ID")
	_codedeployCmd.Flags().StringSliceVarP(&_codedeployInstanceIds, "instance-ids", "", nil, "Instance Ids")
	_codedeployCmd.Flags().StringVarP(&_codedeployInstanceName, "instance-name", "", "", "Instance Name")
	_codedeployCmd.Flags().StringSliceVarP(&_codedeployInstanceNames, "instance-names", "", nil, "Instance Names")
	_codedeployCmd.Flags().StringVarP(&_codedeployInstanceStatusFilter, "instance-status-filter", "", "", "Instance Status Filter")
	_codedeployCmd.Flags().StringVarP(&_codedeployInstanceTypeFilter, "instance-type-filter", "", "", "Instance Type Filter")
	_codedeployCmd.Flags().StringVarP(&_codedeployLifecycleEventHookExecutionId, "lifecycle-event-hook-execution-id", "", "", "Lifecycle Event Hook Execution ID")
	_codedeployCmd.Flags().StringVarP(&_codedeployLoadBalancerInfo, "load-balancer-info", "", "", "Load Balancer Info")
	_codedeployCmd.Flags().StringVarP(&_codedeployMinimumHealthyHosts, "minimum-healthy-hosts", "", "", "Minimum Healthy Hosts")
	_codedeployCmd.Flags().StringVarP(&_codedeployNewApplicationName, "new-application-name", "", "", "New Application Name")
	_codedeployCmd.Flags().StringVarP(&_codedeployNewDeploymentGroupName, "new-deployment-group-name", "", "", "New Deployment Group Name")
	_codedeployCmd.Flags().StringVarP(&_codedeployNextToken, "next-token", "", "", "Next Token")
	_codedeployCmd.Flags().StringVarP(&_codedeployOnPremisesInstanceTagFilters, "on-premises-instance-tag-filters", "", "", "On Premises Instance Tag Filters")
	_codedeployCmd.Flags().StringVarP(&_codedeployOnPremisesTagSet, "on-premises-tag-set", "", "", "On Premises Tag Set")
	_codedeployCmd.Flags().StringVarP(&_codedeployOutdatedInstancesStrategy, "outdated-instances-strategy", "", "", "Outdated Instances Strategy")
	_codedeployCmd.Flags().StringVarP(&_codedeployOverrideAlarmConfiguration, "override-alarm-configuration", "", "", "Override Alarm Configuration")
	_codedeployCmd.Flags().StringVarP(&_codedeployRegistrationStatus, "registration-status", "", "", "Registration Status")
	_codedeployCmd.Flags().StringVarP(&_codedeployResourceArn, "resource-arn", "", "", "Resource ARN")
	_codedeployCmd.Flags().StringVarP(&_codedeployRevision, "revision", "", "", "Revision")
	_codedeployCmd.Flags().StringVarP(&_codedeployRevisions, "revisions", "", "", "Revisions")
	_codedeployCmd.Flags().StringVarP(&_codedeployS3Bucket, "s3-bucket", "", "", "S3 Bucket")
	_codedeployCmd.Flags().StringVarP(&_codedeployS3KeyPrefix, "s3-key-prefix", "", "", "S3 Key Prefix")
	_codedeployCmd.Flags().StringVarP(&_codedeployServiceRoleArn, "service-role-arn", "", "", "Service Role ARN")
	_codedeployCmd.Flags().StringVarP(&_codedeploySortBy, "sort-by", "", "", "Sort By")
	_codedeployCmd.Flags().StringVarP(&_codedeploySortOrder, "sort-order", "", "", "Sort Order")
	_codedeployCmd.Flags().StringVarP(&_codedeployStatus, "status", "", "", "Status")
	_codedeployCmd.Flags().StringVarP(&_codedeployTagFilters, "tag-filters", "", "", "Tag Filters")
	_codedeployCmd.Flags().StringSliceVarP(&_codedeployTagKeys, "tag-keys", "", nil, "Tag Keys")
	_codedeployCmd.Flags().StringVarP(&_codedeployTags, "tags", "", "", "Tags")
	_codedeployCmd.Flags().StringVarP(&_codedeployTargetFilters, "target-filters", "", "", "Target Filters")
	_codedeployCmd.Flags().StringVarP(&_codedeployTargetId, "target-id", "", "", "Target ID")
	_codedeployCmd.Flags().StringSliceVarP(&_codedeployTargetIds, "target-ids", "", nil, "Target Ids")
	_codedeployCmd.Flags().StringVarP(&_codedeployTargetInstances, "target-instances", "", "", "Target Instances")
	_codedeployCmd.Flags().StringVarP(&_codedeployTerminationHookEnabled, "termination-hook-enabled", "", "", "Termination Hook Enabled")
	_codedeployCmd.Flags().StringVarP(&_codedeployTokenName, "token-name", "", "", "Token Name")
	_codedeployCmd.Flags().StringVarP(&_codedeployTrafficRoutingConfig, "traffic-routing-config", "", "", "Traffic Routing Config")
	_codedeployCmd.Flags().StringVarP(&_codedeployTriggerConfigurations, "trigger-configurations", "", "", "Trigger Configurations")
	_codedeployCmd.Flags().StringVarP(&_codedeployUpdateOutdatedInstancesOnly, "update-outdated-instances-only", "", "", "Update Outdated Instances Only")
	_codedeployCmd.Flags().StringVarP(&_codedeployZonalConfig, "zonal-config", "", "", "Zonal Config")

	_codedeployCmd.Flags().BoolVarP(&_codedeployAddTagsToOnPremisesInstances, "add-tags-to-on-premises-instances", "", false, "Add Tags To On Premises Instances")
	_codedeployCmd.Flags().BoolVarP(&_codedeployBatchGetApplicationRevisions, "batch-get-application-revisions", "", false, "Batch Get Application Revisions")
	_codedeployCmd.Flags().BoolVarP(&_codedeployBatchGetApplications, "batch-get-applications", "", false, "Batch Get Applications")
	_codedeployCmd.Flags().BoolVarP(&_codedeployBatchGetDeploymentGroups, "batch-get-deployment-groups", "", false, "Batch Get Deployment Groups")
	_codedeployCmd.Flags().BoolVarP(&_codedeployBatchGetDeploymentInstances, "batch-get-deployment-instances", "", false, "Batch Get Deployment Instances")
	_codedeployCmd.Flags().BoolVarP(&_codedeployBatchGetDeploymentTargets, "batch-get-deployment-targets", "", false, "Batch Get Deployment Targets")
	_codedeployCmd.Flags().BoolVarP(&_codedeployBatchGetDeployments, "batch-get-deployments", "", false, "Batch Get Deployments")
	_codedeployCmd.Flags().BoolVarP(&_codedeployBatchGetOnPremisesInstances, "batch-get-on-premises-instances", "", false, "Batch Get On Premises Instances")
	_codedeployCmd.Flags().BoolVarP(&_codedeployContinueDeployment, "continue-deployment", "", false, "Continue Deployment")
	_codedeployCmd.Flags().BoolVarP(&_codedeployCreateApplication, "create-application", "", false, "Create Application")
	_codedeployCmd.Flags().BoolVarP(&_codedeployCreateDeployment, "create-deployment", "", false, "Create Deployment")
	_codedeployCmd.Flags().BoolVarP(&_codedeployCreateDeploymentConfig, "create-deployment-config", "", false, "Create Deployment Config")
	_codedeployCmd.Flags().BoolVarP(&_codedeployCreateDeploymentGroup, "create-deployment-group", "", false, "Create Deployment Group")
	_codedeployCmd.Flags().BoolVarP(&_codedeployDeleteApplication, "delete-application", "", false, "Delete Application")
	_codedeployCmd.Flags().BoolVarP(&_codedeployDeleteDeploymentConfig, "delete-deployment-config", "", false, "Delete Deployment Config")
	_codedeployCmd.Flags().BoolVarP(&_codedeployDeleteDeploymentGroup, "delete-deployment-group", "", false, "Delete Deployment Group")
	_codedeployCmd.Flags().BoolVarP(&_codedeployDeleteGitHubAccountToken, "delete-git-hub-account-token", "", false, "Delete Git Hub Account Token")
	_codedeployCmd.Flags().BoolVarP(&_codedeployDeleteResourcesByExternalId, "delete-resources-by-external-id", "", false, "Delete Resources By External ID")
	_codedeployCmd.Flags().BoolVarP(&_codedeployDeregisterOnPremisesInstance, "deregister-on-premises-instance", "", false, "Deregister On Premises Instance")
	_codedeployCmd.Flags().BoolVarP(&_codedeployGetApplication, "get-application", "", false, "Get Application")
	_codedeployCmd.Flags().BoolVarP(&_codedeployGetApplicationRevision, "get-application-revision", "", false, "Get Application Revision")
	_codedeployCmd.Flags().BoolVarP(&_codedeployGetDeployment, "get-deployment", "", false, "Get Deployment")
	_codedeployCmd.Flags().BoolVarP(&_codedeployGetDeploymentConfig, "get-deployment-config", "", false, "Get Deployment Config")
	_codedeployCmd.Flags().BoolVarP(&_codedeployGetDeploymentGroup, "get-deployment-group", "", false, "Get Deployment Group")
	_codedeployCmd.Flags().BoolVarP(&_codedeployGetDeploymentInstance, "get-deployment-instance", "", false, "Get Deployment Instance")
	_codedeployCmd.Flags().BoolVarP(&_codedeployGetDeploymentTarget, "get-deployment-target", "", false, "Get Deployment Target")
	_codedeployCmd.Flags().BoolVarP(&_codedeployGetOnPremisesInstance, "get-on-premises-instance", "", false, "Get On Premises Instance")
	_codedeployCmd.Flags().BoolVarP(&_codedeployListApplicationRevisions, "list-application-revisions", "", false, "List Application Revisions")
	_codedeployCmd.Flags().BoolVarP(&_codedeployListApplications, "list-applications", "", false, "List Applications")
	_codedeployCmd.Flags().BoolVarP(&_codedeployListDeploymentConfigs, "list-deployment-configs", "", false, "List Deployment Configs")
	_codedeployCmd.Flags().BoolVarP(&_codedeployListDeploymentGroups, "list-deployment-groups", "", false, "List Deployment Groups")
	_codedeployCmd.Flags().BoolVarP(&_codedeployListDeploymentInstances, "list-deployment-instances", "", false, "List Deployment Instances")
	_codedeployCmd.Flags().BoolVarP(&_codedeployListDeploymentTargets, "list-deployment-targets", "", false, "List Deployment Targets")
	_codedeployCmd.Flags().BoolVarP(&_codedeployListDeployments, "list-deployments", "", false, "List Deployments")
	_codedeployCmd.Flags().BoolVarP(&_codedeployListGitHubAccountTokenNames, "list-git-hub-account-token-names", "", false, "List Git Hub Account Token Names")
	_codedeployCmd.Flags().BoolVarP(&_codedeployListOnPremisesInstances, "list-on-premises-instances", "", false, "List On Premises Instances")
	_codedeployCmd.Flags().BoolVarP(&_codedeployListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_codedeployCmd.Flags().BoolVarP(&_codedeployPutLifecycleEventHookExecutionStatus, "put-lifecycle-event-hook-execution-status", "", false, "Put Lifecycle Event Hook Execution Status")
	_codedeployCmd.Flags().BoolVarP(&_codedeployRegisterApplicationRevision, "register-application-revision", "", false, "Register Application Revision")
	_codedeployCmd.Flags().BoolVarP(&_codedeployRegisterOnPremisesInstance, "register-on-premises-instance", "", false, "Register On Premises Instance")
	_codedeployCmd.Flags().BoolVarP(&_codedeployRemoveTagsFromOnPremisesInstances, "remove-tags-from-on-premises-instances", "", false, "Remove Tags From On Premises Instances")
	_codedeployCmd.Flags().BoolVarP(&_codedeploySkipWaitTimeForInstanceTermination, "skip-wait-time-for-instance-termination", "", false, "Skip Wait Time For Instance Termination")
	_codedeployCmd.Flags().BoolVarP(&_codedeployStopDeployment, "stop-deployment", "", false, "Stop Deployment")
	_codedeployCmd.Flags().BoolVarP(&_codedeployTagResource, "tag-resource", "", false, "Tag Resource")
	_codedeployCmd.Flags().BoolVarP(&_codedeployUntagResource, "untag-resource", "", false, "Untag Resource")
	_codedeployCmd.Flags().BoolVarP(&_codedeployUpdateApplication, "update-application", "", false, "Update Application")
	_codedeployCmd.Flags().BoolVarP(&_codedeployUpdateDeploymentGroup, "update-deployment-group", "", false, "Update Deployment Group")

}
