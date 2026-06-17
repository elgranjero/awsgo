package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workspacesinstances"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// workspacesinstancesCmd represents the workspacesinstances command
var _workspacesinstancesCmd = &cobra.Command{
	Use:   "workspacesinstances",
	Short: "AWS workspacesinstances CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := workspacesinstances.NewFromConfig(cfg)
		if _workspacesinstancesAssociateVolume {
			workspacesinstances_AssociateVolume(cfg, client)
			return
		}
		if _workspacesinstancesCreateVolume {
			workspacesinstances_CreateVolume(cfg, client)
			return
		}
		if _workspacesinstancesCreateWorkspaceInstance {
			workspacesinstances_CreateWorkspaceInstance(cfg, client)
			return
		}
		if _workspacesinstancesDeleteVolume {
			workspacesinstances_DeleteVolume(cfg, client)
			return
		}
		if _workspacesinstancesDeleteWorkspaceInstance {
			workspacesinstances_DeleteWorkspaceInstance(cfg, client)
			return
		}
		if _workspacesinstancesDisassociateVolume {
			workspacesinstances_DisassociateVolume(cfg, client)
			return
		}
		if _workspacesinstancesGetWorkspaceInstance {
			workspacesinstances_GetWorkspaceInstance(cfg, client)
			return
		}
		if _workspacesinstancesListInstanceTypes {
			workspacesinstances_ListInstanceTypes(cfg, client)
			return
		}
		if _workspacesinstancesListRegions {
			workspacesinstances_ListRegions(cfg, client)
			return
		}
		if _workspacesinstancesListTagsForResource {
			workspacesinstances_ListTagsForResource(cfg, client)
			return
		}
		if _workspacesinstancesListWorkspaceInstances {
			workspacesinstances_ListWorkspaceInstances(cfg, client)
			return
		}
		if _workspacesinstancesTagResource {
			workspacesinstances_TagResource(cfg, client)
			return
		}
		if _workspacesinstancesUntagResource {
			workspacesinstances_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_workspacesinstancesAssociateVolume         bool
	_workspacesinstancesCreateVolume            bool
	_workspacesinstancesCreateWorkspaceInstance bool
	_workspacesinstancesDeleteVolume            bool
	_workspacesinstancesDeleteWorkspaceInstance bool
	_workspacesinstancesDisassociateVolume      bool
	_workspacesinstancesGetWorkspaceInstance    bool
	_workspacesinstancesListInstanceTypes       bool
	_workspacesinstancesListRegions             bool
	_workspacesinstancesListTagsForResource     bool
	_workspacesinstancesListWorkspaceInstances  bool
	_workspacesinstancesTagResource             bool
	_workspacesinstancesUntagResource           bool

	_workspacesinstancesAvailabilityZone            string
	_workspacesinstancesBillingConfiguration        string
	_workspacesinstancesClientToken                 string
	_workspacesinstancesDevice                      string
	_workspacesinstancesDisassociateMode            string
	_workspacesinstancesEncrypted                   string
	_workspacesinstancesInstanceConfigurationFilter string
	_workspacesinstancesIops                        string
	_workspacesinstancesKmsKeyId                    string
	_workspacesinstancesManagedInstance             string
	_workspacesinstancesMaxResults                  string
	_workspacesinstancesNextToken                   string
	_workspacesinstancesProvisionStates             string
	_workspacesinstancesSizeInGB                    string
	_workspacesinstancesSnapshotId                  string
	_workspacesinstancesTagKeys                     []string
	_workspacesinstancesTagSpecifications           string
	_workspacesinstancesTags                        string
	_workspacesinstancesThroughput                  string
	_workspacesinstancesVolumeId                    string
	_workspacesinstancesVolumeType                  string
	_workspacesinstancesWorkspaceInstanceId         string
)

// Attaches a volume to a WorkSpace Instance.
func workspacesinstances_AssociateVolume(cfg aws.Config, client *workspacesinstances.Client) {
	input := &workspacesinstances.AssociateVolumeInput{
		// Device: *string, // Required
		// VolumeId: *string, // Required
		// WorkspaceInstanceId: *string, // Required
	}

	if len(_workspacesinstancesDevice) > 0 {
		input.Device = aws.String(_workspacesinstancesDevice)
	}
	if len(_workspacesinstancesVolumeId) > 0 {
		input.VolumeId = aws.String(_workspacesinstancesVolumeId)
	}
	if len(_workspacesinstancesWorkspaceInstanceId) > 0 {
		input.WorkspaceInstanceId = aws.String(_workspacesinstancesWorkspaceInstanceId)
	}

	if resp, err := client.AssociateVolume(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new volume for WorkSpace Instances.
func workspacesinstances_CreateVolume(cfg aws.Config, client *workspacesinstances.Client) {
	input := &workspacesinstances.CreateVolumeInput{
		// AvailabilityZone: *string, // Required
	}

	if len(_workspacesinstancesAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_workspacesinstancesAvailabilityZone)
	}
	if len(_workspacesinstancesClientToken) > 0 {
		input.ClientToken = aws.String(_workspacesinstancesClientToken)
	}
	if len(_workspacesinstancesEncrypted) > 0 {
		if err := assignInputField(input, "Encrypted", _workspacesinstancesEncrypted); err != nil {
			log.Errorf("invalid --encrypted: %s", err.Error())
			return
		}
	}
	if len(_workspacesinstancesIops) > 0 {
		if err := assignInputField(input, "Iops", _workspacesinstancesIops); err != nil {
			log.Errorf("invalid --iops: %s", err.Error())
			return
		}
	}
	if len(_workspacesinstancesKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_workspacesinstancesKmsKeyId)
	}
	if len(_workspacesinstancesSizeInGB) > 0 {
		if err := assignInputField(input, "SizeInGB", _workspacesinstancesSizeInGB); err != nil {
			log.Errorf("invalid --size-in-gb: %s", err.Error())
			return
		}
	}
	if len(_workspacesinstancesSnapshotId) > 0 {
		input.SnapshotId = aws.String(_workspacesinstancesSnapshotId)
	}
	if len(_workspacesinstancesTagSpecifications) > 0 {
		if err := assignInputField(input, "TagSpecifications", _workspacesinstancesTagSpecifications); err != nil {
			log.Errorf("invalid --tag-specifications: %s", err.Error())
			return
		}
	}
	if len(_workspacesinstancesThroughput) > 0 {
		if err := assignInputField(input, "Throughput", _workspacesinstancesThroughput); err != nil {
			log.Errorf("invalid --throughput: %s", err.Error())
			return
		}
	}
	if len(_workspacesinstancesVolumeType) > 0 {
		if err := assignInputField(input, "VolumeType", _workspacesinstancesVolumeType); err != nil {
			log.Errorf("invalid --volume-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVolume(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Launches a new WorkSpace Instance with specified configuration parameters,
// enabling programmatic workspace deployment.
func workspacesinstances_CreateWorkspaceInstance(cfg aws.Config, client *workspacesinstances.Client) {
	input := &workspacesinstances.CreateWorkspaceInstanceInput{
		// ManagedInstance: *types.ManagedInstanceRequest, // Required
	}

	if len(_workspacesinstancesManagedInstance) > 0 {
		if err := assignInputField(input, "ManagedInstance", _workspacesinstancesManagedInstance); err != nil {
			log.Errorf("invalid --managed-instance: %s", err.Error())
			return
		}
	}
	if len(_workspacesinstancesBillingConfiguration) > 0 {
		if err := assignInputField(input, "BillingConfiguration", _workspacesinstancesBillingConfiguration); err != nil {
			log.Errorf("invalid --billing-configuration: %s", err.Error())
			return
		}
	}
	if len(_workspacesinstancesClientToken) > 0 {
		input.ClientToken = aws.String(_workspacesinstancesClientToken)
	}
	if len(_workspacesinstancesTags) > 0 {
		if err := assignInputField(input, "Tags", _workspacesinstancesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWorkspaceInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified volume.
func workspacesinstances_DeleteVolume(cfg aws.Config, client *workspacesinstances.Client) {
	input := &workspacesinstances.DeleteVolumeInput{
		// VolumeId: *string, // Required
	}

	if len(_workspacesinstancesVolumeId) > 0 {
		input.VolumeId = aws.String(_workspacesinstancesVolumeId)
	}

	if resp, err := client.DeleteVolume(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified WorkSpace
// Usage of this API will result in deletion of the resource in question.
func workspacesinstances_DeleteWorkspaceInstance(cfg aws.Config, client *workspacesinstances.Client) {
	input := &workspacesinstances.DeleteWorkspaceInstanceInput{
		// WorkspaceInstanceId: *string, // Required
	}

	if len(_workspacesinstancesWorkspaceInstanceId) > 0 {
		input.WorkspaceInstanceId = aws.String(_workspacesinstancesWorkspaceInstanceId)
	}

	if resp, err := client.DeleteWorkspaceInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detaches a volume from a WorkSpace Instance.
func workspacesinstances_DisassociateVolume(cfg aws.Config, client *workspacesinstances.Client) {
	input := &workspacesinstances.DisassociateVolumeInput{
		// VolumeId: *string, // Required
		// WorkspaceInstanceId: *string, // Required
	}

	if len(_workspacesinstancesVolumeId) > 0 {
		input.VolumeId = aws.String(_workspacesinstancesVolumeId)
	}
	if len(_workspacesinstancesWorkspaceInstanceId) > 0 {
		input.WorkspaceInstanceId = aws.String(_workspacesinstancesWorkspaceInstanceId)
	}
	if len(_workspacesinstancesDevice) > 0 {
		input.Device = aws.String(_workspacesinstancesDevice)
	}
	if len(_workspacesinstancesDisassociateMode) > 0 {
		if err := assignInputField(input, "DisassociateMode", _workspacesinstancesDisassociateMode); err != nil {
			log.Errorf("invalid --disassociate-mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.DisassociateVolume(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specific WorkSpace Instance.
func workspacesinstances_GetWorkspaceInstance(cfg aws.Config, client *workspacesinstances.Client) {
	input := &workspacesinstances.GetWorkspaceInstanceInput{
		// WorkspaceInstanceId: *string, // Required
	}

	if len(_workspacesinstancesWorkspaceInstanceId) > 0 {
		input.WorkspaceInstanceId = aws.String(_workspacesinstancesWorkspaceInstanceId)
	}

	if resp, err := client.GetWorkspaceInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of instance types supported by Amazon WorkSpaces Instances,
// enabling precise workspace infrastructure configuration.
func workspacesinstances_ListInstanceTypes(cfg aws.Config, client *workspacesinstances.Client) {
	input := &workspacesinstances.ListInstanceTypesInput{}

	if len(_workspacesinstancesInstanceConfigurationFilter) > 0 {
		if err := assignInputField(input, "InstanceConfigurationFilter", _workspacesinstancesInstanceConfigurationFilter); err != nil {
			log.Errorf("invalid --instance-configuration-filter: %s", err.Error())
			return
		}
	}
	if len(_workspacesinstancesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspacesinstancesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspacesinstancesNextToken) > 0 {
		input.NextToken = aws.String(_workspacesinstancesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInstanceTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workspacesinstances.ListInstanceTypesOutput
	p := workspacesinstances.NewListInstanceTypesPaginator(client, input)
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

// Retrieves a list of AWS regions supported by Amazon WorkSpaces Instances,
// enabling region discovery for workspace deployments.
func workspacesinstances_ListRegions(cfg aws.Config, client *workspacesinstances.Client) {
	input := &workspacesinstances.ListRegionsInput{}

	if len(_workspacesinstancesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspacesinstancesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspacesinstancesNextToken) > 0 {
		input.NextToken = aws.String(_workspacesinstancesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRegions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workspacesinstances.ListRegionsOutput
	p := workspacesinstances.NewListRegionsPaginator(client, input)
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

// Retrieves tags for a WorkSpace Instance.
func workspacesinstances_ListTagsForResource(cfg aws.Config, client *workspacesinstances.Client) {
	input := &workspacesinstances.ListTagsForResourceInput{
		// WorkspaceInstanceId: *string, // Required
	}

	if len(_workspacesinstancesWorkspaceInstanceId) > 0 {
		input.WorkspaceInstanceId = aws.String(_workspacesinstancesWorkspaceInstanceId)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a collection of WorkSpaces Instances based on specified filters.
func workspacesinstances_ListWorkspaceInstances(cfg aws.Config, client *workspacesinstances.Client) {
	input := &workspacesinstances.ListWorkspaceInstancesInput{}

	if len(_workspacesinstancesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspacesinstancesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspacesinstancesNextToken) > 0 {
		input.NextToken = aws.String(_workspacesinstancesNextToken)
	}
	if len(_workspacesinstancesProvisionStates) > 0 {
		if err := assignInputField(input, "ProvisionStates", _workspacesinstancesProvisionStates); err != nil {
			log.Errorf("invalid --provision-states: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListWorkspaceInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workspacesinstances.ListWorkspaceInstancesOutput
	p := workspacesinstances.NewListWorkspaceInstancesPaginator(client, input)
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

// Adds tags to a WorkSpace Instance.
func workspacesinstances_TagResource(cfg aws.Config, client *workspacesinstances.Client) {
	input := &workspacesinstances.TagResourceInput{
		// Tags: []types.Tag, // Required
		// WorkspaceInstanceId: *string, // Required
	}

	if len(_workspacesinstancesTags) > 0 {
		if err := assignInputField(input, "Tags", _workspacesinstancesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_workspacesinstancesWorkspaceInstanceId) > 0 {
		input.WorkspaceInstanceId = aws.String(_workspacesinstancesWorkspaceInstanceId)
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes tags from a WorkSpace Instance.
func workspacesinstances_UntagResource(cfg aws.Config, client *workspacesinstances.Client) {
	input := &workspacesinstances.UntagResourceInput{
		// TagKeys: []string, // Required
		// WorkspaceInstanceId: *string, // Required
	}

	if len(_workspacesinstancesTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _workspacesinstancesTagKeys...)
	}
	if len(_workspacesinstancesWorkspaceInstanceId) > 0 {
		input.WorkspaceInstanceId = aws.String(_workspacesinstancesWorkspaceInstanceId)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_workspacesinstancesCmd)
	_workspacesinstancesCmd.Flags().SortFlags = false

	_workspacesinstancesCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_workspacesinstancesCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_workspacesinstancesCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_workspacesinstancesCmd.Flags().StringVarP(&_workspacesinstancesAvailabilityZone, "availability-zone", "", "", "Availability Zone")
	_workspacesinstancesCmd.Flags().StringVarP(&_workspacesinstancesBillingConfiguration, "billing-configuration", "", "", "Billing Configuration")
	_workspacesinstancesCmd.Flags().StringVarP(&_workspacesinstancesClientToken, "client-token", "", "", "Client Token")
	_workspacesinstancesCmd.Flags().StringVarP(&_workspacesinstancesDevice, "device", "", "", "Device")
	_workspacesinstancesCmd.Flags().StringVarP(&_workspacesinstancesDisassociateMode, "disassociate-mode", "", "", "Disassociate Mode")
	_workspacesinstancesCmd.Flags().StringVarP(&_workspacesinstancesEncrypted, "encrypted", "", "", "Encrypted")
	_workspacesinstancesCmd.Flags().StringVarP(&_workspacesinstancesInstanceConfigurationFilter, "instance-configuration-filter", "", "", "Instance Configuration Filter")
	_workspacesinstancesCmd.Flags().StringVarP(&_workspacesinstancesIops, "iops", "", "", "IOPS")
	_workspacesinstancesCmd.Flags().StringVarP(&_workspacesinstancesKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_workspacesinstancesCmd.Flags().StringVarP(&_workspacesinstancesManagedInstance, "managed-instance", "", "", "Managed Instance")
	_workspacesinstancesCmd.Flags().StringVarP(&_workspacesinstancesMaxResults, "max-results", "", "", "Max Results")
	_workspacesinstancesCmd.Flags().StringVarP(&_workspacesinstancesNextToken, "next-token", "", "", "Next Token")
	_workspacesinstancesCmd.Flags().StringVarP(&_workspacesinstancesProvisionStates, "provision-states", "", "", "Provision States")
	_workspacesinstancesCmd.Flags().StringVarP(&_workspacesinstancesSizeInGB, "size-in-gb", "", "", "Size In Gb")
	_workspacesinstancesCmd.Flags().StringVarP(&_workspacesinstancesSnapshotId, "snapshot-id", "", "", "Snapshot ID")
	_workspacesinstancesCmd.Flags().StringSliceVarP(&_workspacesinstancesTagKeys, "tag-keys", "", nil, "Tag Keys")
	_workspacesinstancesCmd.Flags().StringVarP(&_workspacesinstancesTagSpecifications, "tag-specifications", "", "", "Tag Specifications")
	_workspacesinstancesCmd.Flags().StringVarP(&_workspacesinstancesTags, "tags", "", "", "Tags")
	_workspacesinstancesCmd.Flags().StringVarP(&_workspacesinstancesThroughput, "throughput", "", "", "Throughput")
	_workspacesinstancesCmd.Flags().StringVarP(&_workspacesinstancesVolumeId, "volume-id", "", "", "Volume ID")
	_workspacesinstancesCmd.Flags().StringVarP(&_workspacesinstancesVolumeType, "volume-type", "", "", "Volume Type")
	_workspacesinstancesCmd.Flags().StringVarP(&_workspacesinstancesWorkspaceInstanceId, "workspace-instance-id", "", "", "Workspace Instance ID")

	_workspacesinstancesCmd.Flags().BoolVarP(&_workspacesinstancesAssociateVolume, "associate-volume", "", false, "Associate Volume")
	_workspacesinstancesCmd.Flags().BoolVarP(&_workspacesinstancesCreateVolume, "create-volume", "", false, "Create Volume")
	_workspacesinstancesCmd.Flags().BoolVarP(&_workspacesinstancesCreateWorkspaceInstance, "create-workspace-instance", "", false, "Create Workspace Instance")
	_workspacesinstancesCmd.Flags().BoolVarP(&_workspacesinstancesDeleteVolume, "delete-volume", "", false, "Delete Volume")
	_workspacesinstancesCmd.Flags().BoolVarP(&_workspacesinstancesDeleteWorkspaceInstance, "delete-workspace-instance", "", false, "Delete Workspace Instance")
	_workspacesinstancesCmd.Flags().BoolVarP(&_workspacesinstancesDisassociateVolume, "disassociate-volume", "", false, "Disassociate Volume")
	_workspacesinstancesCmd.Flags().BoolVarP(&_workspacesinstancesGetWorkspaceInstance, "get-workspace-instance", "", false, "Get Workspace Instance")
	_workspacesinstancesCmd.Flags().BoolVarP(&_workspacesinstancesListInstanceTypes, "list-instance-types", "", false, "List Instance Types")
	_workspacesinstancesCmd.Flags().BoolVarP(&_workspacesinstancesListRegions, "list-regions", "", false, "List Regions")
	_workspacesinstancesCmd.Flags().BoolVarP(&_workspacesinstancesListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_workspacesinstancesCmd.Flags().BoolVarP(&_workspacesinstancesListWorkspaceInstances, "list-workspace-instances", "", false, "List Workspace Instances")
	_workspacesinstancesCmd.Flags().BoolVarP(&_workspacesinstancesTagResource, "tag-resource", "", false, "Tag Resource")
	_workspacesinstancesCmd.Flags().BoolVarP(&_workspacesinstancesUntagResource, "untag-resource", "", false, "Untag Resource")

}
