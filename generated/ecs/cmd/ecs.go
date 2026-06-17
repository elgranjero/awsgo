package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ecsCmd represents the ecs command
var _ecsCmd = &cobra.Command{
	Use:   "ecs",
	Short: "AWS ecs CLI",
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
		client := ecs.NewFromConfig(cfg)
		if _ecsCreateCapacityProvider {
			ecs_CreateCapacityProvider(cfg, client)
			return
		}
		if _ecsCreateCluster {
			ecs_CreateCluster(cfg, client)
			return
		}
		if _ecsCreateExpressGatewayService {
			ecs_CreateExpressGatewayService(cfg, client)
			return
		}
		if _ecsCreateService {
			ecs_CreateService(cfg, client)
			return
		}
		if _ecsCreateTaskSet {
			ecs_CreateTaskSet(cfg, client)
			return
		}
		if _ecsDeleteAccountSetting {
			ecs_DeleteAccountSetting(cfg, client)
			return
		}
		if _ecsDeleteAttributes {
			ecs_DeleteAttributes(cfg, client)
			return
		}
		if _ecsDeleteCapacityProvider {
			ecs_DeleteCapacityProvider(cfg, client)
			return
		}
		if _ecsDeleteCluster {
			ecs_DeleteCluster(cfg, client)
			return
		}
		if _ecsDeleteExpressGatewayService {
			ecs_DeleteExpressGatewayService(cfg, client)
			return
		}
		if _ecsDeleteService {
			ecs_DeleteService(cfg, client)
			return
		}
		if _ecsDeleteTaskDefinitions {
			ecs_DeleteTaskDefinitions(cfg, client)
			return
		}
		if _ecsDeleteTaskSet {
			ecs_DeleteTaskSet(cfg, client)
			return
		}
		if _ecsDeregisterContainerInstance {
			ecs_DeregisterContainerInstance(cfg, client)
			return
		}
		if _ecsDeregisterTaskDefinition {
			ecs_DeregisterTaskDefinition(cfg, client)
			return
		}
		if _ecsDescribeCapacityProviders {
			ecs_DescribeCapacityProviders(cfg, client)
			return
		}
		if _ecsDescribeClusters {
			ecs_DescribeClusters(cfg, client)
			return
		}
		if _ecsDescribeContainerInstances {
			ecs_DescribeContainerInstances(cfg, client)
			return
		}
		if _ecsDescribeExpressGatewayService {
			ecs_DescribeExpressGatewayService(cfg, client)
			return
		}
		if _ecsDescribeServiceDeployments {
			ecs_DescribeServiceDeployments(cfg, client)
			return
		}
		if _ecsDescribeServiceRevisions {
			ecs_DescribeServiceRevisions(cfg, client)
			return
		}
		if _ecsDescribeServices {
			ecs_DescribeServices(cfg, client)
			return
		}
		if _ecsDescribeTaskDefinition {
			ecs_DescribeTaskDefinition(cfg, client)
			return
		}
		if _ecsDescribeTaskSets {
			ecs_DescribeTaskSets(cfg, client)
			return
		}
		if _ecsDescribeTasks {
			ecs_DescribeTasks(cfg, client)
			return
		}
		if _ecsDiscoverPollEndpoint {
			ecs_DiscoverPollEndpoint(cfg, client)
			return
		}
		if _ecsExecuteCommand {
			ecs_ExecuteCommand(cfg, client)
			return
		}
		if _ecsGetTaskProtection {
			ecs_GetTaskProtection(cfg, client)
			return
		}
		if _ecsListAccountSettings {
			ecs_ListAccountSettings(cfg, client)
			return
		}
		if _ecsListAttributes {
			ecs_ListAttributes(cfg, client)
			return
		}
		if _ecsListClusters {
			ecs_ListClusters(cfg, client)
			return
		}
		if _ecsListContainerInstances {
			ecs_ListContainerInstances(cfg, client)
			return
		}
		if _ecsListServiceDeployments {
			ecs_ListServiceDeployments(cfg, client)
			return
		}
		if _ecsListServices {
			ecs_ListServices(cfg, client)
			return
		}
		if _ecsListServicesByNamespace {
			ecs_ListServicesByNamespace(cfg, client)
			return
		}
		if _ecsListTagsForResource {
			ecs_ListTagsForResource(cfg, client)
			return
		}
		if _ecsListTaskDefinitionFamilies {
			ecs_ListTaskDefinitionFamilies(cfg, client)
			return
		}
		if _ecsListTaskDefinitions {
			ecs_ListTaskDefinitions(cfg, client)
			return
		}
		if _ecsListTasks {
			ecs_ListTasks(cfg, client)
			return
		}
		if _ecsPutAccountSetting {
			ecs_PutAccountSetting(cfg, client)
			return
		}
		if _ecsPutAccountSettingDefault {
			ecs_PutAccountSettingDefault(cfg, client)
			return
		}
		if _ecsPutAttributes {
			ecs_PutAttributes(cfg, client)
			return
		}
		if _ecsPutClusterCapacityProviders {
			ecs_PutClusterCapacityProviders(cfg, client)
			return
		}
		if _ecsRegisterContainerInstance {
			ecs_RegisterContainerInstance(cfg, client)
			return
		}
		if _ecsRegisterTaskDefinition {
			ecs_RegisterTaskDefinition(cfg, client)
			return
		}
		if _ecsRunTask {
			ecs_RunTask(cfg, client)
			return
		}
		if _ecsStartTask {
			ecs_StartTask(cfg, client)
			return
		}
		if _ecsStopServiceDeployment {
			ecs_StopServiceDeployment(cfg, client)
			return
		}
		if _ecsStopTask {
			ecs_StopTask(cfg, client)
			return
		}
		if _ecsSubmitAttachmentStateChanges {
			ecs_SubmitAttachmentStateChanges(cfg, client)
			return
		}
		if _ecsSubmitContainerStateChange {
			ecs_SubmitContainerStateChange(cfg, client)
			return
		}
		if _ecsSubmitTaskStateChange {
			ecs_SubmitTaskStateChange(cfg, client)
			return
		}
		if _ecsTagResource {
			ecs_TagResource(cfg, client)
			return
		}
		if _ecsUntagResource {
			ecs_UntagResource(cfg, client)
			return
		}
		if _ecsUpdateCapacityProvider {
			ecs_UpdateCapacityProvider(cfg, client)
			return
		}
		if _ecsUpdateCluster {
			ecs_UpdateCluster(cfg, client)
			return
		}
		if _ecsUpdateClusterSettings {
			ecs_UpdateClusterSettings(cfg, client)
			return
		}
		if _ecsUpdateContainerAgent {
			ecs_UpdateContainerAgent(cfg, client)
			return
		}
		if _ecsUpdateContainerInstancesState {
			ecs_UpdateContainerInstancesState(cfg, client)
			return
		}
		if _ecsUpdateExpressGatewayService {
			ecs_UpdateExpressGatewayService(cfg, client)
			return
		}
		if _ecsUpdateService {
			ecs_UpdateService(cfg, client)
			return
		}
		if _ecsUpdateServicePrimaryTaskSet {
			ecs_UpdateServicePrimaryTaskSet(cfg, client)
			return
		}
		if _ecsUpdateTaskProtection {
			ecs_UpdateTaskProtection(cfg, client)
			return
		}
		if _ecsUpdateTaskSet {
			ecs_UpdateTaskSet(cfg, client)
			return
		}

	},
}

var (
	_ecsCreateCapacityProvider        bool
	_ecsCreateCluster                 bool
	_ecsCreateExpressGatewayService   bool
	_ecsCreateService                 bool
	_ecsCreateTaskSet                 bool
	_ecsDeleteAccountSetting          bool
	_ecsDeleteAttributes              bool
	_ecsDeleteCapacityProvider        bool
	_ecsDeleteCluster                 bool
	_ecsDeleteExpressGatewayService   bool
	_ecsDeleteService                 bool
	_ecsDeleteTaskDefinitions         bool
	_ecsDeleteTaskSet                 bool
	_ecsDeregisterContainerInstance   bool
	_ecsDeregisterTaskDefinition      bool
	_ecsDescribeCapacityProviders     bool
	_ecsDescribeClusters              bool
	_ecsDescribeContainerInstances    bool
	_ecsDescribeExpressGatewayService bool
	_ecsDescribeServiceDeployments    bool
	_ecsDescribeServiceRevisions      bool
	_ecsDescribeServices              bool
	_ecsDescribeTaskDefinition        bool
	_ecsDescribeTaskSets              bool
	_ecsDescribeTasks                 bool
	_ecsDiscoverPollEndpoint          bool
	_ecsExecuteCommand                bool
	_ecsGetTaskProtection             bool
	_ecsListAccountSettings           bool
	_ecsListAttributes                bool
	_ecsListClusters                  bool
	_ecsListContainerInstances        bool
	_ecsListServiceDeployments        bool
	_ecsListServices                  bool
	_ecsListServicesByNamespace       bool
	_ecsListTagsForResource           bool
	_ecsListTaskDefinitionFamilies    bool
	_ecsListTaskDefinitions           bool
	_ecsListTasks                     bool
	_ecsPutAccountSetting             bool
	_ecsPutAccountSettingDefault      bool
	_ecsPutAttributes                 bool
	_ecsPutClusterCapacityProviders   bool
	_ecsRegisterContainerInstance     bool
	_ecsRegisterTaskDefinition        bool
	_ecsRunTask                       bool
	_ecsStartTask                     bool
	_ecsStopServiceDeployment         bool
	_ecsStopTask                      bool
	_ecsSubmitAttachmentStateChanges  bool
	_ecsSubmitContainerStateChange    bool
	_ecsSubmitTaskStateChange         bool
	_ecsTagResource                   bool
	_ecsUntagResource                 bool
	_ecsUpdateCapacityProvider        bool
	_ecsUpdateCluster                 bool
	_ecsUpdateClusterSettings         bool
	_ecsUpdateContainerAgent          bool
	_ecsUpdateContainerInstancesState bool
	_ecsUpdateExpressGatewayService   bool
	_ecsUpdateService                 bool
	_ecsUpdateServicePrimaryTaskSet   bool
	_ecsUpdateTaskProtection          bool
	_ecsUpdateTaskSet                 bool

	_ecsAttachments                       string
	_ecsAttributeName                     string
	_ecsAttributeValue                    string
	_ecsAttributes                        string
	_ecsAutoScalingGroupProvider          string
	_ecsAvailabilityZoneRebalancing       string
	_ecsCapacityProvider                  string
	_ecsCapacityProviderStrategy          string
	_ecsCapacityProviders                 []string
	_ecsClientToken                       string
	_ecsCluster                           string
	_ecsClusterName                       string
	_ecsClusters                          []string
	_ecsCommand                           string
	_ecsConfiguration                     string
	_ecsContainer                         string
	_ecsContainerDefinitions              string
	_ecsContainerInstance                 string
	_ecsContainerInstanceArn              string
	_ecsContainerInstances                []string
	_ecsContainerName                     string
	_ecsContainers                        string
	_ecsCount                             string
	_ecsCpu                               string
	_ecsCreatedAt                         string
	_ecsDefaultCapacityProviderStrategy   string
	_ecsDeploymentConfiguration           string
	_ecsDeploymentController              string
	_ecsDesiredCount                      string
	_ecsDesiredStatus                     string
	_ecsEffectiveSettings                 string
	_ecsEnableECSManagedTags              string
	_ecsEnableExecuteCommand              string
	_ecsEnableFaultInjection              string
	_ecsEphemeralStorage                  string
	_ecsExecutionRoleArn                  string
	_ecsExecutionStoppedAt                string
	_ecsExitCode                          string
	_ecsExpiresInMinutes                  string
	_ecsExternalId                        string
	_ecsFamily                            string
	_ecsFamilyPrefix                      string
	_ecsFilter                            string
	_ecsForce                             string
	_ecsForceNewDeployment                string
	_ecsGroup                             string
	_ecsHealthCheckGracePeriodSeconds     string
	_ecsHealthCheckPath                   string
	_ecsInclude                           string
	_ecsInferenceAccelerators             string
	_ecsInfrastructureRoleArn             string
	_ecsInstanceIdentityDocument          string
	_ecsInstanceIdentityDocumentSignature string
	_ecsInteractive                       string
	_ecsIpcMode                           string
	_ecsLaunchType                        string
	_ecsLoadBalancers                     string
	_ecsManagedAgents                     string
	_ecsManagedInstancesProvider          string
	_ecsMaxResults                        string
	_ecsMemory                            string
	_ecsName                              string
	_ecsNamespace                         string
	_ecsNetworkBindings                   string
	_ecsNetworkConfiguration              string
	_ecsNetworkMode                       string
	_ecsNextToken                         string
	_ecsOverrides                         string
	_ecsPidMode                           string
	_ecsPlacementConstraints              string
	_ecsPlacementStrategy                 string
	_ecsPlatformDevices                   string
	_ecsPlatformVersion                   string
	_ecsPrimaryContainer                  string
	_ecsPrimaryTaskSet                    string
	_ecsPrincipalArn                      string
	_ecsPropagateTags                     string
	_ecsProtectionEnabled                 string
	_ecsProxyConfiguration                string
	_ecsPullStartedAt                     string
	_ecsPullStoppedAt                     string
	_ecsReason                            string
	_ecsReferenceId                       string
	_ecsRequiresCompatibilities           string
	_ecsResourceArn                       string
	_ecsResourceManagementType            string
	_ecsRole                              string
	_ecsRuntimeId                         string
	_ecsRuntimePlatform                   string
	_ecsScale                             string
	_ecsScalingTarget                     string
	_ecsSchedulingStrategy                string
	_ecsService                           string
	_ecsServiceArn                        string
	_ecsServiceConnectConfiguration       string
	_ecsServiceConnectDefaults            string
	_ecsServiceDeploymentArn              string
	_ecsServiceDeploymentArns             []string
	_ecsServiceName                       string
	_ecsServiceRegistries                 string
	_ecsServiceRevisionArns               []string
	_ecsServices                          []string
	_ecsSettings                          string
	_ecsSort                              string
	_ecsStartedBy                         string
	_ecsStatus                            string
	_ecsStopType                          string
	_ecsTagKeys                           []string
	_ecsTags                              string
	_ecsTargetType                        string
	_ecsTask                              string
	_ecsTaskDefinition                    string
	_ecsTaskDefinitions                   []string
	_ecsTaskRoleArn                       string
	_ecsTaskSet                           string
	_ecsTaskSets                          []string
	_ecsTasks                             []string
	_ecsTotalResources                    string
	_ecsValue                             string
	_ecsVersionInfo                       string
	_ecsVolumeConfigurations              string
	_ecsVolumes                           string
	_ecsVpcLatticeConfigurations          string
)

