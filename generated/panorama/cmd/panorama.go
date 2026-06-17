package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/panorama"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// panoramaCmd represents the panorama command
var _panoramaCmd = &cobra.Command{
	Use:   "panorama",
	Short: "AWS panorama CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := panorama.NewFromConfig(cfg)
		if _panoramaCreateApplicationInstance {
			panorama_CreateApplicationInstance(cfg, client)
			return
		}
		if _panoramaCreateJobForDevices {
			panorama_CreateJobForDevices(cfg, client)
			return
		}
		if _panoramaCreateNodeFromTemplateJob {
			panorama_CreateNodeFromTemplateJob(cfg, client)
			return
		}
		if _panoramaCreatePackage {
			panorama_CreatePackage(cfg, client)
			return
		}
		if _panoramaCreatePackageImportJob {
			panorama_CreatePackageImportJob(cfg, client)
			return
		}
		if _panoramaDeleteDevice {
			panorama_DeleteDevice(cfg, client)
			return
		}
		if _panoramaDeletePackage {
			panorama_DeletePackage(cfg, client)
			return
		}
		if _panoramaDeregisterPackageVersion {
			panorama_DeregisterPackageVersion(cfg, client)
			return
		}
		if _panoramaDescribeApplicationInstance {
			panorama_DescribeApplicationInstance(cfg, client)
			return
		}
		if _panoramaDescribeApplicationInstanceDetails {
			panorama_DescribeApplicationInstanceDetails(cfg, client)
			return
		}
		if _panoramaDescribeDevice {
			panorama_DescribeDevice(cfg, client)
			return
		}
		if _panoramaDescribeDeviceJob {
			panorama_DescribeDeviceJob(cfg, client)
			return
		}
		if _panoramaDescribeNode {
			panorama_DescribeNode(cfg, client)
			return
		}
		if _panoramaDescribeNodeFromTemplateJob {
			panorama_DescribeNodeFromTemplateJob(cfg, client)
			return
		}
		if _panoramaDescribePackage {
			panorama_DescribePackage(cfg, client)
			return
		}
		if _panoramaDescribePackageImportJob {
			panorama_DescribePackageImportJob(cfg, client)
			return
		}
		if _panoramaDescribePackageVersion {
			panorama_DescribePackageVersion(cfg, client)
			return
		}
		if _panoramaListApplicationInstanceDependencies {
			panorama_ListApplicationInstanceDependencies(cfg, client)
			return
		}
		if _panoramaListApplicationInstanceNodeInstances {
			panorama_ListApplicationInstanceNodeInstances(cfg, client)
			return
		}
		if _panoramaListApplicationInstances {
			panorama_ListApplicationInstances(cfg, client)
			return
		}
		if _panoramaListDevices {
			panorama_ListDevices(cfg, client)
			return
		}
		if _panoramaListDevicesJobs {
			panorama_ListDevicesJobs(cfg, client)
			return
		}
		if _panoramaListNodeFromTemplateJobs {
			panorama_ListNodeFromTemplateJobs(cfg, client)
			return
		}
		if _panoramaListNodes {
			panorama_ListNodes(cfg, client)
			return
		}
		if _panoramaListPackageImportJobs {
			panorama_ListPackageImportJobs(cfg, client)
			return
		}
		if _panoramaListPackages {
			panorama_ListPackages(cfg, client)
			return
		}
		if _panoramaListTagsForResource {
			panorama_ListTagsForResource(cfg, client)
			return
		}
		if _panoramaProvisionDevice {
			panorama_ProvisionDevice(cfg, client)
			return
		}
		if _panoramaRegisterPackageVersion {
			panorama_RegisterPackageVersion(cfg, client)
			return
		}
		if _panoramaRemoveApplicationInstance {
			panorama_RemoveApplicationInstance(cfg, client)
			return
		}
		if _panoramaSignalApplicationInstanceNodeInstances {
			panorama_SignalApplicationInstanceNodeInstances(cfg, client)
			return
		}
		if _panoramaTagResource {
			panorama_TagResource(cfg, client)
			return
		}
		if _panoramaUntagResource {
			panorama_UntagResource(cfg, client)
			return
		}
		if _panoramaUpdateDeviceMetadata {
			panorama_UpdateDeviceMetadata(cfg, client)
			return
		}

	},
}

