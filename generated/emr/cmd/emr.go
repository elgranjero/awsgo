package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// emrCmd represents the emr command
var _emrCmd = &cobra.Command{
	Use:   "emr",
	Short: "AWS emr CLI",
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
		client := emr.NewFromConfig(cfg)
		if _emrAddInstanceFleet {
			emr_AddInstanceFleet(cfg, client)
			return
		}
		if _emrAddInstanceGroups {
			emr_AddInstanceGroups(cfg, client)
			return
		}
		if _emrAddJobFlowSteps {
			emr_AddJobFlowSteps(cfg, client)
			return
		}
		if _emrAddTags {
			emr_AddTags(cfg, client)
			return
		}
		if _emrCancelSteps {
			emr_CancelSteps(cfg, client)
			return
		}
		if _emrCreatePersistentAppUI {
			emr_CreatePersistentAppUI(cfg, client)
			return
		}
		if _emrCreateSecurityConfiguration {
			emr_CreateSecurityConfiguration(cfg, client)
			return
		}
		if _emrCreateStudio {
			emr_CreateStudio(cfg, client)
			return
		}
		if _emrCreateStudioSessionMapping {
			emr_CreateStudioSessionMapping(cfg, client)
			return
		}
		if _emrDeleteSecurityConfiguration {
			emr_DeleteSecurityConfiguration(cfg, client)
			return
		}
		if _emrDeleteStudio {
			emr_DeleteStudio(cfg, client)
			return
		}
		if _emrDeleteStudioSessionMapping {
			emr_DeleteStudioSessionMapping(cfg, client)
			return
		}
		if _emrDescribeCluster {
			emr_DescribeCluster(cfg, client)
			return
		}
		if _emrDescribeJobFlows {
			emr_DescribeJobFlows(cfg, client)
			return
		}
		if _emrDescribeNotebookExecution {
			emr_DescribeNotebookExecution(cfg, client)
			return
		}
		if _emrDescribePersistentAppUI {
			emr_DescribePersistentAppUI(cfg, client)
			return
		}
		if _emrDescribeReleaseLabel {
			emr_DescribeReleaseLabel(cfg, client)
			return
		}
		if _emrDescribeSecurityConfiguration {
			emr_DescribeSecurityConfiguration(cfg, client)
			return
		}
		if _emrDescribeStep {
			emr_DescribeStep(cfg, client)
			return
		}
		if _emrDescribeStudio {
			emr_DescribeStudio(cfg, client)
			return
		}
		if _emrGetAutoTerminationPolicy {
			emr_GetAutoTerminationPolicy(cfg, client)
			return
		}
		if _emrGetBlockPublicAccessConfiguration {
			emr_GetBlockPublicAccessConfiguration(cfg, client)
			return
		}
		if _emrGetClusterSessionCredentials {
			emr_GetClusterSessionCredentials(cfg, client)
			return
		}
		if _emrGetManagedScalingPolicy {
			emr_GetManagedScalingPolicy(cfg, client)
			return
		}
		if _emrGetOnClusterAppUIPresignedURL {
			emr_GetOnClusterAppUIPresignedURL(cfg, client)
			return
		}
		if _emrGetPersistentAppUIPresignedURL {
			emr_GetPersistentAppUIPresignedURL(cfg, client)
			return
		}
		if _emrGetStudioSessionMapping {
			emr_GetStudioSessionMapping(cfg, client)
			return
		}
		if _emrListBootstrapActions {
			emr_ListBootstrapActions(cfg, client)
			return
		}
		if _emrListClusters {
			emr_ListClusters(cfg, client)
			return
		}
		if _emrListInstanceFleets {
			emr_ListInstanceFleets(cfg, client)
			return
		}
		if _emrListInstanceGroups {
			emr_ListInstanceGroups(cfg, client)
			return
		}
		if _emrListInstances {
			emr_ListInstances(cfg, client)
			return
		}
		if _emrListNotebookExecutions {
			emr_ListNotebookExecutions(cfg, client)
			return
		}
		if _emrListReleaseLabels {
			emr_ListReleaseLabels(cfg, client)
			return
		}
		if _emrListSecurityConfigurations {
			emr_ListSecurityConfigurations(cfg, client)
			return
		}
		if _emrListSteps {
			emr_ListSteps(cfg, client)
			return
		}
		if _emrListStudioSessionMappings {
			emr_ListStudioSessionMappings(cfg, client)
			return
		}
		if _emrListStudios {
			emr_ListStudios(cfg, client)
			return
		}
		if _emrListSupportedInstanceTypes {
			emr_ListSupportedInstanceTypes(cfg, client)
			return
		}
		if _emrModifyCluster {
			emr_ModifyCluster(cfg, client)
			return
		}
		if _emrModifyInstanceFleet {
			emr_ModifyInstanceFleet(cfg, client)
			return
		}
		if _emrModifyInstanceGroups {
			emr_ModifyInstanceGroups(cfg, client)
			return
		}
		if _emrPutAutoScalingPolicy {
			emr_PutAutoScalingPolicy(cfg, client)
			return
		}
		if _emrPutAutoTerminationPolicy {
			emr_PutAutoTerminationPolicy(cfg, client)
			return
		}
		if _emrPutBlockPublicAccessConfiguration {
			emr_PutBlockPublicAccessConfiguration(cfg, client)
			return
		}
		if _emrPutManagedScalingPolicy {
			emr_PutManagedScalingPolicy(cfg, client)
			return
		}
		if _emrRemoveAutoScalingPolicy {
			emr_RemoveAutoScalingPolicy(cfg, client)
			return
		}
		if _emrRemoveAutoTerminationPolicy {
			emr_RemoveAutoTerminationPolicy(cfg, client)
			return
		}
		if _emrRemoveManagedScalingPolicy {
			emr_RemoveManagedScalingPolicy(cfg, client)
			return
		}
		if _emrRemoveTags {
			emr_RemoveTags(cfg, client)
			return
		}
		if _emrRunJobFlow {
			emr_RunJobFlow(cfg, client)
			return
		}
		if _emrSetKeepJobFlowAliveWhenNoSteps {
			emr_SetKeepJobFlowAliveWhenNoSteps(cfg, client)
			return
		}
		if _emrSetTerminationProtection {
			emr_SetTerminationProtection(cfg, client)
			return
		}
		if _emrSetUnhealthyNodeReplacement {
			emr_SetUnhealthyNodeReplacement(cfg, client)
			return
		}
		if _emrSetVisibleToAllUsers {
			emr_SetVisibleToAllUsers(cfg, client)
			return
		}
		if _emrStartNotebookExecution {
			emr_StartNotebookExecution(cfg, client)
			return
		}
		if _emrStopNotebookExecution {
			emr_StopNotebookExecution(cfg, client)
			return
		}
		if _emrTerminateJobFlows {
			emr_TerminateJobFlows(cfg, client)
			return
		}
		if _emrUpdateStudio {
			emr_UpdateStudio(cfg, client)
			return
		}
		if _emrUpdateStudioSessionMapping {
			emr_UpdateStudioSessionMapping(cfg, client)
			return
		}

	},
}

var (
	_emrAddInstanceFleet                  bool
	_emrAddInstanceGroups                 bool
	_emrAddJobFlowSteps                   bool
	_emrAddTags                           bool
	_emrCancelSteps                       bool
	_emrCreatePersistentAppUI             bool
	_emrCreateSecurityConfiguration       bool
	_emrCreateStudio                      bool
	_emrCreateStudioSessionMapping        bool
	_emrDeleteSecurityConfiguration       bool
	_emrDeleteStudio                      bool
	_emrDeleteStudioSessionMapping        bool
	_emrDescribeCluster                   bool
	_emrDescribeJobFlows                  bool
	_emrDescribeNotebookExecution         bool
	_emrDescribePersistentAppUI           bool
	_emrDescribeReleaseLabel              bool
	_emrDescribeSecurityConfiguration     bool
	_emrDescribeStep                      bool
	_emrDescribeStudio                    bool
	_emrGetAutoTerminationPolicy          bool
	_emrGetBlockPublicAccessConfiguration bool
	_emrGetClusterSessionCredentials      bool
	_emrGetManagedScalingPolicy           bool
	_emrGetOnClusterAppUIPresignedURL     bool
	_emrGetPersistentAppUIPresignedURL    bool
	_emrGetStudioSessionMapping           bool
	_emrListBootstrapActions              bool
	_emrListClusters                      bool
	_emrListInstanceFleets                bool
	_emrListInstanceGroups                bool
	_emrListInstances                     bool
	_emrListNotebookExecutions            bool
	_emrListReleaseLabels                 bool
	_emrListSecurityConfigurations        bool
	_emrListSteps                         bool
	_emrListStudioSessionMappings         bool
	_emrListStudios                       bool
	_emrListSupportedInstanceTypes        bool
	_emrModifyCluster                     bool
	_emrModifyInstanceFleet               bool
	_emrModifyInstanceGroups              bool
	_emrPutAutoScalingPolicy              bool
	_emrPutAutoTerminationPolicy          bool
	_emrPutBlockPublicAccessConfiguration bool
	_emrPutManagedScalingPolicy           bool
	_emrRemoveAutoScalingPolicy           bool
	_emrRemoveAutoTerminationPolicy       bool
	_emrRemoveManagedScalingPolicy        bool
	_emrRemoveTags                        bool
	_emrRunJobFlow                        bool
	_emrSetKeepJobFlowAliveWhenNoSteps    bool
	_emrSetTerminationProtection          bool
	_emrSetUnhealthyNodeReplacement       bool
	_emrSetVisibleToAllUsers              bool
	_emrStartNotebookExecution            bool
	_emrStopNotebookExecution             bool
	_emrTerminateJobFlows                 bool
	_emrUpdateStudio                      bool
	_emrUpdateStudioSessionMapping        bool

	_emrAdditionalInfo                    string
	_emrAmiVersion                        string
	_emrApplicationId                     string
	_emrApplications                      string
	_emrAuthMode                          string
	_emrAuthProxyCall                     string
	_emrAutoScalingPolicy                 string
	_emrAutoScalingRole                   string
	_emrAutoTerminationPolicy             string
	_emrBlockPublicAccessConfiguration    string
	_emrBootstrapActions                  string
	_emrClusterId                         string
	_emrClusterStates                     string
	_emrConfigurations                    string
	_emrCreatedAfter                      string
	_emrCreatedBefore                     string
	_emrCustomAmiId                       string
	_emrDefaultS3Location                 string
	_emrDescription                       string
	_emrDryRun                            string
	_emrEbsRootVolumeIops                 string
	_emrEbsRootVolumeSize                 string
	_emrEbsRootVolumeThroughput           string
	_emrEditorId                          string
	_emrEMRContainersConfig               string
	_emrEncryptionKeyArn                  string
	_emrEngineSecurityGroupId             string
	_emrEnvironmentVariables              string
	_emrExecutionEngine                   string
	_emrExecutionEngineId                 string
	_emrExecutionRoleArn                  string
	_emrExtendedSupport                   string
	_emrFilters                           string
	_emrFrom                              string
	_emrIdcInstanceArn                    string
	_emrIdcUserAssignment                 string
	_emrIdentityId                        string
	_emrIdentityName                      string
	_emrIdentityType                      string
	_emrIdpAuthUrl                        string
	_emrIdpRelayStateParameterName        string
	_emrInstanceFleet                     string
	_emrInstanceFleetId                   string
	_emrInstanceFleetType                 string
	_emrInstanceGroupId                   string
	_emrInstanceGroupTypes                string
	_emrInstanceGroups                    string
	_emrInstanceStates                    string
	_emrInstances                         string
	_emrJobFlowId                         string
	_emrJobFlowIds                        []string
	_emrJobFlowRole                       string
	_emrJobFlowStates                     string
	_emrKeepJobFlowAliveWhenNoSteps       string
	_emrKerberosAttributes                string
	_emrLogEncryptionKmsKeyId             string
	_emrLogUri                            string
	_emrManagedScalingPolicy              string
	_emrMarker                            string
	_emrMaxResults                        string
	_emrMonitoringConfiguration           string
	_emrName                              string
	_emrNewSupportedProducts              string
	_emrNextToken                         string
	_emrNotebookExecutionId               string
	_emrNotebookExecutionName             string
	_emrNotebookInstanceSecurityGroupId   string
	_emrNotebookParams                    string
	_emrNotebookS3Location                string
	_emrOnClusterAppUIType                string
	_emrOSReleaseLabel                    string
	_emrOutputNotebookFormat              string
	_emrOutputNotebookS3Location          string
	_emrPersistentAppUIType               string
	_emrPersistentAppUIId                 string
	_emrPlacementGroupConfigs             string
	_emrProfilerType                      string
	_emrRelativePath                      string
	_emrReleaseLabel                      string
	_emrRepoUpgradeOnBoot                 string
	_emrResourceId                        string
	_emrScaleDownBehavior                 string
	_emrSecurityConfiguration             string
	_emrServiceRole                       string
	_emrSessionPolicyArn                  string
	_emrStatus                            string
	_emrStepCancellationOption            string
	_emrStepConcurrencyLevel              string
	_emrStepId                            string
	_emrStepIds                           []string
	_emrStepStates                        string
	_emrSteps                             string
	_emrStudioId                          string
	_emrSubnetIds                         []string
	_emrSupportedProducts                 []string
	_emrTagKeys                           []string
	_emrTags                              string
	_emrTargetResourceArn                 string
	_emrTerminationProtected              string
	_emrTo                                string
	_emrTrustedIdentityPropagationEnabled string
	_emrUnhealthyNodeReplacement          string
	_emrUserRole                          string
	_emrVisibleToAllUsers                 string
	_emrVpcId                             string
	_emrWorkspaceSecurityGroupId          string
	_emrXReferer                          string
)

