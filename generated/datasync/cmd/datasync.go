package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/datasync"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// datasyncCmd represents the datasync command
var _datasyncCmd = &cobra.Command{
	Use:   "datasync",
	Short: "AWS datasync CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := datasync.NewFromConfig(cfg)
		if _datasyncCancelTaskExecution {
			datasync_CancelTaskExecution(cfg, client)
			return
		}
		if _datasyncCreateAgent {
			datasync_CreateAgent(cfg, client)
			return
		}
		if _datasyncCreateLocationAzureBlob {
			datasync_CreateLocationAzureBlob(cfg, client)
			return
		}
		if _datasyncCreateLocationEfs {
			datasync_CreateLocationEfs(cfg, client)
			return
		}
		if _datasyncCreateLocationFsxLustre {
			datasync_CreateLocationFsxLustre(cfg, client)
			return
		}
		if _datasyncCreateLocationFsxOntap {
			datasync_CreateLocationFsxOntap(cfg, client)
			return
		}
		if _datasyncCreateLocationFsxOpenZfs {
			datasync_CreateLocationFsxOpenZfs(cfg, client)
			return
		}
		if _datasyncCreateLocationFsxWindows {
			datasync_CreateLocationFsxWindows(cfg, client)
			return
		}
		if _datasyncCreateLocationHdfs {
			datasync_CreateLocationHdfs(cfg, client)
			return
		}
		if _datasyncCreateLocationNfs {
			datasync_CreateLocationNfs(cfg, client)
			return
		}
		if _datasyncCreateLocationObjectStorage {
			datasync_CreateLocationObjectStorage(cfg, client)
			return
		}
		if _datasyncCreateLocationS3 {
			datasync_CreateLocationS3(cfg, client)
			return
		}
		if _datasyncCreateLocationSmb {
			datasync_CreateLocationSmb(cfg, client)
			return
		}
		if _datasyncCreateTask {
			datasync_CreateTask(cfg, client)
			return
		}
		if _datasyncDeleteAgent {
			datasync_DeleteAgent(cfg, client)
			return
		}
		if _datasyncDeleteLocation {
			datasync_DeleteLocation(cfg, client)
			return
		}
		if _datasyncDeleteTask {
			datasync_DeleteTask(cfg, client)
			return
		}
		if _datasyncDescribeAgent {
			datasync_DescribeAgent(cfg, client)
			return
		}
		if _datasyncDescribeLocationAzureBlob {
			datasync_DescribeLocationAzureBlob(cfg, client)
			return
		}
		if _datasyncDescribeLocationEfs {
			datasync_DescribeLocationEfs(cfg, client)
			return
		}
		if _datasyncDescribeLocationFsxLustre {
			datasync_DescribeLocationFsxLustre(cfg, client)
			return
		}
		if _datasyncDescribeLocationFsxOntap {
			datasync_DescribeLocationFsxOntap(cfg, client)
			return
		}
		if _datasyncDescribeLocationFsxOpenZfs {
			datasync_DescribeLocationFsxOpenZfs(cfg, client)
			return
		}
		if _datasyncDescribeLocationFsxWindows {
			datasync_DescribeLocationFsxWindows(cfg, client)
			return
		}
		if _datasyncDescribeLocationHdfs {
			datasync_DescribeLocationHdfs(cfg, client)
			return
		}
		if _datasyncDescribeLocationNfs {
			datasync_DescribeLocationNfs(cfg, client)
			return
		}
		if _datasyncDescribeLocationObjectStorage {
			datasync_DescribeLocationObjectStorage(cfg, client)
			return
		}
		if _datasyncDescribeLocationS3 {
			datasync_DescribeLocationS3(cfg, client)
			return
		}
		if _datasyncDescribeLocationSmb {
			datasync_DescribeLocationSmb(cfg, client)
			return
		}
		if _datasyncDescribeTask {
			datasync_DescribeTask(cfg, client)
			return
		}
		if _datasyncDescribeTaskExecution {
			datasync_DescribeTaskExecution(cfg, client)
			return
		}
		if _datasyncListAgents {
			datasync_ListAgents(cfg, client)
			return
		}
		if _datasyncListLocations {
			datasync_ListLocations(cfg, client)
			return
		}
		if _datasyncListTagsForResource {
			datasync_ListTagsForResource(cfg, client)
			return
		}
		if _datasyncListTaskExecutions {
			datasync_ListTaskExecutions(cfg, client)
			return
		}
		if _datasyncListTasks {
			datasync_ListTasks(cfg, client)
			return
		}
		if _datasyncStartTaskExecution {
			datasync_StartTaskExecution(cfg, client)
			return
		}
		if _datasyncTagResource {
			datasync_TagResource(cfg, client)
			return
		}
		if _datasyncUntagResource {
			datasync_UntagResource(cfg, client)
			return
		}
		if _datasyncUpdateAgent {
			datasync_UpdateAgent(cfg, client)
			return
		}
		if _datasyncUpdateLocationAzureBlob {
			datasync_UpdateLocationAzureBlob(cfg, client)
			return
		}
		if _datasyncUpdateLocationEfs {
			datasync_UpdateLocationEfs(cfg, client)
			return
		}
		if _datasyncUpdateLocationFsxLustre {
			datasync_UpdateLocationFsxLustre(cfg, client)
			return
		}
		if _datasyncUpdateLocationFsxOntap {
			datasync_UpdateLocationFsxOntap(cfg, client)
			return
		}
		if _datasyncUpdateLocationFsxOpenZfs {
			datasync_UpdateLocationFsxOpenZfs(cfg, client)
			return
		}
		if _datasyncUpdateLocationFsxWindows {
			datasync_UpdateLocationFsxWindows(cfg, client)
			return
		}
		if _datasyncUpdateLocationHdfs {
			datasync_UpdateLocationHdfs(cfg, client)
			return
		}
		if _datasyncUpdateLocationNfs {
			datasync_UpdateLocationNfs(cfg, client)
			return
		}
		if _datasyncUpdateLocationObjectStorage {
			datasync_UpdateLocationObjectStorage(cfg, client)
			return
		}
		if _datasyncUpdateLocationS3 {
			datasync_UpdateLocationS3(cfg, client)
			return
		}
		if _datasyncUpdateLocationSmb {
			datasync_UpdateLocationSmb(cfg, client)
			return
		}
		if _datasyncUpdateTask {
			datasync_UpdateTask(cfg, client)
			return
		}
		if _datasyncUpdateTaskExecution {
			datasync_UpdateTaskExecution(cfg, client)
			return
		}

	},
}

var (
	_datasyncCancelTaskExecution           bool
	_datasyncCreateAgent                   bool
	_datasyncCreateLocationAzureBlob       bool
	_datasyncCreateLocationEfs             bool
	_datasyncCreateLocationFsxLustre       bool
	_datasyncCreateLocationFsxOntap        bool
	_datasyncCreateLocationFsxOpenZfs      bool
	_datasyncCreateLocationFsxWindows      bool
	_datasyncCreateLocationHdfs            bool
	_datasyncCreateLocationNfs             bool
	_datasyncCreateLocationObjectStorage   bool
	_datasyncCreateLocationS3              bool
	_datasyncCreateLocationSmb             bool
	_datasyncCreateTask                    bool
	_datasyncDeleteAgent                   bool
	_datasyncDeleteLocation                bool
	_datasyncDeleteTask                    bool
	_datasyncDescribeAgent                 bool
	_datasyncDescribeLocationAzureBlob     bool
	_datasyncDescribeLocationEfs           bool
	_datasyncDescribeLocationFsxLustre     bool
	_datasyncDescribeLocationFsxOntap      bool
	_datasyncDescribeLocationFsxOpenZfs    bool
	_datasyncDescribeLocationFsxWindows    bool
	_datasyncDescribeLocationHdfs          bool
	_datasyncDescribeLocationNfs           bool
	_datasyncDescribeLocationObjectStorage bool
	_datasyncDescribeLocationS3            bool
	_datasyncDescribeLocationSmb           bool
	_datasyncDescribeTask                  bool
	_datasyncDescribeTaskExecution         bool
	_datasyncListAgents                    bool
	_datasyncListLocations                 bool
	_datasyncListTagsForResource           bool
	_datasyncListTaskExecutions            bool
	_datasyncListTasks                     bool
	_datasyncStartTaskExecution            bool
	_datasyncTagResource                   bool
	_datasyncUntagResource                 bool
	_datasyncUpdateAgent                   bool
	_datasyncUpdateLocationAzureBlob       bool
	_datasyncUpdateLocationEfs             bool
	_datasyncUpdateLocationFsxLustre       bool
	_datasyncUpdateLocationFsxOntap        bool
	_datasyncUpdateLocationFsxOpenZfs      bool
	_datasyncUpdateLocationFsxWindows      bool
	_datasyncUpdateLocationHdfs            bool
	_datasyncUpdateLocationNfs             bool
	_datasyncUpdateLocationObjectStorage   bool
	_datasyncUpdateLocationS3              bool
	_datasyncUpdateLocationSmb             bool
	_datasyncUpdateTask                    bool
	_datasyncUpdateTaskExecution           bool

	_datasyncAccessKey                string
	_datasyncAccessPointArn           string
	_datasyncAccessTier               string
	_datasyncActivationKey            string
	_datasyncAgentArn                 string
	_datasyncAgentArns                []string
	_datasyncAgentName                string
	_datasyncAuthenticationType       string
	_datasyncBlobType                 string
	_datasyncBlockSize                string
	_datasyncBucketName               string
	_datasyncCloudWatchLogGroupArn    string
	_datasyncCmkSecretConfig          string
	_datasyncContainerUrl             string
	_datasyncCustomSecretConfig       string
	_datasyncDestinationLocationArn   string
	_datasyncDnsIpAddresses           []string
	_datasyncDomain                   string
	_datasyncEc2Config                string
	_datasyncEfsFilesystemArn         string
	_datasyncExcludes                 string
	_datasyncFileSystemAccessRoleArn  string
	_datasyncFilters                  string
	_datasyncFsxFilesystemArn         string
	_datasyncInTransitEncryption      string
	_datasyncIncludes                 string
	_datasyncKerberosKeytab           string
	_datasyncKerberosKrb5Conf         string
	_datasyncKerberosPrincipal        string
	_datasyncKeys                     []string
	_datasyncKmsKeyProviderUri        string
	_datasyncLocationArn              string
	_datasyncManifestConfig           string
	_datasyncMaxResults               string
	_datasyncMountOptions             string
	_datasyncName                     string
	_datasyncNameNodes                string
	_datasyncNextToken                string
	_datasyncOnPremConfig             string
	_datasyncOptions                  string
	_datasyncOverrideOptions          string
	_datasyncPassword                 string
	_datasyncProtocol                 string
	_datasyncQopConfiguration         string
	_datasyncReplicationFactor        string
	_datasyncResourceArn              string
	_datasyncS3BucketArn              string
	_datasyncS3Config                 string
	_datasyncS3StorageClass           string
	_datasyncSasConfiguration         string
	_datasyncSchedule                 string
	_datasyncSecretKey                string
	_datasyncSecurityGroupArns        []string
	_datasyncServerCertificate        string
	_datasyncServerHostname           string
	_datasyncServerPort               string
	_datasyncServerProtocol           string
	_datasyncSimpleUser               string
	_datasyncSourceLocationArn        string
	_datasyncStorageVirtualMachineArn string
	_datasyncSubdirectory             string
	_datasyncSubnetArns               []string
	_datasyncTags                     string
	_datasyncTaskArn                  string
	_datasyncTaskExecutionArn         string
	_datasyncTaskMode                 string
	_datasyncTaskReportConfig         string
	_datasyncUser                     string
	_datasyncVpcEndpointId            string
)