var (
	_panoramaCreateApplicationInstance              bool
	_panoramaCreateJobForDevices                    bool
	_panoramaCreateNodeFromTemplateJob              bool
	_panoramaCreatePackage                          bool
	_panoramaCreatePackageImportJob                 bool
	_panoramaDeleteDevice                           bool
	_panoramaDeletePackage                          bool
	_panoramaDeregisterPackageVersion               bool
	_panoramaDescribeApplicationInstance            bool
	_panoramaDescribeApplicationInstanceDetails     bool
	_panoramaDescribeDevice                         bool
	_panoramaDescribeDeviceJob                      bool
	_panoramaDescribeNode                           bool
	_panoramaDescribeNodeFromTemplateJob            bool
	_panoramaDescribePackage                        bool
	_panoramaDescribePackageImportJob               bool
	_panoramaDescribePackageVersion                 bool
	_panoramaListApplicationInstanceDependencies    bool
	_panoramaListApplicationInstanceNodeInstances   bool
	_panoramaListApplicationInstances               bool
	_panoramaListDevices                            bool
	_panoramaListDevicesJobs                        bool
	_panoramaListNodeFromTemplateJobs               bool
	_panoramaListNodes                              bool
	_panoramaListPackageImportJobs                  bool
	_panoramaListPackages                           bool
	_panoramaListTagsForResource                    bool
	_panoramaProvisionDevice                        bool
	_panoramaRegisterPackageVersion                 bool
	_panoramaRemoveApplicationInstance              bool
	_panoramaSignalApplicationInstanceNodeInstances bool
	_panoramaTagResource                            bool
	_panoramaUntagResource                          bool
	_panoramaUpdateDeviceMetadata                   bool

	_panoramaApplicationInstanceId          string
	_panoramaApplicationInstanceIdToReplace string
	_panoramaCategory                       string
	_panoramaClientToken                    string
	_panoramaDefaultRuntimeContextDevice    string
	_panoramaDescription                    string
	_panoramaDeviceAggregatedStatusFilter   string
	_panoramaDeviceId                       string
	_panoramaDeviceIds                      []string
	_panoramaDeviceJobConfig                string
	_panoramaForceDelete                    string
	_panoramaInputConfig                    string
	_panoramaJobId                          string
	_panoramaJobTags                        string
	_panoramaJobType                        string
	_panoramaManifestOverridesPayload       string
	_panoramaManifestPayload                string
	_panoramaMarkLatest                     string
	_panoramaMaxResults                     string
	_panoramaName                           string
	_panoramaNameFilter                     string
	_panoramaNetworkingConfiguration        string
	_panoramaNextToken                      string
	_panoramaNodeDescription                string
	_panoramaNodeId                         string
	_panoramaNodeName                       string
	_panoramaNodeSignals                    string
	_panoramaOutputConfig                   string
	_panoramaOutputPackageName              string
	_panoramaOutputPackageVersion           string
	_panoramaOwnerAccount                   string
	_panoramaPackageId                      string
	_panoramaPackageName                    string
	_panoramaPackageVersion                 string
	_panoramaPatchVersion                   string
	_panoramaResourceArn                    string
	_panoramaRuntimeRoleArn                 string
	_panoramaSortBy                         string
	_panoramaSortOrder                      string
	_panoramaStatusFilter                   string
	_panoramaTagKeys                        []string
	_panoramaTags                           string
	_panoramaTemplateParameters             string
	_panoramaTemplateType                   string
	_panoramaUpdatedLatestPatchVersion      string
)