// Adds an instance fleet to a running cluster.
// The instance fleet configuration is available only in Amazon EMR releases 4.8.0
// and later, excluding 5.0.x.
func emr_AddInstanceFleet(cfg aws.Config, client *emr.Client) {
	input := &emr.AddInstanceFleetInput{
		// ClusterId: *string, // Required
		// InstanceFleet: *types.InstanceFleetConfig, // Required
	}

	if len(_emrClusterId) > 0 {
		input.ClusterId = aws.String(_emrClusterId)
	}
	if len(_emrInstanceFleet) > 0 {
		if err := assignInputField(input, "InstanceFleet", _emrInstanceFleet); err != nil {
			log.Errorf("invalid --instance-fleet: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddInstanceFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more instance groups to a running cluster.
func emr_AddInstanceGroups(cfg aws.Config, client *emr.Client) {
	input := &emr.AddInstanceGroupsInput{
		// InstanceGroups: []types.InstanceGroupConfig, // Required
		// JobFlowId: *string, // Required
	}

	if len(_emrInstanceGroups) > 0 {
		if err := assignInputField(input, "InstanceGroups", _emrInstanceGroups); err != nil {
			log.Errorf("invalid --instance-groups: %s", err.Error())
			return
		}
	}
	if len(_emrJobFlowId) > 0 {
		input.JobFlowId = aws.String(_emrJobFlowId)
	}

	if resp, err := client.AddInstanceGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// AddJobFlowSteps adds new steps to a running cluster. A maximum of 256 steps are
// allowed in each job flow.
//
// If your cluster is long-running (such as a Hive data warehouse) or complex, you
// may require more than 256 steps to process your data. You can bypass the
// 256-step limitation in various ways, including using SSH to connect to the
// master node and submitting queries directly to the software running on the
// master node, such as Hive and Hadoop.
//
// A step specifies the location of a JAR file stored either on the master node of
// the cluster or in Amazon S3. Each step is performed by the main function of the
// main class of the JAR file. The main class can be specified either in the
// manifest of the JAR or by using the MainFunction parameter of the step.
//
// Amazon EMR executes each step in the order listed. For a step to be considered
// complete, the main function must exit with a zero exit code and all Hadoop jobs
// started while the step was running must have completed and run successfully.
//
// You can only add steps to a cluster that is in one of the following states:
// STARTING, BOOTSTRAPPING, RUNNING, or WAITING.
//
// The string values passed into HadoopJarStep object cannot exceed a total of
// 10240 characters.
func emr_AddJobFlowSteps(cfg aws.Config, client *emr.Client) {
	input := &emr.AddJobFlowStepsInput{
		// JobFlowId: *string, // Required
		// Steps: []types.StepConfig, // Required
	}

	if len(_emrJobFlowId) > 0 {
		input.JobFlowId = aws.String(_emrJobFlowId)
	}
	if len(_emrSteps) > 0 {
		if err := assignInputField(input, "Steps", _emrSteps); err != nil {
			log.Errorf("invalid --steps: %s", err.Error())
			return
		}
	}
	if len(_emrExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_emrExecutionRoleArn)
	}

	if resp, err := client.AddJobFlowSteps(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds tags to an Amazon EMR resource, such as a cluster or an Amazon EMR Studio.
// Tags make it easier to associate resources in various ways, such as grouping
// clusters to track your Amazon EMR resource allocation costs. For more
// information, see [Tag Clusters].
//
// [Tag Clusters]: https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-plan-tags.html
func emr_AddTags(cfg aws.Config, client *emr.Client) {
	input := &emr.AddTagsInput{
		// ResourceId: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_emrResourceId) > 0 {
		input.ResourceId = aws.String(_emrResourceId)
	}
	if len(_emrTags) > 0 {
		if err := assignInputField(input, "Tags", _emrTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a pending step or steps in a running cluster. Available only in Amazon
// EMR versions 4.8.0 and later, excluding version 5.0.0. A maximum of 256 steps
// are allowed in each CancelSteps request. CancelSteps is idempotent but
// asynchronous; it does not guarantee that a step will be canceled, even if the
// request is successfully submitted. When you use Amazon EMR releases 5.28.0 and
// later, you can cancel steps that are in a PENDING or RUNNING state. In earlier
// versions of Amazon EMR, you can only cancel steps that are in a PENDING state.
func emr_CancelSteps(cfg aws.Config, client *emr.Client) {
	input := &emr.CancelStepsInput{
		// ClusterId: *string, // Required
		// StepIds: []string, // Required
	}

	if len(_emrClusterId) > 0 {
		input.ClusterId = aws.String(_emrClusterId)
	}
	if len(_emrStepIds) > 0 {
		input.StepIds = append([]string(nil), _emrStepIds...)
	}
	if len(_emrStepCancellationOption) > 0 {
		if err := assignInputField(input, "StepCancellationOption", _emrStepCancellationOption); err != nil {
			log.Errorf("invalid --step-cancellation-option: %s", err.Error())
			return
		}
	}

	if resp, err := client.CancelSteps(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a persistent application user interface.
func emr_CreatePersistentAppUI(cfg aws.Config, client *emr.Client) {
	input := &emr.CreatePersistentAppUIInput{
		// TargetResourceArn: *string, // Required
	}

	if len(_emrTargetResourceArn) > 0 {
		input.TargetResourceArn = aws.String(_emrTargetResourceArn)
	}
	if len(_emrEMRContainersConfig) > 0 {
		if err := assignInputField(input, "EMRContainersConfig", _emrEMRContainersConfig); err != nil {
			log.Errorf("invalid --emr-containers-config: %s", err.Error())
			return
		}
	}
	if len(_emrProfilerType) > 0 {
		if err := assignInputField(input, "ProfilerType", _emrProfilerType); err != nil {
			log.Errorf("invalid --profiler-type: %s", err.Error())
			return
		}
	}
	if len(_emrTags) > 0 {
		if err := assignInputField(input, "Tags", _emrTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_emrXReferer) > 0 {
		input.XReferer = aws.String(_emrXReferer)
	}

	if resp, err := client.CreatePersistentAppUI(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a security configuration, which is stored in the service and can be
// specified when a cluster is created.
func emr_CreateSecurityConfiguration(cfg aws.Config, client *emr.Client) {
	input := &emr.CreateSecurityConfigurationInput{
		// Name: *string, // Required
		// SecurityConfiguration: *string, // Required
	}

	if len(_emrName) > 0 {
		input.Name = aws.String(_emrName)
	}
	if len(_emrSecurityConfiguration) > 0 {
		input.SecurityConfiguration = aws.String(_emrSecurityConfiguration)
	}

	if resp, err := client.CreateSecurityConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Amazon EMR Studio.
func emr_CreateStudio(cfg aws.Config, client *emr.Client) {
	input := &emr.CreateStudioInput{
		// AuthMode: types.AuthMode, // Required
		// DefaultS3Location: *string, // Required
		// EngineSecurityGroupId: *string, // Required
		// Name: *string, // Required
		// ServiceRole: *string, // Required
		// SubnetIds: []string, // Required
		// VpcId: *string, // Required
		// WorkspaceSecurityGroupId: *string, // Required
	}

	if len(_emrAuthMode) > 0 {
		if err := assignInputField(input, "AuthMode", _emrAuthMode); err != nil {
			log.Errorf("invalid --auth-mode: %s", err.Error())
			return
		}
	}
	if len(_emrDefaultS3Location) > 0 {
		input.DefaultS3Location = aws.String(_emrDefaultS3Location)
	}
	if len(_emrEngineSecurityGroupId) > 0 {
		input.EngineSecurityGroupId = aws.String(_emrEngineSecurityGroupId)
	}
	if len(_emrName) > 0 {
		input.Name = aws.String(_emrName)
	}
	if len(_emrServiceRole) > 0 {
		input.ServiceRole = aws.String(_emrServiceRole)
	}
	if len(_emrSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _emrSubnetIds...)
	}
	if len(_emrVpcId) > 0 {
		input.VpcId = aws.String(_emrVpcId)
	}
	if len(_emrWorkspaceSecurityGroupId) > 0 {
		input.WorkspaceSecurityGroupId = aws.String(_emrWorkspaceSecurityGroupId)
	}
	if len(_emrDescription) > 0 {
		input.Description = aws.String(_emrDescription)
	}
	if len(_emrEncryptionKeyArn) > 0 {
		input.EncryptionKeyArn = aws.String(_emrEncryptionKeyArn)
	}
	if len(_emrIdcInstanceArn) > 0 {
		input.IdcInstanceArn = aws.String(_emrIdcInstanceArn)
	}
	if len(_emrIdcUserAssignment) > 0 {
		if err := assignInputField(input, "IdcUserAssignment", _emrIdcUserAssignment); err != nil {
			log.Errorf("invalid --idc-user-assignment: %s", err.Error())
			return
		}
	}
	if len(_emrIdpAuthUrl) > 0 {
		input.IdpAuthUrl = aws.String(_emrIdpAuthUrl)
	}
	if len(_emrIdpRelayStateParameterName) > 0 {
		input.IdpRelayStateParameterName = aws.String(_emrIdpRelayStateParameterName)
	}
	if len(_emrTags) > 0 {
		if err := assignInputField(input, "Tags", _emrTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_emrTrustedIdentityPropagationEnabled) > 0 {
		if err := assignInputField(input, "TrustedIdentityPropagationEnabled", _emrTrustedIdentityPropagationEnabled); err != nil {
			log.Errorf("invalid --trusted-identity-propagation-enabled: %s", err.Error())
			return
		}
	}
	if len(_emrUserRole) > 0 {
		input.UserRole = aws.String(_emrUserRole)
	}

	if resp, err := client.CreateStudio(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Maps a user or group to the Amazon EMR Studio specified by StudioId , and
// applies a session policy to refine Studio permissions for that user or group.
// Use CreateStudioSessionMapping to assign users to a Studio when you use IAM
// Identity Center authentication. For instructions on how to assign users to a
// Studio when you use IAM authentication, see [Assign a user or group to your EMR Studio].
//
// [Assign a user or group to your EMR Studio]: https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-studio-manage-users.html#emr-studio-assign-users-groups
func emr_CreateStudioSessionMapping(cfg aws.Config, client *emr.Client) {
	input := &emr.CreateStudioSessionMappingInput{
		// IdentityType: types.IdentityType, // Required
		// SessionPolicyArn: *string, // Required
		// StudioId: *string, // Required
	}

	if len(_emrIdentityType) > 0 {
		if err := assignInputField(input, "IdentityType", _emrIdentityType); err != nil {
			log.Errorf("invalid --identity-type: %s", err.Error())
			return
		}
	}
	if len(_emrSessionPolicyArn) > 0 {
		input.SessionPolicyArn = aws.String(_emrSessionPolicyArn)
	}
	if len(_emrStudioId) > 0 {
		input.StudioId = aws.String(_emrStudioId)
	}
	if len(_emrIdentityId) > 0 {
		input.IdentityId = aws.String(_emrIdentityId)
	}
	if len(_emrIdentityName) > 0 {
		input.IdentityName = aws.String(_emrIdentityName)
	}

	if resp, err := client.CreateStudioSessionMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a security configuration.
func emr_DeleteSecurityConfiguration(cfg aws.Config, client *emr.Client) {
	input := &emr.DeleteSecurityConfigurationInput{
		// Name: *string, // Required
	}

	if len(_emrName) > 0 {
		input.Name = aws.String(_emrName)
	}

	if resp, err := client.DeleteSecurityConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes an Amazon EMR Studio from the Studio metadata store.
func emr_DeleteStudio(cfg aws.Config, client *emr.Client) {
	input := &emr.DeleteStudioInput{
		// StudioId: *string, // Required
	}

	if len(_emrStudioId) > 0 {
		input.StudioId = aws.String(_emrStudioId)
	}

	if resp, err := client.DeleteStudio(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a user or group from an Amazon EMR Studio.
func emr_DeleteStudioSessionMapping(cfg aws.Config, client *emr.Client) {
	input := &emr.DeleteStudioSessionMappingInput{
		// IdentityType: types.IdentityType, // Required
		// StudioId: *string, // Required
	}

	if len(_emrIdentityType) > 0 {
		if err := assignInputField(input, "IdentityType", _emrIdentityType); err != nil {
			log.Errorf("invalid --identity-type: %s", err.Error())
			return
		}
	}
	if len(_emrStudioId) > 0 {
		input.StudioId = aws.String(_emrStudioId)
	}
	if len(_emrIdentityId) > 0 {
		input.IdentityId = aws.String(_emrIdentityId)
	}
	if len(_emrIdentityName) > 0 {
		input.IdentityName = aws.String(_emrIdentityName)
	}

	if resp, err := client.DeleteStudioSessionMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides cluster-level details including status, hardware and software
// configuration, VPC settings, and so on.
func emr_DescribeCluster(cfg aws.Config, client *emr.Client) {
	input := &emr.DescribeClusterInput{
		// ClusterId: *string, // Required
	}

	if len(_emrClusterId) > 0 {
		input.ClusterId = aws.String(_emrClusterId)
	}

	if resp, err := client.DescribeCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is no longer supported and will eventually be removed. We recommend
// you use ListClusters, DescribeCluster, ListSteps, ListInstanceGroups and ListBootstrapActions instead.
//
// DescribeJobFlows returns a list of job flows that match all of the supplied
// parameters. The parameters can include a list of job flow IDs, job flow states,
// and restrictions on job flow creation date and time.
//
// Regardless of supplied parameters, only job flows created within the last two
// months are returned.
//
// If no parameters are supplied, then job flows matching either of the following
// criteria are returned:
//
// - Job flows created and completed in the last two weeks
//
// - Job flows created within the last two months that are in one of the
// following states: RUNNING , WAITING , SHUTTING_DOWN , STARTING
//
// Amazon EMR can return a maximum of 512 job flow descriptions.
//
// Deprecated: This operation has been deprecated.
func emr_DescribeJobFlows(cfg aws.Config, client *emr.Client) {
	input := &emr.DescribeJobFlowsInput{}

	if len(_emrCreatedAfter) > 0 {
		if err := assignInputField(input, "CreatedAfter", _emrCreatedAfter); err != nil {
			log.Errorf("invalid --created-after: %s", err.Error())
			return
		}
	}
	if len(_emrCreatedBefore) > 0 {
		if err := assignInputField(input, "CreatedBefore", _emrCreatedBefore); err != nil {
			log.Errorf("invalid --created-before: %s", err.Error())
			return
		}
	}
	if len(_emrJobFlowIds) > 0 {
		input.JobFlowIds = append([]string(nil), _emrJobFlowIds...)
	}
	if len(_emrJobFlowStates) > 0 {
		if err := assignInputField(input, "JobFlowStates", _emrJobFlowStates); err != nil {
			log.Errorf("invalid --job-flow-states: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeJobFlows(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides details of a notebook execution.
func emr_DescribeNotebookExecution(cfg aws.Config, client *emr.Client) {
	input := &emr.DescribeNotebookExecutionInput{
		// NotebookExecutionId: *string, // Required
	}

	if len(_emrNotebookExecutionId) > 0 {
		input.NotebookExecutionId = aws.String(_emrNotebookExecutionId)
	}

	if resp, err := client.DescribeNotebookExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a persistent application user interface.
func emr_DescribePersistentAppUI(cfg aws.Config, client *emr.Client) {
	input := &emr.DescribePersistentAppUIInput{
		// PersistentAppUIId: *string, // Required
	}

	if len(_emrPersistentAppUIId) > 0 {
		input.PersistentAppUIId = aws.String(_emrPersistentAppUIId)
	}

	if resp, err := client.DescribePersistentAppUI(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides Amazon EMR release label details, such as the releases available the
// Region where the API request is run, and the available applications for a
// specific Amazon EMR release label. Can also list Amazon EMR releases that
// support a specified version of Spark.
func emr_DescribeReleaseLabel(cfg aws.Config, client *emr.Client) {
	input := &emr.DescribeReleaseLabelInput{}

	if len(_emrMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _emrMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_emrNextToken) > 0 {
		input.NextToken = aws.String(_emrNextToken)
	}
	if len(_emrReleaseLabel) > 0 {
		input.ReleaseLabel = aws.String(_emrReleaseLabel)
	}

	if resp, err := client.DescribeReleaseLabel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides the details of a security configuration by returning the configuration
// JSON.
func emr_DescribeSecurityConfiguration(cfg aws.Config, client *emr.Client) {
	input := &emr.DescribeSecurityConfigurationInput{
		// Name: *string, // Required
	}

	if len(_emrName) > 0 {
		input.Name = aws.String(_emrName)
	}

	if resp, err := client.DescribeSecurityConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides more detail about the cluster step.
func emr_DescribeStep(cfg aws.Config, client *emr.Client) {
	input := &emr.DescribeStepInput{
		// ClusterId: *string, // Required
		// StepId: *string, // Required
	}

	if len(_emrClusterId) > 0 {
		input.ClusterId = aws.String(_emrClusterId)
	}
	if len(_emrStepId) > 0 {
		input.StepId = aws.String(_emrStepId)
	}

	if resp, err := client.DescribeStep(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details for the specified Amazon EMR Studio including ID, Name, VPC,
// Studio access URL, and so on.
func emr_DescribeStudio(cfg aws.Config, client *emr.Client) {
	input := &emr.DescribeStudioInput{
		// StudioId: *string, // Required
	}

	if len(_emrStudioId) > 0 {
		input.StudioId = aws.String(_emrStudioId)
	}

	if resp, err := client.DescribeStudio(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the auto-termination policy for an Amazon EMR cluster.
func emr_GetAutoTerminationPolicy(cfg aws.Config, client *emr.Client) {
	input := &emr.GetAutoTerminationPolicyInput{
		// ClusterId: *string, // Required
	}

	if len(_emrClusterId) > 0 {
		input.ClusterId = aws.String(_emrClusterId)
	}

	if resp, err := client.GetAutoTerminationPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the Amazon EMR block public access configuration for your Amazon Web
// Services account in the current Region. For more information see [Configure Block Public Access for Amazon EMR]in the Amazon
// EMR Management Guide.
//
// [Configure Block Public Access for Amazon EMR]: https://docs.aws.amazon.com/emr/latest/ManagementGuide/configure-block-public-access.html
func emr_GetBlockPublicAccessConfiguration(cfg aws.Config, client *emr.Client) {
	input := &emr.GetBlockPublicAccessConfigurationInput{}

	if resp, err := client.GetBlockPublicAccessConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides temporary, HTTP basic credentials that are associated with a given
// runtime IAM role and used by a cluster with fine-grained access control
// activated. You can use these credentials to connect to cluster endpoints that
// support username and password authentication.
func emr_GetClusterSessionCredentials(cfg aws.Config, client *emr.Client) {
	input := &emr.GetClusterSessionCredentialsInput{
		// ClusterId: *string, // Required
	}

	if len(_emrClusterId) > 0 {
		input.ClusterId = aws.String(_emrClusterId)
	}
	if len(_emrExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_emrExecutionRoleArn)
	}

	if resp, err := client.GetClusterSessionCredentials(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Fetches the attached managed scaling policy for an Amazon EMR cluster.
func emr_GetManagedScalingPolicy(cfg aws.Config, client *emr.Client) {
	input := &emr.GetManagedScalingPolicyInput{
		// ClusterId: *string, // Required
	}

	if len(_emrClusterId) > 0 {
		input.ClusterId = aws.String(_emrClusterId)
	}

	if resp, err := client.GetManagedScalingPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The presigned URL properties for the cluster's application user interface.
func emr_GetOnClusterAppUIPresignedURL(cfg aws.Config, client *emr.Client) {
	input := &emr.GetOnClusterAppUIPresignedURLInput{
		// ClusterId: *string, // Required
	}

	if len(_emrClusterId) > 0 {
		input.ClusterId = aws.String(_emrClusterId)
	}
	if len(_emrApplicationId) > 0 {
		input.ApplicationId = aws.String(_emrApplicationId)
	}
	if len(_emrDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _emrDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_emrExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_emrExecutionRoleArn)
	}
	if len(_emrOnClusterAppUIType) > 0 {
		if err := assignInputField(input, "OnClusterAppUIType", _emrOnClusterAppUIType); err != nil {
			log.Errorf("invalid --on-cluster-app-ui-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetOnClusterAppUIPresignedURL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The presigned URL properties for the cluster's application user interface.
func emr_GetPersistentAppUIPresignedURL(cfg aws.Config, client *emr.Client) {
	input := &emr.GetPersistentAppUIPresignedURLInput{
		// PersistentAppUIId: *string, // Required
	}

	if len(_emrPersistentAppUIId) > 0 {
		input.PersistentAppUIId = aws.String(_emrPersistentAppUIId)
	}
	if len(_emrApplicationId) > 0 {
		input.ApplicationId = aws.String(_emrApplicationId)
	}
	if len(_emrAuthProxyCall) > 0 {
		if err := assignInputField(input, "AuthProxyCall", _emrAuthProxyCall); err != nil {
			log.Errorf("invalid --auth-proxy-call: %s", err.Error())
			return
		}
	}
	if len(_emrExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_emrExecutionRoleArn)
	}
	if len(_emrPersistentAppUIType) > 0 {
		if err := assignInputField(input, "PersistentAppUIType", _emrPersistentAppUIType); err != nil {
			log.Errorf("invalid --persistent-app-ui-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetPersistentAppUIPresignedURL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Fetches mapping details for the specified Amazon EMR Studio and identity (user
// or group).
func emr_GetStudioSessionMapping(cfg aws.Config, client *emr.Client) {
	input := &emr.GetStudioSessionMappingInput{
		// IdentityType: types.IdentityType, // Required
		// StudioId: *string, // Required
	}

	if len(_emrIdentityType) > 0 {
		if err := assignInputField(input, "IdentityType", _emrIdentityType); err != nil {
			log.Errorf("invalid --identity-type: %s", err.Error())
			return
		}
	}
	if len(_emrStudioId) > 0 {
		input.StudioId = aws.String(_emrStudioId)
	}
	if len(_emrIdentityId) > 0 {
		input.IdentityId = aws.String(_emrIdentityId)
	}
	if len(_emrIdentityName) > 0 {
		input.IdentityName = aws.String(_emrIdentityName)
	}

	if resp, err := client.GetStudioSessionMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about the bootstrap actions associated with a cluster.
func emr_ListBootstrapActions(cfg aws.Config, client *emr.Client) {
	input := &emr.ListBootstrapActionsInput{
		// ClusterId: *string, // Required
	}

	if len(_emrClusterId) > 0 {
		input.ClusterId = aws.String(_emrClusterId)
	}
	if len(_emrMarker) > 0 {
		input.Marker = aws.String(_emrMarker)
	}

	if disablePaginator() {
		if resp, err := client.ListBootstrapActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*emr.ListBootstrapActionsOutput
	p := emr.NewListBootstrapActionsPaginator(client, input)
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

// Provides the status of all clusters visible to this Amazon Web Services
// account. Allows you to filter the list of clusters based on certain criteria;
// for example, filtering by cluster creation date and time or by status. This call
// returns a maximum of 50 clusters in unsorted order per call, but returns a
// marker to track the paging of the cluster list across multiple ListClusters
// calls.
func emr_ListClusters(cfg aws.Config, client *emr.Client) {
	input := &emr.ListClustersInput{}

	if len(_emrClusterStates) > 0 {
		if err := assignInputField(input, "ClusterStates", _emrClusterStates); err != nil {
			log.Errorf("invalid --cluster-states: %s", err.Error())
			return
		}
	}
	if len(_emrCreatedAfter) > 0 {
		if err := assignInputField(input, "CreatedAfter", _emrCreatedAfter); err != nil {
			log.Errorf("invalid --created-after: %s", err.Error())
			return
		}
	}
	if len(_emrCreatedBefore) > 0 {
		if err := assignInputField(input, "CreatedBefore", _emrCreatedBefore); err != nil {
			log.Errorf("invalid --created-before: %s", err.Error())
			return
		}
	}
	if len(_emrMarker) > 0 {
		input.Marker = aws.String(_emrMarker)
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

	var results []*emr.ListClustersOutput
	p := emr.NewListClustersPaginator(client, input)
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

// Lists all available details about the instance fleets in a cluster.
// The instance fleet configuration is available only in Amazon EMR releases 4.8.0
// and later, excluding 5.0.x versions.
func emr_ListInstanceFleets(cfg aws.Config, client *emr.Client) {
	input := &emr.ListInstanceFleetsInput{
		// ClusterId: *string, // Required
	}

	if len(_emrClusterId) > 0 {
		input.ClusterId = aws.String(_emrClusterId)
	}
	if len(_emrMarker) > 0 {
		input.Marker = aws.String(_emrMarker)
	}

	if disablePaginator() {
		if resp, err := client.ListInstanceFleets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*emr.ListInstanceFleetsOutput
	p := emr.NewListInstanceFleetsPaginator(client, input)
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

// Provides all available details about the instance groups in a cluster.
func emr_ListInstanceGroups(cfg aws.Config, client *emr.Client) {
	input := &emr.ListInstanceGroupsInput{
		// ClusterId: *string, // Required
	}

	if len(_emrClusterId) > 0 {
		input.ClusterId = aws.String(_emrClusterId)
	}
	if len(_emrMarker) > 0 {
		input.Marker = aws.String(_emrMarker)
	}

	if disablePaginator() {
		if resp, err := client.ListInstanceGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*emr.ListInstanceGroupsOutput
	p := emr.NewListInstanceGroupsPaginator(client, input)
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

// Provides information for all active Amazon EC2 instances and Amazon EC2
// instances terminated in the last 30 days, up to a maximum of 2,000. Amazon EC2
// instances in any of the following states are considered active:
// AWAITING_FULFILLMENT, PROVISIONING, BOOTSTRAPPING, RUNNING.
func emr_ListInstances(cfg aws.Config, client *emr.Client) {
	input := &emr.ListInstancesInput{
		// ClusterId: *string, // Required
	}

	if len(_emrClusterId) > 0 {
		input.ClusterId = aws.String(_emrClusterId)
	}
	if len(_emrInstanceFleetId) > 0 {
		input.InstanceFleetId = aws.String(_emrInstanceFleetId)
	}
	if len(_emrInstanceFleetType) > 0 {
		if err := assignInputField(input, "InstanceFleetType", _emrInstanceFleetType); err != nil {
			log.Errorf("invalid --instance-fleet-type: %s", err.Error())
			return
		}
	}
	if len(_emrInstanceGroupId) > 0 {
		input.InstanceGroupId = aws.String(_emrInstanceGroupId)
	}
	if len(_emrInstanceGroupTypes) > 0 {
		if err := assignInputField(input, "InstanceGroupTypes", _emrInstanceGroupTypes); err != nil {
			log.Errorf("invalid --instance-group-types: %s", err.Error())
			return
		}
	}
	if len(_emrInstanceStates) > 0 {
		if err := assignInputField(input, "InstanceStates", _emrInstanceStates); err != nil {
			log.Errorf("invalid --instance-states: %s", err.Error())
			return
		}
	}
	if len(_emrMarker) > 0 {
		input.Marker = aws.String(_emrMarker)
	}

	if disablePaginator() {
		if resp, err := client.ListInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*emr.ListInstancesOutput
	p := emr.NewListInstancesPaginator(client, input)
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

// Provides summaries of all notebook executions. You can filter the list based on
// multiple criteria such as status, time range, and editor id. Returns a maximum
// of 50 notebook executions and a marker to track the paging of a longer notebook
// execution list across multiple ListNotebookExecutions calls.
func emr_ListNotebookExecutions(cfg aws.Config, client *emr.Client) {
	input := &emr.ListNotebookExecutionsInput{}

	if len(_emrEditorId) > 0 {
		input.EditorId = aws.String(_emrEditorId)
	}
	if len(_emrExecutionEngineId) > 0 {
		input.ExecutionEngineId = aws.String(_emrExecutionEngineId)
	}
	if len(_emrFrom) > 0 {
		if err := assignInputField(input, "From", _emrFrom); err != nil {
			log.Errorf("invalid --from: %s", err.Error())
			return
		}
	}
	if len(_emrMarker) > 0 {
		input.Marker = aws.String(_emrMarker)
	}
	if len(_emrStatus) > 0 {
		if err := assignInputField(input, "Status", _emrStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_emrTo) > 0 {
		if err := assignInputField(input, "To", _emrTo); err != nil {
			log.Errorf("invalid --to: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListNotebookExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*emr.ListNotebookExecutionsOutput
	p := emr.NewListNotebookExecutionsPaginator(client, input)
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

// Retrieves release labels of Amazon EMR services in the Region where the API is
// called.
func emr_ListReleaseLabels(cfg aws.Config, client *emr.Client) {
	input := &emr.ListReleaseLabelsInput{}

	if len(_emrFilters) > 0 {
		if err := assignInputField(input, "Filters", _emrFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_emrMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _emrMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_emrNextToken) > 0 {
		input.NextToken = aws.String(_emrNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListReleaseLabels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*emr.ListReleaseLabelsOutput
	p := emr.NewListReleaseLabelsPaginator(client, input)
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

// Lists all the security configurations visible to this account, providing their
// creation dates and times, and their names. This call returns a maximum of 50
// clusters per call, but returns a marker to track the paging of the cluster list
// across multiple ListSecurityConfigurations calls.
func emr_ListSecurityConfigurations(cfg aws.Config, client *emr.Client) {
	input := &emr.ListSecurityConfigurationsInput{}

	if len(_emrMarker) > 0 {
		input.Marker = aws.String(_emrMarker)
	}

	if disablePaginator() {
		if resp, err := client.ListSecurityConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*emr.ListSecurityConfigurationsOutput
	p := emr.NewListSecurityConfigurationsPaginator(client, input)
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

// Provides a list of steps for the cluster in reverse order unless you specify
// stepIds with the request or filter by StepStates . You can specify a maximum of
// 10 stepIDs . The CLI automatically paginates results to return a list greater
// than 50 steps. To return more than 50 steps using the CLI, specify a Marker ,
// which is a pagination token that indicates the next set of steps to retrieve.
func emr_ListSteps(cfg aws.Config, client *emr.Client) {
	input := &emr.ListStepsInput{
		// ClusterId: *string, // Required
	}

	if len(_emrClusterId) > 0 {
		input.ClusterId = aws.String(_emrClusterId)
	}
	if len(_emrMarker) > 0 {
		input.Marker = aws.String(_emrMarker)
	}
	if len(_emrStepIds) > 0 {
		input.StepIds = append([]string(nil), _emrStepIds...)
	}
	if len(_emrStepStates) > 0 {
		if err := assignInputField(input, "StepStates", _emrStepStates); err != nil {
			log.Errorf("invalid --step-states: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSteps(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*emr.ListStepsOutput
	p := emr.NewListStepsPaginator(client, input)
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

// Returns a list of all user or group session mappings for the Amazon EMR Studio
// specified by StudioId .
func emr_ListStudioSessionMappings(cfg aws.Config, client *emr.Client) {
	input := &emr.ListStudioSessionMappingsInput{}

	if len(_emrIdentityType) > 0 {
		if err := assignInputField(input, "IdentityType", _emrIdentityType); err != nil {
			log.Errorf("invalid --identity-type: %s", err.Error())
			return
		}
	}
	if len(_emrMarker) > 0 {
		input.Marker = aws.String(_emrMarker)
	}
	if len(_emrStudioId) > 0 {
		input.StudioId = aws.String(_emrStudioId)
	}

	if disablePaginator() {
		if resp, err := client.ListStudioSessionMappings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*emr.ListStudioSessionMappingsOutput
	p := emr.NewListStudioSessionMappingsPaginator(client, input)
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

// Returns a list of all Amazon EMR Studios associated with the Amazon Web
// Services account. The list includes details such as ID, Studio Access URL, and
// creation time for each Studio.
func emr_ListStudios(cfg aws.Config, client *emr.Client) {
	input := &emr.ListStudiosInput{}

	if len(_emrMarker) > 0 {
		input.Marker = aws.String(_emrMarker)
	}

	if disablePaginator() {
		if resp, err := client.ListStudios(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*emr.ListStudiosOutput
	p := emr.NewListStudiosPaginator(client, input)
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

// A list of the instance types that Amazon EMR supports. You can filter the list
// by Amazon Web Services Region and Amazon EMR release.
func emr_ListSupportedInstanceTypes(cfg aws.Config, client *emr.Client) {
	input := &emr.ListSupportedInstanceTypesInput{
		// ReleaseLabel: *string, // Required
	}

	if len(_emrReleaseLabel) > 0 {
		input.ReleaseLabel = aws.String(_emrReleaseLabel)
	}
	if len(_emrMarker) > 0 {
		input.Marker = aws.String(_emrMarker)
	}

	if disablePaginator() {
		if resp, err := client.ListSupportedInstanceTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*emr.ListSupportedInstanceTypesOutput
	p := emr.NewListSupportedInstanceTypesPaginator(client, input)
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

// Modifies the number of steps that can be executed concurrently for the cluster
// specified using ClusterID.
func emr_ModifyCluster(cfg aws.Config, client *emr.Client) {
	input := &emr.ModifyClusterInput{
		// ClusterId: *string, // Required
	}

	if len(_emrClusterId) > 0 {
		input.ClusterId = aws.String(_emrClusterId)
	}
	if len(_emrExtendedSupport) > 0 {
		if err := assignInputField(input, "ExtendedSupport", _emrExtendedSupport); err != nil {
			log.Errorf("invalid --extended-support: %s", err.Error())
			return
		}
	}
	if len(_emrStepConcurrencyLevel) > 0 {
		if err := assignInputField(input, "StepConcurrencyLevel", _emrStepConcurrencyLevel); err != nil {
			log.Errorf("invalid --step-concurrency-level: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the target On-Demand and target Spot capacities for the instance fleet
// with the specified InstanceFleetID within the cluster specified using ClusterID.
// The call either succeeds or fails atomically.
//
// The instance fleet configuration is available only in Amazon EMR releases 4.8.0
// and later, excluding 5.0.x versions.
func emr_ModifyInstanceFleet(cfg aws.Config, client *emr.Client) {
	input := &emr.ModifyInstanceFleetInput{
		// ClusterId: *string, // Required
		// InstanceFleet: *types.InstanceFleetModifyConfig, // Required
	}

	if len(_emrClusterId) > 0 {
		input.ClusterId = aws.String(_emrClusterId)
	}
	if len(_emrInstanceFleet) > 0 {
		if err := assignInputField(input, "InstanceFleet", _emrInstanceFleet); err != nil {
			log.Errorf("invalid --instance-fleet: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyInstanceFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// ModifyInstanceGroups modifies the number of nodes and configuration settings of
// an instance group. The input parameters include the new target instance count
// for the group and the instance group ID. The call will either succeed or fail
// atomically.
func emr_ModifyInstanceGroups(cfg aws.Config, client *emr.Client) {
	input := &emr.ModifyInstanceGroupsInput{}

	if len(_emrClusterId) > 0 {
		input.ClusterId = aws.String(_emrClusterId)
	}
	if len(_emrInstanceGroups) > 0 {
		if err := assignInputField(input, "InstanceGroups", _emrInstanceGroups); err != nil {
			log.Errorf("invalid --instance-groups: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyInstanceGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates an automatic scaling policy for a core instance group or
// task instance group in an Amazon EMR cluster. The automatic scaling policy
// defines how an instance group dynamically adds and terminates Amazon EC2
// instances in response to the value of a CloudWatch metric.
func emr_PutAutoScalingPolicy(cfg aws.Config, client *emr.Client) {
	input := &emr.PutAutoScalingPolicyInput{
		// AutoScalingPolicy: *types.AutoScalingPolicy, // Required
		// ClusterId: *string, // Required
		// InstanceGroupId: *string, // Required
	}

	if len(_emrAutoScalingPolicy) > 0 {
		if err := assignInputField(input, "AutoScalingPolicy", _emrAutoScalingPolicy); err != nil {
			log.Errorf("invalid --auto-scaling-policy: %s", err.Error())
			return
		}
	}
	if len(_emrClusterId) > 0 {
		input.ClusterId = aws.String(_emrClusterId)
	}
	if len(_emrInstanceGroupId) > 0 {
		input.InstanceGroupId = aws.String(_emrInstanceGroupId)
	}

	if resp, err := client.PutAutoScalingPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Auto-termination is supported in Amazon EMR releases 5.30.0 and 6.1.0 and
// later. For more information, see [Using an auto-termination policy].
//
// Creates or updates an auto-termination policy for an Amazon EMR cluster. An
// auto-termination policy defines the amount of idle time in seconds after which a
// cluster automatically terminates. For alternative cluster termination options,
// see [Control cluster termination].
//
// [Using an auto-termination policy]: https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-auto-termination-policy.html
// [Control cluster termination]: https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-plan-termination.html
func emr_PutAutoTerminationPolicy(cfg aws.Config, client *emr.Client) {
	input := &emr.PutAutoTerminationPolicyInput{
		// ClusterId: *string, // Required
	}

	if len(_emrClusterId) > 0 {
		input.ClusterId = aws.String(_emrClusterId)
	}
	if len(_emrAutoTerminationPolicy) > 0 {
		if err := assignInputField(input, "AutoTerminationPolicy", _emrAutoTerminationPolicy); err != nil {
			log.Errorf("invalid --auto-termination-policy: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutAutoTerminationPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates an Amazon EMR block public access configuration for your
// Amazon Web Services account in the current Region. For more information see [Configure Block Public Access for Amazon EMR]in
// the Amazon EMR Management Guide.
//
// [Configure Block Public Access for Amazon EMR]: https://docs.aws.amazon.com/emr/latest/ManagementGuide/configure-block-public-access.html
func emr_PutBlockPublicAccessConfiguration(cfg aws.Config, client *emr.Client) {
	input := &emr.PutBlockPublicAccessConfigurationInput{
		// BlockPublicAccessConfiguration: *types.BlockPublicAccessConfiguration, // Required
	}

	if len(_emrBlockPublicAccessConfiguration) > 0 {
		if err := assignInputField(input, "BlockPublicAccessConfiguration", _emrBlockPublicAccessConfiguration); err != nil {
			log.Errorf("invalid --block-public-access-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutBlockPublicAccessConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a managed scaling policy for an Amazon EMR cluster. The
// managed scaling policy defines the limits for resources, such as Amazon EC2
// instances that can be added or terminated from a cluster. The policy only
// applies to the core and task nodes. The master node cannot be scaled after
// initial configuration.
func emr_PutManagedScalingPolicy(cfg aws.Config, client *emr.Client) {
	input := &emr.PutManagedScalingPolicyInput{
		// ClusterId: *string, // Required
		// ManagedScalingPolicy: *types.ManagedScalingPolicy, // Required
	}

	if len(_emrClusterId) > 0 {
		input.ClusterId = aws.String(_emrClusterId)
	}
	if len(_emrManagedScalingPolicy) > 0 {
		if err := assignInputField(input, "ManagedScalingPolicy", _emrManagedScalingPolicy); err != nil {
			log.Errorf("invalid --managed-scaling-policy: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutManagedScalingPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes an automatic scaling policy from a specified instance group within an
// Amazon EMR cluster.
func emr_RemoveAutoScalingPolicy(cfg aws.Config, client *emr.Client) {
	input := &emr.RemoveAutoScalingPolicyInput{
		// ClusterId: *string, // Required
		// InstanceGroupId: *string, // Required
	}

	if len(_emrClusterId) > 0 {
		input.ClusterId = aws.String(_emrClusterId)
	}
	if len(_emrInstanceGroupId) > 0 {
		input.InstanceGroupId = aws.String(_emrInstanceGroupId)
	}

	if resp, err := client.RemoveAutoScalingPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes an auto-termination policy from an Amazon EMR cluster.
func emr_RemoveAutoTerminationPolicy(cfg aws.Config, client *emr.Client) {
	input := &emr.RemoveAutoTerminationPolicyInput{
		// ClusterId: *string, // Required
	}

	if len(_emrClusterId) > 0 {
		input.ClusterId = aws.String(_emrClusterId)
	}

	if resp, err := client.RemoveAutoTerminationPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a managed scaling policy from a specified Amazon EMR cluster.
func emr_RemoveManagedScalingPolicy(cfg aws.Config, client *emr.Client) {
	input := &emr.RemoveManagedScalingPolicyInput{
		// ClusterId: *string, // Required
	}

	if len(_emrClusterId) > 0 {
		input.ClusterId = aws.String(_emrClusterId)
	}

	if resp, err := client.RemoveManagedScalingPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes tags from an Amazon EMR resource, such as a cluster or Amazon EMR
// Studio. Tags make it easier to associate resources in various ways, such as
// grouping clusters to track your Amazon EMR resource allocation costs. For more
// information, see [Tag Clusters].
//
// The following example removes the stack tag with value Prod from a cluster:
//
// [Tag Clusters]: https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-plan-tags.html
func emr_RemoveTags(cfg aws.Config, client *emr.Client) {
	input := &emr.RemoveTagsInput{
		// ResourceId: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_emrResourceId) > 0 {
		input.ResourceId = aws.String(_emrResourceId)
	}
	if len(_emrTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _emrTagKeys...)
	}

	if resp, err := client.RemoveTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// RunJobFlow creates and starts running a new cluster (job flow). The cluster
// runs the steps specified. After the steps complete, the cluster stops and the
// HDFS partition is lost. To prevent loss of data, configure the last step of the
// job flow to store results in Amazon S3. If the JobFlowInstancesConfigKeepJobFlowAliveWhenNoSteps
// parameter is set to TRUE , the cluster transitions to the WAITING state rather
// than shutting down after the steps have completed.
//
// For additional protection, you can set the JobFlowInstancesConfigTerminationProtected parameter to
// TRUE to lock the cluster and prevent it from being terminated by API call, user
// intervention, or in the event of a job flow error.
//
// A maximum of 256 steps are allowed in each job flow.
//
// If your cluster is long-running (such as a Hive data warehouse) or complex, you
// may require more than 256 steps to process your data. You can bypass the
// 256-step limitation in various ways, including using the SSH shell to connect to
// the master node and submitting queries directly to the software running on the
// master node, such as Hive and Hadoop.
//
// For long-running clusters, we recommend that you periodically store your
// results.
//
// The instance fleets configuration is available only in Amazon EMR releases
// 4.8.0 and later, excluding 5.0.x versions. The RunJobFlow request can contain
// InstanceFleets parameters or InstanceGroups parameters, but not both.
func emr_RunJobFlow(cfg aws.Config, client *emr.Client) {
	input := &emr.RunJobFlowInput{
		// Instances: *types.JobFlowInstancesConfig, // Required
		// Name: *string, // Required
	}

	if len(_emrInstances) > 0 {
		if err := assignInputField(input, "Instances", _emrInstances); err != nil {
			log.Errorf("invalid --instances: %s", err.Error())
			return
		}
	}
	if len(_emrName) > 0 {
		input.Name = aws.String(_emrName)
	}
	if len(_emrAdditionalInfo) > 0 {
		input.AdditionalInfo = aws.String(_emrAdditionalInfo)
	}
	if len(_emrAmiVersion) > 0 {
		input.AmiVersion = aws.String(_emrAmiVersion)
	}
	if len(_emrApplications) > 0 {
		if err := assignInputField(input, "Applications", _emrApplications); err != nil {
			log.Errorf("invalid --applications: %s", err.Error())
			return
		}
	}
	if len(_emrAutoScalingRole) > 0 {
		input.AutoScalingRole = aws.String(_emrAutoScalingRole)
	}
	if len(_emrAutoTerminationPolicy) > 0 {
		if err := assignInputField(input, "AutoTerminationPolicy", _emrAutoTerminationPolicy); err != nil {
			log.Errorf("invalid --auto-termination-policy: %s", err.Error())
			return
		}
	}
	if len(_emrBootstrapActions) > 0 {
		if err := assignInputField(input, "BootstrapActions", _emrBootstrapActions); err != nil {
			log.Errorf("invalid --bootstrap-actions: %s", err.Error())
			return
		}
	}
	if len(_emrConfigurations) > 0 {
		if err := assignInputField(input, "Configurations", _emrConfigurations); err != nil {
			log.Errorf("invalid --configurations: %s", err.Error())
			return
		}
	}
	if len(_emrCustomAmiId) > 0 {
		input.CustomAmiId = aws.String(_emrCustomAmiId)
	}
	if len(_emrEbsRootVolumeIops) > 0 {
		if err := assignInputField(input, "EbsRootVolumeIops", _emrEbsRootVolumeIops); err != nil {
			log.Errorf("invalid --ebs-root-volume-iops: %s", err.Error())
			return
		}
	}
	if len(_emrEbsRootVolumeSize) > 0 {
		if err := assignInputField(input, "EbsRootVolumeSize", _emrEbsRootVolumeSize); err != nil {
			log.Errorf("invalid --ebs-root-volume-size: %s", err.Error())
			return
		}
	}
	if len(_emrEbsRootVolumeThroughput) > 0 {
		if err := assignInputField(input, "EbsRootVolumeThroughput", _emrEbsRootVolumeThroughput); err != nil {
			log.Errorf("invalid --ebs-root-volume-throughput: %s", err.Error())
			return
		}
	}
	if len(_emrExtendedSupport) > 0 {
		if err := assignInputField(input, "ExtendedSupport", _emrExtendedSupport); err != nil {
			log.Errorf("invalid --extended-support: %s", err.Error())
			return
		}
	}
	if len(_emrJobFlowRole) > 0 {
		input.JobFlowRole = aws.String(_emrJobFlowRole)
	}
	if len(_emrKerberosAttributes) > 0 {
		if err := assignInputField(input, "KerberosAttributes", _emrKerberosAttributes); err != nil {
			log.Errorf("invalid --kerberos-attributes: %s", err.Error())
			return
		}
	}
	if len(_emrLogEncryptionKmsKeyId) > 0 {
		input.LogEncryptionKmsKeyId = aws.String(_emrLogEncryptionKmsKeyId)
	}
	if len(_emrLogUri) > 0 {
		input.LogUri = aws.String(_emrLogUri)
	}
	if len(_emrManagedScalingPolicy) > 0 {
		if err := assignInputField(input, "ManagedScalingPolicy", _emrManagedScalingPolicy); err != nil {
			log.Errorf("invalid --managed-scaling-policy: %s", err.Error())
			return
		}
	}
	if len(_emrMonitoringConfiguration) > 0 {
		if err := assignInputField(input, "MonitoringConfiguration", _emrMonitoringConfiguration); err != nil {
			log.Errorf("invalid --monitoring-configuration: %s", err.Error())
			return
		}
	}
	if len(_emrNewSupportedProducts) > 0 {
		if err := assignInputField(input, "NewSupportedProducts", _emrNewSupportedProducts); err != nil {
			log.Errorf("invalid --new-supported-products: %s", err.Error())
			return
		}
	}
	if len(_emrOSReleaseLabel) > 0 {
		input.OSReleaseLabel = aws.String(_emrOSReleaseLabel)
	}
	if len(_emrPlacementGroupConfigs) > 0 {
		if err := assignInputField(input, "PlacementGroupConfigs", _emrPlacementGroupConfigs); err != nil {
			log.Errorf("invalid --placement-group-configs: %s", err.Error())
			return
		}
	}
	if len(_emrReleaseLabel) > 0 {
		input.ReleaseLabel = aws.String(_emrReleaseLabel)
	}
	if len(_emrRepoUpgradeOnBoot) > 0 {
		if err := assignInputField(input, "RepoUpgradeOnBoot", _emrRepoUpgradeOnBoot); err != nil {
			log.Errorf("invalid --repo-upgrade-on-boot: %s", err.Error())
			return
		}
	}
	if len(_emrScaleDownBehavior) > 0 {
		if err := assignInputField(input, "ScaleDownBehavior", _emrScaleDownBehavior); err != nil {
			log.Errorf("invalid --scale-down-behavior: %s", err.Error())
			return
		}
	}
	if len(_emrSecurityConfiguration) > 0 {
		input.SecurityConfiguration = aws.String(_emrSecurityConfiguration)
	}
	if len(_emrServiceRole) > 0 {
		input.ServiceRole = aws.String(_emrServiceRole)
	}
	if len(_emrStepConcurrencyLevel) > 0 {
		if err := assignInputField(input, "StepConcurrencyLevel", _emrStepConcurrencyLevel); err != nil {
			log.Errorf("invalid --step-concurrency-level: %s", err.Error())
			return
		}
	}
	if len(_emrSteps) > 0 {
		if err := assignInputField(input, "Steps", _emrSteps); err != nil {
			log.Errorf("invalid --steps: %s", err.Error())
			return
		}
	}
	if len(_emrSupportedProducts) > 0 {
		input.SupportedProducts = append([]string(nil), _emrSupportedProducts...)
	}
	if len(_emrTags) > 0 {
		if err := assignInputField(input, "Tags", _emrTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_emrVisibleToAllUsers) > 0 {
		if err := assignInputField(input, "VisibleToAllUsers", _emrVisibleToAllUsers); err != nil {
			log.Errorf("invalid --visible-to-all-users: %s", err.Error())
			return
		}
	}

	if resp, err := client.RunJobFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// You can use the SetKeepJobFlowAliveWhenNoSteps to configure a cluster (job
// flow) to terminate after the step execution, i.e., all your steps are executed.
// If you want a transient cluster that shuts down after the last of the current
// executing steps are completed, you can configure SetKeepJobFlowAliveWhenNoSteps
// to false. If you want a long running cluster, configure
// SetKeepJobFlowAliveWhenNoSteps to true.
//
// For more information, see [Managing Cluster Termination] in the Amazon EMR Management Guide.
//
// [Managing Cluster Termination]: https://docs.aws.amazon.com/emr/latest/ManagementGuide/UsingEMR_TerminationProtection.html
func emr_SetKeepJobFlowAliveWhenNoSteps(cfg aws.Config, client *emr.Client) {
	input := &emr.SetKeepJobFlowAliveWhenNoStepsInput{
		// JobFlowIds: []string, // Required
		// KeepJobFlowAliveWhenNoSteps: *bool, // Required
	}

	if len(_emrJobFlowIds) > 0 {
		input.JobFlowIds = append([]string(nil), _emrJobFlowIds...)
	}
	if len(_emrKeepJobFlowAliveWhenNoSteps) > 0 {
		if err := assignInputField(input, "KeepJobFlowAliveWhenNoSteps", _emrKeepJobFlowAliveWhenNoSteps); err != nil {
			log.Errorf("invalid --keep-job-flow-alive-when-no-steps: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetKeepJobFlowAliveWhenNoSteps(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// SetTerminationProtection locks a cluster (job flow) so the Amazon EC2 instances
// in the cluster cannot be terminated by user intervention, an API call, or in the
// event of a job-flow error. The cluster still terminates upon successful
// completion of the job flow. Calling SetTerminationProtection on a cluster is
// similar to calling the Amazon EC2 DisableAPITermination API on all Amazon EC2
// instances in a cluster.
//
// SetTerminationProtection is used to prevent accidental termination of a cluster
// and to ensure that in the event of an error, the instances persist so that you
// can recover any data stored in their ephemeral instance storage.
//
// To terminate a cluster that has been locked by setting SetTerminationProtection
// to true , you must first unlock the job flow by a subsequent call to
// SetTerminationProtection in which you set the value to false .
//
// For more information, see [Managing Cluster Termination] in the Amazon EMR Management Guide.
//
// [Managing Cluster Termination]: https://docs.aws.amazon.com/emr/latest/ManagementGuide/UsingEMR_TerminationProtection.html
func emr_SetTerminationProtection(cfg aws.Config, client *emr.Client) {
	input := &emr.SetTerminationProtectionInput{
		// JobFlowIds: []string, // Required
		// TerminationProtected: *bool, // Required
	}

	if len(_emrJobFlowIds) > 0 {
		input.JobFlowIds = append([]string(nil), _emrJobFlowIds...)
	}
	if len(_emrTerminationProtected) > 0 {
		if err := assignInputField(input, "TerminationProtected", _emrTerminationProtected); err != nil {
			log.Errorf("invalid --termination-protected: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetTerminationProtection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specify whether to enable unhealthy node replacement, which lets Amazon EMR
// gracefully replace core nodes on a cluster if any nodes become unhealthy. For
// example, a node becomes unhealthy if disk usage is above 90%. If unhealthy node
// replacement is on and TerminationProtected are off, Amazon EMR immediately
// terminates the unhealthy core nodes. To use unhealthy node replacement and
// retain unhealthy core nodes, use to turn on termination protection. In such
// cases, Amazon EMR adds the unhealthy nodes to a denylist, reducing job
// interruptions and failures.
//
// If unhealthy node replacement is on, Amazon EMR notifies YARN and other
// applications on the cluster to stop scheduling tasks with these nodes, moves the
// data, and then terminates the nodes.
//
// For more information, see [graceful node replacement] in the Amazon EMR Management Guide.
//
// [graceful node replacement]: https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-plan-node-replacement.html
func emr_SetUnhealthyNodeReplacement(cfg aws.Config, client *emr.Client) {
	input := &emr.SetUnhealthyNodeReplacementInput{
		// JobFlowIds: []string, // Required
		// UnhealthyNodeReplacement: *bool, // Required
	}

	if len(_emrJobFlowIds) > 0 {
		input.JobFlowIds = append([]string(nil), _emrJobFlowIds...)
	}
	if len(_emrUnhealthyNodeReplacement) > 0 {
		if err := assignInputField(input, "UnhealthyNodeReplacement", _emrUnhealthyNodeReplacement); err != nil {
			log.Errorf("invalid --unhealthy-node-replacement: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetUnhealthyNodeReplacement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The SetVisibleToAllUsers parameter is no longer supported. Your cluster may be
// visible to all users in your account. To restrict cluster access using an IAM
// policy, see [Identity and Access Management for Amazon EMR].
//
// Sets the Cluster$VisibleToAllUsers value for an Amazon EMR cluster. When true , IAM principals in the
// Amazon Web Services account can perform Amazon EMR cluster actions that their
// IAM policies allow. When false , only the IAM principal that created the cluster
// and the Amazon Web Services account root user can perform Amazon EMR actions on
// the cluster, regardless of IAM permissions policies attached to other IAM
// principals.
//
// This action works on running clusters. When you create a cluster, use the RunJobFlowInput$VisibleToAllUsers
// parameter.
//
// For more information, see [Understanding the Amazon EMR Cluster VisibleToAllUsers Setting] in the Amazon EMR Management Guide.
//
// [Understanding the Amazon EMR Cluster VisibleToAllUsers Setting]: https://docs.aws.amazon.com/emr/latest/ManagementGuide/security_IAM_emr-with-IAM.html#security_set_visible_to_all_users
// [Identity and Access Management for Amazon EMR]: https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-plan-access-IAM.html
func emr_SetVisibleToAllUsers(cfg aws.Config, client *emr.Client) {
	input := &emr.SetVisibleToAllUsersInput{
		// JobFlowIds: []string, // Required
		// VisibleToAllUsers: *bool, // Required
	}

	if len(_emrJobFlowIds) > 0 {
		input.JobFlowIds = append([]string(nil), _emrJobFlowIds...)
	}
	if len(_emrVisibleToAllUsers) > 0 {
		if err := assignInputField(input, "VisibleToAllUsers", _emrVisibleToAllUsers); err != nil {
			log.Errorf("invalid --visible-to-all-users: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetVisibleToAllUsers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a notebook execution.
func emr_StartNotebookExecution(cfg aws.Config, client *emr.Client) {
	input := &emr.StartNotebookExecutionInput{
		// ExecutionEngine: *types.ExecutionEngineConfig, // Required
		// ServiceRole: *string, // Required
	}

	if len(_emrExecutionEngine) > 0 {
		if err := assignInputField(input, "ExecutionEngine", _emrExecutionEngine); err != nil {
			log.Errorf("invalid --execution-engine: %s", err.Error())
			return
		}
	}
	if len(_emrServiceRole) > 0 {
		input.ServiceRole = aws.String(_emrServiceRole)
	}
	if len(_emrEditorId) > 0 {
		input.EditorId = aws.String(_emrEditorId)
	}
	if len(_emrEnvironmentVariables) > 0 {
		if err := assignInputField(input, "EnvironmentVariables", _emrEnvironmentVariables); err != nil {
			log.Errorf("invalid --environment-variables: %s", err.Error())
			return
		}
	}
	if len(_emrNotebookExecutionName) > 0 {
		input.NotebookExecutionName = aws.String(_emrNotebookExecutionName)
	}
	if len(_emrNotebookInstanceSecurityGroupId) > 0 {
		input.NotebookInstanceSecurityGroupId = aws.String(_emrNotebookInstanceSecurityGroupId)
	}
	if len(_emrNotebookParams) > 0 {
		input.NotebookParams = aws.String(_emrNotebookParams)
	}
	if len(_emrNotebookS3Location) > 0 {
		if err := assignInputField(input, "NotebookS3Location", _emrNotebookS3Location); err != nil {
			log.Errorf("invalid --notebook-s3-location: %s", err.Error())
			return
		}
	}
	if len(_emrOutputNotebookFormat) > 0 {
		if err := assignInputField(input, "OutputNotebookFormat", _emrOutputNotebookFormat); err != nil {
			log.Errorf("invalid --output-notebook-format: %s", err.Error())
			return
		}
	}
	if len(_emrOutputNotebookS3Location) > 0 {
		if err := assignInputField(input, "OutputNotebookS3Location", _emrOutputNotebookS3Location); err != nil {
			log.Errorf("invalid --output-notebook-s3-location: %s", err.Error())
			return
		}
	}
	if len(_emrRelativePath) > 0 {
		input.RelativePath = aws.String(_emrRelativePath)
	}
	if len(_emrTags) > 0 {
		if err := assignInputField(input, "Tags", _emrTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartNotebookExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a notebook execution.
func emr_StopNotebookExecution(cfg aws.Config, client *emr.Client) {
	input := &emr.StopNotebookExecutionInput{
		// NotebookExecutionId: *string, // Required
	}

	if len(_emrNotebookExecutionId) > 0 {
		input.NotebookExecutionId = aws.String(_emrNotebookExecutionId)
	}

	if resp, err := client.StopNotebookExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// TerminateJobFlows shuts a list of clusters (job flows) down. When a job flow is
// shut down, any step not yet completed is canceled and the Amazon EC2 instances
// on which the cluster is running are stopped. Any log files not already saved are
// uploaded to Amazon S3 if a LogUri was specified when the cluster was created.
//
// The maximum number of clusters allowed is 10. The call to TerminateJobFlows is
// asynchronous. Depending on the configuration of the cluster, it may take up to
// 1-5 minutes for the cluster to completely terminate and release allocated
// resources, such as Amazon EC2 instances.
func emr_TerminateJobFlows(cfg aws.Config, client *emr.Client) {
	input := &emr.TerminateJobFlowsInput{
		// JobFlowIds: []string, // Required
	}

	if len(_emrJobFlowIds) > 0 {
		input.JobFlowIds = append([]string(nil), _emrJobFlowIds...)
	}

	if resp, err := client.TerminateJobFlows(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Amazon EMR Studio configuration, including attributes such as name,
// description, and subnets.
func emr_UpdateStudio(cfg aws.Config, client *emr.Client) {
	input := &emr.UpdateStudioInput{
		// StudioId: *string, // Required
	}

	if len(_emrStudioId) > 0 {
		input.StudioId = aws.String(_emrStudioId)
	}
	if len(_emrDefaultS3Location) > 0 {
		input.DefaultS3Location = aws.String(_emrDefaultS3Location)
	}
	if len(_emrDescription) > 0 {
		input.Description = aws.String(_emrDescription)
	}
	if len(_emrEncryptionKeyArn) > 0 {
		input.EncryptionKeyArn = aws.String(_emrEncryptionKeyArn)
	}
	if len(_emrName) > 0 {
		input.Name = aws.String(_emrName)
	}
	if len(_emrSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _emrSubnetIds...)
	}

	if resp, err := client.UpdateStudio(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the session policy attached to the user or group for the specified
// Amazon EMR Studio.
func emr_UpdateStudioSessionMapping(cfg aws.Config, client *emr.Client) {
	input := &emr.UpdateStudioSessionMappingInput{
		// IdentityType: types.IdentityType, // Required
		// SessionPolicyArn: *string, // Required
		// StudioId: *string, // Required
	}

	if len(_emrIdentityType) > 0 {
		if err := assignInputField(input, "IdentityType", _emrIdentityType); err != nil {
			log.Errorf("invalid --identity-type: %s", err.Error())
			return
		}
	}
	if len(_emrSessionPolicyArn) > 0 {
		input.SessionPolicyArn = aws.String(_emrSessionPolicyArn)
	}
	if len(_emrStudioId) > 0 {
		input.StudioId = aws.String(_emrStudioId)
	}
	if len(_emrIdentityId) > 0 {
		input.IdentityId = aws.String(_emrIdentityId)
	}
	if len(_emrIdentityName) > 0 {
		input.IdentityName = aws.String(_emrIdentityName)
	}

	if resp, err := client.UpdateStudioSessionMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_emrCmd)
	_emrCmd.Flags().SortFlags = false

	_emrCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_emrCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_emrCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_emrCmd.Flags().StringVarP(&_emrAdditionalInfo, "additional-info", "", "", "Additional Info")
	_emrCmd.Flags().StringVarP(&_emrAmiVersion, "ami-version", "", "", "AMI Version")
	_emrCmd.Flags().StringVarP(&_emrApplicationId, "application-id", "", "", "Application ID")
	_emrCmd.Flags().StringVarP(&_emrApplications, "applications", "", "", "Applications")
	_emrCmd.Flags().StringVarP(&_emrAuthMode, "auth-mode", "", "", "Auth Mode")
	_emrCmd.Flags().StringVarP(&_emrAuthProxyCall, "auth-proxy-call", "", "", "Auth Proxy Call")
	_emrCmd.Flags().StringVarP(&_emrAutoScalingPolicy, "auto-scaling-policy", "", "", "Auto Scaling Policy")
	_emrCmd.Flags().StringVarP(&_emrAutoScalingRole, "auto-scaling-role", "", "", "Auto Scaling Role")
	_emrCmd.Flags().StringVarP(&_emrAutoTerminationPolicy, "auto-termination-policy", "", "", "Auto Termination Policy")
	_emrCmd.Flags().StringVarP(&_emrBlockPublicAccessConfiguration, "block-public-access-configuration", "", "", "Block Public Access Configuration")
	_emrCmd.Flags().StringVarP(&_emrBootstrapActions, "bootstrap-actions", "", "", "Bootstrap Actions")
	_emrCmd.Flags().StringVarP(&_emrClusterId, "cluster-id", "", "", "Cluster ID")
	_emrCmd.Flags().StringVarP(&_emrClusterStates, "cluster-states", "", "", "Cluster States")
	_emrCmd.Flags().StringVarP(&_emrConfigurations, "configurations", "", "", "Configurations")
	_emrCmd.Flags().StringVarP(&_emrCreatedAfter, "created-after", "", "", "Created After")
	_emrCmd.Flags().StringVarP(&_emrCreatedBefore, "created-before", "", "", "Created Before")
	_emrCmd.Flags().StringVarP(&_emrCustomAmiId, "custom-ami-id", "", "", "Custom AMI ID")
	_emrCmd.Flags().StringVarP(&_emrDefaultS3Location, "default-s3-location", "", "", "Default S3 Location")
	_emrCmd.Flags().StringVarP(&_emrDescription, "description", "", "", "Description")
	_emrCmd.Flags().StringVarP(&_emrDryRun, "dry-run", "", "", "Dry Run")
	_emrCmd.Flags().StringVarP(&_emrEbsRootVolumeIops, "ebs-root-volume-iops", "", "", "Ebs Root Volume IOPS")
	_emrCmd.Flags().StringVarP(&_emrEbsRootVolumeSize, "ebs-root-volume-size", "", "", "Ebs Root Volume Size")
	_emrCmd.Flags().StringVarP(&_emrEbsRootVolumeThroughput, "ebs-root-volume-throughput", "", "", "Ebs Root Volume Throughput")
	_emrCmd.Flags().StringVarP(&_emrEditorId, "editor-id", "", "", "Editor ID")
	_emrCmd.Flags().StringVarP(&_emrEMRContainersConfig, "emr-containers-config", "", "", "Emr Containers Config")
	_emrCmd.Flags().StringVarP(&_emrEncryptionKeyArn, "encryption-key-arn", "", "", "Encryption Key ARN")
	_emrCmd.Flags().StringVarP(&_emrEngineSecurityGroupId, "engine-security-group-id", "", "", "Engine Security Group ID")
	_emrCmd.Flags().StringVarP(&_emrEnvironmentVariables, "environment-variables", "", "", "Environment Variables")
	_emrCmd.Flags().StringVarP(&_emrExecutionEngine, "execution-engine", "", "", "Execution Engine")
	_emrCmd.Flags().StringVarP(&_emrExecutionEngineId, "execution-engine-id", "", "", "Execution Engine ID")
	_emrCmd.Flags().StringVarP(&_emrExecutionRoleArn, "execution-role-arn", "", "", "Execution Role ARN")
	_emrCmd.Flags().StringVarP(&_emrExtendedSupport, "extended-support", "", "", "Extended Support")
	_emrCmd.Flags().StringVarP(&_emrFilters, "filters", "", "", "Filters")
	_emrCmd.Flags().StringVarP(&_emrFrom, "from", "", "", "From")
	_emrCmd.Flags().StringVarP(&_emrIdcInstanceArn, "idc-instance-arn", "", "", "Idc Instance ARN")
	_emrCmd.Flags().StringVarP(&_emrIdcUserAssignment, "idc-user-assignment", "", "", "Idc User Assignment")
	_emrCmd.Flags().StringVarP(&_emrIdentityId, "identity-id", "", "", "Identity ID")
	_emrCmd.Flags().StringVarP(&_emrIdentityName, "identity-name", "", "", "Identity Name")
	_emrCmd.Flags().StringVarP(&_emrIdentityType, "identity-type", "", "", "Identity Type")
	_emrCmd.Flags().StringVarP(&_emrIdpAuthUrl, "idp-auth-url", "", "", "Idp Auth URL")
	_emrCmd.Flags().StringVarP(&_emrIdpRelayStateParameterName, "idp-relay-state-parameter-name", "", "", "Idp Relay State Parameter Name")
	_emrCmd.Flags().StringVarP(&_emrInstanceFleet, "instance-fleet", "", "", "Instance Fleet")
	_emrCmd.Flags().StringVarP(&_emrInstanceFleetId, "instance-fleet-id", "", "", "Instance Fleet ID")
	_emrCmd.Flags().StringVarP(&_emrInstanceFleetType, "instance-fleet-type", "", "", "Instance Fleet Type")
	_emrCmd.Flags().StringVarP(&_emrInstanceGroupId, "instance-group-id", "", "", "Instance Group ID")
	_emrCmd.Flags().StringVarP(&_emrInstanceGroupTypes, "instance-group-types", "", "", "Instance Group Types")
	_emrCmd.Flags().StringVarP(&_emrInstanceGroups, "instance-groups", "", "", "Instance Groups")
	_emrCmd.Flags().StringVarP(&_emrInstanceStates, "instance-states", "", "", "Instance States")
	_emrCmd.Flags().StringVarP(&_emrInstances, "instances", "", "", "Instances")
	_emrCmd.Flags().StringVarP(&_emrJobFlowId, "job-flow-id", "", "", "Job Flow ID")
	_emrCmd.Flags().StringSliceVarP(&_emrJobFlowIds, "job-flow-ids", "", nil, "Job Flow Ids")
	_emrCmd.Flags().StringVarP(&_emrJobFlowRole, "job-flow-role", "", "", "Job Flow Role")
	_emrCmd.Flags().StringVarP(&_emrJobFlowStates, "job-flow-states", "", "", "Job Flow States")
	_emrCmd.Flags().StringVarP(&_emrKeepJobFlowAliveWhenNoSteps, "keep-job-flow-alive-when-no-steps", "", "", "Keep Job Flow Alive When No Steps")
	_emrCmd.Flags().StringVarP(&_emrKerberosAttributes, "kerberos-attributes", "", "", "Kerberos Attributes")
	_emrCmd.Flags().StringVarP(&_emrLogEncryptionKmsKeyId, "log-encryption-kms-key-id", "", "", "Log Encryption KMS Key ID")
	_emrCmd.Flags().StringVarP(&_emrLogUri, "log-uri", "", "", "Log URI")
	_emrCmd.Flags().StringVarP(&_emrManagedScalingPolicy, "managed-scaling-policy", "", "", "Managed Scaling Policy")
	_emrCmd.Flags().StringVarP(&_emrMarker, "marker", "", "", "Marker")
	_emrCmd.Flags().StringVarP(&_emrMaxResults, "max-results", "", "", "Max Results")
	_emrCmd.Flags().StringVarP(&_emrMonitoringConfiguration, "monitoring-configuration", "", "", "Monitoring Configuration")
	_emrCmd.Flags().StringVarP(&_emrName, "name", "", "", "Name")
	_emrCmd.Flags().StringVarP(&_emrNewSupportedProducts, "new-supported-products", "", "", "New Supported Products")
	_emrCmd.Flags().StringVarP(&_emrNextToken, "next-token", "", "", "Next Token")
	_emrCmd.Flags().StringVarP(&_emrNotebookExecutionId, "notebook-execution-id", "", "", "Notebook Execution ID")
	_emrCmd.Flags().StringVarP(&_emrNotebookExecutionName, "notebook-execution-name", "", "", "Notebook Execution Name")
	_emrCmd.Flags().StringVarP(&_emrNotebookInstanceSecurityGroupId, "notebook-instance-security-group-id", "", "", "Notebook Instance Security Group ID")
	_emrCmd.Flags().StringVarP(&_emrNotebookParams, "notebook-params", "", "", "Notebook Params")
	_emrCmd.Flags().StringVarP(&_emrNotebookS3Location, "notebook-s3-location", "", "", "Notebook S3 Location")
	_emrCmd.Flags().StringVarP(&_emrOnClusterAppUIType, "on-cluster-app-ui-type", "", "", "On Cluster App Ui Type")
	_emrCmd.Flags().StringVarP(&_emrOSReleaseLabel, "os-release-label", "", "", "OS Release Label")
	_emrCmd.Flags().StringVarP(&_emrOutputNotebookFormat, "output-notebook-format", "", "", "Output Notebook Format")
	_emrCmd.Flags().StringVarP(&_emrOutputNotebookS3Location, "output-notebook-s3-location", "", "", "Output Notebook S3 Location")
	_emrCmd.Flags().StringVarP(&_emrPersistentAppUIType, "persistent-app-ui-type", "", "", "Persistent App Ui Type")
	_emrCmd.Flags().StringVarP(&_emrPersistentAppUIId, "persistent-app-uiid", "", "", "Persistent App Uiid")
	_emrCmd.Flags().StringVarP(&_emrPlacementGroupConfigs, "placement-group-configs", "", "", "Placement Group Configs")
	_emrCmd.Flags().StringVarP(&_emrProfilerType, "profiler-type", "", "", "Profiler Type")
	_emrCmd.Flags().StringVarP(&_emrRelativePath, "relative-path", "", "", "Relative Path")
	_emrCmd.Flags().StringVarP(&_emrReleaseLabel, "release-label", "", "", "Release Label")
	_emrCmd.Flags().StringVarP(&_emrRepoUpgradeOnBoot, "repo-upgrade-on-boot", "", "", "Repo Upgrade On Boot")
	_emrCmd.Flags().StringVarP(&_emrResourceId, "resource-id", "", "", "Resource ID")
	_emrCmd.Flags().StringVarP(&_emrScaleDownBehavior, "scale-down-behavior", "", "", "Scale Down Behavior")
	_emrCmd.Flags().StringVarP(&_emrSecurityConfiguration, "security-configuration", "", "", "Security Configuration")
	_emrCmd.Flags().StringVarP(&_emrServiceRole, "service-role", "", "", "Service Role")
	_emrCmd.Flags().StringVarP(&_emrSessionPolicyArn, "session-policy-arn", "", "", "Session Policy ARN")
	_emrCmd.Flags().StringVarP(&_emrStatus, "status", "", "", "Status")
	_emrCmd.Flags().StringVarP(&_emrStepCancellationOption, "step-cancellation-option", "", "", "Step Cancellation Option")
	_emrCmd.Flags().StringVarP(&_emrStepConcurrencyLevel, "step-concurrency-level", "", "", "Step Concurrency Level")
	_emrCmd.Flags().StringVarP(&_emrStepId, "step-id", "", "", "Step ID")
	_emrCmd.Flags().StringSliceVarP(&_emrStepIds, "step-ids", "", nil, "Step Ids")
	_emrCmd.Flags().StringVarP(&_emrStepStates, "step-states", "", "", "Step States")
	_emrCmd.Flags().StringVarP(&_emrSteps, "steps", "", "", "Steps")
	_emrCmd.Flags().StringVarP(&_emrStudioId, "studio-id", "", "", "Studio ID")
	_emrCmd.Flags().StringSliceVarP(&_emrSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_emrCmd.Flags().StringSliceVarP(&_emrSupportedProducts, "supported-products", "", nil, "Supported Products")
	_emrCmd.Flags().StringSliceVarP(&_emrTagKeys, "tag-keys", "", nil, "Tag Keys")
	_emrCmd.Flags().StringVarP(&_emrTags, "tags", "", "", "Tags")
	_emrCmd.Flags().StringVarP(&_emrTargetResourceArn, "target-resource-arn", "", "", "Target Resource ARN")
	_emrCmd.Flags().StringVarP(&_emrTerminationProtected, "termination-protected", "", "", "Termination Protected")
	_emrCmd.Flags().StringVarP(&_emrTo, "to", "", "", "To")
	_emrCmd.Flags().StringVarP(&_emrTrustedIdentityPropagationEnabled, "trusted-identity-propagation-enabled", "", "", "Trusted Identity Propagation Enabled")
	_emrCmd.Flags().StringVarP(&_emrUnhealthyNodeReplacement, "unhealthy-node-replacement", "", "", "Unhealthy Node Replacement")
	_emrCmd.Flags().StringVarP(&_emrUserRole, "user-role", "", "", "User Role")
	_emrCmd.Flags().StringVarP(&_emrVisibleToAllUsers, "visible-to-all-users", "", "", "Visible To All Users")
	_emrCmd.Flags().StringVarP(&_emrVpcId, "vpc-id", "", "", "VPC ID")
	_emrCmd.Flags().StringVarP(&_emrWorkspaceSecurityGroupId, "workspace-security-group-id", "", "", "Workspace Security Group ID")
	_emrCmd.Flags().StringVarP(&_emrXReferer, "xreferer", "", "", "Xreferer")

	_emrCmd.Flags().BoolVarP(&_emrAddInstanceFleet, "add-instance-fleet", "", false, "Add Instance Fleet")
	_emrCmd.Flags().BoolVarP(&_emrAddInstanceGroups, "add-instance-groups", "", false, "Add Instance Groups")
	_emrCmd.Flags().BoolVarP(&_emrAddJobFlowSteps, "add-job-flow-steps", "", false, "Add Job Flow Steps")
	_emrCmd.Flags().BoolVarP(&_emrAddTags, "add-tags", "", false, "Add Tags")
	_emrCmd.Flags().BoolVarP(&_emrCancelSteps, "cancel-steps", "", false, "Cancel Steps")
	_emrCmd.Flags().BoolVarP(&_emrCreatePersistentAppUI, "create-persistent-app-ui", "", false, "Create Persistent App Ui")
	_emrCmd.Flags().BoolVarP(&_emrCreateSecurityConfiguration, "create-security-configuration", "", false, "Create Security Configuration")
	_emrCmd.Flags().BoolVarP(&_emrCreateStudio, "create-studio", "", false, "Create Studio")
	_emrCmd.Flags().BoolVarP(&_emrCreateStudioSessionMapping, "create-studio-session-mapping", "", false, "Create Studio Session Mapping")
	_emrCmd.Flags().BoolVarP(&_emrDeleteSecurityConfiguration, "delete-security-configuration", "", false, "Delete Security Configuration")
	_emrCmd.Flags().BoolVarP(&_emrDeleteStudio, "delete-studio", "", false, "Delete Studio")
	_emrCmd.Flags().BoolVarP(&_emrDeleteStudioSessionMapping, "delete-studio-session-mapping", "", false, "Delete Studio Session Mapping")
	_emrCmd.Flags().BoolVarP(&_emrDescribeCluster, "describe-cluster", "", false, "Describe Cluster")
	_emrCmd.Flags().BoolVarP(&_emrDescribeJobFlows, "describe-job-flows", "", false, "Describe Job Flows")
	_emrCmd.Flags().BoolVarP(&_emrDescribeNotebookExecution, "describe-notebook-execution", "", false, "Describe Notebook Execution")
	_emrCmd.Flags().BoolVarP(&_emrDescribePersistentAppUI, "describe-persistent-app-ui", "", false, "Describe Persistent App Ui")
	_emrCmd.Flags().BoolVarP(&_emrDescribeReleaseLabel, "describe-release-label", "", false, "Describe Release Label")
	_emrCmd.Flags().BoolVarP(&_emrDescribeSecurityConfiguration, "describe-security-configuration", "", false, "Describe Security Configuration")
	_emrCmd.Flags().BoolVarP(&_emrDescribeStep, "describe-step", "", false, "Describe Step")
	_emrCmd.Flags().BoolVarP(&_emrDescribeStudio, "describe-studio", "", false, "Describe Studio")
	_emrCmd.Flags().BoolVarP(&_emrGetAutoTerminationPolicy, "get-auto-termination-policy", "", false, "Get Auto Termination Policy")
	_emrCmd.Flags().BoolVarP(&_emrGetBlockPublicAccessConfiguration, "get-block-public-access-configuration", "", false, "Get Block Public Access Configuration")
	_emrCmd.Flags().BoolVarP(&_emrGetClusterSessionCredentials, "get-cluster-session-credentials", "", false, "Get Cluster Session Credentials")
	_emrCmd.Flags().BoolVarP(&_emrGetManagedScalingPolicy, "get-managed-scaling-policy", "", false, "Get Managed Scaling Policy")
	_emrCmd.Flags().BoolVarP(&_emrGetOnClusterAppUIPresignedURL, "get-on-cluster-app-ui-presigned-url", "", false, "Get On Cluster App Ui Presigned URL")
	_emrCmd.Flags().BoolVarP(&_emrGetPersistentAppUIPresignedURL, "get-persistent-app-ui-presigned-url", "", false, "Get Persistent App Ui Presigned URL")
	_emrCmd.Flags().BoolVarP(&_emrGetStudioSessionMapping, "get-studio-session-mapping", "", false, "Get Studio Session Mapping")
	_emrCmd.Flags().BoolVarP(&_emrListBootstrapActions, "list-bootstrap-actions", "", false, "List Bootstrap Actions")
	_emrCmd.Flags().BoolVarP(&_emrListClusters, "list-clusters", "", false, "List Clusters")
	_emrCmd.Flags().BoolVarP(&_emrListInstanceFleets, "list-instance-fleets", "", false, "List Instance Fleets")
	_emrCmd.Flags().BoolVarP(&_emrListInstanceGroups, "list-instance-groups", "", false, "List Instance Groups")
	_emrCmd.Flags().BoolVarP(&_emrListInstances, "list-instances", "", false, "List Instances")
	_emrCmd.Flags().BoolVarP(&_emrListNotebookExecutions, "list-notebook-executions", "", false, "List Notebook Executions")
	_emrCmd.Flags().BoolVarP(&_emrListReleaseLabels, "list-release-labels", "", false, "List Release Labels")
	_emrCmd.Flags().BoolVarP(&_emrListSecurityConfigurations, "list-security-configurations", "", false, "List Security Configurations")
	_emrCmd.Flags().BoolVarP(&_emrListSteps, "list-steps", "", false, "List Steps")
	_emrCmd.Flags().BoolVarP(&_emrListStudioSessionMappings, "list-studio-session-mappings", "", false, "List Studio Session Mappings")
	_emrCmd.Flags().BoolVarP(&_emrListStudios, "list-studios", "", false, "List Studios")
	_emrCmd.Flags().BoolVarP(&_emrListSupportedInstanceTypes, "list-supported-instance-types", "", false, "List Supported Instance Types")
	_emrCmd.Flags().BoolVarP(&_emrModifyCluster, "modify-cluster", "", false, "Modify Cluster")
	_emrCmd.Flags().BoolVarP(&_emrModifyInstanceFleet, "modify-instance-fleet", "", false, "Modify Instance Fleet")
	_emrCmd.Flags().BoolVarP(&_emrModifyInstanceGroups, "modify-instance-groups", "", false, "Modify Instance Groups")
	_emrCmd.Flags().BoolVarP(&_emrPutAutoScalingPolicy, "put-auto-scaling-policy", "", false, "Put Auto Scaling Policy")
	_emrCmd.Flags().BoolVarP(&_emrPutAutoTerminationPolicy, "put-auto-termination-policy", "", false, "Put Auto Termination Policy")
	_emrCmd.Flags().BoolVarP(&_emrPutBlockPublicAccessConfiguration, "put-block-public-access-configuration", "", false, "Put Block Public Access Configuration")
	_emrCmd.Flags().BoolVarP(&_emrPutManagedScalingPolicy, "put-managed-scaling-policy", "", false, "Put Managed Scaling Policy")
	_emrCmd.Flags().BoolVarP(&_emrRemoveAutoScalingPolicy, "remove-auto-scaling-policy", "", false, "Remove Auto Scaling Policy")
	_emrCmd.Flags().BoolVarP(&_emrRemoveAutoTerminationPolicy, "remove-auto-termination-policy", "", false, "Remove Auto Termination Policy")
	_emrCmd.Flags().BoolVarP(&_emrRemoveManagedScalingPolicy, "remove-managed-scaling-policy", "", false, "Remove Managed Scaling Policy")
	_emrCmd.Flags().BoolVarP(&_emrRemoveTags, "remove-tags", "", false, "Remove Tags")
	_emrCmd.Flags().BoolVarP(&_emrRunJobFlow, "run-job-flow", "", false, "Run Job Flow")
	_emrCmd.Flags().BoolVarP(&_emrSetKeepJobFlowAliveWhenNoSteps, "set-keep-job-flow-alive-when-no-steps", "", false, "Set Keep Job Flow Alive When No Steps")
	_emrCmd.Flags().BoolVarP(&_emrSetTerminationProtection, "set-termination-protection", "", false, "Set Termination Protection")
	_emrCmd.Flags().BoolVarP(&_emrSetUnhealthyNodeReplacement, "set-unhealthy-node-replacement", "", false, "Set Unhealthy Node Replacement")
	_emrCmd.Flags().BoolVarP(&_emrSetVisibleToAllUsers, "set-visible-to-all-users", "", false, "Set Visible To All Users")
	_emrCmd.Flags().BoolVarP(&_emrStartNotebookExecution, "start-notebook-execution", "", false, "Start Notebook Execution")
	_emrCmd.Flags().BoolVarP(&_emrStopNotebookExecution, "stop-notebook-execution", "", false, "Stop Notebook Execution")
	_emrCmd.Flags().BoolVarP(&_emrTerminateJobFlows, "terminate-job-flows", "", false, "Terminate Job Flows")
	_emrCmd.Flags().BoolVarP(&_emrUpdateStudio, "update-studio", "", false, "Update Studio")
	_emrCmd.Flags().BoolVarP(&_emrUpdateStudioSessionMapping, "update-studio-session-mapping", "", false, "Update Studio Session Mapping")

}