// Stops an DataSync task execution that's in progress. The transfer of some files
// are abruptly interrupted. File contents that're transferred to the destination
// might be incomplete or inconsistent with the source files.
//
// However, if you start a new task execution using the same task and allow it to
// finish, file content on the destination will be complete and consistent. This
// applies to other unexpected failures that interrupt a task execution. In all of
// these cases, DataSync successfully completes the transfer when you start the
// next task execution.
func datasync_CancelTaskExecution(cfg aws.Config, client *datasync.Client) {
	input := &datasync.CancelTaskExecutionInput{
		// TaskExecutionArn: *string, // Required
	}

	if len(_datasyncTaskExecutionArn) > 0 {
		input.TaskExecutionArn = aws.String(_datasyncTaskExecutionArn)
	}

	if resp, err := client.CancelTaskExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Activates an DataSync agent that you deploy in your storage environment. The
// activation process associates the agent with your Amazon Web Services account.
//
// If you haven't deployed an agent yet, see [Do I need a DataSync agent?]
//
// [Do I need a DataSync agent?]: https://docs.aws.amazon.com/datasync/latest/userguide/do-i-need-datasync-agent.html
func datasync_CreateAgent(cfg aws.Config, client *datasync.Client) {
	input := &datasync.CreateAgentInput{
		// ActivationKey: *string, // Required
	}

	if len(_datasyncActivationKey) > 0 {
		input.ActivationKey = aws.String(_datasyncActivationKey)
	}
	if len(_datasyncAgentName) > 0 {
		input.AgentName = aws.String(_datasyncAgentName)
	}
	if len(_datasyncSecurityGroupArns) > 0 {
		input.SecurityGroupArns = append([]string(nil), _datasyncSecurityGroupArns...)
	}
	if len(_datasyncSubnetArns) > 0 {
		input.SubnetArns = append([]string(nil), _datasyncSubnetArns...)
	}
	if len(_datasyncTags) > 0 {
		if err := assignInputField(input, "Tags", _datasyncTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_datasyncVpcEndpointId) > 0 {
		input.VpcEndpointId = aws.String(_datasyncVpcEndpointId)
	}

	if resp, err := client.CreateAgent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a transfer location for a Microsoft Azure Blob Storage container.
// DataSync can use this location as a transfer source or destination. You can make
// transfers with or without a [DataSync agent]that connects to your container.
//
// Before you begin, make sure you know [how DataSync accesses Azure Blob Storage] and works with [access tiers] and [blob types].
//
// [DataSync agent]: https://docs.aws.amazon.com/datasync/latest/userguide/creating-azure-blob-location.html#azure-blob-creating-agent
// [blob types]: https://docs.aws.amazon.com/datasync/latest/userguide/creating-azure-blob-location.html#blob-types
// [how DataSync accesses Azure Blob Storage]: https://docs.aws.amazon.com/datasync/latest/userguide/creating-azure-blob-location.html#azure-blob-access
// [access tiers]: https://docs.aws.amazon.com/datasync/latest/userguide/creating-azure-blob-location.html#azure-blob-access-tiers
func datasync_CreateLocationAzureBlob(cfg aws.Config, client *datasync.Client) {
	input := &datasync.CreateLocationAzureBlobInput{
		// AuthenticationType: types.AzureBlobAuthenticationType, // Required
		// ContainerUrl: *string, // Required
	}

	if len(_datasyncAuthenticationType) > 0 {
		if err := assignInputField(input, "AuthenticationType", _datasyncAuthenticationType); err != nil {
			log.Errorf("invalid --authentication-type: %s", err.Error())
			return
		}
	}
	if len(_datasyncContainerUrl) > 0 {
		input.ContainerUrl = aws.String(_datasyncContainerUrl)
	}
	if len(_datasyncAccessTier) > 0 {
		if err := assignInputField(input, "AccessTier", _datasyncAccessTier); err != nil {
			log.Errorf("invalid --access-tier: %s", err.Error())
			return
		}
	}
	if len(_datasyncAgentArns) > 0 {
		input.AgentArns = append([]string(nil), _datasyncAgentArns...)
	}
	if len(_datasyncBlobType) > 0 {
		if err := assignInputField(input, "BlobType", _datasyncBlobType); err != nil {
			log.Errorf("invalid --blob-type: %s", err.Error())
			return
		}
	}
	if len(_datasyncCmkSecretConfig) > 0 {
		if err := assignInputField(input, "CmkSecretConfig", _datasyncCmkSecretConfig); err != nil {
			log.Errorf("invalid --cmk-secret-config: %s", err.Error())
			return
		}
	}
	if len(_datasyncCustomSecretConfig) > 0 {
		if err := assignInputField(input, "CustomSecretConfig", _datasyncCustomSecretConfig); err != nil {
			log.Errorf("invalid --custom-secret-config: %s", err.Error())
			return
		}
	}
	if len(_datasyncSasConfiguration) > 0 {
		if err := assignInputField(input, "SasConfiguration", _datasyncSasConfiguration); err != nil {
			log.Errorf("invalid --sas-configuration: %s", err.Error())
			return
		}
	}
	if len(_datasyncSubdirectory) > 0 {
		input.Subdirectory = aws.String(_datasyncSubdirectory)
	}
	if len(_datasyncTags) > 0 {
		if err := assignInputField(input, "Tags", _datasyncTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLocationAzureBlob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a transfer location for an Amazon EFS file system. DataSync can use
// this location as a source or destination for transferring data.
//
// Before you begin, make sure that you understand how DataSync [accesses Amazon EFS file systems].
//
// [accesses Amazon EFS file systems]: https://docs.aws.amazon.com/datasync/latest/userguide/create-efs-location.html#create-efs-location-access
func datasync_CreateLocationEfs(cfg aws.Config, client *datasync.Client) {
	input := &datasync.CreateLocationEfsInput{
		// Ec2Config: *types.Ec2Config, // Required
		// EfsFilesystemArn: *string, // Required
	}

	if len(_datasyncEc2Config) > 0 {
		if err := assignInputField(input, "Ec2Config", _datasyncEc2Config); err != nil {
			log.Errorf("invalid --ec2-config: %s", err.Error())
			return
		}
	}
	if len(_datasyncEfsFilesystemArn) > 0 {
		input.EfsFilesystemArn = aws.String(_datasyncEfsFilesystemArn)
	}
	if len(_datasyncAccessPointArn) > 0 {
		input.AccessPointArn = aws.String(_datasyncAccessPointArn)
	}
	if len(_datasyncFileSystemAccessRoleArn) > 0 {
		input.FileSystemAccessRoleArn = aws.String(_datasyncFileSystemAccessRoleArn)
	}
	if len(_datasyncInTransitEncryption) > 0 {
		if err := assignInputField(input, "InTransitEncryption", _datasyncInTransitEncryption); err != nil {
			log.Errorf("invalid --in-transit-encryption: %s", err.Error())
			return
		}
	}
	if len(_datasyncSubdirectory) > 0 {
		input.Subdirectory = aws.String(_datasyncSubdirectory)
	}
	if len(_datasyncTags) > 0 {
		if err := assignInputField(input, "Tags", _datasyncTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLocationEfs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a transfer location for an Amazon FSx for Lustre file system. DataSync
// can use this location as a source or destination for transferring data.
//
// Before you begin, make sure that you understand how DataSync [accesses FSx for Lustre file systems].
//
// [accesses FSx for Lustre file systems]: https://docs.aws.amazon.com/datasync/latest/userguide/create-lustre-location.html#create-lustre-location-access
func datasync_CreateLocationFsxLustre(cfg aws.Config, client *datasync.Client) {
	input := &datasync.CreateLocationFsxLustreInput{
		// FsxFilesystemArn: *string, // Required
		// SecurityGroupArns: []string, // Required
	}

	if len(_datasyncFsxFilesystemArn) > 0 {
		input.FsxFilesystemArn = aws.String(_datasyncFsxFilesystemArn)
	}
	if len(_datasyncSecurityGroupArns) > 0 {
		input.SecurityGroupArns = append([]string(nil), _datasyncSecurityGroupArns...)
	}
	if len(_datasyncSubdirectory) > 0 {
		input.Subdirectory = aws.String(_datasyncSubdirectory)
	}
	if len(_datasyncTags) > 0 {
		if err := assignInputField(input, "Tags", _datasyncTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLocationFsxLustre(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a transfer location for an Amazon FSx for NetApp ONTAP file system.
// DataSync can use this location as a source or destination for transferring data.
//
// Before you begin, make sure that you understand how DataSync [accesses FSx for ONTAP file systems].
//
// [accesses FSx for ONTAP file systems]: https://docs.aws.amazon.com/datasync/latest/userguide/create-ontap-location.html#create-ontap-location-access
func datasync_CreateLocationFsxOntap(cfg aws.Config, client *datasync.Client) {
	input := &datasync.CreateLocationFsxOntapInput{
		// Protocol: *types.FsxProtocol, // Required
		// SecurityGroupArns: []string, // Required
		// StorageVirtualMachineArn: *string, // Required
	}

	if len(_datasyncProtocol) > 0 {
		if err := assignInputField(input, "Protocol", _datasyncProtocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}
	if len(_datasyncSecurityGroupArns) > 0 {
		input.SecurityGroupArns = append([]string(nil), _datasyncSecurityGroupArns...)
	}
	if len(_datasyncStorageVirtualMachineArn) > 0 {
		input.StorageVirtualMachineArn = aws.String(_datasyncStorageVirtualMachineArn)
	}
	if len(_datasyncSubdirectory) > 0 {
		input.Subdirectory = aws.String(_datasyncSubdirectory)
	}
	if len(_datasyncTags) > 0 {
		if err := assignInputField(input, "Tags", _datasyncTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLocationFsxOntap(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a transfer location for an Amazon FSx for OpenZFS file system. DataSync
// can use this location as a source or destination for transferring data.
//
// Before you begin, make sure that you understand how DataSync [accesses FSx for OpenZFS file systems].
//
// Request parameters related to SMB aren't supported with the
// CreateLocationFsxOpenZfs operation.
//
// [accesses FSx for OpenZFS file systems]: https://docs.aws.amazon.com/datasync/latest/userguide/create-openzfs-location.html#create-openzfs-access
func datasync_CreateLocationFsxOpenZfs(cfg aws.Config, client *datasync.Client) {
	input := &datasync.CreateLocationFsxOpenZfsInput{
		// FsxFilesystemArn: *string, // Required
		// Protocol: *types.FsxProtocol, // Required
		// SecurityGroupArns: []string, // Required
	}

	if len(_datasyncFsxFilesystemArn) > 0 {
		input.FsxFilesystemArn = aws.String(_datasyncFsxFilesystemArn)
	}
	if len(_datasyncProtocol) > 0 {
		if err := assignInputField(input, "Protocol", _datasyncProtocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}
	if len(_datasyncSecurityGroupArns) > 0 {
		input.SecurityGroupArns = append([]string(nil), _datasyncSecurityGroupArns...)
	}
	if len(_datasyncSubdirectory) > 0 {
		input.Subdirectory = aws.String(_datasyncSubdirectory)
	}
	if len(_datasyncTags) > 0 {
		if err := assignInputField(input, "Tags", _datasyncTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLocationFsxOpenZfs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a transfer location for an Amazon FSx for Windows File Server file
// system. DataSync can use this location as a source or destination for
// transferring data.
//
// Before you begin, make sure that you understand how DataSync [accesses FSx for Windows File Server file systems].
//
// [accesses FSx for Windows File Server file systems]: https://docs.aws.amazon.com/datasync/latest/userguide/create-fsx-location.html#create-fsx-location-access
func datasync_CreateLocationFsxWindows(cfg aws.Config, client *datasync.Client) {
	input := &datasync.CreateLocationFsxWindowsInput{
		// FsxFilesystemArn: *string, // Required
		// Password: *string, // Required
		// SecurityGroupArns: []string, // Required
		// User: *string, // Required
	}

	if len(_datasyncFsxFilesystemArn) > 0 {
		input.FsxFilesystemArn = aws.String(_datasyncFsxFilesystemArn)
	}
	if len(_datasyncPassword) > 0 {
		input.Password = aws.String(_datasyncPassword)
	}
	if len(_datasyncSecurityGroupArns) > 0 {
		input.SecurityGroupArns = append([]string(nil), _datasyncSecurityGroupArns...)
	}
	if len(_datasyncUser) > 0 {
		input.User = aws.String(_datasyncUser)
	}
	if len(_datasyncDomain) > 0 {
		input.Domain = aws.String(_datasyncDomain)
	}
	if len(_datasyncSubdirectory) > 0 {
		input.Subdirectory = aws.String(_datasyncSubdirectory)
	}
	if len(_datasyncTags) > 0 {
		if err := assignInputField(input, "Tags", _datasyncTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLocationFsxWindows(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a transfer location for a Hadoop Distributed File System (HDFS).
// DataSync can use this location as a source or destination for transferring data.
//
// Before you begin, make sure that you understand how DataSync [accesses HDFS clusters].
//
// [accesses HDFS clusters]: https://docs.aws.amazon.com/datasync/latest/userguide/create-hdfs-location.html#accessing-hdfs
func datasync_CreateLocationHdfs(cfg aws.Config, client *datasync.Client) {
	input := &datasync.CreateLocationHdfsInput{
		// AgentArns: []string, // Required
		// AuthenticationType: types.HdfsAuthenticationType, // Required
		// NameNodes: []types.HdfsNameNode, // Required
	}

	if len(_datasyncAgentArns) > 0 {
		input.AgentArns = append([]string(nil), _datasyncAgentArns...)
	}
	if len(_datasyncAuthenticationType) > 0 {
		if err := assignInputField(input, "AuthenticationType", _datasyncAuthenticationType); err != nil {
			log.Errorf("invalid --authentication-type: %s", err.Error())
			return
		}
	}
	if len(_datasyncNameNodes) > 0 {
		if err := assignInputField(input, "NameNodes", _datasyncNameNodes); err != nil {
			log.Errorf("invalid --name-nodes: %s", err.Error())
			return
		}
	}
	if len(_datasyncBlockSize) > 0 {
		if err := assignInputField(input, "BlockSize", _datasyncBlockSize); err != nil {
			log.Errorf("invalid --block-size: %s", err.Error())
			return
		}
	}
	if len(_datasyncKerberosKeytab) > 0 {
		if err := assignInputField(input, "KerberosKeytab", _datasyncKerberosKeytab); err != nil {
			log.Errorf("invalid --kerberos-keytab: %s", err.Error())
			return
		}
	}
	if len(_datasyncKerberosKrb5Conf) > 0 {
		if err := assignInputField(input, "KerberosKrb5Conf", _datasyncKerberosKrb5Conf); err != nil {
			log.Errorf("invalid --kerberos-krb5-conf: %s", err.Error())
			return
		}
	}
	if len(_datasyncKerberosPrincipal) > 0 {
		input.KerberosPrincipal = aws.String(_datasyncKerberosPrincipal)
	}
	if len(_datasyncKmsKeyProviderUri) > 0 {
		input.KmsKeyProviderUri = aws.String(_datasyncKmsKeyProviderUri)
	}
	if len(_datasyncQopConfiguration) > 0 {
		if err := assignInputField(input, "QopConfiguration", _datasyncQopConfiguration); err != nil {
			log.Errorf("invalid --qop-configuration: %s", err.Error())
			return
		}
	}
	if len(_datasyncReplicationFactor) > 0 {
		if err := assignInputField(input, "ReplicationFactor", _datasyncReplicationFactor); err != nil {
			log.Errorf("invalid --replication-factor: %s", err.Error())
			return
		}
	}
	if len(_datasyncSimpleUser) > 0 {
		input.SimpleUser = aws.String(_datasyncSimpleUser)
	}
	if len(_datasyncSubdirectory) > 0 {
		input.Subdirectory = aws.String(_datasyncSubdirectory)
	}
	if len(_datasyncTags) > 0 {
		if err := assignInputField(input, "Tags", _datasyncTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLocationHdfs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a transfer location for a Network File System (NFS) file server.
// DataSync can use this location as a source or destination for transferring data.
//
// Before you begin, make sure that you understand how DataSync [accesses NFS file servers].
//
// [accesses NFS file servers]: https://docs.aws.amazon.com/datasync/latest/userguide/create-nfs-location.html#accessing-nfs
func datasync_CreateLocationNfs(cfg aws.Config, client *datasync.Client) {
	input := &datasync.CreateLocationNfsInput{
		// OnPremConfig: *types.OnPremConfig, // Required
		// ServerHostname: *string, // Required
		// Subdirectory: *string, // Required
	}

	if len(_datasyncOnPremConfig) > 0 {
		if err := assignInputField(input, "OnPremConfig", _datasyncOnPremConfig); err != nil {
			log.Errorf("invalid --on-prem-config: %s", err.Error())
			return
		}
	}
	if len(_datasyncServerHostname) > 0 {
		input.ServerHostname = aws.String(_datasyncServerHostname)
	}
	if len(_datasyncSubdirectory) > 0 {
		input.Subdirectory = aws.String(_datasyncSubdirectory)
	}
	if len(_datasyncMountOptions) > 0 {
		if err := assignInputField(input, "MountOptions", _datasyncMountOptions); err != nil {
			log.Errorf("invalid --mount-options: %s", err.Error())
			return
		}
	}
	if len(_datasyncTags) > 0 {
		if err := assignInputField(input, "Tags", _datasyncTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLocationNfs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a transfer location for an object storage system. DataSync can use this
// location as a source or destination for transferring data. You can make
// transfers with or without a [DataSync agent].
//
// Before you begin, make sure that you understand the [prerequisites] for DataSync to work with
// object storage systems.
//
// [DataSync agent]: https://docs.aws.amazon.com/datasync/latest/userguide/do-i-need-datasync-agent.html#when-agent-required
// [prerequisites]: https://docs.aws.amazon.com/datasync/latest/userguide/create-object-location.html#create-object-location-prerequisites
func datasync_CreateLocationObjectStorage(cfg aws.Config, client *datasync.Client) {
	input := &datasync.CreateLocationObjectStorageInput{
		// BucketName: *string, // Required
		// ServerHostname: *string, // Required
	}

	if len(_datasyncBucketName) > 0 {
		input.BucketName = aws.String(_datasyncBucketName)
	}
	if len(_datasyncServerHostname) > 0 {
		input.ServerHostname = aws.String(_datasyncServerHostname)
	}
	if len(_datasyncAccessKey) > 0 {
		input.AccessKey = aws.String(_datasyncAccessKey)
	}
	if len(_datasyncAgentArns) > 0 {
		input.AgentArns = append([]string(nil), _datasyncAgentArns...)
	}
	if len(_datasyncCmkSecretConfig) > 0 {
		if err := assignInputField(input, "CmkSecretConfig", _datasyncCmkSecretConfig); err != nil {
			log.Errorf("invalid --cmk-secret-config: %s", err.Error())
			return
		}
	}
	if len(_datasyncCustomSecretConfig) > 0 {
		if err := assignInputField(input, "CustomSecretConfig", _datasyncCustomSecretConfig); err != nil {
			log.Errorf("invalid --custom-secret-config: %s", err.Error())
			return
		}
	}
	if len(_datasyncSecretKey) > 0 {
		input.SecretKey = aws.String(_datasyncSecretKey)
	}
	if len(_datasyncServerCertificate) > 0 {
		if err := assignInputField(input, "ServerCertificate", _datasyncServerCertificate); err != nil {
			log.Errorf("invalid --server-certificate: %s", err.Error())
			return
		}
	}
	if len(_datasyncServerPort) > 0 {
		if err := assignInputField(input, "ServerPort", _datasyncServerPort); err != nil {
			log.Errorf("invalid --server-port: %s", err.Error())
			return
		}
	}
	if len(_datasyncServerProtocol) > 0 {
		if err := assignInputField(input, "ServerProtocol", _datasyncServerProtocol); err != nil {
			log.Errorf("invalid --server-protocol: %s", err.Error())
			return
		}
	}
	if len(_datasyncSubdirectory) > 0 {
		input.Subdirectory = aws.String(_datasyncSubdirectory)
	}
	if len(_datasyncTags) > 0 {
		if err := assignInputField(input, "Tags", _datasyncTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLocationObjectStorage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a transfer location for an Amazon S3 bucket. DataSync can use this
// location as a source or destination for transferring data.
//
// Before you begin, make sure that you read the following topics:
//
// [Storage class considerations with Amazon S3 locations]
//
// [Evaluating S3 request costs when using DataSync]
//
// For more information, see [Configuring transfers with Amazon S3].
//
// [Storage class considerations with Amazon S3 locations]: https://docs.aws.amazon.com/datasync/latest/userguide/create-s3-location.html#using-storage-classes
// [Configuring transfers with Amazon S3]: https://docs.aws.amazon.com/datasync/latest/userguide/create-s3-location.html
// [Evaluating S3 request costs when using DataSync]: https://docs.aws.amazon.com/datasync/latest/userguide/create-s3-location.html#create-s3-location-s3-requests
func datasync_CreateLocationS3(cfg aws.Config, client *datasync.Client) {
	input := &datasync.CreateLocationS3Input{
		// S3BucketArn: *string, // Required
		// S3Config: *types.S3Config, // Required
	}

	if len(_datasyncS3BucketArn) > 0 {
		input.S3BucketArn = aws.String(_datasyncS3BucketArn)
	}
	if len(_datasyncS3Config) > 0 {
		if err := assignInputField(input, "S3Config", _datasyncS3Config); err != nil {
			log.Errorf("invalid --s3-config: %s", err.Error())
			return
		}
	}
	if len(_datasyncAgentArns) > 0 {
		input.AgentArns = append([]string(nil), _datasyncAgentArns...)
	}
	if len(_datasyncS3StorageClass) > 0 {
		if err := assignInputField(input, "S3StorageClass", _datasyncS3StorageClass); err != nil {
			log.Errorf("invalid --s3-storage-class: %s", err.Error())
			return
		}
	}
	if len(_datasyncSubdirectory) > 0 {
		input.Subdirectory = aws.String(_datasyncSubdirectory)
	}
	if len(_datasyncTags) > 0 {
		if err := assignInputField(input, "Tags", _datasyncTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLocationS3(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a transfer location for a Server Message Block (SMB) file server.
// DataSync can use this location as a source or destination for transferring data.
//
// Before you begin, make sure that you understand how DataSync accesses SMB file
// servers. For more information, see [Providing DataSync access to SMB file servers].
//
// [Providing DataSync access to SMB file servers]: https://docs.aws.amazon.com/datasync/latest/userguide/create-smb-location.html#configuring-smb-permissions
func datasync_CreateLocationSmb(cfg aws.Config, client *datasync.Client) {
	input := &datasync.CreateLocationSmbInput{
		// AgentArns: []string, // Required
		// ServerHostname: *string, // Required
		// Subdirectory: *string, // Required
	}

	if len(_datasyncAgentArns) > 0 {
		input.AgentArns = append([]string(nil), _datasyncAgentArns...)
	}
	if len(_datasyncServerHostname) > 0 {
		input.ServerHostname = aws.String(_datasyncServerHostname)
	}
	if len(_datasyncSubdirectory) > 0 {
		input.Subdirectory = aws.String(_datasyncSubdirectory)
	}
	if len(_datasyncAuthenticationType) > 0 {
		if err := assignInputField(input, "AuthenticationType", _datasyncAuthenticationType); err != nil {
			log.Errorf("invalid --authentication-type: %s", err.Error())
			return
		}
	}
	if len(_datasyncCmkSecretConfig) > 0 {
		if err := assignInputField(input, "CmkSecretConfig", _datasyncCmkSecretConfig); err != nil {
			log.Errorf("invalid --cmk-secret-config: %s", err.Error())
			return
		}
	}
	if len(_datasyncCustomSecretConfig) > 0 {
		if err := assignInputField(input, "CustomSecretConfig", _datasyncCustomSecretConfig); err != nil {
			log.Errorf("invalid --custom-secret-config: %s", err.Error())
			return
		}
	}
	if len(_datasyncDnsIpAddresses) > 0 {
		input.DnsIpAddresses = append([]string(nil), _datasyncDnsIpAddresses...)
	}
	if len(_datasyncDomain) > 0 {
		input.Domain = aws.String(_datasyncDomain)
	}
	if len(_datasyncKerberosKeytab) > 0 {
		if err := assignInputField(input, "KerberosKeytab", _datasyncKerberosKeytab); err != nil {
			log.Errorf("invalid --kerberos-keytab: %s", err.Error())
			return
		}
	}
	if len(_datasyncKerberosKrb5Conf) > 0 {
		if err := assignInputField(input, "KerberosKrb5Conf", _datasyncKerberosKrb5Conf); err != nil {
			log.Errorf("invalid --kerberos-krb5-conf: %s", err.Error())
			return
		}
	}
	if len(_datasyncKerberosPrincipal) > 0 {
		input.KerberosPrincipal = aws.String(_datasyncKerberosPrincipal)
	}
	if len(_datasyncMountOptions) > 0 {
		if err := assignInputField(input, "MountOptions", _datasyncMountOptions); err != nil {
			log.Errorf("invalid --mount-options: %s", err.Error())
			return
		}
	}
	if len(_datasyncPassword) > 0 {
		input.Password = aws.String(_datasyncPassword)
	}
	if len(_datasyncTags) > 0 {
		if err := assignInputField(input, "Tags", _datasyncTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_datasyncUser) > 0 {
		input.User = aws.String(_datasyncUser)
	}

	if resp, err := client.CreateLocationSmb(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configures a task, which defines where and how DataSync transfers your data.
// A task includes a source location, destination location, and transfer options
// (such as bandwidth limits, scheduling, and more).
//
// If you're planning to transfer data to or from an Amazon S3 location, review [how DataSync can affect your S3 request charges]
// and the [DataSync pricing page]before you begin.
//
// [how DataSync can affect your S3 request charges]: https://docs.aws.amazon.com/datasync/latest/userguide/create-s3-location.html#create-s3-location-s3-requests
// [DataSync pricing page]: http://aws.amazon.com/datasync/pricing/
func datasync_CreateTask(cfg aws.Config, client *datasync.Client) {
	input := &datasync.CreateTaskInput{
		// DestinationLocationArn: *string, // Required
		// SourceLocationArn: *string, // Required
	}

	if len(_datasyncDestinationLocationArn) > 0 {
		input.DestinationLocationArn = aws.String(_datasyncDestinationLocationArn)
	}
	if len(_datasyncSourceLocationArn) > 0 {
		input.SourceLocationArn = aws.String(_datasyncSourceLocationArn)
	}
	if len(_datasyncCloudWatchLogGroupArn) > 0 {
		input.CloudWatchLogGroupArn = aws.String(_datasyncCloudWatchLogGroupArn)
	}
	if len(_datasyncExcludes) > 0 {
		if err := assignInputField(input, "Excludes", _datasyncExcludes); err != nil {
			log.Errorf("invalid --excludes: %s", err.Error())
			return
		}
	}
	if len(_datasyncIncludes) > 0 {
		if err := assignInputField(input, "Includes", _datasyncIncludes); err != nil {
			log.Errorf("invalid --includes: %s", err.Error())
			return
		}
	}
	if len(_datasyncManifestConfig) > 0 {
		if err := assignInputField(input, "ManifestConfig", _datasyncManifestConfig); err != nil {
			log.Errorf("invalid --manifest-config: %s", err.Error())
			return
		}
	}
	if len(_datasyncName) > 0 {
		input.Name = aws.String(_datasyncName)
	}
	if len(_datasyncOptions) > 0 {
		if err := assignInputField(input, "Options", _datasyncOptions); err != nil {
			log.Errorf("invalid --options: %s", err.Error())
			return
		}
	}
	if len(_datasyncSchedule) > 0 {
		if err := assignInputField(input, "Schedule", _datasyncSchedule); err != nil {
			log.Errorf("invalid --schedule: %s", err.Error())
			return
		}
	}
	if len(_datasyncTags) > 0 {
		if err := assignInputField(input, "Tags", _datasyncTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_datasyncTaskMode) > 0 {
		if err := assignInputField(input, "TaskMode", _datasyncTaskMode); err != nil {
			log.Errorf("invalid --task-mode: %s", err.Error())
			return
		}
	}
	if len(_datasyncTaskReportConfig) > 0 {
		if err := assignInputField(input, "TaskReportConfig", _datasyncTaskReportConfig); err != nil {
			log.Errorf("invalid --task-report-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes an DataSync agent resource from your Amazon Web Services account.
// Keep in mind that this operation (which can't be undone) doesn't remove the
// agent's virtual machine (VM) or Amazon EC2 instance from your storage
// environment. For next steps, you can delete the VM or instance from your storage
// environment or reuse it to [activate a new agent].
//
// [activate a new agent]: https://docs.aws.amazon.com/datasync/latest/userguide/activate-agent.html
func datasync_DeleteAgent(cfg aws.Config, client *datasync.Client) {
	input := &datasync.DeleteAgentInput{
		// AgentArn: *string, // Required
	}

	if len(_datasyncAgentArn) > 0 {
		input.AgentArn = aws.String(_datasyncAgentArn)
	}

	if resp, err := client.DeleteAgent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a transfer location resource from DataSync.
func datasync_DeleteLocation(cfg aws.Config, client *datasync.Client) {
	input := &datasync.DeleteLocationInput{
		// LocationArn: *string, // Required
	}

	if len(_datasyncLocationArn) > 0 {
		input.LocationArn = aws.String(_datasyncLocationArn)
	}

	if resp, err := client.DeleteLocation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a transfer task resource from DataSync.
func datasync_DeleteTask(cfg aws.Config, client *datasync.Client) {
	input := &datasync.DeleteTaskInput{
		// TaskArn: *string, // Required
	}

	if len(_datasyncTaskArn) > 0 {
		input.TaskArn = aws.String(_datasyncTaskArn)
	}

	if resp, err := client.DeleteTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an DataSync agent, such as its name, service endpoint
// type, and status.
func datasync_DescribeAgent(cfg aws.Config, client *datasync.Client) {
	input := &datasync.DescribeAgentInput{
		// AgentArn: *string, // Required
	}

	if len(_datasyncAgentArn) > 0 {
		input.AgentArn = aws.String(_datasyncAgentArn)
	}

	if resp, err := client.DescribeAgent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides details about how an DataSync transfer location for Microsoft Azure
// Blob Storage is configured.
func datasync_DescribeLocationAzureBlob(cfg aws.Config, client *datasync.Client) {
	input := &datasync.DescribeLocationAzureBlobInput{
		// LocationArn: *string, // Required
	}

	if len(_datasyncLocationArn) > 0 {
		input.LocationArn = aws.String(_datasyncLocationArn)
	}

	if resp, err := client.DescribeLocationAzureBlob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides details about how an DataSync transfer location for an Amazon EFS file
// system is configured.
func datasync_DescribeLocationEfs(cfg aws.Config, client *datasync.Client) {
	input := &datasync.DescribeLocationEfsInput{
		// LocationArn: *string, // Required
	}

	if len(_datasyncLocationArn) > 0 {
		input.LocationArn = aws.String(_datasyncLocationArn)
	}

	if resp, err := client.DescribeLocationEfs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides details about how an DataSync transfer location for an Amazon FSx for
// Lustre file system is configured.
func datasync_DescribeLocationFsxLustre(cfg aws.Config, client *datasync.Client) {
	input := &datasync.DescribeLocationFsxLustreInput{
		// LocationArn: *string, // Required
	}

	if len(_datasyncLocationArn) > 0 {
		input.LocationArn = aws.String(_datasyncLocationArn)
	}

	if resp, err := client.DescribeLocationFsxLustre(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides details about how an DataSync transfer location for an Amazon FSx for
// NetApp ONTAP file system is configured.
//
// If your location uses SMB, the DescribeLocationFsxOntap operation doesn't
// actually return a Password .
func datasync_DescribeLocationFsxOntap(cfg aws.Config, client *datasync.Client) {
	input := &datasync.DescribeLocationFsxOntapInput{
		// LocationArn: *string, // Required
	}

	if len(_datasyncLocationArn) > 0 {
		input.LocationArn = aws.String(_datasyncLocationArn)
	}

	if resp, err := client.DescribeLocationFsxOntap(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides details about how an DataSync transfer location for an Amazon FSx for
// OpenZFS file system is configured.
//
// Response elements related to SMB aren't supported with the
// DescribeLocationFsxOpenZfs operation.
func datasync_DescribeLocationFsxOpenZfs(cfg aws.Config, client *datasync.Client) {
	input := &datasync.DescribeLocationFsxOpenZfsInput{
		// LocationArn: *string, // Required
	}

	if len(_datasyncLocationArn) > 0 {
		input.LocationArn = aws.String(_datasyncLocationArn)
	}

	if resp, err := client.DescribeLocationFsxOpenZfs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides details about how an DataSync transfer location for an Amazon FSx for
// Windows File Server file system is configured.
func datasync_DescribeLocationFsxWindows(cfg aws.Config, client *datasync.Client) {
	input := &datasync.DescribeLocationFsxWindowsInput{
		// LocationArn: *string, // Required
	}

	if len(_datasyncLocationArn) > 0 {
		input.LocationArn = aws.String(_datasyncLocationArn)
	}

	if resp, err := client.DescribeLocationFsxWindows(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides details about how an DataSync transfer location for a Hadoop
// Distributed File System (HDFS) is configured.
func datasync_DescribeLocationHdfs(cfg aws.Config, client *datasync.Client) {
	input := &datasync.DescribeLocationHdfsInput{
		// LocationArn: *string, // Required
	}

	if len(_datasyncLocationArn) > 0 {
		input.LocationArn = aws.String(_datasyncLocationArn)
	}

	if resp, err := client.DescribeLocationHdfs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides details about how an DataSync transfer location for a Network File
// System (NFS) file server is configured.
func datasync_DescribeLocationNfs(cfg aws.Config, client *datasync.Client) {
	input := &datasync.DescribeLocationNfsInput{
		// LocationArn: *string, // Required
	}

	if len(_datasyncLocationArn) > 0 {
		input.LocationArn = aws.String(_datasyncLocationArn)
	}

	if resp, err := client.DescribeLocationNfs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides details about how an DataSync transfer location for an object storage
// system is configured.
func datasync_DescribeLocationObjectStorage(cfg aws.Config, client *datasync.Client) {
	input := &datasync.DescribeLocationObjectStorageInput{
		// LocationArn: *string, // Required
	}

	if len(_datasyncLocationArn) > 0 {
		input.LocationArn = aws.String(_datasyncLocationArn)
	}

	if resp, err := client.DescribeLocationObjectStorage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides details about how an DataSync transfer location for an S3 bucket is
// configured.
func datasync_DescribeLocationS3(cfg aws.Config, client *datasync.Client) {
	input := &datasync.DescribeLocationS3Input{
		// LocationArn: *string, // Required
	}

	if len(_datasyncLocationArn) > 0 {
		input.LocationArn = aws.String(_datasyncLocationArn)
	}

	if resp, err := client.DescribeLocationS3(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides details about how an DataSync transfer location for a Server Message
// Block (SMB) file server is configured.
func datasync_DescribeLocationSmb(cfg aws.Config, client *datasync.Client) {
	input := &datasync.DescribeLocationSmbInput{
		// LocationArn: *string, // Required
	}

	if len(_datasyncLocationArn) > 0 {
		input.LocationArn = aws.String(_datasyncLocationArn)
	}

	if resp, err := client.DescribeLocationSmb(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about a task, which defines where and how DataSync
// transfers your data.
func datasync_DescribeTask(cfg aws.Config, client *datasync.Client) {
	input := &datasync.DescribeTaskInput{
		// TaskArn: *string, // Required
	}

	if len(_datasyncTaskArn) > 0 {
		input.TaskArn = aws.String(_datasyncTaskArn)
	}

	if resp, err := client.DescribeTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about an execution of your DataSync task. You can use this
// operation to help monitor the progress of an ongoing data transfer or check the
// results of the transfer.
//
// Some DescribeTaskExecution response elements are only relevant to a specific
// task mode. For information, see [Understanding task mode differences]and [Understanding data transfer performance counters].
//
// [Understanding task mode differences]: https://docs.aws.amazon.com/datasync/latest/userguide/choosing-task-mode.html#task-mode-differences
// [Understanding data transfer performance counters]: https://docs.aws.amazon.com/datasync/latest/userguide/transfer-performance-counters.html
func datasync_DescribeTaskExecution(cfg aws.Config, client *datasync.Client) {
	input := &datasync.DescribeTaskExecutionInput{
		// TaskExecutionArn: *string, // Required
	}

	if len(_datasyncTaskExecutionArn) > 0 {
		input.TaskExecutionArn = aws.String(_datasyncTaskExecutionArn)
	}

	if resp, err := client.DescribeTaskExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of DataSync agents that belong to an Amazon Web Services account
// in the Amazon Web Services Region specified in the request.
//
// With pagination, you can reduce the number of agents returned in a response. If
// you get a truncated list of agents in a response, the response contains a marker
// that you can specify in your next request to fetch the next page of agents.
//
// ListAgents is eventually consistent. This means the result of running the
// operation might not reflect that you just created or deleted an agent. For
// example, if you create an agent with [CreateAgent]and then immediately run ListAgents , that
// agent might not show up in the list right away. In situations like this, you can
// always confirm whether an agent has been created (or deleted) by using [DescribeAgent].
//
// [DescribeAgent]: https://docs.aws.amazon.com/datasync/latest/userguide/API_DescribeAgent.html
// [CreateAgent]: https://docs.aws.amazon.com/datasync/latest/userguide/API_CreateAgent.html
func datasync_ListAgents(cfg aws.Config, client *datasync.Client) {
	input := &datasync.ListAgentsInput{}

	if len(_datasyncMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datasyncMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datasyncNextToken) > 0 {
		input.NextToken = aws.String(_datasyncNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAgents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datasync.ListAgentsOutput
	p := datasync.NewListAgentsPaginator(client, input)
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

// Returns a list of source and destination locations.
// If you have more locations than are returned in a response (that is, the
// response returns only a truncated list of your agents), the response contains a
// token that you can specify in your next request to fetch the next page of
// locations.
func datasync_ListLocations(cfg aws.Config, client *datasync.Client) {
	input := &datasync.ListLocationsInput{}

	if len(_datasyncFilters) > 0 {
		if err := assignInputField(input, "Filters", _datasyncFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_datasyncMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datasyncMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datasyncNextToken) > 0 {
		input.NextToken = aws.String(_datasyncNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLocations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datasync.ListLocationsOutput
	p := datasync.NewListLocationsPaginator(client, input)
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

// Returns all the tags associated with an Amazon Web Services resource.
func datasync_ListTagsForResource(cfg aws.Config, client *datasync.Client) {
	input := &datasync.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_datasyncResourceArn) > 0 {
		input.ResourceArn = aws.String(_datasyncResourceArn)
	}
	if len(_datasyncMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datasyncMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datasyncNextToken) > 0 {
		input.NextToken = aws.String(_datasyncNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datasync.ListTagsForResourceOutput
	p := datasync.NewListTagsForResourcePaginator(client, input)
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

// Returns a list of executions for an DataSync transfer task.
func datasync_ListTaskExecutions(cfg aws.Config, client *datasync.Client) {
	input := &datasync.ListTaskExecutionsInput{}

	if len(_datasyncMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datasyncMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datasyncNextToken) > 0 {
		input.NextToken = aws.String(_datasyncNextToken)
	}
	if len(_datasyncTaskArn) > 0 {
		input.TaskArn = aws.String(_datasyncTaskArn)
	}

	if disablePaginator() {
		if resp, err := client.ListTaskExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datasync.ListTaskExecutionsOutput
	p := datasync.NewListTaskExecutionsPaginator(client, input)
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

// Returns a list of the DataSync tasks you created.
func datasync_ListTasks(cfg aws.Config, client *datasync.Client) {
	input := &datasync.ListTasksInput{}

	if len(_datasyncFilters) > 0 {
		if err := assignInputField(input, "Filters", _datasyncFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_datasyncMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datasyncMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datasyncNextToken) > 0 {
		input.NextToken = aws.String(_datasyncNextToken)
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

	var results []*datasync.ListTasksOutput
	p := datasync.NewListTasksPaginator(client, input)
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

// Starts an DataSync transfer task. For each task, you can only run one task
// execution at a time.
//
// There are several steps to a task execution. For more information, see [Task execution statuses].
//
// If you're planning to transfer data to or from an Amazon S3 location, review [how DataSync can affect your S3 request charges]
// and the [DataSync pricing page]before you begin.
//
// [Task execution statuses]: https://docs.aws.amazon.com/datasync/latest/userguide/working-with-task-executions.html#understand-task-execution-statuses
// [how DataSync can affect your S3 request charges]: https://docs.aws.amazon.com/datasync/latest/userguide/create-s3-location.html#create-s3-location-s3-requests
// [DataSync pricing page]: http://aws.amazon.com/datasync/pricing/
func datasync_StartTaskExecution(cfg aws.Config, client *datasync.Client) {
	input := &datasync.StartTaskExecutionInput{
		// TaskArn: *string, // Required
	}

	if len(_datasyncTaskArn) > 0 {
		input.TaskArn = aws.String(_datasyncTaskArn)
	}
	if len(_datasyncExcludes) > 0 {
		if err := assignInputField(input, "Excludes", _datasyncExcludes); err != nil {
			log.Errorf("invalid --excludes: %s", err.Error())
			return
		}
	}
	if len(_datasyncIncludes) > 0 {
		if err := assignInputField(input, "Includes", _datasyncIncludes); err != nil {
			log.Errorf("invalid --includes: %s", err.Error())
			return
		}
	}
	if len(_datasyncManifestConfig) > 0 {
		if err := assignInputField(input, "ManifestConfig", _datasyncManifestConfig); err != nil {
			log.Errorf("invalid --manifest-config: %s", err.Error())
			return
		}
	}
	if len(_datasyncOverrideOptions) > 0 {
		if err := assignInputField(input, "OverrideOptions", _datasyncOverrideOptions); err != nil {
			log.Errorf("invalid --override-options: %s", err.Error())
			return
		}
	}
	if len(_datasyncTags) > 0 {
		if err := assignInputField(input, "Tags", _datasyncTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_datasyncTaskReportConfig) > 0 {
		if err := assignInputField(input, "TaskReportConfig", _datasyncTaskReportConfig); err != nil {
			log.Errorf("invalid --task-report-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartTaskExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies a tag to an Amazon Web Services resource. Tags are key-value pairs that
// can help you manage, filter, and search for your resources.
//
// These include DataSync resources, such as locations, tasks, and task executions.
func datasync_TagResource(cfg aws.Config, client *datasync.Client) {
	input := &datasync.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.TagListEntry, // Required
	}

	if len(_datasyncResourceArn) > 0 {
		input.ResourceArn = aws.String(_datasyncResourceArn)
	}
	if len(_datasyncTags) > 0 {
		if err := assignInputField(input, "Tags", _datasyncTags); err != nil {
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

// Removes tags from an Amazon Web Services resource.
func datasync_UntagResource(cfg aws.Config, client *datasync.Client) {
	input := &datasync.UntagResourceInput{
		// Keys: []string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_datasyncKeys) > 0 {
		input.Keys = append([]string(nil), _datasyncKeys...)
	}
	if len(_datasyncResourceArn) > 0 {
		input.ResourceArn = aws.String(_datasyncResourceArn)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the name of an DataSync agent.
func datasync_UpdateAgent(cfg aws.Config, client *datasync.Client) {
	input := &datasync.UpdateAgentInput{
		// AgentArn: *string, // Required
	}

	if len(_datasyncAgentArn) > 0 {
		input.AgentArn = aws.String(_datasyncAgentArn)
	}
	if len(_datasyncName) > 0 {
		input.Name = aws.String(_datasyncName)
	}

	if resp, err := client.UpdateAgent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the following configurations of the Microsoft Azure Blob Storage
// transfer location that you're using with DataSync.
//
// For more information, see [Configuring DataSync transfers with Azure Blob Storage].
//
// [Configuring DataSync transfers with Azure Blob Storage]: https://docs.aws.amazon.com/datasync/latest/userguide/creating-azure-blob-location.html
func datasync_UpdateLocationAzureBlob(cfg aws.Config, client *datasync.Client) {
	input := &datasync.UpdateLocationAzureBlobInput{
		// LocationArn: *string, // Required
	}

	if len(_datasyncLocationArn) > 0 {
		input.LocationArn = aws.String(_datasyncLocationArn)
	}
	if len(_datasyncAccessTier) > 0 {
		if err := assignInputField(input, "AccessTier", _datasyncAccessTier); err != nil {
			log.Errorf("invalid --access-tier: %s", err.Error())
			return
		}
	}
	if len(_datasyncAgentArns) > 0 {
		input.AgentArns = append([]string(nil), _datasyncAgentArns...)
	}
	if len(_datasyncAuthenticationType) > 0 {
		if err := assignInputField(input, "AuthenticationType", _datasyncAuthenticationType); err != nil {
			log.Errorf("invalid --authentication-type: %s", err.Error())
			return
		}
	}
	if len(_datasyncBlobType) > 0 {
		if err := assignInputField(input, "BlobType", _datasyncBlobType); err != nil {
			log.Errorf("invalid --blob-type: %s", err.Error())
			return
		}
	}
	if len(_datasyncCmkSecretConfig) > 0 {
		if err := assignInputField(input, "CmkSecretConfig", _datasyncCmkSecretConfig); err != nil {
			log.Errorf("invalid --cmk-secret-config: %s", err.Error())
			return
		}
	}
	if len(_datasyncCustomSecretConfig) > 0 {
		if err := assignInputField(input, "CustomSecretConfig", _datasyncCustomSecretConfig); err != nil {
			log.Errorf("invalid --custom-secret-config: %s", err.Error())
			return
		}
	}
	if len(_datasyncSasConfiguration) > 0 {
		if err := assignInputField(input, "SasConfiguration", _datasyncSasConfiguration); err != nil {
			log.Errorf("invalid --sas-configuration: %s", err.Error())
			return
		}
	}
	if len(_datasyncSubdirectory) > 0 {
		input.Subdirectory = aws.String(_datasyncSubdirectory)
	}

	if resp, err := client.UpdateLocationAzureBlob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the following configuration parameters of the Amazon EFS transfer
// location that you're using with DataSync.
//
// For more information, see [Configuring DataSync transfers with Amazon EFS].
//
// [Configuring DataSync transfers with Amazon EFS]: https://docs.aws.amazon.com/datasync/latest/userguide/create-efs-location.html
func datasync_UpdateLocationEfs(cfg aws.Config, client *datasync.Client) {
	input := &datasync.UpdateLocationEfsInput{
		// LocationArn: *string, // Required
	}

	if len(_datasyncLocationArn) > 0 {
		input.LocationArn = aws.String(_datasyncLocationArn)
	}
	if len(_datasyncAccessPointArn) > 0 {
		input.AccessPointArn = aws.String(_datasyncAccessPointArn)
	}
	if len(_datasyncFileSystemAccessRoleArn) > 0 {
		input.FileSystemAccessRoleArn = aws.String(_datasyncFileSystemAccessRoleArn)
	}
	if len(_datasyncInTransitEncryption) > 0 {
		if err := assignInputField(input, "InTransitEncryption", _datasyncInTransitEncryption); err != nil {
			log.Errorf("invalid --in-transit-encryption: %s", err.Error())
			return
		}
	}
	if len(_datasyncSubdirectory) > 0 {
		input.Subdirectory = aws.String(_datasyncSubdirectory)
	}

	if resp, err := client.UpdateLocationEfs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the following configuration parameters of the Amazon FSx for Lustre
// transfer location that you're using with DataSync.
//
// For more information, see [Configuring DataSync transfers with FSx for Lustre].
//
// [Configuring DataSync transfers with FSx for Lustre]: https://docs.aws.amazon.com/datasync/latest/userguide/create-lustre-location.html
func datasync_UpdateLocationFsxLustre(cfg aws.Config, client *datasync.Client) {
	input := &datasync.UpdateLocationFsxLustreInput{
		// LocationArn: *string, // Required
	}

	if len(_datasyncLocationArn) > 0 {
		input.LocationArn = aws.String(_datasyncLocationArn)
	}
	if len(_datasyncSubdirectory) > 0 {
		input.Subdirectory = aws.String(_datasyncSubdirectory)
	}

	if resp, err := client.UpdateLocationFsxLustre(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the following configuration parameters of the Amazon FSx for NetApp
// ONTAP transfer location that you're using with DataSync.
//
// For more information, see [Configuring DataSync transfers with FSx for ONTAP].
//
// [Configuring DataSync transfers with FSx for ONTAP]: https://docs.aws.amazon.com/datasync/latest/userguide/create-ontap-location.html
func datasync_UpdateLocationFsxOntap(cfg aws.Config, client *datasync.Client) {
	input := &datasync.UpdateLocationFsxOntapInput{
		// LocationArn: *string, // Required
	}

	if len(_datasyncLocationArn) > 0 {
		input.LocationArn = aws.String(_datasyncLocationArn)
	}
	if len(_datasyncProtocol) > 0 {
		if err := assignInputField(input, "Protocol", _datasyncProtocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}
	if len(_datasyncSubdirectory) > 0 {
		input.Subdirectory = aws.String(_datasyncSubdirectory)
	}

	if resp, err := client.UpdateLocationFsxOntap(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the following configuration parameters of the Amazon FSx for OpenZFS
// transfer location that you're using with DataSync.
//
// For more information, see [Configuring DataSync transfers with FSx for OpenZFS].
//
// Request parameters related to SMB aren't supported with the
// UpdateLocationFsxOpenZfs operation.
//
// [Configuring DataSync transfers with FSx for OpenZFS]: https://docs.aws.amazon.com/datasync/latest/userguide/create-openzfs-location.html
func datasync_UpdateLocationFsxOpenZfs(cfg aws.Config, client *datasync.Client) {
	input := &datasync.UpdateLocationFsxOpenZfsInput{
		// LocationArn: *string, // Required
	}

	if len(_datasyncLocationArn) > 0 {
		input.LocationArn = aws.String(_datasyncLocationArn)
	}
	if len(_datasyncProtocol) > 0 {
		if err := assignInputField(input, "Protocol", _datasyncProtocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}
	if len(_datasyncSubdirectory) > 0 {
		input.Subdirectory = aws.String(_datasyncSubdirectory)
	}

	if resp, err := client.UpdateLocationFsxOpenZfs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the following configuration parameters of the Amazon FSx for Windows
// File Server transfer location that you're using with DataSync.
//
// For more information, see [Configuring DataSync transfers with FSx for Windows File Server].
//
// [Configuring DataSync transfers with FSx for Windows File Server]: https://docs.aws.amazon.com/datasync/latest/userguide/create-fsx-location.html
func datasync_UpdateLocationFsxWindows(cfg aws.Config, client *datasync.Client) {
	input := &datasync.UpdateLocationFsxWindowsInput{
		// LocationArn: *string, // Required
	}

	if len(_datasyncLocationArn) > 0 {
		input.LocationArn = aws.String(_datasyncLocationArn)
	}
	if len(_datasyncDomain) > 0 {
		input.Domain = aws.String(_datasyncDomain)
	}
	if len(_datasyncPassword) > 0 {
		input.Password = aws.String(_datasyncPassword)
	}
	if len(_datasyncSubdirectory) > 0 {
		input.Subdirectory = aws.String(_datasyncSubdirectory)
	}
	if len(_datasyncUser) > 0 {
		input.User = aws.String(_datasyncUser)
	}

	if resp, err := client.UpdateLocationFsxWindows(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the following configuration parameters of the Hadoop Distributed File
// System (HDFS) transfer location that you're using with DataSync.
//
// For more information, see [Configuring DataSync transfers with an HDFS cluster].
//
// [Configuring DataSync transfers with an HDFS cluster]: https://docs.aws.amazon.com/datasync/latest/userguide/create-hdfs-location.html
func datasync_UpdateLocationHdfs(cfg aws.Config, client *datasync.Client) {
	input := &datasync.UpdateLocationHdfsInput{
		// LocationArn: *string, // Required
	}

	if len(_datasyncLocationArn) > 0 {
		input.LocationArn = aws.String(_datasyncLocationArn)
	}
	if len(_datasyncAgentArns) > 0 {
		input.AgentArns = append([]string(nil), _datasyncAgentArns...)
	}
	if len(_datasyncAuthenticationType) > 0 {
		if err := assignInputField(input, "AuthenticationType", _datasyncAuthenticationType); err != nil {
			log.Errorf("invalid --authentication-type: %s", err.Error())
			return
		}
	}
	if len(_datasyncBlockSize) > 0 {
		if err := assignInputField(input, "BlockSize", _datasyncBlockSize); err != nil {
			log.Errorf("invalid --block-size: %s", err.Error())
			return
		}
	}
	if len(_datasyncKerberosKeytab) > 0 {
		if err := assignInputField(input, "KerberosKeytab", _datasyncKerberosKeytab); err != nil {
			log.Errorf("invalid --kerberos-keytab: %s", err.Error())
			return
		}
	}
	if len(_datasyncKerberosKrb5Conf) > 0 {
		if err := assignInputField(input, "KerberosKrb5Conf", _datasyncKerberosKrb5Conf); err != nil {
			log.Errorf("invalid --kerberos-krb5-conf: %s", err.Error())
			return
		}
	}
	if len(_datasyncKerberosPrincipal) > 0 {
		input.KerberosPrincipal = aws.String(_datasyncKerberosPrincipal)
	}
	if len(_datasyncKmsKeyProviderUri) > 0 {
		input.KmsKeyProviderUri = aws.String(_datasyncKmsKeyProviderUri)
	}
	if len(_datasyncNameNodes) > 0 {
		if err := assignInputField(input, "NameNodes", _datasyncNameNodes); err != nil {
			log.Errorf("invalid --name-nodes: %s", err.Error())
			return
		}
	}
	if len(_datasyncQopConfiguration) > 0 {
		if err := assignInputField(input, "QopConfiguration", _datasyncQopConfiguration); err != nil {
			log.Errorf("invalid --qop-configuration: %s", err.Error())
			return
		}
	}
	if len(_datasyncReplicationFactor) > 0 {
		if err := assignInputField(input, "ReplicationFactor", _datasyncReplicationFactor); err != nil {
			log.Errorf("invalid --replication-factor: %s", err.Error())
			return
		}
	}
	if len(_datasyncSimpleUser) > 0 {
		input.SimpleUser = aws.String(_datasyncSimpleUser)
	}
	if len(_datasyncSubdirectory) > 0 {
		input.Subdirectory = aws.String(_datasyncSubdirectory)
	}

	if resp, err := client.UpdateLocationHdfs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the following configuration parameters of the Network File System
// (NFS) transfer location that you're using with DataSync.
//
// For more information, see [Configuring transfers with an NFS file server].
//
// [Configuring transfers with an NFS file server]: https://docs.aws.amazon.com/datasync/latest/userguide/create-nfs-location.html
func datasync_UpdateLocationNfs(cfg aws.Config, client *datasync.Client) {
	input := &datasync.UpdateLocationNfsInput{
		// LocationArn: *string, // Required
	}

	if len(_datasyncLocationArn) > 0 {
		input.LocationArn = aws.String(_datasyncLocationArn)
	}
	if len(_datasyncMountOptions) > 0 {
		if err := assignInputField(input, "MountOptions", _datasyncMountOptions); err != nil {
			log.Errorf("invalid --mount-options: %s", err.Error())
			return
		}
	}
	if len(_datasyncOnPremConfig) > 0 {
		if err := assignInputField(input, "OnPremConfig", _datasyncOnPremConfig); err != nil {
			log.Errorf("invalid --on-prem-config: %s", err.Error())
			return
		}
	}
	if len(_datasyncServerHostname) > 0 {
		input.ServerHostname = aws.String(_datasyncServerHostname)
	}
	if len(_datasyncSubdirectory) > 0 {
		input.Subdirectory = aws.String(_datasyncSubdirectory)
	}

	if resp, err := client.UpdateLocationNfs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the following configuration parameters of the object storage transfer
// location that you're using with DataSync.
//
// For more information, see [Configuring DataSync transfers with an object storage system].
//
// [Configuring DataSync transfers with an object storage system]: https://docs.aws.amazon.com/datasync/latest/userguide/create-object-location.html
func datasync_UpdateLocationObjectStorage(cfg aws.Config, client *datasync.Client) {
	input := &datasync.UpdateLocationObjectStorageInput{
		// LocationArn: *string, // Required
	}

	if len(_datasyncLocationArn) > 0 {
		input.LocationArn = aws.String(_datasyncLocationArn)
	}
	if len(_datasyncAccessKey) > 0 {
		input.AccessKey = aws.String(_datasyncAccessKey)
	}
	if len(_datasyncAgentArns) > 0 {
		input.AgentArns = append([]string(nil), _datasyncAgentArns...)
	}
	if len(_datasyncCmkSecretConfig) > 0 {
		if err := assignInputField(input, "CmkSecretConfig", _datasyncCmkSecretConfig); err != nil {
			log.Errorf("invalid --cmk-secret-config: %s", err.Error())
			return
		}
	}
	if len(_datasyncCustomSecretConfig) > 0 {
		if err := assignInputField(input, "CustomSecretConfig", _datasyncCustomSecretConfig); err != nil {
			log.Errorf("invalid --custom-secret-config: %s", err.Error())
			return
		}
	}
	if len(_datasyncSecretKey) > 0 {
		input.SecretKey = aws.String(_datasyncSecretKey)
	}
	if len(_datasyncServerCertificate) > 0 {
		if err := assignInputField(input, "ServerCertificate", _datasyncServerCertificate); err != nil {
			log.Errorf("invalid --server-certificate: %s", err.Error())
			return
		}
	}
	if len(_datasyncServerHostname) > 0 {
		input.ServerHostname = aws.String(_datasyncServerHostname)
	}
	if len(_datasyncServerPort) > 0 {
		if err := assignInputField(input, "ServerPort", _datasyncServerPort); err != nil {
			log.Errorf("invalid --server-port: %s", err.Error())
			return
		}
	}
	if len(_datasyncServerProtocol) > 0 {
		if err := assignInputField(input, "ServerProtocol", _datasyncServerProtocol); err != nil {
			log.Errorf("invalid --server-protocol: %s", err.Error())
			return
		}
	}
	if len(_datasyncSubdirectory) > 0 {
		input.Subdirectory = aws.String(_datasyncSubdirectory)
	}

	if resp, err := client.UpdateLocationObjectStorage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the following configuration parameters of the Amazon S3 transfer
// location that you're using with DataSync.
//
// Before you begin, make sure that you read the following topics:
//
// [Storage class considerations with Amazon S3 locations]
//
// [Evaluating S3 request costs when using DataSync]
//
// [Storage class considerations with Amazon S3 locations]: https://docs.aws.amazon.com/datasync/latest/userguide/create-s3-location.html#using-storage-classes
// [Evaluating S3 request costs when using DataSync]: https://docs.aws.amazon.com/datasync/latest/userguide/create-s3-location.html#create-s3-location-s3-requests
func datasync_UpdateLocationS3(cfg aws.Config, client *datasync.Client) {
	input := &datasync.UpdateLocationS3Input{
		// LocationArn: *string, // Required
	}

	if len(_datasyncLocationArn) > 0 {
		input.LocationArn = aws.String(_datasyncLocationArn)
	}
	if len(_datasyncS3Config) > 0 {
		if err := assignInputField(input, "S3Config", _datasyncS3Config); err != nil {
			log.Errorf("invalid --s3-config: %s", err.Error())
			return
		}
	}
	if len(_datasyncS3StorageClass) > 0 {
		if err := assignInputField(input, "S3StorageClass", _datasyncS3StorageClass); err != nil {
			log.Errorf("invalid --s3-storage-class: %s", err.Error())
			return
		}
	}
	if len(_datasyncSubdirectory) > 0 {
		input.Subdirectory = aws.String(_datasyncSubdirectory)
	}

	if resp, err := client.UpdateLocationS3(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the following configuration parameters of the Server Message Block
// (SMB) transfer location that you're using with DataSync.
//
// For more information, see [Configuring DataSync transfers with an SMB file server].
//
// [Configuring DataSync transfers with an SMB file server]: https://docs.aws.amazon.com/datasync/latest/userguide/create-smb-location.html
func datasync_UpdateLocationSmb(cfg aws.Config, client *datasync.Client) {
	input := &datasync.UpdateLocationSmbInput{
		// LocationArn: *string, // Required
	}

	if len(_datasyncLocationArn) > 0 {
		input.LocationArn = aws.String(_datasyncLocationArn)
	}
	if len(_datasyncAgentArns) > 0 {
		input.AgentArns = append([]string(nil), _datasyncAgentArns...)
	}
	if len(_datasyncAuthenticationType) > 0 {
		if err := assignInputField(input, "AuthenticationType", _datasyncAuthenticationType); err != nil {
			log.Errorf("invalid --authentication-type: %s", err.Error())
			return
		}
	}
	if len(_datasyncCmkSecretConfig) > 0 {
		if err := assignInputField(input, "CmkSecretConfig", _datasyncCmkSecretConfig); err != nil {
			log.Errorf("invalid --cmk-secret-config: %s", err.Error())
			return
		}
	}
	if len(_datasyncCustomSecretConfig) > 0 {
		if err := assignInputField(input, "CustomSecretConfig", _datasyncCustomSecretConfig); err != nil {
			log.Errorf("invalid --custom-secret-config: %s", err.Error())
			return
		}
	}
	if len(_datasyncDnsIpAddresses) > 0 {
		input.DnsIpAddresses = append([]string(nil), _datasyncDnsIpAddresses...)
	}
	if len(_datasyncDomain) > 0 {
		input.Domain = aws.String(_datasyncDomain)
	}
	if len(_datasyncKerberosKeytab) > 0 {
		if err := assignInputField(input, "KerberosKeytab", _datasyncKerberosKeytab); err != nil {
			log.Errorf("invalid --kerberos-keytab: %s", err.Error())
			return
		}
	}
	if len(_datasyncKerberosKrb5Conf) > 0 {
		if err := assignInputField(input, "KerberosKrb5Conf", _datasyncKerberosKrb5Conf); err != nil {
			log.Errorf("invalid --kerberos-krb5-conf: %s", err.Error())
			return
		}
	}
	if len(_datasyncKerberosPrincipal) > 0 {
		input.KerberosPrincipal = aws.String(_datasyncKerberosPrincipal)
	}
	if len(_datasyncMountOptions) > 0 {
		if err := assignInputField(input, "MountOptions", _datasyncMountOptions); err != nil {
			log.Errorf("invalid --mount-options: %s", err.Error())
			return
		}
	}
	if len(_datasyncPassword) > 0 {
		input.Password = aws.String(_datasyncPassword)
	}
	if len(_datasyncServerHostname) > 0 {
		input.ServerHostname = aws.String(_datasyncServerHostname)
	}
	if len(_datasyncSubdirectory) > 0 {
		input.Subdirectory = aws.String(_datasyncSubdirectory)
	}
	if len(_datasyncUser) > 0 {
		input.User = aws.String(_datasyncUser)
	}

	if resp, err := client.UpdateLocationSmb(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of a task, which defines where and how DataSync
// transfers your data.
func datasync_UpdateTask(cfg aws.Config, client *datasync.Client) {
	input := &datasync.UpdateTaskInput{
		// TaskArn: *string, // Required
	}

	if len(_datasyncTaskArn) > 0 {
		input.TaskArn = aws.String(_datasyncTaskArn)
	}
	if len(_datasyncCloudWatchLogGroupArn) > 0 {
		input.CloudWatchLogGroupArn = aws.String(_datasyncCloudWatchLogGroupArn)
	}
	if len(_datasyncExcludes) > 0 {
		if err := assignInputField(input, "Excludes", _datasyncExcludes); err != nil {
			log.Errorf("invalid --excludes: %s", err.Error())
			return
		}
	}
	if len(_datasyncIncludes) > 0 {
		if err := assignInputField(input, "Includes", _datasyncIncludes); err != nil {
			log.Errorf("invalid --includes: %s", err.Error())
			return
		}
	}
	if len(_datasyncManifestConfig) > 0 {
		if err := assignInputField(input, "ManifestConfig", _datasyncManifestConfig); err != nil {
			log.Errorf("invalid --manifest-config: %s", err.Error())
			return
		}
	}
	if len(_datasyncName) > 0 {
		input.Name = aws.String(_datasyncName)
	}
	if len(_datasyncOptions) > 0 {
		if err := assignInputField(input, "Options", _datasyncOptions); err != nil {
			log.Errorf("invalid --options: %s", err.Error())
			return
		}
	}
	if len(_datasyncSchedule) > 0 {
		if err := assignInputField(input, "Schedule", _datasyncSchedule); err != nil {
			log.Errorf("invalid --schedule: %s", err.Error())
			return
		}
	}
	if len(_datasyncTaskReportConfig) > 0 {
		if err := assignInputField(input, "TaskReportConfig", _datasyncTaskReportConfig); err != nil {
			log.Errorf("invalid --task-report-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of a running DataSync task execution.
// Currently, the only Option that you can modify with UpdateTaskExecution is [BytesPerSecond],
// which throttles bandwidth for a running or queued task execution.
//
// [BytesPerSecond]: https://docs.aws.amazon.com/datasync/latest/userguide/API_Options.html#DataSync-Type-Options-BytesPerSecond
func datasync_UpdateTaskExecution(cfg aws.Config, client *datasync.Client) {
	input := &datasync.UpdateTaskExecutionInput{
		// Options: *types.Options, // Required
		// TaskExecutionArn: *string, // Required
	}

	if len(_datasyncOptions) > 0 {
		if err := assignInputField(input, "Options", _datasyncOptions); err != nil {
			log.Errorf("invalid --options: %s", err.Error())
			return
		}
	}
	if len(_datasyncTaskExecutionArn) > 0 {
		input.TaskExecutionArn = aws.String(_datasyncTaskExecutionArn)
	}

	if resp, err := client.UpdateTaskExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_datasyncCmd)
	_datasyncCmd.Flags().SortFlags = false

	_datasyncCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_datasyncCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_datasyncCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_datasyncCmd.Flags().StringVarP(&_datasyncAccessKey, "access-key", "", "", "Access Key")
	_datasyncCmd.Flags().StringVarP(&_datasyncAccessPointArn, "access-point-arn", "", "", "Access Point ARN")
	_datasyncCmd.Flags().StringVarP(&_datasyncAccessTier, "access-tier", "", "", "Access Tier")
	_datasyncCmd.Flags().StringVarP(&_datasyncActivationKey, "activation-key", "", "", "Activation Key")
	_datasyncCmd.Flags().StringVarP(&_datasyncAgentArn, "agent-arn", "", "", "Agent ARN")
	_datasyncCmd.Flags().StringSliceVarP(&_datasyncAgentArns, "agent-arns", "", nil, "Agent Arns")
	_datasyncCmd.Flags().StringVarP(&_datasyncAgentName, "agent-name", "", "", "Agent Name")
	_datasyncCmd.Flags().StringVarP(&_datasyncAuthenticationType, "authentication-type", "", "", "Authentication Type")
	_datasyncCmd.Flags().StringVarP(&_datasyncBlobType, "blob-type", "", "", "Blob Type")
	_datasyncCmd.Flags().StringVarP(&_datasyncBlockSize, "block-size", "", "", "Block Size")
	_datasyncCmd.Flags().StringVarP(&_datasyncBucketName, "bucket-name", "", "", "Bucket Name")
	_datasyncCmd.Flags().StringVarP(&_datasyncCloudWatchLogGroupArn, "cloud-watch-log-group-arn", "", "", "Cloud Watch Log Group ARN")
	_datasyncCmd.Flags().StringVarP(&_datasyncCmkSecretConfig, "cmk-secret-config", "", "", "Cmk Secret Config")
	_datasyncCmd.Flags().StringVarP(&_datasyncContainerUrl, "container-url", "", "", "Container URL")
	_datasyncCmd.Flags().StringVarP(&_datasyncCustomSecretConfig, "custom-secret-config", "", "", "Custom Secret Config")
	_datasyncCmd.Flags().StringVarP(&_datasyncDestinationLocationArn, "destination-location-arn", "", "", "Destination Location ARN")
	_datasyncCmd.Flags().StringSliceVarP(&_datasyncDnsIpAddresses, "dns-ip-addresses", "", nil, "DNS IP Addresses")
	_datasyncCmd.Flags().StringVarP(&_datasyncDomain, "domain", "", "", "Domain")
	_datasyncCmd.Flags().StringVarP(&_datasyncEc2Config, "ec2-config", "", "", "EC2 Config")
	_datasyncCmd.Flags().StringVarP(&_datasyncEfsFilesystemArn, "efs-filesystem-arn", "", "", "EFS Filesystem ARN")
	_datasyncCmd.Flags().StringVarP(&_datasyncExcludes, "excludes", "", "", "Excludes")
	_datasyncCmd.Flags().StringVarP(&_datasyncFileSystemAccessRoleArn, "file-system-access-role-arn", "", "", "File System Access Role ARN")
	_datasyncCmd.Flags().StringVarP(&_datasyncFilters, "filters", "", "", "Filters")
	_datasyncCmd.Flags().StringVarP(&_datasyncFsxFilesystemArn, "fsx-filesystem-arn", "", "", "Fsx Filesystem ARN")
	_datasyncCmd.Flags().StringVarP(&_datasyncInTransitEncryption, "in-transit-encryption", "", "", "In Transit Encryption")
	_datasyncCmd.Flags().StringVarP(&_datasyncIncludes, "includes", "", "", "Includes")
	_datasyncCmd.Flags().StringVarP(&_datasyncKerberosKeytab, "kerberos-keytab", "", "", "Kerberos Keytab")
	_datasyncCmd.Flags().StringVarP(&_datasyncKerberosKrb5Conf, "kerberos-krb5-conf", "", "", "Kerberos Krb5 Conf")
	_datasyncCmd.Flags().StringVarP(&_datasyncKerberosPrincipal, "kerberos-principal", "", "", "Kerberos Principal")
	_datasyncCmd.Flags().StringSliceVarP(&_datasyncKeys, "keys", "", nil, "Keys")
	_datasyncCmd.Flags().StringVarP(&_datasyncKmsKeyProviderUri, "kms-key-provider-uri", "", "", "KMS Key Provider URI")
	_datasyncCmd.Flags().StringVarP(&_datasyncLocationArn, "location-arn", "", "", "Location ARN")
	_datasyncCmd.Flags().StringVarP(&_datasyncManifestConfig, "manifest-config", "", "", "Manifest Config")
	_datasyncCmd.Flags().StringVarP(&_datasyncMaxResults, "max-results", "", "", "Max Results")
	_datasyncCmd.Flags().StringVarP(&_datasyncMountOptions, "mount-options", "", "", "Mount Options")
	_datasyncCmd.Flags().StringVarP(&_datasyncName, "name", "", "", "Name")
	_datasyncCmd.Flags().StringVarP(&_datasyncNameNodes, "name-nodes", "", "", "Name Nodes")
	_datasyncCmd.Flags().StringVarP(&_datasyncNextToken, "next-token", "", "", "Next Token")
	_datasyncCmd.Flags().StringVarP(&_datasyncOnPremConfig, "on-prem-config", "", "", "On Prem Config")
	_datasyncCmd.Flags().StringVarP(&_datasyncOptions, "options", "", "", "Options")
	_datasyncCmd.Flags().StringVarP(&_datasyncOverrideOptions, "override-options", "", "", "Override Options")
	_datasyncCmd.Flags().StringVarP(&_datasyncPassword, "password", "", "", "Password")
	_datasyncCmd.Flags().StringVarP(&_datasyncProtocol, "protocol", "", "", "Protocol")
	_datasyncCmd.Flags().StringVarP(&_datasyncQopConfiguration, "qop-configuration", "", "", "Qop Configuration")
	_datasyncCmd.Flags().StringVarP(&_datasyncReplicationFactor, "replication-factor", "", "", "Replication Factor")
	_datasyncCmd.Flags().StringVarP(&_datasyncResourceArn, "resource-arn", "", "", "Resource ARN")
	_datasyncCmd.Flags().StringVarP(&_datasyncS3BucketArn, "s3-bucket-arn", "", "", "S3 Bucket ARN")
	_datasyncCmd.Flags().StringVarP(&_datasyncS3Config, "s3-config", "", "", "S3 Config")
	_datasyncCmd.Flags().StringVarP(&_datasyncS3StorageClass, "s3-storage-class", "", "", "S3 Storage Class")
	_datasyncCmd.Flags().StringVarP(&_datasyncSasConfiguration, "sas-configuration", "", "", "Sas Configuration")
	_datasyncCmd.Flags().StringVarP(&_datasyncSchedule, "schedule", "", "", "Schedule")
	_datasyncCmd.Flags().StringVarP(&_datasyncSecretKey, "secret-key", "", "", "Secret Key")
	_datasyncCmd.Flags().StringSliceVarP(&_datasyncSecurityGroupArns, "security-group-arns", "", nil, "Security Group Arns")
	_datasyncCmd.Flags().StringVarP(&_datasyncServerCertificate, "server-certificate", "", "", "Server Certificate")
	_datasyncCmd.Flags().StringVarP(&_datasyncServerHostname, "server-hostname", "", "", "Server Hostname")
	_datasyncCmd.Flags().StringVarP(&_datasyncServerPort, "server-port", "", "", "Server Port")
	_datasyncCmd.Flags().StringVarP(&_datasyncServerProtocol, "server-protocol", "", "", "Server Protocol")
	_datasyncCmd.Flags().StringVarP(&_datasyncSimpleUser, "simple-user", "", "", "Simple User")
	_datasyncCmd.Flags().StringVarP(&_datasyncSourceLocationArn, "source-location-arn", "", "", "Source Location ARN")
	_datasyncCmd.Flags().StringVarP(&_datasyncStorageVirtualMachineArn, "storage-virtual-machine-arn", "", "", "Storage Virtual Machine ARN")
	_datasyncCmd.Flags().StringVarP(&_datasyncSubdirectory, "subdirectory", "", "", "Subdirectory")
	_datasyncCmd.Flags().StringSliceVarP(&_datasyncSubnetArns, "subnet-arns", "", nil, "Subnet Arns")
	_datasyncCmd.Flags().StringVarP(&_datasyncTags, "tags", "", "", "Tags")
	_datasyncCmd.Flags().StringVarP(&_datasyncTaskArn, "task-arn", "", "", "Task ARN")
	_datasyncCmd.Flags().StringVarP(&_datasyncTaskExecutionArn, "task-execution-arn", "", "", "Task Execution ARN")
	_datasyncCmd.Flags().StringVarP(&_datasyncTaskMode, "task-mode", "", "", "Task Mode")
	_datasyncCmd.Flags().StringVarP(&_datasyncTaskReportConfig, "task-report-config", "", "", "Task Report Config")
	_datasyncCmd.Flags().StringVarP(&_datasyncUser, "user", "", "", "User")
	_datasyncCmd.Flags().StringVarP(&_datasyncVpcEndpointId, "vpc-endpoint-id", "", "", "VPC Endpoint ID")

	_datasyncCmd.Flags().BoolVarP(&_datasyncCancelTaskExecution, "cancel-task-execution", "", false, "Cancel Task Execution")
	_datasyncCmd.Flags().BoolVarP(&_datasyncCreateAgent, "create-agent", "", false, "Create Agent")
	_datasyncCmd.Flags().BoolVarP(&_datasyncCreateLocationAzureBlob, "create-location-azure-blob", "", false, "Create Location Azure Blob")
	_datasyncCmd.Flags().BoolVarP(&_datasyncCreateLocationEfs, "create-location-efs", "", false, "Create Location EFS")
	_datasyncCmd.Flags().BoolVarP(&_datasyncCreateLocationFsxLustre, "create-location-fsx-lustre", "", false, "Create Location Fsx Lustre")
	_datasyncCmd.Flags().BoolVarP(&_datasyncCreateLocationFsxOntap, "create-location-fsx-ontap", "", false, "Create Location Fsx Ontap")
	_datasyncCmd.Flags().BoolVarP(&_datasyncCreateLocationFsxOpenZfs, "create-location-fsx-open-zfs", "", false, "Create Location Fsx Open Zfs")
	_datasyncCmd.Flags().BoolVarP(&_datasyncCreateLocationFsxWindows, "create-location-fsx-windows", "", false, "Create Location Fsx Windows")
	_datasyncCmd.Flags().BoolVarP(&_datasyncCreateLocationHdfs, "create-location-hdfs", "", false, "Create Location Hdfs")
	_datasyncCmd.Flags().BoolVarP(&_datasyncCreateLocationNfs, "create-location-nfs", "", false, "Create Location Nfs")
	_datasyncCmd.Flags().BoolVarP(&_datasyncCreateLocationObjectStorage, "create-location-object-storage", "", false, "Create Location Object Storage")
	_datasyncCmd.Flags().BoolVarP(&_datasyncCreateLocationS3, "create-location-s3", "", false, "Create Location S3")
	_datasyncCmd.Flags().BoolVarP(&_datasyncCreateLocationSmb, "create-location-smb", "", false, "Create Location Smb")
	_datasyncCmd.Flags().BoolVarP(&_datasyncCreateTask, "create-task", "", false, "Create Task")
	_datasyncCmd.Flags().BoolVarP(&_datasyncDeleteAgent, "delete-agent", "", false, "Delete Agent")
	_datasyncCmd.Flags().BoolVarP(&_datasyncDeleteLocation, "delete-location", "", false, "Delete Location")
	_datasyncCmd.Flags().BoolVarP(&_datasyncDeleteTask, "delete-task", "", false, "Delete Task")
	_datasyncCmd.Flags().BoolVarP(&_datasyncDescribeAgent, "describe-agent", "", false, "Describe Agent")
	_datasyncCmd.Flags().BoolVarP(&_datasyncDescribeLocationAzureBlob, "describe-location-azure-blob", "", false, "Describe Location Azure Blob")
	_datasyncCmd.Flags().BoolVarP(&_datasyncDescribeLocationEfs, "describe-location-efs", "", false, "Describe Location EFS")
	_datasyncCmd.Flags().BoolVarP(&_datasyncDescribeLocationFsxLustre, "describe-location-fsx-lustre", "", false, "Describe Location Fsx Lustre")
	_datasyncCmd.Flags().BoolVarP(&_datasyncDescribeLocationFsxOntap, "describe-location-fsx-ontap", "", false, "Describe Location Fsx Ontap")
	_datasyncCmd.Flags().BoolVarP(&_datasyncDescribeLocationFsxOpenZfs, "describe-location-fsx-open-zfs", "", false, "Describe Location Fsx Open Zfs")
	_datasyncCmd.Flags().BoolVarP(&_datasyncDescribeLocationFsxWindows, "describe-location-fsx-windows", "", false, "Describe Location Fsx Windows")
	_datasyncCmd.Flags().BoolVarP(&_datasyncDescribeLocationHdfs, "describe-location-hdfs", "", false, "Describe Location Hdfs")
	_datasyncCmd.Flags().BoolVarP(&_datasyncDescribeLocationNfs, "describe-location-nfs", "", false, "Describe Location Nfs")
	_datasyncCmd.Flags().BoolVarP(&_datasyncDescribeLocationObjectStorage, "describe-location-object-storage", "", false, "Describe Location Object Storage")
	_datasyncCmd.Flags().BoolVarP(&_datasyncDescribeLocationS3, "describe-location-s3", "", false, "Describe Location S3")
	_datasyncCmd.Flags().BoolVarP(&_datasyncDescribeLocationSmb, "describe-location-smb", "", false, "Describe Location Smb")
	_datasyncCmd.Flags().BoolVarP(&_datasyncDescribeTask, "describe-task", "", false, "Describe Task")
	_datasyncCmd.Flags().BoolVarP(&_datasyncDescribeTaskExecution, "describe-task-execution", "", false, "Describe Task Execution")
	_datasyncCmd.Flags().BoolVarP(&_datasyncListAgents, "list-agents", "", false, "List Agents")
	_datasyncCmd.Flags().BoolVarP(&_datasyncListLocations, "list-locations", "", false, "List Locations")
	_datasyncCmd.Flags().BoolVarP(&_datasyncListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_datasyncCmd.Flags().BoolVarP(&_datasyncListTaskExecutions, "list-task-executions", "", false, "List Task Executions")
	_datasyncCmd.Flags().BoolVarP(&_datasyncListTasks, "list-tasks", "", false, "List Tasks")
	_datasyncCmd.Flags().BoolVarP(&_datasyncStartTaskExecution, "start-task-execution", "", false, "Start Task Execution")
	_datasyncCmd.Flags().BoolVarP(&_datasyncTagResource, "tag-resource", "", false, "Tag Resource")
	_datasyncCmd.Flags().BoolVarP(&_datasyncUntagResource, "untag-resource", "", false, "Untag Resource")
	_datasyncCmd.Flags().BoolVarP(&_datasyncUpdateAgent, "update-agent", "", false, "Update Agent")
	_datasyncCmd.Flags().BoolVarP(&_datasyncUpdateLocationAzureBlob, "update-location-azure-blob", "", false, "Update Location Azure Blob")
	_datasyncCmd.Flags().BoolVarP(&_datasyncUpdateLocationEfs, "update-location-efs", "", false, "Update Location EFS")
	_datasyncCmd.Flags().BoolVarP(&_datasyncUpdateLocationFsxLustre, "update-location-fsx-lustre", "", false, "Update Location Fsx Lustre")
	_datasyncCmd.Flags().BoolVarP(&_datasyncUpdateLocationFsxOntap, "update-location-fsx-ontap", "", false, "Update Location Fsx Ontap")
	_datasyncCmd.Flags().BoolVarP(&_datasyncUpdateLocationFsxOpenZfs, "update-location-fsx-open-zfs", "", false, "Update Location Fsx Open Zfs")
	_datasyncCmd.Flags().BoolVarP(&_datasyncUpdateLocationFsxWindows, "update-location-fsx-windows", "", false, "Update Location Fsx Windows")
	_datasyncCmd.Flags().BoolVarP(&_datasyncUpdateLocationHdfs, "update-location-hdfs", "", false, "Update Location Hdfs")
	_datasyncCmd.Flags().BoolVarP(&_datasyncUpdateLocationNfs, "update-location-nfs", "", false, "Update Location Nfs")
	_datasyncCmd.Flags().BoolVarP(&_datasyncUpdateLocationObjectStorage, "update-location-object-storage", "", false, "Update Location Object Storage")
	_datasyncCmd.Flags().BoolVarP(&_datasyncUpdateLocationS3, "update-location-s3", "", false, "Update Location S3")
	_datasyncCmd.Flags().BoolVarP(&_datasyncUpdateLocationSmb, "update-location-smb", "", false, "Update Location Smb")
	_datasyncCmd.Flags().BoolVarP(&_datasyncUpdateTask, "update-task", "", false, "Update Task")
	_datasyncCmd.Flags().BoolVarP(&_datasyncUpdateTaskExecution, "update-task-execution", "", false, "Update Task Execution")

}