// Creates an application instance and deploys it to a device.
func panorama_CreateApplicationInstance(cfg aws.Config, client *panorama.Client) {
	input := &panorama.CreateApplicationInstanceInput{
		// DefaultRuntimeContextDevice: *string, // Required
		// ManifestPayload: types.ManifestPayload, // Required
	}

	if len(_panoramaDefaultRuntimeContextDevice) > 0 {
		input.DefaultRuntimeContextDevice = aws.String(_panoramaDefaultRuntimeContextDevice)
	}
	if len(_panoramaManifestPayload) > 0 {
		if err := assignInputField(input, "ManifestPayload", _panoramaManifestPayload); err != nil {
			log.Errorf("invalid --manifest-payload: %s", err.Error())
			return
		}
	}
	if len(_panoramaApplicationInstanceIdToReplace) > 0 {
		input.ApplicationInstanceIdToReplace = aws.String(_panoramaApplicationInstanceIdToReplace)
	}
	if len(_panoramaDescription) > 0 {
		input.Description = aws.String(_panoramaDescription)
	}
	if len(_panoramaManifestOverridesPayload) > 0 {
		if err := assignInputField(input, "ManifestOverridesPayload", _panoramaManifestOverridesPayload); err != nil {
			log.Errorf("invalid --manifest-overrides-payload: %s", err.Error())
			return
		}
	}
	if len(_panoramaName) > 0 {
		input.Name = aws.String(_panoramaName)
	}
	if len(_panoramaRuntimeRoleArn) > 0 {
		input.RuntimeRoleArn = aws.String(_panoramaRuntimeRoleArn)
	}
	if len(_panoramaTags) > 0 {
		if err := assignInputField(input, "Tags", _panoramaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateApplicationInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a job to run on a device. A job can update a device's software or
// reboot it.
func panorama_CreateJobForDevices(cfg aws.Config, client *panorama.Client) {
	input := &panorama.CreateJobForDevicesInput{
		// DeviceIds: []string, // Required
		// JobType: types.JobType, // Required
	}

	if len(_panoramaDeviceIds) > 0 {
		input.DeviceIds = append([]string(nil), _panoramaDeviceIds...)
	}
	if len(_panoramaJobType) > 0 {
		if err := assignInputField(input, "JobType", _panoramaJobType); err != nil {
			log.Errorf("invalid --job-type: %s", err.Error())
			return
		}
	}
	if len(_panoramaDeviceJobConfig) > 0 {
		if err := assignInputField(input, "DeviceJobConfig", _panoramaDeviceJobConfig); err != nil {
			log.Errorf("invalid --device-job-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateJobForDevices(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a camera stream node.
func panorama_CreateNodeFromTemplateJob(cfg aws.Config, client *panorama.Client) {
	input := &panorama.CreateNodeFromTemplateJobInput{
		// NodeName: *string, // Required
		// OutputPackageName: *string, // Required
		// OutputPackageVersion: *string, // Required
		// TemplateParameters: map[string]string, // Required
		// TemplateType: types.TemplateType, // Required
	}

	if len(_panoramaNodeName) > 0 {
		input.NodeName = aws.String(_panoramaNodeName)
	}
	if len(_panoramaOutputPackageName) > 0 {
		input.OutputPackageName = aws.String(_panoramaOutputPackageName)
	}
	if len(_panoramaOutputPackageVersion) > 0 {
		input.OutputPackageVersion = aws.String(_panoramaOutputPackageVersion)
	}
	if len(_panoramaTemplateParameters) > 0 {
		if err := assignInputField(input, "TemplateParameters", _panoramaTemplateParameters); err != nil {
			log.Errorf("invalid --template-parameters: %s", err.Error())
			return
		}
	}
	if len(_panoramaTemplateType) > 0 {
		if err := assignInputField(input, "TemplateType", _panoramaTemplateType); err != nil {
			log.Errorf("invalid --template-type: %s", err.Error())
			return
		}
	}
	if len(_panoramaJobTags) > 0 {
		if err := assignInputField(input, "JobTags", _panoramaJobTags); err != nil {
			log.Errorf("invalid --job-tags: %s", err.Error())
			return
		}
	}
	if len(_panoramaNodeDescription) > 0 {
		input.NodeDescription = aws.String(_panoramaNodeDescription)
	}

	if resp, err := client.CreateNodeFromTemplateJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a package and storage location in an Amazon S3 access point.
func panorama_CreatePackage(cfg aws.Config, client *panorama.Client) {
	input := &panorama.CreatePackageInput{
		// PackageName: *string, // Required
	}

	if len(_panoramaPackageName) > 0 {
		input.PackageName = aws.String(_panoramaPackageName)
	}
	if len(_panoramaTags) > 0 {
		if err := assignInputField(input, "Tags", _panoramaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports a node package.
func panorama_CreatePackageImportJob(cfg aws.Config, client *panorama.Client) {
	input := &panorama.CreatePackageImportJobInput{
		// ClientToken: *string, // Required
		// InputConfig: *types.PackageImportJobInputConfig, // Required
		// JobType: types.PackageImportJobType, // Required
		// OutputConfig: *types.PackageImportJobOutputConfig, // Required
	}

	if len(_panoramaClientToken) > 0 {
		input.ClientToken = aws.String(_panoramaClientToken)
	}
	if len(_panoramaInputConfig) > 0 {
		if err := assignInputField(input, "InputConfig", _panoramaInputConfig); err != nil {
			log.Errorf("invalid --input-config: %s", err.Error())
			return
		}
	}
	if len(_panoramaJobType) > 0 {
		if err := assignInputField(input, "JobType", _panoramaJobType); err != nil {
			log.Errorf("invalid --job-type: %s", err.Error())
			return
		}
	}
	if len(_panoramaOutputConfig) > 0 {
		if err := assignInputField(input, "OutputConfig", _panoramaOutputConfig); err != nil {
			log.Errorf("invalid --output-config: %s", err.Error())
			return
		}
	}
	if len(_panoramaJobTags) > 0 {
		if err := assignInputField(input, "JobTags", _panoramaJobTags); err != nil {
			log.Errorf("invalid --job-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePackageImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a device.
func panorama_DeleteDevice(cfg aws.Config, client *panorama.Client) {
	input := &panorama.DeleteDeviceInput{
		// DeviceId: *string, // Required
	}

	if len(_panoramaDeviceId) > 0 {
		input.DeviceId = aws.String(_panoramaDeviceId)
	}

	if resp, err := client.DeleteDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a package.
// To delete a package, you need permission to call s3:DeleteObject in addition to
// permissions for the AWS Panorama API.
func panorama_DeletePackage(cfg aws.Config, client *panorama.Client) {
	input := &panorama.DeletePackageInput{
		// PackageId: *string, // Required
	}

	if len(_panoramaPackageId) > 0 {
		input.PackageId = aws.String(_panoramaPackageId)
	}
	if len(_panoramaForceDelete) > 0 {
		if err := assignInputField(input, "ForceDelete", _panoramaForceDelete); err != nil {
			log.Errorf("invalid --force-delete: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeletePackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregisters a package version.
func panorama_DeregisterPackageVersion(cfg aws.Config, client *panorama.Client) {
	input := &panorama.DeregisterPackageVersionInput{
		// PackageId: *string, // Required
		// PackageVersion: *string, // Required
		// PatchVersion: *string, // Required
	}

	if len(_panoramaPackageId) > 0 {
		input.PackageId = aws.String(_panoramaPackageId)
	}
	if len(_panoramaPackageVersion) > 0 {
		input.PackageVersion = aws.String(_panoramaPackageVersion)
	}
	if len(_panoramaPatchVersion) > 0 {
		input.PatchVersion = aws.String(_panoramaPatchVersion)
	}
	if len(_panoramaOwnerAccount) > 0 {
		input.OwnerAccount = aws.String(_panoramaOwnerAccount)
	}
	if len(_panoramaUpdatedLatestPatchVersion) > 0 {
		input.UpdatedLatestPatchVersion = aws.String(_panoramaUpdatedLatestPatchVersion)
	}

	if resp, err := client.DeregisterPackageVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an application instance on a device.
func panorama_DescribeApplicationInstance(cfg aws.Config, client *panorama.Client) {
	input := &panorama.DescribeApplicationInstanceInput{
		// ApplicationInstanceId: *string, // Required
	}

	if len(_panoramaApplicationInstanceId) > 0 {
		input.ApplicationInstanceId = aws.String(_panoramaApplicationInstanceId)
	}

	if resp, err := client.DescribeApplicationInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an application instance's configuration manifest.
func panorama_DescribeApplicationInstanceDetails(cfg aws.Config, client *panorama.Client) {
	input := &panorama.DescribeApplicationInstanceDetailsInput{
		// ApplicationInstanceId: *string, // Required
	}

	if len(_panoramaApplicationInstanceId) > 0 {
		input.ApplicationInstanceId = aws.String(_panoramaApplicationInstanceId)
	}

	if resp, err := client.DescribeApplicationInstanceDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a device.
func panorama_DescribeDevice(cfg aws.Config, client *panorama.Client) {
	input := &panorama.DescribeDeviceInput{
		// DeviceId: *string, // Required
	}

	if len(_panoramaDeviceId) > 0 {
		input.DeviceId = aws.String(_panoramaDeviceId)
	}

	if resp, err := client.DescribeDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a device job.
func panorama_DescribeDeviceJob(cfg aws.Config, client *panorama.Client) {
	input := &panorama.DescribeDeviceJobInput{
		// JobId: *string, // Required
	}

	if len(_panoramaJobId) > 0 {
		input.JobId = aws.String(_panoramaJobId)
	}

	if resp, err := client.DescribeDeviceJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a node.
func panorama_DescribeNode(cfg aws.Config, client *panorama.Client) {
	input := &panorama.DescribeNodeInput{
		// NodeId: *string, // Required
	}

	if len(_panoramaNodeId) > 0 {
		input.NodeId = aws.String(_panoramaNodeId)
	}
	if len(_panoramaOwnerAccount) > 0 {
		input.OwnerAccount = aws.String(_panoramaOwnerAccount)
	}

	if resp, err := client.DescribeNode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a job to create a camera stream node.
func panorama_DescribeNodeFromTemplateJob(cfg aws.Config, client *panorama.Client) {
	input := &panorama.DescribeNodeFromTemplateJobInput{
		// JobId: *string, // Required
	}

	if len(_panoramaJobId) > 0 {
		input.JobId = aws.String(_panoramaJobId)
	}

	if resp, err := client.DescribeNodeFromTemplateJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a package.
func panorama_DescribePackage(cfg aws.Config, client *panorama.Client) {
	input := &panorama.DescribePackageInput{
		// PackageId: *string, // Required
	}

	if len(_panoramaPackageId) > 0 {
		input.PackageId = aws.String(_panoramaPackageId)
	}

	if resp, err := client.DescribePackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a package import job.
func panorama_DescribePackageImportJob(cfg aws.Config, client *panorama.Client) {
	input := &panorama.DescribePackageImportJobInput{
		// JobId: *string, // Required
	}

	if len(_panoramaJobId) > 0 {
		input.JobId = aws.String(_panoramaJobId)
	}

	if resp, err := client.DescribePackageImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a package version.
func panorama_DescribePackageVersion(cfg aws.Config, client *panorama.Client) {
	input := &panorama.DescribePackageVersionInput{
		// PackageId: *string, // Required
		// PackageVersion: *string, // Required
	}

	if len(_panoramaPackageId) > 0 {
		input.PackageId = aws.String(_panoramaPackageId)
	}
	if len(_panoramaPackageVersion) > 0 {
		input.PackageVersion = aws.String(_panoramaPackageVersion)
	}
	if len(_panoramaOwnerAccount) > 0 {
		input.OwnerAccount = aws.String(_panoramaOwnerAccount)
	}
	if len(_panoramaPatchVersion) > 0 {
		input.PatchVersion = aws.String(_panoramaPatchVersion)
	}

	if resp, err := client.DescribePackageVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of application instance dependencies.
func panorama_ListApplicationInstanceDependencies(cfg aws.Config, client *panorama.Client) {
	input := &panorama.ListApplicationInstanceDependenciesInput{
		// ApplicationInstanceId: *string, // Required
	}

	if len(_panoramaApplicationInstanceId) > 0 {
		input.ApplicationInstanceId = aws.String(_panoramaApplicationInstanceId)
	}
	if len(_panoramaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _panoramaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_panoramaNextToken) > 0 {
		input.NextToken = aws.String(_panoramaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApplicationInstanceDependencies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*panorama.ListApplicationInstanceDependenciesOutput
	p := panorama.NewListApplicationInstanceDependenciesPaginator(client, input)
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

// Returns a list of application node instances.
func panorama_ListApplicationInstanceNodeInstances(cfg aws.Config, client *panorama.Client) {
	input := &panorama.ListApplicationInstanceNodeInstancesInput{
		// ApplicationInstanceId: *string, // Required
	}

	if len(_panoramaApplicationInstanceId) > 0 {
		input.ApplicationInstanceId = aws.String(_panoramaApplicationInstanceId)
	}
	if len(_panoramaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _panoramaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_panoramaNextToken) > 0 {
		input.NextToken = aws.String(_panoramaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApplicationInstanceNodeInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*panorama.ListApplicationInstanceNodeInstancesOutput
	p := panorama.NewListApplicationInstanceNodeInstancesPaginator(client, input)
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

// Returns a list of application instances.
func panorama_ListApplicationInstances(cfg aws.Config, client *panorama.Client) {
	input := &panorama.ListApplicationInstancesInput{}

	if len(_panoramaDeviceId) > 0 {
		input.DeviceId = aws.String(_panoramaDeviceId)
	}
	if len(_panoramaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _panoramaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_panoramaNextToken) > 0 {
		input.NextToken = aws.String(_panoramaNextToken)
	}
	if len(_panoramaStatusFilter) > 0 {
		if err := assignInputField(input, "StatusFilter", _panoramaStatusFilter); err != nil {
			log.Errorf("invalid --status-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListApplicationInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*panorama.ListApplicationInstancesOutput
	p := panorama.NewListApplicationInstancesPaginator(client, input)
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

// Returns a list of devices.
func panorama_ListDevices(cfg aws.Config, client *panorama.Client) {
	input := &panorama.ListDevicesInput{}

	if len(_panoramaDeviceAggregatedStatusFilter) > 0 {
		if err := assignInputField(input, "DeviceAggregatedStatusFilter", _panoramaDeviceAggregatedStatusFilter); err != nil {
			log.Errorf("invalid --device-aggregated-status-filter: %s", err.Error())
			return
		}
	}
	if len(_panoramaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _panoramaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_panoramaNameFilter) > 0 {
		input.NameFilter = aws.String(_panoramaNameFilter)
	}
	if len(_panoramaNextToken) > 0 {
		input.NextToken = aws.String(_panoramaNextToken)
	}
	if len(_panoramaSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _panoramaSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_panoramaSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _panoramaSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDevices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*panorama.ListDevicesOutput
	p := panorama.NewListDevicesPaginator(client, input)
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

// Returns a list of jobs.
func panorama_ListDevicesJobs(cfg aws.Config, client *panorama.Client) {
	input := &panorama.ListDevicesJobsInput{}

	if len(_panoramaDeviceId) > 0 {
		input.DeviceId = aws.String(_panoramaDeviceId)
	}
	if len(_panoramaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _panoramaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_panoramaNextToken) > 0 {
		input.NextToken = aws.String(_panoramaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDevicesJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*panorama.ListDevicesJobsOutput
	p := panorama.NewListDevicesJobsPaginator(client, input)
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

// Returns a list of camera stream node jobs.
func panorama_ListNodeFromTemplateJobs(cfg aws.Config, client *panorama.Client) {
	input := &panorama.ListNodeFromTemplateJobsInput{}

	if len(_panoramaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _panoramaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_panoramaNextToken) > 0 {
		input.NextToken = aws.String(_panoramaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListNodeFromTemplateJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*panorama.ListNodeFromTemplateJobsOutput
	p := panorama.NewListNodeFromTemplateJobsPaginator(client, input)
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

// Returns a list of nodes.
func panorama_ListNodes(cfg aws.Config, client *panorama.Client) {
	input := &panorama.ListNodesInput{}

	if len(_panoramaCategory) > 0 {
		if err := assignInputField(input, "Category", _panoramaCategory); err != nil {
			log.Errorf("invalid --category: %s", err.Error())
			return
		}
	}
	if len(_panoramaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _panoramaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_panoramaNextToken) > 0 {
		input.NextToken = aws.String(_panoramaNextToken)
	}
	if len(_panoramaOwnerAccount) > 0 {
		input.OwnerAccount = aws.String(_panoramaOwnerAccount)
	}
	if len(_panoramaPackageName) > 0 {
		input.PackageName = aws.String(_panoramaPackageName)
	}
	if len(_panoramaPackageVersion) > 0 {
		input.PackageVersion = aws.String(_panoramaPackageVersion)
	}
	if len(_panoramaPatchVersion) > 0 {
		input.PatchVersion = aws.String(_panoramaPatchVersion)
	}

	if disablePaginator() {
		if resp, err := client.ListNodes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*panorama.ListNodesOutput
	p := panorama.NewListNodesPaginator(client, input)
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

// Returns a list of package import jobs.
func panorama_ListPackageImportJobs(cfg aws.Config, client *panorama.Client) {
	input := &panorama.ListPackageImportJobsInput{}

	if len(_panoramaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _panoramaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_panoramaNextToken) > 0 {
		input.NextToken = aws.String(_panoramaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPackageImportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*panorama.ListPackageImportJobsOutput
	p := panorama.NewListPackageImportJobsPaginator(client, input)
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

// Returns a list of packages.
func panorama_ListPackages(cfg aws.Config, client *panorama.Client) {
	input := &panorama.ListPackagesInput{}

	if len(_panoramaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _panoramaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_panoramaNextToken) > 0 {
		input.NextToken = aws.String(_panoramaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPackages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*panorama.ListPackagesOutput
	p := panorama.NewListPackagesPaginator(client, input)
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

// Returns a list of tags for a resource.
func panorama_ListTagsForResource(cfg aws.Config, client *panorama.Client) {
	input := &panorama.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_panoramaResourceArn) > 0 {
		input.ResourceArn = aws.String(_panoramaResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a device and returns a configuration archive. The configuration archive
// is a ZIP file that contains a provisioning certificate that is valid for 5
// minutes. Name the configuration archive certificates-omni_device-name.zip and
// transfer it to the device within 5 minutes. Use the included USB storage device
// and connect it to the USB 3.0 port next to the HDMI output.
func panorama_ProvisionDevice(cfg aws.Config, client *panorama.Client) {
	input := &panorama.ProvisionDeviceInput{
		// Name: *string, // Required
	}

	if len(_panoramaName) > 0 {
		input.Name = aws.String(_panoramaName)
	}
	if len(_panoramaDescription) > 0 {
		input.Description = aws.String(_panoramaDescription)
	}
	if len(_panoramaNetworkingConfiguration) > 0 {
		if err := assignInputField(input, "NetworkingConfiguration", _panoramaNetworkingConfiguration); err != nil {
			log.Errorf("invalid --networking-configuration: %s", err.Error())
			return
		}
	}
	if len(_panoramaTags) > 0 {
		if err := assignInputField(input, "Tags", _panoramaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.ProvisionDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers a package version.
func panorama_RegisterPackageVersion(cfg aws.Config, client *panorama.Client) {
	input := &panorama.RegisterPackageVersionInput{
		// PackageId: *string, // Required
		// PackageVersion: *string, // Required
		// PatchVersion: *string, // Required
	}

	if len(_panoramaPackageId) > 0 {
		input.PackageId = aws.String(_panoramaPackageId)
	}
	if len(_panoramaPackageVersion) > 0 {
		input.PackageVersion = aws.String(_panoramaPackageVersion)
	}
	if len(_panoramaPatchVersion) > 0 {
		input.PatchVersion = aws.String(_panoramaPatchVersion)
	}
	if len(_panoramaMarkLatest) > 0 {
		if err := assignInputField(input, "MarkLatest", _panoramaMarkLatest); err != nil {
			log.Errorf("invalid --mark-latest: %s", err.Error())
			return
		}
	}
	if len(_panoramaOwnerAccount) > 0 {
		input.OwnerAccount = aws.String(_panoramaOwnerAccount)
	}

	if resp, err := client.RegisterPackageVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes an application instance.
func panorama_RemoveApplicationInstance(cfg aws.Config, client *panorama.Client) {
	input := &panorama.RemoveApplicationInstanceInput{
		// ApplicationInstanceId: *string, // Required
	}

	if len(_panoramaApplicationInstanceId) > 0 {
		input.ApplicationInstanceId = aws.String(_panoramaApplicationInstanceId)
	}

	if resp, err := client.RemoveApplicationInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Signal camera nodes to stop or resume.
func panorama_SignalApplicationInstanceNodeInstances(cfg aws.Config, client *panorama.Client) {
	input := &panorama.SignalApplicationInstanceNodeInstancesInput{
		// ApplicationInstanceId: *string, // Required
		// NodeSignals: []types.NodeSignal, // Required
	}

	if len(_panoramaApplicationInstanceId) > 0 {
		input.ApplicationInstanceId = aws.String(_panoramaApplicationInstanceId)
	}
	if len(_panoramaNodeSignals) > 0 {
		if err := assignInputField(input, "NodeSignals", _panoramaNodeSignals); err != nil {
			log.Errorf("invalid --node-signals: %s", err.Error())
			return
		}
	}

	if resp, err := client.SignalApplicationInstanceNodeInstances(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tags a resource.
func panorama_TagResource(cfg aws.Config, client *panorama.Client) {
	input := &panorama.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_panoramaResourceArn) > 0 {
		input.ResourceArn = aws.String(_panoramaResourceArn)
	}
	if len(_panoramaTags) > 0 {
		if err := assignInputField(input, "Tags", _panoramaTags); err != nil {
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

// Removes tags from a resource.
func panorama_UntagResource(cfg aws.Config, client *panorama.Client) {
	input := &panorama.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_panoramaResourceArn) > 0 {
		input.ResourceArn = aws.String(_panoramaResourceArn)
	}
	if len(_panoramaTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _panoramaTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a device's metadata.
func panorama_UpdateDeviceMetadata(cfg aws.Config, client *panorama.Client) {
	input := &panorama.UpdateDeviceMetadataInput{
		// DeviceId: *string, // Required
	}

	if len(_panoramaDeviceId) > 0 {
		input.DeviceId = aws.String(_panoramaDeviceId)
	}
	if len(_panoramaDescription) > 0 {
		input.Description = aws.String(_panoramaDescription)
	}

	if resp, err := client.UpdateDeviceMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_panoramaCmd)
	_panoramaCmd.Flags().SortFlags = false

	_panoramaCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_panoramaCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_panoramaCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_panoramaCmd.Flags().StringVarP(&_panoramaApplicationInstanceId, "application-instance-id", "", "", "Application Instance ID")
	_panoramaCmd.Flags().StringVarP(&_panoramaApplicationInstanceIdToReplace, "application-instance-id-to-replace", "", "", "Application Instance ID To Replace")
	_panoramaCmd.Flags().StringVarP(&_panoramaCategory, "category", "", "", "Category")
	_panoramaCmd.Flags().StringVarP(&_panoramaClientToken, "client-token", "", "", "Client Token")
	_panoramaCmd.Flags().StringVarP(&_panoramaDefaultRuntimeContextDevice, "default-runtime-context-device", "", "", "Default Runtime Context Device")
	_panoramaCmd.Flags().StringVarP(&_panoramaDescription, "description", "", "", "Description")
	_panoramaCmd.Flags().StringVarP(&_panoramaDeviceAggregatedStatusFilter, "device-aggregated-status-filter", "", "", "Device Aggregated Status Filter")
	_panoramaCmd.Flags().StringVarP(&_panoramaDeviceId, "device-id", "", "", "Device ID")
	_panoramaCmd.Flags().StringSliceVarP(&_panoramaDeviceIds, "device-ids", "", nil, "Device Ids")
	_panoramaCmd.Flags().StringVarP(&_panoramaDeviceJobConfig, "device-job-config", "", "", "Device Job Config")
	_panoramaCmd.Flags().StringVarP(&_panoramaForceDelete, "force-delete", "", "", "Force Delete")
	_panoramaCmd.Flags().StringVarP(&_panoramaInputConfig, "input-config", "", "", "Input Config")
	_panoramaCmd.Flags().StringVarP(&_panoramaJobId, "job-id", "", "", "Job ID")
	_panoramaCmd.Flags().StringVarP(&_panoramaJobTags, "job-tags", "", "", "Job Tags")
	_panoramaCmd.Flags().StringVarP(&_panoramaJobType, "job-type", "", "", "Job Type")
	_panoramaCmd.Flags().StringVarP(&_panoramaManifestOverridesPayload, "manifest-overrides-payload", "", "", "Manifest Overrides Payload")
	_panoramaCmd.Flags().StringVarP(&_panoramaManifestPayload, "manifest-payload", "", "", "Manifest Payload")
	_panoramaCmd.Flags().StringVarP(&_panoramaMarkLatest, "mark-latest", "", "", "Mark Latest")
	_panoramaCmd.Flags().StringVarP(&_panoramaMaxResults, "max-results", "", "", "Max Results")
	_panoramaCmd.Flags().StringVarP(&_panoramaName, "name", "", "", "Name")
	_panoramaCmd.Flags().StringVarP(&_panoramaNameFilter, "name-filter", "", "", "Name Filter")
	_panoramaCmd.Flags().StringVarP(&_panoramaNetworkingConfiguration, "networking-configuration", "", "", "Networking Configuration")
	_panoramaCmd.Flags().StringVarP(&_panoramaNextToken, "next-token", "", "", "Next Token")
	_panoramaCmd.Flags().StringVarP(&_panoramaNodeDescription, "node-description", "", "", "Node Description")
	_panoramaCmd.Flags().StringVarP(&_panoramaNodeId, "node-id", "", "", "Node ID")
	_panoramaCmd.Flags().StringVarP(&_panoramaNodeName, "node-name", "", "", "Node Name")
	_panoramaCmd.Flags().StringVarP(&_panoramaNodeSignals, "node-signals", "", "", "Node Signals")
	_panoramaCmd.Flags().StringVarP(&_panoramaOutputConfig, "output-config", "", "", "Output Config")
	_panoramaCmd.Flags().StringVarP(&_panoramaOutputPackageName, "output-package-name", "", "", "Output Package Name")
	_panoramaCmd.Flags().StringVarP(&_panoramaOutputPackageVersion, "output-package-version", "", "", "Output Package Version")
	_panoramaCmd.Flags().StringVarP(&_panoramaOwnerAccount, "owner-account", "", "", "Owner Account")
	_panoramaCmd.Flags().StringVarP(&_panoramaPackageId, "package-id", "", "", "Package ID")
	_panoramaCmd.Flags().StringVarP(&_panoramaPackageName, "package-name", "", "", "Package Name")
	_panoramaCmd.Flags().StringVarP(&_panoramaPackageVersion, "package-version", "", "", "Package Version")
	_panoramaCmd.Flags().StringVarP(&_panoramaPatchVersion, "patch-version", "", "", "Patch Version")
	_panoramaCmd.Flags().StringVarP(&_panoramaResourceArn, "resource-arn", "", "", "Resource ARN")
	_panoramaCmd.Flags().StringVarP(&_panoramaRuntimeRoleArn, "runtime-role-arn", "", "", "Runtime Role ARN")
	_panoramaCmd.Flags().StringVarP(&_panoramaSortBy, "sort-by", "", "", "Sort By")
	_panoramaCmd.Flags().StringVarP(&_panoramaSortOrder, "sort-order", "", "", "Sort Order")
	_panoramaCmd.Flags().StringVarP(&_panoramaStatusFilter, "status-filter", "", "", "Status Filter")
	_panoramaCmd.Flags().StringSliceVarP(&_panoramaTagKeys, "tag-keys", "", nil, "Tag Keys")
	_panoramaCmd.Flags().StringVarP(&_panoramaTags, "tags", "", "", "Tags")
	_panoramaCmd.Flags().StringVarP(&_panoramaTemplateParameters, "template-parameters", "", "", "Template Parameters")
	_panoramaCmd.Flags().StringVarP(&_panoramaTemplateType, "template-type", "", "", "Template Type")
	_panoramaCmd.Flags().StringVarP(&_panoramaUpdatedLatestPatchVersion, "updated-latest-patch-version", "", "", "Updated Latest Patch Version")

	_panoramaCmd.Flags().BoolVarP(&_panoramaCreateApplicationInstance, "create-application-instance", "", false, "Create Application Instance")
	_panoramaCmd.Flags().BoolVarP(&_panoramaCreateJobForDevices, "create-job-for-devices", "", false, "Create Job For Devices")
	_panoramaCmd.Flags().BoolVarP(&_panoramaCreateNodeFromTemplateJob, "create-node-from-template-job", "", false, "Create Node From Template Job")
	_panoramaCmd.Flags().BoolVarP(&_panoramaCreatePackage, "create-package", "", false, "Create Package")
	_panoramaCmd.Flags().BoolVarP(&_panoramaCreatePackageImportJob, "create-package-import-job", "", false, "Create Package Import Job")
	_panoramaCmd.Flags().BoolVarP(&_panoramaDeleteDevice, "delete-device", "", false, "Delete Device")
	_panoramaCmd.Flags().BoolVarP(&_panoramaDeletePackage, "delete-package", "", false, "Delete Package")
	_panoramaCmd.Flags().BoolVarP(&_panoramaDeregisterPackageVersion, "deregister-package-version", "", false, "Deregister Package Version")
	_panoramaCmd.Flags().BoolVarP(&_panoramaDescribeApplicationInstance, "describe-application-instance", "", false, "Describe Application Instance")
	_panoramaCmd.Flags().BoolVarP(&_panoramaDescribeApplicationInstanceDetails, "describe-application-instance-details", "", false, "Describe Application Instance Details")
	_panoramaCmd.Flags().BoolVarP(&_panoramaDescribeDevice, "describe-device", "", false, "Describe Device")
	_panoramaCmd.Flags().BoolVarP(&_panoramaDescribeDeviceJob, "describe-device-job", "", false, "Describe Device Job")
	_panoramaCmd.Flags().BoolVarP(&_panoramaDescribeNode, "describe-node", "", false, "Describe Node")
	_panoramaCmd.Flags().BoolVarP(&_panoramaDescribeNodeFromTemplateJob, "describe-node-from-template-job", "", false, "Describe Node From Template Job")
	_panoramaCmd.Flags().BoolVarP(&_panoramaDescribePackage, "describe-package", "", false, "Describe Package")
	_panoramaCmd.Flags().BoolVarP(&_panoramaDescribePackageImportJob, "describe-package-import-job", "", false, "Describe Package Import Job")
	_panoramaCmd.Flags().BoolVarP(&_panoramaDescribePackageVersion, "describe-package-version", "", false, "Describe Package Version")
	_panoramaCmd.Flags().BoolVarP(&_panoramaListApplicationInstanceDependencies, "list-application-instance-dependencies", "", false, "List Application Instance Dependencies")
	_panoramaCmd.Flags().BoolVarP(&_panoramaListApplicationInstanceNodeInstances, "list-application-instance-node-instances", "", false, "List Application Instance Node Instances")
	_panoramaCmd.Flags().BoolVarP(&_panoramaListApplicationInstances, "list-application-instances", "", false, "List Application Instances")
	_panoramaCmd.Flags().BoolVarP(&_panoramaListDevices, "list-devices", "", false, "List Devices")
	_panoramaCmd.Flags().BoolVarP(&_panoramaListDevicesJobs, "list-devices-jobs", "", false, "List Devices Jobs")
	_panoramaCmd.Flags().BoolVarP(&_panoramaListNodeFromTemplateJobs, "list-node-from-template-jobs", "", false, "List Node From Template Jobs")
	_panoramaCmd.Flags().BoolVarP(&_panoramaListNodes, "list-nodes", "", false, "List Nodes")
	_panoramaCmd.Flags().BoolVarP(&_panoramaListPackageImportJobs, "list-package-import-jobs", "", false, "List Package Import Jobs")
	_panoramaCmd.Flags().BoolVarP(&_panoramaListPackages, "list-packages", "", false, "List Packages")
	_panoramaCmd.Flags().BoolVarP(&_panoramaListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_panoramaCmd.Flags().BoolVarP(&_panoramaProvisionDevice, "provision-device", "", false, "Provision Device")
	_panoramaCmd.Flags().BoolVarP(&_panoramaRegisterPackageVersion, "register-package-version", "", false, "Register Package Version")
	_panoramaCmd.Flags().BoolVarP(&_panoramaRemoveApplicationInstance, "remove-application-instance", "", false, "Remove Application Instance")
	_panoramaCmd.Flags().BoolVarP(&_panoramaSignalApplicationInstanceNodeInstances, "signal-application-instance-node-instances", "", false, "Signal Application Instance Node Instances")
	_panoramaCmd.Flags().BoolVarP(&_panoramaTagResource, "tag-resource", "", false, "Tag Resource")
	_panoramaCmd.Flags().BoolVarP(&_panoramaUntagResource, "untag-resource", "", false, "Untag Resource")
	_panoramaCmd.Flags().BoolVarP(&_panoramaUpdateDeviceMetadata, "update-device-metadata", "", false, "Update Device Metadata")

}