// Creates a capacity provider. Capacity providers are associated with a cluster
// and are used in capacity provider strategies to facilitate cluster auto scaling.
// You can create capacity providers for Amazon ECS Managed Instances and EC2
// instances. Fargate has the predefined FARGATE and FARGATE_SPOT capacity
// providers.
func ecs_CreateCapacityProvider(cfg aws.Config, client *ecs.Client) {
	input := &ecs.CreateCapacityProviderInput{
		// Name: *string, // Required
	}

	if len(_ecsName) > 0 {
		input.Name = aws.String(_ecsName)
	}
	if len(_ecsAutoScalingGroupProvider) > 0 {
		if err := assignInputField(input, "AutoScalingGroupProvider", _ecsAutoScalingGroupProvider); err != nil {
			log.Errorf("invalid --auto-scaling-group-provider: %s", err.Error())
			return
		}
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsManagedInstancesProvider) > 0 {
		if err := assignInputField(input, "ManagedInstancesProvider", _ecsManagedInstancesProvider); err != nil {
			log.Errorf("invalid --managed-instances-provider: %s", err.Error())
			return
		}
	}
	if len(_ecsTags) > 0 {
		if err := assignInputField(input, "Tags", _ecsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCapacityProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Amazon ECS cluster. By default, your account receives a default
// cluster when you launch your first container instance. However, you can create
// your own cluster with a unique name.
//
// When you call the [CreateCluster] API operation, Amazon ECS attempts to create the Amazon ECS
// service-linked role for your account. This is so that it can manage required
// resources in other Amazon Web Services services on your behalf. However, if the
// user that makes the call doesn't have permissions to create the service-linked
// role, it isn't created. For more information, see [Using service-linked roles for Amazon ECS]in the Amazon Elastic
// Container Service Developer Guide.
//
// [Using service-linked roles for Amazon ECS]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using-service-linked-roles.html
// [CreateCluster]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateCluster.html
func ecs_CreateCluster(cfg aws.Config, client *ecs.Client) {
	input := &ecs.CreateClusterInput{}

	if len(_ecsCapacityProviders) > 0 {
		input.CapacityProviders = append([]string(nil), _ecsCapacityProviders...)
	}
	if len(_ecsClusterName) > 0 {
		input.ClusterName = aws.String(_ecsClusterName)
	}
	if len(_ecsConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _ecsConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_ecsDefaultCapacityProviderStrategy) > 0 {
		if err := assignInputField(input, "DefaultCapacityProviderStrategy", _ecsDefaultCapacityProviderStrategy); err != nil {
			log.Errorf("invalid --default-capacity-provider-strategy: %s", err.Error())
			return
		}
	}
	if len(_ecsServiceConnectDefaults) > 0 {
		if err := assignInputField(input, "ServiceConnectDefaults", _ecsServiceConnectDefaults); err != nil {
			log.Errorf("invalid --service-connect-defaults: %s", err.Error())
			return
		}
	}
	if len(_ecsSettings) > 0 {
		if err := assignInputField(input, "Settings", _ecsSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}
	if len(_ecsTags) > 0 {
		if err := assignInputField(input, "Tags", _ecsTags); err != nil {
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

// Creates an Express service that simplifies deploying containerized web
// applications on Amazon ECS with managed Amazon Web Services infrastructure. This
// operation provisions and configures Application Load Balancers, target groups,
// security groups, and auto-scaling policies automatically.
//
// Specify a primary container configuration with your application image and basic
// settings. Amazon ECS creates the necessary Amazon Web Services resources for
// traffic distribution, health monitoring, network access control, and capacity
// management.
//
// Provide an execution role for task operations and an infrastructure role for
// managing Amazon Web Services resources on your behalf.
func ecs_CreateExpressGatewayService(cfg aws.Config, client *ecs.Client) {
	input := &ecs.CreateExpressGatewayServiceInput{
		// ExecutionRoleArn: *string, // Required
		// InfrastructureRoleArn: *string, // Required
		// PrimaryContainer: *types.ExpressGatewayContainer, // Required
	}

	if len(_ecsExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_ecsExecutionRoleArn)
	}
	if len(_ecsInfrastructureRoleArn) > 0 {
		input.InfrastructureRoleArn = aws.String(_ecsInfrastructureRoleArn)
	}
	if len(_ecsPrimaryContainer) > 0 {
		if err := assignInputField(input, "PrimaryContainer", _ecsPrimaryContainer); err != nil {
			log.Errorf("invalid --primary-container: %s", err.Error())
			return
		}
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsCpu) > 0 {
		input.Cpu = aws.String(_ecsCpu)
	}
	if len(_ecsHealthCheckPath) > 0 {
		input.HealthCheckPath = aws.String(_ecsHealthCheckPath)
	}
	if len(_ecsMemory) > 0 {
		input.Memory = aws.String(_ecsMemory)
	}
	if len(_ecsNetworkConfiguration) > 0 {
		if err := assignInputField(input, "NetworkConfiguration", _ecsNetworkConfiguration); err != nil {
			log.Errorf("invalid --network-configuration: %s", err.Error())
			return
		}
	}
	if len(_ecsScalingTarget) > 0 {
		if err := assignInputField(input, "ScalingTarget", _ecsScalingTarget); err != nil {
			log.Errorf("invalid --scaling-target: %s", err.Error())
			return
		}
	}
	if len(_ecsServiceName) > 0 {
		input.ServiceName = aws.String(_ecsServiceName)
	}
	if len(_ecsTags) > 0 {
		if err := assignInputField(input, "Tags", _ecsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_ecsTaskRoleArn) > 0 {
		input.TaskRoleArn = aws.String(_ecsTaskRoleArn)
	}

	if resp, err := client.CreateExpressGatewayService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Runs and maintains your desired number of tasks from a specified task
// definition. If the number of tasks running in a service drops below the
// desiredCount , Amazon ECS runs another copy of the task in the specified
// cluster. To update an existing service, use [UpdateService].
//
// On March 21, 2024, a change was made to resolve the task definition revision
// before authorization. When a task definition revision is not specified,
// authorization will occur using the latest revision of a task definition.
//
// Amazon Elastic Inference (EI) is no longer available to customers.
//
// In addition to maintaining the desired count of tasks in your service, you can
// optionally run your service behind one or more load balancers. The load
// balancers distribute traffic across the tasks that are associated with the
// service. For more information, see [Service load balancing]in the Amazon Elastic Container Service
// Developer Guide.
//
// You can attach Amazon EBS volumes to Amazon ECS tasks by configuring the volume
// when creating or updating a service. volumeConfigurations is only supported for
// REPLICA service and not DAEMON service. For more information, see [Amazon EBS volumes]in the Amazon
// Elastic Container Service Developer Guide.
//
// Tasks for services that don't use a load balancer are considered healthy if
// they're in the RUNNING state. Tasks for services that use a load balancer are
// considered healthy if they're in the RUNNING state and are reported as healthy
// by the load balancer.
//
// There are two service scheduler strategies available:
//
// - REPLICA - The replica scheduling strategy places and maintains your desired
// number of tasks across your cluster. By default, the service scheduler spreads
// tasks across Availability Zones. You can use task placement strategies and
// constraints to customize task placement decisions. For more information, see [Service scheduler concepts]
// in the Amazon Elastic Container Service Developer Guide.
//
// - DAEMON - The daemon scheduling strategy deploys exactly one task on each
// active container instance that meets all of the task placement constraints that
// you specify in your cluster. The service scheduler also evaluates the task
// placement constraints for running tasks. It also stops tasks that don't meet the
// placement constraints. When using this strategy, you don't need to specify a
// desired number of tasks, a task placement strategy, or use Service Auto Scaling
// policies. For more information, see [Amazon ECS services]in the Amazon Elastic Container Service
// Developer Guide.
//
// The deployment controller is the mechanism that determines how tasks are
// deployed for your service. The valid options are:
//
// - ECS
//
// # When you create a service which uses the ECS deployment controller, you can
//
// choose between the following deployment strategies (which you can set in the “
// strategy ” field in “ deploymentConfiguration ”): :
//
// - ROLLING : When you create a service which uses the rolling update ( ROLLING
// ) deployment strategy, the Amazon ECS service scheduler replaces the currently
// running tasks with new tasks. The number of tasks that Amazon ECS adds or
// removes from the service during a rolling update is controlled by the service
// deployment configuration. For more information, see [Deploy Amazon ECS services by replacing tasks]in the Amazon Elastic
// Container Service Developer Guide.
//
// Rolling update deployments are best suited for the following scenarios:
//
// - Gradual service updates: You need to update your service incrementally
// without taking the entire service offline at once.
//
// - Limited resource requirements: You want to avoid the additional resource
// costs of running two complete environments simultaneously (as required by
// blue/green deployments).
//
// - Acceptable deployment time: Your application can tolerate a longer
// deployment process, as rolling updates replace tasks one by one.
//
// - No need for instant roll back: Your service can tolerate a rollback process
// that takes minutes rather than seconds.
//
// - Simple deployment process: You prefer a straightforward deployment approach
// without the complexity of managing multiple environments, target groups, and
// listeners.
//
// - No load balancer requirement: Your service doesn't use or require a load
// balancer, Application Load Balancer, Network Load Balancer, or Service Connect
// (which are required for blue/green deployments).
//
// - Stateful applications: Your application maintains state that makes it
// difficult to run two parallel environments.
//
// - Cost sensitivity: You want to minimize deployment costs by not running
// duplicate environments during deployment.
//
// # Rolling updates are the default deployment strategy for services and provide a
//
// balance between deployment safety and resource efficiency for many common
// application scenarios.
//
// - BLUE_GREEN : A blue/green deployment strategy ( BLUE_GREEN ) is a release
// methodology that reduces downtime and risk by running two identical production
// environments called blue and green. With Amazon ECS blue/green deployments, you
// can validate new service revisions before directing production traffic to them.
// This approach provides a safer way to deploy changes with the ability to quickly
// roll back if needed. For more information, see [Amazon ECS blue/green deployments]in the Amazon Elastic
// Container Service Developer Guide.
//
// Amazon ECS blue/green deployments are best suited for the following scenarios:
//
// - Service validation: When you need to validate new service revisions before
// directing production traffic to them
//
// - Zero downtime: When your service requires zero-downtime deployments
//
// - Instant roll back: When you need the ability to quickly roll back if issues
// are detected
//
// - Load balancer requirement: When your service uses Application Load
// Balancer, Network Load Balancer, or Service Connect
//
// - LINEAR : A linear deployment strategy ( LINEAR ) gradually shifts traffic
// from the current production environment to a new environment in equal percentage
// increments. With Amazon ECS linear deployments, you can control the pace of
// traffic shifting and validate new service revisions with increasing amounts of
// production traffic.
//
// Linear deployments are best suited for the following scenarios:
//
// - Gradual validation: When you want to gradually validate your new service
// version with increasing traffic
//
// - Performance monitoring: When you need time to monitor metrics and
// performance during the deployment
//
// - Risk minimization: When you want to minimize risk by exposing the new
// version to production traffic incrementally
//
// - Load balancer requirement: When your service uses Application Load Balancer
// or Service Connect
//
// - CANARY : A canary deployment strategy ( CANARY ) shifts a small percentage
// of traffic to the new service revision first, then shifts the remaining traffic
// all at once after a specified time period. This allows you to test the new
// version with a subset of users before full deployment.
//
// Canary deployments are best suited for the following scenarios:
//
// - Feature testing: When you want to test new features with a small subset of
// users before full rollout
//
// - Production validation: When you need to validate performance and
// functionality with real production traffic
//
// - Blast radius control: When you want to minimize blast radius if issues are
// discovered in the new version
//
// - Load balancer requirement: When your service uses Application Load Balancer
// or Service Connect
//
// - External
//
// Use a third-party deployment controller.
//
// - Blue/green deployment (powered by CodeDeploy)
//
// # CodeDeploy installs an updated version of the application as a new replacement
//
// task set and reroutes production traffic from the original application task set
// to the replacement task set. The original task set is terminated after a
// successful deployment. Use this deployment controller to verify a new deployment
// of a service before sending production traffic to it.
//
// When creating a service that uses the EXTERNAL deployment controller, you can
// specify only parameters that aren't controlled at the task set level. The only
// required parameter is the service name. You control your services using the [CreateTaskSet].
// For more information, see [Amazon ECS deployment types]in the Amazon Elastic Container Service Developer
// Guide.
//
// When the service scheduler launches new tasks, it determines task placement.
// For information about task placement and task placement strategies, see [Amazon ECS task placement]in the
// Amazon Elastic Container Service Developer Guide
//
// [Amazon ECS task placement]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement.html
// [Service scheduler concepts]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_services.html
// [Amazon ECS deployment types]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html
// [UpdateService]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_UpdateService.html
// [CreateTaskSet]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateTaskSet.html
// [Amazon ECS services]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_services.html
// [Service load balancing]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-load-balancing.html
// [Amazon EBS volumes]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ebs-volumes.html#ebs-volume-types
// [Amazon ECS blue/green deployments]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-blue-green.html
// [Deploy Amazon ECS services by replacing tasks]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-ecs.html
func ecs_CreateService(cfg aws.Config, client *ecs.Client) {
	input := &ecs.CreateServiceInput{
		// ServiceName: *string, // Required
	}

	if len(_ecsServiceName) > 0 {
		input.ServiceName = aws.String(_ecsServiceName)
	}
	if len(_ecsAvailabilityZoneRebalancing) > 0 {
		if err := assignInputField(input, "AvailabilityZoneRebalancing", _ecsAvailabilityZoneRebalancing); err != nil {
			log.Errorf("invalid --availability-zone-rebalancing: %s", err.Error())
			return
		}
	}
	if len(_ecsCapacityProviderStrategy) > 0 {
		if err := assignInputField(input, "CapacityProviderStrategy", _ecsCapacityProviderStrategy); err != nil {
			log.Errorf("invalid --capacity-provider-strategy: %s", err.Error())
			return
		}
	}
	if len(_ecsClientToken) > 0 {
		input.ClientToken = aws.String(_ecsClientToken)
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsDeploymentConfiguration) > 0 {
		if err := assignInputField(input, "DeploymentConfiguration", _ecsDeploymentConfiguration); err != nil {
			log.Errorf("invalid --deployment-configuration: %s", err.Error())
			return
		}
	}
	if len(_ecsDeploymentController) > 0 {
		if err := assignInputField(input, "DeploymentController", _ecsDeploymentController); err != nil {
			log.Errorf("invalid --deployment-controller: %s", err.Error())
			return
		}
	}
	if len(_ecsDesiredCount) > 0 {
		if err := assignInputField(input, "DesiredCount", _ecsDesiredCount); err != nil {
			log.Errorf("invalid --desired-count: %s", err.Error())
			return
		}
	}
	if len(_ecsEnableECSManagedTags) > 0 {
		if err := assignInputField(input, "EnableECSManagedTags", _ecsEnableECSManagedTags); err != nil {
			log.Errorf("invalid --enable-ecs-managed-tags: %s", err.Error())
			return
		}
	}
	if len(_ecsEnableExecuteCommand) > 0 {
		if err := assignInputField(input, "EnableExecuteCommand", _ecsEnableExecuteCommand); err != nil {
			log.Errorf("invalid --enable-execute-command: %s", err.Error())
			return
		}
	}
	if len(_ecsHealthCheckGracePeriodSeconds) > 0 {
		if err := assignInputField(input, "HealthCheckGracePeriodSeconds", _ecsHealthCheckGracePeriodSeconds); err != nil {
			log.Errorf("invalid --health-check-grace-period-seconds: %s", err.Error())
			return
		}
	}
	if len(_ecsLaunchType) > 0 {
		if err := assignInputField(input, "LaunchType", _ecsLaunchType); err != nil {
			log.Errorf("invalid --launch-type: %s", err.Error())
			return
		}
	}
	if len(_ecsLoadBalancers) > 0 {
		if err := assignInputField(input, "LoadBalancers", _ecsLoadBalancers); err != nil {
			log.Errorf("invalid --load-balancers: %s", err.Error())
			return
		}
	}
	if len(_ecsNetworkConfiguration) > 0 {
		if err := assignInputField(input, "NetworkConfiguration", _ecsNetworkConfiguration); err != nil {
			log.Errorf("invalid --network-configuration: %s", err.Error())
			return
		}
	}
	if len(_ecsPlacementConstraints) > 0 {
		if err := assignInputField(input, "PlacementConstraints", _ecsPlacementConstraints); err != nil {
			log.Errorf("invalid --placement-constraints: %s", err.Error())
			return
		}
	}
	if len(_ecsPlacementStrategy) > 0 {
		if err := assignInputField(input, "PlacementStrategy", _ecsPlacementStrategy); err != nil {
			log.Errorf("invalid --placement-strategy: %s", err.Error())
			return
		}
	}
	if len(_ecsPlatformVersion) > 0 {
		input.PlatformVersion = aws.String(_ecsPlatformVersion)
	}
	if len(_ecsPropagateTags) > 0 {
		if err := assignInputField(input, "PropagateTags", _ecsPropagateTags); err != nil {
			log.Errorf("invalid --propagate-tags: %s", err.Error())
			return
		}
	}
	if len(_ecsRole) > 0 {
		input.Role = aws.String(_ecsRole)
	}
	if len(_ecsSchedulingStrategy) > 0 {
		if err := assignInputField(input, "SchedulingStrategy", _ecsSchedulingStrategy); err != nil {
			log.Errorf("invalid --scheduling-strategy: %s", err.Error())
			return
		}
	}
	if len(_ecsServiceConnectConfiguration) > 0 {
		if err := assignInputField(input, "ServiceConnectConfiguration", _ecsServiceConnectConfiguration); err != nil {
			log.Errorf("invalid --service-connect-configuration: %s", err.Error())
			return
		}
	}
	if len(_ecsServiceRegistries) > 0 {
		if err := assignInputField(input, "ServiceRegistries", _ecsServiceRegistries); err != nil {
			log.Errorf("invalid --service-registries: %s", err.Error())
			return
		}
	}
	if len(_ecsTags) > 0 {
		if err := assignInputField(input, "Tags", _ecsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_ecsTaskDefinition) > 0 {
		input.TaskDefinition = aws.String(_ecsTaskDefinition)
	}
	if len(_ecsVolumeConfigurations) > 0 {
		if err := assignInputField(input, "VolumeConfigurations", _ecsVolumeConfigurations); err != nil {
			log.Errorf("invalid --volume-configurations: %s", err.Error())
			return
		}
	}
	if len(_ecsVpcLatticeConfigurations) > 0 {
		if err := assignInputField(input, "VpcLatticeConfigurations", _ecsVpcLatticeConfigurations); err != nil {
			log.Errorf("invalid --vpc-lattice-configurations: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a task set in the specified cluster and service. This is used when a
// service uses the EXTERNAL deployment controller type. For more information, see [Amazon ECS deployment types]
// in the Amazon Elastic Container Service Developer Guide.
//
// On March 21, 2024, a change was made to resolve the task definition revision
// before authorization. When a task definition revision is not specified,
// authorization will occur using the latest revision of a task definition.
//
// For information about the maximum number of task sets and other quotas, see [Amazon ECS service quotas] in
// the Amazon Elastic Container Service Developer Guide.
//
// [Amazon ECS deployment types]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html
// [Amazon ECS service quotas]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-quotas.html
func ecs_CreateTaskSet(cfg aws.Config, client *ecs.Client) {
	input := &ecs.CreateTaskSetInput{
		// Cluster: *string, // Required
		// Service: *string, // Required
		// TaskDefinition: *string, // Required
	}

	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsService) > 0 {
		input.Service = aws.String(_ecsService)
	}
	if len(_ecsTaskDefinition) > 0 {
		input.TaskDefinition = aws.String(_ecsTaskDefinition)
	}
	if len(_ecsCapacityProviderStrategy) > 0 {
		if err := assignInputField(input, "CapacityProviderStrategy", _ecsCapacityProviderStrategy); err != nil {
			log.Errorf("invalid --capacity-provider-strategy: %s", err.Error())
			return
		}
	}
	if len(_ecsClientToken) > 0 {
		input.ClientToken = aws.String(_ecsClientToken)
	}
	if len(_ecsExternalId) > 0 {
		input.ExternalId = aws.String(_ecsExternalId)
	}
	if len(_ecsLaunchType) > 0 {
		if err := assignInputField(input, "LaunchType", _ecsLaunchType); err != nil {
			log.Errorf("invalid --launch-type: %s", err.Error())
			return
		}
	}
	if len(_ecsLoadBalancers) > 0 {
		if err := assignInputField(input, "LoadBalancers", _ecsLoadBalancers); err != nil {
			log.Errorf("invalid --load-balancers: %s", err.Error())
			return
		}
	}
	if len(_ecsNetworkConfiguration) > 0 {
		if err := assignInputField(input, "NetworkConfiguration", _ecsNetworkConfiguration); err != nil {
			log.Errorf("invalid --network-configuration: %s", err.Error())
			return
		}
	}
	if len(_ecsPlatformVersion) > 0 {
		input.PlatformVersion = aws.String(_ecsPlatformVersion)
	}
	if len(_ecsScale) > 0 {
		if err := assignInputField(input, "Scale", _ecsScale); err != nil {
			log.Errorf("invalid --scale: %s", err.Error())
			return
		}
	}
	if len(_ecsServiceRegistries) > 0 {
		if err := assignInputField(input, "ServiceRegistries", _ecsServiceRegistries); err != nil {
			log.Errorf("invalid --service-registries: %s", err.Error())
			return
		}
	}
	if len(_ecsTags) > 0 {
		if err := assignInputField(input, "Tags", _ecsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTaskSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables an account setting for a specified user, role, or the root user for an
// account.
func ecs_DeleteAccountSetting(cfg aws.Config, client *ecs.Client) {
	input := &ecs.DeleteAccountSettingInput{
		// Name: types.SettingName, // Required
	}

	if len(_ecsName) > 0 {
		if err := assignInputField(input, "Name", _ecsName); err != nil {
			log.Errorf("invalid --name: %s", err.Error())
			return
		}
	}
	if len(_ecsPrincipalArn) > 0 {
		input.PrincipalArn = aws.String(_ecsPrincipalArn)
	}

	if resp, err := client.DeleteAccountSetting(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes one or more custom attributes from an Amazon ECS resource.
func ecs_DeleteAttributes(cfg aws.Config, client *ecs.Client) {
	input := &ecs.DeleteAttributesInput{
		// Attributes: []types.Attribute, // Required
	}

	if len(_ecsAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _ecsAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}

	if resp, err := client.DeleteAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified capacity provider.
// The FARGATE and FARGATE_SPOT capacity providers are reserved and can't be
// deleted. You can disassociate them from a cluster using either [PutClusterCapacityProviders]or by deleting
// the cluster.
//
// Prior to a capacity provider being deleted, the capacity provider must be
// removed from the capacity provider strategy from all services. The [UpdateService]API can be
// used to remove a capacity provider from a service's capacity provider strategy.
// When updating a service, the forceNewDeployment option can be used to ensure
// that any tasks using the Amazon EC2 instance capacity provided by the capacity
// provider are transitioned to use the capacity from the remaining capacity
// providers. Only capacity providers that aren't associated with a cluster can be
// deleted. To remove a capacity provider from a cluster, you can either use [PutClusterCapacityProviders]or
// delete the cluster.
//
// [PutClusterCapacityProviders]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutClusterCapacityProviders.html
// [UpdateService]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_UpdateService.html
func ecs_DeleteCapacityProvider(cfg aws.Config, client *ecs.Client) {
	input := &ecs.DeleteCapacityProviderInput{
		// CapacityProvider: *string, // Required
	}

	if len(_ecsCapacityProvider) > 0 {
		input.CapacityProvider = aws.String(_ecsCapacityProvider)
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}

	if resp, err := client.DeleteCapacityProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified cluster. The cluster transitions to the INACTIVE state.
// Clusters with an INACTIVE status might remain discoverable in your account for
// a period of time. However, this behavior is subject to change in the future. We
// don't recommend that you rely on INACTIVE clusters persisting.
//
// You must deregister all container instances from this cluster before you may
// delete it. You can list the container instances in a cluster with [ListContainerInstances]and
// deregister them with [DeregisterContainerInstance].
//
// [ListContainerInstances]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ListContainerInstances.html
// [DeregisterContainerInstance]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DeregisterContainerInstance.html
func ecs_DeleteCluster(cfg aws.Config, client *ecs.Client) {
	input := &ecs.DeleteClusterInput{
		// Cluster: *string, // Required
	}

	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}

	if resp, err := client.DeleteCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Express service and removes all associated Amazon Web Services
// resources. This operation stops service tasks, removes the Application Load
// Balancer, target groups, security groups, auto-scaling policies, and other
// managed infrastructure components.
//
// The service enters a DRAINING state where existing tasks complete current
// requests without starting new tasks. After all tasks stop, the service and
// infrastructure are permanently removed.
//
// This operation cannot be reversed. Back up important data and verify the
// service is no longer needed before deletion.
func ecs_DeleteExpressGatewayService(cfg aws.Config, client *ecs.Client) {
	input := &ecs.DeleteExpressGatewayServiceInput{
		// ServiceArn: *string, // Required
	}

	if len(_ecsServiceArn) > 0 {
		input.ServiceArn = aws.String(_ecsServiceArn)
	}

	if resp, err := client.DeleteExpressGatewayService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified service within a cluster. You can delete a service if you
// have no running tasks in it and the desired task count is zero. If the service
// is actively maintaining tasks, you can't delete it, and you must update the
// service to a desired task count of zero. For more information, see [UpdateService].
//
// When you delete a service, if there are still running tasks that require
// cleanup, the service status moves from ACTIVE to DRAINING , and the service is
// no longer visible in the console or in the [ListServices]API operation. After all tasks have
// transitioned to either STOPPING or STOPPED status, the service status moves
// from DRAINING to INACTIVE . Services in the DRAINING or INACTIVE status can
// still be viewed with the [DescribeServices]API operation. However, in the future, INACTIVE
// services may be cleaned up and purged from Amazon ECS record keeping, and [DescribeServices]calls
// on those services return a ServiceNotFoundException error.
//
// If you attempt to create a new service with the same name as an existing
// service in either ACTIVE or DRAINING status, you receive an error.
//
// [UpdateService]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_UpdateService.html
// [ListServices]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ListServices.html
// [DescribeServices]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DescribeServices.html
func ecs_DeleteService(cfg aws.Config, client *ecs.Client) {
	input := &ecs.DeleteServiceInput{
		// Service: *string, // Required
	}

	if len(_ecsService) > 0 {
		input.Service = aws.String(_ecsService)
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsForce) > 0 {
		if err := assignInputField(input, "Force", _ecsForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes one or more task definitions.
// You must deregister a task definition revision before you delete it. For more
// information, see [DeregisterTaskDefinition].
//
// When you delete a task definition revision, it is immediately transitions from
// the INACTIVE to DELETE_IN_PROGRESS . Existing tasks and services that reference
// a DELETE_IN_PROGRESS task definition revision continue to run without
// disruption. Existing services that reference a DELETE_IN_PROGRESS task
// definition revision can still scale up or down by modifying the service's
// desired count.
//
// You can't use a DELETE_IN_PROGRESS task definition revision to run new tasks or
// create new services. You also can't update an existing service to reference a
// DELETE_IN_PROGRESS task definition revision.
//
// A task definition revision will stay in DELETE_IN_PROGRESS status until all the
// associated tasks and services have been terminated.
//
// When you delete all INACTIVE task definition revisions, the task definition
// name is not displayed in the console and not returned in the API. If a task
// definition revisions are in the DELETE_IN_PROGRESS state, the task definition
// name is displayed in the console and returned in the API. The task definition
// name is retained by Amazon ECS and the revision is incremented the next time you
// create a task definition with that name.
//
// [DeregisterTaskDefinition]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DeregisterTaskDefinition.html
func ecs_DeleteTaskDefinitions(cfg aws.Config, client *ecs.Client) {
	input := &ecs.DeleteTaskDefinitionsInput{
		// TaskDefinitions: []string, // Required
	}

	if len(_ecsTaskDefinitions) > 0 {
		input.TaskDefinitions = append([]string(nil), _ecsTaskDefinitions...)
	}

	if resp, err := client.DeleteTaskDefinitions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified task set within a service. This is used when a service uses
// the EXTERNAL deployment controller type. For more information, see [Amazon ECS deployment types] in the
// Amazon Elastic Container Service Developer Guide.
//
// [Amazon ECS deployment types]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html
func ecs_DeleteTaskSet(cfg aws.Config, client *ecs.Client) {
	input := &ecs.DeleteTaskSetInput{
		// Cluster: *string, // Required
		// Service: *string, // Required
		// TaskSet: *string, // Required
	}

	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsService) > 0 {
		input.Service = aws.String(_ecsService)
	}
	if len(_ecsTaskSet) > 0 {
		input.TaskSet = aws.String(_ecsTaskSet)
	}
	if len(_ecsForce) > 0 {
		if err := assignInputField(input, "Force", _ecsForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteTaskSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregisters an Amazon ECS container instance from the specified cluster. This
// instance is no longer available to run tasks.
//
// If you intend to use the container instance for some other purpose after
// deregistration, we recommend that you stop all of the tasks running on the
// container instance before deregistration. That prevents any orphaned tasks from
// consuming resources.
//
// Deregistering a container instance removes the instance from a cluster, but it
// doesn't terminate the EC2 instance. If you are finished using the instance, be
// sure to terminate it in the Amazon EC2 console to stop billing.
//
// If you terminate a running container instance, Amazon ECS automatically
// deregisters the instance from your cluster (stopped container instances or
// instances with disconnected agents aren't automatically deregistered when
// terminated).
func ecs_DeregisterContainerInstance(cfg aws.Config, client *ecs.Client) {
	input := &ecs.DeregisterContainerInstanceInput{
		// ContainerInstance: *string, // Required
	}

	if len(_ecsContainerInstance) > 0 {
		input.ContainerInstance = aws.String(_ecsContainerInstance)
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsForce) > 0 {
		if err := assignInputField(input, "Force", _ecsForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeregisterContainerInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregisters the specified task definition by family and revision. Upon
// deregistration, the task definition is marked as INACTIVE . Existing tasks and
// services that reference an INACTIVE task definition continue to run without
// disruption. Existing services that reference an INACTIVE task definition can
// still scale up or down by modifying the service's desired count. If you want to
// delete a task definition revision, you must first deregister the task definition
// revision.
//
// You can't use an INACTIVE task definition to run new tasks or create new
// services, and you can't update an existing service to reference an INACTIVE
// task definition. However, there may be up to a 10-minute window following
// deregistration where these restrictions have not yet taken effect.
//
// At this time, INACTIVE task definitions remain discoverable in your account
// indefinitely. However, this behavior is subject to change in the future. We
// don't recommend that you rely on INACTIVE task definitions persisting beyond
// the lifecycle of any associated tasks and services.
//
// You must deregister a task definition revision before you delete it. For more
// information, see [DeleteTaskDefinitions].
//
// [DeleteTaskDefinitions]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DeleteTaskDefinitions.html
func ecs_DeregisterTaskDefinition(cfg aws.Config, client *ecs.Client) {
	input := &ecs.DeregisterTaskDefinitionInput{
		// TaskDefinition: *string, // Required
	}

	if len(_ecsTaskDefinition) > 0 {
		input.TaskDefinition = aws.String(_ecsTaskDefinition)
	}

	if resp, err := client.DeregisterTaskDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes one or more of your capacity providers.
func ecs_DescribeCapacityProviders(cfg aws.Config, client *ecs.Client) {
	input := &ecs.DescribeCapacityProvidersInput{}

	if len(_ecsCapacityProviders) > 0 {
		input.CapacityProviders = append([]string(nil), _ecsCapacityProviders...)
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsInclude) > 0 {
		if err := assignInputField(input, "Include", _ecsInclude); err != nil {
			log.Errorf("invalid --include: %s", err.Error())
			return
		}
	}
	if len(_ecsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ecsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ecsNextToken) > 0 {
		input.NextToken = aws.String(_ecsNextToken)
	}

	if resp, err := client.DescribeCapacityProviders(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes one or more of your clusters.
// For CLI examples, see [describe-clusters.rst] on GitHub.
//
// [describe-clusters.rst]: https://github.com/aws/aws-cli/blob/develop/awscli/examples/ecs/describe-clusters.rst
func ecs_DescribeClusters(cfg aws.Config, client *ecs.Client) {
	input := &ecs.DescribeClustersInput{}

	if len(_ecsClusters) > 0 {
		input.Clusters = append([]string(nil), _ecsClusters...)
	}
	if len(_ecsInclude) > 0 {
		if err := assignInputField(input, "Include", _ecsInclude); err != nil {
			log.Errorf("invalid --include: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeClusters(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes one or more container instances. Returns metadata about each
// container instance requested.
func ecs_DescribeContainerInstances(cfg aws.Config, client *ecs.Client) {
	input := &ecs.DescribeContainerInstancesInput{
		// ContainerInstances: []string, // Required
	}

	if len(_ecsContainerInstances) > 0 {
		input.ContainerInstances = append([]string(nil), _ecsContainerInstances...)
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsInclude) > 0 {
		if err := assignInputField(input, "Include", _ecsInclude); err != nil {
			log.Errorf("invalid --include: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeContainerInstances(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about an Express service, including current
// status, configuration, managed infrastructure, and service revisions.
//
// Returns comprehensive service details, active service revisions, ingress paths
// with endpoints, and managed Amazon Web Services resource status including load
// balancers and auto-scaling policies.
//
// Use the include parameter to retrieve additional information such as resource
// tags.
func ecs_DescribeExpressGatewayService(cfg aws.Config, client *ecs.Client) {
	input := &ecs.DescribeExpressGatewayServiceInput{
		// ServiceArn: *string, // Required
	}

	if len(_ecsServiceArn) > 0 {
		input.ServiceArn = aws.String(_ecsServiceArn)
	}
	if len(_ecsInclude) > 0 {
		if err := assignInputField(input, "Include", _ecsInclude); err != nil {
			log.Errorf("invalid --include: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeExpressGatewayService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes one or more of your service deployments.
// A service deployment happens when you release a software update for the
// service. For more information, see [View service history using Amazon ECS service deployments].
//
// [View service history using Amazon ECS service deployments]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-deployment.html
func ecs_DescribeServiceDeployments(cfg aws.Config, client *ecs.Client) {
	input := &ecs.DescribeServiceDeploymentsInput{
		// ServiceDeploymentArns: []string, // Required
	}

	if len(_ecsServiceDeploymentArns) > 0 {
		input.ServiceDeploymentArns = append([]string(nil), _ecsServiceDeploymentArns...)
	}

	if resp, err := client.DescribeServiceDeployments(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes one or more service revisions.
// A service revision is a version of the service that includes the values for the
// Amazon ECS resources (for example, task definition) and the environment
// resources (for example, load balancers, subnets, and security groups). For more
// information, see [Amazon ECS service revisions].
//
// You can't describe a service revision that was created before October 25, 2024.
//
// [Amazon ECS service revisions]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-revision.html
func ecs_DescribeServiceRevisions(cfg aws.Config, client *ecs.Client) {
	input := &ecs.DescribeServiceRevisionsInput{
		// ServiceRevisionArns: []string, // Required
	}

	if len(_ecsServiceRevisionArns) > 0 {
		input.ServiceRevisionArns = append([]string(nil), _ecsServiceRevisionArns...)
	}

	if resp, err := client.DescribeServiceRevisions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified services running in your cluster.
func ecs_DescribeServices(cfg aws.Config, client *ecs.Client) {
	input := &ecs.DescribeServicesInput{
		// Services: []string, // Required
	}

	if len(_ecsServices) > 0 {
		input.Services = append([]string(nil), _ecsServices...)
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsInclude) > 0 {
		if err := assignInputField(input, "Include", _ecsInclude); err != nil {
			log.Errorf("invalid --include: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeServices(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a task definition. You can specify a family and revision to find
// information about a specific task definition, or you can simply specify the
// family to find the latest ACTIVE revision in that family.
//
// You can only describe INACTIVE task definitions while an active task or service
// references them.
func ecs_DescribeTaskDefinition(cfg aws.Config, client *ecs.Client) {
	input := &ecs.DescribeTaskDefinitionInput{
		// TaskDefinition: *string, // Required
	}

	if len(_ecsTaskDefinition) > 0 {
		input.TaskDefinition = aws.String(_ecsTaskDefinition)
	}
	if len(_ecsInclude) > 0 {
		if err := assignInputField(input, "Include", _ecsInclude); err != nil {
			log.Errorf("invalid --include: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeTaskDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the task sets in the specified cluster and service. This is used when
// a service uses the EXTERNAL deployment controller type. For more information,
// see [Amazon ECS Deployment Types]in the Amazon Elastic Container Service Developer Guide.
//
// [Amazon ECS Deployment Types]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html
func ecs_DescribeTaskSets(cfg aws.Config, client *ecs.Client) {
	input := &ecs.DescribeTaskSetsInput{
		// Cluster: *string, // Required
		// Service: *string, // Required
	}

	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsService) > 0 {
		input.Service = aws.String(_ecsService)
	}
	if len(_ecsInclude) > 0 {
		if err := assignInputField(input, "Include", _ecsInclude); err != nil {
			log.Errorf("invalid --include: %s", err.Error())
			return
		}
	}
	if len(_ecsTaskSets) > 0 {
		input.TaskSets = append([]string(nil), _ecsTaskSets...)
	}

	if resp, err := client.DescribeTaskSets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a specified task or tasks.
// Currently, stopped tasks appear in the returned results for at least one hour.
//
// If you have tasks with tags, and then delete the cluster, the tagged tasks are
// returned in the response. If you create a new cluster with the same name as the
// deleted cluster, the tagged tasks are not included in the response.
func ecs_DescribeTasks(cfg aws.Config, client *ecs.Client) {
	input := &ecs.DescribeTasksInput{
		// Tasks: []string, // Required
	}

	if len(_ecsTasks) > 0 {
		input.Tasks = append([]string(nil), _ecsTasks...)
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsInclude) > 0 {
		if err := assignInputField(input, "Include", _ecsInclude); err != nil {
			log.Errorf("invalid --include: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeTasks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action is only used by the Amazon ECS agent, and it is not intended for
// use outside of the agent.
//
// Returns an endpoint for the Amazon ECS agent to poll for updates.
func ecs_DiscoverPollEndpoint(cfg aws.Config, client *ecs.Client) {
	input := &ecs.DiscoverPollEndpointInput{}

	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsContainerInstance) > 0 {
		input.ContainerInstance = aws.String(_ecsContainerInstance)
	}

	if resp, err := client.DiscoverPollEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Runs a command remotely on a container within a task.
// If you use a condition key in your IAM policy to refine the conditions for the
// policy statement, for example limit the actions to a specific cluster, you
// receive an AccessDeniedException when there is a mismatch between the condition
// key value and the corresponding parameter value.
//
// For information about required permissions and considerations, see [Using Amazon ECS Exec for debugging] in the
// Amazon ECS Developer Guide.
//
// [Using Amazon ECS Exec for debugging]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-exec.html
func ecs_ExecuteCommand(cfg aws.Config, client *ecs.Client) {
	input := &ecs.ExecuteCommandInput{
		// Command: *string, // Required
		// Interactive: bool, // Required
		// Task: *string, // Required
	}

	if len(_ecsCommand) > 0 {
		input.Command = aws.String(_ecsCommand)
	}
	if len(_ecsInteractive) > 0 {
		if err := assignInputField(input, "Interactive", _ecsInteractive); err != nil {
			log.Errorf("invalid --interactive: %s", err.Error())
			return
		}
	}
	if len(_ecsTask) > 0 {
		input.Task = aws.String(_ecsTask)
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsContainer) > 0 {
		input.Container = aws.String(_ecsContainer)
	}

	if resp, err := client.ExecuteCommand(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the protection status of tasks in an Amazon ECS service.
func ecs_GetTaskProtection(cfg aws.Config, client *ecs.Client) {
	input := &ecs.GetTaskProtectionInput{
		// Cluster: *string, // Required
	}

	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsTasks) > 0 {
		input.Tasks = append([]string(nil), _ecsTasks...)
	}

	if resp, err := client.GetTaskProtection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the account settings for a specified principal.
func ecs_ListAccountSettings(cfg aws.Config, client *ecs.Client) {
	input := &ecs.ListAccountSettingsInput{}

	if len(_ecsEffectiveSettings) > 0 {
		if err := assignInputField(input, "EffectiveSettings", _ecsEffectiveSettings); err != nil {
			log.Errorf("invalid --effective-settings: %s", err.Error())
			return
		}
	}
	if len(_ecsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ecsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ecsName) > 0 {
		if err := assignInputField(input, "Name", _ecsName); err != nil {
			log.Errorf("invalid --name: %s", err.Error())
			return
		}
	}
	if len(_ecsNextToken) > 0 {
		input.NextToken = aws.String(_ecsNextToken)
	}
	if len(_ecsPrincipalArn) > 0 {
		input.PrincipalArn = aws.String(_ecsPrincipalArn)
	}
	if len(_ecsValue) > 0 {
		input.Value = aws.String(_ecsValue)
	}

	if disablePaginator() {
		if resp, err := client.ListAccountSettings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ecs.ListAccountSettingsOutput
	p := ecs.NewListAccountSettingsPaginator(client, input)
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

// Lists the attributes for Amazon ECS resources within a specified target type
// and cluster. When you specify a target type and cluster, ListAttributes returns
// a list of attribute objects, one for each attribute on each resource. You can
// filter the list of results to a single attribute name to only return results
// that have that name. You can also filter the results by attribute name and
// value. You can do this, for example, to see which container instances in a
// cluster are running a Linux AMI ( ecs.os-type=linux ).
func ecs_ListAttributes(cfg aws.Config, client *ecs.Client) {
	input := &ecs.ListAttributesInput{
		// TargetType: types.TargetType, // Required
	}

	if len(_ecsTargetType) > 0 {
		if err := assignInputField(input, "TargetType", _ecsTargetType); err != nil {
			log.Errorf("invalid --target-type: %s", err.Error())
			return
		}
	}
	if len(_ecsAttributeName) > 0 {
		input.AttributeName = aws.String(_ecsAttributeName)
	}
	if len(_ecsAttributeValue) > 0 {
		input.AttributeValue = aws.String(_ecsAttributeValue)
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ecsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ecsNextToken) > 0 {
		input.NextToken = aws.String(_ecsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAttributes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ecs.ListAttributesOutput
	p := ecs.NewListAttributesPaginator(client, input)
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

// Returns a list of existing clusters.
func ecs_ListClusters(cfg aws.Config, client *ecs.Client) {
	input := &ecs.ListClustersInput{}

	if len(_ecsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ecsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ecsNextToken) > 0 {
		input.NextToken = aws.String(_ecsNextToken)
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

	var results []*ecs.ListClustersOutput
	p := ecs.NewListClustersPaginator(client, input)
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

// Returns a list of container instances in a specified cluster. You can filter
// the results of a ListContainerInstances operation with cluster query language
// statements inside the filter parameter. For more information, see [Cluster Query Language] in the
// Amazon Elastic Container Service Developer Guide.
//
// [Cluster Query Language]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html
func ecs_ListContainerInstances(cfg aws.Config, client *ecs.Client) {
	input := &ecs.ListContainerInstancesInput{}

	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsFilter) > 0 {
		input.Filter = aws.String(_ecsFilter)
	}
	if len(_ecsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ecsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ecsNextToken) > 0 {
		input.NextToken = aws.String(_ecsNextToken)
	}
	if len(_ecsStatus) > 0 {
		if err := assignInputField(input, "Status", _ecsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListContainerInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ecs.ListContainerInstancesOutput
	p := ecs.NewListContainerInstancesPaginator(client, input)
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

// This operation lists all the service deployments that meet the specified filter
// criteria.
//
// A service deployment happens when you release a software update for the
// service. You route traffic from the running service revisions to the new service
// revison and control the number of running tasks.
//
// This API returns the values that you use for the request parameters in [DescribeServiceRevisions].
//
// [DescribeServiceRevisions]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DescribeServiceRevisions.html
func ecs_ListServiceDeployments(cfg aws.Config, client *ecs.Client) {
	input := &ecs.ListServiceDeploymentsInput{
		// Service: *string, // Required
	}

	if len(_ecsService) > 0 {
		input.Service = aws.String(_ecsService)
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsCreatedAt) > 0 {
		if err := assignInputField(input, "CreatedAt", _ecsCreatedAt); err != nil {
			log.Errorf("invalid --created-at: %s", err.Error())
			return
		}
	}
	if len(_ecsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ecsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ecsNextToken) > 0 {
		input.NextToken = aws.String(_ecsNextToken)
	}
	if len(_ecsStatus) > 0 {
		if err := assignInputField(input, "Status", _ecsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListServiceDeployments(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of services. You can filter the results by cluster, launch type,
// and scheduling strategy.
func ecs_ListServices(cfg aws.Config, client *ecs.Client) {
	input := &ecs.ListServicesInput{}

	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsLaunchType) > 0 {
		if err := assignInputField(input, "LaunchType", _ecsLaunchType); err != nil {
			log.Errorf("invalid --launch-type: %s", err.Error())
			return
		}
	}
	if len(_ecsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ecsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ecsNextToken) > 0 {
		input.NextToken = aws.String(_ecsNextToken)
	}
	if len(_ecsResourceManagementType) > 0 {
		if err := assignInputField(input, "ResourceManagementType", _ecsResourceManagementType); err != nil {
			log.Errorf("invalid --resource-management-type: %s", err.Error())
			return
		}
	}
	if len(_ecsSchedulingStrategy) > 0 {
		if err := assignInputField(input, "SchedulingStrategy", _ecsSchedulingStrategy); err != nil {
			log.Errorf("invalid --scheduling-strategy: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListServices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ecs.ListServicesOutput
	p := ecs.NewListServicesPaginator(client, input)
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

// This operation lists all of the services that are associated with a Cloud Map
// namespace. This list might include services in different clusters. In contrast,
// ListServices can only list services in one cluster at a time. If you need to
// filter the list of services in a single cluster by various parameters, use
// ListServices . For more information, see [Service Connect] in the Amazon Elastic Container
// Service Developer Guide.
//
// [Service Connect]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html
func ecs_ListServicesByNamespace(cfg aws.Config, client *ecs.Client) {
	input := &ecs.ListServicesByNamespaceInput{
		// Namespace: *string, // Required
	}

	if len(_ecsNamespace) > 0 {
		input.Namespace = aws.String(_ecsNamespace)
	}
	if len(_ecsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ecsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ecsNextToken) > 0 {
		input.NextToken = aws.String(_ecsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServicesByNamespace(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ecs.ListServicesByNamespaceOutput
	p := ecs.NewListServicesByNamespacePaginator(client, input)
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

// List the tags for an Amazon ECS resource.
func ecs_ListTagsForResource(cfg aws.Config, client *ecs.Client) {
	input := &ecs.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_ecsResourceArn) > 0 {
		input.ResourceArn = aws.String(_ecsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of task definition families that are registered to your account.
// This list includes task definition families that no longer have any ACTIVE task
// definition revisions.
//
// You can filter out task definition families that don't contain any ACTIVE task
// definition revisions by setting the status parameter to ACTIVE . You can also
// filter the results with the familyPrefix parameter.
func ecs_ListTaskDefinitionFamilies(cfg aws.Config, client *ecs.Client) {
	input := &ecs.ListTaskDefinitionFamiliesInput{}

	if len(_ecsFamilyPrefix) > 0 {
		input.FamilyPrefix = aws.String(_ecsFamilyPrefix)
	}
	if len(_ecsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ecsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ecsNextToken) > 0 {
		input.NextToken = aws.String(_ecsNextToken)
	}
	if len(_ecsStatus) > 0 {
		if err := assignInputField(input, "Status", _ecsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTaskDefinitionFamilies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ecs.ListTaskDefinitionFamiliesOutput
	p := ecs.NewListTaskDefinitionFamiliesPaginator(client, input)
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

// Returns a list of task definitions that are registered to your account. You can
// filter the results by family name with the familyPrefix parameter or by status
// with the status parameter.
func ecs_ListTaskDefinitions(cfg aws.Config, client *ecs.Client) {
	input := &ecs.ListTaskDefinitionsInput{}

	if len(_ecsFamilyPrefix) > 0 {
		input.FamilyPrefix = aws.String(_ecsFamilyPrefix)
	}
	if len(_ecsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ecsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ecsNextToken) > 0 {
		input.NextToken = aws.String(_ecsNextToken)
	}
	if len(_ecsSort) > 0 {
		if err := assignInputField(input, "Sort", _ecsSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}
	if len(_ecsStatus) > 0 {
		if err := assignInputField(input, "Status", _ecsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTaskDefinitions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ecs.ListTaskDefinitionsOutput
	p := ecs.NewListTaskDefinitionsPaginator(client, input)
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

// Returns a list of tasks. You can filter the results by cluster, task definition
// family, container instance, launch type, what IAM principal started the task, or
// by the desired status of the task.
//
// Recently stopped tasks might appear in the returned results.
func ecs_ListTasks(cfg aws.Config, client *ecs.Client) {
	input := &ecs.ListTasksInput{}

	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsContainerInstance) > 0 {
		input.ContainerInstance = aws.String(_ecsContainerInstance)
	}
	if len(_ecsDesiredStatus) > 0 {
		if err := assignInputField(input, "DesiredStatus", _ecsDesiredStatus); err != nil {
			log.Errorf("invalid --desired-status: %s", err.Error())
			return
		}
	}
	if len(_ecsFamily) > 0 {
		input.Family = aws.String(_ecsFamily)
	}
	if len(_ecsLaunchType) > 0 {
		if err := assignInputField(input, "LaunchType", _ecsLaunchType); err != nil {
			log.Errorf("invalid --launch-type: %s", err.Error())
			return
		}
	}
	if len(_ecsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ecsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ecsNextToken) > 0 {
		input.NextToken = aws.String(_ecsNextToken)
	}
	if len(_ecsServiceName) > 0 {
		input.ServiceName = aws.String(_ecsServiceName)
	}
	if len(_ecsStartedBy) > 0 {
		input.StartedBy = aws.String(_ecsStartedBy)
	}

	if disablePaginator() {
		if resp, err := client.ListTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ecs.ListTasksOutput
	p := ecs.NewListTasksPaginator(client, input)
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

// Modifies an account setting. Account settings are set on a per-Region basis.
// If you change the root user account setting, the default settings are reset for
// users and roles that do not have specified individual account settings. For more
// information, see [Account Settings]in the Amazon Elastic Container Service Developer Guide.
//
// [Account Settings]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html
func ecs_PutAccountSetting(cfg aws.Config, client *ecs.Client) {
	input := &ecs.PutAccountSettingInput{
		// Name: types.SettingName, // Required
		// Value: *string, // Required
	}

	if len(_ecsName) > 0 {
		if err := assignInputField(input, "Name", _ecsName); err != nil {
			log.Errorf("invalid --name: %s", err.Error())
			return
		}
	}
	if len(_ecsValue) > 0 {
		input.Value = aws.String(_ecsValue)
	}
	if len(_ecsPrincipalArn) > 0 {
		input.PrincipalArn = aws.String(_ecsPrincipalArn)
	}

	if resp, err := client.PutAccountSetting(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an account setting for all users on an account for whom no individual
// account setting has been specified. Account settings are set on a per-Region
// basis.
func ecs_PutAccountSettingDefault(cfg aws.Config, client *ecs.Client) {
	input := &ecs.PutAccountSettingDefaultInput{
		// Name: types.SettingName, // Required
		// Value: *string, // Required
	}

	if len(_ecsName) > 0 {
		if err := assignInputField(input, "Name", _ecsName); err != nil {
			log.Errorf("invalid --name: %s", err.Error())
			return
		}
	}
	if len(_ecsValue) > 0 {
		input.Value = aws.String(_ecsValue)
	}

	if resp, err := client.PutAccountSettingDefault(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create or update an attribute on an Amazon ECS resource. If the attribute
// doesn't exist, it's created. If the attribute exists, its value is replaced with
// the specified value. To delete an attribute, use [DeleteAttributes]. For more information, see [Attributes]
// in the Amazon Elastic Container Service Developer Guide.
//
// [Attributes]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-constraints.html#attributes
// [DeleteAttributes]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DeleteAttributes.html
func ecs_PutAttributes(cfg aws.Config, client *ecs.Client) {
	input := &ecs.PutAttributesInput{
		// Attributes: []types.Attribute, // Required
	}

	if len(_ecsAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _ecsAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}

	if resp, err := client.PutAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the available capacity providers and the default capacity provider
// strategy for a cluster.
//
// You must specify both the available capacity providers and a default capacity
// provider strategy for the cluster. If the specified cluster has existing
// capacity providers associated with it, you must specify all existing capacity
// providers in addition to any new ones you want to add. Any existing capacity
// providers that are associated with a cluster that are omitted from a [PutClusterCapacityProviders]API call
// will be disassociated with the cluster. You can only disassociate an existing
// capacity provider from a cluster if it's not being used by any existing tasks.
//
// When creating a service or running a task on a cluster, if no capacity provider
// or launch type is specified, then the cluster's default capacity provider
// strategy is used. We recommend that you define a default capacity provider
// strategy for your cluster. However, you must specify an empty array ( [] ) to
// bypass defining a default strategy.
//
// Amazon ECS Managed Instances doesn't support this, because when you create a
// capacity provider with Amazon ECS Managed Instances, it becomes available only
// within the specified cluster.
//
// [PutClusterCapacityProviders]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutClusterCapacityProviders.html
func ecs_PutClusterCapacityProviders(cfg aws.Config, client *ecs.Client) {
	input := &ecs.PutClusterCapacityProvidersInput{
		// CapacityProviders: []string, // Required
		// Cluster: *string, // Required
		// DefaultCapacityProviderStrategy: []types.CapacityProviderStrategyItem, // Required
	}

	if len(_ecsCapacityProviders) > 0 {
		input.CapacityProviders = append([]string(nil), _ecsCapacityProviders...)
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsDefaultCapacityProviderStrategy) > 0 {
		if err := assignInputField(input, "DefaultCapacityProviderStrategy", _ecsDefaultCapacityProviderStrategy); err != nil {
			log.Errorf("invalid --default-capacity-provider-strategy: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutClusterCapacityProviders(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action is only used by the Amazon ECS agent, and it is not intended for
// use outside of the agent.
//
// Registers an EC2 instance into the specified cluster. This instance becomes
// available to place containers on.
func ecs_RegisterContainerInstance(cfg aws.Config, client *ecs.Client) {
	input := &ecs.RegisterContainerInstanceInput{}

	if len(_ecsAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _ecsAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsContainerInstanceArn) > 0 {
		input.ContainerInstanceArn = aws.String(_ecsContainerInstanceArn)
	}
	if len(_ecsInstanceIdentityDocument) > 0 {
		input.InstanceIdentityDocument = aws.String(_ecsInstanceIdentityDocument)
	}
	if len(_ecsInstanceIdentityDocumentSignature) > 0 {
		input.InstanceIdentityDocumentSignature = aws.String(_ecsInstanceIdentityDocumentSignature)
	}
	if len(_ecsPlatformDevices) > 0 {
		if err := assignInputField(input, "PlatformDevices", _ecsPlatformDevices); err != nil {
			log.Errorf("invalid --platform-devices: %s", err.Error())
			return
		}
	}
	if len(_ecsTags) > 0 {
		if err := assignInputField(input, "Tags", _ecsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_ecsTotalResources) > 0 {
		if err := assignInputField(input, "TotalResources", _ecsTotalResources); err != nil {
			log.Errorf("invalid --total-resources: %s", err.Error())
			return
		}
	}
	if len(_ecsVersionInfo) > 0 {
		if err := assignInputField(input, "VersionInfo", _ecsVersionInfo); err != nil {
			log.Errorf("invalid --version-info: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterContainerInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers a new task definition from the supplied family and
// containerDefinitions . Optionally, you can add data volumes to your containers
// with the volumes parameter. For more information about task definition
// parameters and defaults, see [Amazon ECS Task Definitions]in the Amazon Elastic Container Service Developer
// Guide.
//
// You can specify a role for your task with the taskRoleArn parameter. When you
// specify a role for a task, its containers can then use the latest versions of
// the CLI or SDKs to make API requests to the Amazon Web Services services that
// are specified in the policy that's associated with the role. For more
// information, see [IAM Roles for Tasks]in the Amazon Elastic Container Service Developer Guide.
//
// You can specify a Docker networking mode for the containers in your task
// definition with the networkMode parameter. If you specify the awsvpc network
// mode, the task is allocated an elastic network interface, and you must specify a
// [NetworkConfiguration]when you create a service or run a task with the task definition. For more
// information, see [Task Networking]in the Amazon Elastic Container Service Developer Guide.
//
// [Amazon ECS Task Definitions]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html
// [Task Networking]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html
// [IAM Roles for Tasks]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html
// [NetworkConfiguration]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_NetworkConfiguration.html
func ecs_RegisterTaskDefinition(cfg aws.Config, client *ecs.Client) {
	input := &ecs.RegisterTaskDefinitionInput{
		// ContainerDefinitions: []types.ContainerDefinition, // Required
		// Family: *string, // Required
	}

	if len(_ecsContainerDefinitions) > 0 {
		if err := assignInputField(input, "ContainerDefinitions", _ecsContainerDefinitions); err != nil {
			log.Errorf("invalid --container-definitions: %s", err.Error())
			return
		}
	}
	if len(_ecsFamily) > 0 {
		input.Family = aws.String(_ecsFamily)
	}
	if len(_ecsCpu) > 0 {
		input.Cpu = aws.String(_ecsCpu)
	}
	if len(_ecsEnableFaultInjection) > 0 {
		if err := assignInputField(input, "EnableFaultInjection", _ecsEnableFaultInjection); err != nil {
			log.Errorf("invalid --enable-fault-injection: %s", err.Error())
			return
		}
	}
	if len(_ecsEphemeralStorage) > 0 {
		if err := assignInputField(input, "EphemeralStorage", _ecsEphemeralStorage); err != nil {
			log.Errorf("invalid --ephemeral-storage: %s", err.Error())
			return
		}
	}
	if len(_ecsExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_ecsExecutionRoleArn)
	}
	if len(_ecsInferenceAccelerators) > 0 {
		if err := assignInputField(input, "InferenceAccelerators", _ecsInferenceAccelerators); err != nil {
			log.Errorf("invalid --inference-accelerators: %s", err.Error())
			return
		}
	}
	if len(_ecsIpcMode) > 0 {
		if err := assignInputField(input, "IpcMode", _ecsIpcMode); err != nil {
			log.Errorf("invalid --ipc-mode: %s", err.Error())
			return
		}
	}
	if len(_ecsMemory) > 0 {
		input.Memory = aws.String(_ecsMemory)
	}
	if len(_ecsNetworkMode) > 0 {
		if err := assignInputField(input, "NetworkMode", _ecsNetworkMode); err != nil {
			log.Errorf("invalid --network-mode: %s", err.Error())
			return
		}
	}
	if len(_ecsPidMode) > 0 {
		if err := assignInputField(input, "PidMode", _ecsPidMode); err != nil {
			log.Errorf("invalid --pid-mode: %s", err.Error())
			return
		}
	}
	if len(_ecsPlacementConstraints) > 0 {
		if err := assignInputField(input, "PlacementConstraints", _ecsPlacementConstraints); err != nil {
			log.Errorf("invalid --placement-constraints: %s", err.Error())
			return
		}
	}
	if len(_ecsProxyConfiguration) > 0 {
		if err := assignInputField(input, "ProxyConfiguration", _ecsProxyConfiguration); err != nil {
			log.Errorf("invalid --proxy-configuration: %s", err.Error())
			return
		}
	}
	if len(_ecsRequiresCompatibilities) > 0 {
		if err := assignInputField(input, "RequiresCompatibilities", _ecsRequiresCompatibilities); err != nil {
			log.Errorf("invalid --requires-compatibilities: %s", err.Error())
			return
		}
	}
	if len(_ecsRuntimePlatform) > 0 {
		if err := assignInputField(input, "RuntimePlatform", _ecsRuntimePlatform); err != nil {
			log.Errorf("invalid --runtime-platform: %s", err.Error())
			return
		}
	}
	if len(_ecsTags) > 0 {
		if err := assignInputField(input, "Tags", _ecsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_ecsTaskRoleArn) > 0 {
		input.TaskRoleArn = aws.String(_ecsTaskRoleArn)
	}
	if len(_ecsVolumes) > 0 {
		if err := assignInputField(input, "Volumes", _ecsVolumes); err != nil {
			log.Errorf("invalid --volumes: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterTaskDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a new task using the specified task definition.
// On March 21, 2024, a change was made to resolve the task definition revision
// before authorization. When a task definition revision is not specified,
// authorization will occur using the latest revision of a task definition.
//
// Amazon Elastic Inference (EI) is no longer available to customers.
//
// You can allow Amazon ECS to place tasks for you, or you can customize how
// Amazon ECS places tasks using placement constraints and placement strategies.
// For more information, see [Scheduling Tasks]in the Amazon Elastic Container Service Developer
// Guide.
//
// Alternatively, you can use StartTask to use your own scheduler or place tasks
// manually on specific container instances.
//
// You can attach Amazon EBS volumes to Amazon ECS tasks by configuring the volume
// when creating or updating a service. For more information, see [Amazon EBS volumes]in the Amazon
// Elastic Container Service Developer Guide.
//
// The Amazon ECS API follows an eventual consistency model. This is because of
// the distributed nature of the system supporting the API. This means that the
// result of an API command you run that affects your Amazon ECS resources might
// not be immediately visible to all subsequent commands you run. Keep this in mind
// when you carry out an API command that immediately follows a previous API
// command.
//
// To manage eventual consistency, you can do the following:
//
// - Confirm the state of the resource before you run a command to modify it.
// Run the DescribeTasks command using an exponential backoff algorithm to ensure
// that you allow enough time for the previous command to propagate through the
// system. To do this, run the DescribeTasks command repeatedly, starting with a
// couple of seconds of wait time and increasing gradually up to five minutes of
// wait time.
//
// - Add wait time between subsequent commands, even if the DescribeTasks
// command returns an accurate response. Apply an exponential backoff algorithm
// starting with a couple of seconds of wait time, and increase gradually up to
// about five minutes of wait time.
//
// If you get a ConflictException error, the RunTask request could not be
// processed due to conflicts. The provided clientToken is already in use with a
// different RunTask request. The resourceIds are the existing task ARNs which are
// already associated with the clientToken .
//
// To fix this issue:
//
// - Run RunTask with a unique clientToken .
//
// - Run RunTask with the clientToken and the original set of parameters
//
// If you get a ClientException error, the RunTask could not be processed because
// you use managed scaling and there is a capacity error because the quota of tasks
// in the PROVISIONING per cluster has been reached. For information about the
// service quotas, see [Amazon ECS service quotas].
//
// [Scheduling Tasks]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/scheduling_tasks.html
// [Amazon EBS volumes]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ebs-volumes.html#ebs-volume-types
// [Amazon ECS service quotas]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-quotas.html
func ecs_RunTask(cfg aws.Config, client *ecs.Client) {
	input := &ecs.RunTaskInput{
		// TaskDefinition: *string, // Required
	}

	if len(_ecsTaskDefinition) > 0 {
		input.TaskDefinition = aws.String(_ecsTaskDefinition)
	}
	if len(_ecsCapacityProviderStrategy) > 0 {
		if err := assignInputField(input, "CapacityProviderStrategy", _ecsCapacityProviderStrategy); err != nil {
			log.Errorf("invalid --capacity-provider-strategy: %s", err.Error())
			return
		}
	}
	if len(_ecsClientToken) > 0 {
		input.ClientToken = aws.String(_ecsClientToken)
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsCount) > 0 {
		if err := assignInputField(input, "Count", _ecsCount); err != nil {
			log.Errorf("invalid --count: %s", err.Error())
			return
		}
	}
	if len(_ecsEnableECSManagedTags) > 0 {
		if err := assignInputField(input, "EnableECSManagedTags", _ecsEnableECSManagedTags); err != nil {
			log.Errorf("invalid --enable-ecs-managed-tags: %s", err.Error())
			return
		}
	}
	if len(_ecsEnableExecuteCommand) > 0 {
		if err := assignInputField(input, "EnableExecuteCommand", _ecsEnableExecuteCommand); err != nil {
			log.Errorf("invalid --enable-execute-command: %s", err.Error())
			return
		}
	}
	if len(_ecsGroup) > 0 {
		input.Group = aws.String(_ecsGroup)
	}
	if len(_ecsLaunchType) > 0 {
		if err := assignInputField(input, "LaunchType", _ecsLaunchType); err != nil {
			log.Errorf("invalid --launch-type: %s", err.Error())
			return
		}
	}
	if len(_ecsNetworkConfiguration) > 0 {
		if err := assignInputField(input, "NetworkConfiguration", _ecsNetworkConfiguration); err != nil {
			log.Errorf("invalid --network-configuration: %s", err.Error())
			return
		}
	}
	if len(_ecsOverrides) > 0 {
		if err := assignInputField(input, "Overrides", _ecsOverrides); err != nil {
			log.Errorf("invalid --overrides: %s", err.Error())
			return
		}
	}
	if len(_ecsPlacementConstraints) > 0 {
		if err := assignInputField(input, "PlacementConstraints", _ecsPlacementConstraints); err != nil {
			log.Errorf("invalid --placement-constraints: %s", err.Error())
			return
		}
	}
	if len(_ecsPlacementStrategy) > 0 {
		if err := assignInputField(input, "PlacementStrategy", _ecsPlacementStrategy); err != nil {
			log.Errorf("invalid --placement-strategy: %s", err.Error())
			return
		}
	}
	if len(_ecsPlatformVersion) > 0 {
		input.PlatformVersion = aws.String(_ecsPlatformVersion)
	}
	if len(_ecsPropagateTags) > 0 {
		if err := assignInputField(input, "PropagateTags", _ecsPropagateTags); err != nil {
			log.Errorf("invalid --propagate-tags: %s", err.Error())
			return
		}
	}
	if len(_ecsReferenceId) > 0 {
		input.ReferenceId = aws.String(_ecsReferenceId)
	}
	if len(_ecsStartedBy) > 0 {
		input.StartedBy = aws.String(_ecsStartedBy)
	}
	if len(_ecsTags) > 0 {
		if err := assignInputField(input, "Tags", _ecsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_ecsVolumeConfigurations) > 0 {
		if err := assignInputField(input, "VolumeConfigurations", _ecsVolumeConfigurations); err != nil {
			log.Errorf("invalid --volume-configurations: %s", err.Error())
			return
		}
	}

	if resp, err := client.RunTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a new task from the specified task definition on the specified container
// instance or instances.
//
// On March 21, 2024, a change was made to resolve the task definition revision
// before authorization. When a task definition revision is not specified,
// authorization will occur using the latest revision of a task definition.
//
// Amazon Elastic Inference (EI) is no longer available to customers.
//
// Alternatively, you can use RunTask to place tasks for you. For more
// information, see [Scheduling Tasks]in the Amazon Elastic Container Service Developer Guide.
//
// You can attach Amazon EBS volumes to Amazon ECS tasks by configuring the volume
// when creating or updating a service. For more information, see [Amazon EBS volumes]in the Amazon
// Elastic Container Service Developer Guide.
//
// [Scheduling Tasks]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/scheduling_tasks.html
// [Amazon EBS volumes]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ebs-volumes.html#ebs-volume-types
func ecs_StartTask(cfg aws.Config, client *ecs.Client) {
	input := &ecs.StartTaskInput{
		// ContainerInstances: []string, // Required
		// TaskDefinition: *string, // Required
	}

	if len(_ecsContainerInstances) > 0 {
		input.ContainerInstances = append([]string(nil), _ecsContainerInstances...)
	}
	if len(_ecsTaskDefinition) > 0 {
		input.TaskDefinition = aws.String(_ecsTaskDefinition)
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsEnableECSManagedTags) > 0 {
		if err := assignInputField(input, "EnableECSManagedTags", _ecsEnableECSManagedTags); err != nil {
			log.Errorf("invalid --enable-ecs-managed-tags: %s", err.Error())
			return
		}
	}
	if len(_ecsEnableExecuteCommand) > 0 {
		if err := assignInputField(input, "EnableExecuteCommand", _ecsEnableExecuteCommand); err != nil {
			log.Errorf("invalid --enable-execute-command: %s", err.Error())
			return
		}
	}
	if len(_ecsGroup) > 0 {
		input.Group = aws.String(_ecsGroup)
	}
	if len(_ecsNetworkConfiguration) > 0 {
		if err := assignInputField(input, "NetworkConfiguration", _ecsNetworkConfiguration); err != nil {
			log.Errorf("invalid --network-configuration: %s", err.Error())
			return
		}
	}
	if len(_ecsOverrides) > 0 {
		if err := assignInputField(input, "Overrides", _ecsOverrides); err != nil {
			log.Errorf("invalid --overrides: %s", err.Error())
			return
		}
	}
	if len(_ecsPropagateTags) > 0 {
		if err := assignInputField(input, "PropagateTags", _ecsPropagateTags); err != nil {
			log.Errorf("invalid --propagate-tags: %s", err.Error())
			return
		}
	}
	if len(_ecsReferenceId) > 0 {
		input.ReferenceId = aws.String(_ecsReferenceId)
	}
	if len(_ecsStartedBy) > 0 {
		input.StartedBy = aws.String(_ecsStartedBy)
	}
	if len(_ecsTags) > 0 {
		if err := assignInputField(input, "Tags", _ecsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_ecsVolumeConfigurations) > 0 {
		if err := assignInputField(input, "VolumeConfigurations", _ecsVolumeConfigurations); err != nil {
			log.Errorf("invalid --volume-configurations: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an ongoing service deployment.
// The following stop types are avaiable:
//
// - ROLLBACK - This option rolls back the service deployment to the previous
// service revision.
//
// You can use this option even if you didn't configure the service deployment for
//
// the rollback option.
//
// For more information, see [Stopping Amazon ECS service deployments] in the Amazon Elastic Container Service Developer
// Guide.
//
// [Stopping Amazon ECS service deployments]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/stop-service-deployment.html
func ecs_StopServiceDeployment(cfg aws.Config, client *ecs.Client) {
	input := &ecs.StopServiceDeploymentInput{
		// ServiceDeploymentArn: *string, // Required
	}

	if len(_ecsServiceDeploymentArn) > 0 {
		input.ServiceDeploymentArn = aws.String(_ecsServiceDeploymentArn)
	}
	if len(_ecsStopType) > 0 {
		if err := assignInputField(input, "StopType", _ecsStopType); err != nil {
			log.Errorf("invalid --stop-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.StopServiceDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a running task. Any tags associated with the task will be deleted.
// When you call StopTask on a task, the equivalent of docker stop is issued to
// the containers running in the task. This results in a stop signal value and a
// default 30-second timeout, after which the SIGKILL value is sent and the
// containers are forcibly stopped. This signal can be defined in your container
// image with the STOPSIGNAL instruction and will default to SIGTERM . If the
// container handles the SIGTERM value gracefully and exits within 30 seconds from
// receiving it, no SIGKILL value is sent.
//
// For Windows containers, POSIX signals do not work and runtime stops the
// container by sending a CTRL_SHUTDOWN_EVENT . For more information, see [Unable to react to graceful shutdown of (Windows) container #25982] on
// GitHub.
//
// The default 30-second timeout can be configured on the Amazon ECS container
// agent with the ECS_CONTAINER_STOP_TIMEOUT variable. For more information, see [Amazon ECS Container Agent Configuration]
// in the Amazon Elastic Container Service Developer Guide.
//
// [Unable to react to graceful shutdown of (Windows) container #25982]: https://github.com/moby/moby/issues/25982
// [Amazon ECS Container Agent Configuration]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-config.html
func ecs_StopTask(cfg aws.Config, client *ecs.Client) {
	input := &ecs.StopTaskInput{
		// Task: *string, // Required
	}

	if len(_ecsTask) > 0 {
		input.Task = aws.String(_ecsTask)
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsReason) > 0 {
		input.Reason = aws.String(_ecsReason)
	}

	if resp, err := client.StopTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action is only used by the Amazon ECS agent, and it is not intended for
// use outside of the agent.
//
// Sent to acknowledge that an attachment changed states.
func ecs_SubmitAttachmentStateChanges(cfg aws.Config, client *ecs.Client) {
	input := &ecs.SubmitAttachmentStateChangesInput{
		// Attachments: []types.AttachmentStateChange, // Required
	}

	if len(_ecsAttachments) > 0 {
		if err := assignInputField(input, "Attachments", _ecsAttachments); err != nil {
			log.Errorf("invalid --attachments: %s", err.Error())
			return
		}
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}

	if resp, err := client.SubmitAttachmentStateChanges(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action is only used by the Amazon ECS agent, and it is not intended for
// use outside of the agent.
//
// Sent to acknowledge that a container changed states.
func ecs_SubmitContainerStateChange(cfg aws.Config, client *ecs.Client) {
	input := &ecs.SubmitContainerStateChangeInput{}

	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsContainerName) > 0 {
		input.ContainerName = aws.String(_ecsContainerName)
	}
	if len(_ecsExitCode) > 0 {
		if err := assignInputField(input, "ExitCode", _ecsExitCode); err != nil {
			log.Errorf("invalid --exit-code: %s", err.Error())
			return
		}
	}
	if len(_ecsNetworkBindings) > 0 {
		if err := assignInputField(input, "NetworkBindings", _ecsNetworkBindings); err != nil {
			log.Errorf("invalid --network-bindings: %s", err.Error())
			return
		}
	}
	if len(_ecsReason) > 0 {
		input.Reason = aws.String(_ecsReason)
	}
	if len(_ecsRuntimeId) > 0 {
		input.RuntimeId = aws.String(_ecsRuntimeId)
	}
	if len(_ecsStatus) > 0 {
		input.Status = aws.String(_ecsStatus)
	}
	if len(_ecsTask) > 0 {
		input.Task = aws.String(_ecsTask)
	}

	if resp, err := client.SubmitContainerStateChange(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action is only used by the Amazon ECS agent, and it is not intended for
// use outside of the agent.
//
// Sent to acknowledge that a task changed states.
func ecs_SubmitTaskStateChange(cfg aws.Config, client *ecs.Client) {
	input := &ecs.SubmitTaskStateChangeInput{}

	if len(_ecsAttachments) > 0 {
		if err := assignInputField(input, "Attachments", _ecsAttachments); err != nil {
			log.Errorf("invalid --attachments: %s", err.Error())
			return
		}
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsContainers) > 0 {
		if err := assignInputField(input, "Containers", _ecsContainers); err != nil {
			log.Errorf("invalid --containers: %s", err.Error())
			return
		}
	}
	if len(_ecsExecutionStoppedAt) > 0 {
		if err := assignInputField(input, "ExecutionStoppedAt", _ecsExecutionStoppedAt); err != nil {
			log.Errorf("invalid --execution-stopped-at: %s", err.Error())
			return
		}
	}
	if len(_ecsManagedAgents) > 0 {
		if err := assignInputField(input, "ManagedAgents", _ecsManagedAgents); err != nil {
			log.Errorf("invalid --managed-agents: %s", err.Error())
			return
		}
	}
	if len(_ecsPullStartedAt) > 0 {
		if err := assignInputField(input, "PullStartedAt", _ecsPullStartedAt); err != nil {
			log.Errorf("invalid --pull-started-at: %s", err.Error())
			return
		}
	}
	if len(_ecsPullStoppedAt) > 0 {
		if err := assignInputField(input, "PullStoppedAt", _ecsPullStoppedAt); err != nil {
			log.Errorf("invalid --pull-stopped-at: %s", err.Error())
			return
		}
	}
	if len(_ecsReason) > 0 {
		input.Reason = aws.String(_ecsReason)
	}
	if len(_ecsStatus) > 0 {
		input.Status = aws.String(_ecsStatus)
	}
	if len(_ecsTask) > 0 {
		input.Task = aws.String(_ecsTask)
	}

	if resp, err := client.SubmitTaskStateChange(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified tags to a resource with the specified resourceArn . If
// existing tags on a resource aren't specified in the request parameters, they
// aren't changed. When a resource is deleted, the tags that are associated with
// that resource are deleted as well.
func ecs_TagResource(cfg aws.Config, client *ecs.Client) {
	input := &ecs.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_ecsResourceArn) > 0 {
		input.ResourceArn = aws.String(_ecsResourceArn)
	}
	if len(_ecsTags) > 0 {
		if err := assignInputField(input, "Tags", _ecsTags); err != nil {
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

// Deletes specified tags from a resource.
func ecs_UntagResource(cfg aws.Config, client *ecs.Client) {
	input := &ecs.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_ecsResourceArn) > 0 {
		input.ResourceArn = aws.String(_ecsResourceArn)
	}
	if len(_ecsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _ecsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the parameters for a capacity provider.
// These changes only apply to new Amazon ECS Managed Instances, or EC2 instances,
// not existing ones.
func ecs_UpdateCapacityProvider(cfg aws.Config, client *ecs.Client) {
	input := &ecs.UpdateCapacityProviderInput{
		// Name: *string, // Required
	}

	if len(_ecsName) > 0 {
		input.Name = aws.String(_ecsName)
	}
	if len(_ecsAutoScalingGroupProvider) > 0 {
		if err := assignInputField(input, "AutoScalingGroupProvider", _ecsAutoScalingGroupProvider); err != nil {
			log.Errorf("invalid --auto-scaling-group-provider: %s", err.Error())
			return
		}
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsManagedInstancesProvider) > 0 {
		if err := assignInputField(input, "ManagedInstancesProvider", _ecsManagedInstancesProvider); err != nil {
			log.Errorf("invalid --managed-instances-provider: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCapacityProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the cluster.
func ecs_UpdateCluster(cfg aws.Config, client *ecs.Client) {
	input := &ecs.UpdateClusterInput{
		// Cluster: *string, // Required
	}

	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _ecsConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_ecsServiceConnectDefaults) > 0 {
		if err := assignInputField(input, "ServiceConnectDefaults", _ecsServiceConnectDefaults); err != nil {
			log.Errorf("invalid --service-connect-defaults: %s", err.Error())
			return
		}
	}
	if len(_ecsSettings) > 0 {
		if err := assignInputField(input, "Settings", _ecsSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
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

// Modifies the settings to use for a cluster.
func ecs_UpdateClusterSettings(cfg aws.Config, client *ecs.Client) {
	input := &ecs.UpdateClusterSettingsInput{
		// Cluster: *string, // Required
		// Settings: []types.ClusterSetting, // Required
	}

	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsSettings) > 0 {
		if err := assignInputField(input, "Settings", _ecsSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateClusterSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the Amazon ECS container agent on a specified container instance.
// Updating the Amazon ECS container agent doesn't interrupt running tasks or
// services on the container instance. The process for updating the agent differs
// depending on whether your container instance was launched with the Amazon
// ECS-optimized AMI or another operating system.
//
// The UpdateContainerAgent API isn't supported for container instances using the
// Amazon ECS-optimized Amazon Linux 2 (arm64) AMI. To update the container agent,
// you can update the ecs-init package. This updates the agent. For more
// information, see [Updating the Amazon ECS container agent]in the Amazon Elastic Container Service Developer Guide.
//
// Agent updates with the UpdateContainerAgent API operation do not apply to
// Windows container instances. We recommend that you launch new container
// instances to update the agent version in your Windows clusters.
//
// The UpdateContainerAgent API requires an Amazon ECS-optimized AMI or Amazon
// Linux AMI with the ecs-init service installed and running. For help updating
// the Amazon ECS container agent on other operating systems, see [Manually updating the Amazon ECS container agent]in the Amazon
// Elastic Container Service Developer Guide.
//
// [Updating the Amazon ECS container agent]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/agent-update-ecs-ami.html
// [Manually updating the Amazon ECS container agent]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-update.html#manually_update_agent
func ecs_UpdateContainerAgent(cfg aws.Config, client *ecs.Client) {
	input := &ecs.UpdateContainerAgentInput{
		// ContainerInstance: *string, // Required
	}

	if len(_ecsContainerInstance) > 0 {
		input.ContainerInstance = aws.String(_ecsContainerInstance)
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}

	if resp, err := client.UpdateContainerAgent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the status of an Amazon ECS container instance.
// Once a container instance has reached an ACTIVE state, you can change the
// status of a container instance to DRAINING to manually remove an instance from
// a cluster, for example to perform system updates, update the Docker daemon, or
// scale down the cluster size.
//
// A container instance can't be changed to DRAINING until it has reached an ACTIVE
// status. If the instance is in any other status, an error will be received.
//
// When you set a container instance to DRAINING , Amazon ECS prevents new tasks
// from being scheduled for placement on the container instance and replacement
// service tasks are started on other container instances in the cluster if the
// resources are available. Service tasks on the container instance that are in the
// PENDING state are stopped immediately.
//
// Service tasks on the container instance that are in the RUNNING state are
// stopped and replaced according to the service's deployment configuration
// parameters, minimumHealthyPercent and maximumPercent . You can change the
// deployment configuration of your service using [UpdateService].
//
// - If minimumHealthyPercent is below 100%, the scheduler can ignore
// desiredCount temporarily during task replacement. For example, desiredCount is
// four tasks, a minimum of 50% allows the scheduler to stop two existing tasks
// before starting two new tasks. If the minimum is 100%, the service scheduler
// can't remove existing tasks until the replacement tasks are considered healthy.
// Tasks for services that do not use a load balancer are considered healthy if
// they're in the RUNNING state. Tasks for services that use a load balancer are
// considered healthy if they're in the RUNNING state and are reported as healthy
// by the load balancer.
//
// - The maximumPercent parameter represents an upper limit on the number of
// running tasks during task replacement. You can use this to define the
// replacement batch size. For example, if desiredCount is four tasks, a maximum
// of 200% starts four new tasks before stopping the four tasks to be drained,
// provided that the cluster resources required to do this are available. If the
// maximum is 100%, then replacement tasks can't start until the draining tasks
// have stopped.
//
// Any PENDING or RUNNING tasks that do not belong to a service aren't affected.
// You must wait for them to finish or stop them manually.
//
// A container instance has completed draining when it has no more RUNNING tasks.
// You can verify this using [ListTasks].
//
// When a container instance has been drained, you can set a container instance to
// ACTIVE status and once it has reached that status the Amazon ECS scheduler can
// begin scheduling tasks on the instance again.
//
// [UpdateService]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_UpdateService.html
// [ListTasks]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ListTasks.html
func ecs_UpdateContainerInstancesState(cfg aws.Config, client *ecs.Client) {
	input := &ecs.UpdateContainerInstancesStateInput{
		// ContainerInstances: []string, // Required
		// Status: types.ContainerInstanceStatus, // Required
	}

	if len(_ecsContainerInstances) > 0 {
		input.ContainerInstances = append([]string(nil), _ecsContainerInstances...)
	}
	if len(_ecsStatus) > 0 {
		if err := assignInputField(input, "Status", _ecsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}

	if resp, err := client.UpdateContainerInstancesState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing Express service configuration. Modifies container settings,
// resource allocation, auto-scaling configuration, and other service parameters
// without recreating the service.
//
// Amazon ECS creates a new service revision with updated configuration and
// performs a rolling deployment to replace existing tasks. The service remains
// available during updates, ensuring zero-downtime deployments.
//
// Some parameters like the infrastructure role cannot be modified after service
// creation and require creating a new service.
func ecs_UpdateExpressGatewayService(cfg aws.Config, client *ecs.Client) {
	input := &ecs.UpdateExpressGatewayServiceInput{
		// ServiceArn: *string, // Required
	}

	if len(_ecsServiceArn) > 0 {
		input.ServiceArn = aws.String(_ecsServiceArn)
	}
	if len(_ecsCpu) > 0 {
		input.Cpu = aws.String(_ecsCpu)
	}
	if len(_ecsExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_ecsExecutionRoleArn)
	}
	if len(_ecsHealthCheckPath) > 0 {
		input.HealthCheckPath = aws.String(_ecsHealthCheckPath)
	}
	if len(_ecsMemory) > 0 {
		input.Memory = aws.String(_ecsMemory)
	}
	if len(_ecsNetworkConfiguration) > 0 {
		if err := assignInputField(input, "NetworkConfiguration", _ecsNetworkConfiguration); err != nil {
			log.Errorf("invalid --network-configuration: %s", err.Error())
			return
		}
	}
	if len(_ecsPrimaryContainer) > 0 {
		if err := assignInputField(input, "PrimaryContainer", _ecsPrimaryContainer); err != nil {
			log.Errorf("invalid --primary-container: %s", err.Error())
			return
		}
	}
	if len(_ecsScalingTarget) > 0 {
		if err := assignInputField(input, "ScalingTarget", _ecsScalingTarget); err != nil {
			log.Errorf("invalid --scaling-target: %s", err.Error())
			return
		}
	}
	if len(_ecsTaskRoleArn) > 0 {
		input.TaskRoleArn = aws.String(_ecsTaskRoleArn)
	}

	if resp, err := client.UpdateExpressGatewayService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the parameters of a service.
// On March 21, 2024, a change was made to resolve the task definition revision
// before authorization. When a task definition revision is not specified,
// authorization will occur using the latest revision of a task definition.
//
// For services using the rolling update ( ECS ) you can update the desired count,
// deployment configuration, network configuration, load balancers, service
// registries, enable ECS managed tags option, propagate tags option, task
// placement constraints and strategies, and task definition. When you update any
// of these parameters, Amazon ECS starts new tasks with the new configuration.
//
// You can attach Amazon EBS volumes to Amazon ECS tasks by configuring the volume
// when starting or running a task, or when creating or updating a service. For
// more information, see [Amazon EBS volumes]in the Amazon Elastic Container Service Developer Guide.
// You can update your volume configurations and trigger a new deployment.
// volumeConfigurations is only supported for REPLICA service and not DAEMON
// service. If you leave volumeConfigurations null , it doesn't trigger a new
// deployment. For more information on volumes, see [Amazon EBS volumes]in the Amazon Elastic
// Container Service Developer Guide.
//
// For services using the blue/green ( CODE_DEPLOY ) deployment controller, only
// the desired count, deployment configuration, health check grace period, task
// placement constraints and strategies, enable ECS managed tags option, and
// propagate tags can be updated using this API. If the network configuration,
// platform version, task definition, or load balancer need to be updated, create a
// new CodeDeploy deployment. For more information, see [CreateDeployment]in the CodeDeploy API
// Reference.
//
// For services using an external deployment controller, you can update only the
// desired count, task placement constraints and strategies, health check grace
// period, enable ECS managed tags option, and propagate tags option, using this
// API. If the launch type, load balancer, network configuration, platform version,
// or task definition need to be updated, create a new task set For more
// information, see [CreateTaskSet].
//
// You can add to or subtract from the number of instantiations of a task
// definition in a service by specifying the cluster that the service is running in
// and a new desiredCount parameter.
//
// You can attach Amazon EBS volumes to Amazon ECS tasks by configuring the volume
// when starting or running a task, or when creating or updating a service. For
// more information, see [Amazon EBS volumes]in the Amazon Elastic Container Service Developer Guide.
//
// If you have updated the container image of your application, you can create a
// new task definition with that image and deploy it to your service. The service
// scheduler uses the minimum healthy percent and maximum percent parameters (in
// the service's deployment configuration) to determine the deployment strategy.
//
// If your updated Docker image uses the same tag as what is in the existing task
// definition for your service (for example, my_image:latest ), you don't need to
// create a new revision of your task definition. You can update the service using
// the forceNewDeployment option. The new tasks launched by the deployment pull
// the current image/tag combination from your repository when they start.
//
// You can also update the deployment configuration of a service. When a
// deployment is triggered by updating the task definition of a service, the
// service scheduler uses the deployment configuration parameters,
// minimumHealthyPercent and maximumPercent , to determine the deployment strategy.
//
// - If minimumHealthyPercent is below 100%, the scheduler can ignore
// desiredCount temporarily during a deployment. For example, if desiredCount is
// four tasks, a minimum of 50% allows the scheduler to stop two existing tasks
// before starting two new tasks. Tasks for services that don't use a load balancer
// are considered healthy if they're in the RUNNING state. Tasks for services
// that use a load balancer are considered healthy if they're in the RUNNING
// state and are reported as healthy by the load balancer.
//
// - The maximumPercent parameter represents an upper limit on the number of
// running tasks during a deployment. You can use it to define the deployment batch
// size. For example, if desiredCount is four tasks, a maximum of 200% starts
// four new tasks before stopping the four older tasks (provided that the cluster
// resources required to do this are available).
//
// When [UpdateService] stops a task during a deployment, the equivalent of docker stop is issued
// to the containers running in the task. This results in a SIGTERM and a
// 30-second timeout. After this, SIGKILL is sent and the containers are forcibly
// stopped. If the container handles the SIGTERM gracefully and exits within 30
// seconds from receiving it, no SIGKILL is sent.
//
// When the service scheduler launches new tasks, it determines task placement in
// your cluster with the following logic.
//
// - Determine which of the container instances in your cluster can support your
// service's task definition. For example, they have the required CPU, memory,
// ports, and container instance attributes.
//
// - By default, the service scheduler attempts to balance tasks across
// Availability Zones in this manner even though you can choose a different
// placement strategy.
//
// - Sort the valid container instances by the fewest number of running tasks
// for this service in the same Availability Zone as the instance. For example, if
// zone A has one running service task and zones B and C each have zero, valid
// container instances in either zone B or C are considered optimal for placement.
//
// - Place the new service task on a valid container instance in an optimal
// Availability Zone (based on the previous steps), favoring container instances
// with the fewest number of running tasks for this service.
//
// When the service scheduler stops running tasks, it attempts to maintain balance
// across the Availability Zones in your cluster using the following logic:
//
// - Sort the container instances by the largest number of running tasks for
// this service in the same Availability Zone as the instance. For example, if zone
// A has one running service task and zones B and C each have two, container
// instances in either zone B or C are considered optimal for termination.
//
// - Stop the task on a container instance in an optimal Availability Zone
// (based on the previous steps), favoring container instances with the largest
// number of running tasks for this service.
//
// [CreateTaskSet]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateTaskSet.html
// [UpdateService]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_UpdateService.html
// [Amazon EBS volumes]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ebs-volumes.html#ebs-volume-types
// [CreateDeployment]: https://docs.aws.amazon.com/codedeploy/latest/APIReference/API_CreateDeployment.html
func ecs_UpdateService(cfg aws.Config, client *ecs.Client) {
	input := &ecs.UpdateServiceInput{
		// Service: *string, // Required
	}

	if len(_ecsService) > 0 {
		input.Service = aws.String(_ecsService)
	}
	if len(_ecsAvailabilityZoneRebalancing) > 0 {
		if err := assignInputField(input, "AvailabilityZoneRebalancing", _ecsAvailabilityZoneRebalancing); err != nil {
			log.Errorf("invalid --availability-zone-rebalancing: %s", err.Error())
			return
		}
	}
	if len(_ecsCapacityProviderStrategy) > 0 {
		if err := assignInputField(input, "CapacityProviderStrategy", _ecsCapacityProviderStrategy); err != nil {
			log.Errorf("invalid --capacity-provider-strategy: %s", err.Error())
			return
		}
	}
	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsDeploymentConfiguration) > 0 {
		if err := assignInputField(input, "DeploymentConfiguration", _ecsDeploymentConfiguration); err != nil {
			log.Errorf("invalid --deployment-configuration: %s", err.Error())
			return
		}
	}
	if len(_ecsDeploymentController) > 0 {
		if err := assignInputField(input, "DeploymentController", _ecsDeploymentController); err != nil {
			log.Errorf("invalid --deployment-controller: %s", err.Error())
			return
		}
	}
	if len(_ecsDesiredCount) > 0 {
		if err := assignInputField(input, "DesiredCount", _ecsDesiredCount); err != nil {
			log.Errorf("invalid --desired-count: %s", err.Error())
			return
		}
	}
	if len(_ecsEnableECSManagedTags) > 0 {
		if err := assignInputField(input, "EnableECSManagedTags", _ecsEnableECSManagedTags); err != nil {
			log.Errorf("invalid --enable-ecs-managed-tags: %s", err.Error())
			return
		}
	}
	if len(_ecsEnableExecuteCommand) > 0 {
		if err := assignInputField(input, "EnableExecuteCommand", _ecsEnableExecuteCommand); err != nil {
			log.Errorf("invalid --enable-execute-command: %s", err.Error())
			return
		}
	}
	if len(_ecsForceNewDeployment) > 0 {
		if err := assignInputField(input, "ForceNewDeployment", _ecsForceNewDeployment); err != nil {
			log.Errorf("invalid --force-new-deployment: %s", err.Error())
			return
		}
	}
	if len(_ecsHealthCheckGracePeriodSeconds) > 0 {
		if err := assignInputField(input, "HealthCheckGracePeriodSeconds", _ecsHealthCheckGracePeriodSeconds); err != nil {
			log.Errorf("invalid --health-check-grace-period-seconds: %s", err.Error())
			return
		}
	}
	if len(_ecsLoadBalancers) > 0 {
		if err := assignInputField(input, "LoadBalancers", _ecsLoadBalancers); err != nil {
			log.Errorf("invalid --load-balancers: %s", err.Error())
			return
		}
	}
	if len(_ecsNetworkConfiguration) > 0 {
		if err := assignInputField(input, "NetworkConfiguration", _ecsNetworkConfiguration); err != nil {
			log.Errorf("invalid --network-configuration: %s", err.Error())
			return
		}
	}
	if len(_ecsPlacementConstraints) > 0 {
		if err := assignInputField(input, "PlacementConstraints", _ecsPlacementConstraints); err != nil {
			log.Errorf("invalid --placement-constraints: %s", err.Error())
			return
		}
	}
	if len(_ecsPlacementStrategy) > 0 {
		if err := assignInputField(input, "PlacementStrategy", _ecsPlacementStrategy); err != nil {
			log.Errorf("invalid --placement-strategy: %s", err.Error())
			return
		}
	}
	if len(_ecsPlatformVersion) > 0 {
		input.PlatformVersion = aws.String(_ecsPlatformVersion)
	}
	if len(_ecsPropagateTags) > 0 {
		if err := assignInputField(input, "PropagateTags", _ecsPropagateTags); err != nil {
			log.Errorf("invalid --propagate-tags: %s", err.Error())
			return
		}
	}
	if len(_ecsServiceConnectConfiguration) > 0 {
		if err := assignInputField(input, "ServiceConnectConfiguration", _ecsServiceConnectConfiguration); err != nil {
			log.Errorf("invalid --service-connect-configuration: %s", err.Error())
			return
		}
	}
	if len(_ecsServiceRegistries) > 0 {
		if err := assignInputField(input, "ServiceRegistries", _ecsServiceRegistries); err != nil {
			log.Errorf("invalid --service-registries: %s", err.Error())
			return
		}
	}
	if len(_ecsTaskDefinition) > 0 {
		input.TaskDefinition = aws.String(_ecsTaskDefinition)
	}
	if len(_ecsVolumeConfigurations) > 0 {
		if err := assignInputField(input, "VolumeConfigurations", _ecsVolumeConfigurations); err != nil {
			log.Errorf("invalid --volume-configurations: %s", err.Error())
			return
		}
	}
	if len(_ecsVpcLatticeConfigurations) > 0 {
		if err := assignInputField(input, "VpcLatticeConfigurations", _ecsVpcLatticeConfigurations); err != nil {
			log.Errorf("invalid --vpc-lattice-configurations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies which task set in a service is the primary task set. Any parameters
// that are updated on the primary task set in a service will transition to the
// service. This is used when a service uses the EXTERNAL deployment controller
// type. For more information, see [Amazon ECS Deployment Types]in the Amazon Elastic Container Service
// Developer Guide.
//
// [Amazon ECS Deployment Types]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html
func ecs_UpdateServicePrimaryTaskSet(cfg aws.Config, client *ecs.Client) {
	input := &ecs.UpdateServicePrimaryTaskSetInput{
		// Cluster: *string, // Required
		// PrimaryTaskSet: *string, // Required
		// Service: *string, // Required
	}

	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsPrimaryTaskSet) > 0 {
		input.PrimaryTaskSet = aws.String(_ecsPrimaryTaskSet)
	}
	if len(_ecsService) > 0 {
		input.Service = aws.String(_ecsService)
	}

	if resp, err := client.UpdateServicePrimaryTaskSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the protection status of a task. You can set protectionEnabled to true
// to protect your task from termination during scale-in events from [Service Autoscaling]or [deployments].
//
// Task-protection, by default, expires after 2 hours at which point Amazon ECS
// clears the protectionEnabled property making the task eligible for termination
// by a subsequent scale-in event.
//
// You can specify a custom expiration period for task protection from 1 minute to
// up to 2,880 minutes (48 hours). To specify the custom expiration period, set the
// expiresInMinutes property. The expiresInMinutes property is always reset when
// you invoke this operation for a task that already has protectionEnabled set to
// true . You can keep extending the protection expiration period of a task by
// invoking this operation repeatedly.
//
// To learn more about Amazon ECS task protection, see [Task scale-in protection] in the Amazon Elastic
// Container Service Developer Guide .
//
// This operation is only supported for tasks belonging to an Amazon ECS service.
// Invoking this operation for a standalone task will result in an TASK_NOT_VALID
// failure. For more information, see [API failure reasons].
//
// If you prefer to set task protection from within the container, we recommend
// using the [Task scale-in protection endpoint].
//
// [deployments]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html
// [API failure reasons]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/api_failures_messages.html
// [Task scale-in protection endpoint]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-scale-in-protection-endpoint.html
// [Task scale-in protection]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-scale-in-protection.html
// [Service Autoscaling]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-auto-scaling.html
func ecs_UpdateTaskProtection(cfg aws.Config, client *ecs.Client) {
	input := &ecs.UpdateTaskProtectionInput{
		// Cluster: *string, // Required
		// ProtectionEnabled: bool, // Required
		// Tasks: []string, // Required
	}

	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsProtectionEnabled) > 0 {
		if err := assignInputField(input, "ProtectionEnabled", _ecsProtectionEnabled); err != nil {
			log.Errorf("invalid --protection-enabled: %s", err.Error())
			return
		}
	}
	if len(_ecsTasks) > 0 {
		input.Tasks = append([]string(nil), _ecsTasks...)
	}
	if len(_ecsExpiresInMinutes) > 0 {
		if err := assignInputField(input, "ExpiresInMinutes", _ecsExpiresInMinutes); err != nil {
			log.Errorf("invalid --expires-in-minutes: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTaskProtection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies a task set. This is used when a service uses the EXTERNAL deployment
// controller type. For more information, see [Amazon ECS Deployment Types]in the Amazon Elastic Container
// Service Developer Guide.
//
// [Amazon ECS Deployment Types]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html
func ecs_UpdateTaskSet(cfg aws.Config, client *ecs.Client) {
	input := &ecs.UpdateTaskSetInput{
		// Cluster: *string, // Required
		// Scale: *types.Scale, // Required
		// Service: *string, // Required
		// TaskSet: *string, // Required
	}

	if len(_ecsCluster) > 0 {
		input.Cluster = aws.String(_ecsCluster)
	}
	if len(_ecsScale) > 0 {
		if err := assignInputField(input, "Scale", _ecsScale); err != nil {
			log.Errorf("invalid --scale: %s", err.Error())
			return
		}
	}
	if len(_ecsService) > 0 {
		input.Service = aws.String(_ecsService)
	}
	if len(_ecsTaskSet) > 0 {
		input.TaskSet = aws.String(_ecsTaskSet)
	}

	if resp, err := client.UpdateTaskSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_ecsCmd)
	_ecsCmd.Flags().SortFlags = false

	_ecsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_ecsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_ecsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_ecsCmd.Flags().StringVarP(&_ecsAttachments, "attachments", "", "", "Attachments")
	_ecsCmd.Flags().StringVarP(&_ecsAttributeName, "attribute-name", "", "", "Attribute Name")
	_ecsCmd.Flags().StringVarP(&_ecsAttributeValue, "attribute-value", "", "", "Attribute Value")
	_ecsCmd.Flags().StringVarP(&_ecsAttributes, "attributes", "", "", "Attributes")
	_ecsCmd.Flags().StringVarP(&_ecsAutoScalingGroupProvider, "auto-scaling-group-provider", "", "", "Auto Scaling Group Provider")
	_ecsCmd.Flags().StringVarP(&_ecsAvailabilityZoneRebalancing, "availability-zone-rebalancing", "", "", "Availability Zone Rebalancing")
	_ecsCmd.Flags().StringVarP(&_ecsCapacityProvider, "capacity-provider", "", "", "Capacity Provider")
	_ecsCmd.Flags().StringVarP(&_ecsCapacityProviderStrategy, "capacity-provider-strategy", "", "", "Capacity Provider Strategy")
	_ecsCmd.Flags().StringSliceVarP(&_ecsCapacityProviders, "capacity-providers", "", nil, "Capacity Providers")
	_ecsCmd.Flags().StringVarP(&_ecsClientToken, "client-token", "", "", "Client Token")
	_ecsCmd.Flags().StringVarP(&_ecsCluster, "cluster", "", "", "Cluster")
	_ecsCmd.Flags().StringVarP(&_ecsClusterName, "cluster-name", "", "", "Cluster Name")
	_ecsCmd.Flags().StringSliceVarP(&_ecsClusters, "clusters", "", nil, "Clusters")
	_ecsCmd.Flags().StringVarP(&_ecsCommand, "command", "", "", "Command")
	_ecsCmd.Flags().StringVarP(&_ecsConfiguration, "configuration", "", "", "Configuration")
	_ecsCmd.Flags().StringVarP(&_ecsContainer, "container", "", "", "Container")
	_ecsCmd.Flags().StringVarP(&_ecsContainerDefinitions, "container-definitions", "", "", "Container Definitions")
	_ecsCmd.Flags().StringVarP(&_ecsContainerInstance, "container-instance", "", "", "Container Instance")
	_ecsCmd.Flags().StringVarP(&_ecsContainerInstanceArn, "container-instance-arn", "", "", "Container Instance ARN")
	_ecsCmd.Flags().StringSliceVarP(&_ecsContainerInstances, "container-instances", "", nil, "Container Instances")
	_ecsCmd.Flags().StringVarP(&_ecsContainerName, "container-name", "", "", "Container Name")
	_ecsCmd.Flags().StringVarP(&_ecsContainers, "containers", "", "", "Containers")
	_ecsCmd.Flags().StringVarP(&_ecsCount, "count", "", "", "Count")
	_ecsCmd.Flags().StringVarP(&_ecsCpu, "cpu", "", "", "CPU")
	_ecsCmd.Flags().StringVarP(&_ecsCreatedAt, "created-at", "", "", "Created At")
	_ecsCmd.Flags().StringVarP(&_ecsDefaultCapacityProviderStrategy, "default-capacity-provider-strategy", "", "", "Default Capacity Provider Strategy")
	_ecsCmd.Flags().StringVarP(&_ecsDeploymentConfiguration, "deployment-configuration", "", "", "Deployment Configuration")
	_ecsCmd.Flags().StringVarP(&_ecsDeploymentController, "deployment-controller", "", "", "Deployment Controller")
	_ecsCmd.Flags().StringVarP(&_ecsDesiredCount, "desired-count", "", "", "Desired Count")
	_ecsCmd.Flags().StringVarP(&_ecsDesiredStatus, "desired-status", "", "", "Desired Status")
	_ecsCmd.Flags().StringVarP(&_ecsEffectiveSettings, "effective-settings", "", "", "Effective Settings")
	_ecsCmd.Flags().StringVarP(&_ecsEnableECSManagedTags, "enable-ecs-managed-tags", "", "", "Enable Ecs Managed Tags")
	_ecsCmd.Flags().StringVarP(&_ecsEnableExecuteCommand, "enable-execute-command", "", "", "Enable Execute Command")
	_ecsCmd.Flags().StringVarP(&_ecsEnableFaultInjection, "enable-fault-injection", "", "", "Enable Fault Injection")
	_ecsCmd.Flags().StringVarP(&_ecsEphemeralStorage, "ephemeral-storage", "", "", "Ephemeral Storage")
	_ecsCmd.Flags().StringVarP(&_ecsExecutionRoleArn, "execution-role-arn", "", "", "Execution Role ARN")
	_ecsCmd.Flags().StringVarP(&_ecsExecutionStoppedAt, "execution-stopped-at", "", "", "Execution Stopped At")
	_ecsCmd.Flags().StringVarP(&_ecsExitCode, "exit-code", "", "", "Exit Code")
	_ecsCmd.Flags().StringVarP(&_ecsExpiresInMinutes, "expires-in-minutes", "", "", "Expires In Minutes")
	_ecsCmd.Flags().StringVarP(&_ecsExternalId, "external-id", "", "", "External ID")
	_ecsCmd.Flags().StringVarP(&_ecsFamily, "family", "", "", "Family")
	_ecsCmd.Flags().StringVarP(&_ecsFamilyPrefix, "family-prefix", "", "", "Family Prefix")
	_ecsCmd.Flags().StringVarP(&_ecsFilter, "filter", "", "", "Filter")
	_ecsCmd.Flags().StringVarP(&_ecsForce, "force", "", "", "Force")
	_ecsCmd.Flags().StringVarP(&_ecsForceNewDeployment, "force-new-deployment", "", "", "Force New Deployment")
	_ecsCmd.Flags().StringVarP(&_ecsGroup, "group", "", "", "Group")
	_ecsCmd.Flags().StringVarP(&_ecsHealthCheckGracePeriodSeconds, "health-check-grace-period-seconds", "", "", "Health Check Grace Period Seconds")
	_ecsCmd.Flags().StringVarP(&_ecsHealthCheckPath, "health-check-path", "", "", "Health Check Path")
	_ecsCmd.Flags().StringVarP(&_ecsInclude, "include", "", "", "Include")
	_ecsCmd.Flags().StringVarP(&_ecsInferenceAccelerators, "inference-accelerators", "", "", "Inference Accelerators")
	_ecsCmd.Flags().StringVarP(&_ecsInfrastructureRoleArn, "infrastructure-role-arn", "", "", "Infrastructure Role ARN")
	_ecsCmd.Flags().StringVarP(&_ecsInstanceIdentityDocument, "instance-identity-document", "", "", "Instance Identity Document")
	_ecsCmd.Flags().StringVarP(&_ecsInstanceIdentityDocumentSignature, "instance-identity-document-signature", "", "", "Instance Identity Document Signature")
	_ecsCmd.Flags().StringVarP(&_ecsInteractive, "interactive", "", "", "Interactive")
	_ecsCmd.Flags().StringVarP(&_ecsIpcMode, "ipc-mode", "", "", "Ipc Mode")
	_ecsCmd.Flags().StringVarP(&_ecsLaunchType, "launch-type", "", "", "Launch Type")
	_ecsCmd.Flags().StringVarP(&_ecsLoadBalancers, "load-balancers", "", "", "Load Balancers")
	_ecsCmd.Flags().StringVarP(&_ecsManagedAgents, "managed-agents", "", "", "Managed Agents")
	_ecsCmd.Flags().StringVarP(&_ecsManagedInstancesProvider, "managed-instances-provider", "", "", "Managed Instances Provider")
	_ecsCmd.Flags().StringVarP(&_ecsMaxResults, "max-results", "", "", "Max Results")
	_ecsCmd.Flags().StringVarP(&_ecsMemory, "memory", "", "", "Memory")
	_ecsCmd.Flags().StringVarP(&_ecsName, "name", "", "", "Name")
	_ecsCmd.Flags().StringVarP(&_ecsNamespace, "namespace", "", "", "Namespace")
	_ecsCmd.Flags().StringVarP(&_ecsNetworkBindings, "network-bindings", "", "", "Network Bindings")
	_ecsCmd.Flags().StringVarP(&_ecsNetworkConfiguration, "network-configuration", "", "", "Network Configuration")
	_ecsCmd.Flags().StringVarP(&_ecsNetworkMode, "network-mode", "", "", "Network Mode")
	_ecsCmd.Flags().StringVarP(&_ecsNextToken, "next-token", "", "", "Next Token")
	_ecsCmd.Flags().StringVarP(&_ecsOverrides, "overrides", "", "", "Overrides")
	_ecsCmd.Flags().StringVarP(&_ecsPidMode, "pid-mode", "", "", "Pid Mode")
	_ecsCmd.Flags().StringVarP(&_ecsPlacementConstraints, "placement-constraints", "", "", "Placement Constraints")
	_ecsCmd.Flags().StringVarP(&_ecsPlacementStrategy, "placement-strategy", "", "", "Placement Strategy")
	_ecsCmd.Flags().StringVarP(&_ecsPlatformDevices, "platform-devices", "", "", "Platform Devices")
	_ecsCmd.Flags().StringVarP(&_ecsPlatformVersion, "platform-version", "", "", "Platform Version")
	_ecsCmd.Flags().StringVarP(&_ecsPrimaryContainer, "primary-container", "", "", "Primary Container")
	_ecsCmd.Flags().StringVarP(&_ecsPrimaryTaskSet, "primary-task-set", "", "", "Primary Task Set")
	_ecsCmd.Flags().StringVarP(&_ecsPrincipalArn, "principal-arn", "", "", "Principal ARN")
	_ecsCmd.Flags().StringVarP(&_ecsPropagateTags, "propagate-tags", "", "", "Propagate Tags")
	_ecsCmd.Flags().StringVarP(&_ecsProtectionEnabled, "protection-enabled", "", "", "Protection Enabled")
	_ecsCmd.Flags().StringVarP(&_ecsProxyConfiguration, "proxy-configuration", "", "", "Proxy Configuration")
	_ecsCmd.Flags().StringVarP(&_ecsPullStartedAt, "pull-started-at", "", "", "Pull Started At")
	_ecsCmd.Flags().StringVarP(&_ecsPullStoppedAt, "pull-stopped-at", "", "", "Pull Stopped At")
	_ecsCmd.Flags().StringVarP(&_ecsReason, "reason", "", "", "Reason")
	_ecsCmd.Flags().StringVarP(&_ecsReferenceId, "reference-id", "", "", "Reference ID")
	_ecsCmd.Flags().StringVarP(&_ecsRequiresCompatibilities, "requires-compatibilities", "", "", "Requires Compatibilities")
	_ecsCmd.Flags().StringVarP(&_ecsResourceArn, "resource-arn", "", "", "Resource ARN")
	_ecsCmd.Flags().StringVarP(&_ecsResourceManagementType, "resource-management-type", "", "", "Resource Management Type")
	_ecsCmd.Flags().StringVarP(&_ecsRole, "role", "", "", "Role")
	_ecsCmd.Flags().StringVarP(&_ecsRuntimeId, "runtime-id", "", "", "Runtime ID")
	_ecsCmd.Flags().StringVarP(&_ecsRuntimePlatform, "runtime-platform", "", "", "Runtime Platform")
	_ecsCmd.Flags().StringVarP(&_ecsScale, "scale", "", "", "Scale")
	_ecsCmd.Flags().StringVarP(&_ecsScalingTarget, "scaling-target", "", "", "Scaling Target")
	_ecsCmd.Flags().StringVarP(&_ecsSchedulingStrategy, "scheduling-strategy", "", "", "Scheduling Strategy")
	_ecsCmd.Flags().StringVarP(&_ecsService, "service", "", "", "Service")
	_ecsCmd.Flags().StringVarP(&_ecsServiceArn, "service-arn", "", "", "Service ARN")
	_ecsCmd.Flags().StringVarP(&_ecsServiceConnectConfiguration, "service-connect-configuration", "", "", "Service Connect Configuration")
	_ecsCmd.Flags().StringVarP(&_ecsServiceConnectDefaults, "service-connect-defaults", "", "", "Service Connect Defaults")
	_ecsCmd.Flags().StringVarP(&_ecsServiceDeploymentArn, "service-deployment-arn", "", "", "Service Deployment ARN")
	_ecsCmd.Flags().StringSliceVarP(&_ecsServiceDeploymentArns, "service-deployment-arns", "", nil, "Service Deployment Arns")
	_ecsCmd.Flags().StringVarP(&_ecsServiceName, "service-name", "", "", "Service Name")
	_ecsCmd.Flags().StringVarP(&_ecsServiceRegistries, "service-registries", "", "", "Service Registries")
	_ecsCmd.Flags().StringSliceVarP(&_ecsServiceRevisionArns, "service-revision-arns", "", nil, "Service Revision Arns")
	_ecsCmd.Flags().StringSliceVarP(&_ecsServices, "services", "", nil, "Services")
	_ecsCmd.Flags().StringVarP(&_ecsSettings, "settings", "", "", "Settings")
	_ecsCmd.Flags().StringVarP(&_ecsSort, "sort", "", "", "Sort")
	_ecsCmd.Flags().StringVarP(&_ecsStartedBy, "started-by", "", "", "Started By")
	_ecsCmd.Flags().StringVarP(&_ecsStatus, "status", "", "", "Status")
	_ecsCmd.Flags().StringVarP(&_ecsStopType, "stop-type", "", "", "Stop Type")
	_ecsCmd.Flags().StringSliceVarP(&_ecsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_ecsCmd.Flags().StringVarP(&_ecsTags, "tags", "", "", "Tags")
	_ecsCmd.Flags().StringVarP(&_ecsTargetType, "target-type", "", "", "Target Type")
	_ecsCmd.Flags().StringVarP(&_ecsTask, "task", "", "", "Task")
	_ecsCmd.Flags().StringVarP(&_ecsTaskDefinition, "task-definition", "", "", "Task Definition")
	_ecsCmd.Flags().StringSliceVarP(&_ecsTaskDefinitions, "task-definitions", "", nil, "Task Definitions")
	_ecsCmd.Flags().StringVarP(&_ecsTaskRoleArn, "task-role-arn", "", "", "Task Role ARN")
	_ecsCmd.Flags().StringVarP(&_ecsTaskSet, "task-set", "", "", "Task Set")
	_ecsCmd.Flags().StringSliceVarP(&_ecsTaskSets, "task-sets", "", nil, "Task Sets")
	_ecsCmd.Flags().StringSliceVarP(&_ecsTasks, "tasks", "", nil, "Tasks")
	_ecsCmd.Flags().StringVarP(&_ecsTotalResources, "total-resources", "", "", "Total Resources")
	_ecsCmd.Flags().StringVarP(&_ecsValue, "value", "", "", "Value")
	_ecsCmd.Flags().StringVarP(&_ecsVersionInfo, "version-info", "", "", "Version Info")
	_ecsCmd.Flags().StringVarP(&_ecsVolumeConfigurations, "volume-configurations", "", "", "Volume Configurations")
	_ecsCmd.Flags().StringVarP(&_ecsVolumes, "volumes", "", "", "Volumes")
	_ecsCmd.Flags().StringVarP(&_ecsVpcLatticeConfigurations, "vpc-lattice-configurations", "", "", "VPC Lattice Configurations")

	_ecsCmd.Flags().BoolVarP(&_ecsCreateCapacityProvider, "create-capacity-provider", "", false, "Create Capacity Provider")
	_ecsCmd.Flags().BoolVarP(&_ecsCreateCluster, "create-cluster", "", false, "Create Cluster")
	_ecsCmd.Flags().BoolVarP(&_ecsCreateExpressGatewayService, "create-express-gateway-service", "", false, "Create Express Gateway Service")
	_ecsCmd.Flags().BoolVarP(&_ecsCreateService, "create-service", "", false, "Create Service")
	_ecsCmd.Flags().BoolVarP(&_ecsCreateTaskSet, "create-task-set", "", false, "Create Task Set")
	_ecsCmd.Flags().BoolVarP(&_ecsDeleteAccountSetting, "delete-account-setting", "", false, "Delete Account Setting")
	_ecsCmd.Flags().BoolVarP(&_ecsDeleteAttributes, "delete-attributes", "", false, "Delete Attributes")
	_ecsCmd.Flags().BoolVarP(&_ecsDeleteCapacityProvider, "delete-capacity-provider", "", false, "Delete Capacity Provider")
	_ecsCmd.Flags().BoolVarP(&_ecsDeleteCluster, "delete-cluster", "", false, "Delete Cluster")
	_ecsCmd.Flags().BoolVarP(&_ecsDeleteExpressGatewayService, "delete-express-gateway-service", "", false, "Delete Express Gateway Service")
	_ecsCmd.Flags().BoolVarP(&_ecsDeleteService, "delete-service", "", false, "Delete Service")
	_ecsCmd.Flags().BoolVarP(&_ecsDeleteTaskDefinitions, "delete-task-definitions", "", false, "Delete Task Definitions")
	_ecsCmd.Flags().BoolVarP(&_ecsDeleteTaskSet, "delete-task-set", "", false, "Delete Task Set")
	_ecsCmd.Flags().BoolVarP(&_ecsDeregisterContainerInstance, "deregister-container-instance", "", false, "Deregister Container Instance")
	_ecsCmd.Flags().BoolVarP(&_ecsDeregisterTaskDefinition, "deregister-task-definition", "", false, "Deregister Task Definition")
	_ecsCmd.Flags().BoolVarP(&_ecsDescribeCapacityProviders, "describe-capacity-providers", "", false, "Describe Capacity Providers")
	_ecsCmd.Flags().BoolVarP(&_ecsDescribeClusters, "describe-clusters", "", false, "Describe Clusters")
	_ecsCmd.Flags().BoolVarP(&_ecsDescribeContainerInstances, "describe-container-instances", "", false, "Describe Container Instances")
	_ecsCmd.Flags().BoolVarP(&_ecsDescribeExpressGatewayService, "describe-express-gateway-service", "", false, "Describe Express Gateway Service")
	_ecsCmd.Flags().BoolVarP(&_ecsDescribeServiceDeployments, "describe-service-deployments", "", false, "Describe Service Deployments")
	_ecsCmd.Flags().BoolVarP(&_ecsDescribeServiceRevisions, "describe-service-revisions", "", false, "Describe Service Revisions")
	_ecsCmd.Flags().BoolVarP(&_ecsDescribeServices, "describe-services", "", false, "Describe Services")
	_ecsCmd.Flags().BoolVarP(&_ecsDescribeTaskDefinition, "describe-task-definition", "", false, "Describe Task Definition")
	_ecsCmd.Flags().BoolVarP(&_ecsDescribeTaskSets, "describe-task-sets", "", false, "Describe Task Sets")
	_ecsCmd.Flags().BoolVarP(&_ecsDescribeTasks, "describe-tasks", "", false, "Describe Tasks")
	_ecsCmd.Flags().BoolVarP(&_ecsDiscoverPollEndpoint, "discover-poll-endpoint", "", false, "Discover Poll Endpoint")
	_ecsCmd.Flags().BoolVarP(&_ecsExecuteCommand, "execute-command", "", false, "Execute Command")
	_ecsCmd.Flags().BoolVarP(&_ecsGetTaskProtection, "get-task-protection", "", false, "Get Task Protection")
	_ecsCmd.Flags().BoolVarP(&_ecsListAccountSettings, "list-account-settings", "", false, "List Account Settings")
	_ecsCmd.Flags().BoolVarP(&_ecsListAttributes, "list-attributes", "", false, "List Attributes")
	_ecsCmd.Flags().BoolVarP(&_ecsListClusters, "list-clusters", "", false, "List Clusters")
	_ecsCmd.Flags().BoolVarP(&_ecsListContainerInstances, "list-container-instances", "", false, "List Container Instances")
	_ecsCmd.Flags().BoolVarP(&_ecsListServiceDeployments, "list-service-deployments", "", false, "List Service Deployments")
	_ecsCmd.Flags().BoolVarP(&_ecsListServices, "list-services", "", false, "List Services")
	_ecsCmd.Flags().BoolVarP(&_ecsListServicesByNamespace, "list-services-by-namespace", "", false, "List Services By Namespace")
	_ecsCmd.Flags().BoolVarP(&_ecsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_ecsCmd.Flags().BoolVarP(&_ecsListTaskDefinitionFamilies, "list-task-definition-families", "", false, "List Task Definition Families")
	_ecsCmd.Flags().BoolVarP(&_ecsListTaskDefinitions, "list-task-definitions", "", false, "List Task Definitions")
	_ecsCmd.Flags().BoolVarP(&_ecsListTasks, "list-tasks", "", false, "List Tasks")
	_ecsCmd.Flags().BoolVarP(&_ecsPutAccountSetting, "put-account-setting", "", false, "Put Account Setting")
	_ecsCmd.Flags().BoolVarP(&_ecsPutAccountSettingDefault, "put-account-setting-default", "", false, "Put Account Setting Default")
	_ecsCmd.Flags().BoolVarP(&_ecsPutAttributes, "put-attributes", "", false, "Put Attributes")
	_ecsCmd.Flags().BoolVarP(&_ecsPutClusterCapacityProviders, "put-cluster-capacity-providers", "", false, "Put Cluster Capacity Providers")
	_ecsCmd.Flags().BoolVarP(&_ecsRegisterContainerInstance, "register-container-instance", "", false, "Register Container Instance")
	_ecsCmd.Flags().BoolVarP(&_ecsRegisterTaskDefinition, "register-task-definition", "", false, "Register Task Definition")
	_ecsCmd.Flags().BoolVarP(&_ecsRunTask, "run-task", "", false, "Run Task")
	_ecsCmd.Flags().BoolVarP(&_ecsStartTask, "start-task", "", false, "Start Task")
	_ecsCmd.Flags().BoolVarP(&_ecsStopServiceDeployment, "stop-service-deployment", "", false, "Stop Service Deployment")
	_ecsCmd.Flags().BoolVarP(&_ecsStopTask, "stop-task", "", false, "Stop Task")
	_ecsCmd.Flags().BoolVarP(&_ecsSubmitAttachmentStateChanges, "submit-attachment-state-changes", "", false, "Submit Attachment State Changes")
	_ecsCmd.Flags().BoolVarP(&_ecsSubmitContainerStateChange, "submit-container-state-change", "", false, "Submit Container State Change")
	_ecsCmd.Flags().BoolVarP(&_ecsSubmitTaskStateChange, "submit-task-state-change", "", false, "Submit Task State Change")
	_ecsCmd.Flags().BoolVarP(&_ecsTagResource, "tag-resource", "", false, "Tag Resource")
	_ecsCmd.Flags().BoolVarP(&_ecsUntagResource, "untag-resource", "", false, "Untag Resource")
	_ecsCmd.Flags().BoolVarP(&_ecsUpdateCapacityProvider, "update-capacity-provider", "", false, "Update Capacity Provider")
	_ecsCmd.Flags().BoolVarP(&_ecsUpdateCluster, "update-cluster", "", false, "Update Cluster")
	_ecsCmd.Flags().BoolVarP(&_ecsUpdateClusterSettings, "update-cluster-settings", "", false, "Update Cluster Settings")
	_ecsCmd.Flags().BoolVarP(&_ecsUpdateContainerAgent, "update-container-agent", "", false, "Update Container Agent")
	_ecsCmd.Flags().BoolVarP(&_ecsUpdateContainerInstancesState, "update-container-instances-state", "", false, "Update Container Instances State")
	_ecsCmd.Flags().BoolVarP(&_ecsUpdateExpressGatewayService, "update-express-gateway-service", "", false, "Update Express Gateway Service")
	_ecsCmd.Flags().BoolVarP(&_ecsUpdateService, "update-service", "", false, "Update Service")
	_ecsCmd.Flags().BoolVarP(&_ecsUpdateServicePrimaryTaskSet, "update-service-primary-task-set", "", false, "Update Service Primary Task Set")
	_ecsCmd.Flags().BoolVarP(&_ecsUpdateTaskProtection, "update-task-protection", "", false, "Update Task Protection")
	_ecsCmd.Flags().BoolVarP(&_ecsUpdateTaskSet, "update-task-set", "", false, "Update Task Set")

}
